// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds a member to a group.
//
// @param request - AddGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupMemberResponse
func (client *Client) AddGroupMemberWithContext(ctx context.Context, request *AddGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	if !dara.IsNil(request.MemberId) {
		body["member_id"] = request.MemberId
	}

	if !dara.IsNil(request.MemberType) {
		body["member_type"] = request.MemberType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGroupMember"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/add_member"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGroupMemberResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 故事添加文件
//
// @param request - AddStoryFilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddStoryFilesResponse
func (client *Client) AddStoryFilesWithContext(ctx context.Context, request *AddStoryFilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddStoryFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	if !dara.IsNil(request.StoryId) {
		body["story_id"] = request.StoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddStoryFiles"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/add_story_files"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddStoryFilesResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns a group administrator role to a user.
//
// Description:
//
// You can call this operation to assign a group administrator role to a user.
//
// @param request - AssignRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignRoleResponse
func (client *Client) AssignRoleWithContext(ctx context.Context, request *AssignRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AssignRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Identity) {
		body["identity"] = request.Identity
	}

	if !dara.IsNil(request.ManageResourceId) {
		body["manage_resource_id"] = request.ManageResourceId
	}

	if !dara.IsNil(request.ManageResourceType) {
		body["manage_resource_type"] = request.ManageResourceType
	}

	if !dara.IsNil(request.RoleId) {
		body["role_id"] = request.RoleId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignRole"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/role/assign"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignRoleResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports audit logs.
//
// Description:
//
// Log audit is a value-added feature that is provided by Drive and Photo Service (PDS) Developer Edition. Before you call this operation, make sure that you learn about the [value-added billable items](https://www.alibabacloud.com/help/document_detail/425220.html).
//
// @param request - AuditLogExportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuditLogExportResponse
func (client *Client) AuditLogExportWithContext(ctx context.Context, request *AuditLogExportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuditLogExportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["file_name"] = request.FileName
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.OrderBy) {
		body["order_by"] = request.OrderBy
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuditLogExport"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/audit_log/export"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuditLogExportResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Requests permissions by using OAuth 2.0.
//
// Description:
//
// For more information, see "OAuth 2.0 For Web Server Applications" at [OAuth 2.0 For Web Server Applications](https://www.alibabacloud.com/help/en/pds/drive-and-photo-service-dev/user-guide/oauth-2-0-access-process-for-web-server-applications) in User Guide.
//
// @param tmpReq - AuthorizeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeResponse
func (client *Client) AuthorizeWithContext(ctx context.Context, tmpReq *AuthorizeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuthorizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AuthorizeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Scope) {
		request.ScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scope, dara.String("scope"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["client_id"] = request.ClientId
	}

	if !dara.IsNil(request.HideConsent) {
		query["hide_consent"] = request.HideConsent
	}

	if !dara.IsNil(request.LoginType) {
		query["login_type"] = request.LoginType
	}

	if !dara.IsNil(request.RedirectUri) {
		query["redirect_uri"] = request.RedirectUri
	}

	if !dara.IsNil(request.ResponseType) {
		query["response_type"] = request.ResponseType
	}

	if !dara.IsNil(request.ScopeShrink) {
		query["scope"] = request.ScopeShrink
	}

	if !dara.IsNil(request.State) {
		query["state"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Authorize"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/oauth/authorize"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("binary"),
	}
	_result = &AuthorizeResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls multiple operations at a time to improve call efficiency.
//
// @param request - BatchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchResponse
func (client *Client) BatchWithContext(ctx context.Context, request *BatchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Requests) {
		body["requests"] = request.Requests
	}

	if !dara.IsNil(request.Resource) {
		body["resource"] = request.Resource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Batch"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/batch"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a role.
//
// Description:
//
// You can cancel only the group administrator role.
//
// @param request - CancelAssignRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAssignRoleResponse
func (client *Client) CancelAssignRoleWithContext(ctx context.Context, request *CancelAssignRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelAssignRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Identity) {
		body["identity"] = request.Identity
	}

	if !dara.IsNil(request.ManageResourceId) {
		body["manage_resource_id"] = request.ManageResourceId
	}

	if !dara.IsNil(request.ManageResourceType) {
		body["manage_resource_type"] = request.ManageResourceType
	}

	if !dara.IsNil(request.RoleId) {
		body["role_id"] = request.RoleId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAssignRole"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/role/cancel_assign"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAssignRoleResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a share link.
//
// @param request - CancelShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelShareLinkResponse
func (client *Client) CancelShareLinkWithContext(ctx context.Context, request *CancelShareLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelShareLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelShareLink"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/share_link/cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelShareLinkResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Empties the recycle bin.
//
// @param request - ClearRecyclebinRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearRecyclebinResponse
func (client *Client) ClearRecyclebinWithContext(ctx context.Context, request *ClearRecyclebinRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ClearRecyclebinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearRecyclebin"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/recyclebin/clear"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearRecyclebinResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Completes the upload of a file.
//
// @param request - CompleteFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteFileResponse
func (client *Client) CompleteFileWithContext(ctx context.Context, request *CompleteFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CompleteFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Crc64Hash) {
		body["crc64_hash"] = request.Crc64Hash
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.UploadId) {
		body["upload_id"] = request.UploadId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/complete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a file or folder.
//
// @param request - CopyFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyFileResponse
func (client *Client) CopyFileWithContext(ctx context.Context, request *CopyFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CopyFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRename) {
		body["auto_rename"] = request.AutoRename
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.ToDriveId) {
		body["to_drive_id"] = request.ToDriveId
	}

	if !dara.IsNil(request.ToParentFileId) {
		body["to_parent_file_id"] = request.ToParentFileId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/copy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义故事
//
// @param request - CreateCustomizedStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomizedStoryResponse
func (client *Client) CreateCustomizedStoryWithContext(ctx context.Context, request *CreateCustomizedStoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCustomizedStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomLabels) {
		body["custom_labels"] = request.CustomLabels
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.StoryCover) {
		body["story_cover"] = request.StoryCover
	}

	if !dara.IsNil(request.StoryFiles) {
		body["story_files"] = request.StoryFiles
	}

	if !dara.IsNil(request.StoryName) {
		body["story_name"] = request.StoryName
	}

	if !dara.IsNil(request.StorySubType) {
		body["story_sub_type"] = request.StorySubType
	}

	if !dara.IsNil(request.StoryType) {
		body["story_type"] = request.StoryType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomizedStory"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/create_customized_story"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomizedStoryResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a domain.
//
// Description:
//
// The description of the domain.
//
// @param request - CreateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithContext(ctx context.Context, request *CreateDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
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

	if !dara.IsNil(request.DomainName) {
		body["domain_name"] = request.DomainName
	}

	if !dara.IsNil(request.InitDriveEnable) {
		body["init_drive_enable"] = request.InitDriveEnable
	}

	if !dara.IsNil(request.InitDriveSize) {
		body["init_drive_size"] = request.InitDriveSize
	}

	if !dara.IsNil(request.ParentDomainId) {
		body["parent_domain_id"] = request.ParentDomainId
	}

	if !dara.IsNil(request.SizeQuota) {
		body["size_quota"] = request.SizeQuota
	}

	if !dara.IsNil(request.StoreRedundancyType) {
		body["store_redundancy_type"] = request.StoreRedundancyType
	}

	if !dara.IsNil(request.UserCountQuota) {
		body["user_count_quota"] = request.UserCountQuota
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomain"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a drive.
//
// @param request - CreateDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDriveResponse
func (client *Client) CreateDriveWithContext(ctx context.Context, request *CreateDriveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Default) {
		body["default"] = request.Default
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DriveName) {
		body["drive_name"] = request.DriveName
	}

	if !dara.IsNil(request.DriveType) {
		body["drive_type"] = request.DriveType
	}

	if !dara.IsNil(request.Owner) {
		body["owner"] = request.Owner
	}

	if !dara.IsNil(request.OwnerType) {
		body["owner_type"] = request.OwnerType
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.TotalSize) {
		body["total_size"] = request.TotalSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDrive"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDriveResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file or folder.
//
// @param request - CreateFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithContext(ctx context.Context, request *CreateFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckNameMode) {
		body["check_name_mode"] = request.CheckNameMode
	}

	if !dara.IsNil(request.ContentHash) {
		body["content_hash"] = request.ContentHash
	}

	if !dara.IsNil(request.ContentHashName) {
		body["content_hash_name"] = request.ContentHashName
	}

	if !dara.IsNil(request.ContentType) {
		body["content_type"] = request.ContentType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.Hidden) {
		body["hidden"] = request.Hidden
	}

	if !dara.IsNil(request.LocalCreatedAt) {
		body["local_created_at"] = request.LocalCreatedAt
	}

	if !dara.IsNil(request.LocalModifiedAt) {
		body["local_modified_at"] = request.LocalModifiedAt
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ParallelUpload) {
		body["parallel_upload"] = request.ParallelUpload
	}

	if !dara.IsNil(request.ParentFileId) {
		body["parent_file_id"] = request.ParentFileId
	}

	if !dara.IsNil(request.PartInfoList) {
		body["part_info_list"] = request.PartInfoList
	}

	if !dara.IsNil(request.PreHash) {
		body["pre_hash"] = request.PreHash
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.Size) {
		body["size"] = request.Size
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.UserTags) {
		body["user_tags"] = request.UserTags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
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

	if !dara.IsNil(request.GroupName) {
		body["group_name"] = request.GroupName
	}

	if !dara.IsNil(request.IsRoot) {
		body["is_root"] = request.IsRoot
	}

	if !dara.IsNil(request.ParentGroupId) {
		body["parent_group_id"] = request.ParentGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroup"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mapping between an entity and a benefit package. You can call this operation to associate a benefit package with a user.
//
// Description:
//
// If you need to manage a large number of users based on Drive and Photo Service, you can control the features and quotas that users can use based on the benefits to which they are entitled. For more information, join the DingTalk group (ID 23146118).
//
// @param request - CreateIdentityToBenefitPkgMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityToBenefitPkgMappingResponse
func (client *Client) CreateIdentityToBenefitPkgMappingWithContext(ctx context.Context, request *CreateIdentityToBenefitPkgMappingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIdentityToBenefitPkgMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		body["amount"] = request.Amount
	}

	if !dara.IsNil(request.BenefitPkgId) {
		body["benefit_pkg_id"] = request.BenefitPkgId
	}

	if !dara.IsNil(request.ExpireTime) {
		body["expire_time"] = request.ExpireTime
	}

	if !dara.IsNil(request.IdentityId) {
		body["identity_id"] = request.IdentityId
	}

	if !dara.IsNil(request.IdentityType) {
		body["identity_type"] = request.IdentityType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdentityToBenefitPkgMapping"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/benefit/identity_to_benefit_pkg_mapping/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建凌霄订单
//
// @param request - CreateOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderResponse
func (client *Client) CreateOrderWithContext(ctx context.Context, request *CreateOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		body["auto_pay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		body["auto_renew"] = request.AutoRenew
	}

	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.InstanceId) {
		body["instance_id"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderType) {
		body["order_type"] = request.OrderType
	}

	if !dara.IsNil(request.Package) {
		body["package"] = request.Package
	}

	if !dara.IsNil(request.Period) {
		body["period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["period_unit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.TotalSize) {
		body["total_size"] = request.TotalSize
	}

	if !dara.IsNil(request.UserCount) {
		body["user_count"] = request.UserCount
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrder"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/create_order"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a share URL.
//
// Description:
//
// A share is a file view container. You can grant anonymous users the permissions to access files in the user drive by using the share URL. Anonymous users can access the files based on the granted permissions.
//
// @param request - CreateShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateShareLinkResponse
func (client *Client) CreateShareLinkWithContext(ctx context.Context, request *CreateShareLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateShareLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Creatable) {
		body["creatable"] = request.Creatable
	}

	if !dara.IsNil(request.CreatableFileIdList) {
		body["creatable_file_id_list"] = request.CreatableFileIdList
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisableDownload) {
		body["disable_download"] = request.DisableDownload
	}

	if !dara.IsNil(request.DisablePreview) {
		body["disable_preview"] = request.DisablePreview
	}

	if !dara.IsNil(request.DisableSave) {
		body["disable_save"] = request.DisableSave
	}

	if !dara.IsNil(request.DownloadLimit) {
		body["download_limit"] = request.DownloadLimit
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Expiration) {
		body["expiration"] = request.Expiration
	}

	if !dara.IsNil(request.FileIdList) {
		body["file_id_list"] = request.FileIdList
	}

	if !dara.IsNil(request.OfficeEditable) {
		body["office_editable"] = request.OfficeEditable
	}

	if !dara.IsNil(request.PreviewLimit) {
		body["preview_limit"] = request.PreviewLimit
	}

	if !dara.IsNil(request.RequireLogin) {
		body["require_login"] = request.RequireLogin
	}

	if !dara.IsNil(request.SaveLimit) {
		body["save_limit"] = request.SaveLimit
	}

	if !dara.IsNil(request.ShareAllFiles) {
		body["share_all_files"] = request.ShareAllFiles
	}

	if !dara.IsNil(request.ShareName) {
		body["share_name"] = request.ShareName
	}

	if !dara.IsNil(request.SharePwd) {
		body["share_pwd"] = request.SharePwd
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateShareLink"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/share_link/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateShareLinkResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建相似图片聚类任务
//
// @param request - CreateSimilarImageClusterTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimilarImageClusterTaskResponse
func (client *Client) CreateSimilarImageClusterTaskWithContext(ctx context.Context, request *CreateSimilarImageClusterTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSimilarImageClusterTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSimilarImageClusterTask"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/create_similar_image_cluster_task"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSimilarImageClusterTaskResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建推荐故事
//
// @param request - CreateStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoryResponse
func (client *Client) CreateStoryWithContext(ctx context.Context, request *CreateStoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		body["address"] = request.Address
	}

	if !dara.IsNil(request.CustomLabels) {
		body["custom_labels"] = request.CustomLabels
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.MaxImageCount) {
		body["max_image_count"] = request.MaxImageCount
	}

	if !dara.IsNil(request.MinImageCount) {
		body["min_image_count"] = request.MinImageCount
	}

	if !dara.IsNil(request.StoryEndTime) {
		body["story_end_time"] = request.StoryEndTime
	}

	if !dara.IsNil(request.StoryId) {
		body["story_id"] = request.StoryId
	}

	if !dara.IsNil(request.StoryName) {
		body["story_name"] = request.StoryName
	}

	if !dara.IsNil(request.StoryStartTime) {
		body["story_start_time"] = request.StoryStartTime
	}

	if !dara.IsNil(request.StorySubType) {
		body["story_sub_type"] = request.StorySubType
	}

	if !dara.IsNil(request.StoryType) {
		body["story_type"] = request.StoryType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStory"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/create_story"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStoryResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user.
//
// @param request - CreateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithContext(ctx context.Context, request *CreateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Avatar) {
		body["avatar"] = request.Avatar
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.GroupInfoList) {
		body["group_info_list"] = request.GroupInfoList
	}

	if !dara.IsNil(request.NickName) {
		body["nick_name"] = request.NickName
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.Role) {
		body["role"] = request.Role
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.UserData) {
		body["user_data"] = request.UserData
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["user_name"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUser"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/user/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件内容安全信息
//
// @param request - CsiGetFileInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CsiGetFileInfoResponse
func (client *Client) CsiGetFileInfoWithContext(ctx context.Context, request *CsiGetFileInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CsiGetFileInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.UrlExpireSec) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CsiGetFileInfo"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/csi/get_file_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CsiGetFileInfoResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete the domain
//
// @param request - DeleteDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithContext(ctx context.Context, request *DeleteDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		body["domain_id"] = request.DomainId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a drive.
//
// @param request - DeleteDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDriveResponse
func (client *Client) DeleteDriveWithContext(ctx context.Context, request *DeleteDriveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDrive"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDriveResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file or folder.
//
// @param request - DeleteFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithContext(ctx context.Context, request *DeleteFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes groups. Before you delete a group, make sure that no other groups or users exist in the group. Otherwise, the group fails to be deleted.
//
// @param request - DeleteGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithContext(ctx context.Context, request *DeleteGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroup"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a historical version of a file. You cannot delete the latest version of a file.
//
// @param request - DeleteRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRevisionResponse
func (client *Client) DeleteRevisionWithContext(ctx context.Context, request *DeleteRevisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRevisionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.RevisionId) {
		body["revision_id"] = request.RevisionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRevision"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/revision/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRevisionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除故事
//
// @param request - DeleteStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStoryResponse
func (client *Client) DeleteStoryWithContext(ctx context.Context, request *DeleteStoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.StoryId) {
		body["story_id"] = request.StoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStory"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/delete_story"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStoryResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
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
// @param request - DeleteUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithContext(ctx context.Context, request *DeleteUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/user/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cursor of incremental information.
//
// @param request - DeltaGetLastCursorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeltaGetLastCursorResponse
func (client *Client) DeltaGetLastCursorWithContext(ctx context.Context, request *DeltaGetLastCursorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeltaGetLastCursorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.SyncRootId) {
		body["sync_root_id"] = request.SyncRootId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeltaGetLastCursor"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/get_last_cursor"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeltaGetLastCursorResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downloads a file.
//
// Description:
//
// For information about best practices for downloading a file.
//
// @param request - DownloadFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadFileResponse
func (client *Client) DownloadFileWithContext(ctx context.Context, request *DownloadFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DownloadFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		query["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		query["file_id"] = request.FileId
	}

	if !dara.IsNil(request.ImageThumbnailProcess) {
		query["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !dara.IsNil(request.OfficeThumbnailProcess) {
		query["office_thumbnail_process"] = request.OfficeThumbnailProcess
	}

	if !dara.IsNil(request.ShareId) {
		query["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.VideoThumbnailProcess) {
		query["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/download"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("binary"),
	}
	_result = &DownloadFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions to access files to a user or group.
//
// @param request - FileAddPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileAddPermissionResponse
func (client *Client) FileAddPermissionWithContext(ctx context.Context, request *FileAddPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FileAddPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.MemberList) {
		body["member_list"] = request.MemberList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FileAddPermission"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/add_permission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FileAddPermissionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes custom tags from a file.
//
// @param request - FileDeleteUserTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileDeleteUserTagsResponse
func (client *Client) FileDeleteUserTagsWithContext(ctx context.Context, request *FileDeleteUserTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FileDeleteUserTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.KeyList) {
		body["key_list"] = request.KeyList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FileDeleteUserTags"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/delete_usertags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FileDeleteUserTagsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the sharing authorization records of a file.
//
// @param request - FileListPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileListPermissionResponse
func (client *Client) FileListPermissionWithContext(ctx context.Context, request *FileListPermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FileListPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FileListPermission"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/list_permission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &FileListPermissionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds custom tags to a file.
//
// Description:
//
// This operation is an incremental update operation. Take note of the following items:
//
//   - If a tag name specified in the request is the same as an existing tag name, the existing tag is overwritten.
//
//   - If a tag name specified in the request is different from the existing tag names, the specified tag is added.
//
//   - The existing tags with unique names are not affected.
//
// @param request - FilePutUserTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FilePutUserTagsResponse
func (client *Client) FilePutUserTagsWithContext(ctx context.Context, request *FilePutUserTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FilePutUserTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.UserTags) {
		body["user_tags"] = request.UserTags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FilePutUserTags"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/put_usertags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FilePutUserTagsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the permissions on a shared file.
//
// @param request - FileRemovePermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileRemovePermissionResponse
func (client *Client) FileRemovePermissionWithContext(ctx context.Context, request *FileRemovePermissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FileRemovePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.MemberList) {
		body["member_list"] = request.MemberList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FileRemovePermission"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/remove_permission"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FileRemovePermissionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an asynchronous task.
//
// @param request - GetAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTaskWithContext(ctx context.Context, request *GetAsyncTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AsyncTaskId) {
		body["async_task_id"] = request.AsyncTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncTask"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/async_task/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default drive of a user.
//
// @param request - GetDefaultDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultDriveResponse
func (client *Client) GetDefaultDriveWithContext(ctx context.Context, request *GetDefaultDriveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDefaultDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDefaultDrive"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/get_default_drive"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDefaultDriveResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get domain information.
//
// @param request - GetDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithContext(ctx context.Context, request *GetDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		body["domain_id"] = request.DomainId
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.GetQuotaUsed) {
		body["get_quota_used"] = request.GetQuotaUsed
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomain"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取domain限额
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainQuotaResponse
func (client *Client) GetDomainQuotaWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDomainQuotaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomainQuota"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/get_quota"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainQuotaResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the download URL of the file.
//
// @param request - GetDownloadUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDownloadUrlResponse
func (client *Client) GetDownloadUrlWithContext(ctx context.Context, request *GetDownloadUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.ExpireSec) {
		body["expire_sec"] = request.ExpireSec
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.FileName) {
		body["file_name"] = request.FileName
	}

	if !dara.IsNil(request.ResponseContentType) {
		body["response_content_type"] = request.ResponseContentType
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDownloadUrl"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/get_download_url"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDownloadUrlResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a drive.
//
// @param request - GetDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDriveResponse
func (client *Client) GetDriveWithContext(ctx context.Context, request *GetDriveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDrive"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDriveResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file.
//
// @param request - GetFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileResponse
func (client *Client) GetFileWithContext(ctx context.Context, request *GetFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.ThumbnailProcesses) {
		body["thumbnail_processes"] = request.ThumbnailProcesses
	}

	if !dara.IsNil(request.UrlExpireSec) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a group.
//
// @param request - GetGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithContext(ctx context.Context, request *GetGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroup"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mapping between an entity and a benefit package. You can call this operation to query the benefit package that is associated with a user.
//
// @param request - GetIdentityToBenefitPkgMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityToBenefitPkgMappingResponse
func (client *Client) GetIdentityToBenefitPkgMappingWithContext(ctx context.Context, request *GetIdentityToBenefitPkgMappingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIdentityToBenefitPkgMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BenefitPkgId) {
		body["benefit_pkg_id"] = request.BenefitPkgId
	}

	if !dara.IsNil(request.IdentityId) {
		body["identity_id"] = request.IdentityId
	}

	if !dara.IsNil(request.IdentityType) {
		body["identity_type"] = request.IdentityType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityToBenefitPkgMapping"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/benefit/identity_to_benefit_pkg_mapping/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an account.
//
// @param request - GetLinkInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLinkInfoResponse
func (client *Client) GetLinkInfoWithContext(ctx context.Context, request *GetLinkInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLinkInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Extra) {
		body["extra"] = request.Extra
	}

	if !dara.IsNil(request.Identity) {
		body["identity"] = request.Identity
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLinkInfo"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/account/get_link_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLinkInfoResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a user based on the user ID.
//
// @param request - GetLinkInfoByUserIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLinkInfoByUserIdResponse
func (client *Client) GetLinkInfoByUserIdWithContext(ctx context.Context, request *GetLinkInfoByUserIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLinkInfoByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLinkInfoByUserId"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/account/get_link_info_by_user_id"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLinkInfoByUserIdResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a version.
//
// @param request - GetRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRevisionResponse
func (client *Client) GetRevisionWithContext(ctx context.Context, request *GetRevisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRevisionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.RevisionId) {
		body["revision_id"] = request.RevisionId
	}

	if !dara.IsNil(request.UrlExpireSec) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRevision"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/revision/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRevisionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the share URL of a file.
//
// @param request - GetShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShareLinkResponse
func (client *Client) GetShareLinkWithContext(ctx context.Context, request *GetShareLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetShareLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetShareLink"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/share_link/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetShareLinkResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a share link anonymously.
//
// @param request - GetShareLinkByAnonymousRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShareLinkByAnonymousResponse
func (client *Client) GetShareLinkByAnonymousWithContext(ctx context.Context, request *GetShareLinkByAnonymousRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetShareLinkByAnonymousResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetShareLinkByAnonymous"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/share_link/get_by_anonymous"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetShareLinkByAnonymousResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a share token anonymously.
//
// Description:
//
// To access a file by using a share link, you must first obtain a share token, even if the value of share_pwd of this share is an empty string, which specifies that the share is not private.
//
// @param request - GetShareLinkTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShareLinkTokenResponse
func (client *Client) GetShareLinkTokenWithContext(ctx context.Context, request *GetShareLinkTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetShareLinkTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireSec) {
		body["expire_sec"] = request.ExpireSec
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.SharePwd) {
		body["share_pwd"] = request.SharePwd
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetShareLinkToken"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/share_link/get_share_token"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetShareLinkTokenResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取故事详情
//
// @param request - GetStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStoryResponse
func (client *Client) GetStoryWithContext(ctx context.Context, request *GetStoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CoverImageThumbnailProcess) {
		body["cover_image_thumbnail_process"] = request.CoverImageThumbnailProcess
	}

	if !dara.IsNil(request.CoverVideoThumbnailProcess) {
		body["cover_video_thumbnail_process"] = request.CoverVideoThumbnailProcess
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.ImageThumbnailProcess) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !dara.IsNil(request.ImageUrlProcess) {
		body["image_url_process"] = request.ImageUrlProcess
	}

	if !dara.IsNil(request.StoryId) {
		body["story_id"] = request.StoryId
	}

	if !dara.IsNil(request.UrlExpireSec) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	if !dara.IsNil(request.VideoThumbnailProcess) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStory"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/get_story"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStoryResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of a value-added asynchronous task. You can call this operation to query the execution status of an asynchronous task that is created by calling the CreateSimilarImageClusterTask operation.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/425220.html) of Drive and Photo Service**.
//
// To call this operation, make sure that the value-added image processing feature is enabled.
//
// Before you call this operation, a value-added asynchronous task must be created. For example, you can call the CreateSimilarImageClusterTask operation to create an asynchronous task. Then, you can call this operation to query the execution status of the asynchronous task based on the task ID.
//
// @param request - GetTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithContext(ctx context.Context, request *GetTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.TaskId) {
		body["task_id"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskStatus"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/get_task_status"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the upload URL of a file.
//
// @param request - GetUploadUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadUrlResponse
func (client *Client) GetUploadUrlWithContext(ctx context.Context, request *GetUploadUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.PartInfoList) {
		body["part_info_list"] = request.PartInfoList
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.UploadId) {
		body["upload_id"] = request.UploadId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadUrl"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/get_upload_url"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadUrlResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a user.
//
// @param request - GetUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/user/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about video playback.
//
// Description:
//
//	  **Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://help.aliyun.com/document_detail/425220.html) of Drive and Photo Service (PDS).**
//
//		- Before you call this operation, make sure that the transcoding mode which you want to specify by using the category parameter is enabled for the domain. To enable the transcoding feature and configure transcoding templates, contact our technical support. For more information, see [Contact us](https://help.aliyun.com/document_detail/175917.html).
//
//		- This operation is a synchronous operation. If the specified file is not transcoded in the specified transcoding mode, the API call returns **202 VideoPreviewWaitAndRetry**, which indicates that you need to wait a moment and try again. If the specified file cannot be transcoded in the specified transcoding mode, the API call returns **404 NotFound.VideoPreviewInfo**.
//
//		- This operation generates transcoding data and stores it in the space that is used to store the value-added data of the tenant domain. This way, end users can play audio and videos online. For specific transcoding modes, this operation provides domain-level deduplication for transcoding.
//
//		- If the transcoding mode is set to quick_video, the playback URL returned by this operation contains the `{` and `}` characters that are not URL-encoded. For development on iOS, decode and encode the returned URL first to avoid decoding failure of the NSURL library of the system.
//
//		- If the transcoding mode is set to quick_video, you cannot use the GET Range method to obtain segments of the M3U8 file in the playback URL.
//
// >
//
// @param request - GetVideoPreviewPlayInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoPreviewPlayInfoResponse
func (client *Client) GetVideoPreviewPlayInfoWithContext(ctx context.Context, request *GetVideoPreviewPlayInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoPreviewPlayInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.GetMasterUrl) {
		body["get_master_url"] = request.GetMasterUrl
	}

	if !dara.IsNil(request.GetWithoutUrl) {
		body["get_without_url"] = request.GetWithoutUrl
	}

	if !dara.IsNil(request.ReTranscode) {
		body["re_transcode"] = request.ReTranscode
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.TemplateId) {
		body["template_id"] = request.TemplateId
	}

	if !dara.IsNil(request.UrlExpireSec) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoPreviewPlayInfo"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/get_video_preview_play_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoPreviewPlayInfoResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the preview metadata of a video.
//
// Description:
//
// For more information about best practices, see [Preview videos online](https://help.aliyun.com/document_detail/427477.html).
//
// @param request - GetVideoPreviewPlayMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoPreviewPlayMetaResponse
func (client *Client) GetVideoPreviewPlayMetaWithContext(ctx context.Context, request *GetVideoPreviewPlayMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVideoPreviewPlayMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoPreviewPlayMeta"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/get_video_preview_play_meta"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoPreviewPlayMetaResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用户组名字
//
// @param request - GroupUpdateNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupUpdateNameResponse
func (client *Client) GroupUpdateNameWithContext(ctx context.Context, request *GroupUpdateNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GroupUpdateNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GroupUpdateName"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/update_name"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GroupUpdateNameResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a user.
//
// @param request - ImportUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportUserResponse
func (client *Client) ImportUserWithContext(ctx context.Context, request *ImportUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ImportUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationDisplayName) {
		body["authentication_display_name"] = request.AuthenticationDisplayName
	}

	if !dara.IsNil(request.AuthenticationType) {
		body["authentication_type"] = request.AuthenticationType
	}

	if !dara.IsNil(request.AutoCreateDrive) {
		body["auto_create_drive"] = request.AutoCreateDrive
	}

	if !dara.IsNil(request.DriveTotalSize) {
		body["drive_total_size"] = request.DriveTotalSize
	}

	if !dara.IsNil(request.Extra) {
		body["extra"] = request.Extra
	}

	if !dara.IsNil(request.Identity) {
		body["identity"] = request.Identity
	}

	if !dara.IsNil(request.NickName) {
		body["nick_name"] = request.NickName
	}

	if !dara.IsNil(request.ParentGroupId) {
		body["parent_group_id"] = request.ParentGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportUser"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/user/import"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportUserResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 送审文件
//
// @param request - InvestigateFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvestigateFileResponse
func (client *Client) InvestigateFileWithContext(ctx context.Context, request *InvestigateFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvestigateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveFileIds) {
		body["drive_file_ids"] = request.DriveFileIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvestigateFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/csi/investigate_file"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvestigateFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an account with a user.
//
// @param request - LinkAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LinkAccountResponse
func (client *Client) LinkAccountWithContext(ctx context.Context, request *LinkAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *LinkAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Extra) {
		body["extra"] = request.Extra
	}

	if !dara.IsNil(request.Identity) {
		body["identity"] = request.Identity
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LinkAccount"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/account/link"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &LinkAccountResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries location-based groups.
//
// @param request - ListAddressGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddressGroupsResponse
func (client *Client) ListAddressGroupsWithContext(ctx context.Context, request *ListAddressGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddressGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.ImageThumbnailProcess) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.VideoThumbnailProcess) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddressGroups"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/list_address_groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddressGroupsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of assigned roles. For example, you can query the administrators of a group by group ID.
//
// @param request - ListAssignmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssignmentResponse
func (client *Client) ListAssignmentWithContext(ctx context.Context, request *ListAssignmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAssignmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.ManageResourceId) {
		body["manage_resource_id"] = request.ManageResourceId
	}

	if !dara.IsNil(request.ManageResourceType) {
		body["manage_resource_type"] = request.ManageResourceType
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAssignment"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/role/list_assignment"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAssignmentResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries incremental information.
//
// @param request - ListDeltaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeltaResponse
func (client *Client) ListDeltaWithContext(ctx context.Context, request *ListDeltaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeltaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		body["cursor"] = request.Cursor
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.SyncRootId) {
		body["sync_root_id"] = request.SyncRootId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDelta"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/list_delta"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeltaResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举 domain
//
// @param request - ListDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithContext(ctx context.Context, request *ListDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.ParentDomainId) {
		body["parent_domain_id"] = request.ParentDomainId
	}

	if !dara.IsNil(request.ServiceCode) {
		body["service_code"] = request.ServiceCode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomains"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of drives.
//
// @param request - ListDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDriveResponse
func (client *Client) ListDriveWithContext(ctx context.Context, request *ListDriveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.Owner) {
		body["owner"] = request.Owner
	}

	if !dara.IsNil(request.OwnerType) {
		body["owner_type"] = request.OwnerType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDrive"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDriveResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries face-based groups.
//
// @param request - ListFacegroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFacegroupsResponse
func (client *Client) ListFacegroupsWithContext(ctx context.Context, request *ListFacegroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFacegroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.Remarks) {
		body["remarks"] = request.Remarks
	}

	if !dara.IsNil(request.ReturnTotalCount) {
		body["return_total_count"] = request.ReturnTotalCount
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFacegroups"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/list_facegroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFacegroupsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of files and folders.
//
// @param request - ListFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileResponse
func (client *Client) ListFileWithContext(ctx context.Context, request *ListFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.OrderBy) {
		body["order_by"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		body["order_direction"] = request.OrderDirection
	}

	if !dara.IsNil(request.ParentFileId) {
		body["parent_file_id"] = request.ParentFileId
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.ThumbnailProcesses) {
		body["thumbnail_processes"] = request.ThumbnailProcesses
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries groups.
//
// @param request - ListGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupResponse
func (client *Client) ListGroupWithContext(ctx context.Context, request *ListGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroup"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the members of a group.
//
// @param request - ListGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupMemberResponse
func (client *Client) ListGroupMemberWithContext(ctx context.Context, request *ListGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.MemberType) {
		body["member_type"] = request.MemberType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupMember"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/list_member"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupMemberResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举用户或团队已分配的角色列表
//
// @param request - ListIdentityRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityRoleResponse
func (client *Client) ListIdentityRoleWithContext(ctx context.Context, request *ListIdentityRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIdentityRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Identity) {
		body["identity"] = request.Identity
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentityRole"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/role/list_identity_role"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityRoleResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mappings between entities and benefit packages. You can call this operation to query the benefit packages that are associated with a user.
//
// @param request - ListIdentityToBenefitPkgMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityToBenefitPkgMappingResponse
func (client *Client) ListIdentityToBenefitPkgMappingWithContext(ctx context.Context, request *ListIdentityToBenefitPkgMappingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIdentityToBenefitPkgMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityId) {
		body["identity_id"] = request.IdentityId
	}

	if !dara.IsNil(request.IdentityType) {
		body["identity_type"] = request.IdentityType
	}

	if !dara.IsNil(request.IncludeExpired) {
		body["include_expired"] = request.IncludeExpired
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentityToBenefitPkgMapping"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/benefit/identity_to_benefit_pkg_mapping/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the drives of the current user.
//
// @param request - ListMyDrivesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyDrivesResponse
func (client *Client) ListMyDrivesWithContext(ctx context.Context, request *ListMyDrivesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMyDrivesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMyDrives"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/list_my_drives"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMyDrivesResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the team drives that can be accessed by the authorized users.
//
// @param request - ListMyGroupDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyGroupDriveResponse
func (client *Client) ListMyGroupDriveWithContext(ctx context.Context, request *ListMyGroupDriveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMyGroupDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveName) {
		body["drive_name"] = request.DriveName
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMyGroupDrive"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/list_my_group_drive"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMyGroupDriveResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of files that are shared with a user. You can call this operation to query a list of files in a personal drive on which a user is granted permissions.
//
// @param request - ListReceivedFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReceivedFileResponse
func (client *Client) ListReceivedFileWithContext(ctx context.Context, request *ListReceivedFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListReceivedFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReceivedFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/list_received_file"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReceivedFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about files and folders in the recycle bin.
//
// @param request - ListRecyclebinRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecyclebinResponse
func (client *Client) ListRecyclebinWithContext(ctx context.Context, request *ListRecyclebinRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRecyclebinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.ThumbnailProcesses) {
		body["thumbnail_processes"] = request.ThumbnailProcesses
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecyclebin"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/recyclebin/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecyclebinResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of a file.
//
// @param request - ListRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRevisionResponse
func (client *Client) ListRevisionWithContext(ctx context.Context, request *ListRevisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRevisionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRevision"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/revision/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRevisionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries shares.
//
// Description:
//
// This operation is discontinued. To query shares, you can call the SearchShareLink operation.
//
// @param request - ListShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShareLinkResponse
func (client *Client) ListShareLinkWithContext(ctx context.Context, request *ListShareLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListShareLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		body["creator"] = request.Creator
	}

	if !dara.IsNil(request.IncludeCancelled) {
		body["include_cancelled"] = request.IncludeCancelled
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.OrderBy) {
		body["order_by"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		body["order_direction"] = request.OrderDirection
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListShareLink"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/share_link/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListShareLinkResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags.
//
// Description:
//
// You can call this operation to query the tags within the specified drive at a time. The top 2,000 tags of the images are returned.
//
// @param request - ListTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithContext(ctx context.Context, request *ListTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.ImageThumbnailProcess) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !dara.IsNil(request.VideoThumbnailProcess) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTags"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/list_tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the file parts that are uploaded.
//
// @param request - ListUploadedPartsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUploadedPartsResponse
func (client *Client) ListUploadedPartsWithContext(ctx context.Context, request *ListUploadedPartsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUploadedPartsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.PartNumberMarker) {
		body["part_number_marker"] = request.PartNumberMarker
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.UploadId) {
		body["upload_id"] = request.UploadId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUploadedParts"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/list_uploaded_parts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUploadedPartsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries users.
//
// @param request - ListUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserResponse
func (client *Client) ListUserWithContext(ctx context.Context, request *ListUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUser"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/user/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves files or folders.
//
// @param request - MoveFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveFileResponse
func (client *Client) MoveFileWithContext(ctx context.Context, request *MoveFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MoveFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckNameMode) {
		body["check_name_mode"] = request.CheckNameMode
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.ToParentFileId) {
		body["to_parent_file_id"] = request.ToParentFileId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/move"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Penalizes files.
//
// @param request - PunishFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PunishFileResponse
func (client *Client) PunishFileWithContext(ctx context.Context, request *PunishFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PunishFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ActionCode) {
		body["action_code"] = request.ActionCode
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.PunishReason) {
		body["punish_reason"] = request.PunishReason
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PunishFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/csi/business/punish_file"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PunishFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询凌霄订单价格
//
// @param request - QueryOrderPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrderPriceResponse
func (client *Client) QueryOrderPriceWithContext(ctx context.Context, request *QueryOrderPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryOrderPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.InstanceId) {
		body["instance_id"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderType) {
		body["order_type"] = request.OrderType
	}

	if !dara.IsNil(request.Package) {
		body["package"] = request.Package
	}

	if !dara.IsNil(request.Period) {
		body["period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["period_unit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.TotalSize) {
		body["total_size"] = request.TotalSize
	}

	if !dara.IsNil(request.UserCount) {
		body["user_count"] = request.UserCount
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOrderPrice"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/query_order_price"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOrderPriceResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从人脸分组中的移除指定的文件
//
// @param request - RemoveFaceGroupFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFaceGroupFileResponse
func (client *Client) RemoveFaceGroupFileWithContext(ctx context.Context, request *RemoveFaceGroupFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveFaceGroupFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FaceGroupId) {
		body["face_group_id"] = request.FaceGroupId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveFaceGroupFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/albums/unassign_facegroup_item"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveFaceGroupFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a member from a group.
//
// @param request - RemoveGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveGroupMemberResponse
func (client *Client) RemoveGroupMemberWithContext(ctx context.Context, request *RemoveGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	if !dara.IsNil(request.MemberId) {
		body["member_id"] = request.MemberId
	}

	if !dara.IsNil(request.MemberType) {
		body["member_type"] = request.MemberType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveGroupMember"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/remove_member"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveGroupMemberResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 故事移除文件
//
// @param request - RemoveStoryFilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveStoryFilesResponse
func (client *Client) RemoveStoryFilesWithContext(ctx context.Context, request *RemoveStoryFilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveStoryFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	if !dara.IsNil(request.StoryId) {
		body["story_id"] = request.StoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveStoryFiles"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/remove_story_files"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveStoryFilesResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a file or folder from the recycle bin.
//
// @param request - RestoreFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreFileResponse
func (client *Client) RestoreFileWithContext(ctx context.Context, request *RestoreFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestoreFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/recyclebin/restore"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a historical version of a file. You cannot restore the latest version of a file.
//
// @param request - RestoreRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreRevisionResponse
func (client *Client) RestoreRevisionWithContext(ctx context.Context, request *RestoreRevisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestoreRevisionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.RevisionId) {
		body["revision_id"] = request.RevisionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreRevision"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/revision/restore"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreRevisionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scans files.
//
// @param request - ScanFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanFileResponse
func (client *Client) ScanFileWithContext(ctx context.Context, request *ScanFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScanFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScanFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/scan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScanFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries location-based groups based on specific locations.
//
// @param request - SearchAddressGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAddressGroupsResponse
func (client *Client) SearchAddressGroupsWithContext(ctx context.Context, request *SearchAddressGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchAddressGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddressLevel) {
		body["address_level"] = request.AddressLevel
	}

	if !dara.IsNil(request.AddressNames) {
		body["address_names"] = request.AddressNames
	}

	if !dara.IsNil(request.BrGeoPoint) {
		body["br_geo_point"] = request.BrGeoPoint
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.ImageThumbnailProcess) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !dara.IsNil(request.TlGeoPoint) {
		body["tl_geo_point"] = request.TlGeoPoint
	}

	if !dara.IsNil(request.VideoThumbnailProcess) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchAddressGroups"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/search_address_groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchAddressGroupsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Searches for domains
//
// @param request - SearchDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDomainsResponse
func (client *Client) SearchDomainsWithContext(ctx context.Context, request *SearchDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		body["order_by"] = request.OrderBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchDomains"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchDomainsResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries drives.
//
// @param request - SearchDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDriveResponse
func (client *Client) SearchDriveWithContext(ctx context.Context, request *SearchDriveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchDriveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveName) {
		body["drive_name"] = request.DriveName
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.Owner) {
		body["owner"] = request.Owner
	}

	if !dara.IsNil(request.OwnerType) {
		body["owner_type"] = request.OwnerType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchDrive"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchDriveResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for files.
//
// @param request - SearchFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFileResponse
func (client *Client) SearchFileWithContext(ctx context.Context, request *SearchFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.OrderBy) {
		body["order_by"] = request.OrderBy
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Recursive) {
		body["recursive"] = request.Recursive
	}

	if !dara.IsNil(request.ReturnTotalCount) {
		body["return_total_count"] = request.ReturnTotalCount
	}

	if !dara.IsNil(request.ThumbnailProcesses) {
		body["thumbnail_processes"] = request.ThumbnailProcesses
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries share URLs.
//
// @param request - SearchShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchShareLinkResponse
func (client *Client) SearchShareLinkWithContext(ctx context.Context, request *SearchShareLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchShareLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Creators) {
		body["creators"] = request.Creators
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.OrderBy) {
		body["order_by"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		body["order_direction"] = request.OrderDirection
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.ReturnTotalCount) {
		body["return_total_count"] = request.ReturnTotalCount
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchShareLink"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/share_link/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchShareLinkResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取相似图片聚类结果
//
// @param request - SearchSimilarImageClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchSimilarImageClustersResponse
func (client *Client) SearchSimilarImageClustersWithContext(ctx context.Context, request *SearchSimilarImageClustersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchSimilarImageClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.ImageThumbnailProcess) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchSimilarImageClusters"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/query_similar_image_clusters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchSimilarImageClustersResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询故事列表
//
// @param request - SearchStoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchStoriesResponse
func (client *Client) SearchStoriesWithContext(ctx context.Context, request *SearchStoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchStoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CoverImageThumbnailProcess) {
		body["cover_image_thumbnail_process"] = request.CoverImageThumbnailProcess
	}

	if !dara.IsNil(request.CoverVideoThumbnailProcess) {
		body["cover_video_thumbnail_process"] = request.CoverVideoThumbnailProcess
	}

	if !dara.IsNil(request.CreateTimeRange) {
		body["create_time_range"] = request.CreateTimeRange
	}

	if !dara.IsNil(request.CustomLabels) {
		body["custom_labels"] = request.CustomLabels
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FaceGroupIds) {
		body["face_group_ids"] = request.FaceGroupIds
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	if !dara.IsNil(request.Sort) {
		body["sort"] = request.Sort
	}

	if !dara.IsNil(request.StoryEndTimeRange) {
		body["story_end_time_range"] = request.StoryEndTimeRange
	}

	if !dara.IsNil(request.StoryId) {
		body["story_id"] = request.StoryId
	}

	if !dara.IsNil(request.StoryName) {
		body["story_name"] = request.StoryName
	}

	if !dara.IsNil(request.StoryStartTimeRange) {
		body["story_start_time_range"] = request.StoryStartTimeRange
	}

	if !dara.IsNil(request.StoryType) {
		body["story_type"] = request.StoryType
	}

	if !dara.IsNil(request.UrlExpireSec) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	if !dara.IsNil(request.WithEmptyStories) {
		body["with_empty_stories"] = request.WithEmptyStories
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchStories"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/find_stories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchStoriesResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for users.
//
// @param request - SearchUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchUserResponse
func (client *Client) SearchUserWithContext(ctx context.Context, request *SearchUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchUserResponse, _err error) {
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

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Marker) {
		body["marker"] = request.Marker
	}

	if !dara.IsNil(request.NickName) {
		body["nick_name"] = request.NickName
	}

	if !dara.IsNil(request.NickNameForFuzzy) {
		body["nick_name_for_fuzzy"] = request.NickNameForFuzzy
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.Role) {
		body["role"] = request.Role
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.UserName) {
		body["user_name"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchUser"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/user/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchUserResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an access token based on Open Authorization (OAuth) 2.0.
//
// Description:
//
// For more information about how to access Drive and Photo Service from a web server application by using OAuth 2.0, visit [OAuth 2.0 For Web Server Applications](https://www.alibabacloud.com/help/zh/pds/drive-and-photo-service-dev/user-guide/oauth-2-0-access-process-for-web-server-applications).
//
// For more information about how to access Drive and Photo Service by using a JSON Web Token (JWT) application, visit [Access process for JWT applications](https://www.alibabacloud.com/help/zh/pds/drive-and-photo-service-dev/user-guide/access-process-for-jwt-applications).
//
// @param request - TokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TokenResponse
func (client *Client) TokenWithContext(ctx context.Context, request *TokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Assertion) {
		body["assertion"] = request.Assertion
	}

	if !dara.IsNil(request.ClientId) {
		body["client_id"] = request.ClientId
	}

	if !dara.IsNil(request.ClientSecret) {
		body["client_secret"] = request.ClientSecret
	}

	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.GrantType) {
		body["grant_type"] = request.GrantType
	}

	if !dara.IsNil(request.RedirectUri) {
		body["redirect_uri"] = request.RedirectUri
	}

	if !dara.IsNil(request.RefreshToken) {
		body["refresh_token"] = request.RefreshToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Token"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/oauth/token"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TokenResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a file or folder to the recycle bin.
//
// @param request - TrashFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrashFileResponse
func (client *Client) TrashFileWithContext(ctx context.Context, request *TrashFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TrashFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrashFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/recyclebin/trash"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TrashFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Unlink Account Binding
//
// @param request - UnLinkAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnLinkAccountResponse
func (client *Client) UnLinkAccountWithContext(ctx context.Context, request *UnLinkAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnLinkAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Extra) {
		body["extra"] = request.Extra
	}

	if !dara.IsNil(request.Identity) {
		body["identity"] = request.Identity
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnLinkAccount"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/account/unlink"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnLinkAccountResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update domain information.
//
// @param request - UpdateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomainWithContext(ctx context.Context, request *UpdateDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDomainResponse, _err error) {
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

	if !dara.IsNil(request.DomainId) {
		body["domain_id"] = request.DomainId
	}

	if !dara.IsNil(request.DomainName) {
		body["domain_name"] = request.DomainName
	}

	if !dara.IsNil(request.InitDriveEnable) {
		body["init_drive_enable"] = request.InitDriveEnable
	}

	if !dara.IsNil(request.InitDriveSize) {
		body["init_drive_size"] = request.InitDriveSize
	}

	if !dara.IsNil(request.PublishedAppAccessStrategy) {
		body["published_app_access_strategy"] = request.PublishedAppAccessStrategy
	}

	if !dara.IsNil(request.SizeQuota) {
		body["size_quota"] = request.SizeQuota
	}

	if !dara.IsNil(request.UserCountQuota) {
		body["user_count_quota"] = request.UserCountQuota
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomain"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/domain/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a drive.
//
// @param request - UpdateDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDriveResponse
func (client *Client) UpdateDriveWithContext(ctx context.Context, request *UpdateDriveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDriveResponse, _err error) {
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

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.DriveName) {
		body["drive_name"] = request.DriveName
	}

	if !dara.IsNil(request.Owner) {
		body["owner"] = request.Owner
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.TotalSize) {
		body["total_size"] = request.TotalSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDrive"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/drive/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDriveResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a face-based group.
//
// @param request - UpdateFacegroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFacegroupResponse
func (client *Client) UpdateFacegroupWithContext(ctx context.Context, request *UpdateFacegroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFacegroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.GroupCoverFaceId) {
		body["group_cover_face_id"] = request.GroupCoverFaceId
	}

	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		body["group_name"] = request.GroupName
	}

	if !dara.IsNil(request.Remarks) {
		body["remarks"] = request.Remarks
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFacegroup"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/update_facegroup_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFacegroupResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a file instead of the file data.
//
// @param request - UpdateFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileResponse
func (client *Client) UpdateFileWithContext(ctx context.Context, request *UpdateFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckNameMode) {
		body["check_name_mode"] = request.CheckNameMode
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.Hidden) {
		body["hidden"] = request.Hidden
	}

	if !dara.IsNil(request.Labels) {
		body["labels"] = request.Labels
	}

	if !dara.IsNil(request.LocalModifiedAt) {
		body["local_modified_at"] = request.LocalModifiedAt
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Starred) {
		body["starred"] = request.Starred
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFile"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a group.
//
// @param request - UpdateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithContext(ctx context.Context, request *UpdateGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		body["group_id"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		body["group_name"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroup"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/group/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the mapping between an entity and a benefit package. You can call this operation to associate a benefit package with a user.
//
// @param request - UpdateIdentityToBenefitPkgMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityToBenefitPkgMappingResponse
func (client *Client) UpdateIdentityToBenefitPkgMappingWithContext(ctx context.Context, request *UpdateIdentityToBenefitPkgMappingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIdentityToBenefitPkgMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		body["amount"] = request.Amount
	}

	if !dara.IsNil(request.BenefitPkgId) {
		body["benefit_pkg_id"] = request.BenefitPkgId
	}

	if !dara.IsNil(request.ExpireTime) {
		body["expire_time"] = request.ExpireTime
	}

	if !dara.IsNil(request.IdentityId) {
		body["identity_id"] = request.IdentityId
	}

	if !dara.IsNil(request.IdentityType) {
		body["identity_type"] = request.IdentityType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdentityToBenefitPkgMapping"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/benefit/identity_to_benefit_pkg_mapping/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version information. You can call this operation to permanently retain a version or modify the description of a version. You can permanently retain up to 50 versions of a file.
//
// @param request - UpdateRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRevisionResponse
func (client *Client) UpdateRevisionWithContext(ctx context.Context, request *UpdateRevisionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRevisionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.FileId) {
		body["file_id"] = request.FileId
	}

	if !dara.IsNil(request.KeepForever) {
		body["keep_forever"] = request.KeepForever
	}

	if !dara.IsNil(request.RevisionDescription) {
		body["revision_description"] = request.RevisionDescription
	}

	if !dara.IsNil(request.RevisionId) {
		body["revision_id"] = request.RevisionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRevision"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/revision/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRevisionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a share link.
//
// @param request - UpdateShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateShareLinkResponse
func (client *Client) UpdateShareLinkWithContext(ctx context.Context, request *UpdateShareLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateShareLinkResponse, _err error) {
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

	if !dara.IsNil(request.DisableDownload) {
		body["disable_download"] = request.DisableDownload
	}

	if !dara.IsNil(request.DisablePreview) {
		body["disable_preview"] = request.DisablePreview
	}

	if !dara.IsNil(request.DisableSave) {
		body["disable_save"] = request.DisableSave
	}

	if !dara.IsNil(request.DownloadCount) {
		body["download_count"] = request.DownloadCount
	}

	if !dara.IsNil(request.DownloadLimit) {
		body["download_limit"] = request.DownloadLimit
	}

	if !dara.IsNil(request.Expiration) {
		body["expiration"] = request.Expiration
	}

	if !dara.IsNil(request.OfficeEditable) {
		body["office_editable"] = request.OfficeEditable
	}

	if !dara.IsNil(request.PreviewCount) {
		body["preview_count"] = request.PreviewCount
	}

	if !dara.IsNil(request.PreviewLimit) {
		body["preview_limit"] = request.PreviewLimit
	}

	if !dara.IsNil(request.ReportCount) {
		body["report_count"] = request.ReportCount
	}

	if !dara.IsNil(request.SaveCount) {
		body["save_count"] = request.SaveCount
	}

	if !dara.IsNil(request.SaveLimit) {
		body["save_limit"] = request.SaveLimit
	}

	if !dara.IsNil(request.ShareId) {
		body["share_id"] = request.ShareId
	}

	if !dara.IsNil(request.ShareName) {
		body["share_name"] = request.ShareName
	}

	if !dara.IsNil(request.SharePwd) {
		body["share_pwd"] = request.SharePwd
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.VideoPreviewCount) {
		body["video_preview_count"] = request.VideoPreviewCount
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateShareLink"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/share_link/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateShareLinkResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新故事
//
// @param request - UpdateStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStoryResponse
func (client *Client) UpdateStoryWithContext(ctx context.Context, request *UpdateStoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateStoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cover) {
		body["cover"] = request.Cover
	}

	if !dara.IsNil(request.CustomLabels) {
		body["custom_labels"] = request.CustomLabels
	}

	if !dara.IsNil(request.DriveId) {
		body["drive_id"] = request.DriveId
	}

	if !dara.IsNil(request.StoryId) {
		body["story_id"] = request.StoryId
	}

	if !dara.IsNil(request.StoryName) {
		body["story_name"] = request.StoryName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStory"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/image/update_story"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStoryResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a user.
//
// @param request - UpdateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithContext(ctx context.Context, request *UpdateUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Avatar) {
		body["avatar"] = request.Avatar
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		body["email"] = request.Email
	}

	if !dara.IsNil(request.GroupInfoList) {
		body["group_info_list"] = request.GroupInfoList
	}

	if !dara.IsNil(request.NickName) {
		body["nick_name"] = request.NickName
	}

	if !dara.IsNil(request.Phone) {
		body["phone"] = request.Phone
	}

	if !dara.IsNil(request.Role) {
		body["role"] = request.Role
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.UserData) {
		body["user_data"] = request.UserData
	}

	if !dara.IsNil(request.UserId) {
		body["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/user/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the digital rights management (DRM) license of a video.
//
// @param request - VideoDRMLicenseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoDRMLicenseResponse
func (client *Client) VideoDRMLicenseWithContext(ctx context.Context, request *VideoDRMLicenseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VideoDRMLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DrmType) {
		body["drmType"] = request.DrmType
	}

	if !dara.IsNil(request.LicenseRequest) {
		body["licenseRequest"] = request.LicenseRequest
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VideoDRMLicense"),
		Version:     dara.String("2022-03-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/file/video_drm_license"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VideoDRMLicenseResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
