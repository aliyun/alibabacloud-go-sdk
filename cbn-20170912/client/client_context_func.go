// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Enables a flow log. After the flow log is enabled, the system collects traffic information about a specified resource.
//
// Description:
//
//	  After you create a flow log, it is enabled by default. You can call this operation to enable a disabled flow log.
//
//		- `ActiveFlowLog` is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//	    	- If a flow log is in the **Modifying*	- state, the flow log is being enabled. In this case, you can query the flow log but cannot perform other operations.
//
//	    	- If a flow log is in the **Active*	- state, the flow log is enabled.
//
// @param request - ActiveFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveFlowLogResponse
func (client *Client) ActiveFlowLogWithContext(ctx context.Context, request *ActiveFlowLogRequest, runtime *dara.RuntimeOptions) (_result *ActiveFlowLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// *AddTrafficMatchRuleToTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of a traffic classification rule.
//
//   - If a traffic classification rule is in the **Creating*	- state, the traffic classification rule is being created. In this case, you can query the traffic classification rule but cannot perform other operations.
//
//   - If a traffic classification rule is in the **Active*	- state, the traffic classification rule is added to the traffic marking policy.
//
// @param request - AddTrafficMatchRuleToTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTrafficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) AddTrafficMatchRuleToTrafficMarkingPolicyWithContext(ctx context.Context, request *AddTrafficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddTrafficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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
// ### Usage notes
//
// The **AddTraficMatchRuleToTrafficMarkingPolicy*	- operation is deprecated and will be discontinued soon. If you need to add a traffic classification rule to a traffic marking policy, call the [AddTrafficMatchRuleToTrafficMarkingPolicy](https://help.aliyun.com/document_detail/427602.html) operation.
//
// @param request - AddTraficMatchRuleToTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTraficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) AddTraficMatchRuleToTrafficMarkingPolicyWithContext(ctx context.Context, request *AddTraficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddTraficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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
// You can associate multiple bandwidth plans with a CEN instance. However, the pair of areas connected by each bandwidth plan must be unique.
//
// For example, if a CEN instance is associated with a bandwidth plan that connects networks in the Chinese mainland, you cannot associate another bandwidth plan that also connects networks in the Chinese mainland with the CEN instance. However, you can associate a bandwidth plan that connects the Chinese mainland to North America with the CEN instance.
//
// @param request - AssociateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateCenBandwidthPackageResponse
func (client *Client) AssociateCenBandwidthPackageWithContext(ctx context.Context, request *AssociateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *AssociateCenBandwidthPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates an associated forwarding correlation.
//
// Description:
//
// After you create a network instance connection on a transit router, you can configure an associated forwarding correlation to associate the network instance connection with the route table of an Enterprise Edition transit router. The Enterprise Edition transit router forwards traffic for the network instance based on the routes in the route table. Before you begin, we recommend that you take note of the following rules:
//
//   - Only route tables of Enterprise Edition transit routers support associated forwarding correlations. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
//   - Each network instance connection can have an associated forwarding correlation with only one route table of only one Enterprise Edition transit router.
//
//   - **AssociateTransitRouterAttachmentWithRouteTable*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterRouteTableAssociations*	- operation to query the status of an associated forwarding correlation.
//
//   - If an associated forwarding correlation is in the **Associating*	- state, the associated forwarding correlation is being created. You can query the associated forwarding correlation but cannot perform other operations.
//
//   - If an associated forwarding correlation is in the **Active*	- state, the associated forwarding correlation is created.
//
// @param request - AssociateTransitRouterAttachmentWithRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateTransitRouterAttachmentWithRouteTableResponse
func (client *Client) AssociateTransitRouterAttachmentWithRouteTableWithContext(ctx context.Context, request *AssociateTransitRouterAttachmentWithRouteTableRequest, runtime *dara.RuntimeOptions) (_result *AssociateTransitRouterAttachmentWithRouteTableResponse, _err error) {
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
// Associates the vSwitch of a virtual private cloud (VPC) with a multicast domain.
//
// Description:
//
// - A vSwitch can be associated with only one multicast domain. Make sure that the vSwitch is not associated with other multicast domains. For more information about how to disassociate a vSwitch from a multicast domain, see [DisassociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429774.html).
//
// - AssociateTransitRouterMulticastDomain is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the ListTransitRouterMulticastDomainAssociations operation to query whether a vSwitch is associated with the specified multicast domain.
//
//   - If the status is Associating, it indicates that the vSwitch is being associated with the specified multicast domain. You can query the vSwitch but cannot perform other operations on the vSwitch.
//
//   - If the status is Associated, the vSwitch is associated with the specified multicast domain.
//
// - The VPC of the vSwitch must be associated with an Enterprise Edition transit router. For more information about how to associate a VPC with an Enterprise Edition transit router, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/468237.html).
//
// @param request - AssociateTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateTransitRouterMulticastDomainResponse
func (client *Client) AssociateTransitRouterMulticastDomainWithContext(ctx context.Context, request *AssociateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *AssociateTransitRouterMulticastDomainResponse, _err error) {
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
// CEN allows you to attach a network instance that belongs to another Alibaba Cloud account to your CEN instance. Before you attach the network instance, CEN must acquire permissions to access the network instance that belongs to another Alibaba Cloud account.
//
//   - For more information about how to grant CEN permissions on virtual private clouds (VPCs) that belong to another Alibaba Cloud account, see [GrantInstanceToCen](https://help.aliyun.com/document_detail/126224.html).
//
//   - For more information about how to grant CEN permissions on Cloud Connect Network (CCN) instances that belong to another Alibaba Cloud account, see [GrantInstanceToCbn](https://help.aliyun.com/document_detail/126141.html).
//
//   - By default, you cannot grant permissions on virtual border routers (VBRs) that belong to another Alibaba Cloud account to a CEN instance. If you need to use this feature, contact your account manager.
//
// @param request - AttachCenChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachCenChildInstanceResponse
func (client *Client) AttachCenChildInstanceWithContext(ctx context.Context, request *AttachCenChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachCenChildInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries whether your Alibaba Cloud account has the transit router feature activated.
//
// @param request - CheckTransitRouterServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckTransitRouterServiceResponse
func (client *Client) CheckTransitRouterServiceWithContext(ctx context.Context, request *CheckTransitRouterServiceRequest, runtime *dara.RuntimeOptions) (_result *CheckTransitRouterServiceResponse, _err error) {
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
// Cloud Enterprise Network (CEN) instances are basic network resources that are used to manage interconnected networks. You can use a CEN instance to manage a network that covers one or multiple regions. Before you can connect network instances, you must first call the CreateCen operation to create a CEN instance.
//
// Description:
//
// *CreateCen*	- is an asynchronous operation. After you a request is sent, the system returns a request ID and runs the task in the background. You can call **DescribeCens*	- to query the status of the task.
//
//   - If a CEN instance is in the **Creating*	- state, the CEN instance is being created. You can query the CEN instance but cannot perform other operations.
//
//   - If a CEN instance is in the **Active*	- state, the CEN instance is created.
//
// @param request - CreateCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenResponse
func (client *Client) CreateCenWithContext(ctx context.Context, request *CreateCenRequest, runtime *dara.RuntimeOptions) (_result *CreateCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a bandwidth plan.
//
// Description:
//
//	  You must specify the areas to be connected when you create a bandwidth plan. An area contains one or more Alibaba Cloud regions. When you select areas for a bandwidth plan, make sure that the areas contain the regions that you want to connect. For more information about the supported areas and regions, see [Purchase a bandwidth plan](https://help.aliyun.com/document_detail/181560.html).
//
//		- For more information about the billing rules, see [Billing](https://help.aliyun.com/document_detail/189836.html).
//
//		- **CreateCenBandwidthPackage*	- is an asynchronous operation. After you send a request, the system returns a bandwidth plan instance ID and runs the task in the background. You can call the **DescribeCenBandwidthPackages*	- operation to query the status of a bandwidth plan. If a bandwidth plan is in the **Idle*	- or **InUse*	- state, the bandwidth plan is created.
//
// @param request - CreateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenBandwidthPackageResponse
func (client *Client) CreateCenBandwidthPackageWithContext(ctx context.Context, request *CreateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateCenBandwidthPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Adds a route entry to a network instance and sets the next hop of the destination CIDR block to the transit router in the current region.
//
// Description:
//
//	  You can add routes only to virtual private clouds (VPCs) or virtual border routers (VBRs) that are connected to an Enterprise Edition transit router.
//
//		- By default, the next hop of the routes is the **transit router connection**, which is the connection between the VBR and the Enterprise Edition transit router. You cannot modify the next hop.
//
//		- **CreateCenChildInstanceRouteEntryToAttachment*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **DescribeRouteEntryList*	- operation to query the status of a route.
//
//	    	- If a route is in the **Pending*	- state, the route is being created. You can query the route but cannot perform other operations.
//
//	    	- If a route is in the **Available*	- state, the route is created.
//
// @param request - CreateCenChildInstanceRouteEntryToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenChildInstanceRouteEntryToAttachmentResponse
func (client *Client) CreateCenChildInstanceRouteEntryToAttachmentWithContext(ctx context.Context, request *CreateCenChildInstanceRouteEntryToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Adds routes to a network instance.
//
// Description:
//
// ## Limits
//
//   - By default, the CreateCenChildInstanceRouteEntryToCen operation is unavailable. To call this operation,[submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
//   - You cannot add a route entry to an Enterprise Edition transit router by calling the CreateCenChildInstanceRouteEntryToCen operation.
//
//   - By default, the next hop of the route entry is the regional gateway of the Cloud Enterprise Network (CEN) instance. You cannot modify the next hop.
//
// @param request - CreateCenChildInstanceRouteEntryToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenChildInstanceRouteEntryToCenResponse
func (client *Client) CreateCenChildInstanceRouteEntryToCenWithContext(ctx context.Context, request *CreateCenChildInstanceRouteEntryToCenRequest, runtime *dara.RuntimeOptions) (_result *CreateCenChildInstanceRouteEntryToCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a quality of service (QoS) policy for an inter-region connection on an Enterprise Edition transit router.
//
// Description:
//
//	  Only inter-region connections created on Enterprise Edition transit routers support QoS policies.
//
//		- Traffic scheduling applies only to outbound traffic on Enterprise Edition transit routers.
//
//	    For example, you create an inter-region connection between the China (Hangzhou) and China (Qingdao) regions, and create a QoS policy for the transit router in the China (Hangzhou) region. In this case, the QoS policy can ensure bandwidth for network traffic from the China (Hangzhou) region to the China (Qingdao) region. However, the QoS policy does not apply to network traffic from the China (Qingdao) region to the China (Hangzhou) region.
//
//		- **CreateCenInterRegionTrafficQosPolicy*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of the task.
//
//	    	- If the QoS policy is in the **Creating*	- state, the QoS policy is being created. You can query the QoS policy but cannot perform other operations on the QoS policy.
//
//	    	- If the QoS policy is in the **Active*	- state, the QoS policy is created.
//
// ### [](#)Prerequisites
//
// Before you call the **CreateCenInterRegionTrafficQosPolicy*	- operation, make sure that the following requirements are met:
//
//   - An inter-region connection is created. For more information, see [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html).
//
//   - A traffic marking policy is created. For more information, see [CreateTrafficMarkingPolicy](https://help.aliyun.com/document_detail/419025.html).
//
// @param request - CreateCenInterRegionTrafficQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenInterRegionTrafficQosPolicyResponse
func (client *Client) CreateCenInterRegionTrafficQosPolicyWithContext(ctx context.Context, request *CreateCenInterRegionTrafficQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateCenInterRegionTrafficQosPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BandwidthGuaranteeMode) {
		query["BandwidthGuaranteeMode"] = request.BandwidthGuaranteeMode
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
// Creates queues in a quality of service (QoS) policy to manage network traffic based on finer granularities, improve service performance, and meet service-level agreements (SLAs).
//
// Description:
//
// The **CreateCenInterRegionTrafficQosQueue*	- operation is an asynchronous operation. After you send a request, the system returns a request ID and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of the QoS policy to determine the status of the queue. When you call this operation, you must set the **TrafficQosPolicyId*	- parameter.
//
// - If a QoS policy is in the **Modifying*	- state, the queue is being created. In this case, you can query the QoS policy and queue but cannot perform other operations.
//
// - If a QoS policy is in the **Active*	- state, the queue is created.
//
// @param request - CreateCenInterRegionTrafficQosQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenInterRegionTrafficQosQueueResponse
func (client *Client) CreateCenInterRegionTrafficQosQueueWithContext(ctx context.Context, request *CreateCenInterRegionTrafficQosQueueRequest, runtime *dara.RuntimeOptions) (_result *CreateCenInterRegionTrafficQosQueueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a routing policy. A routing policy filters routing information and facilitates network management.
//
// Description:
//
// Routing policies are sorted by priority. A smaller value indicates a higher priority. Each routing policy is a collection of conditional statements and execution statements. Starting from the routing policy with the highest priority, the system matches routes against the match conditions specified by routing policies. If a route meets all the match conditions of a routing policy, the system permits or denies the route based on the action specified in the routing policy. You can also modify the attributes of permitted routes. By default, the system permits routes that meet none of the match conditions. For more information, see [Routing policy overview](https://help.aliyun.com/document_detail/124157.html).
//
// `CreateCenRouteMap` is an asynchronous operation. After you send a request, the routing policy ID is returned but the operation is still being performed in the system background. You can call `DescribeCenRouteMaps` to query the status of a routing policy.
//
//   - If a routing policy is in the **Creating*	- state, the routing policy is being created. In this case, you can query the routing policy but cannot perform other operations.
//
//   - If a routing policy is in the **Active*	- state, the routing policy is created.
//
// @param request - CreateCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenRouteMapResponse
func (client *Client) CreateCenRouteMapWithContext(ctx context.Context, request *CreateCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *CreateCenRouteMapResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Flow logs can be used to capture traffic information about transit routers and network instance connections, including inter-region connections, virtual private cloud (VPC) connections, VPN connections, Express Connect Router (ECR) connections, and virtual border router (VBR) connections. Before you create a flow log, take note of the following items:
//
//   - Flow logs are supported only by Enterprise Edition transit routers.
//
//   - Flow logs are used to capture information about outbound traffic on transit routers. Information about inbound traffic on transit routers is not captured.
//
//     For example, an Elastic Compute Service (ECS) instance in the US (Silicon Valley) region accesses an ECS instance in the US (Virginia) region through Cloud Enterprise Network (CEN). After you enable the flow log feature for the transit router in the US (Virginia) region, you can check the log entries about packets sent from the ECS instance in the US (Virginia) region to the ECS instance in the US (Silicon Valley) region. However, packets sent from the ECS instance in the US (Silicon Valley) region to the ECS instance in the US (Virginia) region are not recorded. If you want to record the packets sent from the ECS instance in the US (Silicon Valley) region to the ECS instance in the US (Virginia) region, you must also enable the flow log feature on the transit router that is in the US (Silicon Valley) region.
//
//   - If you use a flow log to capture traffic information about VPC connections, the flow log captures information only about traffic on the elastic network interface (ENI) of the transit router. For more information about how to view traffic information about other ENIs in the VPC, see [VPC flow log overview](https://help.aliyun.com/document_detail/127150.html).
//
//   - `CreateFlowLog` is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//   - If the flow log is in the **Creating*	- state, the flow log is being created. In this case, you can query the flow log but cannot perform other operations.
//
//   - If the flow log is in the **Active*	- state, the flow log is created.
//
// ### [](#)Prerequisites
//
// Required resources are created. For more information about how to create resources, see the following topics:
//
//   - [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/468237.html)
//
//   - [CreateTransitRouterEcrAttachment](https://help.aliyun.com/document_detail/2715446.html)
//
//   - [CreateTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/468249.html)
//
//   - [CreateTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/468243.html)
//
//   - [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/468270.html)
//
//   - [CreateTransitRouter](https://help.aliyun.com/document_detail/468222.html)
//
// @param request - CreateFlowlogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowlogResponse
func (client *Client) CreateFlowlogWithContext(ctx context.Context, request *CreateFlowlogRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowlogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a traffic marking policy. A traffic marking policy captures network traffic based on traffic classification rules and marks the traffic with the Differentiated Services Code Point (DSCP) values that you specify.
//
// Description:
//
//	  Only Enterprise Edition transit routers support traffic marking policies.
//
//		- **CreateTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a traffic marking policy ID and runs the task in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of a traffic marking policy.
//
//	    	- If a traffic marking policy is in the **Creating*	- state, the traffic marking policy is being created. You can query the traffic marking policy but cannot perform other operations.
//
//	    	- If a traffic marking policy is in the **Active*	- state, the traffic marking policy is created.
//
// @param request - CreateTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficMarkingPolicyResponse
func (client *Client) CreateTrafficMarkingPolicyWithContext(ctx context.Context, request *CreateTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateTrafficMarkingPolicyResponse, _err error) {
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
// After you add an aggregate route to a route table of an Enterprise Edition transit router, the transit router advertises its routes only to route tables of associated virtual private clouds (VPCs) and have route synchronization enabled.
//
// Perform the following operations before you create an aggregate route. Otherwise, the Enterprise Edition transit router does not advertise routes to VPC route tables:
//
//   - Associated forwarding is enabled between the VPCs and the Enterprise Edition transit router. For more information, see [AssociateTransitRouterAttachmentWithRouteTable](https://help.aliyun.com/document_detail/261242.html).
//
//   - Route synchronization is enabled for the VPCs. For more information, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261358.html).
//
// @param tmpReq - CreateTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouteTableAggregationResponse
func (client *Client) CreateTransitRouteTableAggregationWithContext(ctx context.Context, tmpReq *CreateTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouteTableAggregationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// Creates an Enterprise Edition transit router.
//
// Description:
//
//	  You can call **CreateTransitRouter*	- to create an Enterprise Edition transit router. For more information about the regions that support Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
//		- **CreateTransitRouter*	- is an asynchronous operation. After you send a request, the transit router ID is returned but the operation is still being performed in the system background. You can call [ListTransitRouters](https://help.aliyun.com/document_detail/261219.html) to query the status of an Enterprise Edition transit router.
//
//	    	- If an Enterprise Edition transit router is in the **Creating*	- state, the Enterprise Edition transit router is being created. In this case, you can query the Enterprise Edition transit router but cannot perform other operations.
//
//	    	- If an Enterprise Edition transit router is in the **Active*	- state, the Enterprise Edition transit router is created.
//
// @param tmpReq - CreateTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterResponse
func (client *Client) CreateTransitRouterWithContext(ctx context.Context, tmpReq *CreateTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a custom CIDR block for a transit router. Custom CIDR blocks of a transit router are similar to the CIDR blocks of the loopback interface of a router.
//
// Description:
//
// You can specify a CIDR block for a transit router. The CIDR block works in a similar way as the CIDR block of the loopback interface on a router. IP addresses within the CIDR block can be assigned to IPsec-VPN connections. For more information, see [Transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// The **CreateTransitRouterCidr*	- operation can be used to create a CIDR block only after you create a transit router.
//
// The CIDR block must meet the following requirements:
//
//   - Only Enterprise Edition transit routers support custom CIDR blocks.
//
//   - For more information, see [Limits in transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
//   - Each transit router supports at most five CIDR blocks. The subnet mask of a CIDR block must be 16 bits to 24 bits in length.
//
//   - The following CIDR blocks and their subnets are not supported: 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, and 169.254.0.0/16.
//
//   - The CIDR block cannot overlap with the CIDR blocks of the network instances that communicate with each other by using the CEN instance.
//
//   - On the same CEN instance, each transit router CIDR block must be unique.
//
//   - When you create the first VPN connection after you add a CIDR block for a transit router, three CIDR blocks within the CIDR block are reserved. An IP address is allocated from the remaining CIDR blocks to the IPsec-VPN connection.
//
//     You can call the [ListTransitRouterCidrAllocation](https://help.aliyun.com/document_detail/464173.html) operation to query reserved CIDR blocks and IP addresses allocated to network connections.
//
// @param request - CreateTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterCidrResponse
func (client *Client) CreateTransitRouterCidrWithContext(ctx context.Context, request *CreateTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterCidrResponse, _err error) {
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
// Attaches an Express Connect Router (ECR) connection to the transit router in the same region.
//
// Description:
//
//	  Only Enterprise Edition transit routers support ECR connections.
//
//		- The following methods describe how to attach an ECR connection to an Enterprise Edition transit router:
//
//	    	- If an Enterprise Edition transit router is created in the region, specify the **EcrId**, **RegionId**, and **TransitRouterId*	- parameters.
//
//	    	- If no Enterprise Edition transit router is created in the region, specify the **EcrId**, **CenId**, and **RegionId*	- parameters. An Enterprise Edition transit router is automatically created when you create an ECR connection.
//
//		- CreateTransitRouterEcrAttachment is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **ListTransitRouterEcrAttachments*	- operation to query the status of an ECR connection.
//
//	    	- If the ECR connection is in the **Attaching*	- state, the ECR connection is being created. In this case, you can query the ECR connection but cannot perform other operations on the ECR connection.
//
//	    	- If the ECR connection is in the **Attached*	- state, the ECR connection is created.
//
//		- After you create an ECR connection, the ECR connection is not in route learning or associated forwarding relationships with Enterprise Edition transit routers.
//
//	    After you enable [route learning](https://help.aliyun.com/document_detail/468300.html) between the ECR connection and an Enterprise Edition transit router, the routes of the ECR are automatically advertised to the route tables of the Enterprise Edition transit router.
//
//		- After you create an ECR connection, the routes in the route tables of the Enterprise Edition transit router to which the ECR connection is attached are automatically advertised to the route table of the ECR.
//
// ### [](#)Prerequisite
//
//   - The Alibaba Cloud account of the Enterprise Edition transit router and the Alibaba Cloud account of the ECR belong to the same enterprise.
//
//   - The Enterprise Edition transit router and ECR can belong to the same Alibaba Cloud account or different Alibaba Cloud accounts. If the Enterprise Edition transit router and ECR belong to different Alibaba Cloud accounts, grant the transit router permissions on the ECR before you can attach the ECR to the transit router. For more information, see [Acquire permissions to connect to a network instance that belongs to another account](https://help.aliyun.com/document_detail/181553.html).
//
//   - **Before you call this operation to attach an ECR connection to an Enterprise Edition transit router, you must call the [CreateExpressConnectRouterAssociation](https://help.aliyun.com/document_detail/2712082.html) operation to create an association between the ECR and transit router.**
//
//     **If you call the DeleteTransitRouterEcrAttachment operation to forcefully delete an ECR connection, the association between the ECR connection and Enterprise Edition transit router is deleted.**
//
// @param request - CreateTransitRouterEcrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterEcrAttachmentResponse
func (client *Client) CreateTransitRouterEcrAttachmentWithContext(ctx context.Context, request *CreateTransitRouterEcrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterEcrAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a multicast domain. A multicast domain is a multicast network in a region. Only resources in the same multicast domain can transmit and receive multicast packets.
//
// Description:
//
// Before you call this operation, read the following rules:
//
//   - Make sure that an Enterprise Edition transit router is deployed in the region where you want to create the multicast domain, and the multicast feature is enabled for the Enterprise Edition transit router. For more information, see [CreateTransitRouter](https://help.aliyun.com/document_detail/261169.html).
//
//     If an Enterprise Edition transit router was created before you apply for multicast resources, the transit router does not support multicast. You can delete the transit router and create a new one. For more information about how to delete an Enterprise Edition transit router, see [DeleteTransitRouter](https://help.aliyun.com/document_detail/261218.html).
//
//   - When you call **CreateTransitRouterMulticastDomain**, if you set **CenId*	- and **RegionId**, you do not need to set **TransitRouterId**. If you set **TransitRouterId**, you do not need to set **CenId*	- or **RegionId**.
//
// @param request - CreateTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterMulticastDomainResponse
func (client *Client) CreateTransitRouterMulticastDomainWithContext(ctx context.Context, request *CreateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterMulticastDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// To connect network instances across regions, such as virtual private clouds (VPCs), virtual border routers (VBRs), and IPsec-VPN connections that are connected to transit routers, you must create an inter-region connection between the network instances that you want to connect. You can call the CreateTransitRouterPeerAttachment operation to create an inter-region connection on an Enterprise Edition transit router.
//
// Description:
//
//	  Enterprise Edition transit routers allow you to allocate bandwidth resources to inter-region connections by using the following methods:
//
//	    	- **Allocate bandwidth resources from a bandwidth plan**:
//
//	        You must purchase a bandwidth plan, and then allocate bandwidth resources from the bandwidth plan to inter-region connections. For more information about how to purchase a bandwidth plan, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html).
//
//	    	- **Use pay-by-data-transfer bandwidth resources**:
//
//	        You can set a maximum bandwidth value for an inter-region connection. Then, you are charged for the amount of data transfer over the connection. For more information, see [Inter-region data transfer](https://help.aliyun.com/document_detail/337827.html).
//
//		- **CreateTransitRouterPeerAttachment*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **ListTransitRouterPeerAttachments*	- operation to query the status of an inter-region connection.
//
//	    	- If the inter-region connection is in the **Attaching*	- state, the inter-region connection is being created. In this case, you can query the connection but cannot perform other operations on the connection.
//
//	    	- If the inter-region connection is in the **Attached*	- state, the inter-region connection is created.
//
// @param request - CreateTransitRouterPeerAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterPeerAttachmentResponse
func (client *Client) CreateTransitRouterPeerAttachmentWithContext(ctx context.Context, request *CreateTransitRouterPeerAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterPeerAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Associates the route table of an Enterprise Edition transit router with a prefix list.
//
// Description:
//
// To associate an Enterprise Edition transit router with a route prefix, you must meet the following requirements:
//
//   - You are familiar with the limits and route compatibility notes of prefix lists. For more information, see [Prefix lists](https://help.aliyun.com/document_detail/445605.html).
//
//   - A prefix list is created. For more information, see [CreateVpcPrefixList](https://help.aliyun.com/document_detail/437367.html).
//
//   - If the prefix list and the Enterprise Edition transit router belong to different Alibaba Cloud accounts, the prefix list is shared with the Alibaba Cloud account that owns the Enterprise Edition transit router. For more information, see [Resource sharing](https://help.aliyun.com/document_detail/160622.html) and [API references for resource sharing](https://help.aliyun.com/document_detail/193445.html).
//
// @param request - CreateTransitRouterPrefixListAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterPrefixListAssociationResponse
func (client *Client) CreateTransitRouterPrefixListAssociationWithContext(ctx context.Context, request *CreateTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterPrefixListAssociationResponse, _err error) {
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
// Adds a route to a route table of an Enterprise Edition transit router.
//
// Description:
//
// *CreateTransitRouterRouteEntry*	- is an asynchronous operation. After you send a request, the route ID is returned but the operation is still being performed in the system background. You can call **ListTransitRouterRouteEntries*	- to query the status of a route.
//
//   - If a route is in the **Creating*	- state, the route is being created. In this case, you can query the route but cannot perform other operations.
//
//   - If a route is in the **Active*	- state, the route is created.
//
// @param request - CreateTransitRouterRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterRouteEntryResponse
func (client *Client) CreateTransitRouterRouteEntryWithContext(ctx context.Context, request *CreateTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterRouteEntryResponse, _err error) {
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
// Creates a custom route table for an Enterprise Edition transit router.
//
// Description:
//
//	  Only Enterprise Edition transit routers support custom route tables. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
//		- **CreateTransitRouterRouteTable*	- is an asynchronous operation. After you send a request, the route table ID is returned but the operation is still being performed in the system background. You can call **ListTransitRouterRouteTables*	- to query the status of a route table.
//
//	    	- If a route table is in the **Creating*	- state, the route table is being created. In this case, you can query the route table but cannot perform other operations.
//
//	    	- If a route table is in the **Active*	- state, the route table is created.
//
// @param request - CreateTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterRouteTableResponse
func (client *Client) CreateTransitRouterRouteTableWithContext(ctx context.Context, request *CreateTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterRouteTableResponse, _err error) {
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
// Connects the virtual border routers (VBRs) among which you want to establish network communication to the transit router in the region. Then, the transit router can exchange data between the VBR and CEN instance over private connections.
//
// Description:
//
//	  For more information about the regions and zones that support Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
//		- You can use the following methods to create a VBR connection on an Enterprise Edition transit router:
//
//	    	- If an Enterprise Edition transit router is already created in the region, specify the **VbrId**, **RegionId**, and **TransitRouterId*	- parameters to create a VBR connection.
//
//	    	- If no Enterprise Edition transit router is already created in the region, specify the **VbrId**, **CenId**, and **RegionId*	- parameters to create a VBR connection. When you create a VBR connection, the system automatically creates an Enterprise Edition transit router in the specified region.
//
//		- **CreateTransitRouterVbrAttachment*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call **ListTransitRouterVbrAttachments*	- to query the status of a VBR connection.
//
//	    	- If the VBR connection is in the **Attaching*	- state, the VBR connection is being created. In this case, you can query the VBR connection but cannot perform other operations.
//
//	    	- If the VBR connection is in the **Attached*	- state, the VBR connection is created.
//
//		- The transit router and the VBR must belong to the same Alibaba Cloud account.
//
//		- Transit routers can connect to VBRs that belong to the same or a different Alibaba Cloud account. To connect a transit router to a VBR that belongs to a different Alibaba Cloud account, grant permissions on the VBR to the transit router. For more information, see [Grant a transit router permissions on a network instance that belongs to another Alibaba Cloud account](https://help.aliyun.com/document_detail/181553.html).
//
//		- After you create a VBR connection, it is not in route learning or associated forwarding relationship with transit router route tables by default.
//
// @param request - CreateTransitRouterVbrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVbrAttachmentResponse
func (client *Client) CreateTransitRouterVbrAttachmentWithContext(ctx context.Context, request *CreateTransitRouterVbrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVbrAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Attaches virtual private clouds (VPCs) that you want to connect to a transit router. After you attach the VPCs to the same transit router, the VPCs can communicate with each other.
//
// Description:
//
//	  You can use the following methods to create a VPC connection from an Enterprise Edition transit router:
//
//	    	- If an Enterprise Edition transit router is already created in the region where you want to create a VPC connection, configure the **VpcId**, **ZoneMappings.N.VSwitchId**, **ZoneMappings.N.ZoneId**, **TransitRouterId**, and **RegionId*	- parameters.
//
//	    	- If no Enterprise Edition transit router is created in the region where you want to create a VPC connection, configure the **VpcId**, **ZoneMappings.N.VSwitchId**, **ZoneMappings.N.ZoneId**, **CenId**, and **RegionId*	- parameters. Then, the system automatically creates an Enterprise Edition transit router in the specified region.
//
//		- **CreateTransitRouterVpcAttachment*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the [ListTransitRouterVpcAttachments](https://help.aliyun.com/document_detail/261222.html) operation to query the status of a VPC connection.
//
//	    	- If the VPC connection is in the **Attaching*	- state, the VPC connection is being created. You can query the VPC connection but cannot perform other operations.
//
//	    	- If the VPC connection is in the **Attached*	- state, the VPC connection is created.
//
//		- By default, route learning and associated forwarding are disabled between transit router route tables and VPC connections.
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The VPC in the zones of the Enterprise Edition transit router contains at least one vSwitch. Each vSwitch must have at least one idle IP address. For more information, see [Regions and zones supported by Enterprise Edition transit routers](https://help.aliyun.com/document_detail/181681.html).
//
//   - To connect to a network instance that belongs to another Alibaba Cloud account, you must first acquire the permissions from the account. For more information, see [Acquire permissions to connect to a network instance that belongs to another account](https://help.aliyun.com/document_detail/181553.html).
//
//   - VPC connections incur fees. Make sure that you understand the billing rules of VPC connections before you create a VPC connection. For more information, see [Billing](https://help.aliyun.com/document_detail/189836.html).
//
// @param tmpReq - CreateTransitRouterVpcAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVpcAttachmentResponse
func (client *Client) CreateTransitRouterVpcAttachmentWithContext(ctx context.Context, tmpReq *CreateTransitRouterVpcAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVpcAttachmentResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateTransitRouterVpcAttachmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
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
// Attaches an IPsec-VPN connection to a transit router.
//
// Description:
//
//	  By default, route learning and associated forwarding are disabled between transit router route tables and IPsec-VPN attachments.
//
//		- When you call `CreateTransitRouterVpnAttachment`, if you set **CenId*	- and **RegionId**, you do not need to set **TransitRouterId**. If you set **TransitRouterId*	- and **RegionId**, you do not need to set **CenId**.
//
// ### Prerequisites
//
//   - Before you attach an IPsec-VPN connection to a transit router, make sure that at least one IPsec-VPN connection is created in the region where the transit router is deployed. Make sure the IPsec-VPN connection is not associated with a resource. For more information, see [CreateVpnAttachment](https://help.aliyun.com/document_detail/442455.html).
//
//   - If the IPsec-VPN connection to be attached to the transit router belongs to a different Alibaba Cloud account, make sure that the transit router has obtained the required permissions from the IPsec-VPN connection. For more information, see [GrantInstanceToTransitRouter](https://help.aliyun.com/document_detail/417520.html).
//
// @param request - CreateTransitRouterVpnAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVpnAttachmentResponse
func (client *Client) CreateTransitRouterVpnAttachmentWithContext(ctx context.Context, request *CreateTransitRouterVpnAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVpnAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Disables a flow log. A disabled flow log no longer captures information about network traffic.
//
// Description:
//
// `DeactiveFlowLog` is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//   - If a flow log is in the **Modifying*	- state, the flow log is being disabled. You can query the flow log but cannot perform other operations.
//
//   - If a flow log is in the **Inactive*	- state, the flow log is disabled.
//
// @param request - DeactiveFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactiveFlowLogResponse
func (client *Client) DeactiveFlowLogWithContext(ctx context.Context, request *DeactiveFlowLogRequest, runtime *dara.RuntimeOptions) (_result *DeactiveFlowLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// *DeleteCen*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call **DescribeCens*	- to query the status of a CEN instance.
//
//   - If the CEN instance is in the **Deleting*	- state, the CEN instance is being deleted. In this case, you can query the CEN instance but cannot perform other operations.
//
//   - If the CEN instance cannot be found, the CEN instance is deleted.
//
// ### [](#)Prerequisites
//
// The CEN instance that you want to delete is not associated with a bandwidth plan, and the transit router associated with the CEN instance does not have a network instance connection or a custom route table.
//
//   - For more information about how to detach a network instance, see the following topics:
//
//   - [DeleteTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261220.html)
//
//   - [DeleteTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/261223.html)
//
//   - [DeleteTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/443992.html)
//
//   - [DeleteTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261227.html)
//
//     > For more information about how to detach network instances from a Basic Edition transit router, see [DetachCenChildInstance](https://help.aliyun.com/document_detail/65915.html).
//
//   - For more information about how to delete custom route tables from an Enterprise Edition transit router, see [DeleteTransitRouterRouteTable](https://help.aliyun.com/document_detail/261235.html).
//
//   - For more information about how to disassociate a bandwidth plan from a CEN instance, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// @param request - DeleteCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenResponse
func (client *Client) DeleteCenWithContext(ctx context.Context, request *DeleteCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - DeleteCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenBandwidthPackageResponse
func (client *Client) DeleteCenBandwidthPackageWithContext(ctx context.Context, request *DeleteCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenBandwidthPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a route of a network instance from an Enterprise Edition transit router.
//
// Description:
//
//	  You can delete routes only from virtual private clouds (VPCs) and virtual border routers (VBRs) whose next hop is an **Enterprise Edition transit router connection**, which is the connection to the network instance.
//
//		- **DeleteCenChildInstanceRouteEntryToAttachment*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **DescribeRouteEntryList*	- operation to query the status of a route.
//
//	    	- If a route is in the **Deleting*	- state, the route is being deleted. You can query the route but cannot perform other operations.
//
//	    	- If a route cannot be found, the route is deleted.
//
// @param request - DeleteCenChildInstanceRouteEntryToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenChildInstanceRouteEntryToAttachmentResponse
func (client *Client) DeleteCenChildInstanceRouteEntryToAttachmentWithContext(ctx context.Context, request *DeleteCenChildInstanceRouteEntryToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a route from a network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// ## Limits
//
//   - By default, the DeleteCenChildInstanceRouteEntryToCen operation is unavailable. To call this operation, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
//   - You cannot delete a route entry from an Enterprise Edition transit router by calling the DeleteCenChildInstanceRouteEntryToCen operation.
//
// @param request - DeleteCenChildInstanceRouteEntryToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenChildInstanceRouteEntryToCenResponse
func (client *Client) DeleteCenChildInstanceRouteEntryToCenWithContext(ctx context.Context, request *DeleteCenChildInstanceRouteEntryToCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenChildInstanceRouteEntryToCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a quality of service (QoS) policy.
//
// Description:
//
//	  Before you delete a QoS policy, you must delete all queues in the QoS policy except the default queue. For more information, see [DeleteCenInterRegionTrafficQosQueue](https://help.aliyun.com/document_detail/419062.html).
//
//		- **DeleteCenInterRegionTrafficQosPolicy*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of a QoS policy.
//
//	    	- If a QoS policy is in the **Deleting*	- state, the QoS policy is being deleted. You can query the QoS policy but cannot perform other operations.
//
//	    	- If a QoS policy cannot be found, the QoS policy is deleted.
//
// @param request - DeleteCenInterRegionTrafficQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenInterRegionTrafficQosPolicyResponse
func (client *Client) DeleteCenInterRegionTrafficQosPolicyWithContext(ctx context.Context, request *DeleteCenInterRegionTrafficQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenInterRegionTrafficQosPolicyResponse, _err error) {
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
// Deletes a queue from a quality of service (QoS) policy.
//
// Description:
//
//	  You cannot delete the default queue.
//
//		- **DeleteCenInterRegionTrafficQosQueue*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of a queue. If a queue cannot be found, the queue is deleted.
//
// @param request - DeleteCenInterRegionTrafficQosQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenInterRegionTrafficQosQueueResponse
func (client *Client) DeleteCenInterRegionTrafficQosQueueWithContext(ctx context.Context, request *DeleteCenInterRegionTrafficQosQueueRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenInterRegionTrafficQosQueueResponse, _err error) {
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
// Deletes a routing policy.
//
// Description:
//
// `DeleteCenRouteMap` is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the `DescribeCenRouteMaps` operation to query the status of a routing policy.
//
//   - If a routing policy is in the **Deleting*	- state, the routing policy is being deleted. You can query the routing policy but cannot perform other operations.
//
//   - If a routing policy cannot be found, it is deleted.
//
// @param request - DeleteCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenRouteMapResponse
func (client *Client) DeleteCenRouteMapWithContext(ctx context.Context, request *DeleteCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenRouteMapResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// `DeleteFlowlog` is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//   - If a flow log is in the **Deleting*	- state, the flow log is being deleted. In this case, you can query the flow log but cannot perform other operations.
//
//   - If a flow log cannot be found, the flow log is deleted.
//
// @param request - DeleteFlowlogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlogWithContext(ctx context.Context, request *DeleteFlowlogRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowlogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes the configuration of a cloud service connected to a Basic Edition transit router.
//
// Description:
//
// *DeleteRouteServiceInCen*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. If the request parameters are invalid, the system returns a request ID, but the cloud service configuration is not deleted. You can call **DescribeRouteServicesInCen*	- to query the status of the task.
//
//   - If a cloud service is in the **Deleting*	- state, the cloud service configuration is being deleted. In this case, you can only query the cloud service configuration and cannot perform other operations.
//
//   - If the specified cloud service configuration cannot be found, the cloud service configuration is deleted.
//
// @param request - DeleteRouteServiceInCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteServiceInCenResponse
func (client *Client) DeleteRouteServiceInCenWithContext(ctx context.Context, request *DeleteRouteServiceInCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteServiceInCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a specified traffic marking policy.
//
// Description:
//
//	  **DeleteTrafficMarkingPolicy*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of a traffic marking policy.
//
//	    	- If a traffic marking policy is in the **Deleting*	- state, the traffic marking policy is being deleted. You can query the traffic marking policy but cannot perform other operations.
//
//	    	- If a traffic marking policy cannot be found, the traffic marking policy is deleted.
//
//		- Before you delete a traffic marking policy, you must delete all traffic classification rules from the policy. For more information, see [RemoveTrafficMatchRuleFromTrafficMarkingPolicy](https://help.aliyun.com/document_detail/468330.html).
//
// @param request - DeleteTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficMarkingPolicyResponse
func (client *Client) DeleteTrafficMarkingPolicyWithContext(ctx context.Context, request *DeleteTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrafficMarkingPolicyResponse, _err error) {
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
//	  Before you delete an aggregate route, make sure that your network has a redundant route to prevent service interruptions.
//
//		- After an aggregate route is deleted, the aggregate route is automatically withdrawn from virtual private clouds (VPCs). Specific routes that fall within the aggregate route are advertised to the VPCs.
//
// @param request - DeleteTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouteTableAggregationResponse
func (client *Client) DeleteTransitRouteTableAggregationWithContext(ctx context.Context, request *DeleteTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouteTableAggregationResponse, _err error) {
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
// Deletes a transit router.
//
// Description:
//
// *DeleteTransitRouter*	- is an asynchronous operation. After you send a request, the **request ID*	- is returned but the operation is still being performed in the system background. You can call **ListTransitRouters*	- to query the status of a transit router.
//
//   - If a transit router is in the **Deleting*	- state, the transit router is being deleted. In this case, you can query the transit router but cannot perform other operations.
//
//   - If a transit router cannot be found, the transit router is deleted.
//
// #### Prerequisites
//
// Before you delete a transit router, make sure that the following prerequisites are met:
//
// - No network instance connections are created on the transit router.
//
//   - For more information about how to delete a virtual private cloud (VPC) connection, see [DeleteTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261220.html).
//
//   - For more information about how to delete a virtual border router (VBR) connection, see [DeleteTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/261223.html).
//
//   - For more information about how to delete a Cloud Connect Network (CCN) connection, see [DetachCenChildInstance](https://help.aliyun.com/document_detail/65915.html).
//
//   - For more information about how to delete a VPN connection, see [DeleteTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/443992.html).
//
//   - For more information about how to delete an inter-region connection, see [DeleteTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261227.html).
//
// - No custom route tables are created on the transit router. For more information about how to delete a custom route table, see [DeleteTransitRouterRouteTable](https://help.aliyun.com/document_detail/261235.html).
//
// @param request - DeleteTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterResponse
func (client *Client) DeleteTransitRouterWithContext(ctx context.Context, request *DeleteTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterResponse, _err error) {
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
// If IP addresses within the CIDR block have been allocated to network instances, the CIDR block cannot be deleted.
//
// @param request - DeleteTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterCidrResponse
func (client *Client) DeleteTransitRouterCidrWithContext(ctx context.Context, request *DeleteTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterCidrResponse, _err error) {
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
// DeleteTransitRouterEcrAttachment is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the ListTransitRouterEcrAttachments operation to query the status of an ECR connection.
//
// If the ECR connection is in the Detaching state, the ECR connection is being deleted. In this case, you can query the ECR connection but cannot perform other operations on the ECR connection. If the ECR connection cannot be found, the ECR connection is deleted. Before you call the DeleteTransitRouterEcrAttachment operation, make sure that all request parameters are valid. If a request is invalid, a request ID is returned but the ECR connection is not deleted.
//
// @param request - DeleteTransitRouterEcrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterEcrAttachmentResponse
func (client *Client) DeleteTransitRouterEcrAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterEcrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterEcrAttachmentResponse, _err error) {
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
// Deletes a multicast domain.
//
// Description:
//
// Before you delete a multicast domain, make sure that the following requirements are met:
//
//   - The multicast domain is disassociated from all vSwitches. For more information, see [DisassociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429774.html).
//
//   - All multicast sources and members are removed from the multicast domain. For more information, see [DeregisterTransitRouterMulticastGroupSources](https://help.aliyun.com/document_detail/429776.html) and [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/429779.html).
//
//   - The multicast domain is not added to other multicast domains as a multicast member. If the multicast domain is added to another multicast domain as a multicast member, you must remove the multicast domain from the other multicast domain. For more information, see [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/429779.html).
//
//   - Make sure all the request parameters are valid. If a request parameter is invalid, a request ID is returned after you call the operation, but the multicast domain is not deleted.
//
// @param request - DeleteTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterMulticastDomainResponse
func (client *Client) DeleteTransitRouterMulticastDomainWithContext(ctx context.Context, request *DeleteTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterMulticastDomainResponse, _err error) {
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
// Deletes an inter-region connection from an Enterprise Edition transit router.
//
// Description:
//
// *DeleteTransitRouterPeerAttachment*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call **ListTransitRouterPeerAttachments*	- to query the status of an inter-region connection.
//
//   - If an inter-region connection is in the **Detaching*	- state, the inter-region connection is being deleted. You can query the inter-region connection but cannot perform other operations.
//
//   - If an inter-region connection cannot be found, the inter-region connection is deleted.
//
// ## Prerequisites
//
// Before you begin, make sure that the Enterprise Edition transit router that you use to create inter-region connections meets the following prerequisites:
//
//   - No associated forwarding correlation is established between the inter-region connection and the route tables of the Enterprise Edition transit router. For more information about how to delete an associated forwarding correlation, see [DissociateTransitRouterAttachmentFromRouteTable](https://help.aliyun.com/document_detail/260944.html).
//
//   - No route learning correlation is established between the inter-region connection and the route tables of the Enterprise Edition transit router. For more information about how to delete a route learning correlation, see [DisableTransitRouterRouteTablePropagation](https://help.aliyun.com/document_detail/260945.html).
//
//   - The route tables of the Enterprise Edition transit router do not contain a custom route entry whose next hop is the network instance connection. For more information about how to delete custom routes from route tables of Enterprise Edition transit routers, see [DeleteTransitRouterRouteEntry](https://help.aliyun.com/document_detail/261240.html).
//
//   - The route table does not contain a route whose next hop is the inter-region connection and that is generated from a prefix list. You can delete routes from a route table by disassociating the route table from the prefix list. For more information, see [DeleteTransitRouterPrefixListAssociation](https://help.aliyun.com/document_detail/445486.html).
//
//   - No quality of service (QoS) policy is configured for the inter-region connection. For more information about how to delete QoS policies, see [DeleteCenInterRegionTrafficQosPolicy](https://help.aliyun.com/document_detail/427547.html).
//
// @param request - DeleteTransitRouterPeerAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterPeerAttachmentResponse
func (client *Client) DeleteTransitRouterPeerAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterPeerAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterPeerAttachmentResponse, _err error) {
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
// Disassociates the route table of an Enterprise Edition transit router from a prefix list.
//
// Description:
//
// After you disassociate a route table of an Enterprise Edition transit router from a prefix list, the routes that point to the CIDR blocks in the prefix list are automatically withdrawn from the route table. Before you disassociate the route table of an Enterprise Edition transit router from a prefix list, you must migrate workloads that use the routes in case services are interrupted.
//
// @param request - DeleteTransitRouterPrefixListAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterPrefixListAssociationResponse
func (client *Client) DeleteTransitRouterPrefixListAssociationWithContext(ctx context.Context, request *DeleteTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterPrefixListAssociationResponse, _err error) {
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
// Deletes blackhole routes and static routes that point to network instance connections from the route tables of an Enterprise Edition transit router.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you specify the **TransitRouterRouteEntryId*	- parameter to delete a specific route, you do not need to configure the **TransitRouterRouteTableId*	- or **TransitRouterRouteEntryDestinationCidrBlock**. Otherwise, parameter conflicts occur.
//
//   - If you do not specify the **TransitRouterRouteEntryId*	- parameter, configure the following parameters based on the next hop type of the route that you want to delete:
//
//   - To delete a blackhole route, configure the following parameters: **TransitRouterRouteTableId**, **TransitRouterRouteEntryDestinationCidrBlock**, and **TransitRouterRouteEntryNextHopType**.
//
//   - To delete routes other than blackhole routes, configure the following parameters: **TransitRouterRouteTableId**, **TransitRouterRouteEntryDestinationCidrBlock**, **TransitRouterRouteEntryNextHopType**, and **TransitRouterRouteEntryNextHopId**.
//
//   - **DeleteTransitRouterRouteEntry*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterRouteEntries*	- operation to query the status of a route entry.
//
//   - If the route entry is in the **Deleting*	- state, the route entry is being deleted. In this case, you can query the route entry but cannot perform other operations.
//
//   - If a route entry cannot be found, it is deleted.
//
// ### [](#)Limits
//
// You can call this operation to delete only static routes. Automatically learned routes are not supported. You can call the [ListTransitRouterRouteEntries](https://help.aliyun.com/document_detail/260941.html) operation to query route types.
//
// @param request - DeleteTransitRouterRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterRouteEntryResponse
func (client *Client) DeleteTransitRouterRouteEntryWithContext(ctx context.Context, request *DeleteTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterRouteEntryResponse, _err error) {
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
// Deletes a custom route table from an Enterprise Edition transit router.
//
// Description:
//
//	  You cannot delete the default route table of an Enterprise Edition transit router.
//
//		- **DeleteTransitRouterRouteTable*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterRouteTables*	- operation to query the status of a custom route table.
//
//	    	- If a custom route table is in the Deleting state, the custom route table is being deleted. In this case, you can query the custom route table but cannot perform other operations.
//
//	    	- If a custom route table cannot be found, the custom route table is deleted.
//
// @param request - DeleteTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterRouteTableResponse
func (client *Client) DeleteTransitRouterRouteTableWithContext(ctx context.Context, request *DeleteTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterRouteTableResponse, _err error) {
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
// *DeleteTransitRouterVbrAttachment*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterVbrAttachments*	- operation to query the status of a VBR connection.
//
//   - If a VBR connection is in the **Detaching*	- state, the VBR connection is being deleted. You can query the VBR connection but cannot perform other operations.
//
//   - If a VBR connection cannot be found, the VBR connection is deleted.
//
// ## Prerequisites
//
// Before you delete a VBR connection for an Enterprise Edition transit router, make sure that the following requirements are met:
//
//   - No associated forwarding correlation is established between the VBR connection and the route tables of the Enterprise Edition transit router. For more information about how to delete an associated forwarding correlation, see [DissociateTransitRouterAttachmentFromRouteTable](https://help.aliyun.com/document_detail/260944.html).
//
//   - No route learning correlation is established between the VBR connection and the route tables of the Enterprise Edition transit router. For more information about how to delete a route learning correlation, see [DisableTransitRouterRouteTablePropagation](https://help.aliyun.com/document_detail/260945.html).
//
//   - The route tables of the Enterprise Edition transit router do not contain a custom route entry whose next hop is the network instance connection. For more information about how to delete custom route entries, see [DeleteTransitRouterRouteEntry](https://help.aliyun.com/document_detail/261240.html).
//
//   - The route tables of the Enterprise Edition transit router do not contain a route whose next hop is the VBR connection and that is generated from a prefix list. You can delete such routes by disassociating the route table from the prefix list. For more information, see [DeleteTransitRouterPrefixListAssociation](https://help.aliyun.com/document_detail/445486.html).
//
// @param request - DeleteTransitRouterVbrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVbrAttachmentResponse
func (client *Client) DeleteTransitRouterVbrAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterVbrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVbrAttachmentResponse, _err error) {
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
// Deletes a virtual private cloud (VPC) connection from an Enterprise Edition transit router.
//
// Description:
//
// *DeleteTransitRouterVpcAttachment*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterVpcAttachments*	- operation to query the status of a VPC connection.
//
//   - If a VPC connection is in the **Detaching*	- state, the VPC connection is being deleted. You can query the VPC connection but cannot perform other operations.
//
//   - If a VPC connection cannot be found, it is deleted.
//
// ## Prerequisites
//
// Before you delete a VPC connection, make sure that the following requirements are met:
//
//   - No associated forwarding correlation is established between the VPC connection and the route tables of the Enterprise Edition transit router. For more information about how to delete an associated forwarding correlation, see [DissociateTransitRouterAttachmentFromRouteTable](https://help.aliyun.com/document_detail/260944.html).
//
//   - No route learning correlation is established between the VPC connection and the route tables of the Enterprise Edition transit router. For more information about how to delete a route learning correlation, see [DisableTransitRouterRouteTablePropagation](https://help.aliyun.com/document_detail/260945.html).
//
//   - The route table of the VPC does not contain routes that point to the VPC connection. For more information about how to delete routes from a VPC route table, see [DeleteRouteEntry](https://help.aliyun.com/document_detail/36013.html).
//
//   - The route tables of the Enterprise Edition transit router do not contain a custom route entry whose next hop is the network instance connection. For more information about how to delete custom routes from the route tables of an Enterprise Edition transit router, see [DeleteTransitRouterRouteEntry](https://help.aliyun.com/document_detail/261240.html).
//
//   - The route tables of the Enterprise Edition transit router do not contain a route that is generated from a prefix list and the next hop is the VPC connection. You can delete such routes by disassociating the route table from the prefix list. For more information, see [DeleteTransitRouterPrefixListAssociation](https://help.aliyun.com/document_detail/445486.html).
//
// @param request - DeleteTransitRouterVpcAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVpcAttachmentResponse
func (client *Client) DeleteTransitRouterVpcAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterVpcAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVpcAttachmentResponse, _err error) {
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
// Deletes a VPN attachment.
//
// Description:
//
// Before you call the **DeleteTransitRouterVpnAttachment*	- operation, make sure that all request parameters are valid. If a request parameter is invalid, a **request ID*	- is returned, but the VPN attachment is not deleted.
//
// @param request - DeleteTransitRouterVpnAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVpnAttachmentResponse
func (client *Client) DeleteTransitRouterVpnAttachmentWithContext(ctx context.Context, request *DeleteTransitRouterVpnAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVpnAttachmentResponse, _err error) {
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
// Removes a multicast member from a multicast group.
//
// Description:
//
// `DeregisterTransitRouterMulticastGroupMembers` is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of a multicast member.
//
//   - If the multicast member is in the **Deregistering*	- state, the multicast member is being removed. In this case, you can query the multicast member but cannot perform other operations on the multicast member.
//
//   - If a multicast member cannot be found, the multicast member is removed from the multicast group.“
//
// Before you call the DeregisterTransitRouterMulticastGroupMembers operation, make sure that all request parameters are valid. If a request parameter is invalid, a request ID is returned but the multicast member is not removed.
//
// @param request - DeregisterTransitRouterMulticastGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterTransitRouterMulticastGroupMembersResponse
func (client *Client) DeregisterTransitRouterMulticastGroupMembersWithContext(ctx context.Context, request *DeregisterTransitRouterMulticastGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *DeregisterTransitRouterMulticastGroupMembersResponse, _err error) {
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
// Deletes a multicast source from a multicast group.
//
// Description:
//
// `DeregisterTransitRouterMulticastGroupSources` is an asynchronous operation. After a request a sent, the system returns a **request ID*	- and runs the task in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of a multicast source.
//
//   - If a multicast source is in the **Deregistering*	- state, the multicast source is being deleted. You can query the multicast source but cannot perform other operations.
//
//   - If a multicast source cannot be found, the multicast source is deleted.
//
// Before you call DeregisterTransitRouterMulticastGroupSources, make sure that all the request parameters are valid. If a request parameter is invalid, a request ID is returned but the multicast source is not deleted.
//
// @param request - DeregisterTransitRouterMulticastGroupSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterTransitRouterMulticastGroupSourcesResponse
func (client *Client) DeregisterTransitRouterMulticastGroupSourcesWithContext(ctx context.Context, request *DeregisterTransitRouterMulticastGroupSourcesRequest, runtime *dara.RuntimeOptions) (_result *DeregisterTransitRouterMulticastGroupSourcesResponse, _err error) {
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
// Queries the information about a network instance, such as a virtual private cloud (VPC), a virtual border router, or a Cloud Connect Network (CCN) instance, that is attached to a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenAttachedChildInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenAttachedChildInstanceAttributeResponse
func (client *Client) DescribeCenAttachedChildInstanceAttributeWithContext(ctx context.Context, request *DescribeCenAttachedChildInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenAttachedChildInstanceAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the network instances that are attached to a CEN instance.
//
// Description:
//
// You can use one of the following methods to query the information about the network instances that are attached to a CEN instance:
//
//   - You can query all the network instances that are attached to a CEN instance by setting the `CenId` parameter.
//
//   - You can query the network instances that are attached to a CEN instance in a specified region by setting the `CenId` and `ChildInstanceRegionId` parameters.
//
//   - You can query a specified type of network instances that are attached to a CEN instance by setting the `CenId` and `ChildInstanceType` parameters.
//
// @param request - DescribeCenAttachedChildInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenAttachedChildInstancesResponse
func (client *Client) DescribeCenAttachedChildInstancesWithContext(ctx context.Context, request *DescribeCenAttachedChildInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenAttachedChildInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about bandwidth plans.
//
// @param request - DescribeCenBandwidthPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenBandwidthPackagesResponse
func (client *Client) DescribeCenBandwidthPackagesWithContext(ctx context.Context, request *DescribeCenBandwidthPackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenBandwidthPackagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the routes of a network instance that is attached to a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenChildInstanceRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenChildInstanceRouteEntriesResponse
func (client *Client) DescribeCenChildInstanceRouteEntriesWithContext(ctx context.Context, request *DescribeCenChildInstanceRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenChildInstanceRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the remaining bandwidth of a bandwidth plan.
//
// @param request - DescribeCenGeographicSpanRemainingBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenGeographicSpanRemainingBandwidthResponse
func (client *Client) DescribeCenGeographicSpanRemainingBandwidthWithContext(ctx context.Context, request *DescribeCenGeographicSpanRemainingBandwidthRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenGeographicSpanRemainingBandwidthResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the areas that a Cloud Enterprise Network (CEN) instance can connect.
//
// @param request - DescribeCenGeographicSpansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenGeographicSpansResponse
func (client *Client) DescribeCenGeographicSpansWithContext(ctx context.Context, request *DescribeCenGeographicSpansRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenGeographicSpansResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the bandwidth of connections between regions.
//
// @param request - DescribeCenInterRegionBandwidthLimitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenInterRegionBandwidthLimitsResponse
func (client *Client) DescribeCenInterRegionBandwidthLimitsWithContext(ctx context.Context, request *DescribeCenInterRegionBandwidthLimitsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenInterRegionBandwidthLimitsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the connections to PrivateZone.
//
// @param request - DescribeCenPrivateZoneRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenPrivateZoneRoutesResponse
func (client *Client) DescribeCenPrivateZoneRoutesWithContext(ctx context.Context, request *DescribeCenPrivateZoneRoutesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenPrivateZoneRoutesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details about routes in a region for a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenRegionDomainRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenRegionDomainRouteEntriesResponse
func (client *Client) DescribeCenRegionDomainRouteEntriesWithContext(ctx context.Context, request *DescribeCenRegionDomainRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenRegionDomainRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries routing policies.
//
// @param request - DescribeCenRouteMapsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenRouteMapsResponse
func (client *Client) DescribeCenRouteMapsWithContext(ctx context.Context, request *DescribeCenRouteMapsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenRouteMapsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the health check configurations of virtual border routers (VBRs) in a region.
//
// @param request - DescribeCenVbrHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenVbrHealthCheckResponse
func (client *Client) DescribeCenVbrHealthCheckWithContext(ctx context.Context, request *DescribeCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenVbrHealthCheckResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about your Cloud Enterprise Network (CEN) instances.
//
// @param request - DescribeCensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCensResponse
func (client *Client) DescribeCensWithContext(ctx context.Context, request *DescribeCensRequest, runtime *dara.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the regions in which you can attach network instances to Cloud Enterprise Network (CEN) instances.
//
// Description:
//
// The regions that support CEN vary based on the network instance type. To query the regions where you can attach a specified type of network instance to CEN, set the `ProductType` parameter. If you do not set the `ProductType` parameter, the system queries all regions in which you can attach network instances to CEN, regardless of the network instance type.
//
// @param request - DescribeChildInstanceRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChildInstanceRegionsResponse
func (client *Client) DescribeChildInstanceRegionsWithContext(ctx context.Context, request *DescribeChildInstanceRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChildInstanceRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries regions in an area.
//
// @param request - DescribeGeographicRegionMembershipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGeographicRegionMembershipResponse
func (client *Client) DescribeGeographicRegionMembershipWithContext(ctx context.Context, request *DescribeGeographicRegionMembershipRequest, runtime *dara.RuntimeOptions) (_result *DescribeGeographicRegionMembershipResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the network instances of other Alibaba Cloud accounts that have granted permissions to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// Before you call the **DescribeGrantRulesToCen*	- operation, make sure that all request parameters are valid. If a request parameter is invalid, a **request ID*	- is returned, but the network instances are not returned.
//
// @param request - DescribeGrantRulesToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantRulesToCenResponse
func (client *Client) DescribeGrantRulesToCenWithContext(ctx context.Context, request *DescribeGrantRulesToCenRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantRulesToCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about the permissions that the Alibaba Cloud account of a network instance granted to a Cloud Enterprise Network (CEN) instance in a different Alibaba Cloud account, the ID of the CEN instance, and the Alibaba Cloud account that pays the fees of the network instance.
//
// @param request - DescribeGrantRulesToResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantRulesToResourceResponse
func (client *Client) DescribeGrantRulesToResourceWithContext(ctx context.Context, request *DescribeGrantRulesToResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantRulesToResourceResponse, _err error) {
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
// Queries whether the routes of virtual private clouds (VPCs) and virtual border routers (VBRs) are advertised to the Cloud Enterprise Network (CEN) instance to which the VCPs and VBRs are attached.
//
// @param request - DescribePublishedRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePublishedRouteEntriesResponse
func (client *Client) DescribePublishedRouteEntriesWithContext(ctx context.Context, request *DescribePublishedRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribePublishedRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries overlapping routes.
//
// @param request - DescribeRouteConflictRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteConflictResponse
func (client *Client) DescribeRouteConflictWithContext(ctx context.Context, request *DescribeRouteConflictRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteConflictResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the configurations of cloud services, such as the cloud service status and the ID of the associated VPC.
//
// @param request - DescribeRouteServicesInCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteServicesInCenResponse
func (client *Client) DescribeRouteServicesInCenWithContext(ctx context.Context, request *DescribeRouteServicesInCenRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteServicesInCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the aggregate routes on an Enterprise Edition transit router.
//
// Description:
//
// You can specify the values of the **TransitRouteTableId*	- and **TransitRouteTableAggregationCidr*	- parameters to query a specified aggregate route. If you specify only the **TransitRouteTableId*	- parameter, all aggregated routes in the route table are queried.
//
// @param request - DescribeTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransitRouteTableAggregationResponse
func (client *Client) DescribeTransitRouteTableAggregationWithContext(ctx context.Context, request *DescribeTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransitRouteTableAggregationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the configuration of an aggregate route.
//
// @param request - DescribeTransitRouteTableAggregationDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransitRouteTableAggregationDetailResponse
func (client *Client) DescribeTransitRouteTableAggregationDetailWithContext(ctx context.Context, request *DescribeTransitRouteTableAggregationDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransitRouteTableAggregationDetailResponse, _err error) {
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
// Detaches a network instance from a Cloud Enterprise Network (CEN) instance.
//
// @param request - DetachCenChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachCenChildInstanceResponse
func (client *Client) DetachCenChildInstanceWithContext(ctx context.Context, request *DetachCenChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DetachCenChildInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Disables the health check feature for a virtual border router (VBR).
//
// Description:
//
// *DisableCenVbrHealthCheck*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **DescribeCenVbrHealthCheck*	- operation to query the status of health check configurations. If the health check configurations cannot be found, the health check configurations are deleted.
//
// @param request - DisableCenVbrHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCenVbrHealthCheckResponse
func (client *Client) DisableCenVbrHealthCheckWithContext(ctx context.Context, request *DisableCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *DisableCenVbrHealthCheckResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Deletes a route learning correlation.
//
// Description:
//
// *DisableTransitRouterRouteTablePropagation*	- is an synchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterRouteTablePropagations*	- operation to query the status of a route learning correlation.
//
//   - If a route learning correlation is in the **Disabling*	- state, the route learning correlation is being deleted. You can query the route learning correlation but cannot perform other operations.
//
//   - If a route learning correlation cannot be found, the route learning correlation is deleted.
//
// @param request - DisableTransitRouterRouteTablePropagationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableTransitRouterRouteTablePropagationResponse
func (client *Client) DisableTransitRouterRouteTablePropagationWithContext(ctx context.Context, request *DisableTransitRouterRouteTablePropagationRequest, runtime *dara.RuntimeOptions) (_result *DisableTransitRouterRouteTablePropagationResponse, _err error) {
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
//	  Before you delete a vSwitch from a multicast domain, make sure that the vSwitch does not contain a multicast source or a multicast member. For more information about how to remove a multicast source or member from a vSwitch, see [DeregisterTransitRouterMulticastGroupSources](https://help.aliyun.com/document_detail/468416.html) and [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/468409.html).
//
//		- If a request parameter is invalid, the system returns a request ID but does not disassociate the vSwitch from the multicast domain.
//
//		- **DisassociateTransitRouterMulticastDomain*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterMulticastDomainAssociations*	- operation to query whether a vSwitch is disassociated from the specified multicast domain.
//
//	    	- If the status is **Dissociating**, it indicates that the vSwitch is being disassociated from the specified multicast domain. You can query the vSwitch but cannot perform other operations on the vSwitch.
//
//	    	- If the vSwitch cannot be found, the vSwitch is disassociated from the multicast domain.
//
// @param request - DisassociateTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateTransitRouterMulticastDomainResponse
func (client *Client) DisassociateTransitRouterMulticastDomainWithContext(ctx context.Context, request *DisassociateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *DisassociateTransitRouterMulticastDomainResponse, _err error) {
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
// Deletes an associated forwarding correlation.
//
// Description:
//
// *DissociateTransitRouterAttachmentFromRouteTable*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterRouteTableAssociations*	- operation to query an associated forwarding correlation between a network instance connection and a route table.
//
//   - If an associated forwarding correlation is in the **Dissociating*	- state, the associated forwarding correlation is being deleted. You can query the associated forwarding correlation but cannot perform other operations.
//
//   - If an associated forwarding correlation cannot be found, the associated forwarding correlation is deleted.
//
// @param request - DissociateTransitRouterAttachmentFromRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateTransitRouterAttachmentFromRouteTableResponse
func (client *Client) DissociateTransitRouterAttachmentFromRouteTableWithContext(ctx context.Context, request *DissociateTransitRouterAttachmentFromRouteTableRequest, runtime *dara.RuntimeOptions) (_result *DissociateTransitRouterAttachmentFromRouteTableResponse, _err error) {
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
// Enables the health check feature for a virtual border router (VBR) to detect anomalies on Express Connect circuits. You can modify the health check configuration of a VBR based on business requirements.
//
// Description:
//
// You can enable the health check feature for a VBR to monitor the Express Connect circuit between your data center and Alibaba Cloud. This helps you detect connection issues in a timely manner.
//
// Before you use the health check feature, take note of the following information:
//
//   - If your VBR uses static routing, you must add a static route for the data center that is connected to the VBR after you configure the health check feature. Set the destination CIDR block to the source IP address of health checks, set the mask length to 32, and set the next hop to the IP address of the VBR on the Alibaba Cloud side.
//
//   - If your VBR uses dynamic Border Gateway Protocol (BGP) routing, you do not need to add routes for the data center.
//
//   - **EnableCenVbrHealthCheck*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **DescribeCenVbrHealthCheck*	- operation to query the status of health check configurations. If health check configurations are returned, health check is configured or modified.
//
// @param request - EnableCenVbrHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCenVbrHealthCheckResponse
func (client *Client) EnableCenVbrHealthCheckWithContext(ctx context.Context, request *EnableCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *EnableCenVbrHealthCheckResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates a route learning correlation.
//
// Description:
//
// After you establish a network instance connection on a transit router, you can create a route learning correlation for the network instance connection. Then, the routes of the connected network instance are automatically advertised to the route table of the transit router. Before you begin, we recommend that you take note of the following rules:
//
//   - You can create route learning correlations only on Enterprise Edition transit routers. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
//   - **EnableTransitRouterRouteTablePropagation*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterRouteTablePropagations*	- operation to query the route learning status between a network instance connection and a route table.
//
//   - **Enabling*	- indicates that a route learning correlation is being created between the network instance connection and route table. You can query the route learning correlation but cannot perform other operations.
//
//   - **Active*	- indicates that the route learning correlation is created between the network instance connection and route table.
//
// @param request - EnableTransitRouterRouteTablePropagationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableTransitRouterRouteTablePropagationResponse
func (client *Client) EnableTransitRouterRouteTablePropagationWithContext(ctx context.Context, request *EnableTransitRouterRouteTablePropagationRequest, runtime *dara.RuntimeOptions) (_result *EnableTransitRouterRouteTablePropagationResponse, _err error) {
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
// Grants a transit router permissions on network instances that belong to another Alibaba Cloud account. To connect a transit router of Account B to a network instance of Account A, you must use Account A to grant permissions to the transit router of Account B.
//
// Description:
//
//	  The `GrantInstanceToTransitRouter` operation can be used to grant transit routers permissions on network instances that belong to other Alibaba Cloud accounts, including virtual private clouds (VPCs), virtual border routers (VBRs), IPsec-VPN connections, and Express Connect Router (ECRs).
//
//	    To grant transit routers permissions on Cloud Connect Network (CCN) instances, call the [GrantInstanceToCbn](https://help.aliyun.com/document_detail/126141.html) operation.
//
//		- Before you call `GrantInstanceToTransitRouter`, take note of the billing rules, permission limits, and prerequisites on permission management of transit routers. For more information, see [Acquire permissions to connect to a network instance that belongs to another account](https://help.aliyun.com/document_detail/181553.html).
//
//		- Before you grant a transit router permissions on a network instance, make sure that the following requirements are met:
//
//	    The account to which the network instance belongs and the account to which the transit router belongs are of the same type.
//
//	    The ID of the Alibaba Cloud account to which the transit router belongs is obtained.
//
//	    The ID of the Cloud Enterprise Network (CEN) instance to which the Enterprise Edition transit router belongs is obtained.
//
//	    Before you grant a transit router permissions on a VBR, contact your account manager to acquire permissions on the VBR.
//
//	    Before you grant a transit router permissions on an IPsec-VPN connection, make sure that the IPsec-VPN connection is not associated with a resource.
//
//	    If the IPsec-VPN connection is attached to a VPN gateway, the IPsec-VPN connection cannot be attached to transit routers within the same account or different accounts.
//
//	    If the IPsec-VPN connection is attached to a transit router, detach the IPsec-VPN connection from the transit router. For more information, see [Delete a network instance connection](https://help.aliyun.com/document_detail/181554.html).
//
// @param request - GrantInstanceToTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantInstanceToTransitRouterResponse
func (client *Client) GrantInstanceToTransitRouterWithContext(ctx context.Context, request *GrantInstanceToTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *GrantInstanceToTransitRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries routes in route tables of network instances that point to network instance connections on Enterprise Edition transit routers.
//
// Description:
//
// Before you call the ListCenChildInstanceRouteEntriesToAttachment operation, make sure that all request parameter values are valid. If a parameter is set to an invalid value, a request ID is returned, but the routes to the network instance are not returned.
//
// @param request - ListCenChildInstanceRouteEntriesToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenChildInstanceRouteEntriesToAttachmentResponse
func (client *Client) ListCenChildInstanceRouteEntriesToAttachmentWithContext(ctx context.Context, request *ListCenChildInstanceRouteEntriesToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *ListCenChildInstanceRouteEntriesToAttachmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries quality of service (QoS) policies.
//
// Description:
//
// Before you call the **ListCenInterRegionTrafficQosPolicies*	- operation, take note of the following information:
//
//   - You must specify at least one of the **TransitRouterId*	- and **TrafficQosPolicyId*	- parameters.
//
//   - If you do not specify a QoS policy ID (**TrafficQosPolicyId**), the system returns information based on the values of the **TransitRouterId**, **TransitRouterAttachmentId**, **TrafficQosPolicyName**, and **TrafficQosPolicyDescription*	- parameters. The information about the queues in the QoS policies is not returned. In this case, the **TrafficQosQueues*	- parameter is not included in the response.
//
//   - If you specify a QoS policy ID (**TrafficQosPolicyId**), the system returns the information about the QoS policy and queues in the QoS policy. In this case, the **TrafficQosQueues*	- parameter is included in the response. If the value of the **TrafficQosQueues*	- parameter is an empty string, it indicates that the QoS policy contains only the default queue.
//
//   - Make sure that all the request parameters are valid. If a request parameter is invalid, a request ID is returned but the information about the QoS policy is not returned.
//
// @param request - ListCenInterRegionTrafficQosPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenInterRegionTrafficQosPoliciesResponse
func (client *Client) ListCenInterRegionTrafficQosPoliciesWithContext(ctx context.Context, request *ListCenInterRegionTrafficQosPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListCenInterRegionTrafficQosPoliciesResponse, _err error) {
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
// Queries the information about quality of service (QoS) queues.
//
// Description:
//
// You must specify at least one of the **TransitRouterId**, **TrafficQosPolicyId**, and **TrafficQosQueueId*	- parameters.
//
// Make sure that all the request parameters are valid. If a request parameter is invalid, a **request ID*	- is returned but the QoS queue information is not returned.
//
// @param request - ListCenInterRegionTrafficQosQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenInterRegionTrafficQosQueuesResponse
func (client *Client) ListCenInterRegionTrafficQosQueuesWithContext(ctx context.Context, request *ListCenInterRegionTrafficQosQueuesRequest, runtime *dara.RuntimeOptions) (_result *ListCenInterRegionTrafficQosQueuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the elastic network interfaces (ENIs) that can be used as multicast sources or members in a specified virtual private cloud (VPC).
//
// Description:
//
// Before you call `ListGrantVSwitchEnis`, make sure that the VPC is attached to a Cloud Enterprise Network (CEN) instance. For more information, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/468237.html).
//
// @param request - ListGrantVSwitchEnisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGrantVSwitchEnisResponse
func (client *Client) ListGrantVSwitchEnisWithContext(ctx context.Context, request *ListGrantVSwitchEnisRequest, runtime *dara.RuntimeOptions) (_result *ListGrantVSwitchEnisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the vSwitches in a virtual private cloud (VPC) that belongs to another Alibaba Cloud account and is attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// Before you call the `ListGrantVSwitchesToCen` operation, make sure that the following requirements are met:
//
//   - The permissions on the VPC are granted to the CEN instance. For more information, see [GrantInstanceToCen](https://help.aliyun.com/document_detail/126224.html).
//
//   - The VPC is attached to the CEN instance.
//
//   - For more information about how to connect an Enterprise Edition transit router to a VPC, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261358.html).
//
//   - For more information about how to connect a Basic Edition transit router to a VPC, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/65902.html).
//
// @param request - ListGrantVSwitchesToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGrantVSwitchesToCenResponse
func (client *Client) ListGrantVSwitchesToCenWithContext(ctx context.Context, request *ListGrantVSwitchesToCenRequest, runtime *dara.RuntimeOptions) (_result *ListGrantVSwitchesToCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about tags that are added to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// To call this operation, you must set at least one of **ResourceId.N*	- and **Tag.N.Key**.
//
//   - If you set only **ResourceId.N**, the tags that are added to the specified CEN instances are returned.
//
//   - If you set only **Tag.N.Key**, the CEN instances that have the specified tags are returned.
//
//   - If you set both **ResourceId.N*	- and **Tag.N.Key**, the specified tags that are added to the specified CEN instances are returned.
//
//   - Make sure that the CEN instance specified by **ResourceId.N*	- has the tag specified by **Tag.N.Key**. Otherwise, the response returns null.
//
//   - If multiple tag keys are specified, the logical operator among these tag keys is **AND**.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details about a traffic marking policy, such as the status and priority.
//
// Description:
//
// Before you call the **ListTrafficMarkingPolicies*	- operation, take note of the following limits:
//
//   - Specify at least one of the **TransitRouterId*	- and **TrafficMarkingPolicyId*	- parameters.
//
//   - If you do not specify a traffic marking policy ID (**TrafficMarkingPolicyId**), the operation queries only the information about the traffic marking policy based on the **TransitRouterId**, **TrafficMarkingPolicyName**, and **TrafficMarkingPolicyDescription*	- parameters. The **TrafficMatchRules*	- parameter that contains the information about the traffic classification rules is not returned.
//
//   - If you specify a traffic marking policy ID (**TrafficMarkingPolicyId**), the operation queries the information about the traffic marking policy and traffic classification rules. The **TrafficMatchRules*	- parameter is returned in the response. If the value of the **TrafficMatchRules*	- parameter is an empty array, the traffic marking policy does not contain a traffic classification rule.
//
// @param request - ListTrafficMarkingPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficMarkingPoliciesResponse
func (client *Client) ListTrafficMarkingPoliciesWithContext(ctx context.Context, request *ListTrafficMarkingPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListTrafficMarkingPoliciesResponse, _err error) {
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
// Queries the zones that support Enterprise Edition transit routers in a region.
//
// Description:
//
//	  You can call the **ListTransitRouterAvailableResource*	- operation to query the zones that support Enterprise Edition transit routers in a specified region.
//
//	    	- If you do not set **SupportMulticast*	- to **true**, general-purpose zones that support Enterprise Edition transit routers are queried.
//
//	    	- If you set **SupportMulticast*	- to **true**, zones in which Enterprise Edition transit routers support multicast are queried.
//
//		- On May 31, 2022, VPC-connected Enterprise Edition transit routers were optimized. Optimized Enterprise Edition transit routers do not require you to specify the primary and secondary zones when you connect VPCs to the Enterprise Edition transit routers. You can specify one or more zones.
//
//	    	- If your Enterprise Edition transit router has not been optimized, you must specify the primary and secondary zones when you connect a VPC to your Enterprise Edition transit router. After you call **ListTransitRouterAvailableResource**, you can call **MasterZones*	- and **SlaveZones*	- to query the primary and secondary zones.
//
//	    	- If your Enterprise Edition transit router has been optimized, you can specify a zone as needed when you connect a VPC to your Enterprise Edition transit router. After you call **ListTransitRouterAvailableResource**, you can call **AvailableZones*	- to query the zones.
//
// For more information about the optimization, see [Announcement: Optimization on VPC-connected Enterprise Edition transit routers](https://help.aliyun.com/document_detail/434191.html).
//
// @param request - ListTransitRouterAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterAvailableResourceResponse
func (client *Client) ListTransitRouterAvailableResourceWithContext(ctx context.Context, request *ListTransitRouterAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterAvailableResourceResponse, _err error) {
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
// Queries the CIDR blocks of a transit router.
//
// @param request - ListTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterCidrResponse
func (client *Client) ListTransitRouterCidrWithContext(ctx context.Context, request *ListTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterCidrResponse, _err error) {
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
// Queries how a CIDR block is allocated.
//
// @param request - ListTransitRouterCidrAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterCidrAllocationResponse
func (client *Client) ListTransitRouterCidrAllocationWithContext(ctx context.Context, request *ListTransitRouterCidrAllocationRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterCidrAllocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about Express Connect Router (ECR) connections on an Enterprise Edition transit router, such as the connection status, connection ID, and the payer of instance fees.
//
// Description:
//
// You can use the following methods to query ECR connection information:
//
//   - Specify the ID of an Enterprise Edition transit router.
//
//   - Specify the ID of an Enterprise Edition transit router and the ID of the region in which the Enterprise Edition transit router is deployed.
//
//   - Configure the **TransitRouterAttachmentId*	- parameter to specify the ECR connection that you want to query.
//
// @param request - ListTransitRouterEcrAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterEcrAttachmentsResponse
func (client *Client) ListTransitRouterEcrAttachmentsWithContext(ctx context.Context, request *ListTransitRouterEcrAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterEcrAttachmentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries whether a multicast domain is associated with a vSwitch.
//
// Description:
//
//	  You must set at least **TransitRouterMulticastDomainId*	- and **TransitRouterAttachmentId**. If you set **TransitRouterAttachmentId**, the information about the vSwitches in a virtual private cloud (VPC) that are associated with a multicast domain is returned. If you set **TransitRouterMulticastDomainId**, the information about the vSwitches that are associated with a multicast domain is returned.
//
//		- Before you call **ListTransitRouterMulticastDomainAssociations**, make sure that all the request parameters are valid. If a request parameter is invalid, the system returns a **request ID*	- but does not return the vSwitches that are associated with the multicast domain.
//
// @param request - ListTransitRouterMulticastDomainAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainAssociationsResponse
func (client *Client) ListTransitRouterMulticastDomainAssociationsWithContext(ctx context.Context, request *ListTransitRouterMulticastDomainAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainAssociationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

// @param request - ListTransitRouterMulticastDomainVSwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainVSwitchesResponse
func (client *Client) ListTransitRouterMulticastDomainVSwitchesWithContext(ctx context.Context, request *ListTransitRouterMulticastDomainVSwitchesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainVSwitchesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about a multicast domain, such as the status, ID, and description.
//
// Description:
//
//	  If you configure one of the RegionId and CenId parameters, you must configure the other parameter. Otherwise, no information about the multicast domain is returned. You can configure only one of the TransitRouterId and TransitRouterMulticastDomainId parameters.
//
//		- Make sure that all the request parameters are valid. If a request parameter is invalid, a **request ID*	- is returned but the information about the multicast domain is not returned.
//
// @param request - ListTransitRouterMulticastDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainsResponse
func (client *Client) ListTransitRouterMulticastDomainsWithContext(ctx context.Context, request *ListTransitRouterMulticastDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the detailed information about the multicast members and sources in a multicast domain.
//
// Description:
//
// You can call the `ListTransitRouterMulticastGroups` operation to query the multicast sources and members in a multicast domain. Multicast sources and members are also known as multicast resources.
//
//   - If you set **GroupIpAddress**, the system queries multicast resources in the multicast domain by multicast group.
//
//   - If you set **VSwitchIds**, the system queries multicast resources in the multicast domain by vSwitch.
//
//   - If you set **PeerTransitRouterMulticastDomains**, the system queries multicast resources that are also deployed in a different region.
//
//   - If you set **ResourceType**, the system queries the multicast resources of the specified type in the multicast domain.
//
//   - If you set **ResourceId**, the system queries multicast resources by resource.
//
//   - If you set only **TransitRouterMulticastDomainId**, the system queries all the multicast resources in the multicast domain.
//
// @param request - ListTransitRouterMulticastGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastGroupsResponse
func (client *Client) ListTransitRouterMulticastGroupsWithContext(ctx context.Context, request *ListTransitRouterMulticastGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries inter-region connections on an Enterprise Edition transit router.
//
// Description:
//
// You can use the following methods to query inter-region connections on an Enterprise Edition transit router:
//
//   - Query all inter-region connections on an Enterprise Edition transit router by specifying the ID of the Enterprise Edition transit router.
//
//   - Query all inter-region connections on an Enterprise Edition transit router by specifying the ID of the Cloud Enterprise Network (CEN) instance and the ID of the region where the transit router is deployed.
//
// @param request - ListTransitRouterPeerAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterPeerAttachmentsResponse
func (client *Client) ListTransitRouterPeerAttachmentsWithContext(ctx context.Context, request *ListTransitRouterPeerAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterPeerAttachmentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the prefix lists that are associated with an Enterprise Edition transit router.
//
// @param request - ListTransitRouterPrefixListAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterPrefixListAssociationResponse
func (client *Client) ListTransitRouterPrefixListAssociationWithContext(ctx context.Context, request *ListTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterPrefixListAssociationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the details about routes in the route tables of an Enterprise Edition transit router.
//
// @param request - ListTransitRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteEntriesResponse
func (client *Client) ListTransitRouterRouteEntriesWithContext(ctx context.Context, request *ListTransitRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteEntriesResponse, _err error) {
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
// Queries the associated forwarding correlations that are created for a route table of an Enterprise Edition transit router or a network instance connection.
//
// Description:
//
// When you call **ListTransitRouterRouteTableAssociations**, you must set at least one of **TransitRouterRouteTableId*	- and **TransitRouterAttachmentId**.
//
//   - If you set only **TransitRouterRouteTableId**, the network instance connections that are in associated forwarding correlation with a route table of an Enterprise Edition transit router are queried.
//
//   - If you set only **TransitRouterAttachmentId**, the route table of an Enterprise Edition transit router that is in associated forwarding correlation with a network instance connection is queried.
//
//   - If you set both **TransitRouterRouteTableId*	- and **TransitRouterAttachmentId**, the associated forwarding correlations between a specified network instance connection and a specified route table of an Enterprise Edition transit router are queried.
//
//   - If an associated forwarding correlation is created between the network instance connection and the route table of the Enterprise Edition transit router, the information about the associated forwarding correlation is returned.
//
//   - If no associated forwarding correlation is created between the network instance connection and the route table of the Enterprise Edition transit router, **TransitRouterAssociations*	- in the response is empty.
//
// @param request - ListTransitRouterRouteTableAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTableAssociationsResponse
func (client *Client) ListTransitRouterRouteTableAssociationsWithContext(ctx context.Context, request *ListTransitRouterRouteTableAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTableAssociationsResponse, _err error) {
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
// Queries the route learning correlations of an Enterprise Edition transit router.
//
// @param request - ListTransitRouterRouteTablePropagationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTablePropagationsResponse
func (client *Client) ListTransitRouterRouteTablePropagationsWithContext(ctx context.Context, request *ListTransitRouterRouteTablePropagationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTablePropagationsResponse, _err error) {
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
// Queries the route tables of an Enterprise Edition transit router.
//
// @param request - ListTransitRouterRouteTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTablesResponse
func (client *Client) ListTransitRouterRouteTablesWithContext(ctx context.Context, request *ListTransitRouterRouteTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTablesResponse, _err error) {
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
// Queries the virtual border router (VBR) connections on an Enterprise Edition transit router.
//
// Description:
//
// You can use the following methods to query VBR connections on an Enterprise Edition transit router:
//
//   - Specify the ID of the Enterprise Edition transit router.
//
//   - Specify the ID of the relevant Cloud Enterprise Network (CEN) instance and the region ID of the Enterprise Edition transit router.
//
// @param request - ListTransitRouterVbrAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVbrAttachmentsResponse
func (client *Client) ListTransitRouterVbrAttachmentsWithContext(ctx context.Context, request *ListTransitRouterVbrAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVbrAttachmentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the status, billing method, zones, vSwitches, and elastic network interfaces (ENIs) of virtual private cloud (VPC) connections.
//
// Description:
//
// You can use the following methods to query VPC connections on an Enterprise Edition transit router:
//
//   - Specify the ID of the Enterprise Edition transit router.
//
//   - Specify the ID of the relevant Cloud Enterprise Network (CEN) instance and the region ID of the Enterprise Edition transit router.
//
//   - Specify the ID of the region where the Enterprise Edition transit router is deployed.
//
// @param request - ListTransitRouterVpcAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVpcAttachmentsResponse
func (client *Client) ListTransitRouterVpcAttachmentsWithContext(ctx context.Context, request *ListTransitRouterVpcAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVpcAttachmentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about VPN attachments, such as the status and billing method of a VPN attachment, and the ID of an IPsec-VPN connection.
//
// Description:
//
// You can use the following methods to call the ListTransitRouterVpnAttachments operation:
//
//   - Specify only the **TransitRouterAttachmentId*	- parameter to query the information about a VPN attachment.
//
//   - Specify only the **TransitRouterId*	- parameter to query the information about all VPN attachments on a transit router.
//
//   - Specify the **CenId*	- and **RegionId*	- parameter to query the information about VPN attachments in a specified region.
//
// Before you call the **ListTransitRouterVpnAttachments*	- operation, make sure that all request parameters are valid. If a request parameter is invalid, a **request ID*	- is returned, but the information about the VPN attachments is not returned.
//
// @param request - ListTransitRouterVpnAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVpnAttachmentsResponse
func (client *Client) ListTransitRouterVpnAttachmentsWithContext(ctx context.Context, request *ListTransitRouterVpnAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVpnAttachmentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Queries the information about transit routers that are connected to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can set the **RegionId*	- and **TransitRouterId*	- parameters based on your requirements.
//
//   - If you do not set **RegionId*	- or **TransitRouterId**, the system queries all transit routers that are connected to the specified CEN instance.
//
//   - If you set only **RegionId**, the system queries transit routers that are deployed in the specified region.
//
//   - If you set only **TransitRouterId**, the system queries the specified transit router.
//
//   - If you set both **RegionId*	- and **TransitRouterId**, the system queries the specified transit router in the specified region.
//
// @param request - ListTransitRoutersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRoutersResponse
func (client *Client) ListTransitRoutersWithContext(ctx context.Context, request *ListTransitRoutersRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRoutersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// *ModifyCenAttribute*	- is an asynchronous operation. After you send a request, the system returns the **request ID*	- but the operation is still being performed in the system background. You can call **DescribeCens*	- to query the status of a CEN instance.
//
//   - If a CEN instance is in the **Modifying*	- state, the CEN instance is being modified. You can query the CEN instance but cannot perform other operations.
//
//   - If a CEN instance is in the **Active*	- state, the CEN instance is modified.
//
// @param request - ModifyCenAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenAttributeResponse
func (client *Client) ModifyCenAttributeWithContext(ctx context.Context, request *ModifyCenAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the name and description of a bandwidth plan.
//
// @param request - ModifyCenBandwidthPackageAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenBandwidthPackageAttributeResponse
func (client *Client) ModifyCenBandwidthPackageAttributeWithContext(ctx context.Context, request *ModifyCenBandwidthPackageAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenBandwidthPackageAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the maximum bandwidth of a bandwidth plan.
//
// @param request - ModifyCenBandwidthPackageSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenBandwidthPackageSpecResponse
func (client *Client) ModifyCenBandwidthPackageSpecWithContext(ctx context.Context, request *ModifyCenBandwidthPackageSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenBandwidthPackageSpecResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies a routing policy of a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// `ModifyCenRouteMap` is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the `DescribeCenRouteMaps` operation to query the status of a routing policy.
//
//   - **Modifying**: indicates that the system is modifying the routing policy. You can only query the routing policy, but cannot perform other operations.
//
//   - **Active**: indicates that the routing policy is modified.
//
// @param request - ModifyCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenRouteMapResponse
func (client *Client) ModifyCenRouteMapWithContext(ctx context.Context, request *ModifyCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenRouteMapResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the name, description, and capture window of a flow log.
//
// Description:
//
// `ModifyFlowLogAttribute` is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//   - If a flow log is in the **Modifying*	- state, the flow log is being modified. In this case, you can query the flow log but cannot perform other operations.
//
//   - If a flow log is in the **Active*	- state, the flow log is modified.
//
// @param request - ModifyFlowLogAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttributeWithContext(ctx context.Context, request *ModifyFlowLogAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyFlowLogAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the name and description of a traffic classification rule.
//
// @param request - ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) ModifyTrafficMatchRuleToTrafficMarkingPolicyWithContext(ctx context.Context, request *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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
// Edit an aggregate route.
//
// @param tmpReq - ModifyTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTransitRouteTableAggregationResponse
func (client *Client) ModifyTransitRouteTableAggregationWithContext(ctx context.Context, tmpReq *ModifyTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouteTableAggregationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the CIDR block of a transit router.
//
// Description:
//
//	  Before you modify the CIDR block of a transit router, we recommend that you read the [limits on transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
//		- If IP addresses within the CIDR block have been allocated to network instances, you cannot modify the CIDR block.
//
//		- When you call **ModifyTransitRouterCidr**, if no parameter of the **PublishCidrRoute*	- operation is modified, ModifyTransitRouterCidr is a synchronous operation. After you call the operation, the new settings are immediately applied.
//
//		- If a parameter of the **PublishCidrRoute*	- operation is modified, **ModifyTransitRouterCidr*	- is an asynchronous operation. After you call the operation, the request ID (**RequestId**) is returned but the operation is still being performed in the system background. You can call **ListTransitRouterCidr*	- to query the status of the CIDR block of the transit router.
//
//	    	- If the CIDR block of the transit router remains unchanged, the CIDR block is still being modified.
//
//	    	- If the CIDR block of the transit router is changed to the one that you specify in the request, the CIDR block has been modified.
//
// @param request - ModifyTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTransitRouterCidrResponse
func (client *Client) ModifyTransitRouterCidrWithContext(ctx context.Context, request *ModifyTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouterCidrResponse, _err error) {
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
// You can call the ModifyTransitRouterMulticastDomain operation to modify the name, description, and feature options of a multicast domain.
//
// @param request - ModifyTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTransitRouterMulticastDomainResponse
func (client *Client) ModifyTransitRouterMulticastDomainWithContext(ctx context.Context, request *ModifyTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouterMulticastDomainResponse, _err error) {
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
// Moves a Cloud Enterprise Network (CEN) instance or a bandwidth plan to another resource group.
//
// Description:
//
// By default, CEN instances and bandwidth plans are in the default resource group. You can call the `MoveResourceGroup` operation to move CEN instances or bandwidth plans to another resource group.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithContext(ctx context.Context, request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
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
// Activates the transit router feature.
//
// Description:
//
// You can call the `OpenTransitRouterService` operation to activate the transit router feature free of charge. After the `OpenTransitRouterService` operation succeeds, an order is automatically generated. You can use the returned order ID to query the order information in [Alibaba Cloud User Center](https://usercenter2-intl.aliyun.com/billing/#/account/overview).
//
// @param request - OpenTransitRouterServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenTransitRouterServiceResponse
func (client *Client) OpenTransitRouterServiceWithContext(ctx context.Context, request *OpenTransitRouterServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenTransitRouterServiceResponse, _err error) {
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
// Cloud Enterprise Network (CEN) supports route advertisement. You can call this operation to advertise routes of virtual private clouds (VPCs) or virtual border routers (VBRs) attached to a CEN instance to the CEN instance. Other network instances attached to the CEN instance can learn the routes if route conflicts do not exist.
//
// Description:
//
// The following table describes whether routes of different types are advertised to CEN by default. You can call the PublishRouteEntries operation to advertise routes to CEN.
//
// |Route|Network instance|Advertised to CEN by default|
//
// |---|---|---|
//
// |Routes that route network traffic to Elastic Compute Service (ECS) instances|VPC|No|
//
// |Routes that route network traffic to VPN gateways|VPC|No|
//
// |Routes that route network traffic to high-availability virtual IP addresses (HAVIPs)|VPC|No|
//
// |Routes that route network traffic to router interfaces|VPC|No|
//
// |Routes that route network traffic to elastic network interfaces (ENIs)|VPC|No|
//
// |Routes that route network traffic to IPv6 gateways|VPC|No|
//
// |Routes that route network traffic to NAT gateways|VPC|No|
//
// |System routes of VPCs|VPC|Yes|
//
// |Routes that route network traffic to data centers|VBR|Yes|
//
// |Border Gateway Protocol (BGP) routes|VBR|Yes|
//
// @param request - PublishRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRouteEntriesResponse
func (client *Client) PublishRouteEntriesWithContext(ctx context.Context, request *PublishRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *PublishRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Re-advertises an aggregate route.
//
// @param request - RefreshTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshTransitRouteTableAggregationResponse
func (client *Client) RefreshTransitRouteTableAggregationWithContext(ctx context.Context, request *RefreshTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *RefreshTransitRouteTableAggregationResponse, _err error) {
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
// Creates or adds a multicast member.
//
// Description:
//
// Enterprise Edition transit routers support only elastic network interfaces (ENIs) as multicast members. You can call the `RegisterTransitRouterMulticastGroupMembers` operation to specify an ENI in the current region or a different region as a multicast member.
//
//   - If you specify a value for the **NetworkInterfaceIds*	- parameter, an ENI in the current region is to be specified as a multicast member. Make sure that the ENI and vSwitch are associated with the multicast group. For more information, see [AssociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429778.html).
//
//   - If you specify a value for the **PeerTransitRouterMulticastDomains**, a multicast member in a multicast group that belongs to another region but has the same IP address as the current multicast group is to be specified as a multicast member for the current multicast group. Make sure that an inter-region connection is established between the regions. For more information, see [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html).
//
//     For example, you created Multicast Group 1 in Multicast Domain 1, which is in the China (Hangzhou) region. You created Multicast Group 2 in Multicast Domain 2, which is in the China (Shanghai) region. Multicast Group 1 and Multicast Group 2 use the same multicast IP address, and Multicast Member 2 is in Multicast Group 2 in the China (Shanghai) region. If you call the `RegisterTransitRouterMulticastGroupMembers` operation to add multicast members to Multicast Group 1 in the China (Hangzhou) region and set **PeerTransitRouterMulticastDomains*	- to the ID of Multicast Group 2, which is in the China (Shanghai) region, Multicast Member 2, which is in Multicast Domain 2 in the China (Shanghai) region is added to Multicast Group 1 in the China (Hangzhou) region.
//
//   - `RegisterTransitRouterMulticastGroupMembers` is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of a multicast member.
//
//   - If the multicast member is in the **Registering**, the multicast member is being created. In this case, you can query the multicast member but cannot perform other operations on the multicast member.
//
//   - If the multicast member is in the **Registered*	- state, the multicast member is created.
//
// @param request - RegisterTransitRouterMulticastGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterTransitRouterMulticastGroupMembersResponse
func (client *Client) RegisterTransitRouterMulticastGroupMembersWithContext(ctx context.Context, request *RegisterTransitRouterMulticastGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *RegisterTransitRouterMulticastGroupMembersResponse, _err error) {
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
// Creates a multicast source for a one-to-many multicast network.
//
// Description:
//
//	  You can specify only elastic network interfaces (ENIs) as multicast sources.
//
//		- `RegisterTransitRouterMulticastGroupSources` is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of a multicast source.
//
//	    	- If a multicast source is in the **Registering*	- state, the multicast source is being created. You can query the multicast source but cannot perform other operations on the multicast source.
//
//	    	- If a multicast source is in the **Registered*	- state, the multicast source is created.
//
// ### Prerequisite
//
// Before you call `RegisterTransitRouterMulticastGroupSources`, make sure that the vSwitch on which the ENI is created is associated with the multicast domain. For more information, see [AssociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429778.html).
//
// @param request - RegisterTransitRouterMulticastGroupSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterTransitRouterMulticastGroupSourcesResponse
func (client *Client) RegisterTransitRouterMulticastGroupSourcesWithContext(ctx context.Context, request *RegisterTransitRouterMulticastGroupSourcesRequest, runtime *dara.RuntimeOptions) (_result *RegisterTransitRouterMulticastGroupSourcesResponse, _err error) {
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
// Deletes specified traffic classification rules from a traffic marking policy.
//
// Description:
//
//	  When you call **RemoveTrafficMatchRuleFromTrafficMarkingPolicy**, take note of the following rules:
//
//	    	- If you specify the ID of a traffic classification rule in the **TrafficMarkRuleIds*	- parameter, the specified traffic classification rule is deleted.
//
//	    	- If you do not specify a traffic classification rule ID in the **TrafficMarkRuleIds*	- parameter, no operation is performed after you call this operation.
//
//	    If you want to delete a traffic classification rule, you must specify the rule ID before you call this operation.
//
//		- **RemoveTrafficMatchRuleFromTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of a traffic classification rule.
//
//	    	- If a traffic classification rule is in the **Deleting*	- state, the traffic classification rule is being deleted. In this case, you can query the traffic classification rule but cannot perform other operations.
//
//	    	- If a traffic classification rule cannot be found, the traffic classification rule is deleted.
//
// @param request - RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse
func (client *Client) RemoveTrafficMatchRuleFromTrafficMarkingPolicyWithContext(ctx context.Context, request *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
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
// Removes a traffic classification rule from a traffic marking policy.
//
// Description:
//
// ### [](#)Precautions
//
// The **RemoveTraficMatchRuleFromTrafficMarkingPolicy*	- operation is deprecated and will be discontinued soon. To delete a traffic classification rule, call the [RemoveTrafficMatchRuleFromTrafficMarkingPolicy](https://help.aliyun.com/document_detail/452726.html) operation. Maintenance on this document has stopped.
//
// @param request - RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse
func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicyWithContext(ctx context.Context, request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
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
// Associates a network instance connection with another route table of a transit router.
//
// @param request - ReplaceTransitRouterRouteTableAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceTransitRouterRouteTableAssociationResponse
func (client *Client) ReplaceTransitRouterRouteTableAssociationWithContext(ctx context.Context, request *ReplaceTransitRouterRouteTableAssociationRequest, runtime *dara.RuntimeOptions) (_result *ReplaceTransitRouterRouteTableAssociationResponse, _err error) {
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
// Connects an on-premises network to a cloud service.
//
// Description:
//
// Cloud services refer to Alibaba Cloud services that use the 100.64.0.0/10 CIDR block to provide services. These cloud services include Object Storage Service (OSS), Simple Log Service (SLS), and Data Transmission Service (DTS). If your on-premises network needs to access a cloud service, you must attach the virtual border router (VBR) or Cloud Connect Network (CCN) instance that is connected to your on-premises network to a Cloud Enterprise Network (CEN) instance. In addition, you must attach a virtual private cloud (VPC) that is deployed in the same region as the cloud service to the CEN instance. This way, your on-premises network can connect to the VPC that is deployed in the same region as the cloud service and access the cloud service through the VPC.
//
//   - This operation is supported only by Basic Edition transit routers. An on-premises network associated with a VBR can use CEN to access only a cloud service that is deployed in the same region.
//
//     For example, if cloud services are deployed in the China (Beijing) region, only on-premises networks connected to VBRs in the China (Beijing) region can access the cloud services.
//
//   - **ResolveAndRouteServiceInCen*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call **DescribeRouteServicesInCen*	- to query the status of a cloud service.
//
//   - If the cloud service is in the **Creating*	- state, the connection to the cloud service is being created. In this case, you can query the cloud service but cannot perform other operations.
//
//   - If the cloud service is in the **Active*	- state, the connection to the cloud service is created.
//
//   - If the cloud service is in the **Failed*	- state, the connection to the cloud service failed.
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the following conditions are met:
//
//   - The VBR or CCN instance to which your on-premises network is connected is attached to a CEN instance.
//
//   - A VPC that is deployed in the same region as the cloud service is attached to the CEN instance. For more information, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/65902.html).
//
// @param request - ResolveAndRouteServiceInCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResolveAndRouteServiceInCenResponse
func (client *Client) ResolveAndRouteServiceInCenWithContext(ctx context.Context, request *ResolveAndRouteServiceInCenRequest, runtime *dara.RuntimeOptions) (_result *ResolveAndRouteServiceInCenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Revokes the permissions that a transit router has on network instances that belong to another Alibaba Cloud account.
//
// Description:
//
// `RevokeInstanceFromTransitRouter` can be used to revoke permissions on virtual private clouds (VPCs), virtual border routers (VBRs), IPsec-VPN connections, and Express Connect Router (ECRs) that belong to another Alibaba Cloud account.
//
// To revoke permissions on Cloud Connect Network (CCN) instances that belong to another Alibaba Cloud account, call the [RevokeInstanceFromCbn](https://help.aliyun.com/document_detail/126142.html) operation.
//
// ### [](#)Prerequisites
//
// Before you call `RevokeInstanceFromTransitRouter`, you must detach the network instances from the transit router.
//
//   - For more information about how to detach VPCs from Enterprise Edition transit routers, see [DeleteTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261220.html).
//
//   - For more information about how to detach VBRs from Enterprise Edition transit routers, see [DeleteTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/261223.html).
//
//   - For more information about how to detach IPsec-VPN connections from Enterprise Edition transit routers, see [DeleteTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/443992.html).
//
//   - For more information about how to detach ECRs from Enterprise Edition transit routers, see [DeleteTransitRouterEcrAttachment](https://help.aliyun.com/document_detail/443992.html).
//
//   - For more information about how to detach network instances from Basic Edition transit routers, see [DetachCenChildInstance](https://help.aliyun.com/document_detail/65915.html).
//
// @param request - RevokeInstanceFromTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromTransitRouterResponse
func (client *Client) RevokeInstanceFromTransitRouterWithContext(ctx context.Context, request *RevokeInstanceFromTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *RevokeInstanceFromTransitRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Configures PrivateZone.
//
// Description:
//
// Alibaba Cloud DNS PrivateZone (PrivateZone) is an Alibaba Cloud private domain name resolution and management service based on Virtual Private Cloud (VPC). After you attach virtual border routers (VBRs) and Cloud Connect Network (CCN) instances to a Cloud Enterprise Network (CEN) instance, you can enable the on-premises networks connected to the VBRs and CCN instances to access PrivateZone through the CEN instance.
//
// #### Usage notes
//
// - The on-premises networks connected to VBRs or CCN instances must be deployed in the same region as the PrivateZone service. For example, if the PrivateZone service is deployed in the China (Beijing) region, only on-premises networks connected to VBRs or CCN instances in the China (Beijing) region can access the PrivateZone service.
//
// - **RoutePrivateZoneInCenToVpc*	- is an asynchronous operation. After you send a request, the **request ID*	- is returned but the operation is still being performed in the system background. You can call **DescribeCenPrivateZoneRoutes*	- to query the status of PrivateZone.
//
//   - If PrivateZone is in the **Creating*	- state, access to PrivateZone is being configured. In this case, you can query PrivateZone configurations but cannot perform other operations.
//
//   - If PrivateZone is in the **Active*	- state, access to PrivateZone is enabled.
//
//   - If PrivateZone is in the **Failed*	- state, configurations of access to PrivateZone failed.
//
// #### Prerequisites
//
// Before you call **RoutePrivateZoneInCenToVpc**, make sure that the following conditions are met:
//
// - PrivateZone is deployed. For more information, see [PrivateZone quick start](https://help.aliyun.com/document_detail/64627.html).
//
// - The following network instances are attached to the same CEN instance: the VPC that is associated with the PrivateZone service, and the VBR and CCN instance that want to access the PrivateZone service. For more information, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/468684.html).
//
// - If your on-premises network uses a CCN instance to connect to Alibaba Cloud and the account that owns the CCN instance is different from the account that owns the VPC or CEN instance, you must grant the CCN instance required permissions. For more information, see [Grant permissions to CCN](https://help.aliyun.com/document_detail/181654.html).
//
// @param request - RoutePrivateZoneInCenToVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RoutePrivateZoneInCenToVpcResponse
func (client *Client) RoutePrivateZoneInCenToVpcWithContext(ctx context.Context, request *RoutePrivateZoneInCenToVpcRequest, runtime *dara.RuntimeOptions) (_result *RoutePrivateZoneInCenToVpcResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Configures, modifies, or deletes the bandwidth of inter-region connections for a Basic Edition transit router.
//
// Description:
//
// This operation is used to manage bandwidth of inter-region connections only for Basic Edition transit routers.
//
// ### [](#)Prerequisites
//
// The Cloud Enterprise Network (CEN) instance is associated with a bandwidth plan. For more information, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html) and [AssociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65934.html).
//
// You can call the **SetCenInterRegionBandwidthLimit*	- operation to configure, change, or remove the bandwidth limit of an inter-region connection.
//
//   - If you set **BandwidthLimit*	- to a value other than 0, the bandwidth of the inter-region connection is set to the specified value.
//
//   - If you set **BandwidthLimit*	- to 0, the bandwidth of the inter-region connection is no longer limited.
//
// ### [](#)Limits
//
//   - The bandwidth limit of an inter-region connection cannot exceed the bandwidth limit of the associated bandwidth plan.
//
//   - The sum of bandwidth limits of all inter-region connections cannot exceed the bandwidth limit of the associated bandwidth plan.
//
//   - If bandwidth multiplexing is enabled for an inter-region connection, you cannot change the bandwidth of the inter-region connection.
//
//   - The **SetCenInterRegionBandwidthLimit*	- operation can be used to configure, modify, or delete the bandwidth of inter-region connections only for Basic Edition transit routers.
//
//     To configure, modify, or delete the bandwidth of inter-region connections for Enterprise Edition transit routers, call the [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html), [UpdateTransitRouterPeerAttachmentAttribute](https://help.aliyun.com/document_detail/261229.html), or [DeleteTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261227.html) operation.
//
// @param request - SetCenInterRegionBandwidthLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCenInterRegionBandwidthLimitResponse
func (client *Client) SetCenInterRegionBandwidthLimitWithContext(ctx context.Context, request *SetCenInterRegionBandwidthLimitRequest, runtime *dara.RuntimeOptions) (_result *SetCenInterRegionBandwidthLimitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Creates tags and adds them to a resource.
//
// Description:
//
//	  Each tag consists of a tag key and a tag value. When you add a tag, you must specify the tag key and tag value.
//
//		- If you want to add multiple tags to a Cloud Enterprise Network (CEN) instance, each tag key must be unique.
//
//		- You can add at most 20 tags to a CEN instance.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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
// 进行云企业网预付费带宽包临时升配
//
// @param request - TempUpgradeCenBandwidthPackageSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TempUpgradeCenBandwidthPackageSpecResponse
func (client *Client) TempUpgradeCenBandwidthPackageSpecWithContext(ctx context.Context, request *TempUpgradeCenBandwidthPackageSpecRequest, runtime *dara.RuntimeOptions) (_result *TempUpgradeCenBandwidthPackageSpecResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Disassociates a Cloud Enterprise Network (CEN) from a bandwidth plan. After the disassociation, the bandwidth can be associated with another CEN instance.
//
// Description:
//
// No inter-region connections are configured in the bandwidth plan. For more information about how to delete inter-region connections, see [SetCenInterRegionBandwidthLimit](https://help.aliyun.com/document_detail/65942.html).
//
// @param request - UnassociateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassociateCenBandwidthPackageResponse
func (client *Client) UnassociateCenBandwidthPackageWithContext(ctx context.Context, request *UnassociateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *UnassociateCenBandwidthPackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// 删除PrivateZone
//
// @param request - UnroutePrivateZoneInCenToVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnroutePrivateZoneInCenToVpcResponse
func (client *Client) UnroutePrivateZoneInCenToVpcWithContext(ctx context.Context, request *UnroutePrivateZoneInCenToVpcRequest, runtime *dara.RuntimeOptions) (_result *UnroutePrivateZoneInCenToVpcResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// The ID of the request.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
// Modifies the name and description of a quality of service (QoS) policy.
//
// @param request - UpdateCenInterRegionTrafficQosPolicyAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCenInterRegionTrafficQosPolicyAttributeResponse
func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttributeWithContext(ctx context.Context, request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, _err error) {
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
// Modifies the name, description, inter-region bandwidth, and Differentiated Services Code Point (DSCP) value of a quality of service (QoS) queue.
//
// @param request - UpdateCenInterRegionTrafficQosQueueAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCenInterRegionTrafficQosQueueAttributeResponse
func (client *Client) UpdateCenInterRegionTrafficQosQueueAttributeWithContext(ctx context.Context, request *UpdateCenInterRegionTrafficQosQueueAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCenInterRegionTrafficQosQueueAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the name, description, and traffic classification rules of a traffic marking policy.
//
// @param request - UpdateTrafficMarkingPolicyAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrafficMarkingPolicyAttributeResponse
func (client *Client) UpdateTrafficMarkingPolicyAttributeWithContext(ctx context.Context, request *UpdateTrafficMarkingPolicyAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTrafficMarkingPolicyAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the name and description of a transit router.
//
// Description:
//
// *UpdateTransitRouter*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouters*	- operation to query the status of a transit router.
//
//   - If a transit router is in the **Modifying*	- state, the configuration of the transit router is being modified. You can query the transit router but cannot perform other operations.
//
//   - If a transit router is in the **Active*	- state, the configuration of the transit router is modified.
//
// @param request - UpdateTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterResponse
func (client *Client) UpdateTransitRouterWithContext(ctx context.Context, request *UpdateTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterResponse, _err error) {
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
// Modifies the name and description of an Express Connect Router (ECR) connection on a Enterprise Edition transit router.
//
// Description:
//
// UpdateTransitRouterEcrAttachmentAttribute is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the ListTransitRouterEcrAttachments operation to query the status of an ECR connection.
//
// If an ECR connection is in the Modifying state, the ECR connection is being modified. In this case, you can query the ECR connection but cannot perform other operations on the ECR connection. If an ECR connection is in the Attached state, the ECR connection is modified.
//
// @param request - UpdateTransitRouterEcrAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterEcrAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterEcrAttachmentAttributeWithContext(ctx context.Context, request *UpdateTransitRouterEcrAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterEcrAttachmentAttributeResponse, _err error) {
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
// Modifies an inter-region connection on an Enterprise Edition transit router.
//
// Description:
//
// *UpdateTransitRouterPeerAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterPeerAttachments*	- operation to query the status of an inter-region connection.
//
//   - If an inter-region connection is in the **Modifying*	- state, the inter-region connection is being modified. You can query the inter-region connection but cannot perform other operations.
//
//   - If an inter-region connection is in the **Attached*	- state, the inter-region connection is modified.
//
// @param request - UpdateTransitRouterPeerAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterPeerAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterPeerAttachmentAttributeWithContext(ctx context.Context, request *UpdateTransitRouterPeerAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterPeerAttachmentAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the name and description of a route in a route table of an Enterprise Edition transit router.
//
// @param request - UpdateTransitRouterRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterRouteEntryResponse
func (client *Client) UpdateTransitRouterRouteEntryWithContext(ctx context.Context, request *UpdateTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterRouteEntryResponse, _err error) {
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
// Modifies the name and description of a route table of an Enterprise Edition transit router and enables or disables multi-region equal-cost multi-path (ECMP) routing.
//
// @param request - UpdateTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterRouteTableResponse
func (client *Client) UpdateTransitRouterRouteTableWithContext(ctx context.Context, request *UpdateTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterRouteTableResponse, _err error) {
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
// Modifies the name, description, and enables or disables automatic route advertisement for a virtual border router (VBR) connection on an Enterprise Edition transit router.
//
// Description:
//
// *UpdateTransitRouterVbrAttachmentAttribute*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterVbrAttachments*	- operation to query the status of a VBR connection.
//
//   - If a VBR connection is in the **Modifying*	- state, the VBR connection is being modified. You can query the VBR connection but cannot perform other operations.
//
//   - If the VBR connection is in the **Attached*	- state, the VBR connection is modified.
//
// @param request - UpdateTransitRouterVbrAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVbrAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVbrAttachmentAttributeWithContext(ctx context.Context, request *UpdateTransitRouterVbrAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVbrAttachmentAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the name and description of a virtual private cloud (VPC) connection on an Enterprise Edition transit router.
//
// Description:
//
// *UpdateTransitRouterVpcAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterVpcAttachments*	- operation to query the status of a VPC connection.
//
//   - If a VPC connection is in the **Modifying*	- state, the VPC connection is being modified. You can query the VPC connection but cannot perform other operations.
//
//   - If a VPC connection is in the **Attached*	- state, the VPC connection is modified.
//
// @param tmpReq - UpdateTransitRouterVpcAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVpcAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVpcAttachmentAttributeWithContext(ctx context.Context, tmpReq *UpdateTransitRouterVpcAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpcAttachmentAttributeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTransitRouterVpcAttachmentAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
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
// Modifies the zones and vSwitches that are associated with a virtual private cloud (VPC) connection.
//
// Description:
//
//	  When you add a zone and a vSwitch for a VPC connection, make sure that the vSwitch has at least one idle IP address. When you modify the zones and vSwitches of a VPC connection, the Enterprise Edition transit router creates an elastic network interface (ENI) in the vSwitch. The ENI occupies one IP address in the vSwitch. The ENI forwards traffic between the VPC and the Enterprise Edition transit router.
//
//		- **UpdateTransitRouterVpcAttachmentZones*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterVpcAttachments*	- operation to query the status of a VPC connection.
//
//	    	- If a VPC connection is in the **Modifying*	- state, the VPC connection is being modified. You can query the VPC connection but cannot perform other operations.
//
//	    	- If a VPC connection is in the **Attached*	- state, the VPC connection is modified.
//
// @param request - UpdateTransitRouterVpcAttachmentZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVpcAttachmentZonesResponse
func (client *Client) UpdateTransitRouterVpcAttachmentZonesWithContext(ctx context.Context, request *UpdateTransitRouterVpcAttachmentZonesRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpcAttachmentZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Modifies the configuration of a VPN attachment.
//
// @param request - UpdateTransitRouterVpnAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVpnAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVpnAttachmentAttributeWithContext(ctx context.Context, request *UpdateTransitRouterVpnAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpnAttachmentAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
// Withdraws the routes of a virtual private cloud (VPC) or a virtual border router (VBR) from a Cloud Enterprise Network (CEN) instance.
//
// @param request - WithdrawPublishedRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawPublishedRouteEntriesResponse
func (client *Client) WithdrawPublishedRouteEntriesWithContext(ctx context.Context, request *WithdrawPublishedRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *WithdrawPublishedRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
