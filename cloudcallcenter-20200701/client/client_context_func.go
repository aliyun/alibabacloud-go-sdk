// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// abort campaign
//
// @param request - AbortCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortCampaignResponse
func (client *Client) AbortCampaignWithContext(ctx context.Context, request *AbortCampaignRequest, runtime *dara.RuntimeOptions) (_result *AbortCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// abort cases
//
// @param tmpReq - AbortCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortCasesResponse
func (client *Client) AbortCasesWithContext(ctx context.Context, tmpReq *AbortCasesRequest, runtime *dara.RuntimeOptions) (_result *AbortCasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AbortCasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PhoneNumberList) {
		request.PhoneNumberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneNumberList, dara.String("PhoneNumberList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PhoneNumberListShrink) {
		query["PhoneNumberList"] = request.PhoneNumberListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortCases"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortCasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// create campaign
//
// @param tmpReq - CreateCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCampaignResponse
func (client *Client) CreateCampaignWithContext(ctx context.Context, tmpReq *CreateCampaignRequest, runtime *dara.RuntimeOptions) (_result *CreateCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCampaignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CaseList) {
		request.CaseListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseList, dara.String("CaseList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NumberList) {
		request.NumberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NumberList, dara.String("NumberList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallableTime) {
		query["CallableTime"] = request.CallableTime
	}

	if !dara.IsNil(request.CaseFileKey) {
		query["CaseFileKey"] = request.CaseFileKey
	}

	if !dara.IsNil(request.CaseListShrink) {
		query["CaseList"] = request.CaseListShrink
	}

	if !dara.IsNil(request.ContactFlowId) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecutingUntilTimeout) {
		query["ExecutingUntilTimeout"] = request.ExecutingUntilTimeout
	}

	if !dara.IsNil(request.FlashSmsParameters) {
		query["FlashSmsParameters"] = request.FlashSmsParameters
	}

	if !dara.IsNil(request.InstGroupId) {
		query["InstGroupId"] = request.InstGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxAttemptCount) {
		query["MaxAttemptCount"] = request.MaxAttemptCount
	}

	if !dara.IsNil(request.MinAttemptInterval) {
		query["MinAttemptInterval"] = request.MinAttemptInterval
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NumberListShrink) {
		query["NumberList"] = request.NumberListShrink
	}

	if !dara.IsNil(request.QueueId) {
		query["QueueId"] = request.QueueId
	}

	if !dara.IsNil(request.Simulation) {
		query["Simulation"] = request.Simulation
	}

	if !dara.IsNil(request.SimulationParameters) {
		query["SimulationParameters"] = request.SimulationParameters
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StrategyParameters) {
		query["StrategyParameters"] = request.StrategyParameters
	}

	if !dara.IsNil(request.StrategyType) {
		query["StrategyType"] = request.StrategyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateChatMediaUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatMediaUrlResponse
func (client *Client) CreateChatMediaUrlWithContext(ctx context.Context, request *CreateChatMediaUrlRequest, runtime *dara.RuntimeOptions) (_result *CreateChatMediaUrlResponse, _err error) {
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

	if !dara.IsNil(request.MimeType) {
		body["MimeType"] = request.MimeType
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatMediaUrl"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatMediaUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateCorpNumberGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCorpNumberGroupResponse
func (client *Client) CreateCorpNumberGroupWithContext(ctx context.Context, request *CreateCorpNumberGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateCorpNumberGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCorpNumberGroup"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCorpNumberGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateDemoInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDemoInstanceResponse
func (client *Client) CreateDemoInstanceWithContext(ctx context.Context, request *CreateDemoInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDemoInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutboundCallWhitelist) {
		query["OutboundCallWhitelist"] = request.OutboundCallWhitelist
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDemoInstance"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDemoInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetAccessChannelOfStaging
//
// @param request - GetAccessChannelOfStagingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessChannelOfStagingResponse
func (client *Client) GetAccessChannelOfStagingWithContext(ctx context.Context, request *GetAccessChannelOfStagingRequest, runtime *dara.RuntimeOptions) (_result *GetAccessChannelOfStagingResponse, _err error) {
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
		Action:      dara.String("GetAccessChannelOfStaging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessChannelOfStagingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动
//
// @param request - GetCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCampaignResponse
func (client *Client) GetCampaignWithContext(ctx context.Context, request *GetCampaignRequest, runtime *dara.RuntimeOptions) (_result *GetCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetDocument
//
// @param request - GetDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentResponse
func (client *Client) GetDocumentWithContext(ctx context.Context, request *GetDocumentRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		body["DocumentId"] = request.DocumentId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SchemaId) {
		body["SchemaId"] = request.SchemaId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocument"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动历史报表
//
// @param request - GetHistoricalCampaignReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoricalCampaignReportResponse
func (client *Client) GetHistoricalCampaignReportWithContext(ctx context.Context, request *GetHistoricalCampaignReportRequest, runtime *dara.RuntimeOptions) (_result *GetHistoricalCampaignReportResponse, _err error) {
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
		Action:      dara.String("GetHistoricalCampaignReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHistoricalCampaignReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据aliyunuid获取实例
//
// @param request - GetInstanceIdsByAliyunUidV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceIdsByAliyunUidV2Response
func (client *Client) GetInstanceIdsByAliyunUidV2WithContext(ctx context.Context, request *GetInstanceIdsByAliyunUidV2Request, runtime *dara.RuntimeOptions) (_result *GetInstanceIdsByAliyunUidV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunUid) {
		query["AliyunUid"] = request.AliyunUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceIdsByAliyunUidV2"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceIdsByAliyunUidV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼实时状态
//
// @param request - GetRealtimeCampaignStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealtimeCampaignStatsResponse
func (client *Client) GetRealtimeCampaignStatsWithContext(ctx context.Context, request *GetRealtimeCampaignStatsRequest, runtime *dara.RuntimeOptions) (_result *GetRealtimeCampaignStatsResponse, _err error) {
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
		Action:      dara.String("GetRealtimeCampaignStats"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRealtimeCampaignStatsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ImportAdminsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportAdminsResponse
func (client *Client) ImportAdminsWithContext(ctx context.Context, request *ImportAdminsRequest, runtime *dara.RuntimeOptions) (_result *ImportAdminsResponse, _err error) {
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

	if !dara.IsNil(request.RamIdList) {
		query["RamIdList"] = request.RamIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportAdmins"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportAdminsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - IssueSoftphoneCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IssueSoftphoneCommandResponse
func (client *Client) IssueSoftphoneCommandWithContext(ctx context.Context, request *IssueSoftphoneCommandRequest, runtime *dara.RuntimeOptions) (_result *IssueSoftphoneCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IssueSoftphoneCommand"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IssueSoftphoneCommandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼呼叫记录
//
// @param request - ListAttemptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAttemptsResponse
func (client *Client) ListAttemptsWithContext(ctx context.Context, request *ListAttemptsRequest, runtime *dara.RuntimeOptions) (_result *ListAttemptsResponse, _err error) {
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
		Action:      dara.String("ListAttempts"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAttemptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取预测式外呼活动趋势报表
//
// @param request - ListCampaignTrendingReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCampaignTrendingReportResponse
func (client *Client) ListCampaignTrendingReportWithContext(ctx context.Context, request *ListCampaignTrendingReportRequest, runtime *dara.RuntimeOptions) (_result *ListCampaignTrendingReportResponse, _err error) {
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
		Action:      dara.String("ListCampaignTrendingReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCampaignTrendingReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询预测式外呼列表
//
// @param request - ListCampaignsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCampaignsResponse
func (client *Client) ListCampaignsWithContext(ctx context.Context, request *ListCampaignsRequest, runtime *dara.RuntimeOptions) (_result *ListCampaignsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActualStartTimeFrom) {
		query["ActualStartTimeFrom"] = request.ActualStartTimeFrom
	}

	if !dara.IsNil(request.ActualStartTimeTo) {
		query["ActualStartTimeTo"] = request.ActualStartTimeTo
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlanedStartTimeFrom) {
		query["PlanedStartTimeFrom"] = request.PlanedStartTimeFrom
	}

	if !dara.IsNil(request.PlanedStartTimeTo) {
		query["PlanedStartTimeTo"] = request.PlanedStartTimeTo
	}

	if !dara.IsNil(request.QueueId) {
		query["QueueId"] = request.QueueId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCampaigns"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCampaignsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// list case
//
// @param request - ListCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCasesResponse
func (client *Client) ListCasesWithContext(ctx context.Context, request *ListCasesRequest, runtime *dara.RuntimeOptions) (_result *ListCasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCases"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - ListFlashSmsSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsSettingsResponse
func (client *Client) ListFlashSmsSettingsWithContext(ctx context.Context, tmpReq *ListFlashSmsSettingsRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListFlashSmsSettingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SkillGroupIdList) {
		request.SkillGroupIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SkillGroupIdList, dara.String("SkillGroupIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillGroupIdListShrink) {
		query["SkillGroupIdList"] = request.SkillGroupIdListShrink
	}

	if !dara.IsNil(request.SkillGroupName) {
		query["SkillGroupName"] = request.SkillGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlashSmsSettings"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsSettingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询坐席技能组报表
//
// @param request - ListHistoricalAgentSkillGroupReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHistoricalAgentSkillGroupReportResponse
func (client *Client) ListHistoricalAgentSkillGroupReportWithContext(ctx context.Context, request *ListHistoricalAgentSkillGroupReportRequest, runtime *dara.RuntimeOptions) (_result *ListHistoricalAgentSkillGroupReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillGroupIdList) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentIdList) {
		body["AgentIdList"] = request.AgentIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHistoricalAgentSkillGroupReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHistoricalAgentSkillGroupReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询坐席技能组报表
//
// @param request - ListIntervalAgentSkillGroupReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntervalAgentSkillGroupReportResponse
func (client *Client) ListIntervalAgentSkillGroupReportWithContext(ctx context.Context, request *ListIntervalAgentSkillGroupReportRequest, runtime *dara.RuntimeOptions) (_result *ListIntervalAgentSkillGroupReportResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.ShowDefaultIfEmpty) {
		query["ShowDefaultIfEmpty"] = request.ShowDefaultIfEmpty
	}

	if !dara.IsNil(request.SkillGroupId) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntervalAgentSkillGroupReport"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntervalAgentSkillGroupReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 录音查询
//
// @param request - ListMonoRecordingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMonoRecordingsResponse
func (client *Client) ListMonoRecordingsWithContext(ctx context.Context, request *ListMonoRecordingsRequest, runtime *dara.RuntimeOptions) (_result *ListMonoRecordingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMonoRecordings"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMonoRecordingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停预测式外呼
//
// @param request - PauseCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseCampaignResponse
func (client *Client) PauseCampaignWithContext(ctx context.Context, request *PauseCampaignRequest, runtime *dara.RuntimeOptions) (_result *PauseCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ProcessAliMeCallbackOfStaging
//
// @param request - ProcessAliMeCallbackOfStagingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProcessAliMeCallbackOfStagingResponse
func (client *Client) ProcessAliMeCallbackOfStagingWithContext(ctx context.Context, request *ProcessAliMeCallbackOfStagingRequest, runtime *dara.RuntimeOptions) (_result *ProcessAliMeCallbackOfStagingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProcessAliMeCallbackOfStaging"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProcessAliMeCallbackOfStagingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ProcessCustomIMCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProcessCustomIMCallbackResponse
func (client *Client) ProcessCustomIMCallbackWithContext(ctx context.Context, request *ProcessCustomIMCallbackRequest, runtime *dara.RuntimeOptions) (_result *ProcessCustomIMCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessChannelId) {
		body["AccessChannelId"] = request.AccessChannelId
	}

	if !dara.IsNil(request.ConversationId) {
		body["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MessageContent) {
		body["MessageContent"] = request.MessageContent
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SenderAvatarMediaId) {
		body["SenderAvatarMediaId"] = request.SenderAvatarMediaId
	}

	if !dara.IsNil(request.SenderId) {
		body["SenderId"] = request.SenderId
	}

	if !dara.IsNil(request.SenderName) {
		body["SenderName"] = request.SenderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProcessCustomIMCallback"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProcessCustomIMCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ccc migration
//
// @param request - ReplaceMigrationAvailableNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceMigrationAvailableNumbersResponse
func (client *Client) ReplaceMigrationAvailableNumbersWithContext(ctx context.Context, request *ReplaceMigrationAvailableNumbersRequest, runtime *dara.RuntimeOptions) (_result *ReplaceMigrationAvailableNumbersResponse, _err error) {
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

	if !dara.IsNil(request.OperatorName) {
		query["OperatorName"] = request.OperatorName
	}

	if !dara.IsNil(request.OperatorUid) {
		query["OperatorUid"] = request.OperatorUid
	}

	if !dara.IsNil(request.V1Numbers) {
		query["V1Numbers"] = request.V1Numbers
	}

	if !dara.IsNil(request.V2Numbers) {
		query["V2Numbers"] = request.V2Numbers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceMigrationAvailableNumbers"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceMigrationAvailableNumbersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// resume campaign
//
// @param request - ResumeCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeCampaignResponse
func (client *Client) ResumeCampaignWithContext(ctx context.Context, request *ResumeCampaignRequest, runtime *dara.RuntimeOptions) (_result *ResumeCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveRTCStatsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveRTCStatsV2Response
func (client *Client) SaveRTCStatsV2WithContext(ctx context.Context, request *SaveRTCStatsV2Request, runtime *dara.RuntimeOptions) (_result *SaveRTCStatsV2Response, _err error) {
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

	if !dara.IsNil(request.GeneralInfo) {
		query["GeneralInfo"] = request.GeneralInfo
	}

	if !dara.IsNil(request.GoogAddress) {
		query["GoogAddress"] = request.GoogAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ReceiverReport) {
		query["ReceiverReport"] = request.ReceiverReport
	}

	if !dara.IsNil(request.SenderReport) {
		query["SenderReport"] = request.SenderReport
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveRTCStatsV2"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveRTCStatsV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// -
//
// @param request - SaveTerminalLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTerminalLogResponse
func (client *Client) SaveTerminalLogWithContext(ctx context.Context, request *SaveTerminalLogRequest, runtime *dara.RuntimeOptions) (_result *SaveTerminalLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MethodName) {
		query["MethodName"] = request.MethodName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UniqueRequestId) {
		query["UniqueRequestId"] = request.UniqueRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTerminalLog"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTerminalLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// -
//
// @param request - SaveWebRtcInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveWebRtcInfoResponse
func (client *Client) SaveWebRtcInfoWithContext(ctx context.Context, request *SaveWebRtcInfoRequest, runtime *dara.RuntimeOptions) (_result *SaveWebRtcInfoResponse, _err error) {
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

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveWebRtcInfo"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveWebRtcInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送消息
//
// @param request - SendNotificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendNotificationResponse
func (client *Client) SendNotificationWithContext(ctx context.Context, request *SendNotificationRequest, runtime *dara.RuntimeOptions) (_result *SendNotificationResponse, _err error) {
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

	if !dara.IsNil(request.MessageBody) {
		query["MessageBody"] = request.MessageBody
	}

	if !dara.IsNil(request.NotificationTarget) {
		query["NotificationTarget"] = request.NotificationTarget
	}

	if !dara.IsNil(request.NotificationType) {
		query["NotificationType"] = request.NotificationType
	}

	if !dara.IsNil(request.ShardingKey) {
		query["ShardingKey"] = request.ShardingKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendNotification"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendNotificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// submit campaign
//
// @param request - SubmitCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCampaignResponse
func (client *Client) SubmitCampaignWithContext(ctx context.Context, request *SubmitCampaignRequest, runtime *dara.RuntimeOptions) (_result *SubmitCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCampaign"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCampaignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除注册设备
//
// @param request - UnregisterDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnregisterDeviceResponse
func (client *Client) UnregisterDeviceWithContext(ctx context.Context, request *UnregisterDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnregisterDeviceResponse, _err error) {
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

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnregisterDevice"),
		Version:     dara.String("2020-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnregisterDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
