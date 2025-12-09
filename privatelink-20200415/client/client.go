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
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("privatelink"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds an account ID to the whitelist of an endpoint service.
//
// Description:
//
//	  Before you add an account ID to the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
//		- You cannot repeatedly call the **AddUserToVpcEndpointService*	- operation to add the ID of an Alibaba Cloud account to the whitelist of an endpoint service within a specified period of time.
//
// @param request - AddUserToVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToVpcEndpointServiceResponse
func (client *Client) AddUserToVpcEndpointServiceWithOptions(request *AddUserToVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *AddUserToVpcEndpointServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an account ID to the whitelist of an endpoint service.
//
// Description:
//
//	  Before you add an account ID to the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
//		- You cannot repeatedly call the **AddUserToVpcEndpointService*	- operation to add the ID of an Alibaba Cloud account to the whitelist of an endpoint service within a specified period of time.
//
// @param request - AddUserToVpcEndpointServiceRequest
//
// @return AddUserToVpcEndpointServiceResponse
func (client *Client) AddUserToVpcEndpointService(request *AddUserToVpcEndpointServiceRequest) (_result *AddUserToVpcEndpointServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddUserToVpcEndpointServiceResponse{}
	_body, _err := client.AddUserToVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a zone to an endpoint.
//
// Description:
//
//	  **AddZoneToVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the state of the zone.
//
//	    	- If the zone is in the **Creating*	- state, the zone is being added.
//
//	    	- If the zone is in the Wait state, the zone is added.
//
//		- You cannot repeatedly call the **AddZoneToVpcEndpoint*	- operation to add a zone to an endpoint within a specified period of time.
//
// @param request - AddZoneToVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddZoneToVpcEndpointResponse
func (client *Client) AddZoneToVpcEndpointWithOptions(request *AddZoneToVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *AddZoneToVpcEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  **AddZoneToVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the state of the zone.
//
//	    	- If the zone is in the **Creating*	- state, the zone is being added.
//
//	    	- If the zone is in the Wait state, the zone is added.
//
//		- You cannot repeatedly call the **AddZoneToVpcEndpoint*	- operation to add a zone to an endpoint within a specified period of time.
//
// @param request - AddZoneToVpcEndpointRequest
//
// @return AddZoneToVpcEndpointResponse
func (client *Client) AddZoneToVpcEndpoint(request *AddZoneToVpcEndpointRequest) (_result *AddZoneToVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddZoneToVpcEndpointResponse{}
	_body, _err := client.AddZoneToVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a service resource to an endpoint service.
//
// Description:
//
//	  Before you add a service resource to an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
//		- You cannot repeatedly call the **AttachResourceToVpcEndpointService*	- operation to add a service resource to an endpoint service within a specified period of time.
//
// @param request - AttachResourceToVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachResourceToVpcEndpointServiceResponse
func (client *Client) AttachResourceToVpcEndpointServiceWithOptions(request *AttachResourceToVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *AttachResourceToVpcEndpointServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  Before you add a service resource to an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
//		- You cannot repeatedly call the **AttachResourceToVpcEndpointService*	- operation to add a service resource to an endpoint service within a specified period of time.
//
// @param request - AttachResourceToVpcEndpointServiceRequest
//
// @return AttachResourceToVpcEndpointServiceResponse
func (client *Client) AttachResourceToVpcEndpointService(request *AttachResourceToVpcEndpointServiceRequest) (_result *AttachResourceToVpcEndpointServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachResourceToVpcEndpointServiceResponse{}
	_body, _err := client.AttachResourceToVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates an endpoint with a security group.
//
// Description:
//
//	  **AttachSecurityGroupToVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) operation to query the state of the endpoint.
//
//	    	- If the endpoint is in the **Pending*	- state, the endpoint is being associated with the security group.
//
//	    	- If the endpoint is in the **Active*	- state, the endpoint is associated with the security group.
//
//		- You cannot repeatedly call the **AttachSecurityGroupToVpcEndpoint*	- operation to associate an endpoint with a security group within a specified period of time.
//
// @param request - AttachSecurityGroupToVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachSecurityGroupToVpcEndpointResponse
func (client *Client) AttachSecurityGroupToVpcEndpointWithOptions(request *AttachSecurityGroupToVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *AttachSecurityGroupToVpcEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  **AttachSecurityGroupToVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) operation to query the state of the endpoint.
//
//	    	- If the endpoint is in the **Pending*	- state, the endpoint is being associated with the security group.
//
//	    	- If the endpoint is in the **Active*	- state, the endpoint is associated with the security group.
//
//		- You cannot repeatedly call the **AttachSecurityGroupToVpcEndpoint*	- operation to associate an endpoint with a security group within a specified period of time.
//
// @param request - AttachSecurityGroupToVpcEndpointRequest
//
// @return AttachSecurityGroupToVpcEndpointResponse
func (client *Client) AttachSecurityGroupToVpcEndpoint(request *AttachSecurityGroupToVpcEndpointRequest) (_result *AttachSecurityGroupToVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachSecurityGroupToVpcEndpointResponse{}
	_body, _err := client.AttachSecurityGroupToVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Queries whether PrivateLink is activated.
//
// @param request - CheckProductOpenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckProductOpenResponse
func (client *Client) CheckProductOpenWithOptions(runtime *dara.RuntimeOptions) (_result *CheckProductOpenResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("CheckProductOpen"),
		Version:     dara.String("2020-04-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckProductOpenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether PrivateLink is activated.
//
// @return CheckProductOpenResponse
func (client *Client) CheckProductOpen() (_result *CheckProductOpenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckProductOpenResponse{}
	_body, _err := client.CheckProductOpenWithOptions(runtime)
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
// Description:
//
// *CreateVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to check whether the endpoint is created.
//
//   - If the endpoint is in the **Creating*	- state, the endpoint is being created.
//
//   - If the endpoint is in the **Active*	- state, the endpoint is created.
//
// @param request - CreateVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcEndpointResponse
func (client *Client) CreateVpcEndpointWithOptions(request *CreateVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
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
// Description:
//
// *CreateVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to check whether the endpoint is created.
//
//   - If the endpoint is in the **Creating*	- state, the endpoint is being created.
//
//   - If the endpoint is in the **Active*	- state, the endpoint is created.
//
// @param request - CreateVpcEndpointRequest
//
// @return CreateVpcEndpointResponse
func (client *Client) CreateVpcEndpoint(request *CreateVpcEndpointRequest) (_result *CreateVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CreateVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an endpoint service.
//
// Description:
//
//	  Before you create an endpoint service, make sure that you have created a Server Load Balancer (SLB) instance that supports PrivateLink. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/174064.html).
//
//		- **CreateVpcEndpointService*	- is an asynchronous operation. After a request is sent, the system returns a request ID and an instance ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to query the status of the endpoint service.
//
//	    	- If the endpoint service is in the **Creating*	- state, the endpoint service is being created.
//
//	    	- If the endpoint service is in the **Active*	- state, the endpoint service is created.
//
// @param request - CreateVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcEndpointServiceResponse
func (client *Client) CreateVpcEndpointServiceWithOptions(request *CreateVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcEndpointServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an endpoint service.
//
// Description:
//
//	  Before you create an endpoint service, make sure that you have created a Server Load Balancer (SLB) instance that supports PrivateLink. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/174064.html).
//
//		- **CreateVpcEndpointService*	- is an asynchronous operation. After a request is sent, the system returns a request ID and an instance ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to query the status of the endpoint service.
//
//	    	- If the endpoint service is in the **Creating*	- state, the endpoint service is being created.
//
//	    	- If the endpoint service is in the **Active*	- state, the endpoint service is created.
//
// @param request - CreateVpcEndpointServiceRequest
//
// @return CreateVpcEndpointServiceResponse
func (client *Client) CreateVpcEndpointService(request *CreateVpcEndpointServiceRequest) (_result *CreateVpcEndpointServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcEndpointServiceResponse{}
	_body, _err := client.CreateVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an endpoint.
//
// Description:
//
//	  Before you delete an endpoint, you must delete the zones that are added to the endpoint.
//
//		- **DeleteVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to check whether the endpoint is deleted.
//
//	    	- If the endpoint is in the **Deleting*	- state, the endpoint is being deleted.
//
//	    	- If the endpoint cannot be queried, the endpoint is deleted.
//
// @param request - DeleteVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointResponse
func (client *Client) DeleteVpcEndpointWithOptions(request *DeleteVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  Before you delete an endpoint, you must delete the zones that are added to the endpoint.
//
//		- **DeleteVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to check whether the endpoint is deleted.
//
//	    	- If the endpoint is in the **Deleting*	- state, the endpoint is being deleted.
//
//	    	- If the endpoint cannot be queried, the endpoint is deleted.
//
// @param request - DeleteVpcEndpointRequest
//
// @return DeleteVpcEndpointResponse
func (client *Client) DeleteVpcEndpoint(request *DeleteVpcEndpointRequest) (_result *DeleteVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.DeleteVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an endpoint service.
//
// Description:
//
//	  Before you delete an endpoint service, you must disconnect the endpoint from the endpoint service and remove the service resources.
//
//		- **DeleteVpcEndpointService*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to check whether the endpoint service is deleted.
//
//	    	- If the endpoint service is in the **Deleting*	- state, the endpoint service is being deleted.
//
//	    	- If the endpoint service cannot be queried, the endpoint service is deleted.
//
//		- You cannot repeatedly call the **DeleteVpcEndpointService*	- operation to delete an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - DeleteVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcEndpointServiceResponse
func (client *Client) DeleteVpcEndpointServiceWithOptions(request *DeleteVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcEndpointServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  Before you delete an endpoint service, you must disconnect the endpoint from the endpoint service and remove the service resources.
//
//		- **DeleteVpcEndpointService*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/183542.html) operation to check whether the endpoint service is deleted.
//
//	    	- If the endpoint service is in the **Deleting*	- state, the endpoint service is being deleted.
//
//	    	- If the endpoint service cannot be queried, the endpoint service is deleted.
//
//		- You cannot repeatedly call the **DeleteVpcEndpointService*	- operation to delete an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - DeleteVpcEndpointServiceRequest
//
// @return DeleteVpcEndpointServiceResponse
func (client *Client) DeleteVpcEndpointService(request *DeleteVpcEndpointServiceRequest) (_result *DeleteVpcEndpointServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVpcEndpointServiceResponse{}
	_body, _err := client.DeleteVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available regions and zones.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available regions and zones.
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
// Queries a list of zones in a specified region.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of zones in a specified region.
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
// Removes a service resource from an endpoint service.
//
// Description:
//
//	  Before you remove a service resource from an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
//		- You cannot repeatedly call the **DetachResourceFromVpcEndpointService*	- operation to remove a service resource from an endpoint service within a specified period of time.
//
// @param request - DetachResourceFromVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachResourceFromVpcEndpointServiceResponse
func (client *Client) DetachResourceFromVpcEndpointServiceWithOptions(request *DetachResourceFromVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *DetachResourceFromVpcEndpointServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  Before you remove a service resource from an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
//		- You cannot repeatedly call the **DetachResourceFromVpcEndpointService*	- operation to remove a service resource from an endpoint service within a specified period of time.
//
// @param request - DetachResourceFromVpcEndpointServiceRequest
//
// @return DetachResourceFromVpcEndpointServiceResponse
func (client *Client) DetachResourceFromVpcEndpointService(request *DetachResourceFromVpcEndpointServiceRequest) (_result *DetachResourceFromVpcEndpointServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachResourceFromVpcEndpointServiceResponse{}
	_body, _err := client.DetachResourceFromVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates an endpoint from a security group.
//
// Description:
//
//	  **DetachSecurityGroupFromVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) to check whether the endpoint is disassociated from the security group.
//
//	    	- If the endpoint is in the **Pending*	- state, the endpoint is being disassociated from the security group.
//
//	    	- If you cannot query the endpoint in the security group, the endpoint is disassociated from the security group.
//
//		- You cannot repeatedly call the **DetachSecurityGroupFromVpcEndpoint*	- operation to disassociate an endpoint from a security group within a specified period of time.
//
// @param request - DetachSecurityGroupFromVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachSecurityGroupFromVpcEndpointResponse
func (client *Client) DetachSecurityGroupFromVpcEndpointWithOptions(request *DetachSecurityGroupFromVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *DetachSecurityGroupFromVpcEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  **DetachSecurityGroupFromVpcEndpoint*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpoints](https://help.aliyun.com/document_detail/183558.html) to check whether the endpoint is disassociated from the security group.
//
//	    	- If the endpoint is in the **Pending*	- state, the endpoint is being disassociated from the security group.
//
//	    	- If you cannot query the endpoint in the security group, the endpoint is disassociated from the security group.
//
//		- You cannot repeatedly call the **DetachSecurityGroupFromVpcEndpoint*	- operation to disassociate an endpoint from a security group within a specified period of time.
//
// @param request - DetachSecurityGroupFromVpcEndpointRequest
//
// @return DetachSecurityGroupFromVpcEndpointResponse
func (client *Client) DetachSecurityGroupFromVpcEndpoint(request *DetachSecurityGroupFromVpcEndpointRequest) (_result *DetachSecurityGroupFromVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachSecurityGroupFromVpcEndpointResponse{}
	_body, _err := client.DetachSecurityGroupFromVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rejects a connection request from an endpoint.
//
// Description:
//
//	  **DisableVpcEndpointConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//	    	- If the endpoint connection is in the **Disconnecting*	- state, the endpoint is being disconnected from the endpoint service.
//
//	    	- If the endpoint connection is in the **Disconnected*	- state, the endpoint is disconnected from the endpoint service.
//
//		- You cannot repeatedly call the **DisableVpcEndpointConnection*	- operation to allow an endpoint service to reject a connection request from an endpoint within a specified period of time.
//
// @param request - DisableVpcEndpointConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVpcEndpointConnectionResponse
func (client *Client) DisableVpcEndpointConnectionWithOptions(request *DisableVpcEndpointConnectionRequest, runtime *dara.RuntimeOptions) (_result *DisableVpcEndpointConnectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  **DisableVpcEndpointConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//	    	- If the endpoint connection is in the **Disconnecting*	- state, the endpoint is being disconnected from the endpoint service.
//
//	    	- If the endpoint connection is in the **Disconnected*	- state, the endpoint is disconnected from the endpoint service.
//
//		- You cannot repeatedly call the **DisableVpcEndpointConnection*	- operation to allow an endpoint service to reject a connection request from an endpoint within a specified period of time.
//
// @param request - DisableVpcEndpointConnectionRequest
//
// @return DisableVpcEndpointConnectionResponse
func (client *Client) DisableVpcEndpointConnection(request *DisableVpcEndpointConnectionRequest) (_result *DisableVpcEndpointConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableVpcEndpointConnectionResponse{}
	_body, _err := client.DisableVpcEndpointConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Closes connections in an endpoint zone.
//
// Description:
//
//	  You can call this operation only when the state of the endpoint is **Connected*	- and the state of the zone associated with the endpoint is **Connected*	- or **Migrated**.
//
//		- **DisableVpcEndpointZoneConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the status of the task.
//
//	    	- If the zone is in the **Disconnecting*	- state, the task is running.
//
//	    	- If the zone is in the **Disconnected*	- state, the task is successful.
//
//		- You cannot repeatedly call the **DisableVpcEndpointZoneConnection*	- operation to allow an endpoint service to reject a connection request from the endpoint in the zone within a specified period of time.
//
// @param request - DisableVpcEndpointZoneConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVpcEndpointZoneConnectionResponse
func (client *Client) DisableVpcEndpointZoneConnectionWithOptions(request *DisableVpcEndpointZoneConnectionRequest, runtime *dara.RuntimeOptions) (_result *DisableVpcEndpointZoneConnectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  You can call this operation only when the state of the endpoint is **Connected*	- and the state of the zone associated with the endpoint is **Connected*	- or **Migrated**.
//
//		- **DisableVpcEndpointZoneConnection*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to query the status of the task.
//
//	    	- If the zone is in the **Disconnecting*	- state, the task is running.
//
//	    	- If the zone is in the **Disconnected*	- state, the task is successful.
//
//		- You cannot repeatedly call the **DisableVpcEndpointZoneConnection*	- operation to allow an endpoint service to reject a connection request from the endpoint in the zone within a specified period of time.
//
// @param request - DisableVpcEndpointZoneConnectionRequest
//
// @return DisableVpcEndpointZoneConnectionResponse
func (client *Client) DisableVpcEndpointZoneConnection(request *DisableVpcEndpointZoneConnectionRequest) (_result *DisableVpcEndpointZoneConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableVpcEndpointZoneConnectionResponse{}
	_body, _err := client.DisableVpcEndpointZoneConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Accepts connection requests from an endpoint.
//
// Description:
//
//	  **EnableVpcEndpointConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//	    	- If the state is **Connecting**, the endpoint connection is being established.
//
//	    	- If the state is **Connected**, the endpoint connection is established.
//
//		- You cannot repeatedly call the **EnableVpcEndpointConnection*	- operation to allow an endpoint service of an Alibaba Cloud account to accept a connection request from an endpoint within a specified period of time.
//
// @param request - EnableVpcEndpointConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableVpcEndpointConnectionResponse
func (client *Client) EnableVpcEndpointConnectionWithOptions(request *EnableVpcEndpointConnectionRequest, runtime *dara.RuntimeOptions) (_result *EnableVpcEndpointConnectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts connection requests from an endpoint.
//
// Description:
//
//	  **EnableVpcEndpointConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointAttribute](https://help.aliyun.com/document_detail/183568.html) operation to query the state of the endpoint connection.
//
//	    	- If the state is **Connecting**, the endpoint connection is being established.
//
//	    	- If the state is **Connected**, the endpoint connection is established.
//
//		- You cannot repeatedly call the **EnableVpcEndpointConnection*	- operation to allow an endpoint service of an Alibaba Cloud account to accept a connection request from an endpoint within a specified period of time.
//
// @param request - EnableVpcEndpointConnectionRequest
//
// @return EnableVpcEndpointConnectionResponse
func (client *Client) EnableVpcEndpointConnection(request *EnableVpcEndpointConnectionRequest) (_result *EnableVpcEndpointConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableVpcEndpointConnectionResponse{}
	_body, _err := client.EnableVpcEndpointConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Allows connections to endpoint zones.
//
// Description:
//
//	  You can call this operation only when the state of the endpoint is **Connected*	- and the state of the associated zone is **Disconnected**.
//
//		- **EnableVpcEndpointZoneConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the endpoint service accepts a connection request from the endpoint in the associated zone.
//
//	    	- If the zone is in the **Connecting*	- state, the endpoint service is accepting the connection request from the endpoint.
//
//	    	- If the zone is in the **Connected*	- state, the endpoint service has accepted the connection request from the endpoint.
//
//		- You cannot repeatedly call the **EnableVpcEndpointZoneConnection*	- operation to allow an endpoint service to accept a connection request from an endpoint in the associated zone within a specified period of time.
//
// @param request - EnableVpcEndpointZoneConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableVpcEndpointZoneConnectionResponse
func (client *Client) EnableVpcEndpointZoneConnectionWithOptions(request *EnableVpcEndpointZoneConnectionRequest, runtime *dara.RuntimeOptions) (_result *EnableVpcEndpointZoneConnectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  You can call this operation only when the state of the endpoint is **Connected*	- and the state of the associated zone is **Disconnected**.
//
//		- **EnableVpcEndpointZoneConnection*	- is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the endpoint service accepts a connection request from the endpoint in the associated zone.
//
//	    	- If the zone is in the **Connecting*	- state, the endpoint service is accepting the connection request from the endpoint.
//
//	    	- If the zone is in the **Connected*	- state, the endpoint service has accepted the connection request from the endpoint.
//
//		- You cannot repeatedly call the **EnableVpcEndpointZoneConnection*	- operation to allow an endpoint service to accept a connection request from an endpoint in the associated zone within a specified period of time.
//
// @param request - EnableVpcEndpointZoneConnectionRequest
//
// @return EnableVpcEndpointZoneConnectionResponse
func (client *Client) EnableVpcEndpointZoneConnection(request *EnableVpcEndpointZoneConnectionRequest) (_result *EnableVpcEndpointZoneConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableVpcEndpointZoneConnectionResponse{}
	_body, _err := client.EnableVpcEndpointZoneConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of an endpoint.
//
// @param request - GetVpcEndpointAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcEndpointAttributeResponse
func (client *Client) GetVpcEndpointAttributeWithOptions(request *GetVpcEndpointAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetVpcEndpointAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of an endpoint.
//
// @param request - GetVpcEndpointAttributeRequest
//
// @return GetVpcEndpointAttributeResponse
func (client *Client) GetVpcEndpointAttribute(request *GetVpcEndpointAttributeRequest) (_result *GetVpcEndpointAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVpcEndpointAttributeResponse{}
	_body, _err := client.GetVpcEndpointAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of an endpoint service.
//
// @param request - GetVpcEndpointServiceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcEndpointServiceAttributeResponse
func (client *Client) GetVpcEndpointServiceAttributeWithOptions(request *GetVpcEndpointServiceAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetVpcEndpointServiceAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of an endpoint service.
//
// @param request - GetVpcEndpointServiceAttributeRequest
//
// @return GetVpcEndpointServiceAttributeResponse
func (client *Client) GetVpcEndpointServiceAttribute(request *GetVpcEndpointServiceAttributeRequest) (_result *GetVpcEndpointServiceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVpcEndpointServiceAttributeResponse{}
	_body, _err := client.GetVpcEndpointServiceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to resources.
//
// Description:
//
//	  You must specify **ResourceId.N*	- or **Tag.N*	- in the request to specify the object that you want to query.
//
//		- **Tag.N*	- is a resource tag that consists of a key-value pair (Tag.N.Key and Tag.N.Value). If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
//		- If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
//		- If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  You must specify **ResourceId.N*	- or **Tag.N*	- in the request to specify the object that you want to query.
//
//		- **Tag.N*	- is a resource tag that consists of a key-value pair (Tag.N.Key and Tag.N.Value). If you specify only **Tag.N.Key**, all tag values that are associated with the specified key are returned. If you specify only **Tag.N.Value**, an error message is returned.
//
//		- If you specify **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
//		- If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
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
// Queries endpoint connections.
//
// @param request - ListVpcEndpointConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointConnectionsResponse
func (client *Client) ListVpcEndpointConnectionsWithOptions(request *ListVpcEndpointConnectionsRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointConnectionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries endpoint connections.
//
// @param request - ListVpcEndpointConnectionsRequest
//
// @return ListVpcEndpointConnectionsResponse
func (client *Client) ListVpcEndpointConnections(request *ListVpcEndpointConnectionsRequest) (_result *ListVpcEndpointConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcEndpointConnectionsResponse{}
	_body, _err := client.ListVpcEndpointConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security groups that are associated with an endpoint.
//
// @param request - ListVpcEndpointSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointSecurityGroupsResponse
func (client *Client) ListVpcEndpointSecurityGroupsWithOptions(request *ListVpcEndpointSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointSecurityGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security groups that are associated with an endpoint.
//
// @param request - ListVpcEndpointSecurityGroupsRequest
//
// @return ListVpcEndpointSecurityGroupsResponse
func (client *Client) ListVpcEndpointSecurityGroups(request *ListVpcEndpointSecurityGroupsRequest) (_result *ListVpcEndpointSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcEndpointSecurityGroupsResponse{}
	_body, _err := client.ListVpcEndpointSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListVpcEndpointServiceResourcesWithOptions(request *ListVpcEndpointServiceResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointServiceResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListVpcEndpointServiceResourcesResponse
func (client *Client) ListVpcEndpointServiceResources(request *ListVpcEndpointServiceResourcesRequest) (_result *ListVpcEndpointServiceResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcEndpointServiceResourcesResponse{}
	_body, _err := client.ListVpcEndpointServiceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListVpcEndpointServiceUsersWithOptions(request *ListVpcEndpointServiceUsersRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointServiceUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListVpcEndpointServiceUsersResponse
func (client *Client) ListVpcEndpointServiceUsers(request *ListVpcEndpointServiceUsersRequest) (_result *ListVpcEndpointServiceUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcEndpointServiceUsersResponse{}
	_body, _err := client.ListVpcEndpointServiceUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of endpoint services.
//
// @param request - ListVpcEndpointServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServicesResponse
func (client *Client) ListVpcEndpointServicesWithOptions(request *ListVpcEndpointServicesRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointServicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of endpoint services.
//
// @param request - ListVpcEndpointServicesRequest
//
// @return ListVpcEndpointServicesResponse
func (client *Client) ListVpcEndpointServices(request *ListVpcEndpointServicesRequest) (_result *ListVpcEndpointServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcEndpointServicesResponse{}
	_body, _err := client.ListVpcEndpointServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of endpoint services that can be associated with the endpoint created within the current account.
//
// @param request - ListVpcEndpointServicesByEndUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointServicesByEndUserResponse
func (client *Client) ListVpcEndpointServicesByEndUserWithOptions(request *ListVpcEndpointServicesByEndUserRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointServicesByEndUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of endpoint services that can be associated with the endpoint created within the current account.
//
// @param request - ListVpcEndpointServicesByEndUserRequest
//
// @return ListVpcEndpointServicesByEndUserResponse
func (client *Client) ListVpcEndpointServicesByEndUser(request *ListVpcEndpointServicesByEndUserRequest) (_result *ListVpcEndpointServicesByEndUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcEndpointServicesByEndUserResponse{}
	_body, _err := client.ListVpcEndpointServicesByEndUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the zones of an endpoint.
//
// @param request - ListVpcEndpointZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointZonesResponse
func (client *Client) ListVpcEndpointZonesWithOptions(request *ListVpcEndpointZonesRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointZonesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones of an endpoint.
//
// @param request - ListVpcEndpointZonesRequest
//
// @return ListVpcEndpointZonesResponse
func (client *Client) ListVpcEndpointZones(request *ListVpcEndpointZonesRequest) (_result *ListVpcEndpointZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcEndpointZonesResponse{}
	_body, _err := client.ListVpcEndpointZonesWithOptions(request, runtime)
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
// @param request - ListVpcEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcEndpointsResponse
func (client *Client) ListVpcEndpointsWithOptions(request *ListVpcEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
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
// @param request - ListVpcEndpointsRequest
//
// @return ListVpcEndpointsResponse
func (client *Client) ListVpcEndpoints(request *ListVpcEndpointsRequest) (_result *ListVpcEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.ListVpcEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OpenPrivateLinkServiceWithOptions(request *OpenPrivateLinkServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenPrivateLinkServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OpenPrivateLinkServiceResponse
func (client *Client) OpenPrivateLinkService(request *OpenPrivateLinkServiceRequest) (_result *OpenPrivateLinkServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenPrivateLinkServiceResponse{}
	_body, _err := client.OpenPrivateLinkServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an account ID from the whitelist of an endpoint service.
//
// Description:
//
//	  Before you remove an account ID from the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
//		- You cannot repeatedly call the **RemoveUserFromVpcEndpointService*	- operation to remove the ID of an Alibaba Cloud account from the whitelist of an endpoint service within a specified period of time.
//
// @param request - RemoveUserFromVpcEndpointServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromVpcEndpointServiceResponse
func (client *Client) RemoveUserFromVpcEndpointServiceWithOptions(request *RemoveUserFromVpcEndpointServiceRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserFromVpcEndpointServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  Before you remove an account ID from the whitelist of an endpoint service, make sure that the endpoint service is in the **Active*	- state. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to query the status of the endpoint service.
//
//		- You cannot repeatedly call the **RemoveUserFromVpcEndpointService*	- operation to remove the ID of an Alibaba Cloud account from the whitelist of an endpoint service within a specified period of time.
//
// @param request - RemoveUserFromVpcEndpointServiceRequest
//
// @return RemoveUserFromVpcEndpointServiceResponse
func (client *Client) RemoveUserFromVpcEndpointService(request *RemoveUserFromVpcEndpointServiceRequest) (_result *RemoveUserFromVpcEndpointServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUserFromVpcEndpointServiceResponse{}
	_body, _err := client.RemoveUserFromVpcEndpointServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a zone of an endpoint.
//
// Description:
//
//	  **RemoveZoneFromVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the zone of the endpoint is deleted.
//
//	    	- If the zone of the endpoint is in the **Deleting*	- state, the zone is being deleted.
//
//	    	- If the zone cannot be queried, the zone is deleted.
//
//		- You cannot repeatedly call the **RemoveZoneFromVpcEndpoint*	- operation to delete a zone of an endpoint within a specified period of time.
//
// @param request - RemoveZoneFromVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveZoneFromVpcEndpointResponse
func (client *Client) RemoveZoneFromVpcEndpointWithOptions(request *RemoveZoneFromVpcEndpointRequest, runtime *dara.RuntimeOptions) (_result *RemoveZoneFromVpcEndpointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//	  **RemoveZoneFromVpcEndpoint*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListVpcEndpointZones](https://help.aliyun.com/document_detail/183560.html) operation to check whether the zone of the endpoint is deleted.
//
//	    	- If the zone of the endpoint is in the **Deleting*	- state, the zone is being deleted.
//
//	    	- If the zone cannot be queried, the zone is deleted.
//
//		- You cannot repeatedly call the **RemoveZoneFromVpcEndpoint*	- operation to delete a zone of an endpoint within a specified period of time.
//
// @param request - RemoveZoneFromVpcEndpointRequest
//
// @return RemoveZoneFromVpcEndpointResponse
func (client *Client) RemoveZoneFromVpcEndpoint(request *RemoveZoneFromVpcEndpointRequest) (_result *RemoveZoneFromVpcEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveZoneFromVpcEndpointResponse{}
	_body, _err := client.RemoveZoneFromVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Removes tags from one or more endpoints or endpoint services at a time.
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Modifies the attributes of an endpoint.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointAttribute*	- operation to modify the attributes of an endpoint that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointAttributeResponse
func (client *Client) UpdateVpcEndpointAttributeWithOptions(request *UpdateVpcEndpointAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an endpoint.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointAttribute*	- operation to modify the attributes of an endpoint that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointAttributeRequest
//
// @return UpdateVpcEndpointAttributeResponse
func (client *Client) UpdateVpcEndpointAttribute(request *UpdateVpcEndpointAttributeRequest) (_result *UpdateVpcEndpointAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVpcEndpointAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an endpoint connection.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointConnectionAttribute*	- operation to modify the bandwidth of an endpoint connection that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointConnectionAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointConnectionAttributeResponse
func (client *Client) UpdateVpcEndpointConnectionAttributeWithOptions(request *UpdateVpcEndpointConnectionAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointConnectionAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an endpoint connection.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointConnectionAttribute*	- operation to modify the bandwidth of an endpoint connection that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointConnectionAttributeRequest
//
// @return UpdateVpcEndpointConnectionAttributeResponse
func (client *Client) UpdateVpcEndpointConnectionAttribute(request *UpdateVpcEndpointConnectionAttributeRequest) (_result *UpdateVpcEndpointConnectionAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVpcEndpointConnectionAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointConnectionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an endpoint service.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointServiceAttribute*	- operation to modify the attributes of an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointServiceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointServiceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceAttributeWithOptions(request *UpdateVpcEndpointServiceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointServiceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.ConnectBandwidth) {
		query["ConnectBandwidth"] = request.ConnectBandwidth
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an endpoint service.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointServiceAttribute*	- operation to modify the attributes of an endpoint service that belongs to an Alibaba Cloud account within a specified period of time.
//
// @param request - UpdateVpcEndpointServiceAttributeRequest
//
// @return UpdateVpcEndpointServiceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceAttribute(request *UpdateVpcEndpointServiceAttributeRequest) (_result *UpdateVpcEndpointServiceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVpcEndpointServiceAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointServiceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a service resource that is added to an endpoint service.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointServiceResourceAttribute*	- operation to modify the attributes of a service resource that is added to an endpoint service within a specified period of time.
//
// @param request - UpdateVpcEndpointServiceResourceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointServiceResourceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceResourceAttributeWithOptions(request *UpdateVpcEndpointServiceResourceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointServiceResourceAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a service resource that is added to an endpoint service.
//
// Description:
//
// You cannot repeatedly call the **UpdateVpcEndpointServiceResourceAttribute*	- operation to modify the attributes of a service resource that is added to an endpoint service within a specified period of time.
//
// @param request - UpdateVpcEndpointServiceResourceAttributeRequest
//
// @return UpdateVpcEndpointServiceResourceAttributeResponse
func (client *Client) UpdateVpcEndpointServiceResourceAttribute(request *UpdateVpcEndpointServiceResourceAttributeRequest) (_result *UpdateVpcEndpointServiceResourceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVpcEndpointServiceResourceAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointServiceResourceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
//   - If the endpoint connection is in the **Disconnected*	- state, you can manually allocate the service resource in the zone.
//
//   - If the endpoint connection is in the **Connected*	- state, you can manually migrate the service resource in the zone. In this case, the connection might be interrupted.
//
//   - **UpdateVpcEndpointZoneConnectionResourceAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to check whether the service resource is modified.
//
//   - If the endpoint service is in the **Pending*	- state, the service resource is being modified.
//
//   - If the endpoint service is in the **Active*	- state, the service resource is modified.
//
//   - You cannot repeatedly call the **UpdateVpcEndpointZoneConnectionResourceAttribute*	- operation to modify a service resource in the zone to which an endpoint connection belongs within a specified period of time.
//
// @param request - UpdateVpcEndpointZoneConnectionResourceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVpcEndpointZoneConnectionResourceAttributeResponse
func (client *Client) UpdateVpcEndpointZoneConnectionResourceAttributeWithOptions(request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
//   - If the endpoint connection is in the **Disconnected*	- state, you can manually allocate the service resource in the zone.
//
//   - If the endpoint connection is in the **Connected*	- state, you can manually migrate the service resource in the zone. In this case, the connection might be interrupted.
//
//   - **UpdateVpcEndpointZoneConnectionResourceAttribute*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [GetVpcEndpointServiceAttribute](https://help.aliyun.com/document_detail/469330.html) operation to check whether the service resource is modified.
//
//   - If the endpoint service is in the **Pending*	- state, the service resource is being modified.
//
//   - If the endpoint service is in the **Active*	- state, the service resource is modified.
//
//   - You cannot repeatedly call the **UpdateVpcEndpointZoneConnectionResourceAttribute*	- operation to modify a service resource in the zone to which an endpoint connection belongs within a specified period of time.
//
// @param request - UpdateVpcEndpointZoneConnectionResourceAttributeRequest
//
// @return UpdateVpcEndpointZoneConnectionResourceAttributeResponse
func (client *Client) UpdateVpcEndpointZoneConnectionResourceAttribute(request *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) (_result *UpdateVpcEndpointZoneConnectionResourceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVpcEndpointZoneConnectionResourceAttributeResponse{}
	_body, _err := client.UpdateVpcEndpointZoneConnectionResourceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
