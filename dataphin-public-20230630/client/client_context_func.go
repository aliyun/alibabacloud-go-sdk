// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 添加数据服务项目用户并设置角色。
//
// @param tmpReq - AddDataServiceProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataServiceProjectMemberResponse
func (client *Client) AddDataServiceProjectMemberWithContext(ctx context.Context, tmpReq *AddDataServiceProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *AddDataServiceProjectMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddDataServiceProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDataServiceProjectMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataServiceProjectMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加项目成员。
//
// @param tmpReq - AddProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddProjectMemberResponse
func (client *Client) AddProjectMemberWithContext(ctx context.Context, tmpReq *AddProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *AddProjectMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddProjectMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddProjectMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增租户成员
//
// @param tmpReq - AddTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersResponse
func (client *Client) AddTenantMembersWithContext(ctx context.Context, tmpReq *AddTenantMembersRequest, runtime *dara.RuntimeOptions) (_result *AddTenantMembersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddTenantMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTenantMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTenantMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过原始用户添加租户成员.
//
// @param tmpReq - AddTenantMembersBySourceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersBySourceUserResponse
func (client *Client) AddTenantMembersBySourceUserWithContext(ctx context.Context, tmpReq *AddTenantMembersBySourceUserRequest, runtime *dara.RuntimeOptions) (_result *AddTenantMembersBySourceUserResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddTenantMembersBySourceUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTenantMembersBySourceUser"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTenantMembersBySourceUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加用户组成员.
//
// @param tmpReq - AddUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMemberWithContext(ctx context.Context, tmpReq *AddUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *AddUserGroupMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddUserGroupMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddCommand) {
		request.AddCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddCommand, dara.String("AddCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddCommandShrink) {
		body["AddCommand"] = request.AddCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserGroupMember"),
		Version:     dara.String("2023-06-30"),
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
// 申请API权限。
//
// @param tmpReq - ApplyDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataServiceApiResponse
func (client *Client) ApplyDataServiceApiWithContext(ctx context.Context, tmpReq *ApplyDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataServiceApiResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ApplyDataServiceApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplyCommand) {
		request.ApplyCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplyCommand, dara.String("ApplyCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyCommandShrink) {
		body["ApplyCommand"] = request.ApplyCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyDataServiceApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请应用权限。
//
// @param tmpReq - ApplyDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataServiceAppResponse
func (client *Client) ApplyDataServiceAppWithContext(ctx context.Context, tmpReq *ApplyDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataServiceAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ApplyDataServiceAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplyCommand) {
		request.ApplyCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplyCommand, dara.String("ApplyCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyCommandShrink) {
		body["ApplyCommand"] = request.ApplyCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyDataServiceAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 项目计算源连通性检查。
//
// @param tmpReq - CheckComputeSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComputeSourceConnectivityResponse
func (client *Client) CheckComputeSourceConnectivityWithContext(ctx context.Context, tmpReq *CheckComputeSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *CheckComputeSourceConnectivityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CheckComputeSourceConnectivityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckCommand) {
		request.CheckCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckCommand, dara.String("CheckCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckCommandShrink) {
		body["CheckCommand"] = request.CheckCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckComputeSourceConnectivity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckComputeSourceConnectivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 计算源连通性检查。
//
// @param request - CheckComputeSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComputeSourceConnectivityByIdResponse
func (client *Client) CheckComputeSourceConnectivityByIdWithContext(ctx context.Context, request *CheckComputeSourceConnectivityByIdRequest, runtime *dara.RuntimeOptions) (_result *CheckComputeSourceConnectivityByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckComputeSourceConnectivityById"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckComputeSourceConnectivityByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查数据源连通性
//
// @param tmpReq - CheckDataSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityResponse
func (client *Client) CheckDataSourceConnectivityWithContext(ctx context.Context, tmpReq *CheckDataSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *CheckDataSourceConnectivityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CheckDataSourceConnectivityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckCommand) {
		request.CheckCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckCommand, dara.String("CheckCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckCommandShrink) {
		body["CheckCommand"] = request.CheckCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDataSourceConnectivity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDataSourceConnectivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查已创建的数据源是否正常连通
//
// @param request - CheckDataSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityByIdResponse
func (client *Client) CheckDataSourceConnectivityByIdWithContext(ctx context.Context, request *CheckDataSourceConnectivityByIdRequest, runtime *dara.RuntimeOptions) (_result *CheckDataSourceConnectivityByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDataSourceConnectivityById"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDataSourceConnectivityByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查项目是否存在依赖。
//
// @param request - CheckProjectHasDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckProjectHasDependencyResponse
func (client *Client) CheckProjectHasDependencyWithContext(ctx context.Context, request *CheckProjectHasDependencyRequest, runtime *dara.RuntimeOptions) (_result *CheckProjectHasDependencyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckProjectHasDependency"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckProjectHasDependencyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验用户是否有指定资源权限点.
//
// @param tmpReq - CheckResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourcePermissionResponse
func (client *Client) CheckResourcePermissionWithContext(ctx context.Context, tmpReq *CheckResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *CheckResourcePermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CheckResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckCommand) {
		request.CheckCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckCommand, dara.String("CheckCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckCommandShrink) {
		body["CheckCommand"] = request.CheckCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckResourcePermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckResourcePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建即席查询文件
//
// @param tmpReq - CreateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdHocFileResponse
func (client *Client) CreateAdHocFileWithContext(ctx context.Context, tmpReq *CreateAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAdHocFileResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateAdHocFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAdHocFile"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAdHocFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建离线计算任务。
//
// @param tmpReq - CreateBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchTaskResponse
func (client *Client) CreateBatchTaskWithContext(ctx context.Context, tmpReq *CreateBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateBatchTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateBatchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBatchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建业务实体。
//
// @param tmpReq - CreateBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizEntityResponse
func (client *Client) CreateBizEntityWithContext(ctx context.Context, tmpReq *CreateBizEntityRequest, runtime *dara.RuntimeOptions) (_result *CreateBizEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateBizEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBizEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据板块。
//
// @param tmpReq - CreateBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizUnitResponse
func (client *Client) CreateBizUnitWithContext(ctx context.Context, tmpReq *CreateBizUnitRequest, runtime *dara.RuntimeOptions) (_result *CreateBizUnitResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateBizUnitShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBizUnit"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBizUnitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建计算源。
//
// @param tmpReq - CreateComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeSourceResponse
func (client *Client) CreateComputeSourceWithContext(ctx context.Context, tmpReq *CreateComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeSourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateComputeSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建主题域。
//
// @param tmpReq - CreateDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataDomainResponse
func (client *Client) CreateDataDomainWithContext(ctx context.Context, tmpReq *CreateDataDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDataDomainResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDataDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataDomain"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建新的数据服务API并提交。
//
// @param tmpReq - CreateDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceApiResponse
func (client *Client) CreateDataServiceApiWithContext(ctx context.Context, tmpReq *CreateDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceApiResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDataServiceApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建数据源
//
// @param tmpReq - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithContext(ctx context.Context, tmpReq *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDataSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建菜单树文件目录
//
// @param tmpReq - CreateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectoryWithContext(ctx context.Context, tmpReq *CreateDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDirectoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDirectory"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用补数据接口 1.会生成补数据实例运行：影响相关产产出表数据 2.会进行任务运行：造成计算的费用以及存储的费用
//
// @param tmpReq - CreateNodeSupplementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeSupplementResponse
func (client *Client) CreateNodeSupplementWithContext(ctx context.Context, tmpReq *CreateNodeSupplementRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeSupplementResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateNodeSupplementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodeSupplement"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeSupplementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集成任务。
//
// @param tmpReq - CreatePipelineNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineNodeResponse
func (client *Client) CreatePipelineNodeWithContext(ctx context.Context, tmpReq *CreatePipelineNodeRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineNodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreatePipelineNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreatePipelineNodeCommand) {
		request.CreatePipelineNodeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreatePipelineNodeCommand, dara.String("CreatePipelineNodeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreatePipelineNodeCommandShrink) {
		body["CreatePipelineNodeCommand"] = request.CreatePipelineNodeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipelineNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源文件。
//
// @param tmpReq - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithContext(ctx context.Context, tmpReq *CreateResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建行级权限
//
// @param tmpReq - CreateRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRowPermissionResponse
func (client *Client) CreateRowPermissionWithContext(ctx context.Context, tmpReq *CreateRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *CreateRowPermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateRowPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateRowPermissionCommand) {
		request.CreateRowPermissionCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateRowPermissionCommand, dara.String("CreateRowPermissionCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateRowPermissionCommandShrink) {
		body["CreateRowPermissionCommand"] = request.CreateRowPermissionCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRowPermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRowPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流批一体任务
//
// @param tmpReq - CreateStreamBatchJobMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamBatchJobMappingResponse
func (client *Client) CreateStreamBatchJobMappingWithContext(ctx context.Context, tmpReq *CreateStreamBatchJobMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamBatchJobMappingResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateStreamBatchJobMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamBatchJobMappingCreateCommand) {
		request.StreamBatchJobMappingCreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamBatchJobMappingCreateCommand, dara.String("StreamBatchJobMappingCreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.StreamBatchJobMappingCreateCommandShrink) {
		body["StreamBatchJobMappingCreateCommand"] = request.StreamBatchJobMappingCreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStreamBatchJobMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStreamBatchJobMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义函数。
//
// @param tmpReq - CreateUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfResponse
func (client *Client) CreateUdfWithContext(ctx context.Context, tmpReq *CreateUdfRequest, runtime *dara.RuntimeOptions) (_result *CreateUdfResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateUdfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUdf"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUdfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建用户组.
//
// @param tmpReq - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithContext(ctx context.Context, tmpReq *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateUserGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserGroup"),
		Version:     dara.String("2023-06-30"),
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
// 删除菜单树即席查询文件
//
// @param request - DeleteAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdHocFileResponse
func (client *Client) DeleteAdHocFileWithContext(ctx context.Context, request *DeleteAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdHocFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAdHocFile"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdHocFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除离线计算任务，如果任务还没下线需要先下线再删除。
//
// @param tmpReq - DeleteBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBatchTaskResponse
func (client *Client) DeleteBatchTaskWithContext(ctx context.Context, tmpReq *DeleteBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteBatchTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteBatchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBatchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除业务实体。
//
// @param request - DeleteBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizEntityResponse
func (client *Client) DeleteBizEntityWithContext(ctx context.Context, request *DeleteBizEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizUnitId) {
		query["BizUnitId"] = request.BizUnitId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBizEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据板块。
//
// @param request - DeleteBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizUnitResponse
func (client *Client) DeleteBizUnitWithContext(ctx context.Context, request *DeleteBizUnitRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizUnitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBizUnit"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBizUnitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除计算源。
//
// @param request - DeleteComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeSourceResponse
func (client *Client) DeleteComputeSourceWithContext(ctx context.Context, request *DeleteComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComputeSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComputeSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除主题域。
//
// @param request - DeleteDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataDomainResponse
func (client *Client) DeleteDataDomainWithContext(ctx context.Context, request *DeleteDataDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizUnitId) {
		query["BizUnitId"] = request.BizUnitId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataDomain"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据源
//
// @param tmpReq - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithContext(ctx context.Context, tmpReq *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteDataSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteCommand) {
		request.DeleteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteCommand, dara.String("DeleteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCommandShrink) {
		body["DeleteCommand"] = request.DeleteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除菜单树文件目录
//
// @param request - DeleteDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectoryWithContext(ctx context.Context, request *DeleteDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDirectory"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源文件。
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithContext(ctx context.Context, request *DeleteResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除行级权限
//
// @param tmpReq - DeleteRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRowPermissionResponse
func (client *Client) DeleteRowPermissionWithContext(ctx context.Context, tmpReq *DeleteRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *DeleteRowPermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteRowPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteRowPermissionCommand) {
		request.DeleteRowPermissionCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteRowPermissionCommand, dara.String("DeleteRowPermissionCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteRowPermissionCommandShrink) {
		body["DeleteRowPermissionCommand"] = request.DeleteRowPermissionCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRowPermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRowPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义函数。
//
// @param request - DeleteUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfResponse
func (client *Client) DeleteUdfWithContext(ctx context.Context, request *DeleteUdfRequest, runtime *dara.RuntimeOptions) (_result *DeleteUdfResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUdf"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUdfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户组.
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithContext(ctx context.Context, request *DeleteUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroup"),
		Version:     dara.String("2023-06-30"),
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
// 执行即席查询任务。
//
// @param tmpReq - ExecuteAdHocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAdHocTaskResponse
func (client *Client) ExecuteAdHocTaskWithContext(ctx context.Context, tmpReq *ExecuteAdHocTaskRequest, runtime *dara.RuntimeOptions) (_result *ExecuteAdHocTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteAdHocTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExecuteCommand) {
		request.ExecuteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecuteCommand, dara.String("ExecuteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecuteCommandShrink) {
		body["ExecuteCommand"] = request.ExecuteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAdHocTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAdHocTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行手动调度节点。
//
// @param tmpReq - ExecuteManualNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteManualNodeResponse
func (client *Client) ExecuteManualNodeWithContext(ctx context.Context, tmpReq *ExecuteManualNodeRequest, runtime *dara.RuntimeOptions) (_result *ExecuteManualNodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteManualNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExecuteCommand) {
		request.ExecuteCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecuteCommand, dara.String("ExecuteCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecuteCommandShrink) {
		body["ExecuteCommand"] = request.ExecuteCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteManualNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteManualNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重跑下游(修复链路数据), 支持强制重跑下游。影响范围: 1. 会产生计算成本；2. 会影响数据产出
//
// @param tmpReq - FixDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FixDataResponse
func (client *Client) FixDataWithContext(ctx context.Context, tmpReq *FixDataRequest, runtime *dara.RuntimeOptions) (_result *FixDataResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FixDataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FixDataCommand) {
		request.FixDataCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FixDataCommand, dara.String("FixDataCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FixDataCommandShrink) {
		body["FixDataCommand"] = request.FixDataCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FixData"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FixDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据行级权限ID获取某一行级权限下的所有授权账号
//
// @param tmpReq - GetAccountByRowPermissionIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountByRowPermissionIdResponse
func (client *Client) GetAccountByRowPermissionIdWithContext(ctx context.Context, tmpReq *GetAccountByRowPermissionIdRequest, runtime *dara.RuntimeOptions) (_result *GetAccountByRowPermissionIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetAccountByRowPermissionIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GetAccountByRowPermissionIdQuery) {
		request.GetAccountByRowPermissionIdQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GetAccountByRowPermissionIdQuery, dara.String("GetAccountByRowPermissionIdQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GetAccountByRowPermissionIdQueryShrink) {
		body["GetAccountByRowPermissionIdQuery"] = request.GetAccountByRowPermissionIdQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountByRowPermissionId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountByRowPermissionIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询即席查询文件。
//
// @param request - GetAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocFileResponse
func (client *Client) GetAdHocFileWithContext(ctx context.Context, request *GetAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdHocFile"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdHocFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取即席查询任务运行日志。
//
// @param request - GetAdHocTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocTaskLogResponse
func (client *Client) GetAdHocTaskLogWithContext(ctx context.Context, request *GetAdHocTaskLogRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocTaskLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SubTaskId) {
		query["SubTaskId"] = request.SubTaskId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdHocTaskLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdHocTaskLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取即席查询的任务运行结果。
//
// @param request - GetAdHocTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocTaskResultResponse
func (client *Client) GetAdHocTaskResultWithContext(ctx context.Context, request *GetAdHocTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocTaskResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SubTaskId) {
		query["SubTaskId"] = request.SubTaskId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdHocTaskResult"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdHocTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取告警事件详情
//
// @param request - GetAlertEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertEventResponse
func (client *Client) GetAlertEventWithContext(ctx context.Context, request *GetAlertEventRequest, runtime *dara.RuntimeOptions) (_result *GetAlertEventResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlertEvent"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取离线计算任务详情。
//
// @param request - GetBatchTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskInfoResponse
func (client *Client) GetBatchTaskInfoWithContext(ctx context.Context, request *GetBatchTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.IncludeAllUpStreams) {
		query["IncludeAllUpStreams"] = request.IncludeAllUpStreams
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取离线计算任务指定版本任务详情。
//
// @param request - GetBatchTaskInfoByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskInfoByVersionResponse
func (client *Client) GetBatchTaskInfoByVersionWithContext(ctx context.Context, request *GetBatchTaskInfoByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskInfoByVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskInfoByVersion"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskInfoByVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取离线任务自定义血缘。
//
// @param request - GetBatchTaskUdfLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskUdfLineagesResponse
func (client *Client) GetBatchTaskUdfLineagesWithContext(ctx context.Context, request *GetBatchTaskUdfLineagesRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskUdfLineagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskUdfLineages"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskUdfLineagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取离线计算任务版本列表。
//
// @param request - GetBatchTaskVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskVersionsResponse
func (client *Client) GetBatchTaskVersionsWithContext(ctx context.Context, request *GetBatchTaskVersionsRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskVersionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTaskVersions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTaskVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取业务实体详情。
//
// @param request - GetBizEntityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizEntityInfoResponse
func (client *Client) GetBizEntityInfoWithContext(ctx context.Context, request *GetBizEntityInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBizEntityInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBizEntityInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizEntityInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定版本的业务实体的详情。
//
// @param request - GetBizEntityInfoByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizEntityInfoByVersionResponse
func (client *Client) GetBizEntityInfoByVersionWithContext(ctx context.Context, request *GetBizEntityInfoByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetBizEntityInfoByVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBizEntityInfoByVersion"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizEntityInfoByVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据板块详情。
//
// @param request - GetBizUnitInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizUnitInfoResponse
func (client *Client) GetBizUnitInfoWithContext(ctx context.Context, request *GetBizUnitInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBizUnitInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBizUnitInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizUnitInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据环境获取集群信息
//
// @param request - GetClusterQueueInfoByEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterQueueInfoByEnvResponse
func (client *Client) GetClusterQueueInfoByEnvWithContext(ctx context.Context, request *GetClusterQueueInfoByEnvRequest, runtime *dara.RuntimeOptions) (_result *GetClusterQueueInfoByEnvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StreamBatchMode) {
		query["StreamBatchMode"] = request.StreamBatchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterQueueInfoByEnv"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterQueueInfoByEnvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取计算源详情。
//
// @param request - GetComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeSourceResponse
func (client *Client) GetComputeSourceWithContext(ctx context.Context, request *GetComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *GetComputeSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主题域详情。
//
// @param request - GetDataDomainInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataDomainInfoResponse
func (client *Client) GetDataDomainInfoWithContext(ctx context.Context, request *GetDataDomainInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDataDomainInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataDomainInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataDomainInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运维监控Api调用汇总统计。
//
// @param request - GetDataServiceApiCallSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiCallSummaryResponse
func (client *Client) GetDataServiceApiCallSummaryWithContext(ctx context.Context, request *GetDataServiceApiCallSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiCallSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiCallSummary"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiCallSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运维监控Api访问趋势分析。
//
// @param request - GetDataServiceApiCallTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiCallTrendResponse
func (client *Client) GetDataServiceApiCallTrendWithContext(ctx context.Context, request *GetDataServiceApiCallTrendRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiCallTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiCallTrend"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiCallTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取API文档。
//
// @param request - GetDataServiceApiDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiDocumentResponse
func (client *Client) GetDataServiceApiDocumentWithContext(ctx context.Context, request *GetDataServiceApiDocumentRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiDocumentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiDocument"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取API异常影响汇总。
//
// @param request - GetDataServiceApiErrorImpactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiErrorImpactResponse
func (client *Client) GetDataServiceApiErrorImpactWithContext(ctx context.Context, request *GetDataServiceApiErrorImpactRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiErrorImpactResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiErrorImpact"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiErrorImpactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据服务API分组列表。
//
// @param request - GetDataServiceApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiGroupsResponse
func (client *Client) GetDataServiceApiGroupsWithContext(ctx context.Context, request *GetDataServiceApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiGroups"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用有权限的用户列表。
//
// @param request - GetDataServiceAppAuthorizedUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppAuthorizedUsersResponse
func (client *Client) GetDataServiceAppAuthorizedUsersWithContext(ctx context.Context, request *GetDataServiceAppAuthorizedUsersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppAuthorizedUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAppAuthorizedUsers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppAuthorizedUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据服务项目的应用分组列表。
//
// @param request - GetDataServiceAppGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppGroupsResponse
func (client *Client) GetDataServiceAppGroupsWithContext(ctx context.Context, request *GetDataServiceAppGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAppGroups"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询分组下应用列表。
//
// @param request - GetDataServiceAppsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppsByGroupIdResponse
func (client *Client) GetDataServiceAppsByGroupIdWithContext(ctx context.Context, request *GetDataServiceAppsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppsByGroupIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAppsByGroupId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppsByGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据App分组Id查询账号有权限的应用列表。
//
// @param request - GetDataServiceAuthorizedAppsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAuthorizedAppsByGroupIdResponse
func (client *Client) GetDataServiceAuthorizedAppsByGroupIdWithContext(ctx context.Context, request *GetDataServiceAuthorizedAppsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAuthorizedAppsByGroupIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAuthorizedAppsByGroupId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAuthorizedAppsByGroupIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询有权限的项目列表。
//
// @param request - GetDataServiceAuthorizedProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAuthorizedProjectsResponse
func (client *Client) GetDataServiceAuthorizedProjectsWithContext(ctx context.Context, request *GetDataServiceAuthorizedProjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAuthorizedProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAuthorizedProjects"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAuthorizedProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 当前登录当前用户作为负责人的项目列表。
//
// @param request - GetDataServiceMyProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceMyProjectsResponse
func (client *Client) GetDataServiceMyProjectsWithContext(ctx context.Context, request *GetDataServiceMyProjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceMyProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceMyProjects"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceMyProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取可添加到项目成员的用户列表。
//
// @param request - GetDataServiceProjectAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceProjectAddableUsersResponse
func (client *Client) GetDataServiceProjectAddableUsersWithContext(ctx context.Context, request *GetDataServiceProjectAddableUsersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceProjectAddableUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceProjectAddableUsers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceProjectAddableUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据源变更影响的集成任务及数据库SQL任务。
//
// @param request - GetDataSourceDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceDependenciesResponse
func (client *Client) GetDataSourceDependenciesWithContext(ctx context.Context, request *GetDataSourceDependenciesRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceDependenciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataSourceDependencies"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceDependenciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询开发态对象上游依赖。
//
// @param request - GetDevObjectDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDevObjectDependencyResponse
func (client *Client) GetDevObjectDependencyWithContext(ctx context.Context, request *GetDevObjectDependencyRequest, runtime *dara.RuntimeOptions) (_result *GetDevObjectDependencyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectFrom) {
		query["ObjectFrom"] = request.ObjectFrom
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDevObjectDependency"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDevObjectDependencyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件夹目录树
//
// @param request - GetDirectoryTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryTreeResponse
func (client *Client) GetDirectoryTreeWithContext(ctx context.Context, request *GetDirectoryTreeRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryTreeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDirectoryTree"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDirectoryTreeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件存储临时读写授权。
//
// @param request - GetFileStorageCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileStorageCredentialResponse
func (client *Client) GetFileStorageCredentialWithContext(ctx context.Context, request *GetFileStorageCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetFileStorageCredentialResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Purpose) {
		query["Purpose"] = request.Purpose
	}

	if !dara.IsNil(request.UseVpcEndpoint) {
		query["UseVpcEndpoint"] = request.UseVpcEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileStorageCredential"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileStorageCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据起始的实例查询该实例的下游
//
// @param tmpReq - GetInstanceDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceDownStreamResponse
func (client *Client) GetInstanceDownStreamWithContext(ctx context.Context, tmpReq *GetInstanceDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceDownStreamResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetInstanceDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceGet) {
		request.InstanceGetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceGet, dara.String("InstanceGet"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DownStreamDepth) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.RunStatus) {
		query["RunStatus"] = request.RunStatus
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceGetShrink) {
		body["InstanceGet"] = request.InstanceGetShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceDownStream"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceDownStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例的上下游，支持逻辑表和代码任务。
//
// @param tmpReq - GetInstanceUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceUpDownStreamResponse
func (client *Client) GetInstanceUpDownStreamWithContext(ctx context.Context, tmpReq *GetInstanceUpDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceUpDownStreamResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetInstanceUpDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceId) {
		request.InstanceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceId, dara.String("InstanceId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DownStreamDepth) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UpStreamDepth) {
		query["UpStreamDepth"] = request.UpStreamDepth
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdShrink) {
		body["InstanceId"] = request.InstanceIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceUpDownStream"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceUpDownStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取最新的待发布记录详情
//
// @param tmpReq - GetLatestSubmitDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLatestSubmitDetailResponse
func (client *Client) GetLatestSubmitDetailWithContext(ctx context.Context, tmpReq *GetLatestSubmitDetailRequest, runtime *dara.RuntimeOptions) (_result *GetLatestSubmitDetailResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetLatestSubmitDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubmitDetailQuery) {
		request.SubmitDetailQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubmitDetailQuery, dara.String("SubmitDetailQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubmitDetailQueryShrink) {
		body["SubmitDetailQuery"] = request.SubmitDetailQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLatestSubmitDetail"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLatestSubmitDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - GetMyRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyRolesResponse
func (client *Client) GetMyRolesWithContext(ctx context.Context, request *GetMyRolesRequest, runtime *dara.RuntimeOptions) (_result *GetMyRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMyRoles"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMyRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前用户归属租户.
//
// @param tmpReq - GetMyTenantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyTenantsResponse
func (client *Client) GetMyTenantsWithContext(ctx context.Context, tmpReq *GetMyTenantsRequest, runtime *dara.RuntimeOptions) (_result *GetMyTenantsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetMyTenantsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FeatureCodeList) {
		request.FeatureCodeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FeatureCodeList, dara.String("FeatureCodeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FeatureCodeListShrink) {
		body["FeatureCodeList"] = request.FeatureCodeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMyTenants"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMyTenantsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用查询节点上下游接口
//
// @param tmpReq - GetNodeUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeUpDownStreamResponse
func (client *Client) GetNodeUpDownStreamWithContext(ctx context.Context, tmpReq *GetNodeUpDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetNodeUpDownStreamResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetNodeUpDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeId) {
		request.NodeIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeId, dara.String("NodeId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DownStreamDepth) {
		query["DownStreamDepth"] = request.DownStreamDepth
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UpStreamDepth) {
		query["UpStreamDepth"] = request.UpStreamDepth
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeIdShrink) {
		body["NodeId"] = request.NodeIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeUpDownStream"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeUpDownStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询补数据提交的状态
//
// @param request - GetOperationSubmitStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationSubmitStatusResponse
func (client *Client) GetOperationSubmitStatusWithContext(ctx context.Context, request *GetOperationSubmitStatusRequest, runtime *dara.RuntimeOptions) (_result *GetOperationSubmitStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOperationSubmitStatus"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOperationSubmitStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询脚本的实例信息, 包括实例状态、运行时间等信息.
//
// @param request - GetPhysicalInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceResponse
func (client *Client) GetPhysicalInstanceWithContext(ctx context.Context, request *GetPhysicalInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalInstance"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例执行的日志，如果实例重跑了多次，则会有多条日志
//
// @param request - GetPhysicalInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceLogResponse
func (client *Client) GetPhysicalInstanceLogWithContext(ctx context.Context, request *GetPhysicalInstanceLogRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalInstanceLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalInstanceLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalInstanceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询物理调度节点。
//
// @param request - GetPhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeResponse
func (client *Client) GetPhysicalNodeWithContext(ctx context.Context, request *GetPhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据输出名查询对应的物理节点。
//
// @param request - GetPhysicalNodeByOutputNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeByOutputNameResponse
func (client *Client) GetPhysicalNodeByOutputNameWithContext(ctx context.Context, request *GetPhysicalNodeByOutputNameRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeByOutputNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.OutputName) {
		query["OutputName"] = request.OutputName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalNodeByOutputName"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalNodeByOutputNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调度节点代码内容。
//
// @param request - GetPhysicalNodeContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeContentResponse
func (client *Client) GetPhysicalNodeContentWithContext(ctx context.Context, request *GetPhysicalNodeContentRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalNodeContent"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalNodeContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询节点的操作日志。
//
// @param request - GetPhysicalNodeOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeOperationLogResponse
func (client *Client) GetPhysicalNodeOperationLogWithContext(ctx context.Context, request *GetPhysicalNodeOperationLogRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeOperationLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalNodeOperationLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalNodeOperationLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目详情。
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithContext(ctx context.Context, request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过项目名获取项目详情。
//
// @param request - GetProjectByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectByNameResponse
func (client *Client) GetProjectByNameWithContext(ctx context.Context, request *GetProjectByNameRequest, runtime *dara.RuntimeOptions) (_result *GetProjectByNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectByName"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectByNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目生产账号
//
// @param request - GetProjectProduceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectProduceUserResponse
func (client *Client) GetProjectProduceUserWithContext(ctx context.Context, request *GetProjectProduceUserRequest, runtime *dara.RuntimeOptions) (_result *GetProjectProduceUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectProduceUser"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectProduceUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目白名单。
//
// @param request - GetProjectWhiteListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectWhiteListsResponse
func (client *Client) GetProjectWhiteListsWithContext(ctx context.Context, request *GetProjectWhiteListsRequest, runtime *dara.RuntimeOptions) (_result *GetProjectWhiteListsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectWhiteLists"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectWhiteListsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据集群ID获取集群版本
//
// @param request - GetQueueEngineVersionByEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueEngineVersionByEnvResponse
func (client *Client) GetQueueEngineVersionByEnvWithContext(ctx context.Context, request *GetQueueEngineVersionByEnvRequest, runtime *dara.RuntimeOptions) (_result *GetQueueEngineVersionByEnvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.QueueName) {
		query["QueueName"] = request.QueueName
	}

	if !dara.IsNil(request.StreamBatchMode) {
		query["StreamBatchMode"] = request.StreamBatchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueueEngineVersionByEnv"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueueEngineVersionByEnvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源文件详情。
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithContext(ctx context.Context, request *GetResourceRequest, runtime *dara.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源文件指定版本详情。
//
// @param request - GetResourceByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceByVersionResponse
func (client *Client) GetResourceByVersionWithContext(ctx context.Context, request *GetResourceByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetResourceByVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceByVersion"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceByVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取计算源对应集群的spark客户信息
//
// @param request - GetSparkLocalClientInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkLocalClientInfoResponse
func (client *Client) GetSparkLocalClientInfoWithContext(ctx context.Context, request *GetSparkLocalClientInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSparkLocalClientInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvEnum) {
		query["EnvEnum"] = request.EnvEnum
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSparkLocalClientInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSparkLocalClientInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取dataphin实时研发任务集合
//
// @param request - GetStreamJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStreamJobsResponse
func (client *Client) GetStreamJobsWithContext(ctx context.Context, request *GetStreamJobsRequest, runtime *dara.RuntimeOptions) (_result *GetStreamJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStreamJobs"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStreamJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取补数据工作流所有业务日期的Dagrun信息。
//
// @param request - GetSupplementDagrunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunResponse
func (client *Client) GetSupplementDagrunWithContext(ctx context.Context, request *GetSupplementDagrunRequest, runtime *dara.RuntimeOptions) (_result *GetSupplementDagrunResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.SupplementId) {
		query["SupplementId"] = request.SupplementId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSupplementDagrun"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupplementDagrunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出补数据工作流下具体一个业务日期的所有节点的实例。
//
// @param request - GetSupplementDagrunInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunInstanceResponse
func (client *Client) GetSupplementDagrunInstanceWithContext(ctx context.Context, request *GetSupplementDagrunInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetSupplementDagrunInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagrunId) {
		query["DagrunId"] = request.DagrunId
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSupplementDagrunInstance"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSupplementDagrunInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询表字段血缘信息
//
// @param tmpReq - GetTableColumnLineageByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnLineageByTaskIdResponse
func (client *Client) GetTableColumnLineageByTaskIdWithContext(ctx context.Context, tmpReq *GetTableColumnLineageByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnLineageByTaskIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetTableColumnLineageByTaskIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableColumnLineageByTaskIdQuery) {
		request.TableColumnLineageByTaskIdQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableColumnLineageByTaskIdQuery, dara.String("TableColumnLineageByTaskIdQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableColumnLineageByTaskIdQueryShrink) {
		body["TableColumnLineageByTaskIdQuery"] = request.TableColumnLineageByTaskIdQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableColumnLineageByTaskId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableColumnLineageByTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询表血缘信息
//
// @param tmpReq - GetTableLineageByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableLineageByTaskIdResponse
func (client *Client) GetTableLineageByTaskIdWithContext(ctx context.Context, tmpReq *GetTableLineageByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetTableLineageByTaskIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetTableLineageByTaskIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableLineageByTaskIdQuery) {
		request.TableLineageByTaskIdQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableLineageByTaskIdQuery, dara.String("TableLineageByTaskIdQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableLineageByTaskIdQueryShrink) {
		body["TableLineageByTaskIdQuery"] = request.TableLineageByTaskIdQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableLineageByTaskId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableLineageByTaskIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义函数详情。
//
// @param request - GetUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfResponse
func (client *Client) GetUdfWithContext(ctx context.Context, request *GetUdfRequest, runtime *dara.RuntimeOptions) (_result *GetUdfResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUdf"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUdfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义函数版本详情。
//
// @param request - GetUdfByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfByVersionResponse
func (client *Client) GetUdfByVersionWithContext(ctx context.Context, request *GetUdfByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetUdfByVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUdfByVersion"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUdfByVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过用户原始Id（如阿里云Id）获取用户详情
//
// @param request - GetUserBySourceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBySourceIdResponse
func (client *Client) GetUserBySourceIdWithContext(ctx context.Context, request *GetUserBySourceIdRequest, runtime *dara.RuntimeOptions) (_result *GetUserBySourceIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserBySourceId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserBySourceIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户组详情.
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithContext(ctx context.Context, request *GetUserGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户详情
//
// @param tmpReq - GetUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUsersResponse
func (client *Client) GetUsersWithContext(ctx context.Context, tmpReq *GetUsersRequest, runtime *dara.RuntimeOptions) (_result *GetUsersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserIdList) {
		request.UserIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIdList, dara.String("UserIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UserIdListShrink) {
		body["UserIdList"] = request.UserIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUsers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// API授权。
//
// @param tmpReq - GrantDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantDataServiceApiResponse
func (client *Client) GrantDataServiceApiWithContext(ctx context.Context, tmpReq *GrantDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *GrantDataServiceApiResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GrantDataServiceApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GrantCommand) {
		request.GrantCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GrantCommand, dara.String("GrantCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GrantCommandShrink) {
		body["GrantCommand"] = request.GrantCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantDataServiceApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过资源点对用户授权
//
// @param tmpReq - GrantResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantResourcePermissionResponse
func (client *Client) GrantResourcePermissionWithContext(ctx context.Context, tmpReq *GrantResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantResourcePermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GrantResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GrantCommand) {
		request.GrantCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GrantCommand, dara.String("GrantCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GrantCommandShrink) {
		body["GrantCommand"] = request.GrantCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantResourcePermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantResourcePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - ListAddableRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableRolesResponse
func (client *Client) ListAddableRolesWithContext(ctx context.Context, request *ListAddableRolesRequest, runtime *dara.RuntimeOptions) (_result *ListAddableRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddableRoles"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddableRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取可加入租户成员列表的用户
//
// @param tmpReq - ListAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableUsersResponse
func (client *Client) ListAddableUsersWithContext(ctx context.Context, tmpReq *ListAddableUsersRequest, runtime *dara.RuntimeOptions) (_result *ListAddableUsersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListAddableUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddableUsers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAddableUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件查询多个告警事件
//
// @param tmpReq - ListAlertEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertEventsResponse
func (client *Client) ListAlertEventsWithContext(ctx context.Context, tmpReq *ListAlertEventsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertEventsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListAlertEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertEvents"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件查询多个推送记录
//
// @param tmpReq - ListAlertNotificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertNotificationsResponse
func (client *Client) ListAlertNotificationsWithContext(ctx context.Context, tmpReq *ListAlertNotificationsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertNotificationsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListAlertNotificationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertNotifications"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertNotificationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据app查询api列表
//
// @param tmpReq - ListApiByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiByAppResponse
func (client *Client) ListApiByAppWithContext(ctx context.Context, tmpReq *ListApiByAppRequest, runtime *dara.RuntimeOptions) (_result *ListApiByAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListApiByAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PageQuery) {
		request.PageQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageQuery, dara.String("PageQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageQueryShrink) {
		body["PageQuery"] = request.PageQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiByApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiByAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用已申请的API的具体的字段列表
//
// @param tmpReq - ListAuthorizedDataServiceApiDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedDataServiceApiDetailsResponse
func (client *Client) ListAuthorizedDataServiceApiDetailsWithContext(ctx context.Context, tmpReq *ListAuthorizedDataServiceApiDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedDataServiceApiDetailsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListAuthorizedDataServiceApiDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedDataServiceApiDetails"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedDataServiceApiDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询业务实体列表。
//
// @param tmpReq - ListBizEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizEntitiesResponse
func (client *Client) ListBizEntitiesWithContext(ctx context.Context, tmpReq *ListBizEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListBizEntitiesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListBizEntitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBizEntities"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBizEntitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前租户下的所有数据板块
//
// @param request - ListBizUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizUnitsResponse
func (client *Client) ListBizUnitsWithContext(ctx context.Context, request *ListBizUnitsRequest, runtime *dara.RuntimeOptions) (_result *ListBizUnitsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBizUnits"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBizUnitsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询计算源列表。
//
// @param tmpReq - ListComputeSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeSourcesResponse
func (client *Client) ListComputeSourcesWithContext(ctx context.Context, tmpReq *ListComputeSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListComputeSourcesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListComputeSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeSources"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取主题域列表。
//
// @param tmpReq - ListDataDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataDomainsResponse
func (client *Client) ListDataDomainsWithContext(ctx context.Context, tmpReq *ListDataDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDataDomainsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataDomainsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataDomains"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询API运维统计信息。
//
// @param tmpReq - ListDataServiceApiCallStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiCallStatisticsResponse
func (client *Client) ListDataServiceApiCallStatisticsWithContext(ctx context.Context, tmpReq *ListDataServiceApiCallStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiCallStatisticsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataServiceApiCallStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApiCallStatistics"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApiCallStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询数据服务调用日志。
//
// @param tmpReq - ListDataServiceApiCallsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiCallsResponse
func (client *Client) ListDataServiceApiCallsWithContext(ctx context.Context, tmpReq *ListDataServiceApiCallsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiCallsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataServiceApiCallsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApiCalls"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApiCallsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// API影响分析列表。
//
// @param tmpReq - ListDataServiceApiImpactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiImpactsResponse
func (client *Client) ListDataServiceApiImpactsWithContext(ctx context.Context, tmpReq *ListDataServiceApiImpactsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiImpactsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataServiceApiImpactsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApiImpacts"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApiImpactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前用户有权限的应用列表。
//
// @param tmpReq - ListDataServiceAuthorizedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceAuthorizedAppsResponse
func (client *Client) ListDataServiceAuthorizedAppsWithContext(ctx context.Context, tmpReq *ListDataServiceAuthorizedAppsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceAuthorizedAppsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataServiceAuthorizedAppsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceAuthorizedApps"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceAuthorizedAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取我管理的API权限列表。
//
// @param tmpReq - ListDataServiceMyApiPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceMyApiPermissionsResponse
func (client *Client) ListDataServiceMyApiPermissionsWithContext(ctx context.Context, tmpReq *ListDataServiceMyApiPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceMyApiPermissionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataServiceMyApiPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceMyApiPermissions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceMyApiPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前用户有权限的应用。
//
// @param tmpReq - ListDataServiceMyAppPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceMyAppPermissionsResponse
func (client *Client) ListDataServiceMyAppPermissionsWithContext(ctx context.Context, tmpReq *ListDataServiceMyAppPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceMyAppPermissionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataServiceMyAppPermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceMyAppPermissions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceMyAppPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询已发布的API列表。
//
// @param tmpReq - ListDataServicePublishedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServicePublishedApisResponse
func (client *Client) ListDataServicePublishedApisWithContext(ctx context.Context, tmpReq *ListDataServicePublishedApisRequest, runtime *dara.RuntimeOptions) (_result *ListDataServicePublishedApisResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataServicePublishedApisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServicePublishedApis"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServicePublishedApisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索数据源，所属结果包含数据源配置项
//
// @param tmpReq - ListDataSourceWithConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceWithConfigResponse
func (client *Client) ListDataSourceWithConfigWithContext(ctx context.Context, tmpReq *ListDataSourceWithConfigRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceWithConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListDataSourceWithConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceWithConfig"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceWithConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 遍历菜单树目录文件。
//
// @param tmpReq - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithContext(ctx context.Context, tmpReq *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFiles"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询实例。
//
// @param tmpReq - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, tmpReq *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2023-06-30"),
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
// 查询节点下游，创建补数据工作流时可以作为数据参考
//
// @param tmpReq - ListNodeDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDownStreamResponse
func (client *Client) ListNodeDownStreamWithContext(ctx context.Context, tmpReq *ListNodeDownStreamRequest, runtime *dara.RuntimeOptions) (_result *ListNodeDownStreamResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListNodeDownStreamShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeDownStream"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeDownStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询调度节点列表。
//
// @param tmpReq - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithContext(ctx context.Context, tmpReq *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目成员列表。
//
// @param tmpReq - ListProjectMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithContext(ctx context.Context, tmpReq *ListProjectMembersRequest, runtime *dara.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 获取项目列表。
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 分页获取发布记录列表
//
// @param tmpReq - ListPublishRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishRecordsResponse
func (client *Client) ListPublishRecordsWithContext(ctx context.Context, tmpReq *ListPublishRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListPublishRecordsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListPublishRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishRecords"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取权限操作列表
//
// @param tmpReq - ListResourcePermissionOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionOperationLogResponse
func (client *Client) ListResourcePermissionOperationLogWithContext(ctx context.Context, tmpReq *ListResourcePermissionOperationLogRequest, runtime *dara.RuntimeOptions) (_result *ListResourcePermissionOperationLogResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListResourcePermissionOperationLogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourcePermissionOperationLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcePermissionOperationLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取权限记录列表
//
// @param tmpReq - ListResourcePermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionsResponse
func (client *Client) ListResourcePermissionsWithContext(ctx context.Context, tmpReq *ListResourcePermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListResourcePermissionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListResourcePermissionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourcePermissions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcePermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询行级权限
//
// @param tmpReq - ListRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRowPermissionResponse
func (client *Client) ListRowPermissionWithContext(ctx context.Context, tmpReq *ListRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *ListRowPermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListRowPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PageRowPermissionQuery) {
		request.PageRowPermissionQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageRowPermissionQuery, dara.String("PageRowPermissionQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageRowPermissionQueryShrink) {
		body["PageRowPermissionQuery"] = request.PageRowPermissionQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRowPermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRowPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询指定用户行级权限
//
// @param tmpReq - ListRowPermissionByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRowPermissionByUserIdResponse
func (client *Client) ListRowPermissionByUserIdWithContext(ctx context.Context, tmpReq *ListRowPermissionByUserIdRequest, runtime *dara.RuntimeOptions) (_result *ListRowPermissionByUserIdResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListRowPermissionByUserIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListRowPermissionByUserIdQuery) {
		request.ListRowPermissionByUserIdQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListRowPermissionByUserIdQuery, dara.String("ListRowPermissionByUserIdQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListRowPermissionByUserIdQueryShrink) {
		body["ListRowPermissionByUserIdQuery"] = request.ListRowPermissionByUserIdQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRowPermissionByUserId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRowPermissionByUserIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取待发布记录列表
//
// @param tmpReq - ListSubmitRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubmitRecordsResponse
func (client *Client) ListSubmitRecordsWithContext(ctx context.Context, tmpReq *ListSubmitRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListSubmitRecordsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListSubmitRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubmitRecords"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubmitRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询租户成员列表
//
// @param tmpReq - ListTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantMembersResponse
func (client *Client) ListTenantMembersWithContext(ctx context.Context, tmpReq *ListTenantMembersRequest, runtime *dara.RuntimeOptions) (_result *ListTenantMembersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListTenantMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTenantMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTenantMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户组成员列表分页查询.
//
// @param tmpReq - ListUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupMembersResponse
func (client *Client) ListUserGroupMembersWithContext(ctx context.Context, tmpReq *ListUserGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupMembersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListUserGroupMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroupMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户组列表分页查询.
//
// @param tmpReq - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithContext(ctx context.Context, tmpReq *ListUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListUserGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListQuery) {
		request.ListQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListQuery, dara.String("ListQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListQueryShrink) {
		body["ListQuery"] = request.ListQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroups"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下线离线计算任务。
//
// @param request - OfflineBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineBatchTaskResponse
func (client *Client) OfflineBatchTaskWithContext(ctx context.Context, request *OfflineBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *OfflineBatchTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineBatchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下线业务实体、
//
// @param tmpReq - OfflineBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineBizEntityResponse
func (client *Client) OfflineBizEntityWithContext(ctx context.Context, tmpReq *OfflineBizEntityRequest, runtime *dara.RuntimeOptions) (_result *OfflineBizEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &OfflineBizEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OfflineCommand) {
		request.OfflineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OfflineCommand, dara.String("OfflineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OfflineCommandShrink) {
		body["OfflineCommand"] = request.OfflineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineBizEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上线业务实体。
//
// @param tmpReq - OnlineBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineBizEntityResponse
func (client *Client) OnlineBizEntityWithContext(ctx context.Context, tmpReq *OnlineBizEntityRequest, runtime *dara.RuntimeOptions) (_result *OnlineBizEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &OnlineBizEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OnlineCommand) {
		request.OnlineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OnlineCommand, dara.String("OnlineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OnlineCommandShrink) {
		body["OnlineCommand"] = request.OnlineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineBizEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运维实例。
//
// @param tmpReq - OperateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateInstanceResponse
func (client *Client) OperateInstanceWithContext(ctx context.Context, tmpReq *OperateInstanceRequest, runtime *dara.RuntimeOptions) (_result *OperateInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &OperateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperateCommand) {
		request.OperateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperateCommand, dara.String("OperateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OperateCommandShrink) {
		body["OperateCommand"] = request.OperateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateInstance"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解析离线计算任务的逻辑表依赖，注意解析结果上游依赖信息中可能包含自依赖节点（上游节点ID和解析代码的任务节点ID相同）需要用户自己进行处理。
//
// @param tmpReq - ParseBatchTaskDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseBatchTaskDependencyResponse
func (client *Client) ParseBatchTaskDependencyWithContext(ctx context.Context, tmpReq *ParseBatchTaskDependencyRequest, runtime *dara.RuntimeOptions) (_result *ParseBatchTaskDependencyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ParseBatchTaskDependencyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ParseCommand) {
		request.ParseCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParseCommand, dara.String("ParseCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ParseCommandShrink) {
		body["ParseCommand"] = request.ParseCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ParseBatchTaskDependency"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ParseBatchTaskDependencyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停物理节点调度。
//
// @param tmpReq - PausePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PausePhysicalNodeResponse
func (client *Client) PausePhysicalNodeWithContext(ctx context.Context, tmpReq *PausePhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *PausePhysicalNodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PausePhysicalNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PauseCommand) {
		request.PauseCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PauseCommand, dara.String("PauseCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PauseCommandShrink) {
		body["PauseCommand"] = request.PauseCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PausePhysicalNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PausePhysicalNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布数据服务API到生产环境。
//
// @param request - PublishDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDataServiceApiResponse
func (client *Client) PublishDataServiceApiWithContext(ctx context.Context, request *PublishDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *PublishDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishDataServiceApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量发布对象
//
// @param tmpReq - PublishObjectListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishObjectListResponse
func (client *Client) PublishObjectListWithContext(ctx context.Context, tmpReq *PublishObjectListRequest, runtime *dara.RuntimeOptions) (_result *PublishObjectListResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PublishObjectListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PublishCommand) {
		request.PublishCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PublishCommand, dara.String("PublishCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PublishCommandShrink) {
		body["PublishCommand"] = request.PublishCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishObjectList"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishObjectListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目成员。
//
// @param tmpReq - RemoveProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveProjectMemberResponse
func (client *Client) RemoveProjectMemberWithContext(ctx context.Context, tmpReq *RemoveProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveProjectMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RemoveProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveCommand) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, dara.String("RemoveCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RemoveCommandShrink) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveProjectMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveProjectMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除租户成员
//
// @param tmpReq - RemoveTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTenantMemberResponse
func (client *Client) RemoveTenantMemberWithContext(ctx context.Context, tmpReq *RemoveTenantMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveTenantMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RemoveTenantMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveCommand) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, dara.String("RemoveCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RemoveCommandShrink) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTenantMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTenantMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除用户组成员.
//
// @param tmpReq - RemoveUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserGroupMemberResponse
func (client *Client) RemoveUserGroupMemberWithContext(ctx context.Context, tmpReq *RemoveUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserGroupMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RemoveUserGroupMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RemoveCommand) {
		request.RemoveCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveCommand, dara.String("RemoveCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RemoveCommandShrink) {
		body["RemoveCommand"] = request.RemoveCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserGroupMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserGroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新项目白名单。
//
// @param tmpReq - ReplaceProjectWhiteListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceProjectWhiteListsResponse
func (client *Client) ReplaceProjectWhiteListsWithContext(ctx context.Context, tmpReq *ReplaceProjectWhiteListsRequest, runtime *dara.RuntimeOptions) (_result *ReplaceProjectWhiteListsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ReplaceProjectWhiteListsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReplaceCommand) {
		request.ReplaceCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReplaceCommand, dara.String("ReplaceCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ReplaceCommandShrink) {
		body["ReplaceCommand"] = request.ReplaceCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceProjectWhiteLists"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceProjectWhiteListsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复物理节点调度。
//
// @param tmpReq - ResumePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumePhysicalNodeResponse
func (client *Client) ResumePhysicalNodeWithContext(ctx context.Context, tmpReq *ResumePhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *ResumePhysicalNodeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ResumePhysicalNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResumeCommand) {
		request.ResumeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResumeCommand, dara.String("ResumeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResumeCommandShrink) {
		body["ResumeCommand"] = request.ResumeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumePhysicalNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumePhysicalNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回收API授权。
//
// @param tmpReq - RevokeDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeDataServiceApiResponse
func (client *Client) RevokeDataServiceApiWithContext(ctx context.Context, tmpReq *RevokeDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *RevokeDataServiceApiResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RevokeDataServiceApiShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RevokeCommand) {
		request.RevokeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RevokeCommand, dara.String("RevokeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RevokeCommandShrink) {
		body["RevokeCommand"] = request.RevokeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeDataServiceApi"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeDataServiceApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回收用户资源授权
//
// @param tmpReq - RevokeResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourcePermissionResponse
func (client *Client) RevokeResourcePermissionWithContext(ctx context.Context, tmpReq *RevokeResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourcePermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RevokeResourcePermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RevokeCommand) {
		request.RevokeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RevokeCommand, dara.String("RevokeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RevokeCommandShrink) {
		body["RevokeCommand"] = request.RevokeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeResourcePermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeResourcePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 终止即席查询任务。
//
// @param request - StopAdHocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAdHocTaskResponse
func (client *Client) StopAdHocTaskWithContext(ctx context.Context, request *StopAdHocTaskRequest, runtime *dara.RuntimeOptions) (_result *StopAdHocTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAdHocTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAdHocTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交离线计算任务。
//
// @param tmpReq - SubmitBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitBatchTaskResponse
func (client *Client) SubmitBatchTaskWithContext(ctx context.Context, tmpReq *SubmitBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitBatchTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitBatchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubmitCommand) {
		request.SubmitCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubmitCommand, dara.String("SubmitCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubmitCommandShrink) {
		body["SubmitCommand"] = request.SubmitCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitBatchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑即席查询文件。
//
// @param tmpReq - UpdateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdHocFileResponse
func (client *Client) UpdateAdHocFileWithContext(ctx context.Context, tmpReq *UpdateAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdHocFileResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAdHocFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdHocFile"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdHocFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑离线计算任务。
//
// @param tmpReq - UpdateBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchTaskResponse
func (client *Client) UpdateBatchTaskWithContext(ctx context.Context, tmpReq *UpdateBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateBatchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBatchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑离线计算任务自定义血缘。
//
// @param tmpReq - UpdateBatchTaskUdfLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchTaskUdfLineagesResponse
func (client *Client) UpdateBatchTaskUdfLineagesWithContext(ctx context.Context, tmpReq *UpdateBatchTaskUdfLineagesRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchTaskUdfLineagesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateBatchTaskUdfLineagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBatchTaskUdfLineages"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBatchTaskUdfLineagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新业务实体、
//
// @param tmpReq - UpdateBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizEntityResponse
func (client *Client) UpdateBizEntityWithContext(ctx context.Context, tmpReq *UpdateBizEntityRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizEntityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateBizEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBizEntity"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBizEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据板块。
//
// @param tmpReq - UpdateBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizUnitResponse
func (client *Client) UpdateBizUnitWithContext(ctx context.Context, tmpReq *UpdateBizUnitRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizUnitResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateBizUnitShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBizUnit"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBizUnitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改计算源。
//
// @param tmpReq - UpdateComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeSourceResponse
func (client *Client) UpdateComputeSourceWithContext(ctx context.Context, tmpReq *UpdateComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeSourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateComputeSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeSource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新主题域。
//
// @param tmpReq - UpdateDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataDomainResponse
func (client *Client) UpdateDataDomainWithContext(ctx context.Context, tmpReq *UpdateDataDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataDomainResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDataDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataDomain"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑数据源基本信息
//
// @param tmpReq - UpdateDataSourceBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceBasicInfoResponse
func (client *Client) UpdateDataSourceBasicInfoWithContext(ctx context.Context, tmpReq *UpdateDataSourceBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceBasicInfoResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDataSourceBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSourceBasicInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceBasicInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑数据源连接配置项
//
// @param tmpReq - UpdateDataSourceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceConfigResponse
func (client *Client) UpdateDataSourceConfigWithContext(ctx context.Context, tmpReq *UpdateDataSourceConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceConfigResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDataSourceConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSourceConfig"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改菜单树文件所在目录
//
// @param request - UpdateFileDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileDirectoryResponse
func (client *Client) UpdateFileDirectoryWithContext(ctx context.Context, request *UpdateFileDirectoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Directory) {
		query["Directory"] = request.Directory
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFileDirectory"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileDirectoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改菜单树文件名称
//
// @param request - UpdateFileNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileNameResponse
func (client *Client) UpdateFileNameWithContext(ctx context.Context, request *UpdateFileNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFileName"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加项目成员。
//
// @param tmpReq - UpdateProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectMemberResponse
func (client *Client) UpdateProjectMemberWithContext(ctx context.Context, tmpReq *UpdateProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 编辑资源文件。
//
// @param tmpReq - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithContext(ctx context.Context, tmpReq *UpdateResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResource"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新行级权限
//
// @param tmpReq - UpdateRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRowPermissionResponse
func (client *Client) UpdateRowPermissionWithContext(ctx context.Context, tmpReq *UpdateRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *UpdateRowPermissionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateRowPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateRowPermissionCommand) {
		request.UpdateRowPermissionCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateRowPermissionCommand, dara.String("UpdateRowPermissionCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateRowPermissionCommandShrink) {
		body["UpdateRowPermissionCommand"] = request.UpdateRowPermissionCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRowPermission"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRowPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改租户计算设置。
//
// @param tmpReq - UpdateTenantComputeEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantComputeEngineResponse
func (client *Client) UpdateTenantComputeEngineWithContext(ctx context.Context, tmpReq *UpdateTenantComputeEngineRequest, runtime *dara.RuntimeOptions) (_result *UpdateTenantComputeEngineResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTenantComputeEngineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTenantComputeEngine"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTenantComputeEngineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑租户成员
//
// @param tmpReq - UpdateTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantMemberResponse
func (client *Client) UpdateTenantMemberWithContext(ctx context.Context, tmpReq *UpdateTenantMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateTenantMemberResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTenantMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTenantMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTenantMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑自定义函数。
//
// @param tmpReq - UpdateUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfResponse
func (client *Client) UpdateUdfWithContext(ctx context.Context, tmpReq *UpdateUdfRequest, runtime *dara.RuntimeOptions) (_result *UpdateUdfResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateUdfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUdf"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUdfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑用户组.
//
// @param tmpReq - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithContext(ctx context.Context, tmpReq *UpdateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateUserGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserGroup"),
		Version:     dara.String("2023-06-30"),
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
// 编辑用户组启用开关.
//
// @param request - UpdateUserGroupSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupSwitchResponse
func (client *Client) UpdateUserGroupSwitchWithContext(ctx context.Context, request *UpdateUserGroupSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupSwitchResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.UserGroupId) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserGroupSwitch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserGroupSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
