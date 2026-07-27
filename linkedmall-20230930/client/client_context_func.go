// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) CancelRefundOrderWithContext(ctx context.Context, disputeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelRefundOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmDisburseResponse
func (client *Client) ConfirmDisburseWithContext(ctx context.Context, request *ConfirmDisburseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConfirmDisburseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGoodsShippingNoticeResponse
func (client *Client) CreateGoodsShippingNoticeWithContext(ctx context.Context, request *CreateGoodsShippingNoticeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGoodsShippingNoticeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePurchaseOrderResponse
func (client *Client) CreatePurchaseOrderWithContext(ctx context.Context, request *CreatePurchaseOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePurchaseOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRefundOrderResponse
func (client *Client) CreateRefundOrderWithContext(ctx context.Context, request *CreateRefundOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRefundOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderResponse
func (client *Client) GetOrderWithContext(ctx context.Context, orderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurchaseOrderStatusResponse
func (client *Client) GetPurchaseOrderStatusWithContext(ctx context.Context, purchaseOrderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPurchaseOrderStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurchaserShopResponse
func (client *Client) GetPurchaserShopWithContext(ctx context.Context, purchaserId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPurchaserShopResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRefundOrderResponse
func (client *Client) GetRefundOrderWithContext(ctx context.Context, disputeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRefundOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSelectionProductResponse
func (client *Client) GetSelectionProductWithContext(ctx context.Context, productId *string, request *GetSelectionProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSelectionProductResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSelectionProductSaleInfoResponse
func (client *Client) GetSelectionProductSaleInfoWithContext(ctx context.Context, productId *string, request *GetSelectionProductSaleInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSelectionProductSaleInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoriesResponse
func (client *Client) ListCategoriesWithContext(ctx context.Context, request *ListCategoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCategoriesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogisticsOrdersResponse
func (client *Client) ListLogisticsOrdersWithContext(ctx context.Context, orderId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListLogisticsOrdersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPurchaserShopsResponse
func (client *Client) ListPurchaserShopsWithContext(ctx context.Context, request *ListPurchaserShopsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPurchaserShopsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionProductSaleInfosResponse
func (client *Client) ListSelectionProductSaleInfosWithContext(ctx context.Context, request *ListSelectionProductSaleInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSelectionProductSaleInfosResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionProductsResponse
func (client *Client) ListSelectionProductsWithContext(ctx context.Context, request *ListSelectionProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSelectionProductsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionSkuSaleInfosResponse
func (client *Client) ListSelectionSkuSaleInfosWithContext(ctx context.Context, request *ListSelectionSkuSaleInfosRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSelectionSkuSaleInfosResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChildDivisionCodeResponse
func (client *Client) QueryChildDivisionCodeWithContext(ctx context.Context, request *QueryChildDivisionCodeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryChildDivisionCodeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrdersResponse
func (client *Client) QueryOrdersWithContext(ctx context.Context, request *QueryOrdersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryOrdersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenderPurchaseOrderResponse
func (client *Client) RenderPurchaseOrderWithContext(ctx context.Context, request *RenderPurchaseOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenderPurchaseOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenderRefundOrderResponse
func (client *Client) RenderRefundOrderWithContext(ctx context.Context, request *RenderRefundOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenderRefundOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchProductsResponse
func (client *Client) SearchProductsWithContext(ctx context.Context, request *SearchProductsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchProductsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectionGroupAddProductResponse
func (client *Client) SelectionGroupAddProductWithContext(ctx context.Context, request *SelectionGroupAddProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectionGroupAddProductResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectionGroupRemoveProductResponse
func (client *Client) SelectionGroupRemoveProductWithContext(ctx context.Context, request *SelectionGroupRemoveProductRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectionGroupRemoveProductResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SplitPurchaseOrderResponse
func (client *Client) SplitPurchaseOrderWithContext(ctx context.Context, request *SplitPurchaseOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SplitPurchaseOrderResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
