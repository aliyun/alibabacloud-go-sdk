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
	client.EndpointMap = map[string]*string{
		"ap-southeast-1": dara.String("cloudfw.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou":    dara.String("cloudfw.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloudfw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates an address book for access control. Supported address book types are IP address books, Elastic Compute Service (ECS) tag-based address books, port address books, and domain address books. An ECS tag-based address book includes the public IP addresses of the ECS instances that have specific tags.
//
// Description:
//
// You can call the AddAddressBook operation to create an address book for access control. The address book can be an IP address book, an ECS tag-based address book, a port address book, or a domain address book.
//
// ## [](#qps)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddAddressBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAddressBookResponse
func (client *Client) AddAddressBookWithOptions(request *AddAddressBookRequest, runtime *dara.RuntimeOptions) (_result *AddAddressBookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AckClusterConnectorId) {
		query["AckClusterConnectorId"] = request.AckClusterConnectorId
	}

	if !dara.IsNil(request.AckLabels) {
		query["AckLabels"] = request.AckLabels
	}

	if !dara.IsNil(request.AckNamespaces) {
		query["AckNamespaces"] = request.AckNamespaces
	}

	if !dara.IsNil(request.AddressList) {
		query["AddressList"] = request.AddressList
	}

	if !dara.IsNil(request.AutoAddTagEcs) {
		query["AutoAddTagEcs"] = request.AutoAddTagEcs
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.TagList) {
		query["TagList"] = request.TagList
	}

	if !dara.IsNil(request.TagRelation) {
		query["TagRelation"] = request.TagRelation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAddressBook"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an address book for access control. Supported address book types are IP address books, Elastic Compute Service (ECS) tag-based address books, port address books, and domain address books. An ECS tag-based address book includes the public IP addresses of the ECS instances that have specific tags.
//
// Description:
//
// You can call the AddAddressBook operation to create an address book for access control. The address book can be an IP address book, an ECS tag-based address book, a port address book, or a domain address book.
//
// ## [](#qps)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddAddressBookRequest
//
// @return AddAddressBookResponse
func (client *Client) AddAddressBook(request *AddAddressBookRequest) (_result *AddAddressBookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAddressBookResponse{}
	_body, _err := client.AddAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access control policy.
//
// Description:
//
// You can call the AddControlPolicy operation to create an access control policy to allow, block, or monitor traffic that reaches Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddControlPolicyResponse
func (client *Client) AddControlPolicyWithOptions(request *AddControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ApplicationNameList) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestPort) {
		query["DestPort"] = request.DestPort
	}

	if !dara.IsNil(request.DestPortGroup) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !dara.IsNil(request.DestPortType) {
		query["DestPortType"] = request.DestPortType
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainResolveType) {
		query["DomainResolveType"] = request.DomainResolveType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewOrder) {
		query["NewOrder"] = request.NewOrder
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatDays) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !dara.IsNil(request.RepeatEndTime) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !dara.IsNil(request.RepeatStartTime) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control policy.
//
// Description:
//
// You can call the AddControlPolicy operation to create an access control policy to allow, block, or monitor traffic that reaches Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddControlPolicyRequest
//
// @return AddControlPolicyResponse
func (client *Client) AddControlPolicy(request *AddControlPolicyRequest) (_result *AddControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddControlPolicyResponse{}
	_body, _err := client.AddControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加DNS防火墙ACL
//
// @param request - AddDnsFirewallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsFirewallPolicyResponse
func (client *Client) AddDnsFirewallPolicyWithOptions(request *AddDnsFirewallPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddDnsFirewallPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsFirewallPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsFirewallPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加DNS防火墙ACL
//
// @param request - AddDnsFirewallPolicyRequest
//
// @return AddDnsFirewallPolicyResponse
func (client *Client) AddDnsFirewallPolicy(request *AddDnsFirewallPolicyRequest) (_result *AddDnsFirewallPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsFirewallPolicyResponse{}
	_body, _err := client.AddDnsFirewallPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds members to Cloud Firewall.
//
// Description:
//
// You can call this operation to add members to Cloud Firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddInstanceMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddInstanceMembersResponse
func (client *Client) AddInstanceMembersWithOptions(request *AddInstanceMembersRequest, runtime *dara.RuntimeOptions) (_result *AddInstanceMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Members) {
		query["Members"] = request.Members
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddInstanceMembers"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds members to Cloud Firewall.
//
// Description:
//
// You can call this operation to add members to Cloud Firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddInstanceMembersRequest
//
// @return AddInstanceMembersResponse
func (client *Client) AddInstanceMembers(request *AddInstanceMembersRequest) (_result *AddInstanceMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddInstanceMembersResponse{}
	_body, _err := client.AddInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加私网DNS域名
//
// @param request - AddPrivateDnsDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPrivateDnsDomainNameResponse
func (client *Client) AddPrivateDnsDomainNameWithOptions(request *AddPrivateDnsDomainNameRequest, runtime *dara.RuntimeOptions) (_result *AddPrivateDnsDomainNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessInstanceId) {
		query["AccessInstanceId"] = request.AccessInstanceId
	}

	if !dara.IsNil(request.DomainNameList) {
		query["DomainNameList"] = request.DomainNameList
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPrivateDnsDomainName"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPrivateDnsDomainNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加私网DNS域名
//
// @param request - AddPrivateDnsDomainNameRequest
//
// @return AddPrivateDnsDomainNameResponse
func (client *Client) AddPrivateDnsDomainName(request *AddPrivateDnsDomainNameRequest) (_result *AddPrivateDnsDomainNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPrivateDnsDomainNameResponse{}
	_body, _err := client.AddPrivateDnsDomainNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Copies all access control policies from a policy group of a source virtual private cloud (VPC) firewall to a policy group of a destination VPC firewall.
//
// Description:
//
// You can call the BatchCopyVpcFirewallControlPolicy operation to copy all access control policies from a policy group of a source VPC firewall to a policy group of a destination VPC firewall.
//
// Before you call this operation, we recommend that you back up access control policies. For more information about how to back up an access control policy, see [Back up an access control policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/back-up-and-roll-back-an-access-control-policy).
//
// After you call this operation, all the access control policies in the policy group of the destination VPC firewall are replaced.
//
// The policy groups of the source VPC firewall and the destination VPC firewall must belong to the same Alibaba Cloud account.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
//
// @param request - BatchCopyVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCopyVpcFirewallControlPolicyResponse
func (client *Client) BatchCopyVpcFirewallControlPolicyWithOptions(request *BatchCopyVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *BatchCopyVpcFirewallControlPolicyResponse, _err error) {
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

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.SourceVpcFirewallId) {
		query["SourceVpcFirewallId"] = request.SourceVpcFirewallId
	}

	if !dara.IsNil(request.TargetVpcFirewallId) {
		query["TargetVpcFirewallId"] = request.TargetVpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCopyVpcFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCopyVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies all access control policies from a policy group of a source virtual private cloud (VPC) firewall to a policy group of a destination VPC firewall.
//
// Description:
//
// You can call the BatchCopyVpcFirewallControlPolicy operation to copy all access control policies from a policy group of a source VPC firewall to a policy group of a destination VPC firewall.
//
// Before you call this operation, we recommend that you back up access control policies. For more information about how to back up an access control policy, see [Back up an access control policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/back-up-and-roll-back-an-access-control-policy).
//
// After you call this operation, all the access control policies in the policy group of the destination VPC firewall are replaced.
//
// The policy groups of the source VPC firewall and the destination VPC firewall must belong to the same Alibaba Cloud account.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
//
// @param request - BatchCopyVpcFirewallControlPolicyRequest
//
// @return BatchCopyVpcFirewallControlPolicyResponse
func (client *Client) BatchCopyVpcFirewallControlPolicy(request *BatchCopyVpcFirewallControlPolicyRequest) (_result *BatchCopyVpcFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCopyVpcFirewallControlPolicyResponse{}
	_body, _err := client.BatchCopyVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple access control policies for a virtual private cloud (VPC) firewall at a time.
//
// @param request - BatchDeleteVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteVpcFirewallControlPolicyResponse
func (client *Client) BatchDeleteVpcFirewallControlPolicyWithOptions(request *BatchDeleteVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteVpcFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuidList) {
		query["AclUuidList"] = request.AclUuidList
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteVpcFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple access control policies for a virtual private cloud (VPC) firewall at a time.
//
// @param request - BatchDeleteVpcFirewallControlPolicyRequest
//
// @return BatchDeleteVpcFirewallControlPolicyResponse
func (client *Client) BatchDeleteVpcFirewallControlPolicy(request *BatchDeleteVpcFirewallControlPolicyRequest) (_result *BatchDeleteVpcFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.BatchDeleteVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建ACK集群连接器
//
// @param request - CreateAckClusterConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAckClusterConnectorResponse
func (client *Client) CreateAckClusterConnectorWithOptions(request *CreateAckClusterConnectorRequest, runtime *dara.RuntimeOptions) (_result *CreateAckClusterConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConnectorName) {
		query["ConnectorName"] = request.ConnectorName
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PrimaryVswitchId) {
		query["PrimaryVswitchId"] = request.PrimaryVswitchId
	}

	if !dara.IsNil(request.PrimaryVswitchIp) {
		query["PrimaryVswitchIp"] = request.PrimaryVswitchIp
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.StandbyVswitchId) {
		query["StandbyVswitchId"] = request.StandbyVswitchId
	}

	if !dara.IsNil(request.StandbyVswitchIp) {
		query["StandbyVswitchIp"] = request.StandbyVswitchIp
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAckClusterConnector"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAckClusterConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建ACK集群连接器
//
// @param request - CreateAckClusterConnectorRequest
//
// @return CreateAckClusterConnectorResponse
func (client *Client) CreateAckClusterConnector(request *CreateAckClusterConnectorRequest) (_result *CreateAckClusterConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAckClusterConnectorResponse{}
	_body, _err := client.CreateAckClusterConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建ACL检查
//
// @param request - CreateAclCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAclCheckResponse
func (client *Client) CreateAclCheckWithOptions(request *CreateAclCheckRequest, runtime *dara.RuntimeOptions) (_result *CreateAclCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.CheckNames) {
		query["CheckNames"] = request.CheckNames
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAclCheck"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAclCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建ACL检查
//
// @param request - CreateAclCheckRequest
//
// @return CreateAclCheckResponse
func (client *Client) CreateAclCheck(request *CreateAclCheckRequest) (_result *CreateAclCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAclCheckResponse{}
	_body, _err := client.CreateAclCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a file download task.
//
// @param request - CreateDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDownloadTaskResponse
func (client *Client) CreateDownloadTaskWithOptions(request *CreateDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDownloadTaskResponse, _err error) {
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

	if !dara.IsNil(request.TaskData) {
		query["TaskData"] = request.TaskData
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TimeZone) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDownloadTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDownloadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file download task.
//
// @param request - CreateDownloadTaskRequest
//
// @return CreateDownloadTaskResponse
func (client *Client) CreateDownloadTask(request *CreateDownloadTaskRequest) (_result *CreateDownloadTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDownloadTaskResponse{}
	_body, _err := client.CreateDownloadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access control policy for a NAT firewall.
//
// Description:
//
// You can call this operation to create a policy that allows, denies, or monitors the traffic that passes through the NAT firewall.
//
// @param request - CreateNatFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNatFirewallControlPolicyResponse
func (client *Client) CreateNatFirewallControlPolicyWithOptions(request *CreateNatFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateNatFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.ApplicationNameList) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestPort) {
		query["DestPort"] = request.DestPort
	}

	if !dara.IsNil(request.DestPortGroup) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !dara.IsNil(request.DestPortType) {
		query["DestPortType"] = request.DestPortType
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainResolveType) {
		query["DomainResolveType"] = request.DomainResolveType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.NewOrder) {
		query["NewOrder"] = request.NewOrder
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatDays) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !dara.IsNil(request.RepeatEndTime) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !dara.IsNil(request.RepeatStartTime) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNatFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNatFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control policy for a NAT firewall.
//
// Description:
//
// You can call this operation to create a policy that allows, denies, or monitors the traffic that passes through the NAT firewall.
//
// @param request - CreateNatFirewallControlPolicyRequest
//
// @return CreateNatFirewallControlPolicyResponse
func (client *Client) CreateNatFirewallControlPolicy(request *CreateNatFirewallControlPolicyRequest) (_result *CreateNatFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNatFirewallControlPolicyResponse{}
	_body, _err := client.CreateNatFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建NAT防火墙预检查
//
// @param request - CreateNatFirewallPreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNatFirewallPreCheckResponse
func (client *Client) CreateNatFirewallPreCheckWithOptions(request *CreateNatFirewallPreCheckRequest, runtime *dara.RuntimeOptions) (_result *CreateNatFirewallPreCheckResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.NatGatewayId) {
		body["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.RegionNo) {
		body["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNatFirewallPreCheck"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNatFirewallPreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建NAT防火墙预检查
//
// @param request - CreateNatFirewallPreCheckRequest
//
// @return CreateNatFirewallPreCheckResponse
func (client *Client) CreateNatFirewallPreCheck(request *CreateNatFirewallPreCheckRequest) (_result *CreateNatFirewallPreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNatFirewallPreCheckResponse{}
	_body, _err := client.CreateNatFirewallPreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建私网DNS终端节点
//
// @param request - CreatePrivateDnsEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateDnsEndpointResponse
func (client *Client) CreatePrivateDnsEndpointWithOptions(request *CreatePrivateDnsEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateDnsEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessInstanceName) {
		query["AccessInstanceName"] = request.AccessInstanceName
	}

	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrimaryDns) {
		query["PrimaryDns"] = request.PrimaryDns
	}

	if !dara.IsNil(request.PrimaryVSwitchId) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !dara.IsNil(request.PrimaryVSwitchIp) {
		query["PrimaryVSwitchIp"] = request.PrimaryVSwitchIp
	}

	if !dara.IsNil(request.PrivateDnsType) {
		query["PrivateDnsType"] = request.PrivateDnsType
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.StandbyDns) {
		query["StandbyDns"] = request.StandbyDns
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !dara.IsNil(request.StandbyVSwitchIp) {
		query["StandbyVSwitchIp"] = request.StandbyVSwitchIp
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateDnsEndpoint"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateDnsEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建私网DNS终端节点
//
// @param request - CreatePrivateDnsEndpointRequest
//
// @return CreatePrivateDnsEndpointResponse
func (client *Client) CreatePrivateDnsEndpoint(request *CreatePrivateDnsEndpointRequest) (_result *CreatePrivateDnsEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePrivateDnsEndpointResponse{}
	_body, _err := client.CreatePrivateDnsEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a NAT firewall.
//
// @param request - CreateSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityProxyResponse
func (client *Client) CreateSecurityProxyWithOptions(request *CreateSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallSwitch) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.NatRouteEntryList) {
		query["NatRouteEntryList"] = request.NatRouteEntryList
	}

	if !dara.IsNil(request.ProxyName) {
		query["ProxyName"] = request.ProxyName
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.StrictMode) {
		query["StrictMode"] = request.StrictMode
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchAuto) {
		query["VswitchAuto"] = request.VswitchAuto
	}

	if !dara.IsNil(request.VswitchCidr) {
		query["VswitchCidr"] = request.VswitchCidr
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecurityProxy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a NAT firewall.
//
// @param request - CreateSecurityProxyRequest
//
// @return CreateSecurityProxyResponse
func (client *Client) CreateSecurityProxy(request *CreateSecurityProxyRequest) (_result *CreateSecurityProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSecurityProxyResponse{}
	_body, _err := client.CreateSecurityProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Cloud Firewall SLS Log Delivery
//
// @param request - CreateSlsLogDispatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSlsLogDispatchResponse
func (client *Client) CreateSlsLogDispatchWithOptions(request *CreateSlsLogDispatchRequest, runtime *dara.RuntimeOptions) (_result *CreateSlsLogDispatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SlsRegionId) {
		body["SlsRegionId"] = request.SlsRegionId
	}

	if !dara.IsNil(request.Ttl) {
		body["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSlsLogDispatch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSlsLogDispatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Cloud Firewall SLS Log Delivery
//
// @param request - CreateSlsLogDispatchRequest
//
// @return CreateSlsLogDispatchResponse
func (client *Client) CreateSlsLogDispatch(request *CreateSlsLogDispatchRequest) (_result *CreateSlsLogDispatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSlsLogDispatchResponse{}
	_body, _err := client.CreateSlsLogDispatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) firewall for a transit router.
//
// @param request - CreateTrFirewallV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrFirewallV2Response
func (client *Client) CreateTrFirewallV2WithOptions(request *CreateTrFirewallV2Request, runtime *dara.RuntimeOptions) (_result *CreateTrFirewallV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.FirewallDescription) {
		query["FirewallDescription"] = request.FirewallDescription
	}

	if !dara.IsNil(request.FirewallName) {
		query["FirewallName"] = request.FirewallName
	}

	if !dara.IsNil(request.FirewallSubnetCidr) {
		query["FirewallSubnetCidr"] = request.FirewallSubnetCidr
	}

	if !dara.IsNil(request.FirewallVpcCidr) {
		query["FirewallVpcCidr"] = request.FirewallVpcCidr
	}

	if !dara.IsNil(request.FirewallVpcId) {
		query["FirewallVpcId"] = request.FirewallVpcId
	}

	if !dara.IsNil(request.FirewallVswitchId) {
		query["FirewallVswitchId"] = request.FirewallVswitchId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.RouteMode) {
		query["RouteMode"] = request.RouteMode
	}

	if !dara.IsNil(request.TrAttachmentMasterCidr) {
		query["TrAttachmentMasterCidr"] = request.TrAttachmentMasterCidr
	}

	if !dara.IsNil(request.TrAttachmentMasterZone) {
		query["TrAttachmentMasterZone"] = request.TrAttachmentMasterZone
	}

	if !dara.IsNil(request.TrAttachmentSlaveCidr) {
		query["TrAttachmentSlaveCidr"] = request.TrAttachmentSlaveCidr
	}

	if !dara.IsNil(request.TrAttachmentSlaveZone) {
		query["TrAttachmentSlaveZone"] = request.TrAttachmentSlaveZone
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrFirewallV2"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrFirewallV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) firewall for a transit router.
//
// @param request - CreateTrFirewallV2Request
//
// @return CreateTrFirewallV2Response
func (client *Client) CreateTrFirewallV2(request *CreateTrFirewallV2Request) (_result *CreateTrFirewallV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTrFirewallV2Response{}
	_body, _err := client.CreateTrFirewallV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a routing policy for a VPC firewall of a transit router.
//
// @param tmpReq - CreateTrFirewallV2RoutePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrFirewallV2RoutePolicyResponse
func (client *Client) CreateTrFirewallV2RoutePolicyWithOptions(tmpReq *CreateTrFirewallV2RoutePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateTrFirewallV2RoutePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTrFirewallV2RoutePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestCandidateList) {
		request.DestCandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestCandidateList, dara.String("DestCandidateList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SrcCandidateList) {
		request.SrcCandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SrcCandidateList, dara.String("SrcCandidateList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DestCandidateListShrink) {
		query["DestCandidateList"] = request.DestCandidateListShrink
	}

	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PolicyDescription) {
		query["PolicyDescription"] = request.PolicyDescription
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.SrcCandidateListShrink) {
		query["SrcCandidateList"] = request.SrcCandidateListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrFirewallV2RoutePolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrFirewallV2RoutePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a routing policy for a VPC firewall of a transit router.
//
// @param request - CreateTrFirewallV2RoutePolicyRequest
//
// @return CreateTrFirewallV2RoutePolicyResponse
func (client *Client) CreateTrFirewallV2RoutePolicy(request *CreateTrFirewallV2RoutePolicyRequest) (_result *CreateTrFirewallV2RoutePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTrFirewallV2RoutePolicyResponse{}
	_body, _err := client.CreateTrFirewallV2RoutePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) firewall to protect traffic between a specified VPC and a network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the CreateVpcFirewallCenConfigure operation to create a VPC firewall. The VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. The VPC firewall cannot protect mutual access traffic between VBRs, between CCN instances, or between VBRs and CCN instances. For more information, see [VPC firewall limits](https://help.aliyun.com/document_detail/172295.html).
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVpcFirewallCenConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallCenConfigureResponse
func (client *Client) CreateVpcFirewallCenConfigureWithOptions(request *CreateVpcFirewallCenConfigureRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallCenConfigureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.FirewallSwitch) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !dara.IsNil(request.FirewallVSwitchCidrBlock) {
		query["FirewallVSwitchCidrBlock"] = request.FirewallVSwitchCidrBlock
	}

	if !dara.IsNil(request.FirewallVpcCidrBlock) {
		query["FirewallVpcCidrBlock"] = request.FirewallVpcCidrBlock
	}

	if !dara.IsNil(request.FirewallVpcStandbyZoneId) {
		query["FirewallVpcStandbyZoneId"] = request.FirewallVpcStandbyZoneId
	}

	if !dara.IsNil(request.FirewallVpcZoneId) {
		query["FirewallVpcZoneId"] = request.FirewallVpcZoneId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.NetworkInstanceId) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcFirewallName) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	if !dara.IsNil(request.VpcRegion) {
		query["VpcRegion"] = request.VpcRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcFirewallCenConfigure"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) firewall to protect traffic between a specified VPC and a network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the CreateVpcFirewallCenConfigure operation to create a VPC firewall. The VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. The VPC firewall cannot protect mutual access traffic between VBRs, between CCN instances, or between VBRs and CCN instances. For more information, see [VPC firewall limits](https://help.aliyun.com/document_detail/172295.html).
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVpcFirewallCenConfigureRequest
//
// @return CreateVpcFirewallCenConfigureResponse
func (client *Client) CreateVpcFirewallCenConfigure(request *CreateVpcFirewallCenConfigureRequest) (_result *CreateVpcFirewallCenConfigureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcFirewallCenConfigureResponse{}
	_body, _err := client.CreateVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建VPC防火墙手动配置
//
// @param request - CreateVpcFirewallCenManualConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallCenManualConfigureResponse
func (client *Client) CreateVpcFirewallCenManualConfigureWithOptions(request *CreateVpcFirewallCenManualConfigureRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallCenManualConfigureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcFirewallName) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcFirewallCenManualConfigure"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcFirewallCenManualConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建VPC防火墙手动配置
//
// @param request - CreateVpcFirewallCenManualConfigureRequest
//
// @return CreateVpcFirewallCenManualConfigureResponse
func (client *Client) CreateVpcFirewallCenManualConfigure(request *CreateVpcFirewallCenManualConfigureRequest) (_result *CreateVpcFirewallCenManualConfigureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcFirewallCenManualConfigureResponse{}
	_body, _err := client.CreateVpcFirewallCenManualConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Virtual Private Cloud (VPC) firewall to protect traffic between two VPCs that are connected by using an Express Connect.
//
// Description:
//
// You can call this operation to create a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. The VPC firewall does not control the mutual access traffic between VPCs that reside in different regions or belong to different Alibaba Cloud accounts. The firewall also does not control the mutual access traffic between VPCs and virtual border routers (VBRs). For more information, see [VPC firewall limits](https://help.aliyun.com/document_detail/172295.html).
//
// ### [](#qps)QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVpcFirewallConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallConfigureResponse
func (client *Client) CreateVpcFirewallConfigureWithOptions(request *CreateVpcFirewallConfigureRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallConfigureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallSwitch) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LocalVpcCidrTableList) {
		query["LocalVpcCidrTableList"] = request.LocalVpcCidrTableList
	}

	if !dara.IsNil(request.LocalVpcId) {
		query["LocalVpcId"] = request.LocalVpcId
	}

	if !dara.IsNil(request.LocalVpcRegion) {
		query["LocalVpcRegion"] = request.LocalVpcRegion
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PeerVpcCidrTableList) {
		query["PeerVpcCidrTableList"] = request.PeerVpcCidrTableList
	}

	if !dara.IsNil(request.PeerVpcId) {
		query["PeerVpcId"] = request.PeerVpcId
	}

	if !dara.IsNil(request.PeerVpcRegion) {
		query["PeerVpcRegion"] = request.PeerVpcRegion
	}

	if !dara.IsNil(request.VpcFirewallName) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcFirewallConfigure"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Virtual Private Cloud (VPC) firewall to protect traffic between two VPCs that are connected by using an Express Connect.
//
// Description:
//
// You can call this operation to create a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. The VPC firewall does not control the mutual access traffic between VPCs that reside in different regions or belong to different Alibaba Cloud accounts. The firewall also does not control the mutual access traffic between VPCs and virtual border routers (VBRs). For more information, see [VPC firewall limits](https://help.aliyun.com/document_detail/172295.html).
//
// ### [](#qps)QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVpcFirewallConfigureRequest
//
// @return CreateVpcFirewallConfigureResponse
func (client *Client) CreateVpcFirewallConfigure(request *CreateVpcFirewallConfigureRequest) (_result *CreateVpcFirewallConfigureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcFirewallConfigureResponse{}
	_body, _err := client.CreateVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access control policy in a specified policy group for a virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call the CreateVpcFirewallControlPolicy operation to create an access control policy in a specified policy group for a VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallControlPolicyResponse
func (client *Client) CreateVpcFirewallControlPolicyWithOptions(request *CreateVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ApplicationNameList) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestPort) {
		query["DestPort"] = request.DestPort
	}

	if !dara.IsNil(request.DestPortGroup) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !dara.IsNil(request.DestPortType) {
		query["DestPortType"] = request.DestPortType
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.DomainResolveType) {
		query["DomainResolveType"] = request.DomainResolveType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.NewOrder) {
		query["NewOrder"] = request.NewOrder
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatDays) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !dara.IsNil(request.RepeatEndTime) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !dara.IsNil(request.RepeatStartTime) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control policy in a specified policy group for a virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call the CreateVpcFirewallControlPolicy operation to create an access control policy in a specified policy group for a VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVpcFirewallControlPolicyRequest
//
// @return CreateVpcFirewallControlPolicyResponse
func (client *Client) CreateVpcFirewallControlPolicy(request *CreateVpcFirewallControlPolicyRequest) (_result *CreateVpcFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcFirewallControlPolicyResponse{}
	_body, _err := client.CreateVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建VPC防火墙开墙前置任务
//
// @param request - CreateVpcFirewallPrecheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallPrecheckResponse
func (client *Client) CreateVpcFirewallPrecheckWithOptions(request *CreateVpcFirewallPrecheckRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallPrecheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.NetworkInstanceType) {
		query["NetworkInstanceType"] = request.NetworkInstanceType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcFirewallPrecheck"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcFirewallPrecheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建VPC防火墙开墙前置任务
//
// @param request - CreateVpcFirewallPrecheckRequest
//
// @return CreateVpcFirewallPrecheckResponse
func (client *Client) CreateVpcFirewallPrecheck(request *CreateVpcFirewallPrecheckRequest) (_result *CreateVpcFirewallPrecheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcFirewallPrecheckResponse{}
	_body, _err := client.CreateVpcFirewallPrecheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建VPC防火墙资产同步任务
//
// @param request - CreateVpcFirewallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallTaskResponse
func (client *Client) CreateVpcFirewallTaskWithOptions(request *CreateVpcFirewallTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.TaskAction) {
		query["TaskAction"] = request.TaskAction
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcFirewallTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcFirewallTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建VPC防火墙资产同步任务
//
// @param request - CreateVpcFirewallTaskRequest
//
// @return CreateVpcFirewallTaskResponse
func (client *Client) CreateVpcFirewallTask(request *CreateVpcFirewallTaskRequest) (_result *CreateVpcFirewallTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcFirewallTaskResponse{}
	_body, _err := client.CreateVpcFirewallTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除ACK集群连接器
//
// @param request - DeleteAckClusterConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAckClusterConnectorResponse
func (client *Client) DeleteAckClusterConnectorWithOptions(request *DeleteAckClusterConnectorRequest, runtime *dara.RuntimeOptions) (_result *DeleteAckClusterConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAckClusterConnector"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAckClusterConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除ACK集群连接器
//
// @param request - DeleteAckClusterConnectorRequest
//
// @return DeleteAckClusterConnectorResponse
func (client *Client) DeleteAckClusterConnector(request *DeleteAckClusterConnectorRequest) (_result *DeleteAckClusterConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAckClusterConnectorResponse{}
	_body, _err := client.DeleteAckClusterConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除ACL备份
//
// @param request - DeleteAclBackupDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAclBackupDataResponse
func (client *Client) DeleteAclBackupDataWithOptions(request *DeleteAclBackupDataRequest, runtime *dara.RuntimeOptions) (_result *DeleteAclBackupDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackUpTime) {
		query["BackUpTime"] = request.BackUpTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAclBackupData"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAclBackupDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除ACL备份
//
// @param request - DeleteAclBackupDataRequest
//
// @return DeleteAclBackupDataResponse
func (client *Client) DeleteAclBackupData(request *DeleteAclBackupDataRequest) (_result *DeleteAclBackupDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAclBackupDataResponse{}
	_body, _err := client.DeleteAclBackupDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address book for access control.
//
// Description:
//
// You can call the DeleteAddressBook operation to delete an address book for access control.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteAddressBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddressBookResponse
func (client *Client) DeleteAddressBookWithOptions(request *DeleteAddressBookRequest, runtime *dara.RuntimeOptions) (_result *DeleteAddressBookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupUuid) {
		query["GroupUuid"] = request.GroupUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAddressBook"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address book for access control.
//
// Description:
//
// You can call the DeleteAddressBook operation to delete an address book for access control.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteAddressBookRequest
//
// @return DeleteAddressBookResponse
func (client *Client) DeleteAddressBook(request *DeleteAddressBookRequest) (_result *DeleteAddressBookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAddressBookResponse{}
	_body, _err := client.DeleteAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access control policy.
//
// Description:
//
// You can call the DeleteControlPolicy operation to delete an access control policy that applies to inbound or outbound traffic.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlPolicyResponse
func (client *Client) DeleteControlPolicyWithOptions(request *DeleteControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control policy.
//
// Description:
//
// You can call the DeleteControlPolicy operation to delete an access control policy that applies to inbound or outbound traffic.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteControlPolicyRequest
//
// @return DeleteControlPolicyResponse
func (client *Client) DeleteControlPolicy(request *DeleteControlPolicyRequest) (_result *DeleteControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.DeleteControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access control policy template.
//
// @param request - DeleteControlPolicyTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlPolicyTemplateResponse
func (client *Client) DeleteControlPolicyTemplateWithOptions(request *DeleteControlPolicyTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteControlPolicyTemplateResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteControlPolicyTemplate"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteControlPolicyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control policy template.
//
// @param request - DeleteControlPolicyTemplateRequest
//
// @return DeleteControlPolicyTemplateResponse
func (client *Client) DeleteControlPolicyTemplate(request *DeleteControlPolicyTemplateRequest) (_result *DeleteControlPolicyTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteControlPolicyTemplateResponse{}
	_body, _err := client.DeleteControlPolicyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除DNS防火墙规则
//
// @param request - DeleteDnsFirewallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnsFirewallPolicyResponse
func (client *Client) DeleteDnsFirewallPolicyWithOptions(request *DeleteDnsFirewallPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnsFirewallPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDnsFirewallPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDnsFirewallPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除DNS防火墙规则
//
// @param request - DeleteDnsFirewallPolicyRequest
//
// @return DeleteDnsFirewallPolicyResponse
func (client *Client) DeleteDnsFirewallPolicy(request *DeleteDnsFirewallPolicyRequest) (_result *DeleteDnsFirewallPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDnsFirewallPolicyResponse{}
	_body, _err := client.DeleteDnsFirewallPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes file download tasks.
//
// Description:
//
// You can call this operation to delete file download tasks and delete the files.
//
// **
//
// **Warning*	- Both tasks and involved files are deleted. You can no longer download the involved files by using the download links. This operation is irreversible. Proceed with caution.
//
// @param request - DeleteDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDownloadTaskResponse
func (client *Client) DeleteDownloadTaskWithOptions(request *DeleteDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteDownloadTaskResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDownloadTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDownloadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes file download tasks.
//
// Description:
//
// You can call this operation to delete file download tasks and delete the files.
//
// **
//
// **Warning*	- Both tasks and involved files are deleted. You can no longer download the involved files by using the download links. This operation is irreversible. Proceed with caution.
//
// @param request - DeleteDownloadTaskRequest
//
// @return DeleteDownloadTaskResponse
func (client *Client) DeleteDownloadTask(request *DeleteDownloadTaskRequest) (_result *DeleteDownloadTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDownloadTaskResponse{}
	_body, _err := client.DeleteDownloadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes routing policies for a virtual private cloud (VPC) firewall of a transit router.
//
// @param request - DeleteFirewallV2RoutePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFirewallV2RoutePoliciesResponse
func (client *Client) DeleteFirewallV2RoutePoliciesWithOptions(request *DeleteFirewallV2RoutePoliciesRequest, runtime *dara.RuntimeOptions) (_result *DeleteFirewallV2RoutePoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TrFirewallRoutePolicyId) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFirewallV2RoutePolicies"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFirewallV2RoutePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes routing policies for a virtual private cloud (VPC) firewall of a transit router.
//
// @param request - DeleteFirewallV2RoutePoliciesRequest
//
// @return DeleteFirewallV2RoutePoliciesResponse
func (client *Client) DeleteFirewallV2RoutePolicies(request *DeleteFirewallV2RoutePoliciesRequest) (_result *DeleteFirewallV2RoutePoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFirewallV2RoutePoliciesResponse{}
	_body, _err := client.DeleteFirewallV2RoutePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes members from Cloud Firewall.
//
// Description:
//
// You can call this operation to remove up to 20 members from Cloud Firewall at a time. Separate multiple members with commas (,). After a member is removed, Cloud Firewall can no longer access the cloud resources of the member. Proceed with caution. Before you call this operation, call the [DescribeInstanceMembers](https://help.aliyun.com/document_detail/271704.html) operation to obtain the information about the members that are added to Cloud Firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteInstanceMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceMembersResponse
func (client *Client) DeleteInstanceMembersWithOptions(request *DeleteInstanceMembersRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberUids) {
		query["MemberUids"] = request.MemberUids
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceMembers"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes members from Cloud Firewall.
//
// Description:
//
// You can call this operation to remove up to 20 members from Cloud Firewall at a time. Separate multiple members with commas (,). After a member is removed, Cloud Firewall can no longer access the cloud resources of the member. Proceed with caution. Before you call this operation, call the [DescribeInstanceMembers](https://help.aliyun.com/document_detail/271704.html) operation to obtain the information about the members that are added to Cloud Firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteInstanceMembersRequest
//
// @return DeleteInstanceMembersResponse
func (client *Client) DeleteInstanceMembers(request *DeleteInstanceMembersRequest) (_result *DeleteInstanceMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceMembersResponse{}
	_body, _err := client.DeleteInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建IPS私网关联信息
//
// @param request - DeleteIpsPrivateAssocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpsPrivateAssocResponse
func (client *Client) DeleteIpsPrivateAssocWithOptions(request *DeleteIpsPrivateAssocRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpsPrivateAssocResponse, _err error) {
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIpsPrivateAssoc"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpsPrivateAssocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建IPS私网关联信息
//
// @param request - DeleteIpsPrivateAssocRequest
//
// @return DeleteIpsPrivateAssocResponse
func (client *Client) DeleteIpsPrivateAssoc(request *DeleteIpsPrivateAssocRequest) (_result *DeleteIpsPrivateAssocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpsPrivateAssocResponse{}
	_body, _err := client.DeleteIpsPrivateAssocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access control policy that is created for a NAT firewall.
//
// Description:
//
// You can use this operation to delete an outbound access control policy that is created for a NAT firewall.
//
// @param request - DeleteNatFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNatFirewallControlPolicyResponse
func (client *Client) DeleteNatFirewallControlPolicyWithOptions(request *DeleteNatFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteNatFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNatFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNatFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control policy that is created for a NAT firewall.
//
// Description:
//
// You can use this operation to delete an outbound access control policy that is created for a NAT firewall.
//
// @param request - DeleteNatFirewallControlPolicyRequest
//
// @return DeleteNatFirewallControlPolicyResponse
func (client *Client) DeleteNatFirewallControlPolicy(request *DeleteNatFirewallControlPolicyRequest) (_result *DeleteNatFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNatFirewallControlPolicyResponse{}
	_body, _err := client.DeleteNatFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes access control policies that are created for a NAT firewall at a time.
//
// @param request - DeleteNatFirewallControlPolicyBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNatFirewallControlPolicyBatchResponse
func (client *Client) DeleteNatFirewallControlPolicyBatchWithOptions(request *DeleteNatFirewallControlPolicyBatchRequest, runtime *dara.RuntimeOptions) (_result *DeleteNatFirewallControlPolicyBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuidList) {
		query["AclUuidList"] = request.AclUuidList
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNatFirewallControlPolicyBatch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNatFirewallControlPolicyBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes access control policies that are created for a NAT firewall at a time.
//
// @param request - DeleteNatFirewallControlPolicyBatchRequest
//
// @return DeleteNatFirewallControlPolicyBatchResponse
func (client *Client) DeleteNatFirewallControlPolicyBatch(request *DeleteNatFirewallControlPolicyBatchRequest) (_result *DeleteNatFirewallControlPolicyBatchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNatFirewallControlPolicyBatchResponse{}
	_body, _err := client.DeleteNatFirewallControlPolicyBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 清空私网DNS域名
//
// @param request - DeletePrivateDnsAllDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateDnsAllDomainNameResponse
func (client *Client) DeletePrivateDnsAllDomainNameWithOptions(request *DeletePrivateDnsAllDomainNameRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateDnsAllDomainNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessInstanceId) {
		query["AccessInstanceId"] = request.AccessInstanceId
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateDnsAllDomainName"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateDnsAllDomainNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清空私网DNS域名
//
// @param request - DeletePrivateDnsAllDomainNameRequest
//
// @return DeletePrivateDnsAllDomainNameResponse
func (client *Client) DeletePrivateDnsAllDomainName(request *DeletePrivateDnsAllDomainNameRequest) (_result *DeletePrivateDnsAllDomainNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrivateDnsAllDomainNameResponse{}
	_body, _err := client.DeletePrivateDnsAllDomainNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除私网DNS域名
//
// @param request - DeletePrivateDnsDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateDnsDomainNameResponse
func (client *Client) DeletePrivateDnsDomainNameWithOptions(request *DeletePrivateDnsDomainNameRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateDnsDomainNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessInstanceId) {
		query["AccessInstanceId"] = request.AccessInstanceId
	}

	if !dara.IsNil(request.DomainNameList) {
		query["DomainNameList"] = request.DomainNameList
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateDnsDomainName"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateDnsDomainNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除私网DNS域名
//
// @param request - DeletePrivateDnsDomainNameRequest
//
// @return DeletePrivateDnsDomainNameResponse
func (client *Client) DeletePrivateDnsDomainName(request *DeletePrivateDnsDomainNameRequest) (_result *DeletePrivateDnsDomainNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrivateDnsDomainNameResponse{}
	_body, _err := client.DeletePrivateDnsDomainNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除私网DNS终端节点
//
// @param request - DeletePrivateDnsEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateDnsEndpointResponse
func (client *Client) DeletePrivateDnsEndpointWithOptions(request *DeletePrivateDnsEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateDnsEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessInstanceId) {
		query["AccessInstanceId"] = request.AccessInstanceId
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateDnsEndpoint"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateDnsEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除私网DNS终端节点
//
// @param request - DeletePrivateDnsEndpointRequest
//
// @return DeletePrivateDnsEndpointResponse
func (client *Client) DeletePrivateDnsEndpoint(request *DeletePrivateDnsEndpointRequest) (_result *DeletePrivateDnsEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrivateDnsEndpointResponse{}
	_body, _err := client.DeletePrivateDnsEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a NAT firewall.
//
// @param request - DeleteSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityProxyResponse
func (client *Client) DeleteSecurityProxyWithOptions(request *DeleteSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityProxyResponse, _err error) {
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

	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityProxy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a NAT firewall.
//
// @param request - DeleteSecurityProxyRequest
//
// @return DeleteSecurityProxyResponse
func (client *Client) DeleteSecurityProxy(request *DeleteSecurityProxyRequest) (_result *DeleteSecurityProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSecurityProxyResponse{}
	_body, _err := client.DeleteSecurityProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virtual private cloud (VPC) firewall that is created for a transit router.
//
// @param request - DeleteTrFirewallV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrFirewallV2Response
func (client *Client) DeleteTrFirewallV2WithOptions(request *DeleteTrFirewallV2Request, runtime *dara.RuntimeOptions) (_result *DeleteTrFirewallV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrFirewallV2"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrFirewallV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual private cloud (VPC) firewall that is created for a transit router.
//
// @param request - DeleteTrFirewallV2Request
//
// @return DeleteTrFirewallV2Response
func (client *Client) DeleteTrFirewallV2(request *DeleteTrFirewallV2Request) (_result *DeleteTrFirewallV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTrFirewallV2Response{}
	_body, _err := client.DeleteTrFirewallV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virtual private cloud (VPC) firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the DeleteVpcFirewallCenConfigure operation to delete a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVpcFirewallCenConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcFirewallCenConfigureResponse
func (client *Client) DeleteVpcFirewallCenConfigureWithOptions(request *DeleteVpcFirewallCenConfigureRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcFirewallCenConfigureResponse, _err error) {
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

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VpcFirewallIdList) {
		query["VpcFirewallIdList"] = request.VpcFirewallIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpcFirewallCenConfigure"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual private cloud (VPC) firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the DeleteVpcFirewallCenConfigure operation to delete a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVpcFirewallCenConfigureRequest
//
// @return DeleteVpcFirewallCenConfigureResponse
func (client *Client) DeleteVpcFirewallCenConfigure(request *DeleteVpcFirewallCenConfigureRequest) (_result *DeleteVpcFirewallCenConfigureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVpcFirewallCenConfigureResponse{}
	_body, _err := client.DeleteVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virtual private cloud (VPC) firewall that controls traffic between two VPCs. The VPCs are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the DeleteVpcFirewallConfigure operation to delete a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVpcFirewallConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcFirewallConfigureResponse
func (client *Client) DeleteVpcFirewallConfigureWithOptions(request *DeleteVpcFirewallConfigureRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcFirewallConfigureResponse, _err error) {
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

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VpcFirewallIdList) {
		query["VpcFirewallIdList"] = request.VpcFirewallIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpcFirewallConfigure"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual private cloud (VPC) firewall that controls traffic between two VPCs. The VPCs are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the DeleteVpcFirewallConfigure operation to delete a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVpcFirewallConfigureRequest
//
// @return DeleteVpcFirewallConfigureResponse
func (client *Client) DeleteVpcFirewallConfigure(request *DeleteVpcFirewallConfigureRequest) (_result *DeleteVpcFirewallConfigureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVpcFirewallConfigureResponse{}
	_body, _err := client.DeleteVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access control policy from a specific policy group for a virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call the DeleteVpcFirewallControlPolicy operation to delete an access control policy from a specific policy group for a VPC firewall. Different access control policies are used for the VPC firewall that is used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewall that is used to protect each Express Connect circuit.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcFirewallControlPolicyResponse
func (client *Client) DeleteVpcFirewallControlPolicyWithOptions(request *DeleteVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpcFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control policy from a specific policy group for a virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call the DeleteVpcFirewallControlPolicy operation to delete an access control policy from a specific policy group for a VPC firewall. Different access control policies are used for the VPC firewall that is used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewall that is used to protect each Express Connect circuit.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVpcFirewallControlPolicyRequest
//
// @return DeleteVpcFirewallControlPolicyResponse
func (client *Client) DeleteVpcFirewallControlPolicy(request *DeleteVpcFirewallControlPolicyRequest) (_result *DeleteVpcFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.DeleteVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on the requests that are blocked by the access control list (ACL) feature.
//
// @param request - DescribeACLProtectTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeACLProtectTrendResponse
func (client *Client) DescribeACLProtectTrendWithOptions(request *DescribeACLProtectTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeACLProtectTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeACLProtectTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeACLProtectTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the requests that are blocked by the access control list (ACL) feature.
//
// @param request - DescribeACLProtectTrendRequest
//
// @return DescribeACLProtectTrendResponse
func (client *Client) DescribeACLProtectTrend(request *DescribeACLProtectTrendRequest) (_result *DescribeACLProtectTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeACLProtectTrendResponse{}
	_body, _err := client.DescribeACLProtectTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询AI流量分析开启状态
//
// @param request - DescribeAITrafficAnalysisStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAITrafficAnalysisStatusResponse
func (client *Client) DescribeAITrafficAnalysisStatusWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeAITrafficAnalysisStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAITrafficAnalysisStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAITrafficAnalysisStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AI流量分析开启状态
//
// @return DescribeAITrafficAnalysisStatusResponse
func (client *Client) DescribeAITrafficAnalysisStatus() (_result *DescribeAITrafficAnalysisStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAITrafficAnalysisStatusResponse{}
	_body, _err := client.DescribeAITrafficAnalysisStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询接入实例地域列表
//
// @param request - DescribeAccessInstanceRegionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceRegionListResponse
func (client *Client) DescribeAccessInstanceRegionListWithOptions(request *DescribeAccessInstanceRegionListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceRegionListResponse, _err error) {
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
		Action:      dara.String("DescribeAccessInstanceRegionList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessInstanceRegionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入实例地域列表
//
// @param request - DescribeAccessInstanceRegionListRequest
//
// @return DescribeAccessInstanceRegionListResponse
func (client *Client) DescribeAccessInstanceRegionList(request *DescribeAccessInstanceRegionListRequest) (_result *DescribeAccessInstanceRegionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessInstanceRegionListResponse{}
	_body, _err := client.DescribeAccessInstanceRegionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询接入实例任务
//
// @param request - DescribeAccessInstanceTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceTaskResponse
func (client *Client) DescribeAccessInstanceTaskWithOptions(request *DescribeAccessInstanceTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceTaskResponse, _err error) {
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
		Action:      dara.String("DescribeAccessInstanceTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessInstanceTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入实例任务
//
// @param request - DescribeAccessInstanceTaskRequest
//
// @return DescribeAccessInstanceTaskResponse
func (client *Client) DescribeAccessInstanceTask(request *DescribeAccessInstanceTaskRequest) (_result *DescribeAccessInstanceTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessInstanceTaskResponse{}
	_body, _err := client.DescribeAccessInstanceTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询接入实例的交换机列表
//
// @param request - DescribeAccessInstanceVSwitchListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceVSwitchListResponse
func (client *Client) DescribeAccessInstanceVSwitchListWithOptions(request *DescribeAccessInstanceVSwitchListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceVSwitchListResponse, _err error) {
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
		Action:      dara.String("DescribeAccessInstanceVSwitchList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessInstanceVSwitchListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入实例的交换机列表
//
// @param request - DescribeAccessInstanceVSwitchListRequest
//
// @return DescribeAccessInstanceVSwitchListResponse
func (client *Client) DescribeAccessInstanceVSwitchList(request *DescribeAccessInstanceVSwitchListRequest) (_result *DescribeAccessInstanceVSwitchListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessInstanceVSwitchListResponse{}
	_body, _err := client.DescribeAccessInstanceVSwitchListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询接入实例VPC列表
//
// @param request - DescribeAccessInstanceVpcListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceVpcListResponse
func (client *Client) DescribeAccessInstanceVpcListWithOptions(request *DescribeAccessInstanceVpcListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceVpcListResponse, _err error) {
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
		Action:      dara.String("DescribeAccessInstanceVpcList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessInstanceVpcListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入实例VPC列表
//
// @param request - DescribeAccessInstanceVpcListRequest
//
// @return DescribeAccessInstanceVpcListResponse
func (client *Client) DescribeAccessInstanceVpcList(request *DescribeAccessInstanceVpcListRequest) (_result *DescribeAccessInstanceVpcListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessInstanceVpcListResponse{}
	_body, _err := client.DescribeAccessInstanceVpcListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询接入实例支持的可用区列表
//
// @param request - DescribeAccessInstanceZoneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceZoneListResponse
func (client *Client) DescribeAccessInstanceZoneListWithOptions(request *DescribeAccessInstanceZoneListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceZoneListResponse, _err error) {
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
		Action:      dara.String("DescribeAccessInstanceZoneList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessInstanceZoneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入实例支持的可用区列表
//
// @param request - DescribeAccessInstanceZoneListRequest
//
// @return DescribeAccessInstanceZoneListResponse
func (client *Client) DescribeAccessInstanceZoneList(request *DescribeAccessInstanceZoneListRequest) (_result *DescribeAccessInstanceZoneListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessInstanceZoneListResponse{}
	_body, _err := client.DescribeAccessInstanceZoneListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定ACK集群连接器
//
// @param request - DescribeAckClusterConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClusterConnectorResponse
func (client *Client) DescribeAckClusterConnectorWithOptions(request *DescribeAckClusterConnectorRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClusterConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAckClusterConnector"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAckClusterConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定ACK集群连接器
//
// @param request - DescribeAckClusterConnectorRequest
//
// @return DescribeAckClusterConnectorResponse
func (client *Client) DescribeAckClusterConnector(request *DescribeAckClusterConnectorRequest) (_result *DescribeAckClusterConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAckClusterConnectorResponse{}
	_body, _err := client.DescribeAckClusterConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询ACK集群连接器列表
//
// @param request - DescribeAckClusterConnectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClusterConnectorsResponse
func (client *Client) DescribeAckClusterConnectorsWithOptions(request *DescribeAckClusterConnectorsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClusterConnectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConnectorName) {
		query["ConnectorName"] = request.ConnectorName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAckClusterConnectors"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAckClusterConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询ACK集群连接器列表
//
// @param request - DescribeAckClusterConnectorsRequest
//
// @return DescribeAckClusterConnectorsResponse
func (client *Client) DescribeAckClusterConnectors(request *DescribeAckClusterConnectorsRequest) (_result *DescribeAckClusterConnectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAckClusterConnectorsResponse{}
	_body, _err := client.DescribeAckClusterConnectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询阿里云K8S容器服务（ACK）集群命名空间
//
// @param request - DescribeAckClusterNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClusterNamespacesResponse
func (client *Client) DescribeAckClusterNamespacesWithOptions(request *DescribeAckClusterNamespacesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClusterNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAckClusterNamespaces"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAckClusterNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询阿里云K8S容器服务（ACK）集群命名空间
//
// @param request - DescribeAckClusterNamespacesRequest
//
// @return DescribeAckClusterNamespacesResponse
func (client *Client) DescribeAckClusterNamespaces(request *DescribeAckClusterNamespacesRequest) (_result *DescribeAckClusterNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAckClusterNamespacesResponse{}
	_body, _err := client.DescribeAckClusterNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询阿里云K8S容器服务（ACK）集群标签
//
// @param request - DescribeAckClusterPodLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClusterPodLabelsResponse
func (client *Client) DescribeAckClusterPodLabelsWithOptions(request *DescribeAckClusterPodLabelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClusterPodLabelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAckClusterPodLabels"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAckClusterPodLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询阿里云K8S容器服务（ACK）集群标签
//
// @param request - DescribeAckClusterPodLabelsRequest
//
// @return DescribeAckClusterPodLabelsResponse
func (client *Client) DescribeAckClusterPodLabels(request *DescribeAckClusterPodLabelsRequest) (_result *DescribeAckClusterPodLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAckClusterPodLabelsResponse{}
	_body, _err := client.DescribeAckClusterPodLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询阿里云K8S容器服务（ACK）集群，查询符合条件的ACK集群（例如指定集群类型、集群规格）列表信息
//
// @param request - DescribeAckClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClustersResponse
func (client *Client) DescribeAckClustersWithOptions(request *DescribeAckClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterSpec) {
		query["ClusterSpec"] = request.ClusterSpec
	}

	if !dara.IsNil(request.ConnectorStatus) {
		query["ConnectorStatus"] = request.ConnectorStatus
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAckClusters"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAckClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询阿里云K8S容器服务（ACK）集群，查询符合条件的ACK集群（例如指定集群类型、集群规格）列表信息
//
// @param request - DescribeAckClustersRequest
//
// @return DescribeAckClustersResponse
func (client *Client) DescribeAckClusters(request *DescribeAckClustersRequest) (_result *DescribeAckClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAckClustersResponse{}
	_body, _err := client.DescribeAckClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询访问控制应用
//
// @param request - DescribeAclAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclAppsResponse
func (client *Client) DescribeAclAppsWithOptions(request *DescribeAclAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclAppsResponse, _err error) {
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
		Action:      dara.String("DescribeAclApps"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询访问控制应用
//
// @param request - DescribeAclAppsRequest
//
// @return DescribeAclAppsResponse
func (client *Client) DescribeAclApps(request *DescribeAclAppsRequest) (_result *DescribeAclAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAclAppsResponse{}
	_body, _err := client.DescribeAclAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询ACL检查详情
//
// @param request - DescribeAclCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclCheckResponse
func (client *Client) DescribeAclCheckWithOptions(request *DescribeAclCheckRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclCheckResponse, _err error) {
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
		Action:      dara.String("DescribeAclCheck"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询ACL检查详情
//
// @param request - DescribeAclCheckRequest
//
// @return DescribeAclCheckResponse
func (client *Client) DescribeAclCheck(request *DescribeAclCheckRequest) (_result *DescribeAclCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAclCheckResponse{}
	_body, _err := client.DescribeAclCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询NAT防火墙预检查结果
//
// @param request - DescribeAclCheckQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclCheckQuotaResponse
func (client *Client) DescribeAclCheckQuotaWithOptions(request *DescribeAclCheckQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclCheckQuotaResponse, _err error) {
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
		Action:      dara.String("DescribeAclCheckQuota"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclCheckQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询NAT防火墙预检查结果
//
// @param request - DescribeAclCheckQuotaRequest
//
// @return DescribeAclCheckQuotaResponse
func (client *Client) DescribeAclCheckQuota(request *DescribeAclCheckQuotaRequest) (_result *DescribeAclCheckQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAclCheckQuotaResponse{}
	_body, _err := client.DescribeAclCheckQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询ACL检查条目
//
// @param request - DescribeAclChecksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclChecksResponse
func (client *Client) DescribeAclChecksWithOptions(request *DescribeAclChecksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclChecksResponse, _err error) {
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
		Action:      dara.String("DescribeAclChecks"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclChecksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询ACL检查条目
//
// @param request - DescribeAclChecksRequest
//
// @return DescribeAclChecksResponse
func (client *Client) DescribeAclChecks(request *DescribeAclChecksRequest) (_result *DescribeAclChecksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAclChecksResponse{}
	_body, _err := client.DescribeAclChecksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取总ACL配置数
//
// @param request - DescribeAclRuleCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclRuleCountResponse
func (client *Client) DescribeAclRuleCountWithOptions(request *DescribeAclRuleCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclRuleCountResponse, _err error) {
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

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAclRuleCount"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclRuleCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取总ACL配置数
//
// @param request - DescribeAclRuleCountRequest
//
// @return DescribeAclRuleCountResponse
func (client *Client) DescribeAclRuleCount(request *DescribeAclRuleCountRequest) (_result *DescribeAclRuleCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAclRuleCountResponse{}
	_body, _err := client.DescribeAclRuleCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取ACL白名单
//
// @param request - DescribeAclWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclWhitelistResponse
func (client *Client) DescribeAclWhitelistWithOptions(request *DescribeAclWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAclWhitelist"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ACL白名单
//
// @param request - DescribeAclWhitelistRequest
//
// @return DescribeAclWhitelistResponse
func (client *Client) DescribeAclWhitelist(request *DescribeAclWhitelistRequest) (_result *DescribeAclWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAclWhitelistResponse{}
	_body, _err := client.DescribeAclWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about an address book for an access control policy.
//
// Description:
//
// You can call this operation to query the details about an address book for an access control policy.
//
// ## [](#qps)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeAddressBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddressBookResponse
func (client *Client) DescribeAddressBookWithOptions(request *DescribeAddressBookRequest, runtime *dara.RuntimeOptions) (_result *DescribeAddressBookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainPort) {
		query["ContainPort"] = request.ContainPort
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAddressBook"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about an address book for an access control policy.
//
// Description:
//
// You can call this operation to query the details about an address book for an access control policy.
//
// ## [](#qps)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeAddressBookRequest
//
// @return DescribeAddressBookResponse
func (client *Client) DescribeAddressBook(request *DescribeAddressBookRequest) (_result *DescribeAddressBookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAddressBookResponse{}
	_body, _err := client.DescribeAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the assets that are protected by Cloud Firewall.
//
// Description:
//
// You can call the DescribeAssetList operation to query the assets that are protected by Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeAssetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetListResponse
func (client *Client) DescribeAssetListWithOptions(request *DescribeAssetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.NewResourceTag) {
		query["NewResourceTag"] = request.NewResourceTag
	}

	if !dara.IsNil(request.OutStatistic) {
		query["OutStatistic"] = request.OutStatistic
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SearchItem) {
		query["SearchItem"] = request.SearchItem
	}

	if !dara.IsNil(request.SensitiveStatus) {
		query["SensitiveStatus"] = request.SensitiveStatus
	}

	if !dara.IsNil(request.SgStatus) {
		query["SgStatus"] = request.SgStatus
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAssetList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAssetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the assets that are protected by Cloud Firewall.
//
// Description:
//
// You can call the DescribeAssetList operation to query the assets that are protected by Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeAssetListRequest
//
// @return DescribeAssetListResponse
func (client *Client) DescribeAssetList(request *DescribeAssetListRequest) (_result *DescribeAssetListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAssetListResponse{}
	_body, _err := client.DescribeAssetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the risk levels of assets.
//
// @param request - DescribeAssetRiskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetRiskListResponse
func (client *Client) DescribeAssetRiskListWithOptions(request *DescribeAssetRiskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetRiskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpAddrList) {
		query["IpAddrList"] = request.IpAddrList
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAssetRiskList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAssetRiskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the risk levels of assets.
//
// @param request - DescribeAssetRiskListRequest
//
// @return DescribeAssetRiskListResponse
func (client *Client) DescribeAssetRiskList(request *DescribeAssetRiskListRequest) (_result *DescribeAssetRiskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAssetRiskListResponse{}
	_body, _err := client.DescribeAssetRiskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries statistics on the assets that are protected by Cloud Firewall.
//
// @param request - DescribeAssetStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetStatisticResponse
func (client *Client) DescribeAssetStatisticWithOptions(request *DescribeAssetStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetStatisticResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAssetStatistic"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAssetStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics on the assets that are protected by Cloud Firewall.
//
// @param request - DescribeAssetStatisticRequest
//
// @return DescribeAssetStatisticResponse
func (client *Client) DescribeAssetStatistic(request *DescribeAssetStatisticRequest) (_result *DescribeAssetStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAssetStatisticResponse{}
	_body, _err := client.DescribeAssetStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the firewall risk level.
//
// @param request - DescribeCfwRiskLevelSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCfwRiskLevelSummaryResponse
func (client *Client) DescribeCfwRiskLevelSummaryWithOptions(request *DescribeCfwRiskLevelSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeCfwRiskLevelSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCfwRiskLevelSummary"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCfwRiskLevelSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the firewall risk level.
//
// @param request - DescribeCfwRiskLevelSummaryRequest
//
// @return DescribeCfwRiskLevelSummaryResponse
func (client *Client) DescribeCfwRiskLevelSummary(request *DescribeCfwRiskLevelSummaryRequest) (_result *DescribeCfwRiskLevelSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCfwRiskLevelSummaryResponse{}
	_body, _err := client.DescribeCfwRiskLevelSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取清空授权信息
//
// @param request - DescribeClearAuthInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClearAuthInfoResponse
func (client *Client) DescribeClearAuthInfoWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeClearAuthInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClearAuthInfo"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClearAuthInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取清空授权信息
//
// @return DescribeClearAuthInfoResponse
func (client *Client) DescribeClearAuthInfo() (_result *DescribeClearAuthInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClearAuthInfoResponse{}
	_body, _err := client.DescribeClearAuthInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about all access control policies.
//
// Description:
//
// You can call the DescribeControlPolicy operation to query the details about access control policies by page.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeControlPolicyResponse
func (client *Client) DescribeControlPolicyWithOptions(request *DescribeControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about all access control policies.
//
// Description:
//
// You can call the DescribeControlPolicy operation to query the details about access control policies by page.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeControlPolicyRequest
//
// @return DescribeControlPolicyResponse
func (client *Client) DescribeControlPolicy(request *DescribeControlPolicyRequest) (_result *DescribeControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeControlPolicyResponse{}
	_body, _err := client.DescribeControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取成员账号列表
//
// @param request - DescribeCtrlInstanceMemberAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCtrlInstanceMemberAccountsResponse
func (client *Client) DescribeCtrlInstanceMemberAccountsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeCtrlInstanceMemberAccountsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCtrlInstanceMemberAccounts"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCtrlInstanceMemberAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取成员账号列表
//
// @return DescribeCtrlInstanceMemberAccountsResponse
func (client *Client) DescribeCtrlInstanceMemberAccounts() (_result *DescribeCtrlInstanceMemberAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCtrlInstanceMemberAccountsResponse{}
	_body, _err := client.DescribeCtrlInstanceMemberAccountsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default intrusion prevention system (IPS) configurations.
//
// @param request - DescribeDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefaultIPSConfigResponse
func (client *Client) DescribeDefaultIPSConfigWithOptions(request *DescribeDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefaultIPSConfigResponse, _err error) {
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
		Action:      dara.String("DescribeDefaultIPSConfig"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default intrusion prevention system (IPS) configurations.
//
// @param request - DescribeDefaultIPSConfigRequest
//
// @return DescribeDefaultIPSConfigResponse
func (client *Client) DescribeDefaultIPSConfig(request *DescribeDefaultIPSConfigRequest) (_result *DescribeDefaultIPSConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefaultIPSConfigResponse{}
	_body, _err := client.DescribeDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取DNS防火墙ACL列表
//
// @param request - DescribeDnsFirewallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsFirewallPolicyResponse
func (client *Client) DescribeDnsFirewallPolicyWithOptions(request *DescribeDnsFirewallPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsFirewallPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsFirewallPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsFirewallPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取DNS防火墙ACL列表
//
// @param request - DescribeDnsFirewallPolicyRequest
//
// @return DescribeDnsFirewallPolicyResponse
func (client *Client) DescribeDnsFirewallPolicy(request *DescribeDnsFirewallPolicyRequest) (_result *DescribeDnsFirewallPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsFirewallPolicyResponse{}
	_body, _err := client.DescribeDnsFirewallPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeDomainResolve is deprecated
//
// Summary:
//
// Queries Domain Name System (DNS) records.
//
// Description:
//
// You can use this operation to query the DNS record of a domain name. This operation can retrieve DNS records only from Alibaba Cloud DNS. Before you can call this operation, make sure that your domain name is hosted on Alibaba Cloud DNS.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDomainResolveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainResolveResponse
func (client *Client) DescribeDomainResolveWithOptions(request *DescribeDomainResolveRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainResolveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainResolve"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainResolveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeDomainResolve is deprecated
//
// Summary:
//
// Queries Domain Name System (DNS) records.
//
// Description:
//
// You can use this operation to query the DNS record of a domain name. This operation can retrieve DNS records only from Alibaba Cloud DNS. Before you can call this operation, make sure that your domain name is hosted on Alibaba Cloud DNS.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDomainResolveRequest
//
// @return DescribeDomainResolveResponse
// Deprecated
func (client *Client) DescribeDomainResolve(request *DescribeDomainResolveRequest) (_result *DescribeDomainResolveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainResolveResponse{}
	_body, _err := client.DescribeDomainResolveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries file download tasks, including the task information and download URLs.
//
// @param request - DescribeDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadTaskResponse
func (client *Client) DescribeDownloadTaskWithOptions(request *DescribeDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDownloadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries file download tasks, including the task information and download URLs.
//
// @param request - DescribeDownloadTaskRequest
//
// @return DescribeDownloadTaskResponse
func (client *Client) DescribeDownloadTask(request *DescribeDownloadTaskRequest) (_result *DescribeDownloadTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDownloadTaskResponse{}
	_body, _err := client.DescribeDownloadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the types of download tasks. The type corresponds to the TaskType fields in the download task-related operations.
//
// @param request - DescribeDownloadTaskTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadTaskTypeResponse
func (client *Client) DescribeDownloadTaskTypeWithOptions(request *DescribeDownloadTaskTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadTaskTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadTaskType"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDownloadTaskTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types of download tasks. The type corresponds to the TaskType fields in the download task-related operations.
//
// @param request - DescribeDownloadTaskTypeRequest
//
// @return DescribeDownloadTaskTypeResponse
func (client *Client) DescribeDownloadTaskType(request *DescribeDownloadTaskTypeRequest) (_result *DescribeDownloadTaskTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDownloadTaskTypeResponse{}
	_body, _err := client.DescribeDownloadTaskTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取防火墙DROP数据统计
//
// @param request - DescribeFirewallDropStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallDropStatisticsResponse
func (client *Client) DescribeFirewallDropStatisticsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeFirewallDropStatisticsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFirewallDropStatistics"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFirewallDropStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取防火墙DROP数据统计
//
// @return DescribeFirewallDropStatisticsResponse
func (client *Client) DescribeFirewallDropStatistics() (_result *DescribeFirewallDropStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFirewallDropStatisticsResponse{}
	_body, _err := client.DescribeFirewallDropStatisticsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取防火墙任务
//
// @param request - DescribeFirewallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallTaskResponse
func (client *Client) DescribeFirewallTaskWithOptions(request *DescribeFirewallTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeFirewallTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeFirewallTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFirewallTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取防火墙任务
//
// @param request - DescribeFirewallTaskRequest
//
// @return DescribeFirewallTaskResponse
func (client *Client) DescribeFirewallTask(request *DescribeFirewallTaskRequest) (_result *DescribeFirewallTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFirewallTaskResponse{}
	_body, _err := client.DescribeFirewallTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取防火墙创建的交换机
//
// @param request - DescribeFirewallVSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallVSwitchResponse
func (client *Client) DescribeFirewallVSwitchWithOptions(request *DescribeFirewallVSwitchRequest, runtime *dara.RuntimeOptions) (_result *DescribeFirewallVSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFirewallVSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFirewallVSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取防火墙创建的交换机
//
// @param request - DescribeFirewallVSwitchRequest
//
// @return DescribeFirewallVSwitchResponse
func (client *Client) DescribeFirewallVSwitch(request *DescribeFirewallVSwitchRequest) (_result *DescribeFirewallVSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFirewallVSwitchResponse{}
	_body, _err := client.DescribeFirewallVSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about members in Cloud Firewall.
//
// Description:
//
// You can use this operation to query the information about members in Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceMembersResponse
func (client *Client) DescribeInstanceMembersWithOptions(request *DescribeInstanceMembersRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.MemberDesc) {
		query["MemberDesc"] = request.MemberDesc
	}

	if !dara.IsNil(request.MemberDisplayName) {
		query["MemberDisplayName"] = request.MemberDisplayName
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceMembers"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about members in Cloud Firewall.
//
// Description:
//
// You can use this operation to query the information about members in Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceMembersRequest
//
// @return DescribeInstanceMembersResponse
func (client *Client) DescribeInstanceMembers(request *DescribeInstanceMembersRequest) (_result *DescribeInstanceMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceMembersResponse{}
	_body, _err := client.DescribeInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the risk levels of instances.
//
// @param request - DescribeInstanceRiskLevelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceRiskLevelsResponse
func (client *Client) DescribeInstanceRiskLevelsWithOptions(request *DescribeInstanceRiskLevelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceRiskLevelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceRiskLevels"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceRiskLevelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the risk levels of instances.
//
// @param request - DescribeInstanceRiskLevelsRequest
//
// @return DescribeInstanceRiskLevelsResponse
func (client *Client) DescribeInstanceRiskLevels(request *DescribeInstanceRiskLevelsRequest) (_result *DescribeInstanceRiskLevelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceRiskLevelsResponse{}
	_body, _err := client.DescribeInstanceRiskLevelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取互联网方向删除会话趋势图
//
// @param request - DescribeInternetDropTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetDropTrafficTrendResponse
func (client *Client) DescribeInternetDropTrafficTrendWithOptions(request *DescribeInternetDropTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetDropTrafficTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceCode) {
		query["SourceCode"] = request.SourceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetDropTrafficTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetDropTrafficTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取互联网方向删除会话趋势图
//
// @param request - DescribeInternetDropTrafficTrendRequest
//
// @return DescribeInternetDropTrafficTrendResponse
func (client *Client) DescribeInternetDropTrafficTrend(request *DescribeInternetDropTrafficTrendRequest) (_result *DescribeInternetDropTrafficTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInternetDropTrafficTrendResponse{}
	_body, _err := client.DescribeInternetDropTrafficTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP addresses that are open to the Internet.
//
// @param request - DescribeInternetOpenIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenIpResponse
func (client *Client) DescribeInternetOpenIpWithOptions(request *DescribeInternetOpenIpRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetsInstanceId) {
		query["AssetsInstanceId"] = request.AssetsInstanceId
	}

	if !dara.IsNil(request.AssetsInstanceName) {
		query["AssetsInstanceName"] = request.AssetsInstanceName
	}

	if !dara.IsNil(request.AssetsType) {
		query["AssetsType"] = request.AssetsType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PublicIp) {
		query["PublicIp"] = request.PublicIp
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetOpenIp"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetOpenIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses that are open to the Internet.
//
// @param request - DescribeInternetOpenIpRequest
//
// @return DescribeInternetOpenIpResponse
func (client *Client) DescribeInternetOpenIp(request *DescribeInternetOpenIpRequest) (_result *DescribeInternetOpenIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInternetOpenIpResponse{}
	_body, _err := client.DescribeInternetOpenIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取互联网开放端口
//
// @param request - DescribeInternetOpenPortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenPortResponse
func (client *Client) DescribeInternetOpenPortWithOptions(request *DescribeInternetOpenPortRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenPortResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceNameFuzzy) {
		query["ServiceNameFuzzy"] = request.ServiceNameFuzzy
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SuggestLevel) {
		query["SuggestLevel"] = request.SuggestLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetOpenPort"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetOpenPortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取互联网开放端口
//
// @param request - DescribeInternetOpenPortRequest
//
// @return DescribeInternetOpenPortResponse
func (client *Client) DescribeInternetOpenPort(request *DescribeInternetOpenPortRequest) (_result *DescribeInternetOpenPortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInternetOpenPortResponse{}
	_body, _err := client.DescribeInternetOpenPortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取互联网开放服务
//
// @param request - DescribeInternetOpenServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenServiceResponse
func (client *Client) DescribeInternetOpenServiceWithOptions(request *DescribeInternetOpenServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceNameFuzzy) {
		query["ServiceNameFuzzy"] = request.ServiceNameFuzzy
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.SuggestLevel) {
		query["SuggestLevel"] = request.SuggestLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetOpenService"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetOpenServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取互联网开放服务
//
// @param request - DescribeInternetOpenServiceRequest
//
// @return DescribeInternetOpenServiceResponse
func (client *Client) DescribeInternetOpenService(request *DescribeInternetOpenServiceRequest) (_result *DescribeInternetOpenServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInternetOpenServiceResponse{}
	_body, _err := client.DescribeInternetOpenServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trends of Internet traffic.
//
// @param request - DescribeInternetTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetTrafficTrendResponse
func (client *Client) DescribeInternetTrafficTrendWithOptions(request *DescribeInternetTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetTrafficTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceCode) {
		query["SourceCode"] = request.SourceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.SrcPrivateIP) {
		query["SrcPrivateIP"] = request.SrcPrivateIP
	}

	if !dara.IsNil(request.SrcPublicIP) {
		query["SrcPublicIP"] = request.SrcPublicIP
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TrafficType) {
		query["TrafficType"] = request.TrafficType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetTrafficTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetTrafficTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trends of Internet traffic.
//
// @param request - DescribeInternetTrafficTrendRequest
//
// @return DescribeInternetTrafficTrendResponse
func (client *Client) DescribeInternetTrafficTrend(request *DescribeInternetTrafficTrendRequest) (_result *DescribeInternetTrafficTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInternetTrafficTrendResponse{}
	_body, _err := client.DescribeInternetTrafficTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the breach awareness events of a firewall.
//
// @param request - DescribeInvadeEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvadeEventListResponse
func (client *Client) DescribeInvadeEventListWithOptions(request *DescribeInvadeEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvadeEventListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetsIP) {
		query["AssetsIP"] = request.AssetsIP
	}

	if !dara.IsNil(request.AssetsInstanceId) {
		query["AssetsInstanceId"] = request.AssetsInstanceId
	}

	if !dara.IsNil(request.AssetsInstanceName) {
		query["AssetsInstanceName"] = request.AssetsInstanceName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventKey) {
		query["EventKey"] = request.EventKey
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
	}

	if !dara.IsNil(request.EventUuid) {
		query["EventUuid"] = request.EventUuid
	}

	if !dara.IsNil(request.IsIgnore) {
		query["IsIgnore"] = request.IsIgnore
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessStatusList) {
		query["ProcessStatusList"] = request.ProcessStatusList
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvadeEventList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvadeEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the breach awareness events of a firewall.
//
// @param request - DescribeInvadeEventListRequest
//
// @return DescribeInvadeEventListResponse
func (client *Client) DescribeInvadeEventList(request *DescribeInvadeEventListRequest) (_result *DescribeInvadeEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInvadeEventListResponse{}
	_body, _err := client.DescribeInvadeEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取漏洞名称列表
//
// @param request - DescribeInvadeEventNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvadeEventNameListResponse
func (client *Client) DescribeInvadeEventNameListWithOptions(request *DescribeInvadeEventNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvadeEventNameListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvadeEventNameList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvadeEventNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取漏洞名称列表
//
// @param request - DescribeInvadeEventNameListRequest
//
// @return DescribeInvadeEventNameListResponse
func (client *Client) DescribeInvadeEventNameList(request *DescribeInvadeEventNameListRequest) (_result *DescribeInvadeEventNameListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInvadeEventNameListResponse{}
	_body, _err := client.DescribeInvadeEventNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取漏洞事件统计
//
// @param request - DescribeInvadeEventStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvadeEventStatisticResponse
func (client *Client) DescribeInvadeEventStatisticWithOptions(request *DescribeInvadeEventStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvadeEventStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvadeEventStatistic"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvadeEventStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取漏洞事件统计
//
// @param request - DescribeInvadeEventStatisticRequest
//
// @return DescribeInvadeEventStatisticResponse
func (client *Client) DescribeInvadeEventStatistic(request *DescribeInvadeEventStatisticRequest) (_result *DescribeInvadeEventStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInvadeEventStatisticResponse{}
	_body, _err := client.DescribeInvadeEventStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Log Service Information
//
// @param request - DescribeLogStoreInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogStoreInfoResponse
func (client *Client) DescribeLogStoreInfoWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeLogStoreInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogStoreInfo"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogStoreInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Log Service Information
//
// @return DescribeLogStoreInfoResponse
func (client *Client) DescribeLogStoreInfo() (_result *DescribeLogStoreInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogStoreInfoResponse{}
	_body, _err := client.DescribeLogStoreInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the pagination status of NAT firewalls.
//
// @param request - DescribeNatAclPageStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatAclPageStatusResponse
func (client *Client) DescribeNatAclPageStatusWithOptions(request *DescribeNatAclPageStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatAclPageStatusResponse, _err error) {
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
		Action:      dara.String("DescribeNatAclPageStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatAclPageStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the pagination status of NAT firewalls.
//
// @param request - DescribeNatAclPageStatusRequest
//
// @return DescribeNatAclPageStatusResponse
func (client *Client) DescribeNatAclPageStatus(request *DescribeNatAclPageStatusRequest) (_result *DescribeNatAclPageStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNatAclPageStatusResponse{}
	_body, _err := client.DescribeNatAclPageStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all access control policies that are created for NAT firewalls.
//
// Description:
//
// You can use this operation to query the information about all access control policies that are created for NAT firewalls by page.
//
// @param request - DescribeNatFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallControlPolicyResponse
func (client *Client) DescribeNatFirewallControlPolicyWithOptions(request *DescribeNatFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNatFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all access control policies that are created for NAT firewalls.
//
// Description:
//
// You can use this operation to query the information about all access control policies that are created for NAT firewalls by page.
//
// @param request - DescribeNatFirewallControlPolicyRequest
//
// @return DescribeNatFirewallControlPolicyResponse
func (client *Client) DescribeNatFirewallControlPolicy(request *DescribeNatFirewallControlPolicyRequest) (_result *DescribeNatFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNatFirewallControlPolicyResponse{}
	_body, _err := client.DescribeNatFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 概览页-nat防火墙拦截趋势
//
// @param request - DescribeNatFirewallDropTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallDropTrafficTrendResponse
func (client *Client) DescribeNatFirewallDropTrafficTrendWithOptions(request *DescribeNatFirewallDropTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallDropTrafficTrendResponse, _err error) {
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
		Action:      dara.String("DescribeNatFirewallDropTrafficTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallDropTrafficTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 概览页-nat防火墙拦截趋势
//
// @param request - DescribeNatFirewallDropTrafficTrendRequest
//
// @return DescribeNatFirewallDropTrafficTrendResponse
func (client *Client) DescribeNatFirewallDropTrafficTrend(request *DescribeNatFirewallDropTrafficTrendRequest) (_result *DescribeNatFirewallDropTrafficTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNatFirewallDropTrafficTrendResponse{}
	_body, _err := client.DescribeNatFirewallDropTrafficTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details of NAT firewalls.
//
// @param request - DescribeNatFirewallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallListResponse
func (client *Client) DescribeNatFirewallListWithOptions(request *DescribeNatFirewallListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallListResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.ProxyName) {
		query["ProxyName"] = request.ProxyName
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNatFirewallList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details of NAT firewalls.
//
// @param request - DescribeNatFirewallListRequest
//
// @return DescribeNatFirewallListResponse
func (client *Client) DescribeNatFirewallList(request *DescribeNatFirewallListRequest) (_result *DescribeNatFirewallListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNatFirewallListResponse{}
	_body, _err := client.DescribeNatFirewallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the priority range of access control policies that are created for a NAT firewall.
//
// Description:
//
// You can use this operation to query the priority range of access control policies that are created for a NAT firewall.
//
// @param request - DescribeNatFirewallPolicyPriorUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallPolicyPriorUsedResponse
func (client *Client) DescribeNatFirewallPolicyPriorUsedWithOptions(request *DescribeNatFirewallPolicyPriorUsedRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallPolicyPriorUsedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNatFirewallPolicyPriorUsed"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallPolicyPriorUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the priority range of access control policies that are created for a NAT firewall.
//
// Description:
//
// You can use this operation to query the priority range of access control policies that are created for a NAT firewall.
//
// @param request - DescribeNatFirewallPolicyPriorUsedRequest
//
// @return DescribeNatFirewallPolicyPriorUsedResponse
func (client *Client) DescribeNatFirewallPolicyPriorUsed(request *DescribeNatFirewallPolicyPriorUsedRequest) (_result *DescribeNatFirewallPolicyPriorUsedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNatFirewallPolicyPriorUsedResponse{}
	_body, _err := client.DescribeNatFirewallPolicyPriorUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取NAT防火墙配额
//
// @param request - DescribeNatFirewallQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallQuotaResponse
func (client *Client) DescribeNatFirewallQuotaWithOptions(request *DescribeNatFirewallQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallQuotaResponse, _err error) {
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
		Action:      dara.String("DescribeNatFirewallQuota"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取NAT防火墙配额
//
// @param request - DescribeNatFirewallQuotaRequest
//
// @return DescribeNatFirewallQuotaResponse
func (client *Client) DescribeNatFirewallQuota(request *DescribeNatFirewallQuotaRequest) (_result *DescribeNatFirewallQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNatFirewallQuotaResponse{}
	_body, _err := client.DescribeNatFirewallQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 概览页-NAT流量趋势
//
// @param request - DescribeNatFirewallTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallTrafficTrendResponse
func (client *Client) DescribeNatFirewallTrafficTrendWithOptions(request *DescribeNatFirewallTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallTrafficTrendResponse, _err error) {
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
		Action:      dara.String("DescribeNatFirewallTrafficTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallTrafficTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 概览页-NAT流量趋势
//
// @param request - DescribeNatFirewallTrafficTrendRequest
//
// @return DescribeNatFirewallTrafficTrendResponse
func (client *Client) DescribeNatFirewallTrafficTrend(request *DescribeNatFirewallTrafficTrendRequest) (_result *DescribeNatFirewallTrafficTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNatFirewallTrafficTrendResponse{}
	_body, _err := client.DescribeNatFirewallTrafficTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网络实例列表
//
// @param request - DescribeNetworkInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkInstanceListResponse
func (client *Client) DescribeNetworkInstanceListWithOptions(request *DescribeNetworkInstanceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkInstanceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ConnectType) {
		query["ConnectType"] = request.ConnectType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkInstanceList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网络实例列表
//
// @param request - DescribeNetworkInstanceListRequest
//
// @return DescribeNetworkInstanceListResponse
func (client *Client) DescribeNetworkInstanceList(request *DescribeNetworkInstanceListRequest) (_result *DescribeNetworkInstanceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkInstanceListResponse{}
	_body, _err := client.DescribeNetworkInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网络实例关系列表
//
// @param request - DescribeNetworkInstanceRelationListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkInstanceRelationListResponse
func (client *Client) DescribeNetworkInstanceRelationListWithOptions(request *DescribeNetworkInstanceRelationListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkInstanceRelationListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectType) {
		query["ConnectType"] = request.ConnectType
	}

	if !dara.IsNil(request.FirewallConfigureStatus) {
		query["FirewallConfigureStatus"] = request.FirewallConfigureStatus
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NetworkInstanceId) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !dara.IsNil(request.PeerNetworkInstanceId) {
		query["PeerNetworkInstanceId"] = request.PeerNetworkInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkInstanceRelationList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkInstanceRelationListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网络实例关系列表
//
// @param request - DescribeNetworkInstanceRelationListRequest
//
// @return DescribeNetworkInstanceRelationListResponse
func (client *Client) DescribeNetworkInstanceRelationList(request *DescribeNetworkInstanceRelationListRequest) (_result *DescribeNetworkInstanceRelationListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkInstanceRelationListResponse{}
	_body, _err := client.DescribeNetworkInstanceRelationListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网络流量TOP环比
//
// @param request - DescribeNetworkTrafficTopRatioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkTrafficTopRatioResponse
func (client *Client) DescribeNetworkTrafficTopRatioWithOptions(request *DescribeNetworkTrafficTopRatioRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkTrafficTopRatioResponse, _err error) {
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

	if !dara.IsNil(request.AssetIP) {
		query["AssetIP"] = request.AssetIP
	}

	if !dara.IsNil(request.AssetRegion) {
		query["AssetRegion"] = request.AssetRegion
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DstIP) {
		query["DstIP"] = request.DstIP
	}

	if !dara.IsNil(request.DstPort) {
		query["DstPort"] = request.DstPort
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IpProperty) {
		query["IpProperty"] = request.IpProperty
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.RuleResult) {
		query["RuleResult"] = request.RuleResult
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SourceCode) {
		query["SourceCode"] = request.SourceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.SrcIP) {
		query["SrcIP"] = request.SrcIP
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkTrafficTopRatio"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkTrafficTopRatioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网络流量TOP环比
//
// @param request - DescribeNetworkTrafficTopRatioRequest
//
// @return DescribeNetworkTrafficTopRatioResponse
func (client *Client) DescribeNetworkTrafficTopRatio(request *DescribeNetworkTrafficTopRatioRequest) (_result *DescribeNetworkTrafficTopRatioResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkTrafficTopRatioResponse{}
	_body, _err := client.DescribeNetworkTrafficTopRatioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取公网IP通过的源状态
//
// @param request - DescribeOpenIpAccessSrcStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenIpAccessSrcStatResponse
func (client *Client) DescribeOpenIpAccessSrcStatWithOptions(request *DescribeOpenIpAccessSrcStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeOpenIpAccessSrcStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DstIp) {
		query["DstIp"] = request.DstIp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOpenIpAccessSrcStat"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOpenIpAccessSrcStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取公网IP通过的源状态
//
// @param request - DescribeOpenIpAccessSrcStatRequest
//
// @return DescribeOpenIpAccessSrcStatResponse
func (client *Client) DescribeOpenIpAccessSrcStat(request *DescribeOpenIpAccessSrcStatRequest) (_result *DescribeOpenIpAccessSrcStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOpenIpAccessSrcStatResponse{}
	_body, _err := client.DescribeOpenIpAccessSrcStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取外联资产列表
//
// @param request - DescribeOutgoingAssetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingAssetListResponse
func (client *Client) DescribeOutgoingAssetListWithOptions(request *DescribeOutgoingAssetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingAssetListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetsRegion) {
		query["AssetsRegion"] = request.AssetsRegion
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IPType) {
		query["IPType"] = request.IPType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.NatGatewayName) {
		query["NatGatewayName"] = request.NatGatewayName
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivateIP) {
		query["PrivateIP"] = request.PrivateIP
	}

	if !dara.IsNil(request.PublicIP) {
		query["PublicIP"] = request.PublicIP
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SecurityRisk) {
		query["SecurityRisk"] = request.SecurityRisk
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOutgoingAssetList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingAssetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取外联资产列表
//
// @param request - DescribeOutgoingAssetListRequest
//
// @return DescribeOutgoingAssetListResponse
func (client *Client) DescribeOutgoingAssetList(request *DescribeOutgoingAssetListRequest) (_result *DescribeOutgoingAssetListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOutgoingAssetListResponse{}
	_body, _err := client.DescribeOutgoingAssetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get details of outgoing destination IPs
//
// @param request - DescribeOutgoingDestinationIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDestinationIPResponse
func (client *Client) DescribeOutgoingDestinationIPWithOptions(request *DescribeOutgoingDestinationIPRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DstIP) {
		query["DstIP"] = request.DstIP
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateIP) {
		query["PrivateIP"] = request.PrivateIP
	}

	if !dara.IsNil(request.PublicIP) {
		query["PublicIP"] = request.PublicIP
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TagIdNew) {
		query["TagIdNew"] = request.TagIdNew
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOutgoingDestinationIP"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingDestinationIPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get details of outgoing destination IPs
//
// @param request - DescribeOutgoingDestinationIPRequest
//
// @return DescribeOutgoingDestinationIPResponse
func (client *Client) DescribeOutgoingDestinationIP(request *DescribeOutgoingDestinationIPRequest) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOutgoingDestinationIPResponse{}
	_body, _err := client.DescribeOutgoingDestinationIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the domain names in outbound connections.
//
// @param request - DescribeOutgoingDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDomainResponse
func (client *Client) DescribeOutgoingDomainWithOptions(request *DescribeOutgoingDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IsAITraffic) {
		query["IsAITraffic"] = request.IsAITraffic
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PublicIP) {
		query["PublicIP"] = request.PublicIP
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TagIdNew) {
		query["TagIdNew"] = request.TagIdNew
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOutgoingDomain"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the domain names in outbound connections.
//
// @param request - DescribeOutgoingDomainRequest
//
// @return DescribeOutgoingDomainResponse
func (client *Client) DescribeOutgoingDomain(request *DescribeOutgoingDomainRequest) (_result *DescribeOutgoingDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOutgoingDomainResponse{}
	_body, _err := client.DescribeOutgoingDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取外联域名详情
//
// @param request - DescribeOutgoingDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDomainDetailResponse
func (client *Client) DescribeOutgoingDomainDetailWithOptions(request *DescribeOutgoingDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDomainDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclCoverage) {
		query["AclCoverage"] = request.AclCoverage
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.DomainList) {
		query["DomainList"] = request.DomainList
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IPType) {
		query["IPType"] = request.IPType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivateIP) {
		query["PrivateIP"] = request.PrivateIP
	}

	if !dara.IsNil(request.PublicIP) {
		query["PublicIP"] = request.PublicIP
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOutgoingDomainDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取外联域名详情
//
// @param request - DescribeOutgoingDomainDetailRequest
//
// @return DescribeOutgoingDomainDetailResponse
func (client *Client) DescribeOutgoingDomainDetail(request *DescribeOutgoingDomainDetailRequest) (_result *DescribeOutgoingDomainDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOutgoingDomainDetailResponse{}
	_body, _err := client.DescribeOutgoingDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取外联统计
//
// @param request - DescribeOutgoingStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingStatisticResponse
func (client *Client) DescribeOutgoingStatisticWithOptions(request *DescribeOutgoingStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOutgoingStatistic"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取外联统计
//
// @param request - DescribeOutgoingStatisticRequest
//
// @return DescribeOutgoingStatisticResponse
func (client *Client) DescribeOutgoingStatistic(request *DescribeOutgoingStatisticRequest) (_result *DescribeOutgoingStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOutgoingStatisticResponse{}
	_body, _err := client.DescribeOutgoingStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取外联标签
//
// @param request - DescribeOutgoingTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingTagResponse
func (client *Client) DescribeOutgoingTagWithOptions(request *DescribeOutgoingTagRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DstType) {
		query["DstType"] = request.DstType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOutgoingTag"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取外联标签
//
// @param request - DescribeOutgoingTagRequest
//
// @return DescribeOutgoingTagResponse
func (client *Client) DescribeOutgoingTag(request *DescribeOutgoingTagRequest) (_result *DescribeOutgoingTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOutgoingTagResponse{}
	_body, _err := client.DescribeOutgoingTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档
//
// @param request - DescribePageDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePageDocumentsResponse
func (client *Client) DescribePageDocumentsWithOptions(request *DescribePageDocumentsRequest, runtime *dara.RuntimeOptions) (_result *DescribePageDocumentsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageName) {
		query["PageName"] = request.PageName
	}

	if !dara.IsNil(request.SourceCode) {
		query["SourceCode"] = request.SourceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.TabName) {
		query["TabName"] = request.TabName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePageDocuments"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePageDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档
//
// @param request - DescribePageDocumentsRequest
//
// @return DescribePageDocumentsResponse
func (client *Client) DescribePageDocuments(request *DescribePageDocumentsRequest) (_result *DescribePageDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePageDocumentsResponse{}
	_body, _err := client.DescribePageDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the strict mode is enabled for an access control policy.
//
// Description:
//
// You can call the DescribePolicyAdvancedConfig operation to query whether the strict mode is enabled for an access control policy.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePolicyAdvancedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyAdvancedConfigResponse
func (client *Client) DescribePolicyAdvancedConfigWithOptions(request *DescribePolicyAdvancedConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyAdvancedConfigResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicyAdvancedConfig"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolicyAdvancedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the strict mode is enabled for an access control policy.
//
// Description:
//
// You can call the DescribePolicyAdvancedConfig operation to query whether the strict mode is enabled for an access control policy.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePolicyAdvancedConfigRequest
//
// @return DescribePolicyAdvancedConfigResponse
func (client *Client) DescribePolicyAdvancedConfig(request *DescribePolicyAdvancedConfigRequest) (_result *DescribePolicyAdvancedConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolicyAdvancedConfigResponse{}
	_body, _err := client.DescribePolicyAdvancedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the priority range of access control policies.
//
// Description:
//
// You can call this operation to query the priority range of the access control policies that match specific query conditions.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePolicyPriorUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyPriorUsedResponse
func (client *Client) DescribePolicyPriorUsedWithOptions(request *DescribePolicyPriorUsedRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyPriorUsedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicyPriorUsed"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolicyPriorUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the priority range of access control policies.
//
// Description:
//
// You can call this operation to query the priority range of the access control policies that match specific query conditions.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePolicyPriorUsedRequest
//
// @return DescribePolicyPriorUsedResponse
func (client *Client) DescribePolicyPriorUsed(request *DescribePolicyPriorUsedRequest) (_result *DescribePolicyPriorUsedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolicyPriorUsedResponse{}
	_body, _err := client.DescribePolicyPriorUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of traffic billed based on the pay-as-you-go billing method.
//
// Description:
//
// If you use Cloud Firewall that uses the pay-as-you-go billing method, you can call this operation to query traffic details accurate to the granularity of specific resource instances. If you use Cloud Firewall that uses the subscription billing method, you can call this operation to query the overall traffic details.
//
// @param request - DescribePostpayTrafficDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayTrafficDetailResponse
func (client *Client) DescribePostpayTrafficDetailWithOptions(request *DescribePostpayTrafficDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayTrafficDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.SearchItem) {
		query["SearchItem"] = request.SearchItem
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TrafficType) {
		query["TrafficType"] = request.TrafficType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePostpayTrafficDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePostpayTrafficDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of traffic billed based on the pay-as-you-go billing method.
//
// Description:
//
// If you use Cloud Firewall that uses the pay-as-you-go billing method, you can call this operation to query traffic details accurate to the granularity of specific resource instances. If you use Cloud Firewall that uses the subscription billing method, you can call this operation to query the overall traffic details.
//
// @param request - DescribePostpayTrafficDetailRequest
//
// @return DescribePostpayTrafficDetailResponse
func (client *Client) DescribePostpayTrafficDetail(request *DescribePostpayTrafficDetailRequest) (_result *DescribePostpayTrafficDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePostpayTrafficDetailResponse{}
	_body, _err := client.DescribePostpayTrafficDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total volume of traffic that is billed based on the pay-as-you-go billing method, including all firewalls within the current account.
//
// Description:
//
// You can call this operation to query statistics of the current Cloud Firewall from the date of purchase.
//
// @param request - DescribePostpayTrafficTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayTrafficTotalResponse
func (client *Client) DescribePostpayTrafficTotalWithOptions(request *DescribePostpayTrafficTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayTrafficTotalResponse, _err error) {
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
		Action:      dara.String("DescribePostpayTrafficTotal"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePostpayTrafficTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total volume of traffic that is billed based on the pay-as-you-go billing method, including all firewalls within the current account.
//
// Description:
//
// You can call this operation to query statistics of the current Cloud Firewall from the date of purchase.
//
// @param request - DescribePostpayTrafficTotalRequest
//
// @return DescribePostpayTrafficTotalResponse
func (client *Client) DescribePostpayTrafficTotal(request *DescribePostpayTrafficTotalRequest) (_result *DescribePostpayTrafficTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePostpayTrafficTotalResponse{}
	_body, _err := client.DescribePostpayTrafficTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Internet Firewall feature in Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - DescribePostpayUserInternetStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayUserInternetStatusResponse
func (client *Client) DescribePostpayUserInternetStatusWithOptions(request *DescribePostpayUserInternetStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayUserInternetStatusResponse, _err error) {
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
		Action:      dara.String("DescribePostpayUserInternetStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePostpayUserInternetStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Internet Firewall feature in Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - DescribePostpayUserInternetStatusRequest
//
// @return DescribePostpayUserInternetStatusResponse
func (client *Client) DescribePostpayUserInternetStatus(request *DescribePostpayUserInternetStatusRequest) (_result *DescribePostpayUserInternetStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePostpayUserInternetStatusResponse{}
	_body, _err := client.DescribePostpayUserInternetStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the NAT Firewall feature in Cloud Firewall that use the pay-as-you-go billing method.
//
// @param request - DescribePostpayUserNatStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayUserNatStatusResponse
func (client *Client) DescribePostpayUserNatStatusWithOptions(request *DescribePostpayUserNatStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayUserNatStatusResponse, _err error) {
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
		Action:      dara.String("DescribePostpayUserNatStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePostpayUserNatStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the NAT Firewall feature in Cloud Firewall that use the pay-as-you-go billing method.
//
// @param request - DescribePostpayUserNatStatusRequest
//
// @return DescribePostpayUserNatStatusResponse
func (client *Client) DescribePostpayUserNatStatus(request *DescribePostpayUserNatStatusRequest) (_result *DescribePostpayUserNatStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePostpayUserNatStatusResponse{}
	_body, _err := client.DescribePostpayUserNatStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the virtual private cloud (VPC) Firewall feature in Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - DescribePostpayUserVpcStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayUserVpcStatusResponse
func (client *Client) DescribePostpayUserVpcStatusWithOptions(request *DescribePostpayUserVpcStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayUserVpcStatusResponse, _err error) {
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
		Action:      dara.String("DescribePostpayUserVpcStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePostpayUserVpcStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the virtual private cloud (VPC) Firewall feature in Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - DescribePostpayUserVpcStatusRequest
//
// @return DescribePostpayUserVpcStatusResponse
func (client *Client) DescribePostpayUserVpcStatus(request *DescribePostpayUserVpcStatusRequest) (_result *DescribePostpayUserVpcStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePostpayUserVpcStatusResponse{}
	_body, _err := client.DescribePostpayUserVpcStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries prefix lists.
//
// @param request - DescribePrefixListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrefixListsResponse
func (client *Client) DescribePrefixListsWithOptions(request *DescribePrefixListsRequest, runtime *dara.RuntimeOptions) (_result *DescribePrefixListsResponse, _err error) {
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
		Action:      dara.String("DescribePrefixLists"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrefixListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries prefix lists.
//
// @param request - DescribePrefixListsRequest
//
// @return DescribePrefixListsResponse
func (client *Client) DescribePrefixLists(request *DescribePrefixListsRequest) (_result *DescribePrefixListsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePrefixListsResponse{}
	_body, _err := client.DescribePrefixListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询私网DNS域名列表
//
// @param request - DescribePrivateDnsDomainNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrivateDnsDomainNameListResponse
func (client *Client) DescribePrivateDnsDomainNameListWithOptions(request *DescribePrivateDnsDomainNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribePrivateDnsDomainNameListResponse, _err error) {
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
		Action:      dara.String("DescribePrivateDnsDomainNameList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrivateDnsDomainNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询私网DNS域名列表
//
// @param request - DescribePrivateDnsDomainNameListRequest
//
// @return DescribePrivateDnsDomainNameListResponse
func (client *Client) DescribePrivateDnsDomainNameList(request *DescribePrivateDnsDomainNameListRequest) (_result *DescribePrivateDnsDomainNameListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePrivateDnsDomainNameListResponse{}
	_body, _err := client.DescribePrivateDnsDomainNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询私网DNS终端节点详情
//
// @param request - DescribePrivateDnsEndpointDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrivateDnsEndpointDetailResponse
func (client *Client) DescribePrivateDnsEndpointDetailWithOptions(request *DescribePrivateDnsEndpointDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribePrivateDnsEndpointDetailResponse, _err error) {
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
		Action:      dara.String("DescribePrivateDnsEndpointDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrivateDnsEndpointDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询私网DNS终端节点详情
//
// @param request - DescribePrivateDnsEndpointDetailRequest
//
// @return DescribePrivateDnsEndpointDetailResponse
func (client *Client) DescribePrivateDnsEndpointDetail(request *DescribePrivateDnsEndpointDetailRequest) (_result *DescribePrivateDnsEndpointDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePrivateDnsEndpointDetailResponse{}
	_body, _err := client.DescribePrivateDnsEndpointDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询私网DNS终端节点列表
//
// @param request - DescribePrivateDnsEndpointListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrivateDnsEndpointListResponse
func (client *Client) DescribePrivateDnsEndpointListWithOptions(request *DescribePrivateDnsEndpointListRequest, runtime *dara.RuntimeOptions) (_result *DescribePrivateDnsEndpointListResponse, _err error) {
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
		Action:      dara.String("DescribePrivateDnsEndpointList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrivateDnsEndpointListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询私网DNS终端节点列表
//
// @param request - DescribePrivateDnsEndpointListRequest
//
// @return DescribePrivateDnsEndpointListResponse
func (client *Client) DescribePrivateDnsEndpointList(request *DescribePrivateDnsEndpointListRequest) (_result *DescribePrivateDnsEndpointListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePrivateDnsEndpointListResponse{}
	_body, _err := client.DescribePrivateDnsEndpointListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取地域信息
//
// @param request - DescribeRegionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionInfoResponse
func (client *Client) DescribeRegionInfoWithOptions(request *DescribeRegionInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionInfoResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceCode) {
		query["SourceCode"] = request.SourceCode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegionInfo"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取地域信息
//
// @param request - DescribeRegionInfoRequest
//
// @return DescribeRegionInfoResponse
func (client *Client) DescribeRegionInfo(request *DescribeRegionInfoRequest) (_result *DescribeRegionInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionInfoResponse{}
	_body, _err := client.DescribeRegionInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资产类型默认引流
//
// @param request - DescribeResourceTypeAutoEnableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceTypeAutoEnableResponse
func (client *Client) DescribeResourceTypeAutoEnableWithOptions(request *DescribeResourceTypeAutoEnableRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceTypeAutoEnableResponse, _err error) {
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
		Action:      dara.String("DescribeResourceTypeAutoEnable"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceTypeAutoEnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资产类型默认引流
//
// @param request - DescribeResourceTypeAutoEnableRequest
//
// @return DescribeResourceTypeAutoEnableResponse
func (client *Client) DescribeResourceTypeAutoEnable(request *DescribeResourceTypeAutoEnableRequest) (_result *DescribeResourceTypeAutoEnableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceTypeAutoEnableResponse{}
	_body, _err := client.DescribeResourceTypeAutoEnableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of intrusion events.
//
// Description:
//
// You can call the DescribeRiskEventGroup operation to query and download the details of intrusion events. We recommend that you query the details of 5 to 10 intrusion events at a time. If you do not need to query the geographical information about IP addresses, you can set the NoLocation parameter to true to prevent query timeout.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRiskEventGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventGroupResponse
func (client *Client) DescribeRiskEventGroupWithOptions(request *DescribeRiskEventGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttackApp) {
		query["AttackApp"] = request.AttackApp
	}

	if !dara.IsNil(request.AttackAppCategory) {
		query["AttackAppCategory"] = request.AttackAppCategory
	}

	if !dara.IsNil(request.AttackType) {
		query["AttackType"] = request.AttackType
	}

	if !dara.IsNil(request.BuyVersion) {
		query["BuyVersion"] = request.BuyVersion
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DstIP) {
		query["DstIP"] = request.DstIP
	}

	if !dara.IsNil(request.DstNetworkInstanceId) {
		query["DstNetworkInstanceId"] = request.DstNetworkInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
	}

	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.IsOnlyPrivateAssoc) {
		query["IsOnlyPrivateAssoc"] = request.IsOnlyPrivateAssoc
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NoLocation) {
		query["NoLocation"] = request.NoLocation
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleResult) {
		query["RuleResult"] = request.RuleResult
	}

	if !dara.IsNil(request.RuleSource) {
		query["RuleSource"] = request.RuleSource
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SrcIP) {
		query["SrcIP"] = request.SrcIP
	}

	if !dara.IsNil(request.SrcNetworkInstanceId) {
		query["SrcNetworkInstanceId"] = request.SrcNetworkInstanceId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VulLevel) {
		query["VulLevel"] = request.VulLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRiskEventGroup"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskEventGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of intrusion events.
//
// Description:
//
// You can call the DescribeRiskEventGroup operation to query and download the details of intrusion events. We recommend that you query the details of 5 to 10 intrusion events at a time. If you do not need to query the geographical information about IP addresses, you can set the NoLocation parameter to true to prevent query timeout.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeRiskEventGroupRequest
//
// @return DescribeRiskEventGroupResponse
func (client *Client) DescribeRiskEventGroup(request *DescribeRiskEventGroupRequest) (_result *DescribeRiskEventGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRiskEventGroupResponse{}
	_body, _err := client.DescribeRiskEventGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attack payloads of intrusion events.
//
// @param request - DescribeRiskEventPayloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventPayloadResponse
func (client *Client) DescribeRiskEventPayloadWithOptions(request *DescribeRiskEventPayloadRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventPayloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DstIP) {
		query["DstIP"] = request.DstIP
	}

	if !dara.IsNil(request.DstVpcId) {
		query["DstVpcId"] = request.DstVpcId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.PublicIP) {
		query["PublicIP"] = request.PublicIP
	}

	if !dara.IsNil(request.SrcIP) {
		query["SrcIP"] = request.SrcIP
	}

	if !dara.IsNil(request.SrcVpcId) {
		query["SrcVpcId"] = request.SrcVpcId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UUID) {
		query["UUID"] = request.UUID
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRiskEventPayload"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskEventPayloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attack payloads of intrusion events.
//
// @param request - DescribeRiskEventPayloadRequest
//
// @return DescribeRiskEventPayloadResponse
func (client *Client) DescribeRiskEventPayload(request *DescribeRiskEventPayloadRequest) (_result *DescribeRiskEventPayloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRiskEventPayloadResponse{}
	_body, _err := client.DescribeRiskEventPayloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Top风险事件资产
//
// @param request - DescribeRiskEventTopAttackAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventTopAttackAssetResponse
func (client *Client) DescribeRiskEventTopAttackAssetWithOptions(request *DescribeRiskEventTopAttackAssetRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventTopAttackAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttackApp) {
		query["AttackApp"] = request.AttackApp
	}

	if !dara.IsNil(request.AttackType) {
		query["AttackType"] = request.AttackType
	}

	if !dara.IsNil(request.BuyVersion) {
		query["BuyVersion"] = request.BuyVersion
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRiskEventTopAttackAsset"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskEventTopAttackAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Top风险事件资产
//
// @param request - DescribeRiskEventTopAttackAssetRequest
//
// @return DescribeRiskEventTopAttackAssetResponse
func (client *Client) DescribeRiskEventTopAttackAsset(request *DescribeRiskEventTopAttackAssetRequest) (_result *DescribeRiskEventTopAttackAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRiskEventTopAttackAssetResponse{}
	_body, _err := client.DescribeRiskEventTopAttackAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取风险安全组详情
//
// @param request - DescribeRiskSecurityGroupDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskSecurityGroupDetailResponse
func (client *Client) DescribeRiskSecurityGroupDetailWithOptions(request *DescribeRiskSecurityGroupDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskSecurityGroupDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleUuid) {
		query["RuleUuid"] = request.RuleUuid
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRiskSecurityGroupDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskSecurityGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取风险安全组详情
//
// @param request - DescribeRiskSecurityGroupDetailRequest
//
// @return DescribeRiskSecurityGroupDetailResponse
func (client *Client) DescribeRiskSecurityGroupDetail(request *DescribeRiskSecurityGroupDetailRequest) (_result *DescribeRiskSecurityGroupDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRiskSecurityGroupDetailResponse{}
	_body, _err := client.DescribeRiskSecurityGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取安全模式
//
// @param request - DescribeSecurityModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityModeResponse
func (client *Client) DescribeSecurityModeWithOptions(request *DescribeSecurityModeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityModeResponse, _err error) {
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

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityMode"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全模式
//
// @param request - DescribeSecurityModeRequest
//
// @return DescribeSecurityModeResponse
func (client *Client) DescribeSecurityMode(request *DescribeSecurityModeRequest) (_result *DescribeSecurityModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityModeResponse{}
	_body, _err := client.DescribeSecurityModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取正向代理
//
// @param request - DescribeSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityProxyResponse
func (client *Client) DescribeSecurityProxyWithOptions(request *DescribeSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityProxyResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.ProxyName) {
		query["ProxyName"] = request.ProxyName
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityProxy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取正向代理
//
// @param request - DescribeSecurityProxyRequest
//
// @return DescribeSecurityProxyResponse
func (client *Client) DescribeSecurityProxy(request *DescribeSecurityProxyRequest) (_result *DescribeSecurityProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityProxyResponse{}
	_body, _err := client.DescribeSecurityProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about signature library versions.
//
// @param request - DescribeSignatureLibVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSignatureLibVersionResponse
func (client *Client) DescribeSignatureLibVersionWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeSignatureLibVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSignatureLibVersion"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSignatureLibVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about signature library versions.
//
// @return DescribeSignatureLibVersionResponse
func (client *Client) DescribeSignatureLibVersion() (_result *DescribeSignatureLibVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSignatureLibVersionResponse{}
	_body, _err := client.DescribeSignatureLibVersionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取SLS开启状态
//
// @param request - DescribeSlsAnalyzeOpenStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsAnalyzeOpenStatusResponse
func (client *Client) DescribeSlsAnalyzeOpenStatusWithOptions(request *DescribeSlsAnalyzeOpenStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsAnalyzeOpenStatusResponse, _err error) {
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
		Action:      dara.String("DescribeSlsAnalyzeOpenStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsAnalyzeOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取SLS开启状态
//
// @param request - DescribeSlsAnalyzeOpenStatusRequest
//
// @return DescribeSlsAnalyzeOpenStatusResponse
func (client *Client) DescribeSlsAnalyzeOpenStatus(request *DescribeSlsAnalyzeOpenStatusRequest) (_result *DescribeSlsAnalyzeOpenStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlsAnalyzeOpenStatusResponse{}
	_body, _err := client.DescribeSlsAnalyzeOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询威胁情报配置的信息
//
// @param request - DescribeThreatIntelligenceSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeThreatIntelligenceSwitchResponse
func (client *Client) DescribeThreatIntelligenceSwitchWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeThreatIntelligenceSwitchResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeThreatIntelligenceSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeThreatIntelligenceSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询威胁情报配置的信息
//
// @return DescribeThreatIntelligenceSwitchResponse
func (client *Client) DescribeThreatIntelligenceSwitch() (_result *DescribeThreatIntelligenceSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeThreatIntelligenceSwitchResponse{}
	_body, _err := client.DescribeThreatIntelligenceSwitchWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the transit routers that are associated with a virtual private cloud (VPC) firewall created for a transit router.
//
// @param tmpReq - DescribeTrFirewallPolicyBackUpAssociationListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallPolicyBackUpAssociationListResponse
func (client *Client) DescribeTrFirewallPolicyBackUpAssociationListWithOptions(tmpReq *DescribeTrFirewallPolicyBackUpAssociationListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallPolicyBackUpAssociationListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CandidateList) {
		request.CandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CandidateList, dara.String("CandidateList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CandidateListShrink) {
		query["CandidateList"] = request.CandidateListShrink
	}

	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TrFirewallRoutePolicyId) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrFirewallPolicyBackUpAssociationList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrFirewallPolicyBackUpAssociationListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the transit routers that are associated with a virtual private cloud (VPC) firewall created for a transit router.
//
// @param request - DescribeTrFirewallPolicyBackUpAssociationListRequest
//
// @return DescribeTrFirewallPolicyBackUpAssociationListResponse
func (client *Client) DescribeTrFirewallPolicyBackUpAssociationList(request *DescribeTrFirewallPolicyBackUpAssociationListRequest) (_result *DescribeTrFirewallPolicyBackUpAssociationListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrFirewallPolicyBackUpAssociationListResponse{}
	_body, _err := client.DescribeTrFirewallPolicyBackUpAssociationListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the routing policies of a virtual private cloud (VPC) firewall that is created for a transit router.
//
// @param request - DescribeTrFirewallV2RoutePolicyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallV2RoutePolicyListResponse
func (client *Client) DescribeTrFirewallV2RoutePolicyListWithOptions(request *DescribeTrFirewallV2RoutePolicyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallV2RoutePolicyListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrFirewallV2RoutePolicyList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrFirewallV2RoutePolicyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routing policies of a virtual private cloud (VPC) firewall that is created for a transit router.
//
// @param request - DescribeTrFirewallV2RoutePolicyListRequest
//
// @return DescribeTrFirewallV2RoutePolicyListResponse
func (client *Client) DescribeTrFirewallV2RoutePolicyList(request *DescribeTrFirewallV2RoutePolicyListRequest) (_result *DescribeTrFirewallV2RoutePolicyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrFirewallV2RoutePolicyListResponse{}
	_body, _err := client.DescribeTrFirewallV2RoutePolicyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the virtual private cloud (VPC) firewalls that are created for transit routers.
//
// @param request - DescribeTrFirewallsV2DetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallsV2DetailResponse
func (client *Client) DescribeTrFirewallsV2DetailWithOptions(request *DescribeTrFirewallsV2DetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallsV2DetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrFirewallsV2Detail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrFirewallsV2DetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the virtual private cloud (VPC) firewalls that are created for transit routers.
//
// @param request - DescribeTrFirewallsV2DetailRequest
//
// @return DescribeTrFirewallsV2DetailResponse
func (client *Client) DescribeTrFirewallsV2Detail(request *DescribeTrFirewallsV2DetailRequest) (_result *DescribeTrFirewallsV2DetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrFirewallsV2DetailResponse{}
	_body, _err := client.DescribeTrFirewallsV2DetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the virtual private cloud (VPC) firewalls that are created for transit routers.
//
// @param request - DescribeTrFirewallsV2ListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallsV2ListResponse
func (client *Client) DescribeTrFirewallsV2ListWithOptions(request *DescribeTrFirewallsV2ListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallsV2ListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.FirewallName) {
		query["FirewallName"] = request.FirewallName
	}

	if !dara.IsNil(request.FirewallSwitchStatus) {
		query["FirewallSwitchStatus"] = request.FirewallSwitchStatus
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.RouteMode) {
		query["RouteMode"] = request.RouteMode
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrFirewallsV2List"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrFirewallsV2ListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual private cloud (VPC) firewalls that are created for transit routers.
//
// @param request - DescribeTrFirewallsV2ListRequest
//
// @return DescribeTrFirewallsV2ListResponse
func (client *Client) DescribeTrFirewallsV2List(request *DescribeTrFirewallsV2ListRequest) (_result *DescribeTrFirewallsV2ListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrFirewallsV2ListResponse{}
	_body, _err := client.DescribeTrFirewallsV2ListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the route tables of the VPC firewalls that are created for transit routers.
//
// @param request - DescribeTrFirewallsV2RouteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallsV2RouteListResponse
func (client *Client) DescribeTrFirewallsV2RouteListWithOptions(request *DescribeTrFirewallsV2RouteListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallsV2RouteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TrFirewallRoutePolicyId) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrFirewallsV2RouteList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrFirewallsV2RouteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route tables of the VPC firewalls that are created for transit routers.
//
// @param request - DescribeTrFirewallsV2RouteListRequest
//
// @return DescribeTrFirewallsV2RouteListResponse
func (client *Client) DescribeTrFirewallsV2RouteList(request *DescribeTrFirewallsV2RouteListRequest) (_result *DescribeTrFirewallsV2RouteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrFirewallsV2RouteListResponse{}
	_body, _err := client.DescribeTrFirewallsV2RouteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取未保护漏洞趋势
//
// @param request - DescribeUnprotectedVulnTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUnprotectedVulnTrendResponse
func (client *Client) DescribeUnprotectedVulnTrendWithOptions(request *DescribeUnprotectedVulnTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeUnprotectedVulnTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUnprotectedVulnTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUnprotectedVulnTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取未保护漏洞趋势
//
// @param request - DescribeUnprotectedVulnTrendRequest
//
// @return DescribeUnprotectedVulnTrendResponse
func (client *Client) DescribeUnprotectedVulnTrend(request *DescribeUnprotectedVulnTrendRequest) (_result *DescribeUnprotectedVulnTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUnprotectedVulnTrendResponse{}
	_body, _err := client.DescribeUnprotectedVulnTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户告警配置
//
// @param request - DescribeUserAlarmConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAlarmConfigResponse
func (client *Client) DescribeUserAlarmConfigWithOptions(request *DescribeUserAlarmConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAlarmConfigResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserAlarmConfig"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserAlarmConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户告警配置
//
// @param request - DescribeUserAlarmConfigRequest
//
// @return DescribeUserAlarmConfigResponse
func (client *Client) DescribeUserAlarmConfig(request *DescribeUserAlarmConfigRequest) (_result *DescribeUserAlarmConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserAlarmConfigResponse{}
	_body, _err := client.DescribeUserAlarmConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the traffic of a specified asset that belongs to your Alibaba Cloud account.
//
// @param request - DescribeUserAssetIPTrafficInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAssetIPTrafficInfoResponse
func (client *Client) DescribeUserAssetIPTrafficInfoWithOptions(request *DescribeUserAssetIPTrafficInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAssetIPTrafficInfoResponse, _err error) {
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
		Action:      dara.String("DescribeUserAssetIPTrafficInfo"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserAssetIPTrafficInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the traffic of a specified asset that belongs to your Alibaba Cloud account.
//
// @param request - DescribeUserAssetIPTrafficInfoRequest
//
// @return DescribeUserAssetIPTrafficInfoResponse
func (client *Client) DescribeUserAssetIPTrafficInfo(request *DescribeUserAssetIPTrafficInfoRequest) (_result *DescribeUserAssetIPTrafficInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserAssetIPTrafficInfoResponse{}
	_body, _err := client.DescribeUserAssetIPTrafficInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the edition information about Cloud Firewall.
//
// Description:
//
// You can call this operation to query the edition information about Cloud Firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeUserBuyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserBuyVersionResponse
func (client *Client) DescribeUserBuyVersionWithOptions(request *DescribeUserBuyVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserBuyVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserBuyVersion"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserBuyVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the edition information about Cloud Firewall.
//
// Description:
//
// You can call this operation to query the edition information about Cloud Firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeUserBuyVersionRequest
//
// @return DescribeUserBuyVersionResponse
func (client *Client) DescribeUserBuyVersion(request *DescribeUserBuyVersionRequest) (_result *DescribeUserBuyVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserBuyVersionResponse{}
	_body, _err := client.DescribeUserBuyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户IPS白名单
//
// @param request - DescribeUserIPSWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserIPSWhitelistResponse
func (client *Client) DescribeUserIPSWhitelistWithOptions(request *DescribeUserIPSWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserIPSWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserIPSWhitelist"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserIPSWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户IPS白名单
//
// @param request - DescribeUserIPSWhitelistRequest
//
// @return DescribeUserIPSWhitelistResponse
func (client *Client) DescribeUserIPSWhitelist(request *DescribeUserIPSWhitelistRequest) (_result *DescribeUserIPSWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserIPSWhitelistResponse{}
	_body, _err := client.DescribeUserIPSWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取VPC防火墙通过详情
//
// @param request - DescribeVpcFirewallAccessDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallAccessDetailResponse
func (client *Client) DescribeVpcFirewallAccessDetailWithOptions(request *DescribeVpcFirewallAccessDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallAccessDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetIP) {
		query["AssetIP"] = request.AssetIP
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IPProtocol) {
		query["IPProtocol"] = request.IPProtocol
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PeerAssetIP) {
		query["PeerAssetIP"] = request.PeerAssetIP
	}

	if !dara.IsNil(request.PeerAssetInstanceId) {
		query["PeerAssetInstanceId"] = request.PeerAssetInstanceId
	}

	if !dara.IsNil(request.PeerAssetInstanceName) {
		query["PeerAssetInstanceName"] = request.PeerAssetInstanceName
	}

	if !dara.IsNil(request.PeerVpcId) {
		query["PeerVpcId"] = request.PeerVpcId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallAccessDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallAccessDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取VPC防火墙通过详情
//
// @param request - DescribeVpcFirewallAccessDetailRequest
//
// @return DescribeVpcFirewallAccessDetailResponse
func (client *Client) DescribeVpcFirewallAccessDetail(request *DescribeVpcFirewallAccessDetailRequest) (_result *DescribeVpcFirewallAccessDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallAccessDetailResponse{}
	_body, _err := client.DescribeVpcFirewallAccessDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all policy groups of access control policies that are created for virtual private cloud (VPC) firewalls.
//
// Description:
//
// You can call the DescribeVpcFirewallAclGroupList operation to query the information about all policy groups of access control policies that are created for VPC firewalls.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallAclGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallAclGroupListResponse
func (client *Client) DescribeVpcFirewallAclGroupListWithOptions(request *DescribeVpcFirewallAclGroupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallAclGroupListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FirewallConfigureStatus) {
		query["FirewallConfigureStatus"] = request.FirewallConfigureStatus
	}

	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallAclGroupList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallAclGroupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all policy groups of access control policies that are created for virtual private cloud (VPC) firewalls.
//
// Description:
//
// You can call the DescribeVpcFirewallAclGroupList operation to query the information about all policy groups of access control policies that are created for VPC firewalls.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallAclGroupListRequest
//
// @return DescribeVpcFirewallAclGroupListResponse
func (client *Client) DescribeVpcFirewallAclGroupList(request *DescribeVpcFirewallAclGroupListRequest) (_result *DescribeVpcFirewallAclGroupListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallAclGroupListResponse{}
	_body, _err := client.DescribeVpcFirewallAclGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a virtual private cloud (VPC) firewall. The VPC firewall protects access traffic between a VPC and a network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the DescribeVpcFirewallCenDetail operation to query the details about a VPC firewall. The VPC firewall protects access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallCenDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallCenDetailResponse
func (client *Client) DescribeVpcFirewallCenDetailWithOptions(request *DescribeVpcFirewallCenDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallCenDetailResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NetworkInstanceId) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallCenDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallCenDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a virtual private cloud (VPC) firewall. The VPC firewall protects access traffic between a VPC and a network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the DescribeVpcFirewallCenDetail operation to query the details about a VPC firewall. The VPC firewall protects access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallCenDetailRequest
//
// @return DescribeVpcFirewallCenDetailResponse
func (client *Client) DescribeVpcFirewallCenDetail(request *DescribeVpcFirewallCenDetailRequest) (_result *DescribeVpcFirewallCenDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallCenDetailResponse{}
	_body, _err := client.DescribeVpcFirewallCenDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries virtual private cloud (VPC) firewalls. Each VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the DescribeVpcFirewallCenList operation to query VPC firewalls. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallCenListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallCenListResponse
func (client *Client) DescribeVpcFirewallCenListWithOptions(request *DescribeVpcFirewallCenListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallCenListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FirewallSwitchStatus) {
		query["FirewallSwitchStatus"] = request.FirewallSwitchStatus
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.NetworkInstanceId) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.RouteMode) {
		query["RouteMode"] = request.RouteMode
	}

	if !dara.IsNil(request.TransitRouterType) {
		query["TransitRouterType"] = request.TransitRouterType
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !dara.IsNil(request.VpcFirewallName) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallCenList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallCenListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries virtual private cloud (VPC) firewalls. Each VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the DescribeVpcFirewallCenList operation to query VPC firewalls. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallCenListRequest
//
// @return DescribeVpcFirewallCenListResponse
func (client *Client) DescribeVpcFirewallCenList(request *DescribeVpcFirewallCenListRequest) (_result *DescribeVpcFirewallCenListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallCenListResponse{}
	_body, _err := client.DescribeVpcFirewallCenListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取VPC的CEN列表
//
// @param request - DescribeVpcFirewallCenSummaryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallCenSummaryListResponse
func (client *Client) DescribeVpcFirewallCenSummaryListWithOptions(request *DescribeVpcFirewallCenSummaryListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallCenSummaryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TransitRouterType) {
		query["TransitRouterType"] = request.TransitRouterType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallCenSummaryList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallCenSummaryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取VPC的CEN列表
//
// @param request - DescribeVpcFirewallCenSummaryListRequest
//
// @return DescribeVpcFirewallCenSummaryListResponse
func (client *Client) DescribeVpcFirewallCenSummaryList(request *DescribeVpcFirewallCenSummaryListRequest) (_result *DescribeVpcFirewallCenSummaryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallCenSummaryListResponse{}
	_body, _err := client.DescribeVpcFirewallCenSummaryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the access control policies for a specified virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call the DescribeVpcFirewallControlPolicy operation to query the information about all access control policies that are created for a specified VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallControlPolicyResponse
func (client *Client) DescribeVpcFirewallControlPolicyWithOptions(request *DescribeVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the access control policies for a specified virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call the DescribeVpcFirewallControlPolicy operation to query the information about all access control policies that are created for a specified VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallControlPolicyRequest
//
// @return DescribeVpcFirewallControlPolicyResponse
func (client *Client) DescribeVpcFirewallControlPolicy(request *DescribeVpcFirewallControlPolicyRequest) (_result *DescribeVpcFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallControlPolicyResponse{}
	_body, _err := client.DescribeVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the intrusion prevention configurations of a virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call the DescribeVpcFirewallDefaultIPSConfig operation to query the intrusion prevention configurations of a VPC firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallDefaultIPSConfigResponse
func (client *Client) DescribeVpcFirewallDefaultIPSConfigWithOptions(request *DescribeVpcFirewallDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallDefaultIPSConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallDefaultIPSConfig"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the intrusion prevention configurations of a virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call the DescribeVpcFirewallDefaultIPSConfig operation to query the intrusion prevention configurations of a VPC firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallDefaultIPSConfigRequest
//
// @return DescribeVpcFirewallDefaultIPSConfigResponse
func (client *Client) DescribeVpcFirewallDefaultIPSConfig(request *DescribeVpcFirewallDefaultIPSConfigRequest) (_result *DescribeVpcFirewallDefaultIPSConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.DescribeVpcFirewallDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a virtual private cloud (VPC) firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the DescribeVpcFirewallDetail operation to query the details about a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallDetailResponse
func (client *Client) DescribeVpcFirewallDetailWithOptions(request *DescribeVpcFirewallDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallDetailResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LocalVpcId) {
		query["LocalVpcId"] = request.LocalVpcId
	}

	if !dara.IsNil(request.PeerVpcId) {
		query["PeerVpcId"] = request.PeerVpcId
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a virtual private cloud (VPC) firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the DescribeVpcFirewallDetail operation to query the details about a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallDetailRequest
//
// @return DescribeVpcFirewallDetailResponse
func (client *Client) DescribeVpcFirewallDetail(request *DescribeVpcFirewallDetailRequest) (_result *DescribeVpcFirewallDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallDetailResponse{}
	_body, _err := client.DescribeVpcFirewallDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IPS whitelist of a virtual private cloud (VPC) firewall.
//
// @param request - DescribeVpcFirewallIPSWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallIPSWhitelistResponse
func (client *Client) DescribeVpcFirewallIPSWhitelistWithOptions(request *DescribeVpcFirewallIPSWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallIPSWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallIPSWhitelist"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallIPSWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IPS whitelist of a virtual private cloud (VPC) firewall.
//
// @param request - DescribeVpcFirewallIPSWhitelistRequest
//
// @return DescribeVpcFirewallIPSWhitelistResponse
func (client *Client) DescribeVpcFirewallIPSWhitelist(request *DescribeVpcFirewallIPSWhitelistRequest) (_result *DescribeVpcFirewallIPSWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallIPSWhitelistResponse{}
	_body, _err := client.DescribeVpcFirewallIPSWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about virtual private cloud (VPC) firewalls by page. Each VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the DescribeVpcFirewallList operation to query the details about VPC firewalls by page. Each VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallListResponse
func (client *Client) DescribeVpcFirewallListWithOptions(request *DescribeVpcFirewallListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectSubType) {
		query["ConnectSubType"] = request.ConnectSubType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FirewallSwitchStatus) {
		query["FirewallSwitchStatus"] = request.FirewallSwitchStatus
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PeerUid) {
		query["PeerUid"] = request.PeerUid
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !dara.IsNil(request.VpcFirewallName) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about virtual private cloud (VPC) firewalls by page. Each VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the DescribeVpcFirewallList operation to query the details about VPC firewalls by page. Each VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallListRequest
//
// @return DescribeVpcFirewallListResponse
func (client *Client) DescribeVpcFirewallList(request *DescribeVpcFirewallListRequest) (_result *DescribeVpcFirewallListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallListResponse{}
	_body, _err := client.DescribeVpcFirewallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the priority range of access control policies that are created for a virtual private cloud (VPC) firewall in a specific policy group.
//
// Description:
//
// You can call this operation to query the priority range of access control policies that are created for a VPC firewall in a specific policy group.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallPolicyPriorUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallPolicyPriorUsedResponse
func (client *Client) DescribeVpcFirewallPolicyPriorUsedWithOptions(request *DescribeVpcFirewallPolicyPriorUsedRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallPolicyPriorUsedResponse, _err error) {
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

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallPolicyPriorUsed"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallPolicyPriorUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the priority range of access control policies that are created for a virtual private cloud (VPC) firewall in a specific policy group.
//
// Description:
//
// You can call this operation to query the priority range of access control policies that are created for a VPC firewall in a specific policy group.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVpcFirewallPolicyPriorUsedRequest
//
// @return DescribeVpcFirewallPolicyPriorUsedResponse
func (client *Client) DescribeVpcFirewallPolicyPriorUsed(request *DescribeVpcFirewallPolicyPriorUsedRequest) (_result *DescribeVpcFirewallPolicyPriorUsedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallPolicyPriorUsedResponse{}
	_body, _err := client.DescribeVpcFirewallPolicyPriorUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取VPC防火墙总结信息
//
// @param request - DescribeVpcFirewallSummaryInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallSummaryInfoResponse
func (client *Client) DescribeVpcFirewallSummaryInfoWithOptions(request *DescribeVpcFirewallSummaryInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallSummaryInfoResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallSummaryInfo"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallSummaryInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取VPC防火墙总结信息
//
// @param request - DescribeVpcFirewallSummaryInfoRequest
//
// @return DescribeVpcFirewallSummaryInfoResponse
func (client *Client) DescribeVpcFirewallSummaryInfo(request *DescribeVpcFirewallSummaryInfoRequest) (_result *DescribeVpcFirewallSummaryInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcFirewallSummaryInfoResponse{}
	_body, _err := client.DescribeVpcFirewallSummaryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries virtual private clouds (VPCs).
//
// @param request - DescribeVpcListLiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcListLiteResponse
func (client *Client) DescribeVpcListLiteWithOptions(request *DescribeVpcListLiteRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcListLiteResponse, _err error) {
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

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcName) {
		query["VpcName"] = request.VpcName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcListLite"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcListLiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries virtual private clouds (VPCs).
//
// @param request - DescribeVpcListLiteRequest
//
// @return DescribeVpcListLiteResponse
func (client *Client) DescribeVpcListLite(request *DescribeVpcListLiteRequest) (_result *DescribeVpcListLiteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcListLiteResponse{}
	_body, _err := client.DescribeVpcListLiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries virtual private cloud (VPC) zones.
//
// @param request - DescribeVpcZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcZoneResponse
func (client *Client) DescribeVpcZoneWithOptions(request *DescribeVpcZoneRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcZone"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries virtual private cloud (VPC) zones.
//
// @param request - DescribeVpcZoneRequest
//
// @return DescribeVpcZoneResponse
func (client *Client) DescribeVpcZone(request *DescribeVpcZoneRequest) (_result *DescribeVpcZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcZoneResponse{}
	_body, _err := client.DescribeVpcZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the vulnerabilities that are supported by Cloud Firewall.
//
// @param request - DescribeVulnerabilityProtectedListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVulnerabilityProtectedListResponse
func (client *Client) DescribeVulnerabilityProtectedListWithOptions(request *DescribeVulnerabilityProtectedListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVulnerabilityProtectedListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttackType) {
		query["AttackType"] = request.AttackType
	}

	if !dara.IsNil(request.BuyVersion) {
		query["BuyVersion"] = request.BuyVersion
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleTag) {
		query["RuleTag"] = request.RuleTag
	}

	if !dara.IsNil(request.SortKey) {
		query["SortKey"] = request.SortKey
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	if !dara.IsNil(request.VulnCveName) {
		query["VulnCveName"] = request.VulnCveName
	}

	if !dara.IsNil(request.VulnLevel) {
		query["VulnLevel"] = request.VulnLevel
	}

	if !dara.IsNil(request.VulnResource) {
		query["VulnResource"] = request.VulnResource
	}

	if !dara.IsNil(request.VulnStatus) {
		query["VulnStatus"] = request.VulnStatus
	}

	if !dara.IsNil(request.VulnType) {
		query["VulnType"] = request.VulnType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVulnerabilityProtectedList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVulnerabilityProtectedListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the vulnerabilities that are supported by Cloud Firewall.
//
// @param request - DescribeVulnerabilityProtectedListRequest
//
// @return DescribeVulnerabilityProtectedListResponse
func (client *Client) DescribeVulnerabilityProtectedList(request *DescribeVulnerabilityProtectedListRequest) (_result *DescribeVulnerabilityProtectedListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVulnerabilityProtectedListResponse{}
	_body, _err := client.DescribeVulnerabilityProtectedListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启资产数据泄露保护
//
// @param request - EnableSdlProtectedAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSdlProtectedAssetResponse
func (client *Client) EnableSdlProtectedAssetWithOptions(request *EnableSdlProtectedAssetRequest, runtime *dara.RuntimeOptions) (_result *EnableSdlProtectedAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpList) {
		query["IpList"] = request.IpList
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSdlProtectedAsset"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSdlProtectedAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启资产数据泄露保护
//
// @param request - EnableSdlProtectedAssetRequest
//
// @return EnableSdlProtectedAssetResponse
func (client *Client) EnableSdlProtectedAsset(request *EnableSdlProtectedAssetRequest) (_result *EnableSdlProtectedAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableSdlProtectedAssetResponse{}
	_body, _err := client.EnableSdlProtectedAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载TLS证书
//
// @param request - GetTlsInspectCertificateDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTlsInspectCertificateDownloadUrlResponse
func (client *Client) GetTlsInspectCertificateDownloadUrlWithOptions(request *GetTlsInspectCertificateDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetTlsInspectCertificateDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaCertId) {
		query["CaCertId"] = request.CaCertId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTlsInspectCertificateDownloadUrl"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTlsInspectCertificateDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载TLS证书
//
// @param request - GetTlsInspectCertificateDownloadUrlRequest
//
// @return GetTlsInspectCertificateDownloadUrlResponse
func (client *Client) GetTlsInspectCertificateDownloadUrl(request *GetTlsInspectCertificateDownloadUrlRequest) (_result *GetTlsInspectCertificateDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTlsInspectCertificateDownloadUrlResponse{}
	_body, _err := client.GetTlsInspectCertificateDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询TLS检查证书
//
// @param request - ListTlsInspectCACertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTlsInspectCACertificatesResponse
func (client *Client) ListTlsInspectCACertificatesWithOptions(request *ListTlsInspectCACertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListTlsInspectCACertificatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaCertId) {
		query["CaCertId"] = request.CaCertId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTlsInspectCACertificates"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTlsInspectCACertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询TLS检查证书
//
// @param request - ListTlsInspectCACertificatesRequest
//
// @return ListTlsInspectCACertificatesResponse
func (client *Client) ListTlsInspectCACertificates(request *ListTlsInspectCACertificatesRequest) (_result *ListTlsInspectCACertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTlsInspectCACertificatesResponse{}
	_body, _err := client.ListTlsInspectCACertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the address book that is specified in an access control policy.
//
// Description:
//
// You can call the ModifyAddressBook operation to modify the address book that is configured for access control.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyAddressBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAddressBookResponse
func (client *Client) ModifyAddressBookWithOptions(request *ModifyAddressBookRequest, runtime *dara.RuntimeOptions) (_result *ModifyAddressBookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AckLabels) {
		query["AckLabels"] = request.AckLabels
	}

	if !dara.IsNil(request.AckNamespaces) {
		query["AckNamespaces"] = request.AckNamespaces
	}

	if !dara.IsNil(request.AddressList) {
		query["AddressList"] = request.AddressList
	}

	if !dara.IsNil(request.AutoAddTagEcs) {
		query["AutoAddTagEcs"] = request.AutoAddTagEcs
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupUuid) {
		query["GroupUuid"] = request.GroupUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.TagList) {
		query["TagList"] = request.TagList
	}

	if !dara.IsNil(request.TagRelation) {
		query["TagRelation"] = request.TagRelation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAddressBook"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the address book that is specified in an access control policy.
//
// Description:
//
// You can call the ModifyAddressBook operation to modify the address book that is configured for access control.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyAddressBookRequest
//
// @return ModifyAddressBookResponse
func (client *Client) ModifyAddressBook(request *ModifyAddressBookRequest) (_result *ModifyAddressBookResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAddressBookResponse{}
	_body, _err := client.ModifyAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an access control policy.
//
// Description:
//
// You can call this operation to modify the configurations of an access control policy. The policy allows Cloud Firewall to allow, deny, or monitor the traffic that passes through Cloud Firewall.
//
// ## [](#qps)Limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyControlPolicyResponse
func (client *Client) ModifyControlPolicyWithOptions(request *ModifyControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ApplicationNameList) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestPort) {
		query["DestPort"] = request.DestPort
	}

	if !dara.IsNil(request.DestPortGroup) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !dara.IsNil(request.DestPortType) {
		query["DestPortType"] = request.DestPortType
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainResolveType) {
		query["DomainResolveType"] = request.DomainResolveType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatDays) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !dara.IsNil(request.RepeatEndTime) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !dara.IsNil(request.RepeatStartTime) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an access control policy.
//
// Description:
//
// You can call this operation to modify the configurations of an access control policy. The policy allows Cloud Firewall to allow, deny, or monitor the traffic that passes through Cloud Firewall.
//
// ## [](#qps)Limit
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyControlPolicyRequest
//
// @return ModifyControlPolicyResponse
func (client *Client) ModifyControlPolicy(request *ModifyControlPolicyRequest) (_result *ModifyControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyControlPolicyResponse{}
	_body, _err := client.ModifyControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the priority of an IPv4 access control policy for the Internet firewall. An IPv4 access control policy refers to a policy whose source IP address and destination IP address are IPv4 addresses.
//
// Description:
//
// You can use this operation to modify the priority of an IPv4 access control policy for the Internet firewall. No API operations are provided for you to modify the priority of an IPv6 access control policy for the Internet firewall.
//
// ## [](#qps)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyControlPolicyPositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyControlPolicyPositionResponse
func (client *Client) ModifyControlPolicyPositionWithOptions(request *ModifyControlPolicyPositionRequest, runtime *dara.RuntimeOptions) (_result *ModifyControlPolicyPositionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewOrder) {
		query["NewOrder"] = request.NewOrder
	}

	if !dara.IsNil(request.OldOrder) {
		query["OldOrder"] = request.OldOrder
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyControlPolicyPosition"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyControlPolicyPositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the priority of an IPv4 access control policy for the Internet firewall. An IPv4 access control policy refers to a policy whose source IP address and destination IP address are IPv4 addresses.
//
// Description:
//
// You can use this operation to modify the priority of an IPv4 access control policy for the Internet firewall. No API operations are provided for you to modify the priority of an IPv6 access control policy for the Internet firewall.
//
// ## [](#qps)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyControlPolicyPositionRequest
//
// @return ModifyControlPolicyPositionResponse
func (client *Client) ModifyControlPolicyPosition(request *ModifyControlPolicyPositionRequest) (_result *ModifyControlPolicyPositionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyControlPolicyPositionResponse{}
	_body, _err := client.ModifyControlPolicyPositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the priority of an access control policy.
//
// @param request - ModifyControlPolicyPriorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyControlPolicyPriorityResponse
func (client *Client) ModifyControlPolicyPriorityWithOptions(request *ModifyControlPolicyPriorityRequest, runtime *dara.RuntimeOptions) (_result *ModifyControlPolicyPriorityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyControlPolicyPriority"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyControlPolicyPriorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the priority of an access control policy.
//
// @param request - ModifyControlPolicyPriorityRequest
//
// @return ModifyControlPolicyPriorityResponse
func (client *Client) ModifyControlPolicyPriority(request *ModifyControlPolicyPriorityRequest) (_result *ModifyControlPolicyPriorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyControlPolicyPriorityResponse{}
	_body, _err := client.ModifyControlPolicyPriorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the default configuration of the intrusion prevention system (IPS).
//
// @param request - ModifyDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefaultIPSConfigResponse
func (client *Client) ModifyDefaultIPSConfigWithOptions(request *ModifyDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefaultIPSConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasicRules) {
		query["BasicRules"] = request.BasicRules
	}

	if !dara.IsNil(request.CtiRules) {
		query["CtiRules"] = request.CtiRules
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxSdl) {
		query["MaxSdl"] = request.MaxSdl
	}

	if !dara.IsNil(request.PatchRules) {
		query["PatchRules"] = request.PatchRules
	}

	if !dara.IsNil(request.RuleClass) {
		query["RuleClass"] = request.RuleClass
	}

	if !dara.IsNil(request.RunMode) {
		query["RunMode"] = request.RunMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefaultIPSConfig"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the default configuration of the intrusion prevention system (IPS).
//
// @param request - ModifyDefaultIPSConfigRequest
//
// @return ModifyDefaultIPSConfigResponse
func (client *Client) ModifyDefaultIPSConfig(request *ModifyDefaultIPSConfigRequest) (_result *ModifyDefaultIPSConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefaultIPSConfigResponse{}
	_body, _err := client.ModifyDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改DNS防火墙规则
//
// @param request - ModifyDnsFirewallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDnsFirewallPolicyResponse
func (client *Client) ModifyDnsFirewallPolicyWithOptions(request *ModifyDnsFirewallPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDnsFirewallPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDnsFirewallPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDnsFirewallPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改DNS防火墙规则
//
// @param request - ModifyDnsFirewallPolicyRequest
//
// @return ModifyDnsFirewallPolicyResponse
func (client *Client) ModifyDnsFirewallPolicy(request *ModifyDnsFirewallPolicyRequest) (_result *ModifyDnsFirewallPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDnsFirewallPolicyResponse{}
	_body, _err := client.ModifyDnsFirewallPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of a routing policy.
//
// @param request - ModifyFirewallV2RoutePolicySwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFirewallV2RoutePolicySwitchResponse
func (client *Client) ModifyFirewallV2RoutePolicySwitchWithOptions(request *ModifyFirewallV2RoutePolicySwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyFirewallV2RoutePolicySwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ShouldRecover) {
		query["ShouldRecover"] = request.ShouldRecover
	}

	if !dara.IsNil(request.TrFirewallRoutePolicyId) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	if !dara.IsNil(request.TrFirewallRoutePolicySwitchStatus) {
		query["TrFirewallRoutePolicySwitchStatus"] = request.TrFirewallRoutePolicySwitchStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFirewallV2RoutePolicySwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFirewallV2RoutePolicySwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of a routing policy.
//
// @param request - ModifyFirewallV2RoutePolicySwitchRequest
//
// @return ModifyFirewallV2RoutePolicySwitchResponse
func (client *Client) ModifyFirewallV2RoutePolicySwitch(request *ModifyFirewallV2RoutePolicySwitchRequest) (_result *ModifyFirewallV2RoutePolicySwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFirewallV2RoutePolicySwitchResponse{}
	_body, _err := client.ModifyFirewallV2RoutePolicySwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about members in Cloud Firewall.
//
// Description:
//
// You can call the ModifyInstanceMemberAttributes operation to update the information about members in Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second for each account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyInstanceMemberAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMemberAttributesResponse
func (client *Client) ModifyInstanceMemberAttributesWithOptions(request *ModifyInstanceMemberAttributesRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMemberAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Members) {
		query["Members"] = request.Members
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceMemberAttributes"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceMemberAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about members in Cloud Firewall.
//
// Description:
//
// You can call the ModifyInstanceMemberAttributes operation to update the information about members in Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second for each account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyInstanceMemberAttributesRequest
//
// @return ModifyInstanceMemberAttributesResponse
func (client *Client) ModifyInstanceMemberAttributes(request *ModifyInstanceMemberAttributesRequest) (_result *ModifyInstanceMemberAttributesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceMemberAttributesResponse{}
	_body, _err := client.ModifyInstanceMemberAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改IPS规则为默认
//
// @param request - ModifyIpsRulesToDefaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpsRulesToDefaultResponse
func (client *Client) ModifyIpsRulesToDefaultWithOptions(request *ModifyIpsRulesToDefaultRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpsRulesToDefaultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttackApp) {
		query["AttackApp"] = request.AttackApp
	}

	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIpsRulesToDefault"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIpsRulesToDefaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改IPS规则为默认
//
// @param request - ModifyIpsRulesToDefaultRequest
//
// @return ModifyIpsRulesToDefaultResponse
func (client *Client) ModifyIpsRulesToDefault(request *ModifyIpsRulesToDefaultRequest) (_result *ModifyIpsRulesToDefaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyIpsRulesToDefaultResponse{}
	_body, _err := client.ModifyIpsRulesToDefaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an access control policy that is created for a NAT firewall.
//
// Description:
//
// You can use this operation to modify the configurations of an access control policy. The policy is used to allow, deny, or monitor traffic that reaches a NAT firewall.
//
// @param request - ModifyNatFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNatFirewallControlPolicyResponse
func (client *Client) ModifyNatFirewallControlPolicyWithOptions(request *ModifyNatFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyNatFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.ApplicationNameList) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestPort) {
		query["DestPort"] = request.DestPort
	}

	if !dara.IsNil(request.DestPortGroup) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !dara.IsNil(request.DestPortType) {
		query["DestPortType"] = request.DestPortType
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainResolveType) {
		query["DomainResolveType"] = request.DomainResolveType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatDays) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !dara.IsNil(request.RepeatEndTime) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !dara.IsNil(request.RepeatStartTime) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNatFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNatFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an access control policy that is created for a NAT firewall.
//
// Description:
//
// You can use this operation to modify the configurations of an access control policy. The policy is used to allow, deny, or monitor traffic that reaches a NAT firewall.
//
// @param request - ModifyNatFirewallControlPolicyRequest
//
// @return ModifyNatFirewallControlPolicyResponse
func (client *Client) ModifyNatFirewallControlPolicy(request *ModifyNatFirewallControlPolicyRequest) (_result *ModifyNatFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNatFirewallControlPolicyResponse{}
	_body, _err := client.ModifyNatFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the priority of an access control policy that is created for a NAT firewall.
//
// @param request - ModifyNatFirewallControlPolicyPositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNatFirewallControlPolicyPositionResponse
func (client *Client) ModifyNatFirewallControlPolicyPositionWithOptions(request *ModifyNatFirewallControlPolicyPositionRequest, runtime *dara.RuntimeOptions) (_result *ModifyNatFirewallControlPolicyPositionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.NewOrder) {
		query["NewOrder"] = request.NewOrder
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNatFirewallControlPolicyPosition"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNatFirewallControlPolicyPositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the priority of an access control policy that is created for a NAT firewall.
//
// @param request - ModifyNatFirewallControlPolicyPositionRequest
//
// @return ModifyNatFirewallControlPolicyPositionResponse
func (client *Client) ModifyNatFirewallControlPolicyPosition(request *ModifyNatFirewallControlPolicyPositionRequest) (_result *ModifyNatFirewallControlPolicyPositionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNatFirewallControlPolicyPositionResponse{}
	_body, _err := client.ModifyNatFirewallControlPolicyPositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about an operation on an object group.
//
// @param request - ModifyObjectGroupOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyObjectGroupOperationResponse
func (client *Client) ModifyObjectGroupOperationWithOptions(request *ModifyObjectGroupOperationRequest, runtime *dara.RuntimeOptions) (_result *ModifyObjectGroupOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ObjectList) {
		query["ObjectList"] = request.ObjectList
	}

	if !dara.IsNil(request.ObjectOperation) {
		query["ObjectOperation"] = request.ObjectOperation
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyObjectGroupOperation"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyObjectGroupOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about an operation on an object group.
//
// @param request - ModifyObjectGroupOperationRequest
//
// @return ModifyObjectGroupOperationResponse
func (client *Client) ModifyObjectGroupOperation(request *ModifyObjectGroupOperationRequest) (_result *ModifyObjectGroupOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyObjectGroupOperationResponse{}
	_body, _err := client.ModifyObjectGroupOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the strict mode for an access control policy.
//
// Description:
//
// You can call the ModifyPolicyAdvancedConfig operation to enable or disable the strict mode for an access control policy.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyPolicyAdvancedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyAdvancedConfigResponse
func (client *Client) ModifyPolicyAdvancedConfigWithOptions(request *ModifyPolicyAdvancedConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyAdvancedConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Eips) {
		query["Eips"] = request.Eips
	}

	if !dara.IsNil(request.InternetSwitch) {
		query["InternetSwitch"] = request.InternetSwitch
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolicyAdvancedConfig"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolicyAdvancedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the strict mode for an access control policy.
//
// Description:
//
// You can call the ModifyPolicyAdvancedConfig operation to enable or disable the strict mode for an access control policy.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyPolicyAdvancedConfigRequest
//
// @return ModifyPolicyAdvancedConfigResponse
func (client *Client) ModifyPolicyAdvancedConfig(request *ModifyPolicyAdvancedConfigRequest) (_result *ModifyPolicyAdvancedConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPolicyAdvancedConfigResponse{}
	_body, _err := client.ModifyPolicyAdvancedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改私网DNS终端节点
//
// @param request - ModifyPrivateDnsEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrivateDnsEndpointResponse
func (client *Client) ModifyPrivateDnsEndpointWithOptions(request *ModifyPrivateDnsEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyPrivateDnsEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessInstanceId) {
		query["AccessInstanceId"] = request.AccessInstanceId
	}

	if !dara.IsNil(request.AccessInstanceName) {
		query["AccessInstanceName"] = request.AccessInstanceName
	}

	if !dara.IsNil(request.PrimaryDns) {
		query["PrimaryDns"] = request.PrimaryDns
	}

	if !dara.IsNil(request.PrivateDnsType) {
		query["PrivateDnsType"] = request.PrivateDnsType
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.StandbyDns) {
		query["StandbyDns"] = request.StandbyDns
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPrivateDnsEndpoint"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPrivateDnsEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改私网DNS终端节点
//
// @param request - ModifyPrivateDnsEndpointRequest
//
// @return ModifyPrivateDnsEndpointResponse
func (client *Client) ModifyPrivateDnsEndpoint(request *ModifyPrivateDnsEndpointRequest) (_result *ModifyPrivateDnsEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPrivateDnsEndpointResponse{}
	_body, _err := client.ModifyPrivateDnsEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改敏感数据开关
//
// @param request - ModifySensitiveSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySensitiveSwitchResponse
func (client *Client) ModifySensitiveSwitchWithOptions(request *ModifySensitiveSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifySensitiveSwitchResponse, _err error) {
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

	if !dara.IsNil(request.SensitiveCategory) {
		query["SensitiveCategory"] = request.SensitiveCategory
	}

	if !dara.IsNil(request.SwitchStatus) {
		query["SwitchStatus"] = request.SwitchStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySensitiveSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySensitiveSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改敏感数据开关
//
// @param request - ModifySensitiveSwitchRequest
//
// @return ModifySensitiveSwitchResponse
func (client *Client) ModifySensitiveSwitch(request *ModifySensitiveSwitchRequest) (_result *ModifySensitiveSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySensitiveSwitchResponse{}
	_body, _err := client.ModifySensitiveSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改SLS投递
//
// @param request - ModifySlsDispatchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySlsDispatchStatusResponse
func (client *Client) ModifySlsDispatchStatusWithOptions(request *ModifySlsDispatchStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifySlsDispatchStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DispatchValue) {
		query["DispatchValue"] = request.DispatchValue
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.FilterKeys) {
		query["FilterKeys"] = request.FilterKeys
	}

	if !dara.IsNil(request.NewRegionId) {
		query["NewRegionId"] = request.NewRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySlsDispatchStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySlsDispatchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改SLS投递
//
// @param request - ModifySlsDispatchStatusRequest
//
// @return ModifySlsDispatchStatusResponse
func (client *Client) ModifySlsDispatchStatus(request *ModifySlsDispatchStatusRequest) (_result *ModifySlsDispatchStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySlsDispatchStatusResponse{}
	_body, _err := client.ModifySlsDispatchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改威胁情报配置的信息
//
// @param request - ModifyThreatIntelligenceSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyThreatIntelligenceSwitchResponse
func (client *Client) ModifyThreatIntelligenceSwitchWithOptions(request *ModifyThreatIntelligenceSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyThreatIntelligenceSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryList) {
		query["CategoryList"] = request.CategoryList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyThreatIntelligenceSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyThreatIntelligenceSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改威胁情报配置的信息
//
// @param request - ModifyThreatIntelligenceSwitchRequest
//
// @return ModifyThreatIntelligenceSwitchResponse
func (client *Client) ModifyThreatIntelligenceSwitch(request *ModifyThreatIntelligenceSwitchRequest) (_result *ModifyThreatIntelligenceSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyThreatIntelligenceSwitchResponse{}
	_body, _err := client.ModifyThreatIntelligenceSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a virtual private cloud (VPC) firewall that is created for a transit router.
//
// @param request - ModifyTrFirewallV2ConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrFirewallV2ConfigurationResponse
func (client *Client) ModifyTrFirewallV2ConfigurationWithOptions(request *ModifyTrFirewallV2ConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrFirewallV2ConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.FirewallName) {
		query["FirewallName"] = request.FirewallName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTrFirewallV2Configuration"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTrFirewallV2ConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a virtual private cloud (VPC) firewall that is created for a transit router.
//
// @param request - ModifyTrFirewallV2ConfigurationRequest
//
// @return ModifyTrFirewallV2ConfigurationResponse
func (client *Client) ModifyTrFirewallV2Configuration(request *ModifyTrFirewallV2ConfigurationRequest) (_result *ModifyTrFirewallV2ConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTrFirewallV2ConfigurationResponse{}
	_body, _err := client.ModifyTrFirewallV2ConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the effective scope of the routing policy created for the VPC firewall for a transit router.
//
// @param tmpReq - ModifyTrFirewallV2RoutePolicyScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrFirewallV2RoutePolicyScopeResponse
func (client *Client) ModifyTrFirewallV2RoutePolicyScopeWithOptions(tmpReq *ModifyTrFirewallV2RoutePolicyScopeRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrFirewallV2RoutePolicyScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyTrFirewallV2RoutePolicyScopeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestCandidateList) {
		request.DestCandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestCandidateList, dara.String("DestCandidateList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SrcCandidateList) {
		request.SrcCandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SrcCandidateList, dara.String("SrcCandidateList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DestCandidateListShrink) {
		query["DestCandidateList"] = request.DestCandidateListShrink
	}

	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ShouldRecover) {
		query["ShouldRecover"] = request.ShouldRecover
	}

	if !dara.IsNil(request.SrcCandidateListShrink) {
		query["SrcCandidateList"] = request.SrcCandidateListShrink
	}

	if !dara.IsNil(request.TrFirewallRoutePolicyId) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTrFirewallV2RoutePolicyScope"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTrFirewallV2RoutePolicyScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the effective scope of the routing policy created for the VPC firewall for a transit router.
//
// @param request - ModifyTrFirewallV2RoutePolicyScopeRequest
//
// @return ModifyTrFirewallV2RoutePolicyScopeResponse
func (client *Client) ModifyTrFirewallV2RoutePolicyScope(request *ModifyTrFirewallV2RoutePolicyScopeRequest) (_result *ModifyTrFirewallV2RoutePolicyScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTrFirewallV2RoutePolicyScopeResponse{}
	_body, _err := client.ModifyTrFirewallV2RoutePolicyScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改用户IPS白名单
//
// @param request - ModifyUserIPSWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserIPSWhitelistResponse
func (client *Client) ModifyUserIPSWhitelistWithOptions(request *ModifyUserIPSWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserIPSWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ListType) {
		query["ListType"] = request.ListType
	}

	if !dara.IsNil(request.ListValue) {
		query["ListValue"] = request.ListValue
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.WhiteType) {
		query["WhiteType"] = request.WhiteType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserIPSWhitelist"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserIPSWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改用户IPS白名单
//
// @param request - ModifyUserIPSWhitelistRequest
//
// @return ModifyUserIPSWhitelistResponse
func (client *Client) ModifyUserIPSWhitelist(request *ModifyUserIPSWhitelistRequest) (_result *ModifyUserIPSWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserIPSWhitelistResponse{}
	_body, _err := client.ModifyUserIPSWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改用户日志存储时间
//
// @param request - ModifyUserSlsLogStorageTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserSlsLogStorageTimeResponse
func (client *Client) ModifyUserSlsLogStorageTimeWithOptions(request *ModifyUserSlsLogStorageTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserSlsLogStorageTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.StorageTime) {
		query["StorageTime"] = request.StorageTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserSlsLogStorageTime"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserSlsLogStorageTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改用户日志存储时间
//
// @param request - ModifyUserSlsLogStorageTimeRequest
//
// @return ModifyUserSlsLogStorageTimeResponse
func (client *Client) ModifyUserSlsLogStorageTime(request *ModifyUserSlsLogStorageTimeRequest) (_result *ModifyUserSlsLogStorageTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserSlsLogStorageTimeResponse{}
	_body, _err := client.ModifyUserSlsLogStorageTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改VPC防火墙ACL引擎模式
//
// @param request - ModifyVpcFirewallAclEngineModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallAclEngineModeResponse
func (client *Client) ModifyVpcFirewallAclEngineModeWithOptions(request *ModifyVpcFirewallAclEngineModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallAclEngineModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.StrictMode) {
		query["StrictMode"] = request.StrictMode
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallAclEngineMode"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallAclEngineModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改VPC防火墙ACL引擎模式
//
// @param request - ModifyVpcFirewallAclEngineModeRequest
//
// @return ModifyVpcFirewallAclEngineModeResponse
func (client *Client) ModifyVpcFirewallAclEngineMode(request *ModifyVpcFirewallAclEngineModeRequest) (_result *ModifyVpcFirewallAclEngineModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallAclEngineModeResponse{}
	_body, _err := client.ModifyVpcFirewallAclEngineModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a virtual private cloud (VPC) firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the ModifyVpcFirewallCenConfigure operation to modify the configurations of a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallCenConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallCenConfigureResponse
func (client *Client) ModifyVpcFirewallCenConfigureWithOptions(request *ModifyVpcFirewallCenConfigureRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallCenConfigureResponse, _err error) {
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

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !dara.IsNil(request.VpcFirewallName) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallCenConfigure"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a virtual private cloud (VPC) firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the ModifyVpcFirewallCenConfigure operation to modify the configurations of a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallCenConfigureRequest
//
// @return ModifyVpcFirewallCenConfigureResponse
func (client *Client) ModifyVpcFirewallCenConfigure(request *ModifyVpcFirewallCenConfigureRequest) (_result *ModifyVpcFirewallCenConfigureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallCenConfigureResponse{}
	_body, _err := client.ModifyVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables a virtual private cloud (VPC) firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the ModifyVpcFirewallCenSwitchStatus operation to enable or disable a VPC firewall. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. After you enable the VPC firewall, the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. After you disable the VPC firewall, the VPC firewall no longer protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance.
//
// Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallCenSwitchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallCenSwitchStatusResponse
func (client *Client) ModifyVpcFirewallCenSwitchStatusWithOptions(request *ModifyVpcFirewallCenSwitchStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallCenSwitchStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallSwitch) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallCenSwitchStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallCenSwitchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a virtual private cloud (VPC) firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can call the ModifyVpcFirewallCenSwitchStatus operation to enable or disable a VPC firewall. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. After you enable the VPC firewall, the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. After you disable the VPC firewall, the VPC firewall no longer protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance.
//
// Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallCenSwitchStatusRequest
//
// @return ModifyVpcFirewallCenSwitchStatusResponse
func (client *Client) ModifyVpcFirewallCenSwitchStatus(request *ModifyVpcFirewallCenSwitchStatusRequest) (_result *ModifyVpcFirewallCenSwitchStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallCenSwitchStatusResponse{}
	_body, _err := client.ModifyVpcFirewallCenSwitchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a virtual private cloud (VPC) firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the ModifyVpcFirewallConfigure operation to modify the configurations of a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallConfigureResponse
func (client *Client) ModifyVpcFirewallConfigureWithOptions(request *ModifyVpcFirewallConfigureRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallConfigureResponse, _err error) {
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

	if !dara.IsNil(request.LocalVpcCidrTableList) {
		query["LocalVpcCidrTableList"] = request.LocalVpcCidrTableList
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PeerVpcCidrTableList) {
		query["PeerVpcCidrTableList"] = request.PeerVpcCidrTableList
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !dara.IsNil(request.VpcFirewallName) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallConfigure"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a virtual private cloud (VPC) firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the ModifyVpcFirewallConfigure operation to modify the configurations of a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallConfigureRequest
//
// @return ModifyVpcFirewallConfigureResponse
func (client *Client) ModifyVpcFirewallConfigure(request *ModifyVpcFirewallConfigureRequest) (_result *ModifyVpcFirewallConfigureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallConfigureResponse{}
	_body, _err := client.ModifyVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an access control policy that is created for a virtual private cloud (VPC) firewall in a specified policy group.
//
// Description:
//
// You can call the ModifyVpcFirewallControlPolicy operation to modify the configurations of an access control policy that is created for a VPC firewall in a specified policy group. Different access control policies are used for the VPC firewalls that are used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewalls that are used to protect each Express Connect circuit.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallControlPolicyResponse
func (client *Client) ModifyVpcFirewallControlPolicyWithOptions(request *ModifyVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclAction) {
		query["AclAction"] = request.AclAction
	}

	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ApplicationNameList) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestPort) {
		query["DestPort"] = request.DestPort
	}

	if !dara.IsNil(request.DestPortGroup) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !dara.IsNil(request.DestPortType) {
		query["DestPortType"] = request.DestPortType
	}

	if !dara.IsNil(request.Destination) {
		query["Destination"] = request.Destination
	}

	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.DomainResolveType) {
		query["DomainResolveType"] = request.DomainResolveType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Proto) {
		query["Proto"] = request.Proto
	}

	if !dara.IsNil(request.Release) {
		query["Release"] = request.Release
	}

	if !dara.IsNil(request.RepeatDays) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !dara.IsNil(request.RepeatEndTime) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !dara.IsNil(request.RepeatStartTime) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallControlPolicy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an access control policy that is created for a virtual private cloud (VPC) firewall in a specified policy group.
//
// Description:
//
// You can call the ModifyVpcFirewallControlPolicy operation to modify the configurations of an access control policy that is created for a VPC firewall in a specified policy group. Different access control policies are used for the VPC firewalls that are used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewalls that are used to protect each Express Connect circuit.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallControlPolicyRequest
//
// @return ModifyVpcFirewallControlPolicyResponse
func (client *Client) ModifyVpcFirewallControlPolicy(request *ModifyVpcFirewallControlPolicyRequest) (_result *ModifyVpcFirewallControlPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallControlPolicyResponse{}
	_body, _err := client.ModifyVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the priority of an access control policy that is created for a virtual private cloud (VPC) firewall in a specific policy group.
//
// Description:
//
// You can use this operation to modify the priority of an access control policy that is created for a VPC firewall in a specific policy group.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallControlPolicyPositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallControlPolicyPositionResponse
func (client *Client) ModifyVpcFirewallControlPolicyPositionWithOptions(request *ModifyVpcFirewallControlPolicyPositionRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyPositionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewOrder) {
		query["NewOrder"] = request.NewOrder
	}

	if !dara.IsNil(request.OldOrder) {
		query["OldOrder"] = request.OldOrder
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallControlPolicyPosition"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallControlPolicyPositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the priority of an access control policy that is created for a virtual private cloud (VPC) firewall in a specific policy group.
//
// Description:
//
// You can use this operation to modify the priority of an access control policy that is created for a VPC firewall in a specific policy group.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallControlPolicyPositionRequest
//
// @return ModifyVpcFirewallControlPolicyPositionResponse
func (client *Client) ModifyVpcFirewallControlPolicyPosition(request *ModifyVpcFirewallControlPolicyPositionRequest) (_result *ModifyVpcFirewallControlPolicyPositionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallControlPolicyPositionResponse{}
	_body, _err := client.ModifyVpcFirewallControlPolicyPositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the intrusion prevention configurations of a virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call this operation to modify the intrusion prevention configurations of a VPC firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallDefaultIPSConfigResponse
func (client *Client) ModifyVpcFirewallDefaultIPSConfigWithOptions(request *ModifyVpcFirewallDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallDefaultIPSConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasicRules) {
		query["BasicRules"] = request.BasicRules
	}

	if !dara.IsNil(request.EnableAllPatch) {
		query["EnableAllPatch"] = request.EnableAllPatch
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.RuleClass) {
		query["RuleClass"] = request.RuleClass
	}

	if !dara.IsNil(request.RunMode) {
		query["RunMode"] = request.RunMode
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallDefaultIPSConfig"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the intrusion prevention configurations of a virtual private cloud (VPC) firewall.
//
// Description:
//
// You can call this operation to modify the intrusion prevention configurations of a VPC firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallDefaultIPSConfigRequest
//
// @return ModifyVpcFirewallDefaultIPSConfigResponse
func (client *Client) ModifyVpcFirewallDefaultIPSConfig(request *ModifyVpcFirewallDefaultIPSConfigRequest) (_result *ModifyVpcFirewallDefaultIPSConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.ModifyVpcFirewallDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IPS whitelist of a virtual private cloud (VPC) firewall.
//
// @param request - ModifyVpcFirewallIPSWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallIPSWhitelistResponse
func (client *Client) ModifyVpcFirewallIPSWhitelistWithOptions(request *ModifyVpcFirewallIPSWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallIPSWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.ListType) {
		query["ListType"] = request.ListType
	}

	if !dara.IsNil(request.ListValue) {
		query["ListValue"] = request.ListValue
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !dara.IsNil(request.WhiteType) {
		query["WhiteType"] = request.WhiteType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallIPSWhitelist"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallIPSWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IPS whitelist of a virtual private cloud (VPC) firewall.
//
// @param request - ModifyVpcFirewallIPSWhitelistRequest
//
// @return ModifyVpcFirewallIPSWhitelistResponse
func (client *Client) ModifyVpcFirewallIPSWhitelist(request *ModifyVpcFirewallIPSWhitelistRequest) (_result *ModifyVpcFirewallIPSWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallIPSWhitelistResponse{}
	_body, _err := client.ModifyVpcFirewallIPSWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables a virtual private cloud (VPC) firewall. The VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the ModifyVpcFirewallSwitchStatus operation to enable or disable a VPC firewall. The VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit. After you enable the VPC firewall, the VPC firewall protects access traffic between two VPCs that are connected by using an Express Connect circuit. After you disable the VPC firewall, the VPC firewall no longer protects access traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallSwitchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallSwitchStatusResponse
func (client *Client) ModifyVpcFirewallSwitchStatusWithOptions(request *ModifyVpcFirewallSwitchStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallSwitchStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallSwitch) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcFirewallSwitchStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcFirewallSwitchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a virtual private cloud (VPC) firewall. The VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Description:
//
// You can call the ModifyVpcFirewallSwitchStatus operation to enable or disable a VPC firewall. The VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit. After you enable the VPC firewall, the VPC firewall protects access traffic between two VPCs that are connected by using an Express Connect circuit. After you disable the VPC firewall, the VPC firewall no longer protects access traffic between two VPCs that are connected by using an Express Connect circuit.
//
// Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyVpcFirewallSwitchStatusRequest
//
// @return ModifyVpcFirewallSwitchStatusResponse
func (client *Client) ModifyVpcFirewallSwitchStatus(request *ModifyVpcFirewallSwitchStatusRequest) (_result *ModifyVpcFirewallSwitchStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcFirewallSwitchStatusResponse{}
	_body, _err := client.ModifyVpcFirewallSwitchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Turns off all firewall switches.
//
// Description:
//
// You can call the PutDisableAllFwSwitch operation to turn off all firewall switches.
//
// ## [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PutDisableAllFwSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDisableAllFwSwitchResponse
func (client *Client) PutDisableAllFwSwitchWithOptions(request *PutDisableAllFwSwitchRequest, runtime *dara.RuntimeOptions) (_result *PutDisableAllFwSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutDisableAllFwSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDisableAllFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Turns off all firewall switches.
//
// Description:
//
// You can call the PutDisableAllFwSwitch operation to turn off all firewall switches.
//
// ## [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PutDisableAllFwSwitchRequest
//
// @return PutDisableAllFwSwitchResponse
func (client *Client) PutDisableAllFwSwitch(request *PutDisableAllFwSwitchRequest) (_result *PutDisableAllFwSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutDisableAllFwSwitchResponse{}
	_body, _err := client.PutDisableAllFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disable a firewall for specific assets.
//
// Description:
//
// You can call the PutDisableFwSwitch operation to disable a firewall for specific assets. After you disable the firewall, traffic does not pass through Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PutDisableFwSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDisableFwSwitchResponse
func (client *Client) PutDisableFwSwitchWithOptions(request *PutDisableFwSwitchRequest, runtime *dara.RuntimeOptions) (_result *PutDisableFwSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpaddrList) {
		query["IpaddrList"] = request.IpaddrList
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionList) {
		query["RegionList"] = request.RegionList
	}

	if !dara.IsNil(request.ResourceTypeList) {
		query["ResourceTypeList"] = request.ResourceTypeList
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutDisableFwSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDisableFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disable a firewall for specific assets.
//
// Description:
//
// You can call the PutDisableFwSwitch operation to disable a firewall for specific assets. After you disable the firewall, traffic does not pass through Cloud Firewall.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PutDisableFwSwitchRequest
//
// @return PutDisableFwSwitchResponse
func (client *Client) PutDisableFwSwitch(request *PutDisableFwSwitchRequest) (_result *PutDisableFwSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutDisableFwSwitchResponse{}
	_body, _err := client.PutDisableFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a firewall for all public IP addresses within your Alibaba Cloud account.
//
// Description:
//
// You can call the PutEnableAllFwSwitch operation to enable a firewall for all public IP addresses within your Alibaba Cloud account.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PutEnableAllFwSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEnableAllFwSwitchResponse
func (client *Client) PutEnableAllFwSwitchWithOptions(request *PutEnableAllFwSwitchRequest, runtime *dara.RuntimeOptions) (_result *PutEnableAllFwSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutEnableAllFwSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutEnableAllFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a firewall for all public IP addresses within your Alibaba Cloud account.
//
// Description:
//
// You can call the PutEnableAllFwSwitch operation to enable a firewall for all public IP addresses within your Alibaba Cloud account.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PutEnableAllFwSwitchRequest
//
// @return PutEnableAllFwSwitchResponse
func (client *Client) PutEnableAllFwSwitch(request *PutEnableAllFwSwitchRequest) (_result *PutEnableAllFwSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutEnableAllFwSwitchResponse{}
	_body, _err := client.PutEnableAllFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables firewalls for specific assets.
//
// Description:
//
// You can call this operation to enable a firewall. After you enable a firewall, traffic passes through Cloud Firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to five times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PutEnableFwSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEnableFwSwitchResponse
func (client *Client) PutEnableFwSwitchWithOptions(request *PutEnableFwSwitchRequest, runtime *dara.RuntimeOptions) (_result *PutEnableFwSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpaddrList) {
		query["IpaddrList"] = request.IpaddrList
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionList) {
		query["RegionList"] = request.RegionList
	}

	if !dara.IsNil(request.ResourceTypeList) {
		query["ResourceTypeList"] = request.ResourceTypeList
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutEnableFwSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutEnableFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables firewalls for specific assets.
//
// Description:
//
// You can call this operation to enable a firewall. After you enable a firewall, traffic passes through Cloud Firewall.
//
// ## [](#qps-)Limits
//
// You can call this operation up to five times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - PutEnableFwSwitchRequest
//
// @return PutEnableFwSwitchResponse
func (client *Client) PutEnableFwSwitch(request *PutEnableFwSwitchRequest) (_result *PutEnableFwSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutEnableFwSwitchResponse{}
	_body, _err := client.PutEnableFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 释放已过期的实例
//
// @param request - ReleaseExpiredInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseExpiredInstanceResponse
func (client *Client) ReleaseExpiredInstanceWithOptions(request *ReleaseExpiredInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseExpiredInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseExpiredInstance"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseExpiredInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 释放已过期的实例
//
// @param request - ReleaseExpiredInstanceRequest
//
// @return ReleaseExpiredInstanceResponse
func (client *Client) ReleaseExpiredInstance(request *ReleaseExpiredInstanceRequest) (_result *ReleaseExpiredInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseExpiredInstanceResponse{}
	_body, _err := client.ReleaseExpiredInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - ReleasePostInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleasePostInstanceResponse
func (client *Client) ReleasePostInstanceWithOptions(request *ReleasePostInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleasePostInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleasePostInstance"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleasePostInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - ReleasePostInstanceRequest
//
// @return ReleasePostInstanceResponse
func (client *Client) ReleasePostInstance(request *ReleasePostInstanceRequest) (_result *ReleasePostInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleasePostInstanceResponse{}
	_body, _err := client.ReleasePostInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the number of NAT firewall hits.
//
// @param request - ResetNatFirewallRuleHitCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetNatFirewallRuleHitCountResponse
func (client *Client) ResetNatFirewallRuleHitCountWithOptions(request *ResetNatFirewallRuleHitCountRequest, runtime *dara.RuntimeOptions) (_result *ResetNatFirewallRuleHitCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetNatFirewallRuleHitCount"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetNatFirewallRuleHitCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the number of NAT firewall hits.
//
// @param request - ResetNatFirewallRuleHitCountRequest
//
// @return ResetNatFirewallRuleHitCountResponse
func (client *Client) ResetNatFirewallRuleHitCount(request *ResetNatFirewallRuleHitCountRequest) (_result *ResetNatFirewallRuleHitCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetNatFirewallRuleHitCountResponse{}
	_body, _err := client.ResetNatFirewallRuleHitCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置规则命中数
//
// @param request - ResetRuleHitCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetRuleHitCountResponse
func (client *Client) ResetRuleHitCountWithOptions(request *ResetRuleHitCountRequest, runtime *dara.RuntimeOptions) (_result *ResetRuleHitCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetRuleHitCount"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetRuleHitCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置规则命中数
//
// @param request - ResetRuleHitCountRequest
//
// @return ResetRuleHitCountResponse
func (client *Client) ResetRuleHitCount(request *ResetRuleHitCountRequest) (_result *ResetRuleHitCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetRuleHitCountResponse{}
	_body, _err := client.ResetRuleHitCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears the count on hits of an access control policy that is created for a virtual private cloud (VPC) firewall in a specific policy group.
//
// Description:
//
// You can call the ResetVpcFirewallRuleHitCount operation to clear the count on hits of an access control policy that is created for a VPC firewall in a specific policy group.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ResetVpcFirewallRuleHitCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetVpcFirewallRuleHitCountResponse
func (client *Client) ResetVpcFirewallRuleHitCountWithOptions(request *ResetVpcFirewallRuleHitCountRequest, runtime *dara.RuntimeOptions) (_result *ResetVpcFirewallRuleHitCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclUuid) {
		query["AclUuid"] = request.AclUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetVpcFirewallRuleHitCount"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetVpcFirewallRuleHitCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the count on hits of an access control policy that is created for a virtual private cloud (VPC) firewall in a specific policy group.
//
// Description:
//
// You can call the ResetVpcFirewallRuleHitCount operation to clear the count on hits of an access control policy that is created for a VPC firewall in a specific policy group.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ResetVpcFirewallRuleHitCountRequest
//
// @return ResetVpcFirewallRuleHitCountResponse
func (client *Client) ResetVpcFirewallRuleHitCount(request *ResetVpcFirewallRuleHitCountRequest) (_result *ResetVpcFirewallRuleHitCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetVpcFirewallRuleHitCountResponse{}
	_body, _err := client.ResetVpcFirewallRuleHitCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables a NAT firewall.
//
// @param request - SwitchSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchSecurityProxyResponse
func (client *Client) SwitchSecurityProxyWithOptions(request *SwitchSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *SwitchSecurityProxyResponse, _err error) {
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

	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.Switch) {
		query["Switch"] = request.Switch
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchSecurityProxy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchSecurityProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a NAT firewall.
//
// @param request - SwitchSecurityProxyRequest
//
// @return SwitchSecurityProxyResponse
func (client *Client) SwitchSecurityProxy(request *SwitchSecurityProxyRequest) (_result *SwitchSecurityProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchSecurityProxyResponse{}
	_body, _err := client.SwitchSecurityProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改AI流量分析开启状态
//
// @param request - UpdateAITrafficAnalysisStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAITrafficAnalysisStatusResponse
func (client *Client) UpdateAITrafficAnalysisStatusWithOptions(request *UpdateAITrafficAnalysisStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAITrafficAnalysisStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAITrafficAnalysisStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAITrafficAnalysisStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改AI流量分析开启状态
//
// @param request - UpdateAITrafficAnalysisStatusRequest
//
// @return UpdateAITrafficAnalysisStatusResponse
func (client *Client) UpdateAITrafficAnalysisStatus(request *UpdateAITrafficAnalysisStatusRequest) (_result *UpdateAITrafficAnalysisStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAITrafficAnalysisStatusResponse{}
	_body, _err := client.UpdateAITrafficAnalysisStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改ACK集群连接器
//
// @param request - UpdateAckClusterConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAckClusterConnectorResponse
func (client *Client) UpdateAckClusterConnectorWithOptions(request *UpdateAckClusterConnectorRequest, runtime *dara.RuntimeOptions) (_result *UpdateAckClusterConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorName) {
		query["ConnectorName"] = request.ConnectorName
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAckClusterConnector"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAckClusterConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改ACK集群连接器
//
// @param request - UpdateAckClusterConnectorRequest
//
// @return UpdateAckClusterConnectorResponse
func (client *Client) UpdateAckClusterConnector(request *UpdateAckClusterConnectorRequest) (_result *UpdateAckClusterConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAckClusterConnectorResponse{}
	_body, _err := client.UpdateAckClusterConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改ACL检查状态
//
// @param request - UpdateAclCheckDetailStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAclCheckDetailStatusResponse
func (client *Client) UpdateAclCheckDetailStatusWithOptions(request *UpdateAclCheckDetailStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAclCheckDetailStatusResponse, _err error) {
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAclCheckDetailStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAclCheckDetailStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改ACL检查状态
//
// @param request - UpdateAclCheckDetailStatusRequest
//
// @return UpdateAclCheckDetailStatusResponse
func (client *Client) UpdateAclCheckDetailStatus(request *UpdateAclCheckDetailStatusRequest) (_result *UpdateAclCheckDetailStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAclCheckDetailStatusResponse{}
	_body, _err := client.UpdateAclCheckDetailStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of the NAT firewall feature for Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - UpdatePostpayUserInternetStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostpayUserInternetStatusResponse
func (client *Client) UpdatePostpayUserInternetStatusWithOptions(request *UpdatePostpayUserInternetStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostpayUserInternetStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Operate) {
		query["Operate"] = request.Operate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePostpayUserInternetStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePostpayUserInternetStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of the NAT firewall feature for Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - UpdatePostpayUserInternetStatusRequest
//
// @return UpdatePostpayUserInternetStatusResponse
func (client *Client) UpdatePostpayUserInternetStatus(request *UpdatePostpayUserInternetStatusRequest) (_result *UpdatePostpayUserInternetStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePostpayUserInternetStatusResponse{}
	_body, _err := client.UpdatePostpayUserInternetStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of the NAT Firewall feature for Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - UpdatePostpayUserNatStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostpayUserNatStatusResponse
func (client *Client) UpdatePostpayUserNatStatusWithOptions(request *UpdatePostpayUserNatStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostpayUserNatStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Operate) {
		query["Operate"] = request.Operate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePostpayUserNatStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePostpayUserNatStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of the NAT Firewall feature for Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - UpdatePostpayUserNatStatusRequest
//
// @return UpdatePostpayUserNatStatusResponse
func (client *Client) UpdatePostpayUserNatStatus(request *UpdatePostpayUserNatStatusRequest) (_result *UpdatePostpayUserNatStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePostpayUserNatStatusResponse{}
	_body, _err := client.UpdatePostpayUserNatStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of the virtual private cloud (VPC) Firewall feature for Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - UpdatePostpayUserVpcStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostpayUserVpcStatusResponse
func (client *Client) UpdatePostpayUserVpcStatusWithOptions(request *UpdatePostpayUserVpcStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostpayUserVpcStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Operate) {
		query["Operate"] = request.Operate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePostpayUserVpcStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePostpayUserVpcStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of the virtual private cloud (VPC) Firewall feature for Cloud Firewall that uses the pay-as-you-go billing method.
//
// @param request - UpdatePostpayUserVpcStatusRequest
//
// @return UpdatePostpayUserVpcStatusResponse
func (client *Client) UpdatePostpayUserVpcStatus(request *UpdatePostpayUserVpcStatusRequest) (_result *UpdatePostpayUserVpcStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePostpayUserVpcStatusResponse{}
	_body, _err := client.UpdatePostpayUserVpcStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新安全正向代理
//
// @param request - UpdateSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityProxyResponse
func (client *Client) UpdateSecurityProxyWithOptions(request *UpdateSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityProxyResponse, _err error) {
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

	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.ProxyName) {
		query["ProxyName"] = request.ProxyName
	}

	if !dara.IsNil(request.StrictMode) {
		query["StrictMode"] = request.StrictMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityProxy"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新安全正向代理
//
// @param request - UpdateSecurityProxyRequest
//
// @return UpdateSecurityProxyResponse
func (client *Client) UpdateSecurityProxy(request *UpdateSecurityProxyRequest) (_result *UpdateSecurityProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSecurityProxyResponse{}
	_body, _err := client.UpdateSecurityProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
