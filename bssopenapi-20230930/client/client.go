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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 dara.String("business.aliyuncs.com"),
		"cn-shanghai":                 dara.String("business.aliyuncs.com"),
		"ap-southeast-1":              dara.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-1":              dara.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2":              dara.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":                  dara.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              dara.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":              dara.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":              dara.String("business.ap-southeast-1.aliyuncs.com"),
		"cn-beijing":                  dara.String("business.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("business.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("business.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("business.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("business.aliyuncs.com"),
		"cn-chengdu":                  dara.String("business.aliyuncs.com"),
		"cn-edge-1":                   dara.String("business.aliyuncs.com"),
		"cn-fujian":                   dara.String("business.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("business.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("business.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("business.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("business.aliyuncs.com"),
		"cn-hongkong":                 dara.String("business.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("business.aliyuncs.com"),
		"cn-huhehaote":                dara.String("business.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("business.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("business.aliyuncs.com"),
		"cn-qingdao":                  dara.String("business.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("business.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("business.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("business.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("business.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("business.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("business.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("business.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("business.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("business.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("business.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("business.aliyuncs.com"),
		"cn-wuhan":                    dara.String("business.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("business.aliyuncs.com"),
		"cn-yushanfang":               dara.String("business.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("business.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("business.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("business.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("business.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("business.aliyuncs.com"),
		"eu-central-1":                dara.String("business.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":                   dara.String("business.ap-southeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("business.ap-southeast-1.aliyuncs.com"),
		"me-east-1":                   dara.String("business.ap-southeast-1.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("business.ap-southeast-1.aliyuncs.com"),
		"us-east-1":                   dara.String("business.ap-southeast-1.aliyuncs.com"),
		"us-west-1":                   dara.String("business.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("bssopenapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加优惠券抵扣标签
//
// @param tmpReq - AddCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCouponDeductTagResponse
func (client *Client) AddCouponDeductTagWithOptions(tmpReq *AddCouponDeductTagRequest, runtime *dara.RuntimeOptions) (_result *AddCouponDeductTagResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AddCouponDeductTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponId) {
		query["CouponId"] = request.CouponId
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCouponDeductTag"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCouponDeductTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加优惠券抵扣标签
//
// @param request - AddCouponDeductTagRequest
//
// @return AddCouponDeductTagResponse
func (client *Client) AddCouponDeductTag(request *AddCouponDeductTagRequest) (_result *AddCouponDeductTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCouponDeductTagResponse{}
	_body, _err := client.AddCouponDeductTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 财务单元实例重分配
//
// @param tmpReq - AllocateCostCenterResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateCostCenterResourceResponse
func (client *Client) AllocateCostCenterResourceWithOptions(tmpReq *AllocateCostCenterResourceRequest, runtime *dara.RuntimeOptions) (_result *AllocateCostCenterResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AllocateCostCenterResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceInstanceList) {
		request.ResourceInstanceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceInstanceList, dara.String("ResourceInstanceList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FromCostCenterId) {
		body["FromCostCenterId"] = request.FromCostCenterId
	}

	if !dara.IsNil(request.FromOwnerAccountId) {
		body["FromOwnerAccountId"] = request.FromOwnerAccountId
	}

	if !dara.IsNil(request.ResourceInstanceListShrink) {
		body["ResourceInstanceList"] = request.ResourceInstanceListShrink
	}

	if !dara.IsNil(request.ToCostCenterId) {
		body["ToCostCenterId"] = request.ToCostCenterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateCostCenterResource"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateCostCenterResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 财务单元实例重分配
//
// @param request - AllocateCostCenterResourceRequest
//
// @return AllocateCostCenterResourceResponse
func (client *Client) AllocateCostCenterResource(request *AllocateCostCenterResourceRequest) (_result *AllocateCostCenterResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateCostCenterResourceResponse{}
	_body, _err := client.AllocateCostCenterResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消资金账户低额预警
//
// @param request - CancelFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelFundAccountLowAvailableAmountAlarmResponse
func (client *Client) CancelFundAccountLowAvailableAmountAlarmWithOptions(request *CancelFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *CancelFundAccountLowAvailableAmountAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelFundAccountLowAvailableAmountAlarm"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消资金账户低额预警
//
// @param request - CancelFundAccountLowAvailableAmountAlarmRequest
//
// @return CancelFundAccountLowAvailableAmountAlarmResponse
func (client *Client) CancelFundAccountLowAvailableAmountAlarm(request *CancelFundAccountLowAvailableAmountAlarmRequest) (_result *CancelFundAccountLowAvailableAmountAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.CancelFundAccountLowAvailableAmountAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建财务单元
//
// @param tmpReq - CreateCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostCenterResponse
func (client *Client) CreateCostCenterWithOptions(tmpReq *CreateCostCenterRequest, runtime *dara.RuntimeOptions) (_result *CreateCostCenterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateCostCenterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CostCenterEntityList) {
		request.CostCenterEntityListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CostCenterEntityList, dara.String("CostCenterEntityList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CostCenterEntityListShrink) {
		query["CostCenterEntityList"] = request.CostCenterEntityListShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCostCenter"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCostCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财务单元
//
// @param request - CreateCostCenterRequest
//
// @return CreateCostCenterResponse
func (client *Client) CreateCostCenter(request *CreateCostCenterRequest) (_result *CreateCostCenterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCostCenterResponse{}
	_body, _err := client.CreateCostCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新建财务单元规则
//
// @param tmpReq - CreateCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostCenterRuleResponse
func (client *Client) CreateCostCenterRuleWithOptions(tmpReq *CreateCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateCostCenterRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateCostCenterRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterExpression) {
		request.FilterExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterExpression, dara.String("FilterExpression"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterExpressionShrink) {
		query["FilterExpression"] = request.FilterExpressionShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CostCenterId) {
		body["CostCenterId"] = request.CostCenterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCostCenterRule"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCostCenterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建财务单元规则
//
// @param request - CreateCostCenterRuleRequest
//
// @return CreateCostCenterRuleResponse
func (client *Client) CreateCostCenterRule(request *CreateCostCenterRuleRequest) (_result *CreateCostCenterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCostCenterRuleResponse{}
	_body, _err := client.CreateCostCenterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资金账户付款关系
//
// @param tmpReq - CreateFundAccountPayRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFundAccountPayRelationResponse
func (client *Client) CreateFundAccountPayRelationWithOptions(tmpReq *CreateFundAccountPayRelationRequest, runtime *dara.RuntimeOptions) (_result *CreateFundAccountPayRelationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateFundAccountPayRelationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFundAccountPayRelation"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFundAccountPayRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资金账户付款关系
//
// @param request - CreateFundAccountPayRelationRequest
//
// @return CreateFundAccountPayRelationResponse
func (client *Client) CreateFundAccountPayRelation(request *CreateFundAccountPayRelationRequest) (_result *CreateFundAccountPayRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFundAccountPayRelationResponse{}
	_body, _err := client.CreateFundAccountPayRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资金账户划拨/回收
//
// @param request - CreateFundAccountTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFundAccountTransferResponse
func (client *Client) CreateFundAccountTransferWithOptions(request *CreateFundAccountTransferRequest, runtime *dara.RuntimeOptions) (_result *CreateFundAccountTransferResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		body["Amount"] = request.Amount
	}

	if !dara.IsNil(request.Currency) {
		body["Currency"] = request.Currency
	}

	if !dara.IsNil(request.FinanceType) {
		body["FinanceType"] = request.FinanceType
	}

	if !dara.IsNil(request.FromFundAccountId) {
		body["FromFundAccountId"] = request.FromFundAccountId
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ToFundAccountId) {
		body["ToFundAccountId"] = request.ToFundAccountId
	}

	if !dara.IsNil(request.TransferType) {
		body["TransferType"] = request.TransferType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFundAccountTransfer"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFundAccountTransferResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资金账户划拨/回收
//
// @param request - CreateFundAccountTransferRequest
//
// @return CreateFundAccountTransferResponse
func (client *Client) CreateFundAccountTransfer(request *CreateFundAccountTransferRequest) (_result *CreateFundAccountTransferResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFundAccountTransferResponse{}
	_body, _err := client.CreateFundAccountTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申请发票
//
// @param tmpReq - CreateInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInvoiceResponse
func (client *Client) CreateInvoiceWithOptions(tmpReq *CreateInvoiceRequest, runtime *dara.RuntimeOptions) (_result *CreateInvoiceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateInvoiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InvoiceCandidateIds) {
		request.InvoiceCandidateIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InvoiceCandidateIds, dara.String("InvoiceCandidateIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RecipientEmails) {
		request.RecipientEmailsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecipientEmails, dara.String("RecipientEmails"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.InvoiceCandidateIdsShrink) {
		query["InvoiceCandidateIds"] = request.InvoiceCandidateIdsShrink
	}

	if !dara.IsNil(request.InvoiceMode) {
		query["InvoiceMode"] = request.InvoiceMode
	}

	if !dara.IsNil(request.InvoiceRemark) {
		query["InvoiceRemark"] = request.InvoiceRemark
	}

	if !dara.IsNil(request.InvoiceTitleId) {
		query["InvoiceTitleId"] = request.InvoiceTitleId
	}

	if !dara.IsNil(request.InvoiceType) {
		query["InvoiceType"] = request.InvoiceType
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.RecipientEmailsShrink) {
		query["RecipientEmails"] = request.RecipientEmailsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInvoice"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请发票
//
// @param request - CreateInvoiceRequest
//
// @return CreateInvoiceResponse
func (client *Client) CreateInvoice(request *CreateInvoiceRequest) (_result *CreateInvoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInvoiceResponse{}
	_body, _err := client.CreateInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建账单订阅
//
// @param request - CreateReportDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportDefinitionResponse
func (client *Client) CreateReportDefinitionWithOptions(request *CreateReportDefinitionRequest, runtime *dara.RuntimeOptions) (_result *CreateReportDefinitionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginBillingCycle) {
		query["BeginBillingCycle"] = request.BeginBillingCycle
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssBucketOwnerAccountId) {
		query["OssBucketOwnerAccountId"] = request.OssBucketOwnerAccountId
	}

	if !dara.IsNil(request.OssBucketPath) {
		query["OssBucketPath"] = request.OssBucketPath
	}

	if !dara.IsNil(request.ReportType) {
		query["ReportType"] = request.ReportType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.McProject) {
		body["McProject"] = request.McProject
	}

	if !dara.IsNil(request.McTableName) {
		body["McTableName"] = request.McTableName
	}

	if !dara.IsNil(request.ReportSourceType) {
		body["ReportSourceType"] = request.ReportSourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReportDefinition"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReportDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建账单订阅
//
// @param request - CreateReportDefinitionRequest
//
// @return CreateReportDefinitionResponse
func (client *Client) CreateReportDefinition(request *CreateReportDefinitionRequest) (_result *CreateReportDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateReportDefinitionResponse{}
	_body, _err := client.CreateReportDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除财务单元
//
// @param request - DeleteCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostCenterResponse
func (client *Client) DeleteCostCenterWithOptions(request *DeleteCostCenterRequest, runtime *dara.RuntimeOptions) (_result *DeleteCostCenterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CostCenterId) {
		query["CostCenterId"] = request.CostCenterId
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.OwnerAccountId) {
		query["OwnerAccountId"] = request.OwnerAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCostCenter"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCostCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除财务单元
//
// @param request - DeleteCostCenterRequest
//
// @return DeleteCostCenterResponse
func (client *Client) DeleteCostCenter(request *DeleteCostCenterRequest) (_result *DeleteCostCenterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCostCenterResponse{}
	_body, _err := client.DeleteCostCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除财务单元规则
//
// @param tmpReq - DeleteCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostCenterRuleResponse
func (client *Client) DeleteCostCenterRuleWithOptions(tmpReq *DeleteCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteCostCenterRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteCostCenterRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterExpression) {
		request.FilterExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterExpression, dara.String("FilterExpression"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterExpressionShrink) {
		query["FilterExpression"] = request.FilterExpressionShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CostCenterId) {
		body["CostCenterId"] = request.CostCenterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCostCenterRule"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCostCenterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除财务单元规则
//
// @param request - DeleteCostCenterRuleRequest
//
// @return DeleteCostCenterRuleResponse
func (client *Client) DeleteCostCenterRule(request *DeleteCostCenterRuleRequest) (_result *DeleteCostCenterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCostCenterRuleResponse{}
	_body, _err := client.DeleteCostCenterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除优惠券的抵扣标签
//
// @param tmpReq - DeleteCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCouponDeductTagResponse
func (client *Client) DeleteCouponDeductTagWithOptions(tmpReq *DeleteCouponDeductTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteCouponDeductTagResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteCouponDeductTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKeys) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, dara.String("TagKeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponId) {
		query["CouponId"] = request.CouponId
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.TagKeysShrink) {
		query["TagKeys"] = request.TagKeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCouponDeductTag"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCouponDeductTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除优惠券的抵扣标签
//
// @param request - DeleteCouponDeductTagRequest
//
// @return DeleteCouponDeductTagResponse
func (client *Client) DeleteCouponDeductTag(request *DeleteCouponDeductTagRequest) (_result *DeleteCouponDeductTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCouponDeductTagResponse{}
	_body, _err := client.DeleteCouponDeductTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消账单订阅
//
// @param request - DeleteReportDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReportDefinitionResponse
func (client *Client) DeleteReportDefinitionWithOptions(request *DeleteReportDefinitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteReportDefinitionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.ReportTaskId) {
		query["ReportTaskId"] = request.ReportTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteReportDefinition"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReportDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消账单订阅
//
// @param request - DeleteReportDefinitionRequest
//
// @return DeleteReportDefinitionResponse
func (client *Client) DeleteReportDefinition(request *DeleteReportDefinitionRequest) (_result *DeleteReportDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteReportDefinitionResponse{}
	_body, _err := client.DeleteReportDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询优惠券列表
//
// @param tmpReq - DescribeCouponRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponResponse
func (client *Client) DescribeCouponWithOptions(tmpReq *DescribeCouponRequest, runtime *dara.RuntimeOptions) (_result *DescribeCouponResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCouponShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCoupon"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCouponResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询优惠券列表
//
// @param request - DescribeCouponRequest
//
// @return DescribeCouponResponse
func (client *Client) DescribeCoupon(request *DescribeCouponRequest) (_result *DescribeCouponResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCouponResponse{}
	_body, _err := client.DescribeCouponWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询优惠券可用商品列表
//
// @param tmpReq - DescribeCouponItemListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponItemListResponse
func (client *Client) DescribeCouponItemListWithOptions(tmpReq *DescribeCouponItemListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCouponItemListResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCouponItemListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCouponItemList"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCouponItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询优惠券可用商品列表
//
// @param request - DescribeCouponItemListRequest
//
// @return DescribeCouponItemListResponse
func (client *Client) DescribeCouponItemList(request *DescribeCouponItemListRequest) (_result *DescribeCouponItemListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCouponItemListResponse{}
	_body, _err := client.DescribeCouponItemListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取客户使用SPN的概述信息
//
// @param tmpReq - DescribeUserSpnSummaryInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserSpnSummaryInfoResponse
func (client *Client) DescribeUserSpnSummaryInfoWithOptions(tmpReq *DescribeUserSpnSummaryInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserSpnSummaryInfoResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeUserSpnSummaryInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserSpnSummaryInfo"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserSpnSummaryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取客户使用SPN的概述信息
//
// @param request - DescribeUserSpnSummaryInfoRequest
//
// @return DescribeUserSpnSummaryInfoResponse
func (client *Client) DescribeUserSpnSummaryInfo(request *DescribeUserSpnSummaryInfoRequest) (_result *DescribeUserSpnSummaryInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserSpnSummaryInfoResponse{}
	_body, _err := client.DescribeUserSpnSummaryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户可用金
//
// @param request - GetFundAccountAvailableAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountAvailableAmountResponse
func (client *Client) GetFundAccountAvailableAmountWithOptions(request *GetFundAccountAvailableAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountAvailableAmountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFundAccountAvailableAmount"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFundAccountAvailableAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户可用金
//
// @param request - GetFundAccountAvailableAmountRequest
//
// @return GetFundAccountAvailableAmountResponse
func (client *Client) GetFundAccountAvailableAmount(request *GetFundAccountAvailableAmountRequest) (_result *GetFundAccountAvailableAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFundAccountAvailableAmountResponse{}
	_body, _err := client.GetFundAccountAvailableAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户可分配信控额度
//
// @param request - GetFundAccountCanAllocateCreditAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanAllocateCreditAmountResponse
func (client *Client) GetFundAccountCanAllocateCreditAmountWithOptions(request *GetFundAccountCanAllocateCreditAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanAllocateCreditAmountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFundAccountCanAllocateCreditAmount"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFundAccountCanAllocateCreditAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户可分配信控额度
//
// @param request - GetFundAccountCanAllocateCreditAmountRequest
//
// @return GetFundAccountCanAllocateCreditAmountResponse
func (client *Client) GetFundAccountCanAllocateCreditAmount(request *GetFundAccountCanAllocateCreditAmountRequest) (_result *GetFundAccountCanAllocateCreditAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFundAccountCanAllocateCreditAmountResponse{}
	_body, _err := client.GetFundAccountCanAllocateCreditAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户可回收金额
//
// @param request - GetFundAccountCanRecycleAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanRecycleAmountResponse
func (client *Client) GetFundAccountCanRecycleAmountWithOptions(request *GetFundAccountCanRecycleAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanRecycleAmountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Currency) {
		body["Currency"] = request.Currency
	}

	if !dara.IsNil(request.RecycleFromFundAccountId) {
		body["RecycleFromFundAccountId"] = request.RecycleFromFundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFundAccountCanRecycleAmount"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFundAccountCanRecycleAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户可回收金额
//
// @param request - GetFundAccountCanRecycleAmountRequest
//
// @return GetFundAccountCanRecycleAmountResponse
func (client *Client) GetFundAccountCanRecycleAmount(request *GetFundAccountCanRecycleAmountRequest) (_result *GetFundAccountCanRecycleAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFundAccountCanRecycleAmountResponse{}
	_body, _err := client.GetFundAccountCanRecycleAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户的可转出金额
//
// @param request - GetFundAccountCanTransferAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanTransferAmountResponse
func (client *Client) GetFundAccountCanTransferAmountWithOptions(request *GetFundAccountCanTransferAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanTransferAmountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Currency) {
		body["Currency"] = request.Currency
	}

	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFundAccountCanTransferAmount"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFundAccountCanTransferAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户的可转出金额
//
// @param request - GetFundAccountCanTransferAmountRequest
//
// @return GetFundAccountCanTransferAmountResponse
func (client *Client) GetFundAccountCanTransferAmount(request *GetFundAccountCanTransferAmountRequest) (_result *GetFundAccountCanTransferAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFundAccountCanTransferAmountResponse{}
	_body, _err := client.GetFundAccountCanTransferAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户可提现金额
//
// @param request - GetFundAccountCanWithdrawAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanWithdrawAmountResponse
func (client *Client) GetFundAccountCanWithdrawAmountWithOptions(request *GetFundAccountCanWithdrawAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanWithdrawAmountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFundAccountCanWithdrawAmount"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFundAccountCanWithdrawAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户可提现金额
//
// @param request - GetFundAccountCanWithdrawAmountRequest
//
// @return GetFundAccountCanWithdrawAmountResponse
func (client *Client) GetFundAccountCanWithdrawAmount(request *GetFundAccountCanWithdrawAmountRequest) (_result *GetFundAccountCanWithdrawAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFundAccountCanWithdrawAmountResponse{}
	_body, _err := client.GetFundAccountCanWithdrawAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户低额预警
//
// @param request - GetFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) GetFundAccountLowAvailableAmountAlarmWithOptions(request *GetFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFundAccountLowAvailableAmountAlarm"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户低额预警
//
// @param request - GetFundAccountLowAvailableAmountAlarmRequest
//
// @return GetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) GetFundAccountLowAvailableAmountAlarm(request *GetFundAccountLowAvailableAmountAlarmRequest) (_result *GetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.GetFundAccountLowAvailableAmountAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户收支明细
//
// @param tmpReq - GetFundAccountTransactionDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountTransactionDetailsResponse
func (client *Client) GetFundAccountTransactionDetailsWithOptions(tmpReq *GetFundAccountTransactionDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountTransactionDetailsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetFundAccountTransactionDetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TransactionChannelList) {
		request.TransactionChannelListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransactionChannelList, dara.String("TransactionChannelList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransactionTypeList) {
		request.TransactionTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransactionTypeList, dara.String("TransactionTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BillNumber) {
		body["BillNumber"] = request.BillNumber
	}

	if !dara.IsNil(request.ChannelTransactionNumber) {
		body["ChannelTransactionNumber"] = request.ChannelTransactionNumber
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TransactionChannelListShrink) {
		body["TransactionChannelList"] = request.TransactionChannelListShrink
	}

	if !dara.IsNil(request.TransactionDirection) {
		body["TransactionDirection"] = request.TransactionDirection
	}

	if !dara.IsNil(request.TransactionNumber) {
		body["TransactionNumber"] = request.TransactionNumber
	}

	if !dara.IsNil(request.TransactionType) {
		body["TransactionType"] = request.TransactionType
	}

	if !dara.IsNil(request.TransactionTypeListShrink) {
		body["TransactionTypeList"] = request.TransactionTypeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFundAccountTransactionDetails"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFundAccountTransactionDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户收支明细
//
// @param request - GetFundAccountTransactionDetailsRequest
//
// @return GetFundAccountTransactionDetailsResponse
func (client *Client) GetFundAccountTransactionDetails(request *GetFundAccountTransactionDetailsRequest) (_result *GetFundAccountTransactionDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFundAccountTransactionDetailsResponse{}
	_body, _err := client.GetFundAccountTransactionDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 订单详情查询
//
// @param request - GetOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderDetailResponse
func (client *Client) GetOrderDetailWithOptions(request *GetOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetOrderDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrderDetail"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订单详情查询
//
// @param request - GetOrderDetailRequest
//
// @return GetOrderDetailResponse
func (client *Client) GetOrderDetail(request *GetOrderDetailRequest) (_result *GetOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOrderDetailResponse{}
	_body, _err := client.GetOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 订单列表查询
//
// @param request - GetOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrdersResponse
func (client *Client) GetOrdersWithOptions(request *GetOrdersRequest, runtime *dara.RuntimeOptions) (_result *GetOrdersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PaymentStatus) {
		query["PaymentStatus"] = request.PaymentStatus
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrders"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订单列表查询
//
// @param request - GetOrdersRequest
//
// @return GetOrdersResponse
func (client *Client) GetOrders(request *GetOrdersRequest) (_result *GetOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOrdersResponse{}
	_body, _err := client.GetOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节省计划及可抵扣商品信息
//
// @param tmpReq - GetSavingPlanDeductableCommodityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanDeductableCommodityResponse
func (client *Client) GetSavingPlanDeductableCommodityWithOptions(tmpReq *GetSavingPlanDeductableCommodityRequest, runtime *dara.RuntimeOptions) (_result *GetSavingPlanDeductableCommodityResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetSavingPlanDeductableCommodityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSavingPlanDeductableCommodity"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSavingPlanDeductableCommodityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节省计划及可抵扣商品信息
//
// @param request - GetSavingPlanDeductableCommodityRequest
//
// @return GetSavingPlanDeductableCommodityResponse
func (client *Client) GetSavingPlanDeductableCommodity(request *GetSavingPlanDeductableCommodityRequest) (_result *GetSavingPlanDeductableCommodityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSavingPlanDeductableCommodityResponse{}
	_body, _err := client.GetSavingPlanDeductableCommodityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节省计划实例共享账号信息
//
// @param tmpReq - GetSavingPlanShareAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanShareAccountsResponse
func (client *Client) GetSavingPlanShareAccountsWithOptions(tmpReq *GetSavingPlanShareAccountsRequest, runtime *dara.RuntimeOptions) (_result *GetSavingPlanShareAccountsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetSavingPlanShareAccountsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SpnInstanceCode) {
		query["SpnInstanceCode"] = request.SpnInstanceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSavingPlanShareAccounts"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSavingPlanShareAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节省计划实例共享账号信息
//
// @param request - GetSavingPlanShareAccountsRequest
//
// @return GetSavingPlanShareAccountsResponse
func (client *Client) GetSavingPlanShareAccounts(request *GetSavingPlanShareAccountsRequest) (_result *GetSavingPlanShareAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSavingPlanShareAccountsResponse{}
	_body, _err := client.GetSavingPlanShareAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节省计划实例客户自定义规则
//
// @param tmpReq - GetSavingPlanUserDeductRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanUserDeductRuleResponse
func (client *Client) GetSavingPlanUserDeductRuleWithOptions(tmpReq *GetSavingPlanUserDeductRuleRequest, runtime *dara.RuntimeOptions) (_result *GetSavingPlanUserDeductRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetSavingPlanUserDeductRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SpnInstanceCode) {
		query["SpnInstanceCode"] = request.SpnInstanceCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSavingPlanUserDeductRule"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSavingPlanUserDeductRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节省计划实例客户自定义规则
//
// @param request - GetSavingPlanUserDeductRuleRequest
//
// @return GetSavingPlanUserDeductRuleResponse
func (client *Client) GetSavingPlanUserDeductRule(request *GetSavingPlanUserDeductRuleRequest) (_result *GetSavingPlanUserDeductRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSavingPlanUserDeductRuleResponse{}
	_body, _err := client.GetSavingPlanUserDeductRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询优惠券设置的抵扣标签
//
// @param tmpReq - ListCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCouponDeductTagResponse
func (client *Client) ListCouponDeductTagWithOptions(tmpReq *ListCouponDeductTagRequest, runtime *dara.RuntimeOptions) (_result *ListCouponDeductTagResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListCouponDeductTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponId) {
		query["CouponId"] = request.CouponId
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCouponDeductTag"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCouponDeductTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询优惠券设置的抵扣标签
//
// @param request - ListCouponDeductTagRequest
//
// @return ListCouponDeductTagResponse
func (client *Client) ListCouponDeductTag(request *ListCouponDeductTagRequest) (_result *ListCouponDeductTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCouponDeductTagResponse{}
	_body, _err := client.ListCouponDeductTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户列表
//
// @param request - ListFundAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFundAccountResponse
func (client *Client) ListFundAccountWithOptions(request *ListFundAccountRequest, runtime *dara.RuntimeOptions) (_result *ListFundAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.QueryOnlyInUse) {
		body["QueryOnlyInUse"] = request.QueryOnlyInUse
	}

	if !dara.IsNil(request.QueryOnlyManage) {
		body["QueryOnlyManage"] = request.QueryOnlyManage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFundAccount"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFundAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户列表
//
// @param request - ListFundAccountRequest
//
// @return ListFundAccountResponse
func (client *Client) ListFundAccount(request *ListFundAccountRequest) (_result *ListFundAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFundAccountResponse{}
	_body, _err := client.ListFundAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资金账户的付款关系
//
// @param request - ListFundAccountPayRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFundAccountPayRelationResponse
func (client *Client) ListFundAccountPayRelationWithOptions(request *ListFundAccountPayRelationRequest, runtime *dara.RuntimeOptions) (_result *ListFundAccountPayRelationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFundAccountPayRelation"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFundAccountPayRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资金账户的付款关系
//
// @param request - ListFundAccountPayRelationRequest
//
// @return ListFundAccountPayRelationResponse
func (client *Client) ListFundAccountPayRelation(request *ListFundAccountPayRelationRequest) (_result *ListFundAccountPayRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFundAccountPayRelationResponse{}
	_body, _err := client.ListFundAccountPayRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对客OpenAPI开票对象查询
//
// @param tmpReq - ListInvoiceCandidateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInvoiceCandidateResponse
func (client *Client) ListInvoiceCandidateWithOptions(tmpReq *ListInvoiceCandidateRequest, runtime *dara.RuntimeOptions) (_result *ListInvoiceCandidateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListInvoiceCandidateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BillingCycles) {
		request.BillingCyclesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BillingCycles, dara.String("BillingCycles"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BusinessIds) {
		request.BusinessIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BusinessIds, dara.String("BusinessIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InvoiceIssuers) {
		request.InvoiceIssuersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InvoiceIssuers, dara.String("InvoiceIssuers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Status) {
		request.StatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Status, dara.String("Status"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingCyclesShrink) {
		query["BillingCycles"] = request.BillingCyclesShrink
	}

	if !dara.IsNil(request.BusinessIdsShrink) {
		query["BusinessIds"] = request.BusinessIdsShrink
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InvoiceIssuersShrink) {
		query["InvoiceIssuers"] = request.InvoiceIssuersShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatusShrink) {
		query["Status"] = request.StatusShrink
	}

	if !dara.IsNil(request.TypesShrink) {
		query["Types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInvoiceCandidate"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInvoiceCandidateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对客OpenAPI开票对象查询
//
// @param request - ListInvoiceCandidateRequest
//
// @return ListInvoiceCandidateResponse
func (client *Client) ListInvoiceCandidate(request *ListInvoiceCandidateRequest) (_result *ListInvoiceCandidateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInvoiceCandidateResponse{}
	_body, _err := client.ListInvoiceCandidateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发票抬头查询服务
//
// @param request - ListInvoiceTitleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInvoiceTitleResponse
func (client *Client) ListInvoiceTitleWithOptions(runtime *dara.RuntimeOptions) (_result *ListInvoiceTitleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListInvoiceTitle"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInvoiceTitleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发票抬头查询服务
//
// @return ListInvoiceTitleResponse
func (client *Client) ListInvoiceTitle() (_result *ListInvoiceTitleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInvoiceTitleResponse{}
	_body, _err := client.ListInvoiceTitleWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看已订阅的报告列表
//
// @param request - ListReportDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportDefinitionsResponse
func (client *Client) ListReportDefinitionsWithOptions(request *ListReportDefinitionsRequest, runtime *dara.RuntimeOptions) (_result *ListReportDefinitionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReportDefinitions"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReportDefinitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看已订阅的报告列表
//
// @param request - ListReportDefinitionsRequest
//
// @return ListReportDefinitionsResponse
func (client *Client) ListReportDefinitions(request *ListReportDefinitionsRequest) (_result *ListReportDefinitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListReportDefinitionsResponse{}
	_body, _err := client.ListReportDefinitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改财务单元
//
// @param tmpReq - ModifyCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostCenterResponse
func (client *Client) ModifyCostCenterWithOptions(tmpReq *ModifyCostCenterRequest, runtime *dara.RuntimeOptions) (_result *ModifyCostCenterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyCostCenterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CostCenterEntityList) {
		request.CostCenterEntityListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CostCenterEntityList, dara.String("CostCenterEntityList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CostCenterEntityListShrink) {
		query["CostCenterEntityList"] = request.CostCenterEntityListShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCostCenter"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCostCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改财务单元
//
// @param request - ModifyCostCenterRequest
//
// @return ModifyCostCenterResponse
func (client *Client) ModifyCostCenter(request *ModifyCostCenterRequest) (_result *ModifyCostCenterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCostCenterResponse{}
	_body, _err := client.ModifyCostCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改财务单元规则
//
// @param tmpReq - ModifyCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostCenterRuleResponse
func (client *Client) ModifyCostCenterRuleWithOptions(tmpReq *ModifyCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyCostCenterRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyCostCenterRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterExpression) {
		request.FilterExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterExpression, dara.String("FilterExpression"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterExpressionShrink) {
		query["FilterExpression"] = request.FilterExpressionShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CostCenterId) {
		body["CostCenterId"] = request.CostCenterId
	}

	if !dara.IsNil(request.OwnerAccountId) {
		body["OwnerAccountId"] = request.OwnerAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCostCenterRule"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCostCenterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改财务单元规则
//
// @param request - ModifyCostCenterRuleRequest
//
// @return ModifyCostCenterRuleResponse
func (client *Client) ModifyCostCenterRule(request *ModifyCostCenterRuleRequest) (_result *ModifyCostCenterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCostCenterRuleResponse{}
	_body, _err := client.ModifyCostCenterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询财务单元
//
// @param tmpReq - QueryCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterResponse
func (client *Client) QueryCostCenterWithOptions(tmpReq *QueryCostCenterRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QueryCostCenterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.OwnerAccountId) {
		query["OwnerAccountId"] = request.OwnerAccountId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentCostCenterId) {
		query["ParentCostCenterId"] = request.ParentCostCenterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCostCenter"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCostCenterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询财务单元
//
// @param request - QueryCostCenterRequest
//
// @return QueryCostCenterResponse
func (client *Client) QueryCostCenter(request *QueryCostCenterRequest) (_result *QueryCostCenterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCostCenterResponse{}
	_body, _err := client.QueryCostCenterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询财务单元下资源信息
//
// @param request - QueryCostCenterResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterResourceResponse
func (client *Client) QueryCostCenterResourceWithOptions(request *QueryCostCenterResourceRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIds) {
		query["EcIdAccountIds"] = request.EcIdAccountIds
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CostCenterId) {
		body["CostCenterId"] = request.CostCenterId
	}

	if !dara.IsNil(request.OwnerAccountId) {
		body["OwnerAccountId"] = request.OwnerAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCostCenterResource"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCostCenterResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询财务单元下资源信息
//
// @param request - QueryCostCenterResourceRequest
//
// @return QueryCostCenterResourceResponse
func (client *Client) QueryCostCenterResource(request *QueryCostCenterResourceRequest) (_result *QueryCostCenterResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCostCenterResourceResponse{}
	_body, _err := client.QueryCostCenterResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询财务单元规则
//
// @param request - QueryCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterRuleResponse
func (client *Client) QueryCostCenterRuleWithOptions(request *QueryCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIds) {
		query["EcIdAccountIds"] = request.EcIdAccountIds
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CostCenterId) {
		body["CostCenterId"] = request.CostCenterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCostCenterRule"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCostCenterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询财务单元规则
//
// @param request - QueryCostCenterRuleRequest
//
// @return QueryCostCenterRuleResponse
func (client *Client) QueryCostCenterRule(request *QueryCostCenterRuleRequest) (_result *QueryCostCenterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCostCenterRuleResponse{}
	_body, _err := client.QueryCostCenterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置资金账户的信控限额
//
// @param request - SetFundAccountCreditAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFundAccountCreditAmountResponse
func (client *Client) SetFundAccountCreditAmountWithOptions(request *SetFundAccountCreditAmountRequest, runtime *dara.RuntimeOptions) (_result *SetFundAccountCreditAmountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreditAmount) {
		body["CreditAmount"] = request.CreditAmount
	}

	if !dara.IsNil(request.Currency) {
		body["Currency"] = request.Currency
	}

	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetFundAccountCreditAmount"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetFundAccountCreditAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置资金账户的信控限额
//
// @param request - SetFundAccountCreditAmountRequest
//
// @return SetFundAccountCreditAmountResponse
func (client *Client) SetFundAccountCreditAmount(request *SetFundAccountCreditAmountRequest) (_result *SetFundAccountCreditAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetFundAccountCreditAmountResponse{}
	_body, _err := client.SetFundAccountCreditAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置资金账户低额预警
//
// @param request - SetFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) SetFundAccountLowAvailableAmountAlarmWithOptions(request *SetFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *SetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FundAccountId) {
		body["FundAccountId"] = request.FundAccountId
	}

	if !dara.IsNil(request.ThresholdAmount) {
		body["ThresholdAmount"] = request.ThresholdAmount
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetFundAccountLowAvailableAmountAlarm"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置资金账户低额预警
//
// @param request - SetFundAccountLowAvailableAmountAlarmRequest
//
// @return SetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) SetFundAccountLowAvailableAmountAlarm(request *SetFundAccountLowAvailableAmountAlarmRequest) (_result *SetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetFundAccountLowAvailableAmountAlarmResponse{}
	_body, _err := client.SetFundAccountLowAvailableAmountAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置节省计划用户级抵扣规则
//
// @param tmpReq - SetSavingPlanUserDeductRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSavingPlanUserDeductRuleResponse
func (client *Client) SetSavingPlanUserDeductRuleWithOptions(tmpReq *SetSavingPlanUserDeductRuleRequest, runtime *dara.RuntimeOptions) (_result *SetSavingPlanUserDeductRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetSavingPlanUserDeductRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserDeductRules) {
		request.UserDeductRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserDeductRules, dara.String("UserDeductRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SpnInstanceCode) {
		body["SpnInstanceCode"] = request.SpnInstanceCode
	}

	if !dara.IsNil(request.UserDeductRulesShrink) {
		body["UserDeductRules"] = request.UserDeductRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSavingPlanUserDeductRule"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSavingPlanUserDeductRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置节省计划用户级抵扣规则
//
// @param request - SetSavingPlanUserDeductRuleRequest
//
// @return SetSavingPlanUserDeductRuleResponse
func (client *Client) SetSavingPlanUserDeductRule(request *SetSavingPlanUserDeductRuleRequest) (_result *SetSavingPlanUserDeductRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSavingPlanUserDeductRuleResponse{}
	_body, _err := client.SetSavingPlanUserDeductRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
