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
	client.Endpoint, _err = client.GetEndpoint(dara.String("eiam-developerapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 将账户加入多个组织
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @param headers - AddUserToOrganizationalUnitsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnitsWithOptions(instanceId *string, applicationId *string, userId *string, request *AddUserToOrganizationalUnitsRequest, headers *AddUserToOrganizationalUnitsHeaders, runtime *dara.RuntimeOptions) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationalUnitIds) {
		body["organizationalUnitIds"] = request.OrganizationalUnitIds
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToOrganizationalUnits"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId)) + "/actions/addUserToOrganizationalUnits"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &AddUserToOrganizationalUnitsResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将账户加入多个组织
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnits(instanceId *string, applicationId *string, userId *string, request *AddUserToOrganizationalUnitsRequest) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddUserToOrganizationalUnitsHeaders{}
	_result = &AddUserToOrganizationalUnitsResponse{}
	_body, _err := client.AddUserToOrganizationalUnitsWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds multiple Employee Identity and Access Management (EIAM) accounts to an EIAM group. If the accounts are already added to the specified group, no update is performed.
//
// @param request - AddUsersToGroupRequest
//
// @param headers - AddUsersToGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroupWithOptions(instanceId *string, applicationId *string, groupId *string, request *AddUsersToGroupRequest, headers *AddUsersToGroupHeaders, runtime *dara.RuntimeOptions) (_result *AddUsersToGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserIds) {
		body["userIds"] = request.UserIds
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUsersToGroup"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/actions/addUsersToGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds multiple Employee Identity and Access Management (EIAM) accounts to an EIAM group. If the accounts are already added to the specified group, no update is performed.
//
// @param request - AddUsersToGroupRequest
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroup(instanceId *string, applicationId *string, groupId *string, request *AddUsersToGroupRequest) (_result *AddUsersToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddUsersToGroupHeaders{}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.AddUsersToGroupWithOptions(instanceId, applicationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a group.
//
// @param request - CreateGroupRequest
//
// @param headers - CreateGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(instanceId *string, applicationId *string, request *CreateGroupRequest, headers *CreateGroupHeaders, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupExternalId) {
		body["groupExternalId"] = request.GroupExternalId
	}

	if !dara.IsNil(request.GroupName) {
		body["groupName"] = request.GroupName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroup"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a group.
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(instanceId *string, applicationId *string, request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateGroupHeaders{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an organizational unit.
//
// @param request - CreateOrganizationalUnitRequest
//
// @param headers - CreateOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnitWithOptions(instanceId *string, applicationId *string, request *CreateOrganizationalUnitRequest, headers *CreateOrganizationalUnitHeaders, runtime *dara.RuntimeOptions) (_result *CreateOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.OrganizationalUnitExternalId) {
		body["organizationalUnitExternalId"] = request.OrganizationalUnitExternalId
	}

	if !dara.IsNil(request.OrganizationalUnitName) {
		body["organizationalUnitName"] = request.OrganizationalUnitName
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrganizationalUnit"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/organizationalUnits"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an organizational unit.
//
// @param request - CreateOrganizationalUnitRequest
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnit(instanceId *string, applicationId *string, request *CreateOrganizationalUnitRequest) (_result *CreateOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateOrganizationalUnitHeaders{}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.CreateOrganizationalUnitWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Employee Identity and Access Management (EIAM) account in an organizational unit.
//
// @param request - CreateUserRequest
//
// @param headers - CreateUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(instanceId *string, applicationId *string, request *CreateUserRequest, headers *CreateUserHeaders, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomFields) {
		body["customFields"] = request.CustomFields
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.EmailVerified) {
		body["emailVerified"] = request.EmailVerified
	}

	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	if !dara.IsNil(request.PasswordInitializationConfig) {
		body["passwordInitializationConfig"] = request.PasswordInitializationConfig
	}

	if !dara.IsNil(request.PhoneNumber) {
		body["phoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PhoneNumberVerified) {
		body["phoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !dara.IsNil(request.PhoneRegion) {
		body["phoneRegion"] = request.PhoneRegion
	}

	if !dara.IsNil(request.PrimaryOrganizationalUnitId) {
		body["primaryOrganizationalUnitId"] = request.PrimaryOrganizationalUnitId
	}

	if !dara.IsNil(request.UserExternalId) {
		body["userExternalId"] = request.UserExternalId
	}

	if !dara.IsNil(request.Username) {
		body["username"] = request.Username
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Employee Identity and Access Management (EIAM) account in an organizational unit.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(instanceId *string, applicationId *string, request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateUserHeaders{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a group.
//
// @param headers - DeleteGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(instanceId *string, applicationId *string, groupId *string, headers *DeleteGroupHeaders, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroup"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a group.
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(instanceId *string, applicationId *string, groupId *string) (_result *DeleteGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteGroupHeaders{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(instanceId, applicationId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an organizational unit.
//
// @param headers - DeleteOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *DeleteOrganizationalUnitHeaders, runtime *dara.RuntimeOptions) (_result *DeleteOrganizationalUnitResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOrganizationalUnit"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/organizationalUnits/" + dara.PercentEncode(dara.StringValue(organizationalUnitId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an organizational unit.
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *DeleteOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteOrganizationalUnitHeaders{}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.DeleteOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) account.
//
// @param headers - DeleteUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *DeleteUserHeaders, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) account.
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(instanceId *string, applicationId *string, userId *string) (_result *DeleteUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteUserHeaders{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an Employee Identity and Access Management (EIAM) account.
//
// @param headers - DisableUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserResponse
func (client *Client) DisableUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *DisableUserHeaders, runtime *dara.RuntimeOptions) (_result *DisableUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableUser"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId)) + "/actions/disable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DisableUserResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an Employee Identity and Access Management (EIAM) account.
//
// @return DisableUserResponse
func (client *Client) DisableUser(instanceId *string, applicationId *string, userId *string) (_result *DisableUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DisableUserHeaders{}
	_result = &DisableUserResponse{}
	_body, _err := client.DisableUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an Employee Identity and Access Management (EIAM) account.
//
// @param headers - EnableUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableUserResponse
func (client *Client) EnableUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *EnableUserHeaders, runtime *dara.RuntimeOptions) (_result *EnableUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableUser"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId)) + "/actions/enable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &EnableUserResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an Employee Identity and Access Management (EIAM) account.
//
// @return EnableUserResponse
func (client *Client) EnableUser(instanceId *string, applicationId *string, userId *string) (_result *EnableUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &EnableUserHeaders{}
	_result = &EnableUserResponse{}
	_body, _err := client.EnableUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a device code.
//
// @param request - GenerateDeviceCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDeviceCodeResponse
func (client *Client) GenerateDeviceCodeWithOptions(instanceId *string, applicationId *string, request *GenerateDeviceCodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateDeviceCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Scope) {
		query["scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateDeviceCode"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/oauth2/device/code"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateDeviceCodeResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a device code.
//
// @param request - GenerateDeviceCodeRequest
//
// @return GenerateDeviceCodeResponse
func (client *Client) GenerateDeviceCode(instanceId *string, applicationId *string, request *GenerateDeviceCodeRequest) (_result *GenerateDeviceCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateDeviceCodeResponse{}
	_body, _err := client.GenerateDeviceCodeWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a token for accessing an application in an instance.
//
// Description:
//
// The following authorization types are supported: authorization code, device code, refresh token, and client credentials.
//
// @param request - GenerateTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTokenResponse
func (client *Client) GenerateTokenWithOptions(instanceId *string, applicationId *string, request *GenerateTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["client_id"] = request.ClientId
	}

	if !dara.IsNil(request.ClientSecret) {
		query["client_secret"] = request.ClientSecret
	}

	if !dara.IsNil(request.Code) {
		query["code"] = request.Code
	}

	if !dara.IsNil(request.CodeVerifier) {
		query["code_verifier"] = request.CodeVerifier
	}

	if !dara.IsNil(request.DeviceCode) {
		query["device_code"] = request.DeviceCode
	}

	if !dara.IsNil(request.ExclusiveTag) {
		query["exclusive_tag"] = request.ExclusiveTag
	}

	if !dara.IsNil(request.GrantType) {
		query["grant_type"] = request.GrantType
	}

	if !dara.IsNil(request.Password) {
		query["password"] = request.Password
	}

	if !dara.IsNil(request.RedirectUri) {
		query["redirect_uri"] = request.RedirectUri
	}

	if !dara.IsNil(request.RefreshToken) {
		query["refresh_token"] = request.RefreshToken
	}

	if !dara.IsNil(request.Scope) {
		query["scope"] = request.Scope
	}

	if !dara.IsNil(request.Username) {
		query["username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateToken"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/oauth2/token"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateTokenResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a token for accessing an application in an instance.
//
// Description:
//
// The following authorization types are supported: authorization code, device code, refresh token, and client credentials.
//
// @param request - GenerateTokenRequest
//
// @return GenerateTokenResponse
func (client *Client) GenerateToken(instanceId *string, applicationId *string, request *GenerateTokenRequest) (_result *GenerateTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateTokenResponse{}
	_body, _err := client.GenerateTokenWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the synchronization scope of an application in an instance.
//
// Description:
//
// >
//
//   - You can go to the Applications page in the IDaaS console to set the synchronization scope. After an application is created, the application has the permission to call this operation by default.
//
// @param headers - GetApplicationProvisioningScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScopeWithOptions(instanceId *string, applicationId *string, headers *GetApplicationProvisioningScopeHeaders, runtime *dara.RuntimeOptions) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationProvisioningScope"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/provisioningScope"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the synchronization scope of an application in an instance.
//
// Description:
//
// >
//
//   - You can go to the Applications page in the IDaaS console to set the synchronization scope. After an application is created, the application has the permission to call this operation by default.
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScope(instanceId *string, applicationId *string) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetApplicationProvisioningScopeHeaders{}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.GetApplicationProvisioningScopeWithOptions(instanceId, applicationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a group.
//
// @param headers - GetGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(instanceId *string, applicationId *string, groupId *string, headers *GetGroupHeaders, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroup"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a group.
//
// @return GetGroupResponse
func (client *Client) GetGroup(instanceId *string, applicationId *string, groupId *string) (_result *GetGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetGroupHeaders{}
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(instanceId, applicationId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an organizational unit.
//
// @param headers - GetOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *GetOrganizationalUnitHeaders, runtime *dara.RuntimeOptions) (_result *GetOrganizationalUnitResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrganizationalUnit"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/organizationalUnits/" + dara.PercentEncode(dara.StringValue(organizationalUnitId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an organizational unit.
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *GetOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetOrganizationalUnitHeaders{}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.GetOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtains the ID of an organizational unit based on the external ID
//
// @param request - GetOrganizationalUnitIdByExternalIdRequest
//
// @param headers - GetOrganizationalUnitIdByExternalIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationalUnitIdByExternalIdResponse
func (client *Client) GetOrganizationalUnitIdByExternalIdWithOptions(instanceId *string, applicationId *string, request *GetOrganizationalUnitIdByExternalIdRequest, headers *GetOrganizationalUnitIdByExternalIdHeaders, runtime *dara.RuntimeOptions) (_result *GetOrganizationalUnitIdByExternalIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationalUnitExternalId) {
		body["organizationalUnitExternalId"] = request.OrganizationalUnitExternalId
	}

	if !dara.IsNil(request.OrganizationalUnitSourceId) {
		body["organizationalUnitSourceId"] = request.OrganizationalUnitSourceId
	}

	if !dara.IsNil(request.OrganizationalUnitSourceType) {
		body["organizationalUnitSourceType"] = request.OrganizationalUnitSourceType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrganizationalUnitIdByExternalId"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/organizationalUnits/_/actions/getOrganizationalUnitIdByExternalId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrganizationalUnitIdByExternalIdResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtains the ID of an organizational unit based on the external ID
//
// @param request - GetOrganizationalUnitIdByExternalIdRequest
//
// @return GetOrganizationalUnitIdByExternalIdResponse
func (client *Client) GetOrganizationalUnitIdByExternalId(instanceId *string, applicationId *string, request *GetOrganizationalUnitIdByExternalIdRequest) (_result *GetOrganizationalUnitIdByExternalIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetOrganizationalUnitIdByExternalIdHeaders{}
	_result = &GetOrganizationalUnitIdByExternalIdResponse{}
	_body, _err := client.GetOrganizationalUnitIdByExternalIdWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an Employee Identity and Access Management (EIAM) account.
//
// @param headers - GetUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(instanceId *string, applicationId *string, userId *string, headers *GetUserHeaders, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an Employee Identity and Access Management (EIAM) account.
//
// @return GetUserResponse
func (client *Client) GetUser(instanceId *string, applicationId *string, userId *string) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(instanceId, applicationId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account by email address.
//
// @param request - GetUserIdByEmailRequest
//
// @param headers - GetUserIdByEmailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByEmailResponse
func (client *Client) GetUserIdByEmailWithOptions(instanceId *string, applicationId *string, request *GetUserIdByEmailRequest, headers *GetUserIdByEmailHeaders, runtime *dara.RuntimeOptions) (_result *GetUserIdByEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserIdByEmail"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/_/actions/getUserIdByEmail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserIdByEmailResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account by email address.
//
// @param request - GetUserIdByEmailRequest
//
// @return GetUserIdByEmailResponse
func (client *Client) GetUserIdByEmail(instanceId *string, applicationId *string, request *GetUserIdByEmailRequest) (_result *GetUserIdByEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserIdByEmailHeaders{}
	_result = &GetUserIdByEmailResponse{}
	_body, _err := client.GetUserIdByEmailWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the mobile number.
//
// @param request - GetUserIdByPhoneNumberRequest
//
// @param headers - GetUserIdByPhoneNumberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByPhoneNumberResponse
func (client *Client) GetUserIdByPhoneNumberWithOptions(instanceId *string, applicationId *string, request *GetUserIdByPhoneNumberRequest, headers *GetUserIdByPhoneNumberHeaders, runtime *dara.RuntimeOptions) (_result *GetUserIdByPhoneNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PhoneNumber) {
		body["phoneNumber"] = request.PhoneNumber
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserIdByPhoneNumber"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/_/actions/getUserIdByPhoneNumber"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserIdByPhoneNumberResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the mobile number.
//
// @param request - GetUserIdByPhoneNumberRequest
//
// @return GetUserIdByPhoneNumberResponse
func (client *Client) GetUserIdByPhoneNumber(instanceId *string, applicationId *string, request *GetUserIdByPhoneNumberRequest) (_result *GetUserIdByPhoneNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserIdByPhoneNumberHeaders{}
	_result = &GetUserIdByPhoneNumberResponse{}
	_body, _err := client.GetUserIdByPhoneNumberWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the external ID.
//
// @param request - GetUserIdByUserExternalIdRequest
//
// @param headers - GetUserIdByUserExternalIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByUserExternalIdResponse
func (client *Client) GetUserIdByUserExternalIdWithOptions(instanceId *string, applicationId *string, request *GetUserIdByUserExternalIdRequest, headers *GetUserIdByUserExternalIdHeaders, runtime *dara.RuntimeOptions) (_result *GetUserIdByUserExternalIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserExternalId) {
		body["userExternalId"] = request.UserExternalId
	}

	if !dara.IsNil(request.UserSourceId) {
		body["userSourceId"] = request.UserSourceId
	}

	if !dara.IsNil(request.UserSourceType) {
		body["userSourceType"] = request.UserSourceType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserIdByUserExternalId"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/_/actions/getUserIdByExternalId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserIdByUserExternalIdResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the external ID.
//
// @param request - GetUserIdByUserExternalIdRequest
//
// @return GetUserIdByUserExternalIdResponse
func (client *Client) GetUserIdByUserExternalId(instanceId *string, applicationId *string, request *GetUserIdByUserExternalIdRequest) (_result *GetUserIdByUserExternalIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserIdByUserExternalIdHeaders{}
	_result = &GetUserIdByUserExternalIdResponse{}
	_body, _err := client.GetUserIdByUserExternalIdWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the username.
//
// @param request - GetUserIdByUsernameRequest
//
// @param headers - GetUserIdByUsernameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserIdByUsernameResponse
func (client *Client) GetUserIdByUsernameWithOptions(instanceId *string, applicationId *string, request *GetUserIdByUsernameRequest, headers *GetUserIdByUsernameHeaders, runtime *dara.RuntimeOptions) (_result *GetUserIdByUsernameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Username) {
		body["username"] = request.Username
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserIdByUsername"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/_/actions/getUserIdByUsername"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserIdByUsernameResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ID of an Employee Identity and Access Management (EIAM) account based on the username.
//
// @param request - GetUserIdByUsernameRequest
//
// @return GetUserIdByUsernameResponse
func (client *Client) GetUserIdByUsername(instanceId *string, applicationId *string, request *GetUserIdByUsernameRequest) (_result *GetUserIdByUsernameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserIdByUsernameHeaders{}
	_result = &GetUserIdByUsernameResponse{}
	_body, _err := client.GetUserIdByUsernameWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a user by using the user token.
//
// @param headers - GetUserInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfoWithOptions(instanceId *string, applicationId *string, headers *GetUserInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserInfo"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/oauth2/userinfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a user by using the user token.
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfo(instanceId *string, applicationId *string) (_result *GetUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserInfoHeaders{}
	_result = &GetUserInfoResponse{}
	_body, _err := client.GetUserInfoWithOptions(instanceId, applicationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about Employee Identity and Access Management (EIAM) groups by page.
//
// @param request - ListGroupsRequest
//
// @param headers - ListGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(instanceId *string, applicationId *string, request *ListGroupsRequest, headers *ListGroupsHeaders, runtime *dara.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupNameStartWith) {
		query["groupNameStartWith"] = request.GroupNameStartWith
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroups"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about Employee Identity and Access Management (EIAM) groups by page.
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(instanceId *string, applicationId *string, request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListGroupsHeaders{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取账户关联组列表
//
// @param request - ListGroupsForUserRequest
//
// @param headers - ListGroupsForUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithOptions(instanceId *string, applicationId *string, userId *string, request *ListGroupsForUserRequest, headers *ListGroupsForUserHeaders, runtime *dara.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupsForUser"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId)) + "/actions/listGroupsForUser"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取账户关联组列表
//
// @param request - ListGroupsForUserRequest
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUser(instanceId *string, applicationId *string, userId *string, request *ListGroupsForUserRequest) (_result *ListGroupsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListGroupsForUserHeaders{}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.ListGroupsForUserWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of all the parent organizational units of an organizational unit.
//
// @param headers - ListOrganizationalUnitParentIdsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitParentIdsResponse
func (client *Client) ListOrganizationalUnitParentIdsWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, headers *ListOrganizationalUnitParentIdsHeaders, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitParentIdsResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationalUnitParentIds"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/organizationalUnits/" + dara.PercentEncode(dara.StringValue(organizationalUnitId)) + "/parentIds"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationalUnitParentIdsResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of all the parent organizational units of an organizational unit.
//
// @return ListOrganizationalUnitParentIdsResponse
func (client *Client) ListOrganizationalUnitParentIds(instanceId *string, applicationId *string, organizationalUnitId *string) (_result *ListOrganizationalUnitParentIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListOrganizationalUnitParentIdsHeaders{}
	_result = &ListOrganizationalUnitParentIdsResponse{}
	_body, _err := client.ListOrganizationalUnitParentIdsWithOptions(instanceId, applicationId, organizationalUnitId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of Employee Identity and Access Management (EIAM) organizational units by page.
//
// @param request - ListOrganizationalUnitsRequest
//
// @param headers - ListOrganizationalUnitsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnitsWithOptions(instanceId *string, applicationId *string, request *ListOrganizationalUnitsRequest, headers *ListOrganizationalUnitsHeaders, runtime *dara.RuntimeOptions) (_result *ListOrganizationalUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentId) {
		query["parentId"] = request.ParentId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationalUnits"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/organizationalUnits"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of Employee Identity and Access Management (EIAM) organizational units by page.
//
// @param request - ListOrganizationalUnitsRequest
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnits(instanceId *string, applicationId *string, request *ListOrganizationalUnitsRequest) (_result *ListOrganizationalUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListOrganizationalUnitsHeaders{}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.ListOrganizationalUnitsWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of Employee Identity and Access Management (EIAM) accounts by page.
//
// @param request - ListUsersRequest
//
// @param headers - ListUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(instanceId *string, applicationId *string, request *ListUsersRequest, headers *ListUsersHeaders, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationalUnitId) {
		query["organizationalUnitId"] = request.OrganizationalUnitId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of Employee Identity and Access Management (EIAM) accounts by page.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(instanceId *string, applicationId *string, request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListUsersHeaders{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries accounts in an Employee Identity and Access Management (EIAM) group.
//
// @param request - ListUsersForGroupRequest
//
// @param headers - ListUsersForGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithOptions(instanceId *string, applicationId *string, groupId *string, request *ListUsersForGroupRequest, headers *ListUsersForGroupHeaders, runtime *dara.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsersForGroup"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/actions/listUsersForGroup"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries accounts in an Employee Identity and Access Management (EIAM) group.
//
// @param request - ListUsersForGroupRequest
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroup(instanceId *string, applicationId *string, groupId *string, request *ListUsersForGroupRequest) (_result *ListUsersForGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListUsersForGroupHeaders{}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.ListUsersForGroupWithOptions(instanceId, applicationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云角色（CloudAccountRole）的临时访问凭证
//
// @param request - ObtainCloudAccountRoleAccessCredentialRequest
//
// @param headers - ObtainCloudAccountRoleAccessCredentialHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainCloudAccountRoleAccessCredentialResponse
func (client *Client) ObtainCloudAccountRoleAccessCredentialWithOptions(instanceId *string, request *ObtainCloudAccountRoleAccessCredentialRequest, headers *ObtainCloudAccountRoleAccessCredentialHeaders, runtime *dara.RuntimeOptions) (_result *ObtainCloudAccountRoleAccessCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudAccountRoleExternalId) {
		query["cloudAccountRoleExternalId"] = request.CloudAccountRoleExternalId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ObtainCloudAccountRoleAccessCredential"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/cloudAccountRoles/_/actions/obtainAccessCredential"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ObtainCloudAccountRoleAccessCredentialResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云角色（CloudAccountRole）的临时访问凭证
//
// @param request - ObtainCloudAccountRoleAccessCredentialRequest
//
// @return ObtainCloudAccountRoleAccessCredentialResponse
func (client *Client) ObtainCloudAccountRoleAccessCredential(instanceId *string, request *ObtainCloudAccountRoleAccessCredentialRequest) (_result *ObtainCloudAccountRoleAccessCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ObtainCloudAccountRoleAccessCredentialHeaders{}
	_result = &ObtainCloudAccountRoleAccessCredentialResponse{}
	_body, _err := client.ObtainCloudAccountRoleAccessCredentialWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about an Employee Identity and Access Management (EIAM) group.
//
// @param request - PatchGroupRequest
//
// @param headers - PatchGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchGroupResponse
func (client *Client) PatchGroupWithOptions(instanceId *string, applicationId *string, groupId *string, request *PatchGroupRequest, headers *PatchGroupHeaders, runtime *dara.RuntimeOptions) (_result *PatchGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["groupName"] = request.GroupName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PatchGroup"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &PatchGroupResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about an Employee Identity and Access Management (EIAM) group.
//
// @param request - PatchGroupRequest
//
// @return PatchGroupResponse
func (client *Client) PatchGroup(instanceId *string, applicationId *string, groupId *string, request *PatchGroupRequest) (_result *PatchGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PatchGroupHeaders{}
	_result = &PatchGroupResponse{}
	_body, _err := client.PatchGroupWithOptions(instanceId, applicationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an EIAM organizational unit.
//
// Description:
//
// The operation conforms to the HTTP PATCH request method. The value of a parameter is modified only if the parameter is specified in the request.
//
// @param request - PatchOrganizationalUnitRequest
//
// @param headers - PatchOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchOrganizationalUnitResponse
func (client *Client) PatchOrganizationalUnitWithOptions(instanceId *string, applicationId *string, organizationalUnitId *string, request *PatchOrganizationalUnitRequest, headers *PatchOrganizationalUnitHeaders, runtime *dara.RuntimeOptions) (_result *PatchOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.OrganizationalUnitName) {
		body["organizationalUnitName"] = request.OrganizationalUnitName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PatchOrganizationalUnit"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/organizationalUnits/" + dara.PercentEncode(dara.StringValue(organizationalUnitId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &PatchOrganizationalUnitResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an EIAM organizational unit.
//
// Description:
//
// The operation conforms to the HTTP PATCH request method. The value of a parameter is modified only if the parameter is specified in the request.
//
// @param request - PatchOrganizationalUnitRequest
//
// @return PatchOrganizationalUnitResponse
func (client *Client) PatchOrganizationalUnit(instanceId *string, applicationId *string, organizationalUnitId *string, request *PatchOrganizationalUnitRequest) (_result *PatchOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PatchOrganizationalUnitHeaders{}
	_result = &PatchOrganizationalUnitResponse{}
	_body, _err := client.PatchOrganizationalUnitWithOptions(instanceId, applicationId, organizationalUnitId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an Employee Identity and Access Management (EIAM) account.
//
// Description:
//
// The operation conforms to the HTTP PATCH request method. The value of a parameter is modified only if the parameter is specified in the request.
//
// @param request - PatchUserRequest
//
// @param headers - PatchUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PatchUserResponse
func (client *Client) PatchUserWithOptions(instanceId *string, applicationId *string, userId *string, request *PatchUserRequest, headers *PatchUserHeaders, runtime *dara.RuntimeOptions) (_result *PatchUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomFields) {
		body["customFields"] = request.CustomFields
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.EmailVerified) {
		body["emailVerified"] = request.EmailVerified
	}

	if !dara.IsNil(request.PhoneNumber) {
		body["phoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.PhoneNumberVerified) {
		body["phoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !dara.IsNil(request.PhoneRegion) {
		body["phoneRegion"] = request.PhoneRegion
	}

	if !dara.IsNil(request.Username) {
		body["username"] = request.Username
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PatchUser"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &PatchUserResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an Employee Identity and Access Management (EIAM) account.
//
// Description:
//
// The operation conforms to the HTTP PATCH request method. The value of a parameter is modified only if the parameter is specified in the request.
//
// @param request - PatchUserRequest
//
// @return PatchUserResponse
func (client *Client) PatchUser(instanceId *string, applicationId *string, userId *string, request *PatchUserRequest) (_result *PatchUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PatchUserHeaders{}
	_result = &PatchUserResponse{}
	_body, _err := client.PatchUserWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将账户从多个组织移除【不支持移除主组织】
//
// @param request - RemoveUserFromOrganizationalUnitsRequest
//
// @param headers - RemoveUserFromOrganizationalUnitsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnitsWithOptions(instanceId *string, applicationId *string, userId *string, request *RemoveUserFromOrganizationalUnitsRequest, headers *RemoveUserFromOrganizationalUnitsHeaders, runtime *dara.RuntimeOptions) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationalUnitIds) {
		body["organizationalUnitIds"] = request.OrganizationalUnitIds
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserFromOrganizationalUnits"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId)) + "/actions/removeUserFromOrganizationalUnits"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &RemoveUserFromOrganizationalUnitsResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将账户从多个组织移除【不支持移除主组织】
//
// @param request - RemoveUserFromOrganizationalUnitsRequest
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnits(instanceId *string, applicationId *string, userId *string, request *RemoveUserFromOrganizationalUnitsRequest) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RemoveUserFromOrganizationalUnitsHeaders{}
	_result = &RemoveUserFromOrganizationalUnitsResponse{}
	_body, _err := client.RemoveUserFromOrganizationalUnitsWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes multiple Employee Identity and Access Management (EIAM) accounts from an EIAM group. If an account does not belong to the group, the removal succeeds by default.
//
// @param request - RemoveUsersFromGroupRequest
//
// @param headers - RemoveUsersFromGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroupWithOptions(instanceId *string, applicationId *string, groupId *string, request *RemoveUsersFromGroupRequest, headers *RemoveUsersFromGroupHeaders, runtime *dara.RuntimeOptions) (_result *RemoveUsersFromGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserIds) {
		body["userIds"] = request.UserIds
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsersFromGroup"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/actions/removeUsersFromGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes multiple Employee Identity and Access Management (EIAM) accounts from an EIAM group. If an account does not belong to the group, the removal succeeds by default.
//
// @param request - RemoveUsersFromGroupRequest
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroup(instanceId *string, applicationId *string, groupId *string, request *RemoveUsersFromGroupRequest) (_result *RemoveUsersFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &RemoveUsersFromGroupHeaders{}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.RemoveUsersFromGroupWithOptions(instanceId, applicationId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes an access token or refresh token.
//
// @param request - RevokeTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTokenResponse
func (client *Client) RevokeTokenWithOptions(instanceId *string, applicationId *string, request *RevokeTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["client_id"] = request.ClientId
	}

	if !dara.IsNil(request.ClientSecret) {
		query["client_secret"] = request.ClientSecret
	}

	if !dara.IsNil(request.Token) {
		query["token"] = request.Token
	}

	if !dara.IsNil(request.TokenTypeHint) {
		query["token_type_hint"] = request.TokenTypeHint
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeToken"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/oauth2/revoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeTokenResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes an access token or refresh token.
//
// @param request - RevokeTokenRequest
//
// @return RevokeTokenResponse
func (client *Client) RevokeToken(instanceId *string, applicationId *string, request *RevokeTokenRequest) (_result *RevokeTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeTokenResponse{}
	_body, _err := client.RevokeTokenWithOptions(instanceId, applicationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将指定组织设置为账户主组织，移除旧主组织，加入新主组织。
//
// @param request - SetUserPrimaryOrganizationalUnitRequest
//
// @param headers - SetUserPrimaryOrganizationalUnitHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnitWithOptions(instanceId *string, applicationId *string, userId *string, request *SetUserPrimaryOrganizationalUnitRequest, headers *SetUserPrimaryOrganizationalUnitHeaders, runtime *dara.RuntimeOptions) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationalUnitId) {
		body["organizationalUnitId"] = request.OrganizationalUnitId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetUserPrimaryOrganizationalUnit"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId)) + "/actions/setUserPrimaryOrganizationalUnit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &SetUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将指定组织设置为账户主组织，移除旧主组织，加入新主组织。
//
// @param request - SetUserPrimaryOrganizationalUnitRequest
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnit(instanceId *string, applicationId *string, userId *string, request *SetUserPrimaryOrganizationalUnitRequest) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SetUserPrimaryOrganizationalUnitHeaders{}
	_result = &SetUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.SetUserPrimaryOrganizationalUnitWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新账户密码
//
// @param request - UpdateUserPasswordRequest
//
// @param headers - UpdateUserPasswordHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPasswordWithOptions(instanceId *string, applicationId *string, userId *string, request *UpdateUserPasswordRequest, headers *UpdateUserPasswordHeaders, runtime *dara.RuntimeOptions) (_result *UpdateUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserPassword"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/" + dara.PercentEncode(dara.StringValue(applicationId)) + "/users/" + dara.PercentEncode(dara.StringValue(userId)) + "/actions/updateUserPassword"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &UpdateUserPasswordResponse{}
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新账户密码
//
// @param request - UpdateUserPasswordRequest
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPassword(instanceId *string, applicationId *string, userId *string, request *UpdateUserPasswordRequest) (_result *UpdateUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateUserPasswordHeaders{}
	_result = &UpdateUserPasswordResponse{}
	_body, _err := client.UpdateUserPasswordWithOptions(instanceId, applicationId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
