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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("account-crm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AccountOneKeyDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountOneKeyDeleteResponse
func (client *Client) AccountOneKeyDeleteWithOptions(request *AccountOneKeyDeleteRequest, runtime *dara.RuntimeOptions) (_result *AccountOneKeyDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccountOneKeyDelete"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AccountOneKeyDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AccountOneKeyDeleteRequest
//
// @return AccountOneKeyDeleteResponse
func (client *Client) AccountOneKeyDelete(request *AccountOneKeyDeleteRequest) (_result *AccountOneKeyDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AccountOneKeyDeleteResponse{}
	_body, _err := client.AccountOneKeyDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomerLabelResponse
func (client *Client) AddCustomerLabelWithOptions(request *AddCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *AddCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Endtime) {
		query["Endtime"] = request.Endtime
	}

	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.LabelTypes) {
		query["LabelTypes"] = request.LabelTypes
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomerLabel"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomerLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddCustomerLabelRequest
//
// @return AddCustomerLabelResponse
func (client *Client) AddCustomerLabel(request *AddCustomerLabelRequest) (_result *AddCustomerLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCustomerLabelResponse{}
	_body, _err := client.AddCustomerLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AllowAgAccountLoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllowAgAccountLoginResponse
func (client *Client) AllowAgAccountLoginWithOptions(request *AllowAgAccountLoginRequest, runtime *dara.RuntimeOptions) (_result *AllowAgAccountLoginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllowAgAccountLogin"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllowAgAccountLoginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AllowAgAccountLoginRequest
//
// @return AllowAgAccountLoginResponse
func (client *Client) AllowAgAccountLogin(request *AllowAgAccountLoginRequest) (_result *AllowAgAccountLoginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllowAgAccountLoginResponse{}
	_body, _err := client.AllowAgAccountLoginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申请ag注销
//
// @param request - ApplyAgOneKeyDeleteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAgOneKeyDeleteTaskResponse
func (client *Client) ApplyAgOneKeyDeleteTaskWithOptions(request *ApplyAgOneKeyDeleteTaskRequest, runtime *dara.RuntimeOptions) (_result *ApplyAgOneKeyDeleteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbandonedDependency) {
		query["AbandonedDependency"] = request.AbandonedDependency
	}

	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyAgOneKeyDeleteTask"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyAgOneKeyDeleteTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请ag注销
//
// @param request - ApplyAgOneKeyDeleteTaskRequest
//
// @return ApplyAgOneKeyDeleteTaskResponse
func (client *Client) ApplyAgOneKeyDeleteTask(request *ApplyAgOneKeyDeleteTaskRequest) (_result *ApplyAgOneKeyDeleteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyAgOneKeyDeleteTaskResponse{}
	_body, _err := client.ApplyAgOneKeyDeleteTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申请ag注销
//
// @param request - ApplyAgOneKeyOnlyCheckerTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAgOneKeyOnlyCheckerTaskResponse
func (client *Client) ApplyAgOneKeyOnlyCheckerTaskWithOptions(request *ApplyAgOneKeyOnlyCheckerTaskRequest, runtime *dara.RuntimeOptions) (_result *ApplyAgOneKeyOnlyCheckerTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyAgOneKeyOnlyCheckerTask"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyAgOneKeyOnlyCheckerTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请ag注销
//
// @param request - ApplyAgOneKeyOnlyCheckerTaskRequest
//
// @return ApplyAgOneKeyOnlyCheckerTaskResponse
func (client *Client) ApplyAgOneKeyOnlyCheckerTask(request *ApplyAgOneKeyOnlyCheckerTaskRequest) (_result *ApplyAgOneKeyOnlyCheckerTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyAgOneKeyOnlyCheckerTaskResponse{}
	_body, _err := client.ApplyAgOneKeyOnlyCheckerTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ApplyIdentityRegistrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyIdentityRegistrationResponse
func (client *Client) ApplyIdentityRegistrationWithOptions(request *ApplyIdentityRegistrationRequest, runtime *dara.RuntimeOptions) (_result *ApplyIdentityRegistrationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.DocBackPic) {
		query["DocBackPic"] = request.DocBackPic
	}

	if !dara.IsNil(request.DocFrontPic) {
		query["DocFrontPic"] = request.DocFrontPic
	}

	if !dara.IsNil(request.DocNum) {
		query["DocNum"] = request.DocNum
	}

	if !dara.IsNil(request.DocType) {
		query["DocType"] = request.DocType
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.FullName) {
		query["FullName"] = request.FullName
	}

	if !dara.IsNil(request.RegisteredAddress) {
		query["RegisteredAddress"] = request.RegisteredAddress
	}

	if !dara.IsNil(request.RegisteredCountry) {
		query["RegisteredCountry"] = request.RegisteredCountry
	}

	if !dara.IsNil(request.RegisteredNum) {
		query["RegisteredNum"] = request.RegisteredNum
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Tel) {
		query["Tel"] = request.Tel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyIdentityRegistration"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyIdentityRegistrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ApplyIdentityRegistrationRequest
//
// @return ApplyIdentityRegistrationResponse
func (client *Client) ApplyIdentityRegistration(request *ApplyIdentityRegistrationRequest) (_result *ApplyIdentityRegistrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyIdentityRegistrationResponse{}
	_body, _err := client.ApplyIdentityRegistrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AsyncCreateAgAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncCreateAgAccountResponse
func (client *Client) AsyncCreateAgAccountWithOptions(request *AsyncCreateAgAccountRequest, runtime *dara.RuntimeOptions) (_result *AsyncCreateAgAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoginEmail) {
		query["LoginEmail"] = request.LoginEmail
	}

	if !dara.IsNil(request.MaserAccountInfo) {
		query["MaserAccountInfo"] = request.MaserAccountInfo
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsyncCreateAgAccount"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsyncCreateAgAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AsyncCreateAgAccountRequest
//
// @return AsyncCreateAgAccountResponse
func (client *Client) AsyncCreateAgAccount(request *AsyncCreateAgAccountRequest) (_result *AsyncCreateAgAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsyncCreateAgAccountResponse{}
	_body, _err := client.AsyncCreateAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AsyncModifyAgLoginEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AsyncModifyAgLoginEmailResponse
func (client *Client) AsyncModifyAgLoginEmailWithOptions(request *AsyncModifyAgLoginEmailRequest, runtime *dara.RuntimeOptions) (_result *AsyncModifyAgLoginEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.NewLoginEmail) {
		query["NewLoginEmail"] = request.NewLoginEmail
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AsyncModifyAgLoginEmail"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AsyncModifyAgLoginEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AsyncModifyAgLoginEmailRequest
//
// @return AsyncModifyAgLoginEmailResponse
func (client *Client) AsyncModifyAgLoginEmail(request *AsyncModifyAgLoginEmailRequest) (_result *AsyncModifyAgLoginEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AsyncModifyAgLoginEmailResponse{}
	_body, _err := client.AsyncModifyAgLoginEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AuthAndActiveWithHidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthAndActiveWithHidResponse
func (client *Client) AuthAndActiveWithHidWithOptions(request *AuthAndActiveWithHidRequest, runtime *dara.RuntimeOptions) (_result *AuthAndActiveWithHidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthAndActiveWithHid"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthAndActiveWithHidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthAndActiveWithHidRequest
//
// @return AuthAndActiveWithHidResponse
func (client *Client) AuthAndActiveWithHid(request *AuthAndActiveWithHidRequest) (_result *AuthAndActiveWithHidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthAndActiveWithHidResponse{}
	_body, _err := client.AuthAndActiveWithHidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AuthAndRefreshLoginTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthAndRefreshLoginTicketResponse
func (client *Client) AuthAndRefreshLoginTicketWithOptions(request *AuthAndRefreshLoginTicketRequest, runtime *dara.RuntimeOptions) (_result *AuthAndRefreshLoginTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthAndRefreshLoginTicket"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthAndRefreshLoginTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthAndRefreshLoginTicketRequest
//
// @return AuthAndRefreshLoginTicketResponse
func (client *Client) AuthAndRefreshLoginTicket(request *AuthAndRefreshLoginTicketRequest) (_result *AuthAndRefreshLoginTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthAndRefreshLoginTicketResponse{}
	_body, _err := client.AuthAndRefreshLoginTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AuthLoginTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthLoginTicketResponse
func (client *Client) AuthLoginTicketWithOptions(request *AuthLoginTicketRequest, runtime *dara.RuntimeOptions) (_result *AuthLoginTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.MinorAuthCode) {
		query["MinorAuthCode"] = request.MinorAuthCode
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthLoginTicket"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthLoginTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthLoginTicketRequest
//
// @return AuthLoginTicketResponse
func (client *Client) AuthLoginTicket(request *AuthLoginTicketRequest) (_result *AuthLoginTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthLoginTicketResponse{}
	_body, _err := client.AuthLoginTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BatchQueryAgAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryAgAccountResponse
func (client *Client) BatchQueryAgAccountWithOptions(request *BatchQueryAgAccountRequest, runtime *dara.RuntimeOptions) (_result *BatchQueryAgAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.PkList) {
		query["PkList"] = request.PkList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchQueryAgAccount"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchQueryAgAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchQueryAgAccountRequest
//
// @return BatchQueryAgAccountResponse
func (client *Client) BatchQueryAgAccount(request *BatchQueryAgAccountRequest) (_result *BatchQueryAgAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchQueryAgAccountResponse{}
	_body, _err := client.BatchQueryAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BatchQueryCreateAccountTraceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryCreateAccountTraceResponse
func (client *Client) BatchQueryCreateAccountTraceWithOptions(request *BatchQueryCreateAccountTraceRequest, runtime *dara.RuntimeOptions) (_result *BatchQueryCreateAccountTraceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.TraceNoList) {
		query["TraceNoList"] = request.TraceNoList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchQueryCreateAccountTrace"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchQueryCreateAccountTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchQueryCreateAccountTraceRequest
//
// @return BatchQueryCreateAccountTraceResponse
func (client *Client) BatchQueryCreateAccountTrace(request *BatchQueryCreateAccountTraceRequest) (_result *BatchQueryCreateAccountTraceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchQueryCreateAccountTraceResponse{}
	_body, _err := client.BatchQueryCreateAccountTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BatchQueryModifyLoginEmailTraceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryModifyLoginEmailTraceResponse
func (client *Client) BatchQueryModifyLoginEmailTraceWithOptions(request *BatchQueryModifyLoginEmailTraceRequest, runtime *dara.RuntimeOptions) (_result *BatchQueryModifyLoginEmailTraceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.TraceNoList) {
		query["TraceNoList"] = request.TraceNoList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchQueryModifyLoginEmailTrace"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchQueryModifyLoginEmailTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchQueryModifyLoginEmailTraceRequest
//
// @return BatchQueryModifyLoginEmailTraceResponse
func (client *Client) BatchQueryModifyLoginEmailTrace(request *BatchQueryModifyLoginEmailTraceRequest) (_result *BatchQueryModifyLoginEmailTraceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchQueryModifyLoginEmailTraceResponse{}
	_body, _err := client.BatchQueryModifyLoginEmailTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CancelAsyncCreateAgAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAsyncCreateAgAccountResponse
func (client *Client) CancelAsyncCreateAgAccountWithOptions(request *CancelAsyncCreateAgAccountRequest, runtime *dara.RuntimeOptions) (_result *CancelAsyncCreateAgAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.TraceNo) {
		query["TraceNo"] = request.TraceNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAsyncCreateAgAccount"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAsyncCreateAgAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CancelAsyncCreateAgAccountRequest
//
// @return CancelAsyncCreateAgAccountResponse
func (client *Client) CancelAsyncCreateAgAccount(request *CancelAsyncCreateAgAccountRequest) (_result *CancelAsyncCreateAgAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelAsyncCreateAgAccountResponse{}
	_body, _err := client.CancelAsyncCreateAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CancelAsyncModifyLoginEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAsyncModifyLoginEmailResponse
func (client *Client) CancelAsyncModifyLoginEmailWithOptions(request *CancelAsyncModifyLoginEmailRequest, runtime *dara.RuntimeOptions) (_result *CancelAsyncModifyLoginEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.TraceNo) {
		query["TraceNo"] = request.TraceNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAsyncModifyLoginEmail"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAsyncModifyLoginEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CancelAsyncModifyLoginEmailRequest
//
// @return CancelAsyncModifyLoginEmailResponse
func (client *Client) CancelAsyncModifyLoginEmail(request *CancelAsyncModifyLoginEmailRequest) (_result *CancelAsyncModifyLoginEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelAsyncModifyLoginEmailResponse{}
	_body, _err := client.CancelAsyncModifyLoginEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ChangeAgAccountNationalityCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAgAccountNationalityCodeResponse
func (client *Client) ChangeAgAccountNationalityCodeWithOptions(request *ChangeAgAccountNationalityCodeRequest, runtime *dara.RuntimeOptions) (_result *ChangeAgAccountNationalityCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.NationalityCode) {
		query["NationalityCode"] = request.NationalityCode
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeAgAccountNationalityCode"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeAgAccountNationalityCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeAgAccountNationalityCodeRequest
//
// @return ChangeAgAccountNationalityCodeResponse
func (client *Client) ChangeAgAccountNationalityCode(request *ChangeAgAccountNationalityCodeRequest) (_result *ChangeAgAccountNationalityCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeAgAccountNationalityCodeResponse{}
	_body, _err := client.ChangeAgAccountNationalityCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ChangeAgSecurityEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAgSecurityEmailResponse
func (client *Client) ChangeAgSecurityEmailWithOptions(request *ChangeAgSecurityEmailRequest, runtime *dara.RuntimeOptions) (_result *ChangeAgSecurityEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.SecurityEmail) {
		query["SecurityEmail"] = request.SecurityEmail
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeAgSecurityEmail"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeAgSecurityEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeAgSecurityEmailRequest
//
// @return ChangeAgSecurityEmailResponse
func (client *Client) ChangeAgSecurityEmail(request *ChangeAgSecurityEmailRequest) (_result *ChangeAgSecurityEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeAgSecurityEmailResponse{}
	_body, _err := client.ChangeAgSecurityEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ChangeAgSecurityMobileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAgSecurityMobileResponse
func (client *Client) ChangeAgSecurityMobileWithOptions(request *ChangeAgSecurityMobileRequest, runtime *dara.RuntimeOptions) (_result *ChangeAgSecurityMobileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.SecurityMobile) {
		query["SecurityMobile"] = request.SecurityMobile
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeAgSecurityMobile"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeAgSecurityMobileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeAgSecurityMobileRequest
//
// @return ChangeAgSecurityMobileResponse
func (client *Client) ChangeAgSecurityMobile(request *ChangeAgSecurityMobileRequest) (_result *ChangeAgSecurityMobileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeAgSecurityMobileResponse{}
	_body, _err := client.ChangeAgSecurityMobileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAccountProfileInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountProfileInfoResponse
func (client *Client) CreateAccountProfileInfoWithOptions(request *CreateAccountProfileInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountProfileInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountJson) {
		query["AccountJson"] = request.AccountJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccountProfileInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccountProfileInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAccountProfileInfoRequest
//
// @return CreateAccountProfileInfoResponse
func (client *Client) CreateAccountProfileInfo(request *CreateAccountProfileInfoRequest) (_result *CreateAccountProfileInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccountProfileInfoResponse{}
	_body, _err := client.CreateAccountProfileInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAgAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgAccountResponse
func (client *Client) CreateAgAccountWithOptions(request *CreateAgAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAgAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoginEmail) {
		query["LoginEmail"] = request.LoginEmail
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.NationCode) {
		query["NationCode"] = request.NationCode
	}

	if !dara.IsNil(request.Own) {
		query["Own"] = request.Own
	}

	if !dara.IsNil(request.RealParentPk) {
		query["RealParentPk"] = request.RealParentPk
	}

	if !dara.IsNil(request.SecurityMobile) {
		query["SecurityMobile"] = request.SecurityMobile
	}

	if !dara.IsNil(request.ShowNickName) {
		query["ShowNickName"] = request.ShowNickName
	}

	if !dara.IsNil(request.SiteNick) {
		query["SiteNick"] = request.SiteNick
	}

	if !dara.IsNil(request.SrcAccountInfo) {
		query["srcAccountInfo"] = request.SrcAccountInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgAccount"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAgAccountRequest
//
// @return CreateAgAccountResponse
func (client *Client) CreateAgAccount(request *CreateAgAccountRequest) (_result *CreateAgAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAgAccountResponse{}
	_body, _err := client.CreateAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateContacterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContacterResponse
func (client *Client) CreateContacterWithOptions(request *CreateContacterRequest, runtime *dara.RuntimeOptions) (_result *CreateContacterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContacterAddress) {
		query["ContacterAddress"] = request.ContacterAddress
	}

	if !dara.IsNil(request.ContacterDingding) {
		query["ContacterDingding"] = request.ContacterDingding
	}

	if !dara.IsNil(request.ContacterEmail) {
		query["ContacterEmail"] = request.ContacterEmail
	}

	if !dara.IsNil(request.ContacterMobile) {
		query["ContacterMobile"] = request.ContacterMobile
	}

	if !dara.IsNil(request.ContacterName) {
		query["ContacterName"] = request.ContacterName
	}

	if !dara.IsNil(request.ContacterPosition) {
		query["ContacterPosition"] = request.ContacterPosition
	}

	if !dara.IsNil(request.ContacterStaffNo) {
		query["ContacterStaffNo"] = request.ContacterStaffNo
	}

	if !dara.IsNil(request.ContacterType) {
		query["ContacterType"] = request.ContacterType
	}

	if !dara.IsNil(request.ContacterWangwang) {
		query["ContacterWangwang"] = request.ContacterWangwang
	}

	if !dara.IsNil(request.EmailConfirmed) {
		query["EmailConfirmed"] = request.EmailConfirmed
	}

	if !dara.IsNil(request.MobileConfirmed) {
		query["MobileConfirmed"] = request.MobileConfirmed
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContacter"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContacterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateContacterRequest
//
// @return CreateContacterResponse
func (client *Client) CreateContacter(request *CreateContacterRequest) (_result *CreateContacterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateContacterResponse{}
	_body, _err := client.CreateContacterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateRealNameCertificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRealNameCertificationResponse
func (client *Client) CreateRealNameCertificationWithOptions(request *CreateRealNameCertificationRequest, runtime *dara.RuntimeOptions) (_result *CreateRealNameCertificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountCertifyType) {
		query["AccountCertifyType"] = request.AccountCertifyType
	}

	if !dara.IsNil(request.CorporateLicenseNumber) {
		query["CorporateLicenseNumber"] = request.CorporateLicenseNumber
	}

	if !dara.IsNil(request.CorporateName) {
		query["CorporateName"] = request.CorporateName
	}

	if !dara.IsNil(request.LicenseNumber) {
		query["LicenseNumber"] = request.LicenseNumber
	}

	if !dara.IsNil(request.LicenseType) {
		query["LicenseType"] = request.LicenseType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRealNameCertification"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRealNameCertificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateRealNameCertificationRequest
//
// @return CreateRealNameCertificationResponse
func (client *Client) CreateRealNameCertification(request *CreateRealNameCertificationRequest) (_result *CreateRealNameCertificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRealNameCertificationResponse{}
	_body, _err := client.CreateRealNameCertificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CustomerSensitiveInfoLogicalDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerSensitiveInfoLogicalDeleteResponse
func (client *Client) CustomerSensitiveInfoLogicalDeleteWithOptions(request *CustomerSensitiveInfoLogicalDeleteRequest, runtime *dara.RuntimeOptions) (_result *CustomerSensitiveInfoLogicalDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CustomerSensitiveInfoLogicalDelete"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CustomerSensitiveInfoLogicalDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CustomerSensitiveInfoLogicalDeleteRequest
//
// @return CustomerSensitiveInfoLogicalDeleteResponse
func (client *Client) CustomerSensitiveInfoLogicalDelete(request *CustomerSensitiveInfoLogicalDeleteRequest) (_result *CustomerSensitiveInfoLogicalDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CustomerSensitiveInfoLogicalDeleteResponse{}
	_body, _err := client.CustomerSensitiveInfoLogicalDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CustomerSensitiveInfoPhysicalDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CustomerSensitiveInfoPhysicalDeleteResponse
func (client *Client) CustomerSensitiveInfoPhysicalDeleteWithOptions(request *CustomerSensitiveInfoPhysicalDeleteRequest, runtime *dara.RuntimeOptions) (_result *CustomerSensitiveInfoPhysicalDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CustomerSensitiveInfoPhysicalDelete"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CustomerSensitiveInfoPhysicalDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CustomerSensitiveInfoPhysicalDeleteRequest
//
// @return CustomerSensitiveInfoPhysicalDeleteResponse
func (client *Client) CustomerSensitiveInfoPhysicalDelete(request *CustomerSensitiveInfoPhysicalDeleteRequest) (_result *CustomerSensitiveInfoPhysicalDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CustomerSensitiveInfoPhysicalDeleteResponse{}
	_body, _err := client.CustomerSensitiveInfoPhysicalDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步Del缓存操作
//
// @param request - DelCacheOperateSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelCacheOperateSyncResponse
func (client *Client) DelCacheOperateSyncWithOptions(request *DelCacheOperateSyncRequest, runtime *dara.RuntimeOptions) (_result *DelCacheOperateSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelCacheOperateSync"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelCacheOperateSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步Del缓存操作
//
// @param request - DelCacheOperateSyncRequest
//
// @return DelCacheOperateSyncResponse
func (client *Client) DelCacheOperateSync(request *DelCacheOperateSyncRequest) (_result *DelCacheOperateSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelCacheOperateSyncResponse{}
	_body, _err := client.DelCacheOperateSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteContacterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContacterResponse
func (client *Client) DeleteContacterWithOptions(request *DeleteContacterRequest, runtime *dara.RuntimeOptions) (_result *DeleteContacterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContacterId) {
		query["ContacterId"] = request.ContacterId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContacter"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContacterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteContacterRequest
//
// @return DeleteContacterResponse
func (client *Client) DeleteContacter(request *DeleteContacterRequest) (_result *DeleteContacterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContacterResponse{}
	_body, _err := client.DeleteContacterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomerLabelResponse
func (client *Client) DeleteCustomerLabelWithOptions(request *DeleteCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.LabelTypes) {
		query["LabelTypes"] = request.LabelTypes
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomerLabel"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomerLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCustomerLabelRequest
//
// @return DeleteCustomerLabelResponse
func (client *Client) DeleteCustomerLabel(request *DeleteCustomerLabelRequest) (_result *DeleteCustomerLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomerLabelResponse{}
	_body, _err := client.DeleteCustomerLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据outerId查询是否存在绑定关系
//
// @param request - ExistBindsByOuterIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExistBindsByOuterIdResponse
func (client *Client) ExistBindsByOuterIdWithOptions(request *ExistBindsByOuterIdRequest, runtime *dara.RuntimeOptions) (_result *ExistBindsByOuterIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExistBindsByOuterId"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExistBindsByOuterIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据outerId查询是否存在绑定关系
//
// @param request - ExistBindsByOuterIdRequest
//
// @return ExistBindsByOuterIdResponse
func (client *Client) ExistBindsByOuterId(request *ExistBindsByOuterIdRequest) (_result *ExistBindsByOuterIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExistBindsByOuterIdResponse{}
	_body, _err := client.ExistBindsByOuterIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FindAllContacterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindAllContacterResponse
func (client *Client) FindAllContacterWithOptions(request *FindAllContacterRequest, runtime *dara.RuntimeOptions) (_result *FindAllContacterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.LocaleString) {
		query["LocaleString"] = request.LocaleString
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindAllContacter"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindAllContacterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FindAllContacterRequest
//
// @return FindAllContacterResponse
func (client *Client) FindAllContacter(request *FindAllContacterRequest) (_result *FindAllContacterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindAllContacterResponse{}
	_body, _err := client.FindAllContacterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FindBizCategoryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindBizCategoryConfigResponse
func (client *Client) FindBizCategoryConfigWithOptions(request *FindBizCategoryConfigRequest, runtime *dara.RuntimeOptions) (_result *FindBizCategoryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LocaleString) {
		query["LocaleString"] = request.LocaleString
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindBizCategoryConfig"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindBizCategoryConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FindBizCategoryConfigRequest
//
// @return FindBizCategoryConfigResponse
func (client *Client) FindBizCategoryConfig(request *FindBizCategoryConfigRequest) (_result *FindBizCategoryConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindBizCategoryConfigResponse{}
	_body, _err := client.FindBizCategoryConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FindContacterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindContacterResponse
func (client *Client) FindContacterWithOptions(request *FindContacterRequest, runtime *dara.RuntimeOptions) (_result *FindContacterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContacterId) {
		query["ContacterId"] = request.ContacterId
	}

	if !dara.IsNil(request.LocaleString) {
		query["LocaleString"] = request.LocaleString
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindContacter"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindContacterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FindContacterRequest
//
// @return FindContacterResponse
func (client *Client) FindContacter(request *FindContacterRequest) (_result *FindContacterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindContacterResponse{}
	_body, _err := client.FindContacterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FindCustomerInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindCustomerInfoResponse
func (client *Client) FindCustomerInfoWithOptions(request *FindCustomerInfoRequest, runtime *dara.RuntimeOptions) (_result *FindCustomerInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindCustomerInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindCustomerInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FindCustomerInfoRequest
//
// @return FindCustomerInfoResponse
func (client *Client) FindCustomerInfo(request *FindCustomerInfoRequest) (_result *FindCustomerInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindCustomerInfoResponse{}
	_body, _err := client.FindCustomerInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FindCustomerSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindCustomerSnapshotResponse
func (client *Client) FindCustomerSnapshotWithOptions(request *FindCustomerSnapshotRequest, runtime *dara.RuntimeOptions) (_result *FindCustomerSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InfoType) {
		query["InfoType"] = request.InfoType
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindCustomerSnapshot"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindCustomerSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FindCustomerSnapshotRequest
//
// @return FindCustomerSnapshotResponse
func (client *Client) FindCustomerSnapshot(request *FindCustomerSnapshotRequest) (_result *FindCustomerSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindCustomerSnapshotResponse{}
	_body, _err := client.FindCustomerSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FindFinanceTaxRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindFinanceTaxResponse
func (client *Client) FindFinanceTaxWithOptions(request *FindFinanceTaxRequest, runtime *dara.RuntimeOptions) (_result *FindFinanceTaxResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HId) {
		query["HId"] = request.HId
	}

	if !dara.IsNil(request.TaxVersion) {
		query["TaxVersion"] = request.TaxVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindFinanceTax"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindFinanceTaxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FindFinanceTaxRequest
//
// @return FindFinanceTaxResponse
func (client *Client) FindFinanceTax(request *FindFinanceTaxRequest) (_result *FindFinanceTaxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindFinanceTaxResponse{}
	_body, _err := client.FindFinanceTaxWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FindFinanceTaxDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindFinanceTaxDetailResponse
func (client *Client) FindFinanceTaxDetailWithOptions(request *FindFinanceTaxDetailRequest, runtime *dara.RuntimeOptions) (_result *FindFinanceTaxDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KpId) {
		query["KpId"] = request.KpId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindFinanceTaxDetail"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindFinanceTaxDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FindFinanceTaxDetailRequest
//
// @return FindFinanceTaxDetailResponse
func (client *Client) FindFinanceTaxDetail(request *FindFinanceTaxDetailRequest) (_result *FindFinanceTaxDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindFinanceTaxDetailResponse{}
	_body, _err := client.FindFinanceTaxDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 登录过程中根据HID匹配Pk的历史逻辑
//
// @param request - FindPkByHidForLoginWithLegacyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindPkByHidForLoginWithLegacyResponse
func (client *Client) FindPkByHidForLoginWithLegacyWithOptions(request *FindPkByHidForLoginWithLegacyRequest, runtime *dara.RuntimeOptions) (_result *FindPkByHidForLoginWithLegacyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindPkByHidForLoginWithLegacy"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindPkByHidForLoginWithLegacyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 登录过程中根据HID匹配Pk的历史逻辑
//
// @param request - FindPkByHidForLoginWithLegacyRequest
//
// @return FindPkByHidForLoginWithLegacyResponse
func (client *Client) FindPkByHidForLoginWithLegacy(request *FindPkByHidForLoginWithLegacyRequest) (_result *FindPkByHidForLoginWithLegacyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindPkByHidForLoginWithLegacyResponse{}
	_body, _err := client.FindPkByHidForLoginWithLegacyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ForbiddenAgAccountLoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForbiddenAgAccountLoginResponse
func (client *Client) ForbiddenAgAccountLoginWithOptions(request *ForbiddenAgAccountLoginRequest, runtime *dara.RuntimeOptions) (_result *ForbiddenAgAccountLoginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForbiddenAgAccountLogin"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForbiddenAgAccountLoginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ForbiddenAgAccountLoginRequest
//
// @return ForbiddenAgAccountLoginResponse
func (client *Client) ForbiddenAgAccountLogin(request *ForbiddenAgAccountLoginRequest) (_result *ForbiddenAgAccountLoginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ForbiddenAgAccountLoginResponse{}
	_body, _err := client.ForbiddenAgAccountLoginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实名认证url
//
// @param request - GenerateAliyunCertUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAliyunCertUrlResponse
func (client *Client) GenerateAliyunCertUrlWithOptions(request *GenerateAliyunCertUrlRequest, runtime *dara.RuntimeOptions) (_result *GenerateAliyunCertUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunPk) {
		query["AliyunPk"] = request.AliyunPk
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ApplyChannel) {
		query["ApplyChannel"] = request.ApplyChannel
	}

	if !dara.IsNil(request.ApplyType) {
		query["ApplyType"] = request.ApplyType
	}

	if !dara.IsNil(request.Callback) {
		query["Callback"] = request.Callback
	}

	if !dara.IsNil(request.CertWay) {
		query["CertWay"] = request.CertWay
	}

	if !dara.IsNil(request.IgnoreAlreadyCert) {
		query["IgnoreAlreadyCert"] = request.IgnoreAlreadyCert
	}

	if !dara.IsNil(request.IsMobile) {
		query["IsMobile"] = request.IsMobile
	}

	if !dara.IsNil(request.IsOpenApp) {
		query["IsOpenApp"] = request.IsOpenApp
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SubjectType) {
		query["SubjectType"] = request.SubjectType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAliyunCertUrl"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAliyunCertUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实名认证url
//
// @param request - GenerateAliyunCertUrlRequest
//
// @return GenerateAliyunCertUrlResponse
func (client *Client) GenerateAliyunCertUrl(request *GenerateAliyunCertUrlRequest) (_result *GenerateAliyunCertUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateAliyunCertUrlResponse{}
	_body, _err := client.GenerateAliyunCertUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAgAccountAkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgAccountAkResponse
func (client *Client) GetAgAccountAkWithOptions(request *GetAgAccountAkRequest, runtime *dara.RuntimeOptions) (_result *GetAgAccountAkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgAccountAk"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgAccountAkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAgAccountAkRequest
//
// @return GetAgAccountAkResponse
func (client *Client) GetAgAccountAk(request *GetAgAccountAkRequest) (_result *GetAgAccountAkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAgAccountAkResponse{}
	_body, _err := client.GetAgAccountAkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申请ag注销
//
// @param request - GetAgOneKeyDeleteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgOneKeyDeleteTaskResponse
func (client *Client) GetAgOneKeyDeleteTaskWithOptions(request *GetAgOneKeyDeleteTaskRequest, runtime *dara.RuntimeOptions) (_result *GetAgOneKeyDeleteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgOneKeyDeleteTask"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgOneKeyDeleteTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请ag注销
//
// @param request - GetAgOneKeyDeleteTaskRequest
//
// @return GetAgOneKeyDeleteTaskResponse
func (client *Client) GetAgOneKeyDeleteTask(request *GetAgOneKeyDeleteTaskRequest) (_result *GetAgOneKeyDeleteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAgOneKeyDeleteTaskResponse{}
	_body, _err := client.GetAgOneKeyDeleteTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAgRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgRelationResponse
func (client *Client) GetAgRelationWithOptions(request *GetAgRelationRequest, runtime *dara.RuntimeOptions) (_result *GetAgRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgRelation"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAgRelationRequest
//
// @return GetAgRelationResponse
func (client *Client) GetAgRelation(request *GetAgRelationRequest) (_result *GetAgRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAgRelationResponse{}
	_body, _err := client.GetAgRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAliyunIdByPkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAliyunIdByPkResponse
func (client *Client) GetAliyunIdByPkWithOptions(request *GetAliyunIdByPkRequest, runtime *dara.RuntimeOptions) (_result *GetAliyunIdByPkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAliyunIdByPk"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAliyunIdByPkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAliyunIdByPkRequest
//
// @return GetAliyunIdByPkResponse
func (client *Client) GetAliyunIdByPk(request *GetAliyunIdByPkRequest) (_result *GetAliyunIdByPkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAliyunIdByPkResponse{}
	_body, _err := client.GetAliyunIdByPkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAliyunPKByAliyunIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAliyunPKByAliyunIDResponse
func (client *Client) GetAliyunPKByAliyunIDWithOptions(request *GetAliyunPKByAliyunIDRequest, runtime *dara.RuntimeOptions) (_result *GetAliyunPKByAliyunIDResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunId) {
		query["AliyunId"] = request.AliyunId
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAliyunPKByAliyunID"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAliyunPKByAliyunIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAliyunPKByAliyunIDRequest
//
// @return GetAliyunPKByAliyunIDResponse
func (client *Client) GetAliyunPKByAliyunID(request *GetAliyunPKByAliyunIDRequest) (_result *GetAliyunPKByAliyunIDResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAliyunPKByAliyunIDResponse{}
	_body, _err := client.GetAliyunPKByAliyunIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetCustomerCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerCategoryResponse
func (client *Client) GetCustomerCategoryWithOptions(request *GetCustomerCategoryRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LocaleString) {
		query["LocaleString"] = request.LocaleString
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerCategory"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetCustomerCategoryRequest
//
// @return GetCustomerCategoryResponse
func (client *Client) GetCustomerCategory(request *GetCustomerCategoryRequest) (_result *GetCustomerCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomerCategoryResponse{}
	_body, _err := client.GetCustomerCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetCustomerCategoryDictionaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerCategoryDictionaryResponse
func (client *Client) GetCustomerCategoryDictionaryWithOptions(request *GetCustomerCategoryDictionaryRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerCategoryDictionaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerCategoryDictionary"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerCategoryDictionaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetCustomerCategoryDictionaryRequest
//
// @return GetCustomerCategoryDictionaryResponse
func (client *Client) GetCustomerCategoryDictionary(request *GetCustomerCategoryDictionaryRequest) (_result *GetCustomerCategoryDictionaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomerCategoryDictionaryResponse{}
	_body, _err := client.GetCustomerCategoryDictionaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetCustomerInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomerInformationResponse
func (client *Client) GetCustomerInformationWithOptions(request *GetCustomerInformationRequest, runtime *dara.RuntimeOptions) (_result *GetCustomerInformationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomerInformation"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomerInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetCustomerInformationRequest
//
// @return GetCustomerInformationResponse
func (client *Client) GetCustomerInformation(request *GetCustomerInformationRequest) (_result *GetCustomerInformationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomerInformationResponse{}
	_body, _err := client.GetCustomerInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetDingTalkUserOrgByAliyunTmpCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingTalkUserOrgByAliyunTmpCodeResponse
func (client *Client) GetDingTalkUserOrgByAliyunTmpCodeWithOptions(request *GetDingTalkUserOrgByAliyunTmpCodeRequest, runtime *dara.RuntimeOptions) (_result *GetDingTalkUserOrgByAliyunTmpCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDingTalkUserOrgByAliyunTmpCode"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDingTalkUserOrgByAliyunTmpCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDingTalkUserOrgByAliyunTmpCodeRequest
//
// @return GetDingTalkUserOrgByAliyunTmpCodeResponse
func (client *Client) GetDingTalkUserOrgByAliyunTmpCode(request *GetDingTalkUserOrgByAliyunTmpCodeRequest) (_result *GetDingTalkUserOrgByAliyunTmpCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDingTalkUserOrgByAliyunTmpCodeResponse{}
	_body, _err := client.GetDingTalkUserOrgByAliyunTmpCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetIdentityRegistrationByCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityRegistrationByCustomerResponse
func (client *Client) GetIdentityRegistrationByCustomerWithOptions(request *GetIdentityRegistrationByCustomerRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityRegistrationByCustomerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityRegistrationByCustomer"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityRegistrationByCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetIdentityRegistrationByCustomerRequest
//
// @return GetIdentityRegistrationByCustomerResponse
func (client *Client) GetIdentityRegistrationByCustomer(request *GetIdentityRegistrationByCustomerRequest) (_result *GetIdentityRegistrationByCustomerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIdentityRegistrationByCustomerResponse{}
	_body, _err := client.GetIdentityRegistrationByCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetProfileTypeByPkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProfileTypeByPkResponse
func (client *Client) GetProfileTypeByPkWithOptions(request *GetProfileTypeByPkRequest, runtime *dara.RuntimeOptions) (_result *GetProfileTypeByPkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProfileTypeByPk"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProfileTypeByPkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetProfileTypeByPkRequest
//
// @return GetProfileTypeByPkResponse
func (client *Client) GetProfileTypeByPk(request *GetProfileTypeByPkRequest) (_result *GetProfileTypeByPkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProfileTypeByPkResponse{}
	_body, _err := client.GetProfileTypeByPkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetUploadIdentityRegistrationDocConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadIdentityRegistrationDocConfigResponse
func (client *Client) GetUploadIdentityRegistrationDocConfigWithOptions(request *GetUploadIdentityRegistrationDocConfigRequest, runtime *dara.RuntimeOptions) (_result *GetUploadIdentityRegistrationDocConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUploadIdentityRegistrationDocConfig"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUploadIdentityRegistrationDocConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUploadIdentityRegistrationDocConfigRequest
//
// @return GetUploadIdentityRegistrationDocConfigResponse
func (client *Client) GetUploadIdentityRegistrationDocConfig(request *GetUploadIdentityRegistrationDocConfigRequest) (_result *GetUploadIdentityRegistrationDocConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUploadIdentityRegistrationDocConfigResponse{}
	_body, _err := client.GetUploadIdentityRegistrationDocConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # IncrBy缓存操作
//
// @param request - IncrByCacheOperateSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IncrByCacheOperateSyncResponse
func (client *Client) IncrByCacheOperateSyncWithOptions(request *IncrByCacheOperateSyncRequest, runtime *dara.RuntimeOptions) (_result *IncrByCacheOperateSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultValue) {
		query["DefaultValue"] = request.DefaultValue
	}

	if !dara.IsNil(request.ExpireSeconds) {
		query["ExpireSeconds"] = request.ExpireSeconds
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Step) {
		query["Step"] = request.Step
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IncrByCacheOperateSync"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IncrByCacheOperateSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # IncrBy缓存操作
//
// @param request - IncrByCacheOperateSyncRequest
//
// @return IncrByCacheOperateSyncResponse
func (client *Client) IncrByCacheOperateSync(request *IncrByCacheOperateSyncRequest) (_result *IncrByCacheOperateSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IncrByCacheOperateSyncResponse{}
	_body, _err := client.IncrByCacheOperateSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申请ag注销
//
// @param request - JudgeAgExistQuietPeriodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JudgeAgExistQuietPeriodResponse
func (client *Client) JudgeAgExistQuietPeriodWithOptions(request *JudgeAgExistQuietPeriodRequest, runtime *dara.RuntimeOptions) (_result *JudgeAgExistQuietPeriodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JudgeAgExistQuietPeriod"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JudgeAgExistQuietPeriodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请ag注销
//
// @param request - JudgeAgExistQuietPeriodRequest
//
// @return JudgeAgExistQuietPeriodResponse
func (client *Client) JudgeAgExistQuietPeriod(request *JudgeAgExistQuietPeriodRequest) (_result *JudgeAgExistQuietPeriodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &JudgeAgExistQuietPeriodResponse{}
	_body, _err := client.JudgeAgExistQuietPeriodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - LoadRealNameInfoByPkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoadRealNameInfoByPkResponse
func (client *Client) LoadRealNameInfoByPkWithOptions(request *LoadRealNameInfoByPkRequest, runtime *dara.RuntimeOptions) (_result *LoadRealNameInfoByPkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoadRealNameInfoByPk"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoadRealNameInfoByPkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - LoadRealNameInfoByPkRequest
//
// @return LoadRealNameInfoByPkResponse
func (client *Client) LoadRealNameInfoByPk(request *LoadRealNameInfoByPkRequest) (_result *LoadRealNameInfoByPkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LoadRealNameInfoByPkResponse{}
	_body, _err := client.LoadRealNameInfoByPkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - MapFromHavanaBindIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MapFromHavanaBindIdResponse
func (client *Client) MapFromHavanaBindIdWithOptions(tmpReq *MapFromHavanaBindIdRequest, runtime *dara.RuntimeOptions) (_result *MapFromHavanaBindIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MapFromHavanaBindIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HavanaBindStations) {
		request.HavanaBindStationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HavanaBindStations, dara.String("HavanaBindStations"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.HavanaBindId) {
		query["HavanaBindId"] = request.HavanaBindId
	}

	if !dara.IsNil(request.HavanaBindStationsShrink) {
		query["HavanaBindStations"] = request.HavanaBindStationsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MapFromHavanaBindId"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MapFromHavanaBindIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MapFromHavanaBindIdRequest
//
// @return MapFromHavanaBindIdResponse
func (client *Client) MapFromHavanaBindId(request *MapFromHavanaBindIdRequest) (_result *MapFromHavanaBindIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MapFromHavanaBindIdResponse{}
	_body, _err := client.MapFromHavanaBindIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MapPkFromHidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MapPkFromHidResponse
func (client *Client) MapPkFromHidWithOptions(request *MapPkFromHidRequest, runtime *dara.RuntimeOptions) (_result *MapPkFromHidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Hid) {
		query["Hid"] = request.Hid
	}

	if !dara.IsNil(request.MappingScenes) {
		query["MappingScenes"] = request.MappingScenes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MapPkFromHid"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MapPkFromHidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MapPkFromHidRequest
//
// @return MapPkFromHidResponse
func (client *Client) MapPkFromHid(request *MapPkFromHidRequest) (_result *MapPkFromHidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MapPkFromHidResponse{}
	_body, _err := client.MapPkFromHidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MapPkToHidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MapPkToHidResponse
func (client *Client) MapPkToHidWithOptions(request *MapPkToHidRequest, runtime *dara.RuntimeOptions) (_result *MapPkToHidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.MappingScenes) {
		query["MappingScenes"] = request.MappingScenes
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MapPkToHid"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MapPkToHidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MapPkToHidRequest
//
// @return MapPkToHidResponse
func (client *Client) MapPkToHid(request *MapPkToHidRequest) (_result *MapPkToHidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MapPkToHidResponse{}
	_body, _err := client.MapPkToHidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - MapToHavanaBindIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MapToHavanaBindIdResponse
func (client *Client) MapToHavanaBindIdWithOptions(tmpReq *MapToHavanaBindIdRequest, runtime *dara.RuntimeOptions) (_result *MapToHavanaBindIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MapToHavanaBindIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HavanaBindStations) {
		request.HavanaBindStationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HavanaBindStations, dara.String("HavanaBindStations"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.HavanaBindStationsShrink) {
		query["HavanaBindStations"] = request.HavanaBindStationsShrink
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MapToHavanaBindId"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MapToHavanaBindIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - MapToHavanaBindIdRequest
//
// @return MapToHavanaBindIdResponse
func (client *Client) MapToHavanaBindId(request *MapToHavanaBindIdRequest) (_result *MapToHavanaBindIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MapToHavanaBindIdResponse{}
	_body, _err := client.MapToHavanaBindIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyBizCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBizCategoryResponse
func (client *Client) ModifyBizCategoryWithOptions(request *ModifyBizCategoryRequest, runtime *dara.RuntimeOptions) (_result *ModifyBizCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamList) {
		query["ParamList"] = request.ParamList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBizCategory"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBizCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyBizCategoryRequest
//
// @return ModifyBizCategoryResponse
func (client *Client) ModifyBizCategory(request *ModifyBizCategoryRequest) (_result *ModifyBizCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBizCategoryResponse{}
	_body, _err := client.ModifyBizCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyContacterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyContacterResponse
func (client *Client) ModifyContacterWithOptions(request *ModifyContacterRequest, runtime *dara.RuntimeOptions) (_result *ModifyContacterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContacterAddress) {
		query["ContacterAddress"] = request.ContacterAddress
	}

	if !dara.IsNil(request.ContacterDingding) {
		query["ContacterDingding"] = request.ContacterDingding
	}

	if !dara.IsNil(request.ContacterEmail) {
		query["ContacterEmail"] = request.ContacterEmail
	}

	if !dara.IsNil(request.ContacterId) {
		query["ContacterId"] = request.ContacterId
	}

	if !dara.IsNil(request.ContacterMobile) {
		query["ContacterMobile"] = request.ContacterMobile
	}

	if !dara.IsNil(request.ContacterName) {
		query["ContacterName"] = request.ContacterName
	}

	if !dara.IsNil(request.ContacterPosition) {
		query["ContacterPosition"] = request.ContacterPosition
	}

	if !dara.IsNil(request.ContacterStaffNo) {
		query["ContacterStaffNo"] = request.ContacterStaffNo
	}

	if !dara.IsNil(request.ContacterType) {
		query["ContacterType"] = request.ContacterType
	}

	if !dara.IsNil(request.ContacterWangwang) {
		query["ContacterWangwang"] = request.ContacterWangwang
	}

	if !dara.IsNil(request.EmailConfirmed) {
		query["EmailConfirmed"] = request.EmailConfirmed
	}

	if !dara.IsNil(request.MobileConfirmed) {
		query["MobileConfirmed"] = request.MobileConfirmed
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyContacter"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyContacterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyContacterRequest
//
// @return ModifyContacterResponse
func (client *Client) ModifyContacter(request *ModifyContacterRequest) (_result *ModifyContacterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyContacterResponse{}
	_body, _err := client.ModifyContacterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyCustomerInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomerInfoResponse
func (client *Client) ModifyCustomerInfoWithOptions(request *ModifyCustomerInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomerInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Biz) {
		query["Biz"] = request.Biz
	}

	if !dara.IsNil(request.CustomerCategory) {
		query["CustomerCategory"] = request.CustomerCategory
	}

	if !dara.IsNil(request.CustomerSubCategory) {
		query["CustomerSubCategory"] = request.CustomerSubCategory
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Website) {
		query["Website"] = request.Website
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomerInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomerInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyCustomerInfoRequest
//
// @return ModifyCustomerInfoResponse
func (client *Client) ModifyCustomerInfo(request *ModifyCustomerInfoRequest) (_result *ModifyCustomerInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCustomerInfoResponse{}
	_body, _err := client.ModifyCustomerInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - OperateFinanceTaxRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateFinanceTaxResponse
func (client *Client) OperateFinanceTaxWithOptions(request *OperateFinanceTaxRequest, runtime *dara.RuntimeOptions) (_result *OperateFinanceTaxResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FinanceTax) {
		query["FinanceTax"] = request.FinanceTax
	}

	if !dara.IsNil(request.FinanceTaxCertificateImgName) {
		query["FinanceTaxCertificateImgName"] = request.FinanceTaxCertificateImgName
	}

	if !dara.IsNil(request.HId) {
		query["HId"] = request.HId
	}

	if !dara.IsNil(request.SecondFinanceTax) {
		query["SecondFinanceTax"] = request.SecondFinanceTax
	}

	if !dara.IsNil(request.SecondFinanceTaxCertificateImgName) {
		query["SecondFinanceTaxCertificateImgName"] = request.SecondFinanceTaxCertificateImgName
	}

	if !dara.IsNil(request.SecondFinanceTaxCertificateImgUrl) {
		query["SecondFinanceTaxCertificateImgUrl"] = request.SecondFinanceTaxCertificateImgUrl
	}

	if !dara.IsNil(request.FinanceTaxCertificateImgUrl) {
		query["financeTaxCertificateImgUrl"] = request.FinanceTaxCertificateImgUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateFinanceTax"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateFinanceTaxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - OperateFinanceTaxRequest
//
// @return OperateFinanceTaxResponse
func (client *Client) OperateFinanceTax(request *OperateFinanceTaxRequest) (_result *OperateFinanceTaxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateFinanceTaxResponse{}
	_body, _err := client.OperateFinanceTaxWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAccountAddressInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountAddressInfoResponse
func (client *Client) QueryAccountAddressInfoWithOptions(request *QueryAccountAddressInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountAddressInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressVersion) {
		query["AddressVersion"] = request.AddressVersion
	}

	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountAddressInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountAddressInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAccountAddressInfoRequest
//
// @return QueryAccountAddressInfoResponse
func (client *Client) QueryAccountAddressInfo(request *QueryAccountAddressInfoRequest) (_result *QueryAccountAddressInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountAddressInfoResponse{}
	_body, _err := client.QueryAccountAddressInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAccountAddressInfoWithoutHavanaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountAddressInfoWithoutHavanaResponse
func (client *Client) QueryAccountAddressInfoWithoutHavanaWithOptions(request *QueryAccountAddressInfoWithoutHavanaRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountAddressInfoWithoutHavanaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressVersion) {
		query["AddressVersion"] = request.AddressVersion
	}

	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountAddressInfoWithoutHavana"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountAddressInfoWithoutHavanaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAccountAddressInfoWithoutHavanaRequest
//
// @return QueryAccountAddressInfoWithoutHavanaResponse
func (client *Client) QueryAccountAddressInfoWithoutHavana(request *QueryAccountAddressInfoWithoutHavanaRequest) (_result *QueryAccountAddressInfoWithoutHavanaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountAddressInfoWithoutHavanaResponse{}
	_body, _err := client.QueryAccountAddressInfoWithoutHavanaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询账号收货地址
//
// @param request - QueryAccountDeliveryAddressInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountDeliveryAddressInfoResponse
func (client *Client) QueryAccountDeliveryAddressInfoWithOptions(request *QueryAccountDeliveryAddressInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountDeliveryAddressInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountDeliveryAddressInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountDeliveryAddressInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账号收货地址
//
// @param request - QueryAccountDeliveryAddressInfoRequest
//
// @return QueryAccountDeliveryAddressInfoResponse
func (client *Client) QueryAccountDeliveryAddressInfo(request *QueryAccountDeliveryAddressInfoRequest) (_result *QueryAccountDeliveryAddressInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountDeliveryAddressInfoResponse{}
	_body, _err := client.QueryAccountDeliveryAddressInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAccountProfileInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountProfileInfoResponse
func (client *Client) QueryAccountProfileInfoWithOptions(request *QueryAccountProfileInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountProfileInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountProfileInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountProfileInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAccountProfileInfoRequest
//
// @return QueryAccountProfileInfoResponse
func (client *Client) QueryAccountProfileInfo(request *QueryAccountProfileInfoRequest) (_result *QueryAccountProfileInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountProfileInfoResponse{}
	_body, _err := client.QueryAccountProfileInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAccountRealNameInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountRealNameInfoResponse
func (client *Client) QueryAccountRealNameInfoWithOptions(request *QueryAccountRealNameInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountRealNameInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountRealNameInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountRealNameInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAccountRealNameInfoRequest
//
// @return QueryAccountRealNameInfoResponse
func (client *Client) QueryAccountRealNameInfo(request *QueryAccountRealNameInfoRequest) (_result *QueryAccountRealNameInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountRealNameInfoResponse{}
	_body, _err := client.QueryAccountRealNameInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAccountSiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountSiteResponse
func (client *Client) QueryAccountSiteWithOptions(request *QueryAccountSiteRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountSiteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountSite"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountSiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAccountSiteRequest
//
// @return QueryAccountSiteResponse
func (client *Client) QueryAccountSite(request *QueryAccountSiteRequest) (_result *QueryAccountSiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountSiteResponse{}
	_body, _err := client.QueryAccountSiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAccountTrueNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAccountTrueNameResponse
func (client *Client) QueryAccountTrueNameWithOptions(request *QueryAccountTrueNameRequest, runtime *dara.RuntimeOptions) (_result *QueryAccountTrueNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAccountTrueName"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAccountTrueNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAccountTrueNameRequest
//
// @return QueryAccountTrueNameResponse
func (client *Client) QueryAccountTrueName(request *QueryAccountTrueNameRequest) (_result *QueryAccountTrueNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAccountTrueNameResponse{}
	_body, _err := client.QueryAccountTrueNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAgAccountLoginPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAgAccountLoginPermissionResponse
func (client *Client) QueryAgAccountLoginPermissionWithOptions(request *QueryAgAccountLoginPermissionRequest, runtime *dara.RuntimeOptions) (_result *QueryAgAccountLoginPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAgAccountLoginPermission"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAgAccountLoginPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAgAccountLoginPermissionRequest
//
// @return QueryAgAccountLoginPermissionResponse
func (client *Client) QueryAgAccountLoginPermission(request *QueryAgAccountLoginPermissionRequest) (_result *QueryAgAccountLoginPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAgAccountLoginPermissionResponse{}
	_body, _err := client.QueryAgAccountLoginPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAgRelationCountAndQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAgRelationCountAndQuotaResponse
func (client *Client) QueryAgRelationCountAndQuotaWithOptions(request *QueryAgRelationCountAndQuotaRequest, runtime *dara.RuntimeOptions) (_result *QueryAgRelationCountAndQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.CallerBid) {
		body["CallerBid"] = request.CallerBid
	}

	if !dara.IsNil(request.CallerParentId) {
		body["CallerParentId"] = request.CallerParentId
	}

	if !dara.IsNil(request.CallerType) {
		body["CallerType"] = request.CallerType
	}

	if !dara.IsNil(request.CallerUid) {
		body["CallerUid"] = request.CallerUid
	}

	if !dara.IsNil(request.Mpk) {
		body["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.NullObject) {
		body["NullObject"] = request.NullObject
	}

	if !dara.IsNil(request.RequestId) {
		body["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SecurityToken) {
		body["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SourceIp) {
		body["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StsTokenCallerBid) {
		body["StsTokenCallerBid"] = request.StsTokenCallerBid
	}

	if !dara.IsNil(request.StsTokenCallerUid) {
		body["StsTokenCallerUid"] = request.StsTokenCallerUid
	}

	if !dara.IsNil(request.StsTokenRoleId) {
		body["StsTokenRoleId"] = request.StsTokenRoleId
	}

	if !dara.IsNil(request.Version) {
		body["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAgRelationCountAndQuota"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAgRelationCountAndQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAgRelationCountAndQuotaRequest
//
// @return QueryAgRelationCountAndQuotaResponse
func (client *Client) QueryAgRelationCountAndQuota(request *QueryAgRelationCountAndQuotaRequest) (_result *QueryAgRelationCountAndQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAgRelationCountAndQuotaResponse{}
	_body, _err := client.QueryAgRelationCountAndQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAgSecurityMobileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAgSecurityMobileResponse
func (client *Client) QueryAgSecurityMobileWithOptions(request *QueryAgSecurityMobileRequest, runtime *dara.RuntimeOptions) (_result *QueryAgSecurityMobileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAgSecurityMobile"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAgSecurityMobileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAgSecurityMobileRequest
//
// @return QueryAgSecurityMobileResponse
func (client *Client) QueryAgSecurityMobile(request *QueryAgSecurityMobileRequest) (_result *QueryAgSecurityMobileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAgSecurityMobileResponse{}
	_body, _err := client.QueryAgSecurityMobileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryBindsByOuterIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBindsByOuterIdResponse
func (client *Client) QueryBindsByOuterIdWithOptions(request *QueryBindsByOuterIdRequest, runtime *dara.RuntimeOptions) (_result *QueryBindsByOuterIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.MinorOuterId) {
		query["MinorOuterId"] = request.MinorOuterId
	}

	if !dara.IsNil(request.OuterId) {
		query["OuterId"] = request.OuterId
	}

	if !dara.IsNil(request.TenantId) {
		query["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBindsByOuterId"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBindsByOuterIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBindsByOuterIdRequest
//
// @return QueryBindsByOuterIdResponse
func (client *Client) QueryBindsByOuterId(request *QueryBindsByOuterIdRequest) (_result *QueryBindsByOuterIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBindsByOuterIdResponse{}
	_body, _err := client.QueryBindsByOuterIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - QueryBindsByPkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBindsByPkResponse
func (client *Client) QueryBindsByPkWithOptions(tmpReq *QueryBindsByPkRequest, runtime *dara.RuntimeOptions) (_result *QueryBindsByPkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryBindsByPkShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TenantIds) {
		request.TenantIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantIds, dara.String("TenantIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.TenantIdsShrink) {
		query["TenantIds"] = request.TenantIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBindsByPk"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBindsByPkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBindsByPkRequest
//
// @return QueryBindsByPkResponse
func (client *Client) QueryBindsByPk(request *QueryBindsByPkRequest) (_result *QueryBindsByPkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBindsByPkResponse{}
	_body, _err := client.QueryBindsByPkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryCustomerLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCustomerLabelResponse
func (client *Client) QueryCustomerLabelWithOptions(request *QueryCustomerLabelRequest, runtime *dara.RuntimeOptions) (_result *QueryCustomerLabelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelSeries) {
		query["LabelSeries"] = request.LabelSeries
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCustomerLabel"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCustomerLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCustomerLabelRequest
//
// @return QueryCustomerLabelResponse
func (client *Client) QueryCustomerLabel(request *QueryCustomerLabelRequest) (_result *QueryCustomerLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCustomerLabelResponse{}
	_body, _err := client.QueryCustomerLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryDeleteTaskCheckDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeleteTaskCheckDataResponse
func (client *Client) QueryDeleteTaskCheckDataWithOptions(request *QueryDeleteTaskCheckDataRequest, runtime *dara.RuntimeOptions) (_result *QueryDeleteTaskCheckDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.LongLang) {
		query["LongLang"] = request.LongLang
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDeleteTaskCheckData"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDeleteTaskCheckDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDeleteTaskCheckDataRequest
//
// @return QueryDeleteTaskCheckDataResponse
func (client *Client) QueryDeleteTaskCheckData(request *QueryDeleteTaskCheckDataRequest) (_result *QueryDeleteTaskCheckDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDeleteTaskCheckDataResponse{}
	_body, _err := client.QueryDeleteTaskCheckDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryEncryptedAccountProfileInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEncryptedAccountProfileInfoResponse
func (client *Client) QueryEncryptedAccountProfileInfoWithOptions(request *QueryEncryptedAccountProfileInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryEncryptedAccountProfileInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEncryptedAccountProfileInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEncryptedAccountProfileInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryEncryptedAccountProfileInfoRequest
//
// @return QueryEncryptedAccountProfileInfoResponse
func (client *Client) QueryEncryptedAccountProfileInfo(request *QueryEncryptedAccountProfileInfoRequest) (_result *QueryEncryptedAccountProfileInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEncryptedAccountProfileInfoResponse{}
	_body, _err := client.QueryEncryptedAccountProfileInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryEnterpriseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEnterpriseInfoResponse
func (client *Client) QueryEnterpriseInfoWithOptions(request *QueryEnterpriseInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryEnterpriseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseVersion) {
		query["EnterpriseVersion"] = request.EnterpriseVersion
	}

	if !dara.IsNil(request.HavanaId) {
		query["HavanaId"] = request.HavanaId
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEnterpriseInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEnterpriseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryEnterpriseInfoRequest
//
// @return QueryEnterpriseInfoResponse
func (client *Client) QueryEnterpriseInfo(request *QueryEnterpriseInfoRequest) (_result *QueryEnterpriseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEnterpriseInfoResponse{}
	_body, _err := client.QueryEnterpriseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryEnumConfigByTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEnumConfigByTypeResponse
func (client *Client) QueryEnumConfigByTypeWithOptions(request *QueryEnumConfigByTypeRequest, runtime *dara.RuntimeOptions) (_result *QueryEnumConfigByTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEnumConfigByType"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEnumConfigByTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryEnumConfigByTypeRequest
//
// @return QueryEnumConfigByTypeResponse
func (client *Client) QueryEnumConfigByType(request *QueryEnumConfigByTypeRequest) (_result *QueryEnumConfigByTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryEnumConfigByTypeResponse{}
	_body, _err := client.QueryEnumConfigByTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryOneKeyDeleteBlockListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOneKeyDeleteBlockListResponse
func (client *Client) QueryOneKeyDeleteBlockListWithOptions(request *QueryOneKeyDeleteBlockListRequest, runtime *dara.RuntimeOptions) (_result *QueryOneKeyDeleteBlockListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryOneKeyDeleteBlockList"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryOneKeyDeleteBlockListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryOneKeyDeleteBlockListRequest
//
// @return QueryOneKeyDeleteBlockListResponse
func (client *Client) QueryOneKeyDeleteBlockList(request *QueryOneKeyDeleteBlockListRequest) (_result *QueryOneKeyDeleteBlockListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryOneKeyDeleteBlockListResponse{}
	_body, _err := client.QueryOneKeyDeleteBlockListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySecurityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySecurityInfoResponse
func (client *Client) QuerySecurityInfoWithOptions(request *QuerySecurityInfoRequest, runtime *dara.RuntimeOptions) (_result *QuerySecurityInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySecurityInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySecurityInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySecurityInfoRequest
//
// @return QuerySecurityInfoResponse
func (client *Client) QuerySecurityInfo(request *QuerySecurityInfoRequest) (_result *QuerySecurityInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySecurityInfoResponse{}
	_body, _err := client.QuerySecurityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RegisterInternalAccountForBucRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterInternalAccountForBucResponse
func (client *Client) RegisterInternalAccountForBucWithOptions(request *RegisterInternalAccountForBucRequest, runtime *dara.RuntimeOptions) (_result *RegisterInternalAccountForBucResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.IsEmailConfirmed) {
		query["IsEmailConfirmed"] = request.IsEmailConfirmed
	}

	if !dara.IsNil(request.IsMobileConfirmed) {
		query["IsMobileConfirmed"] = request.IsMobileConfirmed
	}

	if !dara.IsNil(request.IsMobileLogin) {
		query["IsMobileLogin"] = request.IsMobileLogin
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.NationalityCode) {
		query["NationalityCode"] = request.NationalityCode
	}

	if !dara.IsNil(request.PlainPassword) {
		query["PlainPassword"] = request.PlainPassword
	}

	if !dara.IsNil(request.PreferredLanguage) {
		query["PreferredLanguage"] = request.PreferredLanguage
	}

	if !dara.IsNil(request.AccountTypeCode) {
		query["accountTypeCode"] = request.AccountTypeCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterInternalAccountForBuc"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterInternalAccountForBucResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RegisterInternalAccountForBucRequest
//
// @return RegisterInternalAccountForBucResponse
func (client *Client) RegisterInternalAccountForBuc(request *RegisterInternalAccountForBucRequest) (_result *RegisterInternalAccountForBucResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterInternalAccountForBucResponse{}
	_body, _err := client.RegisterInternalAccountForBucWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReleaseAgAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseAgAccountResponse
func (client *Client) ReleaseAgAccountWithOptions(request *ReleaseAgAccountRequest, runtime *dara.RuntimeOptions) (_result *ReleaseAgAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.ReleaseReason) {
		query["ReleaseReason"] = request.ReleaseReason
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseAgAccount"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseAgAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReleaseAgAccountRequest
//
// @return ReleaseAgAccountResponse
func (client *Client) ReleaseAgAccount(request *ReleaseAgAccountRequest) (_result *ReleaseAgAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseAgAccountResponse{}
	_body, _err := client.ReleaseAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResendAsyncCreateAgAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendAsyncCreateAgAccountResponse
func (client *Client) ResendAsyncCreateAgAccountWithOptions(request *ResendAsyncCreateAgAccountRequest, runtime *dara.RuntimeOptions) (_result *ResendAsyncCreateAgAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.TraceNo) {
		query["TraceNo"] = request.TraceNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResendAsyncCreateAgAccount"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResendAsyncCreateAgAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResendAsyncCreateAgAccountRequest
//
// @return ResendAsyncCreateAgAccountResponse
func (client *Client) ResendAsyncCreateAgAccount(request *ResendAsyncCreateAgAccountRequest) (_result *ResendAsyncCreateAgAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResendAsyncCreateAgAccountResponse{}
	_body, _err := client.ResendAsyncCreateAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResendAsyncModifyLoginEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResendAsyncModifyLoginEmailResponse
func (client *Client) ResendAsyncModifyLoginEmailWithOptions(request *ResendAsyncModifyLoginEmailRequest, runtime *dara.RuntimeOptions) (_result *ResendAsyncModifyLoginEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.TraceNo) {
		query["TraceNo"] = request.TraceNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResendAsyncModifyLoginEmail"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResendAsyncModifyLoginEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResendAsyncModifyLoginEmailRequest
//
// @return ResendAsyncModifyLoginEmailResponse
func (client *Client) ResendAsyncModifyLoginEmail(request *ResendAsyncModifyLoginEmailRequest) (_result *ResendAsyncModifyLoginEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResendAsyncModifyLoginEmailResponse{}
	_body, _err := client.ResendAsyncModifyLoginEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SeparateAgRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SeparateAgRelationResponse
func (client *Client) SeparateAgRelationWithOptions(request *SeparateAgRelationRequest, runtime *dara.RuntimeOptions) (_result *SeparateAgRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SeparateAgRelation"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SeparateAgRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SeparateAgRelationRequest
//
// @return SeparateAgRelationResponse
func (client *Client) SeparateAgRelation(request *SeparateAgRelationRequest) (_result *SeparateAgRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SeparateAgRelationResponse{}
	_body, _err := client.SeparateAgRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步set操作
//
// @param request - SetCacheOperateSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCacheOperateSyncResponse
func (client *Client) SetCacheOperateSyncWithOptions(request *SetCacheOperateSyncRequest, runtime *dara.RuntimeOptions) (_result *SetCacheOperateSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExceptVersion) {
		query["ExceptVersion"] = request.ExceptVersion
	}

	if !dara.IsNil(request.ExpireSeconds) {
		query["ExpireSeconds"] = request.ExpireSeconds
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.SetType) {
		query["SetType"] = request.SetType
	}

	if !dara.IsNil(request.ValueClazz) {
		query["ValueClazz"] = request.ValueClazz
	}

	if !dara.IsNil(request.ValueString) {
		query["ValueString"] = request.ValueString
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCacheOperateSync"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCacheOperateSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步set操作
//
// @param request - SetCacheOperateSyncRequest
//
// @return SetCacheOperateSyncResponse
func (client *Client) SetCacheOperateSync(request *SetCacheOperateSyncRequest) (_result *SetCacheOperateSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCacheOperateSyncResponse{}
	_body, _err := client.SetCacheOperateSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateAccountAddressInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountAddressInfoResponse
func (client *Client) UpdateAccountAddressInfoWithOptions(tmpReq *UpdateAccountAddressInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccountAddressInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAccountAddressInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CityJsonString) {
		request.CityJsonStringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CityJsonString, dara.String("CityJsonString"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DistrictJsonString) {
		request.DistrictJsonStringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DistrictJsonString, dara.String("DistrictJsonString"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProvinceJsonString) {
		request.ProvinceJsonStringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProvinceJsonString, dara.String("ProvinceJsonString"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.Address2) {
		query["Address2"] = request.Address2
	}

	if !dara.IsNil(request.CityJsonStringShrink) {
		query["CityJsonString"] = request.CityJsonStringShrink
	}

	if !dara.IsNil(request.DistrictJsonStringShrink) {
		query["DistrictJsonString"] = request.DistrictJsonStringShrink
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.PostCode) {
		query["PostCode"] = request.PostCode
	}

	if !dara.IsNil(request.ProvinceJsonStringShrink) {
		query["ProvinceJsonString"] = request.ProvinceJsonStringShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccountAddressInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccountAddressInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAccountAddressInfoRequest
//
// @return UpdateAccountAddressInfoResponse
func (client *Client) UpdateAccountAddressInfo(request *UpdateAccountAddressInfoRequest) (_result *UpdateAccountAddressInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAccountAddressInfoResponse{}
	_body, _err := client.UpdateAccountAddressInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateAccountProfileInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountProfileInfoResponse
func (client *Client) UpdateAccountProfileInfoWithOptions(tmpReq *UpdateAccountProfileInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccountProfileInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAccountProfileInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CityJsonString) {
		request.CityJsonStringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CityJsonString, dara.String("CityJsonString"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DistrictJsonString) {
		request.DistrictJsonStringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DistrictJsonString, dara.String("DistrictJsonString"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProvinceJsonString) {
		request.ProvinceJsonStringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProvinceJsonString, dara.String("ProvinceJsonString"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountAttribute) {
		query["AccountAttribute"] = request.AccountAttribute
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.Address2) {
		query["Address2"] = request.Address2
	}

	if !dara.IsNil(request.BindAlipayNo) {
		query["BindAlipayNo"] = request.BindAlipayNo
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.CityJsonStringShrink) {
		query["CityJsonString"] = request.CityJsonStringShrink
	}

	if !dara.IsNil(request.ContactMethod) {
		query["ContactMethod"] = request.ContactMethod
	}

	if !dara.IsNil(request.DistrictJsonStringShrink) {
		query["DistrictJsonString"] = request.DistrictJsonStringShrink
	}

	if !dara.IsNil(request.Fax) {
		query["Fax"] = request.Fax
	}

	if !dara.IsNil(request.FirstName) {
		query["FirstName"] = request.FirstName
	}

	if !dara.IsNil(request.Head) {
		query["Head"] = request.Head
	}

	if !dara.IsNil(request.HeadColor) {
		query["HeadColor"] = request.HeadColor
	}

	if !dara.IsNil(request.LastName) {
		query["LastName"] = request.LastName
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.PostCode) {
		query["PostCode"] = request.PostCode
	}

	if !dara.IsNil(request.ProvinceJsonStringShrink) {
		query["ProvinceJsonString"] = request.ProvinceJsonStringShrink
	}

	if !dara.IsNil(request.SelfServicingBusinessRegNum) {
		query["SelfServicingBusinessRegNum"] = request.SelfServicingBusinessRegNum
	}

	if !dara.IsNil(request.SelfServicingIdentificationNum) {
		query["SelfServicingIdentificationNum"] = request.SelfServicingIdentificationNum
	}

	if !dara.IsNil(request.TrueName) {
		query["TrueName"] = request.TrueName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccountProfileInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccountProfileInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAccountProfileInfoRequest
//
// @return UpdateAccountProfileInfoResponse
func (client *Client) UpdateAccountProfileInfo(request *UpdateAccountProfileInfoRequest) (_result *UpdateAccountProfileInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAccountProfileInfoResponse{}
	_body, _err := client.UpdateAccountProfileInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateAgAccountAddressInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgAccountAddressInfoResponse
func (client *Client) UpdateAgAccountAddressInfoWithOptions(request *UpdateAgAccountAddressInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgAccountAddressInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.Address2) {
		query["Address2"] = request.Address2
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.PostCode) {
		query["PostCode"] = request.PostCode
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgAccountAddressInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgAccountAddressInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAgAccountAddressInfoRequest
//
// @return UpdateAgAccountAddressInfoResponse
func (client *Client) UpdateAgAccountAddressInfo(request *UpdateAgAccountAddressInfoRequest) (_result *UpdateAgAccountAddressInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAgAccountAddressInfoResponse{}
	_body, _err := client.UpdateAgAccountAddressInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateAgServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgServiceStatusResponse
func (client *Client) UpdateAgServiceStatusWithOptions(request *UpdateAgServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgAccountType) {
		query["AgAccountType"] = request.AgAccountType
	}

	if !dara.IsNil(request.Mpk) {
		query["Mpk"] = request.Mpk
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgServiceStatus"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAgServiceStatusRequest
//
// @return UpdateAgServiceStatusResponse
func (client *Client) UpdateAgServiceStatus(request *UpdateAgServiceStatusRequest) (_result *UpdateAgServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAgServiceStatusResponse{}
	_body, _err := client.UpdateAgServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCustomerCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomerCategoryResponse
func (client *Client) UpdateCustomerCategoryWithOptions(request *UpdateCustomerCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomerCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParamList) {
		query["ParamList"] = request.ParamList
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomerCategory"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomerCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCustomerCategoryRequest
//
// @return UpdateCustomerCategoryResponse
func (client *Client) UpdateCustomerCategory(request *UpdateCustomerCategoryRequest) (_result *UpdateCustomerCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomerCategoryResponse{}
	_body, _err := client.UpdateCustomerCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCustomerInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomerInformationResponse
func (client *Client) UpdateCustomerInformationWithOptions(request *UpdateCustomerInformationRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomerInformationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Biz) {
		query["Biz"] = request.Biz
	}

	if !dara.IsNil(request.CustomerCategory) {
		query["CustomerCategory"] = request.CustomerCategory
	}

	if !dara.IsNil(request.CustomerSubCategory) {
		query["CustomerSubCategory"] = request.CustomerSubCategory
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Website) {
		query["Website"] = request.Website
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomerInformation"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomerInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCustomerInformationRequest
//
// @return UpdateCustomerInformationResponse
func (client *Client) UpdateCustomerInformation(request *UpdateCustomerInformationRequest) (_result *UpdateCustomerInformationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomerInformationResponse{}
	_body, _err := client.UpdateCustomerInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateOrInsertEnterpriseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrInsertEnterpriseInfoResponse
func (client *Client) UpdateOrInsertEnterpriseInfoWithOptions(tmpReq *UpdateOrInsertEnterpriseInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateOrInsertEnterpriseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateOrInsertEnterpriseInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CityJsonString) {
		request.CityJsonStringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CityJsonString, dara.String("CityJsonString"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProvinceJsonString) {
		request.ProvinceJsonStringShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProvinceJsonString, dara.String("ProvinceJsonString"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.Alias) {
		query["Alias"] = request.Alias
	}

	if !dara.IsNil(request.CityJsonStringShrink) {
		query["CityJsonString"] = request.CityJsonStringShrink
	}

	if !dara.IsNil(request.EnterpriseSize) {
		query["EnterpriseSize"] = request.EnterpriseSize
	}

	if !dara.IsNil(request.Fax) {
		query["Fax"] = request.Fax
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PK) {
		query["PK"] = request.PK
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.ProvinceJsonStringShrink) {
		query["ProvinceJsonString"] = request.ProvinceJsonStringShrink
	}

	if !dara.IsNil(request.Years) {
		query["Years"] = request.Years
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOrInsertEnterpriseInfo"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOrInsertEnterpriseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateOrInsertEnterpriseInfoRequest
//
// @return UpdateOrInsertEnterpriseInfoResponse
func (client *Client) UpdateOrInsertEnterpriseInfo(request *UpdateOrInsertEnterpriseInfoRequest) (_result *UpdateOrInsertEnterpriseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOrInsertEnterpriseInfoResponse{}
	_body, _err := client.UpdateOrInsertEnterpriseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DoLogicalDeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DoLogicalDeleteResourceResponse
func (client *Client) DoLogicalDeleteResourceWithOptions(request *DoLogicalDeleteResourceRequest, runtime *dara.RuntimeOptions) (_result *DoLogicalDeleteResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.GmtWakeup) {
		query["GmtWakeup"] = request.GmtWakeup
	}

	if !dara.IsNil(request.Hid) {
		query["Hid"] = request.Hid
	}

	if !dara.IsNil(request.Interrupt) {
		query["Interrupt"] = request.Interrupt
	}

	if !dara.IsNil(request.Invoker) {
		query["Invoker"] = request.Invoker
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.Success) {
		query["Success"] = request.Success
	}

	if !dara.IsNil(request.TaskExtraData) {
		query["TaskExtraData"] = request.TaskExtraData
	}

	if !dara.IsNil(request.TaskIdentifier) {
		query["TaskIdentifier"] = request.TaskIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("doLogicalDeleteResource"),
		Version:     dara.String("2016-06-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DoLogicalDeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DoLogicalDeleteResourceRequest
//
// @return DoLogicalDeleteResourceResponse
func (client *Client) DoLogicalDeleteResource(request *DoLogicalDeleteResourceRequest) (_result *DoLogicalDeleteResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DoLogicalDeleteResourceResponse{}
	_body, _err := client.DoLogicalDeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
