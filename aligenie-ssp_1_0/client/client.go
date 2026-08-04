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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("aligenie"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Collect or remove from favorites.
//
// @param tmpReq - AddAndRemoveFavoriteContentRequest
//
// @param headers - AddAndRemoveFavoriteContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAndRemoveFavoriteContentResponse
func (client *Client) AddAndRemoveFavoriteContentWithOptions(tmpReq *AddAndRemoveFavoriteContentRequest, headers *AddAndRemoveFavoriteContentHeaders, runtime *dara.RuntimeOptions) (_result *AddAndRemoveFavoriteContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddAndRemoveFavoriteContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenAddAndRemoveFavoriteContentRequest) {
		request.OpenAddAndRemoveFavoriteContentRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenAddAndRemoveFavoriteContentRequest, dara.String("OpenAddAndRemoveFavoriteContentRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenAddAndRemoveFavoriteContentRequestShrink) {
		body["OpenAddAndRemoveFavoriteContentRequest"] = request.OpenAddAndRemoveFavoriteContentRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAndRemoveFavoriteContent"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/AddAndRemoveFavoriteContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAndRemoveFavoriteContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Collect or remove from favorites.
//
// @param request - AddAndRemoveFavoriteContentRequest
//
// @return AddAndRemoveFavoriteContentResponse
func (client *Client) AddAndRemoveFavoriteContent(request *AddAndRemoveFavoriteContentRequest) (_result *AddAndRemoveFavoriteContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddAndRemoveFavoriteContentHeaders{}
	_result = &AddAndRemoveFavoriteContentResponse{}
	_body, _err := client.AddAndRemoveFavoriteContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Subscribe to an album.
//
// @param tmpReq - AddSubRequest
//
// @param headers - AddSubHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSubResponse
func (client *Client) AddSubWithOptions(tmpReq *AddSubRequest, headers *AddSubHeaders, runtime *dara.RuntimeOptions) (_result *AddSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddSubShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddSubscriptionInfoRequest) {
		request.AddSubscriptionInfoRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddSubscriptionInfoRequest, dara.String("AddSubscriptionInfoRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AddSubscriptionInfoRequestShrink) {
		query["AddSubscriptionInfoRequest"] = request.AddSubscriptionInfoRequestShrink
	}

	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSub"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/addSub"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Subscribe to an album.
//
// @param request - AddSubRequest
//
// @return AddSubResponse
func (client *Client) AddSub(request *AddSubRequest) (_result *AddSubResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AddSubHeaders{}
	_result = &AddSubResponse{}
	_body, _err := client.AddSubWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorize logon by specifying an Aligenie account.
//
// @param request - AuthLoginWithAligenieUserInfoRequest
//
// @param headers - AuthLoginWithAligenieUserInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthLoginWithAligenieUserInfoResponse
func (client *Client) AuthLoginWithAligenieUserInfoWithOptions(request *AuthLoginWithAligenieUserInfoRequest, headers *AuthLoginWithAligenieUserInfoHeaders, runtime *dara.RuntimeOptions) (_result *AuthLoginWithAligenieUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EncryptedAligenieUserIdentifier) {
		body["EncryptedAligenieUserIdentifier"] = request.EncryptedAligenieUserIdentifier
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthLoginWithAligenieUserInfo"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/authLoginWithAligenieUserInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthLoginWithAligenieUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorize logon by specifying an Aligenie account.
//
// @param request - AuthLoginWithAligenieUserInfoRequest
//
// @return AuthLoginWithAligenieUserInfoResponse
func (client *Client) AuthLoginWithAligenieUserInfo(request *AuthLoginWithAligenieUserInfoRequest) (_result *AuthLoginWithAligenieUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AuthLoginWithAligenieUserInfoHeaders{}
	_result = &AuthLoginWithAligenieUserInfoResponse{}
	_body, _err := client.AuthLoginWithAligenieUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generate a Genie account by phone number for authorization logon.
//
// @param request - AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest
//
// @param headers - AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse
func (client *Client) AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberWithOptions(request *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest, headers *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders, runtime *dara.RuntimeOptions) (_result *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthLoginWithAligenieUserInfoGeneratedByPhoneNumber"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/authLoginWithAligenieUserInfoGeneratedByPhoneNumber"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generate a Genie account by phone number for authorization logon.
//
// @param request - AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest
//
// @return AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse
func (client *Client) AuthLoginWithAligenieUserInfoGeneratedByPhoneNumber(request *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberRequest) (_result *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders{}
	_result = &AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberResponse{}
	_body, _err := client.AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Authorize logon by specifying a Taobao account
//
// @param request - AuthLoginWithTaobaoUserInfoRequest
//
// @param headers - AuthLoginWithTaobaoUserInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthLoginWithTaobaoUserInfoResponse
func (client *Client) AuthLoginWithTaobaoUserInfoWithOptions(request *AuthLoginWithTaobaoUserInfoRequest, headers *AuthLoginWithTaobaoUserInfoHeaders, runtime *dara.RuntimeOptions) (_result *AuthLoginWithTaobaoUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EncryptedTaobaoUserIdentifier) {
		body["EncryptedTaobaoUserIdentifier"] = request.EncryptedTaobaoUserIdentifier
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthLoginWithTaobaoUserInfo"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/authLoginWithTaobaoUserInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthLoginWithTaobaoUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Authorize logon by specifying a Taobao account
//
// @param request - AuthLoginWithTaobaoUserInfoRequest
//
// @return AuthLoginWithTaobaoUserInfoResponse
func (client *Client) AuthLoginWithTaobaoUserInfo(request *AuthLoginWithTaobaoUserInfoRequest) (_result *AuthLoginWithTaobaoUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AuthLoginWithTaobaoUserInfoHeaders{}
	_result = &AuthLoginWithTaobaoUserInfoResponse{}
	_body, _err := client.AuthLoginWithTaobaoUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorize logon by using third-party user information.
//
// @param tmpReq - AuthLoginWithThirdUserInfoRequest
//
// @param headers - AuthLoginWithThirdUserInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthLoginWithThirdUserInfoResponse
func (client *Client) AuthLoginWithThirdUserInfoWithOptions(tmpReq *AuthLoginWithThirdUserInfoRequest, headers *AuthLoginWithThirdUserInfoHeaders, runtime *dara.RuntimeOptions) (_result *AuthLoginWithThirdUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AuthLoginWithThirdUserInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtInfo) {
		request.ExtInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtInfo, dara.String("ExtInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtInfoShrink) {
		body["ExtInfo"] = request.ExtInfoShrink
	}

	if !dara.IsNil(request.SceneCode) {
		body["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.ThirdUserIdentifier) {
		body["ThirdUserIdentifier"] = request.ThirdUserIdentifier
	}

	if !dara.IsNil(request.ThirdUserType) {
		body["ThirdUserType"] = request.ThirdUserType
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthLoginWithThirdUserInfo"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/authLoginWithThirdUserInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthLoginWithThirdUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorize logon by using third-party user information.
//
// @param request - AuthLoginWithThirdUserInfoRequest
//
// @return AuthLoginWithThirdUserInfoResponse
func (client *Client) AuthLoginWithThirdUserInfo(request *AuthLoginWithThirdUserInfoRequest) (_result *AuthLoginWithThirdUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AuthLoginWithThirdUserInfoHeaders{}
	_result = &AuthLoginWithThirdUserInfoResponse{}
	_body, _err := client.AuthLoginWithThirdUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查并拨打voip电话【酒店业务】
//
// @param tmpReq - CheckAndDoVoipCallForHotelRequest
//
// @param headers - CheckAndDoVoipCallForHotelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAndDoVoipCallForHotelResponse
func (client *Client) CheckAndDoVoipCallForHotelWithOptions(tmpReq *CheckAndDoVoipCallForHotelRequest, headers *CheckAndDoVoipCallForHotelHeaders, runtime *dara.RuntimeOptions) (_result *CheckAndDoVoipCallForHotelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CheckAndDoVoipCallForHotelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizData) {
		body["BizData"] = request.BizData
	}

	if !dara.IsNil(request.CalleeNick) {
		body["CalleeNick"] = request.CalleeNick
	}

	if !dara.IsNil(request.CalleePhoneNum) {
		body["CalleePhoneNum"] = request.CalleePhoneNum
	}

	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAndDoVoipCallForHotel"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/checkAndDoVoipCallForHotel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAndDoVoipCallForHotelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查并拨打voip电话【酒店业务】
//
// @param request - CheckAndDoVoipCallForHotelRequest
//
// @return CheckAndDoVoipCallForHotelResponse
func (client *Client) CheckAndDoVoipCallForHotel(request *CheckAndDoVoipCallForHotelRequest) (_result *CheckAndDoVoipCallForHotelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CheckAndDoVoipCallForHotelHeaders{}
	_result = &CheckAndDoVoipCallForHotelResponse{}
	_body, _err := client.CheckAndDoVoipCallForHotelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Poll the result of the activation attachment.
//
// @param tmpReq - CheckAuthCodeBindForExtRequest
//
// @param headers - CheckAuthCodeBindForExtHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAuthCodeBindForExtResponse
func (client *Client) CheckAuthCodeBindForExtWithOptions(tmpReq *CheckAuthCodeBindForExtRequest, headers *CheckAuthCodeBindForExtHeaders, runtime *dara.RuntimeOptions) (_result *CheckAuthCodeBindForExtResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CheckAuthCodeBindForExtShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.EncodeKey) {
		query["EncodeKey"] = request.EncodeKey
	}

	if !dara.IsNil(request.EncodeType) {
		query["EncodeType"] = request.EncodeType
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAuthCodeBindForExt"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/checkAuthCodeBindForExt"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAuthCodeBindForExtResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Poll the result of the activation attachment.
//
// @param request - CheckAuthCodeBindForExtRequest
//
// @return CheckAuthCodeBindForExtResponse
func (client *Client) CheckAuthCodeBindForExt(request *CheckAuthCodeBindForExtRequest) (_result *CheckAuthCodeBindForExtResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CheckAuthCodeBindForExtHeaders{}
	_result = &CheckAuthCodeBindForExtResponse{}
	_body, _err := client.CheckAuthCodeBindForExtWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cloud Push Song: Create a playlist by providing the device openUuid and the Tmall Genie openUserId.
//
// Description:
//
// After synchronously validating the input parameters, the system asynchronously creates the playlist.
//
// @param tmpReq - CloudPlayerRequest
//
// @param headers - CloudPlayerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudPlayerResponse
func (client *Client) CloudPlayerWithOptions(tmpReq *CloudPlayerRequest, headers *CloudPlayerHeaders, runtime *dara.RuntimeOptions) (_result *CloudPlayerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CloudPlayerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SongIdList) {
		request.SongIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SongIdList, dara.String("SongIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurPlayIndex) {
		query["CurPlayIndex"] = request.CurPlayIndex
	}

	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PlayMode) {
		query["PlayMode"] = request.PlayMode
	}

	if !dara.IsNil(request.SongId) {
		query["SongId"] = request.SongId
	}

	if !dara.IsNil(request.SongIdListShrink) {
		query["SongIdList"] = request.SongIdListShrink
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudPlayer"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/cloud/player"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudPlayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cloud Push Song: Create a playlist by providing the device openUuid and the Tmall Genie openUserId.
//
// Description:
//
// After synchronously validating the input parameters, the system asynchronously creates the playlist.
//
// @param request - CloudPlayerRequest
//
// @return CloudPlayerResponse
func (client *Client) CloudPlayer(request *CloudPlayerRequest) (_result *CloudPlayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CloudPlayerHeaders{}
	_result = &CloudPlayerResponse{}
	_body, _err := client.CloudPlayerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create an alarm that supports three types: one-time, weekly, and legal working days.
//
// @param tmpReq - CreateAlarmRequest
//
// @param headers - CreateAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlarmResponse
func (client *Client) CreateAlarmWithOptions(tmpReq *CreateAlarmRequest, headers *CreateAlarmHeaders, runtime *dara.RuntimeOptions) (_result *CreateAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlarm"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/createAlarm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an alarm that supports three types: one-time, weekly, and legal working days.
//
// @param request - CreateAlarmRequest
//
// @return CreateAlarmResponse
func (client *Client) CreateAlarm(request *CreateAlarmRequest) (_result *CreateAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateAlarmHeaders{}
	_result = &CreateAlarmResponse{}
	_body, _err := client.CreateAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a playlist and play the specified content.
//
// @param tmpReq - CreatePlayingListRequest
//
// @param headers - CreatePlayingListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePlayingListResponse
func (client *Client) CreatePlayingListWithOptions(tmpReq *CreatePlayingListRequest, headers *CreatePlayingListHeaders, runtime *dara.RuntimeOptions) (_result *CreatePlayingListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePlayingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenCreatePlayingListRequest) {
		request.OpenCreatePlayingListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenCreatePlayingListRequest, dara.String("OpenCreatePlayingListRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenCreatePlayingListRequestShrink) {
		body["OpenCreatePlayingListRequest"] = request.OpenCreatePlayingListRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePlayingList"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/CreatePlayingList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePlayingListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a playlist and play the specified content.
//
// @param request - CreatePlayingListRequest
//
// @return CreatePlayingListResponse
func (client *Client) CreatePlayingList(request *CreatePlayingListRequest) (_result *CreatePlayingListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreatePlayingListHeaders{}
	_result = &CreatePlayingListResponse{}
	_body, _err := client.CreatePlayingListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Playlist creation uses OAuth2 authorization.
//
// @param tmpReq - CreatePlayingListOAuth2Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePlayingListOAuth2Response
func (client *Client) CreatePlayingListOAuth2WithOptions(tmpReq *CreatePlayingListOAuth2Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePlayingListOAuth2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePlayingListOAuth2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenCreatePlayingListRequest) {
		request.OpenCreatePlayingListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenCreatePlayingListRequest, dara.String("OpenCreatePlayingListRequest"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenCreatePlayingListRequestShrink) {
		body["OpenCreatePlayingListRequest"] = request.OpenCreatePlayingListRequestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePlayingListOAuth2"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/CreatePlayingListOAuth2"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePlayingListOAuth2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Playlist creation uses OAuth2 authorization.
//
// @param request - CreatePlayingListOAuth2Request
//
// @return CreatePlayingListOAuth2Response
func (client *Client) CreatePlayingListOAuth2(request *CreatePlayingListOAuth2Request) (_result *CreatePlayingListOAuth2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePlayingListOAuth2Response{}
	_body, _err := client.CreatePlayingListOAuth2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a scheduled task.
//
// Description:
//
// Configure a triggering task for Custom instruction data in the ecosystem based on scheduling rules such as one-time, daily, or business days.
//
// @param tmpReq - CreateScheduleTaskRequest
//
// @param headers - CreateScheduleTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduleTaskResponse
func (client *Client) CreateScheduleTaskWithOptions(tmpReq *CreateScheduleTaskRequest, headers *CreateScheduleTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateScheduleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateScheduleTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduleTask"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/CreateScheduleTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a scheduled task.
//
// Description:
//
// Configure a triggering task for Custom instruction data in the ecosystem based on scheduling rules such as one-time, daily, or business days.
//
// @param request - CreateScheduleTaskRequest
//
// @return CreateScheduleTaskResponse
func (client *Client) CreateScheduleTask(request *CreateScheduleTaskRequest) (_result *CreateScheduleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateScheduleTaskHeaders{}
	_result = &CreateScheduleTaskResponse{}
	_body, _err := client.CreateScheduleTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete alarms. You can delete multiple alarms at the same time.
//
// @param tmpReq - DeleteAlarmsRequest
//
// @param headers - DeleteAlarmsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlarmsResponse
func (client *Client) DeleteAlarmsWithOptions(tmpReq *DeleteAlarmsRequest, headers *DeleteAlarmsHeaders, runtime *dara.RuntimeOptions) (_result *DeleteAlarmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAlarmsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlarms"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/deleteAlarms"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlarmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete alarms. You can delete multiple alarms at the same time.
//
// @param request - DeleteAlarmsRequest
//
// @return DeleteAlarmsResponse
func (client *Client) DeleteAlarms(request *DeleteAlarmsRequest) (_result *DeleteAlarmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteAlarmsHeaders{}
	_result = &DeleteAlarmsResponse{}
	_body, _err := client.DeleteAlarmsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a scheduled task.
//
// Description:
//
// Deletes a specified scheduled task by its ID.
//
// @param tmpReq - DeleteScheduleTaskRequest
//
// @param headers - DeleteScheduleTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduleTaskResponse
func (client *Client) DeleteScheduleTaskWithOptions(tmpReq *DeleteScheduleTaskRequest, headers *DeleteScheduleTaskHeaders, runtime *dara.RuntimeOptions) (_result *DeleteScheduleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteScheduleTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheduleTask"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/DeleteScheduleTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a scheduled task.
//
// Description:
//
// Deletes a specified scheduled task by its ID.
//
// @param request - DeleteScheduleTaskRequest
//
// @return DeleteScheduleTaskResponse
func (client *Client) DeleteScheduleTask(request *DeleteScheduleTaskRequest) (_result *DeleteScheduleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteScheduleTaskHeaders{}
	_result = &DeleteScheduleTaskResponse{}
	_body, _err := client.DeleteScheduleTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a subscribed album.
//
// @param request - DeleteSubRequest
//
// @param headers - DeleteSubHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubResponse
func (client *Client) DeleteSubWithOptions(request *DeleteSubRequest, headers *DeleteSubHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SubId) {
		query["SubId"] = request.SubId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSub"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/deleteSub"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a subscribed album.
//
// @param request - DeleteSubRequest
//
// @return DeleteSubResponse
func (client *Client) DeleteSub(request *DeleteSubRequest) (_result *DeleteSubResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteSubHeaders{}
	_result = &DeleteSubResponse{}
	_body, _err := client.DeleteSubWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Control property values such as device volume.
//
// Description:
//
// Specify a device to modify property values such as volume.
//
// @param tmpReq - DeviceControlRequest
//
// @param headers - DeviceControlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeviceControlResponse
func (client *Client) DeviceControlWithOptions(tmpReq *DeviceControlRequest, headers *DeviceControlHeaders, runtime *dara.RuntimeOptions) (_result *DeviceControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeviceControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ControlRequest) {
		request.ControlRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ControlRequest, dara.String("ControlRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ControlRequestShrink) {
		body["ControlRequest"] = request.ControlRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeviceControl"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/control"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeviceControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Control property values such as device volume.
//
// Description:
//
// Specify a device to modify property values such as volume.
//
// @param request - DeviceControlRequest
//
// @return DeviceControlResponse
func (client *Client) DeviceControl(request *DeviceControlRequest) (_result *DeviceControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeviceControlHeaders{}
	_result = &DeviceControlResponse{}
	_body, _err := client.DeviceControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Use the login-state access token for authentication to obtain the third-party user identifier and the Tmall Genie user\\"s openId.
//
// Description:
//
// Note: When the login state expires, the integrator must reinitiate the authorization login flow to obtain a new login state (entry service API: AuthLoginWithThirdUserInfo).
//
// - Invoke the service API with the login-state access credential (LoginStateAccessToken) parameter by using the Java SDK (for SDKs in other languages, consult Tmall Genie developers during integration).
//
// If the API call throws a TeaException and TeaException.code is InvalidAuthentication, the login state has expired.
//
// - Invoke the service API by sending an HTTP request with the login-state access credential (LoginStateAccessToken) parameter.
//
// If the HTTP response status code returned by the API call is 400, the login state has expired.
//
// @param request - EcologyOpennessAuthenticateRequest
//
// @param headers - EcologyOpennessAuthenticateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EcologyOpennessAuthenticateResponse
func (client *Client) EcologyOpennessAuthenticateWithOptions(request *EcologyOpennessAuthenticateRequest, headers *EcologyOpennessAuthenticateHeaders, runtime *dara.RuntimeOptions) (_result *EcologyOpennessAuthenticateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EncodeKey) {
		body["EncodeKey"] = request.EncodeKey
	}

	if !dara.IsNil(request.EncodeType) {
		body["EncodeType"] = request.EncodeType
	}

	if !dara.IsNil(request.LoginStateAccessToken) {
		body["LoginStateAccessToken"] = request.LoginStateAccessToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EcologyOpennessAuthenticate"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ecologyOpennessAuthenticate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EcologyOpennessAuthenticateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use the login-state access token for authentication to obtain the third-party user identifier and the Tmall Genie user\\"s openId.
//
// Description:
//
// Note: When the login state expires, the integrator must reinitiate the authorization login flow to obtain a new login state (entry service API: AuthLoginWithThirdUserInfo).
//
// - Invoke the service API with the login-state access credential (LoginStateAccessToken) parameter by using the Java SDK (for SDKs in other languages, consult Tmall Genie developers during integration).
//
// If the API call throws a TeaException and TeaException.code is InvalidAuthentication, the login state has expired.
//
// - Invoke the service API by sending an HTTP request with the login-state access credential (LoginStateAccessToken) parameter.
//
// If the HTTP response status code returned by the API call is 400, the login state has expired.
//
// @param request - EcologyOpennessAuthenticateRequest
//
// @return EcologyOpennessAuthenticateResponse
func (client *Client) EcologyOpennessAuthenticate(request *EcologyOpennessAuthenticateRequest) (_result *EcologyOpennessAuthenticateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &EcologyOpennessAuthenticateHeaders{}
	_result = &EcologyOpennessAuthenticateResponse{}
	_body, _err := client.EcologyOpennessAuthenticateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Ecosystem Open API for sending SMS verification codes
//
// @param request - EcologyOpennessSendVerificationCodeRequest
//
// @param headers - EcologyOpennessSendVerificationCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EcologyOpennessSendVerificationCodeResponse
func (client *Client) EcologyOpennessSendVerificationCodeWithOptions(request *EcologyOpennessSendVerificationCodeRequest, headers *EcologyOpennessSendVerificationCodeHeaders, runtime *dara.RuntimeOptions) (_result *EcologyOpennessSendVerificationCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PhoneNumber) {
		body["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EcologyOpennessSendVerificationCode"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ecologyOpennessSendVerificationCode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EcologyOpennessSendVerificationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Ecosystem Open API for sending SMS verification codes
//
// @param request - EcologyOpennessSendVerificationCodeRequest
//
// @return EcologyOpennessSendVerificationCodeResponse
func (client *Client) EcologyOpennessSendVerificationCode(request *EcologyOpennessSendVerificationCodeRequest) (_result *EcologyOpennessSendVerificationCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &EcologyOpennessSendVerificationCodeHeaders{}
	_result = &EcologyOpennessSendVerificationCodeResponse{}
	_body, _err := client.EcologyOpennessSendVerificationCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Find a list of accounts that can be authorized to log on by phone number
//
// @param request - FindUserlistToAuthLoginWithPhoneNumberRequest
//
// @param headers - FindUserlistToAuthLoginWithPhoneNumberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindUserlistToAuthLoginWithPhoneNumberResponse
func (client *Client) FindUserlistToAuthLoginWithPhoneNumberWithOptions(request *FindUserlistToAuthLoginWithPhoneNumberRequest, headers *FindUserlistToAuthLoginWithPhoneNumberHeaders, runtime *dara.RuntimeOptions) (_result *FindUserlistToAuthLoginWithPhoneNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindUserlistToAuthLoginWithPhoneNumber"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/findUserlistToAuthLoginWithPhoneNumber"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FindUserlistToAuthLoginWithPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Find a list of accounts that can be authorized to log on by phone number
//
// @param request - FindUserlistToAuthLoginWithPhoneNumberRequest
//
// @return FindUserlistToAuthLoginWithPhoneNumberResponse
func (client *Client) FindUserlistToAuthLoginWithPhoneNumber(request *FindUserlistToAuthLoginWithPhoneNumberRequest) (_result *FindUserlistToAuthLoginWithPhoneNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &FindUserlistToAuthLoginWithPhoneNumberHeaders{}
	_result = &FindUserlistToAuthLoginWithPhoneNumberResponse{}
	_body, _err := client.FindUserlistToAuthLoginWithPhoneNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the detailed information of a specific alarm clock.
//
// @param tmpReq - GetAlarmRequest
//
// @param headers - GetAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlarmResponse
func (client *Client) GetAlarmWithOptions(tmpReq *GetAlarmRequest, headers *GetAlarmHeaders, runtime *dara.RuntimeOptions) (_result *GetAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlarm"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getAlarm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the detailed information of a specific alarm clock.
//
// @param request - GetAlarmRequest
//
// @return GetAlarmResponse
func (client *Client) GetAlarm(request *GetAlarmRequest) (_result *GetAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAlarmHeaders{}
	_result = &GetAlarmResponse{}
	_body, _err := client.GetAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve album information by a specific ID.
//
// @param request - GetAlbumRequest
//
// @param headers - GetAlbumHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlbumResponse
func (client *Client) GetAlbumWithOptions(request *GetAlbumRequest, headers *GetAlbumHeaders, runtime *dara.RuntimeOptions) (_result *GetAlbumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlbum"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/GetAlbum"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlbumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve album information by a specific ID.
//
// @param request - GetAlbumRequest
//
// @return GetAlbumResponse
func (client *Client) GetAlbum(request *GetAlbumRequest) (_result *GetAlbumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAlbumHeaders{}
	_result = &GetAlbumResponse{}
	_body, _err := client.GetAlbumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve album information.
//
// @param request - GetAlbumDetailByIdRequest
//
// @param headers - GetAlbumDetailByIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlbumDetailByIdResponse
func (client *Client) GetAlbumDetailByIdWithOptions(request *GetAlbumDetailByIdRequest, headers *GetAlbumDetailByIdHeaders, runtime *dara.RuntimeOptions) (_result *GetAlbumDetailByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlbumId) {
		query["AlbumId"] = request.AlbumId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlbumDetailById"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getAlbumDetailById"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlbumDetailByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve album information.
//
// @param request - GetAlbumDetailByIdRequest
//
// @return GetAlbumDetailByIdResponse
func (client *Client) GetAlbumDetailById(request *GetAlbumDetailByIdRequest) (_result *GetAlbumDetailByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAlbumDetailByIdHeaders{}
	_result = &GetAlbumDetailByIdResponse{}
	_body, _err := client.GetAlbumDetailByIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtain the Aligenie account information bound to a third party by using the logon state credential
//
// Description:
//
// Note: When the logon state expires, the integrator must initiate the authorization logon flow again to obtain a new logon state (entry service API: AuthLoginWithThirdUserInfo).
//
// - Invoke the service API with the logon state access credential (LoginStateAccessToken) parameter by using the Java software development kit (SDK). For SDKs in other languages, consult the Tmall Genie developer team during integration.
//
// If the API call throws a TeaException and TeaException.code is InvalidAuthentication, the logon state has expired.
//
// - Invoke the service API by sending an HTTP request with the logon state access credential (LoginStateAccessToken) parameter.
//
// If the HTTP response status code returned by the API call is 400, the logon state has expired.
//
// @param request - GetAligenieUserInfoRequest
//
// @param headers - GetAligenieUserInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAligenieUserInfoResponse
func (client *Client) GetAligenieUserInfoWithOptions(request *GetAligenieUserInfoRequest, headers *GetAligenieUserInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetAligenieUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoginStateAccessToken) {
		query["LoginStateAccessToken"] = request.LoginStateAccessToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAligenieUserInfo"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getAligenieUserInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAligenieUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain the Aligenie account information bound to a third party by using the logon state credential
//
// Description:
//
// Note: When the logon state expires, the integrator must initiate the authorization logon flow again to obtain a new logon state (entry service API: AuthLoginWithThirdUserInfo).
//
// - Invoke the service API with the logon state access credential (LoginStateAccessToken) parameter by using the Java software development kit (SDK). For SDKs in other languages, consult the Tmall Genie developer team during integration.
//
// If the API call throws a TeaException and TeaException.code is InvalidAuthentication, the logon state has expired.
//
// - Invoke the service API by sending an HTTP request with the logon state access credential (LoginStateAccessToken) parameter.
//
// If the HTTP response status code returned by the API call is 400, the logon state has expired.
//
// @param request - GetAligenieUserInfoRequest
//
// @return GetAligenieUserInfoResponse
func (client *Client) GetAligenieUserInfo(request *GetAligenieUserInfoRequest) (_result *GetAligenieUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetAligenieUserInfoHeaders{}
	_result = &GetAligenieUserInfoResponse{}
	_body, _err := client.GetAligenieUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the authCode.
//
// Description:
//
// Obtain the corresponding authCode by specifying the user ID.
//
// @param tmpReq - GetCodeEnhanceRequest
//
// @param headers - GetCodeEnhanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCodeEnhanceResponse
func (client *Client) GetCodeEnhanceWithOptions(tmpReq *GetCodeEnhanceRequest, headers *GetCodeEnhanceHeaders, runtime *dara.RuntimeOptions) (_result *GetCodeEnhanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCodeEnhanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChannelInfo) {
		request.ChannelInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelInfo, dara.String("ChannelInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelInfoShrink) {
		query["ChannelInfo"] = request.ChannelInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCodeEnhance"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getCodeEnhance"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCodeEnhanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the authCode.
//
// Description:
//
// Obtain the corresponding authCode by specifying the user ID.
//
// @param request - GetCodeEnhanceRequest
//
// @return GetCodeEnhanceResponse
func (client *Client) GetCodeEnhance(request *GetCodeEnhanceRequest) (_result *GetCodeEnhanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetCodeEnhanceHeaders{}
	_result = &GetCodeEnhanceResponse{}
	_body, _err := client.GetCodeEnhanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve content information by a specific ID
//
// @param request - GetContentRequest
//
// @param headers - GetContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContentResponse
func (client *Client) GetContentWithOptions(request *GetContentRequest, headers *GetContentHeaders, runtime *dara.RuntimeOptions) (_result *GetContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContent"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/GetContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve content information by a specific ID
//
// @param request - GetContentRequest
//
// @return GetContentResponse
func (client *Client) GetContent(request *GetContentRequest) (_result *GetContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetContentHeaders{}
	_result = &GetContentResponse{}
	_body, _err := client.GetContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the current playback item.
//
// @param tmpReq - GetCurrentPlayingItemRequest
//
// @param headers - GetCurrentPlayingItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCurrentPlayingItemResponse
func (client *Client) GetCurrentPlayingItemWithOptions(tmpReq *GetCurrentPlayingItemRequest, headers *GetCurrentPlayingItemHeaders, runtime *dara.RuntimeOptions) (_result *GetCurrentPlayingItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCurrentPlayingItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCurrentPlayingItem"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/GetCurrentPlayingItem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCurrentPlayingItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the current playback item.
//
// @param request - GetCurrentPlayingItemRequest
//
// @return GetCurrentPlayingItemResponse
func (client *Client) GetCurrentPlayingItem(request *GetCurrentPlayingItemRequest) (_result *GetCurrentPlayingItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetCurrentPlayingItemHeaders{}
	_result = &GetCurrentPlayingItemResponse{}
	_body, _err := client.GetCurrentPlayingItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the current playlist.
//
// @param tmpReq - GetCurrentPlayingListRequest
//
// @param headers - GetCurrentPlayingListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCurrentPlayingListResponse
func (client *Client) GetCurrentPlayingListWithOptions(tmpReq *GetCurrentPlayingListRequest, headers *GetCurrentPlayingListHeaders, runtime *dara.RuntimeOptions) (_result *GetCurrentPlayingListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetCurrentPlayingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenQueryPlayListRequest) {
		request.OpenQueryPlayListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenQueryPlayListRequest, dara.String("OpenQueryPlayListRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenQueryPlayListRequestShrink) {
		body["OpenQueryPlayListRequest"] = request.OpenQueryPlayListRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCurrentPlayingList"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/GetCurrentPlayingList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCurrentPlayingListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the current playlist.
//
// @param request - GetCurrentPlayingListRequest
//
// @return GetCurrentPlayingListResponse
func (client *Client) GetCurrentPlayingList(request *GetCurrentPlayingListRequest) (_result *GetCurrentPlayingListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetCurrentPlayingListHeaders{}
	_result = &GetCurrentPlayingListResponse{}
	_body, _err := client.GetCurrentPlayingListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain basic device information by Device ID.
//
// @param tmpReq - GetDeviceBasicInfoRequest
//
// @param headers - GetDeviceBasicInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceBasicInfoResponse
func (client *Client) GetDeviceBasicInfoWithOptions(tmpReq *GetDeviceBasicInfoRequest, headers *GetDeviceBasicInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetDeviceBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDeviceBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceBasicInfo"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getDeviceBasicInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain basic device information by Device ID.
//
// @param request - GetDeviceBasicInfoRequest
//
// @return GetDeviceBasicInfoResponse
func (client *Client) GetDeviceBasicInfo(request *GetDeviceBasicInfoRequest) (_result *GetDeviceBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeviceBasicInfoHeaders{}
	_result = &GetDeviceBasicInfoResponse{}
	_body, _err := client.GetDeviceBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the Device ID based on the device MAC or SN.
//
// @param request - GetDeviceIdByIdentityRequest
//
// @param headers - GetDeviceIdByIdentityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceIdByIdentityResponse
func (client *Client) GetDeviceIdByIdentityWithOptions(request *GetDeviceIdByIdentityRequest, headers *GetDeviceIdByIdentityHeaders, runtime *dara.RuntimeOptions) (_result *GetDeviceIdByIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncodeKey) {
		query["EncodeKey"] = request.EncodeKey
	}

	if !dara.IsNil(request.EncodeType) {
		query["EncodeType"] = request.EncodeType
	}

	if !dara.IsNil(request.IdentityId) {
		query["IdentityId"] = request.IdentityId
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceIdByIdentity"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getDeviceIdByIdentity"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceIdByIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the Device ID based on the device MAC or SN.
//
// @param request - GetDeviceIdByIdentityRequest
//
// @return GetDeviceIdByIdentityResponse
func (client *Client) GetDeviceIdByIdentity(request *GetDeviceIdByIdentityRequest) (_result *GetDeviceIdByIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeviceIdByIdentityHeaders{}
	_result = &GetDeviceIdByIdentityResponse{}
	_body, _err := client.GetDeviceIdByIdentityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the user settings of a device.
//
// @param tmpReq - GetDeviceSettingRequest
//
// @param headers - GetDeviceSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceSettingResponse
func (client *Client) GetDeviceSettingWithOptions(tmpReq *GetDeviceSettingRequest, headers *GetDeviceSettingHeaders, runtime *dara.RuntimeOptions) (_result *GetDeviceSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDeviceSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Keys) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, dara.String("Keys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.KeysShrink) {
		query["Keys"] = request.KeysShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceSetting"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getDeviceSetting"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the user settings of a device.
//
// @param request - GetDeviceSettingRequest
//
// @return GetDeviceSettingResponse
func (client *Client) GetDeviceSetting(request *GetDeviceSettingRequest) (_result *GetDeviceSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeviceSettingHeaders{}
	_result = &GetDeviceSettingResponse{}
	_body, _err := client.GetDeviceSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain device status details.
//
// @param tmpReq - GetDeviceStatusDetailRequest
//
// @param headers - GetDeviceStatusDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceStatusDetailResponse
func (client *Client) GetDeviceStatusDetailWithOptions(tmpReq *GetDeviceStatusDetailRequest, headers *GetDeviceStatusDetailHeaders, runtime *dara.RuntimeOptions) (_result *GetDeviceStatusDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDeviceStatusDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Keys) {
		request.KeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keys, dara.String("Keys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.KeysShrink) {
		query["Keys"] = request.KeysShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceStatusDetail"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getDeviceStatusDetail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceStatusDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain device status details.
//
// @param request - GetDeviceStatusDetailRequest
//
// @return GetDeviceStatusDetailResponse
func (client *Client) GetDeviceStatusDetail(request *GetDeviceStatusDetailRequest) (_result *GetDeviceStatusDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeviceStatusDetailHeaders{}
	_result = &GetDeviceStatusDetailResponse{}
	_body, _err := client.GetDeviceStatusDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the online status of a device based on its openID.
//
// @param tmpReq - GetDeviceStatusInfoRequest
//
// @param headers - GetDeviceStatusInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceStatusInfoResponse
func (client *Client) GetDeviceStatusInfoWithOptions(tmpReq *GetDeviceStatusInfoRequest, headers *GetDeviceStatusInfoHeaders, runtime *dara.RuntimeOptions) (_result *GetDeviceStatusInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDeviceStatusInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceStatusInfo"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getDeviceStatusInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceStatusInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the online status of a device based on its openID.
//
// @param request - GetDeviceStatusInfoRequest
//
// @return GetDeviceStatusInfoResponse
func (client *Client) GetDeviceStatusInfo(request *GetDeviceStatusInfoRequest) (_result *GetDeviceStatusInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeviceStatusInfoHeaders{}
	_result = &GetDeviceStatusInfoResponse{}
	_body, _err := client.GetDeviceStatusInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the tag information of a device.
//
// @param tmpReq - GetDeviceTagRequest
//
// @param headers - GetDeviceTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceTagResponse
func (client *Client) GetDeviceTagWithOptions(tmpReq *GetDeviceTagRequest, headers *GetDeviceTagHeaders, runtime *dara.RuntimeOptions) (_result *GetDeviceTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDeviceTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceTag"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getDeviceTag"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the tag information of a device.
//
// @param request - GetDeviceTagRequest
//
// @return GetDeviceTagResponse
func (client *Client) GetDeviceTag(request *GetDeviceTagRequest) (_result *GetDeviceTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetDeviceTagHeaders{}
	_result = &GetDeviceTagResponse{}
	_body, _err := client.GetDeviceTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Derived requirements from Jiangsu Telecom\\"s self-developed desktop launcher project. HaoBai expects to collect statistics on the activation rate, active user rate, and top user usage patterns after the project is published, and to generate daily and weekly reports. Daily reports are synchronized automatically via API, while weekly reports are synchronized offline.
//
// @param request - GetJiangSuTelecomDataRequest
//
// @param headers - GetJiangSuTelecomDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJiangSuTelecomDataResponse
func (client *Client) GetJiangSuTelecomDataWithOptions(request *GetJiangSuTelecomDataRequest, headers *GetJiangSuTelecomDataHeaders, runtime *dara.RuntimeOptions) (_result *GetJiangSuTelecomDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJiangSuTelecomData"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/GetJiangSuTelecomData"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJiangSuTelecomDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Derived requirements from Jiangsu Telecom\\"s self-developed desktop launcher project. HaoBai expects to collect statistics on the activation rate, active user rate, and top user usage patterns after the project is published, and to generate daily and weekly reports. Daily reports are synchronized automatically via API, while weekly reports are synchronized offline.
//
// @param request - GetJiangSuTelecomDataRequest
//
// @return GetJiangSuTelecomDataResponse
func (client *Client) GetJiangSuTelecomData(request *GetJiangSuTelecomDataRequest) (_result *GetJiangSuTelecomDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetJiangSuTelecomDataHeaders{}
	_result = &GetJiangSuTelecomDataResponse{}
	_body, _err := client.GetJiangSuTelecomDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain scheduled task information.
//
// Description:
//
// Query the details of a specified scheduled task by its ID.
//
// @param tmpReq - GetScheduleTaskRequest
//
// @param headers - GetScheduleTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScheduleTaskResponse
func (client *Client) GetScheduleTaskWithOptions(tmpReq *GetScheduleTaskRequest, headers *GetScheduleTaskHeaders, runtime *dara.RuntimeOptions) (_result *GetScheduleTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetScheduleTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheduleTask"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/GetScheduleTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScheduleTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain scheduled task information.
//
// Description:
//
// Query the details of a specified scheduled task by its ID.
//
// @param request - GetScheduleTaskRequest
//
// @return GetScheduleTaskResponse
func (client *Client) GetScheduleTask(request *GetScheduleTaskRequest) (_result *GetScheduleTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetScheduleTaskHeaders{}
	_result = &GetScheduleTaskResponse{}
	_body, _err := client.GetScheduleTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the number of unread messages for a user.
//
// @param tmpReq - GetUnreadMessageCountRequest
//
// @param headers - GetUnreadMessageCountHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUnreadMessageCountResponse
func (client *Client) GetUnreadMessageCountWithOptions(tmpReq *GetUnreadMessageCountRequest, headers *GetUnreadMessageCountHeaders, runtime *dara.RuntimeOptions) (_result *GetUnreadMessageCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetUnreadMessageCountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUnreadMessageCount"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getUnreadMessageCount"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUnreadMessageCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the number of unread messages for a user.
//
// @param request - GetUnreadMessageCountRequest
//
// @return GetUnreadMessageCountResponse
func (client *Client) GetUnreadMessageCount(request *GetUnreadMessageCountRequest) (_result *GetUnreadMessageCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUnreadMessageCountHeaders{}
	_result = &GetUnreadMessageCountResponse{}
	_body, _err := client.GetUnreadMessageCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the User ID attached to a device by the device\\"s openID.
//
// @param tmpReq - GetUserByDeviceIdRequest
//
// @param headers - GetUserByDeviceIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserByDeviceIdResponse
func (client *Client) GetUserByDeviceIdWithOptions(tmpReq *GetUserByDeviceIdRequest, headers *GetUserByDeviceIdHeaders, runtime *dara.RuntimeOptions) (_result *GetUserByDeviceIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetUserByDeviceIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserByDeviceId"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/getUserByDeviceId"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserByDeviceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the User ID attached to a device by the device\\"s openID.
//
// @param request - GetUserByDeviceIdRequest
//
// @return GetUserByDeviceIdResponse
func (client *Client) GetUserByDeviceId(request *GetUserByDeviceIdRequest) (_result *GetUserByDeviceIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetUserByDeviceIdHeaders{}
	_result = &GetUserByDeviceIdResponse{}
	_body, _err := client.GetUserByDeviceIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the weather.
//
// Description:
//
// Query the weather data for the location of the device based on the Device Information obtained from system parameters.
//
// @param tmpReq - GetWeatherRequest
//
// @param headers - GetWeatherHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWeatherResponse
func (client *Client) GetWeatherWithOptions(tmpReq *GetWeatherRequest, headers *GetWeatherHeaders, runtime *dara.RuntimeOptions) (_result *GetWeatherResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetWeatherShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWeather"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/GetWeather"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWeatherResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the weather.
//
// Description:
//
// Query the weather data for the location of the device based on the Device Information obtained from system parameters.
//
// @param request - GetWeatherRequest
//
// @return GetWeatherResponse
func (client *Client) GetWeather(request *GetWeatherRequest) (_result *GetWeatherResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetWeatherHeaders{}
	_result = &GetWeatherResponse{}
	_body, _err := client.GetWeatherWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Click a playlist to start playback.
//
// @param tmpReq - IndexControlPlayingListRequest
//
// @param headers - IndexControlPlayingListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IndexControlPlayingListResponse
func (client *Client) IndexControlPlayingListWithOptions(tmpReq *IndexControlPlayingListRequest, headers *IndexControlPlayingListHeaders, runtime *dara.RuntimeOptions) (_result *IndexControlPlayingListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &IndexControlPlayingListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenIndexControlRequest) {
		request.OpenIndexControlRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenIndexControlRequest, dara.String("OpenIndexControlRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenIndexControlRequestShrink) {
		body["OpenIndexControlRequest"] = request.OpenIndexControlRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IndexControlPlayingList"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/IndexControlPlayingList"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IndexControlPlayingListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Click a playlist to start playback.
//
// @param request - IndexControlPlayingListRequest
//
// @return IndexControlPlayingListResponse
func (client *Client) IndexControlPlayingList(request *IndexControlPlayingListRequest) (_result *IndexControlPlayingListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &IndexControlPlayingListHeaders{}
	_result = &IndexControlPlayingListResponse{}
	_body, _err := client.IndexControlPlayingListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// When a user operation (not initiated from the device side) causes the login session of a third-party application on the device to become invalid (for example, the user logs on to the same account from a new device, invalidating the session on the old device, or the user changes their password, which invalidates the session on the old device), the third party must invoke this API to notify Tmall Genie.
//
// Description:
//
// > Modify message notification template
//
// > - You need to query the templateId first, then modify other parameters based on the templateId.
//
// @param tmpReq - InvalidateThirdPartyAppLoginStateRequest
//
// @param headers - InvalidateThirdPartyAppLoginStateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvalidateThirdPartyAppLoginStateResponse
func (client *Client) InvalidateThirdPartyAppLoginStateWithOptions(tmpReq *InvalidateThirdPartyAppLoginStateRequest, headers *InvalidateThirdPartyAppLoginStateHeaders, runtime *dara.RuntimeOptions) (_result *InvalidateThirdPartyAppLoginStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InvalidateThirdPartyAppLoginStateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.ThirdPartyAppId) {
		body["ThirdPartyAppId"] = request.ThirdPartyAppId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvalidateThirdPartyAppLoginState"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/invalidateThirdPartyAppLoginState"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvalidateThirdPartyAppLoginStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When a user operation (not initiated from the device side) causes the login session of a third-party application on the device to become invalid (for example, the user logs on to the same account from a new device, invalidating the session on the old device, or the user changes their password, which invalidates the session on the old device), the third party must invoke this API to notify Tmall Genie.
//
// Description:
//
// > Modify message notification template
//
// > - You need to query the templateId first, then modify other parameters based on the templateId.
//
// @param request - InvalidateThirdPartyAppLoginStateRequest
//
// @return InvalidateThirdPartyAppLoginStateResponse
func (client *Client) InvalidateThirdPartyAppLoginState(request *InvalidateThirdPartyAppLoginStateRequest) (_result *InvalidateThirdPartyAppLoginStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &InvalidateThirdPartyAppLoginStateHeaders{}
	_result = &InvalidateThirdPartyAppLoginStateResponse{}
	_body, _err := client.InvalidateThirdPartyAppLoginStateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Perform a paged query to retrieve the alarm clock list of a specific User on a device.
//
// @param tmpReq - ListAlarmsRequest
//
// @param headers - ListAlarmsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlarmsResponse
func (client *Client) ListAlarmsWithOptions(tmpReq *ListAlarmsRequest, headers *ListAlarmsHeaders, runtime *dara.RuntimeOptions) (_result *ListAlarmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlarmsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlarms"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listAlarm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlarmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Perform a paged query to retrieve the alarm clock list of a specific User on a device.
//
// @param request - ListAlarmsRequest
//
// @return ListAlarmsResponse
func (client *Client) ListAlarms(request *ListAlarmsRequest) (_result *ListAlarmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListAlarmsHeaders{}
	_result = &ListAlarmsResponse{}
	_body, _err := client.ListAlarmsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the content list within a music audio album.
//
// @param request - ListAlbumDetailRequest
//
// @param headers - ListAlbumDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlbumDetailResponse
func (client *Client) ListAlbumDetailWithOptions(request *ListAlbumDetailRequest, headers *ListAlbumDetailHeaders, runtime *dara.RuntimeOptions) (_result *ListAlbumDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlbumDetail"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ListAlbumDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlbumDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the content list within a music audio album.
//
// @param request - ListAlbumDetailRequest
//
// @return ListAlbumDetailResponse
func (client *Client) ListAlbumDetail(request *ListAlbumDetailRequest) (_result *ListAlbumDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListAlbumDetailHeaders{}
	_result = &ListAlbumDetailResponse{}
	_body, _err := client.ListAlbumDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Whether the album is subscribed.
//
// @param tmpReq - ListAlbumIsAddedRequest
//
// @param headers - ListAlbumIsAddedHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlbumIsAddedResponse
func (client *Client) ListAlbumIsAddedWithOptions(tmpReq *ListAlbumIsAddedRequest, headers *ListAlbumIsAddedHeaders, runtime *dara.RuntimeOptions) (_result *ListAlbumIsAddedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlbumIsAddedShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlbumIdList) {
		request.AlbumIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlbumIdList, dara.String("AlbumIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlbumIdListShrink) {
		query["AlbumIdList"] = request.AlbumIdListShrink
	}

	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlbumIsAdded"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listAlbumIsAdded"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlbumIsAddedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Whether the album is subscribed.
//
// @param request - ListAlbumIsAddedRequest
//
// @return ListAlbumIsAddedResponse
func (client *Client) ListAlbumIsAdded(request *ListAlbumIsAddedRequest) (_result *ListAlbumIsAddedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListAlbumIsAddedHeaders{}
	_result = &ListAlbumIsAddedResponse{}
	_body, _err := client.ListAlbumIsAddedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the content under a specific category according to the specified sorting order.
//
// @param tmpReq - ListCateContentRequest
//
// @param headers - ListCateContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCateContentResponse
func (client *Client) ListCateContentWithOptions(tmpReq *ListCateContentRequest, headers *ListCateContentHeaders, runtime *dara.RuntimeOptions) (_result *ListCateContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCateContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCateContent"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ListCateContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCateContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the content under a specific category according to the specified sorting order.
//
// @param request - ListCateContentRequest
//
// @return ListCateContentResponse
func (client *Client) ListCateContent(request *ListCateContentRequest) (_result *ListCateContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListCateContentHeaders{}
	_result = &ListCateContentResponse{}
	_body, _err := client.ListCateContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain category information.
//
// @param request - ListCateInfoRequest
//
// @param headers - ListCateInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCateInfoResponse
func (client *Client) ListCateInfoWithOptions(request *ListCateInfoRequest, headers *ListCateInfoHeaders, runtime *dara.RuntimeOptions) (_result *ListCateInfoResponse, _err error) {
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

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCateInfo"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ListCateInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain category information.
//
// @param request - ListCateInfoRequest
//
// @return ListCateInfoResponse
func (client *Client) ListCateInfo(request *ListCateInfoRequest) (_result *ListCateInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListCateInfoHeaders{}
	_result = &ListCateInfoResponse{}
	_body, _err := client.ListCateInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the top-level category list for music or audio.
//
// @param request - ListCommonCateFirstFloorRequest
//
// @param headers - ListCommonCateFirstFloorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommonCateFirstFloorResponse
func (client *Client) ListCommonCateFirstFloorWithOptions(request *ListCommonCateFirstFloorRequest, headers *ListCommonCateFirstFloorHeaders, runtime *dara.RuntimeOptions) (_result *ListCommonCateFirstFloorResponse, _err error) {
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

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCommonCateFirstFloor"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ListCommonCateFirstFloor"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCommonCateFirstFloorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the top-level category list for music or audio.
//
// @param request - ListCommonCateFirstFloorRequest
//
// @return ListCommonCateFirstFloorResponse
func (client *Client) ListCommonCateFirstFloor(request *ListCommonCateFirstFloorRequest) (_result *ListCommonCateFirstFloorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListCommonCateFirstFloorHeaders{}
	_result = &ListCommonCateFirstFloorResponse{}
	_body, _err := client.ListCommonCateFirstFloorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the list of level-2 categories under a specified level-1 category.
//
// @param request - ListCommonCateSecondFloorRequest
//
// @param headers - ListCommonCateSecondFloorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCommonCateSecondFloorResponse
func (client *Client) ListCommonCateSecondFloorWithOptions(request *ListCommonCateSecondFloorRequest, headers *ListCommonCateSecondFloorHeaders, runtime *dara.RuntimeOptions) (_result *ListCommonCateSecondFloorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParentCateId) {
		query["ParentCateId"] = request.ParentCateId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCommonCateSecondFloor"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ListCommonCateSecondFloor"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCommonCateSecondFloorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the list of level-2 categories under a specified level-1 category.
//
// @param request - ListCommonCateSecondFloorRequest
//
// @return ListCommonCateSecondFloorResponse
func (client *Client) ListCommonCateSecondFloor(request *ListCommonCateSecondFloorRequest) (_result *ListCommonCateSecondFloorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListCommonCateSecondFloorHeaders{}
	_result = &ListCommonCateSecondFloorResponse{}
	_body, _err := client.ListCommonCateSecondFloorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch retrieve basic device information based on a list of device openIDs.
//
// @param tmpReq - ListDeviceBasicInfoRequest
//
// @param headers - ListDeviceBasicInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceBasicInfoResponse
func (client *Client) ListDeviceBasicInfoWithOptions(tmpReq *ListDeviceBasicInfoRequest, headers *ListDeviceBasicInfoHeaders, runtime *dara.RuntimeOptions) (_result *ListDeviceBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDeviceBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfos) {
		request.DeviceInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfos, dara.String("DeviceInfos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfosShrink) {
		query["DeviceInfos"] = request.DeviceInfosShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceBasicInfo"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listDeviceBasicInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch retrieve basic device information based on a list of device openIDs.
//
// @param request - ListDeviceBasicInfoRequest
//
// @return ListDeviceBasicInfoResponse
func (client *Client) ListDeviceBasicInfo(request *ListDeviceBasicInfoRequest) (_result *ListDeviceBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDeviceBasicInfoHeaders{}
	_result = &ListDeviceBasicInfoResponse{}
	_body, _err := client.ListDeviceBasicInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the list of attached Device IDs based on the user\\"s openID.
//
// @param tmpReq - ListDeviceByUserIdRequest
//
// @param headers - ListDeviceByUserIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceByUserIdResponse
func (client *Client) ListDeviceByUserIdWithOptions(tmpReq *ListDeviceByUserIdRequest, headers *ListDeviceByUserIdHeaders, runtime *dara.RuntimeOptions) (_result *ListDeviceByUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDeviceByUserIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceByUserId"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listDeviceByUserId"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceByUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the list of attached Device IDs based on the user\\"s openID.
//
// @param request - ListDeviceByUserIdRequest
//
// @return ListDeviceByUserIdResponse
func (client *Client) ListDeviceByUserId(request *ListDeviceByUserIdRequest) (_result *ListDeviceByUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDeviceByUserIdHeaders{}
	_result = &ListDeviceByUserIdResponse{}
	_body, _err := client.ListDeviceByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the device list for a specified channel.
//
// @param tmpReq - ListDeviceByUserIdAndChanelRequest
//
// @param headers - ListDeviceByUserIdAndChanelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceByUserIdAndChanelResponse
func (client *Client) ListDeviceByUserIdAndChanelWithOptions(tmpReq *ListDeviceByUserIdAndChanelRequest, headers *ListDeviceByUserIdAndChanelHeaders, runtime *dara.RuntimeOptions) (_result *ListDeviceByUserIdAndChanelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDeviceByUserIdAndChanelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChannelInfo) {
		request.ChannelInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelInfo, dara.String("ChannelInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelInfoShrink) {
		query["ChannelInfo"] = request.ChannelInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceByUserIdAndChanel"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listDeviceByUserIdAndChanel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceByUserIdAndChanelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the device list for a specified channel.
//
// @param request - ListDeviceByUserIdAndChanelRequest
//
// @return ListDeviceByUserIdAndChanelResponse
func (client *Client) ListDeviceByUserIdAndChanel(request *ListDeviceByUserIdAndChanelRequest) (_result *ListDeviceByUserIdAndChanelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDeviceByUserIdAndChanelHeaders{}
	_result = &ListDeviceByUserIdAndChanelResponse{}
	_body, _err := client.ListDeviceByUserIdAndChanelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch retrieve Device IDs by device MAC/SN.
//
// @param tmpReq - ListDeviceIdByIdentitiesRequest
//
// @param headers - ListDeviceIdByIdentitiesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceIdByIdentitiesResponse
func (client *Client) ListDeviceIdByIdentitiesWithOptions(tmpReq *ListDeviceIdByIdentitiesRequest, headers *ListDeviceIdByIdentitiesHeaders, runtime *dara.RuntimeOptions) (_result *ListDeviceIdByIdentitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDeviceIdByIdentitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IdentityIds) {
		request.IdentityIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IdentityIds, dara.String("IdentityIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EncodeKey) {
		query["EncodeKey"] = request.EncodeKey
	}

	if !dara.IsNil(request.EncodeType) {
		query["EncodeType"] = request.EncodeType
	}

	if !dara.IsNil(request.IdentityIdsShrink) {
		query["IdentityIds"] = request.IdentityIdsShrink
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceIdByIdentities"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listDeviceIdByIdentities"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceIdByIdentitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch retrieve Device IDs by device MAC/SN.
//
// @param request - ListDeviceIdByIdentitiesRequest
//
// @return ListDeviceIdByIdentitiesResponse
func (client *Client) ListDeviceIdByIdentities(request *ListDeviceIdByIdentitiesRequest) (_result *ListDeviceIdByIdentitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListDeviceIdByIdentitiesHeaders{}
	_result = &ListDeviceIdByIdentitiesResponse{}
	_body, _err := client.ListDeviceIdByIdentitiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of ringtones under a specified music category.
//
// @param tmpReq - ListMusicRequest
//
// @param headers - ListMusicHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMusicResponse
func (client *Client) ListMusicWithOptions(tmpReq *ListMusicRequest, headers *ListMusicHeaders, runtime *dara.RuntimeOptions) (_result *ListMusicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMusicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMusic"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listMusic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMusicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of ringtones under a specified music category.
//
// @param request - ListMusicRequest
//
// @return ListMusicResponse
func (client *Client) ListMusic(request *ListMusicRequest) (_result *ListMusicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListMusicHeaders{}
	_result = &ListMusicResponse{}
	_body, _err := client.ListMusicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the history playback API.
//
// @param tmpReq - ListPlayHistoryRequest
//
// @param headers - ListPlayHistoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPlayHistoryResponse
func (client *Client) ListPlayHistoryWithOptions(tmpReq *ListPlayHistoryRequest, headers *ListPlayHistoryHeaders, runtime *dara.RuntimeOptions) (_result *ListPlayHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPlayHistoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPlayHistory"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ListPlayHistory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPlayHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the history playback API.
//
// @param request - ListPlayHistoryRequest
//
// @return ListPlayHistoryResponse
func (client *Client) ListPlayHistory(request *ListPlayHistoryRequest) (_result *ListPlayHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListPlayHistoryHeaders{}
	_result = &ListPlayHistoryResponse{}
	_body, _err := client.ListPlayHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain daily recommended music or audio.
//
// @param tmpReq - ListRecommendContentRequest
//
// @param headers - ListRecommendContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecommendContentResponse
func (client *Client) ListRecommendContentWithOptions(tmpReq *ListRecommendContentRequest, headers *ListRecommendContentHeaders, runtime *dara.RuntimeOptions) (_result *ListRecommendContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListRecommendContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecommendContent"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ListRecommendContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecommendContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain daily recommended music or audio.
//
// @param request - ListRecommendContentRequest
//
// @return ListRecommendContentResponse
func (client *Client) ListRecommendContent(request *ListRecommendContentRequest) (_result *ListRecommendContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListRecommendContentHeaders{}
	_result = &ListRecommendContentResponse{}
	_body, _err := client.ListRecommendContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The user has subscribed to the list.
//
// @param tmpReq - ListSubRequest
//
// @param headers - ListSubHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubResponse
func (client *Client) ListSubWithOptions(tmpReq *ListSubRequest, headers *ListSubHeaders, runtime *dara.RuntimeOptions) (_result *ListSubResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSubShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Page) {
		request.PageShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Page, dara.String("Page"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PageShrink) {
		query["Page"] = request.PageShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSub"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listSub"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The user has subscribed to the list.
//
// @param request - ListSubRequest
//
// @return ListSubResponse
func (client *Client) ListSub(request *ListSubRequest) (_result *ListSubResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListSubHeaders{}
	_result = &ListSubResponse{}
	_body, _err := client.ListSubWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Album List.
//
// @param tmpReq - ListSubAlbumRequest
//
// @param headers - ListSubAlbumHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubAlbumResponse
func (client *Client) ListSubAlbumWithOptions(tmpReq *ListSubAlbumRequest, headers *ListSubAlbumHeaders, runtime *dara.RuntimeOptions) (_result *ListSubAlbumResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSubAlbumShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QuerySubscriptionAlbumRequest) {
		request.QuerySubscriptionAlbumRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QuerySubscriptionAlbumRequest, dara.String("QuerySubscriptionAlbumRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.QuerySubscriptionAlbumRequestShrink) {
		query["QuerySubscriptionAlbumRequest"] = request.QuerySubscriptionAlbumRequestShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubAlbum"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listSubAlbum"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubAlbumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Album List.
//
// @param request - ListSubAlbumRequest
//
// @return ListSubAlbumResponse
func (client *Client) ListSubAlbum(request *ListSubAlbumRequest) (_result *ListSubAlbumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListSubAlbumHeaders{}
	_result = &ListSubAlbumResponse{}
	_body, _err := client.ListSubAlbumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Album categorization.
//
// @param request - ListSubscriptionAlbumCategoryRequest
//
// @param headers - ListSubscriptionAlbumCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubscriptionAlbumCategoryResponse
func (client *Client) ListSubscriptionAlbumCategoryWithOptions(request *ListSubscriptionAlbumCategoryRequest, headers *ListSubscriptionAlbumCategoryHeaders, runtime *dara.RuntimeOptions) (_result *ListSubscriptionAlbumCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryName) {
		query["CategoryName"] = request.CategoryName
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubscriptionAlbumCategory"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listSubscriptionAlbumCategory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubscriptionAlbumCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Album categorization.
//
// @param request - ListSubscriptionAlbumCategoryRequest
//
// @return ListSubscriptionAlbumCategoryResponse
func (client *Client) ListSubscriptionAlbumCategory(request *ListSubscriptionAlbumCategoryRequest) (_result *ListSubscriptionAlbumCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListSubscriptionAlbumCategoryHeaders{}
	_result = &ListSubscriptionAlbumCategoryResponse{}
	_body, _err := client.ListSubscriptionAlbumCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of user message records based on the query conditions, including the start time and the number of entries to return.
//
// @param tmpReq - ListUserMessageRequest
//
// @param headers - ListUserMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserMessageResponse
func (client *Client) ListUserMessageWithOptions(tmpReq *ListUserMessageRequest, headers *ListUserMessageHeaders, runtime *dara.RuntimeOptions) (_result *ListUserMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListUserMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BeforeTime) {
		query["BeforeTime"] = request.BeforeTime
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserMessage"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/listUserMessage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of user message records based on the query conditions, including the start time and the number of entries to return.
//
// @param request - ListUserMessageRequest
//
// @return ListUserMessageResponse
func (client *Client) ListUserMessage(request *ListUserMessageRequest) (_result *ListUserMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListUserMessageHeaders{}
	_result = &ListUserMessageResponse{}
	_body, _err := client.ListUserMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Shift the project to obtain recommended songs, including daily recommendations, popular songs, and genre-based playlists.
//
// Description:
//
// Different scenarios (such as daily recommendations and popular songs) can be distinguished by the Type parameter.
//
// @param tmpReq - MobileRecommendRequest
//
// @param headers - MobileRecommendHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MobileRecommendResponse
func (client *Client) MobileRecommendWithOptions(tmpReq *MobileRecommendRequest, headers *MobileRecommendHeaders, runtime *dara.RuntimeOptions) (_result *MobileRecommendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MobileRecommendShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		query["BotId"] = request.BotId
	}

	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.Style) {
		query["Style"] = request.Style
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MobileRecommend"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/mobile/recommend/music"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MobileRecommendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shift the project to obtain recommended songs, including daily recommendations, popular songs, and genre-based playlists.
//
// Description:
//
// Different scenarios (such as daily recommendations and popular songs) can be distinguished by the Type parameter.
//
// @param request - MobileRecommendRequest
//
// @return MobileRecommendResponse
func (client *Client) MobileRecommend(request *MobileRecommendRequest) (_result *MobileRecommendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &MobileRecommendHeaders{}
	_result = &MobileRecommendResponse{}
	_body, _err := client.MobileRecommendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Playback and pause control.
//
// @param tmpReq - PlayAndPauseControlRequest
//
// @param headers - PlayAndPauseControlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PlayAndPauseControlResponse
func (client *Client) PlayAndPauseControlWithOptions(tmpReq *PlayAndPauseControlRequest, headers *PlayAndPauseControlHeaders, runtime *dara.RuntimeOptions) (_result *PlayAndPauseControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PlayAndPauseControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenPlayAndPauseControlParam) {
		request.OpenPlayAndPauseControlParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenPlayAndPauseControlParam, dara.String("OpenPlayAndPauseControlParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenPlayAndPauseControlParamShrink) {
		body["OpenPlayAndPauseControlParam"] = request.OpenPlayAndPauseControlParamShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PlayAndPauseControl"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/PlayAndPauseControl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PlayAndPauseControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Playback and pause control.
//
// @param request - PlayAndPauseControlRequest
//
// @return PlayAndPauseControlResponse
func (client *Client) PlayAndPauseControl(request *PlayAndPauseControlRequest) (_result *PlayAndPauseControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PlayAndPauseControlHeaders{}
	_result = &PlayAndPauseControlResponse{}
	_body, _err := client.PlayAndPauseControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switch playback mode.
//
// @param tmpReq - PlayModeControlRequest
//
// @param headers - PlayModeControlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PlayModeControlResponse
func (client *Client) PlayModeControlWithOptions(tmpReq *PlayModeControlRequest, headers *PlayModeControlHeaders, runtime *dara.RuntimeOptions) (_result *PlayModeControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PlayModeControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenPlayModeControlRequest) {
		request.OpenPlayModeControlRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenPlayModeControlRequest, dara.String("OpenPlayModeControlRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenPlayModeControlRequestShrink) {
		body["OpenPlayModeControlRequest"] = request.OpenPlayModeControlRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PlayModeControl"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/PlayModeControl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PlayModeControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switch playback mode.
//
// @param request - PlayModeControlRequest
//
// @return PlayModeControlResponse
func (client *Client) PlayModeControl(request *PlayModeControlRequest) (_result *PlayModeControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PlayModeControlHeaders{}
	_result = &PlayModeControlResponse{}
	_body, _err := client.PlayModeControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Control the next or previous track in the created playlist.
//
// @param tmpReq - PreviousAndNextControlRequest
//
// @param headers - PreviousAndNextControlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviousAndNextControlResponse
func (client *Client) PreviousAndNextControlWithOptions(tmpReq *PreviousAndNextControlRequest, headers *PreviousAndNextControlHeaders, runtime *dara.RuntimeOptions) (_result *PreviousAndNextControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PreviousAndNextControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenControlPlayingListRequest) {
		request.OpenControlPlayingListRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenControlPlayingListRequest, dara.String("OpenControlPlayingListRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenControlPlayingListRequestShrink) {
		body["OpenControlPlayingListRequest"] = request.OpenControlPlayingListRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviousAndNextControl"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/PreviousAndNextControl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviousAndNextControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Control the next or previous track in the created playlist.
//
// @param request - PreviousAndNextControlRequest
//
// @return PreviousAndNextControlResponse
func (client *Client) PreviousAndNextControl(request *PreviousAndNextControlRequest) (_result *PreviousAndNextControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PreviousAndNextControlHeaders{}
	_result = &PreviousAndNextControlResponse{}
	_body, _err := client.PreviousAndNextControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Progress control.
//
// @param tmpReq - ProgressControlRequest
//
// @param headers - ProgressControlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProgressControlResponse
func (client *Client) ProgressControlWithOptions(tmpReq *ProgressControlRequest, headers *ProgressControlHeaders, runtime *dara.RuntimeOptions) (_result *ProgressControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ProgressControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OpenProgressControlRequest) {
		request.OpenProgressControlRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OpenProgressControlRequest, dara.String("OpenProgressControlRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenProgressControlRequestShrink) {
		body["OpenProgressControlRequest"] = request.OpenProgressControlRequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProgressControl"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/ProgressControl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProgressControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Progress control.
//
// @param request - ProgressControlRequest
//
// @return ProgressControlResponse
func (client *Client) ProgressControl(request *ProgressControlRequest) (_result *ProgressControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ProgressControlHeaders{}
	_result = &ProgressControlResponse{}
	_body, _err := client.ProgressControlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the collection of alarm music types supported by the device.
//
// @param tmpReq - QueryMusicTypeRequest
//
// @param headers - QueryMusicTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMusicTypeResponse
func (client *Client) QueryMusicTypeWithOptions(tmpReq *QueryMusicTypeRequest, headers *QueryMusicTypeHeaders, runtime *dara.RuntimeOptions) (_result *QueryMusicTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QueryMusicTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMusicType"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/queryMusicType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMusicTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the collection of alarm music types supported by the device.
//
// @param request - QueryMusicTypeRequest
//
// @return QueryMusicTypeResponse
func (client *Client) QueryMusicType(request *QueryMusicTypeRequest) (_result *QueryMusicTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryMusicTypeHeaders{}
	_result = &QueryMusicTypeResponse{}
	_body, _err := client.QueryMusicTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtain the list of authorized Tmall Genie users and their devices by using the TME user ID, for use by KuGou Cloud Push Song service
//
// @param request - QueryUserDeviceListByTmeUserIdRequest
//
// @param headers - QueryUserDeviceListByTmeUserIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserDeviceListByTmeUserIdResponse
func (client *Client) QueryUserDeviceListByTmeUserIdWithOptions(request *QueryUserDeviceListByTmeUserIdRequest, headers *QueryUserDeviceListByTmeUserIdHeaders, runtime *dara.RuntimeOptions) (_result *QueryUserDeviceListByTmeUserIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Sp) {
		query["Sp"] = request.Sp
	}

	if !dara.IsNil(request.TmeUserId) {
		query["TmeUserId"] = request.TmeUserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserDeviceListByTmeUserId"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/queryUserDeviceListByTmeUserId"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserDeviceListByTmeUserIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain the list of authorized Tmall Genie users and their devices by using the TME user ID, for use by KuGou Cloud Push Song service
//
// @param request - QueryUserDeviceListByTmeUserIdRequest
//
// @return QueryUserDeviceListByTmeUserIdResponse
func (client *Client) QueryUserDeviceListByTmeUserId(request *QueryUserDeviceListByTmeUserIdRequest) (_result *QueryUserDeviceListByTmeUserIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &QueryUserDeviceListByTmeUserIdHeaders{}
	_result = &QueryUserDeviceListByTmeUserIdResponse{}
	_body, _err := client.QueryUserDeviceListByTmeUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Set a message to read status.
//
// @param tmpReq - ReadMessageRequest
//
// @param headers - ReadMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageResponse
func (client *Client) ReadMessageWithOptions(tmpReq *ReadMessageRequest, headers *ReadMessageHeaders, runtime *dara.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReadMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessage"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/readMessage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Set a message to read status.
//
// @param request - ReadMessageRequest
//
// @return ReadMessageResponse
func (client *Client) ReadMessage(request *ReadMessageRequest) (_result *ReadMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ReadMessageHeaders{}
	_result = &ReadMessageResponse{}
	_body, _err := client.ReadMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A screen-equipped device completes activation and attachment between the user and the device by scanning a QR code.
//
// Description:
//
// Perform activation and attachment by scanning the QR code related to activation and attachment displayed on the device.
//
// @param tmpReq - ScanCodeBindRequest
//
// @param headers - ScanCodeBindHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanCodeBindResponse
func (client *Client) ScanCodeBindWithOptions(tmpReq *ScanCodeBindRequest, headers *ScanCodeBindHeaders, runtime *dara.RuntimeOptions) (_result *ScanCodeBindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ScanCodeBindShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BindReq) {
		request.BindReqShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BindReq, dara.String("BindReq"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BindReqShrink) {
		body["BindReq"] = request.BindReqShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScanCodeBind"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/scanCode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ScanCodeBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A screen-equipped device completes activation and attachment between the user and the device by scanning a QR code.
//
// Description:
//
// Perform activation and attachment by scanning the QR code related to activation and attachment displayed on the device.
//
// @param request - ScanCodeBindRequest
//
// @return ScanCodeBindResponse
func (client *Client) ScanCodeBind(request *ScanCodeBindRequest) (_result *ScanCodeBindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ScanCodeBindHeaders{}
	_result = &ScanCodeBindResponse{}
	_body, _err := client.ScanCodeBindWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the content of a selection pool by its ID.
//
// @param tmpReq - ScgSearchRequest
//
// @param headers - ScgSearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScgSearchResponse
func (client *Client) ScgSearchWithOptions(tmpReq *ScgSearchRequest, headers *ScgSearchHeaders, runtime *dara.RuntimeOptions) (_result *ScgSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ScgSearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ScgFilter) {
		request.ScgFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScgFilter, dara.String("ScgFilter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ScgFilterShrink) {
		query["ScgFilter"] = request.ScgFilterShrink
	}

	if !dara.IsNil(request.TopicId) {
		query["TopicId"] = request.TopicId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScgSearch"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/scgSearch"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScgSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the content of a selection pool by its ID.
//
// @param request - ScgSearchRequest
//
// @return ScgSearchResponse
func (client *Client) ScgSearch(request *ScgSearchRequest) (_result *ScgSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ScgSearchHeaders{}
	_result = &ScgSearchResponse{}
	_body, _err := client.ScgSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Search according to specific search criteria.
//
// @param tmpReq - SearchContentRequest
//
// @param headers - SearchContentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchContentResponse
func (client *Client) SearchContentWithOptions(tmpReq *SearchContentRequest, headers *SearchContentHeaders, runtime *dara.RuntimeOptions) (_result *SearchContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Request) {
		request.RequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Request, dara.String("Request"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestShrink) {
		body["Request"] = request.RequestShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchContent"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/SearchContent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search according to specific search criteria.
//
// @param request - SearchContentRequest
//
// @return SearchContentResponse
func (client *Client) SearchContent(request *SearchContentRequest) (_result *SearchContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SearchContentHeaders{}
	_result = &SearchContentResponse{}
	_body, _err := client.SearchContentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Send a message.
//
// @param tmpReq - SendMessageRequest
//
// @param headers - SendMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithOptions(tmpReq *SendMessageRequest, headers *SendMessageHeaders, runtime *dara.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	if !dara.IsNil(request.UserInfoShrink) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessage"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/sendMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Send a message.
//
// @param request - SendMessageRequest
//
// @return SendMessageResponse
func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SendMessageHeaders{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the user settings of a device.
//
// @param tmpReq - SetDeviceSettingRequest
//
// @param headers - SetDeviceSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeviceSettingResponse
func (client *Client) SetDeviceSettingWithOptions(tmpReq *SetDeviceSettingRequest, headers *SetDeviceSettingHeaders, runtime *dara.RuntimeOptions) (_result *SetDeviceSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetDeviceSettingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		body["Key"] = request.Key
	}

	if !dara.IsNil(request.Value) {
		body["Value"] = request.Value
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDeviceSetting"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/setDeviceSetting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDeviceSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the user settings of a device.
//
// @param request - SetDeviceSettingRequest
//
// @return SetDeviceSettingResponse
func (client *Client) SetDeviceSetting(request *SetDeviceSettingRequest) (_result *SetDeviceSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SetDeviceSettingHeaders{}
	_result = &SetDeviceSettingResponse{}
	_body, _err := client.SetDeviceSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Third-party real-time information data change event push
//
// @param request - ThirdImmediateMsgPushRequest
//
// @param headers - ThirdImmediateMsgPushHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ThirdImmediateMsgPushResponse
func (client *Client) ThirdImmediateMsgPushWithOptions(request *ThirdImmediateMsgPushRequest, headers *ThirdImmediateMsgPushHeaders, runtime *dara.RuntimeOptions) (_result *ThirdImmediateMsgPushResponse, _err error) {
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

	if !dara.IsNil(request.ChangeDetail) {
		query["ChangeDetail"] = request.ChangeDetail
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PsgIds) {
		query["PsgIds"] = request.PsgIds
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TrafficChangeType) {
		query["TrafficChangeType"] = request.TrafficChangeType
	}

	if !dara.IsNil(request.TrafficChangeTypeDesc) {
		query["TrafficChangeTypeDesc"] = request.TrafficChangeTypeDesc
	}

	if !dara.IsNil(request.TrafficJourneyIds) {
		query["TrafficJourneyIds"] = request.TrafficJourneyIds
	}

	if !dara.IsNil(request.TrafficSubOrderIds) {
		query["TrafficSubOrderIds"] = request.TrafficSubOrderIds
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ThirdImmediateMsgPush"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/thirdImmediateMsgPush"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ThirdImmediateMsgPushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Third-party real-time information data change event push
//
// @param request - ThirdImmediateMsgPushRequest
//
// @return ThirdImmediateMsgPushResponse
func (client *Client) ThirdImmediateMsgPush(request *ThirdImmediateMsgPushRequest) (_result *ThirdImmediateMsgPushResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ThirdImmediateMsgPushHeaders{}
	_result = &ThirdImmediateMsgPushResponse{}
	_body, _err := client.ThirdImmediateMsgPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detach the relationship between a third-party account and a Tmall Genie account by using a logon credential.
//
// Description:
//
// Note: When the logon credential expires, the integrator must initiate the authorization logon flow again to obtain a new logon credential (entry service API: AuthLoginWithThirdUserInfo).
//
// - Invoke the service API with the logon credential (LoginStateAccessToken) parameter via the Java software development kit (SDK). For SDKs in other languages, consult Tmall Genie developers during integration.
//
// If the API call throws a TeaException and TeaException.code is "InvalidAuthentication", the logon credential has expired.
//
// - Invoke the service API by sending an HTTP request with the logon credential (LoginStateAccessToken) parameter.
//
// If the HTTP response status code returned by the API call is 400, the logon credential has expired.
//
// @param request - UnbindAligenieUserRequest
//
// @param headers - UnbindAligenieUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAligenieUserResponse
func (client *Client) UnbindAligenieUserWithOptions(request *UnbindAligenieUserRequest, headers *UnbindAligenieUserHeaders, runtime *dara.RuntimeOptions) (_result *UnbindAligenieUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LoginStateAccessToken) {
		body["LoginStateAccessToken"] = request.LoginStateAccessToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindAligenieUser"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/unbindAligenieUser"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAligenieUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach the relationship between a third-party account and a Tmall Genie account by using a logon credential.
//
// Description:
//
// Note: When the logon credential expires, the integrator must initiate the authorization logon flow again to obtain a new logon credential (entry service API: AuthLoginWithThirdUserInfo).
//
// - Invoke the service API with the logon credential (LoginStateAccessToken) parameter via the Java software development kit (SDK). For SDKs in other languages, consult Tmall Genie developers during integration.
//
// If the API call throws a TeaException and TeaException.code is "InvalidAuthentication", the logon credential has expired.
//
// - Invoke the service API by sending an HTTP request with the logon credential (LoginStateAccessToken) parameter.
//
// If the HTTP response status code returned by the API call is 400, the logon credential has expired.
//
// @param request - UnbindAligenieUserRequest
//
// @return UnbindAligenieUserResponse
func (client *Client) UnbindAligenieUser(request *UnbindAligenieUserRequest) (_result *UnbindAligenieUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UnbindAligenieUserHeaders{}
	_result = &UnbindAligenieUserResponse{}
	_body, _err := client.UnbindAligenieUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detach the binding relationship based on the user\\"s open information and the device\\"s open information.
//
// @param tmpReq - UnbindDeviceRequest
//
// @param headers - UnbindDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDeviceResponse
func (client *Client) UnbindDeviceWithOptions(tmpReq *UnbindDeviceRequest, headers *UnbindDeviceHeaders, runtime *dara.RuntimeOptions) (_result *UnbindDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UnbindDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindDevice"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/unbindDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach the binding relationship based on the user\\"s open information and the device\\"s open information.
//
// @param request - UnbindDeviceRequest
//
// @return UnbindDeviceResponse
func (client *Client) UnbindDevice(request *UnbindDeviceRequest) (_result *UnbindDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UnbindDeviceHeaders{}
	_result = &UnbindDeviceResponse{}
	_body, _err := client.UnbindDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update an alarm, including its trigger time, loop type, and music.
//
// @param tmpReq - UpdateAlarmRequest
//
// @param headers - UpdateAlarmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlarmResponse
func (client *Client) UpdateAlarmWithOptions(tmpReq *UpdateAlarmRequest, headers *UpdateAlarmHeaders, runtime *dara.RuntimeOptions) (_result *UpdateAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAlarmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeviceInfo) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, dara.String("DeviceInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Payload) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, dara.String("Payload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserInfo) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, dara.String("UserInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceInfoShrink) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !dara.IsNil(request.PayloadShrink) {
		body["Payload"] = request.PayloadShrink
	}

	if !dara.IsNil(request.UserInfoShrink) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAligenieAccessToken) {
		realHeaders["x-acs-aligenie-access-token"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAligenieAccessToken)))
	}

	if !dara.IsNil(headers.Authorization) {
		realHeaders["Authorization"] = dara.String(dara.ToString(dara.StringValue(headers.Authorization)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlarm"),
		Version:     dara.String("ssp_1.0"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1.0/ssp/updateAlarm"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlarmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update an alarm, including its trigger time, loop type, and music.
//
// @param request - UpdateAlarmRequest
//
// @return UpdateAlarmResponse
func (client *Client) UpdateAlarm(request *UpdateAlarmRequest) (_result *UpdateAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateAlarmHeaders{}
	_result = &UpdateAlarmResponse{}
	_body, _err := client.UpdateAlarmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
