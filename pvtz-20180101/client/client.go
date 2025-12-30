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
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("pvtz"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a custom line.
//
// @param request - AddCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomLineResponse
func (client *Client) AddCustomLineWithOptions(request *AddCustomLineRequest, runtime *dara.RuntimeOptions) (_result *AddCustomLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DnsCategory) {
		query["DnsCategory"] = request.DnsCategory
	}

	if !dara.IsNil(request.Ipv4s) {
		query["Ipv4s"] = request.Ipv4s
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ShareScope) {
		query["ShareScope"] = request.ShareScope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomLine"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom line.
//
// @param request - AddCustomLineRequest
//
// @return AddCustomLineResponse
func (client *Client) AddCustomLine(request *AddCustomLineRequest) (_result *AddCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCustomLineResponse{}
	_body, _err := client.AddCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an endpoint.
//
// @param request - AddResolverEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResolverEndpointResponse
func (client *Client) AddResolverEndpointWithOptions(request *AddResolverEndpointRequest, runtime *dara.RuntimeOptions) (_result *AddResolverEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpConfig) {
		query["IpConfig"] = request.IpConfig
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcRegionId) {
		query["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddResolverEndpoint"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddResolverEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an endpoint.
//
// @param request - AddResolverEndpointRequest
//
// @return AddResolverEndpointResponse
func (client *Client) AddResolverEndpoint(request *AddResolverEndpointRequest) (_result *AddResolverEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddResolverEndpointResponse{}
	_body, _err := client.AddResolverEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a forwarding rule.
//
// Description:
//
// #### [](#)**Precautions**
//
// If a virtual private cloud (VPC) serves as both an inbound VPC and an outbound VPC, the IP addresses of external Domain Name System (DNS) servers cannot be the same as the IP addresses of the inbound endpoint in the VPC. The IP addresses of the external DNS servers are specified in the forwarding rule associated with the outbound endpoint in the same VPC. If the IP addresses are the same, the DNS requests sent from the IP addresses of the inbound endpoint are returned to the VPC. This results in resolution failures.
//
// @param request - AddResolverRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResolverRuleResponse
func (client *Client) AddResolverRuleWithOptions(request *AddResolverRuleRequest, runtime *dara.RuntimeOptions) (_result *AddResolverRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EdgeDnsClusters) {
		query["EdgeDnsClusters"] = request.EdgeDnsClusters
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.ForwardIp) {
		query["ForwardIp"] = request.ForwardIp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Vpcs) {
		query["Vpcs"] = request.Vpcs
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddResolverRule"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddResolverRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a forwarding rule.
//
// Description:
//
// #### [](#)**Precautions**
//
// If a virtual private cloud (VPC) serves as both an inbound VPC and an outbound VPC, the IP addresses of external Domain Name System (DNS) servers cannot be the same as the IP addresses of the inbound endpoint in the VPC. The IP addresses of the external DNS servers are specified in the forwarding rule associated with the outbound endpoint in the same VPC. If the IP addresses are the same, the DNS requests sent from the IP addresses of the inbound endpoint are returned to the VPC. This results in resolution failures.
//
// @param request - AddResolverRuleRequest
//
// @return AddResolverRuleResponse
func (client *Client) AddResolverRule(request *AddResolverRuleRequest) (_result *AddResolverRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddResolverRuleResponse{}
	_body, _err := client.AddResolverRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds another account to associate one or more virtual private clouds (VPCs) of the current account with a private zone.
//
// Description:
//
// #### **Limits**
//
//   - You can set an effective scope across accounts only by using an Alibaba Cloud account instead of a RAM user. You can set an effective scope across accounts registered on the same site. For example, you can perform the operation across accounts that are both registered on the Alibaba Cloud China site or Alibaba Cloud international site. You cannot set an effective scope across accounts registered on different sites. For example, you cannot perform the operation across accounts that are separately registered on the Alibaba Cloud China site and Alibaba Cloud international site.
//
//   - No API operation is provided for sending the verification codes that are required for authorization.
//
// #### **Precautions**
//
// If you set an effective scope across accounts, bills are settled within the account that is used to perform routine management on built-in authoritative zones.
//
// @param request - AddUserVpcAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserVpcAuthorizationResponse
func (client *Client) AddUserVpcAuthorizationWithOptions(request *AddUserVpcAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *AddUserVpcAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthChannel) {
		query["AuthChannel"] = request.AuthChannel
	}

	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.AuthorizedUserId) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserVpcAuthorization"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserVpcAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds another account to associate one or more virtual private clouds (VPCs) of the current account with a private zone.
//
// Description:
//
// #### **Limits**
//
//   - You can set an effective scope across accounts only by using an Alibaba Cloud account instead of a RAM user. You can set an effective scope across accounts registered on the same site. For example, you can perform the operation across accounts that are both registered on the Alibaba Cloud China site or Alibaba Cloud international site. You cannot set an effective scope across accounts registered on different sites. For example, you cannot perform the operation across accounts that are separately registered on the Alibaba Cloud China site and Alibaba Cloud international site.
//
//   - No API operation is provided for sending the verification codes that are required for authorization.
//
// #### **Precautions**
//
// If you set an effective scope across accounts, bills are settled within the account that is used to perform routine management on built-in authoritative zones.
//
// @param request - AddUserVpcAuthorizationRequest
//
// @return AddUserVpcAuthorizationResponse
func (client *Client) AddUserVpcAuthorization(request *AddUserVpcAuthorizationRequest) (_result *AddUserVpcAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserVpcAuthorizationResponse{}
	_body, _err := client.AddUserVpcAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a built-in authoritative zone in the regular module or acceleration module.
//
// @param request - AddZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneResponse
func (client *Client) AddZoneWithOptions(request *AddZoneRequest, runtime *dara.RuntimeOptions) (_result *AddZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DnsGroup) {
		query["DnsGroup"] = request.DnsGroup
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProxyPattern) {
		query["ProxyPattern"] = request.ProxyPattern
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	if !dara.IsNil(request.ZoneTag) {
		query["ZoneTag"] = request.ZoneTag
	}

	if !dara.IsNil(request.ZoneType) {
		query["ZoneType"] = request.ZoneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddZone"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a built-in authoritative zone in the regular module or acceleration module.
//
// @param request - AddZoneRequest
//
// @return AddZoneResponse
func (client *Client) AddZone(request *AddZoneRequest) (_result *AddZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddZoneResponse{}
	_body, _err := client.AddZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a Domain Name System (DNS) record for a built-in authoritative zone. Within the effective scope, the intranet DNS records rather than the Internet DNS records take effect for the zone.
//
// @param request - AddZoneRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneRecordResponse
func (client *Client) AddZoneRecordWithOptions(request *AddZoneRecordRequest, runtime *dara.RuntimeOptions) (_result *AddZoneRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddZoneRecord"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddZoneRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a Domain Name System (DNS) record for a built-in authoritative zone. Within the effective scope, the intranet DNS records rather than the Internet DNS records take effect for the zone.
//
// @param request - AddZoneRecordRequest
//
// @return AddZoneRecordResponse
func (client *Client) AddZoneRecord(request *AddZoneRecordRequest) (_result *AddZoneRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddZoneRecordResponse{}
	_body, _err := client.AddZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a forwarding rule with virtual private clouds (VPCs).
//
// @param request - BindResolverRuleVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindResolverRuleVpcResponse
func (client *Client) BindResolverRuleVpcWithOptions(request *BindResolverRuleVpcRequest, runtime *dara.RuntimeOptions) (_result *BindResolverRuleVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.Vpc) {
		query["Vpc"] = request.Vpc
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindResolverRuleVpc"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindResolverRuleVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a forwarding rule with virtual private clouds (VPCs).
//
// @param request - BindResolverRuleVpcRequest
//
// @return BindResolverRuleVpcResponse
func (client *Client) BindResolverRuleVpc(request *BindResolverRuleVpcRequest) (_result *BindResolverRuleVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindResolverRuleVpcResponse{}
	_body, _err := client.BindResolverRuleVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates or dissociates virtual private clouds (VPCs) from a zone to set the effective scope of the zone.
//
// Description:
//
// ##### [](#)Precautions:
//
// We recommend that you set the effective scope of a zone after you configure all Domain Name System (DNS) records. If you set an effective scope before you configure DNS records, the DNS resolution for the zone within the effective scope will fail unless you enable the recursive resolution proxy for subdomain names.
//
// @param request - BindZoneVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindZoneVpcResponse
func (client *Client) BindZoneVpcWithOptions(request *BindZoneVpcRequest, runtime *dara.RuntimeOptions) (_result *BindZoneVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Vpcs) {
		query["Vpcs"] = request.Vpcs
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindZoneVpc"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindZoneVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates or dissociates virtual private clouds (VPCs) from a zone to set the effective scope of the zone.
//
// Description:
//
// ##### [](#)Precautions:
//
// We recommend that you set the effective scope of a zone after you configure all Domain Name System (DNS) records. If you set an effective scope before you configure DNS records, the DNS resolution for the zone within the effective scope will fail unless you enable the recursive resolution proxy for subdomain names.
//
// @param request - BindZoneVpcRequest
//
// @return BindZoneVpcResponse
func (client *Client) BindZoneVpc(request *BindZoneVpcRequest) (_result *BindZoneVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindZoneVpcResponse{}
	_body, _err := client.BindZoneVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the logical location of a zone.
//
// @param request - ChangeZoneDnsGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeZoneDnsGroupResponse
func (client *Client) ChangeZoneDnsGroupWithOptions(request *ChangeZoneDnsGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeZoneDnsGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DnsGroup) {
		query["DnsGroup"] = request.DnsGroup
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeZoneDnsGroup"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeZoneDnsGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the logical location of a zone.
//
// @param request - ChangeZoneDnsGroupRequest
//
// @return ChangeZoneDnsGroupResponse
func (client *Client) ChangeZoneDnsGroup(request *ChangeZoneDnsGroupRequest) (_result *ChangeZoneDnsGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeZoneDnsGroupResponse{}
	_body, _err := client.ChangeZoneDnsGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a zone name can be added based on a rule.
//
// @param request - CheckZoneNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckZoneNameResponse
func (client *Client) CheckZoneNameWithOptions(request *CheckZoneNameRequest, runtime *dara.RuntimeOptions) (_result *CheckZoneNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckZoneName"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckZoneNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a zone name can be added based on a rule.
//
// @param request - CheckZoneNameRequest
//
// @return CheckZoneNameResponse
func (client *Client) CheckZoneName(request *CheckZoneNameRequest) (_result *CheckZoneNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckZoneNameResponse{}
	_body, _err := client.CheckZoneNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom line.
//
// @param request - DeleteCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomLineResponse
func (client *Client) DeleteCustomLineWithOptions(request *DeleteCustomLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineId) {
		query["LineId"] = request.LineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomLine"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom line.
//
// @param request - DeleteCustomLineRequest
//
// @return DeleteCustomLineResponse
func (client *Client) DeleteCustomLine(request *DeleteCustomLineRequest) (_result *DeleteCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomLineResponse{}
	_body, _err := client.DeleteCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an endpoint based on the endpoint ID.
//
// @param request - DeleteResolverEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResolverEndpointResponse
func (client *Client) DeleteResolverEndpointWithOptions(request *DeleteResolverEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteResolverEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResolverEndpoint"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResolverEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an endpoint based on the endpoint ID.
//
// @param request - DeleteResolverEndpointRequest
//
// @return DeleteResolverEndpointResponse
func (client *Client) DeleteResolverEndpoint(request *DeleteResolverEndpointRequest) (_result *DeleteResolverEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResolverEndpointResponse{}
	_body, _err := client.DeleteResolverEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a forwarding rule based on the rule ID.
//
// @param request - DeleteResolverRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResolverRuleResponse
func (client *Client) DeleteResolverRuleWithOptions(request *DeleteResolverRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteResolverRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResolverRule"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResolverRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a forwarding rule based on the rule ID.
//
// @param request - DeleteResolverRuleRequest
//
// @return DeleteResolverRuleResponse
func (client *Client) DeleteResolverRule(request *DeleteResolverRuleRequest) (_result *DeleteResolverRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResolverRuleResponse{}
	_body, _err := client.DeleteResolverRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an account from the central management of private Domain Name System (DNS) resolution based on the account ID and authorization type.
//
// @param request - DeleteUserVpcAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserVpcAuthorizationResponse
func (client *Client) DeleteUserVpcAuthorizationWithOptions(request *DeleteUserVpcAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserVpcAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.AuthorizedUserId) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserVpcAuthorization"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserVpcAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an account from the central management of private Domain Name System (DNS) resolution based on the account ID and authorization type.
//
// @param request - DeleteUserVpcAuthorizationRequest
//
// @return DeleteUserVpcAuthorizationResponse
func (client *Client) DeleteUserVpcAuthorization(request *DeleteUserVpcAuthorizationRequest) (_result *DeleteUserVpcAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserVpcAuthorizationResponse{}
	_body, _err := client.DeleteUserVpcAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an idle built-in authoritative zone.
//
// Description:
//
// #### [](#)Precautions
//
// If you want to delete a built-in authoritative zone whose effective scope is configured, you must disassociate the zone from the effective scope first.
//
// @param request - DeleteZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteZoneResponse
func (client *Client) DeleteZoneWithOptions(request *DeleteZoneRequest, runtime *dara.RuntimeOptions) (_result *DeleteZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteZone"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an idle built-in authoritative zone.
//
// Description:
//
// #### [](#)Precautions
//
// If you want to delete a built-in authoritative zone whose effective scope is configured, you must disassociate the zone from the effective scope first.
//
// @param request - DeleteZoneRequest
//
// @return DeleteZoneResponse
func (client *Client) DeleteZone(request *DeleteZoneRequest) (_result *DeleteZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteZoneResponse{}
	_body, _err := client.DeleteZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Domain Name System (DNS) record based on the ID of the DNS record.
//
// Description:
//
// #### **Precautions**
//
// Deleted DNS records cannot be restored. Exercise caution when you perform this operation.
//
// @param request - DeleteZoneRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteZoneRecordResponse
func (client *Client) DeleteZoneRecordWithOptions(request *DeleteZoneRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteZoneRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteZoneRecord"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteZoneRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Domain Name System (DNS) record based on the ID of the DNS record.
//
// Description:
//
// #### **Precautions**
//
// Deleted DNS records cannot be restored. Exercise caution when you perform this operation.
//
// @param request - DeleteZoneRecordRequest
//
// @return DeleteZoneRecordResponse
func (client *Client) DeleteZoneRecord(request *DeleteZoneRecordRequest) (_result *DeleteZoneRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteZoneRecordResponse{}
	_body, _err := client.DeleteZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs of Private DNS. Operation logs record operations in modules such as the built-in authoritative module, cache module, forward module, and service address module and record the queries for Domain Name System (DNS) records. You can query operation logs by operation or operation content.
//
// Description:
//
// #### **Precautions**
//
// You can query the operation logs of Private DNS that are generated within the last six months.
//
// @param request - DescribeChangeLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChangeLogsResponse
func (client *Client) DescribeChangeLogsWithOptions(request *DescribeChangeLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChangeLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChangeLogs"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChangeLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs of Private DNS. Operation logs record operations in modules such as the built-in authoritative module, cache module, forward module, and service address module and record the queries for Domain Name System (DNS) records. You can query operation logs by operation or operation content.
//
// Description:
//
// #### **Precautions**
//
// You can query the operation logs of Private DNS that are generated within the last six months.
//
// @param request - DescribeChangeLogsRequest
//
// @return DescribeChangeLogsResponse
func (client *Client) DescribeChangeLogs(request *DescribeChangeLogsRequest) (_result *DescribeChangeLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChangeLogsResponse{}
	_body, _err := client.DescribeChangeLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a custom line.
//
// @param request - DescribeCustomLineInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLineInfoResponse
func (client *Client) DescribeCustomLineInfoWithOptions(request *DescribeCustomLineInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLineInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineId) {
		query["LineId"] = request.LineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomLineInfo"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomLineInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a custom line.
//
// @param request - DescribeCustomLineInfoRequest
//
// @return DescribeCustomLineInfoResponse
func (client *Client) DescribeCustomLineInfo(request *DescribeCustomLineInfoRequest) (_result *DescribeCustomLineInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomLineInfoResponse{}
	_body, _err := client.DescribeCustomLineInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of custom lines.
//
// @param request - DescribeCustomLinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLinesResponse
func (client *Client) DescribeCustomLinesWithOptions(request *DescribeCustomLinesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeCustomLines"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of custom lines.
//
// @param request - DescribeCustomLinesRequest
//
// @return DescribeCustomLinesResponse
func (client *Client) DescribeCustomLines(request *DescribeCustomLinesRequest) (_result *DescribeCustomLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomLinesResponse{}
	_body, _err := client.DescribeCustomLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of regions for selection based on the scenario and virtual private cloud (VPC) type.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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

	if !dara.IsNil(request.AuthorizedUserId) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.VpcType) {
		query["VpcType"] = request.VpcType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of regions for selection based on the scenario and virtual private cloud (VPC) type.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about Domain Name System (DNS) requests based on conditions such as the time range.
//
// @param request - DescribeRequestGraphRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRequestGraphResponse
func (client *Client) DescribeRequestGraphWithOptions(request *DescribeRequestGraphRequest, runtime *dara.RuntimeOptions) (_result *DescribeRequestGraphResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRequestGraph"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRequestGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Domain Name System (DNS) requests based on conditions such as the time range.
//
// @param request - DescribeRequestGraphRequest
//
// @return DescribeRequestGraphResponse
func (client *Client) DescribeRequestGraph(request *DescribeRequestGraphRequest) (_result *DescribeRequestGraphResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRequestGraphResponse{}
	_body, _err := client.DescribeRequestGraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of available zones.
//
// @param request - DescribeResolverAvailableZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverAvailableZonesResponse
func (client *Client) DescribeResolverAvailableZonesWithOptions(request *DescribeResolverAvailableZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverAvailableZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AzId) {
		query["AzId"] = request.AzId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ResolverRegionId) {
		query["ResolverRegionId"] = request.ResolverRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResolverAvailableZones"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResolverAvailableZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of available zones.
//
// @param request - DescribeResolverAvailableZonesRequest
//
// @return DescribeResolverAvailableZonesResponse
func (client *Client) DescribeResolverAvailableZones(request *DescribeResolverAvailableZonesRequest) (_result *DescribeResolverAvailableZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResolverAvailableZonesResponse{}
	_body, _err := client.DescribeResolverAvailableZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an endpoint based on the endpoint ID.
//
// @param request - DescribeResolverEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverEndpointResponse
func (client *Client) DescribeResolverEndpointWithOptions(request *DescribeResolverEndpointRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResolverEndpoint"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResolverEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an endpoint based on the endpoint ID.
//
// @param request - DescribeResolverEndpointRequest
//
// @return DescribeResolverEndpointResponse
func (client *Client) DescribeResolverEndpoint(request *DescribeResolverEndpointRequest) (_result *DescribeResolverEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResolverEndpointResponse{}
	_body, _err := client.DescribeResolverEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of endpoints.
//
// @param request - DescribeResolverEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverEndpointsResponse
func (client *Client) DescribeResolverEndpointsWithOptions(request *DescribeResolverEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VpcRegionId) {
		query["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResolverEndpoints"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResolverEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of endpoints.
//
// @param request - DescribeResolverEndpointsRequest
//
// @return DescribeResolverEndpointsResponse
func (client *Client) DescribeResolverEndpoints(request *DescribeResolverEndpointsRequest) (_result *DescribeResolverEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResolverEndpointsResponse{}
	_body, _err := client.DescribeResolverEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a forwarding rule based on the ID of the forwarding rule.
//
// @param request - DescribeResolverRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverRuleResponse
func (client *Client) DescribeResolverRuleWithOptions(request *DescribeResolverRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResolverRule"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResolverRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a forwarding rule based on the ID of the forwarding rule.
//
// @param request - DescribeResolverRuleRequest
//
// @return DescribeResolverRuleResponse
func (client *Client) DescribeResolverRule(request *DescribeResolverRuleRequest) (_result *DescribeResolverRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResolverRuleResponse{}
	_body, _err := client.DescribeResolverRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of forwarding rules.
//
// @param request - DescribeResolverRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverRulesResponse
func (client *Client) DescribeResolverRulesWithOptions(request *DescribeResolverRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NeedDetailAttributes) {
		query["NeedDetailAttributes"] = request.NeedDetailAttributes
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
		Action:      dara.String("DescribeResolverRules"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResolverRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of forwarding rules.
//
// @param request - DescribeResolverRulesRequest
//
// @return DescribeResolverRulesResponse
func (client *Client) DescribeResolverRules(request *DescribeResolverRulesRequest) (_result *DescribeResolverRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResolverRulesResponse{}
	_body, _err := client.DescribeResolverRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on Domain Name System (DNS) requests received on the previous day, including the top three zones and virtual private clouds (VPCs) with the largest number of DNS requests.
//
// @param request - DescribeStatisticSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStatisticSummaryResponse
func (client *Client) DescribeStatisticSummaryWithOptions(request *DescribeStatisticSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeStatisticSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStatisticSummary"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStatisticSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on Domain Name System (DNS) requests received on the previous day, including the top three zones and virtual private clouds (VPCs) with the largest number of DNS requests.
//
// @param request - DescribeStatisticSummaryRequest
//
// @return DescribeStatisticSummaryResponse
func (client *Client) DescribeStatisticSummary(request *DescribeStatisticSummaryRequest) (_result *DescribeStatisticSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStatisticSummaryResponse{}
	_body, _err := client.DescribeStatisticSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a hostname synchronization task based on a zone ID.
//
// Description:
//
// You can call the DescribeSyncEcsHostTask operation to query the information about a hostname synchronization task based on a zone ID.
//
// @param request - DescribeSyncEcsHostTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSyncEcsHostTaskResponse
func (client *Client) DescribeSyncEcsHostTaskWithOptions(request *DescribeSyncEcsHostTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeSyncEcsHostTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSyncEcsHostTask"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSyncEcsHostTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a hostname synchronization task based on a zone ID.
//
// Description:
//
// You can call the DescribeSyncEcsHostTask operation to query the information about a hostname synchronization task based on a zone ID.
//
// @param request - DescribeSyncEcsHostTaskRequest
//
// @return DescribeSyncEcsHostTaskResponse
func (client *Client) DescribeSyncEcsHostTask(request *DescribeSyncEcsHostTaskRequest) (_result *DescribeSyncEcsHostTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSyncEcsHostTaskResponse{}
	_body, _err := client.DescribeSyncEcsHostTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags added to zones.
//
// Description:
//
// #### **Precautions**
//
// You can call this API operation to query the information about tags added only to zones.
//
// @param request - DescribeTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tags added to zones.
//
// Description:
//
// #### **Precautions**
//
// You can call this API operation to query the information about tags added only to zones.
//
// @param request - DescribeTagsRequest
//
// @return DescribeTagsResponse
func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the current user\\"s service status, such as whether the service is activated, whether there are any unpaid fees, etc.
//
// @param request - DescribeUserServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserServiceStatusResponse
func (client *Client) DescribeUserServiceStatusWithOptions(request *DescribeUserServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserServiceStatus"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the current user\\"s service status, such as whether the service is activated, whether there are any unpaid fees, etc.
//
// @param request - DescribeUserServiceStatusRequest
//
// @return DescribeUserServiceStatusResponse
func (client *Client) DescribeUserServiceStatus(request *DescribeUserServiceStatusRequest) (_result *DescribeUserServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserServiceStatusResponse{}
	_body, _err := client.DescribeUserServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of accounts whose virtual private clouds (VPCs) are associated with a private zone.
//
// @param request - DescribeUserVpcAuthorizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserVpcAuthorizationsResponse
func (client *Client) DescribeUserVpcAuthorizationsWithOptions(request *DescribeUserVpcAuthorizationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserVpcAuthorizationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.AuthorizedUserId) {
		query["AuthorizedUserId"] = request.AuthorizedUserId
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
		Action:      dara.String("DescribeUserVpcAuthorizations"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserVpcAuthorizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of accounts whose virtual private clouds (VPCs) are associated with a private zone.
//
// @param request - DescribeUserVpcAuthorizationsRequest
//
// @return DescribeUserVpcAuthorizationsResponse
func (client *Client) DescribeUserVpcAuthorizations(request *DescribeUserVpcAuthorizationsRequest) (_result *DescribeUserVpcAuthorizationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserVpcAuthorizationsResponse{}
	_body, _err := client.DescribeUserVpcAuthorizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a built-in authoritative zone, such as the virtual private clouds (VPCs) that are associated with the zone.
//
// @param request - DescribeZoneInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneInfoResponse
func (client *Client) DescribeZoneInfoWithOptions(request *DescribeZoneInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeZoneInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZoneInfo"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZoneInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a built-in authoritative zone, such as the virtual private clouds (VPCs) that are associated with the zone.
//
// @param request - DescribeZoneInfoRequest
//
// @return DescribeZoneInfoResponse
func (client *Client) DescribeZoneInfo(request *DescribeZoneInfoRequest) (_result *DescribeZoneInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZoneInfoResponse{}
	_body, _err := client.DescribeZoneInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a Domain Name System (DNS) record.
//
// @param request - DescribeZoneRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneRecordResponse
func (client *Client) DescribeZoneRecordWithOptions(request *DescribeZoneRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeZoneRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZoneRecord"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZoneRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Domain Name System (DNS) record.
//
// @param request - DescribeZoneRecordRequest
//
// @return DescribeZoneRecordResponse
func (client *Client) DescribeZoneRecord(request *DescribeZoneRecordRequest) (_result *DescribeZoneRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZoneRecordResponse{}
	_body, _err := client.DescribeZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Domain Name System (DNS) records.
//
// @param request - DescribeZoneRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneRecordsResponse
func (client *Client) DescribeZoneRecordsWithOptions(request *DescribeZoneRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeZoneRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZoneRecords"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZoneRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Domain Name System (DNS) records.
//
// @param request - DescribeZoneRecordsRequest
//
// @return DescribeZoneRecordsResponse
func (client *Client) DescribeZoneRecords(request *DescribeZoneRecordsRequest) (_result *DescribeZoneRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZoneRecordsResponse{}
	_body, _err := client.DescribeZoneRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of zones within the current account and a list of virtual private clouds (VPCs) associated with the zones.
//
// Description:
//
// We recommend that you do not call this API operation due to its poor performance. Instead, you can call the DescribeZones operation to query a list of zones. If you want to query the information about VPCs with which a zone is associated, you can call the DescribeZoneInfo operation based on the zone ID.
//
// @param request - DescribeZoneVpcTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneVpcTreeResponse
func (client *Client) DescribeZoneVpcTreeWithOptions(request *DescribeZoneVpcTreeRequest, runtime *dara.RuntimeOptions) (_result *DescribeZoneVpcTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZoneVpcTree"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZoneVpcTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of zones within the current account and a list of virtual private clouds (VPCs) associated with the zones.
//
// Description:
//
// We recommend that you do not call this API operation due to its poor performance. Instead, you can call the DescribeZones operation to query a list of zones. If you want to query the information about VPCs with which a zone is associated, you can call the DescribeZoneInfo operation based on the zone ID.
//
// @param request - DescribeZoneVpcTreeRequest
//
// @return DescribeZoneVpcTreeResponse
func (client *Client) DescribeZoneVpcTree(request *DescribeZoneVpcTreeRequest) (_result *DescribeZoneVpcTreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZoneVpcTreeResponse{}
	_body, _err := client.DescribeZoneVpcTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of zones within the current account.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryRegionId) {
		query["QueryRegionId"] = request.QueryRegionId
	}

	if !dara.IsNil(request.QueryVpcId) {
		query["QueryVpcId"] = request.QueryVpcId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceTag) {
		query["ResourceTag"] = request.ResourceTag
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.ZoneTag) {
		query["ZoneTag"] = request.ZoneTag
	}

	if !dara.IsNil(request.ZoneType) {
		query["ZoneType"] = request.ZoneType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of zones within the current account.
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags added to zones.
//
// Description:
//
// #### [](#)**Precautions**
//
// You can call this API operation to query tags added only to zones.
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
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2018-01-01"),
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
// Queries a list of tags added to zones.
//
// Description:
//
// #### [](#)**Precautions**
//
// You can call this API operation to query tags added only to zones.
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
// Changes a resource group.
//
// Description:
//
// #### [](#)Precautions
//
// You can call this API operation to change a resource group only for a zone.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes a resource group.
//
// Description:
//
// #### [](#)Precautions
//
// You can call this API operation to change a resource group only for a zone.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of custom lines.
//
// @param request - SearchCustomLinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCustomLinesResponse
func (client *Client) SearchCustomLinesWithOptions(request *SearchCustomLinesRequest, runtime *dara.RuntimeOptions) (_result *SearchCustomLinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimestampEnd) {
		query["CreateTimestampEnd"] = request.CreateTimestampEnd
	}

	if !dara.IsNil(request.CreateTimestampStart) {
		query["CreateTimestampStart"] = request.CreateTimestampStart
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.Ipv4) {
		query["Ipv4"] = request.Ipv4
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.UpdateTimestampEnd) {
		query["UpdateTimestampEnd"] = request.UpdateTimestampEnd
	}

	if !dara.IsNil(request.UpdateTimestampStart) {
		query["UpdateTimestampStart"] = request.UpdateTimestampStart
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCustomLines"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCustomLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of custom lines.
//
// @param request - SearchCustomLinesRequest
//
// @return SearchCustomLinesResponse
func (client *Client) SearchCustomLines(request *SearchCustomLinesRequest) (_result *SearchCustomLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCustomLinesResponse{}
	_body, _err := client.SearchCustomLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the recursive resolution proxy for subdomain names.
//
// @param request - SetProxyPatternRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetProxyPatternResponse
func (client *Client) SetProxyPatternWithOptions(request *SetProxyPatternRequest, runtime *dara.RuntimeOptions) (_result *SetProxyPatternResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProxyPattern) {
		query["ProxyPattern"] = request.ProxyPattern
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetProxyPattern"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetProxyPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the recursive resolution proxy for subdomain names.
//
// @param request - SetProxyPatternRequest
//
// @return SetProxyPatternResponse
func (client *Client) SetProxyPattern(request *SetProxyPatternRequest) (_result *SetProxyPatternResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetProxyPatternResponse{}
	_body, _err := client.SetProxyPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables a Domain Name System (DNS) record.
//
// @param request - SetZoneRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetZoneRecordStatusResponse
func (client *Client) SetZoneRecordStatusWithOptions(request *SetZoneRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *SetZoneRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetZoneRecordStatus"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetZoneRecordStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a Domain Name System (DNS) record.
//
// @param request - SetZoneRecordStatusRequest
//
// @return SetZoneRecordStatusResponse
func (client *Client) SetZoneRecordStatus(request *SetZoneRecordStatusRequest) (_result *SetZoneRecordStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetZoneRecordStatusResponse{}
	_body, _err := client.SetZoneRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies tags for zones.
//
// Description:
//
// ##### [](#)Precautions
//
// You can configure tags only for zones.
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
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OverWrite) {
		query["OverWrite"] = request.OverWrite
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
		Version:     dara.String("2018-01-01"),
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
// Adds or modifies tags for zones.
//
// Description:
//
// ##### [](#)Precautions
//
// You can configure tags only for zones.
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
// Removes the tags of multiple zones at a time.
//
// Description:
//
// #### [](#)**Precautions**
//
// You can call this API operation to remove tags added only to zones.
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
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Version:     dara.String("2018-01-01"),
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
// Removes the tags of multiple zones at a time.
//
// Description:
//
// #### [](#)**Precautions**
//
// You can call this API operation to remove tags added only to zones.
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
// Modifies a custom line.
//
// @param request - UpdateCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomLineResponse
func (client *Client) UpdateCustomLineWithOptions(request *UpdateCustomLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DnsCategory) {
		query["DnsCategory"] = request.DnsCategory
	}

	if !dara.IsNil(request.Ipv4s) {
		query["Ipv4s"] = request.Ipv4s
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineId) {
		query["LineId"] = request.LineId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomLine"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom line.
//
// @param request - UpdateCustomLineRequest
//
// @return UpdateCustomLineResponse
func (client *Client) UpdateCustomLine(request *UpdateCustomLineRequest) (_result *UpdateCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomLineResponse{}
	_body, _err := client.UpdateCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a Domain Name System (DNS) record based on the record ID.
//
// @param request - UpdateRecordRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordRemarkResponse
func (client *Client) UpdateRecordRemarkWithOptions(request *UpdateRecordRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecordRemark"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecordRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a Domain Name System (DNS) record based on the record ID.
//
// @param request - UpdateRecordRemarkRequest
//
// @return UpdateRecordRemarkResponse
func (client *Client) UpdateRecordRemark(request *UpdateRecordRemarkRequest) (_result *UpdateRecordRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecordRemarkResponse{}
	_body, _err := client.UpdateRecordRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an endpoint.
//
// @param request - UpdateResolverEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResolverEndpointResponse
func (client *Client) UpdateResolverEndpointWithOptions(request *UpdateResolverEndpointRequest, runtime *dara.RuntimeOptions) (_result *UpdateResolverEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.IpConfig) {
		query["IpConfig"] = request.IpConfig
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResolverEndpoint"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResolverEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an endpoint.
//
// @param request - UpdateResolverEndpointRequest
//
// @return UpdateResolverEndpointResponse
func (client *Client) UpdateResolverEndpoint(request *UpdateResolverEndpointRequest) (_result *UpdateResolverEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResolverEndpointResponse{}
	_body, _err := client.UpdateResolverEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a forwarding rule.
//
// @param request - UpdateResolverRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResolverRuleResponse
func (client *Client) UpdateResolverRuleWithOptions(request *UpdateResolverRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateResolverRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.ForwardIp) {
		query["ForwardIp"] = request.ForwardIp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PriorityForwardConfigs) {
		query["PriorityForwardConfigs"] = request.PriorityForwardConfigs
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResolverRule"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResolverRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a forwarding rule.
//
// @param request - UpdateResolverRuleRequest
//
// @return UpdateResolverRuleResponse
func (client *Client) UpdateResolverRule(request *UpdateResolverRuleRequest) (_result *UpdateResolverRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResolverRuleResponse{}
	_body, _err := client.UpdateResolverRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or updates a hostname synchronization task.
//
// @param request - UpdateSyncEcsHostTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSyncEcsHostTaskResponse
func (client *Client) UpdateSyncEcsHostTaskWithOptions(request *UpdateSyncEcsHostTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateSyncEcsHostTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSyncEcsHostTask"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSyncEcsHostTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or updates a hostname synchronization task.
//
// @param request - UpdateSyncEcsHostTaskRequest
//
// @return UpdateSyncEcsHostTaskResponse
func (client *Client) UpdateSyncEcsHostTask(request *UpdateSyncEcsHostTaskRequest) (_result *UpdateSyncEcsHostTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSyncEcsHostTaskResponse{}
	_body, _err := client.UpdateSyncEcsHostTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a Domain Name System (DNS) record of a zone, including the hostname, record value, and weight value of the DNS record.
//
// Description:
//
// #### **Precautions**
//
// The DNS record modification for a zone in the regular module takes effect only after the time to live (TTL) expires. The DNS record modification for a zone in the acceleration module takes effect immediately.
//
// @param request - UpdateZoneRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateZoneRecordResponse
func (client *Client) UpdateZoneRecordWithOptions(request *UpdateZoneRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateZoneRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateZoneRecord"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateZoneRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Domain Name System (DNS) record of a zone, including the hostname, record value, and weight value of the DNS record.
//
// Description:
//
// #### **Precautions**
//
// The DNS record modification for a zone in the regular module takes effect only after the time to live (TTL) expires. The DNS record modification for a zone in the acceleration module takes effect immediately.
//
// @param request - UpdateZoneRecordRequest
//
// @return UpdateZoneRecordResponse
func (client *Client) UpdateZoneRecord(request *UpdateZoneRecordRequest) (_result *UpdateZoneRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateZoneRecordResponse{}
	_body, _err := client.UpdateZoneRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a built-in authoritative zone.
//
// @param request - UpdateZoneRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateZoneRemarkResponse
func (client *Client) UpdateZoneRemarkWithOptions(request *UpdateZoneRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateZoneRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateZoneRemark"),
		Version:     dara.String("2018-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateZoneRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a built-in authoritative zone.
//
// @param request - UpdateZoneRemarkRequest
//
// @return UpdateZoneRemarkResponse
func (client *Client) UpdateZoneRemark(request *UpdateZoneRemarkRequest) (_result *UpdateZoneRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateZoneRemarkResponse{}
	_body, _err := client.UpdateZoneRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
