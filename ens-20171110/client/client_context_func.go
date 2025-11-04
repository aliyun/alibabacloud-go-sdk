// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Associates a network access control list (ACL) with a network.
//
// @param request - AccosicateNetworkAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccosicateNetworkAclResponse
func (client *Client) AccosicateNetworkAclWithContext(ctx context.Context, request *AccosicateNetworkAclRequest, runtime *dara.RuntimeOptions) (_result *AccosicateNetworkAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkAclId) {
		query["NetworkAclId"] = request.NetworkAclId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AccosicateNetworkAcl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AccosicateNetworkAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds backend servers.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param tmpReq - AddBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBackendServersResponse
func (client *Client) AddBackendServersWithContext(ctx context.Context, tmpReq *AddBackendServersRequest, runtime *dara.RuntimeOptions) (_result *AddBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddBackendServersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BackendServers) {
		request.BackendServersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackendServers, dara.String("BackendServers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServersShrink) {
		query["BackendServers"] = request.BackendServersShrink
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBackendServers"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBackendServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an IPv6 network interface controller (NIC). A public IP address is automatically assigned at the same time.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 5 times per second per user.
//
//		- Internal networks and IPv4 addresses are not supported.
//
// @param request - AddNetworkInterfaceToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddNetworkInterfaceToInstanceResponse
func (client *Client) AddNetworkInterfaceToInstanceWithContext(ctx context.Context, request *AddNetworkInterfaceToInstanceRequest, runtime *dara.RuntimeOptions) (_result *AddNetworkInterfaceToInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Networks) {
		query["Networks"] = request.Networks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddNetworkInterfaceToInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddNetworkInterfaceToInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an elastic IP address (EIP) to a Source Network Address Translation (SNAT) entry.
//
// @param request - AddSnatIpForSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSnatIpForSnatEntryResponse
func (client *Client) AddSnatIpForSnatEntryWithContext(ctx context.Context, request *AddSnatIpForSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *AddSnatIpForSnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	if !dara.IsNil(request.SnatIp) {
		query["SnatIp"] = request.SnatIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSnatIpForSnatEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSnatIpForSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns secondary private IP addresses to an elastic network interface (ENI).
//
// @param request - AssignPrivateIpAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignPrivateIpAddressesResponse
func (client *Client) AssignPrivateIpAddressesWithContext(ctx context.Context, request *AssignPrivateIpAddressesRequest, runtime *dara.RuntimeOptions) (_result *AssignPrivateIpAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignPrivateIpAddresses"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignPrivateIpAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an elastic IP address (EIP) with a cloud resource that is deployed in the same region.
//
// @param request - AssociateEnsEipAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateEnsEipAddressResponse
func (client *Client) AssociateEnsEipAddressWithContext(ctx context.Context, request *AssociateEnsEipAddressRequest, runtime *dara.RuntimeOptions) (_result *AssociateEnsEipAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationId) {
		query["AllocationId"] = request.AllocationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Standby) {
		query["Standby"] = request.Standby
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateEnsEipAddress"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateEnsEipAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a high-availability virtual IP address (HAVIP) with an Edge Node Service (ENS) instance or elastic network interface (ENI).
//
// Description:
//
// When you call this operation to associate an HAVIP, take note of the following items:
//
//   - An HAVIP immediately takes effect after it is associated. You do not need to restart the ENS instance. However, you need to associate the HAVIP with the ENI of the ENS instance.
//
//   - The HAVIP and ENS instance must belong to the same vSwitch.
//
//   - The ENS instance must be in the Running or Stopped state.
//
//   - The HAVIP must be in the Available or InUse state.
//
//   - AssociateHaVip is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the DescribeHaVips operation to query the status of an HAVIP:
//
//   - If the HAVIP is in the Associating state, the HAVIP is being associated.
//
//     <!---->
//
//   - If the HAVIP is in the InUse state, the HAVIP is associated.
//
// @param request - AssociateHaVipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateHaVipResponse
func (client *Client) AssociateHaVipWithContext(ctx context.Context, request *AssociateHaVipRequest, runtime *dara.RuntimeOptions) (_result *AssociateHaVipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HaVipId) {
		query["HaVipId"] = request.HaVipId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateHaVip"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateHaVipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a disk to an Edge Node Service (ENS) instance.
//
// @param request - AttachDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDiskResponse
func (client *Client) AttachDiskWithContext(ctx context.Context, request *AttachDiskRequest, runtime *dara.RuntimeOptions) (_result *AttachDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteWithInstance) {
		query["DeleteWithInstance"] = request.DeleteWithInstance
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDisk"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an Edge Node Service (ENS) instance to Container Service for Kubernetes (ACK).
//
// Description:
//
// # [](#)Usage notes
//
//   - You can call this operation up to 10 times per second per account.
//
//   - After you execute the command, the instance restarts loading.
//
//   - Limits: The instance has at least two vCPUs and 4 GB memory. An image of CentOS 7.4 or later is required.
//
// @param request - AttachEnsInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachEnsInstancesResponse
func (client *Client) AttachEnsInstancesWithContext(ctx context.Context, request *AttachEnsInstancesRequest, runtime *dara.RuntimeOptions) (_result *AttachEnsInstancesResponse, _err error) {
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

	if !dara.IsNil(request.Scripts) {
		query["Scripts"] = request.Scripts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachEnsInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachEnsInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a shared data group (SDG) to the corresponding Android in Container (AIC) instance.
//
// @param tmpReq - AttachInstanceSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachInstanceSDGResponse
func (client *Client) AttachInstanceSDGWithContext(ctx context.Context, tmpReq *AttachInstanceSDGRequest, runtime *dara.RuntimeOptions) (_result *AttachInstanceSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachInstanceSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LoadOpt) {
		request.LoadOptShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LoadOpt, dara.String("LoadOpt"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskAccessProtocol) {
		query["DiskAccessProtocol"] = request.DiskAccessProtocol
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.LoadOptShrink) {
		query["LoadOpt"] = request.LoadOptShrink
	}

	if !dara.IsNil(request.SDGId) {
		query["SDGId"] = request.SDGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachInstanceSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachInstanceSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches an Elastic Network Interface (ENI) to an Edge Node Service (ECS) instance.
//
// Description:
//
// When you call this operation, take note of the following limits:
//
//   - The ENI must be in the Available state.
//
//   - An ENI can be attached to only one instance that is the same zone and the same Virtual Private Cloud (VPC).
//
//   - The instance must be in the Stopped state.
//
//   - A maximum of 10 ENIs can be attached to an instance.
//
//   - This operation is an asynchronous operation. After you call this operation to attach an ENI, you can view the status of the ENI to check whether the ENI is attached.
//
// @param request - AttachNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachNetworkInterfaceResponse
func (client *Client) AttachNetworkInterfaceWithContext(ctx context.Context, request *AttachNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *AttachNetworkInterfaceResponse, _err error) {
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

	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachNetworkInterface"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an inbound security group rule. This operation allows or denies the inbound traffic from other devices to instances in the security group.
//
// @param request - AuthorizeSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeSecurityGroupResponse
func (client *Client) AuthorizeSecurityGroupWithContext(ctx context.Context, request *AuthorizeSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeSecurityGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeSecurityGroup"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an outbound security group rule. This operation allows or denies the outbound traffic from the instances in the security group to other devices.
//
// Description:
//
// In the security group-related API documents, outbound traffic refers to the traffic that is sent by the source device and received at the destination device.
//
// @param request - AuthorizeSecurityGroupEgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeSecurityGroupEgressResponse
func (client *Client) AuthorizeSecurityGroupEgressWithContext(ctx context.Context, request *AuthorizeSecurityGroupEgressRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeSecurityGroupEgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestCidrIp) {
		query["DestCidrIp"] = request.DestCidrIp
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeSecurityGroupEgress"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeSecurityGroupEgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates multiple instances in a specified event at a time. You can execute the task immediately or schedule the task execution.
//
// Description:
//
// ## [](#)Request description
//
//   - This O\\&M operation is supported only by the Instance:SystemUpgrade.Migrate event.
//
// @param tmpReq - BatchEventMigrateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchEventMigrateInstanceResponse
func (client *Client) BatchEventMigrateInstanceWithContext(ctx context.Context, tmpReq *BatchEventMigrateInstanceRequest, runtime *dara.RuntimeOptions) (_result *BatchEventMigrateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchEventMigrateInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventInfos) {
		request.EventInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventInfos, dara.String("EventInfos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EventInfosShrink) {
		query["EventInfos"] = request.EventInfosShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchEventMigrateInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchEventMigrateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # The event that is used to immediately redeploy specified resources in batches or by appointment
//
// Description:
//
//	This O\\&M operation supports only the following event types: Instance:SystemMaintenance.Reboot (instance reboot due to system problems)
//
// @param tmpReq - BatchEventRebootInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchEventRebootInstanceResponse
func (client *Client) BatchEventRebootInstanceWithContext(ctx context.Context, tmpReq *BatchEventRebootInstanceRequest, runtime *dara.RuntimeOptions) (_result *BatchEventRebootInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchEventRebootInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventInfos) {
		request.EventInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventInfos, dara.String("EventInfos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EventInfosShrink) {
		query["EventInfos"] = request.EventInfosShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchEventRebootInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchEventRebootInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch redeployment
//
// Description:
//
// - This operation currently only supports event types: Instance:SystemFailure.Redeploy (redeploy instance due to system issues), Instance:SystemMaintenance.Redeploy (redeploy instance due to system maintenance)
//
// @param tmpReq - BatchEventRedeployInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchEventRedeployInstanceResponse
func (client *Client) BatchEventRedeployInstanceWithContext(ctx context.Context, tmpReq *BatchEventRedeployInstanceRequest, runtime *dara.RuntimeOptions) (_result *BatchEventRedeployInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchEventRedeployInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventInfos) {
		request.EventInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventInfos, dara.String("EventInfos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EventInfosShrink) {
		query["EventInfos"] = request.EventInfosShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchEventRedeployInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchEventRedeployInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清理分发数据
//
// @param request - CleanDistDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CleanDistDataResponse
func (client *Client) CleanDistDataWithContext(ctx context.Context, request *CleanDistDataRequest, runtime *dara.RuntimeOptions) (_result *CleanDistDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataName) {
		query["DataName"] = request.DataName
	}

	if !dara.IsNil(request.DataVersion) {
		query["DataVersion"] = request.DataVersion
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CleanDistData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CleanDistDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a shared data group (SDG) across nodes.
//
// @param tmpReq - CopySDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopySDGResponse
func (client *Client) CopySDGWithContext(ctx context.Context, tmpReq *CopySDGRequest, runtime *dara.RuntimeOptions) (_result *CopySDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CopySDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestinationRegionIds) {
		request.DestinationRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationRegionIds, dara.String("DestinationRegionIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopySDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopySDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a snapshot across nodes.
//
// @param tmpReq - CopySnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopySnapshotResponse
func (client *Client) CopySnapshotWithContext(ctx context.Context, tmpReq *CopySnapshotRequest, runtime *dara.RuntimeOptions) (_result *CopySnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CopySnapshotShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestinationRegionIds) {
		request.DestinationRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationRegionIds, dara.String("DestinationRegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationRegionIdsShrink) {
		query["DestinationRegionIds"] = request.DestinationRegionIdsShrink
	}

	if !dara.IsNil(request.DestinationSnapshotDescription) {
		query["DestinationSnapshotDescription"] = request.DestinationSnapshotDescription
	}

	if !dara.IsNil(request.DestinationSnapshotName) {
		query["DestinationSnapshotName"] = request.DestinationSnapshotName
	}

	if !dara.IsNil(request.InstanceBillingCycle) {
		query["InstanceBillingCycle"] = request.InstanceBillingCycle
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopySnapshot"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopySnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ARM server.
//
// @param request - CreateARMServerInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateARMServerInstancesResponse
func (client *Client) CreateARMServerInstancesWithContext(ctx context.Context, request *CreateARMServerInstancesRequest, runtime *dara.RuntimeOptions) (_result *CreateARMServerInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.Cidr) {
		query["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnvironmentVar) {
		query["EnvironmentVar"] = request.EnvironmentVar
	}

	if !dara.IsNil(request.Frequency) {
		query["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceBillingCycle) {
		query["InstanceBillingCycle"] = request.InstanceBillingCycle
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.NameSpace) {
		query["NameSpace"] = request.NameSpace
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.Resolution) {
		query["Resolution"] = request.Resolution
	}

	if !dara.IsNil(request.ServerName) {
		query["ServerName"] = request.ServerName
	}

	if !dara.IsNil(request.ServerType) {
		query["ServerType"] = request.ServerType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateARMServerInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateARMServerInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an edge application that allows you to manage Edge Node Service (ENS) nodes in containers, bare metal instances, and virtual machines.
//
// @param request - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithContext(ctx context.Context, request *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplication"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Creates a classic network
//
// @param request - CreateClassicNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClassicNetworkResponse
func (client *Client) CreateClassicNetworkWithContext(ctx context.Context, request *CreateClassicNetworkRequest, runtime *dara.RuntimeOptions) (_result *CreateClassicNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.NetworkName) {
		query["NetworkName"] = request.NetworkName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClassicNetwork"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClassicNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Container Service for Kubernetes (ACK) edge cluster.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- Creating a cluster is an asynchronous operation. After this operation returns the response, it takes 10 to 20 minutes to initialize the cluster. You can call the DescribeCluster operation to query the cluster status. After you create a cluster, you can call the DescribeClusterKubeConfig operation to obtain the cluster certificate.
//
// @param request - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithContext(ctx context.Context, request *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterVersion) {
		query["ClusterVersion"] = request.ClusterVersion
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go or subscription data disk.
//
// @param request - CreateDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiskResponse
func (client *Client) CreateDiskWithContext(ctx context.Context, request *CreateDiskRequest, runtime *dara.RuntimeOptions) (_result *CreateDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.DiskName) {
		query["DiskName"] = request.DiskName
	}

	if !dara.IsNil(request.Encrypted) {
		query["Encrypted"] = request.Encrypted
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.InstanceBillingCycle) {
		query["InstanceBillingCycle"] = request.InstanceBillingCycle
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.KMSKeyId) {
		query["KMSKeyId"] = request.KMSKeyId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDisk"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for an elastic IP address (EIP).
//
// Description:
//
//	  You can call this operation up to 5,000 times per second per account.
//
//		- You can call this operation up to 50 times per second per user.
//
// @param request - CreateEipInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEipInstanceResponse
func (client *Client) CreateEipInstanceWithContext(ctx context.Context, request *CreateEipInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateEipInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEipInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEipInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom route entry.
//
// @param request - CreateEnsRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnsRouteEntryResponse
func (client *Client) CreateEnsRouteEntryWithContext(ctx context.Context, request *CreateEnsRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateEnsRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.NextHopId) {
		query["NextHopId"] = request.NextHopId
	}

	if !dara.IsNil(request.NextHopType) {
		query["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.RouteEntryName) {
		query["RouteEntryName"] = request.RouteEntryName
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	if !dara.IsNil(request.SourceCidrBlock) {
		query["SourceCidrBlock"] = request.SourceCidrBlock
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnsRouteEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnsRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建售卖约束
//
// @param tmpReq - CreateEnsSaleControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnsSaleControlResponse
func (client *Client) CreateEnsSaleControlWithContext(ctx context.Context, tmpReq *CreateEnsSaleControlRequest, runtime *dara.RuntimeOptions) (_result *CreateEnsSaleControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEnsSaleControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SaleControls) {
		request.SaleControlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SaleControls, dara.String("SaleControls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUidAccount) {
		query["AliUidAccount"] = request.AliUidAccount
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CustomAccount) {
		query["CustomAccount"] = request.CustomAccount
	}

	if !dara.IsNil(request.SaleControlsShrink) {
		query["SaleControls"] = request.SaleControlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnsSaleControl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnsSaleControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an edge service.
//
// @param request - CreateEnsServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnsServiceResponse
func (client *Client) CreateEnsServiceWithContext(ctx context.Context, request *CreateEnsServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateEnsServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsServiceId) {
		query["EnsServiceId"] = request.EnsServiceId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnsService"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnsServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an edge private network (EPN) instance.
//
// @param request - CreateEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEpnInstanceResponse
func (client *Client) CreateEpnInstanceWithContext(ctx context.Context, request *CreateEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceName) {
		query["EPNInstanceName"] = request.EPNInstanceName
	}

	if !dara.IsNil(request.EPNInstanceType) {
		query["EPNInstanceType"] = request.EPNInstanceType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.InternetMaxBandwidthOut) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !dara.IsNil(request.NetworkingModel) {
		query["NetworkingModel"] = request.NetworkingModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a NAS file system.
//
// @param tmpReq - CreateFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileSystemResponse
func (client *Client) CreateFileSystemWithContext(ctx context.Context, tmpReq *CreateFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFileSystemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderDetails) {
		request.OrderDetailsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderDetails, dara.String("OrderDetails"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFileSystem"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a Destination Network Address Translation (DNAT) entry to a DNAT table.
//
// @param request - CreateForwardEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateForwardEntryResponse
func (client *Client) CreateForwardEntryWithContext(ctx context.Context, request *CreateForwardEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateForwardEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalIp) {
		query["ExternalIp"] = request.ExternalIp
	}

	if !dara.IsNil(request.ExternalPort) {
		query["ExternalPort"] = request.ExternalPort
	}

	if !dara.IsNil(request.ForwardEntryName) {
		query["ForwardEntryName"] = request.ForwardEntryName
	}

	if !dara.IsNil(request.HealthCheckPort) {
		query["HealthCheckPort"] = request.HealthCheckPort
	}

	if !dara.IsNil(request.InternalIp) {
		query["InternalIp"] = request.InternalIp
	}

	if !dara.IsNil(request.InternalPort) {
		query["InternalPort"] = request.InternalPort
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.StandbyExternalIp) {
		query["StandbyExternalIp"] = request.StandbyExternalIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateForwardEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateForwardEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a high-availability virtual IP address (HAVIP).
//
// @param request - CreateHaVipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHaVipResponse
func (client *Client) CreateHaVipWithContext(ctx context.Context, request *CreateHaVipRequest, runtime *dara.RuntimeOptions) (_result *CreateHaVipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IpAddress) {
		query["IpAddress"] = request.IpAddress
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHaVip"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHaVipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image from an instance.
//
// @param request - CreateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageResponse
func (client *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest, runtime *dara.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteAfterImageUpload) {
		query["DeleteAfterImageUpload"] = request.DeleteAfterImageUpload
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.TargetOSSRegionId) {
		query["TargetOSSRegionId"] = request.TargetOSSRegionId
	}

	if !dara.IsNil(request.WithDataDisks) {
		query["WithDataDisks"] = request.WithDataDisks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImage"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- We recommend that you increase the request time because instance creation is an asynchronous operation. If the return code of the API operation is 0, it indicates that the request is successful, but does not indicate that the instance is created. If the request is successful, an instance ID is returned. You can check whether the instance is created based on the instance ID.
//
//		- InvalidUserData.NotInWhiteList operation restriction: You can create an instance only if you are in the whitelist in which members have the purchase permissions. Otherwise, an error is returned.
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.IpType) {
		query["IpType"] = request.IpType
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordInherit) {
		query["PasswordInherit"] = request.PasswordInherit
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.PublicIpIdentification) {
		query["PublicIpIdentification"] = request.PublicIpIdentification
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
	}

	if !dara.IsNil(request.UniqueSuffix) {
		query["UniqueSuffix"] = request.UniqueSuffix
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.DataDisk) {
		query["DataDisk"] = request.DataDisk
	}

	if !dara.IsNil(request.SystemDisk) {
		query["SystemDisk"] = request.SystemDisk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用CreateInstanceOpsTask来针对一个实例或实例运维组发起运维任务
//
// @param tmpReq - CreateInstanceActiveOpsTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceActiveOpsTaskResponse
func (client *Client) CreateInstanceActiveOpsTaskWithContext(ctx context.Context, tmpReq *CreateInstanceActiveOpsTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceActiveOpsTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateInstanceActiveOpsTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceActiveOpsTask"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceActiveOpsTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an SSH key pair.
//
// Description:
//
// An SSH key pair consists of a public key and a private key. ENS stores the public key and returns the unencrypted private key that is PEM-encoded in the PKCS#8 format. You must securely lock away the private key.
//
// @param request - CreateKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyPairResponse
func (client *Client) CreateKeyPairWithContext(ctx context.Context, request *CreateKeyPairRequest, runtime *dara.RuntimeOptions) (_result *CreateKeyPairResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKeyPair"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Edge Load Balancer (ELB) instance.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 5 times per second per user.
//
// @param request - CreateLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerResponse
func (client *Client) CreateLoadBalancerWithContext(ctx context.Context, request *CreateLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.LoadBalancerName) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !dara.IsNil(request.LoadBalancerSpec) {
		query["LoadBalancerSpec"] = request.LoadBalancerSpec
	}

	if !dara.IsNil(request.LoadBalancerType) {
		query["LoadBalancerType"] = request.LoadBalancerType
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancer"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an HTTP listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - CreateLoadBalancerHTTPListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerHTTPListenerResponse
func (client *Client) CreateLoadBalancerHTTPListenerWithContext(ctx context.Context, request *CreateLoadBalancerHTTPListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerHTTPListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ForwardPort) {
		query["ForwardPort"] = request.ForwardPort
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerForward) {
		query["ListenerForward"] = request.ListenerForward
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.XForwardedFor) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerHTTPListener"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerHTTPListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an HTTPS listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - CreateLoadBalancerHTTPSListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerHTTPSListenerResponse
func (client *Client) CreateLoadBalancerHTTPSListenerWithContext(ctx context.Context, request *CreateLoadBalancerHTTPSListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerHTTPSListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Cookie) {
		query["Cookie"] = request.Cookie
	}

	if !dara.IsNil(request.CookieTimeout) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ForwardPort) {
		query["ForwardPort"] = request.ForwardPort
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerForward) {
		query["ListenerForward"] = request.ListenerForward
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !dara.IsNil(request.StickySessionType) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerHTTPSListener"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerHTTPSListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Transmission Control Protocol (TCP) listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - CreateLoadBalancerTCPListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerTCPListenerResponse
func (client *Client) CreateLoadBalancerTCPListenerWithContext(ctx context.Context, request *CreateLoadBalancerTCPListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerTCPListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EipTransmit) {
		query["EipTransmit"] = request.EipTransmit
	}

	if !dara.IsNil(request.EstablishedTimeout) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckConnectTimeout) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckType) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.PersistenceTimeout) {
		query["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerTCPListener"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerTCPListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a User Datagram Protocol (UDP) listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - CreateLoadBalancerUDPListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerUDPListenerResponse
func (client *Client) CreateLoadBalancerUDPListenerWithContext(ctx context.Context, request *CreateLoadBalancerUDPListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerUDPListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EipTransmit) {
		query["EipTransmit"] = request.EipTransmit
	}

	if !dara.IsNil(request.EstablishedTimeout) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckConnectTimeout) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !dara.IsNil(request.HealthCheckExp) {
		query["HealthCheckExp"] = request.HealthCheckExp
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckReq) {
		query["HealthCheckReq"] = request.HealthCheckReq
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerUDPListener"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerUDPListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mount target.
//
// Description:
//
// ## [](#)Precautions
//
// After you call this operation, a mount target is not immediately created. Therefore, we recommend that you call the DescribeMountTargets operation to query the status of the mount target. If the mount target is in the Active state, you can then mount the file system. Otherwise, the file system may fail to be mounted.
//
// @param request - CreateMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMountTargetResponse
func (client *Client) CreateMountTargetWithContext(ctx context.Context, request *CreateMountTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetName) {
		query["MountTargetName"] = request.MountTargetName
	}

	if !dara.IsNil(request.NetWorkId) {
		query["NetWorkId"] = request.NetWorkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMountTarget"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network address translation (NAT) gateway.
//
// @param request - CreateNatGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNatGatewayResponse
func (client *Client) CreateNatGatewayWithContext(ctx context.Context, request *CreateNatGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateNatGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.InstanceBillingCycle) {
		query["InstanceBillingCycle"] = request.InstanceBillingCycle
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNatGateway"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNatGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC).
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 5 times per second per user.
//
// @param request - CreateNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkResponse
func (client *Client) CreateNetworkWithContext(ctx context.Context, request *CreateNetworkRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.NetworkName) {
		query["NetworkName"] = request.NetworkName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetwork"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network access control list (ACL).
//
// @param request - CreateNetworkAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkAclResponse
func (client *Client) CreateNetworkAclWithContext(ctx context.Context, request *CreateNetworkAclRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NetworkAclName) {
		query["NetworkAclName"] = request.NetworkAclName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkAcl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network access control list (ACL) rule.
//
// @param request - CreateNetworkAclEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkAclEntryResponse
func (client *Client) CreateNetworkAclEntryWithContext(ctx context.Context, request *CreateNetworkAclEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkAclEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.NetworkAclEntryName) {
		query["NetworkAclEntryName"] = request.NetworkAclEntryName
	}

	if !dara.IsNil(request.NetworkAclId) {
		query["NetworkAclId"] = request.NetworkAclId
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkAclEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkAclEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an Elastic Network Interface (ENI).
//
// @param tmpReq - CreateNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkInterfaceResponse
func (client *Client) CreateNetworkInterfaceWithContext(ctx context.Context, tmpReq *CreateNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNetworkInterfaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SecurityGroupIds) {
		request.SecurityGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityGroupIds, dara.String("SecurityGroupIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SecurityGroupIdsShrink) {
		query["SecurityGroupIds"] = request.SecurityGroupIdsShrink
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkInterface"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a shared data group (SDG) on an Android in Container (AIC) instance.
//
// Description:
//
// A SDG can be regarded as a data partition image of a virtual device. You can save a data partition of a virtual device as an SDG. A created SDG can be deployed in data partitions of other virtual devices to achieve rapid data distribution and application. The procedure for calling SDG-related API operations:
//
//   - Call the [CreateSDG](~~CreateSDG~~) operation to create an SDG, which is bound to AIC Instance A (InstanceId). After you create the SDG, a blank cloud disk (also known as an original cloud disk) is attached to Device A (InstanceId).
//
//   - Install applications on and deliver files to AIC Instance A (InstanceId).
//
//   - Call the [SaveSDG](~~SaveSDG~~) operation to save the data disk of AIC instance A as SDG A.
//
//   - Call the [DeploySDG](~~DeploySDG~~) operation to deploy SDG A to AIC Instance B. This operattion is executed asynchronously. You can call the [DescribeARMServerInstances](~~DescribeARMServerInstances~~) operation to query the status of AIC Instance B. If the status of AIC Instance B changes to success, AIC insance B is available, and AIC Instances A and B have the same applications running.
//
// @param request - CreateSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSDGResponse
func (client *Client) CreateSDGWithContext(ctx context.Context, request *CreateSDGRequest, runtime *dara.RuntimeOptions) (_result *CreateSDGResponse, _err error) {
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
		Action:      dara.String("CreateSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a security group.
//
// @param tmpReq - CreateSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityGroupResponse
func (client *Client) CreateSecurityGroupWithContext(ctx context.Context, tmpReq *CreateSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Permissions) {
		request.PermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permissions, dara.String("Permissions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.PermissionsShrink) {
		query["Permissions"] = request.PermissionsShrink
	}

	if !dara.IsNil(request.SecurityGroupName) {
		query["SecurityGroupName"] = request.SecurityGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecurityGroup"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a snapshot.
//
// @param request - CreateSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.InstanceBillingCycle) {
		query["InstanceBillingCycle"] = request.InstanceBillingCycle
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSnapshot"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a source network address translation (SNAT) entry to a specified SNAT table.
//
// @param request - CreateSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnatEntryResponse
func (client *Client) CreateSnatEntryWithContext(ctx context.Context, request *CreateSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateSnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EipAffinity) {
		query["EipAffinity"] = request.EipAffinity
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.IspAffinity) {
		query["IspAffinity"] = request.IspAffinity
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.SnatEntryName) {
		query["SnatEntryName"] = request.SnatEntryName
	}

	if !dara.IsNil(request.SnatIp) {
		query["SnatIp"] = request.SnatIp
	}

	if !dara.IsNil(request.SourceCIDR) {
		query["SourceCIDR"] = request.SourceCIDR
	}

	if !dara.IsNil(request.SourceNetworkId) {
		query["SourceNetworkId"] = request.SourceNetworkId
	}

	if !dara.IsNil(request.SourceVSwitchId) {
		query["SourceVSwitchId"] = request.SourceVSwitchId
	}

	if !dara.IsNil(request.StandbySnatIp) {
		query["StandbySnatIp"] = request.StandbySnatIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSnatEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a storage gateway.
//
// @param tmpReq - CreateStorageGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStorageGatewayResponse
func (client *Client) CreateStorageGatewayWithContext(ctx context.Context, tmpReq *CreateStorageGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateStorageGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStorageGatewayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderDetails) {
		request.OrderDetailsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderDetails, dara.String("OrderDetails"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderDetailsShrink) {
		query["OrderDetails"] = request.OrderDetailsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStorageGateway"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStorageGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a storage volume.
//
// @param request - CreateStorageVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStorageVolumeResponse
func (client *Client) CreateStorageVolumeWithContext(ctx context.Context, request *CreateStorageVolumeRequest, runtime *dara.RuntimeOptions) (_result *CreateStorageVolumeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthPassword) {
		query["AuthPassword"] = request.AuthPassword
	}

	if !dara.IsNil(request.AuthProtocol) {
		query["AuthProtocol"] = request.AuthProtocol
	}

	if !dara.IsNil(request.AuthUser) {
		query["AuthUser"] = request.AuthUser
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.IsAuth) {
		query["IsAuth"] = request.IsAuth
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	if !dara.IsNil(request.StorageId) {
		query["StorageId"] = request.StorageId
	}

	if !dara.IsNil(request.VolumeName) {
		query["VolumeName"] = request.VolumeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStorageVolume"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStorageVolumeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A vSwitch is created.
//
// @param request - CreateVSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVSwitchResponse
func (client *Client) CreateVSwitchWithContext(ctx context.Context, request *CreateVSwitchRequest, runtime *dara.RuntimeOptions) (_result *CreateVSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchName) {
		query["VSwitchName"] = request.VSwitchName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVSwitch"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除托管公钥
//
// @param request - DeleteAICPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAICPublicKeyResponse
func (client *Client) DeleteAICPublicKeyWithContext(ctx context.Context, request *DeleteAICPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAICPublicKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAICPublicKey"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAICPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases all containers and resource instances related to a specific application in an asynchronous manner.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithContext(ctx context.Context, request *DeleteApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplication"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a bucket.
//
// Description:
//
//	  Only the Alibaba Cloud Account ID owner of a bucket can delete the bucket from the account.
//
//		- You cannot delete buckets that store objects. You can only delete empty buckets.
//
// @param request - DeleteBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBucketResponse
func (client *Client) DeleteBucketWithContext(ctx context.Context, request *DeleteBucketRequest, runtime *dara.RuntimeOptions) (_result *DeleteBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBucket"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBucketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the lifecycle rules for objects in a bucket.
//
// @param request - DeleteBucketLifecycleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBucketLifecycleResponse
func (client *Client) DeleteBucketLifecycleWithContext(ctx context.Context, request *DeleteBucketLifecycleRequest, runtime *dara.RuntimeOptions) (_result *DeleteBucketLifecycleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBucketLifecycle"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBucketLifecycleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a disk.
//
// Description:
//
// When you release a disk, the disk must be in the Available state.
//
// @param request - DeleteDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDiskResponse
func (client *Client) DeleteDiskWithContext(ctx context.Context, request *DeleteDiskRequest, runtime *dara.RuntimeOptions) (_result *DeleteDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDisk"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a pay-as-you-go elastic IP address (EIP).
//
// @param request - DeleteEipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEipResponse
func (client *Client) DeleteEipWithContext(ctx context.Context, request *DeleteEipRequest, runtime *dara.RuntimeOptions) (_result *DeleteEipResponse, _err error) {
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
		Action:      dara.String("DeleteEip"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom route entry.
//
// @param request - DeleteEnsRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnsRouteEntryResponse
func (client *Client) DeleteEnsRouteEntryWithContext(ctx context.Context, request *DeleteEnsRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnsRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RouteEntryId) {
		query["RouteEntryId"] = request.RouteEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnsRouteEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnsRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除售卖约束的条件约束
//
// @param tmpReq - DeleteEnsSaleConditionControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnsSaleConditionControlResponse
func (client *Client) DeleteEnsSaleConditionControlWithContext(ctx context.Context, tmpReq *DeleteEnsSaleConditionControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnsSaleConditionControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteEnsSaleConditionControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SaleControls) {
		request.SaleControlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SaleControls, dara.String("SaleControls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUidAccount) {
		query["AliUidAccount"] = request.AliUidAccount
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CustomAccount) {
		query["CustomAccount"] = request.CustomAccount
	}

	if !dara.IsNil(request.SaleControlsShrink) {
		query["SaleControls"] = request.SaleControlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnsSaleConditionControl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnsSaleConditionControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除售卖约束基础约束
//
// @param tmpReq - DeleteEnsSaleControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnsSaleControlResponse
func (client *Client) DeleteEnsSaleControlWithContext(ctx context.Context, tmpReq *DeleteEnsSaleControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnsSaleControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteEnsSaleControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SaleControls) {
		request.SaleControlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SaleControls, dara.String("SaleControls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUidAccount) {
		query["AliUidAccount"] = request.AliUidAccount
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CustomAccount) {
		query["CustomAccount"] = request.CustomAccount
	}

	if !dara.IsNil(request.SaleControlsShrink) {
		query["SaleControls"] = request.SaleControlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnsSaleControl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnsSaleControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an edge private network (EPN) instance.
//
// Description:
//
// You can delete an EPN instance only when the instance group information is empty.
//
// @param request - DeleteEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEpnInstanceResponse
func (client *Client) DeleteEpnInstanceWithContext(ctx context.Context, request *DeleteEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a File Storage NAS file system.
//
// @param request - DeleteFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileSystemResponse
func (client *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
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
		Action:      dara.String("DeleteFileSystem"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Destination Network Address Translation (DNAT) entry from a specified DNAT table.
//
// @param request - DeleteForwardEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteForwardEntryResponse
func (client *Client) DeleteForwardEntryWithContext(ctx context.Context, request *DeleteForwardEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteForwardEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForwardEntryId) {
		query["ForwardEntryId"] = request.ForwardEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteForwardEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteForwardEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a high-availability VIP (HAVIP).
//
// @param tmpReq - DeleteHaVipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHaVipsResponse
func (client *Client) DeleteHaVipsWithContext(ctx context.Context, tmpReq *DeleteHaVipsRequest, runtime *dara.RuntimeOptions) (_result *DeleteHaVipsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteHaVipsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HaVipIds) {
		request.HaVipIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HaVipIds, dara.String("HaVipIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.HaVipIdsShrink) {
		query["HaVipIds"] = request.HaVipIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHaVips"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHaVipsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom image.
//
// @param request - DeleteImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageResponse
func (client *Client) DeleteImageWithContext(ctx context.Context, request *DeleteImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImage"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes SSH key pairs.
//
// Description:
//
//	  After you delete an SSH key pair, you can no longer query the key pair by calling the DescribeKeyPairs operation.
//
//		- If you delete an SSH key pair that is bound to an Edge Node Service (ENS) instance, ENS no longer stores the SSH key pair. However, you can still use the key pair to access the instance. When you call the DescribeInstance operation to query instance information, no other information but the name of the key pair (**KeyPairName**) is returned.
//
// @param request - DeleteKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeyPairsResponse
func (client *Client) DeleteKeyPairsWithContext(ctx context.Context, request *DeleteKeyPairsRequest, runtime *dara.RuntimeOptions) (_result *DeleteKeyPairsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairId) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKeyPairs"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DeleteLoadBalancerListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoadBalancerListenerResponse
func (client *Client) DeleteLoadBalancerListenerWithContext(ctx context.Context, request *DeleteLoadBalancerListenerRequest, runtime *dara.RuntimeOptions) (_result *DeleteLoadBalancerListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLoadBalancerListener"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLoadBalancerListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a mount target.
//
// Description:
//
// After you delete a mount target, the mount target cannot be restored. Proceed with caution.
//
// @param request - DeleteMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMountTargetResponse
func (client *Client) DeleteMountTargetWithContext(ctx context.Context, request *DeleteMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetName) {
		query["MountTargetName"] = request.MountTargetName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMountTarget"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Internet network address translation (NAT) gateway.
//
// @param request - DeleteNatGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNatGatewayResponse
func (client *Client) DeleteNatGatewayWithContext(ctx context.Context, request *DeleteNatGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteNatGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceDelete) {
		query["ForceDelete"] = request.ForceDelete
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNatGateway"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNatGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual private cloud (VPC).
//
// @param request - DeleteNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkResponse
func (client *Client) DeleteNetworkWithContext(ctx context.Context, request *DeleteNetworkRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetwork"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a network access control list (ACL).
//
// @param request - DeleteNetworkAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkAclResponse
func (client *Client) DeleteNetworkAclWithContext(ctx context.Context, request *DeleteNetworkAclRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkAclId) {
		query["NetworkAclId"] = request.NetworkAclId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkAcl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a network access control list (ACL) rule.
//
// @param request - DeleteNetworkAclEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkAclEntryResponse
func (client *Client) DeleteNetworkAclEntryWithContext(ctx context.Context, request *DeleteNetworkAclEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkAclEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkAclEntryId) {
		query["NetworkAclEntryId"] = request.NetworkAclEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkAclEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkAclEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an elastic network interface (ENI).
//
// @param tmpReq - DeleteNetworkInterfacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkInterfacesResponse
func (client *Client) DeleteNetworkInterfacesWithContext(ctx context.Context, tmpReq *DeleteNetworkInterfacesRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkInterfacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteNetworkInterfacesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NetworkInterfaceIds) {
		request.NetworkInterfaceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkInterfaceIds, dara.String("NetworkInterfaceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkInterfaceIdsShrink) {
		query["NetworkInterfaceIds"] = request.NetworkInterfaceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkInterfaces"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkInterfacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an object.
//
// @param request - DeleteObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteObjectResponse
func (client *Client) DeleteObjectWithContext(ctx context.Context, request *DeleteObjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteObject"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteObjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a shared data group (SDG).
//
// Description:
//
// If all the SDGs corresponding to the original disk are deleted, the original disk is automatically cleared.
//
// @param tmpReq - DeleteSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSDGResponse
func (client *Client) DeleteSDGWithContext(ctx context.Context, tmpReq *DeleteSDGRequest, runtime *dara.RuntimeOptions) (_result *DeleteSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SDGId) {
		request.SDGIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SDGId, dara.String("SDGId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SDGIdShrink) {
		query["SDGId"] = request.SDGIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a security group.
//
// Description:
//
// Before you delete a security group, make sure that no instances exist in the security group.
//
// @param request - DeleteSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityGroupResponse
func (client *Client) DeleteSecurityGroupWithContext(ctx context.Context, request *DeleteSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityGroup"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a snapshot.
//
// @param request - DeleteSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshot"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a source network address translation (SNAT) entry from a specified SNAT table.
//
// @param request - DeleteSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnatEntryResponse
func (client *Client) DeleteSnatEntryWithContext(ctx context.Context, request *DeleteSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnatEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an elastic IP address (EIP) from a source network address translation (SNAT) entry.
//
// @param request - DeleteSnatIpForSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnatIpForSnatEntryResponse
func (client *Client) DeleteSnatIpForSnatEntryWithContext(ctx context.Context, request *DeleteSnatIpForSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnatIpForSnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	if !dara.IsNil(request.SnatIp) {
		query["SnatIp"] = request.SnatIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnatIpForSnatEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnatIpForSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a storage gateway.
//
// @param request - DeleteStorageGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStorageGatewayResponse
func (client *Client) DeleteStorageGatewayWithContext(ctx context.Context, request *DeleteStorageGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteStorageGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStorageGateway"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStorageGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a storage volume.
//
// @param request - DeleteStorageVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStorageVolumeResponse
func (client *Client) DeleteStorageVolumeWithContext(ctx context.Context, request *DeleteStorageVolumeRequest, runtime *dara.RuntimeOptions) (_result *DeleteStorageVolumeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VolumeId) {
		query["VolumeId"] = request.VolumeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStorageVolume"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStorageVolumeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a vSwitch.
//
// Description:
//
// Before you delete a vSwitch, make sure that no instances exist in the vSwitch.
//
// @param request - DeleteVSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVSwitchResponse
func (client *Client) DeleteVSwitchWithContext(ctx context.Context, request *DeleteVSwitchRequest, runtime *dara.RuntimeOptions) (_result *DeleteVSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVSwitch"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys a shared data group (SDG) to compute instances.
//
// @param tmpReq - DeployInstanceSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployInstanceSDGResponse
func (client *Client) DeployInstanceSDGWithContext(ctx context.Context, tmpReq *DeployInstanceSDGRequest, runtime *dara.RuntimeOptions) (_result *DeployInstanceSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeployInstanceSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentType) {
		query["DeploymentType"] = request.DeploymentType
	}

	if !dara.IsNil(request.DiskAccessProtocol) {
		query["DiskAccessProtocol"] = request.DiskAccessProtocol
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.SDGId) {
		query["SDGId"] = request.SDGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployInstanceSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployInstanceSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys shared data groups (SDGs).
//
// @param tmpReq - DeploySDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeploySDGResponse
func (client *Client) DeploySDGWithContext(ctx context.Context, tmpReq *DeploySDGRequest, runtime *dara.RuntimeOptions) (_result *DeploySDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeploySDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeploySDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeploySDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about images of the Android in Container (AIC) instance.
//
// @param request - DescribeAICImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAICImagesResponse
func (client *Client) DescribeAICImagesWithContext(ctx context.Context, request *DescribeAICImagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAICImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.MaxDate) {
		query["MaxDate"] = request.MaxDate
	}

	if !dara.IsNil(request.MinDate) {
		query["MinDate"] = request.MinDate
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAICImages"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAICImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about ARM servers and Android in Container (AIC) instances.
//
// @param tmpReq - DescribeARMServerInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeARMServerInstancesResponse
func (client *Client) DescribeARMServerInstancesWithContext(ctx context.Context, tmpReq *DescribeARMServerInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeARMServerInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeARMServerInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AICSpecs) {
		request.AICSpecsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AICSpecs, dara.String("AICSpecs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EnsRegionIds) {
		request.EnsRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnsRegionIds, dara.String("EnsRegionIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ServerIds) {
		request.ServerIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerIds, dara.String("ServerIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ServerSpecs) {
		request.ServerSpecsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerSpecs, dara.String("ServerSpecs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.States) {
		request.StatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.States, dara.String("States"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeARMServerInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeARMServerInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic properties, resources, and container status of an application.
//
// @param request - DescribeApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationResponse
func (client *Client) DescribeApplicationWithContext(ctx context.Context, request *DescribeApplicationRequest, runtime *dara.RuntimeOptions) (_result *DescribeApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppVersions) {
		query["AppVersions"] = request.AppVersions
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.OutDetailStatParams) {
		query["OutDetailStatParams"] = request.OutDetailStatParams
	}

	if !dara.IsNil(request.ResourceSelector) {
		query["ResourceSelector"] = request.ResourceSelector
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplication"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metering method for the bandwidth.
//
// @param request - DescribeBandwitdhByInternetChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBandwitdhByInternetChargeTypeResponse
func (client *Client) DescribeBandwitdhByInternetChargeTypeWithContext(ctx context.Context, request *DescribeBandwitdhByInternetChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeBandwitdhByInternetChargeTypeResponse, _err error) {
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

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBandwitdhByInternetChargeType"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The specifications of available resources are queried when you create a disk.
//
// @param tmpReq - DescribeCloudDiskTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudDiskTypesResponse
func (client *Client) DescribeCloudDiskTypesWithContext(ctx context.Context, tmpReq *DescribeCloudDiskTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudDiskTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeCloudDiskTypesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EnsRegionIds) {
		request.EnsRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnsRegionIds, dara.String("EnsRegionIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudDiskTypes"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudDiskTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Container Service for Kubernetes (ACK) edge clusters.
//
// Description:
//
//	You can call this operation up to 100 times per second per account.
//
// @param request - DescribeClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResponse
func (client *Client) DescribeClusterWithContext(ctx context.Context, request *DescribeClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCluster"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate of a Container Service for Kubernetes (ACK) edge cluster.
//
// Description:
//
//	The maximum number of times that each user can call this operation per second is 100.
//
// @param request - DescribeClusterKubeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterKubeConfigResponse
func (client *Client) DescribeClusterKubeConfigWithContext(ctx context.Context, request *DescribeClusterKubeConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterKubeConfigResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterKubeConfig"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterKubeConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of creating an instance.
//
// @param request - DescribeCreatePrePaidInstanceResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCreatePrePaidInstanceResultResponse
func (client *Client) DescribeCreatePrePaidInstanceResultWithContext(ctx context.Context, request *DescribeCreatePrePaidInstanceResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeCreatePrePaidInstanceResultResponse, _err error) {
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
		Action:      dara.String("DescribeCreatePrePaidInstanceResult"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCreatePrePaidInstanceResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution status of data files on edge instances of an application.
//
// @param tmpReq - DescribeDataDistResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataDistResultResponse
func (client *Client) DescribeDataDistResultWithContext(ctx context.Context, tmpReq *DescribeDataDistResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataDistResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeDataDistResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EnsRegionIds) {
		request.EnsRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnsRegionIds, dara.String("EnsRegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataNames) {
		query["DataNames"] = request.DataNames
	}

	if !dara.IsNil(request.DataVersions) {
		query["DataVersions"] = request.DataVersions
	}

	if !dara.IsNil(request.EnsRegionIdsShrink) {
		query["EnsRegionIds"] = request.EnsRegionIdsShrink
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.MaxDate) {
		query["MaxDate"] = request.MaxDate
	}

	if !dara.IsNil(request.MinDate) {
		query["MinDate"] = request.MinDate
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
		Action:      dara.String("DescribeDataDistResult"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataDistResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URLs of application data on file servers and returns the file servers on which data is pushed.
//
// @param request - DescribeDataDownloadURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataDownloadURLResponse
func (client *Client) DescribeDataDownloadURLWithContext(ctx context.Context, request *DescribeDataDownloadURLRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataDownloadURLResponse, _err error) {
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
		Action:      dara.String("DescribeDataDownloadURL"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataDownloadURLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the push status of application data files on Edge Node Service (ENS) nodes.
//
// @param request - DescribeDataPushResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataPushResultResponse
func (client *Client) DescribeDataPushResultWithContext(ctx context.Context, request *DescribeDataPushResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataPushResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DataNames) {
		query["DataNames"] = request.DataNames
	}

	if !dara.IsNil(request.DataVersions) {
		query["DataVersions"] = request.DataVersions
	}

	if !dara.IsNil(request.MaxDate) {
		query["MaxDate"] = request.MaxDate
	}

	if !dara.IsNil(request.MinDate) {
		query["MinDate"] = request.MinDate
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionIds) {
		query["RegionIds"] = request.RegionIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataPushResult"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataPushResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the properties of instances and virtual devices in a specific edge application.
//
// @param request - DescribeDeviceServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceServiceResponse
func (client *Client) DescribeDeviceServiceWithContext(ctx context.Context, request *DescribeDeviceServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceServiceResponse, _err error) {
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
		Action:      dara.String("DescribeDeviceService"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the disk IOPS monitoring data.
//
// @param request - DescribeDiskIopsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiskIopsListResponse
func (client *Client) DescribeDiskIopsListWithContext(ctx context.Context, request *DescribeDiskIopsListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiskIopsListResponse, _err error) {
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
		Action:      dara.String("DescribeDiskIopsList"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiskIopsListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more disks.
//
// @param request - DescribeDisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisksResponse
func (client *Client) DescribeDisksWithContext(ctx context.Context, request *DescribeDisksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDisksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.DiskChargeType) {
		query["DiskChargeType"] = request.DiskChargeType
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.DiskIds) {
		query["DiskIds"] = request.DiskIds
	}

	if !dara.IsNil(request.DiskName) {
		query["DiskName"] = request.DiskName
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderByParams) {
		query["OrderByParams"] = request.OrderByParams
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
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
		Action:      dara.String("DescribeDisks"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDisksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 描述商品code
//
// @param request - DescribeEnsCommodityCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsCommodityCodeResponse
func (client *Client) DescribeEnsCommodityCodeWithContext(ctx context.Context, request *DescribeEnsCommodityCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsCommodityCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsCommodityCode"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsCommodityCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模块code
//
// @param request - DescribeEnsCommodityModuleCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsCommodityModuleCodeResponse
func (client *Client) DescribeEnsCommodityModuleCodeWithContext(ctx context.Context, request *DescribeEnsCommodityModuleCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsCommodityModuleCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.ModuleCode) {
		query["ModuleCode"] = request.ModuleCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsCommodityModuleCode"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsCommodityModuleCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries elastic IP addresses (EIPs).
//
// @param request - DescribeEnsEipAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsEipAddressesResponse
func (client *Client) DescribeEnsEipAddressesWithContext(ctx context.Context, request *DescribeEnsEipAddressesRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsEipAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationId) {
		query["AllocationId"] = request.AllocationId
	}

	if !dara.IsNil(request.AssociatedInstanceId) {
		query["AssociatedInstanceId"] = request.AssociatedInstanceId
	}

	if !dara.IsNil(request.AssociatedInstanceType) {
		query["AssociatedInstanceType"] = request.AssociatedInstanceType
	}

	if !dara.IsNil(request.EipAddress) {
		query["EipAddress"] = request.EipAddress
	}

	if !dara.IsNil(request.EipName) {
		query["EipName"] = request.EipName
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
	}

	if !dara.IsNil(request.IcmpReplyEnabled) {
		query["IcmpReplyEnabled"] = request.IcmpReplyEnabled
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Standby) {
		query["Standby"] = request.Standby
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsEipAddresses"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsEipAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regions in which ENS resources can be created.
//
// @param request - DescribeEnsNetDistrictRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsNetDistrictResponse
func (client *Client) DescribeEnsNetDistrictWithContext(ctx context.Context, request *DescribeEnsNetDistrictRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsNetDistrictResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetDistrictCode) {
		query["NetDistrictCode"] = request.NetDistrictCode
	}

	if !dara.IsNil(request.NetDistrictCodeNode) {
		query["NetDistrictCodeNode"] = request.NetDistrictCodeNode
	}

	if !dara.IsNil(request.NetLevelCode) {
		query["NetLevelCode"] = request.NetLevelCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsNetDistrict"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsNetDistrictResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about ISPs and number of ISPs in an area.
//
// @param request - DescribeEnsNetSaleDistrictRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsNetSaleDistrictResponse
func (client *Client) DescribeEnsNetSaleDistrictWithContext(ctx context.Context, request *DescribeEnsNetSaleDistrictRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsNetSaleDistrictResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetDistrictCode) {
		query["NetDistrictCode"] = request.NetDistrictCode
	}

	if !dara.IsNil(request.NetLevelCode) {
		query["NetLevelCode"] = request.NetLevelCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsNetSaleDistrict"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsNetSaleDistrictResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a node supports IPv6.
//
// @param request - DescribeEnsRegionIdIpv6InfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsRegionIdIpv6InfoResponse
func (client *Client) DescribeEnsRegionIdIpv6InfoWithContext(ctx context.Context, request *DescribeEnsRegionIdIpv6InfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsRegionIdIpv6InfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsRegionIdIpv6Info"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsRegionIdIpv6InfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries node resources.
//
// Description:
//
// ***
//
// @param request - DescribeEnsRegionIdResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsRegionIdResourceResponse
func (client *Client) DescribeEnsRegionIdResourceWithContext(ctx context.Context, request *DescribeEnsRegionIdResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsRegionIdResourceResponse, _err error) {
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

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.OrderByParams) {
		query["OrderByParams"] = request.OrderByParams
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
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
		Action:      dara.String("DescribeEnsRegionIdResource"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsRegionIdResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Edge Node Service (ENS) nodes that you can use.
//
// @param request - DescribeEnsRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsRegionsResponse
func (client *Client) DescribeEnsRegionsWithContext(ctx context.Context, request *DescribeEnsRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsRegions"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage summary of ENS virtual machines (VMs), disks, and networks.
//
// @param request - DescribeEnsResourceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsResourceUsageResponse
func (client *Client) DescribeEnsResourceUsageWithContext(ctx context.Context, request *DescribeEnsResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsResourceUsageResponse, _err error) {
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
		Action:      dara.String("DescribeEnsResourceUsage"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsResourceUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries route entries.
//
// @param request - DescribeEnsRouteEntryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsRouteEntryListResponse
func (client *Client) DescribeEnsRouteEntryListWithContext(ctx context.Context, request *DescribeEnsRouteEntryListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsRouteEntryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.NextHopId) {
		query["NextHopId"] = request.NextHopId
	}

	if !dara.IsNil(request.NextHopType) {
		query["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RouteEntryId) {
		query["RouteEntryId"] = request.RouteEntryId
	}

	if !dara.IsNil(request.RouteEntryName) {
		query["RouteEntryName"] = request.RouteEntryName
	}

	if !dara.IsNil(request.RouteEntryType) {
		query["RouteEntryType"] = request.RouteEntryType
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsRouteEntryList"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsRouteEntryListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries route tables.
//
// @param request - DescribeEnsRouteTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsRouteTablesResponse
func (client *Client) DescribeEnsRouteTablesWithContext(ctx context.Context, request *DescribeEnsRouteTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsRouteTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociateType) {
		query["AssociateType"] = request.AssociateType
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	if !dara.IsNil(request.RouteTableName) {
		query["RouteTableName"] = request.RouteTableName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsRouteTables"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsRouteTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示配置的售卖约束信息
//
// @param request - DescribeEnsSaleControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsSaleControlResponse
func (client *Client) DescribeEnsSaleControlWithContext(ctx context.Context, request *DescribeEnsSaleControlRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsSaleControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUidAccount) {
		query["AliUidAccount"] = request.AliUidAccount
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CustomAccount) {
		query["CustomAccount"] = request.CustomAccount
	}

	if !dara.IsNil(request.ModuleCode) {
		query["ModuleCode"] = request.ModuleCode
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsSaleControl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsSaleControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取可用资源
//
// @param request - DescribeEnsSaleControlAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsSaleControlAvailableResourceResponse
func (client *Client) DescribeEnsSaleControlAvailableResourceWithContext(ctx context.Context, request *DescribeEnsSaleControlAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsSaleControlAvailableResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CustomAccount) {
		query["CustomAccount"] = request.CustomAccount
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsSaleControlAvailableResource"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsSaleControlAvailableResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取售卖约束库存
//
// @param request - DescribeEnsSaleControlStockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnsSaleControlStockResponse
func (client *Client) DescribeEnsSaleControlStockWithContext(ctx context.Context, request *DescribeEnsSaleControlStockRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnsSaleControlStockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUidAccount) {
		query["AliUidAccount"] = request.AliUidAccount
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CustomAccount) {
		query["CustomAccount"] = request.CustomAccount
	}

	if !dara.IsNil(request.ModuleCode) {
		query["ModuleCode"] = request.ModuleCode
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnsSaleControlStock"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnsSaleControlStockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the EPN bandwidth usage.
//
// @param request - DescribeEpnBandWidthDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEpnBandWidthDataResponse
func (client *Client) DescribeEpnBandWidthDataWithContext(ctx context.Context, request *DescribeEpnBandWidthDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeEpnBandWidthDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.NetworkingModel) {
		query["NetworkingModel"] = request.NetworkingModel
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEpnBandWidthData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEpnBandWidthDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metering method of the EPN bandwidth within a time period.
//
// @param request - DescribeEpnBandwitdhByInternetChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEpnBandwitdhByInternetChargeTypeResponse
func (client *Client) DescribeEpnBandwitdhByInternetChargeTypeWithContext(ctx context.Context, request *DescribeEpnBandwitdhByInternetChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEpnBandwitdhByInternetChargeTypeResponse, _err error) {
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

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.NetworkingModel) {
		query["NetworkingModel"] = request.NetworkingModel
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEpnBandwitdhByInternetChargeType"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEpnBandwitdhByInternetChargeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries networking information about an EPN instance.
//
// Description:
//
// In internal networking mode, the value of Instances is empty in the response. In public networking mode, the value of VSwitches is empty in the response.
//
// @param request - DescribeEpnInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEpnInstanceAttributeResponse
func (client *Client) DescribeEpnInstanceAttributeWithContext(ctx context.Context, request *DescribeEpnInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEpnInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEpnInstanceAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEpnInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries EPN instances.
//
// @param request - DescribeEpnInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEpnInstancesResponse
func (client *Client) DescribeEpnInstancesWithContext(ctx context.Context, request *DescribeEpnInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeEpnInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	if !dara.IsNil(request.EPNInstanceName) {
		query["EPNInstanceName"] = request.EPNInstanceName
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
		Action:      dara.String("DescribeEpnInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEpnInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metering data of the edge private network (EPN).
//
// @param request - DescribeEpnMeasurementDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEpnMeasurementDataResponse
func (client *Client) DescribeEpnMeasurementDataWithContext(ctx context.Context, request *DescribeEpnMeasurementDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeEpnMeasurementDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEpnMeasurementData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEpnMeasurementDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the export result of an image.
//
// Description:
//
//	  You can call this operation to query information about all custom images in your account. The information include the image properties, image export status, and the Object Storage Service (OSS) download links.
//
//		- Empty strings are returned for images that are not exported.
//
//		- The download links may become invalid if you delete objects in OSS.
//
// @param request - DescribeExportImageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExportImageInfoResponse
func (client *Client) DescribeExportImageInfoWithContext(ctx context.Context, request *DescribeExportImageInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeExportImageInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
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
		Action:      dara.String("DescribeExportImageInfo"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExportImageInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the export status of an image.
//
// @param request - DescribeExportImageStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExportImageStatusResponse
func (client *Client) DescribeExportImageStatusWithContext(ctx context.Context, request *DescribeExportImageStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeExportImageStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExportImageStatus"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExportImageStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about file systems.
//
// @param request - DescribeFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileSystemsResponse
func (client *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileSystemsResponse, _err error) {
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
		Action:      dara.String("DescribeFileSystems"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Destination Network Address Translation (DNAT) entries that you created.
//
// @param request - DescribeForwardTableEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeForwardTableEntriesResponse
func (client *Client) DescribeForwardTableEntriesWithContext(ctx context.Context, request *DescribeForwardTableEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeForwardTableEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalIp) {
		query["ExternalIp"] = request.ExternalIp
	}

	if !dara.IsNil(request.ForwardEntryId) {
		query["ForwardEntryId"] = request.ForwardEntryId
	}

	if !dara.IsNil(request.ForwardEntryName) {
		query["ForwardEntryName"] = request.ForwardEntryName
	}

	if !dara.IsNil(request.InternalIp) {
		query["InternalIp"] = request.InternalIp
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
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
		Action:      dara.String("DescribeForwardTableEntries"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeForwardTableEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries high-availability virtual IP addresses (HAVIPs).
//
// @param request - DescribeHaVipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHaVipsResponse
func (client *Client) DescribeHaVipsWithContext(ctx context.Context, request *DescribeHaVipsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHaVipsResponse, _err error) {
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
		Action:      dara.String("DescribeHaVips"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHaVipsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instance system events.
//
// Description:
//
//	You must specify an event type to query. You can query multiple event types at the same time.
//
// @param tmpReq - DescribeHistoryEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryEventsResponse
func (client *Client) DescribeHistoryEventsWithContext(ctx context.Context, tmpReq *DescribeHistoryEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeHistoryEventsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventLevels) {
		request.EventLevelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventLevels, dara.String("EventLevels"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.EventStatus) {
		request.EventStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventStatus, dara.String("EventStatus"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.EventTypes) {
		request.EventTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventTypes, dara.String("EventTypes"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ResourceIds) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, dara.String("ResourceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EventLevelsShrink) {
		query["EventLevels"] = request.EventLevelsShrink
	}

	if !dara.IsNil(request.EventStatusShrink) {
		query["EventStatus"] = request.EventStatusShrink
	}

	if !dara.IsNil(request.EventTypesShrink) {
		query["EventTypes"] = request.EventTypesShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceIdsShrink) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryEvents"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available images.
//
// @param request - DescribeImageInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageInfosResponse
func (client *Client) DescribeImageInfosWithContext(ctx context.Context, request *DescribeImageInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImageInfos"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accounts with which you share an image specified by the ImageId parameter.
//
// @param request - DescribeImageSharePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageSharePermissionResponse
func (client *Client) DescribeImageSharePermissionWithContext(ctx context.Context, request *DescribeImageSharePermissionRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageSharePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunId) {
		query["AliyunId"] = request.AliyunId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
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
		Action:      dara.String("DescribeImageSharePermission"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageSharePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available images.
//
// @param request - DescribeImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImagesResponse
func (client *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImages"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query whether auto-renewal is enabled for an instance.
//
// @param request - DescribeInstanceAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAutoRenewAttributeResponse
func (client *Client) DescribeInstanceAutoRenewAttributeWithContext(ctx context.Context, request *DescribeInstanceAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAutoRenewAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAutoRenewAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAutoRenewAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed bandwidth data of an instance, which is collected every 5 minutes.
//
// Description:
//
//	  You can call this operation up to 800 times per second per account.
//
//		- You can call this operation up to 100 times per second per user.
//
//		- You can specify multiple request parameters to filter query results. Specified request parameters have logical AND relations. Only the specified parameters are included in the filter conditions. However, if InstanceIds is set to an empty JSON array, this parameter is regarded as a valid filter condition and an empty result is returned.
//
// @param request - DescribeInstanceBandwidthDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceBandwidthDetailResponse
func (client *Client) DescribeInstanceBandwidthDetailWithContext(ctx context.Context, request *DescribeInstanceBandwidthDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceBandwidthDetailResponse, _err error) {
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
		Action:      dara.String("DescribeInstanceBandwidthDetail"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceBandwidthDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the boot configuration of a heterogeneous PC Farm bare metal instance.
//
// Description:
//
// Queries the boot configuration of a heterogeneous PC Farm bare metal instance.
//
// @param request - DescribeInstanceBootConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceBootConfigurationResponse
func (client *Client) DescribeInstanceBootConfigurationWithContext(ctx context.Context, request *DescribeInstanceBootConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceBootConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BootSet) {
		query["BootSet"] = request.BootSet
	}

	if !dara.IsNil(request.BootType) {
		query["BootType"] = request.BootType
	}

	if !dara.IsNil(request.DiskSet) {
		query["DiskSet"] = request.DiskSet
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceBootConfiguration"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceBootConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the vCPU and memory usage of an instance.
//
// @param request - DescribeInstanceMonitorDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceMonitorDataResponse
func (client *Client) DescribeInstanceMonitorDataWithContext(ctx context.Context, request *DescribeInstanceMonitorDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceMonitorDataResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceMonitorData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceMonitorDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries shared data groups (SDGs) that are mounted to an Android in Container (AIC) instance.
//
// @param tmpReq - DescribeInstanceSDGStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSDGStatusResponse
func (client *Client) DescribeInstanceSDGStatusWithContext(ctx context.Context, tmpReq *DescribeInstanceSDGStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSDGStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeInstanceSDGStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SDGIds) {
		request.SDGIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SDGIds, dara.String("SDGIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSDGStatus"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSDGStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Virtual Network Computing (VNC) URL of an Edge Node Service (ENS) instance.
//
// @param request - DescribeInstanceVncUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceVncUrlResponse
func (client *Client) DescribeInstanceVncUrlWithContext(ctx context.Context, request *DescribeInstanceVncUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceVncUrlResponse, _err error) {
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
		Action:      dara.String("DescribeInstanceVncUrl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceVncUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to view the details of one or more instances.
//
// Description:
//
//	  You can call this operation up to 800 times per second per account.
//
//		- You can call this operation up to 100 times per second per user.
//
//		- You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are included in the filter conditions. However, if InstanceIds is set to an empty JSON array, it is regarded as a valid filter condition and an empty result is returned.
//
// @param tmpReq - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithContext(ctx context.Context, tmpReq *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EipAddresses) {
		request.EipAddressesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EipAddresses, dara.String("EipAddresses"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ServiceStatus) {
		request.ServiceStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceStatus, dara.String("ServiceStatus"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EipAddressesShrink) {
		query["EipAddresses"] = request.EipAddressesShrink
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
	}

	if !dara.IsNil(request.EnsServiceId) {
		query["EnsServiceId"] = request.EnsServiceId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceResourceType) {
		query["InstanceResourceType"] = request.InstanceResourceType
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.IntranetIp) {
		query["IntranetIp"] = request.IntranetIp
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.OrderByParams) {
		query["OrderByParams"] = request.OrderByParams
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.ServiceStatusShrink) {
		query["ServiceStatus"] = request.ServiceStatusShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more key pairs.
//
// @param request - DescribeKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKeyPairsResponse
func (client *Client) DescribeKeyPairsWithContext(ctx context.Context, request *DescribeKeyPairsRequest, runtime *dara.RuntimeOptions) (_result *DescribeKeyPairsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairId) {
		query["KeyPairId"] = request.KeyPairId
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
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
		Action:      dara.String("DescribeKeyPairs"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKeyPairsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about an Edge Load Balancer (ELB) instance.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeLoadBalancerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerAttributeResponse
func (client *Client) DescribeLoadBalancerAttributeWithContext(ctx context.Context, request *DescribeLoadBalancerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of an HTTP listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DescribeLoadBalancerHTTPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerHTTPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithContext(ctx context.Context, request *DescribeLoadBalancerHTTPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerHTTPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerHTTPListenerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an HTTPS listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DescribeLoadBalancerHTTPSListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerHTTPSListenerAttributeResponse
func (client *Client) DescribeLoadBalancerHTTPSListenerAttributeWithContext(ctx context.Context, request *DescribeLoadBalancerHTTPSListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerHTTPSListenerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries monitoring data of an edge load balancer (ELB) instance.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DescribeLoadBalancerListenMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerListenMonitorResponse
func (client *Client) DescribeLoadBalancerListenMonitorWithContext(ctx context.Context, request *DescribeLoadBalancerListenMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerListenMonitorResponse, _err error) {
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
		Action:      dara.String("DescribeLoadBalancerListenMonitor"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerListenMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries listeners of Edge Load Balancer (ELB) instances.
//
// @param request - DescribeLoadBalancerListenersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerListenersResponse
func (client *Client) DescribeLoadBalancerListenersWithContext(ctx context.Context, request *DescribeLoadBalancerListenersRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerListenersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
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
		Action:      dara.String("DescribeLoadBalancerListeners"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerListenersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications of an Edge Load Balancer (ELB) instance.
//
// @param request - DescribeLoadBalancerSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerSpecResponse
func (client *Client) DescribeLoadBalancerSpecWithContext(ctx context.Context, request *DescribeLoadBalancerSpecRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerSpecResponse, _err error) {
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
		Action:      dara.String("DescribeLoadBalancerSpec"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a Transmission Control Protocol (TCP) listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DescribeLoadBalancerTCPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerTCPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithContext(ctx context.Context, request *DescribeLoadBalancerTCPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerTCPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerTCPListenerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of a User Datagram Protocol (UDP) listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DescribeLoadBalancerUDPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerUDPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithContext(ctx context.Context, request *DescribeLoadBalancerUDPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerUDPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerUDPListenerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Edge Load Balance (ELB) instances that you have created.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DescribeLoadBalancersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancersResponse
func (client *Client) DescribeLoadBalancersWithContext(ctx context.Context, request *DescribeLoadBalancersRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancersResponse, _err error) {
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
		Action:      dara.String("DescribeLoadBalancers"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metering data of the user.
//
// @param request - DescribeMeasurementDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMeasurementDataResponse
func (client *Client) DescribeMeasurementDataWithContext(ctx context.Context, request *DescribeMeasurementDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeMeasurementDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMeasurementData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMeasurementDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about mount targets.
//
// @param request - DescribeMountTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMountTargetsResponse
func (client *Client) DescribeMountTargetsWithContext(ctx context.Context, request *DescribeMountTargetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetName) {
		query["MountTargetName"] = request.MountTargetName
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
		Action:      dara.String("DescribeMountTargets"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据筛选条件获取指定NC属性和资源量信息
//
// @param request - DescribeNCInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNCInformationResponse
func (client *Client) DescribeNCInformationWithContext(ctx context.Context, request *DescribeNCInformationRequest, runtime *dara.RuntimeOptions) (_result *DescribeNCInformationResponse, _err error) {
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
		Action:      dara.String("DescribeNCInformation"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNCInformationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries created Network Address Translation (NAT) gateways.
//
// @param request - DescribeNatGatewaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNatGatewaysResponse
func (client *Client) DescribeNatGatewaysWithContext(ctx context.Context, request *DescribeNatGatewaysRequest, runtime *dara.RuntimeOptions) (_result *DescribeNatGatewaysResponse, _err error) {
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
		Action:      dara.String("DescribeNatGateways"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNatGatewaysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries network access control lists (ACLs).
//
// @param request - DescribeNetworkAclsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkAclsResponse
func (client *Client) DescribeNetworkAclsWithContext(ctx context.Context, request *DescribeNetworkAclsRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkAclsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkAclId) {
		query["NetworkAclId"] = request.NetworkAclId
	}

	if !dara.IsNil(request.NetworkAclName) {
		query["NetworkAclName"] = request.NetworkAclName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkAcls"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkAclsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of a specified network.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeNetworkAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkAttributeResponse
func (client *Client) DescribeNetworkAttributeWithContext(ctx context.Context, request *DescribeNetworkAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Elastic Network Interfaces (ENIs).
//
// @param request - DescribeNetworkInterfacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkInterfacesResponse
func (client *Client) DescribeNetworkInterfacesWithContext(ctx context.Context, request *DescribeNetworkInterfacesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkInterfacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ipv6Address) {
		query["Ipv6Address"] = request.Ipv6Address
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.NetworkInterfaceIds) {
		query["NetworkInterfaceIds"] = request.NetworkInterfaceIds
	}

	if !dara.IsNil(request.NetworkInterfaceName) {
		query["NetworkInterfaceName"] = request.NetworkInterfaceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrimaryIpAddress) {
		query["PrimaryIpAddress"] = request.PrimaryIpAddress
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkInterfaces"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkInterfacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the network list.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeNetworksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworksResponse
func (client *Client) DescribeNetworksWithContext(ctx context.Context, request *DescribeNetworksRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.NetworkIds) {
		query["NetworkIds"] = request.NetworkIds
	}

	if !dara.IsNil(request.NetworkName) {
		query["NetworkName"] = request.NetworkName
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
		Action:      dara.String("DescribeNetworks"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications of resources that can be purchased in subscription billing mode.
//
// @param request - DescribePrePaidInstanceStockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePrePaidInstanceStockResponse
func (client *Client) DescribePrePaidInstanceStockWithContext(ctx context.Context, request *DescribePrePaidInstanceStockRequest, runtime *dara.RuntimeOptions) (_result *DescribePrePaidInstanceStockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataDiskSize) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.SystemDiskSize) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrePaidInstanceStock"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePrePaidInstanceStockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the most recent price of an Edge Node Service (ENS) instance.
//
// @param tmpReq - DescribePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceResponse
func (client *Client) DescribePriceWithContext(ctx context.Context, tmpReq *DescribePriceRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataDisks) {
		request.DataDisksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataDisks, dara.String("DataDisks"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataDisksShrink) {
		query["DataDisks"] = request.DataDisksShrink
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
	}

	if !dara.IsNil(request.DataDisk) {
		query["DataDisk"] = request.DataDisk
	}

	if !dara.IsNil(request.SystemDisk) {
		query["SystemDisk"] = request.SystemDisk
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrice"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the Internet service providers (ISPs) of edge nodes.
//
// @param request - DescribeRegionIspsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionIspsResponse
func (client *Client) DescribeRegionIspsWithContext(ctx context.Context, request *DescribeRegionIspsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionIspsResponse, _err error) {
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
		Action:      dara.String("DescribeRegionIsps"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionIspsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取区域节点资源量信息
//
// @param request - DescribeRegionResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionResourceResponse
func (client *Client) DescribeRegionResourceWithContext(ctx context.Context, request *DescribeRegionResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionResourceResponse, _err error) {
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
		Action:      dara.String("DescribeRegionResource"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源状态变化时间线
//
// @param request - DescribeResourceTimelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceTimelineResponse
func (client *Client) DescribeResourceTimelineWithContext(ctx context.Context, request *DescribeResourceTimelineRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceTimelineResponse, _err error) {
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
		Action:      dara.String("DescribeResourceTimeline"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceTimelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries basic information about shared data groups (SDGs), including node preload information.
//
// @param tmpReq - DescribeSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSDGResponse
func (client *Client) DescribeSDGWithContext(ctx context.Context, tmpReq *DescribeSDGRequest, runtime *dara.RuntimeOptions) (_result *DescribeSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SDGIds) {
		request.SDGIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SDGIds, dara.String("SDGIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SDGIdsShrink) {
		query["SDGIds"] = request.SDGIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deployment status of the shared data group (SDG).
//
// @param tmpReq - DescribeSDGDeploymentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSDGDeploymentStatusResponse
func (client *Client) DescribeSDGDeploymentStatusWithContext(ctx context.Context, tmpReq *DescribeSDGDeploymentStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSDGDeploymentStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSDGDeploymentStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DiskIds) {
		request.DiskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DiskIds, dara.String("DiskIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegionIds) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, dara.String("RegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentType) {
		query["DeploymentType"] = request.DeploymentType
	}

	if !dara.IsNil(request.DiskIdsShrink) {
		query["DiskIds"] = request.DiskIdsShrink
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionIdsShrink) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !dara.IsNil(request.SDGId) {
		query["SDGId"] = request.SDGId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSDGDeploymentStatus"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSDGDeploymentStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询SDG下的共享盘
//
// @param request - DescribeSDGSharedDisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSDGSharedDisksResponse
func (client *Client) DescribeSDGSharedDisksWithContext(ctx context.Context, request *DescribeSDGSharedDisksRequest, runtime *dara.RuntimeOptions) (_result *DescribeSDGSharedDisksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.SdgId) {
		query["SdgId"] = request.SdgId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSDGSharedDisks"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSDGSharedDisksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about created shared data groups (SDGs).
//
// @param tmpReq - DescribeSDGsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSDGsResponse
func (client *Client) DescribeSDGsWithContext(ctx context.Context, tmpReq *DescribeSDGsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSDGsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeSDGsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SDGIds) {
		request.SDGIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SDGIds, dara.String("SDGIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.SDGIdsShrink) {
		query["SDGIds"] = request.SDGIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSDGs"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSDGsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries created secondary public IP addresses.
//
// @param request - DescribeSecondaryPublicIpAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecondaryPublicIpAddressesResponse
func (client *Client) DescribeSecondaryPublicIpAddressesWithContext(ctx context.Context, request *DescribeSecondaryPublicIpAddressesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecondaryPublicIpAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecondaryPublicIpAddress) {
		query["SecondaryPublicIpAddress"] = request.SecondaryPublicIpAddress
	}

	if !dara.IsNil(request.SecondaryPublicIpId) {
		query["SecondaryPublicIpId"] = request.SecondaryPublicIpId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecondaryPublicIpAddresses"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecondaryPublicIpAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rules of a security group.
//
// @param request - DescribeSecurityGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityGroupAttributeResponse
func (client *Client) DescribeSecurityGroupAttributeWithContext(ctx context.Context, request *DescribeSecurityGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityGroupAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about created security groups.
//
// @param request - DescribeSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityGroupsResponse
func (client *Client) DescribeSecurityGroupsWithContext(ctx context.Context, request *DescribeSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityGroupName) {
		query["SecurityGroupName"] = request.SecurityGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityGroups"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom images.
//
// @param request - DescribeSelfImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSelfImagesResponse
func (client *Client) DescribeSelfImagesWithContext(ctx context.Context, request *DescribeSelfImagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSelfImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSelfImages"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSelfImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeServcieSchedule to query the real-time status of the instance device or container that is being occupied by the UUID.
//
// @param request - DescribeServcieScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServcieScheduleResponse
func (client *Client) DescribeServcieScheduleWithContext(ctx context.Context, request *DescribeServcieScheduleRequest, runtime *dara.RuntimeOptions) (_result *DescribeServcieScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PodConfigName) {
		query["PodConfigName"] = request.PodConfigName
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServcieSchedule"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServcieScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of an edge load balancer (ELB) instance based on the listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DescribeServerLoadBalancerListenMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServerLoadBalancerListenMonitorResponse
func (client *Client) DescribeServerLoadBalancerListenMonitorWithContext(ctx context.Context, request *DescribeServerLoadBalancerListenMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeServerLoadBalancerListenMonitorResponse, _err error) {
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
		Action:      dara.String("DescribeServerLoadBalancerListenMonitor"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServerLoadBalancerListenMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request monitoring data of an edge load balancer (ELB) instance.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - DescribeServerLoadBalancerMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServerLoadBalancerMonitorResponse
func (client *Client) DescribeServerLoadBalancerMonitorWithContext(ctx context.Context, request *DescribeServerLoadBalancerMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeServerLoadBalancerMonitorResponse, _err error) {
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
		Action:      dara.String("DescribeServerLoadBalancerMonitor"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServerLoadBalancerMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about snapshots.
//
// @param request - DescribeSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
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

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnapshots"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific source network address translation (SNAT) entry.
//
// @param request - DescribeSnatAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnatAttributeResponse
func (client *Client) DescribeSnatAttributeWithContext(ctx context.Context, request *DescribeSnatAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnatAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnatAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnatAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries source network address translation (SNAT) entries.
//
// @param request - DescribeSnatTableEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnatTableEntriesResponse
func (client *Client) DescribeSnatTableEntriesWithContext(ctx context.Context, request *DescribeSnatTableEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnatTableEntriesResponse, _err error) {
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
		Action:      dara.String("DescribeSnatTableEntries"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnatTableEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries storage gateways.
//
// @param request - DescribeStorageGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStorageGatewayResponse
func (client *Client) DescribeStorageGatewayWithContext(ctx context.Context, request *DescribeStorageGatewayRequest, runtime *dara.RuntimeOptions) (_result *DescribeStorageGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["GatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStorageGateway"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStorageGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries volumes.
//
// @param request - DescribeStorageVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStorageVolumeResponse
func (client *Client) DescribeStorageVolumeWithContext(ctx context.Context, request *DescribeStorageVolumeRequest, runtime *dara.RuntimeOptions) (_result *DescribeStorageVolumeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StorageId) {
		query["StorageId"] = request.StorageId
	}

	if !dara.IsNil(request.VolumeId) {
		query["VolumeId"] = request.VolumeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStorageVolume"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStorageVolumeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth that you use within a specified period of time.
//
// @param request - DescribeUserBandWidthDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserBandWidthDataResponse
func (client *Client) DescribeUserBandWidthDataWithContext(ctx context.Context, request *DescribeUserBandWidthDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserBandWidthDataResponse, _err error) {
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

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserBandWidthData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserBandWidthDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeVSwitchAttributes interface to query the configuration of a specified VSwitch.
//
// @param request - DescribeVSwitchAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVSwitchAttributesResponse
func (client *Client) DescribeVSwitchAttributesWithContext(ctx context.Context, request *DescribeVSwitchAttributesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVSwitchAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVSwitchAttributes"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVSwitchAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about available vSwitches.
//
// @param request - DescribeVSwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVSwitchesResponse
func (client *Client) DescribeVSwitchesWithContext(ctx context.Context, request *DescribeVSwitchesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.EnsRegionIds) {
		query["EnsRegionIds"] = request.EnsRegionIds
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VSwitchName) {
		query["VSwitchName"] = request.VSwitchName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVSwitches"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a pay-as-you-go disk from an Edge Node Service (ENS) instance. You cannot call this operation to detach a disk that is created together with an instance.
//
// @param request - DetachDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDiskResponse
func (client *Client) DetachDiskWithContext(ctx context.Context, request *DetachDiskRequest, runtime *dara.RuntimeOptions) (_result *DetachDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDisk"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a shared data group (SDG).
//
// @param tmpReq - DetachInstanceSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachInstanceSDGResponse
func (client *Client) DetachInstanceSDGWithContext(ctx context.Context, tmpReq *DetachInstanceSDGRequest, runtime *dara.RuntimeOptions) (_result *DetachInstanceSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachInstanceSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.SDGId) {
		query["SDGId"] = request.SDGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachInstanceSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachInstanceSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach an elastic network interface (ENI) from an instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - You cannot detach a primary ENI from an instance.
//
//   - The ENI must be in the InUse state.
//
//   - The instances are in the Stopped state.
//
//   - This operation is an asynchronous operation. After this operation is called to detach an ENI, you can check the state of the ENI to determine whether the ENI is detached.
//
// @param request - DetachNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachNetworkInterfaceResponse
func (client *Client) DetachNetworkInterfaceWithContext(ctx context.Context, request *DetachNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *DetachNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachNetworkInterface"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Distributes pushed data to the Edge Node Service (ENS) instances of the application. You can specify multiple canary release policies, decompress files, and restart containers.
//
// @param request - DistApplicationDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DistApplicationDataResponse
func (client *Client) DistApplicationDataWithContext(ctx context.Context, request *DistApplicationDataRequest, runtime *dara.RuntimeOptions) (_result *DistApplicationDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.DistStrategy) {
		query["DistStrategy"] = request.DistStrategy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DistApplicationData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DistApplicationDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates the instance across nodes after an O\\\\\\&M event occurs on an instance.
//
// Description:
//
//	This O\\&M operation is supported only by the Instance:SystemUpgrade.Migrate event.
//
// @param request - EventMigrateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EventMigrateInstanceResponse
func (client *Client) EventMigrateInstanceWithContext(ctx context.Context, request *EventMigrateInstanceRequest, runtime *dara.RuntimeOptions) (_result *EventMigrateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataPolicy) {
		query["DataPolicy"] = request.DataPolicy
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.OpsType) {
		query["OpsType"] = request.OpsType
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PlanTime) {
		query["PlanTime"] = request.PlanTime
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EventMigrateInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EventMigrateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restart the instance across nodes after an O\\\\\\&M event occurs on an instance.
//
// Description:
//
//	This O\\&M operation supports only the Instance:SystemMaintenance.Reboot event
//
// @param request - EventRebootInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EventRebootInstanceResponse
func (client *Client) EventRebootInstanceWithContext(ctx context.Context, request *EventRebootInstanceRequest, runtime *dara.RuntimeOptions) (_result *EventRebootInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.OpsType) {
		query["OpsType"] = request.OpsType
	}

	if !dara.IsNil(request.PlanTime) {
		query["PlanTime"] = request.PlanTime
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EventRebootInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EventRebootInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # The event that is used to immediately redeploy a specified resource or by appointment
//
// Description:
//
//	This O\\&M operation supports only the following event types: Instance:SystemFailure.Redeploy (instance redeployment due to system problems) and Instance:SystemMaintenance.Redeploy (instance redeployment due to system maintenance).
//
// @param request - EventRedeployInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EventRedeployInstanceResponse
func (client *Client) EventRedeployInstanceWithContext(ctx context.Context, request *EventRedeployInstanceRequest, runtime *dara.RuntimeOptions) (_result *EventRedeployInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.OpsType) {
		query["OpsType"] = request.OpsType
	}

	if !dara.IsNil(request.PlanTime) {
		query["PlanTime"] = request.PlanTime
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EventRedeployInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EventRedeployInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports billing details within a specific time range.
//
// @param request - ExportBillDetailDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportBillDetailDataResponse
func (client *Client) ExportBillDetailDataWithContext(ctx context.Context, request *ExportBillDetailDataRequest, runtime *dara.RuntimeOptions) (_result *ExportBillDetailDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportBillDetailData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportBillDetailDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports a custom image to an Object Storage Service (OSS) bucket in the same region.
//
// @param request - ExportImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportImageResponse
func (client *Client) ExportImageWithContext(ctx context.Context, request *ExportImageRequest, runtime *dara.RuntimeOptions) (_result *ExportImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.OSSBucket) {
		query["OSSBucket"] = request.OSSBucket
	}

	if !dara.IsNil(request.OSSPrefix) {
		query["OSSPrefix"] = request.OSSPrefix
	}

	if !dara.IsNil(request.OSSRegionId) {
		query["OSSRegionId"] = request.OSSRegionId
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportImage"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports the metering data within a specific time range.
//
// @param request - ExportMeasurementDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportMeasurementDataResponse
func (client *Client) ExportMeasurementDataWithContext(ctx context.Context, request *ExportMeasurementDataRequest, runtime *dara.RuntimeOptions) (_result *ExportMeasurementDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportMeasurementData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportMeasurementDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control list (ACL) of a bucket.
//
// @param request - GetBucketAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBucketAclResponse
func (client *Client) GetBucketAclWithContext(ctx context.Context, request *GetBucketAclRequest, runtime *dara.RuntimeOptions) (_result *GetBucketAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBucketAcl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBucketAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a bucket.
//
// @param request - GetBucketInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBucketInfoResponse
func (client *Client) GetBucketInfoWithContext(ctx context.Context, request *GetBucketInfoRequest, runtime *dara.RuntimeOptions) (_result *GetBucketInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBucketInfo"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBucketInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries lifecycle rules.
//
// @param request - GetBucketLifecycleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBucketLifecycleResponse
func (client *Client) GetBucketLifecycleWithContext(ctx context.Context, request *GetBucketLifecycleRequest, runtime *dara.RuntimeOptions) (_result *GetBucketLifecycleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBucketLifecycle"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBucketLifecycleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage usage in the previous billing cycle and the cumulative number of calls in this month.
//
// @param request - GetOssStorageAndAccByBucketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssStorageAndAccByBucketsResponse
func (client *Client) GetOssStorageAndAccByBucketsWithContext(ctx context.Context, request *GetOssStorageAndAccByBucketsRequest, runtime *dara.RuntimeOptions) (_result *GetOssStorageAndAccByBucketsResponse, _err error) {
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
		Action:      dara.String("GetOssStorageAndAccByBuckets"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssStorageAndAccByBucketsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage and bandwidth usage within a specific time range.
//
// Description:
//
// The query and aggregation granularity of bandwidth and storage usage cannot exceed one day. Data aggregation is to collect the maximum values of usage data within a period of time.
//
// @param request - GetOssUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssUsageDataResponse
func (client *Client) GetOssUsageDataWithContext(ctx context.Context, request *GetOssUsageDataRequest, runtime *dara.RuntimeOptions) (_result *GetOssUsageDataResponse, _err error) {
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
		Action:      dara.String("GetOssUsageData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssUsageDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call ImportImage to import your image file to the cloud server.
//
// @param tmpReq - ImportImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportImageResponse
func (client *Client) ImportImageWithContext(ctx context.Context, tmpReq *ImportImageRequest, runtime *dara.RuntimeOptions) (_result *ImportImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DiskDeviceMapping) {
		request.DiskDeviceMappingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DiskDeviceMapping, dara.String("DiskDeviceMapping"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Architecture) {
		query["Architecture"] = request.Architecture
	}

	if !dara.IsNil(request.ComputeType) {
		query["ComputeType"] = request.ComputeType
	}

	if !dara.IsNil(request.DiskDeviceMappingShrink) {
		query["DiskDeviceMapping"] = request.DiskDeviceMappingShrink
	}

	if !dara.IsNil(request.ImageFormat) {
		query["ImageFormat"] = request.ImageFormat
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.LicenseType) {
		query["LicenseType"] = request.LicenseType
	}

	if !dara.IsNil(request.OSSBucket) {
		query["OSSBucket"] = request.OSSBucket
	}

	if !dara.IsNil(request.OSSObject) {
		query["OSSObject"] = request.OSSObject
	}

	if !dara.IsNil(request.OSSRegion) {
		query["OSSRegion"] = request.OSSRegion
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
	}

	if !dara.IsNil(request.OSVersion) {
		query["OSVersion"] = request.OSVersion
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.TargetOSSRegionId) {
		query["TargetOSSRegionId"] = request.TargetOSSRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportImage"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports the public key of a Rivest–Shamir–Adleman (RSA)-encrypted key pair that is generated by a third-party tool.
//
// Description:
//
//	  After the key pair is imported, ENS stores the public key. You must securely store the private key.
//
//		- The key pair can be only in the ssh-rsa format.
//
// @param request - ImportKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportKeyPairResponse
func (client *Client) ImportKeyPairWithContext(ctx context.Context, request *ImportKeyPairRequest, runtime *dara.RuntimeOptions) (_result *ImportKeyPairResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.PublicKeyBody) {
		query["PublicKeyBody"] = request.PublicKeyBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportKeyPair"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns public IP addresses to an EPN instance.
//
// @param request - JoinPublicIpsToEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinPublicIpsToEpnInstanceResponse
func (client *Client) JoinPublicIpsToEpnInstanceWithContext(ctx context.Context, request *JoinPublicIpsToEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *JoinPublicIpsToEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	if !dara.IsNil(request.InstanceInfos) {
		query["InstanceInfos"] = request.InstanceInfos
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JoinPublicIpsToEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinPublicIpsToEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an Edge Node Service (ENS) instance to a specified security group.
//
// Description:
//
// Before you call this operation to add an instance to a security group, make sure that the instance is in the Stopped or Running state.
//
// @param request - JoinSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinSecurityGroupResponse
func (client *Client) JoinSecurityGroupWithContext(ctx context.Context, request *JoinSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *JoinSecurityGroupResponse, _err error) {
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

	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JoinSecurityGroup"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs networking.
//
// Description:
//
// For the Internal Connection mode and the Intelligent Acceleration and Internal Connection mode, instances of the vSwitch take effect automatically. You do not need to manually add instances. For public connections such as intelligent acceleration, you need to call an operation to manually add the instances to Internet-facing instances.
//
// @param request - JoinVSwitchesToEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinVSwitchesToEpnInstanceResponse
func (client *Client) JoinVSwitchesToEpnInstanceWithContext(ctx context.Context, request *JoinVSwitchesToEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *JoinVSwitchesToEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	if !dara.IsNil(request.VSwitchesInfo) {
		query["VSwitchesInfo"] = request.VSwitchesInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JoinVSwitchesToEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinVSwitchesToEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an instance from a security group.
//
// Description:
//
// Before you remove an instance from a security group, the instance must be in the Stopped or Running state.
//
// @param request - LeaveSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LeaveSecurityGroupResponse
func (client *Client) LeaveSecurityGroupWithContext(ctx context.Context, request *LeaveSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *LeaveSecurityGroupResponse, _err error) {
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

	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LeaveSecurityGroup"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LeaveSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询公钥下发信息
//
// @param request - ListAICPublicKeyDeliveriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAICPublicKeyDeliveriesResponse
func (client *Client) ListAICPublicKeyDeliveriesWithContext(ctx context.Context, request *ListAICPublicKeyDeliveriesRequest, runtime *dara.RuntimeOptions) (_result *ListAICPublicKeyDeliveriesResponse, _err error) {
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

	if !dara.IsNil(request.KeyGroup) {
		query["KeyGroup"] = request.KeyGroup
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
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
		Action:      dara.String("ListAICPublicKeyDeliveries"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAICPublicKeyDeliveriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有托管的公钥
//
// @param request - ListAICPublicKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAICPublicKeysResponse
func (client *Client) ListAICPublicKeysWithContext(ctx context.Context, request *ListAICPublicKeysRequest, runtime *dara.RuntimeOptions) (_result *ListAICPublicKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyGroup) {
		query["KeyGroup"] = request.KeyGroup
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
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
		Action:      dara.String("ListAICPublicKeys"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAICPublicKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the created applications.
//
// @param request - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithContext(ctx context.Context, request *ListApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppVersions) {
		query["AppVersions"] = request.AppVersions
	}

	if !dara.IsNil(request.ClusterNames) {
		query["ClusterNames"] = request.ClusterNames
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MaxDate) {
		query["MaxDate"] = request.MaxDate
	}

	if !dara.IsNil(request.MinDate) {
		query["MinDate"] = request.MinDate
	}

	if !dara.IsNil(request.OutAppInfoParams) {
		query["OutAppInfoParams"] = request.OutAppInfoParams
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
		Action:      dara.String("ListApplications"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all buckets of a user.
//
// @param request - ListBucketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBucketsResponse
func (client *Client) ListBucketsWithContext(ctx context.Context, request *ListBucketsRequest, runtime *dara.RuntimeOptions) (_result *ListBucketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBuckets"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBucketsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all objects in a bucket.
//
// @param request - ListObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListObjectsResponse
func (client *Client) ListObjectsWithContext(ctx context.Context, request *ListObjectsRequest, runtime *dara.RuntimeOptions) (_result *ListObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.ContinuationToken) {
		query["ContinuationToken"] = request.ContinuationToken
	}

	if !dara.IsNil(request.EncodingType) {
		query["EncodingType"] = request.EncodingType
	}

	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.MaxKeys) {
		query["MaxKeys"] = request.MaxKeys
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	if !dara.IsNil(request.StartAfter) {
		query["StartAfter"] = request.StartAfter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListObjects"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListObjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags that are added to Edge Node Service (ENS) instances.
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
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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
		Version:     dara.String("2017-11-10"),
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
// # AIC公钥登入管理
//
// @param request - ManageAICLoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageAICLoginResponse
func (client *Client) ManageAICLoginWithContext(ctx context.Context, request *ManageAICLoginRequest, runtime *dara.RuntimeOptions) (_result *ManageAICLoginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionName) {
		query["ActionName"] = request.ActionName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.KeyGroup) {
		query["KeyGroup"] = request.KeyGroup
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManageAICLogin"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManageAICLoginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and peak bandwidth of a specified elastic IP address (EIP).
//
// @param request - ModifyEnsEipAddressAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEnsEipAddressAttributeResponse
func (client *Client) ModifyEnsEipAddressAttributeWithContext(ctx context.Context, request *ModifyEnsEipAddressAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyEnsEipAddressAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationId) {
		query["AllocationId"] = request.AllocationId
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEnsEipAddressAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEnsEipAddressAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a custom route.
//
// @param request - ModifyEnsRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEnsRouteEntryResponse
func (client *Client) ModifyEnsRouteEntryWithContext(ctx context.Context, request *ModifyEnsRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *ModifyEnsRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.RouteEntryId) {
		query["RouteEntryId"] = request.RouteEntryId
	}

	if !dara.IsNil(request.RouteEntryName) {
		query["RouteEntryName"] = request.RouteEntryName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEnsRouteEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEnsRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an Edge Private Network (EPN) instance.
//
// @param request - ModifyEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEpnInstanceResponse
func (client *Client) ModifyEpnInstanceWithContext(ctx context.Context, request *ModifyEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	if !dara.IsNil(request.EPNInstanceName) {
		query["EPNInstanceName"] = request.EPNInstanceName
	}

	if !dara.IsNil(request.InternetMaxBandwidthOut) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !dara.IsNil(request.NetworkingModel) {
		query["NetworkingModel"] = request.NetworkingModel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a NAS file system.
//
// @param request - ModifyFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFileSystemResponse
func (client *Client) ModifyFileSystemWithContext(ctx context.Context, request *ModifyFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFileSystem"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Destination Network Address Translation (DNAT) rule.
//
// @param request - ModifyForwardEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyForwardEntryResponse
func (client *Client) ModifyForwardEntryWithContext(ctx context.Context, request *ModifyForwardEntryRequest, runtime *dara.RuntimeOptions) (_result *ModifyForwardEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalIp) {
		query["ExternalIp"] = request.ExternalIp
	}

	if !dara.IsNil(request.ExternalPort) {
		query["ExternalPort"] = request.ExternalPort
	}

	if !dara.IsNil(request.ForwardEntryId) {
		query["ForwardEntryId"] = request.ForwardEntryId
	}

	if !dara.IsNil(request.ForwardEntryName) {
		query["ForwardEntryName"] = request.ForwardEntryName
	}

	if !dara.IsNil(request.HealthCheckPort) {
		query["HealthCheckPort"] = request.HealthCheckPort
	}

	if !dara.IsNil(request.InternalIp) {
		query["InternalIp"] = request.InternalIp
	}

	if !dara.IsNil(request.InternalPort) {
		query["InternalPort"] = request.InternalPort
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyForwardEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyForwardEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a high-availability virtual IP address (HAVIP).
//
// @param request - ModifyHaVipAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHaVipAttributeResponse
func (client *Client) ModifyHaVipAttributeWithContext(ctx context.Context, request *ModifyHaVipAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyHaVipAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HaVipId) {
		query["HaVipId"] = request.HaVipId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHaVipAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHaVipAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the image attributes.
//
// @param request - ModifyImageAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyImageAttributeResponse
func (client *Client) ModifyImageAttributeWithContext(ctx context.Context, request *ModifyImageAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyImageAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyImageAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyImageAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shares or unshares an image.
//
// @param request - ModifyImageSharePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyImageSharePermissionResponse
func (client *Client) ModifyImageSharePermissionWithContext(ctx context.Context, request *ModifyImageSharePermissionRequest, runtime *dara.RuntimeOptions) (_result *ModifyImageSharePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddAccounts) {
		query["AddAccounts"] = request.AddAccounts
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.RemoveAccounts) {
		query["RemoveAccounts"] = request.RemoveAccounts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyImageSharePermission"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyImageSharePermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the attributes of an instance, such as the name and the password.
//
// Description:
//
//	  If an instance is in the Starting state, you cannot reset the password of the instance.
//
//		- When the instance is in the Running state, you cannot change the password of the instance.
//
//		- After resetting the password, you must Restart the instance in the ECS console or call the RebootInstance operation to validate the modifications. The restart operation within the instance does not validate the modifications.
//
// @param request - ModifyInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttributeWithContext(ctx context.Context, request *ModifyInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures auto-renewal for instances.
//
// @param request - ModifyInstanceAutoRenewAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAutoRenewAttributeResponse
func (client *Client) ModifyInstanceAutoRenewAttributeWithContext(ctx context.Context, request *ModifyInstanceAutoRenewAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAutoRenewAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RenewalStatus) {
		query["RenewalStatus"] = request.RenewalStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAutoRenewAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAutoRenewAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the boot configuration of a heterogeneous PC Farm bare metal instance.
//
// Description:
//
//	  If an instance is in the Starting state, you cannot reset the password of the instance.
//
//		- If the instance is in the Running state, you cannot change the password of the instance.
//
//		- After resetting the password, you must restart the instance in the ENS console or call the RebootInstance operation to apply the change. The restart operation within the instance does not apply the change.
//
// @param request - ModifyInstanceBootConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceBootConfigurationResponse
func (client *Client) ModifyInstanceBootConfigurationWithContext(ctx context.Context, request *ModifyInstanceBootConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceBootConfigurationResponse, _err error) {
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
		Action:      dara.String("ModifyInstanceBootConfiguration"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceBootConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of Edge Node Service (ENS) instances. You can switch between the pay-as-you-go and subscription billing methods for instances. You can also change the billing method for disks that you created with pay-as-you-go instances to subscription.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and pricing of ENS.
//
// The instances must be in the Running or Stopped state, and you have no overdue payments for them.
//
// @param tmpReq - ModifyInstanceChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceChargeTypeResponse
func (client *Client) ModifyInstanceChargeTypeWithContext(ctx context.Context, tmpReq *ModifyInstanceChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceChargeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyInstanceChargeTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.IncludeDataDisks) {
		query["IncludeDataDisks"] = request.IncludeDataDisks
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceChargeType"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceChargeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about an Edge Load Balancer (ELB) instance.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - ModifyLoadBalancerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLoadBalancerAttributeResponse
func (client *Client) ModifyLoadBalancerAttributeWithContext(ctx context.Context, request *ModifyLoadBalancerAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyLoadBalancerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerName) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLoadBalancerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLoadBalancerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the network information.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 5 times per second per user.
//
// @param request - ModifyNetworkAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNetworkAttributeResponse
func (client *Client) ModifyNetworkAttributeWithContext(ctx context.Context, request *ModifyNetworkAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NetworkId) {
		query["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.NetworkName) {
		query["NetworkName"] = request.NetworkName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNetworkAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNetworkAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an elastic network interface (ENI), such as its name and description.
//
// @param request - ModifyNetworkInterfaceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNetworkInterfaceAttributeResponse
func (client *Client) ModifyNetworkInterfaceAttributeWithContext(ctx context.Context, request *ModifyNetworkInterfaceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkInterfaceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.NetworkInterfaceName) {
		query["NetworkInterfaceName"] = request.NetworkInterfaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNetworkInterfaceAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNetworkInterfaceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades the instance type of a subscription Edge Node Service (ENS) instance. The new instance type takes effect for the remaining lifecycle of the instance.
//
// @param request - ModifyPrepayInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPrepayInstanceSpecResponse
func (client *Client) ModifyPrepayInstanceSpecWithContext(ctx context.Context, request *ModifyPrepayInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyPrepayInstanceSpecResponse, _err error) {
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

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPrepayInstanceSpec"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPrepayInstanceSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a security group.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 5 times per second per user.
//
// @param request - ModifySecurityGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityGroupAttributeResponse
func (client *Client) ModifySecurityGroupAttributeWithContext(ctx context.Context, request *ModifySecurityGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityGroupName) {
		query["SecurityGroupName"] = request.SecurityGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityGroupAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityGroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a snapshot.
//
// @param request - ModifySnapshotAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySnapshotAttributeResponse
func (client *Client) ModifySnapshotAttributeWithContext(ctx context.Context, request *ModifySnapshotAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifySnapshotAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySnapshotAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySnapshotAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改snat规则
//
// @param request - ModifySnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySnatEntryResponse
func (client *Client) ModifySnatEntryWithContext(ctx context.Context, request *ModifySnatEntryRequest, runtime *dara.RuntimeOptions) (_result *ModifySnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EipAffinity) {
		query["EipAffinity"] = request.EipAffinity
	}

	if !dara.IsNil(request.IspAffinity) {
		query["IspAffinity"] = request.IspAffinity
	}

	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	if !dara.IsNil(request.SnatEntryName) {
		query["SnatEntryName"] = request.SnatEntryName
	}

	if !dara.IsNil(request.SnatIp) {
		query["SnatIp"] = request.SnatIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySnatEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about a vSwitch.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 5 times per second per user.
//
// @param request - ModifyVSwitchAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVSwitchAttributeResponse
func (client *Client) ModifyVSwitchAttributeWithContext(ctx context.Context, request *ModifyVSwitchAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyVSwitchAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VSwitchName) {
		query["VSwitchName"] = request.VSwitchName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVSwitchAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVSwitchAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys the SDG that has been attached to instances on the corresponding Android in Container (AIC) instance.
//
// @param tmpReq - MountInstanceSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MountInstanceSDGResponse
func (client *Client) MountInstanceSDGWithContext(ctx context.Context, tmpReq *MountInstanceSDGRequest, runtime *dara.RuntimeOptions) (_result *MountInstanceSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MountInstanceSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.SDGId) {
		query["SDGId"] = request.SDGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MountInstanceSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MountInstanceSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Preloads a shared data group (SDG).
//
// @param tmpReq - PreloadRegionSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadRegionSDGResponse
func (client *Client) PreloadRegionSDGWithContext(ctx context.Context, tmpReq *PreloadRegionSDGRequest, runtime *dara.RuntimeOptions) (_result *PreloadRegionSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PreloadRegionSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestinationRegionIds) {
		request.DestinationRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationRegionIds, dara.String("DestinationRegionIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Namespaces) {
		request.NamespacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Namespaces, dara.String("Namespaces"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationRegionIdsShrink) {
		query["DestinationRegionIds"] = request.DestinationRegionIdsShrink
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.NamespacesShrink) {
		query["Namespaces"] = request.NamespacesShrink
	}

	if !dara.IsNil(request.RedundantNum) {
		query["RedundantNum"] = request.RedundantNum
	}

	if !dara.IsNil(request.SDGId) {
		query["SDGId"] = request.SDGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreloadRegionSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreloadRegionSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prepares the upload and obtains the location of the bucket.
//
// @param request - PrepareUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrepareUploadResponse
func (client *Client) PrepareUploadWithContext(ctx context.Context, request *PrepareUploadRequest, runtime *dara.RuntimeOptions) (_result *PrepareUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrepareUpload"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PrepareUploadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pushes the business or service data of an application to file servers.
//
// @param request - PushApplicationDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushApplicationDataResponse
func (client *Client) PushApplicationDataWithContext(ctx context.Context, request *PushApplicationDataRequest, runtime *dara.RuntimeOptions) (_result *PushApplicationDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.PushStrategy) {
		query["PushStrategy"] = request.PushStrategy
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushApplicationData"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushApplicationDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Edge Object Storage (EOS) bucket.
//
// @param request - PutBucketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutBucketResponse
func (client *Client) PutBucketWithContext(ctx context.Context, request *PutBucketRequest, runtime *dara.RuntimeOptions) (_result *PutBucketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketAcl) {
		query["BucketAcl"] = request.BucketAcl
	}

	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DispatchScope) {
		query["DispatchScope"] = request.DispatchScope
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.LogicalBucketType) {
		query["LogicalBucketType"] = request.LogicalBucketType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutBucket"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutBucketResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the access control list (ACL) of a bucket.
//
// @param request - PutBucketAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutBucketAclResponse
func (client *Client) PutBucketAclWithContext(ctx context.Context, request *PutBucketAclRequest, runtime *dara.RuntimeOptions) (_result *PutBucketAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketAcl) {
		query["BucketAcl"] = request.BucketAcl
	}

	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutBucketAcl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutBucketAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures lifecycle rules for objects.
//
// Description:
//
//	  You can configure up to 1,000 rules.
//
//		- If an object meets multiple rules, the rule that has the earliest expiration time prevails.
//
// @param request - PutBucketLifecycleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutBucketLifecycleResponse
func (client *Client) PutBucketLifecycleWithContext(ctx context.Context, request *PutBucketLifecycleRequest, runtime *dara.RuntimeOptions) (_result *PutBucketLifecycleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowSameActionOverlap) {
		query["AllowSameActionOverlap"] = request.AllowSameActionOverlap
	}

	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.CreatedBeforeDate) {
		query["CreatedBeforeDate"] = request.CreatedBeforeDate
	}

	if !dara.IsNil(request.ExpirationDays) {
		query["ExpirationDays"] = request.ExpirationDays
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutBucketLifecycle"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutBucketLifecycleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initializes a disk.
//
// @param request - ReInitDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReInitDiskResponse
func (client *Client) ReInitDiskWithContext(ctx context.Context, request *ReInitDiskRequest, runtime *dara.RuntimeOptions) (_result *ReInitDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReInitDisk"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReInitDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an Android in Container (AIC) instance.
//
// @param tmpReq - RebootAICInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootAICInstanceResponse
func (client *Client) RebootAICInstanceWithContext(ctx context.Context, tmpReq *RebootAICInstanceRequest, runtime *dara.RuntimeOptions) (_result *RebootAICInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RebootAICInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootAICInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootAICInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reboots an Android in Container (AIC) server.
//
// @param request - RebootARMServerInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootARMServerInstanceResponse
func (client *Client) RebootARMServerInstanceWithContext(ctx context.Context, request *RebootARMServerInstanceRequest, runtime *dara.RuntimeOptions) (_result *RebootARMServerInstanceResponse, _err error) {
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
		Action:      dara.String("RebootARMServerInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootARMServerInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reboots an instance.
//
// Description:
//
//	  Only instances that are in the Running state can be restarted.
//
//		- If the operation is successful, the status of the instance becomes Starting.
//
// @param request - RebootInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootInstanceResponse
func (client *Client) RebootInstanceWithContext(ctx context.Context, request *RebootInstanceRequest, runtime *dara.RuntimeOptions) (_result *RebootInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceStop) {
		query["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 选择批量实例进行重启操作
//
// @param tmpReq - RebootInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootInstancesResponse
func (client *Client) RebootInstancesWithContext(ctx context.Context, tmpReq *RebootInstancesRequest, runtime *dara.RuntimeOptions) (_result *RebootInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RebootInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Recovers an Android in Container (AIC) instance on the server.
//
// @param request - RecoverAICInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverAICInstanceResponse
func (client *Client) RecoverAICInstanceWithContext(ctx context.Context, request *RecoverAICInstanceRequest, runtime *dara.RuntimeOptions) (_result *RecoverAICInstanceResponse, _err error) {
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
		Action:      dara.String("RecoverAICInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoverAICInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets an instance based on specific parameters.
//
// @param request - ReinitInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReinitInstanceResponse
func (client *Client) ReinitInstanceWithContext(ctx context.Context, request *ReinitInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReinitInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReinitInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReinitInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 选择批量实例进行重置操作
//
// @param tmpReq - ReinitInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReinitInstancesResponse
func (client *Client) ReinitInstancesWithContext(ctx context.Context, tmpReq *ReinitInstancesRequest, runtime *dara.RuntimeOptions) (_result *ReinitInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReinitInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReinitInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReinitInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases an Android in Container (AIC) instance from the server.
//
// @param request - ReleaseAICInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseAICInstanceResponse
func (client *Client) ReleaseAICInstanceWithContext(ctx context.Context, request *ReleaseAICInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseAICInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ServerId) {
		query["ServerId"] = request.ServerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseAICInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseAICInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases an ARM server.
//
// @param request - ReleaseARMServerInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseARMServerInstanceResponse
func (client *Client) ReleaseARMServerInstanceWithContext(ctx context.Context, request *ReleaseARMServerInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseARMServerInstanceResponse, _err error) {
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
		Action:      dara.String("ReleaseARMServerInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseARMServerInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases an instance. You can call this operation to release only Elastic IP Addresses (EIPs), Edge Load Balancer (ELB) instances, and cloud disk-based instances. We recommend that you call service-specific operations to release or unsubscribe from instances.
//
// Description:
//
//	  You can call this operation up to 10,000 times per second per account.
//
//		- The maximum number of times that each user can call this operation per second is 50.
//
// @param request - ReleaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithContext(ctx context.Context, request *ReleaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
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
		Action:      dara.String("ReleaseInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a pay-as-you-go instance.
//
// @param request - ReleasePostPaidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleasePostPaidInstanceResponse
func (client *Client) ReleasePostPaidInstanceWithContext(ctx context.Context, request *ReleasePostPaidInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleasePostPaidInstanceResponse, _err error) {
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
		Action:      dara.String("ReleasePostPaidInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleasePostPaidInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call ReleasePrePaidInstance to delete a subscription instance.
//
// @param request - ReleasePrePaidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleasePrePaidInstanceResponse
func (client *Client) ReleasePrePaidInstanceWithContext(ctx context.Context, request *ReleasePrePaidInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleasePrePaidInstanceResponse, _err error) {
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
		Action:      dara.String("ReleasePrePaidInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleasePrePaidInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes backend servers.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param tmpReq - RemoveBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveBackendServersResponse
func (client *Client) RemoveBackendServersWithContext(ctx context.Context, tmpReq *RemoveBackendServersRequest, runtime *dara.RuntimeOptions) (_result *RemoveBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveBackendServersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BackendServers) {
		request.BackendServersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackendServers, dara.String("BackendServers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServersShrink) {
		query["BackendServers"] = request.BackendServersShrink
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveBackendServers"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveBackendServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a shared data group (SDG) that is attached to the compute instance.
//
// @param tmpReq - RemoveInstanceSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInstanceSDGResponse
func (client *Client) RemoveInstanceSDGWithContext(ctx context.Context, tmpReq *RemoveInstanceSDGRequest, runtime *dara.RuntimeOptions) (_result *RemoveInstanceSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveInstanceSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveInstanceSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveInstanceSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes IP addresses from an edge private network (EPN) instance.
//
// @param request - RemovePublicIpsFromEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePublicIpsFromEpnInstanceResponse
func (client *Client) RemovePublicIpsFromEpnInstanceWithContext(ctx context.Context, request *RemovePublicIpsFromEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *RemovePublicIpsFromEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	if !dara.IsNil(request.InstanceInfos) {
		query["InstanceInfos"] = request.InstanceInfos
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePublicIpsFromEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePublicIpsFromEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a deployed shared data group (SDG) and restore local storage.
//
// @param tmpReq - RemoveSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSDGResponse
func (client *Client) RemoveSDGWithContext(ctx context.Context, tmpReq *RemoveSDGRequest, runtime *dara.RuntimeOptions) (_result *RemoveSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes all versions of SDG and restores the mount to the local disk.
//
// @param tmpReq - RemoveSDGsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSDGsResponse
func (client *Client) RemoveSDGsWithContext(ctx context.Context, tmpReq *RemoveSDGsRequest, runtime *dara.RuntimeOptions) (_result *RemoveSDGsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveSDGsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SdgIds) {
		request.SdgIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SdgIds, dara.String("SdgIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.SdgIdsShrink) {
		query["SdgIds"] = request.SdgIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSDGs"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSDGsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the networking information. This operation is applicable only for instances that reside in the internal network.
//
// @param request - RemoveVSwitchesFromEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveVSwitchesFromEpnInstanceResponse
func (client *Client) RemoveVSwitchesFromEpnInstanceWithContext(ctx context.Context, request *RemoveVSwitchesFromEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *RemoveVSwitchesFromEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	if !dara.IsNil(request.VSwitchesInfo) {
		query["VSwitchesInfo"] = request.VSwitchesInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveVSwitchesFromEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveVSwitchesFromEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a subscription Android in Container (AIC) instance.
//
// @param request - RenewARMServerInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewARMServerInstanceResponse
func (client *Client) RenewARMServerInstanceWithContext(ctx context.Context, request *RenewARMServerInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewARMServerInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewARMServerInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewARMServerInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a subscription instance.
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
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
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales resources in an asynchronous manner and deploys or releases the container.
//
// @param request - RescaleApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RescaleApplicationResponse
func (client *Client) RescaleApplicationWithContext(ctx context.Context, request *RescaleApplicationRequest, runtime *dara.RuntimeOptions) (_result *RescaleApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.RescaleLevel) {
		query["RescaleLevel"] = request.RescaleLevel
	}

	if !dara.IsNil(request.RescaleType) {
		query["RescaleType"] = request.RescaleType
	}

	if !dara.IsNil(request.ResourceSelector) {
		query["ResourceSelector"] = request.ResourceSelector
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.ToAppVersion) {
		query["ToAppVersion"] = request.ToAppVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RescaleApplication"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RescaleApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales out a bare metal device.
//
// @param request - RescaleDeviceServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RescaleDeviceServiceResponse
func (client *Client) RescaleDeviceServiceWithContext(ctx context.Context, request *RescaleDeviceServiceRequest, runtime *dara.RuntimeOptions) (_result *RescaleDeviceServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.IpType) {
		query["IpType"] = request.IpType
	}

	if !dara.IsNil(request.RescaleLevel) {
		query["RescaleLevel"] = request.RescaleLevel
	}

	if !dara.IsNil(request.RescaleType) {
		query["RescaleType"] = request.RescaleType
	}

	if !dara.IsNil(request.ResourceSpec) {
		query["ResourceSpec"] = request.ResourceSpec
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceInfo) {
		body["ResourceInfo"] = request.ResourceInfo
	}

	if !dara.IsNil(request.ResourceSelector) {
		body["ResourceSelector"] = request.ResourceSelector
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RescaleDeviceService"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RescaleDeviceServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets an Android in Container (AIC) instance.
//
// @param request - ResetAICInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAICInstanceResponse
func (client *Client) ResetAICInstanceWithContext(ctx context.Context, request *ResetAICInstanceRequest, runtime *dara.RuntimeOptions) (_result *ResetAICInstanceResponse, _err error) {
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
		Action:      dara.String("ResetAICInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAICInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a disk by using a snapshot.
//
// Description:
//
// When you call this operation, take note of the following items:
//
//   - The disk must be in the In Use (In_Use) or Unattached (Available) state.
//
//   - The instance to which the disk is attached must be in the Stopped (Stopped) state. You can call the [StopInstance](~~StopInstance~~) operation to stop an instance.
//
//   - The specified snapshot must be created from the disk specified by the DiskId parameter.
//
//   - If the response contains `{"OperationLocks": {"LockReason" : "security"}}` when you query information about an ENS instance by calling the [DescribeInstances](~~DescribeInstances~~) operation, the instance is locked for security reasons and no operations are allowed on the instance.
//
// @param request - ResetDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDiskResponse
func (client *Client) ResetDiskWithContext(ctx context.Context, request *ResetDiskRequest, runtime *dara.RuntimeOptions) (_result *ResetDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetDisk"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resizes a pay-as-you-go disk that you purchase.
//
// @param request - ResizeDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeDiskResponse
func (client *Client) ResizeDiskWithContext(ctx context.Context, request *ResizeDiskRequest, runtime *dara.RuntimeOptions) (_result *ResizeDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.NewSize) {
		query["NewSize"] = request.NewSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResizeDisk"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an inbound security group rule. After the rule is deleted, the access control implemented by the rule is removed.
//
// Description:
//
//	  In the security group-related API documents, inbound traffic refers to the traffic sent by the source and received by the destination.
//
//		- You can determine an inbound security group rule by specifying one of the following groups of parameters. You cannot determine a security group rule by specifying only one parameter.
//
//		- You can specify one or more of the following parameters to remove access control for a CIDR block: IpProtocol, PortRange, Policy, and SourceCidrIp.
//
// @param request - RevokeSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeSecurityGroupResponse
func (client *Client) RevokeSecurityGroupWithContext(ctx context.Context, request *RevokeSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *RevokeSecurityGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeSecurityGroup"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an outbound security group rule. After the rule is deleted, the access control implemented by the rule is removed.
//
// Description:
//
// >  In the security group-related API documents, outbound traffic refers to the traffic sent by the source and received by the destination.
//
// @param request - RevokeSecurityGroupEgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeSecurityGroupEgressResponse
func (client *Client) RevokeSecurityGroupEgressWithContext(ctx context.Context, request *RevokeSecurityGroupEgressRequest, runtime *dara.RuntimeOptions) (_result *RevokeSecurityGroupEgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestCidrIp) {
		query["DestCidrIp"] = request.DestCidrIp
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeSecurityGroupEgress"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeSecurityGroupEgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back the container version of a specific application.
//
// @param request - RollbackApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackApplicationResponse
func (client *Client) RollbackApplicationWithContext(ctx context.Context, request *RollbackApplicationRequest, runtime *dara.RuntimeOptions) (_result *RollbackApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.FromAppVersion) {
		query["FromAppVersion"] = request.FromAppVersion
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.ToAppVersion) {
		query["ToAppVersion"] = request.ToAppVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackApplication"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go or subscription ENS instance.
//
// @param tmpReq - RunInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunInstancesResponse
func (client *Client) RunInstancesWithContext(ctx context.Context, tmpReq *RunInstancesRequest, runtime *dara.RuntimeOptions) (_result *RunInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataDisk) {
		request.DataDiskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataDisk, dara.String("DataDisk"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SystemDisk) {
		request.SystemDiskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SystemDisk, dara.String("SystemDisk"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AutoReleaseTime) {
		query["AutoReleaseTime"] = request.AutoReleaseTime
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BillingCycle) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !dara.IsNil(request.Carrier) {
		query["Carrier"] = request.Carrier
	}

	if !dara.IsNil(request.DataDiskShrink) {
		query["DataDisk"] = request.DataDiskShrink
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.EnsRegionId) {
		query["EnsRegionId"] = request.EnsRegionId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceChargeStrategy) {
		query["InstanceChargeStrategy"] = request.InstanceChargeStrategy
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.InternetMaxBandwidthOut) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !dara.IsNil(request.IpType) {
		query["IpType"] = request.IpType
	}

	if !dara.IsNil(request.Ipv6AddressCount) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.LaunchTemplateId) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !dara.IsNil(request.LaunchTemplateName) {
		query["LaunchTemplateName"] = request.LaunchTemplateName
	}

	if !dara.IsNil(request.LaunchTemplateVersion) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !dara.IsNil(request.NetDistrictCode) {
		query["NetDistrictCode"] = request.NetDistrictCode
	}

	if !dara.IsNil(request.NetWorkId) {
		query["NetWorkId"] = request.NetWorkId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordInherit) {
		query["PasswordInherit"] = request.PasswordInherit
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.PublicIpIdentification) {
		query["PublicIpIdentification"] = request.PublicIpIdentification
	}

	if !dara.IsNil(request.ScheduleAreaLevel) {
		query["ScheduleAreaLevel"] = request.ScheduleAreaLevel
	}

	if !dara.IsNil(request.SchedulingPriceStrategy) {
		query["SchedulingPriceStrategy"] = request.SchedulingPriceStrategy
	}

	if !dara.IsNil(request.SchedulingStrategy) {
		query["SchedulingStrategy"] = request.SchedulingStrategy
	}

	if !dara.IsNil(request.SecurityId) {
		query["SecurityId"] = request.SecurityId
	}

	if !dara.IsNil(request.SpotDuration) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.SystemDiskShrink) {
		query["SystemDisk"] = request.SystemDiskShrink
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UniqueSuffix) {
		query["UniqueSuffix"] = request.UniqueSuffix
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Schedules the nearest idle resources including instances and pods for your device based on the user ID and IP address and initializes the virtual environment.
//
// @param request - RunServiceScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunServiceScheduleResponse
func (client *Client) RunServiceScheduleWithContext(ctx context.Context, request *RunServiceScheduleRequest, runtime *dara.RuntimeOptions) (_result *RunServiceScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.Directorys) {
		query["Directorys"] = request.Directorys
	}

	if !dara.IsNil(request.PodConfigName) {
		query["PodConfigName"] = request.PodConfigName
	}

	if !dara.IsNil(request.PreLockedTimeout) {
		query["PreLockedTimeout"] = request.PreLockedTimeout
	}

	if !dara.IsNil(request.ScheduleStrategy) {
		query["ScheduleStrategy"] = request.ScheduleStrategy
	}

	if !dara.IsNil(request.ServiceAction) {
		query["ServiceAction"] = request.ServiceAction
	}

	if !dara.IsNil(request.ServiceCommands) {
		query["ServiceCommands"] = request.ServiceCommands
	}

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunServiceSchedule"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunServiceScheduleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves the disk of a specific device as a shared data group (SDG).
//
// @param request - SaveSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSDGResponse
func (client *Client) SaveSDGWithContext(ctx context.Context, request *SaveSDGRequest, runtime *dara.RuntimeOptions) (_result *SaveSDGResponse, _err error) {
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
		Action:      dara.String("SaveSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the weights of backend servers.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param tmpReq - SetBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetBackendServersResponse
func (client *Client) SetBackendServersWithContext(ctx context.Context, tmpReq *SetBackendServersRequest, runtime *dara.RuntimeOptions) (_result *SetBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetBackendServersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BackendServers) {
		request.BackendServersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackendServers, dara.String("BackendServers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServersShrink) {
		query["BackendServers"] = request.BackendServersShrink
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetBackendServers"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetBackendServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an HTTP listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - SetLoadBalancerHTTPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerHTTPListenerAttributeResponse
func (client *Client) SetLoadBalancerHTTPListenerAttributeWithContext(ctx context.Context, request *SetLoadBalancerHTTPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerHTTPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.XForwardedFor) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerHTTPListenerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an HTTPS listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - SetLoadBalancerHTTPSListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerHTTPSListenerAttributeResponse
func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithContext(ctx context.Context, request *SetLoadBalancerHTTPSListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerHTTPSListenerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of an Edge Load Balancer (ELB) instance.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - SetLoadBalancerStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerStatusResponse
func (client *Client) SetLoadBalancerStatusWithContext(ctx context.Context, request *SetLoadBalancerStatusRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerStatus) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerStatus"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a TCP listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - SetLoadBalancerTCPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerTCPListenerAttributeResponse
func (client *Client) SetLoadBalancerTCPListenerAttributeWithContext(ctx context.Context, request *SetLoadBalancerTCPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerTCPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EipTransmit) {
		query["EipTransmit"] = request.EipTransmit
	}

	if !dara.IsNil(request.EstablishedTimeout) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckConnectTimeout) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckType) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.PersistenceTimeout) {
		query["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerTCPListenerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a UDP listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - SetLoadBalancerUDPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerUDPListenerAttributeResponse
func (client *Client) SetLoadBalancerUDPListenerAttributeWithContext(ctx context.Context, request *SetLoadBalancerUDPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerUDPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EipTransmit) {
		query["EipTransmit"] = request.EipTransmit
	}

	if !dara.IsNil(request.EstablishedTimeout) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckConnectTimeout) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !dara.IsNil(request.HealthCheckExp) {
		query["HealthCheckExp"] = request.HealthCheckExp
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckReq) {
		query["HealthCheckReq"] = request.HealthCheckReq
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerUDPListenerAttribute"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 共享AIC镜像
//
// @param tmpReq - ShareAICImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShareAICImageResponse
func (client *Client) ShareAICImageWithContext(ctx context.Context, tmpReq *ShareAICImageRequest, runtime *dara.RuntimeOptions) (_result *ShareAICImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ShareAICImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Users) {
		request.UsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Users, dara.String("Users"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.UsersShrink) {
		query["Users"] = request.UsersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ShareAICImage"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ShareAICImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an edge network instance.
//
// @param request - StartEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartEpnInstanceResponse
func (client *Client) StartEpnInstanceWithContext(ctx context.Context, request *StartEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an instance.
//
// Description:
//
//	  You can call the operation only when the instance is in the Stopped state.
//
//		- If the operation is successful, the status of the instance becomes Starting.
//
// @param request - StartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithContext(ctx context.Context, request *StartInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
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
		Action:      dara.String("StartInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 选择批量实例进行启动操作
//
// @param tmpReq - StartInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstancesResponse
func (client *Client) StartInstancesWithContext(ctx context.Context, tmpReq *StartInstancesRequest, runtime *dara.RuntimeOptions) (_result *StartInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - StartLoadBalancerListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLoadBalancerListenerResponse
func (client *Client) StartLoadBalancerListenerWithContext(ctx context.Context, request *StartLoadBalancerListenerRequest, runtime *dara.RuntimeOptions) (_result *StartLoadBalancerListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLoadBalancerListener"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLoadBalancerListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the elastic IP address (EIP) specified in a source network address translation (SNAT) entry.
//
// @param request - StartSnatIpForSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSnatIpForSnatEntryResponse
func (client *Client) StartSnatIpForSnatEntryWithContext(ctx context.Context, request *StartSnatIpForSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *StartSnatIpForSnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	if !dara.IsNil(request.SnatIp) {
		query["SnatIp"] = request.SnatIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSnatIpForSnatEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSnatIpForSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an EPN instance.
//
// @param request - StopEpnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopEpnInstanceResponse
func (client *Client) StopEpnInstanceWithContext(ctx context.Context, request *StopEpnInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopEpnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EPNInstanceId) {
		query["EPNInstanceId"] = request.EPNInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopEpnInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopEpnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an instance.
//
// Description:
//
//	  You can call this operation to stop instances that are only in the Running state.
//
//		- If the call is successful, the state of the instance becomes Stopping.
//
//		- Once the instance is stopped, the state of the instance becomes Stopped.
//
//		- Force stop is supported, which is equivalent to power-off. Data that is not written to disks on the instance may be lost.
//
// @param request - StopInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithContext(ctx context.Context, request *StopInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceStop) {
		query["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInstance"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 选择批量实例进行停止操作
//
// @param tmpReq - StopInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstancesResponse
func (client *Client) StopInstancesWithContext(ctx context.Context, tmpReq *StopInstancesRequest, runtime *dara.RuntimeOptions) (_result *StopInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInstances"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a listener.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can call this operation up to 10 times per second per user.
//
// @param request - StopLoadBalancerListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLoadBalancerListenerResponse
func (client *Client) StopLoadBalancerListenerWithContext(ctx context.Context, request *StopLoadBalancerListenerRequest, runtime *dara.RuntimeOptions) (_result *StopLoadBalancerListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLoadBalancerListener"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLoadBalancerListenerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an elastic IP address (EIP) in a source network address translation (SNAT) entry.
//
// @param request - StopSnatIpForSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSnatIpForSnatEntryResponse
func (client *Client) StopSnatIpForSnatEntryWithContext(ctx context.Context, request *StopSnatIpForSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *StopSnatIpForSnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnatEntryId) {
		query["SnatEntryId"] = request.SnatEntryId
	}

	if !dara.IsNil(request.SnatIp) {
		query["SnatIp"] = request.SnatIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopSnatIpForSnatEntry"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopSnatIpForSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and adds tags to specific Edge Node Service (ENS) resources.
//
// Description:
//
// Before you add tags to a resource, Alibaba Cloud checks the number of existing tags on the resource. If the number exceeds the upper limit, an error message is returned. Only instance resources, such as virtual machines and bare machines, are supported.
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
		Version:     dara.String("2017-11-10"),
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
// Disassociates an elastic IP address (EIP) from an instance.
//
// @param request - UnAssociateEnsEipAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnAssociateEnsEipAddressResponse
func (client *Client) UnAssociateEnsEipAddressWithContext(ctx context.Context, request *UnAssociateEnsEipAddressRequest, runtime *dara.RuntimeOptions) (_result *UnAssociateEnsEipAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationId) {
		query["AllocationId"] = request.AllocationId
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnAssociateEnsEipAddress"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnAssociateEnsEipAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unassigns secondary private IP addresses from an elastic network interface (ENI).
//
// @param request - UnassignPrivateIpAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassignPrivateIpAddressesResponse
func (client *Client) UnassignPrivateIpAddressesWithContext(ctx context.Context, request *UnassignPrivateIpAddressesRequest, runtime *dara.RuntimeOptions) (_result *UnassignPrivateIpAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnassignPrivateIpAddresses"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnassignPrivateIpAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a high-availability virtual IP address (HAVIP) from an Edge Node Service (ENS) instance or Elastic Network Interface (ENI).
//
// @param request - UnassociateHaVipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassociateHaVipResponse
func (client *Client) UnassociateHaVipWithContext(ctx context.Context, request *UnassociateHaVipRequest, runtime *dara.RuntimeOptions) (_result *UnassociateHaVipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HaVipId) {
		query["HaVipId"] = request.HaVipId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnassociateHaVip"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnassociateHaVipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a network access control list (ACL) from a network.
//
// @param request - UnassociateNetworkAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassociateNetworkAclResponse
func (client *Client) UnassociateNetworkAclWithContext(ctx context.Context, request *UnassociateNetworkAclRequest, runtime *dara.RuntimeOptions) (_result *UnassociateNetworkAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkAclId) {
		query["NetworkAclId"] = request.NetworkAclId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnassociateNetworkAcl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnassociateNetworkAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes preloaded data.
//
// @param tmpReq - UnloadRegionSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnloadRegionSDGResponse
func (client *Client) UnloadRegionSDGWithContext(ctx context.Context, tmpReq *UnloadRegionSDGRequest, runtime *dara.RuntimeOptions) (_result *UnloadRegionSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UnloadRegionSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestinationRegionIds) {
		request.DestinationRegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationRegionIds, dara.String("DestinationRegionIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Namespaces) {
		request.NamespacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Namespaces, dara.String("Namespaces"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationRegionIdsShrink) {
		query["DestinationRegionIds"] = request.DestinationRegionIdsShrink
	}

	if !dara.IsNil(request.NamespacesShrink) {
		query["Namespaces"] = request.NamespacesShrink
	}

	if !dara.IsNil(request.SDGId) {
		query["SDGId"] = request.SDGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnloadRegionSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnloadRegionSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unmounts a shared data group (SDG) from instances.
//
// @param tmpReq - UnmountInstanceSDGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnmountInstanceSDGResponse
func (client *Client) UnmountInstanceSDGWithContext(ctx context.Context, tmpReq *UnmountInstanceSDGRequest, runtime *dara.RuntimeOptions) (_result *UnmountInstanceSDGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UnmountInstanceSDGShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.SDGId) {
		query["SDGId"] = request.SDGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnmountInstanceSDG"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnmountInstanceSDGResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
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
		Version:     dara.String("2017-11-10"),
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
// 修改售卖约束
//
// @param tmpReq - UpdateEnsSaleControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnsSaleControlResponse
func (client *Client) UpdateEnsSaleControlWithContext(ctx context.Context, tmpReq *UpdateEnsSaleControlRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnsSaleControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateEnsSaleControlShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SaleControls) {
		request.SaleControlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SaleControls, dara.String("SaleControls"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUidAccount) {
		query["AliUidAccount"] = request.AliUidAccount
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CustomAccount) {
		query["CustomAccount"] = request.CustomAccount
	}

	if !dara.IsNil(request.SaleControlsShrink) {
		query["SaleControls"] = request.SaleControlsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnsSaleControl"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnsSaleControlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the image of an Android in Container (AIC) instance.
//
// @param tmpReq - UpgradeAICInstanceImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAICInstanceImageResponse
func (client *Client) UpgradeAICInstanceImageWithContext(ctx context.Context, tmpReq *UpgradeAICInstanceImageRequest, runtime *dara.RuntimeOptions) (_result *UpgradeAICInstanceImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpgradeAICInstanceImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServerIds) {
		request.ServerIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerIds, dara.String("ServerIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeAICInstanceImage"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeAICInstanceImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the container in an asynchronous manner. You can configure multiple canary release policies.
//
// @param request - UpgradeApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeApplicationResponse
func (client *Client) UpgradeApplicationWithContext(ctx context.Context, request *UpgradeApplicationRequest, runtime *dara.RuntimeOptions) (_result *UpgradeApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeApplication"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传公钥
//
// @param request - UploadAICPublicKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadAICPublicKeyResponse
func (client *Client) UploadAICPublicKeyWithContext(ctx context.Context, request *UploadAICPublicKeyRequest, runtime *dara.RuntimeOptions) (_result *UploadAICPublicKeyResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.KeyGroup) {
		query["KeyGroup"] = request.KeyGroup
	}

	if !dara.IsNil(request.KeyName) {
		query["KeyName"] = request.KeyName
	}

	if !dara.IsNil(request.KeyType) {
		query["KeyType"] = request.KeyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadAICPublicKey"),
		Version:     dara.String("2017-11-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadAICPublicKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
