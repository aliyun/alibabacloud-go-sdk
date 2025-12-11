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
	client.EndpointRule = dara.String("central")
	client.EndpointMap = map[string]*string{
		"ap-southeast-1": dara.String("dysmsapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5": dara.String("dysmsapi.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":     dara.String("dysmsapi-proxy.cn-beijing.aliyuncs.com"),
		"cn-hongkong":    dara.String("dysmsapi-xman.cn-hongkong.aliyuncs.com"),
		"eu-central-1":   dara.String("dysmsapi.eu-central-1.aliyuncs.com"),
		"us-east-1":      dara.String("dysmsapi.us-east-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dysmsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加验证码签名信息
//
// @param request - AddExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddExtCodeSignResponse
func (client *Client) AddExtCodeSignWithOptions(request *AddExtCodeSignRequest, runtime *dara.RuntimeOptions) (_result *AddExtCodeSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtCode) {
		query["ExtCode"] = request.ExtCode
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddExtCodeSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddExtCodeSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加验证码签名信息
//
// @param request - AddExtCodeSignRequest
//
// @return AddExtCodeSignResponse
func (client *Client) AddExtCodeSign(request *AddExtCodeSignRequest) (_result *AddExtCodeSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddExtCodeSignResponse{}
	_body, _err := client.AddExtCodeSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a short URL.
//
// Description:
//
//	  Before you call this operation, you must register the primary domain name of the source URL in the Short Message Service (SMS) console. After the domain name is registered, you can call this operation to create a short URL. For more information, see [Domain name registration](https://help.aliyun.com/document_detail/302325.html#title-mau-zdh-hd0).
//
//		- You can create up to 3,000 short URLs within a natural day.
//
//		- After a short URL is generated, a security review is required. Generally, the review takes 10 minutes to 2 hours to complete. Before the security review is passed, the short URL cannot be directly accessed.
//
// @param request - AddShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShortUrlResponse
func (client *Client) AddShortUrlWithOptions(request *AddShortUrlRequest, runtime *dara.RuntimeOptions) (_result *AddShortUrlResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveDays) {
		body["EffectiveDays"] = request.EffectiveDays
	}

	if !dara.IsNil(request.ShortUrlName) {
		body["ShortUrlName"] = request.ShortUrlName
	}

	if !dara.IsNil(request.SourceUrl) {
		body["SourceUrl"] = request.SourceUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddShortUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a short URL.
//
// Description:
//
//	  Before you call this operation, you must register the primary domain name of the source URL in the Short Message Service (SMS) console. After the domain name is registered, you can call this operation to create a short URL. For more information, see [Domain name registration](https://help.aliyun.com/document_detail/302325.html#title-mau-zdh-hd0).
//
//		- You can create up to 3,000 short URLs within a natural day.
//
//		- After a short URL is generated, a security review is required. Generally, the review takes 10 minutes to 2 hours to complete. Before the security review is passed, the short URL cannot be directly accessed.
//
// @param request - AddShortUrlRequest
//
// @return AddShortUrlResponse
func (client *Client) AddShortUrl(request *AddShortUrlRequest) (_result *AddShortUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddShortUrlResponse{}
	_body, _err := client.AddShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a signature.
//
// Description:
//
// You can call the AddSmsSign operation or use the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm#/overview) to create an SMS signature. The signature must comply with the [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html). You can call the QuerySmsSign operation or use the SMS console to query the review status of the signature.
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limit
//
// You can call this operation only once per second. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// >
//
//   - You cannot cancel the review of a signature.
//
//   - Individual users can create only one verification code signature, and can create only one general-purpose signature within a natural day. If you need to apply for multiple signatures, we recommend that you upgrade your account to an enterprise user.
//
//   - If you need to use the same signature for messages sent to recipients both in and outside the Chinese mainland, the signature must be a general-purpose signature.
//
//   - If you apply for a signature or message template, you must specify the signature scenario or template type. You must also provide the information of your services, such as a website URL, a domain name with an ICP filing, an application download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
//
//   - An SMS signature must undergo a thorough review process before it can be approved for use.
//
// @param request - AddSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSmsSignResponse
func (client *Client) AddSmsSignWithOptions(request *AddSmsSignRequest, runtime *dara.RuntimeOptions) (_result *AddSmsSignResponse, _err error) {
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

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.SignSource) {
		query["SignSource"] = request.SignSource
	}

	if !dara.IsNil(request.SignType) {
		query["SignType"] = request.SignType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SignFileList) {
		body["SignFileList"] = request.SignFileList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSmsSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a signature.
//
// Description:
//
// You can call the AddSmsSign operation or use the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm#/overview) to create an SMS signature. The signature must comply with the [SMS signature specifications](https://help.aliyun.com/document_detail/108076.html). You can call the QuerySmsSign operation or use the SMS console to query the review status of the signature.
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limit
//
// You can call this operation only once per second. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// >
//
//   - You cannot cancel the review of a signature.
//
//   - Individual users can create only one verification code signature, and can create only one general-purpose signature within a natural day. If you need to apply for multiple signatures, we recommend that you upgrade your account to an enterprise user.
//
//   - If you need to use the same signature for messages sent to recipients both in and outside the Chinese mainland, the signature must be a general-purpose signature.
//
//   - If you apply for a signature or message template, you must specify the signature scenario or template type. You must also provide the information of your services, such as a website URL, a domain name with an ICP filing, an application download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
//
//   - An SMS signature must undergo a thorough review process before it can be approved for use.
//
// @param request - AddSmsSignRequest
//
// @return AddSmsSignResponse
func (client *Client) AddSmsSign(request *AddSmsSignRequest) (_result *AddSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSmsSignResponse{}
	_body, _err := client.AddSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI AddSmsTemplate is deprecated, please use Dysmsapi::2017-05-25::CreateSmsTemplate instead.
//
// Summary:
//
// Creates a message template.
//
// Description:
//
// You can call the operation or use the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview) to apply for a message template. The template must comply with the [message template specifications](https://help.aliyun.com/document_detail/108253.html). You can call the [QuerySmsTemplate](https://help.aliyun.com/document_detail/419289.html) operation or use the Alibaba Cloud SMS console to check whether the message template is approved.
//
// >
//
//   - Message templates pending approval can be withdrawn. You can withdraw a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
//   - Message templates that have been approved can be deleted, and cannot be modified. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
//   - If you call the AddSmsTemplate operation, you can apply for a maximum of 100 message templates in a calendar day. After you apply for a message template, we recommend that you wait for at least 30 seconds before you apply for another one. If you use the Alibaba Cloud SMS console, you can apply for an unlimited number of message templates.
//
//   - Messages sent to the Chinese mainland and messages sent to countries or regions outside the Chinese mainland use separate message templates. Create message templates based on your needs.
//
//   - If you apply for a signature or message template, you must specify the signature scenario or template type. You must also provide the information of your services, such as a website URL, a domain name with an ICP filing, an application download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
//
//   - A signature must undergo a thorough review process before it can be approved for use. For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSmsTemplateResponse
func (client *Client) AddSmsTemplateWithOptions(request *AddSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddSmsTemplateResponse, _err error) {
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

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateContent) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddSmsTemplate is deprecated, please use Dysmsapi::2017-05-25::CreateSmsTemplate instead.
//
// Summary:
//
// Creates a message template.
//
// Description:
//
// You can call the operation or use the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview) to apply for a message template. The template must comply with the [message template specifications](https://help.aliyun.com/document_detail/108253.html). You can call the [QuerySmsTemplate](https://help.aliyun.com/document_detail/419289.html) operation or use the Alibaba Cloud SMS console to check whether the message template is approved.
//
// >
//
//   - Message templates pending approval can be withdrawn. You can withdraw a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
//   - Message templates that have been approved can be deleted, and cannot be modified. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
//   - If you call the AddSmsTemplate operation, you can apply for a maximum of 100 message templates in a calendar day. After you apply for a message template, we recommend that you wait for at least 30 seconds before you apply for another one. If you use the Alibaba Cloud SMS console, you can apply for an unlimited number of message templates.
//
//   - Messages sent to the Chinese mainland and messages sent to countries or regions outside the Chinese mainland use separate message templates. Create message templates based on your needs.
//
//   - If you apply for a signature or message template, you must specify the signature scenario or template type. You must also provide the information of your services, such as a website URL, a domain name with an ICP filing, an application download URL, or the name of your WeChat official account or mini program. For sign-in scenarios, you must also provide an account and password for tests. A detailed description can improve the review efficiency of signatures and templates.
//
//   - A signature must undergo a thorough review process before it can be approved for use. For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddSmsTemplateRequest
//
// @return AddSmsTemplateResponse
// Deprecated
func (client *Client) AddSmsTemplate(request *AddSmsTemplateRequest) (_result *AddSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSmsTemplateResponse{}
	_body, _err := client.AddSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更换签名的资质和授权书
//
// @param request - ChangeSignatureQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeSignatureQualificationResponse
func (client *Client) ChangeSignatureQualificationWithOptions(request *ChangeSignatureQualificationRequest, runtime *dara.RuntimeOptions) (_result *ChangeSignatureQualificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationLetterId) {
		query["AuthorizationLetterId"] = request.AuthorizationLetterId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignatureName) {
		query["SignatureName"] = request.SignatureName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeSignatureQualification"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeSignatureQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更换签名的资质和授权书
//
// @param request - ChangeSignatureQualificationRequest
//
// @return ChangeSignatureQualificationResponse
func (client *Client) ChangeSignatureQualification(request *ChangeSignatureQualificationRequest) (_result *ChangeSignatureQualificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeSignatureQualificationResponse{}
	_body, _err := client.ChangeSignatureQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a mobile phone number can receive card messages.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 2,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CheckMobilesCardSupportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckMobilesCardSupportResponse
func (client *Client) CheckMobilesCardSupportWithOptions(request *CheckMobilesCardSupportRequest, runtime *dara.RuntimeOptions) (_result *CheckMobilesCardSupportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mobiles) {
		query["Mobiles"] = request.Mobiles
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckMobilesCardSupport"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckMobilesCardSupportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a mobile phone number can receive card messages.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 2,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CheckMobilesCardSupportRequest
//
// @return CheckMobilesCardSupportResponse
func (client *Client) CheckMobilesCardSupport(request *CheckMobilesCardSupportRequest) (_result *CheckMobilesCardSupportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckMobilesCardSupportResponse{}
	_body, _err := client.CheckMobilesCardSupportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends conversion rate information to Alibaba Cloud SMS.
//
// @param request - ConversionDataIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConversionDataIntlResponse
func (client *Client) ConversionDataIntlWithOptions(request *ConversionDataIntlRequest, runtime *dara.RuntimeOptions) (_result *ConversionDataIntlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversionRate) {
		query["ConversionRate"] = request.ConversionRate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReportTime) {
		query["ReportTime"] = request.ReportTime
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
		Action:      dara.String("ConversionDataIntl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConversionDataIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends conversion rate information to Alibaba Cloud SMS.
//
// @param request - ConversionDataIntlRequest
//
// @return ConversionDataIntlResponse
func (client *Client) ConversionDataIntl(request *ConversionDataIntlRequest) (_result *ConversionDataIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConversionDataIntlResponse{}
	_body, _err := client.ConversionDataIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a card message template.
//
// Description:
//
//	  The CreateCardSmsTemplate operation saves the card message template information, submits it to the mobile phone manufacturer for approval, and returns the message template ID.
//
//		- If the type of the message template is not supported or events that are not supported by the mobile phone manufacturer are specified, the template is not submitted. For more information, see [Supported message templates](https://help.aliyun.com/document_detail/434611.html).
//
//		- For information about sample card message templates, see [Sample card message templates](https://help.aliyun.com/document_detail/435361.html).
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param tmpReq - CreateCardSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCardSmsTemplateResponse
func (client *Client) CreateCardSmsTemplateWithOptions(tmpReq *CreateCardSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateCardSmsTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCardSmsTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Factorys) {
		query["Factorys"] = request.Factorys
	}

	if !dara.IsNil(request.Memo) {
		query["Memo"] = request.Memo
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCardSmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCardSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a card message template.
//
// Description:
//
//	  The CreateCardSmsTemplate operation saves the card message template information, submits it to the mobile phone manufacturer for approval, and returns the message template ID.
//
//		- If the type of the message template is not supported or events that are not supported by the mobile phone manufacturer are specified, the template is not submitted. For more information, see [Supported message templates](https://help.aliyun.com/document_detail/434611.html).
//
//		- For information about sample card message templates, see [Sample card message templates](https://help.aliyun.com/document_detail/435361.html).
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateCardSmsTemplateRequest
//
// @return CreateCardSmsTemplateResponse
func (client *Client) CreateCardSmsTemplate(request *CreateCardSmsTemplateRequest) (_result *CreateCardSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCardSmsTemplateResponse{}
	_body, _err := client.CreateCardSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建短链
//
// @param request - CreateSmartShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmartShortUrlResponse
func (client *Client) CreateSmartShortUrlWithOptions(request *CreateSmartShortUrlRequest, runtime *dara.RuntimeOptions) (_result *CreateSmartShortUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumbers) {
		query["PhoneNumbers"] = request.PhoneNumbers
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceUrl) {
		query["SourceUrl"] = request.SourceUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmartShortUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmartShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建短链
//
// @param request - CreateSmartShortUrlRequest
//
// @return CreateSmartShortUrlResponse
func (client *Client) CreateSmartShortUrl(request *CreateSmartShortUrlRequest) (_result *CreateSmartShortUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSmartShortUrlResponse{}
	_body, _err := client.CreateSmartShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建委托授权书
//
// @param tmpReq - CreateSmsAuthorizationLetterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsAuthorizationLetterResponse
func (client *Client) CreateSmsAuthorizationLetterWithOptions(tmpReq *CreateSmsAuthorizationLetterRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsAuthorizationLetterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSmsAuthorizationLetterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SignList) {
		request.SignListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SignList, dara.String("SignList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Authorization) {
		query["Authorization"] = request.Authorization
	}

	if !dara.IsNil(request.AuthorizationLetterExpDate) {
		query["AuthorizationLetterExpDate"] = request.AuthorizationLetterExpDate
	}

	if !dara.IsNil(request.AuthorizationLetterName) {
		query["AuthorizationLetterName"] = request.AuthorizationLetterName
	}

	if !dara.IsNil(request.AuthorizationLetterPic) {
		query["AuthorizationLetterPic"] = request.AuthorizationLetterPic
	}

	if !dara.IsNil(request.OrganizationCode) {
		query["OrganizationCode"] = request.OrganizationCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProxyAuthorization) {
		query["ProxyAuthorization"] = request.ProxyAuthorization
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignListShrink) {
		query["SignList"] = request.SignListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmsAuthorizationLetter"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsAuthorizationLetterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建委托授权书
//
// @param request - CreateSmsAuthorizationLetterRequest
//
// @return CreateSmsAuthorizationLetterResponse
func (client *Client) CreateSmsAuthorizationLetter(request *CreateSmsAuthorizationLetterRequest) (_result *CreateSmsAuthorizationLetterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSmsAuthorizationLetterResponse{}
	_body, _err := client.CreateSmsAuthorizationLetterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create SMS Signature
//
// Description:
//
// - For details about the announcement of changes to the new and original interfaces, see [Announcement on Updates to SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Individual authenticated users can apply for one formal signature per natural day under the same Alibaba Cloud account, while enterprise authenticated users have no current restrictions. For details on the differences in rights between individual and enterprise users, please refer to [User Guide](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm).
//
// - Signature information applied through the interface will be synchronized in the SMS service console. For operations related to signatures in the console, see [SMS Signatures](https://help.aliyun.com/zh/sms/user-guide/create-signatures?spm).
//
// - After submitting the signature application, you can query the signature review status and details via the [GetSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-getsmssign?spm) interface. You can also [Configure Receipt Messages](https://help.aliyun.com/zh/sms/developer-reference/configure-delivery-receipts-1?spm) and obtain signature review status messages through SignSmsReport.
//
// @param tmpReq - CreateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSignWithOptions(tmpReq *CreateSmsSignRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSmsSignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MoreData) {
		request.MoreDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MoreData, dara.String("MoreData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIcpRecordId) {
		query["AppIcpRecordId"] = request.AppIcpRecordId
	}

	if !dara.IsNil(request.ApplySceneContent) {
		query["ApplySceneContent"] = request.ApplySceneContent
	}

	if !dara.IsNil(request.AuthorizationLetterId) {
		query["AuthorizationLetterId"] = request.AuthorizationLetterId
	}

	if !dara.IsNil(request.MoreDataShrink) {
		query["MoreData"] = request.MoreDataShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.SignSource) {
		query["SignSource"] = request.SignSource
	}

	if !dara.IsNil(request.SignType) {
		query["SignType"] = request.SignType
	}

	if !dara.IsNil(request.ThirdParty) {
		query["ThirdParty"] = request.ThirdParty
	}

	if !dara.IsNil(request.TrademarkId) {
		query["TrademarkId"] = request.TrademarkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmsSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create SMS Signature
//
// Description:
//
// - For details about the announcement of changes to the new and original interfaces, see [Announcement on Updates to SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Individual authenticated users can apply for one formal signature per natural day under the same Alibaba Cloud account, while enterprise authenticated users have no current restrictions. For details on the differences in rights between individual and enterprise users, please refer to [User Guide](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm).
//
// - Signature information applied through the interface will be synchronized in the SMS service console. For operations related to signatures in the console, see [SMS Signatures](https://help.aliyun.com/zh/sms/user-guide/create-signatures?spm).
//
// - After submitting the signature application, you can query the signature review status and details via the [GetSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-getsmssign?spm) interface. You can also [Configure Receipt Messages](https://help.aliyun.com/zh/sms/developer-reference/configure-delivery-receipts-1?spm) and obtain signature review status messages through SignSmsReport.
//
// @param request - CreateSmsSignRequest
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSign(request *CreateSmsSignRequest) (_result *CreateSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CreateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create SMS Template
//
// Description:
//
// - For details about the changes of this new interface compared to the original one, please refer to [Announcement on the Update of SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - It is recommended to apply for SMS templates via the API with at least a 30-second interval between each request.
//
// - The template information applied through the API will be synchronized in the SMS service console. For operations related to templates in the console, please refer to SMS Templates.
//
// - After submitting the template application, you can query the audit status and details using the GetSmsTemplate interface. You can also configure delivery receipts to obtain the audit status messages via TemplateSmsReport.
//
// - Domestic SMS templates are not interchangeable with international/Hong Kong, Macao, and Taiwan SMS templates. Please apply for templates based on your business scenario.
//
// - Only enterprise-verified users can apply for promotional messages and international/Hong Kong, Macao, and Taiwan messages. For differences in rights between personal and enterprise users, please refer to Usage Instructions.
//
// @param tmpReq - CreateSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsTemplateResponse
func (client *Client) CreateSmsTemplateWithOptions(tmpReq *CreateSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSmsTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MoreData) {
		request.MoreDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MoreData, dara.String("MoreData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplySceneContent) {
		query["ApplySceneContent"] = request.ApplySceneContent
	}

	if !dara.IsNil(request.IntlType) {
		query["IntlType"] = request.IntlType
	}

	if !dara.IsNil(request.MoreDataShrink) {
		query["MoreData"] = request.MoreDataShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RelatedSignName) {
		query["RelatedSignName"] = request.RelatedSignName
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateContent) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateRule) {
		query["TemplateRule"] = request.TemplateRule
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.TrafficDriving) {
		query["TrafficDriving"] = request.TrafficDriving
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create SMS Template
//
// Description:
//
// - For details about the changes of this new interface compared to the original one, please refer to [Announcement on the Update of SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - It is recommended to apply for SMS templates via the API with at least a 30-second interval between each request.
//
// - The template information applied through the API will be synchronized in the SMS service console. For operations related to templates in the console, please refer to SMS Templates.
//
// - After submitting the template application, you can query the audit status and details using the GetSmsTemplate interface. You can also configure delivery receipts to obtain the audit status messages via TemplateSmsReport.
//
// - Domestic SMS templates are not interchangeable with international/Hong Kong, Macao, and Taiwan SMS templates. Please apply for templates based on your business scenario.
//
// - Only enterprise-verified users can apply for promotional messages and international/Hong Kong, Macao, and Taiwan messages. For differences in rights between personal and enterprise users, please refer to Usage Instructions.
//
// @param request - CreateSmsTemplateRequest
//
// @return CreateSmsTemplateResponse
func (client *Client) CreateSmsTemplate(request *CreateSmsTemplateRequest) (_result *CreateSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSmsTemplateResponse{}
	_body, _err := client.CreateSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除验证码签名
//
// @param request - DeleteExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExtCodeSignResponse
func (client *Client) DeleteExtCodeSignWithOptions(request *DeleteExtCodeSignRequest, runtime *dara.RuntimeOptions) (_result *DeleteExtCodeSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtCode) {
		query["ExtCode"] = request.ExtCode
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExtCodeSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExtCodeSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除验证码签名
//
// @param request - DeleteExtCodeSignRequest
//
// @return DeleteExtCodeSignResponse
func (client *Client) DeleteExtCodeSign(request *DeleteExtCodeSignRequest) (_result *DeleteExtCodeSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExtCodeSignResponse{}
	_body, _err := client.DeleteExtCodeSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a short URL. After you delete a short URL, it cannot be changed to its original state.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteShortUrlResponse
func (client *Client) DeleteShortUrlWithOptions(request *DeleteShortUrlRequest, runtime *dara.RuntimeOptions) (_result *DeleteShortUrlResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceUrl) {
		body["SourceUrl"] = request.SourceUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteShortUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a short URL. After you delete a short URL, it cannot be changed to its original state.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteShortUrlRequest
//
// @return DeleteShortUrlResponse
func (client *Client) DeleteShortUrl(request *DeleteShortUrlRequest) (_result *DeleteShortUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteShortUrlResponse{}
	_body, _err := client.DeleteShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资质对客openAPI
//
// @param request - DeleteSmsQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmsQualificationResponse
func (client *Client) DeleteSmsQualificationWithOptions(request *DeleteSmsQualificationRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmsQualificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationGroupId) {
		query["QualificationGroupId"] = request.QualificationGroupId
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
		Action:      dara.String("DeleteSmsQualification"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSmsQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资质对客openAPI
//
// @param request - DeleteSmsQualificationRequest
//
// @return DeleteSmsQualificationResponse
func (client *Client) DeleteSmsQualification(request *DeleteSmsQualificationRequest) (_result *DeleteSmsQualificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSmsQualificationResponse{}
	_body, _err := client.DeleteSmsQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a signature.
//
// Description:
//
//	  You cannot delete a signature that has not been approved.
//
//		- After you delete a signature, you cannot recover it. Proceed with caution.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmsSignResponse
func (client *Client) DeleteSmsSignWithOptions(request *DeleteSmsSignRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmsSignResponse, _err error) {
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSmsSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a signature.
//
// Description:
//
//	  You cannot delete a signature that has not been approved.
//
//		- After you delete a signature, you cannot recover it. Proceed with caution.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSmsSignRequest
//
// @return DeleteSmsSignResponse
func (client *Client) DeleteSmsSign(request *DeleteSmsSignRequest) (_result *DeleteSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSmsSignResponse{}
	_body, _err := client.DeleteSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a message template.
//
// Description:
//
//	  Message templates pending approval can be withdrawn. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
//		- Message templates that have been approved can be deleted, and cannot be modified. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
//		- You cannot recover deleted message templates. Proceed with caution.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmsTemplateResponse
func (client *Client) DeleteSmsTemplateWithOptions(request *DeleteSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmsTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a message template.
//
// Description:
//
//	  Message templates pending approval can be withdrawn. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
//		- Message templates that have been approved can be deleted, and cannot be modified. You can delete a message template pending approval on the Message Templates tab in the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview).
//
//		- You cannot recover deleted message templates. Proceed with caution.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteSmsTemplateRequest
//
// @return DeleteSmsTemplateResponse
func (client *Client) DeleteSmsTemplate(request *DeleteSmsTemplateRequest) (_result *DeleteSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSmsTemplateResponse{}
	_body, _err := client.DeleteSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query card sending details
//
// @param request - GetCardSmsDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardSmsDetailsResponse
func (client *Client) GetCardSmsDetailsWithOptions(request *GetCardSmsDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetCardSmsDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizCardId) {
		query["BizCardId"] = request.BizCardId
	}

	if !dara.IsNil(request.BizDigitId) {
		query["BizDigitId"] = request.BizDigitId
	}

	if !dara.IsNil(request.BizSmsId) {
		query["BizSmsId"] = request.BizSmsId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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

	if !dara.IsNil(request.SendDate) {
		query["SendDate"] = request.SendDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCardSmsDetails"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCardSmsDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query card sending details
//
// @param request - GetCardSmsDetailsRequest
//
// @return GetCardSmsDetailsResponse
func (client *Client) GetCardSmsDetails(request *GetCardSmsDetailsRequest) (_result *GetCardSmsDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCardSmsDetailsResponse{}
	_body, _err := client.GetCardSmsDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the short URLs of a card messages template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCardSmsLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardSmsLinkResponse
func (client *Client) GetCardSmsLinkWithOptions(request *GetCardSmsLinkRequest, runtime *dara.RuntimeOptions) (_result *GetCardSmsLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CardCodeType) {
		query["CardCodeType"] = request.CardCodeType
	}

	if !dara.IsNil(request.CardLinkType) {
		query["CardLinkType"] = request.CardLinkType
	}

	if !dara.IsNil(request.CardTemplateCode) {
		query["CardTemplateCode"] = request.CardTemplateCode
	}

	if !dara.IsNil(request.CardTemplateParamJson) {
		query["CardTemplateParamJson"] = request.CardTemplateParamJson
	}

	if !dara.IsNil(request.CustomShortCodeJson) {
		query["CustomShortCodeJson"] = request.CustomShortCodeJson
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.PhoneNumberJson) {
		query["PhoneNumberJson"] = request.PhoneNumberJson
	}

	if !dara.IsNil(request.SignNameJson) {
		query["SignNameJson"] = request.SignNameJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCardSmsLink"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCardSmsLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the short URLs of a card messages template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetCardSmsLinkRequest
//
// @return GetCardSmsLinkResponse
func (client *Client) GetCardSmsLink(request *GetCardSmsLinkRequest) (_result *GetCardSmsLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCardSmsLinkResponse{}
	_body, _err := client.GetCardSmsLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Converts a resource uploaded to the specified Object Storage Service (OSS) bucket for unified management. Then, a resource ID is returned. You can manage the resource based on the ID.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetMediaResourceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaResourceIdResponse
func (client *Client) GetMediaResourceIdWithOptions(request *GetMediaResourceIdRequest, runtime *dara.RuntimeOptions) (_result *GetMediaResourceIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtendInfo) {
		query["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.FileSize) {
		query["FileSize"] = request.FileSize
	}

	if !dara.IsNil(request.Memo) {
		query["Memo"] = request.Memo
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaResourceId"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaResourceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts a resource uploaded to the specified Object Storage Service (OSS) bucket for unified management. Then, a resource ID is returned. You can manage the resource based on the ID.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetMediaResourceIdRequest
//
// @return GetMediaResourceIdResponse
func (client *Client) GetMediaResourceId(request *GetMediaResourceIdRequest) (_result *GetMediaResourceIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaResourceIdResponse{}
	_body, _err := client.GetMediaResourceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the OSS configuration information about card messages.
//
// Description:
//
// Resources such as images and videos used for card message templates can be uploaded to Object Storage Service (OSS) buckets for storage. For more information, see [Upload files to OSS](https://help.aliyun.com/document_detail/437303.html).
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetOSSInfoForCardTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSInfoForCardTemplateResponse
func (client *Client) GetOSSInfoForCardTemplateWithOptions(runtime *dara.RuntimeOptions) (_result *GetOSSInfoForCardTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetOSSInfoForCardTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOSSInfoForCardTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the OSS configuration information about card messages.
//
// Description:
//
// Resources such as images and videos used for card message templates can be uploaded to Object Storage Service (OSS) buckets for storage. For more information, see [Upload files to OSS](https://help.aliyun.com/document_detail/437303.html).
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @return GetOSSInfoForCardTemplateResponse
func (client *Client) GetOSSInfoForCardTemplate() (_result *GetOSSInfoForCardTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOSSInfoForCardTemplateResponse{}
	_body, _err := client.GetOSSInfoForCardTemplateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SMS File Upload, Get Authorization Info
//
// Description:
//
// - When creating signatures or templates, you can upload materials such as login pages with links, backend page screenshots, software copyrights, supplementary agreements, etc. This helps the review personnel understand your business details. If there are multiple materials, they can be combined into one file, supporting png, jpg, jpeg, doc, docx, pdf formats.
//
// - For additional materials needed when creating signatures or templates, you can upload them to the OSS file system for storage. For file upload operations, refer to [OSS File Upload](https://help.aliyun.com/zh/sms/upload-files-through-oss).
//
// @param request - GetOSSInfoForUploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSInfoForUploadFileResponse
func (client *Client) GetOSSInfoForUploadFileWithOptions(request *GetOSSInfoForUploadFileRequest, runtime *dara.RuntimeOptions) (_result *GetOSSInfoForUploadFileResponse, _err error) {
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
		Action:      dara.String("GetOSSInfoForUploadFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOSSInfoForUploadFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SMS File Upload, Get Authorization Info
//
// Description:
//
// - When creating signatures or templates, you can upload materials such as login pages with links, backend page screenshots, software copyrights, supplementary agreements, etc. This helps the review personnel understand your business details. If there are multiple materials, they can be combined into one file, supporting png, jpg, jpeg, doc, docx, pdf formats.
//
// - For additional materials needed when creating signatures or templates, you can upload them to the OSS file system for storage. For file upload operations, refer to [OSS File Upload](https://help.aliyun.com/zh/sms/upload-files-through-oss).
//
// @param request - GetOSSInfoForUploadFileRequest
//
// @return GetOSSInfoForUploadFileResponse
func (client *Client) GetOSSInfoForUploadFile(request *GetOSSInfoForUploadFileRequest) (_result *GetOSSInfoForUploadFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOSSInfoForUploadFileResponse{}
	_body, _err := client.GetOSSInfoForUploadFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传文件获取oss配置
//
// @param request - GetQualificationOssInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualificationOssInfoResponse
func (client *Client) GetQualificationOssInfoWithOptions(request *GetQualificationOssInfoRequest, runtime *dara.RuntimeOptions) (_result *GetQualificationOssInfoResponse, _err error) {
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
		Action:      dara.String("GetQualificationOssInfo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualificationOssInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传文件获取oss配置
//
// @param request - GetQualificationOssInfoRequest
//
// @return GetQualificationOssInfoResponse
func (client *Client) GetQualificationOssInfo(request *GetQualificationOssInfoRequest) (_result *GetQualificationOssInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualificationOssInfoResponse{}
	_body, _err := client.GetQualificationOssInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取OCR的OSS信息
//
// @param request - GetSmsOcrOssInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsOcrOssInfoResponse
func (client *Client) GetSmsOcrOssInfoWithOptions(request *GetSmsOcrOssInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSmsOcrOssInfoResponse, _err error) {
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

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmsOcrOssInfo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmsOcrOssInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取OCR的OSS信息
//
// @param request - GetSmsOcrOssInfoRequest
//
// @return GetSmsOcrOssInfoResponse
func (client *Client) GetSmsOcrOssInfo(request *GetSmsOcrOssInfoRequest) (_result *GetSmsOcrOssInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSmsOcrOssInfoResponse{}
	_body, _err := client.GetSmsOcrOssInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query SMS Signature Details
//
// Description:
//
// - For details about the changes of this new interface and the original one, please refer to [Announcement on the Update of SMS Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Review Time: Generally, after submitting the signature, Alibaba Cloud expects to complete the review within 2 hours (Review Business Hours: Monday to Sunday 9:00~21:00, with legal holidays postponed). It is recommended to submit your application before 18:00.
//
// - If the signature fails the review, the reason for the failure will be returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm), invoke the [UpdateSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-updatesmssign?spm) API, or modify the unapproved SMS signature on the [Signature Management](https://dysms.console.aliyun.com/domestic/text/sign) page.
//
// @param request - GetSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsSignResponse
func (client *Client) GetSmsSignWithOptions(request *GetSmsSignRequest, runtime *dara.RuntimeOptions) (_result *GetSmsSignResponse, _err error) {
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmsSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query SMS Signature Details
//
// Description:
//
// - For details about the changes of this new interface and the original one, please refer to [Announcement on the Update of SMS Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Review Time: Generally, after submitting the signature, Alibaba Cloud expects to complete the review within 2 hours (Review Business Hours: Monday to Sunday 9:00~21:00, with legal holidays postponed). It is recommended to submit your application before 18:00.
//
// - If the signature fails the review, the reason for the failure will be returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm), invoke the [UpdateSmsSign](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-updatesmssign?spm) API, or modify the unapproved SMS signature on the [Signature Management](https://dysms.console.aliyun.com/domestic/text/sign) page.
//
// @param request - GetSmsSignRequest
//
// @return GetSmsSignResponse
func (client *Client) GetSmsSign(request *GetSmsSignRequest) (_result *GetSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSmsSignResponse{}
	_body, _err := client.GetSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Text SMS Template Details
//
// Description:
//
// - For details about the announcement of changes to the new and original interfaces, see [Announcement on Updates to SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Review Time: Under normal circumstances, Alibaba Cloud expects to complete the review within 2 hours after template submission (review working hours: Monday to Sunday 9:00~21:00, with statutory holidays postponed). It is recommended to submit your application before 18:00.
//
// - If the template fails the review, the reason for the failure will be returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.41fd339f3bPSCQ), invoke the [ModifySmsTemplate](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-modifysmstemplate?spm=a2c4g.11186623.0.0.5b1f6e8bQloFit) API or modify the SMS template on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
//
// - The current QuerySmsTemplate interface queries the audit details of a single template by template code. The [QuerySmsTemplateList](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-querysmstemplatelist?spm=a2c4g.11186623.0.0.24086e8bO8cFn4) interface can query the template details of all templates under your current account.
//
// @param request - GetSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsTemplateResponse
func (client *Client) GetSmsTemplateWithOptions(request *GetSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetSmsTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Text SMS Template Details
//
// Description:
//
// - For details about the announcement of changes to the new and original interfaces, see [Announcement on Updates to SMS Service Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Review Time: Under normal circumstances, Alibaba Cloud expects to complete the review within 2 hours after template submission (review working hours: Monday to Sunday 9:00~21:00, with statutory holidays postponed). It is recommended to submit your application before 18:00.
//
// - If the template fails the review, the reason for the failure will be returned. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.41fd339f3bPSCQ), invoke the [ModifySmsTemplate](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-modifysmstemplate?spm=a2c4g.11186623.0.0.5b1f6e8bQloFit) API or modify the SMS template on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
//
// - The current QuerySmsTemplate interface queries the audit details of a single template by template code. The [QuerySmsTemplateList](https://help.aliyun.com/zh/sms/developer-reference/api-dysmsapi-2017-05-25-querysmstemplatelist?spm=a2c4g.11186623.0.0.24086e8bO8cFn4) interface can query the template details of all templates under your current account.
//
// @param request - GetSmsTemplateRequest
//
// @return GetSmsTemplateResponse
func (client *Client) GetSmsTemplate(request *GetSmsTemplateRequest) (_result *GetSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSmsTemplateResponse{}
	_body, _err := client.GetSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a rejected signature and submit it for approval. Signatures that are pending approval or have been approved cannot be modified.
//
// Description:
//
// You can call the operation or use the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview) to modify an existing signature and submit the signature for approval. The signature must comply with the [signature specifications](https://help.aliyun.com/document_detail/108076.html).
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// >
//
//   - Signatures pending approval cannot be modified.
//
//   - You cannot modify a signature after it is approved. If you no longer need the signature, you can delete it.
//
//   - If you are an individual user, you cannot apply for a new signature on the same day that your signature is rejected or deleted. We recommend that you modify the rejected signature and submit it again.
//
// @param request - ModifySmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmsSignResponse
func (client *Client) ModifySmsSignWithOptions(request *ModifySmsSignRequest, runtime *dara.RuntimeOptions) (_result *ModifySmsSignResponse, _err error) {
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

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.SignSource) {
		query["SignSource"] = request.SignSource
	}

	if !dara.IsNil(request.SignType) {
		query["SignType"] = request.SignType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SignFileList) {
		body["SignFileList"] = request.SignFileList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySmsSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a rejected signature and submit it for approval. Signatures that are pending approval or have been approved cannot be modified.
//
// Description:
//
// You can call the operation or use the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm#/overview) to modify an existing signature and submit the signature for approval. The signature must comply with the [signature specifications](https://help.aliyun.com/document_detail/108076.html).
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// >
//
//   - Signatures pending approval cannot be modified.
//
//   - You cannot modify a signature after it is approved. If you no longer need the signature, you can delete it.
//
//   - If you are an individual user, you cannot apply for a new signature on the same day that your signature is rejected or deleted. We recommend that you modify the rejected signature and submit it again.
//
// @param request - ModifySmsSignRequest
//
// @return ModifySmsSignResponse
func (client *Client) ModifySmsSign(request *ModifySmsSignRequest) (_result *ModifySmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySmsSignResponse{}
	_body, _err := client.ModifySmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifySmsTemplate is deprecated, please use Dysmsapi::2017-05-25::UpdateSmsTemplate instead.
//
// Summary:
//
// Modifies the information of an unapproved message template and submits it for review again.
//
// Description:
//
// After you apply for a message template, if the template fails to pass the review, you can call this operation to modify the template and submit the template again. You can call this operation to modify only a template for a specific message type.
//
// The template content must comply with the [SMS template specifications](https://help.aliyun.com/document_detail/108253.html).
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifySmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmsTemplateResponse
func (client *Client) ModifySmsTemplateWithOptions(request *ModifySmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifySmsTemplateResponse, _err error) {
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

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateContent) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifySmsTemplate is deprecated, please use Dysmsapi::2017-05-25::UpdateSmsTemplate instead.
//
// Summary:
//
// Modifies the information of an unapproved message template and submits it for review again.
//
// Description:
//
// After you apply for a message template, if the template fails to pass the review, you can call this operation to modify the template and submit the template again. You can call this operation to modify only a template for a specific message type.
//
// The template content must comply with the [SMS template specifications](https://help.aliyun.com/document_detail/108253.html).
//
// For more information, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifySmsTemplateRequest
//
// @return ModifySmsTemplateResponse
// Deprecated
func (client *Client) ModifySmsTemplate(request *ModifySmsTemplateRequest) (_result *ModifySmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySmsTemplateResponse{}
	_body, _err := client.ModifySmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the review status of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryCardSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCardSmsTemplateResponse
func (client *Client) QueryCardSmsTemplateWithOptions(request *QueryCardSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryCardSmsTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCardSmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCardSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the review status of a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryCardSmsTemplateRequest
//
// @return QueryCardSmsTemplateResponse
func (client *Client) QueryCardSmsTemplate(request *QueryCardSmsTemplateRequest) (_result *QueryCardSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCardSmsTemplateResponse{}
	_body, _err := client.QueryCardSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries sent card messages.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryCardSmsTemplateReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCardSmsTemplateReportResponse
func (client *Client) QueryCardSmsTemplateReportWithOptions(request *QueryCardSmsTemplateReportRequest, runtime *dara.RuntimeOptions) (_result *QueryCardSmsTemplateReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TemplateCodes) {
		query["TemplateCodes"] = request.TemplateCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCardSmsTemplateReport"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCardSmsTemplateReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sent card messages.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 300 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryCardSmsTemplateReportRequest
//
// @return QueryCardSmsTemplateReportResponse
func (client *Client) QueryCardSmsTemplateReport(request *QueryCardSmsTemplateReportRequest) (_result *QueryCardSmsTemplateReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCardSmsTemplateReportResponse{}
	_body, _err := client.QueryCardSmsTemplateReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询验证码签名
//
// @param request - QueryExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryExtCodeSignResponse
func (client *Client) QueryExtCodeSignWithOptions(request *QueryExtCodeSignRequest, runtime *dara.RuntimeOptions) (_result *QueryExtCodeSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtCode) {
		query["ExtCode"] = request.ExtCode
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryExtCodeSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryExtCodeSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询验证码签名
//
// @param request - QueryExtCodeSignRequest
//
// @return QueryExtCodeSignResponse
func (client *Client) QueryExtCodeSign(request *QueryExtCodeSignRequest) (_result *QueryExtCodeSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryExtCodeSignResponse{}
	_body, _err := client.QueryExtCodeSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a mobile phone number can receive card messages.
//
// @param tmpReq - QueryMobilesCardSupportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMobilesCardSupportResponse
func (client *Client) QueryMobilesCardSupportWithOptions(tmpReq *QueryMobilesCardSupportRequest, runtime *dara.RuntimeOptions) (_result *QueryMobilesCardSupportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMobilesCardSupportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Mobiles) {
		request.MobilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Mobiles, dara.String("Mobiles"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.MobilesShrink) {
		query["Mobiles"] = request.MobilesShrink
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMobilesCardSupport"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMobilesCardSupportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a mobile phone number can receive card messages.
//
// @param request - QueryMobilesCardSupportRequest
//
// @return QueryMobilesCardSupportResponse
func (client *Client) QueryMobilesCardSupport(request *QueryMobilesCardSupportRequest) (_result *QueryMobilesCardSupportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMobilesCardSupportResponse{}
	_body, _err := client.QueryMobilesCardSupportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 点击明细查询
//
// @param request - QueryPageSmartShortUrlLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPageSmartShortUrlLogResponse
func (client *Client) QueryPageSmartShortUrlLogWithOptions(request *QueryPageSmartShortUrlLogRequest, runtime *dara.RuntimeOptions) (_result *QueryPageSmartShortUrlLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateDateEnd) {
		query["CreateDateEnd"] = request.CreateDateEnd
	}

	if !dara.IsNil(request.CreateDateStart) {
		query["CreateDateStart"] = request.CreateDateStart
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

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ShortUrl) {
		query["ShortUrl"] = request.ShortUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPageSmartShortUrlLog"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPageSmartShortUrlLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 点击明细查询
//
// @param request - QueryPageSmartShortUrlLogRequest
//
// @return QueryPageSmartShortUrlLogResponse
func (client *Client) QueryPageSmartShortUrlLog(request *QueryPageSmartShortUrlLogRequest) (_result *QueryPageSmartShortUrlLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPageSmartShortUrlLogResponse{}
	_body, _err := client.QueryPageSmartShortUrlLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a message.
//
// @param request - QuerySendDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendDetailsResponse
func (client *Client) QuerySendDetailsWithOptions(request *QuerySendDetailsRequest, runtime *dara.RuntimeOptions) (_result *QuerySendDetailsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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

	if !dara.IsNil(request.SendDate) {
		query["SendDate"] = request.SendDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySendDetails"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySendDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a message.
//
// @param request - QuerySendDetailsRequest
//
// @return QuerySendDetailsResponse
func (client *Client) QuerySendDetails(request *QuerySendDetailsRequest) (_result *QuerySendDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySendDetailsResponse{}
	_body, _err := client.QuerySendDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries message delivery details.
//
// Description:
//
// You can call the operation to query message delivery details, including the number of delivered messages, the number of messages with delivery receipts, and the time that a message is sent. If a large number of messages are sent on the specified date, you can specify the number of items displayed on each page and the number of pages to view the details by page.
//
// ### QPS limits
//
// You can call this operation up to 20 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySendStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendStatisticsResponse
func (client *Client) QuerySendStatisticsWithOptions(request *QuerySendStatisticsRequest, runtime *dara.RuntimeOptions) (_result *QuerySendStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.IsGlobe) {
		query["IsGlobe"] = request.IsGlobe
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySendStatistics"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySendStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries message delivery details.
//
// Description:
//
// You can call the operation to query message delivery details, including the number of delivered messages, the number of messages with delivery receipts, and the time that a message is sent. If a large number of messages are sent on the specified date, you can specify the number of items displayed on each page and the number of pages to view the details by page.
//
// ### QPS limits
//
// You can call this operation up to 20 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySendStatisticsRequest
//
// @return QuerySendStatisticsResponse
func (client *Client) QuerySendStatistics(request *QuerySendStatisticsRequest) (_result *QuerySendStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySendStatisticsResponse{}
	_body, _err := client.QuerySendStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a short URL.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 20 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryShortUrlResponse
func (client *Client) QueryShortUrlWithOptions(request *QueryShortUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryShortUrlResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.ShortUrl) {
		body["ShortUrl"] = request.ShortUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryShortUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryShortUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a short URL.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 20 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryShortUrlRequest
//
// @return QueryShortUrlResponse
func (client *Client) QueryShortUrl(request *QueryShortUrlRequest) (_result *QueryShortUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryShortUrlResponse{}
	_body, _err := client.QueryShortUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个资质详情
//
// @param request - QuerySingleSmsQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySingleSmsQualificationResponse
func (client *Client) QuerySingleSmsQualificationWithOptions(request *QuerySingleSmsQualificationRequest, runtime *dara.RuntimeOptions) (_result *QuerySingleSmsQualificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationGroupId) {
		query["QualificationGroupId"] = request.QualificationGroupId
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
		Action:      dara.String("QuerySingleSmsQualification"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySingleSmsQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个资质详情
//
// @param request - QuerySingleSmsQualificationRequest
//
// @return QuerySingleSmsQualificationResponse
func (client *Client) QuerySingleSmsQualification(request *QuerySingleSmsQualificationRequest) (_result *QuerySingleSmsQualificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySingleSmsQualificationResponse{}
	_body, _err := client.QuerySingleSmsQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询委托授权书
//
// @param tmpReq - QuerySmsAuthorizationLetterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsAuthorizationLetterResponse
func (client *Client) QuerySmsAuthorizationLetterWithOptions(tmpReq *QuerySmsAuthorizationLetterRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsAuthorizationLetterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QuerySmsAuthorizationLetterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthorizationLetterIdList) {
		request.AuthorizationLetterIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthorizationLetterIdList, dara.String("AuthorizationLetterIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationLetterIdListShrink) {
		query["AuthorizationLetterIdList"] = request.AuthorizationLetterIdListShrink
	}

	if !dara.IsNil(request.OrganizationCode) {
		query["OrganizationCode"] = request.OrganizationCode
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmsAuthorizationLetter"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmsAuthorizationLetterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询委托授权书
//
// @param request - QuerySmsAuthorizationLetterRequest
//
// @return QuerySmsAuthorizationLetterResponse
func (client *Client) QuerySmsAuthorizationLetter(request *QuerySmsAuthorizationLetterRequest) (_result *QuerySmsAuthorizationLetterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySmsAuthorizationLetterResponse{}
	_body, _err := client.QuerySmsAuthorizationLetterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资质审核列表页
//
// @param request - QuerySmsQualificationRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsQualificationRecordResponse
func (client *Client) QuerySmsQualificationRecordWithOptions(request *QuerySmsQualificationRecordRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsQualificationRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompanyName) {
		query["CompanyName"] = request.CompanyName
	}

	if !dara.IsNil(request.LegalPersonName) {
		query["LegalPersonName"] = request.LegalPersonName
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

	if !dara.IsNil(request.QualificationGroupName) {
		query["QualificationGroupName"] = request.QualificationGroupName
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.UseBySelf) {
		query["UseBySelf"] = request.UseBySelf
	}

	if !dara.IsNil(request.WorkOrderId) {
		query["WorkOrderId"] = request.WorkOrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmsQualificationRecord"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmsQualificationRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资质审核列表页
//
// @param request - QuerySmsQualificationRecordRequest
//
// @return QuerySmsQualificationRecordResponse
func (client *Client) QuerySmsQualificationRecord(request *QuerySmsQualificationRecordRequest) (_result *QuerySmsQualificationRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySmsQualificationRecordResponse{}
	_body, _err := client.QuerySmsQualificationRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a signature.
//
// Description:
//
// After you apply for an SMS signature, you can query its status by using the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm) or calling the operation. If the signature is rejected, you can modify the signature based on the reason why it is rejected.
//
// ### QPS limits
//
// You can call this API operation up to 500 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsSignResponse
func (client *Client) QuerySmsSignWithOptions(request *QuerySmsSignRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsSignResponse, _err error) {
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmsSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a signature.
//
// Description:
//
// After you apply for an SMS signature, you can query its status by using the [Alibaba Cloud SMS console](https://dysms.console.aliyun.com/dysms.htm) or calling the operation. If the signature is rejected, you can modify the signature based on the reason why it is rejected.
//
// ### QPS limits
//
// You can call this API operation up to 500 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsSignRequest
//
// @return QuerySmsSignResponse
func (client *Client) QuerySmsSign(request *QuerySmsSignRequest) (_result *QuerySmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySmsSignResponse{}
	_body, _err := client.QuerySmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries message signatures by page.
//
// Description:
//
// You can call this operation to query the details of message signatures, including the name, creation time, and approval status of each signature. If a message template is rejected, the reason is returned. Modify the message signature based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsSignListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsSignListResponse
func (client *Client) QuerySmsSignListWithOptions(request *QuerySmsSignListRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsSignListResponse, _err error) {
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

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
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
		Action:      dara.String("QuerySmsSignList"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmsSignListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries message signatures by page.
//
// Description:
//
// You can call this operation to query the details of message signatures, including the name, creation time, and approval status of each signature. If a message template is rejected, the reason is returned. Modify the message signature based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsSignListRequest
//
// @return QuerySmsSignListResponse
func (client *Client) QuerySmsSignList(request *QuerySmsSignListRequest) (_result *QuerySmsSignListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySmsSignListResponse{}
	_body, _err := client.QuerySmsSignListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI QuerySmsTemplate is deprecated, please use Dysmsapi::2017-05-25::GetSmsTemplate instead.
//
// Summary:
//
// Queries the approval status of a message template.
//
// Description:
//
// After you create a message template, you can call this operation to query the approval status of the template. If a message template is rejected, the reason is returned. Modify the message template based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsTemplateResponse
func (client *Client) QuerySmsTemplateWithOptions(request *QuerySmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI QuerySmsTemplate is deprecated, please use Dysmsapi::2017-05-25::GetSmsTemplate instead.
//
// Summary:
//
// Queries the approval status of a message template.
//
// Description:
//
// After you create a message template, you can call this operation to query the approval status of the template. If a message template is rejected, the reason is returned. Modify the message template based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsTemplateRequest
//
// @return QuerySmsTemplateResponse
// Deprecated
func (client *Client) QuerySmsTemplate(request *QuerySmsTemplateRequest) (_result *QuerySmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySmsTemplateResponse{}
	_body, _err := client.QuerySmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries message templates.
//
// Description:
//
// You can call this operation to query the details of message templates, including the name, creation time, and approval status of each template. If a message template is rejected, the reason is returned. Modify the message template based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsTemplateListResponse
func (client *Client) QuerySmsTemplateListWithOptions(request *QuerySmsTemplateListRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsTemplateListResponse, _err error) {
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

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
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
		Action:      dara.String("QuerySmsTemplateList"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmsTemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries message templates.
//
// Description:
//
// You can call this operation to query the details of message templates, including the name, creation time, and approval status of each template. If a message template is rejected, the reason is returned. Modify the message template based on the reason.
//
// ### QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QuerySmsTemplateListRequest
//
// @return QuerySmsTemplateListResponse
func (client *Client) QuerySmsTemplateList(request *QuerySmsTemplateListRequest) (_result *QuerySmsTemplateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySmsTemplateListResponse{}
	_body, _err := client.QuerySmsTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证手机验证码
//
// @param request - RequiredPhoneCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RequiredPhoneCodeResponse
func (client *Client) RequiredPhoneCodeWithOptions(request *RequiredPhoneCodeRequest, runtime *dara.RuntimeOptions) (_result *RequiredPhoneCodeResponse, _err error) {
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

	if !dara.IsNil(request.PhoneNo) {
		query["PhoneNo"] = request.PhoneNo
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
		Action:      dara.String("RequiredPhoneCode"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RequiredPhoneCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证手机验证码
//
// @param request - RequiredPhoneCodeRequest
//
// @return RequiredPhoneCodeResponse
func (client *Client) RequiredPhoneCode(request *RequiredPhoneCodeRequest) (_result *RequiredPhoneCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RequiredPhoneCodeResponse{}
	_body, _err := client.RequiredPhoneCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends multiple card messages at a time.
//
// Description:
//
// You can call the operation to send multiple card messages to a maximum of mobile phone numbers at a time. Different signatures and rollback settings can be specified for the mobile phone numbers.
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendBatchCardSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendBatchCardSmsResponse
func (client *Client) SendBatchCardSmsWithOptions(request *SendBatchCardSmsRequest, runtime *dara.RuntimeOptions) (_result *SendBatchCardSmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CardTemplateCode) {
		query["CardTemplateCode"] = request.CardTemplateCode
	}

	if !dara.IsNil(request.CardTemplateParamJson) {
		query["CardTemplateParamJson"] = request.CardTemplateParamJson
	}

	if !dara.IsNil(request.DigitalTemplateCode) {
		query["DigitalTemplateCode"] = request.DigitalTemplateCode
	}

	if !dara.IsNil(request.DigitalTemplateParamJson) {
		query["DigitalTemplateParamJson"] = request.DigitalTemplateParamJson
	}

	if !dara.IsNil(request.FallbackType) {
		query["FallbackType"] = request.FallbackType
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.PhoneNumberJson) {
		query["PhoneNumberJson"] = request.PhoneNumberJson
	}

	if !dara.IsNil(request.SignNameJson) {
		query["SignNameJson"] = request.SignNameJson
	}

	if !dara.IsNil(request.SmsTemplateCode) {
		query["SmsTemplateCode"] = request.SmsTemplateCode
	}

	if !dara.IsNil(request.SmsTemplateParamJson) {
		query["SmsTemplateParamJson"] = request.SmsTemplateParamJson
	}

	if !dara.IsNil(request.SmsUpExtendCodeJson) {
		query["SmsUpExtendCodeJson"] = request.SmsUpExtendCodeJson
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateParamJson) {
		query["TemplateParamJson"] = request.TemplateParamJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendBatchCardSms"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendBatchCardSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends multiple card messages at a time.
//
// Description:
//
// You can call the operation to send multiple card messages to a maximum of mobile phone numbers at a time. Different signatures and rollback settings can be specified for the mobile phone numbers.
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendBatchCardSmsRequest
//
// @return SendBatchCardSmsResponse
func (client *Client) SendBatchCardSms(request *SendBatchCardSmsRequest) (_result *SendBatchCardSmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendBatchCardSmsResponse{}
	_body, _err := client.SendBatchCardSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uses a single message template and multiple signatures to send messages to multiple recipients.
//
// Description:
//
// You can call the operation to send messages to a maximum of 100 recipients at a time.
//
// @param request - SendBatchSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendBatchSmsResponse
func (client *Client) SendBatchSmsWithOptions(request *SendBatchSmsRequest, runtime *dara.RuntimeOptions) (_result *SendBatchSmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PhoneNumberJson) {
		body["PhoneNumberJson"] = request.PhoneNumberJson
	}

	if !dara.IsNil(request.SignNameJson) {
		body["SignNameJson"] = request.SignNameJson
	}

	if !dara.IsNil(request.SmsUpExtendCodeJson) {
		body["SmsUpExtendCodeJson"] = request.SmsUpExtendCodeJson
	}

	if !dara.IsNil(request.TemplateParamJson) {
		body["TemplateParamJson"] = request.TemplateParamJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendBatchSms"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendBatchSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uses a single message template and multiple signatures to send messages to multiple recipients.
//
// Description:
//
// You can call the operation to send messages to a maximum of 100 recipients at a time.
//
// @param request - SendBatchSmsRequest
//
// @return SendBatchSmsResponse
func (client *Client) SendBatchSms(request *SendBatchSmsRequest) (_result *SendBatchSmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendBatchSmsResponse{}
	_body, _err := client.SendBatchSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a card message.
//
// Description:
//
//	  Make sure that the message template that you want to use has been approved. If the mobile phone number of a recipient does not support card messages, the SendCardSms operation allows the rollback feature to ensure successful delivery.
//
//		- When you call the SendCardSms operation to send card messages, the operation checks whether the mobile phone numbers of the recipients support card messages. If the mobile phone numbers do not support card messages, you can specify whether to enable rollback. Otherwise, the card message cannot be delivered.
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendCardSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendCardSmsResponse
func (client *Client) SendCardSmsWithOptions(request *SendCardSmsRequest, runtime *dara.RuntimeOptions) (_result *SendCardSmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CardObjects) {
		query["CardObjects"] = request.CardObjects
	}

	if !dara.IsNil(request.CardTemplateCode) {
		query["CardTemplateCode"] = request.CardTemplateCode
	}

	if !dara.IsNil(request.DigitalTemplateCode) {
		query["DigitalTemplateCode"] = request.DigitalTemplateCode
	}

	if !dara.IsNil(request.DigitalTemplateParam) {
		query["DigitalTemplateParam"] = request.DigitalTemplateParam
	}

	if !dara.IsNil(request.FallbackType) {
		query["FallbackType"] = request.FallbackType
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.SmsTemplateCode) {
		query["SmsTemplateCode"] = request.SmsTemplateCode
	}

	if !dara.IsNil(request.SmsTemplateParam) {
		query["SmsTemplateParam"] = request.SmsTemplateParam
	}

	if !dara.IsNil(request.SmsUpExtendCode) {
		query["SmsUpExtendCode"] = request.SmsUpExtendCode
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateParam) {
		query["TemplateParam"] = request.TemplateParam
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendCardSms"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendCardSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a card message.
//
// Description:
//
//	  Make sure that the message template that you want to use has been approved. If the mobile phone number of a recipient does not support card messages, the SendCardSms operation allows the rollback feature to ensure successful delivery.
//
//		- When you call the SendCardSms operation to send card messages, the operation checks whether the mobile phone numbers of the recipients support card messages. If the mobile phone numbers do not support card messages, you can specify whether to enable rollback. Otherwise, the card message cannot be delivered.
//
// ### QPS limit
//
// You can call this operation up to 1,000 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SendCardSmsRequest
//
// @return SendCardSmsResponse
func (client *Client) SendCardSms(request *SendCardSmsRequest) (_result *SendCardSmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendCardSmsResponse{}
	_body, _err := client.SendCardSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送物流短信
//
// @param request - SendLogisticsSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLogisticsSmsResponse
func (client *Client) SendLogisticsSmsWithOptions(request *SendLogisticsSmsRequest, runtime *dara.RuntimeOptions) (_result *SendLogisticsSmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpressCompanyCode) {
		query["ExpressCompanyCode"] = request.ExpressCompanyCode
	}

	if !dara.IsNil(request.MailNo) {
		query["MailNo"] = request.MailNo
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlatformCompanyCode) {
		query["PlatformCompanyCode"] = request.PlatformCompanyCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateParam) {
		query["TemplateParam"] = request.TemplateParam
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendLogisticsSms"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendLogisticsSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送物流短信
//
// @param request - SendLogisticsSmsRequest
//
// @return SendLogisticsSmsResponse
func (client *Client) SendLogisticsSms(request *SendLogisticsSmsRequest) (_result *SendLogisticsSmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendLogisticsSmsResponse{}
	_body, _err := client.SendLogisticsSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a message. Before you call this operation, submit a message signature and message template, and make sure that the signature and template are approved.
//
// Description:
//
//	  This operation is mainly used to send a single message. In special scenarios, you can send multiple messages with the same content to a maximum of 1,000 mobile numbers. Note that group sending may be delayed.
//
//		- To send messages with different signatures and template content to multiple mobile numbers in a single request, call the [SendBatchSms](https://help.aliyun.com/document_detail/102364.html) operation.
//
//		- You are charged for using Alibaba Cloud Short Message Service (SMS) based on the amount of messages sent. For more information, see [Pricing](https://www.aliyun.com/price/product#/sms/detail).
//
//		- If your verification code signature and general-purpose signature have the same name, the system uses the general-purpose signature to send messages by default.
//
// @param request - SendSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSmsResponse
func (client *Client) SendSmsWithOptions(request *SendSmsRequest, runtime *dara.RuntimeOptions) (_result *SendSmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumbers) {
		query["PhoneNumbers"] = request.PhoneNumbers
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.SmsUpExtendCode) {
		query["SmsUpExtendCode"] = request.SmsUpExtendCode
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateParam) {
		query["TemplateParam"] = request.TemplateParam
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendSms"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendSmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message. Before you call this operation, submit a message signature and message template, and make sure that the signature and template are approved.
//
// Description:
//
//	  This operation is mainly used to send a single message. In special scenarios, you can send multiple messages with the same content to a maximum of 1,000 mobile numbers. Note that group sending may be delayed.
//
//		- To send messages with different signatures and template content to multiple mobile numbers in a single request, call the [SendBatchSms](https://help.aliyun.com/document_detail/102364.html) operation.
//
//		- You are charged for using Alibaba Cloud Short Message Service (SMS) based on the amount of messages sent. For more information, see [Pricing](https://www.aliyun.com/price/product#/sms/detail).
//
//		- If your verification code signature and general-purpose signature have the same name, the system uses the general-purpose signature to send messages by default.
//
// @param request - SendSmsRequest
//
// @return SendSmsResponse
func (client *Client) SendSms(request *SendSmsRequest) (_result *SendSmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendSmsResponse{}
	_body, _err := client.SendSmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reports the status of an OTP message to Alibaba Cloud SMS.
//
// Description:
//
// Metrics:
//
//   - Requested OTP messages
//
//   - Verified OTP messages
//
// An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
//
// > If you call the SmsConversion operation to query OTP conversion rates, your business may be affected. We recommend that you perform the following operations: 1. Call the SmsConversion operation in an asynchronous manner by configuring queues or events. 2. Manually degrade your services or use a circuit breaker to automatically degrade services.
//
// @param request - SmsConversionIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmsConversionIntlResponse
func (client *Client) SmsConversionIntlWithOptions(request *SmsConversionIntlRequest, runtime *dara.RuntimeOptions) (_result *SmsConversionIntlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversionTime) {
		query["ConversionTime"] = request.ConversionTime
	}

	if !dara.IsNil(request.Delivered) {
		query["Delivered"] = request.Delivered
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
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
		Action:      dara.String("SmsConversionIntl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmsConversionIntlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports the status of an OTP message to Alibaba Cloud SMS.
//
// Description:
//
// Metrics:
//
//   - Requested OTP messages
//
//   - Verified OTP messages
//
// An OTP conversion rate is calculated based on the following formula: OTP conversion rate = Number of verified OTP messages/Number of requested OTP messages.
//
// > If you call the SmsConversion operation to query OTP conversion rates, your business may be affected. We recommend that you perform the following operations: 1. Call the SmsConversion operation in an asynchronous manner by configuring queues or events. 2. Manually degrade your services or use a circuit breaker to automatically degrade services.
//
// @param request - SmsConversionIntlRequest
//
// @return SmsConversionIntlResponse
func (client *Client) SmsConversionIntl(request *SmsConversionIntlRequest) (_result *SmsConversionIntlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SmsConversionIntlResponse{}
	_body, _err := client.SmsConversionIntlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资质对客openAPI
//
// @param tmpReq - SubmitSmsQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSmsQualificationResponse
func (client *Client) SubmitSmsQualificationWithOptions(tmpReq *SubmitSmsQualificationRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmsQualificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SubmitSmsQualificationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BusinessLicensePics) {
		request.BusinessLicensePicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BusinessLicensePics, dara.String("BusinessLicensePics"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OtherFiles) {
		request.OtherFilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OtherFiles, dara.String("OtherFiles"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminIDCardExpDate) {
		query["AdminIDCardExpDate"] = request.AdminIDCardExpDate
	}

	if !dara.IsNil(request.AdminIDCardFrontFace) {
		query["AdminIDCardFrontFace"] = request.AdminIDCardFrontFace
	}

	if !dara.IsNil(request.AdminIDCardNo) {
		query["AdminIDCardNo"] = request.AdminIDCardNo
	}

	if !dara.IsNil(request.AdminIDCardPic) {
		query["AdminIDCardPic"] = request.AdminIDCardPic
	}

	if !dara.IsNil(request.AdminIDCardType) {
		query["AdminIDCardType"] = request.AdminIDCardType
	}

	if !dara.IsNil(request.AdminName) {
		query["AdminName"] = request.AdminName
	}

	if !dara.IsNil(request.AdminPhoneNo) {
		query["AdminPhoneNo"] = request.AdminPhoneNo
	}

	if !dara.IsNil(request.BusinessLicensePicsShrink) {
		query["BusinessLicensePics"] = request.BusinessLicensePicsShrink
	}

	if !dara.IsNil(request.BussinessLicenseExpDate) {
		query["BussinessLicenseExpDate"] = request.BussinessLicenseExpDate
	}

	if !dara.IsNil(request.CertifyCode) {
		query["CertifyCode"] = request.CertifyCode
	}

	if !dara.IsNil(request.CompanyName) {
		query["CompanyName"] = request.CompanyName
	}

	if !dara.IsNil(request.CompanyType) {
		query["CompanyType"] = request.CompanyType
	}

	if !dara.IsNil(request.LegalPersonIDCardNo) {
		query["LegalPersonIDCardNo"] = request.LegalPersonIDCardNo
	}

	if !dara.IsNil(request.LegalPersonIDCardType) {
		query["LegalPersonIDCardType"] = request.LegalPersonIDCardType
	}

	if !dara.IsNil(request.LegalPersonIdCardBackSide) {
		query["LegalPersonIdCardBackSide"] = request.LegalPersonIdCardBackSide
	}

	if !dara.IsNil(request.LegalPersonIdCardEffTime) {
		query["LegalPersonIdCardEffTime"] = request.LegalPersonIdCardEffTime
	}

	if !dara.IsNil(request.LegalPersonIdCardFrontSide) {
		query["LegalPersonIdCardFrontSide"] = request.LegalPersonIdCardFrontSide
	}

	if !dara.IsNil(request.LegalPersonName) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !dara.IsNil(request.OrganizationCode) {
		query["OrganizationCode"] = request.OrganizationCode
	}

	if !dara.IsNil(request.OtherFilesShrink) {
		query["OtherFiles"] = request.OtherFilesShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationName) {
		query["QualificationName"] = request.QualificationName
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UseBySelf) {
		query["UseBySelf"] = request.UseBySelf
	}

	if !dara.IsNil(request.WhetherShare) {
		query["WhetherShare"] = request.WhetherShare
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSmsQualification"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSmsQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资质对客openAPI
//
// @param request - SubmitSmsQualificationRequest
//
// @return SubmitSmsQualificationResponse
func (client *Client) SubmitSmsQualification(request *SubmitSmsQualificationRequest) (_result *SubmitSmsQualificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSmsQualificationResponse{}
	_body, _err := client.SubmitSmsQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches tags to a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches tags to a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes tags from a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes tags from a message template.
//
// Description:
//
// ### QPS limit
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改验证码签名
//
// @param request - UpdateExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtCodeSignResponse
func (client *Client) UpdateExtCodeSignWithOptions(request *UpdateExtCodeSignRequest, runtime *dara.RuntimeOptions) (_result *UpdateExtCodeSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExistExtCode) {
		query["ExistExtCode"] = request.ExistExtCode
	}

	if !dara.IsNil(request.NewExtCode) {
		query["NewExtCode"] = request.NewExtCode
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

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExtCodeSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExtCodeSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改验证码签名
//
// @param request - UpdateExtCodeSignRequest
//
// @return UpdateExtCodeSignResponse
func (client *Client) UpdateExtCodeSign(request *UpdateExtCodeSignRequest) (_result *UpdateExtCodeSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateExtCodeSignResponse{}
	_body, _err := client.UpdateExtCodeSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改资质对客openAPI
//
// @param tmpReq - UpdateSmsQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsQualificationResponse
func (client *Client) UpdateSmsQualificationWithOptions(tmpReq *UpdateSmsQualificationRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmsQualificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSmsQualificationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BusinessLicensePics) {
		request.BusinessLicensePicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BusinessLicensePics, dara.String("BusinessLicensePics"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OtherFiles) {
		request.OtherFilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OtherFiles, dara.String("OtherFiles"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminIDCardExpDate) {
		query["AdminIDCardExpDate"] = request.AdminIDCardExpDate
	}

	if !dara.IsNil(request.AdminIDCardFrontFace) {
		query["AdminIDCardFrontFace"] = request.AdminIDCardFrontFace
	}

	if !dara.IsNil(request.AdminIDCardNo) {
		query["AdminIDCardNo"] = request.AdminIDCardNo
	}

	if !dara.IsNil(request.AdminIDCardPic) {
		query["AdminIDCardPic"] = request.AdminIDCardPic
	}

	if !dara.IsNil(request.AdminIDCardType) {
		query["AdminIDCardType"] = request.AdminIDCardType
	}

	if !dara.IsNil(request.AdminName) {
		query["AdminName"] = request.AdminName
	}

	if !dara.IsNil(request.AdminPhoneNo) {
		query["AdminPhoneNo"] = request.AdminPhoneNo
	}

	if !dara.IsNil(request.BusinessLicensePicsShrink) {
		query["BusinessLicensePics"] = request.BusinessLicensePicsShrink
	}

	if !dara.IsNil(request.BussinessLicenseExpDate) {
		query["BussinessLicenseExpDate"] = request.BussinessLicenseExpDate
	}

	if !dara.IsNil(request.CertifyCode) {
		query["CertifyCode"] = request.CertifyCode
	}

	if !dara.IsNil(request.CompanyName) {
		query["CompanyName"] = request.CompanyName
	}

	if !dara.IsNil(request.LegalPersonIDCardNo) {
		query["LegalPersonIDCardNo"] = request.LegalPersonIDCardNo
	}

	if !dara.IsNil(request.LegalPersonIDCardType) {
		query["LegalPersonIDCardType"] = request.LegalPersonIDCardType
	}

	if !dara.IsNil(request.LegalPersonIdCardBackSide) {
		query["LegalPersonIdCardBackSide"] = request.LegalPersonIdCardBackSide
	}

	if !dara.IsNil(request.LegalPersonIdCardEffTime) {
		query["LegalPersonIdCardEffTime"] = request.LegalPersonIdCardEffTime
	}

	if !dara.IsNil(request.LegalPersonIdCardFrontSide) {
		query["LegalPersonIdCardFrontSide"] = request.LegalPersonIdCardFrontSide
	}

	if !dara.IsNil(request.LegalPersonName) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OtherFilesShrink) {
		query["OtherFiles"] = request.OtherFilesShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationGroupId) {
		query["QualificationGroupId"] = request.QualificationGroupId
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
		Action:      dara.String("UpdateSmsQualification"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmsQualificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改资质对客openAPI
//
// @param request - UpdateSmsQualificationRequest
//
// @return UpdateSmsQualificationResponse
func (client *Client) UpdateSmsQualification(request *UpdateSmsQualificationRequest) (_result *UpdateSmsQualificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSmsQualificationResponse{}
	_body, _err := client.UpdateSmsQualificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Text SMS Signature
//
// Description:
//
// - For details about the changes of this new interface and the original one, please refer to [Announcement on the Update of SMS Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Only signatures that have not passed the review can be modified. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm) and call this interface to modify and resubmit for review after modification.
//
// - Signature information applied through the interface will be synchronized in the SMS service console. For operations related to signatures in the console, please see [SMS Signatures](https://help.aliyun.com/zh/sms/user-guide/create-signatures?spm).
//
// @param tmpReq - UpdateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsSignResponse
func (client *Client) UpdateSmsSignWithOptions(tmpReq *UpdateSmsSignRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmsSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSmsSignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MoreData) {
		request.MoreDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MoreData, dara.String("MoreData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIcpRecordId) {
		query["AppIcpRecordId"] = request.AppIcpRecordId
	}

	if !dara.IsNil(request.ApplySceneContent) {
		query["ApplySceneContent"] = request.ApplySceneContent
	}

	if !dara.IsNil(request.AuthorizationLetterId) {
		query["AuthorizationLetterId"] = request.AuthorizationLetterId
	}

	if !dara.IsNil(request.MoreDataShrink) {
		query["MoreData"] = request.MoreDataShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.SignSource) {
		query["SignSource"] = request.SignSource
	}

	if !dara.IsNil(request.SignType) {
		query["SignType"] = request.SignType
	}

	if !dara.IsNil(request.ThirdParty) {
		query["ThirdParty"] = request.ThirdParty
	}

	if !dara.IsNil(request.TrademarkId) {
		query["TrademarkId"] = request.TrademarkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmsSign"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Text SMS Signature
//
// Description:
//
// - For details about the changes of this new interface and the original one, please refer to [Announcement on the Update of SMS Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Only signatures that have not passed the review can be modified. Please refer to [Handling Suggestions for Failed SMS Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm) and call this interface to modify and resubmit for review after modification.
//
// - Signature information applied through the interface will be synchronized in the SMS service console. For operations related to signatures in the console, please see [SMS Signatures](https://help.aliyun.com/zh/sms/user-guide/create-signatures?spm).
//
// @param request - UpdateSmsSignRequest
//
// @return UpdateSmsSignResponse
func (client *Client) UpdateSmsSign(request *UpdateSmsSignRequest) (_result *UpdateSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSmsSignResponse{}
	_body, _err := client.UpdateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Text SMS Template
//
// Description:
//
// - For details about the changes of this new interface compared to the original one, please refer to [Announcement on SMS Service Update: Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Only templates that have not passed the review can be modified. Please refer to [Handling Suggestions for Failed SMS Template Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.4bf5561ajcFtMQ) and call this interface to modify and resubmit for review.
//
// - Modifications made through the interface will be synchronized in the SMS service console. For related operations on templates in the console, see [SMS Templates](https://help.aliyun.com/zh/sms/user-guide/message-templates/?spm=a2c4g.11186623.0.0.35a947464Itaxp).
//
// ### QPS Limit
//
// The single-user QPS limit for this interface is 1000 times/second. Exceeding this limit will result in API throttling, which may impact your business. Please make calls reasonably.
//
// @param tmpReq - UpdateSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsTemplateResponse
func (client *Client) UpdateSmsTemplateWithOptions(tmpReq *UpdateSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmsTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSmsTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MoreData) {
		request.MoreDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MoreData, dara.String("MoreData"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplySceneContent) {
		query["ApplySceneContent"] = request.ApplySceneContent
	}

	if !dara.IsNil(request.IntlType) {
		query["IntlType"] = request.IntlType
	}

	if !dara.IsNil(request.MoreDataShrink) {
		query["MoreData"] = request.MoreDataShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RelatedSignName) {
		query["RelatedSignName"] = request.RelatedSignName
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateContent) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateRule) {
		query["TemplateRule"] = request.TemplateRule
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.TrafficDriving) {
		query["TrafficDriving"] = request.TrafficDriving
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Text SMS Template
//
// Description:
//
// - For details about the changes of this new interface compared to the original one, please refer to [Announcement on SMS Service Update: Signature & Template Interfaces](https://help.aliyun.com/zh/sms/product-overview/announcement-on-sms-service-update-signature-template-interface).
//
// - Only templates that have not passed the review can be modified. Please refer to [Handling Suggestions for Failed SMS Template Reviews](https://help.aliyun.com/zh/sms/user-guide/causes-of-application-failures-and-suggestions?spm=a2c4g.11186623.0.0.4bf5561ajcFtMQ) and call this interface to modify and resubmit for review.
//
// - Modifications made through the interface will be synchronized in the SMS service console. For related operations on templates in the console, see [SMS Templates](https://help.aliyun.com/zh/sms/user-guide/message-templates/?spm=a2c4g.11186623.0.0.35a947464Itaxp).
//
// ### QPS Limit
//
// The single-user QPS limit for this interface is 1000 times/second. Exceeding this limit will result in API throttling, which may impact your business. Please make calls reasonably.
//
// @param request - UpdateSmsTemplateRequest
//
// @return UpdateSmsTemplateResponse
func (client *Client) UpdateSmsTemplate(request *UpdateSmsTemplateRequest) (_result *UpdateSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSmsTemplateResponse{}
	_body, _err := client.UpdateSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送手机验证码
//
// @param request - ValidPhoneCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidPhoneCodeResponse
func (client *Client) ValidPhoneCodeWithOptions(request *ValidPhoneCodeRequest, runtime *dara.RuntimeOptions) (_result *ValidPhoneCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertifyCode) {
		query["CertifyCode"] = request.CertifyCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNo) {
		query["PhoneNo"] = request.PhoneNo
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
		Action:      dara.String("ValidPhoneCode"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidPhoneCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送手机验证码
//
// @param request - ValidPhoneCodeRequest
//
// @return ValidPhoneCodeResponse
func (client *Client) ValidPhoneCode(request *ValidPhoneCodeRequest) (_result *ValidPhoneCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidPhoneCodeResponse{}
	_body, _err := client.ValidPhoneCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 物流短信运单号校验
//
// @param request - VerifyLogisticsSmsMailNoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyLogisticsSmsMailNoResponse
func (client *Client) VerifyLogisticsSmsMailNoWithOptions(request *VerifyLogisticsSmsMailNoRequest, runtime *dara.RuntimeOptions) (_result *VerifyLogisticsSmsMailNoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpressCompanyCode) {
		query["ExpressCompanyCode"] = request.ExpressCompanyCode
	}

	if !dara.IsNil(request.MailNo) {
		query["MailNo"] = request.MailNo
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlatformCompanyCode) {
		query["PlatformCompanyCode"] = request.PlatformCompanyCode
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
		Action:      dara.String("VerifyLogisticsSmsMailNo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyLogisticsSmsMailNoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 物流短信运单号校验
//
// @param request - VerifyLogisticsSmsMailNoRequest
//
// @return VerifyLogisticsSmsMailNoResponse
func (client *Client) VerifyLogisticsSmsMailNo(request *VerifyLogisticsSmsMailNoRequest) (_result *VerifyLogisticsSmsMailNoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyLogisticsSmsMailNoResponse{}
	_body, _err := client.VerifyLogisticsSmsMailNoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
