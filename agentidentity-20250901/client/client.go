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
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("agentidentity"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建应用
//
// @param request - AddSAMLIdentityProviderCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSAMLIdentityProviderCertificateResponse
func (client *Client) AddSAMLIdentityProviderCertificateWithOptions(request *AddSAMLIdentityProviderCertificateRequest, runtime *dara.RuntimeOptions) (_result *AddSAMLIdentityProviderCertificateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AddSAMLIdentityProviderCertificateRequest
//
// @return AddSAMLIdentityProviderCertificateResponse
func (client *Client) AddSAMLIdentityProviderCertificate(request *AddSAMLIdentityProviderCertificateRequest) (_result *AddSAMLIdentityProviderCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSAMLIdentityProviderCertificateResponse{}
	_body, _err := client.AddSAMLIdentityProviderCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a policy set to a gateway.
//
// Description:
//
// Attaches a policy set to a target gateway.
//
// The attachment relationship is 1:1. A gateway can have only one policy set attached at a time. If you attach a policy set multiple times, the latest attachment takes effect. You can specify the policy enforcement mode by using the EnforcementMode parameter. Valid values: ENFORCE (blocks non-compliant requests) and LOG_ONLY (records non-compliant requests but allows them, which is commonly used for canary release observation).
//
// @param request - AttachPolicySetToGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicySetToGatewayResponse
func (client *Client) AttachPolicySetToGatewayWithOptions(request *AttachPolicySetToGatewayRequest, runtime *dara.RuntimeOptions) (_result *AttachPolicySetToGatewayResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a policy set to a gateway.
//
// Description:
//
// Attaches a policy set to a target gateway.
//
// The attachment relationship is 1:1. A gateway can have only one policy set attached at a time. If you attach a policy set multiple times, the latest attachment takes effect. You can specify the policy enforcement mode by using the EnforcementMode parameter. Valid values: ENFORCE (blocks non-compliant requests) and LOG_ONLY (records non-compliant requests but allows them, which is commonly used for canary release observation).
//
// @param request - AttachPolicySetToGatewayRequest
//
// @return AttachPolicySetToGatewayResponse
func (client *Client) AttachPolicySetToGateway(request *AttachPolicySetToGatewayRequest) (_result *AttachPolicySetToGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachPolicySetToGatewayResponse{}
	_body, _err := client.AttachPolicySetToGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an API key 
//
// @param request - CreateAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAPIKeyCredentialProviderResponse
func (client *Client) CreateAPIKeyCredentialProviderWithOptions(request *CreateAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateAPIKeyCredentialProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API key credential provider.
//
// @param request - CreateAPIKeyCredentialProviderRequest
//
// @return CreateAPIKeyCredentialProviderResponse
func (client *Client) CreateAPIKeyCredentialProvider(request *CreateAPIKeyCredentialProviderRequest) (_result *CreateAPIKeyCredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAPIKeyCredentialProviderResponse{}
	_body, _err := client.CreateAPIKeyCredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workload identity.
//
// Description:
//
// The response includes the plaintext client secret (ClientSecretValue). For security reasons, the plaintext secret is returned only once upon creation and cannot be retrieved later. Store it securely.
//
// @param request - CreateClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientSecretResponse
func (client *Client) CreateClientSecretWithOptions(request *CreateClientSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateClientSecretResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workload identity.
//
// Description:
//
// The response includes the plaintext client secret (ClientSecretValue). For security reasons, the plaintext secret is returned only once upon creation and cannot be retrieved later. Store it securely.
//
// @param request - CreateClientSecretRequest
//
// @return CreateClientSecretResponse
func (client *Client) CreateClientSecret(request *CreateClientSecretRequest) (_result *CreateClientSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClientSecretResponse{}
	_body, _err := client.CreateClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateIdentityProviderWithOptions(tmpReq *CreateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateIdentityProviderRequest
//
// @return CreateIdentityProviderResponse
func (client *Client) CreateIdentityProvider(request *CreateIdentityProviderRequest) (_result *CreateIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIdentityProviderResponse{}
	_body, _err := client.CreateIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an OAuth 2.0 credential provider.
//
// @param tmpReq - CreateOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOAuth2CredentialProviderResponse
func (client *Client) CreateOAuth2CredentialProviderWithOptions(tmpReq *CreateOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateOAuth2CredentialProviderResponse, _err error) {
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

	if !dara.IsNil(request.OAuthType) {
		body["OAuthType"] = request.OAuthType
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an OAuth 2.0 credential provider.
//
// @param request - CreateOAuth2CredentialProviderRequest
//
// @return CreateOAuth2CredentialProviderResponse
func (client *Client) CreateOAuth2CredentialProvider(request *CreateOAuth2CredentialProviderRequest) (_result *CreateOAuth2CredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOAuth2CredentialProviderResponse{}
	_body, _err := client.CreateOAuth2CredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access policy.
//
// Description:
//
// Calls the CreatePolicy operation to add a policy to an existing policy set. The policy defines specific control rules, such as throttling thresholds, admission conditions, and allow or block actions, through the Definition field. After a policy set is mounted to a gateway, requests on the gateway are evaluated and controlled based on the merged results of all policies in the policy set.
//
// @param tmpReq - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(tmpReq *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access policy.
//
// Description:
//
// Calls the CreatePolicy operation to add a policy to an existing policy set. The policy defines specific control rules, such as throttling thresholds, admission conditions, and allow or block actions, through the Definition field. After a policy set is mounted to a gateway, requests on the gateway are evaluated and controlled based on the merged results of all policies in the policy set.
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access policy set.
//
// Description:
//
// Calls CreatePolicySet to create a policy set in the current account. A policy set is a container for a group of policies. After a policy set is mounted to a gateway, all policies in the set evaluate requests to the gateway as a whole.
//
// @param request - CreatePolicySetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicySetResponse
func (client *Client) CreatePolicySetWithOptions(request *CreatePolicySetRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicySetResponse, _err error) {
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

	if !dara.IsNil(request.SourcePlatform) {
		body["SourcePlatform"] = request.SourcePlatform
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access policy set.
//
// Description:
//
// Calls CreatePolicySet to create a policy set in the current account. A policy set is a container for a group of policies. After a policy set is mounted to a gateway, all policies in the set evaluate requests to the gateway as a whole.
//
// @param request - CreatePolicySetRequest
//
// @return CreatePolicySetResponse
func (client *Client) CreatePolicySet(request *CreatePolicySetRequest) (_result *CreatePolicySetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicySetResponse{}
	_body, _err := client.CreatePolicySetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a role.
//
// Description:
//
// Users in a user pool can be allocated roles to enable quick authorization. Administrators can specify permissions for roles in an access policy.
//
// When a user pool is created, the platform automatically provisions a default role, and newly created users are automatically allocated this default role. This operation creates a custom role of the Manual type, which must be explicitly allocated to users by calling CreateRoleAssignment.
//
// @param request - CreateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a role.
//
// Description:
//
// Users in a user pool can be allocated roles to enable quick authorization. Administrators can specify permissions for roles in an access policy.
//
// When a user pool is created, the platform automatically provisions a default role, and newly created users are automatically allocated this default role. This operation creates a custom role of the Manual type, which must be explicitly allocated to users by calling CreateRoleAssignment.
//
// @param request - CreateRoleRequest
//
// @return CreateRoleResponse
func (client *Client) CreateRole(request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a policy set with a gateway.
//
// Description:
//
// Users must be assigned roles to complete authorization. When a user is created through JIT, a default role is automatically assigned. Administrators can use this operation to assign additional custom roles to users.
//
// PrincipalType currently supports only User. A single identity can have multiple role assignments.
//
// @param request - CreateRoleAssignmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleAssignmentResponse
func (client *Client) CreateRoleAssignmentWithOptions(request *CreateRoleAssignmentRequest, runtime *dara.RuntimeOptions) (_result *CreateRoleAssignmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a policy set with a gateway.
//
// Description:
//
// Users must be assigned roles to complete authorization. When a user is created through JIT, a default role is automatically assigned. Administrators can use this operation to assign additional custom roles to users.
//
// PrincipalType currently supports only User. A single identity can have multiple role assignments.
//
// @param request - CreateRoleAssignmentRequest
//
// @return CreateRoleAssignmentResponse
func (client *Client) CreateRoleAssignment(request *CreateRoleAssignmentRequest) (_result *CreateRoleAssignmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRoleAssignmentResponse{}
	_body, _err := client.CreateRoleAssignmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a token vault.
//
// Description:
//
// Calls the CreateTokenVault operation to create a token vault within the current Alibaba Cloud account. A token vault is a container used by the AgentIdentity service to securely store sensitive credentials such as OAuth2 tokens and API keys. Each token is encrypted when it is written to the vault.
//
// You can specify the encryption configuration when you create the vault. By default, a service-managed key (SERVICE_MANAGED_KEY) is used. For greater control over the key, you can use a customer-managed Key Management Service (KMS) key (CUSTOMER_MANAGED_KEY). To use a customer-managed key, specify the Alibaba Cloud Resource Name (ARN) of the KMS key and use the RoleArn parameter to authorize the AgentIdentity service to access the KMS key.
//
// @param tmpReq - CreateTokenVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTokenVaultResponse
func (client *Client) CreateTokenVaultWithOptions(tmpReq *CreateTokenVaultRequest, runtime *dara.RuntimeOptions) (_result *CreateTokenVaultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a token vault.
//
// Description:
//
// Calls the CreateTokenVault operation to create a token vault within the current Alibaba Cloud account. A token vault is a container used by the AgentIdentity service to securely store sensitive credentials such as OAuth2 tokens and API keys. Each token is encrypted when it is written to the vault.
//
// You can specify the encryption configuration when you create the vault. By default, a service-managed key (SERVICE_MANAGED_KEY) is used. For greater control over the key, you can use a customer-managed Key Management Service (KMS) key (CUSTOMER_MANAGED_KEY). To use a customer-managed key, specify the Alibaba Cloud Resource Name (ARN) of the KMS key and use the RoleArn parameter to authorize the AgentIdentity service to access the KMS key.
//
// @param request - CreateTokenVaultRequest
//
// @return CreateTokenVaultResponse
func (client *Client) CreateTokenVault(request *CreateTokenVaultRequest) (_result *CreateTokenVaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTokenVaultResponse{}
	_body, _err := client.CreateTokenVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user in a user pool.
//
// Description:
//
// Creates a user in a specified user pool. The type of the created user is Manual, and the default status is Enabled. The username must be unique within the user pool and can contain letters, digits, underscores (_), periods (.), and hyphens (-), and must be 3 to 128 characters in length. The email address must be unique within the user pool. After the user is created, you can call the SetUserPassword operation to set a logon password for the user and call the CreateRoleAssignment operation to assign a role to the user.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		body["Email"] = request.Email
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

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
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2025-09-01"),
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
// Creates a user in a user pool.
//
// Description:
//
// Creates a user in a specified user pool. The type of the created user is Manual, and the default status is Enabled. The username must be unique within the user pool and can contain letters, digits, underscores (_), periods (.), and hyphens (-), and must be 3 to 128 characters in length. The email address must be unique within the user pool. After the user is created, you can call the SetUserPassword operation to set a logon password for the user and call the CreateRoleAssignment operation to assign a role to the user.
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
// Creates a user pool.
//
// Description:
//
// A user pool is a secure and isolated user folder designed by the Agent Identity platform for the AI agent ecosystem. You can create multiple user pools under a single Alibaba Cloud account for the independent user identity management of different AI agents. When you create a user pool, you only need to specify the name (UserPoolName). The description is optional. After the user pool is created, the platform automatically generates a default role for the user pool. Users subsequently created through SAML SSO just-in-time (JIT) synchronization are automatically allocated this default role.
//
// @param request - CreateUserPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserPoolResponse
func (client *Client) CreateUserPoolWithOptions(request *CreateUserPoolRequest, runtime *dara.RuntimeOptions) (_result *CreateUserPoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user pool.
//
// Description:
//
// A user pool is a secure and isolated user folder designed by the Agent Identity platform for the AI agent ecosystem. You can create multiple user pools under a single Alibaba Cloud account for the independent user identity management of different AI agents. When you create a user pool, you only need to specify the name (UserPoolName). The description is optional. After the user pool is created, the platform automatically generates a default role for the user pool. Users subsequently created through SAML SSO just-in-time (JIT) synchronization are automatically allocated this default role.
//
// @param request - CreateUserPoolRequest
//
// @return CreateUserPoolResponse
func (client *Client) CreateUserPool(request *CreateUserPoolRequest) (_result *CreateUserPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserPoolResponse{}
	_body, _err := client.CreateUserPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a WorkloadIdentity.
//
// Description:
//
// A client represents an application that is integrated with a user pool. When you create a client, you must specify the client name and redirect URIs. You can configure whether an application secret is required, whether to enforce Proof Key for Code Exchange (PKCE), the access token validity period, and the refresh token validity period. The application scopes are generated by the platform by default. If SecretRequired is set to true, you must call CreateClientSecret to create a client secret.
//
// @param tmpReq - CreateUserPoolClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserPoolClientResponse
func (client *Client) CreateUserPoolClientWithOptions(tmpReq *CreateUserPoolClientRequest, runtime *dara.RuntimeOptions) (_result *CreateUserPoolClientResponse, _err error) {
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

	if !dara.IsNil(tmpReq.Scopes) {
		request.ScopesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scopes, dara.String("Scopes"), dara.String("json"))
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

	if !dara.IsNil(request.ScopesShrink) {
		body["Scopes"] = request.ScopesShrink
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a WorkloadIdentity.
//
// Description:
//
// A client represents an application that is integrated with a user pool. When you create a client, you must specify the client name and redirect URIs. You can configure whether an application secret is required, whether to enforce Proof Key for Code Exchange (PKCE), the access token validity period, and the refresh token validity period. The application scopes are generated by the platform by default. If SecretRequired is set to true, you must call CreateClientSecret to create a client secret.
//
// @param request - CreateUserPoolClientRequest
//
// @return CreateUserPoolClientResponse
func (client *Client) CreateUserPoolClient(request *CreateUserPoolClientRequest) (_result *CreateUserPoolClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserPoolClientResponse{}
	_body, _err := client.CreateUserPoolClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a workload identity for an AI Agent to uniquely identify it.
//
// @param tmpReq - CreateWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkloadIdentityResponse
func (client *Client) CreateWorkloadIdentityWithOptions(tmpReq *CreateWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkloadIdentityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedConsentCallbackURLs) {
		request.AllowedConsentCallbackURLsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedConsentCallbackURLs, dara.String("AllowedConsentCallbackURLs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AllowedResourceOAuth2ReturnURLs) {
		request.AllowedResourceOAuth2ReturnURLsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedResourceOAuth2ReturnURLs, dara.String("AllowedResourceOAuth2ReturnURLs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedConsentCallbackURLsShrink) {
		body["AllowedConsentCallbackURLs"] = request.AllowedConsentCallbackURLsShrink
	}

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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a workload identity for an AI Agent to uniquely identify it.
//
// @param request - CreateWorkloadIdentityRequest
//
// @return CreateWorkloadIdentityResponse
func (client *Client) CreateWorkloadIdentity(request *CreateWorkloadIdentityRequest) (_result *CreateWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWorkloadIdentityResponse{}
	_body, _err := client.CreateWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified API key credential provider.
//
// @param request - DeleteAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAPIKeyCredentialProviderResponse
func (client *Client) DeleteAPIKeyCredentialProviderWithOptions(request *DeleteAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteAPIKeyCredentialProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified API key credential provider.
//
// @param request - DeleteAPIKeyCredentialProviderRequest
//
// @return DeleteAPIKeyCredentialProviderResponse
func (client *Client) DeleteAPIKeyCredentialProvider(request *DeleteAPIKeyCredentialProviderRequest) (_result *DeleteAPIKeyCredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAPIKeyCredentialProviderResponse{}
	_body, _err := client.DeleteAPIKeyCredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an identity provider.
//
// Description:
//
// Deletes a specified client secret by client ID and key ID. The secret becomes invalid immediately after deletion, and all OAuth authentication flows that use this secret will fail.
//
// @param request - DeleteClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientSecretResponse
func (client *Client) DeleteClientSecretWithOptions(request *DeleteClientSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientSecretResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an identity provider.
//
// Description:
//
// Deletes a specified client secret by client ID and key ID. The secret becomes invalid immediately after deletion, and all OAuth authentication flows that use this secret will fail.
//
// @param request - DeleteClientSecretRequest
//
// @return DeleteClientSecretResponse
func (client *Client) DeleteClientSecret(request *DeleteClientSecretRequest) (_result *DeleteClientSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClientSecretResponse{}
	_body, _err := client.DeleteClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified identity provider.
//
// @param request - DeleteIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdentityProviderResponse
func (client *Client) DeleteIdentityProviderWithOptions(request *DeleteIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteIdentityProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified identity provider.
//
// @param request - DeleteIdentityProviderRequest
//
// @return DeleteIdentityProviderResponse
func (client *Client) DeleteIdentityProvider(request *DeleteIdentityProviderRequest) (_result *DeleteIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIdentityProviderResponse{}
	_body, _err := client.DeleteIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified OAuth 2.0 credential provider.
//
// @param request - DeleteOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOAuth2CredentialProviderResponse
func (client *Client) DeleteOAuth2CredentialProviderWithOptions(request *DeleteOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteOAuth2CredentialProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified OAuth 2.0 credential provider.
//
// @param request - DeleteOAuth2CredentialProviderRequest
//
// @return DeleteOAuth2CredentialProviderResponse
func (client *Client) DeleteOAuth2CredentialProvider(request *DeleteOAuth2CredentialProviderRequest) (_result *DeleteOAuth2CredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOAuth2CredentialProviderResponse{}
	_body, _err := client.DeleteOAuth2CredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access policy.
//
// Description:
//
// Calls DeletePolicy to remove a policy from a specified policy set. The policy becomes invalid immediately after deletion, and gateways that depend on the policy set do not need to be remounted.
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access policy.
//
// Description:
//
// Calls DeletePolicy to remove a policy from a specified policy set. The policy becomes invalid immediately after deletion, and gateways that depend on the policy set do not need to be remounted.
//
// @param request - DeletePolicyRequest
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access policy set.
//
// Description:
//
// Calls DeletePolicySet to delete a policy set that is no longer in use.
//
// Before deleting the policy set, ensure that it is unmounted from all gateways (you can query this by calling ListPolicySetAttachedGateways) and contains no policies. Otherwise, the deletion fails.
//
// @param request - DeletePolicySetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicySetResponse
func (client *Client) DeletePolicySetWithOptions(request *DeletePolicySetRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicySetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access policy set.
//
// Description:
//
// Calls DeletePolicySet to delete a policy set that is no longer in use.
//
// Before deleting the policy set, ensure that it is unmounted from all gateways (you can query this by calling ListPolicySetAttachedGateways) and contains no policies. Otherwise, the deletion fails.
//
// @param request - DeletePolicySetRequest
//
// @return DeletePolicySetResponse
func (client *Client) DeletePolicySet(request *DeletePolicySetRequest) (_result *DeletePolicySetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicySetResponse{}
	_body, _err := client.DeletePolicySetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a role.
//
// Description:
//
// Before deletion, the system checks whether the role has any bound users. If role assignments exist, the deletion is rejected. You must call DeleteRoleAssignment to remove the assignments before deleting the role. Default roles (Type=Default) cannot be deleted.
//
// @param request - DeleteRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a role.
//
// Description:
//
// Before deletion, the system checks whether the role has any bound users. If role assignments exist, the deletion is rejected. You must call DeleteRoleAssignment to remove the assignments before deleting the role. Default roles (Type=Default) cannot be deleted.
//
// @param request - DeleteRoleRequest
//
// @return DeleteRoleResponse
func (client *Client) DeleteRole(request *DeleteRoleRequest) (_result *DeleteRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoleResponse{}
	_body, _err := client.DeleteRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a policy set from a gateway.
//
// Description:
//
// Removes the role assignment relationship between an identity and a role.
//
// @param request - DeleteRoleAssignmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleAssignmentResponse
func (client *Client) DeleteRoleAssignmentWithOptions(request *DeleteRoleAssignmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoleAssignmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a policy set from a gateway.
//
// Description:
//
// Removes the role assignment relationship between an identity and a role.
//
// @param request - DeleteRoleAssignmentRequest
//
// @return DeleteRoleAssignmentResponse
func (client *Client) DeleteRoleAssignment(request *DeleteRoleAssignmentRequest) (_result *DeleteRoleAssignmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoleAssignmentResponse{}
	_body, _err := client.DeleteRoleAssignmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workload identity.
//
// Description:
//
// Cleans up expired or unused certificates after certificate rotation is complete. Before deleting the certificate, ensure that it is no longer used for SAML assertion signature verification. Otherwise, SSO logon may fail.
//
// @param request - DeleteSAMLIdentityProviderCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSAMLIdentityProviderCertificateResponse
func (client *Client) DeleteSAMLIdentityProviderCertificateWithOptions(request *DeleteSAMLIdentityProviderCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSAMLIdentityProviderCertificateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workload identity.
//
// Description:
//
// Cleans up expired or unused certificates after certificate rotation is complete. Before deleting the certificate, ensure that it is no longer used for SAML assertion signature verification. Otherwise, SSO logon may fail.
//
// @param request - DeleteSAMLIdentityProviderCertificateRequest
//
// @return DeleteSAMLIdentityProviderCertificateResponse
func (client *Client) DeleteSAMLIdentityProviderCertificate(request *DeleteSAMLIdentityProviderCertificateRequest) (_result *DeleteSAMLIdentityProviderCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSAMLIdentityProviderCertificateResponse{}
	_body, _err := client.DeleteSAMLIdentityProviderCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified token vault.
//
// Description:
//
// Call DeleteTokenVault to delete a token vault that is no longer in use.
//
// Before deletion, ensure that the token vault is not referenced by any credential provider. Otherwise, the deletion may cause credentials that depend on the token vault to become unavailable. This operation cannot be undone.
//
// @param request - DeleteTokenVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTokenVaultResponse
func (client *Client) DeleteTokenVaultWithOptions(request *DeleteTokenVaultRequest, runtime *dara.RuntimeOptions) (_result *DeleteTokenVaultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified token vault.
//
// Description:
//
// Call DeleteTokenVault to delete a token vault that is no longer in use.
//
// Before deletion, ensure that the token vault is not referenced by any credential provider. Otherwise, the deletion may cause credentials that depend on the token vault to become unavailable. This operation cannot be undone.
//
// @param request - DeleteTokenVaultRequest
//
// @return DeleteTokenVaultResponse
func (client *Client) DeleteTokenVault(request *DeleteTokenVaultRequest) (_result *DeleteTokenVaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTokenVaultResponse{}
	_body, _err := client.DeleteTokenVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user.
//
// Description:
//
// Deleting a user forcibly clears all role assignments of the user in the user pool. This operation is irreversible. Proceed with caution.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user.
//
// Description:
//
// Deleting a user forcibly clears all role assignments of the user in the user pool. This operation is irreversible. Proceed with caution.
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
// Deletes a user pool.
//
// Description:
//
// Before deleting a user pool, you must clean up all resources in the user pool, including all users, roles, clients, SAML IdPs, and role assignments. Calls to this operation are rejected if the resources are not cleaned up. This operation is irreversible. Proceed with caution.
//
// @param request - DeleteUserPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPoolResponse
func (client *Client) DeleteUserPoolWithOptions(request *DeleteUserPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserPoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user pool.
//
// Description:
//
// Before deleting a user pool, you must clean up all resources in the user pool, including all users, roles, clients, SAML IdPs, and role assignments. Calls to this operation are rejected if the resources are not cleaned up. This operation is irreversible. Proceed with caution.
//
// @param request - DeleteUserPoolRequest
//
// @return DeleteUserPoolResponse
func (client *Client) DeleteUserPool(request *DeleteUserPoolRequest) (_result *DeleteUserPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserPoolResponse{}
	_body, _err := client.DeleteUserPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a WorkloadIdentity.
//
// Description:
//
// Deleting a client also removes all its associated client secrets. After deletion, the client can no longer be used for OAuth authentication and authorization flows. Proceed with caution.
//
// @param request - DeleteUserPoolClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPoolClientResponse
func (client *Client) DeleteUserPoolClientWithOptions(request *DeleteUserPoolClientRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserPoolClientResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a WorkloadIdentity.
//
// Description:
//
// Deleting a client also removes all its associated client secrets. After deletion, the client can no longer be used for OAuth authentication and authorization flows. Proceed with caution.
//
// @param request - DeleteUserPoolClientRequest
//
// @return DeleteUserPoolClientResponse
func (client *Client) DeleteUserPoolClient(request *DeleteUserPoolClientRequest) (_result *DeleteUserPoolClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserPoolClientResponse{}
	_body, _err := client.DeleteUserPoolClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified workload identity.
//
// @param request - DeleteWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkloadIdentityResponse
func (client *Client) DeleteWorkloadIdentityWithOptions(request *DeleteWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkloadIdentityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified workload identity.
//
// @param request - DeleteWorkloadIdentityRequest
//
// @return DeleteWorkloadIdentityResponse
func (client *Client) DeleteWorkloadIdentity(request *DeleteWorkloadIdentityRequest) (_result *DeleteWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWorkloadIdentityResponse{}
	_body, _err := client.DeleteWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches a policy set from a gateway.
//
// Description:
//
// Detaches a policy set from a gateway. After the policy set is detached, requests on the gateway are no longer governed by the policies in the policy set.
//
// @param request - DetachPolicySetFromGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicySetFromGatewayResponse
func (client *Client) DetachPolicySetFromGatewayWithOptions(request *DetachPolicySetFromGatewayRequest, runtime *dara.RuntimeOptions) (_result *DetachPolicySetFromGatewayResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a policy set from a gateway.
//
// Description:
//
// Detaches a policy set from a gateway. After the policy set is detached, requests on the gateway are no longer governed by the policies in the policy set.
//
// @param request - DetachPolicySetFromGatewayRequest
//
// @return DetachPolicySetFromGatewayResponse
func (client *Client) DetachPolicySetFromGateway(request *DetachPolicySetFromGatewayRequest) (_result *DetachPolicySetFromGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachPolicySetFromGatewayResponse{}
	_body, _err := client.DetachPolicySetFromGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified API key credential provider.
//
// @param request - GetAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAPIKeyCredentialProviderResponse
func (client *Client) GetAPIKeyCredentialProviderWithOptions(request *GetAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *GetAPIKeyCredentialProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified API key credential provider.
//
// @param request - GetAPIKeyCredentialProviderRequest
//
// @return GetAPIKeyCredentialProviderResponse
func (client *Client) GetAPIKeyCredentialProvider(request *GetAPIKeyCredentialProviderRequest) (_result *GetAPIKeyCredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAPIKeyCredentialProviderResponse{}
	_body, _err := client.GetAPIKeyCredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the gateway policy configuration.
//
// Description:
//
// Calls the GetGatewayPolicyConfig operation to query the currently effective policy configuration of a gateway, including the ARN of the mounted policy set and the enforcement mode.
//
// @param request - GetGatewayPolicyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayPolicyConfigResponse
func (client *Client) GetGatewayPolicyConfigWithOptions(request *GetGatewayPolicyConfigRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayPolicyConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the gateway policy configuration.
//
// Description:
//
// Calls the GetGatewayPolicyConfig operation to query the currently effective policy configuration of a gateway, including the ARN of the mounted policy set and the enforcement mode.
//
// @param request - GetGatewayPolicyConfigRequest
//
// @return GetGatewayPolicyConfigResponse
func (client *Client) GetGatewayPolicyConfig(request *GetGatewayPolicyConfigRequest) (_result *GetGatewayPolicyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayPolicyConfigResponse{}
	_body, _err := client.GetGatewayPolicyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetIdentityProviderWithOptions(request *GetIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetIdentityProviderResponse
func (client *Client) GetIdentityProvider(request *GetIdentityProviderRequest) (_result *GetIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIdentityProviderResponse{}
	_body, _err := client.GetIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询身份源同步范围与统计（已选部门有效性 + 展开部门数/预计人数）
//
// Description:
//
// 目标用户池需已配置指定类型的身份源。本接口使用已保存的身份源配置，无需传入钉钉或飞书应用密钥。
//
// 同步范围支持 AllAuthorized 和 SelectedDepartments。未显式配置范围时默认为 AllAuthorized，表示第三方应用授权可见范围内的全部用户。SelectedDepartments 包含选中部门及其全部后代部门；人员有任一当前所属部门命中展开集合，即命中同步范围。
//
// 本接口不修改范围、不创建人员同步任务。查询可能触发后台统计；EstimateStatus 表示范围统计状态，不是人员同步任务状态。统计结果可能复用缓存，也可能来自同一范围最近一次同步结果；跨部门重复用户可能重复计数，不应视为实时、精确去重人数。
//
// 对应身份源未配置时返回 EntityNotExists.IdpConfiguration；目录权限或通讯录可见范围不足时，返回 DingTalkApiPermissionDenied 或 FeishuApiPermissionDenied，不会将权限问题当作部门不存在。
//
// @param request - GetIdentityProviderSyncScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderSyncScopeResponse
func (client *Client) GetIdentityProviderSyncScopeWithOptions(request *GetIdentityProviderSyncScopeRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderSyncScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderType) {
		body["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityProviderSyncScope"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderSyncScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询身份源同步范围与统计（已选部门有效性 + 展开部门数/预计人数）
//
// Description:
//
// 目标用户池需已配置指定类型的身份源。本接口使用已保存的身份源配置，无需传入钉钉或飞书应用密钥。
//
// 同步范围支持 AllAuthorized 和 SelectedDepartments。未显式配置范围时默认为 AllAuthorized，表示第三方应用授权可见范围内的全部用户。SelectedDepartments 包含选中部门及其全部后代部门；人员有任一当前所属部门命中展开集合，即命中同步范围。
//
// 本接口不修改范围、不创建人员同步任务。查询可能触发后台统计；EstimateStatus 表示范围统计状态，不是人员同步任务状态。统计结果可能复用缓存，也可能来自同一范围最近一次同步结果；跨部门重复用户可能重复计数，不应视为实时、精确去重人数。
//
// 对应身份源未配置时返回 EntityNotExists.IdpConfiguration；目录权限或通讯录可见范围不足时，返回 DingTalkApiPermissionDenied 或 FeishuApiPermissionDenied，不会将权限问题当作部门不存在。
//
// @param request - GetIdentityProviderSyncScopeRequest
//
// @return GetIdentityProviderSyncScopeResponse
func (client *Client) GetIdentityProviderSyncScope(request *GetIdentityProviderSyncScopeRequest) (_result *GetIdentityProviderSyncScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIdentityProviderSyncScopeResponse{}
	_body, _err := client.GetIdentityProviderSyncScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the logon-related configurations of a user pool.
//
// Description:
//
// Retrieves the logon configuration and password policy of a specified user pool. The logon configuration includes preference settings such as whether to enable password logon. The password policy includes rules for password complexity, length limits, and logon failure lockouts. The password policy is read-only and does not support custom modifications.
//
// @param request - GetLoginPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginPreferenceResponse
func (client *Client) GetLoginPreferenceWithOptions(request *GetLoginPreferenceRequest, runtime *dara.RuntimeOptions) (_result *GetLoginPreferenceResponse, _err error) {
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
		Action:      dara.String("GetLoginPreference"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoginPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the logon-related configurations of a user pool.
//
// Description:
//
// Retrieves the logon configuration and password policy of a specified user pool. The logon configuration includes preference settings such as whether to enable password logon. The password policy includes rules for password complexity, length limits, and logon failure lockouts. The password policy is read-only and does not support custom modifications.
//
// @param request - GetLoginPreferenceRequest
//
// @return GetLoginPreferenceResponse
func (client *Client) GetLoginPreference(request *GetLoginPreferenceRequest) (_result *GetLoginPreferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLoginPreferenceResponse{}
	_body, _err := client.GetLoginPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified OAuth 2.0 credential provider.
//
// @param request - GetOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOAuth2CredentialProviderResponse
func (client *Client) GetOAuth2CredentialProviderWithOptions(request *GetOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *GetOAuth2CredentialProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified OAuth 2.0 credential provider.
//
// @param request - GetOAuth2CredentialProviderRequest
//
// @return GetOAuth2CredentialProviderResponse
func (client *Client) GetOAuth2CredentialProvider(request *GetOAuth2CredentialProviderRequest) (_result *GetOAuth2CredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOAuth2CredentialProviderResponse{}
	_body, _err := client.GetOAuth2CredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an access policy.
//
// Description:
//
// Calls GetPolicy to retrieve the complete configuration of a single policy, including its definition, associated policy set, and creation and update times.
//
// @param request - GetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an access policy.
//
// Description:
//
// Calls GetPolicy to retrieve the complete configuration of a single policy, including its definition, associated policy set, and creation and update times.
//
// @param request - GetPolicyRequest
//
// @return GetPolicyResponse
func (client *Client) GetPolicy(request *GetPolicyRequest) (_result *GetPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an access policy set.
//
// Description:
//
// Calls GetPolicySet to query the complete configuration of a single policy set, including the ARN, description, creation time, and update time.
//
// @param request - GetPolicySetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicySetResponse
func (client *Client) GetPolicySetWithOptions(request *GetPolicySetRequest, runtime *dara.RuntimeOptions) (_result *GetPolicySetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an access policy set.
//
// Description:
//
// Calls GetPolicySet to query the complete configuration of a single policy set, including the ARN, description, creation time, and update time.
//
// @param request - GetPolicySetRequest
//
// @return GetPolicySetResponse
func (client *Client) GetPolicySet(request *GetPolicySetRequest) (_result *GetPolicySetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicySetResponse{}
	_body, _err := client.GetPolicySetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a role.
//
// Description:
//
// Returns the complete information about a specified role, including the role ID, name, description, type (Default or Manual), and creation and update times.
//
// @param request - GetRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleResponse
func (client *Client) GetRoleWithOptions(request *GetRoleRequest, runtime *dara.RuntimeOptions) (_result *GetRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a role.
//
// Description:
//
// Returns the complete information about a specified role, including the role ID, name, description, type (Default or Manual), and creation and update times.
//
// @param request - GetRoleRequest
//
// @return GetRoleResponse
func (client *Client) GetRole(request *GetRoleRequest) (_result *GetRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRoleResponse{}
	_body, _err := client.GetRoleWithOptions(request, runtime)
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
// Description:
//
// Returns the complete configuration of the SAML IdP, including the EntityId, logon URL, X509 certificate list, and the enabling status of SSO and JIT synchronization.
//
// @param request - GetSAMLIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSAMLIdentityProviderResponse
func (client *Client) GetSAMLIdentityProviderWithOptions(request *GetSAMLIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetSAMLIdentityProviderResponse, _err error) {
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
// Description:
//
// Returns the complete configuration of the SAML IdP, including the EntityId, logon URL, X509 certificate list, and the enabling status of SSO and JIT synchronization.
//
// @param request - GetSAMLIdentityProviderRequest
//
// @return GetSAMLIdentityProviderResponse
func (client *Client) GetSAMLIdentityProvider(request *GetSAMLIdentityProviderRequest) (_result *GetSAMLIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSAMLIdentityProviderResponse{}
	_body, _err := client.GetSAMLIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetSAMLServiceProviderInfoWithOptions(request *GetSAMLServiceProviderInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSAMLServiceProviderInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetSAMLServiceProviderInfoResponse
func (client *Client) GetSAMLServiceProviderInfo(request *GetSAMLServiceProviderInfoRequest) (_result *GetSAMLServiceProviderInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSAMLServiceProviderInfoResponse{}
	_body, _err := client.GetSAMLServiceProviderInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the specified identity provider.
//
// Description:
//
// Queries the SSO configuration of a specific type of identity provider in a specified user pool. Returns information such as the enabling status, synchronization callback status, and metadata of the IdP. For security reasons, sensitive fields such as appSecret, encryptKey, and verificationToken are not returned in the response.
//
// @param request - GetSpecificIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpecificIdentityProviderResponse
func (client *Client) GetSpecificIdentityProviderWithOptions(request *GetSpecificIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetSpecificIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderType) {
		body["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSpecificIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSpecificIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the specified identity provider.
//
// Description:
//
// Queries the SSO configuration of a specific type of identity provider in a specified user pool. Returns information such as the enabling status, synchronization callback status, and metadata of the IdP. For security reasons, sensitive fields such as appSecret, encryptKey, and verificationToken are not returned in the response.
//
// @param request - GetSpecificIdentityProviderRequest
//
// @return GetSpecificIdentityProviderResponse
func (client *Client) GetSpecificIdentityProvider(request *GetSpecificIdentityProviderRequest) (_result *GetSpecificIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSpecificIdentityProviderResponse{}
	_body, _err := client.GetSpecificIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the detailed configuration of a specified token vault.
//
// Description:
//
// Calls GetTokenVault to retrieve the complete configuration of a single token vault, including the encryption method, associated RAM role ARN, creation time, and update time.
//
// @param request - GetTokenVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenVaultResponse
func (client *Client) GetTokenVaultWithOptions(request *GetTokenVaultRequest, runtime *dara.RuntimeOptions) (_result *GetTokenVaultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the detailed configuration of a specified token vault.
//
// Description:
//
// Calls GetTokenVault to retrieve the complete configuration of a single token vault, including the encryption method, associated RAM role ARN, creation time, and update time.
//
// @param request - GetTokenVaultRequest
//
// @return GetTokenVaultResponse
func (client *Client) GetTokenVault(request *GetTokenVaultRequest) (_result *GetTokenVaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTokenVaultResponse{}
	_body, _err := client.GetTokenVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a user.
//
// Description:
//
// Returns the complete information of a specified user, including the user ID, username, display name, description, and creation and update times.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	if !dara.IsNil(request.UserPoolId) {
		body["UserPoolId"] = request.UserPoolId
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a user.
//
// Description:
//
// Returns the complete information of a specified user, including the user ID, username, display name, description, and creation and update times.
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
// Retrieves a user pool.
//
// Description:
//
// Returns the complete configuration information of a specified user pool, including the user pool ID, name, description, creation time, and update time.
//
// @param request - GetUserPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserPoolResponse
func (client *Client) GetUserPoolWithOptions(request *GetUserPoolRequest, runtime *dara.RuntimeOptions) (_result *GetUserPoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a user pool.
//
// Description:
//
// Returns the complete configuration information of a specified user pool, including the user pool ID, name, description, creation time, and update time.
//
// @param request - GetUserPoolRequest
//
// @return GetUserPoolResponse
func (client *Client) GetUserPool(request *GetUserPoolRequest) (_result *GetUserPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserPoolResponse{}
	_body, _err := client.GetUserPoolWithOptions(request, runtime)
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
// Description:
//
// Returns the complete configuration information of a client, including the client ID, name, redirect URIs, token validity periods, SecretRequired and EnforcePKCE settings, and default client scopes.
//
// @param request - GetUserPoolClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserPoolClientResponse
func (client *Client) GetUserPoolClientWithOptions(request *GetUserPoolClientRequest, runtime *dara.RuntimeOptions) (_result *GetUserPoolClientResponse, _err error) {
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
// Description:
//
// Returns the complete configuration information of a client, including the client ID, name, redirect URIs, token validity periods, SecretRequired and EnforcePKCE settings, and default client scopes.
//
// @param request - GetUserPoolClientRequest
//
// @return GetUserPoolClientResponse
func (client *Client) GetUserPoolClient(request *GetUserPoolClientRequest) (_result *GetUserPoolClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserPoolClientResponse{}
	_body, _err := client.GetUserPoolClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a synchronization task in a user pool.
//
// Description:
//
// Queries the execution details of a single synchronization task in a specified user pool, including the task status, start and end times, synchronization result statistics (the number of created, updated, deleted, and unchanged users), and error messages.
//
// @param request - GetUserPoolSyncJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserPoolSyncJobResponse
func (client *Client) GetUserPoolSyncJobWithOptions(request *GetUserPoolSyncJobRequest, runtime *dara.RuntimeOptions) (_result *GetUserPoolSyncJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SynchronizationJobId) {
		body["SynchronizationJobId"] = request.SynchronizationJobId
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserPoolSyncJob"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserPoolSyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a synchronization task in a user pool.
//
// Description:
//
// Queries the execution details of a single synchronization task in a specified user pool, including the task status, start and end times, synchronization result statistics (the number of created, updated, deleted, and unchanged users), and error messages.
//
// @param request - GetUserPoolSyncJobRequest
//
// @return GetUserPoolSyncJobResponse
func (client *Client) GetUserPoolSyncJob(request *GetUserPoolSyncJobRequest) (_result *GetUserPoolSyncJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserPoolSyncJobResponse{}
	_body, _err := client.GetUserPoolSyncJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed configuration of a workload identity by name.
//
// @param request - GetWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkloadIdentityResponse
func (client *Client) GetWorkloadIdentityWithOptions(request *GetWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadIdentityResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed configuration of a workload identity by name.
//
// @param request - GetWorkloadIdentityRequest
//
// @return GetWorkloadIdentityResponse
func (client *Client) GetWorkloadIdentity(request *GetWorkloadIdentityRequest) (_result *GetWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkloadIdentityResponse{}
	_body, _err := client.GetWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all API key credential providers in the account.
//
// @param request - ListAPIKeyCredentialProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAPIKeyCredentialProvidersResponse
func (client *Client) ListAPIKeyCredentialProvidersWithOptions(request *ListAPIKeyCredentialProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListAPIKeyCredentialProvidersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all API key credential providers in the account.
//
// @param request - ListAPIKeyCredentialProvidersRequest
//
// @return ListAPIKeyCredentialProvidersResponse
func (client *Client) ListAPIKeyCredentialProviders(request *ListAPIKeyCredentialProvidersRequest) (_result *ListAPIKeyCredentialProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAPIKeyCredentialProvidersResponse{}
	_body, _err := client.ListAPIKeyCredentialProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists identity providers.
//
// Description:
//
// Returns summary information about all secrets of a client, including the client key ID and creation time. For security reasons, the plaintext secret (ClientSecretValue) is not returned in the list.
//
// @param request - ListClientSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientSecretsResponse
func (client *Client) ListClientSecretsWithOptions(request *ListClientSecretsRequest, runtime *dara.RuntimeOptions) (_result *ListClientSecretsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists identity providers.
//
// Description:
//
// Returns summary information about all secrets of a client, including the client key ID and creation time. For security reasons, the plaintext secret (ClientSecretValue) is not returned in the list.
//
// @param request - ListClientSecretsRequest
//
// @return ListClientSecretsResponse
func (client *Client) ListClientSecrets(request *ListClientSecretsRequest) (_result *ListClientSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClientSecretsResponse{}
	_body, _err := client.ListClientSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询身份源部门树（按层懒加载，供控制台勾选同步范围）
//
// Description:
//
// 用户池需已配置相应身份源，且应用具有目标部门的通讯录访问权限。本接口使用已保存的身份源配置，无需传入应用密钥。
//
// 本接口仅返回指定父部门的直属子部门，不包含父部门本身，不递归返回整棵部门树。未传父部门时，钉钉默认查询根部门1下的直属子部门，飞书默认查询根部门0下的直属子部门。部门浏览受应用授权可见范围限制，不受当前保存的同步范围限制。
//
// 本接口不修改同步范围，不触发人员同步。
//
// @param request - ListIdentityProviderDepartmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProviderDepartmentsResponse
func (client *Client) ListIdentityProviderDepartmentsWithOptions(request *ListIdentityProviderDepartmentsRequest, runtime *dara.RuntimeOptions) (_result *ListIdentityProviderDepartmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderType) {
		body["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentDepartmentId) {
		body["ParentDepartmentId"] = request.ParentDepartmentId
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentityProviderDepartments"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityProviderDepartmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询身份源部门树（按层懒加载，供控制台勾选同步范围）
//
// Description:
//
// 用户池需已配置相应身份源，且应用具有目标部门的通讯录访问权限。本接口使用已保存的身份源配置，无需传入应用密钥。
//
// 本接口仅返回指定父部门的直属子部门，不包含父部门本身，不递归返回整棵部门树。未传父部门时，钉钉默认查询根部门1下的直属子部门，飞书默认查询根部门0下的直属子部门。部门浏览受应用授权可见范围限制，不受当前保存的同步范围限制。
//
// 本接口不修改同步范围，不触发人员同步。
//
// @param request - ListIdentityProviderDepartmentsRequest
//
// @return ListIdentityProviderDepartmentsResponse
func (client *Client) ListIdentityProviderDepartments(request *ListIdentityProviderDepartmentsRequest) (_result *ListIdentityProviderDepartmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIdentityProviderDepartmentsResponse{}
	_body, _err := client.ListIdentityProviderDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListIdentityProvidersWithOptions(request *ListIdentityProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListIdentityProvidersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListIdentityProvidersResponse
func (client *Client) ListIdentityProviders(request *ListIdentityProvidersRequest) (_result *ListIdentityProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIdentityProvidersResponse{}
	_body, _err := client.ListIdentityProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all OAuth 2.0 credential providers in the account.
//
// @param request - ListOAuth2CredentialProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOAuth2CredentialProvidersResponse
func (client *Client) ListOAuth2CredentialProvidersWithOptions(request *ListOAuth2CredentialProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListOAuth2CredentialProvidersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all OAuth 2.0 credential providers in the account.
//
// @param request - ListOAuth2CredentialProvidersRequest
//
// @return ListOAuth2CredentialProvidersResponse
func (client *Client) ListOAuth2CredentialProviders(request *ListOAuth2CredentialProvidersRequest) (_result *ListOAuth2CredentialProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOAuth2CredentialProvidersResponse{}
	_body, _err := client.ListOAuth2CredentialProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPoliciesResponse
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists gateway policy configurations.
//
// Description:
//
// Calls ListPolicySetAttachedGateways to query the gateways to which a specific policy set is currently attached. Call this operation before deleting a policy set to verify that no residual attachments exist.
//
// @param request - ListPolicySetAttachedGatewaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicySetAttachedGatewaysResponse
func (client *Client) ListPolicySetAttachedGatewaysWithOptions(request *ListPolicySetAttachedGatewaysRequest, runtime *dara.RuntimeOptions) (_result *ListPolicySetAttachedGatewaysResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists gateway policy configurations.
//
// Description:
//
// Calls ListPolicySetAttachedGateways to query the gateways to which a specific policy set is currently attached. Call this operation before deleting a policy set to verify that no residual attachments exist.
//
// @param request - ListPolicySetAttachedGatewaysRequest
//
// @return ListPolicySetAttachedGatewaysResponse
func (client *Client) ListPolicySetAttachedGateways(request *ListPolicySetAttachedGatewaysRequest) (_result *ListPolicySetAttachedGatewaysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicySetAttachedGatewaysResponse{}
	_body, _err := client.ListPolicySetAttachedGatewaysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPolicySetsWithOptions(request *ListPolicySetsRequest, runtime *dara.RuntimeOptions) (_result *ListPolicySetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPolicySetsResponse
func (client *Client) ListPolicySets(request *ListPolicySetsRequest) (_result *ListPolicySetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicySetsResponse{}
	_body, _err := client.ListPolicySetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRoleAssignmentsWithOptions(request *ListRoleAssignmentsRequest, runtime *dara.RuntimeOptions) (_result *ListRoleAssignmentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRoleAssignmentsResponse
func (client *Client) ListRoleAssignments(request *ListRoleAssignmentsRequest) (_result *ListRoleAssignmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRoleAssignmentsResponse{}
	_body, _err := client.ListRoleAssignmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists roles.
//
// Description:
//
// Returns summary information about all roles in a user pool, including the default role Default and custom roles Manual. The information includes the role ID, name, description, creation time, update time, and role type.
//
// @param request - ListRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithOptions(request *ListRolesRequest, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists roles.
//
// Description:
//
// Returns summary information about all roles in a user pool, including the default role Default and custom roles Manual. The information includes the role ID, name, description, creation time, update time, and role type.
//
// @param request - ListRolesRequest
//
// @return ListRolesResponse
func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists identity providers.
//
// Description:
//
// Returns a list of all X.509 certificates under the SAML IdP configuration of the user pool to help administrators manage the certificate lifecycle and perform certificate rotation.
//
// @param request - ListSAMLIdentityProviderCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSAMLIdentityProviderCertificatesResponse
func (client *Client) ListSAMLIdentityProviderCertificatesWithOptions(request *ListSAMLIdentityProviderCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListSAMLIdentityProviderCertificatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists identity providers.
//
// Description:
//
// Returns a list of all X.509 certificates under the SAML IdP configuration of the user pool to help administrators manage the certificate lifecycle and perform certificate rotation.
//
// @param request - ListSAMLIdentityProviderCertificatesRequest
//
// @return ListSAMLIdentityProviderCertificatesResponse
func (client *Client) ListSAMLIdentityProviderCertificates(request *ListSAMLIdentityProviderCertificatesRequest) (_result *ListSAMLIdentityProviderCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSAMLIdentityProviderCertificatesResponse{}
	_body, _err := client.ListSAMLIdentityProviderCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTokenVaultsWithOptions(request *ListTokenVaultsRequest, runtime *dara.RuntimeOptions) (_result *ListTokenVaultsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTokenVaultsResponse
func (client *Client) ListTokenVaults(request *ListTokenVaultsRequest) (_result *ListTokenVaultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTokenVaultsResponse{}
	_body, _err := client.ListTokenVaultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists identity providers.
//
// Description:
//
// Returns summary information about all clients in a user pool, including the client ID, name, redirect URIs, creation time, and configurations such as PKCE, client secret, and token validity periods.
//
// @param request - ListUserPoolClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPoolClientsResponse
func (client *Client) ListUserPoolClientsWithOptions(request *ListUserPoolClientsRequest, runtime *dara.RuntimeOptions) (_result *ListUserPoolClientsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists identity providers.
//
// Description:
//
// Returns summary information about all clients in a user pool, including the client ID, name, redirect URIs, creation time, and configurations such as PKCE, client secret, and token validity periods.
//
// @param request - ListUserPoolClientsRequest
//
// @return ListUserPoolClientsResponse
func (client *Client) ListUserPoolClients(request *ListUserPoolClientsRequest) (_result *ListUserPoolClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserPoolClientsResponse{}
	_body, _err := client.ListUserPoolClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListUserPoolSyncJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPoolSyncJobsResponse
func (client *Client) ListUserPoolSyncJobsWithOptions(request *ListUserPoolSyncJobsRequest, runtime *dara.RuntimeOptions) (_result *ListUserPoolSyncJobsResponse, _err error) {
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
		Action:      dara.String("ListUserPoolSyncJobs"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPoolSyncJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListUserPoolSyncJobsRequest
//
// @return ListUserPoolSyncJobsResponse
func (client *Client) ListUserPoolSyncJobs(request *ListUserPoolSyncJobsRequest) (_result *ListUserPoolSyncJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserPoolSyncJobsResponse{}
	_body, _err := client.ListUserPoolSyncJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists identity providers.
//
// Description:
//
// Returns summary information about all user pools within the current account, including the user pool ID, name, description, and creation and update times. Supports pagination by using NextToken.
//
// @param request - ListUserPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPoolsResponse
func (client *Client) ListUserPoolsWithOptions(request *ListUserPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListUserPoolsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists identity providers.
//
// Description:
//
// Returns summary information about all user pools within the current account, including the user pool ID, name, description, and creation and update times. Supports pagination by using NextToken.
//
// @param request - ListUserPoolsRequest
//
// @return ListUserPoolsResponse
func (client *Client) ListUserPools(request *ListUserPoolsRequest) (_result *ListUserPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserPoolsResponse{}
	_body, _err := client.ListUserPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists users.
//
// Description:
//
// Returns the summary information of all users in a user pool, including the user ID, username, display name, description, creation time, and update time.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists users.
//
// Description:
//
// Returns the summary information of all users in a user pool, including the user ID, username, display name, description, creation time, and update time.
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
// Lists workload identities with pagination.
//
// @param request - ListWorkloadIdentitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkloadIdentitiesResponse
func (client *Client) ListWorkloadIdentitiesWithOptions(request *ListWorkloadIdentitiesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkloadIdentitiesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists workload identities with pagination.
//
// @param request - ListWorkloadIdentitiesRequest
//
// @return ListWorkloadIdentitiesResponse
func (client *Client) ListWorkloadIdentities(request *ListWorkloadIdentitiesRequest) (_result *ListWorkloadIdentitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkloadIdentitiesResponse{}
	_body, _err := client.ListWorkloadIdentitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers a user synchronization task in a user pool.
//
// Description:
//
// Creates and immediately executes a user synchronization task in a specified user pool. The system pulls user data from the specified identity provider and synchronizes the data to the user pool. The task runs asynchronously. The operation completes after the task ID is returned. You can call GetUserPoolSyncJob to query the execution progress and results. Only one synchronization task can run in a user pool at a time.
//
// @param request - RunUserPoolSyncJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunUserPoolSyncJobResponse
func (client *Client) RunUserPoolSyncJobWithOptions(request *RunUserPoolSyncJobRequest, runtime *dara.RuntimeOptions) (_result *RunUserPoolSyncJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderType) {
		body["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.MaxSyncUsers) {
		body["MaxSyncUsers"] = request.MaxSyncUsers
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunUserPoolSyncJob"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunUserPoolSyncJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a user synchronization task in a user pool.
//
// Description:
//
// Creates and immediately executes a user synchronization task in a specified user pool. The system pulls user data from the specified identity provider and synchronizes the data to the user pool. The task runs asynchronously. The operation completes after the task ID is returned. You can call GetUserPoolSyncJob to query the execution progress and results. Only one synchronization task can run in a user pool at a time.
//
// @param request - RunUserPoolSyncJobRequest
//
// @return RunUserPoolSyncJobResponse
func (client *Client) RunUserPoolSyncJob(request *RunUserPoolSyncJobRequest) (_result *RunUserPoolSyncJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunUserPoolSyncJobResponse{}
	_body, _err := client.RunUserPoolSyncJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workload identity.
//
// Description:
//
// This operation sets or updates the core configurations of a SAML identity provider (IdP), including the IdP identifier (EntityId), logon URL (LoginURL), X.509 certificates, single sign-on (SSO) status, and Just-In-Time (JIT) user provisioning and update settings. If you specify X509Certificates, all existing certificates in the user pool are replaced. To append certificates while retaining the existing ones, use AddSAMLIdentityProviderCertificate.
//
// @param tmpReq - SetSAMLIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSAMLIdentityProviderResponse
func (client *Client) SetSAMLIdentityProviderWithOptions(tmpReq *SetSAMLIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *SetSAMLIdentityProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workload identity.
//
// Description:
//
// This operation sets or updates the core configurations of a SAML identity provider (IdP), including the IdP identifier (EntityId), logon URL (LoginURL), X.509 certificates, single sign-on (SSO) status, and Just-In-Time (JIT) user provisioning and update settings. If you specify X509Certificates, all existing certificates in the user pool are replaced. To append certificates while retaining the existing ones, use AddSAMLIdentityProviderCertificate.
//
// @param request - SetSAMLIdentityProviderRequest
//
// @return SetSAMLIdentityProviderResponse
func (client *Client) SetSAMLIdentityProvider(request *SetSAMLIdentityProviderRequest) (_result *SetSAMLIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSAMLIdentityProviderResponse{}
	_body, _err := client.SetSAMLIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the SSO settings for a specified type of identity provider in a user pool.
//
// Description:
//
// Configures a specific type of identity provider for a specified user pool. This operation is used to configure SSO logon and address book synchronization for non-SAML IdPs such as DingTalk. A user pool does not support enabling multiple identity providers at the same time. You must disable other enabled IdPs before enabling a new IdP.
//
// @param request - SetSpecificIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSpecificIdentityProviderResponse
func (client *Client) SetSpecificIdentityProviderWithOptions(request *SetSpecificIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *SetSpecificIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IDPMetadata) {
		body["IDPMetadata"] = request.IDPMetadata
	}

	if !dara.IsNil(request.IdentityProviderType) {
		body["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.SSOStatus) {
		body["SSOStatus"] = request.SSOStatus
	}

	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSpecificIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSpecificIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the SSO settings for a specified type of identity provider in a user pool.
//
// Description:
//
// Configures a specific type of identity provider for a specified user pool. This operation is used to configure SSO logon and address book synchronization for non-SAML IdPs such as DingTalk. A user pool does not support enabling multiple identity providers at the same time. You must disable other enabled IdPs before enabling a new IdP.
//
// @param request - SetSpecificIdentityProviderRequest
//
// @return SetSpecificIdentityProviderResponse
func (client *Client) SetSpecificIdentityProvider(request *SetSpecificIdentityProviderRequest) (_result *SetSpecificIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSpecificIdentityProviderResponse{}
	_body, _err := client.SetSpecificIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the logon password of a user.
//
// Description:
//
// Sets or resets the logon password for a user in a specified user pool. Two methods are supported: specifying a password (Password) or generating a random password (GenerateRandomPassword). The two methods are mutually exclusive, and you must select exactly one. If you specify a password, it must meet the password policy requirements: a minimum length of 8 characters, containing at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters, and it must not contain the username. Passwords are managed independently of the user creation process and are not passed through CreateUser.
//
// @param request - SetUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserPasswordResponse
func (client *Client) SetUserPasswordWithOptions(request *SetUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *SetUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GenerateRandomPassword) {
		body["GenerateRandomPassword"] = request.GenerateRandomPassword
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

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
		Action:      dara.String("SetUserPassword"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the logon password of a user.
//
// Description:
//
// Sets or resets the logon password for a user in a specified user pool. Two methods are supported: specifying a password (Password) or generating a random password (GenerateRandomPassword). The two methods are mutually exclusive, and you must select exactly one. If you specify a password, it must meet the password policy requirements: a minimum length of 8 characters, containing at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters, and it must not contain the username. Passwords are managed independently of the user creation process and are not passed through CreateUser.
//
// @param request - SetUserPasswordRequest
//
// @return SetUserPasswordResponse
func (client *Client) SetUserPassword(request *SetUserPasswordRequest) (_result *SetUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetUserPasswordResponse{}
	_body, _err := client.SetUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an existing API key credential provider.
//
// @param request - UpdateAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAPIKeyCredentialProviderResponse
func (client *Client) UpdateAPIKeyCredentialProviderWithOptions(request *UpdateAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateAPIKeyCredentialProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an existing API key credential provider.
//
// @param request - UpdateAPIKeyCredentialProviderRequest
//
// @return UpdateAPIKeyCredentialProviderResponse
func (client *Client) UpdateAPIKeyCredentialProvider(request *UpdateAPIKeyCredentialProviderRequest) (_result *UpdateAPIKeyCredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAPIKeyCredentialProviderResponse{}
	_body, _err := client.UpdateAPIKeyCredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the gateway policy configuration.
//
// Description:
//
// Calls UpdateGatewayPolicyConfig to adjust the policy enforcement mode on a gateway. Common usage: Use LOG_ONLY for observation and verification during the initial rollout of a policy. After confirming that everything is correct, switch the gateway to ENFORCE to enforce interception. This operation does not modify the attached policy set itself. To replace the policy set, call AttachPolicySetToGateway.
//
// @param request - UpdateGatewayPolicyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayPolicyConfigResponse
func (client *Client) UpdateGatewayPolicyConfigWithOptions(request *UpdateGatewayPolicyConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayPolicyConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the gateway policy configuration.
//
// Description:
//
// Calls UpdateGatewayPolicyConfig to adjust the policy enforcement mode on a gateway. Common usage: Use LOG_ONLY for observation and verification during the initial rollout of a policy. After confirming that everything is correct, switch the gateway to ENFORCE to enforce interception. This operation does not modify the attached policy set itself. To replace the policy set, call AttachPolicySetToGateway.
//
// @param request - UpdateGatewayPolicyConfigRequest
//
// @return UpdateGatewayPolicyConfigResponse
func (client *Client) UpdateGatewayPolicyConfig(request *UpdateGatewayPolicyConfigRequest) (_result *UpdateGatewayPolicyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayPolicyConfigResponse{}
	_body, _err := client.UpdateGatewayPolicyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the specified identity provider.
//
// @param tmpReq - UpdateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProviderWithOptions(tmpReq *UpdateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateIdentityProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the specified identity provider.
//
// @param request - UpdateIdentityProviderRequest
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProvider(request *UpdateIdentityProviderRequest) (_result *UpdateIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIdentityProviderResponse{}
	_body, _err := client.UpdateIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the logon configuration of a user pool.
//
// Description:
//
// Updates the logon configuration of a specified user pool. Currently, only logon preferences, such as whether password logon is enabled, can be updated. The password policy cannot be modified. Specify only the fields to be updated. Unspecified fields remain unchanged.
//
// @param tmpReq - UpdateLoginPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoginPreferenceResponse
func (client *Client) UpdateLoginPreferenceWithOptions(tmpReq *UpdateLoginPreferenceRequest, runtime *dara.RuntimeOptions) (_result *UpdateLoginPreferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLoginPreferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedPostLogoutRedirectUris) {
		request.AllowedPostLogoutRedirectUrisShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedPostLogoutRedirectUris, dara.String("AllowedPostLogoutRedirectUris"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LoginPreference) {
		request.LoginPreferenceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LoginPreference, dara.String("LoginPreference"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowedPostLogoutRedirectUrisShrink) {
		query["AllowedPostLogoutRedirectUris"] = request.AllowedPostLogoutRedirectUrisShrink
	}

	if !dara.IsNil(request.LoginPreferenceShrink) {
		query["LoginPreference"] = request.LoginPreferenceShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UserPoolName) {
		body["UserPoolName"] = request.UserPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLoginPreference"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLoginPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the logon configuration of a user pool.
//
// Description:
//
// Updates the logon configuration of a specified user pool. Currently, only logon preferences, such as whether password logon is enabled, can be updated. The password policy cannot be modified. Specify only the fields to be updated. Unspecified fields remain unchanged.
//
// @param request - UpdateLoginPreferenceRequest
//
// @return UpdateLoginPreferenceResponse
func (client *Client) UpdateLoginPreference(request *UpdateLoginPreferenceRequest) (_result *UpdateLoginPreferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLoginPreferenceResponse{}
	_body, _err := client.UpdateLoginPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of an existing OAuth 2.0 credential provider.
//
// @param tmpReq - UpdateOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOAuth2CredentialProviderResponse
func (client *Client) UpdateOAuth2CredentialProviderWithOptions(tmpReq *UpdateOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateOAuth2CredentialProviderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an existing OAuth 2.0 credential provider.
//
// @param request - UpdateOAuth2CredentialProviderRequest
//
// @return UpdateOAuth2CredentialProviderResponse
func (client *Client) UpdateOAuth2CredentialProvider(request *UpdateOAuth2CredentialProviderRequest) (_result *UpdateOAuth2CredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOAuth2CredentialProviderResponse{}
	_body, _err := client.UpdateOAuth2CredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an access policy.
//
// Description:
//
// Calls the UpdatePolicy operation to update the description and definition of an existing policy. The updated definition takes effect immediately for all gateways to which the policy set is mounted.
//
// @param tmpReq - UpdatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicyWithOptions(tmpReq *UpdatePolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an access policy.
//
// Description:
//
// Calls the UpdatePolicy operation to update the description and definition of an existing policy. The updated definition takes effect immediately for all gateways to which the policy set is mounted.
//
// @param request - UpdatePolicyRequest
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicy(request *UpdatePolicyRequest) (_result *UpdatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.UpdatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an access policy set.
//
// Description:
//
// Calls UpdatePolicySet to update the description of an existing policy set. To modify specific policies in the policy set, use UpdatePolicy. To change the mount relationship, use AttachPolicySetToGateway or DetachPolicySetFromGateway.
//
// @param request - UpdatePolicySetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicySetResponse
func (client *Client) UpdatePolicySetWithOptions(request *UpdatePolicySetRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolicySetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an access policy set.
//
// Description:
//
// Calls UpdatePolicySet to update the description of an existing policy set. To modify specific policies in the policy set, use UpdatePolicy. To change the mount relationship, use AttachPolicySetToGateway or DetachPolicySetFromGateway.
//
// @param request - UpdatePolicySetRequest
//
// @return UpdatePolicySetResponse
func (client *Client) UpdatePolicySet(request *UpdatePolicySetRequest) (_result *UpdatePolicySetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePolicySetResponse{}
	_body, _err := client.UpdatePolicySetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a role.
//
// Description:
//
// Updates the Description field of a role. The RoleName field serves as a unique identifier and cannot be modified.
//
// @param request - UpdateRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a role.
//
// Description:
//
// Updates the Description field of a role. The RoleName field serves as a unique identifier and cannot be modified.
//
// @param request - UpdateRoleRequest
//
// @return UpdateRoleResponse
func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a token vault.
//
// Description:
//
// Updates the description and role configuration of an existing token vault. Note: The encryption configuration (KeyType or KmsKeyArn) of a token vault cannot be changed after it is created.
//
// @param request - UpdateTokenVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTokenVaultResponse
func (client *Client) UpdateTokenVaultWithOptions(request *UpdateTokenVaultRequest, runtime *dara.RuntimeOptions) (_result *UpdateTokenVaultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a token vault.
//
// Description:
//
// Updates the description and role configuration of an existing token vault. Note: The encryption configuration (KeyType or KmsKeyArn) of a token vault cannot be changed after it is created.
//
// @param request - UpdateTokenVaultRequest
//
// @return UpdateTokenVaultResponse
func (client *Client) UpdateTokenVault(request *UpdateTokenVaultRequest) (_result *UpdateTokenVaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTokenVaultResponse{}
	_body, _err := client.UpdateTokenVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a user in a user pool.
//
// Description:
//
// Updates the attribute information of a user in a specified user pool. Specify only the fields that you want to update. Unspecified fields remain unchanged. UserName is the user identifier and cannot be modified. You can update the display name, description, email address, and status. Email addresses must be unique within the user pool.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		body["Email"] = request.Email
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

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
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2025-09-01"),
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
// Updates a user in a user pool.
//
// Description:
//
// Updates the attribute information of a user in a specified user pool. Specify only the fields that you want to update. Unspecified fields remain unchanged. UserName is the user identifier and cannot be modified. You can update the display name, description, email address, and status. Email addresses must be unique within the user pool.
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

// Summary:
//
// 更新IdentityProvider
//
// @param request - UpdateUserPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPoolResponse
func (client *Client) UpdateUserPoolWithOptions(request *UpdateUserPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserPoolResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateUserPoolResponse
func (client *Client) UpdateUserPool(request *UpdateUserPoolRequest) (_result *UpdateUserPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserPoolResponse{}
	_body, _err := client.UpdateUserPoolWithOptions(request, runtime)
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
// Description:
//
// Updates the callback URLs (RedirectURIs), whether an application secret is required (SecretRequired), whether Proof Key for Code Exchange (PKCE) is enforced (EnforcePKCE), and the validity periods of the access token and refresh token. The client name (ClientName) serves as a unique identifier and cannot be modified.
//
// @param tmpReq - UpdateUserPoolClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPoolClientResponse
func (client *Client) UpdateUserPoolClientWithOptions(tmpReq *UpdateUserPoolClientRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserPoolClientResponse, _err error) {
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

	if !dara.IsNil(tmpReq.Scopes) {
		request.ScopesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scopes, dara.String("Scopes"), dara.String("json"))
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

	if !dara.IsNil(request.ScopesShrink) {
		body["Scopes"] = request.ScopesShrink
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
// Description:
//
// Updates the callback URLs (RedirectURIs), whether an application secret is required (SecretRequired), whether Proof Key for Code Exchange (PKCE) is enforced (EnforcePKCE), and the validity periods of the access token and refresh token. The client name (ClientName) serves as a unique identifier and cannot be modified.
//
// @param request - UpdateUserPoolClientRequest
//
// @return UpdateUserPoolClientResponse
func (client *Client) UpdateUserPoolClient(request *UpdateUserPoolClientRequest) (_result *UpdateUserPoolClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserPoolClientResponse{}
	_body, _err := client.UpdateUserPoolClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the specified workload identity.
//
// @param tmpReq - UpdateWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkloadIdentityResponse
func (client *Client) UpdateWorkloadIdentityWithOptions(tmpReq *UpdateWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkloadIdentityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedConsentCallbackURLs) {
		request.AllowedConsentCallbackURLsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedConsentCallbackURLs, dara.String("AllowedConsentCallbackURLs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AllowedResourceOAuth2ReturnURLs) {
		request.AllowedResourceOAuth2ReturnURLsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedResourceOAuth2ReturnURLs, dara.String("AllowedResourceOAuth2ReturnURLs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedConsentCallbackURLsShrink) {
		body["AllowedConsentCallbackURLs"] = request.AllowedConsentCallbackURLsShrink
	}

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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the specified workload identity.
//
// @param request - UpdateWorkloadIdentityRequest
//
// @return UpdateWorkloadIdentityResponse
func (client *Client) UpdateWorkloadIdentity(request *UpdateWorkloadIdentityRequest) (_result *UpdateWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkloadIdentityResponse{}
	_body, _err := client.UpdateWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
