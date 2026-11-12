// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("AliGenie")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
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
// Used to report application usage duration on Tmall Genie devices.
//
// @param tmpReq - AppUseTimeReportRequest
//
// @param headers - AppUseTimeReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppUseTimeReportResponse
func (client *Client) AppUseTimeReportWithOptions(tmpReq *AppUseTimeReportRequest, headers *AppUseTimeReportHeaders, runtime *dara.RuntimeOptions) (_result *AppUseTimeReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used to report application usage duration on Tmall Genie devices.
//
// @param request - AppUseTimeReportRequest
//
// @return AppUseTimeReportResponse
func (client *Client) AppUseTimeReport(request *AppUseTimeReportRequest) (_result *AppUseTimeReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AppUseTimeReportHeaders{}
	_result = &AppUseTimeReportResponse{}
	_body, _err := client.AppUseTimeReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Third-party membership redemption callback API.
//
// @param tmpReq - CallBackThirdRightSendPlanRequest
//
// @param headers - CallBackThirdRightSendPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CallBackThirdRightSendPlanResponse
func (client *Client) CallBackThirdRightSendPlanWithOptions(tmpReq *CallBackThirdRightSendPlanRequest, headers *CallBackThirdRightSendPlanHeaders, runtime *dara.RuntimeOptions) (_result *CallBackThirdRightSendPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Third-party membership redemption callback API.
//
// @param request - CallBackThirdRightSendPlanRequest
//
// @return CallBackThirdRightSendPlanResponse
func (client *Client) CallBackThirdRightSendPlan(request *CallBackThirdRightSendPlanRequest) (_result *CallBackThirdRightSendPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CallBackThirdRightSendPlanHeaders{}
	_result = &CallBackThirdRightSendPlanResponse{}
	_body, _err := client.CallBackThirdRightSendPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Commercial mobile screen third-party app card issuance validation.
//
// @param tmpReq - CheckThirdRightSendPlanRequest
//
// @param headers - CheckThirdRightSendPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckThirdRightSendPlanResponse
func (client *Client) CheckThirdRightSendPlanWithOptions(tmpReq *CheckThirdRightSendPlanRequest, headers *CheckThirdRightSendPlanHeaders, runtime *dara.RuntimeOptions) (_result *CheckThirdRightSendPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Commercial mobile screen third-party app card issuance validation.
//
// @param request - CheckThirdRightSendPlanRequest
//
// @return CheckThirdRightSendPlanResponse
func (client *Client) CheckThirdRightSendPlan(request *CheckThirdRightSendPlanRequest) (_result *CheckThirdRightSendPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CheckThirdRightSendPlanHeaders{}
	_result = &CheckThirdRightSendPlanResponse{}
	_body, _err := client.CheckThirdRightSendPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Used to create a reminder on a Tmall Genie device.
//
// @param tmpReq - CreateReminderRequest
//
// @param headers - CreateReminderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReminderResponse
func (client *Client) CreateReminderWithOptions(tmpReq *CreateReminderRequest, headers *CreateReminderHeaders, runtime *dara.RuntimeOptions) (_result *CreateReminderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used to create a reminder on a Tmall Genie device.
//
// @param request - CreateReminderRequest
//
// @return CreateReminderResponse
func (client *Client) CreateReminder(request *CreateReminderRequest) (_result *CreateReminderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateReminderHeaders{}
	_result = &CreateReminderResponse{}
	_body, _err := client.CreateReminderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 携程订单状态变更推送
//
// @param request - CtripOrderPushRequest
//
// @param headers - CtripOrderPushHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CtripOrderPushResponse
func (client *Client) CtripOrderPushWithOptions(request *CtripOrderPushRequest, headers *CtripOrderPushHeaders, runtime *dara.RuntimeOptions) (_result *CtripOrderPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JourneyId) {
		query["JourneyId"] = request.JourneyId
	}

	if !dara.IsNil(request.ScheduleItem) {
		query["ScheduleItem"] = request.ScheduleItem
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
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
		Action:      dara.String("CtripOrderPush"),
		Version:     dara.String("iap_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/iap/ctripOrderPush"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CtripOrderPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 携程订单状态变更推送
//
// @param request - CtripOrderPushRequest
//
// @return CtripOrderPushResponse
func (client *Client) CtripOrderPush(request *CtripOrderPushRequest) (_result *CtripOrderPushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CtripOrderPushHeaders{}
	_result = &CtripOrderPushResponse{}
	_body, _err := client.CtripOrderPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// User deletes a reminder created on a Tmall Genie device.
//
// @param tmpReq - DeleteReminderRequest
//
// @param headers - DeleteReminderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReminderResponse
func (client *Client) DeleteReminderWithOptions(tmpReq *DeleteReminderRequest, headers *DeleteReminderHeaders, runtime *dara.RuntimeOptions) (_result *DeleteReminderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// User deletes a reminder created on a Tmall Genie device.
//
// @param request - DeleteReminderRequest
//
// @return DeleteReminderResponse
func (client *Client) DeleteReminder(request *DeleteReminderRequest) (_result *DeleteReminderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteReminderHeaders{}
	_result = &DeleteReminderResponse{}
	_body, _err := client.DeleteReminderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Used by developers to obtain membership information in business scenarios that depend on CatGenie Membership.
//
// @param tmpReq - GetAccountForAppRequest
//
// @param headers - GetAccountForAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountForAppResponse
func (client *Client) GetAccountForAppWithOptions(tmpReq *GetAccountForAppRequest, headers *GetAccountForAppHeaders, runtime *dara.RuntimeOptions) (_result *GetAccountForAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used by developers to obtain membership information in business scenarios that depend on CatGenie Membership.
//
// @param request - GetAccountForAppRequest
//
// @return GetAccountForAppResponse
func (client *Client) GetAccountForApp(request *GetAccountForAppRequest) (_result *GetAccountForAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAccountForAppHeaders{}
	_result = &GetAccountForAppResponse{}
	_body, _err := client.GetAccountForAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain application configuration.
//
// @param tmpReq - GetBusAppConfigRequest
//
// @param headers - GetBusAppConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusAppConfigResponse
func (client *Client) GetBusAppConfigWithOptions(tmpReq *GetBusAppConfigRequest, headers *GetBusAppConfigHeaders, runtime *dara.RuntimeOptions) (_result *GetBusAppConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain application configuration.
//
// @param request - GetBusAppConfigRequest
//
// @return GetBusAppConfigResponse
func (client *Client) GetBusAppConfig(request *GetBusAppConfigRequest) (_result *GetBusAppConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetBusAppConfigHeaders{}
	_result = &GetBusAppConfigResponse{}
	_body, _err := client.GetBusAppConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the phone number bound to the user\\"s Maojing device. You must first request this capability on the Maojing Skill Application Platform and obtain user authorization on the device side.
//
// @param tmpReq - GetPhoneNumberRequest
//
// @param headers - GetPhoneNumberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhoneNumberResponse
func (client *Client) GetPhoneNumberWithOptions(tmpReq *GetPhoneNumberRequest, headers *GetPhoneNumberHeaders, runtime *dara.RuntimeOptions) (_result *GetPhoneNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the phone number bound to the user\\"s Maojing device. You must first request this capability on the Maojing Skill Application Platform and obtain user authorization on the device side.
//
// @param request - GetPhoneNumberRequest
//
// @return GetPhoneNumberResponse
func (client *Client) GetPhoneNumber(request *GetPhoneNumberRequest) (_result *GetPhoneNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetPhoneNumberHeaders{}
	_result = &GetPhoneNumberResponse{}
	_body, _err := client.GetPhoneNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// User retrieves reminders created on Tmall Genie devices.
//
// @param tmpReq - GetReminderRequest
//
// @param headers - GetReminderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReminderResponse
func (client *Client) GetReminderWithOptions(tmpReq *GetReminderRequest, headers *GetReminderHeaders, runtime *dara.RuntimeOptions) (_result *GetReminderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// User retrieves reminders created on Tmall Genie devices.
//
// @param request - GetReminderRequest
//
// @return GetReminderResponse
func (client *Client) GetReminder(request *GetReminderRequest) (_result *GetReminderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetReminderHeaders{}
	_result = &GetReminderResponse{}
	_body, _err := client.GetReminderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The User batch retrieves reminders created on Tmall Genie devices.
//
// @param tmpReq - ListRemindersRequest
//
// @param headers - ListRemindersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemindersResponse
func (client *Client) ListRemindersWithOptions(tmpReq *ListRemindersRequest, headers *ListRemindersHeaders, runtime *dara.RuntimeOptions) (_result *ListRemindersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The User batch retrieves reminders created on Tmall Genie devices.
//
// @param request - ListRemindersRequest
//
// @return ListRemindersResponse
func (client *Client) ListReminders(request *ListRemindersRequest) (_result *ListRemindersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListRemindersHeaders{}
	_result = &ListRemindersResponse{}
	_body, _err := client.ListRemindersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The user launches the checkout page on a Tmall Genie device.
//
// @param tmpReq - PullCashierRequest
//
// @param headers - PullCashierHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PullCashierResponse
func (client *Client) PullCashierWithOptions(tmpReq *PullCashierRequest, headers *PullCashierHeaders, runtime *dara.RuntimeOptions) (_result *PullCashierResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The user launches the checkout page on a Tmall Genie device.
//
// @param request - PullCashierRequest
//
// @return PullCashierResponse
func (client *Client) PullCashier(request *PullCashierRequest) (_result *PullCashierResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PullCashierHeaders{}
	_result = &PullCashierResponse{}
	_body, _err := client.PullCashierWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Execute the message push feature of Tmall Genie. You must first request the capability and configure the message template on the Tmall Genie Skill Application Platform.
//
// @param tmpReq - PushNotificationsRequest
//
// @param headers - PushNotificationsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushNotificationsResponse
func (client *Client) PushNotificationsWithOptions(tmpReq *PushNotificationsRequest, headers *PushNotificationsHeaders, runtime *dara.RuntimeOptions) (_result *PushNotificationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Execute the message push feature of Tmall Genie. You must first request the capability and configure the message template on the Tmall Genie Skill Application Platform.
//
// @param request - PushNotificationsRequest
//
// @return PushNotificationsResponse
func (client *Client) PushNotifications(request *PushNotificationsRequest) (_result *PushNotificationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PushNotificationsHeaders{}
	_result = &PushNotificationsResponse{}
	_body, _err := client.PushNotificationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Execute the message push feature of Tmall Genie. You must first request the capability and configure the message template in the Tmall Genie Skill Application Platform.
//
// @param tmpReq - SendNotificationsRequest
//
// @param headers - SendNotificationsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendNotificationsResponse
func (client *Client) SendNotificationsWithOptions(tmpReq *SendNotificationsRequest, headers *SendNotificationsHeaders, runtime *dara.RuntimeOptions) (_result *SendNotificationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Execute the message push feature of Tmall Genie. You must first request the capability and configure the message template in the Tmall Genie Skill Application Platform.
//
// @param request - SendNotificationsRequest
//
// @return SendNotificationsResponse
func (client *Client) SendNotifications(request *SendNotificationsRequest) (_result *SendNotificationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SendNotificationsHeaders{}
	_result = &SendNotificationsResponse{}
	_body, _err := client.SendNotificationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ThirdImmediateMsgPushWithOptions(request *ThirdImmediateMsgPushRequest, headers *ThirdImmediateMsgPushHeaders, runtime *dara.RuntimeOptions) (_result *ThirdImmediateMsgPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ThirdImmediateMsgPushResponse
func (client *Client) ThirdImmediateMsgPush(request *ThirdImmediateMsgPushRequest) (_result *ThirdImmediateMsgPushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ThirdImmediateMsgPushHeaders{}
	_result = &ThirdImmediateMsgPushResponse{}
	_body, _err := client.ThirdImmediateMsgPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The user updates a reminder created on a Tmall Genie device.
//
// @param tmpReq - UpdateReminderRequest
//
// @param headers - UpdateReminderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateReminderResponse
func (client *Client) UpdateReminderWithOptions(tmpReq *UpdateReminderRequest, headers *UpdateReminderHeaders, runtime *dara.RuntimeOptions) (_result *UpdateReminderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The user updates a reminder created on a Tmall Genie device.
//
// @param request - UpdateReminderRequest
//
// @return UpdateReminderResponse
func (client *Client) UpdateReminder(request *UpdateReminderRequest) (_result *UpdateReminderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateReminderHeaders{}
	_result = &UpdateReminderResponse{}
	_body, _err := client.UpdateReminderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Used by developers to report membership information in scenarios involving membership collaboration with CatGenie.
//
// @param tmpReq - VideoAppReportRequest
//
// @param headers - VideoAppReportHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoAppReportResponse
func (client *Client) VideoAppReportWithOptions(tmpReq *VideoAppReportRequest, headers *VideoAppReportHeaders, runtime *dara.RuntimeOptions) (_result *VideoAppReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used by developers to report membership information in scenarios involving membership collaboration with CatGenie.
//
// @param request - VideoAppReportRequest
//
// @return VideoAppReportResponse
func (client *Client) VideoAppReport(request *VideoAppReportRequest) (_result *VideoAppReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &VideoAppReportHeaders{}
	_result = &VideoAppReportResponse{}
	_body, _err := client.VideoAppReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// On a smart speaker with a screen, proactively launch the MeowGenie screen-display application. You must first request this capability on the MeowGenie skill application platform.
//
// @param request - WakeUpAppRequest
//
// @param headers - WakeUpAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WakeUpAppResponse
func (client *Client) WakeUpAppWithOptions(request *WakeUpAppRequest, headers *WakeUpAppHeaders, runtime *dara.RuntimeOptions) (_result *WakeUpAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// On a smart speaker with a screen, proactively launch the MeowGenie screen-display application. You must first request this capability on the MeowGenie skill application platform.
//
// @param request - WakeUpAppRequest
//
// @return WakeUpAppResponse
func (client *Client) WakeUpApp(request *WakeUpAppRequest) (_result *WakeUpAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &WakeUpAppHeaders{}
	_result = &WakeUpAppResponse{}
	_body, _err := client.WakeUpAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
