// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a custom line.
//
// @param request - AddCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomLineResponse
func (client *Client) AddCustomLineWithContext(ctx context.Context, request *AddCustomLineRequest, runtime *dara.RuntimeOptions) (_result *AddCustomLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResolverEndpointResponse
func (client *Client) AddResolverEndpointWithContext(ctx context.Context, request *AddResolverEndpointRequest, runtime *dara.RuntimeOptions) (_result *AddResolverEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddResolverRuleResponse
func (client *Client) AddResolverRuleWithContext(ctx context.Context, request *AddResolverRuleRequest, runtime *dara.RuntimeOptions) (_result *AddResolverRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserVpcAuthorizationResponse
func (client *Client) AddUserVpcAuthorizationWithContext(ctx context.Context, request *AddUserVpcAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *AddUserVpcAuthorizationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneResponse
func (client *Client) AddZoneWithContext(ctx context.Context, request *AddZoneRequest, runtime *dara.RuntimeOptions) (_result *AddZoneResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneRecordResponse
func (client *Client) AddZoneRecordWithContext(ctx context.Context, request *AddZoneRecordRequest, runtime *dara.RuntimeOptions) (_result *AddZoneRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindResolverRuleVpcResponse
func (client *Client) BindResolverRuleVpcWithContext(ctx context.Context, request *BindResolverRuleVpcRequest, runtime *dara.RuntimeOptions) (_result *BindResolverRuleVpcResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindZoneVpcResponse
func (client *Client) BindZoneVpcWithContext(ctx context.Context, request *BindZoneVpcRequest, runtime *dara.RuntimeOptions) (_result *BindZoneVpcResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeZoneDnsGroupResponse
func (client *Client) ChangeZoneDnsGroupWithContext(ctx context.Context, request *ChangeZoneDnsGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeZoneDnsGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckZoneNameResponse
func (client *Client) CheckZoneNameWithContext(ctx context.Context, request *CheckZoneNameRequest, runtime *dara.RuntimeOptions) (_result *CheckZoneNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomLineResponse
func (client *Client) DeleteCustomLineWithContext(ctx context.Context, request *DeleteCustomLineRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResolverEndpointResponse
func (client *Client) DeleteResolverEndpointWithContext(ctx context.Context, request *DeleteResolverEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteResolverEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResolverRuleResponse
func (client *Client) DeleteResolverRuleWithContext(ctx context.Context, request *DeleteResolverRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteResolverRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserVpcAuthorizationResponse
func (client *Client) DeleteUserVpcAuthorizationWithContext(ctx context.Context, request *DeleteUserVpcAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserVpcAuthorizationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteZoneResponse
func (client *Client) DeleteZoneWithContext(ctx context.Context, request *DeleteZoneRequest, runtime *dara.RuntimeOptions) (_result *DeleteZoneResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteZoneRecordResponse
func (client *Client) DeleteZoneRecordWithContext(ctx context.Context, request *DeleteZoneRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteZoneRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChangeLogsResponse
func (client *Client) DescribeChangeLogsWithContext(ctx context.Context, request *DescribeChangeLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChangeLogsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLineInfoResponse
func (client *Client) DescribeCustomLineInfoWithContext(ctx context.Context, request *DescribeCustomLineInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLineInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLinesResponse
func (client *Client) DescribeCustomLinesWithContext(ctx context.Context, request *DescribeCustomLinesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLinesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRequestGraphResponse
func (client *Client) DescribeRequestGraphWithContext(ctx context.Context, request *DescribeRequestGraphRequest, runtime *dara.RuntimeOptions) (_result *DescribeRequestGraphResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverAvailableZonesResponse
func (client *Client) DescribeResolverAvailableZonesWithContext(ctx context.Context, request *DescribeResolverAvailableZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverAvailableZonesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverEndpointResponse
func (client *Client) DescribeResolverEndpointWithContext(ctx context.Context, request *DescribeResolverEndpointRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverEndpointsResponse
func (client *Client) DescribeResolverEndpointsWithContext(ctx context.Context, request *DescribeResolverEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverEndpointsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverRuleResponse
func (client *Client) DescribeResolverRuleWithContext(ctx context.Context, request *DescribeResolverRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResolverRulesResponse
func (client *Client) DescribeResolverRulesWithContext(ctx context.Context, request *DescribeResolverRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeResolverRulesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStatisticSummaryResponse
func (client *Client) DescribeStatisticSummaryWithContext(ctx context.Context, request *DescribeStatisticSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeStatisticSummaryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSyncEcsHostTaskResponse
func (client *Client) DescribeSyncEcsHostTaskWithContext(ctx context.Context, request *DescribeSyncEcsHostTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeSyncEcsHostTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func (client *Client) DescribeTagsWithContext(ctx context.Context, request *DescribeTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserServiceStatusResponse
func (client *Client) DescribeUserServiceStatusWithContext(ctx context.Context, request *DescribeUserServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserServiceStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserVpcAuthorizationsResponse
func (client *Client) DescribeUserVpcAuthorizationsWithContext(ctx context.Context, request *DescribeUserVpcAuthorizationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserVpcAuthorizationsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneInfoResponse
func (client *Client) DescribeZoneInfoWithContext(ctx context.Context, request *DescribeZoneInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeZoneInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneRecordResponse
func (client *Client) DescribeZoneRecordWithContext(ctx context.Context, request *DescribeZoneRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeZoneRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneRecordsResponse
func (client *Client) DescribeZoneRecordsWithContext(ctx context.Context, request *DescribeZoneRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeZoneRecordsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZoneVpcTreeResponse
func (client *Client) DescribeZoneVpcTreeWithContext(ctx context.Context, request *DescribeZoneVpcTreeRequest, runtime *dara.RuntimeOptions) (_result *DescribeZoneVpcTreeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithContext(ctx context.Context, request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCustomLinesResponse
func (client *Client) SearchCustomLinesWithContext(ctx context.Context, request *SearchCustomLinesRequest, runtime *dara.RuntimeOptions) (_result *SearchCustomLinesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetProxyPatternResponse
func (client *Client) SetProxyPatternWithContext(ctx context.Context, request *SetProxyPatternRequest, runtime *dara.RuntimeOptions) (_result *SetProxyPatternResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetZoneRecordStatusResponse
func (client *Client) SetZoneRecordStatusWithContext(ctx context.Context, request *SetZoneRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *SetZoneRecordStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomLineResponse
func (client *Client) UpdateCustomLineWithContext(ctx context.Context, request *UpdateCustomLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomLineResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordRemarkResponse
func (client *Client) UpdateRecordRemarkWithContext(ctx context.Context, request *UpdateRecordRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordRemarkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResolverEndpointResponse
func (client *Client) UpdateResolverEndpointWithContext(ctx context.Context, request *UpdateResolverEndpointRequest, runtime *dara.RuntimeOptions) (_result *UpdateResolverEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResolverRuleResponse
func (client *Client) UpdateResolverRuleWithContext(ctx context.Context, request *UpdateResolverRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateResolverRuleResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSyncEcsHostTaskResponse
func (client *Client) UpdateSyncEcsHostTaskWithContext(ctx context.Context, request *UpdateSyncEcsHostTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateSyncEcsHostTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateZoneRecordResponse
func (client *Client) UpdateZoneRecordWithContext(ctx context.Context, request *UpdateZoneRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateZoneRecordResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateZoneRemarkResponse
func (client *Client) UpdateZoneRemarkWithContext(ctx context.Context, request *UpdateZoneRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateZoneRemarkResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
