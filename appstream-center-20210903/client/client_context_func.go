// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Retrieves the connection credentials for cloud applications.
//
// Description:
//
// It is a time-consuming operation to issue connection credentials. Therefore, you may need to invoke this API operation multiple times to obtain the connection credentials.
//
// 1. When you invoke the API for the first time, an application instance is allocated to the specified convenience account, and the application is started. The start node ID (`TaskId`) is returned.
//
// 2. When you invoke the API subsequently, you must pass the `TaskId` parameter as a request parameter to check whether the node is completed. When the returned node status (`TaskStatus`) is `Finished`, the connection ticket (`Ticket`) is also returned.
//
// >
//
// > - If the startup process is fast enough, the connection ticket (`Ticket`) may also be returned when you invoke the API for the first time.
//
// @param request - GetConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicketWithContext(ctx context.Context, request *GetConnectionTicketRequest, runtime *dara.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		body["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceId) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.AppPolicyId) {
		body["AppPolicyId"] = request.AppPolicyId
	}

	if !dara.IsNil(request.AppVersion) {
		body["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.AutoConnectInQueue) {
		body["AutoConnectInQueue"] = request.AutoConnectInQueue
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.ConnectionProperties) {
		body["ConnectionProperties"] = request.ConnectionProperties
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.EnvironmentConfig) {
		body["EnvironmentConfig"] = request.EnvironmentConfig
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.Param) {
		body["Param"] = request.Param
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RequireUuidValidation) {
		body["RequireUuidValidation"] = request.RequireUuidValidation
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnectionTicket"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of published applications.
//
// @param request - ListPublishedAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedAppInfoResponse
func (client *Client) ListPublishedAppInfoWithContext(ctx context.Context, request *ListPublishedAppInfoRequest, runtime *dara.RuntimeOptions) (_result *ListPublishedAppInfoResponse, _err error) {
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

	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CategoryType) {
		query["CategoryType"] = request.CategoryType
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginRegionId) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OrderParam) {
		query["OrderParam"] = request.OrderParam
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishedAppInfo"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishedAppInfoResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries running applications.
//
// Description:
//
// ## Prerequisites
//
// This operation is an anonymous API operation, but the server still verifies the logon status. The `LoginToken` and `SessionId` request parameters are actually required. If either parameter is missing, an HTTP 400 error is returned with the error code `InvalidParameter.loginToken` or `InvalidParameter.sessionId`.
//
// ## Obtain credentials
//
// To obtain the `LoginToken` and `SessionId` by using `GetLoginToken`, call the `GetLoginToken` operation of the `appstream-center` 2021-02-20 version. The 2021-09-03 version does not provide this operation. You can also call `RefreshLoginToken` to update and obtain the `LoginToken`.
//
// @param request - ListRunningAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRunningAppsResponse
func (client *Client) ListRunningAppsWithContext(ctx context.Context, request *ListRunningAppsRequest, runtime *dara.RuntimeOptions) (_result *ListRunningAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRegionId) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginRegionId) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TenantId) {
		query["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRunningApps"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRunningAppsResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets application resources.
//
// @param request - ResetAppResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAppResourcesResponse
func (client *Client) ResetAppResourcesWithContext(ctx context.Context, request *ResetAppResourcesRequest, runtime *dara.RuntimeOptions) (_result *ResetAppResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAppResources"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAppResourcesResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts application resources.
//
// @param request - RestartAppResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartAppResourcesResponse
func (client *Client) RestartAppResourcesWithContext(ctx context.Context, request *RestartAppResourcesRequest, runtime *dara.RuntimeOptions) (_result *RestartAppResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartAppResources"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartAppResourcesResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts application resources.
//
// @param request - StartAppResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAppResourcesResponse
func (client *Client) StartAppResourcesWithContext(ctx context.Context, request *StartAppResourcesRequest, runtime *dara.RuntimeOptions) (_result *StartAppResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAppResources"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAppResourcesResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a specified application.
//
// Description:
//
// ## Credential acquisition
//
// To obtain the `LoginToken` and `SessionId` by using `GetLoginToken`, call the `GetLoginToken` operation of the `appstream-center` 2021-02-20 version. The 2021-09-03 version does not provide this operation. You can also update and obtain the `LoginToken` by using `RefreshLoginToken`.
//
// ### Required parameters and authentication invocation flow for GetLoginToken (2021-02-20 version)
//
// API documentation: [GetLoginToken (2021-02-20)](https://www.alibabacloud.com/help/en/wuying-appstreaming/developer-reference/api-appstream-center-2021-02-20-getlogintoken)
//
// Required metric descriptions:
//
// - `ClientId`: Required. The unique ID generated by the client, serving as the identity for the client that initiates the logon. Reuse the same value within the same session.
//
// - `LoginIdentifier`: The enterprise identity (i.e., enterprise alias). This parameter is required when you use the enterprise identity logon method. Note that the parameter name is `LoginIdentifier`, not `TenantAlias`.
//
// - When you use a convenience account or an AD domain account for authentication, the Settings for `CurrentStage` must be `ADPassword`, and you must provide both `EndUserId` (username) and `Password` (password).
//
// Complete authentication invocation flow:
//
// 1. Invoke `GetLoginToken` for the first time with `ClientId`, `LoginIdentifier`, `CurrentStage` based on the logon method, and the corresponding credentials. Save the returned `LoginToken` and `SessionId`. `SessionId` is returned only during the first invocation in the same session and must be reused in subsequent stage invocations.
//
// 2. If `NextStage` (such as `MFAVerify`, `TokenVerify`, or `ChangePassword`) is returned, carry the same `SessionId` and `ClientId`, update the Settings of `CurrentStage` to the value of `NextStage`, and provide the parameters required for this stage (such as `AuthenticationCode` for the `MFAVerify` stage and `TokenCode` for the `TokenVerify` stage). Repeat the invocation until the final `LoginToken` is obtained.
//
// 3. Use the final `LoginToken` and `SessionId` to invoke this operation (`StopApp`).
//
// @param request - StopAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAppResponse
func (client *Client) StopAppWithContext(ctx context.Context, request *StopAppRequest, runtime *dara.RuntimeOptions) (_result *StopAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		body["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.ApiType) {
		body["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceId) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ClientChannel) {
		body["ClientChannel"] = request.ClientChannel
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ForceStop) {
		body["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.IdpId) {
		body["IdpId"] = request.IdpId
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	if !dara.IsNil(request.WyId) {
		body["WyId"] = request.WyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopApp"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAppResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops application resources.
//
// Description:
//
// ## Before you begin
//
// Before you call this operation, you must complete user logon authentication to obtain the `LoginToken` and `SessionId`. This `StopAppResources` operation uses the **2021-09-03*	- version of `appstream-center`. The prerequisite logon operation uses the **2021-02-20*	- version of [GetLoginToken](https://www.alibabacloud.com/help/en/wuying-appstreaming/developer-reference/api-appstream-center-2021-02-20-getlogintoken). Use the respective API versions for each operation.
//
// ### 1. Prepare the enterprise identifier and account information
//
// The following procedure applies to password authentication for convenience accounts or AD domain accounts that log on by using an enterprise identifier.
//
// | Parameter | Description |
//
// | --- | --- |
//
// | `ClientId` | The unique ID generated by the client. This parameter is required when you call `GetLoginToken`. The same client ID is reused in the following call sequence. |
//
// | `LoginIdentifier` | The enterprise identifier for logon, which is the `TenantAlias`. The parameter name in the logon request is `LoginIdentifier`, not `TenantAlias`. To query the logon configuration of this identifier, you can call [FindIdpListByLoginIdentifier (2021-02-20)](https://www.alibabacloud.com/help/en/wuying-appstreaming/developer-reference/api-appstream-center-2021-02-20-findidplistbyloginidentifier). The `TenantAliasInfo.TenantAlias` field in the response indicates the enterprise identifier. |
//
// | `CurrentStage` | Set this parameter to `ADPassword` for password authentication. This value applies to both convenience accounts and AD domain accounts, not only AD domain accounts. |
//
// | `EndUserId` | The username of the convenience account or AD domain account. This parameter is required when `CurrentStage` is set to `ADPassword`. |
//
// | `Password` | The password of the corresponding account. This parameter is required when `CurrentStage` is set to `ADPassword`. |
//
// ### 2. Continue logon based on the returned authentication stage
//
// After you invoke `GetLoginToken` for the first time, save the returned `SessionId`. This value is returned only during the first invocation in the same logon session. If you need to authenticate further, include the saved `SessionId` in subsequent requests. Do not clear this value even if it is not returned in subsequent responses.
//
// Determine whether you still need to authenticate based on the `NextStage` field in the response, and set the required next authentication stage as the `CurrentStage` for subsequent requests. Examples:
//
// | Next authentication stage | What to do next |
//
// | --- | --- |
//
// | `MFAVerify` | Include the `SessionId` and the `AuthenticationCode` generated by the virtual MFA device to continue to invoke `GetLoginToken`. |
//
// | `TokenVerify` | Include the `SessionId` and the client secondary authentication verification code `TokenCode` to continue to authenticate. |
//
// | `ChangePassword` | Specify `OldPassword` and `NewPassword` as required to complete the password change stage. |
//
// | `MFABind` or other authentication stages | Complete the corresponding operations based on the actual returned stage and the `GetLoginToken` documentation. Do not directly apply the request parameters of `MFAVerify`. |
//
// The authentication stages are determined by the account configuration and the result of the current logon. Not every logon must go through all the preceding stages in sequence. Invoke `StopAppResources` only after logon authentication is complete and the `LoginToken` is obtained. Do not proceed with the shutdown of resources if authentication failed or if there are pending authentication stages.
//
// ### 3. Call sequence example
//
// The following is pseudocode for parameter passing and is not an executable SDK example. `clientId`, `tenantAlias`, `endUserId`, `password`, and `authenticationCode` are runtime variables of the caller. `resourceIds`, `productType`, and `appInstanceGroupId` are business parameters for this operation. Do not write real passwords, verification codes, or logon credentials in examples or logs.
//
// ```text
//
// login1 = GetLoginToken(version="2021-02-20",
//
//	ClientId=clientId, LoginIdentifier=tenantAlias,
//
//	CurrentStage="ADPassword", EndUserId=endUserId, Password=password)
//
// sessionId = login1.SessionId
//
// loginResult = login1
//
// if loginResult.NextStage == "MFAVerify":
//
//	loginResult = GetLoginToken(version="2021-02-20",
//
//	    ClientId=clientId, LoginIdentifier=tenantAlias,
//
//	    SessionId=sessionId, CurrentStage="MFAVerify",
//
//	    AuthenticationCode=authenticationCode)
//
// ```
//
// If other authentication stages are still required, continue processing based on `NextStage`. Do not treat the preceding single MFA call as a fixed endpoint. After logon is complete, call this operation by using the `LoginToken` obtained during this logon and the saved `SessionId`:
//
// ```text
//
// StopAppResources(version="2021-09-03",
//
//	ClientId=clientId, LoginToken=loginResult.LoginToken, SessionId=sessionId,
//
//	ProductType=productType, AppInstanceGroupId=appInstanceGroupId,
//
//	ResourceIds=resourceIds)
//
// ```
//
// `Password`, `LoginIdentifier`, and `CurrentStage` belong to the prerequisite logon request and are not request parameters of `StopAppResources`. For the rules on how to specify other parameters and resource IDs for this operation, see the request parameter table below.
//
// @param request - StopAppResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAppResourcesResponse
func (client *Client) StopAppResourcesWithContext(ctx context.Context, request *StopAppResourcesRequest, runtime *dara.RuntimeOptions) (_result *StopAppResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.BizRegionId) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceIds) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAppResources"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAppResourcesResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an application instance from a user.
//
// @param request - UnbindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindResponse
func (client *Client) UnbindWithContext(ctx context.Context, request *UnbindRequest, runtime *dara.RuntimeOptions) (_result *UnbindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppInstanceGroupId) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !dara.IsNil(request.AppInstanceId) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !dara.IsNil(request.AppInstancePersistentId) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SessionToken) {
		body["SessionToken"] = request.SessionToken
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Unbind"),
		Version:     dara.String("2021-09-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
