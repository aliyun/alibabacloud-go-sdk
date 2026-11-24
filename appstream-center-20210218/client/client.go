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
	client.ProductId = dara.String("appstream-center")
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
	client.Endpoint, _err = client.GetEndpoint(dara.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a time-limited static access token for a specified user in the current account to access a single cloud desktop.
//
// Description:
//
// When you call this operation, you must provide Resource and specify both ResourceType and ResourceId. Only a single cloud desktop is supported. You cannot omit Resource to authorize or query all cloud desktops.
//
// AuthorizerId is the WyId of the user and cannot be replaced by a username. BusinessChannel supports Enterprise and Business. If omitted, the default value is Enterprise.
//
// Before creation, the user must belong to the current calling account and have access permissions for the target cloud desktop. This operation only creates an access token. It does not create a user or add an authorization relationship between the user and the cloud desktop. Each successful call creates a new time-limited token. To obtain an existing token, call GetLatestThirdAppAccessToken and pass the same AuthorizerId and Resource.
//
// The unit of Expiration is seconds. Valid values: 1 to 15552000. If omitted, the default value is 15552000, which equals 180 days. Permanent validity is not supported. Keep the returned AccessToken secure and avoid displaying the full token in logs or public content.
//
// @param tmpReq - CreateThirdAppAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateThirdAppAccessTokenResponse
func (client *Client) CreateThirdAppAccessTokenWithOptions(tmpReq *CreateThirdAppAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateThirdAppAccessTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateThirdAppAccessTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Resource) {
		request.ResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resource, dara.String("Resource"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizerId) {
		body["AuthorizerId"] = request.AuthorizerId
	}

	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Expiration) {
		body["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ResourceShrink) {
		body["Resource"] = request.ResourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateThirdAppAccessToken"),
		Version:     dara.String("2021-02-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateThirdAppAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a time-limited static access token for a specified user in the current account to access a single cloud desktop.
//
// Description:
//
// When you call this operation, you must provide Resource and specify both ResourceType and ResourceId. Only a single cloud desktop is supported. You cannot omit Resource to authorize or query all cloud desktops.
//
// AuthorizerId is the WyId of the user and cannot be replaced by a username. BusinessChannel supports Enterprise and Business. If omitted, the default value is Enterprise.
//
// Before creation, the user must belong to the current calling account and have access permissions for the target cloud desktop. This operation only creates an access token. It does not create a user or add an authorization relationship between the user and the cloud desktop. Each successful call creates a new time-limited token. To obtain an existing token, call GetLatestThirdAppAccessToken and pass the same AuthorizerId and Resource.
//
// The unit of Expiration is seconds. Valid values: 1 to 15552000. If omitted, the default value is 15552000, which equals 180 days. Permanent validity is not supported. Keep the returned AccessToken secure and avoid displaying the full token in logs or public content.
//
// @param request - CreateThirdAppAccessTokenRequest
//
// @return CreateThirdAppAccessTokenResponse
func (client *Client) CreateThirdAppAccessToken(request *CreateThirdAppAccessTokenRequest) (_result *CreateThirdAppAccessTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateThirdAppAccessTokenResponse{}
	_body, _err := client.CreateThirdAppAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invalidates a Wuying Workspace logon session before it naturally expires.
//
// @param request - ExpireLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpireLoginTokenResponse
func (client *Client) ExpireLoginTokenWithOptions(request *ExpireLoginTokenRequest, runtime *dara.RuntimeOptions) (_result *ExpireLoginTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		body["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExpireLoginToken"),
		Version:     dara.String("2021-02-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExpireLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invalidates a Wuying Workspace logon session before it naturally expires.
//
// @param request - ExpireLoginTokenRequest
//
// @return ExpireLoginTokenResponse
func (client *Client) ExpireLoginToken(request *ExpireLoginTokenRequest) (_result *ExpireLoginTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExpireLoginTokenResponse{}
	_body, _err := client.ExpireLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves an authorization code that contains user identity and permission information, which can be used to invoke cloud applications in integration scenarios.
//
// Description:
//
// Retrieves an authorization code that contains user identity and permission information, which can be used to invoke cloud applications in integration scenarios.
//
// @param request - GetAuthCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthCodeResponse
func (client *Client) GetAuthCodeWithOptions(request *GetAuthCodeRequest, runtime *dara.RuntimeOptions) (_result *GetAuthCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TokenType) {
		query["TokenType"] = request.TokenType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountType) {
		body["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.AdDomain) {
		body["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.AdPassword) {
		body["AdPassword"] = request.AdPassword
	}

	if !dara.IsNil(request.AutoCreateUser) {
		body["AutoCreateUser"] = request.AutoCreateUser
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.Policy) {
		body["Policy"] = request.Policy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthCode"),
		Version:     dara.String("2021-02-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an authorization code that contains user identity and permission information, which can be used to invoke cloud applications in integration scenarios.
//
// Description:
//
// Retrieves an authorization code that contains user identity and permission information, which can be used to invoke cloud applications in integration scenarios.
//
// @param request - GetAuthCodeRequest
//
// @return GetAuthCodeResponse
func (client *Client) GetAuthCode(request *GetAuthCodeRequest) (_result *GetAuthCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAuthCodeResponse{}
	_body, _err := client.GetAuthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the current valid static access token for a specified user and a single cloud desktop under the current account, keeping the original expiration time.
//
// Description:
//
// You must provide Resource and specify both ResourceType and ResourceId when calling this operation. Only a single cloud desktop is supported. You cannot omit Resource to authorize or query all cloud desktops.
//
// AuthorizerId is the WyId of the user and cannot be replaced by a username. BusinessChannel supports Enterprise and Business. If omitted, the default value is Enterprise.
//
// The AuthorizerId and Resource in the query conditions must be consistent with those used when the token was created. This operation only returns the current valid token for the current account, user, and cloud desktop. It does not create a new token, renew the token, or restore or roll back to a historical token. If no valid token is found, the token has expired, or the corresponding cache has been cleared, AccessToken.NotFound is returned. To obtain a new token, call CreateThirdAppAccessToken.
//
// Keep the returned AccessToken secure and avoid displaying the full token in logs or public content.
//
// @param tmpReq - GetLatestThirdAppAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLatestThirdAppAccessTokenResponse
func (client *Client) GetLatestThirdAppAccessTokenWithOptions(tmpReq *GetLatestThirdAppAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *GetLatestThirdAppAccessTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetLatestThirdAppAccessTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Resource) {
		request.ResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resource, dara.String("Resource"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizerId) {
		body["AuthorizerId"] = request.AuthorizerId
	}

	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.ResourceShrink) {
		body["Resource"] = request.ResourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLatestThirdAppAccessToken"),
		Version:     dara.String("2021-02-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLatestThirdAppAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the current valid static access token for a specified user and a single cloud desktop under the current account, keeping the original expiration time.
//
// Description:
//
// You must provide Resource and specify both ResourceType and ResourceId when calling this operation. Only a single cloud desktop is supported. You cannot omit Resource to authorize or query all cloud desktops.
//
// AuthorizerId is the WyId of the user and cannot be replaced by a username. BusinessChannel supports Enterprise and Business. If omitted, the default value is Enterprise.
//
// The AuthorizerId and Resource in the query conditions must be consistent with those used when the token was created. This operation only returns the current valid token for the current account, user, and cloud desktop. It does not create a new token, renew the token, or restore or roll back to a historical token. If no valid token is found, the token has expired, or the corresponding cache has been cleared, AccessToken.NotFound is returned. To obtain a new token, call CreateThirdAppAccessToken.
//
// Keep the returned AccessToken secure and avoid displaying the full token in logs or public content.
//
// @param request - GetLatestThirdAppAccessTokenRequest
//
// @return GetLatestThirdAppAccessTokenResponse
func (client *Client) GetLatestThirdAppAccessToken(request *GetLatestThirdAppAccessTokenRequest) (_result *GetLatestThirdAppAccessTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLatestThirdAppAccessTokenResponse{}
	_body, _err := client.GetLatestThirdAppAccessTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a Wuying STS token (temporary identity credential).
//
// @param request - GetStsTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStsTokenResponse
func (client *Client) GetStsTokenWithOptions(request *GetStsTokenRequest, runtime *dara.RuntimeOptions) (_result *GetStsTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.Expiration) {
		body["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExternalId) {
		body["ExternalId"] = request.ExternalId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStsToken"),
		Version:     dara.String("2021-02-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStsTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a Wuying STS token (temporary identity credential).
//
// @param request - GetStsTokenRequest
//
// @return GetStsTokenResponse
func (client *Client) GetStsToken(request *GetStsTokenRequest) (_result *GetStsTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStsTokenResponse{}
	_body, _err := client.GetStsTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
