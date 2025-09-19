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
	client.Endpoint, _err = client.GetEndpoint(dara.String("websitebuild"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 绑定应用域名
//
// @param request - BindAppDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAppDomainResponse
func (client *Client) BindAppDomainWithOptions(request *BindAppDomainRequest, runtime *dara.RuntimeOptions) (_result *BindAppDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定应用域名
//
// @param request - BindAppDomainRequest
//
// @return BindAppDomainResponse
func (client *Client) BindAppDomain(request *BindAppDomainRequest) (_result *BindAppDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAppDomainResponse{}
	_body, _err := client.BindAppDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateLogoTaskWithOptions(request *CreateLogoTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateLogoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateLogoTaskResponse
func (client *Client) CreateLogoTask(request *CreateLogoTaskRequest) (_result *CreateLogoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLogoTaskResponse{}
	_body, _err := client.CreateLogoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAppDomainCertificateWithOptions(request *DeleteAppDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppDomainCertificateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAppDomainCertificateResponse
func (client *Client) DeleteAppDomainCertificate(request *DeleteAppDomainCertificateRequest) (_result *DeleteAppDomainCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppDomainCertificateResponse{}
	_body, _err := client.DeleteAppDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAppDomainRedirectWithOptions(request *DeleteAppDomainRedirectRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppDomainRedirectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAppDomainRedirectResponse
func (client *Client) DeleteAppDomainRedirect(request *DeleteAppDomainRedirectRequest) (_result *DeleteAppDomainRedirectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppDomainRedirectResponse{}
	_body, _err := client.DeleteAppDomainRedirectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeAppDomainDnsRecordWithOptions(request *DescribeAppDomainDnsRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppDomainDnsRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeAppDomainDnsRecordResponse
func (client *Client) DescribeAppDomainDnsRecord(request *DescribeAppDomainDnsRecordRequest) (_result *DescribeAppDomainDnsRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppDomainDnsRecordResponse{}
	_body, _err := client.DescribeAppDomainDnsRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCreateLogoTaskWithOptions(request *GetCreateLogoTaskRequest, runtime *dara.RuntimeOptions) (_result *GetCreateLogoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCreateLogoTaskResponse
func (client *Client) GetCreateLogoTask(request *GetCreateLogoTaskRequest) (_result *GetCreateLogoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCreateLogoTaskResponse{}
	_body, _err := client.GetCreateLogoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDomainInfoForPartnerWithOptions(request *GetDomainInfoForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetDomainInfoForPartnerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDomainInfoForPartnerResponse
func (client *Client) GetDomainInfoForPartner(request *GetDomainInfoForPartnerRequest) (_result *GetDomainInfoForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDomainInfoForPartnerResponse{}
	_body, _err := client.GetDomainInfoForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetIcpFilingInfoForPartnerWithOptions(request *GetIcpFilingInfoForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetIcpFilingInfoForPartnerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetIcpFilingInfoForPartnerResponse
func (client *Client) GetIcpFilingInfoForPartner(request *GetIcpFilingInfoForPartnerRequest) (_result *GetIcpFilingInfoForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIcpFilingInfoForPartnerResponse{}
	_body, _err := client.GetIcpFilingInfoForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUserTmpIdentityForPartnerWithOptions(request *GetUserTmpIdentityForPartnerRequest, runtime *dara.RuntimeOptions) (_result *GetUserTmpIdentityForPartnerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetUserTmpIdentityForPartnerResponse
func (client *Client) GetUserTmpIdentityForPartner(request *GetUserTmpIdentityForPartnerRequest) (_result *GetUserTmpIdentityForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserTmpIdentityForPartnerResponse{}
	_body, _err := client.GetUserTmpIdentityForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAppDomainRedirectRecordsWithOptions(request *ListAppDomainRedirectRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListAppDomainRedirectRecordsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAppDomainRedirectRecordsResponse
func (client *Client) ListAppDomainRedirectRecords(request *ListAppDomainRedirectRecordsRequest) (_result *ListAppDomainRedirectRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppDomainRedirectRecordsResponse{}
	_body, _err := client.ListAppDomainRedirectRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAppInstanceDomainsWithOptions(request *ListAppInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListAppInstanceDomainsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAppInstanceDomainsResponse
func (client *Client) ListAppInstanceDomains(request *ListAppInstanceDomainsRequest) (_result *ListAppInstanceDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppInstanceDomainsResponse{}
	_body, _err := client.ListAppInstanceDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateAppInstanceForPartnerWithOptions(request *OperateAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateAppInstanceForPartnerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateAppInstanceForPartnerResponse
func (client *Client) OperateAppInstanceForPartner(request *OperateAppInstanceForPartnerRequest) (_result *OperateAppInstanceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateAppInstanceForPartnerResponse{}
	_body, _err := client.OperateAppInstanceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateAppServiceForPartnerWithOptions(request *OperateAppServiceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateAppServiceForPartnerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateAppServiceForPartnerResponse
func (client *Client) OperateAppServiceForPartner(request *OperateAppServiceForPartnerRequest) (_result *OperateAppServiceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateAppServiceForPartnerResponse{}
	_body, _err := client.OperateAppServiceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchImageWithOptions(tmpReq *SearchImageRequest, runtime *dara.RuntimeOptions) (_result *SearchImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SearchImageRequest
//
// @return SearchImageResponse
func (client *Client) SearchImage(request *SearchImageRequest) (_result *SearchImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchImageResponse{}
	_body, _err := client.SearchImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetAppDomainCertificateWithOptions(request *SetAppDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetAppDomainCertificateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetAppDomainCertificateResponse
func (client *Client) SetAppDomainCertificate(request *SetAppDomainCertificateRequest) (_result *SetAppDomainCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAppDomainCertificateResponse{}
	_body, _err := client.SetAppDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SyncAppInstanceForPartnerWithOptions(tmpReq *SyncAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *SyncAppInstanceForPartnerResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - SyncAppInstanceForPartnerRequest
//
// @return SyncAppInstanceForPartnerResponse
func (client *Client) SyncAppInstanceForPartner(request *SyncAppInstanceForPartnerRequest) (_result *SyncAppInstanceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncAppInstanceForPartnerResponse{}
	_body, _err := client.SyncAppInstanceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UnbindAppDomainWithOptions(request *UnbindAppDomainRequest, runtime *dara.RuntimeOptions) (_result *UnbindAppDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UnbindAppDomainResponse
func (client *Client) UnbindAppDomain(request *UnbindAppDomainRequest) (_result *UnbindAppDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindAppDomainResponse{}
	_body, _err := client.UnbindAppDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
