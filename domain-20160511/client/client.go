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
	EnableValidate  *bool
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
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("domain"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 域名检查
//
// @param request - CheckDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDomainResponse
func (client *Client) CheckDomainWithOptions(request *CheckDomainRequest, runtime *dara.RuntimeOptions) (_result *CheckDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDomain"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 域名检查
//
// @param request - CheckDomainRequest
//
// @return CheckDomainResponse
func (client *Client) CheckDomain(request *CheckDomainRequest) (_result *CheckDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDomainResponse{}
	_body, _err := client.CheckDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除联系人模板
//
// @param request - DeleteContactTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactTemplateResponse
func (client *Client) DeleteContactTemplateWithOptions(request *DeleteContactTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactTemplate"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除联系人模板
//
// @param request - DeleteContactTemplateRequest
//
// @return DeleteContactTemplateResponse
func (client *Client) DeleteContactTemplate(request *DeleteContactTemplateRequest) (_result *DeleteContactTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContactTemplateResponse{}
	_body, _err := client.DeleteContactTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务详情列表
//
// @param request - QueryBatchTaskDetailListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBatchTaskDetailListResponse
func (client *Client) QueryBatchTaskDetailListWithOptions(request *QueryBatchTaskDetailListRequest, runtime *dara.RuntimeOptions) (_result *QueryBatchTaskDetailListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SaleId) {
		query["SaleId"] = request.SaleId
	}

	if !dara.IsNil(request.TaskNo) {
		query["TaskNo"] = request.TaskNo
	}

	if !dara.IsNil(request.TaskStatus) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBatchTaskDetailList"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBatchTaskDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务详情列表
//
// @param request - QueryBatchTaskDetailListRequest
//
// @return QueryBatchTaskDetailListResponse
func (client *Client) QueryBatchTaskDetailList(request *QueryBatchTaskDetailListRequest) (_result *QueryBatchTaskDetailListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBatchTaskDetailListResponse{}
	_body, _err := client.QueryBatchTaskDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param request - QueryBatchTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBatchTaskListResponse
func (client *Client) QueryBatchTaskListWithOptions(request *QueryBatchTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryBatchTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginCreateTime) {
		query["BeginCreateTime"] = request.BeginCreateTime
	}

	if !dara.IsNil(request.EndCreateTime) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBatchTaskList"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBatchTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param request - QueryBatchTaskListRequest
//
// @return QueryBatchTaskListResponse
func (client *Client) QueryBatchTaskList(request *QueryBatchTaskListRequest) (_result *QueryBatchTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBatchTaskListResponse{}
	_body, _err := client.QueryBatchTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询联系人
//
// @param request - QueryContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContactResponse
func (client *Client) QueryContactWithOptions(request *QueryContactRequest, runtime *dara.RuntimeOptions) (_result *QueryContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactType) {
		query["ContactType"] = request.ContactType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryContact"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QueryContactRequest
//
// @return QueryContactResponse
func (client *Client) QueryContact(request *QueryContactRequest) (_result *QueryContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryContactResponse{}
	_body, _err := client.QueryContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询联系人模板
//
// @param request - QueryContactTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContactTemplateResponse
func (client *Client) QueryContactTemplateWithOptions(request *QueryContactTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryContactTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.CCompany) {
		query["CCompany"] = request.CCompany
	}

	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.DefaultTemplate) {
		query["DefaultTemplate"] = request.DefaultTemplate
	}

	if !dara.IsNil(request.ECompany) {
		query["ECompany"] = request.ECompany
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegType) {
		query["RegType"] = request.RegType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryContactTemplate"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryContactTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人模板
//
// @param request - QueryContactTemplateRequest
//
// @return QueryContactTemplateResponse
func (client *Client) QueryContactTemplate(request *QueryContactTemplateRequest) (_result *QueryContactTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryContactTemplateResponse{}
	_body, _err := client.QueryContactTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据saleId查询域名信息
//
// @param request - QueryDomainBySaleIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainBySaleIdResponse
func (client *Client) QueryDomainBySaleIdWithOptions(request *QueryDomainBySaleIdRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainBySaleIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SaleId) {
		query["SaleId"] = request.SaleId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainBySaleId"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainBySaleIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据saleId查询域名信息
//
// @param request - QueryDomainBySaleIdRequest
//
// @return QueryDomainBySaleIdResponse
func (client *Client) QueryDomainBySaleId(request *QueryDomainBySaleIdRequest) (_result *QueryDomainBySaleIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDomainBySaleIdResponse{}
	_body, _err := client.QueryDomainBySaleIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询域名列表
//
// @param request - QueryDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainListResponse
func (client *Client) QueryDomainListWithOptions(request *QueryDomainListRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeadEndDate) {
		query["DeadEndDate"] = request.DeadEndDate
	}

	if !dara.IsNil(request.DeadStartDate) {
		query["DeadStartDate"] = request.DeadStartDate
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderByType) {
		query["OrderByType"] = request.OrderByType
	}

	if !dara.IsNil(request.OrderKeyType) {
		query["OrderKeyType"] = request.OrderKeyType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductDomainType) {
		query["ProductDomainType"] = request.ProductDomainType
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.RegEndDate) {
		query["RegEndDate"] = request.RegEndDate
	}

	if !dara.IsNil(request.RegStartDate) {
		query["RegStartDate"] = request.RegStartDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainList"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名列表
//
// @param request - QueryDomainListRequest
//
// @return QueryDomainListResponse
func (client *Client) QueryDomainList(request *QueryDomainListRequest) (_result *QueryDomainListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDomainListResponse{}
	_body, _err := client.QueryDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询失败原因列表
//
// @param request - QueryFailReasonListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFailReasonListResponse
func (client *Client) QueryFailReasonListWithOptions(request *QueryFailReasonListRequest, runtime *dara.RuntimeOptions) (_result *QueryFailReasonListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SaleId) {
		query["SaleId"] = request.SaleId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFailReasonList"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFailReasonListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询失败原因列表
//
// @param request - QueryFailReasonListRequest
//
// @return QueryFailReasonListResponse
func (client *Client) QueryFailReasonList(request *QueryFailReasonListRequest) (_result *QueryFailReasonListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFailReasonListResponse{}
	_body, _err := client.QueryFailReasonListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存联系人模板
//
// @param request - SaveContactTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveContactTemplateResponse
func (client *Client) SaveContactTemplateWithOptions(request *SaveContactTemplateRequest, runtime *dara.RuntimeOptions) (_result *SaveContactTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CCity) {
		query["CCity"] = request.CCity
	}

	if !dara.IsNil(request.CCompany) {
		query["CCompany"] = request.CCompany
	}

	if !dara.IsNil(request.CCountry) {
		query["CCountry"] = request.CCountry
	}

	if !dara.IsNil(request.CName) {
		query["CName"] = request.CName
	}

	if !dara.IsNil(request.CProvince) {
		query["CProvince"] = request.CProvince
	}

	if !dara.IsNil(request.CVenu) {
		query["CVenu"] = request.CVenu
	}

	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.DefaultTemplate) {
		query["DefaultTemplate"] = request.DefaultTemplate
	}

	if !dara.IsNil(request.ECity) {
		query["ECity"] = request.ECity
	}

	if !dara.IsNil(request.ECompany) {
		query["ECompany"] = request.ECompany
	}

	if !dara.IsNil(request.EName) {
		query["EName"] = request.EName
	}

	if !dara.IsNil(request.EProvince) {
		query["EProvince"] = request.EProvince
	}

	if !dara.IsNil(request.EVenu) {
		query["EVenu"] = request.EVenu
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PostalCode) {
		query["PostalCode"] = request.PostalCode
	}

	if !dara.IsNil(request.RegType) {
		query["RegType"] = request.RegType
	}

	if !dara.IsNil(request.TelArea) {
		query["TelArea"] = request.TelArea
	}

	if !dara.IsNil(request.TelExt) {
		query["TelExt"] = request.TelExt
	}

	if !dara.IsNil(request.TelMain) {
		query["TelMain"] = request.TelMain
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveContactTemplate"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveContactTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存联系人模板
//
// @param request - SaveContactTemplateRequest
//
// @return SaveContactTemplateResponse
func (client *Client) SaveContactTemplate(request *SaveContactTemplateRequest) (_result *SaveContactTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveContactTemplateResponse{}
	_body, _err := client.SaveContactTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存联系人模板实名资料
//
// @param request - SaveContactTemplateCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveContactTemplateCredentialResponse
func (client *Client) SaveContactTemplateCredentialWithOptions(request *SaveContactTemplateCredentialRequest, runtime *dara.RuntimeOptions) (_result *SaveContactTemplateCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.Credential) {
		query["Credential"] = request.Credential
	}

	if !dara.IsNil(request.CredentialNo) {
		query["CredentialNo"] = request.CredentialNo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveContactTemplateCredential"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveContactTemplateCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存联系人模板实名资料
//
// @param request - SaveContactTemplateCredentialRequest
//
// @return SaveContactTemplateCredentialResponse
func (client *Client) SaveContactTemplateCredential(request *SaveContactTemplateCredentialRequest) (_result *SaveContactTemplateCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveContactTemplateCredentialResponse{}
	_body, _err := client.SaveContactTemplateCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改域名dns任务,对外开放接口，用于domain中
//
// @param request - SaveTaskForModifyingDomainDnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForModifyingDomainDnsResponse
func (client *Client) SaveTaskForModifyingDomainDnsWithOptions(request *SaveTaskForModifyingDomainDnsRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForModifyingDomainDnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunDns) {
		query["AliyunDns"] = request.AliyunDns
	}

	if !dara.IsNil(request.DnsList) {
		query["DnsList"] = request.DnsList
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SaleId) {
		query["SaleId"] = request.SaleId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForModifyingDomainDns"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForModifyingDomainDnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改域名dns任务,对外开放接口，用于domain中
//
// @param request - SaveTaskForModifyingDomainDnsRequest
//
// @return SaveTaskForModifyingDomainDnsResponse
func (client *Client) SaveTaskForModifyingDomainDns(request *SaveTaskForModifyingDomainDnsRequest) (_result *SaveTaskForModifyingDomainDnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveTaskForModifyingDomainDnsResponse{}
	_body, _err := client.SaveTaskForModifyingDomainDnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存域名实名资料任务
//
// @param request - SaveTaskForSubmittingDomainNameCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForSubmittingDomainNameCredentialResponse
func (client *Client) SaveTaskForSubmittingDomainNameCredentialWithOptions(request *SaveTaskForSubmittingDomainNameCredentialRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForSubmittingDomainNameCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Credential) {
		query["Credential"] = request.Credential
	}

	if !dara.IsNil(request.CredentialNo) {
		query["CredentialNo"] = request.CredentialNo
	}

	if !dara.IsNil(request.CredentialType) {
		query["CredentialType"] = request.CredentialType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SaleId) {
		query["SaleId"] = request.SaleId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForSubmittingDomainNameCredential"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForSubmittingDomainNameCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存域名实名资料任务
//
// @param request - SaveTaskForSubmittingDomainNameCredentialRequest
//
// @return SaveTaskForSubmittingDomainNameCredentialResponse
func (client *Client) SaveTaskForSubmittingDomainNameCredential(request *SaveTaskForSubmittingDomainNameCredentialRequest) (_result *SaveTaskForSubmittingDomainNameCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveTaskForSubmittingDomainNameCredentialResponse{}
	_body, _err := client.SaveTaskForSubmittingDomainNameCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据模板保存域名的实名认证信息
//
// @param request - SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse
func (client *Client) SaveTaskForSubmittingDomainNameCredentialByTemplateIdWithOptions(request *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SaleId) {
		query["SaleId"] = request.SaleId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForSubmittingDomainNameCredentialByTemplateId"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据模板保存域名的实名认证信息
//
// @param request - SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest
//
// @return SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse
func (client *Client) SaveTaskForSubmittingDomainNameCredentialByTemplateId(request *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) (_result *SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveTaskForSubmittingDomainNameCredentialByTemplateIdResponse{}
	_body, _err := client.SaveTaskForSubmittingDomainNameCredentialByTemplateIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存修改联系人的任务
//
// @param request - SaveTaskForUpdatingContactByTempateIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForUpdatingContactByTempateIdResponse
func (client *Client) SaveTaskForUpdatingContactByTempateIdWithOptions(request *SaveTaskForUpdatingContactByTempateIdRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForUpdatingContactByTempateIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddTransferLock) {
		query["AddTransferLock"] = request.AddTransferLock
	}

	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.ContactType) {
		query["ContactType"] = request.ContactType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SaleId) {
		query["SaleId"] = request.SaleId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForUpdatingContactByTempateId"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForUpdatingContactByTempateIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存修改联系人的任务
//
// @param request - SaveTaskForUpdatingContactByTempateIdRequest
//
// @return SaveTaskForUpdatingContactByTempateIdResponse
func (client *Client) SaveTaskForUpdatingContactByTempateId(request *SaveTaskForUpdatingContactByTempateIdRequest) (_result *SaveTaskForUpdatingContactByTempateIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveTaskForUpdatingContactByTempateIdResponse{}
	_body, _err := client.SaveTaskForUpdatingContactByTempateIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存修改联系人的任务
//
// @param request - SaveTaskForUpdatingContactByTemplateIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForUpdatingContactByTemplateIdResponse
func (client *Client) SaveTaskForUpdatingContactByTemplateIdWithOptions(request *SaveTaskForUpdatingContactByTemplateIdRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForUpdatingContactByTemplateIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddTransferLock) {
		query["AddTransferLock"] = request.AddTransferLock
	}

	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.ContactType) {
		query["ContactType"] = request.ContactType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SaleId) {
		query["SaleId"] = request.SaleId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForUpdatingContactByTemplateId"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForUpdatingContactByTemplateIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存修改联系人的任务
//
// @param request - SaveTaskForUpdatingContactByTemplateIdRequest
//
// @return SaveTaskForUpdatingContactByTemplateIdResponse
func (client *Client) SaveTaskForUpdatingContactByTemplateId(request *SaveTaskForUpdatingContactByTemplateIdRequest) (_result *SaveTaskForUpdatingContactByTemplateIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveTaskForUpdatingContactByTemplateIdResponse{}
	_body, _err := client.SaveTaskForUpdatingContactByTemplateIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启或者关闭whois保护
//
// @param request - WhoisProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WhoisProtectionResponse
func (client *Client) WhoisProtectionWithOptions(request *WhoisProtectionRequest, runtime *dara.RuntimeOptions) (_result *WhoisProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataContent) {
		query["DataContent"] = request.DataContent
	}

	if !dara.IsNil(request.DataSource) {
		query["DataSource"] = request.DataSource
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.WhoisProtect) {
		query["WhoisProtect"] = request.WhoisProtect
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WhoisProtection"),
		Version:     dara.String("2016-05-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WhoisProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启或者关闭whois保护
//
// @param request - WhoisProtectionRequest
//
// @return WhoisProtectionResponse
func (client *Client) WhoisProtection(request *WhoisProtectionRequest) (_result *WhoisProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &WhoisProtectionResponse{}
	_body, _err := client.WhoisProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
