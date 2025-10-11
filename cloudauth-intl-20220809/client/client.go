// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloudauth-intl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := dara.NewRequest()
	boundary := dara.GetBoundary()
	request_.Protocol = dara.String("HTTPS")
	request_.Method = dara.String("POST")
	request_.Pathname = dara.String("/")
	request_.Headers = map[string]*string{
		"host":       dara.String(dara.ToString(form["host"])),
		"date":       openapiutil.GetDateUTCString(),
		"user-agent": openapiutil.GetUserAgent(dara.String("")),
	}
	request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
	request_.Body = dara.ToFileForm(form, boundary)
	response_, _err := dara.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}

	_result, _err = _postOSSObject_opResponse(response_)
	if _err != nil {
		return nil, _err
	}

	return _result, nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Address Similarity Comparison
//
// Description:
//
// API for comparing two addresses, standardizing and checking address consistency.
//
// @param request - AddressCompareIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddressCompareIntlResponse
func (client *Client) AddressCompareIntlWithOptions(request *AddressCompareIntlRequest, runtime *dara.RuntimeOptions) (_result *AddressCompareIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Address Similarity Comparison
//
// Description:
//
// API for comparing two addresses, standardizing and checking address consistency.
//
// @param request - AddressCompareIntlRequest
//
// @return AddressCompareIntlResponse
func (client *Client) AddressCompareIntl(request *AddressCompareIntlRequest) (_result *AddressCompareIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddressCompareIntlResponse{}
	_body, _err := client.AddressCompareIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddressVerifyIntlWithOptions(request *AddressVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *AddressVerifyIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddressVerifyIntlResponse
func (client *Client) AddressVerifyIntl(request *AddressVerifyIntlRequest) (_result *AddressVerifyIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddressVerifyIntlResponse{}
	_body, _err := client.AddressVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the address verification API operation, which verifies the region and address of a device using device data and carrier big data capabilities.
//
// @param request - AddressVerifyV2IntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddressVerifyV2IntlResponse
func (client *Client) AddressVerifyV2IntlWithOptions(request *AddressVerifyV2IntlRequest, runtime *dara.RuntimeOptions) (_result *AddressVerifyV2IntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes the address verification API operation, which verifies the region and address of a device using device data and carrier big data capabilities.
//
// @param request - AddressVerifyV2IntlRequest
//
// @return AddressVerifyV2IntlResponse
func (client *Client) AddressVerifyV2Intl(request *AddressVerifyV2IntlRequest) (_result *AddressVerifyV2IntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddressVerifyV2IntlResponse{}
	_body, _err := client.AddressVerifyV2IntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BankMetaVerifyIntlWithOptions(request *BankMetaVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *BankMetaVerifyIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BankMetaVerifyIntlResponse
func (client *Client) BankMetaVerifyIntl(request *BankMetaVerifyIntlRequest) (_result *BankMetaVerifyIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BankMetaVerifyIntlResponse{}
	_body, _err := client.BankMetaVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CardOcrWithOptions(request *CardOcrRequest, runtime *dara.RuntimeOptions) (_result *CardOcrResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CardOcrResponse
// Deprecated
func (client *Client) CardOcr(request *CardOcrRequest) (_result *CardOcrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CardOcrResponse{}
	_body, _err := client.CardOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckResultWithOptions(request *CheckResultRequest, runtime *dara.RuntimeOptions) (_result *CheckResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckResultResponse
func (client *Client) CheckResult(request *CheckResultRequest) (_result *CheckResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckResultResponse{}
	_body, _err := client.CheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckVerifyLogWithOptions(request *CheckVerifyLogRequest, runtime *dara.RuntimeOptions) (_result *CheckVerifyLogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckVerifyLogResponse
func (client *Client) CheckVerifyLog(request *CheckVerifyLogRequest) (_result *CheckVerifyLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckVerifyLogResponse{}
	_body, _err := client.CheckVerifyLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 凭证识别查询
//
// @param request - CredentialGetResultIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialGetResultIntlResponse
func (client *Client) CredentialGetResultIntlWithOptions(request *CredentialGetResultIntlRequest, runtime *dara.RuntimeOptions) (_result *CredentialGetResultIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TransactionId) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CredentialGetResultIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialGetResultIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 凭证识别查询
//
// @param request - CredentialGetResultIntlRequest
//
// @return CredentialGetResultIntlResponse
func (client *Client) CredentialGetResultIntl(request *CredentialGetResultIntlRequest) (_result *CredentialGetResultIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CredentialGetResultIntlResponse{}
	_body, _err := client.CredentialGetResultIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Credential Recognition
//
// Description:
//
// Detects whether a voucher (such as water, electricity, gas, credit card, etc., e-bills) is forged using AI technology and extracts key information from the voucher.
//
// @param request - CredentialRecognitionIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialRecognitionIntlResponse
func (client *Client) CredentialRecognitionIntlWithOptions(request *CredentialRecognitionIntlRequest, runtime *dara.RuntimeOptions) (_result *CredentialRecognitionIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Credential Recognition
//
// Description:
//
// Detects whether a voucher (such as water, electricity, gas, credit card, etc., e-bills) is forged using AI technology and extracts key information from the voucher.
//
// @param request - CredentialRecognitionIntlRequest
//
// @return CredentialRecognitionIntlResponse
func (client *Client) CredentialRecognitionIntl(request *CredentialRecognitionIntlRequest) (_result *CredentialRecognitionIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CredentialRecognitionIntlResponse{}
	_body, _err := client.CredentialRecognitionIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 凭证识别提交
//
// @param request - CredentialSubmitIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CredentialSubmitIntlResponse
func (client *Client) CredentialSubmitIntlWithOptions(request *CredentialSubmitIntlRequest, runtime *dara.RuntimeOptions) (_result *CredentialSubmitIntlResponse, _err error) {
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

	if !dara.IsNil(request.MerchantBizId) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.OcrArea) {
		query["OcrArea"] = request.OcrArea
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
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
		Action:      dara.String("CredentialSubmitIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CredentialSubmitIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 凭证识别提交
//
// @param request - CredentialSubmitIntlRequest
//
// @return CredentialSubmitIntlResponse
func (client *Client) CredentialSubmitIntl(request *CredentialSubmitIntlRequest) (_result *CredentialSubmitIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CredentialSubmitIntlResponse{}
	_body, _err := client.CredentialSubmitIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CredentialVerifyIntlWithOptions(request *CredentialVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *CredentialVerifyIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CredentialVerifyIntlResponse
func (client *Client) CredentialVerifyIntl(request *CredentialVerifyIntlRequest) (_result *CredentialVerifyIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CredentialVerifyIntlResponse{}
	_body, _err := client.CredentialVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CredentialVerifyIntlAdvance(request *CredentialVerifyIntlAdvanceRequest, runtime *dara.RuntimeOptions) (_result *CredentialVerifyIntlResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("Cloudauth-intl"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	credentialVerifyIntlReq := &CredentialVerifyIntlRequest{}
	openapiutil.Convert(request, credentialVerifyIntlReq)
	if !dara.IsNil(request.ImageFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.ImageFileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		credentialVerifyIntlReq.ImageFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	credentialVerifyIntlResp, _err := client.CredentialVerifyIntlWithOptions(credentialVerifyIntlReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = credentialVerifyIntlResp
	return _result, _err
}

// Summary:
//
// # Face Credential Verification
//
// Description:
//
// Input a face image and use the algorithm to detect if there is a risk of deep forgery. This includes risk scenarios such as AIGC-generated faces, deepfake face swapping, template faces, and rephotographed faces, and outputs risk labels and confidence levels.
//
// @param request - DeepfakeDetectIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepfakeDetectIntlResponse
func (client *Client) DeepfakeDetectIntlWithOptions(request *DeepfakeDetectIntlRequest, runtime *dara.RuntimeOptions) (_result *DeepfakeDetectIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Face Credential Verification
//
// Description:
//
// Input a face image and use the algorithm to detect if there is a risk of deep forgery. This includes risk scenarios such as AIGC-generated faces, deepfake face swapping, template faces, and rephotographed faces, and outputs risk labels and confidence levels.
//
// @param request - DeepfakeDetectIntlRequest
//
// @return DeepfakeDetectIntlResponse
func (client *Client) DeepfakeDetectIntl(request *DeepfakeDetectIntlRequest) (_result *DeepfakeDetectIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeepfakeDetectIntlResponse{}
	_body, _err := client.DeepfakeDetectIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// deepfake文件流api
//
// @param request - DeepfakeDetectIntlStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeepfakeDetectIntlStreamResponse
func (client *Client) DeepfakeDetectIntlStreamWithOptions(request *DeepfakeDetectIntlStreamRequest, runtime *dara.RuntimeOptions) (_result *DeepfakeDetectIntlStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FaceBase64) {
		body["FaceBase64"] = request.FaceBase64
	}

	if !dara.IsNil(request.FaceFile) {
		body["FaceFile"] = request.FaceFile
	}

	if !dara.IsNil(request.FaceInputType) {
		body["FaceInputType"] = request.FaceInputType
	}

	if !dara.IsNil(request.FaceUrl) {
		body["FaceUrl"] = request.FaceUrl
	}

	if !dara.IsNil(request.MerchantBizId) {
		body["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SceneCode) {
		body["SceneCode"] = request.SceneCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeepfakeDetectIntlStream"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeepfakeDetectIntlStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// deepfake文件流api
//
// @param request - DeepfakeDetectIntlStreamRequest
//
// @return DeepfakeDetectIntlStreamResponse
func (client *Client) DeepfakeDetectIntlStream(request *DeepfakeDetectIntlStreamRequest) (_result *DeepfakeDetectIntlStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeepfakeDetectIntlStreamResponse{}
	_body, _err := client.DeepfakeDetectIntlStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeepfakeDetectIntlStreamAdvance(request *DeepfakeDetectIntlStreamAdvanceRequest, runtime *dara.RuntimeOptions) (_result *DeepfakeDetectIntlStreamResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("Cloudauth-intl"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	deepfakeDetectIntlStreamReq := &DeepfakeDetectIntlStreamRequest{}
	openapiutil.Convert(request, deepfakeDetectIntlStreamReq)
	if !dara.IsNil(request.FaceFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FaceFileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		deepfakeDetectIntlStreamReq.FaceFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	deepfakeDetectIntlStreamResp, _err := client.DeepfakeDetectIntlStreamWithOptions(deepfakeDetectIntlStreamReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = deepfakeDetectIntlStreamResp
	return _result, _err
}

// Summary:
//
// # Delete Face Group
//
// @param request - DeleteFaceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceGroupResponse
func (client *Client) DeleteFaceGroupWithOptions(request *DeleteFaceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceGroup"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Face Group
//
// @param request - DeleteFaceGroupRequest
//
// @return DeleteFaceGroupResponse
func (client *Client) DeleteFaceGroup(request *DeleteFaceGroupRequest) (_result *DeleteFaceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaceGroupResponse{}
	_body, _err := client.DeleteFaceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Face
//
// @param request - DeleteFaceRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFaceRecordResponse
func (client *Client) DeleteFaceRecordWithOptions(request *DeleteFaceRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteFaceRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFaceRecord"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFaceRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Face
//
// @param request - DeleteFaceRecordRequest
//
// @return DeleteFaceRecordResponse
func (client *Client) DeleteFaceRecord(request *DeleteFaceRecordRequest) (_result *DeleteFaceRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFaceRecordResponse{}
	_body, _err := client.DeleteFaceRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteVerifyResultWithOptions(request *DeleteVerifyResultRequest, runtime *dara.RuntimeOptions) (_result *DeleteVerifyResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteVerifyResultResponse
func (client *Client) DeleteVerifyResult(request *DeleteVerifyResultRequest) (_result *DeleteVerifyResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVerifyResultResponse{}
	_body, _err := client.DeleteVerifyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DocOcrWithOptions(request *DocOcrRequest, runtime *dara.RuntimeOptions) (_result *DocOcrResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DocOcrResponse
func (client *Client) DocOcr(request *DocOcrRequest) (_result *DocOcrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DocOcrResponse{}
	_body, _err := client.DocOcrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Global Document OCR Recognition Interface
//
// @param request - DocOcrMaxRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DocOcrMaxResponse
func (client *Client) DocOcrMaxWithOptions(request *DocOcrMaxRequest, runtime *dara.RuntimeOptions) (_result *DocOcrMaxResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Global Document OCR Recognition Interface
//
// @param request - DocOcrMaxRequest
//
// @return DocOcrMaxResponse
func (client *Client) DocOcrMax(request *DocOcrMaxRequest) (_result *DocOcrMaxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DocOcrMaxResponse{}
	_body, _err := client.DocOcrMaxWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes how to integrate with ID Verification using only the server-side API.
//
// @param request - EkycVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EkycVerifyResponse
func (client *Client) EkycVerifyWithOptions(request *EkycVerifyRequest, runtime *dara.RuntimeOptions) (_result *EkycVerifyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes how to integrate with ID Verification using only the server-side API.
//
// @param request - EkycVerifyRequest
//
// @return EkycVerifyResponse
func (client *Client) EkycVerify(request *EkycVerifyRequest) (_result *EkycVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EkycVerifyResponse{}
	_body, _err := client.EkycVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes how to integrate FaceCompare using only the server-side API.
//
// @param request - FaceCompareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceCompareResponse
func (client *Client) FaceCompareWithOptions(request *FaceCompareRequest, runtime *dara.RuntimeOptions) (_result *FaceCompareResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes how to integrate FaceCompare using only the server-side API.
//
// @param request - FaceCompareRequest
//
// @return FaceCompareResponse
func (client *Client) FaceCompare(request *FaceCompareRequest) (_result *FaceCompareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FaceCompareResponse{}
	_body, _err := client.FaceCompareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人脸交叉比对
//
// @param request - FaceCrossCompareIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceCrossCompareIntlResponse
func (client *Client) FaceCrossCompareIntlWithOptions(request *FaceCrossCompareIntlRequest, runtime *dara.RuntimeOptions) (_result *FaceCrossCompareIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompareModel) {
		query["CompareModel"] = request.CompareModel
	}

	if !dara.IsNil(request.FaceVerifyThreshold) {
		query["FaceVerifyThreshold"] = request.FaceVerifyThreshold
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

	if !dara.IsNil(request.SourceAFacePicture) {
		query["SourceAFacePicture"] = request.SourceAFacePicture
	}

	if !dara.IsNil(request.SourceAFacePictureUrl) {
		query["SourceAFacePictureUrl"] = request.SourceAFacePictureUrl
	}

	if !dara.IsNil(request.SourceBFacePicture) {
		query["SourceBFacePicture"] = request.SourceBFacePicture
	}

	if !dara.IsNil(request.SourceBFacePictureUrl) {
		query["SourceBFacePictureUrl"] = request.SourceBFacePictureUrl
	}

	if !dara.IsNil(request.SourceCFacePicture) {
		query["SourceCFacePicture"] = request.SourceCFacePicture
	}

	if !dara.IsNil(request.SourceCFacePictureUrl) {
		query["SourceCFacePictureUrl"] = request.SourceCFacePictureUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FaceCrossCompareIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FaceCrossCompareIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸交叉比对
//
// @param request - FaceCrossCompareIntlRequest
//
// @return FaceCrossCompareIntlResponse
func (client *Client) FaceCrossCompareIntl(request *FaceCrossCompareIntlRequest) (_result *FaceCrossCompareIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FaceCrossCompareIntlResponse{}
	_body, _err := client.FaceCrossCompareIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Face Duplication Detection API
//
// @param request - FaceDuplicationCheckIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceDuplicationCheckIntlResponse
func (client *Client) FaceDuplicationCheckIntlWithOptions(request *FaceDuplicationCheckIntlRequest, runtime *dara.RuntimeOptions) (_result *FaceDuplicationCheckIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRegistration) {
		body["AutoRegistration"] = request.AutoRegistration
	}

	if !dara.IsNil(request.FaceGroupCodes) {
		body["FaceGroupCodes"] = request.FaceGroupCodes
	}

	if !dara.IsNil(request.FaceRegisterGroupCode) {
		body["FaceRegisterGroupCode"] = request.FaceRegisterGroupCode
	}

	if !dara.IsNil(request.FaceVerifyThreshold) {
		body["FaceVerifyThreshold"] = request.FaceVerifyThreshold
	}

	if !dara.IsNil(request.Liveness) {
		body["Liveness"] = request.Liveness
	}

	if !dara.IsNil(request.MerchantBizId) {
		body["MerchantBizId"] = request.MerchantBizId
	}

	if !dara.IsNil(request.MerchantUserId) {
		body["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.ReturnFaces) {
		body["ReturnFaces"] = request.ReturnFaces
	}

	if !dara.IsNil(request.SaveFacePicture) {
		body["SaveFacePicture"] = request.SaveFacePicture
	}

	if !dara.IsNil(request.SceneCode) {
		body["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.SourceFacePicture) {
		body["SourceFacePicture"] = request.SourceFacePicture
	}

	if !dara.IsNil(request.SourceFacePictureUrl) {
		body["SourceFacePictureUrl"] = request.SourceFacePictureUrl
	}

	if !dara.IsNil(request.TargetFacePicture) {
		body["TargetFacePicture"] = request.TargetFacePicture
	}

	if !dara.IsNil(request.TargetFacePictureUrl) {
		body["TargetFacePictureUrl"] = request.TargetFacePictureUrl
	}

	if !dara.IsNil(request.VerifyModel) {
		body["VerifyModel"] = request.VerifyModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FaceDuplicationCheckIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FaceDuplicationCheckIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Face Duplication Detection API
//
// @param request - FaceDuplicationCheckIntlRequest
//
// @return FaceDuplicationCheckIntlResponse
func (client *Client) FaceDuplicationCheckIntl(request *FaceDuplicationCheckIntlRequest) (_result *FaceDuplicationCheckIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FaceDuplicationCheckIntlResponse{}
	_body, _err := client.FaceDuplicationCheckIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes how to set up the server for FACE_GUARD.
//
// @param request - FaceGuardRiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceGuardRiskResponse
func (client *Client) FaceGuardRiskWithOptions(request *FaceGuardRiskRequest, runtime *dara.RuntimeOptions) (_result *FaceGuardRiskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic describes how to set up the server for FACE_GUARD.
//
// @param request - FaceGuardRiskRequest
//
// @return FaceGuardRiskResponse
func (client *Client) FaceGuardRisk(request *FaceGuardRiskRequest) (_result *FaceGuardRiskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FaceGuardRiskResponse{}
	_body, _err := client.FaceGuardRiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Passive liveness detection (FaceLiveness) is a service that detects whether a pre-captured facial image, submitted to an API operation, is a real face. The algorithm primarily detects presentation attacks, such as screen replays and printed photos. This service is suitable for low-risk business scen
//
// @param request - FaceLivenessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FaceLivenessResponse
func (client *Client) FaceLivenessWithOptions(request *FaceLivenessRequest, runtime *dara.RuntimeOptions) (_result *FaceLivenessResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Passive liveness detection (FaceLiveness) is a service that detects whether a pre-captured facial image, submitted to an API operation, is a real face. The algorithm primarily detects presentation attacks, such as screen replays and printed photos. This service is suitable for low-risk business scen
//
// @param request - FaceLivenessRequest
//
// @return FaceLivenessResponse
func (client *Client) FaceLiveness(request *FaceLivenessRequest) (_result *FaceLivenessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FaceLivenessResponse{}
	_body, _err := client.FaceLivenessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Anti-Fraud Callback Interface
//
// @param request - FraudResultCallBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FraudResultCallBackResponse
func (client *Client) FraudResultCallBackWithOptions(request *FraudResultCallBackRequest, runtime *dara.RuntimeOptions) (_result *FraudResultCallBackResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Anti-Fraud Callback Interface
//
// @param request - FraudResultCallBackRequest
//
// @return FraudResultCallBackResponse
func (client *Client) FraudResultCallBack(request *FraudResultCallBackRequest) (_result *FraudResultCallBackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FraudResultCallBackResponse{}
	_body, _err := client.FraudResultCallBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation verifies the authenticity and consistency of a name, ID card number, and the start and end dates of the ID card\\"s validity period against an authoritative source.
//
// @param request - Id2MetaPeriodVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaPeriodVerifyIntlResponse
func (client *Client) Id2MetaPeriodVerifyIntlWithOptions(request *Id2MetaPeriodVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaPeriodVerifyIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation verifies the authenticity and consistency of a name, ID card number, and the start and end dates of the ID card\\"s validity period against an authoritative source.
//
// @param request - Id2MetaPeriodVerifyIntlRequest
//
// @return Id2MetaPeriodVerifyIntlResponse
func (client *Client) Id2MetaPeriodVerifyIntl(request *Id2MetaPeriodVerifyIntlRequest) (_result *Id2MetaPeriodVerifyIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Id2MetaPeriodVerifyIntlResponse{}
	_body, _err := client.Id2MetaPeriodVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies that a name and an ID card number are consistent.
//
// @param request - Id2MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Id2MetaVerifyIntlResponse
func (client *Client) Id2MetaVerifyIntlWithOptions(request *Id2MetaVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *Id2MetaVerifyIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies that a name and an ID card number are consistent.
//
// @param request - Id2MetaVerifyIntlRequest
//
// @return Id2MetaVerifyIntlResponse
func (client *Client) Id2MetaVerifyIntl(request *Id2MetaVerifyIntlRequest) (_result *Id2MetaVerifyIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Id2MetaVerifyIntlResponse{}
	_body, _err := client.Id2MetaVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Authentication Initialization
//
// @param tmpReq - InitializeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeResponse
func (client *Client) InitializeWithOptions(tmpReq *InitializeRequest, runtime *dara.RuntimeOptions) (_result *InitializeResponse, _err error) {
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

	if !dara.IsNil(request.AutoRegistration) {
		query["AutoRegistration"] = request.AutoRegistration
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

	if !dara.IsNil(request.FaceGroupCodes) {
		query["FaceGroupCodes"] = request.FaceGroupCodes
	}

	if !dara.IsNil(request.FacePictureUrl) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !dara.IsNil(request.FaceRegisterGroupCode) {
		query["FaceRegisterGroupCode"] = request.FaceRegisterGroupCode
	}

	if !dara.IsNil(request.FaceVerifyThreshold) {
		query["FaceVerifyThreshold"] = request.FaceVerifyThreshold
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

	if !dara.IsNil(request.ReturnFaces) {
		query["ReturnFaces"] = request.ReturnFaces
	}

	if !dara.IsNil(request.ReturnUrl) {
		query["ReturnUrl"] = request.ReturnUrl
	}

	if !dara.IsNil(request.SaveFacePicture) {
		query["SaveFacePicture"] = request.SaveFacePicture
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

	if !dara.IsNil(request.TargetFacePicture) {
		query["TargetFacePicture"] = request.TargetFacePicture
	}

	if !dara.IsNil(request.TargetFacePictureUrl) {
		query["TargetFacePictureUrl"] = request.TargetFacePictureUrl
	}

	if !dara.IsNil(request.UseNFC) {
		query["UseNFC"] = request.UseNFC
	}

	if !dara.IsNil(request.VerifyModel) {
		query["VerifyModel"] = request.VerifyModel
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Authentication Initialization
//
// @param request - InitializeRequest
//
// @return InitializeResponse
func (client *Client) Initialize(request *InitializeRequest) (_result *InitializeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitializeResponse{}
	_body, _err := client.InitializeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If your server makes infrequent calls to the ID Verification API, you can call the KeepaliveIntl operation to maintain the client connection.
//
// @param request - KeepaliveIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KeepaliveIntlResponse
func (client *Client) KeepaliveIntlWithOptions(runtime *dara.RuntimeOptions) (_result *KeepaliveIntlResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("KeepaliveIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KeepaliveIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If your server makes infrequent calls to the ID Verification API, you can call the KeepaliveIntl operation to maintain the client connection.
//
// @return KeepaliveIntlResponse
func (client *Client) KeepaliveIntl() (_result *KeepaliveIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KeepaliveIntlResponse{}
	_body, _err := client.KeepaliveIntlWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of a mobile number and name against an authoritative data source.
//
// @param request - Mobile2MetaVerifyIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return Mobile2MetaVerifyIntlResponse
func (client *Client) Mobile2MetaVerifyIntlWithOptions(request *Mobile2MetaVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *Mobile2MetaVerifyIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the authenticity and consistency of a mobile number and name against an authoritative data source.
//
// @param request - Mobile2MetaVerifyIntlRequest
//
// @return Mobile2MetaVerifyIntlResponse
func (client *Client) Mobile2MetaVerifyIntl(request *Mobile2MetaVerifyIntlRequest) (_result *Mobile2MetaVerifyIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Mobile2MetaVerifyIntlResponse{}
	_body, _err := client.Mobile2MetaVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) Mobile3MetaVerifyIntlWithOptions(request *Mobile3MetaVerifyIntlRequest, runtime *dara.RuntimeOptions) (_result *Mobile3MetaVerifyIntlResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return Mobile3MetaVerifyIntlResponse
func (client *Client) Mobile3MetaVerifyIntl(request *Mobile3MetaVerifyIntlRequest) (_result *Mobile3MetaVerifyIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &Mobile3MetaVerifyIntlResponse{}
	_body, _err := client.Mobile3MetaVerifyIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改人脸库
//
// @param request - ModifyFaceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFaceGroupResponse
func (client *Client) ModifyFaceGroupWithOptions(request *ModifyFaceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyFaceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFaceGroup"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFaceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改人脸库
//
// @param request - ModifyFaceGroupRequest
//
// @return ModifyFaceGroupResponse
func (client *Client) ModifyFaceGroup(request *ModifyFaceGroupRequest) (_result *ModifyFaceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFaceGroupResponse{}
	_body, _err := client.ModifyFaceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增人脸
//
// @param request - ModifyFaceRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFaceRecordResponse
func (client *Client) ModifyFaceRecordWithOptions(request *ModifyFaceRecordRequest, runtime *dara.RuntimeOptions) (_result *ModifyFaceRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FaceGroupCode) {
		body["FaceGroupCode"] = request.FaceGroupCode
	}

	if !dara.IsNil(request.ImgOssInfos) {
		body["ImgOssInfos"] = request.ImgOssInfos
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFaceRecord"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFaceRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增人脸
//
// @param request - ModifyFaceRecordRequest
//
// @return ModifyFaceRecordResponse
func (client *Client) ModifyFaceRecord(request *ModifyFaceRecordRequest) (_result *ModifyFaceRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFaceRecordResponse{}
	_body, _err := client.ModifyFaceRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询人脸库
//
// @param request - QueryFaceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceGroupResponse
func (client *Client) QueryFaceGroupWithOptions(request *QueryFaceGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.GroupCode) {
		query["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceGroup"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询人脸库
//
// @param request - QueryFaceGroupRequest
//
// @return QueryFaceGroupResponse
func (client *Client) QueryFaceGroup(request *QueryFaceGroupRequest) (_result *QueryFaceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceGroupResponse{}
	_body, _err := client.QueryFaceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询人脸记录
//
// @param request - QueryFaceRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFaceRecordResponse
func (client *Client) QueryFaceRecordWithOptions(request *QueryFaceRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryFaceRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FaceGroupCode) {
		query["FaceGroupCode"] = request.FaceGroupCode
	}

	if !dara.IsNil(request.FaceId) {
		query["FaceId"] = request.FaceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MerchantUserId) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegistrationType) {
		query["RegistrationType"] = request.RegistrationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryFaceRecord"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryFaceRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询人脸记录
//
// @param request - QueryFaceRecordRequest
//
// @return QueryFaceRecordResponse
func (client *Client) QueryFaceRecord(request *QueryFaceRecordRequest) (_result *QueryFaceRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryFaceRecordResponse{}
	_body, _err := client.QueryFaceRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取临时token
//
// @param request - TempAccessTokenIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TempAccessTokenIntlResponse
func (client *Client) TempAccessTokenIntlWithOptions(request *TempAccessTokenIntlRequest, runtime *dara.RuntimeOptions) (_result *TempAccessTokenIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TempAccessTokenIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TempAccessTokenIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取临时token
//
// @param request - TempAccessTokenIntlRequest
//
// @return TempAccessTokenIntlResponse
func (client *Client) TempAccessTokenIntl(request *TempAccessTokenIntlRequest) (_result *TempAccessTokenIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TempAccessTokenIntlResponse{}
	_body, _err := client.TempAccessTokenIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件临时地址
//
// @param request - TempOssUrlIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TempOssUrlIntlResponse
func (client *Client) TempOssUrlIntlWithOptions(request *TempOssUrlIntlRequest, runtime *dara.RuntimeOptions) (_result *TempOssUrlIntlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ObjectName) {
		body["ObjectName"] = request.ObjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TempOssUrlIntl"),
		Version:     dara.String("2022-08-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TempOssUrlIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件临时地址
//
// @param request - TempOssUrlIntlRequest
//
// @return TempOssUrlIntlResponse
func (client *Client) TempOssUrlIntl(request *TempOssUrlIntlRequest) (_result *TempOssUrlIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TempOssUrlIntlResponse{}
	_body, _err := client.TempOssUrlIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}
