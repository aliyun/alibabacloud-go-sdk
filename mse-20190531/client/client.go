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
	client.Endpoint, _err = client.GetEndpoint(dara.String("mse"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 新增服务鉴权规则
//
// @param request - AddAuthPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAuthPolicyResponse
func (client *Client) AddAuthPolicyWithOptions(request *AddAuthPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddAuthPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AuthRule) {
		query["AuthRule"] = request.AuthRule
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.K8sNamespace) {
		query["K8sNamespace"] = request.K8sNamespace
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAuthPolicy"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAuthPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增服务鉴权规则
//
// @param request - AddAuthPolicyRequest
//
// @return AddAuthPolicyResponse
func (client *Client) AddAuthPolicy(request *AddAuthPolicyRequest) (_result *AddAuthPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAuthPolicyResponse{}
	_body, _err := client.AddAuthPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates authorization information for a gateway.
//
// @param tmpReq - AddAuthResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAuthResourceResponse
func (client *Client) AddAuthResourceWithOptions(tmpReq *AddAuthResourceRequest, runtime *dara.RuntimeOptions) (_result *AddAuthResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddAuthResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthResourceHeaderList) {
		request.AuthResourceHeaderListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthResourceHeaderList, dara.String("AuthResourceHeaderList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AuthId) {
		query["AuthId"] = request.AuthId
	}

	if !dara.IsNil(request.AuthResourceHeaderListShrink) {
		query["AuthResourceHeaderList"] = request.AuthResourceHeaderListShrink
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.IgnoreCase) {
		query["IgnoreCase"] = request.IgnoreCase
	}

	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAuthResource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAuthResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates authorization information for a gateway.
//
// @param request - AddAuthResourceRequest
//
// @return AddAuthResourceResponse
func (client *Client) AddAuthResource(request *AddAuthResourceRequest) (_result *AddAuthResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAuthResourceResponse{}
	_body, _err := client.AddAuthResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a blacklist or a whitelist.
//
// @param request - AddBlackWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBlackWhiteListResponse
func (client *Client) AddBlackWhiteListWithOptions(request *AddBlackWhiteListRequest, runtime *dara.RuntimeOptions) (_result *AddBlackWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.IsWhite) {
		query["IsWhite"] = request.IsWhite
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.ResourceIdJsonList) {
		query["ResourceIdJsonList"] = request.ResourceIdJsonList
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBlackWhiteList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBlackWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a blacklist or a whitelist.
//
// @param request - AddBlackWhiteListRequest
//
// @return AddBlackWhiteListResponse
func (client *Client) AddBlackWhiteList(request *AddBlackWhiteListRequest) (_result *AddBlackWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBlackWhiteListResponse{}
	_body, _err := client.AddBlackWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a gateway.
//
// @param tmpReq - AddGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewayResponse
func (client *Client) AddGatewayWithOptions(tmpReq *AddGatewayRequest, runtime *dara.RuntimeOptions) (_result *AddGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddGatewayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ZoneInfo) {
		request.ZoneInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ZoneInfo, dara.String("ZoneInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClbNetworkType) {
		query["ClbNetworkType"] = request.ClbNetworkType
	}

	if !dara.IsNil(request.EnableHardwareAcceleration) {
		query["EnableHardwareAcceleration"] = request.EnableHardwareAcceleration
	}

	if !dara.IsNil(request.EnableSls) {
		query["EnableSls"] = request.EnableSls
	}

	if !dara.IsNil(request.EnableXtrace) {
		query["EnableXtrace"] = request.EnableXtrace
	}

	if !dara.IsNil(request.EnterpriseSecurityGroup) {
		query["EnterpriseSecurityGroup"] = request.EnterpriseSecurityGroup
	}

	if !dara.IsNil(request.InternetSlbSpec) {
		query["InternetSlbSpec"] = request.InternetSlbSpec
	}

	if !dara.IsNil(request.ManagedEntryNetworkType) {
		query["ManagedEntryNetworkType"] = request.ManagedEntryNetworkType
	}

	if !dara.IsNil(request.MserVersion) {
		query["MserVersion"] = request.MserVersion
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NlbNetworkType) {
		query["NlbNetworkType"] = request.NlbNetworkType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Replica) {
		query["Replica"] = request.Replica
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SlbSpec) {
		query["SlbSpec"] = request.SlbSpec
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VSwitchId2) {
		query["VSwitchId2"] = request.VSwitchId2
	}

	if !dara.IsNil(request.Vpc) {
		query["Vpc"] = request.Vpc
	}

	if !dara.IsNil(request.XtraceRatio) {
		query["XtraceRatio"] = request.XtraceRatio
	}

	if !dara.IsNil(request.ZoneInfoShrink) {
		query["ZoneInfo"] = request.ZoneInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGateway"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a gateway.
//
// @param request - AddGatewayRequest
//
// @return AddGatewayResponse
func (client *Client) AddGateway(request *AddGatewayRequest) (_result *AddGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGatewayResponse{}
	_body, _err := client.AddGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an authentication configuration for a gateway.
//
// @param tmpReq - AddGatewayAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewayAuthResponse
func (client *Client) AddGatewayAuthWithOptions(tmpReq *AddGatewayAuthRequest, runtime *dara.RuntimeOptions) (_result *AddGatewayAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddGatewayAuthShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthResourceList) {
		request.AuthResourceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthResourceList, dara.String("AuthResourceList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExternalAuthZJSON) {
		request.ExternalAuthZJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalAuthZJSON, dara.String("ExternalAuthZJSON"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScopesList) {
		request.ScopesListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScopesList, dara.String("ScopesList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AuthResourceConfig) {
		query["AuthResourceConfig"] = request.AuthResourceConfig
	}

	if !dara.IsNil(request.AuthResourceListShrink) {
		query["AuthResourceList"] = request.AuthResourceListShrink
	}

	if !dara.IsNil(request.AuthResourceMode) {
		query["AuthResourceMode"] = request.AuthResourceMode
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientSecret) {
		query["ClientSecret"] = request.ClientSecret
	}

	if !dara.IsNil(request.CookieDomain) {
		query["CookieDomain"] = request.CookieDomain
	}

	if !dara.IsNil(request.ExternalAuthZJSONShrink) {
		query["ExternalAuthZJSON"] = request.ExternalAuthZJSONShrink
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.IsWhite) {
		query["IsWhite"] = request.IsWhite
	}

	if !dara.IsNil(request.Issuer) {
		query["Issuer"] = request.Issuer
	}

	if !dara.IsNil(request.Jwks) {
		query["Jwks"] = request.Jwks
	}

	if !dara.IsNil(request.LoginUrl) {
		query["LoginUrl"] = request.LoginUrl
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RedirectUrl) {
		query["RedirectUrl"] = request.RedirectUrl
	}

	if !dara.IsNil(request.ScopesListShrink) {
		query["ScopesList"] = request.ScopesListShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Sub) {
		query["Sub"] = request.Sub
	}

	if !dara.IsNil(request.TokenName) {
		query["TokenName"] = request.TokenName
	}

	if !dara.IsNil(request.TokenNamePrefix) {
		query["TokenNamePrefix"] = request.TokenNamePrefix
	}

	if !dara.IsNil(request.TokenPass) {
		query["TokenPass"] = request.TokenPass
	}

	if !dara.IsNil(request.TokenPosition) {
		query["TokenPosition"] = request.TokenPosition
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewayAuth"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewayAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an authentication configuration for a gateway.
//
// @param request - AddGatewayAuthRequest
//
// @return AddGatewayAuthResponse
func (client *Client) AddGatewayAuth(request *AddGatewayAuthRequest) (_result *AddGatewayAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGatewayAuthResponse{}
	_body, _err := client.AddGatewayAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a consumer on which a gateway performs authentication operations.
//
// @param request - AddGatewayAuthConsumerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewayAuthConsumerResponse
func (client *Client) AddGatewayAuthConsumerWithOptions(request *AddGatewayAuthConsumerRequest, runtime *dara.RuntimeOptions) (_result *AddGatewayAuthConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EncodeType) {
		query["EncodeType"] = request.EncodeType
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Jwks) {
		query["Jwks"] = request.Jwks
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.KeyValue) {
		query["KeyValue"] = request.KeyValue
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TokenName) {
		query["TokenName"] = request.TokenName
	}

	if !dara.IsNil(request.TokenPass) {
		query["TokenPass"] = request.TokenPass
	}

	if !dara.IsNil(request.TokenPosition) {
		query["TokenPosition"] = request.TokenPosition
	}

	if !dara.IsNil(request.TokenPrefix) {
		query["TokenPrefix"] = request.TokenPrefix
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewayAuthConsumer"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewayAuthConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a consumer on which a gateway performs authentication operations.
//
// @param request - AddGatewayAuthConsumerRequest
//
// @return AddGatewayAuthConsumerResponse
func (client *Client) AddGatewayAuthConsumer(request *AddGatewayAuthConsumerRequest) (_result *AddGatewayAuthConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGatewayAuthConsumerResponse{}
	_body, _err := client.AddGatewayAuthConsumerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a domain name with a gateway.
//
// @param tmpReq - AddGatewayDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewayDomainResponse
func (client *Client) AddGatewayDomainWithOptions(tmpReq *AddGatewayDomainRequest, runtime *dara.RuntimeOptions) (_result *AddGatewayDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddGatewayDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TlsCipherSuitesConfigJSON) {
		request.TlsCipherSuitesConfigJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TlsCipherSuitesConfigJSON, dara.String("TlsCipherSuitesConfigJSON"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Http2) {
		query["Http2"] = request.Http2
	}

	if !dara.IsNil(request.MustHttps) {
		query["MustHttps"] = request.MustHttps
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.TlsCipherSuitesConfigJSONShrink) {
		query["TlsCipherSuitesConfigJSON"] = request.TlsCipherSuitesConfigJSONShrink
	}

	if !dara.IsNil(request.TlsMax) {
		query["TlsMax"] = request.TlsMax
	}

	if !dara.IsNil(request.TlsMin) {
		query["TlsMin"] = request.TlsMin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewayDomain"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewayDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a domain name with a gateway.
//
// @param request - AddGatewayDomainRequest
//
// @return AddGatewayDomainResponse
func (client *Client) AddGatewayDomain(request *AddGatewayDomainRequest) (_result *AddGatewayDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGatewayDomainResponse{}
	_body, _err := client.AddGatewayDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a gateway route.
//
// @param tmpReq - AddGatewayRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewayRouteResponse
func (client *Client) AddGatewayRouteWithOptions(tmpReq *AddGatewayRouteRequest, runtime *dara.RuntimeOptions) (_result *AddGatewayRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddGatewayRouteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DirectResponseJSON) {
		request.DirectResponseJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DirectResponseJSON, dara.String("DirectResponseJSON"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FallbackServices) {
		request.FallbackServicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FallbackServices, dara.String("FallbackServices"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Predicates) {
		request.PredicatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Predicates, dara.String("Predicates"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RedirectJSON) {
		request.RedirectJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RedirectJSON, dara.String("RedirectJSON"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Services) {
		request.ServicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Services, dara.String("Services"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.DirectResponseJSONShrink) {
		query["DirectResponseJSON"] = request.DirectResponseJSONShrink
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.DomainIdListJSON) {
		query["DomainIdListJSON"] = request.DomainIdListJSON
	}

	if !dara.IsNil(request.EnableWaf) {
		query["EnableWaf"] = request.EnableWaf
	}

	if !dara.IsNil(request.Fallback) {
		query["Fallback"] = request.Fallback
	}

	if !dara.IsNil(request.FallbackServicesShrink) {
		query["FallbackServices"] = request.FallbackServicesShrink
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Policies) {
		query["Policies"] = request.Policies
	}

	if !dara.IsNil(request.PredicatesShrink) {
		query["Predicates"] = request.PredicatesShrink
	}

	if !dara.IsNil(request.RedirectJSONShrink) {
		query["RedirectJSON"] = request.RedirectJSONShrink
	}

	if !dara.IsNil(request.RouteOrder) {
		query["RouteOrder"] = request.RouteOrder
	}

	if !dara.IsNil(request.RouteType) {
		query["RouteType"] = request.RouteType
	}

	if !dara.IsNil(request.ServicesShrink) {
		query["Services"] = request.ServicesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewayRoute"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a gateway route.
//
// @param request - AddGatewayRouteRequest
//
// @return AddGatewayRouteResponse
func (client *Client) AddGatewayRoute(request *AddGatewayRouteRequest) (_result *AddGatewayRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGatewayRouteResponse{}
	_body, _err := client.AddGatewayRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a service version.
//
// @param request - AddGatewayServiceVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewayServiceVersionResponse
func (client *Client) AddGatewayServiceVersionWithOptions(request *AddGatewayServiceVersionRequest, runtime *dara.RuntimeOptions) (_result *AddGatewayServiceVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewayServiceVersion"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewayServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a service version.
//
// @param request - AddGatewayServiceVersionRequest
//
// @return AddGatewayServiceVersionResponse
func (client *Client) AddGatewayServiceVersion(request *AddGatewayServiceVersionRequest) (_result *AddGatewayServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGatewayServiceVersionResponse{}
	_body, _err := client.AddGatewayServiceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a Server Load Balancer (SLB) instance with a gateway.
//
// @param tmpReq - AddGatewaySlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewaySlbResponse
func (client *Client) AddGatewaySlbWithOptions(tmpReq *AddGatewaySlbRequest, runtime *dara.RuntimeOptions) (_result *AddGatewaySlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddGatewaySlbShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VServiceList) {
		request.VServiceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VServiceList, dara.String("VServiceList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.HttpPort) {
		query["HttpPort"] = request.HttpPort
	}

	if !dara.IsNil(request.HttpsPort) {
		query["HttpsPort"] = request.HttpsPort
	}

	if !dara.IsNil(request.HttpsVServerGroupId) {
		query["HttpsVServerGroupId"] = request.HttpsVServerGroupId
	}

	if !dara.IsNil(request.ServiceWeight) {
		query["ServiceWeight"] = request.ServiceWeight
	}

	if !dara.IsNil(request.SlbId) {
		query["SlbId"] = request.SlbId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.VServiceListShrink) {
		query["VServiceList"] = request.VServiceListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewaySlb"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewaySlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Server Load Balancer (SLB) instance with a gateway.
//
// @param request - AddGatewaySlbRequest
//
// @return AddGatewaySlbResponse
func (client *Client) AddGatewaySlb(request *AddGatewaySlbRequest) (_result *AddGatewaySlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGatewaySlbResponse{}
	_body, _err := client.AddGatewaySlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a new migration task
//
// @param request - AddMigrationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMigrationTaskResponse
func (client *Client) AddMigrationTaskWithOptions(request *AddMigrationTaskRequest, runtime *dara.RuntimeOptions) (_result *AddMigrationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.OriginInstanceAddress) {
		query["OriginInstanceAddress"] = request.OriginInstanceAddress
	}

	if !dara.IsNil(request.OriginInstanceName) {
		query["OriginInstanceName"] = request.OriginInstanceName
	}

	if !dara.IsNil(request.OriginInstanceNamespace) {
		query["OriginInstanceNamespace"] = request.OriginInstanceNamespace
	}

	if !dara.IsNil(request.ProjectDesc) {
		query["ProjectDesc"] = request.ProjectDesc
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.SyncType) {
		query["SyncType"] = request.SyncType
	}

	if !dara.IsNil(request.TargetClusterName) {
		query["TargetClusterName"] = request.TargetClusterName
	}

	if !dara.IsNil(request.TargetClusterUrl) {
		query["TargetClusterUrl"] = request.TargetClusterUrl
	}

	if !dara.IsNil(request.TargetInstanceId) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMigrationTask"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMigrationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a new migration task
//
// @param request - AddMigrationTaskRequest
//
// @return AddMigrationTaskResponse
func (client *Client) AddMigrationTask(request *AddMigrationTaskRequest) (_result *AddMigrationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMigrationTaskResponse{}
	_body, _err := client.AddMigrationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a mock rule.
//
// @param request - AddMockRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMockRuleResponse
func (client *Client) AddMockRuleWithOptions(request *AddMockRuleRequest, runtime *dara.RuntimeOptions) (_result *AddMockRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConsumerAppIds) {
		query["ConsumerAppIds"] = request.ConsumerAppIds
	}

	if !dara.IsNil(request.DubboMockItems) {
		query["DubboMockItems"] = request.DubboMockItems
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.ExtraJson) {
		query["ExtraJson"] = request.ExtraJson
	}

	if !dara.IsNil(request.MockType) {
		query["MockType"] = request.MockType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProviderAppId) {
		query["ProviderAppId"] = request.ProviderAppId
	}

	if !dara.IsNil(request.ProviderAppName) {
		query["ProviderAppName"] = request.ProviderAppName
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ScMockItems) {
		query["ScMockItems"] = request.ScMockItems
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMockRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMockRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mock rule.
//
// @param request - AddMockRuleRequest
//
// @return AddMockRuleResponse
func (client *Client) AddMockRule(request *AddMockRuleRequest) (_result *AddMockRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMockRuleResponse{}
	_body, _err := client.AddMockRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a certificate with a domain name of a gateway. You can specify a certificate that is hosted in Alibaba Cloud Security.
//
// @param request - AddSSLCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSSLCertResponse
func (client *Client) AddSSLCertWithOptions(request *AddSSLCertRequest, runtime *dara.RuntimeOptions) (_result *AddSSLCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSSLCert"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSSLCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a certificate with a domain name of a gateway. You can specify a certificate that is hosted in Alibaba Cloud Security.
//
// @param request - AddSSLCertRequest
//
// @return AddSSLCertResponse
func (client *Client) AddSSLCert(request *AddSSLCertRequest) (_result *AddSSLCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSSLCertResponse{}
	_body, _err := client.AddSSLCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a security group rule for a gateway.
//
// @param request - AddSecurityGroupRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSecurityGroupRuleResponse
func (client *Client) AddSecurityGroupRuleWithOptions(request *AddSecurityGroupRuleRequest, runtime *dara.RuntimeOptions) (_result *AddSecurityGroupRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSecurityGroupRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSecurityGroupRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a security group rule for a gateway.
//
// @param request - AddSecurityGroupRuleRequest
//
// @return AddSecurityGroupRuleResponse
func (client *Client) AddSecurityGroupRule(request *AddSecurityGroupRuleRequest) (_result *AddSecurityGroupRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddSecurityGroupRuleResponse{}
	_body, _err := client.AddSecurityGroupRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a Nacos service source.
//
// @param tmpReq - AddServiceSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddServiceSourceResponse
func (client *Client) AddServiceSourceWithOptions(tmpReq *AddServiceSourceRequest, runtime *dara.RuntimeOptions) (_result *AddServiceSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddServiceSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GroupList) {
		request.GroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupList, dara.String("GroupList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IngressOptionsRequest) {
		request.IngressOptionsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IngressOptionsRequest, dara.String("IngressOptionsRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PathList) {
		request.PathListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PathList, dara.String("PathList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ToAuthorizeSecurityGroups) {
		request.ToAuthorizeSecurityGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToAuthorizeSecurityGroups, dara.String("ToAuthorizeSecurityGroups"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.GroupListShrink) {
		query["GroupList"] = request.GroupListShrink
	}

	if !dara.IsNil(request.IngressOptionsRequestShrink) {
		query["IngressOptionsRequest"] = request.IngressOptionsRequestShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PathListShrink) {
		query["PathList"] = request.PathListShrink
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.ToAuthorizeSecurityGroupsShrink) {
		query["ToAuthorizeSecurityGroups"] = request.ToAuthorizeSecurityGroupsShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddServiceSource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddServiceSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a Nacos service source.
//
// @param request - AddServiceSourceRequest
//
// @return AddServiceSourceResponse
func (client *Client) AddServiceSource(request *AddServiceSourceRequest) (_result *AddServiceSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddServiceSourceResponse{}
	_body, _err := client.AddServiceSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a route for a gateway.
//
// @param request - ApplyGatewayRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyGatewayRouteResponse
func (client *Client) ApplyGatewayRouteWithOptions(request *ApplyGatewayRouteRequest, runtime *dara.RuntimeOptions) (_result *ApplyGatewayRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyGatewayRoute"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a route for a gateway.
//
// @param request - ApplyGatewayRouteRequest
//
// @return ApplyGatewayRouteResponse
func (client *Client) ApplyGatewayRoute(request *ApplyGatewayRouteRequest) (_result *ApplyGatewayRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyGatewayRouteResponse{}
	_body, _err := client.ApplyGatewayRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a tag-based routing rule.
//
// @param tmpReq - ApplyTagPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyTagPoliciesResponse
func (client *Client) ApplyTagPoliciesWithOptions(tmpReq *ApplyTagPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ApplyTagPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ApplyTagPoliciesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RulesShrink) {
		query["Rules"] = request.RulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyTagPolicies"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyTagPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a tag-based routing rule.
//
// @param request - ApplyTagPoliciesRequest
//
// @return ApplyTagPoliciesResponse
func (client *Client) ApplyTagPolicies(request *ApplyTagPoliciesRequest) (_result *ApplyTagPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyTagPoliciesResponse{}
	_body, _err := client.ApplyTagPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds traffic protection behavior.
//
// @param request - BindSentinelBlockFallbackDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindSentinelBlockFallbackDefinitionResponse
func (client *Client) BindSentinelBlockFallbackDefinitionWithOptions(request *BindSentinelBlockFallbackDefinitionRequest, runtime *dara.RuntimeOptions) (_result *BindSentinelBlockFallbackDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.FallbackId) {
		query["FallbackId"] = request.FallbackId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindSentinelBlockFallbackDefinition"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindSentinelBlockFallbackDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds traffic protection behavior.
//
// @param request - BindSentinelBlockFallbackDefinitionRequest
//
// @return BindSentinelBlockFallbackDefinitionResponse
func (client *Client) BindSentinelBlockFallbackDefinition(request *BindSentinelBlockFallbackDefinitionRequest) (_result *BindSentinelBlockFallbackDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindSentinelBlockFallbackDefinitionResponse{}
	_body, _err := client.BindSentinelBlockFallbackDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Resource Transfer
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Resource Transfer
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Copy Nacos Configuration
//
// Description:
//
// > This OpenAPI is not the Nacos-SDK API. For information related to the Nacos-SDK API, please refer to the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - CloneNacosConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneNacosConfigResponse
func (client *Client) CloneNacosConfigWithOptions(request *CloneNacosConfigRequest, runtime *dara.RuntimeOptions) (_result *CloneNacosConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DataIds) {
		query["DataIds"] = request.DataIds
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OriginNamespaceId) {
		query["OriginNamespaceId"] = request.OriginNamespaceId
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.TargetNamespaceId) {
		query["TargetNamespaceId"] = request.TargetNamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneNacosConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Copy Nacos Configuration
//
// Description:
//
// > This OpenAPI is not the Nacos-SDK API. For information related to the Nacos-SDK API, please refer to the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - CloneNacosConfigRequest
//
// @return CloneNacosConfigResponse
func (client *Client) CloneNacosConfig(request *CloneNacosConfigRequest) (_result *CloneNacosConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloneNacosConfigResponse{}
	_body, _err := client.CloneNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones rules from Application High Availability Service.
//
// @param request - CloneSentinelRuleFromAhasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneSentinelRuleFromAhasResponse
func (client *Client) CloneSentinelRuleFromAhasWithOptions(request *CloneSentinelRuleFromAhasRequest, runtime *dara.RuntimeOptions) (_result *CloneSentinelRuleFromAhasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AhasNamespace) {
		query["AhasNamespace"] = request.AhasNamespace
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.IsAHASPublicRegion) {
		query["IsAHASPublicRegion"] = request.IsAHASPublicRegion
	}

	if !dara.IsNil(request.MseAppName) {
		query["MseAppName"] = request.MseAppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneSentinelRuleFromAhas"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneSentinelRuleFromAhasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones rules from Application High Availability Service.
//
// @param request - CloneSentinelRuleFromAhasRequest
//
// @return CloneSentinelRuleFromAhasResponse
func (client *Client) CloneSentinelRuleFromAhas(request *CloneSentinelRuleFromAhasRequest) (_result *CloneSentinelRuleFromAhasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloneSentinelRuleFromAhasResponse{}
	_body, _err := client.CloneSentinelRuleFromAhasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application.
//
// @param tmpReq - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(tmpReq *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SentinelEnable) {
		query["SentinelEnable"] = request.SentinelEnable
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SwitchEnable) {
		query["SwitchEnable"] = request.SwitchEnable
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplication"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application.
//
// @param request - CreateApplicationRequest
//
// @return CreateApplicationResponse
func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a circuit breaking rule.
//
// @param request - CreateCircuitBreakerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCircuitBreakerRuleResponse
func (client *Client) CreateCircuitBreakerRuleWithOptions(request *CreateCircuitBreakerRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateCircuitBreakerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.HalfOpenBaseAmountPerStep) {
		query["HalfOpenBaseAmountPerStep"] = request.HalfOpenBaseAmountPerStep
	}

	if !dara.IsNil(request.HalfOpenRecoveryStepNum) {
		query["HalfOpenRecoveryStepNum"] = request.HalfOpenRecoveryStepNum
	}

	if !dara.IsNil(request.MaxAllowedRtMs) {
		query["MaxAllowedRtMs"] = request.MaxAllowedRtMs
	}

	if !dara.IsNil(request.MinRequestAmount) {
		query["MinRequestAmount"] = request.MinRequestAmount
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RetryTimeoutMs) {
		query["RetryTimeoutMs"] = request.RetryTimeoutMs
	}

	if !dara.IsNil(request.StatIntervalMs) {
		query["StatIntervalMs"] = request.StatIntervalMs
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCircuitBreakerRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCircuitBreakerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a circuit breaking rule.
//
// @param request - CreateCircuitBreakerRuleRequest
//
// @return CreateCircuitBreakerRuleResponse
func (client *Client) CreateCircuitBreakerRule(request *CreateCircuitBreakerRuleRequest) (_result *CreateCircuitBreakerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCircuitBreakerRuleResponse{}
	_body, _err := client.CreateCircuitBreakerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create an MSE registration and configuration center instance
//
// Description:
//
// Please ensure that you fully understand the billing method and pricing of the MSE (Microservice Engine) product before using this interface.
//
// @param request - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClusterSpecification) {
		query["ClusterSpecification"] = request.ClusterSpecification
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.ClusterVersion) {
		query["ClusterVersion"] = request.ClusterVersion
	}

	if !dara.IsNil(request.ConnectionType) {
		query["ConnectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.EipEnabled) {
		query["EipEnabled"] = request.EipEnabled
	}

	if !dara.IsNil(request.InstanceCount) {
		query["InstanceCount"] = request.InstanceCount
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.MseVersion) {
		query["MseVersion"] = request.MseVersion
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
	}

	if !dara.IsNil(request.PrivateSlbSpecification) {
		query["PrivateSlbSpecification"] = request.PrivateSlbSpecification
	}

	if !dara.IsNil(request.PubNetworkFlow) {
		query["PubNetworkFlow"] = request.PubNetworkFlow
	}

	if !dara.IsNil(request.PubSlbSpecification) {
		query["PubSlbSpecification"] = request.PubSlbSpecification
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupType) {
		query["SecurityGroupType"] = request.SecurityGroupType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create an MSE registration and configuration center instance
//
// Description:
//
// Please ensure that you fully understand the billing method and pricing of the MSE (Microservice Engine) product before using this interface.
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a namespace in a Nacos instance.
//
// @param request - CreateEngineNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEngineNamespaceResponse
func (client *Client) CreateEngineNamespaceWithOptions(request *CreateEngineNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateEngineNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ServiceCount) {
		query["ServiceCount"] = request.ServiceCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEngineNamespace"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEngineNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a namespace in a Nacos instance.
//
// @param request - CreateEngineNamespaceRequest
//
// @return CreateEngineNamespaceResponse
func (client *Client) CreateEngineNamespace(request *CreateEngineNamespaceRequest) (_result *CreateEngineNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEngineNamespaceResponse{}
	_body, _err := client.CreateEngineNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a throttling rule.
//
// @param request - CreateFlowRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowRuleResponse
func (client *Client) CreateFlowRuleWithOptions(request *CreateFlowRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ControlBehavior) {
		query["ControlBehavior"] = request.ControlBehavior
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.LimitApp) {
		query["LimitApp"] = request.LimitApp
	}

	if !dara.IsNil(request.MaxQueueingTimeMs) {
		query["MaxQueueingTimeMs"] = request.MaxQueueingTimeMs
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a throttling rule.
//
// @param request - CreateFlowRuleRequest
//
// @return CreateFlowRuleResponse
func (client *Client) CreateFlowRule(request *CreateFlowRuleRequest) (_result *CreateFlowRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlowRuleResponse{}
	_body, _err := client.CreateFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关路由熔断规则
//
// @param request - CreateGatewayCircuitBreakerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayCircuitBreakerRuleResponse
func (client *Client) CreateGatewayCircuitBreakerRuleWithOptions(request *CreateGatewayCircuitBreakerRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateGatewayCircuitBreakerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.BehaviorType) {
		query["BehaviorType"] = request.BehaviorType
	}

	if !dara.IsNil(request.BodyEncoding) {
		query["BodyEncoding"] = request.BodyEncoding
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.MaxAllowedMs) {
		query["MaxAllowedMs"] = request.MaxAllowedMs
	}

	if !dara.IsNil(request.MinRequestAmount) {
		query["MinRequestAmount"] = request.MinRequestAmount
	}

	if !dara.IsNil(request.RecoveryTimeoutSec) {
		query["RecoveryTimeoutSec"] = request.RecoveryTimeoutSec
	}

	if !dara.IsNil(request.ResponseContentBody) {
		query["ResponseContentBody"] = request.ResponseContentBody
	}

	if !dara.IsNil(request.ResponseRedirectUrl) {
		query["ResponseRedirectUrl"] = request.ResponseRedirectUrl
	}

	if !dara.IsNil(request.ResponseStatusCode) {
		query["ResponseStatusCode"] = request.ResponseStatusCode
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.StatDurationSec) {
		query["StatDurationSec"] = request.StatDurationSec
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	if !dara.IsNil(request.TriggerRatio) {
		query["TriggerRatio"] = request.TriggerRatio
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGatewayCircuitBreakerRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayCircuitBreakerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关路由熔断规则
//
// @param request - CreateGatewayCircuitBreakerRuleRequest
//
// @return CreateGatewayCircuitBreakerRuleResponse
func (client *Client) CreateGatewayCircuitBreakerRule(request *CreateGatewayCircuitBreakerRuleRequest) (_result *CreateGatewayCircuitBreakerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGatewayCircuitBreakerRuleResponse{}
	_body, _err := client.CreateGatewayCircuitBreakerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a throttling rule for a gateway.
//
// @param request - CreateGatewayFlowRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayFlowRuleResponse
func (client *Client) CreateGatewayFlowRuleWithOptions(request *CreateGatewayFlowRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateGatewayFlowRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.BehaviorType) {
		query["BehaviorType"] = request.BehaviorType
	}

	if !dara.IsNil(request.BodyEncoding) {
		query["BodyEncoding"] = request.BodyEncoding
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ResponseContentBody) {
		query["ResponseContentBody"] = request.ResponseContentBody
	}

	if !dara.IsNil(request.ResponseRedirectUrl) {
		query["ResponseRedirectUrl"] = request.ResponseRedirectUrl
	}

	if !dara.IsNil(request.ResponseStatusCode) {
		query["ResponseStatusCode"] = request.ResponseStatusCode
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGatewayFlowRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a throttling rule for a gateway.
//
// @param request - CreateGatewayFlowRuleRequest
//
// @return CreateGatewayFlowRuleResponse
func (client *Client) CreateGatewayFlowRule(request *CreateGatewayFlowRuleRequest) (_result *CreateGatewayFlowRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGatewayFlowRuleResponse{}
	_body, _err := client.CreateGatewayFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网关路由隔离规则
//
// @param request - CreateGatewayIsolationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayIsolationRuleResponse
func (client *Client) CreateGatewayIsolationRuleWithOptions(request *CreateGatewayIsolationRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateGatewayIsolationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.BehaviorType) {
		query["BehaviorType"] = request.BehaviorType
	}

	if !dara.IsNil(request.BodyEncoding) {
		query["BodyEncoding"] = request.BodyEncoding
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.MaxConcurrency) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.ResponseContentBody) {
		query["ResponseContentBody"] = request.ResponseContentBody
	}

	if !dara.IsNil(request.ResponseRedirectUrl) {
		query["ResponseRedirectUrl"] = request.ResponseRedirectUrl
	}

	if !dara.IsNil(request.ResponseStatusCode) {
		query["ResponseStatusCode"] = request.ResponseStatusCode
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGatewayIsolationRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网关路由隔离规则
//
// @param request - CreateGatewayIsolationRuleRequest
//
// @return CreateGatewayIsolationRuleResponse
func (client *Client) CreateGatewayIsolationRule(request *CreateGatewayIsolationRuleRequest) (_result *CreateGatewayIsolationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGatewayIsolationRuleResponse{}
	_body, _err := client.CreateGatewayIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建隔离规则
//
// @param request - CreateIsolationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIsolationRuleResponse
func (client *Client) CreateIsolationRuleWithOptions(request *CreateIsolationRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateIsolationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.LimitApp) {
		query["LimitApp"] = request.LimitApp
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIsolationRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建隔离规则
//
// @param request - CreateIsolationRuleRequest
//
// @return CreateIsolationRuleResponse
func (client *Client) CreateIsolationRule(request *CreateIsolationRuleRequest) (_result *CreateIsolationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIsolationRuleResponse{}
	_body, _err := client.CreateIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateMseServiceApplication is deprecated, please use mse::2019-05-31::CreateApplication instead.
//
// Summary:
//
// Creates an application.
//
// @param request - CreateMseServiceApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMseServiceApplicationResponse
func (client *Client) CreateMseServiceApplicationWithOptions(request *CreateMseServiceApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateMseServiceApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.MseVersion) {
		query["MseVersion"] = request.MseVersion
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SentinelEnable) {
		query["SentinelEnable"] = request.SentinelEnable
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SwitchEnable) {
		query["SwitchEnable"] = request.SwitchEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMseServiceApplication"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMseServiceApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateMseServiceApplication is deprecated, please use mse::2019-05-31::CreateApplication instead.
//
// Summary:
//
// Creates an application.
//
// @param request - CreateMseServiceApplicationRequest
//
// @return CreateMseServiceApplicationResponse
// Deprecated
func (client *Client) CreateMseServiceApplication(request *CreateMseServiceApplicationRequest) (_result *CreateMseServiceApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMseServiceApplicationResponse{}
	_body, _err := client.CreateMseServiceApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Nacos configuration.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - CreateNacosConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNacosConfigResponse
func (client *Client) CreateNacosConfigWithOptions(request *CreateNacosConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateNacosConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BetaIps) {
		query["BetaIps"] = request.BetaIps
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNacosConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Nacos configuration.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - CreateNacosConfigRequest
//
// @return CreateNacosConfigResponse
func (client *Client) CreateNacosConfig(request *CreateNacosConfigRequest) (_result *CreateNacosConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNacosConfigResponse{}
	_body, _err := client.CreateNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - CreateNacosInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNacosInstanceResponse
func (client *Client) CreateNacosInstanceWithOptions(request *CreateNacosInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateNacosInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Ephemeral) {
		query["Ephemeral"] = request.Ephemeral
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Metadata) {
		body["Metadata"] = request.Metadata
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNacosInstance"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNacosInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - CreateNacosInstanceRequest
//
// @return CreateNacosInstanceResponse
func (client *Client) CreateNacosInstance(request *CreateNacosInstanceRequest) (_result *CreateNacosInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNacosInstanceResponse{}
	_body, _err := client.CreateNacosInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个MCP Server
//
// @param request - CreateNacosMcpServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNacosMcpServerResponse
func (client *Client) CreateNacosMcpServerWithOptions(request *CreateNacosMcpServerRequest, runtime *dara.RuntimeOptions) (_result *CreateNacosMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.EncryptToolSpec) {
		query["EncryptToolSpec"] = request.EncryptToolSpec
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.ServerName) {
		query["ServerName"] = request.ServerName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndpointSpecification) {
		body["EndpointSpecification"] = request.EndpointSpecification
	}

	if !dara.IsNil(request.ServerSpecification) {
		body["ServerSpecification"] = request.ServerSpecification
	}

	if !dara.IsNil(request.ToolSpecification) {
		body["ToolSpecification"] = request.ToolSpecification
	}

	if !dara.IsNil(request.YamlConfig) {
		body["YamlConfig"] = request.YamlConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNacosMcpServer"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNacosMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个MCP Server
//
// @param request - CreateNacosMcpServerRequest
//
// @return CreateNacosMcpServerResponse
func (client *Client) CreateNacosMcpServer(request *CreateNacosMcpServerRequest) (_result *CreateNacosMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNacosMcpServerResponse{}
	_body, _err := client.CreateNacosMcpServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Nacos service.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - CreateNacosServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNacosServiceResponse
func (client *Client) CreateNacosServiceWithOptions(request *CreateNacosServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateNacosServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Ephemeral) {
		query["Ephemeral"] = request.Ephemeral
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.ProtectThreshold) {
		query["ProtectThreshold"] = request.ProtectThreshold
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNacosService"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNacosServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Nacos service.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - CreateNacosServiceRequest
//
// @return CreateNacosServiceResponse
func (client *Client) CreateNacosService(request *CreateNacosServiceRequest) (_result *CreateNacosServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNacosServiceResponse{}
	_body, _err := client.CreateNacosServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CreateNamespace
//
// @param tmpReq - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithOptions(tmpReq *CreateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNamespaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Describe) {
		query["Describe"] = request.Describe
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNamespace"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateNamespace
//
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates a lane for end-to-end canary release.
//
// @param tmpReq - CreateOrUpdateSwimmingLaneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateSwimmingLaneResponse
func (client *Client) CreateOrUpdateSwimmingLaneWithOptions(tmpReq *CreateOrUpdateSwimmingLaneRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateSwimmingLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOrUpdateSwimmingLaneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GatewaySwimmingLaneRouteJson) {
		request.GatewaySwimmingLaneRouteJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GatewaySwimmingLaneRouteJson, dara.String("GatewaySwimmingLaneRouteJson"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.EnableRules) {
		query["EnableRules"] = request.EnableRules
	}

	if !dara.IsNil(request.EntryRule) {
		query["EntryRule"] = request.EntryRule
	}

	if !dara.IsNil(request.GatewaySwimmingLaneRouteJsonShrink) {
		query["GatewaySwimmingLaneRouteJson"] = request.GatewaySwimmingLaneRouteJsonShrink
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PathIndependentPercentageEnable) {
		query["PathIndependentPercentageEnable"] = request.PathIndependentPercentageEnable
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntryRules) {
		body["EntryRules"] = request.EntryRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateSwimmingLane"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateSwimmingLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a lane for end-to-end canary release.
//
// @param request - CreateOrUpdateSwimmingLaneRequest
//
// @return CreateOrUpdateSwimmingLaneResponse
func (client *Client) CreateOrUpdateSwimmingLane(request *CreateOrUpdateSwimmingLaneRequest) (_result *CreateOrUpdateSwimmingLaneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOrUpdateSwimmingLaneResponse{}
	_body, _err := client.CreateOrUpdateSwimmingLaneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates a lane group for end-to-end canary release.
//
// @param tmpReq - CreateOrUpdateSwimmingLaneGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateSwimmingLaneGroupResponse
func (client *Client) CreateOrUpdateSwimmingLaneGroupWithOptions(tmpReq *CreateOrUpdateSwimmingLaneGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateOrUpdateSwimmingLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOrUpdateSwimmingLaneGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Paths) {
		request.PathsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Paths, dara.String("Paths"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RouteIds) {
		request.RouteIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RouteIds, dara.String("RouteIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.CanaryModel) {
		query["CanaryModel"] = request.CanaryModel
	}

	if !dara.IsNil(request.DbGrayEnable) {
		query["DbGrayEnable"] = request.DbGrayEnable
	}

	if !dara.IsNil(request.EntryApp) {
		query["EntryApp"] = request.EntryApp
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MessageQueueFilterSide) {
		query["MessageQueueFilterSide"] = request.MessageQueueFilterSide
	}

	if !dara.IsNil(request.MessageQueueGrayEnable) {
		query["MessageQueueGrayEnable"] = request.MessageQueueGrayEnable
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PathsShrink) {
		query["Paths"] = request.PathsShrink
	}

	if !dara.IsNil(request.RecordCanaryDetail) {
		query["RecordCanaryDetail"] = request.RecordCanaryDetail
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RouteIdsShrink) {
		query["RouteIds"] = request.RouteIdsShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SwimVersion) {
		query["SwimVersion"] = request.SwimVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrUpdateSwimmingLaneGroup"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrUpdateSwimmingLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or updates a lane group for end-to-end canary release.
//
// @param request - CreateOrUpdateSwimmingLaneGroupRequest
//
// @return CreateOrUpdateSwimmingLaneGroupResponse
func (client *Client) CreateOrUpdateSwimmingLaneGroup(request *CreateOrUpdateSwimmingLaneGroupRequest) (_result *CreateOrUpdateSwimmingLaneGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOrUpdateSwimmingLaneGroupResponse{}
	_body, _err := client.CreateOrUpdateSwimmingLaneGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a plug-in configuration.
//
// @param tmpReq - CreatePluginConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePluginConfigResponse
func (client *Client) CreatePluginConfigWithOptions(tmpReq *CreatePluginConfigRequest, runtime *dara.RuntimeOptions) (_result *CreatePluginConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePluginConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceIdList) {
		request.ResourceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIdList, dara.String("ResourceIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConfigLevel) {
		query["ConfigLevel"] = request.ConfigLevel
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.ResourceIdListShrink) {
		query["ResourceIdList"] = request.ResourceIdListShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePluginConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePluginConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a plug-in configuration.
//
// @param request - CreatePluginConfigRequest
//
// @return CreatePluginConfigResponse
func (client *Client) CreatePluginConfig(request *CreatePluginConfigRequest) (_result *CreatePluginConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePluginConfigResponse{}
	_body, _err := client.CreatePluginConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建行为管理
//
// @param request - CreateSentinelBlockFallbackDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSentinelBlockFallbackDefinitionResponse
func (client *Client) CreateSentinelBlockFallbackDefinitionWithOptions(request *CreateSentinelBlockFallbackDefinitionRequest, runtime *dara.RuntimeOptions) (_result *CreateSentinelBlockFallbackDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.FallbackBehavior) {
		query["FallbackBehavior"] = request.FallbackBehavior
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceClassification) {
		query["ResourceClassification"] = request.ResourceClassification
	}

	if !dara.IsNil(request.Scenario) {
		query["Scenario"] = request.Scenario
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSentinelBlockFallbackDefinition"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSentinelBlockFallbackDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建行为管理
//
// @param request - CreateSentinelBlockFallbackDefinitionRequest
//
// @return CreateSentinelBlockFallbackDefinitionResponse
func (client *Client) CreateSentinelBlockFallbackDefinition(request *CreateSentinelBlockFallbackDefinitionRequest) (_result *CreateSentinelBlockFallbackDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSentinelBlockFallbackDefinitionResponse{}
	_body, _err := client.CreateSentinelBlockFallbackDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建热点参数防护规则（HTTP 请求）
//
// @param request - CreateWebFlowRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWebFlowRuleResponse
func (client *Client) CreateWebFlowRuleWithOptions(request *CreateWebFlowRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateWebFlowRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Burst) {
		query["Burst"] = request.Burst
	}

	if !dara.IsNil(request.ControlBehavior) {
		query["ControlBehavior"] = request.ControlBehavior
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.MaxQueueingTimeMs) {
		query["MaxQueueingTimeMs"] = request.MaxQueueingTimeMs
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ParamItem) {
		query["ParamItem"] = request.ParamItem
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceMode) {
		query["ResourceMode"] = request.ResourceMode
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StatIntervalMs) {
		query["StatIntervalMs"] = request.StatIntervalMs
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWebFlowRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWebFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建热点参数防护规则（HTTP 请求）
//
// @param request - CreateWebFlowRuleRequest
//
// @return CreateWebFlowRuleResponse
func (client *Client) CreateWebFlowRule(request *CreateWebFlowRuleRequest) (_result *CreateWebFlowRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWebFlowRuleResponse{}
	_body, _err := client.CreateWebFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a znode.
//
// @param request - CreateZnodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateZnodeResponse
func (client *Client) CreateZnodeWithOptions(request *CreateZnodeRequest, runtime *dara.RuntimeOptions) (_result *CreateZnodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateZnode"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateZnodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a znode.
//
// @param request - CreateZnodeRequest
//
// @return CreateZnodeResponse
func (client *Client) CreateZnode(request *CreateZnodeRequest) (_result *CreateZnodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateZnodeResponse{}
	_body, _err := client.CreateZnodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an authorized resource.
//
// @param request - DeleteAuthResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAuthResourceResponse
func (client *Client) DeleteAuthResourceWithOptions(request *DeleteAuthResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteAuthResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAuthResource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAuthResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an authorized resource.
//
// @param request - DeleteAuthResourceRequest
//
// @return DeleteAuthResourceResponse
func (client *Client) DeleteAuthResource(request *DeleteAuthResourceRequest) (_result *DeleteAuthResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAuthResourceResponse{}
	_body, _err := client.DeleteAuthResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除黑白名单
//
// @param request - DeleteBlackWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBlackWhiteListResponse
func (client *Client) DeleteBlackWhiteListWithOptions(request *DeleteBlackWhiteListRequest, runtime *dara.RuntimeOptions) (_result *DeleteBlackWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBlackWhiteList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBlackWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除黑白名单
//
// @param request - DeleteBlackWhiteListRequest
//
// @return DeleteBlackWhiteListResponse
func (client *Client) DeleteBlackWhiteList(request *DeleteBlackWhiteListRequest) (_result *DeleteBlackWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBlackWhiteListResponse{}
	_body, _err := client.DeleteBlackWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes circuit breaking rules.
//
// @param tmpReq - DeleteCircuitBreakerRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCircuitBreakerRulesResponse
func (client *Client) DeleteCircuitBreakerRulesWithOptions(tmpReq *DeleteCircuitBreakerRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCircuitBreakerRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteCircuitBreakerRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCircuitBreakerRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCircuitBreakerRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes circuit breaking rules.
//
// @param request - DeleteCircuitBreakerRulesRequest
//
// @return DeleteCircuitBreakerRulesResponse
func (client *Client) DeleteCircuitBreakerRules(request *DeleteCircuitBreakerRulesRequest) (_result *DeleteCircuitBreakerRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCircuitBreakerRulesResponse{}
	_body, _err := client.DeleteCircuitBreakerRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Microservices Engine (MSE) instance.
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Microservices Engine (MSE) instance.
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a namespace from a Nacos instance.
//
// @param request - DeleteEngineNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEngineNamespaceResponse
func (client *Client) DeleteEngineNamespaceWithOptions(request *DeleteEngineNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteEngineNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEngineNamespace"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEngineNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace from a Nacos instance.
//
// @param request - DeleteEngineNamespaceRequest
//
// @return DeleteEngineNamespaceResponse
func (client *Client) DeleteEngineNamespace(request *DeleteEngineNamespaceRequest) (_result *DeleteEngineNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEngineNamespaceResponse{}
	_body, _err := client.DeleteEngineNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes throttling rules.
//
// @param tmpReq - DeleteFlowRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowRulesResponse
func (client *Client) DeleteFlowRulesWithOptions(tmpReq *DeleteFlowRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteFlowRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlowRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes throttling rules.
//
// @param request - DeleteFlowRulesRequest
//
// @return DeleteFlowRulesResponse
func (client *Client) DeleteFlowRules(request *DeleteFlowRulesRequest) (_result *DeleteFlowRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlowRulesResponse{}
	_body, _err := client.DeleteFlowRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a gateway.
//
// @param request - DeleteGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithOptions(request *DeleteGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DeleteSlb) {
		query["DeleteSlb"] = request.DeleteSlb
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGateway"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a gateway.
//
// @param request - DeleteGatewayRequest
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGateway(request *DeleteGatewayRequest) (_result *DeleteGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关认证
//
// @param request - DeleteGatewayAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayAuthResponse
func (client *Client) DeleteGatewayAuthWithOptions(request *DeleteGatewayAuthRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayAuth"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关认证
//
// @param request - DeleteGatewayAuthRequest
//
// @return DeleteGatewayAuthResponse
func (client *Client) DeleteGatewayAuth(request *DeleteGatewayAuthRequest) (_result *DeleteGatewayAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayAuthResponse{}
	_body, _err := client.DeleteGatewayAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a consumer on which a gateway performs authentication operations.
//
// @param request - DeleteGatewayAuthConsumerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayAuthConsumerResponse
func (client *Client) DeleteGatewayAuthConsumerWithOptions(request *DeleteGatewayAuthConsumerRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayAuthConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayAuthConsumer"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayAuthConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a consumer on which a gateway performs authentication operations.
//
// @param request - DeleteGatewayAuthConsumerRequest
//
// @return DeleteGatewayAuthConsumerResponse
func (client *Client) DeleteGatewayAuthConsumer(request *DeleteGatewayAuthConsumerRequest) (_result *DeleteGatewayAuthConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayAuthConsumerResponse{}
	_body, _err := client.DeleteGatewayAuthConsumerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes resource permissions from the consumer on which a gateway performs authentication operations.
//
// @param request - DeleteGatewayAuthConsumerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayAuthConsumerResourceResponse
func (client *Client) DeleteGatewayAuthConsumerResourceWithOptions(request *DeleteGatewayAuthConsumerResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayAuthConsumerResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.IdList) {
		query["IdList"] = request.IdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayAuthConsumerResource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayAuthConsumerResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes resource permissions from the consumer on which a gateway performs authentication operations.
//
// @param request - DeleteGatewayAuthConsumerResourceRequest
//
// @return DeleteGatewayAuthConsumerResourceResponse
func (client *Client) DeleteGatewayAuthConsumerResource(request *DeleteGatewayAuthConsumerResourceRequest) (_result *DeleteGatewayAuthConsumerResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayAuthConsumerResourceResponse{}
	_body, _err := client.DeleteGatewayAuthConsumerResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关路由熔断规则
//
// @param request - DeleteGatewayCircuitBreakerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayCircuitBreakerRuleResponse
func (client *Client) DeleteGatewayCircuitBreakerRuleWithOptions(request *DeleteGatewayCircuitBreakerRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayCircuitBreakerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayCircuitBreakerRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayCircuitBreakerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关路由熔断规则
//
// @param request - DeleteGatewayCircuitBreakerRuleRequest
//
// @return DeleteGatewayCircuitBreakerRuleResponse
func (client *Client) DeleteGatewayCircuitBreakerRule(request *DeleteGatewayCircuitBreakerRuleRequest) (_result *DeleteGatewayCircuitBreakerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayCircuitBreakerRuleResponse{}
	_body, _err := client.DeleteGatewayCircuitBreakerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a domain name from a gateway.
//
// @param request - DeleteGatewayDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayDomainResponse
func (client *Client) DeleteGatewayDomainWithOptions(request *DeleteGatewayDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayDomain"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a domain name from a gateway.
//
// @param request - DeleteGatewayDomainRequest
//
// @return DeleteGatewayDomainResponse
func (client *Client) DeleteGatewayDomain(request *DeleteGatewayDomainRequest) (_result *DeleteGatewayDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayDomainResponse{}
	_body, _err := client.DeleteGatewayDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关路由流控规则
//
// @param request - DeleteGatewayFlowRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayFlowRuleResponse
func (client *Client) DeleteGatewayFlowRuleWithOptions(request *DeleteGatewayFlowRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayFlowRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayFlowRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关路由流控规则
//
// @param request - DeleteGatewayFlowRuleRequest
//
// @return DeleteGatewayFlowRuleResponse
func (client *Client) DeleteGatewayFlowRule(request *DeleteGatewayFlowRuleRequest) (_result *DeleteGatewayFlowRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayFlowRuleResponse{}
	_body, _err := client.DeleteGatewayFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关路由隔离规则
//
// @param request - DeleteGatewayIsolationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayIsolationRuleResponse
func (client *Client) DeleteGatewayIsolationRuleWithOptions(request *DeleteGatewayIsolationRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayIsolationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayIsolationRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关路由隔离规则
//
// @param request - DeleteGatewayIsolationRuleRequest
//
// @return DeleteGatewayIsolationRuleResponse
func (client *Client) DeleteGatewayIsolationRule(request *DeleteGatewayIsolationRuleRequest) (_result *DeleteGatewayIsolationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayIsolationRuleResponse{}
	_body, _err := client.DeleteGatewayIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a route from a gateway.
//
// @param request - DeleteGatewayRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayRouteResponse
func (client *Client) DeleteGatewayRouteWithOptions(request *DeleteGatewayRouteRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayRoute"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a route from a gateway.
//
// @param request - DeleteGatewayRouteRequest
//
// @return DeleteGatewayRouteResponse
func (client *Client) DeleteGatewayRoute(request *DeleteGatewayRouteRequest) (_result *DeleteGatewayRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayRouteResponse{}
	_body, _err := client.DeleteGatewayRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a service from a gateway.
//
// @param request - DeleteGatewayServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayServiceResponse
func (client *Client) DeleteGatewayServiceWithOptions(request *DeleteGatewayServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayService"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service from a gateway.
//
// @param request - DeleteGatewayServiceRequest
//
// @return DeleteGatewayServiceResponse
func (client *Client) DeleteGatewayService(request *DeleteGatewayServiceRequest) (_result *DeleteGatewayServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayServiceResponse{}
	_body, _err := client.DeleteGatewayServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a service version from a gateway.
//
// @param request - DeleteGatewayServiceVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayServiceVersionResponse
func (client *Client) DeleteGatewayServiceVersionWithOptions(request *DeleteGatewayServiceVersionRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewayServiceVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayServiceVersion"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service version from a gateway.
//
// @param request - DeleteGatewayServiceVersionRequest
//
// @return DeleteGatewayServiceVersionResponse
func (client *Client) DeleteGatewayServiceVersion(request *DeleteGatewayServiceVersionRequest) (_result *DeleteGatewayServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewayServiceVersionResponse{}
	_body, _err := client.DeleteGatewayServiceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the Server Load Balancer (SLB) instance that is associated with a gateway.
//
// @param request - DeleteGatewaySlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewaySlbResponse
func (client *Client) DeleteGatewaySlbWithOptions(request *DeleteGatewaySlbRequest, runtime *dara.RuntimeOptions) (_result *DeleteGatewaySlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DeleteSlb) {
		query["DeleteSlb"] = request.DeleteSlb
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.SlbId) {
		query["SlbId"] = request.SlbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewaySlb"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewaySlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the Server Load Balancer (SLB) instance that is associated with a gateway.
//
// @param request - DeleteGatewaySlbRequest
//
// @return DeleteGatewaySlbResponse
func (client *Client) DeleteGatewaySlb(request *DeleteGatewaySlbRequest) (_result *DeleteGatewaySlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGatewaySlbResponse{}
	_body, _err := client.DeleteGatewaySlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除隔离规则
//
// @param tmpReq - DeleteIsolationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIsolationRulesResponse
func (client *Client) DeleteIsolationRulesWithOptions(tmpReq *DeleteIsolationRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteIsolationRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteIsolationRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIsolationRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIsolationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除隔离规则
//
// @param request - DeleteIsolationRulesRequest
//
// @return DeleteIsolationRulesResponse
func (client *Client) DeleteIsolationRules(request *DeleteIsolationRulesRequest) (_result *DeleteIsolationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIsolationRulesResponse{}
	_body, _err := client.DeleteIsolationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a migration task.
//
// @param request - DeleteMigrationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMigrationTaskResponse
func (client *Client) DeleteMigrationTaskWithOptions(request *DeleteMigrationTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteMigrationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMigrationTask"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMigrationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a migration task.
//
// @param request - DeleteMigrationTaskRequest
//
// @return DeleteMigrationTaskResponse
func (client *Client) DeleteMigrationTask(request *DeleteMigrationTaskRequest) (_result *DeleteMigrationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMigrationTaskResponse{}
	_body, _err := client.DeleteMigrationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete specified Nacos configuration
//
// Description:
//
// > The current OpenAPI is not the Nacos-SDK API. For information related to the Nacos-SDK API, please refer to the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - DeleteNacosConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNacosConfigResponse
func (client *Client) DeleteNacosConfigWithOptions(request *DeleteNacosConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteNacosConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Beta) {
		query["Beta"] = request.Beta
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNacosConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete specified Nacos configuration
//
// Description:
//
// > The current OpenAPI is not the Nacos-SDK API. For information related to the Nacos-SDK API, please refer to the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - DeleteNacosConfigRequest
//
// @return DeleteNacosConfigResponse
func (client *Client) DeleteNacosConfig(request *DeleteNacosConfigRequest) (_result *DeleteNacosConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNacosConfigResponse{}
	_body, _err := client.DeleteNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple Nacos configurations at a time.
//
// Description:
//
// >  The current API operation is not provided in Nacos SDK. For more information about the Nacos-SDK API, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - DeleteNacosConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNacosConfigsResponse
func (client *Client) DeleteNacosConfigsWithOptions(request *DeleteNacosConfigsRequest, runtime *dara.RuntimeOptions) (_result *DeleteNacosConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNacosConfigs"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNacosConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple Nacos configurations at a time.
//
// Description:
//
// >  The current API operation is not provided in Nacos SDK. For more information about the Nacos-SDK API, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - DeleteNacosConfigsRequest
//
// @return DeleteNacosConfigsResponse
func (client *Client) DeleteNacosConfigs(request *DeleteNacosConfigsRequest) (_result *DeleteNacosConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNacosConfigsResponse{}
	_body, _err := client.DeleteNacosConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a persistent application instance from a Microservices Engine (MSE) Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - DeleteNacosInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNacosInstanceResponse
func (client *Client) DeleteNacosInstanceWithOptions(request *DeleteNacosInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNacosInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.Ephemeral) {
		query["Ephemeral"] = request.Ephemeral
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNacosInstance"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNacosInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a persistent application instance from a Microservices Engine (MSE) Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - DeleteNacosInstanceRequest
//
// @return DeleteNacosInstanceResponse
func (client *Client) DeleteNacosInstance(request *DeleteNacosInstanceRequest) (_result *DeleteNacosInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNacosInstanceResponse{}
	_body, _err := client.DeleteNacosInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除一个MCP Server
//
// @param request - DeleteNacosMcpServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNacosMcpServerResponse
func (client *Client) DeleteNacosMcpServerWithOptions(request *DeleteNacosMcpServerRequest, runtime *dara.RuntimeOptions) (_result *DeleteNacosMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.McpServerId) {
		query["McpServerId"] = request.McpServerId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNacosMcpServer"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNacosMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个MCP Server
//
// @param request - DeleteNacosMcpServerRequest
//
// @return DeleteNacosMcpServerResponse
func (client *Client) DeleteNacosMcpServer(request *DeleteNacosMcpServerRequest) (_result *DeleteNacosMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNacosMcpServerResponse{}
	_body, _err := client.DeleteNacosMcpServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Nacos service.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - DeleteNacosServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNacosServiceResponse
func (client *Client) DeleteNacosServiceWithOptions(request *DeleteNacosServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNacosServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNacosService"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNacosServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Nacos service.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - DeleteNacosServiceRequest
//
// @return DeleteNacosServiceResponse
func (client *Client) DeleteNacosService(request *DeleteNacosServiceRequest) (_result *DeleteNacosServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNacosServiceResponse{}
	_body, _err := client.DeleteNacosServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除MSE命名空间
//
// @param request - DeleteNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNamespace"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除MSE命名空间
//
// @param request - DeleteNamespaceRequest
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a plug-in configuration.
//
// @param request - DeletePluginConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePluginConfigResponse
func (client *Client) DeletePluginConfigWithOptions(request *DeletePluginConfigRequest, runtime *dara.RuntimeOptions) (_result *DeletePluginConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.PluginConfigId) {
		query["PluginConfigId"] = request.PluginConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePluginConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePluginConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a plug-in configuration.
//
// @param request - DeletePluginConfigRequest
//
// @return DeletePluginConfigResponse
func (client *Client) DeletePluginConfig(request *DeletePluginConfigRequest) (_result *DeletePluginConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePluginConfigResponse{}
	_body, _err := client.DeletePluginConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a security group rule from a gateway.
//
// @param request - DeleteSecurityGroupRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityGroupRuleResponse
func (client *Client) DeleteSecurityGroupRuleWithOptions(request *DeleteSecurityGroupRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityGroupRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CascadingDelete) {
		query["CascadingDelete"] = request.CascadingDelete
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityGroupRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityGroupRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a security group rule from a gateway.
//
// @param request - DeleteSecurityGroupRuleRequest
//
// @return DeleteSecurityGroupRuleResponse
func (client *Client) DeleteSecurityGroupRule(request *DeleteSecurityGroupRuleRequest) (_result *DeleteSecurityGroupRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecurityGroupRuleResponse{}
	_body, _err := client.DeleteSecurityGroupRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a gateway service source.
//
// @param request - DeleteServiceSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceSourceResponse
func (client *Client) DeleteServiceSourceWithOptions(request *DeleteServiceSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceSource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a gateway service source.
//
// @param request - DeleteServiceSourceRequest
//
// @return DeleteServiceSourceResponse
func (client *Client) DeleteServiceSource(request *DeleteServiceSourceRequest) (_result *DeleteServiceSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteServiceSourceResponse{}
	_body, _err := client.DeleteServiceSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a lane.
//
// @param request - DeleteSwimmingLaneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSwimmingLaneResponse
func (client *Client) DeleteSwimmingLaneWithOptions(request *DeleteSwimmingLaneRequest, runtime *dara.RuntimeOptions) (_result *DeleteSwimmingLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.LaneId) {
		query["LaneId"] = request.LaneId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSwimmingLane"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSwimmingLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lane.
//
// @param request - DeleteSwimmingLaneRequest
//
// @return DeleteSwimmingLaneResponse
func (client *Client) DeleteSwimmingLane(request *DeleteSwimmingLaneRequest) (_result *DeleteSwimmingLaneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSwimmingLaneResponse{}
	_body, _err := client.DeleteSwimmingLaneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a lane group.
//
// @param request - DeleteSwimmingLaneGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSwimmingLaneGroupResponse
func (client *Client) DeleteSwimmingLaneGroupWithOptions(request *DeleteSwimmingLaneGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteSwimmingLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSwimmingLaneGroup"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSwimmingLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lane group.
//
// @param request - DeleteSwimmingLaneGroupRequest
//
// @return DeleteSwimmingLaneGroupResponse
func (client *Client) DeleteSwimmingLaneGroup(request *DeleteSwimmingLaneGroupRequest) (_result *DeleteSwimmingLaneGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSwimmingLaneGroupResponse{}
	_body, _err := client.DeleteSwimmingLaneGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除热点参数防护规则（HTTP 请求）
//
// @param request - DeleteWebFlowRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebFlowRulesResponse
func (client *Client) DeleteWebFlowRulesWithOptions(request *DeleteWebFlowRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebFlowRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebFlowRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebFlowRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除热点参数防护规则（HTTP 请求）
//
// @param request - DeleteWebFlowRulesRequest
//
// @return DeleteWebFlowRulesResponse
func (client *Client) DeleteWebFlowRules(request *DeleteWebFlowRulesRequest) (_result *DeleteWebFlowRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWebFlowRulesResponse{}
	_body, _err := client.DeleteWebFlowRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a ZooKeeper node.
//
// @param request - DeleteZnodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteZnodeResponse
func (client *Client) DeleteZnodeWithOptions(request *DeleteZnodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteZnodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteZnode"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteZnodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a ZooKeeper node.
//
// @param request - DeleteZnodeRequest
//
// @return DeleteZnodeResponse
func (client *Client) DeleteZnode(request *DeleteZnodeRequest) (_result *DeleteZnodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteZnodeResponse{}
	_body, _err := client.DeleteZnodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables HTTP/2 for negotiation between the server and client. The modification takes effect in one to two minutes.
//
// @param request - EnableHttp2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHttp2Response
func (client *Client) EnableHttp2WithOptions(request *EnableHttp2Request, runtime *dara.RuntimeOptions) (_result *EnableHttp2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.EnableHttp2) {
		query["EnableHttp2"] = request.EnableHttp2
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableHttp2"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableHttp2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables HTTP/2 for negotiation between the server and client. The modification takes effect in one to two minutes.
//
// @param request - EnableHttp2Request
//
// @return EnableHttp2Response
func (client *Client) EnableHttp2(request *EnableHttp2Request) (_result *EnableHttp2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableHttp2Response{}
	_body, _err := client.EnableHttp2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the proxy protocol. When an NLB instance is used as an ingress, you cannot obtain the real IP address of the client if you do not enable the proxy protocol. After you enable the proxy protocol, non-proxy requests are not adversely affected.
//
// @param request - EnableProxyProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableProxyProtocolResponse
func (client *Client) EnableProxyProtocolWithOptions(request *EnableProxyProtocolRequest, runtime *dara.RuntimeOptions) (_result *EnableProxyProtocolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.EnableProxyProtocol) {
		query["EnableProxyProtocol"] = request.EnableProxyProtocol
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableProxyProtocol"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableProxyProtocolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the proxy protocol. When an NLB instance is used as an ingress, you cannot obtain the real IP address of the client if you do not enable the proxy protocol. After you enable the proxy protocol, non-proxy requests are not adversely affected.
//
// @param request - EnableProxyProtocolRequest
//
// @return EnableProxyProtocolResponse
func (client *Client) EnableProxyProtocol(request *EnableProxyProtocolRequest) (_result *EnableProxyProtocolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableProxyProtocolResponse{}
	_body, _err := client.EnableProxyProtocolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Exports specified Nacos configurations.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ExportNacosConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportNacosConfigResponse
func (client *Client) ExportNacosConfigWithOptions(request *ExportNacosConfigRequest, runtime *dara.RuntimeOptions) (_result *ExportNacosConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.DataIds) {
		query["DataIds"] = request.DataIds
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportNacosConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports specified Nacos configurations.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ExportNacosConfigRequest
//
// @return ExportNacosConfigResponse
func (client *Client) ExportNacosConfig(request *ExportNacosConfigRequest) (_result *ExportNacosConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportNacosConfigResponse{}
	_body, _err := client.ExportNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a task to export ZooKeeper data.
//
// Description:
//
// Only one task can run at a time.
//
// @param request - ExportZookeeperDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportZookeeperDataResponse
func (client *Client) ExportZookeeperDataWithOptions(request *ExportZookeeperDataRequest, runtime *dara.RuntimeOptions) (_result *ExportZookeeperDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ExportType) {
		query["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportZookeeperData"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportZookeeperDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a task to export ZooKeeper data.
//
// Description:
//
// Only one task can run at a time.
//
// @param request - ExportZookeeperDataRequest
//
// @return ExportZookeeperDataResponse
func (client *Client) ExportZookeeperData(request *ExportZookeeperDataRequest) (_result *ExportZookeeperDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportZookeeperDataResponse{}
	_body, _err := client.ExportZookeeperDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the rules for graceful start and shutdown.
//
// Description:
//
// You can call this operation to query the rules for graceful start and shutdown.
//
// @param request - FetchLosslessRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FetchLosslessRuleListResponse
func (client *Client) FetchLosslessRuleListWithOptions(request *FetchLosslessRuleListRequest, runtime *dara.RuntimeOptions) (_result *FetchLosslessRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FetchLosslessRuleList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FetchLosslessRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the rules for graceful start and shutdown.
//
// Description:
//
// You can call this operation to query the rules for graceful start and shutdown.
//
// @param request - FetchLosslessRuleListRequest
//
// @return FetchLosslessRuleListResponse
func (client *Client) FetchLosslessRuleList(request *FetchLosslessRuleListRequest) (_result *FetchLosslessRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FetchLosslessRuleListResponse{}
	_body, _err := client.FetchLosslessRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of blacklists and whitelists of a gateway.
//
// @param tmpReq - GatewayBlackWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GatewayBlackWhiteListResponse
func (client *Client) GatewayBlackWhiteListWithOptions(tmpReq *GatewayBlackWhiteListRequest, runtime *dara.RuntimeOptions) (_result *GatewayBlackWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GatewayBlackWhiteListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParams) {
		request.FilterParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParams, dara.String("FilterParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DescSort) {
		query["DescSort"] = request.DescSort
	}

	if !dara.IsNil(request.FilterParamsShrink) {
		query["FilterParams"] = request.FilterParamsShrink
	}

	if !dara.IsNil(request.OrderItem) {
		query["OrderItem"] = request.OrderItem
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GatewayBlackWhiteList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GatewayBlackWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of blacklists and whitelists of a gateway.
//
// @param request - GatewayBlackWhiteListRequest
//
// @return GatewayBlackWhiteListResponse
func (client *Client) GatewayBlackWhiteList(request *GatewayBlackWhiteListRequest) (_result *GatewayBlackWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GatewayBlackWhiteListResponse{}
	_body, _err := client.GatewayBlackWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about canary release for messaging of an application.
//
// @param request - GetAppMessageQueueRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppMessageQueueRouteResponse
func (client *Client) GetAppMessageQueueRouteWithOptions(request *GetAppMessageQueueRouteRequest, runtime *dara.RuntimeOptions) (_result *GetAppMessageQueueRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppMessageQueueRoute"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppMessageQueueRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about canary release for messaging of an application.
//
// @param request - GetAppMessageQueueRouteRequest
//
// @return GetAppMessageQueueRouteResponse
func (client *Client) GetAppMessageQueueRoute(request *GetAppMessageQueueRouteRequest) (_result *GetAppMessageQueueRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppMessageQueueRouteResponse{}
	_body, _err := client.GetAppMessageQueueRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of microservice application instances.
//
// @param request - GetApplicationInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationInstanceListResponse
func (client *Client) GetApplicationInstanceListWithOptions(request *GetApplicationInstanceListRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationInstanceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationInstanceList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of microservice application instances.
//
// @param request - GetApplicationInstanceListRequest
//
// @return GetApplicationInstanceListResponse
func (client *Client) GetApplicationInstanceList(request *GetApplicationInstanceListRequest) (_result *GetApplicationInstanceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationInstanceListResponse{}
	_body, _err := client.GetApplicationInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list of applications.
//
// @param tmpReq - GetApplicationListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationListResponse
func (client *Client) GetApplicationListWithOptions(tmpReq *GetApplicationListRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetApplicationListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SentinelEnable) {
		query["SentinelEnable"] = request.SentinelEnable
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SwitchEnable) {
		query["SwitchEnable"] = request.SwitchEnable
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of applications.
//
// @param request - GetApplicationListRequest
//
// @return GetApplicationListResponse
func (client *Client) GetApplicationList(request *GetApplicationListRequest) (_result *GetApplicationListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationListResponse{}
	_body, _err := client.GetApplicationListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetBlackWhiteList is deprecated, please use mse::2019-05-31::GatewayBlackWhiteList instead.
//
// Summary:
//
// Queries the blacklist or whitelist of a gateway.
//
// @param request - GetBlackWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBlackWhiteListResponse
func (client *Client) GetBlackWhiteListWithOptions(request *GetBlackWhiteListRequest, runtime *dara.RuntimeOptions) (_result *GetBlackWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.IsWhite) {
		query["IsWhite"] = request.IsWhite
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBlackWhiteList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBlackWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetBlackWhiteList is deprecated, please use mse::2019-05-31::GatewayBlackWhiteList instead.
//
// Summary:
//
// Queries the blacklist or whitelist of a gateway.
//
// @param request - GetBlackWhiteListRequest
//
// @return GetBlackWhiteListResponse
// Deprecated
func (client *Client) GetBlackWhiteList(request *GetBlackWhiteListRequest) (_result *GetBlackWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBlackWhiteListResponse{}
	_body, _err := client.GetBlackWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the namespaces of a Nacos instance.
//
// @param request - GetEngineNamepaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEngineNamepaceResponse
func (client *Client) GetEngineNamepaceWithOptions(request *GetEngineNamepaceRequest, runtime *dara.RuntimeOptions) (_result *GetEngineNamepaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEngineNamepace"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEngineNamepaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the namespaces of a Nacos instance.
//
// @param request - GetEngineNamepaceRequest
//
// @return GetEngineNamepaceResponse
func (client *Client) GetEngineNamepace(request *GetEngineNamepaceRequest) (_result *GetEngineNamepaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEngineNamepaceResponse{}
	_body, _err := client.GetEngineNamepaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the basic information about a gateway, such as the virtual private cloud (VPC) and vSwitch to which the gateway belongs.
//
// @param request - GetGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayResponse
func (client *Client) GetGatewayWithOptions(request *GetGatewayRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGateway"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the basic information about a gateway, such as the virtual private cloud (VPC) and vSwitch to which the gateway belongs.
//
// @param request - GetGatewayRequest
//
// @return GetGatewayResponse
func (client *Client) GetGateway(request *GetGatewayRequest) (_result *GetGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayResponse{}
	_body, _err := client.GetGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the consumer on which a gateway performs authentication operations.
//
// @param request - GetGatewayAuthConsumerDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayAuthConsumerDetailResponse
func (client *Client) GetGatewayAuthConsumerDetailWithOptions(request *GetGatewayAuthConsumerDetailRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayAuthConsumerDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayAuthConsumerDetail"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayAuthConsumerDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the consumer on which a gateway performs authentication operations.
//
// @param request - GetGatewayAuthConsumerDetailRequest
//
// @return GetGatewayAuthConsumerDetailResponse
func (client *Client) GetGatewayAuthConsumerDetail(request *GetGatewayAuthConsumerDetailRequest) (_result *GetGatewayAuthConsumerDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayAuthConsumerDetailResponse{}
	_body, _err := client.GetGatewayAuthConsumerDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询网关认证详情
//
// @param request - GetGatewayAuthDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayAuthDetailResponse
func (client *Client) GetGatewayAuthDetailWithOptions(request *GetGatewayAuthDetailRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayAuthDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayAuthDetail"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayAuthDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网关认证详情
//
// @param request - GetGatewayAuthDetailRequest
//
// @return GetGatewayAuthDetailResponse
func (client *Client) GetGatewayAuthDetail(request *GetGatewayAuthDetailRequest) (_result *GetGatewayAuthDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayAuthDetailResponse{}
	_body, _err := client.GetGatewayAuthDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网关全局配置
//
// @param request - GetGatewayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayConfigResponse
func (client *Client) GetGatewayConfigWithOptions(request *GetGatewayConfigRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网关全局配置
//
// @param request - GetGatewayConfigRequest
//
// @return GetGatewayConfigResponse
func (client *Client) GetGatewayConfig(request *GetGatewayConfigRequest) (_result *GetGatewayConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayConfigResponse{}
	_body, _err := client.GetGatewayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a domain name associated with a gateway.
//
// @param request - GetGatewayDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayDomainDetailResponse
func (client *Client) GetGatewayDomainDetailWithOptions(request *GetGatewayDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayDomainDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayDomainDetail"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a domain name associated with a gateway.
//
// @param request - GetGatewayDomainDetailRequest
//
// @return GetGatewayDomainDetailResponse
func (client *Client) GetGatewayDomainDetail(request *GetGatewayDomainDetailRequest) (_result *GetGatewayDomainDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayDomainDetailResponse{}
	_body, _err := client.GetGatewayDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the global parameters of a gateway.
//
// @param request - GetGatewayOptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayOptionResponse
func (client *Client) GetGatewayOptionWithOptions(request *GetGatewayOptionRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayOptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayOption"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the global parameters of a gateway.
//
// @param request - GetGatewayOptionRequest
//
// @return GetGatewayOptionResponse
func (client *Client) GetGatewayOption(request *GetGatewayOptionRequest) (_result *GetGatewayOptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayOptionResponse{}
	_body, _err := client.GetGatewayOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a route for a gateway.
//
// @param request - GetGatewayRouteDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayRouteDetailResponse
func (client *Client) GetGatewayRouteDetailWithOptions(request *GetGatewayRouteDetailRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayRouteDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayRouteDetail"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayRouteDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a route for a gateway.
//
// @param request - GetGatewayRouteDetailRequest
//
// @return GetGatewayRouteDetailResponse
func (client *Client) GetGatewayRouteDetail(request *GetGatewayRouteDetailRequest) (_result *GetGatewayRouteDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayRouteDetailResponse{}
	_body, _err := client.GetGatewayRouteDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a service.
//
// @param request - GetGatewayServiceDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayServiceDetailResponse
func (client *Client) GetGatewayServiceDetailWithOptions(request *GetGatewayServiceDetailRequest, runtime *dara.RuntimeOptions) (_result *GetGatewayServiceDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayServiceDetail"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayServiceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a service.
//
// @param request - GetGatewayServiceDetailRequest
//
// @return GetGatewayServiceDetailResponse
func (client *Client) GetGatewayServiceDetail(request *GetGatewayServiceDetailRequest) (_result *GetGatewayServiceDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGatewayServiceDetailResponse{}
	_body, _err := client.GetGatewayServiceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a Container Service for Kubernetes (ACK) cluster for which Microservices Governance is enabled.
//
// @param request - GetGovernanceKubernetesClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGovernanceKubernetesClusterResponse
func (client *Client) GetGovernanceKubernetesClusterWithOptions(request *GetGovernanceKubernetesClusterRequest, runtime *dara.RuntimeOptions) (_result *GetGovernanceKubernetesClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGovernanceKubernetesCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGovernanceKubernetesClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Container Service for Kubernetes (ACK) cluster for which Microservices Governance is enabled.
//
// @param request - GetGovernanceKubernetesClusterRequest
//
// @return GetGovernanceKubernetesClusterResponse
func (client *Client) GetGovernanceKubernetesCluster(request *GetGovernanceKubernetesClusterRequest) (_result *GetGovernanceKubernetesClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGovernanceKubernetesClusterResponse{}
	_body, _err := client.GetGovernanceKubernetesClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the maximum version number to which the current version can be upgraded.
//
// @param request - GetImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
func (client *Client) GetImageWithOptions(request *GetImageRequest, runtime *dara.RuntimeOptions) (_result *GetImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.VersionCode) {
		query["VersionCode"] = request.VersionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImage"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum version number to which the current version can be upgraded.
//
// @param request - GetImageRequest
//
// @return GetImageResponse
func (client *Client) GetImage(request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the URL that is used to upload a configuration file when you import the configuration file into a Microservices Engine (MSE) Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).\\n
//
// @param request - GetImportFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImportFileUrlResponse
func (client *Client) GetImportFileUrlWithOptions(request *GetImportFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetImportFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImportFileUrl"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImportFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the URL that is used to upload a configuration file when you import the configuration file into a Microservices Engine (MSE) Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).\\n
//
// @param request - GetImportFileUrlRequest
//
// @return GetImportFileUrlResponse
func (client *Client) GetImportFileUrl(request *GetImportFileUrlRequest) (_result *GetImportFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImportFileUrlResponse{}
	_body, _err := client.GetImportFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains sources of all Container Service for Kubernetes (ACK) services in a gateway.
//
// @param request - GetKubernetesSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKubernetesSourceResponse
func (client *Client) GetKubernetesSourceWithOptions(request *GetKubernetesSourceRequest, runtime *dara.RuntimeOptions) (_result *GetKubernetesSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.IsAll) {
		query["IsAll"] = request.IsAll
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKubernetesSource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKubernetesSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains sources of all Container Service for Kubernetes (ACK) services in a gateway.
//
// @param request - GetKubernetesSourceRequest
//
// @return GetKubernetesSourceResponse
func (client *Client) GetKubernetesSource(request *GetKubernetesSourceRequest) (_result *GetKubernetesSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetKubernetesSourceResponse{}
	_body, _err := client.GetKubernetesSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取同AZ路由规则
//
// @param request - GetLocalityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLocalityRuleResponse
func (client *Client) GetLocalityRuleWithOptions(request *GetLocalityRuleRequest, runtime *dara.RuntimeOptions) (_result *GetLocalityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLocalityRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLocalityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取同AZ路由规则
//
// @param request - GetLocalityRuleRequest
//
// @return GetLocalityRuleResponse
func (client *Client) GetLocalityRule(request *GetLocalityRuleRequest) (_result *GetLocalityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLocalityRuleResponse{}
	_body, _err := client.GetLocalityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rules for graceful start and shutdown of an application.
//
// Description:
//
// You can call this operation to query the rules for graceful start and shutdown of an application.
//
// You can query the rules for graceful start and shutdown of an application preferentially by using the AppId parameter.
//
// If the AppId parameter is left empty, you can use the RegionId, Namespace, and AppName parameters to query the rules for graceful start and shutdown of an application.
//
// @param request - GetLosslessRuleByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLosslessRuleByAppResponse
func (client *Client) GetLosslessRuleByAppWithOptions(request *GetLosslessRuleByAppRequest, runtime *dara.RuntimeOptions) (_result *GetLosslessRuleByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLosslessRuleByApp"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLosslessRuleByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rules for graceful start and shutdown of an application.
//
// Description:
//
// You can call this operation to query the rules for graceful start and shutdown of an application.
//
// You can query the rules for graceful start and shutdown of an application preferentially by using the AppId parameter.
//
// If the AppId parameter is left empty, you can use the RegionId, Namespace, and AppName parameters to query the rules for graceful start and shutdown of an application.
//
// @param request - GetLosslessRuleByAppRequest
//
// @return GetLosslessRuleByAppResponse
func (client *Client) GetLosslessRuleByApp(request *GetLosslessRuleByAppRequest) (_result *GetLosslessRuleByAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLosslessRuleByAppResponse{}
	_body, _err := client.GetLosslessRuleByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the MSE feature switch.
//
// @param request - GetMseFeatureSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMseFeatureSwitchResponse
func (client *Client) GetMseFeatureSwitchWithOptions(request *GetMseFeatureSwitchRequest, runtime *dara.RuntimeOptions) (_result *GetMseFeatureSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMseFeatureSwitch"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMseFeatureSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the MSE feature switch.
//
// @param request - GetMseFeatureSwitchRequest
//
// @return GetMseFeatureSwitchResponse
func (client *Client) GetMseFeatureSwitch(request *GetMseFeatureSwitchRequest) (_result *GetMseFeatureSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMseFeatureSwitchResponse{}
	_body, _err := client.GetMseFeatureSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the existing Microservices Engine (MSE) Nacos instances that are service sources of a gateway.
//
// @param request - GetMseSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMseSourceResponse
func (client *Client) GetMseSourceWithOptions(request *GetMseSourceRequest, runtime *dara.RuntimeOptions) (_result *GetMseSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMseSource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMseSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the existing Microservices Engine (MSE) Nacos instances that are service sources of a gateway.
//
// @param request - GetMseSourceRequest
//
// @return GetMseSourceResponse
func (client *Client) GetMseSource(request *GetMseSourceRequest) (_result *GetMseSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMseSourceResponse{}
	_body, _err := client.GetMseSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Nacos Configuration
//
// Description:
//
// > This OpenAPI is not the Nacos-SDK API. For information related to the Nacos-SDK API, please refer to the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - GetNacosConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNacosConfigResponse
func (client *Client) GetNacosConfigWithOptions(request *GetNacosConfigRequest, runtime *dara.RuntimeOptions) (_result *GetNacosConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Beta) {
		query["Beta"] = request.Beta
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNacosConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Nacos Configuration
//
// Description:
//
// > This OpenAPI is not the Nacos-SDK API. For information related to the Nacos-SDK API, please refer to the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - GetNacosConfigRequest
//
// @return GetNacosConfigResponse
func (client *Client) GetNacosConfig(request *GetNacosConfigRequest) (_result *GetNacosConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNacosConfigResponse{}
	_body, _err := client.GetNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical details of Nacos configuration changes.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - GetNacosHistoryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNacosHistoryConfigResponse
func (client *Client) GetNacosHistoryConfigWithOptions(request *GetNacosHistoryConfigRequest, runtime *dara.RuntimeOptions) (_result *GetNacosHistoryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Nid) {
		query["Nid"] = request.Nid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNacosHistoryConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNacosHistoryConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical details of Nacos configuration changes.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - GetNacosHistoryConfigRequest
//
// @return GetNacosHistoryConfigResponse
func (client *Client) GetNacosHistoryConfig(request *GetNacosHistoryConfigRequest) (_result *GetNacosHistoryConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNacosHistoryConfigResponse{}
	_body, _err := client.GetNacosHistoryConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取MCP Server的详情
//
// @param request - GetNacosMcpServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNacosMcpServerResponse
func (client *Client) GetNacosMcpServerWithOptions(request *GetNacosMcpServerRequest, runtime *dara.RuntimeOptions) (_result *GetNacosMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.McpServerId) {
		query["McpServerId"] = request.McpServerId
	}

	if !dara.IsNil(request.McpServerVersion) {
		query["McpServerVersion"] = request.McpServerVersion
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNacosMcpServer"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNacosMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取MCP Server的详情
//
// @param request - GetNacosMcpServerRequest
//
// @return GetNacosMcpServerResponse
func (client *Client) GetNacosMcpServer(request *GetNacosMcpServerRequest) (_result *GetNacosMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNacosMcpServerResponse{}
	_body, _err := client.GetNacosMcpServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about service governance.
//
// Description:
//
// You can call this operation to query overview information about service governance.
//
// @param request - GetOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOverviewResponse
func (client *Client) GetOverviewWithOptions(request *GetOverviewRequest, runtime *dara.RuntimeOptions) (_result *GetOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOverview"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about service governance.
//
// Description:
//
// You can call this operation to query overview information about service governance.
//
// @param request - GetOverviewRequest
//
// @return GetOverviewResponse
func (client *Client) GetOverview(request *GetOverviewRequest) (_result *GetOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOverviewResponse{}
	_body, _err := client.GetOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains plug-in configurations.
//
// @param request - GetPluginConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPluginConfigResponse
func (client *Client) GetPluginConfigWithOptions(request *GetPluginConfigRequest, runtime *dara.RuntimeOptions) (_result *GetPluginConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPluginConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPluginConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains plug-in configurations.
//
// @param request - GetPluginConfigRequest
//
// @return GetPluginConfigResponse
func (client *Client) GetPluginConfig(request *GetPluginConfigRequest) (_result *GetPluginConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPluginConfigResponse{}
	_body, _err := client.GetPluginConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains plug-ins.
//
// @param request - GetPluginsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPluginsResponse
func (client *Client) GetPluginsWithOptions(request *GetPluginsRequest, runtime *dara.RuntimeOptions) (_result *GetPluginsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.EnableOnly) {
		query["EnableOnly"] = request.EnableOnly
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPlugins"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains plug-ins.
//
// @param request - GetPluginsRequest
//
// @return GetPluginsResponse
func (client *Client) GetPlugins(request *GetPluginsRequest) (_result *GetPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPluginsResponse{}
	_body, _err := client.GetPluginsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the services of an application.
//
// @param request - GetServiceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceListResponse
func (client *Client) GetServiceListWithOptions(request *GetServiceListRequest, runtime *dara.RuntimeOptions) (_result *GetServiceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the services of an application.
//
// @param request - GetServiceListRequest
//
// @return GetServiceListResponse
func (client *Client) GetServiceList(request *GetServiceListRequest) (_result *GetServiceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceListResponse{}
	_body, _err := client.GetServiceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the version of a microservices application.
//
// @param request - GetServiceListPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceListPageResponse
func (client *Client) GetServiceListPageWithOptions(request *GetServiceListPageRequest, runtime *dara.RuntimeOptions) (_result *GetServiceListPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceListPage"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceListPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version of a microservices application.
//
// @param request - GetServiceListPageRequest
//
// @return GetServiceListPageResponse
func (client *Client) GetServiceListPage(request *GetServiceListPageRequest) (_result *GetServiceListPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceListPageResponse{}
	_body, _err := client.GetServiceListPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of listeners for the destination service.
//
// @param request - GetServiceListenersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceListenersResponse
func (client *Client) GetServiceListenersWithOptions(request *GetServiceListenersRequest, runtime *dara.RuntimeOptions) (_result *GetServiceListenersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.HasIpCount) {
		query["HasIpCount"] = request.HasIpCount
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceListeners"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of listeners for the destination service.
//
// @param request - GetServiceListenersRequest
//
// @return GetServiceListenersResponse
func (client *Client) GetServiceListeners(request *GetServiceListenersRequest) (_result *GetServiceListenersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceListenersResponse{}
	_body, _err := client.GetServiceListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务接口列表
//
// @param request - GetServiceMethodPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceMethodPageResponse
func (client *Client) GetServiceMethodPageWithOptions(request *GetServiceMethodPageRequest, runtime *dara.RuntimeOptions) (_result *GetServiceMethodPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.MethodController) {
		query["MethodController"] = request.MethodController
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ServiceGroup) {
		query["ServiceGroup"] = request.ServiceGroup
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceMethodPage"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceMethodPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务接口列表
//
// @param request - GetServiceMethodPageRequest
//
// @return GetServiceMethodPageResponse
func (client *Client) GetServiceMethodPage(request *GetServiceMethodPageRequest) (_result *GetServiceMethodPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceMethodPageResponse{}
	_body, _err := client.GetServiceMethodPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains all tags in the current lane group.
//
// @param request - GetTagsBySwimmingLaneGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTagsBySwimmingLaneGroupIdResponse
func (client *Client) GetTagsBySwimmingLaneGroupIdWithOptions(request *GetTagsBySwimmingLaneGroupIdRequest, runtime *dara.RuntimeOptions) (_result *GetTagsBySwimmingLaneGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTagsBySwimmingLaneGroupId"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTagsBySwimmingLaneGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains all tags in the current lane group.
//
// @param request - GetTagsBySwimmingLaneGroupIdRequest
//
// @return GetTagsBySwimmingLaneGroupIdResponse
func (client *Client) GetTagsBySwimmingLaneGroupId(request *GetTagsBySwimmingLaneGroupIdRequest) (_result *GetTagsBySwimmingLaneGroupIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTagsBySwimmingLaneGroupIdResponse{}
	_body, _err := client.GetTagsBySwimmingLaneGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// mse-200-105
//
// @param request - GetZookeeperDataImportUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetZookeeperDataImportUrlResponse
func (client *Client) GetZookeeperDataImportUrlWithOptions(request *GetZookeeperDataImportUrlRequest, runtime *dara.RuntimeOptions) (_result *GetZookeeperDataImportUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetZookeeperDataImportUrl"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetZookeeperDataImportUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// mse-200-105
//
// @param request - GetZookeeperDataImportUrlRequest
//
// @return GetZookeeperDataImportUrlResponse
func (client *Client) GetZookeeperDataImportUrl(request *GetZookeeperDataImportUrlRequest) (_result *GetZookeeperDataImportUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetZookeeperDataImportUrlResponse{}
	_body, _err := client.GetZookeeperDataImportUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports Nacos configurations as a file.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ImportNacosConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportNacosConfigResponse
func (client *Client) ImportNacosConfigWithOptions(request *ImportNacosConfigRequest, runtime *dara.RuntimeOptions) (_result *ImportNacosConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportNacosConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports Nacos configurations as a file.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ImportNacosConfigRequest
//
// @return ImportNacosConfigResponse
func (client *Client) ImportNacosConfig(request *ImportNacosConfigRequest) (_result *ImportNacosConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportNacosConfigResponse{}
	_body, _err := client.ImportNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports services to a gateway.
//
// @param tmpReq - ImportServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportServicesResponse
func (client *Client) ImportServicesWithOptions(tmpReq *ImportServicesRequest, runtime *dara.RuntimeOptions) (_result *ImportServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportServicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServiceList) {
		request.ServiceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceList, dara.String("ServiceList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.FcAlias) {
		query["FcAlias"] = request.FcAlias
	}

	if !dara.IsNil(request.FcServiceName) {
		query["FcServiceName"] = request.FcServiceName
	}

	if !dara.IsNil(request.FcVersion) {
		query["FcVersion"] = request.FcVersion
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ServiceListShrink) {
		query["ServiceList"] = request.ServiceListShrink
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.TlsSetting) {
		query["TlsSetting"] = request.TlsSetting
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportServices"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports services to a gateway.
//
// @param request - ImportServicesRequest
//
// @return ImportServicesResponse
func (client *Client) ImportServices(request *ImportServicesRequest) (_result *ImportServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportServicesResponse{}
	_body, _err := client.ImportServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a task to import data from a destination URL to a Microservices Engine (MSE) ZooKeeper instance.
//
// Description:
//
// *Danger*	- This operation clears existing data. Exercise caution when you call this API operation.
//
// @param request - ImportZookeeperDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportZookeeperDataResponse
func (client *Client) ImportZookeeperDataWithOptions(request *ImportZookeeperDataRequest, runtime *dara.RuntimeOptions) (_result *ImportZookeeperDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportZookeeperData"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportZookeeperDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a task to import data from a destination URL to a Microservices Engine (MSE) ZooKeeper instance.
//
// Description:
//
// *Danger*	- This operation clears existing data. Exercise caution when you call this API operation.
//
// @param request - ImportZookeeperDataRequest
//
// @return ImportZookeeperDataResponse
func (client *Client) ImportZookeeperData(request *ImportZookeeperDataRequest) (_result *ImportZookeeperDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportZookeeperDataResponse{}
	_body, _err := client.ImportZookeeperDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户授权mseSLR
//
// @param request - InitializeServiceLinkRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeServiceLinkRoleResponse
func (client *Client) InitializeServiceLinkRoleWithOptions(request *InitializeServiceLinkRoleRequest, runtime *dara.RuntimeOptions) (_result *InitializeServiceLinkRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitializeServiceLinkRole"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeServiceLinkRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户授权mseSLR
//
// @param request - InitializeServiceLinkRoleRequest
//
// @return InitializeServiceLinkRoleResponse
func (client *Client) InitializeServiceLinkRole(request *InitializeServiceLinkRoleRequest) (_result *InitializeServiceLinkRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitializeServiceLinkRoleResponse{}
	_body, _err := client.InitializeServiceLinkRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries application instances that are registered with a Microservices Engine (MSE) Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListAnsInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnsInstancesResponse
func (client *Client) ListAnsInstancesWithOptions(request *ListAnsInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListAnsInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnsInstances"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries application instances that are registered with a Microservices Engine (MSE) Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListAnsInstancesRequest
//
// @return ListAnsInstancesResponse
func (client *Client) ListAnsInstances(request *ListAnsInstancesRequest) (_result *ListAnsInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAnsInstancesResponse{}
	_body, _err := client.ListAnsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the clusters of a Nacos service.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListAnsServiceClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnsServiceClustersResponse
func (client *Client) ListAnsServiceClustersWithOptions(request *ListAnsServiceClustersRequest, runtime *dara.RuntimeOptions) (_result *ListAnsServiceClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnsServiceClusters"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnsServiceClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the clusters of a Nacos service.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListAnsServiceClustersRequest
//
// @return ListAnsServiceClustersResponse
func (client *Client) ListAnsServiceClusters(request *ListAnsServiceClustersRequest) (_result *ListAnsServiceClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAnsServiceClustersResponse{}
	_body, _err := client.ListAnsServiceClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Nacos services.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListAnsServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnsServicesResponse
func (client *Client) ListAnsServicesWithOptions(request *ListAnsServicesRequest, runtime *dara.RuntimeOptions) (_result *ListAnsServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.HasIpCount) {
		query["HasIpCount"] = request.HasIpCount
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnsServices"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnsServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Nacos services.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListAnsServicesRequest
//
// @return ListAnsServicesResponse
func (client *Client) ListAnsServices(request *ListAnsServicesRequest) (_result *ListAnsServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAnsServicesResponse{}
	_body, _err := client.ListAnsServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListAppBySwimmingLaneGroupTag is deprecated, please use mse::2019-05-31::ListAppBySwimmingLaneGroupTags instead.
//
// Summary:
//
// Queries the applications in a lane group by tag.
//
// @param request - ListAppBySwimmingLaneGroupTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppBySwimmingLaneGroupTagResponse
func (client *Client) ListAppBySwimmingLaneGroupTagWithOptions(request *ListAppBySwimmingLaneGroupTagRequest, runtime *dara.RuntimeOptions) (_result *ListAppBySwimmingLaneGroupTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppBySwimmingLaneGroupTag"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppBySwimmingLaneGroupTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListAppBySwimmingLaneGroupTag is deprecated, please use mse::2019-05-31::ListAppBySwimmingLaneGroupTags instead.
//
// Summary:
//
// Queries the applications in a lane group by tag.
//
// @param request - ListAppBySwimmingLaneGroupTagRequest
//
// @return ListAppBySwimmingLaneGroupTagResponse
// Deprecated
func (client *Client) ListAppBySwimmingLaneGroupTag(request *ListAppBySwimmingLaneGroupTagRequest) (_result *ListAppBySwimmingLaneGroupTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppBySwimmingLaneGroupTagResponse{}
	_body, _err := client.ListAppBySwimmingLaneGroupTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists applications by tag in a specified lane group.
//
// @param tmpReq - ListAppBySwimmingLaneGroupTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppBySwimmingLaneGroupTagsResponse
func (client *Client) ListAppBySwimmingLaneGroupTagsWithOptions(tmpReq *ListAppBySwimmingLaneGroupTagsRequest, runtime *dara.RuntimeOptions) (_result *ListAppBySwimmingLaneGroupTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAppBySwimmingLaneGroupTagsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAppBySwimmingLaneGroupTags"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAppBySwimmingLaneGroupTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists applications by tag in a specified lane group.
//
// @param request - ListAppBySwimmingLaneGroupTagsRequest
//
// @return ListAppBySwimmingLaneGroupTagsResponse
func (client *Client) ListAppBySwimmingLaneGroupTags(request *ListAppBySwimmingLaneGroupTagsRequest) (_result *ListAppBySwimmingLaneGroupTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAppBySwimmingLaneGroupTagsResponse{}
	_body, _err := client.ListAppBySwimmingLaneGroupTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the routing rules of an application.
//
// @param request - ListApplicationsWithTagRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsWithTagRulesResponse
func (client *Client) ListApplicationsWithTagRulesWithOptions(request *ListApplicationsWithTagRulesRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsWithTagRulesResponse, _err error) {
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
		Action:      dara.String("ListApplicationsWithTagRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsWithTagRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routing rules of an application.
//
// @param request - ListApplicationsWithTagRulesRequest
//
// @return ListApplicationsWithTagRulesResponse
func (client *Client) ListApplicationsWithTagRules(request *ListApplicationsWithTagRulesRequest) (_result *ListApplicationsWithTagRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsWithTagRulesResponse{}
	_body, _err := client.ListApplicationsWithTagRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of service authentication rules.
//
// @param request - ListAuthPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthPolicyResponse
func (client *Client) ListAuthPolicyWithOptions(request *ListAuthPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListAuthPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthPolicy"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of service authentication rules.
//
// @param request - ListAuthPolicyRequest
//
// @return ListAuthPolicyResponse
func (client *Client) ListAuthPolicy(request *ListAuthPolicyRequest) (_result *ListAuthPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuthPolicyResponse{}
	_body, _err := client.ListAuthPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of circuit breaking rules.
//
// @param request - ListCircuitBreakerRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCircuitBreakerRulesResponse
func (client *Client) ListCircuitBreakerRulesWithOptions(request *ListCircuitBreakerRulesRequest, runtime *dara.RuntimeOptions) (_result *ListCircuitBreakerRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceSearchKey) {
		query["ResourceSearchKey"] = request.ResourceSearchKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCircuitBreakerRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCircuitBreakerRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of circuit breaking rules.
//
// @param request - ListCircuitBreakerRulesRequest
//
// @return ListCircuitBreakerRulesResponse
func (client *Client) ListCircuitBreakerRules(request *ListCircuitBreakerRulesRequest) (_result *ListCircuitBreakerRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCircuitBreakerRulesResponse{}
	_body, _err := client.ListCircuitBreakerRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available cluster connection types.
//
// @param request - ListClusterConnectionTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterConnectionTypesResponse
func (client *Client) ListClusterConnectionTypesWithOptions(request *ListClusterConnectionTypesRequest, runtime *dara.RuntimeOptions) (_result *ListClusterConnectionTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterConnectionTypes"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterConnectionTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available cluster connection types.
//
// @param request - ListClusterConnectionTypesRequest
//
// @return ListClusterConnectionTypesResponse
func (client *Client) ListClusterConnectionTypes(request *ListClusterConnectionTypesRequest) (_result *ListClusterConnectionTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClusterConnectionTypesResponse{}
	_body, _err := client.ListClusterConnectionTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains information about historical health check tasks.
//
// @param request - ListClusterHealthCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterHealthCheckTaskResponse
func (client *Client) ListClusterHealthCheckTaskWithOptions(request *ListClusterHealthCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *ListClusterHealthCheckTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterHealthCheckTask"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterHealthCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains information about historical health check tasks.
//
// @param request - ListClusterHealthCheckTaskRequest
//
// @return ListClusterHealthCheckTaskResponse
func (client *Client) ListClusterHealthCheckTask(request *ListClusterHealthCheckTaskRequest) (_result *ListClusterHealthCheckTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClusterHealthCheckTaskResponse{}
	_body, _err := client.ListClusterHealthCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the engine types that can be activated.
//
// @param request - ListClusterTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterTypesResponse
func (client *Client) ListClusterTypesWithOptions(request *ListClusterTypesRequest, runtime *dara.RuntimeOptions) (_result *ListClusterTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConnectType) {
		query["ConnectType"] = request.ConnectType
	}

	if !dara.IsNil(request.MseVersion) {
		query["MseVersion"] = request.MseVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterTypes"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the engine types that can be activated.
//
// @param request - ListClusterTypesRequest
//
// @return ListClusterTypesResponse
func (client *Client) ListClusterTypes(request *ListClusterTypesRequest) (_result *ListClusterTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClusterTypesResponse{}
	_body, _err := client.ListClusterTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about supported instance versions.
//
// @param request - ListClusterVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterVersionsResponse
func (client *Client) ListClusterVersionsWithOptions(request *ListClusterVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListClusterVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.MseVersion) {
		query["MseVersion"] = request.MseVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterVersions"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about supported instance versions.
//
// @param request - ListClusterVersionsRequest
//
// @return ListClusterVersionsResponse
func (client *Client) ListClusterVersions(request *ListClusterVersionsRequest) (_result *ListClusterVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClusterVersionsResponse{}
	_body, _err := client.ListClusterVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Microservices Engine (MSE) instances.
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterAliasName) {
		query["ClusterAliasName"] = request.ClusterAliasName
	}

	if !dara.IsNil(request.KeyId) {
		query["KeyId"] = request.KeyId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Microservices Engine (MSE) instances.
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the track data of a Nacos configuration center.
//
// @param request - ListConfigTrackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigTrackResponse
func (client *Client) ListConfigTrackWithOptions(request *ListConfigTrackRequest, runtime *dara.RuntimeOptions) (_result *ListConfigTrackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigTrack"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigTrackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the track data of a Nacos configuration center.
//
// @param request - ListConfigTrackRequest
//
// @return ListConfigTrackResponse
func (client *Client) ListConfigTrack(request *ListConfigTrackRequest) (_result *ListConfigTrackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConfigTrackResponse{}
	_body, _err := client.ListConfigTrackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the namespaces of a Nacos instance.
//
// @param request - ListEngineNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEngineNamespacesResponse
func (client *Client) ListEngineNamespacesWithOptions(request *ListEngineNamespacesRequest, runtime *dara.RuntimeOptions) (_result *ListEngineNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEngineNamespaces"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEngineNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the namespaces of a Nacos instance.
//
// @param request - ListEngineNamespacesRequest
//
// @return ListEngineNamespacesResponse
func (client *Client) ListEngineNamespaces(request *ListEngineNamespacesRequest) (_result *ListEngineNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEngineNamespacesResponse{}
	_body, _err := client.ListEngineNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Eureka instances.
//
// @param request - ListEurekaInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEurekaInstancesResponse
func (client *Client) ListEurekaInstancesWithOptions(request *ListEurekaInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListEurekaInstancesResponse, _err error) {
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
		Action:      dara.String("ListEurekaInstances"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEurekaInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Eureka instances.
//
// @param request - ListEurekaInstancesRequest
//
// @return ListEurekaInstancesResponse
func (client *Client) ListEurekaInstances(request *ListEurekaInstancesRequest) (_result *ListEurekaInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEurekaInstancesResponse{}
	_body, _err := client.ListEurekaInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Eureka services.
//
// @param request - ListEurekaServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEurekaServicesResponse
func (client *Client) ListEurekaServicesWithOptions(request *ListEurekaServicesRequest, runtime *dara.RuntimeOptions) (_result *ListEurekaServicesResponse, _err error) {
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
		Action:      dara.String("ListEurekaServices"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEurekaServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Eureka services.
//
// @param request - ListEurekaServicesRequest
//
// @return ListEurekaServicesResponse
func (client *Client) ListEurekaServices(request *ListEurekaServicesRequest) (_result *ListEurekaServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEurekaServicesResponse{}
	_body, _err := client.ListEurekaServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists historical data export tasks of a Microservices Engine (MSE) Zookeeper instance.
//
// @param request - ListExportZookeeperDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExportZookeeperDataResponse
func (client *Client) ListExportZookeeperDataWithOptions(request *ListExportZookeeperDataRequest, runtime *dara.RuntimeOptions) (_result *ListExportZookeeperDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExportZookeeperData"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExportZookeeperDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists historical data export tasks of a Microservices Engine (MSE) Zookeeper instance.
//
// @param request - ListExportZookeeperDataRequest
//
// @return ListExportZookeeperDataResponse
func (client *Client) ListExportZookeeperData(request *ListExportZookeeperDataRequest) (_result *ListExportZookeeperDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExportZookeeperDataResponse{}
	_body, _err := client.ListExportZookeeperDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of throttling rules.
//
// @param request - ListFlowRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowRulesResponse
func (client *Client) ListFlowRulesWithOptions(request *ListFlowRulesRequest, runtime *dara.RuntimeOptions) (_result *ListFlowRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceSearchKey) {
		query["ResourceSearchKey"] = request.ResourceSearchKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlowRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of throttling rules.
//
// @param request - ListFlowRulesRequest
//
// @return ListFlowRulesResponse
func (client *Client) ListFlowRules(request *ListFlowRulesRequest) (_result *ListFlowRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlowRulesResponse{}
	_body, _err := client.ListFlowRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of gateways.
//
// @param tmpReq - ListGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayResponse
func (client *Client) ListGatewayWithOptions(tmpReq *ListGatewayRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListGatewayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParams) {
		request.FilterParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParams, dara.String("FilterParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DescSort) {
		query["DescSort"] = request.DescSort
	}

	if !dara.IsNil(request.FilterParamsShrink) {
		query["FilterParams"] = request.FilterParamsShrink
	}

	if !dara.IsNil(request.OrderItem) {
		query["OrderItem"] = request.OrderItem
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGateway"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of gateways.
//
// @param request - ListGatewayRequest
//
// @return ListGatewayResponse
func (client *Client) ListGateway(request *ListGatewayRequest) (_result *ListGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayResponse{}
	_body, _err := client.ListGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询网关认证
//
// @param tmpReq - ListGatewayAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayAuthResponse
func (client *Client) ListGatewayAuthWithOptions(tmpReq *ListGatewayAuthRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListGatewayAuthShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParams) {
		request.FilterParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParams, dara.String("FilterParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DescSort) {
		query["DescSort"] = request.DescSort
	}

	if !dara.IsNil(request.FilterParamsShrink) {
		query["FilterParams"] = request.FilterParamsShrink
	}

	if !dara.IsNil(request.OrderItem) {
		query["OrderItem"] = request.OrderItem
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayAuth"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询网关认证
//
// @param request - ListGatewayAuthRequest
//
// @return ListGatewayAuthResponse
func (client *Client) ListGatewayAuth(request *ListGatewayAuthRequest) (_result *ListGatewayAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayAuthResponse{}
	_body, _err := client.ListGatewayAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of consumers on which a gateway performs authentication operations.
//
// @param request - ListGatewayAuthConsumerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayAuthConsumerResponse
func (client *Client) ListGatewayAuthConsumerWithOptions(request *ListGatewayAuthConsumerRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayAuthConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConsumerStatus) {
		query["ConsumerStatus"] = request.ConsumerStatus
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayAuthConsumer"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayAuthConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of consumers on which a gateway performs authentication operations.
//
// @param request - ListGatewayAuthConsumerRequest
//
// @return ListGatewayAuthConsumerResponse
func (client *Client) ListGatewayAuthConsumer(request *ListGatewayAuthConsumerRequest) (_result *ListGatewayAuthConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayAuthConsumerResponse{}
	_body, _err := client.ListGatewayAuthConsumerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of authorized resources for the consumer on which a gateway performs authentication operations.
//
// @param request - ListGatewayAuthConsumerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayAuthConsumerResourceResponse
func (client *Client) ListGatewayAuthConsumerResourceWithOptions(request *ListGatewayAuthConsumerResourceRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayAuthConsumerResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceStatus) {
		query["ResourceStatus"] = request.ResourceStatus
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayAuthConsumerResource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayAuthConsumerResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of authorized resources for the consumer on which a gateway performs authentication operations.
//
// @param request - ListGatewayAuthConsumerResourceRequest
//
// @return ListGatewayAuthConsumerResourceResponse
func (client *Client) ListGatewayAuthConsumerResource(request *ListGatewayAuthConsumerResourceRequest) (_result *ListGatewayAuthConsumerResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayAuthConsumerResourceResponse{}
	_body, _err := client.ListGatewayAuthConsumerResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看网关路由熔断规则
//
// @param request - ListGatewayCircuitBreakerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayCircuitBreakerRuleResponse
func (client *Client) ListGatewayCircuitBreakerRuleWithOptions(request *ListGatewayCircuitBreakerRuleRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayCircuitBreakerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.FilterParams) {
		query["FilterParams"] = request.FilterParams
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayCircuitBreakerRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayCircuitBreakerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看网关路由熔断规则
//
// @param request - ListGatewayCircuitBreakerRuleRequest
//
// @return ListGatewayCircuitBreakerRuleResponse
func (client *Client) ListGatewayCircuitBreakerRule(request *ListGatewayCircuitBreakerRuleRequest) (_result *ListGatewayCircuitBreakerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayCircuitBreakerRuleResponse{}
	_body, _err := client.ListGatewayCircuitBreakerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain names that are associated with a gateway.
//
// @param request - ListGatewayDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayDomainResponse
func (client *Client) ListGatewayDomainWithOptions(request *ListGatewayDomainRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayDomain"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names that are associated with a gateway.
//
// @param request - ListGatewayDomainRequest
//
// @return ListGatewayDomainResponse
func (client *Client) ListGatewayDomain(request *ListGatewayDomainRequest) (_result *ListGatewayDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayDomainResponse{}
	_body, _err := client.ListGatewayDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看网关路由流控规则
//
// @param request - ListGatewayFlowRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayFlowRuleResponse
func (client *Client) ListGatewayFlowRuleWithOptions(request *ListGatewayFlowRuleRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayFlowRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.FilterParams) {
		query["FilterParams"] = request.FilterParams
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayFlowRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看网关路由流控规则
//
// @param request - ListGatewayFlowRuleRequest
//
// @return ListGatewayFlowRuleResponse
func (client *Client) ListGatewayFlowRule(request *ListGatewayFlowRuleRequest) (_result *ListGatewayFlowRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayFlowRuleResponse{}
	_body, _err := client.ListGatewayFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看网关路由隔离规则
//
// @param request - ListGatewayIsolationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayIsolationRuleResponse
func (client *Client) ListGatewayIsolationRuleWithOptions(request *ListGatewayIsolationRuleRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayIsolationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.FilterParams) {
		query["FilterParams"] = request.FilterParams
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayIsolationRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看网关路由隔离规则
//
// @param request - ListGatewayIsolationRuleRequest
//
// @return ListGatewayIsolationRuleResponse
func (client *Client) ListGatewayIsolationRule(request *ListGatewayIsolationRuleRequest) (_result *ListGatewayIsolationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayIsolationRuleResponse{}
	_body, _err := client.ListGatewayIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the routes of a gateway.
//
// @param tmpReq - ListGatewayRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayRouteResponse
func (client *Client) ListGatewayRouteWithOptions(tmpReq *ListGatewayRouteRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListGatewayRouteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParams) {
		request.FilterParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParams, dara.String("FilterParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DescSort) {
		query["DescSort"] = request.DescSort
	}

	if !dara.IsNil(request.FilterParamsShrink) {
		query["FilterParams"] = request.FilterParamsShrink
	}

	if !dara.IsNil(request.OrderItem) {
		query["OrderItem"] = request.OrderItem
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayRoute"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routes of a gateway.
//
// @param request - ListGatewayRouteRequest
//
// @return ListGatewayRouteResponse
func (client *Client) ListGatewayRoute(request *ListGatewayRouteRequest) (_result *ListGatewayRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayRouteResponse{}
	_body, _err := client.ListGatewayRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of routes for which authentication is enabled.
//
// @param request - ListGatewayRouteOnAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayRouteOnAuthResponse
func (client *Client) ListGatewayRouteOnAuthWithOptions(request *ListGatewayRouteOnAuthRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayRouteOnAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayRouteOnAuth"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayRouteOnAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of routes for which authentication is enabled.
//
// @param request - ListGatewayRouteOnAuthRequest
//
// @return ListGatewayRouteOnAuthResponse
func (client *Client) ListGatewayRouteOnAuth(request *ListGatewayRouteOnAuthRequest) (_result *ListGatewayRouteOnAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayRouteOnAuthResponse{}
	_body, _err := client.ListGatewayRouteOnAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of services that are subscribed with a gateway.
//
// @param tmpReq - ListGatewayServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayServiceResponse
func (client *Client) ListGatewayServiceWithOptions(tmpReq *ListGatewayServiceRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListGatewayServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterParams) {
		request.FilterParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterParams, dara.String("FilterParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DescSort) {
		query["DescSort"] = request.DescSort
	}

	if !dara.IsNil(request.FilterParamsShrink) {
		query["FilterParams"] = request.FilterParamsShrink
	}

	if !dara.IsNil(request.OrderItem) {
		query["OrderItem"] = request.OrderItem
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayService"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of services that are subscribed with a gateway.
//
// @param request - ListGatewayServiceRequest
//
// @return ListGatewayServiceResponse
func (client *Client) ListGatewayService(request *ListGatewayServiceRequest) (_result *ListGatewayServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayServiceResponse{}
	_body, _err := client.ListGatewayServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Server Load Balancer (SLB) instances that are associated with a gateway.
//
// @param request - ListGatewaySlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewaySlbResponse
func (client *Client) ListGatewaySlbWithOptions(request *ListGatewaySlbRequest, runtime *dara.RuntimeOptions) (_result *ListGatewaySlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewaySlb"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewaySlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Server Load Balancer (SLB) instances that are associated with a gateway.
//
// @param request - ListGatewaySlbRequest
//
// @return ListGatewaySlbResponse
func (client *Client) ListGatewaySlb(request *ListGatewaySlbRequest) (_result *ListGatewaySlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewaySlbResponse{}
	_body, _err := client.ListGatewaySlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of zones where a gateway is available.
//
// @param request - ListGatewayZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayZoneResponse
func (client *Client) ListGatewayZoneWithOptions(request *ListGatewayZoneRequest, runtime *dara.RuntimeOptions) (_result *ListGatewayZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayZone"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of zones where a gateway is available.
//
// @param request - ListGatewayZoneRequest
//
// @return ListGatewayZoneResponse
func (client *Client) ListGatewayZone(request *ListGatewayZoneRequest) (_result *ListGatewayZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGatewayZoneResponse{}
	_body, _err := client.ListGatewayZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Displays the number of nodes that can be deployed for an instance.
//
// @param request - ListInstanceCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceCountResponse
func (client *Client) ListInstanceCountWithOptions(request *ListInstanceCountRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.MseVersion) {
		query["MseVersion"] = request.MseVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceCount"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays the number of nodes that can be deployed for an instance.
//
// @param request - ListInstanceCountRequest
//
// @return ListInstanceCountResponse
func (client *Client) ListInstanceCount(request *ListInstanceCountRequest) (_result *ListInstanceCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceCountResponse{}
	_body, _err := client.ListInstanceCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询隔离规则
//
// @param request - ListIsolationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIsolationRulesResponse
func (client *Client) ListIsolationRulesWithOptions(request *ListIsolationRulesRequest, runtime *dara.RuntimeOptions) (_result *ListIsolationRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceSearchKey) {
		query["ResourceSearchKey"] = request.ResourceSearchKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIsolationRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIsolationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询隔离规则
//
// @param request - ListIsolationRulesRequest
//
// @return ListIsolationRulesResponse
func (client *Client) ListIsolationRules(request *ListIsolationRulesRequest) (_result *ListIsolationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIsolationRulesResponse{}
	_body, _err := client.ListIsolationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries listeners based on configuration information.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param tmpReq - ListListenersByConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListListenersByConfigResponse
func (client *Client) ListListenersByConfigWithOptions(tmpReq *ListListenersByConfigRequest, runtime *dara.RuntimeOptions) (_result *ListListenersByConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListListenersByConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtGrayRules) {
		request.ExtGrayRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtGrayRules, dara.String("ExtGrayRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.ExtGrayRulesShrink) {
		query["ExtGrayRules"] = request.ExtGrayRulesShrink
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListListenersByConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListListenersByConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries listeners based on configuration information.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListListenersByConfigRequest
//
// @return ListListenersByConfigResponse
func (client *Client) ListListenersByConfig(request *ListListenersByConfigRequest) (_result *ListListenersByConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListListenersByConfigResponse{}
	_body, _err := client.ListListenersByConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about listeners based on IP addresses.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListListenersByIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListListenersByIpResponse
func (client *Client) ListListenersByIpWithOptions(request *ListListenersByIpRequest, runtime *dara.RuntimeOptions) (_result *ListListenersByIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListListenersByIp"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListListenersByIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about listeners based on IP addresses.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListListenersByIpRequest
//
// @return ListListenersByIpResponse
func (client *Client) ListListenersByIp(request *ListListenersByIpRequest) (_result *ListListenersByIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListListenersByIpResponse{}
	_body, _err := client.ListListenersByIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a migration task.
//
// @param request - ListMigrationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMigrationTaskResponse
func (client *Client) ListMigrationTaskWithOptions(request *ListMigrationTaskRequest, runtime *dara.RuntimeOptions) (_result *ListMigrationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.OriginInstanceName) {
		query["OriginInstanceName"] = request.OriginInstanceName
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMigrationTask"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMigrationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a migration task.
//
// @param request - ListMigrationTaskRequest
//
// @return ListMigrationTaskResponse
func (client *Client) ListMigrationTask(request *ListMigrationTaskRequest) (_result *ListMigrationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMigrationTaskResponse{}
	_body, _err := client.ListMigrationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Nacos configurations.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListNacosConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNacosConfigsResponse
func (client *Client) ListNacosConfigsWithOptions(request *ListNacosConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListNacosConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNacosConfigs"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNacosConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Nacos configurations.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListNacosConfigsRequest
//
// @return ListNacosConfigsResponse
func (client *Client) ListNacosConfigs(request *ListNacosConfigsRequest) (_result *ListNacosConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNacosConfigsResponse{}
	_body, _err := client.ListNacosConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration history of a Microservices Engine (MSE) Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListNacosHistoryConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNacosHistoryConfigsResponse
func (client *Client) ListNacosHistoryConfigsWithOptions(request *ListNacosHistoryConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListNacosHistoryConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNacosHistoryConfigs"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNacosHistoryConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration history of a Microservices Engine (MSE) Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - ListNacosHistoryConfigsRequest
//
// @return ListNacosHistoryConfigsResponse
func (client *Client) ListNacosHistoryConfigs(request *ListNacosHistoryConfigsRequest) (_result *ListNacosHistoryConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNacosHistoryConfigsResponse{}
	_body, _err := client.ListNacosHistoryConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取McpServer列表
//
// @param request - ListNacosMcpServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNacosMcpServersResponse
func (client *Client) ListNacosMcpServersWithOptions(request *ListNacosMcpServersRequest, runtime *dara.RuntimeOptions) (_result *ListNacosMcpServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNacosMcpServers"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNacosMcpServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取McpServer列表
//
// @param request - ListNacosMcpServersRequest
//
// @return ListNacosMcpServersResponse
func (client *Client) ListNacosMcpServers(request *ListNacosMcpServersRequest) (_result *ListNacosMcpServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNacosMcpServersResponse{}
	_body, _err := client.ListNacosMcpServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示命名空间列表
//
// @param tmpReq - ListNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespacesResponse
func (client *Client) ListNamespacesWithOptions(tmpReq *ListNamespacesRequest, runtime *dara.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNamespacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNamespaces"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示命名空间列表
//
// @param request - ListNamespacesRequest
//
// @return ListNamespacesResponse
func (client *Client) ListNamespaces(request *ListNamespacesRequest) (_result *ListNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNamespacesResponse{}
	_body, _err := client.ListNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the trajectory data of a Nacos registry.
//
// @param request - ListNamingTrackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamingTrackResponse
func (client *Client) ListNamingTrackWithOptions(request *ListNamingTrackRequest, runtime *dara.RuntimeOptions) (_result *ListNamingTrackResponse, _err error) {
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
		Action:      dara.String("ListNamingTrack"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNamingTrackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the trajectory data of a Nacos registry.
//
// @param request - ListNamingTrackRequest
//
// @return ListNamingTrackResponse
func (client *Client) ListNamingTrack(request *ListNamingTrackRequest) (_result *ListNamingTrackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNamingTrackResponse{}
	_body, _err := client.ListNamingTrackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of gateway certificates.
//
// @param request - ListSSLCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSSLCertResponse
func (client *Client) ListSSLCertWithOptions(request *ListSSLCertRequest, runtime *dara.RuntimeOptions) (_result *ListSSLCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSSLCert"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSSLCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of gateway certificates.
//
// @param request - ListSSLCertRequest
//
// @return ListSSLCertResponse
func (client *Client) ListSSLCert(request *ListSSLCertRequest) (_result *ListSSLCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSSLCertResponse{}
	_body, _err := client.ListSSLCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about security groups.
//
// @param request - ListSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityGroupResponse
func (client *Client) ListSecurityGroupWithOptions(request *ListSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecurityGroup"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about security groups.
//
// @param request - ListSecurityGroupRequest
//
// @return ListSecurityGroupResponse
func (client *Client) ListSecurityGroup(request *ListSecurityGroupRequest) (_result *ListSecurityGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecurityGroupResponse{}
	_body, _err := client.ListSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security group rules of a gateway.
//
// @param request - ListSecurityGroupRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityGroupRuleResponse
func (client *Client) ListSecurityGroupRuleWithOptions(request *ListSecurityGroupRuleRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityGroupRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecurityGroupRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityGroupRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security group rules of a gateway.
//
// @param request - ListSecurityGroupRuleRequest
//
// @return ListSecurityGroupRuleResponse
func (client *Client) ListSecurityGroupRule(request *ListSecurityGroupRuleRequest) (_result *ListSecurityGroupRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSecurityGroupRuleResponse{}
	_body, _err := client.ListSecurityGroupRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the custom behavior of traffic protection.
//
// @param tmpReq - ListSentinelBlockFallbackDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSentinelBlockFallbackDefinitionsResponse
func (client *Client) ListSentinelBlockFallbackDefinitionsWithOptions(tmpReq *ListSentinelBlockFallbackDefinitionsRequest, runtime *dara.RuntimeOptions) (_result *ListSentinelBlockFallbackDefinitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSentinelBlockFallbackDefinitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClassificationSet) {
		request.ClassificationSetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClassificationSet, dara.String("ClassificationSet"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClassificationSetShrink) {
		query["ClassificationSet"] = request.ClassificationSetShrink
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSentinelBlockFallbackDefinitions"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSentinelBlockFallbackDefinitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the custom behavior of traffic protection.
//
// @param request - ListSentinelBlockFallbackDefinitionsRequest
//
// @return ListSentinelBlockFallbackDefinitionsResponse
func (client *Client) ListSentinelBlockFallbackDefinitions(request *ListSentinelBlockFallbackDefinitionsRequest) (_result *ListSentinelBlockFallbackDefinitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSentinelBlockFallbackDefinitionsResponse{}
	_body, _err := client.ListSentinelBlockFallbackDefinitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of associated sources.
//
// @param request - ListServiceSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceSourceResponse
func (client *Client) ListServiceSourceWithOptions(request *ListServiceSourceRequest, runtime *dara.RuntimeOptions) (_result *ListServiceSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceSource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of associated sources.
//
// @param request - ListServiceSourceRequest
//
// @return ListServiceSourceResponse
func (client *Client) ListServiceSource(request *ListServiceSourceRequest) (_result *ListServiceSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceSourceResponse{}
	_body, _err := client.ListServiceSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tagged resources.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tagged resources.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询热点参数防护规则（HTTP 请求）
//
// @param request - ListWebFlowRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWebFlowRulesResponse
func (client *Client) ListWebFlowRulesWithOptions(request *ListWebFlowRulesRequest, runtime *dara.RuntimeOptions) (_result *ListWebFlowRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceSearchKey) {
		query["ResourceSearchKey"] = request.ResourceSearchKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWebFlowRules"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWebFlowRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询热点参数防护规则（HTTP 请求）
//
// @param request - ListWebFlowRulesRequest
//
// @return ListWebFlowRulesResponse
func (client *Client) ListWebFlowRules(request *ListWebFlowRulesRequest) (_result *ListWebFlowRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWebFlowRulesResponse{}
	_body, _err := client.ListWebFlowRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the track data of a ZooKeeper instance.
//
// @param request - ListZkTrackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListZkTrackResponse
func (client *Client) ListZkTrackWithOptions(request *ListZkTrackRequest, runtime *dara.RuntimeOptions) (_result *ListZkTrackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.Reverse) {
		query["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListZkTrack"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListZkTrackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the track data of a ZooKeeper instance.
//
// @param request - ListZkTrackRequest
//
// @return ListZkTrackResponse
func (client *Client) ListZkTrack(request *ListZkTrackRequest) (_result *ListZkTrackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListZkTrackResponse{}
	_body, _err := client.ListZkTrackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the child nodes of a ZooKeeper node.
//
// @param request - ListZnodeChildrenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListZnodeChildrenResponse
func (client *Client) ListZnodeChildrenWithOptions(request *ListZnodeChildrenRequest, runtime *dara.RuntimeOptions) (_result *ListZnodeChildrenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListZnodeChildren"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListZnodeChildrenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the child nodes of a ZooKeeper node.
//
// @param request - ListZnodeChildrenRequest
//
// @return ListZnodeChildrenResponse
func (client *Client) ListZnodeChildren(request *ListZnodeChildrenRequest) (_result *ListZnodeChildrenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListZnodeChildrenResponse{}
	_body, _err := client.ListZnodeChildrenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a cluster for which Microservice Governance is enabled.
//
// @param tmpReq - ModifyGovernanceKubernetesClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGovernanceKubernetesClusterResponse
func (client *Client) ModifyGovernanceKubernetesClusterWithOptions(tmpReq *ModifyGovernanceKubernetesClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyGovernanceKubernetesClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyGovernanceKubernetesClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NamespaceInfos) {
		request.NamespaceInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NamespaceInfos, dara.String("NamespaceInfos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceInfosShrink) {
		body["NamespaceInfos"] = request.NamespaceInfosShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGovernanceKubernetesCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGovernanceKubernetesClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a cluster for which Microservice Governance is enabled.
//
// @param request - ModifyGovernanceKubernetesClusterRequest
//
// @return ModifyGovernanceKubernetesClusterResponse
func (client *Client) ModifyGovernanceKubernetesCluster(request *ModifyGovernanceKubernetesClusterRequest) (_result *ModifyGovernanceKubernetesClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGovernanceKubernetesClusterResponse{}
	_body, _err := client.ModifyGovernanceKubernetesClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies configurations of the lossless online and offline feature.
//
// @param request - ModifyLosslessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLosslessRuleResponse
func (client *Client) ModifyLosslessRuleWithOptions(request *ModifyLosslessRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyLosslessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Aligned) {
		query["Aligned"] = request.Aligned
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DelayTime) {
		query["DelayTime"] = request.DelayTime
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.FuncType) {
		query["FuncType"] = request.FuncType
	}

	if !dara.IsNil(request.LossLessDetail) {
		query["LossLessDetail"] = request.LossLessDetail
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Notice) {
		query["Notice"] = request.Notice
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Related) {
		query["Related"] = request.Related
	}

	if !dara.IsNil(request.WarmupTime) {
		query["WarmupTime"] = request.WarmupTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLosslessRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLosslessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies configurations of the lossless online and offline feature.
//
// @param request - ModifyLosslessRuleRequest
//
// @return ModifyLosslessRuleResponse
func (client *Client) ModifyLosslessRule(request *ModifyLosslessRuleRequest) (_result *ModifyLosslessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLosslessRuleResponse{}
	_body, _err := client.ModifyLosslessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unpublishes a route for a gateway.
//
// @param request - OfflineGatewayRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineGatewayRouteResponse
func (client *Client) OfflineGatewayRouteWithOptions(request *OfflineGatewayRouteRequest, runtime *dara.RuntimeOptions) (_result *OfflineGatewayRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineGatewayRoute"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unpublishes a route for a gateway.
//
// @param request - OfflineGatewayRouteRequest
//
// @return OfflineGatewayRouteResponse
func (client *Client) OfflineGatewayRoute(request *OfflineGatewayRouteRequest) (_result *OfflineGatewayRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineGatewayRouteResponse{}
	_body, _err := client.OfflineGatewayRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Subscribes to the notification feature if a risk is detected during a health check.
//
// @param request - OrderClusterHealthCheckRiskNoticeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OrderClusterHealthCheckRiskNoticeResponse
func (client *Client) OrderClusterHealthCheckRiskNoticeWithOptions(request *OrderClusterHealthCheckRiskNoticeRequest, runtime *dara.RuntimeOptions) (_result *OrderClusterHealthCheckRiskNoticeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mute) {
		query["Mute"] = request.Mute
	}

	if !dara.IsNil(request.NoticeType) {
		query["NoticeType"] = request.NoticeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.RiskCode) {
		query["RiskCode"] = request.RiskCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OrderClusterHealthCheckRiskNotice"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OrderClusterHealthCheckRiskNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Subscribes to the notification feature if a risk is detected during a health check.
//
// @param request - OrderClusterHealthCheckRiskNoticeRequest
//
// @return OrderClusterHealthCheckRiskNoticeResponse
func (client *Client) OrderClusterHealthCheckRiskNotice(request *OrderClusterHealthCheckRiskNoticeRequest) (_result *OrderClusterHealthCheckRiskNoticeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OrderClusterHealthCheckRiskNoticeResponse{}
	_body, _err := client.OrderClusterHealthCheckRiskNoticeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies whether to convert all letters of a header into lowercase letters. For requests and responses, HTTP/1.1 headers are not case-sensitive. By default, all letters of headers are converted into lowercase letters.
//
// @param request - PreserveHeaderFormatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreserveHeaderFormatResponse
func (client *Client) PreserveHeaderFormatWithOptions(request *PreserveHeaderFormatRequest, runtime *dara.RuntimeOptions) (_result *PreserveHeaderFormatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.PreserveHeaderFormat) {
		query["PreserveHeaderFormat"] = request.PreserveHeaderFormat
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreserveHeaderFormat"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreserveHeaderFormatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies whether to convert all letters of a header into lowercase letters. For requests and responses, HTTP/1.1 headers are not case-sensitive. By default, all letters of headers are converted into lowercase letters.
//
// @param request - PreserveHeaderFormatRequest
//
// @return PreserveHeaderFormatResponse
func (client *Client) PreserveHeaderFormat(request *PreserveHeaderFormatRequest) (_result *PreserveHeaderFormatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreserveHeaderFormatResponse{}
	_body, _err := client.PreserveHeaderFormatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all the microservices of a service source.
//
// @param request - PullServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PullServicesResponse
func (client *Client) PullServicesWithOptions(request *PullServicesRequest, runtime *dara.RuntimeOptions) (_result *PullServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PullServices"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PullServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all the microservices of a service source.
//
// @param request - PullServicesRequest
//
// @return PullServicesResponse
func (client *Client) PullServices(request *PullServicesRequest) (_result *PullServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PullServicesResponse{}
	_body, _err := client.PullServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a task to check risk evaluation for an instance.
//
// @param request - PutClusterHealthCheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutClusterHealthCheckTaskResponse
func (client *Client) PutClusterHealthCheckTaskWithOptions(request *PutClusterHealthCheckTaskRequest, runtime *dara.RuntimeOptions) (_result *PutClusterHealthCheckTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutClusterHealthCheckTask"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutClusterHealthCheckTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a task to check risk evaluation for an instance.
//
// @param request - PutClusterHealthCheckTaskRequest
//
// @return PutClusterHealthCheckTaskResponse
func (client *Client) PutClusterHealthCheckTask(request *PutClusterHealthCheckTaskRequest) (_result *PutClusterHealthCheckTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutClusterHealthCheckTaskResponse{}
	_body, _err := client.PutClusterHealthCheckTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all the lanes in a lane group.
//
// @param request - QueryAllSwimmingLaneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllSwimmingLaneResponse
func (client *Client) QueryAllSwimmingLaneWithOptions(request *QueryAllSwimmingLaneRequest, runtime *dara.RuntimeOptions) (_result *QueryAllSwimmingLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAllSwimmingLane"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAllSwimmingLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all the lanes in a lane group.
//
// @param request - QueryAllSwimmingLaneRequest
//
// @return QueryAllSwimmingLaneResponse
func (client *Client) QueryAllSwimmingLane(request *QueryAllSwimmingLaneRequest) (_result *QueryAllSwimmingLaneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAllSwimmingLaneResponse{}
	_body, _err := client.QueryAllSwimmingLaneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all lane groups.
//
// @param request - QueryAllSwimmingLaneGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllSwimmingLaneGroupResponse
func (client *Client) QueryAllSwimmingLaneGroupWithOptions(request *QueryAllSwimmingLaneGroupRequest, runtime *dara.RuntimeOptions) (_result *QueryAllSwimmingLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAllSwimmingLaneGroup"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAllSwimmingLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all lane groups.
//
// @param request - QueryAllSwimmingLaneGroupRequest
//
// @return QueryAllSwimmingLaneGroupResponse
func (client *Client) QueryAllSwimmingLaneGroup(request *QueryAllSwimmingLaneGroupRequest) (_result *QueryAllSwimmingLaneGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAllSwimmingLaneGroupResponse{}
	_body, _err := client.QueryAllSwimmingLaneGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about regions.
//
// @param request - QueryBusinessLocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBusinessLocationsResponse
func (client *Client) QueryBusinessLocationsWithOptions(request *QueryBusinessLocationsRequest, runtime *dara.RuntimeOptions) (_result *QueryBusinessLocationsResponse, _err error) {
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
		Action:      dara.String("QueryBusinessLocations"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBusinessLocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about regions.
//
// @param request - QueryBusinessLocationsRequest
//
// @return QueryBusinessLocationsResponse
func (client *Client) QueryBusinessLocations(request *QueryBusinessLocationsRequest) (_result *QueryBusinessLocationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBusinessLocationsResponse{}
	_body, _err := client.QueryBusinessLocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an instance.
//
// @param request - QueryClusterDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryClusterDetailResponse
func (client *Client) QueryClusterDetailWithOptions(request *QueryClusterDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryClusterDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AclSwitch) {
		query["AclSwitch"] = request.AclSwitch
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryClusterDetail"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryClusterDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an instance.
//
// @param request - QueryClusterDetailRequest
//
// @return QueryClusterDetailResponse
func (client *Client) QueryClusterDetail(request *QueryClusterDetailRequest) (_result *QueryClusterDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryClusterDetailResponse{}
	_body, _err := client.QueryClusterDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries disk specifications that are supported by an instance.
//
// @param request - QueryClusterDiskSpecificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryClusterDiskSpecificationResponse
func (client *Client) QueryClusterDiskSpecificationWithOptions(request *QueryClusterDiskSpecificationRequest, runtime *dara.RuntimeOptions) (_result *QueryClusterDiskSpecificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryClusterDiskSpecification"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryClusterDiskSpecificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries disk specifications that are supported by an instance.
//
// @param request - QueryClusterDiskSpecificationRequest
//
// @return QueryClusterDiskSpecificationResponse
func (client *Client) QueryClusterDiskSpecification(request *QueryClusterDiskSpecificationRequest) (_result *QueryClusterDiskSpecificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryClusterDiskSpecificationResponse{}
	_body, _err := client.QueryClusterDiskSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the static information of an instance.
//
// @param request - QueryClusterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryClusterInfoResponse
func (client *Client) QueryClusterInfoWithOptions(request *QueryClusterInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryClusterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AclSwitch) {
		query["AclSwitch"] = request.AclSwitch
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryClusterInfo"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryClusterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the static information of an instance.
//
// @param request - QueryClusterInfoRequest
//
// @return QueryClusterInfoResponse
func (client *Client) QueryClusterInfo(request *QueryClusterInfoRequest) (_result *QueryClusterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryClusterInfoResponse{}
	_body, _err := client.QueryClusterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of supported instance specifications.
//
// @param request - QueryClusterSpecificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryClusterSpecificationResponse
func (client *Client) QueryClusterSpecificationWithOptions(request *QueryClusterSpecificationRequest, runtime *dara.RuntimeOptions) (_result *QueryClusterSpecificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConnectType) {
		query["ConnectType"] = request.ConnectType
	}

	if !dara.IsNil(request.MseVersion) {
		query["MseVersion"] = request.MseVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryClusterSpecification"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryClusterSpecificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of supported instance specifications.
//
// @param request - QueryClusterSpecificationRequest
//
// @return QueryClusterSpecificationResponse
func (client *Client) QueryClusterSpecification(request *QueryClusterSpecificationRequest) (_result *QueryClusterSpecificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryClusterSpecificationResponse{}
	_body, _err := client.QueryClusterSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries configuration information of an instance.
//
// @param request - QueryConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConfigResponse
func (client *Client) QueryConfigWithOptions(request *QueryConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NeedRunningConf) {
		query["NeedRunningConf"] = request.NeedRunningConf
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries configuration information of an instance.
//
// @param request - QueryConfigRequest
//
// @return QueryConfigResponse
func (client *Client) QueryConfig(request *QueryConfigRequest) (_result *QueryConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryConfigResponse{}
	_body, _err := client.QueryConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions supported by a gateway.
//
// @param request - QueryGatewayRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGatewayRegionResponse
func (client *Client) QueryGatewayRegionWithOptions(request *QueryGatewayRegionRequest, runtime *dara.RuntimeOptions) (_result *QueryGatewayRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGatewayRegion"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGatewayRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions supported by a gateway.
//
// @param request - QueryGatewayRegionRequest
//
// @return QueryGatewayRegionResponse
func (client *Client) QueryGatewayRegion(request *QueryGatewayRegionRequest) (_result *QueryGatewayRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryGatewayRegionResponse{}
	_body, _err := client.QueryGatewayRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available gateway types.
//
// @param request - QueryGatewayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGatewayTypeResponse
func (client *Client) QueryGatewayTypeWithOptions(request *QueryGatewayTypeRequest, runtime *dara.RuntimeOptions) (_result *QueryGatewayTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGatewayType"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGatewayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available gateway types.
//
// @param request - QueryGatewayTypeRequest
//
// @return QueryGatewayTypeResponse
func (client *Client) QueryGatewayType(request *QueryGatewayTypeRequest) (_result *QueryGatewayTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryGatewayTypeResponse{}
	_body, _err := client.QueryGatewayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Kubernetes clusters for which Microservices Governance is activated.
//
// @param request - QueryGovernanceKubernetesClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGovernanceKubernetesClusterResponse
func (client *Client) QueryGovernanceKubernetesClusterWithOptions(request *QueryGovernanceKubernetesClusterRequest, runtime *dara.RuntimeOptions) (_result *QueryGovernanceKubernetesClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGovernanceKubernetesCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGovernanceKubernetesClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Kubernetes clusters for which Microservices Governance is activated.
//
// @param request - QueryGovernanceKubernetesClusterRequest
//
// @return QueryGovernanceKubernetesClusterResponse
func (client *Client) QueryGovernanceKubernetesCluster(request *QueryGovernanceKubernetesClusterRequest) (_result *QueryGovernanceKubernetesClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryGovernanceKubernetesClusterResponse{}
	_body, _err := client.QueryGovernanceKubernetesClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the runtime data of a specified cluster.
//
// @param request - QueryInstancesInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstancesInfoResponse
func (client *Client) QueryInstancesInfoWithOptions(request *QueryInstancesInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryInstancesInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInstancesInfo"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstancesInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the runtime data of a specified cluster.
//
// @param request - QueryInstancesInfoRequest
//
// @return QueryInstancesInfoResponse
func (client *Client) QueryInstancesInfo(request *QueryInstancesInfoRequest) (_result *QueryInstancesInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInstancesInfoResponse{}
	_body, _err := client.QueryInstancesInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries monitoring information.
//
// @param request - QueryMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMonitorResponse
func (client *Client) QueryMonitorWithOptions(request *QueryMonitorRequest, runtime *dara.RuntimeOptions) (_result *QueryMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MonitorType) {
		query["MonitorType"] = request.MonitorType
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Step) {
		query["Step"] = request.Step
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMonitor"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries monitoring information.
//
// @param request - QueryMonitorRequest
//
// @return QueryMonitorResponse
func (client *Client) QueryMonitor(request *QueryMonitorRequest) (_result *QueryMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMonitorResponse{}
	_body, _err := client.QueryMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询nacos灰度配置
//
// @param request - QueryNacosGrayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryNacosGrayConfigResponse
func (client *Client) QueryNacosGrayConfigWithOptions(request *QueryNacosGrayConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryNacosGrayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.GrayName) {
		query["GrayName"] = request.GrayName
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryNacosGrayConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryNacosGrayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询nacos灰度配置
//
// @param request - QueryNacosGrayConfigRequest
//
// @return QueryNacosGrayConfigResponse
func (client *Client) QueryNacosGrayConfig(request *QueryNacosGrayConfigRequest) (_result *QueryNacosGrayConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryNacosGrayConfigResponse{}
	_body, _err := client.QueryNacosGrayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询MSE命名空间
//
// @param request - QueryNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryNamespaceResponse
func (client *Client) QueryNamespaceWithOptions(request *QueryNamespaceRequest, runtime *dara.RuntimeOptions) (_result *QueryNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryNamespace"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询MSE命名空间
//
// @param request - QueryNamespaceRequest
//
// @return QueryNamespaceResponse
func (client *Client) QueryNamespace(request *QueryNamespaceRequest) (_result *QueryNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryNamespaceResponse{}
	_body, _err := client.QueryNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the type of a Server Load Balancer (SLB) instance.
//
// @param request - QuerySlbSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySlbSpecResponse
func (client *Client) QuerySlbSpecWithOptions(request *QuerySlbSpecRequest, runtime *dara.RuntimeOptions) (_result *QuerySlbSpecResponse, _err error) {
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
		Action:      dara.String("QuerySlbSpec"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySlbSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the type of a Server Load Balancer (SLB) instance.
//
// @param request - QuerySlbSpecRequest
//
// @return QuerySlbSpecResponse
func (client *Client) QuerySlbSpec(request *QuerySlbSpecRequest) (_result *QuerySlbSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySlbSpecResponse{}
	_body, _err := client.QuerySlbSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a lane based on the lane ID.
//
// @param request - QuerySwimmingLaneByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySwimmingLaneByIdResponse
func (client *Client) QuerySwimmingLaneByIdWithOptions(request *QuerySwimmingLaneByIdRequest, runtime *dara.RuntimeOptions) (_result *QuerySwimmingLaneByIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.LaneId) {
		query["LaneId"] = request.LaneId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySwimmingLaneById"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySwimmingLaneByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a lane based on the lane ID.
//
// @param request - QuerySwimmingLaneByIdRequest
//
// @return QuerySwimmingLaneByIdResponse
func (client *Client) QuerySwimmingLaneById(request *QuerySwimmingLaneByIdRequest) (_result *QuerySwimmingLaneByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySwimmingLaneByIdResponse{}
	_body, _err := client.QuerySwimmingLaneByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a ZooKeeper node.
//
// @param request - QueryZnodeDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryZnodeDetailResponse
func (client *Client) QueryZnodeDetailWithOptions(request *QueryZnodeDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryZnodeDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryZnodeDetail"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryZnodeDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a ZooKeeper node.
//
// @param request - QueryZnodeDetailRequest
//
// @return QueryZnodeDetailResponse
func (client *Client) QueryZnodeDetail(request *QueryZnodeDetailRequest) (_result *QueryZnodeDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryZnodeDetailResponse{}
	_body, _err := client.QueryZnodeDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除单个应用
//
// @param request - RemoveApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApplicationResponse
func (client *Client) RemoveApplicationWithOptions(request *RemoveApplicationRequest, runtime *dara.RuntimeOptions) (_result *RemoveApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApplication"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除单个应用
//
// @param request - RemoveApplicationRequest
//
// @return RemoveApplicationResponse
func (client *Client) RemoveApplication(request *RemoveApplicationRequest) (_result *RemoveApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveApplicationResponse{}
	_body, _err := client.RemoveApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemoveAuthPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAuthPolicyResponse
func (client *Client) RemoveAuthPolicyWithOptions(request *RemoveAuthPolicyRequest, runtime *dara.RuntimeOptions) (_result *RemoveAuthPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAuthPolicy"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAuthPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveAuthPolicyRequest
//
// @return RemoveAuthPolicyResponse
func (client *Client) RemoveAuthPolicy(request *RemoveAuthPolicyRequest) (_result *RemoveAuthPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveAuthPolicyResponse{}
	_body, _err := client.RemoveAuthPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a registry.
//
// @param request - RestartClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartClusterResponse
func (client *Client) RestartClusterWithOptions(request *RestartClusterRequest, runtime *dara.RuntimeOptions) (_result *RestartClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PodNameList) {
		query["PodNameList"] = request.PodNameList
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a registry.
//
// @param request - RestartClusterRequest
//
// @return RestartClusterResponse
func (client *Client) RestartCluster(request *RestartClusterRequest) (_result *RestartClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartClusterResponse{}
	_body, _err := client.RestartClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retries a cluster.
//
// @param request - RetryClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryClusterResponse
func (client *Client) RetryClusterWithOptions(request *RetryClusterRequest, runtime *dara.RuntimeOptions) (_result *RetryClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retries a cluster.
//
// @param request - RetryClusterRequest
//
// @return RetryClusterResponse
func (client *Client) RetryCluster(request *RetryClusterRequest) (_result *RetryClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryClusterResponse{}
	_body, _err := client.RetryClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an idle Server Load Balancer (SLB) instance that is associated with a gateway.
//
// @param request - SelectGatewaySlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectGatewaySlbResponse
func (client *Client) SelectGatewaySlbWithOptions(request *SelectGatewaySlbRequest, runtime *dara.RuntimeOptions) (_result *SelectGatewaySlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectGatewaySlb"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectGatewaySlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an idle Server Load Balancer (SLB) instance that is associated with a gateway.
//
// @param request - SelectGatewaySlbRequest
//
// @return SelectGatewaySlbResponse
func (client *Client) SelectGatewaySlb(request *SelectGatewaySlbRequest) (_result *SelectGatewaySlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SelectGatewaySlbResponse{}
	_body, _err := client.SelectGatewaySlbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tags a specified resource.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tags a specified resource.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Untags resources.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
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
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Untags resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an IP address whitelist.
//
// @param request - UpdateAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAclResponse
func (client *Client) UpdateAclWithOptions(request *UpdateAclRequest, runtime *dara.RuntimeOptions) (_result *UpdateAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AclEntryList) {
		query["AclEntryList"] = request.AclEntryList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAcl"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an IP address whitelist.
//
// @param request - UpdateAclRequest
//
// @return UpdateAclResponse
func (client *Client) UpdateAcl(request *UpdateAclRequest) (_result *UpdateAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAclResponse{}
	_body, _err := client.UpdateAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a service authentication rule.
//
// @param request - UpdateAuthPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthPolicyResponse
func (client *Client) UpdateAuthPolicyWithOptions(request *UpdateAuthPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AuthRule) {
		query["AuthRule"] = request.AuthRule
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.K8sNamespace) {
		query["K8sNamespace"] = request.K8sNamespace
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthPolicy"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a service authentication rule.
//
// @param request - UpdateAuthPolicyRequest
//
// @return UpdateAuthPolicyResponse
func (client *Client) UpdateAuthPolicy(request *UpdateAuthPolicyRequest) (_result *UpdateAuthPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAuthPolicyResponse{}
	_body, _err := client.UpdateAuthPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the blacklist or whitelist of a gateway.
//
// @param request - UpdateBlackWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBlackWhiteListResponse
func (client *Client) UpdateBlackWhiteListWithOptions(request *UpdateBlackWhiteListRequest, runtime *dara.RuntimeOptions) (_result *UpdateBlackWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.IsWhite) {
		query["IsWhite"] = request.IsWhite
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.ResourceIdJsonList) {
		query["ResourceIdJsonList"] = request.ResourceIdJsonList
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBlackWhiteList"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBlackWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the blacklist or whitelist of a gateway.
//
// @param request - UpdateBlackWhiteListRequest
//
// @return UpdateBlackWhiteListResponse
func (client *Client) UpdateBlackWhiteList(request *UpdateBlackWhiteListRequest) (_result *UpdateBlackWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBlackWhiteListResponse{}
	_body, _err := client.UpdateBlackWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a circuit breaking rule.
//
// @param request - UpdateCircuitBreakerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCircuitBreakerRuleResponse
func (client *Client) UpdateCircuitBreakerRuleWithOptions(request *UpdateCircuitBreakerRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateCircuitBreakerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.HalfOpenBaseAmountPerStep) {
		query["HalfOpenBaseAmountPerStep"] = request.HalfOpenBaseAmountPerStep
	}

	if !dara.IsNil(request.HalfOpenRecoveryStepNum) {
		query["HalfOpenRecoveryStepNum"] = request.HalfOpenRecoveryStepNum
	}

	if !dara.IsNil(request.MaxAllowedRtMs) {
		query["MaxAllowedRtMs"] = request.MaxAllowedRtMs
	}

	if !dara.IsNil(request.MinRequestAmount) {
		query["MinRequestAmount"] = request.MinRequestAmount
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RetryTimeoutMs) {
		query["RetryTimeoutMs"] = request.RetryTimeoutMs
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.StatIntervalMs) {
		query["StatIntervalMs"] = request.StatIntervalMs
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCircuitBreakerRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCircuitBreakerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a circuit breaking rule.
//
// @param request - UpdateCircuitBreakerRuleRequest
//
// @return UpdateCircuitBreakerRuleResponse
func (client *Client) UpdateCircuitBreakerRule(request *UpdateCircuitBreakerRuleRequest) (_result *UpdateCircuitBreakerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCircuitBreakerRuleResponse{}
	_body, _err := client.UpdateCircuitBreakerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about an instance.
//
// @param request - UpdateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterResponse
func (client *Client) UpdateClusterWithOptions(request *UpdateClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterAliasName) {
		query["ClusterAliasName"] = request.ClusterAliasName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaintenanceEndTime) {
		query["MaintenanceEndTime"] = request.MaintenanceEndTime
	}

	if !dara.IsNil(request.MaintenanceStartTime) {
		query["MaintenanceStartTime"] = request.MaintenanceStartTime
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about an instance.
//
// @param request - UpdateClusterRequest
//
// @return UpdateClusterResponse
func (client *Client) UpdateCluster(request *UpdateClusterRequest) (_result *UpdateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClusterResponse{}
	_body, _err := client.UpdateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the number or specifications of nodes in a pay-as-you-go Microservices Engine (MSE) instance.
//
// Description:
//
// You can call this operation to update the number or specifications of nodes in a pay-as-you-go MSE instance. You are charged when you add nodes or upgrade node specifications. For more information, see [Pricing] (`~~1806469~~`).
//
// @param request - UpdateClusterSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterSpecResponse
func (client *Client) UpdateClusterSpecWithOptions(request *UpdateClusterSpecRequest, runtime *dara.RuntimeOptions) (_result *UpdateClusterSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterSpecification) {
		query["ClusterSpecification"] = request.ClusterSpecification
	}

	if !dara.IsNil(request.InstanceCount) {
		query["InstanceCount"] = request.InstanceCount
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MseVersion) {
		query["MseVersion"] = request.MseVersion
	}

	if !dara.IsNil(request.PubNetworkFlow) {
		query["PubNetworkFlow"] = request.PubNetworkFlow
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClusterSpec"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the number or specifications of nodes in a pay-as-you-go Microservices Engine (MSE) instance.
//
// Description:
//
// You can call this operation to update the number or specifications of nodes in a pay-as-you-go MSE instance. You are charged when you add nodes or upgrade node specifications. For more information, see [Pricing] (`~~1806469~~`).
//
// @param request - UpdateClusterSpecRequest
//
// @return UpdateClusterSpecResponse
func (client *Client) UpdateClusterSpec(request *UpdateClusterSpecRequest) (_result *UpdateClusterSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClusterSpecResponse{}
	_body, _err := client.UpdateClusterSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of an instance.
//
// @param request - UpdateConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigResponse
func (client *Client) UpdateConfigWithOptions(request *UpdateConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AuthEnabled) {
		query["AuthEnabled"] = request.AuthEnabled
	}

	if !dara.IsNil(request.AutopurgePurgeInterval) {
		query["AutopurgePurgeInterval"] = request.AutopurgePurgeInterval
	}

	if !dara.IsNil(request.AutopurgeSnapRetainCount) {
		query["AutopurgeSnapRetainCount"] = request.AutopurgeSnapRetainCount
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConfigAuthEnabled) {
		query["ConfigAuthEnabled"] = request.ConfigAuthEnabled
	}

	if !dara.IsNil(request.ConfigSecretEnabled) {
		query["ConfigSecretEnabled"] = request.ConfigSecretEnabled
	}

	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.ConsoleUIEnabled) {
		query["ConsoleUIEnabled"] = request.ConsoleUIEnabled
	}

	if !dara.IsNil(request.Enable4lw) {
		query["Enable4lw"] = request.Enable4lw
	}

	if !dara.IsNil(request.EurekaSupported) {
		query["EurekaSupported"] = request.EurekaSupported
	}

	if !dara.IsNil(request.ExtendedTypesEnable) {
		query["ExtendedTypesEnable"] = request.ExtendedTypesEnable
	}

	if !dara.IsNil(request.InitLimit) {
		query["InitLimit"] = request.InitLimit
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JuteMaxbuffer) {
		query["JuteMaxbuffer"] = request.JuteMaxbuffer
	}

	if !dara.IsNil(request.MCPEnabled) {
		query["MCPEnabled"] = request.MCPEnabled
	}

	if !dara.IsNil(request.MaxClientCnxns) {
		query["MaxClientCnxns"] = request.MaxClientCnxns
	}

	if !dara.IsNil(request.MaxSessionTimeout) {
		query["MaxSessionTimeout"] = request.MaxSessionTimeout
	}

	if !dara.IsNil(request.MinSessionTimeout) {
		query["MinSessionTimeout"] = request.MinSessionTimeout
	}

	if !dara.IsNil(request.NamingAuthEnabled) {
		query["NamingAuthEnabled"] = request.NamingAuthEnabled
	}

	if !dara.IsNil(request.PassWord) {
		query["PassWord"] = request.PassWord
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.SnapshotCount) {
		query["SnapshotCount"] = request.SnapshotCount
	}

	if !dara.IsNil(request.SyncLimit) {
		query["SyncLimit"] = request.SyncLimit
	}

	if !dara.IsNil(request.TLSEnabled) {
		query["TLSEnabled"] = request.TLSEnabled
	}

	if !dara.IsNil(request.TickTime) {
		query["TickTime"] = request.TickTime
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OpenSuperAcl) {
		body["OpenSuperAcl"] = request.OpenSuperAcl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of an instance.
//
// @param request - UpdateConfigRequest
//
// @return UpdateConfigResponse
func (client *Client) UpdateConfig(request *UpdateConfigRequest) (_result *UpdateConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConfigResponse{}
	_body, _err := client.UpdateConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a namespace for the Nacos engine.
//
// @param request - UpdateEngineNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEngineNamespaceResponse
func (client *Client) UpdateEngineNamespaceWithOptions(request *UpdateEngineNamespaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateEngineNamespaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ServiceCount) {
		query["ServiceCount"] = request.ServiceCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEngineNamespace"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEngineNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a namespace for the Nacos engine.
//
// @param request - UpdateEngineNamespaceRequest
//
// @return UpdateEngineNamespaceResponse
func (client *Client) UpdateEngineNamespace(request *UpdateEngineNamespaceRequest) (_result *UpdateEngineNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEngineNamespaceResponse{}
	_body, _err := client.UpdateEngineNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a throttling rule.
//
// @param request - UpdateFlowRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowRuleResponse
func (client *Client) UpdateFlowRuleWithOptions(request *UpdateFlowRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlowRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ControlBehavior) {
		query["ControlBehavior"] = request.ControlBehavior
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.LimitApp) {
		query["LimitApp"] = request.LimitApp
	}

	if !dara.IsNil(request.MaxQueueingTimeMs) {
		query["MaxQueueingTimeMs"] = request.MaxQueueingTimeMs
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlowRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a throttling rule.
//
// @param request - UpdateFlowRuleRequest
//
// @return UpdateFlowRuleResponse
func (client *Client) UpdateFlowRule(request *UpdateFlowRuleRequest) (_result *UpdateFlowRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFlowRuleResponse{}
	_body, _err := client.UpdateFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新网关认证
//
// @param tmpReq - UpdateGatewayAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayAuthResponse
func (client *Client) UpdateGatewayAuthWithOptions(tmpReq *UpdateGatewayAuthRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayAuthShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthResourceList) {
		request.AuthResourceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthResourceList, dara.String("AuthResourceList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeleteResourceIdList) {
		request.DeleteResourceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeleteResourceIdList, dara.String("DeleteResourceIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExternalAuthZJSON) {
		request.ExternalAuthZJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalAuthZJSON, dara.String("ExternalAuthZJSON"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScopesList) {
		request.ScopesListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScopesList, dara.String("ScopesList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AuthResourceConfig) {
		query["AuthResourceConfig"] = request.AuthResourceConfig
	}

	if !dara.IsNil(request.AuthResourceListShrink) {
		query["AuthResourceList"] = request.AuthResourceListShrink
	}

	if !dara.IsNil(request.AuthResourceMode) {
		query["AuthResourceMode"] = request.AuthResourceMode
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientSecret) {
		query["ClientSecret"] = request.ClientSecret
	}

	if !dara.IsNil(request.CookieDomain) {
		query["CookieDomain"] = request.CookieDomain
	}

	if !dara.IsNil(request.DeleteResourceIdListShrink) {
		query["DeleteResourceIdList"] = request.DeleteResourceIdListShrink
	}

	if !dara.IsNil(request.ExternalAuthZJSONShrink) {
		query["ExternalAuthZJSON"] = request.ExternalAuthZJSONShrink
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.IsWhite) {
		query["IsWhite"] = request.IsWhite
	}

	if !dara.IsNil(request.Issuer) {
		query["Issuer"] = request.Issuer
	}

	if !dara.IsNil(request.Jwks) {
		query["Jwks"] = request.Jwks
	}

	if !dara.IsNil(request.LoginUrl) {
		query["LoginUrl"] = request.LoginUrl
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RedirectUrl) {
		query["RedirectUrl"] = request.RedirectUrl
	}

	if !dara.IsNil(request.ScopesListShrink) {
		query["ScopesList"] = request.ScopesListShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Sub) {
		query["Sub"] = request.Sub
	}

	if !dara.IsNil(request.TokenName) {
		query["TokenName"] = request.TokenName
	}

	if !dara.IsNil(request.TokenNamePrefix) {
		query["TokenNamePrefix"] = request.TokenNamePrefix
	}

	if !dara.IsNil(request.TokenPass) {
		query["TokenPass"] = request.TokenPass
	}

	if !dara.IsNil(request.TokenPosition) {
		query["TokenPosition"] = request.TokenPosition
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayAuth"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新网关认证
//
// @param request - UpdateGatewayAuthRequest
//
// @return UpdateGatewayAuthResponse
func (client *Client) UpdateGatewayAuth(request *UpdateGatewayAuthRequest) (_result *UpdateGatewayAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayAuthResponse{}
	_body, _err := client.UpdateGatewayAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the consumer on which a gateway performs authentication operations.
//
// @param request - UpdateGatewayAuthConsumerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayAuthConsumerResponse
func (client *Client) UpdateGatewayAuthConsumerWithOptions(request *UpdateGatewayAuthConsumerRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayAuthConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EncodeType) {
		query["EncodeType"] = request.EncodeType
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Jwks) {
		query["Jwks"] = request.Jwks
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.KeyValue) {
		query["KeyValue"] = request.KeyValue
	}

	if !dara.IsNil(request.TokenName) {
		query["TokenName"] = request.TokenName
	}

	if !dara.IsNil(request.TokenPass) {
		query["TokenPass"] = request.TokenPass
	}

	if !dara.IsNil(request.TokenPosition) {
		query["TokenPosition"] = request.TokenPosition
	}

	if !dara.IsNil(request.TokenPrefix) {
		query["TokenPrefix"] = request.TokenPrefix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayAuthConsumer"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayAuthConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the consumer on which a gateway performs authentication operations.
//
// @param request - UpdateGatewayAuthConsumerRequest
//
// @return UpdateGatewayAuthConsumerResponse
func (client *Client) UpdateGatewayAuthConsumer(request *UpdateGatewayAuthConsumerRequest) (_result *UpdateGatewayAuthConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayAuthConsumerResponse{}
	_body, _err := client.UpdateGatewayAuthConsumerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a list of resources on which permissions are granted to a gateway authentication consumer.
//
// @param tmpReq - UpdateGatewayAuthConsumerResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayAuthConsumerResourceResponse
func (client *Client) UpdateGatewayAuthConsumerResourceWithOptions(tmpReq *UpdateGatewayAuthConsumerResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayAuthConsumerResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayAuthConsumerResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceList) {
		request.ResourceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceList, dara.String("ResourceList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ResourceListShrink) {
		query["ResourceList"] = request.ResourceListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayAuthConsumerResource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayAuthConsumerResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a list of resources on which permissions are granted to a gateway authentication consumer.
//
// @param request - UpdateGatewayAuthConsumerResourceRequest
//
// @return UpdateGatewayAuthConsumerResourceResponse
func (client *Client) UpdateGatewayAuthConsumerResource(request *UpdateGatewayAuthConsumerResourceRequest) (_result *UpdateGatewayAuthConsumerResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayAuthConsumerResourceResponse{}
	_body, _err := client.UpdateGatewayAuthConsumerResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the resource authorization status for the consumer on which a gateway performs authentication operations.
//
// @param request - UpdateGatewayAuthConsumerResourceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayAuthConsumerResourceStatusResponse
func (client *Client) UpdateGatewayAuthConsumerResourceStatusWithOptions(request *UpdateGatewayAuthConsumerResourceStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayAuthConsumerResourceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.IdList) {
		query["IdList"] = request.IdList
	}

	if !dara.IsNil(request.ResourceStatus) {
		query["ResourceStatus"] = request.ResourceStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayAuthConsumerResourceStatus"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayAuthConsumerResourceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the resource authorization status for the consumer on which a gateway performs authentication operations.
//
// @param request - UpdateGatewayAuthConsumerResourceStatusRequest
//
// @return UpdateGatewayAuthConsumerResourceStatusResponse
func (client *Client) UpdateGatewayAuthConsumerResourceStatus(request *UpdateGatewayAuthConsumerResourceStatusRequest) (_result *UpdateGatewayAuthConsumerResourceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayAuthConsumerResourceStatusResponse{}
	_body, _err := client.UpdateGatewayAuthConsumerResourceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of the consumer on which a gateway performs authentication operations.
//
// @param request - UpdateGatewayAuthConsumerStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayAuthConsumerStatusResponse
func (client *Client) UpdateGatewayAuthConsumerStatusWithOptions(request *UpdateGatewayAuthConsumerStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayAuthConsumerStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConsumerStatus) {
		query["ConsumerStatus"] = request.ConsumerStatus
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayAuthConsumerStatus"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayAuthConsumerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of the consumer on which a gateway performs authentication operations.
//
// @param request - UpdateGatewayAuthConsumerStatusRequest
//
// @return UpdateGatewayAuthConsumerStatusResponse
func (client *Client) UpdateGatewayAuthConsumerStatus(request *UpdateGatewayAuthConsumerStatusRequest) (_result *UpdateGatewayAuthConsumerStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayAuthConsumerStatusResponse{}
	_body, _err := client.UpdateGatewayAuthConsumerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新网关路由熔断规则
//
// @param request - UpdateGatewayCircuitBreakerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayCircuitBreakerRuleResponse
func (client *Client) UpdateGatewayCircuitBreakerRuleWithOptions(request *UpdateGatewayCircuitBreakerRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayCircuitBreakerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.BehaviorType) {
		query["BehaviorType"] = request.BehaviorType
	}

	if !dara.IsNil(request.BodyEncoding) {
		query["BodyEncoding"] = request.BodyEncoding
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MaxAllowedMs) {
		query["MaxAllowedMs"] = request.MaxAllowedMs
	}

	if !dara.IsNil(request.MinRequestAmount) {
		query["MinRequestAmount"] = request.MinRequestAmount
	}

	if !dara.IsNil(request.RecoveryTimeoutSec) {
		query["RecoveryTimeoutSec"] = request.RecoveryTimeoutSec
	}

	if !dara.IsNil(request.ResponseContentBody) {
		query["ResponseContentBody"] = request.ResponseContentBody
	}

	if !dara.IsNil(request.ResponseRedirectUrl) {
		query["ResponseRedirectUrl"] = request.ResponseRedirectUrl
	}

	if !dara.IsNil(request.ResponseStatusCode) {
		query["ResponseStatusCode"] = request.ResponseStatusCode
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.StatDurationSec) {
		query["StatDurationSec"] = request.StatDurationSec
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	if !dara.IsNil(request.TriggerRatio) {
		query["TriggerRatio"] = request.TriggerRatio
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayCircuitBreakerRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayCircuitBreakerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新网关路由熔断规则
//
// @param request - UpdateGatewayCircuitBreakerRuleRequest
//
// @return UpdateGatewayCircuitBreakerRuleResponse
func (client *Client) UpdateGatewayCircuitBreakerRule(request *UpdateGatewayCircuitBreakerRuleRequest) (_result *UpdateGatewayCircuitBreakerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayCircuitBreakerRuleResponse{}
	_body, _err := client.UpdateGatewayCircuitBreakerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新网关配置
//
// @param request - UpdateGatewayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayConfigResponse
func (client *Client) UpdateGatewayConfigWithOptions(request *UpdateGatewayConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.ConfigValue) {
		query["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新网关配置
//
// @param request - UpdateGatewayConfigRequest
//
// @return UpdateGatewayConfigResponse
func (client *Client) UpdateGatewayConfig(request *UpdateGatewayConfigRequest) (_result *UpdateGatewayConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayConfigResponse{}
	_body, _err := client.UpdateGatewayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about the domain name associated with a gateway.
//
// @param tmpReq - UpdateGatewayDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayDomainResponse
func (client *Client) UpdateGatewayDomainWithOptions(tmpReq *UpdateGatewayDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TlsCipherSuitesConfigJSON) {
		request.TlsCipherSuitesConfigJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TlsCipherSuitesConfigJSON, dara.String("TlsCipherSuitesConfigJSON"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Http2) {
		query["Http2"] = request.Http2
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MustHttps) {
		query["MustHttps"] = request.MustHttps
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.TlsCipherSuitesConfigJSONShrink) {
		query["TlsCipherSuitesConfigJSON"] = request.TlsCipherSuitesConfigJSONShrink
	}

	if !dara.IsNil(request.TlsMax) {
		query["TlsMax"] = request.TlsMax
	}

	if !dara.IsNil(request.TlsMin) {
		query["TlsMin"] = request.TlsMin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayDomain"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about the domain name associated with a gateway.
//
// @param request - UpdateGatewayDomainRequest
//
// @return UpdateGatewayDomainResponse
func (client *Client) UpdateGatewayDomain(request *UpdateGatewayDomainRequest) (_result *UpdateGatewayDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayDomainResponse{}
	_body, _err := client.UpdateGatewayDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新网关路由流控规则
//
// @param request - UpdateGatewayFlowRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayFlowRuleResponse
func (client *Client) UpdateGatewayFlowRuleWithOptions(request *UpdateGatewayFlowRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayFlowRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.BehaviorType) {
		query["BehaviorType"] = request.BehaviorType
	}

	if !dara.IsNil(request.BodyEncoding) {
		query["BodyEncoding"] = request.BodyEncoding
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ResponseContentBody) {
		query["ResponseContentBody"] = request.ResponseContentBody
	}

	if !dara.IsNil(request.ResponseRedirectUrl) {
		query["ResponseRedirectUrl"] = request.ResponseRedirectUrl
	}

	if !dara.IsNil(request.ResponseStatusCode) {
		query["ResponseStatusCode"] = request.ResponseStatusCode
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayFlowRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新网关路由流控规则
//
// @param request - UpdateGatewayFlowRuleRequest
//
// @return UpdateGatewayFlowRuleResponse
func (client *Client) UpdateGatewayFlowRule(request *UpdateGatewayFlowRuleRequest) (_result *UpdateGatewayFlowRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayFlowRuleResponse{}
	_body, _err := client.UpdateGatewayFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新网关路由隔离规则
//
// @param request - UpdateGatewayIsolationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayIsolationRuleResponse
func (client *Client) UpdateGatewayIsolationRuleWithOptions(request *UpdateGatewayIsolationRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayIsolationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.BehaviorType) {
		query["BehaviorType"] = request.BehaviorType
	}

	if !dara.IsNil(request.BodyEncoding) {
		query["BodyEncoding"] = request.BodyEncoding
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MaxConcurrency) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.ResponseContentBody) {
		query["ResponseContentBody"] = request.ResponseContentBody
	}

	if !dara.IsNil(request.ResponseRedirectUrl) {
		query["ResponseRedirectUrl"] = request.ResponseRedirectUrl
	}

	if !dara.IsNil(request.ResponseStatusCode) {
		query["ResponseStatusCode"] = request.ResponseStatusCode
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	if !dara.IsNil(request.RouteName) {
		query["RouteName"] = request.RouteName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayIsolationRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新网关路由隔离规则
//
// @param request - UpdateGatewayIsolationRuleRequest
//
// @return UpdateGatewayIsolationRuleResponse
func (client *Client) UpdateGatewayIsolationRule(request *UpdateGatewayIsolationRuleRequest) (_result *UpdateGatewayIsolationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayIsolationRuleResponse{}
	_body, _err := client.UpdateGatewayIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a gateway.
//
// @param request - UpdateGatewayNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayNameResponse
func (client *Client) UpdateGatewayNameWithOptions(request *UpdateGatewayNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayName"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a gateway.
//
// @param request - UpdateGatewayNameRequest
//
// @return UpdateGatewayNameResponse
func (client *Client) UpdateGatewayName(request *UpdateGatewayNameRequest) (_result *UpdateGatewayNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayNameResponse{}
	_body, _err := client.UpdateGatewayNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of a gateway.
//
// @param tmpReq - UpdateGatewayOptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayOptionResponse
func (client *Client) UpdateGatewayOptionWithOptions(tmpReq *UpdateGatewayOptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayOptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayOptionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GatewayOption) {
		request.GatewayOptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GatewayOption, dara.String("GatewayOption"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayOptionShrink) {
		query["GatewayOption"] = request.GatewayOptionShrink
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayOption"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a gateway.
//
// @param request - UpdateGatewayOptionRequest
//
// @return UpdateGatewayOptionResponse
func (client *Client) UpdateGatewayOption(request *UpdateGatewayOptionRequest) (_result *UpdateGatewayOptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayOptionResponse{}
	_body, _err := client.UpdateGatewayOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a route for a gateway.
//
// @param tmpReq - UpdateGatewayRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteResponse
func (client *Client) UpdateGatewayRouteWithOptions(tmpReq *UpdateGatewayRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayRouteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DirectResponseJSON) {
		request.DirectResponseJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DirectResponseJSON, dara.String("DirectResponseJSON"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FallbackServices) {
		request.FallbackServicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FallbackServices, dara.String("FallbackServices"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Predicates) {
		request.PredicatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Predicates, dara.String("Predicates"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RedirectJSON) {
		request.RedirectJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RedirectJSON, dara.String("RedirectJSON"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Services) {
		request.ServicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Services, dara.String("Services"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.DirectResponseJSONShrink) {
		query["DirectResponseJSON"] = request.DirectResponseJSONShrink
	}

	if !dara.IsNil(request.DomainIdListJSON) {
		query["DomainIdListJSON"] = request.DomainIdListJSON
	}

	if !dara.IsNil(request.EnableWaf) {
		query["EnableWaf"] = request.EnableWaf
	}

	if !dara.IsNil(request.Fallback) {
		query["Fallback"] = request.Fallback
	}

	if !dara.IsNil(request.FallbackServicesShrink) {
		query["FallbackServices"] = request.FallbackServicesShrink
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PredicatesShrink) {
		query["Predicates"] = request.PredicatesShrink
	}

	if !dara.IsNil(request.RedirectJSONShrink) {
		query["RedirectJSON"] = request.RedirectJSONShrink
	}

	if !dara.IsNil(request.RouteOrder) {
		query["RouteOrder"] = request.RouteOrder
	}

	if !dara.IsNil(request.ServicesShrink) {
		query["Services"] = request.ServicesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayRoute"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a route for a gateway.
//
// @param request - UpdateGatewayRouteRequest
//
// @return UpdateGatewayRouteResponse
func (client *Client) UpdateGatewayRoute(request *UpdateGatewayRouteRequest) (_result *UpdateGatewayRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayRouteResponse{}
	_body, _err := client.UpdateGatewayRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the authentication configurations of a route.
//
// @param tmpReq - UpdateGatewayRouteAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteAuthResponse
func (client *Client) UpdateGatewayRouteAuthWithOptions(tmpReq *UpdateGatewayRouteAuthRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayRouteAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayRouteAuthShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthJSON) {
		request.AuthJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthJSON, dara.String("AuthJSON"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AuthJSONShrink) {
		query["AuthJSON"] = request.AuthJSONShrink
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayRouteAuth"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayRouteAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the authentication configurations of a route.
//
// @param request - UpdateGatewayRouteAuthRequest
//
// @return UpdateGatewayRouteAuthResponse
func (client *Client) UpdateGatewayRouteAuth(request *UpdateGatewayRouteAuthRequest) (_result *UpdateGatewayRouteAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayRouteAuthResponse{}
	_body, _err := client.UpdateGatewayRouteAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the cross-origin resource sharing (CORS) policy of a route.
//
// @param tmpReq - UpdateGatewayRouteCORSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteCORSResponse
func (client *Client) UpdateGatewayRouteCORSWithOptions(tmpReq *UpdateGatewayRouteCORSRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayRouteCORSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayRouteCORSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CorsJSON) {
		request.CorsJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CorsJSON, dara.String("CorsJSON"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CorsJSONShrink) {
		query["CorsJSON"] = request.CorsJSONShrink
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayRouteCORS"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayRouteCORSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the cross-origin resource sharing (CORS) policy of a route.
//
// @param request - UpdateGatewayRouteCORSRequest
//
// @return UpdateGatewayRouteCORSResponse
func (client *Client) UpdateGatewayRouteCORS(request *UpdateGatewayRouteCORSRequest) (_result *UpdateGatewayRouteCORSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayRouteCORSResponse{}
	_body, _err := client.UpdateGatewayRouteCORSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the rewrite policy of a route for a gateway.
//
// @param request - UpdateGatewayRouteHTTPRewriteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteHTTPRewriteResponse
func (client *Client) UpdateGatewayRouteHTTPRewriteWithOptions(request *UpdateGatewayRouteHTTPRewriteRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayRouteHTTPRewriteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.HttpRewriteJSON) {
		query["HttpRewriteJSON"] = request.HttpRewriteJSON
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayRouteHTTPRewrite"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayRouteHTTPRewriteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the rewrite policy of a route for a gateway.
//
// @param request - UpdateGatewayRouteHTTPRewriteRequest
//
// @return UpdateGatewayRouteHTTPRewriteResponse
func (client *Client) UpdateGatewayRouteHTTPRewrite(request *UpdateGatewayRouteHTTPRewriteRequest) (_result *UpdateGatewayRouteHTTPRewriteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayRouteHTTPRewriteResponse{}
	_body, _err := client.UpdateGatewayRouteHTTPRewriteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the header configuration policy of a route.
//
// @param request - UpdateGatewayRouteHeaderOpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteHeaderOpResponse
func (client *Client) UpdateGatewayRouteHeaderOpWithOptions(request *UpdateGatewayRouteHeaderOpRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayRouteHeaderOpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.HeaderOpJSON) {
		query["HeaderOpJSON"] = request.HeaderOpJSON
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayRouteHeaderOp"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayRouteHeaderOpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the header configuration policy of a route.
//
// @param request - UpdateGatewayRouteHeaderOpRequest
//
// @return UpdateGatewayRouteHeaderOpResponse
func (client *Client) UpdateGatewayRouteHeaderOp(request *UpdateGatewayRouteHeaderOpRequest) (_result *UpdateGatewayRouteHeaderOpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayRouteHeaderOpResponse{}
	_body, _err := client.UpdateGatewayRouteHeaderOpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the retry policy of a route.
//
// @param tmpReq - UpdateGatewayRouteRetryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteRetryResponse
func (client *Client) UpdateGatewayRouteRetryWithOptions(tmpReq *UpdateGatewayRouteRetryRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayRouteRetryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayRouteRetryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RetryJSON) {
		request.RetryJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RetryJSON, dara.String("RetryJSON"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.RetryJSONShrink) {
		query["RetryJSON"] = request.RetryJSONShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayRouteRetry"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayRouteRetryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the retry policy of a route.
//
// @param request - UpdateGatewayRouteRetryRequest
//
// @return UpdateGatewayRouteRetryResponse
func (client *Client) UpdateGatewayRouteRetry(request *UpdateGatewayRouteRetryRequest) (_result *UpdateGatewayRouteRetryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayRouteRetryResponse{}
	_body, _err := client.UpdateGatewayRouteRetryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the timeout policy of a route.
//
// @param tmpReq - UpdateGatewayRouteTimeoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteTimeoutResponse
func (client *Client) UpdateGatewayRouteTimeoutWithOptions(tmpReq *UpdateGatewayRouteTimeoutRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayRouteTimeoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayRouteTimeoutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TimeoutJSON) {
		request.TimeoutJSONShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TimeoutJSON, dara.String("TimeoutJSON"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.TimeoutJSONShrink) {
		query["TimeoutJSON"] = request.TimeoutJSONShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayRouteTimeout"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayRouteTimeoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the timeout policy of a route.
//
// @param request - UpdateGatewayRouteTimeoutRequest
//
// @return UpdateGatewayRouteTimeoutResponse
func (client *Client) UpdateGatewayRouteTimeout(request *UpdateGatewayRouteTimeoutRequest) (_result *UpdateGatewayRouteTimeoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayRouteTimeoutResponse{}
	_body, _err := client.UpdateGatewayRouteTimeoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the WAF status of a route.
//
// @param request - UpdateGatewayRouteWafStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayRouteWafStatusResponse
func (client *Client) UpdateGatewayRouteWafStatusWithOptions(request *UpdateGatewayRouteWafStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayRouteWafStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.EnableWaf) {
		query["EnableWaf"] = request.EnableWaf
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.RouteId) {
		query["RouteId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayRouteWafStatus"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayRouteWafStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the WAF status of a route.
//
// @param request - UpdateGatewayRouteWafStatusRequest
//
// @return UpdateGatewayRouteWafStatusResponse
func (client *Client) UpdateGatewayRouteWafStatus(request *UpdateGatewayRouteWafStatusRequest) (_result *UpdateGatewayRouteWafStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayRouteWafStatusResponse{}
	_body, _err := client.UpdateGatewayRouteWafStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新服务
//
// @param tmpReq - UpdateGatewayServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayServiceResponse
func (client *Client) UpdateGatewayServiceWithOptions(tmpReq *UpdateGatewayServiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DnsServerList) {
		request.DnsServerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DnsServerList, dara.String("DnsServerList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IpList) {
		request.IpListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpList, dara.String("IpList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.DnsServerListShrink) {
		query["DnsServerList"] = request.DnsServerListShrink
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.IpListShrink) {
		query["IpList"] = request.IpListShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ServicePort) {
		query["ServicePort"] = request.ServicePort
	}

	if !dara.IsNil(request.ServiceProtocol) {
		query["ServiceProtocol"] = request.ServiceProtocol
	}

	if !dara.IsNil(request.TlsSetting) {
		query["TlsSetting"] = request.TlsSetting
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayService"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新服务
//
// @param request - UpdateGatewayServiceRequest
//
// @return UpdateGatewayServiceResponse
func (client *Client) UpdateGatewayService(request *UpdateGatewayServiceRequest) (_result *UpdateGatewayServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayServiceResponse{}
	_body, _err := client.UpdateGatewayServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the health check policy of a specified service in a cloud-native gateway.
//
// @param tmpReq - UpdateGatewayServiceCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayServiceCheckResponse
func (client *Client) UpdateGatewayServiceCheckWithOptions(tmpReq *UpdateGatewayServiceCheckRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayServiceCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayServiceCheckShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExpectedStatuses) {
		request.ExpectedStatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExpectedStatuses, dara.String("ExpectedStatuses"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Check) {
		query["Check"] = request.Check
	}

	if !dara.IsNil(request.ExpectedStatusesShrink) {
		query["ExpectedStatuses"] = request.ExpectedStatusesShrink
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.HttpHost) {
		query["HttpHost"] = request.HttpHost
	}

	if !dara.IsNil(request.HttpPath) {
		query["HttpPath"] = request.HttpPath
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayServiceCheck"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayServiceCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the health check policy of a specified service in a cloud-native gateway.
//
// @param request - UpdateGatewayServiceCheckRequest
//
// @return UpdateGatewayServiceCheckResponse
func (client *Client) UpdateGatewayServiceCheck(request *UpdateGatewayServiceCheckRequest) (_result *UpdateGatewayServiceCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayServiceCheckResponse{}
	_body, _err := client.UpdateGatewayServiceCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the traffic policy of a service.
//
// @param tmpReq - UpdateGatewayServiceTrafficPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayServiceTrafficPolicyResponse
func (client *Client) UpdateGatewayServiceTrafficPolicyWithOptions(tmpReq *UpdateGatewayServiceTrafficPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayServiceTrafficPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateGatewayServiceTrafficPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GatewayTrafficPolicy) {
		request.GatewayTrafficPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GatewayTrafficPolicy, dara.String("GatewayTrafficPolicy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayTrafficPolicyShrink) {
		query["GatewayTrafficPolicy"] = request.GatewayTrafficPolicyShrink
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayServiceTrafficPolicy"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayServiceTrafficPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the traffic policy of a service.
//
// @param request - UpdateGatewayServiceTrafficPolicyRequest
//
// @return UpdateGatewayServiceTrafficPolicyResponse
func (client *Client) UpdateGatewayServiceTrafficPolicy(request *UpdateGatewayServiceTrafficPolicyRequest) (_result *UpdateGatewayServiceTrafficPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayServiceTrafficPolicyResponse{}
	_body, _err := client.UpdateGatewayServiceTrafficPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the version of a service.
//
// @param request - UpdateGatewayServiceVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayServiceVersionResponse
func (client *Client) UpdateGatewayServiceVersionWithOptions(request *UpdateGatewayServiceVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewayServiceVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayServiceVersion"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the version of a service.
//
// @param request - UpdateGatewayServiceVersionRequest
//
// @return UpdateGatewayServiceVersionResponse
func (client *Client) UpdateGatewayServiceVersion(request *UpdateGatewayServiceVersionRequest) (_result *UpdateGatewayServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewayServiceVersionResponse{}
	_body, _err := client.UpdateGatewayServiceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the number of nodes or the specifications of nodes in a pay-as-you-go or subscription cloud-native gateway.
//
// Description:
//
// You can call this operation to update the number of nodes or the specifications of nodes in a pay-as-you-go or subscription cloud-native gateway. If you add nodes or increase the specifications, you will incur fees. For more information, see [Pricing](https://help.aliyun.com/document_detail/250950.html).
//
// @param request - UpdateGatewaySpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewaySpecResponse
func (client *Client) UpdateGatewaySpecWithOptions(request *UpdateGatewaySpecRequest, runtime *dara.RuntimeOptions) (_result *UpdateGatewaySpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Replica) {
		query["Replica"] = request.Replica
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewaySpec"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewaySpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the number of nodes or the specifications of nodes in a pay-as-you-go or subscription cloud-native gateway.
//
// Description:
//
// You can call this operation to update the number of nodes or the specifications of nodes in a pay-as-you-go or subscription cloud-native gateway. If you add nodes or increase the specifications, you will incur fees. For more information, see [Pricing](https://help.aliyun.com/document_detail/250950.html).
//
// @param request - UpdateGatewaySpecRequest
//
// @return UpdateGatewaySpecResponse
func (client *Client) UpdateGatewaySpec(request *UpdateGatewaySpecRequest) (_result *UpdateGatewaySpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGatewaySpecResponse{}
	_body, _err := client.UpdateGatewaySpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the version number of the destination cluster.
//
// @param request - UpdateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageResponse
func (client *Client) UpdateImageWithOptions(request *UpdateImageRequest, runtime *dara.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.VersionCode) {
		query["VersionCode"] = request.VersionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateImage"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version number of the destination cluster.
//
// @param request - UpdateImageRequest
//
// @return UpdateImageResponse
func (client *Client) UpdateImage(request *UpdateImageRequest) (_result *UpdateImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateImageResponse{}
	_body, _err := client.UpdateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新隔离规则
//
// @param request - UpdateIsolationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIsolationRuleResponse
func (client *Client) UpdateIsolationRuleWithOptions(request *UpdateIsolationRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateIsolationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.LimitApp) {
		query["LimitApp"] = request.LimitApp
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIsolationRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIsolationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新隔离规则
//
// @param request - UpdateIsolationRuleRequest
//
// @return UpdateIsolationRuleResponse
func (client *Client) UpdateIsolationRule(request *UpdateIsolationRuleRequest) (_result *UpdateIsolationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIsolationRuleResponse{}
	_body, _err := client.UpdateIsolationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新同AZ路由规则
//
// @param request - UpdateLocalityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLocalityRuleResponse
func (client *Client) UpdateLocalityRuleWithOptions(request *UpdateLocalityRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateLocalityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLocalityRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLocalityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新同AZ路由规则
//
// @param request - UpdateLocalityRuleRequest
//
// @return UpdateLocalityRuleResponse
func (client *Client) UpdateLocalityRule(request *UpdateLocalityRuleRequest) (_result *UpdateLocalityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLocalityRuleResponse{}
	_body, _err := client.UpdateLocalityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a canary release for messaging of an application.
//
// @param tmpReq - UpdateMessageQueueRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageQueueRouteResponse
func (client *Client) UpdateMessageQueueRouteWithOptions(tmpReq *UpdateMessageQueueRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateMessageQueueRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMessageQueueRouteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GrayBaseTags) {
		request.GrayBaseTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GrayBaseTags, dara.String("GrayBaseTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.FilterSide) {
		query["FilterSide"] = request.FilterSide
	}

	if !dara.IsNil(request.GrayBaseTagsShrink) {
		query["GrayBaseTags"] = request.GrayBaseTagsShrink
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMessageQueueRoute"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMessageQueueRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a canary release for messaging of an application.
//
// @param request - UpdateMessageQueueRouteRequest
//
// @return UpdateMessageQueueRouteResponse
func (client *Client) UpdateMessageQueueRoute(request *UpdateMessageQueueRouteRequest) (_result *UpdateMessageQueueRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMessageQueueRouteResponse{}
	_body, _err := client.UpdateMessageQueueRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a migration task.
//
// @param request - UpdateMigrationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMigrationTaskResponse
func (client *Client) UpdateMigrationTaskWithOptions(request *UpdateMigrationTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateMigrationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OriginInstanceAddress) {
		query["OriginInstanceAddress"] = request.OriginInstanceAddress
	}

	if !dara.IsNil(request.OriginInstanceName) {
		query["OriginInstanceName"] = request.OriginInstanceName
	}

	if !dara.IsNil(request.OriginInstanceNamespace) {
		query["OriginInstanceNamespace"] = request.OriginInstanceNamespace
	}

	if !dara.IsNil(request.ProjectDesc) {
		query["ProjectDesc"] = request.ProjectDesc
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.SyncType) {
		query["SyncType"] = request.SyncType
	}

	if !dara.IsNil(request.TargetClusterName) {
		query["TargetClusterName"] = request.TargetClusterName
	}

	if !dara.IsNil(request.TargetClusterUrl) {
		query["TargetClusterUrl"] = request.TargetClusterUrl
	}

	if !dara.IsNil(request.TargetInstanceId) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMigrationTask"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMigrationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a migration task.
//
// @param request - UpdateMigrationTaskRequest
//
// @return UpdateMigrationTaskResponse
func (client *Client) UpdateMigrationTask(request *UpdateMigrationTaskRequest) (_result *UpdateMigrationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMigrationTaskResponse{}
	_body, _err := client.UpdateMigrationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a Nacos cluster.
//
// @param request - UpdateNacosClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacosClusterResponse
func (client *Client) UpdateNacosClusterWithOptions(request *UpdateNacosClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateNacosClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CheckPort) {
		query["CheckPort"] = request.CheckPort
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.HealthChecker) {
		query["HealthChecker"] = request.HealthChecker
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.UseInstancePortForCheck) {
		query["UseInstancePortForCheck"] = request.UseInstancePortForCheck
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNacosCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNacosClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a Nacos cluster.
//
// @param request - UpdateNacosClusterRequest
//
// @return UpdateNacosClusterResponse
func (client *Client) UpdateNacosCluster(request *UpdateNacosClusterRequest) (_result *UpdateNacosClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNacosClusterResponse{}
	_body, _err := client.UpdateNacosClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a Nacos configuration.
//
// Description:
//
// >  The current API operation is not provided in Nacos SDK. For more information about Nacos SDK, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - UpdateNacosConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacosConfigResponse
func (client *Client) UpdateNacosConfigWithOptions(request *UpdateNacosConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateNacosConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BetaIps) {
		query["BetaIps"] = request.BetaIps
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.EncryptedDataKey) {
		query["EncryptedDataKey"] = request.EncryptedDataKey
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNacosConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNacosConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a Nacos configuration.
//
// Description:
//
// >  The current API operation is not provided in Nacos SDK. For more information about Nacos SDK, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - UpdateNacosConfigRequest
//
// @return UpdateNacosConfigResponse
func (client *Client) UpdateNacosConfig(request *UpdateNacosConfigRequest) (_result *UpdateNacosConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNacosConfigResponse{}
	_body, _err := client.UpdateNacosConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新nacos灰度配置
//
// @param request - UpdateNacosGrayConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacosGrayConfigResponse
func (client *Client) UpdateNacosGrayConfigWithOptions(request *UpdateNacosGrayConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateNacosGrayConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	if !dara.IsNil(request.GrayRule) {
		query["GrayRule"] = request.GrayRule
	}

	if !dara.IsNil(request.GrayRuleName) {
		query["GrayRuleName"] = request.GrayRuleName
	}

	if !dara.IsNil(request.GrayRulePriority) {
		query["GrayRulePriority"] = request.GrayRulePriority
	}

	if !dara.IsNil(request.GrayType) {
		query["GrayType"] = request.GrayType
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.OpType) {
		query["OpType"] = request.OpType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.StopGray) {
		query["StopGray"] = request.StopGray
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNacosGrayConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNacosGrayConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新nacos灰度配置
//
// @param request - UpdateNacosGrayConfigRequest
//
// @return UpdateNacosGrayConfigResponse
func (client *Client) UpdateNacosGrayConfig(request *UpdateNacosGrayConfigRequest) (_result *UpdateNacosGrayConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNacosGrayConfigResponse{}
	_body, _err := client.UpdateNacosGrayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about application instances that are registered with a Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - UpdateNacosInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacosInstanceResponse
func (client *Client) UpdateNacosInstanceWithOptions(request *UpdateNacosInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateNacosInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Ephemeral) {
		query["Ephemeral"] = request.Ephemeral
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Metadata) {
		body["Metadata"] = request.Metadata
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNacosInstance"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNacosInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about application instances that are registered with a Nacos instance.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - UpdateNacosInstanceRequest
//
// @return UpdateNacosInstanceResponse
func (client *Client) UpdateNacosInstance(request *UpdateNacosInstanceRequest) (_result *UpdateNacosInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNacosInstanceResponse{}
	_body, _err := client.UpdateNacosInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a Nacos service.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - UpdateNacosServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacosServiceResponse
func (client *Client) UpdateNacosServiceWithOptions(request *UpdateNacosServiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateNacosServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.ProtectThreshold) {
		query["ProtectThreshold"] = request.ProtectThreshold
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNacosService"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNacosServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a Nacos service.
//
// Description:
//
// > The operation is not provided in Nacos SDKs. For information about Nacos SDKs, see the [official documentation](https://nacos.io/zh-cn/docs/sdk.html).
//
// @param request - UpdateNacosServiceRequest
//
// @return UpdateNacosServiceResponse
func (client *Client) UpdateNacosService(request *UpdateNacosServiceRequest) (_result *UpdateNacosServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNacosServiceResponse{}
	_body, _err := client.UpdateNacosServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a plug-in.
//
// @param tmpReq - UpdatePluginConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePluginConfigResponse
func (client *Client) UpdatePluginConfigWithOptions(tmpReq *UpdatePluginConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdatePluginConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePluginConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceIdList) {
		request.ResourceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIdList, dara.String("ResourceIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ConfigLevel) {
		query["ConfigLevel"] = request.ConfigLevel
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.GmtCreate) {
		query["GmtCreate"] = request.GmtCreate
	}

	if !dara.IsNil(request.GmtModified) {
		query["GmtModified"] = request.GmtModified
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.ResourceIdListShrink) {
		query["ResourceIdList"] = request.ResourceIdListShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePluginConfig"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePluginConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a plug-in.
//
// @param request - UpdatePluginConfigRequest
//
// @return UpdatePluginConfigResponse
func (client *Client) UpdatePluginConfig(request *UpdatePluginConfigRequest) (_result *UpdatePluginConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePluginConfigResponse{}
	_body, _err := client.UpdatePluginConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a certificate.
//
// @param request - UpdateSSLCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSSLCertResponse
func (client *Client) UpdateSSLCertWithOptions(request *UpdateSSLCertRequest, runtime *dara.RuntimeOptions) (_result *UpdateSSLCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSSLCert"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSSLCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a certificate.
//
// @param request - UpdateSSLCertRequest
//
// @return UpdateSSLCertResponse
func (client *Client) UpdateSSLCert(request *UpdateSSLCertRequest) (_result *UpdateSSLCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSSLCertResponse{}
	_body, _err := client.UpdateSSLCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies service sources of a cloud-native gateway. You can modify only Container Service for Kubernetes (ACK) service sources that contain configurations related to Ingress resource monitoring.
//
// @param tmpReq - UpdateServiceSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceSourceResponse
func (client *Client) UpdateServiceSourceWithOptions(tmpReq *UpdateServiceSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateServiceSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateServiceSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IngressOptionsRequest) {
		request.IngressOptionsRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IngressOptionsRequest, dara.String("IngressOptionsRequest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PathList) {
		request.PathListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PathList, dara.String("PathList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayUniqueId) {
		query["GatewayUniqueId"] = request.GatewayUniqueId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.IngressOptionsRequestShrink) {
		query["IngressOptionsRequest"] = request.IngressOptionsRequestShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PathListShrink) {
		query["PathList"] = request.PathListShrink
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceSource"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies service sources of a cloud-native gateway. You can modify only Container Service for Kubernetes (ACK) service sources that contain configurations related to Ingress resource monitoring.
//
// @param request - UpdateServiceSourceRequest
//
// @return UpdateServiceSourceResponse
func (client *Client) UpdateServiceSource(request *UpdateServiceSourceRequest) (_result *UpdateServiceSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateServiceSourceResponse{}
	_body, _err := client.UpdateServiceSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新热点参数防护规则（HTTP 请求）
//
// @param request - UpdateWebFlowRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWebFlowRuleResponse
func (client *Client) UpdateWebFlowRuleWithOptions(request *UpdateWebFlowRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateWebFlowRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Burst) {
		query["Burst"] = request.Burst
	}

	if !dara.IsNil(request.ControlBehavior) {
		query["ControlBehavior"] = request.ControlBehavior
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.MaxQueueingTimeMs) {
		query["MaxQueueingTimeMs"] = request.MaxQueueingTimeMs
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ParamItem) {
		query["ParamItem"] = request.ParamItem
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceMode) {
		query["ResourceMode"] = request.ResourceMode
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.StatIntervalMs) {
		query["StatIntervalMs"] = request.StatIntervalMs
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWebFlowRule"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWebFlowRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新热点参数防护规则（HTTP 请求）
//
// @param request - UpdateWebFlowRuleRequest
//
// @return UpdateWebFlowRuleResponse
func (client *Client) UpdateWebFlowRule(request *UpdateWebFlowRuleRequest) (_result *UpdateWebFlowRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWebFlowRuleResponse{}
	_body, _err := client.UpdateWebFlowRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a ZooKeeper node.
//
// @param request - UpdateZnodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateZnodeResponse
func (client *Client) UpdateZnodeWithOptions(request *UpdateZnodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateZnodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateZnode"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateZnodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a ZooKeeper node.
//
// @param request - UpdateZnodeRequest
//
// @return UpdateZnodeResponse
func (client *Client) UpdateZnode(request *UpdateZnodeRequest) (_result *UpdateZnodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateZnodeResponse{}
	_body, _err := client.UpdateZnodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the version of a cluster.
//
// @param request - UpgradeClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClusterResponse
func (client *Client) UpgradeClusterWithOptions(request *UpgradeClusterRequest, runtime *dara.RuntimeOptions) (_result *UpgradeClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestPars) {
		query["RequestPars"] = request.RequestPars
	}

	if !dara.IsNil(request.UpgradeVersion) {
		query["UpgradeVersion"] = request.UpgradeVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeCluster"),
		Version:     dara.String("2019-05-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the version of a cluster.
//
// @param request - UpgradeClusterRequest
//
// @return UpgradeClusterResponse
func (client *Client) UpgradeCluster(request *UpgradeClusterRequest) (_result *UpgradeClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeClusterResponse{}
	_body, _err := client.UpgradeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
