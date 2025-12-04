// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Apply for a secondary private IP address for the current Lingjun Elastic Network Interface. You can automatically assign a secondary private IP address.
//
// Description:
//
// Apply for a secondary private IP address for the specified Lingjun Elastic Network Interface.
//
//   - If the PrivateIp field is empty, a secondary private IP address is automatically assigned and the unique identifier of the IP address is returned.
//
//   - You can use the GetLeniPrivateIpAddress or ListLeniPrivateIpAddresses interface to check whether the secondary private IP address is assigned.
//
// @param request - AssignLeniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignLeniPrivateIpAddressResponse
func (client *Client) AssignLeniPrivateIpAddressWithContext(ctx context.Context, request *AssignLeniPrivateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *AssignLeniPrivateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignLeniPrivateIpAddress"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignLeniPrivateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a private secondary IP address for the current LNI. You can also call this operation to assign a secondary MAC address to the current LNI.
//
// Description:
//
// >  Apply for secondary private IP addresses
//
//   - By default, each network interface controller can apply for three secondary private IP addresses. If the quota is exceeded, contact the administrator.
//
//   - The secondary private IP address is allocated from the Lingjun subnet to which the current network interface controller belongs. The first address and the last two addresses belong to reserved addresses and do not participate in the allocation.
//
// @param request - AssignPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignPrivateIpAddressResponse
func (client *Client) AssignPrivateIpAddressWithContext(ctx context.Context, request *AssignPrivateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *AssignPrivateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssignMac) {
		body["AssignMac"] = request.AssignMac
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SkipConfig) {
		body["SkipConfig"] = request.SkipConfig
	}

	if !dara.IsNil(request.SubnetId) {
		body["SubnetId"] = request.SubnetId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignPrivateIpAddress"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignPrivateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the VPD primary network segment address is not enough to allocate, you can choose to create an additional network segment as the additional network segment of the VPD primary network segment.
//
// Description:
//
// >  **Add a CIDR block**
//
//   - The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
//
//   - The secondary IPv4 CIDR block must not overlap with the primary IPv4 CIDR block of the Lingjun CIDR block and the added secondary IPv4 CIDR block.
//
//   - You cannot use 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 as the CIDR block of Lingjun. Example: In the Lingjun CIDR block whose primary IPv4 CIDR block is 192.168.0.0/16, you cannot add the following CIDR blocks as additional IPv4 CIDR blocks. The CIDR block that is in the same range as 192.168.0.0/16. A CIDR block that is larger than 192.168.0.0/16. Example: 192.168.0.0/8. A CIDR block that is smaller than 192.168.0.0/16. Example: 192.168.0.0/24.
//
//   - By default, each tenant can create three additional CIDR blocks in each region.
//
// @param request - AssociateVpdCidrBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateVpdCidrBlockResponse
func (client *Client) AssociateVpdCidrBlockWithContext(ctx context.Context, request *AssociateVpdCidrBlockRequest, runtime *dara.RuntimeOptions) (_result *AssociateVpdCidrBlockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecondaryCidrBlock) {
		body["SecondaryCidrBlock"] = request.SecondaryCidrBlock
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateVpdCidrBlock"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateVpdCidrBlockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lingjun ENI is bound to NC.
//
// Description:
//
// This interface is an asynchronous interface. You need to use the query interface to wait for the Lingjun Elastic Network Interface to reach the available state.
//
// @param request - AttachElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachElasticNetworkInterfaceResponse
func (client *Client) AttachElasticNetworkInterfaceWithContext(ctx context.Context, request *AttachElasticNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *AttachElasticNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachElasticNetworkInterface"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachElasticNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an LENI.
//
// @param request - CreateElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateElasticNetworkInterfaceResponse
func (client *Client) CreateElasticNetworkInterfaceWithContext(ctx context.Context, request *CreateElasticNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *CreateElasticNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableJumboFrame) {
		body["EnableJumboFrame"] = request.EnableJumboFrame
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateElasticNetworkInterface"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateElasticNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a Lingjun HUB.
//
// Description:
//
// When you call this operation to create a Lingjun HUB, note that:
//
//   - Make sure that you have sufficient Lingjun HUB quota.
//
//   - This interface is an asynchronous interface. After this interface is called, the system will return the ID of a Lingjun HUB. At this time, the Lingjun HUB instance may not be created yet, and the system background creation task is still in progress. You can call the ListErs or GetEr operation to query the status of the Lingjun HUB.
//
//   - If the status of the Lingjun HUB is Executing, it indicates that it is being created.
//
//   - If the status of the Lingjun HUB is Available, the creation is successful.
//
// @param request - CreateErRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateErResponse
func (client *Client) CreateErWithContext(ctx context.Context, request *CreateErRequest, runtime *dara.RuntimeOptions) (_result *CreateErResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ErName) {
		body["ErName"] = request.ErName
	}

	if !dara.IsNil(request.MasterZoneId) {
		body["MasterZoneId"] = request.MasterZoneId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEr"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateErResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a network instance connection.
//
// Description:
//
// When you call this operation to create a network instance connection, note that:
//
//   - Make sure that you have created a Lingjun HUB instance.
//
//   - Make sure that you have sufficient quota for network instance connections.
//
//   - This operation is an asynchronous operation. After you call this operation, the system returns the ID of the network instance connection. In this case, the network instance connection may not be created yet, and the system is still creating the network instance in the background. You can query the connection status of a network instance by ListErAttachments or GetErAttachment:
//
//   - If the connection status of the network instance is Executing, the network instance is being created.
//
//   - If the connection status of the network instance is Available, the network instance is created.
//
// @param request - CreateErAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateErAttachmentResponse
func (client *Client) CreateErAttachmentWithContext(ctx context.Context, request *CreateErAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateErAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoReceiveAllRoute) {
		body["AutoReceiveAllRoute"] = request.AutoReceiveAllRoute
	}

	if !dara.IsNil(request.ErAttachmentName) {
		body["ErAttachmentName"] = request.ErAttachmentName
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceTenantId) {
		body["ResourceTenantId"] = request.ResourceTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateErAttachment"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateErAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Users can use this API to create routing policy by specifying the network instance connection under Lingjun HUB.
//
// Description:
//
// When you call this operation to create a routing policy, note that:
//
//   - Make sure that you have created a Lingjun HUB instance.
//
//   - Make sure that you have created a network instance connection.
//
//   - This operation is an asynchronous operation. After you call this operation, the system returns the ID of the routing policy. In this case, the routing policy instance may not be created yet, and the system background creation task is still in progress. You can use ListErRouteMaps or GetErRouteMap to query the status of a routing policy.
//
//   - If the status of the routing policy is Execute, the system is creating the instance.
//
//   - If the status of the routing policy is Available, the creation is successful.
//
// @param request - CreateErRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateErRouteMapResponse
func (client *Client) CreateErRouteMapWithContext(ctx context.Context, request *CreateErRouteMapRequest, runtime *dara.RuntimeOptions) (_result *CreateErRouteMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ReceptionInstanceId) {
		body["ReceptionInstanceId"] = request.ReceptionInstanceId
	}

	if !dara.IsNil(request.ReceptionInstanceOwner) {
		body["ReceptionInstanceOwner"] = request.ReceptionInstanceOwner
	}

	if !dara.IsNil(request.ReceptionInstanceType) {
		body["ReceptionInstanceType"] = request.ReceptionInstanceType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RouteMapAction) {
		body["RouteMapAction"] = request.RouteMapAction
	}

	if !dara.IsNil(request.RouteMapNum) {
		body["RouteMapNum"] = request.RouteMapNum
	}

	if !dara.IsNil(request.TransmissionInstanceId) {
		body["TransmissionInstanceId"] = request.TransmissionInstanceId
	}

	if !dara.IsNil(request.TransmissionInstanceOwner) {
		body["TransmissionInstanceOwner"] = request.TransmissionInstanceOwner
	}

	if !dara.IsNil(request.TransmissionInstanceType) {
		body["TransmissionInstanceType"] = request.TransmissionInstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateErRouteMap"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateErRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Users can use this API to create a Lingjun subnet under the Lingjun network segment.
//
// Description:
//
// When you call this operation to create a Lingjun subnet, note that:
//
//   - You have created a Lingjun CIDR block.
//
//   - Only one network segment can be specified for a Lingjun subnet.
//
//   - The network segment cannot be modified after the Lingjun subnet is created.
//
//   - Make sure that you have sufficient Lingjun subnet quota.
//
//   - This interface is an asynchronous interface. After calling this interface, the system will return the ID of a Lingjun subnet. At this time, the Lingjun network segment may not be created yet, and the system background creation task is still in progress. You can call the ListSubnets or GetSubnet operation to query the status of the CIDR block of Lingjun.
//
//   - If the status of the Lingjun subnet is Executed, it indicates that it is being created.
//
//   - If the status of the Lingjun subnet is Available, the creation is successful.
//
// @param request - CreateSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubnetResponse
func (client *Client) CreateSubnetWithContext(ctx context.Context, request *CreateSubnetRequest, runtime *dara.RuntimeOptions) (_result *CreateSubnetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cidr) {
		body["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubnetName) {
		body["SubnetName"] = request.SubnetName
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubnet"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubnetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can create a Lingjun connection to connect Lingjun network environment and Alibaba Cloud network environment.
//
// Description:
//
// When you call this operation to create a Lingjun connection, note that:
//
//   - When you specify the vccId parameter, the system will configure the purchased Lingjun connection for you. When the default vccId parameter is set, the system will automatically place an order and configure the Lingjun connection for you.
//
//   - Make sure that you have called the InitializeVcc operation to grant permissions.
//
//   - This interface is an asynchronous interface. After this interface is called, the system will return the Lingjun connection ID, but the Lingjun connection instance may not be created yet, and the system background creation task is still in progress. You can call the ListVccs or GetVcc operation to query the status of the Lingjun connection.
//
//   - If the status of the Lingjun connection is Executed, the Lingjun connection is being created.
//
//   - If the status of the Lingjun connection is Available, the Lingjun connection is created.
//
// @param request - CreateVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVccResponse
func (client *Client) CreateVccWithContext(ctx context.Context, request *CreateVccRequest, runtime *dara.RuntimeOptions) (_result *CreateVccResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessCouldService) {
		body["AccessCouldService"] = request.AccessCouldService
	}

	if !dara.IsNil(request.Bandwidth) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BgpAsn) {
		body["BgpAsn"] = request.BgpAsn
	}

	if !dara.IsNil(request.BgpCidr) {
		body["BgpCidr"] = request.BgpCidr
	}

	if !dara.IsNil(request.CenId) {
		body["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CenOwnerId) {
		body["CenOwnerId"] = request.CenOwnerId
	}

	if !dara.IsNil(request.ConnectionType) {
		body["ConnectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	if !dara.IsNil(request.VccName) {
		body["VccName"] = request.VccName
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVcc"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVccResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Users can use this API to connect Lingjun instance to the Lingjun HUB instance of the target account. After authorization, the target account can be associated with your Lingjun connection by using the authorized Lingjun HUB instance.
//
// Description:
//
// When you call this operation to create cross-account authorization for Lingjun HUB, note that:
//
//   - Make sure that the Alibaba Cloud ID and Lingjun HUB instance that you want to authorize are correct.
//
//   - If you authorize the account of the other party, the account of the other party can load your local network instance to its Lingjun HUB, and the other party\\"s network will be connected to your network. Please proceed with caution.
//
// @param request - CreateVccGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVccGrantRuleResponse
func (client *Client) CreateVccGrantRuleWithContext(ctx context.Context, request *CreateVccGrantRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateVccGrantRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.GrantTenantId) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVccGrantRule"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVccGrantRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a Lingjun connection route entry.
//
// Description:
//
// When you call this operation to create a VBR route entry, take note of the following items:
//
//   - After you call this operation, static route entries and BGP network announcements are created on the VBR to which the Lingjun connection belongs.
//
//   - This operation is an asynchronous operation. After you call this operation, the VBR static route entry may not be created yet, and the system still creates the static route entry in the background. You can query the status of VBR static route entries by ListVccRouteEntries or GetVccRouteEntry:
//
//   - If the VBR static route entry is in the Executing state, it indicates that it is being created.
//
//   - If the status of the VBR static route entry is Available, the VBR is created.
//
// @param request - CreateVccRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVccRouteEntryResponse
func (client *Client) CreateVccRouteEntryWithContext(ctx context.Context, request *CreateVccRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateVccRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVccRouteEntry"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVccRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a private Lingjun CIDR block. This CIDR block has an independent network environment.
//
// Description:
//
// When you call this operation to create a CIDR block for Lingjun, take note of the following:
//
//   - A Lingjun network segment can specify an additional network segment in addition to a main network segment.
//
//   - After the Lingjun network segment is created, the network segment cannot be modified.
//
//   - Make sure that you have a sufficient quota of Lingjun CIDR blocks.
//
//   - This interface is an asynchronous interface. After calling this interface, the system will return the ID of a Lingjun network segment. At this time, the Lingjun network segment may not be created yet, and the system background creation task is still in progress. You can call the ListVpds or GetVpd operation to query the status of the CIDR block of Lingjun.
//
//   - If the status of the Lingjun CIDR block is Executed, the CIDR block is being created.
//
//   - If the status of the Lingjun CIDR block is Available, the creation is successful.
//
// @param request - CreateVpdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpdResponse
func (client *Client) CreateVpdWithContext(ctx context.Context, request *CreateVpdRequest, runtime *dara.RuntimeOptions) (_result *CreateVpdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cidr) {
		body["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Subnets) {
		body["Subnets"] = request.Subnets
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpdName) {
		body["VpdName"] = request.VpdName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpd"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Users can use this API to authorize Lingjun HUB instances of the target account. After authorization, the target account can be associated with your Lingjun CIDR block by using the authorized Lingjun HUB instance.
//
// Description:
//
// When you call this operation to create cross-account authorization for Lingjun HUB, note that:
//
//   - Make sure that the Alibaba Cloud ID and Lingjun HUB instance that you want to authorize are correct.
//
//   - If you authorize the account of the other party, the account of the other party can load your local network instance to its Lingjun HUB, and the other party\\"s network will be connected to your network. Please proceed with caution.
//
// @param request - CreateVpdGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpdGrantRuleResponse
func (client *Client) CreateVpdGrantRuleWithContext(ctx context.Context, request *CreateVpdGrantRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateVpdGrantRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.GrantTenantId) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpdGrantRule"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpdGrantRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete Lingjun Elastic Network Interface. After deletion, all relevant data will be lost and cannot be recovered. Please operate with caution.
//
// @param request - DeleteElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteElasticNetworkInterfaceResponse
func (client *Client) DeleteElasticNetworkInterfaceWithContext(ctx context.Context, request *DeleteElasticNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteElasticNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteElasticNetworkInterface"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteElasticNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After you delete a Lingjun HUB instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete the Lingjun HUB, note that:
//
//   - Before you delete the instance, make sure that no network instance is connected to the Lingjun HUB instance.
//
//   - After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
//   - This interface is an asynchronous interface. After this interface is called, the Lingjun HUB instance may not be deleted, and the system background deletion task is still in progress. You can call the ListErs or GetEr operation to query the deletion status of the Lingjun HUB.
//
//   - If the status of the Lingjun HUB is Deleting, the Lingjun HUB instance is being deleted.
//
//   - If no Lingjun HUB instance is recorded, the Lingjun HUB instance has been deleted.
//
// @param request - DeleteErRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteErResponse
func (client *Client) DeleteErWithContext(ctx context.Context, request *DeleteErRequest, runtime *dara.RuntimeOptions) (_result *DeleteErResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEr"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteErResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you delete a network instance that is connected to an instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a network instance connection, take note of the following:
//
//   - Before you delete the instance, make sure that no routing policy exists under the network instance connection instance.
//
//   - After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
//   - This operation is an asynchronous operation. After you call this operation, the network instance that is connected to the instance may not be deleted. The system still deletes the instance in the background. You can call the ListErAttachments or GetErAttachment to query the deletion status of network instance connections:
//
//   - If the status of the network instance connection is Deleting, the network instance connection is being deleted.
//
//   - If there is no connection record for the network instance, the connection to the network instance has been deleted.
//
// @param request - DeleteErAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteErAttachmentResponse
func (client *Client) DeleteErAttachmentWithContext(ctx context.Context, request *DeleteErAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteErAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErAttachmentId) {
		body["ErAttachmentId"] = request.ErAttachmentId
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteErAttachment"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteErAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you delete a routing policy instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a routing policy, note that:
//
//   - After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
//   - This interface is an asynchronous interface. After this interface is called, the routing policy instance may not be deleted yet, and the system background deletion task is still in progress. You can call the ListErRouteMaps or GetErRouteMap operation to query the deletion status of a routing policy.
//
//   - If the routing policy is in the Deleting state, the routing policy instance is being deleted.
//
//   - If no routing policy instance is recorded, the routing policy instance has been deleted.
//
// @param request - DeleteErRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteErRouteMapResponse
func (client *Client) DeleteErRouteMapWithContext(ctx context.Context, request *DeleteErRouteMapRequest, runtime *dara.RuntimeOptions) (_result *DeleteErRouteMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ErRouteMapId) {
		body["ErRouteMapId"] = request.ErRouteMapId
	}

	if !dara.IsNil(request.ErRouteMapIds) {
		body["ErRouteMapIds"] = request.ErRouteMapIds
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteErRouteMap"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteErRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you delete a Lingjun subnet instance, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a Lingjun subnet, note that:
//
//   - After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
//   - This interface is an asynchronous interface. After this interface is called, the Lingjun subnet instance may not be deleted, and the system background deletion task is still in progress. You can call the ListSubnets or GetSubnet operation to query the deletion status of the subnet.
//
//   - If the status of the Lingjun subnet is Deleting, the Lingjun subnet instance is being deleted.
//
//   - If there is no record of the Lingjun subnet instance, the Lingjun subnet instance has been deleted.
//
// @param request - DeleteSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubnetResponse
func (client *Client) DeleteSubnetWithContext(ctx context.Context, request *DeleteSubnetRequest, runtime *dara.RuntimeOptions) (_result *DeleteSubnetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubnetId) {
		body["SubnetId"] = request.SubnetId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubnet"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubnetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If you delete a Lingjun HUB cross-account authorization that is connected to Lingjun, the related data is lost and cannot be recovered.
//
// @param request - DeleteVccGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVccGrantRuleResponse
func (client *Client) DeleteVccGrantRuleWithContext(ctx context.Context, request *DeleteVccGrantRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteVccGrantRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.GrantRuleId) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVccGrantRule"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVccGrantRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a Lingjun connection route entry.
//
// Description:
//
// When you call this operation to delete a VBR static route entry, note that:
//
//   - After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
//   - This operation is an asynchronous operation. After you call this operation, the VBR static route entries may not be deleted. The system still deletes the VBR static route entries in the background. You can call the ListVccRouteEntries or GetVccRouteEntry to query the deletion status of VBR static route entries:
//
//   - If the VBR static route entry is in the Deleting state, the VBR static route entry is being deleted.
//
//   - If no VBR static route entry instance is recorded, the VBR static route entry instance has been deleted.
//
// @param request - DeleteVccRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVccRouteEntryResponse
func (client *Client) DeleteVccRouteEntryWithContext(ctx context.Context, request *DeleteVccRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteVccRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	if !dara.IsNil(request.VccRouteEntryId) {
		body["VccRouteEntryId"] = request.VccRouteEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVccRouteEntry"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVccRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After you delete a Lingjun CIDR block, the related data is lost and cannot be recovered.
//
// Description:
//
// When you call this operation to delete a Lingjun CIDR block, take note of the following items:
//
//   - After deletion, all related data is lost and cannot be recovered. Exercise caution when performing this operation.
//
//   - Before deleting, make sure that all Lingjun subnet instances under the Lingjun CIDR block have been deleted.
//
//   - This interface is an asynchronous interface. After this interface is called, the Lingjun network segment instance may not be deleted, and the system background deletion task is still in progress. You can call the ListVpds or GetVpd operation to query the deletion status of the CIDR block.
//
//   - If the status of the Lingjun CIDR block is Deleting, the Lingjun CIDR block is being deleted.
//
//   - If there is no record of the Lingjun CIDR block instance, the Lingjun CIDR block instance has been deleted.
//
// @param request - DeleteVpdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpdResponse
func (client *Client) DeleteVpdWithContext(ctx context.Context, request *DeleteVpdRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpd"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the Lingjun HUB cross-account authorization for a Lingjun CIDR block. After the deletion, the related data is lost and cannot be recovered.
//
// @param request - DeleteVpdGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpdGrantRuleResponse
func (client *Client) DeleteVpdGrantRuleWithContext(ctx context.Context, request *DeleteVpdGrantRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpdGrantRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.GrantRuleId) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !dara.IsNil(request.GrantTenantId) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpdGrantRule"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpdGrantRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query whether the user has the SLR role-AliyunServiceRoleForEfloVcc required for Lingjun connection.
//
// Description:
//
// You can call this operation to query whether a user account has a **AliyunServiceRoleForEfloVcc*	- role.
//
// >  If you do not have a **AliyunServiceRoleForEfloVcc*	- role, you need to use the initializeVcc interface to complete authorization, otherwise users will not be able to use Lingjun to connect to the product.
//
// @param request - DescribeSlrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlrResponse
func (client *Client) DescribeSlrWithContext(ctx context.Context, request *DescribeSlrRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlr"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbind Lingjun ENI from NC.
//
// Description:
//
// This interface is an asynchronous interface, and you need to use the query interface to wait for the Lingjun Elastic Network Interface to reach the unbound state.
//
// @param request - DetachElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachElasticNetworkInterfaceResponse
func (client *Client) DetachElasticNetworkInterfaceWithContext(ctx context.Context, request *DetachElasticNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *DetachElasticNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachElasticNetworkInterface"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachElasticNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Users can use this API to query the destination CIDR block of the source policy instance when creating a routing strategy.
//
// @param request - GetDestinationCidrBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDestinationCidrBlockResponse
func (client *Client) GetDestinationCidrBlockWithContext(ctx context.Context, request *GetDestinationCidrBlockRequest, runtime *dara.RuntimeOptions) (_result *GetDestinationCidrBlockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDestinationCidrBlock"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDestinationCidrBlockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an LENI.
//
// @param request - GetElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetElasticNetworkInterfaceResponse
func (client *Client) GetElasticNetworkInterfaceWithContext(ctx context.Context, request *GetElasticNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *GetElasticNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetElasticNetworkInterface"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetElasticNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Lingjun HUB.
//
// @param request - GetErRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErResponse
func (client *Client) GetErWithContext(ctx context.Context, request *GetErRequest, runtime *dara.RuntimeOptions) (_result *GetErResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEr"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries network instance connections.
//
// @param request - GetErAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErAttachmentResponse
func (client *Client) GetErAttachmentWithContext(ctx context.Context, request *GetErAttachmentRequest, runtime *dara.RuntimeOptions) (_result *GetErAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErAttachmentId) {
		body["ErAttachmentId"] = request.ErAttachmentId
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetErAttachment"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of Lingjun HUB route entries.
//
// @param request - GetErRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErRouteEntryResponse
func (client *Client) GetErRouteEntryWithContext(ctx context.Context, request *GetErRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *GetErRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ErRouteEntryId) {
		body["ErRouteEntryId"] = request.ErRouteEntryId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetErRouteEntry"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// query lingjun hub routing policy details.
//
// @param request - GetErRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErRouteMapResponse
func (client *Client) GetErRouteMapWithContext(ctx context.Context, request *GetErRouteMapRequest, runtime *dara.RuntimeOptions) (_result *GetErRouteMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ErRouteMapId) {
		body["ErRouteMapId"] = request.ErRouteMapId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetErRouteMap"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the physical topology information of Lingjun network interface controller and Lingjun nodes under VPD.
//
// @param request - GetFabricTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFabricTopologyResponse
func (client *Client) GetFabricTopologyWithContext(ctx context.Context, request *GetFabricTopologyRequest, runtime *dara.RuntimeOptions) (_result *GetFabricTopologyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.LniIds) {
		body["LniIds"] = request.LniIds
	}

	if !dara.IsNil(request.NodeIds) {
		body["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFabricTopology"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFabricTopologyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of the secondary private IP address of a specified Lingjun Elastic Network Interface.
//
// @param request - GetLeniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLeniPrivateIpAddressResponse
func (client *Client) GetLeniPrivateIpAddressWithContext(ctx context.Context, request *GetLeniPrivateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *GetLeniPrivateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.IpName) {
		body["IpName"] = request.IpName
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLeniPrivateIpAddress"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLeniPrivateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details about the secondary private IP address.
//
// @param request - GetLniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLniPrivateIpAddressResponse
func (client *Client) GetLniPrivateIpAddressWithContext(ctx context.Context, request *GetLniPrivateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *GetLniPrivateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IpName) {
		body["IpName"] = request.IpName
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLniPrivateIpAddress"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLniPrivateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an LNI.
//
// @param request - GetNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkInterfaceResponse
func (client *Client) GetNetworkInterfaceWithContext(ctx context.Context, request *GetNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NetworkInterfaceId) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubnetId) {
		body["SubnetId"] = request.SubnetId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNetworkInterface"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network information of a node.
//
// @param request - GetNodeInfoForPodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeInfoForPodResponse
func (client *Client) GetNodeInfoForPodWithContext(ctx context.Context, request *GetNodeInfoForPodRequest, runtime *dara.RuntimeOptions) (_result *GetNodeInfoForPodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeInfoForPod"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeInfoForPodResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Lingjun subnet, including the type, CIDR block, instance ID, instance status, and number of NCs.
//
// @param request - GetSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubnetResponse
func (client *Client) GetSubnetWithContext(ctx context.Context, request *GetSubnetRequest, runtime *dara.RuntimeOptions) (_result *GetSubnetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubnetId) {
		body["SubnetId"] = request.SubnetId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubnet"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubnetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Lingjun connection, including the specification, Express Connect circuit access port type, instance status, bandwidth, and BGP CIDR block.
//
// @param request - GetVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVccResponse
func (client *Client) GetVccWithContext(ctx context.Context, request *GetVccRequest, runtime *dara.RuntimeOptions) (_result *GetVccResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVcc"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVccResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of cross-account resource authorization for a Lingjun connection, including the authorized tenant ID, Lingjun HUB instance ID, and network instance ID.
//
// @param request - GetVccGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVccGrantRuleResponse
func (client *Client) GetVccGrantRuleWithContext(ctx context.Context, request *GetVccGrantRuleRequest, runtime *dara.RuntimeOptions) (_result *GetVccGrantRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.GrantRuleId) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !dara.IsNil(request.GrantTenantId) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVccGrantRule"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVccGrantRuleResponse{}
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
// @param request - GetVccRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVccRouteEntryResponse
func (client *Client) GetVccRouteEntryWithContext(ctx context.Context, request *GetVccRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *GetVccRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	if !dara.IsNil(request.VccRouteEntryId) {
		body["VccRouteEntryId"] = request.VccRouteEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVccRouteEntry"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVccRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Lingjun CIDR block, including the status of the Lingjun CIDR block, the CIDR block, the number of subnets and NCs.
//
// @param request - GetVpdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpdResponse
func (client *Client) GetVpdWithContext(ctx context.Context, request *GetVpdRequest, runtime *dara.RuntimeOptions) (_result *GetVpdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVpd"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVpdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of cross-account resource authorization for a Lingjun CIDR block, including the authorized tenant ID, Lingjun HUB instance ID, and network instance ID.
//
// @param request - GetVpdGrantRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpdGrantRuleResponse
func (client *Client) GetVpdGrantRuleWithContext(ctx context.Context, request *GetVpdGrantRuleRequest, runtime *dara.RuntimeOptions) (_result *GetVpdGrantRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.GrantRuleId) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !dara.IsNil(request.GrantTenantId) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVpdGrantRule"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVpdGrantRuleResponse{}
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
// @param request - GetVpdRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpdRouteEntryResponse
func (client *Client) GetVpdRouteEntryWithContext(ctx context.Context, request *GetVpdRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *GetVpdRouteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.VpdRouteEntryId) {
		body["VpdRouteEntryId"] = request.VpdRouteEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVpdRouteEntry"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVpdRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initialize the Lingjun connection and authorize Intelligent Computing Lingjun to create an SLR in your account.
//
// @param request - InitializeVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeVccResponse
func (client *Client) InitializeVccWithContext(ctx context.Context, request *InitializeVccRequest, runtime *dara.RuntimeOptions) (_result *InitializeVccResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitializeVcc"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitializeVccResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the LENIs that are associated with a Lingjun node.
//
// @param request - ListElasticNetworkInterfacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListElasticNetworkInterfacesResponse
func (client *Client) ListElasticNetworkInterfacesWithContext(ctx context.Context, request *ListElasticNetworkInterfacesRequest, runtime *dara.RuntimeOptions) (_result *ListElasticNetworkInterfacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.Ip) {
		body["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NetworkType) {
		body["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.VSwitchId) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListElasticNetworkInterfaces"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListElasticNetworkInterfacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries network instance connections.
//
// @param request - ListErAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListErAttachmentsResponse
func (client *Client) ListErAttachmentsWithContext(ctx context.Context, request *ListErAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListErAttachmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoReceiveAllRoute) {
		body["AutoReceiveAllRoute"] = request.AutoReceiveAllRoute
	}

	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.ErAttachmentId) {
		body["ErAttachmentId"] = request.ErAttachmentId
	}

	if !dara.IsNil(request.ErAttachmentName) {
		body["ErAttachmentName"] = request.ErAttachmentName
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceTenantId) {
		body["ResourceTenantId"] = request.ResourceTenantId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListErAttachments"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListErAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route entries of the Lingjun HUB.
//
// @param request - ListErRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListErRouteEntriesResponse
func (client *Client) ListErRouteEntriesWithContext(ctx context.Context, request *ListErRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *ListErRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.IgnoreDetailedRouteEntry) {
		body["IgnoreDetailedRouteEntry"] = request.IgnoreDetailedRouteEntry
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NextHopId) {
		body["NextHopId"] = request.NextHopId
	}

	if !dara.IsNil(request.NextHopType) {
		body["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RouteType) {
		body["RouteType"] = request.RouteType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListErRouteEntries"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListErRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Routing policies are queried.
//
// @param request - ListErRouteMapsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListErRouteMapsResponse
func (client *Client) ListErRouteMapsWithContext(ctx context.Context, request *ListErRouteMapsRequest, runtime *dara.RuntimeOptions) (_result *ListErRouteMapsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ErRouteMapId) {
		body["ErRouteMapId"] = request.ErRouteMapId
	}

	if !dara.IsNil(request.ErRouteMapNum) {
		body["ErRouteMapNum"] = request.ErRouteMapNum
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReceptionInstanceId) {
		body["ReceptionInstanceId"] = request.ReceptionInstanceId
	}

	if !dara.IsNil(request.ReceptionInstanceName) {
		body["ReceptionInstanceName"] = request.ReceptionInstanceName
	}

	if !dara.IsNil(request.ReceptionInstanceType) {
		body["ReceptionInstanceType"] = request.ReceptionInstanceType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RouteMapAction) {
		body["RouteMapAction"] = request.RouteMapAction
	}

	if !dara.IsNil(request.TransmissionInstanceId) {
		body["TransmissionInstanceId"] = request.TransmissionInstanceId
	}

	if !dara.IsNil(request.TransmissionInstanceName) {
		body["TransmissionInstanceName"] = request.TransmissionInstanceName
	}

	if !dara.IsNil(request.TransmissionInstanceType) {
		body["TransmissionInstanceType"] = request.TransmissionInstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListErRouteMaps"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListErRouteMapsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Lingjun HUB.
//
// @param request - ListErsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListErsResponse
func (client *Client) ListErsWithContext(ctx context.Context, request *ListErsRequest, runtime *dara.RuntimeOptions) (_result *ListErsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ErName) {
		body["ErName"] = request.ErName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.MasterZoneId) {
		body["MasterZoneId"] = request.MasterZoneId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListErs"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListErsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the GPU node list of a specified GPU node whose communication distance does not exceed the specified NCD.
//
// @param request - ListInstancesByNcdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesByNcdResponse
func (client *Client) ListInstancesByNcdWithContext(ctx context.Context, request *ListInstancesByNcdRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesByNcdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.MaxNcd) {
		body["MaxNcd"] = request.MaxNcd
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstancesByNcd"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesByNcdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of secondary private IP addresses of Lingjun Elastic Network Interface.
//
// @param request - ListLeniPrivateIpAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLeniPrivateIpAddressesResponse
func (client *Client) ListLeniPrivateIpAddressesWithContext(ctx context.Context, request *ListLeniPrivateIpAddressesRequest, runtime *dara.RuntimeOptions) (_result *ListLeniPrivateIpAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.IpName) {
		body["IpName"] = request.IpName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLeniPrivateIpAddresses"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLeniPrivateIpAddressesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of secondary private IP addresses of Lingjun network interface controller.
//
// @param request - ListLniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLniPrivateIpAddressResponse
func (client *Client) ListLniPrivateIpAddressWithContext(ctx context.Context, request *ListLniPrivateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *ListLniPrivateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.Ip) {
		body["Ip"] = request.Ip
	}

	if !dara.IsNil(request.IpName) {
		body["IpName"] = request.IpName
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLniPrivateIpAddress"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLniPrivateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Lingjun network interfaces (LNIs).
//
// @param request - ListNetworkInterfacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkInterfacesResponse
func (client *Client) ListNetworkInterfacesWithContext(ctx context.Context, request *ListNetworkInterfacesRequest, runtime *dara.RuntimeOptions) (_result *ListNetworkInterfacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.Ip) {
		body["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SubnetId) {
		body["SubnetId"] = request.SubnetId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetworkInterfaces"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworkInterfacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries node network information.
//
// @param request - ListNodeInfosForPodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeInfosForPodResponse
func (client *Client) ListNodeInfosForPodWithContext(ctx context.Context, request *ListNodeInfosForPodRequest, runtime *dara.RuntimeOptions) (_result *ListNodeInfosForPodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeInfosForPod"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeInfosForPodResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the details of one or more Lingjun subnets, including the Lingjun subnet type, network address segment, and instance ID of the Lingjun CIDR block.
//
// @param request - ListSubnetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubnetsResponse
func (client *Client) ListSubnetsWithContext(ctx context.Context, request *ListSubnetsRequest, runtime *dara.RuntimeOptions) (_result *ListSubnetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.SubnetId) {
		body["SubnetId"] = request.SubnetId
	}

	if !dara.IsNil(request.SubnetName) {
		body["SubnetName"] = request.SubnetName
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubnets"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubnetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic rate of a Lingjun connection.
//
// @param request - ListVccFlowInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVccFlowInfosResponse
func (client *Client) ListVccFlowInfosWithContext(ctx context.Context, request *ListVccFlowInfosRequest, runtime *dara.RuntimeOptions) (_result *ListVccFlowInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		body["Direction"] = request.Direction
	}

	if !dara.IsNil(request.From) {
		body["From"] = request.From
	}

	if !dara.IsNil(request.MetricName) {
		body["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.To) {
		body["To"] = request.To
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVccFlowInfos"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVccFlowInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Lingjun connection authorization, including the authorized tenant ID, region, and Lingjun HUB instance information.
//
// @param request - ListVccGrantRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVccGrantRulesResponse
func (client *Client) ListVccGrantRulesWithContext(ctx context.Context, request *ListVccGrantRulesRequest, runtime *dara.RuntimeOptions) (_result *ListVccGrantRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ForSelect) {
		body["ForSelect"] = request.ForSelect
	}

	if !dara.IsNil(request.GrantRuleId) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !dara.IsNil(request.GrantTenantId) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVccGrantRules"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVccGrantRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Lingjun connection route entries.
//
// @param request - ListVccRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVccRouteEntriesResponse
func (client *Client) ListVccRouteEntriesWithContext(ctx context.Context, request *ListVccRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *ListVccRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.IgnoreDetailedRouteEntry) {
		body["IgnoreDetailedRouteEntry"] = request.IgnoreDetailedRouteEntry
	}

	if !dara.IsNil(request.NextHopId) {
		body["NextHopId"] = request.NextHopId
	}

	if !dara.IsNil(request.NextHopType) {
		body["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RouteType) {
		body["RouteType"] = request.RouteType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	if !dara.IsNil(request.VpdRouteEntryId) {
		body["VpdRouteEntryId"] = request.VpdRouteEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVccRouteEntries"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVccRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// query the details of one or more lingjun connections, including the specification, Express Connect circuit access port type, instance status, bandwidth, and bgp network segment.
//
// @param request - ListVccsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVccsResponse
func (client *Client) ListVccsWithContext(ctx context.Context, request *ListVccsRequest, runtime *dara.RuntimeOptions) (_result *ListVccsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.CenId) {
		body["CenId"] = request.CenId
	}

	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.ExStatus) {
		body["ExStatus"] = request.ExStatus
	}

	if !dara.IsNil(request.FilterErId) {
		body["FilterErId"] = request.FilterErId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVccs"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVccsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more route entries in the CIDR block of Lingjun, including the route type, route entry status, destination CIDR block, and instance information of the next route entry.
//
// @param request - ListVpdGrantRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpdGrantRulesResponse
func (client *Client) ListVpdGrantRulesWithContext(ctx context.Context, request *ListVpdGrantRulesRequest, runtime *dara.RuntimeOptions) (_result *ListVpdGrantRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ForSelect) {
		body["ForSelect"] = request.ForSelect
	}

	if !dara.IsNil(request.GrantRuleId) {
		body["GrantRuleId"] = request.GrantRuleId
	}

	if !dara.IsNil(request.GrantTenantId) {
		body["GrantTenantId"] = request.GrantTenantId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpdGrantRules"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpdGrantRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route entries of the Lingjun CIDR block.
//
// @param request - ListVpdRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpdRouteEntriesResponse
func (client *Client) ListVpdRouteEntriesWithContext(ctx context.Context, request *ListVpdRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *ListVpdRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.IgnoreDetailedRouteEntry) {
		body["IgnoreDetailedRouteEntry"] = request.IgnoreDetailedRouteEntry
	}

	if !dara.IsNil(request.NextHopId) {
		body["NextHopId"] = request.NextHopId
	}

	if !dara.IsNil(request.NextHopType) {
		body["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RouteType) {
		body["RouteType"] = request.RouteType
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.VpdRouteEntryId) {
		body["VpdRouteEntryId"] = request.VpdRouteEntryId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpdRouteEntries"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpdRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more Lingjun CIDR blocks, including the status of Lingjun CIDR blocks, Cidr addresses, service CIDR blocks, and Subnet.
//
// @param request - ListVpdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpdsResponse
func (client *Client) ListVpdsWithContext(ctx context.Context, request *ListVpdsRequest, runtime *dara.RuntimeOptions) (_result *ListVpdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnablePage) {
		body["EnablePage"] = request.EnablePage
	}

	if !dara.IsNil(request.FilterErId) {
		body["FilterErId"] = request.FilterErId
	}

	if !dara.IsNil(request.ForSelect) {
		body["ForSelect"] = request.ForSelect
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.VpdName) {
		body["VpdName"] = request.VpdName
	}

	if !dara.IsNil(request.WithDependence) {
		body["WithDependence"] = request.WithDependence
	}

	if !dara.IsNil(request.WithoutVcc) {
		body["WithoutVcc"] = request.WithoutVcc
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpds"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpdsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the network communication distance (Network Communication Distance,NCD) between instances (Lingjun node, Lingjun network interface controller).
//
// @param request - QueryInstanceNcdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceNcdResponse
func (client *Client) QueryInstanceNcdWithContext(ctx context.Context, request *QueryInstanceNcdRequest, runtime *dara.RuntimeOptions) (_result *QueryInstanceNcdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId1) {
		body["InstanceId1"] = request.InstanceId1
	}

	if !dara.IsNil(request.InstanceId2) {
		body["InstanceId2"] = request.InstanceId2
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInstanceNcd"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstanceNcdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Unsubscribe inactive Lingjun connection
//
// Description:
//
// Only unsubscribable for Lingjun connections in the prepayment status.
//
// @param request - RefundVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundVccResponse
func (client *Client) RefundVccWithContext(ctx context.Context, request *RefundVccRequest, runtime *dara.RuntimeOptions) (_result *RefundVccResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefundVcc"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefundVccResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retry trying to create /delete a Lingjun connection.
//
// Description:
//
// This operation allows the user to retry the operation if the Lingjun connection creation and deletion processes fail. Only the Lingjun connection in the creation failure and deletion failure state can be retried
//
// @param request - RetryVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryVccResponse
func (client *Client) RetryVccWithContext(ctx context.Context, request *RetryVccRequest, runtime *dara.RuntimeOptions) (_result *RetryVccResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryVcc"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryVccResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Switch the VCC connection instance or type
//
// @param request - SwitchVccConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchVccConnectionResponse
func (client *Client) SwitchVccConnectionWithContext(ctx context.Context, request *SwitchVccConnectionRequest, runtime *dara.RuntimeOptions) (_result *SwitchVccConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		body["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ConnectionType) {
		body["ConnectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchVccConnection"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchVccConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an assigned secondary private IP address.
//
// @param request - UnAssignPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnAssignPrivateIpAddressResponse
func (client *Client) UnAssignPrivateIpAddressWithContext(ctx context.Context, request *UnAssignPrivateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *UnAssignPrivateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IpName) {
		body["IpName"] = request.IpName
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		body["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		body["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubnetId) {
		body["SubnetId"] = request.SubnetId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnAssignPrivateIpAddress"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnAssignPrivateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This function can be used to delete the additional network segment of VPD.
//
// Description:
//
// *
//
// **Warning*	- If the attached CIDR block has Lingjun subnet resources, you must delete the dependent resources before you can delete the attached CIDR block.
//
// @param request - UnAssociateVpdCidrBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnAssociateVpdCidrBlockResponse
func (client *Client) UnAssociateVpdCidrBlockWithContext(ctx context.Context, request *UnAssociateVpdCidrBlockRequest, runtime *dara.RuntimeOptions) (_result *UnAssociateVpdCidrBlockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecondaryCidrBlock) {
		body["SecondaryCidrBlock"] = request.SecondaryCidrBlock
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnAssociateVpdCidrBlock"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnAssociateVpdCidrBlockResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the assigned secondary private IP address of Lingjun Elastic Network Interface.
//
// @param request - UnassignLeniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassignLeniPrivateIpAddressResponse
func (client *Client) UnassignLeniPrivateIpAddressWithContext(ctx context.Context, request *UnassignLeniPrivateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *UnassignLeniPrivateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.IpName) {
		body["IpName"] = request.IpName
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnassignLeniPrivateIpAddress"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnassignLeniPrivateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update Lingjun Elastic Network Interface information.
//
// @param request - UpdateElasticNetworkInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateElasticNetworkInterfaceResponse
func (client *Client) UpdateElasticNetworkInterfaceWithContext(ctx context.Context, request *UpdateElasticNetworkInterfaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateElasticNetworkInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateElasticNetworkInterface"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateElasticNetworkInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updated Lingjun HUB.
//
// @param request - UpdateErRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateErResponse
func (client *Client) UpdateErWithContext(ctx context.Context, request *UpdateErRequest, runtime *dara.RuntimeOptions) (_result *UpdateErResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ErName) {
		body["ErName"] = request.ErName
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEr"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateErResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a network instance connection.
//
// @param request - UpdateErAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateErAttachmentResponse
func (client *Client) UpdateErAttachmentWithContext(ctx context.Context, request *UpdateErAttachmentRequest, runtime *dara.RuntimeOptions) (_result *UpdateErAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ErAttachmentId) {
		body["ErAttachmentId"] = request.ErAttachmentId
	}

	if !dara.IsNil(request.ErAttachmentName) {
		body["ErAttachmentName"] = request.ErAttachmentName
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateErAttachment"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateErAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update some information about the routing policy, including the description and name of the routing policy.
//
// @param request - UpdateErRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateErRouteMapResponse
func (client *Client) UpdateErRouteMapWithContext(ctx context.Context, request *UpdateErRouteMapRequest, runtime *dara.RuntimeOptions) (_result *UpdateErRouteMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ErId) {
		body["ErId"] = request.ErId
	}

	if !dara.IsNil(request.ErRouteMapId) {
		body["ErRouteMapId"] = request.ErRouteMapId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateErRouteMap"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateErRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updated the description of the secondary private network assigned by the Lingjun Elastic Network Interface.
//
// @param request - UpdateLeniPrivateIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLeniPrivateIpAddressResponse
func (client *Client) UpdateLeniPrivateIpAddressWithContext(ctx context.Context, request *UpdateLeniPrivateIpAddressRequest, runtime *dara.RuntimeOptions) (_result *UpdateLeniPrivateIpAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ElasticNetworkInterfaceId) {
		body["ElasticNetworkInterfaceId"] = request.ElasticNetworkInterfaceId
	}

	if !dara.IsNil(request.IpName) {
		body["IpName"] = request.IpName
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLeniPrivateIpAddress"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLeniPrivateIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates some information about a specified subnet instance, including the name of the subnet instance.
//
// @param request - UpdateSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubnetResponse
func (client *Client) UpdateSubnetWithContext(ctx context.Context, request *UpdateSubnetRequest, runtime *dara.RuntimeOptions) (_result *UpdateSubnetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SubnetId) {
		body["SubnetId"] = request.SubnetId
	}

	if !dara.IsNil(request.SubnetName) {
		body["SubnetName"] = request.SubnetName
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSubnet"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSubnetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a Lingjun connection instance, including the peak bandwidth and name of the Lingjun connection instance.
//
// @param request - UpdateVccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVccResponse
func (client *Client) UpdateVccWithContext(ctx context.Context, request *UpdateVccRequest, runtime *dara.RuntimeOptions) (_result *UpdateVccResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.OrderId) {
		body["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VccId) {
		body["VccId"] = request.VccId
	}

	if !dara.IsNil(request.VccName) {
		body["VccName"] = request.VccName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVcc"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVccResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about the Lingjun CIDR block, including the name of the Lingjun CIDR block.
//
// @param request - UpdateVpdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpdResponse
func (client *Client) UpdateVpdWithContext(ctx context.Context, request *UpdateVpdRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpdId) {
		body["VpdId"] = request.VpdId
	}

	if !dara.IsNil(request.VpdName) {
		body["VpdName"] = request.VpdName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVpd"),
		Version:     dara.String("2022-05-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVpdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
