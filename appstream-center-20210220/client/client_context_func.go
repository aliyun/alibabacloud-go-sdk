// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 身份认证查询接口
//
// @param tmpReq - FindIdpListByLoginIdentifierRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindIdpListByLoginIdentifierResponse
func (client *Client) FindIdpListByLoginIdentifierWithContext(ctx context.Context, tmpReq *FindIdpListByLoginIdentifierRequest, runtime *dara.RuntimeOptions) (_result *FindIdpListByLoginIdentifierResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &FindIdpListByLoginIdentifierShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AvailableFeatures) {
		request.AvailableFeaturesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvailableFeatures, dara.String("AvailableFeatures"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AvailableFeaturesShrink) {
		query["AvailableFeatures"] = request.AvailableFeaturesShrink
	}

	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientChannel) {
		body["ClientChannel"] = request.ClientChannel
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.LoginIdentifier) {
		body["LoginIdentifier"] = request.LoginIdentifier
	}

	if !dara.IsNil(request.SupportTypes) {
		body["SupportTypes"] = request.SupportTypes
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindIdpListByLoginIdentifier"),
		Version:     dara.String("2021-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindIdpListByLoginIdentifierResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetLoginToken
//
// @param tmpReq - GetLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginTokenResponse
func (client *Client) GetLoginTokenWithContext(ctx context.Context, tmpReq *GetLoginTokenRequest, runtime *dara.RuntimeOptions) (_result *GetLoginTokenResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetLoginTokenShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AvailableFeatures) {
		request.AvailableFeaturesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvailableFeatures, dara.String("AvailableFeatures"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationCode) {
		query["AuthenticationCode"] = request.AuthenticationCode
	}

	if !dara.IsNil(request.AvailableFeaturesShrink) {
		query["AvailableFeatures"] = request.AvailableFeaturesShrink
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientName) {
		query["ClientName"] = request.ClientName
	}

	if !dara.IsNil(request.ClientOS) {
		query["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientVersion) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.CurrentStage) {
		query["CurrentStage"] = request.CurrentStage
	}

	if !dara.IsNil(request.DirectoryId) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EncryptedFingerPrintData) {
		query["EncryptedFingerPrintData"] = request.EncryptedFingerPrintData
	}

	if !dara.IsNil(request.EncryptedKey) {
		query["EncryptedKey"] = request.EncryptedKey
	}

	if !dara.IsNil(request.EncryptedPassword) {
		query["EncryptedPassword"] = request.EncryptedPassword
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.FingerPrintData) {
		query["FingerPrintData"] = request.FingerPrintData
	}

	if !dara.IsNil(request.IdpId) {
		query["IdpId"] = request.IdpId
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.KeepAlive) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !dara.IsNil(request.KeepAliveToken) {
		query["KeepAliveToken"] = request.KeepAliveToken
	}

	if !dara.IsNil(request.LoginIdentifier) {
		query["LoginIdentifier"] = request.LoginIdentifier
	}

	if !dara.IsNil(request.LoginName) {
		query["LoginName"] = request.LoginName
	}

	if !dara.IsNil(request.MfaType) {
		query["MfaType"] = request.MfaType
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NewPassword) {
		query["NewPassword"] = request.NewPassword
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.OldPassword) {
		query["OldPassword"] = request.OldPassword
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.PhoneVerifyCode) {
		query["PhoneVerifyCode"] = request.PhoneVerifyCode
	}

	if !dara.IsNil(request.ProfileRegion) {
		query["ProfileRegion"] = request.ProfileRegion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SsoExtendsCookies) {
		query["SsoExtendsCookies"] = request.SsoExtendsCookies
	}

	if !dara.IsNil(request.SsoSessionToken) {
		query["SsoSessionToken"] = request.SsoSessionToken
	}

	if !dara.IsNil(request.TokenCode) {
		query["TokenCode"] = request.TokenCode
	}

	if !dara.IsNil(request.UmidToken) {
		query["UmidToken"] = request.UmidToken
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLoginToken"),
		Version:     dara.String("2021-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取无影StsToken
//
// @param request - GetStsTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStsTokenResponse
func (client *Client) GetStsTokenWithContext(ctx context.Context, request *GetStsTokenRequest, runtime *dara.RuntimeOptions) (_result *GetStsTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		body["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStsToken"),
		Version:     dara.String("2021-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStsTokenResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefreshLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshLoginTokenResponse
func (client *Client) RefreshLoginTokenWithContext(ctx context.Context, request *RefreshLoginTokenRequest, runtime *dara.RuntimeOptions) (_result *RefreshLoginTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.EndUserId) {
		query["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginIdentifier) {
		query["LoginIdentifier"] = request.LoginIdentifier
	}

	if !dara.IsNil(request.LoginToken) {
		query["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.ProfileRegion) {
		query["ProfileRegion"] = request.ProfileRegion
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshLoginToken"),
		Version:     dara.String("2021-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
