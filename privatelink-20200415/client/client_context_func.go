// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Adds an account ID to the whitelist of an endpoint service.
//
// Description:
//
// - Before you add an account ID to the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// - You cannot repeatedly call the **AddUserToVpcEndpointService*	- operation to add the ID of an Alibaba Cloud account to the whitelist of an endpoint service within a specified period of time.
//
// @param request - AddUserToVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToVpcEndpointServiceResponse
func (client *Client) AddUserToVpcEndpointServiceWithContext(ctx context.Context, request *AddUserToVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *AddUserToVpcEndpointServiceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.UserARN) {
		query["UserARN"] = request.UserARN
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToVpcEndpointService"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToVpcEndpointServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a zone to an endpoint.
//
// Description:
//
// - **AddZoneToVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the state of the zone.
//
//   - If the zone is in the **Creating*	- state, the zone is being added.
//
//   - If the zone is in the Wait state, the zone is added.
//
// - You cannot repeatedly call the **AddZoneToVpcEndpoint*	- operation to add a zone to an endpoint within a specified period of time.
//
// @param request - AddZoneToVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneToVpcEndpointResponse
func (client *Client) AddZoneToVpcEndpointWithContext(ctx context.Context, request *AddZoneToVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *AddZoneToVpcEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.Ipv6Address) {
		query["Ipv6Address"] = request.Ipv6Address
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.Ip) {
		query["ip"] = request.Ip
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddZoneToVpcEndpoint"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddZoneToVpcEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a service resource to an endpoint service.
//
// Description:
//
// - Before you add a service resource to an endpoint service, ensure that the endpoint service is in the **Active*	- state. Call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// - The **AttachResourceToVpcEndpointService*	- operation does not support concurrently adding service resources to the same endpoint service instance.
//
// @param request - AttachResourceToVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachResourceToVpcEndpointServiceResponse
func (client *Client) AttachResourceToVpcEndpointServiceWithContext(ctx context.Context, request *AttachResourceToVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *AttachResourceToVpcEndpointServiceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachResourceToVpcEndpointService"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachResourceToVpcEndpointServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an endpoint with a security group.
//
// Description:
//
// - **AttachSecurityGroupToVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) operation to query the state of the endpoint.
//
//   - If the endpoint is in the **Pending*	- state, the endpoint is being associated with the security group.
//
//   - If the endpoint is in the **Active*	- state, the endpoint is associated with the security group.
//
// - You cannot repeatedly call the **AttachSecurityGroupToVpcEndpoint*	- operation to associate an endpoint with a security group within a specified period of time.
//
// @param request - AttachSecurityGroupToVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachSecurityGroupToVpcEndpointResponse
func (client *Client) AttachSecurityGroupToVpcEndpointWithContext(ctx context.Context, request *AttachSecurityGroupToVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *AttachSecurityGroupToVpcEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachSecurityGroupToVpcEndpoint"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachSecurityGroupToVpcEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a resource group.
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
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2020-04-15"),
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
// Creates an endpoint.
//
// Description:
//
// *CreateVpcEndpoint*	- is an asynchronous operation. After you send the request, the system returns an instance ID and starts creating the endpoint in the background. You can call [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/469349.html) to check the endpoint\\"s creation status:
//
// - **Creating**: The endpoint is being created.
//
// - **Active**: The endpoint is available.
//
// @param request - CreateVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcEndpointResponse
func (client *Client) CreateVpcEndpointWithContext(ctx context.Context, request *CreateVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressIpVersion) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossRegionBandwidth) {
		query["CrossRegionBandwidth"] = request.CrossRegionBandwidth
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointDescription) {
		query["EndpointDescription"] = request.EndpointDescription
	}

	if !dara.IsNil(request.EndpointName) {
		query["EndpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
	}

	if !dara.IsNil(request.PolicyDocument) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !dara.IsNil(request.ProtectedEnabled) {
		query["ProtectedEnabled"] = request.ProtectedEnabled
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.Zone) {
		query["Zone"] = request.Zone
	}

	if !dara.IsNil(request.ZoneAffinityEnabled) {
		query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	}

	if !dara.IsNil(request.ZonePrivateIpAddressCount) {
		query["ZonePrivateIpAddressCount"] = request.ZonePrivateIpAddressCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcEndpoint"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an endpoint service by calling the CreateVpcEndpointService operation.
//
// Description:
//
// - Before you create an endpoint service, make sure that you have created a load balancer resource that supports PrivateLink. For more information, see [Create a load balancer instance](https://help.aliyun.com/document_detail/174064.html).
//
// - **CreateVpcEndpointService*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. The endpoint service is not created at the time the request is returned. You can call [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) to query the creation status of the endpoint service:
//
//   - If the endpoint service is in the **Creating*	- state, the endpoint service is being created.
//
//   - If the endpoint service is in the **Active*	- state, the endpoint service is created.
//
// @param request - CreateVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcEndpointServiceResponse
func (client *Client) CreateVpcEndpointServiceWithContext(ctx context.Context, request *CreateVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcEndpointServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressIpVersion) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !dara.IsNil(request.AutoAcceptEnabled) {
		query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Payer) {
		query["Payer"] = request.Payer
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceDescription) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !dara.IsNil(request.ServiceResourceType) {
		query["ServiceResourceType"] = request.ServiceResourceType
	}

	if !dara.IsNil(request.ServiceSupportIPv6) {
		query["ServiceSupportIPv6"] = request.ServiceSupportIPv6
	}

	if !dara.IsNil(request.SupportedRegionList) {
		query["SupportedRegionList"] = request.SupportedRegionList
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.ZoneAffinityEnabled) {
		query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcEndpointService"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcEndpointServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an endpoint.
//
// Description:
//
// - Before you delete an endpoint, make sure that all zones associated with the endpoint have been deleted.
//
// - **DeleteVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and deletes the endpoint in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/469349.html) operation to query the status of the endpoint.
//
//   - If the endpoint status is **Deleting**, the endpoint is being deleted.
//
//   - If the endpoint cannot be queried, the endpoint has been deleted.
//
// @param request - DeleteVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointResponse
func (client *Client) DeleteVpcEndpointWithContext(ctx context.Context, request *DeleteVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpcEndpoint"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an endpoint service.
//
// Description:
//
// - Before you delete an endpoint service, you must disconnect the endpoint from the endpoint service and remove the service resources.
//
// - **DeleteVpcEndpointService*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to check whether the endpoint service is deleted.
//
//   - If the endpoint service is in the **Deleting*	- state, the endpoint service is being deleted.
//
//   - If the endpoint service cannot be queried, the endpoint service is deleted.
//
// - You cannot repeatedly call the **DeleteVpcEndpointService*	- operation to delete an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - DeleteVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointServiceResponse
func (client *Client) DeleteVpcEndpointServiceWithContext(ctx context.Context, request *DeleteVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcEndpointServiceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpcEndpointService"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpcEndpointServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeRegions operation to query the list of regions supported by PrivateLink.
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
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceResourceType) {
		query["ServiceResourceType"] = request.ServiceResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2020-04-15"),
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
// Queries the zones supported by PrivateLink in a specified region. In cross-region scenarios, the supported zones on the endpoint side and the endpoint service side are independent of each other. This operation allows you to query the supported zones for the endpoint side and the endpoint service side separately.
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
	if !dara.IsNil(request.CrossRegion) {
		query["CrossRegion"] = request.CrossRegion
	}

	if !dara.IsNil(request.CrossRegionSide) {
		query["CrossRegionSide"] = request.CrossRegionSide
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceResourceType) {
		query["ServiceResourceType"] = request.ServiceResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2020-04-15"),
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
// Removes a service resource from an endpoint service.
//
// Description:
//
// - Before you remove a service resource from an endpoint service, make sure that the endpoint service is in the **Active*	- status. Call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// - The **DetachResourceFromVpcEndpointService*	- operation does not support the concurrent removal of service resources from the same endpoint service instance.
//
// @param request - DetachResourceFromVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachResourceFromVpcEndpointServiceResponse
func (client *Client) DetachResourceFromVpcEndpointServiceWithContext(ctx context.Context, request *DetachResourceFromVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *DetachResourceFromVpcEndpointServiceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachResourceFromVpcEndpointService"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachResourceFromVpcEndpointServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an endpoint from a security group.
//
// Description:
//
// - **DetachSecurityGroupFromVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) to check whether the endpoint is disassociated from the security group.
//
//   - If the endpoint is in the **Pending*	- state, the endpoint is being disassociated from the security group.
//
//   - If you cannot query the endpoint in the security group, the endpoint is disassociated from the security group.
//
// - You cannot repeatedly call the **DetachSecurityGroupFromVpcEndpoint*	- operation to disassociate an endpoint from a security group within a specified period of time.
//
// @param request - DetachSecurityGroupFromVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachSecurityGroupFromVpcEndpointResponse
func (client *Client) DetachSecurityGroupFromVpcEndpointWithContext(ctx context.Context, request *DetachSecurityGroupFromVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *DetachSecurityGroupFromVpcEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachSecurityGroupFromVpcEndpoint"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachSecurityGroupFromVpcEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rejects a connection request from an endpoint.
//
// Description:
//
// - **DisableVpcEndpointConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//   - If the endpoint connection is in the **Disconnecting*	- state, the endpoint is being disconnected from the endpoint service.
//
//   - If the endpoint connection is in the **Disconnected*	- state, the endpoint is disconnected from the endpoint service.
//
// - You cannot repeatedly call the **DisableVpcEndpointConnection*	- operation to allow an endpoint service to reject a connection request from an endpoint within a specified period of time.
//
// @param request - DisableVpcEndpointConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVpcEndpointConnectionResponse
func (client *Client) DisableVpcEndpointConnectionWithContext(ctx context.Context, request *DisableVpcEndpointConnectionRequest, runtime *dara.RuntimeOptions) (_result *DisableVpcEndpointConnectionResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableVpcEndpointConnection"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableVpcEndpointConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Closes connections in an endpoint zone.
//
// Description:
//
// - You can call this operation only when the state of the endpoint is **Connected*	- and the state of the zone associated with the endpoint is **Connected*	- or **Migrated**.
//
// - **DisableVpcEndpointZoneConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the status of the task.
//
//   - If the zone is in the **Disconnecting*	- state, the task is running.
//
//   - If the zone is in the **Disconnected*	- state, the task is successful.
//
// - You cannot repeatedly call the **DisableVpcEndpointZoneConnection*	- operation to allow an endpoint service to reject a connection request from the endpoint in the zone within a specified period of time.
//
// @param request - DisableVpcEndpointZoneConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVpcEndpointZoneConnectionResponse
func (client *Client) DisableVpcEndpointZoneConnectionWithContext(ctx context.Context, request *DisableVpcEndpointZoneConnectionRequest, runtime *dara.RuntimeOptions) (_result *DisableVpcEndpointZoneConnectionResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplacedResource) {
		query["ReplacedResource"] = request.ReplacedResource
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableVpcEndpointZoneConnection"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableVpcEndpointZoneConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the EnableVpcEndpointConnection operation, and the endpoint service accepts the connection request from the endpoint.
//
// Description:
//
// - **EnableVpcEndpointConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID. However, the endpoint service connection is not yet established, and the connection task is still being run in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the status of the endpoint service connection:
//
//   - If the connection is in the **Connecting*	- state, the endpoint service connection is being established.
//
//   - If the connection is in the **Connected*	- state, the endpoint service connection is established.
//
// - The **EnableVpcEndpointConnection*	- operation does not support concurrently accepting endpoint connection requests under the same Alibaba Cloud account (primary account).
//
// @param request - EnableVpcEndpointConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableVpcEndpointConnectionResponse
func (client *Client) EnableVpcEndpointConnectionWithContext(ctx context.Context, request *EnableVpcEndpointConnectionRequest, runtime *dara.RuntimeOptions) (_result *EnableVpcEndpointConnectionResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.TrafficControlMode) {
		query["TrafficControlMode"] = request.TrafficControlMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableVpcEndpointConnection"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableVpcEndpointConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Allows connections to endpoint zones.
//
// Description:
//
// - You can call this operation only when the state of the endpoint is **Connected*	- and the state of the associated zone is **Disconnected**.
//
// - **EnableVpcEndpointZoneConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the endpoint service accepts a connection request from the endpoint in the associated zone.
//
//   - If the zone is in the **Connecting*	- state, the endpoint service is accepting the connection request from the endpoint.
//
//   - If the zone is in the **Connected*	- state, the endpoint service has accepted the connection request from the endpoint.
//
// - You cannot repeatedly call the **EnableVpcEndpointZoneConnection*	- operation to allow an endpoint service to accept a connection request from an endpoint in the associated zone within a specified period of time.
//
// @param request - EnableVpcEndpointZoneConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableVpcEndpointZoneConnectionResponse
func (client *Client) EnableVpcEndpointZoneConnectionWithContext(ctx context.Context, request *EnableVpcEndpointZoneConnectionRequest, runtime *dara.RuntimeOptions) (_result *EnableVpcEndpointZoneConnectionResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableVpcEndpointZoneConnection"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableVpcEndpointZoneConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use `GetVpcEndpointAttribute` to query the attributes of a specified endpoint.
//
// @param request - GetVpcEndpointAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcEndpointAttributeResponse
func (client *Client) GetVpcEndpointAttributeWithContext(ctx context.Context, request *GetVpcEndpointAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetVpcEndpointAttributeResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVpcEndpointAttribute"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVpcEndpointAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call `GetVpcEndpointServiceAttribute` to retrieve the properties of an endpoint service in your account.
//
// @param request - GetVpcEndpointServiceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcEndpointServiceAttributeResponse
func (client *Client) GetVpcEndpointServiceAttributeWithContext(ctx context.Context, request *GetVpcEndpointServiceAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetVpcEndpointServiceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVpcEndpointServiceAttribute"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVpcEndpointServiceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to resources.
//
// Description:
//
// - You must specify **ResourceId.N*	- or **Tag.N*	- in the request to specify the object that you want to query.
//
// - **Tag.N*	- is a resource tag that consists of a key-value pair (Tag.N.Key and Tag.N.Value). If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
// - If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
// - If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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
		Version:     dara.String("2020-04-15"),
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
// Queries endpoint connections by calling the ListVpcEndpointConnections operation.
//
// @param request - ListVpcEndpointConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointConnectionsResponse
func (client *Client) ListVpcEndpointConnectionsWithContext(ctx context.Context, request *ListVpcEndpointConnectionsRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointConnectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStatus) {
		query["ConnectionStatus"] = request.ConnectionStatus
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.EndpointOwnerId) {
		query["EndpointOwnerId"] = request.EndpointOwnerId
	}

	if !dara.IsNil(request.EniId) {
		query["EniId"] = request.EniId
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplacedResourceId) {
		query["ReplacedResourceId"] = request.ReplacedResourceId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcEndpointConnections"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointConnectionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security group that is associated with an endpoint.
//
// @param request - ListVpcEndpointSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointSecurityGroupsResponse
func (client *Client) ListVpcEndpointSecurityGroupsWithContext(ctx context.Context, request *ListVpcEndpointSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointSecurityGroupsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcEndpointSecurityGroups"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointSecurityGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the service resources that are added to an endpoint service.
//
// @param request - ListVpcEndpointServiceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServiceResourcesResponse
func (client *Client) ListVpcEndpointServiceResourcesWithContext(ctx context.Context, request *ListVpcEndpointServiceResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointServiceResourcesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcEndpointServiceResources"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointServiceResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the whitelist of an endpoint service.
//
// @param request - ListVpcEndpointServiceUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServiceUsersResponse
func (client *Client) ListVpcEndpointServiceUsersWithContext(ctx context.Context, request *ListVpcEndpointServiceUsersRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointServiceUsersResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserListType) {
		query["UserListType"] = request.UserListType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcEndpointServiceUsers"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointServiceUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the ListVpcEndpointServices operation to query the endpoint services owned by your account.
//
// @param request - ListVpcEndpointServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServicesResponse
func (client *Client) ListVpcEndpointServicesWithContext(ctx context.Context, request *ListVpcEndpointServicesRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressIpVersion) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !dara.IsNil(request.AutoAcceptEnabled) {
		query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ServiceBusinessStatus) {
		query["ServiceBusinessStatus"] = request.ServiceBusinessStatus
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceResourceType) {
		query["ServiceResourceType"] = request.ServiceResourceType
	}

	if !dara.IsNil(request.ServiceStatus) {
		query["ServiceStatus"] = request.ServiceStatus
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.ZoneAffinityEnabled) {
		query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcEndpointServices"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointServicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the ListVpcEndpointServicesByEndUser operation to retrieve a list of endpoint services for which the current account is whitelisted. This list includes services that you can connect to by creating new endpoints, as well as services that are already connected to your existing ones.
//
// @param request - ListVpcEndpointServicesByEndUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServicesByEndUserResponse
func (client *Client) ListVpcEndpointServicesByEndUserWithContext(ctx context.Context, request *ListVpcEndpointServicesByEndUserRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointServicesByEndUserResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcEndpointServicesByEndUser"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointServicesByEndUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the ListVpcEndpointZones operation to query the zones of an endpoint.
//
// @param request - ListVpcEndpointZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointZonesResponse
func (client *Client) ListVpcEndpointZonesWithContext(ctx context.Context, request *ListVpcEndpointZonesRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointZonesResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcEndpointZones"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The `ListVpcEndpoints` operation lists endpoints.
//
// @param request - ListVpcEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointsResponse
func (client *Client) ListVpcEndpointsWithContext(ctx context.Context, request *ListVpcEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressIpVersion) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !dara.IsNil(request.ConnectionStatus) {
		query["ConnectionStatus"] = request.ConnectionStatus
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.EndpointName) {
		query["EndpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.EndpointStatus) {
		query["EndpointStatus"] = request.EndpointStatus
	}

	if !dara.IsNil(request.EndpointType) {
		query["EndpointType"] = request.EndpointType
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
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
		Action:      dara.String("ListVpcEndpoints"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates PrivateLink.
//
// @param request - OpenPrivateLinkServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenPrivateLinkServiceResponse
func (client *Client) OpenPrivateLinkServiceWithContext(ctx context.Context, request *OpenPrivateLinkServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenPrivateLinkServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenPrivateLinkService"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenPrivateLinkServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an account ID from the whitelist of an endpoint service.
//
// Description:
//
// - Before you remove an account ID from the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// - You cannot repeatedly call the **RemoveUserFromVpcEndpointService*	- operation to remove the ID of an Alibaba Cloud account from the whitelist of an endpoint service within a specified period of time.
//
// @param request - RemoveUserFromVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromVpcEndpointServiceResponse
func (client *Client) RemoveUserFromVpcEndpointServiceWithContext(ctx context.Context, request *RemoveUserFromVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromVpcEndpointServiceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.UserARN) {
		query["UserARN"] = request.UserARN
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserFromVpcEndpointService"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserFromVpcEndpointServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a zone of an endpoint.
//
// Description:
//
// - **RemoveZoneFromVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the zone of the endpoint is deleted.
//
//   - If the zone of the endpoint is in the **Deleting*	- state, the zone is being deleted.
//
//   - If the zone cannot be queried, the zone is deleted.
//
// - You cannot repeatedly call the **RemoveZoneFromVpcEndpoint*	- operation to delete a zone of an endpoint within a specified period of time.
//
// @param request - RemoveZoneFromVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveZoneFromVpcEndpointResponse
func (client *Client) RemoveZoneFromVpcEndpointWithContext(ctx context.Context, request *RemoveZoneFromVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *RemoveZoneFromVpcEndpointResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveZoneFromVpcEndpoint"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveZoneFromVpcEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to resources. You can call this API operation to add tags to one or more endpoints or endpoint services.
//
// Description:
//
// > You can add up to 20 tags to an instance. Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number is reached, an error message is returned.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2020-04-15"),
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
// Removes tags from one or more endpoints or endpoint services at a time.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		body["All"] = request.All
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		bodyFlat["TagKey"] = request.TagKey
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2020-04-15"),
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
// The `UpdateVpcEndpointAttribute` operation modifies the properties of an endpoint.
//
// Description:
//
// The **UpdateVpcEndpointAttribute*	- operation cannot be called concurrently to modify endpoint properties for the same Alibaba Cloud main account.
//
// @param request - UpdateVpcEndpointAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointAttributeResponse
func (client *Client) UpdateVpcEndpointAttributeWithContext(ctx context.Context, request *UpdateVpcEndpointAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressIpVersion) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossRegionBandwidth) {
		query["CrossRegionBandwidth"] = request.CrossRegionBandwidth
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointDescription) {
		query["EndpointDescription"] = request.EndpointDescription
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.EndpointName) {
		query["EndpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.PolicyDocument) {
		query["PolicyDocument"] = request.PolicyDocument
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResetPolicy) {
		query["ResetPolicy"] = request.ResetPolicy
	}

	if !dara.IsNil(request.ZoneAffinityEnabled) {
		query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVpcEndpointAttribute"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVpcEndpointAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the UpdateVpcEndpointConnectionAttribute operation to modify the attributes of an endpoint connection.
//
// Description:
//
// The **UpdateVpcEndpointConnectionAttribute*	- operation does not support concurrent modifications of the endpoint connection bandwidth within the same Alibaba Cloud account (primary account).
//
// @param request - UpdateVpcEndpointConnectionAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointConnectionAttributeResponse
func (client *Client) UpdateVpcEndpointConnectionAttributeWithContext(ctx context.Context, request *UpdateVpcEndpointConnectionAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointConnectionAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.TrafficControlMode) {
		query["TrafficControlMode"] = request.TrafficControlMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVpcEndpointConnectionAttribute"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVpcEndpointConnectionAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The UpdateVpcEndpointServiceAttribute operation modifies the attributes of an endpoint service.
//
// Description:
//
// - Before you modify the attributes of an endpoint service, ensure that the service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to check the status of the endpoint service.
//
// - The **UpdateVpcEndpointServiceAttribute*	- operation does not support concurrent modifications to the attributes of an endpoint service within the same Alibaba Cloud account (root user).
//
// @param request - UpdateVpcEndpointServiceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointServiceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceAttributeWithContext(ctx context.Context, request *UpdateVpcEndpointServiceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointServiceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddSupportedRegionSet) {
		query["AddSupportedRegionSet"] = request.AddSupportedRegionSet
	}

	if !dara.IsNil(request.AddressIpVersion) {
		query["AddressIpVersion"] = request.AddressIpVersion
	}

	if !dara.IsNil(request.AutoAcceptEnabled) {
		query["AutoAcceptEnabled"] = request.AutoAcceptEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectBandwidth) {
		query["ConnectBandwidth"] = request.ConnectBandwidth
	}

	if !dara.IsNil(request.DeleteSupportedRegionSet) {
		query["DeleteSupportedRegionSet"] = request.DeleteSupportedRegionSet
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceDescription) {
		query["ServiceDescription"] = request.ServiceDescription
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceSupportIPv6) {
		query["ServiceSupportIPv6"] = request.ServiceSupportIPv6
	}

	if !dara.IsNil(request.ZoneAffinityEnabled) {
		query["ZoneAffinityEnabled"] = request.ZoneAffinityEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVpcEndpointServiceAttribute"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVpcEndpointServiceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The UpdateVpcEndpointServiceResourceAttribute operation enables or disables automatic allocation for a service resource of an endpoint service.
//
// Description:
//
// - Before you modify the properties of a service resource, ensure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
// - The **UpdateVpcEndpointServiceResourceAttribute*	- operation does not support concurrent modifications to the properties of service resources that belong to the same endpoint service.
//
// @param request - UpdateVpcEndpointServiceResourceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointServiceResourceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceResourceAttributeWithContext(ctx context.Context, request *UpdateVpcEndpointServiceResourceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointServiceResourceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoAllocatedEnabled) {
		query["AutoAllocatedEnabled"] = request.AutoAllocatedEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVpcEndpointServiceResourceAttribute"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVpcEndpointServiceResourceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a service resource of a zone to which an endpoint connection belongs.
//
// Description:
//
// ### Prerequisites
//
// By default, the feature of modifying a service resource of a zone to which an endpoint connection belongs is unavailable. To use this feature, log on to the [Quota Center console](https://quotas.console.aliyun.com/white-list-products/privatelink/quotas). Click Whitelist Quotas in the left-side navigation pane and click PrivateLink in the Networking section. On the page that appears, search for `privatelink_whitelist/svc_res_mgt_uat` and click Apply in the Actions column.
//
// ### Usage notes
//
// - If the endpoint connection is in the **Disconnected*	- state, you can manually allocate the service resource in the zone.
//
// - If the endpoint connection is in the **Connected*	- state, you can manually migrate the service resource in the zone. In this case, the connection might be interrupted.
//
// - **UpdateVpcEndpointZoneConnectionResourceAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to check whether the service resource is modified.
//
//   - If the endpoint service is in the **Pending*	- state, the service resource is being modified.
//
//   - If the endpoint service is in the **Active*	- state, the service resource is modified.
//
// - You cannot repeatedly call the **UpdateVpcEndpointZoneConnectionResourceAttribute*	- operation to modify a service resource in the zone to which an endpoint connection belongs within a specified period of time.
//
// @param request - UpdateVpcEndpointZoneConnectionResourceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointZoneConnectionResourceAttributeResponse
func (client *Client) UpdateVpcEndpointZoneConnectionResourceAttributeWithContext(ctx context.Context, request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.EndpointId) {
		query["EndpointId"] = request.EndpointId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceAllocateMode) {
		query["ResourceAllocateMode"] = request.ResourceAllocateMode
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceReplaceMode) {
		query["ResourceReplaceMode"] = request.ResourceReplaceMode
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceId) {
		query["ServiceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVpcEndpointZoneConnectionResourceAttribute"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVpcEndpointZoneConnectionResourceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
