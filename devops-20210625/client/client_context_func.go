// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 添加组成员
//
// @param request - AddGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupMemberResponse
func (client *Client) AddGroupMemberWithContext(ctx context.Context, groupId *string, request *AddGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessLevel) {
		body["accessLevel"] = request.AccessLevel
	}

	if !dara.IsNil(request.AliyunPks) {
		body["aliyunPks"] = request.AliyunPks
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGroupMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/members/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加流水线关联
//
// @param request - AddPipelineRelationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPipelineRelationsResponse
func (client *Client) AddPipelineRelationsWithContext(ctx context.Context, organizationId *string, pipelineId *string, request *AddPipelineRelationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddPipelineRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RelObjectIds) {
		query["relObjectIds"] = request.RelObjectIds
	}

	if !dara.IsNil(request.RelObjectType) {
		query["relObjectType"] = request.RelObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPipelineRelations"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRelations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPipelineRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加代码库成员
//
// @param request - AddRepositoryMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRepositoryMemberResponse
func (client *Client) AddRepositoryMemberWithContext(ctx context.Context, repositoryId *string, request *AddRepositoryMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddRepositoryMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessLevel) {
		body["accessLevel"] = request.AccessLevel
	}

	if !dara.IsNil(request.AliyunPks) {
		body["aliyunPks"] = request.AliyunPks
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRepositoryMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/members"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRepositoryMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加代码库Webhook
//
// @param request - AddWebhookRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWebhookResponse
func (client *Client) AddWebhookWithContext(ctx context.Context, repositoryId *string, request *AddWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnableSslVerification) {
		body["enableSslVerification"] = request.EnableSslVerification
	}

	if !dara.IsNil(request.MergeRequestsEvents) {
		body["mergeRequestsEvents"] = request.MergeRequestsEvents
	}

	if !dara.IsNil(request.NoteEvents) {
		body["noteEvents"] = request.NoteEvents
	}

	if !dara.IsNil(request.PushEvents) {
		body["pushEvents"] = request.PushEvents
	}

	if !dara.IsNil(request.SecretToken) {
		body["secretToken"] = request.SecretToken
	}

	if !dara.IsNil(request.TagPushEvents) {
		body["tagPushEvents"] = request.TagPushEvents
	}

	if !dara.IsNil(request.Url) {
		body["url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWebhook"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/webhooks/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消执行研发阶段流水线
//
// @param request - CancelExecutionReleaseStageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelExecutionReleaseStageResponse
func (client *Client) CancelExecutionReleaseStageWithContext(ctx context.Context, appName *string, releaseWorkflowSn *string, releaseStageSn *string, executionNumber *string, request *CancelExecutionReleaseStageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelExecutionReleaseStageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelExecutionReleaseStage"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/releaseWorkflows/" + dara.PercentEncode(dara.StringValue(releaseWorkflowSn)) + "/releaseStages/" + dara.PercentEncode(dara.StringValue(releaseStageSn)) + "/executions/" + dara.PercentEncode(dara.StringValue(executionNumber)) + "%3Acancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelExecutionReleaseStageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭代码评审
//
// @param request - CloseMergeRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseMergeRequestResponse
func (client *Client) CloseMergeRequestWithContext(ctx context.Context, repositoryId *string, localId *string, request *CloseMergeRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseMergeRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseMergeRequest"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/merge_requests/" + dara.PercentEncode(dara.StringValue(localId)) + "/close"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseMergeRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加应用成员
//
// @param request - CreateAppMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppMembersResponse
func (client *Client) CreateAppMembersWithContext(ctx context.Context, appName *string, request *CreateAppMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAppMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PlayerList) {
		body["playerList"] = request.PlayerList
	}

	if !dara.IsNil(request.RoleNames) {
		body["roleNames"] = request.RoleNames
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppMembers"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/members"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &CreateAppMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建分支
//
// @param request - CreateBranchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBranchResponse
func (client *Client) CreateBranchWithContext(ctx context.Context, repositoryId *string, request *CreateBranchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBranchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BranchName) {
		body["branchName"] = request.BranchName
	}

	if !dara.IsNil(request.Ref) {
		body["ref"] = request.Ref
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBranch"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/branches"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBranchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建变更
//
// @param request - CreateChangeRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChangeRequestResponse
func (client *Client) CreateChangeRequestWithContext(ctx context.Context, appName *string, request *CreateChangeRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateChangeRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppCodeRepoSn) {
		body["appCodeRepoSn"] = request.AppCodeRepoSn
	}

	if !dara.IsNil(request.AutoDeleteBranchWhenEnd) {
		body["autoDeleteBranchWhenEnd"] = request.AutoDeleteBranchWhenEnd
	}

	if !dara.IsNil(request.BranchName) {
		body["branchName"] = request.BranchName
	}

	if !dara.IsNil(request.CreateBranch) {
		body["createBranch"] = request.CreateBranch
	}

	if !dara.IsNil(request.OwnerAccountId) {
		body["ownerAccountId"] = request.OwnerAccountId
	}

	if !dara.IsNil(request.OwnerId) {
		body["ownerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChangeRequest"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/changeRequests"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChangeRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加检查运行记录
//
// @param request - CreateCheckRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCheckRunResponse
func (client *Client) CreateCheckRunWithContext(ctx context.Context, request *CreateCheckRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCheckRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Annotations) {
		body["annotations"] = request.Annotations
	}

	if !dara.IsNil(request.CompletedAt) {
		body["completedAt"] = request.CompletedAt
	}

	if !dara.IsNil(request.Conclusion) {
		body["conclusion"] = request.Conclusion
	}

	if !dara.IsNil(request.DetailsUrl) {
		body["detailsUrl"] = request.DetailsUrl
	}

	if !dara.IsNil(request.ExternalId) {
		body["externalId"] = request.ExternalId
	}

	if !dara.IsNil(request.HeadSha) {
		body["headSha"] = request.HeadSha
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.StartedAt) {
		body["startedAt"] = request.StartedAt
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCheckRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/check_runs/create_check_run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCheckRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建评论
//
// @param request - CreateCommentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCommentResponse
func (client *Client) CreateCommentWithContext(ctx context.Context, request *CreateCommentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCommentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.LocalId) {
		query["localId"] = request.LocalId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommentType) {
		body["commentType"] = request.CommentType
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Draft) {
		body["draft"] = request.Draft
	}

	if !dara.IsNil(request.FilePath) {
		body["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.FromPachSetBizId) {
		body["fromPachSetBizId"] = request.FromPachSetBizId
	}

	if !dara.IsNil(request.LineNumber) {
		body["lineNumber"] = request.LineNumber
	}

	if !dara.IsNil(request.ParentCommentBizId) {
		body["parentCommentBizId"] = request.ParentCommentBizId
	}

	if !dara.IsNil(request.PatchSetBizId) {
		body["patchSetBizId"] = request.PatchSetBizId
	}

	if !dara.IsNil(request.Resolved) {
		body["resolved"] = request.Resolved
	}

	if !dara.IsNil(request.ToPatchSetBizId) {
		body["toPatchSetBizId"] = request.ToPatchSetBizId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComment"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/code_reviews/comments/create_comment"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCommentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建提交状态记录
//
// @param request - CreateCommitStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCommitStatusResponse
func (client *Client) CreateCommitStatusWithContext(ctx context.Context, request *CreateCommitStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCommitStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	if !dara.IsNil(request.Sha) {
		query["sha"] = request.Sha
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Context) {
		body["context"] = request.Context
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.State) {
		body["state"] = request.State
	}

	if !dara.IsNil(request.TargetUrl) {
		body["targetUrl"] = request.TargetUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCommitStatus"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/repository/commit_statuses/create_commit_status"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCommitStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 单提交变更多个文件
//
// @param request - CreateCommitWithMultipleFilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCommitWithMultipleFilesResponse
func (client *Client) CreateCommitWithMultipleFilesWithContext(ctx context.Context, request *CreateCommitWithMultipleFilesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCommitWithMultipleFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Actions) {
		body["actions"] = request.Actions
	}

	if !dara.IsNil(request.Branch) {
		body["branch"] = request.Branch
	}

	if !dara.IsNil(request.CommitMessage) {
		body["commitMessage"] = request.CommitMessage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCommitWithMultipleFiles"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/repository/commits/files"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCommitWithMultipleFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建部署密钥
//
// @param request - CreateDeployKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeployKeyResponse
func (client *Client) CreateDeployKeyWithContext(ctx context.Context, repositoryId *string, request *CreateDeployKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDeployKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		body["key"] = request.Key
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeployKey"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/keys/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeployKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文件
//
// @param request - CreateFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithContext(ctx context.Context, repositoryId *string, request *CreateFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BranchName) {
		body["branchName"] = request.BranchName
	}

	if !dara.IsNil(request.CommitMessage) {
		body["commitMessage"] = request.CommitMessage
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Encoding) {
		body["encoding"] = request.Encoding
	}

	if !dara.IsNil(request.FilePath) {
		body["filePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFile"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/files"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建标签
//
// @param request - CreateFlowTagRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowTagResponse
func (client *Client) CreateFlowTagWithContext(ctx context.Context, organizationId *string, request *CreateFlowTagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFlowTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Color) {
		query["color"] = request.Color
	}

	if !dara.IsNil(request.FlowTagGroupId) {
		query["flowTagGroupId"] = request.FlowTagGroupId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowTag"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/flow/tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建标签分类
//
// @param request - CreateFlowTagGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowTagGroupResponse
func (client *Client) CreateFlowTagGroupWithContext(ctx context.Context, organizationId *string, request *CreateFlowTagGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFlowTagGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowTagGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/flow/tagGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowTagGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建主机组
//
// @param request - CreateHostGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostGroupResponse
func (client *Client) CreateHostGroupWithContext(ctx context.Context, organizationId *string, request *CreateHostGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHostGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunRegion) {
		body["aliyunRegion"] = request.AliyunRegion
	}

	if !dara.IsNil(request.EcsLabelKey) {
		body["ecsLabelKey"] = request.EcsLabelKey
	}

	if !dara.IsNil(request.EcsLabelValue) {
		body["ecsLabelValue"] = request.EcsLabelValue
	}

	if !dara.IsNil(request.EcsType) {
		body["ecsType"] = request.EcsType
	}

	if !dara.IsNil(request.EnvId) {
		body["envId"] = request.EnvId
	}

	if !dara.IsNil(request.MachineInfos) {
		body["machineInfos"] = request.MachineInfos
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ServiceConnectionId) {
		body["serviceConnectionId"] = request.ServiceConnectionId
	}

	if !dara.IsNil(request.TagIds) {
		body["tagIds"] = request.TagIds
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/hostGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建代码评审
//
// @param request - CreateMergeRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMergeRequestResponse
func (client *Client) CreateMergeRequestWithContext(ctx context.Context, repositoryId *string, request *CreateMergeRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMergeRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateFrom) {
		body["createFrom"] = request.CreateFrom
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ReviewerIds) {
		body["reviewerIds"] = request.ReviewerIds
	}

	if !dara.IsNil(request.SourceBranch) {
		body["sourceBranch"] = request.SourceBranch
	}

	if !dara.IsNil(request.SourceProjectId) {
		body["sourceProjectId"] = request.SourceProjectId
	}

	if !dara.IsNil(request.TargetBranch) {
		body["targetBranch"] = request.TargetBranch
	}

	if !dara.IsNil(request.TargetProjectId) {
		body["targetProjectId"] = request.TargetProjectId
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	if !dara.IsNil(request.WorkItemIds) {
		body["workItemIds"] = request.WorkItemIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMergeRequest"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/merge_requests"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMergeRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建OAuth令牌
//
// @param request - CreateOAuthTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOAuthTokenResponse
func (client *Client) CreateOAuthTokenWithContext(ctx context.Context, request *CreateOAuthTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateOAuthTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientSecret) {
		body["clientSecret"] = request.ClientSecret
	}

	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.GrantType) {
		body["grantType"] = request.GrantType
	}

	if !dara.IsNil(request.Login) {
		body["login"] = request.Login
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOAuthToken"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/login/oauth/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOAuthTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流水线。
//
// @param request - CreatePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineResponse
func (client *Client) CreatePipelineWithContext(ctx context.Context, organizationId *string, request *CreatePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipeline"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流水线分组
//
// @param request - CreatePipelineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineGroupResponse
func (client *Client) CreatePipelineGroupWithContext(ctx context.Context, organizationId *string, request *CreatePipelineGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePipelineGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipelineGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelineGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建项目
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithContext(ctx context.Context, organizationId *string, request *CreateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomCode) {
		body["customCode"] = request.CustomCode
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.TemplateIdentifier) {
		body["templateIdentifier"] = request.TemplateIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/projects/createProject"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建代码库Label
//
// @param request - CreateProjectLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectLabelResponse
func (client *Client) CreateProjectLabelWithContext(ctx context.Context, request *CreateProjectLabelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Color) {
		body["color"] = request.Color
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProjectLabel"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/labels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建保护分支
//
// @param request - CreateProtectdBranchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtectdBranchResponse
func (client *Client) CreateProtectdBranchWithContext(ctx context.Context, repositoryId *string, request *CreateProtectdBranchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProtectdBranchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowMergeRoles) {
		body["allowMergeRoles"] = request.AllowMergeRoles
	}

	if !dara.IsNil(request.AllowMergeUserIds) {
		body["allowMergeUserIds"] = request.AllowMergeUserIds
	}

	if !dara.IsNil(request.AllowPushRoles) {
		body["allowPushRoles"] = request.AllowPushRoles
	}

	if !dara.IsNil(request.AllowPushUserIds) {
		body["allowPushUserIds"] = request.AllowPushUserIds
	}

	if !dara.IsNil(request.Branch) {
		body["branch"] = request.Branch
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.MergeRequestSetting) {
		body["mergeRequestSetting"] = request.MergeRequestSetting
	}

	if !dara.IsNil(request.TestSettingDTO) {
		body["testSettingDTO"] = request.TestSettingDTO
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProtectdBranch"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/protect_branches"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProtectdBranchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建推送规则
//
// @param request - CreatePushRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePushRuleResponse
func (client *Client) CreatePushRuleWithContext(ctx context.Context, repositoryId *string, request *CreatePushRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePushRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleInfos) {
		body["ruleInfos"] = request.RuleInfos
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePushRule"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/push_rule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePushRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建(导入)代码库
//
// @param request - CreateRepositoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepositoryResponse
func (client *Client) CreateRepositoryWithContext(ctx context.Context, request *CreateRepositoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.CreateParentPath) {
		query["createParentPath"] = request.CreateParentPath
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Sync) {
		query["sync"] = request.Sync
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AvatarUrl) {
		body["avatarUrl"] = request.AvatarUrl
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GitignoreType) {
		body["gitignoreType"] = request.GitignoreType
	}

	if !dara.IsNil(request.ImportAccount) {
		body["importAccount"] = request.ImportAccount
	}

	if !dara.IsNil(request.ImportDemoProject) {
		body["importDemoProject"] = request.ImportDemoProject
	}

	if !dara.IsNil(request.ImportRepoType) {
		body["importRepoType"] = request.ImportRepoType
	}

	if !dara.IsNil(request.ImportToken) {
		body["importToken"] = request.ImportToken
	}

	if !dara.IsNil(request.ImportTokenEncrypted) {
		body["importTokenEncrypted"] = request.ImportTokenEncrypted
	}

	if !dara.IsNil(request.ImportUrl) {
		body["importUrl"] = request.ImportUrl
	}

	if !dara.IsNil(request.InitStandardService) {
		body["initStandardService"] = request.InitStandardService
	}

	if !dara.IsNil(request.IsCryptoEnabled) {
		body["isCryptoEnabled"] = request.IsCryptoEnabled
	}

	if !dara.IsNil(request.LocalImportUrl) {
		body["localImportUrl"] = request.LocalImportUrl
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NamespaceId) {
		body["namespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Path) {
		body["path"] = request.Path
	}

	if !dara.IsNil(request.ReadmeType) {
		body["readmeType"] = request.ReadmeType
	}

	if !dara.IsNil(request.VisibilityLevel) {
		body["visibilityLevel"] = request.VisibilityLevel
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepository"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建代码组
//
// @param request - CreateRepositoryGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepositoryGroupResponse
func (client *Client) CreateRepositoryGroupWithContext(ctx context.Context, request *CreateRepositoryGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRepositoryGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AvatarUrl) {
		body["avatarUrl"] = request.AvatarUrl
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.Path) {
		body["path"] = request.Path
	}

	if !dara.IsNil(request.VisibilityLevel) {
		body["visibilityLevel"] = request.VisibilityLevel
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepositoryGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepositoryGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插入资源成员
//
// @param request - CreateResourceMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceMemberResponse
func (client *Client) CreateResourceMemberWithContext(ctx context.Context, organizationId *string, resourceType *string, resourceId *string, request *CreateResourceMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		body["accountId"] = request.AccountId
	}

	if !dara.IsNil(request.RoleName) {
		body["roleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/" + dara.PercentEncode(dara.StringValue(resourceType)) + "/" + dara.PercentEncode(dara.StringValue(resourceId)) + "/members"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务授权
//
// @param request - CreateServiceAuthRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceAuthResponse
func (client *Client) CreateServiceAuthWithContext(ctx context.Context, organizationId *string, request *CreateServiceAuthRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceAuthType) {
		query["serviceAuthType"] = request.ServiceAuthType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceAuth"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/serviceAuths"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务连接
//
// @param request - CreateServiceConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceConnectionResponse
func (client *Client) CreateServiceConnectionWithContext(ctx context.Context, organizationId *string, request *CreateServiceConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		body["authType"] = request.AuthType
	}

	if !dara.IsNil(request.ConnectionName) {
		body["connectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.ConnectionType) {
		body["connectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.ServiceAuthId) {
		body["serviceAuthId"] = request.ServiceAuthId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceConnection"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/createServiceConnection"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户名密码类型的证书
//
// @param request - CreateServiceCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceCredentialResponse
func (client *Client) CreateServiceCredentialWithContext(ctx context.Context, organizationId *string, request *CreateServiceCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.Username) {
		body["username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceCredential"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/serviceCredentials"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建迭代
//
// @param request - CreateSprintRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSprintResponse
func (client *Client) CreateSprintWithContext(ctx context.Context, organizationId *string, request *CreateSprintRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSprintResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		body["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SpaceIdentifier) {
		body["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.StaffIds) {
		body["staffIds"] = request.StaffIds
	}

	if !dara.IsNil(request.StartDate) {
		body["startDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSprint"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/sprints/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSprintResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建企业公钥
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSshKeyResponse
func (client *Client) CreateSshKeyWithContext(ctx context.Context, organizationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSshKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSshKey"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/sshKey"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSshKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建标签Tag
//
// @param request - CreateTagRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTagResponse
func (client *Client) CreateTagWithContext(ctx context.Context, repositoryId *string, request *CreateTagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Message) {
		body["message"] = request.Message
	}

	if !dara.IsNil(request.Ref) {
		body["ref"] = request.Ref
	}

	if !dara.IsNil(request.TagName) {
		body["tagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTag"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/tags/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建测试用例
//
// @param request - CreateTestCaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTestCaseResponse
func (client *Client) CreateTestCaseWithContext(ctx context.Context, organizationId *string, request *CreateTestCaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTestCaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssignedTo) {
		body["assignedTo"] = request.AssignedTo
	}

	if !dara.IsNil(request.DirectoryIdentifier) {
		body["directoryIdentifier"] = request.DirectoryIdentifier
	}

	if !dara.IsNil(request.FieldValueList) {
		body["fieldValueList"] = request.FieldValueList
	}

	if !dara.IsNil(request.Priority) {
		body["priority"] = request.Priority
	}

	if !dara.IsNil(request.SpaceIdentifier) {
		body["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.Subject) {
		body["subject"] = request.Subject
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.TestcaseStepContentInfo) {
		body["testcaseStepContentInfo"] = request.TestcaseStepContentInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTestCase"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/testhub/testcase"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTestCaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建SSH Key密钥
//
// @param request - CreateUserKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserKeyResponse
func (client *Client) CreateUserKeyWithContext(ctx context.Context, request *CreateUserKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateUserKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireTime) {
		body["expireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.KeyScope) {
		body["keyScope"] = request.KeyScope
	}

	if !dara.IsNil(request.PublicKey) {
		body["publicKey"] = request.PublicKey
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserKey"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v3/user/keys/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建变量组
//
// @param request - CreateVariableGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVariableGroupResponse
func (client *Client) CreateVariableGroupWithContext(ctx context.Context, organizationId *string, request *CreateVariableGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVariableGroupResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVariableGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/variableGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVariableGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建工作项
//
// @param request - CreateWorkitemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkitemResponse
func (client *Client) CreateWorkitemWithContext(ctx context.Context, organizationId *string, request *CreateWorkitemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkitemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssignedTo) {
		body["assignedTo"] = request.AssignedTo
	}

	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DescriptionFormat) {
		body["descriptionFormat"] = request.DescriptionFormat
	}

	if !dara.IsNil(request.FieldValueList) {
		body["fieldValueList"] = request.FieldValueList
	}

	if !dara.IsNil(request.Parent) {
		body["parent"] = request.Parent
	}

	if !dara.IsNil(request.Participant) {
		body["participant"] = request.Participant
	}

	if !dara.IsNil(request.Space) {
		body["space"] = request.Space
	}

	if !dara.IsNil(request.SpaceIdentifier) {
		body["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.SpaceType) {
		body["spaceType"] = request.SpaceType
	}

	if !dara.IsNil(request.Sprint) {
		body["sprint"] = request.Sprint
	}

	if !dara.IsNil(request.Subject) {
		body["subject"] = request.Subject
	}

	if !dara.IsNil(request.Tracker) {
		body["tracker"] = request.Tracker
	}

	if !dara.IsNil(request.Verifier) {
		body["verifier"] = request.Verifier
	}

	if !dara.IsNil(request.WorkitemType) {
		body["workitemType"] = request.WorkitemType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkitem"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkitemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个评论
//
// @param request - CreateWorkitemCommentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkitemCommentResponse
func (client *Client) CreateWorkitemCommentWithContext(ctx context.Context, organizationId *string, request *CreateWorkitemCommentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkitemCommentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.FormatType) {
		body["formatType"] = request.FormatType
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.WorkitemIdentifier) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkitemComment"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/comment"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkitemCommentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 登记预计工时
//
// @param request - CreateWorkitemEstimateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkitemEstimateResponse
func (client *Client) CreateWorkitemEstimateWithContext(ctx context.Context, organizationId *string, request *CreateWorkitemEstimateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkitemEstimateResponse, _err error) {
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

	if !dara.IsNil(request.RecordUserIdentifier) {
		body["recordUserIdentifier"] = request.RecordUserIdentifier
	}

	if !dara.IsNil(request.SpentTime) {
		body["spentTime"] = request.SpentTime
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.WorkitemIdentifier) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkitemEstimate"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/estimate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkitemEstimateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 登记实际工时
//
// @param request - CreateWorkitemRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkitemRecordResponse
func (client *Client) CreateWorkitemRecordWithContext(ctx context.Context, organizationId *string, request *CreateWorkitemRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkitemRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ActualTime) {
		body["actualTime"] = request.ActualTime
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GmtEnd) {
		body["gmtEnd"] = request.GmtEnd
	}

	if !dara.IsNil(request.GmtStart) {
		body["gmtStart"] = request.GmtStart
	}

	if !dara.IsNil(request.RecordUserIdentifier) {
		body["recordUserIdentifier"] = request.RecordUserIdentifier
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.WorkitemIdentifier) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkitemRecord"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/record"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkitemRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作项
//
// @param request - CreateWorkitemV2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkitemV2Response
func (client *Client) CreateWorkitemV2WithContext(ctx context.Context, organizationId *string, request *CreateWorkitemV2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkitemV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssignedTo) {
		body["assignedTo"] = request.AssignedTo
	}

	if !dara.IsNil(request.Category) {
		body["category"] = request.Category
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.FieldValueList) {
		body["fieldValueList"] = request.FieldValueList
	}

	if !dara.IsNil(request.ParentIdentifier) {
		body["parentIdentifier"] = request.ParentIdentifier
	}

	if !dara.IsNil(request.Participants) {
		body["participants"] = request.Participants
	}

	if !dara.IsNil(request.SpaceIdentifier) {
		body["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.SprintIdentifier) {
		body["sprintIdentifier"] = request.SprintIdentifier
	}

	if !dara.IsNil(request.Subject) {
		body["subject"] = request.Subject
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.Trackers) {
		body["trackers"] = request.Trackers
	}

	if !dara.IsNil(request.Verifier) {
		body["verifier"] = request.Verifier
	}

	if !dara.IsNil(request.Versions) {
		body["versions"] = request.Versions
	}

	if !dara.IsNil(request.WorkitemTypeIdentifier) {
		body["workitemTypeIdentifier"] = request.WorkitemTypeIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkitemV2"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkitemV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用成员
//
// @param request - DeleteAppMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppMemberResponse
func (client *Client) DeleteAppMemberWithContext(ctx context.Context, appName *string, request *DeleteAppMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAppMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.SubjectId) {
		query["subjectId"] = request.SubjectId
	}

	if !dara.IsNil(request.SubjectType) {
		query["subjectType"] = request.SubjectType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/members"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &DeleteAppMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除分支
//
// @param request - DeleteBranchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBranchResponse
func (client *Client) DeleteBranchWithContext(ctx context.Context, repositoryId *string, request *DeleteBranchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBranchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.BranchName) {
		query["branchName"] = request.BranchName
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBranch"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/branches/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBranchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件
//
// @param request - DeleteFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithContext(ctx context.Context, repositoryId *string, request *DeleteFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.BranchName) {
		query["branchName"] = request.BranchName
	}

	if !dara.IsNil(request.CommitMessage) {
		query["commitMessage"] = request.CommitMessage
	}

	if !dara.IsNil(request.FilePath) {
		query["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFile"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/files/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowTagResponse
func (client *Client) DeleteFlowTagWithContext(ctx context.Context, organizationId *string, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFlowTagResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlowTag"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/flow/tags/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签分类
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowTagGroupResponse
func (client *Client) DeleteFlowTagGroupWithContext(ctx context.Context, organizationId *string, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFlowTagGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlowTagGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/flow/tagGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowTagGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除组成员
//
// @param request - DeleteGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupMemberResponse
func (client *Client) DeleteGroupMemberWithContext(ctx context.Context, groupId *string, request *DeleteGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.AliyunPk) {
		query["aliyunPk"] = request.AliyunPk
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MemberType) {
		body["memberType"] = request.MemberType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroupMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/members/remove/aliyun_pk"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除主机组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostGroupResponse
func (client *Client) DeleteHostGroupWithContext(ctx context.Context, organizationId *string, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHostGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHostGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/hostGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流水线
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineResponse
func (client *Client) DeletePipelineWithContext(ctx context.Context, organizationId *string, pipelineId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePipelineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePipeline"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流水线分组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineGroupResponse
func (client *Client) DeletePipelineGroupWithContext(ctx context.Context, organizationId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePipelineGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePipelineGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelineGroups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePipelineGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流水线关联
//
// @param request - DeletePipelineRelationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineRelationsResponse
func (client *Client) DeletePipelineRelationsWithContext(ctx context.Context, organizationId *string, pipelineId *string, request *DeletePipelineRelationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePipelineRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RelObjectId) {
		query["relObjectId"] = request.RelObjectId
	}

	if !dara.IsNil(request.RelObjectType) {
		query["relObjectType"] = request.RelObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePipelineRelations"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRelations"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePipelineRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目
//
// @param request - DeleteProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithContext(ctx context.Context, organizationId *string, request *DeleteProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identifier) {
		query["identifier"] = request.Identifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/projects/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代码库Label
//
// @param request - DeleteProjectLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectLabelResponse
func (client *Client) DeleteProjectLabelWithContext(ctx context.Context, labelId *string, request *DeleteProjectLabelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProjectLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProjectLabel"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/labels/" + dara.PercentEncode(dara.StringValue(labelId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除保护分支
//
// @param request - DeleteProtectedBranchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtectedBranchResponse
func (client *Client) DeleteProtectedBranchWithContext(ctx context.Context, repositoryId *string, protectedBranchId *string, request *DeleteProtectedBranchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProtectedBranchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProtectedBranch"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/protect_branches/" + dara.PercentEncode(dara.StringValue(protectedBranchId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProtectedBranchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除推送规则
//
// @param request - DeletePushRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePushRuleResponse
func (client *Client) DeletePushRuleWithContext(ctx context.Context, repositoryId *string, pushRuleId *string, request *DeletePushRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePushRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePushRule"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/push_rule/" + dara.PercentEncode(dara.StringValue(pushRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePushRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代码库
//
// @param request - DeleteRepositoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepositoryResponse
func (client *Client) DeleteRepositoryWithContext(ctx context.Context, repositoryId *string, request *DeleteRepositoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepository"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/remove"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代码组
//
// @param request - DeleteRepositoryGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepositoryGroupResponse
func (client *Client) DeleteRepositoryGroupWithContext(ctx context.Context, groupId *string, request *DeleteRepositoryGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRepositoryGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepositoryGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/remove"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepositoryGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代码库成员
//
// @param request - DeleteRepositoryMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepositoryMemberResponse
func (client *Client) DeleteRepositoryMemberWithContext(ctx context.Context, repositoryId *string, aliyunPk *string, request *DeleteRepositoryMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRepositoryMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MemberType) {
		body["memberType"] = request.MemberType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepositoryMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/members/delete/" + dara.PercentEncode(dara.StringValue(aliyunPk))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepositoryMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代码库Webhook
//
// @param request - DeleteRepositoryWebhookRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepositoryWebhookResponse
func (client *Client) DeleteRepositoryWebhookWithContext(ctx context.Context, repositoryId *string, hookId *string, request *DeleteRepositoryWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRepositoryWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRepositoryWebhook"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/hooks/" + dara.PercentEncode(dara.StringValue(hookId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRepositoryWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源成员
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceMemberResponse
func (client *Client) DeleteResourceMemberWithContext(ctx context.Context, organizationId *string, resourceType *string, resourceId *string, accountId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceMemberResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/" + dara.PercentEncode(dara.StringValue(resourceType)) + "/" + dara.PercentEncode(dara.StringValue(resourceId)) + "/members/" + dara.PercentEncode(dara.StringValue(accountId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - DeleteTagRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTagResponse
func (client *Client) DeleteTagWithContext(ctx context.Context, repositoryId *string, request *DeleteTagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.TagName) {
		query["tagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTag"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/tags/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户的SSH Key
//
// @param request - DeleteUserKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserKeyResponse
func (client *Client) DeleteUserKeyWithContext(ctx context.Context, keyId *string, request *DeleteUserKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUserKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserKey"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v3/user/keys/" + dara.PercentEncode(dara.StringValue(keyId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除变量组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVariableGroupResponse
func (client *Client) DeleteVariableGroupWithContext(ctx context.Context, organizationId *string, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteVariableGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVariableGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/variableGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVariableGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作项
//
// @param request - DeleteWorkitemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkitemResponse
func (client *Client) DeleteWorkitemWithContext(ctx context.Context, organizationId *string, request *DeleteWorkitemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkitemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identifier) {
		query["identifier"] = request.Identifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkitem"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitem/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkitemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除所有评论
//
// @param request - DeleteWorkitemAllCommentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkitemAllCommentResponse
func (client *Client) DeleteWorkitemAllCommentWithContext(ctx context.Context, organizationId *string, request *DeleteWorkitemAllCommentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkitemAllCommentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identifier) {
		query["identifier"] = request.Identifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkitemAllComment"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/deleteAllComment"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkitemAllCommentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除单条评论
//
// @param request - DeleteWorkitemCommentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkitemCommentResponse
func (client *Client) DeleteWorkitemCommentWithContext(ctx context.Context, organizationId *string, request *DeleteWorkitemCommentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkitemCommentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CommentId) {
		body["commentId"] = request.CommentId
	}

	if !dara.IsNil(request.Identifier) {
		body["identifier"] = request.Identifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkitemComment"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/deleteComent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkitemCommentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用部署密钥
//
// @param request - EnableDeployKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDeployKeyResponse
func (client *Client) EnableDeployKeyWithContext(ctx context.Context, repositoryId *string, keyId *string, request *EnableDeployKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableDeployKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDeployKey"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/keys/" + dara.PercentEncode(dara.StringValue(keyId)) + "/enable"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDeployKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行研发阶段流水线
//
// @param request - ExecuteChangeRequestReleaseStageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteChangeRequestReleaseStageResponse
func (client *Client) ExecuteChangeRequestReleaseStageWithContext(ctx context.Context, appName *string, releaseWorkflowSn *string, releaseStageSn *string, request *ExecuteChangeRequestReleaseStageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteChangeRequestReleaseStageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteChangeRequestReleaseStage"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/releaseWorkflows/" + dara.PercentEncode(dara.StringValue(releaseWorkflowSn)) + "/releaseStages/" + dara.PercentEncode(dara.StringValue(releaseStageSn)) + "%3Aexecute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteChangeRequestReleaseStageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight custom_value表
//
// @param request - ExportInsightCustomValueRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightCustomValueResponse
func (client *Client) ExportInsightCustomValueWithContext(ctx context.Context, organizationId *string, request *ExportInsightCustomValueRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightCustomValueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightCustomValue"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/customValues"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightCustomValueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight expected_work_time表数据
//
// @param request - ExportInsightExpectedWorkTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightExpectedWorkTimeResponse
func (client *Client) ExportInsightExpectedWorkTimeWithContext(ctx context.Context, organizationId *string, request *ExportInsightExpectedWorkTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightExpectedWorkTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightExpectedWorkTime"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/expectedWorkTimes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightExpectedWorkTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight field表
//
// @param request - ExportInsightFieldRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightFieldResponse
func (client *Client) ExportInsightFieldWithContext(ctx context.Context, organizationId *string, request *ExportInsightFieldRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightField"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/fields"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight space表数据
//
// @param request - ExportInsightSpaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightSpaceResponse
func (client *Client) ExportInsightSpaceWithContext(ctx context.Context, organizationId *string, request *ExportInsightSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightSpace"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/spaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight space_ref表数据
//
// @param request - ExportInsightSpaceRefRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightSpaceRefResponse
func (client *Client) ExportInsightSpaceRefWithContext(ctx context.Context, organizationId *string, request *ExportInsightSpaceRefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightSpaceRefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightSpaceRef"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/spaceRefs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightSpaceRefResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight sprint表数据
//
// @param request - ExportInsightSprintRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightSprintResponse
func (client *Client) ExportInsightSprintWithContext(ctx context.Context, organizationId *string, request *ExportInsightSprintRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightSprintResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightSprint"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/sprints"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightSprintResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight tag_ref表数据
//
// @param request - ExportInsightTagRefRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightTagRefResponse
func (client *Client) ExportInsightTagRefWithContext(ctx context.Context, organizationId *string, request *ExportInsightTagRefRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightTagRefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightTagRef"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/tagRefs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightTagRefResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight work_time表数据
//
// @param request - ExportInsightWorkTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightWorkTimeResponse
func (client *Client) ExportInsightWorkTimeWithContext(ctx context.Context, organizationId *string, request *ExportInsightWorkTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightWorkTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightWorkTime"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/workTimes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightWorkTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight workitem_stauts表数据
//
// @param request - ExportInsightWorkitemStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightWorkitemStatusResponse
func (client *Client) ExportInsightWorkitemStatusWithContext(ctx context.Context, organizationId *string, request *ExportInsightWorkitemStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightWorkitemStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightWorkitemStatus"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/workitemStatuses"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightWorkitemStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight workitem_stauts表 join workitem_defect_extra表表数据
//
// @param request - ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse
func (client *Client) ExportInsightWorkitemStatusJoinWorkitemDefectExtraWithContext(ctx context.Context, organizationId *string, request *ExportInsightWorkitemStatusJoinWorkitemDefectExtraRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightWorkitemStatusJoinWorkitemDefectExtra"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/workitemStatusJoinWorkitemDefectExtras"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightWorkitemStatusJoinWorkitemDefectExtraResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight workitem_version表数据
//
// @param request - ExportInsightWorkitemVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportInsightWorkitemVersionResponse
func (client *Client) ExportInsightWorkitemVersionWithContext(ctx context.Context, organizationId *string, request *ExportInsightWorkitemVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportInsightWorkitemVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportInsightWorkitemVersion"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/workitemVersions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportInsightWorkitemVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出Insight workitem_activity表数据
//
// @param request - ExportWorkitemActivityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportWorkitemActivityResponse
func (client *Client) ExportWorkitemActivityWithContext(ctx context.Context, organizationId *string, request *ExportWorkitemActivityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportWorkitemActivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		body["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportWorkitemActivity"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/data/workitemActivities"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportWorkitemActivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查找应用详情
//
// @param request - GetApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithContext(ctx context.Context, appName *string, request *GetApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplication"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 查询单个分支
//
// @param request - GetBranchInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBranchInfoResponse
func (client *Client) GetBranchInfoWithContext(ctx context.Context, repositoryId *string, request *GetBranchInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBranchInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.BranchName) {
		query["branchName"] = request.BranchName
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBranchInfo"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/branches/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBranchInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询检查运行
//
// @param request - GetCheckRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCheckRunResponse
func (client *Client) GetCheckRunWithContext(ctx context.Context, request *GetCheckRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCheckRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.CheckRunId) {
		query["checkRunId"] = request.CheckRunId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCheckRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/check_runs/get_check_run"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCheckRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业信息
//
// @param request - GetCodeupOrganizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCodeupOrganizationResponse
func (client *Client) GetCodeupOrganizationWithContext(ctx context.Context, identity *string, request *GetCodeupOrganizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCodeupOrganizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCodeupOrganization"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/organization/" + dara.PercentEncode(dara.StringValue(identity))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCodeupOrganizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取比较详情
//
// @param request - GetCompareDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompareDetailResponse
func (client *Client) GetCompareDetailWithContext(ctx context.Context, repositoryId *string, request *GetCompareDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCompareDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.MaxDiffByte) {
		query["maxDiffByte"] = request.MaxDiffByte
	}

	if !dara.IsNil(request.MaxDiffFile) {
		query["maxDiffFile"] = request.MaxDiffFile
	}

	if !dara.IsNil(request.MergeBase) {
		query["mergeBase"] = request.MergeBase
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCompareDetail"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/commits/compare/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCompareDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义字段的选项值
//
// @param request - GetCustomFieldOptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomFieldOptionResponse
func (client *Client) GetCustomFieldOptionWithContext(ctx context.Context, organizationId *string, fieldId *string, request *GetCustomFieldOptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCustomFieldOptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SpaceIdentifier) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.SpaceType) {
		query["spaceType"] = request.SpaceType
	}

	if !dara.IsNil(request.WorkitemTypeIdentifier) {
		query["workitemTypeIdentifier"] = request.WorkitemTypeIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomFieldOption"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/fields/" + dara.PercentEncode(dara.StringValue(fieldId)) + "/getCustomOption"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomFieldOptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文件
//
// @param request - GetFileBlobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileBlobsResponse
func (client *Client) GetFileBlobsWithContext(ctx context.Context, repositoryId *string, request *GetFileBlobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFileBlobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.FilePath) {
		query["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Ref) {
		query["ref"] = request.Ref
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileBlobs"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/files/blobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileBlobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件上一次提交信息
//
// @param request - GetFileLastCommitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileLastCommitResponse
func (client *Client) GetFileLastCommitWithContext(ctx context.Context, repositoryId *string, request *GetFileLastCommitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFileLastCommitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.FilePath) {
		query["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Sha) {
		query["sha"] = request.Sha
	}

	if !dara.IsNil(request.ShowSignature) {
		query["showSignature"] = request.ShowSignature
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileLastCommit"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/files/lastCommit"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileLastCommitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标签分类
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowTagGroupResponse
func (client *Client) GetFlowTagGroupWithContext(ctx context.Context, organizationId *string, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFlowTagGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFlowTagGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/flow/tagGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowTagGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据路径查询代码组
//
// @param request - GetGroupByPathRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupByPathResponse
func (client *Client) GetGroupByPathWithContext(ctx context.Context, request *GetGroupByPathRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGroupByPathResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identity) {
		query["identity"] = request.Identity
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroupByPath"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/4/groups/find_by_path"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupByPathResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码组信息
//
// @param request - GetGroupDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupDetailResponse
func (client *Client) GetGroupDetailWithContext(ctx context.Context, request *GetGroupDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGroupDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGroupDetail"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/get_detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGroupDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主机组信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostGroupResponse
func (client *Client) GetHostGroupWithContext(ctx context.Context, organizationId *string, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHostGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHostGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/hostGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHostGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合并请求详情
//
// @param request - GetMergeRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMergeRequestResponse
func (client *Client) GetMergeRequestWithContext(ctx context.Context, repositoryId *string, localId *string, request *GetMergeRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMergeRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMergeRequest"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/merge_requests/" + dara.PercentEncode(dara.StringValue(localId)) + "/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMergeRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合并请求的变更信息
//
// @param request - GetMergeRequestChangeTreeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMergeRequestChangeTreeResponse
func (client *Client) GetMergeRequestChangeTreeWithContext(ctx context.Context, request *GetMergeRequestChangeTreeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMergeRequestChangeTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.FromPatchSetBizId) {
		query["fromPatchSetBizId"] = request.FromPatchSetBizId
	}

	if !dara.IsNil(request.LocalId) {
		query["localId"] = request.LocalId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	if !dara.IsNil(request.ToPatchSetBizId) {
		query["toPatchSetBizId"] = request.ToPatchSetBizId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMergeRequestChangeTree"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/merge_requests/diffs/change_tree"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMergeRequestChangeTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业成员
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationMemberResponse
func (client *Client) GetOrganizationMemberWithContext(ctx context.Context, organizationId *string, accountId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrganizationMemberResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrganizationMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/members/" + dara.PercentEncode(dara.StringValue(accountId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrganizationMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineResponse
func (client *Client) GetPipelineWithContext(ctx context.Context, organizationId *string, pipelineId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipeline"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取构建物下载链接
//
// @param request - GetPipelineArtifactUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineArtifactUrlResponse
func (client *Client) GetPipelineArtifactUrlWithContext(ctx context.Context, organizationId *string, request *GetPipelineArtifactUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineArtifactUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FilePath) {
		query["filePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineArtifactUrl"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipeline/getArtifactDownloadUrl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineArtifactUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取emase构建物下载链接
//
// @param request - GetPipelineEmasArtifactUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineEmasArtifactUrlResponse
func (client *Client) GetPipelineEmasArtifactUrlWithContext(ctx context.Context, organizationId *string, emasJobInstanceId *string, md5 *string, pipelineId *string, pipelineRunId *string, request *GetPipelineEmasArtifactUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineEmasArtifactUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceConnectionId) {
		query["serviceConnectionId"] = request.ServiceConnectionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineEmasArtifactUrl"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRun/" + dara.PercentEncode(dara.StringValue(pipelineRunId)) + "/emas/artifact/" + dara.PercentEncode(dara.StringValue(emasJobInstanceId)) + "/" + dara.PercentEncode(dara.StringValue(md5))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineEmasArtifactUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线分组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineGroupResponse
func (client *Client) GetPipelineGroupWithContext(ctx context.Context, organizationId *string, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelineGroups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线运行信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineRunResponse
func (client *Client) GetPipelineRunWithContext(ctx context.Context, organizationId *string, pipelineId *string, pipelineRunId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineRunResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRuns/" + dara.PercentEncode(dara.StringValue(pipelineRunId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取扫描报告下载链接
//
// @param request - GetPipelineScanReportUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineScanReportUrlResponse
func (client *Client) GetPipelineScanReportUrlWithContext(ctx context.Context, organizationId *string, request *GetPipelineScanReportUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineScanReportUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ReportPath) {
		body["reportPath"] = request.ReportPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineScanReportUrl"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipeline/getPipelineScanReportUrl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineScanReportUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据id获取项目详情-Projex
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectInfoResponse
func (client *Client) GetProjectInfoWithContext(ctx context.Context, organizationId *string, projectId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectInfo"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/project/" + dara.PercentEncode(dara.StringValue(projectId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库成员
//
// @param request - GetProjectMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectMemberResponse
func (client *Client) GetProjectMemberWithContext(ctx context.Context, repositoryId *string, aliyunPk *string, request *GetProjectMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/members/get/" + dara.PercentEncode(dara.StringValue(aliyunPk))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取研发阶段流水线运行实例
//
// @param request - GetReleaseStagePipelineRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReleaseStagePipelineRunResponse
func (client *Client) GetReleaseStagePipelineRunWithContext(ctx context.Context, appName *string, releaseWorkflowSn *string, releaseStageSn *string, executionNumber *string, request *GetReleaseStagePipelineRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReleaseStagePipelineRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReleaseStagePipelineRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/releaseWorkflows/" + dara.PercentEncode(dara.StringValue(releaseWorkflowSn)) + "/releaseStages/" + dara.PercentEncode(dara.StringValue(releaseStageSn)) + "/executions/" + dara.PercentEncode(dara.StringValue(executionNumber)) + "%3AgetPipelineRun"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReleaseStagePipelineRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 使用代码库ID或路径查询代码库信息
//
// @param request - GetRepositoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepositoryResponse
func (client *Client) GetRepositoryWithContext(ctx context.Context, request *GetRepositoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.Identity) {
		query["identity"] = request.Identity
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepository"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库提交信息
//
// @param request - GetRepositoryCommitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepositoryCommitResponse
func (client *Client) GetRepositoryCommitWithContext(ctx context.Context, repositoryId *string, sha *string, request *GetRepositoryCommitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRepositoryCommitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.ShowSignature) {
		query["showSignature"] = request.ShowSignature
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepositoryCommit"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/commits/" + dara.PercentEncode(dara.StringValue(sha))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepositoryCommitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个标签
//
// @param request - GetRepositoryTagRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepositoryTagResponse
func (client *Client) GetRepositoryTagWithContext(ctx context.Context, repositoryId *string, request *GetRepositoryTagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRepositoryTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.TagName) {
		query["tagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRepositoryTag"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/tag/info"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRepositoryTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预览代码片段
//
// @param request - GetSearchCodePreviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSearchCodePreviewResponse
func (client *Client) GetSearchCodePreviewWithContext(ctx context.Context, request *GetSearchCodePreviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSearchCodePreviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		query["docId"] = request.DocId
	}

	if !dara.IsNil(request.IsDsl) {
		query["isDsl"] = request.IsDsl
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSearchCodePreview"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/search/code_preview"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSearchCodePreviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取迭代详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSprintInfoResponse
func (client *Client) GetSprintInfoWithContext(ctx context.Context, organizationId *string, sprintId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSprintInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSprintInfo"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/sprints/" + dara.PercentEncode(dara.StringValue(sprintId)) + "/getSprintinfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSprintInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取测试计划中的测试用例列表
//
// @param request - GetTestResultListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTestResultListResponse
func (client *Client) GetTestResultListWithContext(ctx context.Context, organizationId *string, testPlanIdentifier *string, request *GetTestResultListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTestResultListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Conditions) {
		body["conditions"] = request.Conditions
	}

	if !dara.IsNil(request.DirectoryIdentifier) {
		body["directoryIdentifier"] = request.DirectoryIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTestResultList"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/testhub/testplan/" + dara.PercentEncode(dara.StringValue(testPlanIdentifier)) + "/testresults"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTestResultListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取测试用例列表
//
// @param request - GetTestcaseListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTestcaseListResponse
func (client *Client) GetTestcaseListWithContext(ctx context.Context, organizationId *string, request *GetTestcaseListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTestcaseListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Conditions) {
		body["conditions"] = request.Conditions
	}

	if !dara.IsNil(request.DirectoryIdentifier) {
		body["directoryIdentifier"] = request.DirectoryIdentifier
	}

	if !dara.IsNil(request.MaxResult) {
		body["maxResult"] = request.MaxResult
	}

	if !dara.IsNil(request.NextToken) {
		body["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SpaceIdentifier) {
		body["spaceIdentifier"] = request.SpaceIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTestcaseList"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/testhub/testcases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTestcaseListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前用户信息
//
// @param request - GetUserInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfoWithContext(ctx context.Context, request *GetUserInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserInfo"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/users/current"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取部署单信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVMDeployOrderResponse
func (client *Client) GetVMDeployOrderWithContext(ctx context.Context, organizationId *string, pipelineId *string, deployOrderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVMDeployOrderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVMDeployOrder"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/deploy/" + dara.PercentEncode(dara.StringValue(deployOrderId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVMDeployOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取变量组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVariableGroupResponse
func (client *Client) GetVariableGroupWithContext(ctx context.Context, organizationId *string, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVariableGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVariableGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/variableGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVariableGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作项动态
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkItemActivityResponse
func (client *Client) GetWorkItemActivityWithContext(ctx context.Context, organizationId *string, workitemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkItemActivityResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkItemActivity"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/" + dara.PercentEncode(dara.StringValue(workitemId)) + "/getActivity"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkItemActivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据id获取工作项详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkItemInfoResponse
func (client *Client) GetWorkItemInfoWithContext(ctx context.Context, organizationId *string, workitemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkItemInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkItemInfo"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/" + dara.PercentEncode(dara.StringValue(workitemId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkItemInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作项工作流信息
//
// @param request - GetWorkItemWorkFlowInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkItemWorkFlowInfoResponse
func (client *Client) GetWorkItemWorkFlowInfoWithContext(ctx context.Context, organizationId *string, workitemId *string, request *GetWorkItemWorkFlowInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkItemWorkFlowInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigurationId) {
		query["configurationId"] = request.ConfigurationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkItemWorkFlowInfo"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/" + dara.PercentEncode(dara.StringValue(workitemId)) + "/getWorkflowInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkItemWorkFlowInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取附件上传的元信息
//
// @param request - GetWorkitemAttachmentCreatemetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkitemAttachmentCreatemetaResponse
func (client *Client) GetWorkitemAttachmentCreatemetaWithContext(ctx context.Context, organizationId *string, workitemIdentifier *string, request *GetWorkitemAttachmentCreatemetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkitemAttachmentCreatemetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkitemAttachmentCreatemeta"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitem/" + dara.PercentEncode(dara.StringValue(workitemIdentifier)) + "/attachment/createmeta"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkitemAttachmentCreatemetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得所有评论
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkitemCommentListResponse
func (client *Client) GetWorkitemCommentListWithContext(ctx context.Context, organizationId *string, workitemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkitemCommentListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkitemCommentList"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/" + dara.PercentEncode(dara.StringValue(workitemId)) + "/commentList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkitemCommentListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作项文件信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkitemFileResponse
func (client *Client) GetWorkitemFileWithContext(ctx context.Context, organizationId *string, workitemIdentifier *string, fileIdentifier *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkitemFileResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkitemFile"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitem/" + dara.PercentEncode(dara.StringValue(workitemIdentifier)) + "/files/" + dara.PercentEncode(dara.StringValue(fileIdentifier))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkitemFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得一个工作项的指定关联项
//
// @param request - GetWorkitemRelationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkitemRelationsResponse
func (client *Client) GetWorkitemRelationsWithContext(ctx context.Context, organizationId *string, workitemId *string, request *GetWorkitemRelationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkitemRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RelationType) {
		query["relationType"] = request.RelationType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkitemRelations"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/" + dara.PercentEncode(dara.StringValue(workitemId)) + "/getRelations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkitemRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得一个企业下所有工时类型
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkitemTimeTypeListResponse
func (client *Client) GetWorkitemTimeTypeListWithContext(ctx context.Context, organizationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkitemTimeTypeListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkitemTimeTypeList"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/type/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkitemTimeTypeListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 加入流水线分组
//
// @param request - JoinPipelineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinPipelineGroupResponse
func (client *Client) JoinPipelineGroupWithContext(ctx context.Context, organizationId *string, request *JoinPipelineGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *JoinPipelineGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.PipelineIds) {
		query["pipelineIds"] = request.PipelineIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JoinPipelineGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelineGroups/join"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinPipelineGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关联合并请求Label
//
// @param request - LinkMergeRequestLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LinkMergeRequestLabelResponse
func (client *Client) LinkMergeRequestLabelWithContext(ctx context.Context, request *LinkMergeRequestLabelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *LinkMergeRequestLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.LocalId) {
		query["localId"] = request.LocalId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelIds) {
		body["labelIds"] = request.LabelIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LinkMergeRequestLabel"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/merge_requests/link_labels"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &LinkMergeRequestLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查找应用下所有的研发流程
//
// @param request - ListAllReleaseWorkflowsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllReleaseWorkflowsResponse
func (client *Client) ListAllReleaseWorkflowsWithContext(ctx context.Context, appName *string, request *ListAllReleaseWorkflowsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAllReleaseWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllReleaseWorkflows"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/releaseWorkflows"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &ListAllReleaseWorkflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询研发阶段执行记录集成变更信息
//
// @param request - ListAppReleaseStageExecutionIntegratedMetadataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppReleaseStageExecutionIntegratedMetadataResponse
func (client *Client) ListAppReleaseStageExecutionIntegratedMetadataWithContext(ctx context.Context, appName *string, releaseWorkflowSn *string, releaseStageSn *string, executionNumber *string, request *ListAppReleaseStageExecutionIntegratedMetadataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAppReleaseStageExecutionIntegratedMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppReleaseStageExecutionIntegratedMetadata"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/releaseWorkflows/" + dara.PercentEncode(dara.StringValue(releaseWorkflowSn)) + "/releaseStages/" + dara.PercentEncode(dara.StringValue(releaseStageSn)) + "/executions/" + dara.PercentEncode(dara.StringValue(executionNumber)) + "/integratedMetadata"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &ListAppReleaseStageExecutionIntegratedMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询研发阶段执行记录
//
// @param request - ListAppReleaseStageExecutionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppReleaseStageExecutionsResponse
func (client *Client) ListAppReleaseStageExecutionsWithContext(ctx context.Context, appName *string, releaseWorkflowSn *string, releaseStageSn *string, request *ListAppReleaseStageExecutionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAppReleaseStageExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Pagination) {
		query["pagination"] = request.Pagination
	}

	if !dara.IsNil(request.PerPage) {
		query["perPage"] = request.PerPage
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppReleaseStageExecutions"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/releaseWorkflows/" + dara.PercentEncode(dara.StringValue(releaseWorkflowSn)) + "/releaseStages/" + dara.PercentEncode(dara.StringValue(releaseStageSn)) + "/executions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppReleaseStageExecutionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查找应用成员列表
//
// @param request - ListApplicationMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationMembersResponse
func (client *Client) ListApplicationMembersWithContext(ctx context.Context, appName *string, request *ListApplicationMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApplicationMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationMembers"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/members"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查找应用详情
//
// @param request - ListApplicationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithContext(ctx context.Context, request *ListApplicationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Pagination) {
		query["pagination"] = request.Pagination
	}

	if !dara.IsNil(request.PerPage) {
		query["perPage"] = request.PerPage
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplications"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps%3Asearch"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 查询变更研发流程运行记录
//
// @param request - ListChangeRequestWorkflowExecutionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChangeRequestWorkflowExecutionsResponse
func (client *Client) ListChangeRequestWorkflowExecutionsWithContext(ctx context.Context, appName *string, sn *string, request *ListChangeRequestWorkflowExecutionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListChangeRequestWorkflowExecutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PerPage) {
		query["perPage"] = request.PerPage
	}

	if !dara.IsNil(request.ReleaseStageSn) {
		query["releaseStageSn"] = request.ReleaseStageSn
	}

	if !dara.IsNil(request.ReleaseWorkflowSn) {
		query["releaseWorkflowSn"] = request.ReleaseWorkflowSn
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChangeRequestWorkflowExecutions"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/changeRequests/" + dara.PercentEncode(dara.StringValue(sn)) + "/executions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChangeRequestWorkflowExecutionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询变更列表
//
// @param tmpReq - ListChangeRequestsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChangeRequestsResponse
func (client *Client) ListChangeRequestsWithContext(ctx context.Context, appName *string, tmpReq *ListChangeRequestsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListChangeRequestsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListChangeRequestsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppNameList) {
		request.AppNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppNameList, dara.String("appNameList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OwnerIdList) {
		request.OwnerIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OwnerIdList, dara.String("ownerIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StateList) {
		request.StateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StateList, dara.String("stateList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppNameListShrink) {
		query["appNameList"] = request.AppNameListShrink
	}

	if !dara.IsNil(request.DisplayNameKeyword) {
		query["displayNameKeyword"] = request.DisplayNameKeyword
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.OwnerIdListShrink) {
		query["ownerIdList"] = request.OwnerIdListShrink
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Pagination) {
		query["pagination"] = request.Pagination
	}

	if !dara.IsNil(request.PerPage) {
		query["perPage"] = request.PerPage
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	if !dara.IsNil(request.StateListShrink) {
		query["stateList"] = request.StateListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChangeRequests"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/changeRequests"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChangeRequestsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询检查运行列表
//
// @param request - ListCheckRunsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCheckRunsResponse
func (client *Client) ListCheckRunsWithContext(ctx context.Context, request *ListCheckRunsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCheckRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Ref) {
		query["ref"] = request.Ref
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCheckRuns"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/check_runs/list_check_runs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCheckRunsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询提交状态列表
//
// @param request - ListCommitStatusesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommitStatusesResponse
func (client *Client) ListCommitStatusesWithContext(ctx context.Context, request *ListCommitStatusesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCommitStatusesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	if !dara.IsNil(request.Sha) {
		query["sha"] = request.Sha
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCommitStatuses"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/repository/commit_statuses/list_commit_statuses"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCommitStatusesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标签分类列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowTagGroupsResponse
func (client *Client) ListFlowTagGroupsWithContext(ctx context.Context, organizationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFlowTagGroupsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlowTagGroups"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/flow/tagGroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowTagGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组成员列表
//
// @param request - ListGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupMemberResponse
func (client *Client) ListGroupMemberWithContext(ctx context.Context, groupId *string, request *ListGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码组下的库列表
//
// @param request - ListGroupRepositoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupRepositoriesResponse
func (client *Client) ListGroupRepositoriesWithContext(ctx context.Context, groupId *string, request *ListGroupRepositoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGroupRepositoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroupRepositories"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/projects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupRepositoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主机组列表
//
// @param request - ListHostGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHostGroupsResponse
func (client *Client) ListHostGroupsWithContext(ctx context.Context, organizationId *string, request *ListHostGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHostGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateEndTime) {
		query["createEndTime"] = request.CreateEndTime
	}

	if !dara.IsNil(request.CreateStartTime) {
		query["createStartTime"] = request.CreateStartTime
	}

	if !dara.IsNil(request.CreatorAccountIds) {
		query["creatorAccountIds"] = request.CreatorAccountIds
	}

	if !dara.IsNil(request.Ids) {
		query["ids"] = request.Ids
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageOrder) {
		query["pageOrder"] = request.PageOrder
	}

	if !dara.IsNil(request.PageSort) {
		query["pageSort"] = request.PageSort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHostGroups"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/hostGroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 当前用户加入的企业列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJoinedOrganizationsResponse
func (client *Client) ListJoinedOrganizationsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJoinedOrganizationsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJoinedOrganizations"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/users/joinedOrgs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJoinedOrganizationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合并请求评论列表
//
// @param request - ListMergeRequestCommentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMergeRequestCommentsResponse
func (client *Client) ListMergeRequestCommentsWithContext(ctx context.Context, request *ListMergeRequestCommentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMergeRequestCommentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.LocalId) {
		query["localId"] = request.LocalId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CommentType) {
		body["commentType"] = request.CommentType
	}

	if !dara.IsNil(request.FilePath) {
		body["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.PatchSetBizIds) {
		body["patchSetBizIds"] = request.PatchSetBizIds
	}

	if !dara.IsNil(request.Resolved) {
		body["resolved"] = request.Resolved
	}

	if !dara.IsNil(request.State) {
		body["state"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMergeRequestComments"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/merge_requests/comments/list_comments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMergeRequestCommentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合并请求文件已读列表
//
// @param request - ListMergeRequestFilesReadsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMergeRequestFilesReadsResponse
func (client *Client) ListMergeRequestFilesReadsWithContext(ctx context.Context, request *ListMergeRequestFilesReadsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMergeRequestFilesReadsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.FromPatchSetBizId) {
		query["fromPatchSetBizId"] = request.FromPatchSetBizId
	}

	if !dara.IsNil(request.LocalId) {
		query["localId"] = request.LocalId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	if !dara.IsNil(request.ToPatchSetBizId) {
		query["toPatchSetBizId"] = request.ToPatchSetBizId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMergeRequestFilesReads"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/merge_requests/diffs/files_read_infos"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMergeRequestFilesReadsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合并请求Label列表
//
// @param request - ListMergeRequestLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMergeRequestLabelsResponse
func (client *Client) ListMergeRequestLabelsWithContext(ctx context.Context, request *ListMergeRequestLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMergeRequestLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.LocalId) {
		query["localId"] = request.LocalId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMergeRequestLabels"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/merge_requests/labels"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMergeRequestLabelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询合并请求的版本列表
//
// @param request - ListMergeRequestPatchSetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMergeRequestPatchSetsResponse
func (client *Client) ListMergeRequestPatchSetsWithContext(ctx context.Context, request *ListMergeRequestPatchSetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMergeRequestPatchSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.LocalId) {
		query["localId"] = request.LocalId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMergeRequestPatchSets"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/merge_requests/diffs/list_patchsets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMergeRequestPatchSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码评审列表
//
// @param request - ListMergeRequestsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMergeRequestsResponse
func (client *Client) ListMergeRequestsWithContext(ctx context.Context, request *ListMergeRequestsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMergeRequestsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.AuthorIds) {
		query["authorIds"] = request.AuthorIds
	}

	if !dara.IsNil(request.CreatedAfter) {
		query["createdAfter"] = request.CreatedAfter
	}

	if !dara.IsNil(request.CreatedBefore) {
		query["createdBefore"] = request.CreatedBefore
	}

	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.GroupIds) {
		query["groupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.LabelIds) {
		query["labelIds"] = request.LabelIds
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectIds) {
		query["projectIds"] = request.ProjectIds
	}

	if !dara.IsNil(request.ReviewerIds) {
		query["reviewerIds"] = request.ReviewerIds
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	if !dara.IsNil(request.State) {
		query["state"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMergeRequests"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/merge_requests/advanced_search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMergeRequestsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业成员列表
//
// @param request - ListOrganizationMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationMembersResponse
func (client *Client) ListOrganizationMembersWithContext(ctx context.Context, organizationId *string, request *ListOrganizationMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOrganizationMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainsExternInfo) {
		query["containsExternInfo"] = request.ContainsExternInfo
	}

	if !dara.IsNil(request.ExternUid) {
		query["externUid"] = request.ExternUid
	}

	if !dara.IsNil(request.JoinTimeFrom) {
		query["joinTimeFrom"] = request.JoinTimeFrom
	}

	if !dara.IsNil(request.JoinTimeTo) {
		query["joinTimeTo"] = request.JoinTimeTo
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrganizationMemberName) {
		query["organizationMemberName"] = request.OrganizationMemberName
	}

	if !dara.IsNil(request.Provider) {
		query["provider"] = request.Provider
	}

	if !dara.IsNil(request.State) {
		query["state"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationMembers"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/members"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户企业列表
//
// @param request - ListOrganizationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationsResponse
func (client *Client) ListOrganizationsWithContext(ctx context.Context, request *ListOrganizationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOrganizationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessLevel) {
		query["accessLevel"] = request.AccessLevel
	}

	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.MinAccessLevel) {
		query["minAccessLevel"] = request.MinAccessLevel
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizations"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organizations/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrganizationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线分组下流水线列表列表
//
// @param request - ListPipelineGroupPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineGroupPipelinesResponse
func (client *Client) ListPipelineGroupPipelinesWithContext(ctx context.Context, organizationId *string, groupId *string, request *ListPipelineGroupPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineGroupPipelinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateEndTime) {
		query["createEndTime"] = request.CreateEndTime
	}

	if !dara.IsNil(request.CreateStartTime) {
		query["createStartTime"] = request.CreateStartTime
	}

	if !dara.IsNil(request.ExecuteEndTime) {
		query["executeEndTime"] = request.ExecuteEndTime
	}

	if !dara.IsNil(request.ExecuteStartTime) {
		query["executeStartTime"] = request.ExecuteStartTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PipelineName) {
		query["pipelineName"] = request.PipelineName
	}

	if !dara.IsNil(request.ResultStatusList) {
		query["resultStatusList"] = request.ResultStatusList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelineGroupPipelines"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelineGroups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/pipelines"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineGroupPipelinesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线分组列表
//
// @param request - ListPipelineGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineGroupsResponse
func (client *Client) ListPipelineGroupsWithContext(ctx context.Context, organizationId *string, request *ListPipelineGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineGroupsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelineGroups"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelineGroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线运行过的任务历史
//
// @param request - ListPipelineJobHistorysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineJobHistorysResponse
func (client *Client) ListPipelineJobHistorysWithContext(ctx context.Context, organizationId *string, pipelineId *string, request *ListPipelineJobHistorysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineJobHistorysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.Identifier) {
		query["identifier"] = request.Identifier
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelineJobHistorys"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/job/historys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineJobHistorysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线运行过的任务
//
// @param request - ListPipelineJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineJobsResponse
func (client *Client) ListPipelineJobsWithContext(ctx context.Context, organizationId *string, pipelineId *string, request *ListPipelineJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelineJobs"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线关联列表
//
// @param request - ListPipelineRelationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineRelationsResponse
func (client *Client) ListPipelineRelationsWithContext(ctx context.Context, organizationId *string, pipelineId *string, request *ListPipelineRelationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RelObjectType) {
		query["relObjectType"] = request.RelObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelineRelations"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRelations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线运行历史
//
// @param request - ListPipelineRunsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineRunsResponse
func (client *Client) ListPipelineRunsWithContext(ctx context.Context, organizationId *string, pipelineId *string, request *ListPipelineRunsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TriggerMode) {
		query["triggerMode"] = request.TriggerMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelineRuns"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRuns"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线列表
//
// @param request - ListPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelinesResponse
func (client *Client) ListPipelinesWithContext(ctx context.Context, organizationId *string, request *ListPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateEndTime) {
		query["createEndTime"] = request.CreateEndTime
	}

	if !dara.IsNil(request.CreateStartTime) {
		query["createStartTime"] = request.CreateStartTime
	}

	if !dara.IsNil(request.CreatorAccountIds) {
		query["creatorAccountIds"] = request.CreatorAccountIds
	}

	if !dara.IsNil(request.ExecuteAccountIds) {
		query["executeAccountIds"] = request.ExecuteAccountIds
	}

	if !dara.IsNil(request.ExecuteEndTime) {
		query["executeEndTime"] = request.ExecuteEndTime
	}

	if !dara.IsNil(request.ExecuteStartTime) {
		query["executeStartTime"] = request.ExecuteStartTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PipelineName) {
		query["pipelineName"] = request.PipelineName
	}

	if !dara.IsNil(request.StatusList) {
		query["statusList"] = request.StatusList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelines"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库Label列表
//
// @param request - ListProjectLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectLabelsResponse
func (client *Client) ListProjectLabelsWithContext(ctx context.Context, request *ListProjectLabelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	if !dara.IsNil(request.WithCounts) {
		query["withCounts"] = request.WithCounts
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectLabels"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/labels"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectLabelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据项目id获取项目所以成员
//
// @param request - ListProjectMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithContext(ctx context.Context, organizationId *string, projectId *string, request *ListProjectMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TargetType) {
		query["targetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectMembers"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/listMembers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目模板列表
//
// @param request - ListProjectTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectTemplatesResponse
func (client *Client) ListProjectTemplatesWithContext(ctx context.Context, organizationId *string, request *ListProjectTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectTemplates"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/projects/listTemplates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目下开启的工作项类型
//
// @param request - ListProjectWorkitemTypesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectWorkitemTypesResponse
func (client *Client) ListProjectWorkitemTypesWithContext(ctx context.Context, organizationId *string, projectId *string, request *ListProjectWorkitemTypesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectWorkitemTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.SpaceType) {
		query["spaceType"] = request.SpaceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectWorkitemTypes"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/getWorkitemType"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectWorkitemTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目列表
//
// @param request - ListProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, organizationId *string, request *ListProjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.Conditions) {
		query["conditions"] = request.Conditions
	}

	if !dara.IsNil(request.ExtraConditions) {
		query["extraConditions"] = request.ExtraConditions
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Scope) {
		query["scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/listProjects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询保护分支列表
//
// @param request - ListProtectedBranchesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProtectedBranchesResponse
func (client *Client) ListProtectedBranchesWithContext(ctx context.Context, repositoryId *string, request *ListProtectedBranchesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProtectedBranchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProtectedBranches"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/protect_branches"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProtectedBranchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询推送规则列表
//
// @param request - ListPushRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPushRulesResponse
func (client *Client) ListPushRulesWithContext(ctx context.Context, repositoryId *string, request *ListPushRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPushRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPushRules"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/push_rule/push_rules/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPushRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库列表
//
// @param request - ListRepositoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoriesResponse
func (client *Client) ListRepositoriesWithContext(ctx context.Context, request *ListRepositoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.Archived) {
		query["archived"] = request.Archived
	}

	if !dara.IsNil(request.MinAccessLevel) {
		query["minAccessLevel"] = request.MinAccessLevel
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PerPage) {
		query["perPage"] = request.PerPage
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositories"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询分支列表
//
// @param request - ListRepositoryBranchesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryBranchesResponse
func (client *Client) ListRepositoryBranchesWithContext(ctx context.Context, repositoryId *string, request *ListRepositoryBranchesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoryBranchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositoryBranches"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/branches"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryBranchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库单个提交（Commit）的提交内容
//
// @param request - ListRepositoryCommitDiffRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryCommitDiffResponse
func (client *Client) ListRepositoryCommitDiffWithContext(ctx context.Context, repositoryId *string, sha *string, request *ListRepositoryCommitDiffRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoryCommitDiffResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.ContextLine) {
		query["contextLine"] = request.ContextLine
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositoryCommitDiff"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/commits/" + dara.PercentEncode(dara.StringValue(sha)) + "/diff"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryCommitDiffResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库提交历史
//
// @param request - ListRepositoryCommitsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryCommitsResponse
func (client *Client) ListRepositoryCommitsWithContext(ctx context.Context, repositoryId *string, request *ListRepositoryCommitsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoryCommitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.End) {
		query["end"] = request.End
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["path"] = request.Path
	}

	if !dara.IsNil(request.RefName) {
		query["refName"] = request.RefName
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.ShowCommentsCount) {
		query["showCommentsCount"] = request.ShowCommentsCount
	}

	if !dara.IsNil(request.ShowSignature) {
		query["showSignature"] = request.ShowSignature
	}

	if !dara.IsNil(request.Start) {
		query["start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositoryCommits"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/commits"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryCommitsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码组列表
//
// @param request - ListRepositoryGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryGroupsResponse
func (client *Client) ListRepositoryGroupsWithContext(ctx context.Context, request *ListRepositoryGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoryGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.IncludePersonal) {
		query["includePersonal"] = request.IncludePersonal
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentId) {
		query["parentId"] = request.ParentId
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositoryGroups"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/get/all"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库成员列表
//
// @param request - ListRepositoryMemberWithInheritedRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryMemberWithInheritedResponse
func (client *Client) ListRepositoryMemberWithInheritedWithContext(ctx context.Context, repositoryId *string, request *ListRepositoryMemberWithInheritedRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoryMemberWithInheritedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositoryMemberWithInherited"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/members/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryMemberWithInheritedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询标签列表
//
// @param request - ListRepositoryTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryTagsResponse
func (client *Client) ListRepositoryTagsWithContext(ctx context.Context, repositoryId *string, request *ListRepositoryTagsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoryTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Search) {
		query["search"] = request.Search
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositoryTags"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/tag/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库文件树
//
// @param request - ListRepositoryTreeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryTreeResponse
func (client *Client) ListRepositoryTreeWithContext(ctx context.Context, repositoryId *string, request *ListRepositoryTreeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoryTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Path) {
		query["path"] = request.Path
	}

	if !dara.IsNil(request.RefName) {
		query["refName"] = request.RefName
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositoryTree"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/files/tree"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询代码库Webhook列表
//
// @param request - ListRepositoryWebhookRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryWebhookResponse
func (client *Client) ListRepositoryWebhookWithContext(ctx context.Context, repositoryId *string, request *ListRepositoryWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRepositoryWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRepositoryWebhook"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/webhooks/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRepositoryWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源成员列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceMembersResponse
func (client *Client) ListResourceMembersWithContext(ctx context.Context, organizationId *string, resourceType *string, resourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceMembersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceMembers"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/" + dara.PercentEncode(dara.StringValue(resourceType)) + "/" + dara.PercentEncode(dara.StringValue(resourceId)) + "/members"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索代码提交数据
//
// @param request - ListSearchCommitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchCommitResponse
func (client *Client) ListSearchCommitWithContext(ctx context.Context, request *ListSearchCommitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSearchCommitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoPath) {
		body["repoPath"] = request.RepoPath
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.Sort) {
		body["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchCommit"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/search/commit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchCommitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索代码仓库数据
//
// @param request - ListSearchRepositoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchRepositoryResponse
func (client *Client) ListSearchRepositoryWithContext(ctx context.Context, request *ListSearchRepositoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSearchRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunPk) {
		body["aliyunPk"] = request.AliyunPk
	}

	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoPath) {
		body["repoPath"] = request.RepoPath
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.Sort) {
		body["sort"] = request.Sort
	}

	if !dara.IsNil(request.VisibilityLevel) {
		body["visibilityLevel"] = request.VisibilityLevel
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchRepository"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/search/repo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索代码片段
//
// @param request - ListSearchSourceCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchSourceCodeResponse
func (client *Client) ListSearchSourceCodeWithContext(ctx context.Context, request *ListSearchSourceCodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSearchSourceCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilePath) {
		body["filePath"] = request.FilePath
	}

	if !dara.IsNil(request.IsCodeBlock) {
		body["isCodeBlock"] = request.IsCodeBlock
	}

	if !dara.IsNil(request.Keyword) {
		body["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RepoPath) {
		body["repoPath"] = request.RepoPath
	}

	if !dara.IsNil(request.Scope) {
		body["scope"] = request.Scope
	}

	if !dara.IsNil(request.Sort) {
		body["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchSourceCode"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/search/code"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchSourceCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务授权列表
//
// @param request - ListServiceAuthsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceAuthsResponse
func (client *Client) ListServiceAuthsWithContext(ctx context.Context, organizationId *string, request *ListServiceAuthsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceAuthsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceAuthType) {
		query["serviceAuthType"] = request.ServiceAuthType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceAuths"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/serviceAuths"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceAuthsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务连接列表
//
// @param request - ListServiceConnectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceConnectionsResponse
func (client *Client) ListServiceConnectionsWithContext(ctx context.Context, organizationId *string, request *ListServiceConnectionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceConnectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SericeConnectionType) {
		query["sericeConnectionType"] = request.SericeConnectionType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceConnections"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/serviceConnections"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceConnectionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务证书列表
//
// @param request - ListServiceCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceCredentialsResponse
func (client *Client) ListServiceCredentialsWithContext(ctx context.Context, organizationId *string, request *ListServiceCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceCredentialType) {
		query["serviceCredentialType"] = request.ServiceCredentialType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceCredentials"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/serviceCredentials"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取迭代列表
//
// @param request - ListSprintsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSprintsResponse
func (client *Client) ListSprintsWithContext(ctx context.Context, organizationId *string, request *ListSprintsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSprintsResponse, _err error) {
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

	if !dara.IsNil(request.SpaceIdentifier) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.SpaceType) {
		query["spaceType"] = request.SpaceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSprints"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/sprints/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSprintsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取测试用例全部字段
//
// @param request - ListTestCaseFieldsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTestCaseFieldsResponse
func (client *Client) ListTestCaseFieldsWithContext(ctx context.Context, organizationId *string, request *ListTestCaseFieldsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTestCaseFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SpaceIdentifier) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTestCaseFields"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/testhub/testcase/fields"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTestCaseFieldsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 测试DrawService
//
// @param request - ListUserDrawRecordByPkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDrawRecordByPkResponse
func (client *Client) ListUserDrawRecordByPkWithContext(ctx context.Context, request *ListUserDrawRecordByPkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserDrawRecordByPkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunPk) {
		query["aliyunPk"] = request.AliyunPk
	}

	if !dara.IsNil(request.DrawGroup) {
		query["drawGroup"] = request.DrawGroup
	}

	if !dara.IsNil(request.DrawPoolName) {
		query["drawPoolName"] = request.DrawPoolName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserDrawRecordByPk"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/listUserDrawRecords"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserDrawRecordByPkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前用户的SSH Key列表
//
// @param request - ListUserKeysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserKeysResponse
func (client *Client) ListUserKeysWithContext(ctx context.Context, request *ListUserKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Sort) {
		query["sort"] = request.Sort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserKeys"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v3/user/keys"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户有权限的资源（代码库、组）
//
// @param request - ListUserResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserResourcesResponse
func (client *Client) ListUserResourcesWithContext(ctx context.Context, request *ListUserResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserIds) {
		query["userIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserResources"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/user/vision/user_resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取变量组列表
//
// @param request - ListVariableGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVariableGroupsResponse
func (client *Client) ListVariableGroupsWithContext(ctx context.Context, organizationId *string, request *ListVariableGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVariableGroupsResponse, _err error) {
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

	if !dara.IsNil(request.PageOrder) {
		query["pageOrder"] = request.PageOrder
	}

	if !dara.IsNil(request.PageSort) {
		query["pageSort"] = request.PageSort
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVariableGroups"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/variableGroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVariableGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目下工作项的所有字段
//
// @param request - ListWorkItemAllFieldsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkItemAllFieldsResponse
func (client *Client) ListWorkItemAllFieldsWithContext(ctx context.Context, organizationId *string, request *ListWorkItemAllFieldsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkItemAllFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SpaceIdentifier) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.SpaceType) {
		query["spaceType"] = request.SpaceType
	}

	if !dara.IsNil(request.WorkitemTypeIdentifier) {
		query["workitemTypeIdentifier"] = request.WorkitemTypeIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkItemAllFields"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/fields/listAll"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkItemAllFieldsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作项工作流的所有状态
//
// @param request - ListWorkItemWorkFlowStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkItemWorkFlowStatusResponse
func (client *Client) ListWorkItemWorkFlowStatusWithContext(ctx context.Context, organizationId *string, request *ListWorkItemWorkFlowStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkItemWorkFlowStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SpaceIdentifier) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.SpaceType) {
		query["spaceType"] = request.SpaceType
	}

	if !dara.IsNil(request.WorkitemCategoryIdentifier) {
		query["workitemCategoryIdentifier"] = request.WorkitemCategoryIdentifier
	}

	if !dara.IsNil(request.WorkitemTypeIdentifier) {
		query["workitemTypeIdentifier"] = request.WorkitemTypeIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkItemWorkFlowStatus"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/workflow/listWorkflowStatus"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkItemWorkFlowStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作项的附件列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkitemAttachmentsResponse
func (client *Client) ListWorkitemAttachmentsWithContext(ctx context.Context, organizationId *string, workitemIdentifier *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkitemAttachmentsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkitemAttachments"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitem/" + dara.PercentEncode(dara.StringValue(workitemIdentifier)) + "/attachments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkitemAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作项预计工时明细列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkitemEstimateResponse
func (client *Client) ListWorkitemEstimateWithContext(ctx context.Context, organizationId *string, workitemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkitemEstimateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkitemEstimate"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/" + dara.PercentEncode(dara.StringValue(workitemId)) + "/estimate/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkitemEstimateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作项工时明细列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkitemTimeResponse
func (client *Client) ListWorkitemTimeWithContext(ctx context.Context, organizationId *string, workitemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkitemTimeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkitemTime"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/" + dara.PercentEncode(dara.StringValue(workitemId)) + "/time/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkitemTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作项列表
//
// @param request - ListWorkitemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkitemsResponse
func (client *Client) ListWorkitemsWithContext(ctx context.Context, organizationId *string, request *ListWorkitemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkitemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["category"] = request.Category
	}

	if !dara.IsNil(request.Conditions) {
		query["conditions"] = request.Conditions
	}

	if !dara.IsNil(request.ExtraConditions) {
		query["extraConditions"] = request.ExtraConditions
	}

	if !dara.IsNil(request.GroupCondition) {
		query["groupCondition"] = request.GroupCondition
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.SearchType) {
		query["searchType"] = request.SearchType
	}

	if !dara.IsNil(request.SpaceIdentifier) {
		query["spaceIdentifier"] = request.SpaceIdentifier
	}

	if !dara.IsNil(request.SpaceType) {
		query["spaceType"] = request.SpaceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkitems"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/listWorkitems"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkitemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流水线运行任务日志
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LogPipelineJobRunResponse
func (client *Client) LogPipelineJobRunWithContext(ctx context.Context, organizationId *string, pipelineId *string, jobId *string, pipelineRunId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *LogPipelineJobRunResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("LogPipelineJobRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRun/" + dara.PercentEncode(dara.StringValue(pipelineRunId)) + "/job/" + dara.PercentEncode(dara.StringValue(jobId)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &LogPipelineJobRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取机器部署日志
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LogVMDeployMachineResponse
func (client *Client) LogVMDeployMachineWithContext(ctx context.Context, organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *LogVMDeployMachineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("LogVMDeployMachine"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/deploy/" + dara.PercentEncode(dara.StringValue(deployOrderId)) + "/machine/" + dara.PercentEncode(dara.StringValue(machineSn)) + "/log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &LogVMDeployMachineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合并代码评审
//
// @param request - MergeMergeRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MergeMergeRequestResponse
func (client *Client) MergeMergeRequestWithContext(ctx context.Context, repositoryId *string, localId *string, request *MergeMergeRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MergeMergeRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MergeMessage) {
		body["mergeMessage"] = request.MergeMessage
	}

	if !dara.IsNil(request.MergeType) {
		body["mergeType"] = request.MergeType
	}

	if !dara.IsNil(request.RemoveSourceBranch) {
		body["removeSourceBranch"] = request.RemoveSourceBranch
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MergeMergeRequest"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/merge_requests/" + dara.PercentEncode(dara.StringValue(localId)) + "/merge"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MergeMergeRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过人工卡点
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PassPipelineValidateResponse
func (client *Client) PassPipelineValidateWithContext(ctx context.Context, organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PassPipelineValidateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PassPipelineValidate"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRuns/" + dara.PercentEncode(dara.StringValue(pipelineRunId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/pass"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PassPipelineValidateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过人工卡点
//
// @param request - PassReleaseStagePipelineValidateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PassReleaseStagePipelineValidateResponse
func (client *Client) PassReleaseStagePipelineValidateWithContext(ctx context.Context, appName *string, releaseWorkflowSn *string, releaseStageSn *string, executionNumber *string, request *PassReleaseStagePipelineValidateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PassReleaseStagePipelineValidateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["jobId"] = request.JobId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PassReleaseStagePipelineValidate"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/releaseWorkflows/" + dara.PercentEncode(dara.StringValue(releaseWorkflowSn)) + "/releaseStages/" + dara.PercentEncode(dara.StringValue(releaseStageSn)) + "/executions/" + dara.PercentEncode(dara.StringValue(executionNumber)) + "%3ApassPipelineValidate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PassReleaseStagePipelineValidateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拒绝人工卡点
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefusePipelineValidateResponse
func (client *Client) RefusePipelineValidateWithContext(ctx context.Context, organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RefusePipelineValidateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefusePipelineValidate"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRuns/" + dara.PercentEncode(dara.StringValue(pipelineRunId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/refuse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RefusePipelineValidateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拒绝人工卡点
//
// @param request - RefuseReleaseStagePipelineValidateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefuseReleaseStagePipelineValidateResponse
func (client *Client) RefuseReleaseStagePipelineValidateWithContext(ctx context.Context, appName *string, releaseWorkflowSn *string, releaseStageSn *string, executionNumber *string, request *RefuseReleaseStagePipelineValidateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RefuseReleaseStagePipelineValidateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["jobId"] = request.JobId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefuseReleaseStagePipelineValidate"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/releaseWorkflows/" + dara.PercentEncode(dara.StringValue(releaseWorkflowSn)) + "/releaseStages/" + dara.PercentEncode(dara.StringValue(releaseStageSn)) + "/executions/" + dara.PercentEncode(dara.StringValue(executionNumber)) + "%3ArefusePipelineValidate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RefuseReleaseStagePipelineValidateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新打开代码评审
//
// @param request - ReopenMergeRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReopenMergeRequestResponse
func (client *Client) ReopenMergeRequestWithContext(ctx context.Context, repositoryId *string, localId *string, request *ReopenMergeRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReopenMergeRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReopenMergeRequest"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/merge_requests/" + dara.PercentEncode(dara.StringValue(localId)) + "/reopen"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReopenMergeRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置企业公钥
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSshKeyResponse
func (client *Client) ResetSshKeyWithContext(ctx context.Context, organizationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetSshKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetSshKey"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/sshKey"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetSshKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 继续部署
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeVMDeployOrderResponse
func (client *Client) ResumeVMDeployOrderWithContext(ctx context.Context, organizationId *string, pipelineId *string, deployOrderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeVMDeployOrderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeVMDeployOrder"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/deploy/" + dara.PercentEncode(dara.StringValue(deployOrderId)) + "/resume"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeVMDeployOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重试流水线运行
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryPipelineJobRunResponse
func (client *Client) RetryPipelineJobRunWithContext(ctx context.Context, organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryPipelineJobRunResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryPipelineJobRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRuns/" + dara.PercentEncode(dara.StringValue(pipelineRunId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryPipelineJobRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重试机器部署
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryVMDeployMachineResponse
func (client *Client) RetryVMDeployMachineWithContext(ctx context.Context, organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryVMDeployMachineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryVMDeployMachine"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/deploy/" + dara.PercentEncode(dara.StringValue(deployOrderId)) + "/machine/" + dara.PercentEncode(dara.StringValue(machineSn)) + "/retry"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryVMDeployMachineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交变更请求的评审意见
//
// @param request - ReviewMergeRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReviewMergeRequestResponse
func (client *Client) ReviewMergeRequestWithContext(ctx context.Context, repositoryId *string, localId *string, request *ReviewMergeRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReviewMergeRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DraftCommentIds) {
		body["draftCommentIds"] = request.DraftCommentIds
	}

	if !dara.IsNil(request.ReviewComment) {
		body["reviewComment"] = request.ReviewComment
	}

	if !dara.IsNil(request.ReviewOpinion) {
		body["reviewOpinion"] = request.ReviewOpinion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReviewMergeRequest"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/merge_requests/" + dara.PercentEncode(dara.StringValue(localId)) + "/submit_review_opinion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReviewMergeRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 跳过流水线任务运行
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipPipelineJobRunResponse
func (client *Client) SkipPipelineJobRunWithContext(ctx context.Context, organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SkipPipelineJobRunResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SkipPipelineJobRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRuns/" + dara.PercentEncode(dara.StringValue(pipelineRunId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/skip"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SkipPipelineJobRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 跳过机器部署
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipVMDeployMachineResponse
func (client *Client) SkipVMDeployMachineWithContext(ctx context.Context, organizationId *string, pipelineId *string, deployOrderId *string, machineSn *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SkipVMDeployMachineResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SkipVMDeployMachine"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/deploy/" + dara.PercentEncode(dara.StringValue(deployOrderId)) + "/machine/" + dara.PercentEncode(dara.StringValue(machineSn)) + "/skip"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SkipVMDeployMachineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行流水线
//
// @param request - StartPipelineRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartPipelineRunResponse
func (client *Client) StartPipelineRunWithContext(ctx context.Context, organizationId *string, pipelineId *string, request *StartPipelineRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartPipelineRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartPipelineRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organizations/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartPipelineRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 终止流水线任务运行
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPipelineJobRunResponse
func (client *Client) StopPipelineJobRunWithContext(ctx context.Context, organizationId *string, pipelineId *string, pipelineRunId *string, jobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopPipelineJobRunResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopPipelineJobRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRuns/" + dara.PercentEncode(dara.StringValue(pipelineRunId)) + "/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopPipelineJobRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 终止流水线运行
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopPipelineRunResponse
func (client *Client) StopPipelineRunWithContext(ctx context.Context, organizationId *string, pipelineId *string, pipelineRunId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopPipelineRunResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopPipelineRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/pipelineRuns/" + dara.PercentEncode(dara.StringValue(pipelineRunId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopPipelineRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消部署单
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopVMDeployOrderResponse
func (client *Client) StopVMDeployOrderWithContext(ctx context.Context, organizationId *string, pipelineId *string, deployOrderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopVMDeployOrderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopVMDeployOrder"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/deploy/" + dara.PercentEncode(dara.StringValue(deployOrderId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopVMDeployOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转移代码库
//
// @param request - TransferRepositoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferRepositoryResponse
func (client *Client) TransferRepositoryWithContext(ctx context.Context, request *TransferRepositoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TransferRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryId) {
		query["repositoryId"] = request.RepositoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferRepository"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/repository/transfer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发仓库同步
//
// @param request - TriggerRepositoryMirrorSyncRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerRepositoryMirrorSyncResponse
func (client *Client) TriggerRepositoryMirrorSyncWithContext(ctx context.Context, repositoryId *string, request *TriggerRepositoryMirrorSyncRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TriggerRepositoryMirrorSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.Account) {
		query["account"] = request.Account
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.Token) {
		query["token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerRepositoryMirrorSync"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/mirror"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerRepositoryMirrorSyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用成员
//
// @param request - UpdateAppMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppMemberResponse
func (client *Client) UpdateAppMemberWithContext(ctx context.Context, appName *string, request *UpdateAppMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAppMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Player) {
		body["player"] = request.Player
	}

	if !dara.IsNil(request.RoleNames) {
		body["roleNames"] = request.RoleNames
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName)) + "/members"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &UpdateAppMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用
//
// @param request - UpdateApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplicationWithContext(ctx context.Context, appName *string, request *UpdateApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccountId) {
		body["ownerAccountId"] = request.OwnerAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplication"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/appstack/apps/" + dara.PercentEncode(dara.StringValue(appName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新检查运行记录
//
// @param request - UpdateCheckRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCheckRunResponse
func (client *Client) UpdateCheckRunWithContext(ctx context.Context, request *UpdateCheckRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCheckRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.CheckRunId) {
		query["checkRunId"] = request.CheckRunId
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Annotations) {
		body["annotations"] = request.Annotations
	}

	if !dara.IsNil(request.CompletedAt) {
		body["completedAt"] = request.CompletedAt
	}

	if !dara.IsNil(request.Conclusion) {
		body["conclusion"] = request.Conclusion
	}

	if !dara.IsNil(request.DetailsUrl) {
		body["detailsUrl"] = request.DetailsUrl
	}

	if !dara.IsNil(request.ExternalId) {
		body["externalId"] = request.ExternalId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.StartedAt) {
		body["startedAt"] = request.StartedAt
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCheckRun"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/check_runs/update_check_run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCheckRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代码库文件
//
// @param request - UpdateFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileResponse
func (client *Client) UpdateFileWithContext(ctx context.Context, repositoryId *string, request *UpdateFileRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BranchName) {
		body["branchName"] = request.BranchName
	}

	if !dara.IsNil(request.CommitMessage) {
		body["commitMessage"] = request.CommitMessage
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Encoding) {
		body["encoding"] = request.Encoding
	}

	if !dara.IsNil(request.NewPath) {
		body["newPath"] = request.NewPath
	}

	if !dara.IsNil(request.OldPath) {
		body["oldPath"] = request.OldPath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFile"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/files/update"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新标签
//
// @param request - UpdateFlowTagRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowTagResponse
func (client *Client) UpdateFlowTagWithContext(ctx context.Context, organizationId *string, id *string, request *UpdateFlowTagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFlowTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Color) {
		query["color"] = request.Color
	}

	if !dara.IsNil(request.FlowTagGroupId) {
		query["flowTagGroupId"] = request.FlowTagGroupId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlowTag"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/flow/tags/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 标签分类
//
// @param request - UpdateFlowTagGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowTagGroupResponse
func (client *Client) UpdateFlowTagGroupWithContext(ctx context.Context, organizationId *string, id *string, request *UpdateFlowTagGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFlowTagGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlowTagGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/flow/tagGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowTagGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新单个代码组信息
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AvatarUrl) {
		body["avatarUrl"] = request.AvatarUrl
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Path) {
		body["path"] = request.Path
	}

	if !dara.IsNil(request.PathWithNamespace) {
		body["pathWithNamespace"] = request.PathWithNamespace
	}

	if !dara.IsNil(request.VisibilityLevel) {
		body["visibilityLevel"] = request.VisibilityLevel
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/groups/modify"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 修改组成员
//
// @param request - UpdateGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupMemberResponse
func (client *Client) UpdateGroupMemberWithContext(ctx context.Context, groupId *string, request *UpdateGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.AliyunPk) {
		query["aliyunPk"] = request.AliyunPk
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessLevel) {
		body["accessLevel"] = request.AccessLevel
	}

	if !dara.IsNil(request.MemberType) {
		body["memberType"] = request.MemberType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGroupMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/groups/" + dara.PercentEncode(dara.StringValue(groupId)) + "/members/update/aliyun_pk"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新主机组
//
// @param request - UpdateHostGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHostGroupResponse
func (client *Client) UpdateHostGroupWithContext(ctx context.Context, organizationId *string, id *string, request *UpdateHostGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHostGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunRegion) {
		body["aliyunRegion"] = request.AliyunRegion
	}

	if !dara.IsNil(request.EcsLabelKey) {
		body["ecsLabelKey"] = request.EcsLabelKey
	}

	if !dara.IsNil(request.EcsLabelValue) {
		body["ecsLabelValue"] = request.EcsLabelValue
	}

	if !dara.IsNil(request.EcsType) {
		body["ecsType"] = request.EcsType
	}

	if !dara.IsNil(request.EnvId) {
		body["envId"] = request.EnvId
	}

	if !dara.IsNil(request.MachineInfos) {
		body["machineInfos"] = request.MachineInfos
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ServiceConnectionId) {
		body["serviceConnectionId"] = request.ServiceConnectionId
	}

	if !dara.IsNil(request.TagIds) {
		body["tagIds"] = request.TagIds
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHostGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/hostGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHostGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代码评审的标题和描述
//
// @param request - UpdateMergeRequestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMergeRequestResponse
func (client *Client) UpdateMergeRequestWithContext(ctx context.Context, repositoryId *string, localId *string, request *UpdateMergeRequestRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMergeRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMergeRequest"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/merge_requests/" + dara.PercentEncode(dara.StringValue(localId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMergeRequestResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代码评审的干系人
//
// @param request - UpdateMergeRequestPersonnelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMergeRequestPersonnelResponse
func (client *Client) UpdateMergeRequestPersonnelWithContext(ctx context.Context, repositoryId *string, localId *string, personType *string, request *UpdateMergeRequestPersonnelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMergeRequestPersonnelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NewUserIdList) {
		body["newUserIdList"] = request.NewUserIdList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMergeRequestPersonnel"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/merge_requests/" + dara.PercentEncode(dara.StringValue(localId)) + "/person/" + dara.PercentEncode(dara.StringValue(personType))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMergeRequestPersonnelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 当前用户加入的企业列表
//
// @param request - UpdateOrganizationMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationMemberResponse
func (client *Client) UpdateOrganizationMemberWithContext(ctx context.Context, organizationId *string, accountId *string, request *UpdateOrganizationMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationMemberName) {
		query["organizationMemberName"] = request.OrganizationMemberName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOrganizationMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/members/" + dara.PercentEncode(dara.StringValue(accountId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOrganizationMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流水线。
//
// @param request - UpdatePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineResponse
func (client *Client) UpdatePipelineWithContext(ctx context.Context, organizationId *string, request *UpdatePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.PipelineId) {
		body["pipelineId"] = request.PipelineId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipeline"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - UpdatePipelineBaseInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineBaseInfoResponse
func (client *Client) UpdatePipelineBaseInfoWithContext(ctx context.Context, organizationId *string, pipelineId *string, request *UpdatePipelineBaseInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelineBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvId) {
		query["envId"] = request.EnvId
	}

	if !dara.IsNil(request.PipelineName) {
		query["pipelineName"] = request.PipelineName
	}

	if !dara.IsNil(request.TagList) {
		query["tagList"] = request.TagList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipelineBaseInfo"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelines/" + dara.PercentEncode(dara.StringValue(pipelineId)) + "/baseInfo"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineBaseInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流水线分组
//
// @param request - UpdatePipelineGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineGroupResponse
func (client *Client) UpdatePipelineGroupWithContext(ctx context.Context, organizationId *string, groupId *string, request *UpdatePipelineGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelineGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipelineGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/pipelineGroups/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新项目属性与字段
//
// @param request - UpdateProjectFieldRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectFieldResponse
func (client *Client) UpdateProjectFieldWithContext(ctx context.Context, organizationId *string, identifier *string, request *UpdateProjectFieldRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.StatusIdentifier) {
		body["statusIdentifier"] = request.StatusIdentifier
	}

	if !dara.IsNil(request.UpdateBasicFieldRequestList) {
		body["updateBasicFieldRequestList"] = request.UpdateBasicFieldRequestList
	}

	if !dara.IsNil(request.UpdateForOpenApiList) {
		body["updateForOpenApiList"] = request.UpdateForOpenApiList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectField"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/project/" + dara.PercentEncode(dara.StringValue(identifier))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代码库Label
//
// @param request - UpdateProjectLabelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectLabelResponse
func (client *Client) UpdateProjectLabelWithContext(ctx context.Context, labelId *string, request *UpdateProjectLabelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RepositoryIdentity) {
		query["repositoryIdentity"] = request.RepositoryIdentity
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Color) {
		body["color"] = request.Color
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectLabel"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/labels/" + dara.PercentEncode(dara.StringValue(labelId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectLabelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加项目成员
//
// @param request - UpdateProjectMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectMemberResponse
func (client *Client) UpdateProjectMemberWithContext(ctx context.Context, organizationId *string, projectId *string, request *UpdateProjectMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleIdentifier) {
		body["roleIdentifier"] = request.RoleIdentifier
	}

	if !dara.IsNil(request.TargetIdentifier) {
		body["targetIdentifier"] = request.TargetIdentifier
	}

	if !dara.IsNil(request.TargetType) {
		body["targetType"] = request.TargetType
	}

	if !dara.IsNil(request.UserIdentifier) {
		body["userIdentifier"] = request.UserIdentifier
	}

	if !dara.IsNil(request.UserType) {
		body["userType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/updateMember"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改保护分支设置
//
// @param request - UpdateProtectedBranchesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProtectedBranchesResponse
func (client *Client) UpdateProtectedBranchesWithContext(ctx context.Context, repositoryId *string, id *string, request *UpdateProtectedBranchesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProtectedBranchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowMergeRoles) {
		body["allowMergeRoles"] = request.AllowMergeRoles
	}

	if !dara.IsNil(request.AllowMergeUserIds) {
		body["allowMergeUserIds"] = request.AllowMergeUserIds
	}

	if !dara.IsNil(request.AllowPushRoles) {
		body["allowPushRoles"] = request.AllowPushRoles
	}

	if !dara.IsNil(request.AllowPushUserIds) {
		body["allowPushUserIds"] = request.AllowPushUserIds
	}

	if !dara.IsNil(request.Branch) {
		body["branch"] = request.Branch
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.MergeRequestSetting) {
		body["mergeRequestSetting"] = request.MergeRequestSetting
	}

	if !dara.IsNil(request.TestSettingDTO) {
		body["testSettingDTO"] = request.TestSettingDTO
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProtectedBranches"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/protect_branches/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProtectedBranchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送评审模式开关
//
// @param request - UpdatePushReviewOnOffRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePushReviewOnOffResponse
func (client *Client) UpdatePushReviewOnOffWithContext(ctx context.Context, repositoryId *string, request *UpdatePushReviewOnOffRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePushReviewOnOffResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.TrunkMode) {
		query["trunkMode"] = request.TrunkMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePushReviewOnOff"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/settings/trunk_mode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePushReviewOnOffResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新推送规则
//
// @param request - UpdatePushRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePushRuleResponse
func (client *Client) UpdatePushRuleWithContext(ctx context.Context, repositoryId *string, pushRuleId *string, request *UpdatePushRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePushRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleInfos) {
		body["ruleInfos"] = request.RuleInfos
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePushRule"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v4/projects/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/push_rule/" + dara.PercentEncode(dara.StringValue(pushRuleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePushRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新代码库名称、路径、描述、默认分支等设置
//
// @param request - UpdateRepositoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepositoryResponse
func (client *Client) UpdateRepositoryWithContext(ctx context.Context, repositoryId *string, request *UpdateRepositoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRepositoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AdminSettingLanguage) {
		body["adminSettingLanguage"] = request.AdminSettingLanguage
	}

	if !dara.IsNil(request.Avatar) {
		body["avatar"] = request.Avatar
	}

	if !dara.IsNil(request.BuildsEnabled) {
		body["buildsEnabled"] = request.BuildsEnabled
	}

	if !dara.IsNil(request.CheckEmail) {
		body["checkEmail"] = request.CheckEmail
	}

	if !dara.IsNil(request.DefaultBranch) {
		body["defaultBranch"] = request.DefaultBranch
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.IssuesEnabled) {
		body["issuesEnabled"] = request.IssuesEnabled
	}

	if !dara.IsNil(request.MergeRequestsEnabled) {
		body["mergeRequestsEnabled"] = request.MergeRequestsEnabled
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OpenCloneDownloadControl) {
		body["openCloneDownloadControl"] = request.OpenCloneDownloadControl
	}

	if !dara.IsNil(request.Path) {
		body["path"] = request.Path
	}

	if !dara.IsNil(request.ProjectCloneDownloadMethodList) {
		body["projectCloneDownloadMethodList"] = request.ProjectCloneDownloadMethodList
	}

	if !dara.IsNil(request.ProjectCloneDownloadRoleList) {
		body["projectCloneDownloadRoleList"] = request.ProjectCloneDownloadRoleList
	}

	if !dara.IsNil(request.SnippetsEnabled) {
		body["snippetsEnabled"] = request.SnippetsEnabled
	}

	if !dara.IsNil(request.VisibilityLevel) {
		body["visibilityLevel"] = request.VisibilityLevel
	}

	if !dara.IsNil(request.WikiEnabled) {
		body["wikiEnabled"] = request.WikiEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRepository"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改代码库成员的权限（角色）
//
// @param request - UpdateRepositoryMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepositoryMemberResponse
func (client *Client) UpdateRepositoryMemberWithContext(ctx context.Context, repositoryId *string, aliyunPk *string, request *UpdateRepositoryMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRepositoryMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["accessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.OrganizationId) {
		query["organizationId"] = request.OrganizationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessLevel) {
		body["accessLevel"] = request.AccessLevel
	}

	if !dara.IsNil(request.ExpireAt) {
		body["expireAt"] = request.ExpireAt
	}

	if !dara.IsNil(request.MemberType) {
		body["memberType"] = request.MemberType
	}

	if !dara.IsNil(request.RelatedId) {
		body["relatedId"] = request.RelatedId
	}

	if !dara.IsNil(request.RelatedInfos) {
		body["relatedInfos"] = request.RelatedInfos
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRepositoryMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/repository/" + dara.PercentEncode(dara.StringValue(repositoryId)) + "/members/" + dara.PercentEncode(dara.StringValue(aliyunPk))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRepositoryMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源成员
//
// @param request - UpdateResourceMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceMemberResponse
func (client *Client) UpdateResourceMemberWithContext(ctx context.Context, organizationId *string, resourceType *string, resourceId *string, accountId *string, request *UpdateResourceMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RoleName) {
		body["roleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceMember"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/" + dara.PercentEncode(dara.StringValue(resourceType)) + "/" + dara.PercentEncode(dara.StringValue(resourceId)) + "/members/" + dara.PercentEncode(dara.StringValue(accountId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新测试用例字段
//
// @param request - UpdateTestCaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTestCaseResponse
func (client *Client) UpdateTestCaseWithContext(ctx context.Context, organizationId *string, testcaseIdentifier *string, request *UpdateTestCaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTestCaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateWorkitemPropertyRequest) {
		body["updateWorkitemPropertyRequest"] = request.UpdateWorkitemPropertyRequest
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTestCase"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/testhub/testcase/" + dara.PercentEncode(dara.StringValue(testcaseIdentifier))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTestCaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新了测试计划中测试用例的执行人与状态
//
// @param request - UpdateTestResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTestResultResponse
func (client *Client) UpdateTestResultWithContext(ctx context.Context, organizationId *string, testPlanIdentifier *string, testcaseIdentifier *string, request *UpdateTestResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTestResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Executor) {
		body["executor"] = request.Executor
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTestResult"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/testhub/testplan/" + dara.PercentEncode(dara.StringValue(testPlanIdentifier)) + "/testresult/" + dara.PercentEncode(dara.StringValue(testcaseIdentifier))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTestResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新变量组
//
// @param request - UpdateVariableGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVariableGroupResponse
func (client *Client) UpdateVariableGroupWithContext(ctx context.Context, organizationId *string, id *string, request *UpdateVariableGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVariableGroupResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVariableGroup"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/variableGroups/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVariableGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工作项信息
//
// @param request - UpdateWorkItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkItemResponse
func (client *Client) UpdateWorkItemWithContext(ctx context.Context, organizationId *string, request *UpdateWorkItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FieldType) {
		body["fieldType"] = request.FieldType
	}

	if !dara.IsNil(request.Identifier) {
		body["identifier"] = request.Identifier
	}

	if !dara.IsNil(request.PropertyKey) {
		body["propertyKey"] = request.PropertyKey
	}

	if !dara.IsNil(request.PropertyValue) {
		body["propertyValue"] = request.PropertyValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkItem"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工作项评论
//
// @param request - UpdateWorkitemCommentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkitemCommentResponse
func (client *Client) UpdateWorkitemCommentWithContext(ctx context.Context, organizationId *string, request *UpdateWorkitemCommentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkitemCommentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CommentId) {
		body["commentId"] = request.CommentId
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.FormatType) {
		body["formatType"] = request.FormatType
	}

	if !dara.IsNil(request.WorkitemIdentifier) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkitemComment"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/commentUpdate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkitemCommentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新工作项字段信息
//
// @param request - UpdateWorkitemFieldRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkitemFieldResponse
func (client *Client) UpdateWorkitemFieldWithContext(ctx context.Context, organizationId *string, request *UpdateWorkitemFieldRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkitemFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateWorkitemPropertyRequest) {
		body["updateWorkitemPropertyRequest"] = request.UpdateWorkitemPropertyRequest
	}

	if !dara.IsNil(request.WorkitemIdentifier) {
		body["workitemIdentifier"] = request.WorkitemIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkitemField"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitems/updateWorkitemField"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkitemFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工作项附件创建
//
// @param request - WorkitemAttachmentCreateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WorkitemAttachmentCreateResponse
func (client *Client) WorkitemAttachmentCreateWithContext(ctx context.Context, organizationId *string, workitemIdentifier *string, request *WorkitemAttachmentCreateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WorkitemAttachmentCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["fileKey"] = request.FileKey
	}

	if !dara.IsNil(request.OriginalFilename) {
		body["originalFilename"] = request.OriginalFilename
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WorkitemAttachmentCreate"),
		Version:     dara.String("2021-06-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/organization/" + dara.PercentEncode(dara.StringValue(organizationId)) + "/workitem/" + dara.PercentEncode(dara.StringValue(workitemIdentifier)) + "/attachment"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &WorkitemAttachmentCreateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
