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
		"ap-southeast-1": dara.String("accountcenter-intl.aliyuncs.com"),
		"cn-hangzhou":    dara.String("accountcenter.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("accountcenter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates an account contact.
//
// Description:
//
// Creates an account contact.
//
// @param request - AccountContactAddRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactAddResponse
func (client *Client) AccountContactAddWithOptions(request *AccountContactAddRequest, runtime *dara.RuntimeOptions) (_result *AccountContactAddResponse, _err error) {
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

	if !dara.IsNil(request.AsyncEmailVerify) {
		body["AsyncEmailVerify"] = request.AsyncEmailVerify
	}

	if !dara.IsNil(request.AsyncMobileVerify) {
		body["AsyncMobileVerify"] = request.AsyncMobileVerify
	}

	if !dara.IsNil(request.ContactEmail) {
		body["ContactEmail"] = request.ContactEmail
	}

	if !dara.IsNil(request.ContactMobile) {
		body["ContactMobile"] = request.ContactMobile
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.ContactPosition) {
		body["ContactPosition"] = request.ContactPosition
	}

	if !dara.IsNil(request.EmailCode) {
		body["EmailCode"] = request.EmailCode
	}

	if !dara.IsNil(request.MobileCode) {
		body["MobileCode"] = request.MobileCode
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.SharedContact) {
		body["SharedContact"] = request.SharedContact
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccountContactAdd"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AccountContactAddResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account contact.
//
// Description:
//
// Creates an account contact.
//
// @param request - AccountContactAddRequest
//
// @return AccountContactAddResponse
func (client *Client) AccountContactAdd(request *AccountContactAddRequest) (_result *AccountContactAddResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AccountContactAddResponse{}
	_body, _err := client.AccountContactAddWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a private contact.
//
// @param request - AccountContactDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactDeleteResponse
func (client *Client) AccountContactDeleteWithOptions(request *AccountContactDeleteRequest, runtime *dara.RuntimeOptions) (_result *AccountContactDeleteResponse, _err error) {
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

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccountContactDelete"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AccountContactDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a private contact.
//
// @param request - AccountContactDeleteRequest
//
// @return AccountContactDeleteResponse
func (client *Client) AccountContactDelete(request *AccountContactDeleteRequest) (_result *AccountContactDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AccountContactDeleteResponse{}
	_body, _err := client.AccountContactDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify a private contact.
//
// @param request - AccountContactEditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactEditResponse
func (client *Client) AccountContactEditWithOptions(request *AccountContactEditRequest, runtime *dara.RuntimeOptions) (_result *AccountContactEditResponse, _err error) {
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

	if !dara.IsNil(request.AsyncEmailVerify) {
		body["AsyncEmailVerify"] = request.AsyncEmailVerify
	}

	if !dara.IsNil(request.AsyncMobileVerify) {
		body["AsyncMobileVerify"] = request.AsyncMobileVerify
	}

	if !dara.IsNil(request.ContactEmail) {
		body["ContactEmail"] = request.ContactEmail
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactMobile) {
		body["ContactMobile"] = request.ContactMobile
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.ContactPosition) {
		body["ContactPosition"] = request.ContactPosition
	}

	if !dara.IsNil(request.EmailCode) {
		body["EmailCode"] = request.EmailCode
	}

	if !dara.IsNil(request.MobileCode) {
		body["MobileCode"] = request.MobileCode
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.SharedContact) {
		body["SharedContact"] = request.SharedContact
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccountContactEdit"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AccountContactEditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify a private contact.
//
// @param request - AccountContactEditRequest
//
// @return AccountContactEditResponse
func (client *Client) AccountContactEdit(request *AccountContactEditRequest) (_result *AccountContactEditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AccountContactEditResponse{}
	_body, _err := client.AccountContactEditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a contact.
//
// @param request - AccountContactQueryDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactQueryDetailResponse
func (client *Client) AccountContactQueryDetailWithOptions(request *AccountContactQueryDetailRequest, runtime *dara.RuntimeOptions) (_result *AccountContactQueryDetailResponse, _err error) {
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

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccountContactQueryDetail"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AccountContactQueryDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a contact.
//
// @param request - AccountContactQueryDetailRequest
//
// @return AccountContactQueryDetailResponse
func (client *Client) AccountContactQueryDetail(request *AccountContactQueryDetailRequest) (_result *AccountContactQueryDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AccountContactQueryDetailResponse{}
	_body, _err := client.AccountContactQueryDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the contact list.
//
// @param request - AccountContactQueryPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactQueryPageListResponse
func (client *Client) AccountContactQueryPageListWithOptions(request *AccountContactQueryPageListRequest, runtime *dara.RuntimeOptions) (_result *AccountContactQueryPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ShowCompleteInfo) {
		query["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivateContact) {
		body["PrivateContact"] = request.PrivateContact
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.SharedContact) {
		body["SharedContact"] = request.SharedContact
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccountContactQueryPageList"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AccountContactQueryPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the contact list.
//
// @param request - AccountContactQueryPageListRequest
//
// @return AccountContactQueryPageListResponse
func (client *Client) AccountContactQueryPageList(request *AccountContactQueryPageListRequest) (_result *AccountContactQueryPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AccountContactQueryPageListResponse{}
	_body, _err := client.AccountContactQueryPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改登录密码
//
// @param request - EnterpriseAccountChangeLoginPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeLoginPasswordResponse
func (client *Client) EnterpriseAccountChangeLoginPasswordWithOptions(request *EnterpriseAccountChangeLoginPasswordRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountChangeLoginPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptPassword) {
		query["EncryptPassword"] = request.EncryptPassword
	}

	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountChangeLoginPassword"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountChangeLoginPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改登录密码
//
// @param request - EnterpriseAccountChangeLoginPasswordRequest
//
// @return EnterpriseAccountChangeLoginPasswordResponse
func (client *Client) EnterpriseAccountChangeLoginPassword(request *EnterpriseAccountChangeLoginPasswordRequest) (_result *EnterpriseAccountChangeLoginPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountChangeLoginPasswordResponse{}
	_body, _err := client.EnterpriseAccountChangeLoginPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改安全邮箱
//
// @param request - EnterpriseAccountChangeSecurityEmailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeSecurityEmailResponse
func (client *Client) EnterpriseAccountChangeSecurityEmailWithOptions(request *EnterpriseAccountChangeSecurityEmailRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountChangeSecurityEmailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SecurityEmail) {
		query["SecurityEmail"] = request.SecurityEmail
	}

	if !dara.IsNil(request.VerifyCode) {
		query["VerifyCode"] = request.VerifyCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountChangeSecurityEmail"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountChangeSecurityEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改安全邮箱
//
// @param request - EnterpriseAccountChangeSecurityEmailRequest
//
// @return EnterpriseAccountChangeSecurityEmailResponse
func (client *Client) EnterpriseAccountChangeSecurityEmail(request *EnterpriseAccountChangeSecurityEmailRequest) (_result *EnterpriseAccountChangeSecurityEmailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountChangeSecurityEmailResponse{}
	_body, _err := client.EnterpriseAccountChangeSecurityEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改成员账号安全手机号
//
// @param request - EnterpriseAccountChangeSecurityMobileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeSecurityMobileResponse
func (client *Client) EnterpriseAccountChangeSecurityMobileWithOptions(request *EnterpriseAccountChangeSecurityMobileRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountChangeSecurityMobileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.NewMobile) {
		query["NewMobile"] = request.NewMobile
	}

	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.VerificationCode) {
		query["VerificationCode"] = request.VerificationCode
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountChangeSecurityMobile"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountChangeSecurityMobileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改成员账号安全手机号
//
// @param request - EnterpriseAccountChangeSecurityMobileRequest
//
// @return EnterpriseAccountChangeSecurityMobileResponse
func (client *Client) EnterpriseAccountChangeSecurityMobile(request *EnterpriseAccountChangeSecurityMobileRequest) (_result *EnterpriseAccountChangeSecurityMobileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountChangeSecurityMobileResponse{}
	_body, _err := client.EnterpriseAccountChangeSecurityMobileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询纳管账号授权角色
//
// @param request - EnterpriseAccountQueryAccountGrantedRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryAccountGrantedRolesResponse
func (client *Client) EnterpriseAccountQueryAccountGrantedRolesWithOptions(request *EnterpriseAccountQueryAccountGrantedRolesRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountQueryAccountGrantedRolesResponse, _err error) {
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

	if !dara.IsNil(request.IsOpenApi) {
		body["IsOpenApi"] = request.IsOpenApi
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.Pk) {
		body["Pk"] = request.Pk
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountQueryAccountGrantedRoles"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountQueryAccountGrantedRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询纳管账号授权角色
//
// @param request - EnterpriseAccountQueryAccountGrantedRolesRequest
//
// @return EnterpriseAccountQueryAccountGrantedRolesResponse
func (client *Client) EnterpriseAccountQueryAccountGrantedRoles(request *EnterpriseAccountQueryAccountGrantedRolesRequest) (_result *EnterpriseAccountQueryAccountGrantedRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountQueryAccountGrantedRolesResponse{}
	_body, _err := client.EnterpriseAccountQueryAccountGrantedRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询纳管账号信息
//
// @param request - EnterpriseAccountQueryAccountsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryAccountsInfoResponse
func (client *Client) EnterpriseAccountQueryAccountsInfoWithOptions(request *EnterpriseAccountQueryAccountsInfoRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountQueryAccountsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PksJson) {
		query["PksJson"] = request.PksJson
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountQueryAccountsInfo"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountQueryAccountsInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询纳管账号信息
//
// @param request - EnterpriseAccountQueryAccountsInfoRequest
//
// @return EnterpriseAccountQueryAccountsInfoResponse
func (client *Client) EnterpriseAccountQueryAccountsInfo(request *EnterpriseAccountQueryAccountsInfoRequest) (_result *EnterpriseAccountQueryAccountsInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountQueryAccountsInfoResponse{}
	_body, _err := client.EnterpriseAccountQueryAccountsInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询纳管账号登录设置
//
// @param request - EnterpriseAccountQueryLoginSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryLoginSettingsResponse
func (client *Client) EnterpriseAccountQueryLoginSettingsWithOptions(request *EnterpriseAccountQueryLoginSettingsRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountQueryLoginSettingsResponse, _err error) {
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

	if !dara.IsNil(request.IsOpenApi) {
		body["IsOpenApi"] = request.IsOpenApi
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.Pk) {
		body["Pk"] = request.Pk
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountQueryLoginSettings"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountQueryLoginSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询纳管账号登录设置
//
// @param request - EnterpriseAccountQueryLoginSettingsRequest
//
// @return EnterpriseAccountQueryLoginSettingsResponse
func (client *Client) EnterpriseAccountQueryLoginSettings(request *EnterpriseAccountQueryLoginSettingsRequest) (_result *EnterpriseAccountQueryLoginSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountQueryLoginSettingsResponse{}
	_body, _err := client.EnterpriseAccountQueryLoginSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除mfa
//
// @param request - EnterpriseAccountRemoveMfaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountRemoveMfaResponse
func (client *Client) EnterpriseAccountRemoveMfaWithOptions(request *EnterpriseAccountRemoveMfaRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountRemoveMfaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountRemoveMfa"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountRemoveMfaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除mfa
//
// @param request - EnterpriseAccountRemoveMfaRequest
//
// @return EnterpriseAccountRemoveMfaResponse
func (client *Client) EnterpriseAccountRemoveMfa(request *EnterpriseAccountRemoveMfaRequest) (_result *EnterpriseAccountRemoveMfaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountRemoveMfaResponse{}
	_body, _err := client.EnterpriseAccountRemoveMfaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 脱离ea
//
// @param request - EnterpriseAccountSeparateEaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountSeparateEaResponse
func (client *Client) EnterpriseAccountSeparateEaWithOptions(request *EnterpriseAccountSeparateEaRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountSeparateEaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountSeparateEa"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountSeparateEaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 脱离ea
//
// @param request - EnterpriseAccountSeparateEaRequest
//
// @return EnterpriseAccountSeparateEaResponse
func (client *Client) EnterpriseAccountSeparateEa(request *EnterpriseAccountSeparateEaRequest) (_result *EnterpriseAccountSeparateEaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountSeparateEaResponse{}
	_body, _err := client.EnterpriseAccountSeparateEaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新账号企业多账号中的别名
//
// @param request - EnterpriseAccountUpdateAccountAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateAccountAliasResponse
func (client *Client) EnterpriseAccountUpdateAccountAliasWithOptions(request *EnterpriseAccountUpdateAccountAliasRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateAccountAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		query["Alias"] = request.Alias
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountUpdateAccountAlias"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountUpdateAccountAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新账号企业多账号中的别名
//
// @param request - EnterpriseAccountUpdateAccountAliasRequest
//
// @return EnterpriseAccountUpdateAccountAliasResponse
func (client *Client) EnterpriseAccountUpdateAccountAlias(request *EnterpriseAccountUpdateAccountAliasRequest) (_result *EnterpriseAccountUpdateAccountAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateAccountAliasResponse{}
	_body, _err := client.EnterpriseAccountUpdateAccountAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新账号授权
//
// @param request - EnterpriseAccountUpdateAccountBizRoleGrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateAccountBizRoleGrantResponse
func (client *Client) EnterpriseAccountUpdateAccountBizRoleGrantWithOptions(request *EnterpriseAccountUpdateAccountBizRoleGrantRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateAccountBizRoleGrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRoleCodeListJson) {
		query["BizRoleCodeListJson"] = request.BizRoleCodeListJson
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountUpdateAccountBizRoleGrant"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountUpdateAccountBizRoleGrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新账号授权
//
// @param request - EnterpriseAccountUpdateAccountBizRoleGrantRequest
//
// @return EnterpriseAccountUpdateAccountBizRoleGrantResponse
func (client *Client) EnterpriseAccountUpdateAccountBizRoleGrant(request *EnterpriseAccountUpdateAccountBizRoleGrantRequest) (_result *EnterpriseAccountUpdateAccountBizRoleGrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateAccountBizRoleGrantResponse{}
	_body, _err := client.EnterpriseAccountUpdateAccountBizRoleGrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置Ip掩码
//
// @param request - EnterpriseAccountUpdateIpMaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateIpMaskResponse
func (client *Client) EnterpriseAccountUpdateIpMaskWithOptions(request *EnterpriseAccountUpdateIpMaskRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateIpMaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpMasksJson) {
		query["IpMasksJson"] = request.IpMasksJson
	}

	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountUpdateIpMask"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountUpdateIpMaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置Ip掩码
//
// @param request - EnterpriseAccountUpdateIpMaskRequest
//
// @return EnterpriseAccountUpdateIpMaskResponse
func (client *Client) EnterpriseAccountUpdateIpMask(request *EnterpriseAccountUpdateIpMaskRequest) (_result *EnterpriseAccountUpdateIpMaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateIpMaskResponse{}
	_body, _err := client.EnterpriseAccountUpdateIpMaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新操作风控
//
// @param request - EnterpriseAccountUpdateOperateRiskControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateOperateRiskControlResponse
func (client *Client) EnterpriseAccountUpdateOperateRiskControlWithOptions(request *EnterpriseAccountUpdateOperateRiskControlRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateOperateRiskControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.ProductLevel) {
		query["ProductLevel"] = request.ProductLevel
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.ValidateType) {
		query["ValidateType"] = request.ValidateType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountUpdateOperateRiskControl"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountUpdateOperateRiskControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新操作风控
//
// @param request - EnterpriseAccountUpdateOperateRiskControlRequest
//
// @return EnterpriseAccountUpdateOperateRiskControlResponse
func (client *Client) EnterpriseAccountUpdateOperateRiskControl(request *EnterpriseAccountUpdateOperateRiskControlRequest) (_result *EnterpriseAccountUpdateOperateRiskControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateOperateRiskControlResponse{}
	_body, _err := client.EnterpriseAccountUpdateOperateRiskControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改安全手机启用状态
//
// @param request - EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateSecurityMobileLoginStatusResponse
func (client *Client) EnterpriseAccountUpdateSecurityMobileLoginStatusWithOptions(request *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountUpdateSecurityMobileLoginStatus"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountUpdateSecurityMobileLoginStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改安全手机启用状态
//
// @param request - EnterpriseAccountUpdateSecurityMobileLoginStatusRequest
//
// @return EnterpriseAccountUpdateSecurityMobileLoginStatusResponse
func (client *Client) EnterpriseAccountUpdateSecurityMobileLoginStatus(request *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest) (_result *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateSecurityMobileLoginStatusResponse{}
	_body, _err := client.EnterpriseAccountUpdateSecurityMobileLoginStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新过期时间
//
// @param request - EnterpriseAccountUpdateSessionExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateSessionExpireTimeResponse
func (client *Client) EnterpriseAccountUpdateSessionExpireTimeWithOptions(request *EnterpriseAccountUpdateSessionExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateSessionExpireTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.Pk) {
		query["Pk"] = request.Pk
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.SessionExpireTime) {
		query["SessionExpireTime"] = request.SessionExpireTime
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseAccountUpdateSessionExpireTime"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseAccountUpdateSessionExpireTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新过期时间
//
// @param request - EnterpriseAccountUpdateSessionExpireTimeRequest
//
// @return EnterpriseAccountUpdateSessionExpireTimeResponse
func (client *Client) EnterpriseAccountUpdateSessionExpireTime(request *EnterpriseAccountUpdateSessionExpireTimeRequest) (_result *EnterpriseAccountUpdateSessionExpireTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseAccountUpdateSessionExpireTimeResponse{}
	_body, _err := client.EnterpriseAccountUpdateSessionExpireTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an enterprise contact.
//
// Description:
//
// Creates an enterprise public contact.
//
// For information about Alibaba Cloud account authorization, refer to [documentation](https://www.alibabacloud.com/help/en/account/user-guide/add-business-address-and-business-contact).
//
// @param request - EnterpriseContactAddRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactAddResponse
func (client *Client) EnterpriseContactAddWithOptions(request *EnterpriseContactAddRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactAddResponse, _err error) {
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

	if !dara.IsNil(request.AsyncEmailVerify) {
		body["AsyncEmailVerify"] = request.AsyncEmailVerify
	}

	if !dara.IsNil(request.AsyncMobileVerify) {
		body["AsyncMobileVerify"] = request.AsyncMobileVerify
	}

	if !dara.IsNil(request.ContactEmail) {
		body["ContactEmail"] = request.ContactEmail
	}

	if !dara.IsNil(request.ContactMobile) {
		body["ContactMobile"] = request.ContactMobile
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.ContactPosition) {
		body["ContactPosition"] = request.ContactPosition
	}

	if !dara.IsNil(request.EmailCode) {
		body["EmailCode"] = request.EmailCode
	}

	if !dara.IsNil(request.MobileCode) {
		body["MobileCode"] = request.MobileCode
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.SharedContact) {
		body["SharedContact"] = request.SharedContact
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseContactAdd"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseContactAddResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an enterprise contact.
//
// Description:
//
// Creates an enterprise public contact.
//
// For information about Alibaba Cloud account authorization, refer to [documentation](https://www.alibabacloud.com/help/en/account/user-guide/add-business-address-and-business-contact).
//
// @param request - EnterpriseContactAddRequest
//
// @return EnterpriseContactAddResponse
func (client *Client) EnterpriseContactAdd(request *EnterpriseContactAddRequest) (_result *EnterpriseContactAddResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseContactAddResponse{}
	_body, _err := client.EnterpriseContactAddWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an enterprise contact.
//
// Description:
//
// Deletes an enterprise public contact. For information about Alibaba Cloud account authorization, refer to the [documentation](https://www.alibabacloud.com/help/en/account/user-guide/add-business-address-and-business-contact).
//
// @param request - EnterpriseContactDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactDeleteResponse
func (client *Client) EnterpriseContactDeleteWithOptions(request *EnterpriseContactDeleteRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactDeleteResponse, _err error) {
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

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseContactDelete"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseContactDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an enterprise contact.
//
// Description:
//
// Deletes an enterprise public contact. For information about Alibaba Cloud account authorization, refer to the [documentation](https://www.alibabacloud.com/help/en/account/user-guide/add-business-address-and-business-contact).
//
// @param request - EnterpriseContactDeleteRequest
//
// @return EnterpriseContactDeleteResponse
func (client *Client) EnterpriseContactDelete(request *EnterpriseContactDeleteRequest) (_result *EnterpriseContactDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseContactDeleteResponse{}
	_body, _err := client.EnterpriseContactDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a private enterprise contact.
//
// Description:
//
// Modifies a public enterprise contact. For information about primary account authorization, see [documentation](https://www.alibabacloud.com/help/en/account/user-guide/add-business-address-and-business-contact).
//
// @param request - EnterpriseContactEditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactEditResponse
func (client *Client) EnterpriseContactEditWithOptions(request *EnterpriseContactEditRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactEditResponse, _err error) {
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

	if !dara.IsNil(request.AsyncEmailVerify) {
		body["AsyncEmailVerify"] = request.AsyncEmailVerify
	}

	if !dara.IsNil(request.AsyncMobileVerify) {
		body["AsyncMobileVerify"] = request.AsyncMobileVerify
	}

	if !dara.IsNil(request.ContactEmail) {
		body["ContactEmail"] = request.ContactEmail
	}

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.ContactMobile) {
		body["ContactMobile"] = request.ContactMobile
	}

	if !dara.IsNil(request.ContactName) {
		body["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.ContactPosition) {
		body["ContactPosition"] = request.ContactPosition
	}

	if !dara.IsNil(request.EmailCode) {
		body["EmailCode"] = request.EmailCode
	}

	if !dara.IsNil(request.MobileCode) {
		body["MobileCode"] = request.MobileCode
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.SharedContact) {
		body["SharedContact"] = request.SharedContact
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseContactEdit"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseContactEditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a private enterprise contact.
//
// Description:
//
// Modifies a public enterprise contact. For information about primary account authorization, see [documentation](https://www.alibabacloud.com/help/en/account/user-guide/add-business-address-and-business-contact).
//
// @param request - EnterpriseContactEditRequest
//
// @return EnterpriseContactEditResponse
func (client *Client) EnterpriseContactEdit(request *EnterpriseContactEditRequest) (_result *EnterpriseContactEditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseContactEditResponse{}
	_body, _err := client.EnterpriseContactEditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a contact.
//
// Description:
//
// Queries the details of a single enterprise contact.
//
// @param request - EnterpriseContactQueryDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactQueryDetailResponse
func (client *Client) EnterpriseContactQueryDetailWithOptions(request *EnterpriseContactQueryDetailRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactQueryDetailResponse, _err error) {
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

	if !dara.IsNil(request.ContactId) {
		body["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseContactQueryDetail"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseContactQueryDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a contact.
//
// Description:
//
// Queries the details of a single enterprise contact.
//
// @param request - EnterpriseContactQueryDetailRequest
//
// @return EnterpriseContactQueryDetailResponse
func (client *Client) EnterpriseContactQueryDetail(request *EnterpriseContactQueryDetailRequest) (_result *EnterpriseContactQueryDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseContactQueryDetailResponse{}
	_body, _err := client.EnterpriseContactQueryDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the contact list.
//
// Description:
//
// Query enterprise contacts by page.
//
// @param request - EnterpriseContactQueryPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactQueryPageListResponse
func (client *Client) EnterpriseContactQueryPageListWithOptions(request *EnterpriseContactQueryPageListRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactQueryPageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ShowCompleteInfo) {
		query["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivateContact) {
		body["PrivateContact"] = request.PrivateContact
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.SharedContact) {
		body["SharedContact"] = request.SharedContact
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseContactQueryPageList"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseContactQueryPageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the contact list.
//
// Description:
//
// Query enterprise contacts by page.
//
// @param request - EnterpriseContactQueryPageListRequest
//
// @return EnterpriseContactQueryPageListResponse
func (client *Client) EnterpriseContactQueryPageList(request *EnterpriseContactQueryPageListRequest) (_result *EnterpriseContactQueryPageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseContactQueryPageListResponse{}
	_body, _err := client.EnterpriseContactQueryPageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建组织节点
//
// @param tmpReq - EnterpriseOrgCreateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseOrgCreateNodeResponse
func (client *Client) EnterpriseOrgCreateNodeWithOptions(tmpReq *EnterpriseOrgCreateNodeRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseOrgCreateNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EnterpriseOrgCreateNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("Ext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.ExtShrink) {
		body["Ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.IsOpenApi) {
		body["IsOpenApi"] = request.IsOpenApi
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeName) {
		body["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.NodeType) {
		body["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.ParentNodeId) {
		body["ParentNodeId"] = request.ParentNodeId
	}

	if !dara.IsNil(request.ParentNodeType) {
		body["ParentNodeType"] = request.ParentNodeType
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !dara.IsNil(request.TreeId) {
		body["TreeId"] = request.TreeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseOrgCreateNode"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseOrgCreateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建组织节点
//
// @param request - EnterpriseOrgCreateNodeRequest
//
// @return EnterpriseOrgCreateNodeResponse
func (client *Client) EnterpriseOrgCreateNode(request *EnterpriseOrgCreateNodeRequest) (_result *EnterpriseOrgCreateNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseOrgCreateNodeResponse{}
	_body, _err := client.EnterpriseOrgCreateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除组织节点
//
// @param tmpReq - EnterpriseOrgDeleteNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseOrgDeleteNodeResponse
func (client *Client) EnterpriseOrgDeleteNodeWithOptions(tmpReq *EnterpriseOrgDeleteNodeRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseOrgDeleteNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EnterpriseOrgDeleteNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("Ext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.ExtShrink) {
		body["Ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.IsOpenApi) {
		body["IsOpenApi"] = request.IsOpenApi
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeType) {
		body["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !dara.IsNil(request.TreeId) {
		body["TreeId"] = request.TreeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseOrgDeleteNode"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseOrgDeleteNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除组织节点
//
// @param request - EnterpriseOrgDeleteNodeRequest
//
// @return EnterpriseOrgDeleteNodeResponse
func (client *Client) EnterpriseOrgDeleteNode(request *EnterpriseOrgDeleteNodeRequest) (_result *EnterpriseOrgDeleteNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseOrgDeleteNodeResponse{}
	_body, _err := client.EnterpriseOrgDeleteNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 组织目录树查询
//
// @param request - EnterpriseOrgQueryLoadTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseOrgQueryLoadTreeResponse
func (client *Client) EnterpriseOrgQueryLoadTreeWithOptions(request *EnterpriseOrgQueryLoadTreeRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseOrgQueryLoadTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.LoadOrgOnly) {
		query["LoadOrgOnly"] = request.LoadOrgOnly
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseOrgQueryLoadTree"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseOrgQueryLoadTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 组织目录树查询
//
// @param request - EnterpriseOrgQueryLoadTreeRequest
//
// @return EnterpriseOrgQueryLoadTreeResponse
func (client *Client) EnterpriseOrgQueryLoadTree(request *EnterpriseOrgQueryLoadTreeRequest) (_result *EnterpriseOrgQueryLoadTreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseOrgQueryLoadTreeResponse{}
	_body, _err := client.EnterpriseOrgQueryLoadTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重命名组织节点
//
// @param tmpReq - EnterpriseOrgRenameNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseOrgRenameNodeResponse
func (client *Client) EnterpriseOrgRenameNodeWithOptions(tmpReq *EnterpriseOrgRenameNodeRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseOrgRenameNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EnterpriseOrgRenameNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("Ext"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.ExtShrink) {
		body["Ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.IsOpenApi) {
		body["IsOpenApi"] = request.IsOpenApi
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeName) {
		body["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.NodeType) {
		body["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !dara.IsNil(request.TreeId) {
		body["TreeId"] = request.TreeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseOrgRenameNode"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseOrgRenameNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重命名组织节点
//
// @param request - EnterpriseOrgRenameNodeRequest
//
// @return EnterpriseOrgRenameNodeResponse
func (client *Client) EnterpriseOrgRenameNode(request *EnterpriseOrgRenameNodeRequest) (_result *EnterpriseOrgRenameNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseOrgRenameNodeResponse{}
	_body, _err := client.EnterpriseOrgRenameNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建成员账号
//
// @param request - EnterpriseRegisterAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRegisterAccountResponse
func (client *Client) EnterpriseRegisterAccountWithOptions(request *EnterpriseRegisterAccountRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRegisterAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		query["Alias"] = request.Alias
	}

	if !dara.IsNil(request.EncryptPassword) {
		query["EncryptPassword"] = request.EncryptPassword
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.LoginEmail) {
		query["LoginEmail"] = request.LoginEmail
	}

	if !dara.IsNil(request.OrganizationId) {
		query["OrganizationId"] = request.OrganizationId
	}

	if !dara.IsNil(request.RequestId) {
		query["RequestId"] = request.RequestId
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		query["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !dara.IsNil(request.SiteNick) {
		query["SiteNick"] = request.SiteNick
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseRegisterAccount"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseRegisterAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建成员账号
//
// @param request - EnterpriseRegisterAccountRequest
//
// @return EnterpriseRegisterAccountResponse
func (client *Client) EnterpriseRegisterAccount(request *EnterpriseRegisterAccountRequest) (_result *EnterpriseRegisterAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseRegisterAccountResponse{}
	_body, _err := client.EnterpriseRegisterAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建业务角色
//
// @param request - EnterpriseRoleCreateBizRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleCreateBizRoleResponse
func (client *Client) EnterpriseRoleCreateBizRoleWithOptions(request *EnterpriseRoleCreateBizRoleRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleCreateBizRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizPermissionCodeListJson) {
		query["BizPermissionCodeListJson"] = request.BizPermissionCodeListJson
	}

	if !dara.IsNil(request.BizRoleDesc) {
		query["BizRoleDesc"] = request.BizRoleDesc
	}

	if !dara.IsNil(request.BizRoleName) {
		query["BizRoleName"] = request.BizRoleName
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseRoleCreateBizRole"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseRoleCreateBizRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建业务角色
//
// @param request - EnterpriseRoleCreateBizRoleRequest
//
// @return EnterpriseRoleCreateBizRoleResponse
func (client *Client) EnterpriseRoleCreateBizRole(request *EnterpriseRoleCreateBizRoleRequest) (_result *EnterpriseRoleCreateBizRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseRoleCreateBizRoleResponse{}
	_body, _err := client.EnterpriseRoleCreateBizRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除业务角色
//
// @param request - EnterpriseRoleDeleteBizRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleDeleteBizRoleResponse
func (client *Client) EnterpriseRoleDeleteBizRoleWithOptions(request *EnterpriseRoleDeleteBizRoleRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleDeleteBizRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRoleCode) {
		query["BizRoleCode"] = request.BizRoleCode
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseRoleDeleteBizRole"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseRoleDeleteBizRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除业务角色
//
// @param request - EnterpriseRoleDeleteBizRoleRequest
//
// @return EnterpriseRoleDeleteBizRoleResponse
func (client *Client) EnterpriseRoleDeleteBizRole(request *EnterpriseRoleDeleteBizRoleRequest) (_result *EnterpriseRoleDeleteBizRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseRoleDeleteBizRoleResponse{}
	_body, _err := client.EnterpriseRoleDeleteBizRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 角色授权场景下分页查询账号
//
// @param request - EnterpriseRoleQueryAccountForRoleGrantByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryAccountForRoleGrantByPageResponse
func (client *Client) EnterpriseRoleQueryAccountForRoleGrantByPageWithOptions(request *EnterpriseRoleQueryAccountForRoleGrantByPageRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleQueryAccountForRoleGrantByPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRoleCode) {
		query["BizRoleCode"] = request.BizRoleCode
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrgId) {
		query["OrgId"] = request.OrgId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		query["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseRoleQueryAccountForRoleGrantByPage"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseRoleQueryAccountForRoleGrantByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 角色授权场景下分页查询账号
//
// @param request - EnterpriseRoleQueryAccountForRoleGrantByPageRequest
//
// @return EnterpriseRoleQueryAccountForRoleGrantByPageResponse
func (client *Client) EnterpriseRoleQueryAccountForRoleGrantByPage(request *EnterpriseRoleQueryAccountForRoleGrantByPageRequest) (_result *EnterpriseRoleQueryAccountForRoleGrantByPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseRoleQueryAccountForRoleGrantByPageResponse{}
	_body, _err := client.EnterpriseRoleQueryAccountForRoleGrantByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询业务角色
//
// @param request - EnterpriseRoleQueryBizRoleByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryBizRoleByPageResponse
func (client *Client) EnterpriseRoleQueryBizRoleByPageWithOptions(request *EnterpriseRoleQueryBizRoleByPageRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleQueryBizRoleByPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrientedLeId) {
		query["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SrcType) {
		query["SrcType"] = request.SrcType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseRoleQueryBizRoleByPage"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseRoleQueryBizRoleByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询业务角色
//
// @param request - EnterpriseRoleQueryBizRoleByPageRequest
//
// @return EnterpriseRoleQueryBizRoleByPageResponse
func (client *Client) EnterpriseRoleQueryBizRoleByPage(request *EnterpriseRoleQueryBizRoleByPageRequest) (_result *EnterpriseRoleQueryBizRoleByPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseRoleQueryBizRoleByPageResponse{}
	_body, _err := client.EnterpriseRoleQueryBizRoleByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询业务角色详情
//
// @param request - EnterpriseRoleQueryBizRoleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryBizRoleDetailResponse
func (client *Client) EnterpriseRoleQueryBizRoleDetailWithOptions(request *EnterpriseRoleQueryBizRoleDetailRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleQueryBizRoleDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizRoleCode) {
		query["BizRoleCode"] = request.BizRoleCode
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseRoleQueryBizRoleDetail"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseRoleQueryBizRoleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询业务角色详情
//
// @param request - EnterpriseRoleQueryBizRoleDetailRequest
//
// @return EnterpriseRoleQueryBizRoleDetailResponse
func (client *Client) EnterpriseRoleQueryBizRoleDetail(request *EnterpriseRoleQueryBizRoleDetailRequest) (_result *EnterpriseRoleQueryBizRoleDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseRoleQueryBizRoleDetailResponse{}
	_body, _err := client.EnterpriseRoleQueryBizRoleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新业务角色
//
// @param request - EnterpriseRoleUpdateBizRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleUpdateBizRoleResponse
func (client *Client) EnterpriseRoleUpdateBizRoleWithOptions(request *EnterpriseRoleUpdateBizRoleRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleUpdateBizRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizPermissionCodeListJson) {
		query["BizPermissionCodeListJson"] = request.BizPermissionCodeListJson
	}

	if !dara.IsNil(request.BizRoleCode) {
		query["BizRoleCode"] = request.BizRoleCode
	}

	if !dara.IsNil(request.BizRoleDesc) {
		query["BizRoleDesc"] = request.BizRoleDesc
	}

	if !dara.IsNil(request.BizRoleName) {
		query["BizRoleName"] = request.BizRoleName
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseRoleUpdateBizRole"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseRoleUpdateBizRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新业务角色
//
// @param request - EnterpriseRoleUpdateBizRoleRequest
//
// @return EnterpriseRoleUpdateBizRoleResponse
func (client *Client) EnterpriseRoleUpdateBizRole(request *EnterpriseRoleUpdateBizRoleRequest) (_result *EnterpriseRoleUpdateBizRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseRoleUpdateBizRoleResponse{}
	_body, _err := client.EnterpriseRoleUpdateBizRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 处理待办项
//
// @param request - EnterpriseTodoDealAccountTodoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoDealAccountTodoResponse
func (client *Client) EnterpriseTodoDealAccountTodoWithOptions(request *EnterpriseTodoDealAccountTodoRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseTodoDealAccountTodoResponse, _err error) {
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

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TodoId) {
		body["TodoId"] = request.TodoId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseTodoDealAccountTodo"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseTodoDealAccountTodoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 处理待办项
//
// @param request - EnterpriseTodoDealAccountTodoRequest
//
// @return EnterpriseTodoDealAccountTodoResponse
func (client *Client) EnterpriseTodoDealAccountTodo(request *EnterpriseTodoDealAccountTodoRequest) (_result *EnterpriseTodoDealAccountTodoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseTodoDealAccountTodoResponse{}
	_body, _err := client.EnterpriseTodoDealAccountTodoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询当前登录用户处理的待办项列表
//
// @param request - EnterpriseTodoQueryAccountTodoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoQueryAccountTodoListResponse
func (client *Client) EnterpriseTodoQueryAccountTodoListWithOptions(request *EnterpriseTodoQueryAccountTodoListRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseTodoQueryAccountTodoListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OperatePk) {
		body["OperatePk"] = request.OperatePk
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.Page) {
		body["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TodoType) {
		body["TodoType"] = request.TodoType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseTodoQueryAccountTodoList"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseTodoQueryAccountTodoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前登录用户处理的待办项列表
//
// @param request - EnterpriseTodoQueryAccountTodoListRequest
//
// @return EnterpriseTodoQueryAccountTodoListResponse
func (client *Client) EnterpriseTodoQueryAccountTodoList(request *EnterpriseTodoQueryAccountTodoListRequest) (_result *EnterpriseTodoQueryAccountTodoListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseTodoQueryAccountTodoListResponse{}
	_body, _err := client.EnterpriseTodoQueryAccountTodoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询当前登录用户发起的待办项列表
//
// @param request - EnterpriseTodoQueryAccountTodoListByApplicantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoQueryAccountTodoListByApplicantResponse
func (client *Client) EnterpriseTodoQueryAccountTodoListByApplicantWithOptions(request *EnterpriseTodoQueryAccountTodoListByApplicantRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseTodoQueryAccountTodoListByApplicantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OperatePk) {
		body["OperatePk"] = request.OperatePk
	}

	if !dara.IsNil(request.OrientedEcId) {
		body["OrientedEcId"] = request.OrientedEcId
	}

	if !dara.IsNil(request.OrientedLeId) {
		body["OrientedLeId"] = request.OrientedLeId
	}

	if !dara.IsNil(request.OrientedNbId) {
		body["OrientedNbId"] = request.OrientedNbId
	}

	if !dara.IsNil(request.Page) {
		body["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShowCompleteInfo) {
		body["ShowCompleteInfo"] = request.ShowCompleteInfo
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TodoType) {
		body["TodoType"] = request.TodoType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseTodoQueryAccountTodoListByApplicant"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseTodoQueryAccountTodoListByApplicantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询当前登录用户发起的待办项列表
//
// @param request - EnterpriseTodoQueryAccountTodoListByApplicantRequest
//
// @return EnterpriseTodoQueryAccountTodoListByApplicantResponse
func (client *Client) EnterpriseTodoQueryAccountTodoListByApplicant(request *EnterpriseTodoQueryAccountTodoListByApplicantRequest) (_result *EnterpriseTodoQueryAccountTodoListByApplicantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseTodoQueryAccountTodoListByApplicantResponse{}
	_body, _err := client.EnterpriseTodoQueryAccountTodoListByApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 管理员邀请纳管
//
// @param request - EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseUninvitedAdminInviteJoinEnterpriseResponse
func (client *Client) EnterpriseUninvitedAdminInviteJoinEnterpriseWithOptions(request *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcId) {
		query["EcId"] = request.EcId
	}

	if !dara.IsNil(request.EncryptTicket) {
		query["EncryptTicket"] = request.EncryptTicket
	}

	if !dara.IsNil(request.InviteePk) {
		query["InviteePk"] = request.InviteePk
	}

	if !dara.IsNil(request.LeId) {
		query["LeId"] = request.LeId
	}

	if !dara.IsNil(request.NbId) {
		query["NbId"] = request.NbId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnterpriseUninvitedAdminInviteJoinEnterprise"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnterpriseUninvitedAdminInviteJoinEnterpriseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 管理员邀请纳管
//
// @param request - EnterpriseUninvitedAdminInviteJoinEnterpriseRequest
//
// @return EnterpriseUninvitedAdminInviteJoinEnterpriseResponse
func (client *Client) EnterpriseUninvitedAdminInviteJoinEnterprise(request *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest) (_result *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnterpriseUninvitedAdminInviteJoinEnterpriseResponse{}
	_body, _err := client.EnterpriseUninvitedAdminInviteJoinEnterpriseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends an asynchronous email verification message to verify the email address of a specified contact.
//
// Description:
//
// Sends an asynchronous verification link for a contact\\"s email address. Additional rate limits apply. The same account and contact information combination cannot exceed 20 requests within 5 minutes. The same account cannot exceed 300 requests globally within 24 hours.
//
// @param request - SendAsyncEmailCaptchaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAsyncEmailCaptchaResponse
func (client *Client) SendAsyncEmailCaptchaWithOptions(request *SendAsyncEmailCaptchaRequest, runtime *dara.RuntimeOptions) (_result *SendAsyncEmailCaptchaResponse, _err error) {
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

	if !dara.IsNil(request.ContactInfo) {
		body["ContactInfo"] = request.ContactInfo
	}

	if !dara.IsNil(request.ContactorId) {
		body["ContactorId"] = request.ContactorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendAsyncEmailCaptcha"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendAsyncEmailCaptchaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends an asynchronous email verification message to verify the email address of a specified contact.
//
// Description:
//
// Sends an asynchronous verification link for a contact\\"s email address. Additional rate limits apply. The same account and contact information combination cannot exceed 20 requests within 5 minutes. The same account cannot exceed 300 requests globally within 24 hours.
//
// @param request - SendAsyncEmailCaptchaRequest
//
// @return SendAsyncEmailCaptchaResponse
func (client *Client) SendAsyncEmailCaptcha(request *SendAsyncEmailCaptchaRequest) (_result *SendAsyncEmailCaptchaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendAsyncEmailCaptchaResponse{}
	_body, _err := client.SendAsyncEmailCaptchaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends an asynchronous verification SMS to a phone number to verify the phone number of a specified contact.
//
// Description:
//
// Sends an asynchronous verification link for a contact\\"s contact information. Additional rate limits apply. The same account and contact information combination cannot exceed 20 requests within 5 minutes. The same account cannot exceed 300 requests globally within 24 hours.
//
// @param request - SendAsyncMobileCaptchaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAsyncMobileCaptchaResponse
func (client *Client) SendAsyncMobileCaptchaWithOptions(request *SendAsyncMobileCaptchaRequest, runtime *dara.RuntimeOptions) (_result *SendAsyncMobileCaptchaResponse, _err error) {
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

	if !dara.IsNil(request.ContactInfo) {
		body["ContactInfo"] = request.ContactInfo
	}

	if !dara.IsNil(request.ContactorId) {
		body["ContactorId"] = request.ContactorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendAsyncMobileCaptcha"),
		Version:     dara.String("2024-12-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendAsyncMobileCaptchaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends an asynchronous verification SMS to a phone number to verify the phone number of a specified contact.
//
// Description:
//
// Sends an asynchronous verification link for a contact\\"s contact information. Additional rate limits apply. The same account and contact information combination cannot exceed 20 requests within 5 minutes. The same account cannot exceed 300 requests globally within 24 hours.
//
// @param request - SendAsyncMobileCaptchaRequest
//
// @return SendAsyncMobileCaptchaResponse
func (client *Client) SendAsyncMobileCaptcha(request *SendAsyncMobileCaptchaRequest) (_result *SendAsyncMobileCaptchaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendAsyncMobileCaptchaResponse{}
	_body, _err := client.SendAsyncMobileCaptchaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
