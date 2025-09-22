// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds a client ID to an OpenID Connect (OIDC) identity provider (IdP).
//
// @param request - AddClientIdToOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddClientIdToOIDCProviderResponse
func (client *Client) AddClientIdToOIDCProviderWithContext(ctx context.Context, request *AddClientIdToOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *AddClientIdToOIDCProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFingerprintToOIDCProviderResponse
func (client *Client) AddFingerprintToOIDCProviderWithContext(ctx context.Context, request *AddFingerprintToOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *AddFingerprintToOIDCProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroupWithContext(ctx context.Context, request *AddUserToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUserToGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindMFADeviceResponse
func (client *Client) BindMFADeviceWithContext(ctx context.Context, request *BindMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *BindMFADeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangePasswordResponse
func (client *Client) ChangePasswordWithContext(ctx context.Context, request *ChangePasswordRequest, runtime *dara.RuntimeOptions) (_result *ChangePasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessKeyResponse
func (client *Client) CreateAccessKeyWithContext(ctx context.Context, request *CreateAccessKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppSecretResponse
func (client *Client) CreateAppSecretWithContext(ctx context.Context, request *CreateAppSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateAppSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoginProfileResponse
func (client *Client) CreateLoginProfileWithContext(ctx context.Context, request *CreateLoginProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateLoginProfileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOIDCProviderResponse
func (client *Client) CreateOIDCProviderWithContext(ctx context.Context, request *CreateOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateOIDCProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSAMLProviderResponse
func (client *Client) CreateSAMLProviderWithContext(ctx context.Context, request *CreateSAMLProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateSAMLProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualMFADeviceResponse
func (client *Client) CreateVirtualMFADeviceWithContext(ctx context.Context, request *CreateVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *CreateVirtualMFADeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessKeyResponse
func (client *Client) DeleteAccessKeyWithContext(ctx context.Context, request *DeleteAccessKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessKeyInRecycleBinResponse
func (client *Client) DeleteAccessKeyInRecycleBinWithContext(ctx context.Context, request *DeleteAccessKeyInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessKeyInRecycleBinResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppSecretResponse
func (client *Client) DeleteAppSecretWithContext(ctx context.Context, request *DeleteAppSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoginProfileResponse
func (client *Client) DeleteLoginProfileWithContext(ctx context.Context, request *DeleteLoginProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteLoginProfileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOIDCProviderResponse
func (client *Client) DeleteOIDCProviderWithContext(ctx context.Context, request *DeleteOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteOIDCProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePasskeyResponse
func (client *Client) DeletePasskeyWithContext(ctx context.Context, request *DeletePasskeyRequest, runtime *dara.RuntimeOptions) (_result *DeletePasskeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSAMLProviderResponse
func (client *Client) DeleteSAMLProviderWithContext(ctx context.Context, request *DeleteSAMLProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteSAMLProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserInRecycleBinResponse
func (client *Client) DeleteUserInRecycleBinWithContext(ctx context.Context, request *DeleteUserInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserInRecycleBinResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADeviceWithContext(ctx context.Context, request *DeleteVirtualMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprovisionApplicationResponse
func (client *Client) DeprovisionApplicationWithContext(ctx context.Context, request *DeprovisionApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeprovisionApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprovisionExternalApplicationResponse
func (client *Client) DeprovisionExternalApplicationWithContext(ctx context.Context, request *DeprovisionExternalApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeprovisionExternalApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVirtualMFAResponse
func (client *Client) DisableVirtualMFAWithContext(ctx context.Context, request *DisableVirtualMFARequest, runtime *dara.RuntimeOptions) (_result *DisableVirtualMFAResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyInfoInRecycleBinResponse
func (client *Client) GetAccessKeyInfoInRecycleBinWithContext(ctx context.Context, request *GetAccessKeyInfoInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyInfoInRecycleBinResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedResponse
func (client *Client) GetAccessKeyLastUsedWithContext(ctx context.Context, request *GetAccessKeyLastUsedRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSecretResponse
func (client *Client) GetAppSecretWithContext(ctx context.Context, request *GetAppSecretRequest, runtime *dara.RuntimeOptions) (_result *GetAppSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithContext(ctx context.Context, request *GetApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisionInfoResponse
func (client *Client) GetApplicationProvisionInfoWithContext(ctx context.Context, request *GetApplicationProvisionInfoRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationProvisionInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialReportResponse
func (client *Client) GetCredentialReportWithContext(ctx context.Context, request *GetCredentialReportRequest, runtime *dara.RuntimeOptions) (_result *GetCredentialReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExternalApplicationResponse
func (client *Client) GetExternalApplicationWithContext(ctx context.Context, request *GetExternalApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetExternalApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGovernanceItemReportResponse
func (client *Client) GetGovernanceItemReportWithContext(ctx context.Context, request *GetGovernanceItemReportRequest, runtime *dara.RuntimeOptions) (_result *GetGovernanceItemReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithContext(ctx context.Context, request *GetGroupRequest, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginProfileResponse
func (client *Client) GetLoginProfileWithContext(ctx context.Context, request *GetLoginProfileRequest, runtime *dara.RuntimeOptions) (_result *GetLoginProfileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOIDCProviderResponse
func (client *Client) GetOIDCProviderWithContext(ctx context.Context, request *GetOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *GetOIDCProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSAMLProviderResponse
func (client *Client) GetSAMLProviderWithContext(ctx context.Context, request *GetSAMLProviderRequest, runtime *dara.RuntimeOptions) (_result *GetSAMLProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInRecycleBinResponse
func (client *Client) GetUserInRecycleBinWithContext(ctx context.Context, request *GetUserInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *GetUserInRecycleBinResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserMFAInfoResponse
func (client *Client) GetUserMFAInfoWithContext(ctx context.Context, request *GetUserMFAInfoRequest, runtime *dara.RuntimeOptions) (_result *GetUserMFAInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVerificationInfoResponse
func (client *Client) GetVerificationInfoWithContext(ctx context.Context, request *GetVerificationInfoRequest, runtime *dara.RuntimeOptions) (_result *GetVerificationInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessKeysResponse
func (client *Client) ListAccessKeysWithContext(ctx context.Context, request *ListAccessKeysRequest, runtime *dara.RuntimeOptions) (_result *ListAccessKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessKeysInRecycleBinResponse
func (client *Client) ListAccessKeysInRecycleBinWithContext(ctx context.Context, request *ListAccessKeysInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *ListAccessKeysInRecycleBinResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppSecretIdsResponse
func (client *Client) ListAppSecretIdsWithContext(ctx context.Context, request *ListAppSecretIdsRequest, runtime *dara.RuntimeOptions) (_result *ListAppSecretIdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationProvisionInfosResponse
func (client *Client) ListApplicationProvisionInfosWithContext(ctx context.Context, request *ListApplicationProvisionInfosRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationProvisionInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithContext(ctx context.Context, request *ListGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithContext(ctx context.Context, request *ListGroupsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOIDCProvidersResponse
func (client *Client) ListOIDCProvidersWithContext(ctx context.Context, request *ListOIDCProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListOIDCProvidersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPasskeysResponse
func (client *Client) ListPasskeysWithContext(ctx context.Context, request *ListPasskeysRequest, runtime *dara.RuntimeOptions) (_result *ListPasskeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPredefinedScopesResponse
func (client *Client) ListPredefinedScopesWithContext(ctx context.Context, request *ListPredefinedScopesRequest, runtime *dara.RuntimeOptions) (_result *ListPredefinedScopesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSAMLProvidersResponse
func (client *Client) ListSAMLProvidersWithContext(ctx context.Context, request *ListSAMLProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListSAMLProvidersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserBasicInfosResponse
func (client *Client) ListUserBasicInfosWithContext(ctx context.Context, request *ListUserBasicInfosRequest, runtime *dara.RuntimeOptions) (_result *ListUserBasicInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithContext(ctx context.Context, request *ListUsersForGroupRequest, runtime *dara.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersInRecycleBinResponse
func (client *Client) ListUsersInRecycleBinWithContext(ctx context.Context, request *ListUsersInRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *ListUsersInRecycleBinResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirtualMFADevicesResponse
func (client *Client) ListVirtualMFADevicesWithContext(ctx context.Context, request *ListVirtualMFADevicesRequest, runtime *dara.RuntimeOptions) (_result *ListVirtualMFADevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProvisionApplicationResponse
func (client *Client) ProvisionApplicationWithContext(ctx context.Context, request *ProvisionApplicationRequest, runtime *dara.RuntimeOptions) (_result *ProvisionApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProvisionExternalApplicationResponse
func (client *Client) ProvisionExternalApplicationWithContext(ctx context.Context, request *ProvisionExternalApplicationRequest, runtime *dara.RuntimeOptions) (_result *ProvisionExternalApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClientIdFromOIDCProviderResponse
func (client *Client) RemoveClientIdFromOIDCProviderWithContext(ctx context.Context, request *RemoveClientIdFromOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *RemoveClientIdFromOIDCProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFingerprintFromOIDCProviderResponse
func (client *Client) RemoveFingerprintFromOIDCProviderWithContext(ctx context.Context, request *RemoveFingerprintFromOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *RemoveFingerprintFromOIDCProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroupWithContext(ctx context.Context, request *RemoveUserFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreAccessKeyFromRecycleBinResponse
func (client *Client) RestoreAccessKeyFromRecycleBinWithContext(ctx context.Context, request *RestoreAccessKeyFromRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *RestoreAccessKeyFromRecycleBinResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreUserFromRecycleBinResponse
func (client *Client) RestoreUserFromRecycleBinWithContext(ctx context.Context, request *RestoreUserFromRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *RestoreUserFromRecycleBinResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomainWithContext(ctx context.Context, request *SetDefaultDomainRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicyWithContext(ctx context.Context, request *SetPasswordPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HardExpire) {
		query["HardExpire"] = request.HardExpire
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - SetSecurityPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSecurityPreferenceResponse
func (client *Client) SetSecurityPreferenceWithContext(ctx context.Context, tmpReq *SetSecurityPreferenceRequest, runtime *dara.RuntimeOptions) (_result *SetSecurityPreferenceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserSsoSettingsResponse
func (client *Client) SetUserSsoSettingsWithContext(ctx context.Context, request *SetUserSsoSettingsRequest, runtime *dara.RuntimeOptions) (_result *SetUserSsoSettingsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVerificationInfoResponse
func (client *Client) SetVerificationInfoWithContext(ctx context.Context, request *SetVerificationInfoRequest, runtime *dara.RuntimeOptions) (_result *SetVerificationInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindMFADeviceResponse
func (client *Client) UnbindMFADeviceWithContext(ctx context.Context, request *UnbindMFADeviceRequest, runtime *dara.RuntimeOptions) (_result *UnbindMFADeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindVerificationResponse
func (client *Client) UnbindVerificationWithContext(ctx context.Context, request *UnbindVerificationRequest, runtime *dara.RuntimeOptions) (_result *UnbindVerificationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccessKeyResponse
func (client *Client) UpdateAccessKeyWithContext(ctx context.Context, request *UpdateAccessKeyRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccessKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplicationWithContext(ctx context.Context, request *UpdateApplicationRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithContext(ctx context.Context, request *UpdateGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoginProfileResponse
func (client *Client) UpdateLoginProfileWithContext(ctx context.Context, request *UpdateLoginProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateLoginProfileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOIDCProviderResponse
func (client *Client) UpdateOIDCProviderWithContext(ctx context.Context, request *UpdateOIDCProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateOIDCProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePasskeyResponse
func (client *Client) UpdatePasskeyWithContext(ctx context.Context, request *UpdatePasskeyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePasskeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSAMLProviderResponse
func (client *Client) UpdateSAMLProviderWithContext(ctx context.Context, request *UpdateSAMLProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateSAMLProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
