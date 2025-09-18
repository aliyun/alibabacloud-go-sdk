// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 地址相似比对
//
// @param request - AddressCompareIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddressCompareIntlResponse
func (client *Client) AddressCompareIntlWithContext(ctx context.Context, request *AddressCompareIntlRequest, runtime *dara.RuntimeOptions) (_result *AddressCompareIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultCountry) {
		query["DefaultCountry"] = request.DefaultCountry
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Text1) {
		query["Text1"] = request.Text1
	}

	if !dara.IsNil(request.Text2) {
		query["Text2"] = request.Text2
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddressCompareIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddressCompareIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Address Verification
//
// Description:
//
// Based on the operator\\"s capabilities, input the phone number and address (or latitude and longitude) to verify whether the provided address is the user\\"s usual residence.
//
// @param request - AddressVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddressVerifyIntlResponse
func (client *Client) AddressVerifyIntlWithContext(ctx context.Context, request *AddressVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *AddressVerifyIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.DefaultCity) {
		query["DefaultCity"] = request.DefaultCity
	}

	if !dara.IsNil(request.DefaultCountry) {
		query["DefaultCountry"] = request.DefaultCountry
	}

	if !dara.IsNil(request.DefaultDistrict) {
		query["DefaultDistrict"] = request.DefaultDistrict
	}

	if !dara.IsNil(request.DefaultProvince) {
		query["DefaultProvince"] = request.DefaultProvince
	}

	if !dara.IsNil(request.Latitude) {
		query["Latitude"] = request.Latitude
	}

	if !dara.IsNil(request.Longitude) {
		query["Longitude"] = request.Longitude
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddressVerifyIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddressVerifyIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 地址核验
//
// @param request - AddressVerifyV2IntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddressVerifyV2IntlResponse
func (client *Client) AddressVerifyV2IntlWithContext(ctx context.Context, request *AddressVerifyV2IntlRequest, runtime *dara.RuntimeOptions) (_result *AddressVerifyV2IntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceToken) {
		query["DeviceToken"] = request.DeviceToken
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RegCountry) {
		query["RegCountry"] = request.RegCountry
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddressVerifyV2Intl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddressVerifyV2IntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Bank Card Verification
//
// Description:
//
// Verification of bank card elements, including: two-element verification (name + bank card number), three-element verification (name + ID number + bank card number), and four-element verification (name + ID number + phone number + bank card number) for consistency.
//
// @param request - BankMetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BankMetaVerifyIntlResponse
func (client *Client) BankMetaVerifyIntlWithContext(ctx context.Context, request *BankMetaVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *BankMetaVerifyIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BankCard) {
		query["BankCard"] = request.BankCard
	}

	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VerifyMode) {
		query["VerifyMode"] = request.VerifyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BankMetaVerifyIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BankMetaVerifyIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CardOcr is deprecated, please use Cloudauth-intl::2022-08-09::DocOcr instead.
//
// Summary:
//
// # Pure server-side interface for document OCR recognition
//
// @param request - CardOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CardOcrResponse
func (client *Client) CardOcrWithContext(ctx context.Context, request *CardOcrRequest, runtime *dara.RuntimeOptions) (_result *CardOcrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocType) {
		query["DocType"] = request.DocType
	}

	if !dara.IsNil(request.IdFaceQuality) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !dara.IsNil(request.IdOcrPictureUrl) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.Ocr) {
		query["Ocr"] = request.Ocr
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Spoof) {
		query["Spoof"] = request.Spoof
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdOcrPictureBase64) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CardOcr"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CardOcrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Result Query
//
// @param request - CheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResultResponse
func (client *Client) CheckResultWithContext(ctx context.Context, request *CheckResultRequest, runtime *dara.RuntimeOptions) (_result *CheckResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtraImageControlList) {
		query["ExtraImageControlList"] = request.ExtraImageControlList
	}

	if !dara.IsNil(request.IsReturnImage) {
		query["IsReturnImage"] = request.IsReturnImage
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.ReturnFiveCategorySpoofResult) {
		query["ReturnFiveCategorySpoofResult"] = request.ReturnFiveCategorySpoofResult
	}

	if !dara.IsNil(request.TransactionId) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckResult"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Authentication Log Query Interface
//
// @param request - CheckVerifyLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckVerifyLogResponse
func (client *Client) CheckVerifyLogWithContext(ctx context.Context, request *CheckVerifyLogRequest, runtime *dara.RuntimeOptions) (_result *CheckVerifyLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MerchantBizId) {
		body["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.TransactionId) {
		body["TransactionId"] = request.TransactionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckVerifyLog"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckVerifyLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 凭证识别
//
// @param request - CredentialRecognitionIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialRecognitionIntlResponse
func (client *Client) CredentialRecognitionIntlWithContext(ctx context.Context, request *CredentialRecognitionIntlRequest, runtime *dara.RuntimeOptions) (_result *CredentialRecognitionIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocType) {
		query["DocType"] = request.DocType
	}

	if !dara.IsNil(request.FraudCheck) {
		query["FraudCheck"] = request.FraudCheck
	}

	if !dara.IsNil(request.OcrArea) {
		query["OcrArea"] = request.OcrArea
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CredentialOcrPictureBase64) {
		body["CredentialOcrPictureBase64"] = request.CredentialOcrPictureBase64
	}

	if !dara.IsNil(request.CredentialOcrPictureUrl) {
		body["CredentialOcrPictureUrl"] = request.CredentialOcrPictureUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialRecognitionIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialRecognitionIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Credential Verification
//
// Description:
//
// Input credential image information, perform image quality, tampering, and forgery detection, and return the detection results.
//
// @param request - CredentialVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialVerifyIntlResponse
func (client *Client) CredentialVerifyIntlWithContext(ctx context.Context, request *CredentialVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *CredentialVerifyIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredName) {
		query["CredName"] = request.CredName
	}

	if !dara.IsNil(request.CredType) {
		query["CredType"] = request.CredType
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageFile) {
		body["ImageFile"] = request.ImageFile
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialVerifyIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialVerifyIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸凭证核验
//
// @param request - DeepfakeDetectIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepfakeDetectIntlResponse
func (client *Client) DeepfakeDetectIntlWithContext(ctx context.Context, request *DeepfakeDetectIntlRequest, runtime *dara.RuntimeOptions) (_result *DeepfakeDetectIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaceInputType) {
		query["FaceInputType"] = request.FaceInputType
	}

	if !dara.IsNil(request.FaceUrl) {
		query["FaceUrl"] = request.FaceUrl
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FaceBase64) {
		body["FaceBase64"] = request.FaceBase64
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeepfakeDetectIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeepfakeDetectIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete user authentication record results
//
// @param request - DeleteVerifyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVerifyResultResponse
func (client *Client) DeleteVerifyResultWithContext(ctx context.Context, request *DeleteVerifyResultRequest, runtime *dara.RuntimeOptions) (_result *DeleteVerifyResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteAfterQuery) {
		query["DeleteAfterQuery"] = request.DeleteAfterQuery
	}

	if !dara.IsNil(request.DeleteType) {
		query["DeleteType"] = request.DeleteType
	}

	if !dara.IsNil(request.TransactionId) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVerifyResult"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVerifyResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Card and document OCR pure server-side
//
// @param request - DocOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocOcrResponse
func (client *Client) DocOcrWithContext(ctx context.Context, request *DocOcrRequest, runtime *dara.RuntimeOptions) (_result *DocOcrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CardSide) {
		query["CardSide"] = request.CardSide
	}

	if !dara.IsNil(request.DocType) {
		query["DocType"] = request.DocType
	}

	if !dara.IsNil(request.IdFaceQuality) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !dara.IsNil(request.IdOcrPictureUrl) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !dara.IsNil(request.IdThreshold) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.Ocr) {
		query["Ocr"] = request.Ocr
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Spoof) {
		query["Spoof"] = request.Spoof
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdOcrPictureBase64) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DocOcr"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DocOcrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全球证件ocr识别接口
//
// @param request - DocOcrMaxRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocOcrMaxResponse
func (client *Client) DocOcrMaxWithContext(ctx context.Context, request *DocOcrMaxRequest, runtime *dara.RuntimeOptions) (_result *DocOcrMaxResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocPage) {
		body["DocPage"] = request.DocPage
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.IdOcrPictureBase64) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	if !dara.IsNil(request.IdOcrPictureUrl) {
		body["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !dara.IsNil(request.IdSpoof) {
		body["IdSpoof"] = request.IdSpoof
	}

	if !dara.IsNil(request.IdThreshold) {
		body["IdThreshold"] = request.IdThreshold
	}

	if !dara.IsNil(request.MerchantBizId) {
		body["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		body["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.OcrModel) {
		body["OcrModel"] = request.OcrModel
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SceneCode) {
		body["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.Spoof) {
		body["Spoof"] = request.Spoof
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DocOcrMax"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DocOcrMaxResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ekyc纯服务端接口
//
// @param request - EkycVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EkycVerifyResponse
func (client *Client) EkycVerifyWithContext(ctx context.Context, request *EkycVerifyRequest, runtime *dara.RuntimeOptions) (_result *EkycVerifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Authorize) {
		query["Authorize"] = request.Authorize
	}

	if !dara.IsNil(request.Crop) {
		query["Crop"] = request.Crop
	}

	if !dara.IsNil(request.DocName) {
		query["DocName"] = request.DocName
	}

	if !dara.IsNil(request.DocNo) {
		query["DocNo"] = request.DocNo
	}

	if !dara.IsNil(request.DocType) {
		query["DocType"] = request.DocType
	}

	if !dara.IsNil(request.FacePictureUrl) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !dara.IsNil(request.IdOcrPictureUrl) {
		query["IdOcrPictureUrl"] = request.IdOcrPictureUrl
	}

	if !dara.IsNil(request.IdThreshold) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FacePictureBase64) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	if !dara.IsNil(request.IdOcrPictureBase64) {
		body["IdOcrPictureBase64"] = request.IdOcrPictureBase64
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EkycVerify"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EkycVerifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸比对
//
// @param request - FaceCompareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceCompareResponse
func (client *Client) FaceCompareWithContext(ctx context.Context, request *FaceCompareRequest, runtime *dara.RuntimeOptions) (_result *FaceCompareResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FacePictureQualityCheck) {
		query["FacePictureQualityCheck"] = request.FacePictureQualityCheck
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.SourceFacePictureUrl) {
		query["SourceFacePictureUrl"] = request.SourceFacePictureUrl
	}

	if !dara.IsNil(request.TargetFacePictureUrl) {
		query["TargetFacePictureUrl"] = request.TargetFacePictureUrl
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceFacePicture) {
		body["SourceFacePicture"] = request.SourceFacePicture
	}

	if !dara.IsNil(request.TargetFacePicture) {
		body["TargetFacePicture"] = request.TargetFacePicture
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FaceCompare"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FaceCompareResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际人脸保镖纯服务端接口
//
// @param request - FaceGuardRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceGuardRiskResponse
func (client *Client) FaceGuardRiskWithContext(ctx context.Context, request *FaceGuardRiskRequest, runtime *dara.RuntimeOptions) (_result *FaceGuardRiskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.DeviceToken) {
		query["DeviceToken"] = request.DeviceToken
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FaceGuardRisk"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FaceGuardRiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 静默活体API 纯服务端
//
// @param request - FaceLivenessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceLivenessResponse
func (client *Client) FaceLivenessWithContext(ctx context.Context, request *FaceLivenessRequest, runtime *dara.RuntimeOptions) (_result *FaceLivenessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Crop) {
		query["Crop"] = request.Crop
	}

	if !dara.IsNil(request.FacePictureUrl) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !dara.IsNil(request.FaceQuality) {
		query["FaceQuality"] = request.FaceQuality
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.Occlusion) {
		query["Occlusion"] = request.Occlusion
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FacePictureBase64) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FaceLiveness"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FaceLivenessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 防伪回调接口
//
// @param request - FraudResultCallBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FraudResultCallBackResponse
func (client *Client) FraudResultCallBackWithContext(ctx context.Context, request *FraudResultCallBackRequest, runtime *dara.RuntimeOptions) (_result *FraudResultCallBackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyId) {
		query["CertifyId"] = request.CertifyId
	}

	if !dara.IsNil(request.ExtParams) {
		query["ExtParams"] = request.ExtParams
	}

	if !dara.IsNil(request.ResultCode) {
		query["ResultCode"] = request.ResultCode
	}

	if !dara.IsNil(request.VerifyDeployEnv) {
		query["VerifyDeployEnv"] = request.VerifyDeployEnv
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FraudResultCallBack"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FraudResultCallBackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份二要素有效期核验
//
// @param request - Id2MetaPeriodVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaPeriodVerifyIntlResponse
func (client *Client) Id2MetaPeriodVerifyIntlWithContext(ctx context.Context, request *Id2MetaPeriodVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaPeriodVerifyIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocName) {
		body["DocName"] = request.DocName
	}

	if !dara.IsNil(request.DocNo) {
		body["DocNo"] = request.DocNo
	}

	if !dara.IsNil(request.DocType) {
		body["DocType"] = request.DocType
	}

	if !dara.IsNil(request.MerchantBizId) {
		body["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		body["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneCode) {
		body["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.ValidityEndDate) {
		body["ValidityEndDate"] = request.ValidityEndDate
	}

	if !dara.IsNil(request.ValidityStartDate) {
		body["ValidityStartDate"] = request.ValidityStartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaPeriodVerifyIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaPeriodVerifyIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 身份二要素国际版接口
//
// @param request - Id2MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyIntlResponse
func (client *Client) Id2MetaVerifyIntlWithContext(ctx context.Context, request *Id2MetaVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaVerifyIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Id2MetaVerifyIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Id2MetaVerifyIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 认证初始化
//
// @param tmpReq - InitializeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeResponse
func (client *Client) InitializeWithContext(ctx context.Context, tmpReq *InitializeRequest, runtime *dara.RuntimeOptions) (_result *InitializeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &InitializeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DocPageConfig) {
		request.DocPageConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocPageConfig, dara.String("DocPageConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppQualityCheck) {
		query["AppQualityCheck"] = request.AppQualityCheck
	}

	if !dara.IsNil(request.Authorize) {
		query["Authorize"] = request.Authorize
	}

	if !dara.IsNil(request.CallbackToken) {
		query["CallbackToken"] = request.CallbackToken
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ChameleonFrameEnable) {
		query["ChameleonFrameEnable"] = request.ChameleonFrameEnable
	}

	if !dara.IsNil(request.Crop) {
		query["Crop"] = request.Crop
	}

	if !dara.IsNil(request.DateOfBirth) {
		query["DateOfBirth"] = request.DateOfBirth
	}

	if !dara.IsNil(request.DateOfExpiry) {
		query["DateOfExpiry"] = request.DateOfExpiry
	}

	if !dara.IsNil(request.DocName) {
		query["DocName"] = request.DocName
	}

	if !dara.IsNil(request.DocNo) {
		query["DocNo"] = request.DocNo
	}

	if !dara.IsNil(request.DocPageConfigShrink) {
		query["DocPageConfig"] = request.DocPageConfigShrink
	}

	if !dara.IsNil(request.DocScanMode) {
		query["DocScanMode"] = request.DocScanMode
	}

	if !dara.IsNil(request.DocType) {
		query["DocType"] = request.DocType
	}

	if !dara.IsNil(request.DocVideo) {
		query["DocVideo"] = request.DocVideo
	}

	if !dara.IsNil(request.DocumentNumber) {
		query["DocumentNumber"] = request.DocumentNumber
	}

	if !dara.IsNil(request.EditOcrResult) {
		query["EditOcrResult"] = request.EditOcrResult
	}

	if !dara.IsNil(request.ExperienceCode) {
		query["ExperienceCode"] = request.ExperienceCode
	}

	if !dara.IsNil(request.FacePictureUrl) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !dara.IsNil(request.IdFaceQuality) {
		query["IdFaceQuality"] = request.IdFaceQuality
	}

	if !dara.IsNil(request.IdSpoof) {
		query["IdSpoof"] = request.IdSpoof
	}

	if !dara.IsNil(request.IdThreshold) {
		query["IdThreshold"] = request.IdThreshold
	}

	if !dara.IsNil(request.LanguageConfig) {
		query["LanguageConfig"] = request.LanguageConfig
	}

	if !dara.IsNil(request.MRTDInput) {
		query["MRTDInput"] = request.MRTDInput
	}

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.MetaInfo) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Ocr) {
		query["Ocr"] = request.Ocr
	}

	if !dara.IsNil(request.Pages) {
		query["Pages"] = request.Pages
	}

	if !dara.IsNil(request.ProcedurePriority) {
		query["ProcedurePriority"] = request.ProcedurePriority
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductFlow) {
		query["ProductFlow"] = request.ProductFlow
	}

	if !dara.IsNil(request.ReturnUrl) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.SecurityLevel) {
		query["SecurityLevel"] = request.SecurityLevel
	}

	if !dara.IsNil(request.ShowAlbumIcon) {
		query["ShowAlbumIcon"] = request.ShowAlbumIcon
	}

	if !dara.IsNil(request.ShowGuidePage) {
		query["ShowGuidePage"] = request.ShowGuidePage
	}

	if !dara.IsNil(request.ShowOcrResult) {
		query["ShowOcrResult"] = request.ShowOcrResult
	}

	if !dara.IsNil(request.StyleConfig) {
		query["StyleConfig"] = request.StyleConfig
	}

	if !dara.IsNil(request.UseNFC) {
		query["UseNFC"] = request.UseNFC
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FacePictureBase64) {
		body["FacePictureBase64"] = request.FacePictureBase64
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Initialize"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 手机号二要素核验API
//
// @param request - Mobile2MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile2MetaVerifyIntlResponse
func (client *Client) Mobile2MetaVerifyIntlWithContext(ctx context.Context, request *Mobile2MetaVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *Mobile2MetaVerifyIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		body["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		body["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile2MetaVerifyIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile2MetaVerifyIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # International Version of Mobile Three Elements API
//
// @param request - Mobile3MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile3MetaVerifyIntlResponse
func (client *Client) Mobile3MetaVerifyIntlWithContext(ctx context.Context, request *Mobile3MetaVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaVerifyIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyNum) {
		query["IdentifyNum"] = request.IdentifyNum
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.ParamType) {
		query["ParamType"] = request.ParamType
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Mobile3MetaVerifyIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &Mobile3MetaVerifyIntlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
