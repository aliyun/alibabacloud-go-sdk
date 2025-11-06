// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 确认任务结果
//
// @param request - AcknowledgeTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcknowledgeTaskResultResponse
func (client *Client) AcknowledgeTaskResultWithContext(ctx context.Context, request *AcknowledgeTaskResultRequest, runtime *dara.RuntimeOptions) (_result *AcknowledgeTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TaskDetailNo) {
		query["TaskDetailNo"] = request.TaskDetailNo
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcknowledgeTaskResult"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcknowledgeTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过关键字进行批量模糊匹配
//
// @param request - BatchFuzzyMatchDomainSensitiveWordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchFuzzyMatchDomainSensitiveWordResponse
func (client *Client) BatchFuzzyMatchDomainSensitiveWordWithContext(ctx context.Context, request *BatchFuzzyMatchDomainSensitiveWordRequest, runtime *dara.RuntimeOptions) (_result *BatchFuzzyMatchDomainSensitiveWordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchFuzzyMatchDomainSensitiveWord"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchFuzzyMatchDomainSensitiveWordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels real-name verification for a domain name.
//
// @param request - CancelDomainVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDomainVerificationResponse
func (client *Client) CancelDomainVerificationWithContext(ctx context.Context, request *CancelDomainVerificationRequest, runtime *dara.RuntimeOptions) (_result *CancelDomainVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDomainVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDomainVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消审核
//
// @param request - CancelOperationAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOperationAuditResponse
func (client *Client) CancelOperationAuditWithContext(ctx context.Context, request *CancelOperationAuditRequest, runtime *dara.RuntimeOptions) (_result *CancelOperationAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditRecordId) {
		query["AuditRecordId"] = request.AuditRecordId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelOperationAudit"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelOperationAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CancelQualificationVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelQualificationVerificationResponse
func (client *Client) CancelQualificationVerificationWithContext(ctx context.Context, request *CancelQualificationVerificationRequest, runtime *dara.RuntimeOptions) (_result *CancelQualificationVerificationResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.QualificationType) {
		query["QualificationType"] = request.QualificationType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelQualificationVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelQualificationVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CancelTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithContext(ctx context.Context, request *CancelTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TaskNo) {
		query["TaskNo"] = request.TaskNo
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例所在资源组
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDomainResponse
func (client *Client) CheckDomainWithContext(ctx context.Context, request *CheckDomainRequest, runtime *dara.RuntimeOptions) (_result *CheckDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FeeCommand) {
		query["FeeCommand"] = request.FeeCommand
	}

	if !dara.IsNil(request.FeeCurrency) {
		query["FeeCurrency"] = request.FeeCurrency
	}

	if !dara.IsNil(request.FeePeriod) {
		query["FeePeriod"] = request.FeePeriod
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDomain"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckDomainSunriseClaimRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDomainSunriseClaimResponse
func (client *Client) CheckDomainSunriseClaimWithContext(ctx context.Context, request *CheckDomainSunriseClaimRequest, runtime *dara.RuntimeOptions) (_result *CheckDomainSunriseClaimResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDomainSunriseClaim"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDomainSunriseClaimResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验在售国际一口价域名状态和询价
//
// @param request - CheckIntlFixPriceDomainStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckIntlFixPriceDomainStatusResponse
func (client *Client) CheckIntlFixPriceDomainStatusWithContext(ctx context.Context, request *CheckIntlFixPriceDomainStatusRequest, runtime *dara.RuntimeOptions) (_result *CheckIntlFixPriceDomainStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckIntlFixPriceDomainStatus"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckIntlFixPriceDomainStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckMaxYearOfServerLockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckMaxYearOfServerLockResponse
func (client *Client) CheckMaxYearOfServerLockWithContext(ctx context.Context, request *CheckMaxYearOfServerLockRequest, runtime *dara.RuntimeOptions) (_result *CheckMaxYearOfServerLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckAction) {
		query["CheckAction"] = request.CheckAction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckMaxYearOfServerLock"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckMaxYearOfServerLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckProcessingServerLockApplyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckProcessingServerLockApplyResponse
func (client *Client) CheckProcessingServerLockApplyWithContext(ctx context.Context, request *CheckProcessingServerLockApplyRequest, runtime *dara.RuntimeOptions) (_result *CheckProcessingServerLockApplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FeePeriod) {
		query["FeePeriod"] = request.FeePeriod
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckProcessingServerLockApply"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckProcessingServerLockApplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckTransferInFeasibilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckTransferInFeasibilityResponse
func (client *Client) CheckTransferInFeasibilityWithContext(ctx context.Context, request *CheckTransferInFeasibilityRequest, runtime *dara.RuntimeOptions) (_result *CheckTransferInFeasibilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TransferAuthorizationCode) {
		query["TransferAuthorizationCode"] = request.TransferAuthorizationCode
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckTransferInFeasibility"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckTransferInFeasibilityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConfirmTransferInEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmTransferInEmailResponse
func (client *Client) ConfirmTransferInEmailWithContext(ctx context.Context, request *ConfirmTransferInEmailRequest, runtime *dara.RuntimeOptions) (_result *ConfirmTransferInEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmTransferInEmail"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmTransferInEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建国际一口价订单
//
// @param request - CreateIntlFixedPriceDomainOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntlFixedPriceDomainOrderResponse
func (client *Client) CreateIntlFixedPriceDomainOrderWithContext(ctx context.Context, request *CreateIntlFixedPriceDomainOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateIntlFixedPriceDomainOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ExpectedPrice) {
		query["ExpectedPrice"] = request.ExpectedPrice
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIntlFixedPriceDomainOrder"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIntlFixedPriceDomainOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除联系人模板
//
// @param request - DeleteContactTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactTemplatesResponse
func (client *Client) DeleteContactTemplatesWithContext(ctx context.Context, request *DeleteContactTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegistrantProfileIds) {
		query["RegistrantProfileIds"] = request.RegistrantProfileIds
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactTemplates"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除域名分组
//
// @param request - DeleteDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainGroupResponse
func (client *Client) DeleteDomainGroupWithContext(ctx context.Context, request *DeleteDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainGroupId) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainGroup"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除邮箱验证
//
// @param request - DeleteEmailVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEmailVerificationResponse
func (client *Client) DeleteEmailVerificationWithContext(ctx context.Context, request *DeleteEmailVerificationRequest, runtime *dara.RuntimeOptions) (_result *DeleteEmailVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEmailVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEmailVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除联系人模板
//
// @param request - DeleteRegistrantProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistrantProfileResponse
func (client *Client) DeleteRegistrantProfileWithContext(ctx context.Context, request *DeleteRegistrantProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteRegistrantProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRegistrantProfile"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRegistrantProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消域名特殊业务流程
//
// @param request - DomainSpecialBizCancelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DomainSpecialBizCancelResponse
func (client *Client) DomainSpecialBizCancelWithContext(ctx context.Context, request *DomainSpecialBizCancelRequest, runtime *dara.RuntimeOptions) (_result *DomainSpecialBizCancelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DomainSpecialBizCancel"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DomainSpecialBizCancelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 邮箱验证通过
//
// @param request - EmailVerifiedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EmailVerifiedResponse
func (client *Client) EmailVerifiedWithContext(ctx context.Context, request *EmailVerifiedRequest, runtime *dara.RuntimeOptions) (_result *EmailVerifiedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EmailVerified"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EmailVerifiedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过关键字进行模糊匹配
//
// @param request - FuzzyMatchDomainSensitiveWordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FuzzyMatchDomainSensitiveWordResponse
func (client *Client) FuzzyMatchDomainSensitiveWordWithContext(ctx context.Context, request *FuzzyMatchDomainSensitiveWordRequest, runtime *dara.RuntimeOptions) (_result *FuzzyMatchDomainSensitiveWordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FuzzyMatchDomainSensitiveWord"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FuzzyMatchDomainSensitiveWordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of domain names for fixed-price orders at the international site (alibabacloud.com).
//
// @param request - GetIntlFixPriceDomainListUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntlFixPriceDomainListUrlResponse
func (client *Client) GetIntlFixPriceDomainListUrlWithContext(ctx context.Context, request *GetIntlFixPriceDomainListUrlRequest, runtime *dara.RuntimeOptions) (_result *GetIntlFixPriceDomainListUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListDate) {
		query["ListDate"] = request.ListDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIntlFixPriceDomainListUrl"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIntlFixPriceDomainListUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetOperationOssUploadPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationOssUploadPolicyResponse
func (client *Client) GetOperationOssUploadPolicyWithContext(ctx context.Context, request *GetOperationOssUploadPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetOperationOssUploadPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditType) {
		query["AuditType"] = request.AuditType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOperationOssUploadPolicy"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOperationOssUploadPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetQualificationUploadPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualificationUploadPolicyResponse
func (client *Client) GetQualificationUploadPolicyWithContext(ctx context.Context, request *GetQualificationUploadPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetQualificationUploadPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualificationUploadPolicy"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualificationUploadPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListEmailVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEmailVerificationResponse
func (client *Client) ListEmailVerificationWithContext(ctx context.Context, request *ListEmailVerificationRequest, runtime *dara.RuntimeOptions) (_result *ListEmailVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginCreateTime) {
		query["BeginCreateTime"] = request.BeginCreateTime
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.EndCreateTime) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.VerificationStatus) {
		query["VerificationStatus"] = request.VerificationStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEmailVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEmailVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about domain names for which registry locks are enabled.
//
// @param request - ListServerLockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServerLockResponse
func (client *Client) ListServerLockWithContext(ctx context.Context, request *ListServerLockRequest, runtime *dara.RuntimeOptions) (_result *ListServerLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginStartDate) {
		query["BeginStartDate"] = request.BeginStartDate
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndExpireDate) {
		query["EndExpireDate"] = request.EndExpireDate
	}

	if !dara.IsNil(request.EndStartDate) {
		query["EndStartDate"] = request.EndStartDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LockProductId) {
		query["LockProductId"] = request.LockProductId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderByType) {
		query["OrderByType"] = request.OrderByType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServerLockStatus) {
		query["ServerLockStatus"] = request.ServerLockStatus
	}

	if !dara.IsNil(request.StartExpireDate) {
		query["StartExpireDate"] = request.StartExpireDate
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServerLock"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServerLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - LookupTmchNoticeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LookupTmchNoticeResponse
func (client *Client) LookupTmchNoticeWithContext(ctx context.Context, request *LookupTmchNoticeRequest, runtime *dara.RuntimeOptions) (_result *LookupTmchNoticeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClaimKey) {
		query["ClaimKey"] = request.ClaimKey
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LookupTmchNotice"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LookupTmchNoticeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PollTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PollTaskResultResponse
func (client *Client) PollTaskResultWithContext(ctx context.Context, request *PollTaskResultRequest, runtime *dara.RuntimeOptions) (_result *PollTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskNo) {
		query["TaskNo"] = request.TaskNo
	}

	if !dara.IsNil(request.TaskResultStatus) {
		query["TaskResultStatus"] = request.TaskResultStatus
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PollTaskResult"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PollTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索域名列表
//
// @param request - QueryAdvancedDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAdvancedDomainListResponse
func (client *Client) QueryAdvancedDomainListWithContext(ctx context.Context, request *QueryAdvancedDomainListRequest, runtime *dara.RuntimeOptions) (_result *QueryAdvancedDomainListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainGroupId) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !dara.IsNil(request.DomainNameSort) {
		query["DomainNameSort"] = request.DomainNameSort
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !dara.IsNil(request.EndExpirationDate) {
		query["EndExpirationDate"] = request.EndExpirationDate
	}

	if !dara.IsNil(request.EndLength) {
		query["EndLength"] = request.EndLength
	}

	if !dara.IsNil(request.EndRegistrationDate) {
		query["EndRegistrationDate"] = request.EndRegistrationDate
	}

	if !dara.IsNil(request.Excluded) {
		query["Excluded"] = request.Excluded
	}

	if !dara.IsNil(request.ExcludedPrefix) {
		query["ExcludedPrefix"] = request.ExcludedPrefix
	}

	if !dara.IsNil(request.ExcludedSuffix) {
		query["ExcludedSuffix"] = request.ExcludedSuffix
	}

	if !dara.IsNil(request.ExpirationDateSort) {
		query["ExpirationDateSort"] = request.ExpirationDateSort
	}

	if !dara.IsNil(request.Form) {
		query["Form"] = request.Form
	}

	if !dara.IsNil(request.IsPremiumDomain) {
		query["IsPremiumDomain"] = request.IsPremiumDomain
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.KeyWordPrefix) {
		query["KeyWordPrefix"] = request.KeyWordPrefix
	}

	if !dara.IsNil(request.KeyWordSuffix) {
		query["KeyWordSuffix"] = request.KeyWordSuffix
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductDomainType) {
		query["ProductDomainType"] = request.ProductDomainType
	}

	if !dara.IsNil(request.ProductDomainTypeSort) {
		query["ProductDomainTypeSort"] = request.ProductDomainTypeSort
	}

	if !dara.IsNil(request.RegistrationDateSort) {
		query["RegistrationDateSort"] = request.RegistrationDateSort
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartExpirationDate) {
		query["StartExpirationDate"] = request.StartExpirationDate
	}

	if !dara.IsNil(request.StartLength) {
		query["StartLength"] = request.StartLength
	}

	if !dara.IsNil(request.StartRegistrationDate) {
		query["StartRegistrationDate"] = request.StartRegistrationDate
	}

	if !dara.IsNil(request.Suffixs) {
		query["Suffixs"] = request.Suffixs
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TradeType) {
		query["TradeType"] = request.TradeType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAdvancedDomainList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAdvancedDomainListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryArtExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryArtExtensionResponse
func (client *Client) QueryArtExtensionWithContext(ctx context.Context, request *QueryArtExtensionRequest, runtime *dara.RuntimeOptions) (_result *QueryArtExtensionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryArtExtension"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryArtExtensionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询操作记录
//
// @param request - QueryChangeLogListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChangeLogListResponse
func (client *Client) QueryChangeLogListWithContext(ctx context.Context, request *QueryChangeLogListRequest, runtime *dara.RuntimeOptions) (_result *QueryChangeLogListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryChangeLogList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryChangeLogListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryContactInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContactInfoResponse
func (client *Client) QueryContactInfoWithContext(ctx context.Context, request *QueryContactInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryContactInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactType) {
		query["ContactType"] = request.ContactType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryContactInfo"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryContactInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDSRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDSRecordResponse
func (client *Client) QueryDSRecordWithContext(ctx context.Context, request *QueryDSRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryDSRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDSRecord"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDSRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDnsHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDnsHostResponse
func (client *Client) QueryDnsHostWithContext(ctx context.Context, request *QueryDnsHostRequest, runtime *dara.RuntimeOptions) (_result *QueryDnsHostResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDnsHost"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDnsHostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDomainAdminDivisionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainAdminDivisionResponse
func (client *Client) QueryDomainAdminDivisionWithContext(ctx context.Context, request *QueryDomainAdminDivisionRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainAdminDivisionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainAdminDivision"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainAdminDivisionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a domain name.
//
// @param request - QueryDomainByDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainByDomainNameResponse
func (client *Client) QueryDomainByDomainNameWithContext(ctx context.Context, request *QueryDomainByDomainNameRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainByDomainNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainByDomainName"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainByDomainNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据实例id查询域名信息
//
// @param request - QueryDomainByInstanceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainByInstanceIdResponse
func (client *Client) QueryDomainByInstanceIdWithContext(ctx context.Context, request *QueryDomainByInstanceIdRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainByInstanceIdResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainByInstanceId"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainByInstanceIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDomainGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainGroupListResponse
func (client *Client) QueryDomainGroupListWithContext(ctx context.Context, request *QueryDomainGroupListRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainGroupListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainGroupName) {
		query["DomainGroupName"] = request.DomainGroupName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ShowDeletingGroup) {
		query["ShowDeletingGroup"] = request.ShowDeletingGroup
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainGroupList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainGroupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of domain names within your Alibaba Cloud account by page.
//
// @param request - QueryDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainListResponse
func (client *Client) QueryDomainListWithContext(ctx context.Context, request *QueryDomainListRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ccompany) {
		query["Ccompany"] = request.Ccompany
	}

	if !dara.IsNil(request.DomainGroupId) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndExpirationDate) {
		query["EndExpirationDate"] = request.EndExpirationDate
	}

	if !dara.IsNil(request.EndRegistrationDate) {
		query["EndRegistrationDate"] = request.EndRegistrationDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderByType) {
		query["OrderByType"] = request.OrderByType
	}

	if !dara.IsNil(request.OrderKeyType) {
		query["OrderKeyType"] = request.OrderKeyType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductDomainType) {
		query["ProductDomainType"] = request.ProductDomainType
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.Registrar) {
		query["Registrar"] = request.Registrar
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartExpirationDate) {
		query["StartExpirationDate"] = request.StartExpirationDate
	}

	if !dara.IsNil(request.StartRegistrationDate) {
		query["StartRegistrationDate"] = request.StartRegistrationDate
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDomainRealNameVerificationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainRealNameVerificationInfoResponse
func (client *Client) QueryDomainRealNameVerificationInfoWithContext(ctx context.Context, request *QueryDomainRealNameVerificationInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainRealNameVerificationInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FetchImage) {
		query["FetchImage"] = request.FetchImage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainRealNameVerificationInfo"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainRealNameVerificationInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时查询域名价格
//
// @param tmpReq - QueryDomainRealTimePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainRealTimePriceResponse
func (client *Client) QueryDomainRealTimePriceWithContext(ctx context.Context, tmpReq *QueryDomainRealTimePriceRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainRealTimePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryDomainRealTimePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DomainItem) {
		request.DomainItemShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DomainItem, dara.String("DomainItem"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Currency) {
		query["Currency"] = request.Currency
	}

	if !dara.IsNil(request.DomainItemShrink) {
		query["DomainItem"] = request.DomainItemShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainRealTimePrice"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainRealTimePriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名特殊业务详情
//
// @param request - QueryDomainSpecialBizDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainSpecialBizDetailResponse
func (client *Client) QueryDomainSpecialBizDetailWithContext(ctx context.Context, request *QueryDomainSpecialBizDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainSpecialBizDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainSpecialBizDetail"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainSpecialBizDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过域名查询域名特殊业务详情
//
// @param request - QueryDomainSpecialBizInfoByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainSpecialBizInfoByDomainResponse
func (client *Client) QueryDomainSpecialBizInfoByDomainWithContext(ctx context.Context, request *QueryDomainSpecialBizInfoByDomainRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainSpecialBizInfoByDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainSpecialBizInfoByDomain"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainSpecialBizInfoByDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDomainSuffixRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainSuffixResponse
func (client *Client) QueryDomainSuffixWithContext(ctx context.Context, request *QueryDomainSuffixRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainSuffixResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainSuffix"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainSuffixResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询邮箱验证状态
//
// @param request - QueryEmailVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEmailVerificationResponse
func (client *Client) QueryEmailVerificationWithContext(ctx context.Context, request *QueryEmailVerificationRequest, runtime *dara.RuntimeOptions) (_result *QueryEmailVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEmailVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEmailVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryEnsAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEnsAssociationResponse
func (client *Client) QueryEnsAssociationWithContext(ctx context.Context, request *QueryEnsAssociationRequest, runtime *dara.RuntimeOptions) (_result *QueryEnsAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEnsAssociation"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEnsAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryFailReasonForDomainRealNameVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFailReasonForDomainRealNameVerificationResponse
func (client *Client) QueryFailReasonForDomainRealNameVerificationWithContext(ctx context.Context, request *QueryFailReasonForDomainRealNameVerificationRequest, runtime *dara.RuntimeOptions) (_result *QueryFailReasonForDomainRealNameVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RealNameVerificationAction) {
		query["RealNameVerificationAction"] = request.RealNameVerificationAction
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFailReasonForDomainRealNameVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFailReasonForDomainRealNameVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryFailReasonForRegistrantProfileRealNameVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFailReasonForRegistrantProfileRealNameVerificationResponse
func (client *Client) QueryFailReasonForRegistrantProfileRealNameVerificationWithContext(ctx context.Context, request *QueryFailReasonForRegistrantProfileRealNameVerificationRequest, runtime *dara.RuntimeOptions) (_result *QueryFailReasonForRegistrantProfileRealNameVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegistrantProfileID) {
		query["RegistrantProfileID"] = request.RegistrantProfileID
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFailReasonForRegistrantProfileRealNameVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFailReasonForRegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryFailingReasonListForQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFailingReasonListForQualificationResponse
func (client *Client) QueryFailingReasonListForQualificationWithContext(ctx context.Context, request *QueryFailingReasonListForQualificationRequest, runtime *dara.RuntimeOptions) (_result *QueryFailingReasonListForQualificationResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.QualificationType) {
		query["QualificationType"] = request.QualificationType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFailingReasonListForQualification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFailingReasonListForQualificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询国际一口价订单列表
//
// @param request - QueryIntlFixedPriceOrderListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIntlFixedPriceOrderListResponse
func (client *Client) QueryIntlFixedPriceOrderListWithContext(ctx context.Context, request *QueryIntlFixedPriceOrderListRequest, runtime *dara.RuntimeOptions) (_result *QueryIntlFixedPriceOrderListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryIntlFixedPriceOrderList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryIntlFixedPriceOrderListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryLocalEnsAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLocalEnsAssociationResponse
func (client *Client) QueryLocalEnsAssociationWithContext(ctx context.Context, request *QueryLocalEnsAssociationRequest, runtime *dara.RuntimeOptions) (_result *QueryLocalEnsAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLocalEnsAssociation"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLocalEnsAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryOperationAuditInfoDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOperationAuditInfoDetailResponse
func (client *Client) QueryOperationAuditInfoDetailWithContext(ctx context.Context, request *QueryOperationAuditInfoDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryOperationAuditInfoDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditRecordId) {
		query["AuditRecordId"] = request.AuditRecordId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOperationAuditInfoDetail"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOperationAuditInfoDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryOperationAuditInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOperationAuditInfoListResponse
func (client *Client) QueryOperationAuditInfoListWithContext(ctx context.Context, request *QueryOperationAuditInfoListRequest, runtime *dara.RuntimeOptions) (_result *QueryOperationAuditInfoListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.AuditType) {
		query["AuditType"] = request.AuditType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOperationAuditInfoList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOperationAuditInfoListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryQualificationDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryQualificationDetailResponse
func (client *Client) QueryQualificationDetailWithContext(ctx context.Context, request *QueryQualificationDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryQualificationDetailResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.QualificationType) {
		query["QualificationType"] = request.QualificationType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryQualificationDetail"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryQualificationDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryRegistrantProfileRealNameVerificationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRegistrantProfileRealNameVerificationInfoResponse
func (client *Client) QueryRegistrantProfileRealNameVerificationInfoWithContext(ctx context.Context, request *QueryRegistrantProfileRealNameVerificationInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryRegistrantProfileRealNameVerificationInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FetchImage) {
		query["FetchImage"] = request.FetchImage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRegistrantProfileRealNameVerificationInfo"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRegistrantProfileRealNameVerificationInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the registrant profiles that belong to your Alibaba Cloud account.
//
// Description:
//
// You can use optional request parameters to specify specific query criteria to query registrant profiles as required. For example:
//
//   - If you know the ID of the profile that you want to query, you can use the registrant profile ID parameter to query the detailed information about the profile.
//
//   - If you do not know the ID of the profile that you want to query, you can use parameters such as the registrant name parameter to query the detailed information about the profile.
//
// @param request - QueryRegistrantProfilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRegistrantProfilesResponse
func (client *Client) QueryRegistrantProfilesWithContext(ctx context.Context, request *QueryRegistrantProfilesRequest, runtime *dara.RuntimeOptions) (_result *QueryRegistrantProfilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultRegistrantProfile) {
		query["DefaultRegistrantProfile"] = request.DefaultRegistrantProfile
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RealNameStatus) {
		query["RealNameStatus"] = request.RealNameStatus
	}

	if !dara.IsNil(request.RegistrantOrganization) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.RegistrantProfileType) {
		query["RegistrantProfileType"] = request.RegistrantProfileType
	}

	if !dara.IsNil(request.RegistrantType) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZhRegistrantOrganization) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRegistrantProfiles"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRegistrantProfilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryServerLockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryServerLockResponse
func (client *Client) QueryServerLockWithContext(ctx context.Context, request *QueryServerLockRequest, runtime *dara.RuntimeOptions) (_result *QueryServerLockResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryServerLock"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryServerLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTaskDetailHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskDetailHistoryResponse
func (client *Client) QueryTaskDetailHistoryWithContext(ctx context.Context, request *QueryTaskDetailHistoryRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskDetailHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainNameCursor) {
		query["DomainNameCursor"] = request.DomainNameCursor
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskDetailNoCursor) {
		query["TaskDetailNoCursor"] = request.TaskDetailNoCursor
	}

	if !dara.IsNil(request.TaskNo) {
		query["TaskNo"] = request.TaskNo
	}

	if !dara.IsNil(request.TaskStatus) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskDetailHistory"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskDetailHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific domain name task by page.
//
// @param request - QueryTaskDetailListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskDetailListResponse
func (client *Client) QueryTaskDetailListWithContext(ctx context.Context, request *QueryTaskDetailListRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskDetailListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskNo) {
		query["TaskNo"] = request.TaskNo
	}

	if !dara.IsNil(request.TaskStatus) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskDetailList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskDetailListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTaskInfoHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskInfoHistoryResponse
func (client *Client) QueryTaskInfoHistoryWithContext(ctx context.Context, request *QueryTaskInfoHistoryRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskInfoHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginCreateTime) {
		query["BeginCreateTime"] = request.BeginCreateTime
	}

	if !dara.IsNil(request.CreateTimeCursor) {
		query["CreateTimeCursor"] = request.CreateTimeCursor
	}

	if !dara.IsNil(request.EndCreateTime) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskNoCursor) {
		query["TaskNoCursor"] = request.TaskNoCursor
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskInfoHistory"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskInfoHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param request - QueryTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskListResponse
func (client *Client) QueryTaskListWithContext(ctx context.Context, request *QueryTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginCreateTime) {
		query["BeginCreateTime"] = request.BeginCreateTime
	}

	if !dara.IsNil(request.EndCreateTime) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTransferInByInstanceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTransferInByInstanceIdResponse
func (client *Client) QueryTransferInByInstanceIdWithContext(ctx context.Context, request *QueryTransferInByInstanceIdRequest, runtime *dara.RuntimeOptions) (_result *QueryTransferInByInstanceIdResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTransferInByInstanceId"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTransferInByInstanceIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTransferInListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTransferInListResponse
func (client *Client) QueryTransferInListWithContext(ctx context.Context, request *QueryTransferInListRequest, runtime *dara.RuntimeOptions) (_result *QueryTransferInListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SimpleTransferInStatus) {
		query["SimpleTransferInStatus"] = request.SimpleTransferInStatus
	}

	if !dara.IsNil(request.SubmissionEndDate) {
		query["SubmissionEndDate"] = request.SubmissionEndDate
	}

	if !dara.IsNil(request.SubmissionStartDate) {
		query["SubmissionStartDate"] = request.SubmissionStartDate
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTransferInList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTransferInListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTransferOutInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTransferOutInfoResponse
func (client *Client) QueryTransferOutInfoWithContext(ctx context.Context, request *QueryTransferOutInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryTransferOutInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTransferOutInfo"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTransferOutInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存联系人模板实名资料
//
// @param request - RegistrantProfileRealNameVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegistrantProfileRealNameVerificationResponse
func (client *Client) RegistrantProfileRealNameVerificationWithContext(ctx context.Context, request *RegistrantProfileRealNameVerificationRequest, runtime *dara.RuntimeOptions) (_result *RegistrantProfileRealNameVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityCredentialNo) {
		query["IdentityCredentialNo"] = request.IdentityCredentialNo
	}

	if !dara.IsNil(request.IdentityCredentialType) {
		query["IdentityCredentialType"] = request.IdentityCredentialType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegistrantProfileID) {
		query["RegistrantProfileID"] = request.RegistrantProfileID
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityCredential) {
		body["IdentityCredential"] = request.IdentityCredential
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegistrantProfileRealNameVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新发送验证邮件
//
// @param request - ResendEmailVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendEmailVerificationResponse
func (client *Client) ResendEmailVerificationWithContext(ctx context.Context, request *ResendEmailVerificationRequest, runtime *dara.RuntimeOptions) (_result *ResendEmailVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResendEmailVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResendEmailVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置资质审核状态
//
// @param request - ResetQualificationVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetQualificationVerificationResponse
func (client *Client) ResetQualificationVerificationWithContext(ctx context.Context, request *ResetQualificationVerificationRequest, runtime *dara.RuntimeOptions) (_result *ResetQualificationVerificationResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetQualificationVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetQualificationVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量保存域名备注信息
//
// @param request - SaveBatchDomainRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchDomainRemarkResponse
func (client *Client) SaveBatchDomainRemarkWithContext(ctx context.Context, request *SaveBatchDomainRemarkRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchDomainRemarkResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchDomainRemark"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchDomainRemarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量申请域名快速转出
//
// @param request - SaveBatchTaskForApplyQuickTransferOutOpenlyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForApplyQuickTransferOutOpenlyResponse
func (client *Client) SaveBatchTaskForApplyQuickTransferOutOpenlyWithContext(ctx context.Context, request *SaveBatchTaskForApplyQuickTransferOutOpenlyRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForApplyQuickTransferOutOpenlyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForApplyQuickTransferOutOpenly"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForApplyQuickTransferOutOpenlyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存批量任务-注册订单
//
// @param request - SaveBatchTaskForCreatingOrderActivateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForCreatingOrderActivateResponse
func (client *Client) SaveBatchTaskForCreatingOrderActivateWithContext(ctx context.Context, request *SaveBatchTaskForCreatingOrderActivateRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForCreatingOrderActivateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderActivateParam) {
		query["OrderActivateParam"] = request.OrderActivateParam
	}

	if !dara.IsNil(request.PromotionNo) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !dara.IsNil(request.UseCoupon) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !dara.IsNil(request.UsePromotion) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForCreatingOrderActivate"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForCreatingOrderActivateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveBatchTaskForCreatingOrderRedeemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForCreatingOrderRedeemResponse
func (client *Client) SaveBatchTaskForCreatingOrderRedeemWithContext(ctx context.Context, request *SaveBatchTaskForCreatingOrderRedeemRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForCreatingOrderRedeemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderRedeemParam) {
		query["OrderRedeemParam"] = request.OrderRedeemParam
	}

	if !dara.IsNil(request.PromotionNo) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !dara.IsNil(request.UseCoupon) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !dara.IsNil(request.UsePromotion) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForCreatingOrderRedeem"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForCreatingOrderRedeemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存批量任务-续费订单
//
// @param request - SaveBatchTaskForCreatingOrderRenewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForCreatingOrderRenewResponse
func (client *Client) SaveBatchTaskForCreatingOrderRenewWithContext(ctx context.Context, request *SaveBatchTaskForCreatingOrderRenewRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForCreatingOrderRenewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderRenewParam) {
		query["OrderRenewParam"] = request.OrderRenewParam
	}

	if !dara.IsNil(request.PromotionNo) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !dara.IsNil(request.UseCoupon) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !dara.IsNil(request.UsePromotion) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForCreatingOrderRenew"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForCreatingOrderRenewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveBatchTaskForCreatingOrderTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForCreatingOrderTransferResponse
func (client *Client) SaveBatchTaskForCreatingOrderTransferWithContext(ctx context.Context, request *SaveBatchTaskForCreatingOrderTransferRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForCreatingOrderTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderTransferParam) {
		query["OrderTransferParam"] = request.OrderTransferParam
	}

	if !dara.IsNil(request.PromotionNo) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !dara.IsNil(request.UseCoupon) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !dara.IsNil(request.UsePromotion) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForCreatingOrderTransfer"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForCreatingOrderTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存批量任务-开启/关闭whois隐私保护锁
//
// @param request - SaveBatchTaskForDomainNameProxyServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForDomainNameProxyServiceResponse
func (client *Client) SaveBatchTaskForDomainNameProxyServiceWithContext(ctx context.Context, request *SaveBatchTaskForDomainNameProxyServiceRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForDomainNameProxyServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForDomainNameProxyService"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForDomainNameProxyServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交批量生成证书的任务
//
// @param tmpReq - SaveBatchTaskForGenerateDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForGenerateDomainCertificateResponse
func (client *Client) SaveBatchTaskForGenerateDomainCertificateWithContext(ctx context.Context, tmpReq *SaveBatchTaskForGenerateDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForGenerateDomainCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SaveBatchTaskForGenerateDomainCertificateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DomainNames) {
		request.DomainNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DomainNames, dara.String("DomainNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNamesShrink) {
		query["DomainNames"] = request.DomainNamesShrink
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForGenerateDomainCertificate"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForGenerateDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改dns
//
// @param request - SaveBatchTaskForModifyingDomainDnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForModifyingDomainDnsResponse
func (client *Client) SaveBatchTaskForModifyingDomainDnsWithContext(ctx context.Context, request *SaveBatchTaskForModifyingDomainDnsRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForModifyingDomainDnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunDns) {
		query["AliyunDns"] = request.AliyunDns
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainNameServer) {
		query["DomainNameServer"] = request.DomainNameServer
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForModifyingDomainDns"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForModifyingDomainDnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交批量预定删除抢注域名任务
//
// @param request - SaveBatchTaskForReserveDropListDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForReserveDropListDomainResponse
func (client *Client) SaveBatchTaskForReserveDropListDomainWithContext(ctx context.Context, request *SaveBatchTaskForReserveDropListDomainRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForReserveDropListDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForReserveDropListDomain"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForReserveDropListDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 基于转移码的批量转出任务提交
//
// @param request - SaveBatchTaskForTransferOutByAuthorizationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForTransferOutByAuthorizationCodeResponse
func (client *Client) SaveBatchTaskForTransferOutByAuthorizationCodeWithContext(ctx context.Context, request *SaveBatchTaskForTransferOutByAuthorizationCodeRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForTransferOutByAuthorizationCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TransferOutParamList) {
		query["TransferOutParamList"] = request.TransferOutParamList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForTransferOutByAuthorizationCode"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForTransferOutByAuthorizationCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存批量任务-开启/关闭禁止转移锁
//
// @param request - SaveBatchTaskForTransferProhibitionLockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForTransferProhibitionLockResponse
func (client *Client) SaveBatchTaskForTransferProhibitionLockWithContext(ctx context.Context, request *SaveBatchTaskForTransferProhibitionLockRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForTransferProhibitionLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForTransferProhibitionLock"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForTransferProhibitionLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveBatchTaskForUpdateProhibitionLockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForUpdateProhibitionLockResponse
func (client *Client) SaveBatchTaskForUpdateProhibitionLockWithContext(ctx context.Context, request *SaveBatchTaskForUpdateProhibitionLockRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForUpdateProhibitionLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForUpdateProhibitionLock"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForUpdateProhibitionLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 使用联系人信息修改联系人的批量任务
//
// @param request - SaveBatchTaskForUpdatingContactInfoByNewContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForUpdatingContactInfoByNewContactResponse
func (client *Client) SaveBatchTaskForUpdatingContactInfoByNewContactWithContext(ctx context.Context, request *SaveBatchTaskForUpdatingContactInfoByNewContactRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForUpdatingContactInfoByNewContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.ContactType) {
		query["ContactType"] = request.ContactType
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PostalCode) {
		query["PostalCode"] = request.PostalCode
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RegistrantName) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !dara.IsNil(request.RegistrantOrganization) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !dara.IsNil(request.RegistrantType) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !dara.IsNil(request.TelArea) {
		query["TelArea"] = request.TelArea
	}

	if !dara.IsNil(request.TelExt) {
		query["TelExt"] = request.TelExt
	}

	if !dara.IsNil(request.Telephone) {
		query["Telephone"] = request.Telephone
	}

	if !dara.IsNil(request.TransferOutProhibited) {
		query["TransferOutProhibited"] = request.TransferOutProhibited
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZhAddress) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !dara.IsNil(request.ZhCity) {
		query["ZhCity"] = request.ZhCity
	}

	if !dara.IsNil(request.ZhProvince) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !dara.IsNil(request.ZhRegistrantName) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !dara.IsNil(request.ZhRegistrantOrganization) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForUpdatingContactInfoByNewContact"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForUpdatingContactInfoByNewContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 使用模板修改联系人的批量任务
//
// @param request - SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse
func (client *Client) SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdWithContext(ctx context.Context, request *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactType) {
		query["ContactType"] = request.ContactType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.TransferOutProhibited) {
		query["TransferOutProhibited"] = request.TransferOutProhibited
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建/更新域名分组
//
// @param request - SaveDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveDomainGroupResponse
func (client *Client) SaveDomainGroupWithContext(ctx context.Context, request *SaveDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *SaveDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainGroupId) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !dara.IsNil(request.DomainGroupName) {
		query["DomainGroupName"] = request.DomainGroupName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveDomainGroup"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveDomainGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存联系人模板
//
// @param request - SaveRegistrantProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveRegistrantProfileResponse
func (client *Client) SaveRegistrantProfileWithContext(ctx context.Context, request *SaveRegistrantProfileRequest, runtime *dara.RuntimeOptions) (_result *SaveRegistrantProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.DefaultRegistrantProfile) {
		query["DefaultRegistrantProfile"] = request.DefaultRegistrantProfile
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PostalCode) {
		query["PostalCode"] = request.PostalCode
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RegistrantName) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !dara.IsNil(request.RegistrantOrganization) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.RegistrantProfileType) {
		query["RegistrantProfileType"] = request.RegistrantProfileType
	}

	if !dara.IsNil(request.RegistrantType) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !dara.IsNil(request.TelArea) {
		query["TelArea"] = request.TelArea
	}

	if !dara.IsNil(request.TelExt) {
		query["TelExt"] = request.TelExt
	}

	if !dara.IsNil(request.Telephone) {
		query["Telephone"] = request.Telephone
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZhAddress) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !dara.IsNil(request.ZhCity) {
		query["ZhCity"] = request.ZhCity
	}

	if !dara.IsNil(request.ZhProvince) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !dara.IsNil(request.ZhRegistrantName) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !dara.IsNil(request.ZhRegistrantOrganization) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveRegistrantProfile"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveRegistrantProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存联系人模板和凭据
//
// @param request - SaveRegistrantProfileRealNameVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveRegistrantProfileRealNameVerificationResponse
func (client *Client) SaveRegistrantProfileRealNameVerificationWithContext(ctx context.Context, request *SaveRegistrantProfileRealNameVerificationRequest, runtime *dara.RuntimeOptions) (_result *SaveRegistrantProfileRealNameVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.IdentityCredential) {
		query["IdentityCredential"] = request.IdentityCredential
	}

	if !dara.IsNil(request.IdentityCredentialNo) {
		query["IdentityCredentialNo"] = request.IdentityCredentialNo
	}

	if !dara.IsNil(request.IdentityCredentialType) {
		query["IdentityCredentialType"] = request.IdentityCredentialType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PostalCode) {
		query["PostalCode"] = request.PostalCode
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RegistrantName) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !dara.IsNil(request.RegistrantOrganization) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.RegistrantProfileType) {
		query["RegistrantProfileType"] = request.RegistrantProfileType
	}

	if !dara.IsNil(request.RegistrantType) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !dara.IsNil(request.TelArea) {
		query["TelArea"] = request.TelArea
	}

	if !dara.IsNil(request.TelExt) {
		query["TelExt"] = request.TelExt
	}

	if !dara.IsNil(request.Telephone) {
		query["Telephone"] = request.Telephone
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZhAddress) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !dara.IsNil(request.ZhCity) {
		query["ZhCity"] = request.ZhCity
	}

	if !dara.IsNil(request.ZhProvince) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !dara.IsNil(request.ZhRegistrantName) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !dara.IsNil(request.ZhRegistrantOrganization) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveRegistrantProfileRealNameVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveRegistrantProfileRealNameVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加dnsSec记录
//
// @param request - SaveSingleTaskForAddingDSRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForAddingDSRecordResponse
func (client *Client) SaveSingleTaskForAddingDSRecordWithContext(ctx context.Context, request *SaveSingleTaskForAddingDSRecordRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForAddingDSRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.Digest) {
		query["Digest"] = request.Digest
	}

	if !dara.IsNil(request.DigestType) {
		query["DigestType"] = request.DigestType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.KeyTag) {
		query["KeyTag"] = request.KeyTag
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForAddingDSRecord"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForAddingDSRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请域名快速转出
//
// @param request - SaveSingleTaskForApplyQuickTransferOutOpenlyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForApplyQuickTransferOutOpenlyResponse
func (client *Client) SaveSingleTaskForApplyQuickTransferOutOpenlyWithContext(ctx context.Context, request *SaveSingleTaskForApplyQuickTransferOutOpenlyRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForApplyQuickTransferOutOpenly"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForApplyQuickTransferOutOpenlyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 确认转出
//
// @param request - SaveSingleTaskForApprovingTransferOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForApprovingTransferOutResponse
func (client *Client) SaveSingleTaskForApprovingTransferOutWithContext(ctx context.Context, request *SaveSingleTaskForApprovingTransferOutRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForApprovingTransferOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForApprovingTransferOut"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForApprovingTransferOutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveSingleTaskForAssociatingEnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForAssociatingEnsResponse
func (client *Client) SaveSingleTaskForAssociatingEnsWithContext(ctx context.Context, request *SaveSingleTaskForAssociatingEnsRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForAssociatingEnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForAssociatingEns"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForAssociatingEnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveSingleTaskForCancelingTransferInRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForCancelingTransferInResponse
func (client *Client) SaveSingleTaskForCancelingTransferInWithContext(ctx context.Context, request *SaveSingleTaskForCancelingTransferInRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForCancelingTransferInResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForCancelingTransferIn"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForCancelingTransferInResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消转出
//
// @param request - SaveSingleTaskForCancelingTransferOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForCancelingTransferOutResponse
func (client *Client) SaveSingleTaskForCancelingTransferOutWithContext(ctx context.Context, request *SaveSingleTaskForCancelingTransferOutRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForCancelingTransferOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForCancelingTransferOut"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForCancelingTransferOutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存创建dns服务器的任务请求
//
// @param request - SaveSingleTaskForCreatingDnsHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForCreatingDnsHostResponse
func (client *Client) SaveSingleTaskForCreatingDnsHostWithContext(ctx context.Context, request *SaveSingleTaskForCreatingDnsHostRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForCreatingDnsHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DnsName) {
		query["DnsName"] = request.DnsName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForCreatingDnsHost"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForCreatingDnsHostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存单个任务-注册订单
//
// @param request - SaveSingleTaskForCreatingOrderActivateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForCreatingOrderActivateResponse
func (client *Client) SaveSingleTaskForCreatingOrderActivateWithContext(ctx context.Context, request *SaveSingleTaskForCreatingOrderActivateRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForCreatingOrderActivateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AliyunDns) {
		query["AliyunDns"] = request.AliyunDns
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.Dns1) {
		query["Dns1"] = request.Dns1
	}

	if !dara.IsNil(request.Dns2) {
		query["Dns2"] = request.Dns2
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.EnableDomainProxy) {
		query["EnableDomainProxy"] = request.EnableDomainProxy
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PermitPremiumActivation) {
		query["PermitPremiumActivation"] = request.PermitPremiumActivation
	}

	if !dara.IsNil(request.PostalCode) {
		query["PostalCode"] = request.PostalCode
	}

	if !dara.IsNil(request.PromotionNo) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RegistrantName) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !dara.IsNil(request.RegistrantOrganization) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.RegistrantType) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubscriptionDuration) {
		query["SubscriptionDuration"] = request.SubscriptionDuration
	}

	if !dara.IsNil(request.TelArea) {
		query["TelArea"] = request.TelArea
	}

	if !dara.IsNil(request.TelExt) {
		query["TelExt"] = request.TelExt
	}

	if !dara.IsNil(request.Telephone) {
		query["Telephone"] = request.Telephone
	}

	if !dara.IsNil(request.TrademarkDomainActivation) {
		query["TrademarkDomainActivation"] = request.TrademarkDomainActivation
	}

	if !dara.IsNil(request.UseCoupon) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !dara.IsNil(request.UsePromotion) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZhAddress) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !dara.IsNil(request.ZhCity) {
		query["ZhCity"] = request.ZhCity
	}

	if !dara.IsNil(request.ZhProvince) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !dara.IsNil(request.ZhRegistrantName) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !dara.IsNil(request.ZhRegistrantOrganization) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForCreatingOrderActivate"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForCreatingOrderActivateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveSingleTaskForCreatingOrderRedeemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForCreatingOrderRedeemResponse
func (client *Client) SaveSingleTaskForCreatingOrderRedeemWithContext(ctx context.Context, request *SaveSingleTaskForCreatingOrderRedeemRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForCreatingOrderRedeemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.CurrentExpirationDate) {
		query["CurrentExpirationDate"] = request.CurrentExpirationDate
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PromotionNo) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !dara.IsNil(request.UseCoupon) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !dara.IsNil(request.UsePromotion) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForCreatingOrderRedeem"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForCreatingOrderRedeemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存单个任务-续费订单
//
// @param request - SaveSingleTaskForCreatingOrderRenewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForCreatingOrderRenewResponse
func (client *Client) SaveSingleTaskForCreatingOrderRenewWithContext(ctx context.Context, request *SaveSingleTaskForCreatingOrderRenewRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForCreatingOrderRenewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.CurrentExpirationDate) {
		query["CurrentExpirationDate"] = request.CurrentExpirationDate
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PromotionNo) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !dara.IsNil(request.SubscriptionDuration) {
		query["SubscriptionDuration"] = request.SubscriptionDuration
	}

	if !dara.IsNil(request.UseCoupon) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !dara.IsNil(request.UsePromotion) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForCreatingOrderRenew"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForCreatingOrderRenewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveSingleTaskForCreatingOrderTransferRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForCreatingOrderTransferResponse
func (client *Client) SaveSingleTaskForCreatingOrderTransferWithContext(ctx context.Context, request *SaveSingleTaskForCreatingOrderTransferRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForCreatingOrderTransferResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationCode) {
		query["AuthorizationCode"] = request.AuthorizationCode
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PermitPremiumTransfer) {
		query["PermitPremiumTransfer"] = request.PermitPremiumTransfer
	}

	if !dara.IsNil(request.PromotionNo) {
		query["PromotionNo"] = request.PromotionNo
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.UseCoupon) {
		query["UseCoupon"] = request.UseCoupon
	}

	if !dara.IsNil(request.UsePromotion) {
		query["UsePromotion"] = request.UsePromotion
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForCreatingOrderTransfer"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForCreatingOrderTransferResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除dnsSec记录
//
// @param request - SaveSingleTaskForDeletingDSRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForDeletingDSRecordResponse
func (client *Client) SaveSingleTaskForDeletingDSRecordWithContext(ctx context.Context, request *SaveSingleTaskForDeletingDSRecordRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForDeletingDSRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.KeyTag) {
		query["KeyTag"] = request.KeyTag
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForDeletingDSRecord"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForDeletingDSRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除DNS HOST任务
//
// @param request - SaveSingleTaskForDeletingDnsHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForDeletingDnsHostResponse
func (client *Client) SaveSingleTaskForDeletingDnsHostWithContext(ctx context.Context, request *SaveSingleTaskForDeletingDnsHostRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForDeletingDnsHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DnsName) {
		query["DnsName"] = request.DnsName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForDeletingDnsHost"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForDeletingDnsHostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveSingleTaskForDisassociatingEnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForDisassociatingEnsResponse
func (client *Client) SaveSingleTaskForDisassociatingEnsWithContext(ctx context.Context, request *SaveSingleTaskForDisassociatingEnsRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForDisassociatingEnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForDisassociatingEns"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForDisassociatingEnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存单个任务-开启/关闭whois隐私保护锁
//
// @param request - SaveSingleTaskForDomainNameProxyServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForDomainNameProxyServiceResponse
func (client *Client) SaveSingleTaskForDomainNameProxyServiceWithContext(ctx context.Context, request *SaveSingleTaskForDomainNameProxyServiceRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForDomainNameProxyServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForDomainNameProxyService"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForDomainNameProxyServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交生成域名证书任务
//
// @param request - SaveSingleTaskForGenerateDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForGenerateDomainCertificateResponse
func (client *Client) SaveSingleTaskForGenerateDomainCertificateWithContext(ctx context.Context, request *SaveSingleTaskForGenerateDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForGenerateDomainCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForGenerateDomainCertificate"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForGenerateDomainCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改DnsSec记录
//
// @param request - SaveSingleTaskForModifyingDSRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForModifyingDSRecordResponse
func (client *Client) SaveSingleTaskForModifyingDSRecordWithContext(ctx context.Context, request *SaveSingleTaskForModifyingDSRecordRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForModifyingDSRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.Digest) {
		query["Digest"] = request.Digest
	}

	if !dara.IsNil(request.DigestType) {
		query["DigestType"] = request.DigestType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.KeyTag) {
		query["KeyTag"] = request.KeyTag
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForModifyingDSRecord"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForModifyingDSRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存修改dns服务器的任务请求
//
// @param request - SaveSingleTaskForModifyingDnsHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForModifyingDnsHostResponse
func (client *Client) SaveSingleTaskForModifyingDnsHostWithContext(ctx context.Context, request *SaveSingleTaskForModifyingDnsHostRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForModifyingDnsHostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DnsName) {
		query["DnsName"] = request.DnsName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForModifyingDnsHost"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForModifyingDnsHostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送转移码
//
// @param request - SaveSingleTaskForQueryingTransferAuthorizationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForQueryingTransferAuthorizationCodeResponse
func (client *Client) SaveSingleTaskForQueryingTransferAuthorizationCodeWithContext(ctx context.Context, request *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForQueryingTransferAuthorizationCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForQueryingTransferAuthorizationCode"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForQueryingTransferAuthorizationCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 单笔抢注批量接口
//
// @param request - SaveSingleTaskForReserveDropListDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForReserveDropListDomainResponse
func (client *Client) SaveSingleTaskForReserveDropListDomainWithContext(ctx context.Context, request *SaveSingleTaskForReserveDropListDomainRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForReserveDropListDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactTemplateId) {
		query["ContactTemplateId"] = request.ContactTemplateId
	}

	if !dara.IsNil(request.Dns1) {
		query["Dns1"] = request.Dns1
	}

	if !dara.IsNil(request.Dns2) {
		query["Dns2"] = request.Dns2
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForReserveDropListDomain"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForReserveDropListDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存art扩展信息任务
//
// @param request - SaveSingleTaskForSaveArtExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForSaveArtExtensionResponse
func (client *Client) SaveSingleTaskForSaveArtExtensionWithContext(ctx context.Context, request *SaveSingleTaskForSaveArtExtensionRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForSaveArtExtensionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DateOrPeriod) {
		query["DateOrPeriod"] = request.DateOrPeriod
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Features) {
		query["Features"] = request.Features
	}

	if !dara.IsNil(request.InscriptionsAndMarkings) {
		query["InscriptionsAndMarkings"] = request.InscriptionsAndMarkings
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Maker) {
		query["Maker"] = request.Maker
	}

	if !dara.IsNil(request.MaterialsAndTechniques) {
		query["MaterialsAndTechniques"] = request.MaterialsAndTechniques
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.Reference) {
		query["Reference"] = request.Reference
	}

	if !dara.IsNil(request.Subject) {
		query["Subject"] = request.Subject
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForSaveArtExtension"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForSaveArtExtensionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步DnsSec记录
//
// @param request - SaveSingleTaskForSynchronizingDSRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForSynchronizingDSRecordResponse
func (client *Client) SaveSingleTaskForSynchronizingDSRecordWithContext(ctx context.Context, request *SaveSingleTaskForSynchronizingDSRecordRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForSynchronizingDSRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForSynchronizingDSRecord"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForSynchronizingDSRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存同步dns服务器的任务请求
//
// @param request - SaveSingleTaskForSynchronizingDnsHostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForSynchronizingDnsHostResponse
func (client *Client) SaveSingleTaskForSynchronizingDnsHostWithContext(ctx context.Context, request *SaveSingleTaskForSynchronizingDnsHostRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForSynchronizingDnsHostResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForSynchronizingDnsHost"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForSynchronizingDnsHostResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submit a single transfer-out task based on the transfer key of domain names.
//
// Description:
//
// The task ID.
//
// @param request - SaveSingleTaskForTransferOutByAuthorizationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForTransferOutByAuthorizationCodeResponse
func (client *Client) SaveSingleTaskForTransferOutByAuthorizationCodeWithContext(ctx context.Context, request *SaveSingleTaskForTransferOutByAuthorizationCodeRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForTransferOutByAuthorizationCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationCode) {
		query["AuthorizationCode"] = request.AuthorizationCode
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForTransferOutByAuthorizationCode"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForTransferOutByAuthorizationCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存单个任务-开启/关闭禁止转移锁
//
// @param request - SaveSingleTaskForTransferProhibitionLockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForTransferProhibitionLockResponse
func (client *Client) SaveSingleTaskForTransferProhibitionLockWithContext(ctx context.Context, request *SaveSingleTaskForTransferProhibitionLockRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForTransferProhibitionLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForTransferProhibitionLock"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForTransferProhibitionLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存单个任务-开启/关闭信息安全锁
//
// @param request - SaveSingleTaskForUpdateProhibitionLockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForUpdateProhibitionLockResponse
func (client *Client) SaveSingleTaskForUpdateProhibitionLockWithContext(ctx context.Context, request *SaveSingleTaskForUpdateProhibitionLockRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForUpdateProhibitionLockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForUpdateProhibitionLock"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForUpdateProhibitionLockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存修改联系人的任务
//
// @param request - SaveSingleTaskForUpdatingContactInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSingleTaskForUpdatingContactInfoResponse
func (client *Client) SaveSingleTaskForUpdatingContactInfoWithContext(ctx context.Context, request *SaveSingleTaskForUpdatingContactInfoRequest, runtime *dara.RuntimeOptions) (_result *SaveSingleTaskForUpdatingContactInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddTransferLock) {
		query["AddTransferLock"] = request.AddTransferLock
	}

	if !dara.IsNil(request.ContactType) {
		query["ContactType"] = request.ContactType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForUpdatingContactInfo"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSingleTaskForUpdatingContactInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存删除域名的任务
//
// @param request - SaveTaskForSubmittingDomainDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForSubmittingDomainDeleteResponse
func (client *Client) SaveTaskForSubmittingDomainDeleteWithContext(ctx context.Context, request *SaveTaskForSubmittingDomainDeleteRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForSubmittingDomainDeleteResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForSubmittingDomainDelete"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForSubmittingDomainDeleteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量提交域名资料
//
// @param request - SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse
func (client *Client) SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialWithContext(ctx context.Context, request *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.IdentityCredentialNo) {
		query["IdentityCredentialNo"] = request.IdentityCredentialNo
	}

	if !dara.IsNil(request.IdentityCredentialType) {
		query["IdentityCredentialType"] = request.IdentityCredentialType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityCredential) {
		body["IdentityCredential"] = request.IdentityCredential
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredential"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForSubmittingDomainRealNameVerificationByIdentityCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据模板保存域名的实名认证信息
//
// @param request - SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse
func (client *Client) SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDWithContext(ctx context.Context, request *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileID"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForSubmittingDomainRealNameVerificationByRegistrantProfileIDResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据联系人信息批量修改注册联系人信息
//
// @param request - SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse
func (client *Client) SaveTaskForUpdatingRegistrantInfoByIdentityCredentialWithContext(ctx context.Context, request *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.IdentityCredentialNo) {
		query["IdentityCredentialNo"] = request.IdentityCredentialNo
	}

	if !dara.IsNil(request.IdentityCredentialType) {
		query["IdentityCredentialType"] = request.IdentityCredentialType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PostalCode) {
		query["PostalCode"] = request.PostalCode
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RegistrantName) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !dara.IsNil(request.RegistrantOrganization) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !dara.IsNil(request.RegistrantType) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !dara.IsNil(request.TelArea) {
		query["TelArea"] = request.TelArea
	}

	if !dara.IsNil(request.TelExt) {
		query["TelExt"] = request.TelExt
	}

	if !dara.IsNil(request.Telephone) {
		query["Telephone"] = request.Telephone
	}

	if !dara.IsNil(request.TransferOutProhibited) {
		query["TransferOutProhibited"] = request.TransferOutProhibited
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZhAddress) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !dara.IsNil(request.ZhCity) {
		query["ZhCity"] = request.ZhCity
	}

	if !dara.IsNil(request.ZhProvince) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !dara.IsNil(request.ZhRegistrantName) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !dara.IsNil(request.ZhRegistrantOrganization) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityCredential) {
		body["IdentityCredential"] = request.IdentityCredential
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForUpdatingRegistrantInfoByIdentityCredential"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据模板批量修改注册联系人
//
// @param request - SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse
func (client *Client) SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDWithContext(ctx context.Context, request *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest, runtime *dara.RuntimeOptions) (_result *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.TransferOutProhibited) {
		query["TransferOutProhibited"] = request.TransferOutProhibited
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveTaskForUpdatingRegistrantInfoByRegistrantProfileID"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Traverses domain names.
//
// Description:
//
// If you have a large number of domain names, a slow response may occur when you call an API operation to query domain names. In this case, you can call this operation to query domain names more quickly. When you call this operation for the first time, specify the request parameters except ScrollId. A scroll ID is returned without other data. In the second request, use the scroll ID obtained from the previous response. In subsequent requests, the newly specified request parameters do not take effect, and the request parameters that are specified in the first request prevail.
//
// @param request - ScrollDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScrollDomainListResponse
func (client *Client) ScrollDomainListWithContext(ctx context.Context, request *ScrollDomainListRequest, runtime *dara.RuntimeOptions) (_result *ScrollDomainListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainGroupId) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !dara.IsNil(request.EndExpirationDate) {
		query["EndExpirationDate"] = request.EndExpirationDate
	}

	if !dara.IsNil(request.EndLength) {
		query["EndLength"] = request.EndLength
	}

	if !dara.IsNil(request.EndRegistrationDate) {
		query["EndRegistrationDate"] = request.EndRegistrationDate
	}

	if !dara.IsNil(request.Excluded) {
		query["Excluded"] = request.Excluded
	}

	if !dara.IsNil(request.ExcludedPrefix) {
		query["ExcludedPrefix"] = request.ExcludedPrefix
	}

	if !dara.IsNil(request.ExcludedSuffix) {
		query["ExcludedSuffix"] = request.ExcludedSuffix
	}

	if !dara.IsNil(request.Form) {
		query["Form"] = request.Form
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.KeyWordPrefix) {
		query["KeyWordPrefix"] = request.KeyWordPrefix
	}

	if !dara.IsNil(request.KeyWordSuffix) {
		query["KeyWordSuffix"] = request.KeyWordSuffix
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductDomainType) {
		query["ProductDomainType"] = request.ProductDomainType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScrollId) {
		query["ScrollId"] = request.ScrollId
	}

	if !dara.IsNil(request.StartExpirationDate) {
		query["StartExpirationDate"] = request.StartExpirationDate
	}

	if !dara.IsNil(request.StartLength) {
		query["StartLength"] = request.StartLength
	}

	if !dara.IsNil(request.StartRegistrationDate) {
		query["StartRegistrationDate"] = request.StartRegistrationDate
	}

	if !dara.IsNil(request.Suffixs) {
		query["Suffixs"] = request.Suffixs
	}

	if !dara.IsNil(request.TradeType) {
		query["TradeType"] = request.TradeType
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScrollDomainList"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ScrollDomainListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置默认模板
//
// @param request - SetDefaultRegistrantProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultRegistrantProfileResponse
func (client *Client) SetDefaultRegistrantProfileWithContext(ctx context.Context, request *SetDefaultRegistrantProfileRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultRegistrantProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegistrantProfileId) {
		query["RegistrantProfileId"] = request.RegistrantProfileId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultRegistrantProfile"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultRegistrantProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 域名设置自动续费
//
// @param request - SetupDomainAutoRenewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetupDomainAutoRenewResponse
func (client *Client) SetupDomainAutoRenewWithContext(ctx context.Context, request *SetupDomainAutoRenewRequest, runtime *dara.RuntimeOptions) (_result *SetupDomainAutoRenewResponse, _err error) {
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

	if !dara.IsNil(request.Operation) {
		query["Operation"] = request.Operation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetupDomainAutoRenew"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetupDomainAutoRenewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 域名特殊业务提交资料
//
// @param request - SubmitDomainSpecialBizCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDomainSpecialBizCredentialsResponse
func (client *Client) SubmitDomainSpecialBizCredentialsWithContext(ctx context.Context, request *SubmitDomainSpecialBizCredentialsRequest, runtime *dara.RuntimeOptions) (_result *SubmitDomainSpecialBizCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Credentials) {
		body["Credentials"] = request.Credentials
	}

	if !dara.IsNil(request.Extend) {
		body["Extend"] = request.Extend
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDomainSpecialBizCredentials"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDomainSpecialBizCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交邮箱验证
//
// @param request - SubmitEmailVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitEmailVerificationResponse
func (client *Client) SubmitEmailVerificationWithContext(ctx context.Context, request *SubmitEmailVerificationRequest, runtime *dara.RuntimeOptions) (_result *SubmitEmailVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SendIfExist) {
		query["SendIfExist"] = request.SendIfExist
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitEmailVerification"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitEmailVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交申请信息
//
// @param request - SubmitOperationAuditInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitOperationAuditInfoResponse
func (client *Client) SubmitOperationAuditInfoWithContext(ctx context.Context, request *SubmitOperationAuditInfoRequest, runtime *dara.RuntimeOptions) (_result *SubmitOperationAuditInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditInfo) {
		query["AuditInfo"] = request.AuditInfo
	}

	if !dara.IsNil(request.AuditType) {
		query["AuditType"] = request.AuditType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitOperationAuditInfo"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitOperationAuditInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交证件资料
//
// @param request - SubmitOperationCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitOperationCredentialsResponse
func (client *Client) SubmitOperationCredentialsWithContext(ctx context.Context, request *SubmitOperationCredentialsRequest, runtime *dara.RuntimeOptions) (_result *SubmitOperationCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditRecordId) {
		query["AuditRecordId"] = request.AuditRecordId
	}

	if !dara.IsNil(request.AuditType) {
		query["AuditType"] = request.AuditType
	}

	if !dara.IsNil(request.Credentials) {
		query["Credentials"] = request.Credentials
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegType) {
		query["RegType"] = request.RegType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitOperationCredentials"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitOperationCredentialsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TransferInCheckMailTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferInCheckMailTokenResponse
func (client *Client) TransferInCheckMailTokenWithContext(ctx context.Context, request *TransferInCheckMailTokenRequest, runtime *dara.RuntimeOptions) (_result *TransferInCheckMailTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferInCheckMailToken"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferInCheckMailTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TransferInReenterTransferAuthorizationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferInReenterTransferAuthorizationCodeResponse
func (client *Client) TransferInReenterTransferAuthorizationCodeWithContext(ctx context.Context, request *TransferInReenterTransferAuthorizationCodeRequest, runtime *dara.RuntimeOptions) (_result *TransferInReenterTransferAuthorizationCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TransferAuthorizationCode) {
		query["TransferAuthorizationCode"] = request.TransferAuthorizationCode
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferInReenterTransferAuthorizationCode"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferInReenterTransferAuthorizationCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TransferInRefetchWhoisEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferInRefetchWhoisEmailResponse
func (client *Client) TransferInRefetchWhoisEmailWithContext(ctx context.Context, request *TransferInRefetchWhoisEmailRequest, runtime *dara.RuntimeOptions) (_result *TransferInRefetchWhoisEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferInRefetchWhoisEmail"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferInRefetchWhoisEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - TransferInResendMailTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferInResendMailTokenResponse
func (client *Client) TransferInResendMailTokenWithContext(ctx context.Context, request *TransferInResendMailTokenRequest, runtime *dara.RuntimeOptions) (_result *TransferInResendMailTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferInResendMailToken"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferInResendMailTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向分组设置域名
//
// @param request - UpdateDomainToDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainToDomainGroupResponse
func (client *Client) UpdateDomainToDomainGroupWithContext(ctx context.Context, request *UpdateDomainToDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainToDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSource) {
		query["DataSource"] = request.DataSource
	}

	if !dara.IsNil(request.DomainGroupId) {
		query["DomainGroupId"] = request.DomainGroupId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Replace) {
		query["Replace"] = request.Replace
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FileToUpload) {
		body["FileToUpload"] = request.FileToUpload
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainToDomainGroup"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainToDomainGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验联系人信息
//
// @param request - VerifyContactFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyContactFieldResponse
func (client *Client) VerifyContactFieldWithContext(ctx context.Context, request *VerifyContactFieldRequest, runtime *dara.RuntimeOptions) (_result *VerifyContactFieldResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PostalCode) {
		query["PostalCode"] = request.PostalCode
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RegistrantName) {
		query["RegistrantName"] = request.RegistrantName
	}

	if !dara.IsNil(request.RegistrantOrganization) {
		query["RegistrantOrganization"] = request.RegistrantOrganization
	}

	if !dara.IsNil(request.RegistrantType) {
		query["RegistrantType"] = request.RegistrantType
	}

	if !dara.IsNil(request.TelArea) {
		query["TelArea"] = request.TelArea
	}

	if !dara.IsNil(request.TelExt) {
		query["TelExt"] = request.TelExt
	}

	if !dara.IsNil(request.Telephone) {
		query["Telephone"] = request.Telephone
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZhAddress) {
		query["ZhAddress"] = request.ZhAddress
	}

	if !dara.IsNil(request.ZhCity) {
		query["ZhCity"] = request.ZhCity
	}

	if !dara.IsNil(request.ZhProvince) {
		query["ZhProvince"] = request.ZhProvince
	}

	if !dara.IsNil(request.ZhRegistrantName) {
		query["ZhRegistrantName"] = request.ZhRegistrantName
	}

	if !dara.IsNil(request.ZhRegistrantOrganization) {
		query["ZhRegistrantOrganization"] = request.ZhRegistrantOrganization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyContactField"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyContactFieldResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证邮箱Token
//
// @param request - VerifyEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyEmailResponse
func (client *Client) VerifyEmailWithContext(ctx context.Context, request *VerifyEmailRequest, runtime *dara.RuntimeOptions) (_result *VerifyEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyEmail"),
		Version:     dara.String("2018-01-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyEmailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
