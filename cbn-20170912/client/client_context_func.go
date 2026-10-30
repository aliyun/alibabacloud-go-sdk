// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Activates a flow log. After the activate flow log operation is complete, it starts to capture traffic of the specified resources.
//
// Description:
//
// - After a flow log is created, it is in the Active state by default. If you stop activate flow log, you can call this operation to restart it.
//
// - The `ActiveFlowLog` operation is asynchronous. After you call this operation, the system returns a **RequestId**, but activate flow log is not fully activated. The activation task is still running in the background. You can call the `DescribeFlowlogs` operation to query the status of activate flow log.
//
//   - If activate flow log is in the **Modifying*	- state, activate flow log is being activated. In this state, you can only perform query operations.
//
//   - If activate flow log is in the **Active*	- state, activate flow log is activated.
//
// @param request - ActiveFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveFlowLogResponse
func (client *Client) ActiveFlowLogWithContext(ctx context.Context, request *ActiveFlowLogRequest, runtime *dara.RuntimeOptions) (_result *ActiveFlowLogResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
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
		Action:      dara.String("ActiveFlowLog"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveFlowLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a traffic classification rule to a traffic marking policy.
//
// Description:
//
// The **AddTrafficMatchRuleToTrafficMarkingPolicy*	- API operation is asynchronous. After you send the request, the system returns a **RequestId**, but the traffic classification rule is not immediately created. The creation task is still running in the background. You can call the **ListTrafficMarkingPolicies*	- API operation to query the status of the traffic classification rule.
//
// - If the traffic classification rule is in the **Creating*	- state, the rule is being created. In this state, you can only query the rule and cannot perform other operations.
//
// - If the traffic classification rule is in the **Active*	- state, the rule is created.
//
// @param request - AddTrafficMatchRuleToTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTrafficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) AddTrafficMatchRuleToTrafficMarkingPolicyWithContext(ctx context.Context, request *AddTrafficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddTrafficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyId) {
		query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	}

	if !dara.IsNil(request.TrafficMatchRules) {
		query["TrafficMatchRules"] = request.TrafficMatchRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTrafficMatchRuleToTrafficMarkingPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTrafficMatchRuleToTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AddTraficMatchRuleToTrafficMarkingPolicy is deprecated, please use Cbn::2017-09-12::AddTrafficMatchRuleToTrafficMarkingPolicy instead.
//
// Summary:
//
// Adds a traffic classification rule to a traffic marking policy.
//
// Description:
//
// ### Precautions
//
// The **AddTraficMatchRuleToTrafficMarkingPolicy*	- operation is deprecated and will be retired soon. To add a traffic classification rule to a traffic marking policy, use the [AddTrafficMatchRuleToTrafficMarkingPolicy](https://help.aliyun.com/document_detail/427602.html) operation. This API document is no longer maintained.
//
// @param request - AddTraficMatchRuleToTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTraficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) AddTraficMatchRuleToTrafficMarkingPolicyWithContext(ctx context.Context, request *AddTraficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddTraficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyId) {
		query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	}

	if !dara.IsNil(request.TrafficMatchRules) {
		query["TrafficMatchRules"] = request.TrafficMatchRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTraficMatchRuleToTrafficMarkingPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTraficMatchRuleToTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a bandwidth plan with a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// A CEN instance can be associated with multiple bandwidth plans, but cannot be associated with multiple bandwidth plans for the same connected areas.
//
// For example, if a CEN instance is already associated with a bandwidth plan for the connected areas between the Chinese mainland and the Chinese mainland, you cannot associate another bandwidth plan for the same connected areas. However, you can associate a bandwidth plan for the connected areas between the Chinese mainland and North America.
//
// @param request - AssociateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateCenBandwidthPackageResponse
func (client *Client) AssociateCenBandwidthPackageWithContext(ctx context.Context, request *AssociateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *AssociateCenBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenBandwidthPackageId) {
		query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("AssociateCenBandwidthPackage"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateCenBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route table association.
//
// Description:
//
// After you create a network instance connection, you must configure the route table association to associate the network instance connection with an Enterprise Edition transit router route table. After the association is established, the Enterprise Edition transit router forwards traffic from the network instance based on the route entry information in the route table. Before you call this operation to create a route table association, take note of the following information:
//
// - Only Enterprise Edition transit router route tables support route table associations. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - A network instance connection can be associated with only one Enterprise Edition transit router route table.
//
// - The **AssociateTransitRouterAttachmentWithRouteTable*	- operation is asynchronous. After you send the request, the system returns a **RequestId**, but the route table association is not immediately established. The system continues to process the association task in the background. You can call the **ListTransitRouterRouteTableAssociations*	- operation to query the association status between the network instance connection and the route table.
//
//   - If the association status is **Associating**, the route table association is being established between the network instance connection and the route table. In this state, you can only query the association status. You cannot perform other operations.
//
//   - If the association status is **Active**, the route table association is successfully established between the network instance connection and the route table.
//
// - Before you perform the association operation, make sure that the network instance connection is in the Attached state.
//
// @param request - AssociateTransitRouterAttachmentWithRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateTransitRouterAttachmentWithRouteTableResponse
func (client *Client) AssociateTransitRouterAttachmentWithRouteTableWithContext(ctx context.Context, request *AssociateTransitRouterAttachmentWithRouteTableRequest, runtime *dara.RuntimeOptions) (_result *AssociateTransitRouterAttachmentWithRouteTableResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateTransitRouterAttachmentWithRouteTable"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateTransitRouterAttachmentWithRouteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a vSwitch in a virtual private cloud (VPC) with a multicast domain to enable multicast communication for resources in the VPC.
//
// Description:
//
// - A vSwitch can be associated with only one multicast domain. Make sure that the vSwitch to be associated is not associated with other multicast domains. To disassociate a vSwitch from a multicast domain, see [DisassociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429774.html).
//
// - **AssociateTransitRouterMulticastDomain*	- is an asynchronous API operation. After you call this operation, the system returns a **RequestId**, but the vSwitch is not immediately associated with the multicast domain. The association task is running in the background. You can call **ListTransitRouterMulticastDomainAssociations*	- to query the association status of the vSwitch and the multicast domain.
//
//   - If the association status is **Associating**, the vSwitch is being associated with the multicast domain. In this state, you can only perform query operations on the vSwitch.
//
//   - If the association status is **Associated**, the vSwitch is associated with the multicast domain.
//
// - The VPC to which the vSwitch belongs must be connected to an Enterprise Edition transit router. To create a VPC connection, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/468237.html).
//
// @param request - AssociateTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateTransitRouterMulticastDomainResponse
func (client *Client) AssociateTransitRouterMulticastDomainWithContext(ctx context.Context, request *AssociateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *AssociateTransitRouterMulticastDomainResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateTransitRouterMulticastDomain"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateTransitRouterMulticastDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a network instance to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// CEN supports attaching cross-account network instances. Before attaching a cross-account network instance, the CEN instance must be authorized by the network instance:
//
// - To authorize a cross-account VPC instance, see [GrantInstanceToCen](https://help.aliyun.com/document_detail/126224.html).
//
// - To authorize a cross-account CCN instance, see [GrantInstanceToCbn](https://help.aliyun.com/document_detail/126141.html).
//
// - The cross-account VBR instance authorization feature is disabled by default. To use this feature, contact your account manager.
//
// This operation applies only to CEN instances that do not have an Enterprise Edition transit router. If the specified CEN instance has an Enterprise Edition transit router, the OperationFailed.InvalidTransitRouter error is returned.
//
// @param request - AttachCenChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachCenChildInstanceResponse
func (client *Client) AttachCenChildInstanceWithContext(ctx context.Context, request *AttachCenChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachCenChildInstanceResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceOwnerId) {
		query["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("AttachCenChildInstance"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachCenChildInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the transit router service is activated for the current Alibaba Cloud account.
//
// @param request - CheckTransitRouterServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckTransitRouterServiceResponse
func (client *Client) CheckTransitRouterServiceWithContext(ctx context.Context, request *CheckTransitRouterServiceRequest, runtime *dara.RuntimeOptions) (_result *CheckTransitRouterServiceResponse, _err error) {
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
		Action:      dara.String("CheckTransitRouterService"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckTransitRouterServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Cloud Enterprise Network (CEN) instance, which is a basic resource for managing an integrated network. A CEN instance manages a network that can span one or more regions. Before you connect network instances, you must call the CreateCen operation to create a CEN instance.
//
// Description:
//
// The **CreateCen*	- operation is asynchronous. The system first returns a CEN instance ID, but the instance is not yet fully created because the creation task is still running in the background. You can call the **DescribeCens*	- operation to query the status of the CEN instance.
//
// - When the CEN instance is in the **Creating*	- state, it is being created. In this state, you can only perform query operations on the CEN instance.
//
// - When the CEN instance is in the **Active*	- state, it is created.
//
// @param request - CreateCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenResponse
func (client *Client) CreateCenWithContext(ctx context.Context, request *CreateCenRequest, runtime *dara.RuntimeOptions) (_result *CreateCenResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProtectionLevel) {
		query["ProtectionLevel"] = request.ProtectionLevel
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
		Action:      dara.String("CreateCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a bandwidth plan. To enable communication between network instances across regions, you must purchase a bandwidth plan.
//
// Description:
//
// - When you create a bandwidth plan instance, you must specify the connected areas. An area is a collection of regions. Each area contains one or more Alibaba Cloud regions. You can select the connected areas based on the regions that you want to connect. For more information about the relationship between areas and regions, see [Purchase a bandwidth plan](https://help.aliyun.com/document_detail/181560.html).
//
// - For more information about the billing details of bandwidth plans, see [Billing overview](https://help.aliyun.com/document_detail/189836.html).
//
// - The **CreateCenBandwidthPackage*	- operation is asynchronous. The system returns the instance ID of the bandwidth plan before the bandwidth plan is created. The creation node is still running in the background. You can invoke the **DescribeCenBandwidthPackages*	- operation to query the status of the bandwidth plan. If the bandwidth plan is in the **Idle*	- or **InUse*	- state, the bandwidth plan is created.
//
// @param request - CreateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenBandwidthPackageResponse
func (client *Client) CreateCenBandwidthPackageWithContext(ctx context.Context, request *CreateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateCenBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewDuration) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BandwidthPackageChargeType) {
		query["BandwidthPackageChargeType"] = request.BandwidthPackageChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GeographicRegionAId) {
		query["GeographicRegionAId"] = request.GeographicRegionAId
	}

	if !dara.IsNil(request.GeographicRegionBId) {
		query["GeographicRegionBId"] = request.GeographicRegionBId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
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
		Action:      dara.String("CreateCenBandwidthPackage"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCenBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route entry for a network instance connected to an Enterprise Edition transit router. The next hop of the destination CIDR block points to the transit router in the current region.
//
// Description:
//
// - You can add route entries only for virtual private cloud (VPC) instances and virtual border router (VBR) instances connected to an Enterprise Edition transit router.
//
// - The next hop of the route entry defaults to the **transit router connection*	- (that is, the network instance connection) and cannot be modified.
//
// - The **CreateCenChildInstanceRouteEntryToAttachment*	- operation is asynchronous. After you call the operation, the system returns a **RequestId**, but the route entry is not created yet and the creation task is still running in the background. You can call the **DescribeRouteEntryList*	- operation of the VPC to query the status of the route entry.
//
//   - If the route entry is in the **Pending*	- state, the route entry is being created. In this state, you can only query the route entry and cannot perform other operations.
//
//   - If the route entry is in the **Available*	- state, the route entry is created.
//
// @param request - CreateCenChildInstanceRouteEntryToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenChildInstanceRouteEntryToAttachmentResponse
func (client *Client) CreateCenChildInstanceRouteEntryToAttachmentWithContext(ctx context.Context, request *CreateCenChildInstanceRouteEntryToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCenChildInstanceRouteEntryToAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCenChildInstanceRouteEntryToAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a route entry to a network instance.
//
// Description:
//
// - By default, the CreateCenChildInstanceRouteEntryToCen operation is not available. To use this operation, <props="china">[submit a ticket](https://selfservice.console.aliyun.com/ticket/category/cbn/today)<props="intl">[submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
// - The CreateCenChildInstanceRouteEntryToCen operation does not support adding route entries to network instances in Enterprise Edition transit routers.
//
// - The next hop of the route entry defaults to the regional gateway of the Cloud Enterprise Network (CEN) and cannot be modified.
//
// @param request - CreateCenChildInstanceRouteEntryToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenChildInstanceRouteEntryToCenResponse
func (client *Client) CreateCenChildInstanceRouteEntryToCenWithContext(ctx context.Context, request *CreateCenChildInstanceRouteEntryToCenRequest, runtime *dara.RuntimeOptions) (_result *CreateCenChildInstanceRouteEntryToCenResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceAliUid) {
		query["ChildInstanceAliUid"] = request.ChildInstanceAliUid
	}

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCenChildInstanceRouteEntryToCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cross-region traffic scheduling policy for an Enterprise Edition transit router. You can invoke the CreateCenInterRegionTrafficQosPolicy operation to optimize and control forward and routing of cross-area network traffic.
//
// Description:
//
// - Only inter-region connections of an Enterprise Edition transit router support inter-region traffic QoS policies.
//
// - The traffic QoS feature applies only to the outbound direction of an Enterprise Edition transit router.
//
//	For example, if you create an inter-region connection between the China (Hangzhou) and China (Qingdao) regions and configure the traffic QoS feature for the transit router in the China (Hangzhou) region, the feature guarantees the bandwidth for various services when traffic flows from the China (Hangzhou) region to the China (Qingdao) region. However, the feature does not guarantee the service bandwidth when traffic flows from the China (Qingdao) region to the China (Hangzhou) region.
//
// - The **CreateCenInterRegionTrafficQosPolicy*	- operation is asynchronous. After you call this operation, the system returns a traffic QoS policy ID, but the policy is not immediately created. The creation task continues in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of the traffic QoS policy.
//
//   - If the traffic QoS policy is in the **Creating*	- state, the policy is being created. In this state, you can only query the policy and cannot perform other operations.
//
//   - If the traffic QoS policy is in the **Active*	- state, the policy is created.
//
// ### Before you begin
//
// Before you call the **CreateCenInterRegionTrafficQosPolicy*	- operation, make sure that the following requirements are met:
//
// - An inter-region connection is created. For more information, see [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html).
//
// - A traffic marking policy is created. For more information, see [CreateTrafficMarkingPolicy](https://help.aliyun.com/document_detail/419025.html).
//
// @param request - CreateCenInterRegionTrafficQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenInterRegionTrafficQosPolicyResponse
func (client *Client) CreateCenInterRegionTrafficQosPolicyWithContext(ctx context.Context, request *CreateCenInterRegionTrafficQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateCenInterRegionTrafficQosPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BandwidthGuaranteeMode) {
		query["BandwidthGuaranteeMode"] = request.BandwidthGuaranteeMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConsoleDryRun) {
		query["ConsoleDryRun"] = request.ConsoleDryRun
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficQosPolicyDescription) {
		query["TrafficQosPolicyDescription"] = request.TrafficQosPolicyDescription
	}

	if !dara.IsNil(request.TrafficQosPolicyName) {
		query["TrafficQosPolicyName"] = request.TrafficQosPolicyName
	}

	if !dara.IsNil(request.TrafficQosQueues) {
		query["TrafficQosQueues"] = request.TrafficQosQueues
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCenInterRegionTrafficQosPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCenInterRegionTrafficQosPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a queue configuration for a traffic scheduling policy. You can invoke the CreateCenInterRegionTrafficQosQueue operation to manage network traffic of different types and priorities in a fine-grained manner, ensuring the performance of critical service traffic and compliance with the service level protocol.
//
// Description:
//
// The **CreateCenInterRegionTrafficQosQueue*	- operation is asynchronous. The system first returns a queue ID, but the queue is not yet created and the creation task is still running in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of the traffic scheduling policy and determine the creation status of the queue. When you query the status, you must specify the **TrafficQosPolicyId*	- parameter.
//
// - If the traffic scheduling policy is in the **Modifying*	- state, the queue is being created. In this state, you can only perform query operations on the traffic scheduling policy and the queue. You cannot perform other operations.
//
// - If the traffic scheduling policy is in the **Active*	- state, the queue is created.
//
// @param request - CreateCenInterRegionTrafficQosQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenInterRegionTrafficQosQueueResponse
func (client *Client) CreateCenInterRegionTrafficQosQueueWithContext(ctx context.Context, request *CreateCenInterRegionTrafficQosQueueRequest, runtime *dara.RuntimeOptions) (_result *CreateCenInterRegionTrafficQosQueueResponse, _err error) {
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

	if !dara.IsNil(request.Dscps) {
		query["Dscps"] = request.Dscps
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QosQueueDescription) {
		query["QosQueueDescription"] = request.QosQueueDescription
	}

	if !dara.IsNil(request.QosQueueName) {
		query["QosQueueName"] = request.QosQueueName
	}

	if !dara.IsNil(request.RemainBandwidthPercent) {
		query["RemainBandwidthPercent"] = request.RemainBandwidthPercent
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficQosPolicyId) {
		query["TrafficQosPolicyId"] = request.TrafficQosPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCenInterRegionTrafficQosQueue"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCenInterRegionTrafficQosQueueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route map for Cloud Enterprise Network (CEN) to filter route information and customize network connectivity management on Alibaba Cloud.
//
// Description:
//
// Route maps are sorted by priority. A smaller priority value indicates a higher priority. Each route map is a collection of conditional statements and execution statements. When a policy is executed, routing starts matching conditional statements from the policy with the highest priority. For routing that meets all match conditions, the policy will permit or deny the routing based on the policy action. Modifying the route property is supported for permitted routing. For routing that does not meet all match conditions, the system permits the routing to pass by default. For more information, see [Route map overview](https://help.aliyun.com/document_detail/124157.html).
//
// The `CreateCenRouteMap` operation is asynchronous. After you invoke this operation, the system returns a route map ID, but the route map is not immediately created. The creation node is still running in the background. You can invoke the `DescribeCenRouteMaps` operation to query the status of the route map.
//
// - If the route map is in the **Creating*	- state, the route map is being created. In this state, you can only execute query operations.
//
// - If the route map is in the **Active*	- state, the route map is created.
//
// @param request - CreateCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenRouteMapResponse
func (client *Client) CreateCenRouteMapWithContext(ctx context.Context, request *CreateCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *CreateCenRouteMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsPathMatchMode) {
		query["AsPathMatchMode"] = request.AsPathMatchMode
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CenRegionId) {
		query["CenRegionId"] = request.CenRegionId
	}

	if !dara.IsNil(request.CidrMatchMode) {
		query["CidrMatchMode"] = request.CidrMatchMode
	}

	if !dara.IsNil(request.CommunityMatchMode) {
		query["CommunityMatchMode"] = request.CommunityMatchMode
	}

	if !dara.IsNil(request.CommunityOperateMode) {
		query["CommunityOperateMode"] = request.CommunityOperateMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationChildInstanceTypes) {
		query["DestinationChildInstanceTypes"] = request.DestinationChildInstanceTypes
	}

	if !dara.IsNil(request.DestinationCidrBlocks) {
		query["DestinationCidrBlocks"] = request.DestinationCidrBlocks
	}

	if !dara.IsNil(request.DestinationInstanceIds) {
		query["DestinationInstanceIds"] = request.DestinationInstanceIds
	}

	if !dara.IsNil(request.DestinationInstanceIdsReverseMatch) {
		query["DestinationInstanceIdsReverseMatch"] = request.DestinationInstanceIdsReverseMatch
	}

	if !dara.IsNil(request.DestinationRegionIds) {
		query["DestinationRegionIds"] = request.DestinationRegionIds
	}

	if !dara.IsNil(request.DestinationRouteTableIds) {
		query["DestinationRouteTableIds"] = request.DestinationRouteTableIds
	}

	if !dara.IsNil(request.MapResult) {
		query["MapResult"] = request.MapResult
	}

	if !dara.IsNil(request.MatchAddressType) {
		query["MatchAddressType"] = request.MatchAddressType
	}

	if !dara.IsNil(request.MatchAsns) {
		query["MatchAsns"] = request.MatchAsns
	}

	if !dara.IsNil(request.MatchCommunitySet) {
		query["MatchCommunitySet"] = request.MatchCommunitySet
	}

	if !dara.IsNil(request.NextPriority) {
		query["NextPriority"] = request.NextPriority
	}

	if !dara.IsNil(request.OperateCommunitySet) {
		query["OperateCommunitySet"] = request.OperateCommunitySet
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Preference) {
		query["Preference"] = request.Preference
	}

	if !dara.IsNil(request.PrependAsPath) {
		query["PrependAsPath"] = request.PrependAsPath
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteTypes) {
		query["RouteTypes"] = request.RouteTypes
	}

	if !dara.IsNil(request.SourceChildInstanceTypes) {
		query["SourceChildInstanceTypes"] = request.SourceChildInstanceTypes
	}

	if !dara.IsNil(request.SourceInstanceIds) {
		query["SourceInstanceIds"] = request.SourceInstanceIds
	}

	if !dara.IsNil(request.SourceInstanceIdsReverseMatch) {
		query["SourceInstanceIdsReverseMatch"] = request.SourceInstanceIdsReverseMatch
	}

	if !dara.IsNil(request.SourceRegionIds) {
		query["SourceRegionIds"] = request.SourceRegionIds
	}

	if !dara.IsNil(request.SourceRouteTableIds) {
		query["SourceRouteTableIds"] = request.SourceRouteTableIds
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	if !dara.IsNil(request.TransmitDirection) {
		query["TransmitDirection"] = request.TransmitDirection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCenRouteMap"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCenRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a flow log.
//
// Description:
//
// Flow logs help you capture traffic information transmitted by transit router instances and network instance connections (inter-region connections, VPC connections, VPN connections, ECR connections, and VBR connections). Before you create a flow log, take note of the following items:
//
// - Only Enterprise Edition transit routers support flow logs.
//
// - For inter-region connections, flow logs capture only outbound traffic from the transit router and do not capture inbound traffic to the transit router.
//
//	For example, if an Elastic Compute Service (ECS) instance in the US (Silicon Valley) region accesses an ECS instance in the US (Virginia) region through Cloud Enterprise Network (CEN), after you configure a flow log for the transit router in the US (Virginia) region, you can view the packet information that the ECS instance in the US (Virginia) region sends to the ECS instance in the US (Silicon Valley) region in the Simple Log Service console. However, you cannot view the packet information that the ECS instance in the US (Silicon Valley) region sends to the ECS instance in the US (Virginia) region. To view the packet information that the ECS instance in the US (Silicon Valley) region sends to the ECS instance in the US (Virginia) region, you must also configure a flow log for the transit router in the US (Silicon Valley) region.
//
// - When a flow log captures traffic information for a VPC connection, it captures only the traffic transmitted by the elastic network interfaces (ENIs) of the transit router. To view the traffic transmitted by other ENIs in the VPC, see [Overview of VPC flow logs](https://help.aliyun.com/document_detail/127150.html).
//
// - The `CreateFlowlog` operation is asynchronous. The system first returns a flow log ID, but the flow log is not yet created. The creation task continues to run in the background. You can call the `DescribeFlowlogs` operation to query the status of the flow log.
//
//   - If the flow log is in the **Creating*	- state, the flow log is being created. In this state, you can only perform query operations.
//
//   - If the flow log is in the **Active*	- state, the flow log is created.
//
// ### Before you begin
//
// Before you create a flow log for a resource, make sure that you have the relevant resources. For more information about how to create each resource, see:
//
// - [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/468237.html)
//
// - [CreateTransitRouterEcrAttachment](https://help.aliyun.com/document_detail/2715446.html)
//
// - [CreateTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/468249.html)
//
// - [CreateTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/468243.html)
//
// - [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/468270.html)
//
// - [CreateTransitRouter](https://help.aliyun.com/document_detail/468222.html)
//
// @param request - CreateFlowlogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowlogResponse
func (client *Client) CreateFlowlogWithContext(ctx context.Context, request *CreateFlowlogRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowlogResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowLogName) {
		query["FlowLogName"] = request.FlowLogName
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.LogFormatString) {
		query["LogFormatString"] = request.LogFormatString
	}

	if !dara.IsNil(request.LogStoreName) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowlog"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowlogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route map to control the routing behavior between network instances.
//
// Description:
//
// The `CreateRouteMap` operation is asynchronous. The system first returns a route map ID, but the route map is not yet created. The creation task continues to run in the background. You can call the `ListRouteMaps` operation to query the status of the route map.
//
// - When the route map is in the **Creating*	- state, the route map is being created. In this state, you can only perform query operations and cannot perform other operations.
//
// - When the route map is in the **Active*	- state, the route map is created.
//
// @param request - CreateRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteMapResponse
func (client *Client) CreateRouteMapWithContext(ctx context.Context, request *CreateRouteMapRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteMapResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRouteMap"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route map rule.
//
// Description:
//
// ## Request description
//
// - `RouteMapId` is a required parameter that specifies the route map to which the new rule belongs.
//
// - `Priority` is a required parameter. A smaller value indicates a higher priority. Rules within the same route map cannot have duplicate priorities.
//
// - `Name` and `Description` are optional parameters used to identify the rule and describe its purpose.
//
// - `MatchConditions` defines the match conditions for the rule, and `SetActions` defines the actions to be executed when a route meets all match conditions. Both are passed as JSON strings.
//
// - `DryRun` is an optional parameter used to precheck the request without actually creating the route map rule.
//
// @param tmpReq - CreateRouteMapRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteMapRuleResponse
func (client *Client) CreateRouteMapRuleWithContext(ctx context.Context, tmpReq *CreateRouteMapRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteMapRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRouteMapRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MatchConditions) {
		request.MatchConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchConditions, dara.String("MatchConditions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SetActions) {
		request.SetActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SetActions, dara.String("SetActions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.MatchConditionsShrink) {
		query["MatchConditions"] = request.MatchConditionsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	if !dara.IsNil(request.SetActionsShrink) {
		query["SetActions"] = request.SetActionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRouteMapRule"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteMapRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A traffic marking policy captures traffic that matches traffic classification rules and adds a Differentiated Services Code Point (DSCP) value to the traffic as a mark. Creates a traffic marking policy by calling CreateTrafficMarkingPolicy.
//
// Description:
//
// - Only Enterprise Edition transit routers support creating traffic marking policies.
//
// - **CreateTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a traffic marking policy ID but the traffic marking policy is not yet created. The system continues to create the traffic marking policy in the background. You can call **ListTrafficMarkingPolicies*	- to query the status of the traffic marking policy.
//
//   - If the traffic marking policy is in the **Creating*	- state, the traffic marking policy is being created. In this state, you can only query the traffic marking policy but cannot perform other operations.
//
//   - If the traffic marking policy is in the **Active*	- state, the traffic marking policy is created.
//
// @param request - CreateTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficMarkingPolicyResponse
func (client *Client) CreateTrafficMarkingPolicyWithContext(ctx context.Context, request *CreateTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateTrafficMarkingPolicyResponse, _err error) {
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

	if !dara.IsNil(request.MarkingDscp) {
		query["MarkingDscp"] = request.MarkingDscp
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyDescription) {
		query["TrafficMarkingPolicyDescription"] = request.TrafficMarkingPolicyDescription
	}

	if !dara.IsNil(request.TrafficMarkingPolicyName) {
		query["TrafficMarkingPolicyName"] = request.TrafficMarkingPolicyName
	}

	if !dara.IsNil(request.TrafficMatchRules) {
		query["TrafficMatchRules"] = request.TrafficMatchRules
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrafficMarkingPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an aggregate route.
//
// Description:
//
// After you add an aggregate route to an Enterprise Edition transit router route table, the Enterprise Edition transit router propagates the aggregate route only to the route tables of VPC-connected instances that are associated with the current Enterprise Edition transit router route table and have route synchronization enabled.
//
// Before creating an aggregate route, make sure that the following operations are completed. Otherwise, the Enterprise Edition transit router does not propagate the aggregate route to VPC instance route tables:
//
// - The VPC instance is associated with the Enterprise Edition transit router route table. For more information, see [AssociateTransitRouterAttachmentWithRouteTable](https://help.aliyun.com/document_detail/261242.html).
//
// - Route synchronization is enabled for the VPC instance. For more information, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261358.html).
//
// @param tmpReq - CreateTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouteTableAggregationResponse
func (client *Client) CreateTransitRouteTableAggregationWithContext(ctx context.Context, tmpReq *CreateTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouteTableAggregationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTransitRouteTableAggregationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TransitRouteTableAggregationScopeList) {
		request.TransitRouteTableAggregationScopeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransitRouteTableAggregationScopeList, dara.String("TransitRouteTableAggregationScopeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouteTableAggregationCidr) {
		query["TransitRouteTableAggregationCidr"] = request.TransitRouteTableAggregationCidr
	}

	if !dara.IsNil(request.TransitRouteTableAggregationDescription) {
		query["TransitRouteTableAggregationDescription"] = request.TransitRouteTableAggregationDescription
	}

	if !dara.IsNil(request.TransitRouteTableAggregationName) {
		query["TransitRouteTableAggregationName"] = request.TransitRouteTableAggregationName
	}

	if !dara.IsNil(request.TransitRouteTableAggregationScope) {
		query["TransitRouteTableAggregationScope"] = request.TransitRouteTableAggregationScope
	}

	if !dara.IsNil(request.TransitRouteTableAggregationScopeListShrink) {
		query["TransitRouteTableAggregationScopeList"] = request.TransitRouteTableAggregationScopeListShrink
	}

	if !dara.IsNil(request.TransitRouteTableId) {
		query["TransitRouteTableId"] = request.TransitRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouteTableAggregation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouteTableAggregationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Enterprise Edition transit router instance.
//
// Description:
//
// - You can invoke the **CreateTransitRouter*	- operation to create an Enterprise Edition transit router instance to forward traffic and perform routing. Enterprise Edition transit routers are available only in specific regions. For more information about the supported regions, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - The **CreateTransitRouter*	- operation is asynchronous. After you invoke this operation, the system returns the Enterprise Edition transit router instance ID. However, the instance is not immediately created, and the creation node is still running in the background. You can invoke the [ListTransitRouters](https://help.aliyun.com/document_detail/261219.html) operation to query the status of the Enterprise Edition transit router instance.
//
//   - If the Enterprise Edition transit router instance is in the **Creating*	- state, the instance is being created. In this state, you can only execute query operations on the instance and cannot perform other operations.
//
//   - If the Enterprise Edition transit router instance is in the **Active*	- state, the instance is created.
//
// - You can create only one transit router instance in each region for a Cloud Enterprise Network (CEN) instance.
//
// @param tmpReq - CreateTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterResponse
func (client *Client) CreateTransitRouterWithContext(ctx context.Context, tmpReq *CreateTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTransitRouterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TransitRouterCidrList) {
		request.TransitRouterCidrListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransitRouterCidrList, dara.String("TransitRouterCidrList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
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

	if !dara.IsNil(request.SupportMulticast) {
		query["SupportMulticast"] = request.SupportMulticast
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterCidrListShrink) {
		query["TransitRouterCidrList"] = request.TransitRouterCidrListShrink
	}

	if !dara.IsNil(request.TransitRouterDescription) {
		query["TransitRouterDescription"] = request.TransitRouterDescription
	}

	if !dara.IsNil(request.TransitRouterName) {
		query["TransitRouterName"] = request.TransitRouterName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouter"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A transit router CIDR block is a custom CIDR block that you define for a transit router, similar to a CIDR block used to assign IP addresses to a router loopback interface. Calls the CreateTransitRouterCidr operation to create a CIDR block for a transit router.
//
// Description:
//
// A transit router CIDR block is a custom CIDR block that you define for a transit router, similar to a CIDR block used to assign IP addresses to a router loopback interface. Transit router CIDR blocks are used to assign addresses to network instance connections. For more information, see [Transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// The **CreateTransitRouterCidr*	- operation is used only to add a CIDR block to a transit router after the transit router is created.
//
// Before you create a transit router CIDR block, take note of the following information:
//
// - Only Enterprise Edition transit routers support transit router CIDR blocks.
//
// - For limits on transit router CIDR blocks, see [Limits on transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// - A maximum of five CIDR blocks can be configured for a transit router. The subnet mask of each CIDR block must be 16 to 24 bits in length.
//
// - CIDR blocks that fall within 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16 and their subnets are not supported.
//
// - Each CIDR block cannot conflict with any CIDR block that needs to communicate within the Cloud Enterprise Network (CEN) instance.
//
// - Each CIDR block must be unique within the same CEN instance.
//
// - After you add a CIDR block to a transit router, the system automatically reserves three CIDR blocks from the CIDR block when you create the first VPN connection on the transit router. The reserved CIDR blocks are used by the system to create VPN connections in the background. The system assigns IP addresses to IPsec connections from the remaining CIDR blocks.
//
//	You can call the [ListTransitRouterCidrAllocation](https://help.aliyun.com/document_detail/464173.html) operation to query the CIDR blocks that are reserved by the system or from which IP addresses are allocated.
//
// @param request - CreateTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterCidrResponse
func (client *Client) CreateTransitRouterCidrWithContext(ctx context.Context, request *CreateTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterCidrResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PublishCidrRoute) {
		query["PublishCidrRoute"] = request.PublishCidrRoute
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

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterCidr"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterCidrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Express Connect Router (ECR) connection to connect an ECR network instance to a transit router in the same region.
//
// Description:
//
// - Only Enterprise Edition transit routers support ECR connections.
//
// - You can create an ECR connection on an Enterprise Edition transit router in one of the following ways:
//
//   - If you have already created an Enterprise Edition transit router in the target region, you can create an ECR connection by specifying **EcrId**, **RegionId**, and **TransitRouterId**.
//
//   - If you do not have an Enterprise Edition transit router in the target region, you can create an ECR connection by specifying **EcrId**, **CenId**, and **RegionId**. During the creation of the ECR connection, the system performs automatic creation of an Enterprise Edition transit router.
//
// - The CreateTransitRouterEcrAttachment operation is asynchronous. The system first returns an ECR connection ID, but the ECR connection is not fully created and the creation node is still running in the background. You can invoke the ListTransitRouterEcrAttachments operation to query the status of the ECR connection.
//
//   - If the ECR connection is in the **Attaching*	- state, the ECR connection is being created. In this state, you can only query the ECR connection and cannot execute other operations.
//
//   - If the ECR connection is in the **Attached*	- state, the ECR connection is created.
//
// - After an ECR connection is created, the ECR connection does not establish routing learning or associated forwarding relationships with any Enterprise Edition transit router routing tables by default.
//
//	After the ECR connection establishes a [routing learning relationship](https://help.aliyun.com/document_detail/468300.html) with an Enterprise Edition transit router routing table, the system automatically propagates routing from the ECR instance to the Enterprise Edition transit router routing table.
//
// - After an ECR connection is created, the system automatically propagates routing from the associated Enterprise Edition transit router routing table to the routing table of the ECR instance.
//
// ### Before you begin
//
// - The Alibaba Cloud account to which the Enterprise Edition transit router belongs and the Alibaba Cloud account to which the ECR instance belongs must belong to the same enterprise.
//
// - Enterprise Edition transit routers support connections to ECR instances within the same account or across accounts. Before you create a cross-account ECR connection, obtain authorization for the cross-account ECR instance. For more information, see [Grant permissions on cross-account network instances](https://help.aliyun.com/document_detail/181553.html).
//
// - **Before you invoke this operation to create an ECR connection, invoke the [CreateExpressConnectRouterAssociation](https://help.aliyun.com/document_detail/2712082.html) operation to associate the ECR instance with the Enterprise Edition transit router instance.**
//
//	**When you invoke the DeleteTransitRouterEcrAttachment operation to force delete an ECR connection, the system also deletes the association between the ECR instance and the Enterprise Edition transit router instance. You do not need to delete the association separately.**
//
// @param request - CreateTransitRouterEcrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterEcrAttachmentResponse
func (client *Client) CreateTransitRouterEcrAttachmentWithContext(ctx context.Context, request *CreateTransitRouterEcrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterEcrAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		query["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.EcrOwnerId) {
		query["EcrOwnerId"] = request.EcrOwnerId
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterEcrAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterEcrAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a multicast domain. A multicast domain defines the scope of a multicast network within a region. Only resources within the multicast domain can send and receive multicast traffic. Resources outside the multicast domain cannot send or receive multicast traffic.
//
// Description:
//
// Before you call this operation, take note of the following information:
//
// - Make sure that you have created an Enterprise Edition transit router in the region where you want to establish the multicast network, and enabled the multicast feature for the Enterprise Edition transit router. For more information, see [CreateTransitRouter](https://help.aliyun.com/document_detail/261169.html).
//
//	If you have already created an Enterprise Edition transit router instance before applying for multicast resources, the multicast feature cannot be enabled for the instance. You can delete the current Enterprise Edition transit router instance and create a new one. For more information about how to delete an Enterprise Edition transit router instance, see [DeleteTransitRouter](https://help.aliyun.com/document_detail/261218.html).
//
// - When you call the **CreateTransitRouterMulticastDomain*	- operation, if you specify the **CenId*	- and **RegionId*	- parameters, you do not need to specify the **TransitRouterId*	- parameter. If you specify the **TransitRouterId*	- parameter, you do not need to specify the **CenId*	- and **RegionId*	- parameters.
//
// @param request - CreateTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterMulticastDomainResponse
func (client *Client) CreateTransitRouterMulticastDomainWithContext(ctx context.Context, request *CreateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterMulticastDomainResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Options) {
		query["Options"] = request.Options
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainDescription) {
		query["TransitRouterMulticastDomainDescription"] = request.TransitRouterMulticastDomainDescription
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainName) {
		query["TransitRouterMulticastDomainName"] = request.TransitRouterMulticastDomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterMulticastDomain"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterMulticastDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After network instances (VPCs, VBRs, and IPsec connections) are connected to a transit router, you must create an inter-region connection between transit routers to enable communication between network instances in different regions. You can call the CreateTransitRouterPeerAttachment operation to create an inter-region connection for an Enterprise Edition transit router instance.
//
// Description:
//
// - Enterprise Edition transit routers support the following two bandwidth allocation methods:
//
//   - **Allocate from bandwidth package**:
//
//     To use this method, you must first purchase a bandwidth package and allocate bandwidth from the bandwidth package to the inter-region connection. For more information about how to purchase a bandwidth package, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html).
//
//   - **Pay-by-data-transfer**:
//
//     To use this method, you must set a bandwidth limit for the inter-region connection. The system charges you based on the actual traffic of the inter-region connection. For more information about billing, see [Inter-region data transfer](https://help.aliyun.com/document_detail/337827.html).
//
// - **CreateTransitRouterPeerAttachment*	- is an asynchronous operation. After you send a request, the system returns an inter-region connection ID, but the inter-region connection is not yet created. The creation task still runs in the background. You can call the **ListTransitRouterPeerAttachments*	- operation to query the status of the inter-region connection.
//
//   - If the inter-region connection is in the **Attaching*	- state, the inter-region connection is being created. In this state, you can only query the inter-region connection but cannot perform other operations on it.
//
//   - If the inter-region connection is in the **Attached*	- state, the inter-region connection is created.
//
// @param request - CreateTransitRouterPeerAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterPeerAttachmentResponse
func (client *Client) CreateTransitRouterPeerAttachmentWithContext(ctx context.Context, request *CreateTransitRouterPeerAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterPeerAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPublishRouteEnabled) {
		query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BandwidthType) {
		query["BandwidthType"] = request.BandwidthType
	}

	if !dara.IsNil(request.CenBandwidthPackageId) {
		query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DefaultLinkType) {
		query["DefaultLinkType"] = request.DefaultLinkType
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PeerTransitRouterId) {
		query["PeerTransitRouterId"] = request.PeerTransitRouterId
	}

	if !dara.IsNil(request.PeerTransitRouterRegionId) {
		query["PeerTransitRouterRegionId"] = request.PeerTransitRouterRegionId
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterPeerAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterPeerAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a prefix list with an Enterprise Edition transit router route table.
//
// Description:
//
// Before you associate a prefix list with an Enterprise Edition transit router route table, make sure that the following conditions are met:
//
// - You have learned about the limits and routing compatibility information of prefix lists. For more information, see [Prefix lists](https://help.aliyun.com/document_detail/445605.html).
//
// - You have created a prefix list. For more information, see [CreateVpcPrefixList](https://help.aliyun.com/document_detail/437367.html).
//
// - If you want to associate a cross-account prefix list with an Enterprise Edition transit router route table, make sure that the prefix list has been shared with the Alibaba Cloud account that owns the Enterprise Edition transit router route table. For more information about how to share a prefix list, see [Overview of resource sharing](https://help.aliyun.com/document_detail/160622.html) and [API reference (Resource Sharing)](https://help.aliyun.com/document_detail/193445.html).
//
// @param request - CreateTransitRouterPrefixListAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterPrefixListAssociationResponse
func (client *Client) CreateTransitRouterPrefixListAssociationWithContext(ctx context.Context, request *CreateTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterPrefixListAssociationResponse, _err error) {
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

	if !dara.IsNil(request.NextHop) {
		query["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.NextHopType) {
		query["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.OwnerUid) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !dara.IsNil(request.PrefixListId) {
		query["PrefixListId"] = request.PrefixListId
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

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterTableId) {
		query["TransitRouterTableId"] = request.TransitRouterTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterPrefixListAssociation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterPrefixListAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a route entry to a route table of an Enterprise Edition transit router.
//
// Description:
//
// *CreateTransitRouterRouteEntry*	- is an asynchronous operation. After you send a request, the system returns a route entry ID but the route entry is not yet created. The creation task continues to run in the background. You can call **ListTransitRouterRouteEntries*	- to query the status of the route entry.
//
// - If the route entry is in the **Creating*	- state, the route entry is being created. In this state, you can only query the route entry. You cannot perform other operations on the route entry.
//
// - If the route entry is in the **Active*	- state, the route entry is created.
//
// @param request - CreateTransitRouterRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterRouteEntryResponse
func (client *Client) CreateTransitRouterRouteEntryWithContext(ctx context.Context, request *CreateTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterRouteEntryResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryDescription) {
		query["TransitRouterRouteEntryDescription"] = request.TransitRouterRouteEntryDescription
	}

	if !dara.IsNil(request.TransitRouterRouteEntryDestinationCidrBlock) {
		query["TransitRouterRouteEntryDestinationCidrBlock"] = request.TransitRouterRouteEntryDestinationCidrBlock
	}

	if !dara.IsNil(request.TransitRouterRouteEntryName) {
		query["TransitRouterRouteEntryName"] = request.TransitRouterRouteEntryName
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNextHopId) {
		query["TransitRouterRouteEntryNextHopId"] = request.TransitRouterRouteEntryNextHopId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNextHopType) {
		query["TransitRouterRouteEntryNextHopType"] = request.TransitRouterRouteEntryNextHopType
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterRouteEntry"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom route table for an Enterprise Edition transit router by calling the CreateTransitRouterRouteTable operation.
//
// Description:
//
// - Only Enterprise Edition transit routers support custom route tables. For information about the regions and zones that support Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html).
//
// - **CreateTransitRouterRouteTable*	- is an asynchronous operation. After you send a request, the system returns a route table ID but the route table is not yet created. The system continues to create the route table in the background. You can call **ListTransitRouterRouteTables*	- to query the status of the route table.
//
//   - If the route table is in the **Creating*	- state, the route table is being created. In this state, you can only perform query operations.
//
//   - If the route table is in the **Active*	- state, the route table is created.
//
// @param request - CreateTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterRouteTableResponse
func (client *Client) CreateTransitRouterRouteTableWithContext(ctx context.Context, request *CreateTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterRouteTableResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteTableOptions) {
		query["RouteTableOptions"] = request.RouteTableOptions
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterRouteTableDescription) {
		query["TransitRouterRouteTableDescription"] = request.TransitRouterRouteTableDescription
	}

	if !dara.IsNil(request.TransitRouterRouteTableName) {
		query["TransitRouterRouteTableName"] = request.TransitRouterRouteTableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterRouteTable"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterRouteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a VBR connection on an Enterprise Edition transit router. You can invoke the CreateTransitRouterVbrAttachment operation to connect a virtual border router (VBR) network instance to a transit router instance that can forward and perform routing in the same region. After the connection is established, the transit router enables private network peering.
//
// Description:
//
// - For information about the regions and zones supported by Enterprise Edition transit routers, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - You can create a VBR connection on an Enterprise Edition transit router by using one of the following methods:
//
//   - If you have already created an Enterprise Edition transit router instance that can forward and perform routing in the target region, you can create a VBR connection by specifying **VbrId**, **RegionId**, and **TransitRouterId**.
//
//   - If you have not created an Enterprise Edition transit router instance in the target region, you can create a VBR connection by specifying **VbrId**, **CenId**, and **RegionId**. The system performs automatic creation of an Enterprise Edition transit router instance that can forward and perform routing when the VBR connection is created.
//
// - **CreateTransitRouterVbrAttachment*	- is an asynchronous operation. The system returns a VBR connection ID, but the VBR connection is not yet created and the creation task is still running in the background. You can call **ListTransitRouterVbrAttachments*	- to query the status of the VBR connection.
//
//   - If the VBR connection is in the **Attaching*	- state, the VBR connection is being created. In this state, you can only query the VBR connection and cannot perform other operations.
//
//   - If the VBR connection is in the **Attached*	- state, the VBR connection is created.
//
// - The Alibaba Cloud account to which the transit router belongs and the Alibaba Cloud account to which the VBR belongs must belong to the same enterprise.
//
// - Transit routers support connecting VBRs within the same account and across accounts. Before creating a cross-account VBR connection, obtain authorization for the cross-account VBR. For more information, see [Grant permissions on cross-account network instances](https://help.aliyun.com/document_detail/181553.html).
//
// - After a VBR connection is created, the VBR connection does not establish route learning or associated forwarding relationships with any transit router route tables by default.
//
// @param request - CreateTransitRouterVbrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVbrAttachmentResponse
func (client *Client) CreateTransitRouterVbrAttachmentWithContext(ctx context.Context, request *CreateTransitRouterVbrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVbrAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPublishRouteEnabled) {
		query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.VbrId) {
		query["VbrId"] = request.VbrId
	}

	if !dara.IsNil(request.VbrOwnerId) {
		query["VbrOwnerId"] = request.VbrOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterVbrAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterVbrAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateTransitRouterVpcAttachment operation to attach a Virtual Private Cloud (VPC) instance to an Enterprise Edition transit router instance to enable private network peering. After the attachment is created, the transit router instance helps you implement private network connectivity through forwarding and routing.
//
// Description:
//
// - You can create a VPC attachment for an Enterprise Edition transit router in two ways:
//
//   - If you have already created an Enterprise Edition transit router instance in the target region, specify **VpcId**, **ZoneMappings.N.VSwitchId**, **ZoneMappings.N.ZoneId**, **TransitRouterId**, and **RegionId*	- to create the VPC attachment.
//
//   - If you have not created an Enterprise Edition transit router instance in the target region, specify **VpcId**, **ZoneMappings.N.VSwitchId**, **ZoneMappings.N.ZoneId**, **CenId**, and **RegionId*	- to create the VPC attachment. The system uses automatic creation to create an Enterprise Edition transit router instance when you create the VPC attachment.
//
// - **CreateTransitRouterVpcAttachment*	- is an asynchronous operation. The system returns a VPC attachment ID before the VPC attachment is created. The creation task runs in the background. Call [ListTransitRouterVpcAttachments](https://help.aliyun.com/document_detail/261222.html) to query the status of the VPC attachment.
//
//   - When the VPC attachment is in the **Attaching*	- state, the VPC attachment is being created. In this state, you can only query the VPC attachment and cannot execute other operations.
//
//   - When the VPC attachment is in the **Attached*	- state, the VPC attachment is created.
//
// - After a VPC attachment is created, it does not associate with any transit router route table for route learning relationships or forwarding relationships by default.
//
// ### Before you begin
//
// Before you call this operation to create a VPC attachment, make sure that the following requirements are met:
//
// - The VPC instance has at least one vSwitch instance with at least one available IP address in a zone supported by the Enterprise Edition transit router. For information about the regions and zones supported by Enterprise Edition transit routers, see [Regions and zones supported by Enterprise Edition transit routers](https://help.aliyun.com/document_detail/181681.html).
//
// - If you want to attach a VPC-connected instance that belongs to a different account, obtain cross-account authorization first. For more information, see [Cross-account VPC-connected instance authorization](https://help.aliyun.com/document_detail/181553.html).
//
// - Creating a VPC attachment incurs fees. Make sure that you understand the billing rules. For more information, see [Billing](https://help.aliyun.com/document_detail/189836.html).
//
// @param tmpReq - CreateTransitRouterVpcAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVpcAttachmentResponse
func (client *Client) CreateTransitRouterVpcAttachmentWithContext(ctx context.Context, tmpReq *CreateTransitRouterVpcAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVpcAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTransitRouterVpcAttachmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("Options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransitRouterVPCAttachmentOptions) {
		request.TransitRouterVPCAttachmentOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransitRouterVPCAttachmentOptions, dara.String("TransitRouterVPCAttachmentOptions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPublishRouteEnabled) {
		query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OptionsShrink) {
		query["Options"] = request.OptionsShrink
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterVPCAttachmentOptionsShrink) {
		query["TransitRouterVPCAttachmentOptions"] = request.TransitRouterVPCAttachmentOptionsShrink
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcOwnerId) {
		query["VpcOwnerId"] = request.VpcOwnerId
	}

	if !dara.IsNil(request.ZoneMappings) {
		query["ZoneMappings"] = request.ZoneMappings
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterVpcAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterVpcAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transit routers support connections to IPsec connections of VPN gateways. An on-premises data center can directly connect to a transit router through an IPsec connection, and then communicate with other networks through the transit router. Calls the CreateTransitRouterVpnAttachment operation to create a VPN connection.
//
// Description:
//
// - After a VPN connection is created, the VPN connection does not establish route learning or association forwarding relationships with any transit router route table by default.
//
// - When you call the `CreateTransitRouterVpnAttachment` operation, if you specify values for the **CenId*	- and **RegionId*	- parameters, you do not need to specify the **TransitRouterId*	- parameter. If you specify values for the **TransitRouterId*	- and **RegionId*	- parameters, you do not need to specify the **CenId*	- parameter.
//
// ### Before you begin
//
// - Before you create a VPN connection, make sure that you have created an IPsec connection in the region where the transit router instance resides and that the IPsec connection is not bindeded to any resource. For more information, see [CreateVpnAttachment](https://help.aliyun.com/document_detail/442455.html).
//
// - If the transit router instance needs to connect to an IPsec connection that belongs to a different Alibaba Cloud account, make sure that the IPsec connection has been authorized to the transit router instance. For more information, see [GrantInstanceToTransitRouter](https://help.aliyun.com/document_detail/417520.html).
//
// - Before you create a VPN connection, make sure that you have configured the TR CIDR block for the transit router. For more information, see [CreateTransitRouterCidr](https://help.aliyun.com/document_detail/468230.html).
//
// @param request - CreateTransitRouterVpnAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVpnAttachmentResponse
func (client *Client) CreateTransitRouterVpnAttachmentWithContext(ctx context.Context, request *CreateTransitRouterVpnAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVpnAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPublishRouteEnabled) {
		query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.VpnId) {
		query["VpnId"] = request.VpnId
	}

	if !dara.IsNil(request.VpnOwnerId) {
		query["VpnOwnerId"] = request.VpnOwnerId
	}

	if !dara.IsNil(request.Zone) {
		query["Zone"] = request.Zone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitRouterVpnAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitRouterVpnAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deactivates a flow log. After the flow log is deactivated, traffic of the specified resource is no longer captured.
//
// Description:
//
// The `DeactiveFlowLog` operation is asynchronous. After you send a request, the system returns a **RequestId*	- but the flow log is not completely deactivated. The deactivation task continues to run in the background. You can call the `DescribeFlowlogs` operation to query the status of the flow log.
//
// - If the flow log is in the **Modifying*	- state, the flow log is being deactivated. In this state, you can only perform query operations.
//
// - If the flow log is in the **Inactive*	- state, the flow log is deactivated.
//
// @param request - DeactiveFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactiveFlowLogResponse
func (client *Client) DeactiveFlowLogWithContext(ctx context.Context, request *DeactiveFlowLogRequest, runtime *dara.RuntimeOptions) (_result *DeactiveFlowLogResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
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
		Action:      dara.String("DeactiveFlowLog"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactiveFlowLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// *DeleteCen*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the CEN instance is not yet deleted. The deletion task continues to run in the background. You can call the **DescribeCens*	- operation to query the status of the CEN instance.
//
// - If the CEN instance is in the **Deleting*	- state, the CEN instance is being deleted. In this state, you can only query the CEN instance. You cannot perform other operations on it.
//
// - If the CEN instance cannot be found, the CEN instance is deleted.
//
// ### Before you begin
//
// Before you delete a CEN instance, make sure that no bandwidth plans exist under the CEN instance, and that no network instance connections or custom route tables exist under the transit routers of the CEN instance:
//
// - To delete network instance connections, see:
//
//   - [DeleteTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261220.html)
//
//   - [DeleteTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/261223.html)
//
//   - [DeleteTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/443992.html)
//
//   - [DeleteTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261227.html)
//
//     > To delete network instance connections under a Basic Edition transit router, see [DetachCenChildInstance](https://help.aliyun.com/document_detail/65915.html).
//
// - To delete custom route tables of an Enterprise Edition transit router, see [DeleteTransitRouterRouteTable](https://help.aliyun.com/document_detail/261235.html).
//
// - To disassociate a bandwidth plan from a CEN instance, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// @param request - DeleteCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenResponse
func (client *Client) DeleteCenWithContext(ctx context.Context, request *DeleteCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DeleteCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a bandwidth plan instance by calling the DeleteCenBandwidthPackage operation.
//
// Description:
//
// <props="china">
//
// - Before you delete a bandwidth plan instance, make sure that the bandwidth plan instance is disassociated from the Cloud Enterprise Network (CEN) instance. For more information, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// - To delete a bandwidth plan instance whose billing method is PREPAY (subscription), go to the [Order Center](https://usercenter2.aliyun.com/refund/refund) to unsubscribe. If you have questions about unsubscription, see [Unsubscription rules](https://www.alibabacloud.com/help/en/user-center/user-guide/unsubscription-rules#p-1qo-3ce-m7z). This operation does not support deleting subscription bandwidth plan instances.
//
// <props="intl">
//
// Before you delete a bandwidth plan instance, make sure that the bandwidth plan instance is disassociated from the Cloud Enterprise Network (CEN) instance. For more information, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// @param request - DeleteCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenBandwidthPackageResponse
func (client *Client) DeleteCenBandwidthPackageWithContext(ctx context.Context, request *DeleteCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenBandwidthPackageId) {
		query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DeleteCenBandwidthPackage"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCenBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a route entry from a network instance connected to an Enterprise Edition transit router.
//
// Description:
//
// - Only route entries whose next hop is a **transit router connection*	- (network instance connection) in Virtual Private Cloud (VPC) instances and Virtual Border Router (VBR) instances can be deleted.
//
// - **DeleteCenChildInstanceRouteEntryToAttachment*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the route entry is not yet deleted. The deletion task runs in the background. You can call the **DescribeRouteEntryList*	- operation of VPC to query the status of the route entry.
//
//   - If the route entry is in the **Deleting*	- state, the route entry is being deleted. In this state, you can only query the route entry but cannot perform other operations on it.
//
//   - If the specified route entry cannot be found, the route entry is deleted.
//
// @param request - DeleteCenChildInstanceRouteEntryToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenChildInstanceRouteEntryToAttachmentResponse
func (client *Client) DeleteCenChildInstanceRouteEntryToAttachmentWithContext(ctx context.Context, request *DeleteCenChildInstanceRouteEntryToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCenChildInstanceRouteEntryToAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCenChildInstanceRouteEntryToAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a route entry from a network instance by calling the DeleteCenChildInstanceRouteEntryToCen operation.
//
// Description:
//
// - The DeleteCenChildInstanceRouteEntryToCen operation is not available by default. To use this operation, <props="china">[submit a ticket](https://selfservice.console.aliyun.com/ticket/category/cbn/today)<props="intl">[submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
// - The DeleteCenChildInstanceRouteEntryToCen operation does not support deleting route entries from network instances attached to an Enterprise Edition transit router.
//
// @param request - DeleteCenChildInstanceRouteEntryToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenChildInstanceRouteEntryToCenResponse
func (client *Client) DeleteCenChildInstanceRouteEntryToCenWithContext(ctx context.Context, request *DeleteCenChildInstanceRouteEntryToCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenChildInstanceRouteEntryToCenResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceAliUid) {
		query["ChildInstanceAliUid"] = request.ChildInstanceAliUid
	}

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteTableId) {
		query["RouteTableId"] = request.RouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCenChildInstanceRouteEntryToCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a traffic scheduling policy by calling the DeleteCenInterRegionTrafficQosPolicy operation.
//
// Description:
//
// - Before you delete a traffic scheduling policy, you must delete all queues except the default queue from the traffic scheduling policy. For more information, see [DeleteCenInterRegionTrafficQosQueue](https://help.aliyun.com/document_detail/419062.html).
//
// - **DeleteCenInterRegionTrafficQosPolicy*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the traffic scheduling policy is not yet deleted. The deletion task runs in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of the traffic scheduling policy.
//
//   - If the traffic scheduling policy is in the **Deleting*	- state, the traffic scheduling policy is being deleted. In this state, you can only query the traffic scheduling policy but cannot perform other operations on it.
//
//   - If the traffic scheduling policy cannot be found, the traffic scheduling policy is deleted.
//
// @param request - DeleteCenInterRegionTrafficQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenInterRegionTrafficQosPolicyResponse
func (client *Client) DeleteCenInterRegionTrafficQosPolicyWithContext(ctx context.Context, request *DeleteCenInterRegionTrafficQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenInterRegionTrafficQosPolicyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficQosPolicyId) {
		query["TrafficQosPolicyId"] = request.TrafficQosPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCenInterRegionTrafficQosPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCenInterRegionTrafficQosPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a queue from a traffic scheduling policy by calling the DeleteCenInterRegionTrafficQosQueue operation.
//
// Description:
//
// - The default queue cannot be deleted.
//
// - **DeleteCenInterRegionTrafficQosQueue*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the queue is not yet deleted because the deletion task is still running in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the queue information. If the specified queue cannot be found, the queue is deleted.
//
// @param request - DeleteCenInterRegionTrafficQosQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenInterRegionTrafficQosQueueResponse
func (client *Client) DeleteCenInterRegionTrafficQosQueueWithContext(ctx context.Context, request *DeleteCenInterRegionTrafficQosQueueRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenInterRegionTrafficQosQueueResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QosQueueId) {
		query["QosQueueId"] = request.QosQueueId
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
		Action:      dara.String("DeleteCenInterRegionTrafficQosQueue"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCenInterRegionTrafficQosQueueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified routing policy.
//
// Description:
//
// DeleteCenRouteMap is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the routing policy is not yet fully deleted because the deletion task is still running in the background. You can call the `DescribeCenRouteMaps` operation to query the status of the routing policy.
//
// - If the routing policy is in the **Deleting*	- state, the routing policy is being deleted. In this state, you can only perform query operations.
//
// - If the routing policy cannot be found by calling the `DescribeCenRouteMaps` operation, the routing policy is fully deleted.
//
// @param request - DeleteCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenRouteMapResponse
func (client *Client) DeleteCenRouteMapWithContext(ctx context.Context, request *DeleteCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenRouteMapResponse, _err error) {
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

	if !dara.IsNil(request.CenRegionId) {
		query["CenRegionId"] = request.CenRegionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCenRouteMap"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCenRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a flow log.
//
// Description:
//
// The `DeleteFlowlog` operation is asynchronous. After you send a request, the system returns a **RequestId*	- but the flow log is not completely deleted. The deletion task continues to run in the background. You can call the `DescribeFlowlogs` operation to query the status of the flow log.
//
// - If the flow log is in the **Deleting*	- state, the flow log is being deleted. In this state, you can only perform query operations but cannot perform other operations.
//
// - If the `DescribeFlowlogs` operation cannot find the flow log, the flow log is completely deleted.
//
// @param request - DeleteFlowlogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlogWithContext(ctx context.Context, request *DeleteFlowlogRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowlogResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
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
		Action:      dara.String("DeleteFlowlog"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowlogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified route map.
//
// Description:
//
// If the route map to be deleted is still associated with any resources, this operation fails.
//
// If route map rules exist under the route map, this operation is not supported.
//
// The DeleteRouteMap operation is asynchronous. The system returns a **RequestId**, but the route map is not completely deleted. The deletion task continues in the background. You can call the ListRouteMaps operation to query the status of the route map.
//
// - If the route map is in the **Deleting*	- state, the route map is being deleted. In this state, you can only perform query operations.
//
// - If the route map cannot be queried by using the ListRouteMaps operation, the route map is completely deleted.
//
// @param request - DeleteRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteMapResponse
func (client *Client) DeleteRouteMapWithContext(ctx context.Context, request *DeleteRouteMapRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteMapResponse, _err error) {
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

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRouteMap"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified route map rule.
//
// Description:
//
// ## Request description
//
// - This API operation deletes an existing route map rule.
//
// - If the rule to be deleted is still associated with any resources, this operation fails.
//
// @param request - DeleteRouteMapRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteMapRuleResponse
func (client *Client) DeleteRouteMapRuleWithContext(ctx context.Context, request *DeleteRouteMapRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteMapRuleResponse, _err error) {
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

	if !dara.IsNil(request.RouteMapRuleId) {
		query["RouteMapRuleId"] = request.RouteMapRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRouteMapRule"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteMapRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DeleteRouteServiceInCen operation to delete the configuration of an Alibaba Cloud service from a Basic Edition transit router.
//
// Description:
//
// *DeleteRouteServiceInCen*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the operation continues in the background. The system returns a request ID even if you specify an invalid parameter. In this case, the Alibaba Cloud service configuration is not deleted. You can call the **DescribeRouteServicesInCen*	- operation to query the status of the Alibaba Cloud service.
//
// - If the Alibaba Cloud service is in the **Deleting*	- state, you can only query its configuration. You cannot perform other operations.
//
// - If the specified Alibaba Cloud service configuration is not found, the configuration has been deleted.
//
// @param request - DeleteRouteServiceInCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteServiceInCenResponse
func (client *Client) DeleteRouteServiceInCenWithContext(ctx context.Context, request *DeleteRouteServiceInCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteServiceInCenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessRegionId) {
		query["AccessRegionId"] = request.AccessRegionId
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.HostRegionId) {
		query["HostRegionId"] = request.HostRegionId
	}

	if !dara.IsNil(request.HostVpcId) {
		query["HostVpcId"] = request.HostVpcId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DeleteRouteServiceInCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteServiceInCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a traffic marking policy.
//
// Description:
//
// - **DeleteTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the traffic marking policy is not yet deleted. The deletion task continues to run in the background. You can call **ListTrafficMarkingPolicies*	- to query the status of the traffic marking policy.
//
//   - If the traffic marking policy is in the **Deleting*	- state, the traffic marking policy is being deleted. In this state, you can only query the traffic marking policy but cannot perform other operations on it.
//
//   - If the specified traffic marking policy cannot be found, the traffic marking policy is deleted.
//
// - Before you delete a traffic marking policy, delete all traffic classification rules from the traffic marking policy. For more information, see [RemoveTrafficMatchRuleFromTrafficMarkingPolicy](https://help.aliyun.com/document_detail/468330.html).
//
// @param request - DeleteTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficMarkingPolicyResponse
func (client *Client) DeleteTrafficMarkingPolicyWithContext(ctx context.Context, request *DeleteTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrafficMarkingPolicyResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyId) {
		query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficMarkingPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an aggregate route.
//
// Description:
//
// - Before you delete an aggregate route, make sure that redundant routes exist in the current network. Otherwise, service breaks may occur.
//
// - After you delete an aggregate route, the system automatically withdraws the aggregate routing that has been propagated to Virtual Private Cloud (VPC)-connected instances and re-propagates the specific routes within the destination CIDR block of the aggregation route to the VPC-connected instances.
//
// @param request - DeleteTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouteTableAggregationResponse
func (client *Client) DeleteTransitRouteTableAggregationWithContext(ctx context.Context, request *DeleteTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouteTableAggregationResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouteTableAggregationCidr) {
		query["TransitRouteTableAggregationCidr"] = request.TransitRouteTableAggregationCidr
	}

	if !dara.IsNil(request.TransitRouteTableId) {
		query["TransitRouteTableId"] = request.TransitRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouteTableAggregation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouteTableAggregationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Enterprise Edition transit router instance.
//
// Description:
//
// *DeleteTransitRouter*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the transit router instance is not immediately deleted. The deletion task runs in the background. You can call the **ListTransitRouters*	- operation to query the status of the transit router instance.
//
// - If the transit router instance is in the **Deleting*	- state, the transit router instance is being deleted. In this state, you can only query the transit router instance but cannot perform other operations on it.
//
// - If the transit router instance cannot be found, the transit router instance is deleted.
//
// ### Before you begin
//
// Before you delete a transit router instance, make sure that the following conditions are met:
//
// - No connections exist on the transit router instance.
//
//   - To delete a virtual private cloud (VPC) connection, see [DeleteTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261220.html).
//
//   - To delete an Express Connect Router (ECR) connection, see [DeleteTransitRouterEcrAttachment](https://help.aliyun.com/document_detail/2715447.html).
//
//   - To delete a Cloud Connect Network (CCN) connection, see [DetachCenChildInstance](https://help.aliyun.com/document_detail/65915.html).
//
//   - To delete a VPN connection, see [DeleteTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/443992.html).
//
//   - To delete an inter-region connection, see [DeleteTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261227.html).
//
//   - To delete a virtual border router (VBR) connection, see [DeleteTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/261223.html).
//
// - No custom route tables exist on the transit router instance. For more information, see [DeleteTransitRouterRouteTable](https://help.aliyun.com/document_detail/261235.html).
//
// - No custom route entries, route prefixes, or aggregate routes exist in the default route table of the transit router instance. For more information, see:
//
//   - To delete custom route entries of an Enterprise Edition transit router, see [DeleteTransitRouterRouteEntry](https://help.aliyun.com/document_detail/468291.html).
//
//   - To disassociate a prefix list, see [DeleteTransitRouterPrefixListAssociation](https://help.aliyun.com/document_detail/468312.html).
//
//   - To delete an aggregate route, see [DeleteTransitRouteTableAggregation](https://help.aliyun.com/document_detail/476070.html).
//
// - No multicast domains exist on the transit router instance. To delete a multicast domain, see [DeleteTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/468386.html).
//
// - No traffic marking policies exist on the transit router instance. To delete a traffic marking policy, see [DeleteTrafficMarkingPolicy](https://help.aliyun.com/document_detail/468324.html).
//
// @param request - DeleteTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterResponse
func (client *Client) DeleteTransitRouterWithContext(ctx context.Context, request *DeleteTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouter"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a CIDR block from a transit router.
//
// Description:
//
// A transit router CIDR block that has allocated IP addresses cannot be deleted.
//
// @param request - DeleteTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterCidrResponse
func (client *Client) DeleteTransitRouterCidrWithContext(ctx context.Context, request *DeleteTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterCidrResponse, _err error) {
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

	if !dara.IsNil(request.TransitRouterCidrId) {
		query["TransitRouterCidrId"] = request.TransitRouterCidrId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterCidr"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterCidrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect Router (ECR) connection from an Enterprise Edition transit router.
//
// Description:
//
// DeleteTransitRouterEcrAttachment is an asynchronous operation. After you call this operation, the system returns a RequestId, but the ECR connection is not immediately deleted. The deletion task runs in the background. You can call the [ListTransitRouterEcrAttachments](~~2361China~~) operation to query the status of the ECR connection.
//
// If the ECR connection is in the **Detaching*	- state, the ECR connection is being deleted. In this state, you can only query the ECR connection but cannot perform other operations on it.
//
// If the specified ECR connection cannot be found, the ECR connection is deleted.
//
// When you call the DeleteTransitRouterEcrAttachment operation, make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the system still returns a RequestId but does not delete the ECR connection from the Enterprise Edition transit router.
//
// @param request - DeleteTransitRouterEcrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterEcrAttachmentResponse
func (client *Client) DeleteTransitRouterEcrAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterEcrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterEcrAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterEcrAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterEcrAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a multicast domain by calling the DeleteTransitRouterMulticastDomain operation.
//
// Description:
//
// Before you delete a multicast domain, make sure that the following conditions are met:
//
// - The multicast domain is not associated with any vSwitches. For more information, see [DisassociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429774.html).
//
// - No multicast sources or multicast members exist in the multicast domain. For more information, see [DeregisterTransitRouterMulticastGroupSources](https://help.aliyun.com/document_detail/429776.html) and [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/429779.html).
//
// - The multicast domain is not associated with other multicast domains as a multicast member. You can delete the multicast member from other multicast domains to dissociate the other multicast domains from the current multicast domain. For more information, see [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/429779.html).
//
// - Make sure that the parameter values you specify are valid when you call this operation. If you specify invalid parameter values, the system still returns a RequestId but does not delete the multicast domain.
//
// @param request - DeleteTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterMulticastDomainResponse
func (client *Client) DeleteTransitRouterMulticastDomainWithContext(ctx context.Context, request *DeleteTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterMulticastDomainResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterMulticastDomain"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterMulticastDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteTransitRouterPeerAttachment operation to delete an inter-region connection from an Enterprise Edition transit router.
//
// Description:
//
// The **DeleteTransitRouterPeerAttachment*	- operation is asynchronous. After you send a request, the system returns a **RequestId*	- but the inter-region connection is not immediately deleted. The deletion task continues to run in the background. You can call the **ListTransitRouterPeerAttachments*	- operation to query the status of the inter-region connection.
//
// - If the inter-region connection is in the **Detaching*	- state, the inter-region connection is being deleted. In this state, you can only query the inter-region connection but cannot perform other operations on it.
//
// - If the specified inter-region connection cannot be found, the inter-region connection is deleted.
//
// When calling the **DeleteTransitRouterPeerAttachment*	- operation, make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the system still returns a **RequestId*	- but does not delete the inter-region connection from the Enterprise Edition transit router.
//
// @param request - DeleteTransitRouterPeerAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterPeerAttachmentResponse
func (client *Client) DeleteTransitRouterPeerAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterPeerAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterPeerAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterPeerAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterPeerAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates a prefix list from an Enterprise Edition transit router route table.
//
// Description:
//
//	Warning: After a prefix list is dissociated from an Enterprise Edition transit router route table, the system automatically withdraws all route entries related to the prefix list from the Enterprise Edition transit router route table. Before dissociating the prefix list, make sure that redundant routes exist in the Enterprise Edition transit router route table. Otherwise, network interruptions may occur.
//
// @param request - DeleteTransitRouterPrefixListAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterPrefixListAssociationResponse
func (client *Client) DeleteTransitRouterPrefixListAssociationWithContext(ctx context.Context, request *DeleteTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterPrefixListAssociationResponse, _err error) {
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

	if !dara.IsNil(request.NextHop) {
		query["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.NextHopType) {
		query["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PrefixListId) {
		query["PrefixListId"] = request.PrefixListId
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

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterTableId) {
		query["TransitRouterTableId"] = request.TransitRouterTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterPrefixListAssociation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterPrefixListAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteTransitRouterRouteEntry operation to delete static route entries of the blackhole or connection instance type from an Enterprise Edition transit router forward route table.
//
// Description:
//
// Before you call this operation to delete a route entry, take note of the following information:
//
// - If you specify **TransitRouterRouteEntryId*	- to delete a route entry, you do not need to specify **TransitRouterRouteTableId*	- or **TransitRouterRouteEntryDestinationCidrBlock**. Otherwise, a parameter conflict error occurs.
//
// - If you do not specify **TransitRouterRouteEntryId*	- to delete a route entry, specify the corresponding parameters based on the next hop type of the route entry:
//
//   - To delete a blackhole route, specify **TransitRouterRouteTableId**, **TransitRouterRouteEntryDestinationCidrBlock**, and **TransitRouterRouteEntryNextHopType**.
//
//   - To delete a non-blackhole route, specify **TransitRouterRouteTableId**, **TransitRouterRouteEntryDestinationCidrBlock**, **TransitRouterRouteEntryNextHopType**, and **TransitRouterRouteEntryNextHopId**.
//
// - **DeleteTransitRouterRouteEntry*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the route entry is not yet deleted. The deletion task runs in the background. You can call **ListTransitRouterRouteEntries*	- to query the status of the route entry.
//
//   - If the route entry is in the **Deleting*	- state, the route entry is being deleted. In this state, you can only query the route entry but cannot perform other operations on it.
//
//   - If the route entry cannot be found, the route entry is deleted.
//
// ### Limits
//
// This operation can delete only static route entries. Automatically learned route entries cannot be deleted. You can call [ListTransitRouterRouteEntries](https://help.aliyun.com/document_detail/260941.html) to query the type of a route entry.
//
// @param request - DeleteTransitRouterRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterRouteEntryResponse
func (client *Client) DeleteTransitRouterRouteEntryWithContext(ctx context.Context, request *DeleteTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterRouteEntryResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryDestinationCidrBlock) {
		query["TransitRouterRouteEntryDestinationCidrBlock"] = request.TransitRouterRouteEntryDestinationCidrBlock
	}

	if !dara.IsNil(request.TransitRouterRouteEntryId) {
		query["TransitRouterRouteEntryId"] = request.TransitRouterRouteEntryId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNextHopId) {
		query["TransitRouterRouteEntryNextHopId"] = request.TransitRouterRouteEntryNextHopId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNextHopType) {
		query["TransitRouterRouteEntryNextHopType"] = request.TransitRouterRouteEntryNextHopType
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterRouteEntry"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom route table from an Enterprise Edition transit router by calling the DeleteTransitRouterRouteTable operation.
//
// Description:
//
// - The default route table of an Enterprise Edition transit router cannot be deleted.
//
// - **DeleteTransitRouterRouteTable*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the custom route table is not yet deleted. The deletion task continues to run in the background. You can call **ListTransitRouterRouteTables*	- to query the status of the custom route table.
//
//   - If the custom route table is in the Deleting state, the custom route table is being deleted. In this state, you can only query the custom route table but cannot perform other operations on it.
//
//   - If the custom route table cannot be found, the custom route table is deleted.
//
// @param request - DeleteTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterRouteTableResponse
func (client *Client) DeleteTransitRouterRouteTableWithContext(ctx context.Context, request *DeleteTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterRouteTableResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterRouteTable"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterRouteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual border router (VBR) connection from an Enterprise Edition transit router.
//
// Description:
//
// *DeleteTransitRouterVbrAttachment*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the VBR connection is not yet deleted. The deletion task continues to run in the background. You can call **ListTransitRouterVbrAttachments*	- to query the status of the VBR connection.
//
// - If the VBR connection is in the **Detaching*	- state, the VBR connection is being deleted. In this state, you can only query the VBR connection but cannot perform other operations on it.
//
// - If the specified VBR connection cannot be found, the VBR connection is deleted.
//
// When calling the DeleteTransitRouterVbrAttachment operation, make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the system still returns a RequestId but does not delete the VBR connection from the Enterprise Edition transit router.
//
// @param request - DeleteTransitRouterVbrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVbrAttachmentResponse
func (client *Client) DeleteTransitRouterVbrAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterVbrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVbrAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterVbrAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterVbrAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteTransitRouterVpcAttachment operation to delete a virtual private cloud (VPC) connection from an Enterprise Edition transit router.
//
// Description:
//
// *DeleteTransitRouterVpcAttachment*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the VPC connection is not yet deleted. The deletion task continues to run in the background. You can call **ListTransitRouterVpcAttachments*	- to query the status of the VPC connection.
//
// - If the VPC connection is in the **Detaching*	- state, the VPC connection is being deleted. In this state, you can only query the VPC connection. You cannot perform other operations on the VPC connection.
//
// - If the VPC connection cannot be found, the VPC connection is deleted.
//
// When you call **DeleteTransitRouterVpcAttachment**, make sure that the parameter values you specify are valid. If you specify invalid parameter values, the system still returns a **RequestId*	- but does not delete the VPC connection from the Enterprise Edition transit router.
//
// ### Before you begin
//
// Before you delete a VPC connection, make sure that you have not configured a routing rule to access PrivateZone by using the VPC-connected instance. To delete the routing rule to PrivateZone, see [UnroutePrivateZoneInCenToVpc](https://help.aliyun.com/document_detail/468375.html).
//
// @param request - DeleteTransitRouterVpcAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVpcAttachmentResponse
func (client *Client) DeleteTransitRouterVpcAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterVpcAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVpcAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterVpcAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterVpcAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除VPN连接
//
// @param request - DeleteTransitRouterVpnAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVpnAttachmentResponse
func (client *Client) DeleteTransitRouterVpnAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterVpnAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVpnAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTransitRouterVpnAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTransitRouterVpnAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a multicast member from a multicast group if the member no longer needs to receive multicast traffic by calling the DeregisterTransitRouterMulticastGroupMembers operation.
//
// Description:
//
// `DeregisterTransitRouterMulticastGroupMembers` is an asynchronous operation. After you call this operation, the system returns a **RequestId*	- but the multicast member is not immediately removed. The removal task continues to run in the background. You can call `ListTransitRouterMulticastGroups` to query the status of the multicast member.
//
// - If the multicast member is in the **Deregistering*	- state, the multicast member is being removed. In this state, you can only query the multicast member but cannot perform other operations on it.
//
// - If the multicast member cannot be found in the multicast domain when you call the `ListTransitRouterMulticastGroups` operation, the multicast member has been removed.
//
// When you call the DeregisterTransitRouterMulticastGroupMembers operation, make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the system still returns a RequestId but does not remove the multicast member.
//
// @param request - DeregisterTransitRouterMulticastGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterTransitRouterMulticastGroupMembersResponse
func (client *Client) DeregisterTransitRouterMulticastGroupMembersWithContext(ctx context.Context, request *DeregisterTransitRouterMulticastGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *DeregisterTransitRouterMulticastGroupMembersResponse, _err error) {
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

	if !dara.IsNil(request.GroupIpAddress) {
		query["GroupIpAddress"] = request.GroupIpAddress
	}

	if !dara.IsNil(request.NetworkInterfaceIds) {
		query["NetworkInterfaceIds"] = request.NetworkInterfaceIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PeerTransitRouterMulticastDomains) {
		query["PeerTransitRouterMulticastDomains"] = request.PeerTransitRouterMulticastDomains
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeregisterTransitRouterMulticastGroupMembers"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeregisterTransitRouterMulticastGroupMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a multicast source from a multicast group if you no longer need the multicast source to send multicast traffic.
//
// Description:
//
// `DeregisterTransitRouterMulticastGroupSources` is an asynchronous operation. After you call this operation, the system returns a **RequestId*	- but the multicast source is not immediately deleted. The deletion task continues to run in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of the multicast source.
//
// - If the multicast source is in the **Deregistering*	- state, the multicast source is being deleted. In this state, you can only query the multicast source but cannot perform other operations on it.
//
// - If the `ListTransitRouterMulticastGroups` operation cannot find the multicast source in the multicast domain, the multicast source has been deleted.
//
// When you call the DeregisterTransitRouterMulticastGroupSources operation, make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the system still returns a RequestId but does not delete the multicast source.
//
// @param request - DeregisterTransitRouterMulticastGroupSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterTransitRouterMulticastGroupSourcesResponse
func (client *Client) DeregisterTransitRouterMulticastGroupSourcesWithContext(ctx context.Context, request *DeregisterTransitRouterMulticastGroupSourcesRequest, runtime *dara.RuntimeOptions) (_result *DeregisterTransitRouterMulticastGroupSourcesResponse, _err error) {
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

	if !dara.IsNil(request.GroupIpAddress) {
		query["GroupIpAddress"] = request.GroupIpAddress
	}

	if !dara.IsNil(request.NetworkInterfaceIds) {
		query["NetworkInterfaceIds"] = request.NetworkInterfaceIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeregisterTransitRouterMulticastGroupSources"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeregisterTransitRouterMulticastGroupSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a network instance (VPC, VBR, or CCN) attached to a Cloud Enterprise Network (CEN) instance, including the attachment status and network instance type.
//
// @param request - DescribeCenAttachedChildInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenAttachedChildInstanceAttributeResponse
func (client *Client) DescribeCenAttachedChildInstanceAttributeWithContext(ctx context.Context, request *DescribeCenAttachedChildInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenAttachedChildInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeCenAttachedChildInstanceAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenAttachedChildInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about network instances attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can use this operation to query network instances attached to a CEN instance in the following ways:
//
// - Specify `CenId` to query all network instances attached to the CEN instance.
//
// - Specify `CenId` and `ChildInstanceRegionId` to query network instances attached to the CEN instance in a specific region.
//
// - Specify `CenId` and `ChildInstanceType` to query network instances of a specific type attached to the CEN instance.
//
// @param request - DescribeCenAttachedChildInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenAttachedChildInstancesResponse
func (client *Client) DescribeCenAttachedChildInstancesWithContext(ctx context.Context, request *DescribeCenAttachedChildInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenAttachedChildInstancesResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeCenAttachedChildInstances"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenAttachedChildInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about bandwidth package instances by calling the DescribeCenBandwidthPackages operation.
//
// @param request - DescribeCenBandwidthPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenBandwidthPackagesResponse
func (client *Client) DescribeCenBandwidthPackagesWithContext(ctx context.Context, request *DescribeCenBandwidthPackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenBandwidthPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.IncludeReservationData) {
		query["IncludeReservationData"] = request.IncludeReservationData
	}

	if !dara.IsNil(request.IsOrKey) {
		query["IsOrKey"] = request.IsOrKey
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeCenBandwidthPackages"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenBandwidthPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DescribeCenChildInstanceRouteEntries operation to query the route entries of a network instance in a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenChildInstanceRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenChildInstanceRouteEntriesResponse
func (client *Client) DescribeCenChildInstanceRouteEntriesWithContext(ctx context.Context, request *DescribeCenChildInstanceRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenChildInstanceRouteEntriesResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceRouteTableId) {
		query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCenChildInstanceRouteEntries"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenChildInstanceRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remaining bandwidth of a specified bandwidth plan instance.
//
// @param request - DescribeCenGeographicSpanRemainingBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenGeographicSpanRemainingBandwidthResponse
func (client *Client) DescribeCenGeographicSpanRemainingBandwidthWithContext(ctx context.Context, request *DescribeCenGeographicSpanRemainingBandwidthRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenGeographicSpanRemainingBandwidthResponse, _err error) {
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

	if !dara.IsNil(request.GeographicRegionAId) {
		query["GeographicRegionAId"] = request.GeographicRegionAId
	}

	if !dara.IsNil(request.GeographicRegionBId) {
		query["GeographicRegionBId"] = request.GeographicRegionBId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeCenGeographicSpanRemainingBandwidth"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenGeographicSpanRemainingBandwidthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about connected areas supported by Cloud Enterprise Network (CEN) by calling the DescribeCenGeographicSpans operation.
//
// @param request - DescribeCenGeographicSpansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenGeographicSpansResponse
func (client *Client) DescribeCenGeographicSpansWithContext(ctx context.Context, request *DescribeCenGeographicSpansRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenGeographicSpansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GeographicSpanId) {
		query["GeographicSpanId"] = request.GeographicSpanId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeCenGeographicSpans"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenGeographicSpansResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the inter-region bandwidth information between regions by calling the DescribeCenInterRegionBandwidthLimits operation.
//
// @param request - DescribeCenInterRegionBandwidthLimitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenInterRegionBandwidthLimitsResponse
func (client *Client) DescribeCenInterRegionBandwidthLimitsWithContext(ctx context.Context, request *DescribeCenInterRegionBandwidthLimitsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenInterRegionBandwidthLimitsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrRegionId) {
		query["TrRegionId"] = request.TrRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCenInterRegionBandwidthLimits"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenInterRegionBandwidthLimitsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the PrivateZone service configuration of a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenPrivateZoneRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenPrivateZoneRoutesResponse
func (client *Client) DescribeCenPrivateZoneRoutesWithContext(ctx context.Context, request *DescribeCenPrivateZoneRoutesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenPrivateZoneRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessRegionId) {
		query["AccessRegionId"] = request.AccessRegionId
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.HostRegionId) {
		query["HostRegionId"] = request.HostRegionId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeCenPrivateZoneRoutes"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenPrivateZoneRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of route entries in a specific region of a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenRegionDomainRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenRegionDomainRouteEntriesResponse
func (client *Client) DescribeCenRegionDomainRouteEntriesWithContext(ctx context.Context, request *DescribeCenRegionDomainRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenRegionDomainRouteEntriesResponse, _err error) {
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

	if !dara.IsNil(request.CenRegionId) {
		query["CenRegionId"] = request.CenRegionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCenRegionDomainRouteEntries"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenRegionDomainRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configured information in route maps by calling the DescribeCenRouteMaps operation.
//
// @param request - DescribeCenRouteMapsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenRouteMapsResponse
func (client *Client) DescribeCenRouteMapsWithContext(ctx context.Context, request *DescribeCenRouteMapsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenRouteMapsResponse, _err error) {
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

	if !dara.IsNil(request.CenRegionId) {
		query["CenRegionId"] = request.CenRegionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	if !dara.IsNil(request.TransmitDirection) {
		query["TransmitDirection"] = request.TransmitDirection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCenRouteMaps"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenRouteMapsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries health check information about virtual border routers (VBRs) in a specified region.
//
// @param request - DescribeCenVbrHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenVbrHealthCheckResponse
func (client *Client) DescribeCenVbrHealthCheckWithContext(ctx context.Context, request *DescribeCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenVbrHealthCheckResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VbrInstanceId) {
		query["VbrInstanceId"] = request.VbrInstanceId
	}

	if !dara.IsNil(request.VbrInstanceOwnerId) {
		query["VbrInstanceOwnerId"] = request.VbrInstanceOwnerId
	}

	if !dara.IsNil(request.VbrInstanceRegionId) {
		query["VbrInstanceRegionId"] = request.VbrInstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCenVbrHealthCheck"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCenVbrHealthCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about Cloud Enterprise Network (CEN) instances under the current Alibaba Cloud account, including the instance status, whether IPv6 is enabled, and the list of bandwidth packages associated with the instances.
//
// @param request - DescribeCensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCensResponse
func (client *Client) DescribeCensWithContext(ctx context.Context, request *DescribeCensRequest, runtime *dara.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeCens"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCensResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions that support network instances loaded to Cloud Enterprise Network (CEN).
//
// Description:
//
// CEN supports different regions for different network instance types. You can specify the ProductType parameter to query the regions supported by CEN for a specific network instance type. If you do not specify the ProductType parameter, the system queries the regions supported by CEN for all network instance types by default.
//
// @param request - DescribeChildInstanceRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChildInstanceRegionsResponse
func (client *Client) DescribeChildInstanceRegionsWithContext(ctx context.Context, request *DescribeChildInstanceRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChildInstanceRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
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
		Action:      dara.String("DescribeChildInstanceRegions"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChildInstanceRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries flow logs.
//
// @param request - DescribeFlowlogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowlogsResponse
func (client *Client) DescribeFlowlogsWithContext(ctx context.Context, request *DescribeFlowlogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowlogsResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.FlowLogName) {
		query["FlowLogName"] = request.FlowLogName
	}

	if !dara.IsNil(request.FlowLogVersion) {
		query["FlowLogVersion"] = request.FlowLogVersion
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.LogStoreName) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowlogs"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowlogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries region information within a specified area.
//
// @param request - DescribeGeographicRegionMembershipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGeographicRegionMembershipResponse
func (client *Client) DescribeGeographicRegionMembershipWithContext(ctx context.Context, request *DescribeGeographicRegionMembershipRequest, runtime *dara.RuntimeOptions) (_result *DescribeGeographicRegionMembershipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GeographicRegionId) {
		query["GeographicRegionId"] = request.GeographicRegionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeGeographicRegionMembership"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGeographicRegionMembershipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about cross-account network instances that are authorized to be associated with a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// When you call the **DescribeGrantRulesToCen*	- operation, make sure that the parameter values you specify are valid. If you specify invalid parameter values, a **RequestId*	- is still returned, but information about the cross-account network instances authorized to the CEN instance is not returned.
//
// @param request - DescribeGrantRulesToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantRulesToCenResponse
func (client *Client) DescribeGrantRulesToCenWithContext(ctx context.Context, request *DescribeGrantRulesToCenRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantRulesToCenResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceOwnerId) {
		query["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	}

	if !dara.IsNil(request.EnabledIpv6) {
		query["EnabledIpv6"] = request.EnabledIpv6
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

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
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
		Action:      dara.String("DescribeGrantRulesToCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGrantRulesToCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization information of a network instance for a cross-account Cloud Enterprise Network (CEN) instance, including the Alibaba Cloud account ID of the CEN instance owner and the payer of the network instance.
//
// @param request - DescribeGrantRulesToResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantRulesToResourceResponse
func (client *Client) DescribeGrantRulesToResourceWithContext(ctx context.Context, request *DescribeGrantRulesToResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantRulesToResourceResponse, _err error) {
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

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGrantRulesToResource"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGrantRulesToResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the publication status in Cloud Enterprise Network (CEN) for routes loaded to CEN network instances such as VPCs and VBRs, the instance type associated with the next hop of each route entry, and whether the route can be published to or withdrawn from CEN.
//
// @param request - DescribePublishedRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePublishedRouteEntriesResponse
func (client *Client) DescribePublishedRouteEntriesWithContext(ctx context.Context, request *DescribePublishedRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribePublishedRouteEntriesResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceRouteTableId) {
		query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribePublishedRouteEntries"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePublishedRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DescribeRouteConflict operation to query the routing information about route entry conflicts in a network instance.
//
// @param request - DescribeRouteConflictRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteConflictResponse
func (client *Client) DescribeRouteConflictWithContext(ctx context.Context, request *DescribeRouteConflictRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteConflictResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceRouteTableId) {
		query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeRouteConflict"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRouteConflictResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud service configurations under a Basic Edition transit router by calling the DescribeRouteServicesInCen operation.
//
// @param request - DescribeRouteServicesInCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteServicesInCenResponse
func (client *Client) DescribeRouteServicesInCenWithContext(ctx context.Context, request *DescribeRouteServicesInCenRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteServicesInCenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessRegionId) {
		query["AccessRegionId"] = request.AccessRegionId
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.HostRegionId) {
		query["HostRegionId"] = request.HostRegionId
	}

	if !dara.IsNil(request.HostVpcId) {
		query["HostVpcId"] = request.HostVpcId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeRouteServicesInCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRouteServicesInCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the aggregate routes in the route table of an Enterprise Edition transit router.
//
// Description:
//
// You can query the information about a specific aggregate route by specifying the **TransitRouteTableId*	- and **TransitRouteTableAggregationCidr*	- parameters. If you specify only the **TransitRouteTableId*	- parameter, the system queries the information about all aggregate routes in the route table of the Enterprise Edition transit router by default.
//
// @param request - DescribeTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransitRouteTableAggregationResponse
func (client *Client) DescribeTransitRouteTableAggregationWithContext(ctx context.Context, request *DescribeTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransitRouteTableAggregationResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouteTableAggregationCidr) {
		query["TransitRouteTableAggregationCidr"] = request.TransitRouteTableAggregationCidr
	}

	if !dara.IsNil(request.TransitRouteTableId) {
		query["TransitRouteTableId"] = request.TransitRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTransitRouteTableAggregation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTransitRouteTableAggregationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration details of a specified aggregate route.
//
// @param request - DescribeTransitRouteTableAggregationDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransitRouteTableAggregationDetailResponse
func (client *Client) DescribeTransitRouteTableAggregationDetailWithContext(ctx context.Context, request *DescribeTransitRouteTableAggregationDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransitRouteTableAggregationDetailResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouteTableAggregationCidr) {
		query["TransitRouteTableAggregationCidr"] = request.TransitRouteTableAggregationCidr
	}

	if !dara.IsNil(request.TransitRouteTableId) {
		query["TransitRouteTableId"] = request.TransitRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTransitRouteTableAggregationDetail"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTransitRouteTableAggregationDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource reservation configuration of a specified transit router.
//
// @param request - DescribeTransitRouterCapacityReservationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransitRouterCapacityReservationsResponse
func (client *Client) DescribeTransitRouterCapacityReservationsWithContext(ctx context.Context, request *DescribeTransitRouterCapacityReservationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransitRouterCapacityReservationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTransitRouterCapacityReservations"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTransitRouterCapacityReservationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a network instance from a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// The **DetachCenChildInstance*	- operation supports detaching only specified network instances from Basic Edition transit routers.
//
// @param request - DetachCenChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachCenChildInstanceResponse
func (client *Client) DetachCenChildInstanceWithContext(ctx context.Context, request *DetachCenChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DetachCenChildInstanceResponse, _err error) {
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

	if !dara.IsNil(request.CenOwnerId) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceOwnerId) {
		query["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DetachCenChildInstance"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachCenChildInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the health check for a specified virtual border router (VBR).
//
// Description:
//
// The **DisableCenVbrHealthCheck*	- operation is asynchronous. The system returns a **RequestId*	- before the health check configuration is completely deleted, and the deletion task continues in the background. You can call the **DescribeCenVbrHealthCheck*	- operation to query the health check configuration. If the specified health check configuration cannot be found, the configuration is deleted.
//
// @param request - DisableCenVbrHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCenVbrHealthCheckResponse
func (client *Client) DisableCenVbrHealthCheckWithContext(ctx context.Context, request *DisableCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *DisableCenVbrHealthCheckResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VbrInstanceId) {
		query["VbrInstanceId"] = request.VbrInstanceId
	}

	if !dara.IsNil(request.VbrInstanceOwnerId) {
		query["VbrInstanceOwnerId"] = request.VbrInstanceOwnerId
	}

	if !dara.IsNil(request.VbrInstanceRegionId) {
		query["VbrInstanceRegionId"] = request.VbrInstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableCenVbrHealthCheck"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableCenVbrHealthCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables route learning between a network instance connection and a transit router route table.
//
// Description:
//
// *DisableTransitRouterRouteTablePropagation*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- before the route learning relationship between the network instance connection and the route table is fully removed. The removal task continues to run in the background. You can call **ListTransitRouterRouteTablePropagations*	- to query the route learning relationship between the network instance connection and the route table.
//
// - If the route learning relationship is in the **Disabling*	- state, the network instance connection and the route table are being disassociated. In this state, you can only query the route learning relationship. You cannot perform other operations.
//
// - If the **ListTransitRouterRouteTableAssociations*	- operation does not return the route learning relationship between the network instance connection and the route table, the route learning relationship is successfully removed.
//
// @param request - DisableTransitRouterRouteTablePropagationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableTransitRouterRouteTablePropagationResponse
func (client *Client) DisableTransitRouterRouteTablePropagationWithContext(ctx context.Context, request *DisableTransitRouterRouteTablePropagationRequest, runtime *dara.RuntimeOptions) (_result *DisableTransitRouterRouteTablePropagationResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableTransitRouterRouteTablePropagation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableTransitRouterRouteTablePropagationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a vSwitch from a multicast domain.
//
// Description:
//
// - Before dissociating a vSwitch from a multicast domain, make sure that no multicast sources or multicast members exist under the vSwitch. To delete multicast sources and multicast members, see [DeregisterTransitRouterMulticastGroupSources](https://help.aliyun.com/document_detail/468416.html) and [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/468409.html).
//
// - If you specify invalid parameters, the system still returns a RequestId but does not dissociate the vSwitch from the multicast domain.
//
// - **DisassociateTransitRouterMulticastDomain*	- is an asynchronous operation. After you invoke this operation, the system returns a **RequestId*	- but the dissociation has not yet completed. The dissociation node continues to run in the background. You can invoke **ListTransitRouterMulticastDomainAssociations*	- to query the associate status between the vSwitch and the multicast domain.
//
//   - If the associate status is **Dissociating**, the vSwitch is being dissociated from the multicast domain. In this state, you can only execute query operations on the vSwitch but cannot execute other operations.
//
//   - If the vSwitch information cannot be found under the multicast domain, the vSwitch has been successfully dissociated from the multicast domain.
//
// @param request - DisassociateTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateTransitRouterMulticastDomainResponse
func (client *Client) DisassociateTransitRouterMulticastDomainWithContext(ctx context.Context, request *DisassociateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *DisassociateTransitRouterMulticastDomainResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateTransitRouterMulticastDomain"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateTransitRouterMulticastDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates a network instance connection from a route table.
//
// Description:
//
// *DissociateTransitRouterAttachmentFromRouteTable*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the dissociation between the network instance connection and the route table is not complete. The dissociation task continues to run in the background. You can call **ListTransitRouterRouteTableAssociations*	- to query the association status between the network instance connection and the route table.
//
// - If the association status is **Dissociating**, the network instance connection is being dissociated from the route table. In this state, you can only query the forwarding association between the network instance connection and the route table. You cannot perform other operations.
//
// - If the **ListTransitRouterRouteTableAssociations*	- operation does not return the forwarding association between the network instance connection and the route table, the dissociation is successful.
//
// @param request - DissociateTransitRouterAttachmentFromRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateTransitRouterAttachmentFromRouteTableResponse
func (client *Client) DissociateTransitRouterAttachmentFromRouteTableWithContext(ctx context.Context, request *DissociateTransitRouterAttachmentFromRouteTableRequest, runtime *dara.RuntimeOptions) (_result *DissociateTransitRouterAttachmentFromRouteTableResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateTransitRouterAttachmentFromRouteTable"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateTransitRouterAttachmentFromRouteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the health check feature for a virtual border router (VBR) to detect faults on physical Express Connect circuits in a timely manner. You can also use this operation to modify the health check configuration of a VBR.
//
// Description:
//
// You can configure health checks for a VBR instance to monitor the connectivity of the physical Express Connect circuit between your on-premises data center and Alibaba Cloud, so that issues can be detected promptly.
//
// Before using the health check feature, note the following information:
//
// - If your VBR instance uses static routing, after you configure the health check, you must add a static route entry in the on-premises data center connected to the VBR instance.
//
//	The destination CIDR block of the static route is the source IP address of the health check with a 32-bit subnet mask, and the next hop is the Alibaba Cloud-side IP address of the VBR instance.
//
// - If your border router instance uses the BGP dynamic routing protocol, you do not need to add a route entry in the on-premises data center.
//
// - The **EnableCenVbrHealthCheck*	- operation is asynchronous. After you send a request, the system returns a **RequestId**, but the health check instance is not yet created or modified. The creation or modification task continues to run in the background. You can call the **DescribeCenVbrHealthCheck*	- operation to query the health check configuration. If the health check configuration is returned, the health check has been created or modified.
//
// @param request - EnableCenVbrHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCenVbrHealthCheckResponse
func (client *Client) EnableCenVbrHealthCheckWithContext(ctx context.Context, request *EnableCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *EnableCenVbrHealthCheckResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckOnly) {
		query["HealthCheckOnly"] = request.HealthCheckOnly
	}

	if !dara.IsNil(request.HealthCheckSourceIp) {
		query["HealthCheckSourceIp"] = request.HealthCheckSourceIp
	}

	if !dara.IsNil(request.HealthCheckTargetIp) {
		query["HealthCheckTargetIp"] = request.HealthCheckTargetIp
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VbrInstanceId) {
		query["VbrInstanceId"] = request.VbrInstanceId
	}

	if !dara.IsNil(request.VbrInstanceOwnerId) {
		query["VbrInstanceOwnerId"] = request.VbrInstanceOwnerId
	}

	if !dara.IsNil(request.VbrInstanceRegionId) {
		query["VbrInstanceRegionId"] = request.VbrInstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableCenVbrHealthCheck"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableCenVbrHealthCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route learning relationship.
//
// Description:
//
// After you create a network instance connection, you can configure a route learning relationship for it. After the configuration, the network instance connection automatically propagates routes of the network instance to its associated route table. Before you call this operation to create a route learning relationship, take note of the following information:
//
// - Only Enterprise Edition transit routers support creating route learning relationships. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - The **EnableTransitRouterRouteTablePropagation*	- operation is asynchronous. After you call the operation, the system returns a **RequestId**, but the route learning relationship between the network instance connection and the route table is not fully established, and the creation task is still running in the background. You can call the **ListTransitRouterRouteTablePropagations*	- operation to query the route learning relationship between the network instance connection and the route table.
//
//   - When the route learning relationship is in the **Enabling*	- state, the route learning relationship between the network instance connection and the route table is being established. In this state, you can only query the route learning relationship between the network instance connection and the route table, and cannot perform other operations.
//
//   - When the route learning relationship is in the **Active*	- state, the route learning relationship between the network instance connection and the route table is successfully established.
//
// @param request - EnableTransitRouterRouteTablePropagationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableTransitRouterRouteTablePropagationResponse
func (client *Client) EnableTransitRouterRouteTablePropagationWithContext(ctx context.Context, request *EnableTransitRouterRouteTablePropagationRequest, runtime *dara.RuntimeOptions) (_result *EnableTransitRouterRouteTablePropagationResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableTransitRouterRouteTablePropagation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableTransitRouterRouteTablePropagationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions to a transit router instance to connect to a network instance that belongs to a different Alibaba Cloud account. Before a transit router instance can connect to a network instance that belongs to a different account, the network instance owner must authorize the transit router instance by calling the GrantInstanceToTransitRouter operation.
//
// Description:
//
// - The GrantInstanceToTransitRouter operation only supports granting a transit router instance permissions to connect to cross-account Virtual Private Cloud (VPC) instances, Virtual Border Router (VBR) instances, IPsec connections, and Express Connect Router (ECR) instances.
//
//	To grant a transit router instance permissions to connect to a Cloud Connect Network (CCN) instance, call the [GrantInstanceToCbn](https://help.aliyun.com/document_detail/126141.html) operation.
//
// - Before you call the GrantInstanceToTransitRouter operation, make sure that you understand the billing rules of transit routers, the limits on authorization operations, and the prerequisites for authorization operations. For more information, see [Cross-account network instance authorization](https://help.aliyun.com/document_detail/181553.html).
//
// - Before you authorize a network instance, make sure that the following operations are completed:
//
//	Confirm that the account to which the network instance belongs and the account to which the transit router instance belongs are of the same type.
//
//	Obtain the Alibaba Cloud account ID of the account to which the transit router instance belongs.
//
//	Obtain the Cloud Enterprise Network (CEN) instance ID to which the transit router instance belongs.
//
//	Before you authorize a VBR instance, contact your account manager to activate the VBR instance authorization feature.
//
//	Before you authorize an IPsec connection, make sure that the IPsec connection is not associated with any resource:
//
//	If the IPsec connection is already associated with a VPN gateway instance, it cannot be associated with a transit router instance in the same account or a different account.
//
//	If the IPsec connection is already associated with a transit router instance, you must disassociate it first. For more information, see [Delete a network instance connection](https://help.aliyun.com/document_detail/181554.html).
//
// @param request - GrantInstanceToTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantInstanceToTransitRouterResponse
func (client *Client) GrantInstanceToTransitRouterWithContext(ctx context.Context, request *GrantInstanceToTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *GrantInstanceToTransitRouterResponse, _err error) {
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

	if !dara.IsNil(request.CenOwnerId) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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
		Action:      dara.String("GrantInstanceToTransitRouter"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantInstanceToTransitRouterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries route entries that forward traffic to a network instance connection in the route table of a network instance associated with an Enterprise Edition transit router.
//
// Description:
//
// When you call the ListCenChildInstanceRouteEntriesToAttachment operation, make sure that the parameter values you specify are valid. If you specify invalid parameter values, the operation returns a RequestId but does not display the route entries of network instances connected to the Enterprise Edition transit router.
//
// @param request - ListCenChildInstanceRouteEntriesToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenChildInstanceRouteEntriesToAttachmentResponse
func (client *Client) ListCenChildInstanceRouteEntriesToAttachmentWithContext(ctx context.Context, request *ListCenChildInstanceRouteEntriesToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *ListCenChildInstanceRouteEntriesToAttachmentResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceRouteTableId) {
		query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteFilter) {
		query["RouteFilter"] = request.RouteFilter
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCenChildInstanceRouteEntriesToAttachment"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCenChildInstanceRouteEntriesToAttachmentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about traffic scheduling policies by calling the ListCenInterRegionTrafficQosPolicies operation.
//
// Description:
//
// When you call the **ListCenInterRegionTrafficQosPolicies*	- operation:
//
// - Specify at least one of the **TransitRouterId*	- and **TrafficQosPolicyId*	- parameters.
//
// - If you do not specify a traffic scheduling policy ID (that is, you do not specify the **TrafficQosPolicyId*	- parameter), the operation returns only the traffic scheduling policy information based on the values of the **TransitRouterId**, **TransitRouterAttachmentId**, **TrafficQosPolicyName**, and **TrafficQosPolicyDescription*	- parameters. The queue information under the traffic scheduling policy is not returned (that is, the response does not include the **TrafficQosQueues*	- field).
//
// - If you specify a traffic scheduling policy ID (that is, you specify the **TrafficQosPolicyId*	- parameter), the operation returns the traffic scheduling policy information and the queue information under the traffic scheduling policy (that is, the response includes the **TrafficQosQueues*	- field). If the **TrafficQosQueues*	- field is an empty array, only the default queue exists under the traffic scheduling policy.
//
// - Make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the operation still returns a RequestId but does not return traffic scheduling policy information.
//
// @param request - ListCenInterRegionTrafficQosPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenInterRegionTrafficQosPoliciesResponse
func (client *Client) ListCenInterRegionTrafficQosPoliciesWithContext(ctx context.Context, request *ListCenInterRegionTrafficQosPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListCenInterRegionTrafficQosPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficQosPolicyDescription) {
		query["TrafficQosPolicyDescription"] = request.TrafficQosPolicyDescription
	}

	if !dara.IsNil(request.TrafficQosPolicyId) {
		query["TrafficQosPolicyId"] = request.TrafficQosPolicyId
	}

	if !dara.IsNil(request.TrafficQosPolicyName) {
		query["TrafficQosPolicyName"] = request.TrafficQosPolicyName
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCenInterRegionTrafficQosPolicies"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCenInterRegionTrafficQosPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about traffic scheduling policy queues by calling the ListCenInterRegionTrafficQosQueues operation.
//
// Description:
//
// When you call this operation, specify at least one of the following parameters: **TransitRouterId**, **TrafficQosPolicyId**, or **TrafficQosQueueId**.
//
// Make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the system returns a **RequestId*	- but does not return information about traffic scheduling policies.
//
// @param request - ListCenInterRegionTrafficQosQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenInterRegionTrafficQosQueuesResponse
func (client *Client) ListCenInterRegionTrafficQosQueuesWithContext(ctx context.Context, request *ListCenInterRegionTrafficQosQueuesRequest, runtime *dara.RuntimeOptions) (_result *ListCenInterRegionTrafficQosQueuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveBandwidthFilter) {
		query["EffectiveBandwidthFilter"] = request.EffectiveBandwidthFilter
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficQosPolicyId) {
		query["TrafficQosPolicyId"] = request.TrafficQosPolicyId
	}

	if !dara.IsNil(request.TrafficQosQueueDescription) {
		query["TrafficQosQueueDescription"] = request.TrafficQosQueueDescription
	}

	if !dara.IsNil(request.TrafficQosQueueId) {
		query["TrafficQosQueueId"] = request.TrafficQosQueueId
	}

	if !dara.IsNil(request.TrafficQosQueueName) {
		query["TrafficQosQueueName"] = request.TrafficQosQueueName
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCenInterRegionTrafficQosQueues"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCenInterRegionTrafficQosQueuesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the elastic network interfaces (ENIs) within a virtual private cloud (VPC) that can serve as multicast sources or multicast members for multicast communication.
//
// Description:
//
// Before you invoke the `ListGrantVSwitchEnis` operation, make sure that the VPC-connected instance is connected to Cloud Enterprise Network (CEN). For more information, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261358.html).
//
// @param request - ListGrantVSwitchEnisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGrantVSwitchEnisResponse
func (client *Client) ListGrantVSwitchEnisWithContext(ctx context.Context, request *ListGrantVSwitchEnisRequest, runtime *dara.RuntimeOptions) (_result *ListGrantVSwitchEnisResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkInterfaceId) {
		query["NetworkInterfaceId"] = request.NetworkInterfaceId
	}

	if !dara.IsNil(request.NetworkInterfaceName) {
		query["NetworkInterfaceName"] = request.NetworkInterfaceName
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

	if !dara.IsNil(request.PrimaryIpAddress) {
		query["PrimaryIpAddress"] = request.PrimaryIpAddress
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("ListGrantVSwitchEnis"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGrantVSwitchEnisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about vSwitches in a cross-account virtual private cloud (VPC) that is connected to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// Before you invoke the `ListGrantVSwitchesToCen` operation, make sure that the CEN instance has been granted authorization to access the cross-account VPC-connected instance. For more information, see [GrantInstanceToCen](https://help.aliyun.com/document_detail/126224.html).
//
// @param request - ListGrantVSwitchesToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGrantVSwitchesToCenResponse
func (client *Client) ListGrantVSwitchesToCenWithContext(ctx context.Context, request *ListGrantVSwitchesToCenRequest, runtime *dara.RuntimeOptions) (_result *ListGrantVSwitchesToCenResponse, _err error) {
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

	if !dara.IsNil(request.EnabledIpv6) {
		query["EnabledIpv6"] = request.EnabledIpv6
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGrantVSwitchesToCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGrantVSwitchesToCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more route map rules.
//
// Description:
//
// ## Request description
//
// - RouteMapId: specifies the route map whose rules you want to query.
//
// - RouteMapRuleIds: the list of rule IDs to query. This parameter and filter conditions can be used together for matching, and they have an OR relationship.
//
// - Filter: the filter. You can filter by attributes such as direction. Supported filter conditions include Status and Direction.
//
// - DryRun: specifies whether to enable the dry run mode. Default value: false.
//
// You must specify at least one of RouteMapId and RouteMapRuleIds. If both are provided, a joint query is performed based on these two parameters. If only RouteMapId is provided, all rules under the route map are returned. If only RouteMapRuleIds is provided, the details of the specified rules are returned.
//
// @param request - ListRouteMapRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRouteMapRulesResponse
func (client *Client) ListRouteMapRulesWithContext(ctx context.Context, request *ListRouteMapRulesRequest, runtime *dara.RuntimeOptions) (_result *ListRouteMapRulesResponse, _err error) {
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

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	if !dara.IsNil(request.RouteMapRuleIds) {
		query["RouteMapRuleIds"] = request.RouteMapRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRouteMapRules"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRouteMapRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more route maps.
//
// Description:
//
// ## Request description
//
// - **RouteMapIds**: optional. Specifies a list of one or more route map IDs to query. If you do not specify this parameter, all route maps are queried by default.
//
// - **Filter**: optional. Filters route maps based on specific conditions such as the name and direction.
//
// - **ResourceGroupId**: optional. Specifies the ID of the resource group to further filter the results.
//
// - **DryRun**: optional. Specifies whether to check the request without actually performing the operation.
//
// - **MaxResults*	- and **NextToken**: pagination parameters. Used to control the number of returned results and obtain the next page of data.
//
// Note: If you specify both `RouteMapIds` and `Filter`, they are connected by an OR logical operator. Route maps that meet either condition are queried.
//
// @param request - ListRouteMapsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRouteMapsResponse
func (client *Client) ListRouteMapsWithContext(ctx context.Context, request *ListRouteMapsRequest, runtime *dara.RuntimeOptions) (_result *ListRouteMapsResponse, _err error) {
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

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RouteMapIds) {
		query["RouteMapIds"] = request.RouteMapIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRouteMaps"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRouteMapsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags bound to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// When you call the ListTagResources operation, you must specify at least one of the request parameters **ResourceId.N*	- and **Tag.N.Key**:
//
// - If you specify only **ResourceId.N**, the tags bound to the specified CEN instance are queried.
//
// - If you specify only **Tag.N.Key**, all CEN instances that have the specified tag key bound are queried.
//
// - If you specify both **ResourceId.N*	- and **Tag.N.Key**, the specified tags bound to the specified CEN instance are queried.
//
//   - Make sure that the values of **ResourceId.N*	- and **Tag.N.Key*	- correspond to each other. Otherwise, an empty result is returned.
//
//   - If you specify multiple tag keys, the tag keys are evaluated by using the logical AND operator.
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Version:     dara.String("2017-09-12"),
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
// Queries details about traffic marking policies, such as the status and priority of the traffic marking policies.
//
// Description:
//
// When you call the **ListTrafficMarkingPolicies*	- operation:
//
// - You must specify at least one of the **TransitRouterId*	- and **TrafficMarkingPolicyId*	- parameters.
//
// - If you do not specify a traffic marking policy ID (that is, you do not specify a value for the **TrafficMarkingPolicyId*	- parameter), the operation returns only the traffic marking policy information based on the values of the **TransitRouterId**, **TrafficMarkingPolicyName**, and **TrafficMarkingPolicyDescription*	- parameters. The traffic classification rule information under the traffic marking policy is not returned (that is, the response does not contain the **TrafficMatchRules*	- field).
//
// - If you specify a traffic marking policy ID (that is, you specify a value for the **TrafficMarkingPolicyId*	- parameter), the operation returns the traffic marking policy information and the traffic classification rule information under the traffic marking policy (that is, the response contains the **TrafficMatchRules*	- field).
//
// If the **TrafficMatchRules*	- field is an empty array, no traffic classification rules exist under the current traffic marking policy.
//
// @param request - ListTrafficMarkingPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficMarkingPoliciesResponse
func (client *Client) ListTrafficMarkingPoliciesWithContext(ctx context.Context, request *ListTrafficMarkingPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListTrafficMarkingPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyDescription) {
		query["TrafficMarkingPolicyDescription"] = request.TrafficMarkingPolicyDescription
	}

	if !dara.IsNil(request.TrafficMarkingPolicyId) {
		query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyName) {
		query["TrafficMarkingPolicyName"] = request.TrafficMarkingPolicyName
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrafficMarkingPolicies"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrafficMarkingPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones supported by Enterprise Edition transit routers in a specified region.
//
// Description:
//
// - You can invoke **ListTransitRouterAvailableResource*	- to query information about regular zones or zones that support the multicast feature for Enterprise Edition transit routers in a specified region.
//
//   - If you do not set **SupportMulticast*	- to **true**, the system queries only the regular zones supported by Enterprise Edition transit routers by default.
//
//   - If you set **SupportMulticast*	- to **true**, the system queries only the zones that support the multicast feature for Enterprise Edition transit routers.
//
// - On May 31, 2022, Cloud Enterprise Network (CEN) performed an optimization upgrade on the mode in which Enterprise Edition transit routers connect to virtual private clouds (VPCs). After the upgrade, you no longer need to specify primary and secondary zones when connecting an Enterprise Edition transit router to a VPC-connected instance. You can specify one or more zones.
//
//   - If your Enterprise Edition transit router has not been upgraded, you must specify primary and secondary zones when connecting the Enterprise Edition transit router to a VPC-connected instance. After you invoke **ListTransitRouterAvailableResource**, you can obtain the primary and secondary zone information from the **MasterZones*	- and **SlaveZones*	- parameters.
//
//   - If your Enterprise Edition transit router has been upgraded, you can specify any zones when connecting the Enterprise Edition transit router to a VPC-connected instance. After you invoke **ListTransitRouterAvailableResource**, you can obtain the zone information supported by the Enterprise Edition transit router from the **AvailableZones*	- parameter.
//
// For more information about the Enterprise Edition transit router upgrade, see [Upgrade the mode in which an Enterprise Edition transit router connects to a VPC](https://help.aliyun.com/document_detail/434191.html).
//
// @param request - ListTransitRouterAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterAvailableResourceResponse
func (client *Client) ListTransitRouterAvailableResourceWithContext(ctx context.Context, request *ListTransitRouterAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterAvailableResourceResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SupportMulticast) {
		query["SupportMulticast"] = request.SupportMulticast
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterAvailableResource"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterAvailableResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about CIDR blocks of a transit router by calling the ListTransitRouterCidr operation.
//
// @param request - ListTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterCidrResponse
func (client *Client) ListTransitRouterCidrWithContext(ctx context.Context, request *ListTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterCidrResponse, _err error) {
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

	if !dara.IsNil(request.TransitRouterCidrId) {
		query["TransitRouterCidrId"] = request.TransitRouterCidrId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterCidr"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterCidrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the allocation information of a CIDR block by calling the ListTransitRouterCidrAllocation operation.
//
// @param request - ListTransitRouterCidrAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterCidrAllocationResponse
func (client *Client) ListTransitRouterCidrAllocationWithContext(ctx context.Context, request *ListTransitRouterCidrAllocationRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterCidrAllocationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentId) {
		query["AttachmentId"] = request.AttachmentId
	}

	if !dara.IsNil(request.AttachmentName) {
		query["AttachmentName"] = request.AttachmentName
	}

	if !dara.IsNil(request.Cidr) {
		query["Cidr"] = request.Cidr
	}

	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DedicatedOwnerId) {
		query["DedicatedOwnerId"] = request.DedicatedOwnerId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
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

	if !dara.IsNil(request.TransitRouterCidrId) {
		query["TransitRouterCidrId"] = request.TransitRouterCidrId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterCidrAllocation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterCidrAllocationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the connection information of Express Connect Router (ECR) connections under an Enterprise Edition transit router, including the total number of entries, connection status, connection ID, and payer of the network instance.
//
// Description:
//
// You can query the information about ECR connections under an Enterprise Edition transit router in the following three ways:
//
// - Query the information about all ECR connections under the instance by using the Enterprise Edition transit router instance ID.
//
// - Query the information about all ECR connections under the instance by using the Cloud Enterprise Network (CEN) instance ID and the region ID of the Enterprise Edition transit router.
//
// - Query specific ECR connection information by using only the **TransitRouterAttachmentId*	- parameter.
//
// @param request - ListTransitRouterEcrAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterEcrAttachmentsResponse
func (client *Client) ListTransitRouterEcrAttachmentsWithContext(ctx context.Context, request *ListTransitRouterEcrAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterEcrAttachmentsResponse, _err error) {
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterEcrAttachments"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterEcrAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the associations between a multicast domain and vSwitches.
//
// Description:
//
// - When calling this operation, you must specify at least one of the request parameters **TransitRouterMulticastDomainId*	- and **TransitRouterAttachmentId**. Specifying **TransitRouterAttachmentId*	- queries the information about vSwitches that are associated with a multicast domain under a VPC-connected instance. Specifying **TransitRouterMulticastDomainId*	- queries the information about vSwitches that are associated with the multicast domain.
//
// - When calling the **ListTransitRouterMulticastDomainAssociations*	- operation, make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the operation still returns a **RequestId*	- but does not return the associations between the multicast domain and vSwitches.
//
// @param request - ListTransitRouterMulticastDomainAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainAssociationsResponse
func (client *Client) ListTransitRouterMulticastDomainAssociationsWithContext(ctx context.Context, request *ListTransitRouterMulticastDomainAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainAssociationsResponse, _err error) {
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

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterMulticastDomainAssociations"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterMulticastDomainAssociationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about vSwitches that are associated with multicast domains in a VPC-connected instance after the VPC-connected instance is connected to an Enterprise Edition transit router.
//
// Description:
//
// When you call the ListTransitRouterMulticastDomainVSwitches operation, make sure that the parameter values you specify are correct. If you specify incorrect parameter values, the system still returns a RequestId but does not return information about vSwitches that are associated with multicast domains in the VPC-connected instance.
//
// @param request - ListTransitRouterMulticastDomainVSwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainVSwitchesResponse
func (client *Client) ListTransitRouterMulticastDomainVSwitchesWithContext(ctx context.Context, request *ListTransitRouterMulticastDomainVSwitchesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainVSwitchesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterMulticastDomainVSwitches"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterMulticastDomainVSwitchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about multicast domains, such as the status, multicast domain ID, and description.
//
// Description:
//
// - RegionId must be used in combination with CenId and cannot be used alone. Otherwise, the multicast domain information is not displayed. However, TransitRouterId and TransitRouterMulticastDomainId can be used alone.
//
// - Make sure that the input parameter values are correct when you call this operation. If the input parameters are invalid, the operation still returns **RequestId**, but the detailed information about the multicast domains is not displayed.
//
// @param request - ListTransitRouterMulticastDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainsResponse
func (client *Client) ListTransitRouterMulticastDomainsWithContext(ctx context.Context, request *ListTransitRouterMulticastDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainsResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterMulticastDomains"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterMulticastDomainsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of multicast members and multicast sources in a multicast domain.
//
// Description:
//
// You can call the `ListTransitRouterMulticastGroups` operation to query information about multicast members and multicast sources (hereinafter referred to as multicast resources) in a multicast domain.
//
// - If you specify the **GroupIpAddress*	- parameter, you can query multicast resources in a specified multicast group within the multicast domain.
//
// - If you specify the **VSwitchIds*	- parameter, you can query multicast resources under a specified vSwitch within the multicast domain.
//
// - If you specify the **PeerTransitRouterMulticastDomains*	- parameter, you can query cross-region multicast resources within the multicast domain.
//
// - If you specify the **ResourceType*	- parameter, you can query multicast resources of a specified resource type within the multicast domain.
//
// - If you specify the **ResourceId*	- parameter, you can query multicast resources associated with a specified resource.
//
// - If you specify only the **TransitRouterMulticastDomainId*	- parameter, you can query all multicast resources within the multicast domain.
//
// @param request - ListTransitRouterMulticastGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastGroupsResponse
func (client *Client) ListTransitRouterMulticastGroupsWithContext(ctx context.Context, request *ListTransitRouterMulticastGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastGroupsResponse, _err error) {
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

	if !dara.IsNil(request.GroupIpAddress) {
		query["GroupIpAddress"] = request.GroupIpAddress
	}

	if !dara.IsNil(request.IsGroupMember) {
		query["IsGroupMember"] = request.IsGroupMember
	}

	if !dara.IsNil(request.IsGroupSource) {
		query["IsGroupSource"] = request.IsGroupSource
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetworkInterfaceIds) {
		query["NetworkInterfaceIds"] = request.NetworkInterfaceIds
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

	if !dara.IsNil(request.PeerTransitRouterMulticastDomains) {
		query["PeerTransitRouterMulticastDomains"] = request.PeerTransitRouterMulticastDomains
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

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterMulticastGroups"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterMulticastGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of inter-region connections under an Enterprise Edition transit router.
//
// Description:
//
// You can query the information about inter-region connections under an Enterprise Edition transit router in the following two ways:
//
// - Query the information about all inter-region connections under the Enterprise Edition transit router by using instance ID of the Enterprise Edition transit router.
//
// - Query the information about all inter-region connections under the Enterprise Edition transit router by using instance ID of the Cloud Enterprise Network (CEN) instance and the region ID of the Enterprise Edition transit router.
//
// @param request - ListTransitRouterPeerAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterPeerAttachmentsResponse
func (client *Client) ListTransitRouterPeerAttachmentsWithContext(ctx context.Context, request *ListTransitRouterPeerAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterPeerAttachmentsResponse, _err error) {
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterPeerAttachments"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterPeerAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about prefix lists associated with an Enterprise Edition transit router route table.
//
// @param request - ListTransitRouterPrefixListAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterPrefixListAssociationResponse
func (client *Client) ListTransitRouterPrefixListAssociationWithContext(ctx context.Context, request *ListTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterPrefixListAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextHop) {
		query["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.NextHopInstanceId) {
		query["NextHopInstanceId"] = request.NextHopInstanceId
	}

	if !dara.IsNil(request.NextHopType) {
		query["NextHopType"] = request.NextHopType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.OwnerUid) {
		query["OwnerUid"] = request.OwnerUid
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PrefixListId) {
		query["PrefixListId"] = request.PrefixListId
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterTableId) {
		query["TransitRouterTableId"] = request.TransitRouterTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterPrefixListAssociation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterPrefixListAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ListTransitRouterRouteEntries operation to query the details of route entries in an Enterprise Edition transit router route table.
//
// @param request - ListTransitRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteEntriesResponse
func (client *Client) ListTransitRouterRouteEntriesWithContext(ctx context.Context, request *ListTransitRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteEntriesResponse, _err error) {
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

	if !dara.IsNil(request.PrefixListId) {
		query["PrefixListId"] = request.PrefixListId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteFilter) {
		query["RouteFilter"] = request.RouteFilter
	}

	if !dara.IsNil(request.TransitRouterRouteEntryDestinationCidrBlock) {
		query["TransitRouterRouteEntryDestinationCidrBlock"] = request.TransitRouterRouteEntryDestinationCidrBlock
	}

	if !dara.IsNil(request.TransitRouterRouteEntryIds) {
		query["TransitRouterRouteEntryIds"] = request.TransitRouterRouteEntryIds
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNames) {
		query["TransitRouterRouteEntryNames"] = request.TransitRouterRouteEntryNames
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNextHopId) {
		query["TransitRouterRouteEntryNextHopId"] = request.TransitRouterRouteEntryNextHopId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNextHopResourceId) {
		query["TransitRouterRouteEntryNextHopResourceId"] = request.TransitRouterRouteEntryNextHopResourceId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNextHopResourceType) {
		query["TransitRouterRouteEntryNextHopResourceType"] = request.TransitRouterRouteEntryNextHopResourceType
	}

	if !dara.IsNil(request.TransitRouterRouteEntryNextHopType) {
		query["TransitRouterRouteEntryNextHopType"] = request.TransitRouterRouteEntryNextHopType
	}

	if !dara.IsNil(request.TransitRouterRouteEntryOriginResourceId) {
		query["TransitRouterRouteEntryOriginResourceId"] = request.TransitRouterRouteEntryOriginResourceId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryOriginResourceType) {
		query["TransitRouterRouteEntryOriginResourceType"] = request.TransitRouterRouteEntryOriginResourceType
	}

	if !dara.IsNil(request.TransitRouterRouteEntryStatus) {
		query["TransitRouterRouteEntryStatus"] = request.TransitRouterRouteEntryStatus
	}

	if !dara.IsNil(request.TransitRouterRouteEntryType) {
		query["TransitRouterRouteEntryType"] = request.TransitRouterRouteEntryType
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterRouteEntries"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route association relationships created for an Enterprise Edition transit router route table or a network instance connection.
//
// Description:
//
// When you call the **ListTransitRouterRouteTableAssociations*	- operation, specify at least one of the request parameters **TransitRouterRouteTableId*	- and **TransitRouterAttachmentId**:
//
// - If you specify only the **TransitRouterRouteTableId*	- parameter, the system queries the network instance connections that have route association relationships with the specified Enterprise Edition transit router route table.
//
// - If you specify only the **TransitRouterAttachmentId*	- parameter, the system queries the Enterprise Edition transit router route tables that have route association relationships with the specified network instance connection.
//
// - If you specify both the **TransitRouterRouteTableId*	- and **TransitRouterAttachmentId*	- parameters, the system queries the route association relationship between the specified network instance connection and the specified Enterprise Edition transit router route table.
//
//   - If a route association relationship exists between the network instance connection and the Enterprise Edition transit router route table, the system returns the information about the route association relationship.
//
//   - If no route association relationship exists between the network instance connection and the Enterprise Edition transit router route table, the **TransitRouterAssociations*	- array is empty.
//
// When you call the **ListTransitRouterRouteTableAssociations*	- operation, make sure that the parameter values you specify are correct.
//
// If you specify incorrect parameter values, the system still returns a **RequestId*	- but does not return the route association relationships created for the Enterprise Edition transit router route table or network instance connection.
//
// @param request - ListTransitRouterRouteTableAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTableAssociationsResponse
func (client *Client) ListTransitRouterRouteTableAssociationsWithContext(ctx context.Context, request *ListTransitRouterRouteTableAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTableAssociationsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterAttachmentResourceId) {
		query["TransitRouterAttachmentResourceId"] = request.TransitRouterAttachmentResourceId
	}

	if !dara.IsNil(request.TransitRouterAttachmentResourceType) {
		query["TransitRouterAttachmentResourceType"] = request.TransitRouterAttachmentResourceType
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterRouteTableAssociations"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterRouteTableAssociationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route propagations of a route table in an Enterprise Edition transit router.
//
// @param request - ListTransitRouterRouteTablePropagationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTablePropagationsResponse
func (client *Client) ListTransitRouterRouteTablePropagationsWithContext(ctx context.Context, request *ListTransitRouterRouteTablePropagationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTablePropagationsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterAttachmentResourceId) {
		query["TransitRouterAttachmentResourceId"] = request.TransitRouterAttachmentResourceId
	}

	if !dara.IsNil(request.TransitRouterAttachmentResourceType) {
		query["TransitRouterAttachmentResourceType"] = request.TransitRouterAttachmentResourceType
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterRouteTablePropagations"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterRouteTablePropagationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of route tables of an Enterprise Edition transit router by calling the ListTransitRouterRouteTables operation.
//
// @param request - ListTransitRouterRouteTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTablesResponse
func (client *Client) ListTransitRouterRouteTablesWithContext(ctx context.Context, request *ListTransitRouterRouteTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTablesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteTableOptions) {
		query["RouteTableOptions"] = request.RouteTableOptions
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterRouteTableIds) {
		query["TransitRouterRouteTableIds"] = request.TransitRouterRouteTableIds
	}

	if !dara.IsNil(request.TransitRouterRouteTableNames) {
		query["TransitRouterRouteTableNames"] = request.TransitRouterRouteTableNames
	}

	if !dara.IsNil(request.TransitRouterRouteTableStatus) {
		query["TransitRouterRouteTableStatus"] = request.TransitRouterRouteTableStatus
	}

	if !dara.IsNil(request.TransitRouterRouteTableType) {
		query["TransitRouterRouteTableType"] = request.TransitRouterRouteTableType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterRouteTables"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterRouteTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the connection information of virtual border routers (VBRs) under an Enterprise Edition transit router, including the total number of entries, connection status, connection IDs, and payers of network instances.
//
// Description:
//
// You can query information about virtual border router (VBR) connections on an Enterprise Edition transit router in the following ways:
//
// - Query information about all VBR connections on an Enterprise Edition transit router by specifying the transit router instance ID.
//
// - Query information about all VBR connections on an Enterprise Edition transit router by specifying the Cloud Enterprise Network (CEN) instance ID and the region ID of the transit router instance.
//
// - Query information about a VBR connection by specifying only the TransitRouterAttachmentId parameter.
//
// @param request - ListTransitRouterVbrAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVbrAttachmentsResponse
func (client *Client) ListTransitRouterVbrAttachmentsWithContext(ctx context.Context, request *ListTransitRouterVbrAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVbrAttachmentsResponse, _err error) {
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterVbrAttachments"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterVbrAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about Virtual Private Cloud (VPC) connections under an Enterprise Edition transit router, including the status, billing type, zone information, and associated vSwitch and network interface controller (NIC) details of VPC connections that are active for forwarding and routing traffic.
//
// Description:
//
// You can query information about VPC connections under an Enterprise Edition transit router in the following three ways:
//
// - Query information about all VPC connections under an Enterprise Edition transit router instance by specifying the transit router instance ID.
//
// - Query information about all VPC connections under an Enterprise Edition transit router instance by specifying the Cloud Enterprise Network (CEN) instance ID and the region ID of the Enterprise Edition transit router instance.
//
// - Query information about all VPC connections in a region by specifying the region ID of the Enterprise Edition transit router instance.
//
// @param request - ListTransitRouterVpcAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVpcAttachmentsResponse
func (client *Client) ListTransitRouterVpcAttachmentsWithContext(ctx context.Context, request *ListTransitRouterVpcAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVpcAttachmentsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
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
		Action:      dara.String("ListTransitRouterVpcAttachments"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterVpcAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询转发路由器下vpn连接的详情
//
// @param request - ListTransitRouterVpnAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVpnAttachmentsResponse
func (client *Client) ListTransitRouterVpnAttachmentsWithContext(ctx context.Context, request *ListTransitRouterVpnAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVpnAttachmentsResponse, _err error) {
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouterVpnAttachments"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRouterVpnAttachmentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about transit router instances under a Cloud Enterprise Network (CEN) instance, including the instance type, instance status, transit router instance ID, and whether the multicast feature is enabled.
//
// Description:
//
// When you call this operation to query information about transit router instances under a CEN instance, you can specify the **RegionId*	- and **TransitRouterId*	- parameters as needed. The following describes the relationship between these two parameters:
//
// - If you do not specify **RegionId*	- or **TransitRouterId**, all transit router instances under the CEN instance are queried.
//
// - If you specify only **RegionId**, transit router instances in the specified region under the CEN instance are queried.
//
// - If you specify only **TransitRouterId**, the specified transit router instance under the CEN instance is queried.
//
// @param request - ListTransitRoutersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRoutersResponse
func (client *Client) ListTransitRoutersWithContext(ctx context.Context, request *ListTransitRoutersRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRoutersResponse, _err error) {
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

	if !dara.IsNil(request.FeatureFilter) {
		query["FeatureFilter"] = request.FeatureFilter
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterName) {
		query["TransitRouterName"] = request.TransitRouterName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTransitRouters"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTransitRoutersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// *ModifyCenAttribute*	- is an asynchronous operation. The system first returns a **RequestId**, but the modification of the Cloud Enterprise Network (CEN) instance is not completed. The modification task continues in the background. You can call the **DescribeCens*	- operation to query the status of the CEN instance.
//
// - If the CEN instance is in the **Modifying*	- state, the instance is being modified. In this state, you can only perform query operations on the CEN instance and cannot perform other operations.
//
// - If the CEN instance is in the **Active*	- state, the modification of the CEN instance is completed.
//
// @param request - ModifyCenAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenAttributeResponse
func (client *Client) ModifyCenAttributeWithContext(ctx context.Context, request *ModifyCenAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenAttributeResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProtectionLevel) {
		query["ProtectionLevel"] = request.ProtectionLevel
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
		Action:      dara.String("ModifyCenAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCenAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a bandwidth plan instance by calling the ModifyCenBandwidthPackageAttribute operation.
//
// @param request - ModifyCenBandwidthPackageAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenBandwidthPackageAttributeResponse
func (client *Client) ModifyCenBandwidthPackageAttributeWithContext(ctx context.Context, request *ModifyCenBandwidthPackageAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenBandwidthPackageAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenBandwidthPackageId) {
		query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyCenBandwidthPackageAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCenBandwidthPackageAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the peak bandwidth of a bandwidth plan instance by calling the ModifyCenBandwidthPackageSpec operation.
//
// @param request - ModifyCenBandwidthPackageSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenBandwidthPackageSpecResponse
func (client *Client) ModifyCenBandwidthPackageSpecWithContext(ctx context.Context, request *ModifyCenBandwidthPackageSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenBandwidthPackageSpecResponse, _err error) {
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

	if !dara.IsNil(request.CenBandwidthPackageId) {
		query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyCenBandwidthPackageSpec"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCenBandwidthPackageSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a routing policy by calling the ModifyCenRouteMap operation.
//
// Description:
//
// The `ModifyCenRouteMap` operation is asynchronous. After you send a request, the system returns a **RequestId*	- but the routing policy has not been modified yet. The modification task runs in the background. You can call the `DescribeCenRouteMaps` operation to query the status of the routing policy.
//
// - If the routing policy is in the **Modifying*	- state, the routing policy is being modified. In this state, you can only perform query operations.
//
// - If the routing policy is in the **Active*	- state, the routing policy has been modified.
//
// @param request - ModifyCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenRouteMapResponse
func (client *Client) ModifyCenRouteMapWithContext(ctx context.Context, request *ModifyCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenRouteMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsPathMatchMode) {
		query["AsPathMatchMode"] = request.AsPathMatchMode
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CenRegionId) {
		query["CenRegionId"] = request.CenRegionId
	}

	if !dara.IsNil(request.CidrMatchMode) {
		query["CidrMatchMode"] = request.CidrMatchMode
	}

	if !dara.IsNil(request.CommunityMatchMode) {
		query["CommunityMatchMode"] = request.CommunityMatchMode
	}

	if !dara.IsNil(request.CommunityOperateMode) {
		query["CommunityOperateMode"] = request.CommunityOperateMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationChildInstanceTypes) {
		query["DestinationChildInstanceTypes"] = request.DestinationChildInstanceTypes
	}

	if !dara.IsNil(request.DestinationCidrBlocks) {
		query["DestinationCidrBlocks"] = request.DestinationCidrBlocks
	}

	if !dara.IsNil(request.DestinationInstanceIds) {
		query["DestinationInstanceIds"] = request.DestinationInstanceIds
	}

	if !dara.IsNil(request.DestinationInstanceIdsReverseMatch) {
		query["DestinationInstanceIdsReverseMatch"] = request.DestinationInstanceIdsReverseMatch
	}

	if !dara.IsNil(request.DestinationRegionIds) {
		query["DestinationRegionIds"] = request.DestinationRegionIds
	}

	if !dara.IsNil(request.DestinationRouteTableIds) {
		query["DestinationRouteTableIds"] = request.DestinationRouteTableIds
	}

	if !dara.IsNil(request.MapResult) {
		query["MapResult"] = request.MapResult
	}

	if !dara.IsNil(request.MatchAddressType) {
		query["MatchAddressType"] = request.MatchAddressType
	}

	if !dara.IsNil(request.MatchAsns) {
		query["MatchAsns"] = request.MatchAsns
	}

	if !dara.IsNil(request.MatchCommunitySet) {
		query["MatchCommunitySet"] = request.MatchCommunitySet
	}

	if !dara.IsNil(request.NextPriority) {
		query["NextPriority"] = request.NextPriority
	}

	if !dara.IsNil(request.OperateCommunitySet) {
		query["OperateCommunitySet"] = request.OperateCommunitySet
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Preference) {
		query["Preference"] = request.Preference
	}

	if !dara.IsNil(request.PrependAsPath) {
		query["PrependAsPath"] = request.PrependAsPath
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	if !dara.IsNil(request.RouteTypes) {
		query["RouteTypes"] = request.RouteTypes
	}

	if !dara.IsNil(request.SourceChildInstanceTypes) {
		query["SourceChildInstanceTypes"] = request.SourceChildInstanceTypes
	}

	if !dara.IsNil(request.SourceInstanceIds) {
		query["SourceInstanceIds"] = request.SourceInstanceIds
	}

	if !dara.IsNil(request.SourceInstanceIdsReverseMatch) {
		query["SourceInstanceIdsReverseMatch"] = request.SourceInstanceIdsReverseMatch
	}

	if !dara.IsNil(request.SourceRegionIds) {
		query["SourceRegionIds"] = request.SourceRegionIds
	}

	if !dara.IsNil(request.SourceRouteTableIds) {
		query["SourceRouteTableIds"] = request.SourceRouteTableIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCenRouteMap"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCenRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and capture window duration of a flow log.
//
// Description:
//
// The `ModifyFlowLogAttribute` operation is asynchronous. After you call this operation, the system returns a **RequestId**, but the modification has not been completed. The modification continues in the background. You can call the `DescribeFlowlogs` operation to query the status of the flow log.
//
// - If the flow log is in the **Modifying*	- state, the flow log is being modified. In this state, you can only perform query operations.
//
// - If the flow log is in the **Active*	- state, the flow log has been modified.
//
// @param request - ModifyFlowLogAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttributeWithContext(ctx context.Context, request *ModifyFlowLogAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyFlowLogAttributeResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.FlowLogName) {
		query["FlowLogName"] = request.FlowLogName
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
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
		Action:      dara.String("ModifyFlowLogAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFlowLogAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the payer for a cross-account network instance connected to a transit router instance.
//
// Description:
//
// The ModifyGrantInstanceToTransitRouter operation only supports modifying the payer for a cross-account virtual private cloud (VPC) instance, virtual border router (VBR) instance, or IPsec instance connected to a transit router instance.
//
// @param request - ModifyGrantInstanceToTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGrantInstanceToTransitRouterResponse
func (client *Client) ModifyGrantInstanceToTransitRouterWithContext(ctx context.Context, request *ModifyGrantInstanceToTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *ModifyGrantInstanceToTransitRouterResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGrantInstanceToTransitRouter"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGrantInstanceToTransitRouterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a traffic classification rule.
//
// @param request - ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) ModifyTrafficMatchRuleToTrafficMarkingPolicyWithContext(ctx context.Context, request *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyId) {
		query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	}

	if !dara.IsNil(request.TrafficMatchRuleDescription) {
		query["TrafficMatchRuleDescription"] = request.TrafficMatchRuleDescription
	}

	if !dara.IsNil(request.TrafficMatchRuleId) {
		query["TrafficMatchRuleId"] = request.TrafficMatchRuleId
	}

	if !dara.IsNil(request.TrafficMatchRuleName) {
		query["TrafficMatchRuleName"] = request.TrafficMatchRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTrafficMatchRuleToTrafficMarkingPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an aggregate route.
//
// @param tmpReq - ModifyTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTransitRouteTableAggregationResponse
func (client *Client) ModifyTransitRouteTableAggregationWithContext(ctx context.Context, tmpReq *ModifyTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouteTableAggregationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyTransitRouteTableAggregationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TransitRouteTableAggregationScopeList) {
		request.TransitRouteTableAggregationScopeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransitRouteTableAggregationScopeList, dara.String("TransitRouteTableAggregationScopeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouteTableAggregationCidr) {
		query["TransitRouteTableAggregationCidr"] = request.TransitRouteTableAggregationCidr
	}

	if !dara.IsNil(request.TransitRouteTableAggregationDescription) {
		query["TransitRouteTableAggregationDescription"] = request.TransitRouteTableAggregationDescription
	}

	if !dara.IsNil(request.TransitRouteTableAggregationName) {
		query["TransitRouteTableAggregationName"] = request.TransitRouteTableAggregationName
	}

	if !dara.IsNil(request.TransitRouteTableAggregationScope) {
		query["TransitRouteTableAggregationScope"] = request.TransitRouteTableAggregationScope
	}

	if !dara.IsNil(request.TransitRouteTableAggregationScopeListShrink) {
		query["TransitRouteTableAggregationScopeList"] = request.TransitRouteTableAggregationScopeListShrink
	}

	if !dara.IsNil(request.TransitRouteTableId) {
		query["TransitRouteTableId"] = request.TransitRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTransitRouteTableAggregation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTransitRouteTableAggregationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the capacity reservation properties of an Enterprise Edition transit router.
//
// @param tmpReq - ModifyTransitRouterCapacityReservationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTransitRouterCapacityReservationsResponse
func (client *Client) ModifyTransitRouterCapacityReservationsWithContext(ctx context.Context, tmpReq *ModifyTransitRouterCapacityReservationsRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouterCapacityReservationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyTransitRouterCapacityReservationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DeploymentSets) {
		request.DeploymentSetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentSets, dara.String("DeploymentSets"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MinimumTransitRouterCapacities) {
		request.MinimumTransitRouterCapacitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MinimumTransitRouterCapacities, dara.String("MinimumTransitRouterCapacities"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResetCapacityReservations) {
		request.ResetCapacityReservationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResetCapacityReservations, dara.String("ResetCapacityReservations"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeploymentSetsShrink) {
		query["DeploymentSets"] = request.DeploymentSetsShrink
	}

	if !dara.IsNil(request.MinimumTransitRouterCapacitiesShrink) {
		query["MinimumTransitRouterCapacities"] = request.MinimumTransitRouterCapacitiesShrink
	}

	if !dara.IsNil(request.ResetCapacityReservationsShrink) {
		query["ResetCapacityReservations"] = request.ResetCapacityReservationsShrink
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTransitRouterCapacityReservations"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTransitRouterCapacityReservationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyTransitRouterCidr operation to modify a CIDR block of a transit router.
//
// Description:
//
// - Before you modify a transit router CIDR block, we recommend that you familiarize yourself with the [usage limits of transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// - A transit router CIDR block that has allocated IP addresses cannot be modified.
//
// - If you call the **ModifyTransitRouterCidr*	- operation without modifying the **PublishCidrRoute*	- parameter, this operation is synchronous and the modification takes effect immediately.
//
// - If you call the **ModifyTransitRouterCidr*	- operation and modify the **PublishCidrRoute*	- parameter, this operation is asynchronous. After you send a request, the system returns a **RequestId*	- but the transit router CIDR block is not yet modified. The modification task runs in the background. You can call the **ListTransitRouterCidr*	- operation to query the modification status of the transit router CIDR block.
//
//   - If the transit router CIDR block still shows the information before the modification, the transit router CIDR block is being modified.
//
//   - If the transit router CIDR block shows the updated information, the transit router CIDR block has been modified.
//
// @param request - ModifyTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTransitRouterCidrResponse
func (client *Client) ModifyTransitRouterCidrWithContext(ctx context.Context, request *ModifyTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouterCidrResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PublishCidrRoute) {
		query["PublishCidrRoute"] = request.PublishCidrRoute
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

	if !dara.IsNil(request.TransitRouterCidrId) {
		query["TransitRouterCidrId"] = request.TransitRouterCidrId
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTransitRouterCidr"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTransitRouterCidrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and feature options of a multicast domain.
//
// @param request - ModifyTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTransitRouterMulticastDomainResponse
func (client *Client) ModifyTransitRouterMulticastDomainWithContext(ctx context.Context, request *ModifyTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouterMulticastDomainResponse, _err error) {
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

	if !dara.IsNil(request.Options) {
		query["Options"] = request.Options
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainDescription) {
		query["TransitRouterMulticastDomainDescription"] = request.TransitRouterMulticastDomainDescription
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainName) {
		query["TransitRouterMulticastDomainName"] = request.TransitRouterMulticastDomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTransitRouterMulticastDomain"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTransitRouterMulticastDomainResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource group of a Cloud Enterprise Network (CEN) instance or a bandwidth plan.
//
// Description:
//
// By default, CEN instances and bandwidth plans belong to the default resource group. You can call the `MoveResourceGroup` operation to change the resource group of a CEN instance or a bandwidth plan.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithContext(ctx context.Context, request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the transit router service.
//
// Description:
//
// You can call the `OpenTransitRouterService` operation to activate the transit router service free of charge. After the transit router service is activated, the system automatically generates an order. You can use the order ID returned by the `OpenTransitRouterService` operation to query order information in the <props="china">[Alibaba Cloud Management Console Order Center](https://usercenter2.aliyun.com/order/list?pageIndex=1&pageSize=20)<props="intl">[Alibaba Cloud Management Console Order Center](https://usercenter2-intl.aliyun.com/order/list).
//
// > Before calling this operation, call [CheckTransitRouterService](~~CheckTransitRouterService~~) to check whether the transit router service is already activated for the current account. If it is already activated, you do not need to call this operation again.
//
// @param request - OpenTransitRouterServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenTransitRouterServiceResponse
func (client *Client) OpenTransitRouterServiceWithContext(ctx context.Context, request *OpenTransitRouterServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenTransitRouterServiceResponse, _err error) {
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
		Action:      dara.String("OpenTransitRouterService"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenTransitRouterServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cloud Enterprise Network (CEN) supports the publish route entry feature. You can publish routing entries from a VPC or VBR loaded into CEN to CEN by invoking the PublishRouteEntries operation. If no route conflict exists, other network instances in CEN can learn the published routes.
//
// Description:
//
// The following table lists the default publish status of each type of route entry in CEN. You can call the PublishRouteEntries operation to publish route entries that are not published to CEN.
//
// | Route entry        | Instance to which the route entry belongs         |Published to CEN by default
//
// |------------- |-----------------------|--------------------|
//
// |Route entry that points to an ECS instance      |VPC       |No |
//
// |Route entry that points to a VPN gateway      |VPC       |No |
//
// |Route entry that points to a high availability (HA) virtual IP address    |VPC    |No |
//
// |Route entry that points to a router interface    |VPC    |No |
//
// |Route entry that points to an elastic network interfaces (ENIs)    |VPC    |No |
//
// |Route entry that points to an IPv6 gateway    |VPC    |No |
//
// |Route entry that points to a NAT gateway    |VPC    |No |
//
// |VPC system route entry      | VPC       | Yes |
//
// |Route entry that points to an on-premises data center      |VBR      |Yes |
//
// |BGP route    |VBR    |Yes |
//
// @param request - PublishRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRouteEntriesResponse
func (client *Client) PublishRouteEntriesWithContext(ctx context.Context, request *PublishRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *PublishRouteEntriesResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceRouteTableId) {
		query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
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
		Action:      dara.String("PublishRouteEntries"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Republishes an aggregate route.
//
// Description:
//
// For aggregate routes that failed to be published or were partially published, you can call the **RefreshTransitRouteTableAggregation*	- operation to republish the aggregate route to Virtual Private Cloud (VPC) instances after you resolve the route issue.
//
// If you resolve the problematic route by using one of the following methods, the system automatically republishes the aggregate route and you do not need to manually republish it:
//
// - Delete the association forwarding relationship.
//
// - Disable the route synchronization feature.
//
// - Delete the VPC route table.
//
// - Delete the aggregate route.
//
// You can call the **DescribeTransitRouteTableAggregationDetail*	- operation to query the propagation status of an aggregate route.
//
// @param request - RefreshTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshTransitRouteTableAggregationResponse
func (client *Client) RefreshTransitRouteTableAggregationWithContext(ctx context.Context, request *RefreshTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *RefreshTransitRouteTableAggregationResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouteTableAggregationCidr) {
		query["TransitRouteTableAggregationCidr"] = request.TransitRouteTableAggregationCidr
	}

	if !dara.IsNil(request.TransitRouteTableId) {
		query["TransitRouteTableId"] = request.TransitRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshTransitRouteTableAggregation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshTransitRouteTableAggregationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the RegisterTransitRouterMulticastGroupMembers operation to create or add multicast members.
//
// Description:
//
// Currently, Enterprise Edition transit routers support only elastic network interfaces (ENIs) as multicast members. You can call the `RegisterTransitRouterMulticastGroupMembers` operation to specify ENIs in the same region or cross-region ENIs as multicast members.
//
// - If you specify the **NetworkInterfaceIds*	- parameter, you want to specify ENIs in the current region as multicast members. Make sure that the vSwitch to which the ENI belongs is associated with the multicast domain. For more information, see [AssociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429778.html).
//
// - If you specify the **PeerTransitRouterMulticastDomains*	- parameter, you want to specify multicast members in a multicast group with the same multicast IP address in a different region as multicast members of your current multicast group. Make sure that you have created an inter-region connection. For more information, see [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html).
//
//	For example, you have Multicast Domain 1 in the China (Hangzhou) region with Multicast Group 1, and Multicast Domain 2 in the China (Shanghai) region with Multicast Group 2. Multicast Group 1 and Multicast Group 2 have the same multicast IP address, and Multicast Group 2 in the China (Shanghai) region has Multicast Member 2. When you call the `RegisterTransitRouterMulticastGroupMembers` operation to create multicast members for Multicast Group 1 in the China (Hangzhou) region, if you set **PeerTransitRouterMulticastDomains*	- to the ID of Multicast Domain 2 in the China (Shanghai) region, Multicast Member 2 in Multicast Group 2 in the China (Shanghai) region also becomes a multicast member of Multicast Group 1 in the China (Hangzhou) region.
//
// - The `RegisterTransitRouterMulticastGroupMembers` operation is asynchronous. After you send a request, the system returns a **RequestId*	- but the multicast member is not completely created. The creation task continues to run in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of the multicast member.
//
//   - If the multicast member is in the **Registering*	- state, the multicast member is being created. In this state, you can only query the multicast member but cannot perform other operations.
//
//   - If the multicast member is in the **Registered*	- state, the multicast member is created.
//
// @param request - RegisterTransitRouterMulticastGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterTransitRouterMulticastGroupMembersResponse
func (client *Client) RegisterTransitRouterMulticastGroupMembersWithContext(ctx context.Context, request *RegisterTransitRouterMulticastGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *RegisterTransitRouterMulticastGroupMembersResponse, _err error) {
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

	if !dara.IsNil(request.GroupIpAddress) {
		query["GroupIpAddress"] = request.GroupIpAddress
	}

	if !dara.IsNil(request.NetworkInterfaceIds) {
		query["NetworkInterfaceIds"] = request.NetworkInterfaceIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PeerTransitRouterMulticastDomains) {
		query["PeerTransitRouterMulticastDomains"] = request.PeerTransitRouterMulticastDomains
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterTransitRouterMulticastGroupMembers"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterTransitRouterMulticastGroupMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multicast sources to implement one-to-many multicast communication.
//
// Description:
//
// - Only elastic network interfaces (ENIs) can be specified as multicast sources.
//
// - RegisterTransitRouterMulticastGroupSources is an asynchronous operation. After a request is sent, the system returns a **RequestId*	- but the multicast source is not completely created. The creation task still runs in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of the multicast source.
//
//   - If the multicast source is in the **Registering*	- state, the multicast source is being created. In this state, you can only query the multicast source but cannot perform other operations.
//
//   - If the multicast source is in the **Registered*	- state, the multicast source is created.
//
// ### Before you begin
//
// Before you invoke the `RegisterTransitRouterMulticastGroupSources` operation to create a multicast source, make sure that the vSwitch to which the network interface controller (NIC) of the elastic network interfaces (ENIs) belongs is associated with the multicast domain. For more information, see [AssociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429778.html).
//
// @param request - RegisterTransitRouterMulticastGroupSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterTransitRouterMulticastGroupSourcesResponse
func (client *Client) RegisterTransitRouterMulticastGroupSourcesWithContext(ctx context.Context, request *RegisterTransitRouterMulticastGroupSourcesRequest, runtime *dara.RuntimeOptions) (_result *RegisterTransitRouterMulticastGroupSourcesResponse, _err error) {
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

	if !dara.IsNil(request.GroupIpAddress) {
		query["GroupIpAddress"] = request.GroupIpAddress
	}

	if !dara.IsNil(request.NetworkInterfaceIds) {
		query["NetworkInterfaceIds"] = request.NetworkInterfaceIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterMulticastDomainId) {
		query["TransitRouterMulticastDomainId"] = request.TransitRouterMulticastDomainId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterTransitRouterMulticastGroupSources"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterTransitRouterMulticastGroupSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes traffic classification rules from a traffic marking policy by calling the RemoveTrafficMatchRuleFromTrafficMarkingPolicy operation.
//
// Description:
//
// - When you call the **RemoveTrafficMatchRuleFromTrafficMarkingPolicy*	- operation:
//
//   - If you specify traffic classification rule IDs (the **TrafficMarkRuleIds*	- parameter), the operation deletes the specified traffic classification rules.
//
//   - If you do not specify traffic classification rule IDs (the **TrafficMarkRuleIds*	- parameter), the operation does not perform any action.
//
//     If you want to delete specific traffic classification rules, make sure that you have specified the IDs of the traffic classification rules before you call this operation.
//
// - **RemoveTrafficMatchRuleFromTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the traffic classification rules are not immediately deleted. The deletion task runs in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of traffic classification rules.
//
//   - If a traffic classification rule is in the **Deleting*	- state, the rule is being deleted. In this state, you can only query the traffic classification rule. You cannot perform other operations on it.
//
//   - If the specified traffic classification rule cannot be found, the rule has been deleted.
//
// @param request - RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse
func (client *Client) RemoveTrafficMatchRuleFromTrafficMarkingPolicyWithContext(ctx context.Context, request *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkRuleIds) {
		query["TrafficMarkRuleIds"] = request.TrafficMarkRuleIds
	}

	if !dara.IsNil(request.TrafficMarkingPolicyId) {
		query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTrafficMatchRuleFromTrafficMarkingPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RemoveTraficMatchRuleFromTrafficMarkingPolicy is deprecated, please use Cbn::2017-09-12::RemoveTrafficMatchRuleFromTrafficMarkingPolicy instead.
//
// Summary:
//
// Deletes traffic classification rules from a traffic marking policy.
//
// Description:
//
// ### Precautions
//
// The **RemoveTraficMatchRuleFromTrafficMarkingPolicy*	- operation is deprecated and will be discontinued. To delete traffic classification rules from a traffic marking policy, use the [RemoveTrafficMatchRuleFromTrafficMarkingPolicy](https://help.aliyun.com/document_detail/452726.html) operation. This API documentation is no longer maintained.
//
// @param request - RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse
func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicyWithContext(ctx context.Context, request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkRuleIds) {
		query["TrafficMarkRuleIds"] = request.TrafficMarkRuleIds
	}

	if !dara.IsNil(request.TrafficMarkingPolicyId) {
		query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTraficMatchRuleFromTrafficMarkingPolicy"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces the transit router route table associated with a network instance connection.
//
// Description:
//
// - Only network instance connections under an Enterprise Edition transit router support changing the associated transit router route table.
//
// - **ReplaceTransitRouterRouteTableAssociation*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the transit router route table associated with the network instance connection has not been replaced yet. The replacement task is still running in the background. You can call **ListTransitRouterRouteTableAssociations*	- to query the association forwarding status between the network instance connection and the new transit router route table.
//
//   - If the association forwarding status is **Replacing**, the network instance connection is changing the associated transit router route table. In this state, you can only query the association forwarding relationship between the network instance connection and the transit router route table. You cannot perform other operations.
//
//   - If the association forwarding status is **Active**, the network instance connection has successfully changed the associated transit router route table.
//
// @param request - ReplaceTransitRouterRouteTableAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceTransitRouterRouteTableAssociationResponse
func (client *Client) ReplaceTransitRouterRouteTableAssociationWithContext(ctx context.Context, request *ReplaceTransitRouterRouteTableAssociationRequest, runtime *dara.RuntimeOptions) (_result *ReplaceTransitRouterRouteTableAssociationResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceTransitRouterRouteTableAssociation"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceTransitRouterRouteTableAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a cloud service to add access configurations for on-premises networks by calling the ResolveAndRouteServiceInCen operation.
//
// Description:
//
// Cloud services refer to Alibaba Cloud services that use the 100.64.0.0/10 CIDR block, such as Object Storage Service (OSS), Simple Log Service (SLS), and Data Transmission Service (DTS). If your on-premises network needs to access cloud services, you must load the Virtual Border Router (VBR) instance or Cloud Connect Network (CCN) instance associated with your on-premises network to a Cloud Enterprise Network (CEN) instance. You must also load a VPC-connected instance in the region where the cloud service resides to the same CEN instance. After loading, your on-premises network can access the VPC-connected instance in the region of the cloud service through CEN, and then access the cloud service through the VPC by forwarding traffic. The CEN handles the routing accordingly.
//
// - Limits: This operation applies only to Basic Edition transit routers. On-premises networks associated with VBR instances can access only cloud services in the same region through CEN.
//
//	For example, if the cloud service resides in the China (Beijing) region, only on-premises networks associated with VBR instances in the China (Beijing) region can access the cloud service.
//
// - The **ResolveAndRouteServiceInCen*	- operation is asynchronous. After you send a request, the system returns a **RequestId*	- but the cloud service configuration is not yet complete. The background node for adding the configuration continues to run. You can invoke the **DescribeRouteServicesInCen*	- operation to query the status of the cloud service.
//
//   - If the cloud service is in the **Creating*	- state, the cloud service configuration is being added. In this state, you can only execute a query on the cloud service configuration and cannot execute other operations.
//
//   - If the cloud service is in the **Active*	- state, the cloud service configuration is added.
//
//   - If the cloud service is in the **Failed*	- state, the cloud service configuration failed to be added.
//
// ### Before you begin
//
// Before you invoke the ResolveAndRouteServiceInCen operation, make sure that the following conditions are met:
//
// - The VBR or CCN instance with network connectivity to your on-premises network is loaded to the CEN instance.
//
// - A VPC-connected instance in the region where the cloud service resides is loaded to the CEN instance. For more information, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/65902.html).
//
// @param request - ResolveAndRouteServiceInCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResolveAndRouteServiceInCenResponse
func (client *Client) ResolveAndRouteServiceInCenWithContext(ctx context.Context, request *ResolveAndRouteServiceInCenRequest, runtime *dara.RuntimeOptions) (_result *ResolveAndRouteServiceInCenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessRegionIds) {
		query["AccessRegionIds"] = request.AccessRegionIds
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.HostRegionId) {
		query["HostRegionId"] = request.HostRegionId
	}

	if !dara.IsNil(request.HostVpcId) {
		query["HostVpcId"] = request.HostVpcId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ResolveAndRouteServiceInCen"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResolveAndRouteServiceInCenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions that allow a transit router to connect to a cross-account network instance.
//
// Description:
//
// The `RevokeInstanceFromTransitRouter` operation only supports revoking the permissions that allow a transit router to connect to cross-account Virtual Private Cloud (VPC) instances, Virtual Border Router (VBR) instances, IPsec connections, and Express Connect Router (ECR) instances.
//
// To revoke the permissions that allow a transit router to connect to a cross-account Cloud Connect Network (CCN) instance, call the [RevokeInstanceFromCbn](https://help.aliyun.com/document_detail/126142.html) operation.
//
// ### Before you begin
//
// Before you call the `RevokeInstanceFromTransitRouter` operation, make sure that the connection between the transit router and the VPC-connected instance is deleted.
//
// - To delete the connection between an Enterprise Edition transit router and a VPC instance, see [DeleteTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261220.html).
//
// - To delete the connection between an Enterprise Edition transit router and a VBR instance, see [DeleteTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/261223.html).
//
// - To delete the connection between an Enterprise Edition transit router and an IPsec connection, see [DeleteTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/443992.html).
//
// - To delete the connection between an Enterprise Edition transit router and an ECR instance, see [DeleteTransitRouterEcrAttachment](https://help.aliyun.com/document_detail/443992.html).
//
// - To delete the connection between a Basic Edition transit router and a VPC-connected instance, see [DetachCenChildInstance](https://help.aliyun.com/document_detail/65915.html).
//
// @param request - RevokeInstanceFromTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromTransitRouterResponse
func (client *Client) RevokeInstanceFromTransitRouterWithContext(ctx context.Context, request *RevokeInstanceFromTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *RevokeInstanceFromTransitRouterResponse, _err error) {
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

	if !dara.IsNil(request.CenOwnerId) {
		query["CenOwnerId"] = request.CenOwnerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
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
		Action:      dara.String("RevokeInstanceFromTransitRouter"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeInstanceFromTransitRouterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the PrivateZone service.
//
// Description:
//
// PrivateZone is a private DNS domain name resolution and management service based on the Alibaba Cloud Virtual Private Cloud (VPC) environment. After a Virtual Border Router (VBR) instance or a Cloud Connect Network (CCN) instance is attached to a Cloud Enterprise Network (CEN) instance, the associated on-premises networks can access the PrivateZone service through the CEN instance.
//
// - The on-premises networks associated with VBR and CCN instances can only access the PrivateZone service in the same region.
//
//	For example, if the PrivateZone service is located in the China (Beijing) region, only the on-premises networks associated with VBR instances in the China (Beijing) region and CCN instances in the Chinese mainland can access the PrivateZone service.
//
// - The **RoutePrivateZoneInCenToVpc*	- operation is asynchronous. The system first returns a **RequestId**, but the PrivateZone service configuration is not yet complete, and the background task is still running. You can call the **DescribeCenPrivateZoneRoutes*	- operation to query the status of the PrivateZone service.
//
//   - When the PrivateZone service is in the **Creating*	- state, the PrivateZone service configuration is being added. In this state, you can only query the PrivateZone service configuration and cannot perform other operations.
//
//   - When the PrivateZone service is in the **Active*	- state, the PrivateZone service configuration is added.
//
//   - When the PrivateZone service is in the **Failed*	- state, the PrivateZone service configuration fails to be added.
//
// #### Before you begin
//
// Before you invoke the **RoutePrivateZoneInCenToVpc*	- operation, make sure that the following conditions are met:
//
// - The PrivateZone service deployment is complete. For more information, see [Quick start for PrivateZone DNS parse](https://help.aliyun.com/document_detail/64627.html).
//
// - Make sure that you load the VPC-connected instance associated with the PrivateZone service, and the VBR or CCN instance in the access region to the same CEN instance. For more information, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/65902.html).
//
// - If your on-premises network is connected to Alibaba Cloud by using a CCN instance, and the account to which the CCN instance belongs is different from the account to which the VPC-connected instance or the CEN instance belongs, you must first complete the authorization for the CCN instance. For more information, see [Authorize a Cloud Connect Network](https://help.aliyun.com/document_detail/106674.html).
//
// @param request - RoutePrivateZoneInCenToVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RoutePrivateZoneInCenToVpcResponse
func (client *Client) RoutePrivateZoneInCenToVpcWithContext(ctx context.Context, request *RoutePrivateZoneInCenToVpcRequest, runtime *dara.RuntimeOptions) (_result *RoutePrivateZoneInCenToVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessRegionId) {
		query["AccessRegionId"] = request.AccessRegionId
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.HostRegionId) {
		query["HostRegionId"] = request.HostRegionId
	}

	if !dara.IsNil(request.HostVpcId) {
		query["HostVpcId"] = request.HostVpcId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("RoutePrivateZoneInCenToVpc"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RoutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets, modifies, or deletes the inter-region bandwidth between two regions in a bandwidth package of a Basic Edition transit router.
//
// Description:
//
// This operation supports setting the inter-region bandwidth between two regions only in bandwidth packages of Basic Edition transit routers.
//
// ### Before you begin
//
// A bandwidth package instance is already associated with the target Cloud Enterprise Network (CEN) instance. For more information, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html) and [AssociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65934.html).
//
// You can call **SetCenInterRegionBandwidthLimit*	- to set, modify, or delete the inter-region bandwidth:
//
// - If **BandwidthLimit*	- is not 0, the inter-region bandwidth is set or modified.
//
// - If **BandwidthLimit*	- is 0, the inter-region bandwidth is deleted.
//
// ### Settings
//
// - The maximum inter-region bandwidth cannot exceed the peak bandwidth of the bandwidth package instance to which it belongs.
//
// - The total inter-region bandwidth under a bandwidth package instance cannot exceed the peak bandwidth of that bandwidth package instance.
//
// - If the bandwidth multiplexing feature is enabled for the inter-region connection, modifying the inter-region bandwidth is not supported.
//
// - The **SetCenInterRegionBandwidthLimit*	- operation supports setting, modifying, or deleting inter-region bandwidth only for Basic Edition transit routers.
//
//	To set, modify, or delete inter-region bandwidth for Enterprise Edition transit routers, see [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html), [UpdateTransitRouterPeerAttachmentAttribute](https://help.aliyun.com/document_detail/261229.html), and [DeleteTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261227.html).
//
// @param request - SetCenInterRegionBandwidthLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCenInterRegionBandwidthLimitResponse
func (client *Client) SetCenInterRegionBandwidthLimitWithContext(ctx context.Context, request *SetCenInterRegionBandwidthLimitRequest, runtime *dara.RuntimeOptions) (_result *SetCenInterRegionBandwidthLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BandwidthLimit) {
		query["BandwidthLimit"] = request.BandwidthLimit
	}

	if !dara.IsNil(request.BandwidthType) {
		query["BandwidthType"] = request.BandwidthType
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.LocalRegionId) {
		query["LocalRegionId"] = request.LocalRegionId
	}

	if !dara.IsNil(request.OppositeRegionId) {
		query["OppositeRegionId"] = request.OppositeRegionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("SetCenInterRegionBandwidthLimit"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCenInterRegionBandwidthLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and attaches tags to resources.
//
// Description:
//
// - A tag consists of a tag key and a tag value. Both the tag key and tag value are required when you add a tag.
//
// - If you want to add multiple tags to a Cloud Enterprise Network (CEN) instance, the tag keys of the tags must be unique within the instance.
//
// - You can attach up to 20 tags to a CEN instance.
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
		Version:     dara.String("2017-09-12"),
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
// Calls the TempUpgradeCenBandwidthPackageSpec operation to temporarily upgrade the specifications of a subscription bandwidth plan for Cloud Enterprise Network (CEN).
//
// Description:
//
// Subscription bandwidth plans support the temporary upgrade feature. You can use this feature to increase the bandwidth value of a bandwidth plan within a specified period to flexibly handle business bandwidth fluctuations.
//
// The minimum interval for a temporary upgrade is 3 hours. After the payment is completed, the bandwidth is upgraded immediately without service interruptions.
//
// > After a temporary upgrade expires, the subscription bandwidth plan reverts to the original peak bandwidth. If the service traffic on the instance exceeds the original peak bandwidth limit, the traffic may be dropped due to throttling. Plan the expiration time of the temporary upgrade properly and make sure that the peak bandwidth matches your business requirements.
//
// - The temporary upgrade feature is not available by default. To use this feature, contact your account manager.
//
// - Pay-as-you-go bandwidth plans and expired subscription bandwidth plans do not support the temporary upgrade feature.
//
// - The **TempUpgradeCenBandwidthPackageSpec*	- operation is asynchronous. After you call this operation, the system returns a **RequestId*	- but the bandwidth plan is not yet upgraded. The upgrade task continues to run in the background. You can call the **DescribeCenBandwidthPackages*	- operation to query the specifications of the bandwidth plan. If the specifications meet your expectations, the upgrade is complete.
//
// @param request - TempUpgradeCenBandwidthPackageSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TempUpgradeCenBandwidthPackageSpecResponse
func (client *Client) TempUpgradeCenBandwidthPackageSpecWithContext(ctx context.Context, request *TempUpgradeCenBandwidthPackageSpecRequest, runtime *dara.RuntimeOptions) (_result *TempUpgradeCenBandwidthPackageSpecResponse, _err error) {
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

	if !dara.IsNil(request.CenBandwidthPackageId) {
		query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("TempUpgradeCenBandwidthPackageSpec"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TempUpgradeCenBandwidthPackageSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a bandwidth plan from a Cloud Enterprise Network (CEN) instance. After the disassociation, the bandwidth plan can be associated with other CEN instances.
//
// Description:
//
// Disassociates a CEN instance from a bandwidth plan. Before you call this operation, ensure that no inter-region bandwidth is configured for the target bandwidth plan. You can call [DescribeCenInterRegionBandwidthLimits](https://help.aliyun.com/document_detail/468275.html) to query the inter-region bandwidth, and then call [SetCenInterRegionBandwidthLimit](https://help.aliyun.com/document_detail/65942.html) to set BandwidthLimit to 0 to delete the configured inter-region bandwidth.
//
// @param request - UnassociateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassociateCenBandwidthPackageResponse
func (client *Client) UnassociateCenBandwidthPackageWithContext(ctx context.Context, request *UnassociateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *UnassociateCenBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CenBandwidthPackageId) {
		query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("UnassociateCenBandwidthPackage"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnassociateCenBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a PrivateZone service configuration.
//
// Description:
//
// The **UnroutePrivateZoneInCenToVpc*	- operation is asynchronous. The system first returns a **RequestId**, but the PrivateZone service configuration is not immediately deleted and the deletion task continues to run in the background. You can call the **DescribeCenPrivateZoneRoutes*	- operation to query the status of the PrivateZone service.
//
// - If the PrivateZone service is in the **Deleting*	- state, the PrivateZone service configuration is being deleted. In this state, you can only perform query operations on the PrivateZone service configuration and cannot perform other operations.
//
// - If the specified PrivateZone service configuration cannot be queried, the PrivateZone service configuration is deleted.
//
// If a PrivateZone configuration exists in an access region that is a Cloud Connect Network region, you must first delete the PrivateZone configuration in the Cloud Connect Network region before you delete PrivateZone configurations in other access regions.
//
// @param request - UnroutePrivateZoneInCenToVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnroutePrivateZoneInCenToVpcResponse
func (client *Client) UnroutePrivateZoneInCenToVpcWithContext(ctx context.Context, request *UnroutePrivateZoneInCenToVpcRequest, runtime *dara.RuntimeOptions) (_result *UnroutePrivateZoneInCenToVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessRegionId) {
		query["AccessRegionId"] = request.AccessRegionId
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("UnroutePrivateZoneInCenToVpc"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnroutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Remove tags from resources.
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
		Version:     dara.String("2017-09-12"),
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
// Modifies the name and description of a traffic scheduling policy.
//
// @param request - UpdateCenInterRegionTrafficQosPolicyAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCenInterRegionTrafficQosPolicyAttributeResponse
func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttributeWithContext(ctx context.Context, request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficQosPolicyDescription) {
		query["TrafficQosPolicyDescription"] = request.TrafficQosPolicyDescription
	}

	if !dara.IsNil(request.TrafficQosPolicyId) {
		query["TrafficQosPolicyId"] = request.TrafficQosPolicyId
	}

	if !dara.IsNil(request.TrafficQosPolicyName) {
		query["TrafficQosPolicyName"] = request.TrafficQosPolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCenInterRegionTrafficQosPolicyAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCenInterRegionTrafficQosPolicyAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, cross-region bandwidth, and DSCP value configurations of a queue in a traffic QoS policy.
//
// @param request - UpdateCenInterRegionTrafficQosQueueAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCenInterRegionTrafficQosQueueAttributeResponse
func (client *Client) UpdateCenInterRegionTrafficQosQueueAttributeWithContext(ctx context.Context, request *UpdateCenInterRegionTrafficQosQueueAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCenInterRegionTrafficQosQueueAttributeResponse, _err error) {
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

	if !dara.IsNil(request.Dscps) {
		query["Dscps"] = request.Dscps
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QosQueueDescription) {
		query["QosQueueDescription"] = request.QosQueueDescription
	}

	if !dara.IsNil(request.QosQueueId) {
		query["QosQueueId"] = request.QosQueueId
	}

	if !dara.IsNil(request.QosQueueName) {
		query["QosQueueName"] = request.QosQueueName
	}

	if !dara.IsNil(request.RemainBandwidthPercent) {
		query["RemainBandwidthPercent"] = request.RemainBandwidthPercent
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
		Action:      dara.String("UpdateCenInterRegionTrafficQosQueueAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCenInterRegionTrafficQosQueueAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of an existing route map.
//
// Description:
//
// The UpdateRouteMap operation is asynchronous. After you call this operation, the system returns a **RequestId**, but the route map is not immediately updated. The update task continues to run in the background. You can call the ListRouteMaps operation to query the status of the route map.
//
// - If the route map is in the **Updating*	- state, the route map is being updated. In this state, you can only perform query operations.
//
// - If the route map is in the **Active*	- state, the route map is updated.
//
// @param request - UpdateRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRouteMapResponse
func (client *Client) UpdateRouteMapWithContext(ctx context.Context, request *UpdateRouteMapRequest, runtime *dara.RuntimeOptions) (_result *UpdateRouteMapResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRouteMap"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRouteMapResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of an existing route map rule.
//
// Description:
//
// ## Request description
//
// - `RouteMapRuleId` specifies the route map rule to be updated.
//
// - `Name`, `Description`, and `Priority` are used to update the name, description, and priority of the rule. Rule priorities under the same route map must be unique.
//
// - `MatchConditions` is used to update the match conditions of the rule, and `SetActions` is used to update the actions to be performed when a route meets the match conditions.
//
// - `DryRun` is an optional parameter used to precheck the request without actually updating the route map rule.
//
// @param tmpReq - UpdateRouteMapRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRouteMapRuleResponse
func (client *Client) UpdateRouteMapRuleWithContext(ctx context.Context, tmpReq *UpdateRouteMapRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRouteMapRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRouteMapRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MatchConditions) {
		request.MatchConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchConditions, dara.String("MatchConditions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SetActions) {
		request.SetActionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SetActions, dara.String("SetActions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.MatchConditionsShrink) {
		query["MatchConditions"] = request.MatchConditionsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RouteMapRuleId) {
		query["RouteMapRuleId"] = request.RouteMapRuleId
	}

	if !dara.IsNil(request.SetActionsShrink) {
		query["SetActions"] = request.SetActionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRouteMapRule"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRouteMapRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and traffic classification rules of a traffic marking policy.
//
// @param request - UpdateTrafficMarkingPolicyAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficMarkingPolicyAttributeResponse
func (client *Client) UpdateTrafficMarkingPolicyAttributeWithContext(ctx context.Context, request *UpdateTrafficMarkingPolicyAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTrafficMarkingPolicyAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddTrafficMatchRules) {
		query["AddTrafficMatchRules"] = request.AddTrafficMatchRules
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeleteTrafficMatchRules) {
		query["DeleteTrafficMatchRules"] = request.DeleteTrafficMatchRules
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyDescription) {
		query["TrafficMarkingPolicyDescription"] = request.TrafficMarkingPolicyDescription
	}

	if !dara.IsNil(request.TrafficMarkingPolicyId) {
		query["TrafficMarkingPolicyId"] = request.TrafficMarkingPolicyId
	}

	if !dara.IsNil(request.TrafficMarkingPolicyName) {
		query["TrafficMarkingPolicyName"] = request.TrafficMarkingPolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrafficMarkingPolicyAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrafficMarkingPolicyAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a transit router instance.
//
// Description:
//
// *UpdateTransitRouter*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the modification of the transit router instance is not yet complete. The modification task continues to run in the background. You can call the **ListTransitRouters*	- operation to query the status of the transit router instance.
//
// - If the transit router instance is in the **Modifying*	- state, the transit router instance is being modified. In this state, you can only query the transit router instance. You cannot perform other operations on it.
//
// - If the transit router instance is in the **Active*	- state, the transit router instance has been modified.
//
// @param request - UpdateTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterResponse
func (client *Client) UpdateTransitRouterWithContext(ctx context.Context, request *UpdateTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterResponse, _err error) {
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

	if !dara.IsNil(request.TransitRouterDescription) {
		query["TransitRouterDescription"] = request.TransitRouterDescription
	}

	if !dara.IsNil(request.TransitRouterId) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterName) {
		query["TransitRouterName"] = request.TransitRouterName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouter"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the UpdateTransitRouterEcrAttachmentAttribute operation to modify the name and description of an Express Connect Router (ECR) connection under an Enterprise Edition transit router.
//
// Description:
//
// UpdateTransitRouterEcrAttachmentAttribute is an asynchronous operation. The system returns a RequestId immediately, but the ECR connection has not been modified yet because the modification task is still running in the background. You can call the ListTransitRouterEcrAttachments operation to query the status of the ECR connection.
//
// If the ECR connection is in the Modifying state, the ECR connection is being modified. In this state, you can only query the ECR connection but cannot perform other operations on it.
//
// If the ECR connection is in the Attached state, the ECR connection has been modified.
//
// @param request - UpdateTransitRouterEcrAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterEcrAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterEcrAttachmentAttributeWithContext(ctx context.Context, request *UpdateTransitRouterEcrAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterEcrAttachmentAttributeResponse, _err error) {
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

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouterEcrAttachmentAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterEcrAttachmentAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an inter-region connection on an Enterprise Edition transit router by calling the UpdateTransitRouterPeerAttachmentAttribute operation.
//
// Description:
//
// *UpdateTransitRouterPeerAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the inter-region connection is not yet modified. The modification task continues to run in the background. You can call **ListTransitRouterPeerAttachments*	- to query the status of the inter-region connection.
//
// - If the inter-region connection is in the **Modifying*	- state, the inter-region connection is being modified. In this state, you can only query the inter-region connection but cannot perform other operations on it.
//
// - If the inter-region connection is in the **Attached*	- state, the inter-region connection is modified.
//
// @param request - UpdateTransitRouterPeerAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterPeerAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterPeerAttachmentAttributeWithContext(ctx context.Context, request *UpdateTransitRouterPeerAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterPeerAttachmentAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPublishRouteEnabled) {
		query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BandwidthType) {
		query["BandwidthType"] = request.BandwidthType
	}

	if !dara.IsNil(request.CenBandwidthPackageId) {
		query["CenBandwidthPackageId"] = request.CenBandwidthPackageId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DefaultLinkType) {
		query["DefaultLinkType"] = request.DefaultLinkType
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouterPeerAttachmentAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterPeerAttachmentAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a route entry in an Enterprise Edition transit router route table.
//
// @param request - UpdateTransitRouterRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterRouteEntryResponse
func (client *Client) UpdateTransitRouterRouteEntryWithContext(ctx context.Context, request *UpdateTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterRouteEntryResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryDescription) {
		query["TransitRouterRouteEntryDescription"] = request.TransitRouterRouteEntryDescription
	}

	if !dara.IsNil(request.TransitRouterRouteEntryId) {
		query["TransitRouterRouteEntryId"] = request.TransitRouterRouteEntryId
	}

	if !dara.IsNil(request.TransitRouterRouteEntryName) {
		query["TransitRouterRouteEntryName"] = request.TransitRouterRouteEntryName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouterRouteEntry"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterRouteEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the UpdateTransitRouterRouteTable operation to modify the name and description of an Enterprise Edition transit router route table and to enable or disable multi-region equal-cost multi-path (ECMP) routing.
//
// @param request - UpdateTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterRouteTableResponse
func (client *Client) UpdateTransitRouterRouteTableWithContext(ctx context.Context, request *UpdateTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterRouteTableResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteTableOptions) {
		query["RouteTableOptions"] = request.RouteTableOptions
	}

	if !dara.IsNil(request.TransitRouterRouteTableDescription) {
		query["TransitRouterRouteTableDescription"] = request.TransitRouterRouteTableDescription
	}

	if !dara.IsNil(request.TransitRouterRouteTableId) {
		query["TransitRouterRouteTableId"] = request.TransitRouterRouteTableId
	}

	if !dara.IsNil(request.TransitRouterRouteTableName) {
		query["TransitRouterRouteTableName"] = request.TransitRouterRouteTableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouterRouteTable"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterRouteTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and automatic route forwarding settings of a virtual border router (VBR) connection on an Enterprise Edition transit router.
//
// Description:
//
// *UpdateTransitRouterVbrAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the VBR connection is not yet modified. The modification task continues to run in the background. You can call **ListTransitRouterVbrAttachments*	- to query the status of the VBR connection.
//
// - If the VBR connection is in the **Modifying*	- state, the VBR connection is being modified. In this state, you can only query the VBR connection but cannot perform other operations.
//
// - If the VBR connection is in the **Attached*	- state, the VBR connection is modified.
//
// @param request - UpdateTransitRouterVbrAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVbrAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVbrAttachmentAttributeWithContext(ctx context.Context, request *UpdateTransitRouterVbrAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVbrAttachmentAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPublishRouteEnabled) {
		query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouterVbrAttachmentAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterVbrAttachmentAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the UpdateTransitRouterVpcAttachmentAttribute operation to modify the name and description of a VPC connection under an Enterprise Edition transit router and specifies whether the Enterprise Edition transit router automatically publishes routing to the VPC-connected instance.
//
// Description:
//
// *UpdateTransitRouterVpcAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- but the VPC connection has not been modified. The modification task continues to run in the background. You can call **ListTransitRouterVpcAttachments*	- to query the status of the VPC connection.
//
// - If the VPC connection is in the **Modifying*	- state, the VPC connection is being modified. In this state, you can only query the VPC connection but cannot perform other operations.
//
// - If the VPC connection is in the **Attached*	- state, the VPC connection is modified.
//
// @param tmpReq - UpdateTransitRouterVpcAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVpcAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVpcAttachmentAttributeWithContext(ctx context.Context, tmpReq *UpdateTransitRouterVpcAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpcAttachmentAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTransitRouterVpcAttachmentAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("Options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransitRouterVPCAttachmentOptions) {
		request.TransitRouterVPCAttachmentOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransitRouterVPCAttachmentOptions, dara.String("TransitRouterVPCAttachmentOptions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPublishRouteEnabled) {
		query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OptionsShrink) {
		query["Options"] = request.OptionsShrink
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	if !dara.IsNil(request.TransitRouterVPCAttachmentOptionsShrink) {
		query["TransitRouterVPCAttachmentOptions"] = request.TransitRouterVPCAttachmentOptionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouterVpcAttachmentAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterVpcAttachmentAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the zones and vSwitches associated with a virtual private cloud (VPC) connection by calling the UpdateTransitRouterVpcAttachmentZones operation.
//
// Description:
//
// - When you add zones and vSwitches to a VPC connection, make sure that the vSwitch has an idle IP address. During the modification procedure, the Enterprise Edition transit router creates an elastic network interfaces (ENIs) in the vSwitch (which occupies one IP address of the vSwitch) as the interface for routing traffic between the VPC-connected instance and the Enterprise Edition transit router.
//
// - The **UpdateTransitRouterVpcAttachmentZones*	- operation is asynchronous. After you send a request, the system returns a **RequestId*	- but the VPC connection is not yet modified. The modification task continues to run in the background. You can invoke the **ListTransitRouterVpcAttachments*	- operation to query the status of the VPC connection.
//
//   - If the VPC connection is in the **Modifying*	- state, the VPC connection is being modified. In this state, you can only query the VPC connection but cannot perform other operations.
//
//   - If the VPC connection is in the **Attached*	- state, the VPC connection is modified.
//
// - At least one zone and vSwitch mapping must be retained under a **VPC connection ID**. You cannot delete all zone and vSwitch mappings.
//
// @param request - UpdateTransitRouterVpcAttachmentZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVpcAttachmentZonesResponse
func (client *Client) UpdateTransitRouterVpcAttachmentZonesWithContext(ctx context.Context, request *UpdateTransitRouterVpcAttachmentZonesRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpcAttachmentZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddZoneMappings) {
		query["AddZoneMappings"] = request.AddZoneMappings
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RemoveZoneMappings) {
		query["RemoveZoneMappings"] = request.RemoveZoneMappings
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouterVpcAttachmentZones"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterVpcAttachmentZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the UpdateTransitRouterVpnAttachmentAttribute operation to modify the name, description, and whether to automatically publish route entries for a VPN connection under an Enterprise Edition transit router.
//
// @param request - UpdateTransitRouterVpnAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVpnAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVpnAttachmentAttributeWithContext(ctx context.Context, request *UpdateTransitRouterVpnAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpnAttachmentAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPublishRouteEnabled) {
		query["AutoPublishRouteEnabled"] = request.AutoPublishRouteEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransitRouterAttachmentDescription) {
		query["TransitRouterAttachmentDescription"] = request.TransitRouterAttachmentDescription
	}

	if !dara.IsNil(request.TransitRouterAttachmentId) {
		query["TransitRouterAttachmentId"] = request.TransitRouterAttachmentId
	}

	if !dara.IsNil(request.TransitRouterAttachmentName) {
		query["TransitRouterAttachmentName"] = request.TransitRouterAttachmentName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTransitRouterVpnAttachmentAttribute"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTransitRouterVpnAttachmentAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the WithdrawPublishedRouteEntries operation to withdraw routing entries that have been published from a virtual private cloud (VPC) or virtual border router (VBR) instance to Cloud Enterprise Network (CEN).
//
// @param request - WithdrawPublishedRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawPublishedRouteEntriesResponse
func (client *Client) WithdrawPublishedRouteEntriesWithContext(ctx context.Context, request *WithdrawPublishedRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *WithdrawPublishedRouteEntriesResponse, _err error) {
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

	if !dara.IsNil(request.ChildInstanceId) {
		query["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		query["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceRouteTableId) {
		query["ChildInstanceRouteTableId"] = request.ChildInstanceRouteTableId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		query["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		query["DestinationCidrBlock"] = request.DestinationCidrBlock
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
		Action:      dara.String("WithdrawPublishedRouteEntries"),
		Version:     dara.String("2017-09-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WithdrawPublishedRouteEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
