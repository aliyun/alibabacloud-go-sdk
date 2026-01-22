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
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ims"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a client ID to an OpenID Connect (OIDC) identity provider (IdP).
//
// @param request - AddClientIdToOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddClientIdToOIDCProviderResponse
func (client *Client) AddClientIdToOIDCProviderWithOptions(request *AddClientIdToOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *AddClientIdToOIDCProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.OIDCProviderName) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddClientIdToOIDCProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddClientIdToOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a client ID to an OpenID Connect (OIDC) identity provider (IdP).
//
// @param request - AddClientIdToOIDCProviderRequest
//
// @return AddClientIdToOIDCProviderResponse
func (client *Client) AddClientIdToOIDCProvider(request *AddClientIdToOIDCProviderRequest) (_result *AddClientIdToOIDCProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddClientIdToOIDCProviderResponse{}
	_body, _err := client.AddClientIdToOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a fingerprint to an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to add the fingerprint `902ef2deeb3c5b13ea4c3d5193629309e231****` to the OIDC IdP named `TestOIDCProvider`.
//
// @param request - AddFingerprintToOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFingerprintToOIDCProviderResponse
func (client *Client) AddFingerprintToOIDCProviderWithOptions(request *AddFingerprintToOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *AddFingerprintToOIDCProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Fingerprint) {
		query["Fingerprint"] = request.Fingerprint
	}

	if !dara.IsNil(request.OIDCProviderName) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFingerprintToOIDCProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFingerprintToOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a fingerprint to an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to add the fingerprint `902ef2deeb3c5b13ea4c3d5193629309e231****` to the OIDC IdP named `TestOIDCProvider`.
//
// @param request - AddFingerprintToOIDCProviderRequest
//
// @return AddFingerprintToOIDCProviderResponse
func (client *Client) AddFingerprintToOIDCProvider(request *AddFingerprintToOIDCProviderRequest) (_result *AddFingerprintToOIDCProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFingerprintToOIDCProviderResponse{}
	_body, _err := client.AddFingerprintToOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a Resource Access Management (RAM) user to a RAM user group.
//
// @param request - AddUserToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroupWithOptions(request *AddUserToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUserToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToGroup"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a Resource Access Management (RAM) user to a RAM user group.
//
// @param request - AddUserToGroupRequest
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroup(request *AddUserToGroupRequest) (_result *AddUserToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.AddUserToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a multi-factor authentication (MFA) device to a Resource Access Management (RAM) user.
//
// @param request - BindMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindMFADeviceResponse
func (client *Client) BindMFADeviceWithOptions(request *BindMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *BindMFADeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationCode1) {
		query["AuthenticationCode1"] = request.AuthenticationCode1
	}

	if !dara.IsNil(request.AuthenticationCode2) {
		query["AuthenticationCode2"] = request.AuthenticationCode2
	}

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindMFADevice"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a multi-factor authentication (MFA) device to a Resource Access Management (RAM) user.
//
// @param request - BindMFADeviceRequest
//
// @return BindMFADeviceResponse
func (client *Client) BindMFADevice(request *BindMFADeviceRequest) (_result *BindMFADeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindMFADeviceResponse{}
	_body, _err := client.BindMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the password that is used to log on to the console for a Resource Access Management (RAM) user.
//
// Description:
//
// >  This operation is available only for RAM users. Before you call this operation, make sure that `AllowUserToChangePassword` in [SetSecurityPreference](https://help.aliyun.com/document_detail/43765.html) is set to `True`. The value True indicates that RAM users can manage their passwords.
//
// @param request - ChangePasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangePasswordResponse
func (client *Client) ChangePasswordWithOptions(request *ChangePasswordRequest, runtime *dara.RuntimeOptions) (_result *ChangePasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewPassword) {
		query["NewPassword"] = request.NewPassword
	}

	if !dara.IsNil(request.OldPassword) {
		query["OldPassword"] = request.OldPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangePassword"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangePasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the password that is used to log on to the console for a Resource Access Management (RAM) user.
//
// Description:
//
// >  This operation is available only for RAM users. Before you call this operation, make sure that `AllowUserToChangePassword` in [SetSecurityPreference](https://help.aliyun.com/document_detail/43765.html) is set to `True`. The value True indicates that RAM users can manage their passwords.
//
// @param request - ChangePasswordRequest
//
// @return ChangePasswordResponse
func (client *Client) ChangePassword(request *ChangePasswordRequest) (_result *ChangePasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangePasswordResponse{}
	_body, _err := client.ChangePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - CreateAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessKeyResponse
func (client *Client) CreateAccessKeyWithOptions(request *CreateAccessKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessKey"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - CreateAccessKeyRequest
//
// @return CreateAccessKeyResponse
func (client *Client) CreateAccessKey(request *CreateAccessKeyRequest) (_result *CreateAccessKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessKeyResponse{}
	_body, _err := client.CreateAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application secret for an application.
//
// @param request - CreateAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppSecretResponse
func (client *Client) CreateAppSecretWithOptions(request *CreateAppSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateAppSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppSecret"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application secret for an application.
//
// @param request - CreateAppSecretRequest
//
// @return CreateAppSecretResponse
func (client *Client) CreateAppSecret(request *CreateAppSecretRequest) (_result *CreateAppSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppSecretResponse{}
	_body, _err := client.CreateAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application.
//
// @param request - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessTokenValidity) {
		query["AccessTokenValidity"] = request.AccessTokenValidity
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.IsMultiTenant) {
		query["IsMultiTenant"] = request.IsMultiTenant
	}

	if !dara.IsNil(request.PredefinedScopes) {
		query["PredefinedScopes"] = request.PredefinedScopes
	}

	if !dara.IsNil(request.ProtocolVersion) {
		query["ProtocolVersion"] = request.ProtocolVersion
	}

	if !dara.IsNil(request.RedirectUris) {
		query["RedirectUris"] = request.RedirectUris
	}

	if !dara.IsNil(request.RefreshTokenValidity) {
		query["RefreshTokenValidity"] = request.RefreshTokenValidity
	}

	if !dara.IsNil(request.RequiredScopes) {
		query["RequiredScopes"] = request.RequiredScopes
	}

	if !dara.IsNil(request.SecretRequired) {
		query["SecretRequired"] = request.SecretRequired
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application.
//
// @param request - CreateApplicationRequest
//
// @return CreateApplicationResponse
func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) user group.
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comments) {
		query["Comments"] = request.Comments
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroup"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) user group.
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables logon to the console for a Resource Access Management (RAM) user.
//
// @param request - CreateLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoginProfileResponse
func (client *Client) CreateLoginProfileWithOptions(request *CreateLoginProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateLoginProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MFABindRequired) {
		query["MFABindRequired"] = request.MFABindRequired
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordResetRequired) {
		query["PasswordResetRequired"] = request.PasswordResetRequired
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoginProfile"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables logon to the console for a Resource Access Management (RAM) user.
//
// @param request - CreateLoginProfileRequest
//
// @return CreateLoginProfileResponse
func (client *Client) CreateLoginProfile(request *CreateLoginProfileRequest) (_result *CreateLoginProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoginProfileResponse{}
	_body, _err := client.CreateLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an OpenID Connect (OIDC) identity provider (IdP) to configure a trust relationship between Alibaba Cloud and an external IdP. This topic provides an example on how to create an IdP named TestOIDCProvider to configure a trust relationship between the external IdP Okta and Alibaba Cloud.
//
// Description:
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the information such as the URL of the issuer, the fingerprints of HTTPS certificate authority (CA) certificates, and the client IDs are obtained from an external IdP, such as Google Workspace or Okta.
//
// ### [](#)Limits
//
//   - You can create a maximum of 100 OIDC IdPs in an Alibaba Cloud account.
//
//   - You can add a maximum of 50 client IDs to an OIDC IdP.
//
//   - You can add a maximum of five fingerprints to an OIDC IdP.
//
// ### [](#)Operation description
//
// This topic provides an example on how to create an IdP named `TestOIDCProvider` to configure a trust relationship between the external IdP and Alibaba Cloud.
//
// @param request - CreateOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOIDCProviderResponse
func (client *Client) CreateOIDCProviderWithOptions(request *CreateOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateOIDCProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIds) {
		query["ClientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Fingerprints) {
		query["Fingerprints"] = request.Fingerprints
	}

	if !dara.IsNil(request.IssuanceLimitTime) {
		query["IssuanceLimitTime"] = request.IssuanceLimitTime
	}

	if !dara.IsNil(request.IssuerUrl) {
		query["IssuerUrl"] = request.IssuerUrl
	}

	if !dara.IsNil(request.OIDCProviderName) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOIDCProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an OpenID Connect (OIDC) identity provider (IdP) to configure a trust relationship between Alibaba Cloud and an external IdP. This topic provides an example on how to create an IdP named TestOIDCProvider to configure a trust relationship between the external IdP Okta and Alibaba Cloud.
//
// Description:
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the information such as the URL of the issuer, the fingerprints of HTTPS certificate authority (CA) certificates, and the client IDs are obtained from an external IdP, such as Google Workspace or Okta.
//
// ### [](#)Limits
//
//   - You can create a maximum of 100 OIDC IdPs in an Alibaba Cloud account.
//
//   - You can add a maximum of 50 client IDs to an OIDC IdP.
//
//   - You can add a maximum of five fingerprints to an OIDC IdP.
//
// ### [](#)Operation description
//
// This topic provides an example on how to create an IdP named `TestOIDCProvider` to configure a trust relationship between the external IdP and Alibaba Cloud.
//
// @param request - CreateOIDCProviderRequest
//
// @return CreateOIDCProviderResponse
func (client *Client) CreateOIDCProvider(request *CreateOIDCProviderRequest) (_result *CreateOIDCProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOIDCProviderResponse{}
	_body, _err := client.CreateOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - CreateSAMLProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSAMLProviderResponse
func (client *Client) CreateSAMLProviderWithOptions(request *CreateSAMLProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateSAMLProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthnSignAlgo) {
		query["AuthnSignAlgo"] = request.AuthnSignAlgo
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EncodedSAMLMetadataDocument) {
		query["EncodedSAMLMetadataDocument"] = request.EncodedSAMLMetadataDocument
	}

	if !dara.IsNil(request.SAMLProviderName) {
		query["SAMLProviderName"] = request.SAMLProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSAMLProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSAMLProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - CreateSAMLProviderRequest
//
// @return CreateSAMLProviderResponse
func (client *Client) CreateSAMLProvider(request *CreateSAMLProviderRequest) (_result *CreateSAMLProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSAMLProviderResponse{}
	_body, _err := client.CreateSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a RAM user.
//
// Description:
//
// This topic provides an example on how to create a RAM user named `test`.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comments) {
		query["Comments"] = request.Comments
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.MobilePhone) {
		query["MobilePhone"] = request.MobilePhone
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a RAM user.
//
// Description:
//
// This topic provides an example on how to create a RAM user named `test`.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a virtual multi-factor authentication (MFA) device.
//
// @param request - CreateVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualMFADeviceResponse
func (client *Client) CreateVirtualMFADeviceWithOptions(request *CreateVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *CreateVirtualMFADeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VirtualMFADeviceName) {
		query["VirtualMFADeviceName"] = request.VirtualMFADeviceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVirtualMFADevice"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual multi-factor authentication (MFA) device.
//
// @param request - CreateVirtualMFADeviceRequest
//
// @return CreateVirtualMFADeviceResponse
func (client *Client) CreateVirtualMFADevice(request *CreateVirtualMFADeviceRequest) (_result *CreateVirtualMFADeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVirtualMFADeviceResponse{}
	_body, _err := client.CreateVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - DeleteAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessKeyResponse
func (client *Client) DeleteAccessKeyWithOptions(request *DeleteAccessKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessKey"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - DeleteAccessKeyRequest
//
// @return DeleteAccessKeyResponse
func (client *Client) DeleteAccessKey(request *DeleteAccessKeyRequest) (_result *DeleteAccessKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessKeyResponse{}
	_body, _err := client.DeleteAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific AccessKey pair that belongs to a Resource Access Management (RAM) user from the recycle bin.
//
// @param request - DeleteAccessKeyInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessKeyInRecycleBinResponse
func (client *Client) DeleteAccessKeyInRecycleBinWithOptions(request *DeleteAccessKeyInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessKeyInRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessKeyInRecycleBin"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessKeyInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific AccessKey pair that belongs to a Resource Access Management (RAM) user from the recycle bin.
//
// @param request - DeleteAccessKeyInRecycleBinRequest
//
// @return DeleteAccessKeyInRecycleBinResponse
func (client *Client) DeleteAccessKeyInRecycleBin(request *DeleteAccessKeyInRecycleBinRequest) (_result *DeleteAccessKeyInRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessKeyInRecycleBinResponse{}
	_body, _err := client.DeleteAccessKeyInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the application secret of an application.
//
// @param request - DeleteAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppSecretResponse
func (client *Client) DeleteAppSecretWithOptions(request *DeleteAppSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppSecretId) {
		query["AppSecretId"] = request.AppSecretId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppSecret"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the application secret of an application.
//
// @param request - DeleteAppSecretRequest
//
// @return DeleteAppSecretResponse
func (client *Client) DeleteAppSecret(request *DeleteAppSecretRequest) (_result *DeleteAppSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppSecretResponse{}
	_body, _err := client.DeleteAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// @param request - DeleteApplicationRequest
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user group.
//
// Description:
//
// Before you delete a RAM user group, make sure that no policies are attached to the group and no RAM users are included in the group.
//
// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroup"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user group.
//
// Description:
//
// Before you delete a RAM user group, make sure that no policies are attached to the group and no RAM users are included in the group.
//
// @param request - DeleteGroupRequest
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables logon to the console for a Resource Access Management (RAM) user.
//
// @param request - DeleteLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoginProfileResponse
func (client *Client) DeleteLoginProfileWithOptions(request *DeleteLoginProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteLoginProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLoginProfile"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables logon to the console for a Resource Access Management (RAM) user.
//
// @param request - DeleteLoginProfileRequest
//
// @return DeleteLoginProfileResponse
func (client *Client) DeleteLoginProfile(request *DeleteLoginProfileRequest) (_result *DeleteLoginProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLoginProfileResponse{}
	_body, _err := client.DeleteLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the OIDC IdP named `TestOIDCProvider`.
//
// @param request - DeleteOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOIDCProviderResponse
func (client *Client) DeleteOIDCProviderWithOptions(request *DeleteOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteOIDCProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OIDCProviderName) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOIDCProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the OIDC IdP named `TestOIDCProvider`.
//
// @param request - DeleteOIDCProviderRequest
//
// @return DeleteOIDCProviderResponse
func (client *Client) DeleteOIDCProvider(request *DeleteOIDCProviderRequest) (_result *DeleteOIDCProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOIDCProviderResponse{}
	_body, _err := client.DeleteOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a passkey for a Resource Access Management (RAM) user.
//
// @param request - DeletePasskeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePasskeyResponse
func (client *Client) DeletePasskeyWithOptions(request *DeletePasskeyRequest, runtime *dara.RuntimeOptions) (_result *DeletePasskeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PasskeyId) {
		query["PasskeyId"] = request.PasskeyId
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePasskey"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePasskeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a passkey for a Resource Access Management (RAM) user.
//
// @param request - DeletePasskeyRequest
//
// @return DeletePasskeyResponse
func (client *Client) DeletePasskey(request *DeletePasskeyRequest) (_result *DeletePasskeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePasskeyResponse{}
	_body, _err := client.DeletePasskeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - DeleteSAMLProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSAMLProviderResponse
func (client *Client) DeleteSAMLProviderWithOptions(request *DeleteSAMLProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteSAMLProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SAMLProviderName) {
		query["SAMLProviderName"] = request.SAMLProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSAMLProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSAMLProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - DeleteSAMLProviderRequest
//
// @return DeleteSAMLProviderResponse
func (client *Client) DeleteSAMLProvider(request *DeleteSAMLProviderRequest) (_result *DeleteSAMLProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSAMLProviderResponse{}
	_body, _err := client.DeleteSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user.
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific Resource Access Management (RAM) user from the recycle bin.
//
// @param request - DeleteUserInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserInRecycleBinResponse
func (client *Client) DeleteUserInRecycleBinWithOptions(request *DeleteUserInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserInRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserInRecycleBin"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific Resource Access Management (RAM) user from the recycle bin.
//
// @param request - DeleteUserInRecycleBinRequest
//
// @return DeleteUserInRecycleBinResponse
func (client *Client) DeleteUserInRecycleBin(request *DeleteUserInRecycleBinRequest) (_result *DeleteUserInRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserInRecycleBinResponse{}
	_body, _err := client.DeleteUserInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a multi-factor authentication (MFA) device.
//
// @param request - DeleteVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVirtualMFADevice"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a multi-factor authentication (MFA) device.
//
// @param request - DeleteVirtualMFADeviceRequest
//
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls an external application or an internal application of the ServerApp type.
//
// Description:
//
// If you want to call this operation to uninstall an internal application, the type of the internal application must be **ServerApp**. Otherwise, an error occurs when you call this operation.
//
// >  For **internal applications**, only internal applications of the ServerApp type need to be **installed or provisioned**. Therefore, only internal applications of the ServerApp type **can be uninstalled**. Internal applications of the WebApp and NativeApp types **do not need to and cannot be uninstalled**.
//
// @param request - DeprovisionApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprovisionApplicationResponse
func (client *Client) DeprovisionApplicationWithOptions(request *DeprovisionApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeprovisionApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeprovisionApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeprovisionApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls an external application or an internal application of the ServerApp type.
//
// Description:
//
// If you want to call this operation to uninstall an internal application, the type of the internal application must be **ServerApp**. Otherwise, an error occurs when you call this operation.
//
// >  For **internal applications**, only internal applications of the ServerApp type need to be **installed or provisioned**. Therefore, only internal applications of the ServerApp type **can be uninstalled**. Internal applications of the WebApp and NativeApp types **do not need to and cannot be uninstalled**.
//
// @param request - DeprovisionApplicationRequest
//
// @return DeprovisionApplicationResponse
func (client *Client) DeprovisionApplication(request *DeprovisionApplicationRequest) (_result *DeprovisionApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeprovisionApplicationResponse{}
	_body, _err := client.DeprovisionApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an installed external application.
//
// @param request - DeprovisionExternalApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprovisionExternalApplicationResponse
func (client *Client) DeprovisionExternalApplicationWithOptions(request *DeprovisionExternalApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeprovisionExternalApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeprovisionExternalApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeprovisionExternalApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an installed external application.
//
// @param request - DeprovisionExternalApplicationRequest
//
// @return DeprovisionExternalApplicationResponse
func (client *Client) DeprovisionExternalApplication(request *DeprovisionExternalApplicationRequest) (_result *DeprovisionExternalApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeprovisionExternalApplicationResponse{}
	_body, _err := client.DeprovisionExternalApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds and deletes a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - DisableVirtualMFARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVirtualMFAResponse
func (client *Client) DisableVirtualMFAWithOptions(request *DisableVirtualMFARequest, runtime *dara.RuntimeOptions) (_result *DisableVirtualMFAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableVirtualMFA"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableVirtualMFAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds and deletes a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - DisableVirtualMFARequest
//
// @return DisableVirtualMFAResponse
func (client *Client) DisableVirtualMFA(request *DisableVirtualMFARequest) (_result *DisableVirtualMFAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableVirtualMFAResponse{}
	_body, _err := client.DisableVirtualMFAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates the user credential report of an Alibaba Cloud account.
//
// @param request - GenerateCredentialReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCredentialReportResponse
func (client *Client) GenerateCredentialReportWithOptions(runtime *dara.RuntimeOptions) (_result *GenerateCredentialReportResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateCredentialReport"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateCredentialReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates the user credential report of an Alibaba Cloud account.
//
// @return GenerateCredentialReportResponse
func (client *Client) GenerateCredentialReport() (_result *GenerateCredentialReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateCredentialReportResponse{}
	_body, _err := client.GenerateCredentialReportWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a check report for Cloud Governance.
//
// @param request - GenerateGovernanceReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateGovernanceReportResponse
func (client *Client) GenerateGovernanceReportWithOptions(runtime *dara.RuntimeOptions) (_result *GenerateGovernanceReportResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateGovernanceReport"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateGovernanceReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a check report for Cloud Governance.
//
// @return GenerateGovernanceReportResponse
func (client *Client) GenerateGovernanceReport() (_result *GenerateGovernanceReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateGovernanceReportResponse{}
	_body, _err := client.GenerateGovernanceReportWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specific AccessKey pair of a Resource Access Management (RAM) user in the recycle bin.
//
// @param request - GetAccessKeyInfoInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyInfoInRecycleBinResponse
func (client *Client) GetAccessKeyInfoInRecycleBinWithOptions(request *GetAccessKeyInfoInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyInfoInRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyInfoInRecycleBin"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyInfoInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specific AccessKey pair of a Resource Access Management (RAM) user in the recycle bin.
//
// @param request - GetAccessKeyInfoInRecycleBinRequest
//
// @return GetAccessKeyInfoInRecycleBinResponse
func (client *Client) GetAccessKeyInfoInRecycleBin(request *GetAccessKeyInfoInRecycleBinRequest) (_result *GetAccessKeyInfoInRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessKeyInfoInRecycleBinResponse{}
	_body, _err := client.GetAccessKeyInfoInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the time when an AccessKey pair was used for the last time.
//
// @param request - GetAccessKeyLastUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedResponse
func (client *Client) GetAccessKeyLastUsedWithOptions(request *GetAccessKeyLastUsedRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsed"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time when an AccessKey pair was used for the last time.
//
// @param request - GetAccessKeyLastUsedRequest
//
// @return GetAccessKeyLastUsedResponse
func (client *Client) GetAccessKeyLastUsed(request *GetAccessKeyLastUsedRequest) (_result *GetAccessKeyLastUsedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedResponse{}
	_body, _err := client.GetAccessKeyLastUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the multi-factor authentication (MFA) devices of an Alibaba Cloud account.
//
// @param request - GetAccountMFAInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountMFAInfoResponse
func (client *Client) GetAccountMFAInfoWithOptions(runtime *dara.RuntimeOptions) (_result *GetAccountMFAInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountMFAInfo"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountMFAInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the multi-factor authentication (MFA) devices of an Alibaba Cloud account.
//
// @return GetAccountMFAInfoResponse
func (client *Client) GetAccountMFAInfo() (_result *GetAccountMFAInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountMFAInfoResponse{}
	_body, _err := client.GetAccountMFAInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security report of an Alibaba Cloud account.
//
// @param request - GetAccountSecurityPracticeReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountSecurityPracticeReportResponse
func (client *Client) GetAccountSecurityPracticeReportWithOptions(runtime *dara.RuntimeOptions) (_result *GetAccountSecurityPracticeReportResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountSecurityPracticeReport"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountSecurityPracticeReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security report of an Alibaba Cloud account.
//
// @return GetAccountSecurityPracticeReportResponse
func (client *Client) GetAccountSecurityPracticeReport() (_result *GetAccountSecurityPracticeReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountSecurityPracticeReportResponse{}
	_body, _err := client.GetAccountSecurityPracticeReportWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the overview information about an Alibaba Cloud account.
//
// @param request - GetAccountSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountSummaryResponse
func (client *Client) GetAccountSummaryWithOptions(runtime *dara.RuntimeOptions) (_result *GetAccountSummaryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountSummary"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the overview information about an Alibaba Cloud account.
//
// @return GetAccountSummaryResponse
func (client *Client) GetAccountSummary() (_result *GetAccountSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountSummaryResponse{}
	_body, _err := client.GetAccountSummaryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an application secret.
//
// @param request - GetAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSecretResponse
func (client *Client) GetAppSecretWithOptions(request *GetAppSecretRequest, runtime *dara.RuntimeOptions) (_result *GetAppSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppSecretId) {
		query["AppSecretId"] = request.AppSecretId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppSecret"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an application secret.
//
// @param request - GetAppSecretRequest
//
// @return GetAppSecretResponse
func (client *Client) GetAppSecret(request *GetAppSecretRequest) (_result *GetAppSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppSecretResponse{}
	_body, _err := client.GetAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration information about an application.
//
// Description:
//
// This topic provides an example on how to query the configurations of an application named `472457090344041****`.
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration information about an application.
//
// Description:
//
// This topic provides an example on how to query the configurations of an application named `472457090344041****`.
//
// @param request - GetApplicationRequest
//
// @return GetApplicationResponse
func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries installation information about a specified installed application.
//
// @param request - GetApplicationProvisionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisionInfoResponse
func (client *Client) GetApplicationProvisionInfoWithOptions(request *GetApplicationProvisionInfoRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationProvisionInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationProvisionInfo"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationProvisionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries installation information about a specified installed application.
//
// @param request - GetApplicationProvisionInfoRequest
//
// @return GetApplicationProvisionInfoResponse
func (client *Client) GetApplicationProvisionInfo(request *GetApplicationProvisionInfoRequest) (_result *GetApplicationProvisionInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationProvisionInfoResponse{}
	_body, _err := client.GetApplicationProvisionInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the user credential reports of an Alibaba Cloud account.
//
// @param request - GetCredentialReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialReportResponse
func (client *Client) GetCredentialReportWithOptions(request *GetCredentialReportRequest, runtime *dara.RuntimeOptions) (_result *GetCredentialReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCredentialReport"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCredentialReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user credential reports of an Alibaba Cloud account.
//
// @param request - GetCredentialReportRequest
//
// @return GetCredentialReportResponse
func (client *Client) GetCredentialReport(request *GetCredentialReportRequest) (_result *GetCredentialReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCredentialReportResponse{}
	_body, _err := client.GetCredentialReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default domain name of an Alibaba Cloud account.
//
// @param request - GetDefaultDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultDomainResponse
func (client *Client) GetDefaultDomainWithOptions(runtime *dara.RuntimeOptions) (_result *GetDefaultDomainResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetDefaultDomain"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDefaultDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default domain name of an Alibaba Cloud account.
//
// @return GetDefaultDomainResponse
func (client *Client) GetDefaultDomain() (_result *GetDefaultDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDefaultDomainResponse{}
	_body, _err := client.GetDefaultDomainWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about an installed external application.
//
// @param request - GetExternalApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExternalApplicationResponse
func (client *Client) GetExternalApplicationWithOptions(request *GetExternalApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetExternalApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExternalApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExternalApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an installed external application.
//
// @param request - GetExternalApplicationRequest
//
// @return GetExternalApplicationResponse
func (client *Client) GetExternalApplication(request *GetExternalApplicationRequest) (_result *GetExternalApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExternalApplicationResponse{}
	_body, _err := client.GetExternalApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户的单项ram治理报告
//
// @param request - GetGovernanceItemReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGovernanceItemReportResponse
func (client *Client) GetGovernanceItemReportWithOptions(request *GetGovernanceItemReportRequest, runtime *dara.RuntimeOptions) (_result *GetGovernanceItemReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GovernanceItemType) {
		query["GovernanceItemType"] = request.GovernanceItemType
	}

	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGovernanceItemReport"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGovernanceItemReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户的单项ram治理报告
//
// @param request - GetGovernanceItemReportRequest
//
// @return GetGovernanceItemReportResponse
func (client *Client) GetGovernanceItemReport(request *GetGovernanceItemReportRequest) (_result *GetGovernanceItemReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGovernanceItemReportResponse{}
	_body, _err := client.GetGovernanceItemReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询成熟度报告状态
//
// @param request - GetGovernanceReportStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGovernanceReportStatusResponse
func (client *Client) GetGovernanceReportStatusWithOptions(runtime *dara.RuntimeOptions) (_result *GetGovernanceReportStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetGovernanceReportStatus"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGovernanceReportStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询成熟度报告状态
//
// @return GetGovernanceReportStatusResponse
func (client *Client) GetGovernanceReportStatus() (_result *GetGovernanceReportStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGovernanceReportStatusResponse{}
	_body, _err := client.GetGovernanceReportStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a Resource Access Management (RAM) user group.
//
// @param request - GetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroup"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Resource Access Management (RAM) user group.
//
// @param request - GetGroupRequest
//
// @return GetGroupResponse
func (client *Client) GetGroup(request *GetGroupRequest) (_result *GetGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logon configurations of a Resource Access Management (RAM) user.
//
// @param request - GetLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginProfileResponse
func (client *Client) GetLoginProfileWithOptions(request *GetLoginProfileRequest, runtime *dara.RuntimeOptions) (_result *GetLoginProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoginProfile"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logon configurations of a Resource Access Management (RAM) user.
//
// @param request - GetLoginProfileRequest
//
// @return GetLoginProfileResponse
func (client *Client) GetLoginProfile(request *GetLoginProfileRequest) (_result *GetLoginProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLoginProfileResponse{}
	_body, _err := client.GetLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an OIDC IdP.
//
// Description:
//
// ###
//
// This topic provides an example on how to query the information about an OpenID Connect (OIDC) identity provider (IdP) named `TestOIDCProvider`.
//
// @param request - GetOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOIDCProviderResponse
func (client *Client) GetOIDCProviderWithOptions(request *GetOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *GetOIDCProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OIDCProviderName) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOIDCProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an OIDC IdP.
//
// Description:
//
// ###
//
// This topic provides an example on how to query the information about an OpenID Connect (OIDC) identity provider (IdP) named `TestOIDCProvider`.
//
// @param request - GetOIDCProviderRequest
//
// @return GetOIDCProviderResponse
func (client *Client) GetOIDCProvider(request *GetOIDCProviderRequest) (_result *GetOIDCProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOIDCProviderResponse{}
	_body, _err := client.GetOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the password policy for RAM users.
//
// @param request - GetPasswordPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordPolicyResponse
func (client *Client) GetPasswordPolicyWithOptions(runtime *dara.RuntimeOptions) (_result *GetPasswordPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetPasswordPolicy"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the password policy for RAM users.
//
// @return GetPasswordPolicyResponse
func (client *Client) GetPasswordPolicy() (_result *GetPasswordPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.GetPasswordPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - GetSAMLProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSAMLProviderResponse
func (client *Client) GetSAMLProviderWithOptions(request *GetSAMLProviderRequest, runtime *dara.RuntimeOptions) (_result *GetSAMLProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SAMLProviderName) {
		query["SAMLProviderName"] = request.SAMLProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSAMLProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSAMLProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - GetSAMLProviderRequest
//
// @return GetSAMLProviderResponse
func (client *Client) GetSAMLProvider(request *GetSAMLProviderRequest) (_result *GetSAMLProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSAMLProviderResponse{}
	_body, _err := client.GetSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security preferences for RAM users.
//
// @param request - GetSecurityPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityPreferenceResponse
func (client *Client) GetSecurityPreferenceWithOptions(runtime *dara.RuntimeOptions) (_result *GetSecurityPreferenceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityPreference"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security preferences for RAM users.
//
// @return GetSecurityPreferenceResponse
func (client *Client) GetSecurityPreference() (_result *GetSecurityPreferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSecurityPreferenceResponse{}
	_body, _err := client.GetSecurityPreferenceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a RAM user.
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a RAM user.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specific Resource Access Management (RAM) user in the recycle bin.
//
// @param request - GetUserInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInRecycleBinResponse
func (client *Client) GetUserInRecycleBinWithOptions(request *GetUserInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *GetUserInRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserInRecycleBin"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specific Resource Access Management (RAM) user in the recycle bin.
//
// @param request - GetUserInRecycleBinRequest
//
// @return GetUserInRecycleBinResponse
func (client *Client) GetUserInRecycleBin(request *GetUserInRecycleBinRequest) (_result *GetUserInRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserInRecycleBinResponse{}
	_body, _err := client.GetUserInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the multi-factor authentication (MFA) device that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetUserMFAInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserMFAInfoResponse
func (client *Client) GetUserMFAInfoWithOptions(request *GetUserMFAInfoRequest, runtime *dara.RuntimeOptions) (_result *GetUserMFAInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserMFAInfo"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserMFAInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the multi-factor authentication (MFA) device that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetUserMFAInfoRequest
//
// @return GetUserMFAInfoResponse
func (client *Client) GetUserMFAInfo(request *GetUserMFAInfoRequest) (_result *GetUserMFAInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserMFAInfoResponse{}
	_body, _err := client.GetUserMFAInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of user-based single sign-on (SSO).
//
// @param request - GetUserSsoSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserSsoSettingsResponse
func (client *Client) GetUserSsoSettingsWithOptions(runtime *dara.RuntimeOptions) (_result *GetUserSsoSettingsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserSsoSettings"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserSsoSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of user-based single sign-on (SSO).
//
// @return GetUserSsoSettingsResponse
func (client *Client) GetUserSsoSettings() (_result *GetUserSsoSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserSsoSettingsResponse{}
	_body, _err := client.GetUserSsoSettingsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the mobile phone or email that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetVerificationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVerificationInfoResponse
func (client *Client) GetVerificationInfoWithOptions(request *GetVerificationInfoRequest, runtime *dara.RuntimeOptions) (_result *GetVerificationInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVerificationInfo"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVerificationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the mobile phone or email that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetVerificationInfoRequest
//
// @return GetVerificationInfoResponse
func (client *Client) GetVerificationInfo(request *GetVerificationInfoRequest) (_result *GetVerificationInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVerificationInfoResponse{}
	_body, _err := client.GetVerificationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs of an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - ListAccessKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessKeysResponse
func (client *Client) ListAccessKeysWithOptions(request *ListAccessKeysRequest, runtime *dara.RuntimeOptions) (_result *ListAccessKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessKeys"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs of an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - ListAccessKeysRequest
//
// @return ListAccessKeysResponse
func (client *Client) ListAccessKeys(request *ListAccessKeysRequest) (_result *ListAccessKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccessKeysResponse{}
	_body, _err := client.ListAccessKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs of a specific Resource Access Management (RAM) user in the recycle bin.
//
// @param request - ListAccessKeysInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessKeysInRecycleBinResponse
func (client *Client) ListAccessKeysInRecycleBinWithOptions(request *ListAccessKeysInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *ListAccessKeysInRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessKeysInRecycleBin"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessKeysInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs of a specific Resource Access Management (RAM) user in the recycle bin.
//
// @param request - ListAccessKeysInRecycleBinRequest
//
// @return ListAccessKeysInRecycleBinResponse
func (client *Client) ListAccessKeysInRecycleBin(request *ListAccessKeysInRecycleBinRequest) (_result *ListAccessKeysInRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccessKeysInRecycleBinResponse{}
	_body, _err := client.ListAccessKeysInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the secret IDs of an application.
//
// @param request - ListAppSecretIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppSecretIdsResponse
func (client *Client) ListAppSecretIdsWithOptions(request *ListAppSecretIdsRequest, runtime *dara.RuntimeOptions) (_result *ListAppSecretIdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppSecretIds"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppSecretIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the secret IDs of an application.
//
// @param request - ListAppSecretIdsRequest
//
// @return ListAppSecretIdsResponse
func (client *Client) ListAppSecretIds(request *ListAppSecretIdsRequest) (_result *ListAppSecretIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppSecretIdsResponse{}
	_body, _err := client.ListAppSecretIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries installation information about all installed applications.
//
// @param request - ListApplicationProvisionInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationProvisionInfosResponse
func (client *Client) ListApplicationProvisionInfosWithOptions(request *ListApplicationProvisionInfosRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationProvisionInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationProvisionInfos"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationProvisionInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries installation information about all installed applications.
//
// @param request - ListApplicationProvisionInfosRequest
//
// @return ListApplicationProvisionInfosResponse
func (client *Client) ListApplicationProvisionInfos(request *ListApplicationProvisionInfosRequest) (_result *ListApplicationProvisionInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationProvisionInfosResponse{}
	_body, _err := client.ListApplicationProvisionInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the created applications.
//
// Description:
//
// This topic provides an example on how to query the applications within the current account. The returned result shows that only one application named `myapp` belongs to the current account.
//
// @param request - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithOptions(runtime *dara.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplications"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the created applications.
//
// Description:
//
// This topic provides an example on how to query the applications within the current account. The returned result shows that only one application named `myapp` belongs to the current account.
//
// @return ListApplicationsResponse
func (client *Client) ListApplications() (_result *ListApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all installed external applications.
//
// @param request - ListExternalApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExternalApplicationsResponse
func (client *Client) ListExternalApplicationsWithOptions(runtime *dara.RuntimeOptions) (_result *ListExternalApplicationsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListExternalApplications"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExternalApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all installed external applications.
//
// @return ListExternalApplicationsResponse
func (client *Client) ListExternalApplications() (_result *ListExternalApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExternalApplicationsResponse{}
	_body, _err := client.ListExternalApplicationsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user groups.
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroups"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user groups.
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) user groups to which a RAM user belongs.
//
// @param request - ListGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithOptions(request *ListGroupsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupsForUser"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) user groups to which a RAM user belongs.
//
// @param request - ListGroupsForUserRequest
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (_result *ListGroupsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.ListGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries OIDC IdPs.
//
// Description:
//
// ###
//
// This topic provides an example on how to query all OpenID Connect (OIDC) identity providers (IdPs) within your Alibaba Cloud account. The response shows that your Alibaba Cloud account has only one OIDC IdP named `TestOIDCProvider`.
//
// @param request - ListOIDCProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOIDCProvidersResponse
func (client *Client) ListOIDCProvidersWithOptions(request *ListOIDCProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListOIDCProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOIDCProviders"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOIDCProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries OIDC IdPs.
//
// Description:
//
// ###
//
// This topic provides an example on how to query all OpenID Connect (OIDC) identity providers (IdPs) within your Alibaba Cloud account. The response shows that your Alibaba Cloud account has only one OIDC IdP named `TestOIDCProvider`.
//
// @param request - ListOIDCProvidersRequest
//
// @return ListOIDCProvidersResponse
func (client *Client) ListOIDCProviders(request *ListOIDCProvidersRequest) (_result *ListOIDCProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOIDCProvidersResponse{}
	_body, _err := client.ListOIDCProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the passkeys that are bound to a Resource Access Management (RAM) user.
//
// @param request - ListPasskeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPasskeysResponse
func (client *Client) ListPasskeysWithOptions(request *ListPasskeysRequest, runtime *dara.RuntimeOptions) (_result *ListPasskeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPasskeys"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPasskeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the passkeys that are bound to a Resource Access Management (RAM) user.
//
// @param request - ListPasskeysRequest
//
// @return ListPasskeysResponse
func (client *Client) ListPasskeys(request *ListPasskeysRequest) (_result *ListPasskeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPasskeysResponse{}
	_body, _err := client.ListPasskeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries predefined application permissions.
//
// @param request - ListPredefinedScopesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPredefinedScopesResponse
func (client *Client) ListPredefinedScopesWithOptions(request *ListPredefinedScopesRequest, runtime *dara.RuntimeOptions) (_result *ListPredefinedScopesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppType) {
		query["AppType"] = request.AppType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPredefinedScopes"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPredefinedScopesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries predefined application permissions.
//
// @param request - ListPredefinedScopesRequest
//
// @return ListPredefinedScopesResponse
func (client *Client) ListPredefinedScopes(request *ListPredefinedScopesRequest) (_result *ListPredefinedScopesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPredefinedScopesResponse{}
	_body, _err := client.ListPredefinedScopesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all metric values in the most recent governance check.
//
// @param request - ListRecentGovernanceMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentGovernanceMetricsResponse
func (client *Client) ListRecentGovernanceMetricsWithOptions(runtime *dara.RuntimeOptions) (_result *ListRecentGovernanceMetricsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecentGovernanceMetrics"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecentGovernanceMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all metric values in the most recent governance check.
//
// @return ListRecentGovernanceMetricsResponse
func (client *Client) ListRecentGovernanceMetrics() (_result *ListRecentGovernanceMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecentGovernanceMetricsResponse{}
	_body, _err := client.ListRecentGovernanceMetricsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about identity providers (IdPs) for role-based single sign-on (SSO).
//
// @param request - ListSAMLProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSAMLProvidersResponse
func (client *Client) ListSAMLProvidersWithOptions(request *ListSAMLProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListSAMLProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSAMLProviders"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSAMLProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about identity providers (IdPs) for role-based single sign-on (SSO).
//
// @param request - ListSAMLProvidersRequest
//
// @return ListSAMLProvidersResponse
func (client *Client) ListSAMLProviders(request *ListSAMLProvidersRequest) (_result *ListSAMLProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSAMLProvidersResponse{}
	_body, _err := client.ListSAMLProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added resources.
//
// Description:
//
// ###
//
// You must specify at least one of the following parameters or parameter pairs in a request to determine a query object:
//
//   - `ResourceId.N`
//
//   - `Tag.N.Key`
//
//   - `Tag.N.Key` and `Tag.N.Value`
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourcePrincipalName) {
		query["ResourcePrincipalName"] = request.ResourcePrincipalName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added resources.
//
// Description:
//
// ###
//
// You must specify at least one of the following parameters or parameter pairs in a request to determine a query object:
//
//   - `ResourceId.N`
//
//   - `Tag.N.Key`
//
//   - `Tag.N.Key` and `Tag.N.Value`
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information about all Resource Access Management (RAM) users.
//
// Description:
//
// You can call the following API operations to query information about all RAM users:
//
//   - ListUsers: queries the details of all RAM users.
//
//   - ListUserBasicInfos: queries the basic information about all RAM users. The basic information includes only the logon names (`UserPrincipalName`), display names (`DisplayName`), and user IDs (`UserId`).
//
// @param request - ListUserBasicInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserBasicInfosResponse
func (client *Client) ListUserBasicInfosWithOptions(request *ListUserBasicInfosRequest, runtime *dara.RuntimeOptions) (_result *ListUserBasicInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserBasicInfos"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserBasicInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about all Resource Access Management (RAM) users.
//
// Description:
//
// You can call the following API operations to query information about all RAM users:
//
//   - ListUsers: queries the details of all RAM users.
//
//   - ListUserBasicInfos: queries the basic information about all RAM users. The basic information includes only the logon names (`UserPrincipalName`), display names (`DisplayName`), and user IDs (`UserId`).
//
// @param request - ListUserBasicInfosRequest
//
// @return ListUserBasicInfosResponse
func (client *Client) ListUserBasicInfos(request *ListUserBasicInfosRequest) (_result *ListUserBasicInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserBasicInfosResponse{}
	_body, _err := client.ListUserBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all Resource Access Management (RAM) users.
//
// Description:
//
// ### [](#)
//
// You can call the following API operations to query the details of all RAM users:
//
//   - ListUsers: queries the details of all RAM users.
//
//   - ListUserBasicInfos: queries the basic information about all RAM users. The basic information includes only the logon names (`UserPrincipalName`), display names (`DisplayName`), and user IDs (`UserId`).
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all Resource Access Management (RAM) users.
//
// Description:
//
// ### [](#)
//
// You can call the following API operations to query the details of all RAM users:
//
//   - ListUsers: queries the details of all RAM users.
//
//   - ListUserBasicInfos: queries the basic information about all RAM users. The basic information includes only the logon names (`UserPrincipalName`), display names (`DisplayName`), and user IDs (`UserId`).
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) users in a RAM user group.
//
// @param request - ListUsersForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithOptions(request *ListUsersForGroupRequest, runtime *dara.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsersForGroup"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) users in a RAM user group.
//
// @param request - ListUsersForGroupRequest
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroup(request *ListUsersForGroupRequest) (_result *ListUsersForGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.ListUsersForGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information about all Resource Access Management (RAM) users in the recycle bin.
//
// @param request - ListUsersInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersInRecycleBinResponse
func (client *Client) ListUsersInRecycleBinWithOptions(request *ListUsersInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *ListUsersInRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsersInRecycleBin"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about all Resource Access Management (RAM) users in the recycle bin.
//
// @param request - ListUsersInRecycleBinRequest
//
// @return ListUsersInRecycleBinResponse
func (client *Client) ListUsersInRecycleBin(request *ListUsersInRecycleBinRequest) (_result *ListUsersInRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersInRecycleBinResponse{}
	_body, _err := client.ListUsersInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries multi-factor authentication (MFA) devices.
//
// @param request - ListVirtualMFADevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirtualMFADevicesResponse
func (client *Client) ListVirtualMFADevicesWithOptions(request *ListVirtualMFADevicesRequest, runtime *dara.RuntimeOptions) (_result *ListVirtualMFADevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirtualMFADevices"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirtualMFADevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries multi-factor authentication (MFA) devices.
//
// @param request - ListVirtualMFADevicesRequest
//
// @return ListVirtualMFADevicesResponse
func (client *Client) ListVirtualMFADevices(request *ListVirtualMFADevicesRequest) (_result *ListVirtualMFADevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirtualMFADevicesResponse{}
	_body, _err := client.ListVirtualMFADevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an application.
//
// @param request - ProvisionApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProvisionApplicationResponse
func (client *Client) ProvisionApplicationWithOptions(request *ProvisionApplicationRequest, runtime *dara.RuntimeOptions) (_result *ProvisionApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Scopes) {
		query["Scopes"] = request.Scopes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProvisionApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProvisionApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an application.
//
// @param request - ProvisionApplicationRequest
//
// @return ProvisionApplicationResponse
func (client *Client) ProvisionApplication(request *ProvisionApplicationRequest) (_result *ProvisionApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ProvisionApplicationResponse{}
	_body, _err := client.ProvisionApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an external application.
//
// @param request - ProvisionExternalApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProvisionExternalApplicationResponse
func (client *Client) ProvisionExternalApplicationWithOptions(request *ProvisionExternalApplicationRequest, runtime *dara.RuntimeOptions) (_result *ProvisionExternalApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Scopes) {
		query["Scopes"] = request.Scopes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProvisionExternalApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProvisionExternalApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an external application.
//
// @param request - ProvisionExternalApplicationRequest
//
// @return ProvisionExternalApplicationResponse
func (client *Client) ProvisionExternalApplication(request *ProvisionExternalApplicationRequest) (_result *ProvisionExternalApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ProvisionExternalApplicationResponse{}
	_body, _err := client.ProvisionExternalApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a client ID from an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the client ID `498469743454717****` from the OIDC IdP named `TestOIDCProvider`.
//
// @param request - RemoveClientIdFromOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClientIdFromOIDCProviderResponse
func (client *Client) RemoveClientIdFromOIDCProviderWithOptions(request *RemoveClientIdFromOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *RemoveClientIdFromOIDCProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.OIDCProviderName) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveClientIdFromOIDCProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveClientIdFromOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a client ID from an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the client ID `498469743454717****` from the OIDC IdP named `TestOIDCProvider`.
//
// @param request - RemoveClientIdFromOIDCProviderRequest
//
// @return RemoveClientIdFromOIDCProviderResponse
func (client *Client) RemoveClientIdFromOIDCProvider(request *RemoveClientIdFromOIDCProviderRequest) (_result *RemoveClientIdFromOIDCProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveClientIdFromOIDCProviderResponse{}
	_body, _err := client.RemoveClientIdFromOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a fingerprint from an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the fingerprint `6938fd4d98bab03faadb97b34396831e3780****` from the OIDC IdP named `TestOIDCProvider`.
//
// @param request - RemoveFingerprintFromOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFingerprintFromOIDCProviderResponse
func (client *Client) RemoveFingerprintFromOIDCProviderWithOptions(request *RemoveFingerprintFromOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *RemoveFingerprintFromOIDCProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Fingerprint) {
		query["Fingerprint"] = request.Fingerprint
	}

	if !dara.IsNil(request.OIDCProviderName) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveFingerprintFromOIDCProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveFingerprintFromOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a fingerprint from an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the fingerprint `6938fd4d98bab03faadb97b34396831e3780****` from the OIDC IdP named `TestOIDCProvider`.
//
// @param request - RemoveFingerprintFromOIDCProviderRequest
//
// @return RemoveFingerprintFromOIDCProviderResponse
func (client *Client) RemoveFingerprintFromOIDCProvider(request *RemoveFingerprintFromOIDCProviderRequest) (_result *RemoveFingerprintFromOIDCProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveFingerprintFromOIDCProviderResponse{}
	_body, _err := client.RemoveFingerprintFromOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a Resource Access Management (RAM) user from a RAM user group.
//
// @param request - RemoveUserFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroupWithOptions(request *RemoveUserFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserFromGroup"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a Resource Access Management (RAM) user from a RAM user group.
//
// @param request - RemoveUserFromGroupRequest
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (_result *RemoveUserFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.RemoveUserFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a specific AccessKey pair that belongs to a Resource Access Management (RAM) user from the recycle bin.
//
// @param request - RestoreAccessKeyFromRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreAccessKeyFromRecycleBinResponse
func (client *Client) RestoreAccessKeyFromRecycleBinWithOptions(request *RestoreAccessKeyFromRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *RestoreAccessKeyFromRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreAccessKeyFromRecycleBin"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreAccessKeyFromRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a specific AccessKey pair that belongs to a Resource Access Management (RAM) user from the recycle bin.
//
// @param request - RestoreAccessKeyFromRecycleBinRequest
//
// @return RestoreAccessKeyFromRecycleBinResponse
func (client *Client) RestoreAccessKeyFromRecycleBin(request *RestoreAccessKeyFromRecycleBinRequest) (_result *RestoreAccessKeyFromRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreAccessKeyFromRecycleBinResponse{}
	_body, _err := client.RestoreAccessKeyFromRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a specific Resource Access Management (RAM) user from the recycle bin.
//
// @param request - RestoreUserFromRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreUserFromRecycleBinResponse
func (client *Client) RestoreUserFromRecycleBinWithOptions(request *RestoreUserFromRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *RestoreUserFromRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreUserFromRecycleBin"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreUserFromRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a specific Resource Access Management (RAM) user from the recycle bin.
//
// @param request - RestoreUserFromRecycleBinRequest
//
// @return RestoreUserFromRecycleBinResponse
func (client *Client) RestoreUserFromRecycleBin(request *RestoreUserFromRecycleBinRequest) (_result *RestoreUserFromRecycleBinResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreUserFromRecycleBinResponse{}
	_body, _err := client.RestoreUserFromRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the default domain name for an Alibaba Cloud account.
//
// @param request - SetDefaultDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomainWithOptions(request *SetDefaultDomainRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultDomainName) {
		query["DefaultDomainName"] = request.DefaultDomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultDomain"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the default domain name for an Alibaba Cloud account.
//
// @param request - SetDefaultDomainRequest
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomain(request *SetDefaultDomainRequest) (_result *SetDefaultDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDefaultDomainResponse{}
	_body, _err := client.SetDefaultDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the password policy for Resource Access Management (RAM) users.
//
// @param request - SetPasswordPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicyWithOptions(request *SetPasswordPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HardExpire) {
		query["HardExpire"] = request.HardExpire
	}

	if !dara.IsNil(request.InitialPasswordAge) {
		query["InitialPasswordAge"] = request.InitialPasswordAge
	}

	if !dara.IsNil(request.MaxLoginAttemps) {
		query["MaxLoginAttemps"] = request.MaxLoginAttemps
	}

	if !dara.IsNil(request.MaxPasswordAge) {
		query["MaxPasswordAge"] = request.MaxPasswordAge
	}

	if !dara.IsNil(request.MinimumPasswordDifferentCharacter) {
		query["MinimumPasswordDifferentCharacter"] = request.MinimumPasswordDifferentCharacter
	}

	if !dara.IsNil(request.MinimumPasswordLength) {
		query["MinimumPasswordLength"] = request.MinimumPasswordLength
	}

	if !dara.IsNil(request.PasswordNotContainUserName) {
		query["PasswordNotContainUserName"] = request.PasswordNotContainUserName
	}

	if !dara.IsNil(request.PasswordReusePrevention) {
		query["PasswordReusePrevention"] = request.PasswordReusePrevention
	}

	if !dara.IsNil(request.RequireLowercaseCharacters) {
		query["RequireLowercaseCharacters"] = request.RequireLowercaseCharacters
	}

	if !dara.IsNil(request.RequireNumbers) {
		query["RequireNumbers"] = request.RequireNumbers
	}

	if !dara.IsNil(request.RequireSymbols) {
		query["RequireSymbols"] = request.RequireSymbols
	}

	if !dara.IsNil(request.RequireUppercaseCharacters) {
		query["RequireUppercaseCharacters"] = request.RequireUppercaseCharacters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPasswordPolicy"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the password policy for Resource Access Management (RAM) users.
//
// @param request - SetPasswordPolicyRequest
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicy(request *SetPasswordPolicyRequest) (_result *SetPasswordPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.SetPasswordPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures security preferences for a RAM user.
//
// Description:
//
// ###
//
// This topic provides an example on how to enable multi-factor authentication (MFA) only for RAM users who initiated unusual logons.
//
// @param tmpReq - SetSecurityPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSecurityPreferenceResponse
func (client *Client) SetSecurityPreferenceWithOptions(tmpReq *SetSecurityPreferenceRequest, runtime *dara.RuntimeOptions) (_result *SetSecurityPreferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetSecurityPreferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VerificationTypes) {
		request.VerificationTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VerificationTypes, dara.String("VerificationTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowUserToChangePassword) {
		query["AllowUserToChangePassword"] = request.AllowUserToChangePassword
	}

	if !dara.IsNil(request.AllowUserToLoginWithPasskey) {
		query["AllowUserToLoginWithPasskey"] = request.AllowUserToLoginWithPasskey
	}

	if !dara.IsNil(request.AllowUserToManageAccessKeys) {
		query["AllowUserToManageAccessKeys"] = request.AllowUserToManageAccessKeys
	}

	if !dara.IsNil(request.AllowUserToManageMFADevices) {
		query["AllowUserToManageMFADevices"] = request.AllowUserToManageMFADevices
	}

	if !dara.IsNil(request.AllowUserToManagePersonalDingTalk) {
		query["AllowUserToManagePersonalDingTalk"] = request.AllowUserToManagePersonalDingTalk
	}

	if !dara.IsNil(request.EnableSaveMFATicket) {
		query["EnableSaveMFATicket"] = request.EnableSaveMFATicket
	}

	if !dara.IsNil(request.LoginNetworkMasks) {
		query["LoginNetworkMasks"] = request.LoginNetworkMasks
	}

	if !dara.IsNil(request.LoginSessionDuration) {
		query["LoginSessionDuration"] = request.LoginSessionDuration
	}

	if !dara.IsNil(request.MFAOperationForLogin) {
		query["MFAOperationForLogin"] = request.MFAOperationForLogin
	}

	if !dara.IsNil(request.MaxIdleDaysForAccessKeys) {
		query["MaxIdleDaysForAccessKeys"] = request.MaxIdleDaysForAccessKeys
	}

	if !dara.IsNil(request.MaxIdleDaysForUsers) {
		query["MaxIdleDaysForUsers"] = request.MaxIdleDaysForUsers
	}

	if !dara.IsNil(request.OperationForRiskLogin) {
		query["OperationForRiskLogin"] = request.OperationForRiskLogin
	}

	if !dara.IsNil(request.VerificationTypesShrink) {
		query["VerificationTypes"] = request.VerificationTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSecurityPreference"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSecurityPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures security preferences for a RAM user.
//
// Description:
//
// ###
//
// This topic provides an example on how to enable multi-factor authentication (MFA) only for RAM users who initiated unusual logons.
//
// @param request - SetSecurityPreferenceRequest
//
// @return SetSecurityPreferenceResponse
func (client *Client) SetSecurityPreference(request *SetSecurityPreferenceRequest) (_result *SetSecurityPreferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSecurityPreferenceResponse{}
	_body, _err := client.SetSecurityPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures information about user-based single sign-on (SSO).
//
// @param request - SetUserSsoSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserSsoSettingsResponse
func (client *Client) SetUserSsoSettingsWithOptions(request *SetUserSsoSettingsRequest, runtime *dara.RuntimeOptions) (_result *SetUserSsoSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthnSignAlgo) {
		query["AuthnSignAlgo"] = request.AuthnSignAlgo
	}

	if !dara.IsNil(request.AuxiliaryDomain) {
		query["AuxiliaryDomain"] = request.AuxiliaryDomain
	}

	if !dara.IsNil(request.MetadataDocument) {
		query["MetadataDocument"] = request.MetadataDocument
	}

	if !dara.IsNil(request.SsoEnabled) {
		query["SsoEnabled"] = request.SsoEnabled
	}

	if !dara.IsNil(request.SsoLoginWithDomain) {
		query["SsoLoginWithDomain"] = request.SsoLoginWithDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetUserSsoSettings"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetUserSsoSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures information about user-based single sign-on (SSO).
//
// @param request - SetUserSsoSettingsRequest
//
// @return SetUserSsoSettingsResponse
func (client *Client) SetUserSsoSettings(request *SetUserSsoSettingsRequest) (_result *SetUserSsoSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetUserSsoSettingsResponse{}
	_body, _err := client.SetUserSsoSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a mobile phone or email to a Resource Access Management (RAM) user.
//
// @param request - SetVerificationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVerificationInfoResponse
func (client *Client) SetVerificationInfoWithOptions(request *SetVerificationInfoRequest, runtime *dara.RuntimeOptions) (_result *SetVerificationInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.MobilePhone) {
		query["MobilePhone"] = request.MobilePhone
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetVerificationInfo"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetVerificationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a mobile phone or email to a Resource Access Management (RAM) user.
//
// @param request - SetVerificationInfoRequest
//
// @return SetVerificationInfoResponse
func (client *Client) SetVerificationInfo(request *SetVerificationInfoRequest) (_result *SetVerificationInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetVerificationInfoResponse{}
	_body, _err := client.SetVerificationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourcePrincipalName) {
		query["ResourcePrincipalName"] = request.ResourcePrincipalName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - UnbindMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindMFADeviceResponse
func (client *Client) UnbindMFADeviceWithOptions(request *UnbindMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *UnbindMFADeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindMFADevice"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - UnbindMFADeviceRequest
//
// @return UnbindMFADeviceResponse
func (client *Client) UnbindMFADevice(request *UnbindMFADeviceRequest) (_result *UnbindMFADeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindMFADeviceResponse{}
	_body, _err := client.UnbindMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a mobile phone or email from a Resource Access Management (RAM) user.
//
// @param request - UnbindVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindVerificationResponse
func (client *Client) UnbindVerificationWithOptions(request *UnbindVerificationRequest, runtime *dara.RuntimeOptions) (_result *UnbindVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.MobilePhone) {
		query["MobilePhone"] = request.MobilePhone
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindVerification"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a mobile phone or email from a Resource Access Management (RAM) user.
//
// @param request - UnbindVerificationRequest
//
// @return UnbindVerificationResponse
func (client *Client) UnbindVerification(request *UnbindVerificationRequest) (_result *UnbindVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindVerificationResponse{}
	_body, _err := client.UnbindVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from a resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourcePrincipalName) {
		query["ResourcePrincipalName"] = request.ResourcePrincipalName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from a resource.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - UpdateAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccessKeyResponse
func (client *Client) UpdateAccessKeyWithOptions(request *UpdateAccessKeyRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccessKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserAccessKeyId) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccessKey"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - UpdateAccessKeyRequest
//
// @return UpdateAccessKeyResponse
func (client *Client) UpdateAccessKey(request *UpdateAccessKeyRequest) (_result *UpdateAccessKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAccessKeyResponse{}
	_body, _err := client.UpdateAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a specified application.
//
// @param request - UpdateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplicationWithOptions(request *UpdateApplicationRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.NewAccessTokenValidity) {
		query["NewAccessTokenValidity"] = request.NewAccessTokenValidity
	}

	if !dara.IsNil(request.NewDisplayName) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !dara.IsNil(request.NewIsMultiTenant) {
		query["NewIsMultiTenant"] = request.NewIsMultiTenant
	}

	if !dara.IsNil(request.NewPredefinedScopes) {
		query["NewPredefinedScopes"] = request.NewPredefinedScopes
	}

	if !dara.IsNil(request.NewRedirectUris) {
		query["NewRedirectUris"] = request.NewRedirectUris
	}

	if !dara.IsNil(request.NewRefreshTokenValidity) {
		query["NewRefreshTokenValidity"] = request.NewRefreshTokenValidity
	}

	if !dara.IsNil(request.NewRequiredScopes) {
		query["NewRequiredScopes"] = request.NewRequiredScopes
	}

	if !dara.IsNil(request.NewSecretRequired) {
		query["NewSecretRequired"] = request.NewSecretRequired
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplication"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a specified application.
//
// @param request - UpdateApplicationRequest
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplication(request *UpdateApplicationRequest) (_result *UpdateApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.UpdateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about a Resource Access Management (RAM) user group.
//
// @param request - UpdateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.NewComments) {
		query["NewComments"] = request.NewComments
	}

	if !dara.IsNil(request.NewDisplayName) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !dara.IsNil(request.NewGroupName) {
		query["NewGroupName"] = request.NewGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroup"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about a Resource Access Management (RAM) user group.
//
// @param request - UpdateGroupRequest
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroup(request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the console logon configurations of a Resource Access Management (RAM) user.
//
// @param request - UpdateLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoginProfileResponse
func (client *Client) UpdateLoginProfileWithOptions(request *UpdateLoginProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateLoginProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MFABindRequired) {
		query["MFABindRequired"] = request.MFABindRequired
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordResetRequired) {
		query["PasswordResetRequired"] = request.PasswordResetRequired
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLoginProfile"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the console logon configurations of a Resource Access Management (RAM) user.
//
// @param request - UpdateLoginProfileRequest
//
// @return UpdateLoginProfileResponse
func (client *Client) UpdateLoginProfile(request *UpdateLoginProfileRequest) (_result *UpdateLoginProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLoginProfileResponse{}
	_body, _err := client.UpdateLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description and client IDs of an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to change the description of the OIDC IdP named `TestOIDCProvider` to `This is a new OIDC Provider.`
//
// @param request - UpdateOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOIDCProviderResponse
func (client *Client) UpdateOIDCProviderWithOptions(request *UpdateOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateOIDCProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIds) {
		query["ClientIds"] = request.ClientIds
	}

	if !dara.IsNil(request.IssuanceLimitTime) {
		query["IssuanceLimitTime"] = request.IssuanceLimitTime
	}

	if !dara.IsNil(request.NewDescription) {
		query["NewDescription"] = request.NewDescription
	}

	if !dara.IsNil(request.OIDCProviderName) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOIDCProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description and client IDs of an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to change the description of the OIDC IdP named `TestOIDCProvider` to `This is a new OIDC Provider.`
//
// @param request - UpdateOIDCProviderRequest
//
// @return UpdateOIDCProviderResponse
func (client *Client) UpdateOIDCProvider(request *UpdateOIDCProviderRequest) (_result *UpdateOIDCProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOIDCProviderResponse{}
	_body, _err := client.UpdateOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name of a passkey.
//
// @param request - UpdatePasskeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePasskeyResponse
func (client *Client) UpdatePasskeyWithOptions(request *UpdatePasskeyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePasskeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PasskeyId) {
		query["PasskeyId"] = request.PasskeyId
	}

	if !dara.IsNil(request.PasskeyName) {
		query["PasskeyName"] = request.PasskeyName
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePasskey"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePasskeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name of a passkey.
//
// @param request - UpdatePasskeyRequest
//
// @return UpdatePasskeyResponse
func (client *Client) UpdatePasskey(request *UpdatePasskeyRequest) (_result *UpdatePasskeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePasskeyResponse{}
	_body, _err := client.UpdatePasskeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about an identity provider (IdP) for role-based single sign-on (SSO).
//
// Description:
//
// This topic provides an example on how to change the description of an IdP named `test-provider` to `This is a new provider.`
//
// @param request - UpdateSAMLProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSAMLProviderResponse
func (client *Client) UpdateSAMLProviderWithOptions(request *UpdateSAMLProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateSAMLProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthnSignAlgo) {
		query["AuthnSignAlgo"] = request.AuthnSignAlgo
	}

	if !dara.IsNil(request.NewDescription) {
		query["NewDescription"] = request.NewDescription
	}

	if !dara.IsNil(request.NewEncodedSAMLMetadataDocument) {
		query["NewEncodedSAMLMetadataDocument"] = request.NewEncodedSAMLMetadataDocument
	}

	if !dara.IsNil(request.SAMLProviderName) {
		query["SAMLProviderName"] = request.SAMLProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSAMLProvider"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSAMLProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about an identity provider (IdP) for role-based single sign-on (SSO).
//
// Description:
//
// This topic provides an example on how to change the description of an IdP named `test-provider` to `This is a new provider.`
//
// @param request - UpdateSAMLProviderRequest
//
// @return UpdateSAMLProviderResponse
func (client *Client) UpdateSAMLProvider(request *UpdateSAMLProviderRequest) (_result *UpdateSAMLProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSAMLProviderResponse{}
	_body, _err := client.UpdateSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a RAM user.
//
// Description:
//
// This topic provides an example to show how to modify the name of a RAM user from `test@example.onaliyun.com` to `new@example.onaliyun.com`.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewComments) {
		query["NewComments"] = request.NewComments
	}

	if !dara.IsNil(request.NewDisplayName) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !dara.IsNil(request.NewEmail) {
		query["NewEmail"] = request.NewEmail
	}

	if !dara.IsNil(request.NewMobilePhone) {
		query["NewMobilePhone"] = request.NewMobilePhone
	}

	if !dara.IsNil(request.NewUserPrincipalName) {
		query["NewUserPrincipalName"] = request.NewUserPrincipalName
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserPrincipalName) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2019-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a RAM user.
//
// Description:
//
// This topic provides an example to show how to modify the name of a RAM user from `test@example.onaliyun.com` to `new@example.onaliyun.com`.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
