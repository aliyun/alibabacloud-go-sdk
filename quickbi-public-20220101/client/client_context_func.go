// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AddDataLevelPermissionRuleUsersWithContext(ctx context.Context, request *AddDataLevelPermissionRuleUsersRequest, runtime *dara.RuntimeOptions) (_result *AddDataLevelPermissionRuleUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataLevelPermissionWhiteListResponse
func (client *Client) AddDataLevelPermissionWhiteListWithContext(ctx context.Context, request *AddDataLevelPermissionWhiteListRequest, runtime *dara.RuntimeOptions) (_result *AddDataLevelPermissionWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据源
//
// @param request - AddDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataSourceResponse
func (client *Client) AddDataSourceWithContext(ctx context.Context, request *AddDataSourceRequest, runtime *dara.RuntimeOptions) (_result *AddDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddModel) {
		query["AddModel"] = request.AddModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDataSource"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShareReportResponse
func (client *Client) AddShareReportWithContext(ctx context.Context, request *AddShareReportRequest, runtime *dara.RuntimeOptions) (_result *AddShareReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserResponse
func (client *Client) AddUserWithContext(ctx context.Context, request *AddUserRequest, runtime *dara.RuntimeOptions) (_result *AddUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMemberWithContext(ctx context.Context, request *AddUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *AddUserGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMembersResponse
func (client *Client) AddUserGroupMembersWithContext(ctx context.Context, request *AddUserGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *AddUserGroupMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserTagMetaResponse
func (client *Client) AddUserTagMetaWithContext(ctx context.Context, request *AddUserTagMetaRequest, runtime *dara.RuntimeOptions) (_result *AddUserTagMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToWorkspaceResponse
func (client *Client) AddUserToWorkspaceWithContext(ctx context.Context, request *AddUserToWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *AddUserToWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWorkspaceUsersResponse
func (client *Client) AddWorkspaceUsersWithContext(ctx context.Context, request *AddWorkspaceUsersRequest, runtime *dara.RuntimeOptions) (_result *AddWorkspaceUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllotDatasetAccelerationTaskResponse
func (client *Client) AllotDatasetAccelerationTaskWithContext(ctx context.Context, request *AllotDatasetAccelerationTaskRequest, runtime *dara.RuntimeOptions) (_result *AllotDatasetAccelerationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeMenuResponse
func (client *Client) AuthorizeMenuWithContext(ctx context.Context, request *AuthorizeMenuRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeMenuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Batch add Feishu users.
//
// @param request - BatchAddFeishuUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddFeishuUsersResponse
func (client *Client) BatchAddFeishuUsersWithContext(ctx context.Context, request *BatchAddFeishuUsersRequest, runtime *dara.RuntimeOptions) (_result *BatchAddFeishuUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAuthorizationMenuResponse
func (client *Client) CancelAuthorizationMenuWithContext(ctx context.Context, request *CancelAuthorizationMenuRequest, runtime *dara.RuntimeOptions) (_result *CancelAuthorizationMenuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCollectionResponse
func (client *Client) CancelCollectionWithContext(ctx context.Context, request *CancelCollectionRequest, runtime *dara.RuntimeOptions) (_result *CancelCollectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelReportShareResponse
func (client *Client) CancelReportShareWithContext(ctx context.Context, request *CancelReportShareRequest, runtime *dara.RuntimeOptions) (_result *CancelReportShareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeVisibilityModelResponse
func (client *Client) ChangeVisibilityModelWithContext(ctx context.Context, request *ChangeVisibilityModelRequest, runtime *dara.RuntimeOptions) (_result *ChangeVisibilityModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查给定的cubeId是否存在
//
// @param request - CheckDatasetExistedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDatasetExistedResponse
func (client *Client) CheckDatasetExistedWithContext(ctx context.Context, request *CheckDatasetExistedRequest, runtime *dara.RuntimeOptions) (_result *CheckDatasetExistedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDatasetExisted"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDatasetExistedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 判断用户是否属于组织
//
// @param request - CheckOrganizationMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckOrganizationMemberResponse
func (client *Client) CheckOrganizationMemberWithContext(ctx context.Context, request *CheckOrganizationMemberRequest, runtime *dara.RuntimeOptions) (_result *CheckOrganizationMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckOrganizationMember"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckOrganizationMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckReadableResponse
func (client *Client) CheckReadableWithContext(ctx context.Context, request *CheckReadableRequest, runtime *dara.RuntimeOptions) (_result *CheckReadableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据自定义sql创建数据集
//
// @param request - CreateCubeBySqlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCubeBySqlResponse
func (client *Client) CreateCubeBySqlWithContext(ctx context.Context, request *CreateCubeBySqlRequest, runtime *dara.RuntimeOptions) (_result *CreateCubeBySqlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Caption) {
		query["Caption"] = request.Caption
	}

	if !dara.IsNil(request.CustomSql) {
		query["CustomSql"] = request.CustomSql
	}

	if !dara.IsNil(request.DsId) {
		query["DsId"] = request.DsId
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
		Action:      dara.String("CreateCubeBySql"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCubeBySqlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据物理表名称创建数据集
//
// @param request - CreateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithContext(ctx context.Context, request *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DsId) {
		query["DsId"] = request.DsId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TargetDirectoryId) {
		query["TargetDirectoryId"] = request.TargetDirectoryId
	}

	if !dara.IsNil(request.UserDefineCubeName) {
		query["UserDefineCubeName"] = request.UserDefineCubeName
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
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithContext(ctx context.Context, request *CreateTicketRequest, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicket4CopilotResponse
func (client *Client) CreateTicket4CopilotWithContext(ctx context.Context, request *CreateTicket4CopilotRequest, runtime *dara.RuntimeOptions) (_result *CreateTicket4CopilotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithContext(ctx context.Context, request *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - CreateWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithContext(ctx context.Context, request *CreateWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowPublish) {
		query["AllowPublish"] = request.AllowPublish
	}

	if !dara.IsNil(request.AllowShare) {
		query["AllowShare"] = request.AllowShare
	}

	if !dara.IsNil(request.AllowViewAll) {
		query["AllowViewAll"] = request.AllowViewAll
	}

	if !dara.IsNil(request.DefaultShareToAll) {
		query["DefaultShareToAll"] = request.DefaultShareToAll
	}

	if !dara.IsNil(request.OnlyAdminCreateDatasource) {
		query["OnlyAdminCreateDatasource"] = request.OnlyAdminCreateDatasource
	}

	if !dara.IsNil(request.UseComment) {
		query["UseComment"] = request.UseComment
	}

	if !dara.IsNil(request.WorkspaceDescription) {
		query["WorkspaceDescription"] = request.WorkspaceDescription
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataInterpretationResponse
func (client *Client) DataInterpretationWithContext(ctx context.Context, request *DataInterpretationRequest, runtime *dara.RuntimeOptions) (_result *DataInterpretationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataSetBloodResponse
func (client *Client) DataSetBloodWithContext(ctx context.Context, request *DataSetBloodRequest, runtime *dara.RuntimeOptions) (_result *DataSetBloodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataSourceBloodResponse
func (client *Client) DataSourceBloodWithContext(ctx context.Context, request *DataSourceBloodRequest, runtime *dara.RuntimeOptions) (_result *DataSourceBloodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelayTicketExpireTimeResponse
func (client *Client) DelayTicketExpireTimeWithContext(ctx context.Context, request *DelayTicketExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *DelayTicketExpireTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLevelPermissionRuleUsersResponse
func (client *Client) DeleteDataLevelPermissionRuleUsersWithContext(ctx context.Context, request *DeleteDataLevelPermissionRuleUsersRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLevelPermissionRuleUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLevelRuleConfigResponse
func (client *Client) DeleteDataLevelRuleConfigWithContext(ctx context.Context, request *DeleteDataLevelRuleConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLevelRuleConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTicketResponse
func (client *Client) DeleteTicketWithContext(ctx context.Context, request *DeleteTicketRequest, runtime *dara.RuntimeOptions) (_result *DeleteTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserFromWorkspaceResponse
func (client *Client) DeleteUserFromWorkspaceWithContext(ctx context.Context, request *DeleteUserFromWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserFromWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithContext(ctx context.Context, request *DeleteUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupMemberResponse
func (client *Client) DeleteUserGroupMemberWithContext(ctx context.Context, request *DeleteUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupMembersResponse
func (client *Client) DeleteUserGroupMembersWithContext(ctx context.Context, request *DeleteUserGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserTagMetaResponse
func (client *Client) DeleteUserTagMetaWithContext(ctx context.Context, request *DeleteUserTagMetaRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserTagMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceConnectionInfoResponse
func (client *Client) GetDataSourceConnectionInfoWithContext(ctx context.Context, request *GetDataSourceConnectionInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceConnectionInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取订阅任务列表信息
//
// @param request - GetMailTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMailTaskListResponse
func (client *Client) GetMailTaskListWithContext(ctx context.Context, request *GetMailTaskListRequest, runtime *dara.RuntimeOptions) (_result *GetMailTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Paused) {
		query["Paused"] = request.Paused
	}

	if !dara.IsNil(request.UserNick) {
		query["UserNick"] = request.UserNick
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMailTaskList"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMailTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMailTaskStatusResponse
func (client *Client) GetMailTaskStatusWithContext(ctx context.Context, request *GetMailTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetMailTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupInfoResponse
func (client *Client) GetUserGroupInfoWithContext(ctx context.Context, request *GetUserGroupInfoRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorksEmbedListResponse
func (client *Client) GetWorksEmbedListWithContext(ctx context.Context, request *GetWorksEmbedListRequest, runtime *dara.RuntimeOptions) (_result *GetWorksEmbedListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取空间下加速引擎管控页任务信息。
//
// @param request - ListAccelerationOfWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccelerationOfWorkspaceResponse
func (client *Client) ListAccelerationOfWorkspaceWithContext(ctx context.Context, request *ListAccelerationOfWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *ListAccelerationOfWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreatorId) {
		query["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.CubeName) {
		query["CubeName"] = request.CubeName
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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
		Action:      dara.String("ListAccelerationOfWorkspace"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccelerationOfWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiDatasourceResponse
func (client *Client) ListApiDatasourceWithContext(ctx context.Context, request *ListApiDatasourceRequest, runtime *dara.RuntimeOptions) (_result *ListApiDatasourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListByUserGroupIdResponse
func (client *Client) ListByUserGroupIdWithContext(ctx context.Context, request *ListByUserGroupIdRequest, runtime *dara.RuntimeOptions) (_result *ListByUserGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of works that a user has favorited.
//
// @param request - ListCollectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCollectionsResponse
func (client *Client) ListCollectionsWithContext(ctx context.Context, request *ListCollectionsRequest, runtime *dara.RuntimeOptions) (_result *ListCollectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCubeDataLevelPermissionConfigResponse
func (client *Client) ListCubeDataLevelPermissionConfigWithContext(ctx context.Context, request *ListCubeDataLevelPermissionConfigRequest, runtime *dara.RuntimeOptions) (_result *ListCubeDataLevelPermissionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLevelPermissionWhiteListResponse
func (client *Client) ListDataLevelPermissionWhiteListWithContext(ctx context.Context, request *ListDataLevelPermissionWhiteListRequest, runtime *dara.RuntimeOptions) (_result *ListDataLevelPermissionWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceResponse
func (client *Client) ListDataSourceWithContext(ctx context.Context, request *ListDataSourceRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFavoriteReportsResponse
func (client *Client) ListFavoriteReportsWithContext(ctx context.Context, request *ListFavoriteReportsRequest, runtime *dara.RuntimeOptions) (_result *ListFavoriteReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationRoleUsersResponse
func (client *Client) ListOrganizationRoleUsersWithContext(ctx context.Context, request *ListOrganizationRoleUsersRequest, runtime *dara.RuntimeOptions) (_result *ListOrganizationRoleUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPortalMenuAuthorizationResponse
func (client *Client) ListPortalMenuAuthorizationWithContext(ctx context.Context, request *ListPortalMenuAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *ListPortalMenuAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPortalMenusResponse
func (client *Client) ListPortalMenusWithContext(ctx context.Context, request *ListPortalMenusRequest, runtime *dara.RuntimeOptions) (_result *ListPortalMenusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentViewReportsResponse
func (client *Client) ListRecentViewReportsWithContext(ctx context.Context, request *ListRecentViewReportsRequest, runtime *dara.RuntimeOptions) (_result *ListRecentViewReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSharedReportsResponse
func (client *Client) ListSharedReportsWithContext(ctx context.Context, request *ListSharedReportsRequest, runtime *dara.RuntimeOptions) (_result *ListSharedReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsByUserIdResponse
func (client *Client) ListUserGroupsByUserIdWithContext(ctx context.Context, request *ListUserGroupsByUserIdRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据门户菜单的白名单列表
//
// @param request - ListWhitePortalMenuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWhitePortalMenuResponse
func (client *Client) ListWhitePortalMenuWithContext(ctx context.Context, request *ListWhitePortalMenuRequest, runtime *dara.RuntimeOptions) (_result *ListWhitePortalMenuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataportalId) {
		query["DataportalId"] = request.DataportalId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWhitePortalMenu"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWhitePortalMenuResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceRoleUsersResponse
func (client *Client) ListWorkspaceRoleUsersWithContext(ctx context.Context, request *ListWorkspaceRoleUsersRequest, runtime *dara.RuntimeOptions) (_result *ListWorkspaceRoleUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceRolesResponse
func (client *Client) ListWorkspaceRolesWithContext(ctx context.Context, request *ListWorkspaceRolesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkspaceRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户所有空间角色列表
//
// @param request - ListWorkspaceUserRolesByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceUserRolesByUserIdResponse
func (client *Client) ListWorkspaceUserRolesByUserIdWithContext(ctx context.Context, request *ListWorkspaceUserRolesByUserIdRequest, runtime *dara.RuntimeOptions) (_result *ListWorkspaceUserRolesByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaceUserRolesByUserId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspaceUserRolesByUserIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualRunMailTaskResponse
func (client *Client) ManualRunMailTaskWithContext(ctx context.Context, request *ManualRunMailTaskRequest, runtime *dara.RuntimeOptions) (_result *ManualRunMailTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiDatasourceParametersResponse
func (client *Client) ModifyApiDatasourceParametersWithContext(ctx context.Context, request *ModifyApiDatasourceParametersRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiDatasourceParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCopilotEmbedConfigResponse
func (client *Client) ModifyCopilotEmbedConfigWithContext(ctx context.Context, request *ModifyCopilotEmbedConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyCopilotEmbedConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量编辑仪表板的小Q问数状态
//
// @param request - ModifyDashboardNl2sqlStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDashboardNl2sqlStatusResponse
func (client *Client) ModifyDashboardNl2sqlStatusWithContext(ctx context.Context, request *ModifyDashboardNl2sqlStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyDashboardNl2sqlStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DashboardIds) {
		query["DashboardIds"] = request.DashboardIds
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDashboardNl2sqlStatus"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDashboardNl2sqlStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定数据集的加速任务运行日志
//
// @param request - QueryAccelerationLogByCubeIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccelerationLogByCubeIdResponse
func (client *Client) QueryAccelerationLogByCubeIdWithContext(ctx context.Context, request *QueryAccelerationLogByCubeIdRequest, runtime *dara.RuntimeOptions) (_result *QueryAccelerationLogByCubeIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccelerationLogByCubeId"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccelerationLogByCubeIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApprovalInfoResponse
func (client *Client) QueryApprovalInfoWithContext(ctx context.Context, request *QueryApprovalInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryApprovalInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAuditLogResponse
func (client *Client) QueryAuditLogWithContext(ctx context.Context, request *QueryAuditLogRequest, runtime *dara.RuntimeOptions) (_result *QueryAuditLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryComponentPerformanceResponse
func (client *Client) QueryComponentPerformanceWithContext(ctx context.Context, request *QueryComponentPerformanceRequest, runtime *dara.RuntimeOptions) (_result *QueryComponentPerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCopilotEmbedConfigResponse
func (client *Client) QueryCopilotEmbedConfigWithContext(ctx context.Context, request *QueryCopilotEmbedConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryCopilotEmbedConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCubeOptimizationResponse
func (client *Client) QueryCubeOptimizationWithContext(ctx context.Context, request *QueryCubeOptimizationRequest, runtime *dara.RuntimeOptions) (_result *QueryCubeOptimizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCubePerformanceResponse
func (client *Client) QueryCubePerformanceWithContext(ctx context.Context, request *QueryCubePerformanceRequest, runtime *dara.RuntimeOptions) (_result *QueryCubePerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Dashboard\\"s Question Resource Information
//
// @param request - QueryDashboardNl2sqlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDashboardNl2sqlResponse
func (client *Client) QueryDashboardNl2sqlWithContext(ctx context.Context, request *QueryDashboardNl2sqlRequest, runtime *dara.RuntimeOptions) (_result *QueryDashboardNl2sqlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("QueryDashboardNl2sql"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDashboardNl2sqlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataResponse
func (client *Client) QueryDataWithContext(ctx context.Context, request *QueryDataRequest, runtime *dara.RuntimeOptions) (_result *QueryDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataRangeResponse
func (client *Client) QueryDataRangeWithContext(ctx context.Context, request *QueryDataRangeRequest, runtime *dara.RuntimeOptions) (_result *QueryDataRangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataServiceResponse
func (client *Client) QueryDataServiceWithContext(ctx context.Context, request *QueryDataServiceRequest, runtime *dara.RuntimeOptions) (_result *QueryDataServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataServiceListResponse
func (client *Client) QueryDataServiceListWithContext(ctx context.Context, request *QueryDataServiceListRequest, runtime *dara.RuntimeOptions) (_result *QueryDataServiceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetDetailInfoResponse
func (client *Client) QueryDatasetDetailInfoWithContext(ctx context.Context, request *QueryDatasetDetailInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetDetailInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetInfoResponse
func (client *Client) QueryDatasetInfoWithContext(ctx context.Context, request *QueryDatasetInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetListResponse
func (client *Client) QueryDatasetListWithContext(ctx context.Context, request *QueryDatasetListRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetSmartqStatusResponse
func (client *Client) QueryDatasetSmartqStatusWithContext(ctx context.Context, request *QueryDatasetSmartqStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetSmartqStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the row-level permission switch status for a specified dataset.
//
// Description:
//
// > This interface only supports the new row and column permission model of Quick BI. If you are still using the old row and column permissions, please migrate to the new row and column permission model before calling this interface. To migrate to the new row and column permission model, follow these steps: In Organization Management -> Security Configuration -> Upgrade Row and Column Permissions, click **One-Click Upgrade*	- to upgrade to the new row-level permissions.
//
// @param request - QueryDatasetSwitchInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDatasetSwitchInfoResponse
func (client *Client) QueryDatasetSwitchInfoWithContext(ctx context.Context, request *QueryDatasetSwitchInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDatasetSwitchInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEmbeddedStatusResponse
func (client *Client) QueryEmbeddedStatusWithContext(ctx context.Context, request *QueryEmbeddedStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryEmbeddedStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the Most Recent Acceleration Task by Dataset ID
//
// @param request - QueryLastAccelerationEngineJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLastAccelerationEngineJobResponse
func (client *Client) QueryLastAccelerationEngineJobWithContext(ctx context.Context, request *QueryLastAccelerationEngineJobRequest, runtime *dara.RuntimeOptions) (_result *QueryLastAccelerationEngineJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLastAccelerationEngineJob"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLastAccelerationEngineJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLlmCubeWithThemeListByUserIdResponse
func (client *Client) QueryLlmCubeWithThemeListByUserIdWithContext(ctx context.Context, request *QueryLlmCubeWithThemeListByUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryLlmCubeWithThemeListByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrganizationRoleConfigResponse
func (client *Client) QueryOrganizationRoleConfigWithContext(ctx context.Context, request *QueryOrganizationRoleConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryOrganizationRoleConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrganizationWorkspaceListResponse
func (client *Client) QueryOrganizationWorkspaceListWithContext(ctx context.Context, request *QueryOrganizationWorkspaceListRequest, runtime *dara.RuntimeOptions) (_result *QueryOrganizationWorkspaceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReadableResourcesListByUserIdResponse
func (client *Client) QueryReadableResourcesListByUserIdWithContext(ctx context.Context, request *QueryReadableResourcesListByUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryReadableResourcesListByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query list of works user has permission to view (new)
//
// @param request - QueryReadableResourcesListByUserIdV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReadableResourcesListByUserIdV2Response
func (client *Client) QueryReadableResourcesListByUserIdV2WithContext(ctx context.Context, request *QueryReadableResourcesListByUserIdV2Request, runtime *dara.RuntimeOptions) (_result *QueryReadableResourcesListByUserIdV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReportPerformanceResponse
func (client *Client) QueryReportPerformanceWithContext(ctx context.Context, request *QueryReportPerformanceRequest, runtime *dara.RuntimeOptions) (_result *QueryReportPerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryShareListResponse
func (client *Client) QueryShareListWithContext(ctx context.Context, request *QueryShareListRequest, runtime *dara.RuntimeOptions) (_result *QueryShareListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySharesToUserListResponse
func (client *Client) QuerySharesToUserListWithContext(ctx context.Context, request *QuerySharesToUserListRequest, runtime *dara.RuntimeOptions) (_result *QuerySharesToUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmartqPermissionByCubeIdResponse
func (client *Client) QuerySmartqPermissionByCubeIdWithContext(ctx context.Context, request *QuerySmartqPermissionByCubeIdRequest, runtime *dara.RuntimeOptions) (_result *QuerySmartqPermissionByCubeIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTicketInfoResponse
func (client *Client) QueryTicketInfoWithContext(ctx context.Context, request *QueryTicketInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryTicketInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据绑定的第三方账号ID查询UserId
//
// @param request - QueryUserByMobileAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserByMobileAccountResponse
func (client *Client) QueryUserByMobileAccountWithContext(ctx context.Context, request *QueryUserByMobileAccountRequest, runtime *dara.RuntimeOptions) (_result *QueryUserByMobileAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MobileType) {
		query["MobileType"] = request.MobileType
	}

	if !dara.IsNil(request.MobileUserId) {
		query["MobileUserId"] = request.MobileUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserByMobileAccount"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserByMobileAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserGroupListByParentIdResponse
func (client *Client) QueryUserGroupListByParentIdWithContext(ctx context.Context, request *QueryUserGroupListByParentIdRequest, runtime *dara.RuntimeOptions) (_result *QueryUserGroupListByParentIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserGroupMemberResponse
func (client *Client) QueryUserGroupMemberWithContext(ctx context.Context, request *QueryUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *QueryUserGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserInfoByAccountResponse
func (client *Client) QueryUserInfoByAccountWithContext(ctx context.Context, request *QueryUserInfoByAccountRequest, runtime *dara.RuntimeOptions) (_result *QueryUserInfoByAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserInfoByUserIdResponse
func (client *Client) QueryUserInfoByUserIdWithContext(ctx context.Context, request *QueryUserInfoByUserIdRequest, runtime *dara.RuntimeOptions) (_result *QueryUserInfoByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserListResponse
func (client *Client) QueryUserListWithContext(ctx context.Context, request *QueryUserListRequest, runtime *dara.RuntimeOptions) (_result *QueryUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserRoleInfoInWorkspaceResponse
func (client *Client) QueryUserRoleInfoInWorkspaceWithContext(ctx context.Context, request *QueryUserRoleInfoInWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *QueryUserRoleInfoInWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserTagValueListResponse
func (client *Client) QueryUserTagValueListWithContext(ctx context.Context, request *QueryUserTagValueListRequest, runtime *dara.RuntimeOptions) (_result *QueryUserTagValueListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksResponse
func (client *Client) QueryWorksWithContext(ctx context.Context, request *QueryWorksRequest, runtime *dara.RuntimeOptions) (_result *QueryWorksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksBloodRelationshipResponse
func (client *Client) QueryWorksBloodRelationshipWithContext(ctx context.Context, request *QueryWorksBloodRelationshipRequest, runtime *dara.RuntimeOptions) (_result *QueryWorksBloodRelationshipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksByOrganizationResponse
func (client *Client) QueryWorksByOrganizationWithContext(ctx context.Context, request *QueryWorksByOrganizationRequest, runtime *dara.RuntimeOptions) (_result *QueryWorksByOrganizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorksByWorkspaceResponse
func (client *Client) QueryWorksByWorkspaceWithContext(ctx context.Context, request *QueryWorksByWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *QueryWorksByWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorkspaceRoleConfigResponse
func (client *Client) QueryWorkspaceRoleConfigWithContext(ctx context.Context, request *QueryWorkspaceRoleConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryWorkspaceRoleConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryWorkspaceUserListResponse
func (client *Client) QueryWorkspaceUserListWithContext(ctx context.Context, request *QueryWorkspaceUserListRequest, runtime *dara.RuntimeOptions) (_result *QueryWorkspaceUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResultCallbackResponse
func (client *Client) ResultCallbackWithContext(ctx context.Context, request *ResultCallbackRequest, runtime *dara.RuntimeOptions) (_result *ResultCallbackResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveFavoritesResponse
func (client *Client) SaveFavoritesWithContext(ctx context.Context, request *SaveFavoritesRequest, runtime *dara.RuntimeOptions) (_result *SaveFavoritesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set extra configuration for row and column permissions.
//
// Description:
//
// > This interface only supports the new version of Quick BI\\"s row and column permission model. If you are still using the old row and column permissions, please migrate to the new row and column permission model before calling this interface. The steps to migrate to the new row and column permission model: In Organization Management --> Security Configuration --> Upgrade Row and Column Permissions to New Version, click **One-Click Upgrade*	- to upgrade to the new row-level permissions.
//
// @param request - SetDataLevelPermissionExtraConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionExtraConfigResponse
func (client *Client) SetDataLevelPermissionExtraConfigWithContext(ctx context.Context, request *SetDataLevelPermissionExtraConfigRequest, runtime *dara.RuntimeOptions) (_result *SetDataLevelPermissionExtraConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set single dataset row and column permission configuration information (addition and update)
//
// Description:
//
// > This interface only supports the new row and column permission model of Quick BI. If you are still using the old row and column permissions, please migrate to the new row and column permission model before calling this interface. Steps to migrate to the new row and column permission model: In Organization Management --> Security Configuration --> Upgrade Row and Column Permissions to New Version, click **One-Click Upgrade*	- to upgrade to the new row-level permissions.
//
// @param request - SetDataLevelPermissionRuleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionRuleConfigResponse
func (client *Client) SetDataLevelPermissionRuleConfigWithContext(ctx context.Context, request *SetDataLevelPermissionRuleConfigRequest, runtime *dara.RuntimeOptions) (_result *SetDataLevelPermissionRuleConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataLevelPermissionWhiteListResponse
func (client *Client) SetDataLevelPermissionWhiteListWithContext(ctx context.Context, request *SetDataLevelPermissionWhiteListRequest, runtime *dara.RuntimeOptions) (_result *SetDataLevelPermissionWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartqAuthTransferResponse
func (client *Client) SmartqAuthTransferWithContext(ctx context.Context, request *SmartqAuthTransferRequest, runtime *dara.RuntimeOptions) (_result *SmartqAuthTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartqAuthorizeResponse
func (client *Client) SmartqAuthorizeWithContext(ctx context.Context, request *SmartqAuthorizeRequest, runtime *dara.RuntimeOptions) (_result *SmartqAuthorizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartqQueryAbilityResponse
func (client *Client) SmartqQueryAbilityWithContext(ctx context.Context, request *SmartqQueryAbilityRequest, runtime *dara.RuntimeOptions) (_result *SmartqQueryAbilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.MultipleCubeIds) {
		query["MultipleCubeIds"] = request.MultipleCubeIds
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新自定义sql数据集
//
// @param request - UpdateCubeBySqlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCubeBySqlResponse
func (client *Client) UpdateCubeBySqlWithContext(ctx context.Context, request *UpdateCubeBySqlRequest, runtime *dara.RuntimeOptions) (_result *UpdateCubeBySqlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CubeId) {
		query["CubeId"] = request.CubeId
	}

	if !dara.IsNil(request.CustomSql) {
		query["CustomSql"] = request.CustomSql
	}

	if !dara.IsNil(request.DsId) {
		query["DsId"] = request.DsId
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
		Action:      dara.String("UpdateCubeBySql"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCubeBySqlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLevelPermissionStatusResponse
func (client *Client) UpdateDataLevelPermissionStatusWithContext(ctx context.Context, request *UpdateDataLevelPermissionStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLevelPermissionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Data Source Configuration
//
// @param request - UpdateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSourceWithContext(ctx context.Context, request *UpdateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UpdateModel) {
		query["UpdateModel"] = request.UpdateModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSource"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEmbeddedStatusResponse
func (client *Client) UpdateEmbeddedStatusWithContext(ctx context.Context, request *UpdateEmbeddedStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateEmbeddedStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTicketNumResponse
func (client *Client) UpdateTicketNumWithContext(ctx context.Context, request *UpdateTicketNumRequest, runtime *dara.RuntimeOptions) (_result *UpdateTicketNumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithContext(ctx context.Context, request *UpdateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserTagMetaResponse
func (client *Client) UpdateUserTagMetaWithContext(ctx context.Context, request *UpdateUserTagMetaRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserTagMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserTagValueResponse
func (client *Client) UpdateUserTagValueWithContext(ctx context.Context, request *UpdateUserTagValueRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserTagValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceUserRoleResponse
func (client *Client) UpdateWorkspaceUserRoleWithContext(ctx context.Context, request *UpdateWorkspaceUserRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceUserRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceUsersRoleResponse
func (client *Client) UpdateWorkspaceUsersRoleWithContext(ctx context.Context, request *UpdateWorkspaceUsersRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceUsersRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawAllUserGroupsResponse
func (client *Client) WithdrawAllUserGroupsWithContext(ctx context.Context, request *WithdrawAllUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *WithdrawAllUserGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
