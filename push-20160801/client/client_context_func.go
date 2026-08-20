// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Attach an alias to a device.
//
// Description:
//
// You can attach up to 10 aliases in a single request. The attachment takes effect immediately.
//
// @param request - BindAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAliasResponse
func (client *Client) BindAliasWithContext(ctx context.Context, request *BindAliasRequest, runtime *dara.RuntimeOptions) (_result *BindAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAlias"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a device to a phone number.
//
// @param request - BindPhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPhoneResponse
func (client *Client) BindPhoneWithContext(ctx context.Context, request *BindPhoneRequest, runtime *dara.RuntimeOptions) (_result *BindPhoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindPhone"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPhoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds tags to specified device targets. Tag bindings take effect within 10 minutes.
//
// @param request - BindTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindTagResponse
func (client *Client) BindTagWithContext(ctx context.Context, request *BindTagRequest, runtime *dara.RuntimeOptions) (_result *BindTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.ClientKey) {
		query["ClientKey"] = request.ClientKey
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindTag"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a scheduled push task that has not yet been executed.
//
// @param request - CancelPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPushResponse
func (client *Client) CancelPushWithContext(ctx context.Context, request *CancelPushRequest, runtime *dara.RuntimeOptions) (_result *CancelPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelPush"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the expiration time and current status of the iOS certificate for a specified app.
//
// Description:
//
// - If the returned ExpireTime value is later than the current timestamp, the certificate is not necessarily valid. Also verify that the Status is OK.
//
// - The REVOKED status originates from the Apple Push Notification service (APNs) server. If a certificate has a REVOKED status, at least one push notification to APNs has failed in the corresponding environment.
//
// @param request - CheckCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCertificateResponse
func (client *Client) CheckCertificateWithContext(ctx context.Context, request *CheckCertificateRequest, runtime *dara.RuntimeOptions) (_result *CheckCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCertificate"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CheckDevice is deprecated, please use Push::2016-08-01::CheckDevices instead.
//
// Summary:
//
// Validates the specified (device).
//
// @param request - CheckDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDeviceResponse
func (client *Client) CheckDeviceWithContext(ctx context.Context, request *CheckDeviceRequest, runtime *dara.RuntimeOptions) (_result *CheckDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDevice"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validate a specified group of devices.
//
// @param request - CheckDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDevicesResponse
func (client *Client) CheckDevicesWithContext(ctx context.Context, request *CheckDevicesRequest, runtime *dara.RuntimeOptions) (_result *CheckDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceIds) {
		query["DeviceIds"] = request.DeviceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDevices"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually ends a continuous push task.
//
// Description:
//
// If you do not call this operation, the continuous push task automatically ends when it reaches its time-to-live (TTL).
//
// @param request - CompleteContinuouslyPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteContinuouslyPushResponse
func (client *Client) CompleteContinuouslyPushWithContext(ctx context.Context, request *CompleteContinuouslyPushRequest, runtime *dara.RuntimeOptions) (_result *CompleteContinuouslyPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteContinuouslyPush"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteContinuouslyPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a predefined continuous push task.
//
// Description:
//
// This API addresses the limitations of the [Push Advanced Push API](https://help.aliyun.com/document_detail/2249916.html), where push-by-device, push-by-account, and push-by-alias operations each have a maximum target count per single call.
//
// - You can use continuous push when your scenario requires sending the same message to many devices. In this case, you can call the continuous push API repeatedly, each time specifying a group of targets for aggregation (the current limit is 1,000 targets per call for device, account, or alias pushes). The total number of pushes for the same MessageId is restricted to 10,000. If you need a higher limit, contact technical support to evaluate your specific scenario.
//
// - Before using this API, you must first call the Push API with Target set to TBD (To Be Determined) and include your message content. This returns a MessageId from the push system. You can then use this MessageId to repeatedly call the continuous push API, specifying different target groups to deliver the same message.
//
// - After calling the Push API with Target set to TBD and obtaining a MessageId, the message is stored in the push system for 24 hours by default. You can use this API to push to specified targets at any time before expiration. Pushes are not allowed after expiration or after reaching the total push limit.
//
// - Each call to this API sends the message immediately. Scheduled pushes are not supported.
//
// @param request - ContinuouslyPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinuouslyPushResponse
func (client *Client) ContinuouslyPushWithContext(ctx context.Context, request *ContinuouslyPushRequest, runtime *dara.RuntimeOptions) (_result *ContinuouslyPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinuouslyPush"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinuouslyPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of an app. A maximum of 100 records are returned.
//
// @param request - ListTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithContext(ctx context.Context, request *ListTagsRequest, runtime *dara.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTags"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Push different messages/notifications to different devices in batch.
//
// Description:
//
// *Before using this API, make sure you fully understand the [billing methods and pricing](https://help.aliyun.com/document_detail/434638.html) of the EMAS Mobile Push service.**
//
// In certain business scenarios, you may need to push different messages to different devices in a very short time. When the number of devices and messages is extremely large, this can generate high QPS and reach the QPS limit per source IP, causing some pushes to fail.
//
// This API is designed to address this problem. It supports up to 100 independent push tasks in a single call, effectively reducing QPS through request aggregation and improving the stability and success rate of large-scale unicast pushes. Batch push is limited to 500 calls per second per account.
//
// Each independent push task only supports three types of push targets: device, account, and alias, and does not currently support SMS integration configuration.
//
// > The SDK must be upgraded to version 3.11.0 or later.
//
// ## PushTask Properties
//
// - The PushTask property format is: PushTask.N.Property. It includes:
//
//   - Push target (destination)
//
//   - Push configuration (config)
//
//   - iOS notification task configuration
//
//   - Android notification task configuration
//
//   - Android auxiliary popup configuration
//
//   - HarmonyOS notification task configuration
//
//   - Push control
//
// - Each PushTask represents an independent push task, with a maximum of 100 supported. Push-related configurations are consistent with the Push API.
//
// - The PushTask.N.Target parameter only supports three types: DEVICE, ACCOUNT, and ALIAS.
//
// - PushTask does not support SMS integration configuration.
//
// - The product of parent nodes and child nodes cannot exceed 10,000; otherwise, the parameters are considered invalid.
//
// @param request - MassPushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MassPushResponse
func (client *Client) MassPushWithContext(ctx context.Context, request *MassPushRequest, runtime *dara.RuntimeOptions) (_result *MassPushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.IdempotentToken) {
		query["IdempotentToken"] = request.IdempotentToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PushTask) {
		body["PushTask"] = request.PushTask
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MassPush"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MassPushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes notifications in batch by using the advanced push API V2.
//
// Description:
//
// Before using this operation, make sure that you fully understand the [billing method and pricing](https://help.aliyun.com/document_detail/434638.html) of EMAS Mobile Push.
//
// @param tmpReq - MassPushV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MassPushV2Response
func (client *Client) MassPushV2WithContext(ctx context.Context, tmpReq *MassPushV2Request, runtime *dara.RuntimeOptions) (_result *MassPushV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MassPushV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PushTasks) {
		request.PushTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PushTasks, dara.String("PushTasks"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.IdempotentToken) {
		query["IdempotentToken"] = request.IdempotentToken
	}

	if !dara.IsNil(request.PushTasksShrink) {
		query["PushTasks"] = request.PushTasksShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MassPushV2"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MassPushV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Advanced Push API that pushes notifications or messages to different device endpoints. This API provides rich push customization parameters to implement push behavior in different scenarios.
//
// Description:
//
// *Please ensure that you have fully understood the [billing methods and pricing](https://help.aliyun.com/document_detail/434638.html) of the EMAS Mobile Push product before using this API.**
//
// This API differentiates between Android, iOS, and HarmonyOS platforms. For push calls on different platforms, you need to pass the corresponding platform\\"s AppKey.
//
// @param tmpReq - PushRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushResponse
func (client *Client) PushWithContext(ctx context.Context, tmpReq *PushRequest, runtime *dara.RuntimeOptions) (_result *PushResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidOppoPrivateContentParameters) {
		request.AndroidOppoPrivateContentParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidOppoPrivateContentParameters, dara.String("AndroidOppoPrivateContentParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidOppoPrivateTitleParameters) {
		request.AndroidOppoPrivateTitleParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidOppoPrivateTitleParameters, dara.String("AndroidOppoPrivateTitleParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidActivity) {
		query["AndroidActivity"] = request.AndroidActivity
	}

	if !dara.IsNil(request.AndroidBadgeAddNum) {
		query["AndroidBadgeAddNum"] = request.AndroidBadgeAddNum
	}

	if !dara.IsNil(request.AndroidBadgeClass) {
		query["AndroidBadgeClass"] = request.AndroidBadgeClass
	}

	if !dara.IsNil(request.AndroidBadgeSetNum) {
		query["AndroidBadgeSetNum"] = request.AndroidBadgeSetNum
	}

	if !dara.IsNil(request.AndroidBigBody) {
		query["AndroidBigBody"] = request.AndroidBigBody
	}

	if !dara.IsNil(request.AndroidBigPictureUrl) {
		query["AndroidBigPictureUrl"] = request.AndroidBigPictureUrl
	}

	if !dara.IsNil(request.AndroidBigTitle) {
		query["AndroidBigTitle"] = request.AndroidBigTitle
	}

	if !dara.IsNil(request.AndroidExtParameters) {
		query["AndroidExtParameters"] = request.AndroidExtParameters
	}

	if !dara.IsNil(request.AndroidHonorTargetUserType) {
		query["AndroidHonorTargetUserType"] = request.AndroidHonorTargetUserType
	}

	if !dara.IsNil(request.AndroidHuaweiBusinessType) {
		query["AndroidHuaweiBusinessType"] = request.AndroidHuaweiBusinessType
	}

	if !dara.IsNil(request.AndroidHuaweiLiveNotificationPayload) {
		query["AndroidHuaweiLiveNotificationPayload"] = request.AndroidHuaweiLiveNotificationPayload
	}

	if !dara.IsNil(request.AndroidHuaweiReceiptId) {
		query["AndroidHuaweiReceiptId"] = request.AndroidHuaweiReceiptId
	}

	if !dara.IsNil(request.AndroidHuaweiTargetUserType) {
		query["AndroidHuaweiTargetUserType"] = request.AndroidHuaweiTargetUserType
	}

	if !dara.IsNil(request.AndroidImageUrl) {
		query["AndroidImageUrl"] = request.AndroidImageUrl
	}

	if !dara.IsNil(request.AndroidInboxBody) {
		query["AndroidInboxBody"] = request.AndroidInboxBody
	}

	if !dara.IsNil(request.AndroidMeizuNoticeMsgType) {
		query["AndroidMeizuNoticeMsgType"] = request.AndroidMeizuNoticeMsgType
	}

	if !dara.IsNil(request.AndroidMessageHuaweiCategory) {
		query["AndroidMessageHuaweiCategory"] = request.AndroidMessageHuaweiCategory
	}

	if !dara.IsNil(request.AndroidMessageHuaweiUrgency) {
		query["AndroidMessageHuaweiUrgency"] = request.AndroidMessageHuaweiUrgency
	}

	if !dara.IsNil(request.AndroidMessageOppoCategory) {
		query["AndroidMessageOppoCategory"] = request.AndroidMessageOppoCategory
	}

	if !dara.IsNil(request.AndroidMessageOppoNotifyLevel) {
		query["AndroidMessageOppoNotifyLevel"] = request.AndroidMessageOppoNotifyLevel
	}

	if !dara.IsNil(request.AndroidMessageVivoCategory) {
		query["AndroidMessageVivoCategory"] = request.AndroidMessageVivoCategory
	}

	if !dara.IsNil(request.AndroidMusic) {
		query["AndroidMusic"] = request.AndroidMusic
	}

	if !dara.IsNil(request.AndroidNotificationBarPriority) {
		query["AndroidNotificationBarPriority"] = request.AndroidNotificationBarPriority
	}

	if !dara.IsNil(request.AndroidNotificationBarType) {
		query["AndroidNotificationBarType"] = request.AndroidNotificationBarType
	}

	if !dara.IsNil(request.AndroidNotificationChannel) {
		query["AndroidNotificationChannel"] = request.AndroidNotificationChannel
	}

	if !dara.IsNil(request.AndroidNotificationGroup) {
		query["AndroidNotificationGroup"] = request.AndroidNotificationGroup
	}

	if !dara.IsNil(request.AndroidNotificationHonorChannel) {
		query["AndroidNotificationHonorChannel"] = request.AndroidNotificationHonorChannel
	}

	if !dara.IsNil(request.AndroidNotificationHuaweiChannel) {
		query["AndroidNotificationHuaweiChannel"] = request.AndroidNotificationHuaweiChannel
	}

	if !dara.IsNil(request.AndroidNotificationNotifyId) {
		query["AndroidNotificationNotifyId"] = request.AndroidNotificationNotifyId
	}

	if !dara.IsNil(request.AndroidNotificationThreadId) {
		query["AndroidNotificationThreadId"] = request.AndroidNotificationThreadId
	}

	if !dara.IsNil(request.AndroidNotificationVivoChannel) {
		query["AndroidNotificationVivoChannel"] = request.AndroidNotificationVivoChannel
	}

	if !dara.IsNil(request.AndroidNotificationXiaomiChannel) {
		query["AndroidNotificationXiaomiChannel"] = request.AndroidNotificationXiaomiChannel
	}

	if !dara.IsNil(request.AndroidNotifyType) {
		query["AndroidNotifyType"] = request.AndroidNotifyType
	}

	if !dara.IsNil(request.AndroidOpenType) {
		query["AndroidOpenType"] = request.AndroidOpenType
	}

	if !dara.IsNil(request.AndroidOpenUrl) {
		query["AndroidOpenUrl"] = request.AndroidOpenUrl
	}

	if !dara.IsNil(request.AndroidOppoDeleteIntentData) {
		query["AndroidOppoDeleteIntentData"] = request.AndroidOppoDeleteIntentData
	}

	if !dara.IsNil(request.AndroidOppoIntelligentIntent) {
		query["AndroidOppoIntelligentIntent"] = request.AndroidOppoIntelligentIntent
	}

	if !dara.IsNil(request.AndroidOppoIntentEnv) {
		query["AndroidOppoIntentEnv"] = request.AndroidOppoIntentEnv
	}

	if !dara.IsNil(request.AndroidOppoPrivateContentParametersShrink) {
		query["AndroidOppoPrivateContentParameters"] = request.AndroidOppoPrivateContentParametersShrink
	}

	if !dara.IsNil(request.AndroidOppoPrivateMsgTemplateId) {
		query["AndroidOppoPrivateMsgTemplateId"] = request.AndroidOppoPrivateMsgTemplateId
	}

	if !dara.IsNil(request.AndroidOppoPrivateTitleParametersShrink) {
		query["AndroidOppoPrivateTitleParameters"] = request.AndroidOppoPrivateTitleParametersShrink
	}

	if !dara.IsNil(request.AndroidPopupActivity) {
		query["AndroidPopupActivity"] = request.AndroidPopupActivity
	}

	if !dara.IsNil(request.AndroidPopupBody) {
		query["AndroidPopupBody"] = request.AndroidPopupBody
	}

	if !dara.IsNil(request.AndroidPopupTitle) {
		query["AndroidPopupTitle"] = request.AndroidPopupTitle
	}

	if !dara.IsNil(request.AndroidRemind) {
		query["AndroidRemind"] = request.AndroidRemind
	}

	if !dara.IsNil(request.AndroidRenderStyle) {
		query["AndroidRenderStyle"] = request.AndroidRenderStyle
	}

	if !dara.IsNil(request.AndroidTargetUserType) {
		query["AndroidTargetUserType"] = request.AndroidTargetUserType
	}

	if !dara.IsNil(request.AndroidVivoLiveMessage) {
		query["AndroidVivoLiveMessage"] = request.AndroidVivoLiveMessage
	}

	if !dara.IsNil(request.AndroidVivoPushMode) {
		query["AndroidVivoPushMode"] = request.AndroidVivoPushMode
	}

	if !dara.IsNil(request.AndroidVivoReceiptId) {
		query["AndroidVivoReceiptId"] = request.AndroidVivoReceiptId
	}

	if !dara.IsNil(request.AndroidXiaoMiActivity) {
		query["AndroidXiaoMiActivity"] = request.AndroidXiaoMiActivity
	}

	if !dara.IsNil(request.AndroidXiaoMiNotifyBody) {
		query["AndroidXiaoMiNotifyBody"] = request.AndroidXiaoMiNotifyBody
	}

	if !dara.IsNil(request.AndroidXiaoMiNotifyTitle) {
		query["AndroidXiaoMiNotifyTitle"] = request.AndroidXiaoMiNotifyTitle
	}

	if !dara.IsNil(request.AndroidXiaomiBigPictureUrl) {
		query["AndroidXiaomiBigPictureUrl"] = request.AndroidXiaomiBigPictureUrl
	}

	if !dara.IsNil(request.AndroidXiaomiFocusParam) {
		query["AndroidXiaomiFocusParam"] = request.AndroidXiaomiFocusParam
	}

	if !dara.IsNil(request.AndroidXiaomiFocusPics) {
		query["AndroidXiaomiFocusPics"] = request.AndroidXiaomiFocusPics
	}

	if !dara.IsNil(request.AndroidXiaomiImageUrl) {
		query["AndroidXiaomiImageUrl"] = request.AndroidXiaomiImageUrl
	}

	if !dara.IsNil(request.AndroidXiaomiTemplateId) {
		query["AndroidXiaomiTemplateId"] = request.AndroidXiaomiTemplateId
	}

	if !dara.IsNil(request.AndroidXiaomiTemplateParams) {
		query["AndroidXiaomiTemplateParams"] = request.AndroidXiaomiTemplateParams
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.HarmonyAction) {
		query["HarmonyAction"] = request.HarmonyAction
	}

	if !dara.IsNil(request.HarmonyActionType) {
		query["HarmonyActionType"] = request.HarmonyActionType
	}

	if !dara.IsNil(request.HarmonyBadgeAddNum) {
		query["HarmonyBadgeAddNum"] = request.HarmonyBadgeAddNum
	}

	if !dara.IsNil(request.HarmonyBadgeSetNum) {
		query["HarmonyBadgeSetNum"] = request.HarmonyBadgeSetNum
	}

	if !dara.IsNil(request.HarmonyCategory) {
		query["HarmonyCategory"] = request.HarmonyCategory
	}

	if !dara.IsNil(request.HarmonyExtParameters) {
		query["HarmonyExtParameters"] = request.HarmonyExtParameters
	}

	if !dara.IsNil(request.HarmonyExtensionExtraData) {
		query["HarmonyExtensionExtraData"] = request.HarmonyExtensionExtraData
	}

	if !dara.IsNil(request.HarmonyExtensionPush) {
		query["HarmonyExtensionPush"] = request.HarmonyExtensionPush
	}

	if !dara.IsNil(request.HarmonyImageUrl) {
		query["HarmonyImageUrl"] = request.HarmonyImageUrl
	}

	if !dara.IsNil(request.HarmonyInboxContent) {
		query["HarmonyInboxContent"] = request.HarmonyInboxContent
	}

	if !dara.IsNil(request.HarmonyLiveViewPayload) {
		query["HarmonyLiveViewPayload"] = request.HarmonyLiveViewPayload
	}

	if !dara.IsNil(request.HarmonyNotificationSlotType) {
		query["HarmonyNotificationSlotType"] = request.HarmonyNotificationSlotType
	}

	if !dara.IsNil(request.HarmonyNotifyId) {
		query["HarmonyNotifyId"] = request.HarmonyNotifyId
	}

	if !dara.IsNil(request.HarmonyReceiptId) {
		query["HarmonyReceiptId"] = request.HarmonyReceiptId
	}

	if !dara.IsNil(request.HarmonyRemind) {
		query["HarmonyRemind"] = request.HarmonyRemind
	}

	if !dara.IsNil(request.HarmonyRemindBody) {
		query["HarmonyRemindBody"] = request.HarmonyRemindBody
	}

	if !dara.IsNil(request.HarmonyRemindTitle) {
		query["HarmonyRemindTitle"] = request.HarmonyRemindTitle
	}

	if !dara.IsNil(request.HarmonyRenderStyle) {
		query["HarmonyRenderStyle"] = request.HarmonyRenderStyle
	}

	if !dara.IsNil(request.HarmonyTestMessage) {
		query["HarmonyTestMessage"] = request.HarmonyTestMessage
	}

	if !dara.IsNil(request.HarmonyUri) {
		query["HarmonyUri"] = request.HarmonyUri
	}

	if !dara.IsNil(request.IdempotentToken) {
		query["IdempotentToken"] = request.IdempotentToken
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.PushTime) {
		query["PushTime"] = request.PushTime
	}

	if !dara.IsNil(request.PushType) {
		query["PushType"] = request.PushType
	}

	if !dara.IsNil(request.SendChannels) {
		query["SendChannels"] = request.SendChannels
	}

	if !dara.IsNil(request.SendSpeed) {
		query["SendSpeed"] = request.SendSpeed
	}

	if !dara.IsNil(request.SmsDelaySecs) {
		query["SmsDelaySecs"] = request.SmsDelaySecs
	}

	if !dara.IsNil(request.SmsParams) {
		query["SmsParams"] = request.SmsParams
	}

	if !dara.IsNil(request.SmsSendPolicy) {
		query["SmsSendPolicy"] = request.SmsSendPolicy
	}

	if !dara.IsNil(request.SmsSignName) {
		query["SmsSignName"] = request.SmsSignName
	}

	if !dara.IsNil(request.SmsTemplateName) {
		query["SmsTemplateName"] = request.SmsTemplateName
	}

	if !dara.IsNil(request.StoreOffline) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.Trim) {
		query["Trim"] = request.Trim
	}

	if !dara.IsNil(request.IOSApnsEnv) {
		query["iOSApnsEnv"] = request.IOSApnsEnv
	}

	if !dara.IsNil(request.IOSBadge) {
		query["iOSBadge"] = request.IOSBadge
	}

	if !dara.IsNil(request.IOSBadgeAutoIncrement) {
		query["iOSBadgeAutoIncrement"] = request.IOSBadgeAutoIncrement
	}

	if !dara.IsNil(request.IOSExtParameters) {
		query["iOSExtParameters"] = request.IOSExtParameters
	}

	if !dara.IsNil(request.IOSInterruptionLevel) {
		query["iOSInterruptionLevel"] = request.IOSInterruptionLevel
	}

	if !dara.IsNil(request.IOSLiveActivityAttributes) {
		query["iOSLiveActivityAttributes"] = request.IOSLiveActivityAttributes
	}

	if !dara.IsNil(request.IOSLiveActivityAttributesType) {
		query["iOSLiveActivityAttributesType"] = request.IOSLiveActivityAttributesType
	}

	if !dara.IsNil(request.IOSLiveActivityContentState) {
		query["iOSLiveActivityContentState"] = request.IOSLiveActivityContentState
	}

	if !dara.IsNil(request.IOSLiveActivityDismissalDate) {
		query["iOSLiveActivityDismissalDate"] = request.IOSLiveActivityDismissalDate
	}

	if !dara.IsNil(request.IOSLiveActivityEvent) {
		query["iOSLiveActivityEvent"] = request.IOSLiveActivityEvent
	}

	if !dara.IsNil(request.IOSLiveActivityId) {
		query["iOSLiveActivityId"] = request.IOSLiveActivityId
	}

	if !dara.IsNil(request.IOSLiveActivityStaleDate) {
		query["iOSLiveActivityStaleDate"] = request.IOSLiveActivityStaleDate
	}

	if !dara.IsNil(request.IOSMusic) {
		query["iOSMusic"] = request.IOSMusic
	}

	if !dara.IsNil(request.IOSMutableContent) {
		query["iOSMutableContent"] = request.IOSMutableContent
	}

	if !dara.IsNil(request.IOSNotificationCategory) {
		query["iOSNotificationCategory"] = request.IOSNotificationCategory
	}

	if !dara.IsNil(request.IOSNotificationCollapseId) {
		query["iOSNotificationCollapseId"] = request.IOSNotificationCollapseId
	}

	if !dara.IsNil(request.IOSNotificationThreadId) {
		query["iOSNotificationThreadId"] = request.IOSNotificationThreadId
	}

	if !dara.IsNil(request.IOSRelevanceScore) {
		query["iOSRelevanceScore"] = request.IOSRelevanceScore
	}

	if !dara.IsNil(request.IOSRemind) {
		query["iOSRemind"] = request.IOSRemind
	}

	if !dara.IsNil(request.IOSRemindBody) {
		query["iOSRemindBody"] = request.IOSRemindBody
	}

	if !dara.IsNil(request.IOSSilentNotification) {
		query["iOSSilentNotification"] = request.IOSSilentNotification
	}

	if !dara.IsNil(request.IOSSubtitle) {
		query["iOSSubtitle"] = request.IOSSubtitle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Push"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message to an Android device through the Alibaba Cloud Mobile Push proprietary channel. After the app on the device receives the message, it must handle subsequent actions, such as implementing business logic or displaying a local notification.
//
// Description:
//
// *This operation will be deprecated soon. Use the [advanced push API](https://help.aliyun.com/document_detail/2249916.html), which provides enhanced push capabilities. To achieve the same result, set the `DeviceType` parameter to `ANDROID` and the `PushType` parameter to `MESSAGE` in the advanced push API.**
//
// **Before using this operation, review the [billing methods and pricing](https://help.aliyun.com/document_detail/434638.html) for EMAS Mobile Push.**
//
// By default, this operation sends messages only to online devices. If a device is offline, set the `StoreOffline` parameter. The push system then stores the message and delivers it automatically when the device comes online.
//
// @param request - PushMessageToAndroidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMessageToAndroidResponse
func (client *Client) PushMessageToAndroidWithContext(ctx context.Context, request *PushMessageToAndroidRequest, runtime *dara.RuntimeOptions) (_result *PushMessageToAndroidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.StoreOffline) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMessageToAndroid"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushMessageToAndroidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes messages to iOS devices. These messages are delivered through the proprietary channel of Alibaba Cloud Mobile Push. After the app on a device receives a message, it must handle subsequent actions, such as implementing business behaviors or creating local notifications.
//
// Description:
//
// *This API is deprecated. Use the [advanced push API](https://help.aliyun.com/document_detail/2249916.html) for more push capabilities. In that API, set the push platform `DeviceType` to `iOS` and the push type `PushType` to `MESSAGE` to achieve the same effect.**
//
// **Before you use this API, review the [billing methods and pricing](https://help.aliyun.com/document_detail/434638.html) for EMAS Mobile Push.**
//
// By default, this API sends messages only to online devices. If a device is offline, you can set the `StoreOffline` parameter. The push system then saves the message and automatically delivers it when the device comes back online.
//
// @param request - PushMessageToiOSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMessageToiOSResponse
func (client *Client) PushMessageToiOSWithContext(ctx context.Context, request *PushMessageToiOSRequest, runtime *dara.RuntimeOptions) (_result *PushMessageToiOSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.StoreOffline) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMessageToiOS"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushMessageToiOSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a notification to Android devices. The notification appears directly in the device’s notification tray and may be delivered through Alibaba Cloud’s proprietary channel or the device manufacturer’s channel, depending on the scenario.
//
// Description:
//
// *This operation is deprecated. Use the [Advanced Push API](https://help.aliyun.com/document_detail/2249916.html) instead. In that API, set the `DeviceType` parameter to `ANDROID` and the `PushType` parameter to `NOTICE`.**
//
// **Before using this operation, review the [pricing and billing model](https://help.aliyun.com/document_detail/434638.html) for EMAS Mobile Push.**
//
// @param request - PushNoticeToAndroidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushNoticeToAndroidResponse
func (client *Client) PushNoticeToAndroidWithContext(ctx context.Context, request *PushNoticeToAndroidRequest, runtime *dara.RuntimeOptions) (_result *PushNoticeToAndroidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.ExtParameters) {
		query["ExtParameters"] = request.ExtParameters
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.StoreOffline) {
		query["StoreOffline"] = request.StoreOffline
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushNoticeToAndroid"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushNoticeToAndroidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Send a notification to iOS devices. The notification uses Apple’s APNs channel and appears directly in the device notification center.
//
// Description:
//
// *This operation is deprecated. Use the [Advanced Push API](https://help.aliyun.com/document_detail/2249916.html) instead. Set the `DeviceType` parameter to `iOS` and the `PushType` parameter to `NOTICE`.**
//
// **Before you use this operation, review the [pricing and billing model](https://help.aliyun.com/document_detail/434638.html) for EMAS Mobile Push.**
//
// @param request - PushNoticeToiOSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushNoticeToiOSResponse
func (client *Client) PushNoticeToiOSWithContext(ctx context.Context, request *PushNoticeToiOSRequest, runtime *dara.RuntimeOptions) (_result *PushNoticeToiOSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApnsEnv) {
		query["ApnsEnv"] = request.ApnsEnv
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Body) {
		query["Body"] = request.Body
	}

	if !dara.IsNil(request.ExtParameters) {
		query["ExtParameters"] = request.ExtParameters
	}

	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushNoticeToiOS"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushNoticeToiOSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a push notification by using the advanced push API V2.
//
// Description:
//
// *Before you use this operation, make sure that you fully understand the [billing rules and pricing](https://help.aliyun.com/document_detail/434638.html) of EMAS mobile push.**
//
// This operation differentiates between the Android, iOS, and HarmonyOS platforms. To send push notifications to different platforms, pass in the AppKey that corresponds to the target platform.
//
// @param tmpReq - PushV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushV2Response
func (client *Client) PushV2WithContext(ctx context.Context, tmpReq *PushV2Request, runtime *dara.RuntimeOptions) (_result *PushV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PushTask) {
		request.PushTaskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PushTask, dara.String("PushTask"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.IdempotentToken) {
		query["IdempotentToken"] = request.IdempotentToken
	}

	if !dara.IsNil(request.PushTaskShrink) {
		query["PushTask"] = request.PushTaskShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushV2"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of aliases attached to a specified device.
//
// @param request - QueryAliasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAliasesResponse
func (client *Client) QueryAliasesWithContext(ctx context.Context, request *QueryAliasesRequest, runtime *dara.RuntimeOptions) (_result *QueryAliasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAliases"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAliasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query details of a specified device.
//
// @param request - QueryDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceInfoResponse
func (client *Client) QueryDeviceInfoWithContext(ctx context.Context, request *QueryDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceInfo"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries device statistics by application dimension.
//
// Description:
//
// > Currently, this API supports only daily data. The daily dimension lets you query data for up to 31 days. Days are calculated based on UTC+8.
//
// @param request - QueryDeviceStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceStatResponse
func (client *Client) QueryDeviceStatWithContext(ctx context.Context, request *QueryDeviceStatRequest, runtime *dara.RuntimeOptions) (_result *QueryDeviceStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceStat"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceStatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of devices associated with an account using the account name.
//
// @param request - QueryDevicesByAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicesByAccountResponse
func (client *Client) QueryDevicesByAccountWithContext(ctx context.Context, request *QueryDevicesByAccountRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicesByAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDevicesByAccount"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDevicesByAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of devices by alias.
//
// @param request - QueryDevicesByAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDevicesByAliasResponse
func (client *Client) QueryDevicesByAliasWithContext(ctx context.Context, request *QueryDevicesByAliasRequest, runtime *dara.RuntimeOptions) (_result *QueryDevicesByAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		query["Alias"] = request.Alias
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDevicesByAlias"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDevicesByAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can query push records with pagination and basic filtering.
//
// @param request - QueryPushRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushRecordsResponse
func (client *Client) QueryPushRecordsWithContext(ctx context.Context, request *QueryPushRecordsRequest, runtime *dara.RuntimeOptions) (_result *QueryPushRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PushType) {
		query["PushType"] = request.PushType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushRecords"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query push statistics for an app.
//
// @param request - QueryPushStatByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushStatByAppResponse
func (client *Client) QueryPushStatByAppWithContext(ctx context.Context, request *QueryPushStatByAppRequest, runtime *dara.RuntimeOptions) (_result *QueryPushStatByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushStatByApp"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushStatByAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries push statistics for a message.
//
// @param request - QueryPushStatByMsgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPushStatByMsgResponse
func (client *Client) QueryPushStatByMsgWithContext(ctx context.Context, request *QueryPushStatByMsgRequest, runtime *dara.RuntimeOptions) (_result *QueryPushStatByMsgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPushStatByMsg"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPushStatByMsgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags for a specified object, such as a device, account, or alias.
//
// @param request - QueryTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagsResponse
func (client *Client) QueryTagsWithContext(ctx context.Context, request *QueryTagsRequest, runtime *dara.RuntimeOptions) (_result *QueryTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.ClientKey) {
		query["ClientKey"] = request.ClientKey
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTags"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain deduplicated device statistics for an app.
//
// Description:
//
// > This operation returns data only at the daily granularity. You can query up to 31 days of data. Deduplicated device counts reset on the first day of each month.
//
// @param request - QueryUniqueDeviceStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUniqueDeviceStatResponse
func (client *Client) QueryUniqueDeviceStatWithContext(ctx context.Context, request *QueryUniqueDeviceStatRequest, runtime *dara.RuntimeOptions) (_result *QueryUniqueDeviceStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUniqueDeviceStat"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUniqueDeviceStatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a tag from an app.
//
// Description:
//
// Deleting a tag takes time. The time required depends on the number of tagged resources. Do not immediately recreate a tag with the same name after you delete it. Wait at least 5 minutes before you recreate a tag in the same app. If you delete multiple tags, wait at least 5 minutes for each deleted tag before you recreate them.
//
// @param request - RemoveTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTagResponse
func (client *Client) RemoveTagWithContext(ctx context.Context, request *RemoveTagRequest, runtime *dara.RuntimeOptions) (_result *RemoveTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTag"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an alias. The change takes effect immediately.
//
// @param request - UnbindAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAliasResponse
func (client *Client) UnbindAliasWithContext(ctx context.Context, request *UnbindAliasRequest, runtime *dara.RuntimeOptions) (_result *UnbindAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasName) {
		query["AliasName"] = request.AliasName
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.UnbindAll) {
		query["UnbindAll"] = request.UnbindAll
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindAlias"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbind the mobile phone number from a specified device.
//
// @param request - UnbindPhoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPhoneResponse
func (client *Client) UnbindPhoneWithContext(ctx context.Context, request *UnbindPhoneRequest, runtime *dara.RuntimeOptions) (_result *UnbindPhoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindPhone"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindPhoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds one or more tags from a specified target.
//
// @param request - UnbindTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindTagResponse
func (client *Client) UnbindTagWithContext(ctx context.Context, request *UnbindTagRequest, runtime *dara.RuntimeOptions) (_result *UnbindTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.ClientKey) {
		query["ClientKey"] = request.ClientKey
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindTag"),
		Version:     dara.String("2016-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
