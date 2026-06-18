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
	client.EndpointMap = map[string]*string{
		"cn-shanghai":    dara.String("eds-user.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1": dara.String("eds-user.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("eds-user"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
func (client *Client) BatchSetDesktopManagerWithOptions(request *BatchSetDesktopManagerRequest, runtime *dara.RuntimeOptions) (_result *BatchSetDesktopManagerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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
// @return BatchSetDesktopManagerResponse
func (client *Client) BatchSetDesktopManager(request *BatchSetDesktopManagerRequest) (_result *BatchSetDesktopManagerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSetDesktopManagerResponse{}
	_body, _err := client.BatchSetDesktopManagerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChangeUserPasswordWithOptions(request *ChangeUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ChangeUserPasswordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChangeUserPasswordResponse
func (client *Client) ChangeUserPassword(request *ChangeUserPasswordRequest) (_result *ChangeUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeUserPasswordResponse{}
	_body, _err := client.ChangeUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckUsedPropertyWithOptions(request *CheckUsedPropertyRequest, runtime *dara.RuntimeOptions) (_result *CheckUsedPropertyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckUsedPropertyResponse
func (client *Client) CheckUsedProperty(request *CheckUsedPropertyRequest) (_result *CheckUsedPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckUsedPropertyResponse{}
	_body, _err := client.CheckUsedPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckUsedPropertyValueWithOptions(request *CheckUsedPropertyValueRequest, runtime *dara.RuntimeOptions) (_result *CheckUsedPropertyValueResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckUsedPropertyValueResponse
func (client *Client) CheckUsedPropertyValue(request *CheckUsedPropertyValueRequest) (_result *CheckUsedPropertyValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckUsedPropertyValueResponse{}
	_body, _err := client.CheckUsedPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Create an organization.
//
// @param request - CreateOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrgResponse
func (client *Client) CreateOrgWithOptions(request *CreateOrgRequest, runtime *dara.RuntimeOptions) (_result *CreateOrgResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateOrgResponse
func (client *Client) CreateOrg(request *CreateOrgRequest) (_result *CreateOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOrgResponse{}
	_body, _err := client.CreateOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreatePropertyWithOptions(request *CreatePropertyRequest, runtime *dara.RuntimeOptions) (_result *CreatePropertyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreatePropertyResponse
func (client *Client) CreateProperty(request *CreatePropertyRequest) (_result *CreatePropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePropertyResponse{}
	_body, _err := client.CreatePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateUsersWithOptions(request *CreateUsersRequest, runtime *dara.RuntimeOptions) (_result *CreateUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateUsersResponse
func (client *Client) CreateUsers(request *CreateUsersRequest) (_result *CreateUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUsersResponse{}
	_body, _err := client.CreateUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteResourceGroupWithOptions(request *DeleteResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteUserPropertyValueWithOptions(request *DeleteUserPropertyValueRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserPropertyValueResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteUserPropertyValueResponse
func (client *Client) DeleteUserPropertyValue(request *DeleteUserPropertyValueRequest) (_result *DeleteUserPropertyValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserPropertyValueResponse{}
	_body, _err := client.DeleteUserPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeGroupUserWithOptions(request *DescribeGroupUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeGroupUserResponse
func (client *Client) DescribeGroupUser(request *DescribeGroupUserRequest) (_result *DescribeGroupUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupUserResponse{}
	_body, _err := client.DescribeGroupUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeGroupsWithOptions(request *DescribeGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeGroupsResponse
func (client *Client) DescribeGroups(request *DescribeGroupsRequest) (_result *DescribeGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.DescribeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeMfaDevicesWithOptions(request *DescribeMfaDevicesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMfaDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeMfaDevicesResponse
func (client *Client) DescribeMfaDevices(request *DescribeMfaDevicesRequest) (_result *DescribeMfaDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMfaDevicesResponse{}
	_body, _err := client.DescribeMfaDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeOrgByLayerWithOptions(request *DescribeOrgByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeOrgByLayerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeOrgByLayerResponse
func (client *Client) DescribeOrgByLayer(request *DescribeOrgByLayerRequest) (_result *DescribeOrgByLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOrgByLayerResponse{}
	_body, _err := client.DescribeOrgByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeOrgsWithOptions(tmpReq *DescribeOrgsRequest, runtime *dara.RuntimeOptions) (_result *DescribeOrgsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DescribeOrgsRequest
//
// @return DescribeOrgsResponse
func (client *Client) DescribeOrgs(request *DescribeOrgsRequest) (_result *DescribeOrgsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOrgsResponse{}
	_body, _err := client.DescribeOrgsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View resource groups.
//
// @param request - DescribeResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceGroupsResponse
func (client *Client) DescribeResourceGroupsWithOptions(request *DescribeResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View resource groups.
//
// @param request - DescribeResourceGroupsRequest
//
// @return DescribeResourceGroupsResponse
func (client *Client) DescribeResourceGroups(request *DescribeResourceGroupsRequest) (_result *DescribeResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceGroupsResponse{}
	_body, _err := client.DescribeResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeUserWithOptions(request *DescribeUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeUserResponse
func (client *Client) DescribeUser(request *DescribeUserRequest) (_result *DescribeUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserResponse{}
	_body, _err := client.DescribeUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeUsersWithOptions(tmpReq *DescribeUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DescribeUsersRequest
//
// @return DescribeUsersResponse
func (client *Client) DescribeUsers(request *DescribeUsersRequest) (_result *DescribeUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUsersResponse{}
	_body, _err := client.DescribeUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FilterUsersWithOptions(tmpReq *FilterUsersRequest, runtime *dara.RuntimeOptions) (_result *FilterUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - FilterUsersRequest
//
// @return FilterUsersResponse
func (client *Client) FilterUsers(request *FilterUsersRequest) (_result *FilterUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FilterUsersResponse{}
	_body, _err := client.FilterUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetManagerInfoByAuthCodeWithOptions(request *GetManagerInfoByAuthCodeRequest, runtime *dara.RuntimeOptions) (_result *GetManagerInfoByAuthCodeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetManagerInfoByAuthCodeResponse
func (client *Client) GetManagerInfoByAuthCode(request *GetManagerInfoByAuthCodeRequest) (_result *GetManagerInfoByAuthCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetManagerInfoByAuthCodeResponse{}
	_body, _err := client.GetManagerInfoByAuthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InitTenantAliasWithOptions(request *InitTenantAliasRequest, runtime *dara.RuntimeOptions) (_result *InitTenantAliasResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return InitTenantAliasResponse
func (client *Client) InitTenantAlias(request *InitTenantAliasRequest) (_result *InitTenantAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitTenantAliasResponse{}
	_body, _err := client.InitTenantAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPropertyWithOptions(request *ListPropertyRequest, runtime *dara.RuntimeOptions) (_result *ListPropertyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPropertyResponse
func (client *Client) ListProperty(request *ListPropertyRequest) (_result *ListPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPropertyResponse{}
	_body, _err := client.ListPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPropertyValueWithOptions(request *ListPropertyValueRequest, runtime *dara.RuntimeOptions) (_result *ListPropertyValueResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPropertyValueResponse
func (client *Client) ListPropertyValue(request *ListPropertyValueRequest) (_result *ListPropertyValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPropertyValueResponse{}
	_body, _err := client.ListPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) LockMfaDeviceWithOptions(request *LockMfaDeviceRequest, runtime *dara.RuntimeOptions) (_result *LockMfaDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return LockMfaDeviceResponse
func (client *Client) LockMfaDevice(request *LockMfaDeviceRequest) (_result *LockMfaDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LockMfaDeviceResponse{}
	_body, _err := client.LockMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) LockUsersWithOptions(request *LockUsersRequest, runtime *dara.RuntimeOptions) (_result *LockUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return LockUsersResponse
func (client *Client) LockUsers(request *LockUsersRequest) (_result *LockUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LockUsersResponse{}
	_body, _err := client.LockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyGroupWithOptions(request *ModifyGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyGroupResponse
func (client *Client) ModifyGroup(request *ModifyGroupRequest) (_result *ModifyGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGroupResponse{}
	_body, _err := client.ModifyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyOrgWithOptions(request *ModifyOrgRequest, runtime *dara.RuntimeOptions) (_result *ModifyOrgResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyOrgResponse
func (client *Client) ModifyOrg(request *ModifyOrgRequest) (_result *ModifyOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyOrgResponse{}
	_body, _err := client.ModifyOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyUserWithOptions(request *ModifyUserRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyUserResponse
func (client *Client) ModifyUser(request *ModifyUserRequest) (_result *ModifyUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserResponse{}
	_body, _err := client.ModifyUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MoveOrgWithOptions(request *MoveOrgRequest, runtime *dara.RuntimeOptions) (_result *MoveOrgResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MoveOrgResponse
func (client *Client) MoveOrg(request *MoveOrgRequest) (_result *MoveOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveOrgResponse{}
	_body, _err := client.MoveOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MoveUserOrgWithOptions(request *MoveUserOrgRequest, runtime *dara.RuntimeOptions) (_result *MoveUserOrgResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MoveUserOrgResponse
func (client *Client) MoveUserOrg(request *MoveUserOrgRequest) (_result *MoveUserOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveUserOrgResponse{}
	_body, _err := client.MoveUserOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QuerySyncStatusByAliUidWithOptions(request *QuerySyncStatusByAliUidRequest, runtime *dara.RuntimeOptions) (_result *QuerySyncStatusByAliUidResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QuerySyncStatusByAliUidResponse
func (client *Client) QuerySyncStatusByAliUid(request *QuerySyncStatusByAliUidRequest) (_result *QuerySyncStatusByAliUidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySyncStatusByAliUidResponse{}
	_body, _err := client.QuerySyncStatusByAliUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveGroupWithOptions(request *RemoveGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveGroupResponse
func (client *Client) RemoveGroup(request *RemoveGroupRequest) (_result *RemoveGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveGroupResponse{}
	_body, _err := client.RemoveGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveMfaDeviceWithOptions(request *RemoveMfaDeviceRequest, runtime *dara.RuntimeOptions) (_result *RemoveMfaDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveMfaDeviceResponse
func (client *Client) RemoveMfaDevice(request *RemoveMfaDeviceRequest) (_result *RemoveMfaDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveMfaDeviceResponse{}
	_body, _err := client.RemoveMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveOrgWithOptions(request *RemoveOrgRequest, runtime *dara.RuntimeOptions) (_result *RemoveOrgResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveOrgResponse
func (client *Client) RemoveOrg(request *RemoveOrgRequest) (_result *RemoveOrgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveOrgResponse{}
	_body, _err := client.RemoveOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemovePropertyWithOptions(request *RemovePropertyRequest, runtime *dara.RuntimeOptions) (_result *RemovePropertyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemovePropertyResponse
func (client *Client) RemoveProperty(request *RemovePropertyRequest) (_result *RemovePropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePropertyResponse{}
	_body, _err := client.RemovePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveUsersWithOptions(request *RemoveUsersRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveUsersResponse
func (client *Client) RemoveUsers(request *RemoveUsersRequest) (_result *RemoveUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUsersResponse{}
	_body, _err := client.RemoveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ResetUserPasswordWithOptions(request *ResetUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPassword(request *ResetUserPasswordRequest) (_result *ResetUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.ResetUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetUserPropertyValueWithOptions(request *SetUserPropertyValueRequest, runtime *dara.RuntimeOptions) (_result *SetUserPropertyValueResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetUserPropertyValueResponse
func (client *Client) SetUserPropertyValue(request *SetUserPropertyValueRequest) (_result *SetUserPropertyValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetUserPropertyValueResponse{}
	_body, _err := client.SetUserPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronize all education information.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncAllEduInfoResponse
func (client *Client) SyncAllEduInfoWithOptions(runtime *dara.RuntimeOptions) (_result *SyncAllEduInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("SyncAllEduInfo"),
		Version:     dara.String("2021-03-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncAllEduInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronize all education information.
//
// @return SyncAllEduInfoResponse
func (client *Client) SyncAllEduInfo() (_result *SyncAllEduInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncAllEduInfoResponse{}
	_body, _err := client.SyncAllEduInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TransferResourcesIntoGroupWithOptions(request *TransferResourcesIntoGroupRequest, runtime *dara.RuntimeOptions) (_result *TransferResourcesIntoGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TransferResourcesIntoGroupResponse
func (client *Client) TransferResourcesIntoGroup(request *TransferResourcesIntoGroupRequest) (_result *TransferResourcesIntoGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferResourcesIntoGroupResponse{}
	_body, _err := client.TransferResourcesIntoGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TransferResourcesOutofGroupWithOptions(request *TransferResourcesOutofGroupRequest, runtime *dara.RuntimeOptions) (_result *TransferResourcesOutofGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TransferResourcesOutofGroupResponse
func (client *Client) TransferResourcesOutofGroup(request *TransferResourcesOutofGroupRequest) (_result *TransferResourcesOutofGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferResourcesOutofGroupResponse{}
	_body, _err := client.TransferResourcesOutofGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnlockMfaDeviceWithOptions(request *UnlockMfaDeviceRequest, runtime *dara.RuntimeOptions) (_result *UnlockMfaDeviceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnlockMfaDeviceResponse
func (client *Client) UnlockMfaDevice(request *UnlockMfaDeviceRequest) (_result *UnlockMfaDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockMfaDeviceResponse{}
	_body, _err := client.UnlockMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnlockUsersWithOptions(request *UnlockUsersRequest, runtime *dara.RuntimeOptions) (_result *UnlockUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnlockUsersResponse
func (client *Client) UnlockUsers(request *UnlockUsersRequest) (_result *UnlockUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnlockUsersResponse{}
	_body, _err := client.UnlockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdatePropertyWithOptions(request *UpdatePropertyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePropertyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdatePropertyResponse
func (client *Client) UpdateProperty(request *UpdatePropertyRequest) (_result *UpdatePropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePropertyResponse{}
	_body, _err := client.UpdatePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UserBatchJoinGroupWithOptions(request *UserBatchJoinGroupRequest, runtime *dara.RuntimeOptions) (_result *UserBatchJoinGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UserBatchJoinGroupResponse
func (client *Client) UserBatchJoinGroup(request *UserBatchJoinGroupRequest) (_result *UserBatchJoinGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UserBatchJoinGroupResponse{}
	_body, _err := client.UserBatchJoinGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UserBatchQuitGroupWithOptions(request *UserBatchQuitGroupRequest, runtime *dara.RuntimeOptions) (_result *UserBatchQuitGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UserBatchQuitGroupResponse
func (client *Client) UserBatchQuitGroup(request *UserBatchQuitGroupRequest) (_result *UserBatchQuitGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UserBatchQuitGroupResponse{}
	_body, _err := client.UserBatchQuitGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
