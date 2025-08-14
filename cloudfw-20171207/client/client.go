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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a file download task.
//
// @param request - CreateDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDownloadTaskResponse
func (client *Client) CreateDownloadTaskWithOptions(request *CreateDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDownloadTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a NAT firewall.
//
// @param request - CreateSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityProxyResponse
func (client *Client) CreateSecurityProxyWithOptions(request *CreateSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityProxyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.TrimSql) {
		query["trimSql"] = request.TrimSql
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the default intrusion prevention system (IPS) configurations.
//
// @param request - DescribeDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefaultIPSConfigResponse
func (client *Client) DescribeDefaultIPSConfigWithOptions(request *DescribeDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefaultIPSConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the IP addresses that are open to the Internet.
//
// @param request - DescribeInternetOpenIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenIpResponse
func (client *Client) DescribeInternetOpenIpWithOptions(request *DescribeInternetOpenIpRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the trends of Internet traffic.
//
// @param request - DescribeInternetTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetTrafficTrendResponse
func (client *Client) DescribeInternetTrafficTrendWithOptions(request *DescribeInternetTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetTrafficTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries details of NAT firewalls.
//
// @param request - DescribeNatFirewallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallListResponse
func (client *Client) DescribeNatFirewallListWithOptions(request *DescribeNatFirewallListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 概览页-NAT流量趋势
//
// @param request - DescribeNatFirewallTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallTrafficTrendResponse
func (client *Client) DescribeNatFirewallTrafficTrendWithOptions(request *DescribeNatFirewallTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallTrafficTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// # Get details of outgoing destination IPs
//
// @param request - DescribeOutgoingDestinationIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDestinationIPResponse
func (client *Client) DescribeOutgoingDestinationIPWithOptions(request *DescribeOutgoingDestinationIPRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries prefix lists.
//
// @param request - DescribePrefixListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrefixListsResponse
func (client *Client) DescribePrefixListsWithOptions(request *DescribePrefixListsRequest, runtime *dara.RuntimeOptions) (_result *DescribePrefixListsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries information about the transit routers that are associated with a virtual private cloud (VPC) firewall created for a transit router.
//
// @param tmpReq - DescribeTrFirewallPolicyBackUpAssociationListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallPolicyBackUpAssociationListResponse
func (client *Client) DescribeTrFirewallPolicyBackUpAssociationListWithOptions(tmpReq *DescribeTrFirewallPolicyBackUpAssociationListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallPolicyBackUpAssociationListResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about the traffic of a specified asset that belongs to your Alibaba Cloud account.
//
// @param request - DescribeUserAssetIPTrafficInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAssetIPTrafficInfoResponse
func (client *Client) DescribeUserAssetIPTrafficInfoWithOptions(request *DescribeUserAssetIPTrafficInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAssetIPTrafficInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries virtual private clouds (VPCs).
//
// @param request - DescribeVpcListLiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcListLiteResponse
func (client *Client) DescribeVpcListLiteWithOptions(request *DescribeVpcListLiteRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcListLiteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the configuration of a virtual private cloud (VPC) firewall that is created for a transit router.
//
// @param request - ModifyTrFirewallV2ConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrFirewallV2ConfigurationResponse
func (client *Client) ModifyTrFirewallV2ConfigurationWithOptions(request *ModifyTrFirewallV2ConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrFirewallV2ConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
