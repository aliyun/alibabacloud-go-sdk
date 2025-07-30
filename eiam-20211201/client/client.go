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
	client.Endpoint, _err = client.GetEndpoint(dara.String("eiam"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds an Employee Identity and Access Management (EIAM) account to multiple EIAM organizations of Identity as a Service (IDaaS). If the account already exists in the organizational unit, the system directly returns a success response.
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnitsWithOptions(request *AddUserToOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitIds) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToOrganizationalUnits"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an Employee Identity and Access Management (EIAM) account to multiple EIAM organizations of Identity as a Service (IDaaS). If the account already exists in the organizational unit, the system directly returns a success response.
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnits(request *AddUserToOrganizationalUnitsRequest) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserToOrganizationalUnitsResponse{}
	_body, _err := client.AddUserToOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds Employee Identity and Access Management (EIAM) accounts to an EIAM group of Identity as a Service (IDaaS).
//
// @param request - AddUsersToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroupWithOptions(request *AddUsersToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUsersToGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUsersToGroup"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds Employee Identity and Access Management (EIAM) accounts to an EIAM group of Identity as a Service (IDaaS).
//
// @param request - AddUsersToGroupRequest
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroup(request *AddUsersToGroupRequest) (_result *AddUsersToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.AddUsersToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants the permissions to access an application to multiple account groups at a time in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - AuthorizeApplicationToGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToGroupsResponse
func (client *Client) AuthorizeApplicationToGroupsWithOptions(request *AuthorizeApplicationToGroupsRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeApplicationToGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeApplicationToGroups"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeApplicationToGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants the permissions to access an application to multiple account groups at a time in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - AuthorizeApplicationToGroupsRequest
//
// @return AuthorizeApplicationToGroupsResponse
func (client *Client) AuthorizeApplicationToGroups(request *AuthorizeApplicationToGroupsRequest) (_result *AuthorizeApplicationToGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthorizeApplicationToGroupsResponse{}
	_body, _err := client.AuthorizeApplicationToGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an application to multiple Employee Identity and Access Management (EIAM) organizations at a time.
//
// @param request - AuthorizeApplicationToOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToOrganizationalUnitsResponse
func (client *Client) AuthorizeApplicationToOrganizationalUnitsWithOptions(request *AuthorizeApplicationToOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeApplicationToOrganizationalUnitsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitIds) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeApplicationToOrganizationalUnits"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeApplicationToOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an application to multiple Employee Identity and Access Management (EIAM) organizations at a time.
//
// @param request - AuthorizeApplicationToOrganizationalUnitsRequest
//
// @return AuthorizeApplicationToOrganizationalUnitsResponse
func (client *Client) AuthorizeApplicationToOrganizationalUnits(request *AuthorizeApplicationToOrganizationalUnitsRequest) (_result *AuthorizeApplicationToOrganizationalUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthorizeApplicationToOrganizationalUnitsResponse{}
	_body, _err := client.AuthorizeApplicationToOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an application to multiple Employee Identity and Access Management (EIAM) accounts at a time.
//
// @param request - AuthorizeApplicationToUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToUsersResponse
func (client *Client) AuthorizeApplicationToUsersWithOptions(request *AuthorizeApplicationToUsersRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeApplicationToUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeApplicationToUsers"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeApplicationToUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an application to multiple Employee Identity and Access Management (EIAM) accounts at a time.
//
// @param request - AuthorizeApplicationToUsersRequest
//
// @return AuthorizeApplicationToUsersResponse
func (client *Client) AuthorizeApplicationToUsers(request *AuthorizeApplicationToUsersRequest) (_result *AuthorizeApplicationToUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthorizeApplicationToUsersResponse{}
	_body, _err := client.AuthorizeApplicationToUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an application to an Enterprise Identity Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// Description:
//
// IDaaS EIAM supports the following two standard single sign-on (SSO) protocols for adding applications: SAML 2.0 and OIDC. You can select an SSO protocol based on your business requirements when you add an application. You cannot change the SSO protocol that you selected after the application is added.
//
// @param request - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ApplicationSourceType) {
		query["ApplicationSourceType"] = request.ApplicationSourceType
	}

	if !dara.IsNil(request.ApplicationTemplateId) {
		query["ApplicationTemplateId"] = request.ApplicationTemplateId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogoUrl) {
		query["LogoUrl"] = request.LogoUrl
	}

	if !dara.IsNil(request.SsoType) {
		query["SsoType"] = request.SsoType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplication"),
		Version:     dara.String("2021-12-01"),
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
// Adds an application to an Enterprise Identity Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// Description:
//
// IDaaS EIAM supports the following two standard single sign-on (SSO) protocols for adding applications: SAML 2.0 and OIDC. You can select an SSO protocol based on your business requirements when you add an application. You cannot change the SSO protocol that you selected after the application is added.
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
// Creates a client key for an Employee Identity and Access Management (EIAM) application. An EIAM application can have up to two client keys.
//
// @param request - CreateApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationClientSecretResponse
func (client *Client) CreateApplicationClientSecretWithOptions(request *CreateApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationClientSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationClientSecret"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a client key for an Employee Identity and Access Management (EIAM) application. An EIAM application can have up to two client keys.
//
// @param request - CreateApplicationClientSecretRequest
//
// @return CreateApplicationClientSecretResponse
func (client *Client) CreateApplicationClientSecret(request *CreateApplicationClientSecretRequest) (_result *CreateApplicationClientSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationClientSecretResponse{}
	_body, _err := client.CreateApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Conditional Access Policy
//
// Description:
//
// # Create Conditional Access Policy
//
// @param request - CreateConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConditionalAccessPolicyResponse
func (client *Client) CreateConditionalAccessPolicyWithOptions(request *CreateConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateConditionalAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConditionalAccessPolicyName) {
		query["ConditionalAccessPolicyName"] = request.ConditionalAccessPolicyName
	}

	if !dara.IsNil(request.ConditionalAccessPolicyType) {
		query["ConditionalAccessPolicyType"] = request.ConditionalAccessPolicyType
	}

	if !dara.IsNil(request.ConditionsConfig) {
		query["ConditionsConfig"] = request.ConditionsConfig
	}

	if !dara.IsNil(request.DecisionConfig) {
		query["DecisionConfig"] = request.DecisionConfig
	}

	if !dara.IsNil(request.DecisionType) {
		query["DecisionType"] = request.DecisionType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EvaluateAt) {
		query["EvaluateAt"] = request.EvaluateAt
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConditionalAccessPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConditionalAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Conditional Access Policy
//
// Description:
//
// # Create Conditional Access Policy
//
// @param request - CreateConditionalAccessPolicyRequest
//
// @return CreateConditionalAccessPolicyResponse
func (client *Client) CreateConditionalAccessPolicy(request *CreateConditionalAccessPolicyRequest) (_result *CreateConditionalAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConditionalAccessPolicyResponse{}
	_body, _err := client.CreateConditionalAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom domain name for an Employee Identity and Access Management (EIAM) instance.
//
// @param request - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Filing) {
		query["Filing"] = request.Filing
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomain"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom domain name for an Employee Identity and Access Management (EIAM) instance.
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a proxy token for a domain name of an Employee Identity and Access Management (EIAM) instance.
//
// @param request - CreateDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainProxyTokenResponse
func (client *Client) CreateDomainProxyTokenWithOptions(request *CreateDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainProxyTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomainProxyToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a proxy token for a domain name of an Employee Identity and Access Management (EIAM) instance.
//
// @param request - CreateDomainProxyTokenRequest
//
// @return CreateDomainProxyTokenResponse
func (client *Client) CreateDomainProxyToken(request *CreateDomainProxyTokenRequest) (_result *CreateDomainProxyTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDomainProxyTokenResponse{}
	_body, _err := client.CreateDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupExternalId) {
		query["GroupExternalId"] = request.GroupExternalId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroup"),
		Version:     dara.String("2021-12-01"),
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
// Creates an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
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
// # Create Identity Provider
//
// @param request - CreateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityProviderResponse
func (client *Client) CreateIdentityProviderWithOptions(request *CreateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthnConfig) {
		query["AuthnConfig"] = request.AuthnConfig
	}

	if !dara.IsNil(request.AutoCreateUserConfig) {
		query["AutoCreateUserConfig"] = request.AutoCreateUserConfig
	}

	if !dara.IsNil(request.AutoUpdateUserConfig) {
		query["AutoUpdateUserConfig"] = request.AutoUpdateUserConfig
	}

	if !dara.IsNil(request.BindingConfig) {
		query["BindingConfig"] = request.BindingConfig
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DingtalkAppConfig) {
		query["DingtalkAppConfig"] = request.DingtalkAppConfig
	}

	if !dara.IsNil(request.IdentityProviderName) {
		query["IdentityProviderName"] = request.IdentityProviderName
	}

	if !dara.IsNil(request.IdentityProviderType) {
		query["IdentityProviderType"] = request.IdentityProviderType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LarkConfig) {
		query["LarkConfig"] = request.LarkConfig
	}

	if !dara.IsNil(request.LdapConfig) {
		query["LdapConfig"] = request.LdapConfig
	}

	if !dara.IsNil(request.LogoUrl) {
		query["LogoUrl"] = request.LogoUrl
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !dara.IsNil(request.OidcConfig) {
		query["OidcConfig"] = request.OidcConfig
	}

	if !dara.IsNil(request.UdPullConfig) {
		query["UdPullConfig"] = request.UdPullConfig
	}

	if !dara.IsNil(request.UdPushConfig) {
		query["UdPushConfig"] = request.UdPushConfig
	}

	if !dara.IsNil(request.WeComConfig) {
		query["WeComConfig"] = request.WeComConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdentityProvider"),
		Version:     dara.String("2021-12-01"),
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
// # Create Identity Provider
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
// Creates an instance based on which all capabilities of Identity as a Service (IDaaS) Enterprise Identity and Access Management (EIAM) are provided.
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance based on which all capabilities of Identity as a Service (IDaaS) Enterprise Identity and Access Management (EIAM) are provided.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dedicated endpoint.
//
// @param request - CreateNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkAccessEndpointResponse
func (client *Client) CreateNetworkAccessEndpointWithOptions(request *CreateNetworkAccessEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkAccessEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkAccessEndpointName) {
		query["NetworkAccessEndpointName"] = request.NetworkAccessEndpointName
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcRegionId) {
		query["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkAccessEndpoint"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkAccessEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dedicated endpoint.
//
// @param request - CreateNetworkAccessEndpointRequest
//
// @return CreateNetworkAccessEndpointResponse
func (client *Client) CreateNetworkAccessEndpoint(request *CreateNetworkAccessEndpointRequest) (_result *CreateNetworkAccessEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkAccessEndpointResponse{}
	_body, _err := client.CreateNetworkAccessEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - CreateOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnitWithOptions(request *CreateOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *CreateOrganizationalUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitExternalId) {
		query["OrganizationalUnitExternalId"] = request.OrganizationalUnitExternalId
	}

	if !dara.IsNil(request.OrganizationalUnitName) {
		query["OrganizationalUnitName"] = request.OrganizationalUnitName
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - CreateOrganizationalUnitRequest
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnit(request *CreateOrganizationalUnitRequest) (_result *CreateOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.CreateOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an account in an Identity as a Service (IDaaS) Enterprise Identity Access Management (EIAM) instance.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CustomFields) {
		query["CustomFields"] = request.CustomFields
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.EmailVerified) {
		query["EmailVerified"] = request.EmailVerified
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitIds) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordInitializationConfig) {
		query["PasswordInitializationConfig"] = request.PasswordInitializationConfig
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PhoneNumberVerified) {
		query["PhoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !dara.IsNil(request.PhoneRegion) {
		query["PhoneRegion"] = request.PhoneRegion
	}

	if !dara.IsNil(request.PrimaryOrganizationalUnitId) {
		query["PrimaryOrganizationalUnitId"] = request.PrimaryOrganizationalUnitId
	}

	if !dara.IsNil(request.UserExternalId) {
		query["UserExternalId"] = request.UserExternalId
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2021-12-01"),
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
// Creates an account in an Identity as a Service (IDaaS) Enterprise Identity Access Management (EIAM) instance.
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
// Deletes an Employee Identity and Access Management (EIAM) application.
//
// Description:
//
// Make sure that the EIAM application that you want to delete is not used before you delete the EIAM application. After you delete the EIAM application, all configurations are deleted and cannot be restored.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplication"),
		Version:     dara.String("2021-12-01"),
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
// Deletes an Employee Identity and Access Management (EIAM) application.
//
// Description:
//
// Make sure that the EIAM application that you want to delete is not used before you delete the EIAM application. After you delete the EIAM application, all configurations are deleted and cannot be restored.
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
// Deletes a client key for an Employee Identity and Access Management (EIAM) application.
//
// @param request - DeleteApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationClientSecretResponse
func (client *Client) DeleteApplicationClientSecretWithOptions(request *DeleteApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationClientSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationClientSecret"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a client key for an Employee Identity and Access Management (EIAM) application.
//
// @param request - DeleteApplicationClientSecretRequest
//
// @return DeleteApplicationClientSecretResponse
func (client *Client) DeleteApplicationClientSecret(request *DeleteApplicationClientSecretRequest) (_result *DeleteApplicationClientSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationClientSecretResponse{}
	_body, _err := client.DeleteApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Conditional Access Policy
//
// Description:
//
// When deleting a specified conditional access policy, please ensure that the policy is no longer in use. After deletion, all configuration data will be removed and cannot be recovered.
//
// @param request - DeleteConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConditionalAccessPolicyResponse
func (client *Client) DeleteConditionalAccessPolicyWithOptions(request *DeleteConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteConditionalAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConditionalAccessPolicyId) {
		query["ConditionalAccessPolicyId"] = request.ConditionalAccessPolicyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConditionalAccessPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConditionalAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Conditional Access Policy
//
// Description:
//
// When deleting a specified conditional access policy, please ensure that the policy is no longer in use. After deletion, all configuration data will be removed and cannot be recovered.
//
// @param request - DeleteConditionalAccessPolicyRequest
//
// @return DeleteConditionalAccessPolicyResponse
func (client *Client) DeleteConditionalAccessPolicy(request *DeleteConditionalAccessPolicyRequest) (_result *DeleteConditionalAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConditionalAccessPolicyResponse{}
	_body, _err := client.DeleteConditionalAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom domain name of an Employee Identity and Access Management (EIAM) instance. You cannot delete the initial domain name and default domain name of the instance.
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom domain name of an Employee Identity and Access Management (EIAM) instance. You cannot delete the initial domain name and default domain name of the instance.
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a proxy token for a domain name of an Employee Identity and Access Management (EIAM) instance. Only the proxy tokens in the disabled state can be deleted.
//
// @param request - DeleteDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainProxyTokenResponse
func (client *Client) DeleteDomainProxyTokenWithOptions(request *DeleteDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainProxyTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.DomainProxyTokenId) {
		query["DomainProxyTokenId"] = request.DomainProxyTokenId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainProxyToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a proxy token for a domain name of an Employee Identity and Access Management (EIAM) instance. Only the proxy tokens in the disabled state can be deleted.
//
// @param request - DeleteDomainProxyTokenRequest
//
// @return DeleteDomainProxyTokenResponse
func (client *Client) DeleteDomainProxyToken(request *DeleteDomainProxyTokenRequest) (_result *DeleteDomainProxyTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainProxyTokenResponse{}
	_body, _err := client.DeleteDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the information of an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroup"),
		Version:     dara.String("2021-12-01"),
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
// Deletes the information of an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
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
// # Delete identity provider
//
// @param request - DeleteIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdentityProviderResponse
func (client *Client) DeleteIdentityProviderWithOptions(request *DeleteIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteIdentityProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIdentityProvider"),
		Version:     dara.String("2021-12-01"),
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
// # Delete identity provider
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
// Deletes an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS) that you do not need.
//
// Description:
//
// Make sure that the instance to be deleted is no longer used. If the instance is deleted, all data related to the instance will be deleted.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS) that you do not need.
//
// Description:
//
// Make sure that the instance to be deleted is no longer used. If the instance is deleted, all data related to the instance will be deleted.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a network endpoint of a specific type.
//
// @param request - DeleteNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkAccessEndpointResponse
func (client *Client) DeleteNetworkAccessEndpointWithOptions(request *DeleteNetworkAccessEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkAccessEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkAccessEndpoint"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkAccessEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a network endpoint of a specific type.
//
// @param request - DeleteNetworkAccessEndpointRequest
//
// @return DeleteNetworkAccessEndpointResponse
func (client *Client) DeleteNetworkAccessEndpoint(request *DeleteNetworkAccessEndpointRequest) (_result *DeleteNetworkAccessEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkAccessEndpointResponse{}
	_body, _err := client.DeleteNetworkAccessEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). If the organization has EIAM accounts or child organizations, the delete operation fails.
//
// @param request - DeleteOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnitWithOptions(request *DeleteOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *DeleteOrganizationalUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). If the organization has EIAM accounts or child organizations, the delete operation fails.
//
// @param request - DeleteOrganizationalUnitRequest
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnit(request *DeleteOrganizationalUnitRequest) (_result *DeleteOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.DeleteOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete organizational unit information, forcibly deleting all accounts and sub-organizations beneath it
//
// @param request - DeleteOrganizationalUnitChildrenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOrganizationalUnitChildrenResponse
func (client *Client) DeleteOrganizationalUnitChildrenWithOptions(request *DeleteOrganizationalUnitChildrenRequest, runtime *dara.RuntimeOptions) (_result *DeleteOrganizationalUnitChildrenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOrganizationalUnitChildren"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOrganizationalUnitChildrenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete organizational unit information, forcibly deleting all accounts and sub-organizations beneath it
//
// @param request - DeleteOrganizationalUnitChildrenRequest
//
// @return DeleteOrganizationalUnitChildrenResponse
func (client *Client) DeleteOrganizationalUnitChildren(request *DeleteOrganizationalUnitChildrenRequest) (_result *DeleteOrganizationalUnitChildrenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOrganizationalUnitChildrenResponse{}
	_body, _err := client.DeleteOrganizationalUnitChildrenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS). The information related to the account is cleared.
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2021-12-01"),
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
// Deletes an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS). The information related to the account is cleared.
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
// Disables an enabled Employee Identity and Access Management (EIAM) application. All features of the EIAM application cannot be used if you disable the EIAM application.
//
// Description:
//
// All features of the EIAM application cannot be used if you disable the EIAM application, such as single sign-on (SSO) and account synchronization. Make sure that you acknowledge the risks of the delete operation.
//
// @param request - DisableApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationResponse
func (client *Client) DisableApplicationWithOptions(request *DisableApplicationRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplication"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an enabled Employee Identity and Access Management (EIAM) application. All features of the EIAM application cannot be used if you disable the EIAM application.
//
// Description:
//
// All features of the EIAM application cannot be used if you disable the EIAM application, such as single sign-on (SSO) and account synchronization. Make sure that you acknowledge the risks of the delete operation.
//
// @param request - DisableApplicationRequest
//
// @return DisableApplicationResponse
func (client *Client) DisableApplication(request *DisableApplicationRequest) (_result *DisableApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableApplicationResponse{}
	_body, _err := client.DisableApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - DisableApplicationApiInvokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationApiInvokeResponse
func (client *Client) DisableApplicationApiInvokeWithOptions(request *DisableApplicationApiInvokeRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationApiInvokeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplicationApiInvoke"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationApiInvokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - DisableApplicationApiInvokeRequest
//
// @return DisableApplicationApiInvokeResponse
func (client *Client) DisableApplicationApiInvoke(request *DisableApplicationApiInvokeRequest) (_result *DisableApplicationApiInvokeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableApplicationApiInvokeResponse{}
	_body, _err := client.DisableApplicationApiInvokeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a client key of an Employee Identity and Access Management (EIAM) application.
//
// @param request - DisableApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationClientSecretResponse
func (client *Client) DisableApplicationClientSecretWithOptions(request *DisableApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationClientSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplicationClientSecret"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a client key of an Employee Identity and Access Management (EIAM) application.
//
// @param request - DisableApplicationClientSecretRequest
//
// @return DisableApplicationClientSecretResponse
func (client *Client) DisableApplicationClientSecret(request *DisableApplicationClientSecretRequest) (_result *DisableApplicationClientSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableApplicationClientSecretResponse{}
	_body, _err := client.DisableApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - DisableApplicationProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationProvisioningResponse
func (client *Client) DisableApplicationProvisioningWithOptions(request *DisableApplicationProvisioningRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationProvisioningResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplicationProvisioning"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - DisableApplicationProvisioningRequest
//
// @return DisableApplicationProvisioningResponse
func (client *Client) DisableApplicationProvisioning(request *DisableApplicationProvisioningRequest) (_result *DisableApplicationProvisioningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableApplicationProvisioningResponse{}
	_body, _err := client.DisableApplicationProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the single sign-on (SSO) feature for an Employee Identity and Access Management (EIAM) application. This way, employees cannot log on to the application by using SSO.
//
// @param request - DisableApplicationSsoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationSsoResponse
func (client *Client) DisableApplicationSsoWithOptions(request *DisableApplicationSsoRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationSsoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplicationSso"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationSsoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the single sign-on (SSO) feature for an Employee Identity and Access Management (EIAM) application. This way, employees cannot log on to the application by using SSO.
//
// @param request - DisableApplicationSsoRequest
//
// @return DisableApplicationSsoResponse
func (client *Client) DisableApplicationSso(request *DisableApplicationSsoRequest) (_result *DisableApplicationSsoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableApplicationSsoResponse{}
	_body, _err := client.DisableApplicationSsoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Disable Conditional Access Policy
//
// Description:
//
// When changing a conditional access policy from an enabled state to a disabled state, the policy will no longer intercept. Please confirm that you are aware of the potential risks associated with this action.
//
// @param request - DisableConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableConditionalAccessPolicyResponse
func (client *Client) DisableConditionalAccessPolicyWithOptions(request *DisableConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *DisableConditionalAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConditionalAccessPolicyId) {
		query["ConditionalAccessPolicyId"] = request.ConditionalAccessPolicyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableConditionalAccessPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableConditionalAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Disable Conditional Access Policy
//
// Description:
//
// When changing a conditional access policy from an enabled state to a disabled state, the policy will no longer intercept. Please confirm that you are aware of the potential risks associated with this action.
//
// @param request - DisableConditionalAccessPolicyRequest
//
// @return DisableConditionalAccessPolicyResponse
func (client *Client) DisableConditionalAccessPolicy(request *DisableConditionalAccessPolicyRequest) (_result *DisableConditionalAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableConditionalAccessPolicyResponse{}
	_body, _err := client.DisableConditionalAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a proxy token for a domain name of an Employee Identity and Access Management (EIAM) instance. After the proxy token is disabled, the domain name may not be used as expected.
//
// @param request - DisableDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDomainProxyTokenResponse
func (client *Client) DisableDomainProxyTokenWithOptions(request *DisableDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *DisableDomainProxyTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.DomainProxyTokenId) {
		query["DomainProxyTokenId"] = request.DomainProxyTokenId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDomainProxyToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a proxy token for a domain name of an Employee Identity and Access Management (EIAM) instance. After the proxy token is disabled, the domain name may not be used as expected.
//
// @param request - DisableDomainProxyTokenRequest
//
// @return DisableDomainProxyTokenResponse
func (client *Client) DisableDomainProxyToken(request *DisableDomainProxyTokenRequest) (_result *DisableDomainProxyTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDomainProxyTokenResponse{}
	_body, _err := client.DisableDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Disable identity provider synchronization
//
// @param request - DisableIdentityProviderUdPullRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableIdentityProviderUdPullResponse
func (client *Client) DisableIdentityProviderUdPullWithOptions(request *DisableIdentityProviderUdPullRequest, runtime *dara.RuntimeOptions) (_result *DisableIdentityProviderUdPullResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableIdentityProviderUdPull"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableIdentityProviderUdPullResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Disable identity provider synchronization
//
// @param request - DisableIdentityProviderUdPullRequest
//
// @return DisableIdentityProviderUdPullResponse
func (client *Client) DisableIdentityProviderUdPull(request *DisableIdentityProviderUdPullRequest) (_result *DisableIdentityProviderUdPullResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableIdentityProviderUdPullResponse{}
	_body, _err := client.DisableIdentityProviderUdPullWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the feature of automatically redirecting the initial domain name to the default domain name for an Employee Identity and Access Management (EIAM) instance. After the feature is disabled, users who visit the portal page by using the initial domain name are not redirected to the default domain name.
//
// @param request - DisableInitDomainAutoRedirectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInitDomainAutoRedirectResponse
func (client *Client) DisableInitDomainAutoRedirectWithOptions(request *DisableInitDomainAutoRedirectRequest, runtime *dara.RuntimeOptions) (_result *DisableInitDomainAutoRedirectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableInitDomainAutoRedirect"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInitDomainAutoRedirectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the feature of automatically redirecting the initial domain name to the default domain name for an Employee Identity and Access Management (EIAM) instance. After the feature is disabled, users who visit the portal page by using the initial domain name are not redirected to the default domain name.
//
// @param request - DisableInitDomainAutoRedirectRequest
//
// @return DisableInitDomainAutoRedirectResponse
func (client *Client) DisableInitDomainAutoRedirect(request *DisableInitDomainAutoRedirectRequest) (_result *DisableInitDomainAutoRedirectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableInitDomainAutoRedirectResponse{}
	_body, _err := client.DisableInitDomainAutoRedirectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account. If the account is disabled, a success message is returned.
//
// @param request - DisableUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserResponse
func (client *Client) DisableUserWithOptions(request *DisableUserRequest, runtime *dara.RuntimeOptions) (_result *DisableUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account. If the account is disabled, a success message is returned.
//
// @param request - DisableUserRequest
//
// @return DisableUserResponse
func (client *Client) DisableUser(request *DisableUserRequest) (_result *DisableUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableUserResponse{}
	_body, _err := client.DisableUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a disabled Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationResponse
func (client *Client) EnableApplicationWithOptions(request *EnableApplicationRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplication"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a disabled Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationRequest
//
// @return EnableApplicationResponse
func (client *Client) EnableApplication(request *EnableApplicationRequest) (_result *EnableApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableApplicationResponse{}
	_body, _err := client.EnableApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationApiInvokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationApiInvokeResponse
func (client *Client) EnableApplicationApiInvokeWithOptions(request *EnableApplicationApiInvokeRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationApiInvokeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplicationApiInvoke"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationApiInvokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationApiInvokeRequest
//
// @return EnableApplicationApiInvokeResponse
func (client *Client) EnableApplicationApiInvoke(request *EnableApplicationApiInvokeRequest) (_result *EnableApplicationApiInvokeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableApplicationApiInvokeResponse{}
	_body, _err := client.EnableApplicationApiInvokeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the client key of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - EnableApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationClientSecretResponse
func (client *Client) EnableApplicationClientSecretWithOptions(request *EnableApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationClientSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplicationClientSecret"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the client key of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - EnableApplicationClientSecretRequest
//
// @return EnableApplicationClientSecretResponse
func (client *Client) EnableApplicationClientSecret(request *EnableApplicationClientSecretRequest) (_result *EnableApplicationClientSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableApplicationClientSecretResponse{}
	_body, _err := client.EnableApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - EnableApplicationProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationProvisioningResponse
func (client *Client) EnableApplicationProvisioningWithOptions(request *EnableApplicationProvisioningRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationProvisioningResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplicationProvisioning"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - EnableApplicationProvisioningRequest
//
// @return EnableApplicationProvisioningResponse
func (client *Client) EnableApplicationProvisioning(request *EnableApplicationProvisioningRequest) (_result *EnableApplicationProvisioningResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableApplicationProvisioningResponse{}
	_body, _err := client.EnableApplicationProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the single sign-on (SSO) feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationSsoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationSsoResponse
func (client *Client) EnableApplicationSsoWithOptions(request *EnableApplicationSsoRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationSsoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplicationSso"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationSsoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the single sign-on (SSO) feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationSsoRequest
//
// @return EnableApplicationSsoResponse
func (client *Client) EnableApplicationSso(request *EnableApplicationSsoRequest) (_result *EnableApplicationSsoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableApplicationSsoResponse{}
	_body, _err := client.EnableApplicationSsoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Enable Conditional Access Policy
//
// Description:
//
// When changing the status of a conditional access policy from enabled to disabled, the policy will no longer intercept. Please confirm that you are aware of the potential risks associated with this action.
//
// @param request - EnableConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableConditionalAccessPolicyResponse
func (client *Client) EnableConditionalAccessPolicyWithOptions(request *EnableConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *EnableConditionalAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConditionalAccessPolicyId) {
		query["ConditionalAccessPolicyId"] = request.ConditionalAccessPolicyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableConditionalAccessPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableConditionalAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enable Conditional Access Policy
//
// Description:
//
// When changing the status of a conditional access policy from enabled to disabled, the policy will no longer intercept. Please confirm that you are aware of the potential risks associated with this action.
//
// @param request - EnableConditionalAccessPolicyRequest
//
// @return EnableConditionalAccessPolicyResponse
func (client *Client) EnableConditionalAccessPolicy(request *EnableConditionalAccessPolicyRequest) (_result *EnableConditionalAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableConditionalAccessPolicyResponse{}
	_body, _err := client.EnableConditionalAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a proxy token for a domain name of an Employee Identity and Access Management (EIAM) instance. The proxy token is used to verify the security of the domain name.
//
// @param request - EnableDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDomainProxyTokenResponse
func (client *Client) EnableDomainProxyTokenWithOptions(request *EnableDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *EnableDomainProxyTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.DomainProxyTokenId) {
		query["DomainProxyTokenId"] = request.DomainProxyTokenId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDomainProxyToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a proxy token for a domain name of an Employee Identity and Access Management (EIAM) instance. The proxy token is used to verify the security of the domain name.
//
// @param request - EnableDomainProxyTokenRequest
//
// @return EnableDomainProxyTokenResponse
func (client *Client) EnableDomainProxyToken(request *EnableDomainProxyTokenRequest) (_result *EnableDomainProxyTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDomainProxyTokenResponse{}
	_body, _err := client.EnableDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enable identity provider synchronization.
//
// @param request - EnableIdentityProviderUdPullRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableIdentityProviderUdPullResponse
func (client *Client) EnableIdentityProviderUdPullWithOptions(request *EnableIdentityProviderUdPullRequest, runtime *dara.RuntimeOptions) (_result *EnableIdentityProviderUdPullResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableIdentityProviderUdPull"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableIdentityProviderUdPullResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enable identity provider synchronization.
//
// @param request - EnableIdentityProviderUdPullRequest
//
// @return EnableIdentityProviderUdPullResponse
func (client *Client) EnableIdentityProviderUdPull(request *EnableIdentityProviderUdPullRequest) (_result *EnableIdentityProviderUdPullResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableIdentityProviderUdPullResponse{}
	_body, _err := client.EnableIdentityProviderUdPullWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the feature of automatically redirecting the initial domain name to the default domain name for an Employee Identity and Access Management (EIAM) instance.
//
// @param request - EnableInitDomainAutoRedirectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInitDomainAutoRedirectResponse
func (client *Client) EnableInitDomainAutoRedirectWithOptions(request *EnableInitDomainAutoRedirectRequest, runtime *dara.RuntimeOptions) (_result *EnableInitDomainAutoRedirectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInitDomainAutoRedirect"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInitDomainAutoRedirectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the feature of automatically redirecting the initial domain name to the default domain name for an Employee Identity and Access Management (EIAM) instance.
//
// @param request - EnableInitDomainAutoRedirectRequest
//
// @return EnableInitDomainAutoRedirectResponse
func (client *Client) EnableInitDomainAutoRedirect(request *EnableInitDomainAutoRedirectRequest) (_result *EnableInitDomainAutoRedirectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableInitDomainAutoRedirectResponse{}
	_body, _err := client.EnableInitDomainAutoRedirectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS).
//
// @param request - EnableUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableUserResponse
func (client *Client) EnableUserWithOptions(request *EnableUserRequest, runtime *dara.RuntimeOptions) (_result *EnableUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS).
//
// @param request - EnableUserRequest
//
// @return EnableUserResponse
func (client *Client) EnableUser(request *EnableUserRequest) (_result *EnableUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableUserResponse{}
	_body, _err := client.EnableUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an Employee Identity and Access Management (EIAM) application.
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplication"),
		Version:     dara.String("2021-12-01"),
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
// Queries the details of an Employee Identity and Access Management (EIAM) application.
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
// Queries the permissions of the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - GetApplicationGrantScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationGrantScopeResponse
func (client *Client) GetApplicationGrantScopeWithOptions(request *GetApplicationGrantScopeRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationGrantScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationGrantScope"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationGrantScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - GetApplicationGrantScopeRequest
//
// @return GetApplicationGrantScopeResponse
func (client *Client) GetApplicationGrantScope(request *GetApplicationGrantScopeRequest) (_result *GetApplicationGrantScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationGrantScopeResponse{}
	_body, _err := client.GetApplicationGrantScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - GetApplicationProvisioningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningConfigResponse
func (client *Client) GetApplicationProvisioningConfigWithOptions(request *GetApplicationProvisioningConfigRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationProvisioningConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationProvisioningConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationProvisioningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - GetApplicationProvisioningConfigRequest
//
// @return GetApplicationProvisioningConfigResponse
func (client *Client) GetApplicationProvisioningConfig(request *GetApplicationProvisioningConfigRequest) (_result *GetApplicationProvisioningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationProvisioningConfigResponse{}
	_body, _err := client.GetApplicationProvisioningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the account synchronization scope of applications in Identity as a Service (IDaaS) Employee IAM (EIAM). This scope is the same as the scope within which developers can call the DeveloperAPI to query and manage accounts.
//
// @param request - GetApplicationProvisioningScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScopeWithOptions(request *GetApplicationProvisioningScopeRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationProvisioningScope"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the account synchronization scope of applications in Identity as a Service (IDaaS) Employee IAM (EIAM). This scope is the same as the scope within which developers can call the DeveloperAPI to query and manage accounts.
//
// @param request - GetApplicationProvisioningScopeRequest
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScope(request *GetApplicationProvisioningScopeRequest) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.GetApplicationProvisioningScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the single sign-on (SSO) configuration attributes of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetApplicationSsoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationSsoConfigResponse
func (client *Client) GetApplicationSsoConfigWithOptions(request *GetApplicationSsoConfigRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationSsoConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationSsoConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationSsoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the single sign-on (SSO) configuration attributes of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetApplicationSsoConfigRequest
//
// @return GetApplicationSsoConfigResponse
func (client *Client) GetApplicationSsoConfig(request *GetApplicationSsoConfigRequest) (_result *GetApplicationSsoConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationSsoConfigResponse{}
	_body, _err := client.GetApplicationSsoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Conditional Access Policy
//
// Description:
//
// # Query Conditional Access Policy
//
// @param request - GetConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConditionalAccessPolicyResponse
func (client *Client) GetConditionalAccessPolicyWithOptions(request *GetConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetConditionalAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConditionalAccessPolicyId) {
		query["ConditionalAccessPolicyId"] = request.ConditionalAccessPolicyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConditionalAccessPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConditionalAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Conditional Access Policy
//
// Description:
//
// # Query Conditional Access Policy
//
// @param request - GetConditionalAccessPolicyRequest
//
// @return GetConditionalAccessPolicyResponse
func (client *Client) GetConditionalAccessPolicy(request *GetConditionalAccessPolicyRequest) (_result *GetConditionalAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConditionalAccessPolicyResponse{}
	_body, _err := client.GetConditionalAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a domain name of an Employee Identity and Access Management (EIAM) instance.
//
// @param request - GetDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithOptions(request *GetDomainRequest, runtime *dara.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomain"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a domain name of an Employee Identity and Access Management (EIAM) instance.
//
// @param request - GetDomainRequest
//
// @return GetDomainResponse
func (client *Client) GetDomain(request *GetDomainRequest) (_result *GetDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain name system (DNS) challenge records of a domain name of an Employee Identity and Access Management (EIAM) instance. The generated records are used to verify the ownership of the domain name.
//
// @param request - GetDomainDnsChallengeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainDnsChallengeResponse
func (client *Client) GetDomainDnsChallengeWithOptions(request *GetDomainDnsChallengeRequest, runtime *dara.RuntimeOptions) (_result *GetDomainDnsChallengeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomainDnsChallenge"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainDnsChallengeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain name system (DNS) challenge records of a domain name of an Employee Identity and Access Management (EIAM) instance. The generated records are used to verify the ownership of the domain name.
//
// @param request - GetDomainDnsChallengeRequest
//
// @return GetDomainDnsChallengeResponse
func (client *Client) GetDomainDnsChallenge(request *GetDomainDnsChallengeRequest) (_result *GetDomainDnsChallengeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDomainDnsChallengeResponse{}
	_body, _err := client.GetDomainDnsChallengeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the forgot password configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetForgetPasswordConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetForgetPasswordConfigurationResponse
func (client *Client) GetForgetPasswordConfigurationWithOptions(request *GetForgetPasswordConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetForgetPasswordConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetForgetPasswordConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetForgetPasswordConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the forgot password configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetForgetPasswordConfigurationRequest
//
// @return GetForgetPasswordConfigurationResponse
func (client *Client) GetForgetPasswordConfiguration(request *GetForgetPasswordConfigurationRequest) (_result *GetForgetPasswordConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetForgetPasswordConfigurationResponse{}
	_body, _err := client.GetForgetPasswordConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - GetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroup"),
		Version:     dara.String("2021-12-01"),
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
// Queries the information of an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
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
// # Get identity provider
//
// @param request - GetIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderResponse
func (client *Client) GetIdentityProviderWithOptions(request *GetIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityProvider"),
		Version:     dara.String("2021-12-01"),
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
// # Get identity provider
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
// # Get IdP Inbound Synchronization Configuration Information
//
// @param request - GetIdentityProviderUdPullConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderUdPullConfigurationResponse
func (client *Client) GetIdentityProviderUdPullConfigurationWithOptions(request *GetIdentityProviderUdPullConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderUdPullConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityProviderUdPullConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderUdPullConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get IdP Inbound Synchronization Configuration Information
//
// @param request - GetIdentityProviderUdPullConfigurationRequest
//
// @return GetIdentityProviderUdPullConfigurationResponse
func (client *Client) GetIdentityProviderUdPullConfiguration(request *GetIdentityProviderUdPullConfigurationRequest) (_result *GetIdentityProviderUdPullConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIdentityProviderUdPullConfigurationResponse{}
	_body, _err := client.GetIdentityProviderUdPullConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the currently effective License information of the instance
//
// Description:
//
// Please ensure that your current instance is no longer in use. When the EIAM instance is deleted, all related data will be deleted.
//
// @param request - GetInstanceLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceLicenseResponse
func (client *Client) GetInstanceLicenseWithOptions(request *GetInstanceLicenseRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceLicenseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceLicense"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the currently effective License information of the instance
//
// Description:
//
// Please ensure that your current instance is no longer in use. When the EIAM instance is deleted, all related data will be deleted.
//
// @param request - GetInstanceLicenseRequest
//
// @return GetInstanceLicenseResponse
func (client *Client) GetInstanceLicense(request *GetInstanceLicenseRequest) (_result *GetInstanceLicenseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceLicenseResponse{}
	_body, _err := client.GetInstanceLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Network Endpoint Information
//
// @param request - GetNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkAccessEndpointResponse
func (client *Client) GetNetworkAccessEndpointWithOptions(request *GetNetworkAccessEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkAccessEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNetworkAccessEndpoint"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkAccessEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Network Endpoint Information
//
// @param request - GetNetworkAccessEndpointRequest
//
// @return GetNetworkAccessEndpointResponse
func (client *Client) GetNetworkAccessEndpoint(request *GetNetworkAccessEndpointRequest) (_result *GetNetworkAccessEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetworkAccessEndpointResponse{}
	_body, _err := client.GetNetworkAccessEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an organizational unit in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnitWithOptions(request *GetOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *GetOrganizationalUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an organizational unit in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetOrganizationalUnitRequest
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnit(request *GetOrganizationalUnitRequest) (_result *GetOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.GetOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password complexity configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordComplexityConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordComplexityConfigurationResponse
func (client *Client) GetPasswordComplexityConfigurationWithOptions(request *GetPasswordComplexityConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordComplexityConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPasswordComplexityConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPasswordComplexityConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password complexity configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordComplexityConfigurationRequest
//
// @return GetPasswordComplexityConfigurationResponse
func (client *Client) GetPasswordComplexityConfiguration(request *GetPasswordComplexityConfigurationRequest) (_result *GetPasswordComplexityConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPasswordComplexityConfigurationResponse{}
	_body, _err := client.GetPasswordComplexityConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password expiration configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordExpirationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordExpirationConfigurationResponse
func (client *Client) GetPasswordExpirationConfigurationWithOptions(request *GetPasswordExpirationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordExpirationConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPasswordExpirationConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPasswordExpirationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password expiration configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordExpirationConfigurationRequest
//
// @return GetPasswordExpirationConfigurationResponse
func (client *Client) GetPasswordExpirationConfiguration(request *GetPasswordExpirationConfigurationRequest) (_result *GetPasswordExpirationConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPasswordExpirationConfigurationResponse{}
	_body, _err := client.GetPasswordExpirationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password history configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordHistoryConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordHistoryConfigurationResponse
func (client *Client) GetPasswordHistoryConfigurationWithOptions(request *GetPasswordHistoryConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordHistoryConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPasswordHistoryConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPasswordHistoryConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password history configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordHistoryConfigurationRequest
//
// @return GetPasswordHistoryConfigurationResponse
func (client *Client) GetPasswordHistoryConfiguration(request *GetPasswordHistoryConfigurationRequest) (_result *GetPasswordHistoryConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPasswordHistoryConfigurationResponse{}
	_body, _err := client.GetPasswordHistoryConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password initialization configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordInitializationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordInitializationConfigurationResponse
func (client *Client) GetPasswordInitializationConfigurationWithOptions(request *GetPasswordInitializationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordInitializationConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPasswordInitializationConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPasswordInitializationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password initialization configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordInitializationConfigurationRequest
//
// @return GetPasswordInitializationConfigurationResponse
func (client *Client) GetPasswordInitializationConfiguration(request *GetPasswordInitializationConfigurationRequest) (_result *GetPasswordInitializationConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPasswordInitializationConfigurationResponse{}
	_body, _err := client.GetPasswordInitializationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the root organizational unit in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetRootOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRootOrganizationalUnitResponse
func (client *Client) GetRootOrganizationalUnitWithOptions(request *GetRootOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *GetRootOrganizationalUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRootOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRootOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the root organizational unit in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetRootOrganizationalUnitRequest
//
// @return GetRootOrganizationalUnitResponse
func (client *Client) GetRootOrganizationalUnit(request *GetRootOrganizationalUnitRequest) (_result *GetRootOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRootOrganizationalUnitResponse{}
	_body, _err := client.GetRootOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about a single synchronization job.
//
// @param request - GetSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSynchronizationJobResponse
func (client *Client) GetSynchronizationJobWithOptions(request *GetSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *GetSynchronizationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SynchronizationJobId) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSynchronizationJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about a single synchronization job.
//
// @param request - GetSynchronizationJobRequest
//
// @return GetSynchronizationJobResponse
func (client *Client) GetSynchronizationJob(request *GetSynchronizationJobRequest) (_result *GetSynchronizationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSynchronizationJobResponse{}
	_body, _err := client.GetSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an account in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2021-12-01"),
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
// Queries the details of an account in Identity as a Service (IDaaS) Employee IAM (EIAM).
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
// Queries all client keys of an Employee Identity and Access Management (EIAM) application. The returned key secret is not masked. If you want to query the key secret that is masked, call the ObtainApplicationClientSecret operation.
//
// @param request - ListApplicationClientSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationClientSecretsResponse
func (client *Client) ListApplicationClientSecretsWithOptions(request *ListApplicationClientSecretsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationClientSecretsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationClientSecrets"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationClientSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all client keys of an Employee Identity and Access Management (EIAM) application. The returned key secret is not masked. If you want to query the key secret that is masked, call the ObtainApplicationClientSecret operation.
//
// @param request - ListApplicationClientSecretsRequest
//
// @return ListApplicationClientSecretsResponse
func (client *Client) ListApplicationClientSecrets(request *ListApplicationClientSecretsRequest) (_result *ListApplicationClientSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationClientSecretsResponse{}
	_body, _err := client.ListApplicationClientSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or multiple Employee Identity and Access Management (EIAM) applications by page.
//
// @param request - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.AuthorizationType) {
		query["AuthorizationType"] = request.AuthorizationType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.M2MClientStatus) {
		query["M2MClientStatus"] = request.M2MClientStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceServerStatus) {
		query["ResourceServerStatus"] = request.ResourceServerStatus
	}

	if !dara.IsNil(request.SsoType) {
		query["SsoType"] = request.SsoType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplications"),
		Version:     dara.String("2021-12-01"),
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
// Queries the information about one or multiple Employee Identity and Access Management (EIAM) applications by page.
//
// @param request - ListApplicationsRequest
//
// @return ListApplicationsResponse
func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applications that an Employee Identity and Access Management (EIAM) organization can access. The return result includes the IDs of the applications. If you want to obtain the details of the applications, call the GetApplication operation.
//
// Description:
//
// You can only query the permissions that are directly granted to the EIAM organization by calling the ListApplicationsForOrganizationalUnit operation. You can filter applications by configuring the **ApplicationIds*	- parameter when you call this operation.
//
// @param request - ListApplicationsForOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForOrganizationalUnitResponse
func (client *Client) ListApplicationsForOrganizationalUnitWithOptions(request *ListApplicationsForOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForOrganizationalUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that an Employee Identity and Access Management (EIAM) organization can access. The return result includes the IDs of the applications. If you want to obtain the details of the applications, call the GetApplication operation.
//
// Description:
//
// You can only query the permissions that are directly granted to the EIAM organization by calling the ListApplicationsForOrganizationalUnit operation. You can filter applications by configuring the **ApplicationIds*	- parameter when you call this operation.
//
// @param request - ListApplicationsForOrganizationalUnitRequest
//
// @return ListApplicationsForOrganizationalUnitResponse
func (client *Client) ListApplicationsForOrganizationalUnit(request *ListApplicationsForOrganizationalUnitRequest) (_result *ListApplicationsForOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsForOrganizationalUnitResponse{}
	_body, _err := client.ListApplicationsForOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applications that an Employee Identity and Access Management (EIAM) account can access. The return result includes the IDs of the applications. If you want to obtain the details of the applications, call the GetApplication operation.
//
// @param request - ListApplicationsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForUserResponse
func (client *Client) ListApplicationsForUserWithOptions(request *ListApplicationsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryMode) {
		query["QueryMode"] = request.QueryMode
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that an Employee Identity and Access Management (EIAM) account can access. The return result includes the IDs of the applications. If you want to obtain the details of the applications, call the GetApplication operation.
//
// @param request - ListApplicationsForUserRequest
//
// @return ListApplicationsForUserResponse
func (client *Client) ListApplicationsForUser(request *ListApplicationsForUserRequest) (_result *ListApplicationsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsForUserResponse{}
	_body, _err := client.ListApplicationsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of Conditional Access Policies
//
// Description:
//
// # Paginated query for the list of conditional access policies
//
// @param request - ListConditionalAccessPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConditionalAccessPoliciesResponse
func (client *Client) ListConditionalAccessPoliciesWithOptions(request *ListConditionalAccessPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListConditionalAccessPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PreviousToken) {
		query["PreviousToken"] = request.PreviousToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConditionalAccessPolicies"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConditionalAccessPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Conditional Access Policies
//
// Description:
//
// # Paginated query for the list of conditional access policies
//
// @param request - ListConditionalAccessPoliciesRequest
//
// @return ListConditionalAccessPoliciesResponse
func (client *Client) ListConditionalAccessPolicies(request *ListConditionalAccessPoliciesRequest) (_result *ListConditionalAccessPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConditionalAccessPoliciesResponse{}
	_body, _err := client.ListConditionalAccessPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Conditional Access Policies Associated with Network Areas
//
// Description:
//
// # List Conditional Access Policies Associated with Network Zones
//
// @param request - ListConditionalAccessPoliciesForNetworkZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConditionalAccessPoliciesForNetworkZoneResponse
func (client *Client) ListConditionalAccessPoliciesForNetworkZoneWithOptions(request *ListConditionalAccessPoliciesForNetworkZoneRequest, runtime *dara.RuntimeOptions) (_result *ListConditionalAccessPoliciesForNetworkZoneResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkZoneId) {
		query["NetworkZoneId"] = request.NetworkZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConditionalAccessPoliciesForNetworkZone"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConditionalAccessPoliciesForNetworkZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List Conditional Access Policies Associated with Network Areas
//
// Description:
//
// # List Conditional Access Policies Associated with Network Zones
//
// @param request - ListConditionalAccessPoliciesForNetworkZoneRequest
//
// @return ListConditionalAccessPoliciesForNetworkZoneResponse
func (client *Client) ListConditionalAccessPoliciesForNetworkZone(request *ListConditionalAccessPoliciesForNetworkZoneRequest) (_result *ListConditionalAccessPoliciesForNetworkZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConditionalAccessPoliciesForNetworkZoneResponse{}
	_body, _err := client.ListConditionalAccessPoliciesForNetworkZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the proxy tokens of a domain name of an Employee Identity and Access Management (EIAM) instance.
//
// @param request - ListDomainProxyTokensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainProxyTokensResponse
func (client *Client) ListDomainProxyTokensWithOptions(request *ListDomainProxyTokensRequest, runtime *dara.RuntimeOptions) (_result *ListDomainProxyTokensResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomainProxyTokens"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainProxyTokensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proxy tokens of a domain name of an Employee Identity and Access Management (EIAM) instance.
//
// @param request - ListDomainProxyTokensRequest
//
// @return ListDomainProxyTokensResponse
func (client *Client) ListDomainProxyTokens(request *ListDomainProxyTokensRequest) (_result *ListDomainProxyTokensResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDomainProxyTokensResponse{}
	_body, _err := client.ListDomainProxyTokensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of domain names of an Employee Identity and Access Management (EIAM) instance. The list contains the initial domain name and custom domain names.
//
// @param request - ListDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomains"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of domain names of an Employee Identity and Access Management (EIAM) instance. The list contains the initial domain name and custom domain names.
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about Employee Identity and Access Management (EIAM) V1.0 instances or EIAM V2.0 instances.
//
// @param request - ListEiamInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEiamInstancesResponse
func (client *Client) ListEiamInstancesWithOptions(request *ListEiamInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListEiamInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceRegionId) {
		query["InstanceRegionId"] = request.InstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEiamInstances"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEiamInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Employee Identity and Access Management (EIAM) V1.0 instances or EIAM V2.0 instances.
//
// @param request - ListEiamInstancesRequest
//
// @return ListEiamInstancesResponse
func (client *Client) ListEiamInstances(request *ListEiamInstancesRequest) (_result *ListEiamInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEiamInstancesResponse{}
	_body, _err := client.ListEiamInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions in which Employee Identity and Access Management (EIAM) V1.0 instances or EIAM V2.0 instances reside.
//
// @param request - ListEiamRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEiamRegionsResponse
func (client *Client) ListEiamRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *ListEiamRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListEiamRegions"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEiamRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions in which Employee Identity and Access Management (EIAM) V1.0 instances or EIAM V2.0 instances reside.
//
// @return ListEiamRegionsResponse
func (client *Client) ListEiamRegions() (_result *ListEiamRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEiamRegionsResponse{}
	_body, _err := client.ListEiamRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of account groups in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupExternalId) {
		query["GroupExternalId"] = request.GroupExternalId
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupNameStartsWith) {
		query["GroupNameStartsWith"] = request.GroupNameStartsWith
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroups"),
		Version:     dara.String("2021-12-01"),
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
// Queries a list of account groups in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
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
// Queries a list of account groups to which the permissions to access an application are granted. The returned results contain the group IDs. You can call the GetGroup operation to query the information about an account group based on the group ID.
//
// @param request - ListGroupsForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForApplicationResponse
func (client *Client) ListGroupsForApplicationWithOptions(request *ListGroupsForApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsForApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupsForApplication"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsForApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of account groups to which the permissions to access an application are granted. The returned results contain the group IDs. You can call the GetGroup operation to query the information about an account group based on the group ID.
//
// @param request - ListGroupsForApplicationRequest
//
// @return ListGroupsForApplicationResponse
func (client *Client) ListGroupsForApplication(request *ListGroupsForApplicationRequest) (_result *ListGroupsForApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGroupsForApplicationResponse{}
	_body, _err := client.ListGroupsForApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of account groups to which an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS) belongs.
//
// @param request - ListGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithOptions(request *ListGroupsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupsForUser"),
		Version:     dara.String("2021-12-01"),
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
// Queries a list of account groups to which an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS) belongs.
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
// Query the list of identity providers.
//
// @param request - ListIdentityProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProvidersResponse
func (client *Client) ListIdentityProvidersWithOptions(request *ListIdentityProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListIdentityProvidersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentityProviders"),
		Version:     dara.String("2021-12-01"),
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
// Query the list of identity providers.
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
// Queries the information of one or more Enterprise Identity and Access Management (EIAM) instances of Identity as a Service (IDaaS).
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of one or more Enterprise Identity and Access Management (EIAM) instances of Identity as a Service (IDaaS).
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get a list of regions that support network access endpoints.
//
// @param request - ListNetworkAccessEndpointAvailableRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessEndpointAvailableRegionsResponse
func (client *Client) ListNetworkAccessEndpointAvailableRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *ListNetworkAccessEndpointAvailableRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetworkAccessEndpointAvailableRegions"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkAccessEndpointAvailableRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get a list of regions that support network access endpoints.
//
// @return ListNetworkAccessEndpointAvailableRegionsResponse
func (client *Client) ListNetworkAccessEndpointAvailableRegions() (_result *ListNetworkAccessEndpointAvailableRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNetworkAccessEndpointAvailableRegionsResponse{}
	_body, _err := client.ListNetworkAccessEndpointAvailableRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取支持NAE的可用区列表
//
// @param request - ListNetworkAccessEndpointAvailableZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessEndpointAvailableZonesResponse
func (client *Client) ListNetworkAccessEndpointAvailableZonesWithOptions(request *ListNetworkAccessEndpointAvailableZonesRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkAccessEndpointAvailableZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NaeRegionId) {
		query["NaeRegionId"] = request.NaeRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetworkAccessEndpointAvailableZones"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkAccessEndpointAvailableZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取支持NAE的可用区列表
//
// @param request - ListNetworkAccessEndpointAvailableZonesRequest
//
// @return ListNetworkAccessEndpointAvailableZonesResponse
func (client *Client) ListNetworkAccessEndpointAvailableZones(request *ListNetworkAccessEndpointAvailableZonesRequest) (_result *ListNetworkAccessEndpointAvailableZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNetworkAccessEndpointAvailableZonesResponse{}
	_body, _err := client.ListNetworkAccessEndpointAvailableZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询专属网络端点。
//
// @param request - ListNetworkAccessEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessEndpointsResponse
func (client *Client) ListNetworkAccessEndpointsWithOptions(request *ListNetworkAccessEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkAccessEndpointsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkAccessEndpointStatus) {
		query["NetworkAccessEndpointStatus"] = request.NetworkAccessEndpointStatus
	}

	if !dara.IsNil(request.NetworkAccessEndpointType) {
		query["NetworkAccessEndpointType"] = request.NetworkAccessEndpointType
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcRegionId) {
		query["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetworkAccessEndpoints"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkAccessEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询专属网络端点。
//
// @param request - ListNetworkAccessEndpointsRequest
//
// @return ListNetworkAccessEndpointsResponse
func (client *Client) ListNetworkAccessEndpoints(request *ListNetworkAccessEndpointsRequest) (_result *ListNetworkAccessEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNetworkAccessEndpointsResponse{}
	_body, _err := client.ListNetworkAccessEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List the access paths under a certain network access endpoint.
//
// @param request - ListNetworkAccessPathsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessPathsResponse
func (client *Client) ListNetworkAccessPathsWithOptions(request *ListNetworkAccessPathsRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkAccessPathsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetworkAccessPaths"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkAccessPathsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List the access paths under a certain network access endpoint.
//
// @param request - ListNetworkAccessPathsRequest
//
// @return ListNetworkAccessPathsResponse
func (client *Client) ListNetworkAccessPaths(request *ListNetworkAccessPathsRequest) (_result *ListNetworkAccessPathsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNetworkAccessPathsResponse{}
	_body, _err := client.ListNetworkAccessPathsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all parent organizations of an Employee Identity and Access Management (EIAM) organization.
//
// @param request - ListOrganizationalUnitParentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitParentsResponse
func (client *Client) ListOrganizationalUnitParentsWithOptions(request *ListOrganizationalUnitParentsRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitParentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationalUnitParents"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationalUnitParentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all parent organizations of an Employee Identity and Access Management (EIAM) organization.
//
// @param request - ListOrganizationalUnitParentsRequest
//
// @return ListOrganizationalUnitParentsResponse
func (client *Client) ListOrganizationalUnitParents(request *ListOrganizationalUnitParentsRequest) (_result *ListOrganizationalUnitParentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOrganizationalUnitParentsResponse{}
	_body, _err := client.ListOrganizationalUnitParentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about organizational units in Identity as a Service (IDaaS) Employee IAM (EIAM) by page.
//
// @param request - ListOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnitsWithOptions(request *ListOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitIds) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !dara.IsNil(request.OrganizationalUnitName) {
		query["OrganizationalUnitName"] = request.OrganizationalUnitName
	}

	if !dara.IsNil(request.OrganizationalUnitNameStartsWith) {
		query["OrganizationalUnitNameStartsWith"] = request.OrganizationalUnitNameStartsWith
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationalUnits"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about organizational units in Identity as a Service (IDaaS) Employee IAM (EIAM) by page.
//
// @param request - ListOrganizationalUnitsRequest
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnits(request *ListOrganizationalUnitsRequest) (_result *ListOrganizationalUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.ListOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the organizations that are allowed to access an Employee Identity and Access Management (EIAM) application by page. The return result includes the IDs of the organizations. If you want to obtain the details of the organizations, call the GetOrganizationalUnit operation.
//
// @param request - ListOrganizationalUnitsForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsForApplicationResponse
func (client *Client) ListOrganizationalUnitsForApplicationWithOptions(request *ListOrganizationalUnitsForApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitsForApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitIds) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationalUnitsForApplication"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationalUnitsForApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the organizations that are allowed to access an Employee Identity and Access Management (EIAM) application by page. The return result includes the IDs of the organizations. If you want to obtain the details of the organizations, call the GetOrganizationalUnit operation.
//
// @param request - ListOrganizationalUnitsForApplicationRequest
//
// @return ListOrganizationalUnitsForApplicationResponse
func (client *Client) ListOrganizationalUnitsForApplication(request *ListOrganizationalUnitsForApplicationRequest) (_result *ListOrganizationalUnitsForApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOrganizationalUnitsForApplicationResponse{}
	_body, _err := client.ListOrganizationalUnitsForApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the supported Alibaba Cloud regions.
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the supported Alibaba Cloud regions.
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询同步任务
//
// @param request - ListSynchronizationJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSynchronizationJobsResponse
func (client *Client) ListSynchronizationJobsWithOptions(request *ListSynchronizationJobsRequest, runtime *dara.RuntimeOptions) (_result *ListSynchronizationJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TargetIds) {
		query["TargetIds"] = request.TargetIds
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSynchronizationJobs"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSynchronizationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询同步任务
//
// @param request - ListSynchronizationJobsRequest
//
// @return ListSynchronizationJobsResponse
func (client *Client) ListSynchronizationJobs(request *ListSynchronizationJobsRequest) (_result *ListSynchronizationJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSynchronizationJobsResponse{}
	_body, _err := client.ListSynchronizationJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of accounts in Identity as a Service (IDaaS) Employee IAM (EIAM) by page.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisplayNameStartsWith) {
		query["DisplayNameStartsWith"] = request.DisplayNameStartsWith
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PhoneRegion) {
		query["PhoneRegion"] = request.PhoneRegion
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserExternalId) {
		query["UserExternalId"] = request.UserExternalId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	if !dara.IsNil(request.UserSourceId) {
		query["UserSourceId"] = request.UserSourceId
	}

	if !dara.IsNil(request.UserSourceType) {
		query["UserSourceType"] = request.UserSourceType
	}

	if !dara.IsNil(request.UsernameStartsWith) {
		query["UsernameStartsWith"] = request.UsernameStartsWith
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2021-12-01"),
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
// Queries the details of accounts in Identity as a Service (IDaaS) Employee IAM (EIAM) by page.
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
// Queries the accounts that are allowed to access an Employee Identity and Access Management (EIAM) application. The return results include the IDs of the accounts. If you need to obtain the details of the accounts, call the GetUser operation.
//
// @param request - ListUsersForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForApplicationResponse
func (client *Client) ListUsersForApplicationWithOptions(request *ListUsersForApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListUsersForApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsersForApplication"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersForApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accounts that are allowed to access an Employee Identity and Access Management (EIAM) application. The return results include the IDs of the accounts. If you need to obtain the details of the accounts, call the GetUser operation.
//
// @param request - ListUsersForApplicationRequest
//
// @return ListUsersForApplicationResponse
func (client *Client) ListUsersForApplication(request *ListUsersForApplicationRequest) (_result *ListUsersForApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersForApplicationResponse{}
	_body, _err := client.ListUsersForApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of accounts in an Employee Identity and Access Management (EIAM) group of Identity as a Service (IDaaS).
//
// @param request - ListUsersForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithOptions(request *ListUsersForGroupRequest, runtime *dara.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsersForGroup"),
		Version:     dara.String("2021-12-01"),
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
// Queries the information of accounts in an Employee Identity and Access Management (EIAM) group of Identity as a Service (IDaaS).
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
// Queries a client key of an Employee Identity and Access Management (EIAM) application. The returned key secret is masked. If you want to query the key secret that is not masked, call the ListApplicationClientSecrets operation.
//
// @param request - ObtainApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainApplicationClientSecretResponse
func (client *Client) ObtainApplicationClientSecretWithOptions(request *ObtainApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *ObtainApplicationClientSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecretId) {
		query["SecretId"] = request.SecretId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ObtainApplicationClientSecret"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ObtainApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a client key of an Employee Identity and Access Management (EIAM) application. The returned key secret is masked. If you want to query the key secret that is not masked, call the ListApplicationClientSecrets operation.
//
// @param request - ObtainApplicationClientSecretRequest
//
// @return ObtainApplicationClientSecretResponse
func (client *Client) ObtainApplicationClientSecret(request *ObtainApplicationClientSecretRequest) (_result *ObtainApplicationClientSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ObtainApplicationClientSecretResponse{}
	_body, _err := client.ObtainApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a proxy token of a domain name of an Employee Identity and Access Management (EIAM) instance.
//
// @param request - ObtainDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainDomainProxyTokenResponse
func (client *Client) ObtainDomainProxyTokenWithOptions(request *ObtainDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *ObtainDomainProxyTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.DomainProxyTokenId) {
		query["DomainProxyTokenId"] = request.DomainProxyTokenId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ObtainDomainProxyToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ObtainDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a proxy token of a domain name of an Employee Identity and Access Management (EIAM) instance.
//
// @param request - ObtainDomainProxyTokenRequest
//
// @return ObtainDomainProxyTokenResponse
func (client *Client) ObtainDomainProxyToken(request *ObtainDomainProxyTokenRequest) (_result *ObtainDomainProxyTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ObtainDomainProxyTokenResponse{}
	_body, _err := client.ObtainDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an Employee Identity and Access Management (EIAM) account from multiple EIAM organizations of Identity as a Service (IDaaS). You cannot remove an account from a primary organization.
//
// @param request - RemoveUserFromOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnitsWithOptions(request *RemoveUserFromOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitIds) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserFromOrganizationalUnits"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserFromOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an Employee Identity and Access Management (EIAM) account from multiple EIAM organizations of Identity as a Service (IDaaS). You cannot remove an account from a primary organization.
//
// @param request - RemoveUserFromOrganizationalUnitsRequest
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnits(request *RemoveUserFromOrganizationalUnitsRequest) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUserFromOrganizationalUnitsResponse{}
	_body, _err := client.RemoveUserFromOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes Employee Identity and Access Management (EIAM) accounts from an EIAM group of Identity as a Service (IDaaS).
//
// @param request - RemoveUsersFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroupWithOptions(request *RemoveUsersFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersFromGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsersFromGroup"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes Employee Identity and Access Management (EIAM) accounts from an EIAM group of Identity as a Service (IDaaS).
//
// @param request - RemoveUsersFromGroupRequest
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroup(request *RemoveUsersFromGroupRequest) (_result *RemoveUsersFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.RemoveUsersFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple account groups at a time in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - RevokeApplicationFromGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromGroupsResponse
func (client *Client) RevokeApplicationFromGroupsWithOptions(request *RevokeApplicationFromGroupsRequest, runtime *dara.RuntimeOptions) (_result *RevokeApplicationFromGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeApplicationFromGroups"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeApplicationFromGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple account groups at a time in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - RevokeApplicationFromGroupsRequest
//
// @return RevokeApplicationFromGroupsResponse
func (client *Client) RevokeApplicationFromGroups(request *RevokeApplicationFromGroupsRequest) (_result *RevokeApplicationFromGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeApplicationFromGroupsResponse{}
	_body, _err := client.RevokeApplicationFromGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple Employee Identity and Access Management (EIAM) organizations at a time.
//
// @param request - RevokeApplicationFromOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromOrganizationalUnitsResponse
func (client *Client) RevokeApplicationFromOrganizationalUnitsWithOptions(request *RevokeApplicationFromOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *RevokeApplicationFromOrganizationalUnitsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitIds) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeApplicationFromOrganizationalUnits"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeApplicationFromOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple Employee Identity and Access Management (EIAM) organizations at a time.
//
// @param request - RevokeApplicationFromOrganizationalUnitsRequest
//
// @return RevokeApplicationFromOrganizationalUnitsResponse
func (client *Client) RevokeApplicationFromOrganizationalUnits(request *RevokeApplicationFromOrganizationalUnitsRequest) (_result *RevokeApplicationFromOrganizationalUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeApplicationFromOrganizationalUnitsResponse{}
	_body, _err := client.RevokeApplicationFromOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple Employee Identity and Access Management (EIAM) accounts at a time.
//
// @param request - RevokeApplicationFromUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromUsersResponse
func (client *Client) RevokeApplicationFromUsersWithOptions(request *RevokeApplicationFromUsersRequest, runtime *dara.RuntimeOptions) (_result *RevokeApplicationFromUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeApplicationFromUsers"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeApplicationFromUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple Employee Identity and Access Management (EIAM) accounts at a time.
//
// @param request - RevokeApplicationFromUsersRequest
//
// @return RevokeApplicationFromUsersResponse
func (client *Client) RevokeApplicationFromUsers(request *RevokeApplicationFromUsersRequest) (_result *RevokeApplicationFromUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeApplicationFromUsersResponse{}
	_body, _err := client.RevokeApplicationFromUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a synchronization job and immediately runs the job.
//
// @param request - RunSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSynchronizationJobResponse
func (client *Client) RunSynchronizationJobWithOptions(request *RunSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *RunSynchronizationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PasswordInitialization) {
		query["PasswordInitialization"] = request.PasswordInitialization
	}

	if !dara.IsNil(request.SynchronizationScopeConfig) {
		query["SynchronizationScopeConfig"] = request.SynchronizationScopeConfig
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.UserIdentityTypes) {
		query["UserIdentityTypes"] = request.UserIdentityTypes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSynchronizationJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a synchronization job and immediately runs the job.
//
// @param request - RunSynchronizationJobRequest
//
// @return RunSynchronizationJobResponse
func (client *Client) RunSynchronizationJob(request *RunSynchronizationJobRequest) (_result *RunSynchronizationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunSynchronizationJobResponse{}
	_body, _err := client.RunSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the permissions of the Developer API feature of an Employee Identity and Access Management (EIAM) application.
//
// @param request - SetApplicationGrantScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationGrantScopeResponse
func (client *Client) SetApplicationGrantScopeWithOptions(request *SetApplicationGrantScopeRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationGrantScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.GrantScopes) {
		query["GrantScopes"] = request.GrantScopes
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApplicationGrantScope"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApplicationGrantScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the permissions of the Developer API feature of an Employee Identity and Access Management (EIAM) application.
//
// @param request - SetApplicationGrantScopeRequest
//
// @return SetApplicationGrantScopeResponse
func (client *Client) SetApplicationGrantScope(request *SetApplicationGrantScopeRequest) (_result *SetApplicationGrantScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetApplicationGrantScopeResponse{}
	_body, _err := client.SetApplicationGrantScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - SetApplicationProvisioningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationProvisioningConfigResponse
func (client *Client) SetApplicationProvisioningConfigWithOptions(request *SetApplicationProvisioningConfigRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationProvisioningConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.CallbackProvisioningConfig) {
		query["CallbackProvisioningConfig"] = request.CallbackProvisioningConfig
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProvisionPassword) {
		query["ProvisionPassword"] = request.ProvisionPassword
	}

	if !dara.IsNil(request.ProvisionProtocolType) {
		query["ProvisionProtocolType"] = request.ProvisionProtocolType
	}

	if !dara.IsNil(request.ScimProvisioningConfig) {
		query["ScimProvisioningConfig"] = request.ScimProvisioningConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApplicationProvisioningConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApplicationProvisioningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - SetApplicationProvisioningConfigRequest
//
// @return SetApplicationProvisioningConfigResponse
func (client *Client) SetApplicationProvisioningConfig(request *SetApplicationProvisioningConfigRequest) (_result *SetApplicationProvisioningConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetApplicationProvisioningConfigResponse{}
	_body, _err := client.SetApplicationProvisioningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the account synchronization scope of applications in Identity as a Service (IDaaS) Employee IAM (EIAM). This scope is the same as the scope within which developers can call the DeveloperAPI to query and manage accounts.
//
// @param request - SetApplicationProvisioningScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationProvisioningScopeResponse
func (client *Client) SetApplicationProvisioningScopeWithOptions(request *SetApplicationProvisioningScopeRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationProvisioningScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitIds) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApplicationProvisioningScope"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApplicationProvisioningScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the account synchronization scope of applications in Identity as a Service (IDaaS) Employee IAM (EIAM). This scope is the same as the scope within which developers can call the DeveloperAPI to query and manage accounts.
//
// @param request - SetApplicationProvisioningScopeRequest
//
// @return SetApplicationProvisioningScopeResponse
func (client *Client) SetApplicationProvisioningScope(request *SetApplicationProvisioningScopeRequest) (_result *SetApplicationProvisioningScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetApplicationProvisioningScopeResponse{}
	_body, _err := client.SetApplicationProvisioningScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies the single sign-on (SSO) configuration attributes of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// Description:
//
// In IDaaS EIAM, the application management feature supports multiple SSO protocols for applications, including SAML 2.0 and OIDC protocols. Each application supports only one protocol, and the protocol cannot be changed after the application is created. You can specify the SSO configuration attributes of an application based on the supported SSO protocol.
//
// @param request - SetApplicationSsoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationSsoConfigResponse
func (client *Client) SetApplicationSsoConfigWithOptions(request *SetApplicationSsoConfigRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationSsoConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InitLoginType) {
		query["InitLoginType"] = request.InitLoginType
	}

	if !dara.IsNil(request.InitLoginUrl) {
		query["InitLoginUrl"] = request.InitLoginUrl
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OidcSsoConfig) {
		query["OidcSsoConfig"] = request.OidcSsoConfig
	}

	if !dara.IsNil(request.SamlSsoConfig) {
		query["SamlSsoConfig"] = request.SamlSsoConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApplicationSsoConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApplicationSsoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the single sign-on (SSO) configuration attributes of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// Description:
//
// In IDaaS EIAM, the application management feature supports multiple SSO protocols for applications, including SAML 2.0 and OIDC protocols. Each application supports only one protocol, and the protocol cannot be changed after the application is created. You can specify the SSO configuration attributes of an application based on the supported SSO protocol.
//
// @param request - SetApplicationSsoConfigRequest
//
// @return SetApplicationSsoConfigResponse
func (client *Client) SetApplicationSsoConfig(request *SetApplicationSsoConfigRequest) (_result *SetApplicationSsoConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetApplicationSsoConfigResponse{}
	_body, _err := client.SetApplicationSsoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets a domain name of an Employee Identity and Access Management (EIAM) instance as the default domain name.
//
// @param request - SetDefaultDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomainWithOptions(request *SetDefaultDomainRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultDomain"),
		Version:     dara.String("2021-12-01"),
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
// Sets a domain name of an Employee Identity and Access Management (EIAM) instance as the default domain name.
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
// Configures a forgot password policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetForgetPasswordConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetForgetPasswordConfigurationResponse
func (client *Client) SetForgetPasswordConfigurationWithOptions(request *SetForgetPasswordConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetForgetPasswordConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationChannels) {
		query["AuthenticationChannels"] = request.AuthenticationChannels
	}

	if !dara.IsNil(request.ForgetPasswordStatus) {
		query["ForgetPasswordStatus"] = request.ForgetPasswordStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetForgetPasswordConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetForgetPasswordConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a forgot password policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetForgetPasswordConfigurationRequest
//
// @return SetForgetPasswordConfigurationResponse
func (client *Client) SetForgetPasswordConfiguration(request *SetForgetPasswordConfigurationRequest) (_result *SetForgetPasswordConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetForgetPasswordConfigurationResponse{}
	_body, _err := client.SetForgetPasswordConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update IdP synchronization configuration.
//
// @param request - SetIdentityProviderUdPullConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdentityProviderUdPullConfigurationResponse
func (client *Client) SetIdentityProviderUdPullConfigurationWithOptions(request *SetIdentityProviderUdPullConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetIdentityProviderUdPullConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupSyncStatus) {
		query["GroupSyncStatus"] = request.GroupSyncStatus
	}

	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.IncrementalCallbackStatus) {
		query["IncrementalCallbackStatus"] = request.IncrementalCallbackStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LdapUdPullConfig) {
		query["LdapUdPullConfig"] = request.LdapUdPullConfig
	}

	if !dara.IsNil(request.PeriodicSyncConfig) {
		query["PeriodicSyncConfig"] = request.PeriodicSyncConfig
	}

	if !dara.IsNil(request.PeriodicSyncStatus) {
		query["PeriodicSyncStatus"] = request.PeriodicSyncStatus
	}

	if !dara.IsNil(request.PullProtectedRule) {
		query["PullProtectedRule"] = request.PullProtectedRule
	}

	if !dara.IsNil(request.UdSyncScopeConfig) {
		query["UdSyncScopeConfig"] = request.UdSyncScopeConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIdentityProviderUdPullConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIdentityProviderUdPullConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update IdP synchronization configuration.
//
// @param request - SetIdentityProviderUdPullConfigurationRequest
//
// @return SetIdentityProviderUdPullConfigurationResponse
func (client *Client) SetIdentityProviderUdPullConfiguration(request *SetIdentityProviderUdPullConfigurationRequest) (_result *SetIdentityProviderUdPullConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetIdentityProviderUdPullConfigurationResponse{}
	_body, _err := client.SetIdentityProviderUdPullConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a password complexity policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordComplexityConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordComplexityConfigurationResponse
func (client *Client) SetPasswordComplexityConfigurationWithOptions(request *SetPasswordComplexityConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordComplexityConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PasswordComplexityRules) {
		query["PasswordComplexityRules"] = request.PasswordComplexityRules
	}

	if !dara.IsNil(request.PasswordMinLength) {
		query["PasswordMinLength"] = request.PasswordMinLength
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPasswordComplexityConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPasswordComplexityConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a password complexity policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordComplexityConfigurationRequest
//
// @return SetPasswordComplexityConfigurationResponse
func (client *Client) SetPasswordComplexityConfiguration(request *SetPasswordComplexityConfigurationRequest) (_result *SetPasswordComplexityConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPasswordComplexityConfigurationResponse{}
	_body, _err := client.SetPasswordComplexityConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a password expiration policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordExpirationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordExpirationConfigurationResponse
func (client *Client) SetPasswordExpirationConfigurationWithOptions(request *SetPasswordExpirationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordExpirationConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveAuthenticationSourceIds) {
		query["EffectiveAuthenticationSourceIds"] = request.EffectiveAuthenticationSourceIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PasswordExpirationAction) {
		query["PasswordExpirationAction"] = request.PasswordExpirationAction
	}

	if !dara.IsNil(request.PasswordExpirationNotificationChannels) {
		query["PasswordExpirationNotificationChannels"] = request.PasswordExpirationNotificationChannels
	}

	if !dara.IsNil(request.PasswordExpirationNotificationDuration) {
		query["PasswordExpirationNotificationDuration"] = request.PasswordExpirationNotificationDuration
	}

	if !dara.IsNil(request.PasswordExpirationNotificationStatus) {
		query["PasswordExpirationNotificationStatus"] = request.PasswordExpirationNotificationStatus
	}

	if !dara.IsNil(request.PasswordExpirationStatus) {
		query["PasswordExpirationStatus"] = request.PasswordExpirationStatus
	}

	if !dara.IsNil(request.PasswordForcedUpdateDuration) {
		query["PasswordForcedUpdateDuration"] = request.PasswordForcedUpdateDuration
	}

	if !dara.IsNil(request.PasswordValidMaxDay) {
		query["PasswordValidMaxDay"] = request.PasswordValidMaxDay
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPasswordExpirationConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPasswordExpirationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a password expiration policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordExpirationConfigurationRequest
//
// @return SetPasswordExpirationConfigurationResponse
func (client *Client) SetPasswordExpirationConfiguration(request *SetPasswordExpirationConfigurationRequest) (_result *SetPasswordExpirationConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPasswordExpirationConfigurationResponse{}
	_body, _err := client.SetPasswordExpirationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a password history policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordHistoryConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordHistoryConfigurationResponse
func (client *Client) SetPasswordHistoryConfigurationWithOptions(request *SetPasswordHistoryConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordHistoryConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PasswordHistoryMaxRetention) {
		query["PasswordHistoryMaxRetention"] = request.PasswordHistoryMaxRetention
	}

	if !dara.IsNil(request.PasswordHistoryStatus) {
		query["PasswordHistoryStatus"] = request.PasswordHistoryStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPasswordHistoryConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPasswordHistoryConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a password history policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordHistoryConfigurationRequest
//
// @return SetPasswordHistoryConfigurationResponse
func (client *Client) SetPasswordHistoryConfiguration(request *SetPasswordHistoryConfigurationRequest) (_result *SetPasswordHistoryConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPasswordHistoryConfigurationResponse{}
	_body, _err := client.SetPasswordHistoryConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the password initialization configurations for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordInitializationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordInitializationConfigurationResponse
func (client *Client) SetPasswordInitializationConfigurationWithOptions(request *SetPasswordInitializationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordInitializationConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PasswordForcedUpdateStatus) {
		query["PasswordForcedUpdateStatus"] = request.PasswordForcedUpdateStatus
	}

	if !dara.IsNil(request.PasswordInitializationNotificationChannels) {
		query["PasswordInitializationNotificationChannels"] = request.PasswordInitializationNotificationChannels
	}

	if !dara.IsNil(request.PasswordInitializationStatus) {
		query["PasswordInitializationStatus"] = request.PasswordInitializationStatus
	}

	if !dara.IsNil(request.PasswordInitializationType) {
		query["PasswordInitializationType"] = request.PasswordInitializationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPasswordInitializationConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPasswordInitializationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the password initialization configurations for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordInitializationConfigurationRequest
//
// @return SetPasswordInitializationConfigurationResponse
func (client *Client) SetPasswordInitializationConfiguration(request *SetPasswordInitializationConfigurationRequest) (_result *SetPasswordInitializationConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPasswordInitializationConfigurationResponse{}
	_body, _err := client.SetPasswordInitializationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the primary organizational unit to which an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account belongs. This account will be removed from the previous primary organizational unit and added to the new primary organization.
//
// @param request - SetUserPrimaryOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnitWithOptions(request *SetUserPrimaryOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetUserPrimaryOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the primary organizational unit to which an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account belongs. This account will be removed from the previous primary organizational unit and added to the new primary organization.
//
// @param request - SetUserPrimaryOrganizationalUnitRequest
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnit(request *SetUserPrimaryOrganizationalUnitRequest) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.SetUserPrimaryOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlocks an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS) that is locked.
//
// @param request - UnlockUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockUserResponse
func (client *Client) UnlockUserWithOptions(request *UnlockUserRequest, runtime *dara.RuntimeOptions) (_result *UnlockUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS) that is locked.
//
// @param request - UnlockUserRequest
//
// @return UnlockUserResponse
func (client *Client) UnlockUser(request *UnlockUserRequest) (_result *UnlockUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockUserResponse{}
	_body, _err := client.UnlockUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the authorization type of an Employee Identity and Access Management (EIAM) application.
//
// @param request - UpdateApplicationAuthorizationTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationAuthorizationTypeResponse
func (client *Client) UpdateApplicationAuthorizationTypeWithOptions(request *UpdateApplicationAuthorizationTypeRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationAuthorizationTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AuthorizationType) {
		query["AuthorizationType"] = request.AuthorizationType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationAuthorizationType"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationAuthorizationTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the authorization type of an Employee Identity and Access Management (EIAM) application.
//
// @param request - UpdateApplicationAuthorizationTypeRequest
//
// @return UpdateApplicationAuthorizationTypeResponse
func (client *Client) UpdateApplicationAuthorizationType(request *UpdateApplicationAuthorizationTypeRequest) (_result *UpdateApplicationAuthorizationTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationAuthorizationTypeResponse{}
	_body, _err := client.UpdateApplicationAuthorizationTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an Employee Identity and Access Management (EIAM) application.
//
// @param request - UpdateApplicationDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationDescriptionResponse
func (client *Client) UpdateApplicationDescriptionWithOptions(request *UpdateApplicationDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Employee Identity and Access Management (EIAM) application.
//
// @param request - UpdateApplicationDescriptionRequest
//
// @return UpdateApplicationDescriptionResponse
func (client *Client) UpdateApplicationDescription(request *UpdateApplicationDescriptionRequest) (_result *UpdateApplicationDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationDescriptionResponse{}
	_body, _err := client.UpdateApplicationDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Conditional Access Policy
//
// Description:
//
// # Update Conditional Access Policy
//
// @param request - UpdateConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConditionalAccessPolicyResponse
func (client *Client) UpdateConditionalAccessPolicyWithOptions(request *UpdateConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateConditionalAccessPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConditionalAccessPolicyId) {
		query["ConditionalAccessPolicyId"] = request.ConditionalAccessPolicyId
	}

	if !dara.IsNil(request.ConditionalAccessPolicyName) {
		query["ConditionalAccessPolicyName"] = request.ConditionalAccessPolicyName
	}

	if !dara.IsNil(request.ConditionsConfig) {
		query["ConditionsConfig"] = request.ConditionsConfig
	}

	if !dara.IsNil(request.DecisionConfig) {
		query["DecisionConfig"] = request.DecisionConfig
	}

	if !dara.IsNil(request.DecisionType) {
		query["DecisionType"] = request.DecisionType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConditionalAccessPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConditionalAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Conditional Access Policy
//
// Description:
//
// # Update Conditional Access Policy
//
// @param request - UpdateConditionalAccessPolicyRequest
//
// @return UpdateConditionalAccessPolicyResponse
func (client *Client) UpdateConditionalAccessPolicy(request *UpdateConditionalAccessPolicyRequest) (_result *UpdateConditionalAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConditionalAccessPolicyResponse{}
	_body, _err := client.UpdateConditionalAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Conditional Access Policy Description
//
// Description:
//
// # Update Conditional Access Policy Description
//
// @param request - UpdateConditionalAccessPolicyDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConditionalAccessPolicyDescriptionResponse
func (client *Client) UpdateConditionalAccessPolicyDescriptionWithOptions(request *UpdateConditionalAccessPolicyDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateConditionalAccessPolicyDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConditionalAccessPolicyId) {
		query["ConditionalAccessPolicyId"] = request.ConditionalAccessPolicyId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConditionalAccessPolicyDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConditionalAccessPolicyDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Conditional Access Policy Description
//
// Description:
//
// # Update Conditional Access Policy Description
//
// @param request - UpdateConditionalAccessPolicyDescriptionRequest
//
// @return UpdateConditionalAccessPolicyDescriptionResponse
func (client *Client) UpdateConditionalAccessPolicyDescription(request *UpdateConditionalAccessPolicyDescriptionRequest) (_result *UpdateConditionalAccessPolicyDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConditionalAccessPolicyDescriptionResponse{}
	_body, _err := client.UpdateConditionalAccessPolicyDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). If the information is empty, the information is not updated by default.
//
// @param request - UpdateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupExternalId) {
		query["GroupExternalId"] = request.GroupExternalId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroup"),
		Version:     dara.String("2021-12-01"),
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
// Updates the information about an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). If the information is empty, the information is not updated by default.
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
// Updates the description of an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account group.
//
// @param request - UpdateGroupDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupDescriptionResponse
func (client *Client) UpdateGroupDescriptionWithOptions(request *UpdateGroupDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateGroupDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroupDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGroupDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account group.
//
// @param request - UpdateGroupDescriptionRequest
//
// @return UpdateGroupDescriptionResponse
func (client *Client) UpdateGroupDescription(request *UpdateGroupDescriptionRequest) (_result *UpdateGroupDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGroupDescriptionResponse{}
	_body, _err := client.UpdateGroupDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新idp基础配置
//
// @param request - UpdateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProviderWithOptions(request *UpdateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateIdentityProviderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DingtalkAppConfig) {
		query["DingtalkAppConfig"] = request.DingtalkAppConfig
	}

	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.IdentityProviderName) {
		query["IdentityProviderName"] = request.IdentityProviderName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LarkConfig) {
		query["LarkConfig"] = request.LarkConfig
	}

	if !dara.IsNil(request.LdapConfig) {
		query["LdapConfig"] = request.LdapConfig
	}

	if !dara.IsNil(request.LogoUrl) {
		query["LogoUrl"] = request.LogoUrl
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !dara.IsNil(request.OidcConfig) {
		query["OidcConfig"] = request.OidcConfig
	}

	if !dara.IsNil(request.WeComConfig) {
		query["WeComConfig"] = request.WeComConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdentityProvider"),
		Version:     dara.String("2021-12-01"),
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
// 更新idp基础配置
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
// Modifies the description of an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - UpdateInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceDescriptionResponse
func (client *Client) UpdateInstanceDescriptionWithOptions(request *UpdateInstanceDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - UpdateInstanceDescriptionRequest
//
// @return UpdateInstanceDescriptionResponse
func (client *Client) UpdateInstanceDescription(request *UpdateInstanceDescriptionRequest) (_result *UpdateInstanceDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceDescriptionResponse{}
	_body, _err := client.UpdateInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新一个专属网络端点的名称。
//
// @param request - UpdateNetworkAccessEndpointNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkAccessEndpointNameResponse
func (client *Client) UpdateNetworkAccessEndpointNameWithOptions(request *UpdateNetworkAccessEndpointNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateNetworkAccessEndpointNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !dara.IsNil(request.NetworkAccessEndpointName) {
		query["NetworkAccessEndpointName"] = request.NetworkAccessEndpointName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNetworkAccessEndpointName"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNetworkAccessEndpointNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个专属网络端点的名称。
//
// @param request - UpdateNetworkAccessEndpointNameRequest
//
// @return UpdateNetworkAccessEndpointNameResponse
func (client *Client) UpdateNetworkAccessEndpointName(request *UpdateNetworkAccessEndpointNameRequest) (_result *UpdateNetworkAccessEndpointNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNetworkAccessEndpointNameResponse{}
	_body, _err := client.UpdateNetworkAccessEndpointNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about an Employee Identity and Access Management (EIAM) organization. The basic information about the organization is not updated by default if no parameter is specified.
//
// @param request - UpdateOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitResponse
func (client *Client) UpdateOrganizationalUnitWithOptions(request *UpdateOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationalUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !dara.IsNil(request.OrganizationalUnitName) {
		query["OrganizationalUnitName"] = request.OrganizationalUnitName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about an Employee Identity and Access Management (EIAM) organization. The basic information about the organization is not updated by default if no parameter is specified.
//
// @param request - UpdateOrganizationalUnitRequest
//
// @return UpdateOrganizationalUnitResponse
func (client *Client) UpdateOrganizationalUnit(request *UpdateOrganizationalUnitRequest) (_result *UpdateOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOrganizationalUnitResponse{}
	_body, _err := client.UpdateOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an Employee Identity and Access Management (EIAM) organization.
//
// @param request - UpdateOrganizationalUnitDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitDescriptionResponse
func (client *Client) UpdateOrganizationalUnitDescriptionWithOptions(request *UpdateOrganizationalUnitDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationalUnitDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOrganizationalUnitDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOrganizationalUnitDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Employee Identity and Access Management (EIAM) organization.
//
// @param request - UpdateOrganizationalUnitDescriptionRequest
//
// @return UpdateOrganizationalUnitDescriptionResponse
func (client *Client) UpdateOrganizationalUnitDescription(request *UpdateOrganizationalUnitDescriptionRequest) (_result *UpdateOrganizationalUnitDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOrganizationalUnitDescriptionResponse{}
	_body, _err := client.UpdateOrganizationalUnitDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the parent organization ID of an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). In this case, the organization is moved from a parent node to a new node.
//
// @param request - UpdateOrganizationalUnitParentIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitParentIdResponse
func (client *Client) UpdateOrganizationalUnitParentIdWithOptions(request *UpdateOrganizationalUnitParentIdRequest, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationalUnitParentIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOrganizationalUnitParentId"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOrganizationalUnitParentIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the parent organization ID of an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). In this case, the organization is moved from a parent node to a new node.
//
// @param request - UpdateOrganizationalUnitParentIdRequest
//
// @return UpdateOrganizationalUnitParentIdResponse
func (client *Client) UpdateOrganizationalUnitParentId(request *UpdateOrganizationalUnitParentIdRequest) (_result *UpdateOrganizationalUnitParentIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOrganizationalUnitParentIdResponse{}
	_body, _err := client.UpdateOrganizationalUnitParentIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS).
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomFields) {
		query["CustomFields"] = request.CustomFields
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.EmailVerified) {
		query["EmailVerified"] = request.EmailVerified
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PhoneNumberVerified) {
		query["PhoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !dara.IsNil(request.PhoneRegion) {
		query["PhoneRegion"] = request.PhoneRegion
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2021-12-01"),
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
// Updates the basic information about an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS).
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
// Modifies the description of an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account.
//
// @param request - UpdateUserDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDescriptionResponse
func (client *Client) UpdateUserDescriptionWithOptions(request *UpdateUserDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserDescriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account.
//
// @param request - UpdateUserDescriptionRequest
//
// @return UpdateUserDescriptionResponse
func (client *Client) UpdateUserDescription(request *UpdateUserDescriptionRequest) (_result *UpdateUserDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserDescriptionResponse{}
	_body, _err := client.UpdateUserDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the password information of an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS). The password must meet the requirements of the password policies that are configured in the IDaaS console.
//
// @param request - UpdateUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPasswordWithOptions(request *UpdateUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserPasswordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordForcedUpdateStatus) {
		query["PasswordForcedUpdateStatus"] = request.PasswordForcedUpdateStatus
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserNotificationChannels) {
		query["UserNotificationChannels"] = request.UserNotificationChannels
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserPassword"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the password information of an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS). The password must meet the requirements of the password policies that are configured in the IDaaS console.
//
// @param request - UpdateUserPasswordRequest
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPassword(request *UpdateUserPasswordRequest) (_result *UpdateUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserPasswordResponse{}
	_body, _err := client.UpdateUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
