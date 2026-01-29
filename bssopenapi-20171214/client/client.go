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
// Creates a financial relationship.
//
// Description:
//
// For more information about a financial relationship, see <props="intl">[Usage notes on the trusteeship]( https://www.alibabacloud.com/help/doc-detail/116383.html).
//
// If enterprise names used by the management account and a member for real-name verification are the same, you do not need to call an API operation for confirmation. Otherwise, you must call the ConfirmRelation operation for confirmation.
//
// @param request - AddAccountRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAccountRelationResponse
func (client *Client) AddAccountRelationWithOptions(request *AddAccountRelationRequest, runtime *dara.RuntimeOptions) (_result *AddAccountRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChildNick) {
		query["ChildNick"] = request.ChildNick
	}

	if !dara.IsNil(request.ChildUserId) {
		query["ChildUserId"] = request.ChildUserId
	}

	if !dara.IsNil(request.ParentUserId) {
		query["ParentUserId"] = request.ParentUserId
	}

	if !dara.IsNil(request.PermissionCodes) {
		query["PermissionCodes"] = request.PermissionCodes
	}

	if !dara.IsNil(request.RelationType) {
		query["RelationType"] = request.RelationType
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.RoleCodes) {
		query["RoleCodes"] = request.RoleCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAccountRelation"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAccountRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a financial relationship.
//
// Description:
//
// For more information about a financial relationship, see <props="intl">[Usage notes on the trusteeship]( https://www.alibabacloud.com/help/doc-detail/116383.html).
//
// If enterprise names used by the management account and a member for real-name verification are the same, you do not need to call an API operation for confirmation. Otherwise, you must call the ConfirmRelation operation for confirmation.
//
// @param request - AddAccountRelationRequest
//
// @return AddAccountRelationResponse
func (client *Client) AddAccountRelation(request *AddAccountRelationRequest) (_result *AddAccountRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAccountRelationResponse{}
	_body, _err := client.AddAccountRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transfers resource instances from the source cost center to the destination cost center.
//
// @param request - AllocateCostUnitResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateCostUnitResourceResponse
func (client *Client) AllocateCostUnitResourceWithOptions(request *AllocateCostUnitResourceRequest, runtime *dara.RuntimeOptions) (_result *AllocateCostUnitResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromUnitId) {
		query["FromUnitId"] = request.FromUnitId
	}

	if !dara.IsNil(request.FromUnitUserId) {
		query["FromUnitUserId"] = request.FromUnitUserId
	}

	if !dara.IsNil(request.ResourceInstanceList) {
		query["ResourceInstanceList"] = request.ResourceInstanceList
	}

	if !dara.IsNil(request.ToUnitId) {
		query["ToUnitId"] = request.ToUnitId
	}

	if !dara.IsNil(request.ToUnitUserId) {
		query["ToUnitUserId"] = request.ToUnitUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateCostUnitResource"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateCostUnitResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfers resource instances from the source cost center to the destination cost center.
//
// @param request - AllocateCostUnitResourceRequest
//
// @return AllocateCostUnitResourceResponse
func (client *Client) AllocateCostUnitResource(request *AllocateCostUnitResourceRequest) (_result *AllocateCostUnitResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateCostUnitResourceResponse{}
	_body, _err := client.AllocateCostUnitResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an application for an invoice.
//
// @param request - ApplyInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyInvoiceResponse
func (client *Client) ApplyInvoiceWithOptions(request *ApplyInvoiceRequest, runtime *dara.RuntimeOptions) (_result *ApplyInvoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ApplyUserNick) {
		query["ApplyUserNick"] = request.ApplyUserNick
	}

	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.InvoiceAmount) {
		query["InvoiceAmount"] = request.InvoiceAmount
	}

	if !dara.IsNil(request.InvoiceByAmount) {
		query["InvoiceByAmount"] = request.InvoiceByAmount
	}

	if !dara.IsNil(request.InvoicingType) {
		query["InvoicingType"] = request.InvoicingType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProcessWay) {
		query["ProcessWay"] = request.ProcessWay
	}

	if !dara.IsNil(request.SelectedIds) {
		query["SelectedIds"] = request.SelectedIds
	}

	if !dara.IsNil(request.UserRemark) {
		query["UserRemark"] = request.UserRemark
	}

	if !dara.IsNil(request.Emails) {
		query["emails"] = request.Emails
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyInvoice"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an application for an invoice.
//
// @param request - ApplyInvoiceRequest
//
// @return ApplyInvoiceResponse
func (client *Client) ApplyInvoice(request *ApplyInvoiceRequest) (_result *ApplyInvoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyInvoiceResponse{}
	_body, _err := client.ApplyInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an unpaid order.
//
// @param request - CancelOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOrderResponse
func (client *Client) CancelOrderWithOptions(request *CancelOrderRequest, runtime *dara.RuntimeOptions) (_result *CancelOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("CancelOrder"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an unpaid order.
//
// @param request - CancelOrderRequest
//
// @return CancelOrderResponse
func (client *Client) CancelOrder(request *CancelOrderRequest) (_result *CancelOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelOrderResponse{}
	_body, _err := client.CancelOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ChangeResellerConsumeAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResellerConsumeAmountResponse
func (client *Client) ChangeResellerConsumeAmountWithOptions(request *ChangeResellerConsumeAmountRequest, runtime *dara.RuntimeOptions) (_result *ChangeResellerConsumeAmountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdjustType) {
		query["AdjustType"] = request.AdjustType
	}

	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Currency) {
		query["Currency"] = request.Currency
	}

	if !dara.IsNil(request.ExtendMap) {
		query["ExtendMap"] = request.ExtendMap
	}

	if !dara.IsNil(request.OutBizId) {
		query["OutBizId"] = request.OutBizId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResellerConsumeAmount"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResellerConsumeAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeResellerConsumeAmountRequest
//
// @return ChangeResellerConsumeAmountResponse
func (client *Client) ChangeResellerConsumeAmount(request *ChangeResellerConsumeAmountRequest) (_result *ChangeResellerConsumeAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResellerConsumeAmountResponse{}
	_body, _err := client.ChangeResellerConsumeAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Confirms the invitation initiated by the master account.
//
// Description:
//
// 1\\. A member needs to confirm an invitation only if a financial management relationship is established between the management account and the member and enterprise names used by the management account and the member for real-name verification are different. 2. The permissions to be confirmed must be the same as those granted to the member when the management account initiates the invitation.
//
// @param request - ConfirmRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmRelationResponse
func (client *Client) ConfirmRelationWithOptions(request *ConfirmRelationRequest, runtime *dara.RuntimeOptions) (_result *ConfirmRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChildUserId) {
		query["ChildUserId"] = request.ChildUserId
	}

	if !dara.IsNil(request.ConfirmCode) {
		query["ConfirmCode"] = request.ConfirmCode
	}

	if !dara.IsNil(request.ParentUserId) {
		query["ParentUserId"] = request.ParentUserId
	}

	if !dara.IsNil(request.PermissionCodes) {
		query["PermissionCodes"] = request.PermissionCodes
	}

	if !dara.IsNil(request.RelationId) {
		query["RelationId"] = request.RelationId
	}

	if !dara.IsNil(request.RelationType) {
		query["RelationType"] = request.RelationType
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmRelation"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Confirms the invitation initiated by the master account.
//
// Description:
//
// 1\\. A member needs to confirm an invitation only if a financial management relationship is established between the management account and the member and enterprise names used by the management account and the member for real-name verification are different. 2. The permissions to be confirmed must be the same as those granted to the member when the management account initiates the invitation.
//
// @param request - ConfirmRelationRequest
//
// @return ConfirmRelationResponse
func (client *Client) ConfirmRelation(request *ConfirmRelationRequest) (_result *ConfirmRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfirmRelationResponse{}
	_body, _err := client.ConfirmRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the billing method of an instance. You can call this operation to switch the billing method from pay-as-you-go to subscription for Server Load Balancer (SLB) instances, elastic IP addresses (EIPs), and NAT gateways, and switch the billing method from subscription to pay-as-you-go for SLB instances and EIPs.
//
// @param request - ConvertChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertChargeTypeResponse
func (client *Client) ConvertChargeTypeWithOptions(request *ConvertChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ConvertChargeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
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
		Action:      dara.String("ConvertChargeType"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of an instance. You can call this operation to switch the billing method from pay-as-you-go to subscription for Server Load Balancer (SLB) instances, elastic IP addresses (EIPs), and NAT gateways, and switch the billing method from subscription to pay-as-you-go for SLB instances and EIPs.
//
// @param request - ConvertChargeTypeRequest
//
// @return ConvertChargeTypeResponse
func (client *Client) ConvertChargeType(request *ConvertChargeTypeRequest) (_result *ConvertChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConvertChargeTypeResponse{}
	_body, _err := client.ConvertChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an account to establish a financial relationship.
//
// Description:
//
// You can call this operation to create an account so as to establish a master-member financial relationship.
//
// @param request - CreateAgAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgAccountResponse
func (client *Client) CreateAgAccountWithOptions(request *CreateAgAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAgAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountAttr) {
		query["AccountAttr"] = request.AccountAttr
	}

	if !dara.IsNil(request.CityName) {
		query["CityName"] = request.CityName
	}

	if !dara.IsNil(request.EnterpriseName) {
		query["EnterpriseName"] = request.EnterpriseName
	}

	if !dara.IsNil(request.FirstName) {
		query["FirstName"] = request.FirstName
	}

	if !dara.IsNil(request.LastName) {
		query["LastName"] = request.LastName
	}

	if !dara.IsNil(request.LoginEmail) {
		query["LoginEmail"] = request.LoginEmail
	}

	if !dara.IsNil(request.NationCode) {
		query["NationCode"] = request.NationCode
	}

	if !dara.IsNil(request.Postcode) {
		query["Postcode"] = request.Postcode
	}

	if !dara.IsNil(request.ProvinceName) {
		query["ProvinceName"] = request.ProvinceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgAccount"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account to establish a financial relationship.
//
// Description:
//
// You can call this operation to create an account so as to establish a master-member financial relationship.
//
// @param request - CreateAgAccountRequest
//
// @return CreateAgAccountResponse
func (client *Client) CreateAgAccount(request *CreateAgAccountRequest) (_result *CreateAgAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAgAccountResponse{}
	_body, _err := client.CreateAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cost center. You can create multiple cost centers at a time.
//
// @param request - CreateCostUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostUnitResponse
func (client *Client) CreateCostUnitWithOptions(request *CreateCostUnitRequest, runtime *dara.RuntimeOptions) (_result *CreateCostUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UnitEntityList) {
		query["UnitEntityList"] = request.UnitEntityList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCostUnit"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCostUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cost center. You can create multiple cost centers at a time.
//
// @param request - CreateCostUnitRequest
//
// @return CreateCostUnitResponse
func (client *Client) CreateCostUnit(request *CreateCostUnitRequest) (_result *CreateCostUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCostUnitResponse{}
	_body, _err := client.CreateCostUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instance. If you call this operation, an order for a new instance is created and the order is automatically paid for. You cannot create Elastic Compute Service (ECS) instances or ApsaraDB RDS instances by calling the operation.
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Logistics) {
		query["Logistics"] = request.Logistics
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Parameter) {
		query["Parameter"] = request.Parameter
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RenewPeriod) {
		query["RenewPeriod"] = request.RenewPeriod
	}

	if !dara.IsNil(request.RenewalStatus) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance. If you call this operation, an order for a new instance is created and the order is automatically paid for. You cannot create Elastic Compute Service (ECS) instances or ApsaraDB RDS instances by calling the operation.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call failed.
//
// @param request - CreateResellerUserQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResellerUserQuotaResponse
func (client *Client) CreateResellerUserQuotaWithOptions(request *CreateResellerUserQuotaRequest, runtime *dara.RuntimeOptions) (_result *CreateResellerUserQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.Currency) {
		query["Currency"] = request.Currency
	}

	if !dara.IsNil(request.OutBizId) {
		query["OutBizId"] = request.OutBizId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResellerUserQuota"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResellerUserQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call failed.
//
// @param request - CreateResellerUserQuotaRequest
//
// @return CreateResellerUserQuotaResponse
func (client *Client) CreateResellerUserQuota(request *CreateResellerUserQuotaRequest) (_result *CreateResellerUserQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResellerUserQuotaResponse{}
	_body, _err := client.CreateResellerUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource plan.
//
// @param request - CreateResourcePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourcePackageResponse
func (client *Client) CreateResourcePackageWithOptions(request *CreateResourcePackageRequest, runtime *dara.RuntimeOptions) (_result *CreateResourcePackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EffectiveDate) {
		query["EffectiveDate"] = request.EffectiveDate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourcePackage"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource plan.
//
// @param request - CreateResourcePackageRequest
//
// @return CreateResourcePackageResponse
func (client *Client) CreateResourcePackage(request *CreateResourcePackageRequest) (_result *CreateResourcePackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourcePackageResponse{}
	_body, _err := client.CreateResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a savings plan. After you call this operation, a savings plan is purchased and paid for.
//
// @param tmpReq - CreateSavingsPlansInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSavingsPlansInstanceResponse
func (client *Client) CreateSavingsPlansInstanceWithOptions(tmpReq *CreateSavingsPlansInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateSavingsPlansInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSavingsPlansInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtendMap) {
		request.ExtendMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendMap, dara.String("ExtendMap"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EffectiveDate) {
		query["EffectiveDate"] = request.EffectiveDate
	}

	if !dara.IsNil(request.ExtendMapShrink) {
		query["ExtendMap"] = request.ExtendMapShrink
	}

	if !dara.IsNil(request.PayMode) {
		query["PayMode"] = request.PayMode
	}

	if !dara.IsNil(request.PoolValue) {
		query["PoolValue"] = request.PoolValue
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSavingsPlansInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSavingsPlansInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a savings plan. After you call this operation, a savings plan is purchased and paid for.
//
// @param request - CreateSavingsPlansInstanceRequest
//
// @return CreateSavingsPlansInstanceResponse
func (client *Client) CreateSavingsPlansInstance(request *CreateSavingsPlansInstanceRequest) (_result *CreateSavingsPlansInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSavingsPlansInstanceResponse{}
	_body, _err := client.CreateSavingsPlansInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a cost center.
//
// @param request - DeleteCostUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostUnitResponse
func (client *Client) DeleteCostUnitWithOptions(request *DeleteCostUnitRequest, runtime *dara.RuntimeOptions) (_result *DeleteCostUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerUid) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !dara.IsNil(request.UnitId) {
		query["UnitId"] = request.UnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCostUnit"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCostUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cost center.
//
// @param request - DeleteCostUnitRequest
//
// @return DeleteCostUnitResponse
func (client *Client) DeleteCostUnit(request *DeleteCostUnitRequest) (_result *DeleteCostUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCostUnitResponse{}
	_body, _err := client.DeleteCostUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the summary information of the user "Cost Management-Budget".
//
// Description:
//
// This operation is in beta testing and is only available for specific users in the whitelist. Excessive calls may result in performance issues. For example, the response times out.
//
// @param request - DescribeCostBudgetsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostBudgetsSummaryResponse
func (client *Client) DescribeCostBudgetsSummaryWithOptions(request *DescribeCostBudgetsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeCostBudgetsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BudgetName) {
		query["BudgetName"] = request.BudgetName
	}

	if !dara.IsNil(request.BudgetStatus) {
		query["BudgetStatus"] = request.BudgetStatus
	}

	if !dara.IsNil(request.BudgetType) {
		query["BudgetType"] = request.BudgetType
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
		Action:      dara.String("DescribeCostBudgetsSummary"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCostBudgetsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the summary information of the user "Cost Management-Budget".
//
// Description:
//
// This operation is in beta testing and is only available for specific users in the whitelist. Excessive calls may result in performance issues. For example, the response times out.
//
// @param request - DescribeCostBudgetsSummaryRequest
//
// @return DescribeCostBudgetsSummaryResponse
func (client *Client) DescribeCostBudgetsSummary(request *DescribeCostBudgetsSummaryRequest) (_result *DescribeCostBudgetsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCostBudgetsSummaryResponse{}
	_body, _err := client.DescribeCostBudgetsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monthly allocated costs of instances by allocation month.
//
// Description:
//
// You can view and export the allocated costs of the current month after 10:00 on the fourth day of the next month. The allocated costs of a single allocation month may involve orders or bills in different billing cycles. If a historical allocated amount is incorrect, the historical allocated costs need to be adjusted. As a result, the allocated costs displayed for a single allocation month may be different at different time points.
//
// @param request - DescribeInstanceAmortizedCostByAmortizationPeriodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAmortizedCostByAmortizationPeriodResponse
func (client *Client) DescribeInstanceAmortizedCostByAmortizationPeriodWithOptions(request *DescribeInstanceAmortizedCostByAmortizationPeriodRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAmortizedCostByAmortizationPeriodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerIdList) {
		body["BillOwnerIdList"] = request.BillOwnerIdList
	}

	if !dara.IsNil(request.BillUserIdList) {
		body["BillUserIdList"] = request.BillUserIdList
	}

	if !dara.IsNil(request.BillingCycle) {
		body["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.ConsumePeriodFilter) {
		body["ConsumePeriodFilter"] = request.ConsumePeriodFilter
	}

	if !dara.IsNil(request.CostUnitCode) {
		body["CostUnitCode"] = request.CostUnitCode
	}

	if !dara.IsNil(request.InstanceIdList) {
		body["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductDetail) {
		body["ProductDetail"] = request.ProductDetail
	}

	if !dara.IsNil(request.SubscriptionType) {
		body["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAmortizedCostByAmortizationPeriod"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAmortizedCostByAmortizationPeriodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monthly allocated costs of instances by allocation month.
//
// Description:
//
// You can view and export the allocated costs of the current month after 10:00 on the fourth day of the next month. The allocated costs of a single allocation month may involve orders or bills in different billing cycles. If a historical allocated amount is incorrect, the historical allocated costs need to be adjusted. As a result, the allocated costs displayed for a single allocation month may be different at different time points.
//
// @param request - DescribeInstanceAmortizedCostByAmortizationPeriodRequest
//
// @return DescribeInstanceAmortizedCostByAmortizationPeriodResponse
func (client *Client) DescribeInstanceAmortizedCostByAmortizationPeriod(request *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) (_result *DescribeInstanceAmortizedCostByAmortizationPeriodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAmortizedCostByAmortizationPeriodResponse{}
	_body, _err := client.DescribeInstanceAmortizedCostByAmortizationPeriodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例摊销日摊销成本
//
// @param request - DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse
func (client *Client) DescribeInstanceAmortizedCostByAmortizationPeriodDateWithOptions(request *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AmortizationDateEnd) {
		body["AmortizationDateEnd"] = request.AmortizationDateEnd
	}

	if !dara.IsNil(request.AmortizationDateStart) {
		body["AmortizationDateStart"] = request.AmortizationDateStart
	}

	if !dara.IsNil(request.BillOwnerIdList) {
		body["BillOwnerIdList"] = request.BillOwnerIdList
	}

	if !dara.IsNil(request.BillUserIdList) {
		body["BillUserIdList"] = request.BillUserIdList
	}

	if !dara.IsNil(request.BillingCycle) {
		body["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.CostUnitCode) {
		body["CostUnitCode"] = request.CostUnitCode
	}

	if !dara.IsNil(request.InstanceIdList) {
		body["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductDetail) {
		body["ProductDetail"] = request.ProductDetail
	}

	if !dara.IsNil(request.SubscriptionType) {
		body["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAmortizedCostByAmortizationPeriodDate"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例摊销日摊销成本
//
// @param request - DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
//
// @return DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse
func (client *Client) DescribeInstanceAmortizedCostByAmortizationPeriodDate(request *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) (_result *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAmortizedCostByAmortizationPeriodDateResponse{}
	_body, _err := client.DescribeInstanceAmortizedCostByAmortizationPeriodDateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例账期月摊销成本
//
// @param request - DescribeInstanceAmortizedCostByConsumePeriodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAmortizedCostByConsumePeriodResponse
func (client *Client) DescribeInstanceAmortizedCostByConsumePeriodWithOptions(request *DescribeInstanceAmortizedCostByConsumePeriodRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAmortizedCostByConsumePeriodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AmortizationPeriodFilter) {
		body["AmortizationPeriodFilter"] = request.AmortizationPeriodFilter
	}

	if !dara.IsNil(request.BillOwnerIdList) {
		body["BillOwnerIdList"] = request.BillOwnerIdList
	}

	if !dara.IsNil(request.BillUserIdList) {
		body["BillUserIdList"] = request.BillUserIdList
	}

	if !dara.IsNil(request.BillingCycle) {
		body["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.CostUnitCode) {
		body["CostUnitCode"] = request.CostUnitCode
	}

	if !dara.IsNil(request.InstanceIdList) {
		body["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductDetail) {
		body["ProductDetail"] = request.ProductDetail
	}

	if !dara.IsNil(request.SubscriptionType) {
		body["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAmortizedCostByConsumePeriod"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAmortizedCostByConsumePeriodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例账期月摊销成本
//
// @param request - DescribeInstanceAmortizedCostByConsumePeriodRequest
//
// @return DescribeInstanceAmortizedCostByConsumePeriodResponse
func (client *Client) DescribeInstanceAmortizedCostByConsumePeriod(request *DescribeInstanceAmortizedCostByConsumePeriodRequest) (_result *DescribeInstanceAmortizedCostByConsumePeriodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceAmortizedCostByConsumePeriodResponse{}
	_body, _err := client.DescribeInstanceAmortizedCostByConsumePeriodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the billing information about instances or billable items in a billing cycle.
//
// Description:
//
//	  Instance bills are generated after the total bill is split. In most cases, the instance bills do not include data generated on the last day of the specified billing cycle.
//
//		- The instance information may change during the billing cycle. The instance configurations and types in monthly bills are subject to the point in time when you query bills. For more information, see the corresponding bill details.
//
//		- You can query data generated after June 2020 for Cloud Communications services. You can query data generated after November 2020 for Alibaba Cloud Domains.
//
// @param request - DescribeInstanceBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceBillResponse
func (client *Client) DescribeInstanceBillWithOptions(request *DescribeInstanceBillRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.BillingDate) {
		query["BillingDate"] = request.BillingDate
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.InstanceID) {
		query["InstanceID"] = request.InstanceID
	}

	if !dara.IsNil(request.IsBillingItem) {
		query["IsBillingItem"] = request.IsBillingItem
	}

	if !dara.IsNil(request.IsHideZeroCharge) {
		query["IsHideZeroCharge"] = request.IsHideZeroCharge
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipCode) {
		query["PipCode"] = request.PipCode
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
		Action:      dara.String("DescribeInstanceBill"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billing information about instances or billable items in a billing cycle.
//
// Description:
//
//	  Instance bills are generated after the total bill is split. In most cases, the instance bills do not include data generated on the last day of the specified billing cycle.
//
//		- The instance information may change during the billing cycle. The instance configurations and types in monthly bills are subject to the point in time when you query bills. For more information, see the corresponding bill details.
//
//		- You can query data generated after June 2020 for Cloud Communications services. You can query data generated after November 2020 for Alibaba Cloud Domains.
//
// @param request - DescribeInstanceBillRequest
//
// @return DescribeInstanceBillResponse
func (client *Client) DescribeInstanceBill(request *DescribeInstanceBillRequest) (_result *DescribeInstanceBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceBillResponse{}
	_body, _err := client.DescribeInstanceBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 实例摊销日抵扣还原摊销成本
//
// @param request - DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse
func (client *Client) DescribeInstanceDeductAmortizedCostByAmortizationPeriodWithOptions(request *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerIdList) {
		body["BillOwnerIdList"] = request.BillOwnerIdList
	}

	if !dara.IsNil(request.BillUserIdList) {
		body["BillUserIdList"] = request.BillUserIdList
	}

	if !dara.IsNil(request.BillingCycle) {
		body["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.CostUnitCode) {
		body["CostUnitCode"] = request.CostUnitCode
	}

	if !dara.IsNil(request.InstanceIdList) {
		body["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductDetail) {
		body["ProductDetail"] = request.ProductDetail
	}

	if !dara.IsNil(request.SubscriptionType) {
		body["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceDeductAmortizedCostByAmortizationPeriod"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实例摊销日抵扣还原摊销成本
//
// @param request - DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
//
// @return DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse
func (client *Client) DescribeInstanceDeductAmortizedCostByAmortizationPeriod(request *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) (_result *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponse{}
	_body, _err := client.DescribeInstanceDeductAmortizedCostByAmortizationPeriodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the pricing information about an Alibaba Cloud service.
//
// @param request - DescribePricingModuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePricingModuleResponse
func (client *Client) DescribePricingModuleWithOptions(request *DescribePricingModuleRequest, runtime *dara.RuntimeOptions) (_result *DescribePricingModuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribePricingModule"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePricingModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the pricing information about an Alibaba Cloud service.
//
// @param request - DescribePricingModuleRequest
//
// @return DescribePricingModuleResponse
func (client *Client) DescribePricingModule(request *DescribePricingModuleRequest) (_result *DescribePricingModuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePricingModuleResponse{}
	_body, _err := client.DescribePricingModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the allocated costs of services by allocation month.
//
// Description:
//
// You can view and export the allocated costs of the current month after 10:00 on the fourth day of the next month. The allocated costs of a single allocation month may involve orders or bills in different billing cycles. If a historical allocated amount is incorrect, the historical allocated costs need to be adjusted. As a result, the allocated costs displayed for a single allocation month may be different at different time points.
//
// @param request - DescribeProductAmortizedCostByAmortizationPeriodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductAmortizedCostByAmortizationPeriodResponse
func (client *Client) DescribeProductAmortizedCostByAmortizationPeriodWithOptions(request *DescribeProductAmortizedCostByAmortizationPeriodRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductAmortizedCostByAmortizationPeriodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerIdList) {
		body["BillOwnerIdList"] = request.BillOwnerIdList
	}

	if !dara.IsNil(request.BillUserIdList) {
		body["BillUserIdList"] = request.BillUserIdList
	}

	if !dara.IsNil(request.BillingCycle) {
		body["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.ConsumePeriodFilter) {
		body["ConsumePeriodFilter"] = request.ConsumePeriodFilter
	}

	if !dara.IsNil(request.CostUnitCode) {
		body["CostUnitCode"] = request.CostUnitCode
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductDetail) {
		body["ProductDetail"] = request.ProductDetail
	}

	if !dara.IsNil(request.SubscriptionType) {
		body["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductAmortizedCostByAmortizationPeriod"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductAmortizedCostByAmortizationPeriodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the allocated costs of services by allocation month.
//
// Description:
//
// You can view and export the allocated costs of the current month after 10:00 on the fourth day of the next month. The allocated costs of a single allocation month may involve orders or bills in different billing cycles. If a historical allocated amount is incorrect, the historical allocated costs need to be adjusted. As a result, the allocated costs displayed for a single allocation month may be different at different time points.
//
// @param request - DescribeProductAmortizedCostByAmortizationPeriodRequest
//
// @return DescribeProductAmortizedCostByAmortizationPeriodResponse
func (client *Client) DescribeProductAmortizedCostByAmortizationPeriod(request *DescribeProductAmortizedCostByAmortizationPeriodRequest) (_result *DescribeProductAmortizedCostByAmortizationPeriodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProductAmortizedCostByAmortizationPeriodResponse{}
	_body, _err := client.DescribeProductAmortizedCostByAmortizationPeriodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 产品账期月摊销成本
//
// @param request - DescribeProductAmortizedCostByConsumePeriodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductAmortizedCostByConsumePeriodResponse
func (client *Client) DescribeProductAmortizedCostByConsumePeriodWithOptions(request *DescribeProductAmortizedCostByConsumePeriodRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductAmortizedCostByConsumePeriodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AmortizationPeriodFilter) {
		body["AmortizationPeriodFilter"] = request.AmortizationPeriodFilter
	}

	if !dara.IsNil(request.BillOwnerIdList) {
		body["BillOwnerIdList"] = request.BillOwnerIdList
	}

	if !dara.IsNil(request.BillUserIdList) {
		body["BillUserIdList"] = request.BillUserIdList
	}

	if !dara.IsNil(request.BillingCycle) {
		body["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.CostUnitCode) {
		body["CostUnitCode"] = request.CostUnitCode
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductDetail) {
		body["ProductDetail"] = request.ProductDetail
	}

	if !dara.IsNil(request.SubscriptionType) {
		body["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductAmortizedCostByConsumePeriod"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductAmortizedCostByConsumePeriodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 产品账期月摊销成本
//
// @param request - DescribeProductAmortizedCostByConsumePeriodRequest
//
// @return DescribeProductAmortizedCostByConsumePeriodResponse
func (client *Client) DescribeProductAmortizedCostByConsumePeriod(request *DescribeProductAmortizedCostByConsumePeriodRequest) (_result *DescribeProductAmortizedCostByConsumePeriodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProductAmortizedCostByConsumePeriodResponse{}
	_body, _err := client.DescribeProductAmortizedCostByConsumePeriodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the coverage details of reserved instances (RIs) or storage capacity units (SCUs).
//
// Description:
//
// 1\\. The queried coverage details are the same as those displayed in the table on the Coverage tab of the Manage Reserved Instances page in the Billing Management console.
//
// 2\\. You can call this operation to query the coverage details of RIs or SCUs.
//
// 3\\. You can call this operation to query coverage details at an hourly, daily, or monthly granularity.
//
// @param request - DescribeResourceCoverageDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceCoverageDetailResponse
func (client *Client) DescribeResourceCoverageDetailWithOptions(request *DescribeResourceCoverageDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceCoverageDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.EndPeriod) {
		query["EndPeriod"] = request.EndPeriod
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartPeriod) {
		query["StartPeriod"] = request.StartPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceCoverageDetail"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceCoverageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the coverage details of reserved instances (RIs) or storage capacity units (SCUs).
//
// Description:
//
// 1\\. The queried coverage details are the same as those displayed in the table on the Coverage tab of the Manage Reserved Instances page in the Billing Management console.
//
// 2\\. You can call this operation to query the coverage details of RIs or SCUs.
//
// 3\\. You can call this operation to query coverage details at an hourly, daily, or monthly granularity.
//
// @param request - DescribeResourceCoverageDetailRequest
//
// @return DescribeResourceCoverageDetailResponse
func (client *Client) DescribeResourceCoverageDetail(request *DescribeResourceCoverageDetailRequest) (_result *DescribeResourceCoverageDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceCoverageDetailResponse{}
	_body, _err := client.DescribeResourceCoverageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total coverage data of reserved instances (RIs) or storage capacity units (SCUs).
//
// Description:
//
// The queried total coverage data is the same as the aggregated data displayed on the Coverage tab of the Manage Reserved Instances page in the Billing Management console.
//
// You can call this operation to query the total coverage data of RIs or SCUs.
//
// @param request - DescribeResourceCoverageTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceCoverageTotalResponse
func (client *Client) DescribeResourceCoverageTotalWithOptions(request *DescribeResourceCoverageTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceCoverageTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.EndPeriod) {
		query["EndPeriod"] = request.EndPeriod
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartPeriod) {
		query["StartPeriod"] = request.StartPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceCoverageTotal"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceCoverageTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total coverage data of reserved instances (RIs) or storage capacity units (SCUs).
//
// Description:
//
// The queried total coverage data is the same as the aggregated data displayed on the Coverage tab of the Manage Reserved Instances page in the Billing Management console.
//
// You can call this operation to query the total coverage data of RIs or SCUs.
//
// @param request - DescribeResourceCoverageTotalRequest
//
// @return DescribeResourceCoverageTotalResponse
func (client *Client) DescribeResourceCoverageTotal(request *DescribeResourceCoverageTotalRequest) (_result *DescribeResourceCoverageTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceCoverageTotalResponse{}
	_body, _err := client.DescribeResourceCoverageTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about resource plans of an Alibaba Cloud service.
//
// @param request - DescribeResourcePackageProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourcePackageProductResponse
func (client *Client) DescribeResourcePackageProductWithOptions(request *DescribeResourcePackageProductRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourcePackageProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourcePackageProduct"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourcePackageProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about resource plans of an Alibaba Cloud service.
//
// @param request - DescribeResourcePackageProductRequest
//
// @return DescribeResourcePackageProductResponse
func (client *Client) DescribeResourcePackageProduct(request *DescribeResourcePackageProductRequest) (_result *DescribeResourcePackageProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourcePackageProductResponse{}
	_body, _err := client.DescribeResourcePackageProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage details of reserved instances (RIs) or storage capacity units (SCUs).
//
// @param request - DescribeResourceUsageDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceUsageDetailResponse
func (client *Client) DescribeResourceUsageDetailWithOptions(request *DescribeResourceUsageDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceUsageDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.EndPeriod) {
		query["EndPeriod"] = request.EndPeriod
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartPeriod) {
		query["StartPeriod"] = request.StartPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceUsageDetail"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceUsageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage details of reserved instances (RIs) or storage capacity units (SCUs).
//
// @param request - DescribeResourceUsageDetailRequest
//
// @return DescribeResourceUsageDetailResponse
func (client *Client) DescribeResourceUsageDetail(request *DescribeResourceUsageDetailRequest) (_result *DescribeResourceUsageDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceUsageDetailResponse{}
	_body, _err := client.DescribeResourceUsageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total usage data of reserved instances or storage capacity units (SCUs).
//
// @param request - DescribeResourceUsageTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceUsageTotalResponse
func (client *Client) DescribeResourceUsageTotalWithOptions(request *DescribeResourceUsageTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceUsageTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.EndPeriod) {
		query["EndPeriod"] = request.EndPeriod
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartPeriod) {
		query["StartPeriod"] = request.StartPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceUsageTotal"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceUsageTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total usage data of reserved instances or storage capacity units (SCUs).
//
// @param request - DescribeResourceUsageTotalRequest
//
// @return DescribeResourceUsageTotalResponse
func (client *Client) DescribeResourceUsageTotal(request *DescribeResourceUsageTotalRequest) (_result *DescribeResourceUsageTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceUsageTotalResponse{}
	_body, _err := client.DescribeResourceUsageTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the coverage details of savings plans.
//
// @param tmpReq - DescribeSavingsPlansCoverageDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSavingsPlansCoverageDetailResponse
func (client *Client) DescribeSavingsPlansCoverageDetailWithOptions(tmpReq *DescribeSavingsPlansCoverageDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeSavingsPlansCoverageDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSavingsPlansCoverageDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParam) {
		request.FilterParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParam, dara.String("FilterParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.EndPeriod) {
		query["EndPeriod"] = request.EndPeriod
	}

	if !dara.IsNil(request.FilterParamShrink) {
		query["FilterParam"] = request.FilterParamShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	if !dara.IsNil(request.StartPeriod) {
		query["StartPeriod"] = request.StartPeriod
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSavingsPlansCoverageDetail"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSavingsPlansCoverageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the coverage details of savings plans.
//
// @param request - DescribeSavingsPlansCoverageDetailRequest
//
// @return DescribeSavingsPlansCoverageDetailResponse
func (client *Client) DescribeSavingsPlansCoverageDetail(request *DescribeSavingsPlansCoverageDetailRequest) (_result *DescribeSavingsPlansCoverageDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSavingsPlansCoverageDetailResponse{}
	_body, _err := client.DescribeSavingsPlansCoverageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the coverage summary of savings plans.
//
// @param tmpReq - DescribeSavingsPlansCoverageTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSavingsPlansCoverageTotalResponse
func (client *Client) DescribeSavingsPlansCoverageTotalWithOptions(tmpReq *DescribeSavingsPlansCoverageTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribeSavingsPlansCoverageTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSavingsPlansCoverageTotalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParam) {
		request.FilterParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParam, dara.String("FilterParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.EndPeriod) {
		query["EndPeriod"] = request.EndPeriod
	}

	if !dara.IsNil(request.FilterParamShrink) {
		query["FilterParam"] = request.FilterParamShrink
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	if !dara.IsNil(request.StartPeriod) {
		query["StartPeriod"] = request.StartPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSavingsPlansCoverageTotal"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSavingsPlansCoverageTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the coverage summary of savings plans.
//
// @param request - DescribeSavingsPlansCoverageTotalRequest
//
// @return DescribeSavingsPlansCoverageTotalResponse
func (client *Client) DescribeSavingsPlansCoverageTotal(request *DescribeSavingsPlansCoverageTotalRequest) (_result *DescribeSavingsPlansCoverageTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSavingsPlansCoverageTotalResponse{}
	_body, _err := client.DescribeSavingsPlansCoverageTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage details of savings plans.
//
// @param tmpReq - DescribeSavingsPlansUsageDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSavingsPlansUsageDetailResponse
func (client *Client) DescribeSavingsPlansUsageDetailWithOptions(tmpReq *DescribeSavingsPlansUsageDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeSavingsPlansUsageDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSavingsPlansUsageDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParam) {
		request.FilterParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParam, dara.String("FilterParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.EndPeriod) {
		query["EndPeriod"] = request.EndPeriod
	}

	if !dara.IsNil(request.FilterParamShrink) {
		query["FilterParam"] = request.FilterParamShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	if !dara.IsNil(request.StartPeriod) {
		query["StartPeriod"] = request.StartPeriod
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSavingsPlansUsageDetail"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSavingsPlansUsageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage details of savings plans.
//
// @param request - DescribeSavingsPlansUsageDetailRequest
//
// @return DescribeSavingsPlansUsageDetailResponse
func (client *Client) DescribeSavingsPlansUsageDetail(request *DescribeSavingsPlansUsageDetailRequest) (_result *DescribeSavingsPlansUsageDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSavingsPlansUsageDetailResponse{}
	_body, _err := client.DescribeSavingsPlansUsageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage summary of savings plans.
//
// @param tmpReq - DescribeSavingsPlansUsageTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSavingsPlansUsageTotalResponse
func (client *Client) DescribeSavingsPlansUsageTotalWithOptions(tmpReq *DescribeSavingsPlansUsageTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribeSavingsPlansUsageTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSavingsPlansUsageTotalShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParam) {
		request.FilterParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParam, dara.String("FilterParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.EndPeriod) {
		query["EndPeriod"] = request.EndPeriod
	}

	if !dara.IsNil(request.FilterParamShrink) {
		query["FilterParam"] = request.FilterParamShrink
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	if !dara.IsNil(request.StartPeriod) {
		query["StartPeriod"] = request.StartPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSavingsPlansUsageTotal"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSavingsPlansUsageTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage summary of savings plans.
//
// @param request - DescribeSavingsPlansUsageTotalRequest
//
// @return DescribeSavingsPlansUsageTotalResponse
func (client *Client) DescribeSavingsPlansUsageTotal(request *DescribeSavingsPlansUsageTotalRequest) (_result *DescribeSavingsPlansUsageTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSavingsPlansUsageTotalResponse{}
	_body, _err := client.DescribeSavingsPlansUsageTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries split bills.
//
// Description:
//
//	  The data that you query by calling this operation is the same as the data that is queried by billing cycles in the Split Bill module of Cost Allocation.
//
//		- You can query split bills that were generated within the last 12 months by calling this operation.
//
//		- You can query split bills only after you enable the [Split Bill](https://usercenter2-intl.aliyun.com/finance/split-bill) service in the User Center console.
//
// @param request - DescribeSplitItemBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSplitItemBillResponse
func (client *Client) DescribeSplitItemBillWithOptions(request *DescribeSplitItemBillRequest, runtime *dara.RuntimeOptions) (_result *DescribeSplitItemBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.BillingDate) {
		query["BillingDate"] = request.BillingDate
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.InstanceID) {
		query["InstanceID"] = request.InstanceID
	}

	if !dara.IsNil(request.IsHideZeroCharge) {
		query["IsHideZeroCharge"] = request.IsHideZeroCharge
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipCode) {
		query["PipCode"] = request.PipCode
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SplitItemID) {
		query["SplitItemID"] = request.SplitItemID
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	if !dara.IsNil(request.TagFilter) {
		query["TagFilter"] = request.TagFilter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSplitItemBill"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSplitItemBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries split bills.
//
// Description:
//
//	  The data that you query by calling this operation is the same as the data that is queried by billing cycles in the Split Bill module of Cost Allocation.
//
//		- You can query split bills that were generated within the last 12 months by calling this operation.
//
//		- You can query split bills only after you enable the [Split Bill](https://usercenter2-intl.aliyun.com/finance/split-bill) service in the User Center console.
//
// @param request - DescribeSplitItemBillRequest
//
// @return DescribeSplitItemBillResponse
func (client *Client) DescribeSplitItemBill(request *DescribeSplitItemBillRequest) (_result *DescribeSplitItemBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSplitItemBillResponse{}
	_body, _err := client.DescribeSplitItemBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a financial relationship.
//
// @param request - GetAccountRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountRelationResponse
func (client *Client) GetAccountRelationWithOptions(request *GetAccountRelationRequest, runtime *dara.RuntimeOptions) (_result *GetAccountRelationResponse, _err error) {
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
		Action:      dara.String("GetAccountRelation"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a financial relationship.
//
// @param request - GetAccountRelationRequest
//
// @return GetAccountRelationResponse
func (client *Client) GetAccountRelation(request *GetAccountRelationRequest) (_result *GetAccountRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountRelationResponse{}
	_body, _err := client.GetAccountRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the account information about a customer of a virtual network operator (VNO).
//
// @param request - GetCustomerAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerAccountInfoResponse
func (client *Client) GetCustomerAccountInfoWithOptions(request *GetCustomerAccountInfoRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerAccountInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerAccountInfo"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the account information about a customer of a virtual network operator (VNO).
//
// @param request - GetCustomerAccountInfoRequest
//
// @return GetCustomerAccountInfoResponse
func (client *Client) GetCustomerAccountInfo(request *GetCustomerAccountInfoRequest) (_result *GetCustomerAccountInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomerAccountInfoResponse{}
	_body, _err := client.GetCustomerAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IDs of customers of a virtual network operator (VNO).
//
// Description:
//
// The system queries the IDs of customers of a VNO based on the AccessKey pair used in the request.
//
// @param request - GetCustomerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerListResponse
func (client *Client) GetCustomerListWithOptions(runtime *dara.RuntimeOptions) (_result *GetCustomerListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IDs of customers of a virtual network operator (VNO).
//
// Description:
//
// The system queries the IDs of customers of a VNO based on the AccessKey pair used in the request.
//
// @return GetCustomerListResponse
func (client *Client) GetCustomerList() (_result *GetCustomerListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomerListResponse{}
	_body, _err := client.GetCustomerListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an order that belongs to your Alibaba Cloud account or distributors.
//
// @param request - GetOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderDetailResponse
func (client *Client) GetOrderDetailWithOptions(request *GetOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Version:     dara.String("2017-12-14"),
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
// Queries the details of an order that belongs to your Alibaba Cloud account or distributors.
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
// Queries the pay-as-you-go price of an Alibaba Cloud service.
//
// Description:
//
// ### Usage notes
//
// 1.  Call the QueryProductList operation to obtain the code of the service. For more information, see [QueryProductList](https://help.aliyun.com/document_detail/95984.html).
//
// 2.  Call the DescribePricingModule operation to obtain the configuration parameters of the service. For more information, see [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
//
// 3.  Call the GetPayAsYouGoPrice operation to obtain the pay-as-you-go price of the service based on the returned configuration parameters.
//
// @param request - GetPayAsYouGoPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPayAsYouGoPriceResponse
func (client *Client) GetPayAsYouGoPriceWithOptions(request *GetPayAsYouGoPriceRequest, runtime *dara.RuntimeOptions) (_result *GetPayAsYouGoPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModuleList) {
		query["ModuleList"] = request.ModuleList
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPayAsYouGoPrice"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPayAsYouGoPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the pay-as-you-go price of an Alibaba Cloud service.
//
// Description:
//
// ### Usage notes
//
// 1.  Call the QueryProductList operation to obtain the code of the service. For more information, see [QueryProductList](https://help.aliyun.com/document_detail/95984.html).
//
// 2.  Call the DescribePricingModule operation to obtain the configuration parameters of the service. For more information, see [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
//
// 3.  Call the GetPayAsYouGoPrice operation to obtain the pay-as-you-go price of the service based on the returned configuration parameters.
//
// @param request - GetPayAsYouGoPriceRequest
//
// @return GetPayAsYouGoPriceResponse
func (client *Client) GetPayAsYouGoPrice(request *GetPayAsYouGoPriceRequest) (_result *GetPayAsYouGoPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPayAsYouGoPriceResponse{}
	_body, _err := client.GetPayAsYouGoPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price of a resource plan.
//
// @param request - GetResourcePackagePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcePackagePriceResponse
func (client *Client) GetResourcePackagePriceWithOptions(request *GetResourcePackagePriceRequest, runtime *dara.RuntimeOptions) (_result *GetResourcePackagePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EffectiveDate) {
		query["EffectiveDate"] = request.EffectiveDate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourcePackagePrice"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourcePackagePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price of a resource plan.
//
// @param request - GetResourcePackagePriceRequest
//
// @return GetResourcePackagePriceResponse
func (client *Client) GetResourcePackagePrice(request *GetResourcePackagePriceRequest) (_result *GetResourcePackagePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourcePackagePriceResponse{}
	_body, _err := client.GetResourcePackagePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the subscription price of an Alibaba Cloud service.
//
// Description:
//
// 1.  Call the QueryProductList operation to obtain the code of the service. For more information, see [QueryProductList](https://help.aliyun.com/document_detail/95984.html).
//
// 2.  Call the DescribePricingModule operation to obtain the configuration parameters of the service. For more information, see [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
//
// 3.  Call the GetSubscriptionPrice operation to obtain the pricing of the service based on the returned configuration parameters.
//
// @param request - GetSubscriptionPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionPriceResponse
func (client *Client) GetSubscriptionPriceWithOptions(request *GetSubscriptionPriceRequest, runtime *dara.RuntimeOptions) (_result *GetSubscriptionPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModuleList) {
		query["ModuleList"] = request.ModuleList
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ServicePeriodQuantity) {
		query["ServicePeriodQuantity"] = request.ServicePeriodQuantity
	}

	if !dara.IsNil(request.ServicePeriodUnit) {
		query["ServicePeriodUnit"] = request.ServicePeriodUnit
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubscriptionPrice"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubscriptionPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscription price of an Alibaba Cloud service.
//
// Description:
//
// 1.  Call the QueryProductList operation to obtain the code of the service. For more information, see [QueryProductList](https://help.aliyun.com/document_detail/95984.html).
//
// 2.  Call the DescribePricingModule operation to obtain the configuration parameters of the service. For more information, see [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
//
// 3.  Call the GetSubscriptionPrice operation to obtain the pricing of the service based on the returned configuration parameters.
//
// @param request - GetSubscriptionPriceRequest
//
// @return GetSubscriptionPriceResponse
func (client *Client) GetSubscriptionPrice(request *GetSubscriptionPriceRequest) (_result *GetSubscriptionPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSubscriptionPriceResponse{}
	_body, _err := client.GetSubscriptionPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the refundable amount for an instance from which you want to unsubscribe.
//
// Description:
//
// 1.  **Check the information about unsubscription and confirm the unsubscription terms and refundable amount. The resource that is unsubscribed cannot be restored.**
//
// 2.  Refunds are applicable only for the actual paid amount. Vouchers used for the purchase are non-refundable.
//
// 3.  For more information, see [Rules for unsubscribing from resources](https://www.alibabacloud.com/help/en/user-center/user-guide/refund-rules).
//
// @param request - InquiryPriceRefundInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InquiryPriceRefundInstanceResponse
func (client *Client) InquiryPriceRefundInstanceWithOptions(request *InquiryPriceRefundInstanceRequest, runtime *dara.RuntimeOptions) (_result *InquiryPriceRefundInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InquiryPriceRefundInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InquiryPriceRefundInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the refundable amount for an instance from which you want to unsubscribe.
//
// Description:
//
// 1.  **Check the information about unsubscription and confirm the unsubscription terms and refundable amount. The resource that is unsubscribed cannot be restored.**
//
// 2.  Refunds are applicable only for the actual paid amount. Vouchers used for the purchase are non-refundable.
//
// 3.  For more information, see [Rules for unsubscribing from resources](https://www.alibabacloud.com/help/en/user-center/user-guide/refund-rules).
//
// @param request - InquiryPriceRefundInstanceRequest
//
// @return InquiryPriceRefundInstanceResponse
func (client *Client) InquiryPriceRefundInstance(request *InquiryPriceRefundInstanceRequest) (_result *InquiryPriceRefundInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InquiryPriceRefundInstanceResponse{}
	_body, _err := client.InquiryPriceRefundInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or removes permissions granted to a member in a financial relationship.
//
// @param request - ModifyAccountRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountRelationResponse
func (client *Client) ModifyAccountRelationWithOptions(request *ModifyAccountRelationRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChildNick) {
		query["ChildNick"] = request.ChildNick
	}

	if !dara.IsNil(request.ChildUserId) {
		query["ChildUserId"] = request.ChildUserId
	}

	if !dara.IsNil(request.ParentUserId) {
		query["ParentUserId"] = request.ParentUserId
	}

	if !dara.IsNil(request.PermissionCodes) {
		query["PermissionCodes"] = request.PermissionCodes
	}

	if !dara.IsNil(request.RelationId) {
		query["RelationId"] = request.RelationId
	}

	if !dara.IsNil(request.RelationOperation) {
		query["RelationOperation"] = request.RelationOperation
	}

	if !dara.IsNil(request.RelationType) {
		query["RelationType"] = request.RelationType
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.RoleCodes) {
		query["RoleCodes"] = request.RoleCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountRelation"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or removes permissions granted to a member in a financial relationship.
//
// @param request - ModifyAccountRelationRequest
//
// @return ModifyAccountRelationResponse
func (client *Client) ModifyAccountRelation(request *ModifyAccountRelationRequest) (_result *ModifyAccountRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccountRelationResponse{}
	_body, _err := client.ModifyAccountRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies one or more cost centers.
//
// @param request - ModifyCostUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostUnitResponse
func (client *Client) ModifyCostUnitWithOptions(request *ModifyCostUnitRequest, runtime *dara.RuntimeOptions) (_result *ModifyCostUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UnitEntityList) {
		query["UnitEntityList"] = request.UnitEntityList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCostUnit"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCostUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies one or more cost centers.
//
// @param request - ModifyCostUnitRequest
//
// @return ModifyCostUnitResponse
func (client *Client) ModifyCostUnit(request *ModifyCostUnitRequest) (_result *ModifyCostUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCostUnitResponse{}
	_body, _err := client.ModifyCostUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an instance. When you call this operation, the system generates a modification order and automatically completes the payment. You cannot call this operation to modify the configurations of an Elastic Compute Service (ECS) instance or ApsaraDB RDS instance. To modify the configurations of an ECS or ApsaraDB RDS instance, call the dedicated operation of the corresponding service.
//
// @param request - ModifyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstanceWithOptions(request *ModifyInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Parameter) {
		query["Parameter"] = request.Parameter
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
		Action:      dara.String("ModifyInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an instance. When you call this operation, the system generates a modification order and automatically completes the payment. You cannot call this operation to modify the configurations of an Elastic Compute Service (ECS) instance or ApsaraDB RDS instance. To modify the configurations of an ECS or ApsaraDB RDS instance, call the dedicated operation of the corresponding service.
//
// @param request - ModifyInstanceRequest
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstance(request *ModifyInstanceRequest) (_result *ModifyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.ModifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对客订单支付接口
//
// @param request - PayOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PayOrderResponse
func (client *Client) PayOrderWithOptions(request *PayOrderRequest, runtime *dara.RuntimeOptions) (_result *PayOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuyerId) {
		query["BuyerId"] = request.BuyerId
	}

	if !dara.IsNil(request.EcIdAccountIds) {
		query["EcIdAccountIds"] = request.EcIdAccountIds
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PaySubmitUid) {
		query["PaySubmitUid"] = request.PaySubmitUid
	}

	if !dara.IsNil(request.PayerId) {
		query["PayerId"] = request.PayerId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PayOrder"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对客订单支付接口
//
// @param request - PayOrderRequest
//
// @return PayOrderResponse
func (client *Client) PayOrder(request *PayOrderRequest) (_result *PayOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PayOrderResponse{}
	_body, _err := client.PayOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the balance of your account.
//
// @param request - QueryAccountBalanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountBalanceResponse
func (client *Client) QueryAccountBalanceWithOptions(runtime *dara.RuntimeOptions) (_result *QueryAccountBalanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountBalance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountBalanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the balance of your account.
//
// @return QueryAccountBalanceResponse
func (client *Client) QueryAccountBalance() (_result *QueryAccountBalanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountBalanceResponse{}
	_body, _err := client.QueryAccountBalanceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bills of your Alibaba Cloud account within a billing cycle. You can summarize the bills by resource owner.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - Account bills are summarized based on instance bills. In most cases, the account bills do not include the data generated on the last day of the specified period.
//
//   - You can query the data generated in June 2020 or later for Cloud Communications services. However, the query results do not include the data of Alibaba Cloud Domains.
//
// @param request - QueryAccountBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountBillResponse
func (client *Client) QueryAccountBillWithOptions(request *QueryAccountBillRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.BillingDate) {
		query["BillingDate"] = request.BillingDate
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.IsGroupByProduct) {
		query["IsGroupByProduct"] = request.IsGroupByProduct
	}

	if !dara.IsNil(request.OwnerID) {
		query["OwnerID"] = request.OwnerID
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountBill"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bills of your Alibaba Cloud account within a billing cycle. You can summarize the bills by resource owner.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - Account bills are summarized based on instance bills. In most cases, the account bills do not include the data generated on the last day of the specified period.
//
//   - You can query the data generated in June 2020 or later for Cloud Communications services. However, the query results do not include the data of Alibaba Cloud Domains.
//
// @param request - QueryAccountBillRequest
//
// @return QueryAccountBillResponse
func (client *Client) QueryAccountBill(request *QueryAccountBillRequest) (_result *QueryAccountBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountBillResponse{}
	_body, _err := client.QueryAccountBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of transactions within your account.
//
// @param request - QueryAccountTransactionDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountTransactionDetailsResponse
func (client *Client) QueryAccountTransactionDetailsWithOptions(request *QueryAccountTransactionDetailsRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountTransactionDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RecordID) {
		query["RecordID"] = request.RecordID
	}

	if !dara.IsNil(request.TransactionChannel) {
		query["TransactionChannel"] = request.TransactionChannel
	}

	if !dara.IsNil(request.TransactionChannelSN) {
		query["TransactionChannelSN"] = request.TransactionChannelSN
	}

	if !dara.IsNil(request.TransactionNumber) {
		query["TransactionNumber"] = request.TransactionNumber
	}

	if !dara.IsNil(request.TransactionType) {
		query["TransactionType"] = request.TransactionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountTransactionDetails"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountTransactionDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of transactions within your account.
//
// @param request - QueryAccountTransactionDetailsRequest
//
// @return QueryAccountTransactionDetailsResponse
func (client *Client) QueryAccountTransactionDetails(request *QueryAccountTransactionDetailsRequest) (_result *QueryAccountTransactionDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountTransactionDetailsResponse{}
	_body, _err := client.QueryAccountTransactionDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries transactions within your Alibaba Cloud account.
//
// @param request - QueryAccountTransactionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountTransactionsResponse
func (client *Client) QueryAccountTransactionsWithOptions(request *QueryAccountTransactionsRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountTransactionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RecordID) {
		query["RecordID"] = request.RecordID
	}

	if !dara.IsNil(request.TransactionChannel) {
		query["TransactionChannel"] = request.TransactionChannel
	}

	if !dara.IsNil(request.TransactionChannelSN) {
		query["TransactionChannelSN"] = request.TransactionChannelSN
	}

	if !dara.IsNil(request.TransactionFlow) {
		query["TransactionFlow"] = request.TransactionFlow
	}

	if !dara.IsNil(request.TransactionNumber) {
		query["TransactionNumber"] = request.TransactionNumber
	}

	if !dara.IsNil(request.TransactionType) {
		query["TransactionType"] = request.TransactionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountTransactions"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountTransactionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries transactions within your Alibaba Cloud account.
//
// @param request - QueryAccountTransactionsRequest
//
// @return QueryAccountTransactionsResponse
func (client *Client) QueryAccountTransactions(request *QueryAccountTransactionsRequest) (_result *QueryAccountTransactionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountTransactionsResponse{}
	_body, _err := client.QueryAccountTransactionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available instances.
//
// @param request - QueryAvailableInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvailableInstancesResponse
func (client *Client) QueryAvailableInstancesWithOptions(request *QueryAvailableInstancesRequest, runtime *dara.RuntimeOptions) (_result *QueryAvailableInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.EndTimeEnd) {
		query["EndTimeEnd"] = request.EndTimeEnd
	}

	if !dara.IsNil(request.EndTimeStart) {
		query["EndTimeStart"] = request.EndTimeStart
	}

	if !dara.IsNil(request.InstanceIDs) {
		query["InstanceIDs"] = request.InstanceIDs
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

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RenewStatus) {
		query["RenewStatus"] = request.RenewStatus
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAvailableInstances"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvailableInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available instances.
//
// @param request - QueryAvailableInstancesRequest
//
// @return QueryAvailableInstancesResponse
func (client *Client) QueryAvailableInstances(request *QueryAvailableInstancesRequest) (_result *QueryAvailableInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAvailableInstancesResponse{}
	_body, _err := client.QueryAvailableInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bills in a billing cycle.
//
// @param request - QueryBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBillResponse
func (client *Client) QueryBillWithOptions(request *QueryBillRequest, runtime *dara.RuntimeOptions) (_result *QueryBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.IsDisplayLocalCurrency) {
		query["IsDisplayLocalCurrency"] = request.IsDisplayLocalCurrency
	}

	if !dara.IsNil(request.IsHideZeroCharge) {
		query["IsHideZeroCharge"] = request.IsHideZeroCharge
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

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBill"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bills in a billing cycle.
//
// @param request - QueryBillRequest
//
// @return QueryBillResponse
func (client *Client) QueryBill(request *QueryBillRequest) (_result *QueryBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBillResponse{}
	_body, _err := client.QueryBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bill overview information in a billing cycle.
//
// @param request - QueryBillOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBillOverviewResponse
func (client *Client) QueryBillOverviewWithOptions(request *QueryBillOverviewRequest, runtime *dara.RuntimeOptions) (_result *QueryBillOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
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
		Action:      dara.String("QueryBillOverview"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBillOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bill overview information in a billing cycle.
//
// @param request - QueryBillOverviewRequest
//
// @return QueryBillOverviewResponse
func (client *Client) QueryBillOverview(request *QueryBillOverviewRequest) (_result *QueryBillOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBillOverviewResponse{}
	_body, _err := client.QueryBillOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the subscribed bills that are stored in Object Storage Service (OSS) bucket.
//
// @param request - QueryBillToOSSSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBillToOSSSubscriptionResponse
func (client *Client) QueryBillToOSSSubscriptionWithOptions(runtime *dara.RuntimeOptions) (_result *QueryBillToOSSSubscriptionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBillToOSSSubscription"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBillToOSSSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscribed bills that are stored in Object Storage Service (OSS) bucket.
//
// @return QueryBillToOSSSubscriptionResponse
func (client *Client) QueryBillToOSSSubscription() (_result *QueryBillToOSSSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBillToOSSSubscriptionResponse{}
	_body, _err := client.QueryBillToOSSSubscriptionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about vouchers.
//
// @param request - QueryCashCouponsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCashCouponsResponse
func (client *Client) QueryCashCouponsWithOptions(request *QueryCashCouponsRequest, runtime *dara.RuntimeOptions) (_result *QueryCashCouponsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveOrNot) {
		query["EffectiveOrNot"] = request.EffectiveOrNot
	}

	if !dara.IsNil(request.ExpiryTimeEnd) {
		query["ExpiryTimeEnd"] = request.ExpiryTimeEnd
	}

	if !dara.IsNil(request.ExpiryTimeStart) {
		query["ExpiryTimeStart"] = request.ExpiryTimeStart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCashCoupons"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCashCouponsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about vouchers.
//
// @param request - QueryCashCouponsRequest
//
// @return QueryCashCouponsResponse
func (client *Client) QueryCashCoupons(request *QueryCashCouponsRequest) (_result *QueryCashCouponsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCashCouponsResponse{}
	_body, _err := client.QueryCashCouponsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a service based on the service code.
//
// Description:
//
// You can call this operation to query the information about a service based on the service code.
//
// @param request - QueryCommodityListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCommodityListResponse
func (client *Client) QueryCommodityListWithOptions(request *QueryCommodityListRequest, runtime *dara.RuntimeOptions) (_result *QueryCommodityListResponse, _err error) {
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
		Action:      dara.String("QueryCommodityList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCommodityListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a service based on the service code.
//
// Description:
//
// You can call this operation to query the information about a service based on the service code.
//
// @param request - QueryCommodityListRequest
//
// @return QueryCommodityListResponse
func (client *Client) QueryCommodityList(request *QueryCommodityListRequest) (_result *QueryCommodityListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCommodityListResponse{}
	_body, _err := client.QueryCommodityListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all cost centers within the current node of the cost center tree. If the ParentUnitId parameter is set to -1, all cost centers are queried.
//
// @param request - QueryCostUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostUnitResponse
func (client *Client) QueryCostUnitWithOptions(request *QueryCostUnitRequest, runtime *dara.RuntimeOptions) (_result *QueryCostUnitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerUid) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentUnitId) {
		query["ParentUnitId"] = request.ParentUnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCostUnit"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCostUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all cost centers within the current node of the cost center tree. If the ParentUnitId parameter is set to -1, all cost centers are queried.
//
// @param request - QueryCostUnitRequest
//
// @return QueryCostUnitResponse
func (client *Client) QueryCostUnit(request *QueryCostUnitRequest) (_result *QueryCostUnitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCostUnitResponse{}
	_body, _err := client.QueryCostUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource instances that are allocated to a cost center. If the unitId parameter is set to 0, the unallocated primary resource instances and sub-resource instances are queried.
//
// @param request - QueryCostUnitResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostUnitResourceResponse
func (client *Client) QueryCostUnitResourceWithOptions(request *QueryCostUnitResourceRequest, runtime *dara.RuntimeOptions) (_result *QueryCostUnitResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerUid) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UnitId) {
		query["UnitId"] = request.UnitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCostUnitResource"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCostUnitResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource instances that are allocated to a cost center. If the unitId parameter is set to 0, the unallocated primary resource instances and sub-resource instances are queried.
//
// @param request - QueryCostUnitResourceRequest
//
// @return QueryCostUnitResourceResponse
func (client *Client) QueryCostUnitResource(request *QueryCostUnitResourceRequest) (_result *QueryCostUnitResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCostUnitResourceResponse{}
	_body, _err := client.QueryCostUnitResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the addresses to which invoices are mailed.
//
// @param request - QueryCustomerAddressListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerAddressListResponse
func (client *Client) QueryCustomerAddressListWithOptions(request *QueryCustomerAddressListRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomerAddressListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCustomerAddressList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomerAddressListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the addresses to which invoices are mailed.
//
// @param request - QueryCustomerAddressListRequest
//
// @return QueryCustomerAddressListResponse
func (client *Client) QueryCustomerAddressList(request *QueryCustomerAddressListRequest) (_result *QueryCustomerAddressListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCustomerAddressListResponse{}
	_body, _err := client.QueryCustomerAddressListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage of resource plans, including reserved instances (RIs) and storage capacity units (SCUs).
//
// Description:
//
// Limits:
//
//   - Only the usage records within the past year can be queried.
//
// @param request - QueryDPUtilizationDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDPUtilizationDetailResponse
func (client *Client) QueryDPUtilizationDetailWithOptions(request *QueryDPUtilizationDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryDPUtilizationDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.DeductedInstanceId) {
		query["DeductedInstanceId"] = request.DeductedInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IncludeShare) {
		query["IncludeShare"] = request.IncludeShare
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.LastToken) {
		query["LastToken"] = request.LastToken
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDPUtilizationDetail"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDPUtilizationDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage of resource plans, including reserved instances (RIs) and storage capacity units (SCUs).
//
// Description:
//
// Limits:
//
//   - Only the usage records within the past year can be queried.
//
// @param request - QueryDPUtilizationDetailRequest
//
// @return QueryDPUtilizationDetailResponse
func (client *Client) QueryDPUtilizationDetail(request *QueryDPUtilizationDetailRequest) (_result *QueryDPUtilizationDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDPUtilizationDetailResponse{}
	_body, _err := client.QueryDPUtilizationDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the orders for which you want to apply for invoices.
//
// @param request - QueryEvaluateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEvaluateListResponse
func (client *Client) QueryEvaluateListWithOptions(request *QueryEvaluateListRequest, runtime *dara.RuntimeOptions) (_result *QueryEvaluateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillCycle) {
		query["BillCycle"] = request.BillCycle
	}

	if !dara.IsNil(request.BizTypeList) {
		query["BizTypeList"] = request.BizTypeList
	}

	if !dara.IsNil(request.EndAmount) {
		query["EndAmount"] = request.EndAmount
	}

	if !dara.IsNil(request.EndBizTime) {
		query["EndBizTime"] = request.EndBizTime
	}

	if !dara.IsNil(request.EndSearchTime) {
		query["EndSearchTime"] = request.EndSearchTime
	}

	if !dara.IsNil(request.OutBizId) {
		query["OutBizId"] = request.OutBizId
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

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.StartAmount) {
		query["StartAmount"] = request.StartAmount
	}

	if !dara.IsNil(request.StartBizTime) {
		query["StartBizTime"] = request.StartBizTime
	}

	if !dara.IsNil(request.StartSearchTime) {
		query["StartSearchTime"] = request.StartSearchTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEvaluateList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEvaluateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the orders for which you want to apply for invoices.
//
// @param request - QueryEvaluateListRequest
//
// @return QueryEvaluateListResponse
func (client *Client) QueryEvaluateList(request *QueryEvaluateListRequest) (_result *QueryEvaluateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEvaluateListResponse{}
	_body, _err := client.QueryEvaluateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a financial account.
//
// @param request - QueryFinancialAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFinancialAccountInfoResponse
func (client *Client) QueryFinancialAccountInfoWithOptions(request *QueryFinancialAccountInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryFinancialAccountInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFinancialAccountInfo"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFinancialAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a financial account.
//
// @param request - QueryFinancialAccountInfoRequest
//
// @return QueryFinancialAccountInfoResponse
func (client *Client) QueryFinancialAccountInfo(request *QueryFinancialAccountInfoRequest) (_result *QueryFinancialAccountInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFinancialAccountInfoResponse{}
	_body, _err := client.QueryFinancialAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bills of instances or billable items in a billing cycle.
//
// Description:
//
// ##
//
//   - This API operation has been upgraded to DescribeInstanceBill. We recommend that you call the [DescribeInstanceBill](https://help.aliyun.com/document_detail/209402.html) operation to query the bills of instances or billable items in a billing cycle. You can call the QueryInstanceBill operation to query a maximum of 50,000 data rows in a bill.
//
//   - Instance bills are generated after bills are split. In most cases, the instance bills do not include data generated on the last day of the specified period.
//
//   - The instance information changes within a billing cycle. The instance configurations and specifications and the time when the instance was used in the billing cycle are all recorded. For more information, see the corresponding bill details.
//
//   - You can query the data generated in June 2020 or later for Cloud Communications services, and the data generated in November 2020 or later for Alibaba Cloud Domains.
//
// @param request - QueryInstanceBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceBillResponse
func (client *Client) QueryInstanceBillWithOptions(request *QueryInstanceBillRequest, runtime *dara.RuntimeOptions) (_result *QueryInstanceBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.BillingDate) {
		query["BillingDate"] = request.BillingDate
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.IsBillingItem) {
		query["IsBillingItem"] = request.IsBillingItem
	}

	if !dara.IsNil(request.IsHideZeroCharge) {
		query["IsHideZeroCharge"] = request.IsHideZeroCharge
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
		Action:      dara.String("QueryInstanceBill"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstanceBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bills of instances or billable items in a billing cycle.
//
// Description:
//
// ##
//
//   - This API operation has been upgraded to DescribeInstanceBill. We recommend that you call the [DescribeInstanceBill](https://help.aliyun.com/document_detail/209402.html) operation to query the bills of instances or billable items in a billing cycle. You can call the QueryInstanceBill operation to query a maximum of 50,000 data rows in a bill.
//
//   - Instance bills are generated after bills are split. In most cases, the instance bills do not include data generated on the last day of the specified period.
//
//   - The instance information changes within a billing cycle. The instance configurations and specifications and the time when the instance was used in the billing cycle are all recorded. For more information, see the corresponding bill details.
//
//   - You can query the data generated in June 2020 or later for Cloud Communications services, and the data generated in November 2020 or later for Alibaba Cloud Domains.
//
// @param request - QueryInstanceBillRequest
//
// @return QueryInstanceBillResponse
func (client *Client) QueryInstanceBill(request *QueryInstanceBillRequest) (_result *QueryInstanceBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInstanceBillResponse{}
	_body, _err := client.QueryInstanceBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instances by tag.
//
// @param request - QueryInstanceByTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceByTagResponse
func (client *Client) QueryInstanceByTagWithOptions(request *QueryInstanceByTagRequest, runtime *dara.RuntimeOptions) (_result *QueryInstanceByTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInstanceByTag"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstanceByTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instances by tag.
//
// @param request - QueryInstanceByTagRequest
//
// @return QueryInstanceByTagResponse
func (client *Client) QueryInstanceByTag(request *QueryInstanceByTagRequest) (_result *QueryInstanceByTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInstanceByTagResponse{}
	_body, _err := client.QueryInstanceByTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The code of the service.
//
// @param request - QueryInstanceGaapCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceGaapCostResponse
func (client *Client) QueryInstanceGaapCostWithOptions(request *QueryInstanceGaapCostRequest, runtime *dara.RuntimeOptions) (_result *QueryInstanceGaapCostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("QueryInstanceGaapCost"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstanceGaapCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The code of the service.
//
// @param request - QueryInstanceGaapCostRequest
//
// @return QueryInstanceGaapCostResponse
func (client *Client) QueryInstanceGaapCost(request *QueryInstanceGaapCostRequest) (_result *QueryInstanceGaapCostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInstanceGaapCostResponse{}
	_body, _err := client.QueryInstanceGaapCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about invoice titles.
//
// @param request - QueryInvoicingCustomerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInvoicingCustomerListResponse
func (client *Client) QueryInvoicingCustomerListWithOptions(request *QueryInvoicingCustomerListRequest, runtime *dara.RuntimeOptions) (_result *QueryInvoicingCustomerListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInvoicingCustomerList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInvoicingCustomerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about invoice titles.
//
// @param request - QueryInvoicingCustomerListRequest
//
// @return QueryInvoicingCustomerListResponse
func (client *Client) QueryInvoicingCustomerList(request *QueryInvoicingCustomerListRequest) (_result *QueryInvoicingCustomerListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInvoicingCustomerListResponse{}
	_body, _err := client.QueryInvoicingCustomerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the orders of your Alibaba Cloud account or distributors. By default, orders within the last hour are queried. To query earlier orders, specify the CreateTimeStart and CreateTimeEnd parameters.
//
// @param request - QueryOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrdersResponse
func (client *Client) QueryOrdersWithOptions(request *QueryOrdersRequest, runtime *dara.RuntimeOptions) (_result *QueryOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
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
		Action:      dara.String("QueryOrders"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the orders of your Alibaba Cloud account or distributors. By default, orders within the last hour are queried. To query earlier orders, specify the CreateTimeStart and CreateTimeEnd parameters.
//
// @param request - QueryOrdersRequest
//
// @return QueryOrdersResponse
func (client *Client) QueryOrders(request *QueryOrdersRequest) (_result *QueryOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryOrdersResponse{}
	_body, _err := client.QueryOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries, by relationship ID, permissions granted to accounts between which a management-member relationship is established.
//
// @param request - QueryPermissionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPermissionListResponse
func (client *Client) QueryPermissionListWithOptions(request *QueryPermissionListRequest, runtime *dara.RuntimeOptions) (_result *QueryPermissionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RelationId) {
		query["RelationId"] = request.RelationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPermissionList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPermissionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries, by relationship ID, permissions granted to accounts between which a management-member relationship is established.
//
// @param request - QueryPermissionListRequest
//
// @return QueryPermissionListResponse
func (client *Client) QueryPermissionList(request *QueryPermissionListRequest) (_result *QueryPermissionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPermissionListResponse{}
	_body, _err := client.QueryPermissionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries prepaid cards.
//
// @param request - QueryPrepaidCardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPrepaidCardsResponse
func (client *Client) QueryPrepaidCardsWithOptions(request *QueryPrepaidCardsRequest, runtime *dara.RuntimeOptions) (_result *QueryPrepaidCardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveOrNot) {
		query["EffectiveOrNot"] = request.EffectiveOrNot
	}

	if !dara.IsNil(request.ExpiryTimeEnd) {
		query["ExpiryTimeEnd"] = request.ExpiryTimeEnd
	}

	if !dara.IsNil(request.ExpiryTimeStart) {
		query["ExpiryTimeStart"] = request.ExpiryTimeStart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPrepaidCards"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPrepaidCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries prepaid cards.
//
// @param request - QueryPrepaidCardsRequest
//
// @return QueryPrepaidCardsResponse
func (client *Client) QueryPrepaidCards(request *QueryPrepaidCardsRequest) (_result *QueryPrepaidCardsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPrepaidCardsResponse{}
	_body, _err := client.QueryPrepaidCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the billable items of a service.
//
// Description:
//
// You can call this operation to query the billable items of a service. A billable item is the minimum unit used to calculate costs.
//
// @param request - QueryPriceEntityListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPriceEntityListResponse
func (client *Client) QueryPriceEntityListWithOptions(request *QueryPriceEntityListRequest, runtime *dara.RuntimeOptions) (_result *QueryPriceEntityListResponse, _err error) {
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
		Action:      dara.String("QueryPriceEntityList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPriceEntityListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billable items of a service.
//
// Description:
//
// You can call this operation to query the billable items of a service. A billable item is the minimum unit used to calculate costs.
//
// @param request - QueryPriceEntityListRequest
//
// @return QueryPriceEntityListResponse
func (client *Client) QueryPriceEntityList(request *QueryPriceEntityListRequest) (_result *QueryPriceEntityListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPriceEntityListResponse{}
	_body, _err := client.QueryPriceEntityListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all Alibaba Cloud services.
//
// @param request - QueryProductListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProductListResponse
func (client *Client) QueryProductListWithOptions(request *QueryProductListRequest, runtime *dara.RuntimeOptions) (_result *QueryProductListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryTotalCount) {
		query["QueryTotalCount"] = request.QueryTotalCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProductList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProductListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all Alibaba Cloud services.
//
// @param request - QueryProductListRequest
//
// @return QueryProductListResponse
func (client *Client) QueryProductList(request *QueryProductListRequest) (_result *QueryProductListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryProductListResponse{}
	_body, _err := client.QueryProductListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage details of a reserved instance (RI).
//
// @param request - QueryRIUtilizationDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRIUtilizationDetailResponse
func (client *Client) QueryRIUtilizationDetailWithOptions(request *QueryRIUtilizationDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryRIUtilizationDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeductedInstanceId) {
		query["DeductedInstanceId"] = request.DeductedInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RICommodityCode) {
		query["RICommodityCode"] = request.RICommodityCode
	}

	if !dara.IsNil(request.RIInstanceId) {
		query["RIInstanceId"] = request.RIInstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRIUtilizationDetail"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRIUtilizationDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage details of a reserved instance (RI).
//
// @param request - QueryRIUtilizationDetailRequest
//
// @return QueryRIUtilizationDetailResponse
func (client *Client) QueryRIUtilizationDetail(request *QueryRIUtilizationDetailRequest) (_result *QueryRIUtilizationDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRIUtilizationDetailResponse{}
	_body, _err := client.QueryRIUtilizationDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a redemption coupon.
//
// @param request - QueryRedeemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRedeemResponse
func (client *Client) QueryRedeemWithOptions(request *QueryRedeemRequest, runtime *dara.RuntimeOptions) (_result *QueryRedeemResponse, _err error) {
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
		Action:      dara.String("QueryRedeem"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRedeemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a redemption coupon.
//
// @param request - QueryRedeemRequest
//
// @return QueryRedeemResponse
func (client *Client) QueryRedeem(request *QueryRedeemRequest) (_result *QueryRedeemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRedeemResponse{}
	_body, _err := client.QueryRedeemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the members of a management account.
//
// @param request - QueryRelationListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRelationListResponse
func (client *Client) QueryRelationListWithOptions(request *QueryRelationListRequest, runtime *dara.RuntimeOptions) (_result *QueryRelationListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRelationList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRelationListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the members of a management account.
//
// @param request - QueryRelationListRequest
//
// @return QueryRelationListResponse
func (client *Client) QueryRelationList(request *QueryRelationListRequest) (_result *QueryRelationListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRelationListResponse{}
	_body, _err := client.QueryRelationListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call failed.
//
// @param request - QueryResellerAvailableQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryResellerAvailableQuotaResponse
func (client *Client) QueryResellerAvailableQuotaWithOptions(request *QueryResellerAvailableQuotaRequest, runtime *dara.RuntimeOptions) (_result *QueryResellerAvailableQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ItemCodes) {
		query["ItemCodes"] = request.ItemCodes
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryResellerAvailableQuota"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryResellerAvailableQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call failed.
//
// @param request - QueryResellerAvailableQuotaRequest
//
// @return QueryResellerAvailableQuotaResponse
func (client *Client) QueryResellerAvailableQuota(request *QueryResellerAvailableQuotaRequest) (_result *QueryResellerAvailableQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryResellerAvailableQuotaResponse{}
	_body, _err := client.QueryResellerAvailableQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户的信控预警阀值,该接口暂未测试启用
//
// @param request - QueryResellerUserAlarmThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryResellerUserAlarmThresholdResponse
func (client *Client) QueryResellerUserAlarmThresholdWithOptions(request *QueryResellerUserAlarmThresholdRequest, runtime *dara.RuntimeOptions) (_result *QueryResellerUserAlarmThresholdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmType) {
		query["AlarmType"] = request.AlarmType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryResellerUserAlarmThreshold"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryResellerUserAlarmThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户的信控预警阀值,该接口暂未测试启用
//
// @param request - QueryResellerUserAlarmThresholdRequest
//
// @return QueryResellerUserAlarmThresholdResponse
func (client *Client) QueryResellerUserAlarmThreshold(request *QueryResellerUserAlarmThresholdRequest) (_result *QueryResellerUserAlarmThresholdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryResellerUserAlarmThresholdResponse{}
	_body, _err := client.QueryResellerUserAlarmThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the instances of a resource plan. You can query the resource plans that are expired within one year.
//
// @param request - QueryResourcePackageInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryResourcePackageInstancesResponse
func (client *Client) QueryResourcePackageInstancesWithOptions(request *QueryResourcePackageInstancesRequest, runtime *dara.RuntimeOptions) (_result *QueryResourcePackageInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpiryTimeEnd) {
		query["ExpiryTimeEnd"] = request.ExpiryTimeEnd
	}

	if !dara.IsNil(request.ExpiryTimeStart) {
		query["ExpiryTimeStart"] = request.ExpiryTimeStart
	}

	if !dara.IsNil(request.IncludePartner) {
		query["IncludePartner"] = request.IncludePartner
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

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryResourcePackageInstances"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryResourcePackageInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instances of a resource plan. You can query the resource plans that are expired within one year.
//
// @param request - QueryResourcePackageInstancesRequest
//
// @return QueryResourcePackageInstancesResponse
func (client *Client) QueryResourcePackageInstances(request *QueryResourcePackageInstancesRequest) (_result *QueryResourcePackageInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryResourcePackageInstancesResponse{}
	_body, _err := client.QueryResourcePackageInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the deduction details of savings plans.
//
// @param request - QuerySavingsPlansDeductLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySavingsPlansDeductLogResponse
func (client *Client) QuerySavingsPlansDeductLogWithOptions(request *QuerySavingsPlansDeductLogRequest, runtime *dara.RuntimeOptions) (_result *QuerySavingsPlansDeductLogResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Locale) {
		query["Locale"] = request.Locale
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySavingsPlansDeductLog"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySavingsPlansDeductLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deduction details of savings plans.
//
// @param request - QuerySavingsPlansDeductLogRequest
//
// @return QuerySavingsPlansDeductLogResponse
func (client *Client) QuerySavingsPlansDeductLog(request *QuerySavingsPlansDeductLogRequest) (_result *QuerySavingsPlansDeductLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySavingsPlansDeductLogResponse{}
	_body, _err := client.QuerySavingsPlansDeductLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries discounts on savings plans.
//
// @param request - QuerySavingsPlansDiscountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySavingsPlansDiscountResponse
func (client *Client) QuerySavingsPlansDiscountWithOptions(request *QuerySavingsPlansDiscountRequest, runtime *dara.RuntimeOptions) (_result *QuerySavingsPlansDiscountResponse, _err error) {
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
		Action:      dara.String("QuerySavingsPlansDiscount"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySavingsPlansDiscountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries discounts on savings plans.
//
// @param request - QuerySavingsPlansDiscountRequest
//
// @return QuerySavingsPlansDiscountResponse
func (client *Client) QuerySavingsPlansDiscount(request *QuerySavingsPlansDiscountRequest) (_result *QuerySavingsPlansDiscountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySavingsPlansDiscountResponse{}
	_body, _err := client.QuerySavingsPlansDiscountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about savings plan instances of the current user.
//
// @param request - QuerySavingsPlansInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySavingsPlansInstanceResponse
func (client *Client) QuerySavingsPlansInstanceWithOptions(request *QuerySavingsPlansInstanceRequest, runtime *dara.RuntimeOptions) (_result *QuerySavingsPlansInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Locale) {
		query["Locale"] = request.Locale
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySavingsPlansInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySavingsPlansInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about savings plan instances of the current user.
//
// @param request - QuerySavingsPlansInstanceRequest
//
// @return QuerySavingsPlansInstanceResponse
func (client *Client) QuerySavingsPlansInstance(request *QuerySavingsPlansInstanceRequest) (_result *QuerySavingsPlansInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySavingsPlansInstanceResponse{}
	_body, _err := client.QuerySavingsPlansInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The code of the service.
//
// @param request - QuerySettleBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySettleBillResponse
func (client *Client) QuerySettleBillWithOptions(request *QuerySettleBillRequest, runtime *dara.RuntimeOptions) (_result *QuerySettleBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.IsDisplayLocalCurrency) {
		query["IsDisplayLocalCurrency"] = request.IsDisplayLocalCurrency
	}

	if !dara.IsNil(request.IsHideZeroCharge) {
		query["IsHideZeroCharge"] = request.IsHideZeroCharge
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RecordID) {
		query["RecordID"] = request.RecordID
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySettleBill"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySettleBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The code of the service.
//
// @param request - QuerySettleBillRequest
//
// @return QuerySettleBillResponse
func (client *Client) QuerySettleBill(request *QuerySettleBillRequest) (_result *QuerySettleBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySettleBillResponse{}
	_body, _err := client.QuerySettleBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the stock keeping units (SKUs) of a service. In most cases, a service has one or more SKUs. A service may even have tens of thousands of SKUs. You can call this operation to query the SKUs of a specific service and the prices of the SKUs. You can configure request parameters to query the specified SKUs based on the configurations of the SKUs.
//
// @param tmpReq - QuerySkuPriceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySkuPriceListResponse
func (client *Client) QuerySkuPriceListWithOptions(tmpReq *QuerySkuPriceListRequest, runtime *dara.RuntimeOptions) (_result *QuerySkuPriceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QuerySkuPriceListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PriceFactorConditionMap) {
		request.PriceFactorConditionMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PriceFactorConditionMap, dara.String("PriceFactorConditionMap"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySkuPriceList"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySkuPriceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the stock keeping units (SKUs) of a service. In most cases, a service has one or more SKUs. A service may even have tens of thousands of SKUs. You can call this operation to query the SKUs of a specific service and the prices of the SKUs. You can configure request parameters to query the specified SKUs based on the configurations of the SKUs.
//
// @param request - QuerySkuPriceListRequest
//
// @return QuerySkuPriceListResponse
func (client *Client) QuerySkuPriceList(request *QuerySkuPriceListRequest) (_result *QuerySkuPriceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySkuPriceListResponse{}
	_body, _err := client.QuerySkuPriceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries split bills.
//
// Description:
//
//	  This API operation has been upgraded to DescribeSplitItemBill. We recommend that you call the [DescribeSplitItemBill](https://help.aliyun.com/document_detail/208169.html) operation to query split bills. You can call the QuerySplitItemBill operation to query a maximum of 50,000 data rows in a bill.
//
//		- The data queried by calling the QuerySplitItemBill operation is consistent with the data that is displayed for the specified billing cycle on the Split Bill page in User Center.
//
//		- You can call this operation to query split bills generated within the last 12 months.
//
//		- This operation returns split bills only after you activate the [Split Bill](https://usercenter2.aliyun.com/finance/split-bill) service in User Center.
//
// @param request - QuerySplitItemBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySplitItemBillResponse
func (client *Client) QuerySplitItemBillWithOptions(request *QuerySplitItemBillRequest, runtime *dara.RuntimeOptions) (_result *QuerySplitItemBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillOwnerId) {
		query["BillOwnerId"] = request.BillOwnerId
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
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
		Action:      dara.String("QuerySplitItemBill"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySplitItemBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries split bills.
//
// Description:
//
//	  This API operation has been upgraded to DescribeSplitItemBill. We recommend that you call the [DescribeSplitItemBill](https://help.aliyun.com/document_detail/208169.html) operation to query split bills. You can call the QuerySplitItemBill operation to query a maximum of 50,000 data rows in a bill.
//
//		- The data queried by calling the QuerySplitItemBill operation is consistent with the data that is displayed for the specified billing cycle on the Split Bill page in User Center.
//
//		- You can call this operation to query split bills generated within the last 12 months.
//
//		- This operation returns split bills only after you activate the [Split Bill](https://usercenter2.aliyun.com/finance/split-bill) service in User Center.
//
// @param request - QuerySplitItemBillRequest
//
// @return QuerySplitItemBillResponse
func (client *Client) QuerySplitItemBill(request *QuerySplitItemBillRequest) (_result *QuerySplitItemBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySplitItemBillResponse{}
	_body, _err := client.QuerySplitItemBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage data of an Alibaba Cloud service.
//
// Description:
//
// You can call this operation to query the usage data of an Alibaba Cloud service. Take note of the following items:
//
//   - The service code that you specify for querying the usage data of a specific Alibaba Cloud service must be valid. You can query the usage data by hour or by day.
//
//   - The time that you specify must follow the ISO8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
//
//   - Latency exists in data pushes. Therefore, if you set the DataType parameter to Hour, the integrity of usage data recorded in the last 24 hours can be ensured. If you set the DataType parameter to Day, the integrity of usage data recorded in the last two days can be ensured.
//
//   - You can query the usage data that is recorded in the last quarter.
//
// @param request - QueryUserOmsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserOmsDataResponse
func (client *Client) QueryUserOmsDataWithOptions(request *QueryUserOmsDataRequest, runtime *dara.RuntimeOptions) (_result *QueryUserOmsDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Table) {
		query["Table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserOmsData"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserOmsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage data of an Alibaba Cloud service.
//
// Description:
//
// You can call this operation to query the usage data of an Alibaba Cloud service. Take note of the following items:
//
//   - The service code that you specify for querying the usage data of a specific Alibaba Cloud service must be valid. You can query the usage data by hour or by day.
//
//   - The time that you specify must follow the ISO8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
//
//   - Latency exists in data pushes. Therefore, if you set the DataType parameter to Hour, the integrity of usage data recorded in the last 24 hours can be ensured. If you set the DataType parameter to Day, the integrity of usage data recorded in the last two days can be ensured.
//
//   - You can query the usage data that is recorded in the last quarter.
//
// @param request - QueryUserOmsDataRequest
//
// @return QueryUserOmsDataResponse
func (client *Client) QueryUserOmsData(request *QueryUserOmsDataRequest) (_result *QueryUserOmsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserOmsDataResponse{}
	_body, _err := client.QueryUserOmsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unsubscribes from an instance that is no longer needed.
//
// Description:
//
// 1.  Refunds are applicable only for the actual paid amount. Vouchers used for the purchase are non-refundable.
//
// 2.  Check the information about unsubscription and confirm the unsubscription terms and refundable amount. The resource that is unsubscribed cannot be restored.
//
// 3.  For more information, see [Rules for unsubscribing from resources](https://www.alibabacloud.com/help/en/user-center/refund-rules).
//
// @param request - RefundInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundInstanceResponse
func (client *Client) RefundInstanceWithOptions(request *RefundInstanceRequest, runtime *dara.RuntimeOptions) (_result *RefundInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ImmediatelyRelease) {
		query["ImmediatelyRelease"] = request.ImmediatelyRelease
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefundInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefundInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unsubscribes from an instance that is no longer needed.
//
// Description:
//
// 1.  Refunds are applicable only for the actual paid amount. Vouchers used for the purchase are non-refundable.
//
// 2.  Check the information about unsubscription and confirm the unsubscription terms and refundable amount. The resource that is unsubscribed cannot be restored.
//
// 3.  For more information, see [Rules for unsubscribing from resources](https://www.alibabacloud.com/help/en/user-center/refund-rules).
//
// @param request - RefundInstanceRequest
//
// @return RefundInstanceResponse
func (client *Client) RefundInstance(request *RefundInstanceRequest) (_result *RefundInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefundInstanceResponse{}
	_body, _err := client.RefundInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases instances by Virtual Network Operators (VNOs).
//
// Description:
//
// This operation is provided for only VNOs to release instances. If a non-specific VNO calls this operation, the request is blocked.
//
// @param request - ReleaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RenewStatus) {
		query["RenewStatus"] = request.RenewStatus
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases instances by Virtual Network Operators (VNOs).
//
// Description:
//
// This operation is provided for only VNOs to release instances. If a non-specific VNO calls this operation, the request is blocked.
//
// @param request - ReleaseInstanceRequest
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates a financial relationship between the management account and a member.
//
// @param request - RelieveAccountRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RelieveAccountRelationResponse
func (client *Client) RelieveAccountRelationWithOptions(request *RelieveAccountRelationRequest, runtime *dara.RuntimeOptions) (_result *RelieveAccountRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChildUserId) {
		query["ChildUserId"] = request.ChildUserId
	}

	if !dara.IsNil(request.ParentUserId) {
		query["ParentUserId"] = request.ParentUserId
	}

	if !dara.IsNil(request.RelationId) {
		query["RelationId"] = request.RelationId
	}

	if !dara.IsNil(request.RelationType) {
		query["RelationType"] = request.RelationType
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RelieveAccountRelation"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RelieveAccountRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a financial relationship between the management account and a member.
//
// @param request - RelieveAccountRelationRequest
//
// @return RelieveAccountRelationResponse
func (client *Client) RelieveAccountRelation(request *RelieveAccountRelationRequest) (_result *RelieveAccountRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RelieveAccountRelationResponse{}
	_body, _err := client.RelieveAccountRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 续费变配接口
//
// @param request - RenewChangeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewChangeInstanceResponse
func (client *Client) RenewChangeInstanceWithOptions(request *RenewChangeInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewChangeInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Parameter) {
		query["Parameter"] = request.Parameter
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RenewPeriod) {
		query["RenewPeriod"] = request.RenewPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewChangeInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewChangeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 续费变配接口
//
// @param request - RenewChangeInstanceRequest
//
// @return RenewChangeInstanceResponse
func (client *Client) RenewChangeInstance(request *RenewChangeInstanceRequest) (_result *RenewChangeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewChangeInstanceResponse{}
	_body, _err := client.RenewChangeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a specified instance. You cannot call this operation to renew Elastic Compute Service (ECS) instances, ApsaraDB RDS instances, or ApsaraDB for Redis instances. To renew these types of instances, call the dedicated operation of the corresponding service.
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RenewPeriod) {
		query["RenewPeriod"] = request.RenewPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a specified instance. You cannot call this operation to renew Elastic Compute Service (ECS) instances, ApsaraDB RDS instances, or ApsaraDB for Redis instances. To renew these types of instances, call the dedicated operation of the corresponding service.
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a resource plan.
//
// @param request - RenewResourcePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewResourcePackageResponse
func (client *Client) RenewResourcePackageWithOptions(request *RenewResourcePackageRequest, runtime *dara.RuntimeOptions) (_result *RenewResourcePackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EffectiveDate) {
		query["EffectiveDate"] = request.EffectiveDate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewResourcePackage"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a resource plan.
//
// @param request - RenewResourcePackageRequest
//
// @return RenewResourcePackageResponse
func (client *Client) RenewResourcePackage(request *RenewResourcePackageRequest) (_result *RenewResourcePackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewResourcePackageResponse{}
	_body, _err := client.RenewResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets an expiration date for all Elastic Compute Service (ECS) instances.
//
// @param request - SetAllExpirationDayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAllExpirationDayResponse
func (client *Client) SetAllExpirationDayWithOptions(request *SetAllExpirationDayRequest, runtime *dara.RuntimeOptions) (_result *SetAllExpirationDayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.UnifyExpireDay) {
		query["UnifyExpireDay"] = request.UnifyExpireDay
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAllExpirationDay"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAllExpirationDayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets an expiration date for all Elastic Compute Service (ECS) instances.
//
// @param request - SetAllExpirationDayRequest
//
// @return SetAllExpirationDayResponse
func (client *Client) SetAllExpirationDay(request *SetAllExpirationDayRequest) (_result *SetAllExpirationDayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAllExpirationDayResponse{}
	_body, _err := client.SetAllExpirationDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables auto-renewal for an instance.
//
// @param request - SetRenewalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRenewalResponse
func (client *Client) SetRenewalWithOptions(request *SetRenewalRequest, runtime *dara.RuntimeOptions) (_result *SetRenewalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIDs) {
		query["InstanceIDs"] = request.InstanceIDs
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RenewalPeriod) {
		query["RenewalPeriod"] = request.RenewalPeriod
	}

	if !dara.IsNil(request.RenewalPeriodUnit) {
		query["RenewalPeriodUnit"] = request.RenewalPeriodUnit
	}

	if !dara.IsNil(request.RenewalStatus) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	if !dara.IsNil(request.SubscriptionType) {
		query["SubscriptionType"] = request.SubscriptionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetRenewal"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetRenewalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables auto-renewal for an instance.
//
// @param request - SetRenewalRequest
//
// @return SetRenewalResponse
func (client *Client) SetRenewal(request *SetRenewalRequest) (_result *SetRenewalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetRenewalResponse{}
	_body, _err := client.SetRenewalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetResellerUserAlarmThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetResellerUserAlarmThresholdResponse
func (client *Client) SetResellerUserAlarmThresholdWithOptions(request *SetResellerUserAlarmThresholdRequest, runtime *dara.RuntimeOptions) (_result *SetResellerUserAlarmThresholdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmThresholds) {
		query["AlarmThresholds"] = request.AlarmThresholds
	}

	if !dara.IsNil(request.AlarmType) {
		query["AlarmType"] = request.AlarmType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetResellerUserAlarmThreshold"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetResellerUserAlarmThresholdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetResellerUserAlarmThresholdRequest
//
// @return SetResellerUserAlarmThresholdResponse
func (client *Client) SetResellerUserAlarmThreshold(request *SetResellerUserAlarmThresholdRequest) (_result *SetResellerUserAlarmThresholdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetResellerUserAlarmThresholdResponse{}
	_body, _err := client.SetResellerUserAlarmThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the quota ledger and consumption ledger.
//
// @param request - SetResellerUserQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetResellerUserQuotaResponse
func (client *Client) SetResellerUserQuotaWithOptions(request *SetResellerUserQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetResellerUserQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.Currency) {
		query["Currency"] = request.Currency
	}

	if !dara.IsNil(request.OutBizId) {
		query["OutBizId"] = request.OutBizId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetResellerUserQuota"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetResellerUserQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the quota ledger and consumption ledger.
//
// @param request - SetResellerUserQuotaRequest
//
// @return SetResellerUserQuotaResponse
func (client *Client) SetResellerUserQuota(request *SetResellerUserQuotaRequest) (_result *SetResellerUserQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetResellerUserQuotaResponse{}
	_body, _err := client.SetResellerUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetResellerUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetResellerUserStatusResponse
func (client *Client) SetResellerUserStatusWithOptions(request *SetResellerUserStatusRequest, runtime *dara.RuntimeOptions) (_result *SetResellerUserStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StopMode) {
		query["StopMode"] = request.StopMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetResellerUserStatus"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetResellerUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetResellerUserStatusRequest
//
// @return SetResellerUserStatusResponse
func (client *Client) SetResellerUserStatus(request *SetResellerUserStatusRequest) (_result *SetResellerUserStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetResellerUserStatusResponse{}
	_body, _err := client.SetResellerUserStatusWithOptions(request, runtime)
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
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Version:     dara.String("2017-12-14"),
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

// Summary:
//
// Subscribes to the bills that are stored in Object Storage Service (OSS) buckets.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - You can subscribe to only one type of bill at a time.
//
//   - The bills generated on the previous day are pushed on a daily basis the next day after you subscribe to the bills. The full-data bills for the previous month are pushed on the fourth day of each month. The monthly bills in the PDF format for the previous month are pushed on the fourth day of each month.
//
//   - The daily bills may be delayed. The delayed bills are pushed the next day after they are generated. The delayed bills may include the bills that should have been pushed on the previous day. We recommend that you query the full-data bills for the previous month at the beginning of each month.
//
//   - The bill subscriber must have the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
//
//   - The SubscribeBillToOSS operation has the same functionality as the Save Expense Details to OSS Bucket feature in User Center.
//
//   - To subscribe to the bills stored in an OSS bucket, make sure that the directory name specified for the OSS bucket conforms to the following naming rules:
//
// 1.  1.  The directory name can contain only UTF-8 characters and cannot contain emoticons.
//
// 2.  2.  Forward slashes (/) are used to separate paths and can be used to create subdirectories with ease. The directory name cannot start with a forward slash (/), a backslash (\\\\), or consecutive forward slashes (/).
//
// 3.  3.  The name of a subdirectory cannot be set to two consecutive periods (..).
//
// 4.  4.  The directory name must be 1 to 254 characters in length.
//
//   - File names:
//
//   - **BillingItemDetailForBillingPeriod*	- (Detailed bills of billable items)
//
//   - File name format for a daily push: `UID_BillingItemDetail_YYYYMMDD`. Example: `169**_BillingItemDetail_20190310`.
//
//   - File name format for a full-data push at the beginning of the next month: `UID_BillingItemDetail_YYYYMM`. Example: `169**_BillingItemDetail_201903`.
//
//   - **InstanceDetailForBillingPeriod*	- (Detailed bills of instances)
//
//   - File name format for a daily push: `UID_InstanceDetail_YYYYMMDD`. Example: `169**_InstanceDetail_20190310`.
//
//   - File name format for a full-data push at the beginning of the next month: `UID_InstanceDetail_YYYYMM`. Example: `169**_InstanceDetail_201903`.
//
//   - **InstanceDetailMonthly*	- (Instance-based bills summarized by billing cycle)
//
//   - File name format for a daily push: `UID_InstanceDetailMonthly_YYYYMM`. Example: `169**_InstanceDetailMonthly_201903`. A bill of this type contains the full data generated from the beginning of the month to the current day, and is updated every day until the fourth day of the next month.
//
//   - **BillingItemDetailMonthly*	- (Billable item-based bills summarized by billing cycle)
//
//   - File name format for a daily push: `UID_BillingItemDetailMonthly_YYYYMM`. Example: `169**_BillingItemDetailMonthly_201903`. A bill of this type contains the full data generated from the beginning of the month to the current day, and is updated every day until the fourth day of the next month.
//
//   - **SplitItemDetailDaily*	- (Split bills summarized by day)
//
//   - File name format for a daily push: `UID_SplitItemDetailDaily_YYYYMM`. Example: `169**_SplitItemDetailDaily_201903`. A bill of this type contains the full data generated from the beginning of the month to the current day, and is updated every day until the fourth day of the next month.
//
//   - **MonthBill*	- (Monthly bill in the PDF format)
//
//   - File name format for a monthly push: `UID_MonthBill_YYYYMM`. Example: `169**_MonthBill_201903`. The bill for the previous month is pushed on the fourth day of each month.
//
//   - The bills of the MonthBill type are PDF files, whereas the bills of other types are CSV files. If the number of data rows in a bill exceeds a threshold, the bill is automatically split into multiple CSV files. Then, the multiple CSV files are automatically merged and compressed into a ZIP file that has the same name format as the original file.
//
// @param request - SubscribeBillToOSSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeBillToOSSResponse
func (client *Client) SubscribeBillToOSSWithOptions(request *SubscribeBillToOSSRequest, runtime *dara.RuntimeOptions) (_result *SubscribeBillToOSSResponse, _err error) {
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

	if !dara.IsNil(request.BucketOwnerId) {
		query["BucketOwnerId"] = request.BucketOwnerId
	}

	if !dara.IsNil(request.BucketPath) {
		query["BucketPath"] = request.BucketPath
	}

	if !dara.IsNil(request.MultAccountRelSubscribe) {
		query["MultAccountRelSubscribe"] = request.MultAccountRelSubscribe
	}

	if !dara.IsNil(request.RowLimitPerFile) {
		query["RowLimitPerFile"] = request.RowLimitPerFile
	}

	if !dara.IsNil(request.SubscribeBucket) {
		query["SubscribeBucket"] = request.SubscribeBucket
	}

	if !dara.IsNil(request.SubscribeType) {
		query["SubscribeType"] = request.SubscribeType
	}

	if !dara.IsNil(request.UsingSsl) {
		query["UsingSsl"] = request.UsingSsl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubscribeBillToOSS"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubscribeBillToOSSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Subscribes to the bills that are stored in Object Storage Service (OSS) buckets.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - You can subscribe to only one type of bill at a time.
//
//   - The bills generated on the previous day are pushed on a daily basis the next day after you subscribe to the bills. The full-data bills for the previous month are pushed on the fourth day of each month. The monthly bills in the PDF format for the previous month are pushed on the fourth day of each month.
//
//   - The daily bills may be delayed. The delayed bills are pushed the next day after they are generated. The delayed bills may include the bills that should have been pushed on the previous day. We recommend that you query the full-data bills for the previous month at the beginning of each month.
//
//   - The bill subscriber must have the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission.
//
//   - The SubscribeBillToOSS operation has the same functionality as the Save Expense Details to OSS Bucket feature in User Center.
//
//   - To subscribe to the bills stored in an OSS bucket, make sure that the directory name specified for the OSS bucket conforms to the following naming rules:
//
// 1.  1.  The directory name can contain only UTF-8 characters and cannot contain emoticons.
//
// 2.  2.  Forward slashes (/) are used to separate paths and can be used to create subdirectories with ease. The directory name cannot start with a forward slash (/), a backslash (\\\\), or consecutive forward slashes (/).
//
// 3.  3.  The name of a subdirectory cannot be set to two consecutive periods (..).
//
// 4.  4.  The directory name must be 1 to 254 characters in length.
//
//   - File names:
//
//   - **BillingItemDetailForBillingPeriod*	- (Detailed bills of billable items)
//
//   - File name format for a daily push: `UID_BillingItemDetail_YYYYMMDD`. Example: `169**_BillingItemDetail_20190310`.
//
//   - File name format for a full-data push at the beginning of the next month: `UID_BillingItemDetail_YYYYMM`. Example: `169**_BillingItemDetail_201903`.
//
//   - **InstanceDetailForBillingPeriod*	- (Detailed bills of instances)
//
//   - File name format for a daily push: `UID_InstanceDetail_YYYYMMDD`. Example: `169**_InstanceDetail_20190310`.
//
//   - File name format for a full-data push at the beginning of the next month: `UID_InstanceDetail_YYYYMM`. Example: `169**_InstanceDetail_201903`.
//
//   - **InstanceDetailMonthly*	- (Instance-based bills summarized by billing cycle)
//
//   - File name format for a daily push: `UID_InstanceDetailMonthly_YYYYMM`. Example: `169**_InstanceDetailMonthly_201903`. A bill of this type contains the full data generated from the beginning of the month to the current day, and is updated every day until the fourth day of the next month.
//
//   - **BillingItemDetailMonthly*	- (Billable item-based bills summarized by billing cycle)
//
//   - File name format for a daily push: `UID_BillingItemDetailMonthly_YYYYMM`. Example: `169**_BillingItemDetailMonthly_201903`. A bill of this type contains the full data generated from the beginning of the month to the current day, and is updated every day until the fourth day of the next month.
//
//   - **SplitItemDetailDaily*	- (Split bills summarized by day)
//
//   - File name format for a daily push: `UID_SplitItemDetailDaily_YYYYMM`. Example: `169**_SplitItemDetailDaily_201903`. A bill of this type contains the full data generated from the beginning of the month to the current day, and is updated every day until the fourth day of the next month.
//
//   - **MonthBill*	- (Monthly bill in the PDF format)
//
//   - File name format for a monthly push: `UID_MonthBill_YYYYMM`. Example: `169**_MonthBill_201903`. The bill for the previous month is pushed on the fourth day of each month.
//
//   - The bills of the MonthBill type are PDF files, whereas the bills of other types are CSV files. If the number of data rows in a bill exceeds a threshold, the bill is automatically split into multiple CSV files. Then, the multiple CSV files are automatically merged and compressed into a ZIP file that has the same name format as the original file.
//
// @param request - SubscribeBillToOSSRequest
//
// @return SubscribeBillToOSSResponse
func (client *Client) SubscribeBillToOSS(request *SubscribeBillToOSSRequest) (_result *SubscribeBillToOSSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubscribeBillToOSSResponse{}
	_body, _err := client.SubscribeBillToOSSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add tags to resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add tags to resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unsubscribes from the bills that are stored in Object Storage Service (OSS) buckets.
//
// @param request - UnsubscribeBillToOSSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnsubscribeBillToOSSResponse
func (client *Client) UnsubscribeBillToOSSWithOptions(request *UnsubscribeBillToOSSRequest, runtime *dara.RuntimeOptions) (_result *UnsubscribeBillToOSSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MultAccountRelSubscribe) {
		query["MultAccountRelSubscribe"] = request.MultAccountRelSubscribe
	}

	if !dara.IsNil(request.SubscribeType) {
		query["SubscribeType"] = request.SubscribeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnsubscribeBillToOSS"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnsubscribeBillToOSSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unsubscribes from the bills that are stored in Object Storage Service (OSS) buckets.
//
// @param request - UnsubscribeBillToOSSRequest
//
// @return UnsubscribeBillToOSSResponse
func (client *Client) UnsubscribeBillToOSS(request *UnsubscribeBillToOSSRequest) (_result *UnsubscribeBillToOSSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnsubscribeBillToOSSResponse{}
	_body, _err := client.UnsubscribeBillToOSSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades a resource plan.
//
// @param request - UpgradeResourcePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeResourcePackageResponse
func (client *Client) UpgradeResourcePackageWithOptions(request *UpgradeResourcePackageRequest, runtime *dara.RuntimeOptions) (_result *UpgradeResourcePackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveDate) {
		query["EffectiveDate"] = request.EffectiveDate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeResourcePackage"),
		Version:     dara.String("2017-12-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a resource plan.
//
// @param request - UpgradeResourcePackageRequest
//
// @return UpgradeResourcePackageResponse
func (client *Client) UpgradeResourcePackage(request *UpgradeResourcePackageRequest) (_result *UpgradeResourcePackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeResourcePackageResponse{}
	_body, _err := client.UpgradeResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
