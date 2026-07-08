// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates an access control backup.
//
// @param request - AddAclBackupDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAclBackupDataResponse
func (client *Client) AddAclBackupDataWithContext(ctx context.Context, request *AddAclBackupDataRequest, runtime *dara.RuntimeOptions) (_result *AddAclBackupDataResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
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
		Action:      dara.String("AddAclBackupData"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAclBackupDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an address book, including IPv4 address books, ECS tag-based address books, IPv6 address books, domain name address books, and ACK address books.
//
// Description:
//
// This operation creates an address book, including IPv4 address books, ECS tag-based address books, IPv6 address books, domain name address books, and ACK address books.
//
// ## Rate limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param tmpReq - AddAddressBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAddressBookResponse
func (client *Client) AddAddressBookWithContext(ctx context.Context, tmpReq *AddAddressBookRequest, runtime *dara.RuntimeOptions) (_result *AddAddressBookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddAddressBookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetMemberUids) {
		request.AssetMemberUidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetMemberUids, dara.String("AssetMemberUids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AssetRegionResourceTypes) {
		request.AssetRegionResourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetRegionResourceTypes, dara.String("AssetRegionResourceTypes"), dara.String("json"))
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

	if !dara.IsNil(request.AssetMemberUidsShrink) {
		query["AssetMemberUids"] = request.AssetMemberUidsShrink
	}

	if !dara.IsNil(request.AssetRegionResourceTypesShrink) {
		query["AssetRegionResourceTypes"] = request.AssetRegionResourceTypesShrink
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an access control policy.
//
// Description:
//
// You can call this operation to create a policy that allows, denies, or monitors traffic that passes through Cloud Firewall.
//
// ## Rate limit
//
// The single-user queries per second (QPS) limit for this operation is 10. If the number of calls per second exceeds the limit, throttling is triggered. Throttling may affect your business. Call this operation as needed.
//
// @param request - AddControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddControlPolicyResponse
func (client *Client) AddControlPolicyWithContext(ctx context.Context, request *AddControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add a DNS firewall ACL
//
// Description:
//
// Use this API to create a policy that allows, denies, or observes traffic passing through a NAT firewall.
//
// @param request - AddDnsFirewallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsFirewallPolicyResponse
func (client *Client) AddDnsFirewallPolicyWithContext(ctx context.Context, request *AddDnsFirewallPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddDnsFirewallPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a real-time domain name resolution task.
//
// @param request - AddDomainResolveRealtimeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainResolveRealtimeTaskResponse
func (client *Client) AddDomainResolveRealtimeTaskWithContext(ctx context.Context, request *AddDomainResolveRealtimeTaskRequest, runtime *dara.RuntimeOptions) (_result *AddDomainResolveRealtimeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomainResolveRealtimeTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainResolveRealtimeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds member accounts to Cloud Firewall.
//
// Description:
//
// This operation is used to add member accounts to Cloud Firewall.
//
// ## QPS limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddInstanceMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddInstanceMembersResponse
func (client *Client) AddInstanceMembersWithContext(ctx context.Context, request *AddInstanceMembersRequest, runtime *dara.RuntimeOptions) (_result *AddInstanceMembersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a private DNS domain name.
//
// Description:
//
// This operation is used to obtain DNS resolution results for a domain name. Currently, only resolution results from Alibaba Cloud DNS are supported. The domain name that you want to query must use Alibaba Cloud DNS before you can obtain its resolution results.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - AddPrivateDnsDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPrivateDnsDomainNameResponse
func (client *Client) AddPrivateDnsDomainNameWithContext(ctx context.Context, request *AddPrivateDnsDomainNameRequest, runtime *dara.RuntimeOptions) (_result *AddPrivateDnsDomainNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI BatchCopyVpcFirewallControlPolicy is deprecated
//
// Summary:
//
// Copies all policies from a source virtual private cloud (VPC) firewall policy group to a destination VPC firewall policy group.
//
// Description:
//
// This operation is used to copy all policies from a source virtual private cloud (VPC) firewall policy group to a destination VPC firewall policy group.
//
// Before performing this operation, back up your policies. For more information, see [policy backup](https://help.aliyun.com/document_detail/170363.html).
//
// After this operation is complete, the policies in the destination VPC firewall policy group are completely replaced with the policies from the source VPC firewall policy group.
//
// The source VPC firewall policy group and the destination VPC firewall policy group must belong to the same Alibaba Cloud account.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. Invoke this operation as appropriate.
//
// @param request - BatchCopyVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCopyVpcFirewallControlPolicyResponse
func (client *Client) BatchCopyVpcFirewallControlPolicyWithContext(ctx context.Context, request *BatchCopyVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *BatchCopyVpcFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch deletes access control policies of a virtual private cloud (VPC) firewall.
//
// @param request - BatchDeleteVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteVpcFirewallControlPolicyResponse
func (client *Client) BatchDeleteVpcFirewallControlPolicyWithContext(ctx context.Context, request *BatchDeleteVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteVpcFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears firewall audit logs.
//
// @param request - ClearLogStoreStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearLogStoreStorageResponse
func (client *Client) ClearLogStoreStorageWithContext(ctx context.Context, request *ClearLogStoreStorageRequest, runtime *dara.RuntimeOptions) (_result *ClearLogStoreStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Site) {
		query["Site"] = request.Site
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearLogStoreStorage"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearLogStoreStorageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ACK cluster connector.
//
// Description:
//
// ## Rate limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. Manage your calls properly.
//
// @param request - CreateAckClusterConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAckClusterConnectorResponse
func (client *Client) CreateAckClusterConnectorWithContext(ctx context.Context, request *CreateAckClusterConnectorRequest, runtime *dara.RuntimeOptions) (_result *CreateAckClusterConnectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ACL check.
//
// @param request - CreateAclCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAclCheckResponse
func (client *Client) CreateAclCheckWithContext(ctx context.Context, request *CreateAclCheckRequest, runtime *dara.RuntimeOptions) (_result *CreateAclCheckResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDownloadTaskResponse
func (client *Client) CreateDownloadTaskWithContext(ctx context.Context, request *CreateDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDownloadTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a synchronization task for Internet assets.
//
// @param request - CreateInstanceSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceSyncTaskResponse
func (client *Client) CreateInstanceSyncTaskWithContext(ctx context.Context, request *CreateInstanceSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceSyncTaskResponse, _err error) {
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
		Action:      dara.String("CreateInstanceSyncTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceSyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an association for private IP traffic tracing with the Intrusion Prevention System (IPS).
//
// @param request - CreateIpsPrivateAssocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpsPrivateAssocResponse
func (client *Client) CreateIpsPrivateAssocWithContext(ctx context.Context, request *CreateIpsPrivateAssocRequest, runtime *dara.RuntimeOptions) (_result *CreateIpsPrivateAssocResponse, _err error) {
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
		Action:      dara.String("CreateIpsPrivateAssoc"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpsPrivateAssocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add an access control policy to a NAT firewall.
//
// Description:
//
// This API creates a policy to allow, deny, or observe traffic through the NAT Firewall.
//
// @param request - CreateNatFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNatFirewallControlPolicyResponse
func (client *Client) CreateNatFirewallControlPolicyWithContext(ctx context.Context, request *CreateNatFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateNatFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dry run for a NAT firewall.
//
// Description:
//
// Creates a policy that allows, denies, or monitors traffic that passes through a NAT firewall.
//
// @param request - CreateNatFirewallPreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNatFirewallPreCheckResponse
func (client *Client) CreateNatFirewallPreCheckWithContext(ctx context.Context, request *CreateNatFirewallPreCheckRequest, runtime *dara.RuntimeOptions) (_result *CreateNatFirewallPreCheckResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an asset synchronization task for a NAT firewall.
//
// @param request - CreateNatFirewallSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNatFirewallSyncTaskResponse
func (client *Client) CreateNatFirewallSyncTaskWithContext(ctx context.Context, request *CreateNatFirewallSyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateNatFirewallSyncTaskResponse, _err error) {
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
		Action:      dara.String("CreateNatFirewallSyncTask"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNatFirewallSyncTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a private DNS endpoint
//
// Description:
//
// Creates a private DNS endpoint for traffic that passes through the NAT Firewall to allow, deny, or monitor the traffic.
//
// @param request - CreatePrivateDnsEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateDnsEndpointResponse
func (client *Client) CreatePrivateDnsEndpointWithContext(ctx context.Context, request *CreatePrivateDnsEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateDnsEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityProxyResponse
func (client *Client) CreateSecurityProxyWithContext(ctx context.Context, request *CreateSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityProxyResponse, _err error) {
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

	if !dara.IsNil(request.FwVswitchZoneId) {
		query["FwVswitchZoneId"] = request.FwVswitchZoneId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a log delivery configuration for Cloud Firewall to Simple Log Service (SLS).
//
// @param request - CreateSlsLogDispatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSlsLogDispatchResponse
func (client *Client) CreateSlsLogDispatchWithContext(ctx context.Context, request *CreateSlsLogDispatchRequest, runtime *dara.RuntimeOptions) (_result *CreateSlsLogDispatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a VPC firewall for a transit router.
//
// @param request - CreateTrFirewallV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrFirewallV2Response
func (client *Client) CreateTrFirewallV2WithContext(ctx context.Context, request *CreateTrFirewallV2Request, runtime *dara.RuntimeOptions) (_result *CreateTrFirewallV2Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a TR firewall routing rule.
//
// @param tmpReq - CreateTrFirewallV2RoutePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrFirewallV2RoutePolicyResponse
func (client *Client) CreateTrFirewallV2RoutePolicyWithContext(ctx context.Context, tmpReq *CreateTrFirewallV2RoutePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateTrFirewallV2RoutePolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) firewall to protect traffic between network instances in a Cloud Enterprise Network (CEN) instance and a specified VPC.
//
// Description:
//
// This operation is used to create a VPC firewall for VPC-connected instances in a CEN instance. The virtual private cloud (VPC) firewall protects traffic between network instances (including VPCs, virtual border routers (VBRs), and Cloud Connect Networks (CCNs)) in the CEN instance and a specified VPC. The VPC firewall does not protect traffic between VBRs, between CCNs, or between VBRs and CCNs. For more information, see [VPC firewall limits](https://help.aliyun.com/document_detail/172295.html).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls per second exceeds the limit, throttling is triggered. Throttling may affect your business. Invoke this operation within the limit.
//
// @param request - CreateVpcFirewallCenConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallCenConfigureResponse
func (client *Client) CreateVpcFirewallCenConfigureWithContext(ctx context.Context, request *CreateVpcFirewallCenConfigureRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallCenConfigureResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateVpcFirewallCenManualConfigure is deprecated
//
// Summary:
//
// Manually creates a VPC border firewall.
//
// Description:
//
// This operation creates a VPC border firewall for a VPC within a Cloud Enterprise Network (CEN) instance. The VPC border firewall protects traffic between the specified VPC and other network instances that are connected to the CEN instance. These network instances include virtual private clouds (VPCs), virtual border routers (VBRs), and Cloud Connect Network (CCN) instances. The VPC border firewall does not protect traffic between VBRs, between CCN instances, or between VBRs and CCN instances. For more information, see [VPC border firewall limits](https://help.aliyun.com/document_detail/172295.html).
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, API calls are throttled. This can affect your business operations. We recommend that you adhere to this limit.
//
// @param request - CreateVpcFirewallCenManualConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallCenManualConfigureResponse
func (client *Client) CreateVpcFirewallCenManualConfigureWithContext(ctx context.Context, request *CreateVpcFirewallCenManualConfigureRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallCenManualConfigureResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) firewall that protects traffic between two VPCs connected through Express Connect.
//
// Description:
//
// This operation is used to create a VPC firewall. This virtual private cloud (VPC) firewall protects traffic between two VPCs connected through Express Connect. This VPC firewall does not support protection for cross-region traffic, cross-account traffic, or traffic between a VPC and a virtual border router (VBR). For more information, see [VPC firewall limits](https://help.aliyun.com/document_detail/172295.html).
//
// ### Rate limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API invocations are throttled, which may affect your business. Manage your invocations appropriately.
//
// @param request - CreateVpcFirewallConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallConfigureResponse
func (client *Client) CreateVpcFirewallConfigureWithContext(ctx context.Context, request *CreateVpcFirewallConfigureRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallConfigureResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an access control policy to a specified VPC firewall policy group.
//
// Description:
//
// This operation is used to add an access control policy to a specified virtual private cloud (VPC) firewall policy group. Different access control policies are used when a VPC firewall protects traffic between two VPCs connected through Cloud Enterprise Network (CEN) or traffic between two VPCs connected through Express Connect.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. Invoke this operation properly.
//
// @param request - CreateVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallControlPolicyResponse
func (client *Client) CreateVpcFirewallControlPolicyWithContext(ctx context.Context, request *CreateVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a precheck task before you create a VPC firewall.
//
// Description:
//
// This operation creates a policy to accept, deny, or monitor traffic that passes through a NAT firewall.
//
// @param request - CreateVpcFirewallPrecheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallPrecheckResponse
func (client *Client) CreateVpcFirewallPrecheckWithContext(ctx context.Context, request *CreateVpcFirewallPrecheckRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallPrecheckResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a sync task for VPC firewall assets.
//
// Description:
//
// This operation creates a VPC firewall that protects traffic between two VPCs connected by an Express Connect circuit. The VPC firewall does not protect cross-region traffic, cross-account traffic, or traffic between a VPC and a Virtual Border Router (VBR). For more information, see [Limits on VPC firewalls](https://help.aliyun.com/document_detail/172295.html).
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, your API calls will be throttled. Plan your calls accordingly.
//
// @param request - CreateVpcFirewallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcFirewallTaskResponse
func (client *Client) CreateVpcFirewallTaskWithContext(ctx context.Context, request *CreateVpcFirewallTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcFirewallTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ACK cluster connector.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DeleteAckClusterConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAckClusterConnectorResponse
func (client *Client) DeleteAckClusterConnectorWithContext(ctx context.Context, request *DeleteAckClusterConnectorRequest, runtime *dara.RuntimeOptions) (_result *DeleteAckClusterConnectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ACL backup.
//
// Description:
//
// This operation deletes a backup of an access control address book.
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) per user. Calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - DeleteAclBackupDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAclBackupDataResponse
func (client *Client) DeleteAclBackupDataWithContext(ctx context.Context, request *DeleteAclBackupDataRequest, runtime *dara.RuntimeOptions) (_result *DeleteAclBackupDataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address book.
//
// Description:
//
// This operation is used to delete an address book from access control.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteAddressBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddressBookResponse
func (client *Client) DeleteAddressBookWithContext(ctx context.Context, request *DeleteAddressBookRequest, runtime *dara.RuntimeOptions) (_result *DeleteAddressBookResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This operation is used to delete an access control policy whose traffic direction is inbound or outbound.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, throttling is triggered, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlPolicyResponse
func (client *Client) DeleteControlPolicyWithContext(ctx context.Context, request *DeleteControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteControlPolicyTemplateResponse
func (client *Client) DeleteControlPolicyTemplateWithContext(ctx context.Context, request *DeleteControlPolicyTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteControlPolicyTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DNS firewall policy.
//
// Description:
//
// You can call this operation to delete a DNS firewall policy.
//
// @param request - DeleteDnsFirewallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnsFirewallPolicyResponse
func (client *Client) DeleteDnsFirewallPolicyWithContext(ctx context.Context, request *DeleteDnsFirewallPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnsFirewallPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file download task.
//
// Description:
//
// Calling this operation immediately deletes the file download task and the downloaded file.
//
//	Danger: The delete operation deletes the corresponding task and file. **The file can no longer be downloaded by using the existing download link. This operation is irreversible. Proceed with caution.**.
//
// @param request - DeleteDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDownloadTaskResponse
func (client *Client) DeleteDownloadTaskWithContext(ctx context.Context, request *DeleteDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteDownloadTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a routing policy for a VPC firewall for a transit router.
//
// @param request - DeleteFirewallV2RoutePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFirewallV2RoutePoliciesResponse
func (client *Client) DeleteFirewallV2RoutePoliciesWithContext(ctx context.Context, request *DeleteFirewallV2RoutePoliciesRequest, runtime *dara.RuntimeOptions) (_result *DeleteFirewallV2RoutePoliciesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes Cloud Firewall member accounts.
//
// Description:
//
// You can delete up to 20 Cloud Firewall member accounts in a single call. Separate the UIDs of multiple member accounts with commas (,). After a member account is deleted, Cloud Firewall can no longer access the cloud resources of that account. Use this operation with caution. Before deleting member accounts, call the [DescribeInstanceMembers](https://help.aliyun.com/document_detail/271704.html) operation to retrieve information about the member accounts.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed the limit, API calls are throttled. This can affect your business operations. Plan your calls accordingly.
//
// @param request - DeleteInstanceMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceMembersResponse
func (client *Client) DeleteInstanceMembersWithContext(ctx context.Context, request *DeleteInstanceMembersRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceMembersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a private network association for an IPS.
//
// @param request - DeleteIpsPrivateAssocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpsPrivateAssocResponse
func (client *Client) DeleteIpsPrivateAssocWithContext(ctx context.Context, request *DeleteIpsPrivateAssocRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpsPrivateAssocResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a NAT firewall access control policy.
//
// Description:
//
// This operation is used to delete an access control policy for outbound traffic of a NAT firewall.
//
// @param request - DeleteNatFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNatFirewallControlPolicyResponse
func (client *Client) DeleteNatFirewallControlPolicyWithContext(ctx context.Context, request *DeleteNatFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteNatFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a batch of NAT firewall policies.
//
// @param request - DeleteNatFirewallControlPolicyBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNatFirewallControlPolicyBatchResponse
func (client *Client) DeleteNatFirewallControlPolicyBatchWithContext(ctx context.Context, request *DeleteNatFirewallControlPolicyBatchRequest, runtime *dara.RuntimeOptions) (_result *DeleteNatFirewallControlPolicyBatchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes all private DNS domain names.
//
// Description:
//
// This API call is used to delete all private domain names.
//
// ## QPS limit
//
// Each user is limited to 10 queries per second (QPS) for this API call. If you exceed this limit, API calls are throttled, which may impact your business. We recommend that you plan your API calls accordingly.
//
// @param request - DeletePrivateDnsAllDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateDnsAllDomainNameResponse
func (client *Client) DeletePrivateDnsAllDomainNameWithContext(ctx context.Context, request *DeletePrivateDnsAllDomainNameRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateDnsAllDomainNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete private DNS domain names
//
// Description:
//
// Deletes domain names that require private DNS resolution.
//
// @param request - DeletePrivateDnsDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateDnsDomainNameResponse
func (client *Client) DeletePrivateDnsDomainNameWithContext(ctx context.Context, request *DeletePrivateDnsDomainNameRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateDnsDomainNameResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a private DNS endpoint.
//
// Description:
//
// You can use this operation to create a policy that allows, denies, or monitors traffic that passes through a NAT firewall.
//
// @param request - DeletePrivateDnsEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateDnsEndpointResponse
func (client *Client) DeletePrivateDnsEndpointWithContext(ctx context.Context, request *DeletePrivateDnsEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateDnsEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified NAT firewall.
//
// @param request - DeleteSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityProxyResponse
func (client *Client) DeleteSecurityProxyWithContext(ctx context.Context, request *DeleteSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityProxyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VPC firewall for a transit router.
//
// @param request - DeleteTrFirewallV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrFirewallV2Response
func (client *Client) DeleteTrFirewallV2WithContext(ctx context.Context, request *DeleteTrFirewallV2Request, runtime *dara.RuntimeOptions) (_result *DeleteTrFirewallV2Response, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VPC firewall that protects traffic between network instances in a Cloud Enterprise Network (CEN) and a specified VPC.
//
// Description:
//
// This operation deletes a VPC firewall. The VPC firewall protects traffic between network instances (including VPCs, virtual border routers (VBRs), and Cloud Connect Networks (CCNs)) in a Cloud Enterprise Network (CEN) and a specified VPC.
//
// Before calling this operation, call [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) to create a VPC firewall.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DeleteVpcFirewallCenConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcFirewallCenConfigureResponse
func (client *Client) DeleteVpcFirewallCenConfigureWithContext(ctx context.Context, request *DeleteVpcFirewallCenConfigureRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcFirewallCenConfigureResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual private cloud (VPC) firewall that protects traffic between two VPCs connected through Express Connect.
//
// Description:
//
// This operation is used to delete a virtual private cloud (VPC) firewall that protects traffic between two VPCs connected through Express Connect.
//
// Before you invoke this operation, you must have already created a VPC firewall by invoking the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## Rate limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - DeleteVpcFirewallConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcFirewallConfigureResponse
func (client *Client) DeleteVpcFirewallConfigureWithContext(ctx context.Context, request *DeleteVpcFirewallConfigureRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcFirewallConfigureResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control policy from a specified VPC firewall policy group.
//
// Description:
//
// This operation is used to delete an access control policy from a specified VPC firewall policy group. The VPC firewall instances that protect Cloud Enterprise Network (CEN) instances and the VPC firewall instances that protect Express Connect circuits use different access control policies.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. Call this operation as appropriate.
//
// @param request - DeleteVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcFirewallControlPolicyResponse
func (client *Client) DeleteVpcFirewallControlPolicyWithContext(ctx context.Context, request *DeleteVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend chart of Internet access control interceptions.
//
// @param request - DescribeACLProtectTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeACLProtectTrendResponse
func (client *Client) DescribeACLProtectTrendWithContext(ctx context.Context, request *DescribeACLProtectTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeACLProtectTrendResponse, _err error) {
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

	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of regions for synchronization nodes.
//
// @param request - DescribeAccessInstanceRegionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceRegionListResponse
func (client *Client) DescribeAccessInstanceRegionListWithContext(ctx context.Context, request *DescribeAccessInstanceRegionListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceRegionListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress of a synchronization task on a node.
//
// @param request - DescribeAccessInstanceTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceTaskResponse
func (client *Client) DescribeAccessInstanceTaskWithContext(ctx context.Context, request *DescribeAccessInstanceTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the vSwitches for synchronization nodes.
//
// @param request - DescribeAccessInstanceVSwitchListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceVSwitchListResponse
func (client *Client) DescribeAccessInstanceVSwitchListWithContext(ctx context.Context, request *DescribeAccessInstanceVSwitchListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceVSwitchListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the VPCs associated with synchronization nodes.
//
// @param request - DescribeAccessInstanceVpcListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceVpcListResponse
func (client *Client) DescribeAccessInstanceVpcListWithContext(ctx context.Context, request *DescribeAccessInstanceVpcListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceVpcListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a list of available zones for access instances.
//
// @param request - DescribeAccessInstanceZoneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessInstanceZoneListResponse
func (client *Client) DescribeAccessInstanceZoneListWithContext(ctx context.Context, request *DescribeAccessInstanceZoneListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessInstanceZoneListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified ACK cluster connector.
//
// @param request - DescribeAckClusterConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClusterConnectorResponse
func (client *Client) DescribeAckClusterConnectorWithContext(ctx context.Context, request *DescribeAckClusterConnectorRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClusterConnectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ACK cluster connectors in batches.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. Manage your calls properly.
//
// @param request - DescribeAckClusterConnectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClusterConnectorsResponse
func (client *Client) DescribeAckClusterConnectorsWithContext(ctx context.Context, request *DescribeAckClusterConnectorsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClusterConnectorsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the namespaces in an Alibaba Cloud Container Service for Kubernetes (ACK) cluster.
//
// @param request - DescribeAckClusterNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClusterNamespacesResponse
func (client *Client) DescribeAckClusterNamespacesWithContext(ctx context.Context, request *DescribeAckClusterNamespacesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClusterNamespacesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the labels within an Alibaba Cloud Container Service for Kubernetes (ACK) cluster.
//
// @param request - DescribeAckClusterPodLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClusterPodLabelsResponse
func (client *Client) DescribeAckClusterPodLabelsWithContext(ctx context.Context, request *DescribeAckClusterPodLabelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClusterPodLabelsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries for Container Service for Kubernetes (ACK) clusters based on specified conditions, such as cluster type and specifications.
//
// @param request - DescribeAckClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAckClustersResponse
func (client *Client) DescribeAckClustersWithContext(ctx context.Context, request *DescribeAckClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeAckClustersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries multiple ACL applications.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 requests per second. Calls that exceed this limit are throttled, which may impact your business.
//
// @param request - DescribeAclAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclAppsResponse
func (client *Client) DescribeAclAppsWithContext(ctx context.Context, request *DescribeAclAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclAppsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ACL backups.
//
// @param request - DescribeAclBackupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclBackupListResponse
func (client *Client) DescribeAclBackupListWithContext(ctx context.Context, request *DescribeAclBackupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclBackupListResponse, _err error) {
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

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAclBackupList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclBackupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an Access Control List (ACL) check.
//
// Description:
//
// ## QPS limit
//
// This API is limited to 10 queries per second (QPS) per user. Calls exceeding this limit are throttled.
//
// @param request - DescribeAclCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclCheckResponse
func (client *Client) DescribeAclCheckWithContext(ctx context.Context, request *DescribeAclCheckRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclCheckResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota for access control list (ACL) checks.
//
// Description:
//
// ## QPS limits
//
// Each user can make up to 10 queries per second (QPS). If you exceed this limit, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - DescribeAclCheckQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclCheckQuotaResponse
func (client *Client) DescribeAclCheckQuotaWithContext(ctx context.Context, request *DescribeAclCheckQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclCheckQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Access Control List (ACL) checks in batches.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If this limit is exceeded, your API calls are throttled. This may affect your business. We recommend that you plan your calls accordingly.
//
// @param request - DescribeAclChecksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclChecksResponse
func (client *Client) DescribeAclChecksWithContext(ctx context.Context, request *DescribeAclChecksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclChecksResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the total number of access control policy configurations.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeAclRuleCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclRuleCountResponse
func (client *Client) DescribeAclRuleCountWithContext(ctx context.Context, request *DescribeAclRuleCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclRuleCountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the ACL whitelist.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - DescribeAclWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclWhitelistResponse
func (client *Client) DescribeAclWhitelistWithContext(ctx context.Context, request *DescribeAclWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclWhitelistResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries address books in batches.
//
// Description:
//
// This operation is used to query the details of access control policy address books.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param tmpReq - DescribeAddressBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAddressBookResponse
func (client *Client) DescribeAddressBookWithContext(ctx context.Context, tmpReq *DescribeAddressBookRequest, runtime *dara.RuntimeOptions) (_result *DescribeAddressBookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeAddressBookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetMemberUids) {
		request.AssetMemberUidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetMemberUids, dara.String("AssetMemberUids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetMemberUidsShrink) {
		query["AssetMemberUids"] = request.AssetMemberUidsShrink
	}

	if !dara.IsNil(request.ContainPort) {
		query["ContainPort"] = request.ContainPort
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.GroupUuid) {
		query["GroupUuid"] = request.GroupUuid
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about assets protected by Cloud Firewall.
//
// Description:
//
// This API is generally used to query information about assets protected by Cloud Firewall with pagination.
//
// ## QPS limit
//
// The single-user QPS limit for this API is 10 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please make calls appropriately.
//
// @param request - DescribeAssetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetListResponse
func (client *Client) DescribeAssetListWithContext(ctx context.Context, request *DescribeAssetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of asset risk levels.
//
// @param request - DescribeAssetRiskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetRiskListResponse
func (client *Client) DescribeAssetRiskListWithContext(ctx context.Context, request *DescribeAssetRiskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetRiskListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics information of assets protected by Cloud Firewall.
//
// @param request - DescribeAssetStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetStatisticResponse
func (client *Client) DescribeAssetStatisticWithContext(ctx context.Context, request *DescribeAssetStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetStatisticResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of attack categories.
//
// Description:
//
// This operation is generally used for paging query of information about assets protected by Cloud Firewall.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Invoke this operation at an appropriate frequency.
//
// @param request - DescribeAttackAppCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAttackAppCategoryResponse
func (client *Client) DescribeAttackAppCategoryWithContext(ctx context.Context, request *DescribeAttackAppCategoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeAttackAppCategoryResponse, _err error) {
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
		Action:      dara.String("DescribeAttackAppCategory"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAttackAppCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of log delivery.
//
// @param request - DescribeBatchSlsDispatchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBatchSlsDispatchStatusResponse
func (client *Client) DescribeBatchSlsDispatchStatusWithContext(ctx context.Context, request *DescribeBatchSlsDispatchStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeBatchSlsDispatchStatusResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBatchSlsDispatchStatus"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBatchSlsDispatchStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeCfwRiskLevelSummary is deprecated
//
// Summary:
//
// Queries a summary of threat levels for Cloud Firewall.
//
// @param request - DescribeCfwRiskLevelSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCfwRiskLevelSummaryResponse
func (client *Client) DescribeCfwRiskLevelSummaryWithContext(ctx context.Context, request *DescribeCfwRiskLevelSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeCfwRiskLevelSummaryResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries outbound IPs.
//
// @param request - DescribeConfiguredDestinationIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfiguredDestinationIPResponse
func (client *Client) DescribeConfiguredDestinationIPWithContext(ctx context.Context, request *DescribeConfiguredDestinationIPRequest, runtime *dara.RuntimeOptions) (_result *DescribeConfiguredDestinationIPResponse, _err error) {
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

	if !dara.IsNil(request.DestinationIP) {
		query["DestinationIP"] = request.DestinationIP
	}

	if !dara.IsNil(request.DestinationISP) {
		query["DestinationISP"] = request.DestinationISP
	}

	if !dara.IsNil(request.DestinationRegion) {
		query["DestinationRegion"] = request.DestinationRegion
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
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
		Action:      dara.String("DescribeConfiguredDestinationIP"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConfiguredDestinationIPResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the active outbound domain names.
//
// @param request - DescribeConfiguredDomainNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfiguredDomainNamesResponse
func (client *Client) DescribeConfiguredDomainNamesWithContext(ctx context.Context, request *DescribeConfiguredDomainNamesRequest, runtime *dara.RuntimeOptions) (_result *DescribeConfiguredDomainNamesResponse, _err error) {
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

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
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
		Action:      dara.String("DescribeConfiguredDomainNames"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConfiguredDomainNamesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about all access control policies.
//
// Description:
//
// This operation performs a paged query for information about access control policies.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for a single user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeControlPolicyResponse
func (client *Client) DescribeControlPolicyWithContext(ctx context.Context, request *DescribeControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain name resolution results of an access control policy.
//
// @param request - DescribeControlPolicyDomainResolveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeControlPolicyDomainResolveResponse
func (client *Client) DescribeControlPolicyDomainResolveWithContext(ctx context.Context, request *DescribeControlPolicyDomainResolveRequest, runtime *dara.RuntimeOptions) (_result *DescribeControlPolicyDomainResolveResponse, _err error) {
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
		Action:      dara.String("DescribeControlPolicyDomainResolve"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeControlPolicyDomainResolveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of created NAT firewalls.
//
// @param request - DescribeCreatedNatFirewallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCreatedNatFirewallResponse
func (client *Client) DescribeCreatedNatFirewallWithContext(ctx context.Context, request *DescribeCreatedNatFirewallRequest, runtime *dara.RuntimeOptions) (_result *DescribeCreatedNatFirewallResponse, _err error) {
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
		Action:      dara.String("DescribeCreatedNatFirewall"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCreatedNatFirewallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides Intrusion Prevention System (IPS) protection for internet traffic.
//
// @param request - DescribeDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefaultIPSConfigResponse
func (client *Client) DescribeDefaultIPSConfigWithContext(ctx context.Context, request *DescribeDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefaultIPSConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of access control lists (ACLs) for the DNS firewall.
//
// @param request - DescribeDnsFirewallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsFirewallPolicyResponse
func (client *Client) DescribeDnsFirewallPolicyWithContext(ctx context.Context, request *DescribeDnsFirewallPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsFirewallPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Retrieves the Domain Name System (DNS) resolution results for a domain name.
//
// Description:
//
// This operation retrieves the DNS resolution result for a domain name. You can retrieve resolution results only for domain names that use Alibaba Cloud DNS.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed this limit, your API calls are throttled, which may impact your business. Call this operation at a reasonable rate to avoid throttling.
//
// @param request - DescribeDomainResolveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainResolveResponse
func (client *Client) DescribeDomainResolveWithContext(ctx context.Context, request *DescribeDomainResolveRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainResolveResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information and download URLs of file download tasks.
//
// @param request - DescribeDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadTaskResponse
func (client *Client) DescribeDownloadTaskWithContext(ctx context.Context, request *DescribeDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types of download tasks. The returned types correspond to the TaskType field in other download-related API operations.
//
// @param request - DescribeDownloadTaskTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadTaskTypeResponse
func (client *Client) DescribeDownloadTaskTypeWithContext(ctx context.Context, request *DescribeDownloadTaskTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadTaskTypeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the overall firewall interception trend, including Internet, VPC, and NAT traffic.
//
// @param request - DescribeFirewallDropTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallDropTrendResponse
func (client *Client) DescribeFirewallDropTrendWithContext(ctx context.Context, request *DescribeFirewallDropTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeFirewallDropTrendResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFirewallDropTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFirewallDropTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a firewall task.
//
// Description:
//
// ### QPS limit
//
// You can make up to 10 queries per second (QPS). If you exceed this limit, API calls are throttled, which may impact your business. We recommend that you plan your calls accordingly.
//
// @param request - DescribeFirewallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallTaskResponse
func (client *Client) DescribeFirewallTaskWithContext(ctx context.Context, request *DescribeFirewallTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeFirewallTaskResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic trend of a firewall.
//
// @param request - DescribeFirewallTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallTrafficTrendResponse
func (client *Client) DescribeFirewallTrafficTrendWithContext(ctx context.Context, request *DescribeFirewallTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeFirewallTrafficTrendResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFirewallTrafficTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFirewallTrafficTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the vSwitches that are created by Cloud Firewall.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for this operation is 10 per user. If you exceed the limit, API calls are throttled, which may affect your business. Therefore, call this operation at a reasonable rate.
//
// @param request - DescribeFirewallVSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallVSwitchResponse
func (client *Client) DescribeFirewallVSwitchWithContext(ctx context.Context, request *DescribeFirewallVSwitchRequest, runtime *dara.RuntimeOptions) (_result *DescribeFirewallVSwitchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the vSwitch resources for Cloud Firewall.
//
// @param request - DescribeFirewallVswitchResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallVswitchResourcesResponse
func (client *Client) DescribeFirewallVswitchResourcesWithContext(ctx context.Context, request *DescribeFirewallVswitchResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeFirewallVswitchResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeFirewallVswitchResources"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFirewallVswitchResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query IPS rules
//
// @param request - DescribeIPSRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIPSRulesResponse
func (client *Client) DescribeIPSRulesWithContext(ctx context.Context, request *DescribeIPSRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeIPSRulesResponse, _err error) {
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

	if !dara.IsNil(request.AttackApps) {
		query["AttackApps"] = request.AttackApps
	}

	if !dara.IsNil(request.AttackType) {
		query["AttackType"] = request.AttackType
	}

	if !dara.IsNil(request.Cve) {
		query["Cve"] = request.Cve
	}

	if !dara.IsNil(request.DefaultAction) {
		query["DefaultAction"] = request.DefaultAction
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryModify) {
		query["QueryModify"] = request.QueryModify
	}

	if !dara.IsNil(request.RuleAction) {
		query["RuleAction"] = request.RuleAction
	}

	if !dara.IsNil(request.RuleClass) {
		query["RuleClass"] = request.RuleClass
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleLevel) {
		query["RuleLevel"] = request.RuleLevel
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
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
		Action:      dara.String("DescribeIPSRules"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIPSRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about the member accounts of Cloud Firewall.
//
// Description:
//
// You can call this operation to query information about the member accounts of Cloud Firewall.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed the limit, API calls are throttled. This may affect your business. Call this operation at a reasonable rate.
//
// @param request - DescribeInstanceMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceMembersResponse
func (client *Client) DescribeInstanceMembersWithContext(ctx context.Context, request *DescribeInstanceMembersRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceMembersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the account in a resource directory for an instance.
//
// @param request - DescribeInstanceRdAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceRdAccountsResponse
func (client *Client) DescribeInstanceRdAccountsWithContext(ctx context.Context, request *DescribeInstanceRdAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceRdAccountsResponse, _err error) {
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

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceRdAccounts"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceRdAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceRiskLevelsResponse
func (client *Client) DescribeInstanceRiskLevelsWithContext(ctx context.Context, request *DescribeInstanceRiskLevelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceRiskLevelsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Internet security trends
//
// Description:
//
// ## QPS limit
//
// This API is limited to 10 requests per second per user. Exceeding this limit triggers throttling, which can disrupt your service. Plan your API calls accordingly.
//
// @param request - DescribeInternetDropTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetDropTrafficTrendResponse
func (client *Client) DescribeInternetDropTrafficTrendWithContext(ctx context.Context, request *DescribeInternetDropTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetDropTrafficTrendResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can obtain details about Internet access.
//
// @param request - DescribeInternetOpenDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenDetailResponse
func (client *Client) DescribeInternetOpenDetailWithContext(ctx context.Context, request *DescribeInternetOpenDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenDetailResponse, _err error) {
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

	if !dara.IsNil(request.ServiceNameFuzzy) {
		query["ServiceNameFuzzy"] = request.ServiceNameFuzzy
	}

	if !dara.IsNil(request.SortList) {
		query["SortList"] = request.SortList
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
		Action:      dara.String("DescribeInternetOpenDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetOpenDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an access control policy group in Cloud Firewall.
//
// @param request - DescribeInternetOpenIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenIpResponse
func (client *Client) DescribeInternetOpenIpWithContext(ctx context.Context, request *DescribeInternetOpenIpRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenIpResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ports that are open to the Internet.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeInternetOpenPortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenPortResponse
func (client *Client) DescribeInternetOpenPortWithContext(ctx context.Context, request *DescribeInternetOpenPortRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenPortResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries services exposed to the Internet.
//
// Description:
//
// ## QPS limits
//
// You can make up to 10 queries per second (QPS). If you exceed this limit, API calls are throttled. This may affect your business. We recommend that you make API calls at a reasonable rate.
//
// @param request - DescribeInternetOpenServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenServiceResponse
func (client *Client) DescribeInternetOpenServiceWithContext(ctx context.Context, request *DescribeInternetOpenServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves statistics about assets exposed to the Internet.
//
// @param request - DescribeInternetOpenStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetOpenStatisticResponse
func (client *Client) DescribeInternetOpenStatisticWithContext(ctx context.Context, request *DescribeInternetOpenStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetOpenStatisticResponse, _err error) {
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
		Action:      dara.String("DescribeInternetOpenStatistic"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetOpenStatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of Internet service names.
//
// @param request - DescribeInternetServiceNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetServiceNameListResponse
func (client *Client) DescribeInternetServiceNameListWithContext(ctx context.Context, request *DescribeInternetServiceNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetServiceNameListResponse, _err error) {
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
		Action:      dara.String("DescribeInternetServiceNameList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetServiceNameListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of Internet-facing SLB instances.
//
// @param request - DescribeInternetSlbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetSlbResponse
func (client *Client) DescribeInternetSlbWithContext(ctx context.Context, request *DescribeInternetSlbRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetSlbResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
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

	if !dara.IsNil(request.PublicIp) {
		query["PublicIp"] = request.PublicIp
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetSlb"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetSlbResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries the Top-N internet traffic over time.
//
// @param request - DescribeInternetTimeTopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetTimeTopResponse
func (client *Client) DescribeInternetTimeTopWithContext(ctx context.Context, request *DescribeInternetTimeTopRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetTimeTopResponse, _err error) {
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

	if !dara.IsNil(request.IPType) {
		query["IPType"] = request.IPType
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NatIP) {
		query["NatIP"] = request.NatIP
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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

	if !dara.IsNil(request.TrafficTime) {
		query["TrafficTime"] = request.TrafficTime
	}

	if !dara.IsNil(request.TrafficType) {
		query["TrafficType"] = request.TrafficType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetTimeTop"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetTimeTopResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top Internet traffic trends.
//
// @param request - DescribeInternetTrafficTopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetTrafficTopResponse
func (client *Client) DescribeInternetTrafficTopWithContext(ctx context.Context, request *DescribeInternetTrafficTopRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetTrafficTopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

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

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.RuleResult) {
		query["RuleResult"] = request.RuleResult
	}

	if !dara.IsNil(request.RuleSource) {
		query["RuleSource"] = request.RuleSource
	}

	if !dara.IsNil(request.ShowCountryName) {
		query["ShowCountryName"] = request.ShowCountryName
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetTrafficTop"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetTrafficTopResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Internet traffic trends.
//
// @param request - DescribeInternetTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetTrafficTrendResponse
func (client *Client) DescribeInternetTrafficTrendWithContext(ctx context.Context, request *DescribeInternetTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetTrafficTrendResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend of vulnerabilities on ECS instances.
//
// @param request - DescribeInvadeEcsTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvadeEcsTrendResponse
func (client *Client) DescribeInvadeEcsTrendWithContext(ctx context.Context, request *DescribeInvadeEcsTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvadeEcsTrendResponse, _err error) {
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
		Action:      dara.String("DescribeInvadeEcsTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvadeEcsTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a threat detection event.
//
// @param request - DescribeInvadeEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvadeEventDetailResponse
func (client *Client) DescribeInvadeEventDetailWithContext(ctx context.Context, request *DescribeInvadeEventDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvadeEventDetailResponse, _err error) {
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

	if !dara.IsNil(request.EventUuid) {
		query["EventUuid"] = request.EventUuid
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PublicIP) {
		query["PublicIP"] = request.PublicIP
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvadeEventDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInvadeEventDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Cloud Firewall threat detection events.
//
// @param request - DescribeInvadeEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvadeEventListResponse
func (client *Client) DescribeInvadeEventListWithContext(ctx context.Context, request *DescribeInvadeEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvadeEventListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of vulnerability names.
//
// @param request - DescribeInvadeEventNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvadeEventNameListResponse
func (client *Client) DescribeInvadeEventNameListWithContext(ctx context.Context, request *DescribeInvadeEventNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvadeEventNameListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics about intrusion events.
//
// @param request - DescribeInvadeEventStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvadeEventStatisticResponse
func (client *Client) DescribeInvadeEventStatisticWithContext(ctx context.Context, request *DescribeInvadeEventStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvadeEventStatisticResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of IPS Private IP Tracing associations.
//
// Description:
//
// This operation queries information about assets that are protected by Cloud Firewall. The results are paginated.
//
// ## Limits
//
// This operation is limited to 10 queries per second (QPS) per user. If you exceed the limit, API calls are throttled. This may affect your business. Call this operation at a reasonable rate.
//
// @param request - DescribeIpsPrivateAssocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpsPrivateAssocResponse
func (client *Client) DescribeIpsPrivateAssocWithContext(ctx context.Context, request *DescribeIpsPrivateAssocRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpsPrivateAssocResponse, _err error) {
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

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PublicIp) {
		query["PublicIp"] = request.PublicIp
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpsPrivateAssoc"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpsPrivateAssocResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about ISPs.
//
// @param request - DescribeIspInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspInfoResponse
func (client *Client) DescribeIspInfoWithContext(ctx context.Context, request *DescribeIspInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspInfoResponse, _err error) {
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
		Action:      dara.String("DescribeIspInfo"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about geographic locations.
//
// @param request - DescribeLocationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLocationInfoResponse
func (client *Client) DescribeLocationInfoWithContext(ctx context.Context, request *DescribeLocationInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeLocationInfoResponse, _err error) {
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
		Action:      dara.String("DescribeLocationInfo"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLocationInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes member information.
//
// @param request - DescribeMemberInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMemberInfoResponse
func (client *Client) DescribeMemberInfoWithContext(ctx context.Context, request *DescribeMemberInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeMemberInfoResponse, _err error) {
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
		Action:      dara.String("DescribeMemberInfo"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMemberInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the paging status of the NAT firewall.
//
// @param request - DescribeNatAclPageStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatAclPageStatusResponse
func (client *Client) DescribeNatAclPageStatusWithContext(ctx context.Context, request *DescribeNatAclPageStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatAclPageStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control policy groups for NAT firewalls.
//
// @param request - DescribeNatFirewallAclGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallAclGroupListResponse
func (client *Client) DescribeNatFirewallAclGroupListWithContext(ctx context.Context, request *DescribeNatFirewallAclGroupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallAclGroupListResponse, _err error) {
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
		Action:      dara.String("DescribeNatFirewallAclGroupList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallAclGroupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control policies for NAT firewalls.
//
// Description:
//
// This operation queries access control policies for NAT firewalls and returns the results in a paginated list.
//
// @param request - DescribeNatFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallControlPolicyResponse
func (client *Client) DescribeNatFirewallControlPolicyWithContext(ctx context.Context, request *DescribeNatFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Overview - NAT firewall blocking trends
//
// @param request - DescribeNatFirewallDropTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallDropTrafficTrendResponse
func (client *Client) DescribeNatFirewallDropTrafficTrendWithContext(ctx context.Context, request *DescribeNatFirewallDropTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallDropTrafficTrendResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries NAT firewall details.
//
// @param request - DescribeNatFirewallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallListResponse
func (client *Client) DescribeNatFirewallListWithContext(ctx context.Context, request *DescribeNatFirewallListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the priority range of an access control policy for a NAT firewall.
//
// Description:
//
// You can call this operation to query the priority range of an access control policy for outbound traffic on a NAT firewall.
//
// @param request - DescribeNatFirewallPolicyPriorUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallPolicyPriorUsedResponse
func (client *Client) DescribeNatFirewallPolicyPriorUsedWithContext(ctx context.Context, request *DescribeNatFirewallPolicyPriorUsedRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallPolicyPriorUsedResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the precheck details for a NAT firewall.
//
// @param request - DescribeNatFirewallPrecheckDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallPrecheckDetailResponse
func (client *Client) DescribeNatFirewallPrecheckDetailWithContext(ctx context.Context, request *DescribeNatFirewallPrecheckDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallPrecheckDetailResponse, _err error) {
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
		Action:      dara.String("DescribeNatFirewallPrecheckDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallPrecheckDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the quotas for a NAT firewall.
//
// @param request - DescribeNatFirewallQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallQuotaResponse
func (client *Client) DescribeNatFirewallQuotaWithContext(ctx context.Context, request *DescribeNatFirewallQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallQuotaResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top traffic data of a NAT firewall at a specific point in time.
//
// @param request - DescribeNatFirewallTimeTopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallTimeTopResponse
func (client *Client) DescribeNatFirewallTimeTopWithContext(ctx context.Context, request *DescribeNatFirewallTimeTopRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallTimeTopResponse, _err error) {
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
		Action:      dara.String("DescribeNatFirewallTimeTop"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatFirewallTimeTopResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Overview: NAT Traffic Trend
//
// @param request - DescribeNatFirewallTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatFirewallTrafficTrendResponse
func (client *Client) DescribeNatFirewallTrafficTrendWithContext(ctx context.Context, request *DescribeNatFirewallTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatFirewallTrafficTrendResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of network instances.
//
// @param request - DescribeNetworkInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkInstanceListResponse
func (client *Client) DescribeNetworkInstanceListWithContext(ctx context.Context, request *DescribeNetworkInstanceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkInstanceListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the relationships between network instances.
//
// @param request - DescribeNetworkInstanceRelationListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkInstanceRelationListResponse
func (client *Client) DescribeNetworkInstanceRelationListWithContext(ctx context.Context, request *DescribeNetworkInstanceRelationListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkInstanceRelationListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ratio of the top network traffic.
//
// @param request - DescribeNetworkTrafficTopRatioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkTrafficTopRatioResponse
func (client *Client) DescribeNetworkTrafficTopRatioWithContext(ctx context.Context, request *DescribeNetworkTrafficTopRatioRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkTrafficTopRatioResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves statistics about access sources for public IP addresses.
//
// @param request - DescribeOpenIpAccessSrcStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenIpAccessSrcStatResponse
func (client *Client) DescribeOpenIpAccessSrcStatWithContext(ctx context.Context, request *DescribeOpenIpAccessSrcStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeOpenIpAccessSrcStatResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of assets with outbound connections.
//
// @param request - DescribeOutgoingAssetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingAssetListResponse
func (client *Client) DescribeOutgoingAssetListWithContext(ctx context.Context, request *DescribeOutgoingAssetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingAssetListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries outbound destinations.
//
// @param request - DescribeOutgoingDestinationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDestinationResponse
func (client *Client) DescribeOutgoingDestinationWithContext(ctx context.Context, request *DescribeOutgoingDestinationRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDestinationResponse, _err error) {
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

	if !dara.IsNil(request.IsAITraffic) {
		query["IsAITraffic"] = request.IsAITraffic
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

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateIP) {
		query["PrivateIP"] = request.PrivateIP
	}

	if !dara.IsNil(request.PublicIP) {
		query["PublicIP"] = request.PublicIP
	}

	if !dara.IsNil(request.SecuritySuggest) {
		query["SecuritySuggest"] = request.SecuritySuggest
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
		Action:      dara.String("DescribeOutgoingDestination"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingDestinationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the categories of outbound connection destinations.
//
// @param request - DescribeOutgoingDestinationCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDestinationCategoryResponse
func (client *Client) DescribeOutgoingDestinationCategoryWithContext(ctx context.Context, request *DescribeOutgoingDestinationCategoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDestinationCategoryResponse, _err error) {
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

	if !dara.IsNil(request.TypeId) {
		query["TypeId"] = request.TypeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOutgoingDestinationCategory"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingDestinationCategoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Displays the destination IP of an active outbound connection.
//
// @param request - DescribeOutgoingDestinationIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDestinationIPResponse
func (client *Client) DescribeOutgoingDestinationIPWithContext(ctx context.Context, request *DescribeOutgoingDestinationIPRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an outbound destination IP address.
//
// @param request - DescribeOutgoingDestinationIPDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDestinationIPDetailResponse
func (client *Client) DescribeOutgoingDestinationIPDetailWithContext(ctx context.Context, request *DescribeOutgoingDestinationIPDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDestinationIPDetailResponse, _err error) {
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

	if !dara.IsNil(request.DstIP) {
		query["DstIP"] = request.DstIP
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
		Action:      dara.String("DescribeOutgoingDestinationIPDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingDestinationIPDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about outbound domain names.
//
// @param request - DescribeOutgoingDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDomainResponse
func (client *Client) DescribeOutgoingDomainWithContext(ctx context.Context, request *DescribeOutgoingDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDomainResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an outbound domain.
//
// @param request - DescribeOutgoingDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingDomainDetailResponse
func (client *Client) DescribeOutgoingDomainDetailWithContext(ctx context.Context, request *DescribeOutgoingDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingDomainDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of intrusion prevention threats.
//
// @param request - DescribeOutgoingRiskDomainAndIpCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingRiskDomainAndIpCountResponse
func (client *Client) DescribeOutgoingRiskDomainAndIpCountWithContext(ctx context.Context, request *DescribeOutgoingRiskDomainAndIpCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingRiskDomainAndIpCountResponse, _err error) {
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
		Action:      dara.String("DescribeOutgoingRiskDomainAndIpCount"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingRiskDomainAndIpCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend of outgoing connection threats.
//
// @param request - DescribeOutgoingRiskTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingRiskTrendResponse
func (client *Client) DescribeOutgoingRiskTrendWithContext(ctx context.Context, request *DescribeOutgoingRiskTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingRiskTrendResponse, _err error) {
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
		Action:      dara.String("DescribeOutgoingRiskTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOutgoingRiskTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves outbound connection statistics.
//
// @param request - DescribeOutgoingStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingStatisticResponse
func (client *Client) DescribeOutgoingStatisticWithContext(ctx context.Context, request *DescribeOutgoingStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingStatisticResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries outbound connection tags.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 per user. If you exceed the limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeOutgoingTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOutgoingTagResponse
func (client *Client) DescribeOutgoingTagWithContext(ctx context.Context, request *DescribeOutgoingTagRequest, runtime *dara.RuntimeOptions) (_result *DescribeOutgoingTagResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribePageDocuments is deprecated
//
// Summary:
//
// Queries the FAQ for a page.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribePageDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePageDocumentsResponse
func (client *Client) DescribePageDocumentsWithContext(ctx context.Context, request *DescribePageDocumentsRequest, runtime *dara.RuntimeOptions) (_result *DescribePageDocumentsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of strict mode for access control policies.
//
// Description:
//
// You can call this operation to query the status of strict mode for access control policies.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) for each user. API calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - DescribePolicyAdvancedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyAdvancedConfigResponse
func (client *Client) DescribePolicyAdvancedConfigWithContext(ctx context.Context, request *DescribePolicyAdvancedConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyAdvancedConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the effective priority range of access control policies.
//
// Description:
//
// This operation queries the effective priority range of access control policies for inbound and outbound traffic.
//
// ## QPS limit
//
// The QPS limit for this operation is 10 requests per second per user. Calls that exceed this limit are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - DescribePolicyPriorUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyPriorUsedResponse
func (client *Client) DescribePolicyPriorUsedWithContext(ctx context.Context, request *DescribePolicyPriorUsedRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyPriorUsedResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries pay-as-you-go 2.0 bills.
//
// Description:
//
// For pay-as-you-go users, the bill details are accurate to the specific resource instance level. For subscription users, only overall queries are supported.
//
// @param request - DescribePostpayBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayBillResponse
func (client *Client) DescribePostpayBillWithContext(ctx context.Context, request *DescribePostpayBillRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayBillResponse, _err error) {
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

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePostpayBill"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePostpayBillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of pay-as-you-go protection.
//
// @param request - DescribePostpayEnabledProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayEnabledProtectionResponse
func (client *Client) DescribePostpayEnabledProtectionWithContext(ctx context.Context, request *DescribePostpayEnabledProtectionRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayEnabledProtectionResponse, _err error) {
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
		Action:      dara.String("DescribePostpayEnabledProtection"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePostpayEnabledProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of pay-as-you-go billing traffic.
//
// Description:
//
// For pay-as-you-go users, the details are accurate to the specific resource instance level. For subscription users, only overall queries are supported.
//
// @param request - DescribePostpayTrafficDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayTrafficDetailResponse
func (client *Client) DescribePostpayTrafficDetailWithContext(ctx context.Context, request *DescribePostpayTrafficDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayTrafficDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total pay-as-you-go traffic for all border firewalls.
//
// Description:
//
// The statistics are for the current Cloud Firewall instance and include all data from the date of purchase.
//
// @param request - DescribePostpayTrafficTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayTrafficTotalResponse
func (client *Client) DescribePostpayTrafficTotalWithContext(ctx context.Context, request *DescribePostpayTrafficTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayTrafficTotalResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Internet Border firewall for a pay-as-you-go instance.
//
// @param request - DescribePostpayUserInternetStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayUserInternetStatusResponse
func (client *Client) DescribePostpayUserInternetStatusWithContext(ctx context.Context, request *DescribePostpayUserInternetStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayUserInternetStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the NAT border firewall status for a pay-as-you-go Cloud Firewall.
//
// @param request - DescribePostpayUserNatStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayUserNatStatusResponse
func (client *Client) DescribePostpayUserNatStatusWithContext(ctx context.Context, request *DescribePostpayUserNatStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayUserNatStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the virtual private cloud (VPC) firewall switch module for a pay-as-you-go user.
//
// @param request - DescribePostpayUserVpcStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostpayUserVpcStatusResponse
func (client *Client) DescribePostpayUserVpcStatusWithContext(ctx context.Context, request *DescribePostpayUserVpcStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePostpayUserVpcStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes one or more prefix lists.
//
// @param request - DescribePrefixListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrefixListsResponse
func (client *Client) DescribePrefixListsWithContext(ctx context.Context, request *DescribePrefixListsRequest, runtime *dara.RuntimeOptions) (_result *DescribePrefixListsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billing overview of a subscription 2.0 instance.
//
// Description:
//
// The statistics cover the current Cloud Firewall instance of the user, including all data since the purchase date.
//
// @param request - DescribePrepayBillTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrepayBillTotalResponse
func (client *Client) DescribePrepayBillTotalWithContext(ctx context.Context, request *DescribePrepayBillTotalRequest, runtime *dara.RuntimeOptions) (_result *DescribePrepayBillTotalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillType) {
		query["BillType"] = request.BillType
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrepayBillTotal"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrepayBillTotalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of private DNS domain names.
//
// Description:
//
// Queries the list of domain names that require private DNS endpoints for domain name resolution.
//
// @param request - DescribePrivateDnsDomainNameListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrivateDnsDomainNameListResponse
func (client *Client) DescribePrivateDnsDomainNameListWithContext(ctx context.Context, request *DescribePrivateDnsDomainNameListRequest, runtime *dara.RuntimeOptions) (_result *DescribePrivateDnsDomainNameListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a private DNS endpoint.
//
// Description:
//
// This operation queries the details of a private DNS endpoint.
//
// @param request - DescribePrivateDnsEndpointDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrivateDnsEndpointDetailResponse
func (client *Client) DescribePrivateDnsEndpointDetailWithContext(ctx context.Context, request *DescribePrivateDnsEndpointDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribePrivateDnsEndpointDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of private DNS endpoints.
//
// @param request - DescribePrivateDnsEndpointListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrivateDnsEndpointListResponse
func (client *Client) DescribePrivateDnsEndpointListWithContext(ctx context.Context, request *DescribePrivateDnsEndpointListRequest, runtime *dara.RuntimeOptions) (_result *DescribePrivateDnsEndpointListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns statistics about private DNS.
//
// @param request - DescribePrivateDnsStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrivateDnsStatisticsResponse
func (client *Client) DescribePrivateDnsStatisticsWithContext(ctx context.Context, request *DescribePrivateDnsStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribePrivateDnsStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNameCreatedEndTime) {
		query["DomainNameCreatedEndTime"] = request.DomainNameCreatedEndTime
	}

	if !dara.IsNil(request.DomainNameCreatedStartTime) {
		query["DomainNameCreatedStartTime"] = request.DomainNameCreatedStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrivateDnsStatistics"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrivateDnsStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves region information.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeRegionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionInfoResponse
func (client *Client) DescribeRegionInfoWithContext(ctx context.Context, request *DescribeRegionInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic redirection settings by region and asset type.
//
// Description:
//
// This operation is used to retrieve DNS resolution results for a domain name. Currently, only resolution results from Alibaba Cloud DNS are supported. The domain name that you want to query must use Alibaba Cloud DNS. Otherwise, the resolution results cannot be retrieved.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - DescribeRegionResourceTypeAutoEnableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionResourceTypeAutoEnableResponse
func (client *Client) DescribeRegionResourceTypeAutoEnableWithContext(ctx context.Context, request *DescribeRegionResourceTypeAutoEnableRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionResourceTypeAutoEnableResponse, _err error) {
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
		Action:      dara.String("DescribeRegionResourceTypeAutoEnable"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionResourceTypeAutoEnableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default traffic redirection for an asset type.
//
// Description:
//
// This operation is used to retrieve DNS resolution results for a domain name. Currently, only resolution results from Alibaba Cloud DNS are supported. The domain name that you want to query must use Alibaba Cloud DNS. Otherwise, the resolution results cannot be retrieved.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeResourceTypeAutoEnableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceTypeAutoEnableResponse
func (client *Client) DescribeResourceTypeAutoEnableWithContext(ctx context.Context, request *DescribeResourceTypeAutoEnableRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceTypeAutoEnableResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of intrusion prevention events.
//
// Description:
//
// You can use this operation to query and download the details of intrusion prevention events. We recommend querying 5 to 10 entries at a time. To prevent query timeouts, set the NoLocation parameter to true if you do not need IP geolocation information.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If you exceed the limit, your API calls are throttled. This may affect your business. Make calls to this operation at a reasonable rate.
//
// @param request - DescribeRiskEventGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventGroupResponse
func (client *Client) DescribeRiskEventGroupWithContext(ctx context.Context, request *DescribeRiskEventGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventPayloadResponse
func (client *Client) DescribeRiskEventPayloadWithContext(ctx context.Context, request *DescribeRiskEventPayloadRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventPayloadResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides statistics for intrusion prevention events.
//
// @param request - DescribeRiskEventStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventStatisticResponse
func (client *Client) DescribeRiskEventStatisticWithContext(ctx context.Context, request *DescribeRiskEventStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventStatisticResponse, _err error) {
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
		Action:      dara.String("DescribeRiskEventStatistic"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskEventStatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the ranking of applications that are targeted by intrusion prevention attacks.
//
// @param request - DescribeRiskEventTopAttackAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventTopAttackAppResponse
func (client *Client) DescribeRiskEventTopAttackAppWithContext(ctx context.Context, request *DescribeRiskEventTopAttackAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventTopAttackAppResponse, _err error) {
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
		Action:      dara.String("DescribeRiskEventTopAttackApp"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskEventTopAttackAppResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top assets targeted by attacks.
//
// Description:
//
// ## QPS limits
//
// You can make up to 10 queries per second (QPS) to this API. If you exceed this limit, your API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeRiskEventTopAttackAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventTopAttackAssetResponse
func (client *Client) DescribeRiskEventTopAttackAssetWithContext(ctx context.Context, request *DescribeRiskEventTopAttackAssetRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventTopAttackAssetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a ranked list of attack types from intrusion prevention events.
//
// @param request - DescribeRiskEventTopAttackTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskEventTopAttackTypeResponse
func (client *Client) DescribeRiskEventTopAttackTypeWithContext(ctx context.Context, request *DescribeRiskEventTopAttackTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskEventTopAttackTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuyVersion) {
		query["BuyVersion"] = request.BuyVersion
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeRiskEventTopAttackType"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskEventTopAttackTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeRiskSecurityGroupDetail is deprecated
//
// Summary:
//
// Retrieves the details of a risk security group.
//
// @param request - DescribeRiskSecurityGroupDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskSecurityGroupDetailResponse
func (client *Client) DescribeRiskSecurityGroupDetailWithContext(ctx context.Context, request *DescribeRiskSecurityGroupDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskSecurityGroupDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data leak event.
//
// @param request - DescribeSdlEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSdlEventDetailResponse
func (client *Client) DescribeSdlEventDetailWithContext(ctx context.Context, request *DescribeSdlEventDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeSdlEventDetailResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SrcIp) {
		query["SrcIp"] = request.SrcIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSdlEventDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSdlEventDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query data breach events.
//
// @param request - DescribeSdlEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSdlEventListResponse
func (client *Client) DescribeSdlEventListWithContext(ctx context.Context, request *DescribeSdlEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSdlEventListResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.OnlyAiEvt) {
		query["OnlyAiEvt"] = request.OnlyAiEvt
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SensitiveLevel) {
		query["SensitiveLevel"] = request.SensitiveLevel
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SrcIp) {
		query["SrcIp"] = request.SrcIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSdlEventList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSdlEventListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of sensitive data involved in data leaks.
//
// @param request - DescribeSdlEventSdListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSdlEventSdListResponse
func (client *Client) DescribeSdlEventSdListWithContext(ctx context.Context, request *DescribeSdlEventSdListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSdlEventSdListResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SrcIp) {
		query["SrcIp"] = request.SrcIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSdlEventSdList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSdlEventSdListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics about data leaks.
//
// @param request - DescribeSdlEventStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSdlEventStatisticResponse
func (client *Client) DescribeSdlEventStatisticWithContext(ctx context.Context, request *DescribeSdlEventStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSdlEventStatisticResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSdlEventStatistic"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSdlEventStatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the packet payload of a sensitive data leak event.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeSdlLastPayloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSdlLastPayloadResponse
func (client *Client) DescribeSdlLastPayloadWithContext(ctx context.Context, request *DescribeSdlLastPayloadRequest, runtime *dara.RuntimeOptions) (_result *DescribeSdlLastPayloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DstIp) {
		query["DstIp"] = request.DstIp
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SensitiveCategory) {
		query["SensitiveCategory"] = request.SensitiveCategory
	}

	if !dara.IsNil(request.SrcIp) {
		query["SrcIp"] = request.SrcIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSdlLastPayload"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSdlLastPayloadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of sensitive data.
//
// @param request - DescribeSdlStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSdlStatisticResponse
func (client *Client) DescribeSdlStatisticWithContext(ctx context.Context, request *DescribeSdlStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSdlStatisticResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSdlStatistic"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSdlStatisticResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of the safe mode.
//
// Description:
//
// You can use this operation to query the safe mode of Cloud Firewall.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) for each user. If you exceed this limit, your API calls are throttled. Throttling can affect your business operations. We recommend that you plan your API calls accordingly.
//
// @param request - DescribeSecurityModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityModeResponse
func (client *Client) DescribeSecurityModeWithContext(ctx context.Context, request *DescribeSecurityModeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityModeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeSecurityProxy is deprecated, please use Cloudfw::2017-12-07::DescribeNatFirewallList instead.
//
// Summary:
//
// Retrieves NAT firewall information.
//
// @param request - DescribeSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityProxyResponse
func (client *Client) DescribeSecurityProxyWithContext(ctx context.Context, request *DescribeSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityProxyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes NAT firewall resources.
//
// @param request - DescribeSecurityProxyResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityProxyResourcesResponse
func (client *Client) DescribeSecurityProxyResourcesWithContext(ctx context.Context, request *DescribeSecurityProxyResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityProxyResourcesResponse, _err error) {
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

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityProxyResources"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityProxyResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the sensitive data detection switch.
//
// @param request - DescribeSensitiveSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveSwitchResponse
func (client *Client) DescribeSensitiveSwitchWithContext(ctx context.Context, request *DescribeSensitiveSwitchRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveSwitchResponse, _err error) {
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

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentCategory) {
		query["ParentCategory"] = request.ParentCategory
	}

	if !dara.IsNil(request.SensitiveCategory) {
		query["SensitiveCategory"] = request.SensitiveCategory
	}

	if !dara.IsNil(request.SensitiveLevel) {
		query["SensitiveLevel"] = request.SensitiveLevel
	}

	if !dara.IsNil(request.SwitchStatus) {
		query["SwitchStatus"] = request.SwitchStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSensitiveSwitch"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization information of a service-linked role (SLR) for a user.
//
// @param request - DescribeSlrGrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlrGrantResponse
func (client *Client) DescribeSlrGrantWithContext(ctx context.Context, request *DescribeSlrGrantRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlrGrantResponse, _err error) {
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
		Action:      dara.String("DescribeSlrGrant"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlrGrantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the enabled status of Log Service (SLS).
//
// @param request - DescribeSlsAnalyzeOpenStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsAnalyzeOpenStatusResponse
func (client *Client) DescribeSlsAnalyzeOpenStatusWithContext(ctx context.Context, request *DescribeSlsAnalyzeOpenStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsAnalyzeOpenStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can obtain an ACL backup for a VPC firewall for a transit router.
//
// @param tmpReq - DescribeTrFirewallPolicyBackUpAssociationListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallPolicyBackUpAssociationListResponse
func (client *Client) DescribeTrFirewallPolicyBackUpAssociationListWithContext(ctx context.Context, tmpReq *DescribeTrFirewallPolicyBackUpAssociationListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallPolicyBackUpAssociationListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of routing policies for a VPC firewall for a transit router.
//
// @param request - DescribeTrFirewallV2RoutePolicyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallV2RoutePolicyListResponse
func (client *Client) DescribeTrFirewallV2RoutePolicyListWithContext(ctx context.Context, request *DescribeTrFirewallV2RoutePolicyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallV2RoutePolicyListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a VPC firewall for a transit router.
//
// @param request - DescribeTrFirewallsV2DetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallsV2DetailResponse
func (client *Client) DescribeTrFirewallsV2DetailWithContext(ctx context.Context, request *DescribeTrFirewallsV2DetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallsV2DetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of VPC firewalls for a transit router.
//
// @param request - DescribeTrFirewallsV2ListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallsV2ListResponse
func (client *Client) DescribeTrFirewallsV2ListWithContext(ctx context.Context, request *DescribeTrFirewallsV2ListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallsV2ListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route tables for a VPC firewall for a transit router.
//
// @param request - DescribeTrFirewallsV2RouteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrFirewallsV2RouteListResponse
func (client *Client) DescribeTrFirewallsV2RouteListWithContext(ctx context.Context, request *DescribeTrFirewallsV2RouteListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrFirewallsV2RouteListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries log traffic information.
//
// @param request - DescribeTrafficLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrafficLogResponse
func (client *Client) DescribeTrafficLogWithContext(ctx context.Context, request *DescribeTrafficLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrafficLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclPreRuleId) {
		query["AclPreRuleId"] = request.AclPreRuleId
	}

	if !dara.IsNil(request.AclPreState) {
		query["AclPreState"] = request.AclPreState
	}

	if !dara.IsNil(request.AppDpiState) {
		query["AppDpiState"] = request.AppDpiState
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AssetRegion) {
		query["AssetRegion"] = request.AssetRegion
	}

	if !dara.IsNil(request.AttackType) {
		query["AttackType"] = request.AttackType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainUrl) {
		query["DomainUrl"] = request.DomainUrl
	}

	if !dara.IsNil(request.DstIP) {
		query["DstIP"] = request.DstIP
	}

	if !dara.IsNil(request.DstPort) {
		query["DstPort"] = request.DstPort
	}

	if !dara.IsNil(request.DstVpcId) {
		query["DstVpcId"] = request.DstVpcId
	}

	if !dara.IsNil(request.DstVpcRegionNo) {
		query["DstVpcRegionNo"] = request.DstVpcRegionNo
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.FlowType) {
		query["FlowType"] = request.FlowType
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
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

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.NatFirewallId) {
		query["NatFirewallId"] = request.NatFirewallId
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleResult) {
		query["RuleResult"] = request.RuleResult
	}

	if !dara.IsNil(request.RuleSource) {
		query["RuleSource"] = request.RuleSource
	}

	if !dara.IsNil(request.RuleSourceFinal) {
		query["RuleSourceFinal"] = request.RuleSourceFinal
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

	if !dara.IsNil(request.SrcPort) {
		query["SrcPort"] = request.SrcPort
	}

	if !dara.IsNil(request.SrcPrivateIP) {
		query["SrcPrivateIP"] = request.SrcPrivateIP
	}

	if !dara.IsNil(request.SrcVpcId) {
		query["SrcVpcId"] = request.SrcVpcId
	}

	if !dara.IsNil(request.SrcVpcRegionNo) {
		query["SrcVpcRegionNo"] = request.SrcVpcRegionNo
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TlsScopeId) {
		query["TlsScopeId"] = request.TlsScopeId
	}

	if !dara.IsNil(request.VpcFirewallId) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !dara.IsNil(request.VulLevel) {
		query["VulLevel"] = request.VulLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrafficLog"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrafficLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Transit Router resources.
//
// @param request - DescribeTransitRouterResourcesListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransitRouterResourcesListResponse
func (client *Client) DescribeTransitRouterResourcesListWithContext(ctx context.Context, request *DescribeTransitRouterResourcesListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransitRouterResourcesListResponse, _err error) {
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

	if !dara.IsNil(request.FirewallId) {
		query["FirewallId"] = request.FirewallId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
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
		Action:      dara.String("DescribeTransitRouterResourcesList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTransitRouterResourcesListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend of unprotected ports.
//
// @param request - DescribeUnprotectedPortTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUnprotectedPortTrendResponse
func (client *Client) DescribeUnprotectedPortTrendWithContext(ctx context.Context, request *DescribeUnprotectedPortTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeUnprotectedPortTrendResponse, _err error) {
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
		Action:      dara.String("DescribeUnprotectedPortTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUnprotectedPortTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend of unprotected vulnerabilities.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, API calls are throttled, which may impact your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeUnprotectedVulnTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUnprotectedVulnTrendResponse
func (client *Client) DescribeUnprotectedVulnTrendWithContext(ctx context.Context, request *DescribeUnprotectedVulnTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeUnprotectedVulnTrendResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the user\\"s alert configuration.
//
// Description:
//
// ## QPS limit
//
// The QPS limit for this interface is 10 calls per second per user. Exceeding this limit throttles API calls and may affect your service. Plan your calls accordingly.
//
// @param request - DescribeUserAlarmConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAlarmConfigResponse
func (client *Client) DescribeUserAlarmConfigWithContext(ctx context.Context, request *DescribeUserAlarmConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAlarmConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeUserAssetIPTrafficInfo is deprecated
//
// Summary:
//
// Queries the traffic information for a specified asset.
//
// @param request - DescribeUserAssetIPTrafficInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAssetIPTrafficInfoResponse
func (client *Client) DescribeUserAssetIPTrafficInfoWithContext(ctx context.Context, request *DescribeUserAssetIPTrafficInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAssetIPTrafficInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves version information for a user.
//
// Description:
//
// This operation queries information about your Cloud Firewall instance.
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable frequency.
//
// @param request - DescribeUserBuyVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserBuyVersionResponse
func (client *Client) DescribeUserBuyVersionWithContext(ctx context.Context, request *DescribeUserBuyVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserBuyVersionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the intrusion prevention system (IPS) whitelist for the Internet Border.
//
// Description:
//
// ## QPS limits
//
// The queries per second (QPS) limit for this API is 10 calls per second for each user. If you exceed this limit, API calls are throttled, which can impact your business. We recommend that you call this API at a reasonable rate.
//
// @param request - DescribeUserIPSWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserIPSWhitelistResponse
func (client *Client) DescribeUserIPSWhitelistWithContext(ctx context.Context, request *DescribeUserIPSWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserIPSWhitelistResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IPS configuration list of a virtual private cloud (VPC) firewall.
//
// @param request - DescribeVfwIPSConfigListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVfwIPSConfigListResponse
func (client *Client) DescribeVfwIPSConfigListWithContext(ctx context.Context, request *DescribeVfwIPSConfigListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVfwIPSConfigListResponse, _err error) {
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
		Action:      dara.String("DescribeVfwIPSConfigList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVfwIPSConfigListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access details of a VPC firewall.
//
// Description:
//
// ## QPS limits
//
// The queries per second (QPS) limit for this API is 10 for each user. If you exceed this limit, API calls are throttled, which can affect your business. We recommend that you call this API at a reasonable rate.
//
// @param request - DescribeVpcFirewallAccessDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallAccessDetailResponse
func (client *Client) DescribeVpcFirewallAccessDetailWithContext(ctx context.Context, request *DescribeVpcFirewallAccessDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallAccessDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all access control policy groups for VPC firewalls.
//
// Description:
//
// This operation queries information about all access control policy groups for VPC firewalls.
//
// ## QPS limit
//
// The QPS limit for this operation is 10 requests per second per user. API calls that exceed this limit are throttled, potentially affecting your business. Plan your calls accordingly.
//
// @param request - DescribeVpcFirewallAclGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallAclGroupListResponse
func (client *Client) DescribeVpcFirewallAclGroupListWithContext(ctx context.Context, request *DescribeVpcFirewallAclGroupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallAclGroupListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the asset list of a VPC firewall.
//
// @param request - DescribeVpcFirewallAssetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallAssetListResponse
func (client *Client) DescribeVpcFirewallAssetListWithContext(ctx context.Context, request *DescribeVpcFirewallAssetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallAssetListResponse, _err error) {
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

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.EcsInstanceName) {
		query["EcsInstanceName"] = request.EcsInstanceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IPProtocol) {
		query["IPProtocol"] = request.IPProtocol
	}

	if !dara.IsNil(request.IsAITraffic) {
		query["IsAITraffic"] = request.IsAITraffic
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
		Action:      dara.String("DescribeVpcFirewallAssetList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallAssetListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions where the VPC firewall is enabled for asset protection.
//
// @param request - DescribeVpcFirewallAssetRegionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallAssetRegionListResponse
func (client *Client) DescribeVpcFirewallAssetRegionListWithContext(ctx context.Context, request *DescribeVpcFirewallAssetRegionListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallAssetRegionListResponse, _err error) {
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
		Action:      dara.String("DescribeVpcFirewallAssetRegionList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallAssetRegionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a VPC firewall that protects traffic between a network instance in a Cloud Enterprise Network (CEN) and a specified VPC.
//
// Description:
//
// You can call this operation to query the details of a VPC firewall. The VPC firewall protects traffic between a specified VPC and a network instance in a Cloud Enterprise Network (CEN). The network instance can be a VPC, a Virtual Border Router (VBR), or a Cloud Connect Network (CCN) instance.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 10 for each user. If you exceed the limit, your API calls are throttled. This may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeVpcFirewallCenDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallCenDetailResponse
func (client *Client) DescribeVpcFirewallCenDetailWithContext(ctx context.Context, request *DescribeVpcFirewallCenDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallCenDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of VPC firewalls that protect traffic between a specified VPC and network instances in a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// This operation queries the details of a VPC firewall. The firewall protects traffic between a specified VPC and a network instance that is attached to a Cloud Enterprise Network (CEN) instance. The network instance can be a VPC, a Virtual Border Router (VBR), or a Cloud Connect Network (CCN) instance.
//
// ## Limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. This may affect your business. We recommend that you plan your calls accordingly.
//
// @param request - DescribeVpcFirewallCenListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallCenListResponse
func (client *Client) DescribeVpcFirewallCenListWithContext(ctx context.Context, request *DescribeVpcFirewallCenListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallCenListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Cloud Enterprise Network (CEN) instances for a VPC.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If you exceed this limit, API calls are throttled, which can affect your business. Plan your calls accordingly.
//
// @param request - DescribeVpcFirewallCenSummaryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallCenSummaryListResponse
func (client *Client) DescribeVpcFirewallCenSummaryListWithContext(ctx context.Context, request *DescribeVpcFirewallCenSummaryListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallCenSummaryListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all access control policies for a specific VPC boundary firewall.
//
// Description:
//
// This operation queries the access control policies for a VPC firewall. A VPC firewall uses different access control policies to protect traffic between two VPCs that are connected via Cloud Enterprise Network (CEN) or Express Connect.
//
// ## QPS limit
//
// The QPS limit for this operation is 10 requests per second per account. If you exceed this limit, your API calls are throttled.
//
// @param request - DescribeVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallControlPolicyResponse
func (client *Client) DescribeVpcFirewallControlPolicyWithContext(ctx context.Context, request *DescribeVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the intrusion prevention configuration of a specified VPC firewall.
//
// Description:
//
// This operation queries the intrusion prevention configuration of a specified VPC firewall. Before you call this operation, you must create a VPC firewall instance.
//
// ## QPS limit
//
// This API operation has a limit of 10 queries per second (QPS) per user. If you exceed this limit, your calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeVpcFirewallDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallDefaultIPSConfigResponse
func (client *Client) DescribeVpcFirewallDefaultIPSConfigWithContext(ctx context.Context, request *DescribeVpcFirewallDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallDefaultIPSConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a VPC firewall that protects traffic between two VPCs connected by an Express Connect circuit.
//
// Description:
//
// This operation queries the details of a VPC firewall. The VPC firewall protects traffic between two VPCs that are connected by an Express Connect circuit. Before you call this operation, you must create a VPC firewall by calling the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 10 calls per second for each user. If you exceed this limit, your API calls are throttled. This can affect your business. Plan your calls accordingly.
//
// @param request - DescribeVpcFirewallDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallDetailResponse
func (client *Client) DescribeVpcFirewallDetailWithContext(ctx context.Context, request *DescribeVpcFirewallDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of domain names accessed through a VPC firewall.
//
// Description:
//
// ###
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed this limit, API calls are throttled. Throttling can affect your business. Call this operation at a reasonable rate.
//
// @param request - DescribeVpcFirewallDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallDomainListResponse
func (client *Client) DescribeVpcFirewallDomainListWithContext(ctx context.Context, request *DescribeVpcFirewallDomainListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallDomainListResponse, _err error) {
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

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallDomainList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallDomainListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries connections to a specified domain name through VPC Firewall.
//
// @param request - DescribeVpcFirewallDomainRelationListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallDomainRelationListResponse
func (client *Client) DescribeVpcFirewallDomainRelationListWithContext(ctx context.Context, request *DescribeVpcFirewallDomainRelationListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallDomainRelationListResponse, _err error) {
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

	if !dara.IsNil(request.DomainList) {
		query["DomainList"] = request.DomainList
	}

	if !dara.IsNil(request.DstIP) {
		query["DstIP"] = request.DstIP
	}

	if !dara.IsNil(request.DstVpcId) {
		query["DstVpcId"] = request.DstVpcId
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

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallDomainRelationList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallDomainRelationListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the blocked traffic trend for the VPC firewall.
//
// @param request - DescribeVpcFirewallDropTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallDropTrafficTrendResponse
func (client *Client) DescribeVpcFirewallDropTrafficTrendWithContext(ctx context.Context, request *DescribeVpcFirewallDropTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallDropTrafficTrendResponse, _err error) {
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
		Action:      dara.String("DescribeVpcFirewallDropTrafficTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallDropTrafficTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the IPS whitelist of a VPC firewall.
//
// @param request - DescribeVpcFirewallIPSWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallIPSWhitelistResponse
func (client *Client) DescribeVpcFirewallIPSWhitelistWithContext(ctx context.Context, request *DescribeVpcFirewallIPSWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallIPSWhitelistResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a VPC firewall that protects traffic between two VPCs connected by an Express Connect circuit.
//
// Description:
//
// This operation queries a paginated list of VPC firewalls. These firewalls protect traffic between two VPCs that are connected using Express Connect.
//
// ### QPS limit
//
// Each Alibaba Cloud account can send up to 10 queries per second (QPS). If this limit is exceeded, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeVpcFirewallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallListResponse
func (client *Client) DescribeVpcFirewallListWithContext(ctx context.Context, request *DescribeVpcFirewallListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of vSwitches for a VPC firewall created in manual mode.
//
// @param request - DescribeVpcFirewallManualVSwitchListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallManualVSwitchListResponse
func (client *Client) DescribeVpcFirewallManualVSwitchListWithContext(ctx context.Context, request *DescribeVpcFirewallManualVSwitchListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallManualVSwitchListResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeVpcFirewallManualVSwitchList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallManualVSwitchListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the effective priority range for access control policies in a specified VPC firewall policy group.
//
// Description:
//
// This operation queries the effective priority range for access control policies in a specified VPC firewall policy group.
//
// ## Limits
//
// The queries per second (QPS) limit for this operation is 10 for each user. If you exceed the limit, API calls are throttled. This may impact your business. Call this operation an appropriate number of times to prevent interruptions.
//
// @param request - DescribeVpcFirewallPolicyPriorUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallPolicyPriorUsedResponse
func (client *Client) DescribeVpcFirewallPolicyPriorUsedWithContext(ctx context.Context, request *DescribeVpcFirewallPolicyPriorUsedRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallPolicyPriorUsedResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a VPC firewall precheck.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled. This may impact your business. Plan your calls accordingly.
//
// @param request - DescribeVpcFirewallPrecheckDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallPrecheckDetailResponse
func (client *Client) DescribeVpcFirewallPrecheckDetailWithContext(ctx context.Context, request *DescribeVpcFirewallPrecheckDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallPrecheckDetailResponse, _err error) {
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
		Action:      dara.String("DescribeVpcFirewallPrecheckDetail"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallPrecheckDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a summary of VPC firewalls.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for this API operation is 10 for each user. If you exceed this limit, API calls are throttled. This can affect your business. Plan your API calls accordingly.
//
// @param request - DescribeVpcFirewallSummaryInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallSummaryInfoResponse
func (client *Client) DescribeVpcFirewallSummaryInfoWithContext(ctx context.Context, request *DescribeVpcFirewallSummaryInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallSummaryInfoResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of assets that access domain names through a VPC firewall.
//
// @param request - DescribeVpcFirewallTrafficAssetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallTrafficAssetListResponse
func (client *Client) DescribeVpcFirewallTrafficAssetListWithContext(ctx context.Context, request *DescribeVpcFirewallTrafficAssetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallTrafficAssetListResponse, _err error) {
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

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
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
		Action:      dara.String("DescribeVpcFirewallTrafficAssetList"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallTrafficAssetListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the traffic trend of a virtual private cloud (VPC) firewall.
//
// @param request - DescribeVpcFirewallTrafficTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallTrafficTrendResponse
func (client *Client) DescribeVpcFirewallTrafficTrendWithContext(ctx context.Context, request *DescribeVpcFirewallTrafficTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallTrafficTrendResponse, _err error) {
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

	if !dara.IsNil(request.PeerVpcId) {
		query["PeerVpcId"] = request.PeerVpcId
	}

	if !dara.IsNil(request.PrivateIP) {
		query["PrivateIP"] = request.PrivateIP
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
		Action:      dara.String("DescribeVpcFirewallTrafficTrend"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallTrafficTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the available zones for a VPC firewall.
//
// @param request - DescribeVpcFirewallZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcFirewallZoneResponse
func (client *Client) DescribeVpcFirewallZoneWithContext(ctx context.Context, request *DescribeVpcFirewallZoneRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcFirewallZoneResponse, _err error) {
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

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcFirewallZone"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcFirewallZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Virtual Private Clouds (VPCs).
//
// @param request - DescribeVpcListLiteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcListLiteResponse
func (client *Client) DescribeVpcListLiteWithContext(ctx context.Context, request *DescribeVpcListLiteRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcListLiteResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones that are available for VPCs.
//
// @param request - DescribeVpcZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcZoneResponse
func (client *Client) DescribeVpcZoneWithContext(ctx context.Context, request *DescribeVpcZoneRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcZoneResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of vulnerabilities that Cloud Firewall can protect against.
//
// @param request - DescribeVulnerabilityProtectedListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVulnerabilityProtectedListResponse
func (client *Client) DescribeVulnerabilityProtectedListWithContext(ctx context.Context, request *DescribeVulnerabilityProtectedListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVulnerabilityProtectedListResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables sensitive data discovery for a protected asset.
//
// @param request - DisableSdlProtectedAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSdlProtectedAssetResponse
func (client *Client) DisableSdlProtectedAssetWithContext(ctx context.Context, request *DisableSdlProtectedAssetRequest, runtime *dara.RuntimeOptions) (_result *DisableSdlProtectedAssetResponse, _err error) {
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
		Action:      dara.String("DisableSdlProtectedAsset"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSdlProtectedAssetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables data breach protection for assets.
//
// @param request - EnableSdlProtectedAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSdlProtectedAssetResponse
func (client *Client) EnableSdlProtectedAssetWithContext(ctx context.Context, request *EnableSdlProtectedAssetRequest, runtime *dara.RuntimeOptions) (_result *EnableSdlProtectedAssetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the download path for the certificate of a Transport Layer Security (TLS) inspection policy.
//
// Description:
//
// This operation returns a temporary download link for the Certificate Authority (CA) certificate. The link is valid for one minute. After the link expires, call this operation again to obtain a new download link.
//
// @param request - GetTlsInspectCertificateDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTlsInspectCertificateDownloadUrlResponse
func (client *Client) GetTlsInspectCertificateDownloadUrlWithContext(ctx context.Context, request *GetTlsInspectCertificateDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetTlsInspectCertificateDownloadUrlResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation lists the Transport Layer Security (TLS) inspection certificate authority (CA) certificates.
//
// @param request - ListTlsInspectCACertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTlsInspectCACertificatesResponse
func (client *Client) ListTlsInspectCACertificatesWithContext(ctx context.Context, request *ListTlsInspectCACertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListTlsInspectCACertificatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an address book.
//
// Description:
//
// This operation is used to modify an address book.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param tmpReq - ModifyAddressBookRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAddressBookResponse
func (client *Client) ModifyAddressBookWithContext(ctx context.Context, tmpReq *ModifyAddressBookRequest, runtime *dara.RuntimeOptions) (_result *ModifyAddressBookResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAddressBookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetMemberUids) {
		request.AssetMemberUidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetMemberUids, dara.String("AssetMemberUids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AssetRegionResourceTypes) {
		request.AssetRegionResourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetRegionResourceTypes, dara.String("AssetRegionResourceTypes"), dara.String("json"))
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

	if !dara.IsNil(request.AssetMemberUidsShrink) {
		query["AssetMemberUids"] = request.AssetMemberUidsShrink
	}

	if !dara.IsNil(request.AssetRegionResourceTypesShrink) {
		query["AssetRegionResourceTypes"] = request.AssetRegionResourceTypesShrink
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates instance information for pay-as-you-go 2.0 users.
//
// Description:
//
// Before calling this operation, ensure that you understand the billing methods and [pricing](https://help.aliyun.com/zh/cloud-firewall/cloudfirewall/product-overview/pay-as-you-go) for the pay-as-you-go edition of Cloud Firewall.
//
// @param request - ModifyCfwInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCfwInstanceResponse
func (client *Client) ModifyCfwInstanceWithContext(ctx context.Context, request *ModifyCfwInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyCfwInstanceResponse, _err error) {
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

	if !dara.IsNil(request.UpdateList) {
		query["UpdateList"] = request.UpdateList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCfwInstance"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCfwInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// This operation modifies the configurations of an access control policy that allows, denies, or monitors traffic passing through Cloud Firewall.
//
// ## QPS limit
//
// Each user can call this operation up to 10 times per second. If the limit is exceeded, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - ModifyControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyControlPolicyResponse
func (client *Client) ModifyControlPolicyWithContext(ctx context.Context, request *ModifyControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyControlPolicyPosition is deprecated, please use Cloudfw::2017-12-07::ModifyControlPolicyPriority instead.
//
// Summary:
//
// Modifies the priority of an IPv4 access control policy for the Internet firewall. For this type of policy, the source and destination IP addresses are in IPv4 format.
//
// Description:
//
// You can call this operation to modify the priority of an IPv4 access control policy for the Internet firewall. This operation does not support modifying the priority of IPv6 access control policies.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for each user. If you exceed the limit, API calls are throttled, which can affect your business. We recommend that you call this operation within this limit.
//
// @param request - ModifyControlPolicyPositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyControlPolicyPositionResponse
func (client *Client) ModifyControlPolicyPositionWithContext(ctx context.Context, request *ModifyControlPolicyPositionRequest, runtime *dara.RuntimeOptions) (_result *ModifyControlPolicyPositionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// Description:
//
// You can call this operation to modify the priority of an access control policy. An access control policy determines whether to allow, deny, or monitor traffic that passes through Cloud Firewall.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. Exceeding this limit triggers throttling, which may affect your business. We recommend that you plan your calls accordingly.
//
// @param request - ModifyControlPolicyPriorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyControlPolicyPriorityResponse
func (client *Client) ModifyControlPolicyPriorityWithContext(ctx context.Context, request *ModifyControlPolicyPriorityRequest, runtime *dara.RuntimeOptions) (_result *ModifyControlPolicyPriorityResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the default intrusion prevention system (IPS) configuration.
//
// @param request - ModifyDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefaultIPSConfigResponse
func (client *Client) ModifyDefaultIPSConfigWithContext(ctx context.Context, request *ModifyDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefaultIPSConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a DNS firewall rule.
//
// Description:
//
// You can use this operation to modify a DNS firewall policy to accept, deny, or monitor traffic.
//
// @param request - ModifyDnsFirewallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDnsFirewallPolicyResponse
func (client *Client) ModifyDnsFirewallPolicyWithContext(ctx context.Context, request *ModifyDnsFirewallPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDnsFirewallPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a routing policy.
//
// @param request - ModifyFirewallV2RoutePolicySwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFirewallV2RoutePolicySwitchResponse
func (client *Client) ModifyFirewallV2RoutePolicySwitchWithContext(ctx context.Context, request *ModifyFirewallV2RoutePolicySwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyFirewallV2RoutePolicySwitchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of member accounts in Cloud Firewall.
//
// Description:
//
// This operation updates the attributes of member accounts in Cloud Firewall.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 10 for each user. If you exceed this limit, API calls are rate-limited. This may affect your business operations. Plan your calls accordingly.
//
// @param request - ModifyInstanceMemberAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceMemberAttributesResponse
func (client *Client) ModifyInstanceMemberAttributesWithContext(ctx context.Context, request *ModifyInstanceMemberAttributesRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceMemberAttributesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies Intrusion Prevention System (IPS) rules.
//
// @param request - ModifyIpsRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpsRulesResponse
func (client *Client) ModifyIpsRulesWithContext(ctx context.Context, request *ModifyIpsRulesRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpsRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirewallType) {
		query["FirewallType"] = request.FirewallType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RuleAction) {
		query["RuleAction"] = request.RuleAction
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
		Action:      dara.String("ModifyIpsRules"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIpsRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets Intrusion Prevention System (IPS) rules to the default settings.
//
// @param request - ModifyIpsRulesToDefaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpsRulesToDefaultResponse
func (client *Client) ModifyIpsRulesToDefaultWithContext(ctx context.Context, request *ModifyIpsRulesToDefaultRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpsRulesToDefaultResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify a NAT Firewall security access control policy.
//
// Description:
//
// This API modifies the configuration of an access control policy that allows, denies, or observes traffic passing through a NAT Firewall.
//
// @param request - ModifyNatFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNatFirewallControlPolicyResponse
func (client *Client) ModifyNatFirewallControlPolicyWithContext(ctx context.Context, request *ModifyNatFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyNatFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the priority of an access control policy for a NAT firewall.
//
// @param request - ModifyNatFirewallControlPolicyPositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNatFirewallControlPolicyPositionResponse
func (client *Client) ModifyNatFirewallControlPolicyPositionWithContext(ctx context.Context, request *ModifyNatFirewallControlPolicyPositionRequest, runtime *dara.RuntimeOptions) (_result *ModifyNatFirewallControlPolicyPositionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the operation for an object group.
//
// @param request - ModifyObjectGroupOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyObjectGroupOperationResponse
func (client *Client) ModifyObjectGroupOperationWithContext(ctx context.Context, request *ModifyObjectGroupOperationRequest, runtime *dara.RuntimeOptions) (_result *ModifyObjectGroupOperationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the strict mode for access control policies.
//
// Description:
//
// This operation enables or disables the strict mode for access control policies.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second per user. If you exceed the limit, API calls are throttled, which can affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - ModifyPolicyAdvancedConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyAdvancedConfigResponse
func (client *Client) ModifyPolicyAdvancedConfigWithContext(ctx context.Context, request *ModifyPolicyAdvancedConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyAdvancedConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a private DNS endpoint.
//
// @param request - ModifyPrivateDnsEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrivateDnsEndpointResponse
func (client *Client) ModifyPrivateDnsEndpointWithContext(ctx context.Context, request *ModifyPrivateDnsEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyPrivateDnsEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the automatic protection settings for new assets.
//
// @param request - ModifyResourceTypeAutoEnableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyResourceTypeAutoEnableResponse
func (client *Client) ModifyResourceTypeAutoEnableWithContext(ctx context.Context, request *ModifyResourceTypeAutoEnableRequest, runtime *dara.RuntimeOptions) (_result *ModifyResourceTypeAutoEnableResponse, _err error) {
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

	if !dara.IsNil(request.ResourceTypeAutoEnable) {
		query["ResourceTypeAutoEnable"] = request.ResourceTypeAutoEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyResourceTypeAutoEnable"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyResourceTypeAutoEnableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the sensitive data switch.
//
// @param request - ModifySensitiveSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySensitiveSwitchResponse
func (client *Client) ModifySensitiveSwitchWithContext(ctx context.Context, request *ModifySensitiveSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifySensitiveSwitchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the log delivery settings for Simple Log Service (SLS).
//
// Description:
//
// ## QPS limit
//
// You can call this API up to 10 times per second per user. If you exceed this limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - ModifySlsDispatchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySlsDispatchStatusResponse
func (client *Client) ModifySlsDispatchStatusWithContext(ctx context.Context, request *ModifySlsDispatchStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifySlsDispatchStatusResponse, _err error) {
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

	if !dara.IsNil(request.Site) {
		query["Site"] = request.Site
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the threat intelligence configuration.
//
// @param request - ModifyThreatIntelligenceSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyThreatIntelligenceSwitchResponse
func (client *Client) ModifyThreatIntelligenceSwitchWithContext(ctx context.Context, request *ModifyThreatIntelligenceSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyThreatIntelligenceSwitchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a VPC firewall for a transit router.
//
// @param request - ModifyTrFirewallV2ConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrFirewallV2ConfigurationResponse
func (client *Client) ModifyTrFirewallV2ConfigurationWithContext(ctx context.Context, request *ModifyTrFirewallV2ConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrFirewallV2ConfigurationResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the scope of a routing policy for a VPC firewall that is created for a Transit Router (TR).
//
// Description:
//
// You can modify the policy scope for *point-to-multipoint	- and *multipoint-to-multipoint	- scenarios, but not for *point-to-point	- scenarios.
//
// @param tmpReq - ModifyTrFirewallV2RoutePolicyScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrFirewallV2RoutePolicyScopeResponse
func (client *Client) ModifyTrFirewallV2RoutePolicyScopeWithContext(ctx context.Context, tmpReq *ModifyTrFirewallV2RoutePolicyScopeRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrFirewallV2RoutePolicyScopeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies user alert configuration.
//
// @param tmpReq - ModifyUserAlarmConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserAlarmConfigResponse
func (client *Client) ModifyUserAlarmConfigWithContext(ctx context.Context, tmpReq *ModifyUserAlarmConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserAlarmConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyUserAlarmConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ContactConfig) {
		request.ContactConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ContactConfig, dara.String("ContactConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmConfig) {
		query["AlarmConfig"] = request.AlarmConfig
	}

	if !dara.IsNil(request.AlarmLang) {
		query["AlarmLang"] = request.AlarmLang
	}

	if !dara.IsNil(request.ContactConfigShrink) {
		query["ContactConfig"] = request.ContactConfigShrink
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

	if !dara.IsNil(request.UseDefaultContact) {
		query["UseDefaultContact"] = request.UseDefaultContact
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserAlarmConfig"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserAlarmConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the intrusion prevention system (IPS) whitelist for the Internet Border.
//
// Description:
//
// ## QPS limit
//
// This API is limited to 10 queries per second (QPS) for each user. If you exceed this limit, API calls are throttled. This can affect your business. We recommend that you call the API at a reasonable rate.
//
// @param request - ModifyUserIPSWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserIPSWhitelistResponse
func (client *Client) ModifyUserIPSWhitelistWithContext(ctx context.Context, request *ModifyUserIPSWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserIPSWhitelistResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the retention period of user logs.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 per user. Calls that exceed this limit are rate-limited, which may affect your business. Plan your calls accordingly.
//
// @param request - ModifyUserSlsLogStorageTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserSlsLogStorageTimeResponse
func (client *Client) ModifyUserSlsLogStorageTimeWithContext(ctx context.Context, request *ModifyUserSlsLogStorageTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserSlsLogStorageTimeResponse, _err error) {
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

	if !dara.IsNil(request.LogVersion) {
		query["LogVersion"] = request.LogVersion
	}

	if !dara.IsNil(request.SlsRegionId) {
		query["SlsRegionId"] = request.SlsRegionId
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the ACL engine mode for a VPC firewall.
//
// Description:
//
// ## QPS limit
//
// This API is limited to 10 queries per second (QPS) per user. Calls that exceed this limit are throttled. This may affect your business. Plan your API calls accordingly.
//
// @param request - ModifyVpcFirewallAclEngineModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallAclEngineModeResponse
func (client *Client) ModifyVpcFirewallAclEngineModeWithContext(ctx context.Context, request *ModifyVpcFirewallAclEngineModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallAclEngineModeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a VPC firewall that protects traffic between network instances in a Cloud Enterprise Network (CEN) and a specified VPC.
//
// Description:
//
// This operation modifies the configuration of a VPC firewall. The VPC firewall protects traffic between network instances in a Cloud Enterprise Network (CEN) and a specified VPC. The network instances include VPCs, virtual border routers (VBRs), and Cloud Connect Network (CCN) instances. Before you call this operation, you must call the [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) operation to create a VPC firewall.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for a single user. If the limit is exceeded, API calls are throttled. This may affect your business. Please plan your API calls accordingly.
//
// @param request - ModifyVpcFirewallCenConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallCenConfigureResponse
func (client *Client) ModifyVpcFirewallCenConfigureWithContext(ctx context.Context, request *ModifyVpcFirewallCenConfigureRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallCenConfigureResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of a VPC firewall that protects traffic between network instances in a Cloud Enterprise Network (CEN) and a specified VPC.
//
// Description:
//
// This operation modifies the status of a VPC firewall. The firewall protects traffic between network instances in a Cloud Enterprise Network (CEN) and a specified Virtual Private Cloud (VPC). The network instances include VPCs, Virtual Border Routers (VBRs), and Cloud Connect Network (CCN) instances. If the firewall is enabled, it protects traffic between the network instances in the CEN and the specified VPC. If the firewall is disabled, it no longer protects this traffic.
//
// Before you call this operation, you must create a VPC firewall by calling the [CreateVpcFirewallCenConfigure](https://help.aliyun.com/document_detail/345772.html) operation.
//
// ## Limits
//
// This operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, API calls are throttled. Throttling may affect your business. Plan your calls accordingly.
//
// @param request - ModifyVpcFirewallCenSwitchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallCenSwitchStatusResponse
func (client *Client) ModifyVpcFirewallCenSwitchStatusWithContext(ctx context.Context, request *ModifyVpcFirewallCenSwitchStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallCenSwitchStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a VPC firewall that protects traffic between two VPCs connected by an Express Connect circuit.
//
// Description:
//
// This operation modifies the configuration of a VPC firewall that protects traffic between two VPCs connected by an Express Connect circuit. Before you call this operation, you must create a VPC firewall by calling the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) operation.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If the number of calls per second exceeds the limit, throttling is triggered. Throttling may affect your business. You should plan your calls accordingly.
//
// @param request - ModifyVpcFirewallConfigureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallConfigureResponse
func (client *Client) ModifyVpcFirewallConfigureWithContext(ctx context.Context, request *ModifyVpcFirewallConfigureRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallConfigureResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an access control policy for a specified VPC firewall policy group.
//
// Description:
//
// This operation modifies the configuration of an access control policy for a specified VPC firewall policy group. VPC firewall instances use different access control policies to protect Cloud Enterprise Network (CEN) instances and Express Connect circuits.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 for a single user. If the number of calls to this operation per second exceeds the limit, rate limiting is triggered. This may affect your business. Plan your calls accordingly.
//
// @param request - ModifyVpcFirewallControlPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallControlPolicyResponse
func (client *Client) ModifyVpcFirewallControlPolicyWithContext(ctx context.Context, request *ModifyVpcFirewallControlPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the priority of an access control policy in a policy group for a VPC firewall.
//
// Description:
//
// You can call this operation to modify the priority of an access control policy in a policy group for a VPC firewall.
//
// ## QPS limit
//
// The limit on the number of queries per second (QPS) for a single user is 10. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - ModifyVpcFirewallControlPolicyPositionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallControlPolicyPositionResponse
func (client *Client) ModifyVpcFirewallControlPolicyPositionWithContext(ctx context.Context, request *ModifyVpcFirewallControlPolicyPositionRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyPositionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the intrusion prevention configuration of a VPC firewall.
//
// Description:
//
// You can call this operation to modify the intrusion prevention configuration of a VPC firewall.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 per user. If the QPS limit is exceeded, API calls are throttled. This may affect your business. We recommend that you take this limit into consideration when you call this operation.
//
// @param request - ModifyVpcFirewallDefaultIPSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallDefaultIPSConfigResponse
func (client *Client) ModifyVpcFirewallDefaultIPSConfigWithContext(ctx context.Context, request *ModifyVpcFirewallDefaultIPSConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallDefaultIPSConfigResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the intrusion prevention system (IPS) whitelist for a VPC firewall.
//
// @param request - ModifyVpcFirewallIPSWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallIPSWhitelistResponse
func (client *Client) ModifyVpcFirewallIPSWhitelistWithContext(ctx context.Context, request *ModifyVpcFirewallIPSWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallIPSWhitelistResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a VPC firewall. A VPC firewall protects traffic between two VPCs that are connected by an Express Connect circuit.
//
// Description:
//
// This API call modifies the status of a VPC firewall. A VPC firewall protects traffic between two virtual private clouds (VPCs) that are connected by an Express Connect circuit. When the VPC firewall is enabled, it protects traffic between the two VPCs. When the VPC firewall is disabled, it no longer protects traffic between the two VPCs.
//
// Before you make this API call, you must create a VPC firewall using the [CreateVpcFirewallConfigure](https://help.aliyun.com/document_detail/342893.html) API call.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this API call is 10 for each Alibaba Cloud account. If you exceed the limit, your API calls are throttled, which may affect your business. Plan your API calls accordingly.
//
// @param request - ModifyVpcFirewallSwitchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcFirewallSwitchStatusResponse
func (client *Client) ModifyVpcFirewallSwitchStatusWithContext(ctx context.Context, request *ModifyVpcFirewallSwitchStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcFirewallSwitchStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables all firewall switches.
//
// Description:
//
// This operation disables all firewall switches.
//
// ## QPS limit
//
// Each user can send up to 10 queries per second (QPS). If you exceed this limit, API calls are throttled, which may affect your business. Plan your calls accordingly.
//
// @param request - PutDisableAllFwSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDisableAllFwSwitchResponse
func (client *Client) PutDisableAllFwSwitchWithContext(ctx context.Context, request *PutDisableAllFwSwitchRequest, runtime *dara.RuntimeOptions) (_result *PutDisableAllFwSwitchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a firewall switch.
//
// Description:
//
// This operation disables a firewall switch. After a firewall switch is disabled, traffic is no longer routed through Cloud Firewall.
//
// ## QPS limit
//
// The QPS limit for this operation is 10 requests per second per user. Calls that exceed this limit are throttled, which may affect your business.
//
// @param request - PutDisableFwSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDisableFwSwitchResponse
func (client *Client) PutDisableFwSwitchWithContext(ctx context.Context, request *PutDisableFwSwitchRequest, runtime *dara.RuntimeOptions) (_result *PutDisableFwSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.IpaddrList) {
		query["IpaddrList"] = request.IpaddrList
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables all firewall switches.
//
// Description:
//
// This API operation protects all public IP addresses of your Alibaba Cloud account.
//
// ## QPS limits
//
// This API operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this API operation at a reasonable rate.
//
// @param request - PutEnableAllFwSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEnableAllFwSwitchResponse
func (client *Client) PutEnableAllFwSwitchWithContext(ctx context.Context, request *PutEnableAllFwSwitchRequest, runtime *dara.RuntimeOptions) (_result *PutEnableAllFwSwitchResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enable the firewall.
//
// Description:
//
// This API enables the firewall switch. Once enabled, traffic is routed through Cloud Firewall.
//
// ## QPS limit
//
// The QPS limit for this API is 5 requests per second for a single user. If you exceed this limit, the system throttles API calls, which may affect your business.
//
// @param request - PutEnableFwSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEnableFwSwitchResponse
func (client *Client) PutEnableFwSwitchWithContext(ctx context.Context, request *PutEnableFwSwitchRequest, runtime *dara.RuntimeOptions) (_result *PutEnableFwSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.IpaddrList) {
		query["IpaddrList"] = request.IpaddrList
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases an expired instance.
//
// @param request - ReleaseExpiredInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseExpiredInstanceResponse
func (client *Client) ReleaseExpiredInstanceWithContext(ctx context.Context, request *ReleaseExpiredInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseExpiredInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go firewall.
//
// @param request - ReleasePostInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleasePostInstanceResponse
func (client *Client) ReleasePostInstanceWithContext(ctx context.Context, request *ReleasePostInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleasePostInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the hit count of a NAT firewall rule.
//
// @param request - ResetNatFirewallRuleHitCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetNatFirewallRuleHitCountResponse
func (client *Client) ResetNatFirewallRuleHitCountWithContext(ctx context.Context, request *ResetNatFirewallRuleHitCountRequest, runtime *dara.RuntimeOptions) (_result *ResetNatFirewallRuleHitCountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the hit count of a rule.
//
// Description:
//
// This operation resets the hit count of an access control policy in a VPC firewall policy group.
//
// ## QPS limit
//
// This operation is limited to 10 queries per second (QPS) per user. If you exceed this limit, API calls are throttled, which may impact your business. Plan your calls accordingly.
//
// @param request - ResetRuleHitCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetRuleHitCountResponse
func (client *Client) ResetRuleHitCountWithContext(ctx context.Context, request *ResetRuleHitCountRequest, runtime *dara.RuntimeOptions) (_result *ResetRuleHitCountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the hit count of an access control policy in a specified VPC firewall policy group to zero.
//
// Description:
//
// This operation resets the hit count of a specific access control policy in a VPC firewall policy group to zero.
//
// ## QPS limit
//
// This operation has a queries per second (QPS) limit of 10 per user. Calls that exceed this limit are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - ResetVpcFirewallRuleHitCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetVpcFirewallRuleHitCountResponse
func (client *Client) ResetVpcFirewallRuleHitCountWithContext(ctx context.Context, request *ResetVpcFirewallRuleHitCountRequest, runtime *dara.RuntimeOptions) (_result *ResetVpcFirewallRuleHitCountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables automatic protection for new assets.
//
// Description:
//
// Each Cloud Firewall instance supports up to 100 associations with TLS inspection policies.
//
// @param request - SetAutoProtectNewAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAutoProtectNewAssetsResponse
func (client *Client) SetAutoProtectNewAssetsWithContext(ctx context.Context, request *SetAutoProtectNewAssetsRequest, runtime *dara.RuntimeOptions) (_result *SetAutoProtectNewAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoProtect) {
		query["AutoProtect"] = request.AutoProtect
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
		Action:      dara.String("SetAutoProtectNewAssets"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAutoProtectNewAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchSecurityProxyResponse
func (client *Client) SwitchSecurityProxyWithContext(ctx context.Context, request *SwitchSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *SwitchSecurityProxyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the enabling status of AI-powered traffic analysis.
//
// Description:
//
// The analysis covers all data for your Cloud Firewall instance from the date of purchase.
//
// @param request - UpdateAITrafficAnalysisStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAITrafficAnalysisStatusResponse
func (client *Client) UpdateAITrafficAnalysisStatusWithContext(ctx context.Context, request *UpdateAITrafficAnalysisStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAITrafficAnalysisStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an ACK cluster connector.
//
// @param request - UpdateAckClusterConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAckClusterConnectorResponse
func (client *Client) UpdateAckClusterConnectorWithContext(ctx context.Context, request *UpdateAckClusterConnectorRequest, runtime *dara.RuntimeOptions) (_result *UpdateAckClusterConnectorResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the status of ACL check details.
//
// Description:
//
// ## QPS Limit
//
// The single-user QPS limit for this API is 10 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call this API appropriately.
//
// @param request - UpdateAclCheckDetailStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAclCheckDetailStatusResponse
func (client *Client) UpdateAclCheckDetailStatusWithContext(ctx context.Context, request *UpdateAclCheckDetailStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAclCheckDetailStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of the Internet Border firewall switch module for a pay-as-you-go user.
//
// @param request - UpdatePostpayUserInternetStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostpayUserInternetStatusResponse
func (client *Client) UpdatePostpayUserInternetStatusWithContext(ctx context.Context, request *UpdatePostpayUserInternetStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostpayUserInternetStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of a NAT border firewall for a pay-as-you-go instance.
//
// @param request - UpdatePostpayUserNatStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostpayUserNatStatusResponse
func (client *Client) UpdatePostpayUserNatStatusWithContext(ctx context.Context, request *UpdatePostpayUserNatStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostpayUserNatStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of the VPC border firewall for a pay-as-you-go user.
//
// @param request - UpdatePostpayUserVpcStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostpayUserVpcStatusResponse
func (client *Client) UpdatePostpayUserVpcStatusWithContext(ctx context.Context, request *UpdatePostpayUserVpcStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostpayUserVpcStatusResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a NAT firewall.
//
// @param request - UpdateSecurityProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityProxyResponse
func (client *Client) UpdateSecurityProxyWithContext(ctx context.Context, request *UpdateSecurityProxyRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityProxyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores an access control backup.
//
// @param request - UseAclBackupDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UseAclBackupDataResponse
func (client *Client) UseAclBackupDataWithContext(ctx context.Context, request *UseAclBackupDataRequest, runtime *dara.RuntimeOptions) (_result *UseAclBackupDataResponse, _err error) {
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
		Action:      dara.String("UseAclBackupData"),
		Version:     dara.String("2017-12-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UseAclBackupDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
