// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds members to an IP Address Manager (IPAM).
//
// Description:
//
// - Only the delegated administrator of an IPAM instance in a resource directory can perform multi-account management.
//
// - An IPAM delegated administrator can use an IPAM instance in only one region for multi-account management. A maximum of 1,000 member accounts can be added.
//
//		Notice:
//
//	If you add a folder as a member, the system counts all member accounts of the resource directory that are in the folder.
//
// - Members can be of the Folder or Account type.
//
//   - Folder: The delegated IPAM administrator can view IP usage in the IPAM effective region for all resource directory member accounts in the folder.
//
//   - Account: The delegated IPAM administrator can view IP usage in the IPAM effective region for the specified resource directory member account.
//
// - A managed member cannot share its resource discovery with the IPAM delegated administrator. The IPAM delegated administrator cannot add a member if that member has already shared its resource discovery.
//
// - Adding the first member enables the IPAM trusted service for the resource directory.
//
// @param request - AddIpamMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpamMembersResponse
func (client *Client) AddIpamMembersWithContext(ctx context.Context, request *AddIpamMembersRequest, runtime *dara.RuntimeOptions) (_result *AddIpamMembersResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Members) {
		query["Members"] = request.Members
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("AddIpamMembers"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddIpamMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provisions a CIDR block for an IPAM pool.
//
// Description:
//
// - Before provisioning a CIDR block, make sure that you have created an IPAM pool. You can call **CreateIpamPool*	- to create an IPAM pool.
//
// - If the parent pool does not have a provisioned CIDR block, the subpool does not support CIDR block provisioning.
//
// - If the parent pool has a provisioned CIDR block, the subpool can have a provisioned CIDR block, and the provisioned CIDR block must be a subset of the parent pool\\"s provisioned CIDR block.
//
// - If the parent pool has a provisioned CIDR block and also has CIDR allocations, the CIDR block provisioned for the subpool must not conflict with the existing CIDR allocations.
//
// - The request to provision a CIDR block for an IPAM pool must be initiated from the IPAM hosted region.
//
// - The CIDR block provisioned for an IPAM pool must not conflict with CIDR blocks provisioned for other pools within the same scope.
//
// - The number of CIDR blocks that can be provisioned for a pool is limited. The default maximum for a public IPv6 top-level pool is 1. The default maximum for other types of pools is 50.
//
// @param request - AddIpamPoolCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpamPoolCidrResponse
func (client *Client) AddIpamPoolCidrWithContext(ctx context.Context, request *AddIpamPoolCidrRequest, runtime *dara.RuntimeOptions) (_result *AddIpamPoolCidrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidr) {
		query["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamPoolId) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !dara.IsNil(request.NetmaskLength) {
		query["NetmaskLength"] = request.NetmaskLength
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddIpamPoolCidr"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddIpamPoolCidrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a resource discovery with an IPAM instance.
//
// Description:
//
// - You can associate a resource discovery instance with an IPAM instance only once.
//
// @param request - AssociateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateIpamResourceDiscoveryResponse
func (client *Client) AssociateIpamResourceDiscoveryWithContext(ctx context.Context, request *AssociateIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *AssociateIpamResourceDiscoveryResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamId) {
		query["IpamId"] = request.IpamId
	}

	if !dara.IsNil(request.IpamResourceDiscoveryId) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("AssociateIpamResourceDiscovery"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group of an IPAM resource.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IP Address Manager (IPAM).
//
// Description:
//
// - You can create only one IPAM with each Alibaba Cloud account in each region.
//
// - Only IPv4 IP addresses can be allocated.
//
// - When you create an IPAM instance:
//
//   - If there is no custom resource discovery in the region, the system creates a default resource discovery associated with the IPAM instance.
//
//   - If there is a custom resource discovery in the region, the system converts it to a default resource discovery and associates it with the IPAM instance.
//
// @param request - CreateIpamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamResponse
func (client *Client) CreateIpamWithContext(ctx context.Context, request *CreateIpamRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamDescription) {
		query["IpamDescription"] = request.IpamDescription
	}

	if !dara.IsNil(request.IpamName) {
		query["IpamName"] = request.IpamName
	}

	if !dara.IsNil(request.OperatingRegionList) {
		query["OperatingRegionList"] = request.OperatingRegionList
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIpam"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an IPAM address pool.
//
// Description:
//
// - The default maximum number of public IPv6 top-level pools per ISP type per region is 1.
//
// @param request - CreateIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamPoolResponse
func (client *Client) CreateIpamPoolWithContext(ctx context.Context, request *CreateIpamPoolRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationDefaultCidrMask) {
		query["AllocationDefaultCidrMask"] = request.AllocationDefaultCidrMask
	}

	if !dara.IsNil(request.AllocationMaxCidrMask) {
		query["AllocationMaxCidrMask"] = request.AllocationMaxCidrMask
	}

	if !dara.IsNil(request.AllocationMinCidrMask) {
		query["AllocationMinCidrMask"] = request.AllocationMinCidrMask
	}

	if !dara.IsNil(request.AutoImport) {
		query["AutoImport"] = request.AutoImport
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.IpamPoolDescription) {
		query["IpamPoolDescription"] = request.IpamPoolDescription
	}

	if !dara.IsNil(request.IpamPoolName) {
		query["IpamPoolName"] = request.IpamPoolName
	}

	if !dara.IsNil(request.IpamScopeId) {
		query["IpamScopeId"] = request.IpamScopeId
	}

	if !dara.IsNil(request.Ipv6Isp) {
		query["Ipv6Isp"] = request.Ipv6Isp
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolRegionId) {
		query["PoolRegionId"] = request.PoolRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceIpamPoolId) {
		query["SourceIpamPoolId"] = request.SourceIpamPoolId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIpamPool"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpamPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom reserved CIDR block from an IPAM pool.
//
// Description:
//
// - Before you create a custom reserved CIDR block, ensure that you have created an IPAM pool and added a CIDR block to it. You can call the **CreateIpamPool*	- operation to create an IPAM pool and the **AddIpamPoolCidr*	- operation to add a CIDR block to the pool.
//
// - When you specify the Cidr or CidrMask parameter to create a custom reserved CIDR block, the mask must be within the range specified for the IPAM pool.
//
// - If an IPAM pool has a region attribute, the request to create a custom reserved CIDR block must be initiated from the region where the pool is located.
//
// - The custom reserved CIDR block must not conflict with existing CIDR block allocations in the IPAM pool.
//
// @param request - CreateIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamPoolAllocationResponse
func (client *Client) CreateIpamPoolAllocationWithContext(ctx context.Context, request *CreateIpamPoolAllocationRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamPoolAllocationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidr) {
		query["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.CidrMask) {
		query["CidrMask"] = request.CidrMask
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamPoolAllocationDescription) {
		query["IpamPoolAllocationDescription"] = request.IpamPoolAllocationDescription
	}

	if !dara.IsNil(request.IpamPoolAllocationName) {
		query["IpamPoolAllocationName"] = request.IpamPoolAllocationName
	}

	if !dara.IsNil(request.IpamPoolId) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIpamPoolAllocation"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpamPoolAllocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource discovery instance of a custom type.
//
// Description:
//
// - Each Alibaba Cloud account can have only one resource discovery instance in each region.
//
// - This operation creates only resource discovery instances of a custom type.
//
// @param request - CreateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamResourceDiscoveryResponse
func (client *Client) CreateIpamResourceDiscoveryWithContext(ctx context.Context, request *CreateIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamResourceDiscoveryResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamResourceDiscoveryDescription) {
		query["IpamResourceDiscoveryDescription"] = request.IpamResourceDiscoveryDescription
	}

	if !dara.IsNil(request.IpamResourceDiscoveryName) {
		query["IpamResourceDiscoveryName"] = request.IpamResourceDiscoveryName
	}

	if !dara.IsNil(request.OperatingRegionList) {
		query["OperatingRegionList"] = request.OperatingRegionList
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIpamResourceDiscovery"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates scopes for IPAM to manage private and public IP addresses.
//
// @param request - CreateIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamScopeResponse
func (client *Client) CreateIpamScopeWithContext(ctx context.Context, request *CreateIpamScopeRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamScopeResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamId) {
		query["IpamId"] = request.IpamId
	}

	if !dara.IsNil(request.IpamScopeDescription) {
		query["IpamScopeDescription"] = request.IpamScopeDescription
	}

	if !dara.IsNil(request.IpamScopeName) {
		query["IpamScopeName"] = request.IpamScopeName
	}

	if !dara.IsNil(request.IpamScopeType) {
		query["IpamScopeType"] = request.IpamScopeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIpamScope"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpamScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IPAM instance.
//
// Description:
//
// ## Prerequisites
//
// - Before you delete an IPAM instance, ensure that all IPAM pools in the instance are deleted. You can call the **DeleteIpamPool*	- operation to delete the IPAM pools.
//
// - Before you delete an IPAM instance, ensure that all custom IPAM scopes in the instance are deleted. You can call the **DeleteIpamScope*	- operation to delete the IPAM scopes.
//
// - Before you delete an IPAM instance, ensure that the default resource discovery instance is not shared.
//
// - Before you delete an IPAM instance, ensure that no shared resource discovery instances are associated with the IPAM instance.
//
// @param request - DeleteIpamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamResponse
func (client *Client) DeleteIpamWithContext(ctx context.Context, request *DeleteIpamRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamId) {
		query["IpamId"] = request.IpamId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("DeleteIpam"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IPAM pool instance.
//
// Description:
//
// ### Usage notes
//
// - Before deleting a parent pool, make sure that all subpools under the parent pool have been deleted.
//
// - When a parent pool has an effective region configured and has addresses that have already been allocated, the parent pool cannot be deleted.
//
// - When a subpool has an effective region configured and has addresses that have already been allocated, the subpool cannot be deleted.
//
// - When a pool has a sharing relationship, the pool cannot be deleted.
//
// @param request - DeleteIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolResponse
func (client *Client) DeleteIpamPoolWithContext(ctx context.Context, request *DeleteIpamPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamPoolResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamPoolId) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("DeleteIpamPool"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpamPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a CIDR allocation from an IP Address Management (IPAM) address pool. Supported allocation types include virtual private cloud (VPC) and custom allocation.
//
// @param request - DeleteIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolAllocationResponse
func (client *Client) DeleteIpamPoolAllocationWithContext(ctx context.Context, request *DeleteIpamPoolAllocationRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamPoolAllocationResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamPoolAllocationId) {
		query["IpamPoolAllocationId"] = request.IpamPoolAllocationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIpamPoolAllocation"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpamPoolAllocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a provisioned CIDR block from an IP Address Manager (IPAM) pool.
//
// Description:
//
// - If CIDR blocks are provisioned in both a parent pool and its sub-pools, delete the CIDR blocks from the sub-pools before you delete the CIDR block from the parent pool.
//
// - If a CIDR block is provisioned only in a parent pool, you can delete the CIDR block directly from the parent pool.
//
// - If allocations exist from the provisioned CIDR block, delete the allocations before you delete the CIDR block.
//
// - Requests to delete a provisioned CIDR block from an IPAM pool must be sent from the region where the IPAM is deployed.
//
// @param request - DeleteIpamPoolCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolCidrResponse
func (client *Client) DeleteIpamPoolCidrWithContext(ctx context.Context, request *DeleteIpamPoolCidrRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamPoolCidrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidr) {
		query["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamPoolId) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIpamPoolCidr"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpamPoolCidrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource discovery instance.
//
// Description:
//
// - A resource discovery instance cannot be deleted if it is shared.
//
// @param request - DeleteIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamResourceDiscoveryResponse
func (client *Client) DeleteIpamResourceDiscoveryWithContext(ctx context.Context, request *DeleteIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamResourceDiscoveryResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamResourceDiscoveryId) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("DeleteIpamResourceDiscovery"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IPAM scope.
//
// Description:
//
// ### Usage notes
//
// - You cannot delete the two default IPAM scopes that the system automatically creates.
//
// - Before you delete a custom IPAM scope, ensure that all IPAM pools in the scope are deleted. You can call the **DeleteIpamPool*	- operation to delete an IPAM pool.
//
// @param request - DeleteIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamScopeResponse
func (client *Client) DeleteIpamScopeWithContext(ctx context.Context, request *DeleteIpamScopeRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamScopeResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamScopeId) {
		query["IpamScopeId"] = request.IpamScopeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("DeleteIpamScope"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpamScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a resource discovery from an IP Address Manager (IPAM) instance.
//
// @param request - DissociateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateIpamResourceDiscoveryResponse
func (client *Client) DissociateIpamResourceDiscoveryWithContext(ctx context.Context, request *DissociateIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *DissociateIpamResourceDiscoveryResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamId) {
		query["IpamId"] = request.IpamId
	}

	if !dara.IsNil(request.IpamResourceDiscoveryId) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("DissociateIpamResourceDiscovery"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specified CIDR block allocation in an IPAM pool.
//
// @param request - GetIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpamPoolAllocationResponse
func (client *Client) GetIpamPoolAllocationWithContext(ctx context.Context, request *GetIpamPoolAllocationRequest, runtime *dara.RuntimeOptions) (_result *GetIpamPoolAllocationResponse, _err error) {
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
		Action:      dara.String("GetIpamPoolAllocation"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIpamPoolAllocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an available CIDR block from an IPAM pool.
//
// @param request - GetIpamPoolNextAvailableCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpamPoolNextAvailableCidrResponse
func (client *Client) GetIpamPoolNextAvailableCidrWithContext(ctx context.Context, request *GetIpamPoolNextAvailableCidrRequest, runtime *dara.RuntimeOptions) (_result *GetIpamPoolNextAvailableCidrResponse, _err error) {
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
		Action:      dara.String("GetIpamPoolNextAvailableCidr"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIpamPoolNextAvailableCidrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the status of the IPAM service.
//
// @param request - GetVpcIpamServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcIpamServiceStatusResponse
func (client *Client) GetVpcIpamServiceStatusWithContext(ctx context.Context, request *GetVpcIpamServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetVpcIpamServiceStatusResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("GetVpcIpamServiceStatus"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVpcIpamServiceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the IP addresses used by discovered resources in a VPC or vSwitch.
//
// Description:
//
// Supported query combinations:
//
// - `VpcId` only
//
// - `VSwitchId` only
//
// - `VpcId` + `VSwitchId`
//
// - `VpcId` + `CidrBlock`
//
// - `VSwitchId` + `CidrBlock`
//
// - `VpcId` + `VSwitchId` + `CidrBlock`
//
// @param request - ListIpamDiscoveredIpAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamDiscoveredIpAddressesResponse
func (client *Client) ListIpamDiscoveredIpAddressesWithContext(ctx context.Context, request *ListIpamDiscoveredIpAddressesRequest, runtime *dara.RuntimeOptions) (_result *ListIpamDiscoveredIpAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidr) {
		query["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.IpamResourceDiscoveryId) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("ListIpamDiscoveredIpAddresses"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamDiscoveredIpAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource information under a resource discovery.
//
// @param request - ListIpamDiscoveredResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamDiscoveredResourceResponse
func (client *Client) ListIpamDiscoveredResourceWithContext(ctx context.Context, request *ListIpamDiscoveredResourceRequest, runtime *dara.RuntimeOptions) (_result *ListIpamDiscoveredResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpamResourceDiscoveryId) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("ListIpamDiscoveredResource"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamDiscoveredResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the members managed by the IPAM trusted service.
//
// @param request - ListIpamMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamMembersResponse
func (client *Client) ListIpamMembersWithContext(ctx context.Context, request *ListIpamMembersRequest, runtime *dara.RuntimeOptions) (_result *ListIpamMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MemberIds) {
		query["MemberIds"] = request.MemberIds
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("ListIpamMembers"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations in an IPAM pool.
//
// @param request - ListIpamPoolAllocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolAllocationsResponse
func (client *Client) ListIpamPoolAllocationsWithContext(ctx context.Context, request *ListIpamPoolAllocationsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamPoolAllocationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidr) {
		query["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.IpamPoolAllocationIds) {
		query["IpamPoolAllocationIds"] = request.IpamPoolAllocationIds
	}

	if !dara.IsNil(request.IpamPoolAllocationName) {
		query["IpamPoolAllocationName"] = request.IpamPoolAllocationName
	}

	if !dara.IsNil(request.IpamPoolId) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIpamPoolAllocations"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamPoolAllocationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the provisioned CIDR blocks of an IPAM pool.
//
// @param request - ListIpamPoolCidrsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolCidrsResponse
func (client *Client) ListIpamPoolCidrsWithContext(ctx context.Context, request *ListIpamPoolCidrsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamPoolCidrsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidr) {
		query["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.IpamPoolId) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIpamPoolCidrs"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamPoolCidrsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IPAM pools.
//
// @param request - ListIpamPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolsResponse
func (client *Client) ListIpamPoolsWithContext(ctx context.Context, request *ListIpamPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamPoolsResponse, _err error) {
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

	if !dara.IsNil(request.IpamPoolIds) {
		query["IpamPoolIds"] = request.IpamPoolIds
	}

	if !dara.IsNil(request.IpamPoolName) {
		query["IpamPoolName"] = request.IpamPoolName
	}

	if !dara.IsNil(request.IpamScopeId) {
		query["IpamScopeId"] = request.IpamScopeId
	}

	if !dara.IsNil(request.Ipv6Isp) {
		query["Ipv6Isp"] = request.Ipv6Isp
	}

	if !dara.IsNil(request.IsShared) {
		query["IsShared"] = request.IsShared
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PoolRegionId) {
		query["PoolRegionId"] = request.PoolRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceIpamPoolId) {
		query["SourceIpamPoolId"] = request.SourceIpamPoolId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIpamPools"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamPoolsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resources within an IPAM scope.
//
// @param request - ListIpamResourceCidrsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceCidrsResponse
func (client *Client) ListIpamResourceCidrsWithContext(ctx context.Context, request *ListIpamResourceCidrsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamResourceCidrsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpamPoolId) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !dara.IsNil(request.IpamScopeId) {
		query["IpamScopeId"] = request.IpamScopeId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIpamResourceCidrs"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamResourceCidrsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of IPAM resource discovery instances.
//
// @param request - ListIpamResourceDiscoveriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceDiscoveriesResponse
func (client *Client) ListIpamResourceDiscoveriesWithContext(ctx context.Context, request *ListIpamResourceDiscoveriesRequest, runtime *dara.RuntimeOptions) (_result *ListIpamResourceDiscoveriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpamResourceDiscoveryIds) {
		query["IpamResourceDiscoveryIds"] = request.IpamResourceDiscoveryIds
	}

	if !dara.IsNil(request.IpamResourceDiscoveryName) {
		query["IpamResourceDiscoveryName"] = request.IpamResourceDiscoveryName
	}

	if !dara.IsNil(request.IsShared) {
		query["IsShared"] = request.IsShared
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIpamResourceDiscoveries"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamResourceDiscoveriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the associations between resource discoveries and IP Address Managers (IPAMs).
//
// @param request - ListIpamResourceDiscoveryAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceDiscoveryAssociationsResponse
func (client *Client) ListIpamResourceDiscoveryAssociationsWithContext(ctx context.Context, request *ListIpamResourceDiscoveryAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamResourceDiscoveryAssociationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpamId) {
		query["IpamId"] = request.IpamId
	}

	if !dara.IsNil(request.IpamResourceDiscoveryId) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("ListIpamResourceDiscoveryAssociations"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamResourceDiscoveryAssociationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IPAM scopes.
//
// @param request - ListIpamScopesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamScopesResponse
func (client *Client) ListIpamScopesWithContext(ctx context.Context, request *ListIpamScopesRequest, runtime *dara.RuntimeOptions) (_result *ListIpamScopesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpamId) {
		query["IpamId"] = request.IpamId
	}

	if !dara.IsNil(request.IpamScopeIds) {
		query["IpamScopeIds"] = request.IpamScopeIds
	}

	if !dara.IsNil(request.IpamScopeName) {
		query["IpamScopeName"] = request.IpamScopeName
	}

	if !dara.IsNil(request.IpamScopeType) {
		query["IpamScopeType"] = request.IpamScopeType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIpamScopes"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamScopesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more IPAMs.
//
// @param request - ListIpamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamsResponse
func (client *Client) ListIpamsWithContext(ctx context.Context, request *ListIpamsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpamIds) {
		query["IpamIds"] = request.IpamIds
	}

	if !dara.IsNil(request.IpamName) {
		query["IpamName"] = request.IpamName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIpams"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIpamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are associated with resources.
//
// Description:
//
// ### Usage notes
//
// - You must specify at least **ResourceId.N*	- or **Tag.N*	- (**Tag.N.Key*	- and **Tag.N.Value**) in a request to identify the resources to query.
//
// - **Tag.N*	- is a resource tag that consists of a key-value pair. If you specify only **Tag.N.Key**, all tag values associated with the tag key are returned. An error is returned if you specify only **Tag.N.Value**.
//
// - If you specify both **Tag.N*	- and **ResourceId.N**, the query returns only the resources that are specified by **ResourceId.N*	- and are associated with all the specified tag key-value pairs.
//
// - If you specify multiple tag key-value pairs, the query returns only resources that are associated with all the specified key-value pairs.
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
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2023-02-28"),
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
// Activates the IP Address Management (IPAM) service.
//
// @param request - OpenVpcIpamServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenVpcIpamServiceResponse
func (client *Client) OpenVpcIpamServiceWithContext(ctx context.Context, request *OpenVpcIpamServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenVpcIpamServiceResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("OpenVpcIpamService"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenVpcIpamServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes members from the IPAM trusted service.
//
// Description:
//
// - If the delegated IPAM administrator removes the last member, the IPAM trusted service is disabled for the resource directory.
//
// @param request - RemoveIpamMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveIpamMembersResponse
func (client *Client) RemoveIpamMembersWithContext(ctx context.Context, request *RemoveIpamMembersRequest, runtime *dara.RuntimeOptions) (_result *RemoveIpamMembersResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Members) {
		query["Members"] = request.Members
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("RemoveIpamMembers"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveIpamMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a tag to a resource.
//
// Description:
//
// ### [](#)Usage notes
//
// Tags are used to classify instances. Each tag consists of a key-value pair. Before you use tags, take note of the following items:
//
//   - Each tag key that is added to an instance must be unique.
//
//   - You cannot create tags without adding them to instances. All tags must be added to instances.
//
//   - You can add at most 20 tags to each instance. Before you add a tag to an instance, the system automatically checks the number of existing tags. An error message is returned if the maximum number of tags is reached.
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
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2023-02-28"),
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
// Removes a tag from a resource.
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2023-02-28"),
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
// Modifies an IP Address Management (IPAM) instance.
//
// Description:
//
// - The managed region of an IPAM instance cannot be removed.
//
// @param request - UpdateIpamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamResponse
func (client *Client) UpdateIpamWithContext(ctx context.Context, request *UpdateIpamRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddOperatingRegion) {
		query["AddOperatingRegion"] = request.AddOperatingRegion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamDescription) {
		query["IpamDescription"] = request.IpamDescription
	}

	if !dara.IsNil(request.IpamId) {
		query["IpamId"] = request.IpamId
	}

	if !dara.IsNil(request.IpamName) {
		query["IpamName"] = request.IpamName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoveOperatingRegion) {
		query["RemoveOperatingRegion"] = request.RemoveOperatingRegion
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
		Action:      dara.String("UpdateIpam"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIpamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information of an IPAM pool.
//
// @param request - UpdateIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamPoolResponse
func (client *Client) UpdateIpamPoolWithContext(ctx context.Context, request *UpdateIpamPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationDefaultCidrMask) {
		query["AllocationDefaultCidrMask"] = request.AllocationDefaultCidrMask
	}

	if !dara.IsNil(request.AllocationMaxCidrMask) {
		query["AllocationMaxCidrMask"] = request.AllocationMaxCidrMask
	}

	if !dara.IsNil(request.AllocationMinCidrMask) {
		query["AllocationMinCidrMask"] = request.AllocationMinCidrMask
	}

	if !dara.IsNil(request.AutoImport) {
		query["AutoImport"] = request.AutoImport
	}

	if !dara.IsNil(request.ClearAllocationDefaultCidrMask) {
		query["ClearAllocationDefaultCidrMask"] = request.ClearAllocationDefaultCidrMask
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamPoolDescription) {
		query["IpamPoolDescription"] = request.IpamPoolDescription
	}

	if !dara.IsNil(request.IpamPoolId) {
		query["IpamPoolId"] = request.IpamPoolId
	}

	if !dara.IsNil(request.IpamPoolName) {
		query["IpamPoolName"] = request.IpamPoolName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("UpdateIpamPool"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIpamPoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a CIDR allocation from an IPAM address pool.
//
// @param request - UpdateIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamPoolAllocationResponse
func (client *Client) UpdateIpamPoolAllocationWithContext(ctx context.Context, request *UpdateIpamPoolAllocationRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamPoolAllocationResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamPoolAllocationDescription) {
		query["IpamPoolAllocationDescription"] = request.IpamPoolAllocationDescription
	}

	if !dara.IsNil(request.IpamPoolAllocationId) {
		query["IpamPoolAllocationId"] = request.IpamPoolAllocationId
	}

	if !dara.IsNil(request.IpamPoolAllocationName) {
		query["IpamPoolAllocationName"] = request.IpamPoolAllocationName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIpamPoolAllocation"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIpamPoolAllocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a resource discovery instance.
//
// Description:
//
// - You can add or remove operating regions only for custom resource discovery instances.
//
// - When you remove an operating region from a resource discovery instance, you cannot remove the managed region of the resource discovery instance.
//
// @param request - UpdateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamResourceDiscoveryResponse
func (client *Client) UpdateIpamResourceDiscoveryWithContext(ctx context.Context, request *UpdateIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamResourceDiscoveryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddOperatingRegion) {
		query["AddOperatingRegion"] = request.AddOperatingRegion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamResourceDiscoveryDescription) {
		query["IpamResourceDiscoveryDescription"] = request.IpamResourceDiscoveryDescription
	}

	if !dara.IsNil(request.IpamResourceDiscoveryId) {
		query["IpamResourceDiscoveryId"] = request.IpamResourceDiscoveryId
	}

	if !dara.IsNil(request.IpamResourceDiscoveryName) {
		query["IpamResourceDiscoveryName"] = request.IpamResourceDiscoveryName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoveOperatingRegion) {
		query["RemoveOperatingRegion"] = request.RemoveOperatingRegion
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
		Action:      dara.String("UpdateIpamResourceDiscovery"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIpamResourceDiscoveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information of an IPAM scope.
//
// @param request - UpdateIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamScopeResponse
func (client *Client) UpdateIpamScopeWithContext(ctx context.Context, request *UpdateIpamScopeRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamScopeResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IpamScopeDescription) {
		query["IpamScopeDescription"] = request.IpamScopeDescription
	}

	if !dara.IsNil(request.IpamScopeId) {
		query["IpamScopeId"] = request.IpamScopeId
	}

	if !dara.IsNil(request.IpamScopeName) {
		query["IpamScopeName"] = request.IpamScopeName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("UpdateIpamScope"),
		Version:     dara.String("2023-02-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIpamScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
