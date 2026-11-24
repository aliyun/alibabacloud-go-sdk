// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Activates an authorization code by calling ActivateLicense.
//
// @param request - ActivateLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicenseWithContext(ctx context.Context, request *ActivateLicenseRequest, runtime *dara.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Identification) {
		query["Identification"] = request.Identification
	}

	if !dara.IsNil(request.LicenseCode) {
		query["LicenseCode"] = request.LicenseCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateLicense"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets auto-renewal.
//
// @param request - AutoRenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AutoRenewInstanceResponse
func (client *Client) AutoRenewInstanceWithContext(ctx context.Context, request *AutoRenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *AutoRenewInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewCycle) {
		body["AutoRenewCycle"] = request.AutoRenewCycle
	}

	if !dara.IsNil(request.AutoRenewDuration) {
		body["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.OrderBizId) {
		body["OrderBizId"] = request.OrderBizId
	}

	if !dara.IsNil(request.OwnerId) {
		body["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AutoRenewInstance"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AutoRenewInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Confirms that the merchant has received and processed a notification message. After this call, the message will no longer be returned in subsequent queries.
//
// Description:
//
// Confirms the receipt of a subscription notification failure message.
//
// - After a successful call, the message will no longer be returned in subsequent queries. The platform considers that the merchant has correctly processed the message.
//
// @param request - ConfirmNotificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmNotificationResponse
func (client *Client) ConfirmNotificationWithContext(ctx context.Context, request *ConfirmNotificationRequest, runtime *dara.RuntimeOptions) (_result *ConfirmNotificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NotificationRequestId) {
		query["NotificationRequestId"] = request.NotificationRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmNotification"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmNotificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateOrder operation to create an order.
//
// Before calling this operation, make sure that you fully understand the billing method and pricing of Alibaba Cloud Marketplace products.
//
// To create an order for a commercial image, refer to [CreateInstance](https://help.aliyun.com/document_detail/25499.html).
//
// Description:
//
// Before calling this operation, make sure that you fully understand the billing method and [pricing](https://market.aliyun.com/) of Alibaba Cloud Marketplace products.
//
// To create an order for a commercial image, refer to [CreateInstance](https://help.aliyun.com/document_detail/25499.html).
//
// @param request - CreateOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderResponse
func (client *Client) CreateOrderWithContext(ctx context.Context, request *CreateOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
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

	if !dara.IsNil(request.Commodity) {
		query["Commodity"] = request.Commodity
	}

	if !dara.IsNil(request.OrderSouce) {
		query["OrderSouce"] = request.OrderSouce
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrder"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query cross-account role authorization information, including the authorized user\\"s Alibaba Cloud account UID, name, authorization time, and authorized role name.
//
// @param request - CrossAccountVerifyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CrossAccountVerifyTokenResponse
func (client *Client) CrossAccountVerifyTokenWithContext(ctx context.Context, request *CrossAccountVerifyTokenRequest, runtime *dara.RuntimeOptions) (_result *CrossAccountVerifyTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CrossAccountVerifyToken"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CrossAccountVerifyTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the API quota, total capacity, and usage data of a user.
//
// Description:
//
// ### **Note: This API operation supports Alibaba Cloud accounts and RAM users in retrieving API quota, total capacity, and usage information.**
//
// @param request - DescribeApiMeteringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiMeteringResponse
func (client *Client) DescribeApiMeteringWithContext(ctx context.Context, request *DescribeApiMeteringRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiMeteringResponse, _err error) {
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
		Action:      dara.String("DescribeApiMetering"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiMeteringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeCurrentNodeInfo to retrieve the process information of the current node.
//
// Description:
//
// *
//
// **.
//
// @param request - DescribeCurrentNodeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCurrentNodeInfoResponse
func (client *Client) DescribeCurrentNodeInfoWithContext(ctx context.Context, request *DescribeCurrentNodeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCurrentNodeInfoResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCurrentNodeInfo"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCurrentNodeInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of promotional products with detailed information.
//
// Description:
//
// Call this operation by using a promotion service provider account.
//
// @param request - DescribeDistributionProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistributionProductsResponse
func (client *Client) DescribeDistributionProductsWithContext(ctx context.Context, request *DescribeDistributionProductsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDistributionProductsResponse, _err error) {
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
		Action:      dara.String("DescribeDistributionProducts"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDistributionProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves promotion links for promoted products.
//
// Description:
//
// Call this operation by using a promotion service provider account. Calls from non-promotion service provider accounts will fail.
//
// @param tmpReq - DescribeDistributionProductsLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistributionProductsLinkResponse
func (client *Client) DescribeDistributionProductsLinkWithContext(ctx context.Context, tmpReq *DescribeDistributionProductsLinkRequest, runtime *dara.RuntimeOptions) (_result *DescribeDistributionProductsLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeDistributionProductsLinkShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Codes) {
		request.CodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Codes, dara.String("Codes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CodesShrink) {
		query["Codes"] = request.CodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDistributionProductsLink"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDistributionProductsLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of failed subscription notifications. After Alibaba Cloud Marketplace fails to call a service provider\\"s message SPI, the service provider can use this operation to proactively query the failed notification information.
//
// Description:
//
// Queries the information about failed subscription notifications for a service provider.
//
// - If no unconfirmed failed notifications exist, TotalCount returns 0.
//
// @param request - DescribeFailedNotificationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFailedNotificationsResponse
func (client *Client) DescribeFailedNotificationsWithContext(ctx context.Context, request *DescribeFailedNotificationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFailedNotificationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFailedNotifications"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFailedNotificationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务商侧查询镜像实例信息
//
// @param request - DescribeImageInstanceForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageInstanceForIsvResponse
func (client *Client) DescribeImageInstanceForIsvWithContext(ctx context.Context, request *DescribeImageInstanceForIsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageInstanceForIsvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerPk) {
		query["CustomerPk"] = request.CustomerPk
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImageInstanceForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageInstanceForIsvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeInstance operation to query instance information.
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
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

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instance information from the service provider side.
//
// @param request - DescribeInstanceForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceForIsvResponse
func (client *Client) DescribeInstanceForIsvWithContext(ctx context.Context, request *DescribeInstanceForIsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceForIsvResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceForIsvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of services purchased by the user.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Codes) {
		query["Codes"] = request.Codes
	}

	if !dara.IsNil(request.ExceptCodes) {
		query["ExceptCodes"] = request.ExceptCodes
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the invoice information of a Cloud Marketplace vendor store.
//
// Description:
//
// Queries business information based on an authorization code and first verifies whether the authorization code is valid.
//
// - If the authorization code is invalid, an error message is returned.
//
// - If the authorization code is valid, the authorization information is returned.
//
// @param request - DescribeInvoiceForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvoiceForIsvResponse
func (client *Client) DescribeInvoiceForIsvWithContext(ctx context.Context, request *DescribeInvoiceForIsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvoiceForIsvResponse, _err error) {
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

	if !dara.IsNil(request.InvoiceId) {
		query["InvoiceId"] = request.InvoiceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvoiceForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvoiceForIsvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeLicense to query license information.
//
// @param request - DescribeLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLicenseResponse
func (client *Client) DescribeLicenseWithContext(ctx context.Context, request *DescribeLicenseRequest, runtime *dara.RuntimeOptions) (_result *DescribeLicenseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LicenseCode) {
		query["LicenseCode"] = request.LicenseCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLicense"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLicenseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified order.
//
// @param request - DescribeOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOrderResponse
func (client *Client) DescribeOrderWithContext(ctx context.Context, request *DescribeOrderRequest, runtime *dara.RuntimeOptions) (_result *DescribeOrderResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOrder"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries order details from the service provider side.
//
// @param request - DescribeOrderForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOrderForIsvResponse
func (client *Client) DescribeOrderForIsvWithContext(ctx context.Context, request *DescribeOrderForIsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeOrderForIsvResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOrderForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOrderForIsvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribePrice to query the price information of the corresponding product.
//
// @param request - DescribePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceResponse
func (client *Client) DescribePriceWithContext(ctx context.Context, request *DescribePriceRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Commodity) {
		query["Commodity"] = request.Commodity
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrice"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified product.
//
// @param request - DescribeProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductResponse
func (client *Client) DescribeProductWithContext(ctx context.Context, request *DescribeProductRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.QueryDraft) {
		query["QueryDraft"] = request.QueryDraft
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProduct"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of products that meet the specified conditions along with their details.
//
// @param request - DescribeProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductsResponse
func (client *Client) DescribeProductsWithContext(ctx context.Context, request *DescribeProductsRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchTerm) {
		query["SearchTerm"] = request.SearchTerm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProducts"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of project attachments.
//
// @param request - DescribeProjectAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectAttachmentsResponse
func (client *Client) DescribeProjectAttachmentsWithContext(ctx context.Context, request *DescribeProjectAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectAttachmentsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectAttachments"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeProjectInfo to query project information.
//
// @param request - DescribeProjectInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectInfoResponse
func (client *Client) DescribeProjectInfoWithContext(ctx context.Context, request *DescribeProjectInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectInfoResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectInfo"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of messages for a project workflow.
//
// @param request - DescribeProjectMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectMessagesResponse
func (client *Client) DescribeProjectMessagesWithContext(ctx context.Context, request *DescribeProjectMessagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectMessagesResponse, _err error) {
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

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectMessages"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectMessagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of project flow nodes by calling DescribeProjectNodes.
//
// Description:
//
// *
//
// **.
//
// @param request - DescribeProjectNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectNodesResponse
func (client *Client) DescribeProjectNodesWithContext(ctx context.Context, request *DescribeProjectNodesRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectNodesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectNodes"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls DescribeProjectOperateLogs to retrieve the operation logs of a project flow.
//
// Description:
//
// *
//
// **.
//
// @param request - DescribeProjectOperateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectOperateLogsResponse
func (client *Client) DescribeProjectOperateLogsWithContext(ctx context.Context, request *DescribeProjectOperateLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectOperateLogsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjectOperateLogs"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectOperateLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Completes the current process by calling FinishCurrentProjectNode.
//
// Description:
//
// *
//
// **.
//
// @param request - FinishCurrentProjectNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishCurrentProjectNodeResponse
func (client *Client) FinishCurrentProjectNodeWithContext(ctx context.Context, request *FinishCurrentProjectNodeRequest, runtime *dara.RuntimeOptions) (_result *FinishCurrentProjectNodeResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.TemplateForm) {
		query["TemplateForm"] = request.TemplateForm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FinishCurrentProjectNode"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FinishCurrentProjectNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Processes invoice acceptance requests.
//
// Description:
//
// Queries information about failed subscription notifications for merchants.
//
// - If no unconfirmed notification failure information exists, TotalCount returns 0.
//
// @param request - ModifyInvoiceForIsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInvoiceForIsvResponse
func (client *Client) ModifyInvoiceForIsvWithContext(ctx context.Context, request *ModifyInvoiceForIsvRequest, runtime *dara.RuntimeOptions) (_result *ModifyInvoiceForIsvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckNotice) {
		query["CheckNotice"] = request.CheckNotice
	}

	if !dara.IsNil(request.ElectronUrl) {
		query["ElectronUrl"] = request.ElectronUrl
	}

	if !dara.IsNil(request.InvoiceId) {
		query["InvoiceId"] = request.InvoiceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInvoiceForIsv"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInvoiceForIsvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls PauseProject to pause a project.
//
// Description:
//
// *
//
// **.
//
// @param request - PauseProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseProjectResponse
func (client *Client) PauseProjectWithContext(ctx context.Context, request *PauseProjectRequest, runtime *dara.RuntimeOptions) (_result *PauseProjectResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseProject"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes metering data.
//
// Description:
//
// ### Encoding and format requirements
//
// - All request parameters must be URL-encoded.
//
// - Ensure that the Metering field in the parameters is submitted strictly according to the examples in the documentation. In the past, many submissions were not valid JSON fields. You can use a JSON formatting tool to check the field before submission. Otherwise, the API call will fail.
//
// ### Request frequency and single write volume limits
//
// - For hourly or daily metering: Alibaba Cloud Marketplace limits requests to once every 60 seconds per InstanceId, and each request supports writing up to 100 metering data records.
//
// - For real-time metering (for large model scenarios): Alibaba Cloud Marketplace has no frequency limit for requests per InstanceId, and each request supports writing up to 10 metering data records.
//
// ### Limits on batch pushing of metering data
//
// - Batch pushing across InstanceIds of different products is not supported. When pushing in batches, ensure that all InstanceIds belong to the same product.
//
// - When pushing metering data in batches, the InstanceId attribute must belong to instances generated by the same product. Pushing instances of multiple products at the same time is not supported.
//
// ### Time requirements for pushing metering data
//
// - For products billed hourly, the billing cutoff time is the next hour, and metering data must be pushed on time. For example, if the metering data is for 08:10 to 08:20, it must be pushed before 09:59. Otherwise, no bill will be generated. The rule for daily billing is similar: the data for the previous day must be pushed before 01:59 the next day. Otherwise, an error will be reported when pushing the metering data, and no bill will be generated.
//
// - For products billed by the minute, real-time pushed data is aggregated by the minute, and bills are generated by the minute.
//
// - If a product is set to real-time billing, StartTime and EndTime can be of any time width, but EndTime must be greater than StartTime.
//
// - If a product is set to non-real-time billing (such as hourly or daily), the interval between StartTime and EndTime must be greater than 5 minutes.
//
// ### Timeliness of pushing data and billing
//
// - If the instance status is expired, metering data pushed after the expiration will not be billed. For example, if an instance expires at 2025-07-09 12:20:00, the metering data pushed for the hour from 2025-07-09 13:00:00 to 2025-07-09 14:00:00 will not be billed. However, the metering data for the hour from 2025-07-09 12:00:00 to 2025-07-09 13:00:00 can still be billed normally.
//
// - After metering data is pushed, bills are not generated immediately. For hourly billing, the bill is generated in the next time interval. For example, the bill for 08:00 to 09:00 is generated after 10:00. For daily billing, the bill is generated the next day. After the bill is generated, customers can view it in [Alibaba Cloud Bill Details](https://billing-cost.console.aliyun.com/finance/expense-report/expense-detail-by-instance).
//
// ### Rounding down of amounts
//
// - If the bill amount is too small and a single bill amount is less than 0.01 CNY, the bill amount for the buyer is 0. For example, if the unit price of a product is 0.001 CNY per time and it is billed by the minute, and 9 times are consumed every minute, the bill amount per minute is 0. If it is billed hourly and the bill amount for a single hour is less than 0.01 CNY, the bill amount for the buyer is 0. The probability of rounding down to zero for minute-level billing may be higher than that for hourly billing.
//
// @param request - PushMeteringDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushMeteringDataResponse
func (client *Client) PushMeteringDataWithContext(ctx context.Context, request *PushMeteringDataRequest, runtime *dara.RuntimeOptions) (_result *PushMeteringDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Metering) {
		query["Metering"] = request.Metering
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushMeteringData"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushMeteringDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes usage deductions for prepaid pay-per-use products to Alibaba Cloud Marketplace.
//
// Description:
//
// ## Request description
//
// - This operation applies to prepaid pay-per-use products. Service providers can use this operation to push user usage deductions to Alibaba Cloud Marketplace.
//
// - This operation uses a forced update mode to directly deduct the specified number of uses. The caller must determine whether the user\\"s usage exceeds the limit.
//
// - All request parameters must be URL-encoded.
//
// - Use the `ClientToken` parameter to ensure idempotence. The same `ClientToken` cannot be submitted repeatedly within 24 hours.
//
// - Retries for the same deduction operation must use the same `ClientToken`. New operations must use a new `ClientToken`.
//
// - If the previous request succeeded, even if no response was received due to a gateway timeout, a retry with the same `ClientToken` returns a success response idempotently without duplicate deductions.
//
// - If the result of the previous request is undetermined and the request is in the PROCESSING state, a retry with the same `ClientToken` returns a `DuplicateClientToken` error. Retry later or contact Alibaba Cloud Marketplace to confirm the push status.
//
// - Only one operation can be executed for the same `InstanceId` at a time. Concurrent requests return a `Throttling` error. Retry later.
//
// - You can directly call this operation in [OpenAPI Explorer](https://api.aliyun.com/#product=Market&api=PushTimesUsage&type=RPC&version=2015-11-01) without the need to manually calculate the signature.
//
// @param request - PushTimesUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushTimesUsageResponse
func (client *Client) PushTimesUsageWithContext(ctx context.Context, request *PushTimesUsageRequest, runtime *dara.RuntimeOptions) (_result *PushTimesUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Adjust) {
		query["Adjust"] = request.Adjust
	}

	if !dara.IsNil(request.AdjustDate) {
		query["AdjustDate"] = request.AdjustDate
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mark) {
		query["Mark"] = request.Mark
	}

	if !dara.IsNil(request.Times) {
		query["Times"] = request.Times
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushTimesUsage"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushTimesUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls ResumeProject to resume a paused project.
//
// Description:
//
// *
//
// **.
//
// @param request - ResumeProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeProjectResponse
func (client *Client) ResumeProjectWithContext(ctx context.Context, request *ResumeProjectRequest, runtime *dara.RuntimeOptions) (_result *ResumeProjectResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeProject"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a rejection of the current process node by calling RollbackCurrentProjectNode.
//
// Description:
//
// *
//
// **.
//
// @param request - RollbackCurrentProjectNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackCurrentProjectNodeResponse
func (client *Client) RollbackCurrentProjectNodeWithContext(ctx context.Context, request *RollbackCurrentProjectNodeRequest, runtime *dara.RuntimeOptions) (_result *RollbackCurrentProjectNodeResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackCurrentProjectNode"),
		Version:     dara.String("2015-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackCurrentProjectNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
