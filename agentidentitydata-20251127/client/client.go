// This file is auto-generated, don't edit it. Thanks.
package client

import (
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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("agentidentitydata"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 获取工作负载访问令牌
//
// @param request - AssumeRoleForWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssumeRoleForWorkloadIdentityResponse
func (client *Client) AssumeRoleForWorkloadIdentityWithOptions(request *AssumeRoleForWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *AssumeRoleForWorkloadIdentityResponse, _err error) {
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
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
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
// @param request - AssumeRoleForWorkloadIdentityRequest
//
// @return AssumeRoleForWorkloadIdentityResponse
func (client *Client) AssumeRoleForWorkloadIdentity(request *AssumeRoleForWorkloadIdentityRequest) (_result *AssumeRoleForWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssumeRoleForWorkloadIdentityResponse{}
	_body, _err := client.AssumeRoleForWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CompleteResourceTokenAuthWithOptions(tmpReq *CompleteResourceTokenAuthRequest, runtime *dara.RuntimeOptions) (_result *CompleteResourceTokenAuthResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CompleteResourceTokenAuthRequest
//
// @return CompleteResourceTokenAuthResponse
func (client *Client) CompleteResourceTokenAuth(request *CompleteResourceTokenAuthRequest) (_result *CompleteResourceTokenAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompleteResourceTokenAuthResponse{}
	_body, _err := client.CompleteResourceTokenAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetResourceAPIKeyWithOptions(request *GetResourceAPIKeyRequest, runtime *dara.RuntimeOptions) (_result *GetResourceAPIKeyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetResourceAPIKeyResponse
func (client *Client) GetResourceAPIKey(request *GetResourceAPIKeyRequest) (_result *GetResourceAPIKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceAPIKeyResponse{}
	_body, _err := client.GetResourceAPIKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetResourceOAuth2TokenWithOptions(tmpReq *GetResourceOAuth2TokenRequest, runtime *dara.RuntimeOptions) (_result *GetResourceOAuth2TokenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - GetResourceOAuth2TokenRequest
//
// @return GetResourceOAuth2TokenResponse
func (client *Client) GetResourceOAuth2Token(request *GetResourceOAuth2TokenRequest) (_result *GetResourceOAuth2TokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceOAuth2TokenResponse{}
	_body, _err := client.GetResourceOAuth2TokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWorkloadAccessTokenWithOptions(request *GetWorkloadAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadAccessTokenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWorkloadAccessTokenResponse
func (client *Client) GetWorkloadAccessToken(request *GetWorkloadAccessTokenRequest) (_result *GetWorkloadAccessTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkloadAccessTokenResponse{}
	_body, _err := client.GetWorkloadAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWorkloadAccessTokenForJWTWithOptions(request *GetWorkloadAccessTokenForJWTRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadAccessTokenForJWTResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWorkloadAccessTokenForJWTResponse
func (client *Client) GetWorkloadAccessTokenForJWT(request *GetWorkloadAccessTokenForJWTRequest) (_result *GetWorkloadAccessTokenForJWTResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkloadAccessTokenForJWTResponse{}
	_body, _err := client.GetWorkloadAccessTokenForJWTWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetWorkloadAccessTokenForUserIdWithOptions(request *GetWorkloadAccessTokenForUserIdRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadAccessTokenForUserIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetWorkloadAccessTokenForUserIdResponse
func (client *Client) GetWorkloadAccessTokenForUserId(request *GetWorkloadAccessTokenForUserIdRequest) (_result *GetWorkloadAccessTokenForUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkloadAccessTokenForUserIdResponse{}
	_body, _err := client.GetWorkloadAccessTokenForUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
