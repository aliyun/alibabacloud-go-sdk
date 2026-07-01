// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 添加验证码签名信息
//
// @param request - AddExtCodeSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddExtCodeSignResponse
func (client *Client) AddExtCodeSignWithContext(ctx context.Context, request *AddExtCodeSignRequest, runtime *dara.RuntimeOptions) (_result *AddExtCodeSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建/编辑5G消息固定菜单
//
// @param request - AddRcsSignMenuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRcsSignMenuResponse
func (client *Client) AddRcsSignMenuWithContext(ctx context.Context, request *AddRcsSignMenuRequest, runtime *dara.RuntimeOptions) (_result *AddRcsSignMenuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MenuContent) {
		query["MenuContent"] = request.MenuContent
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRcsSignMenu"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRcsSignMenuResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
//	Notice:
//
// Short Message Service does not currently support this API operation.
//
// - You can create up to 3,000 short URLs per calendar day.
//
// - After a short URL is generated, it must pass a security review, which typically takes 10 minutes to 2 hours. You cannot access the short URL until it passes this review.
//
// @param request - AddShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddShortUrlResponse
func (client *Client) AddShortUrlWithContext(ctx context.Context, request *AddShortUrlRequest, runtime *dara.RuntimeOptions) (_result *AddShortUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API has been discontinued.
//
// Description:
//
// - In accordance with the regulations of the Ministry of Industry and Information Technology (MIIT) and the [relevant requirements](https://help.aliyun.com/document_detail/2806975.html) of carriers, Alibaba Cloud has made functional modifications to signature-related APIs. To improve the review efficiency and approval rate of your signatures, use the new API [CreateSmsSign - Apply for an SMS signature](https://help.aliyun.com/document_detail/2807427.html).
//
// - An individual user can apply for one signature per natural day under the same Alibaba Cloud account. Enterprise users have no limit on the number of applications. For details about the differences in rights and interests between individual users and enterprise users, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// - The signature information applied for through the API is synchronized to the SMS console. For operations on signatures in the console, see [SMS signatures](https://help.aliyun.com/document_detail/108073.html).
//
// - After you submit a signature application, you can call the [QuerySmsSign](https://help.aliyun.com/document_detail/419283.html) API to query the review status and details of the signature. You can also [configure receipt messages](https://help.aliyun.com/document_detail/101508.html) and use [SignSmsReport](https://help.aliyun.com/document_detail/120998.html) to obtain signature review status messages.
//
// ### QPS limit
//
// The single-user QPS limit for this API is 1 call per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Call this API at a reasonable rate.
//
// @param request - AddSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSmsSignResponse
func (client *Client) AddSmsSignWithContext(ctx context.Context, request *AddSmsSignRequest, runtime *dara.RuntimeOptions) (_result *AddSmsSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// An SMS template is the detailed content received by the recipient, including variables and template content. You can apply for verification code, notification, or promotional SMS templates based on your business needs. SMS can only be sent after the template is approved.
//
// Description:
//
// - In accordance with the regulations of the Ministry of Industry and Information Technology and the [related requirements](https://help.aliyun.com/document_detail/2806975.html) of carriers, Alibaba Cloud has revamped the functionality of template-related APIs. To improve the review efficiency and approval rate of your templates, please use the new operation [CreateSmsTemplate - Apply for SMS template](https://help.aliyun.com/document_detail/2807431.html).
//
// - You can submit a maximum of 100 SMS template applications per natural day via the API. It is recommended that each application be submitted at intervals of at least 30 seconds. There is no limit on the number of submissions when applying for SMS templates through the [console](https://dysms.console.aliyun.com/domestic/text/template).
//
// - Template information applied for through the API is synchronized to the SMS service console. For related template operations in the console, see [SMS templates](https://help.aliyun.com/document_detail/108085.html).
//
// - After submitting the template application, you can query the template review status and details through the [QuerySmsTemplate](https://help.aliyun.com/document_detail/419289.html) operation. You can also [configure receipt messages](https://help.aliyun.com/document_detail/101508.html) and obtain the template review status messages through [TemplateSmsReport](https://help.aliyun.com/document_detail/120999.html).
//
// - Domestic SMS templates and international/Hong Kong, Macao, and Taiwan SMS templates are not interchangeable (cannot be mixed). Please apply for templates based on your business usage scenarios.
//
// - Only enterprise-verified users can apply for promotional SMS and international/Hong Kong, Macao, and Taiwan messages. For details about the differences between individual and enterprise user privileges, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// ### QPS limits
//
// The per-user QPS limit for this operation is 1,000 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call the operation reasonably.
//
// @param request - AddSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSmsTemplateResponse
func (client *Client) AddSmsTemplateWithContext(ctx context.Context, request *AddSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddSmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the qualification and authorization letter for a signature.
//
// @param request - ChangeSignatureQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeSignatureQualificationResponse
func (client *Client) ChangeSignatureQualificationWithContext(ctx context.Context, request *ChangeSignatureQualificationRequest, runtime *dara.RuntimeOptions) (_result *ChangeSignatureQualificationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether phone numbers support card SMS.
//
// Description:
//
// - Alibaba Cloud accounts that have not activated card SMS cannot call this API.
//
// - Card SMS is currently in the internal invitation phase. Contact your Alibaba Cloud account manager to apply for activation or [contact Alibaba Cloud pre-sales](https://help.aliyun.com/document_detail/464625.html).
//
// - We recommend that you use the new API [QueryMobilesCardSupport](~~QueryMobilesCardSupport~~) to query whether phone numbers support card SMS.
//
// ### QPS limit
//
// The per-user QPS limit for this API is 2,000 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make calls reasonably.
//
// @param request - CheckMobilesCardSupportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckMobilesCardSupportResponse
func (client *Client) CheckMobilesCardSupportWithContext(ctx context.Context, request *CheckMobilesCardSupportRequest, runtime *dara.RuntimeOptions) (_result *CheckMobilesCardSupportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports SMS conversion rate statistics to the Alibaba Cloud SMS platform.
//
// Description:
//
// 指标说明：转化率=OTP 转化量/OTP 发送量。
//
// - OTP发送量：验证码发送量。
//
// - OTP转化量：验证码转换量。（用户成功获取验证码，并进行回传）
//
// >转化率反馈功能会对业务系统有一定的侵入性，为了防止调用转化率 API 的抖动影响业务逻辑，请考虑：
//
// >- 使用异步模式（例如：队列或事件驱动）调用 API。
//
// >- 添加可降级的方案保护业务逻辑（例如：手动降级开工或者使用断路器自动降级）。
//
// @param request - ConversionDataIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConversionDataIntlResponse
func (client *Client) ConversionDataIntlWithContext(ctx context.Context, request *ConversionDataIntlRequest, runtime *dara.RuntimeOptions) (_result *ConversionDataIntlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a card SMS template.
//
// Description:
//
// - The card SMS feature is currently available by invitation only. To enable this feature, contact your Alibaba Cloud business manager or our [pre-sales consultation](https://help.aliyun.com/document_detail/464625.html?spm=a2c4g.11186623.0.0.213273d4Xe6UEu#section-81n-72q-ybm) team.
//
// - This operation saves a card SMS template, submits it to mobile phone vendors for review, and returns a template code.
//
// - If a card SMS template contains a type or event that a vendor does not support, the system does not submit the template to that vendor for review. For more information, see [Supported template types by vendor](https://help.aliyun.com/document_detail/434611.html).
//
// - For more examples of card SMS templates, see [Card SMS template examples](https://help.aliyun.com/document_detail/435361.html).
//
// ### QPS limit
//
// The QPS limit for a single user is 300. API calls that exceed this limit are throttled, which may impact your business. Plan your calls accordingly.
//
// @param tmpReq - CreateCardSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCardSmsTemplateResponse
func (client *Client) CreateCardSmsTemplateWithContext(ctx context.Context, tmpReq *CreateCardSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateCardSmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an order to add, update, or delete a digital message signature.
//
// Description:
//
// Creates, updates, or deletes a signature.
//
// @param tmpReq - CreateDigitalSignOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDigitalSignOrderResponse
func (client *Client) CreateDigitalSignOrderWithContext(ctx context.Context, tmpReq *CreateDigitalSignOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateDigitalSignOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDigitalSignOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderContext) {
		request.OrderContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderContext, dara.String("OrderContext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtendMessage) {
		query["ExtendMessage"] = request.ExtendMessage
	}

	if !dara.IsNil(request.OrderContextShrink) {
		query["OrderContext"] = request.OrderContextShrink
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.QualificationVersion) {
		query["QualificationVersion"] = request.QualificationVersion
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SignId) {
		query["SignId"] = request.SignId
	}

	if !dara.IsNil(request.SignIndustry) {
		query["SignIndustry"] = request.SignIndustry
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.SignSource) {
		query["SignSource"] = request.SignSource
	}

	if !dara.IsNil(request.Submitter) {
		query["Submitter"] = request.Submitter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDigitalSignOrder"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDigitalSignOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a digital SMS template.
//
// Description:
//
// Use this operation to create a reusable template for your digital SMS messages.
//
// @param request - CreateDigitalSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDigitalSmsTemplateResponse
func (client *Client) CreateDigitalSmsTemplateWithContext(ctx context.Context, request *CreateDigitalSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateDigitalSmsTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateContents) {
		query["TemplateContents"] = request.TemplateContents
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDigitalSmsTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDigitalSmsTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建终端5G适配情况查询任务
//
// @param request - CreateRCSMobileCapableTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRCSMobileCapableTaskResponse
func (client *Client) CreateRCSMobileCapableTaskWithContext(ctx context.Context, request *CreateRCSMobileCapableTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRCSMobileCapableTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PhoneNumbersFile) {
		query["PhoneNumbersFile"] = request.PhoneNumbersFile
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRCSMobileCapableTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRCSMobileCapableTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建5G消息模板
//
// @param request - CreateRCSTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRCSTemplateResponse
func (client *Client) CreateRCSTemplateWithContext(ctx context.Context, request *CreateRCSTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateRCSTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RelatedSignNames) {
		query["RelatedSignNames"] = request.RelatedSignNames
	}

	if !dara.IsNil(request.TemplateContent) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !dara.IsNil(request.TemplateFormat) {
		query["TemplateFormat"] = request.TemplateFormat
	}

	if !dara.IsNil(request.TemplateMenu) {
		query["TemplateMenu"] = request.TemplateMenu
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRCSTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRCSTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmartShortUrlResponse
func (client *Client) CreateSmartShortUrlWithContext(ctx context.Context, request *CreateSmartShortUrlRequest, runtime *dara.RuntimeOptions) (_result *CreateSmartShortUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The process for using a live app as a signature source has changed. If you use an app as the signature source, you must call this operation to create an ICP filing record for it.
//
// Description:
//
//	Notice: To use a **live app*	- as a signature source, you must now provide its ICP filing information. This requires you to upload a screenshot of the app\\"s ICP filing.
//
// @param request - CreateSmsAppIcpRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsAppIcpRecordResponse
func (client *Client) CreateSmsAppIcpRecordWithContext(ctx context.Context, request *CreateSmsAppIcpRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsAppIcpRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppApprovalDate) {
		query["AppApprovalDate"] = request.AppApprovalDate
	}

	if !dara.IsNil(request.AppIcpLicenseNumber) {
		query["AppIcpLicenseNumber"] = request.AppIcpLicenseNumber
	}

	if !dara.IsNil(request.AppIcpRecordPic) {
		query["AppIcpRecordPic"] = request.AppIcpRecordPic
	}

	if !dara.IsNil(request.AppPrincipalUnitName) {
		query["AppPrincipalUnitName"] = request.AppPrincipalUnitName
	}

	if !dara.IsNil(request.AppRuntimePic) {
		query["AppRuntimePic"] = request.AppRuntimePic
	}

	if !dara.IsNil(request.AppServiceName) {
		query["AppServiceName"] = request.AppServiceName
	}

	if !dara.IsNil(request.AppStoreDownloadPic) {
		query["AppStoreDownloadPic"] = request.AppStoreDownloadPic
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
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
		Action:      dara.String("CreateSmsAppIcpRecord"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsAppIcpRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If the qualification is intended for use by a third party or the requested signature involves third-party rights, you must obtain third-party authorization and create an authorization letter before submitting the application.
//
// Description:
//
// - Before use, please read the [Authorization Letter Specifications](https://help.aliyun.com/document_detail/56741.html). Download the [Authorization Letter Template](https://help-static-aliyun-doc.aliyuncs.com/file-manage-files/zh-CN/20250414/bvpcmo/%E6%8E%88%E6%9D%83%E5%A7%94%E6%89%98%E4%B9%A6%E6%A8%A1%E7%89%88.doc), fill it out and stamp it according to the specifications, and then upload it.
//
// - The authorization letter you create can be used when applying for SMS qualifications or SMS signatures. If your qualification or signature is intended for use by a third party, you must create and submit an authorization letter.
//
// - After creating an authorization letter, you can call [QuerySmsAuthorizationLetter](~~QuerySmsAuthorizationLetter~~) to query the details of the created authorization letter. Authorization letter information created through the API is synchronized to the Short Message Service console.
//
// @param tmpReq - CreateSmsAuthorizationLetterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsAuthorizationLetterResponse
func (client *Client) CreateSmsAuthorizationLetterWithContext(ctx context.Context, tmpReq *CreateSmsAuthorizationLetterRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsAuthorizationLetterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An SMS signature identifies the sender of an SMS message. Before sending SMS messages, you must apply for a signature and a template. The system prepends the approved SMS signature to the message content and sends them together to the recipient.
//
// Description:
//
// - For details about the changes between the new and original operations, see [Announcement on updating signature and template operations for Short Message Service](https://help.aliyun.com/document_detail/2806975.html).
//
// - Users who verify your identity - Individual account can apply for one formal signature per calendar day per Alibaba Cloud account. Users who verify your identity - Enterprise account currently have no such limit. For details about the differences between individual and enterprise user privileges, see [Before you begin](https://help.aliyun.com/document_detail/55324.html).
//
// - Read the [Signature specifications](https://help.aliyun.com/document_detail/108076.html) to learn about the SMS signature review standards.
//
// - Signatures applied for through the API are synchronized to the Short Message Service console. For console-related operations, see [SMS signatures](https://help.aliyun.com/document_detail/108073.html).
//
// - After you submit a signature application, you can call the [GetSmsSign](https://help.aliyun.com/document_detail/2807429.html) operation to query the signature review status and details. You can also [configure receipt messages](https://help.aliyun.com/document_detail/101508.html) and use [SignSmsReport](https://help.aliyun.com/document_detail/120998.html) to obtain the signature review status messages.
//
// @param tmpReq - CreateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSignWithContext(ctx context.Context, tmpReq *CreateSmsSignRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A message template defines the content of an SMS message. This content includes the message text and any variables. You can create templates for various business needs, such as sending verification codes, notifications, or promotional messages. A template must be approved before you can use it to send messages.
//
// Description:
//
// - For details on the API changes for signatures and templates, see the [Announcement on Signature and Template API Updates for Short Message Service](https://help.aliyun.com/document_detail/2806975.html).
//
// - Wait at least 30 seconds between API calls when applying for a message template.
//
// - Message templates you apply for via the API are synchronized to the Short Message Service console. For details on how to manage message templates in the console, see [Message templates](https://help.aliyun.com/document_detail/108085.html).
//
// - After you submit a template for review, you can call the [GetSmsTemplate](https://help.aliyun.com/document_detail/2807433.html) API to query the template\\"s review status and details. You can also [configure status reports](https://help.aliyun.com/document_detail/101508.html) to receive the template\\"s review status through [TemplateSmsReport](https://help.aliyun.com/document_detail/120999.html).
//
// - Message templates for Chinese mainland messages and international messages are not interchangeable. Apply for message templates based on your use case.
//
// - Only enterprise-verified users can apply for message templates for promotional messages and international messages. For details on the permission differences between individual and enterprise users, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
//
// @param tmpReq - CreateSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsTemplateResponse
func (client *Client) CreateSmsTemplateWithContext(ctx context.Context, tmpReq *CreateSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trademark entity. This operation is used when you need to upload trademark information when the signature source is set to trademark.
//
// Description:
//
// The trademark must be searchable on the China Trademark Network of the Trademark Office of the China National Intellectual Property Administration, and the trademark owner must match the enterprise name.
//
// @param request - CreateSmsTrademarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsTrademarkResponse
func (client *Client) CreateSmsTrademarkWithContext(ctx context.Context, request *CreateSmsTrademarkRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsTrademarkResponse, _err error) {
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

	if !dara.IsNil(request.TrademarkApplicantName) {
		query["TrademarkApplicantName"] = request.TrademarkApplicantName
	}

	if !dara.IsNil(request.TrademarkEffExpDate) {
		query["TrademarkEffExpDate"] = request.TrademarkEffExpDate
	}

	if !dara.IsNil(request.TrademarkName) {
		query["TrademarkName"] = request.TrademarkName
	}

	if !dara.IsNil(request.TrademarkPic) {
		query["TrademarkPic"] = request.TrademarkPic
	}

	if !dara.IsNil(request.TrademarkRegistrationNumber) {
		query["TrademarkRegistrationNumber"] = request.TrademarkRegistrationNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmsTrademark"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsTrademarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExtCodeSignResponse
func (client *Client) DeleteExtCodeSignWithContext(ctx context.Context, request *DeleteExtCodeSignRequest, runtime *dara.RuntimeOptions) (_result *DeleteExtCodeSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a short URL. After deletion, the short URL is no longer usable and cannot be resolved to the source URL.
//
// Description:
//
//	Notice:
//
// Short Message Service does not currently support this API operation.
//
// ### QPS limit
//
// The QPS limit for a single user is 100. Calls that exceed this limit are subject to rate limiting, which may affect your business. To prevent disruptions, call this operation at a reasonable frequency.
//
// @param request - DeleteShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteShortUrlResponse
func (client *Client) DeleteShortUrlWithContext(ctx context.Context, request *DeleteShortUrlRequest, runtime *dara.RuntimeOptions) (_result *DeleteShortUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you no longer use an SMS qualification or need to delete it for other reasons, call this API or delete the SMS qualification in the Short Message Service console.
//
// Description:
//
//	Warning: Once a qualification is deleted, it cannot be restored or selected when you apply for signatures later. Proceed with caution.
//
// - Qualifications under review cannot be modified or deleted. You can withdraw the application in the Short Message Service [console](https://dysms.console.aliyun.com/domestic/text/qualification) before performing the operation.
//
// - Approved qualifications that have been bound to signatures cannot be deleted.
//
// - Rejected qualifications can be directly resubmitted for review after you [modify the qualification information](~~UpdateSmsQualification~~).
//
// @param request - DeleteSmsQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmsQualificationResponse
func (client *Client) DeleteSmsQualificationWithContext(ctx context.Context, request *DeleteSmsQualificationRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmsQualificationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you no longer use an SMS signature and need to delete it, call this operation or delete the SMS signature in the SMS Service console.
//
// Description:
//
// - You can delete signatures that are in the Withdrawn, Failed, or Approved state. You cannot delete signatures that are in the Pending Approval state.
//
// - Deleted SMS signatures cannot be recovered, and the signature can no longer be used to send SMS messages. Proceed with caution.
//
// - Signature deletion operations performed via the API are synchronized to the SMS Service console. For information about how to manage templates in the console, see [SMS signatures](https://help.aliyun.com/document_detail/108073.html).
//
// @param request - DeleteSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmsSignResponse
func (client *Client) DeleteSmsSignWithContext(ctx context.Context, request *DeleteSmsSignRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmsSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an SMS template that you no longer need.
//
// Description:
//
// - 支持删除已撤回、审核失败或审核通过的模板，审核中的短信模板不支持删除。
//
// - 删除短信模板后不可恢复，且后续不可再使用该模板发送短信，请谨慎操作。
//
// - 通过接口删除模板的操作会在短信服务控制台同步，在控制台对模板的相关操作，请参见[短信模板](https://help.aliyun.com/document_detail/108085.html)。
//
// ### QPS限制
//
// 本接口的单用户QPS限制为1000次/秒。超过限制，API调用会被限流，这可能会影响您的业务，请合理调用。
//
// @param request - DeleteSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmsTemplateResponse
func (client *Client) DeleteSmsTemplateWithContext(ctx context.Context, request *DeleteSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the card SMS sending records and sending status of a single phone number.
//
// @param request - GetCardSmsDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardSmsDetailsResponse
func (client *Client) GetCardSmsDetailsWithContext(ctx context.Context, request *GetCardSmsDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetCardSmsDetailsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the short URL for a card message.
//
// Description:
//
// - 目前卡片短信在内部邀约阶段，请联系您的阿里云商务经理申请开通或联系[阿里云售前咨询](https://help.aliyun.com/document_detail/464625.html?spm=a2c4g.11186623.0.0.213273d4Xe6UEu#section-81n-72q-ybm)。
//
// ### QPS限制
//
// - 本接口的单用户QPS限制为1000次/秒。超过限制，API调用会被限流，这可能会影响您的业务，请合理调用。
//
// @param request - GetCardSmsLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCardSmsLinkResponse
func (client *Client) GetCardSmsLinkWithContext(ctx context.Context, request *GetCardSmsLinkRequest, runtime *dara.RuntimeOptions) (_result *GetCardSmsLinkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts images and videos uploaded to OSS storage for card SMS into resource data for unified management, and returns a resource ID. You can manage the returned resource ID as needed.
//
// Description:
//
// ### QPS限制
//
// 本接口的单用户QPS限制为300次/秒。超过限制，API调用会被限流，这可能会影响您的业务，请合理调用。
//
// @param request - GetMediaResourceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaResourceIdResponse
func (client *Client) GetMediaResourceIdWithContext(ctx context.Context, request *GetMediaResourceIdRequest, runtime *dara.RuntimeOptions) (_result *GetMediaResourceIdResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains OSS resource configuration information, which will be used for subsequent OSS file upload operations.
//
// Description:
//
// - When you create a signature or template, you can upload materials such as login pages with links, backend page screenshots, software copyrights, and supplementary agreements. These materials help reviewers understand the details of your business. If you have multiple materials, you can combine them into one file. The supported formats are png, jpg, jpeg, doc, docx, and pdf.
//
// - Additional materials required for creating a signature or template can be uploaded to the OSS file system for storage. For information about file upload operations, see [Upload files to OSS](https://help.aliyun.com/document_detail/2833114.html).
//
// @param request - GetOSSInfoForUploadFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOSSInfoForUploadFileResponse
func (client *Client) GetOSSInfoForUploadFileWithContext(ctx context.Context, request *GetOSSInfoForUploadFileRequest, runtime *dara.RuntimeOptions) (_result *GetOSSInfoForUploadFileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the OSS resource configuration information for qualification materials. This configuration information will be used for subsequent uploads of qualification files such as authorization letters and enterprise certificates.
//
// Description:
//
// - When you apply for a qualification or signature, if the purpose is for other use or involves a third party, you must provide an [authorization letter](https://help.aliyun.com/document_detail/56741.html).
//
// - After you use this API to obtain the OSS resource configuration information, upload the related qualification materials through OSS. For more information, see [Upload files through OSS](https://help.aliyun.com/document_detail/2833114.html).
//
// - The names of files to be uploaded cannot contain Chinese characters or special symbols.
//
// @param request - GetQualificationOssInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualificationOssInfoResponse
func (client *Client) GetQualificationOssInfoWithContext(ctx context.Context, request *GetQualificationOssInfoRequest, runtime *dara.RuntimeOptions) (_result *GetQualificationOssInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询5g短信详情
//
// @param request - GetRCSSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRCSSignatureResponse
func (client *Client) GetRCSSignatureWithContext(ctx context.Context, request *GetRCSSignatureRequest, runtime *dara.RuntimeOptions) (_result *GetRCSSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRCSSignature"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRCSSignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the OSS information for OCR.
//
// @param request - GetSmsOcrOssInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsOcrOssInfoResponse
func (client *Client) GetSmsOcrOssInfoWithContext(ctx context.Context, request *GetSmsOcrOssInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSmsOcrOssInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the review details of a signature after you apply for it.
//
// Description:
//
// - 仅可查询**首次创建**的签名资料或者**最新审核通过**的资料。
//
// - 新接口和原接口变更的公告详情请参见[关于短信服务更新签名&模板接口的公告](https://help.aliyun.com/document_detail/2806975.html)。
//
// - 审核时间：一般情况下，签名提交后，阿里云预计在 2 个小时内审核完成（审核工作时间：周一至周日 9:00~21:00，法定节假日顺延）。
//
// - 如果签名未通过审核，会返回审核失败的原因，请参考[短信审核失败的处理建议](https://help.aliyun.com/document_detail/65990.html)，调用[UpdateSmsSign](https://help.aliyun.com/document_detail/2807428.html)接口或在[签名管理](https://dysms.console.aliyun.com/domestic/text/sign)页面修改未审核通过的短信签名。
//
// - [QuerySmsSignList](~~QuerySmsSignList~~)接口可以查询您账号下的所有签名，包括签名审核状态、签名类型、签名名称等。
//
// - 本接口的单用户QPS限制为150次/秒。超过限制，API调用将会被限流，这可能会影响您的业务，请合理调用。
//
// @param request - GetSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsSignResponse
func (client *Client) GetSmsSignWithContext(ctx context.Context, request *GetSmsSignRequest, runtime *dara.RuntimeOptions) (_result *GetSmsSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the approval details of a template after you submit a template application. You can view the approval status of the template.
//
// Description:
//
// - 新接口和原接口变更的公告详情请参见[关于短信服务更新签名&模板接口的公告](https://help.aliyun.com/document_detail/2806975.html)。
//
// - 审核时间：一般情况下，模板提交后，阿里云预计在2个小时内审核完成（审核工作时间：周一至周日9:00~21:00，法定节假日顺延）。
//
// - 如果模板未通过审核，会返回审核失败的原因，请参考[短信审核失败的处理建议](https://help.aliyun.com/document_detail/65990.html)，调用[UpdateSmsTemplate](~~UpdateSmsTemplate~~)接口或在[模板管理](https://dysms.console.aliyun.com/domestic/text/template)页面修改短信模板。
//
// - 当前接口是通过模板Code查询单个模板的审核详情。[QuerySmsTemplateList](https://help.aliyun.com/document_detail/419288.html)接口可以查询您当前账号下所有模板的模板详情。
//
// @param request - GetSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsTemplateResponse
func (client *Client) GetSmsTemplateWithContext(ctx context.Context, request *GetSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetSmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模板列表详情（新接口）
//
// @param tmpReq - GetSmsTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsTemplateListResponse
func (client *Client) GetSmsTemplateListWithContext(ctx context.Context, tmpReq *GetSmsTemplateListRequest, runtime *dara.RuntimeOptions) (_result *GetSmsTemplateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetSmsTemplateListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UsableStateList) {
		request.UsableStateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UsableStateList, dara.String("UsableStateList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
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

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateTag) {
		query["TemplateTag"] = request.TemplateTag
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.UsableStateListShrink) {
		query["UsableStateList"] = request.UsableStateListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmsTemplateList"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmsTemplateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tags are markers that you add to templates. Each tag consists of a key-value pair (Key-Value). You can query the template codes bound to a tag key-value pair, or query all tags bound to a specific template.
//
// Description:
//
// You can log on to the Short Message Service (SMS) console and go to the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page to filter SMS templates that are bound to tag key-value pairs, or click **Details*	- in the Actions column to view the tags that are bound to the current template.
//
// ### QPS limit
//
// The per-user QPS limit of this operation is 50 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Call this operation properly.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation is discontinued.
//
// Description:
//
// - 根据工信部规定与运营商[相关要求](https://help.aliyun.com/document_detail/2806975.html)，阿里云进行了签名相关API的功能改造。为提升您签名的审核效率和审核通过率，请您使用新接口[UpdateSmsSign-修改短信签名](https://help.aliyun.com/document_detail/2807428.html)。
//
// - 仅支持修改未通过审核的签名，请参考[短信审核失败的处理建议](https://help.aliyun.com/document_detail/65990.html)，调用此接口修改后自动提交审核，签名进入待审核状态。
//
// - 通过接口修改签名的操作会在短信服务控制台同步，在控制台对签名的相关操作，请参见[短信签名](https://help.aliyun.com/document_detail/108073.html)。
//
// @param request - ModifySmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmsSignResponse
func (client *Client) ModifySmsSignWithContext(ctx context.Context, request *ModifySmsSignRequest, runtime *dara.RuntimeOptions) (_result *ModifySmsSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This operation is discontinued.
//
// Description:
//
// - 根据工信部规定与运营商[相关要求](https://help.aliyun.com/document_detail/2806975.html)，阿里云进行了模板相关API的功能改造。为提升您模板的审核效率和审核通过率，请您使用新接口[UpdateSmsTemplate-修改短信模板](https://help.aliyun.com/document_detail/2807432.html)。
//
// - 仅支持修改未通过审核的模板，请参考[短信审核失败的处理建议](https://help.aliyun.com/document_detail/65990.html)，调用此接口修改后重新提交审核。
//
// - 通过接口修改模板的操作会在短信服务控制台同步，在控制台对模板的相关操作，请参见[短信模板](https://help.aliyun.com/document_detail/108085.html)。
//
// ### QPS限制
//
// 本接口的单用户QPS限制为1000次/秒。超过限制，API调用会被限流，这可能会影响您的业务，请合理调用。
//
// @param request - ModifySmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmsTemplateResponse
func (client *Client) ModifySmsTemplateWithContext(ctx context.Context, request *ModifySmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifySmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the review status of a card SMS template and returns information about the review by mobile phone vendors.
//
// Description:
//
// - Alibaba Cloud accounts that have not activated the card SMS service cannot call this API.
//
// - The card SMS service is currently in the internal invitation phase. Contact your Alibaba Cloud business manager to apply for activation or [contact Alibaba Cloud pre-sales consultation](https://help.aliyun.com/document_detail/464625.html).
//
// - You can also log on to the [Domestic Card SMS](https://dysms.console.aliyun.com/domestic/card) page in the console and query the review status of card SMS templates on the Template Management tab.
//
// ### QPS limit
//
// The per-user QPS limit for this operation is 300 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation responsibly.
//
// @param request - QueryCardSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCardSmsTemplateResponse
func (client *Client) QueryCardSmsTemplateWithContext(ctx context.Context, request *QueryCardSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryCardSmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parsing data of a specified card SMS template. The parsing data includes the number of successful SMS parsing receipts, the number of message exposures, and the number of message clicks.
//
// Description:
//
// ### QPS limit
//
// The QPS limit on this operation is 300 calls per second per user. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. Call this operation at a reasonable pace.
//
// @param request - QueryCardSmsTemplateReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCardSmsTemplateReportResponse
func (client *Client) QueryCardSmsTemplateReportWithContext(ctx context.Context, request *QueryCardSmsTemplateReportRequest, runtime *dara.RuntimeOptions) (_result *QueryCardSmsTemplateReportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a digital SMS signature by its name.
//
// Description:
//
// You can query only the digital SMS signatures that belong to your Alibaba Cloud account.
//
// @param request - QueryDigitalSignByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDigitalSignByNameResponse
func (client *Client) QueryDigitalSignByNameWithContext(ctx context.Context, request *QueryDigitalSignByNameRequest, runtime *dara.RuntimeOptions) (_result *QueryDigitalSignByNameResponse, _err error) {
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
		Action:      dara.String("QueryDigitalSignByName"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDigitalSignByNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryExtCodeSignResponse
func (client *Client) QueryExtCodeSignWithContext(ctx context.Context, request *QueryExtCodeSignRequest, runtime *dara.RuntimeOptions) (_result *QueryExtCodeSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a phone number supports card SMS messages.
//
// Description:
//
// - 未开通卡片短信业务的阿里云账号无法调用此API。
//
// - 目前卡片短信在内部邀约阶段，请联系您的阿里云商务经理申请开通或[联系阿里云售前咨询](https://help.aliyun.com/document_detail/464625.html)。
//
// @param tmpReq - QueryMobilesCardSupportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMobilesCardSupportResponse
func (client *Client) QueryMobilesCardSupportWithContext(ctx context.Context, tmpReq *QueryMobilesCardSupportRequest, runtime *dara.RuntimeOptions) (_result *QueryMobilesCardSupportResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPageSmartShortUrlLogResponse
func (client *Client) QueryPageSmartShortUrlLogWithContext(ctx context.Context, request *QueryPageSmartShortUrlLogRequest, runtime *dara.RuntimeOptions) (_result *QueryPageSmartShortUrlLogResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询终端5G适配情况
//
// @param request - QueryRCSMobileCapableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRCSMobileCapableResponse
func (client *Client) QueryRCSMobileCapableWithContext(ctx context.Context, request *QueryRCSMobileCapableRequest, runtime *dara.RuntimeOptions) (_result *QueryRCSMobileCapableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PhoneNumbers) {
		query["PhoneNumbers"] = request.PhoneNumbers
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRCSMobileCapable"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRCSMobileCapableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询终端5G适配情况任务结果
//
// @param request - QueryRCSMobileCapableTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRCSMobileCapableTaskResultResponse
func (client *Client) QueryRCSMobileCapableTaskResultWithContext(ctx context.Context, request *QueryRCSMobileCapableTaskResultRequest, runtime *dara.RuntimeOptions) (_result *QueryRCSMobileCapableTaskResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRCSMobileCapableTaskResult"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRCSMobileCapableTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询5G模板详情
//
// @param request - QueryRCSTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRCSTemplateResponse
func (client *Client) QueryRCSTemplateWithContext(ctx context.Context, request *QueryRCSTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryRCSTemplateResponse, _err error) {
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
		Action:      dara.String("QueryRCSTemplate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRCSTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定版本查看5G消息固定菜单详情
//
// @param request - QueryRcsSignMenuByVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRcsSignMenuByVersionResponse
func (client *Client) QueryRcsSignMenuByVersionWithContext(ctx context.Context, request *QueryRcsSignMenuByVersionRequest, runtime *dara.RuntimeOptions) (_result *QueryRcsSignMenuByVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RcsMenuVersion) {
		query["RcsMenuVersion"] = request.RcsMenuVersion
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRcsSignMenuByVersion"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRcsSignMenuByVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the delivery records and status of SMS messages sent to a single phone number.
//
// Description:
//
// - This operation queries the details of SMS messages sent to a specific phone number on a given date. You can also provide the `BizId` to query a specific delivery record.
//
// - This operation queries delivery details for a single phone number at a time. To view delivery details in bulk, use the [QuerySendStatistics](https://help.aliyun.com/document_detail/419276.html) operation to query delivery statistics, or log in to the [Delivery Receipts](https://dysms.console.aliyun.com/record) page in the console.
//
// ### QPS limit
//
// The QPS limit for this operation is 5,000 requests per second per user. Calls that exceed this limit are throttled.
//
// @param request - QuerySendDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendDetailsResponse
func (client *Client) QuerySendDetailsWithContext(ctx context.Context, request *QuerySendDetailsRequest, runtime *dara.RuntimeOptions) (_result *QuerySendDetailsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries message delivery statistics, including send dates, the number of successfully sent messages, and the number of received delivery receipts.
//
// Description:
//
// - If you query data over a long time range, the results are paginated. You can specify the page size and page number to navigate through the Delivery Logs.
//
// - You can also view delivery details by logging in to the [Short Message Service console](https://dysms.console.aliyun.com/dysms.htm#/overview) and navigating to the **Business Statistics*	- > **Delivery Logs*	- page.
//
// @param request - QuerySendStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendStatisticsResponse
func (client *Client) QuerySendStatisticsWithContext(ctx context.Context, request *QuerySendStatisticsRequest, runtime *dara.RuntimeOptions) (_result *QuerySendStatisticsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the status and availability of a short link.
//
// Description:
//
//	Notice:
//
// This API is not currently supported by Short Message Service.
//
// ### QPS limit
//
// The QPS limit for this API is 20 queries per second per user. API calls that exceed this limit are rate-limited, which can impact your business. Plan your calls accordingly.
//
// @param request - QueryShortUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryShortUrlResponse
func (client *Client) QueryShortUrlWithContext(ctx context.Context, request *QueryShortUrlRequest, runtime *dara.RuntimeOptions) (_result *QueryShortUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After you apply for SMS qualifications, you can call this operation to query the details of a single qualification.
//
// Description:
//
// - This API queries the details of a single qualification (enterprise information, legal representative information, and administrator information).
//
// - To query all qualification information under your current account, or to query review details, call the [QuerySmsQualificationRecord](~~QuerySmsQualificationRecord~~) operation.
//
// - Affected by the SMS signature real-name registration requirements, the volume of qualification review tickets is growing rapidly, and the review time may be extended. Please be patient. The review is expected to be completed within 2 business days (review hours: Monday to Sunday 9:00-21:00, postponed for legal holidays). In special cases, the review time may be extended. Please be patient.
//
// @param request - QuerySingleSmsQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySingleSmsQualificationResponse
func (client *Client) QuerySingleSmsQualificationWithContext(ctx context.Context, request *QuerySingleSmsQualificationRequest, runtime *dara.RuntimeOptions) (_result *QuerySingleSmsQualificationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries icp record details.
//
// Description:
//
// Pass a list of icp record IDs to retrieve their details.
//
// For example, call the QuerySmsSignList or GetSmsSign API to obtain the required icp record IDs.
//
// @param tmpReq - QuerySmsAppIcpRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsAppIcpRecordResponse
func (client *Client) QuerySmsAppIcpRecordWithContext(ctx context.Context, tmpReq *QuerySmsAppIcpRecordRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsAppIcpRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QuerySmsAppIcpRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppIcpRecordIdList) {
		request.AppIcpRecordIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppIcpRecordIdList, dara.String("AppIcpRecordIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIcpRecordIdListShrink) {
		query["AppIcpRecordIdList"] = request.AppIcpRecordIdListShrink
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
		Action:      dara.String("QuerySmsAppIcpRecord"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmsAppIcpRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries created letters of authorization. You can view the review status and authorized signature scope of the letters.
//
// Description:
//
// - Supports full query or conditional query:
//
//   - **Full query**: Queries the information of all letters of authorization under your current account. No parameters need to be passed. Full query is performed by default.
//
//   - **Conditional query**: Supports queries by letter of authorization ID, signature name, and review status of the letter of authorization. Pass the parameters by which you want to filter.
//
// - Review duration: Affected by the real-name registration requirements for SMS signatures, the volume of qualification review tickets is currently increasing rapidly, and the review duration may be extended. Please wait patiently. The review is expected to be completed within 2 working days. SMS signatures and templates are expected to be reviewed within 2 hours after submission. Reviews involving governments and enterprises are generally completed within 2 working days. If verification upgrades, a large number of review tasks, or non-working hours are encountered, the review duration may be extended. Please wait patiently. (Review working hours: Monday to Sunday, 9:00–21:00, postponed for statutory holidays.)
//
// @param tmpReq - QuerySmsAuthorizationLetterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsAuthorizationLetterResponse
func (client *Client) QuerySmsAuthorizationLetterWithContext(ctx context.Context, tmpReq *QuerySmsAuthorizationLetterRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsAuthorizationLetterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of SMS qualifications and their review details after you submit qualification applications. This operation supports filtered query.
//
// Description:
//
// - 支持全量查询或条件查询：
//
//   - **全量查询**：查询您当前帐户下所有短信资质，无需传参。默认全量查询。
//
//   - **条件查询**：支持根据资质名称、企业名称、法人姓名、审核状态、审核工单ID、资质用途进行查询，传入您希望筛选的参数即可。
//
// - 本接口用于查询资质及其审核信息，如果需要查询单个资质的具体信息（企业信息、法人信息、管理员信息），请调用[查询单个资质详情](~~QuerySingleSmsQualification~~)接口。
//
// - 受短信签名实名制报备要求影响，当前资质审核工单量增长快速，审核时间可能会延长，请耐心等待，预计2个工作日内完成（审核工作时间：周一至周日 9:00~21:00，法定节假日顺延）。特殊情况可能延长审核时间，请耐心等待。
//
// - 如果资质未通过审核，审核备注`AuditRemark`会返回审核失败的原因，请参考[审核失败的处理建议](~~2384377#a96cc318b94x1~~)，调用[修改短信资质](~~UpdateSmsQualification~~)接口或在控制台[资质管理](https://dysms.console.aliyun.com/domestic/text/qualification)页面修改资质信息后，重新发起审核。
//
// @param request - QuerySmsQualificationRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsQualificationRecordResponse
func (client *Client) QuerySmsQualificationRecordWithContext(ctx context.Context, request *QuerySmsQualificationRecordRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsQualificationRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the review status of an SMS signature.
//
// Description:
//
// - To comply with regulations from the Ministry of Industry and Information Technology (MIIT) and [related requirements](https://help.aliyun.com/document_detail/2806975.html) from carriers, Alibaba Cloud has upgraded its SMS signature management APIs. We recommend using the new [GetSmsSign - Query Signature Details](https://help.aliyun.com/document_detail/2807429.html) API, which returns more detailed information about signatures than this API.
//
// - We typically review signatures within two hours of submission. The review service is available from 9:00 to 21:00, Monday to Sunday. Reviews may be delayed during public holidays. We recommend submitting your application before 18:00 for a timely review.
//
// - If a signature is rejected, the response includes the review reason. For troubleshooting information, see [Troubleshooting Signature Review Failures](https://help.aliyun.com/document_detail/65990.html). You can then call the [ModifySmsTemplate](https://help.aliyun.com/document_detail/419287.html) API or modify the SMS signature on the [Signature Management](https://dysms.console.aliyun.com/domestic/text) page.
//
// - This API queries the review details for a single signature by name. To query all signatures in your account, call the [QuerySmsSignList](https://help.aliyun.com/document_detail/419288.html) API.
//
// @param request - QuerySmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsSignResponse
func (client *Client) QuerySmsSignWithContext(ctx context.Context, request *QuerySmsSignRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query all signatures under your account, including signature audit status, signature type, and signature name.
//
// Description:
//
// This operation queries the signature information that was **first created*	- or the **most recently approved*	- signature details under your current account. If you need to query more information such as application scenario content or files uploaded during application, you can call the [GetSmsSign](~~GetSmsSign~~) operation to query the audit details of a single signature by signature name.
//
// @param request - QuerySmsSignListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsSignListResponse
func (client *Client) QuerySmsSignListWithContext(ctx context.Context, request *QuerySmsSignListRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsSignListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This API has been deprecated.
//
// Description:
//
// - Alibaba Cloud has updated its template-related APIs to comply with regulatory and [carrier requirements](https://help.aliyun.com/document_detail/2806975.html). We recommend that you use the new [GetSmsTemplate - Query template review details](https://help.aliyun.com/document_detail/2807433.html) API. The new API returns more detailed template information in its response.
//
// - Review timeline: After you submit a template, Alibaba Cloud typically completes the review within two hours. Review hours are 9:00 to 21:00 (UTC+8) from Monday to Sunday. Reviews are postponed during public holidays. We recommend that you submit your templates before 18:00 (UTC+8).
//
// - If a template fails review, the response includes the reason for the rejection. For more information, see [Suggestions for handling a failed review](https://help.aliyun.com/document_detail/65990.html). You can then call the [ModifySmsTemplate](https://help.aliyun.com/document_detail/419287.html) API or modify the template on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
//
// - QuerySmsTemplate queries the review details of a single template by its template code. To query the details of all templates in your account, call the [QuerySmsTemplateList](https://help.aliyun.com/document_detail/419288.html) API.
//
// @param request - QuerySmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsTemplateResponse
func (client *Client) QuerySmsTemplateWithContext(ctx context.Context, request *QuerySmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query all templates under your account. This way, you can view template details, including the template approval status, template type, and template content.
//
// Description:
//
// - This operation queries the template details of all templates under your current account. To query more details such as the template variable content and the file information uploaded during application, you can call the [GetSmsTemplate](~~GetSmsTemplate~~) operation to query the approval details of a single template by template code.
//
// - You can also log on to the Short Message Service (SMS) console and view the template details of all templates under your current account on the [Template Management](https://dysms.console.aliyun.com/domestic/text/template) page.
//
// @param request - QuerySmsTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsTemplateListResponse
func (client *Client) QuerySmsTemplateListWithContext(ctx context.Context, request *QuerySmsTemplateListRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsTemplateListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more trademarks.
//
// Description:
//
// This operation retrieves the details of trademarks by using a list of trademark IDs.
//
// For example, you can obtain trademark IDs by calling signature query operations such as `QuerySmsSignList` or `GetSmsSign`. You can then use this operation to retrieve the details of each trademark.
//
// @param tmpReq - QuerySmsTrademarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmsTrademarkResponse
func (client *Client) QuerySmsTrademarkWithContext(ctx context.Context, tmpReq *QuerySmsTrademarkRequest, runtime *dara.RuntimeOptions) (_result *QuerySmsTrademarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QuerySmsTrademarkShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TrademarkIdList) {
		request.TrademarkIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TrademarkIdList, dara.String("TrademarkIdList"), dara.String("json"))
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

	if !dara.IsNil(request.TrademarkIdListShrink) {
		query["TrademarkIdList"] = request.TrademarkIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmsTrademark"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmsTrademarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When applying for SMS qualification, the administrator\\"s phone number must be verified. Use this operation to obtain an SMS verification code.
//
// Description:
//
// - After you receive the phone verification code, pass it to the `CertifyCode` parameter of the [SubmitSmsQualification](~~SubmitSmsQualification~~) or [UpdateSmsQualification](~~UpdateSmsQualification~~) operation.
//
// - You can call the [ValidPhoneCode](~~ValidPhoneCode~~) operation to verify whether the SMS verification code is correct.
//
// - This operation is subject to [throttling](~~44335#section-0wh-xn6-0t7~~). Do not call it too frequently. For the same phone number, a maximum of 1 message per minute, 5 messages per hour, and 10 messages per day are supported.
//
// @param request - RequiredPhoneCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RequiredPhoneCodeResponse
func (client *Client) RequiredPhoneCodeWithContext(ctx context.Context, request *RequiredPhoneCodeRequest, runtime *dara.RuntimeOptions) (_result *RequiredPhoneCodeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends card SMS messages in batches.
//
// Description:
//
// - Sending card SMS messages is a billable operation. You are not charged if a card SMS message fails to be sent or rendered. For more information, see [Multimedia SMS pricing](https://help.aliyun.com/document_detail/437951.html).
//
// - The card SMS feature is currently in the internal invitation phase. Contact your Alibaba Cloud business manager to apply for activation, or contact [Alibaba Cloud pre-sales consulting](https://help.aliyun.com/document_detail/464625.html?spm=a2c4g.11186623.0.0.213219fcSn2Ykj#section-81n-72q-ybm).
//
// - We recommend that you set the timeout period for card SMS messages to a value greater than or equal to 3 seconds. If a timeout failure occurs, we recommend that you check the delivery status before deciding whether to retry. We also recommend that you do not enable SDK retry logic when calling this operation; otherwise, multiple sending attempts may occur. For more information about timeout and retry settings, see [Timeout mechanism](https://help.aliyun.com/document_detail/262079.html) and [Retry mechanism](https://help.aliyun.com/document_detail/262080.html).
//
// - Domestic SMS, international SMS, and multimedia SMS do not currently support idempotency. Implement idempotency control to prevent duplicate operations caused by multiple retries.
//
// - Before you send a card SMS message, you must add and obtain approval for a card SMS template. When this operation is called to send an SMS message, the system checks whether the phone number supports card SMS messages. If the phone number does not support card SMS messages, you can configure whether to accept fallback to digital SMS or text SMS in the operation to improve the delivery rate.
//
// - When you send card SMS messages in batches, each phone number can use a different signature and a different fallback. In a single request, you can send card SMS messages to a maximum of 100 phone numbers.
//
// ### QPS limit
//
// The QPS limit per user for this operation is 1,000 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation in a reasonable manner.
//
// @param request - SendBatchCardSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendBatchCardSmsResponse
func (client *Client) SendBatchCardSmsWithContext(ctx context.Context, request *SendBatchCardSmsRequest, runtime *dara.RuntimeOptions) (_result *SendBatchCardSmsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation sends messages to different phone numbers using a single template, with different signatures and template variables for each recipient.
//
// Description:
//
// ### Basic information
//
// - You can send messages to a maximum of 100 phone numbers in a single call.
//
// - The global [endpoint](https://help.aliyun.com/document_detail/419270.html) is `dysmsapi.aliyuncs.com`. For a list of region-specific endpoints, see [Endpoints](https://help.aliyun.com/document_detail/419270.html).
//
// ### API calls
//
// - We recommend calling this operation using an SDK. For more information, see [Make your first API call](https://help.aliyun.com/document_detail/2841024.html).
//
// - To send messages from the console, see [Send group messages](https://help.aliyun.com/document_detail/108266.html).
//
// - To build your own API requests, see [V3 request body and signature mechanism](https://help.aliyun.com/document_detail/2593177.html).
//
// ### Usage notes
//
// - For domestic SMS, we recommend setting the timeout period to 1 second or longer. If a timeout occurs, check the delivery receipt status before you retry the request. For more information about timeout and retry settings, see [timeout mechanism](https://help.aliyun.com/document_detail/262079.html) and [retry mechanism](https://help.aliyun.com/document_detail/262080.html).
//
// - This operation does not support idempotence for domestic SMS, international SMS, or Multimedia Messaging Service (MMS) messages. You must implement your own idempotence controls to prevent duplicate operations caused by multiple retries.
//
// - This is a billable operation. For domestic SMS, you are charged based on the delivery receipt status from the carrier. You are not charged for messages that are successfully submitted but fail carrier delivery. For more information, see [Billing overview](https://help.aliyun.com/document_detail/44340.html).
//
//		Warning:
//
//	Batch message delivery may be delayed due to system capacity limits. For time-sensitive messages, such as verification codes or alert notifications, use the SendSms operation to send messages individually.
//
// ### QPS limit
//
// The Queries Per Second (QPS) limit for a single user is 5,000. Calls that exceed this limit are throttled. Plan your usage accordingly.
//
// @param request - SendBatchSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendBatchSmsResponse
func (client *Client) SendBatchSmsWithContext(ctx context.Context, request *SendBatchSmsRequest, runtime *dara.RuntimeOptions) (_result *SendBatchSmsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// - 发送卡片短信为计费接口，卡片短信发送失败或渲染失败时不计费，详情请参见[多媒体短信定价](https://help.aliyun.com/document_detail/437951.html)。
//
// - 目前卡片短信在内部邀约阶段，请联系您的阿里云商务经理申请开通或联系[阿里云售前咨询](https://help.aliyun.com/document_detail/464625.html?spm=a2c4g.11186623.0.0.213219fcSn2Ykj#section-81n-72q-ybm)。
//
// - 卡片短信超时时间建议设置为≥3S；发生超时失败的情况时，建议查看回执状态后再判断是否重试。同时建议您在调用此接口时，不要开启SDK重试逻辑，否则可能会造成多次发送的情况。超时和重试的相关设置，请参见[超时机制](https://help.aliyun.com/document_detail/262079.html)、[重试机制](https://help.aliyun.com/document_detail/262080.html)。
//
// - 国内短信、国际短信及多媒体短信目前均不支持幂等的能力，请您做好幂等控制，防止因多次重试而导致的重复操作问题。
//
// - 发送卡片短信前需添加卡片短信模板且模板审核通过。本接口在发送短信时，会校验号码是否支持卡片短信。如果该手机号不支持发送卡片短信，可在接口中设置是否接受数字短信和文本短信的回落，提升发送的触达率。
//
// ### QPS限制
//
// 本接口的单用户QPS限制为1000次/秒。超过限制，API调用会被限流，这可能会影响您的业务，请合理调用。
//
// @param request - SendCardSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendCardSmsResponse
func (client *Client) SendCardSmsWithContext(ctx context.Context, request *SendCardSmsRequest, runtime *dara.RuntimeOptions) (_result *SendCardSmsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLogisticsSmsResponse
func (client *Client) SendLogisticsSmsWithContext(ctx context.Context, request *SendLogisticsSmsRequest, runtime *dara.RuntimeOptions) (_result *SendLogisticsSmsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 5G消息首次下行
//
// @param request - SendRCSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRCSResponse
func (client *Client) SendRCSWithContext(ctx context.Context, request *SendRCSRequest, runtime *dara.RuntimeOptions) (_result *SendRCSResponse, _err error) {
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

	if !dara.IsNil(request.PhoneNumbers) {
		query["PhoneNumbers"] = request.PhoneNumbers
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
		Action:      dara.String("SendRCS"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendRCSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 5G消息交互下行
//
// @param request - SendRCSReplyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRCSReplyResponse
func (client *Client) SendRCSReplyWithContext(ctx context.Context, request *SendRCSReplyRequest, runtime *dara.RuntimeOptions) (_result *SendRCSReplyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InReplyToRcsID) {
		query["InReplyToRcsID"] = request.InReplyToRcsID
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.PhoneNumbers) {
		query["PhoneNumbers"] = request.PhoneNumbers
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
		Action:      dara.String("SendRCSReply"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendRCSReplyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends an SMS message to one or more specified mobile numbers.
//
// Description:
//
// Use this API to send an SMS message to a single mobile number. This API also supports sending messages with the same signature and template variables to multiple mobile numbers, up to 1,000 per request. Note that bulk sending may experience some latency. If you need to send messages with different signatures or template variables to multiple recipients, use the [SendBatchSms](https://help.aliyun.com/document_detail/419274.html) API, which supports up to 100 mobile numbers per request.
//
// ### Usage notes
//
// - For SMS messages sent to the Chinese mainland, we recommend setting the timeout period to 1 second or longer. If a timeout occurs, check the delivery receipt status before retrying the request. For more information about timeout and retry settings, see [Timeout mechanism](https://help.aliyun.com/document_detail/262079.html) and [Retry mechanism](https://help.aliyun.com/document_detail/262080.html).
//
// - This API does not support idempotence for domestic, international, or multimedia SMS messages. You must implement your own idempotence controls to prevent sending duplicate messages during retries.
//
// - This is a billable API. For messages sent to the Chinese mainland, billing is based on the delivery receipt status from the carrier. You are not charged if the API call is successful but the carrier fails to deliver the message. For more information, see [Billing](https://help.aliyun.com/document_detail/44340.html).
//
// ### QPS limit
//
// This API has a queries per second (QPS) limit of 5,000 for each user. The system throttles calls that exceed this limit. To avoid throttling, use this API within the specified limit.
//
// @param request - SendSmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSmsResponse
func (client *Client) SendSmsWithContext(ctx context.Context, request *SendSmsRequest, runtime *dara.RuntimeOptions) (_result *SendSmsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Feeds back the SMS delivery status corresponding to each message ID (MessageId) to the Alibaba Cloud International SMS platform.
//
// Description:
//
// Metric definitions:
//
// - OTP send volume: the number of verification codes sent.
//
// - OTP conversion volume: the number of verification codes converted (the number of times a user successfully obtained a verification code and reported it back).
//
// Conversion rate = OTP conversion volume / OTP send volume.
//
// > The conversion rate feedback feature has a certain level of intrusiveness on the business system. To prevent jitter in conversion rate API calls from affecting business logic, please consider the following:  - Call the API in asynchronous mode (for example, using a queue or event-driven approach).  - Add a degradable solution to protect business logic (for example, manual degradation, or automatic degradation using a circuit breaker).
//
// @param request - SmsConversionIntlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmsConversionIntlResponse
func (client *Client) SmsConversionIntlWithContext(ctx context.Context, request *SmsConversionIntlRequest, runtime *dara.RuntimeOptions) (_result *SmsConversionIntlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an SMS qualification application. As required by the Ministry of Industry and Information Technology (MIIT) and carriers for real-name SMS sending, domestic SMS services require qualification credential information of the signature owner. Apply for an SMS qualification first, and then apply for signatures and templates.
//
// Description:
//
// - Before submitting an application, read [Qualification material description](https://help.aliyun.com/document_detail/2384377.html) and prepare the required qualification materials.
//
// - Currently, only users who have completed **verify your identity - Enterprise account*	- can use the API to apply for SMS qualifications. If your Alibaba Cloud account has completed verify your identity - Individual account, apply for qualifications through the Short Message Service [console](https://dysms.console.aliyun.com/domestic/text/qualification/add), or [upgrade to verify your identity - Enterprise account](https://help.aliyun.com/document_detail/37178.html). [View my account verification type](https://myaccount.console.aliyun.com/cert-info)
//
// - Batch qualification applications are not supported. Wait at least 5 seconds between applications.
//
// @param tmpReq - SubmitSmsQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSmsQualificationResponse
func (client *Client) SubmitSmsQualificationWithContext(ctx context.Context, tmpReq *SubmitSmsQualificationRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmsQualificationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tags can mark resources, allowing enterprises or individuals to classify templates of the same type for easier search and resource aggregation. Call this operation to bind tags to SMS templates.
//
// Description:
//
// - Each template can be bound to up to 20 tags.
//
// - The tag key (Key) must be unique within the same template. If a template has two tags with the same Key but different Values, the new value overwrites the old value.
//
// - This feature is only available for domestic text messages of Short Message Service on the China site.
//
// ### QPS limit
//
// The per-user QPS limit of this operation is 50 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tags can mark resources, allowing enterprises or individuals to categorize templates of the same type for easier search and resource aggregation. If a template is no longer applicable to its currently bound tags, you can unbind the tags from the template. You can delete a single tag or delete tags in batches.
//
// Description:
//
// ### QPS limit
//
// The QPS limit per user for this operation is 50 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call the operation reasonably.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExtCodeSignResponse
func (client *Client) UpdateExtCodeSignWithContext(ctx context.Context, request *UpdateExtCodeSignRequest, runtime *dara.RuntimeOptions) (_result *UpdateExtCodeSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑5g签名
//
// @param request - UpdateRCSSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRCSSignatureResponse
func (client *Client) UpdateRCSSignatureWithContext(ctx context.Context, request *UpdateRCSSignatureRequest, runtime *dara.RuntimeOptions) (_result *UpdateRCSSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackgroundImage) {
		query["BackgroundImage"] = request.BackgroundImage
	}

	if !dara.IsNil(request.BubbleColor) {
		query["BubbleColor"] = request.BubbleColor
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Latitude) {
		query["Latitude"] = request.Latitude
	}

	if !dara.IsNil(request.Logo) {
		query["Logo"] = request.Logo
	}

	if !dara.IsNil(request.Longitude) {
		query["Longitude"] = request.Longitude
	}

	if !dara.IsNil(request.OfficeAddress) {
		query["OfficeAddress"] = request.OfficeAddress
	}

	if !dara.IsNil(request.ServiceEmail) {
		query["ServiceEmail"] = request.ServiceEmail
	}

	if !dara.IsNil(request.ServicePhone) {
		query["ServicePhone"] = request.ServicePhone
	}

	if !dara.IsNil(request.ServiceTerms) {
		query["ServiceTerms"] = request.ServiceTerms
	}

	if !dara.IsNil(request.ServiceWebsite) {
		query["ServiceWebsite"] = request.ServiceWebsite
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRCSSignature"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRCSSignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you need to update SMS qualification information, you can submit a modification request through this API. After submission, it will re-enter the review process.
//
// Description:
//
// - Qualifications under review do not support modification. Please wait for the review process to finish, or [withdraw the application](https://dysms.console.aliyun.com/domestic/text/qualification) in the SMS Service console before making modifications.
//
// - The modified SMS qualification **must be re-reviewed*	- (including qualifications that have already passed review). Please upload materials that meet the specifications according to the [Qualification Material Description](https://help.aliyun.com/document_detail/2384377.html).
//
// - **Modification is not supported*	- for the qualification name, application purpose, or unified social credit code.
//
// - Batch modification of SMS qualifications is not supported. It is recommended to leave at least 5 seconds between modifications.
//
// @param tmpReq - UpdateSmsQualificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsQualificationResponse
func (client *Client) UpdateSmsQualificationWithContext(ctx context.Context, tmpReq *UpdateSmsQualificationRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmsQualificationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can modify rejected or approved signatures. A modified signature is automatically submitted for review, and its status changes to pending review.
//
// Description:
//
// - For details about the updates to the signature and template APIs, see [Announcement on Updating Signature & Template APIs for Short Message Service](https://help.aliyun.com/document_detail/2806975.html).
//
// - You can modify signatures that are either **rejected*	- or **approved**. For guidance on handling review failures, see [Handling signature review failures](https://help.aliyun.com/document_detail/65990.html). Call this API to modify and resubmit the signature for review.
//
// - You cannot use this API to edit the name of a **rejected*	- signature. To edit the name, go to the [Short Message Service console](https://dysms.console.aliyun.com/domestic/text/sign).
//
// - Signatures you request using this API are synchronized with the Short Message Service console. For information on managing signatures in the console, see [Signatures](https://help.aliyun.com/document_detail/108073.html).
//
// @param tmpReq - UpdateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsSignResponse
func (client *Client) UpdateSmsSignWithContext(ctx context.Context, tmpReq *UpdateSmsSignRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmsSignResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API modifies a template that failed review and automatically resubmits it.
//
// Description:
//
// - For details about the changes to the signature and template APIs, see [Announcement on Updating Signature & Template APIs for Short Message Service](https://help.aliyun.com/document_detail/2806975.html).
//
// - You can only modify templates that have failed review. For troubleshooting, see [Suggestions for handling failed SMS template reviews](https://help.aliyun.com/document_detail/65990.html). After modifying a template with this API, you must resubmit it for review.
//
// - Template changes made using this API are synchronized with the Short Message Service console. To learn more about managing templates in the console, see [SMS templates](https://help.aliyun.com/document_detail/108085.html).
//
// ### QPS limit
//
// The QPS limit for this API is 1,000 queries per second per user. If you exceed this limit, your API calls will be throttled. This can affect your business, so please use the API responsibly.
//
// @param tmpReq - UpdateSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsTemplateResponse
func (client *Client) UpdateSmsTemplateWithContext(ctx context.Context, tmpReq *UpdateSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmsTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级文本短信为5g签名
//
// @param request - UpgradeToRCSSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeToRCSSignatureResponse
func (client *Client) UpgradeToRCSSignatureWithContext(ctx context.Context, request *UpgradeToRCSSignatureRequest, runtime *dara.RuntimeOptions) (_result *UpgradeToRCSSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackgroundImage) {
		query["BackgroundImage"] = request.BackgroundImage
	}

	if !dara.IsNil(request.BubbleColor) {
		query["BubbleColor"] = request.BubbleColor
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Latitude) {
		query["Latitude"] = request.Latitude
	}

	if !dara.IsNil(request.Logo) {
		query["Logo"] = request.Logo
	}

	if !dara.IsNil(request.Longitude) {
		query["Longitude"] = request.Longitude
	}

	if !dara.IsNil(request.OfficeAddress) {
		query["OfficeAddress"] = request.OfficeAddress
	}

	if !dara.IsNil(request.ServiceEmail) {
		query["ServiceEmail"] = request.ServiceEmail
	}

	if !dara.IsNil(request.ServicePhone) {
		query["ServicePhone"] = request.ServicePhone
	}

	if !dara.IsNil(request.ServiceTerms) {
		query["ServiceTerms"] = request.ServiceTerms
	}

	if !dara.IsNil(request.ServiceWebsite) {
		query["ServiceWebsite"] = request.ServiceWebsite
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeToRCSSignature"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeToRCSSignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When applying for SMS qualification, the administrator\\"s phone number must be verified. This operation verifies the phone number and the received verification code.
//
// Description:
//
// - Call the [RequiredPhoneCode](~~RequiredPhoneCode~~) operation first. Alibaba Cloud sends an SMS verification code to the phone number that you provided.
//
// - This operation does not affect the SMS qualification application process and is used only to verify the SMS verification code. When you submit the actual application, pass the verification code into the `CertifyCode` parameter of the [SubmitSmsQualification](~~SubmitSmsQualification~~) operation.
//
// @param request - ValidPhoneCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidPhoneCodeResponse
func (client *Client) ValidPhoneCodeWithContext(ctx context.Context, request *ValidPhoneCodeRequest, runtime *dara.RuntimeOptions) (_result *ValidPhoneCodeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyLogisticsSmsMailNoResponse
func (client *Client) VerifyLogisticsSmsMailNoWithContext(ctx context.Context, request *VerifyLogisticsSmsMailNoRequest, runtime *dara.RuntimeOptions) (_result *VerifyLogisticsSmsMailNoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
