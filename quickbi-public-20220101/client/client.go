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
	client.Endpoint, _err = client.GetEndpoint(dara.String("quickbi-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Add selected groups of people incrementally for a single row and column permission rule.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.\\n
//
// @param request - AddDataLevelPermissionRuleUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataLevelPermissionRuleUsersResponse
func (client *Client) AddDataLevelPermissionRuleUsersWithOptions(request *AddDataLevelPermissionRuleUsersRequest, runtime *dara.RuntimeOptions) (_result *AddDataLevelPermissionRuleUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddUserModel) {
		query["AddUserModel"] = request.AddUserModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDataLevelPermissionRuleUsers"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataLevelPermissionRuleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add selected groups of people incrementally for a single row and column permission rule.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.\\n
//
// @param request - AddDataLevelPermissionRuleUsersRequest
//
// @return AddDataLevelPermissionRuleUsersResponse
func (client *Client) AddDataLevelPermissionRuleUsers(request *AddDataLevelPermissionRuleUsersRequest) (_result *AddDataLevelPermissionRuleUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDataLevelPermissionRuleUsersResponse{}
	_body, _err := client.AddDataLevelPermissionRuleUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 43342***435,1553a****41231
//
// Description:
//
// ROW_LEVEL
//
// @param request - AddDataLevelPermissionWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataLevelPermissionWhiteListResponse
func (client *Client) AddDataLevelPermissionWhiteListWithOptions(request *AddDataLevelPermissionWhiteListRequest, runtime *dara.RuntimeOptions) (_result *AddDataLevelPermissionWhiteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
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
		Action:      dara.String("AddDataLevelPermissionWhiteList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataLevelPermissionWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 43342***435,1553a****41231
//
// Description:
//
// ROW_LEVEL
//
// @param request - AddDataLevelPermissionWhiteListRequest
//
// @return AddDataLevelPermissionWhiteListResponse
func (client *Client) AddDataLevelPermissionWhiteList(request *AddDataLevelPermissionWhiteListRequest) (_result *AddDataLevelPermissionWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDataLevelPermissionWhiteListResponse{}
	_body, _err := client.AddDataLevelPermissionWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a sharing configuration for data works.
//
// @param request - AddShareReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShareReportResponse
func (client *Client) AddShareReportWithOptions(request *AddShareReportRequest, runtime *dara.RuntimeOptions) (_result *AddShareReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthPoint) {
		query["AuthPoint"] = request.AuthPoint
	}

	if !dara.IsNil(request.ExpireDate) {
		query["ExpireDate"] = request.ExpireDate
	}

	if !dara.IsNil(request.ShareToId) {
		query["ShareToId"] = request.ShareToId
	}

	if !dara.IsNil(request.ShareToType) {
		query["ShareToType"] = request.ShareToType
	}

	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddShareReport"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddShareReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a sharing configuration for data works.
//
// @param request - AddShareReportRequest
//
// @return AddShareReportResponse
func (client *Client) AddShareReport(request *AddShareReportRequest) (_result *AddShareReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddShareReportResponse{}
	_body, _err := client.AddShareReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add an organization member.
//
// @param request - AddUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserResponse
func (client *Client) AddUserWithOptions(request *AddUserRequest, runtime *dara.RuntimeOptions) (_result *AddUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AdminUser) {
		query["AdminUser"] = request.AdminUser
	}

	if !dara.IsNil(request.AuthAdminUser) {
		query["AuthAdminUser"] = request.AuthAdminUser
	}

	if !dara.IsNil(request.NickName) {
		query["NickName"] = request.NickName
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleIds) {
		body["RoleIds"] = request.RoleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUser"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add an organization member.
//
// @param request - AddUserRequest
//
// @return AddUserResponse
func (client *Client) AddUser(request *AddUserRequest) (_result *AddUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserResponse{}
	_body, _err := client.AddUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an organization member to a specified user group.
//
// @param request - AddUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMemberWithOptions(request *AddUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *AddUserGroupMemberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserIdList) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserGroupMember"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an organization member to a specified user group.
//
// @param request - AddUserGroupMemberRequest
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMember(request *AddUserGroupMemberRequest) (_result *AddUserGroupMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserGroupMemberResponse{}
	_body, _err := client.AddUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add users to a specified user group at a time.
//
// @param request - AddUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMembersResponse
func (client *Client) AddUserGroupMembersWithOptions(request *AddUserGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *AddUserGroupMembersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserGroupMembers"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add users to a specified user group at a time.
//
// @param request - AddUserGroupMembersRequest
//
// @return AddUserGroupMembersResponse
func (client *Client) AddUserGroupMembers(request *AddUserGroupMembersRequest) (_result *AddUserGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserGroupMembersResponse{}
	_body, _err := client.AddUserGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add organization member tag metadata.
//
// @param request - AddUserTagMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserTagMetaResponse
func (client *Client) AddUserTagMetaWithOptions(request *AddUserTagMetaRequest, runtime *dara.RuntimeOptions) (_result *AddUserTagMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TagDescription) {
		query["TagDescription"] = request.TagDescription
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserTagMeta"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserTagMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add organization member tag metadata.
//
// @param request - AddUserTagMetaRequest
//
// @return AddUserTagMetaResponse
func (client *Client) AddUserTagMeta(request *AddUserTagMetaRequest) (_result *AddUserTagMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserTagMetaResponse{}
	_body, _err := client.AddUserTagMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a member to the specified workspace.
//
// @param request - AddUserToWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToWorkspaceResponse
func (client *Client) AddUserToWorkspaceWithOptions(request *AddUserToWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *AddUserToWorkspaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToWorkspace"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a member to the specified workspace.
//
// @param request - AddUserToWorkspaceRequest
//
// @return AddUserToWorkspaceResponse
func (client *Client) AddUserToWorkspace(request *AddUserToWorkspaceRequest) (_result *AddUserToWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserToWorkspaceResponse{}
	_body, _err := client.AddUserToWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch add members to the workspace.
//
// @param request - AddWorkspaceUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWorkspaceUsersResponse
func (client *Client) AddWorkspaceUsersWithOptions(request *AddWorkspaceUsersRequest, runtime *dara.RuntimeOptions) (_result *AddWorkspaceUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWorkspaceUsers"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWorkspaceUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch add members to the workspace.
//
// @param request - AddWorkspaceUsersRequest
//
// @return AddWorkspaceUsersResponse
func (client *Client) AddWorkspaceUsers(request *AddWorkspaceUsersRequest) (_result *AddWorkspaceUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddWorkspaceUsersResponse{}
	_body, _err := client.AddWorkspaceUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Trigger the collection acceleration of the Quick engine for datasets.
//
// @param request - AllotDatasetAccelerationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllotDatasetAccelerationTaskResponse
func (client *Client) AllotDatasetAccelerationTaskWithOptions(request *AllotDatasetAccelerationTaskRequest, runtime *dara.RuntimeOptions) (_result *AllotDatasetAccelerationTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllotDatasetAccelerationTask"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllotDatasetAccelerationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Trigger the collection acceleration of the Quick engine for datasets.
//
// @param request - AllotDatasetAccelerationTaskRequest
//
// @return AllotDatasetAccelerationTaskResponse
func (client *Client) AllotDatasetAccelerationTask(request *AllotDatasetAccelerationTaskRequest) (_result *AllotDatasetAccelerationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllotDatasetAccelerationTaskResponse{}
	_body, _err := client.AllotDatasetAccelerationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch authorization of BI portal menu will be skipped automatically.
//
// @param request - AuthorizeMenuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeMenuResponse
func (client *Client) AuthorizeMenuWithOptions(request *AuthorizeMenuRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeMenuResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthPointsValue) {
		query["AuthPointsValue"] = request.AuthPointsValue
	}

	if !dara.IsNil(request.DataPortalId) {
		query["DataPortalId"] = request.DataPortalId
	}

	if !dara.IsNil(request.MenuIds) {
		query["MenuIds"] = request.MenuIds
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
		Action:      dara.String("AuthorizeMenu"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeMenuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch authorization of BI portal menu will be skipped automatically.
//
// @param request - AuthorizeMenuRequest
//
// @return AuthorizeMenuResponse
func (client *Client) AuthorizeMenu(request *AuthorizeMenuRequest) (_result *AuthorizeMenuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthorizeMenuResponse{}
	_body, _err := client.AuthorizeMenuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI BatchAddFeishuUsers is deprecated
//
// Summary:
//
// 批量添加飞书用户。
//
// @param request - BatchAddFeishuUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddFeishuUsersResponse
// Deprecated
func (client *Client) BatchAddFeishuUsersWithOptions(request *BatchAddFeishuUsersRequest, runtime *dara.RuntimeOptions) (_result *BatchAddFeishuUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeishuUsers) {
		query["FeishuUsers"] = request.FeishuUsers
	}

	if !dara.IsNil(request.IsAdmin) {
		query["IsAdmin"] = request.IsAdmin
	}

	if !dara.IsNil(request.IsAuthAdmin) {
		query["IsAuthAdmin"] = request.IsAuthAdmin
	}

	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAddFeishuUsers"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAddFeishuUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI BatchAddFeishuUsers is deprecated
//
// Summary:
//
// 批量添加飞书用户。
//
// @param request - BatchAddFeishuUsersRequest
//
// @return BatchAddFeishuUsersResponse
// Deprecated
func (client *Client) BatchAddFeishuUsers(request *BatchAddFeishuUsersRequest) (_result *BatchAddFeishuUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchAddFeishuUsersResponse{}
	_body, _err := client.BatchAddFeishuUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancel the authorization records for specified users and user groups based on the portal menu ID.
//
// @param request - CancelAuthorizationMenuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAuthorizationMenuResponse
func (client *Client) CancelAuthorizationMenuWithOptions(request *CancelAuthorizationMenuRequest, runtime *dara.RuntimeOptions) (_result *CancelAuthorizationMenuResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataPortalId) {
		query["DataPortalId"] = request.DataPortalId
	}

	if !dara.IsNil(request.MenuIds) {
		query["MenuIds"] = request.MenuIds
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
		Action:      dara.String("CancelAuthorizationMenu"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAuthorizationMenuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancel the authorization records for specified users and user groups based on the portal menu ID.
//
// @param request - CancelAuthorizationMenuRequest
//
// @return CancelAuthorizationMenuResponse
func (client *Client) CancelAuthorizationMenu(request *CancelAuthorizationMenuRequest) (_result *CancelAuthorizationMenuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelAuthorizationMenuResponse{}
	_body, _err := client.CancelAuthorizationMenuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancel the data works from the user\\"s collection.
//
// @param request - CancelCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCollectionResponse
func (client *Client) CancelCollectionWithOptions(request *CancelCollectionRequest, runtime *dara.RuntimeOptions) (_result *CancelCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCollection"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancel the data works from the user\\"s collection.
//
// @param request - CancelCollectionRequest
//
// @return CancelCollectionResponse
func (client *Client) CancelCollection(request *CancelCollectionRequest) (_result *CancelCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCollectionResponse{}
	_body, _err := client.CancelCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a share authorization for a data work.
//
// @param request - CancelReportShareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelReportShareResponse
func (client *Client) CancelReportShareWithOptions(request *CancelReportShareRequest, runtime *dara.RuntimeOptions) (_result *CancelReportShareResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.ShareToIds) {
		query["ShareToIds"] = request.ShareToIds
	}

	if !dara.IsNil(request.ShareToType) {
		query["ShareToType"] = request.ShareToType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelReportShare"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelReportShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a share authorization for a data work.
//
// @param request - CancelReportShareRequest
//
// @return CancelReportShareResponse
func (client *Client) CancelReportShare(request *CancelReportShareRequest) (_result *CancelReportShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelReportShareResponse{}
	_body, _err := client.CancelReportShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the visibility mode of the BI portal menu and whether the menu is only authorized to be visible.
//
// @param request - ChangeVisibilityModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeVisibilityModelResponse
func (client *Client) ChangeVisibilityModelWithOptions(request *ChangeVisibilityModelRequest, runtime *dara.RuntimeOptions) (_result *ChangeVisibilityModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataPortalId) {
		query["DataPortalId"] = request.DataPortalId
	}

	if !dara.IsNil(request.MenuIds) {
		query["MenuIds"] = request.MenuIds
	}

	if !dara.IsNil(request.ShowOnlyWithAccess) {
		query["ShowOnlyWithAccess"] = request.ShowOnlyWithAccess
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeVisibilityModel"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeVisibilityModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the visibility mode of the BI portal menu and whether the menu is only authorized to be visible.
//
// @param request - ChangeVisibilityModelRequest
//
// @return ChangeVisibilityModelResponse
func (client *Client) ChangeVisibilityModel(request *ChangeVisibilityModelRequest) (_result *ChangeVisibilityModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeVisibilityModelResponse{}
	_body, _err := client.ChangeVisibilityModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether a user has permissions to view data works, such as dashboards and workbooks.
//
// @param request - CheckReadableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckReadableResponse
func (client *Client) CheckReadableWithOptions(request *CheckReadableRequest, runtime *dara.RuntimeOptions) (_result *CheckReadableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckReadable"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckReadableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a user has permissions to view data works, such as dashboards and workbooks.
//
// @param request - CheckReadableRequest
//
// @return CheckReadableResponse
func (client *Client) CheckReadable(request *CheckReadableRequest) (_result *CheckReadableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckReadableResponse{}
	_body, _err := client.CheckReadableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generate a ticket for third-party embedding.
//
// Description:
//
// For detailed usage, please refer to [Report Embedding Data Permission Control and Parameter Passing Security Enhancement Solution](https://help.aliyun.com/document_detail/391291.html).
//
// @param request - CreateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.CmptId) {
		query["CmptId"] = request.CmptId
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.GlobalParam) {
		query["GlobalParam"] = request.GlobalParam
	}

	if !dara.IsNil(request.TicketNum) {
		query["TicketNum"] = request.TicketNum
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WatermarkParam) {
		query["WatermarkParam"] = request.WatermarkParam
	}

	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTicket"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generate a ticket for third-party embedding.
//
// Description:
//
// For detailed usage, please refer to [Report Embedding Data Permission Control and Parameter Passing Security Enhancement Solution](https://help.aliyun.com/document_detail/391291.html).
//
// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generate an embedding ticket for Smart Q.
//
// @param request - CreateTicket4CopilotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicket4CopilotResponse
func (client *Client) CreateTicket4CopilotWithOptions(request *CreateTicket4CopilotRequest, runtime *dara.RuntimeOptions) (_result *CreateTicket4CopilotResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.CopilotId) {
		query["CopilotId"] = request.CopilotId
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.TicketNum) {
		query["TicketNum"] = request.TicketNum
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTicket4Copilot"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTicket4CopilotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generate an embedding ticket for Smart Q.
//
// @param request - CreateTicket4CopilotRequest
//
// @return CreateTicket4CopilotResponse
func (client *Client) CreateTicket4Copilot(request *CreateTicket4CopilotRequest) (_result *CreateTicket4CopilotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTicket4CopilotResponse{}
	_body, _err := client.CreateTicket4CopilotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a user group. You can specify a parent user group.
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
	if !dara.IsNil(request.ParentUserGroupId) {
		query["ParentUserGroupId"] = request.ParentUserGroupId
	}

	if !dara.IsNil(request.UserGroupDescription) {
		query["UserGroupDescription"] = request.UserGroupDescription
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
		Action:      dara.String("CreateUserGroup"),
		Version:     dara.String("2022-01-01"),
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
// Create a user group. You can specify a parent user group.
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
// 数据解读开放接口
//
// @param request - DataInterpretationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataInterpretationResponse
func (client *Client) DataInterpretationWithOptions(request *DataInterpretationRequest, runtime *dara.RuntimeOptions) (_result *DataInterpretationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.ModelCode) {
		query["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.PromptForceOverride) {
		query["PromptForceOverride"] = request.PromptForceOverride
	}

	if !dara.IsNil(request.UserPrompt) {
		query["UserPrompt"] = request.UserPrompt
	}

	if !dara.IsNil(request.UserQuestion) {
		query["UserQuestion"] = request.UserQuestion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DataInterpretation"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DataInterpretationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据解读开放接口
//
// @param request - DataInterpretationRequest
//
// @return DataInterpretationResponse
func (client *Client) DataInterpretation(request *DataInterpretationRequest) (_result *DataInterpretationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DataInterpretationResponse{}
	_body, _err := client.DataInterpretationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query works information under the specified dataset.
//
// @param request - DataSetBloodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataSetBloodResponse
func (client *Client) DataSetBloodWithOptions(request *DataSetBloodRequest, runtime *dara.RuntimeOptions) (_result *DataSetBloodResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSetIds) {
		query["DataSetIds"] = request.DataSetIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorksType) {
		query["WorksType"] = request.WorksType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DataSetBlood"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DataSetBloodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query works information under the specified dataset.
//
// @param request - DataSetBloodRequest
//
// @return DataSetBloodResponse
func (client *Client) DataSetBlood(request *DataSetBloodRequest) (_result *DataSetBloodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DataSetBloodResponse{}
	_body, _err := client.DataSetBloodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query dataset information under the specified data source
//
// @param request - DataSourceBloodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataSourceBloodResponse
func (client *Client) DataSourceBloodWithOptions(request *DataSourceBloodRequest, runtime *dara.RuntimeOptions) (_result *DataSourceBloodResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DataSourceBlood"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DataSourceBloodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query dataset information under the specified data source
//
// @param request - DataSourceBloodRequest
//
// @return DataSourceBloodResponse
func (client *Client) DataSourceBlood(request *DataSourceBloodRequest) (_result *DataSourceBloodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DataSourceBloodResponse{}
	_body, _err := client.DataSourceBloodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the expiration time of the ticket embedded in the report.
//
// @param request - DelayTicketExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelayTicketExpireTimeResponse
func (client *Client) DelayTicketExpireTimeWithOptions(request *DelayTicketExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *DelayTicketExpireTimeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.Ticket) {
		query["Ticket"] = request.Ticket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelayTicketExpireTime"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelayTicketExpireTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the expiration time of the ticket embedded in the report.
//
// @param request - DelayTicketExpireTimeRequest
//
// @return DelayTicketExpireTimeResponse
func (client *Client) DelayTicketExpireTime(request *DelayTicketExpireTimeRequest) (_result *DelayTicketExpireTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelayTicketExpireTimeResponse{}
	_body, _err := client.DelayTicketExpireTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// { "ruleId": "a5bb24da- ***-a891683e14da", // The ID of the row-column permission rule. "cubeId": "7c7223ae- ***-3c744528014b", // The ID of the dataset. "delModel": { "userGroups": [ "0d5fb19b- ***-1248 fc27ca51", // Delete the user group ID of the user group. "3d2c23d4-***-f6390f325c2d" ], "users": [ "4334 ***358", // Delete the UserID of the user group. "Huang***3fa822" ] } }
//
// Description:
//
// {"ruleId":"a5bb24da-***-a891683e14da","cubeId":"7c7223ae-***-3c744528014b","delModel":{"userGroups":["0d5fb19b-***-1248fc27ca51","3d2c23d4-***-f6390f325c2d"],"users":["4334***358","Huang***3fa822"]}}
//
// @param request - DeleteDataLevelPermissionRuleUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLevelPermissionRuleUsersResponse
func (client *Client) DeleteDataLevelPermissionRuleUsersWithOptions(request *DeleteDataLevelPermissionRuleUsersRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLevelPermissionRuleUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteUserModel) {
		query["DeleteUserModel"] = request.DeleteUserModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLevelPermissionRuleUsers"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLevelPermissionRuleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// { "ruleId": "a5bb24da- ***-a891683e14da", // The ID of the row-column permission rule. "cubeId": "7c7223ae- ***-3c744528014b", // The ID of the dataset. "delModel": { "userGroups": [ "0d5fb19b- ***-1248 fc27ca51", // Delete the user group ID of the user group. "3d2c23d4-***-f6390f325c2d" ], "users": [ "4334 ***358", // Delete the UserID of the user group. "Huang***3fa822" ] } }
//
// Description:
//
// {"ruleId":"a5bb24da-***-a891683e14da","cubeId":"7c7223ae-***-3c744528014b","delModel":{"userGroups":["0d5fb19b-***-1248fc27ca51","3d2c23d4-***-f6390f325c2d"],"users":["4334***358","Huang***3fa822"]}}
//
// @param request - DeleteDataLevelPermissionRuleUsersRequest
//
// @return DeleteDataLevelPermissionRuleUsersResponse
func (client *Client) DeleteDataLevelPermissionRuleUsers(request *DeleteDataLevelPermissionRuleUsersRequest) (_result *DeleteDataLevelPermissionRuleUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLevelPermissionRuleUsersResponse{}
	_body, _err := client.DeleteDataLevelPermissionRuleUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// The ID of the training dataset that you want to remove from the specified custom linguistic model.
//
// @param request - DeleteDataLevelRuleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLevelRuleConfigResponse
func (client *Client) DeleteDataLevelRuleConfigWithOptions(request *DeleteDataLevelRuleConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLevelRuleConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLevelRuleConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLevelRuleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// The ID of the training dataset that you want to remove from the specified custom linguistic model.
//
// @param request - DeleteDataLevelRuleConfigRequest
//
// @return DeleteDataLevelRuleConfigResponse
func (client *Client) DeleteDataLevelRuleConfig(request *DeleteDataLevelRuleConfigRequest) (_result *DeleteDataLevelRuleConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLevelRuleConfigResponse{}
	_body, _err := client.DeleteDataLevelRuleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Third-Party Embedded Ticket
//
// @param request - DeleteTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTicketResponse
func (client *Client) DeleteTicketWithOptions(request *DeleteTicketRequest, runtime *dara.RuntimeOptions) (_result *DeleteTicketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ticket) {
		query["Ticket"] = request.Ticket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTicket"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Third-Party Embedded Ticket
//
// @param request - DeleteTicketRequest
//
// @return DeleteTicketResponse
func (client *Client) DeleteTicket(request *DeleteTicketRequest) (_result *DeleteTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTicketResponse{}
	_body, _err := client.DeleteTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified organization user.
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
	if !dara.IsNil(request.TransferUserId) {
		query["TransferUserId"] = request.TransferUserId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2022-01-01"),
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
// Delete the specified organization user.
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
// Delete a member from the specified workspace.
//
// @param request - DeleteUserFromWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserFromWorkspaceResponse
func (client *Client) DeleteUserFromWorkspaceWithOptions(request *DeleteUserFromWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserFromWorkspaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserFromWorkspace"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserFromWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a member from the specified workspace.
//
// @param request - DeleteUserFromWorkspaceRequest
//
// @return DeleteUserFromWorkspaceResponse
func (client *Client) DeleteUserFromWorkspace(request *DeleteUserFromWorkspaceRequest) (_result *DeleteUserFromWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserFromWorkspaceResponse{}
	_body, _err := client.DeleteUserFromWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user group in an organization.
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
	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroup"),
		Version:     dara.String("2022-01-01"),
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
// Deletes a user group in an organization.
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
// Deletes a specified member from a specified user group.
//
// @param request - DeleteUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupMemberResponse
func (client *Client) DeleteUserGroupMemberWithOptions(request *DeleteUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupMemberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroupMember"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified member from a specified user group.
//
// @param request - DeleteUserGroupMemberRequest
//
// @return DeleteUserGroupMemberResponse
func (client *Client) DeleteUserGroupMember(request *DeleteUserGroupMemberRequest) (_result *DeleteUserGroupMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserGroupMemberResponse{}
	_body, _err := client.DeleteUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch remove specified users from user groups.
//
// @param request - DeleteUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupMembersResponse
func (client *Client) DeleteUserGroupMembersWithOptions(request *DeleteUserGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupMembersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroupMembers"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserGroupMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch remove specified users from user groups.
//
// @param request - DeleteUserGroupMembersRequest
//
// @return DeleteUserGroupMembersResponse
func (client *Client) DeleteUserGroupMembers(request *DeleteUserGroupMembersRequest) (_result *DeleteUserGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserGroupMembersResponse{}
	_body, _err := client.DeleteUserGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the tag metadata of an organization member.
//
// @param request - DeleteUserTagMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserTagMetaResponse
func (client *Client) DeleteUserTagMetaWithOptions(request *DeleteUserTagMetaRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserTagMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserTagMeta"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserTagMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the tag metadata of an organization member.
//
// @param request - DeleteUserTagMetaRequest
//
// @return DeleteUserTagMetaResponse
func (client *Client) DeleteUserTagMeta(request *DeleteUserTagMetaRequest) (_result *DeleteUserTagMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserTagMetaResponse{}
	_body, _err := client.DeleteUserTagMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Data Source Information
//
// @param request - GetDataSourceConnectionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceConnectionInfoResponse
func (client *Client) GetDataSourceConnectionInfoWithOptions(request *GetDataSourceConnectionInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceConnectionInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DsId) {
		query["DsId"] = request.DsId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataSourceConnectionInfo"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceConnectionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Data Source Information
//
// @param request - GetDataSourceConnectionInfoRequest
//
// @return GetDataSourceConnectionInfoResponse
func (client *Client) GetDataSourceConnectionInfo(request *GetDataSourceConnectionInfoRequest) (_result *GetDataSourceConnectionInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataSourceConnectionInfoResponse{}
	_body, _err := client.GetDataSourceConnectionInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check the running status of mail tasks within an organization
//
// @param request - GetMailTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMailTaskStatusResponse
func (client *Client) GetMailTaskStatusWithOptions(request *GetMailTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetMailTaskStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MailId) {
		query["MailId"] = request.MailId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMailTaskStatus"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMailTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check the running status of mail tasks within an organization
//
// @param request - GetMailTaskStatusRequest
//
// @return GetMailTaskStatusResponse
func (client *Client) GetMailTaskStatus(request *GetMailTaskStatusRequest) (_result *GetMailTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMailTaskStatusResponse{}
	_body, _err := client.GetMailTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Search for user group information based on the keyword of the user group name.
//
// @param request - GetUserGroupInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupInfoResponse
func (client *Client) GetUserGroupInfoWithOptions(request *GetUserGroupInfoRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserGroupInfo"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserGroupInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search for user group information based on the keyword of the user group name.
//
// @param request - GetUserGroupInfoRequest
//
// @return GetUserGroupInfoResponse
func (client *Client) GetUserGroupInfo(request *GetUserGroupInfoRequest) (_result *GetUserGroupInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserGroupInfoResponse{}
	_body, _err := client.GetUserGroupInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of embedded reports
//
// @param request - GetWorksEmbedListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorksEmbedListResponse
func (client *Client) GetWorksEmbedListWithOptions(request *GetWorksEmbedListRequest, runtime *dara.RuntimeOptions) (_result *GetWorksEmbedListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorksType) {
		query["WorksType"] = request.WorksType
	}

	if !dara.IsNil(request.WsId) {
		query["WsId"] = request.WsId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorksEmbedList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorksEmbedListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of embedded reports
//
// @param request - GetWorksEmbedListRequest
//
// @return GetWorksEmbedListResponse
func (client *Client) GetWorksEmbedList(request *GetWorksEmbedListRequest) (_result *GetWorksEmbedListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorksEmbedListResponse{}
	_body, _err := client.GetWorksEmbedListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries API data sources.
//
// Description:
//
// For more information about the parameters, see [Create an API data source](https://help.aliyun.com/document_detail/409330.html).
//
// @param request - ListApiDatasourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiDatasourceResponse
func (client *Client) ListApiDatasourceWithOptions(request *ListApiDatasourceRequest, runtime *dara.RuntimeOptions) (_result *ListApiDatasourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiDatasource"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiDatasourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries API data sources.
//
// Description:
//
// For more information about the parameters, see [Create an API data source](https://help.aliyun.com/document_detail/409330.html).
//
// @param request - ListApiDatasourceRequest
//
// @return ListApiDatasourceResponse
func (client *Client) ListApiDatasource(request *ListApiDatasourceRequest) (_result *ListApiDatasourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApiDatasourceResponse{}
	_body, _err := client.ListApiDatasourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user group information at a time by user group ID.
//
// @param request - ListByUserGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListByUserGroupIdResponse
func (client *Client) ListByUserGroupIdWithOptions(request *ListByUserGroupIdRequest, runtime *dara.RuntimeOptions) (_result *ListByUserGroupIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		query["UserGroupIds"] = request.UserGroupIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListByUserGroupId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListByUserGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user group information at a time by user group ID.
//
// @param request - ListByUserGroupIdRequest
//
// @return ListByUserGroupIdResponse
func (client *Client) ListByUserGroupId(request *ListByUserGroupIdRequest) (_result *ListByUserGroupIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListByUserGroupIdResponse{}
	_body, _err := client.ListByUserGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the work.
//
// @param request - ListCollectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCollectionsResponse
func (client *Client) ListCollectionsWithOptions(request *ListCollectionsRequest, runtime *dara.RuntimeOptions) (_result *ListCollectionsResponse, _err error) {
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
		Action:      dara.String("ListCollections"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the work.
//
// @param request - ListCollectionsRequest
//
// @return ListCollectionsResponse
func (client *Client) ListCollections(request *ListCollectionsRequest) (_result *ListCollectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCollectionsResponse{}
	_body, _err := client.ListCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can this operation to obtain a list of row and column permission configurations for a specified dataset.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.
//
// @param request - ListCubeDataLevelPermissionConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCubeDataLevelPermissionConfigResponse
func (client *Client) ListCubeDataLevelPermissionConfigWithOptions(request *ListCubeDataLevelPermissionConfigRequest, runtime *dara.RuntimeOptions) (_result *ListCubeDataLevelPermissionConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCubeDataLevelPermissionConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCubeDataLevelPermissionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can this operation to obtain a list of row and column permission configurations for a specified dataset.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.
//
// @param request - ListCubeDataLevelPermissionConfigRequest
//
// @return ListCubeDataLevelPermissionConfigResponse
func (client *Client) ListCubeDataLevelPermissionConfig(request *ListCubeDataLevelPermissionConfigRequest) (_result *ListCubeDataLevelPermissionConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCubeDataLevelPermissionConfigResponse{}
	_body, _err := client.ListCubeDataLevelPermissionConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the whitelist for dataset row and column permissions based on the type of permission.
//
// Description:
//
// > This API only supports the new row and column permission model of Quick BI. If you are still using the old row and column permissions, please migrate to the new row and column permission model before calling this interface. To migrate to the new row and column permission model, follow these steps: In Organization Management -> Security Configuration -> Upgrade Row and Column Permissions, click **One-Click Upgrade*	- to upgrade to the new row-level permissions.
//
// @param request - ListDataLevelPermissionWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLevelPermissionWhiteListResponse
func (client *Client) ListDataLevelPermissionWhiteListWithOptions(request *ListDataLevelPermissionWhiteListRequest, runtime *dara.RuntimeOptions) (_result *ListDataLevelPermissionWhiteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLevelPermissionWhiteList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLevelPermissionWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the whitelist for dataset row and column permissions based on the type of permission.
//
// Description:
//
// > This API only supports the new row and column permission model of Quick BI. If you are still using the old row and column permissions, please migrate to the new row and column permission model before calling this interface. To migrate to the new row and column permission model, follow these steps: In Organization Management -> Security Configuration -> Upgrade Row and Column Permissions, click **One-Click Upgrade*	- to upgrade to the new row-level permissions.
//
// @param request - ListDataLevelPermissionWhiteListRequest
//
// @return ListDataLevelPermissionWhiteListResponse
func (client *Client) ListDataLevelPermissionWhiteList(request *ListDataLevelPermissionWhiteListRequest) (_result *ListDataLevelPermissionWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLevelPermissionWhiteListResponse{}
	_body, _err := client.ListDataLevelPermissionWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query all data sources under the specified space
//
// @param request - ListDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceResponse
func (client *Client) ListDataSourceWithOptions(request *ListDataSourceRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DsType) {
		query["DsType"] = request.DsType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSource"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query all data sources under the specified space
//
// @param request - ListDataSourceRequest
//
// @return ListDataSourceResponse
func (client *Client) ListDataSource(request *ListDataSourceRequest) (_result *ListDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSourceResponse{}
	_body, _err := client.ListDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Overview
//
// @param request - ListFavoriteReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFavoriteReportsResponse
func (client *Client) ListFavoriteReportsWithOptions(request *ListFavoriteReportsRequest, runtime *dara.RuntimeOptions) (_result *ListFavoriteReportsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TreeType) {
		query["TreeType"] = request.TreeType
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFavoriteReports"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFavoriteReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Overview
//
// @param request - ListFavoriteReportsRequest
//
// @return ListFavoriteReportsResponse
func (client *Client) ListFavoriteReports(request *ListFavoriteReportsRequest) (_result *ListFavoriteReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFavoriteReportsResponse{}
	_body, _err := client.ListFavoriteReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get user list under the specified organization role.
//
// @param request - ListOrganizationRoleUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationRoleUsersResponse
func (client *Client) ListOrganizationRoleUsersWithOptions(request *ListOrganizationRoleUsersRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationRoleUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationRoleUsers"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationRoleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get user list under the specified organization role.
//
// @param request - ListOrganizationRoleUsersRequest
//
// @return ListOrganizationRoleUsersResponse
func (client *Client) ListOrganizationRoleUsers(request *ListOrganizationRoleUsersRequest) (_result *ListOrganizationRoleUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOrganizationRoleUsersResponse{}
	_body, _err := client.ListOrganizationRoleUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the list of custom roles at the organization level.
//
// @param request - ListOrganizationRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationRolesResponse
func (client *Client) ListOrganizationRolesWithOptions(runtime *dara.RuntimeOptions) (_result *ListOrganizationRolesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationRoles"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of custom roles at the organization level.
//
// @return ListOrganizationRolesResponse
func (client *Client) ListOrganizationRoles() (_result *ListOrganizationRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOrganizationRolesResponse{}
	_body, _err := client.ListOrganizationRolesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list of authorization details for a BI portal menu.
//
// @param request - ListPortalMenuAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPortalMenuAuthorizationResponse
func (client *Client) ListPortalMenuAuthorizationWithOptions(request *ListPortalMenuAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *ListPortalMenuAuthorizationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataPortalId) {
		query["DataPortalId"] = request.DataPortalId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPortalMenuAuthorization"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPortalMenuAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of authorization details for a BI portal menu.
//
// @param request - ListPortalMenuAuthorizationRequest
//
// @return ListPortalMenuAuthorizationResponse
func (client *Client) ListPortalMenuAuthorization(request *ListPortalMenuAuthorizationRequest) (_result *ListPortalMenuAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPortalMenuAuthorizationResponse{}
	_body, _err := client.ListPortalMenuAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets a hierarchical list of menus under a specific BI portal.
//
// @param request - ListPortalMenusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPortalMenusResponse
func (client *Client) ListPortalMenusWithOptions(request *ListPortalMenusRequest, runtime *dara.RuntimeOptions) (_result *ListPortalMenusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataPortalId) {
		query["DataPortalId"] = request.DataPortalId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPortalMenus"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPortalMenusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a hierarchical list of menus under a specific BI portal.
//
// @param request - ListPortalMenusRequest
//
// @return ListPortalMenusResponse
func (client *Client) ListPortalMenus(request *ListPortalMenusRequest) (_result *ListPortalMenusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPortalMenusResponse{}
	_body, _err := client.ListPortalMenusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to obtain a list of the most frequently viewed and footsteps displayed in the homepage dashboard for a specified user.
//
// @param request - ListRecentViewReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentViewReportsResponse
func (client *Client) ListRecentViewReportsWithOptions(request *ListRecentViewReportsRequest, runtime *dara.RuntimeOptions) (_result *ListRecentViewReportsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OffsetDay) {
		query["OffsetDay"] = request.OffsetDay
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryMode) {
		query["QueryMode"] = request.QueryMode
	}

	if !dara.IsNil(request.TreeType) {
		query["TreeType"] = request.TreeType
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecentViewReports"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecentViewReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to obtain a list of the most frequently viewed and footsteps displayed in the homepage dashboard for a specified user.
//
// @param request - ListRecentViewReportsRequest
//
// @return ListRecentViewReportsResponse
func (client *Client) ListRecentViewReports(request *ListRecentViewReportsRequest) (_result *ListRecentViewReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecentViewReportsResponse{}
	_body, _err := client.ListRecentViewReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can this operation to obtain the list of authorized works displayed on the homepage of a specified user.
//
// @param request - ListSharedReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSharedReportsResponse
func (client *Client) ListSharedReportsWithOptions(request *ListSharedReportsRequest, runtime *dara.RuntimeOptions) (_result *ListSharedReportsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TreeType) {
		query["TreeType"] = request.TreeType
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSharedReports"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSharedReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can this operation to obtain the list of authorized works displayed on the homepage of a specified user.
//
// @param request - ListSharedReportsRequest
//
// @return ListSharedReportsResponse
func (client *Client) ListSharedReports(request *ListSharedReportsRequest) (_result *ListSharedReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSharedReportsResponse{}
	_body, _err := client.ListSharedReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all user groups to which a user belongs based on the user ID.
//
// @param request - ListUserGroupsByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsByUserIdResponse
func (client *Client) ListUserGroupsByUserIdWithOptions(request *ListUserGroupsByUserIdRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsByUserIdResponse, _err error) {
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
		Action:      dara.String("ListUserGroupsByUserId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all user groups to which a user belongs based on the user ID.
//
// @param request - ListUserGroupsByUserIdRequest
//
// @return ListUserGroupsByUserIdResponse
func (client *Client) ListUserGroupsByUserId(request *ListUserGroupsByUserIdRequest) (_result *ListUserGroupsByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsByUserIdResponse{}
	_body, _err := client.ListUserGroupsByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get user list under the specified workspace role.
//
// @param request - ListWorkspaceRoleUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceRoleUsersResponse
func (client *Client) ListWorkspaceRoleUsersWithOptions(request *ListWorkspaceRoleUsersRequest, runtime *dara.RuntimeOptions) (_result *ListWorkspaceRoleUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaceRoleUsers"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspaceRoleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get user list under the specified workspace role.
//
// @param request - ListWorkspaceRoleUsersRequest
//
// @return ListWorkspaceRoleUsersResponse
func (client *Client) ListWorkspaceRoleUsers(request *ListWorkspaceRoleUsersRequest) (_result *ListWorkspaceRoleUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkspaceRoleUsersResponse{}
	_body, _err := client.ListWorkspaceRoleUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the list of workspace roles.
//
// @param request - ListWorkspaceRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceRolesResponse
func (client *Client) ListWorkspaceRolesWithOptions(request *ListWorkspaceRolesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkspaceRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaceRoles"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspaceRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the list of workspace roles.
//
// @param request - ListWorkspaceRolesRequest
//
// @return ListWorkspaceRolesResponse
func (client *Client) ListWorkspaceRoles(request *ListWorkspaceRolesRequest) (_result *ListWorkspaceRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkspaceRolesResponse{}
	_body, _err := client.ListWorkspaceRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Manually Execute Email Task
//
// @param request - ManualRunMailTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualRunMailTaskResponse
func (client *Client) ManualRunMailTaskWithOptions(request *ManualRunMailTaskRequest, runtime *dara.RuntimeOptions) (_result *ManualRunMailTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MailId) {
		query["MailId"] = request.MailId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManualRunMailTask"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManualRunMailTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Manually Execute Email Task
//
// @param request - ManualRunMailTaskRequest
//
// @return ManualRunMailTaskResponse
func (client *Client) ManualRunMailTask(request *ManualRunMailTaskRequest) (_result *ManualRunMailTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ManualRunMailTaskResponse{}
	_body, _err := client.ManualRunMailTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a specified API data source.
//
// Description:
//
// When you modify a query statement, you can modify only the top-level JsonObject. You cannot modify parameters that are nested in multiple layers. For more information about the parameters, see [Create an API data source](https://help.aliyun.com/document_detail/409330.html).
//
// @param request - ModifyApiDatasourceParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiDatasourceParametersResponse
func (client *Client) ModifyApiDatasourceParametersWithOptions(request *ModifyApiDatasourceParametersRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiDatasourceParametersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiDatasourceParameters"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiDatasourceParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a specified API data source.
//
// Description:
//
// When you modify a query statement, you can modify only the top-level JsonObject. You cannot modify parameters that are nested in multiple layers. For more information about the parameters, see [Create an API data source](https://help.aliyun.com/document_detail/409330.html).
//
// @param request - ModifyApiDatasourceParametersRequest
//
// @return ModifyApiDatasourceParametersResponse
func (client *Client) ModifyApiDatasourceParameters(request *ModifyApiDatasourceParametersRequest) (_result *ModifyApiDatasourceParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApiDatasourceParametersResponse{}
	_body, _err := client.ModifyApiDatasourceParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Intelligent Query Embedding Configuration
//
// @param request - ModifyCopilotEmbedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCopilotEmbedConfigResponse
func (client *Client) ModifyCopilotEmbedConfigWithOptions(request *ModifyCopilotEmbedConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyCopilotEmbedConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.CopilotId) {
		query["CopilotId"] = request.CopilotId
	}

	if !dara.IsNil(request.DataRange) {
		query["DataRange"] = request.DataRange
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCopilotEmbedConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCopilotEmbedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Intelligent Query Embedding Configuration
//
// @param request - ModifyCopilotEmbedConfigRequest
//
// @return ModifyCopilotEmbedConfigResponse
func (client *Client) ModifyCopilotEmbedConfig(request *ModifyCopilotEmbedConfigRequest) (_result *ModifyCopilotEmbedConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCopilotEmbedConfigResponse{}
	_body, _err := client.ModifyCopilotEmbedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get approval flow information based on the approver.
//
// @param request - QueryApprovalInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApprovalInfoResponse
func (client *Client) QueryApprovalInfoWithOptions(request *QueryApprovalInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryApprovalInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryApprovalInfo"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryApprovalInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get approval flow information based on the approver.
//
// @param request - QueryApprovalInfoRequest
//
// @return QueryApprovalInfoResponse
func (client *Client) QueryApprovalInfo(request *QueryApprovalInfoRequest) (_result *QueryApprovalInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryApprovalInfoResponse{}
	_body, _err := client.QueryApprovalInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query audit log information.
//
// @param request - QueryAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAuditLogResponse
func (client *Client) QueryAuditLogWithOptions(request *QueryAuditLogRequest, runtime *dara.RuntimeOptions) (_result *QueryAuditLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessSourceFlag) {
		query["AccessSourceFlag"] = request.AccessSourceFlag
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.OperatorId) {
		query["OperatorId"] = request.OperatorId
	}

	if !dara.IsNil(request.OperatorTypes) {
		query["OperatorTypes"] = request.OperatorTypes
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserAccessDevice) {
		query["UserAccessDevice"] = request.UserAccessDevice
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAuditLog"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query audit log information.
//
// @param request - QueryAuditLogRequest
//
// @return QueryAuditLogResponse
func (client *Client) QueryAuditLog(request *QueryAuditLogRequest) (_result *QueryAuditLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAuditLogResponse{}
	_body, _err := client.QueryAuditLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries component performance logs.
//
// @param request - QueryComponentPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryComponentPerformanceResponse
func (client *Client) QueryComponentPerformanceWithOptions(request *QueryComponentPerformanceRequest, runtime *dara.RuntimeOptions) (_result *QueryComponentPerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CostTimeAvgMin) {
		query["CostTimeAvgMin"] = request.CostTimeAvgMin
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryComponentPerformance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryComponentPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries component performance logs.
//
// @param request - QueryComponentPerformanceRequest
//
// @return QueryComponentPerformanceResponse
func (client *Client) QueryComponentPerformance(request *QueryComponentPerformanceRequest) (_result *QueryComponentPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryComponentPerformanceResponse{}
	_body, _err := client.QueryComponentPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get the List of Configurations for Activating XiaoQ Embedding
//
// @param request - QueryCopilotEmbedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCopilotEmbedConfigResponse
func (client *Client) QueryCopilotEmbedConfigWithOptions(request *QueryCopilotEmbedConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryCopilotEmbedConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCopilotEmbedConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCopilotEmbedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the List of Configurations for Activating XiaoQ Embedding
//
// @param request - QueryCopilotEmbedConfigRequest
//
// @return QueryCopilotEmbedConfigResponse
func (client *Client) QueryCopilotEmbedConfig(request *QueryCopilotEmbedConfigRequest) (_result *QueryCopilotEmbedConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCopilotEmbedConfigResponse{}
	_body, _err := client.QueryCopilotEmbedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries dataset optimization suggestions.
//
// @param request - QueryCubeOptimizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCubeOptimizationResponse
func (client *Client) QueryCubeOptimizationWithOptions(request *QueryCubeOptimizationRequest, runtime *dara.RuntimeOptions) (_result *QueryCubeOptimizationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCubeOptimization"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCubeOptimizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries dataset optimization suggestions.
//
// @param request - QueryCubeOptimizationRequest
//
// @return QueryCubeOptimizationResponse
func (client *Client) QueryCubeOptimization(request *QueryCubeOptimizationRequest) (_result *QueryCubeOptimizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCubeOptimizationResponse{}
	_body, _err := client.QueryCubeOptimizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the performance logs of a dataset.
//
// @param request - QueryCubePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCubePerformanceResponse
func (client *Client) QueryCubePerformanceWithOptions(request *QueryCubePerformanceRequest, runtime *dara.RuntimeOptions) (_result *QueryCubePerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CostTimeAvgMin) {
		query["CostTimeAvgMin"] = request.CostTimeAvgMin
	}

	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCubePerformance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCubePerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance logs of a dataset.
//
// @param request - QueryCubePerformanceRequest
//
// @return QueryCubePerformanceResponse
func (client *Client) QueryCubePerformance(request *QueryCubePerformanceRequest) (_result *QueryCubePerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCubePerformanceResponse{}
	_body, _err := client.QueryCubePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invoke the open data service API.
//
// Description:
//
// ### Prerequisites
//
// You need to create a data service API through Quick BI\\"s data service. For more details, see: [Data Service](https://help.aliyun.com/document_detail/144980.html).
//
// ### Usage Restrictions
//
//   - The data service feature is only available to professional edition customers.
//
//   - The timeout for data service API calls is 60s, and the QPS for a single API is 10 times/second.
//
//   - If row-level permissions are enabled on the dataset referenced by the data service API, the API call will also be intercepted by the row-level permission policy.
//
// @param request - QueryDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataResponse
func (client *Client) QueryDataWithOptions(request *QueryDataRequest, runtime *dara.RuntimeOptions) (_result *QueryDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.Conditions) {
		query["Conditions"] = request.Conditions
	}

	if !dara.IsNil(request.ReturnFields) {
		query["ReturnFields"] = request.ReturnFields
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryData"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke the open data service API.
//
// Description:
//
// ### Prerequisites
//
// You need to create a data service API through Quick BI\\"s data service. For more details, see: [Data Service](https://help.aliyun.com/document_detail/144980.html).
//
// ### Usage Restrictions
//
//   - The data service feature is only available to professional edition customers.
//
//   - The timeout for data service API calls is 60s, and the QPS for a single API is 10 times/second.
//
//   - If row-level permissions are enabled on the dataset referenced by the data service API, the API call will also be intercepted by the row-level permission policy.
//
// @param request - QueryDataRequest
//
// @return QueryDataResponse
func (client *Client) QueryData(request *QueryDataRequest) (_result *QueryDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDataResponse{}
	_body, _err := client.QueryDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Data Range Catalog List
//
// @param request - QueryDataRangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataRangeResponse
func (client *Client) QueryDataRangeWithOptions(request *QueryDataRangeRequest, runtime *dara.RuntimeOptions) (_result *QueryDataRangeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDataRange"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDataRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Data Range Catalog List
//
// @param request - QueryDataRangeRequest
//
// @return QueryDataRangeResponse
func (client *Client) QueryDataRange(request *QueryDataRangeRequest) (_result *QueryDataRangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDataRangeResponse{}
	_body, _err := client.QueryDataRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI QueryDataService is deprecated, please use quickbi-public::2022-01-01::QueryData instead.
//
// Summary:
//
// Invoke an already created API in the data service.
//
// Description:
//
// #### Prerequisites
//
// You create the data service API through Quick BI\\"s data service. For more details, see [Data Service](https://help.aliyun.com/document_detail/144980.html).
//
// #### Usage Restrictions
//
//   - The data service feature is only available to professional edition customers.
//
//   - The timeout for data service API calls is 60s, and the QPS for a single API is 10 times/second.
//
//   - If row-level permissions are enabled on the dataset referenced by the data service API, the API call may be intercepted by the row-level permission policy.
//
// @param request - QueryDataServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataServiceResponse
// Deprecated
func (client *Client) QueryDataServiceWithOptions(request *QueryDataServiceRequest, runtime *dara.RuntimeOptions) (_result *QueryDataServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.Conditions) {
		query["Conditions"] = request.Conditions
	}

	if !dara.IsNil(request.ReturnFields) {
		query["ReturnFields"] = request.ReturnFields
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDataService"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDataServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI QueryDataService is deprecated, please use quickbi-public::2022-01-01::QueryData instead.
//
// Summary:
//
// Invoke an already created API in the data service.
//
// Description:
//
// #### Prerequisites
//
// You create the data service API through Quick BI\\"s data service. For more details, see [Data Service](https://help.aliyun.com/document_detail/144980.html).
//
// #### Usage Restrictions
//
//   - The data service feature is only available to professional edition customers.
//
//   - The timeout for data service API calls is 60s, and the QPS for a single API is 10 times/second.
//
//   - If row-level permissions are enabled on the dataset referenced by the data service API, the API call may be intercepted by the row-level permission policy.
//
// @param request - QueryDataServiceRequest
//
// @return QueryDataServiceResponse
// Deprecated
func (client *Client) QueryDataService(request *QueryDataServiceRequest) (_result *QueryDataServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDataServiceResponse{}
	_body, _err := client.QueryDataServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Data Service API List
//
// @param request - QueryDataServiceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataServiceListResponse
func (client *Client) QueryDataServiceListWithOptions(request *QueryDataServiceListRequest, runtime *dara.RuntimeOptions) (_result *QueryDataServiceListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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
		Action:      dara.String("QueryDataServiceList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDataServiceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Data Service API List
//
// @param request - QueryDataServiceListRequest
//
// @return QueryDataServiceListResponse
func (client *Client) QueryDataServiceList(request *QueryDataServiceListRequest) (_result *QueryDataServiceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDataServiceListResponse{}
	_body, _err := client.QueryDataServiceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified dataset, including the data source, directory, and dataset model.
//
// Description:
//
// The data source, directory, and dataset model (including dimensions, measures, physical fields, custom SQL text, and association relationships).
//
// @param request - QueryDatasetDetailInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetDetailInfoResponse
func (client *Client) QueryDatasetDetailInfoWithOptions(request *QueryDatasetDetailInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetDetailInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDatasetDetailInfo"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDatasetDetailInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified dataset, including the data source, directory, and dataset model.
//
// Description:
//
// The data source, directory, and dataset model (including dimensions, measures, physical fields, custom SQL text, and association relationships).
//
// @param request - QueryDatasetDetailInfoRequest
//
// @return QueryDatasetDetailInfoResponse
func (client *Client) QueryDatasetDetailInfo(request *QueryDatasetDetailInfoRequest) (_result *QueryDatasetDetailInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDatasetDetailInfoResponse{}
	_body, _err := client.QueryDatasetDetailInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Indicates whether the table is a custom SQL table. Valid values:
//
// \\	- true: custom SQL table
//
// \\	- false: non-custom SQL table
//
// @param request - QueryDatasetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetInfoResponse
func (client *Client) QueryDatasetInfoWithOptions(request *QueryDatasetInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDatasetInfo"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDatasetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Indicates whether the table is a custom SQL table. Valid values:
//
// \\	- true: custom SQL table
//
// \\	- false: non-custom SQL table
//
// @param request - QueryDatasetInfoRequest
//
// @return QueryDatasetInfoResponse
func (client *Client) QueryDatasetInfo(request *QueryDatasetInfoRequest) (_result *QueryDatasetInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDatasetInfoResponse{}
	_body, _err := client.QueryDatasetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The name of the training dataset.
//
// @param request - QueryDatasetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetListResponse
func (client *Client) QueryDatasetListWithOptions(request *QueryDatasetListRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WithChildren) {
		query["WithChildren"] = request.WithChildren
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDatasetList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDatasetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The name of the training dataset.
//
// @param request - QueryDatasetListRequest
//
// @return QueryDatasetListResponse
func (client *Client) QueryDatasetList(request *QueryDatasetListRequest) (_result *QueryDatasetListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDatasetListResponse{}
	_body, _err := client.QueryDatasetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check if the Dataset has Enabled Smart Query
//
// @param request - QueryDatasetSmartqStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetSmartqStatusResponse
func (client *Client) QueryDatasetSmartqStatusWithOptions(request *QueryDatasetSmartqStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetSmartqStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDatasetSmartqStatus"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDatasetSmartqStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check if the Dataset has Enabled Smart Query
//
// @param request - QueryDatasetSmartqStatusRequest
//
// @return QueryDatasetSmartqStatusResponse
func (client *Client) QueryDatasetSmartqStatus(request *QueryDatasetSmartqStatusRequest) (_result *QueryDatasetSmartqStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDatasetSmartqStatusResponse{}
	_body, _err := client.QueryDatasetSmartqStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定数据集的行级权限开关状态。
//
// @param request - QueryDatasetSwitchInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetSwitchInfoResponse
func (client *Client) QueryDatasetSwitchInfoWithOptions(request *QueryDatasetSwitchInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetSwitchInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDatasetSwitchInfo"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDatasetSwitchInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定数据集的行级权限开关状态。
//
// @param request - QueryDatasetSwitchInfoRequest
//
// @return QueryDatasetSwitchInfoResponse
func (client *Client) QueryDatasetSwitchInfo(request *QueryDatasetSwitchInfoRequest) (_result *QueryDatasetSwitchInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDatasetSwitchInfoResponse{}
	_body, _err := client.QueryDatasetSwitchInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the embedding configuration in the organization, including the maximum number of embeddings and the number of embeddings.
//
// @param request - QueryEmbeddedInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEmbeddedInfoResponse
func (client *Client) QueryEmbeddedInfoWithOptions(runtime *dara.RuntimeOptions) (_result *QueryEmbeddedInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEmbeddedInfo"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEmbeddedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the embedding configuration in the organization, including the maximum number of embeddings and the number of embeddings.
//
// @return QueryEmbeddedInfoResponse
func (client *Client) QueryEmbeddedInfo() (_result *QueryEmbeddedInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEmbeddedInfoResponse{}
	_body, _err := client.QueryEmbeddedInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether embedding is enabled for a report.
//
// @param request - QueryEmbeddedStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEmbeddedStatusResponse
func (client *Client) QueryEmbeddedStatusWithOptions(request *QueryEmbeddedStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryEmbeddedStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEmbeddedStatus"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEmbeddedStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether embedding is enabled for a report.
//
// @param request - QueryEmbeddedStatusRequest
//
// @return QueryEmbeddedStatusResponse
func (client *Client) QueryEmbeddedStatus(request *QueryEmbeddedStatusRequest) (_result *QueryEmbeddedStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEmbeddedStatusResponse{}
	_body, _err := client.QueryEmbeddedStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check which datasets and analysis themes the user has question authorization for
//
// @param request - QueryLlmCubeWithThemeListByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLlmCubeWithThemeListByUserIdResponse
func (client *Client) QueryLlmCubeWithThemeListByUserIdWithOptions(request *QueryLlmCubeWithThemeListByUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryLlmCubeWithThemeListByUserIdResponse, _err error) {
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
		Action:      dara.String("QueryLlmCubeWithThemeListByUserId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLlmCubeWithThemeListByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check which datasets and analysis themes the user has question authorization for
//
// @param request - QueryLlmCubeWithThemeListByUserIdRequest
//
// @return QueryLlmCubeWithThemeListByUserIdResponse
func (client *Client) QueryLlmCubeWithThemeListByUserId(request *QueryLlmCubeWithThemeListByUserIdRequest) (_result *QueryLlmCubeWithThemeListByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryLlmCubeWithThemeListByUserIdResponse{}
	_body, _err := client.QueryLlmCubeWithThemeListByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the configuration of the specified organization role.
//
// @param request - QueryOrganizationRoleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrganizationRoleConfigResponse
func (client *Client) QueryOrganizationRoleConfigWithOptions(request *QueryOrganizationRoleConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryOrganizationRoleConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOrganizationRoleConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOrganizationRoleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the configuration of the specified organization role.
//
// @param request - QueryOrganizationRoleConfigRequest
//
// @return QueryOrganizationRoleConfigResponse
func (client *Client) QueryOrganizationRoleConfig(request *QueryOrganizationRoleConfigRequest) (_result *QueryOrganizationRoleConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryOrganizationRoleConfigResponse{}
	_body, _err := client.QueryOrganizationRoleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the list of workspaces under the current organization.
//
// @param request - QueryOrganizationWorkspaceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrganizationWorkspaceListResponse
func (client *Client) QueryOrganizationWorkspaceListWithOptions(request *QueryOrganizationWorkspaceListRequest, runtime *dara.RuntimeOptions) (_result *QueryOrganizationWorkspaceListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
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
		Action:      dara.String("QueryOrganizationWorkspaceList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOrganizationWorkspaceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of workspaces under the current organization.
//
// @param request - QueryOrganizationWorkspaceListRequest
//
// @return QueryOrganizationWorkspaceListResponse
func (client *Client) QueryOrganizationWorkspaceList(request *QueryOrganizationWorkspaceListRequest) (_result *QueryOrganizationWorkspaceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryOrganizationWorkspaceListResponse{}
	_body, _err := client.QueryOrganizationWorkspaceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI QueryReadableResourcesListByUserId is deprecated, please use quickbi-public::2022-01-01::QueryReadableResourcesListByUserIdV2 instead.
//
// Summary:
//
// Queries the list of works that a user has the permission to view, including the statements that are authorized to share in a space.
//
// @param request - QueryReadableResourcesListByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReadableResourcesListByUserIdResponse
// Deprecated
func (client *Client) QueryReadableResourcesListByUserIdWithOptions(request *QueryReadableResourcesListByUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryReadableResourcesListByUserIdResponse, _err error) {
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
		Action:      dara.String("QueryReadableResourcesListByUserId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReadableResourcesListByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI QueryReadableResourcesListByUserId is deprecated, please use quickbi-public::2022-01-01::QueryReadableResourcesListByUserIdV2 instead.
//
// Summary:
//
// Queries the list of works that a user has the permission to view, including the statements that are authorized to share in a space.
//
// @param request - QueryReadableResourcesListByUserIdRequest
//
// @return QueryReadableResourcesListByUserIdResponse
// Deprecated
func (client *Client) QueryReadableResourcesListByUserId(request *QueryReadableResourcesListByUserIdRequest) (_result *QueryReadableResourcesListByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryReadableResourcesListByUserIdResponse{}
	_body, _err := client.QueryReadableResourcesListByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户有权查看的作品列表(新)
//
// @param request - QueryReadableResourcesListByUserIdV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReadableResourcesListByUserIdV2Response
func (client *Client) QueryReadableResourcesListByUserIdV2WithOptions(request *QueryReadableResourcesListByUserIdV2Request, runtime *dara.RuntimeOptions) (_result *QueryReadableResourcesListByUserIdV2Response, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkType) {
		query["WorkType"] = request.WorkType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryReadableResourcesListByUserIdV2"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReadableResourcesListByUserIdV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户有权查看的作品列表(新)
//
// @param request - QueryReadableResourcesListByUserIdV2Request
//
// @return QueryReadableResourcesListByUserIdV2Response
func (client *Client) QueryReadableResourcesListByUserIdV2(request *QueryReadableResourcesListByUserIdV2Request) (_result *QueryReadableResourcesListByUserIdV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryReadableResourcesListByUserIdV2Response{}
	_body, _err := client.QueryReadableResourcesListByUserIdV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries report performance logs.
//
// @param request - QueryReportPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReportPerformanceResponse
func (client *Client) QueryReportPerformanceWithOptions(request *QueryReportPerformanceRequest, runtime *dara.RuntimeOptions) (_result *QueryReportPerformanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CostTimeAvgMin) {
		query["CostTimeAvgMin"] = request.CostTimeAvgMin
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryReportPerformance"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReportPerformanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries report performance logs.
//
// @param request - QueryReportPerformanceRequest
//
// @return QueryReportPerformanceResponse
func (client *Client) QueryReportPerformance(request *QueryReportPerformanceRequest) (_result *QueryReportPerformanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryReportPerformanceResponse{}
	_body, _err := client.QueryReportPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of objects to which a work has been shared, returning only the sharing configurations that are still within their validity period.
//
// @param request - QueryShareListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryShareListResponse
func (client *Client) QueryShareListWithOptions(request *QueryShareListRequest, runtime *dara.RuntimeOptions) (_result *QueryShareListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryShareList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryShareListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of objects to which a work has been shared, returning only the sharing configurations that are still within their validity period.
//
// @param request - QueryShareListRequest
//
// @return QueryShareListResponse
func (client *Client) QueryShareList(request *QueryShareListRequest) (_result *QueryShareListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryShareListResponse{}
	_body, _err := client.QueryShareListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to query the list of works authorized to a user.
//
// @param request - QuerySharesToUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySharesToUserListResponse
func (client *Client) QuerySharesToUserListWithOptions(request *QuerySharesToUserListRequest, runtime *dara.RuntimeOptions) (_result *QuerySharesToUserListResponse, _err error) {
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
		Action:      dara.String("QuerySharesToUserList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySharesToUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the list of works authorized to a user.
//
// @param request - QuerySharesToUserListRequest
//
// @return QuerySharesToUserListResponse
func (client *Client) QuerySharesToUserList(request *QuerySharesToUserListRequest) (_result *QuerySharesToUserListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySharesToUserListResponse{}
	_body, _err := client.QuerySharesToUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check if a user has permission for a specific smart question dataset
//
// @param request - QuerySmartqPermissionByCubeIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmartqPermissionByCubeIdResponse
func (client *Client) QuerySmartqPermissionByCubeIdWithOptions(request *QuerySmartqPermissionByCubeIdRequest, runtime *dara.RuntimeOptions) (_result *QuerySmartqPermissionByCubeIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmartqPermissionByCubeId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmartqPermissionByCubeIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check if a user has permission for a specific smart question dataset
//
// @param request - QuerySmartqPermissionByCubeIdRequest
//
// @return QuerySmartqPermissionByCubeIdResponse
func (client *Client) QuerySmartqPermissionByCubeId(request *QuerySmartqPermissionByCubeIdRequest) (_result *QuerySmartqPermissionByCubeIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySmartqPermissionByCubeIdResponse{}
	_body, _err := client.QuerySmartqPermissionByCubeIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a specified ticket for a report that is not embedded in the report.
//
// @param request - QueryTicketInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTicketInfoResponse
func (client *Client) QueryTicketInfoWithOptions(request *QueryTicketInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryTicketInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ticket) {
		query["Ticket"] = request.Ticket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTicketInfo"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTicketInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a specified ticket for a report that is not embedded in the report.
//
// @param request - QueryTicketInfoRequest
//
// @return QueryTicketInfoResponse
func (client *Client) QueryTicketInfo(request *QueryTicketInfoRequest) (_result *QueryTicketInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTicketInfoResponse{}
	_body, _err := client.QueryTicketInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can this operation to obtain information about child user groups under a specified parent user group.
//
// @param request - QueryUserGroupListByParentIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserGroupListByParentIdResponse
func (client *Client) QueryUserGroupListByParentIdWithOptions(request *QueryUserGroupListByParentIdRequest, runtime *dara.RuntimeOptions) (_result *QueryUserGroupListByParentIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParentUserGroupId) {
		query["ParentUserGroupId"] = request.ParentUserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserGroupListByParentId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserGroupListByParentIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can this operation to obtain information about child user groups under a specified parent user group.
//
// @param request - QueryUserGroupListByParentIdRequest
//
// @return QueryUserGroupListByParentIdResponse
func (client *Client) QueryUserGroupListByParentId(request *QueryUserGroupListByParentIdRequest) (_result *QueryUserGroupListByParentIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserGroupListByParentIdResponse{}
	_body, _err := client.QueryUserGroupListByParentIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the list of members under a user group.
//
// @param request - QueryUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserGroupMemberResponse
func (client *Client) QueryUserGroupMemberWithOptions(request *QueryUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *QueryUserGroupMemberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserGroupMember"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserGroupMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of members under a user group.
//
// @param request - QueryUserGroupMemberRequest
//
// @return QueryUserGroupMemberResponse
func (client *Client) QueryUserGroupMember(request *QueryUserGroupMemberRequest) (_result *QueryUserGroupMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserGroupMemberResponse{}
	_body, _err := client.QueryUserGroupMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user information based on the Alibaba Cloud ID or Alibaba Cloud account name.
//
// @param request - QueryUserInfoByAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserInfoByAccountResponse
func (client *Client) QueryUserInfoByAccountWithOptions(request *QueryUserInfoByAccountRequest, runtime *dara.RuntimeOptions) (_result *QueryUserInfoByAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.ParentAccountName) {
		query["ParentAccountName"] = request.ParentAccountName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserInfoByAccount"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserInfoByAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user information based on the Alibaba Cloud ID or Alibaba Cloud account name.
//
// @param request - QueryUserInfoByAccountRequest
//
// @return QueryUserInfoByAccountResponse
func (client *Client) QueryUserInfoByAccount(request *QueryUserInfoByAccountRequest) (_result *QueryUserInfoByAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserInfoByAccountResponse{}
	_body, _err := client.QueryUserInfoByAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user information based on the user ID.
//
// @param request - QueryUserInfoByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserInfoByUserIdResponse
func (client *Client) QueryUserInfoByUserIdWithOptions(request *QueryUserInfoByUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryUserInfoByUserIdResponse, _err error) {
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
		Action:      dara.String("QueryUserInfoByUserId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserInfoByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user information based on the user ID.
//
// @param request - QueryUserInfoByUserIdRequest
//
// @return QueryUserInfoByUserIdResponse
func (client *Client) QueryUserInfoByUserId(request *QueryUserInfoByUserIdRequest) (_result *QueryUserInfoByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserInfoByUserIdResponse{}
	_body, _err := client.QueryUserInfoByUserIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the members of an organization.
//
// @param request - QueryUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserListResponse
func (client *Client) QueryUserListWithOptions(request *QueryUserListRequest, runtime *dara.RuntimeOptions) (_result *QueryUserListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the members of an organization.
//
// @param request - QueryUserListRequest
//
// @return QueryUserListResponse
func (client *Client) QueryUserList(request *QueryUserListRequest) (_result *QueryUserListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserListResponse{}
	_body, _err := client.QueryUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the preset workspace role information for a specified workspace member.
//
// @param request - QueryUserRoleInfoInWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserRoleInfoInWorkspaceResponse
func (client *Client) QueryUserRoleInfoInWorkspaceWithOptions(request *QueryUserRoleInfoInWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *QueryUserRoleInfoInWorkspaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserRoleInfoInWorkspace"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserRoleInfoInWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the preset workspace role information for a specified workspace member.
//
// @param request - QueryUserRoleInfoInWorkspaceRequest
//
// @return QueryUserRoleInfoInWorkspaceResponse
func (client *Client) QueryUserRoleInfoInWorkspace(request *QueryUserRoleInfoInWorkspaceRequest) (_result *QueryUserRoleInfoInWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserRoleInfoInWorkspaceResponse{}
	_body, _err := client.QueryUserRoleInfoInWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata list of member tags in an organization.
//
// @param request - QueryUserTagMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserTagMetaListResponse
func (client *Client) QueryUserTagMetaListWithOptions(runtime *dara.RuntimeOptions) (_result *QueryUserTagMetaListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserTagMetaList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserTagMetaListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata list of member tags in an organization.
//
// @return QueryUserTagMetaListResponse
func (client *Client) QueryUserTagMetaList() (_result *QueryUserTagMetaListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserTagMetaListResponse{}
	_body, _err := client.QueryUserTagMetaListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of specific user tag values.
//
// @param request - QueryUserTagValueListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserTagValueListResponse
func (client *Client) QueryUserTagValueListWithOptions(request *QueryUserTagValueListRequest, runtime *dara.RuntimeOptions) (_result *QueryUserTagValueListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserTagValueList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserTagValueListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of specific user tag values.
//
// @param request - QueryUserTagValueListRequest
//
// @return QueryUserTagValueListResponse
func (client *Client) QueryUserTagValueList(request *QueryUserTagValueListRequest) (_result *QueryUserTagValueListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserTagValueListResponse{}
	_body, _err := client.QueryUserTagValueListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specified data work.
//
// @param request - QueryWorksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksResponse
func (client *Client) QueryWorksWithOptions(request *QueryWorksRequest, runtime *dara.RuntimeOptions) (_result *QueryWorksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryWorks"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWorksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specified data work.
//
// @param request - QueryWorksRequest
//
// @return QueryWorksResponse
func (client *Client) QueryWorks(request *QueryWorksRequest) (_result *QueryWorksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryWorksResponse{}
	_body, _err := client.QueryWorksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the kinship of a data work, including the datasets referenced by each component and query field information. Currently, only supported data works include dashboards, workbooks, and self-service data retrieval.
//
// @param request - QueryWorksBloodRelationshipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksBloodRelationshipResponse
func (client *Client) QueryWorksBloodRelationshipWithOptions(request *QueryWorksBloodRelationshipRequest, runtime *dara.RuntimeOptions) (_result *QueryWorksBloodRelationshipResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryWorksBloodRelationship"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWorksBloodRelationshipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the kinship of a data work, including the datasets referenced by each component and query field information. Currently, only supported data works include dashboards, workbooks, and self-service data retrieval.
//
// @param request - QueryWorksBloodRelationshipRequest
//
// @return QueryWorksBloodRelationshipResponse
func (client *Client) QueryWorksBloodRelationship(request *QueryWorksBloodRelationshipRequest) (_result *QueryWorksBloodRelationshipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryWorksBloodRelationshipResponse{}
	_body, _err := client.QueryWorksBloodRelationshipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query all works under the entire organization, with the option to specify the type of work.
//
// @param request - QueryWorksByOrganizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksByOrganizationResponse
func (client *Client) QueryWorksByOrganizationWithOptions(request *QueryWorksByOrganizationRequest, runtime *dara.RuntimeOptions) (_result *QueryWorksByOrganizationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.ThirdPartAuthFlag) {
		query["ThirdPartAuthFlag"] = request.ThirdPartAuthFlag
	}

	if !dara.IsNil(request.WorksType) {
		query["WorksType"] = request.WorksType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryWorksByOrganization"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWorksByOrganizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query all works under the entire organization, with the option to specify the type of work.
//
// @param request - QueryWorksByOrganizationRequest
//
// @return QueryWorksByOrganizationResponse
func (client *Client) QueryWorksByOrganization(request *QueryWorksByOrganizationRequest) (_result *QueryWorksByOrganizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryWorksByOrganizationResponse{}
	_body, _err := client.QueryWorksByOrganizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all works in a workspace under an organization. You can specify the type of work.
//
// @param request - QueryWorksByWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksByWorkspaceResponse
func (client *Client) QueryWorksByWorkspaceWithOptions(request *QueryWorksByWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *QueryWorksByWorkspaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.ThirdPartAuthFlag) {
		query["ThirdPartAuthFlag"] = request.ThirdPartAuthFlag
	}

	if !dara.IsNil(request.WorksType) {
		query["WorksType"] = request.WorksType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryWorksByWorkspace"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWorksByWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all works in a workspace under an organization. You can specify the type of work.
//
// @param request - QueryWorksByWorkspaceRequest
//
// @return QueryWorksByWorkspaceResponse
func (client *Client) QueryWorksByWorkspace(request *QueryWorksByWorkspaceRequest) (_result *QueryWorksByWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryWorksByWorkspaceResponse{}
	_body, _err := client.QueryWorksByWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Configuration Information for a Specified Workspace Role
//
// @param request - QueryWorkspaceRoleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorkspaceRoleConfigResponse
func (client *Client) QueryWorkspaceRoleConfigWithOptions(request *QueryWorkspaceRoleConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryWorkspaceRoleConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryWorkspaceRoleConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWorkspaceRoleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Configuration Information for a Specified Workspace Role
//
// @param request - QueryWorkspaceRoleConfigRequest
//
// @return QueryWorkspaceRoleConfigResponse
func (client *Client) QueryWorkspaceRoleConfig(request *QueryWorkspaceRoleConfigRequest) (_result *QueryWorkspaceRoleConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryWorkspaceRoleConfigResponse{}
	_body, _err := client.QueryWorkspaceRoleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of members under a specified workspace.
//
// @param request - QueryWorkspaceUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorkspaceUserListResponse
func (client *Client) QueryWorkspaceUserListWithOptions(request *QueryWorkspaceUserListRequest, runtime *dara.RuntimeOptions) (_result *QueryWorkspaceUserListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryWorkspaceUserList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryWorkspaceUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of members under a specified workspace.
//
// @param request - QueryWorkspaceUserListRequest
//
// @return QueryWorkspaceUserListResponse
func (client *Client) QueryWorkspaceUserList(request *QueryWorkspaceUserListRequest) (_result *QueryWorkspaceUserListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryWorkspaceUserListResponse{}
	_body, _err := client.QueryWorkspaceUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can customize the callback interface for approval processes to process Quick BI approval processes.
//
// @param request - ResultCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResultCallbackResponse
func (client *Client) ResultCallbackWithOptions(request *ResultCallbackRequest, runtime *dara.RuntimeOptions) (_result *ResultCallbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.HandleReason) {
		query["HandleReason"] = request.HandleReason
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResultCallback"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResultCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can customize the callback interface for approval processes to process Quick BI approval processes.
//
// @param request - ResultCallbackRequest
//
// @return ResultCallbackResponse
func (client *Client) ResultCallback(request *ResultCallbackRequest) (_result *ResultCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResultCallbackResponse{}
	_body, _err := client.ResultCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a user\\"s favorite work
//
// @param request - SaveFavoritesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveFavoritesResponse
func (client *Client) SaveFavoritesWithOptions(request *SaveFavoritesRequest, runtime *dara.RuntimeOptions) (_result *SaveFavoritesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveFavorites"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveFavoritesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a user\\"s favorite work
//
// @param request - SaveFavoritesRequest
//
// @return SaveFavoritesResponse
func (client *Client) SaveFavorites(request *SaveFavoritesRequest) (_result *SaveFavoritesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveFavoritesResponse{}
	_body, _err := client.SaveFavoritesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置行列权限的额外配置
//
// @param request - SetDataLevelPermissionExtraConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionExtraConfigResponse
func (client *Client) SetDataLevelPermissionExtraConfigWithOptions(request *SetDataLevelPermissionExtraConfigRequest, runtime *dara.RuntimeOptions) (_result *SetDataLevelPermissionExtraConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.MissHitPolicy) {
		query["MissHitPolicy"] = request.MissHitPolicy
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDataLevelPermissionExtraConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDataLevelPermissionExtraConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置行列权限的额外配置
//
// @param request - SetDataLevelPermissionExtraConfigRequest
//
// @return SetDataLevelPermissionExtraConfigResponse
func (client *Client) SetDataLevelPermissionExtraConfig(request *SetDataLevelPermissionExtraConfigRequest) (_result *SetDataLevelPermissionExtraConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDataLevelPermissionExtraConfigResponse{}
	_body, _err := client.SetDataLevelPermissionExtraConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置单条数据集行列权限配置信息（新增和更新）
//
// @param request - SetDataLevelPermissionRuleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionRuleConfigResponse
func (client *Client) SetDataLevelPermissionRuleConfigWithOptions(request *SetDataLevelPermissionRuleConfigRequest, runtime *dara.RuntimeOptions) (_result *SetDataLevelPermissionRuleConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleModel) {
		query["RuleModel"] = request.RuleModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDataLevelPermissionRuleConfig"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDataLevelPermissionRuleConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置单条数据集行列权限配置信息（新增和更新）
//
// @param request - SetDataLevelPermissionRuleConfigRequest
//
// @return SetDataLevelPermissionRuleConfigResponse
func (client *Client) SetDataLevelPermissionRuleConfig(request *SetDataLevelPermissionRuleConfigRequest) (_result *SetDataLevelPermissionRuleConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDataLevelPermissionRuleConfigResponse{}
	_body, _err := client.SetDataLevelPermissionRuleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the whitelist for the specified row-level permissions.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.
//
// @param request - SetDataLevelPermissionWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionWhiteListResponse
func (client *Client) SetDataLevelPermissionWhiteListWithOptions(request *SetDataLevelPermissionWhiteListRequest, runtime *dara.RuntimeOptions) (_result *SetDataLevelPermissionWhiteListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WhiteListModel) {
		query["WhiteListModel"] = request.WhiteListModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDataLevelPermissionWhiteList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDataLevelPermissionWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the whitelist for the specified row-level permissions.
//
// Description:
//
// > : You can only Quick BI the new row-column permission model. If you are still using the old row-column permission model, migrate to the new row-column permission model before you call this operation. To migrate row-level permissions to the new row-level permission model, perform the following steps: Choose Organizations> Security Configurations> Upgrade Row-Level Permissions. On the Upgrade Row-Level Permissions page, click **Upgrade**.
//
// @param request - SetDataLevelPermissionWhiteListRequest
//
// @return SetDataLevelPermissionWhiteListResponse
func (client *Client) SetDataLevelPermissionWhiteList(request *SetDataLevelPermissionWhiteListRequest) (_result *SetDataLevelPermissionWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDataLevelPermissionWhiteListResponse{}
	_body, _err := client.SetDataLevelPermissionWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Synchronize the question count permissions of a specified user to other users
//
// @param request - SmartqAuthTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartqAuthTransferResponse
func (client *Client) SmartqAuthTransferWithOptions(request *SmartqAuthTransferRequest, runtime *dara.RuntimeOptions) (_result *SmartqAuthTransferResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OriginUserId) {
		query["OriginUserId"] = request.OriginUserId
	}

	if !dara.IsNil(request.TargetUserIds) {
		query["TargetUserIds"] = request.TargetUserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SmartqAuthTransfer"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmartqAuthTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Synchronize the question count permissions of a specified user to other users
//
// @param request - SmartqAuthTransferRequest
//
// @return SmartqAuthTransferResponse
func (client *Client) SmartqAuthTransfer(request *SmartqAuthTransferRequest) (_result *SmartqAuthTransferResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SmartqAuthTransferResponse{}
	_body, _err := client.SmartqAuthTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch Management of Smart Q\\\\\\&A Authorizations
//
// Description:
//
// Used for batch management of smart Q&A authorizations. Repeatedly adding an authorization will be treated as a new addition; repeatedly deleting an authorization will be skipped by default and will not be recorded in the audit log.
//
// @param request - SmartqAuthorizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartqAuthorizeResponse
func (client *Client) SmartqAuthorizeWithOptions(request *SmartqAuthorizeRequest, runtime *dara.RuntimeOptions) (_result *SmartqAuthorizeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeIds) {
		query["CubeIds"] = request.CubeIds
	}

	if !dara.IsNil(request.ExpireDay) {
		query["ExpireDay"] = request.ExpireDay
	}

	if !dara.IsNil(request.LlmCubeThemes) {
		query["LlmCubeThemes"] = request.LlmCubeThemes
	}

	if !dara.IsNil(request.LlmCubes) {
		query["LlmCubes"] = request.LlmCubes
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SmartqAuthorize"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmartqAuthorizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch Management of Smart Q\\\\\\&A Authorizations
//
// Description:
//
// Used for batch management of smart Q&A authorizations. Repeatedly adding an authorization will be treated as a new addition; repeatedly deleting an authorization will be skipped by default and will not be recorded in the audit log.
//
// @param request - SmartqAuthorizeRequest
//
// @return SmartqAuthorizeResponse
func (client *Client) SmartqAuthorize(request *SmartqAuthorizeRequest) (_result *SmartqAuthorizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SmartqAuthorizeResponse{}
	_body, _err := client.SmartqAuthorizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Capability Open
//
// Description:
//
// Special Note: When a user is authorized to call this API, it is assumed that the user has the permission to query the corresponding data by passing in the userId as that user.
//
// @param request - SmartqQueryAbilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartqQueryAbilityResponse
func (client *Client) SmartqQueryAbilityWithOptions(request *SmartqQueryAbilityRequest, runtime *dara.RuntimeOptions) (_result *SmartqQueryAbilityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserQuestion) {
		query["UserQuestion"] = request.UserQuestion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SmartqQueryAbility"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmartqQueryAbilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Capability Open
//
// Description:
//
// Special Note: When a user is authorized to call this API, it is assumed that the user has the permission to query the corresponding data by passing in the userId as that user.
//
// @param request - SmartqQueryAbilityRequest
//
// @return SmartqQueryAbilityResponse
func (client *Client) SmartqQueryAbility(request *SmartqQueryAbilityRequest) (_result *SmartqQueryAbilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SmartqQueryAbilityResponse{}
	_body, _err := client.SmartqQueryAbilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Indicates whether the request is successful. Valid values:
//
//   - true: The request was successful.
//
//   - false: The request failed.
//
// Description:
//
// The execution result of the interface. Valid values:
//
//   - true: The request was successful.
//
//   - false: The request failed.
//
// @param request - UpdateDataLevelPermissionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLevelPermissionStatusResponse
func (client *Client) UpdateDataLevelPermissionStatusWithOptions(request *UpdateDataLevelPermissionStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLevelPermissionStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.IsOpen) {
		query["IsOpen"] = request.IsOpen
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLevelPermissionStatus"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLevelPermissionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Indicates whether the request is successful. Valid values:
//
//   - true: The request was successful.
//
//   - false: The request failed.
//
// Description:
//
// The execution result of the interface. Valid values:
//
//   - true: The request was successful.
//
//   - false: The request failed.
//
// @param request - UpdateDataLevelPermissionStatusRequest
//
// @return UpdateDataLevelPermissionStatusResponse
func (client *Client) UpdateDataLevelPermissionStatus(request *UpdateDataLevelPermissionStatusRequest) (_result *UpdateDataLevelPermissionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLevelPermissionStatusResponse{}
	_body, _err := client.UpdateDataLevelPermissionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Change the embedding status of a report, turn on embedding, or turn off embedding.
//
// @param request - UpdateEmbeddedStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEmbeddedStatusResponse
func (client *Client) UpdateEmbeddedStatusWithOptions(request *UpdateEmbeddedStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateEmbeddedStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ThirdPartAuthFlag) {
		query["ThirdPartAuthFlag"] = request.ThirdPartAuthFlag
	}

	if !dara.IsNil(request.WorksId) {
		query["WorksId"] = request.WorksId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEmbeddedStatus"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEmbeddedStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Change the embedding status of a report, turn on embedding, or turn off embedding.
//
// @param request - UpdateEmbeddedStatusRequest
//
// @return UpdateEmbeddedStatusResponse
func (client *Client) UpdateEmbeddedStatus(request *UpdateEmbeddedStatusRequest) (_result *UpdateEmbeddedStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEmbeddedStatusResponse{}
	_body, _err := client.UpdateEmbeddedStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the ticket quantity on the specified ticket used for the exemption embedded report.
//
// @param request - UpdateTicketNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketNumResponse
func (client *Client) UpdateTicketNumWithOptions(request *UpdateTicketNumRequest, runtime *dara.RuntimeOptions) (_result *UpdateTicketNumResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ticket) {
		query["Ticket"] = request.Ticket
	}

	if !dara.IsNil(request.TicketNum) {
		query["TicketNum"] = request.TicketNum
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTicketNum"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTicketNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the ticket quantity on the specified ticket used for the exemption embedded report.
//
// @param request - UpdateTicketNumRequest
//
// @return UpdateTicketNumResponse
func (client *Client) UpdateTicketNum(request *UpdateTicketNumRequest) (_result *UpdateTicketNumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTicketNumResponse{}
	_body, _err := client.UpdateTicketNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information of a specified member in an organization.
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
	if !dara.IsNil(request.AdminUser) {
		query["AdminUser"] = request.AdminUser
	}

	if !dara.IsNil(request.AuthAdminUser) {
		query["AuthAdminUser"] = request.AuthAdminUser
	}

	if !dara.IsNil(request.IsDeleted) {
		query["IsDeleted"] = request.IsDeleted
	}

	if !dara.IsNil(request.NickName) {
		query["NickName"] = request.NickName
	}

	if !dara.IsNil(request.RoleIds) {
		query["RoleIds"] = request.RoleIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2022-01-01"),
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
// Updates the information of a specified member in an organization.
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
// Updates information about a specified user group in an organization.
//
// @param request - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithOptions(request *UpdateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupDescription) {
		query["UserGroupDescription"] = request.UserGroupDescription
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
		Action:      dara.String("UpdateUserGroup"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information about a specified user group in an organization.
//
// @param request - UpdateUserGroupRequest
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (_result *UpdateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.UpdateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Used for updating the metadata of organization member tags
//
// @param request - UpdateUserTagMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserTagMetaResponse
func (client *Client) UpdateUserTagMetaWithOptions(request *UpdateUserTagMetaRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserTagMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TagDescription) {
		query["TagDescription"] = request.TagDescription
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserTagMeta"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserTagMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Used for updating the metadata of organization member tags
//
// @param request - UpdateUserTagMetaRequest
//
// @return UpdateUserTagMetaResponse
func (client *Client) UpdateUserTagMeta(request *UpdateUserTagMetaRequest) (_result *UpdateUserTagMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserTagMetaResponse{}
	_body, _err := client.UpdateUserTagMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the tag value of an organization member.
//
// @param request - UpdateUserTagValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserTagValueResponse
func (client *Client) UpdateUserTagValueWithOptions(request *UpdateUserTagValueRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserTagValueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	if !dara.IsNil(request.TagValue) {
		query["TagValue"] = request.TagValue
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserTagValue"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserTagValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the tag value of an organization member.
//
// @param request - UpdateUserTagValueRequest
//
// @return UpdateUserTagValueResponse
func (client *Client) UpdateUserTagValue(request *UpdateUserTagValueRequest) (_result *UpdateUserTagValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserTagValueResponse{}
	_body, _err := client.UpdateUserTagValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the role of a specified member under the workspace, existing roles will be overwritten.
//
// @param request - UpdateWorkspaceUserRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceUserRoleResponse
func (client *Client) UpdateWorkspaceUserRoleWithOptions(request *UpdateWorkspaceUserRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceUserRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.RoleIds) {
		query["RoleIds"] = request.RoleIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspaceUserRole"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceUserRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the role of a specified member under the workspace, existing roles will be overwritten.
//
// @param request - UpdateWorkspaceUserRoleRequest
//
// @return UpdateWorkspaceUserRoleResponse
func (client *Client) UpdateWorkspaceUserRole(request *UpdateWorkspaceUserRoleRequest) (_result *UpdateWorkspaceUserRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkspaceUserRoleResponse{}
	_body, _err := client.UpdateWorkspaceUserRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch update the role information of workspace members, existing roles will be overwritten
//
// @param request - UpdateWorkspaceUsersRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceUsersRoleResponse
func (client *Client) UpdateWorkspaceUsersRoleWithOptions(request *UpdateWorkspaceUsersRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceUsersRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspaceUsersRole"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceUsersRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch update the role information of workspace members, existing roles will be overwritten
//
// @param request - UpdateWorkspaceUsersRoleRequest
//
// @return UpdateWorkspaceUsersRoleResponse
func (client *Client) UpdateWorkspaceUsersRole(request *UpdateWorkspaceUsersRoleRequest) (_result *UpdateWorkspaceUsersRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkspaceUsersRoleResponse{}
	_body, _err := client.UpdateWorkspaceUsersRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Make the user exit all user groups. This process is irreversible. Exercise caution when performing this operation.
//
// @param request - WithdrawAllUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawAllUserGroupsResponse
func (client *Client) WithdrawAllUserGroupsWithOptions(request *WithdrawAllUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *WithdrawAllUserGroupsResponse, _err error) {
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
		Action:      dara.String("WithdrawAllUserGroups"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WithdrawAllUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Make the user exit all user groups. This process is irreversible. Exercise caution when performing this operation.
//
// @param request - WithdrawAllUserGroupsRequest
//
// @return WithdrawAllUserGroupsResponse
func (client *Client) WithdrawAllUserGroups(request *WithdrawAllUserGroupsRequest) (_result *WithdrawAllUserGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &WithdrawAllUserGroupsResponse{}
	_body, _err := client.WithdrawAllUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
