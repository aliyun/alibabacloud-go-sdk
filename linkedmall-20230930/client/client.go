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
		"cn-hangzhou":                 dara.String("linkedmall.aliyuncs.com"),
		"cn-shanghai":                 dara.String("linkedmall.aliyuncs.com"),
		"ap-northeast-1":              dara.String("linkedmall.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("linkedmall.aliyuncs.com"),
		"ap-south-1":                  dara.String("linkedmall.aliyuncs.com"),
		"ap-southeast-1":              dara.String("linkedmall.aliyuncs.com"),
		"ap-southeast-2":              dara.String("linkedmall.aliyuncs.com"),
		"ap-southeast-3":              dara.String("linkedmall.aliyuncs.com"),
		"ap-southeast-5":              dara.String("linkedmall.aliyuncs.com"),
		"cn-beijing":                  dara.String("linkedmall.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("linkedmall.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("linkedmall.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("linkedmall.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("linkedmall.aliyuncs.com"),
		"cn-chengdu":                  dara.String("linkedmall.aliyuncs.com"),
		"cn-edge-1":                   dara.String("linkedmall.aliyuncs.com"),
		"cn-fujian":                   dara.String("linkedmall.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("linkedmall.aliyuncs.com"),
		"cn-hongkong":                 dara.String("linkedmall.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("linkedmall.aliyuncs.com"),
		"cn-huhehaote":                dara.String("linkedmall.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("linkedmall.aliyuncs.com"),
		"cn-qingdao":                  dara.String("linkedmall.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("linkedmall.aliyuncs.com"),
		"cn-wuhan":                    dara.String("linkedmall.aliyuncs.com"),
		"cn-yushanfang":               dara.String("linkedmall.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("linkedmall.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("linkedmall.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("linkedmall.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("linkedmall.aliyuncs.com"),
		"eu-central-1":                dara.String("linkedmall.aliyuncs.com"),
		"eu-west-1":                   dara.String("linkedmall.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("linkedmall.aliyuncs.com"),
		"me-east-1":                   dara.String("linkedmall.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("linkedmall.aliyuncs.com"),
		"us-east-1":                   dara.String("linkedmall.aliyuncs.com"),
		"us-west-1":                   dara.String("linkedmall.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("linkedmall"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Cancels a refund order.
//
// Description:
//
// Cancel a refund order.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRefundOrderResponse
func (client *Client) CancelRefundOrderWithOptions(disputeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelRefundOrderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelRefundOrder"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds/" + dara.PercentEncode(dara.StringValue(disputeId)) + "/commands/cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a refund order.
//
// Description:
//
// Cancel a refund order.
//
// @return CancelRefundOrderResponse
func (client *Client) CancelRefundOrder(disputeId *string) (_result *CancelRefundOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelRefundOrderResponse{}
	_body, _err := client.CancelRefundOrderWithOptions(disputeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Confirms the receipt of goods.
//
// Description:
//
// Confirms the receipt of goods.
//
// @param request - ConfirmDisburseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmDisburseResponse
func (client *Client) ConfirmDisburseWithOptions(request *ConfirmDisburseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConfirmDisburseResponse, _err error) {
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
		Action:      dara.String("ConfirmDisburse"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/orders/commands/confirmDisburse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmDisburseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Confirms the receipt of goods.
//
// Description:
//
// Confirms the receipt of goods.
//
// @param request - ConfirmDisburseRequest
//
// @return ConfirmDisburseResponse
func (client *Client) ConfirmDisburse(request *ConfirmDisburseRequest) (_result *ConfirmDisburseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConfirmDisburseResponse{}
	_body, _err := client.ConfirmDisburseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Backfill shipping notice information.
//
// Description:
//
// Backfill shipping notice information.
//
// @param request - CreateGoodsShippingNoticeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGoodsShippingNoticeResponse
func (client *Client) CreateGoodsShippingNoticeWithOptions(request *CreateGoodsShippingNoticeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGoodsShippingNoticeResponse, _err error) {
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
		Action:      dara.String("CreateGoodsShippingNotice"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds/command/createGoodsShippingNotice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGoodsShippingNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Backfill shipping notice information.
//
// Description:
//
// Backfill shipping notice information.
//
// @param request - CreateGoodsShippingNoticeRequest
//
// @return CreateGoodsShippingNoticeResponse
func (client *Client) CreateGoodsShippingNotice(request *CreateGoodsShippingNoticeRequest) (_result *CreateGoodsShippingNoticeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGoodsShippingNoticeResponse{}
	_body, _err := client.CreateGoodsShippingNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a purchase order and returns the purchase order ID. The specific result of order creation is communicated through messages. After the order is created, you can query the order details associated with the purchase order using the order API.
//
// Description:
//
// Creates a purchase order and returns the purchase order ID. Messages communicate the specific result of order creation. After the order is created, you can query the order details associated with the purchase order using the order API.
//
//	Warning: Note: Purchase order creation is an asynchronous task. If a distributor calls this API and receives an abnormal status (such as error code 503), do not immediately process customer refunds. Distributors must wait for and consume the PurchaseOrderCreate message (the purchase order creation result message) to determine the order status—for example, by consuming the order status synchronization message—before proceeding with business logic. This prevents financial losses.
//
//	Notice: Note: If you do not receive the PurchaseOrderCreate message (the purchase order creation result message) after calling the purchase order creation API, submit a ticket to the technical support team to inquire about the cause.
//
// @param request - CreatePurchaseOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePurchaseOrderResponse
func (client *Client) CreatePurchaseOrderWithOptions(request *CreatePurchaseOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePurchaseOrderResponse, _err error) {
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
		Action:      dara.String("CreatePurchaseOrder"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaseOrders"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePurchaseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a purchase order and returns the purchase order ID. The specific result of order creation is communicated through messages. After the order is created, you can query the order details associated with the purchase order using the order API.
//
// Description:
//
// Creates a purchase order and returns the purchase order ID. Messages communicate the specific result of order creation. After the order is created, you can query the order details associated with the purchase order using the order API.
//
//	Warning: Note: Purchase order creation is an asynchronous task. If a distributor calls this API and receives an abnormal status (such as error code 503), do not immediately process customer refunds. Distributors must wait for and consume the PurchaseOrderCreate message (the purchase order creation result message) to determine the order status—for example, by consuming the order status synchronization message—before proceeding with business logic. This prevents financial losses.
//
//	Notice: Note: If you do not receive the PurchaseOrderCreate message (the purchase order creation result message) after calling the purchase order creation API, submit a ticket to the technical support team to inquire about the cause.
//
// @param request - CreatePurchaseOrderRequest
//
// @return CreatePurchaseOrderResponse
func (client *Client) CreatePurchaseOrder(request *CreatePurchaseOrderRequest) (_result *CreatePurchaseOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePurchaseOrderResponse{}
	_body, _err := client.CreatePurchaseOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a support ticket.
//
// Description:
//
// Creates a refund order.
//
// @param request - CreateRefundOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRefundOrderResponse
func (client *Client) CreateRefundOrderWithOptions(request *CreateRefundOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRefundOrderResponse, _err error) {
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
		Action:      dara.String("CreateRefundOrder"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a support ticket.
//
// Description:
//
// Creates a refund order.
//
// @param request - CreateRefundOrderRequest
//
// @return CreateRefundOrderResponse
func (client *Client) CreateRefundOrder(request *CreateRefundOrderRequest) (_result *CreateRefundOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRefundOrderResponse{}
	_body, _err := client.CreateRefundOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an order.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderResponse
func (client *Client) GetOrderWithOptions(orderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrder"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/orders/" + dara.PercentEncode(dara.StringValue(orderId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an order.
//
// @return GetOrderResponse
func (client *Client) GetOrder(orderId *string) (_result *GetOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrderResponse{}
	_body, _err := client.GetOrderWithOptions(orderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the purchase order status.
//
// Description:
//
// Retrieve the transaction order status.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurchaseOrderStatusResponse
func (client *Client) GetPurchaseOrderStatusWithOptions(purchaseOrderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPurchaseOrderStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPurchaseOrderStatus"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaseOrders/" + dara.PercentEncode(dara.StringValue(purchaseOrderId)) + "/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPurchaseOrderStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the purchase order status.
//
// Description:
//
// Retrieve the transaction order status.
//
// @return GetPurchaseOrderStatusResponse
func (client *Client) GetPurchaseOrderStatus(purchaseOrderId *string) (_result *GetPurchaseOrderStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPurchaseOrderStatusResponse{}
	_body, _err := client.GetPurchaseOrderStatusWithOptions(purchaseOrderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the purchaser\\"s shop.
//
// Description:
//
// Retrieves the purchaser\\"s shop.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurchaserShopResponse
func (client *Client) GetPurchaserShopWithOptions(purchaserId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPurchaserShopResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPurchaserShop"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaserShops/" + dara.PercentEncode(dara.StringValue(purchaserId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPurchaserShopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the purchaser\\"s shop.
//
// Description:
//
// Retrieves the purchaser\\"s shop.
//
// @return GetPurchaserShopResponse
func (client *Client) GetPurchaserShop(purchaserId *string) (_result *GetPurchaserShopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPurchaserShopResponse{}
	_body, _err := client.GetPurchaserShopWithOptions(purchaserId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve details of an after-sales order.
//
// Description:
//
// # Retrieve after-sales order details
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRefundOrderResponse
func (client *Client) GetRefundOrderWithOptions(disputeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRefundOrderResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRefundOrder"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds/" + dara.PercentEncode(dara.StringValue(disputeId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve details of an after-sales order.
//
// Description:
//
// # Retrieve after-sales order details
//
// @return GetRefundOrderResponse
func (client *Client) GetRefundOrder(disputeId *string) (_result *GetRefundOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRefundOrderResponse{}
	_body, _err := client.GetRefundOrderWithOptions(disputeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the details of a product in the selection pool.
//
// Description:
//
// Retrieve product details from the selection pool using the product ID. You can also specify a region code to check regional inventory.
//
// @param request - GetSelectionProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSelectionProductResponse
func (client *Client) GetSelectionProductWithOptions(productId *string, request *GetSelectionProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSelectionProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DivisionCode) {
		query["divisionCode"] = request.DivisionCode
	}

	if !dara.IsNil(request.PurchaserId) {
		query["purchaserId"] = request.PurchaserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSelectionProduct"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/products/" + dara.PercentEncode(dara.StringValue(productId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSelectionProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the details of a product in the selection pool.
//
// Description:
//
// Retrieve product details from the selection pool using the product ID. You can also specify a region code to check regional inventory.
//
// @param request - GetSelectionProductRequest
//
// @return GetSelectionProductResponse
func (client *Client) GetSelectionProduct(productId *string, request *GetSelectionProductRequest) (_result *GetSelectionProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSelectionProductResponse{}
	_body, _err := client.GetSelectionProductWithOptions(productId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries sales information for products in the selection pool.
//
// Description:
//
// Queries sales information for products in the selection pool. Distributors can call this operation to check product sales details, such as product status. Use the divisionCode input parameter to check whether a product is available for sale in a specific region. We recommend using a five-level administrative division code (township or subdistrict level).
//
// @param request - GetSelectionProductSaleInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSelectionProductSaleInfoResponse
func (client *Client) GetSelectionProductSaleInfoWithOptions(productId *string, request *GetSelectionProductSaleInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSelectionProductSaleInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DivisionCode) {
		query["divisionCode"] = request.DivisionCode
	}

	if !dara.IsNil(request.PurchaserId) {
		query["purchaserId"] = request.PurchaserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSelectionProductSaleInfo"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/products/" + dara.PercentEncode(dara.StringValue(productId)) + "/saleInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSelectionProductSaleInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sales information for products in the selection pool.
//
// Description:
//
// Queries sales information for products in the selection pool. Distributors can call this operation to check product sales details, such as product status. Use the divisionCode input parameter to check whether a product is available for sale in a specific region. We recommend using a five-level administrative division code (township or subdistrict level).
//
// @param request - GetSelectionProductSaleInfoRequest
//
// @return GetSelectionProductSaleInfoResponse
func (client *Client) GetSelectionProductSaleInfo(productId *string, request *GetSelectionProductSaleInfoRequest) (_result *GetSelectionProductSaleInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSelectionProductSaleInfoResponse{}
	_body, _err := client.GetSelectionProductSaleInfoWithOptions(productId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists categories.
//
// Description:
//
// Retrieves all subcategories for a parent category ID, or the details for a specific category ID.
//
// If the parent category ID (parentCategoryId) is 0, the API returns the top-level categories under the root category.
//
// @param request - ListCategoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoriesResponse
func (client *Client) ListCategoriesWithOptions(request *ListCategoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCategoriesResponse, _err error) {
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
		Action:      dara.String("ListCategories"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/categories/commands/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists categories.
//
// Description:
//
// Retrieves all subcategories for a parent category ID, or the details for a specific category ID.
//
// If the parent category ID (parentCategoryId) is 0, the API returns the top-level categories under the root category.
//
// @param request - ListCategoriesRequest
//
// @return ListCategoriesResponse
func (client *Client) ListCategories(request *ListCategoriesRequest) (_result *ListCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCategoriesResponse{}
	_body, _err := client.ListCategoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query logistics information for an order.
//
// Description:
//
// Retrieves logistics information for an order.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogisticsOrdersResponse
func (client *Client) ListLogisticsOrdersWithOptions(orderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogisticsOrdersResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogisticsOrders"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/orders/" + dara.PercentEncode(dara.StringValue(orderId)) + "/logisticsOrders"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogisticsOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query logistics information for an order.
//
// Description:
//
// Retrieves logistics information for an order.
//
// @return ListLogisticsOrdersResponse
func (client *Client) ListLogisticsOrders(orderId *string) (_result *ListLogisticsOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogisticsOrdersResponse{}
	_body, _err := client.ListLogisticsOrdersWithOptions(orderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists purchaser shops.
//
// Description:
//
// Lists purchaser shops.
//
// @param request - ListPurchaserShopsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPurchaserShopsResponse
func (client *Client) ListPurchaserShopsWithOptions(request *ListPurchaserShopsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPurchaserShopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("ListPurchaserShops"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaserShops"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPurchaserShopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists purchaser shops.
//
// Description:
//
// Lists purchaser shops.
//
// @param request - ListPurchaserShopsRequest
//
// @return ListPurchaserShopsResponse
func (client *Client) ListPurchaserShops(request *ListPurchaserShopsRequest) (_result *ListPurchaserShopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPurchaserShopsResponse{}
	_body, _err := client.ListPurchaserShopsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query product sales information for the selection pool in batches.
//
// Description:
//
// You can query product sales information for the selection pool in batches. Distributors can call this operation to retrieve product sales details, such as product status. Use the divisionCode input parameter to check whether products are available for sale in a specific region. We recommend that you pass a five-level address code (town or street level).
//
// @param request - ListSelectionProductSaleInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionProductSaleInfosResponse
func (client *Client) ListSelectionProductSaleInfosWithOptions(request *ListSelectionProductSaleInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSelectionProductSaleInfosResponse, _err error) {
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
		Action:      dara.String("ListSelectionProductSaleInfos"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/products/saleInfo/commands/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSelectionProductSaleInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query product sales information for the selection pool in batches.
//
// Description:
//
// You can query product sales information for the selection pool in batches. Distributors can call this operation to retrieve product sales details, such as product status. Use the divisionCode input parameter to check whether products are available for sale in a specific region. We recommend that you pass a five-level address code (town or street level).
//
// @param request - ListSelectionProductSaleInfosRequest
//
// @return ListSelectionProductSaleInfosResponse
func (client *Client) ListSelectionProductSaleInfos(request *ListSelectionProductSaleInfosRequest) (_result *ListSelectionProductSaleInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSelectionProductSaleInfosResponse{}
	_body, _err := client.ListSelectionProductSaleInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of products from a product selection pool.
//
// @param request - ListSelectionProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionProductsResponse
func (client *Client) ListSelectionProductsWithOptions(request *ListSelectionProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSelectionProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PurchaserId) {
		query["purchaserId"] = request.PurchaserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSelectionProducts"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/products"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSelectionProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of products from a product selection pool.
//
// @param request - ListSelectionProductsRequest
//
// @return ListSelectionProductsResponse
func (client *Client) ListSelectionProducts(request *ListSelectionProductsRequest) (_result *ListSelectionProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSelectionProductsResponse{}
	_body, _err := client.ListSelectionProductsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query SKU sales information for items in the selection pool in batch.
//
// Description:
//
// Query SKU sales information for items in the selection pool in batch. Distributors can call this API to retrieve batch details about SKU sales status and other attributes. To determine whether SKUs are sellable in a specific region, use the divisionCode parameter—preferably a five-level administrative division code for townships or subdistricts.
//
// @param request - ListSelectionSkuSaleInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionSkuSaleInfosResponse
func (client *Client) ListSelectionSkuSaleInfosWithOptions(request *ListSelectionSkuSaleInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSelectionSkuSaleInfosResponse, _err error) {
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
		Action:      dara.String("ListSelectionSkuSaleInfos"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/skus/saleInfo/commands/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSelectionSkuSaleInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query SKU sales information for items in the selection pool in batch.
//
// Description:
//
// Query SKU sales information for items in the selection pool in batch. Distributors can call this API to retrieve batch details about SKU sales status and other attributes. To determine whether SKUs are sellable in a specific region, use the divisionCode parameter—preferably a five-level administrative division code for townships or subdistricts.
//
// @param request - ListSelectionSkuSaleInfosRequest
//
// @return ListSelectionSkuSaleInfosResponse
func (client *Client) ListSelectionSkuSaleInfos(request *ListSelectionSkuSaleInfosRequest) (_result *ListSelectionSkuSaleInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSelectionSkuSaleInfosResponse{}
	_body, _err := client.ListSelectionSkuSaleInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries child division codes.
//
// Description:
//
// Queries child division codes.
//
// @param request - QueryChildDivisionCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChildDivisionCodeResponse
func (client *Client) QueryChildDivisionCodeWithOptions(request *QueryChildDivisionCodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryChildDivisionCodeResponse, _err error) {
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
		Action:      dara.String("QueryChildDivisionCode"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/division/commands/queryChildDivisionCode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryChildDivisionCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries child division codes.
//
// Description:
//
// Queries child division codes.
//
// @param request - QueryChildDivisionCodeRequest
//
// @return QueryChildDivisionCodeResponse
func (client *Client) QueryChildDivisionCode(request *QueryChildDivisionCodeRequest) (_result *QueryChildDivisionCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryChildDivisionCodeResponse{}
	_body, _err := client.QueryChildDivisionCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists orders.
//
// Description:
//
// Lists orders.
//
// @param request - QueryOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrdersResponse
func (client *Client) QueryOrdersWithOptions(request *QueryOrdersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryOrdersResponse, _err error) {
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
		Action:      dara.String("QueryOrders"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/orders/commands/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Lists orders.
//
// Description:
//
// Lists orders.
//
// @param request - QueryOrdersRequest
//
// @return QueryOrdersResponse
func (client *Client) QueryOrders(request *QueryOrdersRequest) (_result *QueryOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryOrdersResponse{}
	_body, _err := client.QueryOrdersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renders a purchase order and returns both sellable and unsellable products. Customers can then select the sellable products to place their orders.
//
// Description:
//
//	Warning:
//
// This API will be offline soon. For purchase order rendering, use the SplitPurchaseOrder API, which supports both purchase order rendering and splitting.
//
// @param request - RenderPurchaseOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenderPurchaseOrderResponse
func (client *Client) RenderPurchaseOrderWithOptions(request *RenderPurchaseOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenderPurchaseOrderResponse, _err error) {
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
		Action:      dara.String("RenderPurchaseOrder"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaseOrders/commands/render"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenderPurchaseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renders a purchase order and returns both sellable and unsellable products. Customers can then select the sellable products to place their orders.
//
// Description:
//
//	Warning:
//
// This API will be offline soon. For purchase order rendering, use the SplitPurchaseOrder API, which supports both purchase order rendering and splitting.
//
// @param request - RenderPurchaseOrderRequest
//
// @return RenderPurchaseOrderResponse
func (client *Client) RenderPurchaseOrder(request *RenderPurchaseOrderRequest) (_result *RenderPurchaseOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenderPurchaseOrderResponse{}
	_body, _err := client.RenderPurchaseOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Reverse Single Rendering
//
// Description:
//
// Renders a refund order.
//
// @param request - RenderRefundOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenderRefundOrderResponse
func (client *Client) RenderRefundOrderWithOptions(request *RenderRefundOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenderRefundOrderResponse, _err error) {
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
		Action:      dara.String("RenderRefundOrder"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds/commands/render"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenderRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Reverse Single Rendering
//
// Description:
//
// Renders a refund order.
//
// @param request - RenderRefundOrderRequest
//
// @return RenderRefundOrderResponse
func (client *Client) RenderRefundOrder(request *RenderRefundOrderRequest) (_result *RenderRefundOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenderRefundOrderResponse{}
	_body, _err := client.RenderRefundOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The product search API is a paginated interface for searching products based on various criteria.
//
// @param request - SearchProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchProductsResponse
func (client *Client) SearchProductsWithOptions(request *SearchProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BrandName) {
		body["brandName"] = request.BrandName
	}

	if !dara.IsNil(request.CategoryIds) {
		body["categoryIds"] = request.CategoryIds
	}

	if !dara.IsNil(request.CreateEndTime) {
		body["createEndTime"] = request.CreateEndTime
	}

	if !dara.IsNil(request.CreateStartTime) {
		body["createStartTime"] = request.CreateStartTime
	}

	if !dara.IsNil(request.DistributionHighPrice) {
		body["distributionHighPrice"] = request.DistributionHighPrice
	}

	if !dara.IsNil(request.DistributionHighPriceRatio) {
		body["distributionHighPriceRatio"] = request.DistributionHighPriceRatio
	}

	if !dara.IsNil(request.DistributionLowPrice) {
		body["distributionLowPrice"] = request.DistributionLowPrice
	}

	if !dara.IsNil(request.DistributionLowPriceRatio) {
		body["distributionLowPriceRatio"] = request.DistributionLowPriceRatio
	}

	if !dara.IsNil(request.HighMarkPrice) {
		body["highMarkPrice"] = request.HighMarkPrice
	}

	if !dara.IsNil(request.HighPrice) {
		body["highPrice"] = request.HighPrice
	}

	if !dara.IsNil(request.InGroup) {
		body["inGroup"] = request.InGroup
	}

	if !dara.IsNil(request.InGroupEndTime) {
		body["inGroupEndTime"] = request.InGroupEndTime
	}

	if !dara.IsNil(request.InGroupStartTime) {
		body["inGroupStartTime"] = request.InGroupStartTime
	}

	if !dara.IsNil(request.InventoryRiskLevel) {
		body["inventoryRiskLevel"] = request.InventoryRiskLevel
	}

	if !dara.IsNil(request.LmItemId) {
		body["lmItemId"] = request.LmItemId
	}

	if !dara.IsNil(request.LowMarkPrice) {
		body["lowMarkPrice"] = request.LowMarkPrice
	}

	if !dara.IsNil(request.LowPrice) {
		body["lowPrice"] = request.LowPrice
	}

	if !dara.IsNil(request.ModifyEndTime) {
		body["modifyEndTime"] = request.ModifyEndTime
	}

	if !dara.IsNil(request.ModifyStartTime) {
		body["modifyStartTime"] = request.ModifyStartTime
	}

	if !dara.IsNil(request.OrderBy) {
		body["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDirection) {
		body["orderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.PageNumber) {
		body["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Platform) {
		body["platform"] = request.Platform
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductName) {
		body["productName"] = request.ProductName
	}

	if !dara.IsNil(request.ProductStatus) {
		body["productStatus"] = request.ProductStatus
	}

	if !dara.IsNil(request.PurchaserId) {
		body["purchaserId"] = request.PurchaserId
	}

	if !dara.IsNil(request.TaxRate) {
		body["taxRate"] = request.TaxRate
	}

	if !dara.IsNil(request.TradeModeAndCredit) {
		body["tradeModeAndCredit"] = request.TradeModeAndCredit
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchProducts"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/selection-group/product/command/searchProduct"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The product search API is a paginated interface for searching products based on various criteria.
//
// @param request - SearchProductsRequest
//
// @return SearchProductsResponse
func (client *Client) SearchProducts(request *SearchProductsRequest) (_result *SearchProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchProductsResponse{}
	_body, _err := client.SearchProductsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The distributor takes delivery of goods.
//
// Description:
//
// Distributors use this API to add products to their selection group.
//
// > We recommend that distributors who onboard on or after January 1, 2025 use this API. For more information about adding products and the related impact, see the [product best practices](https://help.aliyun.com/zh/linkedmall/user-guide/product-interface-best-practices?spm=a2c4g.11186623.help-menu-88587.d_2_2_0_8_0.58122056oN3crP\\&scm=20140722.H_2869668._.OR_help-T_cn~zh-V_1#lFENl).
//
// @param request - SelectionGroupAddProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectionGroupAddProductResponse
func (client *Client) SelectionGroupAddProductWithOptions(request *SelectionGroupAddProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectionGroupAddProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductIds) {
		body["productIds"] = request.ProductIds
	}

	if !dara.IsNil(request.PurchaserId) {
		body["purchaserId"] = request.PurchaserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectionGroupAddProduct"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/selection-group/product/command/addProduct"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectionGroupAddProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The distributor takes delivery of goods.
//
// Description:
//
// Distributors use this API to add products to their selection group.
//
// > We recommend that distributors who onboard on or after January 1, 2025 use this API. For more information about adding products and the related impact, see the [product best practices](https://help.aliyun.com/zh/linkedmall/user-guide/product-interface-best-practices?spm=a2c4g.11186623.help-menu-88587.d_2_2_0_8_0.58122056oN3crP\\&scm=20140722.H_2869668._.OR_help-T_cn~zh-V_1#lFENl).
//
// @param request - SelectionGroupAddProductRequest
//
// @return SelectionGroupAddProductResponse
func (client *Client) SelectionGroupAddProduct(request *SelectionGroupAddProductRequest) (_result *SelectionGroupAddProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectionGroupAddProductResponse{}
	_body, _err := client.SelectionGroupAddProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes products from a distributor\\"s stock.
//
// Description:
//
// Distributors use this API to remove products from their stock.
//
// @param request - SelectionGroupRemoveProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectionGroupRemoveProductResponse
func (client *Client) SelectionGroupRemoveProductWithOptions(request *SelectionGroupRemoveProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectionGroupRemoveProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductIds) {
		body["productIds"] = request.ProductIds
	}

	if !dara.IsNil(request.PurchaserId) {
		body["purchaserId"] = request.PurchaserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectionGroupRemoveProduct"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/selection-group/product/command/removeProduct"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectionGroupRemoveProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes products from a distributor\\"s stock.
//
// Description:
//
// Distributors use this API to remove products from their stock.
//
// @param request - SelectionGroupRemoveProductRequest
//
// @return SelectionGroupRemoveProductResponse
func (client *Client) SelectionGroupRemoveProduct(request *SelectionGroupRemoveProductRequest) (_result *SelectionGroupRemoveProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectionGroupRemoveProductResponse{}
	_body, _err := client.SelectionGroupRemoveProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Splits a purchase order and renders the resulting parent-child order structure. This API returns a list of items based on the final parent-child order structure. Distributors can use this response to render the final parent-child order layout, which simplifies receiving the purchase order creation success message and backfilling parent-child order information later.
//
// Description:
//
// Call this API before creating a purchase order. It returns two lists: one for sellable items and one for unsellable items. The sellable items list follows the final parent-child order split structure.
//
// @param request - SplitPurchaseOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SplitPurchaseOrderResponse
func (client *Client) SplitPurchaseOrderWithOptions(request *SplitPurchaseOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SplitPurchaseOrderResponse, _err error) {
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
		Action:      dara.String("SplitPurchaseOrder"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaseOrders/commands/split"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SplitPurchaseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Splits a purchase order and renders the resulting parent-child order structure. This API returns a list of items based on the final parent-child order structure. Distributors can use this response to render the final parent-child order layout, which simplifies receiving the purchase order creation success message and backfilling parent-child order information later.
//
// Description:
//
// Call this API before creating a purchase order. It returns two lists: one for sellable items and one for unsellable items. The sellable items list follows the final parent-child order split structure.
//
// @param request - SplitPurchaseOrderRequest
//
// @return SplitPurchaseOrderResponse
func (client *Client) SplitPurchaseOrder(request *SplitPurchaseOrderRequest) (_result *SplitPurchaseOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SplitPurchaseOrderResponse{}
	_body, _err := client.SplitPurchaseOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
