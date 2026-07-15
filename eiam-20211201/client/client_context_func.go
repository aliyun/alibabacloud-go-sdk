// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds an application account for a specified employee under the current application.
//
// @param request - AddApplicationAccountToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddApplicationAccountToUserResponse
func (client *Client) AddApplicationAccountToUserWithContext(ctx context.Context, request *AddApplicationAccountToUserRequest, runtime *dara.RuntimeOptions) (_result *AddApplicationAccountToUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationUsername) {
		query["ApplicationUsername"] = request.ApplicationUsername
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
		Action:      dara.String("AddApplicationAccountToUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddApplicationAccountToUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an application to an authorization rule.
//
// @param request - AddApplicationToAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddApplicationToAuthorizationRuleResponse
func (client *Client) AddApplicationToAuthorizationRuleWithContext(ctx context.Context, request *AddApplicationToAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *AddApplicationToAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	if !dara.IsNil(request.ValidityType) {
		query["ValidityType"] = request.ValidityType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddApplicationToAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddApplicationToAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds terms to a brand.
//
// @param request - AddCustomPrivacyPoliciesToBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomPrivacyPoliciesToBrandResponse
func (client *Client) AddCustomPrivacyPoliciesToBrandWithContext(ctx context.Context, request *AddCustomPrivacyPoliciesToBrandRequest, runtime *dara.RuntimeOptions) (_result *AddCustomPrivacyPoliciesToBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.CustomPrivacyPolicyIds) {
		query["CustomPrivacyPolicyIds"] = request.CustomPrivacyPolicyIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomPrivacyPoliciesToBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomPrivacyPoliciesToBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a group to an authorization rule.
//
// @param request - AddGroupToAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupToAuthorizationRuleResponse
func (client *Client) AddGroupToAuthorizationRuleWithContext(ctx context.Context, request *AddGroupToAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *AddGroupToAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	if !dara.IsNil(request.ValidityType) {
		query["ValidityType"] = request.ValidityType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGroupToAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGroupToAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an account to an authorization rule.
//
// @param request - AddUserToAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToAuthorizationRuleResponse
func (client *Client) AddUserToAuthorizationRuleWithContext(ctx context.Context, request *AddUserToAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *AddUserToAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	if !dara.IsNil(request.ValidityType) {
		query["ValidityType"] = request.ValidityType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a specified EIAM account to multiple EIAM organizations. If the account already exists in an organization, the operation returns a success response directly.
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnitsWithContext(ctx context.Context, request *AddUserToOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds multiple Employee Identity and Access Management (EIAM) accounts to a specified EIAM account group.
//
// @param request - AddUsersToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroupWithContext(ctx context.Context, request *AddUsersToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddUsersToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants multiple EIAM groups access to an application in a batch.
//
// @param request - AuthorizeApplicationToGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToGroupsResponse
func (client *Client) AuthorizeApplicationToGroupsWithContext(ctx context.Context, request *AuthorizeApplicationToGroupsRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeApplicationToGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants access to an application for multiple EIAM organizations in a batch operation.
//
// @param request - AuthorizeApplicationToOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToOrganizationalUnitsResponse
func (client *Client) AuthorizeApplicationToOrganizationalUnitsWithContext(ctx context.Context, request *AuthorizeApplicationToOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeApplicationToOrganizationalUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants multiple EIAM accounts access to an application in a batch.
//
// @param request - AuthorizeApplicationToUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToUsersResponse
func (client *Client) AuthorizeApplicationToUsersWithContext(ctx context.Context, request *AuthorizeApplicationToUsersRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeApplicationToUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants Scope permissions under a specified ResourceServer to a Client application.
//
// @param request - AuthorizeResourceServerScopesToClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeResourceServerScopesToClientResponse
func (client *Client) AuthorizeResourceServerScopesToClientWithContext(ctx context.Context, request *AuthorizeResourceServerScopesToClientRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeResourceServerScopesToClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientApplicationId) {
		query["ClientApplicationId"] = request.ClientApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerApplicationId) {
		query["ResourceServerApplicationId"] = request.ResourceServerApplicationId
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeResourceServerScopesToClient"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeResourceServerScopesToClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants Scope permissions under a specified ResourceServer to a group.
//
// @param request - AuthorizeResourceServerScopesToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeResourceServerScopesToGroupResponse
func (client *Client) AuthorizeResourceServerScopesToGroupWithContext(ctx context.Context, request *AuthorizeResourceServerScopesToGroupRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeResourceServerScopesToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeResourceServerScopesToGroup"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeResourceServerScopesToGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants scope permissions under a specified ResourceServer to an organizational unit.
//
// @param request - AuthorizeResourceServerScopesToOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeResourceServerScopesToOrganizationalUnitResponse
func (client *Client) AuthorizeResourceServerScopesToOrganizationalUnitWithContext(ctx context.Context, request *AuthorizeResourceServerScopesToOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeResourceServerScopesToOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeResourceServerScopesToOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeResourceServerScopesToOrganizationalUnitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants Scope permissions under a specified ResourceServer to an account.
//
// @param request - AuthorizeResourceServerScopesToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeResourceServerScopesToUserResponse
func (client *Client) AuthorizeResourceServerScopesToUserWithContext(ctx context.Context, request *AuthorizeResourceServerScopesToUserRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeResourceServerScopesToUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeResourceServerScopesToUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeResourceServerScopesToUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a specified ResourceServer for a Client application.
//
// @param request - AuthorizeResourceServerToClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeResourceServerToClientResponse
func (client *Client) AuthorizeResourceServerToClientWithContext(ctx context.Context, request *AuthorizeResourceServerToClientRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeResourceServerToClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientApplicationId) {
		query["ClientApplicationId"] = request.ClientApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerApplicationId) {
		query["ResourceServerApplicationId"] = request.ResourceServerApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeResourceServerToClient"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeResourceServerToClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a third-party logon account.
//
// @param request - BindUserAuthnSourceMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindUserAuthnSourceMappingResponse
func (client *Client) BindUserAuthnSourceMappingWithContext(ctx context.Context, request *BindUserAuthnSourceMappingRequest, runtime *dara.RuntimeOptions) (_result *BindUserAuthnSourceMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserExternalId) {
		query["UserExternalId"] = request.UserExternalId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindUserAuthnSourceMapping"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindUserAuthnSourceMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the primary organization for application synchronization is within the application synchronization scope.
//
// @param request - CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse
func (client *Client) CheckApplicationProvisioningUserPrimaryOrganizationalUnitWithContext(ctx context.Context, request *CheckApplicationProvisioningUserPrimaryOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserPrimaryOrganizationalUnitId) {
		query["UserPrimaryOrganizationalUnitId"] = request.UserPrimaryOrganizationalUnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckApplicationProvisioningUserPrimaryOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckApplicationProvisioningUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether an instance can be deleted.
//
// @param request - CheckInstanceForDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceForDeleteResponse
func (client *Client) CheckInstanceForDeleteWithContext(ctx context.Context, request *CheckInstanceForDeleteRequest, runtime *dara.RuntimeOptions) (_result *CheckInstanceForDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckInstanceForDelete"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInstanceForDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Determines whether an instance has the feature of a specific module.
//
// @param request - CheckInstanceModuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceModuleStatusResponse
func (client *Client) CheckInstanceModuleStatusWithContext(ctx context.Context, request *CheckInstanceModuleStatusRequest, runtime *dara.RuntimeOptions) (_result *CheckInstanceModuleStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureKey) {
		query["FeatureKey"] = request.FeatureKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleKey) {
		query["ModuleKey"] = request.ModuleKey
	}

	if !dara.IsNil(request.SubFeatureKey) {
		query["SubFeatureKey"] = request.SubFeatureKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckInstanceModuleStatus"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInstanceModuleStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application resource in a specified EIAM instance.
//
// Description:
//
// EIAM supports two standard SSO protocols for application access: SAML 2.0 and OIDC. This parameter cannot be changed after it is specified. Create the application based on the SSO protocol type used in your business scenario.
//
// @param request - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIdentityType) {
		query["ApplicationIdentityType"] = request.ApplicationIdentityType
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ApplicationOwner) {
		query["ApplicationOwner"] = request.ApplicationOwner
	}

	if !dara.IsNil(request.ApplicationSourceType) {
		query["ApplicationSourceType"] = request.ApplicationSourceType
	}

	if !dara.IsNil(request.ApplicationTemplateId) {
		query["ApplicationTemplateId"] = request.ApplicationTemplateId
	}

	if !dara.IsNil(request.CustomFields) {
		query["CustomFields"] = request.CustomFields
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a client secret for an EIAM application. You can create a maximum of two client secrets for each EIAM application.
//
// @param request - CreateApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationClientSecretResponse
func (client *Client) CreateApplicationClientSecretWithContext(ctx context.Context, request *CreateApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationClientSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ExpirationTime) {
		query["ExpirationTime"] = request.ExpirationTime
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application federated identity credential.
//
// @param request - CreateApplicationFederatedCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationFederatedCredentialResponse
func (client *Client) CreateApplicationFederatedCredentialWithContext(ctx context.Context, request *CreateApplicationFederatedCredentialRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationFederatedCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationFederatedCredentialName) {
		query["ApplicationFederatedCredentialName"] = request.ApplicationFederatedCredentialName
	}

	if !dara.IsNil(request.ApplicationFederatedCredentialType) {
		query["ApplicationFederatedCredentialType"] = request.ApplicationFederatedCredentialType
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AttributeMappings) {
		query["AttributeMappings"] = request.AttributeMappings
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FederatedCredentialProviderId) {
		query["FederatedCredentialProviderId"] = request.FederatedCredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OidcVerificationConfig) {
		query["OidcVerificationConfig"] = request.OidcVerificationConfig
	}

	if !dara.IsNil(request.Pkcs7VerificationConfig) {
		query["Pkcs7VerificationConfig"] = request.Pkcs7VerificationConfig
	}

	if !dara.IsNil(request.VerificationCondition) {
		query["VerificationCondition"] = request.VerificationCondition
	}

	if !dara.IsNil(request.VerificationMode) {
		query["VerificationMode"] = request.VerificationMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationFederatedCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationFederatedCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application role.
//
// @param request - CreateApplicationRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationRoleResponse
func (client *Client) CreateApplicationRoleWithContext(ctx context.Context, request *CreateApplicationRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleName) {
		query["ApplicationRoleName"] = request.ApplicationRoleName
	}

	if !dara.IsNil(request.ApplicationRoleValue) {
		query["ApplicationRoleValue"] = request.ApplicationRoleValue
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application token.
//
// @param request - CreateApplicationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationTokenResponse
func (client *Client) CreateApplicationTokenWithContext(ctx context.Context, request *CreateApplicationTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationTokenType) {
		query["ApplicationTokenType"] = request.ApplicationTokenType
	}

	if !dara.IsNil(request.ExpirationTime) {
		query["ExpirationTime"] = request.ExpirationTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an authorization resource.
//
// @param request - CreateAuthorizationResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthorizationResourceResponse
func (client *Client) CreateAuthorizationResourceWithContext(ctx context.Context, request *CreateAuthorizationResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateAuthorizationResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceEntityId) {
		query["AuthorizationResourceEntityId"] = request.AuthorizationResourceEntityId
	}

	if !dara.IsNil(request.AuthorizationResourceEntityType) {
		query["AuthorizationResourceEntityType"] = request.AuthorizationResourceEntityType
	}

	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAuthorizationResource"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAuthorizationResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an authorization rule.
//
// @param request - CreateAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthorizationRuleResponse
func (client *Client) CreateAuthorizationRuleWithContext(ctx context.Context, request *CreateAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceScope) {
		query["AuthorizationResourceScope"] = request.AuthorizationResourceScope
	}

	if !dara.IsNil(request.AuthorizationRuleName) {
		query["AuthorizationRuleName"] = request.AuthorizationRuleName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a brand.
//
// @param request - CreateBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBrandResponse
func (client *Client) CreateBrandWithContext(ctx context.Context, request *CreateBrandRequest, runtime *dara.RuntimeOptions) (_result *CreateBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandName) {
		query["BrandName"] = request.BrandName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ClientPublicKey for an application. When an M2M client uses the PRIVATE_KEY_JWT method to request the token endpoint, this public key is used by the M2M authorization server to verify the assertion carried in the client\\"s token endpoint request.
//
// @param request - CreateClientPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientPublicKeyResponse
func (client *Client) CreateClientPublicKeyWithContext(ctx context.Context, request *CreateClientPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateClientPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlgorithmType) {
		query["AlgorithmType"] = request.AlgorithmType
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PublicKey) {
		query["PublicKey"] = request.PublicKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientPublicKey"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud account resource in a specified EIAM instance.
//
// Description:
//
// *Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.aliyun.com/price/product#/ecs/detail) of IDaaS EIAM.**
//
// @param request - CreateCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudAccountResponse
func (client *Client) CreateCloudAccountWithContext(ctx context.Context, request *CreateCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudAccountExternalId) {
		query["CloudAccountExternalId"] = request.CloudAccountExternalId
	}

	if !dara.IsNil(request.CloudAccountName) {
		query["CloudAccountName"] = request.CloudAccountName
	}

	if !dara.IsNil(request.CloudAccountProviderName) {
		query["CloudAccountProviderName"] = request.CloudAccountProviderName
	}

	if !dara.IsNil(request.CloudAccountSite) {
		query["CloudAccountSite"] = request.CloudAccountSite
	}

	if !dara.IsNil(request.CloudAccountVendorType) {
		query["CloudAccountVendorType"] = request.CloudAccountVendorType
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
		Action:      dara.String("CreateCloudAccount"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud role for a specified Alibaba Cloud account.
//
// @param request - CreateCloudAccountRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudAccountRoleResponse
func (client *Client) CreateCloudAccountRoleWithContext(ctx context.Context, request *CreateCloudAccountRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudAccountRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.CloudAccountRoleName) {
		query["CloudAccountRoleName"] = request.CloudAccountRoleName
	}

	if !dara.IsNil(request.CloudAccountRoleType) {
		query["CloudAccountRoleType"] = request.CloudAccountRoleType
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
		Action:      dara.String("CreateCloudAccountRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudAccountRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a conditional access policy.
//
// Description:
//
// Creates a conditional access policy.
//
// @param request - CreateConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConditionalAccessPolicyResponse
func (client *Client) CreateConditionalAccessPolicyWithContext(ctx context.Context, request *CreateConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateConditionalAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a credential resource in a specified EIAM instance.
//
// @param request - CreateCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCredentialResponse
func (client *Client) CreateCredentialWithContext(ctx context.Context, request *CreateCredentialRequest, runtime *dara.RuntimeOptions) (_result *CreateCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredentialContent) {
		query["CredentialContent"] = request.CredentialContent
	}

	if !dara.IsNil(request.CredentialExternalId) {
		query["CredentialExternalId"] = request.CredentialExternalId
	}

	if !dara.IsNil(request.CredentialIdentifier) {
		query["CredentialIdentifier"] = request.CredentialIdentifier
	}

	if !dara.IsNil(request.CredentialName) {
		query["CredentialName"] = request.CredentialName
	}

	if !dara.IsNil(request.CredentialScenarioLabel) {
		query["CredentialScenarioLabel"] = request.CredentialScenarioLabel
	}

	if !dara.IsNil(request.CredentialSharingScope) {
		query["CredentialSharingScope"] = request.CredentialSharingScope
	}

	if !dara.IsNil(request.CredentialSubjectId) {
		query["CredentialSubjectId"] = request.CredentialSubjectId
	}

	if !dara.IsNil(request.CredentialSubjectType) {
		query["CredentialSubjectType"] = request.CredentialSubjectType
	}

	if !dara.IsNil(request.CredentialType) {
		query["CredentialType"] = request.CredentialType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExclusiveUserId) {
		query["ExclusiveUserId"] = request.ExclusiveUserId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a credential provider.
//
// @param request - CreateCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCredentialProviderResponse
func (client *Client) CreateCredentialProviderWithContext(ctx context.Context, request *CreateCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredentialProviderConfig) {
		query["CredentialProviderConfig"] = request.CredentialProviderConfig
	}

	if !dara.IsNil(request.CredentialProviderIdentifier) {
		query["CredentialProviderIdentifier"] = request.CredentialProviderIdentifier
	}

	if !dara.IsNil(request.CredentialProviderName) {
		query["CredentialProviderName"] = request.CredentialProviderName
	}

	if !dara.IsNil(request.CredentialProviderType) {
		query["CredentialProviderType"] = request.CredentialProviderType
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
		Action:      dara.String("CreateCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an extended field.
//
// @param request - CreateCustomFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomFieldResponse
func (client *Client) CreateCustomFieldWithContext(ctx context.Context, request *CreateCustomFieldRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultValue) {
		query["DefaultValue"] = request.DefaultValue
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Encrypted) {
		query["Encrypted"] = request.Encrypted
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.FieldDataConfig) {
		query["FieldDataConfig"] = request.FieldDataConfig
	}

	if !dara.IsNil(request.FieldDataType) {
		query["FieldDataType"] = request.FieldDataType
	}

	if !dara.IsNil(request.FieldDisplayName) {
		query["FieldDisplayName"] = request.FieldDisplayName
	}

	if !dara.IsNil(request.FieldDisplayType) {
		query["FieldDisplayType"] = request.FieldDisplayType
	}

	if !dara.IsNil(request.FieldName) {
		query["FieldName"] = request.FieldName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Required) {
		query["Required"] = request.Required
	}

	if !dara.IsNil(request.Unique) {
		query["Unique"] = request.Unique
	}

	if !dara.IsNil(request.UserPermission) {
		query["UserPermission"] = request.UserPermission
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomField"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can create custom terms.
//
// @param request - CreateCustomPrivacyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomPrivacyPolicyResponse
func (client *Client) CreateCustomPrivacyPolicyWithContext(ctx context.Context, request *CreateCustomPrivacyPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomPrivacyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CustomPrivacyPolicyContents) {
		query["CustomPrivacyPolicyContents"] = request.CustomPrivacyPolicyContents
	}

	if !dara.IsNil(request.CustomPrivacyPolicyName) {
		query["CustomPrivacyPolicyName"] = request.CustomPrivacyPolicyName
	}

	if !dara.IsNil(request.DefaultLanguageCode) {
		query["DefaultLanguageCode"] = request.DefaultLanguageCode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserConsentType) {
		query["UserConsentType"] = request.UserConsentType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomPrivacyPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomPrivacyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom domain name for an EIAM (Entity and Identity Access Management) instance.
//
// @param request - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainProxyTokenResponse
func (client *Client) CreateDomainProxyTokenWithContext(ctx context.Context, request *CreateDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainProxyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a federated trust source.
//
// @param request - CreateFederatedCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFederatedCredentialProviderResponse
func (client *Client) CreateFederatedCredentialProviderWithContext(ctx context.Context, request *CreateFederatedCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateFederatedCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudIdPProviderConfig) {
		query["CloudIdPProviderConfig"] = request.CloudIdPProviderConfig
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FederatedCredentialProviderName) {
		query["FederatedCredentialProviderName"] = request.FederatedCredentialProviderName
	}

	if !dara.IsNil(request.FederatedCredentialProviderType) {
		query["FederatedCredentialProviderType"] = request.FederatedCredentialProviderType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !dara.IsNil(request.OidcProviderConfig) {
		query["OidcProviderConfig"] = request.OidcProviderConfig
	}

	if !dara.IsNil(request.Pkcs7ProviderConfig) {
		query["Pkcs7ProviderConfig"] = request.Pkcs7ProviderConfig
	}

	if !dara.IsNil(request.PrivateCaProviderConfig) {
		query["PrivateCaProviderConfig"] = request.PrivateCaProviderConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFederatedCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFederatedCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an EIAM account group.
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an identity provider.
//
// @param request - CreateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityProviderResponse
func (client *Client) CreateIdentityProviderWithContext(ctx context.Context, request *CreateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.SamlConfig) {
		query["SamlConfig"] = request.SamlConfig
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a status check job for an identity provider.
//
// @param request - CreateIdentityProviderStatusCheckJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityProviderStatusCheckJobResponse
func (client *Client) CreateIdentityProviderStatusCheckJobWithContext(ctx context.Context, request *CreateIdentityProviderStatusCheckJobRequest, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderStatusCheckJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("CreateIdentityProviderStatusCheckJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdentityProviderStatusCheckJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trial License for an instance.
//
// @param request - CreateInstanceTrialLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceTrialLicenseResponse
func (client *Client) CreateInstanceTrialLicenseWithContext(ctx context.Context, request *CreateInstanceTrialLicenseRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceTrialLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceTrialLicense"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceTrialLicenseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network access endpoint.
//
// @param request - CreateNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkAccessEndpointResponse
func (client *Client) CreateNetworkAccessEndpointWithContext(ctx context.Context, request *CreateNetworkAccessEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkAccessEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network zone object.
//
// @param request - CreateNetworkZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkZoneResponse
func (client *Client) CreateNetworkZoneWithContext(ctx context.Context, request *CreateNetworkZoneRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ipv4Cidrs) {
		query["Ipv4Cidrs"] = request.Ipv4Cidrs
	}

	if !dara.IsNil(request.Ipv6Cidrs) {
		query["Ipv6Cidrs"] = request.Ipv6Cidrs
	}

	if !dara.IsNil(request.NetworkZoneName) {
		query["NetworkZoneName"] = request.NetworkZoneName
	}

	if !dara.IsNil(request.NetworkZoneType) {
		query["NetworkZoneType"] = request.NetworkZoneType
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkZone"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an EIAM organization under a specified organization.
//
// @param request - CreateOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnitWithContext(ctx context.Context, request *CreateOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *CreateOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scope permission for a specified resource server.
//
// @param request - CreateResourceServerScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceServerScopeResponse
func (client *Client) CreateResourceServerScopeWithContext(ctx context.Context, request *CreateResourceServerScopeRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceServerScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.ResourceServerScopeName) {
		query["ResourceServerScopeName"] = request.ResourceServerScopeName
	}

	if !dara.IsNil(request.ResourceServerScopeType) {
		query["ResourceServerScopeType"] = request.ResourceServerScopeType
	}

	if !dara.IsNil(request.ResourceServerScopeValue) {
		query["ResourceServerScopeValue"] = request.ResourceServerScopeValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceServerScope"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceServerScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an EIAM account in a specified EIAM instance.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an EIAM application resource.
//
// Description:
//
// Before deleting a specified EIAM application, make sure the application is no longer in use. After the application is deleted, all configuration data is permanently deleted and cannot be recovered.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a client secret of an EIAM application.
//
// @param request - DeleteApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationClientSecretResponse
func (client *Client) DeleteApplicationClientSecretWithContext(ctx context.Context, request *DeleteApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationClientSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a federated identity credential of an application.
//
// @param request - DeleteApplicationFederatedCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationFederatedCredentialResponse
func (client *Client) DeleteApplicationFederatedCredentialWithContext(ctx context.Context, request *DeleteApplicationFederatedCredentialRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationFederatedCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationFederatedCredentialId) {
		query["ApplicationFederatedCredentialId"] = request.ApplicationFederatedCredentialId
	}

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
		Action:      dara.String("DeleteApplicationFederatedCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationFederatedCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application role.
//
// @param request - DeleteApplicationRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationRoleResponse
func (client *Client) DeleteApplicationRoleWithContext(ctx context.Context, request *DeleteApplicationRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ApplicationToken.
//
// @param request - DeleteApplicationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationTokenResponse
func (client *Client) DeleteApplicationTokenWithContext(ctx context.Context, request *DeleteApplicationTokenRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationTokenId) {
		query["ApplicationTokenId"] = request.ApplicationTokenId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an authorized resource.
//
// @param request - DeleteAuthorizationResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAuthorizationResourceResponse
func (client *Client) DeleteAuthorizationResourceWithContext(ctx context.Context, request *DeleteAuthorizationResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteAuthorizationResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceId) {
		query["AuthorizationResourceId"] = request.AuthorizationResourceId
	}

	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAuthorizationResource"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAuthorizationResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an authorization rule. The authorization rule can be deleted only when it is in the disabled state.
//
// @param request - DeleteAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAuthorizationRuleResponse
func (client *Client) DeleteAuthorizationRuleWithContext(ctx context.Context, request *DeleteAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Brand
//
// @param request - DeleteBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBrandResponse
func (client *Client) DeleteBrandWithContext(ctx context.Context, request *DeleteBrandRequest, runtime *dara.RuntimeOptions) (_result *DeleteBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified application ClientPublicKey.
//
// @param request - DeleteClientPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientPublicKeyResponse
func (client *Client) DeleteClientPublicKeyWithContext(ctx context.Context, request *DeleteClientPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientPublicKeyId) {
		query["ClientPublicKeyId"] = request.ClientPublicKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClientPublicKey"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Alibaba Cloud account resource.
//
// @param request - DeleteCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudAccountResponse
func (client *Client) DeleteCloudAccountWithContext(ctx context.Context, request *DeleteCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudAccount"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cloud role from a specified Alibaba Cloud account.
//
// Description:
//
// You must disable the cloud role before you delete it. After you delete the role, all related data is also deleted and cannot be recovered.
//
// @param request - DeleteCloudAccountRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudAccountRoleResponse
func (client *Client) DeleteCloudAccountRoleWithContext(ctx context.Context, request *DeleteCloudAccountRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudAccountRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.CloudAccountRoleId) {
		query["CloudAccountRoleId"] = request.CloudAccountRoleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudAccountRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudAccountRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a conditional access policy.
//
// Description:
//
// Before you delete a conditional access policy, make sure that the policy is no longer in use. After the policy is deleted, all configuration data is permanently removed and cannot be recovered.
//
// @param request - DeleteConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConditionalAccessPolicyResponse
func (client *Client) DeleteConditionalAccessPolicyWithContext(ctx context.Context, request *DeleteConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteConditionalAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a credential resource.
//
// @param request - DeleteCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredentialWithContext(ctx context.Context, request *DeleteCredentialRequest, runtime *dara.RuntimeOptions) (_result *DeleteCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a credential provider.
//
// @param request - DeleteCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCredentialProviderResponse
func (client *Client) DeleteCredentialProviderWithContext(ctx context.Context, request *DeleteCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialProviderId) {
		query["CredentialProviderId"] = request.CredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an extended character field.
//
// @param request - DeleteCustomFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomFieldResponse
func (client *Client) DeleteCustomFieldWithContext(ctx context.Context, request *DeleteCustomFieldRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FieldId) {
		query["FieldId"] = request.FieldId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomField"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom clause.
//
// @param request - DeleteCustomPrivacyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomPrivacyPolicyResponse
func (client *Client) DeleteCustomPrivacyPolicyWithContext(ctx context.Context, request *DeleteCustomPrivacyPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomPrivacyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrivacyPolicyId) {
		query["CustomPrivacyPolicyId"] = request.CustomPrivacyPolicyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomPrivacyPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomPrivacyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom domain name of a specified Employee Identity and Access Management (EIAM) instance. Deletion of the instance initialization domain name or default domain name is rejected.
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a proxy token for a specified domain name of an EIAM instance. Only proxy tokens in the disabled state can be deleted.
//
// @param request - DeleteDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainProxyTokenResponse
func (client *Client) DeleteDomainProxyTokenWithContext(ctx context.Context, request *DeleteDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainProxyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a federated credential provider.
//
// @param request - DeleteFederatedCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFederatedCredentialProviderResponse
func (client *Client) DeleteFederatedCredentialProviderWithContext(ctx context.Context, request *DeleteFederatedCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteFederatedCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederatedCredentialProviderId) {
		query["FederatedCredentialProviderId"] = request.FederatedCredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFederatedCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFederatedCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an EIAM account group.
//
// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) instance that you no longer use.
//
// Description:
//
// Make sure that the instance is no longer in use. After an EIAM instance is deleted, all related data is deleted.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dedicated network access endpoint.
//
// @param request - DeleteNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkAccessEndpointResponse
func (client *Client) DeleteNetworkAccessEndpointWithContext(ctx context.Context, request *DeleteNetworkAccessEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkAccessEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a network zone object.
//
// @param request - DeleteNetworkZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkZoneResponse
func (client *Client) DeleteNetworkZoneWithContext(ctx context.Context, request *DeleteNetworkZoneRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DeleteNetworkZone"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified EIAM organizational unit. The deletion fails if the organizational unit contains EIAM accounts or child organizational units.
//
// @param request - DeleteOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnitWithContext(ctx context.Context, request *DeleteOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *DeleteOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes all data of a specified organization. This operation can force delete all accounts and sub-organizations under the specified organization.
//
// @param request - DeleteOrganizationalUnitChildrenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOrganizationalUnitChildrenResponse
func (client *Client) DeleteOrganizationalUnitChildrenWithContext(ctx context.Context, request *DeleteOrganizationalUnitChildrenRequest, runtime *dara.RuntimeOptions) (_result *DeleteOrganizationalUnitChildrenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scope permission under a specified ResourceServer.
//
// @param request - DeleteResourceServerScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceServerScopeResponse
func (client *Client) DeleteResourceServerScopeWithContext(ctx context.Context, request *DeleteResourceServerScopeRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceServerScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerScopeId) {
		query["ResourceServerScopeId"] = request.ResourceServerScopeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceServerScope"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceServerScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified EIAM account and purges all information associated with the account.
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple accounts in a batch.
//
// @param request - DeleteUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUsersResponse
func (client *Client) DeleteUsersWithContext(ctx context.Context, request *DeleteUsersRequest, runtime *dara.RuntimeOptions) (_result *DeleteUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DeleteUsers"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified WebAuthn authenticator.
//
// @param request - DeleteWebAuthnAuthenticatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebAuthnAuthenticatorResponse
func (client *Client) DeleteWebAuthnAuthenticatorWithContext(ctx context.Context, request *DeleteWebAuthnAuthenticatorRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebAuthnAuthenticatorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticatorId) {
		query["AuthenticatorId"] = request.AuthenticatorId
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
		Action:      dara.String("DeleteWebAuthnAuthenticator"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebAuthnAuthenticatorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an application by changing its status from enabled to disabled, making all capabilities of the application unavailable.
//
// Description:
//
// When you change an application from the enabled state to the disabled state, all features of the application become unavailable (such as SSO and account synchronization). Confirm that you are aware of the risks that this operation may cause.
//
// @param request - DisableApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationResponse
func (client *Client) DisableApplicationWithContext(ctx context.Context, request *DisableApplicationRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationApiInvokeResponse
func (client *Client) DisableApplicationApiInvokeWithContext(ctx context.Context, request *DisableApplicationApiInvokeRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationApiInvokeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationClientSecretResponse
func (client *Client) DisableApplicationClientSecretWithContext(ctx context.Context, request *DisableApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationClientSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a federated identity credential for an application.
//
// @param request - DisableApplicationFederatedCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationFederatedCredentialResponse
func (client *Client) DisableApplicationFederatedCredentialWithContext(ctx context.Context, request *DisableApplicationFederatedCredentialRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationFederatedCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationFederatedCredentialId) {
		query["ApplicationFederatedCredentialId"] = request.ApplicationFederatedCredentialId
	}

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
		Action:      dara.String("DisableApplicationFederatedCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationFederatedCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the machine-to-machine (M2M) client feature for an application. This feature enables an application to act as an OAuth client and make calls to access resources.
//
// @param request - DisableApplicationM2MClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationM2MClientResponse
func (client *Client) DisableApplicationM2MClientWithContext(ctx context.Context, request *DisableApplicationM2MClientRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationM2MClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DisableApplicationM2MClient"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationM2MClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - DisableApplicationProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationProvisioningResponse
func (client *Client) DisableApplicationProvisioningWithContext(ctx context.Context, request *DisableApplicationProvisioningRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationProvisioningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the ResourceServer capability of a specified application.
//
// @param request - DisableApplicationResourceServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationResourceServerResponse
func (client *Client) DisableApplicationResourceServerWithContext(ctx context.Context, request *DisableApplicationResourceServerRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationResourceServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DisableApplicationResourceServer"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationResourceServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationSsoResponse
func (client *Client) DisableApplicationSsoWithContext(ctx context.Context, request *DisableApplicationSsoRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationSsoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an application token.
//
// Description:
//
// When you change an application from the enabled state to the disabled state, all features of the application become unavailable, such as SSO and account synchronization. Make sure that you are aware of the risks that this operation may cause.
//
// @param request - DisableApplicationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationTokenResponse
func (client *Client) DisableApplicationTokenWithContext(ctx context.Context, request *DisableApplicationTokenRequest, runtime *dara.RuntimeOptions) (_result *DisableApplicationTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationTokenId) {
		query["ApplicationTokenId"] = request.ApplicationTokenId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplicationToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an authorization rule.
//
// @param request - DisableAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAuthorizationRuleResponse
func (client *Client) DisableAuthorizationRuleWithContext(ctx context.Context, request *DisableAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a brand.
//
// @param request - DisableBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableBrandResponse
func (client *Client) DisableBrandWithContext(ctx context.Context, request *DisableBrandRequest, runtime *dara.RuntimeOptions) (_result *DisableBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a specified application ClientPublicKey.
//
// @param request - DisableClientPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableClientPublicKeyResponse
func (client *Client) DisableClientPublicKeyWithContext(ctx context.Context, request *DisableClientPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *DisableClientPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientPublicKeyId) {
		query["ClientPublicKeyId"] = request.ClientPublicKeyId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableClientPublicKey"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableClientPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a cloud role for a specified Alibaba Cloud account.
//
// @param request - DisableCloudAccountRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCloudAccountRoleResponse
func (client *Client) DisableCloudAccountRoleWithContext(ctx context.Context, request *DisableCloudAccountRoleRequest, runtime *dara.RuntimeOptions) (_result *DisableCloudAccountRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.CloudAccountRoleId) {
		query["CloudAccountRoleId"] = request.CloudAccountRoleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCloudAccountRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCloudAccountRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a conditional access policy.
//
// Description:
//
// When you change a conditional access policy from the enabled state to the disabled state, the policy no longer intercepts requests. Confirm that you are aware of the security risks that this operation may cause.
//
// @param request - DisableConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableConditionalAccessPolicyResponse
func (client *Client) DisableConditionalAccessPolicyWithContext(ctx context.Context, request *DisableConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *DisableConditionalAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a credential resource.
//
// @param request - DisableCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCredentialResponse
func (client *Client) DisableCredentialWithContext(ctx context.Context, request *DisableCredentialRequest, runtime *dara.RuntimeOptions) (_result *DisableCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a credential provider.
//
// @param request - DisableCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCredentialProviderResponse
func (client *Client) DisableCredentialProviderWithContext(ctx context.Context, request *DisableCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DisableCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialProviderId) {
		query["CredentialProviderId"] = request.CredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a field.
//
// @param request - DisableCustomFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCustomFieldResponse
func (client *Client) DisableCustomFieldWithContext(ctx context.Context, request *DisableCustomFieldRequest, runtime *dara.RuntimeOptions) (_result *DisableCustomFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FieldId) {
		query["FieldId"] = request.FieldId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCustomField"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCustomFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a custom term.
//
// @param request - DisableCustomPrivacyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCustomPrivacyPolicyResponse
func (client *Client) DisableCustomPrivacyPolicyWithContext(ctx context.Context, request *DisableCustomPrivacyPolicyRequest, runtime *dara.RuntimeOptions) (_result *DisableCustomPrivacyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrivacyPolicyId) {
		query["CustomPrivacyPolicyId"] = request.CustomPrivacyPolicyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCustomPrivacyPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCustomPrivacyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a specified domain proxy token for an EIAM instance. After the token is disabled, domain usage is affected.
//
// @param request - DisableDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDomainProxyTokenResponse
func (client *Client) DisableDomainProxyTokenWithContext(ctx context.Context, request *DisableDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *DisableDomainProxyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a federated credential provider.
//
// @param request - DisableFederatedCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableFederatedCredentialProviderResponse
func (client *Client) DisableFederatedCredentialProviderWithContext(ctx context.Context, request *DisableFederatedCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DisableFederatedCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederatedCredentialProviderId) {
		query["FederatedCredentialProviderId"] = request.FederatedCredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableFederatedCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableFederatedCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the advanced configuration.
//
// @param request - DisableIdentityProviderAdvancedAbilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableIdentityProviderAdvancedAbilityResponse
func (client *Client) DisableIdentityProviderAdvancedAbilityWithContext(ctx context.Context, request *DisableIdentityProviderAdvancedAbilityRequest, runtime *dara.RuntimeOptions) (_result *DisableIdentityProviderAdvancedAbilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DisableIdentityProviderAdvancedAbility"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableIdentityProviderAdvancedAbilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables authentication.
//
// Description:
//
// When you change an application from the enabled state to the disabled state, all features of the application become unavailable, such as SSO and account synchronization. Make sure that you are aware of the risks that this operation may cause.
//
// @param request - DisableIdentityProviderAuthnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableIdentityProviderAuthnResponse
func (client *Client) DisableIdentityProviderAuthnWithContext(ctx context.Context, request *DisableIdentityProviderAuthnRequest, runtime *dara.RuntimeOptions) (_result *DisableIdentityProviderAuthnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DisableIdentityProviderAuthn"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableIdentityProviderAuthnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the inbound synchronization feature of an identity provider.
//
// @param request - DisableIdentityProviderUdPullRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableIdentityProviderUdPullResponse
func (client *Client) DisableIdentityProviderUdPullWithContext(ctx context.Context, request *DisableIdentityProviderUdPullRequest, runtime *dara.RuntimeOptions) (_result *DisableIdentityProviderUdPullResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables automatic redirect from the initialization domain to the default domain for an EIAM instance. After disabling, accessing the portal via the initialization domain will no longer redirect to the default domain.
//
// @param request - DisableInitDomainAutoRedirectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInitDomainAutoRedirectResponse
func (client *Client) DisableInitDomainAutoRedirectWithContext(ctx context.Context, request *DisableInitDomainAutoRedirectRequest, runtime *dara.RuntimeOptions) (_result *DisableInitDomainAutoRedirectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an internal authentication source.
//
// @param request - DisableInternalAuthenticationSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInternalAuthenticationSourceResponse
func (client *Client) DisableInternalAuthenticationSourceWithContext(ctx context.Context, request *DisableInternalAuthenticationSourceRequest, runtime *dara.RuntimeOptions) (_result *DisableInternalAuthenticationSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationSourceId) {
		query["AuthenticationSourceId"] = request.AuthenticationSourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableInternalAuthenticationSource"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInternalAuthenticationSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the custom subject feature for a specified resource server.
//
// @param request - DisableResourceServerCustomSubjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableResourceServerCustomSubjectResponse
func (client *Client) DisableResourceServerCustomSubjectWithContext(ctx context.Context, request *DisableResourceServerCustomSubjectRequest, runtime *dara.RuntimeOptions) (_result *DisableResourceServerCustomSubjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DisableResourceServerCustomSubject"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableResourceServerCustomSubjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets an account status to disabled. If the account is already disabled, the operation returns success directly.
//
// @param request - DisableUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserResponse
func (client *Client) DisableUserWithContext(ctx context.Context, request *DisableUserRequest, runtime *dara.RuntimeOptions) (_result *DisableUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationResponse
func (client *Client) EnableApplicationWithContext(ctx context.Context, request *EnableApplicationRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the EnableApplicationApiInvoke operation to enable Developer API calls for an EIAM application.
//
// @param request - EnableApplicationApiInvokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationApiInvokeResponse
func (client *Client) EnableApplicationApiInvokeWithContext(ctx context.Context, request *EnableApplicationApiInvokeRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationApiInvokeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the client key of an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - EnableApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationClientSecretResponse
func (client *Client) EnableApplicationClientSecretWithContext(ctx context.Context, request *EnableApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationClientSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an application federated credential.
//
// @param request - EnableApplicationFederatedCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationFederatedCredentialResponse
func (client *Client) EnableApplicationFederatedCredentialWithContext(ctx context.Context, request *EnableApplicationFederatedCredentialRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationFederatedCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationFederatedCredentialId) {
		query["ApplicationFederatedCredentialId"] = request.ApplicationFederatedCredentialId
	}

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
		Action:      dara.String("EnableApplicationFederatedCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationFederatedCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the machine-to-machine (M2M) client feature for an application. This allows the application to act as a caller (an OAuth client) to access resources.
//
// @param request - EnableApplicationM2MClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationM2MClientResponse
func (client *Client) EnableApplicationM2MClientWithContext(ctx context.Context, request *EnableApplicationM2MClientRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationM2MClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("EnableApplicationM2MClient"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationM2MClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - EnableApplicationProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationProvisioningResponse
func (client *Client) EnableApplicationProvisioningWithContext(ctx context.Context, request *EnableApplicationProvisioningRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationProvisioningResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the ResourceServer feature for a specified application.
//
// @param request - EnableApplicationResourceServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationResourceServerResponse
func (client *Client) EnableApplicationResourceServerWithContext(ctx context.Context, request *EnableApplicationResourceServerRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationResourceServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("EnableApplicationResourceServer"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationResourceServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the SSO feature for an EIAM application.
//
// @param request - EnableApplicationSsoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationSsoResponse
func (client *Client) EnableApplicationSsoWithContext(ctx context.Context, request *EnableApplicationSsoRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationSsoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an application token.
//
// @param request - EnableApplicationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationTokenResponse
func (client *Client) EnableApplicationTokenWithContext(ctx context.Context, request *EnableApplicationTokenRequest, runtime *dara.RuntimeOptions) (_result *EnableApplicationTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationTokenId) {
		query["ApplicationTokenId"] = request.ApplicationTokenId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplicationToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an authorization rule.
//
// @param request - EnableAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAuthorizationRuleResponse
func (client *Client) EnableAuthorizationRuleWithContext(ctx context.Context, request *EnableAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a brand.
//
// @param request - EnableBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableBrandResponse
func (client *Client) EnableBrandWithContext(ctx context.Context, request *EnableBrandRequest, runtime *dara.RuntimeOptions) (_result *EnableBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a specified application ClientPublicKey.
//
// @param request - EnableClientPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableClientPublicKeyResponse
func (client *Client) EnableClientPublicKeyWithContext(ctx context.Context, request *EnableClientPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *EnableClientPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientPublicKeyId) {
		query["ClientPublicKeyId"] = request.ClientPublicKeyId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableClientPublicKey"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableClientPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a cloud role for a specified Alibaba Cloud account.
//
// @param request - EnableCloudAccountRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCloudAccountRoleResponse
func (client *Client) EnableCloudAccountRoleWithContext(ctx context.Context, request *EnableCloudAccountRoleRequest, runtime *dara.RuntimeOptions) (_result *EnableCloudAccountRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.CloudAccountRoleId) {
		query["CloudAccountRoleId"] = request.CloudAccountRoleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCloudAccountRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCloudAccountRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a conditional access policy.
//
// Description:
//
// When you change a conditional access policy from enabled to disabled, the policy no longer intercepts requests. Confirm that you are aware of the risks that this operation may cause.
//
// @param request - EnableConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableConditionalAccessPolicyResponse
func (client *Client) EnableConditionalAccessPolicyWithContext(ctx context.Context, request *EnableConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *EnableConditionalAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a credential resource.
//
// @param request - EnableCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCredentialResponse
func (client *Client) EnableCredentialWithContext(ctx context.Context, request *EnableCredentialRequest, runtime *dara.RuntimeOptions) (_result *EnableCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a credential provider.
//
// @param request - EnableCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCredentialProviderResponse
func (client *Client) EnableCredentialProviderWithContext(ctx context.Context, request *EnableCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *EnableCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialProviderId) {
		query["CredentialProviderId"] = request.CredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a field.
//
// Description:
//
// *Before you use this operation, make sure that you fully understand the billing method and [pricing](https://www.aliyun.com/price/product#/ecs/detail) of IDaaS EIAM.**.
//
// @param request - EnableCustomFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCustomFieldResponse
func (client *Client) EnableCustomFieldWithContext(ctx context.Context, request *EnableCustomFieldRequest, runtime *dara.RuntimeOptions) (_result *EnableCustomFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FieldId) {
		query["FieldId"] = request.FieldId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCustomField"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCustomFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a custom clause.
//
// @param request - EnableCustomPrivacyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCustomPrivacyPolicyResponse
func (client *Client) EnableCustomPrivacyPolicyWithContext(ctx context.Context, request *EnableCustomPrivacyPolicyRequest, runtime *dara.RuntimeOptions) (_result *EnableCustomPrivacyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrivacyPolicyId) {
		query["CustomPrivacyPolicyId"] = request.CustomPrivacyPolicyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCustomPrivacyPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCustomPrivacyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a domain proxy token for a specified EIAM instance domain name to verify domain security.
//
// @param request - EnableDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDomainProxyTokenResponse
func (client *Client) EnableDomainProxyTokenWithContext(ctx context.Context, request *EnableDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *EnableDomainProxyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a federated credential provider.
//
// @param request - EnableFederatedCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableFederatedCredentialProviderResponse
func (client *Client) EnableFederatedCredentialProviderWithContext(ctx context.Context, request *EnableFederatedCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *EnableFederatedCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederatedCredentialProviderId) {
		query["FederatedCredentialProviderId"] = request.FederatedCredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableFederatedCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableFederatedCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables advanced configuration.
//
// @param request - EnableIdentityProviderAdvancedAbilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableIdentityProviderAdvancedAbilityResponse
func (client *Client) EnableIdentityProviderAdvancedAbilityWithContext(ctx context.Context, request *EnableIdentityProviderAdvancedAbilityRequest, runtime *dara.RuntimeOptions) (_result *EnableIdentityProviderAdvancedAbilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("EnableIdentityProviderAdvancedAbility"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableIdentityProviderAdvancedAbilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables authentication.
//
// Description:
//
// When you change a conditional access policy from enabled to disabled, the policy no longer intercepts requests. Make sure that you are aware of the risks that this operation may cause.
//
// @param request - EnableIdentityProviderAuthnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableIdentityProviderAuthnResponse
func (client *Client) EnableIdentityProviderAuthnWithContext(ctx context.Context, request *EnableIdentityProviderAuthnRequest, runtime *dara.RuntimeOptions) (_result *EnableIdentityProviderAuthnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("EnableIdentityProviderAuthn"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableIdentityProviderAuthnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the inbound synchronization feature for an identity provider.
//
// @param request - EnableIdentityProviderUdPullRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableIdentityProviderUdPullResponse
func (client *Client) EnableIdentityProviderUdPullWithContext(ctx context.Context, request *EnableIdentityProviderUdPullRequest, runtime *dara.RuntimeOptions) (_result *EnableIdentityProviderUdPullResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables automatic redirection from the initialization domain to the default domain for an EIAM instance. After this feature is enabled, accessing the portal through the initialization domain will redirect to the default domain address.
//
// @param request - EnableInitDomainAutoRedirectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInitDomainAutoRedirectResponse
func (client *Client) EnableInitDomainAutoRedirectWithContext(ctx context.Context, request *EnableInitDomainAutoRedirectRequest, runtime *dara.RuntimeOptions) (_result *EnableInitDomainAutoRedirectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an internal authentication source.
//
// @param request - EnableInternalAuthenticationSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInternalAuthenticationSourceResponse
func (client *Client) EnableInternalAuthenticationSourceWithContext(ctx context.Context, request *EnableInternalAuthenticationSourceRequest, runtime *dara.RuntimeOptions) (_result *EnableInternalAuthenticationSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationSourceId) {
		query["AuthenticationSourceId"] = request.AuthenticationSourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInternalAuthenticationSource"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInternalAuthenticationSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the custom subject feature for a specified resource server. After this feature is enabled, the subject of an issued access token changes from <clientId> to <clientId>:<client.activeSubjectUrn>. The `client.activeSubjectUrn` is set in the attribute mapping of the application\\"s federated identity provider.
//
// @param request - EnableResourceServerCustomSubjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableResourceServerCustomSubjectResponse
func (client *Client) EnableResourceServerCustomSubjectWithContext(ctx context.Context, request *EnableResourceServerCustomSubjectRequest, runtime *dara.RuntimeOptions) (_result *EnableResourceServerCustomSubjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("EnableResourceServerCustomSubject"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableResourceServerCustomSubjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets an account to the enabled state.
//
// @param request - EnableUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableUserResponse
func (client *Client) EnableUserWithContext(ctx context.Context, request *EnableUserRequest, runtime *dara.RuntimeOptions) (_result *EnableUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resolves the metadata for an identity provider.
//
// @param request - ExecIdentityProviderMetadataUrlResolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecIdentityProviderMetadataUrlResolutionResponse
func (client *Client) ExecIdentityProviderMetadataUrlResolutionWithContext(ctx context.Context, request *ExecIdentityProviderMetadataUrlResolutionRequest, runtime *dara.RuntimeOptions) (_result *ExecIdentityProviderMetadataUrlResolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !dara.IsNil(request.OidcIssuer) {
		query["OidcIssuer"] = request.OidcIssuer
	}

	if !dara.IsNil(request.SamlMetadataUrl) {
		query["SamlMetadataUrl"] = request.SamlMetadataUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecIdentityProviderMetadataUrlResolution"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecIdentityProviderMetadataUrlResolutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a download URL for file import results.
//
// @param request - GenerateDownloadUrlForSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDownloadUrlForSynchronizationJobResponse
func (client *Client) GenerateDownloadUrlForSynchronizationJobWithContext(ctx context.Context, request *GenerateDownloadUrlForSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *GenerateDownloadUrlForSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GenerateDownloadUrlForSynchronizationJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateDownloadUrlForSynchronizationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a file import template.
//
// @param request - GenerateFileImportTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateFileImportTemplateResponse
func (client *Client) GenerateFileImportTemplateWithContext(ctx context.Context, request *GenerateFileImportTemplateRequest, runtime *dara.RuntimeOptions) (_result *GenerateFileImportTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateFileImportTemplate"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateFileImportTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an access token for accessing a resource server by using a specified application as the client identity.
//
// @param request - GenerateOauthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateOauthTokenResponse
func (client *Client) GenerateOauthTokenWithContext(ctx context.Context, request *GenerateOauthTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateOauthTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Audience) {
		query["Audience"] = request.Audience
	}

	if !dara.IsNil(request.GrantType) {
		query["GrantType"] = request.GrantType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestedTokenType) {
		query["RequestedTokenType"] = request.RequestedTokenType
	}

	if !dara.IsNil(request.ScopeValues) {
		query["ScopeValues"] = request.ScopeValues
	}

	if !dara.IsNil(request.SubjectToken) {
		query["SubjectToken"] = request.SubjectToken
	}

	if !dara.IsNil(request.SubjectTokenType) {
		query["SubjectTokenType"] = request.SubjectTokenType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateOauthToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateOauthTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Upload Authentication
//
// @param request - GenerateUploadAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadAuthResponse
func (client *Client) GenerateUploadAuthWithContext(ctx context.Context, request *GenerateUploadAuthRequest, runtime *dara.RuntimeOptions) (_result *GenerateUploadAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Purpose) {
		query["Purpose"] = request.Purpose
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUploadAuth"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUploadAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a WebAuthn authenticator registration URL.
//
// @param request - GenerateWebAuthnAuthenticatorRegistrationUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateWebAuthnAuthenticatorRegistrationUrlResponse
func (client *Client) GenerateWebAuthnAuthenticatorRegistrationUrlWithContext(ctx context.Context, request *GenerateWebAuthnAuthenticatorRegistrationUrlRequest, runtime *dara.RuntimeOptions) (_result *GenerateWebAuthnAuthenticatorRegistrationUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
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
		Action:      dara.String("GenerateWebAuthnAuthenticatorRegistrationUrl"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateWebAuthnAuthenticatorRegistrationUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified EIAM application.
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithContext(ctx context.Context, request *GetApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the advanced configuration of an application.
//
// @param request - GetApplicationAdvancedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationAdvancedConfigResponse
func (client *Client) GetApplicationAdvancedConfigWithContext(ctx context.Context, request *GetApplicationAdvancedConfigRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationAdvancedConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetApplicationAdvancedConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationAdvancedConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the federated identity credential of an application.
//
// @param request - GetApplicationFederatedCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationFederatedCredentialResponse
func (client *Client) GetApplicationFederatedCredentialWithContext(ctx context.Context, request *GetApplicationFederatedCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationFederatedCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationFederatedCredentialId) {
		query["ApplicationFederatedCredentialId"] = request.ApplicationFederatedCredentialId
	}

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
		Action:      dara.String("GetApplicationFederatedCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationFederatedCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Developer API authorization scope of an EIAM application.
//
// @param request - GetApplicationGrantScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationGrantScopeResponse
func (client *Client) GetApplicationGrantScopeWithContext(ctx context.Context, request *GetApplicationGrantScopeRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationGrantScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the account synchronization configuration for an Entity Identity and Access Management (EIAM) application.
//
// @param request - GetApplicationProvisioningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningConfigResponse
func (client *Client) GetApplicationProvisioningConfigWithContext(ctx context.Context, request *GetApplicationProvisioningConfigRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationProvisioningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScopeWithContext(ctx context.Context, request *GetApplicationProvisioningScopeRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the synchronization configuration of a specified application.
//
// @param request - GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse
func (client *Client) GetApplicationProvisioningUserPrimaryOrganizationalUnitWithContext(ctx context.Context, request *GetApplicationProvisioningUserPrimaryOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetApplicationProvisioningUserPrimaryOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationProvisioningUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves application role information.
//
// @param request - GetApplicationRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationRoleResponse
func (client *Client) GetApplicationRoleWithContext(ctx context.Context, request *GetApplicationRoleRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the single sign-on (SSO) configuration for an application in EIAM.
//
// @param request - GetApplicationSsoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationSsoConfigResponse
func (client *Client) GetApplicationSsoConfigWithContext(ctx context.Context, request *GetApplicationSsoConfigRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationSsoConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves application template information.
//
// @param request - GetApplicationTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationTemplateResponse
func (client *Client) GetApplicationTemplateWithContext(ctx context.Context, request *GetApplicationTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationTemplateId) {
		query["ApplicationTemplateId"] = request.ApplicationTemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationTemplate"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an authorized resource.
//
// @param request - GetAuthorizationResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorizationResourceResponse
func (client *Client) GetAuthorizationResourceWithContext(ctx context.Context, request *GetAuthorizationResourceRequest, runtime *dara.RuntimeOptions) (_result *GetAuthorizationResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceId) {
		query["AuthorizationResourceId"] = request.AuthorizationResourceId
	}

	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthorizationResource"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthorizationResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query information about an authorization rule.
//
// @param request - GetAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorizationRuleResponse
func (client *Client) GetAuthorizationRuleWithContext(ctx context.Context, request *GetAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an authorization server.
//
// @param request - GetAuthorizationServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorizationServerResponse
func (client *Client) GetAuthorizationServerWithContext(ctx context.Context, request *GetAuthorizationServerRequest, runtime *dara.RuntimeOptions) (_result *GetAuthorizationServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationServerId) {
		query["AuthorizationServerId"] = request.AuthorizationServerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthorizationServer"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthorizationServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Brand Details
//
// @param request - GetBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBrandResponse
func (client *Client) GetBrandWithContext(ctx context.Context, request *GetBrandRequest, runtime *dara.RuntimeOptions) (_result *GetBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ClientPublicKey of a specified application.
//
// @param request - GetClientPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientPublicKeyResponse
func (client *Client) GetClientPublicKeyWithContext(ctx context.Context, request *GetClientPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *GetClientPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientPublicKeyId) {
		query["ClientPublicKeyId"] = request.ClientPublicKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClientPublicKey"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource information of a cloud account.
//
// @param request - GetCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudAccountResponse
func (client *Client) GetCloudAccountWithContext(ctx context.Context, request *GetCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *GetCloudAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCloudAccount"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCloudAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource information of a cloud role.
//
// @param request - GetCloudAccountRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudAccountRoleResponse
func (client *Client) GetCloudAccountRoleWithContext(ctx context.Context, request *GetCloudAccountRoleRequest, runtime *dara.RuntimeOptions) (_result *GetCloudAccountRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.CloudAccountRoleId) {
		query["CloudAccountRoleId"] = request.CloudAccountRoleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCloudAccountRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCloudAccountRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get a conditional access policy.
//
// Description:
//
// Query a conditional access policy.
//
// @param request - GetConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConditionalAccessPolicyResponse
func (client *Client) GetConditionalAccessPolicyWithContext(ctx context.Context, request *GetConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetConditionalAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a credential resource.
//
// @param request - GetCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialResponse
func (client *Client) GetCredentialWithContext(ctx context.Context, request *GetCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a credential provider.
//
// @param request - GetCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialProviderResponse
func (client *Client) GetCredentialProviderWithContext(ctx context.Context, request *GetCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *GetCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialProviderId) {
		query["CredentialProviderId"] = request.CredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information about an extended field.
//
// @param request - GetCustomFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomFieldResponse
func (client *Client) GetCustomFieldWithContext(ctx context.Context, request *GetCustomFieldRequest, runtime *dara.RuntimeOptions) (_result *GetCustomFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FieldId) {
		query["FieldId"] = request.FieldId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomField"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a custom privacy policy.
//
// @param request - GetCustomPrivacyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomPrivacyPolicyResponse
func (client *Client) GetCustomPrivacyPolicyWithContext(ctx context.Context, request *GetCustomPrivacyPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetCustomPrivacyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrivacyPolicyId) {
		query["CustomPrivacyPolicyId"] = request.CustomPrivacyPolicyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomPrivacyPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomPrivacyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a domain name for an EIAM instance.
//
// @param request - GetDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithContext(ctx context.Context, request *GetDomainRequest, runtime *dara.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the DNS Challenge record for a specified EIAM domain name. This record is used to verify domain ownership.
//
// @param request - GetDomainDnsChallengeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainDnsChallengeResponse
func (client *Client) GetDomainDnsChallengeWithContext(ctx context.Context, request *GetDomainDnsChallengeRequest, runtime *dara.RuntimeOptions) (_result *GetDomainDnsChallengeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve a federated credential provider.
//
// @param request - GetFederatedCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFederatedCredentialProviderResponse
func (client *Client) GetFederatedCredentialProviderWithContext(ctx context.Context, request *GetFederatedCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *GetFederatedCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederatedCredentialProviderId) {
		query["FederatedCredentialProviderId"] = request.FederatedCredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFederatedCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFederatedCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the forgot password policy of a specified EIAM instance.
//
// @param request - GetForgetPasswordConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetForgetPasswordConfigurationResponse
func (client *Client) GetForgetPasswordConfigurationWithContext(ctx context.Context, request *GetForgetPasswordConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetForgetPasswordConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an EIAM account group.
//
// @param request - GetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithContext(ctx context.Context, request *GetGroupRequest, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an identity provider.
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves advanced configuration information.
//
// @param request - GetIdentityProviderAdvancedConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderAdvancedConfigurationResponse
func (client *Client) GetIdentityProviderAdvancedConfigurationWithContext(ctx context.Context, request *GetIdentityProviderAdvancedConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderAdvancedConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetIdentityProviderAdvancedConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderAdvancedConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an IdP check task.
//
// @param request - GetIdentityProviderStatusCheckJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderStatusCheckJobResponse
func (client *Client) GetIdentityProviderStatusCheckJobWithContext(ctx context.Context, request *GetIdentityProviderStatusCheckJobRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderStatusCheckJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.IdentityProviderStatusCheckJobId) {
		query["IdentityProviderStatusCheckJobId"] = request.IdentityProviderStatusCheckJobId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityProviderStatusCheckJob"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderStatusCheckJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the inbound synchronization configuration of an identity provider (IdP).
//
// @param request - GetIdentityProviderUdPullConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderUdPullConfigurationResponse
func (client *Client) GetIdentityProviderUdPullConfigurationWithContext(ctx context.Context, request *GetIdentityProviderUdPullConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderUdPullConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the IdP outbound synchronization configuration.
//
// @param request - GetIdentityProviderUdPushConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderUdPushConfigurationResponse
func (client *Client) GetIdentityProviderUdPushConfigurationWithContext(ctx context.Context, request *GetIdentityProviderUdPushConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderUdPushConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetIdentityProviderUdPushConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderUdPushConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an EIAM instance.
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instance control items.
//
// @param request - GetInstanceControlConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceControlConfigurationResponse
func (client *Client) GetInstanceControlConfigurationWithContext(ctx context.Context, request *GetInstanceControlConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceControlConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ElementName) {
		query["ElementName"] = request.ElementName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceControlConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceControlConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the language and time zone information of an instance.
//
// @param request - GetInstanceGlobalizationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceGlobalizationConfigResponse
func (client *Client) GetInstanceGlobalizationConfigWithContext(ctx context.Context, request *GetInstanceGlobalizationConfigRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceGlobalizationConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceGlobalizationConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceGlobalizationConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the active license information for an instance.
//
// Description:
//
// Ensure the instance is not in use before deletion. Deleting an EIAM instance permanently removes all of its associated data.
//
// @param request - GetInstanceLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceLicenseResponse
func (client *Client) GetInstanceLicenseWithContext(ctx context.Context, request *GetInstanceLicenseRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all module information under a first-level module.
//
// @param request - GetInstanceModuleInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceModuleInfoResponse
func (client *Client) GetInstanceModuleInfoWithContext(ctx context.Context, request *GetInstanceModuleInfoRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceModuleInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleKey) {
		query["ModuleKey"] = request.ModuleKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceModuleInfo"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceModuleInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the quota of a single type for an instance.
//
// @param request - GetInstanceQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceQuotaResponse
func (client *Client) GetInstanceQuotaWithContext(ctx context.Context, request *GetInstanceQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.QuotaKey) {
		query["QuotaKey"] = request.QuotaKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceQuota"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the trial status of an instance.
//
// @param request - GetInstanceTrialStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceTrialStatusResponse
func (client *Client) GetInstanceTrialStatusWithContext(ctx context.Context, request *GetInstanceTrialStatusRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceTrialStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceTrialStatus"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceTrialStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the post-logon redirect application for a brand.
//
// @param request - GetLoginRedirectApplicationForBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginRedirectApplicationForBrandResponse
func (client *Client) GetLoginRedirectApplicationForBrandWithContext(ctx context.Context, request *GetLoginRedirectApplicationForBrandRequest, runtime *dara.RuntimeOptions) (_result *GetLoginRedirectApplicationForBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoginRedirectApplicationForBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoginRedirectApplicationForBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a specified network access endpoint.
//
// @param request - GetNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkAccessEndpointResponse
func (client *Client) GetNetworkAccessEndpointWithContext(ctx context.Context, request *GetNetworkAccessEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkAccessEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a network zone object.
//
// @param request - GetNetworkZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkZoneResponse
func (client *Client) GetNetworkZoneWithContext(ctx context.Context, request *GetNetworkZoneRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetNetworkZone"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an EIAM organizational unit.
//
// @param request - GetOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnitWithContext(ctx context.Context, request *GetOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *GetOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the password complexity policy for a specified EIAM instance.
//
// @param request - GetPasswordComplexityConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordComplexityConfigurationResponse
func (client *Client) GetPasswordComplexityConfigurationWithContext(ctx context.Context, request *GetPasswordComplexityConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordComplexityConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password expiration policy of a specified EIAM instance.
//
// @param request - GetPasswordExpirationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordExpirationConfigurationResponse
func (client *Client) GetPasswordExpirationConfigurationWithContext(ctx context.Context, request *GetPasswordExpirationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordExpirationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password history policy of a specified EIAM instance.
//
// @param request - GetPasswordHistoryConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordHistoryConfigurationResponse
func (client *Client) GetPasswordHistoryConfigurationWithContext(ctx context.Context, request *GetPasswordHistoryConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordHistoryConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password initialization policy of Employee Identity and Access Management (EIAM).
//
// @param request - GetPasswordInitializationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordInitializationConfigurationResponse
func (client *Client) GetPasswordInitializationConfigurationWithContext(ctx context.Context, request *GetPasswordInitializationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetPasswordInitializationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scope permissions under a specified ResourceServer.
//
// @param request - GetResourceServerScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceServerScopeResponse
func (client *Client) GetResourceServerScopeWithContext(ctx context.Context, request *GetResourceServerScopeRequest, runtime *dara.RuntimeOptions) (_result *GetResourceServerScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerScopeId) {
		query["ResourceServerScopeId"] = request.ResourceServerScopeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceServerScope"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceServerScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the root organizational unit information of EIAM.
//
// @param request - GetRootOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRootOrganizationalUnitResponse
func (client *Client) GetRootOrganizationalUnitWithContext(ctx context.Context, request *GetRootOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *GetRootOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the service quota.
//
// @param request - GetServiceQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceQuotaResponse
func (client *Client) GetServiceQuotaWithContext(ctx context.Context, request *GetServiceQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetServiceQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QuotaType) {
		query["QuotaType"] = request.QuotaType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceQuota"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSynchronizationJobResponse
func (client *Client) GetSynchronizationJobWithContext(ctx context.Context, request *GetSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *GetSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an EIAM account.
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View the list of invocation events.
//
// @param request - ListActionTrackEventTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionTrackEventTypesResponse
func (client *Client) ListActionTrackEventTypesWithContext(ctx context.Context, request *ListActionTrackEventTypesRequest, runtime *dara.RuntimeOptions) (_result *ListActionTrackEventTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("ListActionTrackEventTypes"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListActionTrackEventTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the application accounts of an application by paging.
//
// @param request - ListApplicationAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationAccountsResponse
func (client *Client) ListApplicationAccountsWithContext(ctx context.Context, request *ListApplicationAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationAccounts"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all accounts that belong to a specified user in an application.
//
// Description:
//
// This operation queries only applications that are directly assigned to an organization. You can use the **ApplicationIds*	- parameter to filter the applications.
//
// @param request - ListApplicationAccountsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationAccountsForUserResponse
func (client *Client) ListApplicationAccountsForUserWithContext(ctx context.Context, request *ListApplicationAccountsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationAccountsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
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
		Action:      dara.String("ListApplicationAccountsForUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationAccountsForUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all client secrets for an EIAM application. The key data in the response is masked. To obtain an unmasked key, call the ObtainApplicationClientSecret operation.
//
// @param request - ListApplicationClientSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationClientSecretsResponse
func (client *Client) ListApplicationClientSecretsWithContext(ctx context.Context, request *ListApplicationClientSecretsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationClientSecretsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of application federated credentials.
//
// @param request - ListApplicationFederatedCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationFederatedCredentialsResponse
func (client *Client) ListApplicationFederatedCredentialsWithContext(ctx context.Context, request *ListApplicationFederatedCredentialsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationFederatedCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationFederatedCredentialType) {
		query["ApplicationFederatedCredentialType"] = request.ApplicationFederatedCredentialType
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
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

	if !dara.IsNil(request.PreviousToken) {
		query["PreviousToken"] = request.PreviousToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationFederatedCredentials"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationFederatedCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of application federated credentials by federated trust source ID.
//
// @param request - ListApplicationFederatedCredentialsForProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationFederatedCredentialsForProviderResponse
func (client *Client) ListApplicationFederatedCredentialsForProviderWithContext(ctx context.Context, request *ListApplicationFederatedCredentialsForProviderRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationFederatedCredentialsForProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederatedCredentialProviderId) {
		query["FederatedCredentialProviderId"] = request.FederatedCredentialProviderId
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

	if !dara.IsNil(request.PreviousToken) {
		query["PreviousToken"] = request.PreviousToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationFederatedCredentialsForProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationFederatedCredentialsForProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of application roles using a cursor.
//
// @param request - ListApplicationRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationRolesResponse
func (client *Client) ListApplicationRolesWithContext(ctx context.Context, request *ListApplicationRolesRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationRoles"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of account synchronization types supported by an application.
//
// @param request - ListApplicationSupportedProvisionProtocolTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationSupportedProvisionProtocolTypesResponse
func (client *Client) ListApplicationSupportedProvisionProtocolTypesWithContext(ctx context.Context, request *ListApplicationSupportedProvisionProtocolTypesRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationSupportedProvisionProtocolTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("ListApplicationSupportedProvisionProtocolTypes"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationSupportedProvisionProtocolTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of application tokens.
//
// @param request - ListApplicationTokensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationTokensResponse
func (client *Client) ListApplicationTokensWithContext(ctx context.Context, request *ListApplicationTokensRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationTokensResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationTokenType) {
		query["ApplicationTokenType"] = request.ApplicationTokenType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationTokens"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationTokensResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query to retrieve information about one or more EIAM applications with paging.
//
// @param request - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithContext(ctx context.Context, request *ListApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationCreationType) {
		query["ApplicationCreationType"] = request.ApplicationCreationType
	}

	if !dara.IsNil(request.ApplicationIdentityType) {
		query["ApplicationIdentityType"] = request.ApplicationIdentityType
	}

	if !dara.IsNil(request.ApplicationIds) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.AuthorizationType) {
		query["AuthorizationType"] = request.AuthorizationType
	}

	if !dara.IsNil(request.CustomFields) {
		query["CustomFields"] = request.CustomFields
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.M2MClientStatus) {
		query["M2MClientStatus"] = request.M2MClientStatus
	}

	if !dara.IsNil(request.ManagedServiceCode) {
		query["ManagedServiceCode"] = request.ManagedServiceCode
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

	if !dara.IsNil(request.ServiceManaged) {
		query["ServiceManaged"] = request.ServiceManaged
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the applications associated with an authorization rule.
//
// @param request - ListApplicationsForAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForAuthorizationRuleResponse
func (client *Client) ListApplicationsForAuthorizationRuleWithContext(ctx context.Context, request *ListApplicationsForAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of applications accessible to an EIAM group.
//
// @param request - ListApplicationsForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForGroupResponse
func (client *Client) ListApplicationsForGroupWithContext(ctx context.Context, request *ListApplicationsForGroupRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		query["ApplicationIds"] = request.ApplicationIds
	}

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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForGroup"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves application information under a network access endpoint.
//
// @param request - ListApplicationsForNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForNetworkAccessEndpointResponse
func (client *Client) ListApplicationsForNetworkAccessEndpointWithContext(ctx context.Context, request *ListApplicationsForNetworkAccessEndpointRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForNetworkAccessEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForNetworkAccessEndpoint"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForNetworkAccessEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of applications associated with a network domain.
//
// @param request - ListApplicationsForNetworkZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForNetworkZoneResponse
func (client *Client) ListApplicationsForNetworkZoneWithContext(ctx context.Context, request *ListApplicationsForNetworkZoneRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForNetworkZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkZoneId) {
		query["NetworkZoneId"] = request.NetworkZoneId
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
		Action:      dara.String("ListApplicationsForNetworkZone"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForNetworkZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of applications accessible to an EIAM organizational unit by paging. The response includes application IDs. To obtain detailed application information, call the GetApplication operation.
//
// Description:
//
// This operation queries only the direct permissions of the organizational unit, that is, applications directly assigned to the organizational unit. When you call this operation, you can use the **ApplicationIds*	- parameter to filter applications.
//
// @param request - ListApplicationsForOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForOrganizationalUnitResponse
func (client *Client) ListApplicationsForOrganizationalUnitWithContext(ctx context.Context, request *ListApplicationsForOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that an EIAM account can access and returns a paginated list of application IDs. To retrieve detailed information about a specific application, call the GetApplication operation.
//
// @param request - ListApplicationsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForUserResponse
func (client *Client) ListApplicationsForUserWithContext(ctx context.Context, request *ListApplicationsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of authorization resource information.
//
// @param request - ListAuthorizationResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizationResourcesResponse
func (client *Client) ListAuthorizationResourcesWithContext(ctx context.Context, request *ListAuthorizationResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizationResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizationResources"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizationResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists authorization rules.
//
// @param request - ListAuthorizationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizationRulesResponse
func (client *Client) ListAuthorizationRulesWithContext(ctx context.Context, request *ListAuthorizationRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizationRulesResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizationRules"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizationRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the authorization rules associated with an application.
//
// @param request - ListAuthorizationRulesForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizationRulesForApplicationResponse
func (client *Client) ListAuthorizationRulesForApplicationWithContext(ctx context.Context, request *ListAuthorizationRulesForApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizationRulesForApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizationRulesForApplication"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizationRulesForApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the authorization rules associated with a group.
//
// @param request - ListAuthorizationRulesForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizationRulesForGroupResponse
func (client *Client) ListAuthorizationRulesForGroupWithContext(ctx context.Context, request *ListAuthorizationRulesForGroupRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizationRulesForGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizationRulesForGroup"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizationRulesForGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the authorization rules associated with an account.
//
// @param request - ListAuthorizationRulesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizationRulesForUserResponse
func (client *Client) ListAuthorizationRulesForUserWithContext(ctx context.Context, request *ListAuthorizationRulesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizationRulesForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizationRulesForUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizationRulesForUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of all authorization servers under an instance.
//
// @param request - ListAuthorizationServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizationServersResponse
func (client *Client) ListAuthorizationServersWithContext(ctx context.Context, request *ListAuthorizationServersRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizationServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizationServers"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizationServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of brands.
//
// @param request - ListBrandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBrandsResponse
func (client *Client) ListBrandsWithContext(ctx context.Context, request *ListBrandsRequest, runtime *dara.RuntimeOptions) (_result *ListBrandsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("ListBrands"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBrandsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the client public keys for a specified application using a cursor.
//
// @param request - ListClientPublicKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientPublicKeysResponse
func (client *Client) ListClientPublicKeysWithContext(ctx context.Context, request *ListClientPublicKeysRequest, runtime *dara.RuntimeOptions) (_result *ListClientPublicKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClientPublicKeys"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientPublicKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more cloud roles by using paging.
//
// @param request - ListCloudAccountRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAccountRolesResponse
func (client *Client) ListCloudAccountRolesWithContext(ctx context.Context, request *ListCloudAccountRolesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAccountRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudAccountRoles"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudAccountRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about one or more cloud accounts by using paging.
//
// @param request - ListCloudAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudAccountsResponse
func (client *Client) ListCloudAccountsWithContext(ctx context.Context, request *ListCloudAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudAccountsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudAccounts"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List conditional access policies.
//
// Description:
//
// Query the list of conditional access policies with pagination.
//
// @param request - ListConditionalAccessPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConditionalAccessPoliciesResponse
func (client *Client) ListConditionalAccessPoliciesWithContext(ctx context.Context, request *ListConditionalAccessPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListConditionalAccessPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List conditional access policies associated with an application.
//
// @param request - ListConditionalAccessPoliciesForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConditionalAccessPoliciesForApplicationResponse
func (client *Client) ListConditionalAccessPoliciesForApplicationWithContext(ctx context.Context, request *ListConditionalAccessPoliciesForApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListConditionalAccessPoliciesForApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("ListConditionalAccessPoliciesForApplication"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConditionalAccessPoliciesForApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of conditional access policies associated with a network zone.
//
// Description:
//
// Retrieves the list of conditional access policies associated with a network zone.
//
// @param request - ListConditionalAccessPoliciesForNetworkZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConditionalAccessPoliciesForNetworkZoneResponse
func (client *Client) ListConditionalAccessPoliciesForNetworkZoneWithContext(ctx context.Context, request *ListConditionalAccessPoliciesForNetworkZoneRequest, runtime *dara.RuntimeOptions) (_result *ListConditionalAccessPoliciesForNetworkZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of conditional access policies associated with a user.
//
// @param request - ListConditionalAccessPoliciesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConditionalAccessPoliciesForUserResponse
func (client *Client) ListConditionalAccessPoliciesForUserWithContext(ctx context.Context, request *ListConditionalAccessPoliciesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListConditionalAccessPoliciesForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("ListConditionalAccessPoliciesForUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConditionalAccessPoliciesForUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists credential providers.
//
// @param request - ListCredentialProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCredentialProvidersResponse
func (client *Client) ListCredentialProvidersWithContext(ctx context.Context, request *ListCredentialProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListCredentialProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialProviderIds) {
		query["CredentialProviderIds"] = request.CredentialProviderIds
	}

	if !dara.IsNil(request.CredentialProviderTypes) {
		query["CredentialProviderTypes"] = request.CredentialProviderTypes
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	if !dara.IsNil(request.Statuses) {
		query["Statuses"] = request.Statuses
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCredentialProviders"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCredentialProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of one or more credentials by using paging.
//
// @param request - ListCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCredentialsResponse
func (client *Client) ListCredentialsWithContext(ctx context.Context, request *ListCredentialsRequest, runtime *dara.RuntimeOptions) (_result *ListCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialExternalIds) {
		query["CredentialExternalIds"] = request.CredentialExternalIds
	}

	if !dara.IsNil(request.CredentialIds) {
		query["CredentialIds"] = request.CredentialIds
	}

	if !dara.IsNil(request.CredentialSharingScopes) {
		query["CredentialSharingScopes"] = request.CredentialSharingScopes
	}

	if !dara.IsNil(request.CredentialTypes) {
		query["CredentialTypes"] = request.CredentialTypes
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	if !dara.IsNil(request.Statuses) {
		query["Statuses"] = request.Statuses
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCredentials"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of custom terms.
//
// @param request - ListCustomPrivacyPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomPrivacyPoliciesResponse
func (client *Client) ListCustomPrivacyPoliciesWithContext(ctx context.Context, request *ListCustomPrivacyPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListCustomPrivacyPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrivacyPolicyNameStartsWith) {
		query["CustomPrivacyPolicyNameStartsWith"] = request.CustomPrivacyPolicyNameStartsWith
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

	if !dara.IsNil(request.PreviousToken) {
		query["PreviousToken"] = request.PreviousToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomPrivacyPolicies"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomPrivacyPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the resources of brand-linked instances.
//
// @param request - ListCustomPrivacyPoliciesForBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomPrivacyPoliciesForBrandResponse
func (client *Client) ListCustomPrivacyPoliciesForBrandWithContext(ctx context.Context, request *ListCustomPrivacyPoliciesForBrandRequest, runtime *dara.RuntimeOptions) (_result *ListCustomPrivacyPoliciesForBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
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

	if !dara.IsNil(request.PreviousToken) {
		query["PreviousToken"] = request.PreviousToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomPrivacyPoliciesForBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomPrivacyPoliciesForBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of domain proxy tokens of an EIAM instance.
//
// @param request - ListDomainProxyTokensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainProxyTokensResponse
func (client *Client) ListDomainProxyTokensWithContext(ctx context.Context, request *ListDomainProxyTokensRequest, runtime *dara.RuntimeOptions) (_result *ListDomainProxyTokensResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names of an EIAM instance, including the default domain name and custom domain names.
//
// @param request - ListDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithContext(ctx context.Context, request *ListDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of EIAM 2.0 and EIAM 1.0 instances.
//
// @param request - ListEiamInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEiamInstancesResponse
func (client *Client) ListEiamInstancesWithContext(ctx context.Context, request *ListEiamInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListEiamInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the event list.
//
// @param request - ListEventTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventTypesResponse
func (client *Client) ListEventTypesWithContext(ctx context.Context, request *ListEventTypesRequest, runtime *dara.RuntimeOptions) (_result *ListEventTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("ListEventTypes"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEventTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of federated trust sources.
//
// @param request - ListFederatedCredentialProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFederatedCredentialProvidersResponse
func (client *Client) ListFederatedCredentialProvidersWithContext(ctx context.Context, request *ListFederatedCredentialProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListFederatedCredentialProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederatedCredentialProviderName) {
		query["FederatedCredentialProviderName"] = request.FederatedCredentialProviderName
	}

	if !dara.IsNil(request.FederatedCredentialProviderType) {
		query["FederatedCredentialProviderType"] = request.FederatedCredentialProviderType
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

	if !dara.IsNil(request.PreviousToken) {
		query["PreviousToken"] = request.PreviousToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFederatedCredentialProviders"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFederatedCredentialProvidersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List EIAM account groups.
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithContext(ctx context.Context, request *ListGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query to list the groups authorized to access an application. The response returns the group IDs. To obtain detailed information for a group, you can call the GetGroup operation.
//
// @param request - ListGroupsForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForApplicationResponse
func (client *Client) ListGroupsForApplicationWithContext(ctx context.Context, request *ListGroupsForApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsForApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the groups associated with an authorization rule.
//
// @param request - ListGroupsForAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForAuthorizationRuleResponse
func (client *Client) ListGroupsForAuthorizationRuleWithContext(ctx context.Context, request *ListGroupsForAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsForAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupsForAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsForAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the scopes authorized for groups on a specified resource server. This operation supports cursor-based pagination.
//
// @param request - ListGroupsForResourceServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForResourceServerResponse
func (client *Client) ListGroupsForResourceServerWithContext(ctx context.Context, request *ListGroupsForResourceServerRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsForResourceServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	if !dara.IsNil(request.ResourceServerScopeId) {
		query["ResourceServerScopeId"] = request.ResourceServerScopeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupsForResourceServer"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsForResourceServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of account groups to which a specified EIAM account belongs.
//
// @param request - ListGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithContext(ctx context.Context, request *ListGroupsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about Identity Providers (IdPs) for a network endpoint.
//
// @param request - ListIdentityProvidersForNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProvidersForNetworkAccessEndpointResponse
func (client *Client) ListIdentityProvidersForNetworkAccessEndpointWithContext(ctx context.Context, request *ListIdentityProvidersForNetworkAccessEndpointRequest, runtime *dara.RuntimeOptions) (_result *ListIdentityProvidersForNetworkAccessEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentityProvidersForNetworkAccessEndpoint"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityProvidersForNetworkAccessEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about one or more EIAM instances.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossRegionReplication) {
		query["CrossRegionReplication"] = request.CrossRegionReplication
	}

	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of zones that support the creation of network endpoints in a specified region for IDaaS EIAM.
//
// @param request - ListNetworkAccessEndpointAvailableZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessEndpointAvailableZonesResponse
func (client *Client) ListNetworkAccessEndpointAvailableZonesWithContext(ctx context.Context, request *ListNetworkAccessEndpointAvailableZonesRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkAccessEndpointAvailableZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the network endpoints for an IDaaS EIAM instance.
//
// @param request - ListNetworkAccessEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessEndpointsResponse
func (client *Client) ListNetworkAccessEndpointsWithContext(ctx context.Context, request *ListNetworkAccessEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkAccessEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of access paths under a specified network access endpoint.
//
// @param request - ListNetworkAccessPathsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessPathsResponse
func (client *Client) ListNetworkAccessPathsWithContext(ctx context.Context, request *ListNetworkAccessPathsRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkAccessPathsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of network zone objects.
//
// @param request - ListNetworkZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkZonesResponse
func (client *Client) ListNetworkZonesWithContext(ctx context.Context, request *ListNetworkZonesRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkZoneIds) {
		query["NetworkZoneIds"] = request.NetworkZoneIds
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
		Action:      dara.String("ListNetworkZones"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all ancestor organizations of a specified EIAM organization. The organizations in the result list are sorted in hierarchical order from the top level to the bottom level.
//
// @param request - ListOrganizationalUnitParentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitParentsResponse
func (client *Client) ListOrganizationalUnitParentsWithContext(ctx context.Context, request *ListOrganizationalUnitParentsRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitParentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query for EIAM organizational units.
//
// @param request - ListOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnitsWithContext(ctx context.Context, request *ListOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query to list the organizations that are granted access to an application. The response returns the IDs of the organizations. To obtain detailed information about a specific organization, call the GetOrganizationalUnit operation.
//
// @param request - ListOrganizationalUnitsForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsForApplicationResponse
func (client *Client) ListOrganizationalUnitsForApplicationWithContext(ctx context.Context, request *ListOrganizationalUnitsForApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitsForApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a cursor-paginated list of scopes that the current resource server has granted to an organization.
//
// @param request - ListOrganizationalUnitsForResourceServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsForResourceServerResponse
func (client *Client) ListOrganizationalUnitsForResourceServerWithContext(ctx context.Context, request *ListOrganizationalUnitsForResourceServerRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitsForResourceServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	if !dara.IsNil(request.ResourceServerScopeId) {
		query["ResourceServerScopeId"] = request.ResourceServerScopeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationalUnitsForResourceServer"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationalUnitsForResourceServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of scope permissions under a specified ResourceServer by using a cursor.
//
// @param request - ListResourceServerScopesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceServerScopesResponse
func (client *Client) ListResourceServerScopesWithContext(ctx context.Context, request *ListResourceServerScopesRequest, runtime *dara.RuntimeOptions) (_result *ListResourceServerScopesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PreviousToken) {
		query["PreviousToken"] = request.PreviousToken
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	if !dara.IsNil(request.ResourceServerScopeName) {
		query["ResourceServerScopeName"] = request.ResourceServerScopeName
	}

	if !dara.IsNil(request.ResourceServerScopeType) {
		query["ResourceServerScopeType"] = request.ResourceServerScopeType
	}

	if !dara.IsNil(request.ResourceServerScopeValue) {
		query["ResourceServerScopeValue"] = request.ResourceServerScopeValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceServerScopes"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceServerScopesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of ResourceServer and Scope permissions granted to the current account by using a cursor-based approach.
//
// @param request - ListResourceServersForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceServersForUserResponse
func (client *Client) ListResourceServersForUserWithContext(ctx context.Context, request *ListResourceServersForUserRequest, runtime *dara.RuntimeOptions) (_result *ListResourceServersForUserResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceServersForUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceServersForUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of synchronization job details.
//
// Description:
//
// The following two methods are supported for viewing returned data:
//
// - Method 1: When querying the first page, you only need to set MaxResults to limit the number of entries returned. The NextToken in the response serves as the credential for querying subsequent pages. When querying subsequent pages, set the NextToken parameter to the NextToken value obtained from the previous response as the query credential, and set MaxResults to limit the number of entries returned. If there are no more pages, NextToken is no longer returned. The maximum value of MaxResults is 100.
//
// - Method 2: Set the number of entries returned per page using PageSize, and set the page number using PageNumber.
//
// You can only choose one of the above two methods. When a large number of entries are returned, Method 1 is recommended. If the MaxResults or NextToken parameter is set, the PageSize and PageNumber request parameters become invalid.
//
// @param request - ListSynchronizationJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSynchronizationJobsResponse
func (client *Client) ListSynchronizationJobsWithContext(ctx context.Context, request *ListSynchronizationJobsRequest, runtime *dara.RuntimeOptions) (_result *ListSynchronizationJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the mappings for third-party logon accounts.
//
// Description:
//
// This operation queries only the applications that are directly assigned to an organization. You can use the **ApplicationIds*	- parameter to filter the applications.
//
// @param request - ListUserAuthnSourceMappingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserAuthnSourceMappingsResponse
func (client *Client) ListUserAuthnSourceMappingsWithContext(ctx context.Context, request *ListUserAuthnSourceMappingsRequest, runtime *dara.RuntimeOptions) (_result *ListUserAuthnSourceMappingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
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

	if !dara.IsNil(request.PreviousToken) {
		query["PreviousToken"] = request.PreviousToken
	}

	if !dara.IsNil(request.UserExternalId) {
		query["UserExternalId"] = request.UserExternalId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserAuthnSourceMappings"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserAuthnSourceMappingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries EIAM account information by using paging.
//
// Description:
//
// This operation queries only the direct permissions of an organizational unit, which are the applications directly assigned to the organizational unit. When you call this operation, you can use the **ApplicationIds*	- parameter to filter applications.
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query to list the accounts that have been granted access to an application. The response includes account IDs. To retrieve detailed information about an account, call the GetUser operation.
//
// @param request - ListUsersForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForApplicationResponse
func (client *Client) ListUsersForApplicationWithContext(ctx context.Context, request *ListUsersForApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListUsersForApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the accounts associated with an authorization rule.
//
// @param request - ListUsersForAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForAuthorizationRuleResponse
func (client *Client) ListUsersForAuthorizationRuleWithContext(ctx context.Context, request *ListUsersForAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *ListUsersForAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsersForAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersForAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the users in a specified EIAM account group.
//
// @param request - ListUsersForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithContext(ctx context.Context, request *ListUsersForGroupRequest, runtime *dara.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List the scope permissions granted by a Resource Server to user accounts using cursor-based pagination.
//
// @param request - ListUsersForResourceServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForResourceServerResponse
func (client *Client) ListUsersForResourceServerWithContext(ctx context.Context, request *ListUsersForResourceServerRequest, runtime *dara.RuntimeOptions) (_result *ListUsersForResourceServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
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

	if !dara.IsNil(request.ResourceServerScopeId) {
		query["ResourceServerScopeId"] = request.ResourceServerScopeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsersForResourceServer"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersForResourceServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the client secret of an EIAM application. The returned secret is not masked. To retrieve masked secrets, call the ListApplicationClientSecrets operation.
//
// @param request - ObtainApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainApplicationClientSecretResponse
func (client *Client) ObtainApplicationClientSecretWithContext(ctx context.Context, request *ObtainApplicationClientSecretRequest, runtime *dara.RuntimeOptions) (_result *ObtainApplicationClientSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specified application token.
//
// Description:
//
// When you change an application from the enabled state to the disabled state, all features of the application become unavailable, such as SSO and account synchronization. Make sure that you are aware of the risks that may result from this operation.
//
// @param request - ObtainApplicationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainApplicationTokenResponse
func (client *Client) ObtainApplicationTokenWithContext(ctx context.Context, request *ObtainApplicationTokenRequest, runtime *dara.RuntimeOptions) (_result *ObtainApplicationTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationTokenId) {
		query["ApplicationTokenId"] = request.ApplicationTokenId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ObtainApplicationToken"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ObtainApplicationTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a credential resource that contains sensitive information.
//
// @param request - ObtainCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainCredentialResponse
func (client *Client) ObtainCredentialWithContext(ctx context.Context, request *ObtainCredentialRequest, runtime *dara.RuntimeOptions) (_result *ObtainCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ObtainCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ObtainCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proxy token of an EIAM instance domain name.
//
// @param request - ObtainDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainDomainProxyTokenResponse
func (client *Client) ObtainDomainProxyTokenWithContext(ctx context.Context, request *ObtainDomainProxyTokenRequest, runtime *dara.RuntimeOptions) (_result *ObtainDomainProxyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application account of a specified employee under the current application.
//
// Description:
//
// This operation only queries the direct permissions of the organization, that is, applications directly assigned to the organization. When you call this operation, you can use the **ApplicationIds*	- parameter to filter applications.
//
// @param request - RemoveApplicationAccountFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApplicationAccountFromUserResponse
func (client *Client) RemoveApplicationAccountFromUserWithContext(ctx context.Context, request *RemoveApplicationAccountFromUserRequest, runtime *dara.RuntimeOptions) (_result *RemoveApplicationAccountFromUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationAccountId) {
		query["ApplicationAccountId"] = request.ApplicationAccountId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
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
		Action:      dara.String("RemoveApplicationAccountFromUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApplicationAccountFromUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an application from an authorization rule.
//
// @param request - RemoveApplicationFromAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApplicationFromAuthorizationRuleResponse
func (client *Client) RemoveApplicationFromAuthorizationRuleWithContext(ctx context.Context, request *RemoveApplicationFromAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *RemoveApplicationFromAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApplicationFromAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApplicationFromAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes custom terms associated with a brand.
//
// @param request - RemoveCustomPrivacyPoliciesFromBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveCustomPrivacyPoliciesFromBrandResponse
func (client *Client) RemoveCustomPrivacyPoliciesFromBrandWithContext(ctx context.Context, request *RemoveCustomPrivacyPoliciesFromBrandRequest, runtime *dara.RuntimeOptions) (_result *RemoveCustomPrivacyPoliciesFromBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.CustomPrivacyPolicyIds) {
		query["CustomPrivacyPolicyIds"] = request.CustomPrivacyPolicyIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveCustomPrivacyPoliciesFromBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveCustomPrivacyPoliciesFromBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an application from an authorization rule.
//
// @param request - RemoveGroupFromAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveGroupFromAuthorizationRuleResponse
func (client *Client) RemoveGroupFromAuthorizationRuleWithContext(ctx context.Context, request *RemoveGroupFromAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *RemoveGroupFromAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
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
		Action:      dara.String("RemoveGroupFromAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveGroupFromAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an account from an authorization rule.
//
// @param request - RemoveUserFromAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromAuthorizationRuleResponse
func (client *Client) RemoveUserFromAuthorizationRuleWithContext(ctx context.Context, request *RemoveUserFromAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
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
		Action:      dara.String("RemoveUserFromAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserFromAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnitsWithContext(ctx context.Context, request *RemoveUserFromOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroupWithContext(ctx context.Context, request *RemoveUsersFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersFromGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a Free Edition instance.
//
// @param request - RenewFreeLicenseEndTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewFreeLicenseEndTimeResponse
func (client *Client) RenewFreeLicenseEndTimeWithContext(ctx context.Context, request *RenewFreeLicenseEndTimeRequest, runtime *dara.RuntimeOptions) (_result *RenewFreeLicenseEndTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewFreeLicenseEndTime"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewFreeLicenseEndTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions of multiple Employee Identity and Access Management (EIAM) groups to access an application in a batch.
//
// @param request - RevokeApplicationFromGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromGroupsResponse
func (client *Client) RevokeApplicationFromGroupsWithContext(ctx context.Context, request *RevokeApplicationFromGroupsRequest, runtime *dara.RuntimeOptions) (_result *RevokeApplicationFromGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes application access from multiple EIAM organizations in a batch operation.
//
// @param request - RevokeApplicationFromOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromOrganizationalUnitsResponse
func (client *Client) RevokeApplicationFromOrganizationalUnitsWithContext(ctx context.Context, request *RevokeApplicationFromOrganizationalUnitsRequest, runtime *dara.RuntimeOptions) (_result *RevokeApplicationFromOrganizationalUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions of multiple Enterprise Identity Access Management (EIAM) accounts to access an application in a batch.
//
// @param request - RevokeApplicationFromUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromUsersResponse
func (client *Client) RevokeApplicationFromUsersWithContext(ctx context.Context, request *RevokeApplicationFromUsersRequest, runtime *dara.RuntimeOptions) (_result *RevokeApplicationFromUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the authorization granted by a specified ResourceServer to a Client application.
//
// @param request - RevokeResourceServerFromClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourceServerFromClientResponse
func (client *Client) RevokeResourceServerFromClientWithContext(ctx context.Context, request *RevokeResourceServerFromClientRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourceServerFromClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientApplicationId) {
		query["ClientApplicationId"] = request.ClientApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerApplicationId) {
		query["ResourceServerApplicationId"] = request.ResourceServerApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeResourceServerFromClient"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeResourceServerFromClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the scope permissions of a specified ResourceServer from a client application.
//
// @param request - RevokeResourceServerScopesFromClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourceServerScopesFromClientResponse
func (client *Client) RevokeResourceServerScopesFromClientWithContext(ctx context.Context, request *RevokeResourceServerScopesFromClientRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourceServerScopesFromClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientApplicationId) {
		query["ClientApplicationId"] = request.ClientApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerApplicationId) {
		query["ResourceServerApplicationId"] = request.ResourceServerApplicationId
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeResourceServerScopesFromClient"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeResourceServerScopesFromClientResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the authorization of Scope permissions under a specified ResourceServer from a group.
//
// @param request - RevokeResourceServerScopesFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourceServerScopesFromGroupResponse
func (client *Client) RevokeResourceServerScopesFromGroupWithContext(ctx context.Context, request *RevokeResourceServerScopesFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourceServerScopesFromGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeResourceServerScopesFromGroup"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeResourceServerScopesFromGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the authorization of Scope permissions under a specified ResourceServer from an organizational unit.
//
// @param request - RevokeResourceServerScopesFromOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourceServerScopesFromOrganizationalUnitResponse
func (client *Client) RevokeResourceServerScopesFromOrganizationalUnitWithContext(ctx context.Context, request *RevokeResourceServerScopesFromOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourceServerScopesFromOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrganizationalUnitId) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeResourceServerScopesFromOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeResourceServerScopesFromOrganizationalUnitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the Scope permissions under a specified ResourceServer from an account.
//
// @param request - RevokeResourceServerScopesFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourceServerScopesFromUserResponse
func (client *Client) RevokeResourceServerScopesFromUserWithContext(ctx context.Context, request *RevokeResourceServerScopesFromUserRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourceServerScopesFromUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerScopeIds) {
		query["ResourceServerScopeIds"] = request.ResourceServerScopeIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeResourceServerScopesFromUser"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeResourceServerScopesFromUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and immediately runs a new synchronization task.
//
// @param request - RunSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSynchronizationJobResponse
func (client *Client) RunSynchronizationJobWithContext(ctx context.Context, request *RunSynchronizationJobRequest, runtime *dara.RuntimeOptions) (_result *RunSynchronizationJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationGrantScopeResponse
func (client *Client) SetApplicationGrantScopeWithContext(ctx context.Context, request *SetApplicationGrantScopeRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationGrantScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the account synchronization configuration for an EIAM application.
//
// @param request - SetApplicationProvisioningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationProvisioningConfigResponse
func (client *Client) SetApplicationProvisioningConfigWithContext(ctx context.Context, request *SetApplicationProvisioningConfigRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationProvisioningConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationProvisioningScopeResponse
func (client *Client) SetApplicationProvisioningScopeWithContext(ctx context.Context, request *SetApplicationProvisioningScopeRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationProvisioningScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the primary organizational unit for an application\\"s user provisioning.
//
// @param request - SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse
func (client *Client) SetApplicationProvisioningUserPrimaryOrganizationalUnitWithContext(ctx context.Context, request *SetApplicationProvisioningUserPrimaryOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserPrimaryOrganizationalUnitId) {
		query["UserPrimaryOrganizationalUnitId"] = request.UserPrimaryOrganizationalUnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApplicationProvisioningUserPrimaryOrganizationalUnit"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApplicationProvisioningUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the unique identifier for a resource server. This identifier is used as the aud (audience) claim in a JSON Web Token (JWT) to specify the service that is intended to accept the token.
//
// @param request - SetApplicationResourceServerIdentifierRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationResourceServerIdentifierResponse
func (client *Client) SetApplicationResourceServerIdentifierWithContext(ctx context.Context, request *SetApplicationResourceServerIdentifierRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationResourceServerIdentifierResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerIdentifier) {
		query["ResourceServerIdentifier"] = request.ResourceServerIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApplicationResourceServerIdentifier"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApplicationResourceServerIdentifierResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the single sign-on (SSO) configuration properties for an EIAM application.
//
// Description:
//
// In EIAM, application management supports adding applications that use multiple single sign-on (SSO) protocols (SAML 2.0 and OIDC). However, each application supports only one SSO protocol, which is specified during creation and cannot be changed. Specify the corresponding SSO configuration property parameters based on the SSO protocol type supported by the application.
//
// @param request - SetApplicationSsoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationSsoConfigResponse
func (client *Client) SetApplicationSsoConfigWithContext(ctx context.Context, request *SetApplicationSsoConfigRequest, runtime *dara.RuntimeOptions) (_result *SetApplicationSsoConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a specified domain name of an EIAM instance as the default domain name.
//
// @param request - SetDefaultDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomainWithContext(ctx context.Context, request *SetDefaultDomainRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the forgot password policy for an EIAM instance.
//
// @param request - SetForgetPasswordConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetForgetPasswordConfigurationResponse
func (client *Client) SetForgetPasswordConfigurationWithContext(ctx context.Context, request *SetForgetPasswordConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetForgetPasswordConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Authentication Information
//
// @param request - SetIdentityProviderAuthnConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdentityProviderAuthnConfigurationResponse
func (client *Client) SetIdentityProviderAuthnConfigurationWithContext(ctx context.Context, request *SetIdentityProviderAuthnConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetIdentityProviderAuthnConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateUserConfig) {
		query["AutoCreateUserConfig"] = request.AutoCreateUserConfig
	}

	if !dara.IsNil(request.AutoUpdateUserConfig) {
		query["AutoUpdateUserConfig"] = request.AutoUpdateUserConfig
	}

	if !dara.IsNil(request.BindingConfig) {
		query["BindingConfig"] = request.BindingConfig
	}

	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LdapAuthnConfig) {
		query["LdapAuthnConfig"] = request.LdapAuthnConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIdentityProviderAuthnConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIdentityProviderAuthnConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the inbound synchronization configuration of an Identity Provider (IdP).
//
// @param request - SetIdentityProviderUdPullConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdentityProviderUdPullConfigurationResponse
func (client *Client) SetIdentityProviderUdPullConfigurationWithContext(ctx context.Context, request *SetIdentityProviderUdPullConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetIdentityProviderUdPullConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the push configuration for an identity provider (IdP).
//
// @param request - SetIdentityProviderUdPushConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdentityProviderUdPushConfigurationResponse
func (client *Client) SetIdentityProviderUdPushConfigurationWithContext(ctx context.Context, request *SetIdentityProviderUdPushConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetIdentityProviderUdPushConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.IncrementalCallbackStatus) {
		query["IncrementalCallbackStatus"] = request.IncrementalCallbackStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LdapUdPushConfig) {
		query["LdapUdPushConfig"] = request.LdapUdPushConfig
	}

	if !dara.IsNil(request.PeriodicSyncConfig) {
		query["PeriodicSyncConfig"] = request.PeriodicSyncConfig
	}

	if !dara.IsNil(request.PeriodicSyncStatus) {
		query["PeriodicSyncStatus"] = request.PeriodicSyncStatus
	}

	if !dara.IsNil(request.UdSyncScopeConfigs) {
		query["UdSyncScopeConfigs"] = request.UdSyncScopeConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIdentityProviderUdPushConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIdentityProviderUdPushConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the control elements of an instance.
//
// @param request - SetInstanceControlConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetInstanceControlConfigurationResponse
func (client *Client) SetInstanceControlConfigurationWithContext(ctx context.Context, request *SetInstanceControlConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetInstanceControlConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlElements) {
		query["ControlElements"] = request.ControlElements
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetInstanceControlConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetInstanceControlConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the language and time zone information for an instance.
//
// @param request - SetInstanceGlobalizationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetInstanceGlobalizationConfigResponse
func (client *Client) SetInstanceGlobalizationConfigWithContext(ctx context.Context, request *SetInstanceGlobalizationConfigRequest, runtime *dara.RuntimeOptions) (_result *SetInstanceGlobalizationConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.TimeZone) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetInstanceGlobalizationConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetInstanceGlobalizationConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Set login redirect application for brand
//
// @param request - SetLoginRedirectApplicationForBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoginRedirectApplicationForBrandResponse
func (client *Client) SetLoginRedirectApplicationForBrandWithContext(ctx context.Context, request *SetLoginRedirectApplicationForBrandRequest, runtime *dara.RuntimeOptions) (_result *SetLoginRedirectApplicationForBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoginRedirectApplicationForBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoginRedirectApplicationForBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the password complexity policy for a specified EIAM instance.
//
// @param request - SetPasswordComplexityConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordComplexityConfigurationResponse
func (client *Client) SetPasswordComplexityConfigurationWithContext(ctx context.Context, request *SetPasswordComplexityConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordComplexityConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DisabledWeakPasswordLogin) {
		query["DisabledWeakPasswordLogin"] = request.DisabledWeakPasswordLogin
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the password expiration policy for a specified Enterprise Identity and Access Management (EIAM) instance.
//
// @param request - SetPasswordExpirationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordExpirationConfigurationResponse
func (client *Client) SetPasswordExpirationConfigurationWithContext(ctx context.Context, request *SetPasswordExpirationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordExpirationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the password history policy for a specified Enterprise Identity and Access Management (EIAM) instance.
//
// @param request - SetPasswordHistoryConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordHistoryConfigurationResponse
func (client *Client) SetPasswordHistoryConfigurationWithContext(ctx context.Context, request *SetPasswordHistoryConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordHistoryConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the password initialization policy for a specified EIAM instance.
//
// @param request - SetPasswordInitializationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordInitializationConfigurationResponse
func (client *Client) SetPasswordInitializationConfigurationWithContext(ctx context.Context, request *SetPasswordInitializationConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetPasswordInitializationConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the specified client public key as the primary key for an application.
//
// @param request - SetPrimaryClientPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPrimaryClientPublicKeyResponse
func (client *Client) SetPrimaryClientPublicKeyWithContext(ctx context.Context, request *SetPrimaryClientPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *SetPrimaryClientPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientPublicKeyId) {
		query["ClientPublicKeyId"] = request.ClientPublicKeyId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPrimaryClientPublicKey"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPrimaryClientPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnitWithContext(ctx context.Context, request *SetUserPrimaryOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the WebAuthn configuration.
//
// @param request - SetWebAuthnConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWebAuthnConfigurationResponse
func (client *Client) SetWebAuthnConfigurationWithContext(ctx context.Context, request *SetWebAuthnConfigurationRequest, runtime *dara.RuntimeOptions) (_result *SetWebAuthnConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Aaguids) {
		query["Aaguids"] = request.Aaguids
	}

	if !dara.IsNil(request.EnableAaguidVerification) {
		query["EnableAaguidVerification"] = request.EnableAaguidVerification
	}

	if !dara.IsNil(request.EnableMetadataServiceVerification) {
		query["EnableMetadataServiceVerification"] = request.EnableMetadataServiceVerification
	}

	if !dara.IsNil(request.EnableUserSelfRegistration) {
		query["EnableUserSelfRegistration"] = request.EnableUserSelfRegistration
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWebAuthnConfiguration"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWebAuthnConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches the TOTP authenticator for a specified user.
//
// @param request - UnbindTotpAuthenticatorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindTotpAuthenticatorResponse
func (client *Client) UnbindTotpAuthenticatorWithContext(ctx context.Context, request *UnbindTotpAuthenticatorRequest, runtime *dara.RuntimeOptions) (_result *UnbindTotpAuthenticatorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("UnbindTotpAuthenticator"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindTotpAuthenticatorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a third-party logon account from a user.
//
// @param request - UnbindUserAuthnSourceMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindUserAuthnSourceMappingResponse
func (client *Client) UnbindUserAuthnSourceMappingWithContext(ctx context.Context, request *UnbindUserAuthnSourceMappingRequest, runtime *dara.RuntimeOptions) (_result *UnbindUserAuthnSourceMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderId) {
		query["IdentityProviderId"] = request.IdentityProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserExternalId) {
		query["UserExternalId"] = request.UserExternalId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindUserAuthnSourceMapping"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindUserAuthnSourceMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks a locked EIAM account.
//
// @param request - UnlockUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockUserResponse
func (client *Client) UnlockUserWithContext(ctx context.Context, request *UnlockUserRequest, runtime *dara.RuntimeOptions) (_result *UnlockUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the advanced configuration of an application.
//
// @param request - UpdateApplicationAdvancedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationAdvancedConfigResponse
func (client *Client) UpdateApplicationAdvancedConfigWithContext(ctx context.Context, request *UpdateApplicationAdvancedConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationAdvancedConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScimServerAdvancedConfig) {
		query["ScimServerAdvancedConfig"] = request.ScimServerAdvancedConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationAdvancedConfig"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationAdvancedConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationAuthorizationTypeResponse
func (client *Client) UpdateApplicationAuthorizationTypeWithContext(ctx context.Context, request *UpdateApplicationAuthorizationTypeRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationAuthorizationTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the expiration time of a specified ClientSecret for an application.
//
// @param request - UpdateApplicationClientSecretExpirationTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationClientSecretExpirationTimeResponse
func (client *Client) UpdateApplicationClientSecretExpirationTimeWithContext(ctx context.Context, request *UpdateApplicationClientSecretExpirationTimeRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationClientSecretExpirationTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ExpirationTime) {
		query["ExpirationTime"] = request.ExpirationTime
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
		Action:      dara.String("UpdateApplicationClientSecretExpirationTime"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationClientSecretExpirationTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationDescriptionResponse
func (client *Client) UpdateApplicationDescriptionWithContext(ctx context.Context, request *UpdateApplicationDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a federated credential for an application.
//
// @param request - UpdateApplicationFederatedCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationFederatedCredentialResponse
func (client *Client) UpdateApplicationFederatedCredentialWithContext(ctx context.Context, request *UpdateApplicationFederatedCredentialRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationFederatedCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationFederatedCredentialId) {
		query["ApplicationFederatedCredentialId"] = request.ApplicationFederatedCredentialId
	}

	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AttributeMappings) {
		query["AttributeMappings"] = request.AttributeMappings
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OidcVerificationConfig) {
		query["OidcVerificationConfig"] = request.OidcVerificationConfig
	}

	if !dara.IsNil(request.Pkcs7VerificationConfig) {
		query["Pkcs7VerificationConfig"] = request.Pkcs7VerificationConfig
	}

	if !dara.IsNil(request.VerificationCondition) {
		query["VerificationCondition"] = request.VerificationCondition
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationFederatedCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationFederatedCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a federated credential for an application.
//
// @param request - UpdateApplicationFederatedCredentialDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationFederatedCredentialDescriptionResponse
func (client *Client) UpdateApplicationFederatedCredentialDescriptionWithContext(ctx context.Context, request *UpdateApplicationFederatedCredentialDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationFederatedCredentialDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationFederatedCredentialId) {
		query["ApplicationFederatedCredentialId"] = request.ApplicationFederatedCredentialId
	}

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
		Action:      dara.String("UpdateApplicationFederatedCredentialDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationFederatedCredentialDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information of an application.
//
// @param request - UpdateApplicationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationInfoResponse
func (client *Client) UpdateApplicationInfoWithContext(ctx context.Context, request *UpdateApplicationInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ApplicationOwner) {
		query["ApplicationOwner"] = request.ApplicationOwner
	}

	if !dara.IsNil(request.ApplicationVisibility) {
		query["ApplicationVisibility"] = request.ApplicationVisibility
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CustomFields) {
		query["CustomFields"] = request.CustomFields
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogoUrl) {
		query["LogoUrl"] = request.LogoUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationInfo"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an application role.
//
// @param request - UpdateApplicationRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationRoleResponse
func (client *Client) UpdateApplicationRoleWithContext(ctx context.Context, request *UpdateApplicationRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
	}

	if !dara.IsNil(request.ApplicationRoleName) {
		query["ApplicationRoleName"] = request.ApplicationRoleName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationRole"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an application role.
//
// @param request - UpdateApplicationRoleDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationRoleDescriptionResponse
func (client *Client) UpdateApplicationRoleDescriptionWithContext(ctx context.Context, request *UpdateApplicationRoleDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationRoleDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationRoleId) {
		query["ApplicationRoleId"] = request.ApplicationRoleId
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
		Action:      dara.String("UpdateApplicationRoleDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationRoleDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the SSO parameters for an application template.
//
// Description:
//
// Updates the single sign-on (SSO) parameters for an application template from the marketplace.
//
// @param request - UpdateApplicationSsoFormParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationSsoFormParamsResponse
func (client *Client) UpdateApplicationSsoFormParamsWithContext(ctx context.Context, request *UpdateApplicationSsoFormParamsRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationSsoFormParamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationTemplateParams) {
		query["ApplicationTemplateParams"] = request.ApplicationTemplateParams
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationSsoFormParams"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationSsoFormParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the expiration time of an application token.
//
// @param request - UpdateApplicationTokenExpirationTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationTokenExpirationTimeResponse
func (client *Client) UpdateApplicationTokenExpirationTimeWithContext(ctx context.Context, request *UpdateApplicationTokenExpirationTimeRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationTokenExpirationTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ApplicationTokenId) {
		query["ApplicationTokenId"] = request.ApplicationTokenId
	}

	if !dara.IsNil(request.ExpirationTime) {
		query["ExpirationTime"] = request.ExpirationTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationTokenExpirationTime"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationTokenExpirationTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic properties of an authorization rule.
//
// @param request - UpdateAuthorizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationRuleResponse
func (client *Client) UpdateAuthorizationRuleWithContext(ctx context.Context, request *UpdateAuthorizationRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceScope) {
		query["AuthorizationResourceScope"] = request.AuthorizationResourceScope
	}

	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.AuthorizationRuleName) {
		query["AuthorizationRuleName"] = request.AuthorizationRuleName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthorizationRule"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthorizationRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the properties of the relationship between an authorization rule and an application.
//
// @param request - UpdateAuthorizationRuleApplicationAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationRuleApplicationAttachmentResponse
func (client *Client) UpdateAuthorizationRuleApplicationAttachmentWithContext(ctx context.Context, request *UpdateAuthorizationRuleApplicationAttachmentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthorizationRuleApplicationAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	if !dara.IsNil(request.ValidityType) {
		query["ValidityType"] = request.ValidityType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthorizationRuleApplicationAttachment"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthorizationRuleApplicationAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of an authorization rule.
//
// @param request - UpdateAuthorizationRuleDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationRuleDescriptionResponse
func (client *Client) UpdateAuthorizationRuleDescriptionWithContext(ctx context.Context, request *UpdateAuthorizationRuleDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthorizationRuleDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthorizationRuleDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthorizationRuleDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the relationship properties between an authorization rule and a group.
//
// @param request - UpdateAuthorizationRuleGroupAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationRuleGroupAttachmentResponse
func (client *Client) UpdateAuthorizationRuleGroupAttachmentWithContext(ctx context.Context, request *UpdateAuthorizationRuleGroupAttachmentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthorizationRuleGroupAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	if !dara.IsNil(request.ValidityType) {
		query["ValidityType"] = request.ValidityType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthorizationRuleGroupAttachment"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthorizationRuleGroupAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the relationship properties between an authorization rule and an account.
//
// @param request - UpdateAuthorizationRuleUserAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationRuleUserAttachmentResponse
func (client *Client) UpdateAuthorizationRuleUserAttachmentWithContext(ctx context.Context, request *UpdateAuthorizationRuleUserAttachmentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthorizationRuleUserAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRuleId) {
		query["AuthorizationRuleId"] = request.AuthorizationRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.ValidityPeriod) {
		query["ValidityPeriod"] = request.ValidityPeriod
	}

	if !dara.IsNil(request.ValidityType) {
		query["ValidityType"] = request.ValidityType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthorizationRuleUserAttachment"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthorizationRuleUserAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an authorization server.
//
// @param request - UpdateAuthorizationServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationServerResponse
func (client *Client) UpdateAuthorizationServerWithContext(ctx context.Context, request *UpdateAuthorizationServerRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthorizationServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationServerId) {
		query["AuthorizationServerId"] = request.AuthorizationServerId
	}

	if !dara.IsNil(request.AuthorizationServerName) {
		query["AuthorizationServerName"] = request.AuthorizationServerName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IssuerDomain) {
		query["IssuerDomain"] = request.IssuerDomain
	}

	if !dara.IsNil(request.IssuerMode) {
		query["IssuerMode"] = request.IssuerMode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthorizationServer"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthorizationServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of an authorization server.
//
// @param request - UpdateAuthorizationServerDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationServerDescriptionResponse
func (client *Client) UpdateAuthorizationServerDescriptionWithContext(ctx context.Context, request *UpdateAuthorizationServerDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthorizationServerDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationServerId) {
		query["AuthorizationServerId"] = request.AuthorizationServerId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthorizationServerDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthorizationServerDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a brand.
//
// @param request - UpdateBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBrandResponse
func (client *Client) UpdateBrandWithContext(ctx context.Context, request *UpdateBrandRequest, runtime *dara.RuntimeOptions) (_result *UpdateBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

	if !dara.IsNil(request.BrandName) {
		query["BrandName"] = request.BrandName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information of an Alibaba Cloud account.
//
// @param request - UpdateCloudAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudAccountResponse
func (client *Client) UpdateCloudAccountWithContext(ctx context.Context, request *UpdateCloudAccountRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.CloudAccountName) {
		query["CloudAccountName"] = request.CloudAccountName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudAccount"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of an Alibaba Cloud account.
//
// @param request - UpdateCloudAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudAccountDescriptionResponse
func (client *Client) UpdateCloudAccountDescriptionWithContext(ctx context.Context, request *UpdateCloudAccountDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudAccountDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
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
		Action:      dara.String("UpdateCloudAccountDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudAccountDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a cloud role.
//
// @param request - UpdateCloudAccountRoleDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudAccountRoleDescriptionResponse
func (client *Client) UpdateCloudAccountRoleDescriptionWithContext(ctx context.Context, request *UpdateCloudAccountRoleDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudAccountRoleDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloudAccountId) {
		query["CloudAccountId"] = request.CloudAccountId
	}

	if !dara.IsNil(request.CloudAccountRoleId) {
		query["CloudAccountRoleId"] = request.CloudAccountRoleId
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
		Action:      dara.String("UpdateCloudAccountRoleDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudAccountRoleDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a conditional access policy.
//
// Description:
//
// Updates a conditional access policy.
//
// @param request - UpdateConditionalAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConditionalAccessPolicyResponse
func (client *Client) UpdateConditionalAccessPolicyWithContext(ctx context.Context, request *UpdateConditionalAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateConditionalAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a conditional access policy.
//
// Description:
//
// Updates the description of a conditional access policy.
//
// @param request - UpdateConditionalAccessPolicyDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConditionalAccessPolicyDescriptionResponse
func (client *Client) UpdateConditionalAccessPolicyDescriptionWithContext(ctx context.Context, request *UpdateConditionalAccessPolicyDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateConditionalAccessPolicyDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update basic information for a credential.
//
// @param request - UpdateCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialResponse
func (client *Client) UpdateCredentialWithContext(ctx context.Context, request *UpdateCredentialRequest, runtime *dara.RuntimeOptions) (_result *UpdateCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredentialContent) {
		query["CredentialContent"] = request.CredentialContent
	}

	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
	}

	if !dara.IsNil(request.CredentialName) {
		query["CredentialName"] = request.CredentialName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCredential"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a credential.
//
// @param request - UpdateCredentialDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialDescriptionResponse
func (client *Client) UpdateCredentialDescriptionWithContext(ctx context.Context, request *UpdateCredentialDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateCredentialDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredentialId) {
		query["CredentialId"] = request.CredentialId
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
		Action:      dara.String("UpdateCredentialDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a credential provider.
//
// @param request - UpdateCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialProviderResponse
func (client *Client) UpdateCredentialProviderWithContext(ctx context.Context, request *UpdateCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CredentialProviderConfig) {
		query["CredentialProviderConfig"] = request.CredentialProviderConfig
	}

	if !dara.IsNil(request.CredentialProviderId) {
		query["CredentialProviderId"] = request.CredentialProviderId
	}

	if !dara.IsNil(request.CredentialProviderName) {
		query["CredentialProviderName"] = request.CredentialProviderName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the description of a credential provider.
//
// @param request - UpdateCredentialProviderDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialProviderDescriptionResponse
func (client *Client) UpdateCredentialProviderDescriptionWithContext(ctx context.Context, request *UpdateCredentialProviderDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateCredentialProviderDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialProviderId) {
		query["CredentialProviderId"] = request.CredentialProviderId
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
		Action:      dara.String("UpdateCredentialProviderDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialProviderDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom privacy policy.
//
// @param request - UpdateCustomPrivacyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomPrivacyPolicyResponse
func (client *Client) UpdateCustomPrivacyPolicyWithContext(ctx context.Context, request *UpdateCustomPrivacyPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomPrivacyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrivacyPolicyContents) {
		query["CustomPrivacyPolicyContents"] = request.CustomPrivacyPolicyContents
	}

	if !dara.IsNil(request.CustomPrivacyPolicyId) {
		query["CustomPrivacyPolicyId"] = request.CustomPrivacyPolicyId
	}

	if !dara.IsNil(request.CustomPrivacyPolicyName) {
		query["CustomPrivacyPolicyName"] = request.CustomPrivacyPolicyName
	}

	if !dara.IsNil(request.DefaultLanguageCode) {
		query["DefaultLanguageCode"] = request.DefaultLanguageCode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.UserConsentType) {
		query["UserConsentType"] = request.UserConsentType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomPrivacyPolicy"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomPrivacyPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the brand associated with a domain name.
//
// @param request - UpdateDomainBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainBrandResponse
func (client *Client) UpdateDomainBrandWithContext(ctx context.Context, request *UpdateDomainBrandRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandId) {
		query["BrandId"] = request.BrandId
	}

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
		Action:      dara.String("UpdateDomainBrand"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainBrandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the ICP filing number of a domain.
//
// @param request - UpdateDomainIcpNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainIcpNumberResponse
func (client *Client) UpdateDomainIcpNumberWithContext(ctx context.Context, request *UpdateDomainIcpNumberRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainIcpNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.IcpNumber) {
		query["IcpNumber"] = request.IcpNumber
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainIcpNumber"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainIcpNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a federated trust source.
//
// @param request - UpdateFederatedCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFederatedCredentialProviderResponse
func (client *Client) UpdateFederatedCredentialProviderWithContext(ctx context.Context, request *UpdateFederatedCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateFederatedCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederatedCredentialProviderId) {
		query["FederatedCredentialProviderId"] = request.FederatedCredentialProviderId
	}

	if !dara.IsNil(request.FederatedCredentialProviderName) {
		query["FederatedCredentialProviderName"] = request.FederatedCredentialProviderName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkAccessEndpointId) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !dara.IsNil(request.OidcProviderConfig) {
		query["OidcProviderConfig"] = request.OidcProviderConfig
	}

	if !dara.IsNil(request.Pkcs7ProviderConfig) {
		query["Pkcs7ProviderConfig"] = request.Pkcs7ProviderConfig
	}

	if !dara.IsNil(request.PrivateCaProviderConfig) {
		query["PrivateCaProviderConfig"] = request.PrivateCaProviderConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFederatedCredentialProvider"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFederatedCredentialProviderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a federated trust source.
//
// @param request - UpdateFederatedCredentialProviderDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFederatedCredentialProviderDescriptionResponse
func (client *Client) UpdateFederatedCredentialProviderDescriptionWithContext(ctx context.Context, request *UpdateFederatedCredentialProviderDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateFederatedCredentialProviderDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FederatedCredentialProviderId) {
		query["FederatedCredentialProviderId"] = request.FederatedCredentialProviderId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFederatedCredentialProviderDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFederatedCredentialProviderDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithContext(ctx context.Context, request *UpdateGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupDescriptionResponse
func (client *Client) UpdateGroupDescriptionWithContext(ctx context.Context, request *UpdateGroupDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateGroupDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic configuration of an identity provider.
//
// @param request - UpdateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProviderWithContext(ctx context.Context, request *UpdateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.SamlConfig) {
		query["SamlConfig"] = request.SamlConfig
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a specified EIAM instance.
//
// @param request - UpdateInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceDescriptionResponse
func (client *Client) UpdateInstanceDescriptionWithContext(ctx context.Context, request *UpdateInstanceDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a private network access endpoint.
//
// @param request - UpdateNetworkAccessEndpointNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkAccessEndpointNameResponse
func (client *Client) UpdateNetworkAccessEndpointNameWithContext(ctx context.Context, request *UpdateNetworkAccessEndpointNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateNetworkAccessEndpointNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a network zone object.
//
// @param request - UpdateNetworkZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkZoneResponse
func (client *Client) UpdateNetworkZoneWithContext(ctx context.Context, request *UpdateNetworkZoneRequest, runtime *dara.RuntimeOptions) (_result *UpdateNetworkZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ipv4Cidrs) {
		query["Ipv4Cidrs"] = request.Ipv4Cidrs
	}

	if !dara.IsNil(request.Ipv6Cidrs) {
		query["Ipv6Cidrs"] = request.Ipv6Cidrs
	}

	if !dara.IsNil(request.NetworkZoneId) {
		query["NetworkZoneId"] = request.NetworkZoneId
	}

	if !dara.IsNil(request.NetworkZoneName) {
		query["NetworkZoneName"] = request.NetworkZoneName
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNetworkZone"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNetworkZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a network zone.
//
// @param request - UpdateNetworkZoneDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkZoneDescriptionResponse
func (client *Client) UpdateNetworkZoneDescriptionWithContext(ctx context.Context, request *UpdateNetworkZoneDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateNetworkZoneDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

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
		Action:      dara.String("UpdateNetworkZoneDescription"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNetworkZoneDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitResponse
func (client *Client) UpdateOrganizationalUnitWithContext(ctx context.Context, request *UpdateOrganizationalUnitRequest, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitDescriptionResponse
func (client *Client) UpdateOrganizationalUnitDescriptionWithContext(ctx context.Context, request *UpdateOrganizationalUnitDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationalUnitDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitParentIdResponse
func (client *Client) UpdateOrganizationalUnitParentIdWithContext(ctx context.Context, request *UpdateOrganizationalUnitParentIdRequest, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationalUnitParentIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a scope permission for a specified resource server.
//
// @param request - UpdateResourceServerScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceServerScopeResponse
func (client *Client) UpdateResourceServerScopeWithContext(ctx context.Context, request *UpdateResourceServerScopeRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceServerScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceServerScopeId) {
		query["ResourceServerScopeId"] = request.ResourceServerScopeId
	}

	if !dara.IsNil(request.ResourceServerScopeName) {
		query["ResourceServerScopeName"] = request.ResourceServerScopeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceServerScope"),
		Version:     dara.String("2021-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceServerScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information of an EIAM account.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a user\\"s description.
//
// @param request - UpdateUserDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDescriptionResponse
func (client *Client) UpdateUserDescriptionWithContext(ctx context.Context, request *UpdateUserDescriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPasswordWithContext(ctx context.Context, request *UpdateUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
