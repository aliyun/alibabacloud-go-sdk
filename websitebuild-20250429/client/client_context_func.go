// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 绑定应用域名
//
// @param request - BindAppDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAppDomainResponse
func (client *Client) BindAppDomainWithContext(ctx context.Context, request *BindAppDomainRequest, runtime *dara.RuntimeOptions) (_result *BindAppDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAppDomain"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAppDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交创建Logo任务
//
// @param request - CreateLogoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogoTaskResponse
func (client *Client) CreateLogoTaskWithContext(ctx context.Context, request *CreateLogoTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateLogoTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogoVersion) {
		query["LogoVersion"] = request.LogoVersion
	}

	if !dara.IsNil(request.NegativePrompt) {
		query["NegativePrompt"] = request.NegativePrompt
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogoTask"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除域名的SSL证书
//
// @param request - DeleteAppDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppDomainCertificateResponse
func (client *Client) DeleteAppDomainCertificateWithContext(ctx context.Context, request *DeleteAppDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppDomainCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppDomainCertificate"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除域名的跳转规则
//
// @param request - DeleteAppDomainRedirectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppDomainRedirectResponse
func (client *Client) DeleteAppDomainRedirectWithContext(ctx context.Context, request *DeleteAppDomainRedirectRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppDomainRedirectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppDomainRedirect"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppDomainRedirectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名的DNS解析记录
//
// @param request - DescribeAppDomainDnsRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppDomainDnsRecordResponse
func (client *Client) DescribeAppDomainDnsRecordWithContext(ctx context.Context, request *DescribeAppDomainDnsRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppDomainDnsRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Purpose) {
		query["Purpose"] = request.Purpose
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppDomainDnsRecord"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppDomainDnsRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Logo创建任务
//
// @param request - GetCreateLogoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateLogoTaskResponse
func (client *Client) GetCreateLogoTaskWithContext(ctx context.Context, request *GetCreateLogoTaskRequest, runtime *dara.RuntimeOptions) (_result *GetCreateLogoTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCreateLogoTask"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCreateLogoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提供给服务商的域名查询接口
//
// @param request - GetDomainInfoForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainInfoForPartnerResponse
func (client *Client) GetDomainInfoForPartnerWithContext(ctx context.Context, request *GetDomainInfoForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetDomainInfoForPartnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomainInfoForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainInfoForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名备案信息
//
// @param request - GetIcpFilingInfoForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIcpFilingInfoForPartnerResponse
func (client *Client) GetIcpFilingInfoForPartnerWithContext(ctx context.Context, request *GetIcpFilingInfoForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetIcpFilingInfoForPartnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIcpFilingInfoForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIcpFilingInfoForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴获取用户SLR角色授权临时凭证
//
// @param request - GetUserTmpIdentityForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserTmpIdentityForPartnerResponse
func (client *Client) GetUserTmpIdentityForPartnerWithContext(ctx context.Context, request *GetUserTmpIdentityForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetUserTmpIdentityForPartnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthPurpose) {
		query["AuthPurpose"] = request.AuthPurpose
	}

	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.ServiceLinkedRole) {
		query["ServiceLinkedRole"] = request.ServiceLinkedRole
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserTmpIdentityForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserTmpIdentityForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名的跳转规则列表
//
// @param request - ListAppDomainRedirectRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppDomainRedirectRecordsResponse
func (client *Client) ListAppDomainRedirectRecordsWithContext(ctx context.Context, request *ListAppDomainRedirectRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListAppDomainRedirectRecordsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppDomainRedirectRecords"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppDomainRedirectRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用实例下的所有域名列表
//
// @param request - ListAppInstanceDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInstanceDomainsResponse
func (client *Client) ListAppInstanceDomainsWithContext(ctx context.Context, request *ListAppInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListAppInstanceDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppInstanceDomains"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppInstanceDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用
//
// @param request - OperateAppInstanceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppInstanceForPartnerResponse
func (client *Client) OperateAppInstanceForPartnerWithContext(ctx context.Context, request *OperateAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateAppInstanceForPartnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.OperateEvent) {
		query["OperateEvent"] = request.OperateEvent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAppInstanceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAppInstanceForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用服务
//
// @param request - OperateAppServiceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppServiceForPartnerResponse
func (client *Client) OperateAppServiceForPartnerWithContext(ctx context.Context, request *OperateAppServiceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateAppServiceForPartnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.OperateEvent) {
		query["OperateEvent"] = request.OperateEvent
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAppServiceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAppServiceForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片检索
//
// @param tmpReq - SearchImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageResponse
func (client *Client) SearchImageWithContext(ctx context.Context, tmpReq *SearchImageRequest, runtime *dara.RuntimeOptions) (_result *SearchImageResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SearchImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ColorHex) {
		query["ColorHex"] = request.ColorHex
	}

	if !dara.IsNil(request.HasPerson) {
		query["HasPerson"] = request.HasPerson
	}

	if !dara.IsNil(request.ImageCategory) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !dara.IsNil(request.ImageRatio) {
		query["ImageRatio"] = request.ImageRatio
	}

	if !dara.IsNil(request.MaxHeight) {
		query["MaxHeight"] = request.MaxHeight
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MaxWidth) {
		query["MaxWidth"] = request.MaxWidth
	}

	if !dara.IsNil(request.MinHeight) {
		query["MinHeight"] = request.MinHeight
	}

	if !dara.IsNil(request.MinWidth) {
		query["MinWidth"] = request.MinWidth
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchImage"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置域名的SSL证书
//
// @param request - SetAppDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAppDomainCertificateResponse
func (client *Client) SetAppDomainCertificateWithContext(ctx context.Context, request *SetAppDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetAppDomainCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.CertificateName) {
		query["CertificateName"] = request.CertificateName
	}

	if !dara.IsNil(request.CertificateType) {
		query["CertificateType"] = request.CertificateType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.PublicKey) {
		query["PublicKey"] = request.PublicKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAppDomainCertificate"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAppDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴同步应用实例
//
// @param tmpReq - SyncAppInstanceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncAppInstanceForPartnerResponse
func (client *Client) SyncAppInstanceForPartnerWithContext(ctx context.Context, tmpReq *SyncAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *SyncAppInstanceForPartnerResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SyncAppInstanceForPartnerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppInstance) {
		request.AppInstanceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppInstance, dara.String("AppInstance"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceShrink) {
		query["AppInstance"] = request.AppInstanceShrink
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Operator) {
		query["Operator"] = request.Operator
	}

	if !dara.IsNil(request.SourceBizId) {
		query["SourceBizId"] = request.SourceBizId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncAppInstanceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncAppInstanceForPartnerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑应用域名
//
// @param request - UnbindAppDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAppDomainResponse
func (client *Client) UnbindAppDomainWithContext(ctx context.Context, request *UnbindAppDomainRequest, runtime *dara.RuntimeOptions) (_result *UnbindAppDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindAppDomain"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAppDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
