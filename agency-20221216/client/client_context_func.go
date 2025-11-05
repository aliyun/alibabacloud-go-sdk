// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 作废优惠券
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
// 优惠券审批状态列表
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
// 创建优惠券模板
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
// This function is designed for create a customer who is to be invited.
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
// Query quota adjustment list of Distribution Customer from International Site. Not available on Domestic Site.
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
// 作废优惠券模板
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
// 额度冲减明细列表导出接口
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
// Return Distribution Customer\\"s account information.
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
// 提供返佣商品API
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
// 查询优惠券模板详情
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
// 国际渠道分销优惠券可抵扣产品
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
// Query Credit Control information of Distribution Customers. The PopCreditInfoJson in the Return Parameter will be empty if the Distribution Customer is an Agency. This function is only available for Resellers and Distributors.
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
// 客户订单查询
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
// Issue Distributor\\"s daily Bill. This function is only available for Resellers and Distributors.
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
// Query invitation status of customer who have been created and invited.
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
// Issue Distributor\\"s Monthly Bill. This function is only available for Resellers and Distributors.
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
// Query all the Unassociated Customer.
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
// 发放优惠券
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
// 优惠券使用量列表查询
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
// 通用查询导出任务列表
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
// 额度冲减明细列表
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
// Resend invitation email.
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
// This function is designed for Sub Account information maintenance, including Nickname and Remark.
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
// Set Credit Line for Distribution Customers. This function is only available for Resellers and Distributors.
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
// You can use this API to set the threshold for the use of credit control. When the customer credit control reaches below the threshold, it will pass through the notification email distributor. This feature is for Reseller and Distributor only.
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
