// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 身份证三要素
//
// @param request - CertNoThreeElementVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertNoThreeElementVerificationResponse
func (client *Client) CertNoThreeElementVerificationWithContext(ctx context.Context, request *CertNoThreeElementVerificationRequest, runtime *dara.RuntimeOptions) (_result *CertNoThreeElementVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.CertPicture) {
		query["CertPicture"] = request.CertPicture
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CertNoThreeElementVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CertNoThreeElementVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份证二要素认证
//
// @param request - CertNoTwoElementVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertNoTwoElementVerificationResponse
func (client *Client) CertNoTwoElementVerificationWithContext(ctx context.Context, request *CertNoTwoElementVerificationRequest, runtime *dara.RuntimeOptions) (_result *CertNoTwoElementVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertNo) {
		query["CertNo"] = request.CertNo
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CertNoTwoElementVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CertNoTwoElementVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether the enterprise name, the business license number, and the name and ID card of the legal representative belong to the same enterprise. The verification is successful only when the preceding four elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to four-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0, 1, or 2.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyFourElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyFourElementsVerificationResponse
func (client *Client) CompanyFourElementsVerificationWithContext(ctx context.Context, request *CompanyFourElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *CompanyFourElementsVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.EpCertName) {
		query["EpCertName"] = request.EpCertName
	}

	if !dara.IsNil(request.EpCertNo) {
		query["EpCertNo"] = request.EpCertNo
	}

	if !dara.IsNil(request.LegalPersonCertName) {
		query["LegalPersonCertName"] = request.LegalPersonCertName
	}

	if !dara.IsNil(request.LegalPersonCertNo) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompanyFourElementsVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompanyFourElementsVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether the enterprise name, the business license number, and the name of the legal representative belong to the same enterprise. The verification is successful only when the three elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to three-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0, 1, or 2.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyThreeElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyThreeElementsVerificationResponse
func (client *Client) CompanyThreeElementsVerificationWithContext(ctx context.Context, request *CompanyThreeElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *CompanyThreeElementsVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.EpCertName) {
		query["EpCertName"] = request.EpCertName
	}

	if !dara.IsNil(request.EpCertNo) {
		query["EpCertNo"] = request.EpCertNo
	}

	if !dara.IsNil(request.LegalPersonCertName) {
		query["LegalPersonCertName"] = request.LegalPersonCertName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompanyThreeElementsVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompanyThreeElementsVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether the enterprise name and the business license number belong to the same enterprise. The verification is successful only when the two elements belong to the same enterprise and the business status of the enterprise is Active.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the billing of services related to two-element verification for enterprises. For more information, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
//		- You are charged only if the value of VerifyResult is true or false and the value of ReasonCode is 0 or 1.
//
//		- Before you call this operation, perform the following operations: Log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC\\&lang=zh). On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - CompanyTwoElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyTwoElementsVerificationResponse
func (client *Client) CompanyTwoElementsVerificationWithContext(ctx context.Context, request *CompanyTwoElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *CompanyTwoElementsVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.EpCertName) {
		query["EpCertName"] = request.EpCertName
	}

	if !dara.IsNil(request.EpCertNo) {
		query["EpCertNo"] = request.EpCertNo
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompanyTwoElementsVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompanyTwoElementsVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警联系人删除
//
// @param request - DeleteContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactsResponse
func (client *Client) DeleteContactsWithContext(ctx context.Context, request *DeleteContactsRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContacts"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Predicts whether a phone number is a nonexistent number by using AI algorithms.
//
// Description:
//
//	  You can call this operation to verify whether a phone number is a nonexistent number. When you call this operation to verify a number, the system charges you CNY 0.01 per verification based on the number of verifications. **Before you call this operation, make sure that you are familiar with the billing of Cell Phone Number Service.**
//
//		- You are charged only if the value of Code is OK and the value of Status is not UNKNOWN.
//
//		- The prediction is not strictly accurate because Cell Phone Number Service predicts the nonexistent number probability by using AI algorithms. The accuracy rate of the prediction and the recall rate of empty numbers are about 95%. **Pay attention to this point when you call this operation**.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ### [](#)Authorization information
//
// By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribeEmptyNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEmptyNumberResponse
func (client *Client) DescribeEmptyNumberWithContext(ctx context.Context, request *DescribeEmptyNumberRequest, runtime *dara.RuntimeOptions) (_result *DescribeEmptyNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEmptyNumber"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEmptyNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码归属地查询v2
//
// @param request - DescribeMobileOperatorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMobileOperatorAttributeResponse
func (client *Client) DescribeMobileOperatorAttributeWithContext(ctx context.Context, request *DescribeMobileOperatorAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMobileOperatorAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMobileOperatorAttribute"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMobileOperatorAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码分析实时查询蚂蚁
//
// @param request - DescribePhoneNumberAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisResponse
func (client *Client) DescribePhoneNumberAnalysisWithContext(ctx context.Context, request *DescribePhoneNumberAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.NumberType) {
		query["NumberType"] = request.NumberType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Rate) {
		query["Rate"] = request.Rate
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberAnalysis"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a phone number.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the Labels page, find the label that you want to use, click Activate Now, enter the required information, and then submit your application. After your application is approved, you can use the label. Before you call this operation, make sure that you are familiar with the billing of Cell Phone Number Service.
//
// @param request - DescribePhoneNumberAnalysisAIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisAIResponse
func (client *Client) DescribePhoneNumberAnalysisAIWithContext(ctx context.Context, request *DescribePhoneNumberAnalysisAIRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAnalysisAIResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.ModelConfig) {
		query["ModelConfig"] = request.ModelConfig
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Rate) {
		query["Rate"] = request.Rate
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberAnalysisAI"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisAIResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码分析服务pai供应商批量查询接口
//
// @param request - DescribePhoneNumberAnalysisPaiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisPaiResponse
func (client *Client) DescribePhoneNumberAnalysisPaiWithContext(ctx context.Context, request *DescribePhoneNumberAnalysisPaiRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAnalysisPaiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.ModelConfig) {
		query["ModelConfig"] = request.ModelConfig
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Rate) {
		query["Rate"] = request.Rate
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberAnalysisPai"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisPaiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 泛行业人群筛选
//
// @param request - DescribePhoneNumberAnalysisTransparentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisTransparentResponse
func (client *Client) DescribePhoneNumberAnalysisTransparentWithContext(ctx context.Context, request *DescribePhoneNumberAnalysisTransparentRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAnalysisTransparentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NumberType) {
		query["NumberType"] = request.NumberType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberAnalysisTransparent"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberAnalysisTransparentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribePhoneNumberAttribute is deprecated, please use Dytnsapi::2020-02-17::DescribePhoneNumberOperatorAttribute instead.
//
// Summary:
//
// Queries the carrier, registration location, and mobile number portability information of a phone number.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 2,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePhoneNumberAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAttributeResponse
func (client *Client) DescribePhoneNumberAttributeWithContext(ctx context.Context, request *DescribePhoneNumberAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberAttribute"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage period of a phone number of a user.
//
// Description:
//
//	  Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePhoneNumberOnlineTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOnlineTimeResponse
func (client *Client) DescribePhoneNumberOnlineTimeWithContext(ctx context.Context, request *DescribePhoneNumberOnlineTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberOnlineTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Carrier) {
		query["Carrier"] = request.Carrier
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberOnlineTime"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberOnlineTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attribute information about a phone number, including the registration province, registration city, basic carrier (such as China Mobile, China Unicom, China Telecom, or China Broadnet), reseller of mobile communications services (such as Alibaba Communications), mobile number portability, and the number segment to which the phone number belongs.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154008.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- You can call this operation to obtain the carrier, registration location, and mobile number portability information about a phone number. You can query phone numbers in **plaintext*	- and phone numbers that are encrypted by using **MD5*	- and **SHA256**.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// @param request - DescribePhoneNumberOperatorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOperatorAttributeResponse
func (client *Client) DescribePhoneNumberOperatorAttributeWithContext(ctx context.Context, request *DescribePhoneNumberOperatorAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResultCount) {
		query["ResultCount"] = request.ResultCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberOperatorAttribute"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberOperatorAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码归属服务(包年包月客户专用)
//
// @param request - DescribePhoneNumberOperatorAttributeAnnualRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOperatorAttributeAnnualResponse
func (client *Client) DescribePhoneNumberOperatorAttributeAnnualWithContext(ctx context.Context, request *DescribePhoneNumberOperatorAttributeAnnualRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeAnnualResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberOperatorAttributeAnnual"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberOperatorAttributeAnnualResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 号码归属服务(包年包月客户专用)
//
// @param request - DescribePhoneNumberOperatorAttributeAnnualUseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOperatorAttributeAnnualUseResponse
func (client *Client) DescribePhoneNumberOperatorAttributeAnnualUseWithContext(ctx context.Context, request *DescribePhoneNumberOperatorAttributeAnnualUseRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeAnnualUseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberOperatorAttributeAnnualUse"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberOperatorAttributeAnnualUseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 风险用户评分
//
// @param request - DescribePhoneNumberRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberRiskResponse
func (client *Client) DescribePhoneNumberRiskWithContext(ctx context.Context, request *DescribePhoneNumberRiskRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberRiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneNumberRisk"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneNumberRiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether a phone number is a reassigned phone number by calling this operation.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- You are charged for phone number verifications only if the value of Code is OK and the value of VerifyResult is not 0.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ## [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// ## [](#)Authorization information
//
// By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribePhoneTwiceTelVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneTwiceTelVerifyResponse
func (client *Client) DescribePhoneTwiceTelVerifyWithContext(ctx context.Context, request *DescribePhoneTwiceTelVerifyRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneTwiceTelVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePhoneTwiceTelVerify"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePhoneTwiceTelVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取UAID申请Token所需的签名字段
//
// @param request - GetUAIDApplyTokenSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUAIDApplyTokenSignResponse
func (client *Client) GetUAIDApplyTokenSignWithContext(ctx context.Context, request *GetUAIDApplyTokenSignRequest, runtime *dara.RuntimeOptions) (_result *GetUAIDApplyTokenSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Carrier) {
		query["Carrier"] = request.Carrier
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Format) {
		query["Format"] = request.Format
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParamKey) {
		query["ParamKey"] = request.ParamKey
	}

	if !dara.IsNil(request.ParamStr) {
		query["ParamStr"] = request.ParamStr
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUAIDApplyTokenSign"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUAIDApplyTokenSignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取号码采集服务申请Token所需的签名字段
//
// @param request - GetUAIDConversionSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUAIDConversionSignResponse
func (client *Client) GetUAIDConversionSignWithContext(ctx context.Context, request *GetUAIDConversionSignRequest, runtime *dara.RuntimeOptions) (_result *GetUAIDConversionSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Carrier) {
		query["Carrier"] = request.Carrier
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Format) {
		query["Format"] = request.Format
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParamKey) {
		query["ParamKey"] = request.ParamKey
	}

	if !dara.IsNil(request.ParamStr) {
		query["ParamStr"] = request.ParamStr
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUAIDConversionSign"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUAIDConversionSignResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Filters invalid phone numbers.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - InvalidPhoneNumberFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidPhoneNumberFilterResponse
func (client *Client) InvalidPhoneNumberFilterWithContext(ctx context.Context, request *InvalidPhoneNumberFilterRequest, runtime *dara.RuntimeOptions) (_result *InvalidPhoneNumberFilterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvalidPhoneNumberFilter"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvalidPhoneNumberFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警联系人查询
//
// @param request - ListContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactsResponse
func (client *Client) ListContactsWithContext(ctx context.Context, request *ListContactsRequest, runtime *dara.RuntimeOptions) (_result *ListContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListContacts"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PhoneNumberConvertServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberConvertServiceResponse
func (client *Client) PhoneNumberConvertServiceWithContext(ctx context.Context, request *PhoneNumberConvertServiceRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberConvertServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PhoneNumberConvertService"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PhoneNumberConvertServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Encrypts the original phone number into a virtual number that starts with 140. Cell Phone Number Service integrates the communications services provided by Alibaba Cloud. This allows you to initiate a call by using a virtual number that starts with 140.
//
// Description:
//
// Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberEncryptResponse
func (client *Client) PhoneNumberEncryptWithContext(ctx context.Context, request *PhoneNumberEncryptRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberEncryptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PhoneNumberEncrypt"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PhoneNumberEncryptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForAccountResponse
func (client *Client) PhoneNumberStatusForAccountWithContext(ctx context.Context, request *PhoneNumberStatusForAccountRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PhoneNumberStatusForAccount"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PhoneNumberStatusForAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForPublicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForPublicResponse
func (client *Client) PhoneNumberStatusForPublicWithContext(ctx context.Context, request *PhoneNumberStatusForPublicRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForPublicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PhoneNumberStatusForPublic"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PhoneNumberStatusForPublicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForRealRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForRealResponse
func (client *Client) PhoneNumberStatusForRealWithContext(ctx context.Context, request *PhoneNumberStatusForRealRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForRealResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PhoneNumberStatusForReal"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PhoneNumberStatusForRealResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForSmsResponse
func (client *Client) PhoneNumberStatusForSmsWithContext(ctx context.Context, request *PhoneNumberStatusForSmsRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForSmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PhoneNumberStatusForSms"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PhoneNumberStatusForSmsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time service state of a phone number. The state includes NORMAL, SHUTDOWN, and NOT_EXIST. You can choose an encryption method for your phone number query, including plaintext, MD5, and SHA256.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- By default, only Alibaba Cloud accounts can call this operation. RAM users can call this operation only after the RAM users are granted the related permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 300 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PhoneNumberStatusForVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForVoiceResponse
func (client *Client) PhoneNumberStatusForVoiceWithContext(ctx context.Context, request *PhoneNumberStatusForVoiceRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PhoneNumberStatusForVoice"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PhoneNumberStatusForVoiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available authorization codes.
//
// @param request - QueryAvailableAuthCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvailableAuthCodeResponse
func (client *Client) QueryAvailableAuthCodeWithContext(ctx context.Context, request *QueryAvailableAuthCodeRequest, runtime *dara.RuntimeOptions) (_result *QueryAvailableAuthCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAvailableAuthCode"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvailableAuthCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 套餐包类型信息查询
//
// @param request - QueryPackageTypeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPackageTypeInfoResponse
func (client *Client) QueryPackageTypeInfoWithContext(ctx context.Context, request *QueryPackageTypeInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryPackageTypeInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPackageTypeInfo"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPackageTypeInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 在网时长专用接口
//
// @param request - QueryPhoneNumberOnlineTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneNumberOnlineTimeResponse
func (client *Client) QueryPhoneNumberOnlineTimeWithContext(ctx context.Context, request *QueryPhoneNumberOnlineTimeRequest, runtime *dara.RuntimeOptions) (_result *QueryPhoneNumberOnlineTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPhoneNumberOnlineTime"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPhoneNumberOnlineTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 二次号携号转网号码查询
//
// @param request - QueryPhoneTwiceTelVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPhoneTwiceTelVerifyResponse
func (client *Client) QueryPhoneTwiceTelVerifyWithContext(ctx context.Context, request *QueryPhoneTwiceTelVerifyRequest, runtime *dara.RuntimeOptions) (_result *QueryPhoneTwiceTelVerifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPhoneTwiceTelVerify"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPhoneTwiceTelVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag application rules.
//
// @param request - QueryTagApplyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagApplyRuleResponse
func (client *Client) QueryTagApplyRuleWithContext(ctx context.Context, request *QueryTagApplyRuleRequest, runtime *dara.RuntimeOptions) (_result *QueryTagApplyRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTagApplyRule"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTagApplyRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about tags.
//
// @param request - QueryTagInfoBySelectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagInfoBySelectionResponse
func (client *Client) QueryTagInfoBySelectionWithContext(ctx context.Context, request *QueryTagInfoBySelectionRequest, runtime *dara.RuntimeOptions) (_result *QueryTagInfoBySelectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndustryId) {
		query["IndustryId"] = request.IndustryId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTagInfoBySelection"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTagInfoBySelectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tags by page.
//
// @param request - QueryTagListPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagListPageResponse
func (client *Client) QueryTagListPageWithContext(ctx context.Context, request *QueryTagListPageRequest, runtime *dara.RuntimeOptions) (_result *QueryTagListPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTagListPage"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTagListPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询离线任务列表
//
// @param tmpReq - QueryTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskListResponse
func (client *Client) QueryTaskListWithContext(ctx context.Context, tmpReq *QueryTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryTaskListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Result) {
		request.ResultShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Result, dara.String("Result"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskType) {
		request.TaskTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskType, dara.String("TaskType"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResultShrink) {
		query["Result"] = request.ResultShrink
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskTypeShrink) {
		query["TaskType"] = request.TaskTypeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskList"),
		Version:     dara.String("2020-02-17"),
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

// Summary:
//
// Queries the usage statistics based on tag IDs.
//
// @param request - QueryUsageStatisticsByTagIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUsageStatisticsByTagIdResponse
func (client *Client) QueryUsageStatisticsByTagIdWithContext(ctx context.Context, request *QueryUsageStatisticsByTagIdRequest, runtime *dara.RuntimeOptions) (_result *QueryUsageStatisticsByTagIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUsageStatisticsByTagId"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUsageStatisticsByTagIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether the name, phone number, and ID card number entered by a user belong to the same user.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- You are charged only if the value of Code is OK and the value of IsConsistent is not 2.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ThreeElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ThreeElementsVerificationResponse
func (client *Client) ThreeElementsVerificationWithContext(ctx context.Context, request *ThreeElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *ThreeElementsVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.CertCode) {
		query["CertCode"] = request.CertCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ThreeElementsVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ThreeElementsVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether the name and phone number entered by a user belong to the same user.
//
// Description:
//
//	  Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
//		- Before you call this operation, perform the following operations: Log on to the Cell Phone Number Service console. On the [Labels](https://dytns.console.aliyun.com/analysis/square) page, find the label that you want to use, click **Activate Now**, enter the required information, and then submit your application. After your application is approved, you can use the label.
//
//		- You are charged only if the value of Code is OK and the value of IsConsistent is not 2.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 200 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TwoElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TwoElementsVerificationResponse
func (client *Client) TwoElementsVerificationWithContext(ctx context.Context, request *TwoElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *TwoElementsVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.InputNumber) {
		query["InputNumber"] = request.InputNumber
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TwoElementsVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TwoElementsVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UAID采集
//
// @param request - UAIDCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UAIDCollectionResponse
func (client *Client) UAIDCollectionWithContext(ctx context.Context, request *UAIDCollectionRequest, runtime *dara.RuntimeOptions) (_result *UAIDCollectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Carrier) {
		query["Carrier"] = request.Carrier
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserGrantId) {
		query["UserGrantId"] = request.UserGrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UAIDCollection"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UAIDCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// uaid号码转换服务
//
// @param request - UAIDConversionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UAIDConversionResponse
func (client *Client) UAIDConversionWithContext(ctx context.Context, request *UAIDConversionRequest, runtime *dara.RuntimeOptions) (_result *UAIDConversionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Carrier) {
		query["Carrier"] = request.Carrier
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UaidList) {
		query["UaidList"] = request.UaidList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UAIDConversion"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UAIDConversionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取UAID
//
// @param request - UAIDVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UAIDVerificationResponse
func (client *Client) UAIDVerificationWithContext(ctx context.Context, request *UAIDVerificationRequest, runtime *dara.RuntimeOptions) (_result *UAIDVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Carrier) {
		query["Carrier"] = request.Carrier
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserGrantId) {
		query["UserGrantId"] = request.UserGrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UAIDVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UAIDVerificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预警联系人更新
//
// @param request - UpdateContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContactsResponse
func (client *Client) UpdateContactsWithContext(ctx context.Context, request *UpdateContactsRequest, runtime *dara.RuntimeOptions) (_result *UpdateContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactEmail) {
		query["ContactEmail"] = request.ContactEmail
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.ContactPhone) {
		query["ContactPhone"] = request.ContactPhone
	}

	if !dara.IsNil(request.MailStatus) {
		query["MailStatus"] = request.MailStatus
	}

	if !dara.IsNil(request.OpenStatusWarning) {
		query["OpenStatusWarning"] = request.OpenStatusWarning
	}

	if !dara.IsNil(request.OpentAttributionWarning) {
		query["OpentAttributionWarning"] = request.OpentAttributionWarning
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneStatus) {
		query["PhoneStatus"] = request.PhoneStatus
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContacts"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
