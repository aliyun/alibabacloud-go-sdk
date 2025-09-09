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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - AllocateCostCenterResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateCostCenterResourceResponse
func (client *Client) AllocateCostCenterResourceWithContext(ctx context.Context, tmpReq *AllocateCostCenterResourceRequest, runtime *dara.RuntimeOptions) (_result *AllocateCostCenterResourceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelFundAccountLowAvailableAmountAlarmResponse
func (client *Client) CancelFundAccountLowAvailableAmountAlarmWithContext(ctx context.Context, request *CancelFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *CancelFundAccountLowAvailableAmountAlarmResponse, _err error) {
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
// 创建财务单元
//
// @param tmpReq - CreateCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostCenterResponse
func (client *Client) CreateCostCenterWithContext(ctx context.Context, tmpReq *CreateCostCenterRequest, runtime *dara.RuntimeOptions) (_result *CreateCostCenterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCostCenterRuleResponse
func (client *Client) CreateCostCenterRuleWithContext(ctx context.Context, tmpReq *CreateCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateCostCenterRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateFundAccountPayRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFundAccountPayRelationResponse
func (client *Client) CreateFundAccountPayRelationWithContext(ctx context.Context, tmpReq *CreateFundAccountPayRelationRequest, runtime *dara.RuntimeOptions) (_result *CreateFundAccountPayRelationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFundAccountTransferResponse
func (client *Client) CreateFundAccountTransferWithContext(ctx context.Context, request *CreateFundAccountTransferRequest, runtime *dara.RuntimeOptions) (_result *CreateFundAccountTransferResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateInvoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInvoiceResponse
func (client *Client) CreateInvoiceWithContext(ctx context.Context, tmpReq *CreateInvoiceRequest, runtime *dara.RuntimeOptions) (_result *CreateInvoiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportDefinitionResponse
func (client *Client) CreateReportDefinitionWithContext(ctx context.Context, request *CreateReportDefinitionRequest, runtime *dara.RuntimeOptions) (_result *CreateReportDefinitionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除财务单元
//
// @param request - DeleteCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostCenterResponse
func (client *Client) DeleteCostCenterWithContext(ctx context.Context, request *DeleteCostCenterRequest, runtime *dara.RuntimeOptions) (_result *DeleteCostCenterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DeleteCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCostCenterRuleResponse
func (client *Client) DeleteCostCenterRuleWithContext(ctx context.Context, tmpReq *DeleteCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteCostCenterRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReportDefinitionResponse
func (client *Client) DeleteReportDefinitionWithContext(ctx context.Context, request *DeleteReportDefinitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteReportDefinitionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DescribeCouponRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponResponse
func (client *Client) DescribeCouponWithContext(ctx context.Context, tmpReq *DescribeCouponRequest, runtime *dara.RuntimeOptions) (_result *DescribeCouponResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DescribeCouponItemListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCouponItemListResponse
func (client *Client) DescribeCouponItemListWithContext(ctx context.Context, tmpReq *DescribeCouponItemListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCouponItemListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountAvailableAmountResponse
func (client *Client) GetFundAccountAvailableAmountWithContext(ctx context.Context, request *GetFundAccountAvailableAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountAvailableAmountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanAllocateCreditAmountResponse
func (client *Client) GetFundAccountCanAllocateCreditAmountWithContext(ctx context.Context, request *GetFundAccountCanAllocateCreditAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanAllocateCreditAmountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanRecycleAmountResponse
func (client *Client) GetFundAccountCanRecycleAmountWithContext(ctx context.Context, request *GetFundAccountCanRecycleAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanRecycleAmountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanTransferAmountResponse
func (client *Client) GetFundAccountCanTransferAmountWithContext(ctx context.Context, request *GetFundAccountCanTransferAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanTransferAmountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountCanWithdrawAmountResponse
func (client *Client) GetFundAccountCanWithdrawAmountWithContext(ctx context.Context, request *GetFundAccountCanWithdrawAmountRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountCanWithdrawAmountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) GetFundAccountLowAvailableAmountAlarmWithContext(ctx context.Context, request *GetFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountLowAvailableAmountAlarmResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetFundAccountTransactionDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFundAccountTransactionDetailsResponse
func (client *Client) GetFundAccountTransactionDetailsWithContext(ctx context.Context, tmpReq *GetFundAccountTransactionDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetFundAccountTransactionDetailsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderDetailResponse
func (client *Client) GetOrderDetailWithContext(ctx context.Context, request *GetOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetOrderDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrdersResponse
func (client *Client) GetOrdersWithContext(ctx context.Context, request *GetOrdersRequest, runtime *dara.RuntimeOptions) (_result *GetOrdersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFundAccountResponse
func (client *Client) ListFundAccountWithContext(ctx context.Context, request *ListFundAccountRequest, runtime *dara.RuntimeOptions) (_result *ListFundAccountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFundAccountPayRelationResponse
func (client *Client) ListFundAccountPayRelationWithContext(ctx context.Context, request *ListFundAccountPayRelationRequest, runtime *dara.RuntimeOptions) (_result *ListFundAccountPayRelationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListInvoiceCandidateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInvoiceCandidateResponse
func (client *Client) ListInvoiceCandidateWithContext(ctx context.Context, tmpReq *ListInvoiceCandidateRequest, runtime *dara.RuntimeOptions) (_result *ListInvoiceCandidateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReportDefinitionsResponse
func (client *Client) ListReportDefinitionsWithContext(ctx context.Context, request *ListReportDefinitionsRequest, runtime *dara.RuntimeOptions) (_result *ListReportDefinitionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostCenterResponse
func (client *Client) ModifyCostCenterWithContext(ctx context.Context, tmpReq *ModifyCostCenterRequest, runtime *dara.RuntimeOptions) (_result *ModifyCostCenterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ModifyCostCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCostCenterRuleResponse
func (client *Client) ModifyCostCenterRuleWithContext(ctx context.Context, tmpReq *ModifyCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyCostCenterRuleResponse, _err error) {
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 查询财务单元
//
// @param tmpReq - QueryCostCenterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterResponse
func (client *Client) QueryCostCenterWithContext(ctx context.Context, tmpReq *QueryCostCenterRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterResourceResponse
func (client *Client) QueryCostCenterResourceWithContext(ctx context.Context, request *QueryCostCenterResourceRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterResourceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterRuleResponse
func (client *Client) QueryCostCenterRuleWithContext(ctx context.Context, request *QueryCostCenterRuleRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询财务单元分摊规则
//
// @param request - QueryCostCenterShareRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCostCenterShareRuleResponse
func (client *Client) QueryCostCenterShareRuleWithContext(ctx context.Context, request *QueryCostCenterShareRuleRequest, runtime *dara.RuntimeOptions) (_result *QueryCostCenterShareRuleResponse, _err error) {
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
// 修改财务单元分摊规则
//
// @param tmpReq - SaveCostCenterShareRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveCostCenterShareRuleResponse
func (client *Client) SaveCostCenterShareRuleWithContext(ctx context.Context, tmpReq *SaveCostCenterShareRuleRequest, runtime *dara.RuntimeOptions) (_result *SaveCostCenterShareRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// 设置资金账户的信控限额
//
// @param request - SetFundAccountCreditAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFundAccountCreditAmountResponse
func (client *Client) SetFundAccountCreditAmountWithContext(ctx context.Context, request *SetFundAccountCreditAmountRequest, runtime *dara.RuntimeOptions) (_result *SetFundAccountCreditAmountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFundAccountLowAvailableAmountAlarmResponse
func (client *Client) SetFundAccountLowAvailableAmountAlarmWithContext(ctx context.Context, request *SetFundAccountLowAvailableAmountAlarmRequest, runtime *dara.RuntimeOptions) (_result *SetFundAccountLowAvailableAmountAlarmResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
