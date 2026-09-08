// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a webhook contact.
//
// @param request - CreateWebhookContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWebhookContactResponse
func (client *Client) CreateWebhookContactWithContext(ctx context.Context, request *CreateWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *CreateWebhookContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.BotSecurityToken) {
		body["BotSecurityToken"] = request.BotSecurityToken
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SecurityToken) {
		body["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServerUrl) {
		body["ServerUrl"] = request.ServerUrl
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	if !dara.IsNil(request.VerificationCode) {
		body["VerificationCode"] = request.VerificationCode
	}

	if !dara.IsNil(request.WebhookType) {
		body["WebhookType"] = request.WebhookType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWebhookContact"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWebhookContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a message.
//
// @param request - DelMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelMessageResponse
func (client *Client) DelMessageWithContext(ctx context.Context, request *DelMessageRequest, runtime *dara.RuntimeOptions) (_result *DelMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes all in-app messages (logical deletion).
//
// @param request - DeleteAllMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAllMessageResponse
func (client *Client) DeleteAllMessageWithContext(ctx context.Context, request *DeleteAllMessageRequest, runtime *dara.RuntimeOptions) (_result *DeleteAllMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAllMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAllMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a webhook contact.
//
// @param request - DeleteWebhookContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebhookContactResponse
func (client *Client) DeleteWebhookContactWithContext(ctx context.Context, request *DeleteWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebhookContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebhookContact"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebhookContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all common contacts across accounts.
//
// @param request - ReadAllCommonContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadAllCommonContactsResponse
func (client *Client) ReadAllCommonContactsWithContext(ctx context.Context, request *ReadAllCommonContactsRequest, runtime *dara.RuntimeOptions) (_result *ReadAllCommonContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadAllCommonContacts"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadAllCommonContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all marketing preference configurations.
//
// @param request - ReadAllMarketingPreferencesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadAllMarketingPreferencesResponse
func (client *Client) ReadAllMarketingPreferencesWithContext(ctx context.Context, request *ReadAllMarketingPreferencesRequest, runtime *dara.RuntimeOptions) (_result *ReadAllMarketingPreferencesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadAllMarketingPreferences"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadAllMarketingPreferencesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Marks all messages as read. If no category is specified, all messages are marked as read.
//
// @param request - ReadAllMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadAllMessageResponse
func (client *Client) ReadAllMessageWithContext(ctx context.Context, request *ReadAllMessageRequest, runtime *dara.RuntimeOptions) (_result *ReadAllMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadAllMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadAllMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the full list of webhooks.
//
// @param request - ReadAllWebhookContactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadAllWebhookContactsResponse
func (client *Client) ReadAllWebhookContactsWithContext(ctx context.Context, request *ReadAllWebhookContactsRequest, runtime *dara.RuntimeOptions) (_result *ReadAllWebhookContactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadAllWebhookContacts"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadAllWebhookContactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get category group list
//
// @param request - ReadCategoryGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadCategoryGroupListResponse
func (client *Client) ReadCategoryGroupListWithContext(ctx context.Context, request *ReadCategoryGroupListRequest, runtime *dara.RuntimeOptions) (_result *ReadCategoryGroupListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ChannelGroupCode) {
		body["ChannelGroupCode"] = request.ChannelGroupCode
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadCategoryGroupList"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadCategoryGroupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the number of read messages for each category.
//
// @param request - ReadClassNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadClassNameResponse
func (client *Client) ReadClassNameWithContext(ctx context.Context, request *ReadClassNameRequest, runtime *dara.RuntimeOptions) (_result *ReadClassNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadClassName"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadClassNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a general contact.
//
// @param request - ReadCommonContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadCommonContactResponse
func (client *Client) ReadCommonContactWithContext(ctx context.Context, request *ReadCommonContactRequest, runtime *dara.RuntimeOptions) (_result *ReadCommonContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadCommonContact"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadCommonContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the marketing preferences of a user.
//
// @param request - ReadMarketingPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMarketingPreferenceResponse
func (client *Client) ReadMarketingPreferenceWithContext(ctx context.Context, request *ReadMarketingPreferenceRequest, runtime *dara.RuntimeOptions) (_result *ReadMarketingPreferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMarketingPreference"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMarketingPreferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Marks a message as read.
//
// @param request - ReadMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageResponse
func (client *Client) ReadMessageWithContext(ctx context.Context, request *ReadMessageRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the body of a message.
//
// @param request - ReadMessageContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageContentResponse
func (client *Client) ReadMessageContentWithContext(ctx context.Context, request *ReadMessageContentRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.History) {
		body["History"] = request.History
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageContent"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the message language.
//
// @param request - ReadMessageLanguageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageLanguageResponse
func (client *Client) ReadMessageLanguageWithContext(ctx context.Context, request *ReadMessageLanguageRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageLanguageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.ReturnDefaultLang) {
		body["ReturnDefaultLang"] = request.ReturnDefaultLang
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageLanguage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageLanguageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of messages.
//
// @param request - ReadMessageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageListResponse
func (client *Client) ReadMessageListWithContext(ctx context.Context, request *ReadMessageListRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.History) {
		body["History"] = request.History
	}

	if !dara.IsNil(request.Loc) {
		body["Loc"] = request.Loc
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		body["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageList"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the total number of unread messages.
//
// @param request - ReadMessageNewTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageNewTotalResponse
func (client *Client) ReadMessageNewTotalWithContext(ctx context.Context, request *ReadMessageNewTotalRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageNewTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageNewTotal"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageNewTotalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the metadata service.
//
// @param request - ReadMetaConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMetaConfigResponse
func (client *Client) ReadMetaConfigWithContext(ctx context.Context, request *ReadMetaConfigRequest, runtime *dara.RuntimeOptions) (_result *ReadMetaConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMetaConfig"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMetaConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the count of read messages for each category.
//
// @param request - ReadNumGroupByClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadNumGroupByClassResponse
func (client *Client) ReadNumGroupByClassWithContext(ctx context.Context, request *ReadNumGroupByClassRequest, runtime *dara.RuntimeOptions) (_result *ReadNumGroupByClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadNumGroupByClass"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadNumGroupByClassResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Method description: Get information under all categories
//
// @param request - ReadNumGroupTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadNumGroupTotalResponse
func (client *Client) ReadNumGroupTotalWithContext(ctx context.Context, request *ReadNumGroupTotalRequest, runtime *dara.RuntimeOptions) (_result *ReadNumGroupTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadNumGroupTotal"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadNumGroupTotalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves revision records.
//
// @param tmpReq - ReadRevisionHistoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadRevisionHistoryListResponse
func (client *Client) ReadRevisionHistoryListWithContext(ctx context.Context, tmpReq *ReadRevisionHistoryListRequest, runtime *dara.RuntimeOptions) (_result *ReadRevisionHistoryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReadRevisionHistoryListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PageInfo) {
		request.PageInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PageInfo, dara.String("PageInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.CategoryCode) {
		body["CategoryCode"] = request.CategoryCode
	}

	if !dara.IsNil(request.ChannelGroupCode) {
		body["ChannelGroupCode"] = request.ChannelGroupCode
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.PageInfoShrink) {
		body["PageInfo"] = request.PageInfoShrink
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadRevisionHistoryList"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadRevisionHistoryListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the subscription list of a user.
//
// @param request - ReadUserSubscriptionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadUserSubscriptionListResponse
func (client *Client) ReadUserSubscriptionListWithContext(ctx context.Context, request *ReadUserSubscriptionListRequest, runtime *dara.RuntimeOptions) (_result *ReadUserSubscriptionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.CategoryGroupCode) {
		body["CategoryGroupCode"] = request.CategoryGroupCode
	}

	if !dara.IsNil(request.ChannelGroupCode) {
		body["ChannelGroupCode"] = request.ChannelGroupCode
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadUserSubscriptionList"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadUserSubscriptionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a webhook contact.
//
// @param request - ReadWebhookContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadWebhookContactResponse
func (client *Client) ReadWebhookContactWithContext(ctx context.Context, request *ReadWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *ReadWebhookContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadWebhookContact"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadWebhookContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of webhook send templates.
//
// @param request - ReadWebhookContactSendTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadWebhookContactSendTemplateListResponse
func (client *Client) ReadWebhookContactSendTemplateListWithContext(ctx context.Context, request *ReadWebhookContactSendTemplateListRequest, runtime *dara.RuntimeOptions) (_result *ReadWebhookContactSendTemplateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadWebhookContactSendTemplateList"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadWebhookContactSendTemplateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets user subscriptions.
//
// @param tmpReq - ResetUserSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetUserSubscriptionResponse
func (client *Client) ResetUserSubscriptionWithContext(ctx context.Context, tmpReq *ResetUserSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *ResetUserSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResetUserSubscriptionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CategoryCodes) {
		request.CategoryCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryCodes, dara.String("CategoryCodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.CategoryCodesShrink) {
		body["CategoryCodes"] = request.CategoryCodesShrink
	}

	if !dara.IsNil(request.ChannelGroupCode) {
		body["ChannelGroupCode"] = request.ChannelGroupCode
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.Remarks) {
		body["Remarks"] = request.Remarks
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetUserSubscription"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetUserSubscriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests a webhook.
//
// @param request - TestWebhookContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestWebhookContactResponse
func (client *Client) TestWebhookContactWithContext(ctx context.Context, request *TestWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *TestWebhookContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.BotSecurityToken) {
		body["BotSecurityToken"] = request.BotSecurityToken
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.ServerUrl) {
		body["ServerUrl"] = request.ServerUrl
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	if !dara.IsNil(request.WebhookType) {
		body["WebhookType"] = request.WebhookType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestWebhookContact"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestWebhookContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates marketing preferences.
//
// @param request - UpdateMarketingPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMarketingPreferenceResponse
func (client *Client) UpdateMarketingPreferenceWithContext(ctx context.Context, request *UpdateMarketingPreferenceRequest, runtime *dara.RuntimeOptions) (_result *UpdateMarketingPreferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AllowMarketing) {
		body["AllowMarketing"] = request.AllowMarketing
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMarketingPreference"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMarketingPreferenceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the message language.
//
// @param request - UpdateMessageLanguageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageLanguageResponse
func (client *Client) UpdateMessageLanguageWithContext(ctx context.Context, request *UpdateMessageLanguageRequest, runtime *dara.RuntimeOptions) (_result *UpdateMessageLanguageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.PreferLang) {
		body["PreferLang"] = request.PreferLang
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMessageLanguage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMessageLanguageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a webhook contact.
//
// @param request - UpdateWebhookContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWebhookContactResponse
func (client *Client) UpdateWebhookContactWithContext(ctx context.Context, request *UpdateWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *UpdateWebhookContactResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.BotSecurityToken) {
		body["BotSecurityToken"] = request.BotSecurityToken
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SecurityToken) {
		body["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServerUrl) {
		body["ServerUrl"] = request.ServerUrl
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	if !dara.IsNil(request.VerificationCode) {
		body["VerificationCode"] = request.VerificationCode
	}

	if !dara.IsNil(request.WebhookType) {
		body["WebhookType"] = request.WebhookType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWebhookContact"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWebhookContactResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
