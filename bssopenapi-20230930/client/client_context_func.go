// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 添加优惠券抵扣标签
//
// @param tmpReq - AddCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCouponDeductTagResponse
func (client *Client) AddCouponDeductTagWithContext(ctx context.Context, tmpReq *AddCouponDeductTagRequest, runtime *dara.RuntimeOptions) (_result *AddCouponDeductTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Allocates resource instances (instance-based and attached-resource-based) from a source cost center to a destination cost center.
//
// @param tmpReq - AllocateCostCenterResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateCostCenterResourceResponse
func (client *Client) AllocateCostCenterResourceWithContext(ctx context.Context, tmpReq *AllocateCostCenterResourceRequest, runtime *dara.RuntimeOptions) (_result *AllocateCostCenterResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the low balance alert for a fund account.
//
// Description:
//
// Cancels the low balance alert for a fund account.
//
// @param request - CancelFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelFundAccountLowAvailableAmountAlarmResponse
func (client *Client) CancelFundAccountLowAvailableAmountAlarmWithContext(ctx context.Context, request *CancelFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *CancelFundAccountLowAvailableAmountAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提货券账户检查是否存在
//
// @param request - CheckAccountExistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountExistResponse
func (client *Client) CheckAccountExistWithContext(ctx context.Context, request *CheckAccountExistRequest, runtime *dara.RuntimeOptions) (_result *CheckAccountExistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIds) {
		query["EcIdAccountIds"] = request.EcIdAccountIds
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ToUserType) {
		body["ToUserType"] = request.ToUserType
	}

	if !dara.IsNil(request.TransferAccount) {
		body["TransferAccount"] = request.TransferAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAccountExist"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAccountExistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a specified budgetName exists.
//
// @param request - CheckBudgetNameExistsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckBudgetNameExistsResponse
func (client *Client) CheckBudgetNameExistsWithContext(ctx context.Context, request *CheckBudgetNameExistsRequest, runtime *dara.RuntimeOptions) (_result *CheckBudgetNameExistsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BudgetName) {
		body["BudgetName"] = request.BudgetName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckBudgetNameExists"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckBudgetNameExistsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a budget.
//
// @param tmpReq - CreateBudgetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBudgetResponse
func (client *Client) CreateBudgetWithContext(ctx context.Context, tmpReq *CreateBudgetRequest, runtime *dara.RuntimeOptions) (_result *CreateBudgetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBudgetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CycleQuota) {
		request.CycleQuotaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CycleQuota, dara.String("CycleQuota"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryFilter) {
		request.QueryFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryFilter, dara.String("QueryFilter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WarnConfs) {
		request.WarnConfsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WarnConfs, dara.String("WarnConfs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BudgetName) {
		body["BudgetName"] = request.BudgetName
	}

	if !dara.IsNil(request.BudgetType) {
		body["BudgetType"] = request.BudgetType
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.CycleEndPeriod) {
		body["CycleEndPeriod"] = request.CycleEndPeriod
	}

	if !dara.IsNil(request.CycleQuotaShrink) {
		body["CycleQuota"] = request.CycleQuotaShrink
	}

	if !dara.IsNil(request.CycleStartPeriod) {
		body["CycleStartPeriod"] = request.CycleStartPeriod
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.Metric) {
		body["Metric"] = request.Metric
	}

	if !dara.IsNil(request.QueryFilterShrink) {
		body["QueryFilter"] = request.QueryFilterShrink
	}

	if !dara.IsNil(request.Quota) {
		body["Quota"] = request.Quota
	}

	if !dara.IsNil(request.QuotaType) {
		body["QuotaType"] = request.QuotaType
	}

	if !dara.IsNil(request.WarnConfsShrink) {
		body["WarnConfs"] = request.WarnConfsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBudget"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBudgetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Cost Center
//
// Description:
//
// Creates one or more cost centers.
//
// @param tmpReq - CreateCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostCenterResponse
func (client *Client) CreateCostCenterWithContext(ctx context.Context, tmpReq *CreateCostCenterRequest, runtime *dara.RuntimeOptions) (_result *CreateCostCenterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a financial unit auto-allocation rule
//
// @param tmpReq - CreateCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostCenterRuleResponse
func (client *Client) CreateCostCenterRuleWithContext(ctx context.Context, tmpReq *CreateCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateCostCenterRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create payment relationships for a fund account
//
// @param tmpReq - CreateFundAccountPayRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFundAccountPayRelationResponse
func (client *Client) CreateFundAccountPayRelationWithContext(ctx context.Context, tmpReq *CreateFundAccountPayRelationRequest, runtime *dara.RuntimeOptions) (_result *CreateFundAccountPayRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account transfer or revocation.
//
// @param request - CreateFundAccountTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFundAccountTransferResponse
func (client *Client) CreateFundAccountTransferWithContext(ctx context.Context, request *CreateFundAccountTransferRequest, runtime *dara.RuntimeOptions) (_result *CreateFundAccountTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Apply for Invoice
//
// @param tmpReq - CreateInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInvoiceResponse
func (client *Client) CreateInvoiceWithContext(ctx context.Context, tmpReq *CreateInvoiceRequest, runtime *dara.RuntimeOptions) (_result *CreateInvoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a bill report subscription.
//
// Description:
//
// When calling this operation, note the following:
//
// - A user can subscribe to one type of bill file at a time.
//
// - Except for monthly bill PDFs, after subscription, starting from the next day, the system pushes a bill file that contains full detailed data from the beginning of the current month to date. Before the 4th of each month, the system pushes the full bill file for the entire previous billing cycle.
//
// - Monthly bill PDFs are pushed before the 4th of each month for the previous month.
//
// - Bill files generated on a daily basis may have latency. Delayed bills are pushed the day after they are generated and may include bills from before the previous day that were delayed until the previous day. Pull the full file for the previous month at the beginning of each month.
//
// > Apply for permissions as described in the documentation: [Bill subscription](https://www.alibabacloud.com/help/en/user-center/user-guide/billing-subscription)
//
// - This subscription is the same feature as Expenses and Costs - Bill Subscription. Subscriptions are shared between the two.
//
// - When subscribing to a directory under a bucket, ensure the directory name complies with the naming conventions:
//
//   - Emojis are not allowed. Use valid UTF-8 characters.
//
//   - / is used to separate paths and can quickly create subdirectories. Do not start with / or \\, and do not use consecutive / characters.
//
//   - Subdirectories named .. are not allowed.
//
//   - The total length must be 1 to 254 characters.
//
// - File names:
//
//   - Example: **consumeDetailBillV2*	- (billing item bill details)
//
//   - Daily push file name format: `{Account UID}_{Sales site ID}_{Bill type}_{YYYYMM|YYYYMMDD}`, for example: `169**_2688801000001_consumeDetailBillV2_20190312`.
//
//   - Full file name format at the beginning of the next month: `{Account UID}_{Sales site ID}_{Bill type}_{YYYYMM|YYYYMM}`, for example: `169**_2688801000001_consumeDetailBillV2_201903`.
//
// - Monthly bill PDF type files are in .pdf format. All other file types are .csv files. When the data volume is large, the system automatically splits the exported bill into multiple files and compresses them into one or more zip files. The zip file name format is the same.
//
// @param tmpReq - CreateReportDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportDefinitionResponse
func (client *Client) CreateReportDefinitionWithContext(ctx context.Context, tmpReq *CreateReportDefinitionRequest, runtime *dara.RuntimeOptions) (_result *CreateReportDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateReportDefinitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SelectedFields) {
		request.SelectedFieldsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SelectedFields, dara.String("SelectedFields"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginBillingCycle) {
		query["BeginBillingCycle"] = request.BeginBillingCycle
	}

	if !dara.IsNil(request.IncludeMembers) {
		query["IncludeMembers"] = request.IncludeMembers
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.NotSendOnNoData) {
		query["NotSendOnNoData"] = request.NotSendOnNoData
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

	if !dara.IsNil(request.SelectedFieldsShrink) {
		query["SelectedFields"] = request.SelectedFieldsShrink
	}

	if !dara.IsNil(request.SendWithAttach) {
		query["SendWithAttach"] = request.SendWithAttach
	}

	if !dara.IsNil(request.SplitFileOnUserId) {
		query["SplitFileOnUserId"] = request.SplitFileOnUserId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a budget.
//
// @param request - DeleteBudgetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBudgetResponse
func (client *Client) DeleteBudgetWithContext(ctx context.Context, request *DeleteBudgetRequest, runtime *dara.RuntimeOptions) (_result *DeleteBudgetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BudgetName) {
		body["BudgetName"] = request.BudgetName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBudget"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBudgetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Cost Center
//
// Description:
//
// This API is in canary release and is only available to whitelisted users. Excessive calls may cause performance issues such as response timeouts.
//
// @param request - DeleteCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostCenterResponse
func (client *Client) DeleteCostCenterWithContext(ctx context.Context, request *DeleteCostCenterRequest, runtime *dara.RuntimeOptions) (_result *DeleteCostCenterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete financial unit automatic allocation rule
//
// Description:
//
// This API is in canary release and is only available to whitelisted users. Excessive calls may cause performance issues such as response timeouts.
//
// @param tmpReq - DeleteCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostCenterRuleResponse
func (client *Client) DeleteCostCenterRuleWithContext(ctx context.Context, tmpReq *DeleteCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteCostCenterRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DeleteCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCouponDeductTagResponse
func (client *Client) DeleteCouponDeductTagWithContext(ctx context.Context, tmpReq *DeleteCouponDeductTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteCouponDeductTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a bill report export subscription.
//
// @param request - DeleteReportDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReportDefinitionResponse
func (client *Client) DeleteReportDefinitionWithContext(ctx context.Context, request *DeleteReportDefinitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteReportDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a Single Budget
//
// @param request - DescribeBudgetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBudgetResponse
func (client *Client) DescribeBudgetWithContext(ctx context.Context, request *DescribeBudgetRequest, runtime *dara.RuntimeOptions) (_result *DescribeBudgetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BudgetName) {
		body["BudgetName"] = request.BudgetName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBudget"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBudgetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of budgets.
//
// @param request - DescribeBudgetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBudgetsResponse
func (client *Client) DescribeBudgetsWithContext(ctx context.Context, request *DescribeBudgetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBudgetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BudgetName) {
		body["BudgetName"] = request.BudgetName
	}

	if !dara.IsNil(request.BudgetType) {
		body["BudgetType"] = request.BudgetType
	}

	if !dara.IsNil(request.ExpireStatus) {
		body["ExpireStatus"] = request.ExpireStatus
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBudgets"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBudgetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of coupons.
//
// @param tmpReq - DescribeCouponRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponResponse
func (client *Client) DescribeCouponWithContext(ctx context.Context, tmpReq *DescribeCouponRequest, runtime *dara.RuntimeOptions) (_result *DescribeCouponResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCouponShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CouponTemplateIdList) {
		request.CouponTemplateIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CouponTemplateIdList, dara.String("CouponTemplateIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponId) {
		query["CouponId"] = request.CouponId
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.CouponTemplateIdListShrink) {
		query["CouponTemplateIdList"] = request.CouponTemplateIdListShrink
	}

	if !dara.IsNil(request.CouponType) {
		query["CouponType"] = request.CouponType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.EffectiveEndTime) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !dara.IsNil(request.EffectiveStartTime) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !dara.IsNil(request.ExpireEndDate) {
		query["ExpireEndDate"] = request.ExpireEndDate
	}

	if !dara.IsNil(request.ExpireStartDate) {
		query["ExpireStartDate"] = request.ExpireStartDate
	}

	if !dara.IsNil(request.IncludeShare) {
		query["IncludeShare"] = request.IncludeShare
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

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCoupon"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCouponResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of products available for a coupon.
//
// @param tmpReq - DescribeCouponItemListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponItemListResponse
func (client *Client) DescribeCouponItemListWithContext(ctx context.Context, tmpReq *DescribeCouponItemListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCouponItemListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCouponItemListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponId) {
		query["CouponId"] = request.CouponId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCouponItemList"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCouponItemListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource plan deduction records.
//
// @param tmpReq - DescribeDeductLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeductLogsResponse
func (client *Client) DescribeDeductLogsWithContext(ctx context.Context, tmpReq *DescribeDeductLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeductLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeDeductLogsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelationAccountIds) {
		request.RelationAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelationAccountIds, dara.String("RelationAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BillInstanceId) {
		query["BillInstanceId"] = request.BillInstanceId
	}

	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BillingCommodityCode) {
		body["BillingCommodityCode"] = request.BillingCommodityCode
	}

	if !dara.IsNil(request.BillingEndTime) {
		body["BillingEndTime"] = request.BillingEndTime
	}

	if !dara.IsNil(request.BillingStartTime) {
		body["BillingStartTime"] = request.BillingStartTime
	}

	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.Group) {
		body["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RelationAccountIdsShrink) {
		body["RelationAccountIds"] = request.RelationAccountIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeductLogs"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeductLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource plan instances.
//
// @param tmpReq - DescribeFrInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFrInstancesResponse
func (client *Client) DescribeFrInstancesWithContext(ctx context.Context, tmpReq *DescribeFrInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeFrInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeFrInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.InstanceTag) {
		query["InstanceTag"] = request.InstanceTag
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CapacityType) {
		body["CapacityType"] = request.CapacityType
	}

	if !dara.IsNil(request.CommodityCode) {
		body["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Group) {
		body["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNum) {
		body["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SortRule) {
		body["SortRule"] = request.SortRule
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFrInstances"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFrInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DescribeUserSpnSummaryInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserSpnSummaryInfoResponse
func (client *Client) DescribeUserSpnSummaryInfoWithContext(ctx context.Context, tmpReq *DescribeUserSpnSummaryInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserSpnSummaryInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query available balance of fund account
//
// @param request - GetFundAccountAvailableAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountAvailableAmountResponse
func (client *Client) GetFundAccountAvailableAmountWithContext(ctx context.Context, request *GetFundAccountAvailableAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountAvailableAmountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query allocatable credit limit of a fund account
//
// @param request - GetFundAccountCanAllocateCreditAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanAllocateCreditAmountResponse
func (client *Client) GetFundAccountCanAllocateCreditAmountWithContext(ctx context.Context, request *GetFundAccountCanAllocateCreditAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanAllocateCreditAmountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the reclaimable amount of a fund account.
//
// @param request - GetFundAccountCanRecycleAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanRecycleAmountResponse
func (client *Client) GetFundAccountCanRecycleAmountWithContext(ctx context.Context, request *GetFundAccountCanRecycleAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanRecycleAmountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the transferable amount of a fund account
//
// @param request - GetFundAccountCanTransferAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanTransferAmountResponse
func (client *Client) GetFundAccountCanTransferAmountWithContext(ctx context.Context, request *GetFundAccountCanTransferAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanTransferAmountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Withdrawable Amount of Fund Account
//
// @param request - GetFundAccountCanWithdrawAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanWithdrawAmountResponse
func (client *Client) GetFundAccountCanWithdrawAmountWithContext(ctx context.Context, request *GetFundAccountCanWithdrawAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanWithdrawAmountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Fund Account Low Balance Alert
//
// @param request - GetFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) GetFundAccountLowAvailableAmountAlarmWithContext(ctx context.Context, request *GetFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query fund account transaction details
//
// @param tmpReq - GetFundAccountTransactionDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountTransactionDetailsResponse
func (client *Client) GetFundAccountTransactionDetailsWithContext(ctx context.Context, tmpReq *GetFundAccountTransactionDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountTransactionDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific order for a user or a reseller\\"s customer.
//
// @param request - GetOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderDetailResponse
func (client *Client) GetOrderDetailWithContext(ctx context.Context, request *GetOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the order list of a user or a reseller customer. By default, this operation queries orders created within the most recent hour. To query orders over a longer time range, set the CreateTimeStart and CreateTimeEnd parameters.
//
// @param request - GetOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrdersResponse
func (client *Client) GetOrdersWithContext(ctx context.Context, request *GetOrdersRequest, runtime *dara.RuntimeOptions) (_result *GetOrdersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetSavingPlanDeductableCommodityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanDeductableCommodityResponse
func (client *Client) GetSavingPlanDeductableCommodityWithContext(ctx context.Context, tmpReq *GetSavingPlanDeductableCommodityRequest, runtime *dara.RuntimeOptions) (_result *GetSavingPlanDeductableCommodityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetSavingPlanShareAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanShareAccountsResponse
func (client *Client) GetSavingPlanShareAccountsWithContext(ctx context.Context, tmpReq *GetSavingPlanShareAccountsRequest, runtime *dara.RuntimeOptions) (_result *GetSavingPlanShareAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetSavingPlanUserDeductRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSavingPlanUserDeductRuleResponse
func (client *Client) GetSavingPlanUserDeductRuleWithContext(ctx context.Context, tmpReq *GetSavingPlanUserDeductRuleRequest, runtime *dara.RuntimeOptions) (_result *GetSavingPlanUserDeductRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListCouponDeductTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCouponDeductTagResponse
func (client *Client) ListCouponDeductTagWithContext(ctx context.Context, tmpReq *ListCouponDeductTagRequest, runtime *dara.RuntimeOptions) (_result *ListCouponDeductTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query fund account list
//
// @param request - ListFundAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFundAccountResponse
func (client *Client) ListFundAccountWithContext(ctx context.Context, request *ListFundAccountRequest, runtime *dara.RuntimeOptions) (_result *ListFundAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query payment relationships of an account
//
// @param request - ListFundAccountPayRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFundAccountPayRelationResponse
func (client *Client) ListFundAccountPayRelationWithContext(ctx context.Context, request *ListFundAccountPayRelationRequest, runtime *dara.RuntimeOptions) (_result *ListFundAccountPayRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query invoice candidate data, which can be used for invoicing.
//
// @param tmpReq - ListInvoiceCandidateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInvoiceCandidateResponse
func (client *Client) ListInvoiceCandidateWithContext(ctx context.Context, tmpReq *ListInvoiceCandidateRequest, runtime *dara.RuntimeOptions) (_result *ListInvoiceCandidateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of subscribed reports.
//
// @param request - ListReportDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportDefinitionsResponse
func (client *Client) ListReportDefinitionsWithContext(ctx context.Context, request *ListReportDefinitionsRequest, runtime *dara.RuntimeOptions) (_result *ListReportDefinitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify cost centers
//
// Description:
//
// Modifies one or more cost centers.
//
// @param tmpReq - ModifyCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostCenterResponse
func (client *Client) ModifyCostCenterWithContext(ctx context.Context, tmpReq *ModifyCostCenterRequest, runtime *dara.RuntimeOptions) (_result *ModifyCostCenterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify financial unit rules
//
// Description:
//
// # Modify one or more financial units
//
// @param tmpReq - ModifyCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostCenterRuleResponse
func (client *Client) ModifyCostCenterRuleWithContext(ctx context.Context, tmpReq *ModifyCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyCostCenterRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PayOrderResponse
func (client *Client) PayOrderWithContext(ctx context.Context, request *PayOrderRequest, runtime *dara.RuntimeOptions) (_result *PayOrderResponse, _err error) {
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
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PayOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query cost center expense overview
//
// Description:
//
// # Query cost center expense overview results for a specified billing period
//
// @param request - QueryCostByCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostByCostCenterResponse
func (client *Client) QueryCostByCostCenterWithContext(ctx context.Context, request *QueryCostByCostCenterRequest, runtime *dara.RuntimeOptions) (_result *QueryCostByCostCenterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingMonth) {
		query["BillingMonth"] = request.BillingMonth
	}

	if !dara.IsNil(request.DisplayZeroAmountBills) {
		query["DisplayZeroAmountBills"] = request.DisplayZeroAmountBills
	}

	if !dara.IsNil(request.GroupByCostCenterLevel) {
		query["GroupByCostCenterLevel"] = request.GroupByCostCenterLevel
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.OwnerAccountId) {
		query["OwnerAccountId"] = request.OwnerAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCostByCostCenter"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCostByCostCenterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries financial units.
//
// Description:
//
// Queries a parent financial unit and its child financial units.
//
// @param tmpReq - QueryCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterResponse
func (client *Client) QueryCostCenterWithContext(ctx context.Context, tmpReq *QueryCostCenterRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of resource instances that belong to a cost center of the user. When CostCenterId is 0, it queries unallocated primary and sub-resource instances.
//
// @param request - QueryCostCenterResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterResourceResponse
func (client *Client) QueryCostCenterResourceWithContext(ctx context.Context, request *QueryCostCenterResourceRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query cost center rules
//
// Description:
//
// Query parent cost center and its child cost centers.
//
// @param request - QueryCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterRuleResponse
func (client *Client) QueryCostCenterRuleWithContext(ctx context.Context, request *QueryCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query cost center sharing rules
//
// Description:
//
// Queries the sharing rules of user cost centers.
//
// @param request - QueryCostCenterShareRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterShareRuleResponse
func (client *Client) QueryCostCenterShareRuleWithContext(ctx context.Context, request *QueryCostCenterShareRuleRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterShareRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.OwnerAccountId) {
		query["OwnerAccountId"] = request.OwnerAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCostCenterShareRule"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCostCenterShareRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SLA compensation list for a user.
//
// Description:
//
// Provides the SLA compensation details list for a user. Only data from the last two months is available.
//
// @param request - QueryMonthlySlaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMonthlySlaListResponse
func (client *Client) QueryMonthlySlaListWithContext(ctx context.Context, request *QueryMonthlySlaListRequest, runtime *dara.RuntimeOptions) (_result *QueryMonthlySlaListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EcIdAccountIds) {
		query["EcIdAccountIds"] = request.EcIdAccountIds
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		body["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Months) {
		body["Months"] = request.Months
	}

	if !dara.IsNil(request.PayStatuses) {
		body["PayStatuses"] = request.PayStatuses
	}

	if !dara.IsNil(request.ProductCodes) {
		body["ProductCodes"] = request.ProductCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMonthlySlaList"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMonthlySlaListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies cost center sharing rules, including creating, modifying, and deleting sharing rules.
//
// @param tmpReq - SaveCostCenterShareRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveCostCenterShareRuleResponse
func (client *Client) SaveCostCenterShareRuleWithContext(ctx context.Context, tmpReq *SaveCostCenterShareRuleRequest, runtime *dara.RuntimeOptions) (_result *SaveCostCenterShareRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveCostCenterShareRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateShareRuleList) {
		request.CreateShareRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateShareRuleList, dara.String("CreateShareRuleList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ModifyShareRuleList) {
		request.ModifyShareRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModifyShareRuleList, dara.String("ModifyShareRuleList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RemoveShareRuleList) {
		request.RemoveShareRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RemoveShareRuleList, dara.String("RemoveShareRuleList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateShareRuleListShrink) {
		query["CreateShareRuleList"] = request.CreateShareRuleListShrink
	}

	if !dara.IsNil(request.ModifyShareRuleListShrink) {
		query["ModifyShareRuleList"] = request.ModifyShareRuleListShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	if !dara.IsNil(request.OwnerAccountId) {
		query["OwnerAccountId"] = request.OwnerAccountId
	}

	if !dara.IsNil(request.RemoveShareRuleListShrink) {
		query["RemoveShareRuleList"] = request.RemoveShareRuleListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveCostCenterShareRule"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveCostCenterShareRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Set the credit control limit for a fund account
//
// @param request - SetFundAccountCreditAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFundAccountCreditAmountResponse
func (client *Client) SetFundAccountCreditAmountWithContext(ctx context.Context, request *SetFundAccountCreditAmountRequest, runtime *dara.RuntimeOptions) (_result *SetFundAccountCreditAmountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Set Fund Account Low Balance Alert
//
// @param request - SetFundAccountLowAvailableAmountAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) SetFundAccountLowAvailableAmountAlarmWithContext(ctx context.Context, request *SetFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *SetFundAccountLowAvailableAmountAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - SetSavingPlanUserDeductRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSavingPlanUserDeductRuleResponse
func (client *Client) SetSavingPlanUserDeductRuleWithContext(ctx context.Context, tmpReq *SetSavingPlanUserDeductRuleRequest, runtime *dara.RuntimeOptions) (_result *SetSavingPlanUserDeductRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// User claims coupons for the last two months.
//
// Description:
//
// 1. Call QueryMonthlySlaList to obtain the claimable months and records.
//
// 2. Claim by month or by record.
//
// Note: Only compensation for the last two months can be claimed. Historical compensation has been automatically issued.
//
// @param request - SubmitSlaCouponApplyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSlaCouponApplyResponse
func (client *Client) SubmitSlaCouponApplyWithContext(ctx context.Context, request *SubmitSlaCouponApplyRequest, runtime *dara.RuntimeOptions) (_result *SubmitSlaCouponApplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIds) {
		query["EcIdAccountIds"] = request.EcIdAccountIds
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DamagedIds) {
		body["DamagedIds"] = request.DamagedIds
	}

	if !dara.IsNil(request.Month) {
		body["Month"] = request.Month
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSlaCouponApply"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSlaCouponApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a budget.
//
// @param tmpReq - UpdateBudgetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBudgetResponse
func (client *Client) UpdateBudgetWithContext(ctx context.Context, tmpReq *UpdateBudgetRequest, runtime *dara.RuntimeOptions) (_result *UpdateBudgetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBudgetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CycleQuota) {
		request.CycleQuotaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CycleQuota, dara.String("CycleQuota"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EcIdAccountIds) {
		request.EcIdAccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EcIdAccountIds, dara.String("EcIdAccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueryFilter) {
		request.QueryFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryFilter, dara.String("QueryFilter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.WarnConfs) {
		request.WarnConfsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WarnConfs, dara.String("WarnConfs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EcIdAccountIdsShrink) {
		query["EcIdAccountIds"] = request.EcIdAccountIdsShrink
	}

	if !dara.IsNil(request.Nbid) {
		query["Nbid"] = request.Nbid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BudgetName) {
		body["BudgetName"] = request.BudgetName
	}

	if !dara.IsNil(request.BudgetType) {
		body["BudgetType"] = request.BudgetType
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.CycleEndPeriod) {
		body["CycleEndPeriod"] = request.CycleEndPeriod
	}

	if !dara.IsNil(request.CycleQuotaShrink) {
		body["CycleQuota"] = request.CycleQuotaShrink
	}

	if !dara.IsNil(request.CycleStartPeriod) {
		body["CycleStartPeriod"] = request.CycleStartPeriod
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.Metric) {
		body["Metric"] = request.Metric
	}

	if !dara.IsNil(request.OriginalBudgetName) {
		body["OriginalBudgetName"] = request.OriginalBudgetName
	}

	if !dara.IsNil(request.QueryFilterShrink) {
		body["QueryFilter"] = request.QueryFilterShrink
	}

	if !dara.IsNil(request.Quota) {
		body["Quota"] = request.Quota
	}

	if !dara.IsNil(request.QuotaType) {
		body["QuotaType"] = request.QuotaType
	}

	if !dara.IsNil(request.WarnConfsShrink) {
		body["WarnConfs"] = request.WarnConfsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBudget"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBudgetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
