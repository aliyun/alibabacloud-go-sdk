// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 应用使用时长上报
//
// @param tmpReq - AppUseTimeReportRequest
//
// @param headers - AppUseTimeReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppUseTimeReportResponse
func (client *Client) AppUseTimeReportWithContext(ctx context.Context, tmpReq *AppUseTimeReportRequest, headers *AppUseTimeReportHeaders, runtime *dara.RuntimeOptions) (_result *AppUseTimeReportResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AppUseTimeReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
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
		Action:      dara.String("AppUseTimeReport"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/vip/use/time/report"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AppUseTimeReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 三方领取回调接口
//
// @param tmpReq - CallBackThirdRightSendPlanRequest
//
// @param headers - CallBackThirdRightSendPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CallBackThirdRightSendPlanResponse
func (client *Client) CallBackThirdRightSendPlanWithContext(ctx context.Context, tmpReq *CallBackThirdRightSendPlanRequest, headers *CallBackThirdRightSendPlanHeaders, runtime *dara.RuntimeOptions) (_result *CallBackThirdRightSendPlanResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CallBackThirdRightSendPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtendInfo) {
		request.ExtendInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendInfo, dara.String("ExtendInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizGroup) {
		query["BizGroup"] = request.BizGroup
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CardType) {
		query["CardType"] = request.CardType
	}

	if !dara.IsNil(request.ErrorMsg) {
		query["ErrorMsg"] = request.ErrorMsg
	}

	if !dara.IsNil(request.ExtendInfoShrink) {
		query["ExtendInfo"] = request.ExtendInfoShrink
	}

	if !dara.IsNil(request.GenieOpenId) {
		query["GenieOpenId"] = request.GenieOpenId
	}

	if !dara.IsNil(request.ReceiveStatus) {
		query["ReceiveStatus"] = request.ReceiveStatus
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	if !dara.IsNil(request.SupplierId) {
		query["SupplierId"] = request.SupplierId
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
		Action:      dara.String("CallBackThirdRightSendPlan"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/business/CallBackThirdRightSendPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CallBackThirdRightSendPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商业化移动屏三方app领卡校验
//
// @param tmpReq - CheckThirdRightSendPlanRequest
//
// @param headers - CheckThirdRightSendPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckThirdRightSendPlanResponse
func (client *Client) CheckThirdRightSendPlanWithContext(ctx context.Context, tmpReq *CheckThirdRightSendPlanRequest, headers *CheckThirdRightSendPlanHeaders, runtime *dara.RuntimeOptions) (_result *CheckThirdRightSendPlanResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CheckThirdRightSendPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtendInfo) {
		request.ExtendInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendInfo, dara.String("ExtendInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizGroup) {
		query["BizGroup"] = request.BizGroup
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ExtendInfoShrink) {
		query["ExtendInfo"] = request.ExtendInfoShrink
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	if !dara.IsNil(request.SupplierId) {
		query["SupplierId"] = request.SupplierId
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
		Action:      dara.String("CheckThirdRightSendPlan"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/business/CheckThirdRightSendPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckThirdRightSendPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建提醒
//
// @param tmpReq - CreateReminderRequest
//
// @param headers - CreateReminderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReminderResponse
func (client *Client) CreateReminderWithContext(ctx context.Context, tmpReq *CreateReminderRequest, headers *CreateReminderHeaders, runtime *dara.RuntimeOptions) (_result *CreateReminderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateReminderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
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
		Action:      dara.String("CreateReminder"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/reminder/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReminderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除提醒
//
// @param tmpReq - DeleteReminderRequest
//
// @param headers - DeleteReminderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReminderResponse
func (client *Client) DeleteReminderWithContext(ctx context.Context, tmpReq *DeleteReminderRequest, headers *DeleteReminderHeaders, runtime *dara.RuntimeOptions) (_result *DeleteReminderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteReminderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

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
		Action:      dara.String("DeleteReminder"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/reminder/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReminderResponse{}
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
// @param tmpReq - GetAccountForAppRequest
//
// @param headers - GetAccountForAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountForAppResponse
func (client *Client) GetAccountForAppWithContext(ctx context.Context, tmpReq *GetAccountForAppRequest, headers *GetAccountForAppHeaders, runtime *dara.RuntimeOptions) (_result *GetAccountForAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetAccountForAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

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
		Action:      dara.String("GetAccountForApp"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/vip/account/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountForAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用配置
//
// @param tmpReq - GetBusAppConfigRequest
//
// @param headers - GetBusAppConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusAppConfigResponse
func (client *Client) GetBusAppConfigWithContext(ctx context.Context, tmpReq *GetBusAppConfigRequest, headers *GetBusAppConfigHeaders, runtime *dara.RuntimeOptions) (_result *GetBusAppConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetBusAppConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

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
		Action:      dara.String("GetBusAppConfig"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/app/config/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBusAppConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户手机号获取
//
// @param tmpReq - GetPhoneNumberRequest
//
// @param headers - GetPhoneNumberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhoneNumberResponse
func (client *Client) GetPhoneNumberWithContext(ctx context.Context, tmpReq *GetPhoneNumberRequest, headers *GetPhoneNumberHeaders, runtime *dara.RuntimeOptions) (_result *GetPhoneNumberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetPhoneNumberShrinkRequest{}
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
		Action:      dara.String("GetPhoneNumber"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/profile/phoneNumber"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhoneNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询提醒
//
// @param tmpReq - GetReminderRequest
//
// @param headers - GetReminderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReminderResponse
func (client *Client) GetReminderWithContext(ctx context.Context, tmpReq *GetReminderRequest, headers *GetReminderHeaders, runtime *dara.RuntimeOptions) (_result *GetReminderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetReminderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

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
		Action:      dara.String("GetReminder"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/reminder/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReminderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询提醒列表
//
// @param tmpReq - ListRemindersRequest
//
// @param headers - ListRemindersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemindersResponse
func (client *Client) ListRemindersWithContext(ctx context.Context, tmpReq *ListRemindersRequest, headers *ListRemindersHeaders, runtime *dara.RuntimeOptions) (_result *ListRemindersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListRemindersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

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
		Action:      dara.String("ListReminders"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/reminder/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRemindersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拉取收银台
//
// @param tmpReq - PullCashierRequest
//
// @param headers - PullCashierHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PullCashierResponse
func (client *Client) PullCashierWithContext(ctx context.Context, tmpReq *PullCashierRequest, headers *PullCashierHeaders, runtime *dara.RuntimeOptions) (_result *PullCashierResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PullCashierShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

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
		Action:      dara.String("PullCashier"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/pull/cashier/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PullCashierResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 猫精系统消息推送
//
// @param tmpReq - PushNotificationsRequest
//
// @param headers - PushNotificationsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushNotificationsResponse
func (client *Client) PushNotificationsWithContext(ctx context.Context, tmpReq *PushNotificationsRequest, headers *PushNotificationsHeaders, runtime *dara.RuntimeOptions) (_result *PushNotificationsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PushNotificationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotificationUnicastRequest) {
		request.NotificationUnicastRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationUnicastRequest, dara.String("NotificationUnicastRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantInfo) {
		request.TenantInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantInfo, dara.String("TenantInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NotificationUnicastRequestShrink) {
		body["NotificationUnicastRequest"] = request.NotificationUnicastRequestShrink
	}

	if !dara.IsNil(request.TenantInfoShrink) {
		body["TenantInfo"] = request.TenantInfoShrink
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
		Action:      dara.String("PushNotifications"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/notifications"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("none"),
	}
	_result = &PushNotificationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息推送服务（普通版）
//
// @param tmpReq - SendNotificationsRequest
//
// @param headers - SendNotificationsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendNotificationsResponse
func (client *Client) SendNotificationsWithContext(ctx context.Context, tmpReq *SendNotificationsRequest, headers *SendNotificationsHeaders, runtime *dara.RuntimeOptions) (_result *SendNotificationsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SendNotificationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NotificationUnicastRequest) {
		request.NotificationUnicastRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationUnicastRequest, dara.String("NotificationUnicastRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TenantInfo) {
		request.TenantInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantInfo, dara.String("TenantInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.NotificationUnicastRequestShrink) {
		body["NotificationUnicastRequest"] = request.NotificationUnicastRequestShrink
	}

	if !dara.IsNil(request.TenantInfoShrink) {
		body["TenantInfo"] = request.TenantInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
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
		Action:      dara.String("SendNotifications"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/general/notifications"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("none"),
	}
	_result = &SendNotificationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 三方即时信息数据变更事件推送
//
// @param request - ThirdImmediateMsgPushRequest
//
// @param headers - ThirdImmediateMsgPushHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ThirdImmediateMsgPushResponse
func (client *Client) ThirdImmediateMsgPushWithContext(ctx context.Context, request *ThirdImmediateMsgPushRequest, headers *ThirdImmediateMsgPushHeaders, runtime *dara.RuntimeOptions) (_result *ThirdImmediateMsgPushResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ChangeDetail) {
		query["ChangeDetail"] = request.ChangeDetail
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PsgIds) {
		query["PsgIds"] = request.PsgIds
	}

	if !dara.IsNil(request.TrafficChangeType) {
		query["TrafficChangeType"] = request.TrafficChangeType
	}

	if !dara.IsNil(request.TrafficChangeTypeDesc) {
		query["TrafficChangeTypeDesc"] = request.TrafficChangeTypeDesc
	}

	if !dara.IsNil(request.TrafficJourneyIds) {
		query["TrafficJourneyIds"] = request.TrafficJourneyIds
	}

	if !dara.IsNil(request.TrafficSubOrderIds) {
		query["TrafficSubOrderIds"] = request.TrafficSubOrderIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
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
		Action:      dara.String("ThirdImmediateMsgPush"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/thirdImmediateMsgPush"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ThirdImmediateMsgPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新提醒
//
// @param tmpReq - UpdateReminderRequest
//
// @param headers - UpdateReminderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateReminderResponse
func (client *Client) UpdateReminderWithContext(ctx context.Context, tmpReq *UpdateReminderRequest, headers *UpdateReminderHeaders, runtime *dara.RuntimeOptions) (_result *UpdateReminderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateReminderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
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
		Action:      dara.String("UpdateReminder"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/reminder/update"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateReminderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频类应用会员信息上报
//
// @param tmpReq - VideoAppReportRequest
//
// @param headers - VideoAppReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoAppReportResponse
func (client *Client) VideoAppReportWithContext(ctx context.Context, tmpReq *VideoAppReportRequest, headers *VideoAppReportHeaders, runtime *dara.RuntimeOptions) (_result *VideoAppReportResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &VideoAppReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
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
		Action:      dara.String("VideoAppReport"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/vip/use/video/report"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VideoAppReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 猫精应用唤起
//
// @param request - WakeUpAppRequest
//
// @param headers - WakeUpAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WakeUpAppResponse
func (client *Client) WakeUpAppWithContext(ctx context.Context, request *WakeUpAppRequest, headers *WakeUpAppHeaders, runtime *dara.RuntimeOptions) (_result *WakeUpAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsDebug) {
		body["IsDebug"] = request.IsDebug
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.TargetInfo) {
		body["TargetInfo"] = request.TargetInfo
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
		Action:      dara.String("WakeUpApp"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/wakeup"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &WakeUpAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
