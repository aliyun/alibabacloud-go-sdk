// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDomain"),
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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

// Summary:
//
// Queries a list of domain names.
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

	if !dara.IsNil(request.StartExpirationDate) {
		query["StartExpirationDate"] = request.StartExpirationDate
	}

	if !dara.IsNil(request.StartRegistrationDate) {
		query["StartRegistrationDate"] = request.StartRegistrationDate
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainList"),
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRegistrantProfiles"),
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForDomainNameProxyService"),
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
// 国际站删除抢注批量接口
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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

// @param request - SaveBatchTaskForUpdatingContactInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBatchTaskForUpdatingContactInfoResponse
func (client *Client) SaveBatchTaskForUpdatingContactInfoWithContext(ctx context.Context, request *SaveBatchTaskForUpdatingContactInfoRequest, runtime *dara.RuntimeOptions) (_result *SaveBatchTaskForUpdatingContactInfoResponse, _err error) {
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
		Action:      dara.String("SaveBatchTaskForUpdatingContactInfo"),
		Version:     dara.String("2017-12-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBatchTaskForUpdatingContactInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBatchTaskForUpdatingContactInfoByNewContact"),
		Version:     dara.String("2017-12-18"),
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveRegistrantProfile"),
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSingleTaskForCreatingOrderActivate"),
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Action:      dara.String("SaveSingleTaskForDeletingDnsHost"),
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyContactField"),
		Version:     dara.String("2017-12-18"),
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
		Version:     dara.String("2017-12-18"),
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
