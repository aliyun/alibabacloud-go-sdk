// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Configures the automatic write-off feature before the end-of-month settlement. When the consumed but undeducted quota reaches the specified automatic write-off amount, the amount is transferred to the deducted quota for offset.
//
// @param request - AutomaticWriteOffRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AutomaticWriteOffResponse
func (client *Client) AutomaticWriteOffWithContext(ctx context.Context, request *AutomaticWriteOffRequest, runtime *dara.RuntimeOptions) (_result *AutomaticWriteOffResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutomaticWriteOffAmount) {
		query["AutomaticWriteOffAmount"] = request.AutomaticWriteOffAmount
	}

	if !dara.IsNil(request.AutomaticWriteOffEnabled) {
		query["AutomaticWriteOffEnabled"] = request.AutomaticWriteOffEnabled
	}

	if !dara.IsNil(request.CustomerUid) {
		query["CustomerUid"] = request.CustomerUid
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AutomaticWriteOff"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AutomaticWriteOffResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invalidates a coupon.
//
// @param request - CancelCouponRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCouponResponse
func (client *Client) CancelCouponWithContext(ctx context.Context, request *CancelCouponRequest, runtime *dara.RuntimeOptions) (_result *CancelCouponResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unsubscribe from multi-level bills for international channel ecosystem partners.
//
// Description:
//
// Make sure that the caller is an international Alibaba Cloud master reseller partner.
//
// You can unsubscribe from one type of bill file at a time.
//
// After the corresponding bill type is unsubscribed, files are no longer pushed to the specified OSS space.
//
// **This content is published only on the international site.**
//
// @param request - CancelSubscriptionBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSubscriptionBillResponse
func (client *Client) CancelSubscriptionBillWithContext(ctx context.Context, request *CancelSubscriptionBillRequest, runtime *dara.RuntimeOptions) (_result *CancelSubscriptionBillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of coupon approval statuses.
//
// Description:
//
// Ensure that the current calling account is a distribution partner.
//
// <notice>Available only on the international site.</notice>
//
// @param request - CouponApprovalStatusListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CouponApprovalStatusListResponse
func (client *Client) CouponApprovalStatusListWithContext(ctx context.Context, request *CouponApprovalStatusListRequest, runtime *dara.RuntimeOptions) (_result *CouponApprovalStatusListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a coupon template.
//
// @param tmpReq - CreateCouponTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCouponTemplateResponse
func (client *Client) CreateCouponTemplateWithContext(ctx context.Context, tmpReq *CreateCouponTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateCouponTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a customer to be invited.
//
// @param request - CreateCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerResponse
func (client *Client) CreateCustomerWithContext(ctx context.Context, request *CreateCustomerRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota adjustment records for distribution customers on the Alibaba Cloud International Website. This operation applies only to international users and cannot be called on the Alibaba Cloud China Website.
//
// Description:
//
// The caller must be a channel master distributor partner or a channel distributor partner on the Alibaba Cloud International Website.
//
// @param request - CustomerQuotaRecordListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerQuotaRecordListResponse
func (client *Client) CustomerQuotaRecordListWithContext(ctx context.Context, request *CustomerQuotaRecordListRequest, runtime *dara.RuntimeOptions) (_result *CustomerQuotaRecordListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reduces the deducted credit of a reseller\\"s customer. For example, if the current deducted credit is 500 and the available credit is 1000, a reduction of 300 changes the deducted credit to 200 and the available credit to 1300.
//
// Note: Sometimes the used credit of a customer may be negative. This indicates that no recovery of the used credit is needed. Wait for the customer to consume the credit. This occurs when a refund is issued while the customer\\"s credit is fully used. As a result, the customer\\"s credit increases beyond the original limit.
//
// For example, if the current maximum available credit of a customer is 1000 with no usage, and a refund of 300 occurs, the used credit becomes -300.
//
// @param request - DeductOutstandingBalanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeductOutstandingBalanceResponse
func (client *Client) DeductOutstandingBalanceWithContext(ctx context.Context, request *DeductOutstandingBalanceRequest, runtime *dara.RuntimeOptions) (_result *DeductOutstandingBalanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invalidates a coupon template.
//
// Description:
//
// Make sure that the current caller is a distribution partner.
//
// <notice>Available only on the China site.</notice>.
//
// @param request - DeleteCouponTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCouponTemplateResponse
func (client *Client) DeleteCouponTemplateWithContext(ctx context.Context, request *DeleteCouponTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCouponTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manages the shutdown status of pay-as-you-go instances for resold subordinate customers by Alibaba Cloud international channel master distributor partners and Alibaba Cloud international channel distributor partners.
//
// Description:
//
// The caller must be an Alibaba Cloud international channel master distributor partner or an Alibaba Cloud international channel distributor partner.
//
// **This content is published only on the China site (Chinese).**.
//
// @param request - EditEndUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditEndUserStatusResponse
func (client *Client) EditEndUserStatusWithContext(ctx context.Context, request *EditEndUserStatusRequest, runtime *dara.RuntimeOptions) (_result *EditEndUserStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the new purchase status of a subordinate reseller customer for an Alibaba Cloud reseller partner.
//
// Description:
//
// The caller must be an Alibaba Cloud reseller partner.
//
// **This content is published only on the China site.**.
//
// @param request - EditNewBuyStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditNewBuyStatusResponse
func (client *Client) EditNewBuyStatusWithContext(ctx context.Context, request *EditNewBuyStatusRequest, runtime *dara.RuntimeOptions) (_result *EditNewBuyStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the shutdown policy of a subordinate customer by an Alibaba Cloud distributor partner or reseller partner.
//
// Description:
//
// The caller must be an Alibaba Cloud international channel distributor partner or an Alibaba Cloud international channel reseller partner.
//
// **This content is published only on the Chinese site.**.
//
// @param request - EditZeroCreditShutdownRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditZeroCreditShutdownResponse
func (client *Client) EditZeroCreditShutdownWithContext(ctx context.Context, request *EditZeroCreditShutdownRequest, runtime *dara.RuntimeOptions) (_result *EditZeroCreditShutdownResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports the quota adjustment records of a distributor\\"s customer, corresponding to the total budget adjustment details export feature on ACPN.
//
// Description:
//
// The caller must be an Alibaba Cloud international site channel master distributor partner or an Alibaba Cloud international site channel distributor partner.
//
// @param request - ExportCustomerQuotaRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportCustomerQuotaRecordResponse
func (client *Client) ExportCustomerQuotaRecordWithContext(ctx context.Context, request *ExportCustomerQuotaRecordRequest, runtime *dara.RuntimeOptions) (_result *ExportCustomerQuotaRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports the reversed deduction details from ACPN.
//
// @param request - ExportReversedDeductionHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportReversedDeductionHistoryResponse
func (client *Client) ExportReversedDeductionHistoryWithContext(ctx context.Context, request *ExportReversedDeductionHistoryRequest, runtime *dara.RuntimeOptions) (_result *ExportReversedDeductionHistoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the account information of distribution customers.
//
// @param request - GetAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfoWithContext(ctx context.Context, request *GetAccountInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAccountInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of all active and delisted commission-eligible commodities for T1-level international distribution partners. This operation is available only for international users and cannot be called on the China site.
//
// Description:
//
// Ensure that the current calling account is a distribution partner.
//
// <notice>Available only on the international site.</notice>
//
// @param tmpReq - GetCommissionableProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommissionableProductsResponse
func (client *Client) GetCommissionableProductsWithContext(ctx context.Context, tmpReq *GetCommissionableProductsRequest, runtime *dara.RuntimeOptions) (_result *GetCommissionableProductsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a coupon template.
//
// Description:
//
// Ensure that the current calling account is a distribution partner.
//
// <notice>Available only on the International site
//
// @param request - GetCouponTemplateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCouponTemplateDetailResponse
func (client *Client) GetCouponTemplateDetailWithContext(ctx context.Context, request *GetCouponTemplateDetailRequest, runtime *dara.RuntimeOptions) (_result *GetCouponTemplateDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the products that can be offset by international channel distribution coupons.
//
// @param request - GetCoupondeductProductCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCoupondeductProductCodeResponse
func (client *Client) GetCoupondeductProductCodeWithContext(ctx context.Context, request *GetCoupondeductProductCodeRequest, runtime *dara.RuntimeOptions) (_result *GetCoupondeductProductCodeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the credit control information of a distributor customer. If the distributor customer is of the Agency type, the PopCreditInfoJson parameter in the response is empty. This feature is available only for Reseller and Distributor.
//
// @param request - GetCreditInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreditInfoResponse
func (client *Client) GetCreditInfoWithContext(ctx context.Context, request *GetCreditInfoRequest, runtime *dara.RuntimeOptions) (_result *GetCreditInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries customer orders on the Alibaba Cloud International site. This operation is only allowed for the international site.
//
// Description:
//
// The caller must be an Alibaba Cloud International site channel master distributor partner or an Alibaba Cloud International site channel distributor partner.
//
// @param request - GetCustomerOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerOrdersResponse
func (client *Client) GetCustomerOrdersWithContext(ctx context.Context, request *GetCustomerOrdersRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerOrdersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the daily bills for resellers and distributors. This feature is only available for resellers and distributors.
//
// @param request - GetDailyBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDailyBillResponse
func (client *Client) GetDailyBillWithContext(ctx context.Context, request *GetDailyBillRequest, runtime *dara.RuntimeOptions) (_result *GetDailyBillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the invitation status of registered customers who have been sent invitations.
//
// @param request - GetInviteStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInviteStatusResponse
func (client *Client) GetInviteStatusWithContext(ctx context.Context, request *GetInviteStatusRequest, runtime *dara.RuntimeOptions) (_result *GetInviteStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the monthly bill for distributors. This feature is available only to resellers and distributors.
//
// @param request - GetMonthlyBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMonthlyBillResponse
func (client *Client) GetMonthlyBillWithContext(ctx context.Context, request *GetMonthlyBillRequest, runtime *dara.RuntimeOptions) (_result *GetMonthlyBillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries order control change records.
//
// Description:
//
// The caller must be an Alibaba Cloud international site channel master distributor partner or an Alibaba Cloud international site channel distributor partner.
//
// @param request - GetPurchaseControlRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurchaseControlRecordResponse
func (client *Client) GetPurchaseControlRecordWithContext(ctx context.Context, request *GetPurchaseControlRecordRequest, runtime *dara.RuntimeOptions) (_result *GetPurchaseControlRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the modification records of the deferred shutdown policy.
//
// Description:
//
// The caller must be an Alibaba Cloud international site channel master distributor partner or an Alibaba Cloud international site channel distributor partner.
//
// **This content is published only on the international site.**.
//
// @param request - GetShutdownPolicyRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShutdownPolicyRecordResponse
func (client *Client) GetShutdownPolicyRecordWithContext(ctx context.Context, request *GetShutdownPolicyRecordRequest, runtime *dara.RuntimeOptions) (_result *GetShutdownPolicyRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a coupon application initiated by a tier-2 distributor. This operation applies only to distributors.
//
// @param request - GetTier2CouponApprovalDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTier2CouponApprovalDetailResponse
func (client *Client) GetTier2CouponApprovalDetailWithContext(ctx context.Context, request *GetTier2CouponApprovalDetailRequest, runtime *dara.RuntimeOptions) (_result *GetTier2CouponApprovalDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationSheetId) {
		query["ApplicationSheetId"] = request.ApplicationSheetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTier2CouponApprovalDetail"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTier2CouponApprovalDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all unassociated customers.
//
// @param request - GetUnassociatedCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUnassociatedCustomerResponse
func (client *Client) GetUnassociatedCustomerWithContext(ctx context.Context, request *GetUnassociatedCustomerRequest, runtime *dara.RuntimeOptions) (_result *GetUnassociatedCustomerResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invites a customer to register as a partner customer.
//
// @param request - InviteSubAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteSubAccountResponse
func (client *Client) InviteSubAccountWithContext(ctx context.Context, request *InviteSubAccountRequest, runtime *dara.RuntimeOptions) (_result *InviteSubAccountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invites T2 Reseller sub-accounts in batch as a T1 partner.
//
// Description:
//
// 1. The caller of this operation must be a channel master reseller partner on Alibaba Cloud International Website.
//
// 2. The system automatically determines whether the invitation is cross-region based on whether the RegisterNation parameter falls within the T1 contract coverage area. You can call the ListCountries operation to query the contract coverage area.
//
// - If the invitation is cross-region, a cross-region invitation approval process is initiated. After Alibaba Cloud approves the request, an invitation registration email is sent to the specified email address.
//
// - If the invitation is not cross-region, an invitation registration email is sent directly.
//
// @param request - InviteSubResellerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InviteSubResellerResponse
func (client *Client) InviteSubResellerWithContext(ctx context.Context, request *InviteSubResellerRequest, runtime *dara.RuntimeOptions) (_result *InviteSubResellerResponse, _err error) {
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
		Action:      dara.String("InviteSubReseller"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InviteSubResellerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Issues coupons.
//
// @param request - IssueCouponForCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IssueCouponForCustomerResponse
func (client *Client) IssueCouponForCustomerWithContext(ctx context.Context, request *IssueCouponForCustomerRequest, runtime *dara.RuntimeOptions) (_result *IssueCouponForCustomerResponse, _err error) {
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

	if !dara.IsNil(request.ApplicationReason) {
		query["ApplicationReason"] = request.ApplicationReason
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage of coupons issued by a distribution partner.
//
// Description:
//
// Ensure that the current calling account is a distribution partner.
//
// <notice>Available only on the international site.</notice>
//
// @param request - ListCouponUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCouponUsageResponse
func (client *Client) ListCouponUsageWithContext(ctx context.Context, request *ListCouponUsageRequest, runtime *dara.RuntimeOptions) (_result *ListCouponUsageResponse, _err error) {
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

	if !dara.IsNil(request.T2PartnerUid) {
		query["T2PartnerUid"] = request.T2PartnerUid
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of exported files.
//
// @param request - ListExportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExportTasksResponse
func (client *Client) ListExportTasksWithContext(ctx context.Context, request *ListExportTasksRequest, runtime *dara.RuntimeOptions) (_result *ListExportTasksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the coupon applications submitted by secondary distributors for their end users. This operation applies only to tier-1 distributors.
//
// @param request - ListTier2CouponApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTier2CouponApprovalResponse
func (client *Client) ListTier2CouponApprovalWithContext(ctx context.Context, request *ListTier2CouponApprovalRequest, runtime *dara.RuntimeOptions) (_result *ListTier2CouponApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationSheetId) {
		query["ApplicationSheetId"] = request.ApplicationSheetId
	}

	if !dara.IsNil(request.ApprovalStatus) {
		query["ApprovalStatus"] = request.ApprovalStatus
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.T2PartnerName) {
		query["T2PartnerName"] = request.T2PartnerName
	}

	if !dara.IsNil(request.T2PartnerUid) {
		query["T2PartnerUid"] = request.T2PartnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTier2CouponApproval"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTier2CouponApprovalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Approves a process that requires partner approval. This operation is applicable only to distributors.
//
// @param request - ProcessApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProcessApprovalResponse
func (client *Client) ProcessApprovalWithContext(ctx context.Context, request *ProcessApprovalRequest, runtime *dara.RuntimeOptions) (_result *ProcessApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationSheetId) {
		query["ApplicationSheetId"] = request.ApplicationSheetId
	}

	if !dara.IsNil(request.ApprovalAction) {
		query["ApprovalAction"] = request.ApprovalAction
	}

	if !dara.IsNil(request.ApprovalComments) {
		query["ApprovalComments"] = request.ApprovalComments
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProcessApproval"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProcessApprovalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the change records of the automatic write-off quota by using paging.
//
// @param request - QueryAutomaticWriteOffChangeRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAutomaticWriteOffChangeRecordsResponse
func (client *Client) QueryAutomaticWriteOffChangeRecordsWithContext(ctx context.Context, request *QueryAutomaticWriteOffChangeRecordsRequest, runtime *dara.RuntimeOptions) (_result *QueryAutomaticWriteOffChangeRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerUid) {
		query["CustomerUid"] = request.CustomerUid
	}

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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAutomaticWriteOffChangeRecords"),
		Version:     dara.String("2022-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAutomaticWriteOffChangeRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota offset details of a specified customer.
//
// @param request - QueryReversedDeductionHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReversedDeductionHistoryResponse
func (client *Client) QueryReversedDeductionHistoryWithContext(ctx context.Context, request *QueryReversedDeductionHistoryRequest, runtime *dara.RuntimeOptions) (_result *QueryReversedDeductionHistoryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota transaction export results for international distribution users. This API operation is available only on the international site.
//
// Description:
//
// The caller must be an Alibaba Cloud International site channel master distributor partner or an Alibaba Cloud International site channel distributor partner.
//
// @param request - QuotaListExportPagedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuotaListExportPagedResponse
func (client *Client) QuotaListExportPagedWithContext(ctx context.Context, request *QuotaListExportPagedRequest, runtime *dara.RuntimeOptions) (_result *QuotaListExportPagedResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resends an invitation email.
//
// @param request - ResendEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendEmailResponse
func (client *Client) ResendEmailWithContext(ctx context.Context, request *ResendEmailRequest, runtime *dara.RuntimeOptions) (_result *ResendEmailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates sub-account information, including the nickname and remarks.
//
// @param request - SetAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAccountInfoResponse
func (client *Client) SetAccountInfoWithContext(ctx context.Context, request *SetAccountInfoRequest, runtime *dara.RuntimeOptions) (_result *SetAccountInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the credit limit for a reseller customer. This feature is available only for Resellers and Distributors.
//
// @param request - SetCreditLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCreditLineResponse
func (client *Client) SetCreditLineWithContext(ctx context.Context, request *SetCreditLineRequest, runtime *dara.RuntimeOptions) (_result *SetCreditLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the credit usage threshold. When a customer\\"s available credit falls below the threshold, the reseller is notified by email. This feature is available only for Resellers and Distributors.
//
// @param request - SetWarningThresholdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWarningThresholdResponse
func (client *Client) SetWarningThresholdWithContext(ctx context.Context, request *SetWarningThresholdRequest, runtime *dara.RuntimeOptions) (_result *SetWarningThresholdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Subscribes to multi-level ecosystem partner bills on Alibaba Cloud International and generates bill files to OSS.
//
// Description:
//
// - Ensure that the caller is an Alibaba Cloud international master reseller partner.
//
// - A user can subscribe to one type of bill file at a time.
//
// - After subscription, bill files generated on the previous day are pushed daily starting from the next day. On the 5th of each month, full bill files for the previous billing cycle are pushed.
//
// - Bill files generated on a daily basis may be delayed. Delayed bill files are pushed the day after they are generated and may include bills that were delayed from earlier days. Pull the full files for the previous month at the beginning of each month.
//
// - Grant the [AliyunConsumeDump2OSSRole](https://ram.console.aliyun.com/?spm=api-workbench.API%20Document.0.0.68c71e0fhmTSJp#/role/authorize?request=%7B%22Requests%22:%20%7B%22request1%22:%20%7B%22RoleName%22:%20%22AliyunConsumeDump2OSSRole%22,%20%22TemplateId%22:%20%22Dump2OSSRole%22%7D%7D,%20%22ReturnUrl%22:%20%22https:%2F%2Fusercenter2.aliyun.com%22,%20%22Service%22:%20%22Consume%22%7D) permission to the subscribing user.
//
// - File names.
//
// ```.
//
// BillingItemDetailForBillingPeriod   (Billing item detail)
//
// Daily push file name format: UID_PartnerBillingItemDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_BillingItemDetail_20190310_0001_01.
//
// Full file name format at the beginning of the next month: UID_PartnerBillingItemDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetail_201903_0001_01.
//
// InstanceDetailForBillingPeriod  (Instance bill detail)
//
//	Daily push file name format: UID_PartnerInstanceDetail_YYYYMMDD_SquenceNo_fileNo. Example: 169**_InstanceDetail_20190310_0001_01.
//
// Full file name format at the beginning of the next month: UID_PartnerInstanceDetail_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetail_201903_1999-0001_01.
//
// BillingItemDetailMonthly   (Billing item monthly bill)
//
// Daily push file name format: UID_PartnerBillingItemDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_BillingItemDetailMonthly_201903_0001_01. This file is updated daily with full data from the beginning of the month to the current day, until the 5th of the next month.
//
// InstanceDetailMonthly  (Instance monthly bill)
//
// Daily push file name format: UID_PartnerInstanceDetailMonthly_YYYYMM_SquenceNo_fileNo. Example: 169**_InstanceDetailMonthly_201903_0001_01. This file is updated daily with full data from the beginning of the month to the current day, until the 5th of the next month.
//
// fileNo exists only when the number of bill rows reaches the split threshold and the file is split.
//
// ```
//
// **This content is published only on the Chinese site.**.
//
// @param request - SubscriptionBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscriptionBillResponse
func (client *Client) SubscriptionBillWithContext(ctx context.Context, request *SubscriptionBillRequest, runtime *dara.RuntimeOptions) (_result *SubscriptionBillResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
