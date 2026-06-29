// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates an account and directly adds it as a member.
//
// @param request - AddOrganizationMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrganizationMemberResponse
func (client *Client) AddOrganizationMemberWithContext(ctx context.Context, request *AddOrganizationMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddOrganizationMemberResponse, _err error) {
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

	if !dara.IsNil(request.OrgRoleCode) {
		query["OrgRoleCode"] = request.OrgRoleCode
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddOrganizationMember"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/organization/member-additions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOrganizationMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns seats in batches to the member level.
//
// @param request - BatchAssignSeatsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAssignSeatsResponse
func (client *Client) BatchAssignSeatsWithContext(ctx context.Context, request *BatchAssignSeatsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchAssignSeatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		query["AccountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.Locale) {
		query["Locale"] = request.Locale
	}

	if !dara.IsNil(request.SeatType) {
		query["SeatType"] = request.SeatType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAssignSeats"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/subscription/seat-assignments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAssignSeatsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes member-level seats in batches.
//
// @param tmpReq - BatchRevokeSeatsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchRevokeSeatsResponse
func (client *Client) BatchRevokeSeatsWithContext(ctx context.Context, tmpReq *BatchRevokeSeatsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchRevokeSeatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchRevokeSeatsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Items) {
		request.ItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Items, dara.String("Items"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ItemsShrink) {
		query["Items"] = request.ItemsShrink
	}

	if !dara.IsNil(request.Locale) {
		query["Locale"] = request.Locale
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchRevokeSeats"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/subscription/seat-revocations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchRevokeSeatsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Before using large models or applications in Alibaba Cloud Model Studio, create an API key as an authentication credential.
//
// @param request - CreateApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKeyWithContext(ctx context.Context, request *CreateApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		body["auth"] = request.Auth
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a TokenPlan member invitation link.
//
// Description:
//
// A user can have only one valid invitation link at a time.
//
// If the user already has a valid invitation link, this operation returns the existing link.
//
// To create a new link, call the RevokeTokenPlanInviteLink operation to invalidate the current link first.
//
// This operation returns only the generated token. The invitation link is assembled in the following format: `https://{host}/accept-invite?token=[token]&orgId=[orgId]`
//
//   - For the China site, the host is tokenplan-enterprise.bailian.aliyunportal.com.
//
//   - For the China site, the host is tokenplan-enterprise.modelstudio.aliyunportal.com.
//
// @param request - CreateTokenPlanInviteLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTokenPlanInviteLinkResponse
func (client *Client) CreateTokenPlanInviteLinkWithContext(ctx context.Context, request *CreateTokenPlanInviteLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTokenPlanInviteLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpireType) {
		query["ExpireType"] = request.ExpireType
	}

	if !dara.IsNil(request.SsoSource) {
		query["SsoSource"] = request.SsoSource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTokenPlanInviteLink"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/invite/link/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTokenPlanInviteLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a UAC API key.
//
// @param request - CreateTokenPlanKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTokenPlanKeyResponse
func (client *Client) CreateTokenPlanKeyWithContext(ctx context.Context, request *CreateTokenPlanKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTokenPlanKeyResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTokenPlanKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/api-keys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTokenPlanKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a business workspace.
//
// @param request - CreateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithContext(ctx context.Context, request *CreateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceSite) {
		query["serviceSite"] = request.ServiceSite
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["workspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/workspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Deletes an authentication credential API key.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKeyWithContext(ctx context.Context, apiKeyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApiKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// Description:
//
// A workspace can be deleted only if the following conditional requirements are met:
//
// 1. The workspace is not the default workspace.
//
// 2. The workspace is not used to purchase other products, such as AMB.
//
// 3. In permission management, the workspace is not granted to Resource Access Management (RAM) users or RAM roles.
//
// 4. The workspace does not contain any resources, such as API keys, model deployments, or knowledge bases.
//
// @param request - DeleteWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithContext(ctx context.Context, workspaceId *string, request *DeleteWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspace"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an API key.
//
// Description:
//
// An API key cannot be disabled if it is already disabled.
//
// @param request - DisableApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApiKeyResponse
func (client *Client) DisableApiKeyWithContext(ctx context.Context, apiKeyId *string, request *DisableApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId)) + "/disable"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an API key.
//
// Description:
//
// An API key that is already enabled cannot be enabled again.
//
// @param request - EnableApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApiKeyResponse
func (client *Client) EnableApiKeyWithContext(ctx context.Context, apiKeyId *string, request *EnableApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId)) + "/enable"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified authentication credential API key.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiKeyResponse
func (client *Client) GetApiKeyWithContext(ctx context.Context, apiKeyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApiKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a specified organization.
//
// Description:
//
// Retrieves information about a specified organization by OrgId.
//
// @param request - GetOrganizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationResponse
func (client *Client) GetOrganizationWithContext(ctx context.Context, request *GetOrganizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrganizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrganization"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/organization"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrganizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries organization member statistics information, including the total number of members, the number of administrators, the number of regular members, the number of members with allocated seats, and the number of members without allocated seats.
//
// @param request - GetOrganizationMemberSeatStatsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationMemberSeatStatsResponse
func (client *Client) GetOrganizationMemberSeatStatsWithContext(ctx context.Context, request *GetOrganizationMemberSeatStatsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrganizationMemberSeatStatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrganizationMemberSeatStats"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/organization/member-seat-stats"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrganizationMemberSeatStatsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries seat details by paging.
//
// @param request - GetSubscriptionSeatDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionSeatDetailsResponse
func (client *Client) GetSubscriptionSeatDetailsWithContext(ctx context.Context, request *GetSubscriptionSeatDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSubscriptionSeatDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryAssigned) {
		query["QueryAssigned"] = request.QueryAssigned
	}

	if !dara.IsNil(request.SeatId) {
		query["SeatId"] = request.SeatId
	}

	if !dara.IsNil(request.SeatType) {
		query["SeatType"] = request.SeatType
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubscriptionSeatDetails"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/subscription/seat-detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubscriptionSeatDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of members and seats for member management.
//
// @param request - GetSubscriptionStatsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionStatsResponse
func (client *Client) GetSubscriptionStatsWithContext(ctx context.Context, request *GetSubscriptionStatsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSubscriptionStatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubscriptionStats"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/subscription/stats"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubscriptionStatsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the TokenPlan account details and organization information.
//
// Description:
//
// Retrieves the TokenPlan management platform account information when the user is logged in.
//
// @param request - GetTokenPlanAccountDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenPlanAccountDetailResponse
func (client *Client) GetTokenPlanAccountDetailWithContext(ctx context.Context, request *GetTokenPlanAccountDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenPlanAccountDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTokenPlanAccountDetail"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/account"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenPlanAccountDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the TokenPlan member invitation link.
//
// Description:
//
// This operation returns only the generated token and expiration time. The invitation link is assembled in the following format: `https://{host}/accept-invite?token=[token]&orgId=[orgId]`
//
//   - For the China site, the host is tokenplan-enterprise.bailian.aliyunportal.com.
//
//   - For the international site, the host is tokenplan-enterprise.modelstudio.aliyunportal.com.
//
// @param request - GetTokenPlanInviteLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenPlanInviteLinkResponse
func (client *Client) GetTokenPlanInviteLinkWithContext(ctx context.Context, request *GetTokenPlanInviteLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenPlanInviteLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTokenPlanInviteLink"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/invite/link/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenPlanInviteLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the TokenPlan member invitation configuration.
//
// @param request - GetTokenPlanOrgInviteConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenPlanOrgInviteConfigResponse
func (client *Client) GetTokenPlanOrgInviteConfigWithContext(ctx context.Context, request *GetTokenPlanOrgInviteConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenPlanOrgInviteConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTokenPlanOrgInviteConfig"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/invite/config/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenPlanOrgInviteConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of API key authentication credentials.
//
// @param request - ListApiKeysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiKeysResponse
func (client *Client) ListApiKeysWithContext(ctx context.Context, request *ListApiKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiKeys"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of organization members including seat information. Supports filtering by name, status, and seat assignment, and supports pagination.
//
// @param request - ListOrganizationMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationMembersResponse
func (client *Client) ListOrganizationMembersWithContext(ctx context.Context, request *ListOrganizationMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOrganizationMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HasSeat) {
		query["HasSeat"] = request.HasSeat
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrganizationMembers"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/organization/members"),
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
// Queries the details of shared packages by paging.
//
// @param request - ListSubscriptionSharedPackagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubscriptionSharedPackagesResponse
func (client *Client) ListSubscriptionSharedPackagesWithContext(ctx context.Context, request *ListSubscriptionSharedPackagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSubscriptionSharedPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubscriptionSharedPackages"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/subscription/shared-packages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubscriptionSharedPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of business workspaces.
//
// @param request - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithContext(ctx context.Context, request *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
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

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["workspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/workspaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes organization members. Before removal, checks whether the member holds a seat. If the member holds a seat, the removal is rejected.
//
// @param request - RemoveOrganizationMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveOrganizationMemberResponse
func (client *Client) RemoveOrganizationMemberWithContext(ctx context.Context, request *RemoveOrganizationMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveOrganizationMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		query["AccountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.Locale) {
		query["Locale"] = request.Locale
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveOrganizationMember"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/organization/member-removals"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveOrganizationMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets an API key.
//
// Description:
//
// Only the API key value changes. The API key ID remains unchanged.
//
// @param request - ResetApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetApiKeyResponse
func (client *Client) ResetApiKeyWithContext(ctx context.Context, apiKeyId *string, request *ResetApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId)) + "/reset"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a TokenPlan member invitation link.
//
// @param request - RevokeTokenPlanInviteLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTokenPlanInviteLinkResponse
func (client *Client) RevokeTokenPlanInviteLinkWithContext(ctx context.Context, request *RevokeTokenPlanInviteLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeTokenPlanInviteLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeTokenPlanInviteLink"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/invite/link/revoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeTokenPlanInviteLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets a UAC API key.
//
// Description:
//
// Only the API Key value changes. The API Key ID remains unchanged.
//
// @param request - RotateTokenPlanKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RotateTokenPlanKeyResponse
func (client *Client) RotateTokenPlanKeyWithContext(ctx context.Context, request *RotateTokenPlanKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RotateTokenPlanKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["ApiKeyId"] = request.ApiKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RotateTokenPlanKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/api-key-rotations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RotateTokenPlanKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the member invitation settings for a TokenPlan.
//
// @param request - SetTokenPlanOrgInviteConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTokenPlanOrgInviteConfigResponse
func (client *Client) SetTokenPlanOrgInviteConfigWithContext(ctx context.Context, request *SetTokenPlanOrgInviteConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SetTokenPlanOrgInviteConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultRoleId) {
		query["DefaultRoleId"] = request.DefaultRoleId
	}

	if !dara.IsNil(request.SeatAssignStrategy) {
		query["SeatAssignStrategy"] = request.SeatAssignStrategy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTokenPlanOrgInviteConfig"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/invite/config/set"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTokenPlanOrgInviteConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits the information of an authentication credential API key.
//
// @param request - UpdateApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiKeyResponse
func (client *Client) UpdateApiKeyWithContext(ctx context.Context, apiKeyId *string, request *UpdateApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		body["auth"] = request.Auth
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApiKey"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/modelstudio/apikeys/" + dara.PercentEncode(dara.StringValue(apiKeyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies organization information.
//
// @param request - UpdateOrganizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationResponse
func (client *Client) UpdateOrganizationWithContext(ctx context.Context, request *UpdateOrganizationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOrganization"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/organization"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOrganizationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改组织成员角色
//
// @param request - UpdateOrganizationMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationMemberResponse
func (client *Client) UpdateOrganizationMemberWithContext(ctx context.Context, request *UpdateOrganizationMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateOrganizationMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		query["AccountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.NewRoleCode) {
		query["NewRoleCode"] = request.NewRoleCode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOrganizationMember"),
		Version:     dara.String("2026-02-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tokenplan/organization/members/update"),
		Method:      dara.String("POST"),
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
