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
	EnableValidate  *bool
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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"us-west-1":             dara.String("dytnsapi.aliyuncs.com"),
		"us-east-1":             dara.String("dytnsapi.aliyuncs.com"),
		"me-east-1":             dara.String("dytnsapi.aliyuncs.com"),
		"eu-west-1":             dara.String("dytnsapi.aliyuncs.com"),
		"eu-central-1":          dara.String("dytnsapi.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("dytnsapi.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("dytnsapi.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("dytnsapi.aliyuncs.com"),
		"cn-shenzhen":           dara.String("dytnsapi.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("dytnsapi.aliyuncs.com"),
		"cn-shanghai":           dara.String("dytnsapi.aliyuncs.com"),
		"cn-qingdao":            dara.String("dytnsapi.aliyuncs.com"),
		"cn-huhehaote":          dara.String("dytnsapi.aliyuncs.com"),
		"cn-hongkong":           dara.String("dytnsapi.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("dytnsapi.aliyuncs.com"),
		"cn-hangzhou":           dara.String("dytnsapi.aliyuncs.com"),
		"cn-chengdu":            dara.String("dytnsapi.aliyuncs.com"),
		"cn-beijing-finance-1":  dara.String("dytnsapi.aliyuncs.com"),
		"cn-beijing":            dara.String("dytnsapi.aliyuncs.com"),
		"ap-southeast-5":        dara.String("dytnsapi.aliyuncs.com"),
		"ap-southeast-3":        dara.String("dytnsapi.aliyuncs.com"),
		"ap-southeast-1":        dara.String("dytnsapi.aliyuncs.com"),
		"ap-northeast-1":        dara.String("dytnsapi.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dytnsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
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
// Three-element (name, ID card number, and ID card photo) verification. Used to compare whether the three pieces of information (the name and ID card number filled in by the user and the uploaded portrait photo) belong to the same user.
//
// Description:
//
// - Before using this API, log on to the Cell Phone Number Service console, go to the Tag Square page, find the corresponding tag, click Apply to Activate, fill in the application materials, and use the tag after the application is approved.
//
// - Before using this API, make sure that you have fully understood the [Cell Phone Number Service pricing](https://help.aliyun.com/document_detail/154751.html).
//
// - Billing applies only when the API return value is Code="OK". Other return results are not billed.
//
// - For the verifiable scope, see [ID Card Three Elements Verification](https://help.aliyun.com/document_detail/2844379.html).
//
// ### QPS limit
//
// - The per-user QPS limit for this API is 200 calls per second. If you exceed the limit, API calls are throttled, which may affect your business. Call the API reasonably.
//
// @param request - CertNoThreeElementVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertNoThreeElementVerificationResponse
func (client *Client) CertNoThreeElementVerificationWithOptions(request *CertNoThreeElementVerificationRequest, runtime *dara.RuntimeOptions) (_result *CertNoThreeElementVerificationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Three-element (name, ID card number, and ID card photo) verification. Used to compare whether the three pieces of information (the name and ID card number filled in by the user and the uploaded portrait photo) belong to the same user.
//
// Description:
//
// - Before using this API, log on to the Cell Phone Number Service console, go to the Tag Square page, find the corresponding tag, click Apply to Activate, fill in the application materials, and use the tag after the application is approved.
//
// - Before using this API, make sure that you have fully understood the [Cell Phone Number Service pricing](https://help.aliyun.com/document_detail/154751.html).
//
// - Billing applies only when the API return value is Code="OK". Other return results are not billed.
//
// - For the verifiable scope, see [ID Card Three Elements Verification](https://help.aliyun.com/document_detail/2844379.html).
//
// ### QPS limit
//
// - The per-user QPS limit for this API is 200 calls per second. If you exceed the limit, API calls are throttled, which may affect your business. Call the API reasonably.
//
// @param request - CertNoThreeElementVerificationRequest
//
// @return CertNoThreeElementVerificationResponse
func (client *Client) CertNoThreeElementVerification(request *CertNoThreeElementVerificationRequest) (_result *CertNoThreeElementVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CertNoThreeElementVerificationResponse{}
	_body, _err := client.CertNoThreeElementVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Two-element (name and ID card number) verification. Used to verify whether the name and ID card number entered by a user belong to the same person.
//
// Description:
//
// - Before using this operation, make sure that you have fully understood the pricing of Cell Phone Number Service.
//
// - Before using this operation, log on to the Cell Phone Number Service console. On the Tag Marketplace page, find the desired tag, click Apply to Activate, and fill in the application information. You can use the operation only after your application is approved.
//
// - The operation is charged when the response contains Code="OK" and IsConsistent != 2. Other response results are not charged.
//
// ## QPS Limit
//
// The per-user QPS limit of this operation is 200 calls per second. If the limit is exceeded, the API calls are throttled, which may affect your business. Call the operation properly.
//
// @param request - CertNoTwoElementVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CertNoTwoElementVerificationResponse
func (client *Client) CertNoTwoElementVerificationWithOptions(request *CertNoTwoElementVerificationRequest, runtime *dara.RuntimeOptions) (_result *CertNoTwoElementVerificationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Two-element (name and ID card number) verification. Used to verify whether the name and ID card number entered by a user belong to the same person.
//
// Description:
//
// - Before using this operation, make sure that you have fully understood the pricing of Cell Phone Number Service.
//
// - Before using this operation, log on to the Cell Phone Number Service console. On the Tag Marketplace page, find the desired tag, click Apply to Activate, and fill in the application information. You can use the operation only after your application is approved.
//
// - The operation is charged when the response contains Code="OK" and IsConsistent != 2. Other response results are not charged.
//
// ## QPS Limit
//
// The per-user QPS limit of this operation is 200 calls per second. If the limit is exceeded, the API calls are throttled, which may affect your business. Call the operation properly.
//
// @param request - CertNoTwoElementVerificationRequest
//
// @return CertNoTwoElementVerificationResponse
func (client *Client) CertNoTwoElementVerification(request *CertNoTwoElementVerificationRequest) (_result *CertNoTwoElementVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CertNoTwoElementVerificationResponse{}
	_body, _err := client.CertNoTwoElementVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the enterprise name, enterprise certificate number, enterprise legal person name, and enterprise legal person ID card number. All four items must be consistent and the enterprise\\"s operating status must be active for verification to pass.
//
// Description:
//
// - Before you use this API, make sure that you fully understand the billing method and prices of the Enterprise Four-Element Verification product. For billing details, see [Product Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
// - Billing applies when VerifyResult returns true/false and ReasonCode is 0, 1, or 2 in the response. Billing does not apply in other cases.
//
// - Before you use this API, log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC&lang=zh), go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply to Activate**, and submit your application. You can use the API after the application is approved.
//
// ### QPS limit
//
// The per-user QPS limit for this API is 200 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the API appropriately.
//
// @param request - CompanyFourElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyFourElementsVerificationResponse
func (client *Client) CompanyFourElementsVerificationWithOptions(request *CompanyFourElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *CompanyFourElementsVerificationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the enterprise name, enterprise certificate number, enterprise legal person name, and enterprise legal person ID card number. All four items must be consistent and the enterprise\\"s operating status must be active for verification to pass.
//
// Description:
//
// - Before you use this API, make sure that you fully understand the billing method and prices of the Enterprise Four-Element Verification product. For billing details, see [Product Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
// - Billing applies when VerifyResult returns true/false and ReasonCode is 0, 1, or 2 in the response. Billing does not apply in other cases.
//
// - Before you use this API, log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC&lang=zh), go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply to Activate**, and submit your application. You can use the API after the application is approved.
//
// ### QPS limit
//
// The per-user QPS limit for this API is 200 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the API appropriately.
//
// @param request - CompanyFourElementsVerificationRequest
//
// @return CompanyFourElementsVerificationResponse
func (client *Client) CompanyFourElementsVerification(request *CompanyFourElementsVerificationRequest) (_result *CompanyFourElementsVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompanyFourElementsVerificationResponse{}
	_body, _err := client.CompanyFourElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Used to verify a company name, company certificate number, and the name of the company\\"s legal representative. Verification passes only when all three are consistent and the company\\"s operating status is in business.
//
// Description:
//
// - Before you use this API, make sure that you have fully understood the billing method and pricing of the Three-Element Company Verification product. For billing details, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
// - In the returned result, charges apply when VerifyResult is true/false and ReasonCode is 0/1/2. No charges apply in other cases.
//
// - Before you use this API, log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC&lang=zh), go to the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, and submit the application materials. You can use the API after the application is approved.
//
// ### QPS limit
//
// The per-user QPS limit for this API is 200 calls per second. Requests that exceed this limit are throttled, which may affect your business. Call this API at a reasonable rate.
//
// @param request - CompanyThreeElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyThreeElementsVerificationResponse
func (client *Client) CompanyThreeElementsVerificationWithOptions(request *CompanyThreeElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *CompanyThreeElementsVerificationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used to verify a company name, company certificate number, and the name of the company\\"s legal representative. Verification passes only when all three are consistent and the company\\"s operating status is in business.
//
// Description:
//
// - Before you use this API, make sure that you have fully understood the billing method and pricing of the Three-Element Company Verification product. For billing details, see [Billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
// - In the returned result, charges apply when VerifyResult is true/false and ReasonCode is 0/1/2. No charges apply in other cases.
//
// - Before you use this API, log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC&lang=zh), go to the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, and submit the application materials. You can use the API after the application is approved.
//
// ### QPS limit
//
// The per-user QPS limit for this API is 200 calls per second. Requests that exceed this limit are throttled, which may affect your business. Call this API at a reasonable rate.
//
// @param request - CompanyThreeElementsVerificationRequest
//
// @return CompanyThreeElementsVerificationResponse
func (client *Client) CompanyThreeElementsVerification(request *CompanyThreeElementsVerificationRequest) (_result *CompanyThreeElementsVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompanyThreeElementsVerificationResponse{}
	_body, _err := client.CompanyThreeElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Used to verify the enterprise name and enterprise certificate number. The name and certificate must be fully consistent, and the enterprise must be in operating status for the verification to pass.
//
// Description:
//
// - Before you use this API, make sure that you fully understand the billing methods and pricing of the Enterprise Two-Element Verification product. For billing details, see [Product billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
// - Billing applies when Code=OK and ReasonCode=0/1/3 in the returned result. Other cases are not billed.
//
// - Before you use this API, log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC&lang=zh), find the corresponding tag on the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, click **Apply to Activate**, and fill in the application materials. After the application is approved, you can use the API.
//
// ### QPS limit
//
// The single-user QPS limit for this API is 200 times/second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call the API reasonably.
//
// @param request - CompanyTwoElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompanyTwoElementsVerificationResponse
func (client *Client) CompanyTwoElementsVerificationWithOptions(request *CompanyTwoElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *CompanyTwoElementsVerificationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used to verify the enterprise name and enterprise certificate number. The name and certificate must be fully consistent, and the enterprise must be in operating status for the verification to pass.
//
// Description:
//
// - Before you use this API, make sure that you fully understand the billing methods and pricing of the Enterprise Two-Element Verification product. For billing details, see [Product billing](https://help.aliyun.com/document_detail/154751.html?spm=a2c4g.154007.0.0.3edd7eb6E90YT4).
//
// - Billing applies when Code=OK and ReasonCode=0/1/3 in the returned result. Other cases are not billed.
//
// - Before you use this API, log on to the [Cell Phone Number Service console](https://account.aliyun.com/login/login.htm?oauth_callback=https%3A%2F%2Fdytns.console.aliyun.com%2Foverview%3Fspm%3Da2c4g.608385.0.0.79847f8b3awqUC&lang=zh), find the corresponding tag on the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, click **Apply to Activate**, and fill in the application materials. After the application is approved, you can use the API.
//
// ### QPS limit
//
// The single-user QPS limit for this API is 200 times/second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call the API reasonably.
//
// @param request - CompanyTwoElementsVerificationRequest
//
// @return CompanyTwoElementsVerificationResponse
func (client *Client) CompanyTwoElementsVerification(request *CompanyTwoElementsVerificationRequest) (_result *CompanyTwoElementsVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CompanyTwoElementsVerificationResponse{}
	_body, _err := client.CompanyTwoElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a contact.
//
// Description:
//
// - Make sure you have activated Phone Number Identity Service before calling this operation.
//
// @param request - DeleteContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactsResponse
func (client *Client) DeleteContactsWithOptions(request *DeleteContactsRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a contact.
//
// Description:
//
// - Make sure you have activated Phone Number Identity Service before calling this operation.
//
// @param request - DeleteContactsRequest
//
// @return DeleteContactsResponse
func (client *Client) DeleteContacts(request *DeleteContactsRequest) (_result *DeleteContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContactsResponse{}
	_body, _err := client.DeleteContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Predicts whether a phone number is invalid by using AI algorithms.
//
// Description:
//
// - 本接口用于验证号码是否为空号。发起调用该接口验证号码请求时，系统会根据验证次数计费，标准价为0.01元/次。**请确保在使用该接口前，已充分了解本产品的收费方式和价格。**
//
// - 当返回结果中：Code="OK" 且 Status != UNKNOWN 时计费，其他情况不计费。
//
// - 由于本产品通过AI算法预测手机号的空号概率，所以无法做到100%准确。当前评估的准确率和召回率约为95%左右。**调用时请注意差别**。
//
// - 使用本接口前，请登录号码百科控制台，在[标签广场](https://dytns.console.aliyun.com/analysis/square)页面，找到对应的标签，单击**申请开通**，填写申请资料，审批通过后即可使用。
//
// ### QPS限制
//
// 本接口的单用户QPS限制为100次/秒。超过限制，API调用会被限流，这可能会影响您的业务，请合理调用。
//
// ### 授权信息
//
// 默认仅限阿里云账号使用本接口，RAM用户只有在被授予了相关API操作权限后方可使用。具体请参见[为RAM用户授权](https://help.aliyun.com/document_detail/154006.html)。
//
// @param request - DescribeEmptyNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEmptyNumberResponse
func (client *Client) DescribeEmptyNumberWithOptions(request *DescribeEmptyNumberRequest, runtime *dara.RuntimeOptions) (_result *DescribeEmptyNumberResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Predicts whether a phone number is invalid by using AI algorithms.
//
// Description:
//
// - 本接口用于验证号码是否为空号。发起调用该接口验证号码请求时，系统会根据验证次数计费，标准价为0.01元/次。**请确保在使用该接口前，已充分了解本产品的收费方式和价格。**
//
// - 当返回结果中：Code="OK" 且 Status != UNKNOWN 时计费，其他情况不计费。
//
// - 由于本产品通过AI算法预测手机号的空号概率，所以无法做到100%准确。当前评估的准确率和召回率约为95%左右。**调用时请注意差别**。
//
// - 使用本接口前，请登录号码百科控制台，在[标签广场](https://dytns.console.aliyun.com/analysis/square)页面，找到对应的标签，单击**申请开通**，填写申请资料，审批通过后即可使用。
//
// ### QPS限制
//
// 本接口的单用户QPS限制为100次/秒。超过限制，API调用会被限流，这可能会影响您的业务，请合理调用。
//
// ### 授权信息
//
// 默认仅限阿里云账号使用本接口，RAM用户只有在被授予了相关API操作权限后方可使用。具体请参见[为RAM用户授权](https://help.aliyun.com/document_detail/154006.html)。
//
// @param request - DescribeEmptyNumberRequest
//
// @return DescribeEmptyNumberResponse
func (client *Client) DescribeEmptyNumber(request *DescribeEmptyNumberRequest) (_result *DescribeEmptyNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEmptyNumberResponse{}
	_body, _err := client.DescribeEmptyNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeMobileOperatorAttributeWithOptions(request *DescribeMobileOperatorAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMobileOperatorAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeMobileOperatorAttributeResponse
func (client *Client) DescribeMobileOperatorAttribute(request *DescribeMobileOperatorAttributeRequest) (_result *DescribeMobileOperatorAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMobileOperatorAttributeResponse{}
	_body, _err := client.DescribeMobileOperatorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the result returned by phone number analysis.
//
// Description:
//
// - Before you use this API, log on to the Phone Number Service console, go to the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply to Activate**, fill in the application materials, and then use the API after the application is approved.
//
// - Before you use this API, make sure that you fully understand the [billing](https://help.aliyun.com/document_detail/154008.html) of Phone Number Service.
//
// ### QPS limit
//
// The QPS limit of this API is 1,000 calls per second per user. If the limit is exceeded, API calls are throttled, which may affect your business. Make calls properly.
//
// ### Authorization information
//
// By default, only Alibaba Cloud accounts can use this API. RAM users can use this API only after they are granted the related API operation permissions. For more information, see [Authorize a RAM user](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribePhoneNumberAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisResponse
func (client *Client) DescribePhoneNumberAnalysisWithOptions(request *DescribePhoneNumberAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the result returned by phone number analysis.
//
// Description:
//
// - Before you use this API, log on to the Phone Number Service console, go to the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply to Activate**, fill in the application materials, and then use the API after the application is approved.
//
// - Before you use this API, make sure that you fully understand the [billing](https://help.aliyun.com/document_detail/154008.html) of Phone Number Service.
//
// ### QPS limit
//
// The QPS limit of this API is 1,000 calls per second per user. If the limit is exceeded, API calls are throttled, which may affect your business. Make calls properly.
//
// ### Authorization information
//
// By default, only Alibaba Cloud accounts can use this API. RAM users can use this API only after they are granted the related API operation permissions. For more information, see [Authorize a RAM user](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribePhoneNumberAnalysisRequest
//
// @return DescribePhoneNumberAnalysisResponse
func (client *Client) DescribePhoneNumberAnalysis(request *DescribePhoneNumberAnalysisRequest) (_result *DescribePhoneNumberAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the phone number analysis result.
//
// Description:
//
// Before using this operation, log on to the Cell Phone Number Service console, go to the Tag Square page, find the corresponding tag, click Apply to activate, and fill in the application materials. You can use the operation after the application is approved.
//
// Make sure that you fully understand the billing of the Cell Phone Number Service before you use this operation.
//
// @param request - DescribePhoneNumberAnalysisAIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAnalysisAIResponse
func (client *Client) DescribePhoneNumberAnalysisAIWithOptions(request *DescribePhoneNumberAnalysisAIRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAnalysisAIResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the phone number analysis result.
//
// Description:
//
// Before using this operation, log on to the Cell Phone Number Service console, go to the Tag Square page, find the corresponding tag, click Apply to activate, and fill in the application materials. You can use the operation after the application is approved.
//
// Make sure that you fully understand the billing of the Cell Phone Number Service before you use this operation.
//
// @param request - DescribePhoneNumberAnalysisAIRequest
//
// @return DescribePhoneNumberAnalysisAIResponse
func (client *Client) DescribePhoneNumberAnalysisAI(request *DescribePhoneNumberAnalysisAIRequest) (_result *DescribePhoneNumberAnalysisAIResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisAIResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisAIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePhoneNumberAnalysisPaiWithOptions(request *DescribePhoneNumberAnalysisPaiRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAnalysisPaiResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePhoneNumberAnalysisPaiResponse
func (client *Client) DescribePhoneNumberAnalysisPai(request *DescribePhoneNumberAnalysisPaiRequest) (_result *DescribePhoneNumberAnalysisPaiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisPaiResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisPaiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePhoneNumberAnalysisTransparentWithOptions(request *DescribePhoneNumberAnalysisTransparentRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAnalysisTransparentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePhoneNumberAnalysisTransparentResponse
func (client *Client) DescribePhoneNumberAnalysisTransparent(request *DescribePhoneNumberAnalysisTransparentRequest) (_result *DescribePhoneNumberAnalysisTransparentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberAnalysisTransparentResponse{}
	_body, _err := client.DescribePhoneNumberAnalysisTransparentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribePhoneNumberAttribute is deprecated, please use Dytnsapi::2020-02-17::DescribePhoneNumberOperatorAttribute instead.
//
// @param request - DescribePhoneNumberAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberAttributeResponse
func (client *Client) DescribePhoneNumberAttributeWithOptions(request *DescribePhoneNumberAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribePhoneNumberAttribute is deprecated, please use Dytnsapi::2020-02-17::DescribePhoneNumberOperatorAttribute instead.
//
// @param request - DescribePhoneNumberAttributeRequest
//
// @return DescribePhoneNumberAttributeResponse
// Deprecated
func (client *Client) DescribePhoneNumberAttribute(request *DescribePhoneNumberAttributeRequest) (_result *DescribePhoneNumberAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberAttributeResponse{}
	_body, _err := client.DescribePhoneNumberAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the duration for which a mobile user has been registered on the carrier\\"s network.
//
// Description:
//
// - Before using this API, log on to the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, fill in the application materials, and use the API after the application is approved.
//
// - Make sure that you fully understand the [product pricing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service before using this API.
//
// ### QPS limit
//
// The per-user QPS limit for this API is 200 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this API at a reasonable rate.
//
// @param request - DescribePhoneNumberOnlineTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOnlineTimeResponse
func (client *Client) DescribePhoneNumberOnlineTimeWithOptions(request *DescribePhoneNumberOnlineTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberOnlineTimeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the duration for which a mobile user has been registered on the carrier\\"s network.
//
// Description:
//
// - Before using this API, log on to the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, fill in the application materials, and use the API after the application is approved.
//
// - Make sure that you fully understand the [product pricing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service before using this API.
//
// ### QPS limit
//
// The per-user QPS limit for this API is 200 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this API at a reasonable rate.
//
// @param request - DescribePhoneNumberOnlineTimeRequest
//
// @return DescribePhoneNumberOnlineTimeResponse
func (client *Client) DescribePhoneNumberOnlineTime(request *DescribePhoneNumberOnlineTimeRequest) (_result *DescribePhoneNumberOnlineTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberOnlineTimeResponse{}
	_body, _err := client.DescribePhoneNumberOnlineTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the carrier information of a phone number, including the province, city, basic carrier (China Mobile, China Unicom, China Telecom, or China Broadnet), mobile virtual network operator (such as Alibaba Cloud Communication), whether the number has been ported, and the number segment.
//
// Description:
//
// - Before you use this API, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154008.html) of Cell Phone Number Service.
//
// - By default, only Alibaba Cloud accounts can use this API. RAM users can use this API only after they are granted the required permissions. For more information, see [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/154006.html).
//
// - This API is used to obtain the current carrier, location, and number portability information of a phone number. The API supports queries for **plaintext**, **MD5**, and **SHA256*	- encrypted phone numbers.
//
// - Before you use this API, log on to the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, and submit the application materials. You can use the API after your application is approved.
//
// ### QPS limit
//
// The QPS limit per user for this API is 2,000 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make calls reasonably.
//
// @param request - DescribePhoneNumberOperatorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneNumberOperatorAttributeResponse
func (client *Client) DescribePhoneNumberOperatorAttributeWithOptions(request *DescribePhoneNumberOperatorAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the carrier information of a phone number, including the province, city, basic carrier (China Mobile, China Unicom, China Telecom, or China Broadnet), mobile virtual network operator (such as Alibaba Cloud Communication), whether the number has been ported, and the number segment.
//
// Description:
//
// - Before you use this API, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/154008.html) of Cell Phone Number Service.
//
// - By default, only Alibaba Cloud accounts can use this API. RAM users can use this API only after they are granted the required permissions. For more information, see [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/154006.html).
//
// - This API is used to obtain the current carrier, location, and number portability information of a phone number. The API supports queries for **plaintext**, **MD5**, and **SHA256*	- encrypted phone numbers.
//
// - Before you use this API, log on to the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, and submit the application materials. You can use the API after your application is approved.
//
// ### QPS limit
//
// The QPS limit per user for this API is 2,000 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make calls reasonably.
//
// @param request - DescribePhoneNumberOperatorAttributeRequest
//
// @return DescribePhoneNumberOperatorAttributeResponse
func (client *Client) DescribePhoneNumberOperatorAttribute(request *DescribePhoneNumberOperatorAttributeRequest) (_result *DescribePhoneNumberOperatorAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberOperatorAttributeResponse{}
	_body, _err := client.DescribePhoneNumberOperatorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePhoneNumberOperatorAttributeAnnualWithOptions(request *DescribePhoneNumberOperatorAttributeAnnualRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeAnnualResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePhoneNumberOperatorAttributeAnnualResponse
func (client *Client) DescribePhoneNumberOperatorAttributeAnnual(request *DescribePhoneNumberOperatorAttributeAnnualRequest) (_result *DescribePhoneNumberOperatorAttributeAnnualResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberOperatorAttributeAnnualResponse{}
	_body, _err := client.DescribePhoneNumberOperatorAttributeAnnualWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePhoneNumberOperatorAttributeAnnualUseWithOptions(request *DescribePhoneNumberOperatorAttributeAnnualUseRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberOperatorAttributeAnnualUseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePhoneNumberOperatorAttributeAnnualUseResponse
func (client *Client) DescribePhoneNumberOperatorAttributeAnnualUse(request *DescribePhoneNumberOperatorAttributeAnnualUseRequest) (_result *DescribePhoneNumberOperatorAttributeAnnualUseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberOperatorAttributeAnnualUseResponse{}
	_body, _err := client.DescribePhoneNumberOperatorAttributeAnnualUseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePhoneNumberRiskWithOptions(request *DescribePhoneNumberRiskRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneNumberRiskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePhoneNumberRiskResponse
func (client *Client) DescribePhoneNumberRisk(request *DescribePhoneNumberRiskRequest) (_result *DescribePhoneNumberRiskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneNumberRiskResponse{}
	_body, _err := client.DescribePhoneNumberRiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Quickly verify in a single request if a mobile phone number is a carrier\\"s secondary number assignment.
//
// Description:
//
// - Before you use this API, make sure you understand the [pricing](https://help.aliyun.com/document_detail/154751.html) for Phone Number Verification Service.
//
// - You are charged only when the API response returns `Code="OK"` and `VerifyResult` is not `0`. No charge is incurred in any other case.
//
// - Before you use this API, log in to the Phone Number Verification Service console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the tag you need, click **Request Activation**, and complete the application form. You can use the API once your application is approved.
//
// ## QPS limits
//
// The QPS limit is 100 per user. If you exceed this limit, the system throttles your API calls, which may affect your business. Plan your API calls accordingly.
//
// ## Authorization
//
// By default, only an Alibaba Cloud account can call this API. A RAM user must be granted the required permissions to call the API. For more information, see [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribePhoneTwiceTelVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePhoneTwiceTelVerifyResponse
func (client *Client) DescribePhoneTwiceTelVerifyWithOptions(request *DescribePhoneTwiceTelVerifyRequest, runtime *dara.RuntimeOptions) (_result *DescribePhoneTwiceTelVerifyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Quickly verify in a single request if a mobile phone number is a carrier\\"s secondary number assignment.
//
// Description:
//
// - Before you use this API, make sure you understand the [pricing](https://help.aliyun.com/document_detail/154751.html) for Phone Number Verification Service.
//
// - You are charged only when the API response returns `Code="OK"` and `VerifyResult` is not `0`. No charge is incurred in any other case.
//
// - Before you use this API, log in to the Phone Number Verification Service console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the tag you need, click **Request Activation**, and complete the application form. You can use the API once your application is approved.
//
// ## QPS limits
//
// The QPS limit is 100 per user. If you exceed this limit, the system throttles your API calls, which may affect your business. Plan your API calls accordingly.
//
// ## Authorization
//
// By default, only an Alibaba Cloud account can call this API. A RAM user must be granted the required permissions to call the API. For more information, see [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/154006.html).
//
// @param request - DescribePhoneTwiceTelVerifyRequest
//
// @return DescribePhoneTwiceTelVerifyResponse
func (client *Client) DescribePhoneTwiceTelVerify(request *DescribePhoneTwiceTelVerifyRequest) (_result *DescribePhoneTwiceTelVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePhoneTwiceTelVerifyResponse{}
	_body, _err := client.DescribePhoneTwiceTelVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API only provides the signing functionality for acquiring an authorization token during UAID retrieval from the three-network operators.
//
// Description:
//
// This API provides only the **signing function for obtaining an authorization token*	- as part of the UAID retrieval process for China\\"s three major operators.
//
// - To generate the other parameters required to obtain the authorization token, refer to the respective operator\\"s documentation.
//
// - To request the authorization token, refer to the respective operator\\"s documentation. Note that the request must be sent from the user\\"s mobile device over a cellular network.
//
// - After obtaining the authorization token, call the [UAIDVerification](~~UAIDVerification~~) API to complete the UAID retrieval.
//
// ### How to request an authorization token
//
//	Notice:
//
// When you request an authorization token, the client or H5 page must be connected to a cellular network.
//
// #### China Mobile (CM)
//
// Protocol: HTTPS + application/json
//
// Method: POST
//
// URL: https\\://msg.cmpassport.com/h5/getMobile
//
// ##### Request parameters
//
// Request example:
//
// `{ "traceId": "mfawsxtcmyplwzpayzzvdvbsowxmkynr", "appId": "300011580392", "sign": "2c61b3c58ffbeed97461e31be4fd931a", "msgId": "redbyxsdetddwaaffajcwwapspykftzx", "expandParams": "", "businessType": "3", "version": "1.0", "timestamp": "20201125101540980" }`
//
// Parameter description:
//
// - `version`: Use `1.0`.
//
// - `timestamp`: The request timestamp, with millisecond precision. This value and its format must match the `Time` input parameter for this API.
//
// - `appId`: Use `300011580392`.
//
// - `businessType`: Use `3`.
//
// - `traceId`: The trace ID. This value must match this API\\"s `OutId` input parameter.
//
// - `sign`: Obtained by calling this API.
//
// - `msgId`: A unique message identifier.
//
// ##### Response parameters
//
// Response example:
//
// `{ "header": { "appId": "300011580392", "msgId": "redbyxsdetddwaaffajcwwapspykftzx", "timestamp": "20201125101607932" }, "body": { "resultCode": "103000", "expandParams": "", "resultDesc": "成功", "token": "H5HTTPS4187AE9743AFCB14F8D99B9D65ED9E01" } }`
//
// Retrieve the `token` from the response `body`.
//
// #### China Unicom (CU)
//
// Obtain the token in two steps.
//
// ##### Step 1: Obtain the authurl
//
// Send a request to the portal server to get the authentication server address (`authurl`).
//
// Protocol: HTTPS + application/json
//
// Method: GET
//
// URL: https\\://nisportal.10010.com:9001/api
//
// ###### Request parameters
//
// Request examples:
//
// JSON request: `?appid=1554778161153`
//
// JSONP request: `?appid=1554778161153&callback=callbackFunction`
//
// Parameter description:
//
// - `appid`: Use 1554778161153.
//
// - `callback`: The name of the JSONP callback function. This parameter is required only for JSONP requests.
//
// ###### Response parameters
//
// Response examples:
//
// JSON response:
//
// `{"authurl": "https://enrichgw.10010.com/d93222629f52ec79"}`
//
// JSONP response:
//
// `callbackFunction({"authurl":"https://enrichgw.10010.com/d93222629f52ec79"})`
//
// Retrieve the `authurl` from the response.
//
// ##### Step 2: Obtain the token
//
// Protocol: HTTPS + application/json
//
// Method: GET
//
// URL: The authurl from Step 1, with /api appended.
//
// Request URL example: `https://enrichgw.10010.com/d93222629f52ec79/api`
//
// ###### Request parameters
//
// Request examples:
//
// JSON request: `?appid=1554778161153`
//
// JSONP request: `?appid=1554778161153&callback=callbackFunction`
//
// Parameter description:
//
// - `appid`: Use 1554778161153.
//
// - `callback`: The name of the JSONP callback function. This parameter is required only for JSONP requests.
//
// ###### Response parameters
//
// Response examples:
//
// JSON response:
//
// `{ "province": "1", "code": "7nHS1nggx2WP613750206700RN6oiRN1" }`
//
// JSONP response:
//
// `callbackFunction({"province":"1","code":"7nHS3Dnkd1BS701851092400RN6oiRN1"})`
//
// Retrieve the `code` from the response.
//
// #### China Telecom (CT)
//
// Protocol: HTTPS + application/x-www-form-urlencoded;charset=UTF-8
//
// Method: GET
//
// URL: https\\://id6.me/gw/preuniq.do
//
// ##### Request parameters
//
// Request example:
//
// `?clientType=30100&appId=9390188202&format=json&sign=D63C166FA19E1996EF********09C6A5397C10B4&paramKey=1D7C25EB8B0B8B4CB3CF8DC60628F6549********786B0AF1FEF93FA1335057A35BF5F0B39A3867EAA9BE14B3898********8B01DE34965060445B6E1F66401D714650E4AB161CD6DCF4A72********3B856F22A192B8B0C39D7A55B961062E68C89C928894F119B25********7C548355FE9DB82852EB93C939F2200B48CD17&paramStr=140********95AF8E138B94754CB4CF83BA6FB********52B258BFDFD38BF233&version=1.1`
//
// Parameter description:
//
// - `appId`: Use `9390188202`.
//
// - `clientType`: The client type. This value must match the `ClientType` input parameter for this API.
//
// - `format`: Use `json` or `jsonp`.
//
// - `version`: Use `1.1`.
//
// - `sign`: Obtained by calling this API.
//
// - `paramKey`: The ciphertext of key A. Key A is a 16-character random string generated by the client. To generate paramKey, encrypt key A using the RSA algorithm and the China Telecom public key. The padding mode is `RSA/ECB/PKCS1Padding`. Download the [China Telecom RSA public key](https://id.189.cn/source/files/API.pem).
//
// - `paramStr`: The ciphertext of a parameter string. This string contains `timeStamp` (a Unix timestamp with millisecond precision, for example `1697791988302`, that corresponds to the value of the `Time` parameter of this API) and `callback` (the name of the JSONP callback function, required only when `format` is set to `jsonp`). To generate `paramStr`, encrypt the string using AES with key A. The padding mode is `AES/CBC/PKCS5Padding`, and the initialization vector is `0000000000000000`.
//
// ##### Response parameters
//
// Response example:
//
// `callback?result=10000&msg=success&data=a35336711c70456cb883f4f224e9a259`
//
// The `data` parameter contains the ciphertext of the business result. To get the result, decrypt the data value using key A, the `AES/CBC/PKCS5Padding` mode, and an initialization vector of `0000000000000000`.
//
// Decrypted business result example:
//
// `{"accessCode": "H5HTTPS4187AE9743AFCB14F8D99B9D65ED9E01"}`
//
// Retrieve the `accessCode` from the decrypted result.
//
// @param request - GetUAIDApplyTokenSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUAIDApplyTokenSignResponse
func (client *Client) GetUAIDApplyTokenSignWithOptions(request *GetUAIDApplyTokenSignRequest, runtime *dara.RuntimeOptions) (_result *GetUAIDApplyTokenSignResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API only provides the signing functionality for acquiring an authorization token during UAID retrieval from the three-network operators.
//
// Description:
//
// This API provides only the **signing function for obtaining an authorization token*	- as part of the UAID retrieval process for China\\"s three major operators.
//
// - To generate the other parameters required to obtain the authorization token, refer to the respective operator\\"s documentation.
//
// - To request the authorization token, refer to the respective operator\\"s documentation. Note that the request must be sent from the user\\"s mobile device over a cellular network.
//
// - After obtaining the authorization token, call the [UAIDVerification](~~UAIDVerification~~) API to complete the UAID retrieval.
//
// ### How to request an authorization token
//
//	Notice:
//
// When you request an authorization token, the client or H5 page must be connected to a cellular network.
//
// #### China Mobile (CM)
//
// Protocol: HTTPS + application/json
//
// Method: POST
//
// URL: https\\://msg.cmpassport.com/h5/getMobile
//
// ##### Request parameters
//
// Request example:
//
// `{ "traceId": "mfawsxtcmyplwzpayzzvdvbsowxmkynr", "appId": "300011580392", "sign": "2c61b3c58ffbeed97461e31be4fd931a", "msgId": "redbyxsdetddwaaffajcwwapspykftzx", "expandParams": "", "businessType": "3", "version": "1.0", "timestamp": "20201125101540980" }`
//
// Parameter description:
//
// - `version`: Use `1.0`.
//
// - `timestamp`: The request timestamp, with millisecond precision. This value and its format must match the `Time` input parameter for this API.
//
// - `appId`: Use `300011580392`.
//
// - `businessType`: Use `3`.
//
// - `traceId`: The trace ID. This value must match this API\\"s `OutId` input parameter.
//
// - `sign`: Obtained by calling this API.
//
// - `msgId`: A unique message identifier.
//
// ##### Response parameters
//
// Response example:
//
// `{ "header": { "appId": "300011580392", "msgId": "redbyxsdetddwaaffajcwwapspykftzx", "timestamp": "20201125101607932" }, "body": { "resultCode": "103000", "expandParams": "", "resultDesc": "成功", "token": "H5HTTPS4187AE9743AFCB14F8D99B9D65ED9E01" } }`
//
// Retrieve the `token` from the response `body`.
//
// #### China Unicom (CU)
//
// Obtain the token in two steps.
//
// ##### Step 1: Obtain the authurl
//
// Send a request to the portal server to get the authentication server address (`authurl`).
//
// Protocol: HTTPS + application/json
//
// Method: GET
//
// URL: https\\://nisportal.10010.com:9001/api
//
// ###### Request parameters
//
// Request examples:
//
// JSON request: `?appid=1554778161153`
//
// JSONP request: `?appid=1554778161153&callback=callbackFunction`
//
// Parameter description:
//
// - `appid`: Use 1554778161153.
//
// - `callback`: The name of the JSONP callback function. This parameter is required only for JSONP requests.
//
// ###### Response parameters
//
// Response examples:
//
// JSON response:
//
// `{"authurl": "https://enrichgw.10010.com/d93222629f52ec79"}`
//
// JSONP response:
//
// `callbackFunction({"authurl":"https://enrichgw.10010.com/d93222629f52ec79"})`
//
// Retrieve the `authurl` from the response.
//
// ##### Step 2: Obtain the token
//
// Protocol: HTTPS + application/json
//
// Method: GET
//
// URL: The authurl from Step 1, with /api appended.
//
// Request URL example: `https://enrichgw.10010.com/d93222629f52ec79/api`
//
// ###### Request parameters
//
// Request examples:
//
// JSON request: `?appid=1554778161153`
//
// JSONP request: `?appid=1554778161153&callback=callbackFunction`
//
// Parameter description:
//
// - `appid`: Use 1554778161153.
//
// - `callback`: The name of the JSONP callback function. This parameter is required only for JSONP requests.
//
// ###### Response parameters
//
// Response examples:
//
// JSON response:
//
// `{ "province": "1", "code": "7nHS1nggx2WP613750206700RN6oiRN1" }`
//
// JSONP response:
//
// `callbackFunction({"province":"1","code":"7nHS3Dnkd1BS701851092400RN6oiRN1"})`
//
// Retrieve the `code` from the response.
//
// #### China Telecom (CT)
//
// Protocol: HTTPS + application/x-www-form-urlencoded;charset=UTF-8
//
// Method: GET
//
// URL: https\\://id6.me/gw/preuniq.do
//
// ##### Request parameters
//
// Request example:
//
// `?clientType=30100&appId=9390188202&format=json&sign=D63C166FA19E1996EF********09C6A5397C10B4&paramKey=1D7C25EB8B0B8B4CB3CF8DC60628F6549********786B0AF1FEF93FA1335057A35BF5F0B39A3867EAA9BE14B3898********8B01DE34965060445B6E1F66401D714650E4AB161CD6DCF4A72********3B856F22A192B8B0C39D7A55B961062E68C89C928894F119B25********7C548355FE9DB82852EB93C939F2200B48CD17&paramStr=140********95AF8E138B94754CB4CF83BA6FB********52B258BFDFD38BF233&version=1.1`
//
// Parameter description:
//
// - `appId`: Use `9390188202`.
//
// - `clientType`: The client type. This value must match the `ClientType` input parameter for this API.
//
// - `format`: Use `json` or `jsonp`.
//
// - `version`: Use `1.1`.
//
// - `sign`: Obtained by calling this API.
//
// - `paramKey`: The ciphertext of key A. Key A is a 16-character random string generated by the client. To generate paramKey, encrypt key A using the RSA algorithm and the China Telecom public key. The padding mode is `RSA/ECB/PKCS1Padding`. Download the [China Telecom RSA public key](https://id.189.cn/source/files/API.pem).
//
// - `paramStr`: The ciphertext of a parameter string. This string contains `timeStamp` (a Unix timestamp with millisecond precision, for example `1697791988302`, that corresponds to the value of the `Time` parameter of this API) and `callback` (the name of the JSONP callback function, required only when `format` is set to `jsonp`). To generate `paramStr`, encrypt the string using AES with key A. The padding mode is `AES/CBC/PKCS5Padding`, and the initialization vector is `0000000000000000`.
//
// ##### Response parameters
//
// Response example:
//
// `callback?result=10000&msg=success&data=a35336711c70456cb883f4f224e9a259`
//
// The `data` parameter contains the ciphertext of the business result. To get the result, decrypt the data value using key A, the `AES/CBC/PKCS5Padding` mode, and an initialization vector of `0000000000000000`.
//
// Decrypted business result example:
//
// `{"accessCode": "H5HTTPS4187AE9743AFCB14F8D99B9D65ED9E01"}`
//
// Retrieve the `accessCode` from the decrypted result.
//
// @param request - GetUAIDApplyTokenSignRequest
//
// @return GetUAIDApplyTokenSignResponse
func (client *Client) GetUAIDApplyTokenSign(request *GetUAIDApplyTokenSignRequest) (_result *GetUAIDApplyTokenSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUAIDApplyTokenSignResponse{}
	_body, _err := client.GetUAIDApplyTokenSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUAIDConversionSignWithOptions(request *GetUAIDConversionSignRequest, runtime *dara.RuntimeOptions) (_result *GetUAIDConversionSignResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetUAIDConversionSignResponse
func (client *Client) GetUAIDConversionSign(request *GetUAIDConversionSignRequest) (_result *GetUAIDConversionSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUAIDConversionSignResponse{}
	_body, _err := client.GetUAIDConversionSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies whether a name, phone number, and ID number belonged to the same user at a specific point in time.
//
// Description:
//
// ## Usage notes
//
// - Before you use this API, ensure you understand its [pricing](https://help.aliyun.com/zh/cpns/product-overview/pricing-of-cpns).
//
// - Before you use this API, log on to the Phone Number Intelligence console and apply for the required tag on the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page.
//
// - You are charged for a call only when the response returns `Code=\\"OK\\"` and `IsConsistent != 0`. No charges are incurred for any other results.
//
// - Verification of China Broadcasting Network numbers (numbers with the 192 prefix) is not supported. If you provide a number with the 192 prefix, an HTTP 400 error is returned.
//
// - Due to number portability, the actual carrier may differ from the current carrier of record. You can use the `Carrier` parameter to route the query to a specific carrier.
//
// - The queries per second (QPS) limit per user is 200. Requests that exceed this limit are throttled.
//
// - In the authorization information, the action is `dytns:HistoryThreeElementsVerification`, the access level is Read, and the resource type is All Resources.
//
// @param request - HistoryThreeElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HistoryThreeElementsVerificationResponse
func (client *Client) HistoryThreeElementsVerificationWithOptions(request *HistoryThreeElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *HistoryThreeElementsVerificationResponse, _err error) {
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

	if !dara.IsNil(request.VerificationTime) {
		query["VerificationTime"] = request.VerificationTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HistoryThreeElementsVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HistoryThreeElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether a name, phone number, and ID number belonged to the same user at a specific point in time.
//
// Description:
//
// ## Usage notes
//
// - Before you use this API, ensure you understand its [pricing](https://help.aliyun.com/zh/cpns/product-overview/pricing-of-cpns).
//
// - Before you use this API, log on to the Phone Number Intelligence console and apply for the required tag on the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page.
//
// - You are charged for a call only when the response returns `Code=\\"OK\\"` and `IsConsistent != 0`. No charges are incurred for any other results.
//
// - Verification of China Broadcasting Network numbers (numbers with the 192 prefix) is not supported. If you provide a number with the 192 prefix, an HTTP 400 error is returned.
//
// - Due to number portability, the actual carrier may differ from the current carrier of record. You can use the `Carrier` parameter to route the query to a specific carrier.
//
// - The queries per second (QPS) limit per user is 200. Requests that exceed this limit are throttled.
//
// - In the authorization information, the action is `dytns:HistoryThreeElementsVerification`, the access level is Read, and the resource type is All Resources.
//
// @param request - HistoryThreeElementsVerificationRequest
//
// @return HistoryThreeElementsVerificationResponse
func (client *Client) HistoryThreeElementsVerification(request *HistoryThreeElementsVerificationRequest) (_result *HistoryThreeElementsVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &HistoryThreeElementsVerificationResponse{}
	_body, _err := client.HistoryThreeElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs invalid phone number filtering operations.
//
// Description:
//
// Before you call this operation, log on to the Cell Phone Number Service console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, and submit the application materials. You can use this operation after the application is approved.
//
// ### QPS limit
//
// The QPS limit per user for this operation is 1,000 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the operation at a reasonable frequency.
//
// @param request - InvalidPhoneNumberFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidPhoneNumberFilterResponse
func (client *Client) InvalidPhoneNumberFilterWithOptions(request *InvalidPhoneNumberFilterRequest, runtime *dara.RuntimeOptions) (_result *InvalidPhoneNumberFilterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs invalid phone number filtering operations.
//
// Description:
//
// Before you call this operation, log on to the Cell Phone Number Service console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, and submit the application materials. You can use this operation after the application is approved.
//
// ### QPS limit
//
// The QPS limit per user for this operation is 1,000 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the operation at a reasonable frequency.
//
// @param request - InvalidPhoneNumberFilterRequest
//
// @return InvalidPhoneNumberFilterResponse
func (client *Client) InvalidPhoneNumberFilter(request *InvalidPhoneNumberFilterRequest) (_result *InvalidPhoneNumberFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InvalidPhoneNumberFilterResponse{}
	_body, _err := client.InvalidPhoneNumberFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the alert contacts for Number Encyclopedia.
//
// Description:
//
// - Before using this API, ensure that you have activated Number Encyclopedia.
//
// @param request - ListContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContactsResponse
func (client *Client) ListContactsWithOptions(request *ListContactsRequest, runtime *dara.RuntimeOptions) (_result *ListContactsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the alert contacts for Number Encyclopedia.
//
// Description:
//
// - Before using this API, ensure that you have activated Number Encyclopedia.
//
// @param request - ListContactsRequest
//
// @return ListContactsResponse
func (client *Client) ListContacts(request *ListContactsRequest) (_result *ListContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListContactsResponse{}
	_body, _err := client.ListContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PhoneNumberConvertServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberConvertServiceResponse
func (client *Client) PhoneNumberConvertServiceWithOptions(request *PhoneNumberConvertServiceRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberConvertServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PhoneNumberConvertServiceRequest
//
// @return PhoneNumberConvertServiceResponse
func (client *Client) PhoneNumberConvertService(request *PhoneNumberConvertServiceRequest) (_result *PhoneNumberConvertServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PhoneNumberConvertServiceResponse{}
	_body, _err := client.PhoneNumberConvertServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Encrypts an original phone number into a virtual phone number that starts with 140. By integrating with Alibaba Cloud communication services, you can use the encrypted 140 phone number to initiate voice calls, achieving the effect of a virtual phone number call.
//
// Description:
//
// Before you use this API, log on to the Cell Phone Number Service console, go to the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, and fill in the application materials. You can use this API after the application is approved.
//
// ### QPS limit
//
// The per-user QPS limit of this API is 1,000 calls per second. If the number of API calls exceeds the limit, the calls will be throttled, which may affect your business. Call the API properly.
//
// @param request - PhoneNumberEncryptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberEncryptResponse
func (client *Client) PhoneNumberEncryptWithOptions(request *PhoneNumberEncryptRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberEncryptResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Encrypts an original phone number into a virtual phone number that starts with 140. By integrating with Alibaba Cloud communication services, you can use the encrypted 140 phone number to initiate voice calls, achieving the effect of a virtual phone number call.
//
// Description:
//
// Before you use this API, log on to the Cell Phone Number Service console, go to the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, and fill in the application materials. You can use this API after the application is approved.
//
// ### QPS limit
//
// The per-user QPS limit of this API is 1,000 calls per second. If the number of API calls exceeds the limit, the calls will be throttled, which may affect your business. Call the API properly.
//
// @param request - PhoneNumberEncryptRequest
//
// @return PhoneNumberEncryptResponse
func (client *Client) PhoneNumberEncrypt(request *PhoneNumberEncryptRequest) (_result *PhoneNumberEncryptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PhoneNumberEncryptResponse{}
	_body, _err := client.PhoneNumberEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the real-time network status of a mobile phone number, such as active, shutdown, or non-existent. You can query numbers that are in plaintext or hashed using MD5 or SHA256.
//
// Description:
//
// - **Before you use this operation, make sure that you fully understand the [pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Intelligence.**
//
// - By default, only an Alibaba Cloud account can call this operation. A RAM user can call this operation only after receiving the required permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// - Before you use this operation, log on to the Phone Number Intelligence console. On the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Apply**, and then submit the required information. You can use this operation once your application is approved.
//
// - The phone number status query feature supports numbers from China Telecom, China Unicom, and China Mobile. This feature does not support numbers from China Broadnet. If you call this operation to query a China Broadnet number, the API returns the error code `OperatorLimit`, which indicates that the query is prohibited by the carrier.
//
// ### QPS limit
//
// This operation has a queries per second (QPS) limit of 300 per user. If you exceed this limit, your API calls are throttled, which may affect your services. We recommend that you call this operation at a reasonable frequency.
//
// @param request - PhoneNumberStatusForAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForAccountResponse
func (client *Client) PhoneNumberStatusForAccountWithOptions(request *PhoneNumberStatusForAccountRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForAccountResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the real-time network status of a mobile phone number, such as active, shutdown, or non-existent. You can query numbers that are in plaintext or hashed using MD5 or SHA256.
//
// Description:
//
// - **Before you use this operation, make sure that you fully understand the [pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Intelligence.**
//
// - By default, only an Alibaba Cloud account can call this operation. A RAM user can call this operation only after receiving the required permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// - Before you use this operation, log on to the Phone Number Intelligence console. On the [Tag Plaza](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Apply**, and then submit the required information. You can use this operation once your application is approved.
//
// - The phone number status query feature supports numbers from China Telecom, China Unicom, and China Mobile. This feature does not support numbers from China Broadnet. If you call this operation to query a China Broadnet number, the API returns the error code `OperatorLimit`, which indicates that the query is prohibited by the carrier.
//
// ### QPS limit
//
// This operation has a queries per second (QPS) limit of 300 per user. If you exceed this limit, your API calls are throttled, which may affect your services. We recommend that you call this operation at a reasonable frequency.
//
// @param request - PhoneNumberStatusForAccountRequest
//
// @return PhoneNumberStatusForAccountResponse
func (client *Client) PhoneNumberStatusForAccount(request *PhoneNumberStatusForAccountRequest) (_result *PhoneNumberStatusForAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PhoneNumberStatusForAccountResponse{}
	_body, _err := client.PhoneNumberStatusForAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time status of a mobile phone number to determine whether it is in service, suspended, or non-existent. This operation supports queries for plaintext numbers or numbers encrypted with MD5 or SHA256.
//
// Description:
//
// - **Before you call this operation, make sure that you fully understand the [pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Intelligence.**
//
// - By default, only an Alibaba Cloud account can call this operation. A RAM user must be granted the required permissions before calling this operation. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// - Before calling this operation, log on to the Phone Number Intelligence console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Apply**, and submit your application. You can use the operation after your application is approved.
//
// - The number status query feature supports numbers from China Telecom, China Unicom, and China Mobile, but does not support numbers from China Broadnet. If you call this operation to query the status of a China Broadnet number, the `OperatorLimit` error code is returned, which indicates that the query is prohibited by the carrier.
//
// ### QPS limit
//
// The queries per second (QPS) limit for each user is 300. API calls that exceed this limit are throttled. To avoid business disruptions, plan your calls accordingly.
//
// @param request - PhoneNumberStatusForPublicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForPublicResponse
func (client *Client) PhoneNumberStatusForPublicWithOptions(request *PhoneNumberStatusForPublicRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForPublicResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time status of a mobile phone number to determine whether it is in service, suspended, or non-existent. This operation supports queries for plaintext numbers or numbers encrypted with MD5 or SHA256.
//
// Description:
//
// - **Before you call this operation, make sure that you fully understand the [pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Intelligence.**
//
// - By default, only an Alibaba Cloud account can call this operation. A RAM user must be granted the required permissions before calling this operation. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// - Before calling this operation, log on to the Phone Number Intelligence console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Apply**, and submit your application. You can use the operation after your application is approved.
//
// - The number status query feature supports numbers from China Telecom, China Unicom, and China Mobile, but does not support numbers from China Broadnet. If you call this operation to query the status of a China Broadnet number, the `OperatorLimit` error code is returned, which indicates that the query is prohibited by the carrier.
//
// ### QPS limit
//
// The queries per second (QPS) limit for each user is 300. API calls that exceed this limit are throttled. To avoid business disruptions, plan your calls accordingly.
//
// @param request - PhoneNumberStatusForPublicRequest
//
// @return PhoneNumberStatusForPublicResponse
func (client *Client) PhoneNumberStatusForPublic(request *PhoneNumberStatusForPublicRequest) (_result *PhoneNumberStatusForPublicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PhoneNumberStatusForPublicResponse{}
	_body, _err := client.PhoneNumberStatusForPublicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time status of a mobile phone number, such as normal, suspended, or not in service. This operation supports queries for phone numbers that are in plaintext or encrypted by using MD5, SHA256, or SM3.
//
// Description:
//
// - **Before calling this operation, ensure you fully understand the [pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Intelligence.**
//
// - By default, only an Alibaba Cloud account can call this operation. To allow a RAM user to call this operation, you must first grant the required permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// - Before you call this operation, log on to the Phone Number Intelligence console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Request Activation**, and then submit your application. You can call this operation only after your application is approved.
//
// - This operation supports phone numbers from China Telecom, China Unicom, and China Mobile. Numbers from China Broadnet are not supported. If you call this operation to query a China Broadnet number, the API returns the error code `OperatorLimit` and an error message indicating that the query is restricted by the carrier.
//
// ### QPS limit
//
// The QPS limit for this operation is 300 queries per second (QPS) per user. The system throttles calls that exceed this limit, which may affect your business. Plan your calls accordingly.
//
// @param request - PhoneNumberStatusForRealRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForRealResponse
func (client *Client) PhoneNumberStatusForRealWithOptions(request *PhoneNumberStatusForRealRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForRealResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time status of a mobile phone number, such as normal, suspended, or not in service. This operation supports queries for phone numbers that are in plaintext or encrypted by using MD5, SHA256, or SM3.
//
// Description:
//
// - **Before calling this operation, ensure you fully understand the [pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Intelligence.**
//
// - By default, only an Alibaba Cloud account can call this operation. To allow a RAM user to call this operation, you must first grant the required permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// - Before you call this operation, log on to the Phone Number Intelligence console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Request Activation**, and then submit your application. You can call this operation only after your application is approved.
//
// - This operation supports phone numbers from China Telecom, China Unicom, and China Mobile. Numbers from China Broadnet are not supported. If you call this operation to query a China Broadnet number, the API returns the error code `OperatorLimit` and an error message indicating that the query is restricted by the carrier.
//
// ### QPS limit
//
// The QPS limit for this operation is 300 queries per second (QPS) per user. The system throttles calls that exceed this limit, which may affect your business. Plan your calls accordingly.
//
// @param request - PhoneNumberStatusForRealRequest
//
// @return PhoneNumberStatusForRealResponse
func (client *Client) PhoneNumberStatusForReal(request *PhoneNumberStatusForRealRequest) (_result *PhoneNumberStatusForRealResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PhoneNumberStatusForRealResponse{}
	_body, _err := client.PhoneNumberStatusForRealWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the real-time network status of a phone number, such as active, shutdown, or non-existent. This operation supports queries for numbers in plaintext or numbers encrypted by using MD5, SHA256, or SM3.
//
// Description:
//
// - **Before you call this operation, make sure you understand the [Product Pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Pedia**.
//
// - By default, only an Alibaba Cloud account can call this operation. A RAM user can call this operation only after being granted the required permissions. For more information, see [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/154006.html).
//
// - Before calling this operation, log on to the Phone Number Pedia console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Request Activation**, and then complete the application form. You can use this operation only after your application is approved.
//
// - This feature supports phone numbers from China Telecom, China Unicom, and China Mobile, but does not support phone numbers from China Broadnet. If you call this operation to query the status of a China Broadnet number, the `OperatorLimit` error code and the "The number is limited by the operator." message are returned.
//
// ### QPS limit
//
// This operation is limited to 300 queries per second (QPS) for each user. Calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - PhoneNumberStatusForSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForSmsResponse
func (client *Client) PhoneNumberStatusForSmsWithOptions(request *PhoneNumberStatusForSmsRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForSmsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the real-time network status of a phone number, such as active, shutdown, or non-existent. This operation supports queries for numbers in plaintext or numbers encrypted by using MD5, SHA256, or SM3.
//
// Description:
//
// - **Before you call this operation, make sure you understand the [Product Pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Pedia**.
//
// - By default, only an Alibaba Cloud account can call this operation. A RAM user can call this operation only after being granted the required permissions. For more information, see [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/154006.html).
//
// - Before calling this operation, log on to the Phone Number Pedia console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Request Activation**, and then complete the application form. You can use this operation only after your application is approved.
//
// - This feature supports phone numbers from China Telecom, China Unicom, and China Mobile, but does not support phone numbers from China Broadnet. If you call this operation to query the status of a China Broadnet number, the `OperatorLimit` error code and the "The number is limited by the operator." message are returned.
//
// ### QPS limit
//
// This operation is limited to 300 queries per second (QPS) for each user. Calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - PhoneNumberStatusForSmsRequest
//
// @return PhoneNumberStatusForSmsResponse
func (client *Client) PhoneNumberStatusForSms(request *PhoneNumberStatusForSmsRequest) (_result *PhoneNumberStatusForSmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PhoneNumberStatusForSmsResponse{}
	_body, _err := client.PhoneNumberStatusForSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time network status of a mobile phone number, such as normal, shutdown, or non-existent. This operation supports queries for numbers in plaintext and numbers encrypted by using MD5, SHA256, or SM3.
//
// Description:
//
// - **Before you use this API, make sure that you understand the [pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Encyclopedia**.
//
// - By default, only Alibaba Cloud accounts can call this API. To allow a RAM user to do so, you must grant them the required permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// - Before you use this API, log in to the Phone Number Encyclopedia console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Apply for Access**, and then complete the application form. You can call this API after your application is approved.
//
// - This feature supports phone numbers from China Telecom, China Unicom, and China Mobile, but not from China Broadnet. If you query a China Broadnet number, the `OperatorLimit` error code and an error message are returned: The number is limited by the operator.
//
// ### QPS limit
//
// The QPS limit for a single user is 300 queries per second. If you exceed this limit, the system throttles your API calls, which may impact your business. To avoid interruptions, call this API at a reasonable rate.
//
// @param request - PhoneNumberStatusForVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PhoneNumberStatusForVoiceResponse
func (client *Client) PhoneNumberStatusForVoiceWithOptions(request *PhoneNumberStatusForVoiceRequest, runtime *dara.RuntimeOptions) (_result *PhoneNumberStatusForVoiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time network status of a mobile phone number, such as normal, shutdown, or non-existent. This operation supports queries for numbers in plaintext and numbers encrypted by using MD5, SHA256, or SM3.
//
// Description:
//
// - **Before you use this API, make sure that you understand the [pricing](https://help.aliyun.com/document_detail/154751.html) of Phone Number Encyclopedia**.
//
// - By default, only Alibaba Cloud accounts can call this API. To allow a RAM user to do so, you must grant them the required permissions. For more information, see [Grant permissions to RAM users](https://help.aliyun.com/document_detail/154006.html).
//
// - Before you use this API, log in to the Phone Number Encyclopedia console. On the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the required tag, click **Apply for Access**, and then complete the application form. You can call this API after your application is approved.
//
// - This feature supports phone numbers from China Telecom, China Unicom, and China Mobile, but not from China Broadnet. If you query a China Broadnet number, the `OperatorLimit` error code and an error message are returned: The number is limited by the operator.
//
// ### QPS limit
//
// The QPS limit for a single user is 300 queries per second. If you exceed this limit, the system throttles your API calls, which may impact your business. To avoid interruptions, call this API at a reasonable rate.
//
// @param request - PhoneNumberStatusForVoiceRequest
//
// @return PhoneNumberStatusForVoiceResponse
func (client *Client) PhoneNumberStatusForVoice(request *PhoneNumberStatusForVoiceRequest) (_result *PhoneNumberStatusForVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PhoneNumberStatusForVoiceResponse{}
	_body, _err := client.PhoneNumberStatusForVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryAvailableAuthCodeWithOptions(request *QueryAvailableAuthCodeRequest, runtime *dara.RuntimeOptions) (_result *QueryAvailableAuthCodeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryAvailableAuthCodeResponse
func (client *Client) QueryAvailableAuthCode(request *QueryAvailableAuthCodeRequest) (_result *QueryAvailableAuthCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAvailableAuthCodeResponse{}
	_body, _err := client.QueryAvailableAuthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries package type information.
//
// @param request - QueryPackageTypeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPackageTypeInfoResponse
func (client *Client) QueryPackageTypeInfoWithOptions(request *QueryPackageTypeInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryPackageTypeInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries package type information.
//
// @param request - QueryPackageTypeInfoRequest
//
// @return QueryPackageTypeInfoResponse
func (client *Client) QueryPackageTypeInfo(request *QueryPackageTypeInfoRequest) (_result *QueryPackageTypeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPackageTypeInfoResponse{}
	_body, _err := client.QueryPackageTypeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPhoneNumberOnlineTimeWithOptions(request *QueryPhoneNumberOnlineTimeRequest, runtime *dara.RuntimeOptions) (_result *QueryPhoneNumberOnlineTimeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPhoneNumberOnlineTimeResponse
func (client *Client) QueryPhoneNumberOnlineTime(request *QueryPhoneNumberOnlineTimeRequest) (_result *QueryPhoneNumberOnlineTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPhoneNumberOnlineTimeResponse{}
	_body, _err := client.QueryPhoneNumberOnlineTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryPhoneTwiceTelVerifyWithOptions(request *QueryPhoneTwiceTelVerifyRequest, runtime *dara.RuntimeOptions) (_result *QueryPhoneTwiceTelVerifyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryPhoneTwiceTelVerifyResponse
func (client *Client) QueryPhoneTwiceTelVerify(request *QueryPhoneTwiceTelVerifyRequest) (_result *QueryPhoneTwiceTelVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPhoneTwiceTelVerifyResponse{}
	_body, _err := client.QueryPhoneTwiceTelVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryTagApplyRuleWithOptions(request *QueryTagApplyRuleRequest, runtime *dara.RuntimeOptions) (_result *QueryTagApplyRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryTagApplyRuleResponse
func (client *Client) QueryTagApplyRule(request *QueryTagApplyRuleRequest) (_result *QueryTagApplyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTagApplyRuleResponse{}
	_body, _err := client.QueryTagApplyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tag information.
//
// @param request - QueryTagInfoBySelectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagInfoBySelectionResponse
func (client *Client) QueryTagInfoBySelectionWithOptions(request *QueryTagInfoBySelectionRequest, runtime *dara.RuntimeOptions) (_result *QueryTagInfoBySelectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag information.
//
// @param request - QueryTagInfoBySelectionRequest
//
// @return QueryTagInfoBySelectionResponse
func (client *Client) QueryTagInfoBySelection(request *QueryTagInfoBySelectionRequest) (_result *QueryTagInfoBySelectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTagInfoBySelectionResponse{}
	_body, _err := client.QueryTagInfoBySelectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag list by page.
//
// @param request - QueryTagListPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagListPageResponse
func (client *Client) QueryTagListPageWithOptions(request *QueryTagListPageRequest, runtime *dara.RuntimeOptions) (_result *QueryTagListPageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag list by page.
//
// @param request - QueryTagListPageRequest
//
// @return QueryTagListPageResponse
func (client *Client) QueryTagListPage(request *QueryTagListPageRequest) (_result *QueryTagListPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTagListPageResponse{}
	_body, _err := client.QueryTagListPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of offline tasks.
//
// @param tmpReq - QueryTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskListResponse
func (client *Client) QueryTaskListWithOptions(tmpReq *QueryTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of offline tasks.
//
// @param request - QueryTaskListRequest
//
// @return QueryTaskListResponse
func (client *Client) QueryTaskList(request *QueryTaskListRequest) (_result *QueryTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTaskListResponse{}
	_body, _err := client.QueryTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries usage statistics by tag ID.
//
// @param request - QueryUsageStatisticsByTagIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUsageStatisticsByTagIdResponse
func (client *Client) QueryUsageStatisticsByTagIdWithOptions(request *QueryUsageStatisticsByTagIdRequest, runtime *dara.RuntimeOptions) (_result *QueryUsageStatisticsByTagIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries usage statistics by tag ID.
//
// @param request - QueryUsageStatisticsByTagIdRequest
//
// @return QueryUsageStatisticsByTagIdResponse
func (client *Client) QueryUsageStatisticsByTagId(request *QueryUsageStatisticsByTagIdRequest) (_result *QueryUsageStatisticsByTagIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUsageStatisticsByTagIdResponse{}
	_body, _err := client.QueryUsageStatisticsByTagIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Saves an alert contact.
//
// Description:
//
// - Before you call this operation, ensure you have activated Phone Number Intelligence.
//
// @param request - SaveContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveContactsResponse
func (client *Client) SaveContactsWithOptions(request *SaveContactsRequest, runtime *dara.RuntimeOptions) (_result *SaveContactsResponse, _err error) {
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

	if !dara.IsNil(request.ContactEmail) {
		query["ContactEmail"] = request.ContactEmail
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
		Action:      dara.String("SaveContacts"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves an alert contact.
//
// Description:
//
// - Before you call this operation, ensure you have activated Phone Number Intelligence.
//
// @param request - SaveContactsRequest
//
// @return SaveContactsResponse
func (client *Client) SaveContacts(request *SaveContactsRequest) (_result *SaveContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveContactsResponse{}
	_body, _err := client.SaveContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verification of three elements (name, mobile phone number, and ID card number). Used to verify whether the name, mobile phone number, and ID card number entered by a user belong to the same user.
//
// Description:
//
// - Before you call this API, make sure that you have fully understood the [pricing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
// - Before you call this API, log on to the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, fill in the application materials, and call the API after the application is approved.
//
// - Billing applies when the API returns Code=\\"OK\\" and IsConsistent != 2. Other return results are not billed.
//
// - The verification of virtual carrier numbers is not supported. Virtual carrier numbers refer to numbers that start with 170, 171, 162, or 165.
//
// ### QPS limit
//
// The QPS limit per user for this API is 200 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the API in a reasonable manner.
//
// @param request - ThreeElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ThreeElementsVerificationResponse
func (client *Client) ThreeElementsVerificationWithOptions(request *ThreeElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *ThreeElementsVerificationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verification of three elements (name, mobile phone number, and ID card number). Used to verify whether the name, mobile phone number, and ID card number entered by a user belong to the same user.
//
// Description:
//
// - Before you call this API, make sure that you have fully understood the [pricing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
// - Before you call this API, log on to the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply for Activation**, fill in the application materials, and call the API after the application is approved.
//
// - Billing applies when the API returns Code=\\"OK\\" and IsConsistent != 2. Other return results are not billed.
//
// - The verification of virtual carrier numbers is not supported. Virtual carrier numbers refer to numbers that start with 170, 171, 162, or 165.
//
// ### QPS limit
//
// The QPS limit per user for this API is 200 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the API in a reasonable manner.
//
// @param request - ThreeElementsVerificationRequest
//
// @return ThreeElementsVerificationResponse
func (client *Client) ThreeElementsVerification(request *ThreeElementsVerificationRequest) (_result *ThreeElementsVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ThreeElementsVerificationResponse{}
	_body, _err := client.ThreeElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Two-element (name and phone number) verification. Used to compare whether the name and phone number entered by the user belong to the same user.
//
// Description:
//
// - Before using this API, ensure that you fully understand the [product pricing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
// - Before using this API, log on to the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply to Enable**, fill in the application materials, and use it after the approval.
//
// - Billing is applied when the API returns Code="OK" and IsConsistent != 2. Other return results are not billed.
//
// - MVNO number verification is not supported. MVNO numbers refer to numbers starting with 170, 171, 162, and 165.
//
// ### QPS limit
//
// The per-user QPS limit of this API is 200 times/second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call the API reasonably.
//
// @param request - TwoElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TwoElementsVerificationResponse
func (client *Client) TwoElementsVerificationWithOptions(request *TwoElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *TwoElementsVerificationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Two-element (name and phone number) verification. Used to compare whether the name and phone number entered by the user belong to the same user.
//
// Description:
//
// - Before using this API, ensure that you fully understand the [product pricing](https://help.aliyun.com/document_detail/154751.html) of Cell Phone Number Service.
//
// - Before using this API, log on to the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, find the corresponding tag, click **Apply to Enable**, fill in the application materials, and use it after the approval.
//
// - Billing is applied when the API returns Code="OK" and IsConsistent != 2. Other return results are not billed.
//
// - MVNO number verification is not supported. MVNO numbers refer to numbers starting with 170, 171, 162, and 165.
//
// ### QPS limit
//
// The per-user QPS limit of this API is 200 times/second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call the API reasonably.
//
// @param request - TwoElementsVerificationRequest
//
// @return TwoElementsVerificationResponse
func (client *Client) TwoElementsVerification(request *TwoElementsVerificationRequest) (_result *TwoElementsVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TwoElementsVerificationResponse{}
	_body, _err := client.TwoElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UAIDCollectionWithOptions(request *UAIDCollectionRequest, runtime *dara.RuntimeOptions) (_result *UAIDCollectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UAIDCollectionResponse
func (client *Client) UAIDCollection(request *UAIDCollectionRequest) (_result *UAIDCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UAIDCollectionResponse{}
	_body, _err := client.UAIDCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UAIDConversionWithOptions(request *UAIDConversionRequest, runtime *dara.RuntimeOptions) (_result *UAIDConversionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UAIDConversionResponse
func (client *Client) UAIDConversion(request *UAIDConversionRequest) (_result *UAIDConversionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UAIDConversionResponse{}
	_body, _err := client.UAIDConversionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a UAID based on the carrier authorization token.
//
// For information about how to obtain the authorization token and its signature, see the GetUAIDApplyTokenSign API documentation.
//
// A UAID is 64 characters in length. The first 32 characters describe the device information, and the last 32 characters describe the phone number information.
//
// Description:
//
// Before you call this API, make sure that you have fully understood the billing method and [pricing](https://www.aliyun.com/price/product#/dytns/detail/dytns_penqbag_public_cn) of Cell Phone Number Service.
//
// Obtains a UAID based on the carrier authorization token.
//
// For information about how to obtain the authorization token and its signature, see the GetUAIDApplyTokenSign API documentation.
//
// @param request - UAIDVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UAIDVerificationResponse
func (client *Client) UAIDVerificationWithOptions(request *UAIDVerificationRequest, runtime *dara.RuntimeOptions) (_result *UAIDVerificationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a UAID based on the carrier authorization token.
//
// For information about how to obtain the authorization token and its signature, see the GetUAIDApplyTokenSign API documentation.
//
// A UAID is 64 characters in length. The first 32 characters describe the device information, and the last 32 characters describe the phone number information.
//
// Description:
//
// Before you call this API, make sure that you have fully understood the billing method and [pricing](https://www.aliyun.com/price/product#/dytns/detail/dytns_penqbag_public_cn) of Cell Phone Number Service.
//
// Obtains a UAID based on the carrier authorization token.
//
// For information about how to obtain the authorization token and its signature, see the GetUAIDApplyTokenSign API documentation.
//
// @param request - UAIDVerificationRequest
//
// @return UAIDVerificationResponse
func (client *Client) UAIDVerification(request *UAIDVerificationRequest) (_result *UAIDVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UAIDVerificationResponse{}
	_body, _err := client.UAIDVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an alert contact.
//
// Description:
//
// - Ensure that you have activated the Phone Number Information Service before calling this operation.
//
// @param request - UpdateContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContactsResponse
func (client *Client) UpdateContactsWithOptions(request *UpdateContactsRequest, runtime *dara.RuntimeOptions) (_result *UpdateContactsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an alert contact.
//
// Description:
//
// - Ensure that you have activated the Phone Number Information Service before calling this operation.
//
// @param request - UpdateContactsRequest
//
// @return UpdateContactsResponse
func (client *Client) UpdateContacts(request *UpdateContactsRequest) (_result *UpdateContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContactsResponse{}
	_body, _err := client.UpdateContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 虚商三要素
//
// @param request - VirtualThreeElementsVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VirtualThreeElementsVerificationResponse
func (client *Client) VirtualThreeElementsVerificationWithOptions(request *VirtualThreeElementsVerificationRequest, runtime *dara.RuntimeOptions) (_result *VirtualThreeElementsVerificationResponse, _err error) {
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

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
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
		Action:      dara.String("VirtualThreeElementsVerification"),
		Version:     dara.String("2020-02-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VirtualThreeElementsVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 虚商三要素
//
// @param request - VirtualThreeElementsVerificationRequest
//
// @return VirtualThreeElementsVerificationResponse
func (client *Client) VirtualThreeElementsVerification(request *VirtualThreeElementsVerificationRequest) (_result *VirtualThreeElementsVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VirtualThreeElementsVerificationResponse{}
	_body, _err := client.VirtualThreeElementsVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
