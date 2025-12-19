// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 获取工作负载访问令牌
//
// @param request - AssumeRoleForWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssumeRoleForWorkloadIdentityResponse
func (client *Client) AssumeRoleForWorkloadIdentityWithContext(ctx context.Context, request *AssumeRoleForWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *AssumeRoleForWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DurationSeconds) {
		body["DurationSeconds"] = request.DurationSeconds
	}

	if !dara.IsNil(request.Policy) {
		body["Policy"] = request.Policy
	}

	if !dara.IsNil(request.RoleSessionName) {
		body["RoleSessionName"] = request.RoleSessionName
	}

	if !dara.IsNil(request.WorkloadAccessToken) {
		body["WorkloadAccessToken"] = request.WorkloadAccessToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssumeRoleForWorkloadIdentity"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssumeRoleForWorkloadIdentityResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 确认用户授权会话已完成
//
// @param tmpReq - CompleteResourceTokenAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteResourceTokenAuthResponse
func (client *Client) CompleteResourceTokenAuthWithContext(ctx context.Context, tmpReq *CompleteResourceTokenAuthRequest, runtime *dara.RuntimeOptions) (_result *CompleteResourceTokenAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CompleteResourceTokenAuthShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIdentifier) {
		request.UserIdentifierShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdentifier, dara.String("UserIdentifier"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionURI) {
		body["SessionURI"] = request.SessionURI
	}

	if !dara.IsNil(request.UserIdentifierShrink) {
		body["UserIdentifier"] = request.UserIdentifierShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteResourceTokenAuth"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteResourceTokenAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取下游资源的 API 密钥
//
// @param request - GetResourceAPIKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceAPIKeyResponse
func (client *Client) GetResourceAPIKeyWithContext(ctx context.Context, request *GetResourceAPIKeyRequest, runtime *dara.RuntimeOptions) (_result *GetResourceAPIKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceCredentialProviderName) {
		body["ResourceCredentialProviderName"] = request.ResourceCredentialProviderName
	}

	if !dara.IsNil(request.WorkloadAccessToken) {
		body["WorkloadAccessToken"] = request.WorkloadAccessToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceAPIKey"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceAPIKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取下游资源的 OAuth 2.0 访问令牌
//
// @param tmpReq - GetResourceOAuth2TokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceOAuth2TokenResponse
func (client *Client) GetResourceOAuth2TokenWithContext(ctx context.Context, tmpReq *GetResourceOAuth2TokenRequest, runtime *dara.RuntimeOptions) (_result *GetResourceOAuth2TokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetResourceOAuth2TokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomParameters) {
		request.CustomParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomParameters, dara.String("CustomParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Scopes) {
		request.ScopesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scopes, dara.String("Scopes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomParametersShrink) {
		body["CustomParameters"] = request.CustomParametersShrink
	}

	if !dara.IsNil(request.CustomState) {
		body["CustomState"] = request.CustomState
	}

	if !dara.IsNil(request.ForceAuthentication) {
		body["ForceAuthentication"] = request.ForceAuthentication
	}

	if !dara.IsNil(request.OAuth2Flow) {
		body["OAuth2Flow"] = request.OAuth2Flow
	}

	if !dara.IsNil(request.ResourceCredentialProviderName) {
		body["ResourceCredentialProviderName"] = request.ResourceCredentialProviderName
	}

	if !dara.IsNil(request.ResourceOAuth2ReturnURL) {
		body["ResourceOAuth2ReturnURL"] = request.ResourceOAuth2ReturnURL
	}

	if !dara.IsNil(request.ScopesShrink) {
		body["Scopes"] = request.ScopesShrink
	}

	if !dara.IsNil(request.SessionURI) {
		body["SessionURI"] = request.SessionURI
	}

	if !dara.IsNil(request.WorkloadAccessToken) {
		body["WorkloadAccessToken"] = request.WorkloadAccessToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceOAuth2Token"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceOAuth2TokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作负载访问令牌
//
// @param request - GetWorkloadAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkloadAccessTokenResponse
func (client *Client) GetWorkloadAccessTokenWithContext(ctx context.Context, request *GetWorkloadAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadAccessTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkloadAccessToken"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkloadAccessTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作负载访问令牌
//
// @param request - GetWorkloadAccessTokenForJWTRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkloadAccessTokenForJWTResponse
func (client *Client) GetWorkloadAccessTokenForJWTWithContext(ctx context.Context, request *GetWorkloadAccessTokenForJWTRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadAccessTokenForJWTResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserToken) {
		body["UserToken"] = request.UserToken
	}

	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkloadAccessTokenForJWT"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkloadAccessTokenForJWTResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作负载访问令牌
//
// @param request - GetWorkloadAccessTokenForUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkloadAccessTokenForUserIdResponse
func (client *Client) GetWorkloadAccessTokenForUserIdWithContext(ctx context.Context, request *GetWorkloadAccessTokenForUserIdRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadAccessTokenForUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkloadAccessTokenForUserId"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkloadAccessTokenForUserIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
