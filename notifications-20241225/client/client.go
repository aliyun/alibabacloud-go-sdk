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
		"ap-southeast-1": dara.String("notifications-intl.aliyuncs.com"),
		"cn-zhangjiakou": dara.String("notifications.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("notifications"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a webhook contact.
//
// @param request - CreateWebhookContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWebhookContactResponse
func (client *Client) CreateWebhookContactWithOptions(request *CreateWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *CreateWebhookContactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a webhook contact.
//
// @param request - CreateWebhookContactRequest
//
// @return CreateWebhookContactResponse
func (client *Client) CreateWebhookContact(request *CreateWebhookContactRequest) (_result *CreateWebhookContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWebhookContactResponse{}
	_body, _err := client.CreateWebhookContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DelMessageWithOptions(request *DelMessageRequest, runtime *dara.RuntimeOptions) (_result *DelMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DelMessageResponse
func (client *Client) DelMessage(request *DelMessageRequest) (_result *DelMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelMessageResponse{}
	_body, _err := client.DelMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAllMessageWithOptions(request *DeleteAllMessageRequest, runtime *dara.RuntimeOptions) (_result *DeleteAllMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAllMessageResponse
func (client *Client) DeleteAllMessage(request *DeleteAllMessageRequest) (_result *DeleteAllMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAllMessageResponse{}
	_body, _err := client.DeleteAllMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteWebhookContactWithOptions(request *DeleteWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebhookContactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteWebhookContactResponse
func (client *Client) DeleteWebhookContact(request *DeleteWebhookContactRequest) (_result *DeleteWebhookContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWebhookContactResponse{}
	_body, _err := client.DeleteWebhookContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadAllCommonContactsWithOptions(request *ReadAllCommonContactsRequest, runtime *dara.RuntimeOptions) (_result *ReadAllCommonContactsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadAllCommonContactsResponse
func (client *Client) ReadAllCommonContacts(request *ReadAllCommonContactsRequest) (_result *ReadAllCommonContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadAllCommonContactsResponse{}
	_body, _err := client.ReadAllCommonContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadAllMarketingPreferencesWithOptions(request *ReadAllMarketingPreferencesRequest, runtime *dara.RuntimeOptions) (_result *ReadAllMarketingPreferencesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadAllMarketingPreferencesResponse
func (client *Client) ReadAllMarketingPreferences(request *ReadAllMarketingPreferencesRequest) (_result *ReadAllMarketingPreferencesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadAllMarketingPreferencesResponse{}
	_body, _err := client.ReadAllMarketingPreferencesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadAllMessageWithOptions(request *ReadAllMessageRequest, runtime *dara.RuntimeOptions) (_result *ReadAllMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadAllMessageResponse
func (client *Client) ReadAllMessage(request *ReadAllMessageRequest) (_result *ReadAllMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadAllMessageResponse{}
	_body, _err := client.ReadAllMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadAllWebhookContactsWithOptions(request *ReadAllWebhookContactsRequest, runtime *dara.RuntimeOptions) (_result *ReadAllWebhookContactsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadAllWebhookContactsResponse
func (client *Client) ReadAllWebhookContacts(request *ReadAllWebhookContactsRequest) (_result *ReadAllWebhookContactsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadAllWebhookContactsResponse{}
	_body, _err := client.ReadAllWebhookContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadCategoryGroupListWithOptions(request *ReadCategoryGroupListRequest, runtime *dara.RuntimeOptions) (_result *ReadCategoryGroupListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadCategoryGroupListResponse
func (client *Client) ReadCategoryGroupList(request *ReadCategoryGroupListRequest) (_result *ReadCategoryGroupListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadCategoryGroupListResponse{}
	_body, _err := client.ReadCategoryGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadClassNameWithOptions(request *ReadClassNameRequest, runtime *dara.RuntimeOptions) (_result *ReadClassNameResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadClassNameResponse
func (client *Client) ReadClassName(request *ReadClassNameRequest) (_result *ReadClassNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadClassNameResponse{}
	_body, _err := client.ReadClassNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadCommonContactWithOptions(request *ReadCommonContactRequest, runtime *dara.RuntimeOptions) (_result *ReadCommonContactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadCommonContactResponse
func (client *Client) ReadCommonContact(request *ReadCommonContactRequest) (_result *ReadCommonContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadCommonContactResponse{}
	_body, _err := client.ReadCommonContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadMarketingPreferenceWithOptions(request *ReadMarketingPreferenceRequest, runtime *dara.RuntimeOptions) (_result *ReadMarketingPreferenceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadMarketingPreferenceResponse
func (client *Client) ReadMarketingPreference(request *ReadMarketingPreferenceRequest) (_result *ReadMarketingPreferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMarketingPreferenceResponse{}
	_body, _err := client.ReadMarketingPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadMessageWithOptions(request *ReadMessageRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadMessageResponse
func (client *Client) ReadMessage(request *ReadMessageRequest) (_result *ReadMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageResponse{}
	_body, _err := client.ReadMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadMessageContentWithOptions(request *ReadMessageContentRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadMessageContentResponse
func (client *Client) ReadMessageContent(request *ReadMessageContentRequest) (_result *ReadMessageContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageContentResponse{}
	_body, _err := client.ReadMessageContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadMessageLanguageWithOptions(request *ReadMessageLanguageRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageLanguageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadMessageLanguageResponse
func (client *Client) ReadMessageLanguage(request *ReadMessageLanguageRequest) (_result *ReadMessageLanguageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageLanguageResponse{}
	_body, _err := client.ReadMessageLanguageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadMessageListWithOptions(request *ReadMessageListRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadMessageListResponse
func (client *Client) ReadMessageList(request *ReadMessageListRequest) (_result *ReadMessageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageListResponse{}
	_body, _err := client.ReadMessageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadMessageNewTotalWithOptions(request *ReadMessageNewTotalRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageNewTotalResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadMessageNewTotalResponse
func (client *Client) ReadMessageNewTotal(request *ReadMessageNewTotalRequest) (_result *ReadMessageNewTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageNewTotalResponse{}
	_body, _err := client.ReadMessageNewTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadMetaConfigWithOptions(request *ReadMetaConfigRequest, runtime *dara.RuntimeOptions) (_result *ReadMetaConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadMetaConfigResponse
func (client *Client) ReadMetaConfig(request *ReadMetaConfigRequest) (_result *ReadMetaConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMetaConfigResponse{}
	_body, _err := client.ReadMetaConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadNumGroupByClassWithOptions(request *ReadNumGroupByClassRequest, runtime *dara.RuntimeOptions) (_result *ReadNumGroupByClassResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadNumGroupByClassResponse
func (client *Client) ReadNumGroupByClass(request *ReadNumGroupByClassRequest) (_result *ReadNumGroupByClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadNumGroupByClassResponse{}
	_body, _err := client.ReadNumGroupByClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadNumGroupTotalWithOptions(request *ReadNumGroupTotalRequest, runtime *dara.RuntimeOptions) (_result *ReadNumGroupTotalResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadNumGroupTotalResponse
func (client *Client) ReadNumGroupTotal(request *ReadNumGroupTotalRequest) (_result *ReadNumGroupTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadNumGroupTotalResponse{}
	_body, _err := client.ReadNumGroupTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadRevisionHistoryListWithOptions(tmpReq *ReadRevisionHistoryListRequest, runtime *dara.RuntimeOptions) (_result *ReadRevisionHistoryListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ReadRevisionHistoryListRequest
//
// @return ReadRevisionHistoryListResponse
func (client *Client) ReadRevisionHistoryList(request *ReadRevisionHistoryListRequest) (_result *ReadRevisionHistoryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadRevisionHistoryListResponse{}
	_body, _err := client.ReadRevisionHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadUserSubscriptionListWithOptions(request *ReadUserSubscriptionListRequest, runtime *dara.RuntimeOptions) (_result *ReadUserSubscriptionListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadUserSubscriptionListResponse
func (client *Client) ReadUserSubscriptionList(request *ReadUserSubscriptionListRequest) (_result *ReadUserSubscriptionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadUserSubscriptionListResponse{}
	_body, _err := client.ReadUserSubscriptionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadWebhookContactWithOptions(request *ReadWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *ReadWebhookContactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadWebhookContactResponse
func (client *Client) ReadWebhookContact(request *ReadWebhookContactRequest) (_result *ReadWebhookContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadWebhookContactResponse{}
	_body, _err := client.ReadWebhookContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReadWebhookContactSendTemplateListWithOptions(request *ReadWebhookContactSendTemplateListRequest, runtime *dara.RuntimeOptions) (_result *ReadWebhookContactSendTemplateListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReadWebhookContactSendTemplateListResponse
func (client *Client) ReadWebhookContactSendTemplateList(request *ReadWebhookContactSendTemplateListRequest) (_result *ReadWebhookContactSendTemplateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadWebhookContactSendTemplateListResponse{}
	_body, _err := client.ReadWebhookContactSendTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ResetUserSubscriptionWithOptions(tmpReq *ResetUserSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *ResetUserSubscriptionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ResetUserSubscriptionRequest
//
// @return ResetUserSubscriptionResponse
func (client *Client) ResetUserSubscription(request *ResetUserSubscriptionRequest) (_result *ResetUserSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetUserSubscriptionResponse{}
	_body, _err := client.ResetUserSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TestWebhookContactWithOptions(request *TestWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *TestWebhookContactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TestWebhookContactResponse
func (client *Client) TestWebhookContact(request *TestWebhookContactRequest) (_result *TestWebhookContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestWebhookContactResponse{}
	_body, _err := client.TestWebhookContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMarketingPreferenceWithOptions(request *UpdateMarketingPreferenceRequest, runtime *dara.RuntimeOptions) (_result *UpdateMarketingPreferenceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateMarketingPreferenceResponse
func (client *Client) UpdateMarketingPreference(request *UpdateMarketingPreferenceRequest) (_result *UpdateMarketingPreferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMarketingPreferenceResponse{}
	_body, _err := client.UpdateMarketingPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMessageLanguageWithOptions(request *UpdateMessageLanguageRequest, runtime *dara.RuntimeOptions) (_result *UpdateMessageLanguageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateMessageLanguageResponse
func (client *Client) UpdateMessageLanguage(request *UpdateMessageLanguageRequest) (_result *UpdateMessageLanguageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMessageLanguageResponse{}
	_body, _err := client.UpdateMessageLanguageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateWebhookContactWithOptions(request *UpdateWebhookContactRequest, runtime *dara.RuntimeOptions) (_result *UpdateWebhookContactResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateWebhookContactResponse
func (client *Client) UpdateWebhookContact(request *UpdateWebhookContactRequest) (_result *UpdateWebhookContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWebhookContactResponse{}
	_body, _err := client.UpdateWebhookContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
