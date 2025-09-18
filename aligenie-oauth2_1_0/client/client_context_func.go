// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建播放列表
//
// @param tmpReq - CreatePlayingListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePlayingListResponse
func (client *Client) CreatePlayingListWithContext(ctx context.Context, tmpReq *CreatePlayingListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePlayingListResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreatePlayingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenCreatePlayingListRequest) {
		request.OpenCreatePlayingListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenCreatePlayingListRequest, dara.String("OpenCreatePlayingListRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenCreatePlayingListRequestShrink) {
		body["OpenCreatePlayingListRequest"] = request.OpenCreatePlayingListRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePlayingList"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/content/playing/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePlayingListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行场景
//
// @param request - ExecuteSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSceneResponse
func (client *Client) ExecuteSceneWithContext(ctx context.Context, request *ExecuteSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteSceneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteScene"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/iot/scene/execute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行场景（全屋）
//
// @param request - ExecuteSmartHomeSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSmartHomeSceneResponse
func (client *Client) ExecuteSmartHomeSceneWithContext(ctx context.Context, request *ExecuteSmartHomeSceneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteSmartHomeSceneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FamilyId) {
		body["FamilyId"] = request.FamilyId
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteSmartHomeScene"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/iot/smart_home/scene/execute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteSmartHomeSceneResponse{}
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSceneListResponse
func (client *Client) GetSceneListWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSceneListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSceneList"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/iot/scene/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSceneListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取场景列表（全屋）
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmartHomeSceneListResponse
func (client *Client) GetSmartHomeSceneListWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSmartHomeSceneListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmartHomeSceneList"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/iot/smart_home/scene/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmartHomeSceneListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBasicInfoResponse
func (client *Client) GetUserBasicInfoWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserBasicInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserBasicInfo"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/users/basic"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserBasicInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取天猫精灵用户绑定的手机号
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserPhoneResponse
func (client *Client) GetUserPhoneWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserPhoneResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserPhone"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/user/profile/phone"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserPhoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OAuth2令牌撤销端点
//
// @param request - OAuth2RevocationEndpointRequest
//
// @param headers - OAuth2RevocationEndpointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OAuth2RevocationEndpointResponse
func (client *Client) OAuth2RevocationEndpointWithContext(ctx context.Context, request *OAuth2RevocationEndpointRequest, headers *OAuth2RevocationEndpointHeaders, runtime *dara.RuntimeOptions) (_result *OAuth2RevocationEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	if !dara.IsNil(request.TokenTypeHint) {
		body["TokenTypeHint"] = request.TokenTypeHint
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
		Action:      dara.String("OAuth2RevocationEndpoint"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/revoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OAuth2RevocationEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # OAuth2令牌端点
//
// @param request - OAuth2TokenEndpointRequest
//
// @param headers - OAuth2TokenEndpointHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OAuth2TokenEndpointResponse
func (client *Client) OAuth2TokenEndpointWithContext(ctx context.Context, request *OAuth2TokenEndpointRequest, headers *OAuth2TokenEndpointHeaders, runtime *dara.RuntimeOptions) (_result *OAuth2TokenEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.GrantType) {
		body["GrantType"] = request.GrantType
	}

	if !dara.IsNil(request.RedirectUri) {
		body["RedirectUri"] = request.RedirectUri
	}

	if !dara.IsNil(request.RefreshToken) {
		body["RefreshToken"] = request.RefreshToken
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
		Action:      dara.String("OAuth2TokenEndpoint"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/token"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OAuth2TokenEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送设备通知
//
// @param tmpReq - PushDeviceNotificationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushDeviceNotificationResponse
func (client *Client) PushDeviceNotificationWithContext(ctx context.Context, tmpReq *PushDeviceNotificationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushDeviceNotificationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PushDeviceNotificationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TenantInfo) {
		request.TenantInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantInfo, dara.String("TenantInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("body"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantInfoShrink) {
		body["TenantInfo"] = request.TenantInfoShrink
	}

	if !dara.IsNil(request.BodyShrink) {
		body["body"] = request.BodyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushDeviceNotification"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/device/notification/push"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushDeviceNotificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeviceListResponse
func (client *Client) QueryDeviceListWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryDeviceListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeviceList"),
		Version:     dara.String("oauth2_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/oauth2/device/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeviceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
