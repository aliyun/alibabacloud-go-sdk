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
// 作废优惠券
//
// @param request - CancelCouponRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCouponResponse
func (client *Client) CancelCouponWithOptions(request *CancelCouponRequest, runtime *dara.RuntimeOptions) (_result *CancelCouponResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponId) {
		query["CouponId"] = request.CouponId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCoupon"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCouponResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作废优惠券
//
// @param request - CancelCouponRequest
//
// @return CancelCouponResponse
func (client *Client) CancelCoupon(request *CancelCouponRequest) (_result *CancelCouponResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCouponResponse{}
	_body, _err := client.CancelCouponWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the subscription to multi-level bills as an Alibaba Cloud eco-partner.
//
// Description:
//
// Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
//
// You can call this operation to cancel the subscription to only one type of bill at a time.
//
// After the subscription to a type of bill is canceled, bills of this type are no longer pushed to the specified Object Storage Service (OSS) bucket.
//
// **This topic is published only on the international site (alibabacloud.com).
//
// @param request - CancelSubscriptionBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSubscriptionBillResponse
func (client *Client) CancelSubscriptionBillWithOptions(request *CancelSubscriptionBillRequest, runtime *dara.RuntimeOptions) (_result *CancelSubscriptionBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SubscribeType) {
		query["SubscribeType"] = request.SubscribeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelSubscriptionBill"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelSubscriptionBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the subscription to multi-level bills as an Alibaba Cloud eco-partner.
//
// Description:
//
// Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
//
// You can call this operation to cancel the subscription to only one type of bill at a time.
//
// After the subscription to a type of bill is canceled, bills of this type are no longer pushed to the specified Object Storage Service (OSS) bucket.
//
// **This topic is published only on the international site (alibabacloud.com).
//
// @param request - CancelSubscriptionBillRequest
//
// @return CancelSubscriptionBillResponse
func (client *Client) CancelSubscriptionBill(request *CancelSubscriptionBillRequest) (_result *CancelSubscriptionBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelSubscriptionBillResponse{}
	_body, _err := client.CancelSubscriptionBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 优惠券审批状态列表
//
// @param request - CouponApprovalStatusListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CouponApprovalStatusListResponse
func (client *Client) CouponApprovalStatusListWithOptions(request *CouponApprovalStatusListRequest, runtime *dara.RuntimeOptions) (_result *CouponApprovalStatusListResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateStatus) {
		query["TemplateStatus"] = request.TemplateStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CouponApprovalStatusList"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CouponApprovalStatusListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 优惠券审批状态列表
//
// @param request - CouponApprovalStatusListRequest
//
// @return CouponApprovalStatusListResponse
func (client *Client) CouponApprovalStatusList(request *CouponApprovalStatusListRequest) (_result *CouponApprovalStatusListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CouponApprovalStatusListResponse{}
	_body, _err := client.CouponApprovalStatusListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建优惠券模板
//
// @param tmpReq - CreateCouponTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCouponTemplateResponse
func (client *Client) CreateCouponTemplateWithOptions(tmpReq *CreateCouponTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateCouponTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCouponTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProductType) {
		request.ProductTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProductType, dara.String("ProductType"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ApplicableProducts) {
		query["ApplicableProducts"] = request.ApplicableProducts
	}

	if !dara.IsNil(request.CostBearer) {
		query["CostBearer"] = request.CostBearer
	}

	if !dara.IsNil(request.CouponDescription) {
		query["CouponDescription"] = request.CouponDescription
	}

	if !dara.IsNil(request.Expireddate) {
		query["Expireddate"] = request.Expireddate
	}

	if !dara.IsNil(request.LimitPerPerson) {
		query["LimitPerPerson"] = request.LimitPerPerson
	}

	if !dara.IsNil(request.ProductTypeShrink) {
		query["ProductType"] = request.ProductTypeShrink
	}

	if !dara.IsNil(request.PurchaseType) {
		query["PurchaseType"] = request.PurchaseType
	}

	if !dara.IsNil(request.ReasonForApplication) {
		query["ReasonForApplication"] = request.ReasonForApplication
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.Vailddate) {
		query["Vailddate"] = request.Vailddate
	}

	if !dara.IsNil(request.Vaildperioddays) {
		query["Vaildperioddays"] = request.Vaildperioddays
	}

	if !dara.IsNil(request.ValidUntil) {
		query["ValidUntil"] = request.ValidUntil
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCouponTemplate"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCouponTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建优惠券模板
//
// @param request - CreateCouponTemplateRequest
//
// @return CreateCouponTemplateResponse
func (client *Client) CreateCouponTemplate(request *CreateCouponTemplateRequest) (_result *CreateCouponTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCouponTemplateResponse{}
	_body, _err := client.CreateCouponTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This function is designed for create a customer who is to be invited.
//
// @param request - CreateCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerResponse
func (client *Client) CreateCustomerWithOptions(request *CreateCustomerRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerName) {
		query["CustomerName"] = request.CustomerName
	}

	if !dara.IsNil(request.CustomerSource) {
		query["CustomerSource"] = request.CustomerSource
	}

	if !dara.IsNil(request.CustomerSubTrade) {
		query["CustomerSubTrade"] = request.CustomerSubTrade
	}

	if !dara.IsNil(request.CustomerTrade) {
		query["CustomerTrade"] = request.CustomerTrade
	}

	if !dara.IsNil(request.Nation) {
		query["Nation"] = request.Nation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomer"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This function is designed for create a customer who is to be invited.
//
// @param request - CreateCustomerRequest
//
// @return CreateCustomerResponse
func (client *Client) CreateCustomer(request *CreateCustomerRequest) (_result *CreateCustomerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CreateCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query quota adjustment list of Distribution Customer from International Site. Not available on Domestic Site.
//
// @param request - CustomerQuotaRecordListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerQuotaRecordListResponse
func (client *Client) CustomerQuotaRecordListWithOptions(request *CustomerQuotaRecordListRequest, runtime *dara.RuntimeOptions) (_result *CustomerQuotaRecordListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CustomerQuotaRecordList"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CustomerQuotaRecordListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query quota adjustment list of Distribution Customer from International Site. Not available on Domestic Site.
//
// @param request - CustomerQuotaRecordListRequest
//
// @return CustomerQuotaRecordListResponse
func (client *Client) CustomerQuotaRecordList(request *CustomerQuotaRecordListRequest) (_result *CustomerQuotaRecordListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CustomerQuotaRecordListResponse{}
	_body, _err := client.CustomerQuotaRecordListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API is used to offset the Deducted Credit of a Distribution Customer. For example, if the current Deducted Credit is 500 and the Available Credit is 1000, by offsetting 300, the Deducted Credit will then become 200, and the Available Credit becomes 1300.
//
// Description:
//
// Note that sometimes you may find that the customer\\"s Used Credit is negative. This indicates that there is no need to restore the Used Credit, and its ready for customer\\"s usage. This phenomenon occurs because a refund is generated while the customer\\"s credit is full, thereby triggered additional increasing on the customer\\"s credit.
//
// For example, if the customer\\"s maximum Available Credit is 1000 with no usage, and a refund of 300 occurs, the Used Credit will become -300.
//
// @param request - DeductOutstandingBalanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeductOutstandingBalanceResponse
func (client *Client) DeductOutstandingBalanceWithOptions(request *DeductOutstandingBalanceRequest, runtime *dara.RuntimeOptions) (_result *DeductOutstandingBalanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeductAmount) {
		query["DeductAmount"] = request.DeductAmount
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeductOutstandingBalance"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeductOutstandingBalanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is used to offset the Deducted Credit of a Distribution Customer. For example, if the current Deducted Credit is 500 and the Available Credit is 1000, by offsetting 300, the Deducted Credit will then become 200, and the Available Credit becomes 1300.
//
// Description:
//
// Note that sometimes you may find that the customer\\"s Used Credit is negative. This indicates that there is no need to restore the Used Credit, and its ready for customer\\"s usage. This phenomenon occurs because a refund is generated while the customer\\"s credit is full, thereby triggered additional increasing on the customer\\"s credit.
//
// For example, if the customer\\"s maximum Available Credit is 1000 with no usage, and a refund of 300 occurs, the Used Credit will become -300.
//
// @param request - DeductOutstandingBalanceRequest
//
// @return DeductOutstandingBalanceResponse
func (client *Client) DeductOutstandingBalance(request *DeductOutstandingBalanceRequest) (_result *DeductOutstandingBalanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeductOutstandingBalanceResponse{}
	_body, _err := client.DeductOutstandingBalanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 作废优惠券模板
//
// @param request - DeleteCouponTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCouponTemplateResponse
func (client *Client) DeleteCouponTemplateWithOptions(request *DeleteCouponTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCouponTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCouponTemplate"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCouponTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作废优惠券模板
//
// @param request - DeleteCouponTemplateRequest
//
// @return DeleteCouponTemplateResponse
func (client *Client) DeleteCouponTemplate(request *DeleteCouponTemplateRequest) (_result *DeleteCouponTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCouponTemplateResponse{}
	_body, _err := client.DeleteCouponTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Set the after-shutdown instance status for post-pay End Users as a Reseller.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditEndUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditEndUserStatusResponse
func (client *Client) EditEndUserStatusWithOptions(request *EditEndUserStatusRequest, runtime *dara.RuntimeOptions) (_result *EditEndUserStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditEndUserStatus"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditEndUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set the after-shutdown instance status for post-pay End Users as a Reseller.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditEndUserStatusRequest
//
// @return EditEndUserStatusResponse
func (client *Client) EditEndUserStatus(request *EditEndUserStatusRequest) (_result *EditEndUserStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditEndUserStatusResponse{}
	_body, _err := client.EditEndUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Set the New Buy status for Sub-Customer as a Partner.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditNewBuyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditNewBuyStatusResponse
func (client *Client) EditNewBuyStatusWithOptions(request *EditNewBuyStatusRequest, runtime *dara.RuntimeOptions) (_result *EditNewBuyStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewBuyStatus) {
		query["NewBuyStatus"] = request.NewBuyStatus
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditNewBuyStatus"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditNewBuyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set the New Buy status for Sub-Customer as a Partner.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditNewBuyStatusRequest
//
// @return EditNewBuyStatusResponse
func (client *Client) EditNewBuyStatus(request *EditNewBuyStatusRequest) (_result *EditNewBuyStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditNewBuyStatusResponse{}
	_body, _err := client.EditNewBuyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the End User\\"s Shutdown Policy as a Reseller.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditZeroCreditShutdownRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditZeroCreditShutdownResponse
func (client *Client) EditZeroCreditShutdownWithOptions(request *EditZeroCreditShutdownRequest, runtime *dara.RuntimeOptions) (_result *EditZeroCreditShutdownResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ShutdownPolicy) {
		query["ShutdownPolicy"] = request.ShutdownPolicy
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditZeroCreditShutdown"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditZeroCreditShutdownResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the End User\\"s Shutdown Policy as a Reseller.
//
// Description:
//
// The caller should be the Partner as identified in the Alibaba Cloud distribution model. </br>
//
// **This content is only published on the international site. **
//
// @param request - EditZeroCreditShutdownRequest
//
// @return EditZeroCreditShutdownResponse
func (client *Client) EditZeroCreditShutdown(request *EditZeroCreditShutdownRequest) (_result *EditZeroCreditShutdownResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditZeroCreditShutdownResponse{}
	_body, _err := client.EditZeroCreditShutdownWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Export quota amount adjustment history as a Distribution Customer from International Site. Only available on International Site.
//
// Description:
//
// Caller must be a Partner from International Site, either Distribution or Reseller will do.
//
// @param request - ExportCustomerQuotaRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportCustomerQuotaRecordResponse
func (client *Client) ExportCustomerQuotaRecordWithOptions(request *ExportCustomerQuotaRecordRequest, runtime *dara.RuntimeOptions) (_result *ExportCustomerQuotaRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.EndUserPk) {
		query["EndUserPk"] = request.EndUserPk
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportCustomerQuotaRecord"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportCustomerQuotaRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Export quota amount adjustment history as a Distribution Customer from International Site. Only available on International Site.
//
// Description:
//
// Caller must be a Partner from International Site, either Distribution or Reseller will do.
//
// @param request - ExportCustomerQuotaRecordRequest
//
// @return ExportCustomerQuotaRecordResponse
func (client *Client) ExportCustomerQuotaRecord(request *ExportCustomerQuotaRecordRequest) (_result *ExportCustomerQuotaRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportCustomerQuotaRecordResponse{}
	_body, _err := client.ExportCustomerQuotaRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 额度冲减明细列表导出接口
//
// @param request - ExportReversedDeductionHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportReversedDeductionHistoryResponse
func (client *Client) ExportReversedDeductionHistoryWithOptions(request *ExportReversedDeductionHistoryRequest, runtime *dara.RuntimeOptions) (_result *ExportReversedDeductionHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ExportUid) {
		query["ExportUid"] = request.ExportUid
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportReversedDeductionHistory"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportReversedDeductionHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 额度冲减明细列表导出接口
//
// @param request - ExportReversedDeductionHistoryRequest
//
// @return ExportReversedDeductionHistoryResponse
func (client *Client) ExportReversedDeductionHistory(request *ExportReversedDeductionHistoryRequest) (_result *ExportReversedDeductionHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportReversedDeductionHistoryResponse{}
	_body, _err := client.ExportReversedDeductionHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Return Distribution Customer\\"s account information.
//
// @param request - GetAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfoWithOptions(request *GetAccountInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAccountInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountInfo"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Return Distribution Customer\\"s account information.
//
// @param request - GetAccountInfoRequest
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfo(request *GetAccountInfoRequest) (_result *GetAccountInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.GetAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提供返佣商品API
//
// @param tmpReq - GetCommissionableProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommissionableProductsResponse
func (client *Client) GetCommissionableProductsWithOptions(tmpReq *GetCommissionableProductsRequest, runtime *dara.RuntimeOptions) (_result *GetCommissionableProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCommissionableProductsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ListShowStatusList) {
		request.ListShowStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ListShowStatusList, dara.String("ListShowStatusList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCodeList) {
		query["CommodityCodeList"] = request.CommodityCodeList
	}

	if !dara.IsNil(request.FiscalYear) {
		query["FiscalYear"] = request.FiscalYear
	}

	if !dara.IsNil(request.ListShowStatusListShrink) {
		query["ListShowStatusList"] = request.ListShowStatusListShrink
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PipCodeList) {
		query["PipCodeList"] = request.PipCodeList
	}

	if !dara.IsNil(request.RealEndMonth) {
		query["RealEndMonth"] = request.RealEndMonth
	}

	if !dara.IsNil(request.RealStartMonth) {
		query["RealStartMonth"] = request.RealStartMonth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCommissionableProducts"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCommissionableProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提供返佣商品API
//
// @param request - GetCommissionableProductsRequest
//
// @return GetCommissionableProductsResponse
func (client *Client) GetCommissionableProducts(request *GetCommissionableProductsRequest) (_result *GetCommissionableProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCommissionableProductsResponse{}
	_body, _err := client.GetCommissionableProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询优惠券模板详情
//
// @param request - GetCouponTemplateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCouponTemplateDetailResponse
func (client *Client) GetCouponTemplateDetailWithOptions(request *GetCouponTemplateDetailRequest, runtime *dara.RuntimeOptions) (_result *GetCouponTemplateDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCouponTemplateDetail"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCouponTemplateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询优惠券模板详情
//
// @param request - GetCouponTemplateDetailRequest
//
// @return GetCouponTemplateDetailResponse
func (client *Client) GetCouponTemplateDetail(request *GetCouponTemplateDetailRequest) (_result *GetCouponTemplateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCouponTemplateDetailResponse{}
	_body, _err := client.GetCouponTemplateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 国际渠道分销优惠券可抵扣产品
//
// @param request - GetCoupondeductProductCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCoupondeductProductCodeResponse
func (client *Client) GetCoupondeductProductCodeWithOptions(request *GetCoupondeductProductCodeRequest, runtime *dara.RuntimeOptions) (_result *GetCoupondeductProductCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCoupondeductProductCode"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCoupondeductProductCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际渠道分销优惠券可抵扣产品
//
// @param request - GetCoupondeductProductCodeRequest
//
// @return GetCoupondeductProductCodeResponse
func (client *Client) GetCoupondeductProductCode(request *GetCoupondeductProductCodeRequest) (_result *GetCoupondeductProductCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCoupondeductProductCodeResponse{}
	_body, _err := client.GetCoupondeductProductCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Credit Control information of Distribution Customers. The PopCreditInfoJson in the Return Parameter will be empty if the Distribution Customer is an Agency. This function is only available for Resellers and Distributors.
//
// @param request - GetCreditInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreditInfoResponse
func (client *Client) GetCreditInfoWithOptions(request *GetCreditInfoRequest, runtime *dara.RuntimeOptions) (_result *GetCreditInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCreditInfo"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCreditInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Credit Control information of Distribution Customers. The PopCreditInfoJson in the Return Parameter will be empty if the Distribution Customer is an Agency. This function is only available for Resellers and Distributors.
//
// @param request - GetCreditInfoRequest
//
// @return GetCreditInfoResponse
func (client *Client) GetCreditInfo(request *GetCreditInfoRequest) (_result *GetCreditInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCreditInfoResponse{}
	_body, _err := client.GetCreditInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 客户订单查询
//
// @param request - GetCustomerOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerOrdersResponse
func (client *Client) GetCustomerOrdersWithOptions(request *GetCustomerOrdersRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerOrders"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 客户订单查询
//
// @param request - GetCustomerOrdersRequest
//
// @return GetCustomerOrdersResponse
func (client *Client) GetCustomerOrders(request *GetCustomerOrdersRequest) (_result *GetCustomerOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomerOrdersResponse{}
	_body, _err := client.GetCustomerOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issue Distributor\\"s daily Bill. This function is only available for Resellers and Distributors.
//
// @param request - GetDailyBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDailyBillResponse
func (client *Client) GetDailyBillWithOptions(request *GetDailyBillRequest, runtime *dara.RuntimeOptions) (_result *GetDailyBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwner) {
		query["BillOwner"] = request.BillOwner
	}

	if !dara.IsNil(request.BillType) {
		query["BillType"] = request.BillType
	}

	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDailyBill"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDailyBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issue Distributor\\"s daily Bill. This function is only available for Resellers and Distributors.
//
// @param request - GetDailyBillRequest
//
// @return GetDailyBillResponse
func (client *Client) GetDailyBill(request *GetDailyBillRequest) (_result *GetDailyBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDailyBillResponse{}
	_body, _err := client.GetDailyBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query invitation status of customer who have been created and invited.
//
// @param request - GetInviteStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInviteStatusResponse
func (client *Client) GetInviteStatusWithOptions(request *GetInviteStatusRequest, runtime *dara.RuntimeOptions) (_result *GetInviteStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InviteStatusList) {
		query["InviteStatusList"] = request.InviteStatusList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInviteStatus"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInviteStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query invitation status of customer who have been created and invited.
//
// @param request - GetInviteStatusRequest
//
// @return GetInviteStatusResponse
func (client *Client) GetInviteStatus(request *GetInviteStatusRequest) (_result *GetInviteStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInviteStatusResponse{}
	_body, _err := client.GetInviteStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Issue Distributor\\"s Monthly Bill. This function is only available for Resellers and Distributors.
//
// @param request - GetMonthlyBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonthlyBillResponse
func (client *Client) GetMonthlyBillWithOptions(request *GetMonthlyBillRequest, runtime *dara.RuntimeOptions) (_result *GetMonthlyBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwner) {
		query["BillOwner"] = request.BillOwner
	}

	if !dara.IsNil(request.BillType) {
		query["BillType"] = request.BillType
	}

	if !dara.IsNil(request.Month) {
		query["Month"] = request.Month
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMonthlyBill"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMonthlyBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issue Distributor\\"s Monthly Bill. This function is only available for Resellers and Distributors.
//
// @param request - GetMonthlyBillRequest
//
// @return GetMonthlyBillResponse
func (client *Client) GetMonthlyBill(request *GetMonthlyBillRequest) (_result *GetMonthlyBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMonthlyBillResponse{}
	_body, _err := client.GetMonthlyBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下单控制记录查询
//
// @param request - GetPurchaseControlRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurchaseControlRecordResponse
func (client *Client) GetPurchaseControlRecordWithOptions(request *GetPurchaseControlRecordRequest, runtime *dara.RuntimeOptions) (_result *GetPurchaseControlRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerUID) {
		query["CustomerUID"] = request.CustomerUID
	}

	if !dara.IsNil(request.OperationTime) {
		query["OperationTime"] = request.OperationTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPurchaseControlRecord"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPurchaseControlRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下单控制记录查询
//
// @param request - GetPurchaseControlRecordRequest
//
// @return GetPurchaseControlRecordResponse
func (client *Client) GetPurchaseControlRecord(request *GetPurchaseControlRecordRequest) (_result *GetPurchaseControlRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPurchaseControlRecordResponse{}
	_body, _err := client.GetPurchaseControlRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询延停策略修改记录
//
// @param request - GetShutdownPolicyRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShutdownPolicyRecordResponse
func (client *Client) GetShutdownPolicyRecordWithOptions(request *GetShutdownPolicyRecordRequest, runtime *dara.RuntimeOptions) (_result *GetShutdownPolicyRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerUID) {
		query["CustomerUID"] = request.CustomerUID
	}

	if !dara.IsNil(request.OperationTime) {
		query["OperationTime"] = request.OperationTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetShutdownPolicyRecord"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetShutdownPolicyRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询延停策略修改记录
//
// @param request - GetShutdownPolicyRecordRequest
//
// @return GetShutdownPolicyRecordResponse
func (client *Client) GetShutdownPolicyRecord(request *GetShutdownPolicyRecordRequest) (_result *GetShutdownPolicyRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetShutdownPolicyRecordResponse{}
	_body, _err := client.GetShutdownPolicyRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query all the Unassociated Customer.
//
// @param request - GetUnassociatedCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUnassociatedCustomerResponse
func (client *Client) GetUnassociatedCustomerWithOptions(request *GetUnassociatedCustomerRequest, runtime *dara.RuntimeOptions) (_result *GetUnassociatedCustomerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUnassociatedCustomer"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUnassociatedCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query all the Unassociated Customer.
//
// @param request - GetUnassociatedCustomerRequest
//
// @return GetUnassociatedCustomerResponse
func (client *Client) GetUnassociatedCustomer(request *GetUnassociatedCustomerRequest) (_result *GetUnassociatedCustomerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUnassociatedCustomerResponse{}
	_body, _err := client.GetUnassociatedCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiate the Partner registration invitation.
//
// Description:
//
// The current API request rate for the Cloud Product has not been disclosed.
//
// @param request - InviteSubAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteSubAccountResponse
func (client *Client) InviteSubAccountWithOptions(request *InviteSubAccountRequest, runtime *dara.RuntimeOptions) (_result *InviteSubAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountInfoList) {
		query["AccountInfoList"] = request.AccountInfoList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InviteSubAccount"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InviteSubAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiate the Partner registration invitation.
//
// Description:
//
// The current API request rate for the Cloud Product has not been disclosed.
//
// @param request - InviteSubAccountRequest
//
// @return InviteSubAccountResponse
func (client *Client) InviteSubAccount(request *InviteSubAccountRequest) (_result *InviteSubAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InviteSubAccountResponse{}
	_body, _err := client.InviteSubAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发放优惠券
//
// @param request - IssueCouponForCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IssueCouponForCustomerResponse
func (client *Client) IssueCouponForCustomerWithOptions(request *IssueCouponForCustomerRequest, runtime *dara.RuntimeOptions) (_result *IssueCouponForCustomerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CouponTemplateId) {
		query["CouponTemplateId"] = request.CouponTemplateId
	}

	if !dara.IsNil(request.IsUseBenefit) {
		query["IsUseBenefit"] = request.IsUseBenefit
	}

	if !dara.IsNil(request.Uidlist) {
		query["Uidlist"] = request.Uidlist
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IssueCouponForCustomer"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IssueCouponForCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发放优惠券
//
// @param request - IssueCouponForCustomerRequest
//
// @return IssueCouponForCustomerResponse
func (client *Client) IssueCouponForCustomer(request *IssueCouponForCustomerRequest) (_result *IssueCouponForCustomerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IssueCouponForCustomerResponse{}
	_body, _err := client.IssueCouponForCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This function is available for all Distributors. It displays the corresponding region code information based on the operable countries as agreed in the Distributor\\"s contract.
//
// Description:
//
// The current API request rate for cloud products has not been disclosed.
//
// @param request - ListCountriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCountriesResponse
func (client *Client) ListCountriesWithOptions(runtime *dara.RuntimeOptions) (_result *ListCountriesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListCountries"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCountriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This function is available for all Distributors. It displays the corresponding region code information based on the operable countries as agreed in the Distributor\\"s contract.
//
// Description:
//
// The current API request rate for cloud products has not been disclosed.
//
// @return ListCountriesResponse
func (client *Client) ListCountries() (_result *ListCountriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCountriesResponse{}
	_body, _err := client.ListCountriesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 优惠券使用量列表查询
//
// @param request - ListCouponUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCouponUsageResponse
func (client *Client) ListCouponUsageWithOptions(request *ListCouponUsageRequest, runtime *dara.RuntimeOptions) (_result *ListCouponUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.CouponTemplateId) {
		query["CouponTemplateId"] = request.CouponTemplateId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCouponUsage"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCouponUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 优惠券使用量列表查询
//
// @param request - ListCouponUsageRequest
//
// @return ListCouponUsageResponse
func (client *Client) ListCouponUsage(request *ListCouponUsageRequest) (_result *ListCouponUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCouponUsageResponse{}
	_body, _err := client.ListCouponUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用查询导出任务列表
//
// @param request - ListExportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExportTasksResponse
func (client *Client) ListExportTasksWithOptions(request *ListExportTasksRequest, runtime *dara.RuntimeOptions) (_result *ListExportTasksResponse, _err error) {
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

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExportTasks"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExportTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用查询导出任务列表
//
// @param request - ListExportTasksRequest
//
// @return ListExportTasksResponse
func (client *Client) ListExportTasks(request *ListExportTasksRequest) (_result *ListExportTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExportTasksResponse{}
	_body, _err := client.ListExportTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 额度冲减明细列表
//
// @param request - QueryReversedDeductionHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReversedDeductionHistoryResponse
func (client *Client) QueryReversedDeductionHistoryWithOptions(request *QueryReversedDeductionHistoryRequest, runtime *dara.RuntimeOptions) (_result *QueryReversedDeductionHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryReversedDeductionHistory"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReversedDeductionHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 额度冲减明细列表
//
// @param request - QueryReversedDeductionHistoryRequest
//
// @return QueryReversedDeductionHistoryResponse
func (client *Client) QueryReversedDeductionHistory(request *QueryReversedDeductionHistoryRequest) (_result *QueryReversedDeductionHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryReversedDeductionHistoryResponse{}
	_body, _err := client.QueryReversedDeductionHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check the result of export quota list as a Distribution Customer from International Site. Only available on International Site.
//
// Description:
//
// Caller must be a Partner from International Site, either Distribution or Reseller will do.
//
// @param request - QuotaListExportPagedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuotaListExportPagedResponse
func (client *Client) QuotaListExportPagedWithOptions(request *QuotaListExportPagedRequest, runtime *dara.RuntimeOptions) (_result *QuotaListExportPagedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuotaListExportPaged"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuotaListExportPagedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check the result of export quota list as a Distribution Customer from International Site. Only available on International Site.
//
// Description:
//
// Caller must be a Partner from International Site, either Distribution or Reseller will do.
//
// @param request - QuotaListExportPagedRequest
//
// @return QuotaListExportPagedResponse
func (client *Client) QuotaListExportPaged(request *QuotaListExportPagedRequest) (_result *QuotaListExportPagedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuotaListExportPagedResponse{}
	_body, _err := client.QuotaListExportPagedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resend invitation email.
//
// @param request - ResendEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendEmailResponse
func (client *Client) ResendEmailWithOptions(request *ResendEmailRequest, runtime *dara.RuntimeOptions) (_result *ResendEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InviteId) {
		query["InviteId"] = request.InviteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResendEmail"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResendEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resend invitation email.
//
// @param request - ResendEmailRequest
//
// @return ResendEmailResponse
func (client *Client) ResendEmail(request *ResendEmailRequest) (_result *ResendEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResendEmailResponse{}
	_body, _err := client.ResendEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This function is designed for Sub Account information maintenance, including Nickname and Remark.
//
// @param request - SetAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAccountInfoResponse
func (client *Client) SetAccountInfoWithOptions(request *SetAccountInfoRequest, runtime *dara.RuntimeOptions) (_result *SetAccountInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountNickname) {
		query["AccountNickname"] = request.AccountNickname
	}

	if !dara.IsNil(request.CustomerBd) {
		query["CustomerBd"] = request.CustomerBd
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAccountInfo"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This function is designed for Sub Account information maintenance, including Nickname and Remark.
//
// @param request - SetAccountInfoRequest
//
// @return SetAccountInfoResponse
func (client *Client) SetAccountInfo(request *SetAccountInfoRequest) (_result *SetAccountInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAccountInfoResponse{}
	_body, _err := client.SetAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Set Credit Line for Distribution Customers. This function is only available for Resellers and Distributors.
//
// @param request - SetCreditLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCreditLineResponse
func (client *Client) SetCreditLineWithOptions(request *SetCreditLineRequest, runtime *dara.RuntimeOptions) (_result *SetCreditLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreditLine) {
		query["CreditLine"] = request.CreditLine
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCreditLine"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCreditLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Set Credit Line for Distribution Customers. This function is only available for Resellers and Distributors.
//
// @param request - SetCreditLineRequest
//
// @return SetCreditLineResponse
func (client *Client) SetCreditLine(request *SetCreditLineRequest) (_result *SetCreditLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCreditLineResponse{}
	_body, _err := client.SetCreditLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can use this API to set the threshold for the use of credit control. When the customer credit control reaches below the threshold, it will pass through the notification email distributor. This feature is for Reseller and Distributor only.
//
// @param request - SetWarningThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWarningThresholdResponse
func (client *Client) SetWarningThresholdWithOptions(request *SetWarningThresholdRequest, runtime *dara.RuntimeOptions) (_result *SetWarningThresholdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	if !dara.IsNil(request.WarningValue) {
		query["WarningValue"] = request.WarningValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWarningThreshold"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWarningThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can use this API to set the threshold for the use of credit control. When the customer credit control reaches below the threshold, it will pass through the notification email distributor. This feature is for Reseller and Distributor only.
//
// @param request - SetWarningThresholdRequest
//
// @return SetWarningThresholdResponse
func (client *Client) SetWarningThreshold(request *SetWarningThresholdRequest) (_result *SetWarningThresholdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetWarningThresholdResponse{}
	_body, _err := client.SetWarningThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates the subscription to multi-level bills as an Alibaba Cloud eco-partner.
//
// Description:
//
//	  Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
//
//		- You can call this operation to subscribe to only one type of bill at a time.
//
//		- After the subscription to a type of bill is generated, the bill for the previous day is pushed on a daily basis from the next day. On the fifth day of each month, the full-data bill for the previous month is pushed.
//
//		- A daily bill may be delayed. The delayed bill is pushed the next day after it is generated. The delayed bill may contain the bill data that is delayed until the previous day. We recommend that you query the full-data bill for the previous month at the beginning of each month.
//
//		- Your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
//
//		- The following file name formats are supported for bills:
//
// ```
//
// # BillingItemDetailForBillingPeriod
//
// File name format of a daily bill: UID_PartnerBillingItemDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_BillingItemDetail_20190310_0001_01.
//
// File name format of a monthly full-data bill: UID_PartnerBillingItemDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetail_201903_0001_01.
//
// # InstanceDetailForBillingPeriod
//
//	File name format of a daily bill: UID_PartnerInstanceDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_InstanceDetail_20190310_0001_01.
//
// File name format of a monthly full-data bill: UID_PartnerInstanceDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetail_201903_1999-0001_01.
//
// # BillingItemDetailMonthly
//
// File name format of a daily bill: UID_PartnerBillingItemDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
//
// # InstanceDetailMonthly
//
// File name format of a daily bill: UID_PartnerInstanceDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
//
// The fileNo field exists only when the number of bill rows reaches the maximum rows in a single bill file and the bill is split into multiple files.
//
// ```
//
// **This topic is published only on the international site (alibabacloud.com).
//
// @param request - SubscriptionBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscriptionBillResponse
func (client *Client) SubscriptionBillWithOptions(request *SubscriptionBillRequest, runtime *dara.RuntimeOptions) (_result *SubscriptionBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginBillingCycle) {
		query["BeginBillingCycle"] = request.BeginBillingCycle
	}

	if !dara.IsNil(request.BillFormat) {
		query["BillFormat"] = request.BillFormat
	}

	if !dara.IsNil(request.BucketOwnerId) {
		query["BucketOwnerId"] = request.BucketOwnerId
	}

	if !dara.IsNil(request.SubscribeBucket) {
		query["SubscribeBucket"] = request.SubscribeBucket
	}

	if !dara.IsNil(request.SubscribeSegmentSize) {
		query["SubscribeSegmentSize"] = request.SubscribeSegmentSize
	}

	if !dara.IsNil(request.SubscribeType) {
		query["SubscribeType"] = request.SubscribeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubscriptionBill"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubscriptionBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates the subscription to multi-level bills as an Alibaba Cloud eco-partner.
//
// Description:
//
//	  Make sure that you are a distributor of the Alibaba Cloud international ecosystem.
//
//		- You can call this operation to subscribe to only one type of bill at a time.
//
//		- After the subscription to a type of bill is generated, the bill for the previous day is pushed on a daily basis from the next day. On the fifth day of each month, the full-data bill for the previous month is pushed.
//
//		- A daily bill may be delayed. The delayed bill is pushed the next day after it is generated. The delayed bill may contain the bill data that is delayed until the previous day. We recommend that you query the full-data bill for the previous month at the beginning of each month.
//
//		- Your account must be granted the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
//
//		- The following file name formats are supported for bills:
//
// ```
//
// # BillingItemDetailForBillingPeriod
//
// File name format of a daily bill: UID_PartnerBillingItemDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_BillingItemDetail_20190310_0001_01.
//
// File name format of a monthly full-data bill: UID_PartnerBillingItemDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetail_201903_0001_01.
//
// # InstanceDetailForBillingPeriod
//
//	File name format of a daily bill: UID_PartnerInstanceDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_InstanceDetail_20190310_0001_01.
//
// File name format of a monthly full-data bill: UID_PartnerInstanceDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetail_201903_1999-0001_01.
//
// # BillingItemDetailMonthly
//
// File name format of a daily bill: UID_PartnerBillingItemDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
//
// # InstanceDetailMonthly
//
// File name format of a daily bill: UID_PartnerInstanceDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetailMonthly_201903_0001_01. This bill contains the bill data that is generated from the beginning of the current month to the fifth day of the next month.
//
// The fileNo field exists only when the number of bill rows reaches the maximum rows in a single bill file and the bill is split into multiple files.
//
// ```
//
// **This topic is published only on the international site (alibabacloud.com).
//
// @param request - SubscriptionBillRequest
//
// @return SubscriptionBillResponse
func (client *Client) SubscriptionBill(request *SubscriptionBillRequest) (_result *SubscriptionBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubscriptionBillResponse{}
	_body, _err := client.SubscriptionBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
