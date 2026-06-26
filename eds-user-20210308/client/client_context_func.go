// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Add or remove local administrator permissions on a cloud computer for a convenience account.
//
// Description:
//
// A convenience account with local administrator permissions can install software or modify certain system settings on the cloud computer.
//
// @param request - BatchSetDesktopManagerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetDesktopManagerResponse
func (client *Client) BatchSetDesktopManagerWithContext(ctx context.Context, request *BatchSetDesktopManagerRequest, runtime *dara.RuntimeOptions) (_result *BatchSetDesktopManagerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.IsDesktopManager) {
		body["IsDesktopManager"] = request.IsDesktopManager
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetDesktopManager"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetDesktopManagerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the administrator activates the mode, you can use this API to modify the user logon password.
//
// @param request - ChangeUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeUserPasswordResponse
func (client *Client) ChangeUserPasswordWithContext(ctx context.Context, request *ChangeUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ChangeUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.NewPassword) {
		body["NewPassword"] = request.NewPassword
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeUserPassword"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeUserPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the number of convenience accounts associated with a specified custom property.
//
// @param request - CheckUsedPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUsedPropertyResponse
func (client *Client) CheckUsedPropertyWithContext(ctx context.Context, request *CheckUsedPropertyRequest, runtime *dara.RuntimeOptions) (_result *CheckUsedPropertyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.PropertyId) {
		query["PropertyId"] = request.PropertyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckUsedProperty"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckUsedPropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the number of convenience accounts associated with a specified custom attribute value.
//
// @param request - CheckUsedPropertyValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUsedPropertyValueResponse
func (client *Client) CheckUsedPropertyValueWithContext(ctx context.Context, request *CheckUsedPropertyValueRequest, runtime *dara.RuntimeOptions) (_result *CheckUsedPropertyValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.PropertyId) {
		query["PropertyId"] = request.PropertyId
	}

	if !dara.IsNil(request.PropertyValueId) {
		query["PropertyValueId"] = request.PropertyValueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckUsedPropertyValue"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckUsedPropertyValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a group.
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
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ParentGroupId) {
		query["ParentGroupId"] = request.ParentGroupId
	}

	if !dara.IsNil(request.SolutionId) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroup"),
		Version:     dara.String("2021-03-08"),
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
// Create an organization.
//
// @param request - CreateOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrgResponse
func (client *Client) CreateOrgWithContext(ctx context.Context, request *CreateOrgRequest, runtime *dara.RuntimeOptions) (_result *CreateOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.OrgName) {
		query["OrgName"] = request.OrgName
	}

	if !dara.IsNil(request.ParentOrgId) {
		query["ParentOrgId"] = request.ParentOrgId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrg"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a user attribute.
//
// Description:
//
// - You can create up to 10 different properties under one Alibaba Cloud account. Each property includes a property name (PropertyKey) and multiple attribute values (PropertyValue).
//
// - You can add up to 50 different attribute values to a single property.
//
// @param request - CreatePropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePropertyResponse
func (client *Client) CreatePropertyWithContext(ctx context.Context, request *CreatePropertyRequest, runtime *dara.RuntimeOptions) (_result *CreatePropertyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.PropertyKey) {
		body["PropertyKey"] = request.PropertyKey
	}

	if !dara.IsNil(request.PropertyValues) {
		body["PropertyValues"] = request.PropertyValues
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProperty"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource group.
//
// @param request - CreateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithContext(ctx context.Context, request *CreateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.EnableAliyunResourceGroup) {
		query["EnableAliyunResourceGroup"] = request.EnableAliyunResourceGroup
	}

	if !dara.IsNil(request.IsResourceGroupWithOfficeSite) {
		query["IsResourceGroupWithOfficeSite"] = request.IsResourceGroupWithOfficeSite
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceClassification) {
		query["ResourceClassification"] = request.ResourceClassification
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceGroup"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Easily create accounts for your end users.
//
// Description:
//
// <props="china">
//
// A convenience account is a dedicated account system in Wuying Workspace for simple use cases that do not require enterprise AD integration. Accounts require a username, and either an email or a phone number.
//
// <props="intl">
//
// A convenience account is a dedicated account system in Wuying Workspace for simple use cases that do not require enterprise AD integration. Accounts require both a username and an email.
//
// @param request - CreateUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUsersResponse
func (client *Client) CreateUsersWithContext(ctx context.Context, request *CreateUsersRequest, runtime *dara.RuntimeOptions) (_result *CreateUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoLockTime) {
		query["AutoLockTime"] = request.AutoLockTime
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.IsLocalAdmin) {
		query["IsLocalAdmin"] = request.IsLocalAdmin
	}

	if !dara.IsNil(request.PasswordExpireDays) {
		query["PasswordExpireDays"] = request.PasswordExpireDays
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUsers"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a resource group.
//
// @param request - DeleteResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithContext(ctx context.Context, request *DeleteResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceGroup"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach a user from a user attribute.
//
// @param request - DeleteUserPropertyValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPropertyValueResponse
func (client *Client) DeleteUserPropertyValueWithContext(ctx context.Context, request *DeleteUserPropertyValueRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserPropertyValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.PropertyId) {
		body["PropertyId"] = request.PropertyId
	}

	if !dara.IsNil(request.PropertyValueId) {
		body["PropertyValueId"] = request.PropertyValueId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserPropertyValue"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserPropertyValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query group members.
//
// @param request - DescribeGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupUserResponse
func (client *Client) DescribeGroupUserWithContext(ctx context.Context, request *DescribeGroupUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SolutionId) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupUser"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query user groups.
//
// @param request - DescribeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupsResponse
func (client *Client) DescribeGroupsWithContext(ctx context.Context, request *DescribeGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.ExcludeAttachedLoginPolicyGroups) {
		query["ExcludeAttachedLoginPolicyGroups"] = request.ExcludeAttachedLoginPolicyGroups
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IdpId) {
		query["IdpId"] = request.IdpId
	}

	if !dara.IsNil(request.LoginPolicyId) {
		query["LoginPolicyId"] = request.LoginPolicyId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SolutionId) {
		query["SolutionId"] = request.SolutionId
	}

	if !dara.IsNil(request.TransferFileNeedApproval) {
		query["TransferFileNeedApproval"] = request.TransferFileNeedApproval
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroups"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists virtual MFA devices bound to directory accounts.
//
// @param request - DescribeMfaDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMfaDevicesResponse
func (client *Client) DescribeMfaDevicesWithContext(ctx context.Context, request *DescribeMfaDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMfaDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		query["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.EndUserIds) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SerialNumbers) {
		query["SerialNumbers"] = request.SerialNumbers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMfaDevices"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMfaDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Find subordinate organizations.
//
// @param request - DescribeOrgByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOrgByLayerResponse
func (client *Client) DescribeOrgByLayerWithContext(ctx context.Context, request *DescribeOrgByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeOrgByLayerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.OrgName) {
		body["OrgName"] = request.OrgName
	}

	if !dara.IsNil(request.ParentOrgId) {
		body["ParentOrgId"] = request.ParentOrgId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOrgByLayer"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOrgByLayerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of organizations.
//
// Description:
//
// Organizations are arranged in a tree-like structure. The root organization ID is org-aliyun-wy-org-id.
//
// @param tmpReq - DescribeOrgsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOrgsResponse
func (client *Client) DescribeOrgsWithContext(ctx context.Context, tmpReq *DescribeOrgsRequest, runtime *dara.RuntimeOptions) (_result *DescribeOrgsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeOrgsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ShowExtras) {
		request.ShowExtrasShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ShowExtras, dara.String("ShowExtras"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.IncludeOrgIds) {
		query["IncludeOrgIds"] = request.IncludeOrgIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrgName) {
		query["OrgName"] = request.OrgName
	}

	if !dara.IsNil(request.ParentOrgId) {
		query["ParentOrgId"] = request.ParentOrgId
	}

	if !dara.IsNil(request.ShowExtrasShrink) {
		query["ShowExtras"] = request.ShowExtrasShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOrgs"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOrgsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource groups.
//
// @param request - DescribeResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceGroupsResponse
func (client *Client) DescribeResourceGroupsWithContext(ctx context.Context, request *DescribeResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.AliyunResourceGroupIds) {
		query["AliyunResourceGroupIds"] = request.AliyunResourceGroupIds
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.NeedContainResourceGroupWithOfficeSite) {
		query["NeedContainResourceGroupWithOfficeSite"] = request.NeedContainResourceGroupWithOfficeSite
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceClassification) {
		query["ResourceClassification"] = request.ResourceClassification
	}

	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceGroups"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query user basic information
//
// @param request - DescribeUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserResponse
func (client *Client) DescribeUserWithContext(ctx context.Context, request *DescribeUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.RequireExtraAttributes) {
		query["RequireExtraAttributes"] = request.RequireExtraAttributes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUser"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves directory account information, including the username, email address, and display name.
//
// @param tmpReq - DescribeUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsersResponse
func (client *Client) DescribeUsersWithContext(ctx context.Context, tmpReq *DescribeUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterMap) {
		request.FilterMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterMap, dara.String("FilterMap"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FilterWithAssignedResource) {
		request.FilterWithAssignedResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterWithAssignedResource, dara.String("FilterWithAssignedResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FilterWithAssignedResources) {
		request.FilterWithAssignedResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterWithAssignedResources, dara.String("FilterWithAssignedResources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ShowExtras) {
		request.ShowExtrasShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ShowExtras, dara.String("ShowExtras"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.FilterMapShrink) {
		query["FilterMap"] = request.FilterMapShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.EndUserIds) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.ExcludeEndUserIds) {
		body["ExcludeEndUserIds"] = request.ExcludeEndUserIds
	}

	if !dara.IsNil(request.ExcludeGroupId) {
		body["ExcludeGroupId"] = request.ExcludeGroupId
	}

	if !dara.IsNil(request.FilterWithAssignedResourceShrink) {
		body["FilterWithAssignedResource"] = request.FilterWithAssignedResourceShrink
	}

	if !dara.IsNil(request.FilterWithAssignedResourcesShrink) {
		body["FilterWithAssignedResources"] = request.FilterWithAssignedResourcesShrink
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsQueryAllSubOrgs) {
		body["IsQueryAllSubOrgs"] = request.IsQueryAllSubOrgs
	}

	if !dara.IsNil(request.OrgId) {
		body["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.ShowExtrasShrink) {
		body["ShowExtras"] = request.ShowExtrasShrink
	}

	if !dara.IsNil(request.SolutionId) {
		body["SolutionId"] = request.SolutionId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsers"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Filter account information by user attribute.
//
// @param tmpReq - FilterUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FilterUsersResponse
func (client *Client) FilterUsersWithContext(ctx context.Context, tmpReq *FilterUsersRequest, runtime *dara.RuntimeOptions) (_result *FilterUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &FilterUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterMap) {
		request.FilterMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterMap, dara.String("FilterMap"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OrderParam) {
		request.OrderParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderParam, dara.String("OrderParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.ExcludeEndUserIds) {
		query["ExcludeEndUserIds"] = request.ExcludeEndUserIds
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.FilterMapShrink) {
		query["FilterMap"] = request.FilterMapShrink
	}

	if !dara.IsNil(request.IncludeDesktopCount) {
		query["IncludeDesktopCount"] = request.IncludeDesktopCount
	}

	if !dara.IsNil(request.IncludeDesktopGroupCount) {
		query["IncludeDesktopGroupCount"] = request.IncludeDesktopGroupCount
	}

	if !dara.IsNil(request.IncludeEndUserIds) {
		query["IncludeEndUserIds"] = request.IncludeEndUserIds
	}

	if !dara.IsNil(request.IncludeOrgInfo) {
		query["IncludeOrgInfo"] = request.IncludeOrgInfo
	}

	if !dara.IsNil(request.IncludeSupportIdps) {
		query["IncludeSupportIdps"] = request.IncludeSupportIdps
	}

	if !dara.IsNil(request.IsQueryAllSubOrgs) {
		query["IsQueryAllSubOrgs"] = request.IsQueryAllSubOrgs
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderParamShrink) {
		query["OrderParam"] = request.OrderParamShrink
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.OwnerType) {
		query["OwnerType"] = request.OwnerType
	}

	if !dara.IsNil(request.PropertyFilterParam) {
		query["PropertyFilterParam"] = request.PropertyFilterParam
	}

	if !dara.IsNil(request.PropertyKeyValueFilterParam) {
		query["PropertyKeyValueFilterParam"] = request.PropertyKeyValueFilterParam
	}

	if !dara.IsNil(request.ShowExtras) {
		query["ShowExtras"] = request.ShowExtras
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FilterUsers"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FilterUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about the current logon administrator based on the authorization code.
//
// @param request - GetManagerInfoByAuthCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagerInfoByAuthCodeResponse
func (client *Client) GetManagerInfoByAuthCodeWithContext(ctx context.Context, request *GetManagerInfoByAuthCodeRequest, runtime *dara.RuntimeOptions) (_result *GetManagerInfoByAuthCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetManagerInfoByAuthCode"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetManagerInfoByAuthCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initialize the organization ID.
//
// @param request - InitTenantAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitTenantAliasResponse
func (client *Client) InitTenantAliasWithContext(ctx context.Context, request *InitTenantAliasRequest, runtime *dara.RuntimeOptions) (_result *InitTenantAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitTenantAlias"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitTenantAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of existing user attributes under the current account.
//
// @param request - ListPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPropertyResponse
func (client *Client) ListPropertyWithContext(ctx context.Context, request *ListPropertyRequest, runtime *dara.RuntimeOptions) (_result *ListPropertyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProperty"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of attribute values for a specific user attribute.
//
// @param request - ListPropertyValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPropertyValueResponse
func (client *Client) ListPropertyValueWithContext(ctx context.Context, request *ListPropertyValueRequest, runtime *dara.RuntimeOptions) (_result *ListPropertyValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.PropertyId) {
		query["PropertyId"] = request.PropertyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPropertyValue"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPropertyValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lock the virtual MFA device attached to a convenience account.
//
// Description:
//
// After locking, the status of the virtual MFA device changes to LOCKED. When the associated convenience account attempts to log on to a WUYING Terminal through an office network with MFA enabled, authentication will fail due to the locked MFA device, preventing successful logon. You can invoke [UnlockMfaDevice](~~UnlockMfaDevice~~) to unlock it.
//
// @param request - LockMfaDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockMfaDeviceResponse
func (client *Client) LockMfaDeviceWithContext(ctx context.Context, request *LockMfaDeviceRequest, runtime *dara.RuntimeOptions) (_result *LockMfaDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		query["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LockMfaDevice"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LockMfaDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Locks one or more convenience accounts. Locked convenience accounts cannot be used to sign in to Wuying clients.
//
// Description:
//
// For security purposes, you can lock convenience accounts. Locked convenience users cannot sign in to Wuying clients, and therefore cannot access any Wuying cloud resources.
//
// > Call the [DescribeUsers](https://help.aliyun.com/document_detail/283609.html) operation to check the lock status of convenience accounts. The `Status` value in the response is 0 for unlocked accounts and 9 for locked accounts.
//
// @param request - LockUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockUsersResponse
func (client *Client) LockUsersWithContext(ctx context.Context, request *LockUsersRequest, runtime *dara.RuntimeOptions) (_result *LockUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.LogoutSession) {
		query["LogoutSession"] = request.LogoutSession
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LockUsers"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LockUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the name and description of a group.
//
// @param request - ModifyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGroupResponse
func (client *Client) ModifyGroupWithContext(ctx context.Context, request *ModifyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NewGroupName) {
		query["NewGroupName"] = request.NewGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGroup"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify an organization.
//
// @param request - ModifyOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOrgResponse
func (client *Client) ModifyOrgWithContext(ctx context.Context, request *ModifyOrgRequest, runtime *dara.RuntimeOptions) (_result *ModifyOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.OrgName) {
		query["OrgName"] = request.OrgName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOrg"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOrgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the contact information of a convenience account.
//
// @param request - ModifyUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserResponse
func (client *Client) ModifyUserWithContext(ctx context.Context, request *ModifyUserRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUser"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shift organization.
//
// @param request - MoveOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveOrgResponse
func (client *Client) MoveOrgWithContext(ctx context.Context, request *MoveOrgRequest, runtime *dara.RuntimeOptions) (_result *MoveOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.NewParentOrgId) {
		body["NewParentOrgId"] = request.NewParentOrgId
	}

	if !dara.IsNil(request.OrgId) {
		body["OrgId"] = request.OrgId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveOrg"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveOrgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shift users to the target organization architecture.
//
// @param request - MoveUserOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveUserOrgResponse
func (client *Client) MoveUserOrgWithContext(ctx context.Context, request *MoveUserOrgRequest, runtime *dara.RuntimeOptions) (_result *MoveUserOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.EndUserIds) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.OrgId) {
		body["OrgId"] = request.OrgId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveUserOrg"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveUserOrgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the user synchronization status.
//
// @param request - QuerySyncStatusByAliUidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySyncStatusByAliUidResponse
func (client *Client) QuerySyncStatusByAliUidWithContext(ctx context.Context, request *QuerySyncStatusByAliUidRequest, runtime *dara.RuntimeOptions) (_result *QuerySyncStatusByAliUidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySyncStatusByAliUid"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySyncStatusByAliUidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a group. Supports batch operations.
//
// @param request - RemoveGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveGroupResponse
func (client *Client) RemoveGroupWithContext(ctx context.Context, request *RemoveGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveGroup"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the virtual MFA device attached to a convenience account.
//
// Description:
//
// Deleting the virtual MFA device attached to a convenience account unbinds the MFA device, which is equivalent to resetting or disabling it. The corresponding convenience account must reattach a new virtual MFA device when logging on to a WUYING Terminal.
//
// @param request - RemoveMfaDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveMfaDeviceResponse
func (client *Client) RemoveMfaDeviceWithContext(ctx context.Context, request *RemoveMfaDeviceRequest, runtime *dara.RuntimeOptions) (_result *RemoveMfaDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		query["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveMfaDevice"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveMfaDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove an organization.
//
// @param request - RemoveOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveOrgResponse
func (client *Client) RemoveOrgWithContext(ctx context.Context, request *RemoveOrgRequest, runtime *dara.RuntimeOptions) (_result *RemoveOrgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.OrgId) {
		body["OrgId"] = request.OrgId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveOrg"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveOrgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a set of user attributes.
//
// @param request - RemovePropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePropertyResponse
func (client *Client) RemovePropertyWithContext(ctx context.Context, request *RemovePropertyRequest, runtime *dara.RuntimeOptions) (_result *RemovePropertyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PropertyId) {
		body["PropertyId"] = request.PropertyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveProperty"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete one or more convenience accounts.
//
// @param request - RemoveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsersWithContext(ctx context.Context, request *RemoveUsersRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsers"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reset the password of a convenience account, including generating a password reset token and sending a password reset email to the mailbox of the convenience account.
//
// @param request - ResetUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPasswordWithContext(ctx context.Context, request *ResetUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.NotifyType) {
		body["NotifyType"] = request.NotifyType
	}

	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetUserPassword"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associate a user attribute with a specific user.
//
// @param request - SetUserPropertyValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserPropertyValueResponse
func (client *Client) SetUserPropertyValueWithContext(ctx context.Context, request *SetUserPropertyValueRequest, runtime *dara.RuntimeOptions) (_result *SetUserPropertyValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.PropertyId) {
		body["PropertyId"] = request.PropertyId
	}

	if !dara.IsNil(request.PropertyValueId) {
		body["PropertyValueId"] = request.PropertyValueId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetUserPropertyValue"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetUserPropertyValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将资源转移到资源组中
//
// @param request - TransferResourcesIntoGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferResourcesIntoGroupResponse
func (client *Client) TransferResourcesIntoGroupWithContext(ctx context.Context, request *TransferResourcesIntoGroupRequest, runtime *dara.RuntimeOptions) (_result *TransferResourcesIntoGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferResourcesIntoGroup"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferResourcesIntoGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将资源从资源组中转出
//
// @param request - TransferResourcesOutofGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferResourcesOutofGroupResponse
func (client *Client) TransferResourcesOutofGroupWithContext(ctx context.Context, request *TransferResourcesOutofGroupRequest, runtime *dara.RuntimeOptions) (_result *TransferResourcesOutofGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.TargetResourceGroupId) {
		query["TargetResourceGroupId"] = request.TargetResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferResourcesOutofGroup"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferResourcesOutofGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlock the virtual MFA device attached to a convenience account.
//
// @param request - UnlockMfaDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockMfaDeviceResponse
func (client *Client) UnlockMfaDeviceWithContext(ctx context.Context, request *UnlockMfaDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnlockMfaDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		query["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockMfaDevice"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockMfaDeviceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlock one or more convenience accounts. After being unlocked, the convenience accounts can log on to WUYING Terminal.
//
// Description:
//
// Locked convenience accounts cannot log on to WUYING Terminal and therefore cannot access any WUYING cloud resources. To allow a convenience account to log on to WUYING Terminal, you must first unlock it.
//
// > You can invoke [DescribeUsers](https://help.aliyun.com/document_detail/283609.html) to query convenience account information. If the value of `Status` in the returned data is 0, the convenience account is not locked. If the value of `Status` is 9, the convenience account is locked.
//
// @param request - UnlockUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockUsersResponse
func (client *Client) UnlockUsersWithContext(ctx context.Context, request *UnlockUsersRequest, runtime *dara.RuntimeOptions) (_result *UnlockUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoLockTime) {
		query["AutoLockTime"] = request.AutoLockTime
	}

	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Users) {
		body["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockUsers"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify User Attributes.
//
// @param request - UpdatePropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePropertyResponse
func (client *Client) UpdatePropertyWithContext(ctx context.Context, request *UpdatePropertyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePropertyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		query["BusinessChannel"] = request.BusinessChannel
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PropertyId) {
		body["PropertyId"] = request.PropertyId
	}

	if !dara.IsNil(request.PropertyKey) {
		body["PropertyKey"] = request.PropertyKey
	}

	if !dara.IsNil(request.PropertyValues) {
		body["PropertyValues"] = request.PropertyValues
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProperty"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePropertyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add users to a group in batch.
//
// @param request - UserBatchJoinGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UserBatchJoinGroupResponse
func (client *Client) UserBatchJoinGroupWithContext(ctx context.Context, request *UserBatchJoinGroupRequest, runtime *dara.RuntimeOptions) (_result *UserBatchJoinGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.EndUserIds) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UserBatchJoinGroup"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UserBatchJoinGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove users from a group in batch.
//
// @param request - UserBatchQuitGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UserBatchQuitGroupResponse
func (client *Client) UserBatchQuitGroupWithContext(ctx context.Context, request *UserBatchQuitGroupRequest, runtime *dara.RuntimeOptions) (_result *UserBatchQuitGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessChannel) {
		body["BusinessChannel"] = request.BusinessChannel
	}

	if !dara.IsNil(request.EndUserIds) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UserBatchQuitGroup"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UserBatchQuitGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
