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
// 地址核验
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
// 地址核验
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
// 银行卡核验
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
// 银行卡核验
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
// 证件OCR识别纯服务端接口
//
// @param request - CardOcrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CardOcrResponse
// Deprecated
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
// 证件OCR识别纯服务端接口
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
// 结果查询
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
// 结果查询
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
// 认证日志查询接口
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
// 认证日志查询接口
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
// 凭证核验
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
// 凭证核验
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
// 人脸凭证核验
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
// 人脸凭证核验
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
// 删除用户认证记录结果
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
// 删除用户认证记录结果
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
// 卡证ocr纯服务端
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
// 卡证ocr纯服务端
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
// 全球证件ocr识别接口
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
// 全球证件ocr识别接口
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
// ekyc纯服务端接口
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
// ekyc纯服务端接口
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
// 人脸比对
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
// 人脸比对
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
// 国际人脸保镖纯服务端接口
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
// 国际人脸保镖纯服务端接口
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
// 静默活体API 纯服务端
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
// 静默活体API 纯服务端
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
// 防伪回调接口
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
// 防伪回调接口
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
// 身份二要素有效期核验
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
// 身份二要素有效期核验
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
// 身份二要素国际版接口
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
// 身份二要素国际版接口
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
// 认证初始化
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
	_body, _err := client.CallApi(params, req, runtime)
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
// 客户端连接保持
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
// 客户端连接保持
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
// 手机号三要素国际版接口
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
// 手机号三要素国际版接口
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
