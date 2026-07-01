// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds a regular member to a data service application. Only the application owner can perform this operation.
//
// Online version: v6.0.0.
//
// @param tmpReq - AddDataServiceAppMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataServiceAppMemberResponse
func (client *Client) AddDataServiceAppMemberWithContext(ctx context.Context, tmpReq *AddDataServiceAppMemberRequest, runtime *dara.RuntimeOptions) (_result *AddDataServiceAppMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddDataServiceAppMemberShrinkRequest{}
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
		Action:      dara.String("AddDataServiceAppMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDataServiceAppMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds users to a data service project and assigns roles to them.
//
// @param tmpReq - AddDataServiceProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDataServiceProjectMemberResponse
func (client *Client) AddDataServiceProjectMemberWithContext(ctx context.Context, tmpReq *AddDataServiceProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *AddDataServiceProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Adds members to a project.
//
// @param tmpReq - AddProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddProjectMemberResponse
func (client *Client) AddProjectMemberWithContext(ctx context.Context, tmpReq *AddProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *AddProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Registers data lineage. Available since version v5.4.0.
//
// @param tmpReq - AddRegisterLineageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRegisterLineageResponse
func (client *Client) AddRegisterLineageWithContext(ctx context.Context, tmpReq *AddRegisterLineageRequest, runtime *dara.RuntimeOptions) (_result *AddRegisterLineageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddRegisterLineageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddRegisterLineageCommand) {
		request.AddRegisterLineageCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddRegisterLineageCommand, dara.String("AddRegisterLineageCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddRegisterLineageCommandShrink) {
		body["AddRegisterLineageCommand"] = request.AddRegisterLineageCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRegisterLineage"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRegisterLineageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds users to a tenant in batches. Only super administrators (SuperAdmin) and system administrators can invoke this API operation.
//
// @param tmpReq - AddTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersResponse
func (client *Client) AddTenantMembersWithContext(ctx context.Context, tmpReq *AddTenantMembersRequest, runtime *dara.RuntimeOptions) (_result *AddTenantMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Adds tenant members by using original user identities.
//
// @param tmpReq - AddTenantMembersBySourceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTenantMembersBySourceUserResponse
func (client *Client) AddTenantMembersBySourceUserWithContext(ctx context.Context, tmpReq *AddTenantMembersBySourceUserRequest, runtime *dara.RuntimeOptions) (_result *AddTenantMembersBySourceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Adds members to a user group.
//
// @param tmpReq - AddUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserGroupMemberResponse
func (client *Client) AddUserGroupMemberWithContext(ctx context.Context, tmpReq *AddUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *AddUserGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Applies for API permissions.
//
// @param tmpReq - ApplyDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataServiceApiResponse
func (client *Client) ApplyDataServiceApiWithContext(ctx context.Context, tmpReq *ApplyDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Applies for application permissions.
//
// @param tmpReq - ApplyDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataServiceAppResponse
func (client *Client) ApplyDataServiceAppWithContext(ctx context.Context, tmpReq *ApplyDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Binds specified quality rules to schedule settings.
//
// Release version: v5.4.2.
//
// @param tmpReq - AssignQualityRuleOfAllRuleScopeSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignQualityRuleOfAllRuleScopeSchedulesResponse
func (client *Client) AssignQualityRuleOfAllRuleScopeSchedulesWithContext(ctx context.Context, tmpReq *AssignQualityRuleOfAllRuleScopeSchedulesRequest, runtime *dara.RuntimeOptions) (_result *AssignQualityRuleOfAllRuleScopeSchedulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AssignQualityRuleOfAllRuleScopeSchedulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssignCommand) {
		request.AssignCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssignCommand, dara.String("AssignCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssignCommandShrink) {
		body["AssignCommand"] = request.AssignCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignQualityRuleOfAllRuleScopeSchedules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignQualityRuleOfAllRuleScopeSchedulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the connectivity of a compute source.
//
// @param tmpReq - CheckComputeSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComputeSourceConnectivityResponse
func (client *Client) CheckComputeSourceConnectivityWithContext(ctx context.Context, tmpReq *CheckComputeSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *CheckComputeSourceConnectivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Checks the connectivity of an existing compute source by compute source ID.
//
// @param request - CheckComputeSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckComputeSourceConnectivityByIdResponse
func (client *Client) CheckComputeSourceConnectivityByIdWithContext(ctx context.Context, request *CheckComputeSourceConnectivityByIdRequest, runtime *dara.RuntimeOptions) (_result *CheckComputeSourceConnectivityByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Checks the connectivity of a data source.
//
// @param tmpReq - CheckDataSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityResponse
func (client *Client) CheckDataSourceConnectivityWithContext(ctx context.Context, tmpReq *CheckDataSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *CheckDataSourceConnectivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Checks the connectivity of a data source.
//
// @param request - CheckDataSourceConnectivityByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataSourceConnectivityByIdResponse
func (client *Client) CheckDataSourceConnectivityByIdWithContext(ctx context.Context, request *CheckDataSourceConnectivityByIdRequest, runtime *dara.RuntimeOptions) (_result *CheckDataSourceConnectivityByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Checks whether a project has data dependencies such as tasks.
//
// @param request - CheckProjectHasDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckProjectHasDependencyResponse
func (client *Client) CheckProjectHasDependencyWithContext(ctx context.Context, request *CheckProjectHasDependencyRequest, runtime *dara.RuntimeOptions) (_result *CheckProjectHasDependencyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Checks whether a user has the permission on a specified resource.
//
// @param tmpReq - CheckResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourcePermissionResponse
func (client *Client) CheckResourcePermissionWithContext(ctx context.Context, tmpReq *CheckResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *CheckResourcePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates an ad hoc query file.
//
// @param tmpReq - CreateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdHocFileResponse
func (client *Client) CreateAdHocFileWithContext(ctx context.Context, tmpReq *CreateAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *CreateAdHocFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a batch task.
//
// @param tmpReq - CreateBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchTaskResponse
func (client *Client) CreateBatchTaskWithContext(ctx context.Context, tmpReq *CreateBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a business entity.
//
// @param tmpReq - CreateBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizEntityResponse
func (client *Client) CreateBizEntityWithContext(ctx context.Context, tmpReq *CreateBizEntityRequest, runtime *dara.RuntimeOptions) (_result *CreateBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a business metric.
//
// Release version: v5.5.0.
//
// @param tmpReq - CreateBizMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizMetricResponse
func (client *Client) CreateBizMetricWithContext(ctx context.Context, tmpReq *CreateBizMetricRequest, runtime *dara.RuntimeOptions) (_result *CreateBizMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBizMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateBizMetricCommand) {
		request.CreateBizMetricCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateBizMetricCommand, dara.String("CreateBizMetricCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateBizMetricCommandShrink) {
		body["CreateBizMetricCommand"] = request.CreateBizMetricCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBizMetric"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBizMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data domain.
//
// @param tmpReq - CreateBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBizUnitResponse
func (client *Client) CreateBizUnitWithContext(ctx context.Context, tmpReq *CreateBizUnitRequest, runtime *dara.RuntimeOptions) (_result *CreateBizUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a compute source. Business unit administrators and project administrators have permissions to perform this operation.
//
// @param tmpReq - CreateComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeSourceResponse
func (client *Client) CreateComputeSourceWithContext(ctx context.Context, tmpReq *CreateComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a data domain.
//
// @param tmpReq - CreateDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataDomainResponse
func (client *Client) CreateDataDomainWithContext(ctx context.Context, tmpReq *CreateDataDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDataDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a data service API and submits it.
//
// @param tmpReq - CreateDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceApiResponse
func (client *Client) CreateDataServiceApiWithContext(ctx context.Context, tmpReq *CreateDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a data service application. Only super administrators or system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param tmpReq - CreateDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceAppResponse
func (client *Client) CreateDataServiceAppWithContext(ctx context.Context, tmpReq *CreateDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataServiceAppShrinkRequest{}
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
		Action:      dara.String("CreateDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - CreateDataServiceAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceAppGroupResponse
func (client *Client) CreateDataServiceAppGroupWithContext(ctx context.Context, request *CreateDataServiceAppGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceAppGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceAppGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create Data Source: Tenant administrators, data administrators, business unit administrators, project administrators, and operations administrators have permission to perform this operation.
//
// @param tmpReq - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithContext(ctx context.Context, tmpReq *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a new dataset under a specified project. Available since v6.2.0.
//
// Description:
//
// ## Operation description
//
// - This API creates a new dataset in a specified project.
//
// - `ProjectId` is a required parameter that specifies the ID of the project in which to create the dataset.
//
// - `CreateCommand` is a complex object that contains the configuration information required to create the dataset.
//
// - `Name`, `Type`, `ContentType`, and `Scenario` are required fields that specify the dataset name, type, content type, and scenarios respectively.
//
// - `FileStorageConfig` and `MetadataStorageConfig` in `VersionConfig` can be configured as needed.
//
// - If you need a real-time meta table configuration, provide the `RealtimeMetaTableConfig` information.
//
// - Ensure that all required fields are correctly specified. Otherwise, the request failed.
//
// @param tmpReq - CreateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithContext(ctx context.Context, tmpReq *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
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
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2023-06-30"),
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
// Creates a menu tree directory. This operation supports features such as compute nodes, data integration, and synchronization tasks.
//
// @param tmpReq - CreateDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirectoryResponse
func (client *Client) CreateDirectoryWithContext(ctx context.Context, tmpReq *CreateDirectoryRequest, runtime *dara.RuntimeOptions) (_result *CreateDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// General-purpose backfill API that supports both list-mode and bulk-mode backfill:
//
// 1. Backfill instances will be generated and executed, affecting the data output of related tables.
//
// 2. Task execution will incur computing costs and storage costs.
//
// @param tmpReq - CreateNodeSupplementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeSupplementResponse
func (client *Client) CreateNodeSupplementWithContext(ctx context.Context, tmpReq *CreateNodeSupplementRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeSupplementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Create an integration pipeline/unstructured workflow task.
//
// @param tmpReq - CreatePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineResponse
func (client *Client) CreatePipelineWithContext(ctx context.Context, tmpReq *CreatePipelineRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePipelineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipeline"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
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
// Asynchronously create a pipeline/unstructured workflow.
//
// @param tmpReq - CreatePipelineByAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineByAsyncResponse
func (client *Client) CreatePipelineByAsyncWithContext(ctx context.Context, tmpReq *CreatePipelineByAsyncRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineByAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePipelineByAsyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CreateCommand) {
		request.CreateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateCommand, dara.String("CreateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.CreateCommandShrink) {
		body["CreateCommand"] = request.CreateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipelineByAsync"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineByAsyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data integration task. Note: This operation is deprecated starting from Dataphin v5.3.1. Use CreatePipeline instead.
//
// @param tmpReq - CreatePipelineNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineNodeResponse
func (client *Client) CreatePipelineNodeWithContext(ctx context.Context, tmpReq *CreatePipelineNodeRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a resource file.
//
// @param tmpReq - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithContext(ctx context.Context, tmpReq *CreateResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a row-level permission.
//
// Description:
//
// You can query detailed information about published APIs based on the appKey.
//
// @param tmpReq - CreateRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRowPermissionResponse
func (client *Client) CreateRowPermissionWithContext(ctx context.Context, tmpReq *CreateRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *CreateRowPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a data classification. Available since v5.4.2.
//
// @param tmpReq - CreateSecurityClassifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityClassifyResponse
func (client *Client) CreateSecurityClassifyWithContext(ctx context.Context, tmpReq *CreateSecurityClassifyRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityClassifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityClassifyShrinkRequest{}
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
		Action:      dara.String("CreateSecurityClassify"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityClassifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data classification folder. Available since v5.4.2.
//
// @param tmpReq - CreateSecurityClassifyCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityClassifyCatalogResponse
func (client *Client) CreateSecurityClassifyCatalogWithContext(ctx context.Context, tmpReq *CreateSecurityClassifyCatalogRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityClassifyCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityClassifyCatalogShrinkRequest{}
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
		Action:      dara.String("CreateSecurityClassifyCatalog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityClassifyCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a security identification result.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateSecurityIdentifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityIdentifyResultResponse
func (client *Client) CreateSecurityIdentifyResultWithContext(ctx context.Context, tmpReq *CreateSecurityIdentifyResultRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityIdentifyResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityIdentifyResultShrinkRequest{}
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
		Action:      dara.String("CreateSecurityIdentifyResult"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityIdentifyResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data classification level. Available since v5.4.2.
//
// @param tmpReq - CreateSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityLevelResponse
func (client *Client) CreateSecurityLevelWithContext(ctx context.Context, tmpReq *CreateSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityLevelShrinkRequest{}
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
		Action:      dara.String("CreateSecurityLevel"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityLevelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a standard.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardResponse
func (client *Client) CreateStandardWithContext(ctx context.Context, tmpReq *CreateStandardRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardShrinkRequest{}
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
		Action:      dara.String("CreateStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data standard lookup table.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardLookupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardLookupTableResponse
func (client *Client) CreateStandardLookupTableWithContext(ctx context.Context, tmpReq *CreateStandardLookupTableRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardLookupTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardLookupTableShrinkRequest{}
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
		Action:      dara.String("CreateStandardLookupTable"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardLookupTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates standard mapping relationships, including valid mappings and invalid mappings.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardMappingResponse
func (client *Client) CreateStandardMappingWithContext(ctx context.Context, tmpReq *CreateStandardMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardMappingShrinkRequest{}
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
		Action:      dara.String("CreateStandardMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a standard association. Release version: v5.4.2.
//
// @param tmpReq - CreateStandardRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardRelationsResponse
func (client *Client) CreateStandardRelationsWithContext(ctx context.Context, tmpReq *CreateStandardRelationsRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardRelationsShrinkRequest{}
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
		Action:      dara.String("CreateStandardRelations"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a standard set.
//
// Available since: v5.4.2.
//
// @param tmpReq - CreateStandardSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardSetResponse
func (client *Client) CreateStandardSetWithContext(ctx context.Context, tmpReq *CreateStandardSetRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardSetShrinkRequest{}
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
		Action:      dara.String("CreateStandardSet"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a data standard template.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardTemplateResponse
func (client *Client) CreateStandardTemplateWithContext(ctx context.Context, tmpReq *CreateStandardTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardTemplateShrinkRequest{}
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
		Action:      dara.String("CreateStandardTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data standard root word.
//
// Release version: v5.4.2.
//
// @param tmpReq - CreateStandardWordRootRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardWordRootResponse
func (client *Client) CreateStandardWordRootWithContext(ctx context.Context, tmpReq *CreateStandardWordRootRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardWordRootResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStandardWordRootShrinkRequest{}
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
		Action:      dara.String("CreateStandardWordRoot"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardWordRootResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a stream-batch integrated node.
//
// @param tmpReq - CreateStreamBatchJobMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStreamBatchJobMappingResponse
func (client *Client) CreateStreamBatchJobMappingWithContext(ctx context.Context, tmpReq *CreateStreamBatchJobMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateStreamBatchJobMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a user-defined function.
//
// @param tmpReq - CreateUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfResponse
func (client *Client) CreateUdfWithContext(ctx context.Context, tmpReq *CreateUdfRequest, runtime *dara.RuntimeOptions) (_result *CreateUdfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a user group.
//
// @param tmpReq - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithContext(ctx context.Context, tmpReq *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes an ad hoc query file from the menu tree.
//
// @param request - DeleteAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdHocFileResponse
func (client *Client) DeleteAdHocFileWithContext(ctx context.Context, request *DeleteAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdHocFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a batch task. If the node has not been offlined, you must offline it before deleting it.
//
// @param tmpReq - DeleteBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBatchTaskResponse
func (client *Client) DeleteBatchTaskWithContext(ctx context.Context, tmpReq *DeleteBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a business entity.
//
// @param request - DeleteBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizEntityResponse
func (client *Client) DeleteBizEntityWithContext(ctx context.Context, request *DeleteBizEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a business metric.
//
// Release version: v5.5.0.
//
// @param tmpReq - DeleteBizMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizMetricResponse
func (client *Client) DeleteBizMetricWithContext(ctx context.Context, tmpReq *DeleteBizMetricRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteBizMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteBizMetricCommand) {
		request.DeleteBizMetricCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteBizMetricCommand, dara.String("DeleteBizMetricCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteBizMetricCommandShrink) {
		body["DeleteBizMetricCommand"] = request.DeleteBizMetricCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBizMetric"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBizMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data domain.
//
// @param request - DeleteBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBizUnitResponse
func (client *Client) DeleteBizUnitWithContext(ctx context.Context, request *DeleteBizUnitRequest, runtime *dara.RuntimeOptions) (_result *DeleteBizUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a compute source.
//
// @param request - DeleteComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeSourceResponse
func (client *Client) DeleteComputeSourceWithContext(ctx context.Context, request *DeleteComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a subject domain.
//
// @param request - DeleteDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataDomainResponse
func (client *Client) DeleteDataDomainWithContext(ctx context.Context, request *DeleteDataDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a data service application. Only superusers, system administrators, or application owners can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - DeleteDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataServiceAppResponse
func (client *Client) DeleteDataServiceAppWithContext(ctx context.Context, request *DeleteDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataServiceAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param request - DeleteDataServiceAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataServiceAppGroupResponse
func (client *Client) DeleteDataServiceAppGroupWithContext(ctx context.Context, request *DeleteDataServiceAppGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataServiceAppGroupResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataServiceAppGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataServiceAppGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data source.
//
// @param tmpReq - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithContext(ctx context.Context, tmpReq *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a dataset. Release version: v6.2.0.
//
// @param request - DeleteDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithContext(ctx context.Context, request *DeleteDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file directory from the menu tree.
//
// @param request - DeleteDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirectoryResponse
func (client *Client) DeleteDirectoryWithContext(ctx context.Context, request *DeleteDirectoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes quality rule objects in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteQualityRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityRulesResponse
func (client *Client) DeleteQualityRulesWithContext(ctx context.Context, tmpReq *DeleteQualityRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQualityRulesShrinkRequest{}
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
		Action:      dara.String("DeleteQualityRules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes quality scheduling objects in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteQualitySchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualitySchedulesResponse
func (client *Client) DeleteQualitySchedulesWithContext(ctx context.Context, tmpReq *DeleteQualitySchedulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualitySchedulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQualitySchedulesShrinkRequest{}
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
		Action:      dara.String("DeleteQualitySchedules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualitySchedulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes quality template objects in batches.
//
// Online version: v5.4.2.
//
// @param tmpReq - DeleteQualityTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityTemplatesResponse
func (client *Client) DeleteQualityTemplatesWithContext(ctx context.Context, tmpReq *DeleteQualityTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQualityTemplatesShrinkRequest{}
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
		Action:      dara.String("DeleteQualityTemplates"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes monitored objects in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteQualityWatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityWatchesResponse
func (client *Client) DeleteQualityWatchesWithContext(ctx context.Context, tmpReq *DeleteQualityWatchesRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityWatchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteQualityWatchesShrinkRequest{}
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
		Action:      dara.String("DeleteQualityWatches"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityWatchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes registered lineage. Available since version v5.4.0.
//
// @param tmpReq - DeleteRegisterLineageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegisterLineageResponse
func (client *Client) DeleteRegisterLineageWithContext(ctx context.Context, tmpReq *DeleteRegisterLineageRequest, runtime *dara.RuntimeOptions) (_result *DeleteRegisterLineageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRegisterLineageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeleteRegisterLineageCommand) {
		request.DeleteRegisterLineageCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteRegisterLineageCommand, dara.String("DeleteRegisterLineageCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeleteRegisterLineageCommandShrink) {
		body["DeleteRegisterLineageCommand"] = request.DeleteRegisterLineageCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRegisterLineage"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRegisterLineageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a resource file.
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithContext(ctx context.Context, request *DeleteResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a row-level permission.
//
// @param tmpReq - DeleteRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRowPermissionResponse
func (client *Client) DeleteRowPermissionWithContext(ctx context.Context, tmpReq *DeleteRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *DeleteRowPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a data categorization. Available since v5.4.2.
//
// @param tmpReq - DeleteSecurityClassifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityClassifyResponse
func (client *Client) DeleteSecurityClassifyWithContext(ctx context.Context, tmpReq *DeleteSecurityClassifyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityClassifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSecurityClassifyShrinkRequest{}
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
		Action:      dara.String("DeleteSecurityClassify"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityClassifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data classification catalog. Release version: v5.4.2.
//
// @param tmpReq - DeleteSecurityClassifyCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityClassifyCatalogResponse
func (client *Client) DeleteSecurityClassifyCatalogWithContext(ctx context.Context, tmpReq *DeleteSecurityClassifyCatalogRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityClassifyCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSecurityClassifyCatalogShrinkRequest{}
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
		Action:      dara.String("DeleteSecurityClassifyCatalog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityClassifyCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes security identification results in batches. Release version: v5.4.2.
//
// @param tmpReq - DeleteSecurityIdentifyResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityIdentifyResultsResponse
func (client *Client) DeleteSecurityIdentifyResultsWithContext(ctx context.Context, tmpReq *DeleteSecurityIdentifyResultsRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityIdentifyResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSecurityIdentifyResultsShrinkRequest{}
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
		Action:      dara.String("DeleteSecurityIdentifyResults"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityIdentifyResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data classification level. Available since v5.4.2.
//
// @param tmpReq - DeleteSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityLevelResponse
func (client *Client) DeleteSecurityLevelWithContext(ctx context.Context, tmpReq *DeleteSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSecurityLevelShrinkRequest{}
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
		Action:      dara.String("DeleteSecurityLevel"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityLevelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a standard.
//
// Online version: v5.4.2.
//
// @param tmpReq - DeleteStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardResponse
func (client *Client) DeleteStandardWithContext(ctx context.Context, tmpReq *DeleteStandardRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStandardShrinkRequest{}
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
		Action:      dara.String("DeleteStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes invalid mapping relationships.
//
// Online version: v5.4.2.
//
// @param tmpReq - DeleteStandardInValidMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardInValidMappingResponse
func (client *Client) DeleteStandardInValidMappingWithContext(ctx context.Context, tmpReq *DeleteStandardInValidMappingRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardInValidMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStandardInValidMappingShrinkRequest{}
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
		Action:      dara.String("DeleteStandardInValidMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardInValidMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data standard lookup table. Release version: v5.4.2.
//
// @param request - DeleteStandardLookupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardLookupTableResponse
func (client *Client) DeleteStandardLookupTableWithContext(ctx context.Context, request *DeleteStandardLookupTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardLookupTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DeleteStandardLookupTable"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardLookupTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes standard associations in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteStandardRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardRelationsResponse
func (client *Client) DeleteStandardRelationsWithContext(ctx context.Context, tmpReq *DeleteStandardRelationsRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStandardRelationsShrinkRequest{}
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
		Action:      dara.String("DeleteStandardRelations"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a standard set.
//
// Online version: v5.4.2.
//
// @param request - DeleteStandardSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardSetResponse
func (client *Client) DeleteStandardSetWithContext(ctx context.Context, request *DeleteStandardSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("DeleteStandardSet"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes valid mapping relationships.
//
// Release version: v5.4.2.
//
// @param tmpReq - DeleteStandardValidMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardValidMappingResponse
func (client *Client) DeleteStandardValidMappingWithContext(ctx context.Context, tmpReq *DeleteStandardValidMappingRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardValidMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStandardValidMappingShrinkRequest{}
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
		Action:      dara.String("DeleteStandardValidMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardValidMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data standard root word.
//
// Online version: v5.4.2.
//
// @param request - DeleteStandardWordRootRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardWordRootResponse
func (client *Client) DeleteStandardWordRootWithContext(ctx context.Context, request *DeleteStandardWordRootRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardWordRootResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandardWordRoot"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardWordRootResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function.
//
// @param request - DeleteUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdfResponse
func (client *Client) DeleteUdfWithContext(ctx context.Context, request *DeleteUdfRequest, runtime *dara.RuntimeOptions) (_result *DeleteUdfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a user group.
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
// Executes an ad hoc query task.
//
// @param tmpReq - ExecuteAdHocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAdHocTaskResponse
func (client *Client) ExecuteAdHocTaskWithContext(ctx context.Context, tmpReq *ExecuteAdHocTaskRequest, runtime *dara.RuntimeOptions) (_result *ExecuteAdHocTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Runs a manually scheduled node.
//
// @param tmpReq - ExecuteManualNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteManualNodeResponse
func (client *Client) ExecuteManualNodeWithContext(ctx context.Context, tmpReq *ExecuteManualNodeRequest, runtime *dara.RuntimeOptions) (_result *ExecuteManualNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Runs a trigger-based node.
//
// @param request - ExecuteTriggerNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteTriggerNodeResponse
func (client *Client) ExecuteTriggerNodeWithContext(ctx context.Context, request *ExecuteTriggerNodeRequest, runtime *dara.RuntimeOptions) (_result *ExecuteTriggerNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		query["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	if !dara.IsNil(request.Index) {
		query["Index"] = request.Index
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
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
		Action:      dara.String("ExecuteTriggerNode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteTriggerNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns downstream nodes to fix data link issues. Supports forced rerun of downstream nodes. Impact: incurs compute costs and affects data output.
//
// @param tmpReq - FixDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FixDataResponse
func (client *Client) FixDataWithContext(ctx context.Context, tmpReq *FixDataRequest, runtime *dara.RuntimeOptions) (_result *FixDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves all authorized accounts under a specific row-level permission by row-level permission ID.
//
// @param tmpReq - GetAccountByRowPermissionIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountByRowPermissionIdResponse
func (client *Client) GetAccountByRowPermissionIdWithContext(ctx context.Context, tmpReq *GetAccountByRowPermissionIdRequest, runtime *dara.RuntimeOptions) (_result *GetAccountByRowPermissionIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries a custom query file in the directory tree.
//
// @param request - GetAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocFileResponse
func (client *Client) GetAdHocFileWithContext(ctx context.Context, request *GetAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the runtime logs of an ad hoc query task.
//
// @param request - GetAdHocTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocTaskLogResponse
func (client *Client) GetAdHocTaskLogWithContext(ctx context.Context, request *GetAdHocTaskLogRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocTaskLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the task execution result of an ad hoc query.
//
// @param request - GetAdHocTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdHocTaskResultResponse
func (client *Client) GetAdHocTaskResultWithContext(ctx context.Context, request *GetAdHocTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetAdHocTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of an alert event.
//
// @param request - GetAlertEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertEventResponse
func (client *Client) GetAlertEventWithContext(ctx context.Context, request *GetAlertEventRequest, runtime *dara.RuntimeOptions) (_result *GetAlertEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries mapping relationships by asset object GUID.
//
// Available since: v5.4.2.
//
// @param tmpReq - GetAssetMappingRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAssetMappingRelationsResponse
func (client *Client) GetAssetMappingRelationsWithContext(ctx context.Context, tmpReq *GetAssetMappingRelationsRequest, runtime *dara.RuntimeOptions) (_result *GetAssetMappingRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAssetMappingRelationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetMappingQuery) {
		request.AssetMappingQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetMappingQuery, dara.String("AssetMappingQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssetMappingQueryShrink) {
		body["AssetMappingQuery"] = request.AssetMappingQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAssetMappingRelations"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAssetMappingRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an offline compute node.
//
// @param request - GetBatchTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskInfoResponse
func (client *Client) GetBatchTaskInfoWithContext(ctx context.Context, request *GetBatchTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of a specified version of a batch task.
//
// @param request - GetBatchTaskInfoByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskInfoByVersionResponse
func (client *Client) GetBatchTaskInfoByVersionWithContext(ctx context.Context, request *GetBatchTaskInfoByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskInfoByVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Obtains the custom lineage of an offline task.
//
// @param request - GetBatchTaskUdfLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskUdfLineagesResponse
func (client *Client) GetBatchTaskUdfLineagesWithContext(ctx context.Context, request *GetBatchTaskUdfLineagesRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskUdfLineagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the version list of a batch task.
//
// @param request - GetBatchTaskVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTaskVersionsResponse
func (client *Client) GetBatchTaskVersionsWithContext(ctx context.Context, request *GetBatchTaskVersionsRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTaskVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// 获取指定离线模板ID版本列表。
//
// @param request - GetBatchTemplateVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchTemplateVersionsResponse
func (client *Client) GetBatchTemplateVersionsWithContext(ctx context.Context, request *GetBatchTemplateVersionsRequest, runtime *dara.RuntimeOptions) (_result *GetBatchTemplateVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchTemplateVersions"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchTemplateVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query mapping relationships by belonging asset GUID.
//
// Release version: v5.4.2.
//
// @param tmpReq - GetBelongAssetMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBelongAssetMappingResponse
func (client *Client) GetBelongAssetMappingWithContext(ctx context.Context, tmpReq *GetBelongAssetMappingRequest, runtime *dara.RuntimeOptions) (_result *GetBelongAssetMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetBelongAssetMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetMappingQuery) {
		request.AssetMappingQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetMappingQuery, dara.String("AssetMappingQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssetMappingQueryShrink) {
		body["AssetMappingQuery"] = request.AssetMappingQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBelongAssetMapping"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBelongAssetMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a business entity.
//
// @param request - GetBizEntityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizEntityInfoResponse
func (client *Client) GetBizEntityInfoWithContext(ctx context.Context, request *GetBizEntityInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBizEntityInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the details of a business entity of a specified version.
//
// @param request - GetBizEntityInfoByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizEntityInfoByVersionResponse
func (client *Client) GetBizEntityInfoByVersionWithContext(ctx context.Context, request *GetBizEntityInfoByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetBizEntityInfoByVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Query business metric details by name.
//
// Release version: v5.5.0.
//
// @param tmpReq - GetBizMetricByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizMetricByNameResponse
func (client *Client) GetBizMetricByNameWithContext(ctx context.Context, tmpReq *GetBizMetricByNameRequest, runtime *dara.RuntimeOptions) (_result *GetBizMetricByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetBizMetricByNameShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BizMetricByNameQuery) {
		request.BizMetricByNameQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BizMetricByNameQuery, dara.String("BizMetricByNameQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizMetricByNameQueryShrink) {
		body["BizMetricByNameQuery"] = request.BizMetricByNameQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBizMetricByName"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBizMetricByNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data domain.
//
// @param request - GetBizUnitInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBizUnitInfoResponse
func (client *Client) GetBizUnitInfoWithContext(ctx context.Context, request *GetBizUnitInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBizUnitInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries asset details. Release version: v6.1.0.
//
// @param tmpReq - GetCatalogAssetDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogAssetDetailsResponse
func (client *Client) GetCatalogAssetDetailsWithContext(ctx context.Context, tmpReq *GetCatalogAssetDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetCatalogAssetDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCatalogAssetDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GetCatalogAssetDetailsQuery) {
		request.GetCatalogAssetDetailsQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GetCatalogAssetDetailsQuery, dara.String("GetCatalogAssetDetailsQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GetCatalogAssetDetailsQueryShrink) {
		body["GetCatalogAssetDetailsQuery"] = request.GetCatalogAssetDetailsQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalogAssetDetails"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogAssetDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of connectivity check tasks for a specified data source ID. This operation includes null value validation and tenant permission verification to prevent cross-tenant access.
//
// Release version: v5.5.0.
//
// Description:
//
// Queries the details of connectivity tasks that have been tested for a specified data source ID.
//
// @param request - GetCheckConnectivityJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCheckConnectivityJobsResponse
func (client *Client) GetCheckConnectivityJobsWithContext(ctx context.Context, request *GetCheckConnectivityJobsRequest, runtime *dara.RuntimeOptions) (_result *GetCheckConnectivityJobsResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCheckConnectivityJobs"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCheckConnectivityJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves cluster information based on the environment.
//
// @param request - GetClusterQueueInfoByEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterQueueInfoByEnvResponse
func (client *Client) GetClusterQueueInfoByEnvWithContext(ctx context.Context, request *GetClusterQueueInfoByEnvRequest, runtime *dara.RuntimeOptions) (_result *GetClusterQueueInfoByEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of a compute source by compute source ID.
//
// @param request - GetComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeSourceResponse
func (client *Client) GetComputeSourceWithContext(ctx context.Context, request *GetComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *GetComputeSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of a data domain.
//
// @param request - GetDataDomainInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataDomainInfoResponse
func (client *Client) GetDataDomainInfoWithContext(ctx context.Context, request *GetDataDomainInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDataDomainInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Service Monitoring: Retrieves the aggregate statistics of API calls.
//
// @param request - GetDataServiceApiCallSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiCallSummaryResponse
func (client *Client) GetDataServiceApiCallSummaryWithContext(ctx context.Context, request *GetDataServiceApiCallSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiCallSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Service Monitoring: Analyzes API access trends.
//
// @param request - GetDataServiceApiCallTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiCallTrendResponse
func (client *Client) GetDataServiceApiCallTrendWithContext(ctx context.Context, request *GetDataServiceApiCallTrendRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiCallTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves API documentation.
//
// @param request - GetDataServiceApiDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiDocumentResponse
func (client *Client) GetDataServiceApiDocumentWithContext(ctx context.Context, request *GetDataServiceApiDocumentRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the summary of API exception impacts.
//
// @param request - GetDataServiceApiErrorImpactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiErrorImpactResponse
func (client *Client) GetDataServiceApiErrorImpactWithContext(ctx context.Context, request *GetDataServiceApiErrorImpactRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiErrorImpactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of API groups in Data Service.
//
// @param request - GetDataServiceApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiGroupsResponse
func (client *Client) GetDataServiceApiGroupsWithContext(ctx context.Context, request *GetDataServiceApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the details of a data service application, including the project, application name, authentication information, and IP whitelist. Only application members can view the details.
//
// Release version: v6.0.0.
//
// @param request - GetDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppResponse
func (client *Client) GetDataServiceAppWithContext(ctx context.Context, request *GetDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of users who have permissions on an application.
//
// @param request - GetDataServiceAppAuthorizedUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppAuthorizedUsersResponse
func (client *Client) GetDataServiceAppAuthorizedUsersWithContext(ctx context.Context, request *GetDataServiceAppAuthorizedUsersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppAuthorizedUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of application groups for a data service project.
//
// @param request - GetDataServiceAppGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppGroupsResponse
func (client *Client) GetDataServiceAppGroupsWithContext(ctx context.Context, request *GetDataServiceAppGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the member list of a data service application, including regular members and owners. Only application owners can call this operation.
//
// Online version: v6.0.0.
//
// @param request - GetDataServiceAppMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppMembersResponse
func (client *Client) GetDataServiceAppMembersWithContext(ctx context.Context, request *GetDataServiceAppMembersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceAppMembers"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceAppMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of applications in a group.
//
// @param request - GetDataServiceAppsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAppsByGroupIdResponse
func (client *Client) GetDataServiceAppsByGroupIdWithContext(ctx context.Context, request *GetDataServiceAppsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAppsByGroupIdResponse, _err error) {
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
// Queries the list of applications that the account has permissions to access based on the app group ID.
//
// @param request - GetDataServiceAuthorizedAppsByGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAuthorizedAppsByGroupIdResponse
func (client *Client) GetDataServiceAuthorizedAppsByGroupIdWithContext(ctx context.Context, request *GetDataServiceAuthorizedAppsByGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAuthorizedAppsByGroupIdResponse, _err error) {
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
// Queries the list of projects that the current user has permissions to access.
//
// @param request - GetDataServiceAuthorizedProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceAuthorizedProjectsResponse
func (client *Client) GetDataServiceAuthorizedProjectsWithContext(ctx context.Context, request *GetDataServiceAuthorizedProjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceAuthorizedProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the list of projects for which the current user is the owner.
//
// @param request - GetDataServiceMyProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceMyProjectsResponse
func (client *Client) GetDataServiceMyProjectsWithContext(ctx context.Context, request *GetDataServiceMyProjectsRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceMyProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the list of users who can be added as project members.
//
// @param request - GetDataServiceProjectAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceProjectAddableUsersResponse
func (client *Client) GetDataServiceProjectAddableUsersWithContext(ctx context.Context, request *GetDataServiceProjectAddableUsersRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceProjectAddableUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the integration tasks and database SQL tasks affected by data source changes.
//
// @param request - GetDataSourceDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceDependenciesResponse
func (client *Client) GetDataSourceDependenciesWithContext(ctx context.Context, request *GetDataSourceDependenciesRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceDependenciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves a dataset. Release version: v6.2.0.
//
// Description:
//
// Queries the details of a tested connectivity task based on the data source ID.
//
// @param request - GetDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithContext(ctx context.Context, request *GetDatasetRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetDataset"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query upstream dependencies of development objects.
//
// @param request - GetDevObjectDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDevObjectDependencyResponse
func (client *Client) GetDevObjectDependencyWithContext(ctx context.Context, request *GetDevObjectDependencyRequest, runtime *dara.RuntimeOptions) (_result *GetDevObjectDependencyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the folder directory tree.
//
// @param request - GetDirectoryTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryTreeResponse
func (client *Client) GetDirectoryTreeWithContext(ctx context.Context, request *GetDirectoryTreeRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Obtains temporary read/write authorization for file storage.
//
// @param request - GetFileStorageCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileStorageCredentialResponse
func (client *Client) GetFileStorageCredentialWithContext(ctx context.Context, request *GetFileStorageCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetFileStorageCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the downstream instances of a specified instance.
//
// @param tmpReq - GetInstanceDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceDownStreamResponse
func (client *Client) GetInstanceDownStreamWithContext(ctx context.Context, tmpReq *GetInstanceDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceDownStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the dag of an instance. Logical tables and code nodes are supported.
//
// @param tmpReq - GetInstanceUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceUpDownStreamResponse
func (client *Client) GetInstanceUpDownStreamWithContext(ctx context.Context, tmpReq *GetInstanceUpDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceUpDownStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of the latest pending submit record.
//
// @param tmpReq - GetLatestSubmitDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLatestSubmitDetailResponse
func (client *Client) GetLatestSubmitDetailWithContext(ctx context.Context, tmpReq *GetLatestSubmitDetailRequest, runtime *dara.RuntimeOptions) (_result *GetLatestSubmitDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the list of roles for the current user.
//
// @param request - GetMyRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyRolesResponse
func (client *Client) GetMyRolesWithContext(ctx context.Context, request *GetMyRolesRequest, runtime *dara.RuntimeOptions) (_result *GetMyRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the tenants to which the current user belongs.
//
// @param tmpReq - GetMyTenantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMyTenantsResponse
func (client *Client) GetMyTenantsWithContext(ctx context.Context, tmpReq *GetMyTenantsRequest, runtime *dara.RuntimeOptions) (_result *GetMyTenantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the dag of a node. This is a general-purpose operation.
//
// @param tmpReq - GetNodeUpDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeUpDownStreamResponse
func (client *Client) GetNodeUpDownStreamWithContext(ctx context.Context, tmpReq *GetNodeUpDownStreamRequest, runtime *dara.RuntimeOptions) (_result *GetNodeUpDownStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// 根据Id查询运行记录
//
// @param tmpReq - GetOperationRecordByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationRecordByIdResponse
func (client *Client) GetOperationRecordByIdWithContext(ctx context.Context, tmpReq *GetOperationRecordByIdRequest, runtime *dara.RuntimeOptions) (_result *GetOperationRecordByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetOperationRecordByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DetailCommand) {
		request.DetailCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetailCommand, dara.String("DetailCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DetailCommandShrink) {
		body["DetailCommand"] = request.DetailCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOperationRecordById"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOperationRecordByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an execution record. Released in version v6.2.0.
//
// @param tmpReq - GetOperationRecordDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationRecordDetailResponse
func (client *Client) GetOperationRecordDetailWithContext(ctx context.Context, tmpReq *GetOperationRecordDetailRequest, runtime *dara.RuntimeOptions) (_result *GetOperationRecordDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetOperationRecordDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordDetailCommand) {
		request.RecordDetailCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordDetailCommand, dara.String("RecordDetailCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RecordDetailCommandShrink) {
		body["RecordDetailCommand"] = request.RecordDetailCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOperationRecordDetail"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOperationRecordDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution code of an operation log. Online version: v6.2.0.
//
// @param tmpReq - GetOperationRecordRunCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationRecordRunCodeResponse
func (client *Client) GetOperationRecordRunCodeWithContext(ctx context.Context, tmpReq *GetOperationRecordRunCodeRequest, runtime *dara.RuntimeOptions) (_result *GetOperationRecordRunCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetOperationRecordRunCodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CodeCommand) {
		request.CodeCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CodeCommand, dara.String("CodeCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CodeCommandShrink) {
		body["CodeCommand"] = request.CodeCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOperationRecordRunCode"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOperationRecordRunCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the submit status of a data backfill request.
//
// @param request - GetOperationSubmitStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationSubmitStatusResponse
func (client *Client) GetOperationSubmitStatusWithContext(ctx context.Context, request *GetOperationSubmitStatusRequest, runtime *dara.RuntimeOptions) (_result *GetOperationSubmitStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Gets instance information.
//
// @param request - GetPhysicalInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceResponse
func (client *Client) GetPhysicalInstanceWithContext(ctx context.Context, request *GetPhysicalInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the execution logs of an instance. If the instance has been rerun multiple times, multiple log entries are returned.
//
// @param request - GetPhysicalInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalInstanceLogResponse
func (client *Client) GetPhysicalInstanceLogWithContext(ctx context.Context, request *GetPhysicalInstanceLogRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalInstanceLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries a physical schedule resource.
//
// @param request - GetPhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeResponse
func (client *Client) GetPhysicalNodeWithContext(ctx context.Context, request *GetPhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries a physical node by output name. Only offline code nodes and integration task nodes are supported.
//
// @param request - GetPhysicalNodeByOutputNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeByOutputNameResponse
func (client *Client) GetPhysicalNodeByOutputNameWithContext(ctx context.Context, request *GetPhysicalNodeByOutputNameRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeByOutputNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the code content of a schedule resource node.
//
// @param request - GetPhysicalNodeContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeContentResponse
func (client *Client) GetPhysicalNodeContentWithContext(ctx context.Context, request *GetPhysicalNodeContentRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the operation logs of a node.
//
// @param request - GetPhysicalNodeOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalNodeOperationLogResponse
func (client *Client) GetPhysicalNodeOperationLogWithContext(ctx context.Context, request *GetPhysicalNodeOperationLogRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalNodeOperationLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the execution result of an asynchronous pipeline task.
//
// @param tmpReq - GetPipelineAsyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineAsyncResultResponse
func (client *Client) GetPipelineAsyncResultWithContext(ctx context.Context, tmpReq *GetPipelineAsyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetPipelineAsyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetPipelineAsyncResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AsyncId) {
		query["AsyncId"] = request.AsyncId
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineAsyncResult"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineAsyncResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a pipeline task by pipeline task ID.
//
// @param tmpReq - GetPipelineByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineByIdResponse
func (client *Client) GetPipelineByIdWithContext(ctx context.Context, tmpReq *GetPipelineByIdRequest, runtime *dara.RuntimeOptions) (_result *GetPipelineByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetPipelineByIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryId) {
		request.QueryIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryId, dara.String("QueryId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.QueryIdShrink) {
		body["QueryId"] = request.QueryIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineById"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get project details by project ID.
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithContext(ctx context.Context, request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves project details by project name.
//
// @param request - GetProjectByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectByNameResponse
func (client *Client) GetProjectByNameWithContext(ctx context.Context, request *GetProjectByNameRequest, runtime *dara.RuntimeOptions) (_result *GetProjectByNameResponse, _err error) {
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
// Retrieves the production account of a project. Only a super administrator (SuperAdmin) can call this API operation.
//
// @param request - GetProjectProduceUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectProduceUserResponse
func (client *Client) GetProjectProduceUserWithContext(ctx context.Context, request *GetProjectProduceUserRequest, runtime *dara.RuntimeOptions) (_result *GetProjectProduceUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the whitelist of a project.
//
// @param request - GetProjectWhiteListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectWhiteListsResponse
func (client *Client) GetProjectWhiteListsWithContext(ctx context.Context, request *GetProjectWhiteListsRequest, runtime *dara.RuntimeOptions) (_result *GetProjectWhiteListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves alert settings by monitored object ID. Release version: v5.4.2.
//
// @param request - GetQualityAlertOfAllRuleScopeByWatchIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityAlertOfAllRuleScopeByWatchIdResponse
func (client *Client) GetQualityAlertOfAllRuleScopeByWatchIdWithContext(ctx context.Context, request *GetQualityAlertOfAllRuleScopeByWatchIdRequest, runtime *dara.RuntimeOptions) (_result *GetQualityAlertOfAllRuleScopeByWatchIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchId) {
		query["WatchId"] = request.WatchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityAlertOfAllRuleScopeByWatchId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityAlertOfAllRuleScopeByWatchIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality rule object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleResponse
func (client *Client) GetQualityRuleWithContext(ctx context.Context, request *GetQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetQualityRule"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a quality node task object. Online version: v5.4.2.
//
// @param request - GetQualityRuleTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleTaskResponse
func (client *Client) GetQualityRuleTaskWithContext(ctx context.Context, request *GetQualityRuleTaskRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.RuleTaskId) {
		query["RuleTaskId"] = request.RuleTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRuleTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the log content of a quality node task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityRuleTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleTaskLogResponse
func (client *Client) GetQualityRuleTaskLogWithContext(ctx context.Context, request *GetQualityRuleTaskLogRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleTaskLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.RuleTaskId) {
		query["RuleTaskId"] = request.RuleTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRuleTaskLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleTaskLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality schedule object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityScheduleResponse
func (client *Client) GetQualityScheduleWithContext(ctx context.Context, request *GetQualityScheduleRequest, runtime *dara.RuntimeOptions) (_result *GetQualityScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetQualitySchedule"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of schedule settings by monitored object ID.
//
// Release version: v5.4.2.
//
// @param request - GetQualitySchedulesByWatchIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualitySchedulesByWatchIdResponse
func (client *Client) GetQualitySchedulesByWatchIdWithContext(ctx context.Context, request *GetQualitySchedulesByWatchIdRequest, runtime *dara.RuntimeOptions) (_result *GetQualitySchedulesByWatchIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchId) {
		query["WatchId"] = request.WatchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualitySchedulesByWatchId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualitySchedulesByWatchIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality template object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityTemplateResponse
func (client *Client) GetQualityTemplateWithContext(ctx context.Context, request *GetQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetQualityTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality monitored object.
//
// Release version: v5.4.2.
//
// @param request - GetQualityWatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityWatchResponse
func (client *Client) GetQualityWatchWithContext(ctx context.Context, request *GetQualityWatchRequest, runtime *dara.RuntimeOptions) (_result *GetQualityWatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetQualityWatch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityWatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a quality watchtask record by the original ID of the monitored object, such as the ID of a datasource, table, or metric.
//
// Release version: v5.4.2.
//
// @param request - GetQualityWatchByObjectIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityWatchByObjectIdResponse
func (client *Client) GetQualityWatchByObjectIdWithContext(ctx context.Context, request *GetQualityWatchByObjectIdRequest, runtime *dara.RuntimeOptions) (_result *GetQualityWatchByObjectIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchObjectId) {
		query["WatchObjectId"] = request.WatchObjectId
	}

	if !dara.IsNil(request.WatchType) {
		query["WatchType"] = request.WatchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityWatchByObjectId"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityWatchByObjectIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a monitoring node task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityWatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityWatchTaskResponse
func (client *Client) GetQualityWatchTaskWithContext(ctx context.Context, request *GetQualityWatchTaskRequest, runtime *dara.RuntimeOptions) (_result *GetQualityWatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchTaskId) {
		query["WatchTaskId"] = request.WatchTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityWatchTask"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityWatchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the log content of a monitoring task object.
//
// Online version: v5.4.2.
//
// @param request - GetQualityWatchTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityWatchTaskLogResponse
func (client *Client) GetQualityWatchTaskLogWithContext(ctx context.Context, request *GetQualityWatchTaskLogRequest, runtime *dara.RuntimeOptions) (_result *GetQualityWatchTaskLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.WatchTaskId) {
		query["WatchTaskId"] = request.WatchTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityWatchTaskLog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityWatchTaskLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the cluster version based on the cluster ID.
//
// @param request - GetQueueEngineVersionByEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueEngineVersionByEnvResponse
func (client *Client) GetQueueEngineVersionByEnvWithContext(ctx context.Context, request *GetQueueEngineVersionByEnvRequest, runtime *dara.RuntimeOptions) (_result *GetQueueEngineVersionByEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of a resource file.
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithContext(ctx context.Context, request *GetResourceRequest, runtime *dara.RuntimeOptions) (_result *GetResourceResponse, _err error) {
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
// Gets the details of a specified version of a resource file.
//
// @param request - GetResourceByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceByVersionResponse
func (client *Client) GetResourceByVersionWithContext(ctx context.Context, request *GetResourceByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetResourceByVersionResponse, _err error) {
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
// Release version: v5.4.2.
//
// @param tmpReq - GetRowPermissionByTableGuidsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRowPermissionByTableGuidsResponse
func (client *Client) GetRowPermissionByTableGuidsWithContext(ctx context.Context, tmpReq *GetRowPermissionByTableGuidsRequest, runtime *dara.RuntimeOptions) (_result *GetRowPermissionByTableGuidsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetRowPermissionByTableGuidsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GetRowPermissionByTableGuidsQuery) {
		request.GetRowPermissionByTableGuidsQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GetRowPermissionByTableGuidsQuery, dara.String("GetRowPermissionByTableGuidsQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GetRowPermissionByTableGuidsQueryShrink) {
		body["GetRowPermissionByTableGuidsQuery"] = request.GetRowPermissionByTableGuidsQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRowPermissionByTableGuids"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRowPermissionByTableGuidsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data classification. Release version: v5.4.2.
//
// @param request - GetSecurityClassifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityClassifyResponse
func (client *Client) GetSecurityClassifyWithContext(ctx context.Context, request *GetSecurityClassifyRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityClassifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetSecurityClassify"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityClassifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an identification result.
//
// Release version: v5.4.2.
//
// @param request - GetSecurityIdentifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityIdentifyResultResponse
func (client *Client) GetSecurityIdentifyResultWithContext(ctx context.Context, request *GetSecurityIdentifyResultRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityIdentifyResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("GetSecurityIdentifyResult"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityIdentifyResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data classification level. Available since v5.4.2.
//
// @param request - GetSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityLevelResponse
func (client *Client) GetSecurityLevelWithContext(ctx context.Context, request *GetSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Index) {
		query["Index"] = request.Index
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityLevel"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityLevelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a key value by key name. Online version: v5.4.2.
//
// @param request - GetSecuritySecretKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecuritySecretKeyResponse
func (client *Client) GetSecuritySecretKeyWithContext(ctx context.Context, request *GetSecuritySecretKeyRequest, runtime *dara.RuntimeOptions) (_result *GetSecuritySecretKeyResponse, _err error) {
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

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecuritySecretKey"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecuritySecretKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the Spark client information of the cluster associated with a compute source.
//
// @param request - GetSparkLocalClientInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkLocalClientInfoResponse
func (client *Client) GetSparkLocalClientInfoWithContext(ctx context.Context, request *GetSparkLocalClientInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSparkLocalClientInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of a standard.
//
// Release version: v5.4.2.
//
// @param tmpReq - GetStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardResponse
func (client *Client) GetStandardWithContext(ctx context.Context, tmpReq *GetStandardRequest, runtime *dara.RuntimeOptions) (_result *GetStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetStandardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StandardGetQuery) {
		request.StandardGetQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StandardGetQuery, dara.String("StandardGetQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.StandardGetQueryShrink) {
		body["StandardGetQuery"] = request.StandardGetQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard lookup table.
//
// Online version: v5.4.2.
//
// @param request - GetStandardLookupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardLookupTableResponse
func (client *Client) GetStandardLookupTableWithContext(ctx context.Context, request *GetStandardLookupTableRequest, runtime *dara.RuntimeOptions) (_result *GetStandardLookupTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Nullable) {
		query["Nullable"] = request.Nullable
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardLookupTable"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardLookupTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a standard set.
//
// Release version: v5.4.2.
//
// @param request - GetStandardSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardSetResponse
func (client *Client) GetStandardSetWithContext(ctx context.Context, request *GetStandardSetRequest, runtime *dara.RuntimeOptions) (_result *GetStandardSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Nullable) {
		query["Nullable"] = request.Nullable
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardSet"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of standards grouped by standard type under a specified folder.
//
// Online version: v5.4.2.
//
// @param tmpReq - GetStandardStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardStatisticsResponse
func (client *Client) GetStandardStatisticsWithContext(ctx context.Context, tmpReq *GetStandardStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetStandardStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetStandardStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StatisticsQuery) {
		request.StatisticsQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatisticsQuery, dara.String("StatisticsQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.StatisticsQueryShrink) {
		body["StatisticsQuery"] = request.StatisticsQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardStatistics"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard template.
//
// Online version: v5.4.2.
//
// @param tmpReq - GetStandardTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardTemplateResponse
func (client *Client) GetStandardTemplateWithContext(ctx context.Context, tmpReq *GetStandardTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetStandardTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetStandardTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterQuery) {
		request.FilterQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterQuery, dara.String("FilterQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Nullable) {
		query["Nullable"] = request.Nullable
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilterQueryShrink) {
		body["FilterQuery"] = request.FilterQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data standard word root.
//
// Online version: v5.4.2.
//
// @param request - GetStandardWordRootRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardWordRootResponse
func (client *Client) GetStandardWordRootWithContext(ctx context.Context, request *GetStandardWordRootRequest, runtime *dara.RuntimeOptions) (_result *GetStandardWordRootResponse, _err error) {
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

	if !dara.IsNil(request.Nullable) {
		query["Nullable"] = request.Nullable
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardWordRoot"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardWordRootResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of real-time development nodes.
//
// @param request - GetStreamJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStreamJobsResponse
func (client *Client) GetStreamJobsWithContext(ctx context.Context, request *GetStreamJobsRequest, runtime *dara.RuntimeOptions) (_result *GetStreamJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves dagrun information for all business dates of a data backfill instance workflow.
//
// @param request - GetSupplementDagrunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunResponse
func (client *Client) GetSupplementDagrunWithContext(ctx context.Context, request *GetSupplementDagrunRequest, runtime *dara.RuntimeOptions) (_result *GetSupplementDagrunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Lists the instances of all nodes for a specific business date in a data backfill workflow.
//
// @param request - GetSupplementDagrunInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSupplementDagrunInstanceResponse
func (client *Client) GetSupplementDagrunInstanceWithContext(ctx context.Context, request *GetSupplementDagrunInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetSupplementDagrunInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries table column lineage information.
//
// @param tmpReq - GetTableColumnLineageByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnLineageByTaskIdResponse
func (client *Client) GetTableColumnLineageByTaskIdWithContext(ctx context.Context, tmpReq *GetTableColumnLineageByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnLineageByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the column-level data lineage of an asset table.
//
// Online version: v5.4.2.
//
// @param tmpReq - GetTableColumnLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnLineagesResponse
func (client *Client) GetTableColumnLineagesWithContext(ctx context.Context, tmpReq *GetTableColumnLineagesRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnLineagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTableColumnLineagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterQuery) {
		request.FilterQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterQuery, dara.String("FilterQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilterQueryShrink) {
		body["FilterQuery"] = request.FilterQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableColumnLineages"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableColumnLineagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries columns of a Dataphin table in the asset inventory. Supported table types: dimension logical table, fact logical table, aggregate logical table, tag logical table, logical table view, physical table, physical view, and materialized view.
//
// Release version: v5.4.2.
//
// @param request - GetTableColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnsResponse
func (client *Client) GetTableColumnsWithContext(ctx context.Context, request *GetTableColumnsRequest, runtime *dara.RuntimeOptions) (_result *GetTableColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Catalog) {
		query["Catalog"] = request.Catalog
	}

	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableColumns"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableColumnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries table lineage information.
//
// @param tmpReq - GetTableLineageByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableLineageByTaskIdResponse
func (client *Client) GetTableLineageByTaskIdWithContext(ctx context.Context, tmpReq *GetTableLineageByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *GetTableLineageByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries lineage information of an asset table.
//
// Release version: v5.4.2.
//
// @param tmpReq - GetTableLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableLineagesResponse
func (client *Client) GetTableLineagesWithContext(ctx context.Context, tmpReq *GetTableLineagesRequest, runtime *dara.RuntimeOptions) (_result *GetTableLineagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTableLineagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterQuery) {
		request.FilterQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterQuery, dara.String("FilterQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FilterQueryShrink) {
		body["FilterQuery"] = request.FilterQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableLineages"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableLineagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress of a transfer task by transfer task ID.
//
// @param request - GetTransferInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransferInfoResponse
func (client *Client) GetTransferInfoWithContext(ctx context.Context, request *GetTransferInfoRequest, runtime *dara.RuntimeOptions) (_result *GetTransferInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProposalId) {
		query["ProposalId"] = request.ProposalId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTransferInfo"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTransferInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a user-defined function.
//
// @param request - GetUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfResponse
func (client *Client) GetUdfWithContext(ctx context.Context, request *GetUdfRequest, runtime *dara.RuntimeOptions) (_result *GetUdfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the details of a specific version of a user-defined function.
//
// @param request - GetUdfByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUdfByVersionResponse
func (client *Client) GetUdfByVersionWithContext(ctx context.Context, request *GetUdfByVersionRequest, runtime *dara.RuntimeOptions) (_result *GetUdfByVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves user details by original user ID.
//
// @param request - GetUserBySourceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserBySourceIdResponse
func (client *Client) GetUserBySourceIdWithContext(ctx context.Context, request *GetUserBySourceIdRequest, runtime *dara.RuntimeOptions) (_result *GetUserBySourceIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the details of a user group.
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithContext(ctx context.Context, request *GetUserGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves user information in batches by user ID.
//
// @param tmpReq - GetUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUsersResponse
func (client *Client) GetUsersWithContext(ctx context.Context, tmpReq *GetUsersRequest, runtime *dara.RuntimeOptions) (_result *GetUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Grants API authorization.
//
// @param tmpReq - GrantDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantDataServiceApiResponse
func (client *Client) GrantDataServiceApiWithContext(ctx context.Context, tmpReq *GrantDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *GrantDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Grants permissions on resources to users by resource point.
//
// @param tmpReq - GrantResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantResourcePermissionResponse
func (client *Client) GrantResourcePermissionWithContext(ctx context.Context, tmpReq *GrantResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantResourcePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the global roles that can be assigned to tenant members. Only built-in global roles are supported. Custom global roles are not supported.
//
// @param request - ListAddableRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableRolesResponse
func (client *Client) ListAddableRolesWithContext(ctx context.Context, request *ListAddableRolesRequest, runtime *dara.RuntimeOptions) (_result *ListAddableRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries users that can be added to a tenant. Only the super administrator (SuperAdmin) and system administrator can call this operation. The users must already exist in the Dataphin instance member list but have not yet been added to the tenant member list.
//
// @param tmpReq - ListAddableUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddableUsersResponse
func (client *Client) ListAddableUsersWithContext(ctx context.Context, tmpReq *ListAddableUsersRequest, runtime *dara.RuntimeOptions) (_result *ListAddableUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Performs a conditional query to list multiple alerting events.
//
// @param tmpReq - ListAlertEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertEventsResponse
func (client *Client) ListAlertEventsWithContext(ctx context.Context, tmpReq *ListAlertEventsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Performs a conditional query to list multiple push records.
//
// @param tmpReq - ListAlertNotificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertNotificationsResponse
func (client *Client) ListAlertNotificationsWithContext(ctx context.Context, tmpReq *ListAlertNotificationsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertNotificationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of APIs by application.
//
// Description:
//
// Queries the detailed information of published APIs by appKey.
//
// @param tmpReq - ListApiByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiByAppResponse
func (client *Client) ListApiByAppWithContext(ctx context.Context, tmpReq *ListApiByAppRequest, runtime *dara.RuntimeOptions) (_result *ListApiByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of specific fields for APIs that an application has requested.
//
// @param tmpReq - ListAuthorizedDataServiceApiDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedDataServiceApiDetailsResponse
func (client *Client) ListAuthorizedDataServiceApiDetailsWithContext(ctx context.Context, tmpReq *ListAuthorizedDataServiceApiDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedDataServiceApiDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries offline computing templates by paging. Online version: v6.2.0.
//
// @param tmpReq - ListBatchTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBatchTemplatesResponse
func (client *Client) ListBatchTemplatesWithContext(ctx context.Context, tmpReq *ListBatchTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListBatchTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListBatchTemplatesShrinkRequest{}
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
		Action:      dara.String("ListBatchTemplates"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBatchTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of business entities.
//
// @param tmpReq - ListBizEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizEntitiesResponse
func (client *Client) ListBizEntitiesWithContext(ctx context.Context, tmpReq *ListBizEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListBizEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves all business units under the current tenant.
//
// @param request - ListBizUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizUnitsResponse
func (client *Client) ListBizUnitsWithContext(ctx context.Context, request *ListBizUnitsRequest, runtime *dara.RuntimeOptions) (_result *ListBizUnitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of asset catalog entries. Online version: v6.1.0.
//
// @param tmpReq - ListCatalogAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCatalogAssetsResponse
func (client *Client) ListCatalogAssetsWithContext(ctx context.Context, tmpReq *ListCatalogAssetsRequest, runtime *dara.RuntimeOptions) (_result *ListCatalogAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCatalogAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListCatalogAssetsQuery) {
		request.ListCatalogAssetsQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListCatalogAssetsQuery, dara.String("ListCatalogAssetsQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListCatalogAssetsQueryShrink) {
		body["ListCatalogAssetsQuery"] = request.ListCatalogAssetsQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCatalogAssets"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCatalogAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of compute sources.
//
// @param tmpReq - ListComputeSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeSourcesResponse
func (client *Client) ListComputeSourcesWithContext(ctx context.Context, tmpReq *ListComputeSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListComputeSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves a list of data domains.
//
// @param tmpReq - ListDataDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataDomainsResponse
func (client *Client) ListDataDomainsWithContext(ctx context.Context, tmpReq *ListDataDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDataDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// O&M analysis: API call statistics.
//
// @param tmpReq - ListDataServiceApiCallStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiCallStatisticsResponse
func (client *Client) ListDataServiceApiCallStatisticsWithContext(ctx context.Context, tmpReq *ListDataServiceApiCallStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiCallStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries data service call logs with pagination.
//
// @param tmpReq - ListDataServiceApiCallsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiCallsResponse
func (client *Client) ListDataServiceApiCallsWithContext(ctx context.Context, tmpReq *ListDataServiceApiCallsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiCallsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Operations analysis: analyzes the impact of abnormal API calls.
//
// @param tmpReq - ListDataServiceApiImpactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiImpactsResponse
func (client *Client) ListDataServiceApiImpactsWithContext(ctx context.Context, tmpReq *ListDataServiceApiImpactsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiImpactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of all applications under a data service tenant. All tenant members can perform this operation.
//
// Release version: v6.0.0.
//
// @param tmpReq - ListDataServiceAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceAppsResponse
func (client *Client) ListDataServiceAppsWithContext(ctx context.Context, tmpReq *ListDataServiceAppsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataServiceAppsShrinkRequest{}
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
		Action:      dara.String("ListDataServiceApps"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceAppsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of applications that the current user has permissions to access.
//
// @param tmpReq - ListDataServiceAuthorizedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceAuthorizedAppsResponse
func (client *Client) ListDataServiceAuthorizedAppsWithContext(ctx context.Context, tmpReq *ListDataServiceAuthorizedAppsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceAuthorizedAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Obtain the list of API permissions managed by me.
//
// @param tmpReq - ListDataServiceMyApiPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceMyApiPermissionsResponse
func (client *Client) ListDataServiceMyApiPermissionsWithContext(ctx context.Context, tmpReq *ListDataServiceMyApiPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceMyApiPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the applications that the current user has permissions to access.
//
// @param tmpReq - ListDataServiceMyAppPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceMyAppPermissionsResponse
func (client *Client) ListDataServiceMyAppPermissionsWithContext(ctx context.Context, tmpReq *ListDataServiceMyAppPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceMyAppPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of published APIs by page.
//
// @param tmpReq - ListDataServicePublishedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServicePublishedApisResponse
func (client *Client) ListDataServicePublishedApisWithContext(ctx context.Context, tmpReq *ListDataServicePublishedApisRequest, runtime *dara.RuntimeOptions) (_result *ListDataServicePublishedApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Search for data sources. The results include data source configuration items.
//
// @param tmpReq - ListDataSourceWithConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceWithConfigResponse
func (client *Client) ListDataSourceWithConfigWithContext(ctx context.Context, tmpReq *ListDataSourceWithConfigRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceWithConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Lists datasets in a project based on specified conditional query criteria. Online version: v6.2.0.
//
// Description:
//
// ## Operation description
//
// This API allows you to retrieve dataset information for a specific project by providing a tenant ID, project ID, and other optional parameters such as keywords and type lists. Paging is supported. The returned data includes basic dataset information and version details. ProjectId is required. Other parameters are optional and can be configured as needed to filter results.
//
// @param tmpReq - ListDatasetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithContext(ctx context.Context, tmpReq *ListDatasetsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDatasetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DatasetQuery) {
		request.DatasetQueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DatasetQuery, dara.String("DatasetQuery"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetQueryShrink) {
		body["DatasetQuery"] = request.DatasetQueryShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasets"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the directory tree file list.
//
// @param tmpReq - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithContext(ctx context.Context, tmpReq *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Paginate and query instances.
//
// @param tmpReq - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, tmpReq *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the downstream nodes of a node. The query results can be used as a data reference when you create a data backfill workflow.
//
// @param tmpReq - ListNodeDownStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDownStreamResponse
func (client *Client) ListNodeDownStreamWithContext(ctx context.Context, tmpReq *ListNodeDownStreamRequest, runtime *dara.RuntimeOptions) (_result *ListNodeDownStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves a list of scheduling nodes.
//
// @param tmpReq - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithContext(ctx context.Context, tmpReq *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Performs a paged query on the list of operation records. Online version: v6.2.0.
//
// @param tmpReq - ListOperationRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationRecordResponse
func (client *Client) ListOperationRecordWithContext(ctx context.Context, tmpReq *ListOperationRecordRequest, runtime *dara.RuntimeOptions) (_result *ListOperationRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListOperationRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListCommand) {
		request.ListCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListCommand, dara.String("ListCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ListCommandShrink) {
		body["ListCommand"] = request.ListCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationRecord"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of project members.
//
// @param tmpReq - ListProjectMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithContext(ctx context.Context, tmpReq *ListProjectMembersRequest, runtime *dara.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves a list of projects.
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves a paginated list of publish records.
//
// @param tmpReq - ListPublishRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishRecordsResponse
func (client *Client) ListPublishRecordsWithContext(ctx context.Context, tmpReq *ListPublishRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListPublishRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries quality rule tasks by paging.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityRuleTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityRuleTasksResponse
func (client *Client) ListQualityRuleTasksWithContext(ctx context.Context, tmpReq *ListQualityRuleTasksRequest, runtime *dara.RuntimeOptions) (_result *ListQualityRuleTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityRuleTasksShrinkRequest{}
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
		Action:      dara.String("ListQualityRuleTasks"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityRuleTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality rules by paging.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityRulesResponse
func (client *Client) ListQualityRulesWithContext(ctx context.Context, tmpReq *ListQualityRulesRequest, runtime *dara.RuntimeOptions) (_result *ListQualityRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityRulesShrinkRequest{}
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
		Action:      dara.String("ListQualityRules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality templates by paging.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityTemplatesResponse
func (client *Client) ListQualityTemplatesWithContext(ctx context.Context, tmpReq *ListQualityTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListQualityTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityTemplatesShrinkRequest{}
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
		Action:      dara.String("ListQualityTemplates"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality monitoring nodes by paged query.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityWatchTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityWatchTasksResponse
func (client *Client) ListQualityWatchTasksWithContext(ctx context.Context, tmpReq *ListQualityWatchTasksRequest, runtime *dara.RuntimeOptions) (_result *ListQualityWatchTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityWatchTasksShrinkRequest{}
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
		Action:      dara.String("ListQualityWatchTasks"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityWatchTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paged query of quality monitored objects.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListQualityWatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityWatchesResponse
func (client *Client) ListQualityWatchesWithContext(ctx context.Context, tmpReq *ListQualityWatchesRequest, runtime *dara.RuntimeOptions) (_result *ListQualityWatchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListQualityWatchesShrinkRequest{}
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
		Action:      dara.String("ListQualityWatches"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityWatchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of permission operation logs.
//
// @param tmpReq - ListResourcePermissionOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionOperationLogResponse
func (client *Client) ListResourcePermissionOperationLogWithContext(ctx context.Context, tmpReq *ListResourcePermissionOperationLogRequest, runtime *dara.RuntimeOptions) (_result *ListResourcePermissionOperationLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves permission authorization records with pagination.
//
// @param tmpReq - ListResourcePermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcePermissionsResponse
func (client *Client) ListResourcePermissionsWithContext(ctx context.Context, tmpReq *ListResourcePermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListResourcePermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Performs a paged query of row-level permissions.
//
// @param tmpReq - ListRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRowPermissionResponse
func (client *Client) ListRowPermissionWithContext(ctx context.Context, tmpReq *ListRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *ListRowPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries row-level permissions of a specified user by paging.
//
// @param tmpReq - ListRowPermissionByUserIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRowPermissionByUserIdResponse
func (client *Client) ListRowPermissionByUserIdWithContext(ctx context.Context, tmpReq *ListRowPermissionByUserIdRequest, runtime *dara.RuntimeOptions) (_result *ListRowPermissionByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries identification records of security identification results by paging.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListSecurityIdentifyRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityIdentifyRecordsResponse
func (client *Client) ListSecurityIdentifyRecordsWithContext(ctx context.Context, tmpReq *ListSecurityIdentifyRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityIdentifyRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSecurityIdentifyRecordsShrinkRequest{}
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
		Action:      dara.String("ListSecurityIdentifyRecords"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityIdentifyRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query security identification results by page.
//
// Release version: v5.4.2.
//
// @param tmpReq - ListSecurityIdentifyResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityIdentifyResultsResponse
func (client *Client) ListSecurityIdentifyResultsWithContext(ctx context.Context, tmpReq *ListSecurityIdentifyResultsRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityIdentifyResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSecurityIdentifyResultsShrinkRequest{}
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
		Action:      dara.String("ListSecurityIdentifyResults"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityIdentifyResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the standard list by page.
//
// Release version: v5.4.2.
//
// @param tmpReq - ListStandardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStandardsResponse
func (client *Client) ListStandardsWithContext(ctx context.Context, tmpReq *ListStandardsRequest, runtime *dara.RuntimeOptions) (_result *ListStandardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListStandardsShrinkRequest{}
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
		Action:      dara.String("ListStandards"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStandardsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Paginate and retrieve the list of pending deployment records.
//
// @param tmpReq - ListSubmitRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubmitRecordsResponse
func (client *Client) ListSubmitRecordsWithContext(ctx context.Context, tmpReq *ListSubmitRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListSubmitRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Performs a paged query to retrieve asset table metadata.
//
// Online version: v5.4.2.
//
// @param tmpReq - ListTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithContext(ctx context.Context, tmpReq *ListTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTablesShrinkRequest{}
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
		Action:      dara.String("ListTables"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of tenant members.
//
// @param tmpReq - ListTenantMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantMembersResponse
func (client *Client) ListTenantMembersWithContext(ctx context.Context, tmpReq *ListTenantMembersRequest, runtime *dara.RuntimeOptions) (_result *ListTenantMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Performs a paging query of user group members.
//
// @param tmpReq - ListUserGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupMembersResponse
func (client *Client) ListUserGroupMembersWithContext(ctx context.Context, tmpReq *ListUserGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries user groups by paging.
//
// @param tmpReq - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithContext(ctx context.Context, tmpReq *ListUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Offlines a batch task.
//
// @param request - OfflineBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineBatchTaskResponse
func (client *Client) OfflineBatchTaskWithContext(ctx context.Context, request *OfflineBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *OfflineBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Offline a business entity.
//
// @param tmpReq - OfflineBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineBizEntityResponse
func (client *Client) OfflineBizEntityWithContext(ctx context.Context, tmpReq *OfflineBizEntityRequest, runtime *dara.RuntimeOptions) (_result *OfflineBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Offlines an integration pipeline node.
//
// @param tmpReq - OfflinePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflinePipelineResponse
func (client *Client) OfflinePipelineWithContext(ctx context.Context, tmpReq *OfflinePipelineRequest, runtime *dara.RuntimeOptions) (_result *OfflinePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflinePipelineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OfflineCommand) {
		request.OfflineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OfflineCommand, dara.String("OfflineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.OfflineCommandShrink) {
		body["OfflineCommand"] = request.OfflineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflinePipeline"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflinePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously offlines an integration pipeline node.
//
// @param tmpReq - OfflinePipelineByAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflinePipelineByAsyncResponse
func (client *Client) OfflinePipelineByAsyncWithContext(ctx context.Context, tmpReq *OfflinePipelineByAsyncRequest, runtime *dara.RuntimeOptions) (_result *OfflinePipelineByAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflinePipelineByAsyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OfflineCommand) {
		request.OfflineCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OfflineCommand, dara.String("OfflineCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.OfflineCommandShrink) {
		body["OfflineCommand"] = request.OfflineCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflinePipelineByAsync"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflinePipelineByAsyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Offlines a standard.
//
// Online version: v5.4.2.
//
// @param tmpReq - OfflineStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineStandardResponse
func (client *Client) OfflineStandardWithContext(ctx context.Context, tmpReq *OfflineStandardRequest, runtime *dara.RuntimeOptions) (_result *OfflineStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OfflineStandardShrinkRequest{}
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
		Action:      dara.String("OfflineStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineStandardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Brings a business entity online.
//
// @param tmpReq - OnlineBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineBizEntityResponse
func (client *Client) OnlineBizEntityWithContext(ctx context.Context, tmpReq *OnlineBizEntityRequest, runtime *dara.RuntimeOptions) (_result *OnlineBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Performs batch O&M operations on instances. Both physical instances and logical table instances are supported.
//
// @param tmpReq - OperateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateInstanceResponse
func (client *Client) OperateInstanceWithContext(ctx context.Context, tmpReq *OperateInstanceRequest, runtime *dara.RuntimeOptions) (_result *OperateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Parses the logical table dependencies of an offline compute node. The parsing result may contain self-dependent nodes in the upstream dependency information, where the upstream node ID is the same as the node ID of the parsed code. You must handle such cases on your own.
//
// @param tmpReq - ParseBatchTaskDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseBatchTaskDependencyResponse
func (client *Client) ParseBatchTaskDependencyWithContext(ctx context.Context, tmpReq *ParseBatchTaskDependencyRequest, runtime *dara.RuntimeOptions) (_result *ParseBatchTaskDependencyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Pauses the scheduling of physical nodes. This stops the scheduling of nodes, and downstream nodes cannot be triggered. Currently, only offline code nodes and integration nodes are supported.
//
// @param tmpReq - PausePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PausePhysicalNodeResponse
func (client *Client) PausePhysicalNodeWithContext(ctx context.Context, tmpReq *PausePhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *PausePhysicalNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Publishes a data service API to the production environment.
//
// @param request - PublishDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDataServiceApiResponse
func (client *Client) PublishDataServiceApiWithContext(ctx context.Context, request *PublishDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *PublishDataServiceApiResponse, _err error) {
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
// Publishes objects in batches.
//
// @param tmpReq - PublishObjectListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishObjectListResponse
func (client *Client) PublishObjectListWithContext(ctx context.Context, tmpReq *PublishObjectListRequest, runtime *dara.RuntimeOptions) (_result *PublishObjectListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Publishes a standard.
//
// Release version: v5.4.2.
//
// @param tmpReq - PublishStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishStandardResponse
func (client *Client) PublishStandardWithContext(ctx context.Context, tmpReq *PublishStandardRequest, runtime *dara.RuntimeOptions) (_result *PublishStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PublishStandardShrinkRequest{}
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
		Action:      dara.String("PublishStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishStandardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove regular members from a data service application. Only the application owner can perform this operation.
//
// Released version: v6.0.0.
//
// @param tmpReq - RemoveDataServiceAppMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDataServiceAppMemberResponse
func (client *Client) RemoveDataServiceAppMemberWithContext(ctx context.Context, tmpReq *RemoveDataServiceAppMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveDataServiceAppMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveDataServiceAppMemberShrinkRequest{}
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
		Action:      dara.String("RemoveDataServiceAppMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDataServiceAppMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a project member.
//
// @param tmpReq - RemoveProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveProjectMemberResponse
func (client *Client) RemoveProjectMemberWithContext(ctx context.Context, tmpReq *RemoveProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes the bindings between quality rules and schedules in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - RemoveQualityRuleSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveQualityRuleSchedulesResponse
func (client *Client) RemoveQualityRuleSchedulesWithContext(ctx context.Context, tmpReq *RemoveQualityRuleSchedulesRequest, runtime *dara.RuntimeOptions) (_result *RemoveQualityRuleSchedulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveQualityRuleSchedulesShrinkRequest{}
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
		Action:      dara.String("RemoveQualityRuleSchedules"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveQualityRuleSchedulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a tenant member. Only superusers and system administrators can call this API operation.
//
// @param tmpReq - RemoveTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTenantMemberResponse
func (client *Client) RemoveTenantMemberWithContext(ctx context.Context, tmpReq *RemoveTenantMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveTenantMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Removes members from a user group.
//
// @param tmpReq - RemoveUserGroupMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserGroupMemberResponse
func (client *Client) RemoveUserGroupMemberWithContext(ctx context.Context, tmpReq *RemoveUserGroupMemberRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates the whitelist of a project.
//
// @param tmpReq - ReplaceProjectWhiteListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceProjectWhiteListsResponse
func (client *Client) ReplaceProjectWhiteListsWithContext(ctx context.Context, tmpReq *ReplaceProjectWhiteListsRequest, runtime *dara.RuntimeOptions) (_result *ReplaceProjectWhiteListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Reset the Data Service application key. Only the application owner can perform this operation.
//
// Release version: v6.0.0.
//
// @param tmpReq - ResetDataServiceAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDataServiceAppSecretResponse
func (client *Client) ResetDataServiceAppSecretWithContext(ctx context.Context, tmpReq *ResetDataServiceAppSecretRequest, runtime *dara.RuntimeOptions) (_result *ResetDataServiceAppSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResetDataServiceAppSecretShrinkRequest{}
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
		Action:      dara.String("ResetDataServiceAppSecret"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetDataServiceAppSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resume physical node scheduling.
//
// @param tmpReq - ResumePhysicalNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumePhysicalNodeResponse
func (client *Client) ResumePhysicalNodeWithContext(ctx context.Context, tmpReq *ResumePhysicalNodeRequest, runtime *dara.RuntimeOptions) (_result *ResumePhysicalNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retransfers a failed transfer task.
//
// @param tmpReq - RetryTransferOwnershipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryTransferOwnershipResponse
func (client *Client) RetryTransferOwnershipWithContext(ctx context.Context, tmpReq *RetryTransferOwnershipRequest, runtime *dara.RuntimeOptions) (_result *RetryTransferOwnershipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RetryTransferOwnershipShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PrivilegeTransferRecord) {
		request.PrivilegeTransferRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PrivilegeTransferRecord, dara.String("PrivilegeTransferRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PrivilegeTransferRecordShrink) {
		body["PrivilegeTransferRecord"] = request.PrivilegeTransferRecordShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryTransferOwnership"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryTransferOwnershipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes API authorization.
//
// @param tmpReq - RevokeDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeDataServiceApiResponse
func (client *Client) RevokeDataServiceApiWithContext(ctx context.Context, tmpReq *RevokeDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *RevokeDataServiceApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Revokes resource authorization from a user.
//
// @param tmpReq - RevokeResourcePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeResourcePermissionResponse
func (client *Client) RevokeResourcePermissionWithContext(ctx context.Context, tmpReq *RevokeResourcePermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeResourcePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Stops an ad hoc query task.
//
// @param request - StopAdHocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAdHocTaskResponse
func (client *Client) StopAdHocTaskWithContext(ctx context.Context, request *StopAdHocTaskRequest, runtime *dara.RuntimeOptions) (_result *StopAdHocTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Submits a batch task.
//
// @param tmpReq - SubmitBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitBatchTaskResponse
func (client *Client) SubmitBatchTaskWithContext(ctx context.Context, tmpReq *SubmitBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Batch submit rule tasks with support for test runs.
//
// Release version: v5.4.2.
//
// @param tmpReq - SubmitQualityRuleTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitQualityRuleTasksResponse
func (client *Client) SubmitQualityRuleTasksWithContext(ctx context.Context, tmpReq *SubmitQualityRuleTasksRequest, runtime *dara.RuntimeOptions) (_result *SubmitQualityRuleTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitQualityRuleTasksShrinkRequest{}
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
		Action:      dara.String("SubmitQualityRuleTasks"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitQualityRuleTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits quality watchtask check tasks in batches.
//
// Online version: v5.4.2.
//
// @param tmpReq - SubmitQualityWatchTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitQualityWatchTasksResponse
func (client *Client) SubmitQualityWatchTasksWithContext(ctx context.Context, tmpReq *SubmitQualityWatchTasksRequest, runtime *dara.RuntimeOptions) (_result *SubmitQualityWatchTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitQualityWatchTasksShrinkRequest{}
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
		Action:      dara.String("SubmitQualityWatchTasks"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitQualityWatchTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes department information.
//
// Description:
//
// Queries the details of a published API operation by AppKey.
//
// @param tmpReq - SyncDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDepartmentResponse
func (client *Client) SyncDepartmentWithContext(ctx context.Context, tmpReq *SyncDepartmentRequest, runtime *dara.RuntimeOptions) (_result *SyncDepartmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SyncDepartmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SyncDepartmentCommand) {
		request.SyncDepartmentCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SyncDepartmentCommand, dara.String("SyncDepartmentCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SyncDepartmentCommandShrink) {
		body["SyncDepartmentCommand"] = request.SyncDepartmentCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDepartment"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDepartmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes department member information.
//
// @param tmpReq - SyncDepartmentUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDepartmentUserResponse
func (client *Client) SyncDepartmentUserWithContext(ctx context.Context, tmpReq *SyncDepartmentUserRequest, runtime *dara.RuntimeOptions) (_result *SyncDepartmentUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SyncDepartmentUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SyncDepartmentUserCommand) {
		request.SyncDepartmentUserCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SyncDepartmentUserCommand, dara.String("SyncDepartmentUserCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SyncDepartmentUserCommandShrink) {
		body["SyncDepartmentUserCommand"] = request.SyncDepartmentUserCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDepartmentUser"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDepartmentUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfers ownership to a new owner in one click.
//
// @param tmpReq - TransferOwnershipForAllObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferOwnershipForAllObjectResponse
func (client *Client) TransferOwnershipForAllObjectWithContext(ctx context.Context, tmpReq *TransferOwnershipForAllObjectRequest, runtime *dara.RuntimeOptions) (_result *TransferOwnershipForAllObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TransferOwnershipForAllObjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PrivilegeTransferRecord) {
		request.PrivilegeTransferRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PrivilegeTransferRecord, dara.String("PrivilegeTransferRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PrivilegeTransferRecordShrink) {
		body["PrivilegeTransferRecord"] = request.PrivilegeTransferRecordShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferOwnershipForAllObject"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferOwnershipForAllObjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an ad hoc query file.
//
// @param tmpReq - UpdateAdHocFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdHocFileResponse
func (client *Client) UpdateAdHocFileWithContext(ctx context.Context, tmpReq *UpdateAdHocFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdHocFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates an offline compute node.
//
// @param tmpReq - UpdateBatchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchTaskResponse
func (client *Client) UpdateBatchTaskWithContext(ctx context.Context, tmpReq *UpdateBatchTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Edits the custom data lineage of a batch task.
//
// @param tmpReq - UpdateBatchTaskUdfLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBatchTaskUdfLineagesResponse
func (client *Client) UpdateBatchTaskUdfLineagesWithContext(ctx context.Context, tmpReq *UpdateBatchTaskUdfLineagesRequest, runtime *dara.RuntimeOptions) (_result *UpdateBatchTaskUdfLineagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates a business entity.
//
// @param tmpReq - UpdateBizEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizEntityResponse
func (client *Client) UpdateBizEntityWithContext(ctx context.Context, tmpReq *UpdateBizEntityRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Update a business metric.
//
// Release version: v5.5.0.
//
// @param tmpReq - UpdateBizMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizMetricResponse
func (client *Client) UpdateBizMetricWithContext(ctx context.Context, tmpReq *UpdateBizMetricRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBizMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateBizMetricCommand) {
		request.UpdateBizMetricCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateBizMetricCommand, dara.String("UpdateBizMetricCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpdateBizMetricCommandShrink) {
		body["UpdateBizMetricCommand"] = request.UpdateBizMetricCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBizMetric"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBizMetricResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data domain.
//
// @param tmpReq - UpdateBizUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizUnitResponse
func (client *Client) UpdateBizUnitWithContext(ctx context.Context, tmpReq *UpdateBizUnitRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Edits a compute source. Business unit administrators and project administrators have permissions to perform this operation.
//
// @param tmpReq - UpdateComputeSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeSourceResponse
func (client *Client) UpdateComputeSourceWithContext(ctx context.Context, tmpReq *UpdateComputeSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates a data domain.
//
// @param tmpReq - UpdateDataDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataDomainResponse
func (client *Client) UpdateDataDomainWithContext(ctx context.Context, tmpReq *UpdateDataDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates a data service application. Only super administrators, system administrators, and application owners can perform this operation.
//
// Release version: v6.0.0.
//
// @param tmpReq - UpdateDataServiceAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataServiceAppResponse
func (client *Client) UpdateDataServiceAppWithContext(ctx context.Context, tmpReq *UpdateDataServiceAppRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataServiceAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataServiceAppShrinkRequest{}
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
		Action:      dara.String("UpdateDataServiceApp"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataServiceAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data service application group. Only superusers and system administrators can perform this operation.
//
// Online version: v6.0.0.
//
// @param tmpReq - UpdateDataServiceAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataServiceAppGroupResponse
func (client *Client) UpdateDataServiceAppGroupWithContext(ctx context.Context, tmpReq *UpdateDataServiceAppGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataServiceAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataServiceAppGroupShrinkRequest{}
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
		Action:      dara.String("UpdateDataServiceAppGroup"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataServiceAppGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the regular members of a data service application. Only the application owner can perform this operation.
//
// Online version: v6.0.0.
//
// @param tmpReq - UpdateDataServiceAppMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataServiceAppMemberResponse
func (client *Client) UpdateDataServiceAppMemberWithContext(ctx context.Context, tmpReq *UpdateDataServiceAppMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataServiceAppMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataServiceAppMemberShrinkRequest{}
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
		Action:      dara.String("UpdateDataServiceAppMember"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataServiceAppMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits the basic information of a data source. Tenant administrators, data administrators, business segment administrators, project administrators, and operations administrators have permissions to perform this operation.
//
// @param tmpReq - UpdateDataSourceBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceBasicInfoResponse
func (client *Client) UpdateDataSourceBasicInfoWithContext(ctx context.Context, tmpReq *UpdateDataSourceBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Edits the connection configuration items of a data source. Tenant administrators, data administrators, business unit administrators, project administrators, and operations administrators have permissions to perform this operation.
//
// @param tmpReq - UpdateDataSourceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceConfigResponse
func (client *Client) UpdateDataSourceConfigWithContext(ctx context.Context, tmpReq *UpdateDataSourceConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Dataphin OpenAPI 模板。
//
// Description:
//
// ## 请求说明
//
// - 该 API 用于更新特定项目下已存在的数据集的详细信息。
//
// - 必须提供 `ProjectId` 和 `UpdateCommand` 参数，其中 `UpdateCommand` 包含了需要更新的数据集的具体字段。
//
// - `UpdateCommand` 中的 `Id` 字段是必需的，用来标识要更新的数据集。
//
// - 其他字段如 `Name`, `Type`, `DataCellId` 等为可选项，根据实际需求选择性填写。
//
// - 版本配置（`VersionConfig`）和实时元表配置（`RealtimeMetaTableConfig`）提供了更详细的设置选项，包括存储路径、表结构等，这些也是可选的。
//
// - 注意确保所有提供的 ID 值（如 `ProjectId`, `Id`, `DataSourceId` 等）在系统中有效且正确关联。
//
// @param tmpReq - UpdateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithContext(ctx context.Context, tmpReq *UpdateDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
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
		Action:      dara.String("UpdateDataset"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves the file position in the menu tree.
//
// @param request - UpdateFileDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileDirectoryResponse
func (client *Client) UpdateFileDirectoryWithContext(ctx context.Context, request *UpdateFileDirectoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileDirectoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates a file name.
//
// @param request - UpdateFileNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileNameResponse
func (client *Client) UpdateFileNameWithContext(ctx context.Context, request *UpdateFileNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates an integration pipeline or unstructured workflow node.
//
// @param tmpReq - UpdatePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineResponse
func (client *Client) UpdatePipelineWithContext(ctx context.Context, tmpReq *UpdatePipelineRequest, runtime *dara.RuntimeOptions) (_result *UpdatePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePipelineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipeline"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
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
// Asynchronously updates a pipeline or unstructured workflow node.
//
// @param tmpReq - UpdatePipelineByAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineByAsyncResponse
func (client *Client) UpdatePipelineByAsyncWithContext(ctx context.Context, tmpReq *UpdatePipelineByAsyncRequest, runtime *dara.RuntimeOptions) (_result *UpdatePipelineByAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePipelineByAsyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Context) {
		request.ContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Context, dara.String("Context"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateCommand) {
		request.UpdateCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateCommand, dara.String("UpdateCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContextShrink) {
		body["Context"] = request.ContextShrink
	}

	if !dara.IsNil(request.UpdateCommandShrink) {
		body["UpdateCommand"] = request.UpdateCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipelineByAsync"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineByAsyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits project members.
//
// @param tmpReq - UpdateProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectMemberResponse
func (client *Client) UpdateProjectMemberWithContext(ctx context.Context, tmpReq *UpdateProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Enables or disables quality rules in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateQualityRuleSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityRuleSwitchResponse
func (client *Client) UpdateQualityRuleSwitchWithContext(ctx context.Context, tmpReq *UpdateQualityRuleSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityRuleSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateQualityRuleSwitchShrinkRequest{}
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
		Action:      dara.String("UpdateQualityRuleSwitch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityRuleSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts or stops quality monitored objects in batches.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateQualityWatchSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityWatchSwitchResponse
func (client *Client) UpdateQualityWatchSwitchWithContext(ctx context.Context, tmpReq *UpdateQualityWatchSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityWatchSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateQualityWatchSwitchShrinkRequest{}
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
		Action:      dara.String("UpdateQualityWatchSwitch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityWatchSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a resource file.
//
// @param tmpReq - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithContext(ctx context.Context, tmpReq *UpdateResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates a row-level permission.
//
// @param tmpReq - UpdateRowPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRowPermissionResponse
func (client *Client) UpdateRowPermissionWithContext(ctx context.Context, tmpReq *UpdateRowPermissionRequest, runtime *dara.RuntimeOptions) (_result *UpdateRowPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates a data classification.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateSecurityClassifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityClassifyResponse
func (client *Client) UpdateSecurityClassifyWithContext(ctx context.Context, tmpReq *UpdateSecurityClassifyRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityClassifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityClassifyShrinkRequest{}
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
		Action:      dara.String("UpdateSecurityClassify"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityClassifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data classification folder. Release version: v5.4.2.
//
// @param tmpReq - UpdateSecurityClassifyCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityClassifyCatalogResponse
func (client *Client) UpdateSecurityClassifyCatalogWithContext(ctx context.Context, tmpReq *UpdateSecurityClassifyCatalogRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityClassifyCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityClassifyCatalogShrinkRequest{}
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
		Action:      dara.String("UpdateSecurityClassifyCatalog"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityClassifyCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the effective status of security identification results.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateSecurityIdentifyResultStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityIdentifyResultStatusResponse
func (client *Client) UpdateSecurityIdentifyResultStatusWithContext(ctx context.Context, tmpReq *UpdateSecurityIdentifyResultStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityIdentifyResultStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityIdentifyResultStatusShrinkRequest{}
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
		Action:      dara.String("UpdateSecurityIdentifyResultStatus"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityIdentifyResultStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates data classification.
//
// Online version: v5.4.2.
//
// @param tmpReq - UpdateSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityLevelResponse
func (client *Client) UpdateSecurityLevelWithContext(ctx context.Context, tmpReq *UpdateSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityLevelShrinkRequest{}
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
		Action:      dara.String("UpdateSecurityLevel"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityLevelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a standard.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardResponse
func (client *Client) UpdateStandardWithContext(ctx context.Context, tmpReq *UpdateStandardRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardShrinkRequest{}
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
		Action:      dara.String("UpdateStandard"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data standard lookup table.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardLookupTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardLookupTableResponse
func (client *Client) UpdateStandardLookupTableWithContext(ctx context.Context, tmpReq *UpdateStandardLookupTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardLookupTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardLookupTableShrinkRequest{}
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
		Action:      dara.String("UpdateStandardLookupTable"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardLookupTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the standard mapping relationship to invalid mapping.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardMappingToInvalidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardMappingToInvalidResponse
func (client *Client) UpdateStandardMappingToInvalidWithContext(ctx context.Context, tmpReq *UpdateStandardMappingToInvalidRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardMappingToInvalidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardMappingToInvalidShrinkRequest{}
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
		Action:      dara.String("UpdateStandardMappingToInvalid"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardMappingToInvalidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update standard set.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardSetResponse
func (client *Client) UpdateStandardSetWithContext(ctx context.Context, tmpReq *UpdateStandardSetRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardSetShrinkRequest{}
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
		Action:      dara.String("UpdateStandardSet"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data standard template.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpdateStandardTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardTemplateResponse
func (client *Client) UpdateStandardTemplateWithContext(ctx context.Context, tmpReq *UpdateStandardTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardTemplateShrinkRequest{}
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
		Action:      dara.String("UpdateStandardTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data standard root word.
//
// Online version: v5.4.2.
//
// @param tmpReq - UpdateStandardWordRootRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardWordRootResponse
func (client *Client) UpdateStandardWordRootWithContext(ctx context.Context, tmpReq *UpdateStandardWordRootRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardWordRootResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStandardWordRootShrinkRequest{}
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
		Action:      dara.String("UpdateStandardWordRoot"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardWordRootResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the compute settings of a tenant.
//
// @param tmpReq - UpdateTenantComputeEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantComputeEngineResponse
func (client *Client) UpdateTenantComputeEngineWithContext(ctx context.Context, tmpReq *UpdateTenantComputeEngineRequest, runtime *dara.RuntimeOptions) (_result *UpdateTenantComputeEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Edits tenant members.
//
// @param tmpReq - UpdateTenantMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantMemberResponse
func (client *Client) UpdateTenantMemberWithContext(ctx context.Context, tmpReq *UpdateTenantMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateTenantMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Edits a user-defined function.
//
// @param tmpReq - UpdateUdfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfResponse
func (client *Client) UpdateUdfWithContext(ctx context.Context, tmpReq *UpdateUdfRequest, runtime *dara.RuntimeOptions) (_result *UpdateUdfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Edits a user group.
//
// @param tmpReq - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithContext(ctx context.Context, tmpReq *UpdateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies the enabled status of a user group.
//
// @param request - UpdateUserGroupSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupSwitchResponse
func (client *Client) UpdateUserGroupSwitchWithContext(ctx context.Context, request *UpdateUserGroupSwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Create or modify a quality rule.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityRuleResponse
func (client *Client) UpsertQualityRuleWithContext(ctx context.Context, tmpReq *UpsertQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualityRule"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates scheduling settings.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityScheduleResponse
func (client *Client) UpsertQualityScheduleWithContext(ctx context.Context, tmpReq *UpsertQualityScheduleRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityScheduleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualitySchedule"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a quality template.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityTemplateResponse
func (client *Client) UpsertQualityTemplateWithContext(ctx context.Context, tmpReq *UpsertQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualityTemplate"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a quality monitored object. You can add multiple types of quality monitored objects, including Dataphin tables, global tables, data sources, metrics, and real-time meta tables.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityWatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityWatchResponse
func (client *Client) UpsertQualityWatchWithContext(ctx context.Context, tmpReq *UpsertQualityWatchRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityWatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityWatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualityWatch"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityWatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates alert settings for a monitored object.
//
// Release version: v5.4.2.
//
// @param tmpReq - UpsertQualityWatchAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertQualityWatchAlertResponse
func (client *Client) UpsertQualityWatchAlertWithContext(ctx context.Context, tmpReq *UpsertQualityWatchAlertRequest, runtime *dara.RuntimeOptions) (_result *UpsertQualityWatchAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpsertQualityWatchAlertShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpsertCommand) {
		request.UpsertCommandShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpsertCommand, dara.String("UpsertCommand"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OpTenantId) {
		query["OpTenantId"] = request.OpTenantId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UpsertCommandShrink) {
		body["UpsertCommand"] = request.UpsertCommandShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertQualityWatchAlert"),
		Version:     dara.String("2023-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertQualityWatchAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
