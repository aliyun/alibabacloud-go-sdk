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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              dara.String("agency.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("agency.aliyuncs.com"),
		"ap-south-1":                  dara.String("agency.aliyuncs.com"),
		"ap-southeast-2":              dara.String("agency.aliyuncs.com"),
		"ap-southeast-3":              dara.String("agency.aliyuncs.com"),
		"ap-southeast-5":              dara.String("agency.aliyuncs.com"),
		"cn-beijing":                  dara.String("agency.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("agency.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("agency.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("agency.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("agency.aliyuncs.com"),
		"cn-chengdu":                  dara.String("agency.aliyuncs.com"),
		"cn-edge-1":                   dara.String("agency.aliyuncs.com"),
		"cn-fujian":                   dara.String("agency.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("agency.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("agency.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("agency.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("agency.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("agency.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("agency.aliyuncs.com"),
		"cn-hongkong":                 dara.String("agency.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("agency.aliyuncs.com"),
		"cn-huhehaote":                dara.String("agency.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("agency.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("agency.aliyuncs.com"),
		"cn-qingdao":                  dara.String("agency.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("agency.aliyuncs.com"),
		"cn-shanghai":                 dara.String("agency.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("agency.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("agency.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("agency.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("agency.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("agency.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("agency.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("agency.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("agency.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("agency.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("agency.aliyuncs.com"),
		"cn-wuhan":                    dara.String("agency.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("agency.aliyuncs.com"),
		"cn-yushanfang":               dara.String("agency.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("agency.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("agency.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("agency.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("agency.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("agency.aliyuncs.com"),
		"eu-central-1":                dara.String("agency.aliyuncs.com"),
		"eu-west-1":                   dara.String("agency.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("agency.aliyuncs.com"),
		"me-east-1":                   dara.String("agency.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("agency.aliyuncs.com"),
		"us-east-1":                   dara.String("agency.aliyuncs.com"),
		"us-west-1":                   dara.String("agency.aliyuncs.com"),
		"ap-southeast-1":              dara.String("agency.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("agency"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Queries exported bill files.
//
// @param request - GetBillDetailFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBillDetailFileListResponse
func (client *Client) GetBillDetailFileListWithOptions(request *GetBillDetailFileListRequest, runtime *dara.RuntimeOptions) (_result *GetBillDetailFileListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillMonth) {
		query["BillMonth"] = request.BillMonth
	}

	if !dara.IsNil(request.OssAccessKeyId) {
		query["OssAccessKeyId"] = request.OssAccessKeyId
	}

	if !dara.IsNil(request.OssAccessKeySecret) {
		query["OssAccessKeySecret"] = request.OssAccessKeySecret
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssRegion) {
		query["OssRegion"] = request.OssRegion
	}

	if !dara.IsNil(request.OssSecurityToken) {
		query["OssSecurityToken"] = request.OssSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBillDetailFileList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBillDetailFileListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries exported bill files.
//
// @param request - GetBillDetailFileListRequest
//
// @return GetBillDetailFileListResponse
func (client *Client) GetBillDetailFileList(request *GetBillDetailFileListRequest) (_result *GetBillDetailFileListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBillDetailFileListResponse{}
	_body, _err := client.GetBillDetailFileListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the commission details of a partner.
//
// @param request - GetCommissionDetailFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommissionDetailFileListResponse
func (client *Client) GetCommissionDetailFileListWithOptions(request *GetCommissionDetailFileListRequest, runtime *dara.RuntimeOptions) (_result *GetCommissionDetailFileListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillMonth) {
		query["BillMonth"] = request.BillMonth
	}

	if !dara.IsNil(request.OssAccessKeyId) {
		query["OssAccessKeyId"] = request.OssAccessKeyId
	}

	if !dara.IsNil(request.OssAccessKeySecret) {
		query["OssAccessKeySecret"] = request.OssAccessKeySecret
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssRegion) {
		query["OssRegion"] = request.OssRegion
	}

	if !dara.IsNil(request.OssSecurityToken) {
		query["OssSecurityToken"] = request.OssSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCommissionDetailFileList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCommissionDetailFileListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the commission details of a partner.
//
// @param request - GetCommissionDetailFileListRequest
//
// @return GetCommissionDetailFileListResponse
func (client *Client) GetCommissionDetailFileList(request *GetCommissionDetailFileListRequest) (_result *GetCommissionDetailFileListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCommissionDetailFileListResponse{}
	_body, _err := client.GetCommissionDetailFileListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries partner customer acquisition orders.
//
// @param tmpReq - GetCustomerOrderListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerOrderListResponse
func (client *Client) GetCustomerOrderListWithOptions(tmpReq *GetCustomerOrderListRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerOrderListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCustomerOrderListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderTypeList) {
		request.OrderTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderTypeList, dara.String("OrderTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerAccount) {
		query["CustomerAccount"] = request.CustomerAccount
	}

	if !dara.IsNil(request.CustomerUid) {
		query["CustomerUid"] = request.CustomerUid
	}

	if !dara.IsNil(request.OrderCreateAfter) {
		query["OrderCreateAfter"] = request.OrderCreateAfter
	}

	if !dara.IsNil(request.OrderCreateBefore) {
		query["OrderCreateBefore"] = request.OrderCreateBefore
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OrderPayAfter) {
		query["OrderPayAfter"] = request.OrderPayAfter
	}

	if !dara.IsNil(request.OrderPayBefore) {
		query["OrderPayBefore"] = request.OrderPayBefore
	}

	if !dara.IsNil(request.OrderStatus) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !dara.IsNil(request.OrderTypeListShrink) {
		query["OrderTypeList"] = request.OrderTypeListShrink
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PayAmountAfter) {
		query["PayAmountAfter"] = request.PayAmountAfter
	}

	if !dara.IsNil(request.PayAmountBefore) {
		query["PayAmountBefore"] = request.PayAmountBefore
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RamAccountForCustomerManager) {
		query["RamAccountForCustomerManager"] = request.RamAccountForCustomerManager
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerOrderList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerOrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries partner customer acquisition orders.
//
// @param request - GetCustomerOrderListRequest
//
// @return GetCustomerOrderListResponse
func (client *Client) GetCustomerOrderList(request *GetCustomerOrderListRequest) (_result *GetCustomerOrderListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomerOrderListResponse{}
	_body, _err := client.GetCustomerOrderListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Downloads the commission details of an international partner.
//
// Description:
//
// Make sure that the current caller identity is a T1 distribution partner.
//
// <notice>Available only for international sites.</notice>.
//
// @param request - GetIntlCommissionDetailFileListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntlCommissionDetailFileListResponse
func (client *Client) GetIntlCommissionDetailFileListWithOptions(request *GetIntlCommissionDetailFileListRequest, runtime *dara.RuntimeOptions) (_result *GetIntlCommissionDetailFileListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillMonth) {
		query["BillMonth"] = request.BillMonth
	}

	if !dara.IsNil(request.OssAccessKeyId) {
		query["OssAccessKeyId"] = request.OssAccessKeyId
	}

	if !dara.IsNil(request.OssAccessKeySecret) {
		query["OssAccessKeySecret"] = request.OssAccessKeySecret
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssRegion) {
		query["OssRegion"] = request.OssRegion
	}

	if !dara.IsNil(request.OssSecurityToken) {
		query["OssSecurityToken"] = request.OssSecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIntlCommissionDetailFileList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIntlCommissionDetailFileListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downloads the commission details of an international partner.
//
// Description:
//
// Make sure that the current caller identity is a T1 distribution partner.
//
// <notice>Available only for international sites.</notice>.
//
// @param request - GetIntlCommissionDetailFileListRequest
//
// @return GetIntlCommissionDetailFileListResponse
func (client *Client) GetIntlCommissionDetailFileList(request *GetIntlCommissionDetailFileListRequest) (_result *GetIntlCommissionDetailFileListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIntlCommissionDetailFileListResponse{}
	_body, _err := client.GetIntlCommissionDetailFileListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the partner renewal rate.
//
// @param request - GetRenewalRateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRenewalRateListResponse
func (client *Client) GetRenewalRateListWithOptions(request *GetRenewalRateListRequest, runtime *dara.RuntimeOptions) (_result *GetRenewalRateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FiscalYearAndQuarter) {
		query["FiscalYearAndQuarter"] = request.FiscalYearAndQuarter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRenewalRateList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRenewalRateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the partner renewal rate.
//
// @param request - GetRenewalRateListRequest
//
// @return GetRenewalRateListResponse
func (client *Client) GetRenewalRateList(request *GetRenewalRateListRequest) (_result *GetRenewalRateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRenewalRateListResponse{}
	_body, _err := client.GetRenewalRateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of secondary distributors.
//
// @param request - GetSubPartnerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubPartnerListResponse
func (client *Client) GetSubPartnerListWithOptions(request *GetSubPartnerListRequest, runtime *dara.RuntimeOptions) (_result *GetSubPartnerListResponse, _err error) {
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

	if !dara.IsNil(request.SubPartnerCompanyName) {
		query["SubPartnerCompanyName"] = request.SubPartnerCompanyName
	}

	if !dara.IsNil(request.SubPartnerPid) {
		query["SubPartnerPid"] = request.SubPartnerPid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubPartnerList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubPartnerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of secondary distributors.
//
// @param request - GetSubPartnerListRequest
//
// @return GetSubPartnerListResponse
func (client *Client) GetSubPartnerList(request *GetSubPartnerListRequest) (_result *GetSubPartnerListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSubPartnerListResponse{}
	_body, _err := client.GetSubPartnerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries channel expansion orders.
//
// @param tmpReq - GetSubPartnerOrderListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubPartnerOrderListResponse
func (client *Client) GetSubPartnerOrderListWithOptions(tmpReq *GetSubPartnerOrderListRequest, runtime *dara.RuntimeOptions) (_result *GetSubPartnerOrderListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSubPartnerOrderListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderTypeList) {
		request.OrderTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderTypeList, dara.String("OrderTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderCreateAfter) {
		query["OrderCreateAfter"] = request.OrderCreateAfter
	}

	if !dara.IsNil(request.OrderCreateBefore) {
		query["OrderCreateBefore"] = request.OrderCreateBefore
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OrderPayAfter) {
		query["OrderPayAfter"] = request.OrderPayAfter
	}

	if !dara.IsNil(request.OrderPayBefore) {
		query["OrderPayBefore"] = request.OrderPayBefore
	}

	if !dara.IsNil(request.OrderStatus) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !dara.IsNil(request.OrderTypeListShrink) {
		query["OrderTypeList"] = request.OrderTypeListShrink
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PayAmountAfter) {
		query["PayAmountAfter"] = request.PayAmountAfter
	}

	if !dara.IsNil(request.PayAmountBefore) {
		query["PayAmountBefore"] = request.PayAmountBefore
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SubPartnerName) {
		query["SubPartnerName"] = request.SubPartnerName
	}

	if !dara.IsNil(request.SubPartnerUid) {
		query["SubPartnerUid"] = request.SubPartnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubPartnerOrderList"),
		Version:     dara.String("2025-02-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubPartnerOrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries channel expansion orders.
//
// @param request - GetSubPartnerOrderListRequest
//
// @return GetSubPartnerOrderListResponse
func (client *Client) GetSubPartnerOrderList(request *GetSubPartnerOrderListRequest) (_result *GetSubPartnerOrderListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSubPartnerOrderListResponse{}
	_body, _err := client.GetSubPartnerOrderListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
