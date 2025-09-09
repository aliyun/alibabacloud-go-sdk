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
	client.Endpoint, _err = client.GetEndpoint(dara.String("yundun-bastionhost"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// If an O\\&M engineer attempts to run a command specified in the Command Approval field on the Create Control Policy page, the administrator is notified to review the command in the Bastionhost console. The command can be run only after it is approved by the administrator.
//
// Description:
//
// You can call this operation as a Bastionhost administrator to approve the request to run a command of an O\\&M engineer.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AcceptApproveCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptApproveCommandResponse
func (client *Client) AcceptApproveCommandWithOptions(request *AcceptApproveCommandRequest, runtime *dara.RuntimeOptions) (_result *AcceptApproveCommandResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandId) {
		query["CommandId"] = request.CommandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptApproveCommand"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptApproveCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If an O\\&M engineer attempts to run a command specified in the Command Approval field on the Create Control Policy page, the administrator is notified to review the command in the Bastionhost console. The command can be run only after it is approved by the administrator.
//
// Description:
//
// You can call this operation as a Bastionhost administrator to approve the request to run a command of an O\\&M engineer.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AcceptApproveCommandRequest
//
// @return AcceptApproveCommandResponse
func (client *Client) AcceptApproveCommand(request *AcceptApproveCommandRequest) (_result *AcceptApproveCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptApproveCommandResponse{}
	_body, _err := client.AcceptApproveCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Approves an O\\\\\\\\\\\\&M application.
//
// Description:
//
// You can call this operation as a Bastionhost administrator to approve an O\\&M application of an O\\&M engineer.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AcceptOperationTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptOperationTicketResponse
func (client *Client) AcceptOperationTicketWithOptions(request *AcceptOperationTicketRequest, runtime *dara.RuntimeOptions) (_result *AcceptOperationTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EffectCount) {
		query["EffectCount"] = request.EffectCount
	}

	if !dara.IsNil(request.EffectEndTime) {
		query["EffectEndTime"] = request.EffectEndTime
	}

	if !dara.IsNil(request.EffectStartTime) {
		query["EffectStartTime"] = request.EffectStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OperationTicketId) {
		query["OperationTicketId"] = request.OperationTicketId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptOperationTicket"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptOperationTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Approves an O\\\\\\\\\\\\&M application.
//
// Description:
//
// You can call this operation as a Bastionhost administrator to approve an O\\&M application of an O\\&M engineer.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AcceptOperationTicketRequest
//
// @return AcceptOperationTicketResponse
func (client *Client) AcceptOperationTicket(request *AcceptOperationTicketRequest) (_result *AcceptOperationTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptOperationTicketResponse{}
	_body, _err := client.AcceptOperationTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds multiple databases to a specified asset group.
//
// @param request - AddDatabasesToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDatabasesToGroupResponse
func (client *Client) AddDatabasesToGroupWithOptions(request *AddDatabasesToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddDatabasesToGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseIds) {
		query["DatabaseIds"] = request.DatabaseIds
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDatabasesToGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDatabasesToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds multiple databases to a specified asset group.
//
// @param request - AddDatabasesToGroupRequest
//
// @return AddDatabasesToGroupResponse
func (client *Client) AddDatabasesToGroup(request *AddDatabasesToGroupRequest) (_result *AddDatabasesToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDatabasesToGroupResponse{}
	_body, _err := client.AddDatabasesToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds one or more hosts to the specified host group.
//
// Description:
//
// You can call this operation to add one or more hosts to a host group. You can add multiple hosts to a host group to manage and grant permissions on the hosts in a centralized manner.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
//
// @param request - AddHostsToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddHostsToGroupResponse
func (client *Client) AddHostsToGroupWithOptions(request *AddHostsToGroupRequest, runtime *dara.RuntimeOptions) (_result *AddHostsToGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddHostsToGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddHostsToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more hosts to the specified host group.
//
// Description:
//
// You can call this operation to add one or more hosts to a host group. You can add multiple hosts to a host group to manage and grant permissions on the hosts in a centralized manner.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
//
// @param request - AddHostsToGroupRequest
//
// @return AddHostsToGroupResponse
func (client *Client) AddHostsToGroup(request *AddHostsToGroupRequest) (_result *AddHostsToGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddHostsToGroupResponse{}
	_body, _err := client.AddHostsToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加RD成员账号
//
// @param request - AddInstanceRdMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddInstanceRdMemberResponse
func (client *Client) AddInstanceRdMemberWithOptions(request *AddInstanceRdMemberRequest, runtime *dara.RuntimeOptions) (_result *AddInstanceRdMemberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddInstanceRdMember"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddInstanceRdMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加RD成员账号
//
// @param request - AddInstanceRdMemberRequest
//
// @return AddInstanceRdMemberResponse
func (client *Client) AddInstanceRdMember(request *AddInstanceRdMemberRequest) (_result *AddInstanceRdMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddInstanceRdMemberResponse{}
	_body, _err := client.AddInstanceRdMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add one or more users to a user group.
//
// Description:
//
// #
//
// You can call this operation to add one or more users to a user group. After you call the [CreateUserGroup](https://help.aliyun.com/document_detail/204596.html) operation to create a user group, you can call the AddUsersToGroup operation to add multiple users to the user group. Then, you can manage and grant permissions to the users at a time.
//
// # Limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUsersToGroup"),
		Version:     dara.String("2019-12-09"),
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
// Add one or more users to a user group.
//
// Description:
//
// #
//
// You can call this operation to add one or more users to a user group. After you call the [CreateUserGroup](https://help.aliyun.com/document_detail/204596.html) operation to create a user group, you can call the AddUsersToGroup operation to add multiple users to the user group. Then, you can manage and grant permissions to the users at a time.
//
// # Limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Authorizes a user to manage databases and database accounts.
//
// @param request - AttachDatabaseAccountsToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDatabaseAccountsToUserResponse
func (client *Client) AttachDatabaseAccountsToUserWithOptions(request *AttachDatabaseAccountsToUserRequest, runtime *dara.RuntimeOptions) (_result *AttachDatabaseAccountsToUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDatabaseAccountsToUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDatabaseAccountsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage databases and database accounts.
//
// @param request - AttachDatabaseAccountsToUserRequest
//
// @return AttachDatabaseAccountsToUserResponse
func (client *Client) AttachDatabaseAccountsToUser(request *AttachDatabaseAccountsToUserRequest) (_result *AttachDatabaseAccountsToUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachDatabaseAccountsToUserResponse{}
	_body, _err := client.AttachDatabaseAccountsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes a user group to manage databases and database accounts.
//
// @param request - AttachDatabaseAccountsToUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDatabaseAccountsToUserGroupResponse
func (client *Client) AttachDatabaseAccountsToUserGroupWithOptions(request *AttachDatabaseAccountsToUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AttachDatabaseAccountsToUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDatabaseAccountsToUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDatabaseAccountsToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a user group to manage databases and database accounts.
//
// @param request - AttachDatabaseAccountsToUserGroupRequest
//
// @return AttachDatabaseAccountsToUserGroupResponse
func (client *Client) AttachDatabaseAccountsToUserGroup(request *AttachDatabaseAccountsToUserGroupRequest) (_result *AttachDatabaseAccountsToUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachDatabaseAccountsToUserGroupResponse{}
	_body, _err := client.AttachDatabaseAccountsToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates host accounts with a shared key.
//
// @param request - AttachHostAccountsToHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostAccountsToHostShareKeyResponse
func (client *Client) AttachHostAccountsToHostShareKeyWithOptions(request *AttachHostAccountsToHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *AttachHostAccountsToHostShareKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountIds) {
		query["HostAccountIds"] = request.HostAccountIds
	}

	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostAccountsToHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostAccountsToHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates host accounts with a shared key.
//
// @param request - AttachHostAccountsToHostShareKeyRequest
//
// @return AttachHostAccountsToHostShareKeyResponse
func (client *Client) AttachHostAccountsToHostShareKey(request *AttachHostAccountsToHostShareKeyRequest) (_result *AttachHostAccountsToHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostAccountsToHostShareKeyResponse{}
	_body, _err := client.AttachHostAccountsToHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage the hosts and host accounts.
//
// @param request - AttachHostAccountsToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostAccountsToUserResponse
func (client *Client) AttachHostAccountsToUserWithOptions(request *AttachHostAccountsToUserRequest, runtime *dara.RuntimeOptions) (_result *AttachHostAccountsToUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostAccountsToUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostAccountsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage the hosts and host accounts.
//
// @param request - AttachHostAccountsToUserRequest
//
// @return AttachHostAccountsToUserResponse
func (client *Client) AttachHostAccountsToUser(request *AttachHostAccountsToUserRequest) (_result *AttachHostAccountsToUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostAccountsToUserResponse{}
	_body, _err := client.AttachHostAccountsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes a user group to manage one or more hosts and host accounts.
//
// Description:
//
// After you authorize a user group to manage specific hosts and host accounts, all the users in the user group have access to the authorized hosts and host accounts.
//
// @param request - AttachHostAccountsToUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostAccountsToUserGroupResponse
func (client *Client) AttachHostAccountsToUserGroupWithOptions(request *AttachHostAccountsToUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AttachHostAccountsToUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostAccountsToUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostAccountsToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a user group to manage one or more hosts and host accounts.
//
// Description:
//
// After you authorize a user group to manage specific hosts and host accounts, all the users in the user group have access to the authorized hosts and host accounts.
//
// @param request - AttachHostAccountsToUserGroupRequest
//
// @return AttachHostAccountsToUserGroupResponse
func (client *Client) AttachHostAccountsToUserGroup(request *AttachHostAccountsToUserGroupRequest) (_result *AttachHostAccountsToUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostAccountsToUserGroupResponse{}
	_body, _err := client.AttachHostAccountsToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage one or more host groups and host accounts.
//
// @param request - AttachHostGroupAccountsToUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostGroupAccountsToUserResponse
func (client *Client) AttachHostGroupAccountsToUserWithOptions(request *AttachHostGroupAccountsToUserRequest, runtime *dara.RuntimeOptions) (_result *AttachHostGroupAccountsToUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostGroupAccountsToUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostGroupAccountsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage one or more host groups and host accounts.
//
// @param request - AttachHostGroupAccountsToUserRequest
//
// @return AttachHostGroupAccountsToUserResponse
func (client *Client) AttachHostGroupAccountsToUser(request *AttachHostGroupAccountsToUserRequest) (_result *AttachHostGroupAccountsToUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostGroupAccountsToUserResponse{}
	_body, _err := client.AttachHostGroupAccountsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage one or more host groups and host accounts.
//
// @param request - AttachHostGroupAccountsToUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachHostGroupAccountsToUserGroupResponse
func (client *Client) AttachHostGroupAccountsToUserGroupWithOptions(request *AttachHostGroupAccountsToUserGroupRequest, runtime *dara.RuntimeOptions) (_result *AttachHostGroupAccountsToUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachHostGroupAccountsToUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachHostGroupAccountsToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a user to manage one or more host groups and host accounts.
//
// @param request - AttachHostGroupAccountsToUserGroupRequest
//
// @return AttachHostGroupAccountsToUserGroupResponse
func (client *Client) AttachHostGroupAccountsToUserGroup(request *AttachHostGroupAccountsToUserGroupRequest) (_result *AttachHostGroupAccountsToUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachHostGroupAccountsToUserGroupResponse{}
	_body, _err := client.AttachHostGroupAccountsToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures security groups for a bastion host.
//
// @param request - ConfigInstanceSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceSecurityGroupsResponse
func (client *Client) ConfigInstanceSecurityGroupsWithOptions(request *ConfigInstanceSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *ConfigInstanceSecurityGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedSecurityGroups) {
		query["AuthorizedSecurityGroups"] = request.AuthorizedSecurityGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigInstanceSecurityGroups"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigInstanceSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures security groups for a bastion host.
//
// @param request - ConfigInstanceSecurityGroupsRequest
//
// @return ConfigInstanceSecurityGroupsResponse
func (client *Client) ConfigInstanceSecurityGroups(request *ConfigInstanceSecurityGroupsRequest) (_result *ConfigInstanceSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigInstanceSecurityGroupsResponse{}
	_body, _err := client.ConfigInstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a whitelist of public IP addresses for a bastion host.
//
// Description:
//
// ## Usage notes
//
// You can call this operation to configure a whitelist of public IP addresses for a bastion host. By default, a bastion host is accessible from all public IP addresses. If you want to allow the requests from specific public IP addresses, you can call this operation to add trusted IP addresses to the whitelist of the bastion host.
//
// ## Limits
//
// You can call this operation up to 30 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ConfigInstanceWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceWhiteListResponse
func (client *Client) ConfigInstanceWhiteListWithOptions(request *ConfigInstanceWhiteListRequest, runtime *dara.RuntimeOptions) (_result *ConfigInstanceWhiteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WhiteList) {
		query["WhiteList"] = request.WhiteList
	}

	if !dara.IsNil(request.WhiteListPolicies) {
		query["WhiteListPolicies"] = request.WhiteListPolicies
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigInstanceWhiteList"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a whitelist of public IP addresses for a bastion host.
//
// Description:
//
// ## Usage notes
//
// You can call this operation to configure a whitelist of public IP addresses for a bastion host. By default, a bastion host is accessible from all public IP addresses. If you want to allow the requests from specific public IP addresses, you can call this operation to add trusted IP addresses to the whitelist of the bastion host.
//
// ## Limits
//
// You can call this operation up to 30 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ConfigInstanceWhiteListRequest
//
// @return ConfigInstanceWhiteListResponse
func (client *Client) ConfigInstanceWhiteList(request *ConfigInstanceWhiteListRequest) (_result *ConfigInstanceWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.ConfigInstanceWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports an ApsaraDB RDS for MySQL instance, ApsaraDB RDS for SQL Server instance, ApsaraDB RDS for PostgreSQL instance, PolarDB for MySQL cluster, PolarDB for PostgreSQL cluster, PolarDB for PostgreSQL (Compatible with Oracle) cluster, self-managed MySQL database, self-managed SQL Server database, self-managed PostgreSQL database, or self-managed Oracle database to a bastion host.
//
// @param request - CreateDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabaseWithOptions(request *CreateDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAddressType) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabasePort) {
		query["DatabasePort"] = request.DatabasePort
	}

	if !dara.IsNil(request.DatabasePrivateAddress) {
		query["DatabasePrivateAddress"] = request.DatabasePrivateAddress
	}

	if !dara.IsNil(request.DatabasePublicAddress) {
		query["DatabasePublicAddress"] = request.DatabasePublicAddress
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.PolarDBEndpointType) {
		query["PolarDBEndpointType"] = request.PolarDBEndpointType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourceInstanceRegionId) {
		query["SourceInstanceRegionId"] = request.SourceInstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabase"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports an ApsaraDB RDS for MySQL instance, ApsaraDB RDS for SQL Server instance, ApsaraDB RDS for PostgreSQL instance, PolarDB for MySQL cluster, PolarDB for PostgreSQL cluster, PolarDB for PostgreSQL (Compatible with Oracle) cluster, self-managed MySQL database, self-managed SQL Server database, self-managed PostgreSQL database, or self-managed Oracle database to a bastion host.
//
// @param request - CreateDatabaseRequest
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabase(request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After a database is created, you can create a database account for the database. After the account is created, O\\&M engineers can use the account to log on to and perform O\\&M operations on the database.
//
// @param request - CreateDatabaseAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseAccountResponse
func (client *Client) CreateDatabaseAccountWithOptions(request *CreateDatabaseAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateDatabaseAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.DatabaseSchema) {
		query["DatabaseSchema"] = request.DatabaseSchema
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LoginAttribute) {
		query["LoginAttribute"] = request.LoginAttribute
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabaseAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After a database is created, you can create a database account for the database. After the account is created, O\\&M engineers can use the account to log on to and perform O\\&M operations on the database.
//
// @param request - CreateDatabaseAccountRequest
//
// @return CreateDatabaseAccountResponse
func (client *Client) CreateDatabaseAccount(request *CreateDatabaseAccountRequest) (_result *CreateDatabaseAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatabaseAccountResponse{}
	_body, _err := client.CreateDatabaseAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateExportConfigJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExportConfigJobResponse
func (client *Client) CreateExportConfigJobWithOptions(request *CreateExportConfigJobRequest, runtime *dara.RuntimeOptions) (_result *CreateExportConfigJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExportConfigJob"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExportConfigJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateExportConfigJobRequest
//
// @return CreateExportConfigJobResponse
func (client *Client) CreateExportConfigJob(request *CreateExportConfigJobRequest) (_result *CreateExportConfigJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExportConfigJobResponse{}
	_body, _err := client.CreateExportConfigJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Bastionhost allows you to perform O\\&M operations on hosts from different sources, such as Alibaba Cloud Elastic Compute Service (ECS) instances, servers in on-premises data centers, and servers on other cloud platforms. Before you perform O\\&M operations on hosts by using a bastion host, you must import the hosts to the bastion host. You can call this operation to import a host to a bastion host.
//
// @param request - CreateHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostResponse
func (client *Client) CreateHostWithOptions(request *CreateHostRequest, runtime *dara.RuntimeOptions) (_result *CreateHostResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAddressType) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.HostPrivateAddress) {
		query["HostPrivateAddress"] = request.HostPrivateAddress
	}

	if !dara.IsNil(request.HostPublicAddress) {
		query["HostPublicAddress"] = request.HostPublicAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceRegionId) {
		query["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Bastionhost allows you to perform O\\&M operations on hosts from different sources, such as Alibaba Cloud Elastic Compute Service (ECS) instances, servers in on-premises data centers, and servers on other cloud platforms. Before you perform O\\&M operations on hosts by using a bastion host, you must import the hosts to the bastion host. You can call this operation to import a host to a bastion host.
//
// @param request - CreateHostRequest
//
// @return CreateHostResponse
func (client *Client) CreateHost(request *CreateHostRequest) (_result *CreateHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostResponse{}
	_body, _err := client.CreateHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After you import a host to a bastion host, you must add an account of the host to the bastion host. This way, O\\&M engineers can use the account to log on to and perform O\\&M operations on the host by using the bastion host.
//
// @param request - CreateHostAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostAccountResponse
func (client *Client) CreateHostAccountWithOptions(request *CreateHostAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateHostAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PassPhrase) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.PrivilegeType) {
		query["PrivilegeType"] = request.PrivilegeType
	}

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RotationMode) {
		query["RotationMode"] = request.RotationMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After you import a host to a bastion host, you must add an account of the host to the bastion host. This way, O\\&M engineers can use the account to log on to and perform O\\&M operations on the host by using the bastion host.
//
// @param request - CreateHostAccountRequest
//
// @return CreateHostAccountResponse
func (client *Client) CreateHostAccount(request *CreateHostAccountRequest) (_result *CreateHostAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostAccountResponse{}
	_body, _err := client.CreateHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can create asset groups based on your business requirements and add assets of the same type to an asset group. This allows you to classify assets and manage multiple assets at a time.
//
// @param request - CreateHostGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostGroupResponse
func (client *Client) CreateHostGroupWithOptions(request *CreateHostGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateHostGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can create asset groups based on your business requirements and add assets of the same type to an asset group. This allows you to classify assets and manage multiple assets at a time.
//
// @param request - CreateHostGroupRequest
//
// @return CreateHostGroupResponse
func (client *Client) CreateHostGroup(request *CreateHostGroupRequest) (_result *CreateHostGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CreateHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Bastionhost provides the shared key feature. This feature allows you to manage the private key that is used to log on to a host in a bastion host. This way, you can associate the private key with multiple accounts of the host to make host account management more efficient.
//
// @param request - CreateHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostShareKeyResponse
func (client *Client) CreateHostShareKeyWithOptions(request *CreateHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateHostShareKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyName) {
		query["HostShareKeyName"] = request.HostShareKeyName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PassPhrase) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Bastionhost provides the shared key feature. This feature allows you to manage the private key that is used to log on to a host in a bastion host. This way, you can associate the private key with multiple accounts of the host to make host account management more efficient.
//
// @param request - CreateHostShareKeyRequest
//
// @return CreateHostShareKeyResponse
func (client *Client) CreateHostShareKey(request *CreateHostShareKeyRequest) (_result *CreateHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostShareKeyResponse{}
	_body, _err := client.CreateHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a network domain.
//
// @param request - CreateNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkDomainResponse
func (client *Client) CreateNetworkDomainWithOptions(request *CreateNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainName) {
		query["NetworkDomainName"] = request.NetworkDomainName
	}

	if !dara.IsNil(request.NetworkDomainType) {
		query["NetworkDomainType"] = request.NetworkDomainType
	}

	if !dara.IsNil(request.Proxies) {
		query["Proxies"] = request.Proxies
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network domain.
//
// @param request - CreateNetworkDomainRequest
//
// @return CreateNetworkDomainResponse
func (client *Client) CreateNetworkDomain(request *CreateNetworkDomainRequest) (_result *CreateNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkDomainResponse{}
	_body, _err := client.CreateNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateOperationTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOperationTicketResponse
func (client *Client) CreateOperationTicketWithOptions(request *CreateOperationTicketRequest, runtime *dara.RuntimeOptions) (_result *CreateOperationTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApproveComment) {
		query["ApproveComment"] = request.ApproveComment
	}

	if !dara.IsNil(request.AssetAccountName) {
		query["AssetAccountName"] = request.AssetAccountName
	}

	if !dara.IsNil(request.AssetId) {
		query["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.EffectEndTime) {
		query["EffectEndTime"] = request.EffectEndTime
	}

	if !dara.IsNil(request.EffectStartTime) {
		query["EffectStartTime"] = request.EffectStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsOneTimeEffect) {
		query["IsOneTimeEffect"] = request.IsOneTimeEffect
	}

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOperationTicket"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOperationTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateOperationTicketRequest
//
// @return CreateOperationTicketResponse
func (client *Client) CreateOperationTicket(request *CreateOperationTicketRequest) (_result *CreateOperationTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOperationTicketResponse{}
	_body, _err := client.CreateOperationTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a command control, command approval, protocol control, or access control policy to manage O\\&M operations. This effectively prevents users from performing high-risk operations or accidental operations to ensure O\\&M security.
//
// @param request - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2019-12-09"),
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
// Configures a command control, command approval, protocol control, or access control policy to manage O\\&M operations. This effectively prevents users from performing high-risk operations or accidental operations to ensure O\\&M security.
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
// You can create authorization rules to authorize multiple users to manage assets. You can also specify a validity period for an authorization rule. This way, you can manage users and assets in a more efficient manner and limit the time periods during which users can access assets.
//
// @param request - CreateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can create authorization rules to authorize multiple users to manage assets. You can also specify a validity period for an authorization rule. This way, you can manage users and assets in a more efficient manner and limit the time periods during which users can access assets.
//
// @param request - CreateRuleRequest
//
// @return CreateRuleResponse
func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a user to a bastion host.
//
// Description:
//
// You can call the CreateUser operation to add local users, Resource Access Management (RAM) users, Active Directory (AD)-authenticated users, or Lightweight Directory Access Protocol (LDAP)-authenticated users to a bastion host. After a Bastionhost administrator adds a user to a bastion host, O\\&M engineers can log on to the bastion host as the user to perform O\\&M operations on the hosts that the user is authorized to manage.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
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
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.LanguageStatus) {
		query["LanguageStatus"] = request.LanguageStatus
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.MobileCountryCode) {
		query["MobileCountryCode"] = request.MobileCountryCode
	}

	if !dara.IsNil(request.NeedResetPassword) {
		query["NeedResetPassword"] = request.NeedResetPassword
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceUserId) {
		query["SourceUserId"] = request.SourceUserId
	}

	if !dara.IsNil(request.TwoFactorMethods) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	if !dara.IsNil(request.TwoFactorStatus) {
		query["TwoFactorStatus"] = request.TwoFactorStatus
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2019-12-09"),
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
// Adds a user to a bastion host.
//
// Description:
//
// You can call the CreateUser operation to add local users, Resource Access Management (RAM) users, Active Directory (AD)-authenticated users, or Lightweight Directory Access Protocol (LDAP)-authenticated users to a bastion host. After a Bastionhost administrator adds a user to a bastion host, O\\&M engineers can log on to the bastion host as the user to perform O\\&M operations on the hosts that the user is authorized to manage.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
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
// Creates a user group for the specified bastion host.
//
// Description:
//
// You can call this operation to create a user group for a bastion host as an administrator. Then, you can call the [AddUsersToGroup](https://help.aliyun.com/document_detail/204600.html) operation to add users to the user group at a time. After you add the users to the user group, you can authorize and manage the users in a centralized manner.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user group for the specified bastion host.
//
// Description:
//
// You can call this operation to create a user group for a bastion host as an administrator. Then, you can call the [AddUsersToGroup](https://help.aliyun.com/document_detail/204600.html) operation to add users to the user group at a time. After you add the users to the user group, you can authorize and manage the users in a centralized manner.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateUserGroupRequest
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a public key for a bastion host user and hosts the public key in the bastion host. This way, O\\&M engineers can use the private key that corresponds to the public key to log on to the bastion host from an O\\&M client.
//
// Description:
//
// You can call the CreateUserPublicKey operation to create a public key for the specified user of a bastion host.
//
// @param request - CreateUserPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserPublicKeyResponse
func (client *Client) CreateUserPublicKeyWithOptions(request *CreateUserPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateUserPublicKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PublicKey) {
		query["PublicKey"] = request.PublicKey
	}

	if !dara.IsNil(request.PublicKeyName) {
		query["PublicKeyName"] = request.PublicKeyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserPublicKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public key for a bastion host user and hosts the public key in the bastion host. This way, O\\&M engineers can use the private key that corresponds to the public key to log on to the bastion host from an O\\&M client.
//
// Description:
//
// You can call the CreateUserPublicKey operation to create a public key for the specified user of a bastion host.
//
// @param request - CreateUserPublicKeyRequest
//
// @return CreateUserPublicKeyResponse
func (client *Client) CreateUserPublicKey(request *CreateUserPublicKeyRequest) (_result *CreateUserPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserPublicKeyResponse{}
	_body, _err := client.CreateUserPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database.
//
// @param request - DeleteDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatabase"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database.
//
// @param request - DeleteDatabaseRequest
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database account.
//
// @param request - DeleteDatabaseAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatabaseAccountResponse
func (client *Client) DeleteDatabaseAccountWithOptions(request *DeleteDatabaseAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatabaseAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountId) {
		query["DatabaseAccountId"] = request.DatabaseAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatabaseAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatabaseAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database account.
//
// @param request - DeleteDatabaseAccountRequest
//
// @return DeleteDatabaseAccountResponse
func (client *Client) DeleteDatabaseAccount(request *DeleteDatabaseAccountRequest) (_result *DeleteDatabaseAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatabaseAccountResponse{}
	_body, _err := client.DeleteDatabaseAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified host.
//
// @param request - DeleteHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostResponse
func (client *Client) DeleteHostWithOptions(request *DeleteHostRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified host.
//
// @param request - DeleteHostRequest
//
// @return DeleteHostResponse
func (client *Client) DeleteHost(request *DeleteHostRequest) (_result *DeleteHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostResponse{}
	_body, _err := client.DeleteHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a host account.
//
// Description:
//
// ## Usage notes
//
// This interface is used to delete individual host accounts. If a host account is no longer in use, you can invoke this interface to delete the host account for that host that has been configured on the bastion.
//
// >  After you remove the host account, you must enter the username and password of the host when you log on to the host in Bastionhost.
//
// ## QPS Limit
//
// The single-user QPS limit of this interface is 10 times/second. If the limit is exceeded, the API call will be stream-limited, which may affect your business, please call reasonably.
//
// @param request - DeleteHostAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostAccountResponse
func (client *Client) DeleteHostAccountWithOptions(request *DeleteHostAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountId) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHostAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a host account.
//
// Description:
//
// ## Usage notes
//
// This interface is used to delete individual host accounts. If a host account is no longer in use, you can invoke this interface to delete the host account for that host that has been configured on the bastion.
//
// >  After you remove the host account, you must enter the username and password of the host when you log on to the host in Bastionhost.
//
// ## QPS Limit
//
// The single-user QPS limit of this interface is 10 times/second. If the limit is exceeded, the API call will be stream-limited, which may affect your business, please call reasonably.
//
// @param request - DeleteHostAccountRequest
//
// @return DeleteHostAccountResponse
func (client *Client) DeleteHostAccount(request *DeleteHostAccountRequest) (_result *DeleteHostAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostAccountResponse{}
	_body, _err := client.DeleteHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a host group.
//
// Description:
//
// You can call this operation to delete a single host group. If you no longer need to perform O\\&M operations on all hosts in a host group, you can call this operation to delete the host group.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteHostGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostGroupResponse
func (client *Client) DeleteHostGroupWithOptions(request *DeleteHostGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHostGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a host group.
//
// Description:
//
// You can call this operation to delete a single host group. If you no longer need to perform O\\&M operations on all hosts in a host group, you can call this operation to delete the host group.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteHostGroupRequest
//
// @return DeleteHostGroupResponse
func (client *Client) DeleteHostGroup(request *DeleteHostGroupRequest) (_result *DeleteHostGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.DeleteHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a shared key.
//
// @param request - DeleteHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostShareKeyResponse
func (client *Client) DeleteHostShareKeyWithOptions(request *DeleteHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostShareKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a shared key.
//
// @param request - DeleteHostShareKeyRequest
//
// @return DeleteHostShareKeyResponse
func (client *Client) DeleteHostShareKey(request *DeleteHostShareKeyRequest) (_result *DeleteHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostShareKeyResponse{}
	_body, _err := client.DeleteHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a network domain.
//
// @param request - DeleteNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkDomainResponse
func (client *Client) DeleteNetworkDomainWithOptions(request *DeleteNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a network domain.
//
// @param request - DeleteNetworkDomainRequest
//
// @return DeleteNetworkDomainResponse
func (client *Client) DeleteNetworkDomain(request *DeleteNetworkDomainRequest) (_result *DeleteNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkDomainResponse{}
	_body, _err := client.DeleteNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a control policy.
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2019-12-09"),
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
// Deletes a control policy.
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
// Deletes an authorization rule.
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an authorization rule.
//
// @param request - DeleteRuleRequest
//
// @return DeleteRuleResponse
func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a bastion host user.
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2019-12-09"),
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
// Deletes a bastion host user.
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
// Deletes a specified user group from a specified bastion host.
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified user group from a specified bastion host.
//
// @param request - DeleteUserGroupRequest
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a public key from the specified user.
//
// Description:
//
// You can call the DeleteUserPublicKey operation to delete a public key from the specified user of a bastion host.
//
// @param request - DeleteUserPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPublicKeyResponse
func (client *Client) DeleteUserPublicKeyWithOptions(request *DeleteUserPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserPublicKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PublicKeyId) {
		query["PublicKeyId"] = request.PublicKeyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserPublicKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a public key from the specified user.
//
// Description:
//
// You can call the DeleteUserPublicKey operation to delete a public key from the specified user of a bastion host.
//
// @param request - DeleteUserPublicKeyRequest
//
// @return DeleteUserPublicKeyResponse
func (client *Client) DeleteUserPublicKey(request *DeleteUserPublicKeyRequest) (_result *DeleteUserPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserPublicKeyResponse{}
	_body, _err := client.DeleteUserPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attribute information about the specified bastion host. The information includes the ID and remarks of the bastion host.
//
// Description:
//
// ***
//
// @param request - DescribeInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAttributeResponse
func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAttribute"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attribute information about the specified bastion host. The information includes the ID and remarks of the bastion host.
//
// Description:
//
// ***
//
// @param request - DescribeInstanceAttributeRequest
//
// @return DescribeInstanceAttributeResponse
func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (_result *DescribeInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DescribeInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 堡垒机实例列表
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 堡垒机实例列表
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available regions where you can create bastion hosts.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available regions where you can create bastion hosts.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on databases and database accounts from a user.
//
// @param request - DetachDatabaseAccountsFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDatabaseAccountsFromUserResponse
func (client *Client) DetachDatabaseAccountsFromUserWithOptions(request *DetachDatabaseAccountsFromUserRequest, runtime *dara.RuntimeOptions) (_result *DetachDatabaseAccountsFromUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDatabaseAccountsFromUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDatabaseAccountsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on databases and database accounts from a user.
//
// @param request - DetachDatabaseAccountsFromUserRequest
//
// @return DetachDatabaseAccountsFromUserResponse
func (client *Client) DetachDatabaseAccountsFromUser(request *DetachDatabaseAccountsFromUserRequest) (_result *DetachDatabaseAccountsFromUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachDatabaseAccountsFromUserResponse{}
	_body, _err := client.DetachDatabaseAccountsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on databases and database accounts from a user group.
//
// @param request - DetachDatabaseAccountsFromUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDatabaseAccountsFromUserGroupResponse
func (client *Client) DetachDatabaseAccountsFromUserGroupWithOptions(request *DetachDatabaseAccountsFromUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DetachDatabaseAccountsFromUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDatabaseAccountsFromUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDatabaseAccountsFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on databases and database accounts from a user group.
//
// @param request - DetachDatabaseAccountsFromUserGroupRequest
//
// @return DetachDatabaseAccountsFromUserGroupResponse
func (client *Client) DetachDatabaseAccountsFromUserGroup(request *DetachDatabaseAccountsFromUserGroupRequest) (_result *DetachDatabaseAccountsFromUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachDatabaseAccountsFromUserGroupResponse{}
	_body, _err := client.DetachDatabaseAccountsFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociate host accounts from a shared key.
//
// @param request - DetachHostAccountsFromHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostAccountsFromHostShareKeyResponse
func (client *Client) DetachHostAccountsFromHostShareKeyWithOptions(request *DetachHostAccountsFromHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *DetachHostAccountsFromHostShareKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountIds) {
		query["HostAccountIds"] = request.HostAccountIds
	}

	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostAccountsFromHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostAccountsFromHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociate host accounts from a shared key.
//
// @param request - DetachHostAccountsFromHostShareKeyRequest
//
// @return DetachHostAccountsFromHostShareKeyResponse
func (client *Client) DetachHostAccountsFromHostShareKey(request *DetachHostAccountsFromHostShareKeyRequest) (_result *DetachHostAccountsFromHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostAccountsFromHostShareKeyResponse{}
	_body, _err := client.DetachHostAccountsFromHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on hosts and host accounts from a user.
//
// @param request - DetachHostAccountsFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostAccountsFromUserResponse
func (client *Client) DetachHostAccountsFromUserWithOptions(request *DetachHostAccountsFromUserRequest, runtime *dara.RuntimeOptions) (_result *DetachHostAccountsFromUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostAccountsFromUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostAccountsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on hosts and host accounts from a user.
//
// @param request - DetachHostAccountsFromUserRequest
//
// @return DetachHostAccountsFromUserResponse
func (client *Client) DetachHostAccountsFromUser(request *DetachHostAccountsFromUserRequest) (_result *DetachHostAccountsFromUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostAccountsFromUserResponse{}
	_body, _err := client.DetachHostAccountsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions on one or more hosts and host accounts from a user group.
//
// @param request - DetachHostAccountsFromUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostAccountsFromUserGroupResponse
func (client *Client) DetachHostAccountsFromUserGroupWithOptions(request *DetachHostAccountsFromUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DetachHostAccountsFromUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostAccountsFromUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostAccountsFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions on one or more hosts and host accounts from a user group.
//
// @param request - DetachHostAccountsFromUserGroupRequest
//
// @return DetachHostAccountsFromUserGroupResponse
func (client *Client) DetachHostAccountsFromUserGroup(request *DetachHostAccountsFromUserGroupRequest) (_result *DetachHostAccountsFromUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostAccountsFromUserGroupResponse{}
	_body, _err := client.DetachHostAccountsFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes host groups and host accounts from the list of host groups and host accounts that a user is authorized to manage.
//
// @param request - DetachHostGroupAccountsFromUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostGroupAccountsFromUserResponse
func (client *Client) DetachHostGroupAccountsFromUserWithOptions(request *DetachHostGroupAccountsFromUserRequest, runtime *dara.RuntimeOptions) (_result *DetachHostGroupAccountsFromUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostGroupAccountsFromUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostGroupAccountsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes host groups and host accounts from the list of host groups and host accounts that a user is authorized to manage.
//
// @param request - DetachHostGroupAccountsFromUserRequest
//
// @return DetachHostGroupAccountsFromUserResponse
func (client *Client) DetachHostGroupAccountsFromUser(request *DetachHostGroupAccountsFromUserRequest) (_result *DetachHostGroupAccountsFromUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostGroupAccountsFromUserResponse{}
	_body, _err := client.DetachHostGroupAccountsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on one or more host groups and host accounts from a user group.
//
// Description:
//
// ***
//
// @param request - DetachHostGroupAccountsFromUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachHostGroupAccountsFromUserGroupResponse
func (client *Client) DetachHostGroupAccountsFromUserGroupWithOptions(request *DetachHostGroupAccountsFromUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DetachHostGroupAccountsFromUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachHostGroupAccountsFromUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachHostGroupAccountsFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on one or more host groups and host accounts from a user group.
//
// Description:
//
// ***
//
// @param request - DetachHostGroupAccountsFromUserGroupRequest
//
// @return DetachHostGroupAccountsFromUserGroupResponse
func (client *Client) DetachHostGroupAccountsFromUserGroup(request *DetachHostGroupAccountsFromUserGroupRequest) (_result *DetachHostGroupAccountsFromUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachHostGroupAccountsFromUserGroupResponse{}
	_body, _err := client.DetachHostGroupAccountsFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables Internet access for a bastion host.
//
// @param request - DisableInstancePublicAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInstancePublicAccessResponse
func (client *Client) DisableInstancePublicAccessWithOptions(request *DisableInstancePublicAccessRequest, runtime *dara.RuntimeOptions) (_result *DisableInstancePublicAccessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableInstancePublicAccess"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables Internet access for a bastion host.
//
// @param request - DisableInstancePublicAccessRequest
//
// @return DisableInstancePublicAccessResponse
func (client *Client) DisableInstancePublicAccess(request *DisableInstancePublicAccessRequest) (_result *DisableInstancePublicAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.DisableInstancePublicAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an authorization rule.
//
// @param request - DisableRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableRuleResponse
func (client *Client) DisableRuleWithOptions(request *DisableRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DisableRuleRequest
//
// @return DisableRuleResponse
func (client *Client) DisableRule(request *DisableRuleRequest) (_result *DisableRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableRuleResponse{}
	_body, _err := client.DisableRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables Internet access for a bastion host.
//
// @param request - EnableInstancePublicAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInstancePublicAccessResponse
func (client *Client) EnableInstancePublicAccessWithOptions(request *EnableInstancePublicAccessRequest, runtime *dara.RuntimeOptions) (_result *EnableInstancePublicAccessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInstancePublicAccess"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables Internet access for a bastion host.
//
// @param request - EnableInstancePublicAccessRequest
//
// @return EnableInstancePublicAccessResponse
func (client *Client) EnableInstancePublicAccess(request *EnableInstancePublicAccessRequest) (_result *EnableInstancePublicAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.EnableInstancePublicAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an authorization rule.
//
// @param request - EnableRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRuleResponse
func (client *Client) EnableRuleWithOptions(request *EnableRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - EnableRuleRequest
//
// @return EnableRuleResponse
func (client *Client) EnableRule(request *EnableRuleRequest) (_result *EnableRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableRuleResponse{}
	_body, _err := client.EnableRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for an O\\&M token.
//
// @param request - GenerateAssetOperationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAssetOperationTokenResponse
func (client *Client) GenerateAssetOperationTokenWithOptions(request *GenerateAssetOperationTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateAssetOperationTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetAccountId) {
		query["AssetAccountId"] = request.AssetAccountId
	}

	if !dara.IsNil(request.AssetAccountName) {
		query["AssetAccountName"] = request.AssetAccountName
	}

	if !dara.IsNil(request.AssetAccountPassword) {
		query["AssetAccountPassword"] = request.AssetAccountPassword
	}

	if !dara.IsNil(request.AssetAccountProtocolName) {
		query["AssetAccountProtocolName"] = request.AssetAccountProtocolName
	}

	if !dara.IsNil(request.AssetId) {
		query["AssetId"] = request.AssetId
	}

	if !dara.IsNil(request.AssetType) {
		query["AssetType"] = request.AssetType
	}

	if !dara.IsNil(request.DatabaseSchema) {
		query["DatabaseSchema"] = request.DatabaseSchema
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LoginAttribute) {
		query["LoginAttribute"] = request.LoginAttribute
	}

	if !dara.IsNil(request.OperationMode) {
		query["OperationMode"] = request.OperationMode
	}

	if !dara.IsNil(request.OperationNote) {
		query["OperationNote"] = request.OperationNote
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SsoClient) {
		query["SsoClient"] = request.SsoClient
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAssetOperationToken"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAssetOperationTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for an O\\&M token.
//
// @param request - GenerateAssetOperationTokenRequest
//
// @return GenerateAssetOperationTokenResponse
func (client *Client) GenerateAssetOperationToken(request *GenerateAssetOperationTokenRequest) (_result *GenerateAssetOperationTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateAssetOperationTokenResponse{}
	_body, _err := client.GenerateAssetOperationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a database.
//
// @param request - GetDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseResponse
func (client *Client) GetDatabaseWithOptions(request *GetDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabase"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a database.
//
// @param request - GetDatabaseRequest
//
// @return GetDatabaseResponse
func (client *Client) GetDatabase(request *GetDatabaseRequest) (_result *GetDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatabaseResponse{}
	_body, _err := client.GetDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a database account.
//
// @param request - GetDatabaseAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseAccountResponse
func (client *Client) GetDatabaseAccountWithOptions(request *GetDatabaseAccountRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountId) {
		query["DatabaseAccountId"] = request.DatabaseAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabaseAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a database account.
//
// @param request - GetDatabaseAccountRequest
//
// @return GetDatabaseAccountResponse
func (client *Client) GetDatabaseAccount(request *GetDatabaseAccountRequest) (_result *GetDatabaseAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatabaseAccountResponse{}
	_body, _err := client.GetDatabaseAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetExportConfigJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExportConfigJobResponse
func (client *Client) GetExportConfigJobWithOptions(request *GetExportConfigJobRequest, runtime *dara.RuntimeOptions) (_result *GetExportConfigJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExportConfigJob"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExportConfigJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetExportConfigJobRequest
//
// @return GetExportConfigJobResponse
func (client *Client) GetExportConfigJob(request *GetExportConfigJobRequest) (_result *GetExportConfigJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExportConfigJobResponse{}
	_body, _err := client.GetExportConfigJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a host, such as the name, source, address, protocol, and service port of the host.
//
// @param request - GetHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostResponse
func (client *Client) GetHostWithOptions(request *GetHostRequest, runtime *dara.RuntimeOptions) (_result *GetHostResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a host, such as the name, source, address, protocol, and service port of the host.
//
// @param request - GetHostRequest
//
// @return GetHostResponse
func (client *Client) GetHost(request *GetHostRequest) (_result *GetHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHostResponse{}
	_body, _err := client.GetHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified host account.
//
// @param request - GetHostAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostAccountResponse
func (client *Client) GetHostAccountWithOptions(request *GetHostAccountRequest, runtime *dara.RuntimeOptions) (_result *GetHostAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountId) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHostAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified host account.
//
// @param request - GetHostAccountRequest
//
// @return GetHostAccountResponse
func (client *Client) GetHostAccount(request *GetHostAccountRequest) (_result *GetHostAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHostAccountResponse{}
	_body, _err := client.GetHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified host group.
//
// @param request - GetHostGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostGroupResponse
func (client *Client) GetHostGroupWithOptions(request *GetHostGroupRequest, runtime *dara.RuntimeOptions) (_result *GetHostGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHostGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified host group.
//
// @param request - GetHostGroupRequest
//
// @return GetHostGroupResponse
func (client *Client) GetHostGroup(request *GetHostGroupRequest) (_result *GetHostGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHostGroupResponse{}
	_body, _err := client.GetHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a shared key.
//
// @param request - GetHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostShareKeyResponse
func (client *Client) GetHostShareKeyWithOptions(request *GetHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *GetHostShareKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a shared key.
//
// @param request - GetHostShareKeyRequest
//
// @return GetHostShareKeyResponse
func (client *Client) GetHostShareKey(request *GetHostShareKeyRequest) (_result *GetHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHostShareKeyResponse{}
	_body, _err := client.GetHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the settings of Active Directory (AD) authentication on a bastion host.
//
// Description:
//
// ###
//
// You can call this operation to query the settings of AD authentication on a bastion host. After you configure AD authentication on a bastion host, you can import AD-authenticated users into the bastion host. After the AD-authenticated users are imported into the bastion host, the AD-authenticated users can log on to the bastion host to perform O\\&M operations on servers.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetInstanceADAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceADAuthServerResponse
func (client *Client) GetInstanceADAuthServerWithOptions(request *GetInstanceADAuthServerRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceADAuthServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceADAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of Active Directory (AD) authentication on a bastion host.
//
// Description:
//
// ###
//
// You can call this operation to query the settings of AD authentication on a bastion host. After you configure AD authentication on a bastion host, you can import AD-authenticated users into the bastion host. After the AD-authenticated users are imported into the bastion host, the AD-authenticated users can log on to the bastion host to perform O\\&M operations on servers.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetInstanceADAuthServerRequest
//
// @return GetInstanceADAuthServerResponse
func (client *Client) GetInstanceADAuthServer(request *GetInstanceADAuthServerRequest) (_result *GetInstanceADAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceADAuthServerResponse{}
	_body, _err := client.GetInstanceADAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the settings of Lightweight Directory Access Protocol (LDAP) authentication on a bastion host.
//
// @param request - GetInstanceLDAPAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceLDAPAuthServerResponse
func (client *Client) GetInstanceLDAPAuthServerWithOptions(request *GetInstanceLDAPAuthServerRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceLDAPAuthServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceLDAPAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of Lightweight Directory Access Protocol (LDAP) authentication on a bastion host.
//
// @param request - GetInstanceLDAPAuthServerRequest
//
// @return GetInstanceLDAPAuthServerResponse
func (client *Client) GetInstanceLDAPAuthServer(request *GetInstanceLDAPAuthServerRequest) (_result *GetInstanceLDAPAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceLDAPAuthServerResponse{}
	_body, _err := client.GetInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetInstanceStoreInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceStoreInfoResponse
func (client *Client) GetInstanceStoreInfoWithOptions(request *GetInstanceStoreInfoRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceStoreInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceStoreInfo"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceStoreInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetInstanceStoreInfoRequest
//
// @return GetInstanceStoreInfoResponse
func (client *Client) GetInstanceStoreInfo(request *GetInstanceStoreInfoRequest) (_result *GetInstanceStoreInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceStoreInfoResponse{}
	_body, _err := client.GetInstanceStoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the settings of two-factor authentication on a bastion host.
//
// Description:
//
// You can call this operation to query the settings of two-factor authentication on a bastion host. After you enable two-factor authentication, Bastionhost sends a verification code to a local user when the local user logs on to a bastion host. A local user can log on to the bastion host only when the local user enters the valid username and password and the verification code. This reduces the security risks caused by account information leaks.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetInstanceTwoFactorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceTwoFactorResponse
func (client *Client) GetInstanceTwoFactorWithOptions(request *GetInstanceTwoFactorRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceTwoFactorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceTwoFactor"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceTwoFactorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of two-factor authentication on a bastion host.
//
// Description:
//
// You can call this operation to query the settings of two-factor authentication on a bastion host. After you enable two-factor authentication, Bastionhost sends a verification code to a local user when the local user logs on to a bastion host. A local user can log on to the bastion host only when the local user enters the valid username and password and the verification code. This reduces the security risks caused by account information leaks.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetInstanceTwoFactorRequest
//
// @return GetInstanceTwoFactorResponse
func (client *Client) GetInstanceTwoFactor(request *GetInstanceTwoFactorRequest) (_result *GetInstanceTwoFactorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceTwoFactorResponse{}
	_body, _err := client.GetInstanceTwoFactorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a network domain.
//
// @param request - GetNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkDomainResponse
func (client *Client) GetNetworkDomainWithOptions(request *GetNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a network domain.
//
// @param request - GetNetworkDomainRequest
//
// @return GetNetworkDomainResponse
func (client *Client) GetNetworkDomain(request *GetNetworkDomainRequest) (_result *GetNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetworkDomainResponse{}
	_body, _err := client.GetNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a control policy.
//
// @param request - GetPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicy"),
		Version:     dara.String("2019-12-09"),
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
// Queries the detailed information about a control policy.
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
// Queries the assets to which a control policy applies.
//
// @param request - GetPolicyAssetScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyAssetScopeResponse
func (client *Client) GetPolicyAssetScopeWithOptions(request *GetPolicyAssetScopeRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyAssetScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicyAssetScope"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyAssetScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the assets to which a control policy applies.
//
// @param request - GetPolicyAssetScopeRequest
//
// @return GetPolicyAssetScopeResponse
func (client *Client) GetPolicyAssetScope(request *GetPolicyAssetScopeRequest) (_result *GetPolicyAssetScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicyAssetScopeResponse{}
	_body, _err := client.GetPolicyAssetScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scope of users to whom a control policy applies.
//
// @param request - GetPolicyUserScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyUserScopeResponse
func (client *Client) GetPolicyUserScopeWithOptions(request *GetPolicyUserScopeRequest, runtime *dara.RuntimeOptions) (_result *GetPolicyUserScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicyUserScope"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyUserScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scope of users to whom a control policy applies.
//
// @param request - GetPolicyUserScopeRequest
//
// @return GetPolicyUserScopeResponse
func (client *Client) GetPolicyUserScope(request *GetPolicyUserScopeRequest) (_result *GetPolicyUserScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPolicyUserScopeResponse{}
	_body, _err := client.GetPolicyUserScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about an authorization rule.
//
// @param request - GetRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleResponse
func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *dara.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about an authorization rule.
//
// @param request - GetRuleRequest
//
// @return GetRuleResponse
func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a user of the specified bastion host.
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2019-12-09"),
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
// Queries the details of a user of the specified bastion host.
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
// Queries the details of a user group in a bastion host.
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a user group in a bastion host.
//
// @param request - GetUserGroupRequest
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroup(request *GetUserGroupRequest) (_result *GetUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserGroupResponse{}
	_body, _err := client.GetUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries commands to be reviewed.
//
// Description:
//
// You can call this operation to query commands to be reviewed by a Bastionhost administrator.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListApproveCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApproveCommandsResponse
func (client *Client) ListApproveCommandsWithOptions(request *ListApproveCommandsRequest, runtime *dara.RuntimeOptions) (_result *ListApproveCommandsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApproveCommands"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApproveCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries commands to be reviewed.
//
// Description:
//
// You can call this operation to query commands to be reviewed by a Bastionhost administrator.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListApproveCommandsRequest
//
// @return ListApproveCommandsResponse
func (client *Client) ListApproveCommands(request *ListApproveCommandsRequest) (_result *ListApproveCommandsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApproveCommandsResponse{}
	_body, _err := client.ListApproveCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the database accounts of a database.
//
// @param request - ListDatabaseAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseAccountsResponse
func (client *Client) ListDatabaseAccountsWithOptions(request *ListDatabaseAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListDatabaseAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabaseAccounts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabaseAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the database accounts of a database.
//
// @param request - ListDatabaseAccountsRequest
//
// @return ListDatabaseAccountsResponse
func (client *Client) ListDatabaseAccounts(request *ListDatabaseAccountsRequest) (_result *ListDatabaseAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabaseAccountsResponse{}
	_body, _err := client.ListDatabaseAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the database accounts of a database and whether a user is authorized to manage each database account.
//
// @param request - ListDatabaseAccountsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseAccountsForUserResponse
func (client *Client) ListDatabaseAccountsForUserWithOptions(request *ListDatabaseAccountsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListDatabaseAccountsForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabaseAccountsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabaseAccountsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the database accounts of a database and whether a user is authorized to manage each database account.
//
// @param request - ListDatabaseAccountsForUserRequest
//
// @return ListDatabaseAccountsForUserResponse
func (client *Client) ListDatabaseAccountsForUser(request *ListDatabaseAccountsForUserRequest) (_result *ListDatabaseAccountsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabaseAccountsForUserResponse{}
	_body, _err := client.ListDatabaseAccountsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the database accounts of a database and whether a user group is authorized to manage each database account.
//
// @param request - ListDatabaseAccountsForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseAccountsForUserGroupResponse
func (client *Client) ListDatabaseAccountsForUserGroupWithOptions(request *ListDatabaseAccountsForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListDatabaseAccountsForUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabaseAccountsForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabaseAccountsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the database accounts of a database and whether a user group is authorized to manage each database account.
//
// @param request - ListDatabaseAccountsForUserGroupRequest
//
// @return ListDatabaseAccountsForUserGroupResponse
func (client *Client) ListDatabaseAccountsForUserGroup(request *ListDatabaseAccountsForUserGroupRequest) (_result *ListDatabaseAccountsForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabaseAccountsForUserGroupResponse{}
	_body, _err := client.ListDatabaseAccountsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases that are managed by a bastion host.
//
// @param request - ListDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithOptions(request *ListDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases that are managed by a bastion host.
//
// @param request - ListDatabasesRequest
//
// @return ListDatabasesResponse
func (client *Client) ListDatabases(request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases that a user is authorized to manage.
//
// @param request - ListDatabasesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesForUserResponse
func (client *Client) ListDatabasesForUserWithOptions(request *ListDatabasesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAddress) {
		query["DatabaseAddress"] = request.DatabaseAddress
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabasesForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases that a user is authorized to manage.
//
// @param request - ListDatabasesForUserRequest
//
// @return ListDatabasesForUserResponse
func (client *Client) ListDatabasesForUser(request *ListDatabasesForUserRequest) (_result *ListDatabasesForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesForUserResponse{}
	_body, _err := client.ListDatabasesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases that a user group is authorized to manage.
//
// @param request - ListDatabasesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesForUserGroupResponse
func (client *Client) ListDatabasesForUserGroupWithOptions(request *ListDatabasesForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesForUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAddress) {
		query["DatabaseAddress"] = request.DatabaseAddress
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabasesForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases that a user group is authorized to manage.
//
// @param request - ListDatabasesForUserGroupRequest
//
// @return ListDatabasesForUserGroupResponse
func (client *Client) ListDatabasesForUserGroup(request *ListDatabasesForUserGroupRequest) (_result *ListDatabasesForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesForUserGroupResponse{}
	_body, _err := client.ListDatabasesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries accounts of a specified host.
//
// @param request - ListHostAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostAccountsResponse
func (client *Client) ListHostAccountsWithOptions(request *ListHostAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListHostAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
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

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostAccounts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries accounts of a specified host.
//
// @param request - ListHostAccountsRequest
//
// @return ListHostAccountsResponse
func (client *Client) ListHostAccounts(request *ListHostAccountsRequest) (_result *ListHostAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostAccountsResponse{}
	_body, _err := client.ListHostAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the host accounts that are associated with a shared key.
//
// @param request - ListHostAccountsForHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostAccountsForHostShareKeyResponse
func (client *Client) ListHostAccountsForHostShareKeyWithOptions(request *ListHostAccountsForHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *ListHostAccountsForHostShareKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostAccountsForHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostAccountsForHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the host accounts that are associated with a shared key.
//
// @param request - ListHostAccountsForHostShareKeyRequest
//
// @return ListHostAccountsForHostShareKeyResponse
func (client *Client) ListHostAccountsForHostShareKey(request *ListHostAccountsForHostShareKeyRequest) (_result *ListHostAccountsForHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostAccountsForHostShareKeyResponse{}
	_body, _err := client.ListHostAccountsForHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the host accounts that the specified user is authorized to manage on the specified host.
//
// @param request - ListHostAccountsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostAccountsForUserResponse
func (client *Client) ListHostAccountsForUserWithOptions(request *ListHostAccountsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListHostAccountsForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostAccountsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostAccountsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the host accounts that the specified user is authorized to manage on the specified host.
//
// @param request - ListHostAccountsForUserRequest
//
// @return ListHostAccountsForUserResponse
func (client *Client) ListHostAccountsForUser(request *ListHostAccountsForUserRequest) (_result *ListHostAccountsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostAccountsForUserResponse{}
	_body, _err := client.ListHostAccountsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the host accounts of the specified host that the specified user group is authorized to manage.
//
// @param request - ListHostAccountsForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostAccountsForUserGroupResponse
func (client *Client) ListHostAccountsForUserGroupWithOptions(request *ListHostAccountsForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListHostAccountsForUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostAccountsForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostAccountsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the host accounts of the specified host that the specified user group is authorized to manage.
//
// @param request - ListHostAccountsForUserGroupRequest
//
// @return ListHostAccountsForUserGroupResponse
func (client *Client) ListHostAccountsForUserGroup(request *ListHostAccountsForUserGroupRequest) (_result *ListHostAccountsForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostAccountsForUserGroupResponse{}
	_body, _err := client.ListHostAccountsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the names of the host accounts that a specified user is authorized to manage in a specified host group.
//
// @param request - ListHostGroupAccountNamesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupAccountNamesForUserResponse
func (client *Client) ListHostGroupAccountNamesForUserWithOptions(request *ListHostGroupAccountNamesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupAccountNamesForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroupAccountNamesForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupAccountNamesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of the host accounts that a specified user is authorized to manage in a specified host group.
//
// @param request - ListHostGroupAccountNamesForUserRequest
//
// @return ListHostGroupAccountNamesForUserResponse
func (client *Client) ListHostGroupAccountNamesForUser(request *ListHostGroupAccountNamesForUserRequest) (_result *ListHostGroupAccountNamesForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupAccountNamesForUserResponse{}
	_body, _err := client.ListHostGroupAccountNamesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the names of the host accounts that a user group is authorized to manage in a host group.
//
// @param request - ListHostGroupAccountNamesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupAccountNamesForUserGroupResponse
func (client *Client) ListHostGroupAccountNamesForUserGroupWithOptions(request *ListHostGroupAccountNamesForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupAccountNamesForUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroupAccountNamesForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupAccountNamesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of the host accounts that a user group is authorized to manage in a host group.
//
// @param request - ListHostGroupAccountNamesForUserGroupRequest
//
// @return ListHostGroupAccountNamesForUserGroupResponse
func (client *Client) ListHostGroupAccountNamesForUserGroup(request *ListHostGroupAccountNamesForUserGroupRequest) (_result *ListHostGroupAccountNamesForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupAccountNamesForUserGroupResponse{}
	_body, _err := client.ListHostGroupAccountNamesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of asset groups that are managed by a bastion host.
//
// @param request - ListHostGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupsResponse
func (client *Client) ListHostGroupsWithOptions(request *ListHostGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroups"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of asset groups that are managed by a bastion host.
//
// @param request - ListHostGroupsRequest
//
// @return ListHostGroupsResponse
func (client *Client) ListHostGroups(request *ListHostGroupsRequest) (_result *ListHostGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.ListHostGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of host groups that a bastion host user is authorized or is not authorized to manage.
//
// @param request - ListHostGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupsForUserResponse
func (client *Client) ListHostGroupsForUserWithOptions(request *ListHostGroupsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupsForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroupsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of host groups that a bastion host user is authorized or is not authorized to manage.
//
// @param request - ListHostGroupsForUserRequest
//
// @return ListHostGroupsForUserResponse
func (client *Client) ListHostGroupsForUser(request *ListHostGroupsForUserRequest) (_result *ListHostGroupsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupsForUserResponse{}
	_body, _err := client.ListHostGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hosts that a specified user group is authorized or not authorized to manage.
//
// @param request - ListHostGroupsForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupsForUserGroupResponse
func (client *Client) ListHostGroupsForUserGroupWithOptions(request *ListHostGroupsForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListHostGroupsForUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroupsForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hosts that a specified user group is authorized or not authorized to manage.
//
// @param request - ListHostGroupsForUserGroupRequest
//
// @return ListHostGroupsForUserGroupResponse
func (client *Client) ListHostGroupsForUserGroup(request *ListHostGroupsForUserGroupRequest) (_result *ListHostGroupsForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostGroupsForUserGroupResponse{}
	_body, _err := client.ListHostGroupsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the shared keys that are associated with a host.
//
// @param request - ListHostShareKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostShareKeysResponse
func (client *Client) ListHostShareKeysWithOptions(request *ListHostShareKeysRequest, runtime *dara.RuntimeOptions) (_result *ListHostShareKeysResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostShareKeys"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostShareKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the shared keys that are associated with a host.
//
// @param request - ListHostShareKeysRequest
//
// @return ListHostShareKeysResponse
func (client *Client) ListHostShareKeys(request *ListHostShareKeysRequest) (_result *ListHostShareKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostShareKeysResponse{}
	_body, _err := client.ListHostShareKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hosts in a bastion host.
//
// @param request - ListHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostsResponse
func (client *Client) ListHostsWithOptions(request *ListHostsRequest, runtime *dara.RuntimeOptions) (_result *ListHostsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourceInstanceState) {
		query["SourceInstanceState"] = request.SourceInstanceState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHosts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hosts in a bastion host.
//
// @param request - ListHostsRequest
//
// @return ListHostsResponse
func (client *Client) ListHosts(request *ListHostsRequest) (_result *ListHostsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostsResponse{}
	_body, _err := client.ListHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hosts that a user group is authorized or not authorized to manage.
//
// @param request - ListHostsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostsForUserResponse
func (client *Client) ListHostsForUserWithOptions(request *ListHostsForUserRequest, runtime *dara.RuntimeOptions) (_result *ListHostsForUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostsForUser"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hosts that a user group is authorized or not authorized to manage.
//
// @param request - ListHostsForUserRequest
//
// @return ListHostsForUserResponse
func (client *Client) ListHostsForUser(request *ListHostsForUserRequest) (_result *ListHostsForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostsForUserResponse{}
	_body, _err := client.ListHostsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hosts that a user group is authorized or not authorized to manage.
//
// @param request - ListHostsForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostsForUserGroupResponse
func (client *Client) ListHostsForUserGroupWithOptions(request *ListHostsForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListHostsForUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostsForUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hosts that a user group is authorized or not authorized to manage.
//
// @param request - ListHostsForUserGroupRequest
//
// @return ListHostsForUserGroupResponse
func (client *Client) ListHostsForUserGroup(request *ListHostsForUserGroupRequest) (_result *ListHostsForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHostsForUserGroupResponse{}
	_body, _err := client.ListHostsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取RD成员账号列表
//
// @param request - ListInstanceRdMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceRdMembersResponse
func (client *Client) ListInstanceRdMembersWithOptions(request *ListInstanceRdMembersRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceRdMembersResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceRdMembers"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceRdMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取RD成员账号列表
//
// @param request - ListInstanceRdMembersRequest
//
// @return ListInstanceRdMembersResponse
func (client *Client) ListInstanceRdMembers(request *ListInstanceRdMembersRequest) (_result *ListInstanceRdMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceRdMembersResponse{}
	_body, _err := client.ListInstanceRdMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network domains created in a bastion host.
//
// @param request - ListNetworkDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkDomainsResponse
func (client *Client) ListNetworkDomainsWithOptions(request *ListNetworkDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainName) {
		query["NetworkDomainName"] = request.NetworkDomainName
	}

	if !dara.IsNil(request.NetworkDomainType) {
		query["NetworkDomainType"] = request.NetworkDomainType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetworkDomains"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network domains created in a bastion host.
//
// @param request - ListNetworkDomainsRequest
//
// @return ListNetworkDomainsResponse
func (client *Client) ListNetworkDomains(request *ListNetworkDomainsRequest) (_result *ListNetworkDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNetworkDomainsResponse{}
	_body, _err := client.ListNetworkDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of database accounts that the current Resource Access Management (RAM) user is authorized to manage.
//
// @param request - ListOperationDatabaseAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationDatabaseAccountsResponse
func (client *Client) ListOperationDatabaseAccountsWithOptions(request *ListOperationDatabaseAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationDatabaseAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationDatabaseAccounts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationDatabaseAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of database accounts that the current Resource Access Management (RAM) user is authorized to manage.
//
// @param request - ListOperationDatabaseAccountsRequest
//
// @return ListOperationDatabaseAccountsResponse
func (client *Client) ListOperationDatabaseAccounts(request *ListOperationDatabaseAccountsRequest) (_result *ListOperationDatabaseAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationDatabaseAccountsResponse{}
	_body, _err := client.ListOperationDatabaseAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of databases that the current Resource Access Management (RAM) user is authorized to manage.
//
// @param request - ListOperationDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationDatabasesResponse
func (client *Client) ListOperationDatabasesWithOptions(request *ListOperationDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListOperationDatabasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAddress) {
		query["DatabaseAddress"] = request.DatabaseAddress
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourceInstanceState) {
		query["SourceInstanceState"] = request.SourceInstanceState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationDatabases"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of databases that the current Resource Access Management (RAM) user is authorized to manage.
//
// @param request - ListOperationDatabasesRequest
//
// @return ListOperationDatabasesResponse
func (client *Client) ListOperationDatabases(request *ListOperationDatabasesRequest) (_result *ListOperationDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationDatabasesResponse{}
	_body, _err := client.ListOperationDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of host accounts that the current Resource Access Management (RAM) user is authorized to manage.
//
// @param request - ListOperationHostAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationHostAccountsResponse
func (client *Client) ListOperationHostAccountsWithOptions(request *ListOperationHostAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationHostAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationHostAccounts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationHostAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of host accounts that the current Resource Access Management (RAM) user is authorized to manage.
//
// @param request - ListOperationHostAccountsRequest
//
// @return ListOperationHostAccountsResponse
func (client *Client) ListOperationHostAccounts(request *ListOperationHostAccountsRequest) (_result *ListOperationHostAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationHostAccountsResponse{}
	_body, _err := client.ListOperationHostAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of hosts that the current Resource Access Management (RAM) user is authorized to manage.
//
// @param request - ListOperationHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationHostsResponse
func (client *Client) ListOperationHostsWithOptions(request *ListOperationHostsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationHostsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAddress) {
		query["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourceInstanceState) {
		query["SourceInstanceState"] = request.SourceInstanceState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationHosts"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of hosts that the current Resource Access Management (RAM) user is authorized to manage.
//
// @param request - ListOperationHostsRequest
//
// @return ListOperationHostsResponse
func (client *Client) ListOperationHosts(request *ListOperationHostsRequest) (_result *ListOperationHostsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationHostsResponse{}
	_body, _err := client.ListOperationHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries O\\\\\\\\\\\\&M applications to be reviewed.
//
// Description:
//
// You can call this operation to query the O\\&M applications to be reviewed by a Bastionhost administrator.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListOperationTicketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationTicketsResponse
func (client *Client) ListOperationTicketsWithOptions(request *ListOperationTicketsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationTicketsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetAddress) {
		query["AssetAddress"] = request.AssetAddress
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationTickets"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries O\\\\\\\\\\\\&M applications to be reviewed.
//
// Description:
//
// You can call this operation to query the O\\&M applications to be reviewed by a Bastionhost administrator.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListOperationTicketsRequest
//
// @return ListOperationTicketsResponse
func (client *Client) ListOperationTickets(request *ListOperationTicketsRequest) (_result *ListOperationTicketsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationTicketsResponse{}
	_body, _err := client.ListOperationTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of control policies.
//
// @param request - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicies"),
		Version:     dara.String("2019-12-09"),
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
// Queries a list of control policies.
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
// Queries a list of authorization rules of a bastion host.
//
// @param request - ListRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRulesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleState) {
		query["RuleState"] = request.RuleState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRules"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of authorization rules of a bastion host.
//
// @param request - ListRulesRequest
//
// @return ListRulesResponse
func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to a resource.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to a resource.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags bound to one or more Bastionhost instances.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Version:     dara.String("2019-12-09"),
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
// Queries the tags bound to one or more Bastionhost instances.
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
// Queries a list of user groups on a bastion host.
//
// @param request - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithOptions(request *ListUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroups"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of user groups on a bastion host.
//
// @param request - ListUserGroupsRequest
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all public keys of the specified user.
//
// @param request - ListUserPublicKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPublicKeysResponse
func (client *Client) ListUserPublicKeysWithOptions(request *ListUserPublicKeysRequest, runtime *dara.RuntimeOptions) (_result *ListUserPublicKeysResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserPublicKeys"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPublicKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all public keys of the specified user.
//
// @param request - ListUserPublicKeysRequest
//
// @return ListUserPublicKeysResponse
func (client *Client) ListUserPublicKeys(request *ListUserPublicKeysRequest) (_result *ListUserPublicKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserPublicKeysResponse{}
	_body, _err := client.ListUserPublicKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of users of a bastion host.
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
	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceUserId) {
		query["SourceUserId"] = request.SourceUserId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.UserState) {
		query["UserState"] = request.UserState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2019-12-09"),
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
// Queries a list of users of a bastion host.
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
// Locks one or more users of a bastion host.
//
// Description:
//
// # Description
//
// You can call this operation to lock one or more users of a bastion host. If a user does not need to use a bastion host to perform O\\&M operations within a specific period of time, you can lock the user. The locked user can no longer log on to or perform O\\&M operations on the hosts on which the user is granted permissions. If you want to unlock the user later, you can call the [UnlockUsers](https://help.aliyun.com/document_detail/204590.html) operation.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - LockUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockUsersResponse
func (client *Client) LockUsersWithOptions(request *LockUsersRequest, runtime *dara.RuntimeOptions) (_result *LockUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LockUsers"),
		Version:     dara.String("2019-12-09"),
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
// Locks one or more users of a bastion host.
//
// Description:
//
// # Description
//
// You can call this operation to lock one or more users of a bastion host. If a user does not need to use a bastion host to perform O\\&M operations within a specific period of time, you can lock the user. The locked user can no longer log on to or perform O\\&M operations on the hosts on which the user is granted permissions. If you want to unlock the user later, you can call the [UnlockUsers](https://help.aliyun.com/document_detail/204590.html) operation.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Modifies the basic information about a database.
//
// @param request - ModifyDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseResponse
func (client *Client) ModifyDatabaseWithOptions(request *ModifyDatabaseRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatabaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAddressType) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DatabasePort) {
		query["DatabasePort"] = request.DatabasePort
	}

	if !dara.IsNil(request.DatabasePrivateAddress) {
		query["DatabasePrivateAddress"] = request.DatabasePrivateAddress
	}

	if !dara.IsNil(request.DatabasePublicAddress) {
		query["DatabasePublicAddress"] = request.DatabasePublicAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatabase"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a database.
//
// @param request - ModifyDatabaseRequest
//
// @return ModifyDatabaseResponse
func (client *Client) ModifyDatabase(request *ModifyDatabaseRequest) (_result *ModifyDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDatabaseResponse{}
	_body, _err := client.ModifyDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a database account.
//
// @param request - ModifyDatabaseAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseAccountResponse
func (client *Client) ModifyDatabaseAccountWithOptions(request *ModifyDatabaseAccountRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatabaseAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseAccountId) {
		query["DatabaseAccountId"] = request.DatabaseAccountId
	}

	if !dara.IsNil(request.DatabaseAccountName) {
		query["DatabaseAccountName"] = request.DatabaseAccountName
	}

	if !dara.IsNil(request.DatabaseSchema) {
		query["DatabaseSchema"] = request.DatabaseSchema
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatabaseAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatabaseAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a database account.
//
// @param request - ModifyDatabaseAccountRequest
//
// @return ModifyDatabaseAccountResponse
func (client *Client) ModifyDatabaseAccount(request *ModifyDatabaseAccountRequest) (_result *ModifyDatabaseAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDatabaseAccountResponse{}
	_body, _err := client.ModifyDatabaseAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about a host. The information includes the address, name, and description of the host and the operating system that the host runs.
//
// Description:
//
// You can call the ModifyHost operation to modify the basic information about a host in a data center, an Elastic Compute Service (ECS) instance, or a host in an ApsaraDB MyBase dedicated cluster.
//
// > The basic information about ECS instances and hosts in ApsaraDB MyBase dedicated clusters within your Alibaba Cloud account is synchronized to Bastionhost on a regular basis. After you modify the basic information about an ECS instance or a host in an ApsaraDB MyBase dedicated cluster, the modification result may be overwritten by the synchronized information.
//
// @param request - ModifyHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostResponse
func (client *Client) ModifyHostWithOptions(request *ModifyHostRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.HostPrivateAddress) {
		query["HostPrivateAddress"] = request.HostPrivateAddress
	}

	if !dara.IsNil(request.HostPublicAddress) {
		query["HostPublicAddress"] = request.HostPublicAddress
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.PrefKex) {
		query["PrefKex"] = request.PrefKex
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHost"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about a host. The information includes the address, name, and description of the host and the operating system that the host runs.
//
// Description:
//
// You can call the ModifyHost operation to modify the basic information about a host in a data center, an Elastic Compute Service (ECS) instance, or a host in an ApsaraDB MyBase dedicated cluster.
//
// > The basic information about ECS instances and hosts in ApsaraDB MyBase dedicated clusters within your Alibaba Cloud account is synchronized to Bastionhost on a regular basis. After you modify the basic information about an ECS instance or a host in an ApsaraDB MyBase dedicated cluster, the modification result may be overwritten by the synchronized information.
//
// @param request - ModifyHostRequest
//
// @return ModifyHostResponse
func (client *Client) ModifyHost(request *ModifyHostRequest) (_result *ModifyHostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostResponse{}
	_body, _err := client.ModifyHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a host account, such as the username, password, and private key of the host account.
//
// @param request - ModifyHostAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostAccountResponse
func (client *Client) ModifyHostAccountWithOptions(request *ModifyHostAccountRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAccountId) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !dara.IsNil(request.HostAccountName) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PassPhrase) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.PrivilegeType) {
		query["PrivilegeType"] = request.PrivilegeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RotationMode) {
		query["RotationMode"] = request.RotationMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostAccount"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a host account, such as the username, password, and private key of the host account.
//
// @param request - ModifyHostAccountRequest
//
// @return ModifyHostAccountResponse
func (client *Client) ModifyHostAccount(request *ModifyHostAccountRequest) (_result *ModifyHostAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostAccountResponse{}
	_body, _err := client.ModifyHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name or description of the specified host group.
//
// @param request - ModifyHostGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostGroupResponse
func (client *Client) ModifyHostGroupWithOptions(request *ModifyHostGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.HostGroupName) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name or description of the specified host group.
//
// @param request - ModifyHostGroupRequest
//
// @return ModifyHostGroupResponse
func (client *Client) ModifyHostGroup(request *ModifyHostGroupRequest) (_result *ModifyHostGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostGroupResponse{}
	_body, _err := client.ModifyHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a shared key.
//
// @param request - ModifyHostShareKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostShareKeyResponse
func (client *Client) ModifyHostShareKeyWithOptions(request *ModifyHostShareKeyRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostShareKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostShareKeyId) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !dara.IsNil(request.HostShareKeyName) {
		query["HostShareKeyName"] = request.HostShareKeyName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PassPhrase) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostShareKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a shared key.
//
// @param request - ModifyHostShareKeyRequest
//
// @return ModifyHostShareKeyResponse
func (client *Client) ModifyHostShareKey(request *ModifyHostShareKeyRequest) (_result *ModifyHostShareKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostShareKeyResponse{}
	_body, _err := client.ModifyHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the portal type of one or more hosts for O\\&M.
//
// @param request - ModifyHostsActiveAddressTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostsActiveAddressTypeResponse
func (client *Client) ModifyHostsActiveAddressTypeWithOptions(request *ModifyHostsActiveAddressTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostsActiveAddressTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAddressType) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostsActiveAddressType"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostsActiveAddressTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the portal type of one or more hosts for O\\&M.
//
// @param request - ModifyHostsActiveAddressTypeRequest
//
// @return ModifyHostsActiveAddressTypeResponse
func (client *Client) ModifyHostsActiveAddressType(request *ModifyHostsActiveAddressTypeRequest) (_result *ModifyHostsActiveAddressTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostsActiveAddressTypeResponse{}
	_body, _err := client.ModifyHostsActiveAddressTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the port for the O\\\\\\\\\\\\&M protocol on one or more hosts.
//
// Description:
//
// ## Usage notes
//
// You can call this operation to change the port for the O&M protocol on one or more hosts. If the standard port for the O&M protocol on your host is vulnerable to attacks, you can call this operation to specify a custom port. For example, the standard port for SSH is port 22.
//
// >  Ports 0 to 1024 are reserved for Bastionhost. Do not change the port for the O&M protocol to a reserved port.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyHostsPortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostsPortResponse
func (client *Client) ModifyHostsPortWithOptions(request *ModifyHostsPortRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostsPortResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostsPort"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostsPortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the port for the O\\\\\\\\\\\\&M protocol on one or more hosts.
//
// Description:
//
// ## Usage notes
//
// You can call this operation to change the port for the O&M protocol on one or more hosts. If the standard port for the O&M protocol on your host is vulnerable to attacks, you can call this operation to specify a custom port. For example, the standard port for SSH is port 22.
//
// >  Ports 0 to 1024 are reserved for Bastionhost. Do not change the port for the O&M protocol to a reserved port.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyHostsPortRequest
//
// @return ModifyHostsPortResponse
func (client *Client) ModifyHostsPort(request *ModifyHostsPortRequest) (_result *ModifyHostsPortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostsPortResponse{}
	_body, _err := client.ModifyHostsPortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the settings of the Active Directory (AD) authentication server of a bastion host.
//
// @param request - ModifyInstanceADAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceADAuthServerResponse
func (client *Client) ModifyInstanceADAuthServerWithOptions(request *ModifyInstanceADAuthServerRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceADAuthServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.BaseDN) {
		query["BaseDN"] = request.BaseDN
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EmailMapping) {
		query["EmailMapping"] = request.EmailMapping
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsSSL) {
		query["IsSSL"] = request.IsSSL
	}

	if !dara.IsNil(request.MobileMapping) {
		query["MobileMapping"] = request.MobileMapping
	}

	if !dara.IsNil(request.NameMapping) {
		query["NameMapping"] = request.NameMapping
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Server) {
		query["Server"] = request.Server
	}

	if !dara.IsNil(request.StandbyServer) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceADAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of the Active Directory (AD) authentication server of a bastion host.
//
// @param request - ModifyInstanceADAuthServerRequest
//
// @return ModifyInstanceADAuthServerResponse
func (client *Client) ModifyInstanceADAuthServer(request *ModifyInstanceADAuthServerRequest) (_result *ModifyInstanceADAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceADAuthServerResponse{}
	_body, _err := client.ModifyInstanceADAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a bastion host.
//
// @param request - ModifyInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAttribute"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a bastion host.
//
// @param request - ModifyInstanceAttributeRequest
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the settings of the Lightweight Directory Access Protocol (LDAP) authentication server of a bastion host.
//
// @param request - ModifyInstanceLDAPAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceLDAPAuthServerResponse
func (client *Client) ModifyInstanceLDAPAuthServerWithOptions(request *ModifyInstanceLDAPAuthServerRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceLDAPAuthServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.BaseDN) {
		query["BaseDN"] = request.BaseDN
	}

	if !dara.IsNil(request.EmailMapping) {
		query["EmailMapping"] = request.EmailMapping
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsSSL) {
		query["IsSSL"] = request.IsSSL
	}

	if !dara.IsNil(request.LoginNameMapping) {
		query["LoginNameMapping"] = request.LoginNameMapping
	}

	if !dara.IsNil(request.MobileMapping) {
		query["MobileMapping"] = request.MobileMapping
	}

	if !dara.IsNil(request.NameMapping) {
		query["NameMapping"] = request.NameMapping
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Server) {
		query["Server"] = request.Server
	}

	if !dara.IsNil(request.StandbyServer) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceLDAPAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of the Lightweight Directory Access Protocol (LDAP) authentication server of a bastion host.
//
// @param request - ModifyInstanceLDAPAuthServerRequest
//
// @return ModifyInstanceLDAPAuthServerResponse
func (client *Client) ModifyInstanceLDAPAuthServer(request *ModifyInstanceLDAPAuthServerRequest) (_result *ModifyInstanceLDAPAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.ModifyInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the two-factor authentication settings of a bastion host.
//
// @param request - ModifyInstanceTwoFactorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceTwoFactorResponse
func (client *Client) ModifyInstanceTwoFactorWithOptions(request *ModifyInstanceTwoFactorRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceTwoFactorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableTwoFactor) {
		query["EnableTwoFactor"] = request.EnableTwoFactor
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SkipTwoFactorTime) {
		query["SkipTwoFactorTime"] = request.SkipTwoFactorTime
	}

	if !dara.IsNil(request.TwoFactorMethods) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceTwoFactor"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceTwoFactorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the two-factor authentication settings of a bastion host.
//
// @param request - ModifyInstanceTwoFactorRequest
//
// @return ModifyInstanceTwoFactorResponse
func (client *Client) ModifyInstanceTwoFactor(request *ModifyInstanceTwoFactorRequest) (_result *ModifyInstanceTwoFactorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceTwoFactorResponse{}
	_body, _err := client.ModifyInstanceTwoFactorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a network domain.
//
// @param request - ModifyNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNetworkDomainResponse
func (client *Client) ModifyNetworkDomainWithOptions(request *ModifyNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.NetworkDomainName) {
		query["NetworkDomainName"] = request.NetworkDomainName
	}

	if !dara.IsNil(request.NetworkDomainType) {
		query["NetworkDomainType"] = request.NetworkDomainType
	}

	if !dara.IsNil(request.Proxies) {
		query["Proxies"] = request.Proxies
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a network domain.
//
// @param request - ModifyNetworkDomainRequest
//
// @return ModifyNetworkDomainResponse
func (client *Client) ModifyNetworkDomain(request *ModifyNetworkDomainRequest) (_result *ModifyNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNetworkDomainResponse{}
	_body, _err := client.ModifyNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a control policy.
//
// @param request - ModifyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyResponse
func (client *Client) ModifyPolicyWithOptions(request *ModifyPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolicy"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about a control policy.
//
// @param request - ModifyPolicyRequest
//
// @return ModifyPolicyResponse
func (client *Client) ModifyPolicy(request *ModifyPolicyRequest) (_result *ModifyPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.ModifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information of an authorization rule.
//
// @param request - ModifyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleResponse
func (client *Client) ModifyRuleWithOptions(request *ModifyRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRule"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information of an authorization rule.
//
// @param request - ModifyRuleRequest
//
// @return ModifyRuleResponse
func (client *Client) ModifyRule(request *ModifyRuleRequest) (_result *ModifyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRuleResponse{}
	_body, _err := client.ModifyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a user of a bastion host.
//
// @param request - ModifyUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserResponse
func (client *Client) ModifyUserWithOptions(request *ModifyUserRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.LanguageStatus) {
		query["LanguageStatus"] = request.LanguageStatus
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.MobileCountryCode) {
		query["MobileCountryCode"] = request.MobileCountryCode
	}

	if !dara.IsNil(request.NeedResetPassword) {
		query["NeedResetPassword"] = request.NeedResetPassword
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TwoFactorMethods) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	if !dara.IsNil(request.TwoFactorStatus) {
		query["TwoFactorStatus"] = request.TwoFactorStatus
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUser"),
		Version:     dara.String("2019-12-09"),
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
// Modifies the information about a user of a bastion host.
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
// Modifies the information about the specified user group.
//
// @param request - ModifyUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserGroupResponse
func (client *Client) ModifyUserGroupWithOptions(request *ModifyUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserGroupName) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about the specified user group.
//
// @param request - ModifyUserGroupRequest
//
// @return ModifyUserGroupResponse
func (client *Client) ModifyUserGroup(request *ModifyUserGroupRequest) (_result *ModifyUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserGroupResponse{}
	_body, _err := client.ModifyUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the public key of the user.
//
// @param request - ModifyUserPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserPublicKeyResponse
func (client *Client) ModifyUserPublicKeyWithOptions(request *ModifyUserPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserPublicKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PublicKey) {
		query["PublicKey"] = request.PublicKey
	}

	if !dara.IsNil(request.PublicKeyId) {
		query["PublicKeyId"] = request.PublicKeyId
	}

	if !dara.IsNil(request.PublicKeyName) {
		query["PublicKeyName"] = request.PublicKeyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserPublicKey"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the public key of the user.
//
// @param request - ModifyUserPublicKeyRequest
//
// @return ModifyUserPublicKeyResponse
func (client *Client) ModifyUserPublicKey(request *ModifyUserPublicKeyRequest) (_result *ModifyUserPublicKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserPublicKeyResponse{}
	_body, _err := client.ModifyUserPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds multiple databases to a network domain at a time.
//
// @param request - MoveDatabasesToNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveDatabasesToNetworkDomainResponse
func (client *Client) MoveDatabasesToNetworkDomainWithOptions(request *MoveDatabasesToNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *MoveDatabasesToNetworkDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseIds) {
		query["DatabaseIds"] = request.DatabaseIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveDatabasesToNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveDatabasesToNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds multiple databases to a network domain at a time.
//
// @param request - MoveDatabasesToNetworkDomainRequest
//
// @return MoveDatabasesToNetworkDomainResponse
func (client *Client) MoveDatabasesToNetworkDomain(request *MoveDatabasesToNetworkDomainRequest) (_result *MoveDatabasesToNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveDatabasesToNetworkDomainResponse{}
	_body, _err := client.MoveDatabasesToNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds multiple hosts to a network domain at a time.
//
// @param request - MoveHostsToNetworkDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveHostsToNetworkDomainResponse
func (client *Client) MoveHostsToNetworkDomainWithOptions(request *MoveHostsToNetworkDomainRequest, runtime *dara.RuntimeOptions) (_result *MoveHostsToNetworkDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkDomainId) {
		query["NetworkDomainId"] = request.NetworkDomainId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveHostsToNetworkDomain"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveHostsToNetworkDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds multiple hosts to a network domain at a time.
//
// @param request - MoveHostsToNetworkDomainRequest
//
// @return MoveHostsToNetworkDomainResponse
func (client *Client) MoveHostsToNetworkDomain(request *MoveHostsToNetworkDomainRequest) (_result *MoveHostsToNetworkDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveHostsToNetworkDomainResponse{}
	_body, _err := client.MoveHostsToNetworkDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a bastion host from one resource group to another resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a bastion host from one resource group to another resource group.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If an O\\&M engineer attempts to run a command specified in the Command Approval section of the Create Control Policy page, the administrator is notified to review the command in the Bastionhost console. The command can be run only after it is approved by the administrator.
//
// Description:
//
// You can call this operation as a Bastionhost administrator to reject the request to run a command of an O\\&M engineer.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - RejectApproveCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectApproveCommandResponse
func (client *Client) RejectApproveCommandWithOptions(request *RejectApproveCommandRequest, runtime *dara.RuntimeOptions) (_result *RejectApproveCommandResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandId) {
		query["CommandId"] = request.CommandId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectApproveCommand"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectApproveCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If an O\\&M engineer attempts to run a command specified in the Command Approval section of the Create Control Policy page, the administrator is notified to review the command in the Bastionhost console. The command can be run only after it is approved by the administrator.
//
// Description:
//
// You can call this operation as a Bastionhost administrator to reject the request to run a command of an O\\&M engineer.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - RejectApproveCommandRequest
//
// @return RejectApproveCommandResponse
func (client *Client) RejectApproveCommand(request *RejectApproveCommandRequest) (_result *RejectApproveCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectApproveCommandResponse{}
	_body, _err := client.RejectApproveCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If a Bastionhost administrator enables O\\\\\\&M Approval on the Create Control Policy page, O\\\\\\&M engineers can log on to assets to perform O\\\\\\&M operations only after the administrator approves their O\\\\\\&M applications.
//
// Description:
//
// You can call this operation to reject an O\\&M application of an O\\&M engineer as a Bastionhost administrator.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - RejectOperationTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectOperationTicketResponse
func (client *Client) RejectOperationTicketWithOptions(request *RejectOperationTicketRequest, runtime *dara.RuntimeOptions) (_result *RejectOperationTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OperationTicketId) {
		query["OperationTicketId"] = request.OperationTicketId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectOperationTicket"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectOperationTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If a Bastionhost administrator enables O\\\\\\&M Approval on the Create Control Policy page, O\\\\\\&M engineers can log on to assets to perform O\\\\\\&M operations only after the administrator approves their O\\\\\\&M applications.
//
// Description:
//
// You can call this operation to reject an O\\&M application of an O\\&M engineer as a Bastionhost administrator.
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - RejectOperationTicketRequest
//
// @return RejectOperationTicketResponse
func (client *Client) RejectOperationTicket(request *RejectOperationTicketRequest) (_result *RejectOperationTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectOperationTicketResponse{}
	_body, _err := client.RejectOperationTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes multiple databases from an asset group at a time.
//
// @param request - RemoveDatabasesFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDatabasesFromGroupResponse
func (client *Client) RemoveDatabasesFromGroupWithOptions(request *RemoveDatabasesFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveDatabasesFromGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseIds) {
		query["DatabaseIds"] = request.DatabaseIds
	}

	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDatabasesFromGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDatabasesFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes multiple databases from an asset group at a time.
//
// @param request - RemoveDatabasesFromGroupRequest
//
// @return RemoveDatabasesFromGroupResponse
func (client *Client) RemoveDatabasesFromGroup(request *RemoveDatabasesFromGroupRequest) (_result *RemoveDatabasesFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveDatabasesFromGroupResponse{}
	_body, _err := client.RemoveDatabasesFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes multiple hosts from an asset group at a time.
//
// Description:
//
// You can call the RemoveHostsFromGroup operation to remove multiple hosts from an asset group at a time. If you no longer need to manage some hosts in an asset group, you can call this operation to remove the hosts from the asset group.
//
// # [](#qps-)QPS limit
//
// You can call this API operation up to 10 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - RemoveHostsFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveHostsFromGroupResponse
func (client *Client) RemoveHostsFromGroupWithOptions(request *RemoveHostsFromGroupRequest, runtime *dara.RuntimeOptions) (_result *RemoveHostsFromGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostGroupId) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !dara.IsNil(request.HostIds) {
		query["HostIds"] = request.HostIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveHostsFromGroup"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveHostsFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes multiple hosts from an asset group at a time.
//
// Description:
//
// You can call the RemoveHostsFromGroup operation to remove multiple hosts from an asset group at a time. If you no longer need to manage some hosts in an asset group, you can call this operation to remove the hosts from the asset group.
//
// # [](#qps-)QPS limit
//
// You can call this API operation up to 10 times per second per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - RemoveHostsFromGroupRequest
//
// @return RemoveHostsFromGroupResponse
func (client *Client) RemoveHostsFromGroup(request *RemoveHostsFromGroupRequest) (_result *RemoveHostsFromGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveHostsFromGroupResponse{}
	_body, _err := client.RemoveHostsFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除RD成员账号
//
// @param request - RemoveInstanceRdMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInstanceRdMemberResponse
func (client *Client) RemoveInstanceRdMemberWithOptions(request *RemoveInstanceRdMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveInstanceRdMemberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveInstanceRdMember"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveInstanceRdMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除RD成员账号
//
// @param request - RemoveInstanceRdMemberRequest
//
// @return RemoveInstanceRdMemberResponse
func (client *Client) RemoveInstanceRdMember(request *RemoveInstanceRdMemberRequest) (_result *RemoveInstanceRdMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveInstanceRdMemberResponse{}
	_body, _err := client.RemoveInstanceRdMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes one or more users from a user group.
//
// Description:
//
// You can call this operation to remove one or more users from a user group. When users in a user group are transferred to a new position, resign, or are switched to another user group, you can call this operation to remove the users from the current user group at a time.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsersFromGroup"),
		Version:     dara.String("2019-12-09"),
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
// Removes one or more users from a user group.
//
// Description:
//
// You can call this operation to remove one or more users from a user group. When users in a user group are transferred to a new position, resign, or are switched to another user group, you can call this operation to remove the users from the current user group at a time.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Renews an O\\&M token for one hour.
//
// @param request - RenewAssetOperationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAssetOperationTokenResponse
func (client *Client) RenewAssetOperationTokenWithOptions(request *RenewAssetOperationTokenRequest, runtime *dara.RuntimeOptions) (_result *RenewAssetOperationTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TokenId) {
		query["TokenId"] = request.TokenId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewAssetOperationToken"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewAssetOperationTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews an O\\&M token for one hour.
//
// @param request - RenewAssetOperationTokenRequest
//
// @return RenewAssetOperationTokenResponse
func (client *Client) RenewAssetOperationToken(request *RenewAssetOperationTokenRequest) (_result *RenewAssetOperationTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewAssetOperationTokenResponse{}
	_body, _err := client.RenewAssetOperationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the logon credential of a specified host account. The logon credential can be the password or Secure Shell (SSH) private key.
//
// @param request - ResetHostAccountCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetHostAccountCredentialResponse
func (client *Client) ResetHostAccountCredentialWithOptions(request *ResetHostAccountCredentialRequest, runtime *dara.RuntimeOptions) (_result *ResetHostAccountCredentialResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialType) {
		query["CredentialType"] = request.CredentialType
	}

	if !dara.IsNil(request.HostAccountId) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetHostAccountCredential"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetHostAccountCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the logon credential of a specified host account. The logon credential can be the password or Secure Shell (SSH) private key.
//
// @param request - ResetHostAccountCredentialRequest
//
// @return ResetHostAccountCredentialResponse
func (client *Client) ResetHostAccountCredential(request *ResetHostAccountCredentialRequest) (_result *ResetHostAccountCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetHostAccountCredentialResponse{}
	_body, _err := client.ResetHostAccountCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the logon period limits in a control policy.
//
// @param tmpReq - SetPolicyAccessTimeRangeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyAccessTimeRangeConfigResponse
func (client *Client) SetPolicyAccessTimeRangeConfigWithOptions(tmpReq *SetPolicyAccessTimeRangeConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyAccessTimeRangeConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetPolicyAccessTimeRangeConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccessTimeRangeConfig) {
		request.AccessTimeRangeConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessTimeRangeConfig, dara.String("AccessTimeRangeConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessTimeRangeConfigShrink) {
		query["AccessTimeRangeConfig"] = request.AccessTimeRangeConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyAccessTimeRangeConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyAccessTimeRangeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the logon period limits in a control policy.
//
// @param request - SetPolicyAccessTimeRangeConfigRequest
//
// @return SetPolicyAccessTimeRangeConfigResponse
func (client *Client) SetPolicyAccessTimeRangeConfig(request *SetPolicyAccessTimeRangeConfigRequest) (_result *SetPolicyAccessTimeRangeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyAccessTimeRangeConfigResponse{}
	_body, _err := client.SetPolicyAccessTimeRangeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the O&M approval setting in a control policy.
//
// @param tmpReq - SetPolicyApprovalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyApprovalConfigResponse
func (client *Client) SetPolicyApprovalConfigWithOptions(tmpReq *SetPolicyApprovalConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyApprovalConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetPolicyApprovalConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApprovalConfig) {
		request.ApprovalConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovalConfig, dara.String("ApprovalConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalConfigShrink) {
		query["ApprovalConfig"] = request.ApprovalConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyApprovalConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyApprovalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the O&M approval setting in a control policy.
//
// @param request - SetPolicyApprovalConfigRequest
//
// @return SetPolicyApprovalConfigResponse
func (client *Client) SetPolicyApprovalConfig(request *SetPolicyApprovalConfigRequest) (_result *SetPolicyApprovalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyApprovalConfigResponse{}
	_body, _err := client.SetPolicyApprovalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies the assets to which a control policy applies.
//
// @param request - SetPolicyAssetScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyAssetScopeResponse
func (client *Client) SetPolicyAssetScopeWithOptions(request *SetPolicyAssetScopeRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyAssetScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Databases) {
		query["Databases"] = request.Databases
	}

	if !dara.IsNil(request.HostGroups) {
		query["HostGroups"] = request.HostGroups
	}

	if !dara.IsNil(request.Hosts) {
		query["Hosts"] = request.Hosts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyAssetScope"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyAssetScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the assets to which a control policy applies.
//
// @param request - SetPolicyAssetScopeRequest
//
// @return SetPolicyAssetScopeResponse
func (client *Client) SetPolicyAssetScope(request *SetPolicyAssetScopeRequest) (_result *SetPolicyAssetScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyAssetScopeResponse{}
	_body, _err := client.SetPolicyAssetScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies the commands that can or cannot be run by the users or on the assets associated with the policy and the commands that must be reviewed.
//
// @param tmpReq - SetPolicyCommandConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyCommandConfigResponse
func (client *Client) SetPolicyCommandConfigWithOptions(tmpReq *SetPolicyCommandConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyCommandConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetPolicyCommandConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CommandConfig) {
		request.CommandConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommandConfig, dara.String("CommandConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandConfigShrink) {
		query["CommandConfig"] = request.CommandConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyCommandConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyCommandConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the commands that can or cannot be run by the users or on the assets associated with the policy and the commands that must be reviewed.
//
// @param request - SetPolicyCommandConfigRequest
//
// @return SetPolicyCommandConfigResponse
func (client *Client) SetPolicyCommandConfig(request *SetPolicyCommandConfigRequest) (_result *SetPolicyCommandConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyCommandConfigResponse{}
	_body, _err := client.SetPolicyCommandConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures access control settings in a control policy.
//
// @param tmpReq - SetPolicyIPAclConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyIPAclConfigResponse
func (client *Client) SetPolicyIPAclConfigWithOptions(tmpReq *SetPolicyIPAclConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyIPAclConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetPolicyIPAclConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IPAclConfig) {
		request.IPAclConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IPAclConfig, dara.String("IPAclConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IPAclConfigShrink) {
		query["IPAclConfig"] = request.IPAclConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyIPAclConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyIPAclConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures access control settings in a control policy.
//
// @param request - SetPolicyIPAclConfigRequest
//
// @return SetPolicyIPAclConfigResponse
func (client *Client) SetPolicyIPAclConfig(request *SetPolicyIPAclConfigRequest) (_result *SetPolicyIPAclConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyIPAclConfigResponse{}
	_body, _err := client.SetPolicyIPAclConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the protocol control settings in a control policy.
//
// @param tmpReq - SetPolicyProtocolConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyProtocolConfigResponse
func (client *Client) SetPolicyProtocolConfigWithOptions(tmpReq *SetPolicyProtocolConfigRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyProtocolConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetPolicyProtocolConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProtocolConfig) {
		request.ProtocolConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProtocolConfig, dara.String("ProtocolConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.ProtocolConfigShrink) {
		query["ProtocolConfig"] = request.ProtocolConfigShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyProtocolConfig"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyProtocolConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the protocol control settings in a control policy.
//
// @param request - SetPolicyProtocolConfigRequest
//
// @return SetPolicyProtocolConfigResponse
func (client *Client) SetPolicyProtocolConfig(request *SetPolicyProtocolConfigRequest) (_result *SetPolicyProtocolConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyProtocolConfigResponse{}
	_body, _err := client.SetPolicyProtocolConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies the users to whom a control policy applies.
//
// @param request - SetPolicyUserScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPolicyUserScopeResponse
func (client *Client) SetPolicyUserScopeWithOptions(request *SetPolicyUserScopeRequest, runtime *dara.RuntimeOptions) (_result *SetPolicyUserScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetPolicyUserScope"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetPolicyUserScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the users to whom a control policy applies.
//
// @param request - SetPolicyUserScopeRequest
//
// @return SetPolicyUserScopeResponse
func (client *Client) SetPolicyUserScope(request *SetPolicyUserScopeRequest) (_result *SetPolicyUserScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetPolicyUserScopeResponse{}
	_body, _err := client.SetPolicyUserScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the specified bastion host.
//
// @param request - StartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientSecurityGroupIds) {
		query["ClientSecurityGroupIds"] = request.ClientSecurityGroupIds
	}

	if !dara.IsNil(request.EnablePortalPrivateAccess) {
		query["EnablePortalPrivateAccess"] = request.EnablePortalPrivateAccess
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupIds) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	if !dara.IsNil(request.SlaveVswitchId) {
		query["SlaveVswitchId"] = request.SlaveVswitchId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartInstance"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the specified bastion host.
//
// @param request - StartInstanceRequest
//
// @return StartInstanceResponse
func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and adds tags to specified bastion hosts.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Version:     dara.String("2019-12-09"),
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
// Creates and adds tags to specified bastion hosts.
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
// Unlocks one or more users of a bastion host.
//
// Description:
//
// After you call the [LockUsers](https://help.aliyun.com/document_detail/204591.html) operation to lock one or more users of a bastion host, you can call this operation to unlock the users. After the users are unlocked, the users can perform O\\&M operations by using the bastion host.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UnlockUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockUsersResponse
func (client *Client) UnlockUsersWithOptions(request *UnlockUsersRequest, runtime *dara.RuntimeOptions) (_result *UnlockUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockUsers"),
		Version:     dara.String("2019-12-09"),
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
// Unlocks one or more users of a bastion host.
//
// Description:
//
// After you call the [LockUsers](https://help.aliyun.com/document_detail/204591.html) operation to lock one or more users of a bastion host, you can call this operation to unlock the users. After the users are unlocked, the users can perform O\\&M operations by using the bastion host.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Removes tags from the specified bastion host and deletes the tags at a time.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Version:     dara.String("2019-12-09"),
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
// Removes tags from the specified bastion host and deletes the tags at a time.
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
// 验证实例AD服务配置
//
// @param request - VerifyInstanceADAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyInstanceADAuthServerResponse
func (client *Client) VerifyInstanceADAuthServerWithOptions(request *VerifyInstanceADAuthServerRequest, runtime *dara.RuntimeOptions) (_result *VerifyInstanceADAuthServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.BaseDN) {
		query["BaseDN"] = request.BaseDN
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsSSL) {
		query["IsSSL"] = request.IsSSL
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Server) {
		query["Server"] = request.Server
	}

	if !dara.IsNil(request.StandbyServer) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyInstanceADAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证实例AD服务配置
//
// @param request - VerifyInstanceADAuthServerRequest
//
// @return VerifyInstanceADAuthServerResponse
func (client *Client) VerifyInstanceADAuthServer(request *VerifyInstanceADAuthServerRequest) (_result *VerifyInstanceADAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyInstanceADAuthServerResponse{}
	_body, _err := client.VerifyInstanceADAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证实例LDAP服务配置
//
// @param request - VerifyInstanceLDAPAuthServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyInstanceLDAPAuthServerResponse
func (client *Client) VerifyInstanceLDAPAuthServerWithOptions(request *VerifyInstanceLDAPAuthServerRequest, runtime *dara.RuntimeOptions) (_result *VerifyInstanceLDAPAuthServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.BaseDN) {
		query["BaseDN"] = request.BaseDN
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsSSL) {
		query["IsSSL"] = request.IsSSL
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Server) {
		query["Server"] = request.Server
	}

	if !dara.IsNil(request.StandbyServer) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyInstanceLDAPAuthServer"),
		Version:     dara.String("2019-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证实例LDAP服务配置
//
// @param request - VerifyInstanceLDAPAuthServerRequest
//
// @return VerifyInstanceLDAPAuthServerResponse
func (client *Client) VerifyInstanceLDAPAuthServer(request *VerifyInstanceLDAPAuthServerRequest) (_result *VerifyInstanceLDAPAuthServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.VerifyInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
