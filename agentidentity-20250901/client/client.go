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
	client.Endpoint, _err = client.GetEndpoint(dara.String("agentidentity"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建一个 API 密钥凭证提供商
//
// @param request - CreateAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAPIKeyCredentialProviderResponse
func (client *Client) CreateAPIKeyCredentialProviderWithOptions(request *CreateAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateAPIKeyCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		body["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.APIKeyCredentialProviderName) {
		body["APIKeyCredentialProviderName"] = request.APIKeyCredentialProviderName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAPIKeyCredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAPIKeyCredentialProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个 API 密钥凭证提供商
//
// @param request - CreateAPIKeyCredentialProviderRequest
//
// @return CreateAPIKeyCredentialProviderResponse
func (client *Client) CreateAPIKeyCredentialProvider(request *CreateAPIKeyCredentialProviderRequest) (_result *CreateAPIKeyCredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAPIKeyCredentialProviderResponse{}
	_body, _err := client.CreateAPIKeyCredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建IdentityProvider
//
// @param tmpReq - CreateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityProviderResponse
func (client *Client) CreateIdentityProviderWithOptions(tmpReq *CreateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateIdentityProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedAudience) {
		request.AllowedAudienceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedAudience, dara.String("AllowedAudience"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedAudienceShrink) {
		body["AllowedAudience"] = request.AllowedAudienceShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DiscoveryURL) {
		body["DiscoveryURL"] = request.DiscoveryURL
	}

	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建IdentityProvider
//
// @param request - CreateIdentityProviderRequest
//
// @return CreateIdentityProviderResponse
func (client *Client) CreateIdentityProvider(request *CreateIdentityProviderRequest) (_result *CreateIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIdentityProviderResponse{}
	_body, _err := client.CreateIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建 OAuth2 凭证提供商
//
// @param tmpReq - CreateOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOAuth2CredentialProviderResponse
func (client *Client) CreateOAuth2CredentialProviderWithOptions(tmpReq *CreateOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *CreateOAuth2CredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOAuth2CredentialProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OAuth2ProviderConfig) {
		request.OAuth2ProviderConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OAuth2ProviderConfig, dara.String("OAuth2ProviderConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackURL) {
		body["CallbackURL"] = request.CallbackURL
	}

	if !dara.IsNil(request.CredentialProviderVendor) {
		body["CredentialProviderVendor"] = request.CredentialProviderVendor
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.OAuth2CredentialProviderName) {
		body["OAuth2CredentialProviderName"] = request.OAuth2CredentialProviderName
	}

	if !dara.IsNil(request.OAuth2ProviderConfigShrink) {
		body["OAuth2ProviderConfig"] = request.OAuth2ProviderConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOAuth2CredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOAuth2CredentialProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 OAuth2 凭证提供商
//
// @param request - CreateOAuth2CredentialProviderRequest
//
// @return CreateOAuth2CredentialProviderResponse
func (client *Client) CreateOAuth2CredentialProvider(request *CreateOAuth2CredentialProviderRequest) (_result *CreateOAuth2CredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOAuth2CredentialProviderResponse{}
	_body, _err := client.CreateOAuth2CredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建WorkloadIdentity
//
// @param tmpReq - CreateWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkloadIdentityResponse
func (client *Client) CreateWorkloadIdentityWithOptions(tmpReq *CreateWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkloadIdentityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedResourceOAuth2ReturnURLs) {
		request.AllowedResourceOAuth2ReturnURLsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedResourceOAuth2ReturnURLs, dara.String("AllowedResourceOAuth2ReturnURLs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedResourceOAuth2ReturnURLsShrink) {
		body["AllowedResourceOAuth2ReturnURLs"] = request.AllowedResourceOAuth2ReturnURLsShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkloadIdentity"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkloadIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建WorkloadIdentity
//
// @param request - CreateWorkloadIdentityRequest
//
// @return CreateWorkloadIdentityResponse
func (client *Client) CreateWorkloadIdentity(request *CreateWorkloadIdentityRequest) (_result *CreateWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWorkloadIdentityResponse{}
	_body, _err := client.CreateWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除一个 API 密钥凭证提供商
//
// @param request - DeleteAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAPIKeyCredentialProviderResponse
func (client *Client) DeleteAPIKeyCredentialProviderWithOptions(request *DeleteAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteAPIKeyCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.APIKeyCredentialProviderName) {
		body["APIKeyCredentialProviderName"] = request.APIKeyCredentialProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAPIKeyCredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAPIKeyCredentialProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个 API 密钥凭证提供商
//
// @param request - DeleteAPIKeyCredentialProviderRequest
//
// @return DeleteAPIKeyCredentialProviderResponse
func (client *Client) DeleteAPIKeyCredentialProvider(request *DeleteAPIKeyCredentialProviderRequest) (_result *DeleteAPIKeyCredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAPIKeyCredentialProviderResponse{}
	_body, _err := client.DeleteAPIKeyCredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除IdentityProvider
//
// @param request - DeleteIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdentityProviderResponse
func (client *Client) DeleteIdentityProviderWithOptions(request *DeleteIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除IdentityProvider
//
// @param request - DeleteIdentityProviderRequest
//
// @return DeleteIdentityProviderResponse
func (client *Client) DeleteIdentityProvider(request *DeleteIdentityProviderRequest) (_result *DeleteIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIdentityProviderResponse{}
	_body, _err := client.DeleteIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 OAuth2 凭证提供商
//
// @param request - DeleteOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOAuth2CredentialProviderResponse
func (client *Client) DeleteOAuth2CredentialProviderWithOptions(request *DeleteOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *DeleteOAuth2CredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OAuth2CredentialProviderName) {
		body["OAuth2CredentialProviderName"] = request.OAuth2CredentialProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOAuth2CredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOAuth2CredentialProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 OAuth2 凭证提供商
//
// @param request - DeleteOAuth2CredentialProviderRequest
//
// @return DeleteOAuth2CredentialProviderResponse
func (client *Client) DeleteOAuth2CredentialProvider(request *DeleteOAuth2CredentialProviderRequest) (_result *DeleteOAuth2CredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOAuth2CredentialProviderResponse{}
	_body, _err := client.DeleteOAuth2CredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除WorkloadIdentity
//
// @param request - DeleteWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkloadIdentityResponse
func (client *Client) DeleteWorkloadIdentityWithOptions(request *DeleteWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkloadIdentity"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkloadIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除WorkloadIdentity
//
// @param request - DeleteWorkloadIdentityRequest
//
// @return DeleteWorkloadIdentityResponse
func (client *Client) DeleteWorkloadIdentity(request *DeleteWorkloadIdentityRequest) (_result *DeleteWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWorkloadIdentityResponse{}
	_body, _err := client.DeleteWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一个 API 密钥凭证提供商
//
// @param request - GetAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAPIKeyCredentialProviderResponse
func (client *Client) GetAPIKeyCredentialProviderWithOptions(request *GetAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *GetAPIKeyCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.APIKeyCredentialProviderName) {
		body["APIKeyCredentialProviderName"] = request.APIKeyCredentialProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAPIKeyCredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAPIKeyCredentialProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一个 API 密钥凭证提供商
//
// @param request - GetAPIKeyCredentialProviderRequest
//
// @return GetAPIKeyCredentialProviderResponse
func (client *Client) GetAPIKeyCredentialProvider(request *GetAPIKeyCredentialProviderRequest) (_result *GetAPIKeyCredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAPIKeyCredentialProviderResponse{}
	_body, _err := client.GetAPIKeyCredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityProviderResponse
func (client *Client) GetIdentityProviderWithOptions(request *GetIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *GetIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetIdentityProviderRequest
//
// @return GetIdentityProviderResponse
func (client *Client) GetIdentityProvider(request *GetIdentityProviderRequest) (_result *GetIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIdentityProviderResponse{}
	_body, _err := client.GetIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 OAuth2 凭证提供商
//
// @param request - GetOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOAuth2CredentialProviderResponse
func (client *Client) GetOAuth2CredentialProviderWithOptions(request *GetOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *GetOAuth2CredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OAuth2CredentialProviderName) {
		body["OAuth2CredentialProviderName"] = request.OAuth2CredentialProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOAuth2CredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOAuth2CredentialProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 OAuth2 凭证提供商
//
// @param request - GetOAuth2CredentialProviderRequest
//
// @return GetOAuth2CredentialProviderResponse
func (client *Client) GetOAuth2CredentialProvider(request *GetOAuth2CredentialProviderRequest) (_result *GetOAuth2CredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOAuth2CredentialProviderResponse{}
	_body, _err := client.GetOAuth2CredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkloadIdentityResponse
func (client *Client) GetWorkloadIdentityWithOptions(request *GetWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *GetWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkloadIdentity"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkloadIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - GetWorkloadIdentityRequest
//
// @return GetWorkloadIdentityResponse
func (client *Client) GetWorkloadIdentity(request *GetWorkloadIdentityRequest) (_result *GetWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkloadIdentityResponse{}
	_body, _err := client.GetWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出 API 密钥凭证提供商
//
// @param request - ListAPIKeyCredentialProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAPIKeyCredentialProvidersResponse
func (client *Client) ListAPIKeyCredentialProvidersWithOptions(request *ListAPIKeyCredentialProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListAPIKeyCredentialProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAPIKeyCredentialProviders"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAPIKeyCredentialProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出 API 密钥凭证提供商
//
// @param request - ListAPIKeyCredentialProvidersRequest
//
// @return ListAPIKeyCredentialProvidersResponse
func (client *Client) ListAPIKeyCredentialProviders(request *ListAPIKeyCredentialProvidersRequest) (_result *ListAPIKeyCredentialProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAPIKeyCredentialProvidersResponse{}
	_body, _err := client.ListAPIKeyCredentialProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListIdentityProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityProvidersResponse
func (client *Client) ListIdentityProvidersWithOptions(request *ListIdentityProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListIdentityProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentityProviders"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentityProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListIdentityProvidersRequest
//
// @return ListIdentityProvidersResponse
func (client *Client) ListIdentityProviders(request *ListIdentityProvidersRequest) (_result *ListIdentityProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIdentityProvidersResponse{}
	_body, _err := client.ListIdentityProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出 OAuth2 凭证提供商
//
// @param request - ListOAuth2CredentialProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOAuth2CredentialProvidersResponse
func (client *Client) ListOAuth2CredentialProvidersWithOptions(request *ListOAuth2CredentialProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListOAuth2CredentialProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOAuth2CredentialProviders"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOAuth2CredentialProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出 OAuth2 凭证提供商
//
// @param request - ListOAuth2CredentialProvidersRequest
//
// @return ListOAuth2CredentialProvidersResponse
func (client *Client) ListOAuth2CredentialProviders(request *ListOAuth2CredentialProvidersRequest) (_result *ListOAuth2CredentialProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOAuth2CredentialProvidersResponse{}
	_body, _err := client.ListOAuth2CredentialProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListWorkloadIdentitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkloadIdentitiesResponse
func (client *Client) ListWorkloadIdentitiesWithOptions(request *ListWorkloadIdentitiesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkloadIdentitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkloadIdentities"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkloadIdentitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出IdentityProvider
//
// @param request - ListWorkloadIdentitiesRequest
//
// @return ListWorkloadIdentitiesResponse
func (client *Client) ListWorkloadIdentities(request *ListWorkloadIdentitiesRequest) (_result *ListWorkloadIdentitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkloadIdentitiesResponse{}
	_body, _err := client.ListWorkloadIdentitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新一个 API 密钥凭证提供商
//
// @param request - UpdateAPIKeyCredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAPIKeyCredentialProviderResponse
func (client *Client) UpdateAPIKeyCredentialProviderWithOptions(request *UpdateAPIKeyCredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateAPIKeyCredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.APIKey) {
		body["APIKey"] = request.APIKey
	}

	if !dara.IsNil(request.APIKeyCredentialProviderName) {
		body["APIKeyCredentialProviderName"] = request.APIKeyCredentialProviderName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAPIKeyCredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAPIKeyCredentialProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个 API 密钥凭证提供商
//
// @param request - UpdateAPIKeyCredentialProviderRequest
//
// @return UpdateAPIKeyCredentialProviderResponse
func (client *Client) UpdateAPIKeyCredentialProvider(request *UpdateAPIKeyCredentialProviderRequest) (_result *UpdateAPIKeyCredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAPIKeyCredentialProviderResponse{}
	_body, _err := client.UpdateAPIKeyCredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新IdentityProvider
//
// @param tmpReq - UpdateIdentityProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProviderWithOptions(tmpReq *UpdateIdentityProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateIdentityProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateIdentityProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedAudience) {
		request.AllowedAudienceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedAudience, dara.String("AllowedAudience"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedAudienceShrink) {
		body["AllowedAudience"] = request.AllowedAudienceShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DiscoveryURL) {
		body["DiscoveryURL"] = request.DiscoveryURL
	}

	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdentityProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIdentityProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新IdentityProvider
//
// @param request - UpdateIdentityProviderRequest
//
// @return UpdateIdentityProviderResponse
func (client *Client) UpdateIdentityProvider(request *UpdateIdentityProviderRequest) (_result *UpdateIdentityProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIdentityProviderResponse{}
	_body, _err := client.UpdateIdentityProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改 OAuth2 凭证提供商
//
// @param tmpReq - UpdateOAuth2CredentialProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOAuth2CredentialProviderResponse
func (client *Client) UpdateOAuth2CredentialProviderWithOptions(tmpReq *UpdateOAuth2CredentialProviderRequest, runtime *dara.RuntimeOptions) (_result *UpdateOAuth2CredentialProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateOAuth2CredentialProviderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OAuth2ProviderConfig) {
		request.OAuth2ProviderConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OAuth2ProviderConfig, dara.String("OAuth2ProviderConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackURL) {
		body["CallbackURL"] = request.CallbackURL
	}

	if !dara.IsNil(request.CredentialProviderVendor) {
		body["CredentialProviderVendor"] = request.CredentialProviderVendor
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.OAuth2CredentialProviderName) {
		body["OAuth2CredentialProviderName"] = request.OAuth2CredentialProviderName
	}

	if !dara.IsNil(request.OAuth2ProviderConfigShrink) {
		body["OAuth2ProviderConfig"] = request.OAuth2ProviderConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateOAuth2CredentialProvider"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateOAuth2CredentialProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改 OAuth2 凭证提供商
//
// @param request - UpdateOAuth2CredentialProviderRequest
//
// @return UpdateOAuth2CredentialProviderResponse
func (client *Client) UpdateOAuth2CredentialProvider(request *UpdateOAuth2CredentialProviderRequest) (_result *UpdateOAuth2CredentialProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateOAuth2CredentialProviderResponse{}
	_body, _err := client.UpdateOAuth2CredentialProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param tmpReq - UpdateWorkloadIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkloadIdentityResponse
func (client *Client) UpdateWorkloadIdentityWithOptions(tmpReq *UpdateWorkloadIdentityRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkloadIdentityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkloadIdentityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AllowedResourceOAuth2ReturnURLs) {
		request.AllowedResourceOAuth2ReturnURLsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AllowedResourceOAuth2ReturnURLs, dara.String("AllowedResourceOAuth2ReturnURLs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedResourceOAuth2ReturnURLsShrink) {
		body["AllowedResourceOAuth2ReturnURLs"] = request.AllowedResourceOAuth2ReturnURLsShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.IdentityProviderName) {
		body["IdentityProviderName"] = request.IdentityProviderName
	}

	if !dara.IsNil(request.RoleArn) {
		body["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.WorkloadIdentityName) {
		body["WorkloadIdentityName"] = request.WorkloadIdentityName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkloadIdentity"),
		Version:     dara.String("2025-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkloadIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用
//
// @param request - UpdateWorkloadIdentityRequest
//
// @return UpdateWorkloadIdentityResponse
func (client *Client) UpdateWorkloadIdentity(request *UpdateWorkloadIdentityRequest) (_result *UpdateWorkloadIdentityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkloadIdentityResponse{}
	_body, _err := client.UpdateWorkloadIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
