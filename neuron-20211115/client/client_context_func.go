// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 微服务分组加入或退出泳道
//
// @param request - AddOrQuitPdpLaneForServiceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrQuitPdpLaneForServiceGroupResponse
func (client *Client) AddOrQuitPdpLaneForServiceGroupWithContext(ctx context.Context, request *AddOrQuitPdpLaneForServiceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddOrQuitPdpLaneForServiceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddOrQuitPdpLaneForServiceGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/commands/add-quit-lane"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOrQuitPdpLaneForServiceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 审批审核记录
//
// @param request - AuditForkReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuditForkReviewResponse
func (client *Client) AuditForkReviewWithContext(ctx context.Context, reviewId *string, request *AuditForkReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuditForkReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuditForkReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/fork-reviews/" + dara.PercentEncode(dara.StringValue(reviewId)) + "/commands/audit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuditForkReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 审批审核记录
//
// @param request - AuditPbcInvokeReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuditPbcInvokeReviewResponse
func (client *Client) AuditPbcInvokeReviewWithContext(ctx context.Context, reviewId *string, request *AuditPbcInvokeReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuditPbcInvokeReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Approve) {
		body["approve"] = request.Approve
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuditPbcInvokeReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-invoke-reviews/" + dara.PercentEncode(dara.StringValue(reviewId)) + "/commands/audit"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuditPbcInvokeReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 公司授权产品
//
// @param request - AuthorizeProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeProductsResponse
func (client *Client) AuthorizeProductsWithContext(ctx context.Context, request *AuthorizeProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuthorizeProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeProducts"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/products/commands/authorize"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量给开发者授权
//
// @param request - BatchGrantRolesToDeveloperRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGrantRolesToDeveloperResponse
func (client *Client) BatchGrantRolesToDeveloperWithContext(ctx context.Context, request *BatchGrantRolesToDeveloperRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGrantRolesToDeveloperResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGrantRolesToDeveloper"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/commands/batch-grant-roles"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &BatchGrantRolesToDeveloperResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验用户角色
//
// @param request - CheckDeveloperRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDeveloperRoleResponse
func (client *Client) CheckDeveloperRoleWithContext(ctx context.Context, request *CheckDeveloperRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckDeveloperRoleResponse, _err error) {
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

	if !dara.IsNil(request.CompanyId) {
		body["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.Platform) {
		body["platform"] = request.Platform
	}

	if !dara.IsNil(request.RoleName) {
		body["roleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDeveloperRole"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/commands/check-role"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDeveloperRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 完成二方包注册
//
// @param request - CompleteRegisterLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteRegisterLibraryResponse
func (client *Client) CompleteRegisterLibraryWithContext(ctx context.Context, id *string, request *CompleteRegisterLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CompleteRegisterLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DependIntegral) {
		body["dependIntegral"] = request.DependIntegral
	}

	if !dara.IsNil(request.MarketId) {
		body["marketId"] = request.MarketId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteRegisterLibrary"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/complete-register"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteRegisterLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 完成pbc版本注册
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteRegistrationPbcVersionResponse
func (client *Client) CompleteRegistrationPbcVersionWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CompleteRegistrationPbcVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteRegistrationPbcVersion"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/complete-register"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteRegistrationPbcVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务信息
//
// @param request - ContinueDeploymentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDeploymentResponse
func (client *Client) ContinueDeploymentWithContext(ctx context.Context, request *ContinueDeploymentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ContinueDeploymentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinueDeployment"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/deployments/commands/continue"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinueDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关注资产
//
// @param request - CreateAssetWatchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAssetWatchResponse
func (client *Client) CreateAssetWatchWithContext(ctx context.Context, request *CreateAssetWatchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAssetWatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAssetWatch"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/asset-watchs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAssetWatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增组件
//
// @param request - CreateComponentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComponentResponse
func (client *Client) CreateComponentWithContext(ctx context.Context, request *CreateComponentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComponent"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/components"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增组件模板
//
// @param request - CreateComponentTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComponentTemplateResponse
func (client *Client) CreateComponentTemplateWithContext(ctx context.Context, request *CreateComponentTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateComponentTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComponentTemplate"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/component-templates"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComponentTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册开发者
//
// @param request - CreateDeveloperRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeveloperResponse
func (client *Client) CreateDeveloperWithContext(ctx context.Context, request *CreateDeveloperRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDeveloperResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeveloper"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/developers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeveloperResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建公司
//
// @param request - CreateEnterpriseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseResponse
func (client *Client) CreateEnterpriseWithContext(ctx context.Context, request *CreateEnterpriseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEnterpriseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnterprise"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/enterprises"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnterpriseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请代码Fork(创建代码Fork的审批流程)
//
// @param request - CreateForkReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateForkReviewResponse
func (client *Client) CreateForkReviewWithContext(ctx context.Context, request *CreateForkReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateForkReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateForkReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/fork-reviews"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateForkReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务灰度分组
//
// @param request - CreateGreyPdpServiceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGreyPdpServiceGroupResponse
func (client *Client) CreateGreyPdpServiceGroupWithContext(ctx context.Context, request *CreateGreyPdpServiceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGreyPdpServiceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGreyPdpServiceGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/groups/grey"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGreyPdpServiceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建二方包
//
// @param request - CreateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLibraryResponse
func (client *Client) CreateLibraryWithContext(ctx context.Context, request *CreateLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLibrary"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建二方包使用说明文档
//
// @param request - CreateLibraryInstructionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLibraryInstructionResponse
func (client *Client) CreateLibraryInstructionWithContext(ctx context.Context, libraryId *string, request *CreateLibraryInstructionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLibraryInstructionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		body["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.Document) {
		body["document"] = request.Document
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.MarketId) {
		body["marketId"] = request.MarketId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLibraryInstruction"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/instructions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLibraryInstructionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交审核
//
// @param request - CreateLibraryReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLibraryReviewResponse
func (client *Client) CreateLibraryReviewWithContext(ctx context.Context, request *CreateLibraryReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLibraryReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLibraryReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/library-reviews"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLibraryReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建二方包规格
//
// @param request - CreateLibrarySchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLibrarySchemaResponse
func (client *Client) CreateLibrarySchemaWithContext(ctx context.Context, request *CreateLibrarySchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLibrarySchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLibrarySchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/schemas"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLibrarySchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用报警任务
//
// @param request - CreateMonitorArmsAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorArmsAlertResponse
func (client *Client) CreateMonitorArmsAlertWithContext(ctx context.Context, request *CreateMonitorArmsAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorArmsAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorArmsAlert"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/createMonitorArmsAlert"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorArmsAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增应用报警任务
//
// @param request - CreateMonitorArmsAlertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorArmsAlertsResponse
func (client *Client) CreateMonitorArmsAlertsWithContext(ctx context.Context, request *CreateMonitorArmsAlertsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorArmsAlertsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorArmsAlerts"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/createMonitorArmsAlerts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorArmsAlertsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建联系人
//
// @param request - CreateMonitorContactRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorContactResponse
func (client *Client) CreateMonitorContactWithContext(ctx context.Context, request *CreateMonitorContactRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorContact"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/contact"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建联系人组
//
// @param request - CreateMonitorContactGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorContactGroupResponse
func (client *Client) CreateMonitorContactGroupWithContext(ctx context.Context, request *CreateMonitorContactGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorContactGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorContactGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 加入联系人
//
// @param request - CreateMonitorGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupMemberResponse
func (client *Client) CreateMonitorGroupMemberWithContext(ctx context.Context, groupId *string, request *CreateMonitorGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroupMember"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/group/" + dara.PercentEncode(dara.StringValue(groupId)) + "/commands/create-member"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateMonitorGroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建MQ报警任务
//
// @param request - CreateMonitorMqAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorMqAlertResponse
func (client *Client) CreateMonitorMqAlertWithContext(ctx context.Context, request *CreateMonitorMqAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorMqAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorMqAlert"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/createMonitorMqAlert"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorMqAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增MQ报警任务
//
// @param request - CreateMonitorMqAlertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorMqAlertsResponse
func (client *Client) CreateMonitorMqAlertsWithContext(ctx context.Context, request *CreateMonitorMqAlertsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorMqAlertsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorMqAlerts"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/createMonitorMqAlerts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorMqAlertsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建日志报警任务
//
// @param request - CreateMonitorSlsAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorSlsAlertResponse
func (client *Client) CreateMonitorSlsAlertWithContext(ctx context.Context, request *CreateMonitorSlsAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorSlsAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorSlsAlert"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/createMonitorSlsAlert"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorSlsAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增建报警任务
//
// @param request - CreateMonitorSlsAlertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorSlsAlertsResponse
func (client *Client) CreateMonitorSlsAlertsWithContext(ctx context.Context, request *CreateMonitorSlsAlertsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorSlsAlertsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorSlsAlerts"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/createMonitorSlsAlerts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorSlsAlertsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建webhook
//
// @param request - CreateMonitorWebhookRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorWebhookResponse
func (client *Client) CreateMonitorWebhookWithContext(ctx context.Context, request *CreateMonitorWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMonitorWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorWebhook"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/webhook"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Group
//
// @param request - CreateMqGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMqGroupResponse
func (client *Client) CreateMqGroupWithContext(ctx context.Context, request *CreateMqGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMqGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMqGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMqGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Topic
//
// @param request - CreateMqTopicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMqTopicResponse
func (client *Client) CreateMqTopicWithContext(ctx context.Context, request *CreateMqTopicRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMqTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMqTopic"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/topic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMqTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PBC
//
// @param request - CreatePbcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePbcResponse
func (client *Client) CreatePbcWithContext(ctx context.Context, request *CreatePbcRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePbcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PBC的Api规格
//
// @param request - CreatePbcApiSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePbcApiSchemaResponse
func (client *Client) CreatePbcApiSchemaWithContext(ctx context.Context, request *CreatePbcApiSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePbcApiSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePbcApiSchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/api-schemas"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePbcApiSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PBC使用说明文档
//
// @param request - CreatePbcInstructionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePbcInstructionResponse
func (client *Client) CreatePbcInstructionWithContext(ctx context.Context, request *CreatePbcInstructionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePbcInstructionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePbcInstruction"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/instructions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePbcInstructionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务调用记录
//
// @param request - CreatePbcInvokeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePbcInvokeResponse
func (client *Client) CreatePbcInvokeWithContext(ctx context.Context, request *CreatePbcInvokeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePbcInvokeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePbcInvoke"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-invokes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePbcInvokeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PBC服务调用审核
//
// @param request - CreatePbcInvokeReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePbcInvokeReviewResponse
func (client *Client) CreatePbcInvokeReviewWithContext(ctx context.Context, request *CreatePbcInvokeReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePbcInvokeReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePbcInvokeReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-invoke-reviews"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePbcInvokeReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交审核
//
// @param request - CreatePbcReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePbcReviewResponse
func (client *Client) CreatePbcReviewWithContext(ctx context.Context, request *CreatePbcReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePbcReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePbcReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-reviews"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePbcReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PBC规格
//
// @param request - CreatePbcSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePbcSchemaResponse
func (client *Client) CreatePbcSchemaWithContext(ctx context.Context, request *CreatePbcSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePbcSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePbcSchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/schemas"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePbcSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PBC版本
//
// @param request - CreatePbcVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePbcVersionResponse
func (client *Client) CreatePbcVersionWithContext(ctx context.Context, request *CreatePbcVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePbcVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePbcVersion"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePbcVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建配置
//
// @param request - CreatePdpConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdpConfigResponse
func (client *Client) CreatePdpConfigWithContext(ctx context.Context, request *CreatePdpConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePdpConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdpConfig"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/configs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdpConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建泳道
//
// @param request - CreatePdpLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdpLaneResponse
func (client *Client) CreatePdpLaneWithContext(ctx context.Context, request *CreatePdpLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePdpLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdpLane"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdpLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PdpPbc
//
// @param request - CreatePdpPbcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdpPbcResponse
func (client *Client) CreatePdpPbcWithContext(ctx context.Context, request *CreatePdpPbcRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePdpPbcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdpPbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/pbcs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdpPbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务
//
// @param request - CreatePdpServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdpServiceResponse
func (client *Client) CreatePdpServiceWithContext(ctx context.Context, request *CreatePdpServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePdpServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdpService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/services"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务分组
//
// @param request - CreatePdpServiceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdpServiceGroupResponse
func (client *Client) CreatePdpServiceGroupWithContext(ctx context.Context, request *CreatePdpServiceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePdpServiceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdpServiceGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdpServiceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给角色添加权限
//
// @param request - CreatePrivilegeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivilegeResponse
func (client *Client) CreatePrivilegeWithContext(ctx context.Context, roleId *string, request *CreatePrivilegeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivilege"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/" + dara.PercentEncode(dara.StringValue(roleId)) + "/privileges"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给角色添加权限
//
// @param request - CreatePrivilegePopRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivilegePopResponse
func (client *Client) CreatePrivilegePopWithContext(ctx context.Context, roleId *string, request *CreatePrivilegePopRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePrivilegePopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivilegePop"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/" + dara.PercentEncode(dara.StringValue(roleId)) + "/privileges/pop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivilegePopResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建产品
//
// @param request - CreateProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductResponse
func (client *Client) CreateProductWithContext(ctx context.Context, request *CreateProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProduct"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/products"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建代码库Fork记录
//
// @param request - CreateRepoForkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoForkResponse
func (client *Client) CreateRepoForkWithContext(ctx context.Context, request *CreateRepoForkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRepoForkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRepoFork"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/repo-forks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRepoForkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增资源
//
// @param request - CreateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithContext(ctx context.Context, request *CreateResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResource"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 生成角色
//
// @param request - CreateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRole"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建入驻申请
//
// @param request - CreateSettledRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSettledResponse
func (client *Client) CreateSettledWithContext(ctx context.Context, request *CreateSettledRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSettledResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSettled"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/settleds"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSettledResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除组件
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComponentResponse
func (client *Client) DeleteComponentWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteComponentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComponent"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/components/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除组件模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComponentTemplateResponse
func (client *Client) DeleteComponentTemplateWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteComponentTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComponentTemplate"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/component-templates/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteComponentTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除开发者
//
// @param request - DeleteDeveloperRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeveloperResponse
func (client *Client) DeleteDeveloperWithContext(ctx context.Context, accountId *string, request *DeleteDeveloperRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDeveloperResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeveloper"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/developers/" + dara.PercentEncode(dara.StringValue(accountId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteDeveloperResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除公司
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnterpriseResponse
func (client *Client) DeleteEnterpriseWithContext(ctx context.Context, enterpriseId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEnterpriseResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnterprise"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/enterprises/" + dara.PercentEncode(dara.StringValue(enterpriseId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteEnterpriseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除二方包
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLibraryResponse
func (client *Client) DeleteLibraryWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLibraryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLibrary"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除二方包使用说明文档
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLibraryInstructionResponse
func (client *Client) DeleteLibraryInstructionWithContext(ctx context.Context, libraryId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLibraryInstructionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLibraryInstruction"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/instructions"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteLibraryInstructionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除二方包规格
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLibrarySchemaResponse
func (client *Client) DeleteLibrarySchemaWithContext(ctx context.Context, libraryId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLibrarySchemaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLibrarySchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/schemas"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteLibrarySchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除报警任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorAlertResponse
func (client *Client) DeleteMonitorAlertWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMonitorAlertResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorAlert"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteMonitorAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除联系人
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorContactResponse
func (client *Client) DeleteMonitorContactWithContext(ctx context.Context, contactId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMonitorContactResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorContact"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/contact/" + dara.PercentEncode(dara.StringValue(contactId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteMonitorContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除联系人组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorContactGroupResponse
func (client *Client) DeleteMonitorContactGroupWithContext(ctx context.Context, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMonitorContactGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorContactGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/group/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteMonitorContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从联系人组中删除联系人
//
// @param request - DeleteMonitorGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupMemberResponse
func (client *Client) DeleteMonitorGroupMemberWithContext(ctx context.Context, groupId *string, request *DeleteMonitorGroupMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactIds) {
		query["contactIds"] = request.ContactIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroupMember"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/group/" + dara.PercentEncode(dara.StringValue(groupId)) + "/commands/delete-member"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteMonitorGroupMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除webhook
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorWebhookResponse
func (client *Client) DeleteMonitorWebhookWithContext(ctx context.Context, webhookId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMonitorWebhookResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorWebhook"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/webhook/" + dara.PercentEncode(dara.StringValue(webhookId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteMonitorWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除MQ Group
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMqGroupResponse
func (client *Client) DeleteMqGroupWithContext(ctx context.Context, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMqGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMqGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/group/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteMqGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Topic
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMqTopicResponse
func (client *Client) DeleteMqTopicWithContext(ctx context.Context, topicId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMqTopicResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMqTopic"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/topic/" + dara.PercentEncode(dara.StringValue(topicId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteMqTopicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消服务调用服务
//
// @param request - DeletePbcInvokeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePbcInvokeResponse
func (client *Client) DeletePbcInvokeWithContext(ctx context.Context, request *DeletePbcInvokeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePbcInvokeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Applicant) {
		query["applicant"] = request.Applicant
	}

	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePbcInvoke"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-invokes"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePbcInvokeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePdpConfigResponse
func (client *Client) DeletePdpConfigWithContext(ctx context.Context, configId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePdpConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePdpConfig"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/configs/" + dara.PercentEncode(dara.StringValue(configId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePdpConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除泳道
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePdpLaneResponse
func (client *Client) DeletePdpLaneWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePdpLaneResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePdpLane"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePdpLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定泳道入口微服务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePdpLaneInletServiceResponse
func (client *Client) DeletePdpLaneInletServiceWithContext(ctx context.Context, laneId *string, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePdpLaneInletServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePdpLaneInletService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/" + dara.PercentEncode(dara.StringValue(laneId)) + "/commands/deleted-inlet-service/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePdpLaneInletServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除所有泳道中含有指定的微服务组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePdpLaneServiceGroupResponse
func (client *Client) DeletePdpLaneServiceGroupWithContext(ctx context.Context, serviceGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePdpLaneServiceGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePdpLaneServiceGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/commands/deleted-service-group/" + dara.PercentEncode(dara.StringValue(serviceGroupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePdpLaneServiceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除PdpPbc
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePdpPbcResponse
func (client *Client) DeletePdpPbcWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePdpPbcResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePdpPbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/pbcs/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePdpPbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePdpServiceResponse
func (client *Client) DeletePdpServiceWithContext(ctx context.Context, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePdpServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePdpService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePdpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务分组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePdpServiceGroupResponse
func (client *Client) DeletePdpServiceGroupWithContext(ctx context.Context, serviceGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePdpServiceGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePdpServiceGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/groups/" + dara.PercentEncode(dara.StringValue(serviceGroupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeletePdpServiceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除权限
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivilegeResponse
func (client *Client) DeletePrivilegeWithContext(ctx context.Context, privilegeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePrivilegeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivilege"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/privileges/" + dara.PercentEncode(dara.StringValue(privilegeId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeletePrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询删除产品
//
// @param request - DeleteProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProductResponse
func (client *Client) DeleteProductWithContext(ctx context.Context, id *string, request *DeleteProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProduct"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/products/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定资源
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResource"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/resources/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
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
// 删除角色
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithContext(ctx context.Context, roleId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRole"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/" + dara.PercentEncode(dara.StringValue(roleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除入驻申请记录
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSettledResponse
func (client *Client) DeleteSettledWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSettledResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSettled"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/settleds/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSettledResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 依赖二方包
//
// @param request - DependLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DependLibraryResponse
func (client *Client) DependLibraryWithContext(ctx context.Context, id *string, request *DependLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DependLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DependLibrary"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/dependency"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DependLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 审批审核
//
// @param request - FeedbackLibraryReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FeedbackLibraryReviewResponse
func (client *Client) FeedbackLibraryReviewWithContext(ctx context.Context, reviewId *string, request *FeedbackLibraryReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FeedbackLibraryReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Approve) {
		body["approve"] = request.Approve
	}

	if !dara.IsNil(request.Feedback) {
		body["feedback"] = request.Feedback
	}

	if !dara.IsNil(request.ReviewId) {
		body["reviewId"] = request.ReviewId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FeedbackLibraryReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/library-reviews/" + dara.PercentEncode(dara.StringValue(reviewId)) + "/commands/feedback"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FeedbackLibraryReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交审核
//
// @param request - FeedbackPbcReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FeedbackPbcReviewResponse
func (client *Client) FeedbackPbcReviewWithContext(ctx context.Context, reviewId *string, request *FeedbackPbcReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FeedbackPbcReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FeedbackPbcReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-reviews/" + dara.PercentEncode(dara.StringValue(reviewId)) + "/commands/feedback"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FeedbackPbcReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Buc公司
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBucEnterpriseResponse
func (client *Client) GetBucEnterpriseWithContext(ctx context.Context, enterpriseId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBucEnterpriseResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBucEnterprise"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/bucs/enterprises/" + dara.PercentEncode(dara.StringValue(enterpriseId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBucEnterpriseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组件
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComponentResponse
func (client *Client) GetComponentWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetComponentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComponent"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/components/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组件模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComponentTemplateResponse
func (client *Client) GetComponentTemplateWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetComponentTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComponentTemplate"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/component-templates/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComponentTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询部署详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentResponse
func (client *Client) GetDeploymentWithContext(ctx context.Context, deploymentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeploymentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeployment"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/deployments/instance/" + dara.PercentEncode(dara.StringValue(deploymentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询开发者
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeveloperResponse
func (client *Client) GetDeveloperWithContext(ctx context.Context, accountId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeveloperResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeveloper"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/developers/" + dara.PercentEncode(dara.StringValue(accountId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeveloperResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询开发者所在公司
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeveloperEnterpriseResponse
func (client *Client) GetDeveloperEnterpriseWithContext(ctx context.Context, accountId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeveloperEnterpriseResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeveloperEnterprise"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/enterprises/developers/" + dara.PercentEncode(dara.StringValue(accountId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeveloperEnterpriseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询公司
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnterpriseResponse
func (client *Client) GetEnterpriseWithContext(ctx context.Context, enterpriseId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEnterpriseResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnterprise"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/enterprises/" + dara.PercentEncode(dara.StringValue(enterpriseId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnterpriseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询开发者
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnterpriseDeveloperResponse
func (client *Client) GetEnterpriseDeveloperWithContext(ctx context.Context, accountId *string, enterpriseId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEnterpriseDeveloperResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnterpriseDeveloper"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/developers/" + dara.PercentEncode(dara.StringValue(accountId)) + "/enterprises/" + dara.PercentEncode(dara.StringValue(enterpriseId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnterpriseDeveloperResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取审核详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetForkReviewResponse
func (client *Client) GetForkReviewWithContext(ctx context.Context, reviewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetForkReviewResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetForkReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/fork-reviews/" + dara.PercentEncode(dara.StringValue(reviewId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetForkReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询历史开发者
//
// @param request - GetHistoryDeveloperRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoryDeveloperResponse
func (client *Client) GetHistoryDeveloperWithContext(ctx context.Context, accountId *string, request *GetHistoryDeveloperRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHistoryDeveloperResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
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
		Action:      dara.String("GetHistoryDeveloper"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/developers/" + dara.PercentEncode(dara.StringValue(accountId)) + "/commonds/history"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHistoryDeveloperResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询最近一次部署配置
//
// @param request - GetLastDeploymentConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLastDeploymentConfigResponse
func (client *Client) GetLastDeploymentConfigWithContext(ctx context.Context, request *GetLastDeploymentConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLastDeploymentConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLastDeploymentConfig"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/deployments/last-config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLastDeploymentConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询最新版本二方库
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryResponse
func (client *Client) GetLibraryWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibrary"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取二方包代码库开发者统计信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryDeveloperRepoMetricsResponse
func (client *Client) GetLibraryDeveloperRepoMetricsWithContext(ctx context.Context, libraryId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryDeveloperRepoMetricsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibraryDeveloperRepoMetrics"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/Librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/code/commands/get-developer-repo-metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryDeveloperRepoMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询二方包使用说明文档
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryInstructionResponse
func (client *Client) GetLibraryInstructionWithContext(ctx context.Context, libraryId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryInstructionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibraryInstruction"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/instructions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryInstructionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取二方包的代码库统计信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryRepoMetricsResponse
func (client *Client) GetLibraryRepoMetricsWithContext(ctx context.Context, libraryId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryRepoMetricsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibraryRepoMetrics"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/Librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/code/commands/get-repo-metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryRepoMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审核信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryReviewResponse
func (client *Client) GetLibraryReviewWithContext(ctx context.Context, reviewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryReviewResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibraryReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/library-reviews/" + dara.PercentEncode(dara.StringValue(reviewId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询二方包规格
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibrarySchemaResponse
func (client *Client) GetLibrarySchemaWithContext(ctx context.Context, libraryId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibrarySchemaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibrarySchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/schemas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibrarySchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志链接
//
// @param request - GetLogUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogUrlResponse
func (client *Client) GetLogUrlWithContext(ctx context.Context, request *GetLogUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLogUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ip) {
		query["ip"] = request.Ip
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLogUrl"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/pdp-log/url"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLogUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取登录的buc用户信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginUserInfoResponse
func (client *Client) GetLoginUserInfoWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLoginUserInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoginUserInfo"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/bucs/logins"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoginUserInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报警任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonitorAlertResponse
func (client *Client) GetMonitorAlertWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMonitorAlertResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMonitorAlert"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMonitorAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报警历史
//
// @param request - GetMonitorAlertHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonitorAlertHistoryResponse
func (client *Client) GetMonitorAlertHistoryWithContext(ctx context.Context, request *GetMonitorAlertHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMonitorAlertHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertRuleName) {
		query["alertRuleName"] = request.AlertRuleName
	}

	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMonitorAlertHistory"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/getMonitorAlertHistory"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMonitorAlertHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonitorContactResponse
func (client *Client) GetMonitorContactWithContext(ctx context.Context, contactId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMonitorContactResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMonitorContact"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/contact/" + dara.PercentEncode(dara.StringValue(contactId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMonitorContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonitorContactGroupResponse
func (client *Client) GetMonitorContactGroupWithContext(ctx context.Context, groupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMonitorContactGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMonitorContactGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/group/" + dara.PercentEncode(dara.StringValue(groupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMonitorContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询webhook
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonitorWebhookResponse
func (client *Client) GetMonitorWebhookWithContext(ctx context.Context, webhookId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMonitorWebhookResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMonitorWebhook"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/webhook/" + dara.PercentEncode(dara.StringValue(webhookId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMonitorWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询最新版本pbc
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcResponse
func (client *Client) GetPbcWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PBC的api规格
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcApiSchemaResponse
func (client *Client) GetPbcApiSchemaWithContext(ctx context.Context, pbcVersionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcApiSchemaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbcApiSchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(pbcVersionId)) + "/api-schemas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcApiSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取PBC代码库开发者统计信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcDeveloperRepoMetricsResponse
func (client *Client) GetPbcDeveloperRepoMetricsWithContext(ctx context.Context, pbcId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcDeveloperRepoMetricsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbcDeveloperRepoMetrics"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/" + dara.PercentEncode(dara.StringValue(pbcId)) + "/code/commands/get-developer-repo-metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcDeveloperRepoMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PBC使用说明文档
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcInstructionResponse
func (client *Client) GetPbcInstructionWithContext(ctx context.Context, pbcVersionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcInstructionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbcInstruction"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(pbcVersionId)) + "/instructions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcInstructionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取审核详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcInvokeReviewResponse
func (client *Client) GetPbcInvokeReviewWithContext(ctx context.Context, reviewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcInvokeReviewResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbcInvokeReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-invoke-reviews/" + dara.PercentEncode(dara.StringValue(reviewId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcInvokeReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取PBC的代码库统计信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcRepoMetricsResponse
func (client *Client) GetPbcRepoMetricsWithContext(ctx context.Context, pbcId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcRepoMetricsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbcRepoMetrics"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/" + dara.PercentEncode(dara.StringValue(pbcId)) + "/code/commands/get-repo-metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcRepoMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询pbc审核信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcReviewResponse
func (client *Client) GetPbcReviewWithContext(ctx context.Context, reviewId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcReviewResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbcReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-reviews/" + dara.PercentEncode(dara.StringValue(reviewId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PBC规格
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcSchemaResponse
func (client *Client) GetPbcSchemaWithContext(ctx context.Context, pbcVersionId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcSchemaResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbcSchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(pbcVersionId)) + "/schemas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询pbc版本信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPbcVersionResponse
func (client *Client) GetPbcVersionWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPbcVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPbcVersion"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPbcVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取配置信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPdpConfigResponse
func (client *Client) GetPdpConfigWithContext(ctx context.Context, configId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPdpConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPdpConfig"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/configs/last/" + dara.PercentEncode(dara.StringValue(configId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPdpConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取历史配置详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPdpHistoryConfigResponse
func (client *Client) GetPdpHistoryConfigWithContext(ctx context.Context, historyConfigId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPdpHistoryConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPdpHistoryConfig"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/configs/history/" + dara.PercentEncode(dara.StringValue(historyConfigId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPdpHistoryConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询泳道详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPdpLaneResponse
func (client *Client) GetPdpLaneWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPdpLaneResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPdpLane"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPdpLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询PdpPbc
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPdpPbcResponse
func (client *Client) GetPdpPbcWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPdpPbcResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPdpPbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/pbcs/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPdpPbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPdpServiceResponse
func (client *Client) GetPdpServiceWithContext(ctx context.Context, serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPdpServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPdpService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPdpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务分组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPdpServiceGroupResponse
func (client *Client) GetPdpServiceGroupWithContext(ctx context.Context, serviceGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPdpServiceGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPdpServiceGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/groups/" + dara.PercentEncode(dara.StringValue(serviceGroupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPdpServiceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询产品
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductResponse
func (client *Client) GetProductWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProductResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProduct"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/products/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResource"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/resources/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 查询角色
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleResponse
func (client *Client) GetRoleWithContext(ctx context.Context, roleId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRole"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/role-id/" + dara.PercentEncode(dara.StringValue(roleId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取调用链方法栈
//
// @param request - GetStackDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackDetailResponse
func (client *Client) GetStackDetailWithContext(ctx context.Context, traceId *string, request *GetStackDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetStackDetailResponse, _err error) {
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

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.RpcId) {
		query["rpcId"] = request.RpcId
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackDetail"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/traces/" + dara.PercentEncode(dara.StringValue(traceId)) + "/commonds/stack"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取token
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithContext(ctx context.Context, request *GetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
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

	if !dara.IsNil(request.PbcId) {
		body["pbcId"] = request.PbcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/pbcs/token"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取调用链详情
//
// @param request - GetTraceDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceDetailResponse
func (client *Client) GetTraceDetailWithContext(ctx context.Context, traceId *string, request *GetTraceDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTraceDetailResponse, _err error) {
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

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTraceDetail"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/traces/" + dara.PercentEncode(dara.StringValue(traceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTraceDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权
//
// @param request - GrantRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantRoleResponse
func (client *Client) GrantRoleWithContext(ctx context.Context, roleId *string, request *GrantRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantRole"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/" + dara.PercentEncode(dara.StringValue(roleId)) + "/commands/grant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &GrantRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询可以授权的产品
//
// @param request - ListAuthorizeProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizeProductsResponse
func (client *Client) ListAuthorizeProductsWithContext(ctx context.Context, request *ListAuthorizeProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAuthorizeProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizeProducts"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/products/commands/list-authorize"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizeProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Buc用户的公司
//
// @param request - ListBucUserEnterpriseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBucUserEnterpriseResponse
func (client *Client) ListBucUserEnterpriseWithContext(ctx context.Context, request *ListBucUserEnterpriseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBucUserEnterpriseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EmpId) {
		query["empId"] = request.EmpId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBucUserEnterprise"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/bucs/enterprises"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBucUserEnterpriseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组件模板列表
//
// @param request - ListComponentTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentTemplatesResponse
func (client *Client) ListComponentTemplatesWithContext(ctx context.Context, request *ListComponentTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComponentTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponentTemplates"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/component-templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组件列表
//
// @param request - ListComponentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentsResponse
func (client *Client) ListComponentsWithContext(ctx context.Context, request *ListComponentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	if !dara.IsNil(request.TemplateId) {
		query["templateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponents"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/components"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询我依赖的二方包列表
//
// @param request - ListDependLibrarysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDependLibrarysResponse
func (client *Client) ListDependLibrarysWithContext(ctx context.Context, request *ListDependLibrarysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDependLibrarysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Applicant) {
		query["applicant"] = request.Applicant
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDependLibrarys"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/commands/list-dependency"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDependLibrarysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询部署记录列表
//
// @param tmpReq - ListDeploymentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentsResponse
func (client *Client) ListDeploymentsWithContext(ctx context.Context, tmpReq *ListDeploymentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDeploymentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExcludeStatus) {
		request.ExcludeStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeStatus, dara.String("excludeStatus"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Status) {
		request.StatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Status, dara.String("status"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ExcludeStatusShrink) {
		query["excludeStatus"] = request.ExcludeStatusShrink
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.StatusShrink) {
		query["status"] = request.StatusShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeployments"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/deployments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询开发者所在公司列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeveloperEnterprisesResponse
func (client *Client) ListDeveloperEnterprisesWithContext(ctx context.Context, accountId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeveloperEnterprisesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeveloperEnterprises"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/enterprises/developers/" + dara.PercentEncode(dara.StringValue(accountId)) + "/commands/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeveloperEnterprisesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前开发者拥有权限的pbc列表
//
// @param request - ListDeveloperPbcsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeveloperPbcsResponse
func (client *Client) ListDeveloperPbcsWithContext(ctx context.Context, request *ListDeveloperPbcsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeveloperPbcsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeveloperPbcs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/commands/allow-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeveloperPbcsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询开发者列表
//
// @param tmpReq - ListDevelopersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDevelopersResponse
func (client *Client) ListDevelopersWithContext(ctx context.Context, tmpReq *ListDevelopersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDevelopersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDevelopersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("accountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		query["accountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RoleId) {
		query["roleId"] = request.RoleId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDevelopers"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/developers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDevelopersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 此PBC依赖的PBC列表(下游游PBC)
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownstreamPbcResponse
func (client *Client) ListDownstreamPbcWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDownstreamPbcResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDownstreamPbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/list-downstream"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDownstreamPbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询公司列表
//
// @param request - ListEnterprisesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterprisesResponse
func (client *Client) ListEnterprisesWithContext(ctx context.Context, request *ListEnterprisesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEnterprisesResponse, _err error) {
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

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterprises"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/enterprises"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterprisesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 产销环境信息列表
//
// @param request - ListEnvInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvInfosResponse
func (client *Client) ListEnvInfosWithContext(ctx context.Context, request *ListEnvInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEnvInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvInfos"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/services/env/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审核列表，支持按照审核人、申请人查询
//
// @param request - ListForkReviewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListForkReviewsResponse
func (client *Client) ListForkReviewsWithContext(ctx context.Context, request *ListForkReviewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListForkReviewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Applicant) {
		query["applicant"] = request.Applicant
	}

	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reviewer) {
		query["reviewer"] = request.Reviewer
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListForkReviews"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/fork-reviews"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListForkReviewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询被授权角色列表
//
// @param request - ListGrantedRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGrantedRolesResponse
func (client *Client) ListGrantedRolesWithContext(ctx context.Context, request *ListGrantedRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGrantedRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizerId) {
		query["authorizerId"] = request.AuthorizerId
	}

	if !dara.IsNil(request.AuthorizerType) {
		query["authorizerType"] = request.AuthorizerType
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGrantedRoles"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/commands/list-granted-roles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGrantedRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询订阅当前组件的pbc的列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInvokePbcsResponse
func (client *Client) ListInvokePbcsWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInvokePbcsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInvokePbcs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/invoke-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInvokePbcsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审核信息列表
//
// @param request - ListLibraryReviewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLibraryReviewsResponse
func (client *Client) ListLibraryReviewsWithContext(ctx context.Context, request *ListLibraryReviewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLibraryReviewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Applicant) {
		query["applicant"] = request.Applicant
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reviewer) {
		query["reviewer"] = request.Reviewer
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLibraryReviews"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/library-reviews"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLibraryReviewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询二方包列表
//
// @param request - ListLibrarysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLibrarysResponse
func (client *Client) ListLibrarysWithContext(ctx context.Context, request *ListLibrarysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLibrarysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Provider) {
		query["provider"] = request.Provider
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLibrarys"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLibrarysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询市场列表
//
// @param request - ListMarketsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMarketsResponse
func (client *Client) ListMarketsWithContext(ctx context.Context, request *ListMarketsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMarketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMarkets"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/markets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMarketsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询/搜索元数据信息列表
//
// @param request - ListMetadataInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetadataInfosResponse
func (client *Client) ListMetadataInfosWithContext(ctx context.Context, request *ListMetadataInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMetadataInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.NamespaceId) {
		query["namespace_id"] = request.NamespaceId
	}

	if !dara.IsNil(request.NamespaceName) {
		query["namespace_name"] = request.NamespaceName
	}

	if !dara.IsNil(request.OrderBy) {
		query["order_by"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["order_direction"] = request.OrderDirection
	}

	if !dara.IsNil(request.OrgId) {
		query["org_id"] = request.OrgId
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.PbcId) {
		query["pbc_id"] = request.PbcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMetadataInfos"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/metadata"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetadataInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询微服务列表
//
// @param request - ListMicroServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMicroServiceResponse
func (client *Client) ListMicroServiceWithContext(ctx context.Context, env *string, request *ListMicroServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMicroServiceResponse, _err error) {
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

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	if !dara.IsNil(request.ServiceIds) {
		query["serviceIds"] = request.ServiceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMicroService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/group/env/" + dara.PercentEncode(dara.StringValue(env))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMicroServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人组列表
//
// @param request - ListMonitorContactGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMonitorContactGroupsResponse
func (client *Client) ListMonitorContactGroupsWithContext(ctx context.Context, request *ListMonitorContactGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMonitorContactGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMonitorContactGroups"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/group"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMonitorContactGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人列表
//
// @param request - ListMonitorContactsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMonitorContactsResponse
func (client *Client) ListMonitorContactsWithContext(ctx context.Context, request *ListMonitorContactsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMonitorContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMonitorContacts"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/contact"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMonitorContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有的联系人/联系人组/webhook列表
//
// @param request - ListMonitorNotifyObjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMonitorNotifyObjectsResponse
func (client *Client) ListMonitorNotifyObjectsWithContext(ctx context.Context, request *ListMonitorNotifyObjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMonitorNotifyObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.WebhookType) {
		query["webhookType"] = request.WebhookType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMonitorNotifyObjects"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/listMonitorNotifyObjects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMonitorNotifyObjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询通知报警列表
//
// @param request - ListMonitorTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMonitorTasksResponse
func (client *Client) ListMonitorTasksWithContext(ctx context.Context, request *ListMonitorTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMonitorTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertName) {
		query["alertName"] = request.AlertName
	}

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMonitorTasks"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/listMonitorTasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMonitorTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询webhook列表
//
// @param request - ListMonitorWebhooksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMonitorWebhooksResponse
func (client *Client) ListMonitorWebhooksWithContext(ctx context.Context, request *ListMonitorWebhooksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMonitorWebhooksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WebhookType) {
		query["webhookType"] = request.WebhookType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMonitorWebhooks"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/webhook"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMonitorWebhooksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Group下的死信消息列表
//
// @param request - ListMqGroupMsgsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMqGroupMsgsResponse
func (client *Client) ListMqGroupMsgsWithContext(ctx context.Context, groupId *string, request *ListMqGroupMsgsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMqGroupMsgsResponse, _err error) {
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

	if !dara.IsNil(request.MsgId) {
		query["msgId"] = request.MsgId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMqGroupMsgs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/group/" + dara.PercentEncode(dara.StringValue(groupId)) + "/commands/msgs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMqGroupMsgsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Topic下的消息列表
//
// @param request - ListMqTopicMsgsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMqTopicMsgsResponse
func (client *Client) ListMqTopicMsgsWithContext(ctx context.Context, topicId *string, request *ListMqTopicMsgsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMqTopicMsgsResponse, _err error) {
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

	if !dara.IsNil(request.MsgId) {
		query["msgId"] = request.MsgId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMqTopicMsgs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/topic/" + dara.PercentEncode(dara.StringValue(topicId)) + "/commands/msgs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMqTopicMsgsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Topic下的在线订阅列表
//
// @param request - ListMqTopicSubsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMqTopicSubsResponse
func (client *Client) ListMqTopicSubsWithContext(ctx context.Context, topicId *string, request *ListMqTopicSubsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMqTopicSubsResponse, _err error) {
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

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMqTopicSubs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/topic/" + dara.PercentEncode(dara.StringValue(topicId)) + "/commands/subs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMqTopicSubsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Topic列表
//
// @param request - ListMqTopicsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMqTopicsResponse
func (client *Client) ListMqTopicsWithContext(ctx context.Context, env *string, pbcId *string, request *ListMqTopicsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMqTopicsResponse, _err error) {
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

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TopicName) {
		query["topicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMqTopics"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/topic/env/" + dara.PercentEncode(dara.StringValue(env)) + "/pbcId/" + dara.PercentEncode(dara.StringValue(pbcId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMqTopicsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务调用审核列表，支持按照审核人、申请人查询
//
// @param request - ListPbcInvokeReviewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPbcInvokeReviewsResponse
func (client *Client) ListPbcInvokeReviewsWithContext(ctx context.Context, request *ListPbcInvokeReviewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPbcInvokeReviewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Applicant) {
		query["applicant"] = request.Applicant
	}

	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.Orderby) {
		query["orderby"] = request.Orderby
	}

	if !dara.IsNil(request.Reviewer) {
		query["reviewer"] = request.Reviewer
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPbcInvokeReviews"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-invoke-reviews"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPbcInvokeReviewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询我调用的pbc列表
//
// @param request - ListPbcInvokesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPbcInvokesResponse
func (client *Client) ListPbcInvokesWithContext(ctx context.Context, request *ListPbcInvokesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPbcInvokesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Applicant) {
		query["applicant"] = request.Applicant
	}

	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPbcInvokes"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-invokes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPbcInvokesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审核信息列表
//
// @param request - ListPbcReviewsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPbcReviewsResponse
func (client *Client) ListPbcReviewsWithContext(ctx context.Context, request *ListPbcReviewsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPbcReviewsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Applicant) {
		query["applicant"] = request.Applicant
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Reviewer) {
		query["reviewer"] = request.Reviewer
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPbcReviews"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-reviews"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPbcReviewsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询我订阅的pbc列表
//
// @param request - ListPbcSubscribeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPbcSubscribeResponse
func (client *Client) ListPbcSubscribeWithContext(ctx context.Context, request *ListPbcSubscribeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPbcSubscribeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderBy) {
		query["order_by"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["order_direction"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPbcSubscribe"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/commands/list-subscribe"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPbcSubscribeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询我创建的资产
//
// @param request - ListPbcVersionBuildRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPbcVersionBuildResponse
func (client *Client) ListPbcVersionBuildWithContext(ctx context.Context, request *ListPbcVersionBuildRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPbcVersionBuildResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["accountId"] = request.AccountId
	}

	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPbcVersionBuild"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/commands/list-build"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPbcVersionBuildResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询pbc的版本号列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPbcVersionNumbersResponse
func (client *Client) ListPbcVersionNumbersWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPbcVersionNumbersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPbcVersionNumbers"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/listPbcVersionNumbers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPbcVersionNumbersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询pbc列表
//
// @param request - ListPbcsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPbcsResponse
func (client *Client) ListPbcsWithContext(ctx context.Context, request *ListPbcsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPbcsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.DeveloperId) {
		query["developerId"] = request.DeveloperId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPbcs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPbcsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务配置列表
//
// @param request - ListPdpConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpConfigsResponse
func (client *Client) ListPdpConfigsWithContext(ctx context.Context, request *ListPdpConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpConfigsResponse, _err error) {
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

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpConfigs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/configs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务历史配置列表
//
// @param request - ListPdpHistoryConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpHistoryConfigsResponse
func (client *Client) ListPdpHistoryConfigsWithContext(ctx context.Context, request *ListPdpHistoryConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpHistoryConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["configId"] = request.ConfigId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpHistoryConfigs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/configs/history"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpHistoryConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务的镜像
//
// @param request - ListPdpImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpImageResponse
func (client *Client) ListPdpImageWithContext(ctx context.Context, request *ListPdpImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpImageResponse, _err error) {
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

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpImage"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/services/instance/images"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询泳道列表
//
// @param request - ListPdpLanesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpLanesResponse
func (client *Client) ListPdpLanesWithContext(ctx context.Context, request *ListPdpLanesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpLanesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.InletServiceId) {
		query["inletServiceId"] = request.InletServiceId
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpLanes"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpLanesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询微服务分组可以加入的泳道列表
//
// @param tmpReq - ListPdpLanesForServiceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpLanesForServiceGroupResponse
func (client *Client) ListPdpLanesForServiceGroupWithContext(ctx context.Context, tmpReq *ListPdpLanesForServiceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpLanesForServiceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPdpLanesForServiceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LaneIds) {
		request.LaneIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LaneIds, dara.String("laneIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LaneIdsShrink) {
		query["laneIds"] = request.LaneIdsShrink
	}

	if !dara.IsNil(request.Operator) {
		query["operator"] = request.Operator
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		query["serviceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpLanesForServiceGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/commands/list-service-group-lane"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpLanesForServiceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日志列表
//
// @param request - ListPdpLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpLogsResponse
func (client *Client) ListPdpLogsWithContext(ctx context.Context, request *ListPdpLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpLogsResponse, _err error) {
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

	if !dara.IsNil(request.Ip) {
		query["ip"] = request.Ip
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpLogs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/pdp-log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询/搜索PdpPbc列表
//
// @param tmpReq - ListPdpPbcsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpPbcsResponse
func (client *Client) ListPdpPbcsWithContext(ctx context.Context, tmpReq *ListPdpPbcsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpPbcsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPdpPbcsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PbcIds) {
		request.PbcIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PbcIds, dara.String("pbcIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.DeveloperId) {
		query["developerId"] = request.DeveloperId
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PbcIdsShrink) {
		query["pbcIds"] = request.PbcIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpPbcs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/pbcs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpPbcsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务分组列表
//
// @param tmpReq - ListPdpServiceGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpServiceGroupsResponse
func (client *Client) ListPdpServiceGroupsWithContext(ctx context.Context, tmpReq *ListPdpServiceGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpServiceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPdpServiceGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		query["alias"] = request.Alias
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.EnvType) {
		query["envType"] = request.EnvType
	}

	if !dara.IsNil(request.GroupType) {
		query["groupType"] = request.GroupType
	}

	if !dara.IsNil(request.IdsShrink) {
		query["ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.OrgType) {
		query["orgType"] = request.OrgType
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.RepoId) {
		query["repoId"] = request.RepoId
	}

	if !dara.IsNil(request.ServiceId) {
		query["serviceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpServiceGroups"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpServiceGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务列表
//
// @param tmpReq - ListPdpServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPdpServicesResponse
func (client *Client) ListPdpServicesWithContext(ctx context.Context, tmpReq *ListPdpServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPdpServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPdpServicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PdpServiceIds) {
		request.PdpServiceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PdpServiceIds, dara.String("pdpServiceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		query["alias"] = request.Alias
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	if !dara.IsNil(request.PdpServiceIdsShrink) {
		query["pdpServiceIds"] = request.PdpServiceIdsShrink
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPdpServices"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPdpServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取权限列表
//
// @param request - ListPermissionResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionResourceResponse
func (client *Client) ListPermissionResourceWithContext(ctx context.Context, request *ListPermissionResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPermissionResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		query["action"] = request.Action
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OperatorId) {
		query["operatorId"] = request.OperatorId
	}

	if !dara.IsNil(request.OperatorType) {
		query["operatorType"] = request.OperatorType
	}

	if !dara.IsNil(request.ResourcePrefix) {
		query["resourcePrefix"] = request.ResourcePrefix
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissionResource"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/permissions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取前端资源权限，如按钮、页面
//
// @param request - ListPermissionResourceForFrontRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionResourceForFrontResponse
func (client *Client) ListPermissionResourceForFrontWithContext(ctx context.Context, request *ListPermissionResourceForFrontRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPermissionResourceForFrontResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		query["action"] = request.Action
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OperatorId) {
		query["operatorId"] = request.OperatorId
	}

	if !dara.IsNil(request.OperatorType) {
		query["operatorType"] = request.OperatorType
	}

	if !dara.IsNil(request.QueryType) {
		query["queryType"] = request.QueryType
	}

	if !dara.IsNil(request.ResourcePrefix) {
		query["resourcePrefix"] = request.ResourcePrefix
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissionResourceForFront"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/permissions/command/front-permission"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionResourceForFrontResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取有权限的资源,请求来源POP
//
// @param request - ListPermissionResourcePopRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionResourcePopResponse
func (client *Client) ListPermissionResourcePopWithContext(ctx context.Context, request *ListPermissionResourcePopRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPermissionResourcePopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Action) {
		query["action"] = request.Action
	}

	if !dara.IsNil(request.OperatorId) {
		query["operatorId"] = request.OperatorId
	}

	if !dara.IsNil(request.OperatorType) {
		query["operatorType"] = request.OperatorType
	}

	if !dara.IsNil(request.ResourcePrefix) {
		query["resourcePrefix"] = request.ResourcePrefix
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissionResourcePop"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/permissions/pop"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionResourcePopResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取该角色下的权限
//
// @param request - ListPrivilegeByRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivilegeByRoleResponse
func (client *Client) ListPrivilegeByRoleWithContext(ctx context.Context, roleId *string, request *ListPrivilegeByRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPrivilegeByRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["accountId"] = request.AccountId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivilegeByRole"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/" + dara.PercentEncode(dara.StringValue(roleId)) + "/privileges"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivilegeByRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询产品环境
//
// @param request - ListProductEnvInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductEnvInfosResponse
func (client *Client) ListProductEnvInfosWithContext(ctx context.Context, request *ListProductEnvInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProductEnvInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProductEnvInfos"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/commands/list-product-env"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductEnvInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询产品环境
//
// @param request - ListProductEnvsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductEnvsResponse
func (client *Client) ListProductEnvsWithContext(ctx context.Context, request *ListProductEnvsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProductEnvsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProductEnvs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/resources/commands/list-product-env"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductEnvsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询/搜索产品列表
//
// @param request - ListProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithContext(ctx context.Context, request *ListProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/products"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源列表
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithContext(ctx context.Context, request *ListResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		query["productId"] = request.ProductId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResources"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有审核人信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReviewersResponse
func (client *Client) ListReviewersWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListReviewersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReviewers"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-reviews/commands/listReviewers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReviewersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询角色列表
//
// @param tmpReq - ListRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithContext(ctx context.Context, tmpReq *ListRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleIds) {
		request.RoleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleIds, dara.String("roleIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizerId) {
		query["authorizerId"] = request.AuthorizerId
	}

	if !dara.IsNil(request.AuthorizerType) {
		query["authorizerType"] = request.AuthorizerType
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Platform) {
		query["platform"] = request.Platform
	}

	if !dara.IsNil(request.RoleIdsShrink) {
		query["roleIds"] = request.RoleIdsShrink
	}

	if !dara.IsNil(request.RoleType) {
		query["roleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoles"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询token列表
//
// @param request - ListRuntimeTokensRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRuntimeTokensResponse
func (client *Client) ListRuntimeTokensWithContext(ctx context.Context, request *ListRuntimeTokensRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRuntimeTokensResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["enterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PbcId) {
		query["pbcId"] = request.PbcId
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRuntimeTokens"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/services/env/tokens"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRuntimeTokensResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务实例列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServerInstancesResponse
func (client *Client) ListServerInstancesWithContext(ctx context.Context, env *string, serviceGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServerInstancesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServerInstances"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/services/" + dara.PercentEncode(dara.StringValue(serviceGroupId)) + "/env/" + dara.PercentEncode(dara.StringValue(env))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServerInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询多个服务监控指标
//
// @param request - ListServiceMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceMetricsResponse
func (client *Client) ListServiceMetricsWithContext(ctx context.Context, request *ListServiceMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceMetricsResponse, _err error) {
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

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.GroupId) {
		query["groupId"] = request.GroupId
	}

	if !dara.IsNil(request.IntervalInSec) {
		query["intervalInSec"] = request.IntervalInSec
	}

	if !dara.IsNil(request.Ip) {
		query["ip"] = request.Ip
	}

	if !dara.IsNil(request.Measures) {
		query["measures"] = request.Measures
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.TopicId) {
		query["topicId"] = request.TopicId
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceMetrics"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询申请记录列表
//
// @param request - ListSettledsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSettledsResponse
func (client *Client) ListSettledsWithContext(ctx context.Context, request *ListSettledsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSettledsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["accountId"] = request.AccountId
	}

	if !dara.IsNil(request.Applicant) {
		query["applicant"] = request.Applicant
	}

	if !dara.IsNil(request.EnterpriseName) {
		query["enterpriseName"] = request.EnterpriseName
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSettleds"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/settleds"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSettledsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询订阅当前组件的pbc的列表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubscribePbcsResponse
func (client *Client) ListSubscribePbcsWithContext(ctx context.Context, pbcName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSubscribePbcsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubscribePbcs"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/" + dara.PercentEncode(dara.StringValue(pbcName)) + "/commands/subscribe-list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubscribePbcsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 依赖此PBC的列表(上游PBC)
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUpstreamPbcResponse
func (client *Client) ListUpstreamPbcWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUpstreamPbcResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUpstreamPbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-/versions/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/list-upstream"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUpstreamPbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询我关注的资产列表
//
// @param request - ListWatchAssetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWatchAssetsResponse
func (client *Client) ListWatchAssetsWithContext(ctx context.Context, request *ListWatchAssetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWatchAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["accountId"] = request.AccountId
	}

	if !dara.IsNil(request.AssetType) {
		query["assetType"] = request.AssetType
	}

	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UpshelfAssetId) {
		query["upshelfAssetId"] = request.UpshelfAssetId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWatchAssets"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/asset-watchs/commands/list-watch"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWatchAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成跳转mq控制台链接
//
// @param request - ObtainMqConsoleLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainMqConsoleLinkResponse
func (client *Client) ObtainMqConsoleLinkWithContext(ctx context.Context, request *ObtainMqConsoleLinkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ObtainMqConsoleLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ObtainMqConsoleLink"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/mq/topic/commonds/obtainMqConsoleLink"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ObtainMqConsoleLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 已有微服务开启分组
//
// @param request - OpenServiceGroupForServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenServiceGroupForServiceResponse
func (client *Client) OpenServiceGroupForServiceWithContext(ctx context.Context, request *OpenServiceGroupForServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenServiceGroupForServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenServiceGroupForService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/commands/open-group"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenServiceGroupForServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预览组件CRD
//
// @param request - PreviewComponentCrdSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewComponentCrdSchemaResponse
func (client *Client) PreviewComponentCrdSchemaWithContext(ctx context.Context, request *PreviewComponentCrdSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PreviewComponentCrdSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewComponentCrdSchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/components/commonds/preview-component-schema"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewComponentCrdSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册Buc用户
//
// @param request - RegisterBucUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterBucUserResponse
func (client *Client) RegisterBucUserWithContext(ctx context.Context, request *RegisterBucUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RegisterBucUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterBucUser"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/bucs/logins/register"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterBucUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消关注资产
//
// @param request - RemoveAssetWatchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAssetWatchResponse
func (client *Client) RemoveAssetWatchWithContext(ctx context.Context, assetId *string, request *RemoveAssetWatchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveAssetWatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetType) {
		query["assetType"] = request.AssetType
	}

	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAssetWatch"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/asset-watchs/" + dara.PercentEncode(dara.StringValue(assetId)) + "/commands/remove-watch"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAssetWatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消依赖二方包
//
// @param request - RemoveDependLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDependLibraryResponse
func (client *Client) RemoveDependLibraryWithContext(ctx context.Context, id *string, request *RemoveDependLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveDependLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDependLibrary"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/remove-dependency"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDependLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用指定版本部署
//
// @param request - RevertPdpServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevertPdpServiceResponse
func (client *Client) RevertPdpServiceWithContext(ctx context.Context, request *RevertPdpServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevertPdpServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevertPdpService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/deployments/commands/revert"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevertPdpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销二方包上架审核
//
// @param request - RevokeLibraryReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeLibraryReviewResponse
func (client *Client) RevokeLibraryReviewWithContext(ctx context.Context, request *RevokeLibraryReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeLibraryReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeLibraryReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/library-reviews/commands/revoke-review"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeLibraryReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销Pbc上架审核
//
// @param request - RevokePbcReviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokePbcReviewResponse
func (client *Client) RevokePbcReviewWithContext(ctx context.Context, request *RevokePbcReviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokePbcReviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokePbcReview"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-reviews/commands/revoke-review"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokePbcReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消授权
//
// @param request - RevokeRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeRoleResponse
func (client *Client) RevokeRoleWithContext(ctx context.Context, roleId *string, request *RevokeRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeRole"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/" + dara.PercentEncode(dara.StringValue(roleId)) + "/commands/revoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &RevokeRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回滚
//
// @param request - RollbackPdpServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackPdpServiceResponse
func (client *Client) RollbackPdpServiceWithContext(ctx context.Context, request *RollbackPdpServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RollbackPdpServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackPdpService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/deployments/commands/rollback"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackPdpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询/搜索资产列表信息
//
// @param tmpReq - SearchAssetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAssetsResponse
func (client *Client) SearchAssetsWithContext(ctx context.Context, tmpReq *SearchAssetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetIndustrys) {
		request.AssetIndustrysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetIndustrys, dara.String("assetIndustrys"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AssetTypes) {
		request.AssetTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetTypes, dara.String("assetTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetIndustrysShrink) {
		query["assetIndustrys"] = request.AssetIndustrysShrink
	}

	if !dara.IsNil(request.AssetName) {
		query["assetName"] = request.AssetName
	}

	if !dara.IsNil(request.AssetTypesShrink) {
		query["assetTypes"] = request.AssetTypesShrink
	}

	if !dara.IsNil(request.CompanyId) {
		query["companyId"] = request.CompanyId
	}

	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchAssets"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/markets/commands/search-asset"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询/搜索pbc资产列表信息
//
// @param request - SearchPbcAssetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchPbcAssetsResponse
func (client *Client) SearchPbcAssetsWithContext(ctx context.Context, request *SearchPbcAssetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchPbcAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Industry) {
		query["industry"] = request.Industry
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.OrderBy) {
		query["order_by"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["order_direction"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["page_number"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["page_size"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchPbcAssets"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/markets/commands/search"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchPbcAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询调用链列表信息
//
// @param request - SearchTracesByPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTracesByPageResponse
func (client *Client) SearchTracesByPageWithContext(ctx context.Context, request *SearchTracesByPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchTracesByPageResponse, _err error) {
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

	if !dara.IsNil(request.Env) {
		query["env"] = request.Env
	}

	if !dara.IsNil(request.MinDuration) {
		query["minDuration"] = request.MinDuration
	}

	if !dara.IsNil(request.OperationName) {
		query["operationName"] = request.OperationName
	}

	if !dara.IsNil(request.OrderBy) {
		query["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		query["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceGroupId) {
		query["serviceGroupId"] = request.ServiceGroupId
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchTracesByPage"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/traces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchTracesByPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ARMS告警联系人发送手机号码验证
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTTSVerifyLinkResponse
func (client *Client) SendTTSVerifyLinkWithContext(ctx context.Context, contactId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendTTSVerifyLinkResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendTTSVerifyLink"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/contact/commands/phoneVerify/" + dara.PercentEncode(dara.StringValue(contactId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SendTTSVerifyLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订阅pbc
//
// @param request - SubscribePbcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribePbcResponse
func (client *Client) SubscribePbcWithContext(ctx context.Context, pbcName *string, request *SubscribePbcRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubscribePbcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubscribePbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbcs/" + dara.PercentEncode(dara.StringValue(pbcName)) + "/commands/subscribe"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubscribePbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步组件的模板配置
//
// @param request - SyncComponentTemplateConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncComponentTemplateConfigResponse
func (client *Client) SyncComponentTemplateConfigWithContext(ctx context.Context, request *SyncComponentTemplateConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncComponentTemplateConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncComponentTemplateConfig"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/components/commonds/sync-template-config"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncComponentTemplateConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转移公司
//
// @param request - TransferEnterpriseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferEnterpriseResponse
func (client *Client) TransferEnterpriseWithContext(ctx context.Context, enterpriseId *string, request *TransferEnterpriseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TransferEnterpriseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferEnterprise"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/enterprises/" + dara.PercentEncode(dara.StringValue(enterpriseId)) + "/commands/transfer"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferEnterpriseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发部署
//
// @param request - TriggerDeploymentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerDeploymentResponse
func (client *Client) TriggerDeploymentWithContext(ctx context.Context, request *TriggerDeploymentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TriggerDeploymentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerDeployment"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/deployments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerDeploymentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上架二方包
//
// @param request - UpShelfLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpShelfLibraryResponse
func (client *Client) UpShelfLibraryWithContext(ctx context.Context, id *string, request *UpShelfLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpShelfLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpShelfLibrary"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/up-shelf-library"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpShelfLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上架pbc版本
//
// @param request - UpShelfPbcVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpShelfPbcVersionResponse
func (client *Client) UpShelfPbcVersionWithContext(ctx context.Context, id *string, request *UpShelfPbcVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpShelfPbcVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MarketId) {
		query["marketId"] = request.MarketId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpShelfPbcVersion"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(id)) + "/commands/upShelf"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpShelfPbcVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新组件
//
// @param request - UpdateComponentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComponentResponse
func (client *Client) UpdateComponentWithContext(ctx context.Context, id *string, request *UpdateComponentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComponent"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/components/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新组件模板
//
// @param request - UpdateComponentTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComponentTemplateResponse
func (client *Client) UpdateComponentTemplateWithContext(ctx context.Context, id *string, request *UpdateComponentTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateComponentTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComponentTemplate"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/component-templates/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComponentTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新开发者信息
//
// @param request - UpdateDeveloperRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeveloperResponse
func (client *Client) UpdateDeveloperWithContext(ctx context.Context, accountId *string, request *UpdateDeveloperRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDeveloperResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeveloper"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/developers/" + dara.PercentEncode(dara.StringValue(accountId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeveloperResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新公司信息
//
// @param request - UpdateEnterpriseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnterpriseResponse
func (client *Client) UpdateEnterpriseWithContext(ctx context.Context, enterpriseId *string, request *UpdateEnterpriseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEnterpriseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnterprise"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/enterprises/" + dara.PercentEncode(dara.StringValue(enterpriseId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnterpriseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改二方包信息
//
// @param request - UpdateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLibraryResponse
func (client *Client) UpdateLibraryWithContext(ctx context.Context, request *UpdateLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLibrary"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新二方包使用说明文档
//
// @param request - UpdateLibraryInstructionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLibraryInstructionResponse
func (client *Client) UpdateLibraryInstructionWithContext(ctx context.Context, libraryId *string, request *UpdateLibraryInstructionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLibraryInstructionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLibraryInstruction"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/instructions"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLibraryInstructionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新二方包规格
//
// @param request - UpdateLibrarySchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLibrarySchemaResponse
func (client *Client) UpdateLibrarySchemaWithContext(ctx context.Context, libraryId *string, request *UpdateLibrarySchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLibrarySchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLibrarySchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/librarys/" + dara.PercentEncode(dara.StringValue(libraryId)) + "/schemas"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLibrarySchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑应用报警任务
//
// @param request - UpdateMonitorArmsAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMonitorArmsAlertResponse
func (client *Client) UpdateMonitorArmsAlertWithContext(ctx context.Context, request *UpdateMonitorArmsAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMonitorArmsAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMonitorArmsAlert"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/updateMonitorArmsAlert"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMonitorArmsAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新联系人
//
// @param request - UpdateMonitorContactRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMonitorContactResponse
func (client *Client) UpdateMonitorContactWithContext(ctx context.Context, request *UpdateMonitorContactRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMonitorContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMonitorContact"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/contact"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMonitorContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新联系人组
//
// @param request - UpdateMonitorContactGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMonitorContactGroupResponse
func (client *Client) UpdateMonitorContactGroupWithContext(ctx context.Context, request *UpdateMonitorContactGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMonitorContactGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMonitorContactGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/group"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMonitorContactGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑MQ报警任务
//
// @param request - UpdateMonitorMqAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMonitorMqAlertResponse
func (client *Client) UpdateMonitorMqAlertWithContext(ctx context.Context, request *UpdateMonitorMqAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMonitorMqAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMonitorMqAlert"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/updateMonitorMqAlert"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMonitorMqAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑日志报警任务
//
// @param request - UpdateMonitorSlsAlertRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMonitorSlsAlertResponse
func (client *Client) UpdateMonitorSlsAlertWithContext(ctx context.Context, request *UpdateMonitorSlsAlertRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMonitorSlsAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMonitorSlsAlert"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/alert/commands/updateMonitorSlsAlert"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMonitorSlsAlertResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新webhook
//
// @param request - UpdateMonitorWebhookRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMonitorWebhookResponse
func (client *Client) UpdateMonitorWebhookWithContext(ctx context.Context, request *UpdateMonitorWebhookRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMonitorWebhookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMonitorWebhook"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-monitor/v1/monitor/webhook"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMonitorWebhookResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新PBC的api规格
//
// @param request - UpdatePbcApiSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePbcApiSchemaResponse
func (client *Client) UpdatePbcApiSchemaWithContext(ctx context.Context, pbcVersionId *string, request *UpdatePbcApiSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePbcApiSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePbcApiSchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(pbcVersionId)) + "/api-schemas"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePbcApiSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新PBC使用说明文档
//
// @param request - UpdatePbcInstructionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePbcInstructionResponse
func (client *Client) UpdatePbcInstructionWithContext(ctx context.Context, pbcVersionId *string, request *UpdatePbcInstructionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePbcInstructionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePbcInstruction"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(pbcVersionId)) + "/instructions"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePbcInstructionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新PBC规格
//
// @param request - UpdatePbcSchemaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePbcSchemaResponse
func (client *Client) UpdatePbcSchemaWithContext(ctx context.Context, pbcVersionId *string, request *UpdatePbcSchemaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePbcSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePbcSchema"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(pbcVersionId)) + "/schemas"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePbcSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新PBC版本
//
// @param request - UpdatePbcVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePbcVersionResponse
func (client *Client) UpdatePbcVersionWithContext(ctx context.Context, id *string, request *UpdatePbcVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePbcVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePbcVersion"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/catalog/v1/pbc-versions/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePbcVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务配置信息
//
// @param request - UpdatePdpConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePdpConfigResponse
func (client *Client) UpdatePdpConfigWithContext(ctx context.Context, request *UpdatePdpConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePdpConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePdpConfig"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/configs"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePdpConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新泳道
//
// @param request - UpdatePdpLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePdpLaneResponse
func (client *Client) UpdatePdpLaneWithContext(ctx context.Context, id *string, request *UpdatePdpLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePdpLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePdpLane"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/lanes/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePdpLaneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新PdpPbc
//
// @param request - UpdatePdpPbcRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePdpPbcResponse
func (client *Client) UpdatePdpPbcWithContext(ctx context.Context, request *UpdatePdpPbcRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePdpPbcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePdpPbc"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/pbcs"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePdpPbcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务信息
//
// @param request - UpdatePdpServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePdpServiceResponse
func (client *Client) UpdatePdpServiceWithContext(ctx context.Context, request *UpdatePdpServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePdpServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePdpService"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/services"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePdpServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务分组信息
//
// @param request - UpdatePdpServiceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePdpServiceGroupResponse
func (client *Client) UpdatePdpServiceGroupWithContext(ctx context.Context, request *UpdatePdpServiceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePdpServiceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePdpServiceGroup"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-service/v1/groups"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePdpServiceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新产品
//
// @param request - UpdateProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProductResponse
func (client *Client) UpdateProductWithContext(ctx context.Context, request *UpdateProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProduct"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-pbc/v1/products"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源
//
// @param request - UpdateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithContext(ctx context.Context, id *string, request *UpdateResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResource"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pdp-resource/v1/resources/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// 更新角色信息
//
// @param request - UpdateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithContext(ctx context.Context, roleId *string, request *UpdateRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRole"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/roles/role-id/" + dara.PercentEncode(dara.StringValue(roleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新申请记录
//
// @param request - UpdateSettledRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSettledResponse
func (client *Client) UpdateSettledWithContext(ctx context.Context, id *string, request *UpdateSettledRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSettledResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSettled"),
		Version:     dara.String("2021-11-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/manager/v1/settleds/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSettledResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
