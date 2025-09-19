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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("vpcipam"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Provisions a CIDR block to an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  Before you provision a CIDR block, make sure that an IPAM pool is created. You can call the **CreateIpamPool*	- operation to create an IPAM pool.
//
//		- If no CIDR block is provisioned to a parent pool, you cannot provision CIDR blocks to its subpools.
//
//		- If a CIDR block is provisioned to a parent pool, you can provision CIDR blocks to its subpools and the CIDR blocks must be subsets of the CIDR block provisioned to the parent pool.
//
//		- If a CIDR block is provisioned to a parent pool and allocations are created, CIDR blocks provisioned to its subpools cannot overlap with existing allocated CIDR blocks.
//
//		- You can provision CIDR blocks to a pool only in the region where the IPAM is hosted.
//
//		- CIDR blocks provisioned to an IPAM pool cannot overlap with the CIDR blocks provisioned to other pools in the same scope.
//
//		- A maximum of 1 CIDR block can be provisioned to a public IPv6 top-level pool, while up to 50 CIDR blocks can be provisioned to other types of address pools.
//
// @param request - AddIpamPoolCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpamPoolCidrResponse
func (client *Client) AddIpamPoolCidrWithOptions(request *AddIpamPoolCidrRequest, runtime *dara.RuntimeOptions) (_result *AddIpamPoolCidrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provisions a CIDR block to an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  Before you provision a CIDR block, make sure that an IPAM pool is created. You can call the **CreateIpamPool*	- operation to create an IPAM pool.
//
//		- If no CIDR block is provisioned to a parent pool, you cannot provision CIDR blocks to its subpools.
//
//		- If a CIDR block is provisioned to a parent pool, you can provision CIDR blocks to its subpools and the CIDR blocks must be subsets of the CIDR block provisioned to the parent pool.
//
//		- If a CIDR block is provisioned to a parent pool and allocations are created, CIDR blocks provisioned to its subpools cannot overlap with existing allocated CIDR blocks.
//
//		- You can provision CIDR blocks to a pool only in the region where the IPAM is hosted.
//
//		- CIDR blocks provisioned to an IPAM pool cannot overlap with the CIDR blocks provisioned to other pools in the same scope.
//
//		- A maximum of 1 CIDR block can be provisioned to a public IPv6 top-level pool, while up to 50 CIDR blocks can be provisioned to other types of address pools.
//
// @param request - AddIpamPoolCidrRequest
//
// @return AddIpamPoolCidrResponse
func (client *Client) AddIpamPoolCidr(request *AddIpamPoolCidrRequest) (_result *AddIpamPoolCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddIpamPoolCidrResponse{}
	_body, _err := client.AddIpamPoolCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates resource discovery with an IPAM instance.
//
// Description:
//
//	The specified resource discovery instance can only be associated with one IPAM instance and associations cannot be duplicated.
//
// @param request - AssociateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateIpamResourceDiscoveryResponse
func (client *Client) AssociateIpamResourceDiscoveryWithOptions(request *AssociateIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *AssociateIpamResourceDiscoveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates resource discovery with an IPAM instance.
//
// Description:
//
//	The specified resource discovery instance can only be associated with one IPAM instance and associations cannot be duplicated.
//
// @param request - AssociateIpamResourceDiscoveryRequest
//
// @return AssociateIpamResourceDiscoveryResponse
func (client *Client) AssociateIpamResourceDiscovery(request *AssociateIpamResourceDiscoveryRequest) (_result *AssociateIpamResourceDiscoveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateIpamResourceDiscoveryResponse{}
	_body, _err := client.AssociateIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
func (client *Client) CreateIpamWithOptions(request *CreateIpamRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateIpamResponse
func (client *Client) CreateIpam(request *CreateIpamRequest) (_result *CreateIpamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIpamResponse{}
	_body, _err := client.CreateIpamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an IP Address Manager (IPAM) pool.
//
// Description:
//
// The number of public IPv6 IPAM top pool for a specific ISP that a user is allowed to create per region is limited to 1.
//
// @param request - CreateIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamPoolResponse
func (client *Client) CreateIpamPoolWithOptions(request *CreateIpamPoolRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamPoolResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IP Address Manager (IPAM) pool.
//
// Description:
//
// The number of public IPv6 IPAM top pool for a specific ISP that a user is allowed to create per region is limited to 1.
//
// @param request - CreateIpamPoolRequest
//
// @return CreateIpamPoolResponse
func (client *Client) CreateIpamPool(request *CreateIpamPoolRequest) (_result *CreateIpamPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIpamPoolResponse{}
	_body, _err := client.CreateIpamPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reserves a custom CIDR block from an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  Before you reserve a custom CIDR block, make sure that an IPAM pool is created and CIDR blocks are added to the pool. You can call **CreateIpamPool*	- to create an IPAM pool and call **AddIpamPoolCidr*	- to add CIDR blocks to the pool.
//
//		- When you specify Cidr or CidrMask to reserve a custom CIDR block, the mask must fall within the range specified by the IPAM pool.
//
//		- If the IPAM pool has the region attribute, you must reserve a custom CIDR block in the region to which the IPAM pool belongs.
//
//		- The custom CIDR block that you want to reserve cannot overlap with existing CIDR blocks created from the IPAM pool.
//
// @param request - CreateIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamPoolAllocationResponse
func (client *Client) CreateIpamPoolAllocationWithOptions(request *CreateIpamPoolAllocationRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamPoolAllocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reserves a custom CIDR block from an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  Before you reserve a custom CIDR block, make sure that an IPAM pool is created and CIDR blocks are added to the pool. You can call **CreateIpamPool*	- to create an IPAM pool and call **AddIpamPoolCidr*	- to add CIDR blocks to the pool.
//
//		- When you specify Cidr or CidrMask to reserve a custom CIDR block, the mask must fall within the range specified by the IPAM pool.
//
//		- If the IPAM pool has the region attribute, you must reserve a custom CIDR block in the region to which the IPAM pool belongs.
//
//		- The custom CIDR block that you want to reserve cannot overlap with existing CIDR blocks created from the IPAM pool.
//
// @param request - CreateIpamPoolAllocationRequest
//
// @return CreateIpamPoolAllocationResponse
func (client *Client) CreateIpamPoolAllocation(request *CreateIpamPoolAllocationRequest) (_result *CreateIpamPoolAllocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIpamPoolAllocationResponse{}
	_body, _err := client.CreateIpamPoolAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom resource discovery instance.
//
// Description:
//
//	  Each Alibaba Cloud account can create only one resource discovery instance in each region.
//
//		- You can create only custom resource discovery instances.
//
// @param request - CreateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamResourceDiscoveryResponse
func (client *Client) CreateIpamResourceDiscoveryWithOptions(request *CreateIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamResourceDiscoveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom resource discovery instance.
//
// Description:
//
//	  Each Alibaba Cloud account can create only one resource discovery instance in each region.
//
//		- You can create only custom resource discovery instances.
//
// @param request - CreateIpamResourceDiscoveryRequest
//
// @return CreateIpamResourceDiscoveryResponse
func (client *Client) CreateIpamResourceDiscovery(request *CreateIpamResourceDiscoveryRequest) (_result *CreateIpamResourceDiscoveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIpamResourceDiscoveryResponse{}
	_body, _err := client.CreateIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a public scope and private scope to respectively manage public and private IP addresses.
//
// @param request - CreateIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpamScopeResponse
func (client *Client) CreateIpamScopeWithOptions(request *CreateIpamScopeRequest, runtime *dara.RuntimeOptions) (_result *CreateIpamScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public scope and private scope to respectively manage public and private IP addresses.
//
// @param request - CreateIpamScopeRequest
//
// @return CreateIpamScopeResponse
func (client *Client) CreateIpamScope(request *CreateIpamScopeRequest) (_result *CreateIpamScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIpamScopeResponse{}
	_body, _err := client.CreateIpamScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM).
//
// Description:
//
// ## [](#)Prerequisites
//
//   - Before you delete an IPAM, make sure that all IPAM pools of the IPAM are deleted. You can call **DeleteIpamPool*	- to delete IPAM pools.
//
//   - Before you delete an IPAM, make sure that all IPAM scopes of the IPAM are deleted. You can call **DeleteIpamScope*	- to delete IPAM scopes.
//
// @param request - DeleteIpamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamResponse
func (client *Client) DeleteIpamWithOptions(request *DeleteIpamRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM).
//
// Description:
//
// ## [](#)Prerequisites
//
//   - Before you delete an IPAM, make sure that all IPAM pools of the IPAM are deleted. You can call **DeleteIpamPool*	- to delete IPAM pools.
//
//   - Before you delete an IPAM, make sure that all IPAM scopes of the IPAM are deleted. You can call **DeleteIpamScope*	- to delete IPAM scopes.
//
// @param request - DeleteIpamRequest
//
// @return DeleteIpamResponse
func (client *Client) DeleteIpam(request *DeleteIpamRequest) (_result *DeleteIpamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpamResponse{}
	_body, _err := client.DeleteIpamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM) scope.
//
// Description:
//
// ### [](#)Usage notes
//
//   - Before you delete a parent pool, make sure that all subpools of the parent pool are deleted.
//
//   - If an effective region is specified for a parent pool and IP addresses are allocated from the parent pool, you cannot delete the parent pool.
//
//   - If an effective region is specified for a subpool and IP addresses are allocated from the subpool, you cannot delete the subpool.
//
// @param request - DeleteIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolResponse
func (client *Client) DeleteIpamPoolWithOptions(request *DeleteIpamPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamPoolResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM) scope.
//
// Description:
//
// ### [](#)Usage notes
//
//   - Before you delete a parent pool, make sure that all subpools of the parent pool are deleted.
//
//   - If an effective region is specified for a parent pool and IP addresses are allocated from the parent pool, you cannot delete the parent pool.
//
//   - If an effective region is specified for a subpool and IP addresses are allocated from the subpool, you cannot delete the subpool.
//
// @param request - DeleteIpamPoolRequest
//
// @return DeleteIpamPoolResponse
func (client *Client) DeleteIpamPool(request *DeleteIpamPoolRequest) (_result *DeleteIpamPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpamPoolResponse{}
	_body, _err := client.DeleteIpamPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom reserved CIDR block from an IP Address Manager (IPAM) pool.
//
// @param request - DeleteIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolAllocationResponse
func (client *Client) DeleteIpamPoolAllocationWithOptions(request *DeleteIpamPoolAllocationRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamPoolAllocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom reserved CIDR block from an IP Address Manager (IPAM) pool.
//
// @param request - DeleteIpamPoolAllocationRequest
//
// @return DeleteIpamPoolAllocationResponse
func (client *Client) DeleteIpamPoolAllocation(request *DeleteIpamPoolAllocationRequest) (_result *DeleteIpamPoolAllocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpamPoolAllocationResponse{}
	_body, _err := client.DeleteIpamPoolAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a CIDR block provisioned to an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  If CIDR blocks are provisioned to a parent pool and its subpools, you must first delete the CIDR blocks provisioned to the subpools before you delete the ones provisioned to the parent pool.
//
//		- If CIDR blocks are provisioned only to the parent pool, directly delete them.
//
//		- If CIDR blocks are allocated from provisioned ones, you must first delete the allocated CIDR blocks before you delete the provisioned ones.
//
//		- You can delete CIDR blocks provisioned to an IPAM pool only in the region where the IPAM is hosted.
//
// @param request - DeleteIpamPoolCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamPoolCidrResponse
func (client *Client) DeleteIpamPoolCidrWithOptions(request *DeleteIpamPoolCidrRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamPoolCidrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a CIDR block provisioned to an IP Address Manager (IPAM) pool.
//
// Description:
//
//	  If CIDR blocks are provisioned to a parent pool and its subpools, you must first delete the CIDR blocks provisioned to the subpools before you delete the ones provisioned to the parent pool.
//
//		- If CIDR blocks are provisioned only to the parent pool, directly delete them.
//
//		- If CIDR blocks are allocated from provisioned ones, you must first delete the allocated CIDR blocks before you delete the provisioned ones.
//
//		- You can delete CIDR blocks provisioned to an IPAM pool only in the region where the IPAM is hosted.
//
// @param request - DeleteIpamPoolCidrRequest
//
// @return DeleteIpamPoolCidrResponse
func (client *Client) DeleteIpamPoolCidr(request *DeleteIpamPoolCidrRequest) (_result *DeleteIpamPoolCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpamPoolCidrResponse{}
	_body, _err := client.DeleteIpamPoolCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom resource discovery instance.
//
// Description:
//
//	If a resource discovery instance is shared, it cannot be deleted.
//
// @param request - DeleteIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamResourceDiscoveryResponse
func (client *Client) DeleteIpamResourceDiscoveryWithOptions(request *DeleteIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamResourceDiscoveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom resource discovery instance.
//
// Description:
//
//	If a resource discovery instance is shared, it cannot be deleted.
//
// @param request - DeleteIpamResourceDiscoveryRequest
//
// @return DeleteIpamResourceDiscoveryResponse
func (client *Client) DeleteIpamResourceDiscovery(request *DeleteIpamResourceDiscoveryRequest) (_result *DeleteIpamResourceDiscoveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpamResourceDiscoveryResponse{}
	_body, _err := client.DeleteIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM) scope.
//
// Description:
//
// ### [](#)Usage notes
//
//   - You cannot delete the private scope and public scope created by the system.
//
//   - Before you delete an IPAM scope, make sure that all pools within the scope are deleted. You can call **DeleteIpamPool*	- to delete IPAM pools.
//
// @param request - DeleteIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpamScopeResponse
func (client *Client) DeleteIpamScopeWithOptions(request *DeleteIpamScopeRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpamScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IP Address Manager (IPAM) scope.
//
// Description:
//
// ### [](#)Usage notes
//
//   - You cannot delete the private scope and public scope created by the system.
//
//   - Before you delete an IPAM scope, make sure that all pools within the scope are deleted. You can call **DeleteIpamPool*	- to delete IPAM pools.
//
// @param request - DeleteIpamScopeRequest
//
// @return DeleteIpamScopeResponse
func (client *Client) DeleteIpamScope(request *DeleteIpamScopeRequest) (_result *DeleteIpamScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpamScopeResponse{}
	_body, _err := client.DeleteIpamScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates resource discovery and IPAM instances.
//
// @param request - DissociateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateIpamResourceDiscoveryResponse
func (client *Client) DissociateIpamResourceDiscoveryWithOptions(request *DissociateIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *DissociateIpamResourceDiscoveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates resource discovery and IPAM instances.
//
// @param request - DissociateIpamResourceDiscoveryRequest
//
// @return DissociateIpamResourceDiscoveryResponse
func (client *Client) DissociateIpamResourceDiscovery(request *DissociateIpamResourceDiscoveryRequest) (_result *DissociateIpamResourceDiscoveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DissociateIpamResourceDiscoveryResponse{}
	_body, _err := client.DissociateIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - GetIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpamPoolAllocationResponse
func (client *Client) GetIpamPoolAllocationWithOptions(request *GetIpamPoolAllocationRequest, runtime *dara.RuntimeOptions) (_result *GetIpamPoolAllocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - GetIpamPoolAllocationRequest
//
// @return GetIpamPoolAllocationResponse
func (client *Client) GetIpamPoolAllocation(request *GetIpamPoolAllocationRequest) (_result *GetIpamPoolAllocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIpamPoolAllocationResponse{}
	_body, _err := client.GetIpamPoolAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the available CIDR blocks of the IPAM pool.
//
// @param request - GetIpamPoolNextAvailableCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpamPoolNextAvailableCidrResponse
func (client *Client) GetIpamPoolNextAvailableCidrWithOptions(request *GetIpamPoolNextAvailableCidrRequest, runtime *dara.RuntimeOptions) (_result *GetIpamPoolNextAvailableCidrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the available CIDR blocks of the IPAM pool.
//
// @param request - GetIpamPoolNextAvailableCidrRequest
//
// @return GetIpamPoolNextAvailableCidrResponse
func (client *Client) GetIpamPoolNextAvailableCidr(request *GetIpamPoolNextAvailableCidrRequest) (_result *GetIpamPoolNextAvailableCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIpamPoolNextAvailableCidrResponse{}
	_body, _err := client.GetIpamPoolNextAvailableCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether IP Address Manager (IPAM) is activated.
//
// @param request - GetVpcIpamServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcIpamServiceStatusResponse
func (client *Client) GetVpcIpamServiceStatusWithOptions(request *GetVpcIpamServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetVpcIpamServiceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether IP Address Manager (IPAM) is activated.
//
// @param request - GetVpcIpamServiceStatusRequest
//
// @return GetVpcIpamServiceStatusResponse
func (client *Client) GetVpcIpamServiceStatus(request *GetVpcIpamServiceStatusRequest) (_result *GetVpcIpamServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVpcIpamServiceStatusResponse{}
	_body, _err := client.GetVpcIpamServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries discovered resources.
//
// @param request - ListIpamDiscoveredResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamDiscoveredResourceResponse
func (client *Client) ListIpamDiscoveredResourceWithOptions(request *ListIpamDiscoveredResourceRequest, runtime *dara.RuntimeOptions) (_result *ListIpamDiscoveredResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries discovered resources.
//
// @param request - ListIpamDiscoveredResourceRequest
//
// @return ListIpamDiscoveredResourceResponse
func (client *Client) ListIpamDiscoveredResource(request *ListIpamDiscoveredResourceRequest) (_result *ListIpamDiscoveredResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamDiscoveredResourceResponse{}
	_body, _err := client.ListIpamDiscoveredResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamPoolAllocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolAllocationsResponse
func (client *Client) ListIpamPoolAllocationsWithOptions(request *ListIpamPoolAllocationsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamPoolAllocationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamPoolAllocationsRequest
//
// @return ListIpamPoolAllocationsResponse
func (client *Client) ListIpamPoolAllocations(request *ListIpamPoolAllocationsRequest) (_result *ListIpamPoolAllocationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamPoolAllocationsResponse{}
	_body, _err := client.ListIpamPoolAllocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries CIDR blocks provisioned to an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamPoolCidrsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolCidrsResponse
func (client *Client) ListIpamPoolCidrsWithOptions(request *ListIpamPoolCidrsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamPoolCidrsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CIDR blocks provisioned to an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamPoolCidrsRequest
//
// @return ListIpamPoolCidrsResponse
func (client *Client) ListIpamPoolCidrs(request *ListIpamPoolCidrsRequest) (_result *ListIpamPoolCidrsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamPoolCidrsResponse{}
	_body, _err := client.ListIpamPoolCidrsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP Address Manager (IPAM) pools.
//
// @param request - ListIpamPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamPoolsResponse
func (client *Client) ListIpamPoolsWithOptions(request *ListIpamPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamPoolsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP Address Manager (IPAM) pools.
//
// @param request - ListIpamPoolsRequest
//
// @return ListIpamPoolsResponse
func (client *Client) ListIpamPools(request *ListIpamPoolsRequest) (_result *ListIpamPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamPoolsResponse{}
	_body, _err := client.ListIpamPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resources in an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamResourceCidrsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceCidrsResponse
func (client *Client) ListIpamResourceCidrsWithOptions(request *ListIpamResourceCidrsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamResourceCidrsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resources in an IP Address Manager (IPAM) pool.
//
// @param request - ListIpamResourceCidrsRequest
//
// @return ListIpamResourceCidrsResponse
func (client *Client) ListIpamResourceCidrs(request *ListIpamResourceCidrsRequest) (_result *ListIpamResourceCidrsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamResourceCidrsResponse{}
	_body, _err := client.ListIpamResourceCidrsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IPAM resource discovery instances.
//
// @param request - ListIpamResourceDiscoveriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceDiscoveriesResponse
func (client *Client) ListIpamResourceDiscoveriesWithOptions(request *ListIpamResourceDiscoveriesRequest, runtime *dara.RuntimeOptions) (_result *ListIpamResourceDiscoveriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IPAM resource discovery instances.
//
// @param request - ListIpamResourceDiscoveriesRequest
//
// @return ListIpamResourceDiscoveriesResponse
func (client *Client) ListIpamResourceDiscoveries(request *ListIpamResourceDiscoveriesRequest) (_result *ListIpamResourceDiscoveriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamResourceDiscoveriesResponse{}
	_body, _err := client.ListIpamResourceDiscoveriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the association between resource discovery and IPAM.
//
// @param request - ListIpamResourceDiscoveryAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamResourceDiscoveryAssociationsResponse
func (client *Client) ListIpamResourceDiscoveryAssociationsWithOptions(request *ListIpamResourceDiscoveryAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamResourceDiscoveryAssociationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association between resource discovery and IPAM.
//
// @param request - ListIpamResourceDiscoveryAssociationsRequest
//
// @return ListIpamResourceDiscoveryAssociationsResponse
func (client *Client) ListIpamResourceDiscoveryAssociations(request *ListIpamResourceDiscoveryAssociationsRequest) (_result *ListIpamResourceDiscoveryAssociationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamResourceDiscoveryAssociationsResponse{}
	_body, _err := client.ListIpamResourceDiscoveryAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP Address Manager (IPAM) scopes.
//
// @param request - ListIpamScopesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamScopesResponse
func (client *Client) ListIpamScopesWithOptions(request *ListIpamScopesRequest, runtime *dara.RuntimeOptions) (_result *ListIpamScopesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP Address Manager (IPAM) scopes.
//
// @param request - ListIpamScopesRequest
//
// @return ListIpamScopesResponse
func (client *Client) ListIpamScopes(request *ListIpamScopesRequest) (_result *ListIpamScopesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamScopesResponse{}
	_body, _err := client.ListIpamScopesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP Address Managers (IPAMs).
//
// @param request - ListIpamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIpamsResponse
func (client *Client) ListIpamsWithOptions(request *ListIpamsRequest, runtime *dara.RuntimeOptions) (_result *ListIpamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP Address Managers (IPAMs).
//
// @param request - ListIpamsRequest
//
// @return ListIpamsResponse
func (client *Client) ListIpams(request *ListIpamsRequest) (_result *ListIpamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIpamsResponse{}
	_body, _err := client.ListIpamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of resource tags.
//
// Description:
//
// ### [](#)Usage notes
//
//   - You must specify **ResourceId.N*	- or **Tag.N*	- that consists of **Tag.N.Key*	- and **Tag.N.Value*	- in the request to specify the object that you want to query.
//
//   - **Tag.N*	- is a resource tag that consists of a key-value pair. If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
//   - If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
//   - If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of resource tags.
//
// Description:
//
// ### [](#)Usage notes
//
//   - You must specify **ResourceId.N*	- or **Tag.N*	- that consists of **Tag.N.Key*	- and **Tag.N.Value*	- in the request to specify the object that you want to query.
//
//   - **Tag.N*	- is a resource tag that consists of a key-value pair. If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
//   - If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
//   - If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
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
// Activates IP Address Manager (IPAM).
//
// @param request - OpenVpcIpamServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenVpcIpamServiceResponse
func (client *Client) OpenVpcIpamServiceWithOptions(request *OpenVpcIpamServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenVpcIpamServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates IP Address Manager (IPAM).
//
// @param request - OpenVpcIpamServiceRequest
//
// @return OpenVpcIpamServiceResponse
func (client *Client) OpenVpcIpamService(request *OpenVpcIpamServiceRequest) (_result *OpenVpcIpamServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenVpcIpamServiceResponse{}
	_body, _err := client.OpenVpcIpamServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Removes a tag from a resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Updates an IP Address Manager (IPAM).
//
// @param request - UpdateIpamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamResponse
func (client *Client) UpdateIpamWithOptions(request *UpdateIpamRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an IP Address Manager (IPAM).
//
// @param request - UpdateIpamRequest
//
// @return UpdateIpamResponse
func (client *Client) UpdateIpam(request *UpdateIpamRequest) (_result *UpdateIpamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIpamResponse{}
	_body, _err := client.UpdateIpamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information about an IP Address Manager (IPAM) pool.
//
// @param request - UpdateIpamPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamPoolResponse
func (client *Client) UpdateIpamPoolWithOptions(request *UpdateIpamPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamPoolResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about an IP Address Manager (IPAM) pool.
//
// @param request - UpdateIpamPoolRequest
//
// @return UpdateIpamPoolResponse
func (client *Client) UpdateIpamPool(request *UpdateIpamPoolRequest) (_result *UpdateIpamPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIpamPoolResponse{}
	_body, _err := client.UpdateIpamPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - UpdateIpamPoolAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamPoolAllocationResponse
func (client *Client) UpdateIpamPoolAllocationWithOptions(request *UpdateIpamPoolAllocationRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamPoolAllocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies CIDR block allocations of an IP Address Manager (IPAM) pool.
//
// @param request - UpdateIpamPoolAllocationRequest
//
// @return UpdateIpamPoolAllocationResponse
func (client *Client) UpdateIpamPoolAllocation(request *UpdateIpamPoolAllocationRequest) (_result *UpdateIpamPoolAllocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIpamPoolAllocationResponse{}
	_body, _err := client.UpdateIpamPoolAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a resource discovery instance.
//
// Description:
//
//	  You can add or remove effective regions only for custom resource discovery instances.
//
//		- When removing effective regions from a resource discovery instance, the hosted region cannot be included.
//
// @param request - UpdateIpamResourceDiscoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamResourceDiscoveryResponse
func (client *Client) UpdateIpamResourceDiscoveryWithOptions(request *UpdateIpamResourceDiscoveryRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamResourceDiscoveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  You can add or remove effective regions only for custom resource discovery instances.
//
//		- When removing effective regions from a resource discovery instance, the hosted region cannot be included.
//
// @param request - UpdateIpamResourceDiscoveryRequest
//
// @return UpdateIpamResourceDiscoveryResponse
func (client *Client) UpdateIpamResourceDiscovery(request *UpdateIpamResourceDiscoveryRequest) (_result *UpdateIpamResourceDiscoveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIpamResourceDiscoveryResponse{}
	_body, _err := client.UpdateIpamResourceDiscoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information about an IP Address Manager (IPAM) scope.
//
// @param request - UpdateIpamScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpamScopeResponse
func (client *Client) UpdateIpamScopeWithOptions(request *UpdateIpamScopeRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpamScopeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information about an IP Address Manager (IPAM) scope.
//
// @param request - UpdateIpamScopeRequest
//
// @return UpdateIpamScopeResponse
func (client *Client) UpdateIpamScope(request *UpdateIpamScopeRequest) (_result *UpdateIpamScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIpamScopeResponse{}
	_body, _err := client.UpdateIpamScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
