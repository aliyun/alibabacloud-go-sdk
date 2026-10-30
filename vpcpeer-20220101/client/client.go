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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("vpcpeer"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Accepts a VPC peering connection request.
//
// Description:
//
// - For a cross-account VPC peering connection, the VPC peering connection instance is activated only after the acceptor accepts the connection request.
//
// - The **AcceptVpcPeerConnection*	- operation is asynchronous. After you call this operation, the system returns a **RequestId**, but the VPC peering connection instance is not immediately activated. The activation task is still running in the background. You can call [GetVpcPeerConnectionAttribute](https://help.aliyun.com/document_detail/426100.html) to query the status of the VPC peering connection instance.
//
//   - If the VPC peering connection instance is in the **Updating*	- state, the instance is being activated.
//
//   - If the VPC peering connection instance is in the **Activated*	- state, the instance is activated.
//
// - You cannot call the **AcceptVpcPeerConnection*	- operation to concurrently accept requests for the same VPC peering connection instance.
//
// @param request - AcceptVpcPeerConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptVpcPeerConnectionResponse
func (client *Client) AcceptVpcPeerConnectionWithOptions(request *AcceptVpcPeerConnectionRequest, runtime *dara.RuntimeOptions) (_result *AcceptVpcPeerConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		body["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptVpcPeerConnection"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Accepts a VPC peering connection request.
//
// Description:
//
// - For a cross-account VPC peering connection, the VPC peering connection instance is activated only after the acceptor accepts the connection request.
//
// - The **AcceptVpcPeerConnection*	- operation is asynchronous. After you call this operation, the system returns a **RequestId**, but the VPC peering connection instance is not immediately activated. The activation task is still running in the background. You can call [GetVpcPeerConnectionAttribute](https://help.aliyun.com/document_detail/426100.html) to query the status of the VPC peering connection instance.
//
//   - If the VPC peering connection instance is in the **Updating*	- state, the instance is being activated.
//
//   - If the VPC peering connection instance is in the **Activated*	- state, the instance is activated.
//
// - You cannot call the **AcceptVpcPeerConnection*	- operation to concurrently accept requests for the same VPC peering connection instance.
//
// @param request - AcceptVpcPeerConnectionRequest
//
// @return AcceptVpcPeerConnectionResponse
func (client *Client) AcceptVpcPeerConnection(request *AcceptVpcPeerConnectionRequest) (_result *AcceptVpcPeerConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptVpcPeerConnectionResponse{}
	_body, _err := client.AcceptVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the CreateVpcPeerConnection operation to create a VPC peering connection.
//
// Description:
//
// Before you create a VPC peering connection, take note of the following items:
//
// - **CreateVpcPeerConnection*	- is an asynchronous operation. After you call this operation, the system returns an **InstanceId**, but the VPC peering connection is not created yet and the creation task is still running in the background. You can call [GetVpcPeerConnectionAttribute](https://help.aliyun.com/document_detail/426095.html) to query the status of the VPC peering connection.
//
//   - If the VPC peering connection is in the **Creating*	- state, the VPC peering connection is being created.
//
//   - If the VPC peering connection is in the **Activated*	- state, the VPC peering connection is created.
//
//   - If the VPC peering connection is in the **Accepting*	- state, the VPC peering connection is created across accounts and must be accepted by the acceptor.
//
// - You cannot call the **CreateVpcPeerConnection*	- operation to concurrently create VPC peering connections in the same VPC.
//
// <props="china">When you create a VPC peering connection, the system automatically activates Cloud Data Transfer (CDT) for non-cross-border data transfer and cross-border data transfer.
//
// <props="intl">When you create a VPC peering connection, the system automatically activates Cloud Data Transfer (CDT) for non-cross-border data transfer.
//
// @param request - CreateVpcPeerConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVpcPeerConnectionResponse
func (client *Client) CreateVpcPeerConnectionWithOptions(request *CreateVpcPeerConnectionRequest, runtime *dara.RuntimeOptions) (_result *CreateVpcPeerConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LinkType) {
		query["LinkType"] = request.LinkType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptingAliUid) {
		body["AcceptingAliUid"] = request.AcceptingAliUid
	}

	if !dara.IsNil(request.AcceptingRegionId) {
		body["AcceptingRegionId"] = request.AcceptingRegionId
	}

	if !dara.IsNil(request.AcceptingVpcId) {
		body["AcceptingVpcId"] = request.AcceptingVpcId
	}

	if !dara.IsNil(request.Bandwidth) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableJumboFrame) {
		body["EnableJumboFrame"] = request.EnableJumboFrame
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVpcPeerConnection"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the CreateVpcPeerConnection operation to create a VPC peering connection.
//
// Description:
//
// Before you create a VPC peering connection, take note of the following items:
//
// - **CreateVpcPeerConnection*	- is an asynchronous operation. After you call this operation, the system returns an **InstanceId**, but the VPC peering connection is not created yet and the creation task is still running in the background. You can call [GetVpcPeerConnectionAttribute](https://help.aliyun.com/document_detail/426095.html) to query the status of the VPC peering connection.
//
//   - If the VPC peering connection is in the **Creating*	- state, the VPC peering connection is being created.
//
//   - If the VPC peering connection is in the **Activated*	- state, the VPC peering connection is created.
//
//   - If the VPC peering connection is in the **Accepting*	- state, the VPC peering connection is created across accounts and must be accepted by the acceptor.
//
// - You cannot call the **CreateVpcPeerConnection*	- operation to concurrently create VPC peering connections in the same VPC.
//
// <props="china">When you create a VPC peering connection, the system automatically activates Cloud Data Transfer (CDT) for non-cross-border data transfer and cross-border data transfer.
//
// <props="intl">When you create a VPC peering connection, the system automatically activates Cloud Data Transfer (CDT) for non-cross-border data transfer.
//
// @param request - CreateVpcPeerConnectionRequest
//
// @return CreateVpcPeerConnectionResponse
func (client *Client) CreateVpcPeerConnection(request *CreateVpcPeerConnectionRequest) (_result *CreateVpcPeerConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVpcPeerConnectionResponse{}
	_body, _err := client.CreateVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a VPC peering connection instance.
//
// Description:
//
// - You can delete a VPC peering connection instance. After the instance is deleted, your services will be interrupted. Perform this operation only when it does not affect your services.
//
//   - If you choose to force delete the instance, the system deletes the route entry that points to the VPC peering connection instance from the VPC routing table.
//
//   - If you do not choose to force delete the instance, the system does not delete the route entry that points to the VPC peering connection instance from the VPC routing table. You must manually delete the route entry.
//
// - The **DeleteVpcPeerConnection*	- operation is asynchronous. After you call the operation, the system returns a **RequestId**, but the VPC peering connection instance is not immediately deleted. The deletion task is still running in the background. You can call the [GetVpcPeerConnectionAttribute](https://help.aliyun.com/document_detail/2523294.html) operation to query the status of the VPC peering connection instance.
//
//   - If the VPC peering connection instance is in the **Deleting*	- state, the instance is being deleted.
//
//   - If the VPC peering connection instance is in the **Deleted*	- state, the instance is deleted.
//
// - You cannot call the **DeleteVpcPeerConnection*	- operation to concurrently delete the same VPC peering connection instance.
//
// @param request - DeleteVpcPeerConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVpcPeerConnectionResponse
func (client *Client) DeleteVpcPeerConnectionWithOptions(request *DeleteVpcPeerConnectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteVpcPeerConnectionResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVpcPeerConnection"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VPC peering connection instance.
//
// Description:
//
// - You can delete a VPC peering connection instance. After the instance is deleted, your services will be interrupted. Perform this operation only when it does not affect your services.
//
//   - If you choose to force delete the instance, the system deletes the route entry that points to the VPC peering connection instance from the VPC routing table.
//
//   - If you do not choose to force delete the instance, the system does not delete the route entry that points to the VPC peering connection instance from the VPC routing table. You must manually delete the route entry.
//
// - The **DeleteVpcPeerConnection*	- operation is asynchronous. After you call the operation, the system returns a **RequestId**, but the VPC peering connection instance is not immediately deleted. The deletion task is still running in the background. You can call the [GetVpcPeerConnectionAttribute](https://help.aliyun.com/document_detail/2523294.html) operation to query the status of the VPC peering connection instance.
//
//   - If the VPC peering connection instance is in the **Deleting*	- state, the instance is being deleted.
//
//   - If the VPC peering connection instance is in the **Deleted*	- state, the instance is deleted.
//
// - You cannot call the **DeleteVpcPeerConnection*	- operation to concurrently delete the same VPC peering connection instance.
//
// @param request - DeleteVpcPeerConnectionRequest
//
// @return DeleteVpcPeerConnectionResponse
func (client *Client) DeleteVpcPeerConnection(request *DeleteVpcPeerConnectionRequest) (_result *DeleteVpcPeerConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVpcPeerConnectionResponse{}
	_body, _err := client.DeleteVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified VPC peering connection.
//
// @param request - GetVpcPeerConnectionAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVpcPeerConnectionAttributeResponse
func (client *Client) GetVpcPeerConnectionAttributeWithOptions(request *GetVpcPeerConnectionAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetVpcPeerConnectionAttributeResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		body["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVpcPeerConnectionAttribute"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVpcPeerConnectionAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified VPC peering connection.
//
// @param request - GetVpcPeerConnectionAttributeRequest
//
// @return GetVpcPeerConnectionAttributeResponse
func (client *Client) GetVpcPeerConnectionAttribute(request *GetVpcPeerConnectionAttributeRequest) (_result *GetVpcPeerConnectionAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVpcPeerConnectionAttributeResponse{}
	_body, _err := client.GetVpcPeerConnectionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags bound to VPC peering connections.
//
// Description:
//
// - Specify at least one of the **ResourceId.N*	- and **Tag.N*	- (**Tag.N.Key*	- and **Tag.N.Value**) parameters in the request to determine the resources to query.
//
// - **Tag.N*	- is the tag of the resource and consists of a key-value pair. If you specify only **Tag.N.Key**, all tag values associated with the tag key are returned. If you specify only **Tag.N.Value**, an error is reported.
//
// - If you specify both **Tag.N*	- and **ResourceId.N*	- to filter tags, the **ResourceId.N*	- parameter must match all the specified tag key-value pairs.
//
// - If you specify multiple tag key-value pairs, the returned resources contain all the specified key-value pairs.
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
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
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
		Version:     dara.String("2022-01-01"),
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
// Queries the tags bound to VPC peering connections.
//
// Description:
//
// - Specify at least one of the **ResourceId.N*	- and **Tag.N*	- (**Tag.N.Key*	- and **Tag.N.Value**) parameters in the request to determine the resources to query.
//
// - **Tag.N*	- is the tag of the resource and consists of a key-value pair. If you specify only **Tag.N.Key**, all tag values associated with the tag key are returned. If you specify only **Tag.N.Value**, an error is reported.
//
// - If you specify both **Tag.N*	- and **ResourceId.N*	- to filter tags, the **ResourceId.N*	- parameter must match all the specified tag key-value pairs.
//
// - If you specify multiple tag key-value pairs, the returned resources contain all the specified key-value pairs.
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
// Queries the details of VPC peering connections.
//
// @param tmpReq - ListVpcPeerConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcPeerConnectionsResponse
func (client *Client) ListVpcPeerConnectionsWithOptions(tmpReq *ListVpcPeerConnectionsRequest, runtime *dara.RuntimeOptions) (_result *ListVpcPeerConnectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListVpcPeerConnectionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VpcId) {
		request.VpcIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VpcId, dara.String("VpcId"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpcIdShrink) {
		body["VpcId"] = request.VpcIdShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpcPeerConnections"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcPeerConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of VPC peering connections.
//
// @param request - ListVpcPeerConnectionsRequest
//
// @return ListVpcPeerConnectionsResponse
func (client *Client) ListVpcPeerConnections(request *ListVpcPeerConnectionsRequest) (_result *ListVpcPeerConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVpcPeerConnectionsResponse{}
	_body, _err := client.ListVpcPeerConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name or description of a VPC peering connection.
//
// Description:
//
// - **ModifyVpcPeerConnection*	- is an asynchronous operation. The system returns a **RequestId**, but the VPC peering connection is not immediately modified. The modification task continues in the background. You can call [GetVpcPeerConnectionAttribute](https://help.aliyun.com/document_detail/426100.html) to query the status of the VPC peering connection.
//
//   - If the VPC peering connection is in the **Updating*	- state, it is being modified.
//
//   - If the VPC peering connection is in the **Activated*	- state, the modification is complete.
//
// - **ModifyVpcPeerConnection*	- does not support concurrent modifications to the name and description of the same VPC peering connection.
//
// @param request - ModifyVpcPeerConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcPeerConnectionResponse
func (client *Client) ModifyVpcPeerConnectionWithOptions(request *ModifyVpcPeerConnectionRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcPeerConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LinkType) {
		query["LinkType"] = request.LinkType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableJumboFrame) {
		body["EnableJumboFrame"] = request.EnableJumboFrame
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcPeerConnection"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name or description of a VPC peering connection.
//
// Description:
//
// - **ModifyVpcPeerConnection*	- is an asynchronous operation. The system returns a **RequestId**, but the VPC peering connection is not immediately modified. The modification task continues in the background. You can call [GetVpcPeerConnectionAttribute](https://help.aliyun.com/document_detail/426100.html) to query the status of the VPC peering connection.
//
//   - If the VPC peering connection is in the **Updating*	- state, it is being modified.
//
//   - If the VPC peering connection is in the **Activated*	- state, the modification is complete.
//
// - **ModifyVpcPeerConnection*	- does not support concurrent modifications to the name and description of the same VPC peering connection.
//
// @param request - ModifyVpcPeerConnectionRequest
//
// @return ModifyVpcPeerConnectionResponse
func (client *Client) ModifyVpcPeerConnection(request *ModifyVpcPeerConnectionRequest) (_result *ModifyVpcPeerConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcPeerConnectionResponse{}
	_body, _err := client.ModifyVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the resource group of a VPC peering connection.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource group of a VPC peering connection.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rejects a request for a VPC peering connection.
//
// Description:
//
// - For a cross-account VPC peering connection instance, the acceptor can reject the connection request. A VPC peering connection instance whose connection request is rejected enters the **Rejected*	- state.
//
// - **RejectVpcPeerConnection*	- does not support concurrent requests to reject the same peering connection instance.
//
// @param request - RejectVpcPeerConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectVpcPeerConnectionResponse
func (client *Client) RejectVpcPeerConnectionWithOptions(request *RejectVpcPeerConnectionRequest, runtime *dara.RuntimeOptions) (_result *RejectVpcPeerConnectionResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		body["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectVpcPeerConnection"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectVpcPeerConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rejects a request for a VPC peering connection.
//
// Description:
//
// - For a cross-account VPC peering connection instance, the acceptor can reject the connection request. A VPC peering connection instance whose connection request is rejected enters the **Rejected*	- state.
//
// - **RejectVpcPeerConnection*	- does not support concurrent requests to reject the same peering connection instance.
//
// @param request - RejectVpcPeerConnectionRequest
//
// @return RejectVpcPeerConnectionResponse
func (client *Client) RejectVpcPeerConnection(request *RejectVpcPeerConnectionRequest) (_result *RejectVpcPeerConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectVpcPeerConnectionResponse{}
	_body, _err := client.RejectVpcPeerConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and attaches tags to one or more specified VPC peering connections.
//
// Description:
//
// Tags are labels that you assign to instances. Instance tag consists of a key-value pair. Take note of the following limits on tags:
//
// - The key of instance tag on an instance must be unique.
//
// - Tags must be attached to instances. Unattached tags are not supported.
//
// - Tag information is not shared across regions.
//
//	For example, tags created in the China (Hangzhou) region are not visible in the China (Shanghai) region.
//
// - Within the same account and region, tag information is shared across different VPC peering connections.
//
//	For example, if a tag is attached to a VPC peering connection within the same account and region, you can directly select and attach this tag to other VPC peering connections on the Edit Tags page without manually entering the key-value pair. You can modify the key and value of a tag, or delete the tag from an instance at any time. If an instance is deleted, all tags attached to the instance are also deleted.
//
// - You can attach up to 20 tags to a single instance. Before attaching tags, Alibaba Cloud verifies the number of existing tags on the resource. If the limit is exceeded, an error message is returned.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-01-01"),
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
// Creates and attaches tags to one or more specified VPC peering connections.
//
// Description:
//
// Tags are labels that you assign to instances. Instance tag consists of a key-value pair. Take note of the following limits on tags:
//
// - The key of instance tag on an instance must be unique.
//
// - Tags must be attached to instances. Unattached tags are not supported.
//
// - Tag information is not shared across regions.
//
//	For example, tags created in the China (Hangzhou) region are not visible in the China (Shanghai) region.
//
// - Within the same account and region, tag information is shared across different VPC peering connections.
//
//	For example, if a tag is attached to a VPC peering connection within the same account and region, you can directly select and attach this tag to other VPC peering connections on the Edit Tags page without manually entering the key-value pair. You can modify the key and value of a tag, or delete the tag from an instance at any time. If an instance is deleted, all tags attached to the instance are also deleted.
//
// - You can attach up to 20 tags to a single instance. Before attaching tags, Alibaba Cloud verifies the number of existing tags on the resource. If the limit is exceeded, an error message is returned.
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
// Removes tags from specified VPC peering connections.
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnTagResources"),
		Version:     dara.String("2022-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from specified VPC peering connections.
//
// @param request - UnTagResourcesRequest
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
