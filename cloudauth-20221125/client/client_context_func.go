// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # Enterprise Element Verification
//
// Description:
//
// Supports only enterprises and individual businesses.
//
// @param request - EntElementVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntElementVerifyResponse
func (client *Client) EntElementVerifyWithContext(ctx context.Context, request *EntElementVerifyRequest, runtime *dara.RuntimeOptions) (_result *EntElementVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntName) {
		query["EntName"] = request.EntName
	}

	if !dara.IsNil(request.InfoVerifyType) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !dara.IsNil(request.LegalPersonCertNo) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !dara.IsNil(request.LegalPersonName) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !dara.IsNil(request.LicenseNo) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.UserAuthorization) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EntElementVerify"),
		Version:     dara.String("2022-11-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EntElementVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 企业要素核验PRO
//
// @param request - EntElementVerifyPRORequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntElementVerifyPROResponse
func (client *Client) EntElementVerifyPROWithContext(ctx context.Context, request *EntElementVerifyPRORequest, runtime *dara.RuntimeOptions) (_result *EntElementVerifyPROResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntName) {
		query["EntName"] = request.EntName
	}

	if !dara.IsNil(request.InfoVerifyType) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !dara.IsNil(request.LegalPersonCertNo) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !dara.IsNil(request.LegalPersonName) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !dara.IsNil(request.LicenseNo) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.UserAuthorization) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EntElementVerifyPRO"),
		Version:     dara.String("2022-11-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EntElementVerifyPROResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enterprise Element Verification V2
//
// Description:
//
// The Enterprise Element Verification API provides a service for verifying the consistency of enterprise element information, used to identify the authenticity of enterprise information.
//
// It supports various institutions including enterprises, individual businesses, farmers\\" professional cooperatives, government agencies, public institutions, social organizations, legal profession institutions, and owners\\" meetings for 2-3 elements;
//
// For 4 elements, it supports enterprises, individual businesses, farmers\\" professional cooperatives, and legal professions.
//
// @param request - EntElementVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntElementVerifyV2Response
func (client *Client) EntElementVerifyV2WithContext(ctx context.Context, request *EntElementVerifyV2Request, runtime *dara.RuntimeOptions) (_result *EntElementVerifyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntName) {
		query["EntName"] = request.EntName
	}

	if !dara.IsNil(request.InfoVerifyType) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !dara.IsNil(request.LegalPersonCertNo) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !dara.IsNil(request.LegalPersonName) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !dara.IsNil(request.LicenseNo) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.UserAuthorization) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EntElementVerifyV2"),
		Version:     dara.String("2022-11-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EntElementVerifyV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Abnormal Business Operation Query
//
// @param request - EntRiskQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntRiskQueryResponse
func (client *Client) EntRiskQueryWithContext(ctx context.Context, request *EntRiskQueryRequest, runtime *dara.RuntimeOptions) (_result *EntRiskQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.ParamValue) {
		query["ParamValue"] = request.ParamValue
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.UserAuthorization) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EntRiskQuery"),
		Version:     dara.String("2022-11-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EntRiskQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enterprise Authentication
//
// @param request - EntVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntVerifyResponse
func (client *Client) EntVerifyWithContext(ctx context.Context, request *EntVerifyRequest, runtime *dara.RuntimeOptions) (_result *EntVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountNo) {
		query["AccountNo"] = request.AccountNo
	}

	if !dara.IsNil(request.EntName) {
		query["EntName"] = request.EntName
	}

	if !dara.IsNil(request.InfoVerifyType) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !dara.IsNil(request.LegalPersonCertNo) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !dara.IsNil(request.LegalPersonMobile) {
		query["LegalPersonMobile"] = request.LegalPersonMobile
	}

	if !dara.IsNil(request.LegalPersonName) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !dara.IsNil(request.LicenseNo) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.RiskModelVersion) {
		query["RiskModelVersion"] = request.RiskModelVersion
	}

	if !dara.IsNil(request.RiskVerifyType) {
		query["RiskVerifyType"] = request.RiskVerifyType
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.UserAuthorization) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EntVerify"),
		Version:     dara.String("2022-11-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EntVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
