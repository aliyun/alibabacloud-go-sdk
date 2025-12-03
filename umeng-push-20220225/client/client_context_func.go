// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 根据消息ID取消发送
//
// @param request - CancelByMsgIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelByMsgIdResponse
func (client *Client) CancelByMsgIdWithContext(ctx context.Context, request *CancelByMsgIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelByMsgIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelByMsgId"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/CancelByMsgId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelByMsgIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息状态查询
//
// @param request - QueryMsgStatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMsgStatResponse
func (client *Client) QueryMsgStatWithContext(ctx context.Context, request *QueryMsgStatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryMsgStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMsgStat"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/QueryMsgStat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMsgStatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定别名发送
//
// @param tmpReq - SendByAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAliasResponse
func (client *Client) SendByAliasWithContext(ctx context.Context, tmpReq *SendByAliasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByAliasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.AliasType) {
		body["AliasType"] = request.AliasType
	}

	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByAlias"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByAlias"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定别名文件发送
//
// @param tmpReq - SendByAliasFileIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAliasFileIdResponse
func (client *Client) SendByAliasFileIdWithContext(ctx context.Context, tmpReq *SendByAliasFileIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByAliasFileIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByAliasFileIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliasType) {
		body["AliasType"] = request.AliasType
	}

	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByAliasFileId"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByAliasFileId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByAliasFileIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 广播
//
// @param tmpReq - SendByAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAppResponse
func (client *Client) SendByAppWithContext(ctx context.Context, tmpReq *SendByAppRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByApp"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByApp"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定设备发送
//
// @param tmpReq - SendByDeviceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByDeviceResponse
func (client *Client) SendByDeviceWithContext(ctx context.Context, tmpReq *SendByDeviceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DeviceTokens) {
		body["DeviceTokens"] = request.DeviceTokens
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByDevice"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定设备文件发送
//
// @param tmpReq - SendByDeviceFileIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByDeviceFileIdResponse
func (client *Client) SendByDeviceFileIdWithContext(ctx context.Context, tmpReq *SendByDeviceFileIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByDeviceFileIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByDeviceFileIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByDeviceFileId"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByDeviceFileId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByDeviceFileIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据筛选条件发送
//
// @param tmpReq - SendByFilterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByFilterResponse
func (client *Client) SendByFilterWithContext(ctx context.Context, tmpReq *SendByFilterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByFilterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByFilterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayload) {
		body["AndroidShortPayload"] = request.AndroidShortPayload
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByFilter"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByFilter"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传设备列表创建设备文件
//
// @param request - UploadDeviceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDeviceResponse
func (client *Client) UploadDeviceWithContext(ctx context.Context, request *UploadDeviceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTokens) {
		body["DeviceTokens"] = request.DeviceTokens
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDevice"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/UploadDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
