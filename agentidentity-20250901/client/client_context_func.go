// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建应用
//
// @param request - AddSAMLIdentityProviderCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSAMLIdentityProviderCertificateResponse
func (client *Client) AddSAMLIdentityProviderCertificateWithContext(ctx context.Context, request *AddSAMLIdentityProviderCertificateRequest, runtime *dara.RuntimeOptions) (_result *AddSAMLIdentityProviderCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	if !dara.IsNil(request.X509Certificate) {
		body["X509Certificate"] = request.X509Certificate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSAMLIdentityProviderCertificate"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSAMLIdentityProviderCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 策略集关联网关
//
// @param request - AttachPolicySetToGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicySetToGatewayResponse
func (client *Client) AttachPolicySetToGatewayWithContext(ctx context.Context, request *AttachPolicySetToGatewayRequest, runtime *dara.RuntimeOptions) (_result *AttachPolicySetToGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnforcementMode) {
		body["EnforcementMode"] = request.EnforcementMode
	}

	if !dara.IsNil(request.GatewayArn) {
		body["GatewayArn"] = request.GatewayArn
	}

	if !dara.IsNil(request.GatewayType) {
		body["GatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachPolicySetToGateway"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachPolicySetToGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个 API 密钥凭证提供商
//
// @param request - CreateAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAPIKeyCredentialProviderResponse
func (client *Client) CreateAPIKeyCredentialProviderWithContext(ctx context.Context, request *CreateAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateAPIKeyCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		body["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.APIKeyCredentialProviderName) {
		body["APIKeyCredentialProviderName"] = request.APIKeyCredentialProviderName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAPIKeyCredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAPIKeyCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建WorkloadIdentity
//
// @param request - CreateClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientSecretResponse
func (client *Client) CreateClientSecretWithContext(ctx context.Context, request *CreateClientSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateClientSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientName) {
		body["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientSecret"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建IdentityProvider
//
// @param tmpReq - CreateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityProviderResponse
func (client *Client) CreateIdentityProviderWithContext(ctx context.Context, tmpReq *CreateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateIdentityProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedAudience) {
		request.AllowedAudienceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedAudience, dara.String("AllowedAudience"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedAudienceShrink) {
		body["AllowedAudience"] = request.AllowedAudienceShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DiscoveryURL) {
		body["DiscoveryURL"] = request.DiscoveryURL
	}

	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 OAuth2 凭证提供商
//
// @param tmpReq - CreateOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOAuth2CredentialProviderResponse
func (client *Client) CreateOAuth2CredentialProviderWithContext(ctx context.Context, tmpReq *CreateOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateOAuth2CredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOAuth2CredentialProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OAuth2ProviderConfig) {
		request.OAuth2ProviderConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OAuth2ProviderConfig, dara.String("OAuth2ProviderConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackURL) {
		body["CallbackURL"] = request.CallbackURL
	}

	if !dara.IsNil(request.CredentialProviderVendor) {
		body["CredentialProviderVendor"] = request.CredentialProviderVendor
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.OAuth2CredentialProviderName) {
		body["OAuth2CredentialProviderName"] = request.OAuth2CredentialProviderName
	}

	if !dara.IsNil(request.OAuth2ProviderConfigShrink) {
		body["OAuth2ProviderConfig"] = request.OAuth2ProviderConfigShrink
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOAuth2CredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOAuth2CredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个权限策略
//
// @param tmpReq - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithContext(ctx context.Context, tmpReq *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Definition) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Definition, dara.String("Definition"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DefinitionShrink) {
		body["Definition"] = request.DefinitionShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个权限策略集合
//
// @param request - CreatePolicySetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicySetResponse
func (client *Client) CreatePolicySetWithContext(ctx context.Context, request *CreatePolicySetRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicySetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicySet"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicySetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Role
//
// @param request - CreateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.RoleName) {
		body["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRole"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 策略集关联网关
//
// @param request - CreateRoleAssignmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleAssignmentResponse
func (client *Client) CreateRoleAssignmentWithContext(ctx context.Context, request *CreateRoleAssignmentRequest, runtime *dara.RuntimeOptions) (_result *CreateRoleAssignmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PrincipalName) {
		body["PrincipalName"] = request.PrincipalName
	}

	if !dara.IsNil(request.PrincipalType) {
		body["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.RoleName) {
		body["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoleAssignment"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoleAssignmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个凭证库
//
// @param tmpReq - CreateTokenVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTokenVaultResponse
func (client *Client) CreateTokenVaultWithContext(ctx context.Context, tmpReq *CreateTokenVaultRequest, runtime *dara.RuntimeOptions) (_result *CreateTokenVaultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTokenVaultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EncryptionConfig) {
		request.EncryptionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncryptionConfig, dara.String("EncryptionConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EncryptionConfigShrink) {
		body["EncryptionConfig"] = request.EncryptionConfigShrink
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTokenVault"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTokenVaultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建UserPool
//
// @param request - CreateUserPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserPoolResponse
func (client *Client) CreateUserPoolWithContext(ctx context.Context, request *CreateUserPoolRequest, runtime *dara.RuntimeOptions) (_result *CreateUserPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserPool"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建WorkloadIdentity
//
// @param tmpReq - CreateUserPoolClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserPoolClientResponse
func (client *Client) CreateUserPoolClientWithContext(ctx context.Context, tmpReq *CreateUserPoolClientRequest, runtime *dara.RuntimeOptions) (_result *CreateUserPoolClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUserPoolClientShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RedirectURIs) {
		request.RedirectURIsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RedirectURIs, dara.String("RedirectURIs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessTokenValidity) {
		body["AccessTokenValidity"] = request.AccessTokenValidity
	}

	if !dara.IsNil(request.ClientName) {
		body["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.EnforcePKCE) {
		body["EnforcePKCE"] = request.EnforcePKCE
	}

	if !dara.IsNil(request.RedirectURIsShrink) {
		body["RedirectURIs"] = request.RedirectURIsShrink
	}

	if !dara.IsNil(request.RefreshTokenValidity) {
		body["RefreshTokenValidity"] = request.RefreshTokenValidity
	}

	if !dara.IsNil(request.SecretRequired) {
		body["SecretRequired"] = request.SecretRequired
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserPoolClient"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserPoolClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建WorkloadIdentity
//
// @param tmpReq - CreateWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkloadIdentityResponse
func (client *Client) CreateWorkloadIdentityWithContext(ctx context.Context, tmpReq *CreateWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkloadIdentityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedResourceOAuth2ReturnURLs) {
		request.AllowedResourceOAuth2ReturnURLsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedResourceOAuth2ReturnURLs, dara.String("AllowedResourceOAuth2ReturnURLs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedResourceOAuth2ReturnURLsShrink) {
		body["AllowedResourceOAuth2ReturnURLs"] = request.AllowedResourceOAuth2ReturnURLsShrink
	}

	if !dara.IsNil(request.CreateRAMRole) {
		body["CreateRAMRole"] = request.CreateRAMRole
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.SessionBindingEnabled) {
		body["SessionBindingEnabled"] = request.SessionBindingEnabled
	}

	if !dara.IsNil(request.SourceAgentArn) {
		body["SourceAgentArn"] = request.SourceAgentArn
	}

	if !dara.IsNil(request.SourcePlatform) {
		body["SourcePlatform"] = request.SourcePlatform
	}

	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkloadIdentity"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkloadIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个 API 密钥凭证提供商
//
// @param request - DeleteAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAPIKeyCredentialProviderResponse
func (client *Client) DeleteAPIKeyCredentialProviderWithContext(ctx context.Context, request *DeleteAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteAPIKeyCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.APIKeyCredentialProviderName) {
		body["APIKeyCredentialProviderName"] = request.APIKeyCredentialProviderName
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAPIKeyCredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAPIKeyCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除IdentityProvider
//
// @param request - DeleteClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientSecretResponse
func (client *Client) DeleteClientSecretWithContext(ctx context.Context, request *DeleteClientSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientName) {
		body["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.ClientSecretId) {
		body["ClientSecretId"] = request.ClientSecretId
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClientSecret"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除IdentityProvider
//
// @param request - DeleteIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdentityProviderResponse
func (client *Client) DeleteIdentityProviderWithContext(ctx context.Context, request *DeleteIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 OAuth2 凭证提供商
//
// @param request - DeleteOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOAuth2CredentialProviderResponse
func (client *Client) DeleteOAuth2CredentialProviderWithContext(ctx context.Context, request *DeleteOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteOAuth2CredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OAuth2CredentialProviderName) {
		body["OAuth2CredentialProviderName"] = request.OAuth2CredentialProviderName
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOAuth2CredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOAuth2CredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个权限策略
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithContext(ctx context.Context, request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个权限策略集合
//
// @param request - DeletePolicySetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicySetResponse
func (client *Client) DeletePolicySetWithContext(ctx context.Context, request *DeletePolicySetRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicySetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicySet"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicySetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Role
//
// @param request - DeleteRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithContext(ctx context.Context, request *DeleteRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleName) {
		body["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRole"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 网关取消关联策略集
//
// @param request - DeleteRoleAssignmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleAssignmentResponse
func (client *Client) DeleteRoleAssignmentWithContext(ctx context.Context, request *DeleteRoleAssignmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoleAssignmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PrincipalName) {
		body["PrincipalName"] = request.PrincipalName
	}

	if !dara.IsNil(request.PrincipalType) {
		body["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.RoleName) {
		body["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoleAssignment"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoleAssignmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除WorkloadIdentity
//
// @param request - DeleteSAMLIdentityProviderCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSAMLIdentityProviderCertificateResponse
func (client *Client) DeleteSAMLIdentityProviderCertificateWithContext(ctx context.Context, request *DeleteSAMLIdentityProviderCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSAMLIdentityProviderCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		body["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSAMLIdentityProviderCertificate"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSAMLIdentityProviderCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个指定的凭证库。
//
// @param request - DeleteTokenVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTokenVaultResponse
func (client *Client) DeleteTokenVaultWithContext(ctx context.Context, request *DeleteTokenVaultRequest, runtime *dara.RuntimeOptions) (_result *DeleteTokenVaultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTokenVault"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTokenVaultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除User
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2025-09-01"),
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
// 删除UserPool
//
// @param request - DeleteUserPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPoolResponse
func (client *Client) DeleteUserPoolWithContext(ctx context.Context, request *DeleteUserPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserPool"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除WorkloadIdentity
//
// @param request - DeleteUserPoolClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPoolClientResponse
func (client *Client) DeleteUserPoolClientWithContext(ctx context.Context, request *DeleteUserPoolClientRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserPoolClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientName) {
		body["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserPoolClient"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserPoolClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除WorkloadIdentity
//
// @param request - DeleteWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkloadIdentityResponse
func (client *Client) DeleteWorkloadIdentityWithContext(ctx context.Context, request *DeleteWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkloadIdentityResponse, _err error) {
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
		Action:      dara.String("DeleteWorkloadIdentity"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkloadIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 网关取消关联策略集
//
// @param request - DetachPolicySetFromGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicySetFromGatewayResponse
func (client *Client) DetachPolicySetFromGatewayWithContext(ctx context.Context, request *DetachPolicySetFromGatewayRequest, runtime *dara.RuntimeOptions) (_result *DetachPolicySetFromGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GatewayArn) {
		body["GatewayArn"] = request.GatewayArn
	}

	if !dara.IsNil(request.GatewayType) {
		body["GatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachPolicySetFromGateway"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachPolicySetFromGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一个 API 密钥凭证提供商
//
// @param request - GetAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAPIKeyCredentialProviderResponse
func (client *Client) GetAPIKeyCredentialProviderWithContext(ctx context.Context, request *GetAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *GetAPIKeyCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.APIKeyCredentialProviderName) {
		body["APIKeyCredentialProviderName"] = request.APIKeyCredentialProviderName
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAPIKeyCredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAPIKeyCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网关策略配置
//
// @param request - GetGatewayPolicyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayPolicyConfigResponse
func (client *Client) GetGatewayPolicyConfigWithContext(ctx context.Context, request *GetGatewayPolicyConfigRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayPolicyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GatewayArn) {
		body["GatewayArn"] = request.GatewayArn
	}

	if !dara.IsNil(request.GatewayType) {
		body["GatewayType"] = request.GatewayType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayPolicyConfig"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayPolicyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderResponse
func (client *Client) GetIdentityProviderWithContext(ctx context.Context, request *GetIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 OAuth2 凭证提供商
//
// @param request - GetOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOAuth2CredentialProviderResponse
func (client *Client) GetOAuth2CredentialProviderWithContext(ctx context.Context, request *GetOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *GetOAuth2CredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OAuth2CredentialProviderName) {
		body["OAuth2CredentialProviderName"] = request.OAuth2CredentialProviderName
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOAuth2CredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOAuth2CredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一个 权限策略
//
// @param request - GetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithContext(ctx context.Context, request *GetPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicy"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一个 权限策略集合
//
// @param request - GetPolicySetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicySetResponse
func (client *Client) GetPolicySetWithContext(ctx context.Context, request *GetPolicySetRequest, runtime *dara.RuntimeOptions) (_result *GetPolicySetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicySet"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicySetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Role
//
// @param request - GetRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleResponse
func (client *Client) GetRoleWithContext(ctx context.Context, request *GetRoleRequest, runtime *dara.RuntimeOptions) (_result *GetRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleName) {
		body["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRole"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetSAMLIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSAMLIdentityProviderResponse
func (client *Client) GetSAMLIdentityProviderWithContext(ctx context.Context, request *GetSAMLIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetSAMLIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSAMLIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSAMLIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetSAMLServiceProviderInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSAMLServiceProviderInfoResponse
func (client *Client) GetSAMLServiceProviderInfoWithContext(ctx context.Context, request *GetSAMLServiceProviderInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSAMLServiceProviderInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSAMLServiceProviderInfo"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSAMLServiceProviderInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定凭证库的详细配置。
//
// @param request - GetTokenVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenVaultResponse
func (client *Client) GetTokenVaultWithContext(ctx context.Context, request *GetTokenVaultRequest, runtime *dara.RuntimeOptions) (_result *GetTokenVaultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTokenVault"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenVaultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2025-09-01"),
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
// 获取UserPool
//
// @param request - GetUserPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserPoolResponse
func (client *Client) GetUserPoolWithContext(ctx context.Context, request *GetUserPoolRequest, runtime *dara.RuntimeOptions) (_result *GetUserPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserPool"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetUserPoolClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserPoolClientResponse
func (client *Client) GetUserPoolClientWithContext(ctx context.Context, request *GetUserPoolClientRequest, runtime *dara.RuntimeOptions) (_result *GetUserPoolClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientName) {
		body["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserPoolClient"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserPoolClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkloadIdentityResponse
func (client *Client) GetWorkloadIdentityWithContext(ctx context.Context, request *GetWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadIdentityResponse, _err error) {
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
		Action:      dara.String("GetWorkloadIdentity"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkloadIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出 API 密钥凭证提供商
//
// @param request - ListAPIKeyCredentialProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAPIKeyCredentialProvidersResponse
func (client *Client) ListAPIKeyCredentialProvidersWithContext(ctx context.Context, request *ListAPIKeyCredentialProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListAPIKeyCredentialProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAPIKeyCredentialProviders"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAPIKeyCredentialProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListClientSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientSecretsResponse
func (client *Client) ListClientSecretsWithContext(ctx context.Context, request *ListClientSecretsRequest, runtime *dara.RuntimeOptions) (_result *ListClientSecretsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientName) {
		body["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClientSecrets"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientSecretsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListIdentityProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProvidersResponse
func (client *Client) ListIdentityProvidersWithContext(ctx context.Context, request *ListIdentityProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListIdentityProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentityProviders"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出 OAuth2 凭证提供商
//
// @param request - ListOAuth2CredentialProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOAuth2CredentialProvidersResponse
func (client *Client) ListOAuth2CredentialProvidersWithContext(ctx context.Context, request *ListOAuth2CredentialProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListOAuth2CredentialProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOAuth2CredentialProviders"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOAuth2CredentialProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出权限策略
//
// @param request - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithContext(ctx context.Context, request *ListPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicies"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出网关策略配置
//
// @param request - ListPolicySetAttachedGatewaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicySetAttachedGatewaysResponse
func (client *Client) ListPolicySetAttachedGatewaysWithContext(ctx context.Context, request *ListPolicySetAttachedGatewaysRequest, runtime *dara.RuntimeOptions) (_result *ListPolicySetAttachedGatewaysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GatewayType) {
		body["GatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicySetAttachedGateways"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicySetAttachedGatewaysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出权限策略集合
//
// @param request - ListPolicySetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicySetsResponse
func (client *Client) ListPolicySetsWithContext(ctx context.Context, request *ListPolicySetsRequest, runtime *dara.RuntimeOptions) (_result *ListPolicySetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicySets"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicySetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出网关策略配置
//
// @param request - ListRoleAssignmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoleAssignmentsResponse
func (client *Client) ListRoleAssignmentsWithContext(ctx context.Context, request *ListRoleAssignmentsRequest, runtime *dara.RuntimeOptions) (_result *ListRoleAssignmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PrincipalName) {
		body["PrincipalName"] = request.PrincipalName
	}

	if !dara.IsNil(request.PrincipalType) {
		body["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.RoleName) {
		body["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoleAssignments"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoleAssignmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出Roles
//
// @param request - ListRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithContext(ctx context.Context, request *ListRolesRequest, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoles"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListSAMLIdentityProviderCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSAMLIdentityProviderCertificatesResponse
func (client *Client) ListSAMLIdentityProviderCertificatesWithContext(ctx context.Context, request *ListSAMLIdentityProviderCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListSAMLIdentityProviderCertificatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSAMLIdentityProviderCertificates"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSAMLIdentityProviderCertificatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页列出账户下所有的 API 密钥凭证
//
// @param request - ListTokenVaultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTokenVaultsResponse
func (client *Client) ListTokenVaultsWithContext(ctx context.Context, request *ListTokenVaultsRequest, runtime *dara.RuntimeOptions) (_result *ListTokenVaultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTokenVaults"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTokenVaultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListUserPoolClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPoolClientsResponse
func (client *Client) ListUserPoolClientsWithContext(ctx context.Context, request *ListUserPoolClientsRequest, runtime *dara.RuntimeOptions) (_result *ListUserPoolClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserPoolClients"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPoolClientsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListUserPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPoolsResponse
func (client *Client) ListUserPoolsWithContext(ctx context.Context, request *ListUserPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListUserPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserPools"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPoolsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出用户
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2025-09-01"),
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
// 列出IdentityProvider
//
// @param request - ListWorkloadIdentitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkloadIdentitiesResponse
func (client *Client) ListWorkloadIdentitiesWithContext(ctx context.Context, request *ListWorkloadIdentitiesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkloadIdentitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkloadIdentities"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkloadIdentitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建WorkloadIdentity
//
// @param tmpReq - SetSAMLIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSAMLIdentityProviderResponse
func (client *Client) SetSAMLIdentityProviderWithContext(ctx context.Context, tmpReq *SetSAMLIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *SetSAMLIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetSAMLIdentityProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.X509Certificates) {
		request.X509CertificatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.X509Certificates, dara.String("X509Certificates"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.JITProvisionStatus) {
		body["JITProvisionStatus"] = request.JITProvisionStatus
	}

	if !dara.IsNil(request.JITUpdateStatus) {
		body["JITUpdateStatus"] = request.JITUpdateStatus
	}

	if !dara.IsNil(request.LoginURL) {
		body["LoginURL"] = request.LoginURL
	}

	if !dara.IsNil(request.SAMLBindingType) {
		body["SAMLBindingType"] = request.SAMLBindingType
	}

	if !dara.IsNil(request.SSOStatus) {
		body["SSOStatus"] = request.SSOStatus
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	if !dara.IsNil(request.X509CertificatesShrink) {
		body["X509Certificates"] = request.X509CertificatesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSAMLIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSAMLIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个 API 密钥凭证提供商
//
// @param request - UpdateAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAPIKeyCredentialProviderResponse
func (client *Client) UpdateAPIKeyCredentialProviderWithContext(ctx context.Context, request *UpdateAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateAPIKeyCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		body["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.APIKeyCredentialProviderName) {
		body["APIKeyCredentialProviderName"] = request.APIKeyCredentialProviderName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAPIKeyCredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAPIKeyCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网关策略配置
//
// @param request - UpdateGatewayPolicyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayPolicyConfigResponse
func (client *Client) UpdateGatewayPolicyConfigWithContext(ctx context.Context, request *UpdateGatewayPolicyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayPolicyConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnforcementMode) {
		body["EnforcementMode"] = request.EnforcementMode
	}

	if !dara.IsNil(request.GatewayArn) {
		body["GatewayArn"] = request.GatewayArn
	}

	if !dara.IsNil(request.GatewayType) {
		body["GatewayType"] = request.GatewayType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayPolicyConfig"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayPolicyConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新IdentityProvider
//
// @param tmpReq - UpdateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProviderWithContext(ctx context.Context, tmpReq *UpdateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateIdentityProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedAudience) {
		request.AllowedAudienceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedAudience, dara.String("AllowedAudience"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedAudienceShrink) {
		body["AllowedAudience"] = request.AllowedAudienceShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DiscoveryURL) {
		body["DiscoveryURL"] = request.DiscoveryURL
	}

	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIdentityProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改 OAuth2 凭证提供商
//
// @param tmpReq - UpdateOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOAuth2CredentialProviderResponse
func (client *Client) UpdateOAuth2CredentialProviderWithContext(ctx context.Context, tmpReq *UpdateOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateOAuth2CredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateOAuth2CredentialProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OAuth2ProviderConfig) {
		request.OAuth2ProviderConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OAuth2ProviderConfig, dara.String("OAuth2ProviderConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackURL) {
		body["CallbackURL"] = request.CallbackURL
	}

	if !dara.IsNil(request.CredentialProviderVendor) {
		body["CredentialProviderVendor"] = request.CredentialProviderVendor
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.OAuth2CredentialProviderName) {
		body["OAuth2CredentialProviderName"] = request.OAuth2CredentialProviderName
	}

	if !dara.IsNil(request.OAuth2ProviderConfigShrink) {
		body["OAuth2ProviderConfig"] = request.OAuth2ProviderConfigShrink
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOAuth2CredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOAuth2CredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个权限策略
//
// @param tmpReq - UpdatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicyWithContext(ctx context.Context, tmpReq *UpdatePolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Definition) {
		request.DefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Definition, dara.String("Definition"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DefinitionShrink) {
		body["Definition"] = request.DefinitionShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolicy"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个权限策略集合
//
// @param request - UpdatePolicySetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicySetResponse
func (client *Client) UpdatePolicySetWithContext(ctx context.Context, request *UpdatePolicySetRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolicySetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.PolicySetName) {
		body["PolicySetName"] = request.PolicySetName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolicySet"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolicySetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Role
//
// @param request - UpdateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithContext(ctx context.Context, request *UpdateRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.RoleName) {
		body["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRole"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新凭证库。
//
// @param request - UpdateTokenVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTokenVaultResponse
func (client *Client) UpdateTokenVaultWithContext(ctx context.Context, request *UpdateTokenVaultRequest, runtime *dara.RuntimeOptions) (_result *UpdateTokenVaultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.TokenVaultName) {
		body["TokenVaultName"] = request.TokenVaultName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTokenVault"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTokenVaultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新IdentityProvider
//
// @param request - UpdateUserPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPoolResponse
func (client *Client) UpdateUserPoolWithContext(ctx context.Context, request *UpdateUserPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserPool"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param tmpReq - UpdateUserPoolClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPoolClientResponse
func (client *Client) UpdateUserPoolClientWithContext(ctx context.Context, tmpReq *UpdateUserPoolClientRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserPoolClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateUserPoolClientShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RedirectURIs) {
		request.RedirectURIsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RedirectURIs, dara.String("RedirectURIs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessTokenValidity) {
		body["AccessTokenValidity"] = request.AccessTokenValidity
	}

	if !dara.IsNil(request.ClientName) {
		body["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.EnforcePKCE) {
		body["EnforcePKCE"] = request.EnforcePKCE
	}

	if !dara.IsNil(request.RedirectURIsShrink) {
		body["RedirectURIs"] = request.RedirectURIsShrink
	}

	if !dara.IsNil(request.RefreshTokenValidity) {
		body["RefreshTokenValidity"] = request.RefreshTokenValidity
	}

	if !dara.IsNil(request.SecretRequired) {
		body["SecretRequired"] = request.SecretRequired
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserPoolClient"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserPoolClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param tmpReq - UpdateWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkloadIdentityResponse
func (client *Client) UpdateWorkloadIdentityWithContext(ctx context.Context, tmpReq *UpdateWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkloadIdentityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedResourceOAuth2ReturnURLs) {
		request.AllowedResourceOAuth2ReturnURLsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedResourceOAuth2ReturnURLs, dara.String("AllowedResourceOAuth2ReturnURLs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedResourceOAuth2ReturnURLsShrink) {
		body["AllowedResourceOAuth2ReturnURLs"] = request.AllowedResourceOAuth2ReturnURLsShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.SessionBindingEnabled) {
		body["SessionBindingEnabled"] = request.SessionBindingEnabled
	}

	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkloadIdentity"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkloadIdentityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
