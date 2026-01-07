// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 添加私有联系人
//
// @param request - AccountContactAddRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactAddResponse
func (client *Client) AccountContactAddWithContext(ctx context.Context, request *AccountContactAddRequest, runtime *dara.RuntimeOptions) (_result *AccountContactAddResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除私有联系人
//
// @param request - AccountContactDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactDeleteResponse
func (client *Client) AccountContactDeleteWithContext(ctx context.Context, request *AccountContactDeleteRequest, runtime *dara.RuntimeOptions) (_result *AccountContactDeleteResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改私有联系人
//
// @param request - AccountContactEditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactEditResponse
func (client *Client) AccountContactEditWithContext(ctx context.Context, request *AccountContactEditRequest, runtime *dara.RuntimeOptions) (_result *AccountContactEditResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人详情
//
// @param request - AccountContactQueryDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactQueryDetailResponse
func (client *Client) AccountContactQueryDetailWithContext(ctx context.Context, request *AccountContactQueryDetailRequest, runtime *dara.RuntimeOptions) (_result *AccountContactQueryDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人列表
//
// @param request - AccountContactQueryPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountContactQueryPageListResponse
func (client *Client) AccountContactQueryPageListWithContext(ctx context.Context, request *AccountContactQueryPageListRequest, runtime *dara.RuntimeOptions) (_result *AccountContactQueryPageListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeLoginPasswordResponse
func (client *Client) EnterpriseAccountChangeLoginPasswordWithContext(ctx context.Context, request *EnterpriseAccountChangeLoginPasswordRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountChangeLoginPasswordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeSecurityEmailResponse
func (client *Client) EnterpriseAccountChangeSecurityEmailWithContext(ctx context.Context, request *EnterpriseAccountChangeSecurityEmailRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountChangeSecurityEmailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountChangeSecurityMobileResponse
func (client *Client) EnterpriseAccountChangeSecurityMobileWithContext(ctx context.Context, request *EnterpriseAccountChangeSecurityMobileRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountChangeSecurityMobileResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryAccountGrantedRolesResponse
func (client *Client) EnterpriseAccountQueryAccountGrantedRolesWithContext(ctx context.Context, request *EnterpriseAccountQueryAccountGrantedRolesRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountQueryAccountGrantedRolesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryAccountsInfoResponse
func (client *Client) EnterpriseAccountQueryAccountsInfoWithContext(ctx context.Context, request *EnterpriseAccountQueryAccountsInfoRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountQueryAccountsInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountQueryLoginSettingsResponse
func (client *Client) EnterpriseAccountQueryLoginSettingsWithContext(ctx context.Context, request *EnterpriseAccountQueryLoginSettingsRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountQueryLoginSettingsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountRemoveMfaResponse
func (client *Client) EnterpriseAccountRemoveMfaWithContext(ctx context.Context, request *EnterpriseAccountRemoveMfaRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountRemoveMfaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountSeparateEaResponse
func (client *Client) EnterpriseAccountSeparateEaWithContext(ctx context.Context, request *EnterpriseAccountSeparateEaRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountSeparateEaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateAccountAliasResponse
func (client *Client) EnterpriseAccountUpdateAccountAliasWithContext(ctx context.Context, request *EnterpriseAccountUpdateAccountAliasRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateAccountAliasResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateAccountBizRoleGrantResponse
func (client *Client) EnterpriseAccountUpdateAccountBizRoleGrantWithContext(ctx context.Context, request *EnterpriseAccountUpdateAccountBizRoleGrantRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateAccountBizRoleGrantResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateIpMaskResponse
func (client *Client) EnterpriseAccountUpdateIpMaskWithContext(ctx context.Context, request *EnterpriseAccountUpdateIpMaskRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateIpMaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateOperateRiskControlResponse
func (client *Client) EnterpriseAccountUpdateOperateRiskControlWithContext(ctx context.Context, request *EnterpriseAccountUpdateOperateRiskControlRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateOperateRiskControlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateSecurityMobileLoginStatusResponse
func (client *Client) EnterpriseAccountUpdateSecurityMobileLoginStatusWithContext(ctx context.Context, request *EnterpriseAccountUpdateSecurityMobileLoginStatusRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateSecurityMobileLoginStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseAccountUpdateSessionExpireTimeResponse
func (client *Client) EnterpriseAccountUpdateSessionExpireTimeWithContext(ctx context.Context, request *EnterpriseAccountUpdateSessionExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseAccountUpdateSessionExpireTimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加企业联系人
//
// @param request - EnterpriseContactAddRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactAddResponse
func (client *Client) EnterpriseContactAddWithContext(ctx context.Context, request *EnterpriseContactAddRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactAddResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除企业联系人
//
// @param request - EnterpriseContactDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactDeleteResponse
func (client *Client) EnterpriseContactDeleteWithContext(ctx context.Context, request *EnterpriseContactDeleteRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactDeleteResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改私企业联系人
//
// @param request - EnterpriseContactEditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactEditResponse
func (client *Client) EnterpriseContactEditWithContext(ctx context.Context, request *EnterpriseContactEditRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactEditResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人详情
//
// @param request - EnterpriseContactQueryDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactQueryDetailResponse
func (client *Client) EnterpriseContactQueryDetailWithContext(ctx context.Context, request *EnterpriseContactQueryDetailRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactQueryDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询联系人列表
//
// @param request - EnterpriseContactQueryPageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseContactQueryPageListResponse
func (client *Client) EnterpriseContactQueryPageListWithContext(ctx context.Context, request *EnterpriseContactQueryPageListRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseContactQueryPageListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseOrgQueryLoadTreeResponse
func (client *Client) EnterpriseOrgQueryLoadTreeWithContext(ctx context.Context, request *EnterpriseOrgQueryLoadTreeRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseOrgQueryLoadTreeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRegisterAccountResponse
func (client *Client) EnterpriseRegisterAccountWithContext(ctx context.Context, request *EnterpriseRegisterAccountRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRegisterAccountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleCreateBizRoleResponse
func (client *Client) EnterpriseRoleCreateBizRoleWithContext(ctx context.Context, request *EnterpriseRoleCreateBizRoleRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleCreateBizRoleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleDeleteBizRoleResponse
func (client *Client) EnterpriseRoleDeleteBizRoleWithContext(ctx context.Context, request *EnterpriseRoleDeleteBizRoleRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleDeleteBizRoleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryAccountForRoleGrantByPageResponse
func (client *Client) EnterpriseRoleQueryAccountForRoleGrantByPageWithContext(ctx context.Context, request *EnterpriseRoleQueryAccountForRoleGrantByPageRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleQueryAccountForRoleGrantByPageResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryBizRoleByPageResponse
func (client *Client) EnterpriseRoleQueryBizRoleByPageWithContext(ctx context.Context, request *EnterpriseRoleQueryBizRoleByPageRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleQueryBizRoleByPageResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleQueryBizRoleDetailResponse
func (client *Client) EnterpriseRoleQueryBizRoleDetailWithContext(ctx context.Context, request *EnterpriseRoleQueryBizRoleDetailRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleQueryBizRoleDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseRoleUpdateBizRoleResponse
func (client *Client) EnterpriseRoleUpdateBizRoleWithContext(ctx context.Context, request *EnterpriseRoleUpdateBizRoleRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseRoleUpdateBizRoleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoDealAccountTodoResponse
func (client *Client) EnterpriseTodoDealAccountTodoWithContext(ctx context.Context, request *EnterpriseTodoDealAccountTodoRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseTodoDealAccountTodoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoQueryAccountTodoListResponse
func (client *Client) EnterpriseTodoQueryAccountTodoListWithContext(ctx context.Context, request *EnterpriseTodoQueryAccountTodoListRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseTodoQueryAccountTodoListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseTodoQueryAccountTodoListByApplicantResponse
func (client *Client) EnterpriseTodoQueryAccountTodoListByApplicantWithContext(ctx context.Context, request *EnterpriseTodoQueryAccountTodoListByApplicantRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseTodoQueryAccountTodoListByApplicantResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnterpriseUninvitedAdminInviteJoinEnterpriseResponse
func (client *Client) EnterpriseUninvitedAdminInviteJoinEnterpriseWithContext(ctx context.Context, request *EnterpriseUninvitedAdminInviteJoinEnterpriseRequest, runtime *dara.RuntimeOptions) (_result *EnterpriseUninvitedAdminInviteJoinEnterpriseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号中心发送异步验证邮件
//
// @param request - SendAsyncEmailCaptchaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAsyncEmailCaptchaResponse
func (client *Client) SendAsyncEmailCaptchaWithContext(ctx context.Context, request *SendAsyncEmailCaptchaRequest, runtime *dara.RuntimeOptions) (_result *SendAsyncEmailCaptchaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号中心发送异步验证短信
//
// @param request - SendAsyncMobileCaptchaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAsyncMobileCaptchaResponse
func (client *Client) SendAsyncMobileCaptchaWithContext(ctx context.Context, request *SendAsyncMobileCaptchaRequest, runtime *dara.RuntimeOptions) (_result *SendAsyncMobileCaptchaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
