// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Verifies SMS verification codes.
//
// @param request - CheckSmsVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSmsVerifyCodeResponse
func (client *Client) CheckSmsVerifyCodeWithContext(ctx context.Context, request *CheckSmsVerifyCodeRequest, runtime *dara.RuntimeOptions) (_result *CheckSmsVerifyCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaseAuthPolicy) {
		query["CaseAuthPolicy"] = request.CaseAuthPolicy
	}

	if !dara.IsNil(request.CountryCode) {
		query["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

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

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
	}

	if !dara.IsNil(request.VerifyCode) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckSmsVerifyCode"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSmsVerifyCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a code for a converged communication authentication service.
//
// @param request - CreateSchemeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchemeConfigResponse
func (client *Client) CreateSchemeConfigWithContext(ctx context.Context, request *CreateSchemeConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSchemeConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPackageName) {
		query["AndroidPackageName"] = request.AndroidPackageName
	}

	if !dara.IsNil(request.AndroidPackageSign) {
		query["AndroidPackageSign"] = request.AndroidPackageSign
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.H5Origin) {
		query["H5Origin"] = request.H5Origin
	}

	if !dara.IsNil(request.H5Url) {
		query["H5Url"] = request.H5Url
	}

	if !dara.IsNil(request.IosBundleId) {
		query["IosBundleId"] = request.IosBundleId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSchemeConfig"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSchemeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVerifySchemeResponse
func (client *Client) CreateVerifySchemeWithContext(ctx context.Context, request *CreateVerifySchemeRequest, runtime *dara.RuntimeOptions) (_result *CreateVerifySchemeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.CmApiCode) {
		query["CmApiCode"] = request.CmApiCode
	}

	if !dara.IsNil(request.CtApiCode) {
		query["CtApiCode"] = request.CtApiCode
	}

	if !dara.IsNil(request.CuApiCode) {
		query["CuApiCode"] = request.CuApiCode
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.HmAppIdentifier) {
		query["HmAppIdentifier"] = request.HmAppIdentifier
	}

	if !dara.IsNil(request.HmPackageName) {
		query["HmPackageName"] = request.HmPackageName
	}

	if !dara.IsNil(request.HmSignName) {
		query["HmSignName"] = request.HmSignName
	}

	if !dara.IsNil(request.IpWhiteList) {
		query["IpWhiteList"] = request.IpWhiteList
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PackName) {
		query["PackName"] = request.PackName
	}

	if !dara.IsNil(request.PackSign) {
		query["PackSign"] = request.PackSign
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
	}

	if !dara.IsNil(request.SmsSignName) {
		query["SmsSignName"] = request.SmsSignName
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVerifyScheme"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVerifySchemeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVerifySchemeResponse
func (client *Client) DeleteVerifySchemeWithContext(ctx context.Context, request *DeleteVerifySchemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteVerifySchemeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
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

	if !dara.IsNil(request.SchemeCode) {
		query["SchemeCode"] = request.SchemeCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVerifyScheme"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVerifySchemeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifySchemeResponse
func (client *Client) DescribeVerifySchemeWithContext(ctx context.Context, request *DescribeVerifySchemeRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifySchemeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
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

	if !dara.IsNil(request.SchemeCode) {
		query["SchemeCode"] = request.SchemeCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyScheme"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifySchemeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the authorization token used for the authentication of the phone number verification for HTML5 pages. You can obtain AccessToken and JwtToken after a successful call.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the phone number verification feature for HTML5 pages](https://help.aliyun.com/document_detail/169786.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetAuthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthTokenResponse
func (client *Client) GetAuthTokenWithContext(ctx context.Context, request *GetAuthTokenRequest, runtime *dara.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CmApiCode) {
		query["CmApiCode"] = request.CmApiCode
	}

	if !dara.IsNil(request.CtApiCode) {
		query["CtApiCode"] = request.CtApiCode
	}

	if !dara.IsNil(request.CuApiCode) {
		query["CuApiCode"] = request.CuApiCode
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
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

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthToken"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the URL for the Alipay account authorization.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account and obtain an Alibaba Cloud AccessKey pair. For more information, see [Process of communication authorization](https://help.aliyun.com/document_detail/196922.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetAuthorizationUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorizationUrlResponse
func (client *Client) GetAuthorizationUrlWithContext(ctx context.Context, request *GetAuthorizationUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAuthorizationUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
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

	if !dara.IsNil(request.SchemeId) {
		query["SchemeId"] = request.SchemeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthorizationUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthorizationUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the verification results by using the token that is obtained from the client SDKs.
//
// @param request - GetFusionAuthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFusionAuthTokenResponse
func (client *Client) GetFusionAuthTokenWithContext(ctx context.Context, request *GetFusionAuthTokenRequest, runtime *dara.RuntimeOptions) (_result *GetFusionAuthTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.DurationSeconds) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PackageName) {
		query["PackageName"] = request.PackageName
	}

	if !dara.IsNil(request.PackageSign) {
		query["PackageSign"] = request.PackageSign
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SchemeCode) {
		query["SchemeCode"] = request.SchemeCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFusionAuthToken"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFusionAuthTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a phone number for one-click logon.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable only to one-click logon or registration. You can call this operation only after you confirm the authorization on the authorization page provided by the SDK for one-click logon. You are prohibited from simulating or bypassing the authorization process. Alibaba Cloud reserves the right to terminate our services and take legal actions against such violations.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetMobileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMobileResponse
func (client *Client) GetMobileWithContext(ctx context.Context, request *GetMobileRequest, runtime *dara.RuntimeOptions) (_result *GetMobileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessToken) {
		query["AccessToken"] = request.AccessToken
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
		Action:      dara.String("GetMobile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMobileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a phone number for one-click logon. This operation is exclusive to HTML5 pages.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable only to one-click logon or registration in HTML5 pages. You can call this operation only after you confirm the authorization on the authorization page provided by the JavaScript SDK. You are prohibited from simulating or bypassing the authorization process. Alibaba Cloud reserves the right to terminate our services and take legal actions against such violations.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPhoneWithTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhoneWithTokenResponse
func (client *Client) GetPhoneWithTokenWithContext(ctx context.Context, request *GetPhoneWithTokenRequest, runtime *dara.RuntimeOptions) (_result *GetPhoneWithTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.SpToken) {
		query["SpToken"] = request.SpToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhoneWithToken"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhoneWithTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the authorization token for an SMS verification code.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the SMS verification feature](https://help.aliyun.com/document_detail/313209.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetSmsAuthTokensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsAuthTokensResponse
func (client *Client) GetSmsAuthTokensWithContext(ctx context.Context, request *GetSmsAuthTokensRequest, runtime *dara.RuntimeOptions) (_result *GetSmsAuthTokensResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.Expire) {
		query["Expire"] = request.Expire
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PackageName) {
		query["PackageName"] = request.PackageName
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.SmsCodeExpire) {
		query["SmsCodeExpire"] = request.SmsCodeExpire
	}

	if !dara.IsNil(request.SmsTemplateCode) {
		query["SmsTemplateCode"] = request.SmsTemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmsAuthTokens"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmsAuthTokensResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fees generated by a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryGateVerifyBillingPublicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGateVerifyBillingPublicResponse
func (client *Client) QueryGateVerifyBillingPublicWithContext(ctx context.Context, request *QueryGateVerifyBillingPublicRequest, runtime *dara.RuntimeOptions) (_result *QueryGateVerifyBillingPublicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.Month) {
		query["Month"] = request.Month
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGateVerifyBillingPublic"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGateVerifyBillingPublicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the calls of Phone Number Verification Service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryGateVerifyStatisticPublicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGateVerifyStatisticPublicResponse
func (client *Client) QueryGateVerifyStatisticPublicWithContext(ctx context.Context, request *QueryGateVerifyStatisticPublicRequest, runtime *dara.RuntimeOptions) (_result *QueryGateVerifyStatisticPublicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGateVerifyStatisticPublic"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGateVerifyStatisticPublicResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI QuerySendDetails is deprecated
//
// Summary:
//
// Queries the delivery status of the SMS verification code. You can query only the delivery status of the SMS verification code that is sent by calling corresponding API operations.
//
// @param request - QuerySendDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendDetailsResponse
func (client *Client) QuerySendDetailsWithContext(ctx context.Context, request *QuerySendDetailsRequest, runtime *dara.RuntimeOptions) (_result *QuerySendDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Sends SMS verification codes.
//
// @param request - SendSmsVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSmsVerifyCodeResponse
func (client *Client) SendSmsVerifyCodeWithContext(ctx context.Context, request *SendSmsVerifyCodeRequest, runtime *dara.RuntimeOptions) (_result *SendSmsVerifyCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRetry) {
		query["AutoRetry"] = request.AutoRetry
	}

	if !dara.IsNil(request.CodeLength) {
		query["CodeLength"] = request.CodeLength
	}

	if !dara.IsNil(request.CodeType) {
		query["CodeType"] = request.CodeType
	}

	if !dara.IsNil(request.CountryCode) {
		query["CountryCode"] = request.CountryCode
	}

	if !dara.IsNil(request.DuplicatePolicy) {
		query["DuplicatePolicy"] = request.DuplicatePolicy
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

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

	if !dara.IsNil(request.ReturnVerifyCode) {
		query["ReturnVerifyCode"] = request.ReturnVerifyCode
	}

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
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

	if !dara.IsNil(request.ValidTime) {
		query["ValidTime"] = request.ValidTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendSmsVerifyCode"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendSmsVerifyCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the phone number that you use.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable to only the verification of thephone number that you use. To obtain a phone number for one-click logon, call [GetMobile](https://help.aliyun.com/document_detail/189865.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyMobileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyMobileResponse
func (client *Client) VerifyMobileWithContext(ctx context.Context, request *VerifyMobileRequest, runtime *dara.RuntimeOptions) (_result *VerifyMobileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessCode) {
		query["AccessCode"] = request.AccessCode
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

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
		Action:      dara.String("VerifyMobile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyMobileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the phone number used in HTML5 pages.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the phone number verification feature for HTML5 pages](https://help.aliyun.com/document_detail/169786.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyPhoneWithTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyPhoneWithTokenResponse
func (client *Client) VerifyPhoneWithTokenWithContext(ctx context.Context, request *VerifyPhoneWithTokenRequest, runtime *dara.RuntimeOptions) (_result *VerifyPhoneWithTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.SpToken) {
		query["SpToken"] = request.SpToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyPhoneWithToken"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyPhoneWithTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies SMS verification codes.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the SMS verification feature](https://help.aliyun.com/document_detail/313209.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifySmsCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySmsCodeResponse
func (client *Client) VerifySmsCodeWithContext(ctx context.Context, request *VerifySmsCodeRequest, runtime *dara.RuntimeOptions) (_result *VerifySmsCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.SmsCode) {
		query["SmsCode"] = request.SmsCode
	}

	if !dara.IsNil(request.SmsToken) {
		query["SmsToken"] = request.SmsToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifySmsCode"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifySmsCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the verification results by using the verification token that is obtained by using the authentication token.
//
// @param request - VerifyWithFusionAuthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyWithFusionAuthTokenResponse
func (client *Client) VerifyWithFusionAuthTokenWithContext(ctx context.Context, request *VerifyWithFusionAuthTokenRequest, runtime *dara.RuntimeOptions) (_result *VerifyWithFusionAuthTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.VerifyToken) {
		query["VerifyToken"] = request.VerifyToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyWithFusionAuthToken"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyWithFusionAuthTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
