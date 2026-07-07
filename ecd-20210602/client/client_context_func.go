// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a tenant skill.
//
// @param tmpReq - CreateTenantSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTenantSkillResponse
func (client *Client) CreateTenantSkillWithContext(ctx context.Context, tmpReq *CreateTenantSkillRequest, runtime *dara.RuntimeOptions) (_result *CreateTenantSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTenantSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EnvVars) {
		request.EnvVarsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvVars, dara.String("EnvVars"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EnvVarsShrink) {
		query["EnvVars"] = request.EnvVarsShrink
	}

	if !dara.IsNil(request.IconETag) {
		query["IconETag"] = request.IconETag
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIcon) {
		query["SkillIcon"] = request.SkillIcon
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	if !dara.IsNil(request.Slug) {
		query["Slug"] = request.Slug
	}

	if !dara.IsNil(request.TaskKey) {
		query["TaskKey"] = request.TaskKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTenantSkill"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTenantSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes skills in batches.
//
// @param request - DeleteTenantSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTenantSkillsResponse
func (client *Client) DeleteTenantSkillsWithContext(ctx context.Context, request *DeleteTenantSkillsRequest, runtime *dara.RuntimeOptions) (_result *DeleteTenantSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTenantSkills"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTenantSkillsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 OSS STS 令牌
//
// Description:
//
// 获取到的SecurityToken有效期为15分钟。
//
// @param request - GetOssStsTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssStsTokenResponse
func (client *Client) GetOssStsTokenWithContext(ctx context.Context, request *GetOssStsTokenRequest, runtime *dara.RuntimeOptions) (_result *GetOssStsTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssStsToken"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssStsTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the parsed content of a skill package.
//
// Description:
//
// Call the ParseSkillPackage operation first. Poll this operation every 3 seconds.
//
// @param request - GetParseProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParseProgressResponse
func (client *Client) GetParseProgressWithContext(ctx context.Context, request *GetParseProgressRequest, runtime *dara.RuntimeOptions) (_result *GetParseProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskKey) {
		query["TaskKey"] = request.TaskKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParseProgress"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParseProgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of identities for which security policies are enabled.
//
// Description:
//
// The resource type supports only cloud computers.
//
// @param request - ListSecureSkillIdentitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecureSkillIdentitiesResponse
func (client *Client) ListSecureSkillIdentitiesWithContext(ctx context.Context, request *ListSecureSkillIdentitiesRequest, runtime *dara.RuntimeOptions) (_result *ListSecureSkillIdentitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecureSkillIdentities"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecureSkillIdentitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of identities authorized for a skill.
//
// Description:
//
// Authorized objects support only cloud computers.
//
// @param request - ListSkillAuthedIdentitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillAuthedIdentitiesResponse
func (client *Client) ListSkillAuthedIdentitiesWithContext(ctx context.Context, request *ListSkillAuthedIdentitiesRequest, runtime *dara.RuntimeOptions) (_result *ListSkillAuthedIdentitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkillAuthedIdentities"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillAuthedIdentitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of skills.
//
// @param request - ListSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithContext(ctx context.Context, request *ListSkillsRequest, runtime *dara.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	if !dara.IsNil(request.SupplierType) {
		query["SupplierType"] = request.SupplierType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkills"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses a skill package.
//
// @param request - ParseSkillPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseSkillPackageResponse
func (client *Client) ParseSkillPackageWithContext(ctx context.Context, request *ParseSkillPackageRequest, runtime *dara.RuntimeOptions) (_result *ParseSkillPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OssObjectETag) {
		query["OssObjectETag"] = request.OssObjectETag
	}

	if !dara.IsNil(request.OssObjectKey) {
		query["OssObjectKey"] = request.OssObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ParseSkillPackage"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ParseSkillPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets skill permissions for an identity.
//
// Description:
//
// The authorized object supports only cloud computers.
//
// @param request - SetIdentitySkillAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdentitySkillAuthResponse
func (client *Client) SetIdentitySkillAuthWithContext(ctx context.Context, request *SetIdentitySkillAuthRequest, runtime *dara.RuntimeOptions) (_result *SetIdentitySkillAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoInstall) {
		query["AutoInstall"] = request.AutoInstall
	}

	if !dara.IsNil(request.Identities) {
		query["Identities"] = request.Identities
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIdentitySkillAuth"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIdentitySkillAuthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the security policy for identity skills.
//
// Description:
//
// The resource type supports only cloud computers.
//
// @param request - SetIdentitySkillSecurityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIdentitySkillSecurityResponse
func (client *Client) SetIdentitySkillSecurityWithContext(ctx context.Context, request *SetIdentitySkillSecurityRequest, runtime *dara.RuntimeOptions) (_result *SetIdentitySkillSecurityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.IdentityIds) {
		query["IdentityIds"] = request.IdentityIds
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIdentitySkillSecurity"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIdentitySkillSecurityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置租户技能启用状态
//
// @param request - SetTenantSkillEnabledRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTenantSkillEnabledResponse
func (client *Client) SetTenantSkillEnabledWithContext(ctx context.Context, request *SetTenantSkillEnabledRequest, runtime *dara.RuntimeOptions) (_result *SetTenantSkillEnabledResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.SkillChannel) {
		query["SkillChannel"] = request.SkillChannel
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTenantSkillEnabled"),
		Version:     dara.String("2021-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTenantSkillEnabledResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
